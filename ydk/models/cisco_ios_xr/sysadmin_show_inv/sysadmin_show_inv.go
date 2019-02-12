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
    Location []*Inventory_Location

    // The type is slice of Inventory_All.
    All []*Inventory_All

    // The type is slice of Inventory_Chassis.
    Chassis []*Inventory_Chassis

    // The type is slice of Inventory_Power.
    Power []*Inventory_Power

    // The type is slice of Inventory_Fan.
    Fan []*Inventory_Fan

    // The type is slice of Inventory_Raw.
    Raw []*Inventory_Raw
}

func (inventory *Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-show-inv"
    inventory.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-show-inv:inventory"
    inventory.EntityData.AbsolutePath = inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range inventory.Location {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.Location[i]), types.YChild{"Location", inventory.Location[i]})
    }
    inventory.EntityData.Children.Append("all", types.YChild{"All", nil})
    for i := range inventory.All {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.All[i]), types.YChild{"All", inventory.All[i]})
    }
    inventory.EntityData.Children.Append("chassis", types.YChild{"Chassis", nil})
    for i := range inventory.Chassis {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.Chassis[i]), types.YChild{"Chassis", inventory.Chassis[i]})
    }
    inventory.EntityData.Children.Append("power", types.YChild{"Power", nil})
    for i := range inventory.Power {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.Power[i]), types.YChild{"Power", inventory.Power[i]})
    }
    inventory.EntityData.Children.Append("fan", types.YChild{"Fan", nil})
    for i := range inventory.Fan {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.Fan[i]), types.YChild{"Fan", inventory.Fan[i]})
    }
    inventory.EntityData.Children.Append("raw", types.YChild{"Raw", nil})
    for i := range inventory.Raw {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.Raw[i]), types.YChild{"Raw", inventory.Raw[i]})
    }
    inventory.EntityData.Leafs = types.NewOrderedMap()

    inventory.EntityData.YListKeys = []string {}

    return &(inventory.EntityData)
}

// Inventory_Location
type Inventory_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node id of the entity. The type is string.
    Loc interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    PID interface{}

    // Version ID of the entity. The type is string.
    VID interface{}

    // Serial Numbe of the entity. The type is string.
    SN interface{}

    // Index for the entity. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (location *Inventory_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "inventory"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Loc, "loc")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-inv:inventory/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", location.Loc})
    location.EntityData.Leafs.Append("name", types.YLeaf{"Name", location.Name})
    location.EntityData.Leafs.Append("Description", types.YLeaf{"Description", location.Description})
    location.EntityData.Leafs.Append("PID", types.YLeaf{"PID", location.PID})
    location.EntityData.Leafs.Append("VID", types.YLeaf{"VID", location.VID})
    location.EntityData.Leafs.Append("SN", types.YLeaf{"SN", location.SN})
    location.EntityData.Leafs.Append("index", types.YLeaf{"Index", location.Index})

    location.EntityData.YListKeys = []string {"Loc"}

    return &(location.EntityData)
}

// Inventory_All
type Inventory_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    PID interface{}

    // Version ID of the entity. The type is string.
    VID interface{}

    // Serial Numbe of the entity. The type is string.
    SN interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (all *Inventory_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "inventory"
    all.EntityData.SegmentPath = "all" + types.AddKeyToken(all.Index, "index")
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-inv:inventory/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("index", types.YLeaf{"Index", all.Index})
    all.EntityData.Leafs.Append("name", types.YLeaf{"Name", all.Name})
    all.EntityData.Leafs.Append("Description", types.YLeaf{"Description", all.Description})
    all.EntityData.Leafs.Append("PID", types.YLeaf{"PID", all.PID})
    all.EntityData.Leafs.Append("VID", types.YLeaf{"VID", all.VID})
    all.EntityData.Leafs.Append("SN", types.YLeaf{"SN", all.SN})
    all.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", all.Loc})

    all.EntityData.YListKeys = []string {"Index"}

    return &(all.EntityData)
}

// Inventory_Chassis
type Inventory_Chassis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    PID interface{}

    // Version ID of the entity. The type is string.
    VID interface{}

    // Serial Numbe of the entity. The type is string.
    SN interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (chassis *Inventory_Chassis) GetEntityData() *types.CommonEntityData {
    chassis.EntityData.YFilter = chassis.YFilter
    chassis.EntityData.YangName = "chassis"
    chassis.EntityData.BundleName = "cisco_ios_xr"
    chassis.EntityData.ParentYangName = "inventory"
    chassis.EntityData.SegmentPath = "chassis" + types.AddKeyToken(chassis.Index, "index")
    chassis.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-inv:inventory/" + chassis.EntityData.SegmentPath
    chassis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassis.EntityData.Children = types.NewOrderedMap()
    chassis.EntityData.Leafs = types.NewOrderedMap()
    chassis.EntityData.Leafs.Append("index", types.YLeaf{"Index", chassis.Index})
    chassis.EntityData.Leafs.Append("name", types.YLeaf{"Name", chassis.Name})
    chassis.EntityData.Leafs.Append("Description", types.YLeaf{"Description", chassis.Description})
    chassis.EntityData.Leafs.Append("PID", types.YLeaf{"PID", chassis.PID})
    chassis.EntityData.Leafs.Append("VID", types.YLeaf{"VID", chassis.VID})
    chassis.EntityData.Leafs.Append("SN", types.YLeaf{"SN", chassis.SN})
    chassis.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", chassis.Loc})

    chassis.EntityData.YListKeys = []string {"Index"}

    return &(chassis.EntityData)
}

// Inventory_Power
type Inventory_Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    PID interface{}

    // Version ID of the entity. The type is string.
    VID interface{}

    // Serial Numbe of the entity. The type is string.
    SN interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (power *Inventory_Power) GetEntityData() *types.CommonEntityData {
    power.EntityData.YFilter = power.YFilter
    power.EntityData.YangName = "power"
    power.EntityData.BundleName = "cisco_ios_xr"
    power.EntityData.ParentYangName = "inventory"
    power.EntityData.SegmentPath = "power" + types.AddKeyToken(power.Index, "index")
    power.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-inv:inventory/" + power.EntityData.SegmentPath
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = types.NewOrderedMap()
    power.EntityData.Leafs = types.NewOrderedMap()
    power.EntityData.Leafs.Append("index", types.YLeaf{"Index", power.Index})
    power.EntityData.Leafs.Append("name", types.YLeaf{"Name", power.Name})
    power.EntityData.Leafs.Append("Description", types.YLeaf{"Description", power.Description})
    power.EntityData.Leafs.Append("PID", types.YLeaf{"PID", power.PID})
    power.EntityData.Leafs.Append("VID", types.YLeaf{"VID", power.VID})
    power.EntityData.Leafs.Append("SN", types.YLeaf{"SN", power.SN})
    power.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", power.Loc})

    power.EntityData.YListKeys = []string {"Index"}

    return &(power.EntityData)
}

// Inventory_Fan
type Inventory_Fan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    PID interface{}

    // Version ID of the entity. The type is string.
    VID interface{}

    // Serial Numbe of the entity. The type is string.
    SN interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (fan *Inventory_Fan) GetEntityData() *types.CommonEntityData {
    fan.EntityData.YFilter = fan.YFilter
    fan.EntityData.YangName = "fan"
    fan.EntityData.BundleName = "cisco_ios_xr"
    fan.EntityData.ParentYangName = "inventory"
    fan.EntityData.SegmentPath = "fan" + types.AddKeyToken(fan.Index, "index")
    fan.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-inv:inventory/" + fan.EntityData.SegmentPath
    fan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fan.EntityData.Children = types.NewOrderedMap()
    fan.EntityData.Leafs = types.NewOrderedMap()
    fan.EntityData.Leafs.Append("index", types.YLeaf{"Index", fan.Index})
    fan.EntityData.Leafs.Append("name", types.YLeaf{"Name", fan.Name})
    fan.EntityData.Leafs.Append("Description", types.YLeaf{"Description", fan.Description})
    fan.EntityData.Leafs.Append("PID", types.YLeaf{"PID", fan.PID})
    fan.EntityData.Leafs.Append("VID", types.YLeaf{"VID", fan.VID})
    fan.EntityData.Leafs.Append("SN", types.YLeaf{"SN", fan.SN})
    fan.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", fan.Loc})

    fan.EntityData.YListKeys = []string {"Index"}

    return &(fan.EntityData)
}

// Inventory_Raw
type Inventory_Raw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index for the entity. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // Name of the entity. The type is string.
    Name interface{}

    // Description of the entity. The type is string.
    Description interface{}

    // Product ID of the entity. The type is string.
    PID interface{}

    // Version ID of the entity. The type is string.
    VID interface{}

    // Serial Numbe of the entity. The type is string.
    SN interface{}

    // Node id of the entity. The type is string.
    Loc interface{}
}

func (raw *Inventory_Raw) GetEntityData() *types.CommonEntityData {
    raw.EntityData.YFilter = raw.YFilter
    raw.EntityData.YangName = "raw"
    raw.EntityData.BundleName = "cisco_ios_xr"
    raw.EntityData.ParentYangName = "inventory"
    raw.EntityData.SegmentPath = "raw" + types.AddKeyToken(raw.Index, "index")
    raw.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-inv:inventory/" + raw.EntityData.SegmentPath
    raw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raw.EntityData.Children = types.NewOrderedMap()
    raw.EntityData.Leafs = types.NewOrderedMap()
    raw.EntityData.Leafs.Append("index", types.YLeaf{"Index", raw.Index})
    raw.EntityData.Leafs.Append("name", types.YLeaf{"Name", raw.Name})
    raw.EntityData.Leafs.Append("Description", types.YLeaf{"Description", raw.Description})
    raw.EntityData.Leafs.Append("PID", types.YLeaf{"PID", raw.PID})
    raw.EntityData.Leafs.Append("VID", types.YLeaf{"VID", raw.VID})
    raw.EntityData.Leafs.Append("SN", types.YLeaf{"SN", raw.SN})
    raw.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", raw.Loc})

    raw.EntityData.YListKeys = []string {"Index"}

    return &(raw.EntityData)
}

