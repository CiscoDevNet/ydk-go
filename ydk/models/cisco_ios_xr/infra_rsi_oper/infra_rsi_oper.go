// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rsi package operational data.
// 
// This module contains definitions
// for the following management objects:
//   vrf-group: VRF group operational data
//   srlg: srlg
//   selective-vrf-download: selective vrf download
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_rsi_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rsi_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-oper vrf-group}", reflect.TypeOf(VrfGroup{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-oper:vrf-group", reflect.TypeOf(VrfGroup{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-oper srlg}", reflect.TypeOf(Srlg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-oper:srlg", reflect.TypeOf(Srlg{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-oper selective-vrf-download}", reflect.TypeOf(SelectiveVrfDownload{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-oper:selective-vrf-download", reflect.TypeOf(SelectiveVrfDownload{}))
}

// Priority represents Priority
type Priority string

const (
    // Critical
    Priority_critical Priority = "critical"

    // High
    Priority_high Priority = "high"

    // Medium
    Priority_medium Priority = "medium"

    // Low
    Priority_low Priority = "low"

    // Very low
    Priority_very_low Priority = "very-low"
)

// Source represents Source
type Source string

const (
    // Configured
    Source_configured Source = "configured"

    // From group
    Source_from_group Source = "from-group"

    // Inherited
    Source_inherited Source = "inherited"

    // From optical
    Source_from_optical Source = "from-optical"

    // Configured and notified
    Source_configured_and_notified Source = "configured-and-notified"

    // From group and notified
    Source_from_group_and_notified Source = "from-group-and-notified"

    // Inherited and notified
    Source_inherited_and_notified Source = "inherited-and-notified"

    // From optical and notified
    Source_from_optical_and_notified Source = "from-optical-and-notified"
)

// VrfGroup
// VRF group operational data
type VrfGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node operational data.
    Nodes VrfGroup_Nodes
}

func (vrfGroup *VrfGroup) GetFilter() yfilter.YFilter { return vrfGroup.YFilter }

func (vrfGroup *VrfGroup) SetFilter(yf yfilter.YFilter) { vrfGroup.YFilter = yf }

func (vrfGroup *VrfGroup) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (vrfGroup *VrfGroup) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-oper:vrf-group"
}

func (vrfGroup *VrfGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &vrfGroup.Nodes
    }
    return nil
}

func (vrfGroup *VrfGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &vrfGroup.Nodes
    return children
}

func (vrfGroup *VrfGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfGroup *VrfGroup) GetBundleName() string { return "cisco_ios_xr" }

func (vrfGroup *VrfGroup) GetYangName() string { return "vrf-group" }

func (vrfGroup *VrfGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfGroup *VrfGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfGroup *VrfGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfGroup *VrfGroup) SetParent(parent types.Entity) { vrfGroup.parent = parent }

func (vrfGroup *VrfGroup) GetParent() types.Entity { return vrfGroup.parent }

func (vrfGroup *VrfGroup) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-oper" }

// VrfGroup_Nodes
// Node operational data
type VrfGroup_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node details. The type is slice of VrfGroup_Nodes_Node.
    Node []VrfGroup_Nodes_Node
}

func (nodes *VrfGroup_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *VrfGroup_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *VrfGroup_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *VrfGroup_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *VrfGroup_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfGroup_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *VrfGroup_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *VrfGroup_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *VrfGroup_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *VrfGroup_Nodes) GetYangName() string { return "nodes" }

func (nodes *VrfGroup_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *VrfGroup_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *VrfGroup_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *VrfGroup_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *VrfGroup_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *VrfGroup_Nodes) GetParentYangName() string { return "vrf-group" }

// VrfGroup_Nodes_Node
// Node details
type VrfGroup_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Group operational data.
    Groups VrfGroup_Nodes_Node_Groups
}

func (node *VrfGroup_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *VrfGroup_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *VrfGroup_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "groups" { return "Groups" }
    return ""
}

func (node *VrfGroup_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *VrfGroup_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &node.Groups
    }
    return nil
}

func (node *VrfGroup_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &node.Groups
    return children
}

func (node *VrfGroup_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *VrfGroup_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *VrfGroup_Nodes_Node) GetYangName() string { return "node" }

func (node *VrfGroup_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *VrfGroup_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *VrfGroup_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *VrfGroup_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *VrfGroup_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *VrfGroup_Nodes_Node) GetParentYangName() string { return "nodes" }

// VrfGroup_Nodes_Node_Groups
// Group operational data
type VrfGroup_Nodes_Node_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group details. The type is slice of VrfGroup_Nodes_Node_Groups_Group.
    Group []VrfGroup_Nodes_Node_Groups_Group
}

func (groups *VrfGroup_Nodes_Node_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *VrfGroup_Nodes_Node_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *VrfGroup_Nodes_Node_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *VrfGroup_Nodes_Node_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *VrfGroup_Nodes_Node_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfGroup_Nodes_Node_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *VrfGroup_Nodes_Node_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *VrfGroup_Nodes_Node_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *VrfGroup_Nodes_Node_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *VrfGroup_Nodes_Node_Groups) GetYangName() string { return "groups" }

func (groups *VrfGroup_Nodes_Node_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *VrfGroup_Nodes_Node_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *VrfGroup_Nodes_Node_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *VrfGroup_Nodes_Node_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *VrfGroup_Nodes_Node_Groups) GetParent() types.Entity { return groups.parent }

func (groups *VrfGroup_Nodes_Node_Groups) GetParentYangName() string { return "node" }

// VrfGroup_Nodes_Node_Groups_Group
// Group details
type VrfGroup_Nodes_Node_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with length: 1..32.
    GroupName interface{}

    // Number of VRFs in this VRF group. The type is interface{} with range:
    // 0..4294967295.
    VrFs interface{}

    // VRF group not present but used. The type is bool.
    ForwardReference interface{}

    // VRF group's VRF. The type is slice of VrfGroup_Nodes_Node_Groups_Group_Vrf.
    Vrf []VrfGroup_Nodes_Node_Groups_Group_Vrf
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *VrfGroup_Nodes_Node_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetGoName(yname string) string {
    if yname == "group-name" { return "GroupName" }
    if yname == "vr-fs" { return "VrFs" }
    if yname == "forward-reference" { return "ForwardReference" }
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range group.Vrf {
            if group.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfGroup_Nodes_Node_Groups_Group_Vrf{}
        group.Vrf = append(group.Vrf, child)
        return &group.Vrf[len(group.Vrf)-1]
    }
    return nil
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range group.Vrf {
        children[group.Vrf[i].GetSegmentPath()] = &group.Vrf[i]
    }
    return children
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-name"] = group.GroupName
    leafs["vr-fs"] = group.VrFs
    leafs["forward-reference"] = group.ForwardReference
    return leafs
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetYangName() string { return "group" }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *VrfGroup_Nodes_Node_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *VrfGroup_Nodes_Node_Groups_Group) GetParentYangName() string { return "groups" }

// VrfGroup_Nodes_Node_Groups_Group_Vrf
// VRF group's VRF
type VrfGroup_Nodes_Node_Groups_Group_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetSegmentPath() string {
    return "vrf"
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetYangName() string { return "vrf" }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetParentYangName() string { return "group" }

// Srlg
// srlg
type Srlg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set of SRLG name, value maps configured.
    SrlgMaps Srlg_SrlgMaps

    // RSI SRLG operational data.
    Nodes Srlg_Nodes

    // Set of SRLG names configured.
    InterfaceSrlgNames Srlg_InterfaceSrlgNames
}

func (srlg *Srlg) GetFilter() yfilter.YFilter { return srlg.YFilter }

func (srlg *Srlg) SetFilter(yf yfilter.YFilter) { srlg.YFilter = yf }

func (srlg *Srlg) GetGoName(yname string) string {
    if yname == "srlg-maps" { return "SrlgMaps" }
    if yname == "nodes" { return "Nodes" }
    if yname == "interface-srlg-names" { return "InterfaceSrlgNames" }
    return ""
}

func (srlg *Srlg) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-oper:srlg"
}

func (srlg *Srlg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-maps" {
        return &srlg.SrlgMaps
    }
    if childYangName == "nodes" {
        return &srlg.Nodes
    }
    if childYangName == "interface-srlg-names" {
        return &srlg.InterfaceSrlgNames
    }
    return nil
}

func (srlg *Srlg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srlg-maps"] = &srlg.SrlgMaps
    children["nodes"] = &srlg.Nodes
    children["interface-srlg-names"] = &srlg.InterfaceSrlgNames
    return children
}

func (srlg *Srlg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srlg *Srlg) GetBundleName() string { return "cisco_ios_xr" }

func (srlg *Srlg) GetYangName() string { return "srlg" }

func (srlg *Srlg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlg *Srlg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlg *Srlg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlg *Srlg) SetParent(parent types.Entity) { srlg.parent = parent }

func (srlg *Srlg) GetParent() types.Entity { return srlg.parent }

func (srlg *Srlg) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-oper" }

// Srlg_SrlgMaps
// Set of SRLG name, value maps configured
type Srlg_SrlgMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of Srlg_SrlgMaps_SrlgMap.
    SrlgMap []Srlg_SrlgMaps_SrlgMap
}

func (srlgMaps *Srlg_SrlgMaps) GetFilter() yfilter.YFilter { return srlgMaps.YFilter }

func (srlgMaps *Srlg_SrlgMaps) SetFilter(yf yfilter.YFilter) { srlgMaps.YFilter = yf }

func (srlgMaps *Srlg_SrlgMaps) GetGoName(yname string) string {
    if yname == "srlg-map" { return "SrlgMap" }
    return ""
}

func (srlgMaps *Srlg_SrlgMaps) GetSegmentPath() string {
    return "srlg-maps"
}

func (srlgMaps *Srlg_SrlgMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-map" {
        for _, c := range srlgMaps.SrlgMap {
            if srlgMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_SrlgMaps_SrlgMap{}
        srlgMaps.SrlgMap = append(srlgMaps.SrlgMap, child)
        return &srlgMaps.SrlgMap[len(srlgMaps.SrlgMap)-1]
    }
    return nil
}

func (srlgMaps *Srlg_SrlgMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srlgMaps.SrlgMap {
        children[srlgMaps.SrlgMap[i].GetSegmentPath()] = &srlgMaps.SrlgMap[i]
    }
    return children
}

func (srlgMaps *Srlg_SrlgMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srlgMaps *Srlg_SrlgMaps) GetBundleName() string { return "cisco_ios_xr" }

func (srlgMaps *Srlg_SrlgMaps) GetYangName() string { return "srlg-maps" }

func (srlgMaps *Srlg_SrlgMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgMaps *Srlg_SrlgMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgMaps *Srlg_SrlgMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgMaps *Srlg_SrlgMaps) SetParent(parent types.Entity) { srlgMaps.parent = parent }

func (srlgMaps *Srlg_SrlgMaps) GetParent() types.Entity { return srlgMaps.parent }

func (srlgMaps *Srlg_SrlgMaps) GetParentYangName() string { return "srlg" }

// Srlg_SrlgMaps_SrlgMap
// Configured SRLG name details 
type Srlg_SrlgMaps_SrlgMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // SRLG name. The type is string.
    SrlgNameXr interface{}
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetFilter() yfilter.YFilter { return srlgMap.YFilter }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) SetFilter(yf yfilter.YFilter) { srlgMap.YFilter = yf }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "srlg-name-xr" { return "SrlgNameXr" }
    return ""
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetSegmentPath() string {
    return "srlg-map" + "[srlg-name='" + fmt.Sprintf("%v", srlgMap.SrlgName) + "']"
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-name"] = srlgMap.SrlgName
    leafs["srlg-value"] = srlgMap.SrlgValue
    leafs["srlg-name-xr"] = srlgMap.SrlgNameXr
    return leafs
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetBundleName() string { return "cisco_ios_xr" }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetYangName() string { return "srlg-map" }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) SetParent(parent types.Entity) { srlgMap.parent = parent }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetParent() types.Entity { return srlgMap.parent }

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetParentYangName() string { return "srlg-maps" }

// Srlg_Nodes
// RSI SRLG operational data
type Srlg_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSI SRLG operational data. The type is slice of Srlg_Nodes_Node.
    Node []Srlg_Nodes_Node
}

func (nodes *Srlg_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Srlg_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Srlg_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Srlg_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Srlg_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Srlg_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Srlg_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Srlg_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Srlg_Nodes) GetYangName() string { return "nodes" }

func (nodes *Srlg_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Srlg_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Srlg_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Srlg_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Srlg_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Srlg_Nodes) GetParentYangName() string { return "srlg" }

// Srlg_Nodes_Node
// RSI SRLG operational data
type Srlg_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Set of SRLG name, value maps configured.
    SrlgMaps Srlg_Nodes_Node_SrlgMaps

    // Set of Groups configured for SRLG.
    Groups Srlg_Nodes_Node_Groups

    // Set of inherit locations configured for SRLG.
    InheritNodes Srlg_Nodes_Node_InheritNodes

    // Set of interfaces configured for SRLG.
    Interfaces Srlg_Nodes_Node_Interfaces

    // Set of interfaces configured for SRLG.
    InterfaceDetails Srlg_Nodes_Node_InterfaceDetails

    // Set of SRLG values configured.
    SrlgValues Srlg_Nodes_Node_SrlgValues

    // Set of SRLG names configured.
    InterfaceSrlgNames Srlg_Nodes_Node_InterfaceSrlgNames
}

func (node *Srlg_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Srlg_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Srlg_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "srlg-maps" { return "SrlgMaps" }
    if yname == "groups" { return "Groups" }
    if yname == "inherit-nodes" { return "InheritNodes" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "interface-details" { return "InterfaceDetails" }
    if yname == "srlg-values" { return "SrlgValues" }
    if yname == "interface-srlg-names" { return "InterfaceSrlgNames" }
    return ""
}

func (node *Srlg_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Srlg_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-maps" {
        return &node.SrlgMaps
    }
    if childYangName == "groups" {
        return &node.Groups
    }
    if childYangName == "inherit-nodes" {
        return &node.InheritNodes
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "interface-details" {
        return &node.InterfaceDetails
    }
    if childYangName == "srlg-values" {
        return &node.SrlgValues
    }
    if childYangName == "interface-srlg-names" {
        return &node.InterfaceSrlgNames
    }
    return nil
}

func (node *Srlg_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srlg-maps"] = &node.SrlgMaps
    children["groups"] = &node.Groups
    children["inherit-nodes"] = &node.InheritNodes
    children["interfaces"] = &node.Interfaces
    children["interface-details"] = &node.InterfaceDetails
    children["srlg-values"] = &node.SrlgValues
    children["interface-srlg-names"] = &node.InterfaceSrlgNames
    return children
}

func (node *Srlg_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Srlg_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Srlg_Nodes_Node) GetYangName() string { return "node" }

func (node *Srlg_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Srlg_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Srlg_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Srlg_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Srlg_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Srlg_Nodes_Node) GetParentYangName() string { return "nodes" }

// Srlg_Nodes_Node_SrlgMaps
// Set of SRLG name, value maps configured
type Srlg_Nodes_Node_SrlgMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_Nodes_Node_SrlgMaps_SrlgMap.
    SrlgMap []Srlg_Nodes_Node_SrlgMaps_SrlgMap
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetFilter() yfilter.YFilter { return srlgMaps.YFilter }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) SetFilter(yf yfilter.YFilter) { srlgMaps.YFilter = yf }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetGoName(yname string) string {
    if yname == "srlg-map" { return "SrlgMap" }
    return ""
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetSegmentPath() string {
    return "srlg-maps"
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-map" {
        for _, c := range srlgMaps.SrlgMap {
            if srlgMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_SrlgMaps_SrlgMap{}
        srlgMaps.SrlgMap = append(srlgMaps.SrlgMap, child)
        return &srlgMaps.SrlgMap[len(srlgMaps.SrlgMap)-1]
    }
    return nil
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srlgMaps.SrlgMap {
        children[srlgMaps.SrlgMap[i].GetSegmentPath()] = &srlgMaps.SrlgMap[i]
    }
    return children
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetBundleName() string { return "cisco_ios_xr" }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetYangName() string { return "srlg-maps" }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) SetParent(parent types.Entity) { srlgMaps.parent = parent }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetParent() types.Entity { return srlgMaps.parent }

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_SrlgMaps_SrlgMap
// Configured SRLG name details 
type Srlg_Nodes_Node_SrlgMaps_SrlgMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // SRLG name. The type is string.
    SrlgNameXr interface{}
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetFilter() yfilter.YFilter { return srlgMap.YFilter }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) SetFilter(yf yfilter.YFilter) { srlgMap.YFilter = yf }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "srlg-name-xr" { return "SrlgNameXr" }
    return ""
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetSegmentPath() string {
    return "srlg-map" + "[srlg-name='" + fmt.Sprintf("%v", srlgMap.SrlgName) + "']"
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-name"] = srlgMap.SrlgName
    leafs["srlg-value"] = srlgMap.SrlgValue
    leafs["srlg-name-xr"] = srlgMap.SrlgNameXr
    return leafs
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetBundleName() string { return "cisco_ios_xr" }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetYangName() string { return "srlg-map" }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) SetParent(parent types.Entity) { srlgMap.parent = parent }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetParent() types.Entity { return srlgMap.parent }

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetParentYangName() string { return "srlg-maps" }

// Srlg_Nodes_Node_Groups
// Set of Groups configured for SRLG
type Srlg_Nodes_Node_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG group details. The type is slice of Srlg_Nodes_Node_Groups_Group.
    Group []Srlg_Nodes_Node_Groups_Group
}

func (groups *Srlg_Nodes_Node_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Srlg_Nodes_Node_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Srlg_Nodes_Node_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Srlg_Nodes_Node_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Srlg_Nodes_Node_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Srlg_Nodes_Node_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Srlg_Nodes_Node_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Srlg_Nodes_Node_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Srlg_Nodes_Node_Groups) GetYangName() string { return "groups" }

func (groups *Srlg_Nodes_Node_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Srlg_Nodes_Node_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Srlg_Nodes_Node_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Srlg_Nodes_Node_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Srlg_Nodes_Node_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Srlg_Nodes_Node_Groups) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_Groups_Group
// SRLG group details
type Srlg_Nodes_Node_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupName interface{}

    // Group name. The type is string.
    GroupNameXr interface{}

    // Group values. The type is interface{} with range: 0..4294967295.
    GroupValues interface{}

    // SRLG attribute. The type is slice of
    // Srlg_Nodes_Node_Groups_Group_SrlgAttribute.
    SrlgAttribute []Srlg_Nodes_Node_Groups_Group_SrlgAttribute
}

func (group *Srlg_Nodes_Node_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Srlg_Nodes_Node_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Srlg_Nodes_Node_Groups_Group) GetGoName(yname string) string {
    if yname == "group-name" { return "GroupName" }
    if yname == "group-name-xr" { return "GroupNameXr" }
    if yname == "group-values" { return "GroupValues" }
    if yname == "srlg-attribute" { return "SrlgAttribute" }
    return ""
}

func (group *Srlg_Nodes_Node_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
}

func (group *Srlg_Nodes_Node_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-attribute" {
        for _, c := range group.SrlgAttribute {
            if group.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_Groups_Group_SrlgAttribute{}
        group.SrlgAttribute = append(group.SrlgAttribute, child)
        return &group.SrlgAttribute[len(group.SrlgAttribute)-1]
    }
    return nil
}

func (group *Srlg_Nodes_Node_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range group.SrlgAttribute {
        children[group.SrlgAttribute[i].GetSegmentPath()] = &group.SrlgAttribute[i]
    }
    return children
}

func (group *Srlg_Nodes_Node_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-name"] = group.GroupName
    leafs["group-name-xr"] = group.GroupNameXr
    leafs["group-values"] = group.GroupValues
    return leafs
}

func (group *Srlg_Nodes_Node_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Srlg_Nodes_Node_Groups_Group) GetYangName() string { return "group" }

func (group *Srlg_Nodes_Node_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Srlg_Nodes_Node_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Srlg_Nodes_Node_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Srlg_Nodes_Node_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Srlg_Nodes_Node_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Srlg_Nodes_Node_Groups_Group) GetParentYangName() string { return "groups" }

// Srlg_Nodes_Node_Groups_Group_SrlgAttribute
// SRLG attribute
type Srlg_Nodes_Node_Groups_Group_SrlgAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetFilter() yfilter.YFilter { return srlgAttribute.YFilter }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) SetFilter(yf yfilter.YFilter) { srlgAttribute.YFilter = yf }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetGoName(yname string) string {
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "priority" { return "Priority" }
    if yname == "srlg-index" { return "SrlgIndex" }
    return ""
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetSegmentPath() string {
    return "srlg-attribute"
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-value"] = srlgAttribute.SrlgValue
    leafs["priority"] = srlgAttribute.Priority
    leafs["srlg-index"] = srlgAttribute.SrlgIndex
    return leafs
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetBundleName() string { return "cisco_ios_xr" }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetYangName() string { return "srlg-attribute" }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) SetParent(parent types.Entity) { srlgAttribute.parent = parent }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetParent() types.Entity { return srlgAttribute.parent }

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetParentYangName() string { return "group" }

// Srlg_Nodes_Node_InheritNodes
// Set of inherit locations configured for SRLG
type Srlg_Nodes_Node_InheritNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG inherit location details. The type is slice of
    // Srlg_Nodes_Node_InheritNodes_InheritNode.
    InheritNode []Srlg_Nodes_Node_InheritNodes_InheritNode
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetFilter() yfilter.YFilter { return inheritNodes.YFilter }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) SetFilter(yf yfilter.YFilter) { inheritNodes.YFilter = yf }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetGoName(yname string) string {
    if yname == "inherit-node" { return "InheritNode" }
    return ""
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetSegmentPath() string {
    return "inherit-nodes"
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inherit-node" {
        for _, c := range inheritNodes.InheritNode {
            if inheritNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_InheritNodes_InheritNode{}
        inheritNodes.InheritNode = append(inheritNodes.InheritNode, child)
        return &inheritNodes.InheritNode[len(inheritNodes.InheritNode)-1]
    }
    return nil
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inheritNodes.InheritNode {
        children[inheritNodes.InheritNode[i].GetSegmentPath()] = &inheritNodes.InheritNode[i]
    }
    return children
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetBundleName() string { return "cisco_ios_xr" }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetYangName() string { return "inherit-nodes" }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) SetParent(parent types.Entity) { inheritNodes.parent = parent }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetParent() types.Entity { return inheritNodes.parent }

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_InheritNodes_InheritNode
// SRLG inherit location details
type Srlg_Nodes_Node_InheritNodes_InheritNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Inherit node. The type is string with pattern:
    // ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    InheritNodeName interface{}

    // Inherit node name. The type is string.
    NodeName interface{}

    // Node values. The type is interface{} with range: 0..4294967295.
    NodeValues interface{}

    // SRLG attribute. The type is slice of
    // Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute.
    SrlgAttribute []Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetFilter() yfilter.YFilter { return inheritNode.YFilter }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) SetFilter(yf yfilter.YFilter) { inheritNode.YFilter = yf }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetGoName(yname string) string {
    if yname == "inherit-node-name" { return "InheritNodeName" }
    if yname == "node-name" { return "NodeName" }
    if yname == "node-values" { return "NodeValues" }
    if yname == "srlg-attribute" { return "SrlgAttribute" }
    return ""
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetSegmentPath() string {
    return "inherit-node" + "[inherit-node-name='" + fmt.Sprintf("%v", inheritNode.InheritNodeName) + "']"
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-attribute" {
        for _, c := range inheritNode.SrlgAttribute {
            if inheritNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute{}
        inheritNode.SrlgAttribute = append(inheritNode.SrlgAttribute, child)
        return &inheritNode.SrlgAttribute[len(inheritNode.SrlgAttribute)-1]
    }
    return nil
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inheritNode.SrlgAttribute {
        children[inheritNode.SrlgAttribute[i].GetSegmentPath()] = &inheritNode.SrlgAttribute[i]
    }
    return children
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inherit-node-name"] = inheritNode.InheritNodeName
    leafs["node-name"] = inheritNode.NodeName
    leafs["node-values"] = inheritNode.NodeValues
    return leafs
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetBundleName() string { return "cisco_ios_xr" }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetYangName() string { return "inherit-node" }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) SetParent(parent types.Entity) { inheritNode.parent = parent }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetParent() types.Entity { return inheritNode.parent }

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetParentYangName() string { return "inherit-nodes" }

// Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute
// SRLG attribute
type Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetFilter() yfilter.YFilter { return srlgAttribute.YFilter }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) SetFilter(yf yfilter.YFilter) { srlgAttribute.YFilter = yf }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetGoName(yname string) string {
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "priority" { return "Priority" }
    if yname == "srlg-index" { return "SrlgIndex" }
    return ""
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetSegmentPath() string {
    return "srlg-attribute"
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-value"] = srlgAttribute.SrlgValue
    leafs["priority"] = srlgAttribute.Priority
    leafs["srlg-index"] = srlgAttribute.SrlgIndex
    return leafs
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetBundleName() string { return "cisco_ios_xr" }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetYangName() string { return "srlg-attribute" }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) SetParent(parent types.Entity) { srlgAttribute.parent = parent }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetParent() types.Entity { return srlgAttribute.parent }

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetParentYangName() string { return "inherit-node" }

// Srlg_Nodes_Node_Interfaces
// Set of interfaces configured for SRLG
type Srlg_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG interface summary. The type is slice of
    // Srlg_Nodes_Node_Interfaces_Interface.
    Interface []Srlg_Nodes_Node_Interfaces_Interface
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Srlg_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Srlg_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Srlg_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_Interfaces_Interface
// SRLG interface summary
type Srlg_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Values. The type is interface{} with range: 0..4294967295.
    ValueCount interface{}

    // Registrations. The type is interface{} with range: 0..4294967295.
    Registrations interface{}

    // SRLG values. The type is slice of interface{} with range: 0..4294967295.
    SrlgValue []interface{}
}

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Srlg_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "value-count" { return "ValueCount" }
    if yname == "registrations" { return "Registrations" }
    if yname == "srlg-value" { return "SrlgValue" }
    return ""
}

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-name-xr"] = self.InterfaceNameXr
    leafs["value-count"] = self.ValueCount
    leafs["registrations"] = self.Registrations
    leafs["srlg-value"] = self.SrlgValue
    return leafs
}

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Srlg_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Srlg_Nodes_Node_InterfaceDetails
// Set of interfaces configured for SRLG
type Srlg_Nodes_Node_InterfaceDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG interface details. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail.
    InterfaceDetail []Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetFilter() yfilter.YFilter { return interfaceDetails.YFilter }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) SetFilter(yf yfilter.YFilter) { interfaceDetails.YFilter = yf }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetGoName(yname string) string {
    if yname == "interface-detail" { return "InterfaceDetail" }
    return ""
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetSegmentPath() string {
    return "interface-details"
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-detail" {
        for _, c := range interfaceDetails.InterfaceDetail {
            if interfaceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail{}
        interfaceDetails.InterfaceDetail = append(interfaceDetails.InterfaceDetail, child)
        return &interfaceDetails.InterfaceDetail[len(interfaceDetails.InterfaceDetail)-1]
    }
    return nil
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceDetails.InterfaceDetail {
        children[interfaceDetails.InterfaceDetail[i].GetSegmentPath()] = &interfaceDetails.InterfaceDetail[i]
    }
    return children
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetYangName() string { return "interface-details" }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) SetParent(parent types.Entity) { interfaceDetails.parent = parent }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetParent() types.Entity { return interfaceDetails.parent }

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail
// SRLG interface details
type Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Groups. The type is interface{} with range: 0..4294967295.
    Groups interface{}

    // Nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // SRLG attributes. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute.
    SrlgAttribute []Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetFilter() yfilter.YFilter { return interfaceDetail.YFilter }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) SetFilter(yf yfilter.YFilter) { interfaceDetail.YFilter = yf }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "groups" { return "Groups" }
    if yname == "nodes" { return "Nodes" }
    if yname == "srlg-attribute" { return "SrlgAttribute" }
    return ""
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetSegmentPath() string {
    return "interface-detail" + "[interface-name='" + fmt.Sprintf("%v", interfaceDetail.InterfaceName) + "']"
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-attribute" {
        for _, c := range interfaceDetail.SrlgAttribute {
            if interfaceDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute{}
        interfaceDetail.SrlgAttribute = append(interfaceDetail.SrlgAttribute, child)
        return &interfaceDetail.SrlgAttribute[len(interfaceDetail.SrlgAttribute)-1]
    }
    return nil
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceDetail.SrlgAttribute {
        children[interfaceDetail.SrlgAttribute[i].GetSegmentPath()] = &interfaceDetail.SrlgAttribute[i]
    }
    return children
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceDetail.InterfaceName
    leafs["groups"] = interfaceDetail.Groups
    leafs["nodes"] = interfaceDetail.Nodes
    return leafs
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetYangName() string { return "interface-detail" }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) SetParent(parent types.Entity) { interfaceDetail.parent = parent }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetParent() types.Entity { return interfaceDetail.parent }

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetParentYangName() string { return "interface-details" }

// Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute
// SRLG attributes
type Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Source. The type is Source.
    Source interface{}

    // Source name. The type is string.
    SourceName interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetFilter() yfilter.YFilter { return srlgAttribute.YFilter }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) SetFilter(yf yfilter.YFilter) { srlgAttribute.YFilter = yf }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetGoName(yname string) string {
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "priority" { return "Priority" }
    if yname == "source" { return "Source" }
    if yname == "source-name" { return "SourceName" }
    if yname == "srlg-index" { return "SrlgIndex" }
    return ""
}

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetSegmentPath() string {
    return "srlg-attribute"
}

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-value"] = srlgAttribute.SrlgValue
    leafs["priority"] = srlgAttribute.Priority
    leafs["source"] = srlgAttribute.Source
    leafs["source-name"] = srlgAttribute.SourceName
    leafs["srlg-index"] = srlgAttribute.SrlgIndex
    return leafs
}

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetBundleName() string { return "cisco_ios_xr" }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetYangName() string { return "srlg-attribute" }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) SetParent(parent types.Entity) { srlgAttribute.parent = parent }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetParent() types.Entity { return srlgAttribute.parent }

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetParentYangName() string { return "interface-detail" }

// Srlg_Nodes_Node_SrlgValues
// Set of SRLG values configured
type Srlg_Nodes_Node_SrlgValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured SRLG value details . The type is slice of
    // Srlg_Nodes_Node_SrlgValues_SrlgValue.
    SrlgValue []Srlg_Nodes_Node_SrlgValues_SrlgValue
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetFilter() yfilter.YFilter { return srlgValues.YFilter }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) SetFilter(yf yfilter.YFilter) { srlgValues.YFilter = yf }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetGoName(yname string) string {
    if yname == "srlg-value" { return "SrlgValue" }
    return ""
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetSegmentPath() string {
    return "srlg-values"
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-value" {
        for _, c := range srlgValues.SrlgValue {
            if srlgValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_SrlgValues_SrlgValue{}
        srlgValues.SrlgValue = append(srlgValues.SrlgValue, child)
        return &srlgValues.SrlgValue[len(srlgValues.SrlgValue)-1]
    }
    return nil
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srlgValues.SrlgValue {
        children[srlgValues.SrlgValue[i].GetSegmentPath()] = &srlgValues.SrlgValue[i]
    }
    return children
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetBundleName() string { return "cisco_ios_xr" }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetYangName() string { return "srlg-values" }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) SetParent(parent types.Entity) { srlgValues.parent = parent }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetParent() types.Entity { return srlgValues.parent }

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_SrlgValues_SrlgValue
// Configured SRLG value details 
type Srlg_Nodes_Node_SrlgValues_SrlgValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG value. The type is interface{} with range:
    // -2147483648..2147483647.
    Value interface{}

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetFilter() yfilter.YFilter { return srlgValue.YFilter }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) SetFilter(yf yfilter.YFilter) { srlgValue.YFilter = yf }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetSegmentPath() string {
    return "srlg-value" + "[value='" + fmt.Sprintf("%v", srlgValue.Value) + "']"
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = srlgValue.Value
    leafs["interface-name"] = srlgValue.InterfaceName
    return leafs
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetBundleName() string { return "cisco_ios_xr" }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetYangName() string { return "srlg-value" }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) SetParent(parent types.Entity) { srlgValue.parent = parent }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetParent() types.Entity { return srlgValue.parent }

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetParentYangName() string { return "srlg-values" }

// Srlg_Nodes_Node_InterfaceSrlgNames
// Set of SRLG names configured
type Srlg_Nodes_Node_InterfaceSrlgNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetFilter() yfilter.YFilter { return interfaceSrlgNames.YFilter }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) SetFilter(yf yfilter.YFilter) { interfaceSrlgNames.YFilter = yf }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetGoName(yname string) string {
    if yname == "interface-srlg-name" { return "InterfaceSrlgName" }
    return ""
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetSegmentPath() string {
    return "interface-srlg-names"
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-srlg-name" {
        for _, c := range interfaceSrlgNames.InterfaceSrlgName {
            if interfaceSrlgNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName{}
        interfaceSrlgNames.InterfaceSrlgName = append(interfaceSrlgNames.InterfaceSrlgName, child)
        return &interfaceSrlgNames.InterfaceSrlgName[len(interfaceSrlgNames.InterfaceSrlgName)-1]
    }
    return nil
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        children[interfaceSrlgNames.InterfaceSrlgName[i].GetSegmentPath()] = &interfaceSrlgNames.InterfaceSrlgName[i]
    }
    return children
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetYangName() string { return "interface-srlg-names" }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) SetParent(parent types.Entity) { interfaceSrlgNames.parent = parent }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetParent() types.Entity { return interfaceSrlgNames.parent }

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetParentYangName() string { return "node" }

// Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName
// Configured SRLG name details 
type Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG name. The type is string.
    SrlgNameXr interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Interfaces information.
    Interfaces Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces
}

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetFilter() yfilter.YFilter { return interfaceSrlgName.YFilter }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) SetFilter(yf yfilter.YFilter) { interfaceSrlgName.YFilter = yf }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    if yname == "srlg-name-xr" { return "SrlgNameXr" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetSegmentPath() string {
    return "interface-srlg-name" + "[srlg-name='" + fmt.Sprintf("%v", interfaceSrlgName.SrlgName) + "']"
}

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceSrlgName.Interfaces
    }
    return nil
}

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceSrlgName.Interfaces
    return children
}

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-name"] = interfaceSrlgName.SrlgName
    leafs["srlg-name-xr"] = interfaceSrlgName.SrlgNameXr
    leafs["srlg-value"] = interfaceSrlgName.SrlgValue
    return leafs
}

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetYangName() string { return "interface-srlg-name" }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) SetParent(parent types.Entity) { interfaceSrlgName.parent = parent }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetParent() types.Entity { return interfaceSrlgName.parent }

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetParentYangName() string { return "interface-srlg-names" }

// Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces
// Interfaces information
type Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaces.InterfaceName
    return leafs
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetParentYangName() string { return "interface-srlg-name" }

// Srlg_InterfaceSrlgNames
// Set of SRLG names configured
type Srlg_InterfaceSrlgNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []Srlg_InterfaceSrlgNames_InterfaceSrlgName
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetFilter() yfilter.YFilter { return interfaceSrlgNames.YFilter }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) SetFilter(yf yfilter.YFilter) { interfaceSrlgNames.YFilter = yf }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetGoName(yname string) string {
    if yname == "interface-srlg-name" { return "InterfaceSrlgName" }
    return ""
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetSegmentPath() string {
    return "interface-srlg-names"
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-srlg-name" {
        for _, c := range interfaceSrlgNames.InterfaceSrlgName {
            if interfaceSrlgNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_InterfaceSrlgNames_InterfaceSrlgName{}
        interfaceSrlgNames.InterfaceSrlgName = append(interfaceSrlgNames.InterfaceSrlgName, child)
        return &interfaceSrlgNames.InterfaceSrlgName[len(interfaceSrlgNames.InterfaceSrlgName)-1]
    }
    return nil
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        children[interfaceSrlgNames.InterfaceSrlgName[i].GetSegmentPath()] = &interfaceSrlgNames.InterfaceSrlgName[i]
    }
    return children
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetYangName() string { return "interface-srlg-names" }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) SetParent(parent types.Entity) { interfaceSrlgNames.parent = parent }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetParent() types.Entity { return interfaceSrlgNames.parent }

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetParentYangName() string { return "srlg" }

// Srlg_InterfaceSrlgNames_InterfaceSrlgName
// Configured SRLG name details 
type Srlg_InterfaceSrlgNames_InterfaceSrlgName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG name. The type is string.
    SrlgNameXr interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Interfaces information.
    Interfaces Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces
}

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetFilter() yfilter.YFilter { return interfaceSrlgName.YFilter }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) SetFilter(yf yfilter.YFilter) { interfaceSrlgName.YFilter = yf }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    if yname == "srlg-name-xr" { return "SrlgNameXr" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetSegmentPath() string {
    return "interface-srlg-name" + "[srlg-name='" + fmt.Sprintf("%v", interfaceSrlgName.SrlgName) + "']"
}

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceSrlgName.Interfaces
    }
    return nil
}

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceSrlgName.Interfaces
    return children
}

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-name"] = interfaceSrlgName.SrlgName
    leafs["srlg-name-xr"] = interfaceSrlgName.SrlgNameXr
    leafs["srlg-value"] = interfaceSrlgName.SrlgValue
    return leafs
}

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetYangName() string { return "interface-srlg-name" }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) SetParent(parent types.Entity) { interfaceSrlgName.parent = parent }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetParent() types.Entity { return interfaceSrlgName.parent }

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetParentYangName() string { return "interface-srlg-names" }

// Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces
// Interfaces information
type Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaces.InterfaceName
    return leafs
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetParentYangName() string { return "interface-srlg-name" }

// SelectiveVrfDownload
// selective vrf download
type SelectiveVrfDownload struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selective VRF Download feature state details.
    State SelectiveVrfDownload_State
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetFilter() yfilter.YFilter { return selectiveVrfDownload.YFilter }

func (selectiveVrfDownload *SelectiveVrfDownload) SetFilter(yf yfilter.YFilter) { selectiveVrfDownload.YFilter = yf }

func (selectiveVrfDownload *SelectiveVrfDownload) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    return ""
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-oper:selective-vrf-download"
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &selectiveVrfDownload.State
    }
    return nil
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &selectiveVrfDownload.State
    return children
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetBundleName() string { return "cisco_ios_xr" }

func (selectiveVrfDownload *SelectiveVrfDownload) GetYangName() string { return "selective-vrf-download" }

func (selectiveVrfDownload *SelectiveVrfDownload) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectiveVrfDownload *SelectiveVrfDownload) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectiveVrfDownload *SelectiveVrfDownload) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectiveVrfDownload *SelectiveVrfDownload) SetParent(parent types.Entity) { selectiveVrfDownload.parent = parent }

func (selectiveVrfDownload *SelectiveVrfDownload) GetParent() types.Entity { return selectiveVrfDownload.parent }

func (selectiveVrfDownload *SelectiveVrfDownload) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-oper" }

// SelectiveVrfDownload_State
// Selective VRF Download feature state details
type SelectiveVrfDownload_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SVD Enabled Operational. The type is bool.
    IsSvdEnabled interface{}

    // Is SVD Enabled Config. The type is bool.
    IsSvdEnabledCfg interface{}
}

func (state *SelectiveVrfDownload_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *SelectiveVrfDownload_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *SelectiveVrfDownload_State) GetGoName(yname string) string {
    if yname == "is-svd-enabled" { return "IsSvdEnabled" }
    if yname == "is-svd-enabled-cfg" { return "IsSvdEnabledCfg" }
    return ""
}

func (state *SelectiveVrfDownload_State) GetSegmentPath() string {
    return "state"
}

func (state *SelectiveVrfDownload_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *SelectiveVrfDownload_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *SelectiveVrfDownload_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-svd-enabled"] = state.IsSvdEnabled
    leafs["is-svd-enabled-cfg"] = state.IsSvdEnabledCfg
    return leafs
}

func (state *SelectiveVrfDownload_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *SelectiveVrfDownload_State) GetYangName() string { return "state" }

func (state *SelectiveVrfDownload_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *SelectiveVrfDownload_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *SelectiveVrfDownload_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *SelectiveVrfDownload_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *SelectiveVrfDownload_State) GetParent() types.Entity { return state.parent }

func (state *SelectiveVrfDownload_State) GetParentYangName() string { return "selective-vrf-download" }

