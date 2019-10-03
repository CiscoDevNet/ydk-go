// This module contains a collection of YANG definitions for
// monitoring vlans in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package vlan_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package vlan_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-vlan-oper vlans}", reflect.TypeOf(Vlans{}))
    ydk.RegisterEntity("Cisco-IOS-XE-vlan-oper:vlans", reflect.TypeOf(Vlans{}))
}

// VlanStatusType represents Operational state of the VLAN
type VlanStatusType string

const (
    VlanStatusType_active VlanStatusType = "active"

    VlanStatusType_suspend VlanStatusType = "suspend"
)

// Vlans
// Informaton about VLANs
type Vlans struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VLANs, keyed by id. The type is slice of Vlans_Vlan.
    Vlan []*Vlans_Vlan
}

func (vlans *Vlans) GetEntityData() *types.CommonEntityData {
    vlans.EntityData.YFilter = vlans.YFilter
    vlans.EntityData.YangName = "vlans"
    vlans.EntityData.BundleName = "cisco_ios_xe"
    vlans.EntityData.ParentYangName = "Cisco-IOS-XE-vlan-oper"
    vlans.EntityData.SegmentPath = "Cisco-IOS-XE-vlan-oper:vlans"
    vlans.EntityData.AbsolutePath = vlans.EntityData.SegmentPath
    vlans.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlans.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlans.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlans.EntityData.Children = types.NewOrderedMap()
    vlans.EntityData.Children.Append("vlan", types.YChild{"Vlan", nil})
    for i := range vlans.Vlan {
        vlans.EntityData.Children.Append(types.GetSegmentPath(vlans.Vlan[i]), types.YChild{"Vlan", vlans.Vlan[i]})
    }
    vlans.EntityData.Leafs = types.NewOrderedMap()

    vlans.EntityData.YListKeys = []string {}

    return &(vlans.EntityData)
}

// Vlans_Vlan
// List of VLANs, keyed by id
type Vlans_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VLAN id. The type is interface{} with range:
    // 0..65535.
    Id interface{}

    // VLAN name. The type is string.
    Name interface{}

    // VLAN status. The type is VlanStatusType.
    Status interface{}

    // Assigned ports. The type is slice of Vlans_Vlan_Ports.
    Ports []*Vlans_Vlan_Ports

    // List of interfaces for a given VLAN. The type is slice of
    // Vlans_Vlan_VlanInterfaces.
    VlanInterfaces []*Vlans_Vlan_VlanInterfaces
}

func (vlan *Vlans_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xe"
    vlan.EntityData.ParentYangName = "vlans"
    vlan.EntityData.SegmentPath = "vlan" + types.AddKeyToken(vlan.Id, "id")
    vlan.EntityData.AbsolutePath = "Cisco-IOS-XE-vlan-oper:vlans/" + vlan.EntityData.SegmentPath
    vlan.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("ports", types.YChild{"Ports", nil})
    for i := range vlan.Ports {
        types.SetYListKey(vlan.Ports[i], i)
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.Ports[i]), types.YChild{"Ports", vlan.Ports[i]})
    }
    vlan.EntityData.Children.Append("vlan-interfaces", types.YChild{"VlanInterfaces", nil})
    for i := range vlan.VlanInterfaces {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.VlanInterfaces[i]), types.YChild{"VlanInterfaces", vlan.VlanInterfaces[i]})
    }
    vlan.EntityData.Leafs = types.NewOrderedMap()
    vlan.EntityData.Leafs.Append("id", types.YLeaf{"Id", vlan.Id})
    vlan.EntityData.Leafs.Append("name", types.YLeaf{"Name", vlan.Name})
    vlan.EntityData.Leafs.Append("status", types.YLeaf{"Status", vlan.Status})

    vlan.EntityData.YListKeys = []string {"Id"}

    return &(vlan.EntityData)
}

// Vlans_Vlan_Ports
// Assigned ports
type Vlans_Vlan_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Assigned interface. The type is string.
    Interface interface{}

    // Assigned subinterface. The type is interface{} with range: 0..4294967295.
    Subinterface interface{}
}

func (ports *Vlans_Vlan_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xe"
    ports.EntityData.ParentYangName = "vlan"
    ports.EntityData.SegmentPath = "ports" + types.AddNoKeyToken(ports)
    ports.EntityData.AbsolutePath = "Cisco-IOS-XE-vlan-oper:vlans/vlan/" + ports.EntityData.SegmentPath
    ports.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ports.EntityData.Children = types.NewOrderedMap()
    ports.EntityData.Leafs = types.NewOrderedMap()
    ports.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", ports.Interface})
    ports.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", ports.Subinterface})

    ports.EntityData.YListKeys = []string {}

    return &(ports.EntityData)
}

// Vlans_Vlan_VlanInterfaces
// List of interfaces for a given VLAN
type Vlans_Vlan_VlanInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Assigned interface to the vlan. The type is
    // string.
    Interface interface{}

    // Assigned subinterface to the vlan. The type is interface{} with range:
    // 0..4294967295.
    Subinterface interface{}
}

func (vlanInterfaces *Vlans_Vlan_VlanInterfaces) GetEntityData() *types.CommonEntityData {
    vlanInterfaces.EntityData.YFilter = vlanInterfaces.YFilter
    vlanInterfaces.EntityData.YangName = "vlan-interfaces"
    vlanInterfaces.EntityData.BundleName = "cisco_ios_xe"
    vlanInterfaces.EntityData.ParentYangName = "vlan"
    vlanInterfaces.EntityData.SegmentPath = "vlan-interfaces" + types.AddKeyToken(vlanInterfaces.Interface, "interface")
    vlanInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XE-vlan-oper:vlans/vlan/" + vlanInterfaces.EntityData.SegmentPath
    vlanInterfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlanInterfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlanInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlanInterfaces.EntityData.Children = types.NewOrderedMap()
    vlanInterfaces.EntityData.Leafs = types.NewOrderedMap()
    vlanInterfaces.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", vlanInterfaces.Interface})
    vlanInterfaces.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", vlanInterfaces.Subinterface})

    vlanInterfaces.EntityData.YListKeys = []string {"Interface"}

    return &(vlanInterfaces.EntityData)
}

