// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-ifib package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lpts-ifib: lpts ifib database
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_ifib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_ifib_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-ifib-oper lpts-ifib}", reflect.TypeOf(LptsIfib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-ifib-oper:lpts-ifib", reflect.TypeOf(LptsIfib{}))
}

// LptsIfib
// lpts ifib database
type LptsIfib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ifib database.
    Nodes LptsIfib_Nodes
}

func (lptsIfib *LptsIfib) GetEntityData() *types.CommonEntityData {
    lptsIfib.EntityData.YFilter = lptsIfib.YFilter
    lptsIfib.EntityData.YangName = "lpts-ifib"
    lptsIfib.EntityData.BundleName = "cisco_ios_xr"
    lptsIfib.EntityData.ParentYangName = "Cisco-IOS-XR-lpts-ifib-oper"
    lptsIfib.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-ifib-oper:lpts-ifib"
    lptsIfib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsIfib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsIfib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsIfib.EntityData.Children = make(map[string]types.YChild)
    lptsIfib.EntityData.Children["nodes"] = types.YChild{"Nodes", &lptsIfib.Nodes}
    lptsIfib.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lptsIfib.EntityData)
}

// LptsIfib_Nodes
// Node ifib database
type LptsIfib_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node slice . The type is slice of LptsIfib_Nodes_Node.
    Node []LptsIfib_Nodes_Node
}

func (nodes *LptsIfib_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "lpts-ifib"
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

// LptsIfib_Nodes_Node
// Per node slice 
type LptsIfib_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Slice specific.
    SliceIds LptsIfib_Nodes_Node_SliceIds
}

func (node *LptsIfib_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["slice-ids"] = types.YChild{"SliceIds", &node.SliceIds}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// LptsIfib_Nodes_Node_SliceIds
// Slice specific
type LptsIfib_Nodes_Node_SliceIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // slice types. The type is slice of LptsIfib_Nodes_Node_SliceIds_SliceId.
    SliceId []LptsIfib_Nodes_Node_SliceIds_SliceId
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetEntityData() *types.CommonEntityData {
    sliceIds.EntityData.YFilter = sliceIds.YFilter
    sliceIds.EntityData.YangName = "slice-ids"
    sliceIds.EntityData.BundleName = "cisco_ios_xr"
    sliceIds.EntityData.ParentYangName = "node"
    sliceIds.EntityData.SegmentPath = "slice-ids"
    sliceIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceIds.EntityData.Children = make(map[string]types.YChild)
    sliceIds.EntityData.Children["slice-id"] = types.YChild{"SliceId", nil}
    for i := range sliceIds.SliceId {
        sliceIds.EntityData.Children[types.GetSegmentPath(&sliceIds.SliceId[i])] = types.YChild{"SliceId", &sliceIds.SliceId[i]}
    }
    sliceIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sliceIds.EntityData)
}

// LptsIfib_Nodes_Node_SliceIds_SliceId
// slice types
type LptsIfib_Nodes_Node_SliceIds_SliceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type value. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SliceName interface{}

    // Data for single pre-ifib entry. The type is slice of
    // LptsIfib_Nodes_Node_SliceIds_SliceId_Entry.
    Entry []LptsIfib_Nodes_Node_SliceIds_SliceId_Entry
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetEntityData() *types.CommonEntityData {
    sliceId.EntityData.YFilter = sliceId.YFilter
    sliceId.EntityData.YangName = "slice-id"
    sliceId.EntityData.BundleName = "cisco_ios_xr"
    sliceId.EntityData.ParentYangName = "slice-ids"
    sliceId.EntityData.SegmentPath = "slice-id" + "[slice-name='" + fmt.Sprintf("%v", sliceId.SliceName) + "']"
    sliceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceId.EntityData.Children = make(map[string]types.YChild)
    sliceId.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range sliceId.Entry {
        sliceId.EntityData.Children[types.GetSegmentPath(&sliceId.Entry[i])] = types.YChild{"Entry", &sliceId.Entry[i]}
    }
    sliceId.EntityData.Leafs = make(map[string]types.YLeaf)
    sliceId.EntityData.Leafs["slice-name"] = types.YLeaf{"SliceName", sliceId.SliceName}
    return &(sliceId.EntityData)
}

// LptsIfib_Nodes_Node_SliceIds_SliceId_Entry
// Data for single pre-ifib entry
type LptsIfib_Nodes_Node_SliceIds_SliceId_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single Pre-ifib entry. The type is interface{}
    // with range: -2147483648..2147483647.
    Entry interface{}

    // Destination Key Type. The type is string.
    DestinationType interface{}

    // Destination Port/ICMP Type/IGMP Type. The type is string.
    DestinationValue interface{}

    // Source port. The type is string.
    SourcePort interface{}

    // Destination IP Address. The type is string.
    DestinationAddr interface{}

    // Source IP Address. The type is string.
    SourceAddr interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    Vid interface{}

    // Layer 3 Protocol. The type is interface{} with range: 0..4294967295.
    L3Protocol interface{}

    // Layer 4 Protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Interface Name. The type is string.
    IntfName interface{}

    // Interface Handle. The type is interface{} with range: 0..4294967295.
    IntfHandle interface{}

    // Is SYN. The type is interface{} with range: 0..255.
    IsSyn interface{}

    // Opcode. The type is string.
    Opcode interface{}

    // Packets matched to accept. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepts interface{}

    // Packets matched to drop. The type is interface{} with range:
    // 0..18446744073709551615.
    Drops interface{}

    // Flow type. The type is string.
    FlowType interface{}

    // Listener Tag. The type is string.
    ListenerTag interface{}

    // Local Flag. The type is interface{} with range: 0..255.
    LocalFlag interface{}

    // Is FGID or not. The type is interface{} with range: 0..255.
    IsFgid interface{}

    // Deliver List Short Format. The type is string.
    DeliverListShort interface{}

    // Deliver List Long Format. The type is string.
    DeliverListLong interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    MinTtl interface{}

    // pending ifib queue delay. The type is interface{} with range:
    // 0..4294967295.
    PendingIfibqDelay interface{}

    // sl_ifibq delay. The type is interface{} with range: 0..4294967295.
    SlIfibqDelay interface{}

    // ifib program time in netio. The type is string.
    IfibProgramTime interface{}
}

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "slice-id"
    entry.EntityData.SegmentPath = "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = make(map[string]types.YChild)
    entry.EntityData.Leafs = make(map[string]types.YLeaf)
    entry.EntityData.Leafs["entry"] = types.YLeaf{"Entry", entry.Entry}
    entry.EntityData.Leafs["destination-type"] = types.YLeaf{"DestinationType", entry.DestinationType}
    entry.EntityData.Leafs["destination-value"] = types.YLeaf{"DestinationValue", entry.DestinationValue}
    entry.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", entry.SourcePort}
    entry.EntityData.Leafs["destination-addr"] = types.YLeaf{"DestinationAddr", entry.DestinationAddr}
    entry.EntityData.Leafs["source-addr"] = types.YLeaf{"SourceAddr", entry.SourceAddr}
    entry.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", entry.VrfName}
    entry.EntityData.Leafs["vid"] = types.YLeaf{"Vid", entry.Vid}
    entry.EntityData.Leafs["l3protocol"] = types.YLeaf{"L3Protocol", entry.L3Protocol}
    entry.EntityData.Leafs["l4protocol"] = types.YLeaf{"L4Protocol", entry.L4Protocol}
    entry.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", entry.IntfName}
    entry.EntityData.Leafs["intf-handle"] = types.YLeaf{"IntfHandle", entry.IntfHandle}
    entry.EntityData.Leafs["is-syn"] = types.YLeaf{"IsSyn", entry.IsSyn}
    entry.EntityData.Leafs["opcode"] = types.YLeaf{"Opcode", entry.Opcode}
    entry.EntityData.Leafs["accepts"] = types.YLeaf{"Accepts", entry.Accepts}
    entry.EntityData.Leafs["drops"] = types.YLeaf{"Drops", entry.Drops}
    entry.EntityData.Leafs["flow-type"] = types.YLeaf{"FlowType", entry.FlowType}
    entry.EntityData.Leafs["listener-tag"] = types.YLeaf{"ListenerTag", entry.ListenerTag}
    entry.EntityData.Leafs["local-flag"] = types.YLeaf{"LocalFlag", entry.LocalFlag}
    entry.EntityData.Leafs["is-fgid"] = types.YLeaf{"IsFgid", entry.IsFgid}
    entry.EntityData.Leafs["deliver-list-short"] = types.YLeaf{"DeliverListShort", entry.DeliverListShort}
    entry.EntityData.Leafs["deliver-list-long"] = types.YLeaf{"DeliverListLong", entry.DeliverListLong}
    entry.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", entry.MinTtl}
    entry.EntityData.Leafs["pending-ifibq-delay"] = types.YLeaf{"PendingIfibqDelay", entry.PendingIfibqDelay}
    entry.EntityData.Leafs["sl-ifibq-delay"] = types.YLeaf{"SlIfibqDelay", entry.SlIfibqDelay}
    entry.EntityData.Leafs["ifib-program-time"] = types.YLeaf{"IfibProgramTime", entry.IfibProgramTime}
    return &(entry.EntityData)
}

