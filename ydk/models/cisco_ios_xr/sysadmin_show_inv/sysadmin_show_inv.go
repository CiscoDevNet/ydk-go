// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Calvados Inventory Service maintain entity database
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_show_inv

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_show_inv"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-show-inv inventory}", reflect.TypeOf(Inventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-show-inv:inventory", reflect.TypeOf(Inventory{}))
}

// Inventory
// show inventory
type Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Inventory_Location.
    Location []Inventory_Location

    // The type is slice of Inventory_All.
    All []Inventory_All

    // The type is slice of Inventory_Chassis.
    Chassis []Inventory_Chassis

    // The type is slice of Inventory_Power.
    Power []Inventory_Power

    // The type is slice of Inventory_Fan.
    Fan []Inventory_Fan

    // The type is slice of Inventory_Raw.
    Raw []Inventory_Raw
}

func (inventory *Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-show-inv"
    inventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-show-inv:inventory"
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = make(map[string]types.YChild)
    inventory.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range inventory.Location {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Location[i])] = types.YChild{"Location", &inventory.Location[i]}
    }
    inventory.EntityData.Children["all"] = types.YChild{"All", nil}
    for i := range inventory.All {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.All[i])] = types.YChild{"All", &inventory.All[i]}
    }
    inventory.EntityData.Children["chassis"] = types.YChild{"Chassis", nil}
    for i := range inventory.Chassis {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Chassis[i])] = types.YChild{"Chassis", &inventory.Chassis[i]}
    }
    inventory.EntityData.Children["power"] = types.YChild{"Power", nil}
    for i := range inventory.Power {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Power[i])] = types.YChild{"Power", &inventory.Power[i]}
    }
    inventory.EntityData.Children["fan"] = types.YChild{"Fan", nil}
    for i := range inventory.Fan {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Fan[i])] = types.YChild{"Fan", &inventory.Fan[i]}
    }
    inventory.EntityData.Children["raw"] = types.YChild{"Raw", nil}
    for i := range inventory.Raw {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Raw[i])] = types.YChild{"Raw", &inventory.Raw[i]}
    }
    inventory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inventory.EntityData)
}

// Inventory_Location
type Inventory_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node id of the entity. The type is string.
    Loc interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    Pid interface{}

    // Version ID of the entity. The type is string.
    Vid interface{}

    // Serial Numbe of the entity. The type is string.
    Sn interface{}

    // Index for the entity. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (location *Inventory_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "inventory"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    location.EntityData.Leafs["name"] = types.YLeaf{"Name", location.Name}
    location.EntityData.Leafs["Description"] = types.YLeaf{"Description", location.Description}
    location.EntityData.Leafs["PID"] = types.YLeaf{"Pid", location.Pid}
    location.EntityData.Leafs["VID"] = types.YLeaf{"Vid", location.Vid}
    location.EntityData.Leafs["SN"] = types.YLeaf{"Sn", location.Sn}
    location.EntityData.Leafs["index"] = types.YLeaf{"Index", location.Index}
    return &(location.EntityData)
}

// Inventory_All
type Inventory_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    Pid interface{}

    // Version ID of the entity. The type is string.
    Vid interface{}

    // Serial Numbe of the entity. The type is string.
    Sn interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (all *Inventory_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "inventory"
    all.EntityData.SegmentPath = "all" + "[index='" + fmt.Sprintf("%v", all.Index) + "']"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    all.EntityData.Leafs["index"] = types.YLeaf{"Index", all.Index}
    all.EntityData.Leafs["name"] = types.YLeaf{"Name", all.Name}
    all.EntityData.Leafs["Description"] = types.YLeaf{"Description", all.Description}
    all.EntityData.Leafs["PID"] = types.YLeaf{"Pid", all.Pid}
    all.EntityData.Leafs["VID"] = types.YLeaf{"Vid", all.Vid}
    all.EntityData.Leafs["SN"] = types.YLeaf{"Sn", all.Sn}
    all.EntityData.Leafs["loc"] = types.YLeaf{"Loc", all.Loc}
    return &(all.EntityData)
}

// Inventory_Chassis
type Inventory_Chassis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    Pid interface{}

    // Version ID of the entity. The type is string.
    Vid interface{}

    // Serial Numbe of the entity. The type is string.
    Sn interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (chassis *Inventory_Chassis) GetEntityData() *types.CommonEntityData {
    chassis.EntityData.YFilter = chassis.YFilter
    chassis.EntityData.YangName = "chassis"
    chassis.EntityData.BundleName = "cisco_ios_xr"
    chassis.EntityData.ParentYangName = "inventory"
    chassis.EntityData.SegmentPath = "chassis" + "[index='" + fmt.Sprintf("%v", chassis.Index) + "']"
    chassis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassis.EntityData.Children = make(map[string]types.YChild)
    chassis.EntityData.Leafs = make(map[string]types.YLeaf)
    chassis.EntityData.Leafs["index"] = types.YLeaf{"Index", chassis.Index}
    chassis.EntityData.Leafs["name"] = types.YLeaf{"Name", chassis.Name}
    chassis.EntityData.Leafs["Description"] = types.YLeaf{"Description", chassis.Description}
    chassis.EntityData.Leafs["PID"] = types.YLeaf{"Pid", chassis.Pid}
    chassis.EntityData.Leafs["VID"] = types.YLeaf{"Vid", chassis.Vid}
    chassis.EntityData.Leafs["SN"] = types.YLeaf{"Sn", chassis.Sn}
    chassis.EntityData.Leafs["loc"] = types.YLeaf{"Loc", chassis.Loc}
    return &(chassis.EntityData)
}

// Inventory_Power
type Inventory_Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    Pid interface{}

    // Version ID of the entity. The type is string.
    Vid interface{}

    // Serial Numbe of the entity. The type is string.
    Sn interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (power *Inventory_Power) GetEntityData() *types.CommonEntityData {
    power.EntityData.YFilter = power.YFilter
    power.EntityData.YangName = "power"
    power.EntityData.BundleName = "cisco_ios_xr"
    power.EntityData.ParentYangName = "inventory"
    power.EntityData.SegmentPath = "power" + "[index='" + fmt.Sprintf("%v", power.Index) + "']"
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = make(map[string]types.YChild)
    power.EntityData.Leafs = make(map[string]types.YLeaf)
    power.EntityData.Leafs["index"] = types.YLeaf{"Index", power.Index}
    power.EntityData.Leafs["name"] = types.YLeaf{"Name", power.Name}
    power.EntityData.Leafs["Description"] = types.YLeaf{"Description", power.Description}
    power.EntityData.Leafs["PID"] = types.YLeaf{"Pid", power.Pid}
    power.EntityData.Leafs["VID"] = types.YLeaf{"Vid", power.Vid}
    power.EntityData.Leafs["SN"] = types.YLeaf{"Sn", power.Sn}
    power.EntityData.Leafs["loc"] = types.YLeaf{"Loc", power.Loc}
    return &(power.EntityData)
}

// Inventory_Fan
type Inventory_Fan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    Pid interface{}

    // Version ID of the entity. The type is string.
    Vid interface{}

    // Serial Numbe of the entity. The type is string.
    Sn interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (fan *Inventory_Fan) GetEntityData() *types.CommonEntityData {
    fan.EntityData.YFilter = fan.YFilter
    fan.EntityData.YangName = "fan"
    fan.EntityData.BundleName = "cisco_ios_xr"
    fan.EntityData.ParentYangName = "inventory"
    fan.EntityData.SegmentPath = "fan" + "[index='" + fmt.Sprintf("%v", fan.Index) + "']"
    fan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fan.EntityData.Children = make(map[string]types.YChild)
    fan.EntityData.Leafs = make(map[string]types.YLeaf)
    fan.EntityData.Leafs["index"] = types.YLeaf{"Index", fan.Index}
    fan.EntityData.Leafs["name"] = types.YLeaf{"Name", fan.Name}
    fan.EntityData.Leafs["Description"] = types.YLeaf{"Description", fan.Description}
    fan.EntityData.Leafs["PID"] = types.YLeaf{"Pid", fan.Pid}
    fan.EntityData.Leafs["VID"] = types.YLeaf{"Vid", fan.Vid}
    fan.EntityData.Leafs["SN"] = types.YLeaf{"Sn", fan.Sn}
    fan.EntityData.Leafs["loc"] = types.YLeaf{"Loc", fan.Loc}
    return &(fan.EntityData)
}

// Inventory_Raw
type Inventory_Raw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    Pid interface{}

    // Version ID of the entity. The type is string.
    Vid interface{}

    // Serial Numbe of the entity. The type is string.
    Sn interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (raw *Inventory_Raw) GetEntityData() *types.CommonEntityData {
    raw.EntityData.YFilter = raw.YFilter
    raw.EntityData.YangName = "raw"
    raw.EntityData.BundleName = "cisco_ios_xr"
    raw.EntityData.ParentYangName = "inventory"
    raw.EntityData.SegmentPath = "raw" + "[index='" + fmt.Sprintf("%v", raw.Index) + "']"
    raw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raw.EntityData.Children = make(map[string]types.YChild)
    raw.EntityData.Leafs = make(map[string]types.YLeaf)
    raw.EntityData.Leafs["index"] = types.YLeaf{"Index", raw.Index}
    raw.EntityData.Leafs["name"] = types.YLeaf{"Name", raw.Name}
    raw.EntityData.Leafs["Description"] = types.YLeaf{"Description", raw.Description}
    raw.EntityData.Leafs["PID"] = types.YLeaf{"Pid", raw.Pid}
    raw.EntityData.Leafs["VID"] = types.YLeaf{"Vid", raw.Vid}
    raw.EntityData.Leafs["SN"] = types.YLeaf{"Sn", raw.Sn}
    raw.EntityData.Leafs["loc"] = types.YLeaf{"Loc", raw.Loc}
    return &(raw.EntityData)
}

