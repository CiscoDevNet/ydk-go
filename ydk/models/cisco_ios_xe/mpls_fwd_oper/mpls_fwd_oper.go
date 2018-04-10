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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of MPLS forwarding table entries. The type is slice of
    // MplsForwardingTable_LocalLabelEntry.
    LocalLabelEntry []MplsForwardingTable_LocalLabelEntry
}

func (mplsForwardingTable *MplsForwardingTable) GetEntityData() *types.CommonEntityData {
    mplsForwardingTable.EntityData.YFilter = mplsForwardingTable.YFilter
    mplsForwardingTable.EntityData.YangName = "mpls-forwarding-table"
    mplsForwardingTable.EntityData.BundleName = "cisco_ios_xe"
    mplsForwardingTable.EntityData.ParentYangName = "Cisco-IOS-XE-mpls-fwd-oper"
    mplsForwardingTable.EntityData.SegmentPath = "Cisco-IOS-XE-mpls-fwd-oper:mpls-forwarding-table"
    mplsForwardingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsForwardingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsForwardingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsForwardingTable.EntityData.Children = make(map[string]types.YChild)
    mplsForwardingTable.EntityData.Children["local-label-entry"] = types.YChild{"LocalLabelEntry", nil}
    for i := range mplsForwardingTable.LocalLabelEntry {
        mplsForwardingTable.EntityData.Children[types.GetSegmentPath(&mplsForwardingTable.LocalLabelEntry[i])] = types.YChild{"LocalLabelEntry", &mplsForwardingTable.LocalLabelEntry[i]}
    }
    mplsForwardingTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsForwardingTable.EntityData)
}

// MplsForwardingTable_LocalLabelEntry
// The list of MPLS forwarding table entries.
type MplsForwardingTable_LocalLabelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Value of local-label. The type is interface{} with
    // range: 0..4294967295.
    LocalLabel interface{}

    // The type is slice of MplsForwardingTable_LocalLabelEntry_ForwardingInfo.
    ForwardingInfo []MplsForwardingTable_LocalLabelEntry_ForwardingInfo
}

func (localLabelEntry *MplsForwardingTable_LocalLabelEntry) GetEntityData() *types.CommonEntityData {
    localLabelEntry.EntityData.YFilter = localLabelEntry.YFilter
    localLabelEntry.EntityData.YangName = "local-label-entry"
    localLabelEntry.EntityData.BundleName = "cisco_ios_xe"
    localLabelEntry.EntityData.ParentYangName = "mpls-forwarding-table"
    localLabelEntry.EntityData.SegmentPath = "local-label-entry" + "[local-label='" + fmt.Sprintf("%v", localLabelEntry.LocalLabel) + "']"
    localLabelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    localLabelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    localLabelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    localLabelEntry.EntityData.Children = make(map[string]types.YChild)
    localLabelEntry.EntityData.Children["forwarding-info"] = types.YChild{"ForwardingInfo", nil}
    for i := range localLabelEntry.ForwardingInfo {
        localLabelEntry.EntityData.Children[types.GetSegmentPath(&localLabelEntry.ForwardingInfo[i])] = types.YChild{"ForwardingInfo", &localLabelEntry.ForwardingInfo[i]}
    }
    localLabelEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    localLabelEntry.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", localLabelEntry.LocalLabel}
    return &(localLabelEntry.EntityData)
}

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    NextHop interface{}

    // The Prefix or tunnel-id info corresponding to this label. Ex: 1) for l2ckt,
    // a number tunnel-id value. 2) for ipv4, a prefix with [V] tag
    // (113.113.113.113/32[V]). 3) for TE, a pefix with [T] tag
    // (113.113.113.113/32[T]).
    ConnectionInfo MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo
}

func (forwardingInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo) GetEntityData() *types.CommonEntityData {
    forwardingInfo.EntityData.YFilter = forwardingInfo.YFilter
    forwardingInfo.EntityData.YangName = "forwarding-info"
    forwardingInfo.EntityData.BundleName = "cisco_ios_xe"
    forwardingInfo.EntityData.ParentYangName = "local-label-entry"
    forwardingInfo.EntityData.SegmentPath = "forwarding-info" + "[outgoing-interface='" + fmt.Sprintf("%v", forwardingInfo.OutgoingInterface) + "']"
    forwardingInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwardingInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwardingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwardingInfo.EntityData.Children = make(map[string]types.YChild)
    forwardingInfo.EntityData.Children["connection-info"] = types.YChild{"ConnectionInfo", &forwardingInfo.ConnectionInfo}
    forwardingInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    forwardingInfo.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", forwardingInfo.OutgoingInterface}
    forwardingInfo.EntityData.Leafs["outgoing-label"] = types.YLeaf{"OutgoingLabel", forwardingInfo.OutgoingLabel}
    forwardingInfo.EntityData.Leafs["label-switched-bytes"] = types.YLeaf{"LabelSwitchedBytes", forwardingInfo.LabelSwitchedBytes}
    forwardingInfo.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", forwardingInfo.NextHop}
    return &(forwardingInfo.EntityData)
}

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo
// The Prefix or tunnel-id info corresponding to this label.
// Ex: 1) for l2ckt, a number tunnel-id value.
// 2) for ipv4, a prefix with [V] tag (113.113.113.113/32[V]).
// 3) for TE, a pefix with [T] tag (113.113.113.113/32[T])
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of connection represented by this label. The type is Type_.
    Type_ interface{}

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (connectionInfo *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo) GetEntityData() *types.CommonEntityData {
    connectionInfo.EntityData.YFilter = connectionInfo.YFilter
    connectionInfo.EntityData.YangName = "connection-info"
    connectionInfo.EntityData.BundleName = "cisco_ios_xe"
    connectionInfo.EntityData.ParentYangName = "forwarding-info"
    connectionInfo.EntityData.SegmentPath = "connection-info"
    connectionInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    connectionInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    connectionInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    connectionInfo.EntityData.Children = make(map[string]types.YChild)
    connectionInfo.EntityData.Children["tunnel-tp"] = types.YChild{"TunnelTp", &connectionInfo.TunnelTp}
    connectionInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    connectionInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", connectionInfo.Type_}
    connectionInfo.EntityData.Leafs["ip"] = types.YLeaf{"Ip", connectionInfo.Ip}
    connectionInfo.EntityData.Leafs["mask"] = types.YLeaf{"Mask", connectionInfo.Mask}
    connectionInfo.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", connectionInfo.TunnelId}
    connectionInfo.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", connectionInfo.VrfId}
    connectionInfo.EntityData.Leafs["nh-id"] = types.YLeaf{"NhId", connectionInfo.NhId}
    connectionInfo.EntityData.Leafs["l2ckt-id"] = types.YLeaf{"L2CktId", connectionInfo.L2CktId}
    return &(connectionInfo.EntityData)
}

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Tunnel interface{}

    
    SrcId MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId

    
    DstId MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId
}

func (tunnelTp *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp) GetEntityData() *types.CommonEntityData {
    tunnelTp.EntityData.YFilter = tunnelTp.YFilter
    tunnelTp.EntityData.YangName = "tunnel-tp"
    tunnelTp.EntityData.BundleName = "cisco_ios_xe"
    tunnelTp.EntityData.ParentYangName = "connection-info"
    tunnelTp.EntityData.SegmentPath = "tunnel-tp"
    tunnelTp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelTp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelTp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelTp.EntityData.Children = make(map[string]types.YChild)
    tunnelTp.EntityData.Children["src-id"] = types.YChild{"SrcId", &tunnelTp.SrcId}
    tunnelTp.EntityData.Children["dst-id"] = types.YChild{"DstId", &tunnelTp.DstId}
    tunnelTp.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelTp.EntityData.Leafs["tunnel"] = types.YLeaf{"Tunnel", tunnelTp.Tunnel}
    return &(tunnelTp.EntityData)
}

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Global interface{}

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Node interface{}
}

func (srcId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_SrcId) GetEntityData() *types.CommonEntityData {
    srcId.EntityData.YFilter = srcId.YFilter
    srcId.EntityData.YangName = "src-id"
    srcId.EntityData.BundleName = "cisco_ios_xe"
    srcId.EntityData.ParentYangName = "tunnel-tp"
    srcId.EntityData.SegmentPath = "src-id"
    srcId.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    srcId.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    srcId.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    srcId.EntityData.Children = make(map[string]types.YChild)
    srcId.EntityData.Leafs = make(map[string]types.YLeaf)
    srcId.EntityData.Leafs["global"] = types.YLeaf{"Global", srcId.Global}
    srcId.EntityData.Leafs["node"] = types.YLeaf{"Node", srcId.Node}
    return &(srcId.EntityData)
}

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Global interface{}

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Node interface{}
}

func (dstId *MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_TunnelTp_DstId) GetEntityData() *types.CommonEntityData {
    dstId.EntityData.YFilter = dstId.YFilter
    dstId.EntityData.YangName = "dst-id"
    dstId.EntityData.BundleName = "cisco_ios_xe"
    dstId.EntityData.ParentYangName = "tunnel-tp"
    dstId.EntityData.SegmentPath = "dst-id"
    dstId.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dstId.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dstId.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dstId.EntityData.Children = make(map[string]types.YChild)
    dstId.EntityData.Leafs = make(map[string]types.YLeaf)
    dstId.EntityData.Leafs["global"] = types.YLeaf{"Global", dstId.Global}
    dstId.EntityData.Leafs["node"] = types.YLeaf{"Node", dstId.Node}
    return &(dstId.EntityData)
}

// MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ represents The type of connection represented by this label
type MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ string

const (
    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__ip MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "ip"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__tunnel MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "tunnel"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__nh_id MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "nh-id"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__l2ckt MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "l2ckt"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__ip_vrf MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "ip-vrf"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__none MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "none"

    MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type__tunnel_tp MplsForwardingTable_LocalLabelEntry_ForwardingInfo_ConnectionInfo_Type_ = "tunnel-tp"
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

