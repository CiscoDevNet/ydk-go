// This module contains a collection of YANG definitions
// for Cisco IOS-XR sse-span package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ssespan: SSE SPAN operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package sse_span_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sse_span_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sse-span-oper ssespan}", reflect.TypeOf(Ssespan{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sse-span-oper:ssespan", reflect.TypeOf(Ssespan{}))
}

// Ssespan
// SSE SPAN operational data
type Ssespan struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node table for node-specific operational data.
    Nodes Ssespan_Nodes
}

func (ssespan *Ssespan) GetFilter() yfilter.YFilter { return ssespan.YFilter }

func (ssespan *Ssespan) SetFilter(yf yfilter.YFilter) { ssespan.YFilter = yf }

func (ssespan *Ssespan) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ssespan *Ssespan) GetSegmentPath() string {
    return "Cisco-IOS-XR-sse-span-oper:ssespan"
}

func (ssespan *Ssespan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ssespan.Nodes
    }
    return nil
}

func (ssespan *Ssespan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ssespan.Nodes
    return children
}

func (ssespan *Ssespan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssespan *Ssespan) GetBundleName() string { return "cisco_ios_xr" }

func (ssespan *Ssespan) GetYangName() string { return "ssespan" }

func (ssespan *Ssespan) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssespan *Ssespan) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssespan *Ssespan) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssespan *Ssespan) SetParent(parent types.Entity) { ssespan.parent = parent }

func (ssespan *Ssespan) GetParent() types.Entity { return ssespan.parent }

func (ssespan *Ssespan) GetParentYangName() string { return "Cisco-IOS-XR-sse-span-oper" }

// Ssespan_Nodes
// Node table for node-specific operational data
type Ssespan_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // Ssespan_Nodes_Node.
    Node []Ssespan_Nodes_Node
}

func (nodes *Ssespan_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ssespan_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ssespan_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ssespan_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ssespan_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssespan_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ssespan_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ssespan_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ssespan_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ssespan_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ssespan_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ssespan_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ssespan_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ssespan_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ssespan_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ssespan_Nodes) GetParentYangName() string { return "ssespan" }

// Ssespan_Nodes_Node
// Node-specific data for a particular node
type Ssespan_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // SPAN SFT entry.
    SpanMirrInfos Ssespan_Nodes_Node_SpanMirrInfos

    // udf info.
    Spanudf Ssespan_Nodes_Node_Spanudf

    // SPAN SDT entry.
    SpanSessInfos Ssespan_Nodes_Node_SpanSessInfos
}

func (node *Ssespan_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ssespan_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ssespan_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "span-mirr-infos" { return "SpanMirrInfos" }
    if yname == "spanudf" { return "Spanudf" }
    if yname == "span-sess-infos" { return "SpanSessInfos" }
    return ""
}

func (node *Ssespan_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *Ssespan_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "span-mirr-infos" {
        return &node.SpanMirrInfos
    }
    if childYangName == "spanudf" {
        return &node.Spanudf
    }
    if childYangName == "span-sess-infos" {
        return &node.SpanSessInfos
    }
    return nil
}

func (node *Ssespan_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["span-mirr-infos"] = &node.SpanMirrInfos
    children["spanudf"] = &node.Spanudf
    children["span-sess-infos"] = &node.SpanSessInfos
    return children
}

func (node *Ssespan_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *Ssespan_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ssespan_Nodes_Node) GetYangName() string { return "node" }

func (node *Ssespan_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ssespan_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ssespan_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ssespan_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ssespan_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ssespan_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ssespan_Nodes_Node_SpanMirrInfos
// SPAN SFT entry
type Ssespan_Nodes_Node_SpanMirrInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mirror info . The type is slice of
    // Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo.
    SpanMirrInfo []Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetFilter() yfilter.YFilter { return spanMirrInfos.YFilter }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) SetFilter(yf yfilter.YFilter) { spanMirrInfos.YFilter = yf }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetGoName(yname string) string {
    if yname == "span-mirr-info" { return "SpanMirrInfo" }
    return ""
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetSegmentPath() string {
    return "span-mirr-infos"
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "span-mirr-info" {
        for _, c := range spanMirrInfos.SpanMirrInfo {
            if spanMirrInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo{}
        spanMirrInfos.SpanMirrInfo = append(spanMirrInfos.SpanMirrInfo, child)
        return &spanMirrInfos.SpanMirrInfo[len(spanMirrInfos.SpanMirrInfo)-1]
    }
    return nil
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spanMirrInfos.SpanMirrInfo {
        children[spanMirrInfos.SpanMirrInfo[i].GetSegmentPath()] = &spanMirrInfos.SpanMirrInfo[i]
    }
    return children
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetBundleName() string { return "cisco_ios_xr" }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetYangName() string { return "span-mirr-infos" }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) SetParent(parent types.Entity) { spanMirrInfos.parent = parent }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetParent() types.Entity { return spanMirrInfos.parent }

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetParentYangName() string { return "node" }

// Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo
// Mirror info 
type Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    IntfName interface{}

    // source IFH. The type is interface{} with range: 0..4294967295.
    SrcIfh interface{}

    // interface name. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    IntfNameXr interface{}

    // ipv4 acl flag. The type is interface{} with range: 0..4294967295.
    V4AclFlag interface{}

    // ipv6 acl flag. The type is interface{} with range: 0..4294967295.
    V6AclFlag interface{}

    // gre acl flag. The type is interface{} with range: 0..4294967295.
    GreAclFlag interface{}

    // ipv4 state. The type is interface{} with range: 0..4294967295.
    V4State interface{}

    // ipv6 state. The type is interface{} with range: 0..4294967295.
    V6State interface{}

    // gre state. The type is interface{} with range: 0..4294967295.
    GreState interface{}

    // ipv4 session Id. The type is interface{} with range: 0..4294967295.
    V4Sessid interface{}

    // ipv6 session Id. The type is interface{} with range: 0..4294967295.
    V6Sessid interface{}

    // gre session Id. The type is interface{} with range: 0..4294967295.
    GreSessid interface{}

    // ipv4 dst type. The type is interface{} with range: 0..4294967295.
    V4DstType interface{}

    // ipv6 dst type. The type is interface{} with range: 0..4294967295.
    V6DstType interface{}

    // gre dst type. The type is interface{} with range: 0..4294967295.
    GredstType interface{}

    // v4 stats ptr. The type is interface{} with range: 0..4294967295.
    V4Statsptr interface{}

    // v6 stats ptr. The type is interface{} with range: 0..4294967295.
    V6Statsptr interface{}

    // mpls ipv4 stats ptr. The type is interface{} with range: 0..4294967295.
    Grev4Statsptr interface{}

    // mpls ipv6 stats ptr. The type is interface{} with range: 0..4294967295.
    Grev6Statsptr interface{}

    // mpls ipv4 pkts. The type is interface{} with range: 0..4294967295.
    Mplsv4Stats interface{}

    // mpls ipv6 pkts. The type is interface{} with range: 0..4294967295.
    Mplsv6Pkts interface{}

    // npu mask. The type is interface{} with range: 0..4294967295.
    NpUmask interface{}

    // uidb array . The type is slice of interface{} with range: 0..4294967295.
    Uidb []interface{}

    // 16x5npu=80 bytes of hw sft entry. The type is slice of interface{} with
    // range: 0..4294967295.
    SftHwData []interface{}
}

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetFilter() yfilter.YFilter { return spanMirrInfo.YFilter }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) SetFilter(yf yfilter.YFilter) { spanMirrInfo.YFilter = yf }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetGoName(yname string) string {
    if yname == "intf-name" { return "IntfName" }
    if yname == "src-ifh" { return "SrcIfh" }
    if yname == "intf-name-xr" { return "IntfNameXr" }
    if yname == "v4-acl-flag" { return "V4AclFlag" }
    if yname == "v6-acl-flag" { return "V6AclFlag" }
    if yname == "gre-acl-flag" { return "GreAclFlag" }
    if yname == "v4state" { return "V4State" }
    if yname == "v6state" { return "V6State" }
    if yname == "gre-state" { return "GreState" }
    if yname == "v4sessid" { return "V4Sessid" }
    if yname == "v6sessid" { return "V6Sessid" }
    if yname == "gre-sessid" { return "GreSessid" }
    if yname == "v4dst-type" { return "V4DstType" }
    if yname == "v6dst-type" { return "V6DstType" }
    if yname == "gredst-type" { return "GredstType" }
    if yname == "v4statsptr" { return "V4Statsptr" }
    if yname == "v6statsptr" { return "V6Statsptr" }
    if yname == "grev4statsptr" { return "Grev4Statsptr" }
    if yname == "grev6statsptr" { return "Grev6Statsptr" }
    if yname == "mplsv4stats" { return "Mplsv4Stats" }
    if yname == "mplsv6pkts" { return "Mplsv6Pkts" }
    if yname == "np-umask" { return "NpUmask" }
    if yname == "uidb" { return "Uidb" }
    if yname == "sft-hw-data" { return "SftHwData" }
    return ""
}

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetSegmentPath() string {
    return "span-mirr-info" + "[intf-name='" + fmt.Sprintf("%v", spanMirrInfo.IntfName) + "']"
}

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intf-name"] = spanMirrInfo.IntfName
    leafs["src-ifh"] = spanMirrInfo.SrcIfh
    leafs["intf-name-xr"] = spanMirrInfo.IntfNameXr
    leafs["v4-acl-flag"] = spanMirrInfo.V4AclFlag
    leafs["v6-acl-flag"] = spanMirrInfo.V6AclFlag
    leafs["gre-acl-flag"] = spanMirrInfo.GreAclFlag
    leafs["v4state"] = spanMirrInfo.V4State
    leafs["v6state"] = spanMirrInfo.V6State
    leafs["gre-state"] = spanMirrInfo.GreState
    leafs["v4sessid"] = spanMirrInfo.V4Sessid
    leafs["v6sessid"] = spanMirrInfo.V6Sessid
    leafs["gre-sessid"] = spanMirrInfo.GreSessid
    leafs["v4dst-type"] = spanMirrInfo.V4DstType
    leafs["v6dst-type"] = spanMirrInfo.V6DstType
    leafs["gredst-type"] = spanMirrInfo.GredstType
    leafs["v4statsptr"] = spanMirrInfo.V4Statsptr
    leafs["v6statsptr"] = spanMirrInfo.V6Statsptr
    leafs["grev4statsptr"] = spanMirrInfo.Grev4Statsptr
    leafs["grev6statsptr"] = spanMirrInfo.Grev6Statsptr
    leafs["mplsv4stats"] = spanMirrInfo.Mplsv4Stats
    leafs["mplsv6pkts"] = spanMirrInfo.Mplsv6Pkts
    leafs["np-umask"] = spanMirrInfo.NpUmask
    leafs["uidb"] = spanMirrInfo.Uidb
    leafs["sft-hw-data"] = spanMirrInfo.SftHwData
    return leafs
}

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetBundleName() string { return "cisco_ios_xr" }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetYangName() string { return "span-mirr-info" }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) SetParent(parent types.Entity) { spanMirrInfo.parent = parent }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetParent() types.Entity { return spanMirrInfo.parent }

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetParentYangName() string { return "span-mirr-infos" }

// Ssespan_Nodes_Node_Spanudf
// udf info
type Ssespan_Nodes_Node_Spanudf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // udf header. The type is slice of interface{} with range: 0..4294967295.
    UdfHdr []interface{}

    // udf type. The type is slice of interface{} with range: 0..4294967295.
    UdfType []interface{}

    // udf len. The type is slice of interface{} with range: 0..4294967295.
    UdfLen []interface{}

    // udf value. The type is slice of interface{} with range: 0..4294967295.
    UdfValue []interface{}

    // 16x5npu=80 bytes of hw udf entry. The type is slice of interface{} with
    // range: 0..4294967295.
    UdfHwData []interface{}
}

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetFilter() yfilter.YFilter { return spanudf.YFilter }

func (spanudf *Ssespan_Nodes_Node_Spanudf) SetFilter(yf yfilter.YFilter) { spanudf.YFilter = yf }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetGoName(yname string) string {
    if yname == "udf-hdr" { return "UdfHdr" }
    if yname == "udf-type" { return "UdfType" }
    if yname == "udf-len" { return "UdfLen" }
    if yname == "udf-value" { return "UdfValue" }
    if yname == "udf-hw-data" { return "UdfHwData" }
    return ""
}

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetSegmentPath() string {
    return "spanudf"
}

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udf-hdr"] = spanudf.UdfHdr
    leafs["udf-type"] = spanudf.UdfType
    leafs["udf-len"] = spanudf.UdfLen
    leafs["udf-value"] = spanudf.UdfValue
    leafs["udf-hw-data"] = spanudf.UdfHwData
    return leafs
}

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetBundleName() string { return "cisco_ios_xr" }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetYangName() string { return "spanudf" }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanudf *Ssespan_Nodes_Node_Spanudf) SetParent(parent types.Entity) { spanudf.parent = parent }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetParent() types.Entity { return spanudf.parent }

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetParentYangName() string { return "node" }

// Ssespan_Nodes_Node_SpanSessInfos
// SPAN SDT entry
type Ssespan_Nodes_Node_SpanSessInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session info . The type is slice of
    // Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo.
    SpanSessInfo []Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetFilter() yfilter.YFilter { return spanSessInfos.YFilter }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) SetFilter(yf yfilter.YFilter) { spanSessInfos.YFilter = yf }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetGoName(yname string) string {
    if yname == "span-sess-info" { return "SpanSessInfo" }
    return ""
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetSegmentPath() string {
    return "span-sess-infos"
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "span-sess-info" {
        for _, c := range spanSessInfos.SpanSessInfo {
            if spanSessInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo{}
        spanSessInfos.SpanSessInfo = append(spanSessInfos.SpanSessInfo, child)
        return &spanSessInfos.SpanSessInfo[len(spanSessInfos.SpanSessInfo)-1]
    }
    return nil
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spanSessInfos.SpanSessInfo {
        children[spanSessInfos.SpanSessInfo[i].GetSegmentPath()] = &spanSessInfos.SpanSessInfo[i]
    }
    return children
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetBundleName() string { return "cisco_ios_xr" }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetYangName() string { return "span-sess-infos" }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) SetParent(parent types.Entity) { spanSessInfos.parent = parent }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetParent() types.Entity { return spanSessInfos.parent }

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetParentYangName() string { return "node" }

// Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo
// Session info 
type Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // This attribute is a key. Session class. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionClass interface{}

    // marks validity of entry. The type is interface{} with range: 0..255.
    Valid interface{}

    // Numerical ID assigned to session. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // session state. The type is interface{} with range: 0..4294967295.
    State interface{}

    // session Class gre,ipv4,ipv6. The type is interface{} with range:
    // 0..4294967295.
    Class interface{}

    // ifhandle of interface. The type is interface{} with range: 0..4294967295.
    Ifhandle interface{}

    // Tunnel mode. The type is interface{} with range: 0..4294967295.
    Mode interface{}

    // IP type v4 or v6 . The type is interface{} with range: 0..4294967295.
    IpType interface{}

    // Vrf Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // type of service. The type is interface{} with range: 0..4294967295.
    TosBit interface{}

    // type of service copied. The type is interface{} with range: 0..4294967295.
    TosBitCopied interface{}

    // TTL. The type is interface{} with range: 0..4294967295.
    Ttl interface{}

    // DF bit. The type is interface{} with range: 0..4294967295.
    Dfbit interface{}

    // src ip v4 or v6. The type is slice of interface{} with range:
    // 0..4294967295.
    SrcIp []interface{}

    // dst ip v4 or v6. The type is slice of interface{} with range:
    // 0..4294967295.
    DstIp []interface{}

    // 16x5npu=80 bytes of hw sdt entry. The type is slice of interface{} with
    // range: 0..4294967295.
    SdtHwData []interface{}
}

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetFilter() yfilter.YFilter { return spanSessInfo.YFilter }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) SetFilter(yf yfilter.YFilter) { spanSessInfo.YFilter = yf }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "session-class" { return "SessionClass" }
    if yname == "valid" { return "Valid" }
    if yname == "id" { return "Id" }
    if yname == "state" { return "State" }
    if yname == "class" { return "Class" }
    if yname == "ifhandle" { return "Ifhandle" }
    if yname == "mode" { return "Mode" }
    if yname == "ip-type" { return "IpType" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "tos-bit" { return "TosBit" }
    if yname == "tos-bit-copied" { return "TosBitCopied" }
    if yname == "ttl" { return "Ttl" }
    if yname == "dfbit" { return "Dfbit" }
    if yname == "src-ip" { return "SrcIp" }
    if yname == "dst-ip" { return "DstIp" }
    if yname == "sdt-hw-data" { return "SdtHwData" }
    return ""
}

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetSegmentPath() string {
    return "span-sess-info" + "[session-id='" + fmt.Sprintf("%v", spanSessInfo.SessionId) + "']" + "[session-class='" + fmt.Sprintf("%v", spanSessInfo.SessionClass) + "']"
}

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = spanSessInfo.SessionId
    leafs["session-class"] = spanSessInfo.SessionClass
    leafs["valid"] = spanSessInfo.Valid
    leafs["id"] = spanSessInfo.Id
    leafs["state"] = spanSessInfo.State
    leafs["class"] = spanSessInfo.Class
    leafs["ifhandle"] = spanSessInfo.Ifhandle
    leafs["mode"] = spanSessInfo.Mode
    leafs["ip-type"] = spanSessInfo.IpType
    leafs["vrf-id"] = spanSessInfo.VrfId
    leafs["tos-bit"] = spanSessInfo.TosBit
    leafs["tos-bit-copied"] = spanSessInfo.TosBitCopied
    leafs["ttl"] = spanSessInfo.Ttl
    leafs["dfbit"] = spanSessInfo.Dfbit
    leafs["src-ip"] = spanSessInfo.SrcIp
    leafs["dst-ip"] = spanSessInfo.DstIp
    leafs["sdt-hw-data"] = spanSessInfo.SdtHwData
    return leafs
}

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetBundleName() string { return "cisco_ios_xr" }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetYangName() string { return "span-sess-info" }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) SetParent(parent types.Entity) { spanSessInfo.parent = parent }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetParent() types.Entity { return spanSessInfo.parent }

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetParentYangName() string { return "span-sess-infos" }

