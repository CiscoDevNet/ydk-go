// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_entity_state_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_entity_state_mib"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-entity-state-mib ENTITY-STATE-MIB}", reflect.TypeOf(ENTITYSTATEMIB{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-entity-state-mib:ENTITY-STATE-MIB", reflect.TypeOf(ENTITYSTATEMIB{}))
}

// ENTITYSTATEMIB
type ENTITYSTATEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    EntStateTable ENTITYSTATEMIB_EntStateTable
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetEntityData() *types.CommonEntityData {
    eNTITYSTATEMIB.EntityData.YFilter = eNTITYSTATEMIB.YFilter
    eNTITYSTATEMIB.EntityData.YangName = "ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.BundleName = "cisco_ios_xr"
    eNTITYSTATEMIB.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-entity-state-mib"
    eNTITYSTATEMIB.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-entity-state-mib:ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.AbsolutePath = eNTITYSTATEMIB.EntityData.SegmentPath
    eNTITYSTATEMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eNTITYSTATEMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eNTITYSTATEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eNTITYSTATEMIB.EntityData.Children = types.NewOrderedMap()
    eNTITYSTATEMIB.EntityData.Children.Append("entStateTable", types.YChild{"EntStateTable", &eNTITYSTATEMIB.EntStateTable})
    eNTITYSTATEMIB.EntityData.Leafs = types.NewOrderedMap()

    eNTITYSTATEMIB.EntityData.YListKeys = []string {}

    return &(eNTITYSTATEMIB.EntityData)
}

// ENTITYSTATEMIB_EntStateTable
type ENTITYSTATEMIB_EntStateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYSTATEMIB_EntStateTable_EntStateEntry.
    EntStateEntry []*ENTITYSTATEMIB_EntStateTable_EntStateEntry
}

func (entStateTable *ENTITYSTATEMIB_EntStateTable) GetEntityData() *types.CommonEntityData {
    entStateTable.EntityData.YFilter = entStateTable.YFilter
    entStateTable.EntityData.YangName = "entStateTable"
    entStateTable.EntityData.BundleName = "cisco_ios_xr"
    entStateTable.EntityData.ParentYangName = "ENTITY-STATE-MIB"
    entStateTable.EntityData.SegmentPath = "entStateTable"
    entStateTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-state-mib:ENTITY-STATE-MIB/" + entStateTable.EntityData.SegmentPath
    entStateTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entStateTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entStateTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entStateTable.EntityData.Children = types.NewOrderedMap()
    entStateTable.EntityData.Children.Append("entStateEntry", types.YChild{"EntStateEntry", nil})
    for i := range entStateTable.EntStateEntry {
        entStateTable.EntityData.Children.Append(types.GetSegmentPath(entStateTable.EntStateEntry[i]), types.YChild{"EntStateEntry", entStateTable.EntStateEntry[i]})
    }
    entStateTable.EntityData.Leafs = types.NewOrderedMap()

    entStateTable.EntityData.YListKeys = []string {}

    return &(entStateTable.EntityData)
}

// ENTITYSTATEMIB_EntStateTable_EntStateEntry
type ENTITYSTATEMIB_EntStateTable_EntStateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    EntStateLastChanged interface{}

    // The type is EntityAdminState.
    EntStateAdmin interface{}

    // The type is EntityOperState.
    EntStateOper interface{}

    // The type is EntityUsageState.
    EntStateUsage interface{}

    // The type is map[string]bool.
    EntStateAlarm interface{}

    // The type is EntityStandbyStatus.
    EntStateStandby interface{}
}

func (entStateEntry *ENTITYSTATEMIB_EntStateTable_EntStateEntry) GetEntityData() *types.CommonEntityData {
    entStateEntry.EntityData.YFilter = entStateEntry.YFilter
    entStateEntry.EntityData.YangName = "entStateEntry"
    entStateEntry.EntityData.BundleName = "cisco_ios_xr"
    entStateEntry.EntityData.ParentYangName = "entStateTable"
    entStateEntry.EntityData.SegmentPath = "entStateEntry" + types.AddKeyToken(entStateEntry.EntPhysicalIndex, "entPhysicalIndex")
    entStateEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-state-mib:ENTITY-STATE-MIB/entStateTable/" + entStateEntry.EntityData.SegmentPath
    entStateEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entStateEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entStateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entStateEntry.EntityData.Children = types.NewOrderedMap()
    entStateEntry.EntityData.Leafs = types.NewOrderedMap()
    entStateEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entStateEntry.EntPhysicalIndex})
    entStateEntry.EntityData.Leafs.Append("entStateLastChanged", types.YLeaf{"EntStateLastChanged", entStateEntry.EntStateLastChanged})
    entStateEntry.EntityData.Leafs.Append("entStateAdmin", types.YLeaf{"EntStateAdmin", entStateEntry.EntStateAdmin})
    entStateEntry.EntityData.Leafs.Append("entStateOper", types.YLeaf{"EntStateOper", entStateEntry.EntStateOper})
    entStateEntry.EntityData.Leafs.Append("entStateUsage", types.YLeaf{"EntStateUsage", entStateEntry.EntStateUsage})
    entStateEntry.EntityData.Leafs.Append("entStateAlarm", types.YLeaf{"EntStateAlarm", entStateEntry.EntStateAlarm})
    entStateEntry.EntityData.Leafs.Append("entStateStandby", types.YLeaf{"EntStateStandby", entStateEntry.EntStateStandby})

    entStateEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(entStateEntry.EntityData)
}

