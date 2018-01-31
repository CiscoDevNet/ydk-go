// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-io package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-network: IPv4 network operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_io_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_io_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-io-oper ipv4-network}", reflect.TypeOf(Ipv4Network{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-io-oper:ipv4-network", reflect.TypeOf(Ipv4Network{}))
}

// RpfMode represents Interface line states
type RpfMode string

const (
    // Strict RPF
    RpfMode_strict RpfMode = "strict"

    // Loose RPF
    RpfMode_loose RpfMode = "loose"
)

// Ipv4MaOperLineState represents Interface line states
type Ipv4MaOperLineState string

const (
    // Interface state is unknown
    Ipv4MaOperLineState_unknown Ipv4MaOperLineState = "unknown"

    // Interface has been shutdown
    Ipv4MaOperLineState_shutdown Ipv4MaOperLineState = "shutdown"

    // Interface state is down
    Ipv4MaOperLineState_down Ipv4MaOperLineState = "down"

    // Interface state is up
    Ipv4MaOperLineState_up Ipv4MaOperLineState = "up"
)

// Ipv4Network
// IPv4 network operational data
type Ipv4Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific IPv4 network operational data.
    Nodes Ipv4Network_Nodes

    // IPv4 network operational interface data.
    Interfaces Ipv4Network_Interfaces
}

func (ipv4Network *Ipv4Network) GetFilter() yfilter.YFilter { return ipv4Network.YFilter }

func (ipv4Network *Ipv4Network) SetFilter(yf yfilter.YFilter) { ipv4Network.YFilter = yf }

func (ipv4Network *Ipv4Network) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "Cisco-IOS-XR-ipv4-ma-oper:interfaces" { return "Interfaces" }
    return ""
}

func (ipv4Network *Ipv4Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-io-oper:ipv4-network"
}

func (ipv4Network *Ipv4Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ipv4Network.Nodes
    }
    if childYangName == "Cisco-IOS-XR-ipv4-ma-oper:interfaces" {
        return &ipv4Network.Interfaces
    }
    return nil
}

func (ipv4Network *Ipv4Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ipv4Network.Nodes
    children["Cisco-IOS-XR-ipv4-ma-oper:interfaces"] = &ipv4Network.Interfaces
    return children
}

func (ipv4Network *Ipv4Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Network *Ipv4Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Network *Ipv4Network) GetYangName() string { return "ipv4-network" }

func (ipv4Network *Ipv4Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Network *Ipv4Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Network *Ipv4Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Network *Ipv4Network) SetParent(parent types.Entity) { ipv4Network.parent = parent }

func (ipv4Network *Ipv4Network) GetParent() types.Entity { return ipv4Network.parent }

func (ipv4Network *Ipv4Network) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-io-oper" }

// Ipv4Network_Nodes
// Node-specific IPv4 network operational data
type Ipv4Network_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 network operational data for a particular node. The type is slice of
    // Ipv4Network_Nodes_Node.
    Node []Ipv4Network_Nodes_Node
}

func (nodes *Ipv4Network_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ipv4Network_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ipv4Network_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ipv4Network_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ipv4Network_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ipv4Network_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ipv4Network_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ipv4Network_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ipv4Network_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ipv4Network_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ipv4Network_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ipv4Network_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ipv4Network_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ipv4Network_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ipv4Network_Nodes) GetParentYangName() string { return "ipv4-network" }

// Ipv4Network_Nodes_Node
// IPv4 network operational data for a particular
// node
type Ipv4Network_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // IPv4 network operational interface data.
    InterfaceData Ipv4Network_Nodes_Node_InterfaceData

    // Statistical IPv4 network operational data for a node.
    Statistics Ipv4Network_Nodes_Node_Statistics
}

func (node *Ipv4Network_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ipv4Network_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ipv4Network_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "interface-data" { return "InterfaceData" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *Ipv4Network_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ipv4Network_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-data" {
        return &node.InterfaceData
    }
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *Ipv4Network_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-data"] = &node.InterfaceData
    children["statistics"] = &node.Statistics
    return children
}

func (node *Ipv4Network_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ipv4Network_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ipv4Network_Nodes_Node) GetYangName() string { return "node" }

func (node *Ipv4Network_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ipv4Network_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ipv4Network_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ipv4Network_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ipv4Network_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ipv4Network_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ipv4Network_Nodes_Node_InterfaceData
// IPv4 network operational interface data
type Ipv4Network_Nodes_Node_InterfaceData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF specific IPv4 network operational interface data.
    Vrfs Ipv4Network_Nodes_Node_InterfaceData_Vrfs

    // Summary of IPv4 network operational interface data on a node.
    Summary Ipv4Network_Nodes_Node_InterfaceData_Summary
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetFilter() yfilter.YFilter { return interfaceData.YFilter }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) SetFilter(yf yfilter.YFilter) { interfaceData.YFilter = yf }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetSegmentPath() string {
    return "interface-data"
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &interfaceData.Vrfs
    }
    if childYangName == "summary" {
        return &interfaceData.Summary
    }
    return nil
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &interfaceData.Vrfs
    children["summary"] = &interfaceData.Summary
    return children
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetYangName() string { return "interface-data" }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) SetParent(parent types.Entity) { interfaceData.parent = parent }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetParent() types.Entity { return interfaceData.parent }

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetParentYangName() string { return "node" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs
// VRF specific IPv4 network operational
// interface data
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name of an interface belong to. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf.
    Vrf []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetParentYangName() string { return "interface-data" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf
// VRF name of an interface belong to
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Brief interface IPv4 network operational data for a node.
    Briefs Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs

    // Detail interface IPv4 network operational data for a node.
    Details Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "briefs" { return "Briefs" }
    if yname == "details" { return "Details" }
    return ""
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "briefs" {
        return &vrf.Briefs
    }
    if childYangName == "details" {
        return &vrf.Details
    }
    return nil
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["briefs"] = &vrf.Briefs
    children["details"] = &vrf.Details
    return children
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs
// Brief interface IPv4 network operational
// data for a node
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief interface IPv4 network operational data for an interface. The type is
    // slice of Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief.
    Brief []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetFilter() yfilter.YFilter { return briefs.YFilter }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) SetFilter(yf yfilter.YFilter) { briefs.YFilter = yf }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetGoName(yname string) string {
    if yname == "brief" { return "Brief" }
    return ""
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetSegmentPath() string {
    return "briefs"
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief" {
        for _, c := range briefs.Brief {
            if briefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief{}
        briefs.Brief = append(briefs.Brief, child)
        return &briefs.Brief[len(briefs.Brief)-1]
    }
    return nil
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefs.Brief {
        children[briefs.Brief[i].GetSegmentPath()] = &briefs.Brief[i]
    }
    return children
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetBundleName() string { return "cisco_ios_xr" }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetYangName() string { return "briefs" }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) SetParent(parent types.Entity) { briefs.parent = parent }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetParent() types.Entity { return briefs.parent }

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetParentYangName() string { return "vrf" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
// Brief interface IPv4 network operational
// data for an interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Primary address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF name of the interface. The type is string.
    VrfName interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "primary-address" { return "PrimaryAddress" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "line-state" { return "LineState" }
    return ""
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetSegmentPath() string {
    return "brief" + "[interface-name='" + fmt.Sprintf("%v", brief.InterfaceName) + "']"
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = brief.InterfaceName
    leafs["primary-address"] = brief.PrimaryAddress
    leafs["vrf-id"] = brief.VrfId
    leafs["vrf-name"] = brief.VrfName
    leafs["line-state"] = brief.LineState
    return leafs
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetYangName() string { return "brief" }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetParent() types.Entity { return brief.parent }

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetParentYangName() string { return "briefs" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
// Detail interface IPv4 network operational
// data for a node
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail interface IPv4 network operational data for an interface. The type
    // is slice of Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail.
    Detail []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    return ""
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        for _, c := range details.Detail {
            if details.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail{}
        details.Detail = append(details.Detail, child)
        return &details.Detail[len(details.Detail)-1]
    }
    return nil
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range details.Detail {
        children[details.Detail[i].GetSegmentPath()] = &details.Detail[i]
    }
    return children
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetYangName() string { return "details" }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetParentYangName() string { return "vrf" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
// Detail interface IPv4 network operational
// data for an interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Primary address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}

    // Prefix length of primary address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route tag associated with the primary address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}

    // IP MTU of the interface. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Are ICMP unreachables sent on the interface?. The type is bool.
    Unreachable interface{}

    // Are ICMP redirects sent on the interface?. The type is bool.
    Redirect interface{}

    // Are direct broadcasts sent on the interface?. The type is bool.
    DirectBroadcast interface{}

    // Are mask replies sent on the interface?. The type is bool.
    MaskReply interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Name of referenced interface (valid if unnumbered). The type is string.
    UnnumberedInterfaceName interface{}

    // Is Proxy ARP disabled on the interface?. The type is bool.
    ProxyArpDisabled interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // ACLs configured on the interface.
    Acl Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl

    // Multi ACLs configured on the interface.
    MultiAcl Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl

    // Helper Addresses configured on the interface.
    HelperAddress Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress

    // RPF config on the interface.
    Rpf Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa

    // Address Publish Time.
    PubUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime

    // IDB Create Time.
    IdbUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime

    // Multicast groups joined on the interface. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup.
    MulticastGroup []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup

    // Secondary addresses on the interface. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress.
    SecondaryAddress []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "primary-address" { return "PrimaryAddress" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "line-state" { return "LineState" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "mtu" { return "Mtu" }
    if yname == "unreachable" { return "Unreachable" }
    if yname == "redirect" { return "Redirect" }
    if yname == "direct-broadcast" { return "DirectBroadcast" }
    if yname == "mask-reply" { return "MaskReply" }
    if yname == "rg-id-exists" { return "RgIdExists" }
    if yname == "mlacp-active" { return "MlacpActive" }
    if yname == "unnumbered-interface-name" { return "UnnumberedInterfaceName" }
    if yname == "proxy-arp-disabled" { return "ProxyArpDisabled" }
    if yname == "flow-tag-src" { return "FlowTagSrc" }
    if yname == "flow-tag-dst" { return "FlowTagDst" }
    if yname == "acl" { return "Acl" }
    if yname == "multi-acl" { return "MultiAcl" }
    if yname == "helper-address" { return "HelperAddress" }
    if yname == "rpf" { return "Rpf" }
    if yname == "bgp-pa" { return "BgpPa" }
    if yname == "pub-utime" { return "PubUtime" }
    if yname == "idb-utime" { return "IdbUtime" }
    if yname == "caps-utime" { return "CapsUtime" }
    if yname == "fwd-en-utime" { return "FwdEnUtime" }
    if yname == "fwd-dis-utime" { return "FwdDisUtime" }
    if yname == "multicast-group" { return "MulticastGroup" }
    if yname == "secondary-address" { return "SecondaryAddress" }
    return ""
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetSegmentPath() string {
    return "detail" + "[interface-name='" + fmt.Sprintf("%v", detail.InterfaceName) + "']"
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl" {
        return &detail.Acl
    }
    if childYangName == "multi-acl" {
        return &detail.MultiAcl
    }
    if childYangName == "helper-address" {
        return &detail.HelperAddress
    }
    if childYangName == "rpf" {
        return &detail.Rpf
    }
    if childYangName == "bgp-pa" {
        return &detail.BgpPa
    }
    if childYangName == "pub-utime" {
        return &detail.PubUtime
    }
    if childYangName == "idb-utime" {
        return &detail.IdbUtime
    }
    if childYangName == "caps-utime" {
        return &detail.CapsUtime
    }
    if childYangName == "fwd-en-utime" {
        return &detail.FwdEnUtime
    }
    if childYangName == "fwd-dis-utime" {
        return &detail.FwdDisUtime
    }
    if childYangName == "multicast-group" {
        for _, c := range detail.MulticastGroup {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup{}
        detail.MulticastGroup = append(detail.MulticastGroup, child)
        return &detail.MulticastGroup[len(detail.MulticastGroup)-1]
    }
    if childYangName == "secondary-address" {
        for _, c := range detail.SecondaryAddress {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress{}
        detail.SecondaryAddress = append(detail.SecondaryAddress, child)
        return &detail.SecondaryAddress[len(detail.SecondaryAddress)-1]
    }
    return nil
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl"] = &detail.Acl
    children["multi-acl"] = &detail.MultiAcl
    children["helper-address"] = &detail.HelperAddress
    children["rpf"] = &detail.Rpf
    children["bgp-pa"] = &detail.BgpPa
    children["pub-utime"] = &detail.PubUtime
    children["idb-utime"] = &detail.IdbUtime
    children["caps-utime"] = &detail.CapsUtime
    children["fwd-en-utime"] = &detail.FwdEnUtime
    children["fwd-dis-utime"] = &detail.FwdDisUtime
    for i := range detail.MulticastGroup {
        children[detail.MulticastGroup[i].GetSegmentPath()] = &detail.MulticastGroup[i]
    }
    for i := range detail.SecondaryAddress {
        children[detail.SecondaryAddress[i].GetSegmentPath()] = &detail.SecondaryAddress[i]
    }
    return children
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = detail.InterfaceName
    leafs["primary-address"] = detail.PrimaryAddress
    leafs["vrf-id"] = detail.VrfId
    leafs["line-state"] = detail.LineState
    leafs["prefix-length"] = detail.PrefixLength
    leafs["route-tag"] = detail.RouteTag
    leafs["mtu"] = detail.Mtu
    leafs["unreachable"] = detail.Unreachable
    leafs["redirect"] = detail.Redirect
    leafs["direct-broadcast"] = detail.DirectBroadcast
    leafs["mask-reply"] = detail.MaskReply
    leafs["rg-id-exists"] = detail.RgIdExists
    leafs["mlacp-active"] = detail.MlacpActive
    leafs["unnumbered-interface-name"] = detail.UnnumberedInterfaceName
    leafs["proxy-arp-disabled"] = detail.ProxyArpDisabled
    leafs["flow-tag-src"] = detail.FlowTagSrc
    leafs["flow-tag-dst"] = detail.FlowTagDst
    return leafs
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetYangName() string { return "detail" }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetParentYangName() string { return "details" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl
// ACLs configured on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    Inbound interface{}

    // ACL applied to outgoing packets. The type is string.
    Outbound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetFilter() yfilter.YFilter { return acl.YFilter }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) SetFilter(yf yfilter.YFilter) { acl.YFilter = yf }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    if yname == "common-in-bound" { return "CommonInBound" }
    if yname == "common-out-bound" { return "CommonOutBound" }
    return ""
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetSegmentPath() string {
    return "acl"
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inbound"] = acl.Inbound
    leafs["outbound"] = acl.Outbound
    leafs["common-in-bound"] = acl.CommonInBound
    leafs["common-out-bound"] = acl.CommonOutBound
    return leafs
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetBundleName() string { return "cisco_ios_xr" }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetYangName() string { return "acl" }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) SetParent(parent types.Entity) { acl.parent = parent }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetParent() types.Entity { return acl.parent }

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl
// Multi ACLs configured on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound.
    Inbound []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound

    // Outbound ACLs. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound.
    Outbound []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound

    // Common ACLs. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common.
    Common []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetFilter() yfilter.YFilter { return multiAcl.YFilter }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) SetFilter(yf yfilter.YFilter) { multiAcl.YFilter = yf }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    if yname == "common" { return "Common" }
    return ""
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetSegmentPath() string {
    return "multi-acl"
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inbound" {
        for _, c := range multiAcl.Inbound {
            if multiAcl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound{}
        multiAcl.Inbound = append(multiAcl.Inbound, child)
        return &multiAcl.Inbound[len(multiAcl.Inbound)-1]
    }
    if childYangName == "outbound" {
        for _, c := range multiAcl.Outbound {
            if multiAcl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound{}
        multiAcl.Outbound = append(multiAcl.Outbound, child)
        return &multiAcl.Outbound[len(multiAcl.Outbound)-1]
    }
    if childYangName == "common" {
        for _, c := range multiAcl.Common {
            if multiAcl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common{}
        multiAcl.Common = append(multiAcl.Common, child)
        return &multiAcl.Common[len(multiAcl.Common)-1]
    }
    return nil
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multiAcl.Inbound {
        children[multiAcl.Inbound[i].GetSegmentPath()] = &multiAcl.Inbound[i]
    }
    for i := range multiAcl.Outbound {
        children[multiAcl.Outbound[i].GetSegmentPath()] = &multiAcl.Outbound[i]
    }
    for i := range multiAcl.Common {
        children[multiAcl.Common[i].GetSegmentPath()] = &multiAcl.Common[i]
    }
    return children
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetBundleName() string { return "cisco_ios_xr" }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetYangName() string { return "multi-acl" }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) SetParent(parent types.Entity) { multiAcl.parent = parent }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetParent() types.Entity { return multiAcl.parent }

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound
// Inbound ACLs
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = inbound.Entry
    return leafs
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetYangName() string { return "inbound" }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetParentYangName() string { return "multi-acl" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound
// Outbound ACLs
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = outbound.Entry
    return leafs
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetYangName() string { return "outbound" }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetParentYangName() string { return "multi-acl" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common
// Common ACLs
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetFilter() yfilter.YFilter { return common.YFilter }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) SetFilter(yf yfilter.YFilter) { common.YFilter = yf }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetSegmentPath() string {
    return "common"
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = common.Entry
    return leafs
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetBundleName() string { return "cisco_ios_xr" }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetYangName() string { return "common" }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) SetParent(parent types.Entity) { common.parent = parent }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetParent() types.Entity { return common.parent }

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetParentYangName() string { return "multi-acl" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress
// Helper Addresses configured on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray.
    AddressArray []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetGoName(yname string) string {
    if yname == "address-array" { return "AddressArray" }
    return ""
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetSegmentPath() string {
    return "helper-address"
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-array" {
        for _, c := range helperAddress.AddressArray {
            if helperAddress.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray{}
        helperAddress.AddressArray = append(helperAddress.AddressArray, child)
        return &helperAddress.AddressArray[len(helperAddress.AddressArray)-1]
    }
    return nil
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddress.AddressArray {
        children[helperAddress.AddressArray[i].GetSegmentPath()] = &helperAddress.AddressArray[i]
    }
    return children
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray
// Helper address
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Entry interface{}
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetFilter() yfilter.YFilter { return addressArray.YFilter }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) SetFilter(yf yfilter.YFilter) { addressArray.YFilter = yf }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetSegmentPath() string {
    return "address-array"
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = addressArray.Entry
    return leafs
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetBundleName() string { return "cisco_ios_xr" }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetYangName() string { return "address-array" }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) SetParent(parent types.Entity) { addressArray.parent = parent }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetParent() types.Entity { return addressArray.parent }

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetParentYangName() string { return "helper-address" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf
// RPF config on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is RpfMode.
    Mode interface{}
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-default-route" { return "AllowDefaultRoute" }
    if yname == "allow-self-ping" { return "AllowSelfPing" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpf.Enable
    leafs["allow-default-route"] = rpf.AllowDefaultRoute
    leafs["allow-self-ping"] = rpf.AllowSelfPing
    leafs["mode"] = rpf.Mode
    return leafs
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetYangName() string { return "rpf" }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa
// BGP PA config on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input

    // BGP PA output config.
    Output Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetFilter() yfilter.YFilter { return bgpPa.YFilter }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) SetFilter(yf yfilter.YFilter) { bgpPa.YFilter = yf }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetSegmentPath() string {
    return "bgp-pa"
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &bgpPa.Input
    }
    if childYangName == "output" {
        return &bgpPa.Output
    }
    return nil
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &bgpPa.Input
    children["output"] = &bgpPa.Output
    return children
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetBundleName() string { return "cisco_ios_xr" }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetYangName() string { return "bgp-pa" }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) SetParent(parent types.Entity) { bgpPa.parent = parent }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetParent() types.Entity { return bgpPa.parent }

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input
// BGP PA input config
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetSegmentPath() string {
    return "input"
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = input.Enable
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetYangName() string { return "input" }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetParent() types.Entity { return input.parent }

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetParentYangName() string { return "bgp-pa" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
// BGP PA output config
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetSegmentPath() string {
    return "output"
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = output.Enable
    leafs["source"] = output.Source
    leafs["destination"] = output.Destination
    return leafs
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetYangName() string { return "output" }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetParent() types.Entity { return output.parent }

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetParentYangName() string { return "bgp-pa" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime
// Address Publish Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetFilter() yfilter.YFilter { return pubUtime.YFilter }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) SetFilter(yf yfilter.YFilter) { pubUtime.YFilter = yf }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetGoName(yname string) string {
    return ""
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetSegmentPath() string {
    return "pub-utime"
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetBundleName() string { return "cisco_ios_xr" }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetYangName() string { return "pub-utime" }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) SetParent(parent types.Entity) { pubUtime.parent = parent }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetParent() types.Entity { return pubUtime.parent }

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime
// IDB Create Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetFilter() yfilter.YFilter { return idbUtime.YFilter }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) SetFilter(yf yfilter.YFilter) { idbUtime.YFilter = yf }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetGoName(yname string) string {
    return ""
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetSegmentPath() string {
    return "idb-utime"
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetBundleName() string { return "cisco_ios_xr" }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetYangName() string { return "idb-utime" }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) SetParent(parent types.Entity) { idbUtime.parent = parent }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetParent() types.Entity { return idbUtime.parent }

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime
// CAPS Add Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetFilter() yfilter.YFilter { return capsUtime.YFilter }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) SetFilter(yf yfilter.YFilter) { capsUtime.YFilter = yf }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetGoName(yname string) string {
    return ""
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetSegmentPath() string {
    return "caps-utime"
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetBundleName() string { return "cisco_ios_xr" }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetYangName() string { return "caps-utime" }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) SetParent(parent types.Entity) { capsUtime.parent = parent }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetParent() types.Entity { return capsUtime.parent }

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime
// FWD ENABLE Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetFilter() yfilter.YFilter { return fwdEnUtime.YFilter }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) SetFilter(yf yfilter.YFilter) { fwdEnUtime.YFilter = yf }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetGoName(yname string) string {
    return ""
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetSegmentPath() string {
    return "fwd-en-utime"
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetYangName() string { return "fwd-en-utime" }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) SetParent(parent types.Entity) { fwdEnUtime.parent = parent }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetParent() types.Entity { return fwdEnUtime.parent }

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime
// FWD DISABLE Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetFilter() yfilter.YFilter { return fwdDisUtime.YFilter }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) SetFilter(yf yfilter.YFilter) { fwdDisUtime.YFilter = yf }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetGoName(yname string) string {
    return ""
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetSegmentPath() string {
    return "fwd-dis-utime"
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetYangName() string { return "fwd-dis-utime" }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) SetParent(parent types.Entity) { fwdDisUtime.parent = parent }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetParent() types.Entity { return fwdDisUtime.parent }

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup
// Multicast groups joined on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of multicast group. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetFilter() yfilter.YFilter { return multicastGroup.YFilter }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) SetFilter(yf yfilter.YFilter) { multicastGroup.YFilter = yf }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    return ""
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetSegmentPath() string {
    return "multicast-group"
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = multicastGroup.GroupAddress
    return leafs
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetBundleName() string { return "cisco_ios_xr" }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetYangName() string { return "multicast-group" }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) SetParent(parent types.Entity) { multicastGroup.parent = parent }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetParent() types.Entity { return multicastGroup.parent }

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress
// Secondary addresses on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix length of address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route Tag associated with this address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetFilter() yfilter.YFilter { return secondaryAddress.YFilter }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) SetFilter(yf yfilter.YFilter) { secondaryAddress.YFilter = yf }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetSegmentPath() string {
    return "secondary-address"
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = secondaryAddress.Address
    leafs["prefix-length"] = secondaryAddress.PrefixLength
    leafs["route-tag"] = secondaryAddress.RouteTag
    return leafs
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetYangName() string { return "secondary-address" }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) SetParent(parent types.Entity) { secondaryAddress.parent = parent }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetParent() types.Entity { return secondaryAddress.parent }

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetParentYangName() string { return "detail" }

// Ipv4Network_Nodes_Node_InterfaceData_Summary
// Summary of IPv4 network operational interface
// data on a node
type Ipv4Network_Nodes_Node_InterfaceData_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces (up,down) with basecaps up. The type is interface{}
    // with range: 0..4294967295.
    IfUpDownBasecapsUp interface{}

    // Number of interfaces (up,up).
    IfUpUp Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp

    // Number of interfaces (up,down).
    IfUpDown Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown

    // Number of interfaces (down,down).
    IfDownDown Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown

    // Number of interfaces (shutdown,down).
    IfShutdownDown Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetGoName(yname string) string {
    if yname == "if-up-down-basecaps-up" { return "IfUpDownBasecapsUp" }
    if yname == "if-up-up" { return "IfUpUp" }
    if yname == "if-up-down" { return "IfUpDown" }
    if yname == "if-down-down" { return "IfDownDown" }
    if yname == "if-shutdown-down" { return "IfShutdownDown" }
    return ""
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "if-up-up" {
        return &summary.IfUpUp
    }
    if childYangName == "if-up-down" {
        return &summary.IfUpDown
    }
    if childYangName == "if-down-down" {
        return &summary.IfDownDown
    }
    if childYangName == "if-shutdown-down" {
        return &summary.IfShutdownDown
    }
    return nil
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["if-up-up"] = &summary.IfUpUp
    children["if-up-down"] = &summary.IfUpDown
    children["if-down-down"] = &summary.IfDownDown
    children["if-shutdown-down"] = &summary.IfShutdownDown
    return children
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-up-down-basecaps-up"] = summary.IfUpDownBasecapsUp
    return leafs
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetYangName() string { return "summary" }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetParentYangName() string { return "interface-data" }

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp
// Number of interfaces (up,up)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetFilter() yfilter.YFilter { return ifUpUp.YFilter }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) SetFilter(yf yfilter.YFilter) { ifUpUp.YFilter = yf }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetSegmentPath() string {
    return "if-up-up"
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifUpUp.IpAssigned
    leafs["ip-unnumbered"] = ifUpUp.IpUnnumbered
    leafs["ip-unassigned"] = ifUpUp.IpUnassigned
    return leafs
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetBundleName() string { return "cisco_ios_xr" }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetYangName() string { return "if-up-up" }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) SetParent(parent types.Entity) { ifUpUp.parent = parent }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetParent() types.Entity { return ifUpUp.parent }

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetParentYangName() string { return "summary" }

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown
// Number of interfaces (up,down)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetFilter() yfilter.YFilter { return ifUpDown.YFilter }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) SetFilter(yf yfilter.YFilter) { ifUpDown.YFilter = yf }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetSegmentPath() string {
    return "if-up-down"
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifUpDown.IpAssigned
    leafs["ip-unnumbered"] = ifUpDown.IpUnnumbered
    leafs["ip-unassigned"] = ifUpDown.IpUnassigned
    return leafs
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetBundleName() string { return "cisco_ios_xr" }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetYangName() string { return "if-up-down" }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) SetParent(parent types.Entity) { ifUpDown.parent = parent }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetParent() types.Entity { return ifUpDown.parent }

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetParentYangName() string { return "summary" }

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown
// Number of interfaces (down,down)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetFilter() yfilter.YFilter { return ifDownDown.YFilter }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) SetFilter(yf yfilter.YFilter) { ifDownDown.YFilter = yf }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetSegmentPath() string {
    return "if-down-down"
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifDownDown.IpAssigned
    leafs["ip-unnumbered"] = ifDownDown.IpUnnumbered
    leafs["ip-unassigned"] = ifDownDown.IpUnassigned
    return leafs
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetBundleName() string { return "cisco_ios_xr" }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetYangName() string { return "if-down-down" }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) SetParent(parent types.Entity) { ifDownDown.parent = parent }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetParent() types.Entity { return ifDownDown.parent }

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetParentYangName() string { return "summary" }

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
// Number of interfaces (shutdown,down)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetFilter() yfilter.YFilter { return ifShutdownDown.YFilter }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) SetFilter(yf yfilter.YFilter) { ifShutdownDown.YFilter = yf }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetSegmentPath() string {
    return "if-shutdown-down"
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifShutdownDown.IpAssigned
    leafs["ip-unnumbered"] = ifShutdownDown.IpUnnumbered
    leafs["ip-unassigned"] = ifShutdownDown.IpUnassigned
    return leafs
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetBundleName() string { return "cisco_ios_xr" }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetYangName() string { return "if-shutdown-down" }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) SetParent(parent types.Entity) { ifShutdownDown.parent = parent }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetParent() types.Entity { return ifShutdownDown.parent }

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetParentYangName() string { return "summary" }

// Ipv4Network_Nodes_Node_Statistics
// Statistical IPv4 network operational data for
// a node
type Ipv4Network_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Traffic statistics for a node.
    Traffic Ipv4Network_Nodes_Node_Statistics_Traffic
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Network_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "traffic" { return "Traffic" }
    return ""
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &statistics.Traffic
    }
    return nil
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &statistics.Traffic
    return children
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Network_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Ipv4Network_Nodes_Node_Statistics_Traffic
// Traffic statistics for a node
type Ipv4Network_Nodes_Node_Statistics_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Network Stats.
    Ipv4Stats Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats

    // ICMP Stats.
    IcmpStats Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetGoName(yname string) string {
    if yname == "ipv4-stats" { return "Ipv4Stats" }
    if yname == "icmp-stats" { return "IcmpStats" }
    return ""
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-stats" {
        return &traffic.Ipv4Stats
    }
    if childYangName == "icmp-stats" {
        return &traffic.IcmpStats
    }
    return nil
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-stats"] = &traffic.Ipv4Stats
    children["icmp-stats"] = &traffic.IcmpStats
    return children
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetYangName() string { return "traffic" }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetParentYangName() string { return "statistics" }

// Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats
// IPv4 Network Stats
type Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Input Packets. The type is interface{} with range: 0..4294967295.
    InputPackets interface{}

    // Received Packets. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Format Errors. The type is interface{} with range: 0..4294967295.
    FormatErrors interface{}

    // Bad Hop Count. The type is interface{} with range: 0..4294967295.
    BadHopCount interface{}

    // Bad Source Address. The type is interface{} with range: 0..4294967295.
    BadSourceAddress interface{}

    // Bad Header. The type is interface{} with range: 0..4294967295.
    BadHeader interface{}

    // No Protocol. The type is interface{} with range: 0..4294967295.
    NoProtocol interface{}

    // No Gateway. The type is interface{} with range: 0..4294967295.
    NoGateway interface{}

    // RaInput. The type is interface{} with range: 0..4294967295.
    ReassembleInput interface{}

    // Reassembled. The type is interface{} with range: 0..4294967295.
    Reassembled interface{}

    // Reassembly Timeout. The type is interface{} with range: 0..4294967295.
    ReassembleTimeout interface{}

    // Reassembly Max Drop. The type is interface{} with range: 0..4294967295.
    ReassembleMaxDrop interface{}

    // Reassembly Failed. The type is interface{} with range: 0..4294967295.
    ReassembleFailed interface{}

    // IP Options Present. The type is interface{} with range: 0..4294967295.
    OptionsPresent interface{}

    // Bad Option. The type is interface{} with range: 0..4294967295.
    BadOption interface{}

    // Unknown Option. The type is interface{} with range: 0..4294967295.
    UnknownOption interface{}

    // Bad Security Option. The type is interface{} with range: 0..4294967295.
    BadSecurityOption interface{}

    // Basic Security Option. The type is interface{} with range: 0..4294967295.
    BasicSecurityOption interface{}

    // Extended Security Option. The type is interface{} with range:
    // 0..4294967295.
    ExtendedSecurityOption interface{}

    // Cipso Option. The type is interface{} with range: 0..4294967295.
    CipsoOption interface{}

    // Strict Source Route Option. The type is interface{} with range:
    // 0..4294967295.
    StrictSourceRouteOption interface{}

    // Loose Source Route Option. The type is interface{} with range:
    // 0..4294967295.
    LooseSourceRouteOption interface{}

    // Record Route Option. The type is interface{} with range: 0..4294967295.
    RecordRouteOption interface{}

    // SID Option. The type is interface{} with range: 0..4294967295.
    SidOption interface{}

    // Timestamp Option. The type is interface{} with range: 0..4294967295.
    TimestampOption interface{}

    // Router Alert Option. The type is interface{} with range: 0..4294967295.
    RouterAlertOption interface{}

    // Noop Option. The type is interface{} with range: 0..4294967295.
    NoopOption interface{}

    // End Option. The type is interface{} with range: 0..4294967295.
    EndOption interface{}

    // Packets Output. The type is interface{} with range: 0..4294967295.
    PacketsOutput interface{}

    // Packets Forwarded. The type is interface{} with range: 0..4294967295.
    PacketsForwarded interface{}

    // Packets Fragmented. The type is interface{} with range: 0..4294967295.
    PacketsFragmented interface{}

    // Fragment Count. The type is interface{} with range: 0..4294967295.
    FragmentCount interface{}

    // Encapsulation Failed. The type is interface{} with range: 0..4294967295.
    EncapsulationFailed interface{}

    // No Router. The type is interface{} with range: 0..4294967295.
    NoRouter interface{}

    // Packet Too Big. The type is interface{} with range: 0..4294967295.
    PacketTooBig interface{}

    // Multicast In. The type is interface{} with range: 0..4294967295.
    MulticastIn interface{}

    // Multicast Out. The type is interface{} with range: 0..4294967295.
    MulticastOut interface{}

    // Broadcast In. The type is interface{} with range: 0..4294967295.
    BroadcastIn interface{}

    // Broadcast Out. The type is interface{} with range: 0..4294967295.
    BroadcastOut interface{}

    // Lisp IPv4 encapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV4Encap interface{}

    // Lisp IPv4 decapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV4Decap interface{}

    // Lisp IPv6 encapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV6Encap interface{}

    // Lisp IPv6 decapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV6Decap interface{}

    // Lisp encap errors. The type is interface{} with range: 0..4294967295.
    LispEncapError interface{}

    // Lisp decap errors. The type is interface{} with range: 0..4294967295.
    LispDecapError interface{}
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetFilter() yfilter.YFilter { return ipv4Stats.YFilter }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) SetFilter(yf yfilter.YFilter) { ipv4Stats.YFilter = yf }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetGoName(yname string) string {
    if yname == "input-packets" { return "InputPackets" }
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "format-errors" { return "FormatErrors" }
    if yname == "bad-hop-count" { return "BadHopCount" }
    if yname == "bad-source-address" { return "BadSourceAddress" }
    if yname == "bad-header" { return "BadHeader" }
    if yname == "no-protocol" { return "NoProtocol" }
    if yname == "no-gateway" { return "NoGateway" }
    if yname == "reassemble-input" { return "ReassembleInput" }
    if yname == "reassembled" { return "Reassembled" }
    if yname == "reassemble-timeout" { return "ReassembleTimeout" }
    if yname == "reassemble-max-drop" { return "ReassembleMaxDrop" }
    if yname == "reassemble-failed" { return "ReassembleFailed" }
    if yname == "options-present" { return "OptionsPresent" }
    if yname == "bad-option" { return "BadOption" }
    if yname == "unknown-option" { return "UnknownOption" }
    if yname == "bad-security-option" { return "BadSecurityOption" }
    if yname == "basic-security-option" { return "BasicSecurityOption" }
    if yname == "extended-security-option" { return "ExtendedSecurityOption" }
    if yname == "cipso-option" { return "CipsoOption" }
    if yname == "strict-source-route-option" { return "StrictSourceRouteOption" }
    if yname == "loose-source-route-option" { return "LooseSourceRouteOption" }
    if yname == "record-route-option" { return "RecordRouteOption" }
    if yname == "sid-option" { return "SidOption" }
    if yname == "timestamp-option" { return "TimestampOption" }
    if yname == "router-alert-option" { return "RouterAlertOption" }
    if yname == "noop-option" { return "NoopOption" }
    if yname == "end-option" { return "EndOption" }
    if yname == "packets-output" { return "PacketsOutput" }
    if yname == "packets-forwarded" { return "PacketsForwarded" }
    if yname == "packets-fragmented" { return "PacketsFragmented" }
    if yname == "fragment-count" { return "FragmentCount" }
    if yname == "encapsulation-failed" { return "EncapsulationFailed" }
    if yname == "no-router" { return "NoRouter" }
    if yname == "packet-too-big" { return "PacketTooBig" }
    if yname == "multicast-in" { return "MulticastIn" }
    if yname == "multicast-out" { return "MulticastOut" }
    if yname == "broadcast-in" { return "BroadcastIn" }
    if yname == "broadcast-out" { return "BroadcastOut" }
    if yname == "lisp-v4-encap" { return "LispV4Encap" }
    if yname == "lisp-v4-decap" { return "LispV4Decap" }
    if yname == "lisp-v6-encap" { return "LispV6Encap" }
    if yname == "lisp-v6-decap" { return "LispV6Decap" }
    if yname == "lisp-encap-error" { return "LispEncapError" }
    if yname == "lisp-decap-error" { return "LispDecapError" }
    return ""
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetSegmentPath() string {
    return "ipv4-stats"
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input-packets"] = ipv4Stats.InputPackets
    leafs["received-packets"] = ipv4Stats.ReceivedPackets
    leafs["format-errors"] = ipv4Stats.FormatErrors
    leafs["bad-hop-count"] = ipv4Stats.BadHopCount
    leafs["bad-source-address"] = ipv4Stats.BadSourceAddress
    leafs["bad-header"] = ipv4Stats.BadHeader
    leafs["no-protocol"] = ipv4Stats.NoProtocol
    leafs["no-gateway"] = ipv4Stats.NoGateway
    leafs["reassemble-input"] = ipv4Stats.ReassembleInput
    leafs["reassembled"] = ipv4Stats.Reassembled
    leafs["reassemble-timeout"] = ipv4Stats.ReassembleTimeout
    leafs["reassemble-max-drop"] = ipv4Stats.ReassembleMaxDrop
    leafs["reassemble-failed"] = ipv4Stats.ReassembleFailed
    leafs["options-present"] = ipv4Stats.OptionsPresent
    leafs["bad-option"] = ipv4Stats.BadOption
    leafs["unknown-option"] = ipv4Stats.UnknownOption
    leafs["bad-security-option"] = ipv4Stats.BadSecurityOption
    leafs["basic-security-option"] = ipv4Stats.BasicSecurityOption
    leafs["extended-security-option"] = ipv4Stats.ExtendedSecurityOption
    leafs["cipso-option"] = ipv4Stats.CipsoOption
    leafs["strict-source-route-option"] = ipv4Stats.StrictSourceRouteOption
    leafs["loose-source-route-option"] = ipv4Stats.LooseSourceRouteOption
    leafs["record-route-option"] = ipv4Stats.RecordRouteOption
    leafs["sid-option"] = ipv4Stats.SidOption
    leafs["timestamp-option"] = ipv4Stats.TimestampOption
    leafs["router-alert-option"] = ipv4Stats.RouterAlertOption
    leafs["noop-option"] = ipv4Stats.NoopOption
    leafs["end-option"] = ipv4Stats.EndOption
    leafs["packets-output"] = ipv4Stats.PacketsOutput
    leafs["packets-forwarded"] = ipv4Stats.PacketsForwarded
    leafs["packets-fragmented"] = ipv4Stats.PacketsFragmented
    leafs["fragment-count"] = ipv4Stats.FragmentCount
    leafs["encapsulation-failed"] = ipv4Stats.EncapsulationFailed
    leafs["no-router"] = ipv4Stats.NoRouter
    leafs["packet-too-big"] = ipv4Stats.PacketTooBig
    leafs["multicast-in"] = ipv4Stats.MulticastIn
    leafs["multicast-out"] = ipv4Stats.MulticastOut
    leafs["broadcast-in"] = ipv4Stats.BroadcastIn
    leafs["broadcast-out"] = ipv4Stats.BroadcastOut
    leafs["lisp-v4-encap"] = ipv4Stats.LispV4Encap
    leafs["lisp-v4-decap"] = ipv4Stats.LispV4Decap
    leafs["lisp-v6-encap"] = ipv4Stats.LispV6Encap
    leafs["lisp-v6-decap"] = ipv4Stats.LispV6Decap
    leafs["lisp-encap-error"] = ipv4Stats.LispEncapError
    leafs["lisp-decap-error"] = ipv4Stats.LispDecapError
    return leafs
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetYangName() string { return "ipv4-stats" }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) SetParent(parent types.Entity) { ipv4Stats.parent = parent }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetParent() types.Entity { return ipv4Stats.parent }

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetParentYangName() string { return "traffic" }

// Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats
// ICMP Stats
type Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICMP Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // ICMP Checksum Errors. The type is interface{} with range: 0..4294967295.
    ChecksumError interface{}

    // ICMP Unknown. The type is interface{} with range: 0..4294967295.
    Unknown interface{}

    // ICMP Transmitted. The type is interface{} with range: 0..4294967295.
    Output interface{}

    // ICMP Admin Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    AdminUnreachableSent interface{}

    // ICMP Network Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    NetworkUnreachableSent interface{}

    // ICMP Host Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    HostUnreachableSent interface{}

    // ICMP Protocol Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    ProtocolUnreachableSent interface{}

    // ICMP Port Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    PortUnreachableSent interface{}

    // ICMP Fragment Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    FragmentUnreachableSent interface{}

    // ICMP Admin Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    AdminUnreachableReceived interface{}

    // ICMP Network Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    NetworkUnreachableReceived interface{}

    // ICMP Host Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    HostUnreachableReceived interface{}

    // ICMP Protocol Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ProtocolUnreachableReceived interface{}

    // ICMP Port Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    PortUnreachableReceived interface{}

    // ICMP Fragment Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    FragmentUnreachableReceived interface{}

    // ICMP Hopcount Sent. The type is interface{} with range: 0..4294967295.
    HopcountSent interface{}

    // ICMP Reassembly Sent. The type is interface{} with range: 0..4294967295.
    ReassemblySent interface{}

    // ICMP Hopcount Received. The type is interface{} with range: 0..4294967295.
    HopcountReceived interface{}

    // ICMP Reassembly Received. The type is interface{} with range:
    // 0..4294967295.
    ReasseblyReceived interface{}

    // ICMP Parameter Error Received. The type is interface{} with range:
    // 0..4294967295.
    ParamErrorReceived interface{}

    // ICMP Parameter Error Sent. The type is interface{} with range:
    // 0..4294967295.
    ParamErrorSend interface{}

    // ICMP Echo Request Sent. The type is interface{} with range: 0..4294967295.
    EchoRequestSent interface{}

    // ICMP Echo Request Sent. The type is interface{} with range: 0..4294967295.
    EchoRequestReceived interface{}

    // ICMP Echo Reply Sent. The type is interface{} with range: 0..4294967295.
    EchoReplySent interface{}

    // ICMP Echo Reply Received. The type is interface{} with range:
    // 0..4294967295.
    EchoReplyReceived interface{}

    // ICMP Mask Sent. The type is interface{} with range: 0..4294967295.
    MaskRequestSent interface{}

    // ICMP Mask Received. The type is interface{} with range: 0..4294967295.
    MaskRequestReceived interface{}

    // ICMP Mask Sent. The type is interface{} with range: 0..4294967295.
    MaskReplySent interface{}

    // ICMP Mask Received. The type is interface{} with range: 0..4294967295.
    MaskReplyReceived interface{}

    // ICMP Source Quench. The type is interface{} with range: 0..4294967295.
    SourceQuenchReceived interface{}

    // ICMP Redirect Received. The type is interface{} with range: 0..4294967295.
    RedirectReceived interface{}

    // ICMP Redirect Sent. The type is interface{} with range: 0..4294967295.
    RedirectSend interface{}

    // ICMP Timestamp Received. The type is interface{} with range: 0..4294967295.
    TimestampReceived interface{}

    // ICMP Timestamp Reply Received. The type is interface{} with range:
    // 0..4294967295.
    TimestampReplyReceived interface{}

    // ICMP Router Advertisement Received. The type is interface{} with range:
    // 0..4294967295.
    RouterAdvertReceived interface{}

    // ICMP Router Solicited Received. The type is interface{} with range:
    // 0..4294967295.
    RouterSolicitReceived interface{}
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetFilter() yfilter.YFilter { return icmpStats.YFilter }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) SetFilter(yf yfilter.YFilter) { icmpStats.YFilter = yf }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetGoName(yname string) string {
    if yname == "received" { return "Received" }
    if yname == "checksum-error" { return "ChecksumError" }
    if yname == "unknown" { return "Unknown" }
    if yname == "output" { return "Output" }
    if yname == "admin-unreachable-sent" { return "AdminUnreachableSent" }
    if yname == "network-unreachable-sent" { return "NetworkUnreachableSent" }
    if yname == "host-unreachable-sent" { return "HostUnreachableSent" }
    if yname == "protocol-unreachable-sent" { return "ProtocolUnreachableSent" }
    if yname == "port-unreachable-sent" { return "PortUnreachableSent" }
    if yname == "fragment-unreachable-sent" { return "FragmentUnreachableSent" }
    if yname == "admin-unreachable-received" { return "AdminUnreachableReceived" }
    if yname == "network-unreachable-received" { return "NetworkUnreachableReceived" }
    if yname == "host-unreachable-received" { return "HostUnreachableReceived" }
    if yname == "protocol-unreachable-received" { return "ProtocolUnreachableReceived" }
    if yname == "port-unreachable-received" { return "PortUnreachableReceived" }
    if yname == "fragment-unreachable-received" { return "FragmentUnreachableReceived" }
    if yname == "hopcount-sent" { return "HopcountSent" }
    if yname == "reassembly-sent" { return "ReassemblySent" }
    if yname == "hopcount-received" { return "HopcountReceived" }
    if yname == "reassebly-received" { return "ReasseblyReceived" }
    if yname == "param-error-received" { return "ParamErrorReceived" }
    if yname == "param-error-send" { return "ParamErrorSend" }
    if yname == "echo-request-sent" { return "EchoRequestSent" }
    if yname == "echo-request-received" { return "EchoRequestReceived" }
    if yname == "echo-reply-sent" { return "EchoReplySent" }
    if yname == "echo-reply-received" { return "EchoReplyReceived" }
    if yname == "mask-request-sent" { return "MaskRequestSent" }
    if yname == "mask-request-received" { return "MaskRequestReceived" }
    if yname == "mask-reply-sent" { return "MaskReplySent" }
    if yname == "mask-reply-received" { return "MaskReplyReceived" }
    if yname == "source-quench-received" { return "SourceQuenchReceived" }
    if yname == "redirect-received" { return "RedirectReceived" }
    if yname == "redirect-send" { return "RedirectSend" }
    if yname == "timestamp-received" { return "TimestampReceived" }
    if yname == "timestamp-reply-received" { return "TimestampReplyReceived" }
    if yname == "router-advert-received" { return "RouterAdvertReceived" }
    if yname == "router-solicit-received" { return "RouterSolicitReceived" }
    return ""
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetSegmentPath() string {
    return "icmp-stats"
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received"] = icmpStats.Received
    leafs["checksum-error"] = icmpStats.ChecksumError
    leafs["unknown"] = icmpStats.Unknown
    leafs["output"] = icmpStats.Output
    leafs["admin-unreachable-sent"] = icmpStats.AdminUnreachableSent
    leafs["network-unreachable-sent"] = icmpStats.NetworkUnreachableSent
    leafs["host-unreachable-sent"] = icmpStats.HostUnreachableSent
    leafs["protocol-unreachable-sent"] = icmpStats.ProtocolUnreachableSent
    leafs["port-unreachable-sent"] = icmpStats.PortUnreachableSent
    leafs["fragment-unreachable-sent"] = icmpStats.FragmentUnreachableSent
    leafs["admin-unreachable-received"] = icmpStats.AdminUnreachableReceived
    leafs["network-unreachable-received"] = icmpStats.NetworkUnreachableReceived
    leafs["host-unreachable-received"] = icmpStats.HostUnreachableReceived
    leafs["protocol-unreachable-received"] = icmpStats.ProtocolUnreachableReceived
    leafs["port-unreachable-received"] = icmpStats.PortUnreachableReceived
    leafs["fragment-unreachable-received"] = icmpStats.FragmentUnreachableReceived
    leafs["hopcount-sent"] = icmpStats.HopcountSent
    leafs["reassembly-sent"] = icmpStats.ReassemblySent
    leafs["hopcount-received"] = icmpStats.HopcountReceived
    leafs["reassebly-received"] = icmpStats.ReasseblyReceived
    leafs["param-error-received"] = icmpStats.ParamErrorReceived
    leafs["param-error-send"] = icmpStats.ParamErrorSend
    leafs["echo-request-sent"] = icmpStats.EchoRequestSent
    leafs["echo-request-received"] = icmpStats.EchoRequestReceived
    leafs["echo-reply-sent"] = icmpStats.EchoReplySent
    leafs["echo-reply-received"] = icmpStats.EchoReplyReceived
    leafs["mask-request-sent"] = icmpStats.MaskRequestSent
    leafs["mask-request-received"] = icmpStats.MaskRequestReceived
    leafs["mask-reply-sent"] = icmpStats.MaskReplySent
    leafs["mask-reply-received"] = icmpStats.MaskReplyReceived
    leafs["source-quench-received"] = icmpStats.SourceQuenchReceived
    leafs["redirect-received"] = icmpStats.RedirectReceived
    leafs["redirect-send"] = icmpStats.RedirectSend
    leafs["timestamp-received"] = icmpStats.TimestampReceived
    leafs["timestamp-reply-received"] = icmpStats.TimestampReplyReceived
    leafs["router-advert-received"] = icmpStats.RouterAdvertReceived
    leafs["router-solicit-received"] = icmpStats.RouterSolicitReceived
    return leafs
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetBundleName() string { return "cisco_ios_xr" }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetYangName() string { return "icmp-stats" }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) SetParent(parent types.Entity) { icmpStats.parent = parent }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetParent() types.Entity { return icmpStats.parent }

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetParentYangName() string { return "traffic" }

// Ipv4Network_Interfaces
// IPv4 network operational interface data
type Ipv4Network_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface names with VRF. The type is slice of
    // Ipv4Network_Interfaces_Interface.
    Interface []Ipv4Network_Interfaces_Interface
}

func (interfaces *Ipv4Network_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Ipv4Network_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Ipv4Network_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Ipv4Network_Interfaces) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ma-oper:interfaces"
}

func (interfaces *Ipv4Network_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Ipv4Network_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Ipv4Network_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Ipv4Network_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Ipv4Network_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Ipv4Network_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Ipv4Network_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Ipv4Network_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Ipv4Network_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Ipv4Network_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Ipv4Network_Interfaces) GetParentYangName() string { return "ipv4-network" }

// Ipv4Network_Interfaces_Interface
// Interface names with VRF
type Ipv4Network_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // List of VRF on the interface.
    Vrfs Ipv4Network_Interfaces_Interface_Vrfs
}

func (self *Ipv4Network_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ipv4Network_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ipv4Network_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (self *Ipv4Network_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Ipv4Network_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &self.Vrfs
    }
    return nil
}

func (self *Ipv4Network_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &self.Vrfs
    return children
}

func (self *Ipv4Network_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *Ipv4Network_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ipv4Network_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Ipv4Network_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ipv4Network_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ipv4Network_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ipv4Network_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ipv4Network_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Ipv4Network_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Ipv4Network_Interfaces_Interface_Vrfs
// List of VRF on the interface
type Ipv4Network_Interfaces_Interface_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF information on the interface. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf.
    Vrf []Ipv4Network_Interfaces_Interface_Vrfs_Vrf
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Interfaces_Interface_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetParentYangName() string { return "interface" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf
// VRF information on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Detail IPv4 network operational data for an interface.
    Detail Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail

    // Brief IPv4 network operational data for an interface.
    Brief Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "detail" { return "Detail" }
    if yname == "brief" { return "Brief" }
    return ""
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &vrf.Detail
    }
    if childYangName == "brief" {
        return &vrf.Brief
    }
    return nil
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &vrf.Detail
    children["brief"] = &vrf.Brief
    return children
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail
// Detail IPv4 network operational data for an
// interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Primary address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}

    // Prefix length of primary address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route tag associated with the primary address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}

    // IP MTU of the interface. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Are ICMP unreachables sent on the interface?. The type is bool.
    Unreachable interface{}

    // Are ICMP redirects sent on the interface?. The type is bool.
    Redirect interface{}

    // Are direct broadcasts sent on the interface?. The type is bool.
    DirectBroadcast interface{}

    // Are mask replies sent on the interface?. The type is bool.
    MaskReply interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Name of referenced interface (valid if unnumbered). The type is string.
    UnnumberedInterfaceName interface{}

    // Is Proxy ARP disabled on the interface?. The type is bool.
    ProxyArpDisabled interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // ACLs configured on the interface.
    Acl Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl

    // Multi ACLs configured on the interface.
    MultiAcl Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl

    // Helper Addresses configured on the interface.
    HelperAddress Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress

    // RPF config on the interface.
    Rpf Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa

    // Address Publish Time.
    PubUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime

    // IDB Create Time.
    IdbUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime

    // Multicast groups joined on the interface. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup.
    MulticastGroup []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup

    // Secondary addresses on the interface. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress.
    SecondaryAddress []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetGoName(yname string) string {
    if yname == "primary-address" { return "PrimaryAddress" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "line-state" { return "LineState" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "mtu" { return "Mtu" }
    if yname == "unreachable" { return "Unreachable" }
    if yname == "redirect" { return "Redirect" }
    if yname == "direct-broadcast" { return "DirectBroadcast" }
    if yname == "mask-reply" { return "MaskReply" }
    if yname == "rg-id-exists" { return "RgIdExists" }
    if yname == "mlacp-active" { return "MlacpActive" }
    if yname == "unnumbered-interface-name" { return "UnnumberedInterfaceName" }
    if yname == "proxy-arp-disabled" { return "ProxyArpDisabled" }
    if yname == "flow-tag-src" { return "FlowTagSrc" }
    if yname == "flow-tag-dst" { return "FlowTagDst" }
    if yname == "acl" { return "Acl" }
    if yname == "multi-acl" { return "MultiAcl" }
    if yname == "helper-address" { return "HelperAddress" }
    if yname == "rpf" { return "Rpf" }
    if yname == "bgp-pa" { return "BgpPa" }
    if yname == "pub-utime" { return "PubUtime" }
    if yname == "idb-utime" { return "IdbUtime" }
    if yname == "caps-utime" { return "CapsUtime" }
    if yname == "fwd-en-utime" { return "FwdEnUtime" }
    if yname == "fwd-dis-utime" { return "FwdDisUtime" }
    if yname == "multicast-group" { return "MulticastGroup" }
    if yname == "secondary-address" { return "SecondaryAddress" }
    return ""
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl" {
        return &detail.Acl
    }
    if childYangName == "multi-acl" {
        return &detail.MultiAcl
    }
    if childYangName == "helper-address" {
        return &detail.HelperAddress
    }
    if childYangName == "rpf" {
        return &detail.Rpf
    }
    if childYangName == "bgp-pa" {
        return &detail.BgpPa
    }
    if childYangName == "pub-utime" {
        return &detail.PubUtime
    }
    if childYangName == "idb-utime" {
        return &detail.IdbUtime
    }
    if childYangName == "caps-utime" {
        return &detail.CapsUtime
    }
    if childYangName == "fwd-en-utime" {
        return &detail.FwdEnUtime
    }
    if childYangName == "fwd-dis-utime" {
        return &detail.FwdDisUtime
    }
    if childYangName == "multicast-group" {
        for _, c := range detail.MulticastGroup {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup{}
        detail.MulticastGroup = append(detail.MulticastGroup, child)
        return &detail.MulticastGroup[len(detail.MulticastGroup)-1]
    }
    if childYangName == "secondary-address" {
        for _, c := range detail.SecondaryAddress {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress{}
        detail.SecondaryAddress = append(detail.SecondaryAddress, child)
        return &detail.SecondaryAddress[len(detail.SecondaryAddress)-1]
    }
    return nil
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl"] = &detail.Acl
    children["multi-acl"] = &detail.MultiAcl
    children["helper-address"] = &detail.HelperAddress
    children["rpf"] = &detail.Rpf
    children["bgp-pa"] = &detail.BgpPa
    children["pub-utime"] = &detail.PubUtime
    children["idb-utime"] = &detail.IdbUtime
    children["caps-utime"] = &detail.CapsUtime
    children["fwd-en-utime"] = &detail.FwdEnUtime
    children["fwd-dis-utime"] = &detail.FwdDisUtime
    for i := range detail.MulticastGroup {
        children[detail.MulticastGroup[i].GetSegmentPath()] = &detail.MulticastGroup[i]
    }
    for i := range detail.SecondaryAddress {
        children[detail.SecondaryAddress[i].GetSegmentPath()] = &detail.SecondaryAddress[i]
    }
    return children
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["primary-address"] = detail.PrimaryAddress
    leafs["vrf-id"] = detail.VrfId
    leafs["line-state"] = detail.LineState
    leafs["prefix-length"] = detail.PrefixLength
    leafs["route-tag"] = detail.RouteTag
    leafs["mtu"] = detail.Mtu
    leafs["unreachable"] = detail.Unreachable
    leafs["redirect"] = detail.Redirect
    leafs["direct-broadcast"] = detail.DirectBroadcast
    leafs["mask-reply"] = detail.MaskReply
    leafs["rg-id-exists"] = detail.RgIdExists
    leafs["mlacp-active"] = detail.MlacpActive
    leafs["unnumbered-interface-name"] = detail.UnnumberedInterfaceName
    leafs["proxy-arp-disabled"] = detail.ProxyArpDisabled
    leafs["flow-tag-src"] = detail.FlowTagSrc
    leafs["flow-tag-dst"] = detail.FlowTagDst
    return leafs
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetYangName() string { return "detail" }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetParentYangName() string { return "vrf" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl
// ACLs configured on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    Inbound interface{}

    // ACL applied to outgoing packets. The type is string.
    Outbound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetFilter() yfilter.YFilter { return acl.YFilter }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) SetFilter(yf yfilter.YFilter) { acl.YFilter = yf }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    if yname == "common-in-bound" { return "CommonInBound" }
    if yname == "common-out-bound" { return "CommonOutBound" }
    return ""
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetSegmentPath() string {
    return "acl"
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inbound"] = acl.Inbound
    leafs["outbound"] = acl.Outbound
    leafs["common-in-bound"] = acl.CommonInBound
    leafs["common-out-bound"] = acl.CommonOutBound
    return leafs
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetBundleName() string { return "cisco_ios_xr" }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetYangName() string { return "acl" }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) SetParent(parent types.Entity) { acl.parent = parent }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetParent() types.Entity { return acl.parent }

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl
// Multi ACLs configured on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of string.
    Inbound []interface{}

    // Outbound ACLs. The type is slice of string.
    Outbound []interface{}

    // Common ACLs. The type is slice of string.
    Common []interface{}
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetFilter() yfilter.YFilter { return multiAcl.YFilter }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) SetFilter(yf yfilter.YFilter) { multiAcl.YFilter = yf }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    if yname == "common" { return "Common" }
    return ""
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetSegmentPath() string {
    return "multi-acl"
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inbound"] = multiAcl.Inbound
    leafs["outbound"] = multiAcl.Outbound
    leafs["common"] = multiAcl.Common
    return leafs
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetBundleName() string { return "cisco_ios_xr" }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetYangName() string { return "multi-acl" }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) SetParent(parent types.Entity) { multiAcl.parent = parent }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetParent() types.Entity { return multiAcl.parent }

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress
// Helper Addresses configured on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Helper address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AddressArray []interface{}
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetGoName(yname string) string {
    if yname == "address-array" { return "AddressArray" }
    return ""
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetSegmentPath() string {
    return "helper-address"
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-array"] = helperAddress.AddressArray
    return leafs
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf
// RPF config on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is RpfMode.
    Mode interface{}
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-default-route" { return "AllowDefaultRoute" }
    if yname == "allow-self-ping" { return "AllowSelfPing" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpf.Enable
    leafs["allow-default-route"] = rpf.AllowDefaultRoute
    leafs["allow-self-ping"] = rpf.AllowSelfPing
    leafs["mode"] = rpf.Mode
    return leafs
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetYangName() string { return "rpf" }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa
// BGP PA config on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input

    // BGP PA output config.
    Output Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetFilter() yfilter.YFilter { return bgpPa.YFilter }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) SetFilter(yf yfilter.YFilter) { bgpPa.YFilter = yf }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetSegmentPath() string {
    return "bgp-pa"
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &bgpPa.Input
    }
    if childYangName == "output" {
        return &bgpPa.Output
    }
    return nil
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &bgpPa.Input
    children["output"] = &bgpPa.Output
    return children
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetBundleName() string { return "cisco_ios_xr" }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetYangName() string { return "bgp-pa" }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) SetParent(parent types.Entity) { bgpPa.parent = parent }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetParent() types.Entity { return bgpPa.parent }

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input
// BGP PA input config
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetSegmentPath() string {
    return "input"
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = input.Enable
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetYangName() string { return "input" }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetParent() types.Entity { return input.parent }

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetParentYangName() string { return "bgp-pa" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output
// BGP PA output config
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetSegmentPath() string {
    return "output"
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = output.Enable
    leafs["source"] = output.Source
    leafs["destination"] = output.Destination
    return leafs
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetYangName() string { return "output" }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetParent() types.Entity { return output.parent }

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetParentYangName() string { return "bgp-pa" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime
// Address Publish Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetFilter() yfilter.YFilter { return pubUtime.YFilter }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) SetFilter(yf yfilter.YFilter) { pubUtime.YFilter = yf }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetGoName(yname string) string {
    return ""
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetSegmentPath() string {
    return "pub-utime"
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetBundleName() string { return "cisco_ios_xr" }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetYangName() string { return "pub-utime" }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) SetParent(parent types.Entity) { pubUtime.parent = parent }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetParent() types.Entity { return pubUtime.parent }

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime
// IDB Create Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetFilter() yfilter.YFilter { return idbUtime.YFilter }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) SetFilter(yf yfilter.YFilter) { idbUtime.YFilter = yf }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetGoName(yname string) string {
    return ""
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetSegmentPath() string {
    return "idb-utime"
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetBundleName() string { return "cisco_ios_xr" }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetYangName() string { return "idb-utime" }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) SetParent(parent types.Entity) { idbUtime.parent = parent }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetParent() types.Entity { return idbUtime.parent }

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime
// CAPS Add Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetFilter() yfilter.YFilter { return capsUtime.YFilter }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) SetFilter(yf yfilter.YFilter) { capsUtime.YFilter = yf }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetGoName(yname string) string {
    return ""
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetSegmentPath() string {
    return "caps-utime"
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetBundleName() string { return "cisco_ios_xr" }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetYangName() string { return "caps-utime" }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) SetParent(parent types.Entity) { capsUtime.parent = parent }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetParent() types.Entity { return capsUtime.parent }

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime
// FWD ENABLE Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetFilter() yfilter.YFilter { return fwdEnUtime.YFilter }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) SetFilter(yf yfilter.YFilter) { fwdEnUtime.YFilter = yf }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetGoName(yname string) string {
    return ""
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetSegmentPath() string {
    return "fwd-en-utime"
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetYangName() string { return "fwd-en-utime" }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) SetParent(parent types.Entity) { fwdEnUtime.parent = parent }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetParent() types.Entity { return fwdEnUtime.parent }

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime
// FWD DISABLE Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetFilter() yfilter.YFilter { return fwdDisUtime.YFilter }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) SetFilter(yf yfilter.YFilter) { fwdDisUtime.YFilter = yf }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetGoName(yname string) string {
    return ""
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetSegmentPath() string {
    return "fwd-dis-utime"
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetYangName() string { return "fwd-dis-utime" }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) SetParent(parent types.Entity) { fwdDisUtime.parent = parent }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetParent() types.Entity { return fwdDisUtime.parent }

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup
// Multicast groups joined on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of multicast group. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetFilter() yfilter.YFilter { return multicastGroup.YFilter }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) SetFilter(yf yfilter.YFilter) { multicastGroup.YFilter = yf }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    return ""
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetSegmentPath() string {
    return "multicast-group"
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = multicastGroup.GroupAddress
    return leafs
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetBundleName() string { return "cisco_ios_xr" }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetYangName() string { return "multicast-group" }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) SetParent(parent types.Entity) { multicastGroup.parent = parent }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetParent() types.Entity { return multicastGroup.parent }

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress
// Secondary addresses on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix length of address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route Tag associated with this address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetFilter() yfilter.YFilter { return secondaryAddress.YFilter }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) SetFilter(yf yfilter.YFilter) { secondaryAddress.YFilter = yf }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetSegmentPath() string {
    return "secondary-address"
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = secondaryAddress.Address
    leafs["prefix-length"] = secondaryAddress.PrefixLength
    leafs["route-tag"] = secondaryAddress.RouteTag
    return leafs
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetYangName() string { return "secondary-address" }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) SetParent(parent types.Entity) { secondaryAddress.parent = parent }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetParent() types.Entity { return secondaryAddress.parent }

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetParentYangName() string { return "detail" }

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief
// Brief IPv4 network operational data for an
// interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Primary address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF name of the interface. The type is string.
    VrfName interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetGoName(yname string) string {
    if yname == "primary-address" { return "PrimaryAddress" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "line-state" { return "LineState" }
    return ""
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetSegmentPath() string {
    return "brief"
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["primary-address"] = brief.PrimaryAddress
    leafs["vrf-id"] = brief.VrfId
    leafs["vrf-name"] = brief.VrfName
    leafs["line-state"] = brief.LineState
    return leafs
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetYangName() string { return "brief" }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetParent() types.Entity { return brief.parent }

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetParentYangName() string { return "vrf" }

