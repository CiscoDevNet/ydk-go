// This module contains a collection of YANG definitions
// for Cisco IOS-XR invmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   inventory-configurations: Configuration for inventory entities
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package invmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package invmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-invmgr-cfg inventory-configurations}", reflect.TypeOf(InventoryConfigurations{}))
    ydk.RegisterEntity("Cisco-IOS-XR-invmgr-cfg:inventory-configurations", reflect.TypeOf(InventoryConfigurations{}))
}

// InventoryConfigurations
// Configuration for inventory entities
type InventoryConfigurations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity name. The type is slice of InventoryConfigurations_Entity.
    Entity []InventoryConfigurations_Entity
}

func (inventoryConfigurations *InventoryConfigurations) GetFilter() yfilter.YFilter { return inventoryConfigurations.YFilter }

func (inventoryConfigurations *InventoryConfigurations) SetFilter(yf yfilter.YFilter) { inventoryConfigurations.YFilter = yf }

func (inventoryConfigurations *InventoryConfigurations) GetGoName(yname string) string {
    if yname == "entity" { return "Entity" }
    return ""
}

func (inventoryConfigurations *InventoryConfigurations) GetSegmentPath() string {
    return "Cisco-IOS-XR-invmgr-cfg:inventory-configurations"
}

func (inventoryConfigurations *InventoryConfigurations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entity" {
        for _, c := range inventoryConfigurations.Entity {
            if inventoryConfigurations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InventoryConfigurations_Entity{}
        inventoryConfigurations.Entity = append(inventoryConfigurations.Entity, child)
        return &inventoryConfigurations.Entity[len(inventoryConfigurations.Entity)-1]
    }
    return nil
}

func (inventoryConfigurations *InventoryConfigurations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventoryConfigurations.Entity {
        children[inventoryConfigurations.Entity[i].GetSegmentPath()] = &inventoryConfigurations.Entity[i]
    }
    return children
}

func (inventoryConfigurations *InventoryConfigurations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventoryConfigurations *InventoryConfigurations) GetBundleName() string { return "cisco_ios_xr" }

func (inventoryConfigurations *InventoryConfigurations) GetYangName() string { return "inventory-configurations" }

func (inventoryConfigurations *InventoryConfigurations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventoryConfigurations *InventoryConfigurations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventoryConfigurations *InventoryConfigurations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventoryConfigurations *InventoryConfigurations) SetParent(parent types.Entity) { inventoryConfigurations.parent = parent }

func (inventoryConfigurations *InventoryConfigurations) GetParent() types.Entity { return inventoryConfigurations.parent }

func (inventoryConfigurations *InventoryConfigurations) GetParentYangName() string { return "Cisco-IOS-XR-invmgr-cfg" }

// InventoryConfigurations_Entity
// Entity name
type InventoryConfigurations_Entity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Entity name. The type is string.
    Name interface{}

    // Entity name. The type is string. This attribute is mandatory.
    NameXr interface{}
}

func (entity *InventoryConfigurations_Entity) GetFilter() yfilter.YFilter { return entity.YFilter }

func (entity *InventoryConfigurations_Entity) SetFilter(yf yfilter.YFilter) { entity.YFilter = yf }

func (entity *InventoryConfigurations_Entity) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "name-xr" { return "NameXr" }
    return ""
}

func (entity *InventoryConfigurations_Entity) GetSegmentPath() string {
    return "entity" + "[name='" + fmt.Sprintf("%v", entity.Name) + "']"
}

func (entity *InventoryConfigurations_Entity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entity *InventoryConfigurations_Entity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entity *InventoryConfigurations_Entity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = entity.Name
    leafs["name-xr"] = entity.NameXr
    return leafs
}

func (entity *InventoryConfigurations_Entity) GetBundleName() string { return "cisco_ios_xr" }

func (entity *InventoryConfigurations_Entity) GetYangName() string { return "entity" }

func (entity *InventoryConfigurations_Entity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entity *InventoryConfigurations_Entity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entity *InventoryConfigurations_Entity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entity *InventoryConfigurations_Entity) SetParent(parent types.Entity) { entity.parent = parent }

func (entity *InventoryConfigurations_Entity) GetParent() types.Entity { return entity.parent }

func (entity *InventoryConfigurations_Entity) GetParentYangName() string { return "inventory-configurations" }

