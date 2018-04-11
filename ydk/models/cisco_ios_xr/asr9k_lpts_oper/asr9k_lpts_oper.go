// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lpts package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform-lptsp-ifib-static: ASR9K platform ifib operational
//     data 
//   platform-lptsp-ifib: platform lptsp ifib
//   platform-lptsp-ifib-np-stats: platform lptsp ifib np stats
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_lpts_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_lpts_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lpts-oper platform-lptsp-ifib-static}", reflect.TypeOf(PlatformLptspIfibStatic{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib-static", reflect.TypeOf(PlatformLptspIfibStatic{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lpts-oper platform-lptsp-ifib}", reflect.TypeOf(PlatformLptspIfib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib", reflect.TypeOf(PlatformLptspIfib{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lpts-oper platform-lptsp-ifib-np-stats}", reflect.TypeOf(PlatformLptspIfibNpStats{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib-np-stats", reflect.TypeOf(PlatformLptspIfibNpStats{}))
}

// PlatformLptspIfibStatic
// ASR9K platform ifib operational data 
type PlatformLptspIfibStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes with platform specific lpts operation data.
    NodeStatics PlatformLptspIfibStatic_NodeStatics
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetEntityData() *types.CommonEntityData {
    platformLptspIfibStatic.EntityData.YFilter = platformLptspIfibStatic.YFilter
    platformLptspIfibStatic.EntityData.YangName = "platform-lptsp-ifib-static"
    platformLptspIfibStatic.EntityData.BundleName = "cisco_ios_xr"
    platformLptspIfibStatic.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lpts-oper"
    platformLptspIfibStatic.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib-static"
    platformLptspIfibStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformLptspIfibStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformLptspIfibStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformLptspIfibStatic.EntityData.Children = make(map[string]types.YChild)
    platformLptspIfibStatic.EntityData.Children["node-statics"] = types.YChild{"NodeStatics", &platformLptspIfibStatic.NodeStatics}
    platformLptspIfibStatic.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformLptspIfibStatic.EntityData)
}

// PlatformLptspIfibStatic_NodeStatics
// List of nodes with platform specific lpts
// operation data
type PlatformLptspIfibStatic_NodeStatics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node with platform specific lpts data. The type is slice of
    // PlatformLptspIfibStatic_NodeStatics_NodeStatic.
    NodeStatic []PlatformLptspIfibStatic_NodeStatics_NodeStatic
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetEntityData() *types.CommonEntityData {
    nodeStatics.EntityData.YFilter = nodeStatics.YFilter
    nodeStatics.EntityData.YangName = "node-statics"
    nodeStatics.EntityData.BundleName = "cisco_ios_xr"
    nodeStatics.EntityData.ParentYangName = "platform-lptsp-ifib-static"
    nodeStatics.EntityData.SegmentPath = "node-statics"
    nodeStatics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeStatics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeStatics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeStatics.EntityData.Children = make(map[string]types.YChild)
    nodeStatics.EntityData.Children["node-static"] = types.YChild{"NodeStatic", nil}
    for i := range nodeStatics.NodeStatic {
        nodeStatics.EntityData.Children[types.GetSegmentPath(&nodeStatics.NodeStatic[i])] = types.YChild{"NodeStatic", &nodeStatics.NodeStatic[i]}
    }
    nodeStatics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodeStatics.EntityData)
}

// PlatformLptspIfibStatic_NodeStatics_NodeStatic
// Node with platform specific lpts data
type PlatformLptspIfibStatic_NodeStatics_NodeStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // pl_pifib police data.
    Police PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police

    // pl_pifib stats.
    Stats PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetEntityData() *types.CommonEntityData {
    nodeStatic.EntityData.YFilter = nodeStatic.YFilter
    nodeStatic.EntityData.YangName = "node-static"
    nodeStatic.EntityData.BundleName = "cisco_ios_xr"
    nodeStatic.EntityData.ParentYangName = "node-statics"
    nodeStatic.EntityData.SegmentPath = "node-static" + "[node-name='" + fmt.Sprintf("%v", nodeStatic.NodeName) + "']"
    nodeStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeStatic.EntityData.Children = make(map[string]types.YChild)
    nodeStatic.EntityData.Children["police"] = types.YChild{"Police", &nodeStatic.Police}
    nodeStatic.EntityData.Children["stats"] = types.YChild{"Stats", &nodeStatic.Stats}
    nodeStatic.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeStatic.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeStatic.NodeName}
    return &(nodeStatic.EntityData)
}

// PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police
// pl_pifib police data
type PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per punt reason info. The type is slice of
    // PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo.
    StaticInfo []PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "node-static"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["static-info"] = types.YChild{"StaticInfo", nil}
    for i := range police.StaticInfo {
        police.EntityData.Children[types.GetSegmentPath(&police.StaticInfo[i])] = types.YChild{"StaticInfo", &police.StaticInfo[i]}
    }
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(police.EntityData)
}

// PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo
// Per punt reason info
type PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // punt reason. The type is interface{} with range: 0..4294967295.
    PuntReason interface{}

    // sid. The type is interface{} with range: 0..4294967295.
    Sid interface{}

    // flow rate. The type is interface{} with range: 0..4294967295.
    FlowRate interface{}

    // burst rate. The type is interface{} with range: 0..4294967295.
    BurstRate interface{}

    // accepted. The type is interface{} with range: 0..18446744073709551615.
    Accepted interface{}

    // dropped. The type is interface{} with range: 0..18446744073709551615.
    Dropped interface{}

    // punt reason string. The type is string with length: 0..50.
    PuntReasonString interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}
}

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetEntityData() *types.CommonEntityData {
    staticInfo.EntityData.YFilter = staticInfo.YFilter
    staticInfo.EntityData.YangName = "static-info"
    staticInfo.EntityData.BundleName = "cisco_ios_xr"
    staticInfo.EntityData.ParentYangName = "police"
    staticInfo.EntityData.SegmentPath = "static-info"
    staticInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticInfo.EntityData.Children = make(map[string]types.YChild)
    staticInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    staticInfo.EntityData.Leafs["punt-reason"] = types.YLeaf{"PuntReason", staticInfo.PuntReason}
    staticInfo.EntityData.Leafs["sid"] = types.YLeaf{"Sid", staticInfo.Sid}
    staticInfo.EntityData.Leafs["flow-rate"] = types.YLeaf{"FlowRate", staticInfo.FlowRate}
    staticInfo.EntityData.Leafs["burst-rate"] = types.YLeaf{"BurstRate", staticInfo.BurstRate}
    staticInfo.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", staticInfo.Accepted}
    staticInfo.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", staticInfo.Dropped}
    staticInfo.EntityData.Leafs["punt-reason-string"] = types.YLeaf{"PuntReasonString", staticInfo.PuntReasonString}
    staticInfo.EntityData.Leafs["change-type"] = types.YLeaf{"ChangeType", staticInfo.ChangeType}
    return &(staticInfo.EntityData)
}

// PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats
// pl_pifib stats
type PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Deleted-entry accepted packets counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepted interface{}

    // Deleted-entry dropped packets counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Dropped interface{}

    // Statistics clear timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTs interface{}

    // No statistics memory error. The type is interface{} with range:
    // 0..18446744073709551615.
    NoStatsMemErr interface{}
}

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "node-static"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", stats.Accepted}
    stats.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", stats.Dropped}
    stats.EntityData.Leafs["clear-ts"] = types.YLeaf{"ClearTs", stats.ClearTs}
    stats.EntityData.Leafs["no-stats-mem-err"] = types.YLeaf{"NoStatsMemErr", stats.NoStatsMemErr}
    return &(stats.EntityData)
}

// PlatformLptspIfib
// platform lptsp ifib
type PlatformLptspIfib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes with platform specific lpts operation data.
    Nodes PlatformLptspIfib_Nodes
}

func (platformLptspIfib *PlatformLptspIfib) GetEntityData() *types.CommonEntityData {
    platformLptspIfib.EntityData.YFilter = platformLptspIfib.YFilter
    platformLptspIfib.EntityData.YangName = "platform-lptsp-ifib"
    platformLptspIfib.EntityData.BundleName = "cisco_ios_xr"
    platformLptspIfib.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lpts-oper"
    platformLptspIfib.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib"
    platformLptspIfib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformLptspIfib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformLptspIfib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformLptspIfib.EntityData.Children = make(map[string]types.YChild)
    platformLptspIfib.EntityData.Children["nodes"] = types.YChild{"Nodes", &platformLptspIfib.Nodes}
    platformLptspIfib.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformLptspIfib.EntityData)
}

// PlatformLptspIfib_Nodes
// List of nodes with platform specific lpts
// operation data
type PlatformLptspIfib_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node with platform specific lpts data. The type is slice of
    // PlatformLptspIfib_Nodes_Node.
    Node []PlatformLptspIfib_Nodes_Node
}

func (nodes *PlatformLptspIfib_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "platform-lptsp-ifib"
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

// PlatformLptspIfib_Nodes_Node
// Node with platform specific lpts data
type PlatformLptspIfib_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // pl_pifib police data.
    Police PlatformLptspIfib_Nodes_Node_Police

    // pl_pifib stats.
    Stats PlatformLptspIfib_Nodes_Node_Stats
}

func (node *PlatformLptspIfib_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["police"] = types.YChild{"Police", &node.Police}
    node.EntityData.Children["stats"] = types.YChild{"Stats", &node.Stats}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// PlatformLptspIfib_Nodes_Node_Police
// pl_pifib police data
type PlatformLptspIfib_Nodes_Node_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // PlatformLptspIfib_Nodes_Node_Police_PoliceInfo.
    PoliceInfo []PlatformLptspIfib_Nodes_Node_Police_PoliceInfo
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "node"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["police-info"] = types.YChild{"PoliceInfo", nil}
    for i := range police.PoliceInfo {
        police.EntityData.Children[types.GetSegmentPath(&police.PoliceInfo[i])] = types.YChild{"PoliceInfo", &police.PoliceInfo[i]}
    }
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(police.EntityData)
}

// PlatformLptspIfib_Nodes_Node_Police_PoliceInfo
// Per flow type police info
type PlatformLptspIfib_Nodes_Node_Police_PoliceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // avgrate. The type is interface{} with range: 0..4294967295.
    Avgrate interface{}

    // burst. The type is interface{} with range: 0..4294967295.
    Burst interface{}

    // static avgrate. The type is interface{} with range: 0..4294967295.
    StaticAvgrate interface{}

    // avgrate type. The type is string with length: 0..50.
    AvgrateType interface{}

    // flow type. The type is string with length: 0..50.
    FlowType interface{}

    // accepted stats. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedStats interface{}

    // dropped stats. The type is interface{} with range: 0..18446744073709551615.
    DroppedStats interface{}

    // policer. The type is interface{} with range: 0..4294967295.
    Policer interface{}

    // str iptos val. The type is string with length: 0..8.
    StrIptosVal interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}

    // acl config. The type is interface{} with range: 0..255.
    AclConfig interface{}

    // acl str. The type is string with length: 0..50.
    AclStr interface{}
}

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetEntityData() *types.CommonEntityData {
    policeInfo.EntityData.YFilter = policeInfo.YFilter
    policeInfo.EntityData.YangName = "police-info"
    policeInfo.EntityData.BundleName = "cisco_ios_xr"
    policeInfo.EntityData.ParentYangName = "police"
    policeInfo.EntityData.SegmentPath = "police-info"
    policeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeInfo.EntityData.Children = make(map[string]types.YChild)
    policeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    policeInfo.EntityData.Leafs["avgrate"] = types.YLeaf{"Avgrate", policeInfo.Avgrate}
    policeInfo.EntityData.Leafs["burst"] = types.YLeaf{"Burst", policeInfo.Burst}
    policeInfo.EntityData.Leafs["static-avgrate"] = types.YLeaf{"StaticAvgrate", policeInfo.StaticAvgrate}
    policeInfo.EntityData.Leafs["avgrate-type"] = types.YLeaf{"AvgrateType", policeInfo.AvgrateType}
    policeInfo.EntityData.Leafs["flow-type"] = types.YLeaf{"FlowType", policeInfo.FlowType}
    policeInfo.EntityData.Leafs["accepted-stats"] = types.YLeaf{"AcceptedStats", policeInfo.AcceptedStats}
    policeInfo.EntityData.Leafs["dropped-stats"] = types.YLeaf{"DroppedStats", policeInfo.DroppedStats}
    policeInfo.EntityData.Leafs["policer"] = types.YLeaf{"Policer", policeInfo.Policer}
    policeInfo.EntityData.Leafs["str-iptos-val"] = types.YLeaf{"StrIptosVal", policeInfo.StrIptosVal}
    policeInfo.EntityData.Leafs["change-type"] = types.YLeaf{"ChangeType", policeInfo.ChangeType}
    policeInfo.EntityData.Leafs["acl-config"] = types.YLeaf{"AclConfig", policeInfo.AclConfig}
    policeInfo.EntityData.Leafs["acl-str"] = types.YLeaf{"AclStr", policeInfo.AclStr}
    return &(policeInfo.EntityData)
}

// PlatformLptspIfib_Nodes_Node_Stats
// pl_pifib stats
type PlatformLptspIfib_Nodes_Node_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Deleted-entry accepted packets counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepted interface{}

    // Deleted-entry dropped packets counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Dropped interface{}

    // Statistics clear timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTs interface{}

    // No statistics memory error. The type is interface{} with range:
    // 0..18446744073709551615.
    NoStatsMemErr interface{}
}

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "node"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", stats.Accepted}
    stats.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", stats.Dropped}
    stats.EntityData.Leafs["clear-ts"] = types.YLeaf{"ClearTs", stats.ClearTs}
    stats.EntityData.Leafs["no-stats-mem-err"] = types.YLeaf{"NoStatsMemErr", stats.NoStatsMemErr}
    return &(stats.EntityData)
}

// PlatformLptspIfibNpStats
// platform lptsp ifib np stats
type PlatformLptspIfibNpStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes with platform specific lpts operation data.
    NodeNpStats PlatformLptspIfibNpStats_NodeNpStats
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetEntityData() *types.CommonEntityData {
    platformLptspIfibNpStats.EntityData.YFilter = platformLptspIfibNpStats.YFilter
    platformLptspIfibNpStats.EntityData.YangName = "platform-lptsp-ifib-np-stats"
    platformLptspIfibNpStats.EntityData.BundleName = "cisco_ios_xr"
    platformLptspIfibNpStats.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lpts-oper"
    platformLptspIfibNpStats.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib-np-stats"
    platformLptspIfibNpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformLptspIfibNpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformLptspIfibNpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformLptspIfibNpStats.EntityData.Children = make(map[string]types.YChild)
    platformLptspIfibNpStats.EntityData.Children["node-np-stats"] = types.YChild{"NodeNpStats", &platformLptspIfibNpStats.NodeNpStats}
    platformLptspIfibNpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformLptspIfibNpStats.EntityData)
}

// PlatformLptspIfibNpStats_NodeNpStats
// List of nodes with platform specific lpts
// operation data
type PlatformLptspIfibNpStats_NodeNpStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node with platform specific lpts data. The type is slice of
    // PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat.
    NodeNpStat []PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetEntityData() *types.CommonEntityData {
    nodeNpStats.EntityData.YFilter = nodeNpStats.YFilter
    nodeNpStats.EntityData.YangName = "node-np-stats"
    nodeNpStats.EntityData.BundleName = "cisco_ios_xr"
    nodeNpStats.EntityData.ParentYangName = "platform-lptsp-ifib-np-stats"
    nodeNpStats.EntityData.SegmentPath = "node-np-stats"
    nodeNpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeNpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeNpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeNpStats.EntityData.Children = make(map[string]types.YChild)
    nodeNpStats.EntityData.Children["node-np-stat"] = types.YChild{"NodeNpStat", nil}
    for i := range nodeNpStats.NodeNpStat {
        nodeNpStats.EntityData.Children[types.GetSegmentPath(&nodeNpStats.NodeNpStat[i])] = types.YChild{"NodeNpStat", &nodeNpStats.NodeNpStat[i]}
    }
    nodeNpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodeNpStats.EntityData)
}

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat
// Node with platform specific lpts data
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // List of all NP.
    Nps PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetEntityData() *types.CommonEntityData {
    nodeNpStat.EntityData.YFilter = nodeNpStat.YFilter
    nodeNpStat.EntityData.YangName = "node-np-stat"
    nodeNpStat.EntityData.BundleName = "cisco_ios_xr"
    nodeNpStat.EntityData.ParentYangName = "node-np-stats"
    nodeNpStat.EntityData.SegmentPath = "node-np-stat" + "[node-name='" + fmt.Sprintf("%v", nodeNpStat.NodeName) + "']"
    nodeNpStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeNpStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeNpStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeNpStat.EntityData.Children = make(map[string]types.YChild)
    nodeNpStat.EntityData.Children["nps"] = types.YChild{"Nps", &nodeNpStat.Nps}
    nodeNpStat.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeNpStat.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeNpStat.NodeName}
    return &(nodeNpStat.EntityData)
}

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps
// List of all NP
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // np0 to np7. The type is slice of
    // PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np.
    Np []PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetEntityData() *types.CommonEntityData {
    nps.EntityData.YFilter = nps.YFilter
    nps.EntityData.YangName = "nps"
    nps.EntityData.BundleName = "cisco_ios_xr"
    nps.EntityData.ParentYangName = "node-np-stat"
    nps.EntityData.SegmentPath = "nps"
    nps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nps.EntityData.Children = make(map[string]types.YChild)
    nps.EntityData.Children["np"] = types.YChild{"Np", nil}
    for i := range nps.Np {
        nps.EntityData.Children[types.GetSegmentPath(&nps.Np[i])] = types.YChild{"Np", &nps.Np[i]}
    }
    nps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nps.EntityData)
}

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np
// np0 to np7
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NP name. The type is string with pattern:
    // b'(np0)|(np1)|(np2)|(np3)|(np4)|(np5)|(np6)|(np7)'.
    NpName interface{}

    // pl_pifib police data.
    NpPolice PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetEntityData() *types.CommonEntityData {
    np.EntityData.YFilter = np.YFilter
    np.EntityData.YangName = "np"
    np.EntityData.BundleName = "cisco_ios_xr"
    np.EntityData.ParentYangName = "nps"
    np.EntityData.SegmentPath = "np" + "[np-name='" + fmt.Sprintf("%v", np.NpName) + "']"
    np.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    np.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    np.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    np.EntityData.Children = make(map[string]types.YChild)
    np.EntityData.Children["np-police"] = types.YChild{"NpPolice", &np.NpPolice}
    np.EntityData.Leafs = make(map[string]types.YLeaf)
    np.EntityData.Leafs["np-name"] = types.YLeaf{"NpName", np.NpName}
    return &(np.EntityData)
}

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice
// pl_pifib police data
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo.
    PoliceInfo []PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetEntityData() *types.CommonEntityData {
    npPolice.EntityData.YFilter = npPolice.YFilter
    npPolice.EntityData.YangName = "np-police"
    npPolice.EntityData.BundleName = "cisco_ios_xr"
    npPolice.EntityData.ParentYangName = "np"
    npPolice.EntityData.SegmentPath = "np-police"
    npPolice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npPolice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npPolice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npPolice.EntityData.Children = make(map[string]types.YChild)
    npPolice.EntityData.Children["police-info"] = types.YChild{"PoliceInfo", nil}
    for i := range npPolice.PoliceInfo {
        npPolice.EntityData.Children[types.GetSegmentPath(&npPolice.PoliceInfo[i])] = types.YChild{"PoliceInfo", &npPolice.PoliceInfo[i]}
    }
    npPolice.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(npPolice.EntityData)
}

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo
// Per flow type police info
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // avgrate. The type is interface{} with range: 0..4294967295.
    Avgrate interface{}

    // burst. The type is interface{} with range: 0..4294967295.
    Burst interface{}

    // static avgrate. The type is interface{} with range: 0..4294967295.
    StaticAvgrate interface{}

    // avgrate type. The type is string with length: 0..50.
    AvgrateType interface{}

    // flow type. The type is string with length: 0..50.
    FlowType interface{}

    // accepted stats. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedStats interface{}

    // dropped stats. The type is interface{} with range: 0..18446744073709551615.
    DroppedStats interface{}

    // policer. The type is interface{} with range: 0..4294967295.
    Policer interface{}

    // str iptos val. The type is string with length: 0..8.
    StrIptosVal interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}

    // acl config. The type is interface{} with range: 0..255.
    AclConfig interface{}

    // acl str. The type is string with length: 0..50.
    AclStr interface{}
}

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetEntityData() *types.CommonEntityData {
    policeInfo.EntityData.YFilter = policeInfo.YFilter
    policeInfo.EntityData.YangName = "police-info"
    policeInfo.EntityData.BundleName = "cisco_ios_xr"
    policeInfo.EntityData.ParentYangName = "np-police"
    policeInfo.EntityData.SegmentPath = "police-info"
    policeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeInfo.EntityData.Children = make(map[string]types.YChild)
    policeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    policeInfo.EntityData.Leafs["avgrate"] = types.YLeaf{"Avgrate", policeInfo.Avgrate}
    policeInfo.EntityData.Leafs["burst"] = types.YLeaf{"Burst", policeInfo.Burst}
    policeInfo.EntityData.Leafs["static-avgrate"] = types.YLeaf{"StaticAvgrate", policeInfo.StaticAvgrate}
    policeInfo.EntityData.Leafs["avgrate-type"] = types.YLeaf{"AvgrateType", policeInfo.AvgrateType}
    policeInfo.EntityData.Leafs["flow-type"] = types.YLeaf{"FlowType", policeInfo.FlowType}
    policeInfo.EntityData.Leafs["accepted-stats"] = types.YLeaf{"AcceptedStats", policeInfo.AcceptedStats}
    policeInfo.EntityData.Leafs["dropped-stats"] = types.YLeaf{"DroppedStats", policeInfo.DroppedStats}
    policeInfo.EntityData.Leafs["policer"] = types.YLeaf{"Policer", policeInfo.Policer}
    policeInfo.EntityData.Leafs["str-iptos-val"] = types.YLeaf{"StrIptosVal", policeInfo.StrIptosVal}
    policeInfo.EntityData.Leafs["change-type"] = types.YLeaf{"ChangeType", policeInfo.ChangeType}
    policeInfo.EntityData.Leafs["acl-config"] = types.YLeaf{"AclConfig", policeInfo.AclConfig}
    policeInfo.EntityData.Leafs["acl-str"] = types.YLeaf{"AclStr", policeInfo.AclStr}
    return &(policeInfo.EntityData)
}

