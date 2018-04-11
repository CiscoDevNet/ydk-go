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

// EntitySensorStatus represents of-range, jittery, or wildly fluctuating readings.
type EntitySensorStatus string

const (
    EntitySensorStatus_ok EntitySensorStatus = "ok"

    EntitySensorStatus_unavailable EntitySensorStatus = "unavailable"

    EntitySensorStatus_nonoperational EntitySensorStatus = "nonoperational"
)

// ENTITYSENSORMIB
type ENTITYSENSORMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains one row per physical sensor represented by an
    // associated row in the entPhysicalTable.
    Entphysensortable ENTITYSENSORMIB_Entphysensortable
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

    eNTITYSENSORMIB.EntityData.Children = make(map[string]types.YChild)
    eNTITYSENSORMIB.EntityData.Children["entPhySensorTable"] = types.YChild{"Entphysensortable", &eNTITYSENSORMIB.Entphysensortable}
    eNTITYSENSORMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eNTITYSENSORMIB.EntityData)
}

// ENTITYSENSORMIB_Entphysensortable
// This table contains one row per physical sensor represented
// by an associated row in the entPhysicalTable.
type ENTITYSENSORMIB_Entphysensortable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical sensor.  An entry in this table
    // describes the present reading of a sensor, the measurement units and scale,
    // and sensor operational status.  Entries are created in this table by the
    // agent.  An entry for each physical sensor SHOULD be created at the same
    // time as the associated entPhysicalEntry.  An entry SHOULD be destroyed if
    // the associated entPhysicalEntry is destroyed. The type is slice of
    // ENTITYSENSORMIB_Entphysensortable_Entphysensorentry.
    Entphysensorentry []ENTITYSENSORMIB_Entphysensortable_Entphysensorentry
}

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetEntityData() *types.CommonEntityData {
    entphysensortable.EntityData.YFilter = entphysensortable.YFilter
    entphysensortable.EntityData.YangName = "entPhySensorTable"
    entphysensortable.EntityData.BundleName = "cisco_ios_xe"
    entphysensortable.EntityData.ParentYangName = "ENTITY-SENSOR-MIB"
    entphysensortable.EntityData.SegmentPath = "entPhySensorTable"
    entphysensortable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entphysensortable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entphysensortable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entphysensortable.EntityData.Children = make(map[string]types.YChild)
    entphysensortable.EntityData.Children["entPhySensorEntry"] = types.YChild{"Entphysensorentry", nil}
    for i := range entphysensortable.Entphysensorentry {
        entphysensortable.EntityData.Children[types.GetSegmentPath(&entphysensortable.Entphysensorentry[i])] = types.YChild{"Entphysensorentry", &entphysensortable.Entphysensorentry[i]}
    }
    entphysensortable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entphysensortable.EntityData)
}

// ENTITYSENSORMIB_Entphysensortable_Entphysensorentry
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
type ENTITYSENSORMIB_Entphysensortable_Entphysensorentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The type of data returned by the associated entPhySensorValue object.  This
    // object SHOULD be set by the agent during entry creation, and the value
    // SHOULD NOT change during operation. The type is EntitySensorDataType.
    Entphysensortype interface{}

    // The exponent to apply to values returned by the associated
    // entPhySensorValue object.  This object SHOULD be set by the agent during
    // entry creation, and the value SHOULD NOT change during operation. The type
    // is EntitySensorDataScale.
    Entphysensorscale interface{}

    // The number of decimal places of precision in fixed-point sensor values
    // returned by the associated entPhySensorValue object.  This object SHOULD be
    // set to '0' when the associated entPhySensorType value is not a fixed-point
    // type: e.g., 'percentRH(9)', 'rpm(10)', 'cmm(11)', or 'truthvalue(12)'. 
    // This object SHOULD be set by the agent during entry creation, and the value
    // SHOULD NOT change during operation. The type is interface{} with range:
    // -8..9.
    Entphysensorprecision interface{}

    // The most recent measurement obtained by the agent for this sensor.  To
    // correctly interpret the value of this object, the associated
    // entPhySensorType, entPhySensorScale, and entPhySensorPrecision objects must
    // also be examined. The type is interface{} with range:
    // -1000000000..1073741823.
    Entphysensorvalue interface{}

    // The operational status of the sensor. The type is EntitySensorStatus.
    Entphysensoroperstatus interface{}

    // A textual description of the data units that should be used in the display
    // of entPhySensorValue. The type is string.
    Entphysensorunitsdisplay interface{}

    // The value of sysUpTime at the time the status and/or value of this sensor
    // was last obtained by the agent. The type is interface{} with range:
    // 0..4294967295.
    Entphysensorvaluetimestamp interface{}

    // An indication of the frequency that the agent updates the associated
    // entPhySensorValue object, representing in milliseconds.  The value zero
    // indicates:      - the sensor value is updated on demand (e.g.,       when
    // polled by the agent for a get-request),     - the sensor value is updated
    // when the sensor       value changes (event-driven),     - the agent does
    // not know the update rate. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Entphysensorvalueupdaterate interface{}
}

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetEntityData() *types.CommonEntityData {
    entphysensorentry.EntityData.YFilter = entphysensorentry.YFilter
    entphysensorentry.EntityData.YangName = "entPhySensorEntry"
    entphysensorentry.EntityData.BundleName = "cisco_ios_xe"
    entphysensorentry.EntityData.ParentYangName = "entPhySensorTable"
    entphysensorentry.EntityData.SegmentPath = "entPhySensorEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysensorentry.Entphysicalindex) + "']"
    entphysensorentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entphysensorentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entphysensorentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entphysensorentry.EntityData.Children = make(map[string]types.YChild)
    entphysensorentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entphysensorentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entphysensorentry.Entphysicalindex}
    entphysensorentry.EntityData.Leafs["entPhySensorType"] = types.YLeaf{"Entphysensortype", entphysensorentry.Entphysensortype}
    entphysensorentry.EntityData.Leafs["entPhySensorScale"] = types.YLeaf{"Entphysensorscale", entphysensorentry.Entphysensorscale}
    entphysensorentry.EntityData.Leafs["entPhySensorPrecision"] = types.YLeaf{"Entphysensorprecision", entphysensorentry.Entphysensorprecision}
    entphysensorentry.EntityData.Leafs["entPhySensorValue"] = types.YLeaf{"Entphysensorvalue", entphysensorentry.Entphysensorvalue}
    entphysensorentry.EntityData.Leafs["entPhySensorOperStatus"] = types.YLeaf{"Entphysensoroperstatus", entphysensorentry.Entphysensoroperstatus}
    entphysensorentry.EntityData.Leafs["entPhySensorUnitsDisplay"] = types.YLeaf{"Entphysensorunitsdisplay", entphysensorentry.Entphysensorunitsdisplay}
    entphysensorentry.EntityData.Leafs["entPhySensorValueTimeStamp"] = types.YLeaf{"Entphysensorvaluetimestamp", entphysensorentry.Entphysensorvaluetimestamp}
    entphysensorentry.EntityData.Leafs["entPhySensorValueUpdateRate"] = types.YLeaf{"Entphysensorvalueupdaterate", entphysensorentry.Entphysensorvalueupdaterate}
    return &(entphysensorentry.EntityData)
}

