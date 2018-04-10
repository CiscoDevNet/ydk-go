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

// VrfGroup
// VRF group operational data
type VrfGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node operational data.
    Nodes VrfGroup_Nodes
}

func (vrfGroup *VrfGroup) GetEntityData() *types.CommonEntityData {
    vrfGroup.EntityData.YFilter = vrfGroup.YFilter
    vrfGroup.EntityData.YangName = "vrf-group"
    vrfGroup.EntityData.BundleName = "cisco_ios_xr"
    vrfGroup.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-oper"
    vrfGroup.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-oper:vrf-group"
    vrfGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfGroup.EntityData.Children = make(map[string]types.YChild)
    vrfGroup.EntityData.Children["nodes"] = types.YChild{"Nodes", &vrfGroup.Nodes}
    vrfGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfGroup.EntityData)
}

// VrfGroup_Nodes
// Node operational data
type VrfGroup_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node details. The type is slice of VrfGroup_Nodes_Node.
    Node []VrfGroup_Nodes_Node
}

func (nodes *VrfGroup_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "vrf-group"
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

// VrfGroup_Nodes_Node
// Node details
type VrfGroup_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Group operational data.
    Groups VrfGroup_Nodes_Node_Groups
}

func (node *VrfGroup_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["groups"] = types.YChild{"Groups", &node.Groups}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// VrfGroup_Nodes_Node_Groups
// Group operational data
type VrfGroup_Nodes_Node_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group details. The type is slice of VrfGroup_Nodes_Node_Groups_Group.
    Group []VrfGroup_Nodes_Node_Groups_Group
}

func (groups *VrfGroup_Nodes_Node_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "node"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// VrfGroup_Nodes_Node_Groups_Group
// Group details
type VrfGroup_Nodes_Node_Groups_Group struct {
    EntityData types.CommonEntityData
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

func (group *VrfGroup_Nodes_Node_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range group.Vrf {
        group.EntityData.Children[types.GetSegmentPath(&group.Vrf[i])] = types.YChild{"Vrf", &group.Vrf[i]}
    }
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", group.GroupName}
    group.EntityData.Leafs["vr-fs"] = types.YLeaf{"VrFs", group.VrFs}
    group.EntityData.Leafs["forward-reference"] = types.YLeaf{"ForwardReference", group.ForwardReference}
    return &(group.EntityData)
}

// VrfGroup_Nodes_Node_Groups_Group_Vrf
// VRF group's VRF
type VrfGroup_Nodes_Node_Groups_Group_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}
}

func (vrf *VrfGroup_Nodes_Node_Groups_Group_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "group"
    vrf.EntityData.SegmentPath = "vrf"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Srlg
// srlg
type Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set of SRLG name, value maps configured.
    SrlgMaps Srlg_SrlgMaps

    // RSI SRLG operational data.
    Nodes Srlg_Nodes

    // Set of SRLG names configured.
    InterfaceSrlgNames Srlg_InterfaceSrlgNames
}

func (srlg *Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "cisco_ios_xr"
    srlg.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-oper"
    srlg.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-oper:srlg"
    srlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlg.EntityData.Children = make(map[string]types.YChild)
    srlg.EntityData.Children["srlg-maps"] = types.YChild{"SrlgMaps", &srlg.SrlgMaps}
    srlg.EntityData.Children["nodes"] = types.YChild{"Nodes", &srlg.Nodes}
    srlg.EntityData.Children["interface-srlg-names"] = types.YChild{"InterfaceSrlgNames", &srlg.InterfaceSrlgNames}
    srlg.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(srlg.EntityData)
}

// Srlg_SrlgMaps
// Set of SRLG name, value maps configured
type Srlg_SrlgMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of Srlg_SrlgMaps_SrlgMap.
    SrlgMap []Srlg_SrlgMaps_SrlgMap
}

func (srlgMaps *Srlg_SrlgMaps) GetEntityData() *types.CommonEntityData {
    srlgMaps.EntityData.YFilter = srlgMaps.YFilter
    srlgMaps.EntityData.YangName = "srlg-maps"
    srlgMaps.EntityData.BundleName = "cisco_ios_xr"
    srlgMaps.EntityData.ParentYangName = "srlg"
    srlgMaps.EntityData.SegmentPath = "srlg-maps"
    srlgMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgMaps.EntityData.Children = make(map[string]types.YChild)
    srlgMaps.EntityData.Children["srlg-map"] = types.YChild{"SrlgMap", nil}
    for i := range srlgMaps.SrlgMap {
        srlgMaps.EntityData.Children[types.GetSegmentPath(&srlgMaps.SrlgMap[i])] = types.YChild{"SrlgMap", &srlgMaps.SrlgMap[i]}
    }
    srlgMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(srlgMaps.EntityData)
}

// Srlg_SrlgMaps_SrlgMap
// Configured SRLG name details 
type Srlg_SrlgMaps_SrlgMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // SRLG name. The type is string.
    SrlgNameXr interface{}
}

func (srlgMap *Srlg_SrlgMaps_SrlgMap) GetEntityData() *types.CommonEntityData {
    srlgMap.EntityData.YFilter = srlgMap.YFilter
    srlgMap.EntityData.YangName = "srlg-map"
    srlgMap.EntityData.BundleName = "cisco_ios_xr"
    srlgMap.EntityData.ParentYangName = "srlg-maps"
    srlgMap.EntityData.SegmentPath = "srlg-map" + "[srlg-name='" + fmt.Sprintf("%v", srlgMap.SrlgName) + "']"
    srlgMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgMap.EntityData.Children = make(map[string]types.YChild)
    srlgMap.EntityData.Leafs = make(map[string]types.YLeaf)
    srlgMap.EntityData.Leafs["srlg-name"] = types.YLeaf{"SrlgName", srlgMap.SrlgName}
    srlgMap.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", srlgMap.SrlgValue}
    srlgMap.EntityData.Leafs["srlg-name-xr"] = types.YLeaf{"SrlgNameXr", srlgMap.SrlgNameXr}
    return &(srlgMap.EntityData)
}

// Srlg_Nodes
// RSI SRLG operational data
type Srlg_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSI SRLG operational data. The type is slice of Srlg_Nodes_Node.
    Node []Srlg_Nodes_Node
}

func (nodes *Srlg_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "srlg"
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

// Srlg_Nodes_Node
// RSI SRLG operational data
type Srlg_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Srlg_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["srlg-maps"] = types.YChild{"SrlgMaps", &node.SrlgMaps}
    node.EntityData.Children["groups"] = types.YChild{"Groups", &node.Groups}
    node.EntityData.Children["inherit-nodes"] = types.YChild{"InheritNodes", &node.InheritNodes}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["interface-details"] = types.YChild{"InterfaceDetails", &node.InterfaceDetails}
    node.EntityData.Children["srlg-values"] = types.YChild{"SrlgValues", &node.SrlgValues}
    node.EntityData.Children["interface-srlg-names"] = types.YChild{"InterfaceSrlgNames", &node.InterfaceSrlgNames}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Srlg_Nodes_Node_SrlgMaps
// Set of SRLG name, value maps configured
type Srlg_Nodes_Node_SrlgMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_Nodes_Node_SrlgMaps_SrlgMap.
    SrlgMap []Srlg_Nodes_Node_SrlgMaps_SrlgMap
}

func (srlgMaps *Srlg_Nodes_Node_SrlgMaps) GetEntityData() *types.CommonEntityData {
    srlgMaps.EntityData.YFilter = srlgMaps.YFilter
    srlgMaps.EntityData.YangName = "srlg-maps"
    srlgMaps.EntityData.BundleName = "cisco_ios_xr"
    srlgMaps.EntityData.ParentYangName = "node"
    srlgMaps.EntityData.SegmentPath = "srlg-maps"
    srlgMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgMaps.EntityData.Children = make(map[string]types.YChild)
    srlgMaps.EntityData.Children["srlg-map"] = types.YChild{"SrlgMap", nil}
    for i := range srlgMaps.SrlgMap {
        srlgMaps.EntityData.Children[types.GetSegmentPath(&srlgMaps.SrlgMap[i])] = types.YChild{"SrlgMap", &srlgMaps.SrlgMap[i]}
    }
    srlgMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(srlgMaps.EntityData)
}

// Srlg_Nodes_Node_SrlgMaps_SrlgMap
// Configured SRLG name details 
type Srlg_Nodes_Node_SrlgMaps_SrlgMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // SRLG name. The type is string.
    SrlgNameXr interface{}
}

func (srlgMap *Srlg_Nodes_Node_SrlgMaps_SrlgMap) GetEntityData() *types.CommonEntityData {
    srlgMap.EntityData.YFilter = srlgMap.YFilter
    srlgMap.EntityData.YangName = "srlg-map"
    srlgMap.EntityData.BundleName = "cisco_ios_xr"
    srlgMap.EntityData.ParentYangName = "srlg-maps"
    srlgMap.EntityData.SegmentPath = "srlg-map" + "[srlg-name='" + fmt.Sprintf("%v", srlgMap.SrlgName) + "']"
    srlgMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgMap.EntityData.Children = make(map[string]types.YChild)
    srlgMap.EntityData.Leafs = make(map[string]types.YLeaf)
    srlgMap.EntityData.Leafs["srlg-name"] = types.YLeaf{"SrlgName", srlgMap.SrlgName}
    srlgMap.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", srlgMap.SrlgValue}
    srlgMap.EntityData.Leafs["srlg-name-xr"] = types.YLeaf{"SrlgNameXr", srlgMap.SrlgNameXr}
    return &(srlgMap.EntityData)
}

// Srlg_Nodes_Node_Groups
// Set of Groups configured for SRLG
type Srlg_Nodes_Node_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG group details. The type is slice of Srlg_Nodes_Node_Groups_Group.
    Group []Srlg_Nodes_Node_Groups_Group
}

func (groups *Srlg_Nodes_Node_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "node"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Srlg_Nodes_Node_Groups_Group
// SRLG group details
type Srlg_Nodes_Node_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    GroupName interface{}

    // Group name. The type is string.
    GroupNameXr interface{}

    // Group values. The type is interface{} with range: 0..4294967295.
    GroupValues interface{}

    // SRLG attribute. The type is slice of
    // Srlg_Nodes_Node_Groups_Group_SrlgAttribute.
    SrlgAttribute []Srlg_Nodes_Node_Groups_Group_SrlgAttribute
}

func (group *Srlg_Nodes_Node_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["srlg-attribute"] = types.YChild{"SrlgAttribute", nil}
    for i := range group.SrlgAttribute {
        group.EntityData.Children[types.GetSegmentPath(&group.SrlgAttribute[i])] = types.YChild{"SrlgAttribute", &group.SrlgAttribute[i]}
    }
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", group.GroupName}
    group.EntityData.Leafs["group-name-xr"] = types.YLeaf{"GroupNameXr", group.GroupNameXr}
    group.EntityData.Leafs["group-values"] = types.YLeaf{"GroupValues", group.GroupValues}
    return &(group.EntityData)
}

// Srlg_Nodes_Node_Groups_Group_SrlgAttribute
// SRLG attribute
type Srlg_Nodes_Node_Groups_Group_SrlgAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Nodes_Node_Groups_Group_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "group"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = make(map[string]types.YChild)
    srlgAttribute.EntityData.Leafs = make(map[string]types.YLeaf)
    srlgAttribute.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue}
    srlgAttribute.EntityData.Leafs["priority"] = types.YLeaf{"Priority", srlgAttribute.Priority}
    srlgAttribute.EntityData.Leafs["srlg-index"] = types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex}
    return &(srlgAttribute.EntityData)
}

// Srlg_Nodes_Node_InheritNodes
// Set of inherit locations configured for SRLG
type Srlg_Nodes_Node_InheritNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG inherit location details. The type is slice of
    // Srlg_Nodes_Node_InheritNodes_InheritNode.
    InheritNode []Srlg_Nodes_Node_InheritNodes_InheritNode
}

func (inheritNodes *Srlg_Nodes_Node_InheritNodes) GetEntityData() *types.CommonEntityData {
    inheritNodes.EntityData.YFilter = inheritNodes.YFilter
    inheritNodes.EntityData.YangName = "inherit-nodes"
    inheritNodes.EntityData.BundleName = "cisco_ios_xr"
    inheritNodes.EntityData.ParentYangName = "node"
    inheritNodes.EntityData.SegmentPath = "inherit-nodes"
    inheritNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNodes.EntityData.Children = make(map[string]types.YChild)
    inheritNodes.EntityData.Children["inherit-node"] = types.YChild{"InheritNode", nil}
    for i := range inheritNodes.InheritNode {
        inheritNodes.EntityData.Children[types.GetSegmentPath(&inheritNodes.InheritNode[i])] = types.YChild{"InheritNode", &inheritNodes.InheritNode[i]}
    }
    inheritNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inheritNodes.EntityData)
}

// Srlg_Nodes_Node_InheritNodes_InheritNode
// SRLG inherit location details
type Srlg_Nodes_Node_InheritNodes_InheritNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Inherit node. The type is string with pattern:
    // b'((([a-zA-Z0-9_]*\\d+)|(\\*))/){2}(([a-zA-Z0-9_]*\\d+)|(\\*))'.
    InheritNodeName interface{}

    // Inherit node name. The type is string.
    NodeName interface{}

    // Node values. The type is interface{} with range: 0..4294967295.
    NodeValues interface{}

    // SRLG attribute. The type is slice of
    // Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute.
    SrlgAttribute []Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetEntityData() *types.CommonEntityData {
    inheritNode.EntityData.YFilter = inheritNode.YFilter
    inheritNode.EntityData.YangName = "inherit-node"
    inheritNode.EntityData.BundleName = "cisco_ios_xr"
    inheritNode.EntityData.ParentYangName = "inherit-nodes"
    inheritNode.EntityData.SegmentPath = "inherit-node" + "[inherit-node-name='" + fmt.Sprintf("%v", inheritNode.InheritNodeName) + "']"
    inheritNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNode.EntityData.Children = make(map[string]types.YChild)
    inheritNode.EntityData.Children["srlg-attribute"] = types.YChild{"SrlgAttribute", nil}
    for i := range inheritNode.SrlgAttribute {
        inheritNode.EntityData.Children[types.GetSegmentPath(&inheritNode.SrlgAttribute[i])] = types.YChild{"SrlgAttribute", &inheritNode.SrlgAttribute[i]}
    }
    inheritNode.EntityData.Leafs = make(map[string]types.YLeaf)
    inheritNode.EntityData.Leafs["inherit-node-name"] = types.YLeaf{"InheritNodeName", inheritNode.InheritNodeName}
    inheritNode.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", inheritNode.NodeName}
    inheritNode.EntityData.Leafs["node-values"] = types.YLeaf{"NodeValues", inheritNode.NodeValues}
    return &(inheritNode.EntityData)
}

// Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute
// SRLG attribute
type Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "inherit-node"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = make(map[string]types.YChild)
    srlgAttribute.EntityData.Leafs = make(map[string]types.YLeaf)
    srlgAttribute.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue}
    srlgAttribute.EntityData.Leafs["priority"] = types.YLeaf{"Priority", srlgAttribute.Priority}
    srlgAttribute.EntityData.Leafs["srlg-index"] = types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex}
    return &(srlgAttribute.EntityData)
}

// Srlg_Nodes_Node_Interfaces
// Set of interfaces configured for SRLG
type Srlg_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG interface summary. The type is slice of
    // Srlg_Nodes_Node_Interfaces_Interface_.
    Interface_ []Srlg_Nodes_Node_Interfaces_Interface
}

func (interfaces *Srlg_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Srlg_Nodes_Node_Interfaces_Interface
// SRLG interface summary
type Srlg_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *Srlg_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr}
    self.EntityData.Leafs["value-count"] = types.YLeaf{"ValueCount", self.ValueCount}
    self.EntityData.Leafs["registrations"] = types.YLeaf{"Registrations", self.Registrations}
    self.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", self.SrlgValue}
    return &(self.EntityData)
}

// Srlg_Nodes_Node_InterfaceDetails
// Set of interfaces configured for SRLG
type Srlg_Nodes_Node_InterfaceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG interface details. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail.
    InterfaceDetail []Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail
}

func (interfaceDetails *Srlg_Nodes_Node_InterfaceDetails) GetEntityData() *types.CommonEntityData {
    interfaceDetails.EntityData.YFilter = interfaceDetails.YFilter
    interfaceDetails.EntityData.YangName = "interface-details"
    interfaceDetails.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetails.EntityData.ParentYangName = "node"
    interfaceDetails.EntityData.SegmentPath = "interface-details"
    interfaceDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetails.EntityData.Children = make(map[string]types.YChild)
    interfaceDetails.EntityData.Children["interface-detail"] = types.YChild{"InterfaceDetail", nil}
    for i := range interfaceDetails.InterfaceDetail {
        interfaceDetails.EntityData.Children[types.GetSegmentPath(&interfaceDetails.InterfaceDetail[i])] = types.YChild{"InterfaceDetail", &interfaceDetails.InterfaceDetail[i]}
    }
    interfaceDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceDetails.EntityData)
}

// Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail
// SRLG interface details
type Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Groups. The type is interface{} with range: 0..4294967295.
    Groups interface{}

    // Nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // SRLG attributes. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute.
    SrlgAttribute []Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetEntityData() *types.CommonEntityData {
    interfaceDetail.EntityData.YFilter = interfaceDetail.YFilter
    interfaceDetail.EntityData.YangName = "interface-detail"
    interfaceDetail.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetail.EntityData.ParentYangName = "interface-details"
    interfaceDetail.EntityData.SegmentPath = "interface-detail" + "[interface-name='" + fmt.Sprintf("%v", interfaceDetail.InterfaceName) + "']"
    interfaceDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetail.EntityData.Children = make(map[string]types.YChild)
    interfaceDetail.EntityData.Children["srlg-attribute"] = types.YChild{"SrlgAttribute", nil}
    for i := range interfaceDetail.SrlgAttribute {
        interfaceDetail.EntityData.Children[types.GetSegmentPath(&interfaceDetail.SrlgAttribute[i])] = types.YChild{"SrlgAttribute", &interfaceDetail.SrlgAttribute[i]}
    }
    interfaceDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDetail.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceDetail.InterfaceName}
    interfaceDetail.EntityData.Leafs["groups"] = types.YLeaf{"Groups", interfaceDetail.Groups}
    interfaceDetail.EntityData.Leafs["nodes"] = types.YLeaf{"Nodes", interfaceDetail.Nodes}
    return &(interfaceDetail.EntityData)
}

// Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute
// SRLG attributes
type Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute struct {
    EntityData types.CommonEntityData
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

func (srlgAttribute *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "interface-detail"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = make(map[string]types.YChild)
    srlgAttribute.EntityData.Leafs = make(map[string]types.YLeaf)
    srlgAttribute.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue}
    srlgAttribute.EntityData.Leafs["priority"] = types.YLeaf{"Priority", srlgAttribute.Priority}
    srlgAttribute.EntityData.Leafs["source"] = types.YLeaf{"Source", srlgAttribute.Source}
    srlgAttribute.EntityData.Leafs["source-name"] = types.YLeaf{"SourceName", srlgAttribute.SourceName}
    srlgAttribute.EntityData.Leafs["srlg-index"] = types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex}
    return &(srlgAttribute.EntityData)
}

// Srlg_Nodes_Node_SrlgValues
// Set of SRLG values configured
type Srlg_Nodes_Node_SrlgValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG value details . The type is slice of
    // Srlg_Nodes_Node_SrlgValues_SrlgValue.
    SrlgValue []Srlg_Nodes_Node_SrlgValues_SrlgValue
}

func (srlgValues *Srlg_Nodes_Node_SrlgValues) GetEntityData() *types.CommonEntityData {
    srlgValues.EntityData.YFilter = srlgValues.YFilter
    srlgValues.EntityData.YangName = "srlg-values"
    srlgValues.EntityData.BundleName = "cisco_ios_xr"
    srlgValues.EntityData.ParentYangName = "node"
    srlgValues.EntityData.SegmentPath = "srlg-values"
    srlgValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgValues.EntityData.Children = make(map[string]types.YChild)
    srlgValues.EntityData.Children["srlg-value"] = types.YChild{"SrlgValue", nil}
    for i := range srlgValues.SrlgValue {
        srlgValues.EntityData.Children[types.GetSegmentPath(&srlgValues.SrlgValue[i])] = types.YChild{"SrlgValue", &srlgValues.SrlgValue[i]}
    }
    srlgValues.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(srlgValues.EntityData)
}

// Srlg_Nodes_Node_SrlgValues_SrlgValue
// Configured SRLG value details 
type Srlg_Nodes_Node_SrlgValues_SrlgValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG value. The type is interface{} with range:
    // -2147483648..2147483647.
    Value interface{}

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetEntityData() *types.CommonEntityData {
    srlgValue.EntityData.YFilter = srlgValue.YFilter
    srlgValue.EntityData.YangName = "srlg-value"
    srlgValue.EntityData.BundleName = "cisco_ios_xr"
    srlgValue.EntityData.ParentYangName = "srlg-values"
    srlgValue.EntityData.SegmentPath = "srlg-value" + "[value='" + fmt.Sprintf("%v", srlgValue.Value) + "']"
    srlgValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgValue.EntityData.Children = make(map[string]types.YChild)
    srlgValue.EntityData.Leafs = make(map[string]types.YLeaf)
    srlgValue.EntityData.Leafs["value"] = types.YLeaf{"Value", srlgValue.Value}
    srlgValue.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", srlgValue.InterfaceName}
    return &(srlgValue.EntityData)
}

// Srlg_Nodes_Node_InterfaceSrlgNames
// Set of SRLG names configured
type Srlg_Nodes_Node_InterfaceSrlgNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName
}

func (interfaceSrlgNames *Srlg_Nodes_Node_InterfaceSrlgNames) GetEntityData() *types.CommonEntityData {
    interfaceSrlgNames.EntityData.YFilter = interfaceSrlgNames.YFilter
    interfaceSrlgNames.EntityData.YangName = "interface-srlg-names"
    interfaceSrlgNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceSrlgNames.EntityData.ParentYangName = "node"
    interfaceSrlgNames.EntityData.SegmentPath = "interface-srlg-names"
    interfaceSrlgNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgNames.EntityData.Children = make(map[string]types.YChild)
    interfaceSrlgNames.EntityData.Children["interface-srlg-name"] = types.YChild{"InterfaceSrlgName", nil}
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        interfaceSrlgNames.EntityData.Children[types.GetSegmentPath(&interfaceSrlgNames.InterfaceSrlgName[i])] = types.YChild{"InterfaceSrlgName", &interfaceSrlgNames.InterfaceSrlgName[i]}
    }
    interfaceSrlgNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSrlgNames.EntityData)
}

// Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName
// Configured SRLG name details 
type Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName struct {
    EntityData types.CommonEntityData
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

func (interfaceSrlgName *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName) GetEntityData() *types.CommonEntityData {
    interfaceSrlgName.EntityData.YFilter = interfaceSrlgName.YFilter
    interfaceSrlgName.EntityData.YangName = "interface-srlg-name"
    interfaceSrlgName.EntityData.BundleName = "cisco_ios_xr"
    interfaceSrlgName.EntityData.ParentYangName = "interface-srlg-names"
    interfaceSrlgName.EntityData.SegmentPath = "interface-srlg-name" + "[srlg-name='" + fmt.Sprintf("%v", interfaceSrlgName.SrlgName) + "']"
    interfaceSrlgName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgName.EntityData.Children = make(map[string]types.YChild)
    interfaceSrlgName.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &interfaceSrlgName.Interfaces}
    interfaceSrlgName.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSrlgName.EntityData.Leafs["srlg-name"] = types.YLeaf{"SrlgName", interfaceSrlgName.SrlgName}
    interfaceSrlgName.EntityData.Leafs["srlg-name-xr"] = types.YLeaf{"SrlgNameXr", interfaceSrlgName.SrlgNameXr}
    interfaceSrlgName.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", interfaceSrlgName.SrlgValue}
    return &(interfaceSrlgName.EntityData)
}

// Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces
// Interfaces information
type Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (interfaces *Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-srlg-name"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaces.InterfaceName}
    return &(interfaces.EntityData)
}

// Srlg_InterfaceSrlgNames
// Set of SRLG names configured
type Srlg_InterfaceSrlgNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []Srlg_InterfaceSrlgNames_InterfaceSrlgName
}

func (interfaceSrlgNames *Srlg_InterfaceSrlgNames) GetEntityData() *types.CommonEntityData {
    interfaceSrlgNames.EntityData.YFilter = interfaceSrlgNames.YFilter
    interfaceSrlgNames.EntityData.YangName = "interface-srlg-names"
    interfaceSrlgNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceSrlgNames.EntityData.ParentYangName = "srlg"
    interfaceSrlgNames.EntityData.SegmentPath = "interface-srlg-names"
    interfaceSrlgNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgNames.EntityData.Children = make(map[string]types.YChild)
    interfaceSrlgNames.EntityData.Children["interface-srlg-name"] = types.YChild{"InterfaceSrlgName", nil}
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        interfaceSrlgNames.EntityData.Children[types.GetSegmentPath(&interfaceSrlgNames.InterfaceSrlgName[i])] = types.YChild{"InterfaceSrlgName", &interfaceSrlgNames.InterfaceSrlgName[i]}
    }
    interfaceSrlgNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSrlgNames.EntityData)
}

// Srlg_InterfaceSrlgNames_InterfaceSrlgName
// Configured SRLG name details 
type Srlg_InterfaceSrlgNames_InterfaceSrlgName struct {
    EntityData types.CommonEntityData
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

func (interfaceSrlgName *Srlg_InterfaceSrlgNames_InterfaceSrlgName) GetEntityData() *types.CommonEntityData {
    interfaceSrlgName.EntityData.YFilter = interfaceSrlgName.YFilter
    interfaceSrlgName.EntityData.YangName = "interface-srlg-name"
    interfaceSrlgName.EntityData.BundleName = "cisco_ios_xr"
    interfaceSrlgName.EntityData.ParentYangName = "interface-srlg-names"
    interfaceSrlgName.EntityData.SegmentPath = "interface-srlg-name" + "[srlg-name='" + fmt.Sprintf("%v", interfaceSrlgName.SrlgName) + "']"
    interfaceSrlgName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgName.EntityData.Children = make(map[string]types.YChild)
    interfaceSrlgName.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &interfaceSrlgName.Interfaces}
    interfaceSrlgName.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSrlgName.EntityData.Leafs["srlg-name"] = types.YLeaf{"SrlgName", interfaceSrlgName.SrlgName}
    interfaceSrlgName.EntityData.Leafs["srlg-name-xr"] = types.YLeaf{"SrlgNameXr", interfaceSrlgName.SrlgNameXr}
    interfaceSrlgName.EntityData.Leafs["srlg-value"] = types.YLeaf{"SrlgValue", interfaceSrlgName.SrlgValue}
    return &(interfaceSrlgName.EntityData)
}

// Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces
// Interfaces information
type Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (interfaces *Srlg_InterfaceSrlgNames_InterfaceSrlgName_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-srlg-name"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaces.InterfaceName}
    return &(interfaces.EntityData)
}

// SelectiveVrfDownload
// selective vrf download
type SelectiveVrfDownload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selective VRF Download feature state details.
    State SelectiveVrfDownload_State
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetEntityData() *types.CommonEntityData {
    selectiveVrfDownload.EntityData.YFilter = selectiveVrfDownload.YFilter
    selectiveVrfDownload.EntityData.YangName = "selective-vrf-download"
    selectiveVrfDownload.EntityData.BundleName = "cisco_ios_xr"
    selectiveVrfDownload.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-oper"
    selectiveVrfDownload.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-oper:selective-vrf-download"
    selectiveVrfDownload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectiveVrfDownload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectiveVrfDownload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectiveVrfDownload.EntityData.Children = make(map[string]types.YChild)
    selectiveVrfDownload.EntityData.Children["state"] = types.YChild{"State", &selectiveVrfDownload.State}
    selectiveVrfDownload.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(selectiveVrfDownload.EntityData)
}

// SelectiveVrfDownload_State
// Selective VRF Download feature state details
type SelectiveVrfDownload_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SVD Enabled Operational. The type is bool.
    IsSvdEnabled interface{}

    // Is SVD Enabled Config. The type is bool.
    IsSvdEnabledCfg interface{}
}

func (state *SelectiveVrfDownload_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "selective-vrf-download"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["is-svd-enabled"] = types.YLeaf{"IsSvdEnabled", state.IsSvdEnabled}
    state.EntityData.Leafs["is-svd-enabled-cfg"] = types.YLeaf{"IsSvdEnabledCfg", state.IsSvdEnabledCfg}
    return &(state.EntityData)
}

