// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-port-mapper package operational data.
// 
// This module contains definitions
// for the following management objects:
//   oor: Out of Resource Data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OOR data for available nodes.
    Nodes Oor_Nodes
}

func (oor *Oor) GetEntityData() *types.CommonEntityData {
    oor.EntityData.YFilter = oor.YFilter
    oor.EntityData.YangName = "oor"
    oor.EntityData.BundleName = "cisco_ios_xr"
    oor.EntityData.ParentYangName = "Cisco-IOS-XR-dnx-port-mapper-oper"
    oor.EntityData.SegmentPath = "Cisco-IOS-XR-dnx-port-mapper-oper:oor"
    oor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oor.EntityData.Children = types.NewOrderedMap()
    oor.EntityData.Children.Append("nodes", types.YChild{"Nodes", &oor.Nodes})
    oor.EntityData.Leafs = types.NewOrderedMap()

    oor.EntityData.YListKeys = []string {}

    return &(oor.EntityData)
}

// Oor_Nodes
// OOR data for available nodes
type Oor_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Oor_Nodes_Node.
    Node []*Oor_Nodes_Node
}

func (nodes *Oor_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "oor"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Oor_Nodes_Node
// DPA operational data for a particular node
type Oor_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Summary.
    Summary Oor_Nodes_Node_Summary

    // OOR Interface Information.
    InterfaceNames Oor_Nodes_Node_InterfaceNames
}

func (node *Oor_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &node.InterfaceNames})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Oor_Nodes_Node_Summary
// Summary
type Oor_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Oor_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("red", types.YLeaf{"Red", summary.Red})
    summary.EntityData.Leafs.Append("green", types.YLeaf{"Green", summary.Green})
    summary.EntityData.Leafs.Append("yel-low", types.YLeaf{"YelLow", summary.YelLow})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Oor_Nodes_Node_InterfaceNames
// OOR Interface Information
type Oor_Nodes_Node_InterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OOR Data for particular interface. The type is slice of
    // Oor_Nodes_Node_InterfaceNames_InterfaceName.
    InterfaceName []*Oor_Nodes_Node_InterfaceNames_InterfaceName
}

func (interfaceNames *Oor_Nodes_Node_InterfaceNames) GetEntityData() *types.CommonEntityData {
    interfaceNames.EntityData.YFilter = interfaceNames.YFilter
    interfaceNames.EntityData.YangName = "interface-names"
    interfaceNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceNames.EntityData.ParentYangName = "node"
    interfaceNames.EntityData.SegmentPath = "interface-names"
    interfaceNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceNames.EntityData.Children = types.NewOrderedMap()
    interfaceNames.EntityData.Children.Append("interface-name", types.YChild{"InterfaceName", nil})
    for i := range interfaceNames.InterfaceName {
        interfaceNames.EntityData.Children.Append(types.GetSegmentPath(interfaceNames.InterfaceName[i]), types.YChild{"InterfaceName", interfaceNames.InterfaceName[i]})
    }
    interfaceNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceNames.EntityData.YListKeys = []string {}

    return &(interfaceNames.EntityData)
}

// Oor_Nodes_Node_InterfaceNames_InterfaceName
// OOR Data for particular interface
type Oor_Nodes_Node_InterfaceNames_InterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // Interface details. The type is slice of
    // Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface.
    Interface []*Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface
}

func (interfaceName *Oor_Nodes_Node_InterfaceNames_InterfaceName) GetEntityData() *types.CommonEntityData {
    interfaceName.EntityData.YFilter = interfaceName.YFilter
    interfaceName.EntityData.YangName = "interface-name"
    interfaceName.EntityData.BundleName = "cisco_ios_xr"
    interfaceName.EntityData.ParentYangName = "interface-names"
    interfaceName.EntityData.SegmentPath = "interface-name" + types.AddKeyToken(interfaceName.InterfaceName, "interface-name")
    interfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceName.EntityData.Children = types.NewOrderedMap()
    interfaceName.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceName.Interface {
        interfaceName.EntityData.Children.Append(types.GetSegmentPath(interfaceName.Interface[i]), types.YChild{"Interface", interfaceName.Interface[i]})
    }
    interfaceName.EntityData.Leafs = types.NewOrderedMap()
    interfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceName.InterfaceName})

    interfaceName.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceName.EntityData)
}

// Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface
// Interface details
type Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Oor_Nodes_Node_InterfaceNames_InterfaceName_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-name"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-status", types.YLeaf{"InterfaceStatus", self.InterfaceStatus})
    self.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", self.TimeStamp})
    self.EntityData.Leafs.Append("npu-id", types.YLeaf{"NpuId", self.NpuId})
    self.EntityData.Leafs.Append("hardware-resource", types.YLeaf{"HardwareResource", self.HardwareResource})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

