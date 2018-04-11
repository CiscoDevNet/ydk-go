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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NODE container class for MPLS IO EA operational data.
    Nodes MplsEa_Nodes
}

func (mplsEa *MplsEa) GetEntityData() *types.CommonEntityData {
    mplsEa.EntityData.YFilter = mplsEa.YFilter
    mplsEa.EntityData.YangName = "mpls-ea"
    mplsEa.EntityData.BundleName = "cisco_ios_xr"
    mplsEa.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-io-oper"
    mplsEa.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-io-oper:mpls-ea"
    mplsEa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsEa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsEa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsEa.EntityData.Children = make(map[string]types.YChild)
    mplsEa.EntityData.Children["nodes"] = types.YChild{"Nodes", &mplsEa.Nodes}
    mplsEa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsEa.EntityData)
}

// MplsEa_Nodes
// NODE container class for MPLS IO EA operational
// data
type MplsEa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node MPLS IO EA operational data. The type is slice of
    // MplsEa_Nodes_Node.
    Node []MplsEa_Nodes_Node
}

func (nodes *MplsEa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "mpls-ea"
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

// MplsEa_Nodes_Node
// Per node MPLS IO EA operational data
type MplsEa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // MPLS IO EA Interfaces information .
    Interfaces MplsEa_Nodes_Node_Interfaces
}

func (node *MplsEa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// MplsEa_Nodes_Node_Interfaces
// MPLS IO EA Interfaces information 
type MplsEa_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS IO EA NODE Interface data . The type is slice of
    // MplsEa_Nodes_Node_Interfaces_Interface_.
    Interface_ []MplsEa_Nodes_Node_Interfaces_Interface
}

func (interfaces *MplsEa_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// MplsEa_Nodes_Node_Interfaces_Interface
// MPLS IO EA NODE Interface data 
type MplsEa_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *MplsEa_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", self.Mtu}
    self.EntityData.Leafs["bkp-label-stack-depth"] = types.YLeaf{"BkpLabelStackDepth", self.BkpLabelStackDepth}
    self.EntityData.Leafs["srte-label-stack-depth"] = types.YLeaf{"SrteLabelStackDepth", self.SrteLabelStackDepth}
    self.EntityData.Leafs["pri-label-stack-depth"] = types.YLeaf{"PriLabelStackDepth", self.PriLabelStackDepth}
    return &(self.EntityData)
}

// MplsMa
// mpls ma
type MplsMa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NODE container class for MPLS IO MA operational data.
    Nodes MplsMa_Nodes
}

func (mplsMa *MplsMa) GetEntityData() *types.CommonEntityData {
    mplsMa.EntityData.YFilter = mplsMa.YFilter
    mplsMa.EntityData.YangName = "mpls-ma"
    mplsMa.EntityData.BundleName = "cisco_ios_xr"
    mplsMa.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-io-oper"
    mplsMa.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-io-oper:mpls-ma"
    mplsMa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsMa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsMa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsMa.EntityData.Children = make(map[string]types.YChild)
    mplsMa.EntityData.Children["nodes"] = types.YChild{"Nodes", &mplsMa.Nodes}
    mplsMa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsMa.EntityData)
}

// MplsMa_Nodes
// NODE container class for MPLS IO MA operational
// data
type MplsMa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node MPLS IO MA operational data. The type is slice of
    // MplsMa_Nodes_Node.
    Node []MplsMa_Nodes_Node
}

func (nodes *MplsMa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "mpls-ma"
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

// MplsMa_Nodes_Node
// Per node MPLS IO MA operational data
type MplsMa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // MPLS IO MA Interfaces information .
    Interfaces MplsMa_Nodes_Node_Interfaces
}

func (node *MplsMa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// MplsMa_Nodes_Node_Interfaces
// MPLS IO MA Interfaces information 
type MplsMa_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS IO MA NODE Interface data . The type is slice of
    // MplsMa_Nodes_Node_Interfaces_Interface_.
    Interface_ []MplsMa_Nodes_Node_Interfaces_Interface
}

func (interfaces *MplsMa_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// MplsMa_Nodes_Node_Interfaces_Interface
// MPLS IO MA NODE Interface data 
type MplsMa_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *MplsMa_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", self.Mtu}
    self.EntityData.Leafs["bkp-label-stack-depth"] = types.YLeaf{"BkpLabelStackDepth", self.BkpLabelStackDepth}
    self.EntityData.Leafs["srte-label-stack-depth"] = types.YLeaf{"SrteLabelStackDepth", self.SrteLabelStackDepth}
    self.EntityData.Leafs["pri-label-stack-depth"] = types.YLeaf{"PriLabelStackDepth", self.PriLabelStackDepth}
    return &(self.EntityData)
}

