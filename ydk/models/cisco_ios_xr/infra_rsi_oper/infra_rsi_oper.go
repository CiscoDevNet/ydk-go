// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rsi package operational data.
// 
// This module contains definitions
// for the following management objects:
//   vrf-group: VRF group operational data
//   srlg: srlg
//   selective-vrf-download: selective vrf download
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // Invalid
    Priority_invald Priority = "invald"
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

    vrfGroup.EntityData.Children = types.NewOrderedMap()
    vrfGroup.EntityData.Children.Append("nodes", types.YChild{"Nodes", &vrfGroup.Nodes})
    vrfGroup.EntityData.Leafs = types.NewOrderedMap()

    vrfGroup.EntityData.YListKeys = []string {}

    return &(vrfGroup.EntityData)
}

// VrfGroup_Nodes
// Node operational data
type VrfGroup_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node details. The type is slice of VrfGroup_Nodes_Node.
    Node []*VrfGroup_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// VrfGroup_Nodes_Node
// Node details
type VrfGroup_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Group operational data.
    Groups VrfGroup_Nodes_Node_Groups
}

func (node *VrfGroup_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("groups", types.YChild{"Groups", &node.Groups})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// VrfGroup_Nodes_Node_Groups
// Group operational data
type VrfGroup_Nodes_Node_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group details. The type is slice of VrfGroup_Nodes_Node_Groups_Group.
    Group []*VrfGroup_Nodes_Node_Groups_Group
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

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

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
    Vrf []*VrfGroup_Nodes_Node_Groups_Group_Vrf
}

func (group *VrfGroup_Nodes_Node_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupName, "group-name")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range group.Vrf {
        group.EntityData.Children.Append(types.GetSegmentPath(group.Vrf[i]), types.YChild{"Vrf", group.Vrf[i]})
    }
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", group.GroupName})
    group.EntityData.Leafs.Append("vr-fs", types.YLeaf{"VrFs", group.VrFs})
    group.EntityData.Leafs.Append("forward-reference", types.YLeaf{"ForwardReference", group.ForwardReference})

    group.EntityData.YListKeys = []string {"GroupName"}

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

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {}

    return &(vrf.EntityData)
}

// Srlg
// srlg
type Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set of Groups configured for SRLG.
    Groups Srlg_Groups

    // Set of interfaces configured for SRLG.
    Interfaces Srlg_Interfaces

    // Set of rsip configured for SRLG.
    Rsips Srlg_Rsips

    // Set of SRLG name, value maps configured.
    SrlgMaps Srlg_SrlgMaps

    // RSI SRLG operational data.
    Nodes Srlg_Nodes

    // Set of SRLG names configured.
    InterfaceSrlgNames Srlg_InterfaceSrlgNames

    // Set of inherit locations configured for SRLG.
    InheritNodes Srlg_InheritNodes

    // Set of SRLG values configured.
    SrlgValues Srlg_SrlgValues

    // Set of interfaces configured for SRLG.
    InterfaceDetails Srlg_InterfaceDetails
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

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Children.Append("groups", types.YChild{"Groups", &srlg.Groups})
    srlg.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &srlg.Interfaces})
    srlg.EntityData.Children.Append("rsips", types.YChild{"Rsips", &srlg.Rsips})
    srlg.EntityData.Children.Append("srlg-maps", types.YChild{"SrlgMaps", &srlg.SrlgMaps})
    srlg.EntityData.Children.Append("nodes", types.YChild{"Nodes", &srlg.Nodes})
    srlg.EntityData.Children.Append("interface-srlg-names", types.YChild{"InterfaceSrlgNames", &srlg.InterfaceSrlgNames})
    srlg.EntityData.Children.Append("inherit-nodes", types.YChild{"InheritNodes", &srlg.InheritNodes})
    srlg.EntityData.Children.Append("srlg-values", types.YChild{"SrlgValues", &srlg.SrlgValues})
    srlg.EntityData.Children.Append("interface-details", types.YChild{"InterfaceDetails", &srlg.InterfaceDetails})
    srlg.EntityData.Leafs = types.NewOrderedMap()

    srlg.EntityData.YListKeys = []string {}

    return &(srlg.EntityData)
}

// Srlg_Groups
// Set of Groups configured for SRLG
type Srlg_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG group details. The type is slice of Srlg_Groups_Group.
    Group []*Srlg_Groups_Group
}

func (groups *Srlg_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "srlg"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Srlg_Groups_Group
// SRLG group details
type Srlg_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupName interface{}

    // Group name. The type is string.
    GroupNameXr interface{}

    // Group values. The type is interface{} with range: 0..4294967295.
    GroupValues interface{}

    // SRLG attribute. The type is slice of Srlg_Groups_Group_SrlgAttribute.
    SrlgAttribute []*Srlg_Groups_Group_SrlgAttribute
}

func (group *Srlg_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupName, "group-name")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range group.SrlgAttribute {
        group.EntityData.Children.Append(types.GetSegmentPath(group.SrlgAttribute[i]), types.YChild{"SrlgAttribute", group.SrlgAttribute[i]})
    }
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", group.GroupName})
    group.EntityData.Leafs.Append("group-name-xr", types.YLeaf{"GroupNameXr", group.GroupNameXr})
    group.EntityData.Leafs.Append("group-values", types.YLeaf{"GroupValues", group.GroupValues})

    group.EntityData.YListKeys = []string {"GroupName"}

    return &(group.EntityData)
}

// Srlg_Groups_Group_SrlgAttribute
// SRLG attribute
type Srlg_Groups_Group_SrlgAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Groups_Group_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "group"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_Interfaces
// Set of interfaces configured for SRLG
type Srlg_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG interface summary. The type is slice of Srlg_Interfaces_Interface.
    Interface []*Srlg_Interfaces_Interface
}

func (interfaces *Srlg_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "srlg"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Srlg_Interfaces_Interface
// SRLG interface summary
type Srlg_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (self *Srlg_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr})
    self.EntityData.Leafs.Append("value-count", types.YLeaf{"ValueCount", self.ValueCount})
    self.EntityData.Leafs.Append("registrations", types.YLeaf{"Registrations", self.Registrations})
    self.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", self.SrlgValue})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Srlg_Rsips
// Set of rsip configured for SRLG
type Srlg_Rsips struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG rsip details. The type is slice of Srlg_Rsips_Rsip.
    Rsip []*Srlg_Rsips_Rsip
}

func (rsips *Srlg_Rsips) GetEntityData() *types.CommonEntityData {
    rsips.EntityData.YFilter = rsips.YFilter
    rsips.EntityData.YangName = "rsips"
    rsips.EntityData.BundleName = "cisco_ios_xr"
    rsips.EntityData.ParentYangName = "srlg"
    rsips.EntityData.SegmentPath = "rsips"
    rsips.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsips.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsips.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsips.EntityData.Children = types.NewOrderedMap()
    rsips.EntityData.Children.Append("rsip", types.YChild{"Rsip", nil})
    for i := range rsips.Rsip {
        rsips.EntityData.Children.Append(types.GetSegmentPath(rsips.Rsip[i]), types.YChild{"Rsip", rsips.Rsip[i]})
    }
    rsips.EntityData.Leafs = types.NewOrderedMap()

    rsips.EntityData.YListKeys = []string {}

    return &(rsips.EntityData)
}

// Srlg_Rsips_Rsip
// SRLG rsip details
type Srlg_Rsips_Rsip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    RsipName interface{}

    // Optical layer interface name. The type is string.
    OpticalLayerInterfaceName interface{}

    // Registrations. The type is interface{} with range: 0..4294967295.
    Registrations interface{}

    // Interface values. The type is interface{} with range: 0..4294967295.
    InterfaceValues interface{}

    // SRLG attribute. The type is slice of Srlg_Rsips_Rsip_SrlgAttribute.
    SrlgAttribute []*Srlg_Rsips_Rsip_SrlgAttribute
}

func (rsip *Srlg_Rsips_Rsip) GetEntityData() *types.CommonEntityData {
    rsip.EntityData.YFilter = rsip.YFilter
    rsip.EntityData.YangName = "rsip"
    rsip.EntityData.BundleName = "cisco_ios_xr"
    rsip.EntityData.ParentYangName = "rsips"
    rsip.EntityData.SegmentPath = "rsip" + types.AddKeyToken(rsip.RsipName, "rsip-name")
    rsip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsip.EntityData.Children = types.NewOrderedMap()
    rsip.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range rsip.SrlgAttribute {
        rsip.EntityData.Children.Append(types.GetSegmentPath(rsip.SrlgAttribute[i]), types.YChild{"SrlgAttribute", rsip.SrlgAttribute[i]})
    }
    rsip.EntityData.Leafs = types.NewOrderedMap()
    rsip.EntityData.Leafs.Append("rsip-name", types.YLeaf{"RsipName", rsip.RsipName})
    rsip.EntityData.Leafs.Append("optical-layer-interface-name", types.YLeaf{"OpticalLayerInterfaceName", rsip.OpticalLayerInterfaceName})
    rsip.EntityData.Leafs.Append("registrations", types.YLeaf{"Registrations", rsip.Registrations})
    rsip.EntityData.Leafs.Append("interface-values", types.YLeaf{"InterfaceValues", rsip.InterfaceValues})

    rsip.EntityData.YListKeys = []string {"RsipName"}

    return &(rsip.EntityData)
}

// Srlg_Rsips_Rsip_SrlgAttribute
// SRLG attribute
type Srlg_Rsips_Rsip_SrlgAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_Rsips_Rsip_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "rsip"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_SrlgMaps
// Set of SRLG name, value maps configured
type Srlg_SrlgMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of Srlg_SrlgMaps_SrlgMap.
    SrlgMap []*Srlg_SrlgMaps_SrlgMap
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

    srlgMaps.EntityData.Children = types.NewOrderedMap()
    srlgMaps.EntityData.Children.Append("srlg-map", types.YChild{"SrlgMap", nil})
    for i := range srlgMaps.SrlgMap {
        srlgMaps.EntityData.Children.Append(types.GetSegmentPath(srlgMaps.SrlgMap[i]), types.YChild{"SrlgMap", srlgMaps.SrlgMap[i]})
    }
    srlgMaps.EntityData.Leafs = types.NewOrderedMap()

    srlgMaps.EntityData.YListKeys = []string {}

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
    srlgMap.EntityData.SegmentPath = "srlg-map" + types.AddKeyToken(srlgMap.SrlgName, "srlg-name")
    srlgMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgMap.EntityData.Children = types.NewOrderedMap()
    srlgMap.EntityData.Leafs = types.NewOrderedMap()
    srlgMap.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", srlgMap.SrlgName})
    srlgMap.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgMap.SrlgValue})
    srlgMap.EntityData.Leafs.Append("srlg-name-xr", types.YLeaf{"SrlgNameXr", srlgMap.SrlgNameXr})

    srlgMap.EntityData.YListKeys = []string {"SrlgName"}

    return &(srlgMap.EntityData)
}

// Srlg_Nodes
// RSI SRLG operational data
type Srlg_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSI SRLG operational data. The type is slice of Srlg_Nodes_Node.
    Node []*Srlg_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Srlg_Nodes_Node
// RSI SRLG operational data
type Srlg_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Srlg_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("srlg-maps", types.YChild{"SrlgMaps", &node.SrlgMaps})
    node.EntityData.Children.Append("groups", types.YChild{"Groups", &node.Groups})
    node.EntityData.Children.Append("inherit-nodes", types.YChild{"InheritNodes", &node.InheritNodes})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("interface-details", types.YChild{"InterfaceDetails", &node.InterfaceDetails})
    node.EntityData.Children.Append("srlg-values", types.YChild{"SrlgValues", &node.SrlgValues})
    node.EntityData.Children.Append("interface-srlg-names", types.YChild{"InterfaceSrlgNames", &node.InterfaceSrlgNames})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Srlg_Nodes_Node_SrlgMaps
// Set of SRLG name, value maps configured
type Srlg_Nodes_Node_SrlgMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_Nodes_Node_SrlgMaps_SrlgMap.
    SrlgMap []*Srlg_Nodes_Node_SrlgMaps_SrlgMap
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

    srlgMaps.EntityData.Children = types.NewOrderedMap()
    srlgMaps.EntityData.Children.Append("srlg-map", types.YChild{"SrlgMap", nil})
    for i := range srlgMaps.SrlgMap {
        srlgMaps.EntityData.Children.Append(types.GetSegmentPath(srlgMaps.SrlgMap[i]), types.YChild{"SrlgMap", srlgMaps.SrlgMap[i]})
    }
    srlgMaps.EntityData.Leafs = types.NewOrderedMap()

    srlgMaps.EntityData.YListKeys = []string {}

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
    srlgMap.EntityData.SegmentPath = "srlg-map" + types.AddKeyToken(srlgMap.SrlgName, "srlg-name")
    srlgMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgMap.EntityData.Children = types.NewOrderedMap()
    srlgMap.EntityData.Leafs = types.NewOrderedMap()
    srlgMap.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", srlgMap.SrlgName})
    srlgMap.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgMap.SrlgValue})
    srlgMap.EntityData.Leafs.Append("srlg-name-xr", types.YLeaf{"SrlgNameXr", srlgMap.SrlgNameXr})

    srlgMap.EntityData.YListKeys = []string {"SrlgName"}

    return &(srlgMap.EntityData)
}

// Srlg_Nodes_Node_Groups
// Set of Groups configured for SRLG
type Srlg_Nodes_Node_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG group details. The type is slice of Srlg_Nodes_Node_Groups_Group.
    Group []*Srlg_Nodes_Node_Groups_Group
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

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Srlg_Nodes_Node_Groups_Group
// SRLG group details
type Srlg_Nodes_Node_Groups_Group struct {
    EntityData types.CommonEntityData
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
    SrlgAttribute []*Srlg_Nodes_Node_Groups_Group_SrlgAttribute
}

func (group *Srlg_Nodes_Node_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupName, "group-name")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range group.SrlgAttribute {
        group.EntityData.Children.Append(types.GetSegmentPath(group.SrlgAttribute[i]), types.YChild{"SrlgAttribute", group.SrlgAttribute[i]})
    }
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", group.GroupName})
    group.EntityData.Leafs.Append("group-name-xr", types.YLeaf{"GroupNameXr", group.GroupNameXr})
    group.EntityData.Leafs.Append("group-values", types.YLeaf{"GroupValues", group.GroupValues})

    group.EntityData.YListKeys = []string {"GroupName"}

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

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_Nodes_Node_InheritNodes
// Set of inherit locations configured for SRLG
type Srlg_Nodes_Node_InheritNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG inherit location details. The type is slice of
    // Srlg_Nodes_Node_InheritNodes_InheritNode.
    InheritNode []*Srlg_Nodes_Node_InheritNodes_InheritNode
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

    inheritNodes.EntityData.Children = types.NewOrderedMap()
    inheritNodes.EntityData.Children.Append("inherit-node", types.YChild{"InheritNode", nil})
    for i := range inheritNodes.InheritNode {
        inheritNodes.EntityData.Children.Append(types.GetSegmentPath(inheritNodes.InheritNode[i]), types.YChild{"InheritNode", inheritNodes.InheritNode[i]})
    }
    inheritNodes.EntityData.Leafs = types.NewOrderedMap()

    inheritNodes.EntityData.YListKeys = []string {}

    return &(inheritNodes.EntityData)
}

// Srlg_Nodes_Node_InheritNodes_InheritNode
// SRLG inherit location details
type Srlg_Nodes_Node_InheritNodes_InheritNode struct {
    EntityData types.CommonEntityData
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
    SrlgAttribute []*Srlg_Nodes_Node_InheritNodes_InheritNode_SrlgAttribute
}

func (inheritNode *Srlg_Nodes_Node_InheritNodes_InheritNode) GetEntityData() *types.CommonEntityData {
    inheritNode.EntityData.YFilter = inheritNode.YFilter
    inheritNode.EntityData.YangName = "inherit-node"
    inheritNode.EntityData.BundleName = "cisco_ios_xr"
    inheritNode.EntityData.ParentYangName = "inherit-nodes"
    inheritNode.EntityData.SegmentPath = "inherit-node" + types.AddKeyToken(inheritNode.InheritNodeName, "inherit-node-name")
    inheritNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNode.EntityData.Children = types.NewOrderedMap()
    inheritNode.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range inheritNode.SrlgAttribute {
        inheritNode.EntityData.Children.Append(types.GetSegmentPath(inheritNode.SrlgAttribute[i]), types.YChild{"SrlgAttribute", inheritNode.SrlgAttribute[i]})
    }
    inheritNode.EntityData.Leafs = types.NewOrderedMap()
    inheritNode.EntityData.Leafs.Append("inherit-node-name", types.YLeaf{"InheritNodeName", inheritNode.InheritNodeName})
    inheritNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", inheritNode.NodeName})
    inheritNode.EntityData.Leafs.Append("node-values", types.YLeaf{"NodeValues", inheritNode.NodeValues})

    inheritNode.EntityData.YListKeys = []string {"InheritNodeName"}

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

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_Nodes_Node_Interfaces
// Set of interfaces configured for SRLG
type Srlg_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG interface summary. The type is slice of
    // Srlg_Nodes_Node_Interfaces_Interface.
    Interface []*Srlg_Nodes_Node_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Srlg_Nodes_Node_Interfaces_Interface
// SRLG interface summary
type Srlg_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr})
    self.EntityData.Leafs.Append("value-count", types.YLeaf{"ValueCount", self.ValueCount})
    self.EntityData.Leafs.Append("registrations", types.YLeaf{"Registrations", self.Registrations})
    self.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", self.SrlgValue})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Srlg_Nodes_Node_InterfaceDetails
// Set of interfaces configured for SRLG
type Srlg_Nodes_Node_InterfaceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG interface details. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail.
    InterfaceDetail []*Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail
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

    interfaceDetails.EntityData.Children = types.NewOrderedMap()
    interfaceDetails.EntityData.Children.Append("interface-detail", types.YChild{"InterfaceDetail", nil})
    for i := range interfaceDetails.InterfaceDetail {
        interfaceDetails.EntityData.Children.Append(types.GetSegmentPath(interfaceDetails.InterfaceDetail[i]), types.YChild{"InterfaceDetail", interfaceDetails.InterfaceDetail[i]})
    }
    interfaceDetails.EntityData.Leafs = types.NewOrderedMap()

    interfaceDetails.EntityData.YListKeys = []string {}

    return &(interfaceDetails.EntityData)
}

// Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail
// SRLG interface details
type Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Groups. The type is interface{} with range: 0..4294967295.
    Groups interface{}

    // Nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // SRLG attributes. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute.
    SrlgAttribute []*Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_SrlgAttribute

    // rsip list. The type is slice of
    // Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_Rsip.
    Rsip []*Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_Rsip
}

func (interfaceDetail *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail) GetEntityData() *types.CommonEntityData {
    interfaceDetail.EntityData.YFilter = interfaceDetail.YFilter
    interfaceDetail.EntityData.YangName = "interface-detail"
    interfaceDetail.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetail.EntityData.ParentYangName = "interface-details"
    interfaceDetail.EntityData.SegmentPath = "interface-detail" + types.AddKeyToken(interfaceDetail.InterfaceName, "interface-name")
    interfaceDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetail.EntityData.Children = types.NewOrderedMap()
    interfaceDetail.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range interfaceDetail.SrlgAttribute {
        interfaceDetail.EntityData.Children.Append(types.GetSegmentPath(interfaceDetail.SrlgAttribute[i]), types.YChild{"SrlgAttribute", interfaceDetail.SrlgAttribute[i]})
    }
    interfaceDetail.EntityData.Children.Append("rsip", types.YChild{"Rsip", nil})
    for i := range interfaceDetail.Rsip {
        interfaceDetail.EntityData.Children.Append(types.GetSegmentPath(interfaceDetail.Rsip[i]), types.YChild{"Rsip", interfaceDetail.Rsip[i]})
    }
    interfaceDetail.EntityData.Leafs = types.NewOrderedMap()
    interfaceDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceDetail.InterfaceName})
    interfaceDetail.EntityData.Leafs.Append("groups", types.YLeaf{"Groups", interfaceDetail.Groups})
    interfaceDetail.EntityData.Leafs.Append("nodes", types.YLeaf{"Nodes", interfaceDetail.Nodes})

    interfaceDetail.EntityData.YListKeys = []string {"InterfaceName"}

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

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("source", types.YLeaf{"Source", srlgAttribute.Source})
    srlgAttribute.EntityData.Leafs.Append("source-name", types.YLeaf{"SourceName", srlgAttribute.SourceName})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_Rsip
// rsip list
type Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_Rsip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of names matching rsip. The type is string.
    RsipName interface{}
}

func (rsip *Srlg_Nodes_Node_InterfaceDetails_InterfaceDetail_Rsip) GetEntityData() *types.CommonEntityData {
    rsip.EntityData.YFilter = rsip.YFilter
    rsip.EntityData.YangName = "rsip"
    rsip.EntityData.BundleName = "cisco_ios_xr"
    rsip.EntityData.ParentYangName = "interface-detail"
    rsip.EntityData.SegmentPath = "rsip"
    rsip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsip.EntityData.Children = types.NewOrderedMap()
    rsip.EntityData.Leafs = types.NewOrderedMap()
    rsip.EntityData.Leafs.Append("rsip-name", types.YLeaf{"RsipName", rsip.RsipName})

    rsip.EntityData.YListKeys = []string {}

    return &(rsip.EntityData)
}

// Srlg_Nodes_Node_SrlgValues
// Set of SRLG values configured
type Srlg_Nodes_Node_SrlgValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG value details . The type is slice of
    // Srlg_Nodes_Node_SrlgValues_SrlgValue.
    SrlgValue []*Srlg_Nodes_Node_SrlgValues_SrlgValue
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

    srlgValues.EntityData.Children = types.NewOrderedMap()
    srlgValues.EntityData.Children.Append("srlg-value", types.YChild{"SrlgValue", nil})
    for i := range srlgValues.SrlgValue {
        srlgValues.EntityData.Children.Append(types.GetSegmentPath(srlgValues.SrlgValue[i]), types.YChild{"SrlgValue", srlgValues.SrlgValue[i]})
    }
    srlgValues.EntityData.Leafs = types.NewOrderedMap()

    srlgValues.EntityData.YListKeys = []string {}

    return &(srlgValues.EntityData)
}

// Srlg_Nodes_Node_SrlgValues_SrlgValue
// Configured SRLG value details 
type Srlg_Nodes_Node_SrlgValues_SrlgValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (srlgValue *Srlg_Nodes_Node_SrlgValues_SrlgValue) GetEntityData() *types.CommonEntityData {
    srlgValue.EntityData.YFilter = srlgValue.YFilter
    srlgValue.EntityData.YangName = "srlg-value"
    srlgValue.EntityData.BundleName = "cisco_ios_xr"
    srlgValue.EntityData.ParentYangName = "srlg-values"
    srlgValue.EntityData.SegmentPath = "srlg-value" + types.AddKeyToken(srlgValue.Value, "value")
    srlgValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgValue.EntityData.Children = types.NewOrderedMap()
    srlgValue.EntityData.Leafs = types.NewOrderedMap()
    srlgValue.EntityData.Leafs.Append("value", types.YLeaf{"Value", srlgValue.Value})
    srlgValue.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", srlgValue.InterfaceName})

    srlgValue.EntityData.YListKeys = []string {"Value"}

    return &(srlgValue.EntityData)
}

// Srlg_Nodes_Node_InterfaceSrlgNames
// Set of SRLG names configured
type Srlg_Nodes_Node_InterfaceSrlgNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []*Srlg_Nodes_Node_InterfaceSrlgNames_InterfaceSrlgName
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

    interfaceSrlgNames.EntityData.Children = types.NewOrderedMap()
    interfaceSrlgNames.EntityData.Children.Append("interface-srlg-name", types.YChild{"InterfaceSrlgName", nil})
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        interfaceSrlgNames.EntityData.Children.Append(types.GetSegmentPath(interfaceSrlgNames.InterfaceSrlgName[i]), types.YChild{"InterfaceSrlgName", interfaceSrlgNames.InterfaceSrlgName[i]})
    }
    interfaceSrlgNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceSrlgNames.EntityData.YListKeys = []string {}

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
    interfaceSrlgName.EntityData.SegmentPath = "interface-srlg-name" + types.AddKeyToken(interfaceSrlgName.SrlgName, "srlg-name")
    interfaceSrlgName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgName.EntityData.Children = types.NewOrderedMap()
    interfaceSrlgName.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceSrlgName.Interfaces})
    interfaceSrlgName.EntityData.Leafs = types.NewOrderedMap()
    interfaceSrlgName.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", interfaceSrlgName.SrlgName})
    interfaceSrlgName.EntityData.Leafs.Append("srlg-name-xr", types.YLeaf{"SrlgNameXr", interfaceSrlgName.SrlgNameXr})
    interfaceSrlgName.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", interfaceSrlgName.SrlgValue})

    interfaceSrlgName.EntityData.YListKeys = []string {"SrlgName"}

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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Leafs = types.NewOrderedMap()
    interfaces.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaces.InterfaceName})

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Srlg_InterfaceSrlgNames
// Set of SRLG names configured
type Srlg_InterfaceSrlgNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG name details . The type is slice of
    // Srlg_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []*Srlg_InterfaceSrlgNames_InterfaceSrlgName
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

    interfaceSrlgNames.EntityData.Children = types.NewOrderedMap()
    interfaceSrlgNames.EntityData.Children.Append("interface-srlg-name", types.YChild{"InterfaceSrlgName", nil})
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        interfaceSrlgNames.EntityData.Children.Append(types.GetSegmentPath(interfaceSrlgNames.InterfaceSrlgName[i]), types.YChild{"InterfaceSrlgName", interfaceSrlgNames.InterfaceSrlgName[i]})
    }
    interfaceSrlgNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceSrlgNames.EntityData.YListKeys = []string {}

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
    interfaceSrlgName.EntityData.SegmentPath = "interface-srlg-name" + types.AddKeyToken(interfaceSrlgName.SrlgName, "srlg-name")
    interfaceSrlgName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgName.EntityData.Children = types.NewOrderedMap()
    interfaceSrlgName.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceSrlgName.Interfaces})
    interfaceSrlgName.EntityData.Leafs = types.NewOrderedMap()
    interfaceSrlgName.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", interfaceSrlgName.SrlgName})
    interfaceSrlgName.EntityData.Leafs.Append("srlg-name-xr", types.YLeaf{"SrlgNameXr", interfaceSrlgName.SrlgNameXr})
    interfaceSrlgName.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", interfaceSrlgName.SrlgValue})

    interfaceSrlgName.EntityData.YListKeys = []string {"SrlgName"}

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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Leafs = types.NewOrderedMap()
    interfaces.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaces.InterfaceName})

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Srlg_InheritNodes
// Set of inherit locations configured for SRLG
type Srlg_InheritNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG inherit location details. The type is slice of
    // Srlg_InheritNodes_InheritNode.
    InheritNode []*Srlg_InheritNodes_InheritNode
}

func (inheritNodes *Srlg_InheritNodes) GetEntityData() *types.CommonEntityData {
    inheritNodes.EntityData.YFilter = inheritNodes.YFilter
    inheritNodes.EntityData.YangName = "inherit-nodes"
    inheritNodes.EntityData.BundleName = "cisco_ios_xr"
    inheritNodes.EntityData.ParentYangName = "srlg"
    inheritNodes.EntityData.SegmentPath = "inherit-nodes"
    inheritNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNodes.EntityData.Children = types.NewOrderedMap()
    inheritNodes.EntityData.Children.Append("inherit-node", types.YChild{"InheritNode", nil})
    for i := range inheritNodes.InheritNode {
        inheritNodes.EntityData.Children.Append(types.GetSegmentPath(inheritNodes.InheritNode[i]), types.YChild{"InheritNode", inheritNodes.InheritNode[i]})
    }
    inheritNodes.EntityData.Leafs = types.NewOrderedMap()

    inheritNodes.EntityData.YListKeys = []string {}

    return &(inheritNodes.EntityData)
}

// Srlg_InheritNodes_InheritNode
// SRLG inherit location details
type Srlg_InheritNodes_InheritNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Inherit Locatio. The type is string with pattern:
    // ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    InheritNodeName interface{}

    // Inherit node name. The type is string.
    NodeName interface{}

    // Node values. The type is interface{} with range: 0..4294967295.
    NodeValues interface{}

    // SRLG attribute. The type is slice of
    // Srlg_InheritNodes_InheritNode_SrlgAttribute.
    SrlgAttribute []*Srlg_InheritNodes_InheritNode_SrlgAttribute
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetEntityData() *types.CommonEntityData {
    inheritNode.EntityData.YFilter = inheritNode.YFilter
    inheritNode.EntityData.YangName = "inherit-node"
    inheritNode.EntityData.BundleName = "cisco_ios_xr"
    inheritNode.EntityData.ParentYangName = "inherit-nodes"
    inheritNode.EntityData.SegmentPath = "inherit-node" + types.AddKeyToken(inheritNode.InheritNodeName, "inherit-node-name")
    inheritNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNode.EntityData.Children = types.NewOrderedMap()
    inheritNode.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range inheritNode.SrlgAttribute {
        inheritNode.EntityData.Children.Append(types.GetSegmentPath(inheritNode.SrlgAttribute[i]), types.YChild{"SrlgAttribute", inheritNode.SrlgAttribute[i]})
    }
    inheritNode.EntityData.Leafs = types.NewOrderedMap()
    inheritNode.EntityData.Leafs.Append("inherit-node-name", types.YLeaf{"InheritNodeName", inheritNode.InheritNodeName})
    inheritNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", inheritNode.NodeName})
    inheritNode.EntityData.Leafs.Append("node-values", types.YLeaf{"NodeValues", inheritNode.NodeValues})

    inheritNode.EntityData.YListKeys = []string {"InheritNodeName"}

    return &(inheritNode.EntityData)
}

// Srlg_InheritNodes_InheritNode_SrlgAttribute
// SRLG attribute
type Srlg_InheritNodes_InheritNode_SrlgAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG value. The type is interface{} with range: 0..4294967295.
    SrlgValue interface{}

    // Priority. The type is Priority.
    Priority interface{}

    // Index. The type is interface{} with range: 0..65535.
    SrlgIndex interface{}
}

func (srlgAttribute *Srlg_InheritNodes_InheritNode_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "inherit-node"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_SrlgValues
// Set of SRLG values configured
type Srlg_SrlgValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured SRLG value details . The type is slice of
    // Srlg_SrlgValues_SrlgValue.
    SrlgValue []*Srlg_SrlgValues_SrlgValue
}

func (srlgValues *Srlg_SrlgValues) GetEntityData() *types.CommonEntityData {
    srlgValues.EntityData.YFilter = srlgValues.YFilter
    srlgValues.EntityData.YangName = "srlg-values"
    srlgValues.EntityData.BundleName = "cisco_ios_xr"
    srlgValues.EntityData.ParentYangName = "srlg"
    srlgValues.EntityData.SegmentPath = "srlg-values"
    srlgValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgValues.EntityData.Children = types.NewOrderedMap()
    srlgValues.EntityData.Children.Append("srlg-value", types.YChild{"SrlgValue", nil})
    for i := range srlgValues.SrlgValue {
        srlgValues.EntityData.Children.Append(types.GetSegmentPath(srlgValues.SrlgValue[i]), types.YChild{"SrlgValue", srlgValues.SrlgValue[i]})
    }
    srlgValues.EntityData.Leafs = types.NewOrderedMap()

    srlgValues.EntityData.YListKeys = []string {}

    return &(srlgValues.EntityData)
}

// Srlg_SrlgValues_SrlgValue
// Configured SRLG value details 
type Srlg_SrlgValues_SrlgValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Interface name. The type is slice of string.
    InterfaceName []interface{}
}

func (srlgValue *Srlg_SrlgValues_SrlgValue) GetEntityData() *types.CommonEntityData {
    srlgValue.EntityData.YFilter = srlgValue.YFilter
    srlgValue.EntityData.YangName = "srlg-value"
    srlgValue.EntityData.BundleName = "cisco_ios_xr"
    srlgValue.EntityData.ParentYangName = "srlg-values"
    srlgValue.EntityData.SegmentPath = "srlg-value" + types.AddKeyToken(srlgValue.Value, "value")
    srlgValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgValue.EntityData.Children = types.NewOrderedMap()
    srlgValue.EntityData.Leafs = types.NewOrderedMap()
    srlgValue.EntityData.Leafs.Append("value", types.YLeaf{"Value", srlgValue.Value})
    srlgValue.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", srlgValue.InterfaceName})

    srlgValue.EntityData.YListKeys = []string {"Value"}

    return &(srlgValue.EntityData)
}

// Srlg_InterfaceDetails
// Set of interfaces configured for SRLG
type Srlg_InterfaceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG interface details. The type is slice of
    // Srlg_InterfaceDetails_InterfaceDetail.
    InterfaceDetail []*Srlg_InterfaceDetails_InterfaceDetail
}

func (interfaceDetails *Srlg_InterfaceDetails) GetEntityData() *types.CommonEntityData {
    interfaceDetails.EntityData.YFilter = interfaceDetails.YFilter
    interfaceDetails.EntityData.YangName = "interface-details"
    interfaceDetails.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetails.EntityData.ParentYangName = "srlg"
    interfaceDetails.EntityData.SegmentPath = "interface-details"
    interfaceDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetails.EntityData.Children = types.NewOrderedMap()
    interfaceDetails.EntityData.Children.Append("interface-detail", types.YChild{"InterfaceDetail", nil})
    for i := range interfaceDetails.InterfaceDetail {
        interfaceDetails.EntityData.Children.Append(types.GetSegmentPath(interfaceDetails.InterfaceDetail[i]), types.YChild{"InterfaceDetail", interfaceDetails.InterfaceDetail[i]})
    }
    interfaceDetails.EntityData.Leafs = types.NewOrderedMap()

    interfaceDetails.EntityData.YListKeys = []string {}

    return &(interfaceDetails.EntityData)
}

// Srlg_InterfaceDetails_InterfaceDetail
// SRLG interface details
type Srlg_InterfaceDetails_InterfaceDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Groups. The type is interface{} with range: 0..4294967295.
    Groups interface{}

    // Nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // SRLG attributes. The type is slice of
    // Srlg_InterfaceDetails_InterfaceDetail_SrlgAttribute.
    SrlgAttribute []*Srlg_InterfaceDetails_InterfaceDetail_SrlgAttribute

    // rsip list. The type is slice of Srlg_InterfaceDetails_InterfaceDetail_Rsip.
    Rsip []*Srlg_InterfaceDetails_InterfaceDetail_Rsip
}

func (interfaceDetail *Srlg_InterfaceDetails_InterfaceDetail) GetEntityData() *types.CommonEntityData {
    interfaceDetail.EntityData.YFilter = interfaceDetail.YFilter
    interfaceDetail.EntityData.YangName = "interface-detail"
    interfaceDetail.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetail.EntityData.ParentYangName = "interface-details"
    interfaceDetail.EntityData.SegmentPath = "interface-detail" + types.AddKeyToken(interfaceDetail.InterfaceName, "interface-name")
    interfaceDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetail.EntityData.Children = types.NewOrderedMap()
    interfaceDetail.EntityData.Children.Append("srlg-attribute", types.YChild{"SrlgAttribute", nil})
    for i := range interfaceDetail.SrlgAttribute {
        interfaceDetail.EntityData.Children.Append(types.GetSegmentPath(interfaceDetail.SrlgAttribute[i]), types.YChild{"SrlgAttribute", interfaceDetail.SrlgAttribute[i]})
    }
    interfaceDetail.EntityData.Children.Append("rsip", types.YChild{"Rsip", nil})
    for i := range interfaceDetail.Rsip {
        interfaceDetail.EntityData.Children.Append(types.GetSegmentPath(interfaceDetail.Rsip[i]), types.YChild{"Rsip", interfaceDetail.Rsip[i]})
    }
    interfaceDetail.EntityData.Leafs = types.NewOrderedMap()
    interfaceDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceDetail.InterfaceName})
    interfaceDetail.EntityData.Leafs.Append("groups", types.YLeaf{"Groups", interfaceDetail.Groups})
    interfaceDetail.EntityData.Leafs.Append("nodes", types.YLeaf{"Nodes", interfaceDetail.Nodes})

    interfaceDetail.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceDetail.EntityData)
}

// Srlg_InterfaceDetails_InterfaceDetail_SrlgAttribute
// SRLG attributes
type Srlg_InterfaceDetails_InterfaceDetail_SrlgAttribute struct {
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

func (srlgAttribute *Srlg_InterfaceDetails_InterfaceDetail_SrlgAttribute) GetEntityData() *types.CommonEntityData {
    srlgAttribute.EntityData.YFilter = srlgAttribute.YFilter
    srlgAttribute.EntityData.YangName = "srlg-attribute"
    srlgAttribute.EntityData.BundleName = "cisco_ios_xr"
    srlgAttribute.EntityData.ParentYangName = "interface-detail"
    srlgAttribute.EntityData.SegmentPath = "srlg-attribute"
    srlgAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgAttribute.EntityData.Children = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs = types.NewOrderedMap()
    srlgAttribute.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgAttribute.SrlgValue})
    srlgAttribute.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", srlgAttribute.Priority})
    srlgAttribute.EntityData.Leafs.Append("source", types.YLeaf{"Source", srlgAttribute.Source})
    srlgAttribute.EntityData.Leafs.Append("source-name", types.YLeaf{"SourceName", srlgAttribute.SourceName})
    srlgAttribute.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", srlgAttribute.SrlgIndex})

    srlgAttribute.EntityData.YListKeys = []string {}

    return &(srlgAttribute.EntityData)
}

// Srlg_InterfaceDetails_InterfaceDetail_Rsip
// rsip list
type Srlg_InterfaceDetails_InterfaceDetail_Rsip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of names matching rsip. The type is string.
    RsipName interface{}
}

func (rsip *Srlg_InterfaceDetails_InterfaceDetail_Rsip) GetEntityData() *types.CommonEntityData {
    rsip.EntityData.YFilter = rsip.YFilter
    rsip.EntityData.YangName = "rsip"
    rsip.EntityData.BundleName = "cisco_ios_xr"
    rsip.EntityData.ParentYangName = "interface-detail"
    rsip.EntityData.SegmentPath = "rsip"
    rsip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsip.EntityData.Children = types.NewOrderedMap()
    rsip.EntityData.Leafs = types.NewOrderedMap()
    rsip.EntityData.Leafs.Append("rsip-name", types.YLeaf{"RsipName", rsip.RsipName})

    rsip.EntityData.YListKeys = []string {}

    return &(rsip.EntityData)
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

    selectiveVrfDownload.EntityData.Children = types.NewOrderedMap()
    selectiveVrfDownload.EntityData.Children.Append("state", types.YChild{"State", &selectiveVrfDownload.State})
    selectiveVrfDownload.EntityData.Leafs = types.NewOrderedMap()

    selectiveVrfDownload.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("is-svd-enabled", types.YLeaf{"IsSvdEnabled", state.IsSvdEnabled})
    state.EntityData.Leafs.Append("is-svd-enabled-cfg", types.YLeaf{"IsSvdEnabledCfg", state.IsSvdEnabledCfg})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

