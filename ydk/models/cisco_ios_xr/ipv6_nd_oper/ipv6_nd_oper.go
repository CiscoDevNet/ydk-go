// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-nd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-node-discovery: IPv6 node discovery operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_nd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_nd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-nd-oper ipv6-node-discovery}", reflect.TypeOf(Ipv6NodeDiscovery{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery", reflect.TypeOf(Ipv6NodeDiscovery{}))
}

// Ipv6NdShVrState represents IPv6 ND VR Entry State Type 
type Ipv6NdShVrState string

const (
    // Delete
    Ipv6NdShVrState_deleted Ipv6NdShVrState = "deleted"

    // Standby
    Ipv6NdShVrState_standby Ipv6NdShVrState = "standby"

    // Active
    Ipv6NdShVrState_active Ipv6NdShVrState = "active"
)

// Ipv6NdBndlState represents IPv6 ND Bundle State
type Ipv6NdBndlState string

const (
    // Running state
    Ipv6NdBndlState_run Ipv6NdBndlState = "run"

    // Error state
    Ipv6NdBndlState_error Ipv6NdBndlState = "error"

    // Wait state
    Ipv6NdBndlState_wait Ipv6NdBndlState = "wait"
)

// Ipv6NdMediaEncap represents IPv6 ND Media Encapsulation Type
type Ipv6NdMediaEncap string

const (
    // No encapsulation
    Ipv6NdMediaEncap_none Ipv6NdMediaEncap = "none"

    // ARPA encapsulation
    Ipv6NdMediaEncap_arpa Ipv6NdMediaEncap = "arpa"

    // SNAP encapsulation
    Ipv6NdMediaEncap_snap Ipv6NdMediaEncap = "snap"

    // 802_1Q encapsulation
    Ipv6NdMediaEncap_ieee802_1q Ipv6NdMediaEncap = "ieee802-1q"

    // SRP encapsulation
    Ipv6NdMediaEncap_srp Ipv6NdMediaEncap = "srp"

    // SRPA encapsulation
    Ipv6NdMediaEncap_srpa Ipv6NdMediaEncap = "srpa"

    // SRPB encapsulation
    Ipv6NdMediaEncap_srpb Ipv6NdMediaEncap = "srpb"

    // PPP encapsulation
    Ipv6NdMediaEncap_ppp Ipv6NdMediaEncap = "ppp"

    // HDLC encapsulation
    Ipv6NdMediaEncap_hdlc Ipv6NdMediaEncap = "hdlc"

    // CHDLC encapsulation
    Ipv6NdMediaEncap_chdlc Ipv6NdMediaEncap = "chdlc"

    // DOT1Q encapsulation
    Ipv6NdMediaEncap_dot1q Ipv6NdMediaEncap = "dot1q"

    // FR encapsulation
    Ipv6NdMediaEncap_fr Ipv6NdMediaEncap = "fr"

    // GRE encapsulation
    Ipv6NdMediaEncap_gre Ipv6NdMediaEncap = "gre"
)

// Ipv6NdNeighborOrigin represents IPv6 ND Neighbor Origin Type
type Ipv6NdNeighborOrigin string

const (
    // Other Address
    Ipv6NdNeighborOrigin_other Ipv6NdNeighborOrigin = "other"

    // Static Address
    Ipv6NdNeighborOrigin_static Ipv6NdNeighborOrigin = "static"

    // Dynamic Address
    Ipv6NdNeighborOrigin_dynamic Ipv6NdNeighborOrigin = "dynamic"
)

// Ipv6NdShState represents IPv6 ND Neighbor Reachability State
type Ipv6NdShState string

const (
    // Incomplete
    Ipv6NdShState_incomplete Ipv6NdShState = "incomplete"

    // Reachable
    Ipv6NdShState_reachable Ipv6NdShState = "reachable"

    // Stale
    Ipv6NdShState_stale Ipv6NdShState = "stale"

    // Glean
    Ipv6NdShState_glean Ipv6NdShState = "glean"

    // Delay
    Ipv6NdShState_delay Ipv6NdShState = "delay"

    // Probe
    Ipv6NdShState_probe Ipv6NdShState = "probe"

    // Delete
    Ipv6NdShState_delete Ipv6NdShState = "delete"
)

// Ipv6NdShVrFlags represents IPv6 ND VR Entry Flags Type 
type Ipv6NdShVrFlags string

const (
    // None
    Ipv6NdShVrFlags_no_flags Ipv6NdShVrFlags = "no-flags"

    // Final RA
    Ipv6NdShVrFlags_final_ra Ipv6NdShVrFlags = "final-ra"
)

// Ipv6NodeDiscovery
// IPv6 node discovery operational data
type Ipv6NodeDiscovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 node discovery list of nodes.
    Nodes Ipv6NodeDiscovery_Nodes
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetFilter() yfilter.YFilter { return ipv6NodeDiscovery.YFilter }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) SetFilter(yf yfilter.YFilter) { ipv6NodeDiscovery.YFilter = yf }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery"
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ipv6NodeDiscovery.Nodes
    }
    return nil
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ipv6NodeDiscovery.Nodes
    return children
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetYangName() string { return "ipv6-node-discovery" }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) SetParent(parent types.Entity) { ipv6NodeDiscovery.parent = parent }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetParent() types.Entity { return ipv6NodeDiscovery.parent }

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-nd-oper" }

// Ipv6NodeDiscovery_Nodes
// IPv6 node discovery list of nodes
type Ipv6NodeDiscovery_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 node discovery operational data for a particular node. The type is
    // slice of Ipv6NodeDiscovery_Nodes_Node.
    Node []Ipv6NodeDiscovery_Nodes_Node
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ipv6NodeDiscovery_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ipv6NodeDiscovery_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ipv6NodeDiscovery_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ipv6NodeDiscovery_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ipv6NodeDiscovery_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ipv6NodeDiscovery_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ipv6NodeDiscovery_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ipv6NodeDiscovery_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ipv6NodeDiscovery_Nodes) GetParentYangName() string { return "ipv6-node-discovery" }

// Ipv6NodeDiscovery_Nodes_Node
// IPv6 node discovery operational data for a
// particular node
type Ipv6NodeDiscovery_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // IPv6 node discovery list of neighbor interfaces.
    NeighborInterfaces Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces

    // IPv6 Neighbor summary.
    NeighborSummary Ipv6NodeDiscovery_Nodes_Node_NeighborSummary

    // IPv6 ND list of bundle nodes for a specific node.
    BundleNodes Ipv6NodeDiscovery_Nodes_Node_BundleNodes

    // IPv6 ND list of bundle interfaces for a specific node.
    BundleInterfaces Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces

    // IPv6 node discovery list of interfaces for a specific node.
    Interfaces Ipv6NodeDiscovery_Nodes_Node_Interfaces

    // IPv6 ND virtual router information for a specific interface.
    NdVirtualRouters Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters

    // IPv6 ND list of SLAAC MGMT interfaces for a specific node.
    SlaacInterfaces Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces
}

func (node *Ipv6NodeDiscovery_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ipv6NodeDiscovery_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "neighbor-interfaces" { return "NeighborInterfaces" }
    if yname == "neighbor-summary" { return "NeighborSummary" }
    if yname == "bundle-nodes" { return "BundleNodes" }
    if yname == "bundle-interfaces" { return "BundleInterfaces" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "nd-virtual-routers" { return "NdVirtualRouters" }
    if yname == "slaac-interfaces" { return "SlaacInterfaces" }
    return ""
}

func (node *Ipv6NodeDiscovery_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ipv6NodeDiscovery_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor-interfaces" {
        return &node.NeighborInterfaces
    }
    if childYangName == "neighbor-summary" {
        return &node.NeighborSummary
    }
    if childYangName == "bundle-nodes" {
        return &node.BundleNodes
    }
    if childYangName == "bundle-interfaces" {
        return &node.BundleInterfaces
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "nd-virtual-routers" {
        return &node.NdVirtualRouters
    }
    if childYangName == "slaac-interfaces" {
        return &node.SlaacInterfaces
    }
    return nil
}

func (node *Ipv6NodeDiscovery_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbor-interfaces"] = &node.NeighborInterfaces
    children["neighbor-summary"] = &node.NeighborSummary
    children["bundle-nodes"] = &node.BundleNodes
    children["bundle-interfaces"] = &node.BundleInterfaces
    children["interfaces"] = &node.Interfaces
    children["nd-virtual-routers"] = &node.NdVirtualRouters
    children["slaac-interfaces"] = &node.SlaacInterfaces
    return children
}

func (node *Ipv6NodeDiscovery_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ipv6NodeDiscovery_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetYangName() string { return "node" }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ipv6NodeDiscovery_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ipv6NodeDiscovery_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces
// IPv6 node discovery list of neighbor
// interfaces
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 node discovery neighbor interface. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface.
    NeighborInterface []Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetFilter() yfilter.YFilter { return neighborInterfaces.YFilter }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) SetFilter(yf yfilter.YFilter) { neighborInterfaces.YFilter = yf }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetGoName(yname string) string {
    if yname == "neighbor-interface" { return "NeighborInterface" }
    return ""
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetSegmentPath() string {
    return "neighbor-interfaces"
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor-interface" {
        for _, c := range neighborInterfaces.NeighborInterface {
            if neighborInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface{}
        neighborInterfaces.NeighborInterface = append(neighborInterfaces.NeighborInterface, child)
        return &neighborInterfaces.NeighborInterface[len(neighborInterfaces.NeighborInterface)-1]
    }
    return nil
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighborInterfaces.NeighborInterface {
        children[neighborInterfaces.NeighborInterface[i].GetSegmentPath()] = &neighborInterfaces.NeighborInterface[i]
    }
    return children
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetYangName() string { return "neighbor-interfaces" }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) SetParent(parent types.Entity) { neighborInterfaces.parent = parent }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetParent() types.Entity { return neighborInterfaces.parent }

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface
// IPv6 node discovery neighbor interface
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IPv6 node discovery list of neighbor host addresses.
    HostAddresses Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetFilter() yfilter.YFilter { return neighborInterface.YFilter }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) SetFilter(yf yfilter.YFilter) { neighborInterface.YFilter = yf }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "host-addresses" { return "HostAddresses" }
    return ""
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetSegmentPath() string {
    return "neighbor-interface" + "[interface-name='" + fmt.Sprintf("%v", neighborInterface.InterfaceName) + "']"
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-addresses" {
        return &neighborInterface.HostAddresses
    }
    return nil
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["host-addresses"] = &neighborInterface.HostAddresses
    return children
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = neighborInterface.InterfaceName
    return leafs
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetBundleName() string { return "cisco_ios_xr" }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetYangName() string { return "neighbor-interface" }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) SetParent(parent types.Entity) { neighborInterface.parent = parent }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetParent() types.Entity { return neighborInterface.parent }

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetParentYangName() string { return "neighbor-interfaces" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses
// IPv6 node discovery list of neighbor host
// addresses
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Neighbor detailed information. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress.
    HostAddress []Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetFilter() yfilter.YFilter { return hostAddresses.YFilter }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) SetFilter(yf yfilter.YFilter) { hostAddresses.YFilter = yf }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetGoName(yname string) string {
    if yname == "host-address" { return "HostAddress" }
    return ""
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetSegmentPath() string {
    return "host-addresses"
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-address" {
        for _, c := range hostAddresses.HostAddress {
            if hostAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress{}
        hostAddresses.HostAddress = append(hostAddresses.HostAddress, child)
        return &hostAddresses.HostAddress[len(hostAddresses.HostAddress)-1]
    }
    return nil
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hostAddresses.HostAddress {
        children[hostAddresses.HostAddress[i].GetSegmentPath()] = &hostAddresses.HostAddress[i]
    }
    return children
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetYangName() string { return "host-addresses" }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) SetParent(parent types.Entity) { hostAddresses.parent = parent }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetParent() types.Entity { return hostAddresses.parent }

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetParentYangName() string { return "neighbor-interface" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress
// IPv6 Neighbor detailed information
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Host Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddress interface{}

    // Current state. The type is Ipv6NdShState.
    ReachabilityState interface{}

    // Link-Layer Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    LinkLayerAddress interface{}

    // Preferred media encap type. The type is Ipv6NdMediaEncap.
    Encapsulation interface{}

    // Selected media encap. The type is Ipv6NdMediaEncap.
    SelectedEncapsulation interface{}

    // Neighbor origin. The type is Ipv6NdNeighborOrigin.
    OriginEncapsulation interface{}

    // Interface name. The type is string.
    InterfaceName interface{}

    // Location where the neighbor entry exists. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // IsRouter. The type is bool.
    IsRouter interface{}

    // ND serg flags for this entry. The type is interface{} with range:
    // 0..4294967295.
    SergFlags interface{}

    // Last time of reachability.
    LastReachedTime Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetFilter() yfilter.YFilter { return hostAddress.YFilter }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) SetFilter(yf yfilter.YFilter) { hostAddress.YFilter = yf }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetGoName(yname string) string {
    if yname == "host-address" { return "HostAddress" }
    if yname == "reachability-state" { return "ReachabilityState" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "selected-encapsulation" { return "SelectedEncapsulation" }
    if yname == "origin-encapsulation" { return "OriginEncapsulation" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "location" { return "Location" }
    if yname == "is-router" { return "IsRouter" }
    if yname == "serg-flags" { return "SergFlags" }
    if yname == "last-reached-time" { return "LastReachedTime" }
    return ""
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetSegmentPath() string {
    return "host-address" + "[host-address='" + fmt.Sprintf("%v", hostAddress.HostAddress) + "']"
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-reached-time" {
        return &hostAddress.LastReachedTime
    }
    return nil
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-reached-time"] = &hostAddress.LastReachedTime
    return children
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-address"] = hostAddress.HostAddress
    leafs["reachability-state"] = hostAddress.ReachabilityState
    leafs["link-layer-address"] = hostAddress.LinkLayerAddress
    leafs["encapsulation"] = hostAddress.Encapsulation
    leafs["selected-encapsulation"] = hostAddress.SelectedEncapsulation
    leafs["origin-encapsulation"] = hostAddress.OriginEncapsulation
    leafs["interface-name"] = hostAddress.InterfaceName
    leafs["location"] = hostAddress.Location
    leafs["is-router"] = hostAddress.IsRouter
    leafs["serg-flags"] = hostAddress.SergFlags
    return leafs
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetBundleName() string { return "cisco_ios_xr" }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetYangName() string { return "host-address" }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) SetParent(parent types.Entity) { hostAddress.parent = parent }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetParent() types.Entity { return hostAddress.parent }

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetParentYangName() string { return "host-addresses" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime
// Last time of reachability
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetFilter() yfilter.YFilter { return lastReachedTime.YFilter }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) SetFilter(yf yfilter.YFilter) { lastReachedTime.YFilter = yf }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetSegmentPath() string {
    return "last-reached-time"
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastReachedTime.Seconds
    return leafs
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetBundleName() string { return "cisco_ios_xr" }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetYangName() string { return "last-reached-time" }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) SetParent(parent types.Entity) { lastReachedTime.parent = parent }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetParent() types.Entity { return lastReachedTime.parent }

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetParentYangName() string { return "host-address" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary
// IPv6 Neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    TotalNeighborEntries interface{}

    // Multicast neighbor summary.
    Multicast Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast

    // Static neighbor summary.
    Static Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static

    // Dynamic neighbor summary.
    Dynamic Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetFilter() yfilter.YFilter { return neighborSummary.YFilter }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) SetFilter(yf yfilter.YFilter) { neighborSummary.YFilter = yf }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetGoName(yname string) string {
    if yname == "total-neighbor-entries" { return "TotalNeighborEntries" }
    if yname == "multicast" { return "Multicast" }
    if yname == "static" { return "Static" }
    if yname == "dynamic" { return "Dynamic" }
    return ""
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetSegmentPath() string {
    return "neighbor-summary"
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multicast" {
        return &neighborSummary.Multicast
    }
    if childYangName == "static" {
        return &neighborSummary.Static
    }
    if childYangName == "dynamic" {
        return &neighborSummary.Dynamic
    }
    return nil
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["multicast"] = &neighborSummary.Multicast
    children["static"] = &neighborSummary.Static
    children["dynamic"] = &neighborSummary.Dynamic
    return children
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-neighbor-entries"] = neighborSummary.TotalNeighborEntries
    return leafs
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetBundleName() string { return "cisco_ios_xr" }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetYangName() string { return "neighbor-summary" }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) SetParent(parent types.Entity) { neighborSummary.parent = parent }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetParent() types.Entity { return neighborSummary.parent }

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast
// Multicast neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total incomplete entries. The type is interface{} with range:
    // 0..4294967295.
    IncompleteEntries interface{}

    // Total reachable entries. The type is interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total stale entries. The type is interface{} with range: 0..4294967295.
    StaleEntries interface{}

    // Total delayed entries. The type is interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total probe entries. The type is interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total deleted entries. The type is interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetFilter() yfilter.YFilter { return multicast.YFilter }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) SetFilter(yf yfilter.YFilter) { multicast.YFilter = yf }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetGoName(yname string) string {
    if yname == "incomplete-entries" { return "IncompleteEntries" }
    if yname == "reachable-entries" { return "ReachableEntries" }
    if yname == "stale-entries" { return "StaleEntries" }
    if yname == "delayed-entries" { return "DelayedEntries" }
    if yname == "probe-entries" { return "ProbeEntries" }
    if yname == "deleted-entries" { return "DeletedEntries" }
    if yname == "subtotal-neighbor-entries" { return "SubtotalNeighborEntries" }
    return ""
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetSegmentPath() string {
    return "multicast"
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["incomplete-entries"] = multicast.IncompleteEntries
    leafs["reachable-entries"] = multicast.ReachableEntries
    leafs["stale-entries"] = multicast.StaleEntries
    leafs["delayed-entries"] = multicast.DelayedEntries
    leafs["probe-entries"] = multicast.ProbeEntries
    leafs["deleted-entries"] = multicast.DeletedEntries
    leafs["subtotal-neighbor-entries"] = multicast.SubtotalNeighborEntries
    return leafs
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetBundleName() string { return "cisco_ios_xr" }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetYangName() string { return "multicast" }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) SetParent(parent types.Entity) { multicast.parent = parent }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetParent() types.Entity { return multicast.parent }

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetParentYangName() string { return "neighbor-summary" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static
// Static neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total incomplete entries. The type is interface{} with range:
    // 0..4294967295.
    IncompleteEntries interface{}

    // Total reachable entries. The type is interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total stale entries. The type is interface{} with range: 0..4294967295.
    StaleEntries interface{}

    // Total delayed entries. The type is interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total probe entries. The type is interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total deleted entries. The type is interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetFilter() yfilter.YFilter { return static.YFilter }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) SetFilter(yf yfilter.YFilter) { static.YFilter = yf }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetGoName(yname string) string {
    if yname == "incomplete-entries" { return "IncompleteEntries" }
    if yname == "reachable-entries" { return "ReachableEntries" }
    if yname == "stale-entries" { return "StaleEntries" }
    if yname == "delayed-entries" { return "DelayedEntries" }
    if yname == "probe-entries" { return "ProbeEntries" }
    if yname == "deleted-entries" { return "DeletedEntries" }
    if yname == "subtotal-neighbor-entries" { return "SubtotalNeighborEntries" }
    return ""
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetSegmentPath() string {
    return "static"
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["incomplete-entries"] = static.IncompleteEntries
    leafs["reachable-entries"] = static.ReachableEntries
    leafs["stale-entries"] = static.StaleEntries
    leafs["delayed-entries"] = static.DelayedEntries
    leafs["probe-entries"] = static.ProbeEntries
    leafs["deleted-entries"] = static.DeletedEntries
    leafs["subtotal-neighbor-entries"] = static.SubtotalNeighborEntries
    return leafs
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetBundleName() string { return "cisco_ios_xr" }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetYangName() string { return "static" }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) SetParent(parent types.Entity) { static.parent = parent }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetParent() types.Entity { return static.parent }

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetParentYangName() string { return "neighbor-summary" }

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic
// Dynamic neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total incomplete entries. The type is interface{} with range:
    // 0..4294967295.
    IncompleteEntries interface{}

    // Total reachable entries. The type is interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total stale entries. The type is interface{} with range: 0..4294967295.
    StaleEntries interface{}

    // Total delayed entries. The type is interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total probe entries. The type is interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total deleted entries. The type is interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetFilter() yfilter.YFilter { return dynamic.YFilter }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) SetFilter(yf yfilter.YFilter) { dynamic.YFilter = yf }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetGoName(yname string) string {
    if yname == "incomplete-entries" { return "IncompleteEntries" }
    if yname == "reachable-entries" { return "ReachableEntries" }
    if yname == "stale-entries" { return "StaleEntries" }
    if yname == "delayed-entries" { return "DelayedEntries" }
    if yname == "probe-entries" { return "ProbeEntries" }
    if yname == "deleted-entries" { return "DeletedEntries" }
    if yname == "subtotal-neighbor-entries" { return "SubtotalNeighborEntries" }
    return ""
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetSegmentPath() string {
    return "dynamic"
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["incomplete-entries"] = dynamic.IncompleteEntries
    leafs["reachable-entries"] = dynamic.ReachableEntries
    leafs["stale-entries"] = dynamic.StaleEntries
    leafs["delayed-entries"] = dynamic.DelayedEntries
    leafs["probe-entries"] = dynamic.ProbeEntries
    leafs["deleted-entries"] = dynamic.DeletedEntries
    leafs["subtotal-neighbor-entries"] = dynamic.SubtotalNeighborEntries
    return leafs
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetBundleName() string { return "cisco_ios_xr" }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetYangName() string { return "dynamic" }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) SetParent(parent types.Entity) { dynamic.parent = parent }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetParent() types.Entity { return dynamic.parent }

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetParentYangName() string { return "neighbor-summary" }

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes
// IPv6 ND list of bundle nodes for a specific
// node
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific bundle node. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode.
    BundleNode []Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetFilter() yfilter.YFilter { return bundleNodes.YFilter }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) SetFilter(yf yfilter.YFilter) { bundleNodes.YFilter = yf }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetGoName(yname string) string {
    if yname == "bundle-node" { return "BundleNode" }
    return ""
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetSegmentPath() string {
    return "bundle-nodes"
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-node" {
        for _, c := range bundleNodes.BundleNode {
            if bundleNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode{}
        bundleNodes.BundleNode = append(bundleNodes.BundleNode, child)
        return &bundleNodes.BundleNode[len(bundleNodes.BundleNode)-1]
    }
    return nil
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bundleNodes.BundleNode {
        children[bundleNodes.BundleNode[i].GetSegmentPath()] = &bundleNodes.BundleNode[i]
    }
    return children
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetBundleName() string { return "cisco_ios_xr" }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetYangName() string { return "bundle-nodes" }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) SetParent(parent types.Entity) { bundleNodes.parent = parent }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetParent() types.Entity { return bundleNodes.parent }

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode
// IPv6 ND operational data for a specific
// bundle node
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The bundle node name. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // Sent sequence num. The type is interface{} with range: 0..4294967295.
    SentSequenceNumber interface{}

    // Received sequence num. The type is interface{} with range: 0..4294967295.
    ReceivedSequenceNumber interface{}

    // State. The type is Ipv6NdBndlState.
    State interface{}

    // State changes. The type is interface{} with range: 0..4294967295.
    StateChanges interface{}

    // Total packet sends. The type is interface{} with range: 0..4294967295.
    SentPackets interface{}

    // Total packet receives. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Uptime of node (secs).
    Age Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetFilter() yfilter.YFilter { return bundleNode.YFilter }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) SetFilter(yf yfilter.YFilter) { bundleNode.YFilter = yf }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "group-id" { return "GroupId" }
    if yname == "process-name" { return "ProcessName" }
    if yname == "sent-sequence-number" { return "SentSequenceNumber" }
    if yname == "received-sequence-number" { return "ReceivedSequenceNumber" }
    if yname == "state" { return "State" }
    if yname == "state-changes" { return "StateChanges" }
    if yname == "sent-packets" { return "SentPackets" }
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "age" { return "Age" }
    return ""
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetSegmentPath() string {
    return "bundle-node" + "[node-name='" + fmt.Sprintf("%v", bundleNode.NodeName) + "']"
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "age" {
        return &bundleNode.Age
    }
    return nil
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["age"] = &bundleNode.Age
    return children
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = bundleNode.NodeName
    leafs["group-id"] = bundleNode.GroupId
    leafs["process-name"] = bundleNode.ProcessName
    leafs["sent-sequence-number"] = bundleNode.SentSequenceNumber
    leafs["received-sequence-number"] = bundleNode.ReceivedSequenceNumber
    leafs["state"] = bundleNode.State
    leafs["state-changes"] = bundleNode.StateChanges
    leafs["sent-packets"] = bundleNode.SentPackets
    leafs["received-packets"] = bundleNode.ReceivedPackets
    return leafs
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetBundleName() string { return "cisco_ios_xr" }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetYangName() string { return "bundle-node" }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) SetParent(parent types.Entity) { bundleNode.parent = parent }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetParent() types.Entity { return bundleNode.parent }

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetParentYangName() string { return "bundle-nodes" }

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age
// Uptime of node (secs)
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetFilter() yfilter.YFilter { return age.YFilter }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) SetFilter(yf yfilter.YFilter) { age.YFilter = yf }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetSegmentPath() string {
    return "age"
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = age.Seconds
    return leafs
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetBundleName() string { return "cisco_ios_xr" }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetYangName() string { return "age" }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) SetParent(parent types.Entity) { age.parent = parent }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetParent() types.Entity { return age.parent }

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetParentYangName() string { return "bundle-node" }

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces
// IPv6 ND list of bundle interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific bundler interface. The type is
    // slice of Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface.
    BundleInterface []Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetFilter() yfilter.YFilter { return bundleInterfaces.YFilter }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) SetFilter(yf yfilter.YFilter) { bundleInterfaces.YFilter = yf }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetGoName(yname string) string {
    if yname == "bundle-interface" { return "BundleInterface" }
    return ""
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetSegmentPath() string {
    return "bundle-interfaces"
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-interface" {
        for _, c := range bundleInterfaces.BundleInterface {
            if bundleInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface{}
        bundleInterfaces.BundleInterface = append(bundleInterfaces.BundleInterface, child)
        return &bundleInterfaces.BundleInterface[len(bundleInterfaces.BundleInterface)-1]
    }
    return nil
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bundleInterfaces.BundleInterface {
        children[bundleInterfaces.BundleInterface[i].GetSegmentPath()] = &bundleInterfaces.BundleInterface[i]
    }
    return children
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetYangName() string { return "bundle-interfaces" }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) SetParent(parent types.Entity) { bundleInterfaces.parent = parent }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetParent() types.Entity { return bundleInterfaces.parent }

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface
// IPv6 ND operational data for a specific
// bundler interface
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Parent interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterfaceName interface{}

    // Interface type. The type is interface{} with range: 0..4294967295.
    Iftype interface{}

    // MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // etype. The type is interface{} with range: 0..4294967295.
    Etype interface{}

    // vlan tag/id/ucv. The type is interface{} with range: 0..65535.
    VlanTag interface{}

    // mac address size. The type is interface{} with range: 0..4294967295.
    MacAddrSize interface{}

    // mac address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddr interface{}

    // If true, interface is enabled. The type is bool.
    IsInterfaceEnabled interface{}

    // If true, IPv6 is enabled. The type is bool.
    IsIpv6Enabled interface{}

    // If true, MPLS is enabled. The type is bool.
    IsMplsEnabled interface{}

    // List of member links. The type is slice of interface{} with range:
    // 0..4294967295.
    MemberLink []interface{}

    // ND interface parameters.
    NdParameters Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters

    // Link local address.
    LocalAddress Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress

    // List of ND global addresses. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress.
    GlobalAddress []Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress

    // List of member nodes. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode.
    MemberNode []Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetFilter() yfilter.YFilter { return bundleInterface.YFilter }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) SetFilter(yf yfilter.YFilter) { bundleInterface.YFilter = yf }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "parent-interface-name" { return "ParentInterfaceName" }
    if yname == "iftype" { return "Iftype" }
    if yname == "mtu" { return "Mtu" }
    if yname == "etype" { return "Etype" }
    if yname == "vlan-tag" { return "VlanTag" }
    if yname == "mac-addr-size" { return "MacAddrSize" }
    if yname == "mac-addr" { return "MacAddr" }
    if yname == "is-interface-enabled" { return "IsInterfaceEnabled" }
    if yname == "is-ipv6-enabled" { return "IsIpv6Enabled" }
    if yname == "is-mpls-enabled" { return "IsMplsEnabled" }
    if yname == "member-link" { return "MemberLink" }
    if yname == "nd-parameters" { return "NdParameters" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "global-address" { return "GlobalAddress" }
    if yname == "member-node" { return "MemberNode" }
    return ""
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetSegmentPath() string {
    return "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nd-parameters" {
        return &bundleInterface.NdParameters
    }
    if childYangName == "local-address" {
        return &bundleInterface.LocalAddress
    }
    if childYangName == "global-address" {
        for _, c := range bundleInterface.GlobalAddress {
            if bundleInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress{}
        bundleInterface.GlobalAddress = append(bundleInterface.GlobalAddress, child)
        return &bundleInterface.GlobalAddress[len(bundleInterface.GlobalAddress)-1]
    }
    if childYangName == "member-node" {
        for _, c := range bundleInterface.MemberNode {
            if bundleInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode{}
        bundleInterface.MemberNode = append(bundleInterface.MemberNode, child)
        return &bundleInterface.MemberNode[len(bundleInterface.MemberNode)-1]
    }
    return nil
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nd-parameters"] = &bundleInterface.NdParameters
    children["local-address"] = &bundleInterface.LocalAddress
    for i := range bundleInterface.GlobalAddress {
        children[bundleInterface.GlobalAddress[i].GetSegmentPath()] = &bundleInterface.GlobalAddress[i]
    }
    for i := range bundleInterface.MemberNode {
        children[bundleInterface.MemberNode[i].GetSegmentPath()] = &bundleInterface.MemberNode[i]
    }
    return children
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bundleInterface.InterfaceName
    leafs["parent-interface-name"] = bundleInterface.ParentInterfaceName
    leafs["iftype"] = bundleInterface.Iftype
    leafs["mtu"] = bundleInterface.Mtu
    leafs["etype"] = bundleInterface.Etype
    leafs["vlan-tag"] = bundleInterface.VlanTag
    leafs["mac-addr-size"] = bundleInterface.MacAddrSize
    leafs["mac-addr"] = bundleInterface.MacAddr
    leafs["is-interface-enabled"] = bundleInterface.IsInterfaceEnabled
    leafs["is-ipv6-enabled"] = bundleInterface.IsIpv6Enabled
    leafs["is-mpls-enabled"] = bundleInterface.IsMplsEnabled
    leafs["member-link"] = bundleInterface.MemberLink
    return leafs
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetYangName() string { return "bundle-interface" }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) SetParent(parent types.Entity) { bundleInterface.parent = parent }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetParent() types.Entity { return bundleInterface.parent }

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetParentYangName() string { return "bundle-interfaces" }

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters
// ND interface parameters
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If true, DAD (D.. A.. D..) is enabled otherwise it is disabled. The type is
    // bool.
    IsDadEnabled interface{}

    // DAD attempt count. The type is interface{} with range: 0..4294967295.
    DadAttempts interface{}

    // ICMP redirect flag. The type is bool.
    IsIcmPv6Redirect interface{}

    // Flag used for utilising DHCP. The type is bool.
    IsDhcpManaged interface{}

    // Flag used to manage routable address. The type is bool.
    IsRouteAddressManaged interface{}

    // Suppress flag. The type is bool.
    IsSuppressed interface{}

    // unicast RA send flag. The type is bool.
    SendUnicastRa interface{}

    // ND retransmit interval in msec. The type is interface{} with range:
    // 0..4294967295.
    NdRetransmitInterval interface{}

    // ND router advertisement minimum transmit interval in sec. The type is
    // interface{} with range: 0..4294967295.
    NdMinTransmitInterval interface{}

    // ND router advertisement maximum transmit interval in sec. The type is
    // interface{} with range: 0..4294967295.
    NdMaxTransmitInterval interface{}

    // ND router advertisement life time in sec. The type is interface{} with
    // range: 0..4294967295.
    NdAdvertisementLifetime interface{}

    // Time to reach ND in msec. The type is interface{} with range:
    // 0..4294967295.
    NdReachableTime interface{}

    // Completed adjacency limit per interface. The type is interface{} with
    // range: 0..4294967295.
    NdCacheLimit interface{}

    // Completed PROTO entry Count. The type is interface{} with range:
    // 0..4294967295.
    CompleteProtocolCount interface{}

    // Completed GLEAN entry count. The type is interface{} with range:
    // 0..4294967295.
    CompleteGleanCount interface{}

    // Incomplete PROTO entry count. The type is interface{} with range:
    // 0..4294967295.
    IncompleteProtocolCount interface{}

    // Incomplete GLEAN entry count. The type is interface{} with range:
    // 0..4294967295.
    IncompleteGleanCount interface{}

    // Dropped PROTO entry request count. The type is interface{} with range:
    // 0..4294967295.
    DroppedProtocolReqCount interface{}

    // Dropped GLEAN entry lequest count. The type is interface{} with range:
    // 0..4294967295.
    DroppedGleanReqCount interface{}
}

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetFilter() yfilter.YFilter { return ndParameters.YFilter }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) SetFilter(yf yfilter.YFilter) { ndParameters.YFilter = yf }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetGoName(yname string) string {
    if yname == "is-dad-enabled" { return "IsDadEnabled" }
    if yname == "dad-attempts" { return "DadAttempts" }
    if yname == "is-icm-pv6-redirect" { return "IsIcmPv6Redirect" }
    if yname == "is-dhcp-managed" { return "IsDhcpManaged" }
    if yname == "is-route-address-managed" { return "IsRouteAddressManaged" }
    if yname == "is-suppressed" { return "IsSuppressed" }
    if yname == "send-unicast-ra" { return "SendUnicastRa" }
    if yname == "nd-retransmit-interval" { return "NdRetransmitInterval" }
    if yname == "nd-min-transmit-interval" { return "NdMinTransmitInterval" }
    if yname == "nd-max-transmit-interval" { return "NdMaxTransmitInterval" }
    if yname == "nd-advertisement-lifetime" { return "NdAdvertisementLifetime" }
    if yname == "nd-reachable-time" { return "NdReachableTime" }
    if yname == "nd-cache-limit" { return "NdCacheLimit" }
    if yname == "complete-protocol-count" { return "CompleteProtocolCount" }
    if yname == "complete-glean-count" { return "CompleteGleanCount" }
    if yname == "incomplete-protocol-count" { return "IncompleteProtocolCount" }
    if yname == "incomplete-glean-count" { return "IncompleteGleanCount" }
    if yname == "dropped-protocol-req-count" { return "DroppedProtocolReqCount" }
    if yname == "dropped-glean-req-count" { return "DroppedGleanReqCount" }
    return ""
}

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetSegmentPath() string {
    return "nd-parameters"
}

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-dad-enabled"] = ndParameters.IsDadEnabled
    leafs["dad-attempts"] = ndParameters.DadAttempts
    leafs["is-icm-pv6-redirect"] = ndParameters.IsIcmPv6Redirect
    leafs["is-dhcp-managed"] = ndParameters.IsDhcpManaged
    leafs["is-route-address-managed"] = ndParameters.IsRouteAddressManaged
    leafs["is-suppressed"] = ndParameters.IsSuppressed
    leafs["send-unicast-ra"] = ndParameters.SendUnicastRa
    leafs["nd-retransmit-interval"] = ndParameters.NdRetransmitInterval
    leafs["nd-min-transmit-interval"] = ndParameters.NdMinTransmitInterval
    leafs["nd-max-transmit-interval"] = ndParameters.NdMaxTransmitInterval
    leafs["nd-advertisement-lifetime"] = ndParameters.NdAdvertisementLifetime
    leafs["nd-reachable-time"] = ndParameters.NdReachableTime
    leafs["nd-cache-limit"] = ndParameters.NdCacheLimit
    leafs["complete-protocol-count"] = ndParameters.CompleteProtocolCount
    leafs["complete-glean-count"] = ndParameters.CompleteGleanCount
    leafs["incomplete-protocol-count"] = ndParameters.IncompleteProtocolCount
    leafs["incomplete-glean-count"] = ndParameters.IncompleteGleanCount
    leafs["dropped-protocol-req-count"] = ndParameters.DroppedProtocolReqCount
    leafs["dropped-glean-req-count"] = ndParameters.DroppedGleanReqCount
    return leafs
}

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetBundleName() string { return "cisco_ios_xr" }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetYangName() string { return "nd-parameters" }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) SetParent(parent types.Entity) { ndParameters.parent = parent }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetParent() types.Entity { return ndParameters.parent }

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetParentYangName() string { return "bundle-interface" }

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress
// Link local address
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetParentYangName() string { return "bundle-interface" }

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress
// List of ND global addresses
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetFilter() yfilter.YFilter { return globalAddress.YFilter }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) SetFilter(yf yfilter.YFilter) { globalAddress.YFilter = yf }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetSegmentPath() string {
    return "global-address"
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = globalAddress.Ipv6Address
    return leafs
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetYangName() string { return "global-address" }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) SetParent(parent types.Entity) { globalAddress.parent = parent }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetParent() types.Entity { return globalAddress.parent }

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetParentYangName() string { return "bundle-interface" }

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode
// List of member nodes
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Number of links on the node. The type is interface{} with range:
    // 0..4294967295.
    TotalLinks interface{}
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetFilter() yfilter.YFilter { return memberNode.YFilter }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) SetFilter(yf yfilter.YFilter) { memberNode.YFilter = yf }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "total-links" { return "TotalLinks" }
    return ""
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetSegmentPath() string {
    return "member-node"
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = memberNode.NodeName
    leafs["total-links"] = memberNode.TotalLinks
    return leafs
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetBundleName() string { return "cisco_ios_xr" }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetYangName() string { return "member-node" }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) SetParent(parent types.Entity) { memberNode.parent = parent }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetParent() types.Entity { return memberNode.parent }

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetParentYangName() string { return "bundle-interface" }

// Ipv6NodeDiscovery_Nodes_Node_Interfaces
// IPv6 node discovery list of interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6  node discovery operational data for a specific node and interface.
    // The type is slice of Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface.
    Interface []Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface
// IPv6  node discovery operational data for a
// specific node and interface
type Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // If true, DAD (D.. A.. D..) is enabled otherwise it is disabled. The type is
    // bool.
    IsDadEnabled interface{}

    // DAD attempt count. The type is interface{} with range: 0..4294967295.
    DadAttempts interface{}

    // ICMP redirect flag. The type is bool.
    IsIcmPv6Redirect interface{}

    // Flag used for utilising DHCP. The type is bool.
    IsDhcpManaged interface{}

    // Flag used to manage routable address. The type is bool.
    IsRouteAddressManaged interface{}

    // Suppress flag. The type is bool.
    IsSuppressed interface{}

    // unicast RA send flag. The type is bool.
    SendUnicastRa interface{}

    // ND retransmit interval in msec. The type is interface{} with range:
    // 0..4294967295.
    NdRetransmitInterval interface{}

    // ND router advertisement minimum transmit interval in sec. The type is
    // interface{} with range: 0..4294967295.
    NdMinTransmitInterval interface{}

    // ND router advertisement maximum transmit interval in sec. The type is
    // interface{} with range: 0..4294967295.
    NdMaxTransmitInterval interface{}

    // ND router advertisement life time in sec. The type is interface{} with
    // range: 0..4294967295.
    NdAdvertisementLifetime interface{}

    // Time to reach ND in msec. The type is interface{} with range:
    // 0..4294967295.
    NdReachableTime interface{}

    // Completed adjacency limit per interface. The type is interface{} with
    // range: 0..4294967295.
    NdCacheLimit interface{}

    // Completed PROTO entry Count. The type is interface{} with range:
    // 0..4294967295.
    CompleteProtocolCount interface{}

    // Completed GLEAN entry count. The type is interface{} with range:
    // 0..4294967295.
    CompleteGleanCount interface{}

    // Incomplete PROTO entry count. The type is interface{} with range:
    // 0..4294967295.
    IncompleteProtocolCount interface{}

    // Incomplete GLEAN entry count. The type is interface{} with range:
    // 0..4294967295.
    IncompleteGleanCount interface{}

    // Dropped PROTO entry request count. The type is interface{} with range:
    // 0..4294967295.
    DroppedProtocolReqCount interface{}

    // Dropped GLEAN entry lequest count. The type is interface{} with range:
    // 0..4294967295.
    DroppedGleanReqCount interface{}
}

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "is-dad-enabled" { return "IsDadEnabled" }
    if yname == "dad-attempts" { return "DadAttempts" }
    if yname == "is-icm-pv6-redirect" { return "IsIcmPv6Redirect" }
    if yname == "is-dhcp-managed" { return "IsDhcpManaged" }
    if yname == "is-route-address-managed" { return "IsRouteAddressManaged" }
    if yname == "is-suppressed" { return "IsSuppressed" }
    if yname == "send-unicast-ra" { return "SendUnicastRa" }
    if yname == "nd-retransmit-interval" { return "NdRetransmitInterval" }
    if yname == "nd-min-transmit-interval" { return "NdMinTransmitInterval" }
    if yname == "nd-max-transmit-interval" { return "NdMaxTransmitInterval" }
    if yname == "nd-advertisement-lifetime" { return "NdAdvertisementLifetime" }
    if yname == "nd-reachable-time" { return "NdReachableTime" }
    if yname == "nd-cache-limit" { return "NdCacheLimit" }
    if yname == "complete-protocol-count" { return "CompleteProtocolCount" }
    if yname == "complete-glean-count" { return "CompleteGleanCount" }
    if yname == "incomplete-protocol-count" { return "IncompleteProtocolCount" }
    if yname == "incomplete-glean-count" { return "IncompleteGleanCount" }
    if yname == "dropped-protocol-req-count" { return "DroppedProtocolReqCount" }
    if yname == "dropped-glean-req-count" { return "DroppedGleanReqCount" }
    return ""
}

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["is-dad-enabled"] = self.IsDadEnabled
    leafs["dad-attempts"] = self.DadAttempts
    leafs["is-icm-pv6-redirect"] = self.IsIcmPv6Redirect
    leafs["is-dhcp-managed"] = self.IsDhcpManaged
    leafs["is-route-address-managed"] = self.IsRouteAddressManaged
    leafs["is-suppressed"] = self.IsSuppressed
    leafs["send-unicast-ra"] = self.SendUnicastRa
    leafs["nd-retransmit-interval"] = self.NdRetransmitInterval
    leafs["nd-min-transmit-interval"] = self.NdMinTransmitInterval
    leafs["nd-max-transmit-interval"] = self.NdMaxTransmitInterval
    leafs["nd-advertisement-lifetime"] = self.NdAdvertisementLifetime
    leafs["nd-reachable-time"] = self.NdReachableTime
    leafs["nd-cache-limit"] = self.NdCacheLimit
    leafs["complete-protocol-count"] = self.CompleteProtocolCount
    leafs["complete-glean-count"] = self.CompleteGleanCount
    leafs["incomplete-protocol-count"] = self.IncompleteProtocolCount
    leafs["incomplete-glean-count"] = self.IncompleteGleanCount
    leafs["dropped-protocol-req-count"] = self.DroppedProtocolReqCount
    leafs["dropped-glean-req-count"] = self.DroppedGleanReqCount
    return leafs
}

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters
// IPv6 ND virtual router information for a
// specific interface
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ND virtual  router operational data for a specific interface. The type
    // is slice of Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter.
    NdVirtualRouter []Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetFilter() yfilter.YFilter { return ndVirtualRouters.YFilter }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) SetFilter(yf yfilter.YFilter) { ndVirtualRouters.YFilter = yf }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetGoName(yname string) string {
    if yname == "nd-virtual-router" { return "NdVirtualRouter" }
    return ""
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetSegmentPath() string {
    return "nd-virtual-routers"
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nd-virtual-router" {
        for _, c := range ndVirtualRouters.NdVirtualRouter {
            if ndVirtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter{}
        ndVirtualRouters.NdVirtualRouter = append(ndVirtualRouters.NdVirtualRouter, child)
        return &ndVirtualRouters.NdVirtualRouter[len(ndVirtualRouters.NdVirtualRouter)-1]
    }
    return nil
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ndVirtualRouters.NdVirtualRouter {
        children[ndVirtualRouters.NdVirtualRouter[i].GetSegmentPath()] = &ndVirtualRouters.NdVirtualRouter[i]
    }
    return children
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetYangName() string { return "nd-virtual-routers" }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) SetParent(parent types.Entity) { ndVirtualRouters.parent = parent }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetParent() types.Entity { return ndVirtualRouters.parent }

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter
// IPv6 ND virtual  router operational data for
// a specific interface
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Link-Layer Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    LinkLayerAddress interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    Context interface{}

    // VR state. The type is Ipv6NdShVrState.
    State interface{}

    // VR Flags. The type is Ipv6NdShVrFlags.
    Flags interface{}

    // Virtual Global Address Count. The type is interface{} with range:
    // 0..4294967295.
    VrGlAddrCt interface{}

    // Link local address.
    LocalAddress Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress

    // List of ND global addresses. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress.
    VrGlobalAddress []Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetFilter() yfilter.YFilter { return ndVirtualRouter.YFilter }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) SetFilter(yf yfilter.YFilter) { ndVirtualRouter.YFilter = yf }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "context" { return "Context" }
    if yname == "state" { return "State" }
    if yname == "flags" { return "Flags" }
    if yname == "vr-gl-addr-ct" { return "VrGlAddrCt" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "vr-global-address" { return "VrGlobalAddress" }
    return ""
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetSegmentPath() string {
    return "nd-virtual-router" + "[interface-name='" + fmt.Sprintf("%v", ndVirtualRouter.InterfaceName) + "']"
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &ndVirtualRouter.LocalAddress
    }
    if childYangName == "vr-global-address" {
        for _, c := range ndVirtualRouter.VrGlobalAddress {
            if ndVirtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress{}
        ndVirtualRouter.VrGlobalAddress = append(ndVirtualRouter.VrGlobalAddress, child)
        return &ndVirtualRouter.VrGlobalAddress[len(ndVirtualRouter.VrGlobalAddress)-1]
    }
    return nil
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &ndVirtualRouter.LocalAddress
    for i := range ndVirtualRouter.VrGlobalAddress {
        children[ndVirtualRouter.VrGlobalAddress[i].GetSegmentPath()] = &ndVirtualRouter.VrGlobalAddress[i]
    }
    return children
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = ndVirtualRouter.InterfaceName
    leafs["link-layer-address"] = ndVirtualRouter.LinkLayerAddress
    leafs["context"] = ndVirtualRouter.Context
    leafs["state"] = ndVirtualRouter.State
    leafs["flags"] = ndVirtualRouter.Flags
    leafs["vr-gl-addr-ct"] = ndVirtualRouter.VrGlAddrCt
    return leafs
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetYangName() string { return "nd-virtual-router" }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) SetParent(parent types.Entity) { ndVirtualRouter.parent = parent }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetParent() types.Entity { return ndVirtualRouter.parent }

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetParentYangName() string { return "nd-virtual-routers" }

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress
// Link local address
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetParentYangName() string { return "nd-virtual-router" }

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress
// List of ND global addresses
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetFilter() yfilter.YFilter { return vrGlobalAddress.YFilter }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) SetFilter(yf yfilter.YFilter) { vrGlobalAddress.YFilter = yf }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetSegmentPath() string {
    return "vr-global-address"
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = vrGlobalAddress.Ipv6Address
    return leafs
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetYangName() string { return "vr-global-address" }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) SetParent(parent types.Entity) { vrGlobalAddress.parent = parent }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetParent() types.Entity { return vrGlobalAddress.parent }

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetParentYangName() string { return "nd-virtual-router" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces
// IPv6 ND list of SLAAC MGMT interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific slaac interface. The type is slice
    // of Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface.
    SlaacInterface []Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetFilter() yfilter.YFilter { return slaacInterfaces.YFilter }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) SetFilter(yf yfilter.YFilter) { slaacInterfaces.YFilter = yf }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetGoName(yname string) string {
    if yname == "slaac-interface" { return "SlaacInterface" }
    return ""
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetSegmentPath() string {
    return "slaac-interfaces"
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slaac-interface" {
        for _, c := range slaacInterfaces.SlaacInterface {
            if slaacInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface{}
        slaacInterfaces.SlaacInterface = append(slaacInterfaces.SlaacInterface, child)
        return &slaacInterfaces.SlaacInterface[len(slaacInterfaces.SlaacInterface)-1]
    }
    return nil
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slaacInterfaces.SlaacInterface {
        children[slaacInterfaces.SlaacInterface[i].GetSegmentPath()] = &slaacInterfaces.SlaacInterface[i]
    }
    return children
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetYangName() string { return "slaac-interfaces" }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) SetParent(parent types.Entity) { slaacInterfaces.parent = parent }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetParent() types.Entity { return slaacInterfaces.parent }

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetParentYangName() string { return "node" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface
// IPv6 ND operational data for a specific slaac
// interface
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IPv6 ND operational data for a specific slaac interface.
    RouterAdvertDetail Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetFilter() yfilter.YFilter { return slaacInterface.YFilter }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) SetFilter(yf yfilter.YFilter) { slaacInterface.YFilter = yf }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "router-advert-detail" { return "RouterAdvertDetail" }
    return ""
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetSegmentPath() string {
    return "slaac-interface" + "[interface-name='" + fmt.Sprintf("%v", slaacInterface.InterfaceName) + "']"
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router-advert-detail" {
        return &slaacInterface.RouterAdvertDetail
    }
    return nil
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router-advert-detail"] = &slaacInterface.RouterAdvertDetail
    return children
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = slaacInterface.InterfaceName
    return leafs
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetBundleName() string { return "cisco_ios_xr" }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetYangName() string { return "slaac-interface" }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) SetParent(parent types.Entity) { slaacInterface.parent = parent }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetParent() types.Entity { return slaacInterface.parent }

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetParentYangName() string { return "slaac-interfaces" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail
// IPv6 ND operational data for a specific
// slaac interface
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // idb. The type is string with pattern: [a-zA-Z0-9./-]+.
    Idb interface{}

    // slaac db. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra.
    Ra []Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetFilter() yfilter.YFilter { return routerAdvertDetail.YFilter }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) SetFilter(yf yfilter.YFilter) { routerAdvertDetail.YFilter = yf }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetGoName(yname string) string {
    if yname == "idb" { return "Idb" }
    if yname == "ra" { return "Ra" }
    return ""
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetSegmentPath() string {
    return "router-advert-detail"
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ra" {
        for _, c := range routerAdvertDetail.Ra {
            if routerAdvertDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra{}
        routerAdvertDetail.Ra = append(routerAdvertDetail.Ra, child)
        return &routerAdvertDetail.Ra[len(routerAdvertDetail.Ra)-1]
    }
    return nil
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routerAdvertDetail.Ra {
        children[routerAdvertDetail.Ra[i].GetSegmentPath()] = &routerAdvertDetail.Ra[i]
    }
    return children
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idb"] = routerAdvertDetail.Idb
    return leafs
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetBundleName() string { return "cisco_ios_xr" }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetYangName() string { return "router-advert-detail" }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) SetParent(parent types.Entity) { routerAdvertDetail.parent = parent }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetParent() types.Entity { return routerAdvertDetail.parent }

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetParentYangName() string { return "slaac-interface" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra
// slaac db
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // hops. The type is interface{} with range: 0..4294967295.
    Hops interface{}

    // flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // lifetime. The type is interface{} with range: 0..4294967295.
    LifeTime interface{}

    // mtu. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // errmsg. The type is bool.
    ErrMsg interface{}

    // vrf id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // tbl id. The type is interface{} with range: 0..4294967295.
    U6TblId interface{}

    // proto id. The type is interface{} with range: 0..65535.
    RibProtoid interface{}

    // router. The type is bool.
    DefaultRouter interface{}

    // reach. The type is interface{} with range: 0..4294967295.
    Reachability interface{}

    // elapsedRATime.
    ElapsedRaTime Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime

    // reachabletime.
    ReachableTime Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime

    // retranstime.
    RetransTime Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime

    // Prefix Queue. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ.
    PrefixQ []Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetFilter() yfilter.YFilter { return ra.YFilter }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) SetFilter(yf yfilter.YFilter) { ra.YFilter = yf }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "hops" { return "Hops" }
    if yname == "flags" { return "Flags" }
    if yname == "life-time" { return "LifeTime" }
    if yname == "mtu" { return "Mtu" }
    if yname == "err-msg" { return "ErrMsg" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "u6-tbl-id" { return "U6TblId" }
    if yname == "rib-protoid" { return "RibProtoid" }
    if yname == "default-router" { return "DefaultRouter" }
    if yname == "reachability" { return "Reachability" }
    if yname == "elapsed-ra-time" { return "ElapsedRaTime" }
    if yname == "reachable-time" { return "ReachableTime" }
    if yname == "retrans-time" { return "RetransTime" }
    if yname == "prefix-q" { return "PrefixQ" }
    return ""
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetSegmentPath() string {
    return "ra"
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "elapsed-ra-time" {
        return &ra.ElapsedRaTime
    }
    if childYangName == "reachable-time" {
        return &ra.ReachableTime
    }
    if childYangName == "retrans-time" {
        return &ra.RetransTime
    }
    if childYangName == "prefix-q" {
        for _, c := range ra.PrefixQ {
            if ra.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ{}
        ra.PrefixQ = append(ra.PrefixQ, child)
        return &ra.PrefixQ[len(ra.PrefixQ)-1]
    }
    return nil
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["elapsed-ra-time"] = &ra.ElapsedRaTime
    children["reachable-time"] = &ra.ReachableTime
    children["retrans-time"] = &ra.RetransTime
    for i := range ra.PrefixQ {
        children[ra.PrefixQ[i].GetSegmentPath()] = &ra.PrefixQ[i]
    }
    return children
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ra.Address
    leafs["hops"] = ra.Hops
    leafs["flags"] = ra.Flags
    leafs["life-time"] = ra.LifeTime
    leafs["mtu"] = ra.Mtu
    leafs["err-msg"] = ra.ErrMsg
    leafs["vrf-id"] = ra.VrfId
    leafs["u6-tbl-id"] = ra.U6TblId
    leafs["rib-protoid"] = ra.RibProtoid
    leafs["default-router"] = ra.DefaultRouter
    leafs["reachability"] = ra.Reachability
    return leafs
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetBundleName() string { return "cisco_ios_xr" }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetYangName() string { return "ra" }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) SetParent(parent types.Entity) { ra.parent = parent }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetParent() types.Entity { return ra.parent }

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetParentYangName() string { return "router-advert-detail" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime
// elapsedRATime
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetFilter() yfilter.YFilter { return elapsedRaTime.YFilter }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) SetFilter(yf yfilter.YFilter) { elapsedRaTime.YFilter = yf }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetSegmentPath() string {
    return "elapsed-ra-time"
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = elapsedRaTime.Seconds
    return leafs
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetBundleName() string { return "cisco_ios_xr" }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetYangName() string { return "elapsed-ra-time" }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) SetParent(parent types.Entity) { elapsedRaTime.parent = parent }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetParent() types.Entity { return elapsedRaTime.parent }

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetParentYangName() string { return "ra" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime
// reachabletime
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetFilter() yfilter.YFilter { return reachableTime.YFilter }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) SetFilter(yf yfilter.YFilter) { reachableTime.YFilter = yf }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetSegmentPath() string {
    return "reachable-time"
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = reachableTime.Seconds
    return leafs
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetBundleName() string { return "cisco_ios_xr" }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetYangName() string { return "reachable-time" }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) SetParent(parent types.Entity) { reachableTime.parent = parent }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetParent() types.Entity { return reachableTime.parent }

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetParentYangName() string { return "ra" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime
// retranstime
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetFilter() yfilter.YFilter { return retransTime.YFilter }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) SetFilter(yf yfilter.YFilter) { retransTime.YFilter = yf }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetSegmentPath() string {
    return "retrans-time"
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = retransTime.Seconds
    return leafs
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetBundleName() string { return "cisco_ios_xr" }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetYangName() string { return "retrans-time" }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) SetParent(parent types.Entity) { retransTime.parent = parent }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetParent() types.Entity { return retransTime.parent }

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetParentYangName() string { return "ra" }

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ
// Prefix Queue
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixAddress interface{}

    // IPv6 Auto generated address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Eui64 interface{}

    // Valid Life Time. The type is interface{} with range: 0..4294967295.
    ValidLifeTime interface{}

    // Preferred Life Time. The type is interface{} with range: 0..4294967295.
    PreferredLifeTime interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLen interface{}

    // IPv6 Address Specific Flags. The type is interface{} with range:
    // 0..4294967295.
    Flags interface{}

    // Prefix Address Specific Flags. The type is interface{} with range:
    // 0..4294967295.
    PfxFlags interface{}
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetFilter() yfilter.YFilter { return prefixQ.YFilter }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) SetFilter(yf yfilter.YFilter) { prefixQ.YFilter = yf }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetGoName(yname string) string {
    if yname == "prefix-address" { return "PrefixAddress" }
    if yname == "eui64" { return "Eui64" }
    if yname == "valid-life-time" { return "ValidLifeTime" }
    if yname == "preferred-life-time" { return "PreferredLifeTime" }
    if yname == "prefix-len" { return "PrefixLen" }
    if yname == "flags" { return "Flags" }
    if yname == "pfx-flags" { return "PfxFlags" }
    return ""
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetSegmentPath() string {
    return "prefix-q"
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-address"] = prefixQ.PrefixAddress
    leafs["eui64"] = prefixQ.Eui64
    leafs["valid-life-time"] = prefixQ.ValidLifeTime
    leafs["preferred-life-time"] = prefixQ.PreferredLifeTime
    leafs["prefix-len"] = prefixQ.PrefixLen
    leafs["flags"] = prefixQ.Flags
    leafs["pfx-flags"] = prefixQ.PfxFlags
    return leafs
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetBundleName() string { return "cisco_ios_xr" }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetYangName() string { return "prefix-q" }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) SetParent(parent types.Entity) { prefixQ.parent = parent }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetParent() types.Entity { return prefixQ.parent }

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetParentYangName() string { return "ra" }

