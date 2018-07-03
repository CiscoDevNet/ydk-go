// This module defines Entity MIB extensions for physical
// sensors.
// 
// Copyright (C) The Internet Society (2002). This version
// of this MIB module is part of RFC 3433; see the RFC
// itself for full legal notices.
package entity_sensor_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package entity_sensor_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:ENTITY-SENSOR-MIB ENTITY-SENSOR-MIB}", reflect.TypeOf(ENTITYSENSORMIB{}))
    ydk.RegisterEntity("ENTITY-SENSOR-MIB:ENTITY-SENSOR-MIB", reflect.TypeOf(ENTITYSENSORMIB{}))
}

// EntitySensorDataType represents    truthvalue(12):  value takes { true(1), false(2) }
type EntitySensorDataType string

const (
    EntitySensorDataType_other EntitySensorDataType = "other"

    EntitySensorDataType_unknown EntitySensorDataType = "unknown"

    EntitySensorDataType_voltsAC EntitySensorDataType = "voltsAC"

    EntitySensorDataType_voltsDC EntitySensorDataType = "voltsDC"

    EntitySensorDataType_amperes EntitySensorDataType = "amperes"

    EntitySensorDataType_watts EntitySensorDataType = "watts"

    EntitySensorDataType_hertz EntitySensorDataType = "hertz"

    EntitySensorDataType_celsius EntitySensorDataType = "celsius"

    EntitySensorDataType_percentRH EntitySensorDataType = "percentRH"

    EntitySensorDataType_rpm EntitySensorDataType = "rpm"

    EntitySensorDataType_cmm EntitySensorDataType = "cmm"

    EntitySensorDataType_truthvalue EntitySensorDataType = "truthvalue"
)

// EntitySensorStatus represents of-range, jittery, or wildly fluctuating readings.
type EntitySensorStatus string

const (
    EntitySensorStatus_ok EntitySensorStatus = "ok"

    EntitySensorStatus_unavailable EntitySensorStatus = "unavailable"

    EntitySensorStatus_nonoperational EntitySensorStatus = "nonoperational"
)

// EntitySensorDataScale represents object of type EntitySensorValue.
type EntitySensorDataScale string

const (
    EntitySensorDataScale_yocto EntitySensorDataScale = "yocto"

    EntitySensorDataScale_zepto EntitySensorDataScale = "zepto"

    EntitySensorDataScale_atto EntitySensorDataScale = "atto"

    EntitySensorDataScale_femto EntitySensorDataScale = "femto"

    EntitySensorDataScale_pico EntitySensorDataScale = "pico"

    EntitySensorDataScale_nano EntitySensorDataScale = "nano"

    EntitySensorDataScale_micro EntitySensorDataScale = "micro"

    EntitySensorDataScale_milli EntitySensorDataScale = "milli"

    EntitySensorDataScale_units EntitySensorDataScale = "units"

    EntitySensorDataScale_kilo EntitySensorDataScale = "kilo"

    EntitySensorDataScale_mega EntitySensorDataScale = "mega"

    EntitySensorDataScale_giga EntitySensorDataScale = "giga"

    EntitySensorDataScale_tera EntitySensorDataScale = "tera"

    EntitySensorDataScale_exa EntitySensorDataScale = "exa"

    EntitySensorDataScale_peta EntitySensorDataScale = "peta"

    EntitySensorDataScale_zetta EntitySensorDataScale = "zetta"

    EntitySensorDataScale_yotta EntitySensorDataScale = "yotta"
)

// ENTITYSENSORMIB
type ENTITYSENSORMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains one row per physical sensor represented by an
    // associated row in the entPhysicalTable.
    EntPhySensorTable ENTITYSENSORMIB_EntPhySensorTable
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetEntityData() *types.CommonEntityData {
    eNTITYSENSORMIB.EntityData.YFilter = eNTITYSENSORMIB.YFilter
    eNTITYSENSORMIB.EntityData.YangName = "ENTITY-SENSOR-MIB"
    eNTITYSENSORMIB.EntityData.BundleName = "cisco_ios_xe"
    eNTITYSENSORMIB.EntityData.ParentYangName = "ENTITY-SENSOR-MIB"
    eNTITYSENSORMIB.EntityData.SegmentPath = "ENTITY-SENSOR-MIB:ENTITY-SENSOR-MIB"
    eNTITYSENSORMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eNTITYSENSORMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eNTITYSENSORMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eNTITYSENSORMIB.EntityData.Children = types.NewOrderedMap()
    eNTITYSENSORMIB.EntityData.Children.Append("entPhySensorTable", types.YChild{"EntPhySensorTable", &eNTITYSENSORMIB.EntPhySensorTable})
    eNTITYSENSORMIB.EntityData.Leafs = types.NewOrderedMap()

    eNTITYSENSORMIB.EntityData.YListKeys = []string {}

    return &(eNTITYSENSORMIB.EntityData)
}

// ENTITYSENSORMIB_EntPhySensorTable
// This table contains one row per physical sensor represented
// by an associated row in the entPhysicalTable.
type ENTITYSENSORMIB_EntPhySensorTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical sensor.  An entry in this table
    // describes the present reading of a sensor, the measurement units and scale,
    // and sensor operational status.  Entries are created in this table by the
    // agent.  An entry for each physical sensor SHOULD be created at the same
    // time as the associated entPhysicalEntry.  An entry SHOULD be destroyed if
    // the associated entPhysicalEntry is destroyed. The type is slice of
    // ENTITYSENSORMIB_EntPhySensorTable_EntPhySensorEntry.
    EntPhySensorEntry []*ENTITYSENSORMIB_EntPhySensorTable_EntPhySensorEntry
}

func (entPhySensorTable *ENTITYSENSORMIB_EntPhySensorTable) GetEntityData() *types.CommonEntityData {
    entPhySensorTable.EntityData.YFilter = entPhySensorTable.YFilter
    entPhySensorTable.EntityData.YangName = "entPhySensorTable"
    entPhySensorTable.EntityData.BundleName = "cisco_ios_xe"
    entPhySensorTable.EntityData.ParentYangName = "ENTITY-SENSOR-MIB"
    entPhySensorTable.EntityData.SegmentPath = "entPhySensorTable"
    entPhySensorTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entPhySensorTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entPhySensorTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entPhySensorTable.EntityData.Children = types.NewOrderedMap()
    entPhySensorTable.EntityData.Children.Append("entPhySensorEntry", types.YChild{"EntPhySensorEntry", nil})
    for i := range entPhySensorTable.EntPhySensorEntry {
        entPhySensorTable.EntityData.Children.Append(types.GetSegmentPath(entPhySensorTable.EntPhySensorEntry[i]), types.YChild{"EntPhySensorEntry", entPhySensorTable.EntPhySensorEntry[i]})
    }
    entPhySensorTable.EntityData.Leafs = types.NewOrderedMap()

    entPhySensorTable.EntityData.YListKeys = []string {}

    return &(entPhySensorTable.EntityData)
}

// ENTITYSENSORMIB_EntPhySensorTable_EntPhySensorEntry
// Information about a particular physical sensor.
// 
// An entry in this table describes the present reading of a
// sensor, the measurement units and scale, and sensor
// operational status.
// 
// Entries are created in this table by the agent.  An entry
// for each physical sensor SHOULD be created at the same time
// as the associated entPhysicalEntry.  An entry SHOULD be
// destroyed if the associated entPhysicalEntry is destroyed.
type ENTITYSENSORMIB_EntPhySensorTable_EntPhySensorEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The type of data returned by the associated entPhySensorValue object.  This
    // object SHOULD be set by the agent during entry creation, and the value
    // SHOULD NOT change during operation. The type is EntitySensorDataType.
    EntPhySensorType interface{}

    // The exponent to apply to values returned by the associated
    // entPhySensorValue object.  This object SHOULD be set by the agent during
    // entry creation, and the value SHOULD NOT change during operation. The type
    // is EntitySensorDataScale.
    EntPhySensorScale interface{}

    // The number of decimal places of precision in fixed-point sensor values
    // returned by the associated entPhySensorValue object.  This object SHOULD be
    // set to '0' when the associated entPhySensorType value is not a fixed-point
    // type: e.g., 'percentRH(9)', 'rpm(10)', 'cmm(11)', or 'truthvalue(12)'. 
    // This object SHOULD be set by the agent during entry creation, and the value
    // SHOULD NOT change during operation. The type is interface{} with range:
    // -8..9.
    EntPhySensorPrecision interface{}

    // The most recent measurement obtained by the agent for this sensor.  To
    // correctly interpret the value of this object, the associated
    // entPhySensorType, entPhySensorScale, and entPhySensorPrecision objects must
    // also be examined. The type is interface{} with range:
    // -1000000000..1073741823.
    EntPhySensorValue interface{}

    // The operational status of the sensor. The type is EntitySensorStatus.
    EntPhySensorOperStatus interface{}

    // A textual description of the data units that should be used in the display
    // of entPhySensorValue. The type is string.
    EntPhySensorUnitsDisplay interface{}

    // The value of sysUpTime at the time the status and/or value of this sensor
    // was last obtained by the agent. The type is interface{} with range:
    // 0..4294967295.
    EntPhySensorValueTimeStamp interface{}

    // An indication of the frequency that the agent updates the associated
    // entPhySensorValue object, representing in milliseconds.  The value zero
    // indicates:      - the sensor value is updated on demand (e.g.,       when
    // polled by the agent for a get-request),     - the sensor value is updated
    // when the sensor       value changes (event-driven),     - the agent does
    // not know the update rate. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    EntPhySensorValueUpdateRate interface{}
}

func (entPhySensorEntry *ENTITYSENSORMIB_EntPhySensorTable_EntPhySensorEntry) GetEntityData() *types.CommonEntityData {
    entPhySensorEntry.EntityData.YFilter = entPhySensorEntry.YFilter
    entPhySensorEntry.EntityData.YangName = "entPhySensorEntry"
    entPhySensorEntry.EntityData.BundleName = "cisco_ios_xe"
    entPhySensorEntry.EntityData.ParentYangName = "entPhySensorTable"
    entPhySensorEntry.EntityData.SegmentPath = "entPhySensorEntry" + types.AddKeyToken(entPhySensorEntry.EntPhysicalIndex, "entPhysicalIndex")
    entPhySensorEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entPhySensorEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entPhySensorEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entPhySensorEntry.EntityData.Children = types.NewOrderedMap()
    entPhySensorEntry.EntityData.Leafs = types.NewOrderedMap()
    entPhySensorEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entPhySensorEntry.EntPhysicalIndex})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorType", types.YLeaf{"EntPhySensorType", entPhySensorEntry.EntPhySensorType})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorScale", types.YLeaf{"EntPhySensorScale", entPhySensorEntry.EntPhySensorScale})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorPrecision", types.YLeaf{"EntPhySensorPrecision", entPhySensorEntry.EntPhySensorPrecision})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorValue", types.YLeaf{"EntPhySensorValue", entPhySensorEntry.EntPhySensorValue})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorOperStatus", types.YLeaf{"EntPhySensorOperStatus", entPhySensorEntry.EntPhySensorOperStatus})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorUnitsDisplay", types.YLeaf{"EntPhySensorUnitsDisplay", entPhySensorEntry.EntPhySensorUnitsDisplay})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorValueTimeStamp", types.YLeaf{"EntPhySensorValueTimeStamp", entPhySensorEntry.EntPhySensorValueTimeStamp})
    entPhySensorEntry.EntityData.Leafs.Append("entPhySensorValueUpdateRate", types.YLeaf{"EntPhySensorValueUpdateRate", entPhySensorEntry.EntPhySensorValueUpdateRate})

    entPhySensorEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(entPhySensorEntry.EntityData)
}

