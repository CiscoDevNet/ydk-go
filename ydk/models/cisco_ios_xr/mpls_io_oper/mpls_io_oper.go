// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-io package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpls-ea: MPLS IO EA operational data
//   mpls-ma: mpls ma
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_io_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_io_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-io-oper mpls-ea}", reflect.TypeOf(MplsEa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-io-oper:mpls-ea", reflect.TypeOf(MplsEa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-io-oper mpls-ma}", reflect.TypeOf(MplsMa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-io-oper:mpls-ma", reflect.TypeOf(MplsMa{}))
}

// MplsEa
// MPLS IO EA operational data
type MplsEa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NODE container class for MPLS IO EA operational data.
    Nodes MplsEa_Nodes
}

func (mplsEa *MplsEa) GetFilter() yfilter.YFilter { return mplsEa.YFilter }

func (mplsEa *MplsEa) SetFilter(yf yfilter.YFilter) { mplsEa.YFilter = yf }

func (mplsEa *MplsEa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (mplsEa *MplsEa) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-io-oper:mpls-ea"
}

func (mplsEa *MplsEa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &mplsEa.Nodes
    }
    return nil
}

func (mplsEa *MplsEa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &mplsEa.Nodes
    return children
}

func (mplsEa *MplsEa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsEa *MplsEa) GetBundleName() string { return "cisco_ios_xr" }

func (mplsEa *MplsEa) GetYangName() string { return "mpls-ea" }

func (mplsEa *MplsEa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsEa *MplsEa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsEa *MplsEa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsEa *MplsEa) SetParent(parent types.Entity) { mplsEa.parent = parent }

func (mplsEa *MplsEa) GetParent() types.Entity { return mplsEa.parent }

func (mplsEa *MplsEa) GetParentYangName() string { return "Cisco-IOS-XR-mpls-io-oper" }

// MplsEa_Nodes
// NODE container class for MPLS IO EA operational
// data
type MplsEa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node MPLS IO EA operational data. The type is slice of
    // MplsEa_Nodes_Node.
    Node []MplsEa_Nodes_Node
}

func (nodes *MplsEa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *MplsEa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *MplsEa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *MplsEa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *MplsEa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsEa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *MplsEa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *MplsEa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *MplsEa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *MplsEa_Nodes) GetYangName() string { return "nodes" }

func (nodes *MplsEa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *MplsEa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *MplsEa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *MplsEa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *MplsEa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *MplsEa_Nodes) GetParentYangName() string { return "mpls-ea" }

// MplsEa_Nodes_Node
// Per node MPLS IO EA operational data
type MplsEa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // MPLS IO EA Interfaces information .
    Interfaces MplsEa_Nodes_Node_Interfaces
}

func (node *MplsEa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *MplsEa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *MplsEa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *MplsEa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *MplsEa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *MplsEa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *MplsEa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *MplsEa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *MplsEa_Nodes_Node) GetYangName() string { return "node" }

func (node *MplsEa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *MplsEa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *MplsEa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *MplsEa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *MplsEa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *MplsEa_Nodes_Node) GetParentYangName() string { return "nodes" }

// MplsEa_Nodes_Node_Interfaces
// MPLS IO EA Interfaces information 
type MplsEa_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS IO EA NODE Interface data . The type is slice of
    // MplsEa_Nodes_Node_Interfaces_Interface.
    Interface []MplsEa_Nodes_Node_Interfaces_Interface
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsEa_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsEa_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsEa_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// MplsEa_Nodes_Node_Interfaces_Interface
// MPLS IO EA NODE Interface data 
type MplsEa_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MTU for fragmentation. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Bkp Label Stack Depth. The type is interface{} with range: 0..255.
    BkpLabelStackDepth interface{}

    // Srte Label Stack Depth. The type is interface{} with range: 0..255.
    SrteLabelStackDepth interface{}

    // Pri Label Stack Depth. The type is interface{} with range: 0..255.
    PriLabelStackDepth interface{}
}

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mtu" { return "Mtu" }
    if yname == "bkp-label-stack-depth" { return "BkpLabelStackDepth" }
    if yname == "srte-label-stack-depth" { return "SrteLabelStackDepth" }
    if yname == "pri-label-stack-depth" { return "PriLabelStackDepth" }
    return ""
}

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["mtu"] = self.Mtu
    leafs["bkp-label-stack-depth"] = self.BkpLabelStackDepth
    leafs["srte-label-stack-depth"] = self.SrteLabelStackDepth
    leafs["pri-label-stack-depth"] = self.PriLabelStackDepth
    return leafs
}

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsMa
// mpls ma
type MplsMa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NODE container class for MPLS IO MA operational data.
    Nodes MplsMa_Nodes
}

func (mplsMa *MplsMa) GetFilter() yfilter.YFilter { return mplsMa.YFilter }

func (mplsMa *MplsMa) SetFilter(yf yfilter.YFilter) { mplsMa.YFilter = yf }

func (mplsMa *MplsMa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (mplsMa *MplsMa) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-io-oper:mpls-ma"
}

func (mplsMa *MplsMa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &mplsMa.Nodes
    }
    return nil
}

func (mplsMa *MplsMa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &mplsMa.Nodes
    return children
}

func (mplsMa *MplsMa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsMa *MplsMa) GetBundleName() string { return "cisco_ios_xr" }

func (mplsMa *MplsMa) GetYangName() string { return "mpls-ma" }

func (mplsMa *MplsMa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsMa *MplsMa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsMa *MplsMa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsMa *MplsMa) SetParent(parent types.Entity) { mplsMa.parent = parent }

func (mplsMa *MplsMa) GetParent() types.Entity { return mplsMa.parent }

func (mplsMa *MplsMa) GetParentYangName() string { return "Cisco-IOS-XR-mpls-io-oper" }

// MplsMa_Nodes
// NODE container class for MPLS IO MA operational
// data
type MplsMa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node MPLS IO MA operational data. The type is slice of
    // MplsMa_Nodes_Node.
    Node []MplsMa_Nodes_Node
}

func (nodes *MplsMa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *MplsMa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *MplsMa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *MplsMa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *MplsMa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsMa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *MplsMa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *MplsMa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *MplsMa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *MplsMa_Nodes) GetYangName() string { return "nodes" }

func (nodes *MplsMa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *MplsMa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *MplsMa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *MplsMa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *MplsMa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *MplsMa_Nodes) GetParentYangName() string { return "mpls-ma" }

// MplsMa_Nodes_Node
// Per node MPLS IO MA operational data
type MplsMa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // MPLS IO MA Interfaces information .
    Interfaces MplsMa_Nodes_Node_Interfaces
}

func (node *MplsMa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *MplsMa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *MplsMa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *MplsMa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *MplsMa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *MplsMa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *MplsMa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *MplsMa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *MplsMa_Nodes_Node) GetYangName() string { return "node" }

func (node *MplsMa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *MplsMa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *MplsMa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *MplsMa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *MplsMa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *MplsMa_Nodes_Node) GetParentYangName() string { return "nodes" }

// MplsMa_Nodes_Node_Interfaces
// MPLS IO MA Interfaces information 
type MplsMa_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS IO MA NODE Interface data . The type is slice of
    // MplsMa_Nodes_Node_Interfaces_Interface.
    Interface []MplsMa_Nodes_Node_Interfaces_Interface
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsMa_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsMa_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsMa_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// MplsMa_Nodes_Node_Interfaces_Interface
// MPLS IO MA NODE Interface data 
type MplsMa_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MTU for fragmentation. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Bkp Label Stack Depth. The type is interface{} with range: 0..255.
    BkpLabelStackDepth interface{}

    // Srte Label Stack Depth. The type is interface{} with range: 0..255.
    SrteLabelStackDepth interface{}

    // Pri Label Stack Depth. The type is interface{} with range: 0..255.
    PriLabelStackDepth interface{}
}

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mtu" { return "Mtu" }
    if yname == "bkp-label-stack-depth" { return "BkpLabelStackDepth" }
    if yname == "srte-label-stack-depth" { return "SrteLabelStackDepth" }
    if yname == "pri-label-stack-depth" { return "PriLabelStackDepth" }
    return ""
}

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["mtu"] = self.Mtu
    leafs["bkp-label-stack-depth"] = self.BkpLabelStackDepth
    leafs["srte-label-stack-depth"] = self.SrteLabelStackDepth
    leafs["pri-label-stack-depth"] = self.PriLabelStackDepth
    return leafs
}

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

