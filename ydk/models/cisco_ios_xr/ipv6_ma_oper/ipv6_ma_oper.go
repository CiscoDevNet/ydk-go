// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-ma package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-network: IPv6 network operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_ma_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ma-oper ipv6-network}", reflect.TypeOf(Ipv6Network{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ma-oper:ipv6-network", reflect.TypeOf(Ipv6Network{}))
}

// Ipv6MaOperState represents Interface oper states
type Ipv6MaOperState string

const (
    // Interface oper state is up
    Ipv6MaOperState_oper_up Ipv6MaOperState = "oper-up"

    // Interface oper state is down
    Ipv6MaOperState_oper_down Ipv6MaOperState = "oper-down"
)

// Ipv6MaIfLineState represents Interface line states
type Ipv6MaIfLineState string

const (
    // Interface state is down
    Ipv6MaIfLineState_down Ipv6MaIfLineState = "down"

    // Interface state is up
    Ipv6MaIfLineState_up Ipv6MaIfLineState = "up"

    // Interface state is unknown
    Ipv6MaIfLineState_unknown Ipv6MaIfLineState = "unknown"

    // Interface state is incorrect
    Ipv6MaIfLineState_error Ipv6MaIfLineState = "error"
)

// Ipv6MaIfAddrState represents Interface address states
type Ipv6MaIfAddrState string

const (
    // This is an active address that can appear as
    // the destination or source address of a packet
    Ipv6MaIfAddrState_active Ipv6MaIfAddrState = "active"

    // This is a valid but deprecated address that
    // should no longer be used as a source address in
    // new communications
    Ipv6MaIfAddrState_deprecated Ipv6MaIfAddrState = "deprecated"

    // This is a duplicate (invalid) address because
    // of conflict
    Ipv6MaIfAddrState_duplicate Ipv6MaIfAddrState = "duplicate"

    // This is not accessible because the interface to
    // which this address is assigned is not
    // operational
    Ipv6MaIfAddrState_inaccessible Ipv6MaIfAddrState = "inaccessible"

    // Status can not be determined for some reason
    Ipv6MaIfAddrState_tentative Ipv6MaIfAddrState = "tentative"
)

// Ipv6Network
// IPv6 network operational data
type Ipv6Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific IPv6 network operational data.
    Nodes Ipv6Network_Nodes
}

func (ipv6Network *Ipv6Network) GetFilter() yfilter.YFilter { return ipv6Network.YFilter }

func (ipv6Network *Ipv6Network) SetFilter(yf yfilter.YFilter) { ipv6Network.YFilter = yf }

func (ipv6Network *Ipv6Network) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ipv6Network *Ipv6Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ma-oper:ipv6-network"
}

func (ipv6Network *Ipv6Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ipv6Network.Nodes
    }
    return nil
}

func (ipv6Network *Ipv6Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ipv6Network.Nodes
    return children
}

func (ipv6Network *Ipv6Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Network *Ipv6Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Network *Ipv6Network) GetYangName() string { return "ipv6-network" }

func (ipv6Network *Ipv6Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Network *Ipv6Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Network *Ipv6Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Network *Ipv6Network) SetParent(parent types.Entity) { ipv6Network.parent = parent }

func (ipv6Network *Ipv6Network) GetParent() types.Entity { return ipv6Network.parent }

func (ipv6Network *Ipv6Network) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ma-oper" }

// Ipv6Network_Nodes
// Node-specific IPv6 network operational data
type Ipv6Network_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 network operational data for a particular node. The type is slice of
    // Ipv6Network_Nodes_Node.
    Node []Ipv6Network_Nodes_Node
}

func (nodes *Ipv6Network_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ipv6Network_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ipv6Network_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ipv6Network_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ipv6Network_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ipv6Network_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ipv6Network_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ipv6Network_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ipv6Network_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ipv6Network_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ipv6Network_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ipv6Network_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ipv6Network_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ipv6Network_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ipv6Network_Nodes) GetParentYangName() string { return "ipv6-network" }

// Ipv6Network_Nodes_Node
// IPv6 network operational data for a particular
// node
type Ipv6Network_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // IPv6 network operational interface data.
    InterfaceData Ipv6Network_Nodes_Node_InterfaceData
}

func (node *Ipv6Network_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ipv6Network_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ipv6Network_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "interface-data" { return "InterfaceData" }
    return ""
}

func (node *Ipv6Network_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ipv6Network_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-data" {
        return &node.InterfaceData
    }
    return nil
}

func (node *Ipv6Network_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-data"] = &node.InterfaceData
    return children
}

func (node *Ipv6Network_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ipv6Network_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ipv6Network_Nodes_Node) GetYangName() string { return "node" }

func (node *Ipv6Network_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ipv6Network_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ipv6Network_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ipv6Network_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ipv6Network_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ipv6Network_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ipv6Network_Nodes_Node_InterfaceData
// IPv6 network operational interface data
type Ipv6Network_Nodes_Node_InterfaceData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF specific IPv6 network operational interface data.
    Vrfs Ipv6Network_Nodes_Node_InterfaceData_Vrfs

    // Summary of IPv6 network operational interface data on a node.
    Summary Ipv6Network_Nodes_Node_InterfaceData_Summary
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetFilter() yfilter.YFilter { return interfaceData.YFilter }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) SetFilter(yf yfilter.YFilter) { interfaceData.YFilter = yf }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetSegmentPath() string {
    return "interface-data"
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &interfaceData.Vrfs
    }
    if childYangName == "summary" {
        return &interfaceData.Summary
    }
    return nil
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &interfaceData.Vrfs
    children["summary"] = &interfaceData.Summary
    return children
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetYangName() string { return "interface-data" }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) SetParent(parent types.Entity) { interfaceData.parent = parent }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetParent() types.Entity { return interfaceData.parent }

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetParentYangName() string { return "node" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs
// VRF specific IPv6 network operational
// interface data
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF ID of an interface belong to. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf.
    Vrf []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetParentYangName() string { return "interface-data" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf
// VRF ID of an interface belong to
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Brief interface IPv6 network operational data for a node.
    Briefs Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs

    // Detail interface IPv4 network operational data for global data.
    GlobalDetails Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails

    // Brief interface IPv6 network operational data from global data.
    GlobalBriefs Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs

    // Detail interface IPv4 network operational data for a node.
    Details Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "briefs" { return "Briefs" }
    if yname == "global-details" { return "GlobalDetails" }
    if yname == "global-briefs" { return "GlobalBriefs" }
    if yname == "details" { return "Details" }
    return ""
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "briefs" {
        return &vrf.Briefs
    }
    if childYangName == "global-details" {
        return &vrf.GlobalDetails
    }
    if childYangName == "global-briefs" {
        return &vrf.GlobalBriefs
    }
    if childYangName == "details" {
        return &vrf.Details
    }
    return nil
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["briefs"] = &vrf.Briefs
    children["global-details"] = &vrf.GlobalDetails
    children["global-briefs"] = &vrf.GlobalBriefs
    children["details"] = &vrf.Details
    return children
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs
// Brief interface IPv6 network operational
// data for a node
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief interface IPv6 network operational data for an interface. The type is
    // slice of Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief.
    Brief []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetFilter() yfilter.YFilter { return briefs.YFilter }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) SetFilter(yf yfilter.YFilter) { briefs.YFilter = yf }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetGoName(yname string) string {
    if yname == "brief" { return "Brief" }
    return ""
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetSegmentPath() string {
    return "briefs"
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief" {
        for _, c := range briefs.Brief {
            if briefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief{}
        briefs.Brief = append(briefs.Brief, child)
        return &briefs.Brief[len(briefs.Brief)-1]
    }
    return nil
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefs.Brief {
        children[briefs.Brief[i].GetSegmentPath()] = &briefs.Brief[i]
    }
    return children
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetBundleName() string { return "cisco_ios_xr" }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetYangName() string { return "briefs" }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) SetParent(parent types.Entity) { briefs.parent = parent }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetParent() types.Entity { return briefs.parent }

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetParentYangName() string { return "vrf" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
// Brief interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "line-state" { return "LineState" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "link-local-address" { return "LinkLocalAddress" }
    if yname == "address" { return "Address" }
    return ""
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetSegmentPath() string {
    return "brief" + "[interface-name='" + fmt.Sprintf("%v", brief.InterfaceName) + "']"
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-local-address" {
        return &brief.LinkLocalAddress
    }
    if childYangName == "address" {
        for _, c := range brief.Address {
            if brief.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address{}
        brief.Address = append(brief.Address, child)
        return &brief.Address[len(brief.Address)-1]
    }
    return nil
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-local-address"] = &brief.LinkLocalAddress
    for i := range brief.Address {
        children[brief.Address[i].GetSegmentPath()] = &brief.Address[i]
    }
    return children
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = brief.InterfaceName
    leafs["line-state"] = brief.LineState
    leafs["vrf-name"] = brief.VrfName
    return leafs
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetYangName() string { return "brief" }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetParent() types.Entity { return brief.parent }

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetParentYangName() string { return "briefs" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetFilter() yfilter.YFilter { return linkLocalAddress.YFilter }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) SetFilter(yf yfilter.YFilter) { linkLocalAddress.YFilter = yf }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetSegmentPath() string {
    return "link-local-address"
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = linkLocalAddress.Address
    leafs["prefix-length"] = linkLocalAddress.PrefixLength
    leafs["address-state"] = linkLocalAddress.AddressState
    leafs["is-anycast"] = linkLocalAddress.IsAnycast
    leafs["route-tag"] = linkLocalAddress.RouteTag
    return leafs
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetYangName() string { return "link-local-address" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) SetParent(parent types.Entity) { linkLocalAddress.parent = parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetParent() types.Entity { return linkLocalAddress.parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetParentYangName() string { return "brief" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = address.Address
    leafs["prefix-length"] = address.PrefixLength
    leafs["address-state"] = address.AddressState
    leafs["is-anycast"] = address.IsAnycast
    leafs["route-tag"] = address.RouteTag
    return leafs
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetYangName() string { return "address" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetParentYangName() string { return "brief" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails
// Detail interface IPv4 network operational
// data for global data
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail interface IPv6 network operational data for an interface. The type
    // is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail.
    GlobalDetail []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetFilter() yfilter.YFilter { return globalDetails.YFilter }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) SetFilter(yf yfilter.YFilter) { globalDetails.YFilter = yf }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetGoName(yname string) string {
    if yname == "global-detail" { return "GlobalDetail" }
    return ""
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetSegmentPath() string {
    return "global-details"
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-detail" {
        for _, c := range globalDetails.GlobalDetail {
            if globalDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail{}
        globalDetails.GlobalDetail = append(globalDetails.GlobalDetail, child)
        return &globalDetails.GlobalDetail[len(globalDetails.GlobalDetail)-1]
    }
    return nil
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalDetails.GlobalDetail {
        children[globalDetails.GlobalDetail[i].GetSegmentPath()] = &globalDetails.GlobalDetail[i]
    }
    return children
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetBundleName() string { return "cisco_ios_xr" }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetYangName() string { return "global-details" }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) SetParent(parent types.Entity) { globalDetails.parent = parent }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetParent() types.Entity { return globalDetails.parent }

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetParentYangName() string { return "vrf" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail
// Detail interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // IPv6 MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // IPv6 Operation State. The type is Ipv6MaOperState.
    OperationState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // ICMP unreach Enable. The type is bool.
    IsIcmpUnreachEnabled interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress

    // IPv6 Access Control List.
    AccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList

    // Multi IPv6 Access Control List.
    MultiAccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList

    // RPF config on the interface.
    Rpf Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa

    // Address Publish Time.
    Utime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime

    // IDB Create Time.
    IdbUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime

    // IPv6 Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup.
    MulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address

    // IPv6 Client Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup.
    ClientMulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetFilter() yfilter.YFilter { return globalDetail.YFilter }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) SetFilter(yf yfilter.YFilter) { globalDetail.YFilter = yf }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "line-state" { return "LineState" }
    if yname == "mtu" { return "Mtu" }
    if yname == "operation-state" { return "OperationState" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "is-icmp-unreach-enabled" { return "IsIcmpUnreachEnabled" }
    if yname == "rg-id-exists" { return "RgIdExists" }
    if yname == "mlacp-active" { return "MlacpActive" }
    if yname == "flow-tag-src" { return "FlowTagSrc" }
    if yname == "flow-tag-dst" { return "FlowTagDst" }
    if yname == "link-local-address" { return "LinkLocalAddress" }
    if yname == "access-control-list" { return "AccessControlList" }
    if yname == "multi-access-control-list" { return "MultiAccessControlList" }
    if yname == "rpf" { return "Rpf" }
    if yname == "bgp-pa" { return "BgpPa" }
    if yname == "utime" { return "Utime" }
    if yname == "idb-utime" { return "IdbUtime" }
    if yname == "caps-utime" { return "CapsUtime" }
    if yname == "fwd-en-utime" { return "FwdEnUtime" }
    if yname == "fwd-dis-utime" { return "FwdDisUtime" }
    if yname == "multicast-group" { return "MulticastGroup" }
    if yname == "address" { return "Address" }
    if yname == "client-multicast-group" { return "ClientMulticastGroup" }
    return ""
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetSegmentPath() string {
    return "global-detail" + "[interface-name='" + fmt.Sprintf("%v", globalDetail.InterfaceName) + "']"
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-local-address" {
        return &globalDetail.LinkLocalAddress
    }
    if childYangName == "access-control-list" {
        return &globalDetail.AccessControlList
    }
    if childYangName == "multi-access-control-list" {
        return &globalDetail.MultiAccessControlList
    }
    if childYangName == "rpf" {
        return &globalDetail.Rpf
    }
    if childYangName == "bgp-pa" {
        return &globalDetail.BgpPa
    }
    if childYangName == "utime" {
        return &globalDetail.Utime
    }
    if childYangName == "idb-utime" {
        return &globalDetail.IdbUtime
    }
    if childYangName == "caps-utime" {
        return &globalDetail.CapsUtime
    }
    if childYangName == "fwd-en-utime" {
        return &globalDetail.FwdEnUtime
    }
    if childYangName == "fwd-dis-utime" {
        return &globalDetail.FwdDisUtime
    }
    if childYangName == "multicast-group" {
        for _, c := range globalDetail.MulticastGroup {
            if globalDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup{}
        globalDetail.MulticastGroup = append(globalDetail.MulticastGroup, child)
        return &globalDetail.MulticastGroup[len(globalDetail.MulticastGroup)-1]
    }
    if childYangName == "address" {
        for _, c := range globalDetail.Address {
            if globalDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address{}
        globalDetail.Address = append(globalDetail.Address, child)
        return &globalDetail.Address[len(globalDetail.Address)-1]
    }
    if childYangName == "client-multicast-group" {
        for _, c := range globalDetail.ClientMulticastGroup {
            if globalDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup{}
        globalDetail.ClientMulticastGroup = append(globalDetail.ClientMulticastGroup, child)
        return &globalDetail.ClientMulticastGroup[len(globalDetail.ClientMulticastGroup)-1]
    }
    return nil
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-local-address"] = &globalDetail.LinkLocalAddress
    children["access-control-list"] = &globalDetail.AccessControlList
    children["multi-access-control-list"] = &globalDetail.MultiAccessControlList
    children["rpf"] = &globalDetail.Rpf
    children["bgp-pa"] = &globalDetail.BgpPa
    children["utime"] = &globalDetail.Utime
    children["idb-utime"] = &globalDetail.IdbUtime
    children["caps-utime"] = &globalDetail.CapsUtime
    children["fwd-en-utime"] = &globalDetail.FwdEnUtime
    children["fwd-dis-utime"] = &globalDetail.FwdDisUtime
    for i := range globalDetail.MulticastGroup {
        children[globalDetail.MulticastGroup[i].GetSegmentPath()] = &globalDetail.MulticastGroup[i]
    }
    for i := range globalDetail.Address {
        children[globalDetail.Address[i].GetSegmentPath()] = &globalDetail.Address[i]
    }
    for i := range globalDetail.ClientMulticastGroup {
        children[globalDetail.ClientMulticastGroup[i].GetSegmentPath()] = &globalDetail.ClientMulticastGroup[i]
    }
    return children
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = globalDetail.InterfaceName
    leafs["line-state"] = globalDetail.LineState
    leafs["mtu"] = globalDetail.Mtu
    leafs["operation-state"] = globalDetail.OperationState
    leafs["vrf-name"] = globalDetail.VrfName
    leafs["is-icmp-unreach-enabled"] = globalDetail.IsIcmpUnreachEnabled
    leafs["rg-id-exists"] = globalDetail.RgIdExists
    leafs["mlacp-active"] = globalDetail.MlacpActive
    leafs["flow-tag-src"] = globalDetail.FlowTagSrc
    leafs["flow-tag-dst"] = globalDetail.FlowTagDst
    return leafs
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetBundleName() string { return "cisco_ios_xr" }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetYangName() string { return "global-detail" }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) SetParent(parent types.Entity) { globalDetail.parent = parent }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetParent() types.Entity { return globalDetail.parent }

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetParentYangName() string { return "global-details" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetFilter() yfilter.YFilter { return linkLocalAddress.YFilter }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) SetFilter(yf yfilter.YFilter) { linkLocalAddress.YFilter = yf }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetSegmentPath() string {
    return "link-local-address"
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = linkLocalAddress.Address
    leafs["prefix-length"] = linkLocalAddress.PrefixLength
    leafs["address-state"] = linkLocalAddress.AddressState
    leafs["is-anycast"] = linkLocalAddress.IsAnycast
    leafs["route-tag"] = linkLocalAddress.RouteTag
    return leafs
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetYangName() string { return "link-local-address" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) SetParent(parent types.Entity) { linkLocalAddress.parent = parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetParent() types.Entity { return linkLocalAddress.parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList
// IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    InBound interface{}

    // ACL applied to outgoing packets. The type is string.
    OutBound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetFilter() yfilter.YFilter { return accessControlList.YFilter }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) SetFilter(yf yfilter.YFilter) { accessControlList.YFilter = yf }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetGoName(yname string) string {
    if yname == "in-bound" { return "InBound" }
    if yname == "out-bound" { return "OutBound" }
    if yname == "common-in-bound" { return "CommonInBound" }
    if yname == "common-out-bound" { return "CommonOutBound" }
    return ""
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetSegmentPath() string {
    return "access-control-list"
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-bound"] = accessControlList.InBound
    leafs["out-bound"] = accessControlList.OutBound
    leafs["common-in-bound"] = accessControlList.CommonInBound
    leafs["common-out-bound"] = accessControlList.CommonOutBound
    return leafs
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetBundleName() string { return "cisco_ios_xr" }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetYangName() string { return "access-control-list" }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) SetParent(parent types.Entity) { accessControlList.parent = parent }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetParent() types.Entity { return accessControlList.parent }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList
// Multi IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of string.
    Inbound []interface{}

    // Outbound ACLs. The type is slice of string.
    Outbound []interface{}

    // Common ACLs. The type is slice of string.
    Common []interface{}
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetFilter() yfilter.YFilter { return multiAccessControlList.YFilter }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) SetFilter(yf yfilter.YFilter) { multiAccessControlList.YFilter = yf }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    if yname == "common" { return "Common" }
    return ""
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetSegmentPath() string {
    return "multi-access-control-list"
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inbound"] = multiAccessControlList.Inbound
    leafs["outbound"] = multiAccessControlList.Outbound
    leafs["common"] = multiAccessControlList.Common
    return leafs
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetBundleName() string { return "cisco_ios_xr" }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetYangName() string { return "multi-access-control-list" }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) SetParent(parent types.Entity) { multiAccessControlList.parent = parent }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetParent() types.Entity { return multiAccessControlList.parent }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf
// RPF config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is interface{} with range: 0..4294967295.
    Mode interface{}
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-default-route" { return "AllowDefaultRoute" }
    if yname == "allow-self-ping" { return "AllowSelfPing" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpf.Enable
    leafs["allow-default-route"] = rpf.AllowDefaultRoute
    leafs["allow-self-ping"] = rpf.AllowSelfPing
    leafs["mode"] = rpf.Mode
    return leafs
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetYangName() string { return "rpf" }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa
// BGP PA config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input

    // BGP PA output config.
    Output Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetFilter() yfilter.YFilter { return bgpPa.YFilter }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) SetFilter(yf yfilter.YFilter) { bgpPa.YFilter = yf }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetSegmentPath() string {
    return "bgp-pa"
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &bgpPa.Input
    }
    if childYangName == "output" {
        return &bgpPa.Output
    }
    return nil
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &bgpPa.Input
    children["output"] = &bgpPa.Output
    return children
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetBundleName() string { return "cisco_ios_xr" }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetYangName() string { return "bgp-pa" }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) SetParent(parent types.Entity) { bgpPa.parent = parent }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetParent() types.Entity { return bgpPa.parent }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input
// BGP PA input config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetSegmentPath() string {
    return "input"
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = input.Enable
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetYangName() string { return "input" }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetParent() types.Entity { return input.parent }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetParentYangName() string { return "bgp-pa" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output
// BGP PA output config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetSegmentPath() string {
    return "output"
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = output.Enable
    leafs["source"] = output.Source
    leafs["destination"] = output.Destination
    return leafs
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetYangName() string { return "output" }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetParent() types.Entity { return output.parent }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetParentYangName() string { return "bgp-pa" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime
// Address Publish Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetFilter() yfilter.YFilter { return utime.YFilter }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) SetFilter(yf yfilter.YFilter) { utime.YFilter = yf }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetGoName(yname string) string {
    return ""
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetSegmentPath() string {
    return "utime"
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetBundleName() string { return "cisco_ios_xr" }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetYangName() string { return "utime" }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) SetParent(parent types.Entity) { utime.parent = parent }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetParent() types.Entity { return utime.parent }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime
// IDB Create Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetFilter() yfilter.YFilter { return idbUtime.YFilter }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) SetFilter(yf yfilter.YFilter) { idbUtime.YFilter = yf }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetGoName(yname string) string {
    return ""
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetSegmentPath() string {
    return "idb-utime"
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetBundleName() string { return "cisco_ios_xr" }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetYangName() string { return "idb-utime" }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) SetParent(parent types.Entity) { idbUtime.parent = parent }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetParent() types.Entity { return idbUtime.parent }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime
// CAPS Add Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetFilter() yfilter.YFilter { return capsUtime.YFilter }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) SetFilter(yf yfilter.YFilter) { capsUtime.YFilter = yf }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetGoName(yname string) string {
    return ""
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetSegmentPath() string {
    return "caps-utime"
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetBundleName() string { return "cisco_ios_xr" }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetYangName() string { return "caps-utime" }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) SetParent(parent types.Entity) { capsUtime.parent = parent }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetParent() types.Entity { return capsUtime.parent }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime
// FWD ENABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetFilter() yfilter.YFilter { return fwdEnUtime.YFilter }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) SetFilter(yf yfilter.YFilter) { fwdEnUtime.YFilter = yf }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetGoName(yname string) string {
    return ""
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetSegmentPath() string {
    return "fwd-en-utime"
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetYangName() string { return "fwd-en-utime" }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) SetParent(parent types.Entity) { fwdEnUtime.parent = parent }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetParent() types.Entity { return fwdEnUtime.parent }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime
// FWD DISABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetFilter() yfilter.YFilter { return fwdDisUtime.YFilter }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) SetFilter(yf yfilter.YFilter) { fwdDisUtime.YFilter = yf }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetGoName(yname string) string {
    return ""
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetSegmentPath() string {
    return "fwd-dis-utime"
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetYangName() string { return "fwd-dis-utime" }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) SetParent(parent types.Entity) { fwdDisUtime.parent = parent }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetParent() types.Entity { return fwdDisUtime.parent }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup
// IPv6 Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetFilter() yfilter.YFilter { return multicastGroup.YFilter }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) SetFilter(yf yfilter.YFilter) { multicastGroup.YFilter = yf }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetSegmentPath() string {
    return "multicast-group"
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = multicastGroup.Address
    return leafs
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetBundleName() string { return "cisco_ios_xr" }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetYangName() string { return "multicast-group" }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) SetParent(parent types.Entity) { multicastGroup.parent = parent }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetParent() types.Entity { return multicastGroup.parent }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = address.Address
    leafs["prefix-length"] = address.PrefixLength
    leafs["address-state"] = address.AddressState
    leafs["is-anycast"] = address.IsAnycast
    leafs["route-tag"] = address.RouteTag
    return leafs
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetYangName() string { return "address" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup
// IPv6 Client Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetFilter() yfilter.YFilter { return clientMulticastGroup.YFilter }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) SetFilter(yf yfilter.YFilter) { clientMulticastGroup.YFilter = yf }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetSegmentPath() string {
    return "client-multicast-group"
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = clientMulticastGroup.Address
    return leafs
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetBundleName() string { return "cisco_ios_xr" }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetYangName() string { return "client-multicast-group" }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) SetParent(parent types.Entity) { clientMulticastGroup.parent = parent }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetParent() types.Entity { return clientMulticastGroup.parent }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetParentYangName() string { return "global-detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs
// Brief interface IPv6 network operational
// data from global data
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief interface IPv6 network operational data for an interface. The type is
    // slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief.
    GlobalBrief []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetFilter() yfilter.YFilter { return globalBriefs.YFilter }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) SetFilter(yf yfilter.YFilter) { globalBriefs.YFilter = yf }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetGoName(yname string) string {
    if yname == "global-brief" { return "GlobalBrief" }
    return ""
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetSegmentPath() string {
    return "global-briefs"
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-brief" {
        for _, c := range globalBriefs.GlobalBrief {
            if globalBriefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief{}
        globalBriefs.GlobalBrief = append(globalBriefs.GlobalBrief, child)
        return &globalBriefs.GlobalBrief[len(globalBriefs.GlobalBrief)-1]
    }
    return nil
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalBriefs.GlobalBrief {
        children[globalBriefs.GlobalBrief[i].GetSegmentPath()] = &globalBriefs.GlobalBrief[i]
    }
    return children
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetBundleName() string { return "cisco_ios_xr" }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetYangName() string { return "global-briefs" }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) SetParent(parent types.Entity) { globalBriefs.parent = parent }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetParent() types.Entity { return globalBriefs.parent }

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetParentYangName() string { return "vrf" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief
// Brief interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetFilter() yfilter.YFilter { return globalBrief.YFilter }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) SetFilter(yf yfilter.YFilter) { globalBrief.YFilter = yf }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "line-state" { return "LineState" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "link-local-address" { return "LinkLocalAddress" }
    if yname == "address" { return "Address" }
    return ""
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetSegmentPath() string {
    return "global-brief" + "[interface-name='" + fmt.Sprintf("%v", globalBrief.InterfaceName) + "']"
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-local-address" {
        return &globalBrief.LinkLocalAddress
    }
    if childYangName == "address" {
        for _, c := range globalBrief.Address {
            if globalBrief.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address{}
        globalBrief.Address = append(globalBrief.Address, child)
        return &globalBrief.Address[len(globalBrief.Address)-1]
    }
    return nil
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-local-address"] = &globalBrief.LinkLocalAddress
    for i := range globalBrief.Address {
        children[globalBrief.Address[i].GetSegmentPath()] = &globalBrief.Address[i]
    }
    return children
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = globalBrief.InterfaceName
    leafs["line-state"] = globalBrief.LineState
    leafs["vrf-name"] = globalBrief.VrfName
    return leafs
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetBundleName() string { return "cisco_ios_xr" }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetYangName() string { return "global-brief" }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) SetParent(parent types.Entity) { globalBrief.parent = parent }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetParent() types.Entity { return globalBrief.parent }

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetParentYangName() string { return "global-briefs" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetFilter() yfilter.YFilter { return linkLocalAddress.YFilter }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) SetFilter(yf yfilter.YFilter) { linkLocalAddress.YFilter = yf }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetSegmentPath() string {
    return "link-local-address"
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = linkLocalAddress.Address
    leafs["prefix-length"] = linkLocalAddress.PrefixLength
    leafs["address-state"] = linkLocalAddress.AddressState
    leafs["is-anycast"] = linkLocalAddress.IsAnycast
    leafs["route-tag"] = linkLocalAddress.RouteTag
    return leafs
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetYangName() string { return "link-local-address" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) SetParent(parent types.Entity) { linkLocalAddress.parent = parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetParent() types.Entity { return linkLocalAddress.parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetParentYangName() string { return "global-brief" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = address.Address
    leafs["prefix-length"] = address.PrefixLength
    leafs["address-state"] = address.AddressState
    leafs["is-anycast"] = address.IsAnycast
    leafs["route-tag"] = address.RouteTag
    return leafs
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetYangName() string { return "address" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetParentYangName() string { return "global-brief" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
// Detail interface IPv4 network operational
// data for a node
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail interface IPv6 network operational data for an interface. The type
    // is slice of Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail.
    Detail []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    return ""
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        for _, c := range details.Detail {
            if details.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail{}
        details.Detail = append(details.Detail, child)
        return &details.Detail[len(details.Detail)-1]
    }
    return nil
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range details.Detail {
        children[details.Detail[i].GetSegmentPath()] = &details.Detail[i]
    }
    return children
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetYangName() string { return "details" }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetParentYangName() string { return "vrf" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
// Detail interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // IPv6 MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // IPv6 Operation State. The type is Ipv6MaOperState.
    OperationState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // ICMP unreach Enable. The type is bool.
    IsIcmpUnreachEnabled interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress

    // IPv6 Access Control List.
    AccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList

    // Multi IPv6 Access Control List.
    MultiAccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList

    // RPF config on the interface.
    Rpf Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa

    // Address Publish Time.
    Utime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime

    // IDB Create Time.
    IdbUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime

    // IPv6 Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup.
    MulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address

    // IPv6 Client Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup.
    ClientMulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "line-state" { return "LineState" }
    if yname == "mtu" { return "Mtu" }
    if yname == "operation-state" { return "OperationState" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "is-icmp-unreach-enabled" { return "IsIcmpUnreachEnabled" }
    if yname == "rg-id-exists" { return "RgIdExists" }
    if yname == "mlacp-active" { return "MlacpActive" }
    if yname == "flow-tag-src" { return "FlowTagSrc" }
    if yname == "flow-tag-dst" { return "FlowTagDst" }
    if yname == "link-local-address" { return "LinkLocalAddress" }
    if yname == "access-control-list" { return "AccessControlList" }
    if yname == "multi-access-control-list" { return "MultiAccessControlList" }
    if yname == "rpf" { return "Rpf" }
    if yname == "bgp-pa" { return "BgpPa" }
    if yname == "utime" { return "Utime" }
    if yname == "idb-utime" { return "IdbUtime" }
    if yname == "caps-utime" { return "CapsUtime" }
    if yname == "fwd-en-utime" { return "FwdEnUtime" }
    if yname == "fwd-dis-utime" { return "FwdDisUtime" }
    if yname == "multicast-group" { return "MulticastGroup" }
    if yname == "address" { return "Address" }
    if yname == "client-multicast-group" { return "ClientMulticastGroup" }
    return ""
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetSegmentPath() string {
    return "detail" + "[interface-name='" + fmt.Sprintf("%v", detail.InterfaceName) + "']"
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-local-address" {
        return &detail.LinkLocalAddress
    }
    if childYangName == "access-control-list" {
        return &detail.AccessControlList
    }
    if childYangName == "multi-access-control-list" {
        return &detail.MultiAccessControlList
    }
    if childYangName == "rpf" {
        return &detail.Rpf
    }
    if childYangName == "bgp-pa" {
        return &detail.BgpPa
    }
    if childYangName == "utime" {
        return &detail.Utime
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
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup{}
        detail.MulticastGroup = append(detail.MulticastGroup, child)
        return &detail.MulticastGroup[len(detail.MulticastGroup)-1]
    }
    if childYangName == "address" {
        for _, c := range detail.Address {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address{}
        detail.Address = append(detail.Address, child)
        return &detail.Address[len(detail.Address)-1]
    }
    if childYangName == "client-multicast-group" {
        for _, c := range detail.ClientMulticastGroup {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup{}
        detail.ClientMulticastGroup = append(detail.ClientMulticastGroup, child)
        return &detail.ClientMulticastGroup[len(detail.ClientMulticastGroup)-1]
    }
    return nil
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-local-address"] = &detail.LinkLocalAddress
    children["access-control-list"] = &detail.AccessControlList
    children["multi-access-control-list"] = &detail.MultiAccessControlList
    children["rpf"] = &detail.Rpf
    children["bgp-pa"] = &detail.BgpPa
    children["utime"] = &detail.Utime
    children["idb-utime"] = &detail.IdbUtime
    children["caps-utime"] = &detail.CapsUtime
    children["fwd-en-utime"] = &detail.FwdEnUtime
    children["fwd-dis-utime"] = &detail.FwdDisUtime
    for i := range detail.MulticastGroup {
        children[detail.MulticastGroup[i].GetSegmentPath()] = &detail.MulticastGroup[i]
    }
    for i := range detail.Address {
        children[detail.Address[i].GetSegmentPath()] = &detail.Address[i]
    }
    for i := range detail.ClientMulticastGroup {
        children[detail.ClientMulticastGroup[i].GetSegmentPath()] = &detail.ClientMulticastGroup[i]
    }
    return children
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = detail.InterfaceName
    leafs["line-state"] = detail.LineState
    leafs["mtu"] = detail.Mtu
    leafs["operation-state"] = detail.OperationState
    leafs["vrf-name"] = detail.VrfName
    leafs["is-icmp-unreach-enabled"] = detail.IsIcmpUnreachEnabled
    leafs["rg-id-exists"] = detail.RgIdExists
    leafs["mlacp-active"] = detail.MlacpActive
    leafs["flow-tag-src"] = detail.FlowTagSrc
    leafs["flow-tag-dst"] = detail.FlowTagDst
    return leafs
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetYangName() string { return "detail" }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetParentYangName() string { return "details" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetFilter() yfilter.YFilter { return linkLocalAddress.YFilter }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) SetFilter(yf yfilter.YFilter) { linkLocalAddress.YFilter = yf }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetSegmentPath() string {
    return "link-local-address"
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = linkLocalAddress.Address
    leafs["prefix-length"] = linkLocalAddress.PrefixLength
    leafs["address-state"] = linkLocalAddress.AddressState
    leafs["is-anycast"] = linkLocalAddress.IsAnycast
    leafs["route-tag"] = linkLocalAddress.RouteTag
    return leafs
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetYangName() string { return "link-local-address" }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) SetParent(parent types.Entity) { linkLocalAddress.parent = parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetParent() types.Entity { return linkLocalAddress.parent }

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList
// IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    InBound interface{}

    // ACL applied to outgoing packets. The type is string.
    OutBound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetFilter() yfilter.YFilter { return accessControlList.YFilter }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) SetFilter(yf yfilter.YFilter) { accessControlList.YFilter = yf }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetGoName(yname string) string {
    if yname == "in-bound" { return "InBound" }
    if yname == "out-bound" { return "OutBound" }
    if yname == "common-in-bound" { return "CommonInBound" }
    if yname == "common-out-bound" { return "CommonOutBound" }
    return ""
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetSegmentPath() string {
    return "access-control-list"
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-bound"] = accessControlList.InBound
    leafs["out-bound"] = accessControlList.OutBound
    leafs["common-in-bound"] = accessControlList.CommonInBound
    leafs["common-out-bound"] = accessControlList.CommonOutBound
    return leafs
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetBundleName() string { return "cisco_ios_xr" }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetYangName() string { return "access-control-list" }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) SetParent(parent types.Entity) { accessControlList.parent = parent }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetParent() types.Entity { return accessControlList.parent }

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList
// Multi IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of string.
    Inbound []interface{}

    // Outbound ACLs. The type is slice of string.
    Outbound []interface{}

    // Common ACLs. The type is slice of string.
    Common []interface{}
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetFilter() yfilter.YFilter { return multiAccessControlList.YFilter }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) SetFilter(yf yfilter.YFilter) { multiAccessControlList.YFilter = yf }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    if yname == "common" { return "Common" }
    return ""
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetSegmentPath() string {
    return "multi-access-control-list"
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inbound"] = multiAccessControlList.Inbound
    leafs["outbound"] = multiAccessControlList.Outbound
    leafs["common"] = multiAccessControlList.Common
    return leafs
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetBundleName() string { return "cisco_ios_xr" }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetYangName() string { return "multi-access-control-list" }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) SetParent(parent types.Entity) { multiAccessControlList.parent = parent }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetParent() types.Entity { return multiAccessControlList.parent }

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf
// RPF config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is interface{} with range: 0..4294967295.
    Mode interface{}
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-default-route" { return "AllowDefaultRoute" }
    if yname == "allow-self-ping" { return "AllowSelfPing" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpf.Enable
    leafs["allow-default-route"] = rpf.AllowDefaultRoute
    leafs["allow-self-ping"] = rpf.AllowSelfPing
    leafs["mode"] = rpf.Mode
    return leafs
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetYangName() string { return "rpf" }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa
// BGP PA config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input

    // BGP PA output config.
    Output Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetFilter() yfilter.YFilter { return bgpPa.YFilter }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) SetFilter(yf yfilter.YFilter) { bgpPa.YFilter = yf }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetSegmentPath() string {
    return "bgp-pa"
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &bgpPa.Input
    }
    if childYangName == "output" {
        return &bgpPa.Output
    }
    return nil
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &bgpPa.Input
    children["output"] = &bgpPa.Output
    return children
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetBundleName() string { return "cisco_ios_xr" }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetYangName() string { return "bgp-pa" }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) SetParent(parent types.Entity) { bgpPa.parent = parent }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetParent() types.Entity { return bgpPa.parent }

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input
// BGP PA input config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetSegmentPath() string {
    return "input"
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = input.Enable
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetYangName() string { return "input" }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetParent() types.Entity { return input.parent }

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetParentYangName() string { return "bgp-pa" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
// BGP PA output config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetSegmentPath() string {
    return "output"
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = output.Enable
    leafs["source"] = output.Source
    leafs["destination"] = output.Destination
    return leafs
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetYangName() string { return "output" }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetParent() types.Entity { return output.parent }

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetParentYangName() string { return "bgp-pa" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime
// Address Publish Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetFilter() yfilter.YFilter { return utime.YFilter }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) SetFilter(yf yfilter.YFilter) { utime.YFilter = yf }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetGoName(yname string) string {
    return ""
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetSegmentPath() string {
    return "utime"
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetBundleName() string { return "cisco_ios_xr" }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetYangName() string { return "utime" }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) SetParent(parent types.Entity) { utime.parent = parent }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetParent() types.Entity { return utime.parent }

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime
// IDB Create Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetFilter() yfilter.YFilter { return idbUtime.YFilter }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) SetFilter(yf yfilter.YFilter) { idbUtime.YFilter = yf }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetGoName(yname string) string {
    return ""
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetSegmentPath() string {
    return "idb-utime"
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetBundleName() string { return "cisco_ios_xr" }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetYangName() string { return "idb-utime" }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) SetParent(parent types.Entity) { idbUtime.parent = parent }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetParent() types.Entity { return idbUtime.parent }

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime
// CAPS Add Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetFilter() yfilter.YFilter { return capsUtime.YFilter }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) SetFilter(yf yfilter.YFilter) { capsUtime.YFilter = yf }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetGoName(yname string) string {
    return ""
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetSegmentPath() string {
    return "caps-utime"
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetBundleName() string { return "cisco_ios_xr" }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetYangName() string { return "caps-utime" }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) SetParent(parent types.Entity) { capsUtime.parent = parent }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetParent() types.Entity { return capsUtime.parent }

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime
// FWD ENABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetFilter() yfilter.YFilter { return fwdEnUtime.YFilter }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) SetFilter(yf yfilter.YFilter) { fwdEnUtime.YFilter = yf }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetGoName(yname string) string {
    return ""
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetSegmentPath() string {
    return "fwd-en-utime"
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetYangName() string { return "fwd-en-utime" }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) SetParent(parent types.Entity) { fwdEnUtime.parent = parent }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetParent() types.Entity { return fwdEnUtime.parent }

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime
// FWD DISABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetFilter() yfilter.YFilter { return fwdDisUtime.YFilter }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) SetFilter(yf yfilter.YFilter) { fwdDisUtime.YFilter = yf }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetGoName(yname string) string {
    return ""
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetSegmentPath() string {
    return "fwd-dis-utime"
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetBundleName() string { return "cisco_ios_xr" }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetYangName() string { return "fwd-dis-utime" }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) SetParent(parent types.Entity) { fwdDisUtime.parent = parent }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetParent() types.Entity { return fwdDisUtime.parent }

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup
// IPv6 Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetFilter() yfilter.YFilter { return multicastGroup.YFilter }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) SetFilter(yf yfilter.YFilter) { multicastGroup.YFilter = yf }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetSegmentPath() string {
    return "multicast-group"
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = multicastGroup.Address
    return leafs
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetBundleName() string { return "cisco_ios_xr" }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetYangName() string { return "multicast-group" }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) SetParent(parent types.Entity) { multicastGroup.parent = parent }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetParent() types.Entity { return multicastGroup.parent }

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-state" { return "AddressState" }
    if yname == "is-anycast" { return "IsAnycast" }
    if yname == "route-tag" { return "RouteTag" }
    return ""
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = address.Address
    leafs["prefix-length"] = address.PrefixLength
    leafs["address-state"] = address.AddressState
    leafs["is-anycast"] = address.IsAnycast
    leafs["route-tag"] = address.RouteTag
    return leafs
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetYangName() string { return "address" }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup
// IPv6 Client Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetFilter() yfilter.YFilter { return clientMulticastGroup.YFilter }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) SetFilter(yf yfilter.YFilter) { clientMulticastGroup.YFilter = yf }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetSegmentPath() string {
    return "client-multicast-group"
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = clientMulticastGroup.Address
    return leafs
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetBundleName() string { return "cisco_ios_xr" }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetYangName() string { return "client-multicast-group" }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) SetParent(parent types.Entity) { clientMulticastGroup.parent = parent }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetParent() types.Entity { return clientMulticastGroup.parent }

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetParentYangName() string { return "detail" }

// Ipv6Network_Nodes_Node_InterfaceData_Summary
// Summary of IPv6 network operational interface
// data on a node
type Ipv6Network_Nodes_Node_InterfaceData_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces (up,down) with basecaps up. The type is interface{}
    // with range: 0..4294967295.
    IfUpDownBasecapsUp interface{}

    // Number of interfaces (up,up).
    IfUpUp Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp

    // Number of interfaces (up,down).
    IfUpDown Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown

    // Number of interfaces (down,down).
    IfDownDown Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown

    // Number of interfaces (shutdown,down).
    IfShutdownDown Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
}

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetGoName(yname string) string {
    if yname == "if-up-down-basecaps-up" { return "IfUpDownBasecapsUp" }
    if yname == "if-up-up" { return "IfUpUp" }
    if yname == "if-up-down" { return "IfUpDown" }
    if yname == "if-down-down" { return "IfDownDown" }
    if yname == "if-shutdown-down" { return "IfShutdownDown" }
    return ""
}

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
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

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["if-up-up"] = &summary.IfUpUp
    children["if-up-down"] = &summary.IfUpDown
    children["if-down-down"] = &summary.IfDownDown
    children["if-shutdown-down"] = &summary.IfShutdownDown
    return children
}

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-up-down-basecaps-up"] = summary.IfUpDownBasecapsUp
    return leafs
}

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetYangName() string { return "summary" }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetParentYangName() string { return "interface-data" }

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp
// Number of interfaces (up,up)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetFilter() yfilter.YFilter { return ifUpUp.YFilter }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) SetFilter(yf yfilter.YFilter) { ifUpUp.YFilter = yf }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetSegmentPath() string {
    return "if-up-up"
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifUpUp.IpAssigned
    leafs["ip-unnumbered"] = ifUpUp.IpUnnumbered
    leafs["ip-unassigned"] = ifUpUp.IpUnassigned
    return leafs
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetBundleName() string { return "cisco_ios_xr" }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetYangName() string { return "if-up-up" }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) SetParent(parent types.Entity) { ifUpUp.parent = parent }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetParent() types.Entity { return ifUpUp.parent }

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetParentYangName() string { return "summary" }

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown
// Number of interfaces (up,down)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetFilter() yfilter.YFilter { return ifUpDown.YFilter }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) SetFilter(yf yfilter.YFilter) { ifUpDown.YFilter = yf }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetSegmentPath() string {
    return "if-up-down"
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifUpDown.IpAssigned
    leafs["ip-unnumbered"] = ifUpDown.IpUnnumbered
    leafs["ip-unassigned"] = ifUpDown.IpUnassigned
    return leafs
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetBundleName() string { return "cisco_ios_xr" }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetYangName() string { return "if-up-down" }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) SetParent(parent types.Entity) { ifUpDown.parent = parent }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetParent() types.Entity { return ifUpDown.parent }

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetParentYangName() string { return "summary" }

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown
// Number of interfaces (down,down)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetFilter() yfilter.YFilter { return ifDownDown.YFilter }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) SetFilter(yf yfilter.YFilter) { ifDownDown.YFilter = yf }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetSegmentPath() string {
    return "if-down-down"
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifDownDown.IpAssigned
    leafs["ip-unnumbered"] = ifDownDown.IpUnnumbered
    leafs["ip-unassigned"] = ifDownDown.IpUnassigned
    return leafs
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetBundleName() string { return "cisco_ios_xr" }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetYangName() string { return "if-down-down" }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) SetParent(parent types.Entity) { ifDownDown.parent = parent }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetParent() types.Entity { return ifDownDown.parent }

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetParentYangName() string { return "summary" }

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
// Number of interfaces (shutdown,down)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetFilter() yfilter.YFilter { return ifShutdownDown.YFilter }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) SetFilter(yf yfilter.YFilter) { ifShutdownDown.YFilter = yf }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetGoName(yname string) string {
    if yname == "ip-assigned" { return "IpAssigned" }
    if yname == "ip-unnumbered" { return "IpUnnumbered" }
    if yname == "ip-unassigned" { return "IpUnassigned" }
    return ""
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetSegmentPath() string {
    return "if-shutdown-down"
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-assigned"] = ifShutdownDown.IpAssigned
    leafs["ip-unnumbered"] = ifShutdownDown.IpUnnumbered
    leafs["ip-unassigned"] = ifShutdownDown.IpUnassigned
    return leafs
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetBundleName() string { return "cisco_ios_xr" }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetYangName() string { return "if-shutdown-down" }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) SetParent(parent types.Entity) { ifShutdownDown.parent = parent }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetParent() types.Entity { return ifShutdownDown.parent }

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetParentYangName() string { return "summary" }

