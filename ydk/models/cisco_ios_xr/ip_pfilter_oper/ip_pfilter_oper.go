// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-pfilter package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pfilter-ma: Root class of PfilterMa Oper schema
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_pfilter_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_pfilter_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-pfilter-oper pfilter-ma}", reflect.TypeOf(PfilterMa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-pfilter-oper:pfilter-ma", reflect.TypeOf(PfilterMa{}))
}

// PfilterMa
// Root class of PfilterMa Oper schema
type PfilterMa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific operational data.
    Nodes PfilterMa_Nodes
}

func (pfilterMa *PfilterMa) GetFilter() yfilter.YFilter { return pfilterMa.YFilter }

func (pfilterMa *PfilterMa) SetFilter(yf yfilter.YFilter) { pfilterMa.YFilter = yf }

func (pfilterMa *PfilterMa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (pfilterMa *PfilterMa) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-oper:pfilter-ma"
}

func (pfilterMa *PfilterMa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &pfilterMa.Nodes
    }
    return nil
}

func (pfilterMa *PfilterMa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &pfilterMa.Nodes
    return children
}

func (pfilterMa *PfilterMa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pfilterMa *PfilterMa) GetBundleName() string { return "cisco_ios_xr" }

func (pfilterMa *PfilterMa) GetYangName() string { return "pfilter-ma" }

func (pfilterMa *PfilterMa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pfilterMa *PfilterMa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pfilterMa *PfilterMa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pfilterMa *PfilterMa) SetParent(parent types.Entity) { pfilterMa.parent = parent }

func (pfilterMa *PfilterMa) GetParent() types.Entity { return pfilterMa.parent }

func (pfilterMa *PfilterMa) GetParentYangName() string { return "Cisco-IOS-XR-ip-pfilter-oper" }

// PfilterMa_Nodes
// Node-specific operational data
type PfilterMa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PfilterMa operational data for a particular node. The type is slice of
    // PfilterMa_Nodes_Node.
    Node []PfilterMa_Nodes_Node
}

func (nodes *PfilterMa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PfilterMa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PfilterMa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PfilterMa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PfilterMa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PfilterMa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PfilterMa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PfilterMa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PfilterMa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PfilterMa_Nodes) GetYangName() string { return "nodes" }

func (nodes *PfilterMa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PfilterMa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PfilterMa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PfilterMa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PfilterMa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PfilterMa_Nodes) GetParentYangName() string { return "pfilter-ma" }

// PfilterMa_Nodes_Node
// PfilterMa operational data for a particular
// node
type PfilterMa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Operational data for pfilter.
    Process PfilterMa_Nodes_Node_Process
}

func (node *PfilterMa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PfilterMa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PfilterMa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "process" { return "Process" }
    return ""
}

func (node *PfilterMa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *PfilterMa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process" {
        return &node.Process
    }
    return nil
}

func (node *PfilterMa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process"] = &node.Process
    return children
}

func (node *PfilterMa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *PfilterMa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PfilterMa_Nodes_Node) GetYangName() string { return "node" }

func (node *PfilterMa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PfilterMa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PfilterMa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PfilterMa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PfilterMa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PfilterMa_Nodes_Node) GetParentYangName() string { return "nodes" }

// PfilterMa_Nodes_Node_Process
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    Ipv6 PfilterMa_Nodes_Node_Process_Ipv6

    // Operational data for pfilter.
    Ipv4 PfilterMa_Nodes_Node_Process_Ipv4
}

func (process *PfilterMa_Nodes_Node_Process) GetFilter() yfilter.YFilter { return process.YFilter }

func (process *PfilterMa_Nodes_Node_Process) SetFilter(yf yfilter.YFilter) { process.YFilter = yf }

func (process *PfilterMa_Nodes_Node_Process) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (process *PfilterMa_Nodes_Node_Process) GetSegmentPath() string {
    return "process"
}

func (process *PfilterMa_Nodes_Node_Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &process.Ipv6
    }
    if childYangName == "ipv4" {
        return &process.Ipv4
    }
    return nil
}

func (process *PfilterMa_Nodes_Node_Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &process.Ipv6
    children["ipv4"] = &process.Ipv4
    return children
}

func (process *PfilterMa_Nodes_Node_Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (process *PfilterMa_Nodes_Node_Process) GetBundleName() string { return "cisco_ios_xr" }

func (process *PfilterMa_Nodes_Node_Process) GetYangName() string { return "process" }

func (process *PfilterMa_Nodes_Node_Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (process *PfilterMa_Nodes_Node_Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (process *PfilterMa_Nodes_Node_Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (process *PfilterMa_Nodes_Node_Process) SetParent(parent types.Entity) { process.parent = parent }

func (process *PfilterMa_Nodes_Node_Process) GetParent() types.Entity { return process.parent }

func (process *PfilterMa_Nodes_Node_Process) GetParentYangName() string { return "node" }

// PfilterMa_Nodes_Node_Process_Ipv6
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    AclInfoTable PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetGoName(yname string) string {
    if yname == "acl-info-table" { return "AclInfoTable" }
    return ""
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-info-table" {
        return &ipv6.AclInfoTable
    }
    return nil
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl-info-table"] = &ipv6.AclInfoTable
    return children
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetParentYangName() string { return "process" }

// PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    InterfaceInfos PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetFilter() yfilter.YFilter { return aclInfoTable.YFilter }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) SetFilter(yf yfilter.YFilter) { aclInfoTable.YFilter = yf }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetGoName(yname string) string {
    if yname == "interface-infos" { return "InterfaceInfos" }
    return ""
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetSegmentPath() string {
    return "acl-info-table"
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-infos" {
        return &aclInfoTable.InterfaceInfos
    }
    return nil
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-infos"] = &aclInfoTable.InterfaceInfos
    return children
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetBundleName() string { return "cisco_ios_xr" }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetYangName() string { return "acl-info-table" }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) SetParent(parent types.Entity) { aclInfoTable.parent = parent }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetParent() types.Entity { return aclInfoTable.parent }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetParentYangName() string { return "ipv6" }

// PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter in bag. The type is slice of
    // PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo.
    InterfaceInfo []PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetFilter() yfilter.YFilter { return interfaceInfos.YFilter }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) SetFilter(yf yfilter.YFilter) { interfaceInfos.YFilter = yf }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetGoName(yname string) string {
    if yname == "interface-info" { return "InterfaceInfo" }
    return ""
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetSegmentPath() string {
    return "interface-infos"
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-info" {
        for _, c := range interfaceInfos.InterfaceInfo {
            if interfaceInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo{}
        interfaceInfos.InterfaceInfo = append(interfaceInfos.InterfaceInfo, child)
        return &interfaceInfos.InterfaceInfo[len(interfaceInfos.InterfaceInfo)-1]
    }
    return nil
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceInfos.InterfaceInfo {
        children[interfaceInfos.InterfaceInfo[i].GetSegmentPath()] = &interfaceInfos.InterfaceInfo[i]
    }
    return children
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetYangName() string { return "interface-infos" }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) SetParent(parent types.Entity) { interfaceInfos.parent = parent }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetParent() types.Entity { return interfaceInfos.parent }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetParentYangName() string { return "acl-info-table" }

// PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo
// Operational data for pfilter in bag
type PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // acl information. The type is string.
    AclInfo interface{}
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetFilter() yfilter.YFilter { return interfaceInfo.YFilter }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) SetFilter(yf yfilter.YFilter) { interfaceInfo.YFilter = yf }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "acl-info" { return "AclInfo" }
    return ""
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetSegmentPath() string {
    return "interface-info" + "[interface-name='" + fmt.Sprintf("%v", interfaceInfo.InterfaceName) + "']"
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceInfo.InterfaceName
    leafs["acl-info"] = interfaceInfo.AclInfo
    return leafs
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetYangName() string { return "interface-info" }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) SetParent(parent types.Entity) { interfaceInfo.parent = parent }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetParent() types.Entity { return interfaceInfo.parent }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetParentYangName() string { return "interface-infos" }

// PfilterMa_Nodes_Node_Process_Ipv4
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    AclInfoTable PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetGoName(yname string) string {
    if yname == "acl-info-table" { return "AclInfoTable" }
    return ""
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-info-table" {
        return &ipv4.AclInfoTable
    }
    return nil
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl-info-table"] = &ipv4.AclInfoTable
    return children
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetParentYangName() string { return "process" }

// PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    InterfaceInfos PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetFilter() yfilter.YFilter { return aclInfoTable.YFilter }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) SetFilter(yf yfilter.YFilter) { aclInfoTable.YFilter = yf }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetGoName(yname string) string {
    if yname == "interface-infos" { return "InterfaceInfos" }
    return ""
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetSegmentPath() string {
    return "acl-info-table"
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-infos" {
        return &aclInfoTable.InterfaceInfos
    }
    return nil
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-infos"] = &aclInfoTable.InterfaceInfos
    return children
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetBundleName() string { return "cisco_ios_xr" }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetYangName() string { return "acl-info-table" }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) SetParent(parent types.Entity) { aclInfoTable.parent = parent }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetParent() types.Entity { return aclInfoTable.parent }

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetParentYangName() string { return "ipv4" }

// PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for pfilter in bag. The type is slice of
    // PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo.
    InterfaceInfo []PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetFilter() yfilter.YFilter { return interfaceInfos.YFilter }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) SetFilter(yf yfilter.YFilter) { interfaceInfos.YFilter = yf }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetGoName(yname string) string {
    if yname == "interface-info" { return "InterfaceInfo" }
    return ""
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetSegmentPath() string {
    return "interface-infos"
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-info" {
        for _, c := range interfaceInfos.InterfaceInfo {
            if interfaceInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo{}
        interfaceInfos.InterfaceInfo = append(interfaceInfos.InterfaceInfo, child)
        return &interfaceInfos.InterfaceInfo[len(interfaceInfos.InterfaceInfo)-1]
    }
    return nil
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceInfos.InterfaceInfo {
        children[interfaceInfos.InterfaceInfo[i].GetSegmentPath()] = &interfaceInfos.InterfaceInfo[i]
    }
    return children
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetYangName() string { return "interface-infos" }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) SetParent(parent types.Entity) { interfaceInfos.parent = parent }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetParent() types.Entity { return interfaceInfos.parent }

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetParentYangName() string { return "acl-info-table" }

// PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo
// Operational data for pfilter in bag
type PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // acl information. The type is string.
    AclInfo interface{}
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetFilter() yfilter.YFilter { return interfaceInfo.YFilter }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) SetFilter(yf yfilter.YFilter) { interfaceInfo.YFilter = yf }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "acl-info" { return "AclInfo" }
    return ""
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetSegmentPath() string {
    return "interface-info" + "[interface-name='" + fmt.Sprintf("%v", interfaceInfo.InterfaceName) + "']"
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceInfo.InterfaceName
    leafs["acl-info"] = interfaceInfo.AclInfo
    return leafs
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetYangName() string { return "interface-info" }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) SetParent(parent types.Entity) { interfaceInfo.parent = parent }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetParent() types.Entity { return interfaceInfo.parent }

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetParentYangName() string { return "interface-infos" }

