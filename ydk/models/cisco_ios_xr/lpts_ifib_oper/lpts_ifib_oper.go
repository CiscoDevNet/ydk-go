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
    parent types.Entity
    YFilter yfilter.YFilter

    // Node ifib database.
    Nodes LptsIfib_Nodes
}

func (lptsIfib *LptsIfib) GetFilter() yfilter.YFilter { return lptsIfib.YFilter }

func (lptsIfib *LptsIfib) SetFilter(yf yfilter.YFilter) { lptsIfib.YFilter = yf }

func (lptsIfib *LptsIfib) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (lptsIfib *LptsIfib) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-ifib-oper:lpts-ifib"
}

func (lptsIfib *LptsIfib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &lptsIfib.Nodes
    }
    return nil
}

func (lptsIfib *LptsIfib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &lptsIfib.Nodes
    return children
}

func (lptsIfib *LptsIfib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsIfib *LptsIfib) GetBundleName() string { return "cisco_ios_xr" }

func (lptsIfib *LptsIfib) GetYangName() string { return "lpts-ifib" }

func (lptsIfib *LptsIfib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsIfib *LptsIfib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsIfib *LptsIfib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsIfib *LptsIfib) SetParent(parent types.Entity) { lptsIfib.parent = parent }

func (lptsIfib *LptsIfib) GetParent() types.Entity { return lptsIfib.parent }

func (lptsIfib *LptsIfib) GetParentYangName() string { return "Cisco-IOS-XR-lpts-ifib-oper" }

// LptsIfib_Nodes
// Node ifib database
type LptsIfib_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node slice . The type is slice of LptsIfib_Nodes_Node.
    Node []LptsIfib_Nodes_Node
}

func (nodes *LptsIfib_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *LptsIfib_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *LptsIfib_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *LptsIfib_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *LptsIfib_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsIfib_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *LptsIfib_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *LptsIfib_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *LptsIfib_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *LptsIfib_Nodes) GetYangName() string { return "nodes" }

func (nodes *LptsIfib_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *LptsIfib_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *LptsIfib_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *LptsIfib_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *LptsIfib_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *LptsIfib_Nodes) GetParentYangName() string { return "lpts-ifib" }

// LptsIfib_Nodes_Node
// Per node slice 
type LptsIfib_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Slice specific.
    SliceIds LptsIfib_Nodes_Node_SliceIds
}

func (node *LptsIfib_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *LptsIfib_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *LptsIfib_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "slice-ids" { return "SliceIds" }
    return ""
}

func (node *LptsIfib_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *LptsIfib_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-ids" {
        return &node.SliceIds
    }
    return nil
}

func (node *LptsIfib_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slice-ids"] = &node.SliceIds
    return children
}

func (node *LptsIfib_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *LptsIfib_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *LptsIfib_Nodes_Node) GetYangName() string { return "node" }

func (node *LptsIfib_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *LptsIfib_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *LptsIfib_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *LptsIfib_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *LptsIfib_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *LptsIfib_Nodes_Node) GetParentYangName() string { return "nodes" }

// LptsIfib_Nodes_Node_SliceIds
// Slice specific
type LptsIfib_Nodes_Node_SliceIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // slice types. The type is slice of LptsIfib_Nodes_Node_SliceIds_SliceId.
    SliceId []LptsIfib_Nodes_Node_SliceIds_SliceId
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetFilter() yfilter.YFilter { return sliceIds.YFilter }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) SetFilter(yf yfilter.YFilter) { sliceIds.YFilter = yf }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    return ""
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetSegmentPath() string {
    return "slice-ids"
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-id" {
        for _, c := range sliceIds.SliceId {
            if sliceIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsIfib_Nodes_Node_SliceIds_SliceId{}
        sliceIds.SliceId = append(sliceIds.SliceId, child)
        return &sliceIds.SliceId[len(sliceIds.SliceId)-1]
    }
    return nil
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceIds.SliceId {
        children[sliceIds.SliceId[i].GetSegmentPath()] = &sliceIds.SliceId[i]
    }
    return children
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetBundleName() string { return "cisco_ios_xr" }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetYangName() string { return "slice-ids" }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) SetParent(parent types.Entity) { sliceIds.parent = parent }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetParent() types.Entity { return sliceIds.parent }

func (sliceIds *LptsIfib_Nodes_Node_SliceIds) GetParentYangName() string { return "node" }

// LptsIfib_Nodes_Node_SliceIds_SliceId
// slice types
type LptsIfib_Nodes_Node_SliceIds_SliceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type value. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SliceName interface{}

    // Data for single pre-ifib entry. The type is slice of
    // LptsIfib_Nodes_Node_SliceIds_SliceId_Entry.
    Entry []LptsIfib_Nodes_Node_SliceIds_SliceId_Entry
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetFilter() yfilter.YFilter { return sliceId.YFilter }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) SetFilter(yf yfilter.YFilter) { sliceId.YFilter = yf }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetGoName(yname string) string {
    if yname == "slice-name" { return "SliceName" }
    if yname == "entry" { return "Entry" }
    return ""
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetSegmentPath() string {
    return "slice-id" + "[slice-name='" + fmt.Sprintf("%v", sliceId.SliceName) + "']"
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry" {
        for _, c := range sliceId.Entry {
            if sliceId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsIfib_Nodes_Node_SliceIds_SliceId_Entry{}
        sliceId.Entry = append(sliceId.Entry, child)
        return &sliceId.Entry[len(sliceId.Entry)-1]
    }
    return nil
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceId.Entry {
        children[sliceId.Entry[i].GetSegmentPath()] = &sliceId.Entry[i]
    }
    return children
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-name"] = sliceId.SliceName
    return leafs
}

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetBundleName() string { return "cisco_ios_xr" }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetYangName() string { return "slice-id" }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) SetParent(parent types.Entity) { sliceId.parent = parent }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetParent() types.Entity { return sliceId.parent }

func (sliceId *LptsIfib_Nodes_Node_SliceIds_SliceId) GetParentYangName() string { return "slice-ids" }

// LptsIfib_Nodes_Node_SliceIds_SliceId_Entry
// Data for single pre-ifib entry
type LptsIfib_Nodes_Node_SliceIds_SliceId_Entry struct {
    parent types.Entity
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

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetFilter() yfilter.YFilter { return entry.YFilter }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) SetFilter(yf yfilter.YFilter) { entry.YFilter = yf }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    if yname == "destination-type" { return "DestinationType" }
    if yname == "destination-value" { return "DestinationValue" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "destination-addr" { return "DestinationAddr" }
    if yname == "source-addr" { return "SourceAddr" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vid" { return "Vid" }
    if yname == "l3protocol" { return "L3Protocol" }
    if yname == "l4protocol" { return "L4Protocol" }
    if yname == "intf-name" { return "IntfName" }
    if yname == "intf-handle" { return "IntfHandle" }
    if yname == "is-syn" { return "IsSyn" }
    if yname == "opcode" { return "Opcode" }
    if yname == "accepts" { return "Accepts" }
    if yname == "drops" { return "Drops" }
    if yname == "flow-type" { return "FlowType" }
    if yname == "listener-tag" { return "ListenerTag" }
    if yname == "local-flag" { return "LocalFlag" }
    if yname == "is-fgid" { return "IsFgid" }
    if yname == "deliver-list-short" { return "DeliverListShort" }
    if yname == "deliver-list-long" { return "DeliverListLong" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "pending-ifibq-delay" { return "PendingIfibqDelay" }
    if yname == "sl-ifibq-delay" { return "SlIfibqDelay" }
    if yname == "ifib-program-time" { return "IfibProgramTime" }
    return ""
}

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetSegmentPath() string {
    return "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
}

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = entry.Entry
    leafs["destination-type"] = entry.DestinationType
    leafs["destination-value"] = entry.DestinationValue
    leafs["source-port"] = entry.SourcePort
    leafs["destination-addr"] = entry.DestinationAddr
    leafs["source-addr"] = entry.SourceAddr
    leafs["vrf-name"] = entry.VrfName
    leafs["vid"] = entry.Vid
    leafs["l3protocol"] = entry.L3Protocol
    leafs["l4protocol"] = entry.L4Protocol
    leafs["intf-name"] = entry.IntfName
    leafs["intf-handle"] = entry.IntfHandle
    leafs["is-syn"] = entry.IsSyn
    leafs["opcode"] = entry.Opcode
    leafs["accepts"] = entry.Accepts
    leafs["drops"] = entry.Drops
    leafs["flow-type"] = entry.FlowType
    leafs["listener-tag"] = entry.ListenerTag
    leafs["local-flag"] = entry.LocalFlag
    leafs["is-fgid"] = entry.IsFgid
    leafs["deliver-list-short"] = entry.DeliverListShort
    leafs["deliver-list-long"] = entry.DeliverListLong
    leafs["min-ttl"] = entry.MinTtl
    leafs["pending-ifibq-delay"] = entry.PendingIfibqDelay
    leafs["sl-ifibq-delay"] = entry.SlIfibqDelay
    leafs["ifib-program-time"] = entry.IfibProgramTime
    return leafs
}

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetBundleName() string { return "cisco_ios_xr" }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetYangName() string { return "entry" }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) SetParent(parent types.Entity) { entry.parent = parent }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetParent() types.Entity { return entry.parent }

func (entry *LptsIfib_Nodes_Node_SliceIds_SliceId_Entry) GetParentYangName() string { return "slice-id" }

