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
    parent types.Entity
    YFilter yfilter.YFilter

    // This table contains one row per physical sensor represented by an
    // associated row in the entPhysicalTable.
    Entphysensortable ENTITYSENSORMIB_Entphysensortable
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetFilter() yfilter.YFilter { return eNTITYSENSORMIB.YFilter }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) SetFilter(yf yfilter.YFilter) { eNTITYSENSORMIB.YFilter = yf }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetGoName(yname string) string {
    if yname == "entPhySensorTable" { return "Entphysensortable" }
    return ""
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetSegmentPath() string {
    return "ENTITY-SENSOR-MIB:ENTITY-SENSOR-MIB"
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entPhySensorTable" {
        return &eNTITYSENSORMIB.Entphysensortable
    }
    return nil
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entPhySensorTable"] = &eNTITYSENSORMIB.Entphysensortable
    return children
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetBundleName() string { return "cisco_ios_xe" }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetYangName() string { return "ENTITY-SENSOR-MIB" }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) SetParent(parent types.Entity) { eNTITYSENSORMIB.parent = parent }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetParent() types.Entity { return eNTITYSENSORMIB.parent }

func (eNTITYSENSORMIB *ENTITYSENSORMIB) GetParentYangName() string { return "ENTITY-SENSOR-MIB" }

// ENTITYSENSORMIB_Entphysensortable
// This table contains one row per physical sensor represented
// by an associated row in the entPhysicalTable.
type ENTITYSENSORMIB_Entphysensortable struct {
    parent types.Entity
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

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetFilter() yfilter.YFilter { return entphysensortable.YFilter }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) SetFilter(yf yfilter.YFilter) { entphysensortable.YFilter = yf }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetGoName(yname string) string {
    if yname == "entPhySensorEntry" { return "Entphysensorentry" }
    return ""
}

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetSegmentPath() string {
    return "entPhySensorTable"
}

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entPhySensorEntry" {
        for _, c := range entphysensortable.Entphysensorentry {
            if entphysensortable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYSENSORMIB_Entphysensortable_Entphysensorentry{}
        entphysensortable.Entphysensorentry = append(entphysensortable.Entphysensorentry, child)
        return &entphysensortable.Entphysensorentry[len(entphysensortable.Entphysensorentry)-1]
    }
    return nil
}

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entphysensortable.Entphysensorentry {
        children[entphysensortable.Entphysensorentry[i].GetSegmentPath()] = &entphysensortable.Entphysensorentry[i]
    }
    return children
}

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetBundleName() string { return "cisco_ios_xe" }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetYangName() string { return "entPhySensorTable" }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) SetParent(parent types.Entity) { entphysensortable.parent = parent }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetParent() types.Entity { return entphysensortable.parent }

func (entphysensortable *ENTITYSENSORMIB_Entphysensortable) GetParentYangName() string { return "ENTITY-SENSOR-MIB" }

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
    parent types.Entity
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

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetFilter() yfilter.YFilter { return entphysensorentry.YFilter }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) SetFilter(yf yfilter.YFilter) { entphysensorentry.YFilter = yf }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entPhySensorType" { return "Entphysensortype" }
    if yname == "entPhySensorScale" { return "Entphysensorscale" }
    if yname == "entPhySensorPrecision" { return "Entphysensorprecision" }
    if yname == "entPhySensorValue" { return "Entphysensorvalue" }
    if yname == "entPhySensorOperStatus" { return "Entphysensoroperstatus" }
    if yname == "entPhySensorUnitsDisplay" { return "Entphysensorunitsdisplay" }
    if yname == "entPhySensorValueTimeStamp" { return "Entphysensorvaluetimestamp" }
    if yname == "entPhySensorValueUpdateRate" { return "Entphysensorvalueupdaterate" }
    return ""
}

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetSegmentPath() string {
    return "entPhySensorEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysensorentry.Entphysicalindex) + "']"
}

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entphysensorentry.Entphysicalindex
    leafs["entPhySensorType"] = entphysensorentry.Entphysensortype
    leafs["entPhySensorScale"] = entphysensorentry.Entphysensorscale
    leafs["entPhySensorPrecision"] = entphysensorentry.Entphysensorprecision
    leafs["entPhySensorValue"] = entphysensorentry.Entphysensorvalue
    leafs["entPhySensorOperStatus"] = entphysensorentry.Entphysensoroperstatus
    leafs["entPhySensorUnitsDisplay"] = entphysensorentry.Entphysensorunitsdisplay
    leafs["entPhySensorValueTimeStamp"] = entphysensorentry.Entphysensorvaluetimestamp
    leafs["entPhySensorValueUpdateRate"] = entphysensorentry.Entphysensorvalueupdaterate
    return leafs
}

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetBundleName() string { return "cisco_ios_xe" }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetYangName() string { return "entPhySensorEntry" }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) SetParent(parent types.Entity) { entphysensorentry.parent = parent }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetParent() types.Entity { return entphysensorentry.parent }

func (entphysensorentry *ENTITYSENSORMIB_Entphysensortable_Entphysensorentry) GetParentYangName() string { return "entPhySensorTable" }

