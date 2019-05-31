// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_entity_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_entity_mib"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-entity-mib ENTITY-MIB}", reflect.TypeOf(ENTITYMIB{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB", reflect.TypeOf(ENTITYMIB{}))
}

// PhysicalClass
type PhysicalClass string

const (
    PhysicalClass_other PhysicalClass = "other"

    PhysicalClass_unknown PhysicalClass = "unknown"

    PhysicalClass_chassis PhysicalClass = "chassis"

    PhysicalClass_backplane PhysicalClass = "backplane"

    PhysicalClass_container PhysicalClass = "container"

    PhysicalClass_powerSupply PhysicalClass = "powerSupply"

    PhysicalClass_fan PhysicalClass = "fan"

    PhysicalClass_sensor PhysicalClass = "sensor"

    PhysicalClass_module PhysicalClass = "module"

    PhysicalClass_port PhysicalClass = "port"

    PhysicalClass_stack PhysicalClass = "stack"
)

// ENTITYMIB
type ENTITYMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    EntityGeneral ENTITYMIB_EntityGeneral

    
    EntPhysicalTable ENTITYMIB_EntPhysicalTable

    
    EntLogicalTable ENTITYMIB_EntLogicalTable

    
    EntLPMappingTable ENTITYMIB_EntLPMappingTable

    
    EntAliasMappingTable ENTITYMIB_EntAliasMappingTable

    
    EntPhysicalContainsTable ENTITYMIB_EntPhysicalContainsTable
}

func (eNTITYMIB *ENTITYMIB) GetEntityData() *types.CommonEntityData {
    eNTITYMIB.EntityData.YFilter = eNTITYMIB.YFilter
    eNTITYMIB.EntityData.YangName = "ENTITY-MIB"
    eNTITYMIB.EntityData.BundleName = "cisco_ios_xr"
    eNTITYMIB.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-entity-mib"
    eNTITYMIB.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB"
    eNTITYMIB.EntityData.AbsolutePath = eNTITYMIB.EntityData.SegmentPath
    eNTITYMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eNTITYMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eNTITYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eNTITYMIB.EntityData.Children = types.NewOrderedMap()
    eNTITYMIB.EntityData.Children.Append("entityGeneral", types.YChild{"EntityGeneral", &eNTITYMIB.EntityGeneral})
    eNTITYMIB.EntityData.Children.Append("entPhysicalTable", types.YChild{"EntPhysicalTable", &eNTITYMIB.EntPhysicalTable})
    eNTITYMIB.EntityData.Children.Append("entLogicalTable", types.YChild{"EntLogicalTable", &eNTITYMIB.EntLogicalTable})
    eNTITYMIB.EntityData.Children.Append("entLPMappingTable", types.YChild{"EntLPMappingTable", &eNTITYMIB.EntLPMappingTable})
    eNTITYMIB.EntityData.Children.Append("entAliasMappingTable", types.YChild{"EntAliasMappingTable", &eNTITYMIB.EntAliasMappingTable})
    eNTITYMIB.EntityData.Children.Append("entPhysicalContainsTable", types.YChild{"EntPhysicalContainsTable", &eNTITYMIB.EntPhysicalContainsTable})
    eNTITYMIB.EntityData.Leafs = types.NewOrderedMap()

    eNTITYMIB.EntityData.YListKeys = []string {}

    return &(eNTITYMIB.EntityData)
}

// ENTITYMIB_EntityGeneral
type ENTITYMIB_EntityGeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    EntLastChangeTime interface{}
}

func (entityGeneral *ENTITYMIB_EntityGeneral) GetEntityData() *types.CommonEntityData {
    entityGeneral.EntityData.YFilter = entityGeneral.YFilter
    entityGeneral.EntityData.YangName = "entityGeneral"
    entityGeneral.EntityData.BundleName = "cisco_ios_xr"
    entityGeneral.EntityData.ParentYangName = "ENTITY-MIB"
    entityGeneral.EntityData.SegmentPath = "entityGeneral"
    entityGeneral.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/" + entityGeneral.EntityData.SegmentPath
    entityGeneral.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityGeneral.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityGeneral.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityGeneral.EntityData.Children = types.NewOrderedMap()
    entityGeneral.EntityData.Leafs = types.NewOrderedMap()
    entityGeneral.EntityData.Leafs.Append("entLastChangeTime", types.YLeaf{"EntLastChangeTime", entityGeneral.EntLastChangeTime})

    entityGeneral.EntityData.YListKeys = []string {}

    return &(entityGeneral.EntityData)
}

// ENTITYMIB_EntPhysicalTable
type ENTITYMIB_EntPhysicalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_EntPhysicalTable_EntPhysicalEntry.
    EntPhysicalEntry []*ENTITYMIB_EntPhysicalTable_EntPhysicalEntry
}

func (entPhysicalTable *ENTITYMIB_EntPhysicalTable) GetEntityData() *types.CommonEntityData {
    entPhysicalTable.EntityData.YFilter = entPhysicalTable.YFilter
    entPhysicalTable.EntityData.YangName = "entPhysicalTable"
    entPhysicalTable.EntityData.BundleName = "cisco_ios_xr"
    entPhysicalTable.EntityData.ParentYangName = "ENTITY-MIB"
    entPhysicalTable.EntityData.SegmentPath = "entPhysicalTable"
    entPhysicalTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/" + entPhysicalTable.EntityData.SegmentPath
    entPhysicalTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhysicalTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhysicalTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhysicalTable.EntityData.Children = types.NewOrderedMap()
    entPhysicalTable.EntityData.Children.Append("entPhysicalEntry", types.YChild{"EntPhysicalEntry", nil})
    for i := range entPhysicalTable.EntPhysicalEntry {
        entPhysicalTable.EntityData.Children.Append(types.GetSegmentPath(entPhysicalTable.EntPhysicalEntry[i]), types.YChild{"EntPhysicalEntry", entPhysicalTable.EntPhysicalEntry[i]})
    }
    entPhysicalTable.EntityData.Leafs = types.NewOrderedMap()

    entPhysicalTable.EntityData.YListKeys = []string {}

    return &(entPhysicalTable.EntityData)
}

// ENTITYMIB_EntPhysicalTable_EntPhysicalEntry
type ENTITYMIB_EntPhysicalTable_EntPhysicalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is string with length: 0..255.
    EntPhysicalDescr interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    EntPhysicalVendorType interface{}

    // The type is interface{} with range: 0..2147483647.
    EntPhysicalContainedIn interface{}

    // The type is PhysicalClass.
    EntPhysicalClass interface{}

    // The type is interface{} with range: -1..2147483647.
    EntPhysicalParentRelPos interface{}

    // The type is string with length: 0..255.
    EntPhysicalName interface{}

    // The type is string with length: 0..255.
    EntPhysicalHardwareRev interface{}

    // The type is string with length: 0..255.
    EntPhysicalFirmwareRev interface{}

    // The type is string with length: 0..255.
    EntPhysicalSoftwareRev interface{}

    // The type is string with length: 0..32.
    EntPhysicalSerialNum interface{}

    // The type is string with length: 0..255.
    EntPhysicalMfgName interface{}

    // The type is string with length: 0..255.
    EntPhysicalModelName interface{}

    // The type is string with length: 0..32.
    EntPhysicalAlias interface{}

    // The type is string with length: 0..32.
    EntPhysicalAssetID interface{}

    // The type is TruthValue.
    EntPhysicalIsFRU interface{}
}

func (entPhysicalEntry *ENTITYMIB_EntPhysicalTable_EntPhysicalEntry) GetEntityData() *types.CommonEntityData {
    entPhysicalEntry.EntityData.YFilter = entPhysicalEntry.YFilter
    entPhysicalEntry.EntityData.YangName = "entPhysicalEntry"
    entPhysicalEntry.EntityData.BundleName = "cisco_ios_xr"
    entPhysicalEntry.EntityData.ParentYangName = "entPhysicalTable"
    entPhysicalEntry.EntityData.SegmentPath = "entPhysicalEntry" + types.AddKeyToken(entPhysicalEntry.EntPhysicalIndex, "entPhysicalIndex")
    entPhysicalEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/entPhysicalTable/" + entPhysicalEntry.EntityData.SegmentPath
    entPhysicalEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhysicalEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhysicalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhysicalEntry.EntityData.Children = types.NewOrderedMap()
    entPhysicalEntry.EntityData.Leafs = types.NewOrderedMap()
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entPhysicalEntry.EntPhysicalIndex})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalDescr", types.YLeaf{"EntPhysicalDescr", entPhysicalEntry.EntPhysicalDescr})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalVendorType", types.YLeaf{"EntPhysicalVendorType", entPhysicalEntry.EntPhysicalVendorType})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalContainedIn", types.YLeaf{"EntPhysicalContainedIn", entPhysicalEntry.EntPhysicalContainedIn})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalClass", types.YLeaf{"EntPhysicalClass", entPhysicalEntry.EntPhysicalClass})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalParentRelPos", types.YLeaf{"EntPhysicalParentRelPos", entPhysicalEntry.EntPhysicalParentRelPos})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalName", types.YLeaf{"EntPhysicalName", entPhysicalEntry.EntPhysicalName})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalHardwareRev", types.YLeaf{"EntPhysicalHardwareRev", entPhysicalEntry.EntPhysicalHardwareRev})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalFirmwareRev", types.YLeaf{"EntPhysicalFirmwareRev", entPhysicalEntry.EntPhysicalFirmwareRev})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalSoftwareRev", types.YLeaf{"EntPhysicalSoftwareRev", entPhysicalEntry.EntPhysicalSoftwareRev})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalSerialNum", types.YLeaf{"EntPhysicalSerialNum", entPhysicalEntry.EntPhysicalSerialNum})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalMfgName", types.YLeaf{"EntPhysicalMfgName", entPhysicalEntry.EntPhysicalMfgName})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalModelName", types.YLeaf{"EntPhysicalModelName", entPhysicalEntry.EntPhysicalModelName})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalAlias", types.YLeaf{"EntPhysicalAlias", entPhysicalEntry.EntPhysicalAlias})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalAssetID", types.YLeaf{"EntPhysicalAssetID", entPhysicalEntry.EntPhysicalAssetID})
    entPhysicalEntry.EntityData.Leafs.Append("entPhysicalIsFRU", types.YLeaf{"EntPhysicalIsFRU", entPhysicalEntry.EntPhysicalIsFRU})

    entPhysicalEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(entPhysicalEntry.EntityData)
}

// ENTITYMIB_EntLogicalTable
type ENTITYMIB_EntLogicalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_EntLogicalTable_EntLogicalEntry.
    EntLogicalEntry []*ENTITYMIB_EntLogicalTable_EntLogicalEntry
}

func (entLogicalTable *ENTITYMIB_EntLogicalTable) GetEntityData() *types.CommonEntityData {
    entLogicalTable.EntityData.YFilter = entLogicalTable.YFilter
    entLogicalTable.EntityData.YangName = "entLogicalTable"
    entLogicalTable.EntityData.BundleName = "cisco_ios_xr"
    entLogicalTable.EntityData.ParentYangName = "ENTITY-MIB"
    entLogicalTable.EntityData.SegmentPath = "entLogicalTable"
    entLogicalTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/" + entLogicalTable.EntityData.SegmentPath
    entLogicalTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entLogicalTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entLogicalTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entLogicalTable.EntityData.Children = types.NewOrderedMap()
    entLogicalTable.EntityData.Children.Append("entLogicalEntry", types.YChild{"EntLogicalEntry", nil})
    for i := range entLogicalTable.EntLogicalEntry {
        entLogicalTable.EntityData.Children.Append(types.GetSegmentPath(entLogicalTable.EntLogicalEntry[i]), types.YChild{"EntLogicalEntry", entLogicalTable.EntLogicalEntry[i]})
    }
    entLogicalTable.EntityData.Leafs = types.NewOrderedMap()

    entLogicalTable.EntityData.YListKeys = []string {}

    return &(entLogicalTable.EntityData)
}

// ENTITYMIB_EntLogicalTable_EntLogicalEntry
type ENTITYMIB_EntLogicalTable_EntLogicalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntLogicalIndex interface{}

    // The type is string with length: 0..255.
    EntLogicalDescr interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    EntLogicalType interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    EntLogicalCommunity interface{}

    // The type is string with pattern: b'(\\d*(.\\d*)*)?'.
    EntLogicalTAddress interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    EntLogicalTDomain interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    EntLogicalContextEngineID interface{}

    // The type is string with length: 0..255.
    EntLogicalContextName interface{}
}

func (entLogicalEntry *ENTITYMIB_EntLogicalTable_EntLogicalEntry) GetEntityData() *types.CommonEntityData {
    entLogicalEntry.EntityData.YFilter = entLogicalEntry.YFilter
    entLogicalEntry.EntityData.YangName = "entLogicalEntry"
    entLogicalEntry.EntityData.BundleName = "cisco_ios_xr"
    entLogicalEntry.EntityData.ParentYangName = "entLogicalTable"
    entLogicalEntry.EntityData.SegmentPath = "entLogicalEntry" + types.AddKeyToken(entLogicalEntry.EntLogicalIndex, "entLogicalIndex")
    entLogicalEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/entLogicalTable/" + entLogicalEntry.EntityData.SegmentPath
    entLogicalEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entLogicalEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entLogicalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entLogicalEntry.EntityData.Children = types.NewOrderedMap()
    entLogicalEntry.EntityData.Leafs = types.NewOrderedMap()
    entLogicalEntry.EntityData.Leafs.Append("entLogicalIndex", types.YLeaf{"EntLogicalIndex", entLogicalEntry.EntLogicalIndex})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalDescr", types.YLeaf{"EntLogicalDescr", entLogicalEntry.EntLogicalDescr})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalType", types.YLeaf{"EntLogicalType", entLogicalEntry.EntLogicalType})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalCommunity", types.YLeaf{"EntLogicalCommunity", entLogicalEntry.EntLogicalCommunity})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalTAddress", types.YLeaf{"EntLogicalTAddress", entLogicalEntry.EntLogicalTAddress})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalTDomain", types.YLeaf{"EntLogicalTDomain", entLogicalEntry.EntLogicalTDomain})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalContextEngineID", types.YLeaf{"EntLogicalContextEngineID", entLogicalEntry.EntLogicalContextEngineID})
    entLogicalEntry.EntityData.Leafs.Append("entLogicalContextName", types.YLeaf{"EntLogicalContextName", entLogicalEntry.EntLogicalContextName})

    entLogicalEntry.EntityData.YListKeys = []string {"EntLogicalIndex"}

    return &(entLogicalEntry.EntityData)
}

// ENTITYMIB_EntLPMappingTable
type ENTITYMIB_EntLPMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_EntLPMappingTable_EntLPMappingEntry.
    EntLPMappingEntry []*ENTITYMIB_EntLPMappingTable_EntLPMappingEntry
}

func (entLPMappingTable *ENTITYMIB_EntLPMappingTable) GetEntityData() *types.CommonEntityData {
    entLPMappingTable.EntityData.YFilter = entLPMappingTable.YFilter
    entLPMappingTable.EntityData.YangName = "entLPMappingTable"
    entLPMappingTable.EntityData.BundleName = "cisco_ios_xr"
    entLPMappingTable.EntityData.ParentYangName = "ENTITY-MIB"
    entLPMappingTable.EntityData.SegmentPath = "entLPMappingTable"
    entLPMappingTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/" + entLPMappingTable.EntityData.SegmentPath
    entLPMappingTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entLPMappingTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entLPMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entLPMappingTable.EntityData.Children = types.NewOrderedMap()
    entLPMappingTable.EntityData.Children.Append("entLPMappingEntry", types.YChild{"EntLPMappingEntry", nil})
    for i := range entLPMappingTable.EntLPMappingEntry {
        entLPMappingTable.EntityData.Children.Append(types.GetSegmentPath(entLPMappingTable.EntLPMappingEntry[i]), types.YChild{"EntLPMappingEntry", entLPMappingTable.EntLPMappingEntry[i]})
    }
    entLPMappingTable.EntityData.Leafs = types.NewOrderedMap()

    entLPMappingTable.EntityData.YListKeys = []string {}

    return &(entLPMappingTable.EntityData)
}

// ENTITYMIB_EntLPMappingTable_EntLPMappingEntry
type ENTITYMIB_EntLPMappingTable_EntLPMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntLogicalIndex interface{}

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntLPPhysicalIndex interface{}
}

func (entLPMappingEntry *ENTITYMIB_EntLPMappingTable_EntLPMappingEntry) GetEntityData() *types.CommonEntityData {
    entLPMappingEntry.EntityData.YFilter = entLPMappingEntry.YFilter
    entLPMappingEntry.EntityData.YangName = "entLPMappingEntry"
    entLPMappingEntry.EntityData.BundleName = "cisco_ios_xr"
    entLPMappingEntry.EntityData.ParentYangName = "entLPMappingTable"
    entLPMappingEntry.EntityData.SegmentPath = "entLPMappingEntry" + types.AddKeyToken(entLPMappingEntry.EntLogicalIndex, "entLogicalIndex") + types.AddKeyToken(entLPMappingEntry.EntLPPhysicalIndex, "entLPPhysicalIndex")
    entLPMappingEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/entLPMappingTable/" + entLPMappingEntry.EntityData.SegmentPath
    entLPMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entLPMappingEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entLPMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entLPMappingEntry.EntityData.Children = types.NewOrderedMap()
    entLPMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    entLPMappingEntry.EntityData.Leafs.Append("entLogicalIndex", types.YLeaf{"EntLogicalIndex", entLPMappingEntry.EntLogicalIndex})
    entLPMappingEntry.EntityData.Leafs.Append("entLPPhysicalIndex", types.YLeaf{"EntLPPhysicalIndex", entLPMappingEntry.EntLPPhysicalIndex})

    entLPMappingEntry.EntityData.YListKeys = []string {"EntLogicalIndex", "EntLPPhysicalIndex"}

    return &(entLPMappingEntry.EntityData)
}

// ENTITYMIB_EntAliasMappingTable
type ENTITYMIB_EntAliasMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_EntAliasMappingTable_EntAliasMappingEntry.
    EntAliasMappingEntry []*ENTITYMIB_EntAliasMappingTable_EntAliasMappingEntry
}

func (entAliasMappingTable *ENTITYMIB_EntAliasMappingTable) GetEntityData() *types.CommonEntityData {
    entAliasMappingTable.EntityData.YFilter = entAliasMappingTable.YFilter
    entAliasMappingTable.EntityData.YangName = "entAliasMappingTable"
    entAliasMappingTable.EntityData.BundleName = "cisco_ios_xr"
    entAliasMappingTable.EntityData.ParentYangName = "ENTITY-MIB"
    entAliasMappingTable.EntityData.SegmentPath = "entAliasMappingTable"
    entAliasMappingTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/" + entAliasMappingTable.EntityData.SegmentPath
    entAliasMappingTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entAliasMappingTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entAliasMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entAliasMappingTable.EntityData.Children = types.NewOrderedMap()
    entAliasMappingTable.EntityData.Children.Append("entAliasMappingEntry", types.YChild{"EntAliasMappingEntry", nil})
    for i := range entAliasMappingTable.EntAliasMappingEntry {
        entAliasMappingTable.EntityData.Children.Append(types.GetSegmentPath(entAliasMappingTable.EntAliasMappingEntry[i]), types.YChild{"EntAliasMappingEntry", entAliasMappingTable.EntAliasMappingEntry[i]})
    }
    entAliasMappingTable.EntityData.Leafs = types.NewOrderedMap()

    entAliasMappingTable.EntityData.YListKeys = []string {}

    return &(entAliasMappingTable.EntityData)
}

// ENTITYMIB_EntAliasMappingTable_EntAliasMappingEntry
type ENTITYMIB_EntAliasMappingTable_EntAliasMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is interface{} with range: 0..2147483647.
    EntAliasLogicalIndexOrZero interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    EntAliasMappingIdentifier interface{}
}

func (entAliasMappingEntry *ENTITYMIB_EntAliasMappingTable_EntAliasMappingEntry) GetEntityData() *types.CommonEntityData {
    entAliasMappingEntry.EntityData.YFilter = entAliasMappingEntry.YFilter
    entAliasMappingEntry.EntityData.YangName = "entAliasMappingEntry"
    entAliasMappingEntry.EntityData.BundleName = "cisco_ios_xr"
    entAliasMappingEntry.EntityData.ParentYangName = "entAliasMappingTable"
    entAliasMappingEntry.EntityData.SegmentPath = "entAliasMappingEntry" + types.AddKeyToken(entAliasMappingEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(entAliasMappingEntry.EntAliasLogicalIndexOrZero, "entAliasLogicalIndexOrZero")
    entAliasMappingEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/entAliasMappingTable/" + entAliasMappingEntry.EntityData.SegmentPath
    entAliasMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entAliasMappingEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entAliasMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entAliasMappingEntry.EntityData.Children = types.NewOrderedMap()
    entAliasMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    entAliasMappingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entAliasMappingEntry.EntPhysicalIndex})
    entAliasMappingEntry.EntityData.Leafs.Append("entAliasLogicalIndexOrZero", types.YLeaf{"EntAliasLogicalIndexOrZero", entAliasMappingEntry.EntAliasLogicalIndexOrZero})
    entAliasMappingEntry.EntityData.Leafs.Append("entAliasMappingIdentifier", types.YLeaf{"EntAliasMappingIdentifier", entAliasMappingEntry.EntAliasMappingIdentifier})

    entAliasMappingEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "EntAliasLogicalIndexOrZero"}

    return &(entAliasMappingEntry.EntityData)
}

// ENTITYMIB_EntPhysicalContainsTable
type ENTITYMIB_EntPhysicalContainsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // ENTITYMIB_EntPhysicalContainsTable_EntPhysicalContainsEntry.
    EntPhysicalContainsEntry []*ENTITYMIB_EntPhysicalContainsTable_EntPhysicalContainsEntry
}

func (entPhysicalContainsTable *ENTITYMIB_EntPhysicalContainsTable) GetEntityData() *types.CommonEntityData {
    entPhysicalContainsTable.EntityData.YFilter = entPhysicalContainsTable.YFilter
    entPhysicalContainsTable.EntityData.YangName = "entPhysicalContainsTable"
    entPhysicalContainsTable.EntityData.BundleName = "cisco_ios_xr"
    entPhysicalContainsTable.EntityData.ParentYangName = "ENTITY-MIB"
    entPhysicalContainsTable.EntityData.SegmentPath = "entPhysicalContainsTable"
    entPhysicalContainsTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/" + entPhysicalContainsTable.EntityData.SegmentPath
    entPhysicalContainsTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhysicalContainsTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhysicalContainsTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhysicalContainsTable.EntityData.Children = types.NewOrderedMap()
    entPhysicalContainsTable.EntityData.Children.Append("entPhysicalContainsEntry", types.YChild{"EntPhysicalContainsEntry", nil})
    for i := range entPhysicalContainsTable.EntPhysicalContainsEntry {
        entPhysicalContainsTable.EntityData.Children.Append(types.GetSegmentPath(entPhysicalContainsTable.EntPhysicalContainsEntry[i]), types.YChild{"EntPhysicalContainsEntry", entPhysicalContainsTable.EntPhysicalContainsEntry[i]})
    }
    entPhysicalContainsTable.EntityData.Leafs = types.NewOrderedMap()

    entPhysicalContainsTable.EntityData.YListKeys = []string {}

    return &(entPhysicalContainsTable.EntityData)
}

// ENTITYMIB_EntPhysicalContainsTable_EntPhysicalContainsEntry
type ENTITYMIB_EntPhysicalContainsTable_EntPhysicalContainsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalChildIndex interface{}
}

func (entPhysicalContainsEntry *ENTITYMIB_EntPhysicalContainsTable_EntPhysicalContainsEntry) GetEntityData() *types.CommonEntityData {
    entPhysicalContainsEntry.EntityData.YFilter = entPhysicalContainsEntry.YFilter
    entPhysicalContainsEntry.EntityData.YangName = "entPhysicalContainsEntry"
    entPhysicalContainsEntry.EntityData.BundleName = "cisco_ios_xr"
    entPhysicalContainsEntry.EntityData.ParentYangName = "entPhysicalContainsTable"
    entPhysicalContainsEntry.EntityData.SegmentPath = "entPhysicalContainsEntry" + types.AddKeyToken(entPhysicalContainsEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(entPhysicalContainsEntry.EntPhysicalChildIndex, "entPhysicalChildIndex")
    entPhysicalContainsEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB/entPhysicalContainsTable/" + entPhysicalContainsEntry.EntityData.SegmentPath
    entPhysicalContainsEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhysicalContainsEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhysicalContainsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhysicalContainsEntry.EntityData.Children = types.NewOrderedMap()
    entPhysicalContainsEntry.EntityData.Leafs = types.NewOrderedMap()
    entPhysicalContainsEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entPhysicalContainsEntry.EntPhysicalIndex})
    entPhysicalContainsEntry.EntityData.Leafs.Append("entPhysicalChildIndex", types.YLeaf{"EntPhysicalChildIndex", entPhysicalContainsEntry.EntPhysicalChildIndex})

    entPhysicalContainsEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "EntPhysicalChildIndex"}

    return &(entPhysicalContainsEntry.EntityData)
}

