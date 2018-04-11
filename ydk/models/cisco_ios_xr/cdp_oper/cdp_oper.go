// This module contains a collection of YANG definitions
// for Cisco IOS-XR cdp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cdp: CDP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    cdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdp.EntityData.Children = make(map[string]types.YChild)
    cdp.EntityData.Children["nodes"] = types.YChild{"Nodes", &cdp.Nodes}
    cdp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdp.EntityData)
}

// Cdp_Nodes
// Per node CDP operational data
type Cdp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The CDP operational data for a particular node. The type is slice of
    // Cdp_Nodes_Node.
    Node []Cdp_Nodes_Node
}

func (nodes *Cdp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cdp"
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

// Cdp_Nodes_Node
// The CDP operational data for a particular node
type Cdp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &node.Neighbors}
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
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
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["details"] = types.YChild{"Details", &neighbors.Details}
    neighbors.EntityData.Children["devices"] = types.YChild{"Devices", &neighbors.Devices}
    neighbors.EntityData.Children["summaries"] = types.YChild{"Summaries", &neighbors.Summaries}
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details
// The detailed CDP neighbor table
type Cdp_Nodes_Node_Neighbors_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail.
    Detail []Cdp_Nodes_Node_Neighbors_Details_Detail
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "neighbors"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range details.Detail {
        details.EntityData.Children[types.GetSegmentPath(&details.Detail[i])] = types.YChild{"Detail", &details.Detail[i]}
    }
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail
// Detailed information about a CDP neighbor
// entry
type Cdp_Nodes_Node_Neighbors_Details_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor.
    CdpNeighbor []Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "details"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["cdp-neighbor"] = types.YChild{"CdpNeighbor", nil}
    for i := range detail.CdpNeighbor {
        detail.EntityData.Children[types.GetSegmentPath(&detail.CdpNeighbor[i])] = types.YChild{"CdpNeighbor", &detail.CdpNeighbor[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", detail.InterfaceName}
    detail.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", detail.DeviceId}
    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
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
    Detail Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetEntityData() *types.CommonEntityData {
    cdpNeighbor.EntityData.YFilter = cdpNeighbor.YFilter
    cdpNeighbor.EntityData.YangName = "cdp-neighbor"
    cdpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    cdpNeighbor.EntityData.ParentYangName = "detail"
    cdpNeighbor.EntityData.SegmentPath = "cdp-neighbor"
    cdpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpNeighbor.EntityData.Children = make(map[string]types.YChild)
    cdpNeighbor.EntityData.Children["detail"] = types.YChild{"Detail", &cdpNeighbor.Detail}
    cdpNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpNeighbor.EntityData.Leafs["receiving-interface-name"] = types.YLeaf{"ReceivingInterfaceName", cdpNeighbor.ReceivingInterfaceName}
    cdpNeighbor.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", cdpNeighbor.DeviceId}
    cdpNeighbor.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", cdpNeighbor.PortId}
    cdpNeighbor.EntityData.Leafs["header-version"] = types.YLeaf{"HeaderVersion", cdpNeighbor.HeaderVersion}
    cdpNeighbor.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", cdpNeighbor.HoldTime}
    cdpNeighbor.EntityData.Leafs["capabilities"] = types.YLeaf{"Capabilities", cdpNeighbor.Capabilities}
    cdpNeighbor.EntityData.Leafs["platform"] = types.YLeaf{"Platform", cdpNeighbor.Platform}
    return &(cdpNeighbor.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ struct {
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
    NetworkAddresses Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses

    // List of protocol hello entries.
    ProtocolHelloList Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList
}

func (detail_ *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_) GetEntityData() *types.CommonEntityData {
    detail_.EntityData.YFilter = detail_.YFilter
    detail_.EntityData.YangName = "detail"
    detail_.EntityData.BundleName = "cisco_ios_xr"
    detail_.EntityData.ParentYangName = "cdp-neighbor"
    detail_.EntityData.SegmentPath = "detail"
    detail_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail_.EntityData.Children = make(map[string]types.YChild)
    detail_.EntityData.Children["network-addresses"] = types.YChild{"NetworkAddresses", &detail_.NetworkAddresses}
    detail_.EntityData.Children["protocol-hello-list"] = types.YChild{"ProtocolHelloList", &detail_.ProtocolHelloList}
    detail_.EntityData.Leafs = make(map[string]types.YLeaf)
    detail_.EntityData.Leafs["version"] = types.YLeaf{"Version", detail_.Version}
    detail_.EntityData.Leafs["vtp-domain"] = types.YLeaf{"VtpDomain", detail_.VtpDomain}
    detail_.EntityData.Leafs["native-vlan"] = types.YLeaf{"NativeVlan", detail_.NativeVlan}
    detail_.EntityData.Leafs["duplex"] = types.YLeaf{"Duplex", detail_.Duplex}
    detail_.EntityData.Leafs["system-name"] = types.YLeaf{"SystemName", detail_.SystemName}
    return &(detail_.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = make(map[string]types.YChild)
    networkAddresses.EntityData.Children["cdp-addr-entry"] = types.YChild{"CdpAddrEntry", nil}
    for i := range networkAddresses.CdpAddrEntry {
        networkAddresses.EntityData.Children[types.GetSegmentPath(&networkAddresses.CdpAddrEntry[i])] = types.YChild{"CdpAddrEntry", &networkAddresses.CdpAddrEntry[i]}
    }
    networkAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networkAddresses.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry) GetEntityData() *types.CommonEntityData {
    cdpAddrEntry.EntityData.YFilter = cdpAddrEntry.YFilter
    cdpAddrEntry.EntityData.YangName = "cdp-addr-entry"
    cdpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpAddrEntry.EntityData.ParentYangName = "network-addresses"
    cdpAddrEntry.EntityData.SegmentPath = "cdp-addr-entry"
    cdpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpAddrEntry.EntityData.Children = make(map[string]types.YChild)
    cdpAddrEntry.EntityData.Children["address"] = types.YChild{"Address", &cdpAddrEntry.Address}
    cdpAddrEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdpAddrEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is CdpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__NetworkAddresses_CdpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "cdp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-type"] = types.YLeaf{"AddressType", address.AddressType}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList) GetEntityData() *types.CommonEntityData {
    protocolHelloList.EntityData.YFilter = protocolHelloList.YFilter
    protocolHelloList.EntityData.YangName = "protocol-hello-list"
    protocolHelloList.EntityData.BundleName = "cisco_ios_xr"
    protocolHelloList.EntityData.ParentYangName = "detail"
    protocolHelloList.EntityData.SegmentPath = "protocol-hello-list"
    protocolHelloList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolHelloList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolHelloList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolHelloList.EntityData.Children = make(map[string]types.YChild)
    protocolHelloList.EntityData.Children["cdp-prot-hello-entry"] = types.YChild{"CdpProtHelloEntry", nil}
    for i := range protocolHelloList.CdpProtHelloEntry {
        protocolHelloList.EntityData.Children[types.GetSegmentPath(&protocolHelloList.CdpProtHelloEntry[i])] = types.YChild{"CdpProtHelloEntry", &protocolHelloList.CdpProtHelloEntry[i]}
    }
    protocolHelloList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocolHelloList.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList_CdpProtHelloEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Hello msg. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail__ProtocolHelloList_CdpProtHelloEntry) GetEntityData() *types.CommonEntityData {
    cdpProtHelloEntry.EntityData.YFilter = cdpProtHelloEntry.YFilter
    cdpProtHelloEntry.EntityData.YangName = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpProtHelloEntry.EntityData.ParentYangName = "protocol-hello-list"
    cdpProtHelloEntry.EntityData.SegmentPath = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpProtHelloEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpProtHelloEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpProtHelloEntry.EntityData.Children = make(map[string]types.YChild)
    cdpProtHelloEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpProtHelloEntry.EntityData.Leafs["hello-message"] = types.YLeaf{"HelloMessage", cdpProtHelloEntry.HelloMessage}
    return &(cdpProtHelloEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices
// The detailed CDP neighbor table
type Cdp_Nodes_Node_Neighbors_Devices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device.
    Device []Cdp_Nodes_Node_Neighbors_Devices_Device
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetEntityData() *types.CommonEntityData {
    devices.EntityData.YFilter = devices.YFilter
    devices.EntityData.YangName = "devices"
    devices.EntityData.BundleName = "cisco_ios_xr"
    devices.EntityData.ParentYangName = "neighbors"
    devices.EntityData.SegmentPath = "devices"
    devices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devices.EntityData.Children = make(map[string]types.YChild)
    devices.EntityData.Children["device"] = types.YChild{"Device", nil}
    for i := range devices.Device {
        devices.EntityData.Children[types.GetSegmentPath(&devices.Device[i])] = types.YChild{"Device", &devices.Device[i]}
    }
    devices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(devices.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device
// Detailed information about a CDP neighbor
// entry
type Cdp_Nodes_Node_Neighbors_Devices_Device struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The neighboring device identifier. The type is
    // string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor.
    CdpNeighbor []Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetEntityData() *types.CommonEntityData {
    device.EntityData.YFilter = device.YFilter
    device.EntityData.YangName = "device"
    device.EntityData.BundleName = "cisco_ios_xr"
    device.EntityData.ParentYangName = "devices"
    device.EntityData.SegmentPath = "device" + "[device-id='" + fmt.Sprintf("%v", device.DeviceId) + "']"
    device.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    device.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    device.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    device.EntityData.Children = make(map[string]types.YChild)
    device.EntityData.Children["cdp-neighbor"] = types.YChild{"CdpNeighbor", nil}
    for i := range device.CdpNeighbor {
        device.EntityData.Children[types.GetSegmentPath(&device.CdpNeighbor[i])] = types.YChild{"CdpNeighbor", &device.CdpNeighbor[i]}
    }
    device.EntityData.Leafs = make(map[string]types.YLeaf)
    device.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", device.DeviceId}
    return &(device.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
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
    cdpNeighbor.EntityData.SegmentPath = "cdp-neighbor"
    cdpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpNeighbor.EntityData.Children = make(map[string]types.YChild)
    cdpNeighbor.EntityData.Children["detail"] = types.YChild{"Detail", &cdpNeighbor.Detail}
    cdpNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpNeighbor.EntityData.Leafs["receiving-interface-name"] = types.YLeaf{"ReceivingInterfaceName", cdpNeighbor.ReceivingInterfaceName}
    cdpNeighbor.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", cdpNeighbor.DeviceId}
    cdpNeighbor.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", cdpNeighbor.PortId}
    cdpNeighbor.EntityData.Leafs["header-version"] = types.YLeaf{"HeaderVersion", cdpNeighbor.HeaderVersion}
    cdpNeighbor.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", cdpNeighbor.HoldTime}
    cdpNeighbor.EntityData.Leafs["capabilities"] = types.YLeaf{"Capabilities", cdpNeighbor.Capabilities}
    cdpNeighbor.EntityData.Leafs["platform"] = types.YLeaf{"Platform", cdpNeighbor.Platform}
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
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["network-addresses"] = types.YChild{"NetworkAddresses", &detail.NetworkAddresses}
    detail.EntityData.Children["protocol-hello-list"] = types.YChild{"ProtocolHelloList", &detail.ProtocolHelloList}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["version"] = types.YLeaf{"Version", detail.Version}
    detail.EntityData.Leafs["vtp-domain"] = types.YLeaf{"VtpDomain", detail.VtpDomain}
    detail.EntityData.Leafs["native-vlan"] = types.YLeaf{"NativeVlan", detail.NativeVlan}
    detail.EntityData.Leafs["duplex"] = types.YLeaf{"Duplex", detail.Duplex}
    detail.EntityData.Leafs["system-name"] = types.YLeaf{"SystemName", detail.SystemName}
    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = make(map[string]types.YChild)
    networkAddresses.EntityData.Children["cdp-addr-entry"] = types.YChild{"CdpAddrEntry", nil}
    for i := range networkAddresses.CdpAddrEntry {
        networkAddresses.EntityData.Children[types.GetSegmentPath(&networkAddresses.CdpAddrEntry[i])] = types.YChild{"CdpAddrEntry", &networkAddresses.CdpAddrEntry[i]}
    }
    networkAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networkAddresses.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetEntityData() *types.CommonEntityData {
    cdpAddrEntry.EntityData.YFilter = cdpAddrEntry.YFilter
    cdpAddrEntry.EntityData.YangName = "cdp-addr-entry"
    cdpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpAddrEntry.EntityData.ParentYangName = "network-addresses"
    cdpAddrEntry.EntityData.SegmentPath = "cdp-addr-entry"
    cdpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpAddrEntry.EntityData.Children = make(map[string]types.YChild)
    cdpAddrEntry.EntityData.Children["address"] = types.YChild{"Address", &cdpAddrEntry.Address}
    cdpAddrEntry.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "cdp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-type"] = types.YLeaf{"AddressType", address.AddressType}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetEntityData() *types.CommonEntityData {
    protocolHelloList.EntityData.YFilter = protocolHelloList.YFilter
    protocolHelloList.EntityData.YangName = "protocol-hello-list"
    protocolHelloList.EntityData.BundleName = "cisco_ios_xr"
    protocolHelloList.EntityData.ParentYangName = "detail"
    protocolHelloList.EntityData.SegmentPath = "protocol-hello-list"
    protocolHelloList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolHelloList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolHelloList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolHelloList.EntityData.Children = make(map[string]types.YChild)
    protocolHelloList.EntityData.Children["cdp-prot-hello-entry"] = types.YChild{"CdpProtHelloEntry", nil}
    for i := range protocolHelloList.CdpProtHelloEntry {
        protocolHelloList.EntityData.Children[types.GetSegmentPath(&protocolHelloList.CdpProtHelloEntry[i])] = types.YChild{"CdpProtHelloEntry", &protocolHelloList.CdpProtHelloEntry[i]}
    }
    protocolHelloList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocolHelloList.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Hello msg. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetEntityData() *types.CommonEntityData {
    cdpProtHelloEntry.EntityData.YFilter = cdpProtHelloEntry.YFilter
    cdpProtHelloEntry.EntityData.YangName = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpProtHelloEntry.EntityData.ParentYangName = "protocol-hello-list"
    cdpProtHelloEntry.EntityData.SegmentPath = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpProtHelloEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpProtHelloEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpProtHelloEntry.EntityData.Children = make(map[string]types.YChild)
    cdpProtHelloEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpProtHelloEntry.EntityData.Leafs["hello-message"] = types.YLeaf{"HelloMessage", cdpProtHelloEntry.HelloMessage}
    return &(cdpProtHelloEntry.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries
// The CDP neighbor summary table
type Cdp_Nodes_Node_Neighbors_Summaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary.
    Summary []Cdp_Nodes_Node_Neighbors_Summaries_Summary
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetEntityData() *types.CommonEntityData {
    summaries.EntityData.YFilter = summaries.YFilter
    summaries.EntityData.YangName = "summaries"
    summaries.EntityData.BundleName = "cisco_ios_xr"
    summaries.EntityData.ParentYangName = "neighbors"
    summaries.EntityData.SegmentPath = "summaries"
    summaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaries.EntityData.Children = make(map[string]types.YChild)
    summaries.EntityData.Children["summary"] = types.YChild{"Summary", nil}
    for i := range summaries.Summary {
        summaries.EntityData.Children[types.GetSegmentPath(&summaries.Summary[i])] = types.YChild{"Summary", &summaries.Summary[i]}
    }
    summaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summaries.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary
// Brief information about a CDP neighbor entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor.
    CdpNeighbor []Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "summaries"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["cdp-neighbor"] = types.YChild{"CdpNeighbor", nil}
    for i := range summary.CdpNeighbor {
        summary.EntityData.Children[types.GetSegmentPath(&summary.CdpNeighbor[i])] = types.YChild{"CdpNeighbor", &summary.CdpNeighbor[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", summary.InterfaceName}
    summary.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", summary.DeviceId}
    return &(summary.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
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
    cdpNeighbor.EntityData.SegmentPath = "cdp-neighbor"
    cdpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpNeighbor.EntityData.Children = make(map[string]types.YChild)
    cdpNeighbor.EntityData.Children["detail"] = types.YChild{"Detail", &cdpNeighbor.Detail}
    cdpNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpNeighbor.EntityData.Leafs["receiving-interface-name"] = types.YLeaf{"ReceivingInterfaceName", cdpNeighbor.ReceivingInterfaceName}
    cdpNeighbor.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", cdpNeighbor.DeviceId}
    cdpNeighbor.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", cdpNeighbor.PortId}
    cdpNeighbor.EntityData.Leafs["header-version"] = types.YLeaf{"HeaderVersion", cdpNeighbor.HeaderVersion}
    cdpNeighbor.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", cdpNeighbor.HoldTime}
    cdpNeighbor.EntityData.Leafs["capabilities"] = types.YLeaf{"Capabilities", cdpNeighbor.Capabilities}
    cdpNeighbor.EntityData.Leafs["platform"] = types.YLeaf{"Platform", cdpNeighbor.Platform}
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
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["network-addresses"] = types.YChild{"NetworkAddresses", &detail.NetworkAddresses}
    detail.EntityData.Children["protocol-hello-list"] = types.YChild{"ProtocolHelloList", &detail.ProtocolHelloList}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["version"] = types.YLeaf{"Version", detail.Version}
    detail.EntityData.Leafs["vtp-domain"] = types.YLeaf{"VtpDomain", detail.VtpDomain}
    detail.EntityData.Leafs["native-vlan"] = types.YLeaf{"NativeVlan", detail.NativeVlan}
    detail.EntityData.Leafs["duplex"] = types.YLeaf{"Duplex", detail.Duplex}
    detail.EntityData.Leafs["system-name"] = types.YLeaf{"SystemName", detail.SystemName}
    return &(detail.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = make(map[string]types.YChild)
    networkAddresses.EntityData.Children["cdp-addr-entry"] = types.YChild{"CdpAddrEntry", nil}
    for i := range networkAddresses.CdpAddrEntry {
        networkAddresses.EntityData.Children[types.GetSegmentPath(&networkAddresses.CdpAddrEntry[i])] = types.YChild{"CdpAddrEntry", &networkAddresses.CdpAddrEntry[i]}
    }
    networkAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networkAddresses.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetEntityData() *types.CommonEntityData {
    cdpAddrEntry.EntityData.YFilter = cdpAddrEntry.YFilter
    cdpAddrEntry.EntityData.YangName = "cdp-addr-entry"
    cdpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpAddrEntry.EntityData.ParentYangName = "network-addresses"
    cdpAddrEntry.EntityData.SegmentPath = "cdp-addr-entry"
    cdpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpAddrEntry.EntityData.Children = make(map[string]types.YChild)
    cdpAddrEntry.EntityData.Children["address"] = types.YChild{"Address", &cdpAddrEntry.Address}
    cdpAddrEntry.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "cdp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-type"] = types.YLeaf{"AddressType", address.AddressType}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetEntityData() *types.CommonEntityData {
    protocolHelloList.EntityData.YFilter = protocolHelloList.YFilter
    protocolHelloList.EntityData.YangName = "protocol-hello-list"
    protocolHelloList.EntityData.BundleName = "cisco_ios_xr"
    protocolHelloList.EntityData.ParentYangName = "detail"
    protocolHelloList.EntityData.SegmentPath = "protocol-hello-list"
    protocolHelloList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolHelloList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolHelloList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolHelloList.EntityData.Children = make(map[string]types.YChild)
    protocolHelloList.EntityData.Children["cdp-prot-hello-entry"] = types.YChild{"CdpProtHelloEntry", nil}
    for i := range protocolHelloList.CdpProtHelloEntry {
        protocolHelloList.EntityData.Children[types.GetSegmentPath(&protocolHelloList.CdpProtHelloEntry[i])] = types.YChild{"CdpProtHelloEntry", &protocolHelloList.CdpProtHelloEntry[i]}
    }
    protocolHelloList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocolHelloList.EntityData)
}

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Hello msg. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetEntityData() *types.CommonEntityData {
    cdpProtHelloEntry.EntityData.YFilter = cdpProtHelloEntry.YFilter
    cdpProtHelloEntry.EntityData.YangName = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.BundleName = "cisco_ios_xr"
    cdpProtHelloEntry.EntityData.ParentYangName = "protocol-hello-list"
    cdpProtHelloEntry.EntityData.SegmentPath = "cdp-prot-hello-entry"
    cdpProtHelloEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdpProtHelloEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdpProtHelloEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdpProtHelloEntry.EntityData.Children = make(map[string]types.YChild)
    cdpProtHelloEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpProtHelloEntry.EntityData.Leafs["hello-message"] = types.YLeaf{"HelloMessage", cdpProtHelloEntry.HelloMessage}
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
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets}
    statistics.EntityData.Leafs["received-packets-v1"] = types.YLeaf{"ReceivedPacketsV1", statistics.ReceivedPacketsV1}
    statistics.EntityData.Leafs["received-packets-v2"] = types.YLeaf{"ReceivedPacketsV2", statistics.ReceivedPacketsV2}
    statistics.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", statistics.TransmittedPackets}
    statistics.EntityData.Leafs["transmitted-packets-v1"] = types.YLeaf{"TransmittedPacketsV1", statistics.TransmittedPacketsV1}
    statistics.EntityData.Leafs["transmitted-packets-v2"] = types.YLeaf{"TransmittedPacketsV2", statistics.TransmittedPacketsV2}
    statistics.EntityData.Leafs["header-errors"] = types.YLeaf{"HeaderErrors", statistics.HeaderErrors}
    statistics.EntityData.Leafs["checksum-errors"] = types.YLeaf{"ChecksumErrors", statistics.ChecksumErrors}
    statistics.EntityData.Leafs["encapsulation-errors"] = types.YLeaf{"EncapsulationErrors", statistics.EncapsulationErrors}
    statistics.EntityData.Leafs["bad-packet-errors"] = types.YLeaf{"BadPacketErrors", statistics.BadPacketErrors}
    statistics.EntityData.Leafs["out-of-memory-errors"] = types.YLeaf{"OutOfMemoryErrors", statistics.OutOfMemoryErrors}
    statistics.EntityData.Leafs["truncated-packet-errors"] = types.YLeaf{"TruncatedPacketErrors", statistics.TruncatedPacketErrors}
    statistics.EntityData.Leafs["header-version-errors"] = types.YLeaf{"HeaderVersionErrors", statistics.HeaderVersionErrors}
    statistics.EntityData.Leafs["open-file-errors"] = types.YLeaf{"OpenFileErrors", statistics.OpenFileErrors}
    return &(statistics.EntityData)
}

// Cdp_Nodes_Node_Interfaces
// The table of interfaces on which CDP is
// running on this node
type Cdp_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for an interface on which CDP is running. The type is
    // slice of Cdp_Nodes_Node_Interfaces_Interface_.
    Interface_ []Cdp_Nodes_Node_Interfaces_Interface
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
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

// Cdp_Nodes_Node_Interfaces_Interface
// Operational data for an interface on which
// CDP is running
type Cdp_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", self.InterfaceHandle}
    self.EntityData.Leafs["basecaps-state"] = types.YLeaf{"BasecapsState", self.BasecapsState}
    self.EntityData.Leafs["cdp-protocol-state"] = types.YLeaf{"CdpProtocolState", self.CdpProtocolState}
    self.EntityData.Leafs["interface-encaps"] = types.YLeaf{"InterfaceEncaps", self.InterfaceEncaps}
    return &(self.EntityData)
}

