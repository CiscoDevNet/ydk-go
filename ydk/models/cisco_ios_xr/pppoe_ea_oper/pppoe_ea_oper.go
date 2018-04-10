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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPOE_EA list of nodes.
    Nodes PppoeEa_Nodes
}

func (pppoeEa *PppoeEa) GetEntityData() *types.CommonEntityData {
    pppoeEa.EntityData.YFilter = pppoeEa.YFilter
    pppoeEa.EntityData.YangName = "pppoe-ea"
    pppoeEa.EntityData.BundleName = "cisco_ios_xr"
    pppoeEa.EntityData.ParentYangName = "Cisco-IOS-XR-pppoe-ea-oper"
    pppoeEa.EntityData.SegmentPath = "Cisco-IOS-XR-pppoe-ea-oper:pppoe-ea"
    pppoeEa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoeEa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoeEa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoeEa.EntityData.Children = make(map[string]types.YChild)
    pppoeEa.EntityData.Children["nodes"] = types.YChild{"Nodes", &pppoeEa.Nodes}
    pppoeEa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pppoeEa.EntityData)
}

// PppoeEa_Nodes
// PPPOE_EA list of nodes
type PppoeEa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPOE-EA operational data for a particular node. The type is slice of
    // PppoeEa_Nodes_Node.
    Node []PppoeEa_Nodes_Node
}

func (nodes *PppoeEa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "pppoe-ea"
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

// PppoeEa_Nodes_Node
// PPPOE-EA operational data for a particular node
type PppoeEa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // PPPoE parent interface info.
    ParentInterfaceIds PppoeEa_Nodes_Node_ParentInterfaceIds

    // PPPoE interface info.
    InterfaceIds PppoeEa_Nodes_Node_InterfaceIds
}

func (node *PppoeEa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["parent-interface-ids"] = types.YChild{"ParentInterfaceIds", &node.ParentInterfaceIds}
    node.EntityData.Children["interface-ids"] = types.YChild{"InterfaceIds", &node.InterfaceIds}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// PppoeEa_Nodes_Node_ParentInterfaceIds
// PPPoE parent interface info
type PppoeEa_Nodes_Node_ParentInterfaceIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE parent interface info. The type is slice of
    // PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId.
    ParentInterfaceId []PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId
}

func (parentInterfaceIds *PppoeEa_Nodes_Node_ParentInterfaceIds) GetEntityData() *types.CommonEntityData {
    parentInterfaceIds.EntityData.YFilter = parentInterfaceIds.YFilter
    parentInterfaceIds.EntityData.YangName = "parent-interface-ids"
    parentInterfaceIds.EntityData.BundleName = "cisco_ios_xr"
    parentInterfaceIds.EntityData.ParentYangName = "node"
    parentInterfaceIds.EntityData.SegmentPath = "parent-interface-ids"
    parentInterfaceIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentInterfaceIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentInterfaceIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentInterfaceIds.EntityData.Children = make(map[string]types.YChild)
    parentInterfaceIds.EntityData.Children["parent-interface-id"] = types.YChild{"ParentInterfaceId", nil}
    for i := range parentInterfaceIds.ParentInterfaceId {
        parentInterfaceIds.EntityData.Children[types.GetSegmentPath(&parentInterfaceIds.ParentInterfaceId[i])] = types.YChild{"ParentInterfaceId", &parentInterfaceIds.ParentInterfaceId[i]}
    }
    parentInterfaceIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(parentInterfaceIds.EntityData)
}

// PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId
// PPPoE parent interface info
type PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    ParentInterfaceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Is in sync. The type is bool.
    IsInSync interface{}

    // SRG VMac-address.
    SrgvMac PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac
}

func (parentInterfaceId *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId) GetEntityData() *types.CommonEntityData {
    parentInterfaceId.EntityData.YFilter = parentInterfaceId.YFilter
    parentInterfaceId.EntityData.YangName = "parent-interface-id"
    parentInterfaceId.EntityData.BundleName = "cisco_ios_xr"
    parentInterfaceId.EntityData.ParentYangName = "parent-interface-ids"
    parentInterfaceId.EntityData.SegmentPath = "parent-interface-id" + "[parent-interface-name='" + fmt.Sprintf("%v", parentInterfaceId.ParentInterfaceName) + "']"
    parentInterfaceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentInterfaceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentInterfaceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentInterfaceId.EntityData.Children = make(map[string]types.YChild)
    parentInterfaceId.EntityData.Children["srgv-mac"] = types.YChild{"SrgvMac", &parentInterfaceId.SrgvMac}
    parentInterfaceId.EntityData.Leafs = make(map[string]types.YLeaf)
    parentInterfaceId.EntityData.Leafs["parent-interface-name"] = types.YLeaf{"ParentInterfaceName", parentInterfaceId.ParentInterfaceName}
    parentInterfaceId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", parentInterfaceId.Interface_}
    parentInterfaceId.EntityData.Leafs["is-in-sync"] = types.YLeaf{"IsInSync", parentInterfaceId.IsInSync}
    return &(parentInterfaceId.EntityData)
}

// PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac
// SRG VMac-address
type PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Macaddr interface{}
}

func (srgvMac *PppoeEa_Nodes_Node_ParentInterfaceIds_ParentInterfaceId_SrgvMac) GetEntityData() *types.CommonEntityData {
    srgvMac.EntityData.YFilter = srgvMac.YFilter
    srgvMac.EntityData.YangName = "srgv-mac"
    srgvMac.EntityData.BundleName = "cisco_ios_xr"
    srgvMac.EntityData.ParentYangName = "parent-interface-id"
    srgvMac.EntityData.SegmentPath = "srgv-mac"
    srgvMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgvMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgvMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgvMac.EntityData.Children = make(map[string]types.YChild)
    srgvMac.EntityData.Leafs = make(map[string]types.YLeaf)
    srgvMac.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", srgvMac.Macaddr}
    return &(srgvMac.EntityData)
}

// PppoeEa_Nodes_Node_InterfaceIds
// PPPoE interface info
type PppoeEa_Nodes_Node_InterfaceIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE interface info. The type is slice of
    // PppoeEa_Nodes_Node_InterfaceIds_InterfaceId.
    InterfaceId []PppoeEa_Nodes_Node_InterfaceIds_InterfaceId
}

func (interfaceIds *PppoeEa_Nodes_Node_InterfaceIds) GetEntityData() *types.CommonEntityData {
    interfaceIds.EntityData.YFilter = interfaceIds.YFilter
    interfaceIds.EntityData.YangName = "interface-ids"
    interfaceIds.EntityData.BundleName = "cisco_ios_xr"
    interfaceIds.EntityData.ParentYangName = "node"
    interfaceIds.EntityData.SegmentPath = "interface-ids"
    interfaceIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIds.EntityData.Children = make(map[string]types.YChild)
    interfaceIds.EntityData.Children["interface-id"] = types.YChild{"InterfaceId", nil}
    for i := range interfaceIds.InterfaceId {
        interfaceIds.EntityData.Children[types.GetSegmentPath(&interfaceIds.InterfaceId[i])] = types.YChild{"InterfaceId", &interfaceIds.InterfaceId[i]}
    }
    interfaceIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceIds.EntityData)
}

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId
// PPPoE interface info
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Session ID. The type is interface{} with range: 0..65535.
    SessionId interface{}

    // Parent Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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

func (interfaceId *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId) GetEntityData() *types.CommonEntityData {
    interfaceId.EntityData.YFilter = interfaceId.YFilter
    interfaceId.EntityData.YangName = "interface-id"
    interfaceId.EntityData.BundleName = "cisco_ios_xr"
    interfaceId.EntityData.ParentYangName = "interface-ids"
    interfaceId.EntityData.SegmentPath = "interface-id" + "[interface-name='" + fmt.Sprintf("%v", interfaceId.InterfaceName) + "']"
    interfaceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceId.EntityData.Children = make(map[string]types.YChild)
    interfaceId.EntityData.Children["peer-mac"] = types.YChild{"PeerMac", &interfaceId.PeerMac}
    interfaceId.EntityData.Children["local-mac"] = types.YChild{"LocalMac", &interfaceId.LocalMac}
    interfaceId.EntityData.Children["srgv-mac"] = types.YChild{"SrgvMac", &interfaceId.SrgvMac}
    interfaceId.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceId.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceId.InterfaceName}
    interfaceId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaceId.Interface_}
    interfaceId.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", interfaceId.SessionId}
    interfaceId.EntityData.Leafs["parent-interface"] = types.YLeaf{"ParentInterface", interfaceId.ParentInterface}
    interfaceId.EntityData.Leafs["is-priority-set"] = types.YLeaf{"IsPrioritySet", interfaceId.IsPrioritySet}
    interfaceId.EntityData.Leafs["priority"] = types.YLeaf{"Priority", interfaceId.Priority}
    interfaceId.EntityData.Leafs["is-in-sync"] = types.YLeaf{"IsInSync", interfaceId.IsInSync}
    interfaceId.EntityData.Leafs["is-platform-created"] = types.YLeaf{"IsPlatformCreated", interfaceId.IsPlatformCreated}
    interfaceId.EntityData.Leafs["vlanid"] = types.YLeaf{"Vlanid", interfaceId.Vlanid}
    return &(interfaceId.EntityData)
}

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac
// Peer Mac-address
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Macaddr interface{}
}

func (peerMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_PeerMac) GetEntityData() *types.CommonEntityData {
    peerMac.EntityData.YFilter = peerMac.YFilter
    peerMac.EntityData.YangName = "peer-mac"
    peerMac.EntityData.BundleName = "cisco_ios_xr"
    peerMac.EntityData.ParentYangName = "interface-id"
    peerMac.EntityData.SegmentPath = "peer-mac"
    peerMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerMac.EntityData.Children = make(map[string]types.YChild)
    peerMac.EntityData.Leafs = make(map[string]types.YLeaf)
    peerMac.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", peerMac.Macaddr}
    return &(peerMac.EntityData)
}

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac
// Local Mac-address
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Macaddr interface{}
}

func (localMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_LocalMac) GetEntityData() *types.CommonEntityData {
    localMac.EntityData.YFilter = localMac.YFilter
    localMac.EntityData.YangName = "local-mac"
    localMac.EntityData.BundleName = "cisco_ios_xr"
    localMac.EntityData.ParentYangName = "interface-id"
    localMac.EntityData.SegmentPath = "local-mac"
    localMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localMac.EntityData.Children = make(map[string]types.YChild)
    localMac.EntityData.Leafs = make(map[string]types.YLeaf)
    localMac.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", localMac.Macaddr}
    return &(localMac.EntityData)
}

// PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac
// SRG VMac-address
type PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Macaddr interface{}
}

func (srgvMac *PppoeEa_Nodes_Node_InterfaceIds_InterfaceId_SrgvMac) GetEntityData() *types.CommonEntityData {
    srgvMac.EntityData.YFilter = srgvMac.YFilter
    srgvMac.EntityData.YangName = "srgv-mac"
    srgvMac.EntityData.BundleName = "cisco_ios_xr"
    srgvMac.EntityData.ParentYangName = "interface-id"
    srgvMac.EntityData.SegmentPath = "srgv-mac"
    srgvMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgvMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgvMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgvMac.EntityData.Children = make(map[string]types.YChild)
    srgvMac.EntityData.Leafs = make(map[string]types.YLeaf)
    srgvMac.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", srgvMac.Macaddr}
    return &(srgvMac.EntityData)
}

