// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-port-mapper package operational data.
// 
// This module contains definitions
// for the following management objects:
//   oor: Out of Resource Data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dnx_port_mapper_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dnx_port_mapper_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dnx-port-mapper-oper oor}", reflect.TypeOf(Oor{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dnx-port-mapper-oper:oor", reflect.TypeOf(Oor{}))
}

// Oor
// Out of Resource Data
type Oor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OOR data for available nodes.
    Nodes Oor_Nodes
}

func (oor *Oor) GetFilter() yfilter.YFilter { return oor.YFilter }

func (oor *Oor) SetFilter(yf yfilter.YFilter) { oor.YFilter = yf }

func (oor *Oor) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (oor *Oor) GetSegmentPath() string {
    return "Cisco-IOS-XR-dnx-port-mapper-oper:oor"
}

func (oor *Oor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &oor.Nodes
    }
    return nil
}

func (oor *Oor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &oor.Nodes
    return children
}

func (oor *Oor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oor *Oor) GetBundleName() string { return "cisco_ios_xr" }

func (oor *Oor) GetYangName() string { return "oor" }

func (oor *Oor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oor *Oor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oor *Oor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oor *Oor) SetParent(parent types.Entity) { oor.parent = parent }

func (oor *Oor) GetParent() types.Entity { return oor.parent }

func (oor *Oor) GetParentYangName() string { return "Cisco-IOS-XR-dnx-port-mapper-oper" }

// Oor_Nodes
// OOR data for available nodes
type Oor_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Oor_Nodes_Node.
    Node []Oor_Nodes_Node
}

func (nodes *Oor_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Oor_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Oor_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Oor_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Oor_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Oor_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Oor_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Oor_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Oor_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Oor_Nodes) GetYangName() string { return "nodes" }

func (nodes *Oor_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Oor_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Oor_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Oor_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Oor_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Oor_Nodes) GetParentYangName() string { return "oor" }

// Oor_Nodes_Node
// DPA operational data for a particular node
type Oor_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Summary.
    Summary Oor_Nodes_Node_Summary

    // OOR Interface Information.
    InterfaceNames Oor_Nodes_Node_InterfaceNames
}

func (node *Oor_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Oor_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Oor_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "summary" { return "Summary" }
    if yname == "interface-names" { return "InterfaceNames" }
    return ""
}

func (node *Oor_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Oor_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &node.Summary
    }
    if childYangName == "interface-names" {
        return &node.InterfaceNames
    }
    return nil
}

func (node *Oor_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &node.Summary
    children["interface-names"] = &node.InterfaceNames
    return children
}

func (node *Oor_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Oor_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Oor_Nodes_Node) GetYangName() string { return "node" }

func (node *Oor_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Oor_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Oor_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Oor_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Oor_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Oor_Nodes_Node) GetParentYangName() string { return "nodes" }

// Oor_Nodes_Node_Summary
// Summary
type Oor_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interfaces in red state. The type is interface{} with range: 0..4294967295.
    Red interface{}

    // interfaces in green state. The type is interface{} with range:
    // 0..4294967295.
    Green interface{}

    // interfaces in yellow state. The type is interface{} with range:
    // 0..4294967295.
    YelLow interface{}
}

func (summary *Oor_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Oor_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Oor_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "red" { return "Red" }
    if yname == "green" { return "Green" }
    if yname == "yel-low" { return "YelLow" }
    return ""
}

func (summary *Oor_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Oor_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Oor_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Oor_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["red"] = summary.Red
    leafs["green"] = summary.Green
    leafs["yel-low"] = summary.YelLow
    return leafs
}

func (summary *Oor_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Oor_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *Oor_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Oor_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Oor_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Oor_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Oor_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Oor_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// Oor_Nodes_Node_InterfaceNames
// OOR Interface Information
type Oor_Nodes_Node_InterfaceNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OOR Data for particular interface. The type is slice of
    // Oor_Nodes_Node_InterfaceNames_InterfaceName.
    InterfaceName []Oor_Nodes_Node_InterfaceNames_InterfaceName
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetFilter() yfilter.YFilter { return interfaceNames.YFilter }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) SetFilter(yf yfilter.YFilter) { interfaceNames.YFilter = yf }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetSegmentPath() string {
    return "interface-names"
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-name" {
        for _, c := range interfaceNames.InterfaceName {
            if interfaceNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Oor_Nodes_Node_InterfaceNames_InterfaceName{}
        interfaceNames.InterfaceName = append(interfaceNames.InterfaceName, child)
        return &interfaceNames.InterfaceName[len(interfaceNames.InterfaceName)-1]
    }
    return nil
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceNames.InterfaceName {
        children[interfaceNames.InterfaceName[i].GetSegmentPath()] = &interfaceNames.InterfaceName[i]
    }
    return children
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetYangName() string { return "interface-names" }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) SetParent(parent types.Entity) { interfaceNames.parent = parent }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetParent() types.Entity { return interfaceNames.parent }

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetParentYangName() string { return "node" }

// Oor_Nodes_Node_InterfaceNames_InterfaceName
// OOR Data for particular interface
type Oor_Nodes_Node_InterfaceNames_InterfaceName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // Interface details. The type is slice of
    // Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface.
    Interface []Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetFilter() yfilter.YFilter { return interfaceName.YFilter }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) SetFilter(yf yfilter.YFilter) { interfaceName.YFilter = yf }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetSegmentPath() string {
    return "interface-name" + "[interface-name='" + fmt.Sprintf("%v", interfaceName.InterfaceName) + "']"
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaceName.Interface {
            if interfaceName.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface{}
        interfaceName.Interface = append(interfaceName.Interface, child)
        return &interfaceName.Interface[len(interfaceName.Interface)-1]
    }
    return nil
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceName.Interface {
        children[interfaceName.Interface[i].GetSegmentPath()] = &interfaceName.Interface[i]
    }
    return children
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceName.InterfaceName
    return leafs
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetYangName() string { return "interface-name" }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) SetParent(parent types.Entity) { interfaceName.parent = parent }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetParent() types.Entity { return interfaceName.parent }

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetParentYangName() string { return "interface-names" }

// Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface
// Interface details
type Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the interface. The type is string.
    InterfaceName interface{}

    // The current state of the interface. The type is string.
    InterfaceStatus interface{}

    // Timestamp. The type is string.
    TimeStamp interface{}

    // Npuid of the interface. The type is string.
    NpuId interface{}

    // Type of harware resoruce. The type is string.
    HardwareResource interface{}
}

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-status" { return "InterfaceStatus" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "npu-id" { return "NpuId" }
    if yname == "hardware-resource" { return "HardwareResource" }
    return ""
}

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-status"] = self.InterfaceStatus
    leafs["time-stamp"] = self.TimeStamp
    leafs["npu-id"] = self.NpuId
    leafs["hardware-resource"] = self.HardwareResource
    return leafs
}

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetYangName() string { return "interface" }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetParent() types.Entity { return self.parent }

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetParentYangName() string { return "interface-name" }

