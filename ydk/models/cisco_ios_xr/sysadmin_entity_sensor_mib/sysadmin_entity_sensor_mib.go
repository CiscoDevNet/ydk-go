// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
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

    
    Entsensorvaluetable CISCOENTITYSENSORMIB_Entsensorvaluetable

    
    Entsensorthresholdtable CISCOENTITYSENSORMIB_Entsensorthresholdtable
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYSENSORMIB.EntityData.YFilter = cISCOENTITYSENSORMIB.YFilter
    cISCOENTITYSENSORMIB.EntityData.YangName = "CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.BundleName = "cisco_ios_xr"
    cISCOENTITYSENSORMIB.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-entity-sensor-mib"
    cISCOENTITYSENSORMIB.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-entity-sensor-mib:CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cISCOENTITYSENSORMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cISCOENTITYSENSORMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cISCOENTITYSENSORMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENTITYSENSORMIB.EntityData.Children["entSensorValueTable"] = types.YChild{"Entsensorvaluetable", &cISCOENTITYSENSORMIB.Entsensorvaluetable}
    cISCOENTITYSENSORMIB.EntityData.Children["entSensorThresholdTable"] = types.YChild{"Entsensorthresholdtable", &cISCOENTITYSENSORMIB.Entsensorthresholdtable}
    cISCOENTITYSENSORMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENTITYSENSORMIB.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorvaluetable
type CISCOENTITYSENSORMIB_Entsensorvaluetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry.
    Entsensorvalueentry []CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetEntityData() *types.CommonEntityData {
    entsensorvaluetable.EntityData.YFilter = entsensorvaluetable.YFilter
    entsensorvaluetable.EntityData.YangName = "entSensorValueTable"
    entsensorvaluetable.EntityData.BundleName = "cisco_ios_xr"
    entsensorvaluetable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entsensorvaluetable.EntityData.SegmentPath = "entSensorValueTable"
    entsensorvaluetable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entsensorvaluetable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entsensorvaluetable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entsensorvaluetable.EntityData.Children = make(map[string]types.YChild)
    entsensorvaluetable.EntityData.Children["entSensorValueEntry"] = types.YChild{"Entsensorvalueentry", nil}
    for i := range entsensorvaluetable.Entsensorvalueentry {
        entsensorvaluetable.EntityData.Children[types.GetSegmentPath(&entsensorvaluetable.Entsensorvalueentry[i])] = types.YChild{"Entsensorvalueentry", &entsensorvaluetable.Entsensorvalueentry[i]}
    }
    entsensorvaluetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entsensorvaluetable.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry
type CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalindex interface{}

    // The type is SensorDataType.
    Entsensortype interface{}

    // The type is SensorDataScale.
    Entsensorscale interface{}

    // The type is interface{} with range: -8..9.
    Entsensorprecision interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    Entsensorvalue interface{}

    // The type is SensorStatus.
    Entsensorstatus interface{}

    // The type is interface{} with range: 0..4294967295.
    Entsensorvaluetimestamp interface{}

    // The type is interface{} with range: 0..999999999.
    Entsensorvalueupdaterate interface{}

    // The type is interface{} with range: 0..2147483647.
    Entsensormeasuredentity interface{}
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetEntityData() *types.CommonEntityData {
    entsensorvalueentry.EntityData.YFilter = entsensorvalueentry.YFilter
    entsensorvalueentry.EntityData.YangName = "entSensorValueEntry"
    entsensorvalueentry.EntityData.BundleName = "cisco_ios_xr"
    entsensorvalueentry.EntityData.ParentYangName = "entSensorValueTable"
    entsensorvalueentry.EntityData.SegmentPath = "entSensorValueEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entsensorvalueentry.Entphysicalindex) + "']"
    entsensorvalueentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entsensorvalueentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entsensorvalueentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entsensorvalueentry.EntityData.Children = make(map[string]types.YChild)
    entsensorvalueentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entsensorvalueentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entsensorvalueentry.Entphysicalindex}
    entsensorvalueentry.EntityData.Leafs["entSensorType"] = types.YLeaf{"Entsensortype", entsensorvalueentry.Entsensortype}
    entsensorvalueentry.EntityData.Leafs["entSensorScale"] = types.YLeaf{"Entsensorscale", entsensorvalueentry.Entsensorscale}
    entsensorvalueentry.EntityData.Leafs["entSensorPrecision"] = types.YLeaf{"Entsensorprecision", entsensorvalueentry.Entsensorprecision}
    entsensorvalueentry.EntityData.Leafs["entSensorValue"] = types.YLeaf{"Entsensorvalue", entsensorvalueentry.Entsensorvalue}
    entsensorvalueentry.EntityData.Leafs["entSensorStatus"] = types.YLeaf{"Entsensorstatus", entsensorvalueentry.Entsensorstatus}
    entsensorvalueentry.EntityData.Leafs["entSensorValueTimeStamp"] = types.YLeaf{"Entsensorvaluetimestamp", entsensorvalueentry.Entsensorvaluetimestamp}
    entsensorvalueentry.EntityData.Leafs["entSensorValueUpdateRate"] = types.YLeaf{"Entsensorvalueupdaterate", entsensorvalueentry.Entsensorvalueupdaterate}
    entsensorvalueentry.EntityData.Leafs["entSensorMeasuredEntity"] = types.YLeaf{"Entsensormeasuredentity", entsensorvalueentry.Entsensormeasuredentity}
    return &(entsensorvalueentry.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorthresholdtable
type CISCOENTITYSENSORMIB_Entsensorthresholdtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry.
    Entsensorthresholdentry []CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetEntityData() *types.CommonEntityData {
    entsensorthresholdtable.EntityData.YFilter = entsensorthresholdtable.YFilter
    entsensorthresholdtable.EntityData.YangName = "entSensorThresholdTable"
    entsensorthresholdtable.EntityData.BundleName = "cisco_ios_xr"
    entsensorthresholdtable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entsensorthresholdtable.EntityData.SegmentPath = "entSensorThresholdTable"
    entsensorthresholdtable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entsensorthresholdtable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entsensorthresholdtable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entsensorthresholdtable.EntityData.Children = make(map[string]types.YChild)
    entsensorthresholdtable.EntityData.Children["entSensorThresholdEntry"] = types.YChild{"Entsensorthresholdentry", nil}
    for i := range entsensorthresholdtable.Entsensorthresholdentry {
        entsensorthresholdtable.EntityData.Children[types.GetSegmentPath(&entsensorthresholdtable.Entsensorthresholdentry[i])] = types.YChild{"Entsensorthresholdentry", &entsensorthresholdtable.Entsensorthresholdentry[i]}
    }
    entsensorthresholdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entsensorthresholdtable.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry
type CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalindex interface{}

    // This attribute is a key. The type is interface{} with range: 1..99999999.
    Entsensorthresholdindex interface{}

    // The type is SensorThresholdSeverity.
    Entsensorthresholdseverity interface{}

    // The type is SensorThresholdRelation.
    Entsensorthresholdrelation interface{}

    // The type is interface{} with range: -1000000000..1000000000.
    Entsensorthresholdvalue interface{}

    // The type is TruthValue.
    Entsensorthresholdevaluation interface{}

    // The type is TruthValue.
    Entsensorthresholdnotificationenable interface{}
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetEntityData() *types.CommonEntityData {
    entsensorthresholdentry.EntityData.YFilter = entsensorthresholdentry.YFilter
    entsensorthresholdentry.EntityData.YangName = "entSensorThresholdEntry"
    entsensorthresholdentry.EntityData.BundleName = "cisco_ios_xr"
    entsensorthresholdentry.EntityData.ParentYangName = "entSensorThresholdTable"
    entsensorthresholdentry.EntityData.SegmentPath = "entSensorThresholdEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entsensorthresholdentry.Entphysicalindex) + "']" + "[entSensorThresholdIndex='" + fmt.Sprintf("%v", entsensorthresholdentry.Entsensorthresholdindex) + "']"
    entsensorthresholdentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entsensorthresholdentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entsensorthresholdentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entsensorthresholdentry.EntityData.Children = make(map[string]types.YChild)
    entsensorthresholdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entsensorthresholdentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entsensorthresholdentry.Entphysicalindex}
    entsensorthresholdentry.EntityData.Leafs["entSensorThresholdIndex"] = types.YLeaf{"Entsensorthresholdindex", entsensorthresholdentry.Entsensorthresholdindex}
    entsensorthresholdentry.EntityData.Leafs["entSensorThresholdSeverity"] = types.YLeaf{"Entsensorthresholdseverity", entsensorthresholdentry.Entsensorthresholdseverity}
    entsensorthresholdentry.EntityData.Leafs["entSensorThresholdRelation"] = types.YLeaf{"Entsensorthresholdrelation", entsensorthresholdentry.Entsensorthresholdrelation}
    entsensorthresholdentry.EntityData.Leafs["entSensorThresholdValue"] = types.YLeaf{"Entsensorthresholdvalue", entsensorthresholdentry.Entsensorthresholdvalue}
    entsensorthresholdentry.EntityData.Leafs["entSensorThresholdEvaluation"] = types.YLeaf{"Entsensorthresholdevaluation", entsensorthresholdentry.Entsensorthresholdevaluation}
    entsensorthresholdentry.EntityData.Leafs["entSensorThresholdNotificationEnable"] = types.YLeaf{"Entsensorthresholdnotificationenable", entsensorthresholdentry.Entsensorthresholdnotificationenable}
    return &(entsensorthresholdentry.EntityData)
}

