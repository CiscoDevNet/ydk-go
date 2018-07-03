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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node (Physical location of the node in R_S_I format). The type is slice of
    // NodePath_Node.
    Node []*NodePath_Node
}

func (nodePath *NodePath) GetEntityData() *types.CommonEntityData {
    nodePath.EntityData.YFilter = nodePath.YFilter
    nodePath.EntityData.YangName = "node-path"
    nodePath.EntityData.BundleName = "cisco_ios_xr"
    nodePath.EntityData.ParentYangName = "Cisco-IOS-XR-slice-mgr-proxy-cfg"
    nodePath.EntityData.SegmentPath = "Cisco-IOS-XR-slice-mgr-proxy-cfg:node-path"
    nodePath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodePath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodePath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodePath.EntityData.Children = types.NewOrderedMap()
    nodePath.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodePath.Node {
        nodePath.EntityData.Children.Append(types.GetSegmentPath(nodePath.Node[i]), types.YChild{"Node", nodePath.Node[i]})
    }
    nodePath.EntityData.Leafs = types.NewOrderedMap()

    nodePath.EntityData.YListKeys = []string {}

    return &(nodePath.EntityData)
}

// NodePath_Node
// Node (Physical location of the node in R_S_I
// format)
type NodePath_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location in R_S_I format. The type is string.
    NodeName interface{}

    // Slice.
    SliceIds NodePath_Node_SliceIds
}

func (node *NodePath_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "node-path"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("slice-ids", types.YChild{"SliceIds", &node.SliceIds})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// NodePath_Node_SliceIds
// Slice
type NodePath_Node_SliceIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slice Id on which configuration will be applied. The type is slice of
    // NodePath_Node_SliceIds_SliceId.
    SliceId []*NodePath_Node_SliceIds_SliceId
}

func (sliceIds *NodePath_Node_SliceIds) GetEntityData() *types.CommonEntityData {
    sliceIds.EntityData.YFilter = sliceIds.YFilter
    sliceIds.EntityData.YangName = "slice-ids"
    sliceIds.EntityData.BundleName = "cisco_ios_xr"
    sliceIds.EntityData.ParentYangName = "node"
    sliceIds.EntityData.SegmentPath = "slice-ids"
    sliceIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceIds.EntityData.Children = types.NewOrderedMap()
    sliceIds.EntityData.Children.Append("slice-id", types.YChild{"SliceId", nil})
    for i := range sliceIds.SliceId {
        sliceIds.EntityData.Children.Append(types.GetSegmentPath(sliceIds.SliceId[i]), types.YChild{"SliceId", sliceIds.SliceId[i]})
    }
    sliceIds.EntityData.Leafs = types.NewOrderedMap()

    sliceIds.EntityData.YListKeys = []string {}

    return &(sliceIds.EntityData)
}

// NodePath_Node_SliceIds_SliceId
// Slice Id on which configuration will be
// applied
type NodePath_Node_SliceIds_SliceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this slice. The type is
    // interface{} with range: 0..4.
    SliceId interface{}

    // set val 0 to shutdown. The type is interface{} with range: 0..4294967295.
    State interface{}

    // 10G Breakout Config. The type is interface{} with range: 0..4294967295.
    Breakout interface{}

    // set val 4 for OTU4 . The type is interface{} with range: 0..4294967295.
    Mode interface{}
}

func (sliceId *NodePath_Node_SliceIds_SliceId) GetEntityData() *types.CommonEntityData {
    sliceId.EntityData.YFilter = sliceId.YFilter
    sliceId.EntityData.YangName = "slice-id"
    sliceId.EntityData.BundleName = "cisco_ios_xr"
    sliceId.EntityData.ParentYangName = "slice-ids"
    sliceId.EntityData.SegmentPath = "slice-id" + types.AddKeyToken(sliceId.SliceId, "slice-id")
    sliceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceId.EntityData.Children = types.NewOrderedMap()
    sliceId.EntityData.Leafs = types.NewOrderedMap()
    sliceId.EntityData.Leafs.Append("slice-id", types.YLeaf{"SliceId", sliceId.SliceId})
    sliceId.EntityData.Leafs.Append("state", types.YLeaf{"State", sliceId.State})
    sliceId.EntityData.Leafs.Append("breakout", types.YLeaf{"Breakout", sliceId.Breakout})
    sliceId.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", sliceId.Mode})

    sliceId.EntityData.YListKeys = []string {"SliceId"}

    return &(sliceId.EntityData)
}

