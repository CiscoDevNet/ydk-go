// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-pre-ifib package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lpts-pifib: lpts pre-ifib data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_pre_ifib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_pre_ifib_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-pre-ifib-oper lpts-pifib}", reflect.TypeOf(LptsPifib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib", reflect.TypeOf(LptsPifib{}))
}

// LptsPifib represents Lpts pifib
type LptsPifib string

const (
    // ISIS packets
    LptsPifib_isis LptsPifib = "isis"

    // IPv4 fragmented packets
    LptsPifib_ipv4_frag LptsPifib = "ipv4-frag"

    // IPv4 ICMP Echo packets
    LptsPifib_ipv4_echo LptsPifib = "ipv4-echo"

    // All IPv4 packets
    LptsPifib_ipv4_any LptsPifib = "ipv4-any"

    // IPv6 fragmented packets
    LptsPifib_ipv6_frag LptsPifib = "ipv6-frag"

    // IPv6 ICMP Echo packets
    LptsPifib_ipv6_echo LptsPifib = "ipv6-echo"

    // IPv6 ND packets
    LptsPifib_ipv6_nd LptsPifib = "ipv6-nd"

    // All IPv6 packets
    LptsPifib_ipv6_any LptsPifib = "ipv6-any"

    // BFD packets
    LptsPifib_bfd_any LptsPifib = "bfd-any"

    // All packets
    LptsPifib_all LptsPifib = "all"
)

// LptsPifib
// lpts pre-ifib data
type LptsPifib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of Pre-ifib Nodes.
    Nodes LptsPifib_Nodes
}

func (lptsPifib *LptsPifib) GetFilter() yfilter.YFilter { return lptsPifib.YFilter }

func (lptsPifib *LptsPifib) SetFilter(yf yfilter.YFilter) { lptsPifib.YFilter = yf }

func (lptsPifib *LptsPifib) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (lptsPifib *LptsPifib) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib"
}

func (lptsPifib *LptsPifib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &lptsPifib.Nodes
    }
    return nil
}

func (lptsPifib *LptsPifib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &lptsPifib.Nodes
    return children
}

func (lptsPifib *LptsPifib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsPifib *LptsPifib) GetBundleName() string { return "cisco_ios_xr" }

func (lptsPifib *LptsPifib) GetYangName() string { return "lpts-pifib" }

func (lptsPifib *LptsPifib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsPifib *LptsPifib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsPifib *LptsPifib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsPifib *LptsPifib) SetParent(parent types.Entity) { lptsPifib.parent = parent }

func (lptsPifib *LptsPifib) GetParent() types.Entity { return lptsPifib.parent }

func (lptsPifib *LptsPifib) GetParentYangName() string { return "Cisco-IOS-XR-lpts-pre-ifib-oper" }

// LptsPifib_Nodes
// List of Pre-ifib Nodes
type LptsPifib_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pre-ifib data for particular node. The type is slice of
    // LptsPifib_Nodes_Node.
    Node []LptsPifib_Nodes_Node
}

func (nodes *LptsPifib_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *LptsPifib_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *LptsPifib_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *LptsPifib_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *LptsPifib_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *LptsPifib_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *LptsPifib_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *LptsPifib_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *LptsPifib_Nodes) GetYangName() string { return "nodes" }

func (nodes *LptsPifib_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *LptsPifib_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *LptsPifib_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *LptsPifib_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *LptsPifib_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *LptsPifib_Nodes) GetParentYangName() string { return "lpts-pifib" }

// LptsPifib_Nodes_Node
// Pre-ifib data for particular node
type LptsPifib_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Type specific.
    TypeValues LptsPifib_Nodes_Node_TypeValues

    // Dynamic Flows Statistics.
    DynamicFlowsStats LptsPifib_Nodes_Node_DynamicFlowsStats

    // Hardware specific.
    Hardware LptsPifib_Nodes_Node_Hardware
}

func (node *LptsPifib_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *LptsPifib_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *LptsPifib_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "type-values" { return "TypeValues" }
    if yname == "dynamic-flows-stats" { return "DynamicFlowsStats" }
    if yname == "Cisco-IOS-XR-platform-pifib-oper:hardware" { return "Hardware" }
    return ""
}

func (node *LptsPifib_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *LptsPifib_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "type-values" {
        return &node.TypeValues
    }
    if childYangName == "dynamic-flows-stats" {
        return &node.DynamicFlowsStats
    }
    if childYangName == "Cisco-IOS-XR-platform-pifib-oper:hardware" {
        return &node.Hardware
    }
    return nil
}

func (node *LptsPifib_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["type-values"] = &node.TypeValues
    children["dynamic-flows-stats"] = &node.DynamicFlowsStats
    children["Cisco-IOS-XR-platform-pifib-oper:hardware"] = &node.Hardware
    return children
}

func (node *LptsPifib_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *LptsPifib_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *LptsPifib_Nodes_Node) GetYangName() string { return "node" }

func (node *LptsPifib_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *LptsPifib_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *LptsPifib_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *LptsPifib_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *LptsPifib_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *LptsPifib_Nodes_Node) GetParentYangName() string { return "nodes" }

// LptsPifib_Nodes_Node_TypeValues
// Type specific
type LptsPifib_Nodes_Node_TypeValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // pifib types. The type is slice of
    // LptsPifib_Nodes_Node_TypeValues_TypeValue.
    TypeValue []LptsPifib_Nodes_Node_TypeValues_TypeValue
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetFilter() yfilter.YFilter { return typeValues.YFilter }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) SetFilter(yf yfilter.YFilter) { typeValues.YFilter = yf }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetGoName(yname string) string {
    if yname == "type-value" { return "TypeValue" }
    return ""
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetSegmentPath() string {
    return "type-values"
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "type-value" {
        for _, c := range typeValues.TypeValue {
            if typeValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_TypeValues_TypeValue{}
        typeValues.TypeValue = append(typeValues.TypeValue, child)
        return &typeValues.TypeValue[len(typeValues.TypeValue)-1]
    }
    return nil
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range typeValues.TypeValue {
        children[typeValues.TypeValue[i].GetSegmentPath()] = &typeValues.TypeValue[i]
    }
    return children
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetBundleName() string { return "cisco_ios_xr" }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetYangName() string { return "type-values" }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) SetParent(parent types.Entity) { typeValues.parent = parent }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetParent() types.Entity { return typeValues.parent }

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetParentYangName() string { return "node" }

// LptsPifib_Nodes_Node_TypeValues_TypeValue
// pifib types
type LptsPifib_Nodes_Node_TypeValues_TypeValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type value. The type is LptsPifib.
    PifibType interface{}

    // Data for single pre-ifib entry. The type is slice of
    // LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry.
    Entry []LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetFilter() yfilter.YFilter { return typeValue.YFilter }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) SetFilter(yf yfilter.YFilter) { typeValue.YFilter = yf }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetGoName(yname string) string {
    if yname == "pifib-type" { return "PifibType" }
    if yname == "entry" { return "Entry" }
    return ""
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetSegmentPath() string {
    return "type-value" + "[pifib-type='" + fmt.Sprintf("%v", typeValue.PifibType) + "']"
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry" {
        for _, c := range typeValue.Entry {
            if typeValue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry{}
        typeValue.Entry = append(typeValue.Entry, child)
        return &typeValue.Entry[len(typeValue.Entry)-1]
    }
    return nil
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range typeValue.Entry {
        children[typeValue.Entry[i].GetSegmentPath()] = &typeValue.Entry[i]
    }
    return children
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pifib-type"] = typeValue.PifibType
    return leafs
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetBundleName() string { return "cisco_ios_xr" }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetYangName() string { return "type-value" }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) SetParent(parent types.Entity) { typeValue.parent = parent }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetParent() types.Entity { return typeValue.parent }

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetParentYangName() string { return "type-values" }

// LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry
// Data for single pre-ifib entry
type LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single Pre-ifib entry. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Entry interface{}

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

    // Destination IP Address. The type is string.
    DestinationAddr interface{}

    // Source IP Address. The type is string.
    SourceAddr interface{}

    // Destination Key Type. The type is string.
    DestinationType interface{}

    // Destination Port/ICMP Type/IGMP Type. The type is string.
    DestinationValue interface{}

    // Source port. The type is string.
    SourcePort interface{}

    // Is Fragment. The type is interface{} with range: 0..255.
    IsFrag interface{}

    // Is SYN. The type is interface{} with range: 0..255.
    IsSyn interface{}

    // Opcode. The type is string.
    Opcode interface{}

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

    // Packets matched to accept. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepts interface{}

    // Packets matched for drop. The type is interface{} with range:
    // 0..18446744073709551615.
    Drops interface{}

    // Is Stale. The type is interface{} with range: 0..255.
    Stale interface{}

    // sub Pre-IFIB type. The type is interface{} with range: 0..255.
    PifibType interface{}

    // Creation or Update Time. The type is string.
    PifibProgramTime interface{}
}

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetFilter() yfilter.YFilter { return entry.YFilter }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) SetFilter(yf yfilter.YFilter) { entry.YFilter = yf }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vid" { return "Vid" }
    if yname == "l3protocol" { return "L3Protocol" }
    if yname == "l4protocol" { return "L4Protocol" }
    if yname == "intf-name" { return "IntfName" }
    if yname == "intf-handle" { return "IntfHandle" }
    if yname == "destination-addr" { return "DestinationAddr" }
    if yname == "source-addr" { return "SourceAddr" }
    if yname == "destination-type" { return "DestinationType" }
    if yname == "destination-value" { return "DestinationValue" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "is-frag" { return "IsFrag" }
    if yname == "is-syn" { return "IsSyn" }
    if yname == "opcode" { return "Opcode" }
    if yname == "flow-type" { return "FlowType" }
    if yname == "listener-tag" { return "ListenerTag" }
    if yname == "local-flag" { return "LocalFlag" }
    if yname == "is-fgid" { return "IsFgid" }
    if yname == "deliver-list-short" { return "DeliverListShort" }
    if yname == "deliver-list-long" { return "DeliverListLong" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "accepts" { return "Accepts" }
    if yname == "drops" { return "Drops" }
    if yname == "stale" { return "Stale" }
    if yname == "pifib-type" { return "PifibType" }
    if yname == "pifib-program-time" { return "PifibProgramTime" }
    return ""
}

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetSegmentPath() string {
    return "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
}

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = entry.Entry
    leafs["vrf-name"] = entry.VrfName
    leafs["vid"] = entry.Vid
    leafs["l3protocol"] = entry.L3Protocol
    leafs["l4protocol"] = entry.L4Protocol
    leafs["intf-name"] = entry.IntfName
    leafs["intf-handle"] = entry.IntfHandle
    leafs["destination-addr"] = entry.DestinationAddr
    leafs["source-addr"] = entry.SourceAddr
    leafs["destination-type"] = entry.DestinationType
    leafs["destination-value"] = entry.DestinationValue
    leafs["source-port"] = entry.SourcePort
    leafs["is-frag"] = entry.IsFrag
    leafs["is-syn"] = entry.IsSyn
    leafs["opcode"] = entry.Opcode
    leafs["flow-type"] = entry.FlowType
    leafs["listener-tag"] = entry.ListenerTag
    leafs["local-flag"] = entry.LocalFlag
    leafs["is-fgid"] = entry.IsFgid
    leafs["deliver-list-short"] = entry.DeliverListShort
    leafs["deliver-list-long"] = entry.DeliverListLong
    leafs["min-ttl"] = entry.MinTtl
    leafs["accepts"] = entry.Accepts
    leafs["drops"] = entry.Drops
    leafs["stale"] = entry.Stale
    leafs["pifib-type"] = entry.PifibType
    leafs["pifib-program-time"] = entry.PifibProgramTime
    return leafs
}

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetBundleName() string { return "cisco_ios_xr" }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetYangName() string { return "entry" }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) SetParent(parent types.Entity) { entry.parent = parent }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetParent() types.Entity { return entry.parent }

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetParentYangName() string { return "type-value" }

// LptsPifib_Nodes_Node_DynamicFlowsStats
// Dynamic Flows Statistics
type LptsPifib_Nodes_Node_DynamicFlowsStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dynamic Flows Enabled. The type is bool.
    DynamicFlowsEnabled interface{}

    // Platform Max. The type is interface{} with range: 0..4294967295.
    PlatformSupportedMax interface{}

    // Platform Config Limit. The type is interface{} with range: 0..4294967295.
    PlatformConfiguredMax interface{}

    // Platform Total Configured. The type is interface{} with range:
    // 0..4294967295.
    PlatformTotalConfigured interface{}

    // Total HW Entries. The type is interface{} with range: 0..4294967295.
    TotalHwEntries interface{}

    // Total SW Entries. The type is interface{} with range: 0..4294967295.
    TotalSwEntries interface{}

    // Flow Datalist. The type is slice of
    // LptsPifib_Nodes_Node_DynamicFlowsStats_Flow.
    Flow []LptsPifib_Nodes_Node_DynamicFlowsStats_Flow
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetFilter() yfilter.YFilter { return dynamicFlowsStats.YFilter }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) SetFilter(yf yfilter.YFilter) { dynamicFlowsStats.YFilter = yf }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetGoName(yname string) string {
    if yname == "dynamic-flows-enabled" { return "DynamicFlowsEnabled" }
    if yname == "platform-supported-max" { return "PlatformSupportedMax" }
    if yname == "platform-configured-max" { return "PlatformConfiguredMax" }
    if yname == "platform-total-configured" { return "PlatformTotalConfigured" }
    if yname == "total-hw-entries" { return "TotalHwEntries" }
    if yname == "total-sw-entries" { return "TotalSwEntries" }
    if yname == "flow" { return "Flow" }
    return ""
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetSegmentPath() string {
    return "dynamic-flows-stats"
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow" {
        for _, c := range dynamicFlowsStats.Flow {
            if dynamicFlowsStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_DynamicFlowsStats_Flow{}
        dynamicFlowsStats.Flow = append(dynamicFlowsStats.Flow, child)
        return &dynamicFlowsStats.Flow[len(dynamicFlowsStats.Flow)-1]
    }
    return nil
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamicFlowsStats.Flow {
        children[dynamicFlowsStats.Flow[i].GetSegmentPath()] = &dynamicFlowsStats.Flow[i]
    }
    return children
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dynamic-flows-enabled"] = dynamicFlowsStats.DynamicFlowsEnabled
    leafs["platform-supported-max"] = dynamicFlowsStats.PlatformSupportedMax
    leafs["platform-configured-max"] = dynamicFlowsStats.PlatformConfiguredMax
    leafs["platform-total-configured"] = dynamicFlowsStats.PlatformTotalConfigured
    leafs["total-hw-entries"] = dynamicFlowsStats.TotalHwEntries
    leafs["total-sw-entries"] = dynamicFlowsStats.TotalSwEntries
    return leafs
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetYangName() string { return "dynamic-flows-stats" }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) SetParent(parent types.Entity) { dynamicFlowsStats.parent = parent }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetParent() types.Entity { return dynamicFlowsStats.parent }

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetParentYangName() string { return "node" }

// LptsPifib_Nodes_Node_DynamicFlowsStats_Flow
// Flow Datalist
type LptsPifib_Nodes_Node_DynamicFlowsStats_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow Name. The type is string.
    FlowName interface{}

    // Is Configurable. The type is bool.
    Configurable interface{}

    // Is Configured. The type is bool.
    Configured interface{}

    // Default Max. The type is interface{} with range: 0..4294967295.
    DefaultMax interface{}

    // Configured Max. The type is string.
    ConfiguredMax interface{}

    // Active Max. The type is interface{} with range: 0..4294967295.
    ActiveMax interface{}

    // Hardware Count. The type is interface{} with range: 0..4294967295.
    HardwareCount interface{}

    // Software Count. The type is interface{} with range: 0..4294967295.
    SoftwareCount interface{}

    // Pending Software Entries. The type is bool.
    PendingSoftwareEntries interface{}
}

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetGoName(yname string) string {
    if yname == "flow-name" { return "FlowName" }
    if yname == "configurable" { return "Configurable" }
    if yname == "configured" { return "Configured" }
    if yname == "default-max" { return "DefaultMax" }
    if yname == "configured-max" { return "ConfiguredMax" }
    if yname == "active-max" { return "ActiveMax" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "software-count" { return "SoftwareCount" }
    if yname == "pending-software-entries" { return "PendingSoftwareEntries" }
    return ""
}

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetSegmentPath() string {
    return "flow"
}

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-name"] = flow.FlowName
    leafs["configurable"] = flow.Configurable
    leafs["configured"] = flow.Configured
    leafs["default-max"] = flow.DefaultMax
    leafs["configured-max"] = flow.ConfiguredMax
    leafs["active-max"] = flow.ActiveMax
    leafs["hardware-count"] = flow.HardwareCount
    leafs["software-count"] = flow.SoftwareCount
    leafs["pending-software-entries"] = flow.PendingSoftwareEntries
    return leafs
}

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetBundleName() string { return "cisco_ios_xr" }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetYangName() string { return "flow" }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetParent() types.Entity { return flow.parent }

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetParentYangName() string { return "dynamic-flows-stats" }

// LptsPifib_Nodes_Node_Hardware
// Hardware specific
type LptsPifib_Nodes_Node_Hardware struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Usage Table options.
    UsageEntries LptsPifib_Nodes_Node_Hardware_UsageEntries

    // Police details.
    Police LptsPifib_Nodes_Node_Hardware_Police

    // Static Police details.
    StaticPolice LptsPifib_Nodes_Node_Hardware_StaticPolice

    // Bfd details.
    Bfd LptsPifib_Nodes_Node_Hardware_Bfd

    // Hardware Entry type.
    Statistics LptsPifib_Nodes_Node_Hardware_Statistics

    // Hardware Entry options.
    IndexEntries LptsPifib_Nodes_Node_Hardware_IndexEntries
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetFilter() yfilter.YFilter { return hardware.YFilter }

func (hardware *LptsPifib_Nodes_Node_Hardware) SetFilter(yf yfilter.YFilter) { hardware.YFilter = yf }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetGoName(yname string) string {
    if yname == "usage-entries" { return "UsageEntries" }
    if yname == "police" { return "Police" }
    if yname == "static-police" { return "StaticPolice" }
    if yname == "bfd" { return "Bfd" }
    if yname == "statistics" { return "Statistics" }
    if yname == "index-entries" { return "IndexEntries" }
    return ""
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetSegmentPath() string {
    return "Cisco-IOS-XR-platform-pifib-oper:hardware"
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usage-entries" {
        return &hardware.UsageEntries
    }
    if childYangName == "police" {
        return &hardware.Police
    }
    if childYangName == "static-police" {
        return &hardware.StaticPolice
    }
    if childYangName == "bfd" {
        return &hardware.Bfd
    }
    if childYangName == "statistics" {
        return &hardware.Statistics
    }
    if childYangName == "index-entries" {
        return &hardware.IndexEntries
    }
    return nil
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["usage-entries"] = &hardware.UsageEntries
    children["police"] = &hardware.Police
    children["static-police"] = &hardware.StaticPolice
    children["bfd"] = &hardware.Bfd
    children["statistics"] = &hardware.Statistics
    children["index-entries"] = &hardware.IndexEntries
    return children
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetBundleName() string { return "cisco_ios_xr" }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetYangName() string { return "hardware" }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardware *LptsPifib_Nodes_Node_Hardware) SetParent(parent types.Entity) { hardware.parent = parent }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetParent() types.Entity { return hardware.parent }

func (hardware *LptsPifib_Nodes_Node_Hardware) GetParentYangName() string { return "node" }

// LptsPifib_Nodes_Node_Hardware_UsageEntries
// Usage Table options
type LptsPifib_Nodes_Node_Hardware_UsageEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Usage details. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry.
    UsageEntry []LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetFilter() yfilter.YFilter { return usageEntries.YFilter }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) SetFilter(yf yfilter.YFilter) { usageEntries.YFilter = yf }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetGoName(yname string) string {
    if yname == "usage-entry" { return "UsageEntry" }
    return ""
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetSegmentPath() string {
    return "usage-entries"
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usage-entry" {
        for _, c := range usageEntries.UsageEntry {
            if usageEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry{}
        usageEntries.UsageEntry = append(usageEntries.UsageEntry, child)
        return &usageEntries.UsageEntry[len(usageEntries.UsageEntry)-1]
    }
    return nil
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usageEntries.UsageEntry {
        children[usageEntries.UsageEntry[i].GetSegmentPath()] = &usageEntries.UsageEntry[i]
    }
    return children
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetBundleName() string { return "cisco_ios_xr" }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetYangName() string { return "usage-entries" }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) SetParent(parent types.Entity) { usageEntries.parent = parent }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetParent() types.Entity { return usageEntries.parent }

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetParentYangName() string { return "hardware" }

// LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry
// Usage details
type LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Region ID. The type is UsageAddressFamily.
    RegionId interface{}

    // Per TCAM type usage info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo.
    UsageInfo []LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetFilter() yfilter.YFilter { return usageEntry.YFilter }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) SetFilter(yf yfilter.YFilter) { usageEntry.YFilter = yf }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetGoName(yname string) string {
    if yname == "region-id" { return "RegionId" }
    if yname == "usage-info" { return "UsageInfo" }
    return ""
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetSegmentPath() string {
    return "usage-entry" + "[region-id='" + fmt.Sprintf("%v", usageEntry.RegionId) + "']"
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usage-info" {
        for _, c := range usageEntry.UsageInfo {
            if usageEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo{}
        usageEntry.UsageInfo = append(usageEntry.UsageInfo, child)
        return &usageEntry.UsageInfo[len(usageEntry.UsageInfo)-1]
    }
    return nil
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usageEntry.UsageInfo {
        children[usageEntry.UsageInfo[i].GetSegmentPath()] = &usageEntry.UsageInfo[i]
    }
    return children
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["region-id"] = usageEntry.RegionId
    return leafs
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetBundleName() string { return "cisco_ios_xr" }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetYangName() string { return "usage-entry" }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) SetParent(parent types.Entity) { usageEntry.parent = parent }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetParent() types.Entity { return usageEntry.parent }

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetParentYangName() string { return "usage-entries" }

// LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo
// Per TCAM type usage info
type LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pipe ID. The type is interface{} with range: 0..255.
    PipeId interface{}

    // Region Type. The type is interface{} with range: 0..255.
    Region interface{}

    // Region ID. The type is interface{} with range: 0..255.
    RegionId interface{}

    // Maximum Number of Entries in the Region. The type is interface{} with
    // range: 0..4294967295.
    Size interface{}

    // Used Number of Entries in the Region. The type is interface{} with range:
    // 0..4294967295.
    Used interface{}
}

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetFilter() yfilter.YFilter { return usageInfo.YFilter }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) SetFilter(yf yfilter.YFilter) { usageInfo.YFilter = yf }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetGoName(yname string) string {
    if yname == "pipe-id" { return "PipeId" }
    if yname == "region" { return "Region" }
    if yname == "region-id" { return "RegionId" }
    if yname == "size" { return "Size" }
    if yname == "used" { return "Used" }
    return ""
}

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetSegmentPath() string {
    return "usage-info"
}

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pipe-id"] = usageInfo.PipeId
    leafs["region"] = usageInfo.Region
    leafs["region-id"] = usageInfo.RegionId
    leafs["size"] = usageInfo.Size
    leafs["used"] = usageInfo.Used
    return leafs
}

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetBundleName() string { return "cisco_ios_xr" }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetYangName() string { return "usage-info" }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) SetParent(parent types.Entity) { usageInfo.parent = parent }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetParent() types.Entity { return usageInfo.parent }

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetParentYangName() string { return "usage-entry" }

// LptsPifib_Nodes_Node_Hardware_Police
// Police details
type LptsPifib_Nodes_Node_Hardware_Police struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo.
    PoliceInfo []LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *LptsPifib_Nodes_Node_Hardware_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetGoName(yname string) string {
    if yname == "police-info" { return "PoliceInfo" }
    return ""
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetSegmentPath() string {
    return "police"
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police-info" {
        for _, c := range police.PoliceInfo {
            if police.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo{}
        police.PoliceInfo = append(police.PoliceInfo, child)
        return &police.PoliceInfo[len(police.PoliceInfo)-1]
    }
    return nil
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range police.PoliceInfo {
        children[police.PoliceInfo[i].GetSegmentPath()] = &police.PoliceInfo[i]
    }
    return children
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetYangName() string { return "police" }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *LptsPifib_Nodes_Node_Hardware_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetParent() types.Entity { return police.parent }

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetParentYangName() string { return "hardware" }

// LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo
// Per flow type police info
type LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // avgrate. The type is interface{} with range: 0..4294967295.
    Avgrate interface{}

    // burst. The type is interface{} with range: 0..4294967295.
    Burst interface{}

    // static avgrate. The type is interface{} with range: 0..4294967295.
    StaticAvgrate interface{}

    // avgrate type. The type is interface{} with range: 0..4294967295.
    AvgrateType interface{}

    // accepted stats. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedStats interface{}

    // dropped stats. The type is interface{} with range: 0..18446744073709551615.
    DroppedStats interface{}

    // policer. The type is interface{} with range: 0..4294967295.
    Policer interface{}

    // iptos value. The type is interface{} with range: 0..255.
    IptosValue interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}

    // acl config. The type is interface{} with range: 0..255.
    AclConfig interface{}

    // acl str. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AclStr interface{}
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetFilter() yfilter.YFilter { return policeInfo.YFilter }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) SetFilter(yf yfilter.YFilter) { policeInfo.YFilter = yf }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetGoName(yname string) string {
    if yname == "avgrate" { return "Avgrate" }
    if yname == "burst" { return "Burst" }
    if yname == "static-avgrate" { return "StaticAvgrate" }
    if yname == "avgrate-type" { return "AvgrateType" }
    if yname == "accepted-stats" { return "AcceptedStats" }
    if yname == "dropped-stats" { return "DroppedStats" }
    if yname == "policer" { return "Policer" }
    if yname == "iptos-value" { return "IptosValue" }
    if yname == "change-type" { return "ChangeType" }
    if yname == "acl-config" { return "AclConfig" }
    if yname == "acl-str" { return "AclStr" }
    return ""
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetSegmentPath() string {
    return "police-info"
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["avgrate"] = policeInfo.Avgrate
    leafs["burst"] = policeInfo.Burst
    leafs["static-avgrate"] = policeInfo.StaticAvgrate
    leafs["avgrate-type"] = policeInfo.AvgrateType
    leafs["accepted-stats"] = policeInfo.AcceptedStats
    leafs["dropped-stats"] = policeInfo.DroppedStats
    leafs["policer"] = policeInfo.Policer
    leafs["iptos-value"] = policeInfo.IptosValue
    leafs["change-type"] = policeInfo.ChangeType
    leafs["acl-config"] = policeInfo.AclConfig
    leafs["acl-str"] = policeInfo.AclStr
    return leafs
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetYangName() string { return "police-info" }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) SetParent(parent types.Entity) { policeInfo.parent = parent }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetParent() types.Entity { return policeInfo.parent }

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetParentYangName() string { return "police" }

// LptsPifib_Nodes_Node_Hardware_StaticPolice
// Static Police details
type LptsPifib_Nodes_Node_Hardware_StaticPolice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per punt reason info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo.
    StaticInfo []LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetFilter() yfilter.YFilter { return staticPolice.YFilter }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) SetFilter(yf yfilter.YFilter) { staticPolice.YFilter = yf }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetGoName(yname string) string {
    if yname == "static-info" { return "StaticInfo" }
    return ""
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetSegmentPath() string {
    return "static-police"
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-info" {
        for _, c := range staticPolice.StaticInfo {
            if staticPolice.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo{}
        staticPolice.StaticInfo = append(staticPolice.StaticInfo, child)
        return &staticPolice.StaticInfo[len(staticPolice.StaticInfo)-1]
    }
    return nil
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticPolice.StaticInfo {
        children[staticPolice.StaticInfo[i].GetSegmentPath()] = &staticPolice.StaticInfo[i]
    }
    return children
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetBundleName() string { return "cisco_ios_xr" }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetYangName() string { return "static-police" }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) SetParent(parent types.Entity) { staticPolice.parent = parent }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetParent() types.Entity { return staticPolice.parent }

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetParentYangName() string { return "hardware" }

// LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo
// Per punt reason info
type LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo struct {
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

    // punt reason string. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PuntReasonString interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}
}

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetFilter() yfilter.YFilter { return staticInfo.YFilter }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) SetFilter(yf yfilter.YFilter) { staticInfo.YFilter = yf }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetGoName(yname string) string {
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

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetSegmentPath() string {
    return "static-info"
}

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetLeafs() map[string]interface{} {
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

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetBundleName() string { return "cisco_ios_xr" }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetYangName() string { return "static-info" }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) SetParent(parent types.Entity) { staticInfo.parent = parent }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetParent() types.Entity { return staticInfo.parent }

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetParentYangName() string { return "static-police" }

// LptsPifib_Nodes_Node_Hardware_Bfd
// Bfd details
type LptsPifib_Nodes_Node_Hardware_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per bfd disc entry info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo.
    BfdEntryInfo []LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetGoName(yname string) string {
    if yname == "bfd-entry-info" { return "BfdEntryInfo" }
    return ""
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-entry-info" {
        for _, c := range bfd.BfdEntryInfo {
            if bfd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo{}
        bfd.BfdEntryInfo = append(bfd.BfdEntryInfo, child)
        return &bfd.BfdEntryInfo[len(bfd.BfdEntryInfo)-1]
    }
    return nil
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfd.BfdEntryInfo {
        children[bfd.BfdEntryInfo[i].GetSegmentPath()] = &bfd.BfdEntryInfo[i]
    }
    return children
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetYangName() string { return "bfd" }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetParentYangName() string { return "hardware" }

// LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo
// Per bfd disc entry info
type LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // index. The type is interface{} with range: 0..255.
    Index interface{}

    // is mcast. The type is interface{} with range: 0..255.
    IsMcast interface{}

    // fgid or vqi. The type is interface{} with range: 0..4294967295.
    FgidOrVqi interface{}

    // is valid. The type is interface{} with range: 0..255.
    IsValid interface{}

    // policer id. The type is interface{} with range: 0..4294967295.
    PolicerId interface{}
}

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetFilter() yfilter.YFilter { return bfdEntryInfo.YFilter }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) SetFilter(yf yfilter.YFilter) { bfdEntryInfo.YFilter = yf }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "is-mcast" { return "IsMcast" }
    if yname == "fgid-or-vqi" { return "FgidOrVqi" }
    if yname == "is-valid" { return "IsValid" }
    if yname == "policer-id" { return "PolicerId" }
    return ""
}

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetSegmentPath() string {
    return "bfd-entry-info"
}

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = bfdEntryInfo.Index
    leafs["is-mcast"] = bfdEntryInfo.IsMcast
    leafs["fgid-or-vqi"] = bfdEntryInfo.FgidOrVqi
    leafs["is-valid"] = bfdEntryInfo.IsValid
    leafs["policer-id"] = bfdEntryInfo.PolicerId
    return leafs
}

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetBundleName() string { return "cisco_ios_xr" }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetYangName() string { return "bfd-entry-info" }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) SetParent(parent types.Entity) { bfdEntryInfo.parent = parent }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetParent() types.Entity { return bfdEntryInfo.parent }

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetParentYangName() string { return "bfd" }

// LptsPifib_Nodes_Node_Hardware_Statistics
// Hardware Entry type
type LptsPifib_Nodes_Node_Hardware_Statistics struct {
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

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetGoName(yname string) string {
    if yname == "accepted" { return "Accepted" }
    if yname == "dropped" { return "Dropped" }
    if yname == "clear-ts" { return "ClearTs" }
    if yname == "no-stats-mem-err" { return "NoStatsMemErr" }
    return ""
}

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accepted"] = statistics.Accepted
    leafs["dropped"] = statistics.Dropped
    leafs["clear-ts"] = statistics.ClearTs
    leafs["no-stats-mem-err"] = statistics.NoStatsMemErr
    return leafs
}

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetYangName() string { return "statistics" }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetParentYangName() string { return "hardware" }

// LptsPifib_Nodes_Node_Hardware_IndexEntries
// Hardware Entry options
type LptsPifib_Nodes_Node_Hardware_IndexEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entry options. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry.
    IndexEntry []LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetFilter() yfilter.YFilter { return indexEntries.YFilter }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) SetFilter(yf yfilter.YFilter) { indexEntries.YFilter = yf }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetGoName(yname string) string {
    if yname == "index-entry" { return "IndexEntry" }
    return ""
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetSegmentPath() string {
    return "index-entries"
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "index-entry" {
        for _, c := range indexEntries.IndexEntry {
            if indexEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry{}
        indexEntries.IndexEntry = append(indexEntries.IndexEntry, child)
        return &indexEntries.IndexEntry[len(indexEntries.IndexEntry)-1]
    }
    return nil
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indexEntries.IndexEntry {
        children[indexEntries.IndexEntry[i].GetSegmentPath()] = &indexEntries.IndexEntry[i]
    }
    return children
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetBundleName() string { return "cisco_ios_xr" }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetYangName() string { return "index-entries" }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) SetParent(parent types.Entity) { indexEntries.parent = parent }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetParent() types.Entity { return indexEntries.parent }

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetParentYangName() string { return "hardware" }

// LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry
// Entry options
type LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Layer 3 Protocol. The type is interface{} with range: 0..4294967295.
    L3Protocol interface{}

    // Layer 4 Protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Interface Handle. The type is interface{} with range: 0..4294967295.
    IntfHandle interface{}

    // Interface Name. The type is string.
    IntfName interface{}

    // Interface uidb index. The type is interface{} with range: 0..4294967295.
    UidbIndex interface{}

    // Local IP Address. The type is string.
    LocalAddr interface{}

    // Local Prefix Length. The type is interface{} with range: 0..4294967295.
    LocalPrefixLen interface{}

    // Remote IP Address. The type is string.
    RemoteAddr interface{}

    // Remote Prefix Length. The type is interface{} with range: 0..4294967295.
    RemotePrefixLen interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Remote Port/ICMP Type/IGMP Type. The type is interface{} with range:
    // 0..4294967295.
    UValue interface{}

    // Union Key Length. The type is interface{} with range: 0..4294967295.
    ULen interface{}

    // Local port. The type is interface{} with range: 0..4294967295.
    LocalPort interface{}

    // Is Fragment. The type is interface{} with range: 0..255.
    IsFrag interface{}

    // Is SYN. The type is interface{} with range: 0..255.
    IsSyn interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // Action String. The type is string.
    ActionString interface{}

    // Listener Tag. The type is interface{} with range: 0..255.
    ListenerTag interface{}

    // Is FGID or not. The type is interface{} with range: 0..255.
    IsFgid interface{}

    // Is VRF or not. The type is interface{} with range: 0..255.
    IsVrf interface{}

    // Is optimized or not. The type is interface{} with range: 0..255.
    IsOptimized interface{}

    // Is uidb set for optimized entry or not. The type is interface{} with range:
    // 0..255.
    IsUidbOptBit interface{}

    // fabric group id or swith fabric port. The type is interface{} with range:
    // 0..4294967295.
    FgidOrSfp interface{}

    // Is entry remote or not. The type is interface{} with range: 0..255.
    RemoteRack interface{}

    // Remote racknum if remote. The type is interface{} with range:
    // 0..4294967295.
    RackId interface{}

    // Remote slotnum if remote. The type is interface{} with range:
    // 0..4294967295.
    Rslot interface{}

    // Committed Information Rate. The type is interface{} with range:
    // 0..18446744073709551615.
    Cir interface{}

    // Flow type. The type is interface{} with range: 0..4294967295.
    FlowType interface{}

    // Flow priority or COS. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Stream ID. The type is interface{} with range: 0..4294967295.
    Sid interface{}

    // Policer avg. rate limit. The type is interface{} with range: 0..4294967295.
    PolicerAvgrate interface{}

    // Policer burst. The type is interface{} with range: 0..4294967295.
    PolicerBurst interface{}

    // Lookup priority. The type is interface{} with range:
    // -2147483648..2147483647.
    LookupPriority interface{}

    // Storage priority. The type is interface{} with range:
    // -2147483648..2147483647.
    StoragePriority interface{}

    // Number of TCAM entries used. The type is interface{} with range:
    // -2147483648..2147483647.
    NumTmEntries interface{}

    // ptr to ifib_entry_st. The type is interface{} with range: 0..4294967295.
    EntryPtr interface{}

    // ptr to ifib_entry_shadow_st. The type is interface{} with range:
    // 0..4294967295.
    EntryShadowPtr interface{}

    // ptr to dlqueue list node. The type is interface{} with range:
    // 0..4294967295.
    ListNodePtr interface{}

    // state of pifib entry. The type is interface{} with range: 0..255.
    State interface{}

    // failure cause. The type is interface{} with range: 0..255.
    RetryFailCause interface{}

    // retries count. The type is interface{} with range: 0..255.
    NumRetries interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    MinTtl interface{}

    // Union Key Type. The type is interface{} with range: 0..255.
    UType interface{}

    // Remote FGID. The type is interface{} with range: 0..4294967295.
    RemoteFgid interface{}

    // Acl name. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AclStr interface{}

    // Stats not available. The type is interface{} with range: 0..255.
    NoStats interface{}

    // Per pipe type hardware info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo.
    HwInfo []LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetFilter() yfilter.YFilter { return indexEntry.YFilter }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) SetFilter(yf yfilter.YFilter) { indexEntry.YFilter = yf }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "l3protocol" { return "L3Protocol" }
    if yname == "l4protocol" { return "L4Protocol" }
    if yname == "intf-handle" { return "IntfHandle" }
    if yname == "intf-name" { return "IntfName" }
    if yname == "uidb-index" { return "UidbIndex" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "local-prefix-len" { return "LocalPrefixLen" }
    if yname == "remote-addr" { return "RemoteAddr" }
    if yname == "remote-prefix-len" { return "RemotePrefixLen" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "u-value" { return "UValue" }
    if yname == "u-len" { return "ULen" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "is-frag" { return "IsFrag" }
    if yname == "is-syn" { return "IsSyn" }
    if yname == "action" { return "Action" }
    if yname == "action-string" { return "ActionString" }
    if yname == "listener-tag" { return "ListenerTag" }
    if yname == "is-fgid" { return "IsFgid" }
    if yname == "is-vrf" { return "IsVrf" }
    if yname == "is-optimized" { return "IsOptimized" }
    if yname == "is-uidb-opt-bit" { return "IsUidbOptBit" }
    if yname == "fgid-or-sfp" { return "FgidOrSfp" }
    if yname == "remote-rack" { return "RemoteRack" }
    if yname == "rack-id" { return "RackId" }
    if yname == "rslot" { return "Rslot" }
    if yname == "cir" { return "Cir" }
    if yname == "flow-type" { return "FlowType" }
    if yname == "priority" { return "Priority" }
    if yname == "sid" { return "Sid" }
    if yname == "policer-avgrate" { return "PolicerAvgrate" }
    if yname == "policer-burst" { return "PolicerBurst" }
    if yname == "lookup-priority" { return "LookupPriority" }
    if yname == "storage-priority" { return "StoragePriority" }
    if yname == "num-tm-entries" { return "NumTmEntries" }
    if yname == "entry-ptr" { return "EntryPtr" }
    if yname == "entry-shadow-ptr" { return "EntryShadowPtr" }
    if yname == "list-node-ptr" { return "ListNodePtr" }
    if yname == "state" { return "State" }
    if yname == "retry-fail-cause" { return "RetryFailCause" }
    if yname == "num-retries" { return "NumRetries" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "u-type" { return "UType" }
    if yname == "remote-fgid" { return "RemoteFgid" }
    if yname == "acl-str" { return "AclStr" }
    if yname == "no-stats" { return "NoStats" }
    if yname == "hw-info" { return "HwInfo" }
    return ""
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetSegmentPath() string {
    return "index-entry" + "[index='" + fmt.Sprintf("%v", indexEntry.Index) + "']"
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-info" {
        for _, c := range indexEntry.HwInfo {
            if indexEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo{}
        indexEntry.HwInfo = append(indexEntry.HwInfo, child)
        return &indexEntry.HwInfo[len(indexEntry.HwInfo)-1]
    }
    return nil
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indexEntry.HwInfo {
        children[indexEntry.HwInfo[i].GetSegmentPath()] = &indexEntry.HwInfo[i]
    }
    return children
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = indexEntry.Index
    leafs["l3protocol"] = indexEntry.L3Protocol
    leafs["l4protocol"] = indexEntry.L4Protocol
    leafs["intf-handle"] = indexEntry.IntfHandle
    leafs["intf-name"] = indexEntry.IntfName
    leafs["uidb-index"] = indexEntry.UidbIndex
    leafs["local-addr"] = indexEntry.LocalAddr
    leafs["local-prefix-len"] = indexEntry.LocalPrefixLen
    leafs["remote-addr"] = indexEntry.RemoteAddr
    leafs["remote-prefix-len"] = indexEntry.RemotePrefixLen
    leafs["vrf-id"] = indexEntry.VrfId
    leafs["u-value"] = indexEntry.UValue
    leafs["u-len"] = indexEntry.ULen
    leafs["local-port"] = indexEntry.LocalPort
    leafs["is-frag"] = indexEntry.IsFrag
    leafs["is-syn"] = indexEntry.IsSyn
    leafs["action"] = indexEntry.Action
    leafs["action-string"] = indexEntry.ActionString
    leafs["listener-tag"] = indexEntry.ListenerTag
    leafs["is-fgid"] = indexEntry.IsFgid
    leafs["is-vrf"] = indexEntry.IsVrf
    leafs["is-optimized"] = indexEntry.IsOptimized
    leafs["is-uidb-opt-bit"] = indexEntry.IsUidbOptBit
    leafs["fgid-or-sfp"] = indexEntry.FgidOrSfp
    leafs["remote-rack"] = indexEntry.RemoteRack
    leafs["rack-id"] = indexEntry.RackId
    leafs["rslot"] = indexEntry.Rslot
    leafs["cir"] = indexEntry.Cir
    leafs["flow-type"] = indexEntry.FlowType
    leafs["priority"] = indexEntry.Priority
    leafs["sid"] = indexEntry.Sid
    leafs["policer-avgrate"] = indexEntry.PolicerAvgrate
    leafs["policer-burst"] = indexEntry.PolicerBurst
    leafs["lookup-priority"] = indexEntry.LookupPriority
    leafs["storage-priority"] = indexEntry.StoragePriority
    leafs["num-tm-entries"] = indexEntry.NumTmEntries
    leafs["entry-ptr"] = indexEntry.EntryPtr
    leafs["entry-shadow-ptr"] = indexEntry.EntryShadowPtr
    leafs["list-node-ptr"] = indexEntry.ListNodePtr
    leafs["state"] = indexEntry.State
    leafs["retry-fail-cause"] = indexEntry.RetryFailCause
    leafs["num-retries"] = indexEntry.NumRetries
    leafs["min-ttl"] = indexEntry.MinTtl
    leafs["u-type"] = indexEntry.UType
    leafs["remote-fgid"] = indexEntry.RemoteFgid
    leafs["acl-str"] = indexEntry.AclStr
    leafs["no-stats"] = indexEntry.NoStats
    return leafs
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetBundleName() string { return "cisco_ios_xr" }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetYangName() string { return "index-entry" }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) SetParent(parent types.Entity) { indexEntry.parent = parent }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetParent() types.Entity { return indexEntry.parent }

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetParentYangName() string { return "index-entries" }

// LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo
// Per pipe type hardware info
type LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer Pointer. The type is interface{} with range: 0..4294967295.
    Policer interface{}

    // Stats Pointer. The type is interface{} with range: 0..4294967295.
    StatsPtr interface{}

    // Accepted Packets Counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepted interface{}

    // Dropped Packets Counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Dropped interface{}

    // Relative position in sorting order. The type is interface{} with range:
    // -2147483648..2147483647.
    SortStartOffset interface{}

    // Relative position in TCAM. The type is interface{} with range:
    // -2147483648..2147483647.
    TmStartOffset interface{}
}

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetFilter() yfilter.YFilter { return hwInfo.YFilter }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) SetFilter(yf yfilter.YFilter) { hwInfo.YFilter = yf }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetGoName(yname string) string {
    if yname == "policer" { return "Policer" }
    if yname == "stats-ptr" { return "StatsPtr" }
    if yname == "accepted" { return "Accepted" }
    if yname == "dropped" { return "Dropped" }
    if yname == "sort-start-offset" { return "SortStartOffset" }
    if yname == "tm-start-offset" { return "TmStartOffset" }
    return ""
}

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetSegmentPath() string {
    return "hw-info"
}

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policer"] = hwInfo.Policer
    leafs["stats-ptr"] = hwInfo.StatsPtr
    leafs["accepted"] = hwInfo.Accepted
    leafs["dropped"] = hwInfo.Dropped
    leafs["sort-start-offset"] = hwInfo.SortStartOffset
    leafs["tm-start-offset"] = hwInfo.TmStartOffset
    return leafs
}

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetBundleName() string { return "cisco_ios_xr" }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetYangName() string { return "hw-info" }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) SetParent(parent types.Entity) { hwInfo.parent = parent }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetParent() types.Entity { return hwInfo.parent }

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetParentYangName() string { return "index-entry" }

