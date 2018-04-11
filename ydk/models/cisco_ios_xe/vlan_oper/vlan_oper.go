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
    Vlan []Vlans_Vlan
}

func (vlans *Vlans) GetEntityData() *types.CommonEntityData {
    vlans.EntityData.YFilter = vlans.YFilter
    vlans.EntityData.YangName = "vlans"
    vlans.EntityData.BundleName = "cisco_ios_xe"
    vlans.EntityData.ParentYangName = "Cisco-IOS-XE-vlan-oper"
    vlans.EntityData.SegmentPath = "Cisco-IOS-XE-vlan-oper:vlans"
    vlans.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlans.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlans.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlans.EntityData.Children = make(map[string]types.YChild)
    vlans.EntityData.Children["vlan"] = types.YChild{"Vlan", nil}
    for i := range vlans.Vlan {
        vlans.EntityData.Children[types.GetSegmentPath(&vlans.Vlan[i])] = types.YChild{"Vlan", &vlans.Vlan[i]}
    }
    vlans.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlans.EntityData)
}

// Vlans_Vlan
// List of VLANs, keyed by id
type Vlans_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VLAN id. The type is interface{} with range:
    // 0..65535.
    Id interface{}

    // VLAN name. The type is string.
    Name interface{}

    // VLAN status. The type is VlanStatusType.
    Status interface{}

    // Assigned ports. The type is slice of Vlans_Vlan_Ports.
    Ports []Vlans_Vlan_Ports
}

func (vlan *Vlans_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xe"
    vlan.EntityData.ParentYangName = "vlans"
    vlan.EntityData.SegmentPath = "vlan" + "[id='" + fmt.Sprintf("%v", vlan.Id) + "']"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlan.EntityData.Children = make(map[string]types.YChild)
    vlan.EntityData.Children["ports"] = types.YChild{"Ports", nil}
    for i := range vlan.Ports {
        vlan.EntityData.Children[types.GetSegmentPath(&vlan.Ports[i])] = types.YChild{"Ports", &vlan.Ports[i]}
    }
    vlan.EntityData.Leafs = make(map[string]types.YLeaf)
    vlan.EntityData.Leafs["id"] = types.YLeaf{"Id", vlan.Id}
    vlan.EntityData.Leafs["name"] = types.YLeaf{"Name", vlan.Name}
    vlan.EntityData.Leafs["status"] = types.YLeaf{"Status", vlan.Status}
    return &(vlan.EntityData)
}

// Vlans_Vlan_Ports
// Assigned ports
type Vlans_Vlan_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assigned interface. The type is string.
    Interface_ interface{}

    // Assigned subinterface. The type is interface{} with range: 0..4294967295.
    Subinterface interface{}
}

func (ports *Vlans_Vlan_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xe"
    ports.EntityData.ParentYangName = "vlan"
    ports.EntityData.SegmentPath = "ports"
    ports.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ports.EntityData.Children = make(map[string]types.YChild)
    ports.EntityData.Leafs = make(map[string]types.YLeaf)
    ports.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", ports.Interface_}
    ports.EntityData.Leafs["subinterface"] = types.YLeaf{"Subinterface", ports.Subinterface}
    return &(ports.EntityData)
}

