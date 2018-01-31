// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-mlan-cmp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mlan-disable-cmp: Disable/Enable CMP
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_mlan_cmp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_mlan_cmp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-mlan-cmp-cfg mlan-disable-cmp}", reflect.TypeOf(MlanDisableCmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-mlan-cmp-cfg:mlan-disable-cmp", reflect.TypeOf(MlanDisableCmp{}))
}

// MlanDisableCmp
// Disable/Enable CMP
type MlanDisableCmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fully qualified RSP4/RSP4s/RP2 card specification.
    Nodes MlanDisableCmp_Nodes
}

func (mlanDisableCmp *MlanDisableCmp) GetFilter() yfilter.YFilter { return mlanDisableCmp.YFilter }

func (mlanDisableCmp *MlanDisableCmp) SetFilter(yf yfilter.YFilter) { mlanDisableCmp.YFilter = yf }

func (mlanDisableCmp *MlanDisableCmp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (mlanDisableCmp *MlanDisableCmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-mlan-cmp-cfg:mlan-disable-cmp"
}

func (mlanDisableCmp *MlanDisableCmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &mlanDisableCmp.Nodes
    }
    return nil
}

func (mlanDisableCmp *MlanDisableCmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &mlanDisableCmp.Nodes
    return children
}

func (mlanDisableCmp *MlanDisableCmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mlanDisableCmp *MlanDisableCmp) GetBundleName() string { return "cisco_ios_xr" }

func (mlanDisableCmp *MlanDisableCmp) GetYangName() string { return "mlan-disable-cmp" }

func (mlanDisableCmp *MlanDisableCmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mlanDisableCmp *MlanDisableCmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mlanDisableCmp *MlanDisableCmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mlanDisableCmp *MlanDisableCmp) SetParent(parent types.Entity) { mlanDisableCmp.parent = parent }

func (mlanDisableCmp *MlanDisableCmp) GetParent() types.Entity { return mlanDisableCmp.parent }

func (mlanDisableCmp *MlanDisableCmp) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-mlan-cmp-cfg" }

// MlanDisableCmp_Nodes
// Fully qualified RSP4/RSP4s/RP2 card
// specification
type MlanDisableCmp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A node. The type is slice of MlanDisableCmp_Nodes_Node.
    Node []MlanDisableCmp_Nodes_Node
}

func (nodes *MlanDisableCmp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *MlanDisableCmp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *MlanDisableCmp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *MlanDisableCmp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *MlanDisableCmp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MlanDisableCmp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *MlanDisableCmp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *MlanDisableCmp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *MlanDisableCmp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *MlanDisableCmp_Nodes) GetYangName() string { return "nodes" }

func (nodes *MlanDisableCmp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *MlanDisableCmp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *MlanDisableCmp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *MlanDisableCmp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *MlanDisableCmp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *MlanDisableCmp_Nodes) GetParentYangName() string { return "mlan-disable-cmp" }

// MlanDisableCmp_Nodes_Node
// A node
type MlanDisableCmp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NodeName. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Disable CMP. The type is interface{}.
    Disable interface{}
}

func (node *MlanDisableCmp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *MlanDisableCmp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *MlanDisableCmp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "disable" { return "Disable" }
    return ""
}

func (node *MlanDisableCmp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *MlanDisableCmp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *MlanDisableCmp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *MlanDisableCmp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["disable"] = node.Disable
    return leafs
}

func (node *MlanDisableCmp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *MlanDisableCmp_Nodes_Node) GetYangName() string { return "node" }

func (node *MlanDisableCmp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *MlanDisableCmp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *MlanDisableCmp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *MlanDisableCmp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *MlanDisableCmp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *MlanDisableCmp_Nodes_Node) GetParentYangName() string { return "nodes" }

