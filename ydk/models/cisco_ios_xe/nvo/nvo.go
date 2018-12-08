// This module contains a collection of YANG definitions
// for Cisco VxLAN feature configuration.
// 
// Copyright (c) 2013-2014 by Cisco Systems, Inc.
// All rights reserved.
package nvo

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package nvo"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:nvo nvo-instances}", reflect.TypeOf(NvoInstances{}))
    ydk.RegisterEntity("nvo:nvo-instances", reflect.TypeOf(NvoInstances{}))
}

type NvgreType struct {
}

func (id NvgreType) String() string {
	return "nvo:nvgre-type"
}

type VxlanType struct {
}

func (id VxlanType) String() string {
	return "nvo:vxlan-type"
}

type OverlayEncapType struct {
}

func (id OverlayEncapType) String() string {
	return "nvo:overlay-encap-type"
}

// NvoInstances
// vxlan instances
type NvoInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of instances. The type is slice of NvoInstances_NvoInstance.
    NvoInstance []*NvoInstances_NvoInstance
}

func (nvoInstances *NvoInstances) GetEntityData() *types.CommonEntityData {
    nvoInstances.EntityData.YFilter = nvoInstances.YFilter
    nvoInstances.EntityData.YangName = "nvo-instances"
    nvoInstances.EntityData.BundleName = "cisco_ios_xe"
    nvoInstances.EntityData.ParentYangName = "nvo"
    nvoInstances.EntityData.SegmentPath = "nvo:nvo-instances"
    nvoInstances.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nvoInstances.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nvoInstances.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nvoInstances.EntityData.Children = types.NewOrderedMap()
    nvoInstances.EntityData.Children.Append("nvo-instance", types.YChild{"NvoInstance", nil})
    for i := range nvoInstances.NvoInstance {
        nvoInstances.EntityData.Children.Append(types.GetSegmentPath(nvoInstances.NvoInstance[i]), types.YChild{"NvoInstance", nvoInstances.NvoInstance[i]})
    }
    nvoInstances.EntityData.Leafs = types.NewOrderedMap()

    nvoInstances.EntityData.YListKeys = []string {}

    return &(nvoInstances.EntityData)
}

// NvoInstances_NvoInstance
// List of instances
type NvoInstances_NvoInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Network Virtualization Overlay Instance 
    // Identifier. The type is interface{} with range: 0..65535.
    NvoId interface{}

    // Encapsulation type. The type is one of the following: NvgreTypeVxlanType.
    OverlayEncapsulation interface{}

    // Source interface name. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name This attribute is mandatory.
    SourceInterface interface{}

    // VNI member attributes. The type is slice of
    // NvoInstances_NvoInstance_VirtualNetwork.
    VirtualNetwork []*NvoInstances_NvoInstance_VirtualNetwork
}

func (nvoInstance *NvoInstances_NvoInstance) GetEntityData() *types.CommonEntityData {
    nvoInstance.EntityData.YFilter = nvoInstance.YFilter
    nvoInstance.EntityData.YangName = "nvo-instance"
    nvoInstance.EntityData.BundleName = "cisco_ios_xe"
    nvoInstance.EntityData.ParentYangName = "nvo-instances"
    nvoInstance.EntityData.SegmentPath = "nvo-instance" + types.AddKeyToken(nvoInstance.NvoId, "nvo-id")
    nvoInstance.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nvoInstance.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nvoInstance.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nvoInstance.EntityData.Children = types.NewOrderedMap()
    nvoInstance.EntityData.Children.Append("virtual-network", types.YChild{"VirtualNetwork", nil})
    for i := range nvoInstance.VirtualNetwork {
        nvoInstance.EntityData.Children.Append(types.GetSegmentPath(nvoInstance.VirtualNetwork[i]), types.YChild{"VirtualNetwork", nvoInstance.VirtualNetwork[i]})
    }
    nvoInstance.EntityData.Leafs = types.NewOrderedMap()
    nvoInstance.EntityData.Leafs.Append("nvo-id", types.YLeaf{"NvoId", nvoInstance.NvoId})
    nvoInstance.EntityData.Leafs.Append("overlay-encapsulation", types.YLeaf{"OverlayEncapsulation", nvoInstance.OverlayEncapsulation})
    nvoInstance.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", nvoInstance.SourceInterface})

    nvoInstance.EntityData.YListKeys = []string {"NvoId"}

    return &(nvoInstance.EntityData)
}

// NvoInstances_NvoInstance_VirtualNetwork
// VNI member attributes
type NvoInstances_NvoInstance_VirtualNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single Virtual Network Identifier  or start of
    // range. The type is interface{} with range: 1..16777214. This attribute is
    // mandatory.
    VniStart interface{}

    // This attribute is a key. End of Virtual Network Identifier range  (make
    // equal to vni-start for single vni. The type is interface{} with range:
    // 1..16777214. This attribute is mandatory.
    VniEnd interface{}

    // VRF Name. The type is string. Refers to
    // ietf_routing.Routing_RoutingInstance_Name
    RoutingInstance interface{}

    // How to peform endpoint discovery. The type is EndHostDiscovery. The default
    // value is flood-and-learn.
    EndHostDiscovery interface{}

    // Use control protocol BGP to discover  peers. The type is interface{}.
    Bgp interface{}

    // Enable ARP request suppression for this VNI. The type is interface{}.
    SuppressArp interface{}

    // Mulitcast group range associated  with the VxLAN segment(s).
    Multicast NvoInstances_NvoInstance_VirtualNetwork_Multicast

    // List of VTEP peers. The type is slice of
    // NvoInstances_NvoInstance_VirtualNetwork_Peers.
    Peers []*NvoInstances_NvoInstance_VirtualNetwork_Peers
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetEntityData() *types.CommonEntityData {
    virtualNetwork.EntityData.YFilter = virtualNetwork.YFilter
    virtualNetwork.EntityData.YangName = "virtual-network"
    virtualNetwork.EntityData.BundleName = "cisco_ios_xe"
    virtualNetwork.EntityData.ParentYangName = "nvo-instance"
    virtualNetwork.EntityData.SegmentPath = "virtual-network" + types.AddKeyToken(virtualNetwork.VniStart, "vni-start") + types.AddKeyToken(virtualNetwork.VniEnd, "vni-end")
    virtualNetwork.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    virtualNetwork.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    virtualNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    virtualNetwork.EntityData.Children = types.NewOrderedMap()
    virtualNetwork.EntityData.Children.Append("multicast", types.YChild{"Multicast", &virtualNetwork.Multicast})
    virtualNetwork.EntityData.Children.Append("peers", types.YChild{"Peers", nil})
    for i := range virtualNetwork.Peers {
        virtualNetwork.EntityData.Children.Append(types.GetSegmentPath(virtualNetwork.Peers[i]), types.YChild{"Peers", virtualNetwork.Peers[i]})
    }
    virtualNetwork.EntityData.Leafs = types.NewOrderedMap()
    virtualNetwork.EntityData.Leafs.Append("vni-start", types.YLeaf{"VniStart", virtualNetwork.VniStart})
    virtualNetwork.EntityData.Leafs.Append("vni-end", types.YLeaf{"VniEnd", virtualNetwork.VniEnd})
    virtualNetwork.EntityData.Leafs.Append("routing-instance", types.YLeaf{"RoutingInstance", virtualNetwork.RoutingInstance})
    virtualNetwork.EntityData.Leafs.Append("end-host-discovery", types.YLeaf{"EndHostDiscovery", virtualNetwork.EndHostDiscovery})
    virtualNetwork.EntityData.Leafs.Append("bgp", types.YLeaf{"Bgp", virtualNetwork.Bgp})
    virtualNetwork.EntityData.Leafs.Append("suppress-arp", types.YLeaf{"SuppressArp", virtualNetwork.SuppressArp})

    virtualNetwork.EntityData.YListKeys = []string {"VniStart", "VniEnd"}

    return &(virtualNetwork.EntityData)
}

// NvoInstances_NvoInstance_VirtualNetwork_Multicast
// Mulitcast group range associated 
// with the VxLAN segment(s)
type NvoInstances_NvoInstance_VirtualNetwork_Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single IPV4 Multicast group  address or start of range. The type is string
    // with pattern:
    // (2((2[4-9])|(3[0-9]))\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    MulticastGroupMin interface{}

    // End of IPV4 Multicast group  address (leave unspecified for single value.
    // The type is string with pattern:
    // (2((2[4-9])|(3[0-9]))\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    MulticastGroupMax interface{}
}

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetEntityData() *types.CommonEntityData {
    multicast.EntityData.YFilter = multicast.YFilter
    multicast.EntityData.YangName = "multicast"
    multicast.EntityData.BundleName = "cisco_ios_xe"
    multicast.EntityData.ParentYangName = "virtual-network"
    multicast.EntityData.SegmentPath = "multicast"
    multicast.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multicast.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multicast.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multicast.EntityData.Children = types.NewOrderedMap()
    multicast.EntityData.Leafs = types.NewOrderedMap()
    multicast.EntityData.Leafs.Append("multicast-group-min", types.YLeaf{"MulticastGroupMin", multicast.MulticastGroupMin})
    multicast.EntityData.Leafs.Append("multicast-group-max", types.YLeaf{"MulticastGroupMax", multicast.MulticastGroupMax})

    multicast.EntityData.YListKeys = []string {}

    return &(multicast.EntityData)
}

// NvoInstances_NvoInstance_VirtualNetwork_Peers
// List of VTEP peers
type NvoInstances_NvoInstance_VirtualNetwork_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VTEP peer IP address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIp interface{}
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xe"
    peers.EntityData.ParentYangName = "virtual-network"
    peers.EntityData.SegmentPath = "peers" + types.AddKeyToken(peers.PeerIp, "peer-ip")
    peers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Leafs = types.NewOrderedMap()
    peers.EntityData.Leafs.Append("peer-ip", types.YLeaf{"PeerIp", peers.PeerIp})

    peers.EntityData.YListKeys = []string {"PeerIp"}

    return &(peers.EntityData)
}

// NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery represents How to peform endpoint discovery
type NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery string

const (
    // Discover end-hosts using data plane 
    //               flood and learn
    NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery_flood_and_learn NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery = "flood-and-learn"

    // Discover end-hosts using bgp-evpn
    NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery_bgp NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery = "bgp"
)

