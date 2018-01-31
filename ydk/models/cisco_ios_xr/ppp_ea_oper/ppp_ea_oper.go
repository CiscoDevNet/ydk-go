// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ea package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pppea: PPPEA operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node PPPEA operational data.
    Nodes Pppea_Nodes
}

func (pppea *Pppea) GetFilter() yfilter.YFilter { return pppea.YFilter }

func (pppea *Pppea) SetFilter(yf yfilter.YFilter) { pppea.YFilter = yf }

func (pppea *Pppea) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (pppea *Pppea) GetSegmentPath() string {
    return "Cisco-IOS-XR-ppp-ea-oper:pppea"
}

func (pppea *Pppea) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &pppea.Nodes
    }
    return nil
}

func (pppea *Pppea) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &pppea.Nodes
    return children
}

func (pppea *Pppea) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pppea *Pppea) GetBundleName() string { return "cisco_ios_xr" }

func (pppea *Pppea) GetYangName() string { return "pppea" }

func (pppea *Pppea) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppea *Pppea) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppea *Pppea) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppea *Pppea) SetParent(parent types.Entity) { pppea.parent = parent }

func (pppea *Pppea) GetParent() types.Entity { return pppea.parent }

func (pppea *Pppea) GetParentYangName() string { return "Cisco-IOS-XR-ppp-ea-oper" }

// Pppea_Nodes
// Per node PPPEA operational data
type Pppea_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The PPPEA operational data for a particular node. The type is slice of
    // Pppea_Nodes_Node.
    Node []Pppea_Nodes_Node
}

func (nodes *Pppea_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Pppea_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Pppea_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Pppea_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Pppea_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppea_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Pppea_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Pppea_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Pppea_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Pppea_Nodes) GetYangName() string { return "nodes" }

func (nodes *Pppea_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Pppea_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Pppea_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Pppea_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Pppea_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Pppea_Nodes) GetParentYangName() string { return "pppea" }

// Pppea_Nodes_Node
// The PPPEA operational data for a particular
// node
type Pppea_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Show interface related information from the PPP EA.
    EaInterfaceNames Pppea_Nodes_Node_EaInterfaceNames
}

func (node *Pppea_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Pppea_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Pppea_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ea-interface-names" { return "EaInterfaceNames" }
    return ""
}

func (node *Pppea_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Pppea_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ea-interface-names" {
        return &node.EaInterfaceNames
    }
    return nil
}

func (node *Pppea_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ea-interface-names"] = &node.EaInterfaceNames
    return children
}

func (node *Pppea_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Pppea_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Pppea_Nodes_Node) GetYangName() string { return "node" }

func (node *Pppea_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Pppea_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Pppea_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Pppea_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Pppea_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Pppea_Nodes_Node) GetParentYangName() string { return "nodes" }

// Pppea_Nodes_Node_EaInterfaceNames
// Show interface related information from the
// PPP EA
type Pppea_Nodes_Node_EaInterfaceNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName.
    EaInterfaceName []Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetFilter() yfilter.YFilter { return eaInterfaceNames.YFilter }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) SetFilter(yf yfilter.YFilter) { eaInterfaceNames.YFilter = yf }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetGoName(yname string) string {
    if yname == "ea-interface-name" { return "EaInterfaceName" }
    return ""
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetSegmentPath() string {
    return "ea-interface-names"
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ea-interface-name" {
        for _, c := range eaInterfaceNames.EaInterfaceName {
            if eaInterfaceNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName{}
        eaInterfaceNames.EaInterfaceName = append(eaInterfaceNames.EaInterfaceName, child)
        return &eaInterfaceNames.EaInterfaceName[len(eaInterfaceNames.EaInterfaceName)-1]
    }
    return nil
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eaInterfaceNames.EaInterfaceName {
        children[eaInterfaceNames.EaInterfaceName[i].GetSegmentPath()] = &eaInterfaceNames.EaInterfaceName[i]
    }
    return children
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetBundleName() string { return "cisco_ios_xr" }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetYangName() string { return "ea-interface-names" }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) SetParent(parent types.Entity) { eaInterfaceNames.parent = parent }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetParent() types.Entity { return eaInterfaceNames.parent }

func (eaInterfaceNames *Pppea_Nodes_Node_EaInterfaceNames) GetParentYangName() string { return "node" }

// Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName
// Interface name
type Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface running PPPEA. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // TRUE if LCP is running in the dataplane for the interface. The type is
    // bool.
    IsLcpRunning interface{}

    // TRUE if IPCP is running in the dataplane for the interface. The type is
    // bool.
    IsIpcpRunning interface{}

    // TRUE if IPV6CP is running in the dataplane for the interface. The type is
    // bool.
    IsIpv6CpRunning interface{}

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
    // string with pattern: [a-zA-Z0-9./-]+.
    MultilinkInterface interface{}

    // L2 Tunnel State. The type is bool.
    L2TunnelEnabled interface{}

    // L2 Provisioned State. The type is bool.
    L2Provisioned interface{}

    // L2 IP Interworking State. The type is bool.
    L2IpInterworkingEnabled interface{}

    // Is VPDN tunneled. The type is bool.
    IsVpdnTunneled interface{}

    // XConnect ID. The type is interface{} with range: 0..4294967295.
    XconnectId interface{}

    // Parent Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterfaceHandle interface{}

    // IPCP VRF Table ID. The type is interface{} with range: 0..4294967295.
    VrfTableId interface{}

    // IPv6CP VRF Table ID. The type is interface{} with range: 0..4294967295.
    Ipv6VrfTableId interface{}

    // L2 adjacency state. The type is PppEaAdjState.
    L2AdjacencyState interface{}

    // L2 IP Interworking adjacency state. The type is PppEaAdjState.
    L2IpInterworkingAdjacencyState interface{}

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

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetFilter() yfilter.YFilter { return eaInterfaceName.YFilter }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) SetFilter(yf yfilter.YFilter) { eaInterfaceName.YFilter = yf }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "is-lcp-running" { return "IsLcpRunning" }
    if yname == "is-ipcp-running" { return "IsIpcpRunning" }
    if yname == "is-ipv6cp-running" { return "IsIpv6CpRunning" }
    if yname == "is-mplscp-running" { return "IsMplscpRunning" }
    if yname == "local-mtu" { return "LocalMtu" }
    if yname == "local-mrru" { return "LocalMrru" }
    if yname == "peer-mrru" { return "PeerMrru" }
    if yname == "local-magic" { return "LocalMagic" }
    if yname == "peer-magic" { return "PeerMagic" }
    if yname == "local-mcmp-classes" { return "LocalMcmpClasses" }
    if yname == "peer-mcmp-classes" { return "PeerMcmpClasses" }
    if yname == "echo-request-interval" { return "EchoRequestInterval" }
    if yname == "echo-request-retry-count" { return "EchoRequestRetryCount" }
    if yname == "is-multilink-bundle" { return "IsMultilinkBundle" }
    if yname == "synchronized" { return "Synchronized" }
    if yname == "forwarding-enabled" { return "ForwardingEnabled" }
    if yname == "multilink-interface" { return "MultilinkInterface" }
    if yname == "l2-tunnel-enabled" { return "L2TunnelEnabled" }
    if yname == "l2-provisioned" { return "L2Provisioned" }
    if yname == "l2ip-interworking-enabled" { return "L2IpInterworkingEnabled" }
    if yname == "is-vpdn-tunneled" { return "IsVpdnTunneled" }
    if yname == "xconnect-id" { return "XconnectId" }
    if yname == "parent-interface-handle" { return "ParentInterfaceHandle" }
    if yname == "vrf-table-id" { return "VrfTableId" }
    if yname == "ipv6vrf-table-id" { return "Ipv6VrfTableId" }
    if yname == "l2-adjacency-state" { return "L2AdjacencyState" }
    if yname == "l2ip-interworking-adjacency-state" { return "L2IpInterworkingAdjacencyState" }
    if yname == "lac-adjacency-state" { return "LacAdjacencyState" }
    if yname == "interface-adjacency-state" { return "InterfaceAdjacencyState" }
    if yname == "ipv4-adjacency-state" { return "Ipv4AdjacencyState" }
    if yname == "ipv6-adjacency-state" { return "Ipv6AdjacencyState" }
    if yname == "mpls-adjacency-state" { return "MplsAdjacencyState" }
    return ""
}

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetSegmentPath() string {
    return "ea-interface-name" + "[interface-name='" + fmt.Sprintf("%v", eaInterfaceName.InterfaceName) + "']"
}

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = eaInterfaceName.InterfaceName
    leafs["interface"] = eaInterfaceName.Interface
    leafs["is-lcp-running"] = eaInterfaceName.IsLcpRunning
    leafs["is-ipcp-running"] = eaInterfaceName.IsIpcpRunning
    leafs["is-ipv6cp-running"] = eaInterfaceName.IsIpv6CpRunning
    leafs["is-mplscp-running"] = eaInterfaceName.IsMplscpRunning
    leafs["local-mtu"] = eaInterfaceName.LocalMtu
    leafs["local-mrru"] = eaInterfaceName.LocalMrru
    leafs["peer-mrru"] = eaInterfaceName.PeerMrru
    leafs["local-magic"] = eaInterfaceName.LocalMagic
    leafs["peer-magic"] = eaInterfaceName.PeerMagic
    leafs["local-mcmp-classes"] = eaInterfaceName.LocalMcmpClasses
    leafs["peer-mcmp-classes"] = eaInterfaceName.PeerMcmpClasses
    leafs["echo-request-interval"] = eaInterfaceName.EchoRequestInterval
    leafs["echo-request-retry-count"] = eaInterfaceName.EchoRequestRetryCount
    leafs["is-multilink-bundle"] = eaInterfaceName.IsMultilinkBundle
    leafs["synchronized"] = eaInterfaceName.Synchronized
    leafs["forwarding-enabled"] = eaInterfaceName.ForwardingEnabled
    leafs["multilink-interface"] = eaInterfaceName.MultilinkInterface
    leafs["l2-tunnel-enabled"] = eaInterfaceName.L2TunnelEnabled
    leafs["l2-provisioned"] = eaInterfaceName.L2Provisioned
    leafs["l2ip-interworking-enabled"] = eaInterfaceName.L2IpInterworkingEnabled
    leafs["is-vpdn-tunneled"] = eaInterfaceName.IsVpdnTunneled
    leafs["xconnect-id"] = eaInterfaceName.XconnectId
    leafs["parent-interface-handle"] = eaInterfaceName.ParentInterfaceHandle
    leafs["vrf-table-id"] = eaInterfaceName.VrfTableId
    leafs["ipv6vrf-table-id"] = eaInterfaceName.Ipv6VrfTableId
    leafs["l2-adjacency-state"] = eaInterfaceName.L2AdjacencyState
    leafs["l2ip-interworking-adjacency-state"] = eaInterfaceName.L2IpInterworkingAdjacencyState
    leafs["lac-adjacency-state"] = eaInterfaceName.LacAdjacencyState
    leafs["interface-adjacency-state"] = eaInterfaceName.InterfaceAdjacencyState
    leafs["ipv4-adjacency-state"] = eaInterfaceName.Ipv4AdjacencyState
    leafs["ipv6-adjacency-state"] = eaInterfaceName.Ipv6AdjacencyState
    leafs["mpls-adjacency-state"] = eaInterfaceName.MplsAdjacencyState
    return leafs
}

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetBundleName() string { return "cisco_ios_xr" }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetYangName() string { return "ea-interface-name" }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) SetParent(parent types.Entity) { eaInterfaceName.parent = parent }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetParent() types.Entity { return eaInterfaceName.parent }

func (eaInterfaceName *Pppea_Nodes_Node_EaInterfaceNames_EaInterfaceName) GetParentYangName() string { return "ea-interface-names" }

