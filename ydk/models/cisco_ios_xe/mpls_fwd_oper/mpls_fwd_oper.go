// This module contains a collection of YANG definitions for
// monitoring memory usage of processes in a Network Element.Copyright (c) 2016-2017 by Cisco Systems, Inc.All rights reserved.
package mpls_fwd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_fwd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mpls-fwd-oper mpls-forwarding-table}", reflect.TypeOf(MplsForwardingTable{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mpls-fwd-oper:mpls-forwarding-table", reflect.TypeOf(MplsForwardingTable{}))
}

// MplsForwardingTable
// Data nodes for MPLS forwarding table entries.
type MplsForwardingTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of MPLS forwarding table entries. The type is slice of
    // MplsForwardingTable_LocalLabelEntry.
    LocalLabelEntry []MplsForwardingTable_LocalLabelEntry
}

func (mplsForwardingTable *MplsForwardingTable) GetFilter() yfilter.YFilter { return mplsForwardingTable.YFilter }

func (mplsForwardingTable *MplsForwardingTable) SetFilter(yf yfilter.YFilter) { mplsForwardingTable.YFilter = yf }

func (mplsForwardingTable *MplsForwardingTable) GetGoName(yname string) string {
    if yname == "local-label-entry" { return "LocalLabelEntry" }
    return ""
}

func (mplsForwardingTable *MplsForwardingTable) GetSegmentPath() string {
    return "Cisco-IOS-XE-mpls-fwd-oper:mpls-forwarding-table"
}

func (mplsForwardingTable *MplsForwardingTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label-entry" {
        for _, c := range mplsForwardingTable.LocalLabelEntry {
            if mplsForwardingTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsForwardingTable_LocalLabelEntry{}
        mplsForwardingTable.LocalLabelEntry = append(mplsForwardingTable.LocalLabelEntry, child)
        return &mplsForwardingTable.LocalLabelEntry[len(mplsForwardingTable.LocalLabelEntry)-1]
    }
    return nil
}

func (mplsForwardingTable *MplsForwardingTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsForwardingTable.LocalLabelEntry {
        children[mplsForwardingTable.LocalLabelEntry[i].GetSegmentPath()] = &mplsForwardingTable.LocalLabelEntry[i]
    }
    return children
}

func (mplsForwardingTable *MplsForwardingTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsForwardingTable *MplsForwardingTable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsForwardingTable *MplsForwardingTable) GetYangName() string { return "mpls-forwarding-table" }

func (mplsForwardingTable *MplsForwardingTable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsForwardingTable *MplsForwardingTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsForwardingTable *MplsForwardingTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsForwardingTable *MplsForwardingTable) SetParent(parent types.Entity) { mplsForwardingTable.parent = parent }

func (mplsForwardingTable *MplsForwardingTable) GetParent() types.Entity { return mplsForwardingTable.parent }

func (mplsForwardingTable *MplsForwardingTable) GetParentYangName() string { return "Cisco-IOS-XE-mpls-fwd-oper" }

// MplsForwardingTable_LocalLabelEntry
// The list of MPLS forwarding table entries.
type MplsForwardingTable_LocalLabelEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Value of local-label. The type is interface{} with
    // range: 0..4294967295.
    LocalLabel interface{}

    // The type is slice of MplsForwardingTable_LocalLabelEntry_ForwardingInfo.
    ForwardingInfo []MplsForwardingTable_LocalLabelEntry_ForwardingInfo
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetFilter() yfilter.YFilter { return localLabelEntry.YFilter }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) SetFilter(yf yfilter.YFilter) { localLabelEntry.YFilter = yf }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    if yname == "forwarding-info" { return "ForwardingInfo" }
    return ""
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetSegmentPath() string {
    return "local-label-entry" + "[local-label='" + fmt.Sprintf("%v", localLabelEntry.LocalLabel) + "']"
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forwarding-info" {
        for _, c := range localLabelEntry.ForwardingInfo {
            if localLabelEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsForwardingTable_LocalLabelEntry_ForwardingInfo{}
        localLabelEntry.ForwardingInfo = append(localLabelEntry.ForwardingInfo, child)
        return &localLabelEntry.ForwardingInfo[len(localLabelEntry.ForwardingInfo)-1]
    }
    return nil
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabelEntry.ForwardingInfo {
        children[localLabelEntry.ForwardingInfo[i].GetSegmentPath()] = &localLabelEntry.ForwardingInfo[i]
    }
    return children
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label"] = localLabelEntry.LocalLabel
    return leafs
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetBundleName() string { return "cisco_ios_xe" }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetYangName() string { return "local-label-entry" }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) SetParent(parent types.Entity) { localLabelEntry.parent = parent }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetParent() types.Entity { return localLabelEntry.parent }

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetParentYangName() string { return "mpls-forwarding-table" }

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the outgoing interface. Example
    // possible values are 1.none, 2.drop, 3.<tunnel-name>, 4.<interface-name>,
    // 5.aggregate/<vrf-name> etc. The type is one of the following types:
    // enumeration
    // MplsForwardingTable.LocalLabelEntry.ForwardingInfo.OutgoingInterface, or
    // string.
    OutgoingInterface interface{}

    // Value of outgoing-label if exists or the type of non-present label. The
    // type is one of the following types: int with range: 0..4294967295, or
    // enumeration
    // MplsForwardingTable.LocalLabelEntry.ForwardingInfo.OutgoingLabel.
    OutgoingLabel interface{}

    // The number of bytes switched using this label. The type is interface{} with
    // range: 0..18446744073709551615.
    LabelSwitchedBytes interface{}

    // Next hop information. Example possible values are 1.point2point,
    // 2.<ip-address>. The type is one of the following types: enumeration
    // MplsForwardingTable.LocalLabelEntry.ForwardingInfo.NextHop, or string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    NextHop interface{}

    // The Prefix or tunnel-id info corresponding to this label. Ex: 1) for l2ckt,
    // a number tunnel-id value. 2) for ipv4, a prefix with [V] tag
    // (113.113.113.113/32[V]). 3) for TE, a pefix with [T] tag
    // (113.113.113.113/32[T]).
    ConnectionInfo MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetFilter() yfilter.YFilter { return forwardingInfo.YFilter }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) SetFilter(yf yfilter.YFilter) { forwardingInfo.YFilter = yf }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetGoName(yname string) string {
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    if yname == "outgoing-label" { return "OutgoingLabel" }
    if yname == "label-switched-bytes" { return "LabelSwitchedBytes" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "connection-info" { return "ConnectionInfo" }
    return ""
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetSegmentPath() string {
    return "forwarding-info" + "[outgoing-interface='" + fmt.Sprintf("%v", forwardingInfo.OutgoingInterface) + "']"
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connection-info" {
        return &forwardingInfo.ConnectionInfo
    }
    return nil
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connection-info"] = &forwardingInfo.ConnectionInfo
    return children
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outgoing-interface"] = forwardingInfo.OutgoingInterface
    leafs["outgoing-label"] = forwardingInfo.OutgoingLabel
    leafs["label-switched-bytes"] = forwardingInfo.LabelSwitchedBytes
    leafs["next-hop"] = forwardingInfo.NextHop
    return leafs
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetBundleName() string { return "cisco_ios_xe" }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetYangName() string { return "forwarding-info" }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) SetParent(parent types.Entity) { forwardingInfo.parent = parent }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetParent() types.Entity { return forwardingInfo.parent }

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetParentYangName() string { return "local-label-entry" }

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo
// The Prefix or tunnel-id info corresponding to this label.
// Ex: 1) for l2ckt, a number tunnel-id value.
// 2) for ipv4, a prefix with [V] tag (113.113.113.113/32[V]).
// 3) for TE, a pefix with [T] tag (113.113.113.113/32[T])
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type of connection represented by this label. The type is Type.
    Type interface{}

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Mask interface{}

    // The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // The type is interface{} with range: 0..4294967295.
    NhId interface{}

    // The type is interface{} with range: 0..4294967295.
    L2CktId interface{}

    
    TunnelTp MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp
}

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetFilter() yfilter.YFilter { return connectionInfo.YFilter }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) SetFilter(yf yfilter.YFilter) { connectionInfo.YFilter = yf }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ip" { return "Ip" }
    if yname == "mask" { return "Mask" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "nh-id" { return "NhId" }
    if yname == "l2ckt-id" { return "L2CktId" }
    if yname == "tunnel-tp" { return "TunnelTp" }
    return ""
}

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetSegmentPath() string {
    return "connection-info"
}

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-tp" {
        return &connectionInfo.TunnelTp
    }
    return nil
}

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnel-tp"] = &connectionInfo.TunnelTp
    return children
}

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = connectionInfo.Type
    leafs["ip"] = connectionInfo.Ip
    leafs["mask"] = connectionInfo.Mask
    leafs["tunnel-id"] = connectionInfo.TunnelId
    leafs["vrf-id"] = connectionInfo.VrfId
    leafs["nh-id"] = connectionInfo.NhId
    leafs["l2ckt-id"] = connectionInfo.L2CktId
    return leafs
}

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetBundleName() string { return "cisco_ios_xe" }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetYangName() string { return "connection-info" }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) SetParent(parent types.Entity) { connectionInfo.parent = parent }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetParent() types.Entity { return connectionInfo.parent }

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetParentYangName() string { return "forwarding-info" }

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Tunnel interface{}

    
    SrcId MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId

    
    DstId MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetFilter() yfilter.YFilter { return tunnelTp.YFilter }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) SetFilter(yf yfilter.YFilter) { tunnelTp.YFilter = yf }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    if yname == "src-id" { return "SrcId" }
    if yname == "dst-id" { return "DstId" }
    return ""
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetSegmentPath() string {
    return "tunnel-tp"
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "src-id" {
        return &tunnelTp.SrcId
    }
    if childYangName == "dst-id" {
        return &tunnelTp.DstId
    }
    return nil
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["src-id"] = &tunnelTp.SrcId
    children["dst-id"] = &tunnelTp.DstId
    return children
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel"] = tunnelTp.Tunnel
    return leafs
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetBundleName() string { return "cisco_ios_xe" }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetYangName() string { return "tunnel-tp" }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) SetParent(parent types.Entity) { tunnelTp.parent = parent }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetParent() types.Entity { return tunnelTp.parent }

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetParentYangName() string { return "connection-info" }

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Global interface{}

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Node interface{}
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetFilter() yfilter.YFilter { return srcId.YFilter }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) SetFilter(yf yfilter.YFilter) { srcId.YFilter = yf }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "node" { return "Node" }
    return ""
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetSegmentPath() string {
    return "src-id"
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global"] = srcId.Global
    leafs["node"] = srcId.Node
    return leafs
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetBundleName() string { return "cisco_ios_xe" }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetYangName() string { return "src-id" }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) SetParent(parent types.Entity) { srcId.parent = parent }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetParent() types.Entity { return srcId.parent }

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetParentYangName() string { return "tunnel-tp" }

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Global interface{}

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Node interface{}
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetFilter() yfilter.YFilter { return dstId.YFilter }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) SetFilter(yf yfilter.YFilter) { dstId.YFilter = yf }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "node" { return "Node" }
    return ""
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetSegmentPath() string {
    return "dst-id"
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global"] = dstId.Global
    leafs["node"] = dstId.Node
    return leafs
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetBundleName() string { return "cisco_ios_xe" }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetYangName() string { return "dst-id" }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) SetParent(parent types.Entity) { dstId.parent = parent }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetParent() types.Entity { return dstId.parent }

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetParentYangName() string { return "tunnel-tp" }

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type represents The type of connection represented by this label
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type string

const (
    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ip MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "ip"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_tunnel MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "tunnel"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_nh_id MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "nh-id"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_l2ckt MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "l2ckt"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ip_vrf MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "ip-vrf"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_none MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "none"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_tunnel_tp MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type = "tunnel-tp"
)

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_NextHop represents 2.<ip-address>
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_NextHop string

const (
    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_NextHop_point2point MplsForwardingTable_LocalLabelEntry_ForwardingInfo_NextHop = "point2point"
)

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface represents 4.<interface-name>, 5.aggregate/<vrf-name> etc.
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface string

const (
    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface_drop MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface = "drop"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface_punt MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface = "punt"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface_aggregate MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface = "aggregate"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface_exception MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface = "exception"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface_none MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingInterface = "none"
)

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel represents the type of non-present label.
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel string

const (
    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel_no_label MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel = "no-label"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel_pop_label MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel = "pop-label"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel_aggregate MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel = "aggregate"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel_explicit_null MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel = "explicit-null"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel_illegal MplsForwardingTable_LocalLabelEntry_ForwardingInfo_OutgoingLabel = "illegal"
)

