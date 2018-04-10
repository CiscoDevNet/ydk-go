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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fully qualified RSP4/RSP4s/RP2 card specification.
    Nodes MlanDisableCmp_Nodes
}

func (mlanDisableCmp *MlanDisableCmp) GetEntityData() *types.CommonEntityData {
    mlanDisableCmp.EntityData.YFilter = mlanDisableCmp.YFilter
    mlanDisableCmp.EntityData.YangName = "mlan-disable-cmp"
    mlanDisableCmp.EntityData.BundleName = "cisco_ios_xr"
    mlanDisableCmp.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-mlan-cmp-cfg"
    mlanDisableCmp.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-mlan-cmp-cfg:mlan-disable-cmp"
    mlanDisableCmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mlanDisableCmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mlanDisableCmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mlanDisableCmp.EntityData.Children = make(map[string]types.YChild)
    mlanDisableCmp.EntityData.Children["nodes"] = types.YChild{"Nodes", &mlanDisableCmp.Nodes}
    mlanDisableCmp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mlanDisableCmp.EntityData)
}

// MlanDisableCmp_Nodes
// Fully qualified RSP4/RSP4s/RP2 card
// specification
type MlanDisableCmp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A node. The type is slice of MlanDisableCmp_Nodes_Node.
    Node []MlanDisableCmp_Nodes_Node
}

func (nodes *MlanDisableCmp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "mlan-disable-cmp"
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

// MlanDisableCmp_Nodes_Node
// A node
type MlanDisableCmp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NodeName. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Disable CMP. The type is interface{}.
    Disable interface{}
}

func (node *MlanDisableCmp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    node.EntityData.Leafs["disable"] = types.YLeaf{"Disable", node.Disable}
    return &(node.EntityData)
}

