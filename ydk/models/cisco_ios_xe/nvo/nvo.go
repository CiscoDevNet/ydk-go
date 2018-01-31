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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of instances. The type is slice of NvoInstances_NvoInstance.
    NvoInstance []NvoInstances_NvoInstance
}

func (nvoInstances *NvoInstances) GetFilter() yfilter.YFilter { return nvoInstances.YFilter }

func (nvoInstances *NvoInstances) SetFilter(yf yfilter.YFilter) { nvoInstances.YFilter = yf }

func (nvoInstances *NvoInstances) GetGoName(yname string) string {
    if yname == "nvo-instance" { return "NvoInstance" }
    return ""
}

func (nvoInstances *NvoInstances) GetSegmentPath() string {
    return "nvo:nvo-instances"
}

func (nvoInstances *NvoInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nvo-instance" {
        for _, c := range nvoInstances.NvoInstance {
            if nvoInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvoInstances_NvoInstance{}
        nvoInstances.NvoInstance = append(nvoInstances.NvoInstance, child)
        return &nvoInstances.NvoInstance[len(nvoInstances.NvoInstance)-1]
    }
    return nil
}

func (nvoInstances *NvoInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nvoInstances.NvoInstance {
        children[nvoInstances.NvoInstance[i].GetSegmentPath()] = &nvoInstances.NvoInstance[i]
    }
    return children
}

func (nvoInstances *NvoInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nvoInstances *NvoInstances) GetBundleName() string { return "cisco_ios_xe" }

func (nvoInstances *NvoInstances) GetYangName() string { return "nvo-instances" }

func (nvoInstances *NvoInstances) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nvoInstances *NvoInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nvoInstances *NvoInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nvoInstances *NvoInstances) SetParent(parent types.Entity) { nvoInstances.parent = parent }

func (nvoInstances *NvoInstances) GetParent() types.Entity { return nvoInstances.parent }

func (nvoInstances *NvoInstances) GetParentYangName() string { return "nvo" }

// NvoInstances_NvoInstance
// List of instances
type NvoInstances_NvoInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Network Virtualization Overlay Instance 
    // Identifier. The type is interface{} with range: 0..65535.
    NvoId interface{}

    // Source interface name. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name This attribute is mandatory.
    SourceInterface interface{}

    // Encapsulation type. The type is one of the following: NvgreTypeVxlanType.
    OverlayEncapsulation interface{}

    // VNI member attributes. The type is slice of
    // NvoInstances_NvoInstance_VirtualNetwork.
    VirtualNetwork []NvoInstances_NvoInstance_VirtualNetwork
}

func (nvoInstance *NvoInstances_NvoInstance) GetFilter() yfilter.YFilter { return nvoInstance.YFilter }

func (nvoInstance *NvoInstances_NvoInstance) SetFilter(yf yfilter.YFilter) { nvoInstance.YFilter = yf }

func (nvoInstance *NvoInstances_NvoInstance) GetGoName(yname string) string {
    if yname == "nvo-id" { return "NvoId" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "overlay-encapsulation" { return "OverlayEncapsulation" }
    if yname == "virtual-network" { return "VirtualNetwork" }
    return ""
}

func (nvoInstance *NvoInstances_NvoInstance) GetSegmentPath() string {
    return "nvo-instance" + "[nvo-id='" + fmt.Sprintf("%v", nvoInstance.NvoId) + "']"
}

func (nvoInstance *NvoInstances_NvoInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-network" {
        for _, c := range nvoInstance.VirtualNetwork {
            if nvoInstance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvoInstances_NvoInstance_VirtualNetwork{}
        nvoInstance.VirtualNetwork = append(nvoInstance.VirtualNetwork, child)
        return &nvoInstance.VirtualNetwork[len(nvoInstance.VirtualNetwork)-1]
    }
    return nil
}

func (nvoInstance *NvoInstances_NvoInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nvoInstance.VirtualNetwork {
        children[nvoInstance.VirtualNetwork[i].GetSegmentPath()] = &nvoInstance.VirtualNetwork[i]
    }
    return children
}

func (nvoInstance *NvoInstances_NvoInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nvo-id"] = nvoInstance.NvoId
    leafs["source-interface"] = nvoInstance.SourceInterface
    leafs["overlay-encapsulation"] = nvoInstance.OverlayEncapsulation
    return leafs
}

func (nvoInstance *NvoInstances_NvoInstance) GetBundleName() string { return "cisco_ios_xe" }

func (nvoInstance *NvoInstances_NvoInstance) GetYangName() string { return "nvo-instance" }

func (nvoInstance *NvoInstances_NvoInstance) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nvoInstance *NvoInstances_NvoInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nvoInstance *NvoInstances_NvoInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nvoInstance *NvoInstances_NvoInstance) SetParent(parent types.Entity) { nvoInstance.parent = parent }

func (nvoInstance *NvoInstances_NvoInstance) GetParent() types.Entity { return nvoInstance.parent }

func (nvoInstance *NvoInstances_NvoInstance) GetParentYangName() string { return "nvo-instances" }

// NvoInstances_NvoInstance_VirtualNetwork
// VNI member attributes
type NvoInstances_NvoInstance_VirtualNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single Virtual Network Identifier  or start of
    // range. The type is interface{} with range: 1..16777214. This attribute is
    // mandatory.
    VniStart interface{}

    // This attribute is a key. End of Virtual Network Identifier range  (make
    // equal to vni-start for single vni. The type is interface{} with range:
    // 1..16777214. This attribute is mandatory.
    VniEnd interface{}

    // How to peform endpoint discovery. The type is EndHostDiscovery. The default
    // value is flood-and-learn.
    EndHostDiscovery interface{}

    // Use control protocol BGP to discover  peers. The type is interface{}.
    Bgp interface{}

    // Enable ARP request suppression for this VNI. The type is interface{}.
    SuppressArp interface{}

    // VRF Name. The type is string. Refers to
    // ietf_routing.Routing_RoutingInstance_Name
    RoutingInstance interface{}

    // Mulitcast group range associated  with the VxLAN segment(s).
    Multicast NvoInstances_NvoInstance_VirtualNetwork_Multicast

    // List of VTEP peers. The type is slice of
    // NvoInstances_NvoInstance_VirtualNetwork_Peers.
    Peers []NvoInstances_NvoInstance_VirtualNetwork_Peers
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetFilter() yfilter.YFilter { return virtualNetwork.YFilter }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) SetFilter(yf yfilter.YFilter) { virtualNetwork.YFilter = yf }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetGoName(yname string) string {
    if yname == "vni-start" { return "VniStart" }
    if yname == "vni-end" { return "VniEnd" }
    if yname == "end-host-discovery" { return "EndHostDiscovery" }
    if yname == "bgp" { return "Bgp" }
    if yname == "suppress-arp" { return "SuppressArp" }
    if yname == "routing-instance" { return "RoutingInstance" }
    if yname == "multicast" { return "Multicast" }
    if yname == "peers" { return "Peers" }
    return ""
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetSegmentPath() string {
    return "virtual-network" + "[vni-start='" + fmt.Sprintf("%v", virtualNetwork.VniStart) + "']" + "[vni-end='" + fmt.Sprintf("%v", virtualNetwork.VniEnd) + "']"
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multicast" {
        return &virtualNetwork.Multicast
    }
    if childYangName == "peers" {
        for _, c := range virtualNetwork.Peers {
            if virtualNetwork.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvoInstances_NvoInstance_VirtualNetwork_Peers{}
        virtualNetwork.Peers = append(virtualNetwork.Peers, child)
        return &virtualNetwork.Peers[len(virtualNetwork.Peers)-1]
    }
    return nil
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["multicast"] = &virtualNetwork.Multicast
    for i := range virtualNetwork.Peers {
        children[virtualNetwork.Peers[i].GetSegmentPath()] = &virtualNetwork.Peers[i]
    }
    return children
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vni-start"] = virtualNetwork.VniStart
    leafs["vni-end"] = virtualNetwork.VniEnd
    leafs["end-host-discovery"] = virtualNetwork.EndHostDiscovery
    leafs["bgp"] = virtualNetwork.Bgp
    leafs["suppress-arp"] = virtualNetwork.SuppressArp
    leafs["routing-instance"] = virtualNetwork.RoutingInstance
    return leafs
}

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetBundleName() string { return "cisco_ios_xe" }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetYangName() string { return "virtual-network" }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) SetParent(parent types.Entity) { virtualNetwork.parent = parent }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetParent() types.Entity { return virtualNetwork.parent }

func (virtualNetwork *NvoInstances_NvoInstance_VirtualNetwork) GetParentYangName() string { return "nvo-instance" }

// NvoInstances_NvoInstance_VirtualNetwork_Multicast
// Mulitcast group range associated 
// with the VxLAN segment(s)
type NvoInstances_NvoInstance_VirtualNetwork_Multicast struct {
    parent types.Entity
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

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetFilter() yfilter.YFilter { return multicast.YFilter }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) SetFilter(yf yfilter.YFilter) { multicast.YFilter = yf }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetGoName(yname string) string {
    if yname == "multicast-group-min" { return "MulticastGroupMin" }
    if yname == "multicast-group-max" { return "MulticastGroupMax" }
    return ""
}

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetSegmentPath() string {
    return "multicast"
}

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multicast-group-min"] = multicast.MulticastGroupMin
    leafs["multicast-group-max"] = multicast.MulticastGroupMax
    return leafs
}

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetBundleName() string { return "cisco_ios_xe" }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetYangName() string { return "multicast" }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) SetParent(parent types.Entity) { multicast.parent = parent }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetParent() types.Entity { return multicast.parent }

func (multicast *NvoInstances_NvoInstance_VirtualNetwork_Multicast) GetParentYangName() string { return "virtual-network" }

// NvoInstances_NvoInstance_VirtualNetwork_Peers
// List of VTEP peers
type NvoInstances_NvoInstance_VirtualNetwork_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VTEP peer IP address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIp interface{}
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetGoName(yname string) string {
    if yname == "peer-ip" { return "PeerIp" }
    return ""
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetSegmentPath() string {
    return "peers" + "[peer-ip='" + fmt.Sprintf("%v", peers.PeerIp) + "']"
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-ip"] = peers.PeerIp
    return leafs
}

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetBundleName() string { return "cisco_ios_xe" }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetYangName() string { return "peers" }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetParent() types.Entity { return peers.parent }

func (peers *NvoInstances_NvoInstance_VirtualNetwork_Peers) GetParentYangName() string { return "virtual-network" }

// NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery represents How to peform endpoint discovery
type NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery string

const (
    // Discover end-hosts using data plane 
    //               flood and learn
    NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery_flood_and_learn NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery = "flood-and-learn"

    // Discover end-hosts using bgp-evpn
    NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery_bgp NvoInstances_NvoInstance_VirtualNetwork_EndHostDiscovery = "bgp"
)

