// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-ipsub package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ip-subscriber: IP subscriber operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_ipsub_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_ipsub_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-ipsub-oper ip-subscriber}", reflect.TypeOf(IpSubscriber{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-ipsub-oper:ip-subscriber", reflect.TypeOf(IpSubscriber{}))
}

// IpsubMaIntfStateData represents Interface states
type IpsubMaIntfStateData string

const (
    // Invalid state
    IpsubMaIntfStateData_invalid IpsubMaIntfStateData = "invalid"

    // Initial state
    IpsubMaIntfStateData_initialized IpsubMaIntfStateData = "initialized"

    // Interface creation started
    IpsubMaIntfStateData_session_creation_started IpsubMaIntfStateData = "session-creation-started"

    // Interface created in IM, AAA session start
    // called
    IpsubMaIntfStateData_control_policy_executing IpsubMaIntfStateData = "control-policy-executing"

    // AAA session created
    IpsubMaIntfStateData_control_policy_executed IpsubMaIntfStateData = "control-policy-executed"

    // Interface config activated
    IpsubMaIntfStateData_session_features_applied IpsubMaIntfStateData = "session-features-applied"

    // Interface address and VRF information received
    // from IPv4
    IpsubMaIntfStateData_vrf_configured IpsubMaIntfStateData = "vrf-configured"

    // VRF configuration received and interface config
    // activated
    IpsubMaIntfStateData_adding_adjacency IpsubMaIntfStateData = "adding-adjacency"

    // Subscriber AIB adjacency added
    IpsubMaIntfStateData_adjacency_added IpsubMaIntfStateData = "adjacency-added"

    // Session up
    IpsubMaIntfStateData_up IpsubMaIntfStateData = "up"

    // Session down
    IpsubMaIntfStateData_down IpsubMaIntfStateData = "down"

    // Session down in progress
    IpsubMaIntfStateData_address_family_down IpsubMaIntfStateData = "address-family-down"

    // Session down complete
    IpsubMaIntfStateData_address_family_down_complete IpsubMaIntfStateData = "address-family-down-complete"

    // Session teardown in progress
    IpsubMaIntfStateData_disconnecting IpsubMaIntfStateData = "disconnecting"

    // Session disconnected
    IpsubMaIntfStateData_disconnected IpsubMaIntfStateData = "disconnected"

    // Session in error state
    IpsubMaIntfStateData_error_ IpsubMaIntfStateData = "error"
)

// IpsubMaParentIntfVlan represents Access interface VLAN type
type IpsubMaParentIntfVlan string

const (
    // Plain
    IpsubMaParentIntfVlan_plain IpsubMaParentIntfVlan = "plain"

    // Ambiguous
    IpsubMaParentIntfVlan_ambiguous IpsubMaParentIntfVlan = "ambiguous"
)

// IpsubMaParentIntfStateData represents Parent interface state
type IpsubMaParentIntfStateData string

const (
    // Interface being deleted
    IpsubMaParentIntfStateData_deleted IpsubMaParentIntfStateData = "deleted"

    // Interface operationally down
    IpsubMaParentIntfStateData_down IpsubMaParentIntfStateData = "down"

    // Interface up
    IpsubMaParentIntfStateData_up IpsubMaParentIntfStateData = "up"
)

// IpsubMaIntfInitiatorData represents Ipsub ma intf initiator data
type IpsubMaIntfInitiatorData string

const (
    // Session creation via DHCP discover packet
    IpsubMaIntfInitiatorData_dhcp IpsubMaIntfInitiatorData = "dhcp"

    // Session creation via unclassified IPv4 packet
    IpsubMaIntfInitiatorData_packet_trigger IpsubMaIntfInitiatorData = "packet-trigger"

    // Invalid Trigger
    IpsubMaIntfInitiatorData_invalid_trigger IpsubMaIntfInitiatorData = "invalid-trigger"
)

// IpSubscriber
// IP subscriber operational data
type IpSubscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP subscriber operational data for a particular location.
    Nodes IpSubscriber_Nodes
}

func (ipSubscriber *IpSubscriber) GetEntityData() *types.CommonEntityData {
    ipSubscriber.EntityData.YFilter = ipSubscriber.YFilter
    ipSubscriber.EntityData.YangName = "ip-subscriber"
    ipSubscriber.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriber.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-ipsub-oper"
    ipSubscriber.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-ipsub-oper:ip-subscriber"
    ipSubscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriber.EntityData.Children = types.NewOrderedMap()
    ipSubscriber.EntityData.Children.Append("nodes", types.YChild{"Nodes", &ipSubscriber.Nodes})
    ipSubscriber.EntityData.Leafs = types.NewOrderedMap()

    ipSubscriber.EntityData.YListKeys = []string {}

    return &(ipSubscriber.EntityData)
}

// IpSubscriber_Nodes
// IP subscriber operational data for a particular
// location
type IpSubscriber_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of IpSubscriber_Nodes_Node.
    Node []*IpSubscriber_Nodes_Node
}

func (nodes *IpSubscriber_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ip-subscriber"
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

// IpSubscriber_Nodes_Node
// Location. For eg., 0/1/CPU0
type IpSubscriber_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node ID to filter on. For eg., 0/1/CPU0. The
    // type is string with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // IP subscriber interface summary.
    Summary IpSubscriber_Nodes_Node_Summary

    // IP subscriber interface table.
    Interfaces IpSubscriber_Nodes_Node_Interfaces

    // IP subscriber access interface table.
    AccessInterfaces IpSubscriber_Nodes_Node_AccessInterfaces
}

func (node *IpSubscriber_Nodes_Node) GetEntityData() *types.CommonEntityData {
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
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("access-interfaces", types.YChild{"AccessInterfaces", &node.AccessInterfaces})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// IpSubscriber_Nodes_Node_Summary
// IP subscriber interface summary
type IpSubscriber_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface summary statistics.
    AccessInterfaceSummary IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary

    // Initiator interface counts.
    InterfaceCounts IpSubscriber_Nodes_Node_Summary_InterfaceCounts

    // Array of VRFs with IPSUB interfaces. The type is slice of
    // IpSubscriber_Nodes_Node_Summary_Vrf.
    Vrf []*IpSubscriber_Nodes_Node_Summary_Vrf
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("access-interface-summary", types.YChild{"AccessInterfaceSummary", &summary.AccessInterfaceSummary})
    summary.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &summary.InterfaceCounts})
    summary.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range summary.Vrf {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Vrf[i]), types.YChild{"Vrf", summary.Vrf[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary
// Access interface summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with subscriber configuration. The type is interface{}
    // with range: 0..4294967295.
    Interfaces interface{}

    // Summary counts per initiator.
    Initiators IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators

    // Summary counts per initiator for ipv6 session.
    Ipv6Initiators IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummary.EntityData.YFilter = accessInterfaceSummary.YFilter
    accessInterfaceSummary.EntityData.YangName = "access-interface-summary"
    accessInterfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummary.EntityData.ParentYangName = "summary"
    accessInterfaceSummary.EntityData.SegmentPath = "access-interface-summary"
    accessInterfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummary.EntityData.Children = types.NewOrderedMap()
    accessInterfaceSummary.EntityData.Children.Append("initiators", types.YChild{"Initiators", &accessInterfaceSummary.Initiators})
    accessInterfaceSummary.EntityData.Children.Append("ipv6-initiators", types.YChild{"Ipv6Initiators", &accessInterfaceSummary.Ipv6Initiators})
    accessInterfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    accessInterfaceSummary.EntityData.Leafs.Append("interfaces", types.YLeaf{"Interfaces", accessInterfaceSummary.Interfaces})

    accessInterfaceSummary.EntityData.YListKeys = []string {}

    return &(accessInterfaceSummary.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators
// Summary counts per initiator
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP summary statistics.
    Dhcp IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp

    // Packet trigger summary statistics.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetEntityData() *types.CommonEntityData {
    initiators.EntityData.YFilter = initiators.YFilter
    initiators.EntityData.YangName = "initiators"
    initiators.EntityData.BundleName = "cisco_ios_xr"
    initiators.EntityData.ParentYangName = "access-interface-summary"
    initiators.EntityData.SegmentPath = "initiators"
    initiators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initiators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initiators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initiators.EntityData.Children = types.NewOrderedMap()
    initiators.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &initiators.Dhcp})
    initiators.EntityData.Children.Append("packet-trigger", types.YChild{"PacketTrigger", &initiators.PacketTrigger})
    initiators.EntityData.Leafs = types.NewOrderedMap()

    initiators.EntityData.YListKeys = []string {}

    return &(initiators.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp
// DHCP summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xr"
    dhcp.EntityData.ParentYangName = "initiators"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Leafs = types.NewOrderedMap()
    dhcp.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", dhcp.FsolPackets})
    dhcp.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", dhcp.FsolBytes})

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger
// Packet trigger summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetEntityData() *types.CommonEntityData {
    packetTrigger.EntityData.YFilter = packetTrigger.YFilter
    packetTrigger.EntityData.YangName = "packet-trigger"
    packetTrigger.EntityData.BundleName = "cisco_ios_xr"
    packetTrigger.EntityData.ParentYangName = "initiators"
    packetTrigger.EntityData.SegmentPath = "packet-trigger"
    packetTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetTrigger.EntityData.Children = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", packetTrigger.FsolPackets})
    packetTrigger.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", packetTrigger.FsolBytes})

    packetTrigger.EntityData.YListKeys = []string {}

    return &(packetTrigger.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators
// Summary counts per initiator for ipv6 session
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP summary statistics.
    Dhcp IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp

    // Packet trigger summary statistics.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetEntityData() *types.CommonEntityData {
    ipv6Initiators.EntityData.YFilter = ipv6Initiators.YFilter
    ipv6Initiators.EntityData.YangName = "ipv6-initiators"
    ipv6Initiators.EntityData.BundleName = "cisco_ios_xr"
    ipv6Initiators.EntityData.ParentYangName = "access-interface-summary"
    ipv6Initiators.EntityData.SegmentPath = "ipv6-initiators"
    ipv6Initiators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Initiators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Initiators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Initiators.EntityData.Children = types.NewOrderedMap()
    ipv6Initiators.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &ipv6Initiators.Dhcp})
    ipv6Initiators.EntityData.Children.Append("packet-trigger", types.YChild{"PacketTrigger", &ipv6Initiators.PacketTrigger})
    ipv6Initiators.EntityData.Leafs = types.NewOrderedMap()

    ipv6Initiators.EntityData.YListKeys = []string {}

    return &(ipv6Initiators.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp
// DHCP summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xr"
    dhcp.EntityData.ParentYangName = "ipv6-initiators"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Leafs = types.NewOrderedMap()
    dhcp.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", dhcp.FsolPackets})
    dhcp.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", dhcp.FsolBytes})

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger
// Packet trigger summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetEntityData() *types.CommonEntityData {
    packetTrigger.EntityData.YFilter = packetTrigger.YFilter
    packetTrigger.EntityData.YangName = "packet-trigger"
    packetTrigger.EntityData.BundleName = "cisco_ios_xr"
    packetTrigger.EntityData.ParentYangName = "ipv6-initiators"
    packetTrigger.EntityData.SegmentPath = "packet-trigger"
    packetTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetTrigger.EntityData.Children = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", packetTrigger.FsolPackets})
    packetTrigger.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", packetTrigger.FsolBytes})

    packetTrigger.EntityData.YListKeys = []string {}

    return &(packetTrigger.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts
// Initiator interface counts
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initiators.
    Initiators IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators

    // IPv6 Initiators.
    Ipv6Initiators IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "summary"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Children.Append("initiators", types.YChild{"Initiators", &interfaceCounts.Initiators})
    interfaceCounts.EntityData.Children.Append("ipv6-initiators", types.YChild{"Ipv6Initiators", &interfaceCounts.Ipv6Initiators})
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators
// Initiators
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP.
    Dhcp IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp

    // Packet trigger.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetEntityData() *types.CommonEntityData {
    initiators.EntityData.YFilter = initiators.YFilter
    initiators.EntityData.YangName = "initiators"
    initiators.EntityData.BundleName = "cisco_ios_xr"
    initiators.EntityData.ParentYangName = "interface-counts"
    initiators.EntityData.SegmentPath = "initiators"
    initiators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initiators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initiators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initiators.EntityData.Children = types.NewOrderedMap()
    initiators.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &initiators.Dhcp})
    initiators.EntityData.Children.Append("packet-trigger", types.YChild{"PacketTrigger", &initiators.PacketTrigger})
    initiators.EntityData.Leafs = types.NewOrderedMap()

    initiators.EntityData.YListKeys = []string {}

    return &(initiators.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp
// DHCP
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid. The type is interface{} with range: 0..4294967295.
    Invalid interface{}

    // Initialized. The type is interface{} with range: 0..4294967295.
    Initialized interface{}

    // Session creation started. The type is interface{} with range:
    // 0..4294967295.
    SessionCreationStarted interface{}

    // Control policy executing. The type is interface{} with range:
    // 0..4294967295.
    ControlPolicyExecuting interface{}

    // Control policy executed. The type is interface{} with range: 0..4294967295.
    ControlPolicyExecuted interface{}

    // Session features applied. The type is interface{} with range:
    // 0..4294967295.
    SessionFeaturesApplied interface{}

    // VRF configured. The type is interface{} with range: 0..4294967295.
    VrfConfigured interface{}

    // Adding adjacency. The type is interface{} with range: 0..4294967295.
    AddingAdjacency interface{}

    // Adjacency added. The type is interface{} with range: 0..4294967295.
    AdjacencyAdded interface{}

    // Up. The type is interface{} with range: 0..4294967295.
    Up interface{}

    // Down. The type is interface{} with range: 0..4294967295.
    Down interface{}

    // Disconnecting. The type is interface{} with range: 0..4294967295.
    Disconnecting interface{}

    // Disconnected. The type is interface{} with range: 0..4294967295.
    Disconnected interface{}

    // Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Total number of interfaces in all states. The type is interface{} with
    // range: 0..4294967295.
    TotalInterfaces interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xr"
    dhcp.EntityData.ParentYangName = "initiators"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Leafs = types.NewOrderedMap()
    dhcp.EntityData.Leafs.Append("invalid", types.YLeaf{"Invalid", dhcp.Invalid})
    dhcp.EntityData.Leafs.Append("initialized", types.YLeaf{"Initialized", dhcp.Initialized})
    dhcp.EntityData.Leafs.Append("session-creation-started", types.YLeaf{"SessionCreationStarted", dhcp.SessionCreationStarted})
    dhcp.EntityData.Leafs.Append("control-policy-executing", types.YLeaf{"ControlPolicyExecuting", dhcp.ControlPolicyExecuting})
    dhcp.EntityData.Leafs.Append("control-policy-executed", types.YLeaf{"ControlPolicyExecuted", dhcp.ControlPolicyExecuted})
    dhcp.EntityData.Leafs.Append("session-features-applied", types.YLeaf{"SessionFeaturesApplied", dhcp.SessionFeaturesApplied})
    dhcp.EntityData.Leafs.Append("vrf-configured", types.YLeaf{"VrfConfigured", dhcp.VrfConfigured})
    dhcp.EntityData.Leafs.Append("adding-adjacency", types.YLeaf{"AddingAdjacency", dhcp.AddingAdjacency})
    dhcp.EntityData.Leafs.Append("adjacency-added", types.YLeaf{"AdjacencyAdded", dhcp.AdjacencyAdded})
    dhcp.EntityData.Leafs.Append("up", types.YLeaf{"Up", dhcp.Up})
    dhcp.EntityData.Leafs.Append("down", types.YLeaf{"Down", dhcp.Down})
    dhcp.EntityData.Leafs.Append("disconnecting", types.YLeaf{"Disconnecting", dhcp.Disconnecting})
    dhcp.EntityData.Leafs.Append("disconnected", types.YLeaf{"Disconnected", dhcp.Disconnected})
    dhcp.EntityData.Leafs.Append("error", types.YLeaf{"Error", dhcp.Error})
    dhcp.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", dhcp.TotalInterfaces})

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger
// Packet trigger
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid. The type is interface{} with range: 0..4294967295.
    Invalid interface{}

    // Initialized. The type is interface{} with range: 0..4294967295.
    Initialized interface{}

    // Session creation started. The type is interface{} with range:
    // 0..4294967295.
    SessionCreationStarted interface{}

    // Control policy executing. The type is interface{} with range:
    // 0..4294967295.
    ControlPolicyExecuting interface{}

    // Control policy executed. The type is interface{} with range: 0..4294967295.
    ControlPolicyExecuted interface{}

    // Session features applied. The type is interface{} with range:
    // 0..4294967295.
    SessionFeaturesApplied interface{}

    // VRF configured. The type is interface{} with range: 0..4294967295.
    VrfConfigured interface{}

    // Adding adjacency. The type is interface{} with range: 0..4294967295.
    AddingAdjacency interface{}

    // Adjacency added. The type is interface{} with range: 0..4294967295.
    AdjacencyAdded interface{}

    // Up. The type is interface{} with range: 0..4294967295.
    Up interface{}

    // Down. The type is interface{} with range: 0..4294967295.
    Down interface{}

    // Disconnecting. The type is interface{} with range: 0..4294967295.
    Disconnecting interface{}

    // Disconnected. The type is interface{} with range: 0..4294967295.
    Disconnected interface{}

    // Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Total number of interfaces in all states. The type is interface{} with
    // range: 0..4294967295.
    TotalInterfaces interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetEntityData() *types.CommonEntityData {
    packetTrigger.EntityData.YFilter = packetTrigger.YFilter
    packetTrigger.EntityData.YangName = "packet-trigger"
    packetTrigger.EntityData.BundleName = "cisco_ios_xr"
    packetTrigger.EntityData.ParentYangName = "initiators"
    packetTrigger.EntityData.SegmentPath = "packet-trigger"
    packetTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetTrigger.EntityData.Children = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs.Append("invalid", types.YLeaf{"Invalid", packetTrigger.Invalid})
    packetTrigger.EntityData.Leafs.Append("initialized", types.YLeaf{"Initialized", packetTrigger.Initialized})
    packetTrigger.EntityData.Leafs.Append("session-creation-started", types.YLeaf{"SessionCreationStarted", packetTrigger.SessionCreationStarted})
    packetTrigger.EntityData.Leafs.Append("control-policy-executing", types.YLeaf{"ControlPolicyExecuting", packetTrigger.ControlPolicyExecuting})
    packetTrigger.EntityData.Leafs.Append("control-policy-executed", types.YLeaf{"ControlPolicyExecuted", packetTrigger.ControlPolicyExecuted})
    packetTrigger.EntityData.Leafs.Append("session-features-applied", types.YLeaf{"SessionFeaturesApplied", packetTrigger.SessionFeaturesApplied})
    packetTrigger.EntityData.Leafs.Append("vrf-configured", types.YLeaf{"VrfConfigured", packetTrigger.VrfConfigured})
    packetTrigger.EntityData.Leafs.Append("adding-adjacency", types.YLeaf{"AddingAdjacency", packetTrigger.AddingAdjacency})
    packetTrigger.EntityData.Leafs.Append("adjacency-added", types.YLeaf{"AdjacencyAdded", packetTrigger.AdjacencyAdded})
    packetTrigger.EntityData.Leafs.Append("up", types.YLeaf{"Up", packetTrigger.Up})
    packetTrigger.EntityData.Leafs.Append("down", types.YLeaf{"Down", packetTrigger.Down})
    packetTrigger.EntityData.Leafs.Append("disconnecting", types.YLeaf{"Disconnecting", packetTrigger.Disconnecting})
    packetTrigger.EntityData.Leafs.Append("disconnected", types.YLeaf{"Disconnected", packetTrigger.Disconnected})
    packetTrigger.EntityData.Leafs.Append("error", types.YLeaf{"Error", packetTrigger.Error})
    packetTrigger.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", packetTrigger.TotalInterfaces})

    packetTrigger.EntityData.YListKeys = []string {}

    return &(packetTrigger.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators
// IPv6 Initiators
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP.
    Dhcp IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp

    // Packet trigger.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetEntityData() *types.CommonEntityData {
    ipv6Initiators.EntityData.YFilter = ipv6Initiators.YFilter
    ipv6Initiators.EntityData.YangName = "ipv6-initiators"
    ipv6Initiators.EntityData.BundleName = "cisco_ios_xr"
    ipv6Initiators.EntityData.ParentYangName = "interface-counts"
    ipv6Initiators.EntityData.SegmentPath = "ipv6-initiators"
    ipv6Initiators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Initiators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Initiators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Initiators.EntityData.Children = types.NewOrderedMap()
    ipv6Initiators.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &ipv6Initiators.Dhcp})
    ipv6Initiators.EntityData.Children.Append("packet-trigger", types.YChild{"PacketTrigger", &ipv6Initiators.PacketTrigger})
    ipv6Initiators.EntityData.Leafs = types.NewOrderedMap()

    ipv6Initiators.EntityData.YListKeys = []string {}

    return &(ipv6Initiators.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp
// DHCP
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid. The type is interface{} with range: 0..4294967295.
    Invalid interface{}

    // Initialized. The type is interface{} with range: 0..4294967295.
    Initialized interface{}

    // Session creation started. The type is interface{} with range:
    // 0..4294967295.
    SessionCreationStarted interface{}

    // Control policy executing. The type is interface{} with range:
    // 0..4294967295.
    ControlPolicyExecuting interface{}

    // Control policy executed. The type is interface{} with range: 0..4294967295.
    ControlPolicyExecuted interface{}

    // Session features applied. The type is interface{} with range:
    // 0..4294967295.
    SessionFeaturesApplied interface{}

    // VRF configured. The type is interface{} with range: 0..4294967295.
    VrfConfigured interface{}

    // Adding adjacency. The type is interface{} with range: 0..4294967295.
    AddingAdjacency interface{}

    // Adjacency added. The type is interface{} with range: 0..4294967295.
    AdjacencyAdded interface{}

    // Up. The type is interface{} with range: 0..4294967295.
    Up interface{}

    // Down. The type is interface{} with range: 0..4294967295.
    Down interface{}

    // Disconnecting. The type is interface{} with range: 0..4294967295.
    Disconnecting interface{}

    // Disconnected. The type is interface{} with range: 0..4294967295.
    Disconnected interface{}

    // Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Total number of interfaces in all states. The type is interface{} with
    // range: 0..4294967295.
    TotalInterfaces interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xr"
    dhcp.EntityData.ParentYangName = "ipv6-initiators"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Leafs = types.NewOrderedMap()
    dhcp.EntityData.Leafs.Append("invalid", types.YLeaf{"Invalid", dhcp.Invalid})
    dhcp.EntityData.Leafs.Append("initialized", types.YLeaf{"Initialized", dhcp.Initialized})
    dhcp.EntityData.Leafs.Append("session-creation-started", types.YLeaf{"SessionCreationStarted", dhcp.SessionCreationStarted})
    dhcp.EntityData.Leafs.Append("control-policy-executing", types.YLeaf{"ControlPolicyExecuting", dhcp.ControlPolicyExecuting})
    dhcp.EntityData.Leafs.Append("control-policy-executed", types.YLeaf{"ControlPolicyExecuted", dhcp.ControlPolicyExecuted})
    dhcp.EntityData.Leafs.Append("session-features-applied", types.YLeaf{"SessionFeaturesApplied", dhcp.SessionFeaturesApplied})
    dhcp.EntityData.Leafs.Append("vrf-configured", types.YLeaf{"VrfConfigured", dhcp.VrfConfigured})
    dhcp.EntityData.Leafs.Append("adding-adjacency", types.YLeaf{"AddingAdjacency", dhcp.AddingAdjacency})
    dhcp.EntityData.Leafs.Append("adjacency-added", types.YLeaf{"AdjacencyAdded", dhcp.AdjacencyAdded})
    dhcp.EntityData.Leafs.Append("up", types.YLeaf{"Up", dhcp.Up})
    dhcp.EntityData.Leafs.Append("down", types.YLeaf{"Down", dhcp.Down})
    dhcp.EntityData.Leafs.Append("disconnecting", types.YLeaf{"Disconnecting", dhcp.Disconnecting})
    dhcp.EntityData.Leafs.Append("disconnected", types.YLeaf{"Disconnected", dhcp.Disconnected})
    dhcp.EntityData.Leafs.Append("error", types.YLeaf{"Error", dhcp.Error})
    dhcp.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", dhcp.TotalInterfaces})

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger
// Packet trigger
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid. The type is interface{} with range: 0..4294967295.
    Invalid interface{}

    // Initialized. The type is interface{} with range: 0..4294967295.
    Initialized interface{}

    // Session creation started. The type is interface{} with range:
    // 0..4294967295.
    SessionCreationStarted interface{}

    // Control policy executing. The type is interface{} with range:
    // 0..4294967295.
    ControlPolicyExecuting interface{}

    // Control policy executed. The type is interface{} with range: 0..4294967295.
    ControlPolicyExecuted interface{}

    // Session features applied. The type is interface{} with range:
    // 0..4294967295.
    SessionFeaturesApplied interface{}

    // VRF configured. The type is interface{} with range: 0..4294967295.
    VrfConfigured interface{}

    // Adding adjacency. The type is interface{} with range: 0..4294967295.
    AddingAdjacency interface{}

    // Adjacency added. The type is interface{} with range: 0..4294967295.
    AdjacencyAdded interface{}

    // Up. The type is interface{} with range: 0..4294967295.
    Up interface{}

    // Down. The type is interface{} with range: 0..4294967295.
    Down interface{}

    // Disconnecting. The type is interface{} with range: 0..4294967295.
    Disconnecting interface{}

    // Disconnected. The type is interface{} with range: 0..4294967295.
    Disconnected interface{}

    // Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Total number of interfaces in all states. The type is interface{} with
    // range: 0..4294967295.
    TotalInterfaces interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetEntityData() *types.CommonEntityData {
    packetTrigger.EntityData.YFilter = packetTrigger.YFilter
    packetTrigger.EntityData.YangName = "packet-trigger"
    packetTrigger.EntityData.BundleName = "cisco_ios_xr"
    packetTrigger.EntityData.ParentYangName = "ipv6-initiators"
    packetTrigger.EntityData.SegmentPath = "packet-trigger"
    packetTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetTrigger.EntityData.Children = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs.Append("invalid", types.YLeaf{"Invalid", packetTrigger.Invalid})
    packetTrigger.EntityData.Leafs.Append("initialized", types.YLeaf{"Initialized", packetTrigger.Initialized})
    packetTrigger.EntityData.Leafs.Append("session-creation-started", types.YLeaf{"SessionCreationStarted", packetTrigger.SessionCreationStarted})
    packetTrigger.EntityData.Leafs.Append("control-policy-executing", types.YLeaf{"ControlPolicyExecuting", packetTrigger.ControlPolicyExecuting})
    packetTrigger.EntityData.Leafs.Append("control-policy-executed", types.YLeaf{"ControlPolicyExecuted", packetTrigger.ControlPolicyExecuted})
    packetTrigger.EntityData.Leafs.Append("session-features-applied", types.YLeaf{"SessionFeaturesApplied", packetTrigger.SessionFeaturesApplied})
    packetTrigger.EntityData.Leafs.Append("vrf-configured", types.YLeaf{"VrfConfigured", packetTrigger.VrfConfigured})
    packetTrigger.EntityData.Leafs.Append("adding-adjacency", types.YLeaf{"AddingAdjacency", packetTrigger.AddingAdjacency})
    packetTrigger.EntityData.Leafs.Append("adjacency-added", types.YLeaf{"AdjacencyAdded", packetTrigger.AdjacencyAdded})
    packetTrigger.EntityData.Leafs.Append("up", types.YLeaf{"Up", packetTrigger.Up})
    packetTrigger.EntityData.Leafs.Append("down", types.YLeaf{"Down", packetTrigger.Down})
    packetTrigger.EntityData.Leafs.Append("disconnecting", types.YLeaf{"Disconnecting", packetTrigger.Disconnecting})
    packetTrigger.EntityData.Leafs.Append("disconnected", types.YLeaf{"Disconnected", packetTrigger.Disconnected})
    packetTrigger.EntityData.Leafs.Append("error", types.YLeaf{"Error", packetTrigger.Error})
    packetTrigger.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", packetTrigger.TotalInterfaces})

    packetTrigger.EntityData.YListKeys = []string {}

    return &(packetTrigger.EntityData)
}

// IpSubscriber_Nodes_Node_Summary_Vrf
// Array of VRFs with IPSUB interfaces
type IpSubscriber_Nodes_Node_Summary_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 VRF. The type is string.
    VrfName interface{}

    // IPv6 VRF. The type is string.
    Ipv6vrfName interface{}

    // Number of IP subscriber interfaces in the VRF table. The type is
    // interface{} with range: 0..18446744073709551615.
    Interfaces interface{}

    // Number of IPv6 subscriber interfaces in the VRF table. The type is
    // interface{} with range: 0..18446744073709551615.
    Ipv6Interfaces interface{}
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "summary"
    vrf.EntityData.SegmentPath = "vrf"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("ipv6vrf-name", types.YLeaf{"Ipv6vrfName", vrf.Ipv6vrfName})
    vrf.EntityData.Leafs.Append("interfaces", types.YLeaf{"Interfaces", vrf.Interfaces})
    vrf.EntityData.Leafs.Append("ipv6-interfaces", types.YLeaf{"Ipv6Interfaces", vrf.Ipv6Interfaces})

    vrf.EntityData.YListKeys = []string {}

    return &(vrf.EntityData)
}

// IpSubscriber_Nodes_Node_Interfaces
// IP subscriber interface table
type IpSubscriber_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP subscriber interface entry. The type is slice of
    // IpSubscriber_Nodes_Node_Interfaces_Interface.
    Interface []*IpSubscriber_Nodes_Node_Interfaces_Interface
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
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

// IpSubscriber_Nodes_Node_Interfaces_Interface
// IP subscriber interface entry
type IpSubscriber_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Access interface through which this subscriber is accessible. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    AccessInterface interface{}

    // IPv4 Address of the subscriber. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubscriberIpv4Address interface{}

    // IPv6 Address of the subscriber. The type is string.
    SubscriberIpv6Address interface{}

    // MAC address of the subscriber. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SubscriberMacAddres interface{}

    // Subscriber label for this subscriber interface. The type is interface{}
    // with range: 0..4294967295.
    SubscriberLabel interface{}

    // Interface creation time in month day hh:mm:ss format. The type is string.
    InterfaceCreationTime interface{}

    // Age in hh:mm:ss format. The type is string.
    Age interface{}

    // Protocol trigger for creation of this subscriber session. The type is
    // IpsubMaIntfInitiatorData.
    Initiator interface{}

    // State of the subscriber session. The type is IpsubMaIntfStateData.
    State interface{}

    // Previous state of the subscriber session. The type is IpsubMaIntfStateData.
    OldState interface{}

    // Interface's last state change time in month day hh:mm:ss format. The type
    // is string.
    LastStateChangeTime interface{}

    // Current change age in hh:mm:ss format. The type is string.
    CurrentChangeAge interface{}

    // Protocol trigger for creation of this subscriber's IPv6 session. The type
    // is IpsubMaIntfInitiatorData.
    Ipv6Initiator interface{}

    // State of the subscriber's IPv6 session. The type is IpsubMaIntfStateData.
    Ipv6State interface{}

    // Previous state of the subscriber's IPv6 session. The type is
    // IpsubMaIntfStateData.
    Ipv6OldState interface{}

    // Interface's IPV6 last state change time in month day hh:mm:ss format. The
    // type is string.
    Ipv6LastStateChangeTime interface{}

    // IPV6 Current change age in hh:mm:ss format. The type is string.
    Ipv6CurrentChangeAge interface{}

    // True if L2 connected. The type is bool.
    IsL2Connected interface{}

    // Session Type. The type is string.
    SessionType interface{}

    // IPv4 VRF details.
    Vrf IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf

    // IPv6 VRF details.
    Ipv6vrf IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6vrf
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("vrf", types.YChild{"Vrf", &self.Vrf})
    self.EntityData.Children.Append("ipv6vrf", types.YChild{"Ipv6vrf", &self.Ipv6vrf})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("access-interface", types.YLeaf{"AccessInterface", self.AccessInterface})
    self.EntityData.Leafs.Append("subscriber-ipv4-address", types.YLeaf{"SubscriberIpv4Address", self.SubscriberIpv4Address})
    self.EntityData.Leafs.Append("subscriber-ipv6-address", types.YLeaf{"SubscriberIpv6Address", self.SubscriberIpv6Address})
    self.EntityData.Leafs.Append("subscriber-mac-addres", types.YLeaf{"SubscriberMacAddres", self.SubscriberMacAddres})
    self.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", self.SubscriberLabel})
    self.EntityData.Leafs.Append("interface-creation-time", types.YLeaf{"InterfaceCreationTime", self.InterfaceCreationTime})
    self.EntityData.Leafs.Append("age", types.YLeaf{"Age", self.Age})
    self.EntityData.Leafs.Append("initiator", types.YLeaf{"Initiator", self.Initiator})
    self.EntityData.Leafs.Append("state", types.YLeaf{"State", self.State})
    self.EntityData.Leafs.Append("old-state", types.YLeaf{"OldState", self.OldState})
    self.EntityData.Leafs.Append("last-state-change-time", types.YLeaf{"LastStateChangeTime", self.LastStateChangeTime})
    self.EntityData.Leafs.Append("current-change-age", types.YLeaf{"CurrentChangeAge", self.CurrentChangeAge})
    self.EntityData.Leafs.Append("ipv6-initiator", types.YLeaf{"Ipv6Initiator", self.Ipv6Initiator})
    self.EntityData.Leafs.Append("ipv6-state", types.YLeaf{"Ipv6State", self.Ipv6State})
    self.EntityData.Leafs.Append("ipv6-old-state", types.YLeaf{"Ipv6OldState", self.Ipv6OldState})
    self.EntityData.Leafs.Append("ipv6-last-state-change-time", types.YLeaf{"Ipv6LastStateChangeTime", self.Ipv6LastStateChangeTime})
    self.EntityData.Leafs.Append("ipv6-current-change-age", types.YLeaf{"Ipv6CurrentChangeAge", self.Ipv6CurrentChangeAge})
    self.EntityData.Leafs.Append("is-l2-connected", types.YLeaf{"IsL2Connected", self.IsL2Connected})
    self.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", self.SessionType})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf
// IPv4 VRF details
type IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}

    // Table. The type is string.
    TableName interface{}
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "interface"
    vrf.EntityData.SegmentPath = "vrf"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("table-name", types.YLeaf{"TableName", vrf.TableName})

    vrf.EntityData.YListKeys = []string {}

    return &(vrf.EntityData)
}

// IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6vrf
// IPv6 VRF details
type IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}

    // Table. The type is string.
    TableName interface{}
}

func (ipv6vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6vrf) GetEntityData() *types.CommonEntityData {
    ipv6vrf.EntityData.YFilter = ipv6vrf.YFilter
    ipv6vrf.EntityData.YangName = "ipv6vrf"
    ipv6vrf.EntityData.BundleName = "cisco_ios_xr"
    ipv6vrf.EntityData.ParentYangName = "interface"
    ipv6vrf.EntityData.SegmentPath = "ipv6vrf"
    ipv6vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6vrf.EntityData.Children = types.NewOrderedMap()
    ipv6vrf.EntityData.Leafs = types.NewOrderedMap()
    ipv6vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6vrf.VrfName})
    ipv6vrf.EntityData.Leafs.Append("table-name", types.YLeaf{"TableName", ipv6vrf.TableName})

    ipv6vrf.EntityData.YListKeys = []string {}

    return &(ipv6vrf.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces
// IP subscriber access interface table
type IpSubscriber_Nodes_Node_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP subscriber access interface entry. The type is slice of
    // IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface.
    AccessInterface []*IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetEntityData() *types.CommonEntityData {
    accessInterfaces.EntityData.YFilter = accessInterfaces.YFilter
    accessInterfaces.EntityData.YangName = "access-interfaces"
    accessInterfaces.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaces.EntityData.ParentYangName = "node"
    accessInterfaces.EntityData.SegmentPath = "access-interfaces"
    accessInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaces.EntityData.Children = types.NewOrderedMap()
    accessInterfaces.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children.Append(types.GetSegmentPath(accessInterfaces.AccessInterface[i]), types.YChild{"AccessInterface", accessInterfaces.AccessInterface[i]})
    }
    accessInterfaces.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaces.EntityData.YListKeys = []string {}

    return &(accessInterfaces.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface
// IP subscriber access interface entry
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface creation time in Month Date HH:MM:SS format. The type is string.
    InterfaceCreationTime interface{}

    // Age in HH:MM:SS format. The type is string.
    Age interface{}

    // Interface Type. The type is string.
    InterfaceType interface{}

    // Operational state of this interface. The type is
    // IpsubMaParentIntfStateData.
    State interface{}

    // Operational ipv6 state of this interface. The type is
    // IpsubMaParentIntfStateData.
    Ipv6State interface{}

    // The VLAN type on the access interface. The type is IpsubMaParentIntfVlan.
    VlanType interface{}

    // Configurational state-statistics for each initiating protocol enabled on
    // this parent interface.
    Initiators IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators

    // Configurational state-statistics for each initiating protocol enabled on
    // this parent interface for IPv6 session.
    Ipv6Initiators IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators

    // Configuration session limits for each session limit source and type.
    SessionLimit IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit
}

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "access-interfaces"
    accessInterface.EntityData.SegmentPath = "access-interface" + types.AddKeyToken(accessInterface.InterfaceName, "interface-name")
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Children.Append("initiators", types.YChild{"Initiators", &accessInterface.Initiators})
    accessInterface.EntityData.Children.Append("ipv6-initiators", types.YChild{"Ipv6Initiators", &accessInterface.Ipv6Initiators})
    accessInterface.EntityData.Children.Append("session-limit", types.YChild{"SessionLimit", &accessInterface.SessionLimit})
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("interface-creation-time", types.YLeaf{"InterfaceCreationTime", accessInterface.InterfaceCreationTime})
    accessInterface.EntityData.Leafs.Append("age", types.YLeaf{"Age", accessInterface.Age})
    accessInterface.EntityData.Leafs.Append("interface-type", types.YLeaf{"InterfaceType", accessInterface.InterfaceType})
    accessInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", accessInterface.State})
    accessInterface.EntityData.Leafs.Append("ipv6-state", types.YLeaf{"Ipv6State", accessInterface.Ipv6State})
    accessInterface.EntityData.Leafs.Append("vlan-type", types.YLeaf{"VlanType", accessInterface.VlanType})

    accessInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterface.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators
// Configurational state-statistics for each
// initiating protocol enabled on this parent
// interface
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP information.
    Dhcp IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp

    // packet trigger information.
    PacketTrigger IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetEntityData() *types.CommonEntityData {
    initiators.EntityData.YFilter = initiators.YFilter
    initiators.EntityData.YangName = "initiators"
    initiators.EntityData.BundleName = "cisco_ios_xr"
    initiators.EntityData.ParentYangName = "access-interface"
    initiators.EntityData.SegmentPath = "initiators"
    initiators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initiators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initiators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initiators.EntityData.Children = types.NewOrderedMap()
    initiators.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &initiators.Dhcp})
    initiators.EntityData.Children.Append("packet-trigger", types.YChild{"PacketTrigger", &initiators.PacketTrigger})
    initiators.EntityData.Leafs = types.NewOrderedMap()

    initiators.EntityData.YListKeys = []string {}

    return &(initiators.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp
// DHCP information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ture if the initiator is configred. The type is bool.
    IsConfigured interface{}

    // True if check for subscriber address uniquenessduring first sign of life is
    // enabled. The type is bool.
    UniqueIpCheck interface{}

    // Number of sessions currently up for each initiator. The type is interface{}
    // with range: 0..4294967295.
    Sessions interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface. The type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    FsolBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface that were dropped. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    FsolDroppedBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding creation rate. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsFlow interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding one or more of the
    // configured session limits. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPacketsSessionLimit interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to duplicate source address. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsDupAddr interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xr"
    dhcp.EntityData.ParentYangName = "initiators"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Leafs = types.NewOrderedMap()
    dhcp.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", dhcp.IsConfigured})
    dhcp.EntityData.Leafs.Append("unique-ip-check", types.YLeaf{"UniqueIpCheck", dhcp.UniqueIpCheck})
    dhcp.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", dhcp.Sessions})
    dhcp.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", dhcp.FsolPackets})
    dhcp.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", dhcp.FsolBytes})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets", types.YLeaf{"FsolDroppedPackets", dhcp.FsolDroppedPackets})
    dhcp.EntityData.Leafs.Append("fsol-dropped-bytes", types.YLeaf{"FsolDroppedBytes", dhcp.FsolDroppedBytes})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets-flow", types.YLeaf{"FsolDroppedPacketsFlow", dhcp.FsolDroppedPacketsFlow})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets-session-limit", types.YLeaf{"FsolDroppedPacketsSessionLimit", dhcp.FsolDroppedPacketsSessionLimit})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets-dup-addr", types.YLeaf{"FsolDroppedPacketsDupAddr", dhcp.FsolDroppedPacketsDupAddr})

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger
// packet trigger information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ture if the initiator is configred. The type is bool.
    IsConfigured interface{}

    // True if check for subscriber address uniquenessduring first sign of life is
    // enabled. The type is bool.
    UniqueIpCheck interface{}

    // Number of sessions currently up for each initiator. The type is interface{}
    // with range: 0..4294967295.
    Sessions interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface. The type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    FsolBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface that were dropped. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    FsolDroppedBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding creation rate. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsFlow interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding one or more of the
    // configured session limits. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPacketsSessionLimit interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to duplicate source address. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsDupAddr interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetEntityData() *types.CommonEntityData {
    packetTrigger.EntityData.YFilter = packetTrigger.YFilter
    packetTrigger.EntityData.YangName = "packet-trigger"
    packetTrigger.EntityData.BundleName = "cisco_ios_xr"
    packetTrigger.EntityData.ParentYangName = "initiators"
    packetTrigger.EntityData.SegmentPath = "packet-trigger"
    packetTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetTrigger.EntityData.Children = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", packetTrigger.IsConfigured})
    packetTrigger.EntityData.Leafs.Append("unique-ip-check", types.YLeaf{"UniqueIpCheck", packetTrigger.UniqueIpCheck})
    packetTrigger.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", packetTrigger.Sessions})
    packetTrigger.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", packetTrigger.FsolPackets})
    packetTrigger.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", packetTrigger.FsolBytes})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets", types.YLeaf{"FsolDroppedPackets", packetTrigger.FsolDroppedPackets})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-bytes", types.YLeaf{"FsolDroppedBytes", packetTrigger.FsolDroppedBytes})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets-flow", types.YLeaf{"FsolDroppedPacketsFlow", packetTrigger.FsolDroppedPacketsFlow})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets-session-limit", types.YLeaf{"FsolDroppedPacketsSessionLimit", packetTrigger.FsolDroppedPacketsSessionLimit})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets-dup-addr", types.YLeaf{"FsolDroppedPacketsDupAddr", packetTrigger.FsolDroppedPacketsDupAddr})

    packetTrigger.EntityData.YListKeys = []string {}

    return &(packetTrigger.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators
// Configurational state-statistics for each
// initiating protocol enabled on this parent
// interface for IPv6 session
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP information.
    Dhcp IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp

    // packet trigger information.
    PacketTrigger IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetEntityData() *types.CommonEntityData {
    ipv6Initiators.EntityData.YFilter = ipv6Initiators.YFilter
    ipv6Initiators.EntityData.YangName = "ipv6-initiators"
    ipv6Initiators.EntityData.BundleName = "cisco_ios_xr"
    ipv6Initiators.EntityData.ParentYangName = "access-interface"
    ipv6Initiators.EntityData.SegmentPath = "ipv6-initiators"
    ipv6Initiators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Initiators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Initiators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Initiators.EntityData.Children = types.NewOrderedMap()
    ipv6Initiators.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &ipv6Initiators.Dhcp})
    ipv6Initiators.EntityData.Children.Append("packet-trigger", types.YChild{"PacketTrigger", &ipv6Initiators.PacketTrigger})
    ipv6Initiators.EntityData.Leafs = types.NewOrderedMap()

    ipv6Initiators.EntityData.YListKeys = []string {}

    return &(ipv6Initiators.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp
// DHCP information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ture if the initiator is configred. The type is bool.
    IsConfigured interface{}

    // True if check for subscriber address uniquenessduring first sign of life is
    // enabled. The type is bool.
    UniqueIpCheck interface{}

    // Number of sessions currently up for each initiator. The type is interface{}
    // with range: 0..4294967295.
    Sessions interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface. The type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    FsolBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface that were dropped. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    FsolDroppedBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding creation rate. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsFlow interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding one or more of the
    // configured session limits. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPacketsSessionLimit interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to duplicate source address. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsDupAddr interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xr"
    dhcp.EntityData.ParentYangName = "ipv6-initiators"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Leafs = types.NewOrderedMap()
    dhcp.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", dhcp.IsConfigured})
    dhcp.EntityData.Leafs.Append("unique-ip-check", types.YLeaf{"UniqueIpCheck", dhcp.UniqueIpCheck})
    dhcp.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", dhcp.Sessions})
    dhcp.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", dhcp.FsolPackets})
    dhcp.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", dhcp.FsolBytes})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets", types.YLeaf{"FsolDroppedPackets", dhcp.FsolDroppedPackets})
    dhcp.EntityData.Leafs.Append("fsol-dropped-bytes", types.YLeaf{"FsolDroppedBytes", dhcp.FsolDroppedBytes})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets-flow", types.YLeaf{"FsolDroppedPacketsFlow", dhcp.FsolDroppedPacketsFlow})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets-session-limit", types.YLeaf{"FsolDroppedPacketsSessionLimit", dhcp.FsolDroppedPacketsSessionLimit})
    dhcp.EntityData.Leafs.Append("fsol-dropped-packets-dup-addr", types.YLeaf{"FsolDroppedPacketsDupAddr", dhcp.FsolDroppedPacketsDupAddr})

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger
// packet trigger information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ture if the initiator is configred. The type is bool.
    IsConfigured interface{}

    // True if check for subscriber address uniquenessduring first sign of life is
    // enabled. The type is bool.
    UniqueIpCheck interface{}

    // Number of sessions currently up for each initiator. The type is interface{}
    // with range: 0..4294967295.
    Sessions interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface. The type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    FsolBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPackets interface{}

    // Number of first sign of life bytes received for initiating protocol on this
    // interface that were dropped. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    FsolDroppedBytes interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding creation rate. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsFlow interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to exceeding one or more of the
    // configured session limits. The type is interface{} with range:
    // 0..4294967295.
    FsolDroppedPacketsSessionLimit interface{}

    // Number of first sign of life packets received for initiating protocol on
    // this interface that were dropped due to duplicate source address. The type
    // is interface{} with range: 0..4294967295.
    FsolDroppedPacketsDupAddr interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetEntityData() *types.CommonEntityData {
    packetTrigger.EntityData.YFilter = packetTrigger.YFilter
    packetTrigger.EntityData.YangName = "packet-trigger"
    packetTrigger.EntityData.BundleName = "cisco_ios_xr"
    packetTrigger.EntityData.ParentYangName = "ipv6-initiators"
    packetTrigger.EntityData.SegmentPath = "packet-trigger"
    packetTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetTrigger.EntityData.Children = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs = types.NewOrderedMap()
    packetTrigger.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", packetTrigger.IsConfigured})
    packetTrigger.EntityData.Leafs.Append("unique-ip-check", types.YLeaf{"UniqueIpCheck", packetTrigger.UniqueIpCheck})
    packetTrigger.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", packetTrigger.Sessions})
    packetTrigger.EntityData.Leafs.Append("fsol-packets", types.YLeaf{"FsolPackets", packetTrigger.FsolPackets})
    packetTrigger.EntityData.Leafs.Append("fsol-bytes", types.YLeaf{"FsolBytes", packetTrigger.FsolBytes})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets", types.YLeaf{"FsolDroppedPackets", packetTrigger.FsolDroppedPackets})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-bytes", types.YLeaf{"FsolDroppedBytes", packetTrigger.FsolDroppedBytes})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets-flow", types.YLeaf{"FsolDroppedPacketsFlow", packetTrigger.FsolDroppedPacketsFlow})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets-session-limit", types.YLeaf{"FsolDroppedPacketsSessionLimit", packetTrigger.FsolDroppedPacketsSessionLimit})
    packetTrigger.EntityData.Leafs.Append("fsol-dropped-packets-dup-addr", types.YLeaf{"FsolDroppedPacketsDupAddr", packetTrigger.FsolDroppedPacketsDupAddr})

    packetTrigger.EntityData.YListKeys = []string {}

    return &(packetTrigger.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit
// Configuration session limits for each session
// limit source and type
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unclassified source session limits.
    UnclassifiedSource IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource

    // All sources session limits.
    Total IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetEntityData() *types.CommonEntityData {
    sessionLimit.EntityData.YFilter = sessionLimit.YFilter
    sessionLimit.EntityData.YangName = "session-limit"
    sessionLimit.EntityData.BundleName = "cisco_ios_xr"
    sessionLimit.EntityData.ParentYangName = "access-interface"
    sessionLimit.EntityData.SegmentPath = "session-limit"
    sessionLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionLimit.EntityData.Children = types.NewOrderedMap()
    sessionLimit.EntityData.Children.Append("unclassified-source", types.YChild{"UnclassifiedSource", &sessionLimit.UnclassifiedSource})
    sessionLimit.EntityData.Children.Append("total", types.YChild{"Total", &sessionLimit.Total})
    sessionLimit.EntityData.Leafs = types.NewOrderedMap()

    sessionLimit.EntityData.YListKeys = []string {}

    return &(sessionLimit.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource
// Unclassified source session limits
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-VLAN limit category. The type is interface{} with range: 0..4294967295.
    PerVlan interface{}
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetEntityData() *types.CommonEntityData {
    unclassifiedSource.EntityData.YFilter = unclassifiedSource.YFilter
    unclassifiedSource.EntityData.YangName = "unclassified-source"
    unclassifiedSource.EntityData.BundleName = "cisco_ios_xr"
    unclassifiedSource.EntityData.ParentYangName = "session-limit"
    unclassifiedSource.EntityData.SegmentPath = "unclassified-source"
    unclassifiedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unclassifiedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unclassifiedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unclassifiedSource.EntityData.Children = types.NewOrderedMap()
    unclassifiedSource.EntityData.Leafs = types.NewOrderedMap()
    unclassifiedSource.EntityData.Leafs.Append("per-vlan", types.YLeaf{"PerVlan", unclassifiedSource.PerVlan})

    unclassifiedSource.EntityData.YListKeys = []string {}

    return &(unclassifiedSource.EntityData)
}

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total
// All sources session limits
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-VLAN limit category. The type is interface{} with range: 0..4294967295.
    PerVlan interface{}
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetEntityData() *types.CommonEntityData {
    total.EntityData.YFilter = total.YFilter
    total.EntityData.YangName = "total"
    total.EntityData.BundleName = "cisco_ios_xr"
    total.EntityData.ParentYangName = "session-limit"
    total.EntityData.SegmentPath = "total"
    total.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    total.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    total.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    total.EntityData.Children = types.NewOrderedMap()
    total.EntityData.Leafs = types.NewOrderedMap()
    total.EntityData.Leafs.Append("per-vlan", types.YLeaf{"PerVlan", total.PerVlan})

    total.EntityData.YListKeys = []string {}

    return &(total.EntityData)
}

