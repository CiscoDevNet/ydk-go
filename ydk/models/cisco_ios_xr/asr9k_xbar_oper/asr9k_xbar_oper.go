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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Nodes.
    Nodes CrossBarStats_Nodes
}

func (crossBarStats *CrossBarStats) GetFilter() yfilter.YFilter { return crossBarStats.YFilter }

func (crossBarStats *CrossBarStats) SetFilter(yf yfilter.YFilter) { crossBarStats.YFilter = yf }

func (crossBarStats *CrossBarStats) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (crossBarStats *CrossBarStats) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-xbar-oper:cross-bar-stats"
}

func (crossBarStats *CrossBarStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &crossBarStats.Nodes
    }
    return nil
}

func (crossBarStats *CrossBarStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &crossBarStats.Nodes
    return children
}

func (crossBarStats *CrossBarStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crossBarStats *CrossBarStats) GetBundleName() string { return "cisco_ios_xr" }

func (crossBarStats *CrossBarStats) GetYangName() string { return "cross-bar-stats" }

func (crossBarStats *CrossBarStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crossBarStats *CrossBarStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crossBarStats *CrossBarStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crossBarStats *CrossBarStats) SetParent(parent types.Entity) { crossBarStats.parent = parent }

func (crossBarStats *CrossBarStats) GetParent() types.Entity { return crossBarStats.parent }

func (crossBarStats *CrossBarStats) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-xbar-oper" }

// CrossBarStats_Nodes
// Table of Nodes
type CrossBarStats_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // CrossBarStats_Nodes_Node.
    Node []CrossBarStats_Nodes_Node
}

func (nodes *CrossBarStats_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *CrossBarStats_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *CrossBarStats_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *CrossBarStats_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *CrossBarStats_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CrossBarStats_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *CrossBarStats_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *CrossBarStats_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *CrossBarStats_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *CrossBarStats_Nodes) GetYangName() string { return "nodes" }

func (nodes *CrossBarStats_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *CrossBarStats_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *CrossBarStats_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *CrossBarStats_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *CrossBarStats_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *CrossBarStats_Nodes) GetParentYangName() string { return "cross-bar-stats" }

// CrossBarStats_Nodes_Node
// Information about a particular node
type CrossBarStats_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of stats information.
    CrossBarTable CrossBarStats_Nodes_Node_CrossBarTable
}

func (node *CrossBarStats_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *CrossBarStats_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *CrossBarStats_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "cross-bar-table" { return "CrossBarTable" }
    return ""
}

func (node *CrossBarStats_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *CrossBarStats_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cross-bar-table" {
        return &node.CrossBarTable
    }
    return nil
}

func (node *CrossBarStats_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cross-bar-table"] = &node.CrossBarTable
    return children
}

func (node *CrossBarStats_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *CrossBarStats_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *CrossBarStats_Nodes_Node) GetYangName() string { return "node" }

func (node *CrossBarStats_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *CrossBarStats_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *CrossBarStats_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *CrossBarStats_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *CrossBarStats_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *CrossBarStats_Nodes_Node) GetParentYangName() string { return "nodes" }

// CrossBarStats_Nodes_Node_CrossBarTable
// Table of stats information
type CrossBarStats_Nodes_Node_CrossBarTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of packet stats.
    PktStats CrossBarStats_Nodes_Node_CrossBarTable_PktStats

    // Table of packet stats for SM15.
    Sm15Stats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetFilter() yfilter.YFilter { return crossBarTable.YFilter }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) SetFilter(yf yfilter.YFilter) { crossBarTable.YFilter = yf }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetGoName(yname string) string {
    if yname == "pkt-stats" { return "PktStats" }
    if yname == "sm15-stats" { return "Sm15Stats" }
    return ""
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetSegmentPath() string {
    return "cross-bar-table"
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pkt-stats" {
        return &crossBarTable.PktStats
    }
    if childYangName == "sm15-stats" {
        return &crossBarTable.Sm15Stats
    }
    return nil
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pkt-stats"] = &crossBarTable.PktStats
    children["sm15-stats"] = &crossBarTable.Sm15Stats
    return children
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetBundleName() string { return "cisco_ios_xr" }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetYangName() string { return "cross-bar-table" }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) SetParent(parent types.Entity) { crossBarTable.parent = parent }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetParent() types.Entity { return crossBarTable.parent }

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetParentYangName() string { return "node" }

// CrossBarStats_Nodes_Node_CrossBarTable_PktStats
// Table of packet stats
type CrossBarStats_Nodes_Node_CrossBarTable_PktStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats information for a particular asic type and port. The type is slice of
    // CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat.
    PktStat []CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetFilter() yfilter.YFilter { return pktStats.YFilter }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) SetFilter(yf yfilter.YFilter) { pktStats.YFilter = yf }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetGoName(yname string) string {
    if yname == "pkt-stat" { return "PktStat" }
    return ""
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetSegmentPath() string {
    return "pkt-stats"
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pkt-stat" {
        for _, c := range pktStats.PktStat {
            if pktStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat{}
        pktStats.PktStat = append(pktStats.PktStat, child)
        return &pktStats.PktStat[len(pktStats.PktStat)-1]
    }
    return nil
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pktStats.PktStat {
        children[pktStats.PktStat[i].GetSegmentPath()] = &pktStats.PktStat[i]
    }
    return children
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetBundleName() string { return "cisco_ios_xr" }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetYangName() string { return "pkt-stats" }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) SetParent(parent types.Entity) { pktStats.parent = parent }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetParent() types.Entity { return pktStats.parent }

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetParentYangName() string { return "cross-bar-table" }

// CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat
// Stats information for a particular asic type
// and port
type CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat struct {
    parent types.Entity
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

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetFilter() yfilter.YFilter { return pktStat.YFilter }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) SetFilter(yf yfilter.YFilter) { pktStat.YFilter = yf }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetGoName(yname string) string {
    if yname == "asic-id" { return "AsicId" }
    if yname == "port" { return "Port" }
    if yname == "internal-error-count" { return "InternalErrorCount" }
    if yname == "input-buffer-queued-packet-count-high" { return "InputBufferQueuedPacketCountHigh" }
    if yname == "ingress-packet-count-since-last-read-high" { return "IngressPacketCountSinceLastReadHigh" }
    if yname == "ingress-channel-utilization-count-high" { return "IngressChannelUtilizationCountHigh" }
    if yname == "input-buffer-back-pressure-count-high" { return "InputBufferBackPressureCountHigh" }
    if yname == "xbar-timeout-drop-count-high" { return "XbarTimeoutDropCountHigh" }
    if yname == "holdrop-count-high" { return "HoldropCountHigh" }
    if yname == "null-fpoe-drop-count-high" { return "NullFpoeDropCountHigh" }
    if yname == "diagnostic-packet-count-high" { return "DiagnosticPacketCountHigh" }
    if yname == "input-buffer-correctable-ecc-error-count-high" { return "InputBufferCorrectableEccErrorCountHigh" }
    if yname == "input-buffer-uncorrectable-ecc-error-count-high" { return "InputBufferUncorrectableEccErrorCountHigh" }
    if yname == "header-crc-error-count-high" { return "HeaderCrcErrorCountHigh" }
    if yname == "short-input-header-error-count-high" { return "ShortInputHeaderErrorCountHigh" }
    if yname == "packet-crc-error-count-high" { return "PacketCrcErrorCountHigh" }
    if yname == "short-packet-error-count-high" { return "ShortPacketErrorCountHigh" }
    if yname == "output-buffer-queued-packet-count-high" { return "OutputBufferQueuedPacketCountHigh" }
    if yname == "egress-packet-count-since-last-read-high" { return "EgressPacketCountSinceLastReadHigh" }
    if yname == "egress-channel-utilization-count-high" { return "EgressChannelUtilizationCountHigh" }
    if yname == "output-buffer-back-pressure-count-high" { return "OutputBufferBackPressureCountHigh" }
    if yname == "output-buffer-correctable-ecc-error-count-high" { return "OutputBufferCorrectableEccErrorCountHigh" }
    if yname == "output-buffer-uncorrectable-ecc-error-count-high" { return "OutputBufferUncorrectableEccErrorCountHigh" }
    if yname == "fpoedb-correctable-ecc-error-count-high" { return "FpoedbCorrectableEccErrorCountHigh" }
    if yname == "fpoedb-uncorrectable-ecc-error-count-high" { return "FpoedbUncorrectableEccErrorCountHigh" }
    if yname == "input-buffer-queued-packet-count-low" { return "InputBufferQueuedPacketCountLow" }
    if yname == "ingress-packet-count-since-last-read-low" { return "IngressPacketCountSinceLastReadLow" }
    if yname == "ingress-channel-utilization-count-low" { return "IngressChannelUtilizationCountLow" }
    if yname == "input-buffer-back-pressure-count-low" { return "InputBufferBackPressureCountLow" }
    if yname == "xbar-timeout-drop-count-low" { return "XbarTimeoutDropCountLow" }
    if yname == "holdrop-count-low" { return "HoldropCountLow" }
    if yname == "null-fpoe-drop-count-low" { return "NullFpoeDropCountLow" }
    if yname == "diagnostic-packet-count-low" { return "DiagnosticPacketCountLow" }
    if yname == "input-buffer-correctable-ecc-error-count-low" { return "InputBufferCorrectableEccErrorCountLow" }
    if yname == "input-buffer-uncorrectable-ecc-error-count-low" { return "InputBufferUncorrectableEccErrorCountLow" }
    if yname == "header-crc-error-count-low" { return "HeaderCrcErrorCountLow" }
    if yname == "short-input-header-error-count-low" { return "ShortInputHeaderErrorCountLow" }
    if yname == "packet-crc-error-count-low" { return "PacketCrcErrorCountLow" }
    if yname == "short-packet-error-count-low" { return "ShortPacketErrorCountLow" }
    if yname == "output-buffer-queued-packet-count-low" { return "OutputBufferQueuedPacketCountLow" }
    if yname == "egress-packet-count-since-last-read-low" { return "EgressPacketCountSinceLastReadLow" }
    if yname == "egress-channel-utilization-count-low" { return "EgressChannelUtilizationCountLow" }
    if yname == "output-buffer-back-pressure-count-low" { return "OutputBufferBackPressureCountLow" }
    if yname == "output-buffer-correctable-ecc-error-count-low" { return "OutputBufferCorrectableEccErrorCountLow" }
    if yname == "output-buffer-uncorrectable-ecc-error-count-low" { return "OutputBufferUncorrectableEccErrorCountLow" }
    if yname == "fpoedb-correctable-ecc-error-count-low" { return "FpoedbCorrectableEccErrorCountLow" }
    if yname == "fpoedb-uncorrectable-ecc-error-count-low" { return "FpoedbUncorrectableEccErrorCountLow" }
    return ""
}

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetSegmentPath() string {
    return "pkt-stat"
}

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asic-id"] = pktStat.AsicId
    leafs["port"] = pktStat.Port
    leafs["internal-error-count"] = pktStat.InternalErrorCount
    leafs["input-buffer-queued-packet-count-high"] = pktStat.InputBufferQueuedPacketCountHigh
    leafs["ingress-packet-count-since-last-read-high"] = pktStat.IngressPacketCountSinceLastReadHigh
    leafs["ingress-channel-utilization-count-high"] = pktStat.IngressChannelUtilizationCountHigh
    leafs["input-buffer-back-pressure-count-high"] = pktStat.InputBufferBackPressureCountHigh
    leafs["xbar-timeout-drop-count-high"] = pktStat.XbarTimeoutDropCountHigh
    leafs["holdrop-count-high"] = pktStat.HoldropCountHigh
    leafs["null-fpoe-drop-count-high"] = pktStat.NullFpoeDropCountHigh
    leafs["diagnostic-packet-count-high"] = pktStat.DiagnosticPacketCountHigh
    leafs["input-buffer-correctable-ecc-error-count-high"] = pktStat.InputBufferCorrectableEccErrorCountHigh
    leafs["input-buffer-uncorrectable-ecc-error-count-high"] = pktStat.InputBufferUncorrectableEccErrorCountHigh
    leafs["header-crc-error-count-high"] = pktStat.HeaderCrcErrorCountHigh
    leafs["short-input-header-error-count-high"] = pktStat.ShortInputHeaderErrorCountHigh
    leafs["packet-crc-error-count-high"] = pktStat.PacketCrcErrorCountHigh
    leafs["short-packet-error-count-high"] = pktStat.ShortPacketErrorCountHigh
    leafs["output-buffer-queued-packet-count-high"] = pktStat.OutputBufferQueuedPacketCountHigh
    leafs["egress-packet-count-since-last-read-high"] = pktStat.EgressPacketCountSinceLastReadHigh
    leafs["egress-channel-utilization-count-high"] = pktStat.EgressChannelUtilizationCountHigh
    leafs["output-buffer-back-pressure-count-high"] = pktStat.OutputBufferBackPressureCountHigh
    leafs["output-buffer-correctable-ecc-error-count-high"] = pktStat.OutputBufferCorrectableEccErrorCountHigh
    leafs["output-buffer-uncorrectable-ecc-error-count-high"] = pktStat.OutputBufferUncorrectableEccErrorCountHigh
    leafs["fpoedb-correctable-ecc-error-count-high"] = pktStat.FpoedbCorrectableEccErrorCountHigh
    leafs["fpoedb-uncorrectable-ecc-error-count-high"] = pktStat.FpoedbUncorrectableEccErrorCountHigh
    leafs["input-buffer-queued-packet-count-low"] = pktStat.InputBufferQueuedPacketCountLow
    leafs["ingress-packet-count-since-last-read-low"] = pktStat.IngressPacketCountSinceLastReadLow
    leafs["ingress-channel-utilization-count-low"] = pktStat.IngressChannelUtilizationCountLow
    leafs["input-buffer-back-pressure-count-low"] = pktStat.InputBufferBackPressureCountLow
    leafs["xbar-timeout-drop-count-low"] = pktStat.XbarTimeoutDropCountLow
    leafs["holdrop-count-low"] = pktStat.HoldropCountLow
    leafs["null-fpoe-drop-count-low"] = pktStat.NullFpoeDropCountLow
    leafs["diagnostic-packet-count-low"] = pktStat.DiagnosticPacketCountLow
    leafs["input-buffer-correctable-ecc-error-count-low"] = pktStat.InputBufferCorrectableEccErrorCountLow
    leafs["input-buffer-uncorrectable-ecc-error-count-low"] = pktStat.InputBufferUncorrectableEccErrorCountLow
    leafs["header-crc-error-count-low"] = pktStat.HeaderCrcErrorCountLow
    leafs["short-input-header-error-count-low"] = pktStat.ShortInputHeaderErrorCountLow
    leafs["packet-crc-error-count-low"] = pktStat.PacketCrcErrorCountLow
    leafs["short-packet-error-count-low"] = pktStat.ShortPacketErrorCountLow
    leafs["output-buffer-queued-packet-count-low"] = pktStat.OutputBufferQueuedPacketCountLow
    leafs["egress-packet-count-since-last-read-low"] = pktStat.EgressPacketCountSinceLastReadLow
    leafs["egress-channel-utilization-count-low"] = pktStat.EgressChannelUtilizationCountLow
    leafs["output-buffer-back-pressure-count-low"] = pktStat.OutputBufferBackPressureCountLow
    leafs["output-buffer-correctable-ecc-error-count-low"] = pktStat.OutputBufferCorrectableEccErrorCountLow
    leafs["output-buffer-uncorrectable-ecc-error-count-low"] = pktStat.OutputBufferUncorrectableEccErrorCountLow
    leafs["fpoedb-correctable-ecc-error-count-low"] = pktStat.FpoedbCorrectableEccErrorCountLow
    leafs["fpoedb-uncorrectable-ecc-error-count-low"] = pktStat.FpoedbUncorrectableEccErrorCountLow
    return leafs
}

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetBundleName() string { return "cisco_ios_xr" }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetYangName() string { return "pkt-stat" }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) SetParent(parent types.Entity) { pktStat.parent = parent }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetParent() types.Entity { return pktStat.parent }

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetParentYangName() string { return "pkt-stats" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats
// Table of packet stats for SM15
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats information for a particular asic type and port. The type is slice of
    // CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat.
    Sm15Stat []CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetFilter() yfilter.YFilter { return sm15Stats.YFilter }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) SetFilter(yf yfilter.YFilter) { sm15Stats.YFilter = yf }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetGoName(yname string) string {
    if yname == "sm15-stat" { return "Sm15Stat" }
    return ""
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetSegmentPath() string {
    return "sm15-stats"
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sm15-stat" {
        for _, c := range sm15Stats.Sm15Stat {
            if sm15Stats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat{}
        sm15Stats.Sm15Stat = append(sm15Stats.Sm15Stat, child)
        return &sm15Stats.Sm15Stat[len(sm15Stats.Sm15Stat)-1]
    }
    return nil
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sm15Stats.Sm15Stat {
        children[sm15Stats.Sm15Stat[i].GetSegmentPath()] = &sm15Stats.Sm15Stat[i]
    }
    return children
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetBundleName() string { return "cisco_ios_xr" }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetYangName() string { return "sm15-stats" }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) SetParent(parent types.Entity) { sm15Stats.parent = parent }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetParent() types.Entity { return sm15Stats.parent }

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetParentYangName() string { return "cross-bar-table" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat
// Stats information for a particular asic type
// and port
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat struct {
    parent types.Entity
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

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetFilter() yfilter.YFilter { return sm15Stat.YFilter }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) SetFilter(yf yfilter.YFilter) { sm15Stat.YFilter = yf }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetGoName(yname string) string {
    if yname == "asic-id" { return "AsicId" }
    if yname == "port" { return "Port" }
    if yname == "internal-err-cnt" { return "InternalErrCnt" }
    if yname == "ua0-stats" { return "Ua0Stats" }
    if yname == "ua1-stats" { return "Ua1Stats" }
    if yname == "ua2-stats" { return "Ua2Stats" }
    if yname == "ma-stats" { return "MaStats" }
    if yname == "ca-stats" { return "CaStats" }
    if yname == "pi-stats" { return "PiStats" }
    if yname == "pe-stats" { return "PeStats" }
    if yname == "pi-uc-stats" { return "PiUcStats" }
    if yname == "pi-mc-stats" { return "PiMcStats" }
    if yname == "pi-cc-stats" { return "PiCcStats" }
    if yname == "pe-uc-stats" { return "PeUcStats" }
    if yname == "pe-mc-stats" { return "PeMcStats" }
    if yname == "pe-cc-stats" { return "PeCcStats" }
    return ""
}

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetSegmentPath() string {
    return "sm15-stat"
}

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ua0-stats" {
        return &sm15Stat.Ua0Stats
    }
    if childYangName == "ua1-stats" {
        return &sm15Stat.Ua1Stats
    }
    if childYangName == "ua2-stats" {
        return &sm15Stat.Ua2Stats
    }
    if childYangName == "ma-stats" {
        return &sm15Stat.MaStats
    }
    if childYangName == "ca-stats" {
        return &sm15Stat.CaStats
    }
    if childYangName == "pi-stats" {
        return &sm15Stat.PiStats
    }
    if childYangName == "pe-stats" {
        return &sm15Stat.PeStats
    }
    if childYangName == "pi-uc-stats" {
        return &sm15Stat.PiUcStats
    }
    if childYangName == "pi-mc-stats" {
        return &sm15Stat.PiMcStats
    }
    if childYangName == "pi-cc-stats" {
        return &sm15Stat.PiCcStats
    }
    if childYangName == "pe-uc-stats" {
        return &sm15Stat.PeUcStats
    }
    if childYangName == "pe-mc-stats" {
        return &sm15Stat.PeMcStats
    }
    if childYangName == "pe-cc-stats" {
        return &sm15Stat.PeCcStats
    }
    return nil
}

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ua0-stats"] = &sm15Stat.Ua0Stats
    children["ua1-stats"] = &sm15Stat.Ua1Stats
    children["ua2-stats"] = &sm15Stat.Ua2Stats
    children["ma-stats"] = &sm15Stat.MaStats
    children["ca-stats"] = &sm15Stat.CaStats
    children["pi-stats"] = &sm15Stat.PiStats
    children["pe-stats"] = &sm15Stat.PeStats
    children["pi-uc-stats"] = &sm15Stat.PiUcStats
    children["pi-mc-stats"] = &sm15Stat.PiMcStats
    children["pi-cc-stats"] = &sm15Stat.PiCcStats
    children["pe-uc-stats"] = &sm15Stat.PeUcStats
    children["pe-mc-stats"] = &sm15Stat.PeMcStats
    children["pe-cc-stats"] = &sm15Stat.PeCcStats
    return children
}

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asic-id"] = sm15Stat.AsicId
    leafs["port"] = sm15Stat.Port
    leafs["internal-err-cnt"] = sm15Stat.InternalErrCnt
    return leafs
}

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetBundleName() string { return "cisco_ios_xr" }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetYangName() string { return "sm15-stat" }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) SetParent(parent types.Entity) { sm15Stat.parent = parent }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetParent() types.Entity { return sm15Stat.parent }

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetParentYangName() string { return "sm15-stats" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats
// ua0 stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats struct {
    parent types.Entity
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

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetFilter() yfilter.YFilter { return ua0Stats.YFilter }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) SetFilter(yf yfilter.YFilter) { ua0Stats.YFilter = yf }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetGoName(yname string) string {
    if yname == "dest-drop-pkt-cnt" { return "DestDropPktCnt" }
    if yname == "src-dest-pkt-cnt" { return "SrcDestPktCnt" }
    if yname == "dest-src-pkt-cnt" { return "DestSrcPktCnt" }
    if yname == "rcv-pkt-cnt" { return "RcvPktCnt" }
    if yname == "tx-pkt-cnt" { return "TxPktCnt" }
    if yname == "rx-drop-pkt-cnt" { return "RxDropPktCnt" }
    if yname == "rx-fabric-to-cnt" { return "RxFabricToCnt" }
    if yname == "ack-wait-cnt" { return "AckWaitCnt" }
    return ""
}

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetSegmentPath() string {
    return "ua0-stats"
}

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dest-drop-pkt-cnt"] = ua0Stats.DestDropPktCnt
    leafs["src-dest-pkt-cnt"] = ua0Stats.SrcDestPktCnt
    leafs["dest-src-pkt-cnt"] = ua0Stats.DestSrcPktCnt
    leafs["rcv-pkt-cnt"] = ua0Stats.RcvPktCnt
    leafs["tx-pkt-cnt"] = ua0Stats.TxPktCnt
    leafs["rx-drop-pkt-cnt"] = ua0Stats.RxDropPktCnt
    leafs["rx-fabric-to-cnt"] = ua0Stats.RxFabricToCnt
    leafs["ack-wait-cnt"] = ua0Stats.AckWaitCnt
    return leafs
}

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetBundleName() string { return "cisco_ios_xr" }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetYangName() string { return "ua0-stats" }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) SetParent(parent types.Entity) { ua0Stats.parent = parent }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetParent() types.Entity { return ua0Stats.parent }

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats
// ua1 stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats struct {
    parent types.Entity
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

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetFilter() yfilter.YFilter { return ua1Stats.YFilter }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) SetFilter(yf yfilter.YFilter) { ua1Stats.YFilter = yf }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetGoName(yname string) string {
    if yname == "dest-drop-pkt-cnt" { return "DestDropPktCnt" }
    if yname == "src-dest-pkt-cnt" { return "SrcDestPktCnt" }
    if yname == "dest-src-pkt-cnt" { return "DestSrcPktCnt" }
    if yname == "rcv-pkt-cnt" { return "RcvPktCnt" }
    if yname == "tx-pkt-cnt" { return "TxPktCnt" }
    if yname == "rx-drop-pkt-cnt" { return "RxDropPktCnt" }
    if yname == "rx-fabric-to-cnt" { return "RxFabricToCnt" }
    if yname == "ack-wait-cnt" { return "AckWaitCnt" }
    return ""
}

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetSegmentPath() string {
    return "ua1-stats"
}

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dest-drop-pkt-cnt"] = ua1Stats.DestDropPktCnt
    leafs["src-dest-pkt-cnt"] = ua1Stats.SrcDestPktCnt
    leafs["dest-src-pkt-cnt"] = ua1Stats.DestSrcPktCnt
    leafs["rcv-pkt-cnt"] = ua1Stats.RcvPktCnt
    leafs["tx-pkt-cnt"] = ua1Stats.TxPktCnt
    leafs["rx-drop-pkt-cnt"] = ua1Stats.RxDropPktCnt
    leafs["rx-fabric-to-cnt"] = ua1Stats.RxFabricToCnt
    leafs["ack-wait-cnt"] = ua1Stats.AckWaitCnt
    return leafs
}

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetBundleName() string { return "cisco_ios_xr" }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetYangName() string { return "ua1-stats" }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) SetParent(parent types.Entity) { ua1Stats.parent = parent }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetParent() types.Entity { return ua1Stats.parent }

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats
// ua2 stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats struct {
    parent types.Entity
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

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetFilter() yfilter.YFilter { return ua2Stats.YFilter }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) SetFilter(yf yfilter.YFilter) { ua2Stats.YFilter = yf }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetGoName(yname string) string {
    if yname == "dest-drop-pkt-cnt" { return "DestDropPktCnt" }
    if yname == "src-dest-pkt-cnt" { return "SrcDestPktCnt" }
    if yname == "dest-src-pkt-cnt" { return "DestSrcPktCnt" }
    if yname == "rcv-pkt-cnt" { return "RcvPktCnt" }
    if yname == "tx-pkt-cnt" { return "TxPktCnt" }
    if yname == "rx-drop-pkt-cnt" { return "RxDropPktCnt" }
    if yname == "rx-fabric-to-cnt" { return "RxFabricToCnt" }
    if yname == "ack-wait-cnt" { return "AckWaitCnt" }
    return ""
}

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetSegmentPath() string {
    return "ua2-stats"
}

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dest-drop-pkt-cnt"] = ua2Stats.DestDropPktCnt
    leafs["src-dest-pkt-cnt"] = ua2Stats.SrcDestPktCnt
    leafs["dest-src-pkt-cnt"] = ua2Stats.DestSrcPktCnt
    leafs["rcv-pkt-cnt"] = ua2Stats.RcvPktCnt
    leafs["tx-pkt-cnt"] = ua2Stats.TxPktCnt
    leafs["rx-drop-pkt-cnt"] = ua2Stats.RxDropPktCnt
    leafs["rx-fabric-to-cnt"] = ua2Stats.RxFabricToCnt
    leafs["ack-wait-cnt"] = ua2Stats.AckWaitCnt
    return leafs
}

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetBundleName() string { return "cisco_ios_xr" }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetYangName() string { return "ua2-stats" }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) SetParent(parent types.Entity) { ua2Stats.parent = parent }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetParent() types.Entity { return ua2Stats.parent }

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats
// ma stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats struct {
    parent types.Entity
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

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetFilter() yfilter.YFilter { return maStats.YFilter }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) SetFilter(yf yfilter.YFilter) { maStats.YFilter = yf }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetGoName(yname string) string {
    if yname == "dest-drop-pkt-cnt" { return "DestDropPktCnt" }
    if yname == "src-dest-pkt-cnt" { return "SrcDestPktCnt" }
    if yname == "dest-src-pkt-cnt" { return "DestSrcPktCnt" }
    if yname == "rcv-pkt-cnt" { return "RcvPktCnt" }
    if yname == "tx-pkt-cnt" { return "TxPktCnt" }
    if yname == "rx-drop-pkt-cnt" { return "RxDropPktCnt" }
    if yname == "rx-re-transmit-cnt" { return "RxReTransmitCnt" }
    if yname == "rx-fabric-to-cnt" { return "RxFabricToCnt" }
    if yname == "rx-hol-to-cnt" { return "RxHolToCnt" }
    return ""
}

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetSegmentPath() string {
    return "ma-stats"
}

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dest-drop-pkt-cnt"] = maStats.DestDropPktCnt
    leafs["src-dest-pkt-cnt"] = maStats.SrcDestPktCnt
    leafs["dest-src-pkt-cnt"] = maStats.DestSrcPktCnt
    leafs["rcv-pkt-cnt"] = maStats.RcvPktCnt
    leafs["tx-pkt-cnt"] = maStats.TxPktCnt
    leafs["rx-drop-pkt-cnt"] = maStats.RxDropPktCnt
    leafs["rx-re-transmit-cnt"] = maStats.RxReTransmitCnt
    leafs["rx-fabric-to-cnt"] = maStats.RxFabricToCnt
    leafs["rx-hol-to-cnt"] = maStats.RxHolToCnt
    return leafs
}

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetBundleName() string { return "cisco_ios_xr" }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetYangName() string { return "ma-stats" }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) SetParent(parent types.Entity) { maStats.parent = parent }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetParent() types.Entity { return maStats.parent }

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats
// ca stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats struct {
    parent types.Entity
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

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetFilter() yfilter.YFilter { return caStats.YFilter }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) SetFilter(yf yfilter.YFilter) { caStats.YFilter = yf }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetGoName(yname string) string {
    if yname == "dest-drop-pkt-cnt" { return "DestDropPktCnt" }
    if yname == "src-dest-pkt-cnt" { return "SrcDestPktCnt" }
    if yname == "dest-src-pkt-cnt" { return "DestSrcPktCnt" }
    if yname == "rcv-pkt-cnt" { return "RcvPktCnt" }
    if yname == "tx-pkt-cnt" { return "TxPktCnt" }
    if yname == "rx-drop-pkt-cnt" { return "RxDropPktCnt" }
    return ""
}

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetSegmentPath() string {
    return "ca-stats"
}

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dest-drop-pkt-cnt"] = caStats.DestDropPktCnt
    leafs["src-dest-pkt-cnt"] = caStats.SrcDestPktCnt
    leafs["dest-src-pkt-cnt"] = caStats.DestSrcPktCnt
    leafs["rcv-pkt-cnt"] = caStats.RcvPktCnt
    leafs["tx-pkt-cnt"] = caStats.TxPktCnt
    leafs["rx-drop-pkt-cnt"] = caStats.RxDropPktCnt
    return leafs
}

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetBundleName() string { return "cisco_ios_xr" }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetYangName() string { return "ca-stats" }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) SetParent(parent types.Entity) { caStats.parent = parent }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetParent() types.Entity { return caStats.parent }

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats
// pi stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats struct {
    parent types.Entity
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

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetFilter() yfilter.YFilter { return piStats.YFilter }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) SetFilter(yf yfilter.YFilter) { piStats.YFilter = yf }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetGoName(yname string) string {
    if yname == "total-rate1-cnt" { return "TotalRate1Cnt" }
    if yname == "total-rate2-cnt" { return "TotalRate2Cnt" }
    if yname == "total-rate3-cnt" { return "TotalRate3Cnt" }
    if yname == "total-calc-rate" { return "TotalCalcRate" }
    return ""
}

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetSegmentPath() string {
    return "pi-stats"
}

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-rate1-cnt"] = piStats.TotalRate1Cnt
    leafs["total-rate2-cnt"] = piStats.TotalRate2Cnt
    leafs["total-rate3-cnt"] = piStats.TotalRate3Cnt
    leafs["total-calc-rate"] = piStats.TotalCalcRate
    return leafs
}

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetBundleName() string { return "cisco_ios_xr" }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetYangName() string { return "pi-stats" }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) SetParent(parent types.Entity) { piStats.parent = parent }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetParent() types.Entity { return piStats.parent }

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats
// pe stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats struct {
    parent types.Entity
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
    Mc2UcPreemptCnt interface{}
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetFilter() yfilter.YFilter { return peStats.YFilter }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) SetFilter(yf yfilter.YFilter) { peStats.YFilter = yf }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetGoName(yname string) string {
    if yname == "total-rate1-cnt" { return "TotalRate1Cnt" }
    if yname == "total-rate2-cnt" { return "TotalRate2Cnt" }
    if yname == "total-rate3-cnt" { return "TotalRate3Cnt" }
    if yname == "total-calc-rate" { return "TotalCalcRate" }
    if yname == "mc2uc-preempt-cnt" { return "Mc2UcPreemptCnt" }
    return ""
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetSegmentPath() string {
    return "pe-stats"
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-rate1-cnt"] = peStats.TotalRate1Cnt
    leafs["total-rate2-cnt"] = peStats.TotalRate2Cnt
    leafs["total-rate3-cnt"] = peStats.TotalRate3Cnt
    leafs["total-calc-rate"] = peStats.TotalCalcRate
    leafs["mc2uc-preempt-cnt"] = peStats.Mc2UcPreemptCnt
    return leafs
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetBundleName() string { return "cisco_ios_xr" }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetYangName() string { return "pe-stats" }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) SetParent(parent types.Entity) { peStats.parent = parent }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetParent() types.Entity { return peStats.parent }

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats
// pi uc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats struct {
    parent types.Entity
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
    Uc0DataMemEcc1BitErrCnt interface{}

    // UC1 DATA MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc1DataMemEcc1BitErrCnt interface{}

    // UC2 DATA MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc2DataMemEcc1BitErrCnt interface{}

    // UC0 DATA MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc0DataMemEcc2BitErrCnt interface{}

    // UC1 DATA MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc1DataMemEcc2BitErrCnt interface{}

    // UC2 DATA MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc2DataMemEcc2BitErrCnt interface{}

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
    FpoeMemEcc1BitErrCnt interface{}

    // FPOE MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEcc2BitErrCnt interface{}

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

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetFilter() yfilter.YFilter { return piUcStats.YFilter }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) SetFilter(yf yfilter.YFilter) { piUcStats.YFilter = yf }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetGoName(yname string) string {
    if yname == "pkt-rcv-cnt" { return "PktRcvCnt" }
    if yname == "pkt-seq-err-cnt" { return "PktSeqErrCnt" }
    if yname == "in-coming-pkt-err-cnt" { return "InComingPktErrCnt" }
    if yname == "min-pkt-len-err-cnt" { return "MinPktLenErrCnt" }
    if yname == "max-pkt-len-err-cnt" { return "MaxPktLenErrCnt" }
    if yname == "line-err-drp-pkt" { return "LineErrDrpPkt" }
    if yname == "pkt-crc-err-cnt" { return "PktCrcErrCnt" }
    if yname == "pkt-cfh-crc-err-cnt" { return "PktCfhCrcErrCnt" }
    if yname == "line-s-written-in-mem0" { return "LineSWrittenInMem0" }
    if yname == "line-s-written-in-mem1" { return "LineSWrittenInMem1" }
    if yname == "line-s-written-in-mem2" { return "LineSWrittenInMem2" }
    if yname == "tail-drp-pkt-cnt" { return "TailDrpPktCnt" }
    if yname == "uc0-data-mem-ecc-1bit-err-cnt" { return "Uc0DataMemEcc1BitErrCnt" }
    if yname == "uc1-data-mem-ecc-1bit-err-cnt" { return "Uc1DataMemEcc1BitErrCnt" }
    if yname == "uc2-data-mem-ecc-1bit-err-cnt" { return "Uc2DataMemEcc1BitErrCnt" }
    if yname == "uc0-data-mem-ecc-2bit-err-cnt" { return "Uc0DataMemEcc2BitErrCnt" }
    if yname == "uc1-data-mem-ecc-2bit-err-cnt" { return "Uc1DataMemEcc2BitErrCnt" }
    if yname == "uc2-data-mem-ecc-2bit-err-cnt" { return "Uc2DataMemEcc2BitErrCnt" }
    if yname == "diag-pkt-cnt" { return "DiagPktCnt" }
    if yname == "pkt-sent-to-disabled-port-cnt" { return "PktSentToDisabledPortCnt" }
    if yname == "pkt-null-poe-sent-ua0-cnt" { return "PktNullPoeSentUa0Cnt" }
    if yname == "pkt-null-poe-sent-ua1-cnt" { return "PktNullPoeSentUa1Cnt" }
    if yname == "pkt-null-poe-sent-ua2-cnt" { return "PktNullPoeSentUa2Cnt" }
    if yname == "pkt-fpoe-addr-rng-hit-cnt" { return "PktFpoeAddrRngHitCnt" }
    if yname == "fpoe-mem-ecc-1bit-err-cnt" { return "FpoeMemEcc1BitErrCnt" }
    if yname == "fpoe-mem-ecc-2bit-err-cnt" { return "FpoeMemEcc2BitErrCnt" }
    if yname == "pkts-sent-to-ux0-cnt" { return "PktsSentToUx0Cnt" }
    if yname == "pkts-sent-to-ux1-cnt" { return "PktsSentToUx1Cnt" }
    if yname == "pkts-sent-to-ux2-cnt" { return "PktsSentToUx2Cnt" }
    if yname == "cpp-head-drop-pkt-cnt" { return "CppHeadDropPktCnt" }
    if yname == "tr-head-drop-pkt-cnt" { return "TrHeadDropPktCnt" }
    if yname == "tr-pkt-sent-to-ux" { return "TrPktSentToUx" }
    if yname == "stop-thrsh-hit-cnt" { return "StopThrshHitCnt" }
    if yname == "rate-cnt" { return "RateCnt" }
    if yname == "calc-rate" { return "CalcRate" }
    if yname == "crc-stomp-pkt-cnt" { return "CrcStompPktCnt" }
    return ""
}

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetSegmentPath() string {
    return "pi-uc-stats"
}

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pkt-rcv-cnt"] = piUcStats.PktRcvCnt
    leafs["pkt-seq-err-cnt"] = piUcStats.PktSeqErrCnt
    leafs["in-coming-pkt-err-cnt"] = piUcStats.InComingPktErrCnt
    leafs["min-pkt-len-err-cnt"] = piUcStats.MinPktLenErrCnt
    leafs["max-pkt-len-err-cnt"] = piUcStats.MaxPktLenErrCnt
    leafs["line-err-drp-pkt"] = piUcStats.LineErrDrpPkt
    leafs["pkt-crc-err-cnt"] = piUcStats.PktCrcErrCnt
    leafs["pkt-cfh-crc-err-cnt"] = piUcStats.PktCfhCrcErrCnt
    leafs["line-s-written-in-mem0"] = piUcStats.LineSWrittenInMem0
    leafs["line-s-written-in-mem1"] = piUcStats.LineSWrittenInMem1
    leafs["line-s-written-in-mem2"] = piUcStats.LineSWrittenInMem2
    leafs["tail-drp-pkt-cnt"] = piUcStats.TailDrpPktCnt
    leafs["uc0-data-mem-ecc-1bit-err-cnt"] = piUcStats.Uc0DataMemEcc1BitErrCnt
    leafs["uc1-data-mem-ecc-1bit-err-cnt"] = piUcStats.Uc1DataMemEcc1BitErrCnt
    leafs["uc2-data-mem-ecc-1bit-err-cnt"] = piUcStats.Uc2DataMemEcc1BitErrCnt
    leafs["uc0-data-mem-ecc-2bit-err-cnt"] = piUcStats.Uc0DataMemEcc2BitErrCnt
    leafs["uc1-data-mem-ecc-2bit-err-cnt"] = piUcStats.Uc1DataMemEcc2BitErrCnt
    leafs["uc2-data-mem-ecc-2bit-err-cnt"] = piUcStats.Uc2DataMemEcc2BitErrCnt
    leafs["diag-pkt-cnt"] = piUcStats.DiagPktCnt
    leafs["pkt-sent-to-disabled-port-cnt"] = piUcStats.PktSentToDisabledPortCnt
    leafs["pkt-null-poe-sent-ua0-cnt"] = piUcStats.PktNullPoeSentUa0Cnt
    leafs["pkt-null-poe-sent-ua1-cnt"] = piUcStats.PktNullPoeSentUa1Cnt
    leafs["pkt-null-poe-sent-ua2-cnt"] = piUcStats.PktNullPoeSentUa2Cnt
    leafs["pkt-fpoe-addr-rng-hit-cnt"] = piUcStats.PktFpoeAddrRngHitCnt
    leafs["fpoe-mem-ecc-1bit-err-cnt"] = piUcStats.FpoeMemEcc1BitErrCnt
    leafs["fpoe-mem-ecc-2bit-err-cnt"] = piUcStats.FpoeMemEcc2BitErrCnt
    leafs["pkts-sent-to-ux0-cnt"] = piUcStats.PktsSentToUx0Cnt
    leafs["pkts-sent-to-ux1-cnt"] = piUcStats.PktsSentToUx1Cnt
    leafs["pkts-sent-to-ux2-cnt"] = piUcStats.PktsSentToUx2Cnt
    leafs["cpp-head-drop-pkt-cnt"] = piUcStats.CppHeadDropPktCnt
    leafs["tr-head-drop-pkt-cnt"] = piUcStats.TrHeadDropPktCnt
    leafs["tr-pkt-sent-to-ux"] = piUcStats.TrPktSentToUx
    leafs["stop-thrsh-hit-cnt"] = piUcStats.StopThrshHitCnt
    leafs["rate-cnt"] = piUcStats.RateCnt
    leafs["calc-rate"] = piUcStats.CalcRate
    leafs["crc-stomp-pkt-cnt"] = piUcStats.CrcStompPktCnt
    return leafs
}

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetBundleName() string { return "cisco_ios_xr" }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetYangName() string { return "pi-uc-stats" }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) SetParent(parent types.Entity) { piUcStats.parent = parent }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetParent() types.Entity { return piUcStats.parent }

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats
// pi mc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats struct {
    parent types.Entity
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
    DataMem0Ecc1BitErrCnt interface{}

    // DATA MEM1 ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem1Ecc1BitErrCnt interface{}

    // DATA MEM2 ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem2Ecc1BitErrCnt interface{}

    // DATA MEM0 ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem0Ecc2BitErrCnt interface{}

    // DATA MEM1 ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem1Ecc2BitErrCnt interface{}

    // DATA MEM2 ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem2Ecc2BitErrCnt interface{}

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
    FpoeMemEcc1BitErrCnt interface{}

    // FPOE MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEcc2BitErrCnt interface{}

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

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetFilter() yfilter.YFilter { return piMcStats.YFilter }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) SetFilter(yf yfilter.YFilter) { piMcStats.YFilter = yf }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetGoName(yname string) string {
    if yname == "pkt-rcv-cnt" { return "PktRcvCnt" }
    if yname == "pkt-seq-err-cnt" { return "PktSeqErrCnt" }
    if yname == "in-coming-pkt-err-cnt" { return "InComingPktErrCnt" }
    if yname == "min-pkt-len-err-cnt" { return "MinPktLenErrCnt" }
    if yname == "max-pkt-len-err-cnt" { return "MaxPktLenErrCnt" }
    if yname == "line-err-drp-pkt" { return "LineErrDrpPkt" }
    if yname == "pkt-crc-err-cnt" { return "PktCrcErrCnt" }
    if yname == "pkt-cfh-crc-err-cnt" { return "PktCfhCrcErrCnt" }
    if yname == "line-s-written-in-mem" { return "LineSWrittenInMem" }
    if yname == "tail-drp-pkt-cnt" { return "TailDrpPktCnt" }
    if yname == "data-mem0-ecc-1bit-err-cnt" { return "DataMem0Ecc1BitErrCnt" }
    if yname == "data-mem1-ecc-1bit-err-cnt" { return "DataMem1Ecc1BitErrCnt" }
    if yname == "data-mem2-ecc-1bit-err-cnt" { return "DataMem2Ecc1BitErrCnt" }
    if yname == "data-mem0-ecc-2bit-err-cnt" { return "DataMem0Ecc2BitErrCnt" }
    if yname == "data-mem1-ecc-2bit-err-cnt" { return "DataMem1Ecc2BitErrCnt" }
    if yname == "data-mem2-ecc-2bit-err-cnt" { return "DataMem2Ecc2BitErrCnt" }
    if yname == "diag-pkt-cnt" { return "DiagPktCnt" }
    if yname == "pkt-sent-to-disabled-port" { return "PktSentToDisabledPort" }
    if yname == "pkt-fpoe-match-hit-cnt" { return "PktFpoeMatchHitCnt" }
    if yname == "pkt-null-poe-sent-cnt" { return "PktNullPoeSentCnt" }
    if yname == "pkt-fpoe-addr-rng-hit-cnt" { return "PktFpoeAddrRngHitCnt" }
    if yname == "di-hdr-len-err-pkt-cnt" { return "DiHdrLenErrPktCnt" }
    if yname == "di-err-pkt-cnt" { return "DiErrPktCnt" }
    if yname == "fpoe-mem-ecc-1bit-err-cnt" { return "FpoeMemEcc1BitErrCnt" }
    if yname == "fpoe-mem-ecc-2bit-err-cnt" { return "FpoeMemEcc2BitErrCnt" }
    if yname == "pkts-sent-to-mx-cnt" { return "PktsSentToMxCnt" }
    if yname == "cpp-head-drop-pkt-from-ma-cnt" { return "CppHeadDropPktFromMaCnt" }
    if yname == "tr-head-drop-pkt-from-ma-cnt" { return "TrHeadDropPktFromMaCnt" }
    if yname == "tr-pkt-sent-to-mx" { return "TrPktSentToMx" }
    if yname == "stop-thrsh-hit-cnt" { return "StopThrshHitCnt" }
    if yname == "rate-cnt" { return "RateCnt" }
    if yname == "calc-rate" { return "CalcRate" }
    if yname == "crc-stomp-pkt-cnt" { return "CrcStompPktCnt" }
    return ""
}

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetSegmentPath() string {
    return "pi-mc-stats"
}

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pkt-rcv-cnt"] = piMcStats.PktRcvCnt
    leafs["pkt-seq-err-cnt"] = piMcStats.PktSeqErrCnt
    leafs["in-coming-pkt-err-cnt"] = piMcStats.InComingPktErrCnt
    leafs["min-pkt-len-err-cnt"] = piMcStats.MinPktLenErrCnt
    leafs["max-pkt-len-err-cnt"] = piMcStats.MaxPktLenErrCnt
    leafs["line-err-drp-pkt"] = piMcStats.LineErrDrpPkt
    leafs["pkt-crc-err-cnt"] = piMcStats.PktCrcErrCnt
    leafs["pkt-cfh-crc-err-cnt"] = piMcStats.PktCfhCrcErrCnt
    leafs["line-s-written-in-mem"] = piMcStats.LineSWrittenInMem
    leafs["tail-drp-pkt-cnt"] = piMcStats.TailDrpPktCnt
    leafs["data-mem0-ecc-1bit-err-cnt"] = piMcStats.DataMem0Ecc1BitErrCnt
    leafs["data-mem1-ecc-1bit-err-cnt"] = piMcStats.DataMem1Ecc1BitErrCnt
    leafs["data-mem2-ecc-1bit-err-cnt"] = piMcStats.DataMem2Ecc1BitErrCnt
    leafs["data-mem0-ecc-2bit-err-cnt"] = piMcStats.DataMem0Ecc2BitErrCnt
    leafs["data-mem1-ecc-2bit-err-cnt"] = piMcStats.DataMem1Ecc2BitErrCnt
    leafs["data-mem2-ecc-2bit-err-cnt"] = piMcStats.DataMem2Ecc2BitErrCnt
    leafs["diag-pkt-cnt"] = piMcStats.DiagPktCnt
    leafs["pkt-sent-to-disabled-port"] = piMcStats.PktSentToDisabledPort
    leafs["pkt-fpoe-match-hit-cnt"] = piMcStats.PktFpoeMatchHitCnt
    leafs["pkt-null-poe-sent-cnt"] = piMcStats.PktNullPoeSentCnt
    leafs["pkt-fpoe-addr-rng-hit-cnt"] = piMcStats.PktFpoeAddrRngHitCnt
    leafs["di-hdr-len-err-pkt-cnt"] = piMcStats.DiHdrLenErrPktCnt
    leafs["di-err-pkt-cnt"] = piMcStats.DiErrPktCnt
    leafs["fpoe-mem-ecc-1bit-err-cnt"] = piMcStats.FpoeMemEcc1BitErrCnt
    leafs["fpoe-mem-ecc-2bit-err-cnt"] = piMcStats.FpoeMemEcc2BitErrCnt
    leafs["pkts-sent-to-mx-cnt"] = piMcStats.PktsSentToMxCnt
    leafs["cpp-head-drop-pkt-from-ma-cnt"] = piMcStats.CppHeadDropPktFromMaCnt
    leafs["tr-head-drop-pkt-from-ma-cnt"] = piMcStats.TrHeadDropPktFromMaCnt
    leafs["tr-pkt-sent-to-mx"] = piMcStats.TrPktSentToMx
    leafs["stop-thrsh-hit-cnt"] = piMcStats.StopThrshHitCnt
    leafs["rate-cnt"] = piMcStats.RateCnt
    leafs["calc-rate"] = piMcStats.CalcRate
    leafs["crc-stomp-pkt-cnt"] = piMcStats.CrcStompPktCnt
    return leafs
}

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetBundleName() string { return "cisco_ios_xr" }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetYangName() string { return "pi-mc-stats" }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) SetParent(parent types.Entity) { piMcStats.parent = parent }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetParent() types.Entity { return piMcStats.parent }

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats
// pi cc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats struct {
    parent types.Entity
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

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetFilter() yfilter.YFilter { return piCcStats.YFilter }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) SetFilter(yf yfilter.YFilter) { piCcStats.YFilter = yf }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetGoName(yname string) string {
    if yname == "in0-ecc-serr-cnt" { return "In0EccSerrCnt" }
    if yname == "in0-ecc-derr-cnt" { return "In0EccDerrCnt" }
    if yname == "in1-ecc-serr-cnt" { return "In1EccSerrCnt" }
    if yname == "in1-ecc-derr-cnt" { return "In1EccDerrCnt" }
    if yname == "data-mem-ecc-serr-cnt" { return "DataMemEccSerrCnt" }
    if yname == "data-mem-ecc-derr-cnt" { return "DataMemEccDerrCnt" }
    if yname == "data-mem-ovf0-cnt" { return "DataMemOvf0Cnt" }
    if yname == "data-mem-ovf1-cnt" { return "DataMemOvf1Cnt" }
    if yname == "fpoe-mem-ecc-serr-cnt" { return "FpoeMemEccSerrCnt" }
    if yname == "fpoe-mem-ecc-derr-cnt" { return "FpoeMemEccDerrCnt" }
    if yname == "null-poe-cnt" { return "NullPoeCnt" }
    if yname == "shut-ack-cnt" { return "ShutAckCnt" }
    if yname == "in0-fnc-err-cnt" { return "In0FncErrCnt" }
    if yname == "in1-fnc-err-cnt" { return "In1FncErrCnt" }
    if yname == "in0-drop-cnt" { return "In0DropCnt" }
    if yname == "in1-drop-cnt" { return "In1DropCnt" }
    if yname == "in0-cong-cnt" { return "In0CongCnt" }
    if yname == "in1-cong-cnt" { return "In1CongCnt" }
    if yname == "in0-shut-cnt" { return "In0ShutCnt" }
    if yname == "in1-shut-cnt" { return "In1ShutCnt" }
    if yname == "tail-drop-msg-cnt" { return "TailDropMsgCnt" }
    if yname == "in0-pkt-cnt" { return "In0PktCnt" }
    if yname == "in1-pkt-cnt" { return "In1PktCnt" }
    if yname == "dmem-rd-cnt" { return "DmemRdCnt" }
    if yname == "in-dmem0-cnt" { return "InDmem0Cnt" }
    if yname == "in-dmem1-cnt" { return "InDmem1Cnt" }
    if yname == "out-pkt-cnt" { return "OutPktCnt" }
    if yname == "stop-thrsh-hit-cnt" { return "StopThrshHitCnt" }
    if yname == "rate-cnt" { return "RateCnt" }
    if yname == "calc-rate" { return "CalcRate" }
    return ""
}

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetSegmentPath() string {
    return "pi-cc-stats"
}

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in0-ecc-serr-cnt"] = piCcStats.In0EccSerrCnt
    leafs["in0-ecc-derr-cnt"] = piCcStats.In0EccDerrCnt
    leafs["in1-ecc-serr-cnt"] = piCcStats.In1EccSerrCnt
    leafs["in1-ecc-derr-cnt"] = piCcStats.In1EccDerrCnt
    leafs["data-mem-ecc-serr-cnt"] = piCcStats.DataMemEccSerrCnt
    leafs["data-mem-ecc-derr-cnt"] = piCcStats.DataMemEccDerrCnt
    leafs["data-mem-ovf0-cnt"] = piCcStats.DataMemOvf0Cnt
    leafs["data-mem-ovf1-cnt"] = piCcStats.DataMemOvf1Cnt
    leafs["fpoe-mem-ecc-serr-cnt"] = piCcStats.FpoeMemEccSerrCnt
    leafs["fpoe-mem-ecc-derr-cnt"] = piCcStats.FpoeMemEccDerrCnt
    leafs["null-poe-cnt"] = piCcStats.NullPoeCnt
    leafs["shut-ack-cnt"] = piCcStats.ShutAckCnt
    leafs["in0-fnc-err-cnt"] = piCcStats.In0FncErrCnt
    leafs["in1-fnc-err-cnt"] = piCcStats.In1FncErrCnt
    leafs["in0-drop-cnt"] = piCcStats.In0DropCnt
    leafs["in1-drop-cnt"] = piCcStats.In1DropCnt
    leafs["in0-cong-cnt"] = piCcStats.In0CongCnt
    leafs["in1-cong-cnt"] = piCcStats.In1CongCnt
    leafs["in0-shut-cnt"] = piCcStats.In0ShutCnt
    leafs["in1-shut-cnt"] = piCcStats.In1ShutCnt
    leafs["tail-drop-msg-cnt"] = piCcStats.TailDropMsgCnt
    leafs["in0-pkt-cnt"] = piCcStats.In0PktCnt
    leafs["in1-pkt-cnt"] = piCcStats.In1PktCnt
    leafs["dmem-rd-cnt"] = piCcStats.DmemRdCnt
    leafs["in-dmem0-cnt"] = piCcStats.InDmem0Cnt
    leafs["in-dmem1-cnt"] = piCcStats.InDmem1Cnt
    leafs["out-pkt-cnt"] = piCcStats.OutPktCnt
    leafs["stop-thrsh-hit-cnt"] = piCcStats.StopThrshHitCnt
    leafs["rate-cnt"] = piCcStats.RateCnt
    leafs["calc-rate"] = piCcStats.CalcRate
    return leafs
}

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetBundleName() string { return "cisco_ios_xr" }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetYangName() string { return "pi-cc-stats" }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) SetParent(parent types.Entity) { piCcStats.parent = parent }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetParent() types.Entity { return piCcStats.parent }

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats
// pe uc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats struct {
    parent types.Entity
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
    Ecc1BitErrUc00Cnt interface{}

    // ECC 1BIT ERR UC0 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrUc01Cnt interface{}

    // ECC 1BIT ERR UC1 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrUc10Cnt interface{}

    // ECC 1BIT ERR UC1 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrUc11Cnt interface{}

    // ECC 1BIT ERR UC2 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrUc20Cnt interface{}

    // ECC 1BIT ERR UC2 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrUc21Cnt interface{}

    // ECC 2BIT ERR UC0 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrUc00Cnt interface{}

    // ECC 2BIT ERR UC0 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrUc01Cnt interface{}

    // ECC 2BIT ERR UC1 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrUc10Cnt interface{}

    // ECC 2BIT ERR UC1 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrUc11Cnt interface{}

    // ECC 2BIT ERR UC2 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrUc20Cnt interface{}

    // ECC 2BIT ERR UC2 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrUc21Cnt interface{}

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

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetFilter() yfilter.YFilter { return peUcStats.YFilter }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) SetFilter(yf yfilter.YFilter) { peUcStats.YFilter = yf }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetGoName(yname string) string {
    if yname == "in-pkt-uc0-cnt" { return "InPktUc0Cnt" }
    if yname == "in-pkt-uc1-cnt" { return "InPktUc1Cnt" }
    if yname == "in-pkt-uc2-cnt" { return "InPktUc2Cnt" }
    if yname == "in-full-line-uc0-cnt" { return "InFullLineUc0Cnt" }
    if yname == "in-full-line-uc1-cnt" { return "InFullLineUc1Cnt" }
    if yname == "in-full-line-uc2-cnt" { return "InFullLineUc2Cnt" }
    if yname == "pkt-trunc-eop-uc0-cnt" { return "PktTruncEopUc0Cnt" }
    if yname == "pkt-trunc-eop-uc1-cnt" { return "PktTruncEopUc1Cnt" }
    if yname == "pkt-trunc-eop-uc2-cnt" { return "PktTruncEopUc2Cnt" }
    if yname == "pkt-sop-drop-uc0-cnt" { return "PktSopDropUc0Cnt" }
    if yname == "pkt-sop-drop-uc1-cnt" { return "PktSopDropUc1Cnt" }
    if yname == "pkt-sop-drop-uc2-cnt" { return "PktSopDropUc2Cnt" }
    if yname == "pkt-ecc-err-drop-uc-cnt" { return "PktEccErrDropUcCnt" }
    if yname == "pkt-ecc-trunc-cnt-uc-cnt" { return "PktEccTruncCntUcCnt" }
    if yname == "ecc-1bit-err-uc0-0-cnt" { return "Ecc1BitErrUc00Cnt" }
    if yname == "ecc-1bit-err-uc0-1-cnt" { return "Ecc1BitErrUc01Cnt" }
    if yname == "ecc-1bit-err-uc1-0-cnt" { return "Ecc1BitErrUc10Cnt" }
    if yname == "ecc-1bit-err-uc1-1-cnt" { return "Ecc1BitErrUc11Cnt" }
    if yname == "ecc-1bit-err-uc2-0-cnt" { return "Ecc1BitErrUc20Cnt" }
    if yname == "ecc-1bit-err-uc2-1-cnt" { return "Ecc1BitErrUc21Cnt" }
    if yname == "ecc-2bit-err-uc0-0-cnt" { return "Ecc2BitErrUc00Cnt" }
    if yname == "ecc-2bit-err-uc0-1-cnt" { return "Ecc2BitErrUc01Cnt" }
    if yname == "ecc-2bit-err-uc1-0-cnt" { return "Ecc2BitErrUc10Cnt" }
    if yname == "ecc-2bit-err-uc1-1-cnt" { return "Ecc2BitErrUc11Cnt" }
    if yname == "ecc-2bit-err-uc2-0-cnt" { return "Ecc2BitErrUc20Cnt" }
    if yname == "ecc-2bit-err-uc2-1-cnt" { return "Ecc2BitErrUc21Cnt" }
    if yname == "out-pkt-uc-cnt" { return "OutPktUcCnt" }
    if yname == "fe-uc-sop-eop-pack-cnt" { return "FeUcSopEopPackCnt" }
    if yname == "fc-uc-0-1-trans-cnt" { return "FcUc01TransCnt" }
    if yname == "rate-cnt" { return "RateCnt" }
    if yname == "calc-rate" { return "CalcRate" }
    return ""
}

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetSegmentPath() string {
    return "pe-uc-stats"
}

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkt-uc0-cnt"] = peUcStats.InPktUc0Cnt
    leafs["in-pkt-uc1-cnt"] = peUcStats.InPktUc1Cnt
    leafs["in-pkt-uc2-cnt"] = peUcStats.InPktUc2Cnt
    leafs["in-full-line-uc0-cnt"] = peUcStats.InFullLineUc0Cnt
    leafs["in-full-line-uc1-cnt"] = peUcStats.InFullLineUc1Cnt
    leafs["in-full-line-uc2-cnt"] = peUcStats.InFullLineUc2Cnt
    leafs["pkt-trunc-eop-uc0-cnt"] = peUcStats.PktTruncEopUc0Cnt
    leafs["pkt-trunc-eop-uc1-cnt"] = peUcStats.PktTruncEopUc1Cnt
    leafs["pkt-trunc-eop-uc2-cnt"] = peUcStats.PktTruncEopUc2Cnt
    leafs["pkt-sop-drop-uc0-cnt"] = peUcStats.PktSopDropUc0Cnt
    leafs["pkt-sop-drop-uc1-cnt"] = peUcStats.PktSopDropUc1Cnt
    leafs["pkt-sop-drop-uc2-cnt"] = peUcStats.PktSopDropUc2Cnt
    leafs["pkt-ecc-err-drop-uc-cnt"] = peUcStats.PktEccErrDropUcCnt
    leafs["pkt-ecc-trunc-cnt-uc-cnt"] = peUcStats.PktEccTruncCntUcCnt
    leafs["ecc-1bit-err-uc0-0-cnt"] = peUcStats.Ecc1BitErrUc00Cnt
    leafs["ecc-1bit-err-uc0-1-cnt"] = peUcStats.Ecc1BitErrUc01Cnt
    leafs["ecc-1bit-err-uc1-0-cnt"] = peUcStats.Ecc1BitErrUc10Cnt
    leafs["ecc-1bit-err-uc1-1-cnt"] = peUcStats.Ecc1BitErrUc11Cnt
    leafs["ecc-1bit-err-uc2-0-cnt"] = peUcStats.Ecc1BitErrUc20Cnt
    leafs["ecc-1bit-err-uc2-1-cnt"] = peUcStats.Ecc1BitErrUc21Cnt
    leafs["ecc-2bit-err-uc0-0-cnt"] = peUcStats.Ecc2BitErrUc00Cnt
    leafs["ecc-2bit-err-uc0-1-cnt"] = peUcStats.Ecc2BitErrUc01Cnt
    leafs["ecc-2bit-err-uc1-0-cnt"] = peUcStats.Ecc2BitErrUc10Cnt
    leafs["ecc-2bit-err-uc1-1-cnt"] = peUcStats.Ecc2BitErrUc11Cnt
    leafs["ecc-2bit-err-uc2-0-cnt"] = peUcStats.Ecc2BitErrUc20Cnt
    leafs["ecc-2bit-err-uc2-1-cnt"] = peUcStats.Ecc2BitErrUc21Cnt
    leafs["out-pkt-uc-cnt"] = peUcStats.OutPktUcCnt
    leafs["fe-uc-sop-eop-pack-cnt"] = peUcStats.FeUcSopEopPackCnt
    leafs["fc-uc-0-1-trans-cnt"] = peUcStats.FcUc01TransCnt
    leafs["rate-cnt"] = peUcStats.RateCnt
    leafs["calc-rate"] = peUcStats.CalcRate
    return leafs
}

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetBundleName() string { return "cisco_ios_xr" }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetYangName() string { return "pe-uc-stats" }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) SetParent(parent types.Entity) { peUcStats.parent = parent }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetParent() types.Entity { return peUcStats.parent }

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats
// pe mc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats struct {
    parent types.Entity
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
    Ecc1BitErrMc0Cnt interface{}

    // ECC 1BIT ERR MC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrMc1Cnt interface{}

    // ECC 1BIT ERR MC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1BitErrMc2Cnt interface{}

    // ECC 2BIT ERR MC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrMc0Cnt interface{}

    // ECC 2BIT ERR MC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrMc1Cnt interface{}

    // ECC 2BIT ERR MC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2BitErrMc2Cnt interface{}

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

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetFilter() yfilter.YFilter { return peMcStats.YFilter }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) SetFilter(yf yfilter.YFilter) { peMcStats.YFilter = yf }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetGoName(yname string) string {
    if yname == "in-pkt-mc-cnt" { return "InPktMcCnt" }
    if yname == "in-full-line-mc-cnt" { return "InFullLineMcCnt" }
    if yname == "pkt-trunc-eop-mc-cnt" { return "PktTruncEopMcCnt" }
    if yname == "pkt-sop-drop-mc-cnt" { return "PktSopDropMcCnt" }
    if yname == "pkt-ecc-err-drop-mc-cnt" { return "PktEccErrDropMcCnt" }
    if yname == "pkt-ecc-err-trunc-cnt-mc-cnt" { return "PktEccErrTruncCntMcCnt" }
    if yname == "ecc-1bit-err-mc0-cnt" { return "Ecc1BitErrMc0Cnt" }
    if yname == "ecc-1bit-err-mc1-cnt" { return "Ecc1BitErrMc1Cnt" }
    if yname == "ecc-1bit-err-mc2-cnt" { return "Ecc1BitErrMc2Cnt" }
    if yname == "ecc-2bit-err-mc0-cnt" { return "Ecc2BitErrMc0Cnt" }
    if yname == "ecc-2bit-err-mc1-cnt" { return "Ecc2BitErrMc1Cnt" }
    if yname == "ecc-2bit-err-mc2-cnt" { return "Ecc2BitErrMc2Cnt" }
    if yname == "out-pkt-mc-cnt" { return "OutPktMcCnt" }
    if yname == "fe-mc-sop-eop-pack-cnt" { return "FeMcSopEopPackCnt" }
    if yname == "fc-mc-0-1-trans-cnt" { return "FcMc01TransCnt" }
    if yname == "rate-cnt" { return "RateCnt" }
    if yname == "calc-rate" { return "CalcRate" }
    return ""
}

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetSegmentPath() string {
    return "pe-mc-stats"
}

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkt-mc-cnt"] = peMcStats.InPktMcCnt
    leafs["in-full-line-mc-cnt"] = peMcStats.InFullLineMcCnt
    leafs["pkt-trunc-eop-mc-cnt"] = peMcStats.PktTruncEopMcCnt
    leafs["pkt-sop-drop-mc-cnt"] = peMcStats.PktSopDropMcCnt
    leafs["pkt-ecc-err-drop-mc-cnt"] = peMcStats.PktEccErrDropMcCnt
    leafs["pkt-ecc-err-trunc-cnt-mc-cnt"] = peMcStats.PktEccErrTruncCntMcCnt
    leafs["ecc-1bit-err-mc0-cnt"] = peMcStats.Ecc1BitErrMc0Cnt
    leafs["ecc-1bit-err-mc1-cnt"] = peMcStats.Ecc1BitErrMc1Cnt
    leafs["ecc-1bit-err-mc2-cnt"] = peMcStats.Ecc1BitErrMc2Cnt
    leafs["ecc-2bit-err-mc0-cnt"] = peMcStats.Ecc2BitErrMc0Cnt
    leafs["ecc-2bit-err-mc1-cnt"] = peMcStats.Ecc2BitErrMc1Cnt
    leafs["ecc-2bit-err-mc2-cnt"] = peMcStats.Ecc2BitErrMc2Cnt
    leafs["out-pkt-mc-cnt"] = peMcStats.OutPktMcCnt
    leafs["fe-mc-sop-eop-pack-cnt"] = peMcStats.FeMcSopEopPackCnt
    leafs["fc-mc-0-1-trans-cnt"] = peMcStats.FcMc01TransCnt
    leafs["rate-cnt"] = peMcStats.RateCnt
    leafs["calc-rate"] = peMcStats.CalcRate
    return leafs
}

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetBundleName() string { return "cisco_ios_xr" }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetYangName() string { return "pe-mc-stats" }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) SetParent(parent types.Entity) { peMcStats.parent = parent }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetParent() types.Entity { return peMcStats.parent }

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetParentYangName() string { return "sm15-stat" }

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats
// pe cc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats struct {
    parent types.Entity
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

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetFilter() yfilter.YFilter { return peCcStats.YFilter }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) SetFilter(yf yfilter.YFilter) { peCcStats.YFilter = yf }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetGoName(yname string) string {
    if yname == "in-pkt-cnt" { return "InPktCnt" }
    if yname == "out-path0-pkt-cnt" { return "OutPath0PktCnt" }
    if yname == "out-path1-pkt-cnt" { return "OutPath1PktCnt" }
    if yname == "xbar-ecc-drop-pkt-cnt" { return "XbarEccDropPktCnt" }
    if yname == "mem0-drop-pkt-cnt" { return "Mem0DropPktCnt" }
    if yname == "mem1-drop-pkt-cnt" { return "Mem1DropPktCnt" }
    if yname == "congn-pkt-cnt" { return "CongnPktCnt" }
    if yname == "xbar-ecc-single-err-cnt" { return "XbarEccSingleErrCnt" }
    if yname == "xbar-ecc-double-err-cnt" { return "XbarEccDoubleErrCnt" }
    if yname == "mem0-ecc-single-err-cnt" { return "Mem0EccSingleErrCnt" }
    if yname == "mem0-ecc-double-err-cnt" { return "Mem0EccDoubleErrCnt" }
    if yname == "mem1-ecc-single-err-cnt" { return "Mem1EccSingleErrCnt" }
    if yname == "mem1-ecc-double-err-cnt" { return "Mem1EccDoubleErrCnt" }
    if yname == "fc-cc-0-1-trans-cnt" { return "FcCc01TransCnt" }
    if yname == "rate-cnt" { return "RateCnt" }
    if yname == "calc-rate" { return "CalcRate" }
    return ""
}

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetSegmentPath() string {
    return "pe-cc-stats"
}

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkt-cnt"] = peCcStats.InPktCnt
    leafs["out-path0-pkt-cnt"] = peCcStats.OutPath0PktCnt
    leafs["out-path1-pkt-cnt"] = peCcStats.OutPath1PktCnt
    leafs["xbar-ecc-drop-pkt-cnt"] = peCcStats.XbarEccDropPktCnt
    leafs["mem0-drop-pkt-cnt"] = peCcStats.Mem0DropPktCnt
    leafs["mem1-drop-pkt-cnt"] = peCcStats.Mem1DropPktCnt
    leafs["congn-pkt-cnt"] = peCcStats.CongnPktCnt
    leafs["xbar-ecc-single-err-cnt"] = peCcStats.XbarEccSingleErrCnt
    leafs["xbar-ecc-double-err-cnt"] = peCcStats.XbarEccDoubleErrCnt
    leafs["mem0-ecc-single-err-cnt"] = peCcStats.Mem0EccSingleErrCnt
    leafs["mem0-ecc-double-err-cnt"] = peCcStats.Mem0EccDoubleErrCnt
    leafs["mem1-ecc-single-err-cnt"] = peCcStats.Mem1EccSingleErrCnt
    leafs["mem1-ecc-double-err-cnt"] = peCcStats.Mem1EccDoubleErrCnt
    leafs["fc-cc-0-1-trans-cnt"] = peCcStats.FcCc01TransCnt
    leafs["rate-cnt"] = peCcStats.RateCnt
    leafs["calc-rate"] = peCcStats.CalcRate
    return leafs
}

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetBundleName() string { return "cisco_ios_xr" }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetYangName() string { return "pe-cc-stats" }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) SetParent(parent types.Entity) { peCcStats.parent = parent }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetParent() types.Entity { return peCcStats.parent }

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetParentYangName() string { return "sm15-stat" }

