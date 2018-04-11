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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node table for node-specific operational data.
    Nodes Ssespan_Nodes
}

func (ssespan *Ssespan) GetEntityData() *types.CommonEntityData {
    ssespan.EntityData.YFilter = ssespan.YFilter
    ssespan.EntityData.YangName = "ssespan"
    ssespan.EntityData.BundleName = "cisco_ios_xr"
    ssespan.EntityData.ParentYangName = "Cisco-IOS-XR-sse-span-oper"
    ssespan.EntityData.SegmentPath = "Cisco-IOS-XR-sse-span-oper:ssespan"
    ssespan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssespan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssespan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssespan.EntityData.Children = make(map[string]types.YChild)
    ssespan.EntityData.Children["nodes"] = types.YChild{"Nodes", &ssespan.Nodes}
    ssespan.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ssespan.EntityData)
}

// Ssespan_Nodes
// Node table for node-specific operational data
type Ssespan_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // Ssespan_Nodes_Node.
    Node []Ssespan_Nodes_Node
}

func (nodes *Ssespan_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ssespan"
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

// Ssespan_Nodes_Node
// Node-specific data for a particular node
type Ssespan_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // SPAN SFT entry.
    SpanMirrInfos Ssespan_Nodes_Node_SpanMirrInfos

    // udf info.
    Spanudf Ssespan_Nodes_Node_Spanudf

    // SPAN SDT entry.
    SpanSessInfos Ssespan_Nodes_Node_SpanSessInfos
}

func (node *Ssespan_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["span-mirr-infos"] = types.YChild{"SpanMirrInfos", &node.SpanMirrInfos}
    node.EntityData.Children["spanudf"] = types.YChild{"Spanudf", &node.Spanudf}
    node.EntityData.Children["span-sess-infos"] = types.YChild{"SpanSessInfos", &node.SpanSessInfos}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node"] = types.YLeaf{"Node", node.Node}
    return &(node.EntityData)
}

// Ssespan_Nodes_Node_SpanMirrInfos
// SPAN SFT entry
type Ssespan_Nodes_Node_SpanMirrInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mirror info . The type is slice of
    // Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo.
    SpanMirrInfo []Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo
}

func (spanMirrInfos *Ssespan_Nodes_Node_SpanMirrInfos) GetEntityData() *types.CommonEntityData {
    spanMirrInfos.EntityData.YFilter = spanMirrInfos.YFilter
    spanMirrInfos.EntityData.YangName = "span-mirr-infos"
    spanMirrInfos.EntityData.BundleName = "cisco_ios_xr"
    spanMirrInfos.EntityData.ParentYangName = "node"
    spanMirrInfos.EntityData.SegmentPath = "span-mirr-infos"
    spanMirrInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMirrInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMirrInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMirrInfos.EntityData.Children = make(map[string]types.YChild)
    spanMirrInfos.EntityData.Children["span-mirr-info"] = types.YChild{"SpanMirrInfo", nil}
    for i := range spanMirrInfos.SpanMirrInfo {
        spanMirrInfos.EntityData.Children[types.GetSegmentPath(&spanMirrInfos.SpanMirrInfo[i])] = types.YChild{"SpanMirrInfo", &spanMirrInfos.SpanMirrInfo[i]}
    }
    spanMirrInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spanMirrInfos.EntityData)
}

// Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo
// Mirror info 
type Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    IntfName interface{}

    // source IFH. The type is interface{} with range: 0..4294967295.
    SrcIfh interface{}

    // interface name. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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

func (spanMirrInfo *Ssespan_Nodes_Node_SpanMirrInfos_SpanMirrInfo) GetEntityData() *types.CommonEntityData {
    spanMirrInfo.EntityData.YFilter = spanMirrInfo.YFilter
    spanMirrInfo.EntityData.YangName = "span-mirr-info"
    spanMirrInfo.EntityData.BundleName = "cisco_ios_xr"
    spanMirrInfo.EntityData.ParentYangName = "span-mirr-infos"
    spanMirrInfo.EntityData.SegmentPath = "span-mirr-info" + "[intf-name='" + fmt.Sprintf("%v", spanMirrInfo.IntfName) + "']"
    spanMirrInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMirrInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMirrInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMirrInfo.EntityData.Children = make(map[string]types.YChild)
    spanMirrInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    spanMirrInfo.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", spanMirrInfo.IntfName}
    spanMirrInfo.EntityData.Leafs["src-ifh"] = types.YLeaf{"SrcIfh", spanMirrInfo.SrcIfh}
    spanMirrInfo.EntityData.Leafs["intf-name-xr"] = types.YLeaf{"IntfNameXr", spanMirrInfo.IntfNameXr}
    spanMirrInfo.EntityData.Leafs["v4-acl-flag"] = types.YLeaf{"V4AclFlag", spanMirrInfo.V4AclFlag}
    spanMirrInfo.EntityData.Leafs["v6-acl-flag"] = types.YLeaf{"V6AclFlag", spanMirrInfo.V6AclFlag}
    spanMirrInfo.EntityData.Leafs["gre-acl-flag"] = types.YLeaf{"GreAclFlag", spanMirrInfo.GreAclFlag}
    spanMirrInfo.EntityData.Leafs["v4state"] = types.YLeaf{"V4State", spanMirrInfo.V4State}
    spanMirrInfo.EntityData.Leafs["v6state"] = types.YLeaf{"V6State", spanMirrInfo.V6State}
    spanMirrInfo.EntityData.Leafs["gre-state"] = types.YLeaf{"GreState", spanMirrInfo.GreState}
    spanMirrInfo.EntityData.Leafs["v4sessid"] = types.YLeaf{"V4Sessid", spanMirrInfo.V4Sessid}
    spanMirrInfo.EntityData.Leafs["v6sessid"] = types.YLeaf{"V6Sessid", spanMirrInfo.V6Sessid}
    spanMirrInfo.EntityData.Leafs["gre-sessid"] = types.YLeaf{"GreSessid", spanMirrInfo.GreSessid}
    spanMirrInfo.EntityData.Leafs["v4dst-type"] = types.YLeaf{"V4DstType", spanMirrInfo.V4DstType}
    spanMirrInfo.EntityData.Leafs["v6dst-type"] = types.YLeaf{"V6DstType", spanMirrInfo.V6DstType}
    spanMirrInfo.EntityData.Leafs["gredst-type"] = types.YLeaf{"GredstType", spanMirrInfo.GredstType}
    spanMirrInfo.EntityData.Leafs["v4statsptr"] = types.YLeaf{"V4Statsptr", spanMirrInfo.V4Statsptr}
    spanMirrInfo.EntityData.Leafs["v6statsptr"] = types.YLeaf{"V6Statsptr", spanMirrInfo.V6Statsptr}
    spanMirrInfo.EntityData.Leafs["grev4statsptr"] = types.YLeaf{"Grev4Statsptr", spanMirrInfo.Grev4Statsptr}
    spanMirrInfo.EntityData.Leafs["grev6statsptr"] = types.YLeaf{"Grev6Statsptr", spanMirrInfo.Grev6Statsptr}
    spanMirrInfo.EntityData.Leafs["mplsv4stats"] = types.YLeaf{"Mplsv4Stats", spanMirrInfo.Mplsv4Stats}
    spanMirrInfo.EntityData.Leafs["mplsv6pkts"] = types.YLeaf{"Mplsv6Pkts", spanMirrInfo.Mplsv6Pkts}
    spanMirrInfo.EntityData.Leafs["np-umask"] = types.YLeaf{"NpUmask", spanMirrInfo.NpUmask}
    spanMirrInfo.EntityData.Leafs["uidb"] = types.YLeaf{"Uidb", spanMirrInfo.Uidb}
    spanMirrInfo.EntityData.Leafs["sft-hw-data"] = types.YLeaf{"SftHwData", spanMirrInfo.SftHwData}
    return &(spanMirrInfo.EntityData)
}

// Ssespan_Nodes_Node_Spanudf
// udf info
type Ssespan_Nodes_Node_Spanudf struct {
    EntityData types.CommonEntityData
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

func (spanudf *Ssespan_Nodes_Node_Spanudf) GetEntityData() *types.CommonEntityData {
    spanudf.EntityData.YFilter = spanudf.YFilter
    spanudf.EntityData.YangName = "spanudf"
    spanudf.EntityData.BundleName = "cisco_ios_xr"
    spanudf.EntityData.ParentYangName = "node"
    spanudf.EntityData.SegmentPath = "spanudf"
    spanudf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanudf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanudf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanudf.EntityData.Children = make(map[string]types.YChild)
    spanudf.EntityData.Leafs = make(map[string]types.YLeaf)
    spanudf.EntityData.Leafs["udf-hdr"] = types.YLeaf{"UdfHdr", spanudf.UdfHdr}
    spanudf.EntityData.Leafs["udf-type"] = types.YLeaf{"UdfType", spanudf.UdfType}
    spanudf.EntityData.Leafs["udf-len"] = types.YLeaf{"UdfLen", spanudf.UdfLen}
    spanudf.EntityData.Leafs["udf-value"] = types.YLeaf{"UdfValue", spanudf.UdfValue}
    spanudf.EntityData.Leafs["udf-hw-data"] = types.YLeaf{"UdfHwData", spanudf.UdfHwData}
    return &(spanudf.EntityData)
}

// Ssespan_Nodes_Node_SpanSessInfos
// SPAN SDT entry
type Ssespan_Nodes_Node_SpanSessInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session info . The type is slice of
    // Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo.
    SpanSessInfo []Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo
}

func (spanSessInfos *Ssespan_Nodes_Node_SpanSessInfos) GetEntityData() *types.CommonEntityData {
    spanSessInfos.EntityData.YFilter = spanSessInfos.YFilter
    spanSessInfos.EntityData.YangName = "span-sess-infos"
    spanSessInfos.EntityData.BundleName = "cisco_ios_xr"
    spanSessInfos.EntityData.ParentYangName = "node"
    spanSessInfos.EntityData.SegmentPath = "span-sess-infos"
    spanSessInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanSessInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanSessInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanSessInfos.EntityData.Children = make(map[string]types.YChild)
    spanSessInfos.EntityData.Children["span-sess-info"] = types.YChild{"SpanSessInfo", nil}
    for i := range spanSessInfos.SpanSessInfo {
        spanSessInfos.EntityData.Children[types.GetSegmentPath(&spanSessInfos.SpanSessInfo[i])] = types.YChild{"SpanSessInfo", &spanSessInfos.SpanSessInfo[i]}
    }
    spanSessInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spanSessInfos.EntityData)
}

// Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo
// Session info 
type Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo struct {
    EntityData types.CommonEntityData
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

func (spanSessInfo *Ssespan_Nodes_Node_SpanSessInfos_SpanSessInfo) GetEntityData() *types.CommonEntityData {
    spanSessInfo.EntityData.YFilter = spanSessInfo.YFilter
    spanSessInfo.EntityData.YangName = "span-sess-info"
    spanSessInfo.EntityData.BundleName = "cisco_ios_xr"
    spanSessInfo.EntityData.ParentYangName = "span-sess-infos"
    spanSessInfo.EntityData.SegmentPath = "span-sess-info" + "[session-id='" + fmt.Sprintf("%v", spanSessInfo.SessionId) + "']" + "[session-class='" + fmt.Sprintf("%v", spanSessInfo.SessionClass) + "']"
    spanSessInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanSessInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanSessInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanSessInfo.EntityData.Children = make(map[string]types.YChild)
    spanSessInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    spanSessInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", spanSessInfo.SessionId}
    spanSessInfo.EntityData.Leafs["session-class"] = types.YLeaf{"SessionClass", spanSessInfo.SessionClass}
    spanSessInfo.EntityData.Leafs["valid"] = types.YLeaf{"Valid", spanSessInfo.Valid}
    spanSessInfo.EntityData.Leafs["id"] = types.YLeaf{"Id", spanSessInfo.Id}
    spanSessInfo.EntityData.Leafs["state"] = types.YLeaf{"State", spanSessInfo.State}
    spanSessInfo.EntityData.Leafs["class"] = types.YLeaf{"Class", spanSessInfo.Class}
    spanSessInfo.EntityData.Leafs["ifhandle"] = types.YLeaf{"Ifhandle", spanSessInfo.Ifhandle}
    spanSessInfo.EntityData.Leafs["mode"] = types.YLeaf{"Mode", spanSessInfo.Mode}
    spanSessInfo.EntityData.Leafs["ip-type"] = types.YLeaf{"IpType", spanSessInfo.IpType}
    spanSessInfo.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", spanSessInfo.VrfId}
    spanSessInfo.EntityData.Leafs["tos-bit"] = types.YLeaf{"TosBit", spanSessInfo.TosBit}
    spanSessInfo.EntityData.Leafs["tos-bit-copied"] = types.YLeaf{"TosBitCopied", spanSessInfo.TosBitCopied}
    spanSessInfo.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", spanSessInfo.Ttl}
    spanSessInfo.EntityData.Leafs["dfbit"] = types.YLeaf{"Dfbit", spanSessInfo.Dfbit}
    spanSessInfo.EntityData.Leafs["src-ip"] = types.YLeaf{"SrcIp", spanSessInfo.SrcIp}
    spanSessInfo.EntityData.Leafs["dst-ip"] = types.YLeaf{"DstIp", spanSessInfo.DstIp}
    spanSessInfo.EntityData.Leafs["sdt-hw-data"] = types.YLeaf{"SdtHwData", spanSessInfo.SdtHwData}
    return &(spanSessInfo.EntityData)
}

