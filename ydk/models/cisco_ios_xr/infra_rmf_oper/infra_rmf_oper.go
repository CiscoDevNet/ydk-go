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

    redundancy.EntityData.Children = make(map[string]types.YChild)
    redundancy.EntityData.Children["nodes"] = types.YChild{"Nodes", &redundancy.Nodes}
    redundancy.EntityData.Children["summary"] = types.YChild{"Summary", &redundancy.Summary}
    redundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redundancy.EntityData)
}

// Redundancy_Nodes
// Location show information
type Redundancy_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Node Information. The type is slice of Redundancy_Nodes_Node.
    Node []Redundancy_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Redundancy_Nodes_Node
// Redundancy Node Information
type Redundancy_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node Location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    Redundancy Redundancy_Nodes_Node_Redundancy_
}

func (node *Redundancy_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["redundancy"] = types.YChild{"Redundancy", &node.Redundancy}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["log"] = types.YLeaf{"Log", node.Log}
    node.EntityData.Leafs["active-reboot-reason"] = types.YLeaf{"ActiveRebootReason", node.ActiveRebootReason}
    node.EntityData.Leafs["standby-reboot-reason"] = types.YLeaf{"StandbyRebootReason", node.StandbyRebootReason}
    node.EntityData.Leafs["err-log"] = types.YLeaf{"ErrLog", node.ErrLog}
    return &(node.EntityData)
}

// Redundancy_Nodes_Node_Redundancy_
// Row information
type Redundancy_Nodes_Node_Redundancy_ struct {
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

    // groupinfo. The type is slice of
    // Redundancy_Nodes_Node_Redundancy__Groupinfo.
    Groupinfo []Redundancy_Nodes_Node_Redundancy__Groupinfo
}

func (redundancy_ *Redundancy_Nodes_Node_Redundancy_) GetEntityData() *types.CommonEntityData {
    redundancy_.EntityData.YFilter = redundancy_.YFilter
    redundancy_.EntityData.YangName = "redundancy"
    redundancy_.EntityData.BundleName = "cisco_ios_xr"
    redundancy_.EntityData.ParentYangName = "node"
    redundancy_.EntityData.SegmentPath = "redundancy"
    redundancy_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy_.EntityData.Children = make(map[string]types.YChild)
    redundancy_.EntityData.Children["groupinfo"] = types.YChild{"Groupinfo", nil}
    for i := range redundancy_.Groupinfo {
        redundancy_.EntityData.Children[types.GetSegmentPath(&redundancy_.Groupinfo[i])] = types.YChild{"Groupinfo", &redundancy_.Groupinfo[i]}
    }
    redundancy_.EntityData.Leafs = make(map[string]types.YLeaf)
    redundancy_.EntityData.Leafs["active"] = types.YLeaf{"Active", redundancy_.Active}
    redundancy_.EntityData.Leafs["standby"] = types.YLeaf{"Standby", redundancy_.Standby}
    redundancy_.EntityData.Leafs["ha-state"] = types.YLeaf{"HaState", redundancy_.HaState}
    redundancy_.EntityData.Leafs["nsr-state"] = types.YLeaf{"NsrState", redundancy_.NsrState}
    return &(redundancy_.EntityData)
}

// Redundancy_Nodes_Node_Redundancy__Groupinfo
// groupinfo
type Redundancy_Nodes_Node_Redundancy__Groupinfo struct {
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

func (groupinfo *Redundancy_Nodes_Node_Redundancy__Groupinfo) GetEntityData() *types.CommonEntityData {
    groupinfo.EntityData.YFilter = groupinfo.YFilter
    groupinfo.EntityData.YangName = "groupinfo"
    groupinfo.EntityData.BundleName = "cisco_ios_xr"
    groupinfo.EntityData.ParentYangName = "redundancy"
    groupinfo.EntityData.SegmentPath = "groupinfo"
    groupinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupinfo.EntityData.Children = make(map[string]types.YChild)
    groupinfo.EntityData.Leafs = make(map[string]types.YLeaf)
    groupinfo.EntityData.Leafs["active"] = types.YLeaf{"Active", groupinfo.Active}
    groupinfo.EntityData.Leafs["standby"] = types.YLeaf{"Standby", groupinfo.Standby}
    groupinfo.EntityData.Leafs["ha-state"] = types.YLeaf{"HaState", groupinfo.HaState}
    groupinfo.EntityData.Leafs["nsr-state"] = types.YLeaf{"NsrState", groupinfo.NsrState}
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
    RedPair []Redundancy_Summary_RedPair
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

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["red-pair"] = types.YChild{"RedPair", nil}
    for i := range summary.RedPair {
        summary.EntityData.Children[types.GetSegmentPath(&summary.RedPair[i])] = types.YChild{"RedPair", &summary.RedPair[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["err-log"] = types.YLeaf{"ErrLog", summary.ErrLog}
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
    Groupinfo []Redundancy_Summary_RedPair_Groupinfo
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

    redPair.EntityData.Children = make(map[string]types.YChild)
    redPair.EntityData.Children["groupinfo"] = types.YChild{"Groupinfo", nil}
    for i := range redPair.Groupinfo {
        redPair.EntityData.Children[types.GetSegmentPath(&redPair.Groupinfo[i])] = types.YChild{"Groupinfo", &redPair.Groupinfo[i]}
    }
    redPair.EntityData.Leafs = make(map[string]types.YLeaf)
    redPair.EntityData.Leafs["active"] = types.YLeaf{"Active", redPair.Active}
    redPair.EntityData.Leafs["standby"] = types.YLeaf{"Standby", redPair.Standby}
    redPair.EntityData.Leafs["ha-state"] = types.YLeaf{"HaState", redPair.HaState}
    redPair.EntityData.Leafs["nsr-state"] = types.YLeaf{"NsrState", redPair.NsrState}
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

    groupinfo.EntityData.Children = make(map[string]types.YChild)
    groupinfo.EntityData.Leafs = make(map[string]types.YLeaf)
    groupinfo.EntityData.Leafs["active"] = types.YLeaf{"Active", groupinfo.Active}
    groupinfo.EntityData.Leafs["standby"] = types.YLeaf{"Standby", groupinfo.Standby}
    groupinfo.EntityData.Leafs["ha-state"] = types.YLeaf{"HaState", groupinfo.HaState}
    groupinfo.EntityData.Leafs["nsr-state"] = types.YLeaf{"NsrState", groupinfo.NsrState}
    return &(groupinfo.EntityData)
}

