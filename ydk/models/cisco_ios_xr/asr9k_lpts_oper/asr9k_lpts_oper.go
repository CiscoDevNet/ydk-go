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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes with platform specific lpts operation data.
    NodeStatics PlatformLptspIfibStatic_NodeStatics
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetFilter() yfilter.YFilter { return platformLptspIfibStatic.YFilter }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) SetFilter(yf yfilter.YFilter) { platformLptspIfibStatic.YFilter = yf }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetGoName(yname string) string {
    if yname == "node-statics" { return "NodeStatics" }
    return ""
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib-static"
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-statics" {
        return &platformLptspIfibStatic.NodeStatics
    }
    return nil
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-statics"] = &platformLptspIfibStatic.NodeStatics
    return children
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetBundleName() string { return "cisco_ios_xr" }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetYangName() string { return "platform-lptsp-ifib-static" }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) SetParent(parent types.Entity) { platformLptspIfibStatic.parent = parent }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetParent() types.Entity { return platformLptspIfibStatic.parent }

func (platformLptspIfibStatic *PlatformLptspIfibStatic) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lpts-oper" }

// PlatformLptspIfibStatic_NodeStatics
// List of nodes with platform specific lpts
// operation data
type PlatformLptspIfibStatic_NodeStatics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node with platform specific lpts data. The type is slice of
    // PlatformLptspIfibStatic_NodeStatics_NodeStatic.
    NodeStatic []PlatformLptspIfibStatic_NodeStatics_NodeStatic
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetFilter() yfilter.YFilter { return nodeStatics.YFilter }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) SetFilter(yf yfilter.YFilter) { nodeStatics.YFilter = yf }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetGoName(yname string) string {
    if yname == "node-static" { return "NodeStatic" }
    return ""
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetSegmentPath() string {
    return "node-statics"
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-static" {
        for _, c := range nodeStatics.NodeStatic {
            if nodeStatics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfibStatic_NodeStatics_NodeStatic{}
        nodeStatics.NodeStatic = append(nodeStatics.NodeStatic, child)
        return &nodeStatics.NodeStatic[len(nodeStatics.NodeStatic)-1]
    }
    return nil
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeStatics.NodeStatic {
        children[nodeStatics.NodeStatic[i].GetSegmentPath()] = &nodeStatics.NodeStatic[i]
    }
    return children
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetBundleName() string { return "cisco_ios_xr" }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetYangName() string { return "node-statics" }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) SetParent(parent types.Entity) { nodeStatics.parent = parent }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetParent() types.Entity { return nodeStatics.parent }

func (nodeStatics *PlatformLptspIfibStatic_NodeStatics) GetParentYangName() string { return "platform-lptsp-ifib-static" }

// PlatformLptspIfibStatic_NodeStatics_NodeStatic
// Node with platform specific lpts data
type PlatformLptspIfibStatic_NodeStatics_NodeStatic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // pl_pifib police data.
    Police PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police

    // pl_pifib stats.
    Stats PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetFilter() yfilter.YFilter { return nodeStatic.YFilter }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) SetFilter(yf yfilter.YFilter) { nodeStatic.YFilter = yf }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "police" { return "Police" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetSegmentPath() string {
    return "node-static" + "[node-name='" + fmt.Sprintf("%v", nodeStatic.NodeName) + "']"
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police" {
        return &nodeStatic.Police
    }
    if childYangName == "stats" {
        return &nodeStatic.Stats
    }
    return nil
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["police"] = &nodeStatic.Police
    children["stats"] = &nodeStatic.Stats
    return children
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeStatic.NodeName
    return leafs
}

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetBundleName() string { return "cisco_ios_xr" }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetYangName() string { return "node-static" }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) SetParent(parent types.Entity) { nodeStatic.parent = parent }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetParent() types.Entity { return nodeStatic.parent }

func (nodeStatic *PlatformLptspIfibStatic_NodeStatics_NodeStatic) GetParentYangName() string { return "node-statics" }

// PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police
// pl_pifib police data
type PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per punt reason info. The type is slice of
    // PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo.
    StaticInfo []PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetGoName(yname string) string {
    if yname == "static-info" { return "StaticInfo" }
    return ""
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-info" {
        for _, c := range police.StaticInfo {
            if police.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo{}
        police.StaticInfo = append(police.StaticInfo, child)
        return &police.StaticInfo[len(police.StaticInfo)-1]
    }
    return nil
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range police.StaticInfo {
        children[police.StaticInfo[i].GetSegmentPath()] = &police.StaticInfo[i]
    }
    return children
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetYangName() string { return "police" }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police) GetParentYangName() string { return "node-static" }

// PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo
// Per punt reason info
type PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo struct {
    parent types.Entity
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

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetFilter() yfilter.YFilter { return staticInfo.YFilter }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) SetFilter(yf yfilter.YFilter) { staticInfo.YFilter = yf }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetGoName(yname string) string {
    if yname == "punt-reason" { return "PuntReason" }
    if yname == "sid" { return "Sid" }
    if yname == "flow-rate" { return "FlowRate" }
    if yname == "burst-rate" { return "BurstRate" }
    if yname == "accepted" { return "Accepted" }
    if yname == "dropped" { return "Dropped" }
    if yname == "punt-reason-string" { return "PuntReasonString" }
    if yname == "change-type" { return "ChangeType" }
    return ""
}

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetSegmentPath() string {
    return "static-info"
}

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["punt-reason"] = staticInfo.PuntReason
    leafs["sid"] = staticInfo.Sid
    leafs["flow-rate"] = staticInfo.FlowRate
    leafs["burst-rate"] = staticInfo.BurstRate
    leafs["accepted"] = staticInfo.Accepted
    leafs["dropped"] = staticInfo.Dropped
    leafs["punt-reason-string"] = staticInfo.PuntReasonString
    leafs["change-type"] = staticInfo.ChangeType
    return leafs
}

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetBundleName() string { return "cisco_ios_xr" }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetYangName() string { return "static-info" }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) SetParent(parent types.Entity) { staticInfo.parent = parent }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetParent() types.Entity { return staticInfo.parent }

func (staticInfo *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Police_StaticInfo) GetParentYangName() string { return "police" }

// PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats
// pl_pifib stats
type PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats struct {
    parent types.Entity
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

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetGoName(yname string) string {
    if yname == "accepted" { return "Accepted" }
    if yname == "dropped" { return "Dropped" }
    if yname == "clear-ts" { return "ClearTs" }
    if yname == "no-stats-mem-err" { return "NoStatsMemErr" }
    return ""
}

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accepted"] = stats.Accepted
    leafs["dropped"] = stats.Dropped
    leafs["clear-ts"] = stats.ClearTs
    leafs["no-stats-mem-err"] = stats.NoStatsMemErr
    return leafs
}

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetYangName() string { return "stats" }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetParent() types.Entity { return stats.parent }

func (stats *PlatformLptspIfibStatic_NodeStatics_NodeStatic_Stats) GetParentYangName() string { return "node-static" }

// PlatformLptspIfib
// platform lptsp ifib
type PlatformLptspIfib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes with platform specific lpts operation data.
    Nodes PlatformLptspIfib_Nodes
}

func (platformLptspIfib *PlatformLptspIfib) GetFilter() yfilter.YFilter { return platformLptspIfib.YFilter }

func (platformLptspIfib *PlatformLptspIfib) SetFilter(yf yfilter.YFilter) { platformLptspIfib.YFilter = yf }

func (platformLptspIfib *PlatformLptspIfib) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (platformLptspIfib *PlatformLptspIfib) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib"
}

func (platformLptspIfib *PlatformLptspIfib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &platformLptspIfib.Nodes
    }
    return nil
}

func (platformLptspIfib *PlatformLptspIfib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &platformLptspIfib.Nodes
    return children
}

func (platformLptspIfib *PlatformLptspIfib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformLptspIfib *PlatformLptspIfib) GetBundleName() string { return "cisco_ios_xr" }

func (platformLptspIfib *PlatformLptspIfib) GetYangName() string { return "platform-lptsp-ifib" }

func (platformLptspIfib *PlatformLptspIfib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformLptspIfib *PlatformLptspIfib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformLptspIfib *PlatformLptspIfib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformLptspIfib *PlatformLptspIfib) SetParent(parent types.Entity) { platformLptspIfib.parent = parent }

func (platformLptspIfib *PlatformLptspIfib) GetParent() types.Entity { return platformLptspIfib.parent }

func (platformLptspIfib *PlatformLptspIfib) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lpts-oper" }

// PlatformLptspIfib_Nodes
// List of nodes with platform specific lpts
// operation data
type PlatformLptspIfib_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node with platform specific lpts data. The type is slice of
    // PlatformLptspIfib_Nodes_Node.
    Node []PlatformLptspIfib_Nodes_Node
}

func (nodes *PlatformLptspIfib_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PlatformLptspIfib_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PlatformLptspIfib_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PlatformLptspIfib_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PlatformLptspIfib_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfib_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PlatformLptspIfib_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PlatformLptspIfib_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PlatformLptspIfib_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PlatformLptspIfib_Nodes) GetYangName() string { return "nodes" }

func (nodes *PlatformLptspIfib_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PlatformLptspIfib_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PlatformLptspIfib_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PlatformLptspIfib_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PlatformLptspIfib_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PlatformLptspIfib_Nodes) GetParentYangName() string { return "platform-lptsp-ifib" }

// PlatformLptspIfib_Nodes_Node
// Node with platform specific lpts data
type PlatformLptspIfib_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // pl_pifib police data.
    Police PlatformLptspIfib_Nodes_Node_Police

    // pl_pifib stats.
    Stats PlatformLptspIfib_Nodes_Node_Stats
}

func (node *PlatformLptspIfib_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PlatformLptspIfib_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PlatformLptspIfib_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "police" { return "Police" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (node *PlatformLptspIfib_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *PlatformLptspIfib_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police" {
        return &node.Police
    }
    if childYangName == "stats" {
        return &node.Stats
    }
    return nil
}

func (node *PlatformLptspIfib_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["police"] = &node.Police
    children["stats"] = &node.Stats
    return children
}

func (node *PlatformLptspIfib_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *PlatformLptspIfib_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PlatformLptspIfib_Nodes_Node) GetYangName() string { return "node" }

func (node *PlatformLptspIfib_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PlatformLptspIfib_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PlatformLptspIfib_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PlatformLptspIfib_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PlatformLptspIfib_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PlatformLptspIfib_Nodes_Node) GetParentYangName() string { return "nodes" }

// PlatformLptspIfib_Nodes_Node_Police
// pl_pifib police data
type PlatformLptspIfib_Nodes_Node_Police struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // PlatformLptspIfib_Nodes_Node_Police_PoliceInfo.
    PoliceInfo []PlatformLptspIfib_Nodes_Node_Police_PoliceInfo
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformLptspIfib_Nodes_Node_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetGoName(yname string) string {
    if yname == "police-info" { return "PoliceInfo" }
    return ""
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police-info" {
        for _, c := range police.PoliceInfo {
            if police.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfib_Nodes_Node_Police_PoliceInfo{}
        police.PoliceInfo = append(police.PoliceInfo, child)
        return &police.PoliceInfo[len(police.PoliceInfo)-1]
    }
    return nil
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range police.PoliceInfo {
        children[police.PoliceInfo[i].GetSegmentPath()] = &police.PoliceInfo[i]
    }
    return children
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (police *PlatformLptspIfib_Nodes_Node_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetYangName() string { return "police" }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformLptspIfib_Nodes_Node_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformLptspIfib_Nodes_Node_Police) GetParentYangName() string { return "node" }

// PlatformLptspIfib_Nodes_Node_Police_PoliceInfo
// Per flow type police info
type PlatformLptspIfib_Nodes_Node_Police_PoliceInfo struct {
    parent types.Entity
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

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetFilter() yfilter.YFilter { return policeInfo.YFilter }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) SetFilter(yf yfilter.YFilter) { policeInfo.YFilter = yf }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetGoName(yname string) string {
    if yname == "avgrate" { return "Avgrate" }
    if yname == "burst" { return "Burst" }
    if yname == "static-avgrate" { return "StaticAvgrate" }
    if yname == "avgrate-type" { return "AvgrateType" }
    if yname == "flow-type" { return "FlowType" }
    if yname == "accepted-stats" { return "AcceptedStats" }
    if yname == "dropped-stats" { return "DroppedStats" }
    if yname == "policer" { return "Policer" }
    if yname == "str-iptos-val" { return "StrIptosVal" }
    if yname == "change-type" { return "ChangeType" }
    if yname == "acl-config" { return "AclConfig" }
    if yname == "acl-str" { return "AclStr" }
    return ""
}

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetSegmentPath() string {
    return "police-info"
}

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["avgrate"] = policeInfo.Avgrate
    leafs["burst"] = policeInfo.Burst
    leafs["static-avgrate"] = policeInfo.StaticAvgrate
    leafs["avgrate-type"] = policeInfo.AvgrateType
    leafs["flow-type"] = policeInfo.FlowType
    leafs["accepted-stats"] = policeInfo.AcceptedStats
    leafs["dropped-stats"] = policeInfo.DroppedStats
    leafs["policer"] = policeInfo.Policer
    leafs["str-iptos-val"] = policeInfo.StrIptosVal
    leafs["change-type"] = policeInfo.ChangeType
    leafs["acl-config"] = policeInfo.AclConfig
    leafs["acl-str"] = policeInfo.AclStr
    return leafs
}

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetYangName() string { return "police-info" }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) SetParent(parent types.Entity) { policeInfo.parent = parent }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetParent() types.Entity { return policeInfo.parent }

func (policeInfo *PlatformLptspIfib_Nodes_Node_Police_PoliceInfo) GetParentYangName() string { return "police" }

// PlatformLptspIfib_Nodes_Node_Stats
// pl_pifib stats
type PlatformLptspIfib_Nodes_Node_Stats struct {
    parent types.Entity
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

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetGoName(yname string) string {
    if yname == "accepted" { return "Accepted" }
    if yname == "dropped" { return "Dropped" }
    if yname == "clear-ts" { return "ClearTs" }
    if yname == "no-stats-mem-err" { return "NoStatsMemErr" }
    return ""
}

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accepted"] = stats.Accepted
    leafs["dropped"] = stats.Dropped
    leafs["clear-ts"] = stats.ClearTs
    leafs["no-stats-mem-err"] = stats.NoStatsMemErr
    return leafs
}

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetYangName() string { return "stats" }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetParent() types.Entity { return stats.parent }

func (stats *PlatformLptspIfib_Nodes_Node_Stats) GetParentYangName() string { return "node" }

// PlatformLptspIfibNpStats
// platform lptsp ifib np stats
type PlatformLptspIfibNpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes with platform specific lpts operation data.
    NodeNpStats PlatformLptspIfibNpStats_NodeNpStats
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetFilter() yfilter.YFilter { return platformLptspIfibNpStats.YFilter }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) SetFilter(yf yfilter.YFilter) { platformLptspIfibNpStats.YFilter = yf }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetGoName(yname string) string {
    if yname == "node-np-stats" { return "NodeNpStats" }
    return ""
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lpts-oper:platform-lptsp-ifib-np-stats"
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-np-stats" {
        return &platformLptspIfibNpStats.NodeNpStats
    }
    return nil
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-np-stats"] = &platformLptspIfibNpStats.NodeNpStats
    return children
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetBundleName() string { return "cisco_ios_xr" }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetYangName() string { return "platform-lptsp-ifib-np-stats" }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) SetParent(parent types.Entity) { platformLptspIfibNpStats.parent = parent }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetParent() types.Entity { return platformLptspIfibNpStats.parent }

func (platformLptspIfibNpStats *PlatformLptspIfibNpStats) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lpts-oper" }

// PlatformLptspIfibNpStats_NodeNpStats
// List of nodes with platform specific lpts
// operation data
type PlatformLptspIfibNpStats_NodeNpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node with platform specific lpts data. The type is slice of
    // PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat.
    NodeNpStat []PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetFilter() yfilter.YFilter { return nodeNpStats.YFilter }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) SetFilter(yf yfilter.YFilter) { nodeNpStats.YFilter = yf }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetGoName(yname string) string {
    if yname == "node-np-stat" { return "NodeNpStat" }
    return ""
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetSegmentPath() string {
    return "node-np-stats"
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-np-stat" {
        for _, c := range nodeNpStats.NodeNpStat {
            if nodeNpStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat{}
        nodeNpStats.NodeNpStat = append(nodeNpStats.NodeNpStat, child)
        return &nodeNpStats.NodeNpStat[len(nodeNpStats.NodeNpStat)-1]
    }
    return nil
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeNpStats.NodeNpStat {
        children[nodeNpStats.NodeNpStat[i].GetSegmentPath()] = &nodeNpStats.NodeNpStat[i]
    }
    return children
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetBundleName() string { return "cisco_ios_xr" }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetYangName() string { return "node-np-stats" }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) SetParent(parent types.Entity) { nodeNpStats.parent = parent }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetParent() types.Entity { return nodeNpStats.parent }

func (nodeNpStats *PlatformLptspIfibNpStats_NodeNpStats) GetParentYangName() string { return "platform-lptsp-ifib-np-stats" }

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat
// Node with platform specific lpts data
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of all NP.
    Nps PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetFilter() yfilter.YFilter { return nodeNpStat.YFilter }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) SetFilter(yf yfilter.YFilter) { nodeNpStat.YFilter = yf }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "nps" { return "Nps" }
    return ""
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetSegmentPath() string {
    return "node-np-stat" + "[node-name='" + fmt.Sprintf("%v", nodeNpStat.NodeName) + "']"
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nps" {
        return &nodeNpStat.Nps
    }
    return nil
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nps"] = &nodeNpStat.Nps
    return children
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeNpStat.NodeName
    return leafs
}

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetBundleName() string { return "cisco_ios_xr" }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetYangName() string { return "node-np-stat" }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) SetParent(parent types.Entity) { nodeNpStat.parent = parent }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetParent() types.Entity { return nodeNpStat.parent }

func (nodeNpStat *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat) GetParentYangName() string { return "node-np-stats" }

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps
// List of all NP
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // np0 to np7. The type is slice of
    // PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np.
    Np []PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetFilter() yfilter.YFilter { return nps.YFilter }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) SetFilter(yf yfilter.YFilter) { nps.YFilter = yf }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetGoName(yname string) string {
    if yname == "np" { return "Np" }
    return ""
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetSegmentPath() string {
    return "nps"
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np" {
        for _, c := range nps.Np {
            if nps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np{}
        nps.Np = append(nps.Np, child)
        return &nps.Np[len(nps.Np)-1]
    }
    return nil
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nps.Np {
        children[nps.Np[i].GetSegmentPath()] = &nps.Np[i]
    }
    return children
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetBundleName() string { return "cisco_ios_xr" }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetYangName() string { return "nps" }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) SetParent(parent types.Entity) { nps.parent = parent }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetParent() types.Entity { return nps.parent }

func (nps *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps) GetParentYangName() string { return "node-np-stat" }

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np
// np0 to np7
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NP name. The type is string with pattern:
    // (np0)|(np1)|(np2)|(np3)|(np4)|(np5)|(np6)|(np7).
    NpName interface{}

    // pl_pifib police data.
    NpPolice PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetFilter() yfilter.YFilter { return np.YFilter }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) SetFilter(yf yfilter.YFilter) { np.YFilter = yf }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetGoName(yname string) string {
    if yname == "np-name" { return "NpName" }
    if yname == "np-police" { return "NpPolice" }
    return ""
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetSegmentPath() string {
    return "np" + "[np-name='" + fmt.Sprintf("%v", np.NpName) + "']"
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np-police" {
        return &np.NpPolice
    }
    return nil
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["np-police"] = &np.NpPolice
    return children
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["np-name"] = np.NpName
    return leafs
}

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetBundleName() string { return "cisco_ios_xr" }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetYangName() string { return "np" }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) SetParent(parent types.Entity) { np.parent = parent }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetParent() types.Entity { return np.parent }

func (np *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np) GetParentYangName() string { return "nps" }

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice
// pl_pifib police data
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo.
    PoliceInfo []PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetFilter() yfilter.YFilter { return npPolice.YFilter }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) SetFilter(yf yfilter.YFilter) { npPolice.YFilter = yf }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetGoName(yname string) string {
    if yname == "police-info" { return "PoliceInfo" }
    return ""
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetSegmentPath() string {
    return "np-police"
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police-info" {
        for _, c := range npPolice.PoliceInfo {
            if npPolice.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo{}
        npPolice.PoliceInfo = append(npPolice.PoliceInfo, child)
        return &npPolice.PoliceInfo[len(npPolice.PoliceInfo)-1]
    }
    return nil
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npPolice.PoliceInfo {
        children[npPolice.PoliceInfo[i].GetSegmentPath()] = &npPolice.PoliceInfo[i]
    }
    return children
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetBundleName() string { return "cisco_ios_xr" }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetYangName() string { return "np-police" }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) SetParent(parent types.Entity) { npPolice.parent = parent }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetParent() types.Entity { return npPolice.parent }

func (npPolice *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice) GetParentYangName() string { return "np" }

// PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo
// Per flow type police info
type PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo struct {
    parent types.Entity
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

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetFilter() yfilter.YFilter { return policeInfo.YFilter }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) SetFilter(yf yfilter.YFilter) { policeInfo.YFilter = yf }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetGoName(yname string) string {
    if yname == "avgrate" { return "Avgrate" }
    if yname == "burst" { return "Burst" }
    if yname == "static-avgrate" { return "StaticAvgrate" }
    if yname == "avgrate-type" { return "AvgrateType" }
    if yname == "flow-type" { return "FlowType" }
    if yname == "accepted-stats" { return "AcceptedStats" }
    if yname == "dropped-stats" { return "DroppedStats" }
    if yname == "policer" { return "Policer" }
    if yname == "str-iptos-val" { return "StrIptosVal" }
    if yname == "change-type" { return "ChangeType" }
    if yname == "acl-config" { return "AclConfig" }
    if yname == "acl-str" { return "AclStr" }
    return ""
}

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetSegmentPath() string {
    return "police-info"
}

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["avgrate"] = policeInfo.Avgrate
    leafs["burst"] = policeInfo.Burst
    leafs["static-avgrate"] = policeInfo.StaticAvgrate
    leafs["avgrate-type"] = policeInfo.AvgrateType
    leafs["flow-type"] = policeInfo.FlowType
    leafs["accepted-stats"] = policeInfo.AcceptedStats
    leafs["dropped-stats"] = policeInfo.DroppedStats
    leafs["policer"] = policeInfo.Policer
    leafs["str-iptos-val"] = policeInfo.StrIptosVal
    leafs["change-type"] = policeInfo.ChangeType
    leafs["acl-config"] = policeInfo.AclConfig
    leafs["acl-str"] = policeInfo.AclStr
    return leafs
}

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetYangName() string { return "police-info" }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) SetParent(parent types.Entity) { policeInfo.parent = parent }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetParent() types.Entity { return policeInfo.parent }

func (policeInfo *PlatformLptspIfibNpStats_NodeNpStats_NodeNpStat_Nps_Np_NpPolice_PoliceInfo) GetParentYangName() string { return "np-police" }

