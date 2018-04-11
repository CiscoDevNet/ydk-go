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

    
    Entstatetable ENTITYSTATEMIB_Entstatetable
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetEntityData() *types.CommonEntityData {
    eNTITYSTATEMIB.EntityData.YFilter = eNTITYSTATEMIB.YFilter
    eNTITYSTATEMIB.EntityData.YangName = "ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.BundleName = "cisco_ios_xr"
    eNTITYSTATEMIB.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-entity-state-mib"
    eNTITYSTATEMIB.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-entity-state-mib:ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eNTITYSTATEMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eNTITYSTATEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eNTITYSTATEMIB.EntityData.Children = make(map[string]types.YChild)
    eNTITYSTATEMIB.EntityData.Children["entStateTable"] = types.YChild{"Entstatetable", &eNTITYSTATEMIB.Entstatetable}
    eNTITYSTATEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eNTITYSTATEMIB.EntityData)
}

// ENTITYSTATEMIB_Entstatetable
type ENTITYSTATEMIB_Entstatetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYSTATEMIB_Entstatetable_Entstateentry.
    Entstateentry []ENTITYSTATEMIB_Entstatetable_Entstateentry
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetEntityData() *types.CommonEntityData {
    entstatetable.EntityData.YFilter = entstatetable.YFilter
    entstatetable.EntityData.YangName = "entStateTable"
    entstatetable.EntityData.BundleName = "cisco_ios_xr"
    entstatetable.EntityData.ParentYangName = "ENTITY-STATE-MIB"
    entstatetable.EntityData.SegmentPath = "entStateTable"
    entstatetable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entstatetable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entstatetable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entstatetable.EntityData.Children = make(map[string]types.YChild)
    entstatetable.EntityData.Children["entStateEntry"] = types.YChild{"Entstateentry", nil}
    for i := range entstatetable.Entstateentry {
        entstatetable.EntityData.Children[types.GetSegmentPath(&entstatetable.Entstateentry[i])] = types.YChild{"Entstateentry", &entstatetable.Entstateentry[i]}
    }
    entstatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entstatetable.EntityData)
}

// ENTITYSTATEMIB_Entstatetable_Entstateentry
type ENTITYSTATEMIB_Entstatetable_Entstateentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalindex interface{}

    // The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    Entstatelastchanged interface{}

    // The type is EntityAdminState.
    Entstateadmin interface{}

    // The type is EntityOperState.
    Entstateoper interface{}

    // The type is EntityUsageState.
    Entstateusage interface{}

    // The type is map[string]bool.
    Entstatealarm interface{}

    // The type is EntityStandbyStatus.
    Entstatestandby interface{}
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetEntityData() *types.CommonEntityData {
    entstateentry.EntityData.YFilter = entstateentry.YFilter
    entstateentry.EntityData.YangName = "entStateEntry"
    entstateentry.EntityData.BundleName = "cisco_ios_xr"
    entstateentry.EntityData.ParentYangName = "entStateTable"
    entstateentry.EntityData.SegmentPath = "entStateEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entstateentry.Entphysicalindex) + "']"
    entstateentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entstateentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entstateentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entstateentry.EntityData.Children = make(map[string]types.YChild)
    entstateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entstateentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entstateentry.Entphysicalindex}
    entstateentry.EntityData.Leafs["entStateLastChanged"] = types.YLeaf{"Entstatelastchanged", entstateentry.Entstatelastchanged}
    entstateentry.EntityData.Leafs["entStateAdmin"] = types.YLeaf{"Entstateadmin", entstateentry.Entstateadmin}
    entstateentry.EntityData.Leafs["entStateOper"] = types.YLeaf{"Entstateoper", entstateentry.Entstateoper}
    entstateentry.EntityData.Leafs["entStateUsage"] = types.YLeaf{"Entstateusage", entstateentry.Entstateusage}
    entstateentry.EntityData.Leafs["entStateAlarm"] = types.YLeaf{"Entstatealarm", entstateentry.Entstatealarm}
    entstateentry.EntityData.Leafs["entStateStandby"] = types.YLeaf{"Entstatestandby", entstateentry.Entstatestandby}
    return &(entstateentry.EntityData)
}

