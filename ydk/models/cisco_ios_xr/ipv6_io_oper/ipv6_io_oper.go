// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-io package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-io: IPv6 IO Operational Data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_io_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_io_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-io-oper ipv6-io}", reflect.TypeOf(Ipv6Io{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-io-oper:ipv6-io", reflect.TypeOf(Ipv6Io{}))
}

// Ipv6Io
// IPv6 IO Operational Data
type Ipv6Io struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific IPv6 IO operational data.
    Nodes Ipv6Io_Nodes
}

func (ipv6Io *Ipv6Io) GetFilter() yfilter.YFilter { return ipv6Io.YFilter }

func (ipv6Io *Ipv6Io) SetFilter(yf yfilter.YFilter) { ipv6Io.YFilter = yf }

func (ipv6Io *Ipv6Io) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ipv6Io *Ipv6Io) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-io-oper:ipv6-io"
}

func (ipv6Io *Ipv6Io) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ipv6Io.Nodes
    }
    return nil
}

func (ipv6Io *Ipv6Io) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ipv6Io.Nodes
    return children
}

func (ipv6Io *Ipv6Io) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Io *Ipv6Io) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Io *Ipv6Io) GetYangName() string { return "ipv6-io" }

func (ipv6Io *Ipv6Io) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Io *Ipv6Io) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Io *Ipv6Io) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Io *Ipv6Io) SetParent(parent types.Entity) { ipv6Io.parent = parent }

func (ipv6Io *Ipv6Io) GetParent() types.Entity { return ipv6Io.parent }

func (ipv6Io *Ipv6Io) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-io-oper" }

// Ipv6Io_Nodes
// Node-specific IPv6 IO operational data
type Ipv6Io_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 network operational data for a particular node. The type is slice of
    // Ipv6Io_Nodes_Node.
    Node []Ipv6Io_Nodes_Node
}

func (nodes *Ipv6Io_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ipv6Io_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ipv6Io_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ipv6Io_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ipv6Io_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Io_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ipv6Io_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ipv6Io_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ipv6Io_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ipv6Io_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ipv6Io_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ipv6Io_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ipv6Io_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ipv6Io_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ipv6Io_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ipv6Io_Nodes) GetParentYangName() string { return "ipv6-io" }

// Ipv6Io_Nodes_Node
// IPv6 network operational data for a particular
// node
type Ipv6Io_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Statistical IPv6 network operational data for a node.
    Statistics Ipv6Io_Nodes_Node_Statistics
}

func (node *Ipv6Io_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ipv6Io_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ipv6Io_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *Ipv6Io_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ipv6Io_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *Ipv6Io_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    return children
}

func (node *Ipv6Io_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ipv6Io_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ipv6Io_Nodes_Node) GetYangName() string { return "node" }

func (node *Ipv6Io_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ipv6Io_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ipv6Io_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ipv6Io_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ipv6Io_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ipv6Io_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ipv6Io_Nodes_Node_Statistics
// Statistical IPv6 network operational data for
// a node
type Ipv6Io_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Traffic statistics for a node.
    Traffic Ipv6Io_Nodes_Node_Statistics_Traffic
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv6Io_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "traffic" { return "Traffic" }
    return ""
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &statistics.Traffic
    }
    return nil
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &statistics.Traffic
    return children
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv6Io_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Ipv6Io_Nodes_Node_Statistics_Traffic
// Traffic statistics for a node
type Ipv6Io_Nodes_Node_Statistics_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Statistics.
    Ipv6 Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6

    // ICMP Statistics.
    Icmp Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp

    // IPv6 Node Discovery Statistics.
    Ipv6NodeDiscovery Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "icmp" { return "Icmp" }
    if yname == "ipv6-node-discovery" { return "Ipv6NodeDiscovery" }
    return ""
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &traffic.Ipv6
    }
    if childYangName == "icmp" {
        return &traffic.Icmp
    }
    if childYangName == "ipv6-node-discovery" {
        return &traffic.Ipv6NodeDiscovery
    }
    return nil
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &traffic.Ipv6
    children["icmp"] = &traffic.Icmp
    children["ipv6-node-discovery"] = &traffic.Ipv6NodeDiscovery
    return children
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetYangName() string { return "traffic" }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetParentYangName() string { return "statistics" }

// Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6
// IPv6 Statistics
type Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total Packets. The type is interface{} with range: 0..4294967295.
    TotalPackets interface{}

    // Local Destination Packets. The type is interface{} with range:
    // 0..4294967295.
    LocalDestinationPackets interface{}

    // Format Errors. The type is interface{} with range: 0..4294967295.
    FormatErrors interface{}

    // Truncated Packets. The type is interface{} with range: 0..4294967295.
    TruncatedPackets interface{}

    // Hop Count Exceeded Packets. The type is interface{} with range:
    // 0..4294967295.
    HopCountExceededPackets interface{}

    // Bad Source Address Packets. The type is interface{} with range:
    // 0..4294967295.
    BadSourceAddressPackets interface{}

    // Bad Header Packets. The type is interface{} with range: 0..4294967295.
    BadHeaderPackets interface{}

    // Unknown Option Type Packets. The type is interface{} with range:
    // 0..4294967295.
    UnknownOptionTypePackets interface{}

    // Unknown Protocol Packets. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPackets interface{}

    // Fragments. The type is interface{} with range: 0..4294967295.
    Fragments interface{}

    // Reassembled Packets. The type is interface{} with range: 0..4294967295.
    ReassembledPackets interface{}

    // Reassembly Timeouts. The type is interface{} with range: 0..4294967295.
    ReassemblyTimeouts interface{}

    // Reassembly Failures. The type is interface{} with range: 0..4294967295.
    ReassemblyFailures interface{}

    // Reassembly Reach Maximum Drop. The type is interface{} with range:
    // 0..4294967295.
    ReassemblyMaximumDrops interface{}

    // Packets Output. The type is interface{} with range: 0..4294967295.
    GeneratedPackets interface{}

    // Packets Forwarded. The type is interface{} with range: 0..4294967295.
    ForwardedPackets interface{}

    // Packets Source Routed. The type is interface{} with range: 0..4294967295.
    SourceRoutedPackets interface{}

    // Packets Fragmented. The type is interface{} with range: 0..4294967295.
    FragmentedPackets interface{}

    // Fragmented Packet Count. The type is interface{} with range: 0..4294967295.
    FragmentCount interface{}

    // Fragment Failures. The type is interface{} with range: 0..4294967295.
    FragmentFailures interface{}

    // No Route Packets. The type is interface{} with range: 0..4294967295.
    NoRoutePackets interface{}

    // Packet Too Big. The type is interface{} with range: 0..4294967295.
    TooBigPackets interface{}

    // Multicast In. The type is interface{} with range: 0..4294967295.
    ReceivedMulticastPackets interface{}

    // Multicast Out. The type is interface{} with range: 0..4294967295.
    SentMulticastPackets interface{}

    // Misc. drops. The type is interface{} with range: 0..4294967295.
    MiscellaneousDrops interface{}

    // Lisp IPv4 Encapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV4EncapPackets interface{}

    // Lisp IPv4 Decapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV4DecapPackets interface{}

    // Lisp IPv6 Encapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV6EncapPackets interface{}

    // Lisp IPv6 Decapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV6DecapPackets interface{}

    // Lisp Encap errors. The type is interface{} with range: 0..4294967295.
    LispEncapErrors interface{}

    // Lisp Decap errors. The type is interface{} with range: 0..4294967295.
    LispDecapErrors interface{}
}

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetGoName(yname string) string {
    if yname == "total-packets" { return "TotalPackets" }
    if yname == "local-destination-packets" { return "LocalDestinationPackets" }
    if yname == "format-errors" { return "FormatErrors" }
    if yname == "truncated-packets" { return "TruncatedPackets" }
    if yname == "hop-count-exceeded-packets" { return "HopCountExceededPackets" }
    if yname == "bad-source-address-packets" { return "BadSourceAddressPackets" }
    if yname == "bad-header-packets" { return "BadHeaderPackets" }
    if yname == "unknown-option-type-packets" { return "UnknownOptionTypePackets" }
    if yname == "unknown-protocol-packets" { return "UnknownProtocolPackets" }
    if yname == "fragments" { return "Fragments" }
    if yname == "reassembled-packets" { return "ReassembledPackets" }
    if yname == "reassembly-timeouts" { return "ReassemblyTimeouts" }
    if yname == "reassembly-failures" { return "ReassemblyFailures" }
    if yname == "reassembly-maximum-drops" { return "ReassemblyMaximumDrops" }
    if yname == "generated-packets" { return "GeneratedPackets" }
    if yname == "forwarded-packets" { return "ForwardedPackets" }
    if yname == "source-routed-packets" { return "SourceRoutedPackets" }
    if yname == "fragmented-packets" { return "FragmentedPackets" }
    if yname == "fragment-count" { return "FragmentCount" }
    if yname == "fragment-failures" { return "FragmentFailures" }
    if yname == "no-route-packets" { return "NoRoutePackets" }
    if yname == "too-big-packets" { return "TooBigPackets" }
    if yname == "received-multicast-packets" { return "ReceivedMulticastPackets" }
    if yname == "sent-multicast-packets" { return "SentMulticastPackets" }
    if yname == "miscellaneous-drops" { return "MiscellaneousDrops" }
    if yname == "lisp-v4-encap-packets" { return "LispV4EncapPackets" }
    if yname == "lisp-v4-decap-packets" { return "LispV4DecapPackets" }
    if yname == "lisp-v6-encap-packets" { return "LispV6EncapPackets" }
    if yname == "lisp-v6-decap-packets" { return "LispV6DecapPackets" }
    if yname == "lisp-encap-errors" { return "LispEncapErrors" }
    if yname == "lisp-decap-errors" { return "LispDecapErrors" }
    return ""
}

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-packets"] = ipv6.TotalPackets
    leafs["local-destination-packets"] = ipv6.LocalDestinationPackets
    leafs["format-errors"] = ipv6.FormatErrors
    leafs["truncated-packets"] = ipv6.TruncatedPackets
    leafs["hop-count-exceeded-packets"] = ipv6.HopCountExceededPackets
    leafs["bad-source-address-packets"] = ipv6.BadSourceAddressPackets
    leafs["bad-header-packets"] = ipv6.BadHeaderPackets
    leafs["unknown-option-type-packets"] = ipv6.UnknownOptionTypePackets
    leafs["unknown-protocol-packets"] = ipv6.UnknownProtocolPackets
    leafs["fragments"] = ipv6.Fragments
    leafs["reassembled-packets"] = ipv6.ReassembledPackets
    leafs["reassembly-timeouts"] = ipv6.ReassemblyTimeouts
    leafs["reassembly-failures"] = ipv6.ReassemblyFailures
    leafs["reassembly-maximum-drops"] = ipv6.ReassemblyMaximumDrops
    leafs["generated-packets"] = ipv6.GeneratedPackets
    leafs["forwarded-packets"] = ipv6.ForwardedPackets
    leafs["source-routed-packets"] = ipv6.SourceRoutedPackets
    leafs["fragmented-packets"] = ipv6.FragmentedPackets
    leafs["fragment-count"] = ipv6.FragmentCount
    leafs["fragment-failures"] = ipv6.FragmentFailures
    leafs["no-route-packets"] = ipv6.NoRoutePackets
    leafs["too-big-packets"] = ipv6.TooBigPackets
    leafs["received-multicast-packets"] = ipv6.ReceivedMulticastPackets
    leafs["sent-multicast-packets"] = ipv6.SentMulticastPackets
    leafs["miscellaneous-drops"] = ipv6.MiscellaneousDrops
    leafs["lisp-v4-encap-packets"] = ipv6.LispV4EncapPackets
    leafs["lisp-v4-decap-packets"] = ipv6.LispV4DecapPackets
    leafs["lisp-v6-encap-packets"] = ipv6.LispV6EncapPackets
    leafs["lisp-v6-decap-packets"] = ipv6.LispV6DecapPackets
    leafs["lisp-encap-errors"] = ipv6.LispEncapErrors
    leafs["lisp-decap-errors"] = ipv6.LispDecapErrors
    return leafs
}

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetParentYangName() string { return "traffic" }

// Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp
// ICMP Statistics
type Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICMP Received. The type is interface{} with range: 0..4294967295.
    TotalMessages interface{}

    // ICMP Too Short Errors. The type is interface{} with range: 0..4294967295.
    TooShortErrorMessages interface{}

    // ICMP Checksum Errors. The type is interface{} with range: 0..4294967295.
    ChecksumErrorMessages interface{}

    // ICMP Unknown Error. The type is interface{} with range: 0..4294967295.
    UnknownErrorTypeMessages interface{}

    // ICMP Transmitted. The type is interface{} with range: 0..4294967295.
    OutputMessages interface{}

    // ICMP Sent Packets Ratelimited. The type is interface{} with range:
    // 0..4294967295.
    SentRateLimitedPackets interface{}

    // ICMP Route Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnreachableRoutingMessages interface{}

    // ICMP Admin Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnreachableAdminMessages interface{}

    // ICMP Host Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnreachableNeighborMessages interface{}

    // ICMP Addr Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnreachableAddressMessages interface{}

    // ICMP Port Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnreachablePortMessages interface{}

    // ICMP Route Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedUnreachableRoutingMessages interface{}

    // ICMP Admin Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedUnreachableAdminMessages interface{}

    // ICMP Host Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedUnreachableNeighborMessages interface{}

    // ICMP Addr Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedUnreachableAddressMessages interface{}

    // ICMP Port Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedUnreachablePortMessages interface{}

    // ICMP Hop Count Expired Sent. The type is interface{} with range:
    // 0..4294967295.
    SentHopCountExpiredMessages interface{}

    // ICMP Reassembly Timeouts. The type is interface{} with range:
    // 0..4294967295.
    SentReassemblyTimeouts interface{}

    // ICMP Hop Count Expired Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedHopCountExpiredMessages interface{}

    // ICMP Reassembly Timeouts. The type is interface{} with range:
    // 0..4294967295.
    ReceivedReassemblyTimeouts interface{}

    // ICMP Too Big Messages Sent. The type is interface{} with range:
    // 0..4294967295.
    SentTooBigMessages interface{}

    // ICMP Too Big Messages Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedTooBigMessages interface{}

    // ICMP Parameter Error Messages Sent. The type is interface{} with range:
    // 0..4294967295.
    SentParameterErrorMessages interface{}

    // ICMP Parameter Next Header Messages Sent. The type is interface{} with
    // range: 0..4294967295.
    SentParameterHeaderMessages interface{}

    // ICMP Parameter Option Messages Sent. The type is interface{} with range:
    // 0..4294967295.
    SentParameterOptionMessages interface{}

    // ICMP Parameter Error Messages Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedParameterErrorMessages interface{}

    // ICMP Parameter Next Header Messages Received. The type is interface{} with
    // range: 0..4294967295.
    ReceivedParameterHeaderMessages interface{}

    // ICMP Parameter Option Problem Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedParameterOptionMessages interface{}

    // ICMP Echo Request Sent. The type is interface{} with range: 0..4294967295.
    SentEchoRequestMessages interface{}

    // ICMP Echo Reply Sent. The type is interface{} with range: 0..4294967295.
    SentEchoReplyMessages interface{}

    // ICMP Echo Request Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedEchoRequestMessages interface{}

    // ICMP Echo Reply Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedEchoReplyMessages interface{}

    // ICMP Unknown Timeout Messages Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnknownTimeoutMessages interface{}

    // ICMP Unknown Timeout Messages Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedUnknownTimeoutMessages interface{}

    // ICMP Parameter Unknown Type Messages Sent. The type is interface{} with
    // range: 0..4294967295.
    SentParameterUnknownTypeMessages interface{}

    // ICMP Parameter Unknown Type Messages Received. The type is interface{} with
    // range: 0..4294967295.
    ReceivedParameterUnknownTypeMessages interface{}

    // ICMP Unreachable Unknown Messages Sent. The type is interface{} with range:
    // 0..4294967295.
    SentUnreachableUnknownTypeMessages interface{}

    // ICMP Unreachable Unknown Messages Received. The type is interface{} with
    // range: 0..4294967295.
    ReceivedUnreachableUnknownTypeMessages interface{}
}

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetFilter() yfilter.YFilter { return icmp.YFilter }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) SetFilter(yf yfilter.YFilter) { icmp.YFilter = yf }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetGoName(yname string) string {
    if yname == "total-messages" { return "TotalMessages" }
    if yname == "too-short-error-messages" { return "TooShortErrorMessages" }
    if yname == "checksum-error-messages" { return "ChecksumErrorMessages" }
    if yname == "unknown-error-type-messages" { return "UnknownErrorTypeMessages" }
    if yname == "output-messages" { return "OutputMessages" }
    if yname == "sent-rate-limited-packets" { return "SentRateLimitedPackets" }
    if yname == "sent-unreachable-routing-messages" { return "SentUnreachableRoutingMessages" }
    if yname == "sent-unreachable-admin-messages" { return "SentUnreachableAdminMessages" }
    if yname == "sent-unreachable-neighbor-messages" { return "SentUnreachableNeighborMessages" }
    if yname == "sent-unreachable-address-messages" { return "SentUnreachableAddressMessages" }
    if yname == "sent-unreachable-port-messages" { return "SentUnreachablePortMessages" }
    if yname == "received-unreachable-routing-messages" { return "ReceivedUnreachableRoutingMessages" }
    if yname == "received-unreachable-admin-messages" { return "ReceivedUnreachableAdminMessages" }
    if yname == "received-unreachable-neighbor-messages" { return "ReceivedUnreachableNeighborMessages" }
    if yname == "received-unreachable-address-messages" { return "ReceivedUnreachableAddressMessages" }
    if yname == "received-unreachable-port-messages" { return "ReceivedUnreachablePortMessages" }
    if yname == "sent-hop-count-expired-messages" { return "SentHopCountExpiredMessages" }
    if yname == "sent-reassembly-timeouts" { return "SentReassemblyTimeouts" }
    if yname == "received-hop-count-expired-messages" { return "ReceivedHopCountExpiredMessages" }
    if yname == "received-reassembly-timeouts" { return "ReceivedReassemblyTimeouts" }
    if yname == "sent-too-big-messages" { return "SentTooBigMessages" }
    if yname == "received-too-big-messages" { return "ReceivedTooBigMessages" }
    if yname == "sent-parameter-error-messages" { return "SentParameterErrorMessages" }
    if yname == "sent-parameter-header-messages" { return "SentParameterHeaderMessages" }
    if yname == "sent-parameter-option-messages" { return "SentParameterOptionMessages" }
    if yname == "received-parameter-error-messages" { return "ReceivedParameterErrorMessages" }
    if yname == "received-parameter-header-messages" { return "ReceivedParameterHeaderMessages" }
    if yname == "received-parameter-option-messages" { return "ReceivedParameterOptionMessages" }
    if yname == "sent-echo-request-messages" { return "SentEchoRequestMessages" }
    if yname == "sent-echo-reply-messages" { return "SentEchoReplyMessages" }
    if yname == "received-echo-request-messages" { return "ReceivedEchoRequestMessages" }
    if yname == "received-echo-reply-messages" { return "ReceivedEchoReplyMessages" }
    if yname == "sent-unknown-timeout-messages" { return "SentUnknownTimeoutMessages" }
    if yname == "received-unknown-timeout-messages" { return "ReceivedUnknownTimeoutMessages" }
    if yname == "sent-parameter-unknown-type-messages" { return "SentParameterUnknownTypeMessages" }
    if yname == "received-parameter-unknown-type-messages" { return "ReceivedParameterUnknownTypeMessages" }
    if yname == "sent-unreachable-unknown-type-messages" { return "SentUnreachableUnknownTypeMessages" }
    if yname == "received-unreachable-unknown-type-messages" { return "ReceivedUnreachableUnknownTypeMessages" }
    return ""
}

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetSegmentPath() string {
    return "icmp"
}

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-messages"] = icmp.TotalMessages
    leafs["too-short-error-messages"] = icmp.TooShortErrorMessages
    leafs["checksum-error-messages"] = icmp.ChecksumErrorMessages
    leafs["unknown-error-type-messages"] = icmp.UnknownErrorTypeMessages
    leafs["output-messages"] = icmp.OutputMessages
    leafs["sent-rate-limited-packets"] = icmp.SentRateLimitedPackets
    leafs["sent-unreachable-routing-messages"] = icmp.SentUnreachableRoutingMessages
    leafs["sent-unreachable-admin-messages"] = icmp.SentUnreachableAdminMessages
    leafs["sent-unreachable-neighbor-messages"] = icmp.SentUnreachableNeighborMessages
    leafs["sent-unreachable-address-messages"] = icmp.SentUnreachableAddressMessages
    leafs["sent-unreachable-port-messages"] = icmp.SentUnreachablePortMessages
    leafs["received-unreachable-routing-messages"] = icmp.ReceivedUnreachableRoutingMessages
    leafs["received-unreachable-admin-messages"] = icmp.ReceivedUnreachableAdminMessages
    leafs["received-unreachable-neighbor-messages"] = icmp.ReceivedUnreachableNeighborMessages
    leafs["received-unreachable-address-messages"] = icmp.ReceivedUnreachableAddressMessages
    leafs["received-unreachable-port-messages"] = icmp.ReceivedUnreachablePortMessages
    leafs["sent-hop-count-expired-messages"] = icmp.SentHopCountExpiredMessages
    leafs["sent-reassembly-timeouts"] = icmp.SentReassemblyTimeouts
    leafs["received-hop-count-expired-messages"] = icmp.ReceivedHopCountExpiredMessages
    leafs["received-reassembly-timeouts"] = icmp.ReceivedReassemblyTimeouts
    leafs["sent-too-big-messages"] = icmp.SentTooBigMessages
    leafs["received-too-big-messages"] = icmp.ReceivedTooBigMessages
    leafs["sent-parameter-error-messages"] = icmp.SentParameterErrorMessages
    leafs["sent-parameter-header-messages"] = icmp.SentParameterHeaderMessages
    leafs["sent-parameter-option-messages"] = icmp.SentParameterOptionMessages
    leafs["received-parameter-error-messages"] = icmp.ReceivedParameterErrorMessages
    leafs["received-parameter-header-messages"] = icmp.ReceivedParameterHeaderMessages
    leafs["received-parameter-option-messages"] = icmp.ReceivedParameterOptionMessages
    leafs["sent-echo-request-messages"] = icmp.SentEchoRequestMessages
    leafs["sent-echo-reply-messages"] = icmp.SentEchoReplyMessages
    leafs["received-echo-request-messages"] = icmp.ReceivedEchoRequestMessages
    leafs["received-echo-reply-messages"] = icmp.ReceivedEchoReplyMessages
    leafs["sent-unknown-timeout-messages"] = icmp.SentUnknownTimeoutMessages
    leafs["received-unknown-timeout-messages"] = icmp.ReceivedUnknownTimeoutMessages
    leafs["sent-parameter-unknown-type-messages"] = icmp.SentParameterUnknownTypeMessages
    leafs["received-parameter-unknown-type-messages"] = icmp.ReceivedParameterUnknownTypeMessages
    leafs["sent-unreachable-unknown-type-messages"] = icmp.SentUnreachableUnknownTypeMessages
    leafs["received-unreachable-unknown-type-messages"] = icmp.ReceivedUnreachableUnknownTypeMessages
    return leafs
}

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetBundleName() string { return "cisco_ios_xr" }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetYangName() string { return "icmp" }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) SetParent(parent types.Entity) { icmp.parent = parent }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetParent() types.Entity { return icmp.parent }

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetParentYangName() string { return "traffic" }

// Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery
// IPv6 Node Discovery Statistics
type Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICMP Router Solicitations Sent. The type is interface{} with range:
    // 0..4294967295.
    SentRouterSolicitationMessages interface{}

    // ICMP Router Advertisements Sent. The type is interface{} with range:
    // 0..4294967295.
    SentRouterAdvertisementMessages interface{}

    // ICMP Neighbor Solicitations Sent. The type is interface{} with range:
    // 0..4294967295.
    SentNeighborSolicitationMessages interface{}

    // ICMP Neighbor Advertisements Sent. The type is interface{} with range:
    // 0..4294967295.
    SentNeighborAdvertisementMessages interface{}

    // ICMP Redirect Sent. The type is interface{} with range: 0..4294967295.
    SentRedirectMessages interface{}

    // ICMP Router Solicitations Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedRouterSolicitationMessages interface{}

    // ICMP Router Advertisements Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedRouterAdvertisementMessages interface{}

    // ICMP Neighbor Solicitations Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedNeighborSolicitationMessages interface{}

    // ICMP Neighbor Advertisements Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedNeighborAdvertisementMessages interface{}

    // ICMP Redirect Received. The type is interface{} with range: 0..4294967295.
    ReceivedRedirectMessages interface{}
}

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetFilter() yfilter.YFilter { return ipv6NodeDiscovery.YFilter }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) SetFilter(yf yfilter.YFilter) { ipv6NodeDiscovery.YFilter = yf }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetGoName(yname string) string {
    if yname == "sent-router-solicitation-messages" { return "SentRouterSolicitationMessages" }
    if yname == "sent-router-advertisement-messages" { return "SentRouterAdvertisementMessages" }
    if yname == "sent-neighbor-solicitation-messages" { return "SentNeighborSolicitationMessages" }
    if yname == "sent-neighbor-advertisement-messages" { return "SentNeighborAdvertisementMessages" }
    if yname == "sent-redirect-messages" { return "SentRedirectMessages" }
    if yname == "received-router-solicitation-messages" { return "ReceivedRouterSolicitationMessages" }
    if yname == "received-router-advertisement-messages" { return "ReceivedRouterAdvertisementMessages" }
    if yname == "received-neighbor-solicitation-messages" { return "ReceivedNeighborSolicitationMessages" }
    if yname == "received-neighbor-advertisement-messages" { return "ReceivedNeighborAdvertisementMessages" }
    if yname == "received-redirect-messages" { return "ReceivedRedirectMessages" }
    return ""
}

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetSegmentPath() string {
    return "ipv6-node-discovery"
}

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent-router-solicitation-messages"] = ipv6NodeDiscovery.SentRouterSolicitationMessages
    leafs["sent-router-advertisement-messages"] = ipv6NodeDiscovery.SentRouterAdvertisementMessages
    leafs["sent-neighbor-solicitation-messages"] = ipv6NodeDiscovery.SentNeighborSolicitationMessages
    leafs["sent-neighbor-advertisement-messages"] = ipv6NodeDiscovery.SentNeighborAdvertisementMessages
    leafs["sent-redirect-messages"] = ipv6NodeDiscovery.SentRedirectMessages
    leafs["received-router-solicitation-messages"] = ipv6NodeDiscovery.ReceivedRouterSolicitationMessages
    leafs["received-router-advertisement-messages"] = ipv6NodeDiscovery.ReceivedRouterAdvertisementMessages
    leafs["received-neighbor-solicitation-messages"] = ipv6NodeDiscovery.ReceivedNeighborSolicitationMessages
    leafs["received-neighbor-advertisement-messages"] = ipv6NodeDiscovery.ReceivedNeighborAdvertisementMessages
    leafs["received-redirect-messages"] = ipv6NodeDiscovery.ReceivedRedirectMessages
    return leafs
}

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetYangName() string { return "ipv6-node-discovery" }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) SetParent(parent types.Entity) { ipv6NodeDiscovery.parent = parent }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetParent() types.Entity { return ipv6NodeDiscovery.parent }

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetParentYangName() string { return "traffic" }

