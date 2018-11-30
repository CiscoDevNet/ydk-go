// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rmf package operational data.
// 
// This module contains definitions
// for the following management objects:
//   redundancy: Redundancy show information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location show information.
    Nodes Redundancy_Nodes

    // Redundancy Summary of Nodes.
    Summary Redundancy_Summary
}

func (redundancy *Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rmf-oper"
    redundancy.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rmf-oper:redundancy"
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = types.NewOrderedMap()
    redundancy.EntityData.Children.Append("nodes", types.YChild{"Nodes", &redundancy.Nodes})
    redundancy.EntityData.Children.Append("summary", types.YChild{"Summary", &redundancy.Summary})
    redundancy.EntityData.Leafs = types.NewOrderedMap()

    redundancy.EntityData.YListKeys = []string {}

    return &(redundancy.EntityData)
}

// Redundancy_Nodes
// Location show information
type Redundancy_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Node Information. The type is slice of Redundancy_Nodes_Node.
    Node []*Redundancy_Nodes_Node
}

func (nodes *Redundancy_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "redundancy"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Redundancy_Nodes_Node
// Redundancy Node Information
type Redundancy_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Redundancy_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("redundancy", types.YChild{"Redundancy", &node.Redundancy})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("log", types.YLeaf{"Log", node.Log})
    node.EntityData.Leafs.Append("active-reboot-reason", types.YLeaf{"ActiveRebootReason", node.ActiveRebootReason})
    node.EntityData.Leafs.Append("standby-reboot-reason", types.YLeaf{"StandbyRebootReason", node.StandbyRebootReason})
    node.EntityData.Leafs.Append("err-log", types.YLeaf{"ErrLog", node.ErrLog})

    node.EntityData.YListKeys = []string {"NodeId"}

    return &(node.EntityData)
}

// Redundancy_Nodes_Node_Redundancy
// Row information
type Redundancy_Nodes_Node_Redundancy struct {
    EntityData types.CommonEntityData
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
    Groupinfo []*Redundancy_Nodes_Node_Redundancy_Groupinfo
}

func (redundancy *Redundancy_Nodes_Node_Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "node"
    redundancy.EntityData.SegmentPath = "redundancy"
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = types.NewOrderedMap()
    redundancy.EntityData.Children.Append("groupinfo", types.YChild{"Groupinfo", nil})
    for i := range redundancy.Groupinfo {
        redundancy.EntityData.Children.Append(types.GetSegmentPath(redundancy.Groupinfo[i]), types.YChild{"Groupinfo", redundancy.Groupinfo[i]})
    }
    redundancy.EntityData.Leafs = types.NewOrderedMap()
    redundancy.EntityData.Leafs.Append("active", types.YLeaf{"Active", redundancy.Active})
    redundancy.EntityData.Leafs.Append("standby", types.YLeaf{"Standby", redundancy.Standby})
    redundancy.EntityData.Leafs.Append("ha-state", types.YLeaf{"HaState", redundancy.HaState})
    redundancy.EntityData.Leafs.Append("nsr-state", types.YLeaf{"NsrState", redundancy.NsrState})

    redundancy.EntityData.YListKeys = []string {}

    return &(redundancy.EntityData)
}

// Redundancy_Nodes_Node_Redundancy_Groupinfo
// groupinfo
type Redundancy_Nodes_Node_Redundancy_Groupinfo struct {
    EntityData types.CommonEntityData
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

func (groupinfo *Redundancy_Nodes_Node_Redundancy_Groupinfo) GetEntityData() *types.CommonEntityData {
    groupinfo.EntityData.YFilter = groupinfo.YFilter
    groupinfo.EntityData.YangName = "groupinfo"
    groupinfo.EntityData.BundleName = "cisco_ios_xr"
    groupinfo.EntityData.ParentYangName = "redundancy"
    groupinfo.EntityData.SegmentPath = "groupinfo"
    groupinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupinfo.EntityData.Children = types.NewOrderedMap()
    groupinfo.EntityData.Leafs = types.NewOrderedMap()
    groupinfo.EntityData.Leafs.Append("active", types.YLeaf{"Active", groupinfo.Active})
    groupinfo.EntityData.Leafs.Append("standby", types.YLeaf{"Standby", groupinfo.Standby})
    groupinfo.EntityData.Leafs.Append("ha-state", types.YLeaf{"HaState", groupinfo.HaState})
    groupinfo.EntityData.Leafs.Append("nsr-state", types.YLeaf{"NsrState", groupinfo.NsrState})

    groupinfo.EntityData.YListKeys = []string {}

    return &(groupinfo.EntityData)
}

// Redundancy_Summary
// Redundancy Summary of Nodes
type Redundancy_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error Log. The type is string.
    ErrLog interface{}

    // Redundancy Pair. The type is slice of Redundancy_Summary_RedPair.
    RedPair []*Redundancy_Summary_RedPair
}

func (summary *Redundancy_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "redundancy"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("red-pair", types.YChild{"RedPair", nil})
    for i := range summary.RedPair {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.RedPair[i]), types.YChild{"RedPair", summary.RedPair[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("err-log", types.YLeaf{"ErrLog", summary.ErrLog})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Redundancy_Summary_RedPair
// Redundancy Pair
type Redundancy_Summary_RedPair struct {
    EntityData types.CommonEntityData
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
    Groupinfo []*Redundancy_Summary_RedPair_Groupinfo
}

func (redPair *Redundancy_Summary_RedPair) GetEntityData() *types.CommonEntityData {
    redPair.EntityData.YFilter = redPair.YFilter
    redPair.EntityData.YangName = "red-pair"
    redPair.EntityData.BundleName = "cisco_ios_xr"
    redPair.EntityData.ParentYangName = "summary"
    redPair.EntityData.SegmentPath = "red-pair"
    redPair.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redPair.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redPair.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redPair.EntityData.Children = types.NewOrderedMap()
    redPair.EntityData.Children.Append("groupinfo", types.YChild{"Groupinfo", nil})
    for i := range redPair.Groupinfo {
        redPair.EntityData.Children.Append(types.GetSegmentPath(redPair.Groupinfo[i]), types.YChild{"Groupinfo", redPair.Groupinfo[i]})
    }
    redPair.EntityData.Leafs = types.NewOrderedMap()
    redPair.EntityData.Leafs.Append("active", types.YLeaf{"Active", redPair.Active})
    redPair.EntityData.Leafs.Append("standby", types.YLeaf{"Standby", redPair.Standby})
    redPair.EntityData.Leafs.Append("ha-state", types.YLeaf{"HaState", redPair.HaState})
    redPair.EntityData.Leafs.Append("nsr-state", types.YLeaf{"NsrState", redPair.NsrState})

    redPair.EntityData.YListKeys = []string {}

    return &(redPair.EntityData)
}

// Redundancy_Summary_RedPair_Groupinfo
// groupinfo
type Redundancy_Summary_RedPair_Groupinfo struct {
    EntityData types.CommonEntityData
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

func (groupinfo *Redundancy_Summary_RedPair_Groupinfo) GetEntityData() *types.CommonEntityData {
    groupinfo.EntityData.YFilter = groupinfo.YFilter
    groupinfo.EntityData.YangName = "groupinfo"
    groupinfo.EntityData.BundleName = "cisco_ios_xr"
    groupinfo.EntityData.ParentYangName = "red-pair"
    groupinfo.EntityData.SegmentPath = "groupinfo"
    groupinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupinfo.EntityData.Children = types.NewOrderedMap()
    groupinfo.EntityData.Leafs = types.NewOrderedMap()
    groupinfo.EntityData.Leafs.Append("active", types.YLeaf{"Active", groupinfo.Active})
    groupinfo.EntityData.Leafs.Append("standby", types.YLeaf{"Standby", groupinfo.Standby})
    groupinfo.EntityData.Leafs.Append("ha-state", types.YLeaf{"HaState", groupinfo.HaState})
    groupinfo.EntityData.Leafs.Append("nsr-state", types.YLeaf{"NsrState", groupinfo.NsrState})

    groupinfo.EntityData.YListKeys = []string {}

    return &(groupinfo.EntityData)
}

