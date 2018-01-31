// This module contains a collection of YANG definitions
// for Cisco IOS-XR cmproxy package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sdr-inventory-vm: Platform VM information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package cmproxy_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cmproxy_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cmproxy-oper sdr-inventory-vm}", reflect.TypeOf(SdrInventoryVm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm", reflect.TypeOf(SdrInventoryVm{}))
}

// SdrInventoryVm
// Platform VM information
type SdrInventoryVm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node directory.
    Nodes SdrInventoryVm_Nodes
}

func (sdrInventoryVm *SdrInventoryVm) GetFilter() yfilter.YFilter { return sdrInventoryVm.YFilter }

func (sdrInventoryVm *SdrInventoryVm) SetFilter(yf yfilter.YFilter) { sdrInventoryVm.YFilter = yf }

func (sdrInventoryVm *SdrInventoryVm) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (sdrInventoryVm *SdrInventoryVm) GetSegmentPath() string {
    return "Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm"
}

func (sdrInventoryVm *SdrInventoryVm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &sdrInventoryVm.Nodes
    }
    return nil
}

func (sdrInventoryVm *SdrInventoryVm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &sdrInventoryVm.Nodes
    return children
}

func (sdrInventoryVm *SdrInventoryVm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sdrInventoryVm *SdrInventoryVm) GetBundleName() string { return "cisco_ios_xr" }

func (sdrInventoryVm *SdrInventoryVm) GetYangName() string { return "sdr-inventory-vm" }

func (sdrInventoryVm *SdrInventoryVm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdrInventoryVm *SdrInventoryVm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdrInventoryVm *SdrInventoryVm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdrInventoryVm *SdrInventoryVm) SetParent(parent types.Entity) { sdrInventoryVm.parent = parent }

func (sdrInventoryVm *SdrInventoryVm) GetParent() types.Entity { return sdrInventoryVm.parent }

func (sdrInventoryVm *SdrInventoryVm) GetParentYangName() string { return "Cisco-IOS-XR-cmproxy-oper" }

// SdrInventoryVm_Nodes
// Node directory
type SdrInventoryVm_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node name. The type is slice of SdrInventoryVm_Nodes_Node.
    Node []SdrInventoryVm_Nodes_Node
}

func (nodes *SdrInventoryVm_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *SdrInventoryVm_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *SdrInventoryVm_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *SdrInventoryVm_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *SdrInventoryVm_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SdrInventoryVm_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *SdrInventoryVm_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *SdrInventoryVm_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *SdrInventoryVm_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *SdrInventoryVm_Nodes) GetYangName() string { return "nodes" }

func (nodes *SdrInventoryVm_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *SdrInventoryVm_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *SdrInventoryVm_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *SdrInventoryVm_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *SdrInventoryVm_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *SdrInventoryVm_Nodes) GetParentYangName() string { return "sdr-inventory-vm" }

// SdrInventoryVm_Nodes_Node
// Node name
type SdrInventoryVm_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // VM Information.
    NodeEntries SdrInventoryVm_Nodes_Node_NodeEntries
}

func (node *SdrInventoryVm_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *SdrInventoryVm_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *SdrInventoryVm_Nodes_Node) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "node-entries" { return "NodeEntries" }
    return ""
}

func (node *SdrInventoryVm_Nodes_Node) GetSegmentPath() string {
    return "node" + "[name='" + fmt.Sprintf("%v", node.Name) + "']"
}

func (node *SdrInventoryVm_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-entries" {
        return &node.NodeEntries
    }
    return nil
}

func (node *SdrInventoryVm_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-entries"] = &node.NodeEntries
    return children
}

func (node *SdrInventoryVm_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = node.Name
    return leafs
}

func (node *SdrInventoryVm_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *SdrInventoryVm_Nodes_Node) GetYangName() string { return "node" }

func (node *SdrInventoryVm_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *SdrInventoryVm_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *SdrInventoryVm_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *SdrInventoryVm_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *SdrInventoryVm_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *SdrInventoryVm_Nodes_Node) GetParentYangName() string { return "nodes" }

// SdrInventoryVm_Nodes_Node_NodeEntries
// VM Information
type SdrInventoryVm_Nodes_Node_NodeEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VM information for a node. The type is slice of
    // SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry.
    NodeEntry []SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetFilter() yfilter.YFilter { return nodeEntries.YFilter }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) SetFilter(yf yfilter.YFilter) { nodeEntries.YFilter = yf }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetGoName(yname string) string {
    if yname == "node-entry" { return "NodeEntry" }
    return ""
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetSegmentPath() string {
    return "node-entries"
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-entry" {
        for _, c := range nodeEntries.NodeEntry {
            if nodeEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry{}
        nodeEntries.NodeEntry = append(nodeEntries.NodeEntry, child)
        return &nodeEntries.NodeEntry[len(nodeEntries.NodeEntry)-1]
    }
    return nil
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeEntries.NodeEntry {
        children[nodeEntries.NodeEntry[i].GetSegmentPath()] = &nodeEntries.NodeEntry[i]
    }
    return children
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetBundleName() string { return "cisco_ios_xr" }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetYangName() string { return "node-entries" }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) SetParent(parent types.Entity) { nodeEntries.parent = parent }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetParent() types.Entity { return nodeEntries.parent }

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetParentYangName() string { return "node" }

// SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry
// VM information for a node
type SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // valid flag. The type is interface{} with range: 0..4294967295.
    Valid interface{}

    // card type. The type is interface{} with range: 0..4294967295.
    CardType interface{}

    // card type string. The type is string with length: 0..32.
    CardTypeString interface{}

    // node ID. The type is interface{} with range: -2147483648..2147483647.
    Nodeid interface{}

    // node name string. The type is string with length: 0..32.
    NodeName interface{}

    // partner node id. The type is interface{} with range:
    // -2147483648..2147483647.
    PartnerId interface{}

    // partner name string. The type is string with length: 0..32.
    PartnerName interface{}

    // redundancy state. The type is interface{} with range: 0..4294967295.
    RedState interface{}

    // redundancy state string. The type is string with length: 0..32.
    RedStateString interface{}

    // current software state. The type is interface{} with range: 0..4294967295.
    NodeSwState interface{}

    // current software state string. The type is string with length: 0..32.
    NodeSwStateString interface{}

    // previous software state. The type is interface{} with range: 0..4294967295.
    PrevSwState interface{}

    // previous software state string. The type is string with length: 0..32.
    PrevSwStateString interface{}

    // node IP address. The type is interface{} with range: 0..4294967295.
    NodeIp interface{}

    // node IPv4 address string. The type is string with length: 0..16.
    NodeIpv4String interface{}
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetFilter() yfilter.YFilter { return nodeEntry.YFilter }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) SetFilter(yf yfilter.YFilter) { nodeEntry.YFilter = yf }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "valid" { return "Valid" }
    if yname == "card-type" { return "CardType" }
    if yname == "card-type-string" { return "CardTypeString" }
    if yname == "nodeid" { return "Nodeid" }
    if yname == "node-name" { return "NodeName" }
    if yname == "partner-id" { return "PartnerId" }
    if yname == "partner-name" { return "PartnerName" }
    if yname == "red-state" { return "RedState" }
    if yname == "red-state-string" { return "RedStateString" }
    if yname == "node-sw-state" { return "NodeSwState" }
    if yname == "node-sw-state-string" { return "NodeSwStateString" }
    if yname == "prev-sw-state" { return "PrevSwState" }
    if yname == "prev-sw-state-string" { return "PrevSwStateString" }
    if yname == "node-ip" { return "NodeIp" }
    if yname == "node-ipv4-string" { return "NodeIpv4String" }
    return ""
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetSegmentPath() string {
    return "node-entry" + "[name='" + fmt.Sprintf("%v", nodeEntry.Name) + "']"
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = nodeEntry.Name
    leafs["valid"] = nodeEntry.Valid
    leafs["card-type"] = nodeEntry.CardType
    leafs["card-type-string"] = nodeEntry.CardTypeString
    leafs["nodeid"] = nodeEntry.Nodeid
    leafs["node-name"] = nodeEntry.NodeName
    leafs["partner-id"] = nodeEntry.PartnerId
    leafs["partner-name"] = nodeEntry.PartnerName
    leafs["red-state"] = nodeEntry.RedState
    leafs["red-state-string"] = nodeEntry.RedStateString
    leafs["node-sw-state"] = nodeEntry.NodeSwState
    leafs["node-sw-state-string"] = nodeEntry.NodeSwStateString
    leafs["prev-sw-state"] = nodeEntry.PrevSwState
    leafs["prev-sw-state-string"] = nodeEntry.PrevSwStateString
    leafs["node-ip"] = nodeEntry.NodeIp
    leafs["node-ipv4-string"] = nodeEntry.NodeIpv4String
    return leafs
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetBundleName() string { return "cisco_ios_xr" }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetYangName() string { return "node-entry" }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) SetParent(parent types.Entity) { nodeEntry.parent = parent }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetParent() types.Entity { return nodeEntry.parent }

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetParentYangName() string { return "node-entries" }

