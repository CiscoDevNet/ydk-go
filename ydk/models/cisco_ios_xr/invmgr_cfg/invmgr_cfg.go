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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity name. The type is slice of InventoryConfigurations_Entity.
    Entity []*InventoryConfigurations_Entity
}

func (inventoryConfigurations *InventoryConfigurations) GetEntityData() *types.CommonEntityData {
    inventoryConfigurations.EntityData.YFilter = inventoryConfigurations.YFilter
    inventoryConfigurations.EntityData.YangName = "inventory-configurations"
    inventoryConfigurations.EntityData.BundleName = "cisco_ios_xr"
    inventoryConfigurations.EntityData.ParentYangName = "Cisco-IOS-XR-invmgr-cfg"
    inventoryConfigurations.EntityData.SegmentPath = "Cisco-IOS-XR-invmgr-cfg:inventory-configurations"
    inventoryConfigurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventoryConfigurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventoryConfigurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventoryConfigurations.EntityData.Children = types.NewOrderedMap()
    inventoryConfigurations.EntityData.Children.Append("entity", types.YChild{"Entity", nil})
    for i := range inventoryConfigurations.Entity {
        inventoryConfigurations.EntityData.Children.Append(types.GetSegmentPath(inventoryConfigurations.Entity[i]), types.YChild{"Entity", inventoryConfigurations.Entity[i]})
    }
    inventoryConfigurations.EntityData.Leafs = types.NewOrderedMap()

    inventoryConfigurations.EntityData.YListKeys = []string {}

    return &(inventoryConfigurations.EntityData)
}

// InventoryConfigurations_Entity
// Entity name
type InventoryConfigurations_Entity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Entity name. The type is string.
    Name interface{}

    // Entity name. The type is string. This attribute is mandatory.
    NameXr interface{}
}

func (entity *InventoryConfigurations_Entity) GetEntityData() *types.CommonEntityData {
    entity.EntityData.YFilter = entity.YFilter
    entity.EntityData.YangName = "entity"
    entity.EntityData.BundleName = "cisco_ios_xr"
    entity.EntityData.ParentYangName = "inventory-configurations"
    entity.EntityData.SegmentPath = "entity" + types.AddKeyToken(entity.Name, "name")
    entity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entity.EntityData.Children = types.NewOrderedMap()
    entity.EntityData.Leafs = types.NewOrderedMap()
    entity.EntityData.Leafs.Append("name", types.YLeaf{"Name", entity.Name})
    entity.EntityData.Leafs.Append("name-xr", types.YLeaf{"NameXr", entity.NameXr})

    entity.EntityData.YListKeys = []string {"Name"}

    return &(entity.EntityData)
}

