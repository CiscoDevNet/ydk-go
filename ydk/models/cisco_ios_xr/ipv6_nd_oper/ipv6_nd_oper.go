// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-nd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-node-discovery: IPv6 node discovery operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// Ipv6NdShVrFlags represents IPv6 ND VR Entry Flags Type 
type Ipv6NdShVrFlags string

const (
    // None
    Ipv6NdShVrFlags_no_flags Ipv6NdShVrFlags = "no-flags"

    // Final RA
    Ipv6NdShVrFlags_final_ra Ipv6NdShVrFlags = "final-ra"
)

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
    Ipv6NdBndlState_error_ Ipv6NdBndlState = "error"

    // Wait state
    Ipv6NdBndlState_wait Ipv6NdBndlState = "wait"
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
    Ipv6NdShState_delete_ Ipv6NdShState = "delete"
)

// Ipv6NodeDiscovery
// IPv6 node discovery operational data
type Ipv6NodeDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 node discovery list of nodes.
    Nodes Ipv6NodeDiscovery_Nodes
}

func (ipv6NodeDiscovery *Ipv6NodeDiscovery) GetEntityData() *types.CommonEntityData {
    ipv6NodeDiscovery.EntityData.YFilter = ipv6NodeDiscovery.YFilter
    ipv6NodeDiscovery.EntityData.YangName = "ipv6-node-discovery"
    ipv6NodeDiscovery.EntityData.BundleName = "cisco_ios_xr"
    ipv6NodeDiscovery.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-nd-oper"
    ipv6NodeDiscovery.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery"
    ipv6NodeDiscovery.EntityData.AbsolutePath = ipv6NodeDiscovery.EntityData.SegmentPath
    ipv6NodeDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6NodeDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6NodeDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6NodeDiscovery.EntityData.Children = types.NewOrderedMap()
    ipv6NodeDiscovery.EntityData.Children.Append("nodes", types.YChild{"Nodes", &ipv6NodeDiscovery.Nodes})
    ipv6NodeDiscovery.EntityData.Leafs = types.NewOrderedMap()

    ipv6NodeDiscovery.EntityData.YListKeys = []string {}

    return &(ipv6NodeDiscovery.EntityData)
}

// Ipv6NodeDiscovery_Nodes
// IPv6 node discovery list of nodes
type Ipv6NodeDiscovery_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 node discovery operational data for a particular node. The type is
    // slice of Ipv6NodeDiscovery_Nodes_Node.
    Node []*Ipv6NodeDiscovery_Nodes_Node
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ipv6-node-discovery"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/" + nodes.EntityData.SegmentPath
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

// Ipv6NodeDiscovery_Nodes_Node
// IPv6 node discovery operational data for a
// particular node
type Ipv6NodeDiscovery_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // IPv6 node discovery list of neighbor interfaces.
    NeighborInterfaces Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces

    // Summary of IPv6 Neighbors.
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

func (node *Ipv6NodeDiscovery_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("neighbor-interfaces", types.YChild{"NeighborInterfaces", &node.NeighborInterfaces})
    node.EntityData.Children.Append("neighbor-summary", types.YChild{"NeighborSummary", &node.NeighborSummary})
    node.EntityData.Children.Append("bundle-nodes", types.YChild{"BundleNodes", &node.BundleNodes})
    node.EntityData.Children.Append("bundle-interfaces", types.YChild{"BundleInterfaces", &node.BundleInterfaces})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("nd-virtual-routers", types.YChild{"NdVirtualRouters", &node.NdVirtualRouters})
    node.EntityData.Children.Append("slaac-interfaces", types.YChild{"SlaacInterfaces", &node.SlaacInterfaces})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces
// IPv6 node discovery list of neighbor
// interfaces
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 node discovery neighbor interface. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface.
    NeighborInterface []*Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetEntityData() *types.CommonEntityData {
    neighborInterfaces.EntityData.YFilter = neighborInterfaces.YFilter
    neighborInterfaces.EntityData.YangName = "neighbor-interfaces"
    neighborInterfaces.EntityData.BundleName = "cisco_ios_xr"
    neighborInterfaces.EntityData.ParentYangName = "node"
    neighborInterfaces.EntityData.SegmentPath = "neighbor-interfaces"
    neighborInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + neighborInterfaces.EntityData.SegmentPath
    neighborInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborInterfaces.EntityData.Children = types.NewOrderedMap()
    neighborInterfaces.EntityData.Children.Append("neighbor-interface", types.YChild{"NeighborInterface", nil})
    for i := range neighborInterfaces.NeighborInterface {
        neighborInterfaces.EntityData.Children.Append(types.GetSegmentPath(neighborInterfaces.NeighborInterface[i]), types.YChild{"NeighborInterface", neighborInterfaces.NeighborInterface[i]})
    }
    neighborInterfaces.EntityData.Leafs = types.NewOrderedMap()

    neighborInterfaces.EntityData.YListKeys = []string {}

    return &(neighborInterfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface
// IPv6 node discovery neighbor interface
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // IPv6 node discovery list of neighbor host addresses.
    HostAddresses Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetEntityData() *types.CommonEntityData {
    neighborInterface.EntityData.YFilter = neighborInterface.YFilter
    neighborInterface.EntityData.YangName = "neighbor-interface"
    neighborInterface.EntityData.BundleName = "cisco_ios_xr"
    neighborInterface.EntityData.ParentYangName = "neighbor-interfaces"
    neighborInterface.EntityData.SegmentPath = "neighbor-interface" + types.AddKeyToken(neighborInterface.InterfaceName, "interface-name")
    neighborInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-interfaces/" + neighborInterface.EntityData.SegmentPath
    neighborInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborInterface.EntityData.Children = types.NewOrderedMap()
    neighborInterface.EntityData.Children.Append("host-addresses", types.YChild{"HostAddresses", &neighborInterface.HostAddresses})
    neighborInterface.EntityData.Leafs = types.NewOrderedMap()
    neighborInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", neighborInterface.InterfaceName})

    neighborInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(neighborInterface.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses
// IPv6 node discovery list of neighbor host
// addresses
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Neighbor detailed information. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress.
    HostAddress []*Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetEntityData() *types.CommonEntityData {
    hostAddresses.EntityData.YFilter = hostAddresses.YFilter
    hostAddresses.EntityData.YangName = "host-addresses"
    hostAddresses.EntityData.BundleName = "cisco_ios_xr"
    hostAddresses.EntityData.ParentYangName = "neighbor-interface"
    hostAddresses.EntityData.SegmentPath = "host-addresses"
    hostAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-interfaces/neighbor-interface/" + hostAddresses.EntityData.SegmentPath
    hostAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAddresses.EntityData.Children = types.NewOrderedMap()
    hostAddresses.EntityData.Children.Append("host-address", types.YChild{"HostAddress", nil})
    for i := range hostAddresses.HostAddress {
        hostAddresses.EntityData.Children.Append(types.GetSegmentPath(hostAddresses.HostAddress[i]), types.YChild{"HostAddress", hostAddresses.HostAddress[i]})
    }
    hostAddresses.EntityData.Leafs = types.NewOrderedMap()

    hostAddresses.EntityData.YListKeys = []string {}

    return &(hostAddresses.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress
// IPv6 Neighbor detailed information
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Host Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    HostAddress interface{}

    // Current state. The type is Ipv6NdShState.
    ReachabilityState interface{}

    // IPV6 Link-Layer Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    LinkLayerAddress interface{}

    // Preferred media encap type. The type is Ipv6NdMediaEncap.
    Encapsulation interface{}

    // Selected media encap. The type is Ipv6NdMediaEncap.
    SelectedEncapsulation interface{}

    // Neighbor origin. The type is Ipv6NdNeighborOrigin.
    OriginEncapsulation interface{}

    // Name of Interface. The type is string.
    InterfaceName interface{}

    // Location where the neighbor entry exists. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Location interface{}

    // IsRouter. The type is bool.
    IsRouter interface{}

    // ND serg flags for this entry. The type is interface{} with range:
    // 0..4294967295.
    SergFlags interface{}

    // VRF name for this entry. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Last time of reachability.
    LastReachedTime Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime
}

func (hostAddress *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress) GetEntityData() *types.CommonEntityData {
    hostAddress.EntityData.YFilter = hostAddress.YFilter
    hostAddress.EntityData.YangName = "host-address"
    hostAddress.EntityData.BundleName = "cisco_ios_xr"
    hostAddress.EntityData.ParentYangName = "host-addresses"
    hostAddress.EntityData.SegmentPath = "host-address" + types.AddKeyToken(hostAddress.HostAddress, "host-address")
    hostAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-interfaces/neighbor-interface/host-addresses/" + hostAddress.EntityData.SegmentPath
    hostAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAddress.EntityData.Children = types.NewOrderedMap()
    hostAddress.EntityData.Children.Append("last-reached-time", types.YChild{"LastReachedTime", &hostAddress.LastReachedTime})
    hostAddress.EntityData.Leafs = types.NewOrderedMap()
    hostAddress.EntityData.Leafs.Append("host-address", types.YLeaf{"HostAddress", hostAddress.HostAddress})
    hostAddress.EntityData.Leafs.Append("reachability-state", types.YLeaf{"ReachabilityState", hostAddress.ReachabilityState})
    hostAddress.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", hostAddress.LinkLayerAddress})
    hostAddress.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", hostAddress.Encapsulation})
    hostAddress.EntityData.Leafs.Append("selected-encapsulation", types.YLeaf{"SelectedEncapsulation", hostAddress.SelectedEncapsulation})
    hostAddress.EntityData.Leafs.Append("origin-encapsulation", types.YLeaf{"OriginEncapsulation", hostAddress.OriginEncapsulation})
    hostAddress.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", hostAddress.InterfaceName})
    hostAddress.EntityData.Leafs.Append("location", types.YLeaf{"Location", hostAddress.Location})
    hostAddress.EntityData.Leafs.Append("is-router", types.YLeaf{"IsRouter", hostAddress.IsRouter})
    hostAddress.EntityData.Leafs.Append("serg-flags", types.YLeaf{"SergFlags", hostAddress.SergFlags})
    hostAddress.EntityData.Leafs.Append("vrfid", types.YLeaf{"Vrfid", hostAddress.Vrfid})

    hostAddress.EntityData.YListKeys = []string {"HostAddress"}

    return &(hostAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime
// Last time of reachability
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (lastReachedTime *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress_LastReachedTime) GetEntityData() *types.CommonEntityData {
    lastReachedTime.EntityData.YFilter = lastReachedTime.YFilter
    lastReachedTime.EntityData.YangName = "last-reached-time"
    lastReachedTime.EntityData.BundleName = "cisco_ios_xr"
    lastReachedTime.EntityData.ParentYangName = "host-address"
    lastReachedTime.EntityData.SegmentPath = "last-reached-time"
    lastReachedTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-interfaces/neighbor-interface/host-addresses/host-address/" + lastReachedTime.EntityData.SegmentPath
    lastReachedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReachedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReachedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReachedTime.EntityData.Children = types.NewOrderedMap()
    lastReachedTime.EntityData.Leafs = types.NewOrderedMap()
    lastReachedTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastReachedTime.Seconds})

    lastReachedTime.EntityData.YListKeys = []string {}

    return &(lastReachedTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary
// Summary of IPv6 Neighbors
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    TotalNeighborEntries interface{}

    // Multicast neighbor summary.
    Multicast Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast

    // Static neighbor summary.
    Static Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static

    // Dynamic neighbor summary.
    Dynamic Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic

    // Sync neighbor summary.
    Sync Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Sync

    // StaticSync neighbor summary.
    StaticSync Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_StaticSync

    // DynamicSync neighbor summary.
    DynamicSync Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_DynamicSync
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetEntityData() *types.CommonEntityData {
    neighborSummary.EntityData.YFilter = neighborSummary.YFilter
    neighborSummary.EntityData.YangName = "neighbor-summary"
    neighborSummary.EntityData.BundleName = "cisco_ios_xr"
    neighborSummary.EntityData.ParentYangName = "node"
    neighborSummary.EntityData.SegmentPath = "neighbor-summary"
    neighborSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + neighborSummary.EntityData.SegmentPath
    neighborSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborSummary.EntityData.Children = types.NewOrderedMap()
    neighborSummary.EntityData.Children.Append("multicast", types.YChild{"Multicast", &neighborSummary.Multicast})
    neighborSummary.EntityData.Children.Append("static", types.YChild{"Static", &neighborSummary.Static})
    neighborSummary.EntityData.Children.Append("dynamic", types.YChild{"Dynamic", &neighborSummary.Dynamic})
    neighborSummary.EntityData.Children.Append("sync", types.YChild{"Sync", &neighborSummary.Sync})
    neighborSummary.EntityData.Children.Append("static-sync", types.YChild{"StaticSync", &neighborSummary.StaticSync})
    neighborSummary.EntityData.Children.Append("dynamic-sync", types.YChild{"DynamicSync", &neighborSummary.DynamicSync})
    neighborSummary.EntityData.Leafs = types.NewOrderedMap()
    neighborSummary.EntityData.Leafs.Append("total-neighbor-entries", types.YLeaf{"TotalNeighborEntries", neighborSummary.TotalNeighborEntries})

    neighborSummary.EntityData.YListKeys = []string {}

    return &(neighborSummary.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast
// Multicast neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ipv6 neighbhors count which are in INCMP state. The type is
    // interface{} with range: 0..4294967295.
    IncompleteEntries interface{}

    // Total ipv6 neighbhors count which are in REACH state. The type is
    // interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total ipv6 neighbhors count which are STALE. The type is interface{} with
    // range: 0..4294967295.
    StaleEntries interface{}

    // Total ipv6 neighbhors count which are in DELAY state. The type is
    // interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total ipv6 neighbhors count which are in PROBE state. The type is
    // interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total ipv6 neighbhors count which are in DELETE state. The type is
    // interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetEntityData() *types.CommonEntityData {
    multicast.EntityData.YFilter = multicast.YFilter
    multicast.EntityData.YangName = "multicast"
    multicast.EntityData.BundleName = "cisco_ios_xr"
    multicast.EntityData.ParentYangName = "neighbor-summary"
    multicast.EntityData.SegmentPath = "multicast"
    multicast.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-summary/" + multicast.EntityData.SegmentPath
    multicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicast.EntityData.Children = types.NewOrderedMap()
    multicast.EntityData.Leafs = types.NewOrderedMap()
    multicast.EntityData.Leafs.Append("incomplete-entries", types.YLeaf{"IncompleteEntries", multicast.IncompleteEntries})
    multicast.EntityData.Leafs.Append("reachable-entries", types.YLeaf{"ReachableEntries", multicast.ReachableEntries})
    multicast.EntityData.Leafs.Append("stale-entries", types.YLeaf{"StaleEntries", multicast.StaleEntries})
    multicast.EntityData.Leafs.Append("delayed-entries", types.YLeaf{"DelayedEntries", multicast.DelayedEntries})
    multicast.EntityData.Leafs.Append("probe-entries", types.YLeaf{"ProbeEntries", multicast.ProbeEntries})
    multicast.EntityData.Leafs.Append("deleted-entries", types.YLeaf{"DeletedEntries", multicast.DeletedEntries})
    multicast.EntityData.Leafs.Append("subtotal-neighbor-entries", types.YLeaf{"SubtotalNeighborEntries", multicast.SubtotalNeighborEntries})

    multicast.EntityData.YListKeys = []string {}

    return &(multicast.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static
// Static neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ipv6 neighbhors count which are in INCMP state. The type is
    // interface{} with range: 0..4294967295.
    IncompleteEntries interface{}

    // Total ipv6 neighbhors count which are in REACH state. The type is
    // interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total ipv6 neighbhors count which are STALE. The type is interface{} with
    // range: 0..4294967295.
    StaleEntries interface{}

    // Total ipv6 neighbhors count which are in DELAY state. The type is
    // interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total ipv6 neighbhors count which are in PROBE state. The type is
    // interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total ipv6 neighbhors count which are in DELETE state. The type is
    // interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xr"
    static.EntityData.ParentYangName = "neighbor-summary"
    static.EntityData.SegmentPath = "static"
    static.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-summary/" + static.EntityData.SegmentPath
    static.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    static.EntityData.Children = types.NewOrderedMap()
    static.EntityData.Leafs = types.NewOrderedMap()
    static.EntityData.Leafs.Append("incomplete-entries", types.YLeaf{"IncompleteEntries", static.IncompleteEntries})
    static.EntityData.Leafs.Append("reachable-entries", types.YLeaf{"ReachableEntries", static.ReachableEntries})
    static.EntityData.Leafs.Append("stale-entries", types.YLeaf{"StaleEntries", static.StaleEntries})
    static.EntityData.Leafs.Append("delayed-entries", types.YLeaf{"DelayedEntries", static.DelayedEntries})
    static.EntityData.Leafs.Append("probe-entries", types.YLeaf{"ProbeEntries", static.ProbeEntries})
    static.EntityData.Leafs.Append("deleted-entries", types.YLeaf{"DeletedEntries", static.DeletedEntries})
    static.EntityData.Leafs.Append("subtotal-neighbor-entries", types.YLeaf{"SubtotalNeighborEntries", static.SubtotalNeighborEntries})

    static.EntityData.YListKeys = []string {}

    return &(static.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic
// Dynamic neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ipv6 neighbhors count which are in INCMP state. The type is
    // interface{} with range: 0..4294967295.
    IncompleteEntries interface{}

    // Total ipv6 neighbhors count which are in REACH state. The type is
    // interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total ipv6 neighbhors count which are STALE. The type is interface{} with
    // range: 0..4294967295.
    StaleEntries interface{}

    // Total ipv6 neighbhors count which are in DELAY state. The type is
    // interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total ipv6 neighbhors count which are in PROBE state. The type is
    // interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total ipv6 neighbhors count which are in DELETE state. The type is
    // interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetEntityData() *types.CommonEntityData {
    dynamic.EntityData.YFilter = dynamic.YFilter
    dynamic.EntityData.YangName = "dynamic"
    dynamic.EntityData.BundleName = "cisco_ios_xr"
    dynamic.EntityData.ParentYangName = "neighbor-summary"
    dynamic.EntityData.SegmentPath = "dynamic"
    dynamic.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-summary/" + dynamic.EntityData.SegmentPath
    dynamic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamic.EntityData.Children = types.NewOrderedMap()
    dynamic.EntityData.Leafs = types.NewOrderedMap()
    dynamic.EntityData.Leafs.Append("incomplete-entries", types.YLeaf{"IncompleteEntries", dynamic.IncompleteEntries})
    dynamic.EntityData.Leafs.Append("reachable-entries", types.YLeaf{"ReachableEntries", dynamic.ReachableEntries})
    dynamic.EntityData.Leafs.Append("stale-entries", types.YLeaf{"StaleEntries", dynamic.StaleEntries})
    dynamic.EntityData.Leafs.Append("delayed-entries", types.YLeaf{"DelayedEntries", dynamic.DelayedEntries})
    dynamic.EntityData.Leafs.Append("probe-entries", types.YLeaf{"ProbeEntries", dynamic.ProbeEntries})
    dynamic.EntityData.Leafs.Append("deleted-entries", types.YLeaf{"DeletedEntries", dynamic.DeletedEntries})
    dynamic.EntityData.Leafs.Append("subtotal-neighbor-entries", types.YLeaf{"SubtotalNeighborEntries", dynamic.SubtotalNeighborEntries})

    dynamic.EntityData.YListKeys = []string {}

    return &(dynamic.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Sync
// Sync neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Sync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ipv6 neighbhors count which are in INCMP state. The type is
    // interface{} with range: 0..4294967295.
    IncompleteEntries interface{}

    // Total ipv6 neighbhors count which are in REACH state. The type is
    // interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total ipv6 neighbhors count which are STALE. The type is interface{} with
    // range: 0..4294967295.
    StaleEntries interface{}

    // Total ipv6 neighbhors count which are in DELAY state. The type is
    // interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total ipv6 neighbhors count which are in PROBE state. The type is
    // interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total ipv6 neighbhors count which are in DELETE state. The type is
    // interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (sync *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Sync) GetEntityData() *types.CommonEntityData {
    sync.EntityData.YFilter = sync.YFilter
    sync.EntityData.YangName = "sync"
    sync.EntityData.BundleName = "cisco_ios_xr"
    sync.EntityData.ParentYangName = "neighbor-summary"
    sync.EntityData.SegmentPath = "sync"
    sync.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-summary/" + sync.EntityData.SegmentPath
    sync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sync.EntityData.Children = types.NewOrderedMap()
    sync.EntityData.Leafs = types.NewOrderedMap()
    sync.EntityData.Leafs.Append("incomplete-entries", types.YLeaf{"IncompleteEntries", sync.IncompleteEntries})
    sync.EntityData.Leafs.Append("reachable-entries", types.YLeaf{"ReachableEntries", sync.ReachableEntries})
    sync.EntityData.Leafs.Append("stale-entries", types.YLeaf{"StaleEntries", sync.StaleEntries})
    sync.EntityData.Leafs.Append("delayed-entries", types.YLeaf{"DelayedEntries", sync.DelayedEntries})
    sync.EntityData.Leafs.Append("probe-entries", types.YLeaf{"ProbeEntries", sync.ProbeEntries})
    sync.EntityData.Leafs.Append("deleted-entries", types.YLeaf{"DeletedEntries", sync.DeletedEntries})
    sync.EntityData.Leafs.Append("subtotal-neighbor-entries", types.YLeaf{"SubtotalNeighborEntries", sync.SubtotalNeighborEntries})

    sync.EntityData.YListKeys = []string {}

    return &(sync.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_StaticSync
// StaticSync neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_StaticSync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ipv6 neighbhors count which are in INCMP state. The type is
    // interface{} with range: 0..4294967295.
    IncompleteEntries interface{}

    // Total ipv6 neighbhors count which are in REACH state. The type is
    // interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total ipv6 neighbhors count which are STALE. The type is interface{} with
    // range: 0..4294967295.
    StaleEntries interface{}

    // Total ipv6 neighbhors count which are in DELAY state. The type is
    // interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total ipv6 neighbhors count which are in PROBE state. The type is
    // interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total ipv6 neighbhors count which are in DELETE state. The type is
    // interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (staticSync *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_StaticSync) GetEntityData() *types.CommonEntityData {
    staticSync.EntityData.YFilter = staticSync.YFilter
    staticSync.EntityData.YangName = "static-sync"
    staticSync.EntityData.BundleName = "cisco_ios_xr"
    staticSync.EntityData.ParentYangName = "neighbor-summary"
    staticSync.EntityData.SegmentPath = "static-sync"
    staticSync.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-summary/" + staticSync.EntityData.SegmentPath
    staticSync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticSync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticSync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticSync.EntityData.Children = types.NewOrderedMap()
    staticSync.EntityData.Leafs = types.NewOrderedMap()
    staticSync.EntityData.Leafs.Append("incomplete-entries", types.YLeaf{"IncompleteEntries", staticSync.IncompleteEntries})
    staticSync.EntityData.Leafs.Append("reachable-entries", types.YLeaf{"ReachableEntries", staticSync.ReachableEntries})
    staticSync.EntityData.Leafs.Append("stale-entries", types.YLeaf{"StaleEntries", staticSync.StaleEntries})
    staticSync.EntityData.Leafs.Append("delayed-entries", types.YLeaf{"DelayedEntries", staticSync.DelayedEntries})
    staticSync.EntityData.Leafs.Append("probe-entries", types.YLeaf{"ProbeEntries", staticSync.ProbeEntries})
    staticSync.EntityData.Leafs.Append("deleted-entries", types.YLeaf{"DeletedEntries", staticSync.DeletedEntries})
    staticSync.EntityData.Leafs.Append("subtotal-neighbor-entries", types.YLeaf{"SubtotalNeighborEntries", staticSync.SubtotalNeighborEntries})

    staticSync.EntityData.YListKeys = []string {}

    return &(staticSync.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_DynamicSync
// DynamicSync neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_DynamicSync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ipv6 neighbhors count which are in INCMP state. The type is
    // interface{} with range: 0..4294967295.
    IncompleteEntries interface{}

    // Total ipv6 neighbhors count which are in REACH state. The type is
    // interface{} with range: 0..4294967295.
    ReachableEntries interface{}

    // Total ipv6 neighbhors count which are STALE. The type is interface{} with
    // range: 0..4294967295.
    StaleEntries interface{}

    // Total ipv6 neighbhors count which are in DELAY state. The type is
    // interface{} with range: 0..4294967295.
    DelayedEntries interface{}

    // Total ipv6 neighbhors count which are in PROBE state. The type is
    // interface{} with range: 0..4294967295.
    ProbeEntries interface{}

    // Total ipv6 neighbhors count which are in DELETE state. The type is
    // interface{} with range: 0..4294967295.
    DeletedEntries interface{}

    // Total number of entries. The type is interface{} with range: 0..4294967295.
    SubtotalNeighborEntries interface{}
}

func (dynamicSync *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_DynamicSync) GetEntityData() *types.CommonEntityData {
    dynamicSync.EntityData.YFilter = dynamicSync.YFilter
    dynamicSync.EntityData.YangName = "dynamic-sync"
    dynamicSync.EntityData.BundleName = "cisco_ios_xr"
    dynamicSync.EntityData.ParentYangName = "neighbor-summary"
    dynamicSync.EntityData.SegmentPath = "dynamic-sync"
    dynamicSync.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/neighbor-summary/" + dynamicSync.EntityData.SegmentPath
    dynamicSync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicSync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicSync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicSync.EntityData.Children = types.NewOrderedMap()
    dynamicSync.EntityData.Leafs = types.NewOrderedMap()
    dynamicSync.EntityData.Leafs.Append("incomplete-entries", types.YLeaf{"IncompleteEntries", dynamicSync.IncompleteEntries})
    dynamicSync.EntityData.Leafs.Append("reachable-entries", types.YLeaf{"ReachableEntries", dynamicSync.ReachableEntries})
    dynamicSync.EntityData.Leafs.Append("stale-entries", types.YLeaf{"StaleEntries", dynamicSync.StaleEntries})
    dynamicSync.EntityData.Leafs.Append("delayed-entries", types.YLeaf{"DelayedEntries", dynamicSync.DelayedEntries})
    dynamicSync.EntityData.Leafs.Append("probe-entries", types.YLeaf{"ProbeEntries", dynamicSync.ProbeEntries})
    dynamicSync.EntityData.Leafs.Append("deleted-entries", types.YLeaf{"DeletedEntries", dynamicSync.DeletedEntries})
    dynamicSync.EntityData.Leafs.Append("subtotal-neighbor-entries", types.YLeaf{"SubtotalNeighborEntries", dynamicSync.SubtotalNeighborEntries})

    dynamicSync.EntityData.YListKeys = []string {}

    return &(dynamicSync.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes
// IPv6 ND list of bundle nodes for a specific
// node
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific bundle node. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode.
    BundleNode []*Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetEntityData() *types.CommonEntityData {
    bundleNodes.EntityData.YFilter = bundleNodes.YFilter
    bundleNodes.EntityData.YangName = "bundle-nodes"
    bundleNodes.EntityData.BundleName = "cisco_ios_xr"
    bundleNodes.EntityData.ParentYangName = "node"
    bundleNodes.EntityData.SegmentPath = "bundle-nodes"
    bundleNodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + bundleNodes.EntityData.SegmentPath
    bundleNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleNodes.EntityData.Children = types.NewOrderedMap()
    bundleNodes.EntityData.Children.Append("bundle-node", types.YChild{"BundleNode", nil})
    for i := range bundleNodes.BundleNode {
        bundleNodes.EntityData.Children.Append(types.GetSegmentPath(bundleNodes.BundleNode[i]), types.YChild{"BundleNode", bundleNodes.BundleNode[i]})
    }
    bundleNodes.EntityData.Leafs = types.NewOrderedMap()

    bundleNodes.EntityData.YListKeys = []string {}

    return &(bundleNodes.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode
// IPv6 ND operational data for a specific
// bundle node
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The bundle node name. The type is string with
    // pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // Name of the process. The type is string.
    ProcessName interface{}

    // Sent sequence number for error detection. The type is interface{} with
    // range: 0..4294967295.
    SentSequenceNumber interface{}

    // Received sequence num for error detection. The type is interface{} with
    // range: 0..4294967295.
    ReceivedSequenceNumber interface{}

    // State. The type is Ipv6NdBndlState.
    State interface{}

    // change of state. The type is interface{} with range: 0..4294967295.
    StateChanges interface{}

    // Total packet sends. The type is interface{} with range: 0..4294967295.
    SentPackets interface{}

    // Total packet receives. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Uptime of node (secs).
    Age Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age
}

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetEntityData() *types.CommonEntityData {
    bundleNode.EntityData.YFilter = bundleNode.YFilter
    bundleNode.EntityData.YangName = "bundle-node"
    bundleNode.EntityData.BundleName = "cisco_ios_xr"
    bundleNode.EntityData.ParentYangName = "bundle-nodes"
    bundleNode.EntityData.SegmentPath = "bundle-node" + types.AddKeyToken(bundleNode.NodeName, "node-name")
    bundleNode.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-nodes/" + bundleNode.EntityData.SegmentPath
    bundleNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleNode.EntityData.Children = types.NewOrderedMap()
    bundleNode.EntityData.Children.Append("age", types.YChild{"Age", &bundleNode.Age})
    bundleNode.EntityData.Leafs = types.NewOrderedMap()
    bundleNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", bundleNode.NodeName})
    bundleNode.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", bundleNode.GroupId})
    bundleNode.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", bundleNode.ProcessName})
    bundleNode.EntityData.Leafs.Append("sent-sequence-number", types.YLeaf{"SentSequenceNumber", bundleNode.SentSequenceNumber})
    bundleNode.EntityData.Leafs.Append("received-sequence-number", types.YLeaf{"ReceivedSequenceNumber", bundleNode.ReceivedSequenceNumber})
    bundleNode.EntityData.Leafs.Append("state", types.YLeaf{"State", bundleNode.State})
    bundleNode.EntityData.Leafs.Append("state-changes", types.YLeaf{"StateChanges", bundleNode.StateChanges})
    bundleNode.EntityData.Leafs.Append("sent-packets", types.YLeaf{"SentPackets", bundleNode.SentPackets})
    bundleNode.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", bundleNode.ReceivedPackets})

    bundleNode.EntityData.YListKeys = []string {"NodeName"}

    return &(bundleNode.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age
// Uptime of node (secs)
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (age *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode_Age) GetEntityData() *types.CommonEntityData {
    age.EntityData.YFilter = age.YFilter
    age.EntityData.YangName = "age"
    age.EntityData.BundleName = "cisco_ios_xr"
    age.EntityData.ParentYangName = "bundle-node"
    age.EntityData.SegmentPath = "age"
    age.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-nodes/bundle-node/" + age.EntityData.SegmentPath
    age.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    age.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    age.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    age.EntityData.Children = types.NewOrderedMap()
    age.EntityData.Leafs = types.NewOrderedMap()
    age.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", age.Seconds})

    age.EntityData.YListKeys = []string {}

    return &(age.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces
// IPv6 ND list of bundle interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific bundler interface. The type is
    // slice of Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface.
    BundleInterface []*Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetEntityData() *types.CommonEntityData {
    bundleInterfaces.EntityData.YFilter = bundleInterfaces.YFilter
    bundleInterfaces.EntityData.YangName = "bundle-interfaces"
    bundleInterfaces.EntityData.BundleName = "cisco_ios_xr"
    bundleInterfaces.EntityData.ParentYangName = "node"
    bundleInterfaces.EntityData.SegmentPath = "bundle-interfaces"
    bundleInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + bundleInterfaces.EntityData.SegmentPath
    bundleInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterfaces.EntityData.Children = types.NewOrderedMap()
    bundleInterfaces.EntityData.Children.Append("bundle-interface", types.YChild{"BundleInterface", nil})
    for i := range bundleInterfaces.BundleInterface {
        bundleInterfaces.EntityData.Children.Append(types.GetSegmentPath(bundleInterfaces.BundleInterface[i]), types.YChild{"BundleInterface", bundleInterfaces.BundleInterface[i]})
    }
    bundleInterfaces.EntityData.Leafs = types.NewOrderedMap()

    bundleInterfaces.EntityData.YListKeys = []string {}

    return &(bundleInterfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface
// IPv6 ND operational data for a specific
// bundler interface
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Name of the Parent interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    ParentInterfaceName interface{}

    // Interface type. The type is interface{} with range: 0..4294967295.
    Iftype interface{}

    // MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // etype field. The type is interface{} with range: 0..4294967295.
    Etype interface{}

    // vlan tag/id/ucv. The type is interface{} with range: 0..65535.
    VlanTag interface{}

    // size of mac address. The type is interface{} with range: 0..4294967295.
    MacAddrSize interface{}

    // media access control address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    // IPV6 Link local address.
    LocalAddress Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress

    // List of ND global addresses. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress.
    GlobalAddress []*Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress

    // List of member nodes. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode.
    MemberNode []*Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetEntityData() *types.CommonEntityData {
    bundleInterface.EntityData.YFilter = bundleInterface.YFilter
    bundleInterface.EntityData.YangName = "bundle-interface"
    bundleInterface.EntityData.BundleName = "cisco_ios_xr"
    bundleInterface.EntityData.ParentYangName = "bundle-interfaces"
    bundleInterface.EntityData.SegmentPath = "bundle-interface" + types.AddKeyToken(bundleInterface.InterfaceName, "interface-name")
    bundleInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-interfaces/" + bundleInterface.EntityData.SegmentPath
    bundleInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterface.EntityData.Children = types.NewOrderedMap()
    bundleInterface.EntityData.Children.Append("nd-parameters", types.YChild{"NdParameters", &bundleInterface.NdParameters})
    bundleInterface.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &bundleInterface.LocalAddress})
    bundleInterface.EntityData.Children.Append("global-address", types.YChild{"GlobalAddress", nil})
    for i := range bundleInterface.GlobalAddress {
        types.SetYListKey(bundleInterface.GlobalAddress[i], i)
        bundleInterface.EntityData.Children.Append(types.GetSegmentPath(bundleInterface.GlobalAddress[i]), types.YChild{"GlobalAddress", bundleInterface.GlobalAddress[i]})
    }
    bundleInterface.EntityData.Children.Append("member-node", types.YChild{"MemberNode", nil})
    for i := range bundleInterface.MemberNode {
        types.SetYListKey(bundleInterface.MemberNode[i], i)
        bundleInterface.EntityData.Children.Append(types.GetSegmentPath(bundleInterface.MemberNode[i]), types.YChild{"MemberNode", bundleInterface.MemberNode[i]})
    }
    bundleInterface.EntityData.Leafs = types.NewOrderedMap()
    bundleInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", bundleInterface.InterfaceName})
    bundleInterface.EntityData.Leafs.Append("parent-interface-name", types.YLeaf{"ParentInterfaceName", bundleInterface.ParentInterfaceName})
    bundleInterface.EntityData.Leafs.Append("iftype", types.YLeaf{"Iftype", bundleInterface.Iftype})
    bundleInterface.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", bundleInterface.Mtu})
    bundleInterface.EntityData.Leafs.Append("etype", types.YLeaf{"Etype", bundleInterface.Etype})
    bundleInterface.EntityData.Leafs.Append("vlan-tag", types.YLeaf{"VlanTag", bundleInterface.VlanTag})
    bundleInterface.EntityData.Leafs.Append("mac-addr-size", types.YLeaf{"MacAddrSize", bundleInterface.MacAddrSize})
    bundleInterface.EntityData.Leafs.Append("mac-addr", types.YLeaf{"MacAddr", bundleInterface.MacAddr})
    bundleInterface.EntityData.Leafs.Append("is-interface-enabled", types.YLeaf{"IsInterfaceEnabled", bundleInterface.IsInterfaceEnabled})
    bundleInterface.EntityData.Leafs.Append("is-ipv6-enabled", types.YLeaf{"IsIpv6Enabled", bundleInterface.IsIpv6Enabled})
    bundleInterface.EntityData.Leafs.Append("is-mpls-enabled", types.YLeaf{"IsMplsEnabled", bundleInterface.IsMplsEnabled})
    bundleInterface.EntityData.Leafs.Append("member-link", types.YLeaf{"MemberLink", bundleInterface.MemberLink})

    bundleInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(bundleInterface.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters
// ND interface parameters
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters struct {
    EntityData types.CommonEntityData
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

func (ndParameters *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_NdParameters) GetEntityData() *types.CommonEntityData {
    ndParameters.EntityData.YFilter = ndParameters.YFilter
    ndParameters.EntityData.YangName = "nd-parameters"
    ndParameters.EntityData.BundleName = "cisco_ios_xr"
    ndParameters.EntityData.ParentYangName = "bundle-interface"
    ndParameters.EntityData.SegmentPath = "nd-parameters"
    ndParameters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-interfaces/bundle-interface/" + ndParameters.EntityData.SegmentPath
    ndParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ndParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ndParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ndParameters.EntityData.Children = types.NewOrderedMap()
    ndParameters.EntityData.Leafs = types.NewOrderedMap()
    ndParameters.EntityData.Leafs.Append("is-dad-enabled", types.YLeaf{"IsDadEnabled", ndParameters.IsDadEnabled})
    ndParameters.EntityData.Leafs.Append("dad-attempts", types.YLeaf{"DadAttempts", ndParameters.DadAttempts})
    ndParameters.EntityData.Leafs.Append("is-icm-pv6-redirect", types.YLeaf{"IsIcmPv6Redirect", ndParameters.IsIcmPv6Redirect})
    ndParameters.EntityData.Leafs.Append("is-dhcp-managed", types.YLeaf{"IsDhcpManaged", ndParameters.IsDhcpManaged})
    ndParameters.EntityData.Leafs.Append("is-route-address-managed", types.YLeaf{"IsRouteAddressManaged", ndParameters.IsRouteAddressManaged})
    ndParameters.EntityData.Leafs.Append("is-suppressed", types.YLeaf{"IsSuppressed", ndParameters.IsSuppressed})
    ndParameters.EntityData.Leafs.Append("send-unicast-ra", types.YLeaf{"SendUnicastRa", ndParameters.SendUnicastRa})
    ndParameters.EntityData.Leafs.Append("nd-retransmit-interval", types.YLeaf{"NdRetransmitInterval", ndParameters.NdRetransmitInterval})
    ndParameters.EntityData.Leafs.Append("nd-min-transmit-interval", types.YLeaf{"NdMinTransmitInterval", ndParameters.NdMinTransmitInterval})
    ndParameters.EntityData.Leafs.Append("nd-max-transmit-interval", types.YLeaf{"NdMaxTransmitInterval", ndParameters.NdMaxTransmitInterval})
    ndParameters.EntityData.Leafs.Append("nd-advertisement-lifetime", types.YLeaf{"NdAdvertisementLifetime", ndParameters.NdAdvertisementLifetime})
    ndParameters.EntityData.Leafs.Append("nd-reachable-time", types.YLeaf{"NdReachableTime", ndParameters.NdReachableTime})
    ndParameters.EntityData.Leafs.Append("nd-cache-limit", types.YLeaf{"NdCacheLimit", ndParameters.NdCacheLimit})
    ndParameters.EntityData.Leafs.Append("complete-protocol-count", types.YLeaf{"CompleteProtocolCount", ndParameters.CompleteProtocolCount})
    ndParameters.EntityData.Leafs.Append("complete-glean-count", types.YLeaf{"CompleteGleanCount", ndParameters.CompleteGleanCount})
    ndParameters.EntityData.Leafs.Append("incomplete-protocol-count", types.YLeaf{"IncompleteProtocolCount", ndParameters.IncompleteProtocolCount})
    ndParameters.EntityData.Leafs.Append("incomplete-glean-count", types.YLeaf{"IncompleteGleanCount", ndParameters.IncompleteGleanCount})
    ndParameters.EntityData.Leafs.Append("dropped-protocol-req-count", types.YLeaf{"DroppedProtocolReqCount", ndParameters.DroppedProtocolReqCount})
    ndParameters.EntityData.Leafs.Append("dropped-glean-req-count", types.YLeaf{"DroppedGleanReqCount", ndParameters.DroppedGleanReqCount})

    ndParameters.EntityData.YListKeys = []string {}

    return &(ndParameters.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress
// IPV6 Link local address
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of type IPV6. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // Valid lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    ValidLifetime interface{}

    // Preffered lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    PrefLifetime interface{}

    // IPV6 Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Address flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "bundle-interface"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-interfaces/bundle-interface/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", localAddress.Ipv6Address})
    localAddress.EntityData.Leafs.Append("valid-lifetime", types.YLeaf{"ValidLifetime", localAddress.ValidLifetime})
    localAddress.EntityData.Leafs.Append("pref-lifetime", types.YLeaf{"PrefLifetime", localAddress.PrefLifetime})
    localAddress.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", localAddress.PrefixLength})
    localAddress.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", localAddress.Flags})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress
// List of ND global addresses
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Address of type IPV6. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // Valid lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    ValidLifetime interface{}

    // Preffered lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    PrefLifetime interface{}

    // IPV6 Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Address flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetEntityData() *types.CommonEntityData {
    globalAddress.EntityData.YFilter = globalAddress.YFilter
    globalAddress.EntityData.YangName = "global-address"
    globalAddress.EntityData.BundleName = "cisco_ios_xr"
    globalAddress.EntityData.ParentYangName = "bundle-interface"
    globalAddress.EntityData.SegmentPath = "global-address" + types.AddNoKeyToken(globalAddress)
    globalAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-interfaces/bundle-interface/" + globalAddress.EntityData.SegmentPath
    globalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAddress.EntityData.Children = types.NewOrderedMap()
    globalAddress.EntityData.Leafs = types.NewOrderedMap()
    globalAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", globalAddress.Ipv6Address})
    globalAddress.EntityData.Leafs.Append("valid-lifetime", types.YLeaf{"ValidLifetime", globalAddress.ValidLifetime})
    globalAddress.EntityData.Leafs.Append("pref-lifetime", types.YLeaf{"PrefLifetime", globalAddress.PrefLifetime})
    globalAddress.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", globalAddress.PrefixLength})
    globalAddress.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", globalAddress.Flags})

    globalAddress.EntityData.YListKeys = []string {}

    return &(globalAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode
// List of member nodes
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Node Name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Number of links on the node. The type is interface{} with range:
    // 0..4294967295.
    TotalLinks interface{}
}

func (memberNode *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode) GetEntityData() *types.CommonEntityData {
    memberNode.EntityData.YFilter = memberNode.YFilter
    memberNode.EntityData.YangName = "member-node"
    memberNode.EntityData.BundleName = "cisco_ios_xr"
    memberNode.EntityData.ParentYangName = "bundle-interface"
    memberNode.EntityData.SegmentPath = "member-node" + types.AddNoKeyToken(memberNode)
    memberNode.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/bundle-interfaces/bundle-interface/" + memberNode.EntityData.SegmentPath
    memberNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberNode.EntityData.Children = types.NewOrderedMap()
    memberNode.EntityData.Leafs = types.NewOrderedMap()
    memberNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", memberNode.NodeName})
    memberNode.EntityData.Leafs.Append("total-links", types.YLeaf{"TotalLinks", memberNode.TotalLinks})

    memberNode.EntityData.YListKeys = []string {}

    return &(memberNode.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_Interfaces
// IPv6 node discovery list of interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6  node discovery operational data for a specific node and interface.
    // The type is slice of Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface.
    Interface []*Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface
// IPv6  node discovery operational data for a
// specific node and interface
type Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
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

func (self *Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("is-dad-enabled", types.YLeaf{"IsDadEnabled", self.IsDadEnabled})
    self.EntityData.Leafs.Append("dad-attempts", types.YLeaf{"DadAttempts", self.DadAttempts})
    self.EntityData.Leafs.Append("is-icm-pv6-redirect", types.YLeaf{"IsIcmPv6Redirect", self.IsIcmPv6Redirect})
    self.EntityData.Leafs.Append("is-dhcp-managed", types.YLeaf{"IsDhcpManaged", self.IsDhcpManaged})
    self.EntityData.Leafs.Append("is-route-address-managed", types.YLeaf{"IsRouteAddressManaged", self.IsRouteAddressManaged})
    self.EntityData.Leafs.Append("is-suppressed", types.YLeaf{"IsSuppressed", self.IsSuppressed})
    self.EntityData.Leafs.Append("send-unicast-ra", types.YLeaf{"SendUnicastRa", self.SendUnicastRa})
    self.EntityData.Leafs.Append("nd-retransmit-interval", types.YLeaf{"NdRetransmitInterval", self.NdRetransmitInterval})
    self.EntityData.Leafs.Append("nd-min-transmit-interval", types.YLeaf{"NdMinTransmitInterval", self.NdMinTransmitInterval})
    self.EntityData.Leafs.Append("nd-max-transmit-interval", types.YLeaf{"NdMaxTransmitInterval", self.NdMaxTransmitInterval})
    self.EntityData.Leafs.Append("nd-advertisement-lifetime", types.YLeaf{"NdAdvertisementLifetime", self.NdAdvertisementLifetime})
    self.EntityData.Leafs.Append("nd-reachable-time", types.YLeaf{"NdReachableTime", self.NdReachableTime})
    self.EntityData.Leafs.Append("nd-cache-limit", types.YLeaf{"NdCacheLimit", self.NdCacheLimit})
    self.EntityData.Leafs.Append("complete-protocol-count", types.YLeaf{"CompleteProtocolCount", self.CompleteProtocolCount})
    self.EntityData.Leafs.Append("complete-glean-count", types.YLeaf{"CompleteGleanCount", self.CompleteGleanCount})
    self.EntityData.Leafs.Append("incomplete-protocol-count", types.YLeaf{"IncompleteProtocolCount", self.IncompleteProtocolCount})
    self.EntityData.Leafs.Append("incomplete-glean-count", types.YLeaf{"IncompleteGleanCount", self.IncompleteGleanCount})
    self.EntityData.Leafs.Append("dropped-protocol-req-count", types.YLeaf{"DroppedProtocolReqCount", self.DroppedProtocolReqCount})
    self.EntityData.Leafs.Append("dropped-glean-req-count", types.YLeaf{"DroppedGleanReqCount", self.DroppedGleanReqCount})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters
// IPv6 ND virtual router information for a
// specific interface
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ND virtual  router operational data for a specific interface. The type
    // is slice of Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter.
    NdVirtualRouter []*Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetEntityData() *types.CommonEntityData {
    ndVirtualRouters.EntityData.YFilter = ndVirtualRouters.YFilter
    ndVirtualRouters.EntityData.YangName = "nd-virtual-routers"
    ndVirtualRouters.EntityData.BundleName = "cisco_ios_xr"
    ndVirtualRouters.EntityData.ParentYangName = "node"
    ndVirtualRouters.EntityData.SegmentPath = "nd-virtual-routers"
    ndVirtualRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + ndVirtualRouters.EntityData.SegmentPath
    ndVirtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ndVirtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ndVirtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ndVirtualRouters.EntityData.Children = types.NewOrderedMap()
    ndVirtualRouters.EntityData.Children.Append("nd-virtual-router", types.YChild{"NdVirtualRouter", nil})
    for i := range ndVirtualRouters.NdVirtualRouter {
        ndVirtualRouters.EntityData.Children.Append(types.GetSegmentPath(ndVirtualRouters.NdVirtualRouter[i]), types.YChild{"NdVirtualRouter", ndVirtualRouters.NdVirtualRouter[i]})
    }
    ndVirtualRouters.EntityData.Leafs = types.NewOrderedMap()

    ndVirtualRouters.EntityData.YListKeys = []string {}

    return &(ndVirtualRouters.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter
// IPv6 ND virtual  router operational data for
// a specific interface
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // IPV6 Link-Layer Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    // IPV6 Link local address.
    LocalAddress Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress

    // List of ND global addresses. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress.
    VrGlobalAddress []*Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetEntityData() *types.CommonEntityData {
    ndVirtualRouter.EntityData.YFilter = ndVirtualRouter.YFilter
    ndVirtualRouter.EntityData.YangName = "nd-virtual-router"
    ndVirtualRouter.EntityData.BundleName = "cisco_ios_xr"
    ndVirtualRouter.EntityData.ParentYangName = "nd-virtual-routers"
    ndVirtualRouter.EntityData.SegmentPath = "nd-virtual-router" + types.AddKeyToken(ndVirtualRouter.InterfaceName, "interface-name")
    ndVirtualRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/nd-virtual-routers/" + ndVirtualRouter.EntityData.SegmentPath
    ndVirtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ndVirtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ndVirtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ndVirtualRouter.EntityData.Children = types.NewOrderedMap()
    ndVirtualRouter.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &ndVirtualRouter.LocalAddress})
    ndVirtualRouter.EntityData.Children.Append("vr-global-address", types.YChild{"VrGlobalAddress", nil})
    for i := range ndVirtualRouter.VrGlobalAddress {
        types.SetYListKey(ndVirtualRouter.VrGlobalAddress[i], i)
        ndVirtualRouter.EntityData.Children.Append(types.GetSegmentPath(ndVirtualRouter.VrGlobalAddress[i]), types.YChild{"VrGlobalAddress", ndVirtualRouter.VrGlobalAddress[i]})
    }
    ndVirtualRouter.EntityData.Leafs = types.NewOrderedMap()
    ndVirtualRouter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ndVirtualRouter.InterfaceName})
    ndVirtualRouter.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", ndVirtualRouter.LinkLayerAddress})
    ndVirtualRouter.EntityData.Leafs.Append("context", types.YLeaf{"Context", ndVirtualRouter.Context})
    ndVirtualRouter.EntityData.Leafs.Append("state", types.YLeaf{"State", ndVirtualRouter.State})
    ndVirtualRouter.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", ndVirtualRouter.Flags})
    ndVirtualRouter.EntityData.Leafs.Append("vr-gl-addr-ct", types.YLeaf{"VrGlAddrCt", ndVirtualRouter.VrGlAddrCt})

    ndVirtualRouter.EntityData.YListKeys = []string {"InterfaceName"}

    return &(ndVirtualRouter.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress
// IPV6 Link local address
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of type IPV6. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // Valid lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    ValidLifetime interface{}

    // Preffered lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    PrefLifetime interface{}

    // IPV6 Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Address flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "nd-virtual-router"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/nd-virtual-routers/nd-virtual-router/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", localAddress.Ipv6Address})
    localAddress.EntityData.Leafs.Append("valid-lifetime", types.YLeaf{"ValidLifetime", localAddress.ValidLifetime})
    localAddress.EntityData.Leafs.Append("pref-lifetime", types.YLeaf{"PrefLifetime", localAddress.PrefLifetime})
    localAddress.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", localAddress.PrefixLength})
    localAddress.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", localAddress.Flags})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress
// List of ND global addresses
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Address of type IPV6. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // Valid lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    ValidLifetime interface{}

    // Preffered lifetime of a Prefix. The type is interface{} with range:
    // 0..4294967295.
    PrefLifetime interface{}

    // IPV6 Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Address flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetEntityData() *types.CommonEntityData {
    vrGlobalAddress.EntityData.YFilter = vrGlobalAddress.YFilter
    vrGlobalAddress.EntityData.YangName = "vr-global-address"
    vrGlobalAddress.EntityData.BundleName = "cisco_ios_xr"
    vrGlobalAddress.EntityData.ParentYangName = "nd-virtual-router"
    vrGlobalAddress.EntityData.SegmentPath = "vr-global-address" + types.AddNoKeyToken(vrGlobalAddress)
    vrGlobalAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/nd-virtual-routers/nd-virtual-router/" + vrGlobalAddress.EntityData.SegmentPath
    vrGlobalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrGlobalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrGlobalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrGlobalAddress.EntityData.Children = types.NewOrderedMap()
    vrGlobalAddress.EntityData.Leafs = types.NewOrderedMap()
    vrGlobalAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", vrGlobalAddress.Ipv6Address})
    vrGlobalAddress.EntityData.Leafs.Append("valid-lifetime", types.YLeaf{"ValidLifetime", vrGlobalAddress.ValidLifetime})
    vrGlobalAddress.EntityData.Leafs.Append("pref-lifetime", types.YLeaf{"PrefLifetime", vrGlobalAddress.PrefLifetime})
    vrGlobalAddress.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", vrGlobalAddress.PrefixLength})
    vrGlobalAddress.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", vrGlobalAddress.Flags})

    vrGlobalAddress.EntityData.YListKeys = []string {}

    return &(vrGlobalAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces
// IPv6 ND list of SLAAC MGMT interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific slaac interface. The type is slice
    // of Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface.
    SlaacInterface []*Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetEntityData() *types.CommonEntityData {
    slaacInterfaces.EntityData.YFilter = slaacInterfaces.YFilter
    slaacInterfaces.EntityData.YangName = "slaac-interfaces"
    slaacInterfaces.EntityData.BundleName = "cisco_ios_xr"
    slaacInterfaces.EntityData.ParentYangName = "node"
    slaacInterfaces.EntityData.SegmentPath = "slaac-interfaces"
    slaacInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/" + slaacInterfaces.EntityData.SegmentPath
    slaacInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaacInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaacInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaacInterfaces.EntityData.Children = types.NewOrderedMap()
    slaacInterfaces.EntityData.Children.Append("slaac-interface", types.YChild{"SlaacInterface", nil})
    for i := range slaacInterfaces.SlaacInterface {
        slaacInterfaces.EntityData.Children.Append(types.GetSegmentPath(slaacInterfaces.SlaacInterface[i]), types.YChild{"SlaacInterface", slaacInterfaces.SlaacInterface[i]})
    }
    slaacInterfaces.EntityData.Leafs = types.NewOrderedMap()

    slaacInterfaces.EntityData.YListKeys = []string {}

    return &(slaacInterfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface
// IPv6 ND operational data for a specific slaac
// interface
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // IPv6 ND operational data for a specific slaac interface.
    RouterAdvertDetail Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetEntityData() *types.CommonEntityData {
    slaacInterface.EntityData.YFilter = slaacInterface.YFilter
    slaacInterface.EntityData.YangName = "slaac-interface"
    slaacInterface.EntityData.BundleName = "cisco_ios_xr"
    slaacInterface.EntityData.ParentYangName = "slaac-interfaces"
    slaacInterface.EntityData.SegmentPath = "slaac-interface" + types.AddKeyToken(slaacInterface.InterfaceName, "interface-name")
    slaacInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/" + slaacInterface.EntityData.SegmentPath
    slaacInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaacInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaacInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaacInterface.EntityData.Children = types.NewOrderedMap()
    slaacInterface.EntityData.Children.Append("router-advert-detail", types.YChild{"RouterAdvertDetail", &slaacInterface.RouterAdvertDetail})
    slaacInterface.EntityData.Leafs = types.NewOrderedMap()
    slaacInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", slaacInterface.InterfaceName})

    slaacInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(slaacInterface.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail
// IPv6 ND operational data for a specific
// slaac interface
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface database. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Idb interface{}

    // slaac db. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra.
    Ra []*Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetEntityData() *types.CommonEntityData {
    routerAdvertDetail.EntityData.YFilter = routerAdvertDetail.YFilter
    routerAdvertDetail.EntityData.YangName = "router-advert-detail"
    routerAdvertDetail.EntityData.BundleName = "cisco_ios_xr"
    routerAdvertDetail.EntityData.ParentYangName = "slaac-interface"
    routerAdvertDetail.EntityData.SegmentPath = "router-advert-detail"
    routerAdvertDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/slaac-interface/" + routerAdvertDetail.EntityData.SegmentPath
    routerAdvertDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerAdvertDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerAdvertDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerAdvertDetail.EntityData.Children = types.NewOrderedMap()
    routerAdvertDetail.EntityData.Children.Append("ra", types.YChild{"Ra", nil})
    for i := range routerAdvertDetail.Ra {
        types.SetYListKey(routerAdvertDetail.Ra[i], i)
        routerAdvertDetail.EntityData.Children.Append(types.GetSegmentPath(routerAdvertDetail.Ra[i]), types.YChild{"Ra", routerAdvertDetail.Ra[i]})
    }
    routerAdvertDetail.EntityData.Leafs = types.NewOrderedMap()
    routerAdvertDetail.EntityData.Leafs.Append("idb", types.YLeaf{"Idb", routerAdvertDetail.Idb})

    routerAdvertDetail.EntityData.YListKeys = []string {}

    return &(routerAdvertDetail.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra
// slaac db
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // address of type IPV6. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // number of intermediate devices between source and destination. The type is
    // interface{} with range: 0..4294967295.
    Hops interface{}

    // RA flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // active time. The type is interface{} with range: 0..4294967295.
    LifeTime interface{}

    // maximum transmission unit. The type is interface{} with range:
    // 0..4294967295.
    Mtu interface{}

    // message having the error info. The type is bool.
    ErrMsg interface{}

    // virtual routing and forwarding id. The type is interface{} with range:
    // 0..4294967295.
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

    // common reachabletime.
    ReachableTime Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime

    // RA retransmit time.
    RetransTime Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime

    // Prefix Queue. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ.
    PrefixQ []*Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ
}

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetEntityData() *types.CommonEntityData {
    ra.EntityData.YFilter = ra.YFilter
    ra.EntityData.YangName = "ra"
    ra.EntityData.BundleName = "cisco_ios_xr"
    ra.EntityData.ParentYangName = "router-advert-detail"
    ra.EntityData.SegmentPath = "ra" + types.AddNoKeyToken(ra)
    ra.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/slaac-interface/router-advert-detail/" + ra.EntityData.SegmentPath
    ra.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ra.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ra.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ra.EntityData.Children = types.NewOrderedMap()
    ra.EntityData.Children.Append("elapsed-ra-time", types.YChild{"ElapsedRaTime", &ra.ElapsedRaTime})
    ra.EntityData.Children.Append("reachable-time", types.YChild{"ReachableTime", &ra.ReachableTime})
    ra.EntityData.Children.Append("retrans-time", types.YChild{"RetransTime", &ra.RetransTime})
    ra.EntityData.Children.Append("prefix-q", types.YChild{"PrefixQ", nil})
    for i := range ra.PrefixQ {
        types.SetYListKey(ra.PrefixQ[i], i)
        ra.EntityData.Children.Append(types.GetSegmentPath(ra.PrefixQ[i]), types.YChild{"PrefixQ", ra.PrefixQ[i]})
    }
    ra.EntityData.Leafs = types.NewOrderedMap()
    ra.EntityData.Leafs.Append("address", types.YLeaf{"Address", ra.Address})
    ra.EntityData.Leafs.Append("hops", types.YLeaf{"Hops", ra.Hops})
    ra.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", ra.Flags})
    ra.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", ra.LifeTime})
    ra.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ra.Mtu})
    ra.EntityData.Leafs.Append("err-msg", types.YLeaf{"ErrMsg", ra.ErrMsg})
    ra.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", ra.VrfId})
    ra.EntityData.Leafs.Append("u6-tbl-id", types.YLeaf{"U6TblId", ra.U6TblId})
    ra.EntityData.Leafs.Append("rib-protoid", types.YLeaf{"RibProtoid", ra.RibProtoid})
    ra.EntityData.Leafs.Append("default-router", types.YLeaf{"DefaultRouter", ra.DefaultRouter})
    ra.EntityData.Leafs.Append("reachability", types.YLeaf{"Reachability", ra.Reachability})

    ra.EntityData.YListKeys = []string {}

    return &(ra.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime
// elapsedRATime
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (elapsedRaTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ElapsedRaTime) GetEntityData() *types.CommonEntityData {
    elapsedRaTime.EntityData.YFilter = elapsedRaTime.YFilter
    elapsedRaTime.EntityData.YangName = "elapsed-ra-time"
    elapsedRaTime.EntityData.BundleName = "cisco_ios_xr"
    elapsedRaTime.EntityData.ParentYangName = "ra"
    elapsedRaTime.EntityData.SegmentPath = "elapsed-ra-time"
    elapsedRaTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/slaac-interface/router-advert-detail/ra/" + elapsedRaTime.EntityData.SegmentPath
    elapsedRaTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    elapsedRaTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    elapsedRaTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    elapsedRaTime.EntityData.Children = types.NewOrderedMap()
    elapsedRaTime.EntityData.Leafs = types.NewOrderedMap()
    elapsedRaTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", elapsedRaTime.Seconds})

    elapsedRaTime.EntityData.YListKeys = []string {}

    return &(elapsedRaTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime
// common reachabletime
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (reachableTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime) GetEntityData() *types.CommonEntityData {
    reachableTime.EntityData.YFilter = reachableTime.YFilter
    reachableTime.EntityData.YangName = "reachable-time"
    reachableTime.EntityData.BundleName = "cisco_ios_xr"
    reachableTime.EntityData.ParentYangName = "ra"
    reachableTime.EntityData.SegmentPath = "reachable-time"
    reachableTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/slaac-interface/router-advert-detail/ra/" + reachableTime.EntityData.SegmentPath
    reachableTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reachableTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reachableTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reachableTime.EntityData.Children = types.NewOrderedMap()
    reachableTime.EntityData.Leafs = types.NewOrderedMap()
    reachableTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", reachableTime.Seconds})

    reachableTime.EntityData.YListKeys = []string {}

    return &(reachableTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime
// RA retransmit time
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Seconds interface{}
}

func (retransTime *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime) GetEntityData() *types.CommonEntityData {
    retransTime.EntityData.YFilter = retransTime.YFilter
    retransTime.EntityData.YangName = "retrans-time"
    retransTime.EntityData.BundleName = "cisco_ios_xr"
    retransTime.EntityData.ParentYangName = "ra"
    retransTime.EntityData.SegmentPath = "retrans-time"
    retransTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/slaac-interface/router-advert-detail/ra/" + retransTime.EntityData.SegmentPath
    retransTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransTime.EntityData.Children = types.NewOrderedMap()
    retransTime.EntityData.Leafs = types.NewOrderedMap()
    retransTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", retransTime.Seconds})

    retransTime.EntityData.YListKeys = []string {}

    return &(retransTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ
// Prefix Queue
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IPV6 Prefix address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixAddress interface{}

    // IPv6 Auto generated address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Eui64 interface{}

    // IPV6 Prefix Valid Life Time. The type is interface{} with range:
    // 0..4294967295.
    ValidLifeTime interface{}

    // IPV6 Prefix Preferred Life Time. The type is interface{} with range:
    // 0..4294967295.
    PreferredLifeTime interface{}

    // IPV6 Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLen interface{}

    // IPv6 Address Specific Flags. The type is interface{} with range:
    // 0..4294967295.
    Flags interface{}

    // Prefix Address Specific Flags. The type is interface{} with range:
    // 0..4294967295.
    PfxFlags interface{}
}

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetEntityData() *types.CommonEntityData {
    prefixQ.EntityData.YFilter = prefixQ.YFilter
    prefixQ.EntityData.YangName = "prefix-q"
    prefixQ.EntityData.BundleName = "cisco_ios_xr"
    prefixQ.EntityData.ParentYangName = "ra"
    prefixQ.EntityData.SegmentPath = "prefix-q" + types.AddNoKeyToken(prefixQ)
    prefixQ.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-oper:ipv6-node-discovery/nodes/node/slaac-interfaces/slaac-interface/router-advert-detail/ra/" + prefixQ.EntityData.SegmentPath
    prefixQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixQ.EntityData.Children = types.NewOrderedMap()
    prefixQ.EntityData.Leafs = types.NewOrderedMap()
    prefixQ.EntityData.Leafs.Append("prefix-address", types.YLeaf{"PrefixAddress", prefixQ.PrefixAddress})
    prefixQ.EntityData.Leafs.Append("eui64", types.YLeaf{"Eui64", prefixQ.Eui64})
    prefixQ.EntityData.Leafs.Append("valid-life-time", types.YLeaf{"ValidLifeTime", prefixQ.ValidLifeTime})
    prefixQ.EntityData.Leafs.Append("preferred-life-time", types.YLeaf{"PreferredLifeTime", prefixQ.PreferredLifeTime})
    prefixQ.EntityData.Leafs.Append("prefix-len", types.YLeaf{"PrefixLen", prefixQ.PrefixLen})
    prefixQ.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", prefixQ.Flags})
    prefixQ.EntityData.Leafs.Append("pfx-flags", types.YLeaf{"PfxFlags", prefixQ.PfxFlags})

    prefixQ.EntityData.YListKeys = []string {}

    return &(prefixQ.EntityData)
}

