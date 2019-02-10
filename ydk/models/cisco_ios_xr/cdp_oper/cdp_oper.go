// This module contains a collection of YANG definitions
// for Cisco IOS-XR cdp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cdp: CDP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package cdp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cdp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cdp-oper cdp}", reflect.TypeOf(Cdp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cdp-oper:cdp", reflect.TypeOf(Cdp{}))
}

// CdpDuplex represents Cdp duplex
type CdpDuplex string

const (
    // cdp dplx none
    CdpDuplex_cdp_dplx_none CdpDuplex = "cdp-dplx-none"

    // cdp dplx half
    CdpDuplex_cdp_dplx_half CdpDuplex = "cdp-dplx-half"

    // cdp dplx full
    CdpDuplex_cdp_dplx_full CdpDuplex = "cdp-dplx-full"
)

// CdpL3AddrProtocol represents Cdp l3 addr protocol
type CdpL3AddrProtocol string

const (
    // IPv4
    CdpL3AddrProtocol_ipv4 CdpL3AddrProtocol = "ipv4"

    // IPv6
    CdpL3AddrProtocol_ipv6 CdpL3AddrProtocol = "ipv6"
)

// Cdp
// CDP operational data
type Cdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node CDP operational data.
    Nodes Cdp_Nodes
}

func (cdp *Cdp) GetEntityData() *types.CommonEntityData {
    cdp.EntityData.YFilter = cdp.YFilter
    cdp.EntityData.YangName = "cdp"
    cdp.EntityData.BundleName = "cisco_ios_xr"
    cdp.EntityData.ParentYangName = "Cisco-IOS-XR-cdp-oper"
    cdp.EntityData.SegmentPath = "Cisco-IOS-XR-cdp-oper:cdp"
    cdp.EntityData.AbsolutePath = cdp.EntityData.SegmentPath
    cdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdp.EntityData.Children = types.NewOrderedMap()
    cdp.EntityData.Children.Append("nodes", types.YChild{"Nodes", &cdp.Nodes})
    cdp.EntityData.Leafs = types.NewOrderedMap()

    cdp.EntityData.YListKeys = []string {}

    return &(cdp.EntityData)
}

// Cdp_Nodes
// Per node CDP operational data
type Cdp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The CDP operational data for a particular node. The type is slice of
    // Cdp_Nodes_Node.
    Node []*Cdp_Nodes_Node
}

func (nodes *Cdp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cdp"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/" + nodes.EntityData.SegmentPath
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

// Cdp_Nodes_Node
// The CDP operational data for a particular node
type Cdp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // The CDP neighbor tables on this node.
    Neighbors Cdp_Nodes_Node_Neighbors

    // The CDP traffic statistics for this node.
    Statistics Cdp_Nodes_Node_Statistics

    // The table of interfaces on which CDP is running on this node.
    Interfaces Cdp_Nodes_Node_Interfaces
}

func (node *Cdp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &node.Neighbors})
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Cdp_Nodes_Node_Neighbors
// The CDP neighbor tables on this node
type Cdp_Nodes_Node_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The detailed CDP neighbor table.
    Details Cdp_Nodes_Node_Neighbors_Details

    // The detailed CDP neighbor table.
    Devices Cdp_Nodes_Node_Neighbors_Devices

    // The CDP neighbor summary table.
    Summaries Cdp_Nodes_Node_Neighbors_Summaries
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "node"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("details", types.YChild{"Details", &neighbors.Details})
    neighbors.EntityData.Children.Append("devices", types.YChild{"Devices", &neighbors.Devices})
    neighbors.EntityData.Children.Append("summaries", types.YChild{"Summaries", &neighbors.Summaries})
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details
// The detailed CDP neighbor table
type Cdp_Nodes_Node_Neighbors_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail.
    Detail []*Cdp_Nodes_Node_Neighbors_Details_Detail
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "neighbors"
    details.EntityData.SegmentPath = "details"
    details.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/" + details.EntityData.SegmentPath
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = types.NewOrderedMap()
    details.EntityData.Children.Append("detail", types.YChild{"Detail", nil})
    for i := range details.Detail {
        types.SetYListKey(details.Detail[i], i)
        details.EntityData.Children.Append(types.GetSegmentPath(details.Detail[i]), types.YChild{"Detail", details.Detail[i]})
    }
    details.EntityData.Leafs = types.NewOrderedMap()

    details.EntityData.YListKeys = []string {}

    return &(details.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail
// Detailed information about a CDP neighbor
// entry
type Cdp_Nodes_Node_Neighbors_Details_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor.
    CdpNeighbor []*Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "details"
    detail.EntityData.SegmentPath = "detail" + types.AddNoKeyToken(detail)
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("cdp-neighbor", types.YChild{"CdpNeighbor", nil})
    for i := range detail.CdpNeighbor {
        types.SetYListKey(detail.CdpNeighbor[i], i)
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.CdpNeighbor[i]), types.YChild{"CdpNeighbor", detail.CdpNeighbor[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", detail.InterfaceName})
    detail.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", detail.DeviceId})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ReceivingInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Outgoing port identifier. The type is string.
    PortId interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Capabilities. The type is string.
    Capabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetEntityData() *types.CommonEntityData {
    cdpNeighbor.EntityData.YFilter = cdpNeighbor.YFilter
    cdpNeighbor.EntityData.YangName = "cdp-neighbor"
    cdpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    cdpNeighbor.EntityData.ParentYangName = "detail"
    cdpNeighbor.EntityData.SegmentPath = "cdp-neighbor" + types.AddNoKeyToken(cdpNeighbor)
    cdpNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/" + cdpNeighbor.EntityData.SegmentPath
    cdpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpNeighbor.EntityData.Children = types.NewOrderedMap()
    cdpNeighbor.EntityData.Children.Append("detail", types.YChild{"Detail", &cdpNeighbor.Detail})
    cdpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    cdpNeighbor.EntityData.Leafs.Append("receiving-interface-name", types.YLeaf{"ReceivingInterfaceName", cdpNeighbor.ReceivingInterfaceName})
    cdpNeighbor.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", cdpNeighbor.DeviceId})
    cdpNeighbor.EntityData.Leafs.Append("port-id", types.YLeaf{"PortId", cdpNeighbor.PortId})
    cdpNeighbor.EntityData.Leafs.Append("header-version", types.YLeaf{"HeaderVersion", cdpNeighbor.HeaderVersion})
    cdpNeighbor.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", cdpNeighbor.HoldTime})
    cdpNeighbor.EntityData.Leafs.Append("capabilities", types.YLeaf{"Capabilities", cdpNeighbor.Capabilities})
    cdpNeighbor.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", cdpNeighbor.Platform})

    cdpNeighbor.EntityData.YListKeys = []string {}

    return &(cdpNeighbor.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version TLV. The type is string.
    Version interface{}

    // VTP domain. The type is string.
    VtpDomain interface{}

    // Native VLAN. The type is interface{} with range: 0..4294967295.
    NativeVlan interface{}

    // Duplex setting. The type is CdpDuplex.
    Duplex interface{}

    // SysName. The type is string.
    SystemName interface{}

    // List of network addresses .
    NetworkAddresses Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses

    // List of protocol hello entries.
    ProtocolHelloList Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "cdp-neighbor"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/cdp-neighbor/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("network-addresses", types.YChild{"NetworkAddresses", &detail.NetworkAddresses})
    detail.EntityData.Children.Append("protocol-hello-list", types.YChild{"ProtocolHelloList", &detail.ProtocolHelloList})
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("version", types.YLeaf{"Version", detail.Version})
    detail.EntityData.Leafs.Append("vtp-domain", types.YLeaf{"VtpDomain", detail.VtpDomain})
    detail.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", detail.NativeVlan})
    detail.EntityData.Leafs.Append("duplex", types.YLeaf{"Duplex", detail.Duplex})
    detail.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", detail.SystemName})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []*Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/cdp-neighbor/detail/" + networkAddresses.EntityData.SegmentPath
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = types.NewOrderedMap()
    networkAddresses.EntityData.Children.Append("cdp-addr-entry", types.YChild{"CdpAddrEntry", nil})
    for i := range networkAddresses.CdpAddrEntry {
        types.SetYListKey(networkAddresses.CdpAddrEntry[i], i)
        networkAddresses.EntityData.Children.Append(types.GetSegmentPath(networkAddresses.CdpAddrEntry[i]), types.YChild{"CdpAddrEntry", networkAddresses.CdpAddrEntry[i]})
    }
    networkAddresses.EntityData.Leafs = types.NewOrderedMap()

    networkAddresses.EntityData.YListKeys = []string {}

    return &(networkAddresses.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetEntityData() *types.CommonEntityData {
    cdpAddrEntry.EntityData.YFilter = cdpAddrEntry.YFilter
    cdpAddrEntry.EntityData.YangName = "cdp-addr-entry"
    cdpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpAddrEntry.EntityData.ParentYangName = "network-addresses"
    cdpAddrEntry.EntityData.SegmentPath = "cdp-addr-entry" + types.AddNoKeyToken(cdpAddrEntry)
    cdpAddrEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/cdp-neighbor/detail/network-addresses/" + cdpAddrEntry.EntityData.SegmentPath
    cdpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpAddrEntry.EntityData.Children = types.NewOrderedMap()
    cdpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &cdpAddrEntry.Address})
    cdpAddrEntry.EntityData.Leafs = types.NewOrderedMap()

    cdpAddrEntry.EntityData.YListKeys = []string {}

    return &(cdpAddrEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is CdpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "cdp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/cdp-neighbor/detail/network-addresses/cdp-addr-entry/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []*Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetEntityData() *types.CommonEntityData {
    protocolHelloList.EntityData.YFilter = protocolHelloList.YFilter
    protocolHelloList.EntityData.YangName = "protocol-hello-list"
    protocolHelloList.EntityData.BundleName = "cisco_ios_xr"
    protocolHelloList.EntityData.ParentYangName = "detail"
    protocolHelloList.EntityData.SegmentPath = "protocol-hello-list"
    protocolHelloList.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/cdp-neighbor/detail/" + protocolHelloList.EntityData.SegmentPath
    protocolHelloList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolHelloList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolHelloList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolHelloList.EntityData.Children = types.NewOrderedMap()
    protocolHelloList.EntityData.Children.Append("cdp-prot-hello-entry", types.YChild{"CdpProtHelloEntry", nil})
    for i := range protocolHelloList.CdpProtHelloEntry {
        types.SetYListKey(protocolHelloList.CdpProtHelloEntry[i], i)
        protocolHelloList.EntityData.Children.Append(types.GetSegmentPath(protocolHelloList.CdpProtHelloEntry[i]), types.YChild{"CdpProtHelloEntry", protocolHelloList.CdpProtHelloEntry[i]})
    }
    protocolHelloList.EntityData.Leafs = types.NewOrderedMap()

    protocolHelloList.EntityData.YListKeys = []string {}

    return &(protocolHelloList.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Protocol Hello msg. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetEntityData() *types.CommonEntityData {
    cdpProtHelloEntry.EntityData.YFilter = cdpProtHelloEntry.YFilter
    cdpProtHelloEntry.EntityData.YangName = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpProtHelloEntry.EntityData.ParentYangName = "protocol-hello-list"
    cdpProtHelloEntry.EntityData.SegmentPath = "cdp-prot-hello-entry" + types.AddNoKeyToken(cdpProtHelloEntry)
    cdpProtHelloEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/details/detail/cdp-neighbor/detail/protocol-hello-list/" + cdpProtHelloEntry.EntityData.SegmentPath
    cdpProtHelloEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpProtHelloEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpProtHelloEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpProtHelloEntry.EntityData.Children = types.NewOrderedMap()
    cdpProtHelloEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpProtHelloEntry.EntityData.Leafs.Append("hello-message", types.YLeaf{"HelloMessage", cdpProtHelloEntry.HelloMessage})

    cdpProtHelloEntry.EntityData.YListKeys = []string {}

    return &(cdpProtHelloEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices
// The detailed CDP neighbor table
type Cdp_Nodes_Node_Neighbors_Devices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device.
    Device []*Cdp_Nodes_Node_Neighbors_Devices_Device
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetEntityData() *types.CommonEntityData {
    devices.EntityData.YFilter = devices.YFilter
    devices.EntityData.YangName = "devices"
    devices.EntityData.BundleName = "cisco_ios_xr"
    devices.EntityData.ParentYangName = "neighbors"
    devices.EntityData.SegmentPath = "devices"
    devices.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/" + devices.EntityData.SegmentPath
    devices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devices.EntityData.Children = types.NewOrderedMap()
    devices.EntityData.Children.Append("device", types.YChild{"Device", nil})
    for i := range devices.Device {
        devices.EntityData.Children.Append(types.GetSegmentPath(devices.Device[i]), types.YChild{"Device", devices.Device[i]})
    }
    devices.EntityData.Leafs = types.NewOrderedMap()

    devices.EntityData.YListKeys = []string {}

    return &(devices.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device
// Detailed information about a CDP neighbor
// entry
type Cdp_Nodes_Node_Neighbors_Devices_Device struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The neighboring device identifier. The type is
    // string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor.
    CdpNeighbor []*Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetEntityData() *types.CommonEntityData {
    device.EntityData.YFilter = device.YFilter
    device.EntityData.YangName = "device"
    device.EntityData.BundleName = "cisco_ios_xr"
    device.EntityData.ParentYangName = "devices"
    device.EntityData.SegmentPath = "device" + types.AddKeyToken(device.DeviceId, "device-id")
    device.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/" + device.EntityData.SegmentPath
    device.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    device.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    device.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    device.EntityData.Children = types.NewOrderedMap()
    device.EntityData.Children.Append("cdp-neighbor", types.YChild{"CdpNeighbor", nil})
    for i := range device.CdpNeighbor {
        types.SetYListKey(device.CdpNeighbor[i], i)
        device.EntityData.Children.Append(types.GetSegmentPath(device.CdpNeighbor[i]), types.YChild{"CdpNeighbor", device.CdpNeighbor[i]})
    }
    device.EntityData.Leafs = types.NewOrderedMap()
    device.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", device.DeviceId})

    device.EntityData.YListKeys = []string {"DeviceId"}

    return &(device.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ReceivingInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Outgoing port identifier. The type is string.
    PortId interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Capabilities. The type is string.
    Capabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetEntityData() *types.CommonEntityData {
    cdpNeighbor.EntityData.YFilter = cdpNeighbor.YFilter
    cdpNeighbor.EntityData.YangName = "cdp-neighbor"
    cdpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    cdpNeighbor.EntityData.ParentYangName = "device"
    cdpNeighbor.EntityData.SegmentPath = "cdp-neighbor" + types.AddNoKeyToken(cdpNeighbor)
    cdpNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/" + cdpNeighbor.EntityData.SegmentPath
    cdpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpNeighbor.EntityData.Children = types.NewOrderedMap()
    cdpNeighbor.EntityData.Children.Append("detail", types.YChild{"Detail", &cdpNeighbor.Detail})
    cdpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    cdpNeighbor.EntityData.Leafs.Append("receiving-interface-name", types.YLeaf{"ReceivingInterfaceName", cdpNeighbor.ReceivingInterfaceName})
    cdpNeighbor.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", cdpNeighbor.DeviceId})
    cdpNeighbor.EntityData.Leafs.Append("port-id", types.YLeaf{"PortId", cdpNeighbor.PortId})
    cdpNeighbor.EntityData.Leafs.Append("header-version", types.YLeaf{"HeaderVersion", cdpNeighbor.HeaderVersion})
    cdpNeighbor.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", cdpNeighbor.HoldTime})
    cdpNeighbor.EntityData.Leafs.Append("capabilities", types.YLeaf{"Capabilities", cdpNeighbor.Capabilities})
    cdpNeighbor.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", cdpNeighbor.Platform})

    cdpNeighbor.EntityData.YListKeys = []string {}

    return &(cdpNeighbor.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version TLV. The type is string.
    Version interface{}

    // VTP domain. The type is string.
    VtpDomain interface{}

    // Native VLAN. The type is interface{} with range: 0..4294967295.
    NativeVlan interface{}

    // Duplex setting. The type is CdpDuplex.
    Duplex interface{}

    // SysName. The type is string.
    SystemName interface{}

    // List of network addresses .
    NetworkAddresses Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses

    // List of protocol hello entries.
    ProtocolHelloList Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList
}

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "cdp-neighbor"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/cdp-neighbor/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("network-addresses", types.YChild{"NetworkAddresses", &detail.NetworkAddresses})
    detail.EntityData.Children.Append("protocol-hello-list", types.YChild{"ProtocolHelloList", &detail.ProtocolHelloList})
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("version", types.YLeaf{"Version", detail.Version})
    detail.EntityData.Leafs.Append("vtp-domain", types.YLeaf{"VtpDomain", detail.VtpDomain})
    detail.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", detail.NativeVlan})
    detail.EntityData.Leafs.Append("duplex", types.YLeaf{"Duplex", detail.Duplex})
    detail.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", detail.SystemName})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []*Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/cdp-neighbor/detail/" + networkAddresses.EntityData.SegmentPath
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = types.NewOrderedMap()
    networkAddresses.EntityData.Children.Append("cdp-addr-entry", types.YChild{"CdpAddrEntry", nil})
    for i := range networkAddresses.CdpAddrEntry {
        types.SetYListKey(networkAddresses.CdpAddrEntry[i], i)
        networkAddresses.EntityData.Children.Append(types.GetSegmentPath(networkAddresses.CdpAddrEntry[i]), types.YChild{"CdpAddrEntry", networkAddresses.CdpAddrEntry[i]})
    }
    networkAddresses.EntityData.Leafs = types.NewOrderedMap()

    networkAddresses.EntityData.YListKeys = []string {}

    return &(networkAddresses.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetEntityData() *types.CommonEntityData {
    cdpAddrEntry.EntityData.YFilter = cdpAddrEntry.YFilter
    cdpAddrEntry.EntityData.YangName = "cdp-addr-entry"
    cdpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpAddrEntry.EntityData.ParentYangName = "network-addresses"
    cdpAddrEntry.EntityData.SegmentPath = "cdp-addr-entry" + types.AddNoKeyToken(cdpAddrEntry)
    cdpAddrEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/cdp-neighbor/detail/network-addresses/" + cdpAddrEntry.EntityData.SegmentPath
    cdpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpAddrEntry.EntityData.Children = types.NewOrderedMap()
    cdpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &cdpAddrEntry.Address})
    cdpAddrEntry.EntityData.Leafs = types.NewOrderedMap()

    cdpAddrEntry.EntityData.YListKeys = []string {}

    return &(cdpAddrEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is CdpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "cdp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/cdp-neighbor/detail/network-addresses/cdp-addr-entry/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []*Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetEntityData() *types.CommonEntityData {
    protocolHelloList.EntityData.YFilter = protocolHelloList.YFilter
    protocolHelloList.EntityData.YangName = "protocol-hello-list"
    protocolHelloList.EntityData.BundleName = "cisco_ios_xr"
    protocolHelloList.EntityData.ParentYangName = "detail"
    protocolHelloList.EntityData.SegmentPath = "protocol-hello-list"
    protocolHelloList.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/cdp-neighbor/detail/" + protocolHelloList.EntityData.SegmentPath
    protocolHelloList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolHelloList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolHelloList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolHelloList.EntityData.Children = types.NewOrderedMap()
    protocolHelloList.EntityData.Children.Append("cdp-prot-hello-entry", types.YChild{"CdpProtHelloEntry", nil})
    for i := range protocolHelloList.CdpProtHelloEntry {
        types.SetYListKey(protocolHelloList.CdpProtHelloEntry[i], i)
        protocolHelloList.EntityData.Children.Append(types.GetSegmentPath(protocolHelloList.CdpProtHelloEntry[i]), types.YChild{"CdpProtHelloEntry", protocolHelloList.CdpProtHelloEntry[i]})
    }
    protocolHelloList.EntityData.Leafs = types.NewOrderedMap()

    protocolHelloList.EntityData.YListKeys = []string {}

    return &(protocolHelloList.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Protocol Hello msg. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetEntityData() *types.CommonEntityData {
    cdpProtHelloEntry.EntityData.YFilter = cdpProtHelloEntry.YFilter
    cdpProtHelloEntry.EntityData.YangName = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpProtHelloEntry.EntityData.ParentYangName = "protocol-hello-list"
    cdpProtHelloEntry.EntityData.SegmentPath = "cdp-prot-hello-entry" + types.AddNoKeyToken(cdpProtHelloEntry)
    cdpProtHelloEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/devices/device/cdp-neighbor/detail/protocol-hello-list/" + cdpProtHelloEntry.EntityData.SegmentPath
    cdpProtHelloEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpProtHelloEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpProtHelloEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpProtHelloEntry.EntityData.Children = types.NewOrderedMap()
    cdpProtHelloEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpProtHelloEntry.EntityData.Leafs.Append("hello-message", types.YLeaf{"HelloMessage", cdpProtHelloEntry.HelloMessage})

    cdpProtHelloEntry.EntityData.YListKeys = []string {}

    return &(cdpProtHelloEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries
// The CDP neighbor summary table
type Cdp_Nodes_Node_Neighbors_Summaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary.
    Summary []*Cdp_Nodes_Node_Neighbors_Summaries_Summary
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetEntityData() *types.CommonEntityData {
    summaries.EntityData.YFilter = summaries.YFilter
    summaries.EntityData.YangName = "summaries"
    summaries.EntityData.BundleName = "cisco_ios_xr"
    summaries.EntityData.ParentYangName = "neighbors"
    summaries.EntityData.SegmentPath = "summaries"
    summaries.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/" + summaries.EntityData.SegmentPath
    summaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaries.EntityData.Children = types.NewOrderedMap()
    summaries.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range summaries.Summary {
        types.SetYListKey(summaries.Summary[i], i)
        summaries.EntityData.Children.Append(types.GetSegmentPath(summaries.Summary[i]), types.YChild{"Summary", summaries.Summary[i]})
    }
    summaries.EntityData.Leafs = types.NewOrderedMap()

    summaries.EntityData.YListKeys = []string {}

    return &(summaries.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary
// Brief information about a CDP neighbor entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor.
    CdpNeighbor []*Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "summaries"
    summary.EntityData.SegmentPath = "summary" + types.AddNoKeyToken(summary)
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("cdp-neighbor", types.YChild{"CdpNeighbor", nil})
    for i := range summary.CdpNeighbor {
        types.SetYListKey(summary.CdpNeighbor[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.CdpNeighbor[i]), types.YChild{"CdpNeighbor", summary.CdpNeighbor[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", summary.InterfaceName})
    summary.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", summary.DeviceId})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ReceivingInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Outgoing port identifier. The type is string.
    PortId interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Capabilities. The type is string.
    Capabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetEntityData() *types.CommonEntityData {
    cdpNeighbor.EntityData.YFilter = cdpNeighbor.YFilter
    cdpNeighbor.EntityData.YangName = "cdp-neighbor"
    cdpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    cdpNeighbor.EntityData.ParentYangName = "summary"
    cdpNeighbor.EntityData.SegmentPath = "cdp-neighbor" + types.AddNoKeyToken(cdpNeighbor)
    cdpNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/" + cdpNeighbor.EntityData.SegmentPath
    cdpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpNeighbor.EntityData.Children = types.NewOrderedMap()
    cdpNeighbor.EntityData.Children.Append("detail", types.YChild{"Detail", &cdpNeighbor.Detail})
    cdpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    cdpNeighbor.EntityData.Leafs.Append("receiving-interface-name", types.YLeaf{"ReceivingInterfaceName", cdpNeighbor.ReceivingInterfaceName})
    cdpNeighbor.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", cdpNeighbor.DeviceId})
    cdpNeighbor.EntityData.Leafs.Append("port-id", types.YLeaf{"PortId", cdpNeighbor.PortId})
    cdpNeighbor.EntityData.Leafs.Append("header-version", types.YLeaf{"HeaderVersion", cdpNeighbor.HeaderVersion})
    cdpNeighbor.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", cdpNeighbor.HoldTime})
    cdpNeighbor.EntityData.Leafs.Append("capabilities", types.YLeaf{"Capabilities", cdpNeighbor.Capabilities})
    cdpNeighbor.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", cdpNeighbor.Platform})

    cdpNeighbor.EntityData.YListKeys = []string {}

    return &(cdpNeighbor.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version TLV. The type is string.
    Version interface{}

    // VTP domain. The type is string.
    VtpDomain interface{}

    // Native VLAN. The type is interface{} with range: 0..4294967295.
    NativeVlan interface{}

    // Duplex setting. The type is CdpDuplex.
    Duplex interface{}

    // SysName. The type is string.
    SystemName interface{}

    // List of network addresses .
    NetworkAddresses Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses

    // List of protocol hello entries.
    ProtocolHelloList Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList
}

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "cdp-neighbor"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/cdp-neighbor/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("network-addresses", types.YChild{"NetworkAddresses", &detail.NetworkAddresses})
    detail.EntityData.Children.Append("protocol-hello-list", types.YChild{"ProtocolHelloList", &detail.ProtocolHelloList})
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("version", types.YLeaf{"Version", detail.Version})
    detail.EntityData.Leafs.Append("vtp-domain", types.YLeaf{"VtpDomain", detail.VtpDomain})
    detail.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", detail.NativeVlan})
    detail.EntityData.Leafs.Append("duplex", types.YLeaf{"Duplex", detail.Duplex})
    detail.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", detail.SystemName})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []*Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/cdp-neighbor/detail/" + networkAddresses.EntityData.SegmentPath
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = types.NewOrderedMap()
    networkAddresses.EntityData.Children.Append("cdp-addr-entry", types.YChild{"CdpAddrEntry", nil})
    for i := range networkAddresses.CdpAddrEntry {
        types.SetYListKey(networkAddresses.CdpAddrEntry[i], i)
        networkAddresses.EntityData.Children.Append(types.GetSegmentPath(networkAddresses.CdpAddrEntry[i]), types.YChild{"CdpAddrEntry", networkAddresses.CdpAddrEntry[i]})
    }
    networkAddresses.EntityData.Leafs = types.NewOrderedMap()

    networkAddresses.EntityData.YListKeys = []string {}

    return &(networkAddresses.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetEntityData() *types.CommonEntityData {
    cdpAddrEntry.EntityData.YFilter = cdpAddrEntry.YFilter
    cdpAddrEntry.EntityData.YangName = "cdp-addr-entry"
    cdpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpAddrEntry.EntityData.ParentYangName = "network-addresses"
    cdpAddrEntry.EntityData.SegmentPath = "cdp-addr-entry" + types.AddNoKeyToken(cdpAddrEntry)
    cdpAddrEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/cdp-neighbor/detail/network-addresses/" + cdpAddrEntry.EntityData.SegmentPath
    cdpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpAddrEntry.EntityData.Children = types.NewOrderedMap()
    cdpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &cdpAddrEntry.Address})
    cdpAddrEntry.EntityData.Leafs = types.NewOrderedMap()

    cdpAddrEntry.EntityData.YListKeys = []string {}

    return &(cdpAddrEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is CdpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "cdp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/cdp-neighbor/detail/network-addresses/cdp-addr-entry/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []*Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetEntityData() *types.CommonEntityData {
    protocolHelloList.EntityData.YFilter = protocolHelloList.YFilter
    protocolHelloList.EntityData.YangName = "protocol-hello-list"
    protocolHelloList.EntityData.BundleName = "cisco_ios_xr"
    protocolHelloList.EntityData.ParentYangName = "detail"
    protocolHelloList.EntityData.SegmentPath = "protocol-hello-list"
    protocolHelloList.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/cdp-neighbor/detail/" + protocolHelloList.EntityData.SegmentPath
    protocolHelloList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolHelloList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolHelloList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolHelloList.EntityData.Children = types.NewOrderedMap()
    protocolHelloList.EntityData.Children.Append("cdp-prot-hello-entry", types.YChild{"CdpProtHelloEntry", nil})
    for i := range protocolHelloList.CdpProtHelloEntry {
        types.SetYListKey(protocolHelloList.CdpProtHelloEntry[i], i)
        protocolHelloList.EntityData.Children.Append(types.GetSegmentPath(protocolHelloList.CdpProtHelloEntry[i]), types.YChild{"CdpProtHelloEntry", protocolHelloList.CdpProtHelloEntry[i]})
    }
    protocolHelloList.EntityData.Leafs = types.NewOrderedMap()

    protocolHelloList.EntityData.YListKeys = []string {}

    return &(protocolHelloList.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Protocol Hello msg. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetEntityData() *types.CommonEntityData {
    cdpProtHelloEntry.EntityData.YFilter = cdpProtHelloEntry.YFilter
    cdpProtHelloEntry.EntityData.YangName = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpProtHelloEntry.EntityData.ParentYangName = "protocol-hello-list"
    cdpProtHelloEntry.EntityData.SegmentPath = "cdp-prot-hello-entry" + types.AddNoKeyToken(cdpProtHelloEntry)
    cdpProtHelloEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/neighbors/summaries/summary/cdp-neighbor/detail/protocol-hello-list/" + cdpProtHelloEntry.EntityData.SegmentPath
    cdpProtHelloEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpProtHelloEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpProtHelloEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpProtHelloEntry.EntityData.Children = types.NewOrderedMap()
    cdpProtHelloEntry.EntityData.Leafs = types.NewOrderedMap()
    cdpProtHelloEntry.EntityData.Leafs.Append("hello-message", types.YLeaf{"HelloMessage", cdpProtHelloEntry.HelloMessage})

    cdpProtHelloEntry.EntityData.YListKeys = []string {}

    return &(cdpProtHelloEntry.EntityData)
}

// Cdp_Nodes_Node_Statistics
// The CDP traffic statistics for this node
type Cdp_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Received v1 packets. The type is interface{} with range: 0..4294967295.
    ReceivedPacketsV1 interface{}

    // Received v2 packets. The type is interface{} with range: 0..4294967295.
    ReceivedPacketsV2 interface{}

    // Transmitted packets. The type is interface{} with range: 0..4294967295.
    TransmittedPackets interface{}

    // Transmitted v1 packets. The type is interface{} with range: 0..4294967295.
    TransmittedPacketsV1 interface{}

    // Transmitted v2 packets. The type is interface{} with range: 0..4294967295.
    TransmittedPacketsV2 interface{}

    // Header syntax errors. The type is interface{} with range: 0..4294967295.
    HeaderErrors interface{}

    // Checksum errors. The type is interface{} with range: 0..4294967295.
    ChecksumErrors interface{}

    // Transmission errors. The type is interface{} with range: 0..4294967295.
    EncapsulationErrors interface{}

    // Bad packet received and dropped. The type is interface{} with range:
    // 0..4294967295.
    BadPacketErrors interface{}

    // Out-of-memory conditions. The type is interface{} with range:
    // 0..4294967295.
    OutOfMemoryErrors interface{}

    // Truncated messages. The type is interface{} with range: 0..4294967295.
    TruncatedPacketErrors interface{}

    // Can't handle receive version. The type is interface{} with range:
    // 0..4294967295.
    HeaderVersionErrors interface{}

    // Cannot open file. The type is interface{} with range: 0..4294967295.
    OpenFileErrors interface{}
}

func (statistics *Cdp_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets})
    statistics.EntityData.Leafs.Append("received-packets-v1", types.YLeaf{"ReceivedPacketsV1", statistics.ReceivedPacketsV1})
    statistics.EntityData.Leafs.Append("received-packets-v2", types.YLeaf{"ReceivedPacketsV2", statistics.ReceivedPacketsV2})
    statistics.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", statistics.TransmittedPackets})
    statistics.EntityData.Leafs.Append("transmitted-packets-v1", types.YLeaf{"TransmittedPacketsV1", statistics.TransmittedPacketsV1})
    statistics.EntityData.Leafs.Append("transmitted-packets-v2", types.YLeaf{"TransmittedPacketsV2", statistics.TransmittedPacketsV2})
    statistics.EntityData.Leafs.Append("header-errors", types.YLeaf{"HeaderErrors", statistics.HeaderErrors})
    statistics.EntityData.Leafs.Append("checksum-errors", types.YLeaf{"ChecksumErrors", statistics.ChecksumErrors})
    statistics.EntityData.Leafs.Append("encapsulation-errors", types.YLeaf{"EncapsulationErrors", statistics.EncapsulationErrors})
    statistics.EntityData.Leafs.Append("bad-packet-errors", types.YLeaf{"BadPacketErrors", statistics.BadPacketErrors})
    statistics.EntityData.Leafs.Append("out-of-memory-errors", types.YLeaf{"OutOfMemoryErrors", statistics.OutOfMemoryErrors})
    statistics.EntityData.Leafs.Append("truncated-packet-errors", types.YLeaf{"TruncatedPacketErrors", statistics.TruncatedPacketErrors})
    statistics.EntityData.Leafs.Append("header-version-errors", types.YLeaf{"HeaderVersionErrors", statistics.HeaderVersionErrors})
    statistics.EntityData.Leafs.Append("open-file-errors", types.YLeaf{"OpenFileErrors", statistics.OpenFileErrors})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Cdp_Nodes_Node_Interfaces
// The table of interfaces on which CDP is
// running on this node
type Cdp_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for an interface on which CDP is running. The type is
    // slice of Cdp_Nodes_Node_Interfaces_Interface.
    Interface []*Cdp_Nodes_Node_Interfaces_Interface
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/" + interfaces.EntityData.SegmentPath
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

// Cdp_Nodes_Node_Interfaces_Interface
// Operational data for an interface on which
// CDP is running
type Cdp_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // Interface basecaps state. The type is interface{} with range:
    // 0..4294967295.
    BasecapsState interface{}

    // CDP protocol state. The type is interface{} with range: 0..4294967295.
    CdpProtocolState interface{}

    // Interface encapsulation. The type is string.
    InterfaceEncaps interface{}
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-cdp-oper:cdp/nodes/node/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", self.InterfaceHandle})
    self.EntityData.Leafs.Append("basecaps-state", types.YLeaf{"BasecapsState", self.BasecapsState})
    self.EntityData.Leafs.Append("cdp-protocol-state", types.YLeaf{"CdpProtocolState", self.CdpProtocolState})
    self.EntityData.Leafs.Append("interface-encaps", types.YLeaf{"InterfaceEncaps", self.InterfaceEncaps})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

