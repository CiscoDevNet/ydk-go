// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rmf package operational data.
// 
// This module contains definitions
// for the following management objects:
//   redundancy: Redundancy show information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_rmf_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rmf_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rmf-oper redundancy}", reflect.TypeOf(Redundancy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rmf-oper:redundancy", reflect.TypeOf(Redundancy{}))
}

// Redundancy
// Redundancy show information
type Redundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location show information.
    Nodes Redundancy_Nodes

    // Redundancy Summary of Nodes.
    Summary Redundancy_Summary
}

func (redundancy *Redundancy) GetFilter() yfilter.YFilter { return redundancy.YFilter }

func (redundancy *Redundancy) SetFilter(yf yfilter.YFilter) { redundancy.YFilter = yf }

func (redundancy *Redundancy) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (redundancy *Redundancy) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rmf-oper:redundancy"
}

func (redundancy *Redundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &redundancy.Nodes
    }
    if childYangName == "summary" {
        return &redundancy.Summary
    }
    return nil
}

func (redundancy *Redundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &redundancy.Nodes
    children["summary"] = &redundancy.Summary
    return children
}

func (redundancy *Redundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redundancy *Redundancy) GetBundleName() string { return "cisco_ios_xr" }

func (redundancy *Redundancy) GetYangName() string { return "redundancy" }

func (redundancy *Redundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancy *Redundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancy *Redundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancy *Redundancy) SetParent(parent types.Entity) { redundancy.parent = parent }

func (redundancy *Redundancy) GetParent() types.Entity { return redundancy.parent }

func (redundancy *Redundancy) GetParentYangName() string { return "Cisco-IOS-XR-infra-rmf-oper" }

// Redundancy_Nodes
// Location show information
type Redundancy_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy Node Information. The type is slice of Redundancy_Nodes_Node.
    Node []Redundancy_Nodes_Node
}

func (nodes *Redundancy_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Redundancy_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Redundancy_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Redundancy_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Redundancy_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Redundancy_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Redundancy_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Redundancy_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Redundancy_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Redundancy_Nodes) GetYangName() string { return "nodes" }

func (nodes *Redundancy_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Redundancy_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Redundancy_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Redundancy_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Redundancy_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Redundancy_Nodes) GetParentYangName() string { return "redundancy" }

// Redundancy_Nodes_Node
// Redundancy Node Information
type Redundancy_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Reload and boot logs. The type is string.
    Log interface{}

    // Active node reload. The type is string.
    ActiveRebootReason interface{}

    // Standby node reload. The type is string.
    StandbyRebootReason interface{}

    // Error Log. The type is string.
    ErrLog interface{}

    // Row information.
    Redundancy Redundancy_Nodes_Node_Redundancy
}

func (node *Redundancy_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Redundancy_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Redundancy_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "log" { return "Log" }
    if yname == "active-reboot-reason" { return "ActiveRebootReason" }
    if yname == "standby-reboot-reason" { return "StandbyRebootReason" }
    if yname == "err-log" { return "ErrLog" }
    if yname == "redundancy" { return "Redundancy" }
    return ""
}

func (node *Redundancy_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *Redundancy_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redundancy" {
        return &node.Redundancy
    }
    return nil
}

func (node *Redundancy_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["redundancy"] = &node.Redundancy
    return children
}

func (node *Redundancy_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["log"] = node.Log
    leafs["active-reboot-reason"] = node.ActiveRebootReason
    leafs["standby-reboot-reason"] = node.StandbyRebootReason
    leafs["err-log"] = node.ErrLog
    return leafs
}

func (node *Redundancy_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Redundancy_Nodes_Node) GetYangName() string { return "node" }

func (node *Redundancy_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Redundancy_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Redundancy_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Redundancy_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Redundancy_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Redundancy_Nodes_Node) GetParentYangName() string { return "nodes" }

// Redundancy_Nodes_Node_Redundancy
// Row information
type Redundancy_Nodes_Node_Redundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active node name R/S/I. The type is string.
    Active interface{}

    // Standby node name R/S/I. The type is string.
    Standby interface{}

    // High Availability state Ready/Not Ready. The type is string.
    HaState interface{}

    // NSR state Configured/Not Configured. The type is string.
    NsrState interface{}

    // groupinfo. The type is slice of Redundancy_Nodes_Node_Redundancy_Groupinfo.
    Groupinfo []Redundancy_Nodes_Node_Redundancy_Groupinfo
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetFilter() yfilter.YFilter { return redundancy.YFilter }

func (redundancy *Redundancy_Nodes_Node_Redundancy) SetFilter(yf yfilter.YFilter) { redundancy.YFilter = yf }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "standby" { return "Standby" }
    if yname == "ha-state" { return "HaState" }
    if yname == "nsr-state" { return "NsrState" }
    if yname == "groupinfo" { return "Groupinfo" }
    return ""
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetSegmentPath() string {
    return "redundancy"
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groupinfo" {
        for _, c := range redundancy.Groupinfo {
            if redundancy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Redundancy_Nodes_Node_Redundancy_Groupinfo{}
        redundancy.Groupinfo = append(redundancy.Groupinfo, child)
        return &redundancy.Groupinfo[len(redundancy.Groupinfo)-1]
    }
    return nil
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redundancy.Groupinfo {
        children[redundancy.Groupinfo[i].GetSegmentPath()] = &redundancy.Groupinfo[i]
    }
    return children
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = redundancy.Active
    leafs["standby"] = redundancy.Standby
    leafs["ha-state"] = redundancy.HaState
    leafs["nsr-state"] = redundancy.NsrState
    return leafs
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetBundleName() string { return "cisco_ios_xr" }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetYangName() string { return "redundancy" }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancy *Redundancy_Nodes_Node_Redundancy) SetParent(parent types.Entity) { redundancy.parent = parent }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetParent() types.Entity { return redundancy.parent }

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetParentYangName() string { return "node" }

// Redundancy_Nodes_Node_Redundancy_Groupinfo
// groupinfo
type Redundancy_Nodes_Node_Redundancy_Groupinfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active. The type is string.
    Active interface{}

    // Standby. The type is string.
    Standby interface{}

    // HAState. The type is string.
    HaState interface{}

    // NSRState. The type is string.
    NsrState interface{}
}

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetFilter() yfilter.YFilter { return groupinfo.YFilter }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) SetFilter(yf yfilter.YFilter) { groupinfo.YFilter = yf }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "standby" { return "Standby" }
    if yname == "ha-state" { return "HaState" }
    if yname == "nsr-state" { return "NsrState" }
    return ""
}

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetSegmentPath() string {
    return "groupinfo"
}

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = groupinfo.Active
    leafs["standby"] = groupinfo.Standby
    leafs["ha-state"] = groupinfo.HaState
    leafs["nsr-state"] = groupinfo.NsrState
    return leafs
}

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetBundleName() string { return "cisco_ios_xr" }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetYangName() string { return "groupinfo" }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) SetParent(parent types.Entity) { groupinfo.parent = parent }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetParent() types.Entity { return groupinfo.parent }

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetParentYangName() string { return "redundancy" }

// Redundancy_Summary
// Redundancy Summary of Nodes
type Redundancy_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error Log. The type is string.
    ErrLog interface{}

    // Redundancy Pair. The type is slice of Redundancy_Summary_RedPair.
    RedPair []Redundancy_Summary_RedPair
}

func (summary *Redundancy_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Redundancy_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Redundancy_Summary) GetGoName(yname string) string {
    if yname == "err-log" { return "ErrLog" }
    if yname == "red-pair" { return "RedPair" }
    return ""
}

func (summary *Redundancy_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Redundancy_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "red-pair" {
        for _, c := range summary.RedPair {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Redundancy_Summary_RedPair{}
        summary.RedPair = append(summary.RedPair, child)
        return &summary.RedPair[len(summary.RedPair)-1]
    }
    return nil
}

func (summary *Redundancy_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.RedPair {
        children[summary.RedPair[i].GetSegmentPath()] = &summary.RedPair[i]
    }
    return children
}

func (summary *Redundancy_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["err-log"] = summary.ErrLog
    return leafs
}

func (summary *Redundancy_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Redundancy_Summary) GetYangName() string { return "summary" }

func (summary *Redundancy_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Redundancy_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Redundancy_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Redundancy_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Redundancy_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Redundancy_Summary) GetParentYangName() string { return "redundancy" }

// Redundancy_Summary_RedPair
// Redundancy Pair
type Redundancy_Summary_RedPair struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active node name R/S/I. The type is string.
    Active interface{}

    // Standby node name R/S/I. The type is string.
    Standby interface{}

    // High Availability state Ready/Not Ready. The type is string.
    HaState interface{}

    // NSR state Configured/Not Configured. The type is string.
    NsrState interface{}

    // groupinfo. The type is slice of Redundancy_Summary_RedPair_Groupinfo.
    Groupinfo []Redundancy_Summary_RedPair_Groupinfo
}

func (redPair *Redundancy_Summary_RedPair) GetFilter() yfilter.YFilter { return redPair.YFilter }

func (redPair *Redundancy_Summary_RedPair) SetFilter(yf yfilter.YFilter) { redPair.YFilter = yf }

func (redPair *Redundancy_Summary_RedPair) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "standby" { return "Standby" }
    if yname == "ha-state" { return "HaState" }
    if yname == "nsr-state" { return "NsrState" }
    if yname == "groupinfo" { return "Groupinfo" }
    return ""
}

func (redPair *Redundancy_Summary_RedPair) GetSegmentPath() string {
    return "red-pair"
}

func (redPair *Redundancy_Summary_RedPair) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groupinfo" {
        for _, c := range redPair.Groupinfo {
            if redPair.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Redundancy_Summary_RedPair_Groupinfo{}
        redPair.Groupinfo = append(redPair.Groupinfo, child)
        return &redPair.Groupinfo[len(redPair.Groupinfo)-1]
    }
    return nil
}

func (redPair *Redundancy_Summary_RedPair) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redPair.Groupinfo {
        children[redPair.Groupinfo[i].GetSegmentPath()] = &redPair.Groupinfo[i]
    }
    return children
}

func (redPair *Redundancy_Summary_RedPair) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = redPair.Active
    leafs["standby"] = redPair.Standby
    leafs["ha-state"] = redPair.HaState
    leafs["nsr-state"] = redPair.NsrState
    return leafs
}

func (redPair *Redundancy_Summary_RedPair) GetBundleName() string { return "cisco_ios_xr" }

func (redPair *Redundancy_Summary_RedPair) GetYangName() string { return "red-pair" }

func (redPair *Redundancy_Summary_RedPair) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redPair *Redundancy_Summary_RedPair) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redPair *Redundancy_Summary_RedPair) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redPair *Redundancy_Summary_RedPair) SetParent(parent types.Entity) { redPair.parent = parent }

func (redPair *Redundancy_Summary_RedPair) GetParent() types.Entity { return redPair.parent }

func (redPair *Redundancy_Summary_RedPair) GetParentYangName() string { return "summary" }

// Redundancy_Summary_RedPair_Groupinfo
// groupinfo
type Redundancy_Summary_RedPair_Groupinfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active. The type is string.
    Active interface{}

    // Standby. The type is string.
    Standby interface{}

    // HAState. The type is string.
    HaState interface{}

    // NSRState. The type is string.
    NsrState interface{}
}

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetFilter() yfilter.YFilter { return groupinfo.YFilter }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) SetFilter(yf yfilter.YFilter) { groupinfo.YFilter = yf }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "standby" { return "Standby" }
    if yname == "ha-state" { return "HaState" }
    if yname == "nsr-state" { return "NsrState" }
    return ""
}

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetSegmentPath() string {
    return "groupinfo"
}

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = groupinfo.Active
    leafs["standby"] = groupinfo.Standby
    leafs["ha-state"] = groupinfo.HaState
    leafs["nsr-state"] = groupinfo.NsrState
    return leafs
}

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetBundleName() string { return "cisco_ios_xr" }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetYangName() string { return "groupinfo" }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) SetParent(parent types.Entity) { groupinfo.parent = parent }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetParent() types.Entity { return groupinfo.parent }

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetParentYangName() string { return "red-pair" }

