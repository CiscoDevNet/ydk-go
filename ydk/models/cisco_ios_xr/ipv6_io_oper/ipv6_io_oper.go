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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific IPv6 IO operational data.
    Nodes Ipv6Io_Nodes
}

func (ipv6Io *Ipv6Io) GetEntityData() *types.CommonEntityData {
    ipv6Io.EntityData.YFilter = ipv6Io.YFilter
    ipv6Io.EntityData.YangName = "ipv6-io"
    ipv6Io.EntityData.BundleName = "cisco_ios_xr"
    ipv6Io.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-io-oper"
    ipv6Io.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-io-oper:ipv6-io"
    ipv6Io.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Io.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Io.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Io.EntityData.Children = make(map[string]types.YChild)
    ipv6Io.EntityData.Children["nodes"] = types.YChild{"Nodes", &ipv6Io.Nodes}
    ipv6Io.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Io.EntityData)
}

// Ipv6Io_Nodes
// Node-specific IPv6 IO operational data
type Ipv6Io_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 network operational data for a particular node. The type is slice of
    // Ipv6Io_Nodes_Node.
    Node []Ipv6Io_Nodes_Node
}

func (nodes *Ipv6Io_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ipv6-io"
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

// Ipv6Io_Nodes_Node
// IPv6 network operational data for a particular
// node
type Ipv6Io_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Statistical IPv6 network operational data for a node.
    Statistics Ipv6Io_Nodes_Node_Statistics
}

func (node *Ipv6Io_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Ipv6Io_Nodes_Node_Statistics
// Statistical IPv6 network operational data for
// a node
type Ipv6Io_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic statistics for a node.
    Traffic Ipv6Io_Nodes_Node_Statistics_Traffic
}

func (statistics *Ipv6Io_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["traffic"] = types.YChild{"Traffic", &statistics.Traffic}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv6Io_Nodes_Node_Statistics_Traffic
// Traffic statistics for a node
type Ipv6Io_Nodes_Node_Statistics_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Statistics.
    Ipv6 Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6

    // ICMP Statistics.
    Icmp Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp

    // IPv6 Node Discovery Statistics.
    Ipv6NodeDiscovery Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery
}

func (traffic *Ipv6Io_Nodes_Node_Statistics_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "statistics"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = make(map[string]types.YChild)
    traffic.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &traffic.Ipv6}
    traffic.EntityData.Children["icmp"] = types.YChild{"Icmp", &traffic.Icmp}
    traffic.EntityData.Children["ipv6-node-discovery"] = types.YChild{"Ipv6NodeDiscovery", &traffic.Ipv6NodeDiscovery}
    traffic.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(traffic.EntityData)
}

// Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6
// IPv6 Statistics
type Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6 struct {
    EntityData types.CommonEntityData
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

func (ipv6 *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "traffic"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["total-packets"] = types.YLeaf{"TotalPackets", ipv6.TotalPackets}
    ipv6.EntityData.Leafs["local-destination-packets"] = types.YLeaf{"LocalDestinationPackets", ipv6.LocalDestinationPackets}
    ipv6.EntityData.Leafs["format-errors"] = types.YLeaf{"FormatErrors", ipv6.FormatErrors}
    ipv6.EntityData.Leafs["truncated-packets"] = types.YLeaf{"TruncatedPackets", ipv6.TruncatedPackets}
    ipv6.EntityData.Leafs["hop-count-exceeded-packets"] = types.YLeaf{"HopCountExceededPackets", ipv6.HopCountExceededPackets}
    ipv6.EntityData.Leafs["bad-source-address-packets"] = types.YLeaf{"BadSourceAddressPackets", ipv6.BadSourceAddressPackets}
    ipv6.EntityData.Leafs["bad-header-packets"] = types.YLeaf{"BadHeaderPackets", ipv6.BadHeaderPackets}
    ipv6.EntityData.Leafs["unknown-option-type-packets"] = types.YLeaf{"UnknownOptionTypePackets", ipv6.UnknownOptionTypePackets}
    ipv6.EntityData.Leafs["unknown-protocol-packets"] = types.YLeaf{"UnknownProtocolPackets", ipv6.UnknownProtocolPackets}
    ipv6.EntityData.Leafs["fragments"] = types.YLeaf{"Fragments", ipv6.Fragments}
    ipv6.EntityData.Leafs["reassembled-packets"] = types.YLeaf{"ReassembledPackets", ipv6.ReassembledPackets}
    ipv6.EntityData.Leafs["reassembly-timeouts"] = types.YLeaf{"ReassemblyTimeouts", ipv6.ReassemblyTimeouts}
    ipv6.EntityData.Leafs["reassembly-failures"] = types.YLeaf{"ReassemblyFailures", ipv6.ReassemblyFailures}
    ipv6.EntityData.Leafs["reassembly-maximum-drops"] = types.YLeaf{"ReassemblyMaximumDrops", ipv6.ReassemblyMaximumDrops}
    ipv6.EntityData.Leafs["generated-packets"] = types.YLeaf{"GeneratedPackets", ipv6.GeneratedPackets}
    ipv6.EntityData.Leafs["forwarded-packets"] = types.YLeaf{"ForwardedPackets", ipv6.ForwardedPackets}
    ipv6.EntityData.Leafs["source-routed-packets"] = types.YLeaf{"SourceRoutedPackets", ipv6.SourceRoutedPackets}
    ipv6.EntityData.Leafs["fragmented-packets"] = types.YLeaf{"FragmentedPackets", ipv6.FragmentedPackets}
    ipv6.EntityData.Leafs["fragment-count"] = types.YLeaf{"FragmentCount", ipv6.FragmentCount}
    ipv6.EntityData.Leafs["fragment-failures"] = types.YLeaf{"FragmentFailures", ipv6.FragmentFailures}
    ipv6.EntityData.Leafs["no-route-packets"] = types.YLeaf{"NoRoutePackets", ipv6.NoRoutePackets}
    ipv6.EntityData.Leafs["too-big-packets"] = types.YLeaf{"TooBigPackets", ipv6.TooBigPackets}
    ipv6.EntityData.Leafs["received-multicast-packets"] = types.YLeaf{"ReceivedMulticastPackets", ipv6.ReceivedMulticastPackets}
    ipv6.EntityData.Leafs["sent-multicast-packets"] = types.YLeaf{"SentMulticastPackets", ipv6.SentMulticastPackets}
    ipv6.EntityData.Leafs["miscellaneous-drops"] = types.YLeaf{"MiscellaneousDrops", ipv6.MiscellaneousDrops}
    ipv6.EntityData.Leafs["lisp-v4-encap-packets"] = types.YLeaf{"LispV4EncapPackets", ipv6.LispV4EncapPackets}
    ipv6.EntityData.Leafs["lisp-v4-decap-packets"] = types.YLeaf{"LispV4DecapPackets", ipv6.LispV4DecapPackets}
    ipv6.EntityData.Leafs["lisp-v6-encap-packets"] = types.YLeaf{"LispV6EncapPackets", ipv6.LispV6EncapPackets}
    ipv6.EntityData.Leafs["lisp-v6-decap-packets"] = types.YLeaf{"LispV6DecapPackets", ipv6.LispV6DecapPackets}
    ipv6.EntityData.Leafs["lisp-encap-errors"] = types.YLeaf{"LispEncapErrors", ipv6.LispEncapErrors}
    ipv6.EntityData.Leafs["lisp-decap-errors"] = types.YLeaf{"LispDecapErrors", ipv6.LispDecapErrors}
    return &(ipv6.EntityData)
}

// Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp
// ICMP Statistics
type Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp struct {
    EntityData types.CommonEntityData
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

func (icmp *Ipv6Io_Nodes_Node_Statistics_Traffic_Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xr"
    icmp.EntityData.ParentYangName = "traffic"
    icmp.EntityData.SegmentPath = "icmp"
    icmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmp.EntityData.Children = make(map[string]types.YChild)
    icmp.EntityData.Leafs = make(map[string]types.YLeaf)
    icmp.EntityData.Leafs["total-messages"] = types.YLeaf{"TotalMessages", icmp.TotalMessages}
    icmp.EntityData.Leafs["too-short-error-messages"] = types.YLeaf{"TooShortErrorMessages", icmp.TooShortErrorMessages}
    icmp.EntityData.Leafs["checksum-error-messages"] = types.YLeaf{"ChecksumErrorMessages", icmp.ChecksumErrorMessages}
    icmp.EntityData.Leafs["unknown-error-type-messages"] = types.YLeaf{"UnknownErrorTypeMessages", icmp.UnknownErrorTypeMessages}
    icmp.EntityData.Leafs["output-messages"] = types.YLeaf{"OutputMessages", icmp.OutputMessages}
    icmp.EntityData.Leafs["sent-rate-limited-packets"] = types.YLeaf{"SentRateLimitedPackets", icmp.SentRateLimitedPackets}
    icmp.EntityData.Leafs["sent-unreachable-routing-messages"] = types.YLeaf{"SentUnreachableRoutingMessages", icmp.SentUnreachableRoutingMessages}
    icmp.EntityData.Leafs["sent-unreachable-admin-messages"] = types.YLeaf{"SentUnreachableAdminMessages", icmp.SentUnreachableAdminMessages}
    icmp.EntityData.Leafs["sent-unreachable-neighbor-messages"] = types.YLeaf{"SentUnreachableNeighborMessages", icmp.SentUnreachableNeighborMessages}
    icmp.EntityData.Leafs["sent-unreachable-address-messages"] = types.YLeaf{"SentUnreachableAddressMessages", icmp.SentUnreachableAddressMessages}
    icmp.EntityData.Leafs["sent-unreachable-port-messages"] = types.YLeaf{"SentUnreachablePortMessages", icmp.SentUnreachablePortMessages}
    icmp.EntityData.Leafs["received-unreachable-routing-messages"] = types.YLeaf{"ReceivedUnreachableRoutingMessages", icmp.ReceivedUnreachableRoutingMessages}
    icmp.EntityData.Leafs["received-unreachable-admin-messages"] = types.YLeaf{"ReceivedUnreachableAdminMessages", icmp.ReceivedUnreachableAdminMessages}
    icmp.EntityData.Leafs["received-unreachable-neighbor-messages"] = types.YLeaf{"ReceivedUnreachableNeighborMessages", icmp.ReceivedUnreachableNeighborMessages}
    icmp.EntityData.Leafs["received-unreachable-address-messages"] = types.YLeaf{"ReceivedUnreachableAddressMessages", icmp.ReceivedUnreachableAddressMessages}
    icmp.EntityData.Leafs["received-unreachable-port-messages"] = types.YLeaf{"ReceivedUnreachablePortMessages", icmp.ReceivedUnreachablePortMessages}
    icmp.EntityData.Leafs["sent-hop-count-expired-messages"] = types.YLeaf{"SentHopCountExpiredMessages", icmp.SentHopCountExpiredMessages}
    icmp.EntityData.Leafs["sent-reassembly-timeouts"] = types.YLeaf{"SentReassemblyTimeouts", icmp.SentReassemblyTimeouts}
    icmp.EntityData.Leafs["received-hop-count-expired-messages"] = types.YLeaf{"ReceivedHopCountExpiredMessages", icmp.ReceivedHopCountExpiredMessages}
    icmp.EntityData.Leafs["received-reassembly-timeouts"] = types.YLeaf{"ReceivedReassemblyTimeouts", icmp.ReceivedReassemblyTimeouts}
    icmp.EntityData.Leafs["sent-too-big-messages"] = types.YLeaf{"SentTooBigMessages", icmp.SentTooBigMessages}
    icmp.EntityData.Leafs["received-too-big-messages"] = types.YLeaf{"ReceivedTooBigMessages", icmp.ReceivedTooBigMessages}
    icmp.EntityData.Leafs["sent-parameter-error-messages"] = types.YLeaf{"SentParameterErrorMessages", icmp.SentParameterErrorMessages}
    icmp.EntityData.Leafs["sent-parameter-header-messages"] = types.YLeaf{"SentParameterHeaderMessages", icmp.SentParameterHeaderMessages}
    icmp.EntityData.Leafs["sent-parameter-option-messages"] = types.YLeaf{"SentParameterOptionMessages", icmp.SentParameterOptionMessages}
    icmp.EntityData.Leafs["received-parameter-error-messages"] = types.YLeaf{"ReceivedParameterErrorMessages", icmp.ReceivedParameterErrorMessages}
    icmp.EntityData.Leafs["received-parameter-header-messages"] = types.YLeaf{"ReceivedParameterHeaderMessages", icmp.ReceivedParameterHeaderMessages}
    icmp.EntityData.Leafs["received-parameter-option-messages"] = types.YLeaf{"ReceivedParameterOptionMessages", icmp.ReceivedParameterOptionMessages}
    icmp.EntityData.Leafs["sent-echo-request-messages"] = types.YLeaf{"SentEchoRequestMessages", icmp.SentEchoRequestMessages}
    icmp.EntityData.Leafs["sent-echo-reply-messages"] = types.YLeaf{"SentEchoReplyMessages", icmp.SentEchoReplyMessages}
    icmp.EntityData.Leafs["received-echo-request-messages"] = types.YLeaf{"ReceivedEchoRequestMessages", icmp.ReceivedEchoRequestMessages}
    icmp.EntityData.Leafs["received-echo-reply-messages"] = types.YLeaf{"ReceivedEchoReplyMessages", icmp.ReceivedEchoReplyMessages}
    icmp.EntityData.Leafs["sent-unknown-timeout-messages"] = types.YLeaf{"SentUnknownTimeoutMessages", icmp.SentUnknownTimeoutMessages}
    icmp.EntityData.Leafs["received-unknown-timeout-messages"] = types.YLeaf{"ReceivedUnknownTimeoutMessages", icmp.ReceivedUnknownTimeoutMessages}
    icmp.EntityData.Leafs["sent-parameter-unknown-type-messages"] = types.YLeaf{"SentParameterUnknownTypeMessages", icmp.SentParameterUnknownTypeMessages}
    icmp.EntityData.Leafs["received-parameter-unknown-type-messages"] = types.YLeaf{"ReceivedParameterUnknownTypeMessages", icmp.ReceivedParameterUnknownTypeMessages}
    icmp.EntityData.Leafs["sent-unreachable-unknown-type-messages"] = types.YLeaf{"SentUnreachableUnknownTypeMessages", icmp.SentUnreachableUnknownTypeMessages}
    icmp.EntityData.Leafs["received-unreachable-unknown-type-messages"] = types.YLeaf{"ReceivedUnreachableUnknownTypeMessages", icmp.ReceivedUnreachableUnknownTypeMessages}
    return &(icmp.EntityData)
}

// Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery
// IPv6 Node Discovery Statistics
type Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery struct {
    EntityData types.CommonEntityData
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

func (ipv6NodeDiscovery *Ipv6Io_Nodes_Node_Statistics_Traffic_Ipv6NodeDiscovery) GetEntityData() *types.CommonEntityData {
    ipv6NodeDiscovery.EntityData.YFilter = ipv6NodeDiscovery.YFilter
    ipv6NodeDiscovery.EntityData.YangName = "ipv6-node-discovery"
    ipv6NodeDiscovery.EntityData.BundleName = "cisco_ios_xr"
    ipv6NodeDiscovery.EntityData.ParentYangName = "traffic"
    ipv6NodeDiscovery.EntityData.SegmentPath = "ipv6-node-discovery"
    ipv6NodeDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6NodeDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6NodeDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6NodeDiscovery.EntityData.Children = make(map[string]types.YChild)
    ipv6NodeDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6NodeDiscovery.EntityData.Leafs["sent-router-solicitation-messages"] = types.YLeaf{"SentRouterSolicitationMessages", ipv6NodeDiscovery.SentRouterSolicitationMessages}
    ipv6NodeDiscovery.EntityData.Leafs["sent-router-advertisement-messages"] = types.YLeaf{"SentRouterAdvertisementMessages", ipv6NodeDiscovery.SentRouterAdvertisementMessages}
    ipv6NodeDiscovery.EntityData.Leafs["sent-neighbor-solicitation-messages"] = types.YLeaf{"SentNeighborSolicitationMessages", ipv6NodeDiscovery.SentNeighborSolicitationMessages}
    ipv6NodeDiscovery.EntityData.Leafs["sent-neighbor-advertisement-messages"] = types.YLeaf{"SentNeighborAdvertisementMessages", ipv6NodeDiscovery.SentNeighborAdvertisementMessages}
    ipv6NodeDiscovery.EntityData.Leafs["sent-redirect-messages"] = types.YLeaf{"SentRedirectMessages", ipv6NodeDiscovery.SentRedirectMessages}
    ipv6NodeDiscovery.EntityData.Leafs["received-router-solicitation-messages"] = types.YLeaf{"ReceivedRouterSolicitationMessages", ipv6NodeDiscovery.ReceivedRouterSolicitationMessages}
    ipv6NodeDiscovery.EntityData.Leafs["received-router-advertisement-messages"] = types.YLeaf{"ReceivedRouterAdvertisementMessages", ipv6NodeDiscovery.ReceivedRouterAdvertisementMessages}
    ipv6NodeDiscovery.EntityData.Leafs["received-neighbor-solicitation-messages"] = types.YLeaf{"ReceivedNeighborSolicitationMessages", ipv6NodeDiscovery.ReceivedNeighborSolicitationMessages}
    ipv6NodeDiscovery.EntityData.Leafs["received-neighbor-advertisement-messages"] = types.YLeaf{"ReceivedNeighborAdvertisementMessages", ipv6NodeDiscovery.ReceivedNeighborAdvertisementMessages}
    ipv6NodeDiscovery.EntityData.Leafs["received-redirect-messages"] = types.YLeaf{"ReceivedRedirectMessages", ipv6NodeDiscovery.ReceivedRedirectMessages}
    return &(ipv6NodeDiscovery.EntityData)
}

