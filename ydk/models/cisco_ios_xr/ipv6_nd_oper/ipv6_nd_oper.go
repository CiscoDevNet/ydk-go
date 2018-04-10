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
    Ipv6NdBndlState_error Ipv6NdBndlState = "error"

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
    Ipv6NdShState_delete Ipv6NdShState = "delete"
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
    ipv6NodeDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6NodeDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6NodeDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6NodeDiscovery.EntityData.Children = make(map[string]types.YChild)
    ipv6NodeDiscovery.EntityData.Children["nodes"] = types.YChild{"Nodes", &ipv6NodeDiscovery.Nodes}
    ipv6NodeDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6NodeDiscovery.EntityData)
}

// Ipv6NodeDiscovery_Nodes
// IPv6 node discovery list of nodes
type Ipv6NodeDiscovery_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 node discovery operational data for a particular node. The type is
    // slice of Ipv6NodeDiscovery_Nodes_Node.
    Node []Ipv6NodeDiscovery_Nodes_Node
}

func (nodes *Ipv6NodeDiscovery_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ipv6-node-discovery"
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

// Ipv6NodeDiscovery_Nodes_Node
// IPv6 node discovery operational data for a
// particular node
type Ipv6NodeDiscovery_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Ipv6NodeDiscovery_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["neighbor-interfaces"] = types.YChild{"NeighborInterfaces", &node.NeighborInterfaces}
    node.EntityData.Children["neighbor-summary"] = types.YChild{"NeighborSummary", &node.NeighborSummary}
    node.EntityData.Children["bundle-nodes"] = types.YChild{"BundleNodes", &node.BundleNodes}
    node.EntityData.Children["bundle-interfaces"] = types.YChild{"BundleInterfaces", &node.BundleInterfaces}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["nd-virtual-routers"] = types.YChild{"NdVirtualRouters", &node.NdVirtualRouters}
    node.EntityData.Children["slaac-interfaces"] = types.YChild{"SlaacInterfaces", &node.SlaacInterfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
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
    NeighborInterface []Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface
}

func (neighborInterfaces *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces) GetEntityData() *types.CommonEntityData {
    neighborInterfaces.EntityData.YFilter = neighborInterfaces.YFilter
    neighborInterfaces.EntityData.YangName = "neighbor-interfaces"
    neighborInterfaces.EntityData.BundleName = "cisco_ios_xr"
    neighborInterfaces.EntityData.ParentYangName = "node"
    neighborInterfaces.EntityData.SegmentPath = "neighbor-interfaces"
    neighborInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborInterfaces.EntityData.Children = make(map[string]types.YChild)
    neighborInterfaces.EntityData.Children["neighbor-interface"] = types.YChild{"NeighborInterface", nil}
    for i := range neighborInterfaces.NeighborInterface {
        neighborInterfaces.EntityData.Children[types.GetSegmentPath(&neighborInterfaces.NeighborInterface[i])] = types.YChild{"NeighborInterface", &neighborInterfaces.NeighborInterface[i]}
    }
    neighborInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighborInterfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface
// IPv6 node discovery neighbor interface
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // IPv6 node discovery list of neighbor host addresses.
    HostAddresses Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses
}

func (neighborInterface *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface) GetEntityData() *types.CommonEntityData {
    neighborInterface.EntityData.YFilter = neighborInterface.YFilter
    neighborInterface.EntityData.YangName = "neighbor-interface"
    neighborInterface.EntityData.BundleName = "cisco_ios_xr"
    neighborInterface.EntityData.ParentYangName = "neighbor-interfaces"
    neighborInterface.EntityData.SegmentPath = "neighbor-interface" + "[interface-name='" + fmt.Sprintf("%v", neighborInterface.InterfaceName) + "']"
    neighborInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborInterface.EntityData.Children = make(map[string]types.YChild)
    neighborInterface.EntityData.Children["host-addresses"] = types.YChild{"HostAddresses", &neighborInterface.HostAddresses}
    neighborInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    neighborInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", neighborInterface.InterfaceName}
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
    HostAddress []Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress
}

func (hostAddresses *Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses) GetEntityData() *types.CommonEntityData {
    hostAddresses.EntityData.YFilter = hostAddresses.YFilter
    hostAddresses.EntityData.YangName = "host-addresses"
    hostAddresses.EntityData.BundleName = "cisco_ios_xr"
    hostAddresses.EntityData.ParentYangName = "neighbor-interface"
    hostAddresses.EntityData.SegmentPath = "host-addresses"
    hostAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAddresses.EntityData.Children = make(map[string]types.YChild)
    hostAddresses.EntityData.Children["host-address"] = types.YChild{"HostAddress", nil}
    for i := range hostAddresses.HostAddress {
        hostAddresses.EntityData.Children[types.GetSegmentPath(&hostAddresses.HostAddress[i])] = types.YChild{"HostAddress", &hostAddresses.HostAddress[i]}
    }
    hostAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hostAddresses.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress
// IPv6 Neighbor detailed information
type Ipv6NodeDiscovery_Nodes_Node_NeighborInterfaces_NeighborInterface_HostAddresses_HostAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Host Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    HostAddress interface{}

    // Current state. The type is Ipv6NdShState.
    ReachabilityState interface{}

    // Link-Layer Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    hostAddress.EntityData.SegmentPath = "host-address" + "[host-address='" + fmt.Sprintf("%v", hostAddress.HostAddress) + "']"
    hostAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAddress.EntityData.Children = make(map[string]types.YChild)
    hostAddress.EntityData.Children["last-reached-time"] = types.YChild{"LastReachedTime", &hostAddress.LastReachedTime}
    hostAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    hostAddress.EntityData.Leafs["host-address"] = types.YLeaf{"HostAddress", hostAddress.HostAddress}
    hostAddress.EntityData.Leafs["reachability-state"] = types.YLeaf{"ReachabilityState", hostAddress.ReachabilityState}
    hostAddress.EntityData.Leafs["link-layer-address"] = types.YLeaf{"LinkLayerAddress", hostAddress.LinkLayerAddress}
    hostAddress.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", hostAddress.Encapsulation}
    hostAddress.EntityData.Leafs["selected-encapsulation"] = types.YLeaf{"SelectedEncapsulation", hostAddress.SelectedEncapsulation}
    hostAddress.EntityData.Leafs["origin-encapsulation"] = types.YLeaf{"OriginEncapsulation", hostAddress.OriginEncapsulation}
    hostAddress.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", hostAddress.InterfaceName}
    hostAddress.EntityData.Leafs["location"] = types.YLeaf{"Location", hostAddress.Location}
    hostAddress.EntityData.Leafs["is-router"] = types.YLeaf{"IsRouter", hostAddress.IsRouter}
    hostAddress.EntityData.Leafs["serg-flags"] = types.YLeaf{"SergFlags", hostAddress.SergFlags}
    hostAddress.EntityData.Leafs["vrfid"] = types.YLeaf{"Vrfid", hostAddress.Vrfid}
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
    lastReachedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReachedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReachedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReachedTime.EntityData.Children = make(map[string]types.YChild)
    lastReachedTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReachedTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastReachedTime.Seconds}
    return &(lastReachedTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary
// IPv6 Neighbor summary
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
}

func (neighborSummary *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary) GetEntityData() *types.CommonEntityData {
    neighborSummary.EntityData.YFilter = neighborSummary.YFilter
    neighborSummary.EntityData.YangName = "neighbor-summary"
    neighborSummary.EntityData.BundleName = "cisco_ios_xr"
    neighborSummary.EntityData.ParentYangName = "node"
    neighborSummary.EntityData.SegmentPath = "neighbor-summary"
    neighborSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborSummary.EntityData.Children = make(map[string]types.YChild)
    neighborSummary.EntityData.Children["multicast"] = types.YChild{"Multicast", &neighborSummary.Multicast}
    neighborSummary.EntityData.Children["static"] = types.YChild{"Static", &neighborSummary.Static}
    neighborSummary.EntityData.Children["dynamic"] = types.YChild{"Dynamic", &neighborSummary.Dynamic}
    neighborSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    neighborSummary.EntityData.Leafs["total-neighbor-entries"] = types.YLeaf{"TotalNeighborEntries", neighborSummary.TotalNeighborEntries}
    return &(neighborSummary.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast
// Multicast neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast struct {
    EntityData types.CommonEntityData
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

func (multicast *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Multicast) GetEntityData() *types.CommonEntityData {
    multicast.EntityData.YFilter = multicast.YFilter
    multicast.EntityData.YangName = "multicast"
    multicast.EntityData.BundleName = "cisco_ios_xr"
    multicast.EntityData.ParentYangName = "neighbor-summary"
    multicast.EntityData.SegmentPath = "multicast"
    multicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicast.EntityData.Children = make(map[string]types.YChild)
    multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    multicast.EntityData.Leafs["incomplete-entries"] = types.YLeaf{"IncompleteEntries", multicast.IncompleteEntries}
    multicast.EntityData.Leafs["reachable-entries"] = types.YLeaf{"ReachableEntries", multicast.ReachableEntries}
    multicast.EntityData.Leafs["stale-entries"] = types.YLeaf{"StaleEntries", multicast.StaleEntries}
    multicast.EntityData.Leafs["delayed-entries"] = types.YLeaf{"DelayedEntries", multicast.DelayedEntries}
    multicast.EntityData.Leafs["probe-entries"] = types.YLeaf{"ProbeEntries", multicast.ProbeEntries}
    multicast.EntityData.Leafs["deleted-entries"] = types.YLeaf{"DeletedEntries", multicast.DeletedEntries}
    multicast.EntityData.Leafs["subtotal-neighbor-entries"] = types.YLeaf{"SubtotalNeighborEntries", multicast.SubtotalNeighborEntries}
    return &(multicast.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static
// Static neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static struct {
    EntityData types.CommonEntityData
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

func (static *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xr"
    static.EntityData.ParentYangName = "neighbor-summary"
    static.EntityData.SegmentPath = "static"
    static.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    static.EntityData.Children = make(map[string]types.YChild)
    static.EntityData.Leafs = make(map[string]types.YLeaf)
    static.EntityData.Leafs["incomplete-entries"] = types.YLeaf{"IncompleteEntries", static.IncompleteEntries}
    static.EntityData.Leafs["reachable-entries"] = types.YLeaf{"ReachableEntries", static.ReachableEntries}
    static.EntityData.Leafs["stale-entries"] = types.YLeaf{"StaleEntries", static.StaleEntries}
    static.EntityData.Leafs["delayed-entries"] = types.YLeaf{"DelayedEntries", static.DelayedEntries}
    static.EntityData.Leafs["probe-entries"] = types.YLeaf{"ProbeEntries", static.ProbeEntries}
    static.EntityData.Leafs["deleted-entries"] = types.YLeaf{"DeletedEntries", static.DeletedEntries}
    static.EntityData.Leafs["subtotal-neighbor-entries"] = types.YLeaf{"SubtotalNeighborEntries", static.SubtotalNeighborEntries}
    return &(static.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic
// Dynamic neighbor summary
type Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic struct {
    EntityData types.CommonEntityData
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

func (dynamic *Ipv6NodeDiscovery_Nodes_Node_NeighborSummary_Dynamic) GetEntityData() *types.CommonEntityData {
    dynamic.EntityData.YFilter = dynamic.YFilter
    dynamic.EntityData.YangName = "dynamic"
    dynamic.EntityData.BundleName = "cisco_ios_xr"
    dynamic.EntityData.ParentYangName = "neighbor-summary"
    dynamic.EntityData.SegmentPath = "dynamic"
    dynamic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamic.EntityData.Children = make(map[string]types.YChild)
    dynamic.EntityData.Leafs = make(map[string]types.YLeaf)
    dynamic.EntityData.Leafs["incomplete-entries"] = types.YLeaf{"IncompleteEntries", dynamic.IncompleteEntries}
    dynamic.EntityData.Leafs["reachable-entries"] = types.YLeaf{"ReachableEntries", dynamic.ReachableEntries}
    dynamic.EntityData.Leafs["stale-entries"] = types.YLeaf{"StaleEntries", dynamic.StaleEntries}
    dynamic.EntityData.Leafs["delayed-entries"] = types.YLeaf{"DelayedEntries", dynamic.DelayedEntries}
    dynamic.EntityData.Leafs["probe-entries"] = types.YLeaf{"ProbeEntries", dynamic.ProbeEntries}
    dynamic.EntityData.Leafs["deleted-entries"] = types.YLeaf{"DeletedEntries", dynamic.DeletedEntries}
    dynamic.EntityData.Leafs["subtotal-neighbor-entries"] = types.YLeaf{"SubtotalNeighborEntries", dynamic.SubtotalNeighborEntries}
    return &(dynamic.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes
// IPv6 ND list of bundle nodes for a specific
// node
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ND operational data for a specific bundle node. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode.
    BundleNode []Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode
}

func (bundleNodes *Ipv6NodeDiscovery_Nodes_Node_BundleNodes) GetEntityData() *types.CommonEntityData {
    bundleNodes.EntityData.YFilter = bundleNodes.YFilter
    bundleNodes.EntityData.YangName = "bundle-nodes"
    bundleNodes.EntityData.BundleName = "cisco_ios_xr"
    bundleNodes.EntityData.ParentYangName = "node"
    bundleNodes.EntityData.SegmentPath = "bundle-nodes"
    bundleNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleNodes.EntityData.Children = make(map[string]types.YChild)
    bundleNodes.EntityData.Children["bundle-node"] = types.YChild{"BundleNode", nil}
    for i := range bundleNodes.BundleNode {
        bundleNodes.EntityData.Children[types.GetSegmentPath(&bundleNodes.BundleNode[i])] = types.YChild{"BundleNode", &bundleNodes.BundleNode[i]}
    }
    bundleNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleNodes.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode
// IPv6 ND operational data for a specific
// bundle node
type Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The bundle node name. The type is string with
    // pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (bundleNode *Ipv6NodeDiscovery_Nodes_Node_BundleNodes_BundleNode) GetEntityData() *types.CommonEntityData {
    bundleNode.EntityData.YFilter = bundleNode.YFilter
    bundleNode.EntityData.YangName = "bundle-node"
    bundleNode.EntityData.BundleName = "cisco_ios_xr"
    bundleNode.EntityData.ParentYangName = "bundle-nodes"
    bundleNode.EntityData.SegmentPath = "bundle-node" + "[node-name='" + fmt.Sprintf("%v", bundleNode.NodeName) + "']"
    bundleNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleNode.EntityData.Children = make(map[string]types.YChild)
    bundleNode.EntityData.Children["age"] = types.YChild{"Age", &bundleNode.Age}
    bundleNode.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleNode.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", bundleNode.NodeName}
    bundleNode.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", bundleNode.GroupId}
    bundleNode.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", bundleNode.ProcessName}
    bundleNode.EntityData.Leafs["sent-sequence-number"] = types.YLeaf{"SentSequenceNumber", bundleNode.SentSequenceNumber}
    bundleNode.EntityData.Leafs["received-sequence-number"] = types.YLeaf{"ReceivedSequenceNumber", bundleNode.ReceivedSequenceNumber}
    bundleNode.EntityData.Leafs["state"] = types.YLeaf{"State", bundleNode.State}
    bundleNode.EntityData.Leafs["state-changes"] = types.YLeaf{"StateChanges", bundleNode.StateChanges}
    bundleNode.EntityData.Leafs["sent-packets"] = types.YLeaf{"SentPackets", bundleNode.SentPackets}
    bundleNode.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bundleNode.ReceivedPackets}
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
    age.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    age.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    age.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    age.EntityData.Children = make(map[string]types.YChild)
    age.EntityData.Leafs = make(map[string]types.YLeaf)
    age.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", age.Seconds}
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
    BundleInterface []Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces) GetEntityData() *types.CommonEntityData {
    bundleInterfaces.EntityData.YFilter = bundleInterfaces.YFilter
    bundleInterfaces.EntityData.YangName = "bundle-interfaces"
    bundleInterfaces.EntityData.BundleName = "cisco_ios_xr"
    bundleInterfaces.EntityData.ParentYangName = "node"
    bundleInterfaces.EntityData.SegmentPath = "bundle-interfaces"
    bundleInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterfaces.EntityData.Children = make(map[string]types.YChild)
    bundleInterfaces.EntityData.Children["bundle-interface"] = types.YChild{"BundleInterface", nil}
    for i := range bundleInterfaces.BundleInterface {
        bundleInterfaces.EntityData.Children[types.GetSegmentPath(&bundleInterfaces.BundleInterface[i])] = types.YChild{"BundleInterface", &bundleInterfaces.BundleInterface[i]}
    }
    bundleInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleInterfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface
// IPv6 ND operational data for a specific
// bundler interface
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Parent interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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

    // Link local address.
    LocalAddress Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress

    // List of ND global addresses. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress.
    GlobalAddress []Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress

    // List of member nodes. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode.
    MemberNode []Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode
}

func (bundleInterface *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface) GetEntityData() *types.CommonEntityData {
    bundleInterface.EntityData.YFilter = bundleInterface.YFilter
    bundleInterface.EntityData.YangName = "bundle-interface"
    bundleInterface.EntityData.BundleName = "cisco_ios_xr"
    bundleInterface.EntityData.ParentYangName = "bundle-interfaces"
    bundleInterface.EntityData.SegmentPath = "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
    bundleInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterface.EntityData.Children = make(map[string]types.YChild)
    bundleInterface.EntityData.Children["nd-parameters"] = types.YChild{"NdParameters", &bundleInterface.NdParameters}
    bundleInterface.EntityData.Children["local-address"] = types.YChild{"LocalAddress", &bundleInterface.LocalAddress}
    bundleInterface.EntityData.Children["global-address"] = types.YChild{"GlobalAddress", nil}
    for i := range bundleInterface.GlobalAddress {
        bundleInterface.EntityData.Children[types.GetSegmentPath(&bundleInterface.GlobalAddress[i])] = types.YChild{"GlobalAddress", &bundleInterface.GlobalAddress[i]}
    }
    bundleInterface.EntityData.Children["member-node"] = types.YChild{"MemberNode", nil}
    for i := range bundleInterface.MemberNode {
        bundleInterface.EntityData.Children[types.GetSegmentPath(&bundleInterface.MemberNode[i])] = types.YChild{"MemberNode", &bundleInterface.MemberNode[i]}
    }
    bundleInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bundleInterface.InterfaceName}
    bundleInterface.EntityData.Leafs["parent-interface-name"] = types.YLeaf{"ParentInterfaceName", bundleInterface.ParentInterfaceName}
    bundleInterface.EntityData.Leafs["iftype"] = types.YLeaf{"Iftype", bundleInterface.Iftype}
    bundleInterface.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", bundleInterface.Mtu}
    bundleInterface.EntityData.Leafs["etype"] = types.YLeaf{"Etype", bundleInterface.Etype}
    bundleInterface.EntityData.Leafs["vlan-tag"] = types.YLeaf{"VlanTag", bundleInterface.VlanTag}
    bundleInterface.EntityData.Leafs["mac-addr-size"] = types.YLeaf{"MacAddrSize", bundleInterface.MacAddrSize}
    bundleInterface.EntityData.Leafs["mac-addr"] = types.YLeaf{"MacAddr", bundleInterface.MacAddr}
    bundleInterface.EntityData.Leafs["is-interface-enabled"] = types.YLeaf{"IsInterfaceEnabled", bundleInterface.IsInterfaceEnabled}
    bundleInterface.EntityData.Leafs["is-ipv6-enabled"] = types.YLeaf{"IsIpv6Enabled", bundleInterface.IsIpv6Enabled}
    bundleInterface.EntityData.Leafs["is-mpls-enabled"] = types.YLeaf{"IsMplsEnabled", bundleInterface.IsMplsEnabled}
    bundleInterface.EntityData.Leafs["member-link"] = types.YLeaf{"MemberLink", bundleInterface.MemberLink}
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
    ndParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ndParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ndParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ndParameters.EntityData.Children = make(map[string]types.YChild)
    ndParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    ndParameters.EntityData.Leafs["is-dad-enabled"] = types.YLeaf{"IsDadEnabled", ndParameters.IsDadEnabled}
    ndParameters.EntityData.Leafs["dad-attempts"] = types.YLeaf{"DadAttempts", ndParameters.DadAttempts}
    ndParameters.EntityData.Leafs["is-icm-pv6-redirect"] = types.YLeaf{"IsIcmPv6Redirect", ndParameters.IsIcmPv6Redirect}
    ndParameters.EntityData.Leafs["is-dhcp-managed"] = types.YLeaf{"IsDhcpManaged", ndParameters.IsDhcpManaged}
    ndParameters.EntityData.Leafs["is-route-address-managed"] = types.YLeaf{"IsRouteAddressManaged", ndParameters.IsRouteAddressManaged}
    ndParameters.EntityData.Leafs["is-suppressed"] = types.YLeaf{"IsSuppressed", ndParameters.IsSuppressed}
    ndParameters.EntityData.Leafs["send-unicast-ra"] = types.YLeaf{"SendUnicastRa", ndParameters.SendUnicastRa}
    ndParameters.EntityData.Leafs["nd-retransmit-interval"] = types.YLeaf{"NdRetransmitInterval", ndParameters.NdRetransmitInterval}
    ndParameters.EntityData.Leafs["nd-min-transmit-interval"] = types.YLeaf{"NdMinTransmitInterval", ndParameters.NdMinTransmitInterval}
    ndParameters.EntityData.Leafs["nd-max-transmit-interval"] = types.YLeaf{"NdMaxTransmitInterval", ndParameters.NdMaxTransmitInterval}
    ndParameters.EntityData.Leafs["nd-advertisement-lifetime"] = types.YLeaf{"NdAdvertisementLifetime", ndParameters.NdAdvertisementLifetime}
    ndParameters.EntityData.Leafs["nd-reachable-time"] = types.YLeaf{"NdReachableTime", ndParameters.NdReachableTime}
    ndParameters.EntityData.Leafs["nd-cache-limit"] = types.YLeaf{"NdCacheLimit", ndParameters.NdCacheLimit}
    ndParameters.EntityData.Leafs["complete-protocol-count"] = types.YLeaf{"CompleteProtocolCount", ndParameters.CompleteProtocolCount}
    ndParameters.EntityData.Leafs["complete-glean-count"] = types.YLeaf{"CompleteGleanCount", ndParameters.CompleteGleanCount}
    ndParameters.EntityData.Leafs["incomplete-protocol-count"] = types.YLeaf{"IncompleteProtocolCount", ndParameters.IncompleteProtocolCount}
    ndParameters.EntityData.Leafs["incomplete-glean-count"] = types.YLeaf{"IncompleteGleanCount", ndParameters.IncompleteGleanCount}
    ndParameters.EntityData.Leafs["dropped-protocol-req-count"] = types.YLeaf{"DroppedProtocolReqCount", ndParameters.DroppedProtocolReqCount}
    ndParameters.EntityData.Leafs["dropped-glean-req-count"] = types.YLeaf{"DroppedGleanReqCount", ndParameters.DroppedGleanReqCount}
    return &(ndParameters.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress
// Link local address
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "bundle-interface"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = make(map[string]types.YChild)
    localAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    localAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", localAddress.Ipv6Address}
    return &(localAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress
// List of ND global addresses
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (globalAddress *Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_GlobalAddress) GetEntityData() *types.CommonEntityData {
    globalAddress.EntityData.YFilter = globalAddress.YFilter
    globalAddress.EntityData.YangName = "global-address"
    globalAddress.EntityData.BundleName = "cisco_ios_xr"
    globalAddress.EntityData.ParentYangName = "bundle-interface"
    globalAddress.EntityData.SegmentPath = "global-address"
    globalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAddress.EntityData.Children = make(map[string]types.YChild)
    globalAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    globalAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", globalAddress.Ipv6Address}
    return &(globalAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode
// List of member nodes
type Ipv6NodeDiscovery_Nodes_Node_BundleInterfaces_BundleInterface_MemberNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    memberNode.EntityData.SegmentPath = "member-node"
    memberNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberNode.EntityData.Children = make(map[string]types.YChild)
    memberNode.EntityData.Leafs = make(map[string]types.YLeaf)
    memberNode.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", memberNode.NodeName}
    memberNode.EntityData.Leafs["total-links"] = types.YLeaf{"TotalLinks", memberNode.TotalLinks}
    return &(memberNode.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_Interfaces
// IPv6 node discovery list of interfaces for a
// specific node
type Ipv6NodeDiscovery_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6  node discovery operational data for a specific node and interface.
    // The type is slice of Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface_.
    Interface_ []Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface
}

func (interfaces *Ipv6NodeDiscovery_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
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

// Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface
// IPv6  node discovery operational data for a
// specific node and interface
type Ipv6NodeDiscovery_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["is-dad-enabled"] = types.YLeaf{"IsDadEnabled", self.IsDadEnabled}
    self.EntityData.Leafs["dad-attempts"] = types.YLeaf{"DadAttempts", self.DadAttempts}
    self.EntityData.Leafs["is-icm-pv6-redirect"] = types.YLeaf{"IsIcmPv6Redirect", self.IsIcmPv6Redirect}
    self.EntityData.Leafs["is-dhcp-managed"] = types.YLeaf{"IsDhcpManaged", self.IsDhcpManaged}
    self.EntityData.Leafs["is-route-address-managed"] = types.YLeaf{"IsRouteAddressManaged", self.IsRouteAddressManaged}
    self.EntityData.Leafs["is-suppressed"] = types.YLeaf{"IsSuppressed", self.IsSuppressed}
    self.EntityData.Leafs["send-unicast-ra"] = types.YLeaf{"SendUnicastRa", self.SendUnicastRa}
    self.EntityData.Leafs["nd-retransmit-interval"] = types.YLeaf{"NdRetransmitInterval", self.NdRetransmitInterval}
    self.EntityData.Leafs["nd-min-transmit-interval"] = types.YLeaf{"NdMinTransmitInterval", self.NdMinTransmitInterval}
    self.EntityData.Leafs["nd-max-transmit-interval"] = types.YLeaf{"NdMaxTransmitInterval", self.NdMaxTransmitInterval}
    self.EntityData.Leafs["nd-advertisement-lifetime"] = types.YLeaf{"NdAdvertisementLifetime", self.NdAdvertisementLifetime}
    self.EntityData.Leafs["nd-reachable-time"] = types.YLeaf{"NdReachableTime", self.NdReachableTime}
    self.EntityData.Leafs["nd-cache-limit"] = types.YLeaf{"NdCacheLimit", self.NdCacheLimit}
    self.EntityData.Leafs["complete-protocol-count"] = types.YLeaf{"CompleteProtocolCount", self.CompleteProtocolCount}
    self.EntityData.Leafs["complete-glean-count"] = types.YLeaf{"CompleteGleanCount", self.CompleteGleanCount}
    self.EntityData.Leafs["incomplete-protocol-count"] = types.YLeaf{"IncompleteProtocolCount", self.IncompleteProtocolCount}
    self.EntityData.Leafs["incomplete-glean-count"] = types.YLeaf{"IncompleteGleanCount", self.IncompleteGleanCount}
    self.EntityData.Leafs["dropped-protocol-req-count"] = types.YLeaf{"DroppedProtocolReqCount", self.DroppedProtocolReqCount}
    self.EntityData.Leafs["dropped-glean-req-count"] = types.YLeaf{"DroppedGleanReqCount", self.DroppedGleanReqCount}
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
    NdVirtualRouter []Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter
}

func (ndVirtualRouters *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters) GetEntityData() *types.CommonEntityData {
    ndVirtualRouters.EntityData.YFilter = ndVirtualRouters.YFilter
    ndVirtualRouters.EntityData.YangName = "nd-virtual-routers"
    ndVirtualRouters.EntityData.BundleName = "cisco_ios_xr"
    ndVirtualRouters.EntityData.ParentYangName = "node"
    ndVirtualRouters.EntityData.SegmentPath = "nd-virtual-routers"
    ndVirtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ndVirtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ndVirtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ndVirtualRouters.EntityData.Children = make(map[string]types.YChild)
    ndVirtualRouters.EntityData.Children["nd-virtual-router"] = types.YChild{"NdVirtualRouter", nil}
    for i := range ndVirtualRouters.NdVirtualRouter {
        ndVirtualRouters.EntityData.Children[types.GetSegmentPath(&ndVirtualRouters.NdVirtualRouter[i])] = types.YChild{"NdVirtualRouter", &ndVirtualRouters.NdVirtualRouter[i]}
    }
    ndVirtualRouters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ndVirtualRouters.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter
// IPv6 ND virtual  router operational data for
// a specific interface
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Link-Layer Address. The type is string with pattern:
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

    // Link local address.
    LocalAddress Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress

    // List of ND global addresses. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress.
    VrGlobalAddress []Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress
}

func (ndVirtualRouter *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter) GetEntityData() *types.CommonEntityData {
    ndVirtualRouter.EntityData.YFilter = ndVirtualRouter.YFilter
    ndVirtualRouter.EntityData.YangName = "nd-virtual-router"
    ndVirtualRouter.EntityData.BundleName = "cisco_ios_xr"
    ndVirtualRouter.EntityData.ParentYangName = "nd-virtual-routers"
    ndVirtualRouter.EntityData.SegmentPath = "nd-virtual-router" + "[interface-name='" + fmt.Sprintf("%v", ndVirtualRouter.InterfaceName) + "']"
    ndVirtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ndVirtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ndVirtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ndVirtualRouter.EntityData.Children = make(map[string]types.YChild)
    ndVirtualRouter.EntityData.Children["local-address"] = types.YChild{"LocalAddress", &ndVirtualRouter.LocalAddress}
    ndVirtualRouter.EntityData.Children["vr-global-address"] = types.YChild{"VrGlobalAddress", nil}
    for i := range ndVirtualRouter.VrGlobalAddress {
        ndVirtualRouter.EntityData.Children[types.GetSegmentPath(&ndVirtualRouter.VrGlobalAddress[i])] = types.YChild{"VrGlobalAddress", &ndVirtualRouter.VrGlobalAddress[i]}
    }
    ndVirtualRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    ndVirtualRouter.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", ndVirtualRouter.InterfaceName}
    ndVirtualRouter.EntityData.Leafs["link-layer-address"] = types.YLeaf{"LinkLayerAddress", ndVirtualRouter.LinkLayerAddress}
    ndVirtualRouter.EntityData.Leafs["context"] = types.YLeaf{"Context", ndVirtualRouter.Context}
    ndVirtualRouter.EntityData.Leafs["state"] = types.YLeaf{"State", ndVirtualRouter.State}
    ndVirtualRouter.EntityData.Leafs["flags"] = types.YLeaf{"Flags", ndVirtualRouter.Flags}
    ndVirtualRouter.EntityData.Leafs["vr-gl-addr-ct"] = types.YLeaf{"VrGlAddrCt", ndVirtualRouter.VrGlAddrCt}
    return &(ndVirtualRouter.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress
// Link local address
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (localAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "nd-virtual-router"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = make(map[string]types.YChild)
    localAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    localAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", localAddress.Ipv6Address}
    return &(localAddress.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress
// List of ND global addresses
type Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (vrGlobalAddress *Ipv6NodeDiscovery_Nodes_Node_NdVirtualRouters_NdVirtualRouter_VrGlobalAddress) GetEntityData() *types.CommonEntityData {
    vrGlobalAddress.EntityData.YFilter = vrGlobalAddress.YFilter
    vrGlobalAddress.EntityData.YangName = "vr-global-address"
    vrGlobalAddress.EntityData.BundleName = "cisco_ios_xr"
    vrGlobalAddress.EntityData.ParentYangName = "nd-virtual-router"
    vrGlobalAddress.EntityData.SegmentPath = "vr-global-address"
    vrGlobalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrGlobalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrGlobalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrGlobalAddress.EntityData.Children = make(map[string]types.YChild)
    vrGlobalAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    vrGlobalAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", vrGlobalAddress.Ipv6Address}
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
    SlaacInterface []Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface
}

func (slaacInterfaces *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces) GetEntityData() *types.CommonEntityData {
    slaacInterfaces.EntityData.YFilter = slaacInterfaces.YFilter
    slaacInterfaces.EntityData.YangName = "slaac-interfaces"
    slaacInterfaces.EntityData.BundleName = "cisco_ios_xr"
    slaacInterfaces.EntityData.ParentYangName = "node"
    slaacInterfaces.EntityData.SegmentPath = "slaac-interfaces"
    slaacInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaacInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaacInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaacInterfaces.EntityData.Children = make(map[string]types.YChild)
    slaacInterfaces.EntityData.Children["slaac-interface"] = types.YChild{"SlaacInterface", nil}
    for i := range slaacInterfaces.SlaacInterface {
        slaacInterfaces.EntityData.Children[types.GetSegmentPath(&slaacInterfaces.SlaacInterface[i])] = types.YChild{"SlaacInterface", &slaacInterfaces.SlaacInterface[i]}
    }
    slaacInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slaacInterfaces.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface
// IPv6 ND operational data for a specific slaac
// interface
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // IPv6 ND operational data for a specific slaac interface.
    RouterAdvertDetail Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail
}

func (slaacInterface *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface) GetEntityData() *types.CommonEntityData {
    slaacInterface.EntityData.YFilter = slaacInterface.YFilter
    slaacInterface.EntityData.YangName = "slaac-interface"
    slaacInterface.EntityData.BundleName = "cisco_ios_xr"
    slaacInterface.EntityData.ParentYangName = "slaac-interfaces"
    slaacInterface.EntityData.SegmentPath = "slaac-interface" + "[interface-name='" + fmt.Sprintf("%v", slaacInterface.InterfaceName) + "']"
    slaacInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaacInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaacInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaacInterface.EntityData.Children = make(map[string]types.YChild)
    slaacInterface.EntityData.Children["router-advert-detail"] = types.YChild{"RouterAdvertDetail", &slaacInterface.RouterAdvertDetail}
    slaacInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    slaacInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", slaacInterface.InterfaceName}
    return &(slaacInterface.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail
// IPv6 ND operational data for a specific
// slaac interface
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // idb. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Idb interface{}

    // slaac db. The type is slice of
    // Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra.
    Ra []Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra
}

func (routerAdvertDetail *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail) GetEntityData() *types.CommonEntityData {
    routerAdvertDetail.EntityData.YFilter = routerAdvertDetail.YFilter
    routerAdvertDetail.EntityData.YangName = "router-advert-detail"
    routerAdvertDetail.EntityData.BundleName = "cisco_ios_xr"
    routerAdvertDetail.EntityData.ParentYangName = "slaac-interface"
    routerAdvertDetail.EntityData.SegmentPath = "router-advert-detail"
    routerAdvertDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerAdvertDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerAdvertDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerAdvertDetail.EntityData.Children = make(map[string]types.YChild)
    routerAdvertDetail.EntityData.Children["ra"] = types.YChild{"Ra", nil}
    for i := range routerAdvertDetail.Ra {
        routerAdvertDetail.EntityData.Children[types.GetSegmentPath(&routerAdvertDetail.Ra[i])] = types.YChild{"Ra", &routerAdvertDetail.Ra[i]}
    }
    routerAdvertDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    routerAdvertDetail.EntityData.Leafs["idb"] = types.YLeaf{"Idb", routerAdvertDetail.Idb}
    return &(routerAdvertDetail.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra
// slaac db
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (ra *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra) GetEntityData() *types.CommonEntityData {
    ra.EntityData.YFilter = ra.YFilter
    ra.EntityData.YangName = "ra"
    ra.EntityData.BundleName = "cisco_ios_xr"
    ra.EntityData.ParentYangName = "router-advert-detail"
    ra.EntityData.SegmentPath = "ra"
    ra.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ra.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ra.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ra.EntityData.Children = make(map[string]types.YChild)
    ra.EntityData.Children["elapsed-ra-time"] = types.YChild{"ElapsedRaTime", &ra.ElapsedRaTime}
    ra.EntityData.Children["reachable-time"] = types.YChild{"ReachableTime", &ra.ReachableTime}
    ra.EntityData.Children["retrans-time"] = types.YChild{"RetransTime", &ra.RetransTime}
    ra.EntityData.Children["prefix-q"] = types.YChild{"PrefixQ", nil}
    for i := range ra.PrefixQ {
        ra.EntityData.Children[types.GetSegmentPath(&ra.PrefixQ[i])] = types.YChild{"PrefixQ", &ra.PrefixQ[i]}
    }
    ra.EntityData.Leafs = make(map[string]types.YLeaf)
    ra.EntityData.Leafs["address"] = types.YLeaf{"Address", ra.Address}
    ra.EntityData.Leafs["hops"] = types.YLeaf{"Hops", ra.Hops}
    ra.EntityData.Leafs["flags"] = types.YLeaf{"Flags", ra.Flags}
    ra.EntityData.Leafs["life-time"] = types.YLeaf{"LifeTime", ra.LifeTime}
    ra.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", ra.Mtu}
    ra.EntityData.Leafs["err-msg"] = types.YLeaf{"ErrMsg", ra.ErrMsg}
    ra.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", ra.VrfId}
    ra.EntityData.Leafs["u6-tbl-id"] = types.YLeaf{"U6TblId", ra.U6TblId}
    ra.EntityData.Leafs["rib-protoid"] = types.YLeaf{"RibProtoid", ra.RibProtoid}
    ra.EntityData.Leafs["default-router"] = types.YLeaf{"DefaultRouter", ra.DefaultRouter}
    ra.EntityData.Leafs["reachability"] = types.YLeaf{"Reachability", ra.Reachability}
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
    elapsedRaTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    elapsedRaTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    elapsedRaTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    elapsedRaTime.EntityData.Children = make(map[string]types.YChild)
    elapsedRaTime.EntityData.Leafs = make(map[string]types.YLeaf)
    elapsedRaTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", elapsedRaTime.Seconds}
    return &(elapsedRaTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_ReachableTime
// reachabletime
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
    reachableTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reachableTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reachableTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reachableTime.EntityData.Children = make(map[string]types.YChild)
    reachableTime.EntityData.Leafs = make(map[string]types.YLeaf)
    reachableTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", reachableTime.Seconds}
    return &(reachableTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_RetransTime
// retranstime
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
    retransTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransTime.EntityData.Children = make(map[string]types.YChild)
    retransTime.EntityData.Leafs = make(map[string]types.YLeaf)
    retransTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", retransTime.Seconds}
    return &(retransTime.EntityData)
}

// Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ
// Prefix Queue
type Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixAddress interface{}

    // IPv6 Auto generated address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (prefixQ *Ipv6NodeDiscovery_Nodes_Node_SlaacInterfaces_SlaacInterface_RouterAdvertDetail_Ra_PrefixQ) GetEntityData() *types.CommonEntityData {
    prefixQ.EntityData.YFilter = prefixQ.YFilter
    prefixQ.EntityData.YangName = "prefix-q"
    prefixQ.EntityData.BundleName = "cisco_ios_xr"
    prefixQ.EntityData.ParentYangName = "ra"
    prefixQ.EntityData.SegmentPath = "prefix-q"
    prefixQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixQ.EntityData.Children = make(map[string]types.YChild)
    prefixQ.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixQ.EntityData.Leafs["prefix-address"] = types.YLeaf{"PrefixAddress", prefixQ.PrefixAddress}
    prefixQ.EntityData.Leafs["eui64"] = types.YLeaf{"Eui64", prefixQ.Eui64}
    prefixQ.EntityData.Leafs["valid-life-time"] = types.YLeaf{"ValidLifeTime", prefixQ.ValidLifeTime}
    prefixQ.EntityData.Leafs["preferred-life-time"] = types.YLeaf{"PreferredLifeTime", prefixQ.PreferredLifeTime}
    prefixQ.EntityData.Leafs["prefix-len"] = types.YLeaf{"PrefixLen", prefixQ.PrefixLen}
    prefixQ.EntityData.Leafs["flags"] = types.YLeaf{"Flags", prefixQ.Flags}
    prefixQ.EntityData.Leafs["pfx-flags"] = types.YLeaf{"PfxFlags", prefixQ.PfxFlags}
    return &(prefixQ.EntityData)
}

