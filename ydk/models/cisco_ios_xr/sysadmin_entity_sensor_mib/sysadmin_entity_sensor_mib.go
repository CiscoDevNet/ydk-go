// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_entity_sensor_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_entity_sensor_mib"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-entity-sensor-mib CISCO-ENTITY-SENSOR-MIB}", reflect.TypeOf(CISCOENTITYSENSORMIB{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB", reflect.TypeOf(CISCOENTITYSENSORMIB{}))
}

// SensorDataType
type SensorDataType string

const (
    SensorDataType_other SensorDataType = "other"

    SensorDataType_unknown SensorDataType = "unknown"

    SensorDataType_voltsAC SensorDataType = "voltsAC"

    SensorDataType_voltsDC SensorDataType = "voltsDC"

    SensorDataType_amperes SensorDataType = "amperes"

    SensorDataType_watts SensorDataType = "watts"

    SensorDataType_hertz SensorDataType = "hertz"

    SensorDataType_celsius SensorDataType = "celsius"

    SensorDataType_percentRH SensorDataType = "percentRH"

    SensorDataType_rpm SensorDataType = "rpm"

    SensorDataType_cmm SensorDataType = "cmm"

    SensorDataType_truthvalue SensorDataType = "truthvalue"

    SensorDataType_specialEnum SensorDataType = "specialEnum"
)

// SensorDataScale
type SensorDataScale string

const (
    SensorDataScale_yocto SensorDataScale = "yocto"

    SensorDataScale_zepto SensorDataScale = "zepto"

    SensorDataScale_atto SensorDataScale = "atto"

    SensorDataScale_femto SensorDataScale = "femto"

    SensorDataScale_pico SensorDataScale = "pico"

    SensorDataScale_nano SensorDataScale = "nano"

    SensorDataScale_micro SensorDataScale = "micro"

    SensorDataScale_milli SensorDataScale = "milli"

    SensorDataScale_units SensorDataScale = "units"

    SensorDataScale_kilo SensorDataScale = "kilo"

    SensorDataScale_mega SensorDataScale = "mega"

    SensorDataScale_giga SensorDataScale = "giga"

    SensorDataScale_tera SensorDataScale = "tera"

    SensorDataScale_exa SensorDataScale = "exa"

    SensorDataScale_peta SensorDataScale = "peta"

    SensorDataScale_zetta SensorDataScale = "zetta"

    SensorDataScale_yotta SensorDataScale = "yotta"
)

// SensorStatus
type SensorStatus string

const (
    SensorStatus_ok SensorStatus = "ok"

    SensorStatus_unavailable SensorStatus = "unavailable"

    SensorStatus_nonoperational SensorStatus = "nonoperational"
)

// SensorThresholdSeverity
type SensorThresholdSeverity string

const (
    SensorThresholdSeverity_other SensorThresholdSeverity = "other"

    SensorThresholdSeverity_minor SensorThresholdSeverity = "minor"

    SensorThresholdSeverity_major SensorThresholdSeverity = "major"

    SensorThresholdSeverity_critical SensorThresholdSeverity = "critical"
)

// SensorThresholdRelation
type SensorThresholdRelation string

const (
    SensorThresholdRelation_lessThan SensorThresholdRelation = "lessThan"

    SensorThresholdRelation_lessOrEqual SensorThresholdRelation = "lessOrEqual"

    SensorThresholdRelation_greaterThan SensorThresholdRelation = "greaterThan"

    SensorThresholdRelation_greaterOrEqual SensorThresholdRelation = "greaterOrEqual"

    SensorThresholdRelation_equalTo SensorThresholdRelation = "equalTo"

    SensorThresholdRelation_notEqualTo SensorThresholdRelation = "notEqualTo"
)

// CISCOENTITYSENSORMIB
type CISCOENTITYSENSORMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    EntSensorValueTable CISCOENTITYSENSORMIB_EntSensorValueTable

    
    EntSensorThresholdTable CISCOENTITYSENSORMIB_EntSensorThresholdTable
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYSENSORMIB.EntityData.YFilter = cISCOENTITYSENSORMIB.YFilter
    cISCOENTITYSENSORMIB.EntityData.YangName = "CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.BundleName = "cisco_ios_xr"
    cISCOENTITYSENSORMIB.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-entity-sensor-mib"
    cISCOENTITYSENSORMIB.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.AbsolutePath = cISCOENTITYSENSORMIB.EntityData.SegmentPath
    cISCOENTITYSENSORMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cISCOENTITYSENSORMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cISCOENTITYSENSORMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cISCOENTITYSENSORMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENTITYSENSORMIB.EntityData.Children.Append("entSensorValueTable", types.YChild{"EntSensorValueTable", &cISCOENTITYSENSORMIB.EntSensorValueTable})
    cISCOENTITYSENSORMIB.EntityData.Children.Append("entSensorThresholdTable", types.YChild{"EntSensorThresholdTable", &cISCOENTITYSENSORMIB.EntSensorThresholdTable})
    cISCOENTITYSENSORMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENTITYSENSORMIB.EntityData.YListKeys = []string {}

    return &(cISCOENTITYSENSORMIB.EntityData)
}

// CISCOENTITYSENSORMIB_EntSensorValueTable
type CISCOENTITYSENSORMIB_EntSensorValueTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry.
    EntSensorValueEntry []*CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry
}

func (entSensorValueTable *CISCOENTITYSENSORMIB_EntSensorValueTable) GetEntityData() *types.CommonEntityData {
    entSensorValueTable.EntityData.YFilter = entSensorValueTable.YFilter
    entSensorValueTable.EntityData.YangName = "entSensorValueTable"
    entSensorValueTable.EntityData.BundleName = "cisco_ios_xr"
    entSensorValueTable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entSensorValueTable.EntityData.SegmentPath = "entSensorValueTable"
    entSensorValueTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB/" + entSensorValueTable.EntityData.SegmentPath
    entSensorValueTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entSensorValueTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entSensorValueTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entSensorValueTable.EntityData.Children = types.NewOrderedMap()
    entSensorValueTable.EntityData.Children.Append("entSensorValueEntry", types.YChild{"EntSensorValueEntry", nil})
    for i := range entSensorValueTable.EntSensorValueEntry {
        entSensorValueTable.EntityData.Children.Append(types.GetSegmentPath(entSensorValueTable.EntSensorValueEntry[i]), types.YChild{"EntSensorValueEntry", entSensorValueTable.EntSensorValueEntry[i]})
    }
    entSensorValueTable.EntityData.Leafs = types.NewOrderedMap()

    entSensorValueTable.EntityData.YListKeys = []string {}

    return &(entSensorValueTable.EntityData)
}

// CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry
type CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // The type is SensorDataType.
    EntSensorType interface{}

    // The type is SensorDataScale.
    EntSensorScale interface{}

    // The type is interface{} with range: -8..9.
    EntSensorPrecision interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    EntSensorValue interface{}

    // The type is SensorStatus.
    EntSensorStatus interface{}

    // The type is interface{} with range: 0..4294967295.
    EntSensorValueTimeStamp interface{}

    // The type is interface{} with range: 0..999999999.
    EntSensorValueUpdateRate interface{}

    // The type is interface{} with range: 0..2147483647.
    EntSensorMeasuredEntity interface{}
}

func (entSensorValueEntry *CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry) GetEntityData() *types.CommonEntityData {
    entSensorValueEntry.EntityData.YFilter = entSensorValueEntry.YFilter
    entSensorValueEntry.EntityData.YangName = "entSensorValueEntry"
    entSensorValueEntry.EntityData.BundleName = "cisco_ios_xr"
    entSensorValueEntry.EntityData.ParentYangName = "entSensorValueTable"
    entSensorValueEntry.EntityData.SegmentPath = "entSensorValueEntry" + types.AddKeyToken(entSensorValueEntry.EntPhysicalIndex, "entPhysicalIndex")
    entSensorValueEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB/entSensorValueTable/" + entSensorValueEntry.EntityData.SegmentPath
    entSensorValueEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entSensorValueEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entSensorValueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entSensorValueEntry.EntityData.Children = types.NewOrderedMap()
    entSensorValueEntry.EntityData.Leafs = types.NewOrderedMap()
    entSensorValueEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entSensorValueEntry.EntPhysicalIndex})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorType", types.YLeaf{"EntSensorType", entSensorValueEntry.EntSensorType})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorScale", types.YLeaf{"EntSensorScale", entSensorValueEntry.EntSensorScale})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorPrecision", types.YLeaf{"EntSensorPrecision", entSensorValueEntry.EntSensorPrecision})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorValue", types.YLeaf{"EntSensorValue", entSensorValueEntry.EntSensorValue})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorStatus", types.YLeaf{"EntSensorStatus", entSensorValueEntry.EntSensorStatus})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorValueTimeStamp", types.YLeaf{"EntSensorValueTimeStamp", entSensorValueEntry.EntSensorValueTimeStamp})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorValueUpdateRate", types.YLeaf{"EntSensorValueUpdateRate", entSensorValueEntry.EntSensorValueUpdateRate})
    entSensorValueEntry.EntityData.Leafs.Append("entSensorMeasuredEntity", types.YLeaf{"EntSensorMeasuredEntity", entSensorValueEntry.EntSensorMeasuredEntity})

    entSensorValueEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(entSensorValueEntry.EntityData)
}

// CISCOENTITYSENSORMIB_EntSensorThresholdTable
type CISCOENTITYSENSORMIB_EntSensorThresholdTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry.
    EntSensorThresholdEntry []*CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry
}

func (entSensorThresholdTable *CISCOENTITYSENSORMIB_EntSensorThresholdTable) GetEntityData() *types.CommonEntityData {
    entSensorThresholdTable.EntityData.YFilter = entSensorThresholdTable.YFilter
    entSensorThresholdTable.EntityData.YangName = "entSensorThresholdTable"
    entSensorThresholdTable.EntityData.BundleName = "cisco_ios_xr"
    entSensorThresholdTable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entSensorThresholdTable.EntityData.SegmentPath = "entSensorThresholdTable"
    entSensorThresholdTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB/" + entSensorThresholdTable.EntityData.SegmentPath
    entSensorThresholdTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entSensorThresholdTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entSensorThresholdTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entSensorThresholdTable.EntityData.Children = types.NewOrderedMap()
    entSensorThresholdTable.EntityData.Children.Append("entSensorThresholdEntry", types.YChild{"EntSensorThresholdEntry", nil})
    for i := range entSensorThresholdTable.EntSensorThresholdEntry {
        entSensorThresholdTable.EntityData.Children.Append(types.GetSegmentPath(entSensorThresholdTable.EntSensorThresholdEntry[i]), types.YChild{"EntSensorThresholdEntry", entSensorThresholdTable.EntSensorThresholdEntry[i]})
    }
    entSensorThresholdTable.EntityData.Leafs = types.NewOrderedMap()

    entSensorThresholdTable.EntityData.YListKeys = []string {}

    return &(entSensorThresholdTable.EntityData)
}

// CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry
type CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is interface{} with range: 1..99999999.
    EntSensorThresholdIndex interface{}

    // The type is SensorThresholdSeverity.
    EntSensorThresholdSeverity interface{}

    // The type is SensorThresholdRelation.
    EntSensorThresholdRelation interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    EntSensorThresholdValue interface{}

    // The type is TruthValue.
    EntSensorThresholdEvaluation interface{}

    // The type is TruthValue.
    EntSensorThresholdNotificationEnable interface{}
}

func (entSensorThresholdEntry *CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry) GetEntityData() *types.CommonEntityData {
    entSensorThresholdEntry.EntityData.YFilter = entSensorThresholdEntry.YFilter
    entSensorThresholdEntry.EntityData.YangName = "entSensorThresholdEntry"
    entSensorThresholdEntry.EntityData.BundleName = "cisco_ios_xr"
    entSensorThresholdEntry.EntityData.ParentYangName = "entSensorThresholdTable"
    entSensorThresholdEntry.EntityData.SegmentPath = "entSensorThresholdEntry" + types.AddKeyToken(entSensorThresholdEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(entSensorThresholdEntry.EntSensorThresholdIndex, "entSensorThresholdIndex")
    entSensorThresholdEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB/entSensorThresholdTable/" + entSensorThresholdEntry.EntityData.SegmentPath
    entSensorThresholdEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entSensorThresholdEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entSensorThresholdEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entSensorThresholdEntry.EntityData.Children = types.NewOrderedMap()
    entSensorThresholdEntry.EntityData.Leafs = types.NewOrderedMap()
    entSensorThresholdEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entSensorThresholdEntry.EntPhysicalIndex})
    entSensorThresholdEntry.EntityData.Leafs.Append("entSensorThresholdIndex", types.YLeaf{"EntSensorThresholdIndex", entSensorThresholdEntry.EntSensorThresholdIndex})
    entSensorThresholdEntry.EntityData.Leafs.Append("entSensorThresholdSeverity", types.YLeaf{"EntSensorThresholdSeverity", entSensorThresholdEntry.EntSensorThresholdSeverity})
    entSensorThresholdEntry.EntityData.Leafs.Append("entSensorThresholdRelation", types.YLeaf{"EntSensorThresholdRelation", entSensorThresholdEntry.EntSensorThresholdRelation})
    entSensorThresholdEntry.EntityData.Leafs.Append("entSensorThresholdValue", types.YLeaf{"EntSensorThresholdValue", entSensorThresholdEntry.EntSensorThresholdValue})
    entSensorThresholdEntry.EntityData.Leafs.Append("entSensorThresholdEvaluation", types.YLeaf{"EntSensorThresholdEvaluation", entSensorThresholdEntry.EntSensorThresholdEvaluation})
    entSensorThresholdEntry.EntityData.Leafs.Append("entSensorThresholdNotificationEnable", types.YLeaf{"EntSensorThresholdNotificationEnable", entSensorThresholdEntry.EntSensorThresholdNotificationEnable})

    entSensorThresholdEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "EntSensorThresholdIndex"}

    return &(entSensorThresholdEntry.EntityData)
}

