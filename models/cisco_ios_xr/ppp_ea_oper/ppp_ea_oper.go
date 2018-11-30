// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ea package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pppea: PPPEA operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ea_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ea_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ppp-ea-oper pppea}", reflect.TypeOf(Pppea{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ppp-ea-oper:pppea", reflect.TypeOf(Pppea{}))
}

// PppEaAdjState represents Ppp ea adj state
type PppEaAdjState string

const (
    // Ajacency not installed in AIB
    PppEaAdjState_ppp_ea_adj_state_not_installed PppEaAdjState = "ppp-ea-adj-state-not-installed"

    // Adjacency installed in AIB
    PppEaAdjState_ppp_ea_adj_state_installed PppEaAdjState = "ppp-ea-adj-state-installed"
)

// Pppea
// PPPEA operational data
type Pppea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node PPPEA operational data.
    Nodes Pppea_Nodes
}

func (pppea *Pppea) GetEntityData() *types.CommonEntityData {
    pppea.EntityData.YFilter = pppea.YFilter
    pppea.EntityData.YangName = "pppea"
    pppea.EntityData.BundleName = "cisco_ios_xr"
    pppea.EntityData.ParentYangName = "Cisco-IOS-XR-ppp-ea-oper"
    pppea.EntityData.SegmentPath = "Cisco-IOS-XR-ppp-ea-oper:pppea"
    pppea.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppea.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppea.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppea.EntityData.Children = types.NewOrderedMap()
    pppea.EntityData.Children.Append("nodes", types.YChild{"Nodes", &pppea.Nodes})
    pppea.EntityData.Leafs = types.NewOrderedMap()

    pppea.EntityData.YListKeys = []string {}

    return &(pppea.EntityData)
}

// Pppea_Nodes
// Per node PPPEA operational data
type Pppea_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The PPPEA operational data for a particular node. The type is slice of
    // Pppea_Nodes_Node.
    Node []*Pppea_Nodes_Node
}

func (nodes *Pppea_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "pppea"
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

// Pppea_Nodes_Node
// The PPPEA operational data for a particular
// node
type Pppea_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Show interface related information from the PPP EA.
    EaInterfaceNames Pppea_Nodes_Node_EaInterfaceNames
}

func (node *Pppea_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("ea-interface-names", types.YChild{"EaInterfaceNames", &node.EaInterfaceNames})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Pppea_Nodes_Node_EaInterfaceNames
// Show interface related information from the
// PPP EA
type Pppea_Nodes_Node_EaInterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName.
    EaInterfaceName []*Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetEntityData() *types.CommonEntityData {
    eaInterfaceNames.EntityData.YFilter = eaInterfaceNames.YFilter
    eaInterfaceNames.EntityData.YangName = "ea-interface-names"
    eaInterfaceNames.EntityData.BundleName = "cisco_ios_xr"
    eaInterfaceNames.EntityData.ParentYangName = "node"
    eaInterfaceNames.EntityData.SegmentPath = "ea-interface-names"
    eaInterfaceNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eaInterfaceNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eaInterfaceNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eaInterfaceNames.EntityData.Children = types.NewOrderedMap()
    eaInterfaceNames.EntityData.Children.Append("ea-interface-name", types.YChild{"EaInterfaceName", nil})
    for i := range eaInterfaceNames.EaInterfaceName {
        eaInterfaceNames.EntityData.Children.Append(types.GetSegmentPath(eaInterfaceNames.EaInterfaceName[i]), types.YChild{"EaInterfaceName", eaInterfaceNames.EaInterfaceName[i]})
    }
    eaInterfaceNames.EntityData.Leafs = types.NewOrderedMap()

    eaInterfaceNames.EntityData.YListKeys = []string {}

    return &(eaInterfaceNames.EntityData)
}

// Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName
// Interface name
type Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface running PPPEA. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // TRUE if LCP is running in the dataplane for the interface. The type is
    // bool.
    IsLcpRunning interface{}

    // TRUE if IPCP is running in the dataplane for the interface. The type is
    // bool.
    IsIpcpRunning interface{}

    // TRUE if IPV6CP is running in the dataplane for the interface. The type is
    // bool.
    IsIpv6cpRunning interface{}

    // TRUE if MPLSCP is running in the dataplane for the interface. The type is
    // bool.
    IsMplscpRunning interface{}

    // Local interface MTU. The type is interface{} with range: 0..65535.
    LocalMtu interface{}

    // Local MRRU. The type is interface{} with range: 0..65535.
    LocalMrru interface{}

    // Peer MRRU. The type is interface{} with range: 0..65535.
    PeerMrru interface{}

    // Local magic number. The type is interface{} with range: 0..4294967295.
    LocalMagic interface{}

    // Peer magic number. The type is interface{} with range: 0..4294967295.
    PeerMagic interface{}

    // Local number of MCMP Suspension classes. The type is interface{} with
    // range: 0..255.
    LocalMcmpClasses interface{}

    // Peer number of MCMP Suspension classes. The type is interface{} with range:
    // 0..255.
    PeerMcmpClasses interface{}

    // Echo-Request interval. The type is interface{} with range: 0..4294967295.
    EchoRequestInterval interface{}

    // Echo-Request retry count. The type is interface{} with range:
    // 0..4294967295.
    EchoRequestRetryCount interface{}

    // TRUE if this is a Multilink bundle interface. The type is bool.
    IsMultilinkBundle interface{}

    // MA synchronization. The type is bool.
    Synchronized interface{}

    // Forwarding State. The type is bool.
    ForwardingEnabled interface{}

    // Multilink interface that this interface is a member of, if any. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    MultilinkInterface interface{}

    // L2 Tunnel State. The type is bool.
    L2TunnelEnabled interface{}

    // L2 Provisioned State. The type is bool.
    L2Provisioned interface{}

    // L2 IP Interworking State. The type is bool.
    L2ipInterworkingEnabled interface{}

    // Is VPDN tunneled. The type is bool.
    IsVpdnTunneled interface{}

    // XConnect ID. The type is interface{} with range: 0..4294967295.
    XconnectId interface{}

    // Parent Interface Handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    ParentInterfaceHandle interface{}

    // IPCP VRF Table ID. The type is interface{} with range: 0..4294967295.
    VrfTableId interface{}

    // IPv6CP VRF Table ID. The type is interface{} with range: 0..4294967295.
    Ipv6vrfTableId interface{}

    // L2 adjacency state. The type is PppEaAdjState.
    L2AdjacencyState interface{}

    // L2 IP Interworking adjacency state. The type is PppEaAdjState.
    L2ipInterworkingAdjacencyState interface{}

    // LAC adjacency state. The type is PppEaAdjState.
    LacAdjacencyState interface{}

    // Interface adjacency state. The type is PppEaAdjState.
    InterfaceAdjacencyState interface{}

    // Ipv4 adjacency state. The type is PppEaAdjState.
    Ipv4AdjacencyState interface{}

    // IPv6 adjacency state. The type is PppEaAdjState.
    Ipv6AdjacencyState interface{}

    // MPLS adjacency state. The type is PppEaAdjState.
    MplsAdjacencyState interface{}
}

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetEntityData() *types.CommonEntityData {
    eaInterfaceName.EntityData.YFilter = eaInterfaceName.YFilter
    eaInterfaceName.EntityData.YangName = "ea-interface-name"
    eaInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    eaInterfaceName.EntityData.ParentYangName = "ea-interface-names"
    eaInterfaceName.EntityData.SegmentPath = "ea-interface-name" + types.AddKeyToken(eaInterfaceName.InterfaceName, "interface-name")
    eaInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eaInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eaInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eaInterfaceName.EntityData.Children = types.NewOrderedMap()
    eaInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    eaInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", eaInterfaceName.InterfaceName})
    eaInterfaceName.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", eaInterfaceName.Interface})
    eaInterfaceName.EntityData.Leafs.Append("is-lcp-running", types.YLeaf{"IsLcpRunning", eaInterfaceName.IsLcpRunning})
    eaInterfaceName.EntityData.Leafs.Append("is-ipcp-running", types.YLeaf{"IsIpcpRunning", eaInterfaceName.IsIpcpRunning})
    eaInterfaceName.EntityData.Leafs.Append("is-ipv6cp-running", types.YLeaf{"IsIpv6cpRunning", eaInterfaceName.IsIpv6cpRunning})
    eaInterfaceName.EntityData.Leafs.Append("is-mplscp-running", types.YLeaf{"IsMplscpRunning", eaInterfaceName.IsMplscpRunning})
    eaInterfaceName.EntityData.Leafs.Append("local-mtu", types.YLeaf{"LocalMtu", eaInterfaceName.LocalMtu})
    eaInterfaceName.EntityData.Leafs.Append("local-mrru", types.YLeaf{"LocalMrru", eaInterfaceName.LocalMrru})
    eaInterfaceName.EntityData.Leafs.Append("peer-mrru", types.YLeaf{"PeerMrru", eaInterfaceName.PeerMrru})
    eaInterfaceName.EntityData.Leafs.Append("local-magic", types.YLeaf{"LocalMagic", eaInterfaceName.LocalMagic})
    eaInterfaceName.EntityData.Leafs.Append("peer-magic", types.YLeaf{"PeerMagic", eaInterfaceName.PeerMagic})
    eaInterfaceName.EntityData.Leafs.Append("local-mcmp-classes", types.YLeaf{"LocalMcmpClasses", eaInterfaceName.LocalMcmpClasses})
    eaInterfaceName.EntityData.Leafs.Append("peer-mcmp-classes", types.YLeaf{"PeerMcmpClasses", eaInterfaceName.PeerMcmpClasses})
    eaInterfaceName.EntityData.Leafs.Append("echo-request-interval", types.YLeaf{"EchoRequestInterval", eaInterfaceName.EchoRequestInterval})
    eaInterfaceName.EntityData.Leafs.Append("echo-request-retry-count", types.YLeaf{"EchoRequestRetryCount", eaInterfaceName.EchoRequestRetryCount})
    eaInterfaceName.EntityData.Leafs.Append("is-multilink-bundle", types.YLeaf{"IsMultilinkBundle", eaInterfaceName.IsMultilinkBundle})
    eaInterfaceName.EntityData.Leafs.Append("synchronized", types.YLeaf{"Synchronized", eaInterfaceName.Synchronized})
    eaInterfaceName.EntityData.Leafs.Append("forwarding-enabled", types.YLeaf{"ForwardingEnabled", eaInterfaceName.ForwardingEnabled})
    eaInterfaceName.EntityData.Leafs.Append("multilink-interface", types.YLeaf{"MultilinkInterface", eaInterfaceName.MultilinkInterface})
    eaInterfaceName.EntityData.Leafs.Append("l2-tunnel-enabled", types.YLeaf{"L2TunnelEnabled", eaInterfaceName.L2TunnelEnabled})
    eaInterfaceName.EntityData.Leafs.Append("l2-provisioned", types.YLeaf{"L2Provisioned", eaInterfaceName.L2Provisioned})
    eaInterfaceName.EntityData.Leafs.Append("l2ip-interworking-enabled", types.YLeaf{"L2ipInterworkingEnabled", eaInterfaceName.L2ipInterworkingEnabled})
    eaInterfaceName.EntityData.Leafs.Append("is-vpdn-tunneled", types.YLeaf{"IsVpdnTunneled", eaInterfaceName.IsVpdnTunneled})
    eaInterfaceName.EntityData.Leafs.Append("xconnect-id", types.YLeaf{"XconnectId", eaInterfaceName.XconnectId})
    eaInterfaceName.EntityData.Leafs.Append("parent-interface-handle", types.YLeaf{"ParentInterfaceHandle", eaInterfaceName.ParentInterfaceHandle})
    eaInterfaceName.EntityData.Leafs.Append("vrf-table-id", types.YLeaf{"VrfTableId", eaInterfaceName.VrfTableId})
    eaInterfaceName.EntityData.Leafs.Append("ipv6vrf-table-id", types.YLeaf{"Ipv6vrfTableId", eaInterfaceName.Ipv6vrfTableId})
    eaInterfaceName.EntityData.Leafs.Append("l2-adjacency-state", types.YLeaf{"L2AdjacencyState", eaInterfaceName.L2AdjacencyState})
    eaInterfaceName.EntityData.Leafs.Append("l2ip-interworking-adjacency-state", types.YLeaf{"L2ipInterworkingAdjacencyState", eaInterfaceName.L2ipInterworkingAdjacencyState})
    eaInterfaceName.EntityData.Leafs.Append("lac-adjacency-state", types.YLeaf{"LacAdjacencyState", eaInterfaceName.LacAdjacencyState})
    eaInterfaceName.EntityData.Leafs.Append("interface-adjacency-state", types.YLeaf{"InterfaceAdjacencyState", eaInterfaceName.InterfaceAdjacencyState})
    eaInterfaceName.EntityData.Leafs.Append("ipv4-adjacency-state", types.YLeaf{"Ipv4AdjacencyState", eaInterfaceName.Ipv4AdjacencyState})
    eaInterfaceName.EntityData.Leafs.Append("ipv6-adjacency-state", types.YLeaf{"Ipv6AdjacencyState", eaInterfaceName.Ipv6AdjacencyState})
    eaInterfaceName.EntityData.Leafs.Append("mpls-adjacency-state", types.YLeaf{"MplsAdjacencyState", eaInterfaceName.MplsAdjacencyState})

    eaInterfaceName.EntityData.YListKeys = []string {"InterfaceName"}

    return &(eaInterfaceName.EntityData)
}

