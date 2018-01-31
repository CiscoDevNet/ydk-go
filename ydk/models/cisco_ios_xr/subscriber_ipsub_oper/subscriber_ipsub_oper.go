// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-ipsub package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ip-subscriber: IP subscriber operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    IpsubMaIntfStateData_error IpsubMaIntfStateData = "error"
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
    parent types.Entity
    YFilter yfilter.YFilter

    // IP subscriber operational data for a particular location.
    Nodes IpSubscriber_Nodes
}

func (ipSubscriber *IpSubscriber) GetFilter() yfilter.YFilter { return ipSubscriber.YFilter }

func (ipSubscriber *IpSubscriber) SetFilter(yf yfilter.YFilter) { ipSubscriber.YFilter = yf }

func (ipSubscriber *IpSubscriber) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ipSubscriber *IpSubscriber) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-ipsub-oper:ip-subscriber"
}

func (ipSubscriber *IpSubscriber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ipSubscriber.Nodes
    }
    return nil
}

func (ipSubscriber *IpSubscriber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ipSubscriber.Nodes
    return children
}

func (ipSubscriber *IpSubscriber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipSubscriber *IpSubscriber) GetBundleName() string { return "cisco_ios_xr" }

func (ipSubscriber *IpSubscriber) GetYangName() string { return "ip-subscriber" }

func (ipSubscriber *IpSubscriber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipSubscriber *IpSubscriber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipSubscriber *IpSubscriber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipSubscriber *IpSubscriber) SetParent(parent types.Entity) { ipSubscriber.parent = parent }

func (ipSubscriber *IpSubscriber) GetParent() types.Entity { return ipSubscriber.parent }

func (ipSubscriber *IpSubscriber) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-ipsub-oper" }

// IpSubscriber_Nodes
// IP subscriber operational data for a particular
// location
type IpSubscriber_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of IpSubscriber_Nodes_Node.
    Node []IpSubscriber_Nodes_Node
}

func (nodes *IpSubscriber_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *IpSubscriber_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *IpSubscriber_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *IpSubscriber_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *IpSubscriber_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpSubscriber_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *IpSubscriber_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *IpSubscriber_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *IpSubscriber_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *IpSubscriber_Nodes) GetYangName() string { return "nodes" }

func (nodes *IpSubscriber_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *IpSubscriber_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *IpSubscriber_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *IpSubscriber_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *IpSubscriber_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *IpSubscriber_Nodes) GetParentYangName() string { return "ip-subscriber" }

// IpSubscriber_Nodes_Node
// Location. For eg., 0/1/CPU0
type IpSubscriber_Nodes_Node struct {
    parent types.Entity
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

func (node *IpSubscriber_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *IpSubscriber_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *IpSubscriber_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "summary" { return "Summary" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "access-interfaces" { return "AccessInterfaces" }
    return ""
}

func (node *IpSubscriber_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *IpSubscriber_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &node.Summary
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "access-interfaces" {
        return &node.AccessInterfaces
    }
    return nil
}

func (node *IpSubscriber_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &node.Summary
    children["interfaces"] = &node.Interfaces
    children["access-interfaces"] = &node.AccessInterfaces
    return children
}

func (node *IpSubscriber_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *IpSubscriber_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *IpSubscriber_Nodes_Node) GetYangName() string { return "node" }

func (node *IpSubscriber_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *IpSubscriber_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *IpSubscriber_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *IpSubscriber_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *IpSubscriber_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *IpSubscriber_Nodes_Node) GetParentYangName() string { return "nodes" }

// IpSubscriber_Nodes_Node_Summary
// IP subscriber interface summary
type IpSubscriber_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access interface summary statistics.
    AccessInterfaceSummary IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary

    // Initiator interface counts.
    InterfaceCounts IpSubscriber_Nodes_Node_Summary_InterfaceCounts

    // Array of VRFs with IPSUB interfaces. The type is slice of
    // IpSubscriber_Nodes_Node_Summary_Vrf.
    Vrf []IpSubscriber_Nodes_Node_Summary_Vrf
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *IpSubscriber_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *IpSubscriber_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "access-interface-summary" { return "AccessInterfaceSummary" }
    if yname == "interface-counts" { return "InterfaceCounts" }
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface-summary" {
        return &summary.AccessInterfaceSummary
    }
    if childYangName == "interface-counts" {
        return &summary.InterfaceCounts
    }
    if childYangName == "vrf" {
        for _, c := range summary.Vrf {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpSubscriber_Nodes_Node_Summary_Vrf{}
        summary.Vrf = append(summary.Vrf, child)
        return &summary.Vrf[len(summary.Vrf)-1]
    }
    return nil
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-interface-summary"] = &summary.AccessInterfaceSummary
    children["interface-counts"] = &summary.InterfaceCounts
    for i := range summary.Vrf {
        children[summary.Vrf[i].GetSegmentPath()] = &summary.Vrf[i]
    }
    return children
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *IpSubscriber_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *IpSubscriber_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *IpSubscriber_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *IpSubscriber_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *IpSubscriber_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *IpSubscriber_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *IpSubscriber_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *IpSubscriber_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary
// Access interface summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with subscriber configuration. The type is interface{}
    // with range: 0..4294967295.
    Interfaces interface{}

    // Summary counts per initiator.
    Initiators IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators

    // Summary counts per initiator for ipv6 session.
    Ipv6Initiators IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetFilter() yfilter.YFilter { return accessInterfaceSummary.YFilter }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) SetFilter(yf yfilter.YFilter) { accessInterfaceSummary.YFilter = yf }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "initiators" { return "Initiators" }
    if yname == "ipv6-initiators" { return "Ipv6Initiators" }
    return ""
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetSegmentPath() string {
    return "access-interface-summary"
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "initiators" {
        return &accessInterfaceSummary.Initiators
    }
    if childYangName == "ipv6-initiators" {
        return &accessInterfaceSummary.Ipv6Initiators
    }
    return nil
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["initiators"] = &accessInterfaceSummary.Initiators
    children["ipv6-initiators"] = &accessInterfaceSummary.Ipv6Initiators
    return children
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interfaces"] = accessInterfaceSummary.Interfaces
    return leafs
}

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetYangName() string { return "access-interface-summary" }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) SetParent(parent types.Entity) { accessInterfaceSummary.parent = parent }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetParent() types.Entity { return accessInterfaceSummary.parent }

func (accessInterfaceSummary *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary) GetParentYangName() string { return "summary" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators
// Summary counts per initiator
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP summary statistics.
    Dhcp IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp

    // Packet trigger summary statistics.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetFilter() yfilter.YFilter { return initiators.YFilter }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) SetFilter(yf yfilter.YFilter) { initiators.YFilter = yf }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetGoName(yname string) string {
    if yname == "dhcp" { return "Dhcp" }
    if yname == "packet-trigger" { return "PacketTrigger" }
    return ""
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetSegmentPath() string {
    return "initiators"
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp" {
        return &initiators.Dhcp
    }
    if childYangName == "packet-trigger" {
        return &initiators.PacketTrigger
    }
    return nil
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp"] = &initiators.Dhcp
    children["packet-trigger"] = &initiators.PacketTrigger
    return children
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetBundleName() string { return "cisco_ios_xr" }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetYangName() string { return "initiators" }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) SetParent(parent types.Entity) { initiators.parent = parent }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetParent() types.Entity { return initiators.parent }

func (initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators) GetParentYangName() string { return "access-interface-summary" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp
// DHCP summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetFilter() yfilter.YFilter { return dhcp.YFilter }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) SetFilter(yf yfilter.YFilter) { dhcp.YFilter = yf }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetGoName(yname string) string {
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    return ""
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetSegmentPath() string {
    return "dhcp"
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fsol-packets"] = dhcp.FsolPackets
    leafs["fsol-bytes"] = dhcp.FsolBytes
    return leafs
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetBundleName() string { return "cisco_ios_xr" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetYangName() string { return "dhcp" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) SetParent(parent types.Entity) { dhcp.parent = parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetParent() types.Entity { return dhcp.parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_Dhcp) GetParentYangName() string { return "initiators" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger
// Packet trigger summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetFilter() yfilter.YFilter { return packetTrigger.YFilter }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) SetFilter(yf yfilter.YFilter) { packetTrigger.YFilter = yf }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetGoName(yname string) string {
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    return ""
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetSegmentPath() string {
    return "packet-trigger"
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fsol-packets"] = packetTrigger.FsolPackets
    leafs["fsol-bytes"] = packetTrigger.FsolBytes
    return leafs
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetBundleName() string { return "cisco_ios_xr" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetYangName() string { return "packet-trigger" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) SetParent(parent types.Entity) { packetTrigger.parent = parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetParent() types.Entity { return packetTrigger.parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Initiators_PacketTrigger) GetParentYangName() string { return "initiators" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators
// Summary counts per initiator for ipv6 session
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP summary statistics.
    Dhcp IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp

    // Packet trigger summary statistics.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetFilter() yfilter.YFilter { return ipv6Initiators.YFilter }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) SetFilter(yf yfilter.YFilter) { ipv6Initiators.YFilter = yf }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetGoName(yname string) string {
    if yname == "dhcp" { return "Dhcp" }
    if yname == "packet-trigger" { return "PacketTrigger" }
    return ""
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetSegmentPath() string {
    return "ipv6-initiators"
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp" {
        return &ipv6Initiators.Dhcp
    }
    if childYangName == "packet-trigger" {
        return &ipv6Initiators.PacketTrigger
    }
    return nil
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp"] = &ipv6Initiators.Dhcp
    children["packet-trigger"] = &ipv6Initiators.PacketTrigger
    return children
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetYangName() string { return "ipv6-initiators" }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) SetParent(parent types.Entity) { ipv6Initiators.parent = parent }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetParent() types.Entity { return ipv6Initiators.parent }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators) GetParentYangName() string { return "access-interface-summary" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp
// DHCP summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetFilter() yfilter.YFilter { return dhcp.YFilter }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) SetFilter(yf yfilter.YFilter) { dhcp.YFilter = yf }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetGoName(yname string) string {
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    return ""
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetSegmentPath() string {
    return "dhcp"
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fsol-packets"] = dhcp.FsolPackets
    leafs["fsol-bytes"] = dhcp.FsolBytes
    return leafs
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetBundleName() string { return "cisco_ios_xr" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetYangName() string { return "dhcp" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) SetParent(parent types.Entity) { dhcp.parent = parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetParent() types.Entity { return dhcp.parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_Dhcp) GetParentYangName() string { return "ipv6-initiators" }

// IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger
// Packet trigger summary statistics
type IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of first sign of life packets received for initiating protocol. The
    // type is interface{} with range: 0..4294967295.
    FsolPackets interface{}

    // Number of first sign of life bytes received for initiating protocol. The
    // type is interface{} with range: 0..4294967295. Units are byte.
    FsolBytes interface{}
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetFilter() yfilter.YFilter { return packetTrigger.YFilter }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) SetFilter(yf yfilter.YFilter) { packetTrigger.YFilter = yf }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetGoName(yname string) string {
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    return ""
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetSegmentPath() string {
    return "packet-trigger"
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fsol-packets"] = packetTrigger.FsolPackets
    leafs["fsol-bytes"] = packetTrigger.FsolBytes
    return leafs
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetBundleName() string { return "cisco_ios_xr" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetYangName() string { return "packet-trigger" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) SetParent(parent types.Entity) { packetTrigger.parent = parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetParent() types.Entity { return packetTrigger.parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_AccessInterfaceSummary_Ipv6Initiators_PacketTrigger) GetParentYangName() string { return "ipv6-initiators" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts
// Initiator interface counts
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initiators.
    Initiators IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators

    // IPv6 Initiators.
    Ipv6Initiators IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetGoName(yname string) string {
    if yname == "initiators" { return "Initiators" }
    if yname == "ipv6-initiators" { return "Ipv6Initiators" }
    return ""
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "initiators" {
        return &interfaceCounts.Initiators
    }
    if childYangName == "ipv6-initiators" {
        return &interfaceCounts.Ipv6Initiators
    }
    return nil
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["initiators"] = &interfaceCounts.Initiators
    children["ipv6-initiators"] = &interfaceCounts.Ipv6Initiators
    return children
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *IpSubscriber_Nodes_Node_Summary_InterfaceCounts) GetParentYangName() string { return "summary" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators
// Initiators
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP.
    Dhcp IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp

    // Packet trigger.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetFilter() yfilter.YFilter { return initiators.YFilter }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) SetFilter(yf yfilter.YFilter) { initiators.YFilter = yf }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetGoName(yname string) string {
    if yname == "dhcp" { return "Dhcp" }
    if yname == "packet-trigger" { return "PacketTrigger" }
    return ""
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetSegmentPath() string {
    return "initiators"
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp" {
        return &initiators.Dhcp
    }
    if childYangName == "packet-trigger" {
        return &initiators.PacketTrigger
    }
    return nil
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp"] = &initiators.Dhcp
    children["packet-trigger"] = &initiators.PacketTrigger
    return children
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetBundleName() string { return "cisco_ios_xr" }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetYangName() string { return "initiators" }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) SetParent(parent types.Entity) { initiators.parent = parent }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetParent() types.Entity { return initiators.parent }

func (initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators) GetParentYangName() string { return "interface-counts" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp
// DHCP
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp struct {
    parent types.Entity
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

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetFilter() yfilter.YFilter { return dhcp.YFilter }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) SetFilter(yf yfilter.YFilter) { dhcp.YFilter = yf }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetGoName(yname string) string {
    if yname == "invalid" { return "Invalid" }
    if yname == "initialized" { return "Initialized" }
    if yname == "session-creation-started" { return "SessionCreationStarted" }
    if yname == "control-policy-executing" { return "ControlPolicyExecuting" }
    if yname == "control-policy-executed" { return "ControlPolicyExecuted" }
    if yname == "session-features-applied" { return "SessionFeaturesApplied" }
    if yname == "vrf-configured" { return "VrfConfigured" }
    if yname == "adding-adjacency" { return "AddingAdjacency" }
    if yname == "adjacency-added" { return "AdjacencyAdded" }
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    if yname == "disconnecting" { return "Disconnecting" }
    if yname == "disconnected" { return "Disconnected" }
    if yname == "error" { return "Error" }
    if yname == "total-interfaces" { return "TotalInterfaces" }
    return ""
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetSegmentPath() string {
    return "dhcp"
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["invalid"] = dhcp.Invalid
    leafs["initialized"] = dhcp.Initialized
    leafs["session-creation-started"] = dhcp.SessionCreationStarted
    leafs["control-policy-executing"] = dhcp.ControlPolicyExecuting
    leafs["control-policy-executed"] = dhcp.ControlPolicyExecuted
    leafs["session-features-applied"] = dhcp.SessionFeaturesApplied
    leafs["vrf-configured"] = dhcp.VrfConfigured
    leafs["adding-adjacency"] = dhcp.AddingAdjacency
    leafs["adjacency-added"] = dhcp.AdjacencyAdded
    leafs["up"] = dhcp.Up
    leafs["down"] = dhcp.Down
    leafs["disconnecting"] = dhcp.Disconnecting
    leafs["disconnected"] = dhcp.Disconnected
    leafs["error"] = dhcp.Error
    leafs["total-interfaces"] = dhcp.TotalInterfaces
    return leafs
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetBundleName() string { return "cisco_ios_xr" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetYangName() string { return "dhcp" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) SetParent(parent types.Entity) { dhcp.parent = parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetParent() types.Entity { return dhcp.parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_Dhcp) GetParentYangName() string { return "initiators" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger
// Packet trigger
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger struct {
    parent types.Entity
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

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetFilter() yfilter.YFilter { return packetTrigger.YFilter }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) SetFilter(yf yfilter.YFilter) { packetTrigger.YFilter = yf }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetGoName(yname string) string {
    if yname == "invalid" { return "Invalid" }
    if yname == "initialized" { return "Initialized" }
    if yname == "session-creation-started" { return "SessionCreationStarted" }
    if yname == "control-policy-executing" { return "ControlPolicyExecuting" }
    if yname == "control-policy-executed" { return "ControlPolicyExecuted" }
    if yname == "session-features-applied" { return "SessionFeaturesApplied" }
    if yname == "vrf-configured" { return "VrfConfigured" }
    if yname == "adding-adjacency" { return "AddingAdjacency" }
    if yname == "adjacency-added" { return "AdjacencyAdded" }
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    if yname == "disconnecting" { return "Disconnecting" }
    if yname == "disconnected" { return "Disconnected" }
    if yname == "error" { return "Error" }
    if yname == "total-interfaces" { return "TotalInterfaces" }
    return ""
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetSegmentPath() string {
    return "packet-trigger"
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["invalid"] = packetTrigger.Invalid
    leafs["initialized"] = packetTrigger.Initialized
    leafs["session-creation-started"] = packetTrigger.SessionCreationStarted
    leafs["control-policy-executing"] = packetTrigger.ControlPolicyExecuting
    leafs["control-policy-executed"] = packetTrigger.ControlPolicyExecuted
    leafs["session-features-applied"] = packetTrigger.SessionFeaturesApplied
    leafs["vrf-configured"] = packetTrigger.VrfConfigured
    leafs["adding-adjacency"] = packetTrigger.AddingAdjacency
    leafs["adjacency-added"] = packetTrigger.AdjacencyAdded
    leafs["up"] = packetTrigger.Up
    leafs["down"] = packetTrigger.Down
    leafs["disconnecting"] = packetTrigger.Disconnecting
    leafs["disconnected"] = packetTrigger.Disconnected
    leafs["error"] = packetTrigger.Error
    leafs["total-interfaces"] = packetTrigger.TotalInterfaces
    return leafs
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetBundleName() string { return "cisco_ios_xr" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetYangName() string { return "packet-trigger" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) SetParent(parent types.Entity) { packetTrigger.parent = parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetParent() types.Entity { return packetTrigger.parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Initiators_PacketTrigger) GetParentYangName() string { return "initiators" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators
// IPv6 Initiators
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP.
    Dhcp IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp

    // Packet trigger.
    PacketTrigger IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetFilter() yfilter.YFilter { return ipv6Initiators.YFilter }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) SetFilter(yf yfilter.YFilter) { ipv6Initiators.YFilter = yf }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetGoName(yname string) string {
    if yname == "dhcp" { return "Dhcp" }
    if yname == "packet-trigger" { return "PacketTrigger" }
    return ""
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetSegmentPath() string {
    return "ipv6-initiators"
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp" {
        return &ipv6Initiators.Dhcp
    }
    if childYangName == "packet-trigger" {
        return &ipv6Initiators.PacketTrigger
    }
    return nil
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp"] = &ipv6Initiators.Dhcp
    children["packet-trigger"] = &ipv6Initiators.PacketTrigger
    return children
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetYangName() string { return "ipv6-initiators" }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) SetParent(parent types.Entity) { ipv6Initiators.parent = parent }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetParent() types.Entity { return ipv6Initiators.parent }

func (ipv6Initiators *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators) GetParentYangName() string { return "interface-counts" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp
// DHCP
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp struct {
    parent types.Entity
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

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetFilter() yfilter.YFilter { return dhcp.YFilter }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) SetFilter(yf yfilter.YFilter) { dhcp.YFilter = yf }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetGoName(yname string) string {
    if yname == "invalid" { return "Invalid" }
    if yname == "initialized" { return "Initialized" }
    if yname == "session-creation-started" { return "SessionCreationStarted" }
    if yname == "control-policy-executing" { return "ControlPolicyExecuting" }
    if yname == "control-policy-executed" { return "ControlPolicyExecuted" }
    if yname == "session-features-applied" { return "SessionFeaturesApplied" }
    if yname == "vrf-configured" { return "VrfConfigured" }
    if yname == "adding-adjacency" { return "AddingAdjacency" }
    if yname == "adjacency-added" { return "AdjacencyAdded" }
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    if yname == "disconnecting" { return "Disconnecting" }
    if yname == "disconnected" { return "Disconnected" }
    if yname == "error" { return "Error" }
    if yname == "total-interfaces" { return "TotalInterfaces" }
    return ""
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetSegmentPath() string {
    return "dhcp"
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["invalid"] = dhcp.Invalid
    leafs["initialized"] = dhcp.Initialized
    leafs["session-creation-started"] = dhcp.SessionCreationStarted
    leafs["control-policy-executing"] = dhcp.ControlPolicyExecuting
    leafs["control-policy-executed"] = dhcp.ControlPolicyExecuted
    leafs["session-features-applied"] = dhcp.SessionFeaturesApplied
    leafs["vrf-configured"] = dhcp.VrfConfigured
    leafs["adding-adjacency"] = dhcp.AddingAdjacency
    leafs["adjacency-added"] = dhcp.AdjacencyAdded
    leafs["up"] = dhcp.Up
    leafs["down"] = dhcp.Down
    leafs["disconnecting"] = dhcp.Disconnecting
    leafs["disconnected"] = dhcp.Disconnected
    leafs["error"] = dhcp.Error
    leafs["total-interfaces"] = dhcp.TotalInterfaces
    return leafs
}

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetBundleName() string { return "cisco_ios_xr" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetYangName() string { return "dhcp" }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) SetParent(parent types.Entity) { dhcp.parent = parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetParent() types.Entity { return dhcp.parent }

func (dhcp *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_Dhcp) GetParentYangName() string { return "ipv6-initiators" }

// IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger
// Packet trigger
type IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger struct {
    parent types.Entity
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

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetFilter() yfilter.YFilter { return packetTrigger.YFilter }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) SetFilter(yf yfilter.YFilter) { packetTrigger.YFilter = yf }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetGoName(yname string) string {
    if yname == "invalid" { return "Invalid" }
    if yname == "initialized" { return "Initialized" }
    if yname == "session-creation-started" { return "SessionCreationStarted" }
    if yname == "control-policy-executing" { return "ControlPolicyExecuting" }
    if yname == "control-policy-executed" { return "ControlPolicyExecuted" }
    if yname == "session-features-applied" { return "SessionFeaturesApplied" }
    if yname == "vrf-configured" { return "VrfConfigured" }
    if yname == "adding-adjacency" { return "AddingAdjacency" }
    if yname == "adjacency-added" { return "AdjacencyAdded" }
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    if yname == "disconnecting" { return "Disconnecting" }
    if yname == "disconnected" { return "Disconnected" }
    if yname == "error" { return "Error" }
    if yname == "total-interfaces" { return "TotalInterfaces" }
    return ""
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetSegmentPath() string {
    return "packet-trigger"
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["invalid"] = packetTrigger.Invalid
    leafs["initialized"] = packetTrigger.Initialized
    leafs["session-creation-started"] = packetTrigger.SessionCreationStarted
    leafs["control-policy-executing"] = packetTrigger.ControlPolicyExecuting
    leafs["control-policy-executed"] = packetTrigger.ControlPolicyExecuted
    leafs["session-features-applied"] = packetTrigger.SessionFeaturesApplied
    leafs["vrf-configured"] = packetTrigger.VrfConfigured
    leafs["adding-adjacency"] = packetTrigger.AddingAdjacency
    leafs["adjacency-added"] = packetTrigger.AdjacencyAdded
    leafs["up"] = packetTrigger.Up
    leafs["down"] = packetTrigger.Down
    leafs["disconnecting"] = packetTrigger.Disconnecting
    leafs["disconnected"] = packetTrigger.Disconnected
    leafs["error"] = packetTrigger.Error
    leafs["total-interfaces"] = packetTrigger.TotalInterfaces
    return leafs
}

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetBundleName() string { return "cisco_ios_xr" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetYangName() string { return "packet-trigger" }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) SetParent(parent types.Entity) { packetTrigger.parent = parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetParent() types.Entity { return packetTrigger.parent }

func (packetTrigger *IpSubscriber_Nodes_Node_Summary_InterfaceCounts_Ipv6Initiators_PacketTrigger) GetParentYangName() string { return "ipv6-initiators" }

// IpSubscriber_Nodes_Node_Summary_Vrf
// Array of VRFs with IPSUB interfaces
type IpSubscriber_Nodes_Node_Summary_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 VRF. The type is string.
    VrfName interface{}

    // IPv6 VRF. The type is string.
    Ipv6VrfName interface{}

    // Number of IP subscriber interfaces in the VRF table. The type is
    // interface{} with range: 0..18446744073709551615.
    Interfaces interface{}

    // Number of IPv6 subscriber interfaces in the VRF table. The type is
    // interface{} with range: 0..18446744073709551615.
    Ipv6Interfaces interface{}
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6vrf-name" { return "Ipv6VrfName" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "ipv6-interfaces" { return "Ipv6Interfaces" }
    return ""
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetSegmentPath() string {
    return "vrf"
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["ipv6vrf-name"] = vrf.Ipv6VrfName
    leafs["interfaces"] = vrf.Interfaces
    leafs["ipv6-interfaces"] = vrf.Ipv6Interfaces
    return leafs
}

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetYangName() string { return "vrf" }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *IpSubscriber_Nodes_Node_Summary_Vrf) GetParentYangName() string { return "summary" }

// IpSubscriber_Nodes_Node_Interfaces
// IP subscriber interface table
type IpSubscriber_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP subscriber interface entry. The type is slice of
    // IpSubscriber_Nodes_Node_Interfaces_Interface.
    Interface []IpSubscriber_Nodes_Node_Interfaces_Interface
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpSubscriber_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *IpSubscriber_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// IpSubscriber_Nodes_Node_Interfaces_Interface
// IP subscriber interface entry
type IpSubscriber_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Access interface through which this subscriber is accessible. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
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
    Ipv6Vrf IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "access-interface" { return "AccessInterface" }
    if yname == "subscriber-ipv4-address" { return "SubscriberIpv4Address" }
    if yname == "subscriber-ipv6-address" { return "SubscriberIpv6Address" }
    if yname == "subscriber-mac-addres" { return "SubscriberMacAddres" }
    if yname == "subscriber-label" { return "SubscriberLabel" }
    if yname == "interface-creation-time" { return "InterfaceCreationTime" }
    if yname == "age" { return "Age" }
    if yname == "initiator" { return "Initiator" }
    if yname == "state" { return "State" }
    if yname == "old-state" { return "OldState" }
    if yname == "last-state-change-time" { return "LastStateChangeTime" }
    if yname == "current-change-age" { return "CurrentChangeAge" }
    if yname == "ipv6-initiator" { return "Ipv6Initiator" }
    if yname == "ipv6-state" { return "Ipv6State" }
    if yname == "ipv6-old-state" { return "Ipv6OldState" }
    if yname == "ipv6-last-state-change-time" { return "Ipv6LastStateChangeTime" }
    if yname == "ipv6-current-change-age" { return "Ipv6CurrentChangeAge" }
    if yname == "is-l2-connected" { return "IsL2Connected" }
    if yname == "session-type" { return "SessionType" }
    if yname == "vrf" { return "Vrf" }
    if yname == "ipv6vrf" { return "Ipv6Vrf" }
    return ""
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        return &self.Vrf
    }
    if childYangName == "ipv6vrf" {
        return &self.Ipv6Vrf
    }
    return nil
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf"] = &self.Vrf
    children["ipv6vrf"] = &self.Ipv6Vrf
    return children
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["access-interface"] = self.AccessInterface
    leafs["subscriber-ipv4-address"] = self.SubscriberIpv4Address
    leafs["subscriber-ipv6-address"] = self.SubscriberIpv6Address
    leafs["subscriber-mac-addres"] = self.SubscriberMacAddres
    leafs["subscriber-label"] = self.SubscriberLabel
    leafs["interface-creation-time"] = self.InterfaceCreationTime
    leafs["age"] = self.Age
    leafs["initiator"] = self.Initiator
    leafs["state"] = self.State
    leafs["old-state"] = self.OldState
    leafs["last-state-change-time"] = self.LastStateChangeTime
    leafs["current-change-age"] = self.CurrentChangeAge
    leafs["ipv6-initiator"] = self.Ipv6Initiator
    leafs["ipv6-state"] = self.Ipv6State
    leafs["ipv6-old-state"] = self.Ipv6OldState
    leafs["ipv6-last-state-change-time"] = self.Ipv6LastStateChangeTime
    leafs["ipv6-current-change-age"] = self.Ipv6CurrentChangeAge
    leafs["is-l2-connected"] = self.IsL2Connected
    leafs["session-type"] = self.SessionType
    return leafs
}

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *IpSubscriber_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf
// IPv4 VRF details
type IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}

    // Table. The type is string.
    TableName interface{}
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "table-name" { return "TableName" }
    return ""
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetSegmentPath() string {
    return "vrf"
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["table-name"] = vrf.TableName
    return leafs
}

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetYangName() string { return "vrf" }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Vrf) GetParentYangName() string { return "interface" }

// IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf
// IPv6 VRF details
type IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}

    // Table. The type is string.
    TableName interface{}
}

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetFilter() yfilter.YFilter { return ipv6Vrf.YFilter }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) SetFilter(yf yfilter.YFilter) { ipv6Vrf.YFilter = yf }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "table-name" { return "TableName" }
    return ""
}

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetSegmentPath() string {
    return "ipv6vrf"
}

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv6Vrf.VrfName
    leafs["table-name"] = ipv6Vrf.TableName
    return leafs
}

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetYangName() string { return "ipv6vrf" }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) SetParent(parent types.Entity) { ipv6Vrf.parent = parent }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetParent() types.Entity { return ipv6Vrf.parent }

func (ipv6Vrf *IpSubscriber_Nodes_Node_Interfaces_Interface_Ipv6Vrf) GetParentYangName() string { return "interface" }

// IpSubscriber_Nodes_Node_AccessInterfaces
// IP subscriber access interface table
type IpSubscriber_Nodes_Node_AccessInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP subscriber access interface entry. The type is slice of
    // IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface.
    AccessInterface []IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetFilter() yfilter.YFilter { return accessInterfaces.YFilter }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) SetFilter(yf yfilter.YFilter) { accessInterfaces.YFilter = yf }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetGoName(yname string) string {
    if yname == "access-interface" { return "AccessInterface" }
    return ""
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetSegmentPath() string {
    return "access-interfaces"
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface" {
        for _, c := range accessInterfaces.AccessInterface {
            if accessInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface{}
        accessInterfaces.AccessInterface = append(accessInterfaces.AccessInterface, child)
        return &accessInterfaces.AccessInterface[len(accessInterfaces.AccessInterface)-1]
    }
    return nil
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessInterfaces.AccessInterface {
        children[accessInterfaces.AccessInterface[i].GetSegmentPath()] = &accessInterfaces.AccessInterface[i]
    }
    return children
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetYangName() string { return "access-interfaces" }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) SetParent(parent types.Entity) { accessInterfaces.parent = parent }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetParent() types.Entity { return accessInterfaces.parent }

func (accessInterfaces *IpSubscriber_Nodes_Node_AccessInterfaces) GetParentYangName() string { return "node" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface
// IP subscriber access interface entry
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetFilter() yfilter.YFilter { return accessInterface.YFilter }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) SetFilter(yf yfilter.YFilter) { accessInterface.YFilter = yf }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-creation-time" { return "InterfaceCreationTime" }
    if yname == "age" { return "Age" }
    if yname == "interface-type" { return "InterfaceType" }
    if yname == "state" { return "State" }
    if yname == "ipv6-state" { return "Ipv6State" }
    if yname == "vlan-type" { return "VlanType" }
    if yname == "initiators" { return "Initiators" }
    if yname == "ipv6-initiators" { return "Ipv6Initiators" }
    if yname == "session-limit" { return "SessionLimit" }
    return ""
}

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetSegmentPath() string {
    return "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
}

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "initiators" {
        return &accessInterface.Initiators
    }
    if childYangName == "ipv6-initiators" {
        return &accessInterface.Ipv6Initiators
    }
    if childYangName == "session-limit" {
        return &accessInterface.SessionLimit
    }
    return nil
}

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["initiators"] = &accessInterface.Initiators
    children["ipv6-initiators"] = &accessInterface.Ipv6Initiators
    children["session-limit"] = &accessInterface.SessionLimit
    return children
}

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = accessInterface.InterfaceName
    leafs["interface-creation-time"] = accessInterface.InterfaceCreationTime
    leafs["age"] = accessInterface.Age
    leafs["interface-type"] = accessInterface.InterfaceType
    leafs["state"] = accessInterface.State
    leafs["ipv6-state"] = accessInterface.Ipv6State
    leafs["vlan-type"] = accessInterface.VlanType
    return leafs
}

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetYangName() string { return "access-interface" }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) SetParent(parent types.Entity) { accessInterface.parent = parent }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetParent() types.Entity { return accessInterface.parent }

func (accessInterface *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface) GetParentYangName() string { return "access-interfaces" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators
// Configurational state-statistics for each
// initiating protocol enabled on this parent
// interface
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP information.
    Dhcp IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp

    // packet trigger information.
    PacketTrigger IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetFilter() yfilter.YFilter { return initiators.YFilter }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) SetFilter(yf yfilter.YFilter) { initiators.YFilter = yf }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetGoName(yname string) string {
    if yname == "dhcp" { return "Dhcp" }
    if yname == "packet-trigger" { return "PacketTrigger" }
    return ""
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetSegmentPath() string {
    return "initiators"
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp" {
        return &initiators.Dhcp
    }
    if childYangName == "packet-trigger" {
        return &initiators.PacketTrigger
    }
    return nil
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp"] = &initiators.Dhcp
    children["packet-trigger"] = &initiators.PacketTrigger
    return children
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetBundleName() string { return "cisco_ios_xr" }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetYangName() string { return "initiators" }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) SetParent(parent types.Entity) { initiators.parent = parent }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetParent() types.Entity { return initiators.parent }

func (initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators) GetParentYangName() string { return "access-interface" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp
// DHCP information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp struct {
    parent types.Entity
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

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetFilter() yfilter.YFilter { return dhcp.YFilter }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) SetFilter(yf yfilter.YFilter) { dhcp.YFilter = yf }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetGoName(yname string) string {
    if yname == "is-configured" { return "IsConfigured" }
    if yname == "unique-ip-check" { return "UniqueIpCheck" }
    if yname == "sessions" { return "Sessions" }
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    if yname == "fsol-dropped-packets" { return "FsolDroppedPackets" }
    if yname == "fsol-dropped-bytes" { return "FsolDroppedBytes" }
    if yname == "fsol-dropped-packets-flow" { return "FsolDroppedPacketsFlow" }
    if yname == "fsol-dropped-packets-session-limit" { return "FsolDroppedPacketsSessionLimit" }
    if yname == "fsol-dropped-packets-dup-addr" { return "FsolDroppedPacketsDupAddr" }
    return ""
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetSegmentPath() string {
    return "dhcp"
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-configured"] = dhcp.IsConfigured
    leafs["unique-ip-check"] = dhcp.UniqueIpCheck
    leafs["sessions"] = dhcp.Sessions
    leafs["fsol-packets"] = dhcp.FsolPackets
    leafs["fsol-bytes"] = dhcp.FsolBytes
    leafs["fsol-dropped-packets"] = dhcp.FsolDroppedPackets
    leafs["fsol-dropped-bytes"] = dhcp.FsolDroppedBytes
    leafs["fsol-dropped-packets-flow"] = dhcp.FsolDroppedPacketsFlow
    leafs["fsol-dropped-packets-session-limit"] = dhcp.FsolDroppedPacketsSessionLimit
    leafs["fsol-dropped-packets-dup-addr"] = dhcp.FsolDroppedPacketsDupAddr
    return leafs
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetBundleName() string { return "cisco_ios_xr" }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetYangName() string { return "dhcp" }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) SetParent(parent types.Entity) { dhcp.parent = parent }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetParent() types.Entity { return dhcp.parent }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_Dhcp) GetParentYangName() string { return "initiators" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger
// packet trigger information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger struct {
    parent types.Entity
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

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetFilter() yfilter.YFilter { return packetTrigger.YFilter }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) SetFilter(yf yfilter.YFilter) { packetTrigger.YFilter = yf }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetGoName(yname string) string {
    if yname == "is-configured" { return "IsConfigured" }
    if yname == "unique-ip-check" { return "UniqueIpCheck" }
    if yname == "sessions" { return "Sessions" }
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    if yname == "fsol-dropped-packets" { return "FsolDroppedPackets" }
    if yname == "fsol-dropped-bytes" { return "FsolDroppedBytes" }
    if yname == "fsol-dropped-packets-flow" { return "FsolDroppedPacketsFlow" }
    if yname == "fsol-dropped-packets-session-limit" { return "FsolDroppedPacketsSessionLimit" }
    if yname == "fsol-dropped-packets-dup-addr" { return "FsolDroppedPacketsDupAddr" }
    return ""
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetSegmentPath() string {
    return "packet-trigger"
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-configured"] = packetTrigger.IsConfigured
    leafs["unique-ip-check"] = packetTrigger.UniqueIpCheck
    leafs["sessions"] = packetTrigger.Sessions
    leafs["fsol-packets"] = packetTrigger.FsolPackets
    leafs["fsol-bytes"] = packetTrigger.FsolBytes
    leafs["fsol-dropped-packets"] = packetTrigger.FsolDroppedPackets
    leafs["fsol-dropped-bytes"] = packetTrigger.FsolDroppedBytes
    leafs["fsol-dropped-packets-flow"] = packetTrigger.FsolDroppedPacketsFlow
    leafs["fsol-dropped-packets-session-limit"] = packetTrigger.FsolDroppedPacketsSessionLimit
    leafs["fsol-dropped-packets-dup-addr"] = packetTrigger.FsolDroppedPacketsDupAddr
    return leafs
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetBundleName() string { return "cisco_ios_xr" }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetYangName() string { return "packet-trigger" }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) SetParent(parent types.Entity) { packetTrigger.parent = parent }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetParent() types.Entity { return packetTrigger.parent }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Initiators_PacketTrigger) GetParentYangName() string { return "initiators" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators
// Configurational state-statistics for each
// initiating protocol enabled on this parent
// interface for IPv6 session
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP information.
    Dhcp IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp

    // packet trigger information.
    PacketTrigger IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetFilter() yfilter.YFilter { return ipv6Initiators.YFilter }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) SetFilter(yf yfilter.YFilter) { ipv6Initiators.YFilter = yf }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetGoName(yname string) string {
    if yname == "dhcp" { return "Dhcp" }
    if yname == "packet-trigger" { return "PacketTrigger" }
    return ""
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetSegmentPath() string {
    return "ipv6-initiators"
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp" {
        return &ipv6Initiators.Dhcp
    }
    if childYangName == "packet-trigger" {
        return &ipv6Initiators.PacketTrigger
    }
    return nil
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp"] = &ipv6Initiators.Dhcp
    children["packet-trigger"] = &ipv6Initiators.PacketTrigger
    return children
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetYangName() string { return "ipv6-initiators" }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) SetParent(parent types.Entity) { ipv6Initiators.parent = parent }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetParent() types.Entity { return ipv6Initiators.parent }

func (ipv6Initiators *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators) GetParentYangName() string { return "access-interface" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp
// DHCP information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp struct {
    parent types.Entity
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

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetFilter() yfilter.YFilter { return dhcp.YFilter }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) SetFilter(yf yfilter.YFilter) { dhcp.YFilter = yf }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetGoName(yname string) string {
    if yname == "is-configured" { return "IsConfigured" }
    if yname == "unique-ip-check" { return "UniqueIpCheck" }
    if yname == "sessions" { return "Sessions" }
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    if yname == "fsol-dropped-packets" { return "FsolDroppedPackets" }
    if yname == "fsol-dropped-bytes" { return "FsolDroppedBytes" }
    if yname == "fsol-dropped-packets-flow" { return "FsolDroppedPacketsFlow" }
    if yname == "fsol-dropped-packets-session-limit" { return "FsolDroppedPacketsSessionLimit" }
    if yname == "fsol-dropped-packets-dup-addr" { return "FsolDroppedPacketsDupAddr" }
    return ""
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetSegmentPath() string {
    return "dhcp"
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-configured"] = dhcp.IsConfigured
    leafs["unique-ip-check"] = dhcp.UniqueIpCheck
    leafs["sessions"] = dhcp.Sessions
    leafs["fsol-packets"] = dhcp.FsolPackets
    leafs["fsol-bytes"] = dhcp.FsolBytes
    leafs["fsol-dropped-packets"] = dhcp.FsolDroppedPackets
    leafs["fsol-dropped-bytes"] = dhcp.FsolDroppedBytes
    leafs["fsol-dropped-packets-flow"] = dhcp.FsolDroppedPacketsFlow
    leafs["fsol-dropped-packets-session-limit"] = dhcp.FsolDroppedPacketsSessionLimit
    leafs["fsol-dropped-packets-dup-addr"] = dhcp.FsolDroppedPacketsDupAddr
    return leafs
}

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetBundleName() string { return "cisco_ios_xr" }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetYangName() string { return "dhcp" }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) SetParent(parent types.Entity) { dhcp.parent = parent }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetParent() types.Entity { return dhcp.parent }

func (dhcp *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_Dhcp) GetParentYangName() string { return "ipv6-initiators" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger
// packet trigger information
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger struct {
    parent types.Entity
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

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetFilter() yfilter.YFilter { return packetTrigger.YFilter }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) SetFilter(yf yfilter.YFilter) { packetTrigger.YFilter = yf }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetGoName(yname string) string {
    if yname == "is-configured" { return "IsConfigured" }
    if yname == "unique-ip-check" { return "UniqueIpCheck" }
    if yname == "sessions" { return "Sessions" }
    if yname == "fsol-packets" { return "FsolPackets" }
    if yname == "fsol-bytes" { return "FsolBytes" }
    if yname == "fsol-dropped-packets" { return "FsolDroppedPackets" }
    if yname == "fsol-dropped-bytes" { return "FsolDroppedBytes" }
    if yname == "fsol-dropped-packets-flow" { return "FsolDroppedPacketsFlow" }
    if yname == "fsol-dropped-packets-session-limit" { return "FsolDroppedPacketsSessionLimit" }
    if yname == "fsol-dropped-packets-dup-addr" { return "FsolDroppedPacketsDupAddr" }
    return ""
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetSegmentPath() string {
    return "packet-trigger"
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-configured"] = packetTrigger.IsConfigured
    leafs["unique-ip-check"] = packetTrigger.UniqueIpCheck
    leafs["sessions"] = packetTrigger.Sessions
    leafs["fsol-packets"] = packetTrigger.FsolPackets
    leafs["fsol-bytes"] = packetTrigger.FsolBytes
    leafs["fsol-dropped-packets"] = packetTrigger.FsolDroppedPackets
    leafs["fsol-dropped-bytes"] = packetTrigger.FsolDroppedBytes
    leafs["fsol-dropped-packets-flow"] = packetTrigger.FsolDroppedPacketsFlow
    leafs["fsol-dropped-packets-session-limit"] = packetTrigger.FsolDroppedPacketsSessionLimit
    leafs["fsol-dropped-packets-dup-addr"] = packetTrigger.FsolDroppedPacketsDupAddr
    return leafs
}

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetBundleName() string { return "cisco_ios_xr" }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetYangName() string { return "packet-trigger" }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) SetParent(parent types.Entity) { packetTrigger.parent = parent }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetParent() types.Entity { return packetTrigger.parent }

func (packetTrigger *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_Ipv6Initiators_PacketTrigger) GetParentYangName() string { return "ipv6-initiators" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit
// Configuration session limits for each session
// limit source and type
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unclassified source session limits.
    UnclassifiedSource IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource

    // All sources session limits.
    Total IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetFilter() yfilter.YFilter { return sessionLimit.YFilter }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) SetFilter(yf yfilter.YFilter) { sessionLimit.YFilter = yf }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetGoName(yname string) string {
    if yname == "unclassified-source" { return "UnclassifiedSource" }
    if yname == "total" { return "Total" }
    return ""
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetSegmentPath() string {
    return "session-limit"
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unclassified-source" {
        return &sessionLimit.UnclassifiedSource
    }
    if childYangName == "total" {
        return &sessionLimit.Total
    }
    return nil
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unclassified-source"] = &sessionLimit.UnclassifiedSource
    children["total"] = &sessionLimit.Total
    return children
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetBundleName() string { return "cisco_ios_xr" }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetYangName() string { return "session-limit" }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) SetParent(parent types.Entity) { sessionLimit.parent = parent }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetParent() types.Entity { return sessionLimit.parent }

func (sessionLimit *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit) GetParentYangName() string { return "access-interface" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource
// Unclassified source session limits
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-VLAN limit category. The type is interface{} with range: 0..4294967295.
    PerVlan interface{}
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetFilter() yfilter.YFilter { return unclassifiedSource.YFilter }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) SetFilter(yf yfilter.YFilter) { unclassifiedSource.YFilter = yf }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetGoName(yname string) string {
    if yname == "per-vlan" { return "PerVlan" }
    return ""
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetSegmentPath() string {
    return "unclassified-source"
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["per-vlan"] = unclassifiedSource.PerVlan
    return leafs
}

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetBundleName() string { return "cisco_ios_xr" }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetYangName() string { return "unclassified-source" }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) SetParent(parent types.Entity) { unclassifiedSource.parent = parent }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetParent() types.Entity { return unclassifiedSource.parent }

func (unclassifiedSource *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_UnclassifiedSource) GetParentYangName() string { return "session-limit" }

// IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total
// All sources session limits
type IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-VLAN limit category. The type is interface{} with range: 0..4294967295.
    PerVlan interface{}
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetFilter() yfilter.YFilter { return total.YFilter }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) SetFilter(yf yfilter.YFilter) { total.YFilter = yf }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetGoName(yname string) string {
    if yname == "per-vlan" { return "PerVlan" }
    return ""
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetSegmentPath() string {
    return "total"
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["per-vlan"] = total.PerVlan
    return leafs
}

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetBundleName() string { return "cisco_ios_xr" }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetYangName() string { return "total" }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) SetParent(parent types.Entity) { total.parent = parent }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetParent() types.Entity { return total.parent }

func (total *IpSubscriber_Nodes_Node_AccessInterfaces_AccessInterface_SessionLimit_Total) GetParentYangName() string { return "session-limit" }

