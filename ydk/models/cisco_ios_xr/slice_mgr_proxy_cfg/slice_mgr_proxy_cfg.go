// This module contains a collection of YANG definitions
// for Cisco IOS-XR slice-mgr-proxy package configuration.
// 
// This module contains definitions
// for the following management objects:
//   node-path: Node act path
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package slice_mgr_proxy_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package slice_mgr_proxy_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-slice-mgr-proxy-cfg node-path}", reflect.TypeOf(NodePath{}))
    ydk.RegisterEntity("Cisco-IOS-XR-slice-mgr-proxy-cfg:node-path", reflect.TypeOf(NodePath{}))
}

// NodePath
// Node act path
type NodePath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node (Physical location of the node in R_S_I format). The type is slice of
    // NodePath_Node.
    Node []NodePath_Node
}

func (nodePath *NodePath) GetFilter() yfilter.YFilter { return nodePath.YFilter }

func (nodePath *NodePath) SetFilter(yf yfilter.YFilter) { nodePath.YFilter = yf }

func (nodePath *NodePath) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodePath *NodePath) GetSegmentPath() string {
    return "Cisco-IOS-XR-slice-mgr-proxy-cfg:node-path"
}

func (nodePath *NodePath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodePath.Node {
            if nodePath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NodePath_Node{}
        nodePath.Node = append(nodePath.Node, child)
        return &nodePath.Node[len(nodePath.Node)-1]
    }
    return nil
}

func (nodePath *NodePath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodePath.Node {
        children[nodePath.Node[i].GetSegmentPath()] = &nodePath.Node[i]
    }
    return children
}

func (nodePath *NodePath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodePath *NodePath) GetBundleName() string { return "cisco_ios_xr" }

func (nodePath *NodePath) GetYangName() string { return "node-path" }

func (nodePath *NodePath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodePath *NodePath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodePath *NodePath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodePath *NodePath) SetParent(parent types.Entity) { nodePath.parent = parent }

func (nodePath *NodePath) GetParent() types.Entity { return nodePath.parent }

func (nodePath *NodePath) GetParentYangName() string { return "Cisco-IOS-XR-slice-mgr-proxy-cfg" }

// NodePath_Node
// Node (Physical location of the node in R_S_I
// format)
type NodePath_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location in R_S_I format. The type is string.
    NodeName interface{}

    // Slice.
    SliceIds NodePath_Node_SliceIds
}

func (node *NodePath_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *NodePath_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *NodePath_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "slice-ids" { return "SliceIds" }
    return ""
}

func (node *NodePath_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *NodePath_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-ids" {
        return &node.SliceIds
    }
    return nil
}

func (node *NodePath_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slice-ids"] = &node.SliceIds
    return children
}

func (node *NodePath_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *NodePath_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *NodePath_Node) GetYangName() string { return "node" }

func (node *NodePath_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *NodePath_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *NodePath_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *NodePath_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *NodePath_Node) GetParent() types.Entity { return node.parent }

func (node *NodePath_Node) GetParentYangName() string { return "node-path" }

// NodePath_Node_SliceIds
// Slice
type NodePath_Node_SliceIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slice Id on which configuration will be applied. The type is slice of
    // NodePath_Node_SliceIds_SliceId.
    SliceId []NodePath_Node_SliceIds_SliceId
}

func (sliceIds *NodePath_Node_SliceIds) GetFilter() yfilter.YFilter { return sliceIds.YFilter }

func (sliceIds *NodePath_Node_SliceIds) SetFilter(yf yfilter.YFilter) { sliceIds.YFilter = yf }

func (sliceIds *NodePath_Node_SliceIds) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    return ""
}

func (sliceIds *NodePath_Node_SliceIds) GetSegmentPath() string {
    return "slice-ids"
}

func (sliceIds *NodePath_Node_SliceIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-id" {
        for _, c := range sliceIds.SliceId {
            if sliceIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NodePath_Node_SliceIds_SliceId{}
        sliceIds.SliceId = append(sliceIds.SliceId, child)
        return &sliceIds.SliceId[len(sliceIds.SliceId)-1]
    }
    return nil
}

func (sliceIds *NodePath_Node_SliceIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceIds.SliceId {
        children[sliceIds.SliceId[i].GetSegmentPath()] = &sliceIds.SliceId[i]
    }
    return children
}

func (sliceIds *NodePath_Node_SliceIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sliceIds *NodePath_Node_SliceIds) GetBundleName() string { return "cisco_ios_xr" }

func (sliceIds *NodePath_Node_SliceIds) GetYangName() string { return "slice-ids" }

func (sliceIds *NodePath_Node_SliceIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceIds *NodePath_Node_SliceIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceIds *NodePath_Node_SliceIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceIds *NodePath_Node_SliceIds) SetParent(parent types.Entity) { sliceIds.parent = parent }

func (sliceIds *NodePath_Node_SliceIds) GetParent() types.Entity { return sliceIds.parent }

func (sliceIds *NodePath_Node_SliceIds) GetParentYangName() string { return "node" }

// NodePath_Node_SliceIds_SliceId
// Slice Id on which configuration will be
// applied
type NodePath_Node_SliceIds_SliceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this slice. The type is
    // interface{} with range: 0..4.
    SliceId interface{}

    // set val 0 to shutdown. The type is interface{} with range:
    // -2147483648..2147483647.
    State interface{}

    // 10G Breakout Config. The type is interface{} with range:
    // -2147483648..2147483647.
    Breakout interface{}

    // set val 4 for OTU4 . The type is interface{} with range:
    // -2147483648..2147483647.
    Mode interface{}
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetFilter() yfilter.YFilter { return sliceId.YFilter }

func (sliceId *NodePath_Node_SliceIds_SliceId) SetFilter(yf yfilter.YFilter) { sliceId.YFilter = yf }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    if yname == "state" { return "State" }
    if yname == "breakout" { return "Breakout" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetSegmentPath() string {
    return "slice-id" + "[slice-id='" + fmt.Sprintf("%v", sliceId.SliceId) + "']"
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-id"] = sliceId.SliceId
    leafs["state"] = sliceId.State
    leafs["breakout"] = sliceId.Breakout
    leafs["mode"] = sliceId.Mode
    return leafs
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetBundleName() string { return "cisco_ios_xr" }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetYangName() string { return "slice-id" }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceId *NodePath_Node_SliceIds_SliceId) SetParent(parent types.Entity) { sliceId.parent = parent }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetParent() types.Entity { return sliceId.parent }

func (sliceId *NodePath_Node_SliceIds_SliceId) GetParentYangName() string { return "slice-ids" }

