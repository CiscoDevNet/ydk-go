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
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node CDP operational data.
    Nodes Cdp_Nodes
}

func (cdp *Cdp) GetFilter() yfilter.YFilter { return cdp.YFilter }

func (cdp *Cdp) SetFilter(yf yfilter.YFilter) { cdp.YFilter = yf }

func (cdp *Cdp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (cdp *Cdp) GetSegmentPath() string {
    return "Cisco-IOS-XR-cdp-oper:cdp"
}

func (cdp *Cdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &cdp.Nodes
    }
    return nil
}

func (cdp *Cdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &cdp.Nodes
    return children
}

func (cdp *Cdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdp *Cdp) GetBundleName() string { return "cisco_ios_xr" }

func (cdp *Cdp) GetYangName() string { return "cdp" }

func (cdp *Cdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdp *Cdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdp *Cdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdp *Cdp) SetParent(parent types.Entity) { cdp.parent = parent }

func (cdp *Cdp) GetParent() types.Entity { return cdp.parent }

func (cdp *Cdp) GetParentYangName() string { return "Cisco-IOS-XR-cdp-oper" }

// Cdp_Nodes
// Per node CDP operational data
type Cdp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The CDP operational data for a particular node. The type is slice of
    // Cdp_Nodes_Node.
    Node []Cdp_Nodes_Node
}

func (nodes *Cdp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Cdp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Cdp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Cdp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Cdp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Cdp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Cdp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Cdp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Cdp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Cdp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Cdp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Cdp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Cdp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Cdp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Cdp_Nodes) GetParentYangName() string { return "cdp" }

// Cdp_Nodes_Node
// The CDP operational data for a particular node
type Cdp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (node *Cdp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Cdp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Cdp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "statistics" { return "Statistics" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *Cdp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Cdp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &node.Neighbors
    }
    if childYangName == "statistics" {
        return &node.Statistics
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *Cdp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &node.Neighbors
    children["statistics"] = &node.Statistics
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *Cdp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Cdp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Cdp_Nodes_Node) GetYangName() string { return "node" }

func (node *Cdp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Cdp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Cdp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Cdp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Cdp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Cdp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Cdp_Nodes_Node_Neighbors
// The CDP neighbor tables on this node
type Cdp_Nodes_Node_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The detailed CDP neighbor table.
    Details Cdp_Nodes_Node_Neighbors_Details

    // The detailed CDP neighbor table.
    Devices Cdp_Nodes_Node_Neighbors_Devices

    // The CDP neighbor summary table.
    Summaries Cdp_Nodes_Node_Neighbors_Summaries
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Cdp_Nodes_Node_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    if yname == "devices" { return "Devices" }
    if yname == "summaries" { return "Summaries" }
    return ""
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &neighbors.Details
    }
    if childYangName == "devices" {
        return &neighbors.Devices
    }
    if childYangName == "summaries" {
        return &neighbors.Summaries
    }
    return nil
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &neighbors.Details
    children["devices"] = &neighbors.Devices
    children["summaries"] = &neighbors.Summaries
    return children
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Cdp_Nodes_Node_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Cdp_Nodes_Node_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Cdp_Nodes_Node_Neighbors) GetParentYangName() string { return "node" }

// Cdp_Nodes_Node_Neighbors_Details
// The detailed CDP neighbor table
type Cdp_Nodes_Node_Neighbors_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail.
    Detail []Cdp_Nodes_Node_Neighbors_Details_Detail
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Cdp_Nodes_Node_Neighbors_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    return ""
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetSegmentPath() string {
    return "details"
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        for _, c := range details.Detail {
            if details.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Details_Detail{}
        details.Detail = append(details.Detail, child)
        return &details.Detail[len(details.Detail)-1]
    }
    return nil
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range details.Detail {
        children[details.Detail[i].GetSegmentPath()] = &details.Detail[i]
    }
    return children
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *Cdp_Nodes_Node_Neighbors_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetYangName() string { return "details" }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Cdp_Nodes_Node_Neighbors_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetParent() types.Entity { return details.parent }

func (details *Cdp_Nodes_Node_Neighbors_Details) GetParentYangName() string { return "neighbors" }

// Cdp_Nodes_Node_Neighbors_Details_Detail
// Detailed information about a CDP neighbor
// entry
type Cdp_Nodes_Node_Neighbors_Details_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor.
    CdpNeighbor []Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "cdp-neighbor" { return "CdpNeighbor" }
    return ""
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-neighbor" {
        for _, c := range detail.CdpNeighbor {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor{}
        detail.CdpNeighbor = append(detail.CdpNeighbor, child)
        return &detail.CdpNeighbor[len(detail.CdpNeighbor)-1]
    }
    return nil
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detail.CdpNeighbor {
        children[detail.CdpNeighbor[i].GetSegmentPath()] = &detail.CdpNeighbor[i]
    }
    return children
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = detail.InterfaceName
    leafs["device-id"] = detail.DeviceId
    return leafs
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetYangName() string { return "detail" }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail) GetParentYangName() string { return "details" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetFilter() yfilter.YFilter { return cdpNeighbor.YFilter }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) SetFilter(yf yfilter.YFilter) { cdpNeighbor.YFilter = yf }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetGoName(yname string) string {
    if yname == "receiving-interface-name" { return "ReceivingInterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "port-id" { return "PortId" }
    if yname == "header-version" { return "HeaderVersion" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "capabilities" { return "Capabilities" }
    if yname == "platform" { return "Platform" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetSegmentPath() string {
    return "cdp-neighbor"
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &cdpNeighbor.Detail
    }
    return nil
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &cdpNeighbor.Detail
    return children
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receiving-interface-name"] = cdpNeighbor.ReceivingInterfaceName
    leafs["device-id"] = cdpNeighbor.DeviceId
    leafs["port-id"] = cdpNeighbor.PortId
    leafs["header-version"] = cdpNeighbor.HeaderVersion
    leafs["hold-time"] = cdpNeighbor.HoldTime
    leafs["capabilities"] = cdpNeighbor.Capabilities
    leafs["platform"] = cdpNeighbor.Platform
    return leafs
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetYangName() string { return "cdp-neighbor" }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) SetParent(parent types.Entity) { cdpNeighbor.parent = parent }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetParent() types.Entity { return cdpNeighbor.parent }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail struct {
    parent types.Entity
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

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "vtp-domain" { return "VtpDomain" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "duplex" { return "Duplex" }
    if yname == "system-name" { return "SystemName" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    if yname == "protocol-hello-list" { return "ProtocolHelloList" }
    return ""
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &detail.NetworkAddresses
    }
    if childYangName == "protocol-hello-list" {
        return &detail.ProtocolHelloList
    }
    return nil
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &detail.NetworkAddresses
    children["protocol-hello-list"] = &detail.ProtocolHelloList
    return children
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = detail.Version
    leafs["vtp-domain"] = detail.VtpDomain
    leafs["native-vlan"] = detail.NativeVlan
    leafs["duplex"] = detail.Duplex
    leafs["system-name"] = detail.SystemName
    return leafs
}

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetYangName() string { return "detail" }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail) GetParentYangName() string { return "cdp-neighbor" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetGoName(yname string) string {
    if yname == "cdp-addr-entry" { return "CdpAddrEntry" }
    return ""
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-addr-entry" {
        for _, c := range networkAddresses.CdpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry{}
        networkAddresses.CdpAddrEntry = append(networkAddresses.CdpAddrEntry, child)
        return &networkAddresses.CdpAddrEntry[len(networkAddresses.CdpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.CdpAddrEntry {
        children[networkAddresses.CdpAddrEntry[i].GetSegmentPath()] = &networkAddresses.CdpAddrEntry[i]
    }
    return children
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetFilter() yfilter.YFilter { return cdpAddrEntry.YFilter }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) SetFilter(yf yfilter.YFilter) { cdpAddrEntry.YFilter = yf }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetSegmentPath() string {
    return "cdp-addr-entry"
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &cdpAddrEntry.Address
    }
    return nil
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &cdpAddrEntry.Address
    return children
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetYangName() string { return "cdp-addr-entry" }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) SetParent(parent types.Entity) { cdpAddrEntry.parent = parent }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetParent() types.Entity { return cdpAddrEntry.parent }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetParentYangName() string { return "network-addresses" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address struct {
    parent types.Entity
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

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetParentYangName() string { return "cdp-addr-entry" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetFilter() yfilter.YFilter { return protocolHelloList.YFilter }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) SetFilter(yf yfilter.YFilter) { protocolHelloList.YFilter = yf }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetGoName(yname string) string {
    if yname == "cdp-prot-hello-entry" { return "CdpProtHelloEntry" }
    return ""
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetSegmentPath() string {
    return "protocol-hello-list"
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-prot-hello-entry" {
        for _, c := range protocolHelloList.CdpProtHelloEntry {
            if protocolHelloList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry{}
        protocolHelloList.CdpProtHelloEntry = append(protocolHelloList.CdpProtHelloEntry, child)
        return &protocolHelloList.CdpProtHelloEntry[len(protocolHelloList.CdpProtHelloEntry)-1]
    }
    return nil
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocolHelloList.CdpProtHelloEntry {
        children[protocolHelloList.CdpProtHelloEntry[i].GetSegmentPath()] = &protocolHelloList.CdpProtHelloEntry[i]
    }
    return children
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetBundleName() string { return "cisco_ios_xr" }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetYangName() string { return "protocol-hello-list" }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) SetParent(parent types.Entity) { protocolHelloList.parent = parent }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetParent() types.Entity { return protocolHelloList.parent }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol Hello msg. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetFilter() yfilter.YFilter { return cdpProtHelloEntry.YFilter }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) SetFilter(yf yfilter.YFilter) { cdpProtHelloEntry.YFilter = yf }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetGoName(yname string) string {
    if yname == "hello-message" { return "HelloMessage" }
    return ""
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetSegmentPath() string {
    return "cdp-prot-hello-entry"
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-message"] = cdpProtHelloEntry.HelloMessage
    return leafs
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetBundleName() string { return "cisco_ios_xr" }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetYangName() string { return "cdp-prot-hello-entry" }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) SetParent(parent types.Entity) { cdpProtHelloEntry.parent = parent }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetParent() types.Entity { return cdpProtHelloEntry.parent }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Details_Detail_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetParentYangName() string { return "protocol-hello-list" }

// Cdp_Nodes_Node_Neighbors_Devices
// The detailed CDP neighbor table
type Cdp_Nodes_Node_Neighbors_Devices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device.
    Device []Cdp_Nodes_Node_Neighbors_Devices_Device
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetFilter() yfilter.YFilter { return devices.YFilter }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) SetFilter(yf yfilter.YFilter) { devices.YFilter = yf }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetGoName(yname string) string {
    if yname == "device" { return "Device" }
    return ""
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetSegmentPath() string {
    return "devices"
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "device" {
        for _, c := range devices.Device {
            if devices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Devices_Device{}
        devices.Device = append(devices.Device, child)
        return &devices.Device[len(devices.Device)-1]
    }
    return nil
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range devices.Device {
        children[devices.Device[i].GetSegmentPath()] = &devices.Device[i]
    }
    return children
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetBundleName() string { return "cisco_ios_xr" }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetYangName() string { return "devices" }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) SetParent(parent types.Entity) { devices.parent = parent }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetParent() types.Entity { return devices.parent }

func (devices *Cdp_Nodes_Node_Neighbors_Devices) GetParentYangName() string { return "neighbors" }

// Cdp_Nodes_Node_Neighbors_Devices_Device
// Detailed information about a CDP neighbor
// entry
type Cdp_Nodes_Node_Neighbors_Devices_Device struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The neighboring device identifier. The type is
    // string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor.
    CdpNeighbor []Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetFilter() yfilter.YFilter { return device.YFilter }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) SetFilter(yf yfilter.YFilter) { device.YFilter = yf }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetGoName(yname string) string {
    if yname == "device-id" { return "DeviceId" }
    if yname == "cdp-neighbor" { return "CdpNeighbor" }
    return ""
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetSegmentPath() string {
    return "device" + "[device-id='" + fmt.Sprintf("%v", device.DeviceId) + "']"
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-neighbor" {
        for _, c := range device.CdpNeighbor {
            if device.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor{}
        device.CdpNeighbor = append(device.CdpNeighbor, child)
        return &device.CdpNeighbor[len(device.CdpNeighbor)-1]
    }
    return nil
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range device.CdpNeighbor {
        children[device.CdpNeighbor[i].GetSegmentPath()] = &device.CdpNeighbor[i]
    }
    return children
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-id"] = device.DeviceId
    return leafs
}

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetBundleName() string { return "cisco_ios_xr" }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetYangName() string { return "device" }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) SetParent(parent types.Entity) { device.parent = parent }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetParent() types.Entity { return device.parent }

func (device *Cdp_Nodes_Node_Neighbors_Devices_Device) GetParentYangName() string { return "devices" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetFilter() yfilter.YFilter { return cdpNeighbor.YFilter }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) SetFilter(yf yfilter.YFilter) { cdpNeighbor.YFilter = yf }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetGoName(yname string) string {
    if yname == "receiving-interface-name" { return "ReceivingInterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "port-id" { return "PortId" }
    if yname == "header-version" { return "HeaderVersion" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "capabilities" { return "Capabilities" }
    if yname == "platform" { return "Platform" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetSegmentPath() string {
    return "cdp-neighbor"
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &cdpNeighbor.Detail
    }
    return nil
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &cdpNeighbor.Detail
    return children
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receiving-interface-name"] = cdpNeighbor.ReceivingInterfaceName
    leafs["device-id"] = cdpNeighbor.DeviceId
    leafs["port-id"] = cdpNeighbor.PortId
    leafs["header-version"] = cdpNeighbor.HeaderVersion
    leafs["hold-time"] = cdpNeighbor.HoldTime
    leafs["capabilities"] = cdpNeighbor.Capabilities
    leafs["platform"] = cdpNeighbor.Platform
    return leafs
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetYangName() string { return "cdp-neighbor" }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) SetParent(parent types.Entity) { cdpNeighbor.parent = parent }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetParent() types.Entity { return cdpNeighbor.parent }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor) GetParentYangName() string { return "device" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail struct {
    parent types.Entity
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

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "vtp-domain" { return "VtpDomain" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "duplex" { return "Duplex" }
    if yname == "system-name" { return "SystemName" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    if yname == "protocol-hello-list" { return "ProtocolHelloList" }
    return ""
}

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &detail.NetworkAddresses
    }
    if childYangName == "protocol-hello-list" {
        return &detail.ProtocolHelloList
    }
    return nil
}

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &detail.NetworkAddresses
    children["protocol-hello-list"] = &detail.ProtocolHelloList
    return children
}

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = detail.Version
    leafs["vtp-domain"] = detail.VtpDomain
    leafs["native-vlan"] = detail.NativeVlan
    leafs["duplex"] = detail.Duplex
    leafs["system-name"] = detail.SystemName
    return leafs
}

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetYangName() string { return "detail" }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail) GetParentYangName() string { return "cdp-neighbor" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetGoName(yname string) string {
    if yname == "cdp-addr-entry" { return "CdpAddrEntry" }
    return ""
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-addr-entry" {
        for _, c := range networkAddresses.CdpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry{}
        networkAddresses.CdpAddrEntry = append(networkAddresses.CdpAddrEntry, child)
        return &networkAddresses.CdpAddrEntry[len(networkAddresses.CdpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.CdpAddrEntry {
        children[networkAddresses.CdpAddrEntry[i].GetSegmentPath()] = &networkAddresses.CdpAddrEntry[i]
    }
    return children
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetFilter() yfilter.YFilter { return cdpAddrEntry.YFilter }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) SetFilter(yf yfilter.YFilter) { cdpAddrEntry.YFilter = yf }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetSegmentPath() string {
    return "cdp-addr-entry"
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &cdpAddrEntry.Address
    }
    return nil
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &cdpAddrEntry.Address
    return children
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetYangName() string { return "cdp-addr-entry" }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) SetParent(parent types.Entity) { cdpAddrEntry.parent = parent }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetParent() types.Entity { return cdpAddrEntry.parent }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetParentYangName() string { return "network-addresses" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address struct {
    parent types.Entity
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

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetParentYangName() string { return "cdp-addr-entry" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetFilter() yfilter.YFilter { return protocolHelloList.YFilter }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) SetFilter(yf yfilter.YFilter) { protocolHelloList.YFilter = yf }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetGoName(yname string) string {
    if yname == "cdp-prot-hello-entry" { return "CdpProtHelloEntry" }
    return ""
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetSegmentPath() string {
    return "protocol-hello-list"
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-prot-hello-entry" {
        for _, c := range protocolHelloList.CdpProtHelloEntry {
            if protocolHelloList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry{}
        protocolHelloList.CdpProtHelloEntry = append(protocolHelloList.CdpProtHelloEntry, child)
        return &protocolHelloList.CdpProtHelloEntry[len(protocolHelloList.CdpProtHelloEntry)-1]
    }
    return nil
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocolHelloList.CdpProtHelloEntry {
        children[protocolHelloList.CdpProtHelloEntry[i].GetSegmentPath()] = &protocolHelloList.CdpProtHelloEntry[i]
    }
    return children
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetBundleName() string { return "cisco_ios_xr" }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetYangName() string { return "protocol-hello-list" }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) SetParent(parent types.Entity) { protocolHelloList.parent = parent }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetParent() types.Entity { return protocolHelloList.parent }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol Hello msg. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetFilter() yfilter.YFilter { return cdpProtHelloEntry.YFilter }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) SetFilter(yf yfilter.YFilter) { cdpProtHelloEntry.YFilter = yf }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetGoName(yname string) string {
    if yname == "hello-message" { return "HelloMessage" }
    return ""
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetSegmentPath() string {
    return "cdp-prot-hello-entry"
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-message"] = cdpProtHelloEntry.HelloMessage
    return leafs
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetBundleName() string { return "cisco_ios_xr" }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetYangName() string { return "cdp-prot-hello-entry" }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) SetParent(parent types.Entity) { cdpProtHelloEntry.parent = parent }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetParent() types.Entity { return cdpProtHelloEntry.parent }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Devices_Device_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetParentYangName() string { return "protocol-hello-list" }

// Cdp_Nodes_Node_Neighbors_Summaries
// The CDP neighbor summary table
type Cdp_Nodes_Node_Neighbors_Summaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about a CDP neighbor entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary.
    Summary []Cdp_Nodes_Node_Neighbors_Summaries_Summary
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetFilter() yfilter.YFilter { return summaries.YFilter }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) SetFilter(yf yfilter.YFilter) { summaries.YFilter = yf }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    return ""
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetSegmentPath() string {
    return "summaries"
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        for _, c := range summaries.Summary {
            if summaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Summaries_Summary{}
        summaries.Summary = append(summaries.Summary, child)
        return &summaries.Summary[len(summaries.Summary)-1]
    }
    return nil
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summaries.Summary {
        children[summaries.Summary[i].GetSegmentPath()] = &summaries.Summary[i]
    }
    return children
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetBundleName() string { return "cisco_ios_xr" }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetYangName() string { return "summaries" }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) SetParent(parent types.Entity) { summaries.parent = parent }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetParent() types.Entity { return summaries.parent }

func (summaries *Cdp_Nodes_Node_Neighbors_Summaries) GetParentYangName() string { return "neighbors" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary
// Brief information about a CDP neighbor entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // cdp neighbor. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor.
    CdpNeighbor []Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "cdp-neighbor" { return "CdpNeighbor" }
    return ""
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-neighbor" {
        for _, c := range summary.CdpNeighbor {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor{}
        summary.CdpNeighbor = append(summary.CdpNeighbor, child)
        return &summary.CdpNeighbor[len(summary.CdpNeighbor)-1]
    }
    return nil
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.CdpNeighbor {
        children[summary.CdpNeighbor[i].GetSegmentPath()] = &summary.CdpNeighbor[i]
    }
    return children
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = summary.InterfaceName
    leafs["device-id"] = summary.DeviceId
    return leafs
}

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetYangName() string { return "summary" }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Cdp_Nodes_Node_Neighbors_Summaries_Summary) GetParentYangName() string { return "summaries" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor
// cdp neighbor
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetFilter() yfilter.YFilter { return cdpNeighbor.YFilter }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) SetFilter(yf yfilter.YFilter) { cdpNeighbor.YFilter = yf }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetGoName(yname string) string {
    if yname == "receiving-interface-name" { return "ReceivingInterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "port-id" { return "PortId" }
    if yname == "header-version" { return "HeaderVersion" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "capabilities" { return "Capabilities" }
    if yname == "platform" { return "Platform" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetSegmentPath() string {
    return "cdp-neighbor"
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &cdpNeighbor.Detail
    }
    return nil
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &cdpNeighbor.Detail
    return children
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receiving-interface-name"] = cdpNeighbor.ReceivingInterfaceName
    leafs["device-id"] = cdpNeighbor.DeviceId
    leafs["port-id"] = cdpNeighbor.PortId
    leafs["header-version"] = cdpNeighbor.HeaderVersion
    leafs["hold-time"] = cdpNeighbor.HoldTime
    leafs["capabilities"] = cdpNeighbor.Capabilities
    leafs["platform"] = cdpNeighbor.Platform
    return leafs
}

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetYangName() string { return "cdp-neighbor" }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) SetParent(parent types.Entity) { cdpNeighbor.parent = parent }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetParent() types.Entity { return cdpNeighbor.parent }

func (cdpNeighbor *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor) GetParentYangName() string { return "summary" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail
// Detailed neighbor info
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail struct {
    parent types.Entity
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

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "vtp-domain" { return "VtpDomain" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "duplex" { return "Duplex" }
    if yname == "system-name" { return "SystemName" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    if yname == "protocol-hello-list" { return "ProtocolHelloList" }
    return ""
}

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &detail.NetworkAddresses
    }
    if childYangName == "protocol-hello-list" {
        return &detail.ProtocolHelloList
    }
    return nil
}

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &detail.NetworkAddresses
    children["protocol-hello-list"] = &detail.ProtocolHelloList
    return children
}

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = detail.Version
    leafs["vtp-domain"] = detail.VtpDomain
    leafs["native-vlan"] = detail.NativeVlan
    leafs["duplex"] = detail.Duplex
    leafs["system-name"] = detail.SystemName
    return leafs
}

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetYangName() string { return "detail" }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail) GetParentYangName() string { return "cdp-neighbor" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses
// List of network addresses 
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cdp addr entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry.
    CdpAddrEntry []Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetGoName(yname string) string {
    if yname == "cdp-addr-entry" { return "CdpAddrEntry" }
    return ""
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-addr-entry" {
        for _, c := range networkAddresses.CdpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry{}
        networkAddresses.CdpAddrEntry = append(networkAddresses.CdpAddrEntry, child)
        return &networkAddresses.CdpAddrEntry[len(networkAddresses.CdpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.CdpAddrEntry {
        children[networkAddresses.CdpAddrEntry[i].GetSegmentPath()] = &networkAddresses.CdpAddrEntry[i]
    }
    return children
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry
// cdp addr entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network layer address.
    Address Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetFilter() yfilter.YFilter { return cdpAddrEntry.YFilter }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) SetFilter(yf yfilter.YFilter) { cdpAddrEntry.YFilter = yf }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetSegmentPath() string {
    return "cdp-addr-entry"
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &cdpAddrEntry.Address
    }
    return nil
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &cdpAddrEntry.Address
    return children
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetYangName() string { return "cdp-addr-entry" }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) SetParent(parent types.Entity) { cdpAddrEntry.parent = parent }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetParent() types.Entity { return cdpAddrEntry.parent }

func (cdpAddrEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry) GetParentYangName() string { return "network-addresses" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address
// Network layer address
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address struct {
    parent types.Entity
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

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_NetworkAddresses_CdpAddrEntry_Address) GetParentYangName() string { return "cdp-addr-entry" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList
// List of protocol hello entries
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cdp prot hello entry. The type is slice of
    // Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry.
    CdpProtHelloEntry []Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetFilter() yfilter.YFilter { return protocolHelloList.YFilter }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) SetFilter(yf yfilter.YFilter) { protocolHelloList.YFilter = yf }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetGoName(yname string) string {
    if yname == "cdp-prot-hello-entry" { return "CdpProtHelloEntry" }
    return ""
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetSegmentPath() string {
    return "protocol-hello-list"
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-prot-hello-entry" {
        for _, c := range protocolHelloList.CdpProtHelloEntry {
            if protocolHelloList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry{}
        protocolHelloList.CdpProtHelloEntry = append(protocolHelloList.CdpProtHelloEntry, child)
        return &protocolHelloList.CdpProtHelloEntry[len(protocolHelloList.CdpProtHelloEntry)-1]
    }
    return nil
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocolHelloList.CdpProtHelloEntry {
        children[protocolHelloList.CdpProtHelloEntry[i].GetSegmentPath()] = &protocolHelloList.CdpProtHelloEntry[i]
    }
    return children
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetBundleName() string { return "cisco_ios_xr" }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetYangName() string { return "protocol-hello-list" }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) SetParent(parent types.Entity) { protocolHelloList.parent = parent }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetParent() types.Entity { return protocolHelloList.parent }

func (protocolHelloList *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList) GetParentYangName() string { return "detail" }

// Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry
// cdp prot hello entry
type Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol Hello msg. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HelloMessage interface{}
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetFilter() yfilter.YFilter { return cdpProtHelloEntry.YFilter }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) SetFilter(yf yfilter.YFilter) { cdpProtHelloEntry.YFilter = yf }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetGoName(yname string) string {
    if yname == "hello-message" { return "HelloMessage" }
    return ""
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetSegmentPath() string {
    return "cdp-prot-hello-entry"
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-message"] = cdpProtHelloEntry.HelloMessage
    return leafs
}

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetBundleName() string { return "cisco_ios_xr" }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetYangName() string { return "cdp-prot-hello-entry" }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) SetParent(parent types.Entity) { cdpProtHelloEntry.parent = parent }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetParent() types.Entity { return cdpProtHelloEntry.parent }

func (cdpProtHelloEntry *Cdp_Nodes_Node_Neighbors_Summaries_Summary_CdpNeighbor_Detail_ProtocolHelloList_CdpProtHelloEntry) GetParentYangName() string { return "protocol-hello-list" }

// Cdp_Nodes_Node_Statistics
// The CDP traffic statistics for this node
type Cdp_Nodes_Node_Statistics struct {
    parent types.Entity
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

func (statistics *Cdp_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Cdp_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Cdp_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "received-packets-v1" { return "ReceivedPacketsV1" }
    if yname == "received-packets-v2" { return "ReceivedPacketsV2" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "transmitted-packets-v1" { return "TransmittedPacketsV1" }
    if yname == "transmitted-packets-v2" { return "TransmittedPacketsV2" }
    if yname == "header-errors" { return "HeaderErrors" }
    if yname == "checksum-errors" { return "ChecksumErrors" }
    if yname == "encapsulation-errors" { return "EncapsulationErrors" }
    if yname == "bad-packet-errors" { return "BadPacketErrors" }
    if yname == "out-of-memory-errors" { return "OutOfMemoryErrors" }
    if yname == "truncated-packet-errors" { return "TruncatedPacketErrors" }
    if yname == "header-version-errors" { return "HeaderVersionErrors" }
    if yname == "open-file-errors" { return "OpenFileErrors" }
    return ""
}

func (statistics *Cdp_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Cdp_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Cdp_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Cdp_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["received-packets-v1"] = statistics.ReceivedPacketsV1
    leafs["received-packets-v2"] = statistics.ReceivedPacketsV2
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["transmitted-packets-v1"] = statistics.TransmittedPacketsV1
    leafs["transmitted-packets-v2"] = statistics.TransmittedPacketsV2
    leafs["header-errors"] = statistics.HeaderErrors
    leafs["checksum-errors"] = statistics.ChecksumErrors
    leafs["encapsulation-errors"] = statistics.EncapsulationErrors
    leafs["bad-packet-errors"] = statistics.BadPacketErrors
    leafs["out-of-memory-errors"] = statistics.OutOfMemoryErrors
    leafs["truncated-packet-errors"] = statistics.TruncatedPacketErrors
    leafs["header-version-errors"] = statistics.HeaderVersionErrors
    leafs["open-file-errors"] = statistics.OpenFileErrors
    return leafs
}

func (statistics *Cdp_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Cdp_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Cdp_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Cdp_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Cdp_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Cdp_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Cdp_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Cdp_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Cdp_Nodes_Node_Interfaces
// The table of interfaces on which CDP is
// running on this node
type Cdp_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for an interface on which CDP is running. The type is
    // slice of Cdp_Nodes_Node_Interfaces_Interface.
    Interface []Cdp_Nodes_Node_Interfaces_Interface
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Cdp_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cdp_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Cdp_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Cdp_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Cdp_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Cdp_Nodes_Node_Interfaces_Interface
// Operational data for an interface on which
// CDP is running
type Cdp_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // Interface basecaps state. The type is interface{} with range:
    // 0..4294967295.
    BasecapsState interface{}

    // CDP protocol state. The type is interface{} with range: 0..4294967295.
    CdpProtocolState interface{}

    // Interface encapsulation. The type is string.
    InterfaceEncaps interface{}
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Cdp_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "basecaps-state" { return "BasecapsState" }
    if yname == "cdp-protocol-state" { return "CdpProtocolState" }
    if yname == "interface-encaps" { return "InterfaceEncaps" }
    return ""
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-handle"] = self.InterfaceHandle
    leafs["basecaps-state"] = self.BasecapsState
    leafs["cdp-protocol-state"] = self.CdpProtocolState
    leafs["interface-encaps"] = self.InterfaceEncaps
    return leafs
}

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Cdp_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Cdp_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

