// This module contains a collection of YANG definitions
// for Cisco IOS-XR pppoe-ea package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pppoe-ea: PPPoE Ea data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pppoe_ea_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pppoe_ea_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pppoe-ea-oper pppoe-ea}", reflect.TypeOf(PppoeEa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pppoe-ea-oper:pppoe-ea", reflect.TypeOf(PppoeEa{}))
}

// PppoeEa
// PPPoE Ea data
type PppoeEa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPOE_EA list of nodes.
    Nodes PppoeEa_Nodes
}

func (pppoeEa *PppoeEa) GetFilter() yfilter.YFilter { return pppoeEa.YFilter }

func (pppoeEa *PppoeEa) SetFilter(yf yfilter.YFilter) { pppoeEa.YFilter = yf }

func (pppoeEa *PppoeEa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (pppoeEa *PppoeEa) GetSegmentPath() string {
    return "Cisco-IOS-XR-pppoe-ea-oper:pppoe-ea"
}

func (pppoeEa *PppoeEa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &pppoeEa.Nodes
    }
    return nil
}

func (pppoeEa *PppoeEa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &pppoeEa.Nodes
    return children
}

func (pppoeEa *PppoeEa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pppoeEa *PppoeEa) GetBundleName() string { return "cisco_ios_xr" }

func (pppoeEa *PppoeEa) GetYangName() string { return "pppoe-ea" }

func (pppoeEa *PppoeEa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoeEa *PppoeEa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoeEa *PppoeEa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoeEa *PppoeEa) SetParent(parent types.Entity) { pppoeEa.parent = parent }

func (pppoeEa *PppoeEa) GetParent() types.Entity { return pppoeEa.parent }

func (pppoeEa *PppoeEa) GetParentYangName() string { return "Cisco-IOS-XR-pppoe-ea-oper" }

// PppoeEa_Nodes
// PPPOE_EA list of nodes
type PppoeEa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPOE-EA operational data for a particular node. The type is slice of
    // PppoeEa_Nodes_Node.
    Node []PppoeEa_Nodes_Node
}

func (nodes *PppoeEa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PppoeEa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PppoeEa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PppoeEa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PppoeEa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeEa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PppoeEa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PppoeEa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PppoeEa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PppoeEa_Nodes) GetYangName() string { return "nodes" }

func (nodes *PppoeEa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PppoeEa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PppoeEa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PppoeEa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PppoeEa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PppoeEa_Nodes) GetParentYangName() string { return "pppoe-ea" }

// PppoeEa_Nodes_Node
// PPPOE-EA operational data for a particular node
type PppoeEa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // PPPoE parent interface info.
    ParentInterfaceIds PppoeEa_Nodes_Node_ParentInterfaceIds

    // PPPoE interface info.
    InterfaceIds PppoeEa_Nodes_Node_InterfaceIds
}

func (node *PppoeEa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PppoeEa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PppoeEa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "parent-interface-ids" { return "ParentInterfaceIds" }
    if yname == "interface-ids" { return "InterfaceIds" }
    return ""
}

func (node *PppoeEa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *PppoeEa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "parent-interface-ids" {
        return &node.ParentInterfaceIds
    }
    if childYangName == "interface-ids" {
        return &node.InterfaceIds
    }
    return nil
}

func (node *PppoeEa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["parent-interface-ids"] = &node.ParentInterfaceIds
    children["interface-ids"] = &node.InterfaceIds
    return children
}

func (node *PppoeEa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *PppoeEa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PppoeEa_Nodes_Node) GetYangName() string { return "node" }

func (node *PppoeEa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PppoeEa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PppoeEa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PppoeEa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PppoeEa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PppoeEa_Nodes_Node) GetParentYangName() string { return "nodes" }

// PppoeEa_Nodes_Node_ParentInterfaceIds
// PPPoE parent interface info
type PppoeEa_Nodes_Node_ParentInterfaceIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE parent interface info. The type is slice of
    // PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId.
    ParentInterfaceId []PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetFilter() yfilter.YFilter { return parentInterfaceIds.YFilter }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) SetFilter(yf yfilter.YFilter) { parentInterfaceIds.YFilter = yf }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetGoName(yname string) string {
    if yname == "parent-interface-id" { return "ParentInterfaceId" }
    return ""
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetSegmentPath() string {
    return "parent-interface-ids"
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "parent-interface-id" {
        for _, c := range parentInterfaceIds.ParentInterfaceId {
            if parentInterfaceIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId{}
        parentInterfaceIds.ParentInterfaceId = append(parentInterfaceIds.ParentInterfaceId, child)
        return &parentInterfaceIds.ParentInterfaceId[len(parentInterfaceIds.ParentInterfaceId)-1]
    }
    return nil
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range parentInterfaceIds.ParentInterfaceId {
        children[parentInterfaceIds.ParentInterfaceId[i].GetSegmentPath()] = &parentInterfaceIds.ParentInterfaceId[i]
    }
    return children
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetBundleName() string { return "cisco_ios_xr" }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetYangName() string { return "parent-interface-ids" }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) SetParent(parent types.Entity) { parentInterfaceIds.parent = parent }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetParent() types.Entity { return parentInterfaceIds.parent }

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetParentYangName() string { return "node" }

// PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId
// PPPoE parent interface info
type PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ParentInterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Is in sync. The type is bool.
    IsInSync interface{}

    // SRG VMac-address.
    SrgvMac PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetFilter() yfilter.YFilter { return parentInterfaceId.YFilter }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) SetFilter(yf yfilter.YFilter) { parentInterfaceId.YFilter = yf }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetGoName(yname string) string {
    if yname == "parent-interface-name" { return "ParentInterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "is-in-sync" { return "IsInSync" }
    if yname == "srgv-mac" { return "SrgvMac" }
    return ""
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetSegmentPath() string {
    return "parent-interface-id" + "[parent-interface-name='" + fmt.Sprintf("%v", parentInterfaceId.ParentInterfaceName) + "']"
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srgv-mac" {
        return &parentInterfaceId.SrgvMac
    }
    return nil
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srgv-mac"] = &parentInterfaceId.SrgvMac
    return children
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["parent-interface-name"] = parentInterfaceId.ParentInterfaceName
    leafs["interface"] = parentInterfaceId.Interface
    leafs["is-in-sync"] = parentInterfaceId.IsInSync
    return leafs
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetBundleName() string { return "cisco_ios_xr" }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetYangName() string { return "parent-interface-id" }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) SetParent(parent types.Entity) { parentInterfaceId.parent = parent }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetParent() types.Entity { return parentInterfaceId.parent }

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetParentYangName() string { return "parent-interface-ids" }

// PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac
// SRG VMac-address
type PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Macaddr interface{}
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetFilter() yfilter.YFilter { return srgvMac.YFilter }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) SetFilter(yf yfilter.YFilter) { srgvMac.YFilter = yf }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetSegmentPath() string {
    return "srgv-mac"
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = srgvMac.Macaddr
    return leafs
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetBundleName() string { return "cisco_ios_xr" }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetYangName() string { return "srgv-mac" }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) SetParent(parent types.Entity) { srgvMac.parent = parent }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetParent() types.Entity { return srgvMac.parent }

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetParentYangName() string { return "parent-interface-id" }

// PppoeEa_Nodes_Node_InterfaceIds
// PPPoE interface info
type PppoeEa_Nodes_Node_InterfaceIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE interface info. The type is slice of
    // PppoeEa_Nodes_Node_InterfaceIds_InterfaceId.
    InterfaceId []PppoeEa_Nodes_Node_InterfaceIds_InterfaceId
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetFilter() yfilter.YFilter { return interfaceIds.YFilter }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) SetFilter(yf yfilter.YFilter) { interfaceIds.YFilter = yf }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetGoName(yname string) string {
    if yname == "interface-id" { return "InterfaceId" }
    return ""
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetSegmentPath() string {
    return "interface-ids"
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-id" {
        for _, c := range interfaceIds.InterfaceId {
            if interfaceIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeEa_Nodes_Node_InterfaceIds_InterfaceId{}
        interfaceIds.InterfaceId = append(interfaceIds.InterfaceId, child)
        return &interfaceIds.InterfaceId[len(interfaceIds.InterfaceId)-1]
    }
    return nil
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceIds.InterfaceId {
        children[interfaceIds.InterfaceId[i].GetSegmentPath()] = &interfaceIds.InterfaceId[i]
    }
    return children
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetYangName() string { return "interface-ids" }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) SetParent(parent types.Entity) { interfaceIds.parent = parent }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetParent() types.Entity { return interfaceIds.parent }

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetParentYangName() string { return "node" }

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId
// PPPoE interface info
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Session ID. The type is interface{} with range: 0..65535.
    SessionId interface{}

    // Parent Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterface interface{}

    // Is Priority Set. The type is bool.
    IsPrioritySet interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Is in sync. The type is bool.
    IsInSync interface{}

    // Is Platform created. The type is bool.
    IsPlatformCreated interface{}

    // VLAN Ids. The type is slice of interface{} with range: 0..65535.
    Vlanid []interface{}

    // Peer Mac-address.
    PeerMac PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac

    // Local Mac-address.
    LocalMac PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac

    // SRG VMac-address.
    SrgvMac PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac
}

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetFilter() yfilter.YFilter { return interfaceId.YFilter }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) SetFilter(yf yfilter.YFilter) { interfaceId.YFilter = yf }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "session-id" { return "SessionId" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "is-priority-set" { return "IsPrioritySet" }
    if yname == "priority" { return "Priority" }
    if yname == "is-in-sync" { return "IsInSync" }
    if yname == "is-platform-created" { return "IsPlatformCreated" }
    if yname == "vlanid" { return "Vlanid" }
    if yname == "peer-mac" { return "PeerMac" }
    if yname == "local-mac" { return "LocalMac" }
    if yname == "srgv-mac" { return "SrgvMac" }
    return ""
}

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetSegmentPath() string {
    return "interface-id" + "[interface-name='" + fmt.Sprintf("%v", interfaceId.InterfaceName) + "']"
}

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-mac" {
        return &interfaceId.PeerMac
    }
    if childYangName == "local-mac" {
        return &interfaceId.LocalMac
    }
    if childYangName == "srgv-mac" {
        return &interfaceId.SrgvMac
    }
    return nil
}

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-mac"] = &interfaceId.PeerMac
    children["local-mac"] = &interfaceId.LocalMac
    children["srgv-mac"] = &interfaceId.SrgvMac
    return children
}

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceId.InterfaceName
    leafs["interface"] = interfaceId.Interface
    leafs["session-id"] = interfaceId.SessionId
    leafs["parent-interface"] = interfaceId.ParentInterface
    leafs["is-priority-set"] = interfaceId.IsPrioritySet
    leafs["priority"] = interfaceId.Priority
    leafs["is-in-sync"] = interfaceId.IsInSync
    leafs["is-platform-created"] = interfaceId.IsPlatformCreated
    leafs["vlanid"] = interfaceId.Vlanid
    return leafs
}

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetYangName() string { return "interface-id" }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) SetParent(parent types.Entity) { interfaceId.parent = parent }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetParent() types.Entity { return interfaceId.parent }

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetParentYangName() string { return "interface-ids" }

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac
// Peer Mac-address
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Macaddr interface{}
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetFilter() yfilter.YFilter { return peerMac.YFilter }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) SetFilter(yf yfilter.YFilter) { peerMac.YFilter = yf }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetSegmentPath() string {
    return "peer-mac"
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = peerMac.Macaddr
    return leafs
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetBundleName() string { return "cisco_ios_xr" }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetYangName() string { return "peer-mac" }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) SetParent(parent types.Entity) { peerMac.parent = parent }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetParent() types.Entity { return peerMac.parent }

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetParentYangName() string { return "interface-id" }

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac
// Local Mac-address
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Macaddr interface{}
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetFilter() yfilter.YFilter { return localMac.YFilter }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) SetFilter(yf yfilter.YFilter) { localMac.YFilter = yf }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetSegmentPath() string {
    return "local-mac"
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = localMac.Macaddr
    return leafs
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetBundleName() string { return "cisco_ios_xr" }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetYangName() string { return "local-mac" }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) SetParent(parent types.Entity) { localMac.parent = parent }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetParent() types.Entity { return localMac.parent }

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetParentYangName() string { return "interface-id" }

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac
// SRG VMac-address
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Macaddr interface{}
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetFilter() yfilter.YFilter { return srgvMac.YFilter }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) SetFilter(yf yfilter.YFilter) { srgvMac.YFilter = yf }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetSegmentPath() string {
    return "srgv-mac"
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = srgvMac.Macaddr
    return leafs
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetBundleName() string { return "cisco_ios_xr" }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetYangName() string { return "srgv-mac" }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) SetParent(parent types.Entity) { srgvMac.parent = parent }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetParent() types.Entity { return srgvMac.parent }

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetParentYangName() string { return "interface-id" }

