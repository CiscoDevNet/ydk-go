// The CISCO-ENTITY-SENSOR-MIB is used to monitor
// the values of sensors in the Entity-MIB (RFC 2037) 
// entPhysicalTable.
package cisco_entity_sensor_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_entity_sensor_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ENTITY-SENSOR-MIB CISCO-ENTITY-SENSOR-MIB}", reflect.TypeOf(CISCOENTITYSENSORMIB{}))
    ydk.RegisterEntity("CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB", reflect.TypeOf(CISCOENTITYSENSORMIB{}))
}

// SensorThresholdSeverity represents               threshold.
type SensorThresholdSeverity string

const (
    SensorThresholdSeverity_other SensorThresholdSeverity = "other"

    SensorThresholdSeverity_minor SensorThresholdSeverity = "minor"

    SensorThresholdSeverity_major SensorThresholdSeverity = "major"

    SensorThresholdSeverity_critical SensorThresholdSeverity = "critical"
)

// SensorStatus represents readings.
type SensorStatus string

const (
    SensorStatus_ok SensorStatus = "ok"

    SensorStatus_unavailable SensorStatus = "unavailable"

    SensorStatus_nonoperational SensorStatus = "nonoperational"
)

// SensorDataType represents dBm(14):         dB relative to 1mW of power
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

    SensorDataType_dBm SensorDataType = "dBm"
)

// SensorThresholdRelation represents                     the threshold value
type SensorThresholdRelation string

const (
    SensorThresholdRelation_lessThan SensorThresholdRelation = "lessThan"

    SensorThresholdRelation_lessOrEqual SensorThresholdRelation = "lessOrEqual"

    SensorThresholdRelation_greaterThan SensorThresholdRelation = "greaterThan"

    SensorThresholdRelation_greaterOrEqual SensorThresholdRelation = "greaterOrEqual"

    SensorThresholdRelation_equalTo SensorThresholdRelation = "equalTo"

    SensorThresholdRelation_notEqualTo SensorThresholdRelation = "notEqualTo"
)

// SensorDataScale represents International System of Units (SI) prefixes.
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

// CISCOENTITYSENSORMIB
type CISCOENTITYSENSORMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Entsensorglobalobjects CISCOENTITYSENSORMIB_Entsensorglobalobjects

    // This table lists the type, scale, and present value of a sensor listed in
    // the Entity-MIB entPhysicalTable.
    Entsensorvaluetable CISCOENTITYSENSORMIB_Entsensorvaluetable

    // This table lists the threshold severity, relation, and comparison value,
    // for a sensor listed in the Entity-MIB  entPhysicalTable.
    Entsensorthresholdtable CISCOENTITYSENSORMIB_Entsensorthresholdtable
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetFilter() yfilter.YFilter { return cISCOENTITYSENSORMIB.YFilter }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) SetFilter(yf yfilter.YFilter) { cISCOENTITYSENSORMIB.YFilter = yf }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetGoName(yname string) string {
    if yname == "entSensorGlobalObjects" { return "Entsensorglobalobjects" }
    if yname == "entSensorValueTable" { return "Entsensorvaluetable" }
    if yname == "entSensorThresholdTable" { return "Entsensorthresholdtable" }
    return ""
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetSegmentPath() string {
    return "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB"
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entSensorGlobalObjects" {
        return &cISCOENTITYSENSORMIB.Entsensorglobalobjects
    }
    if childYangName == "entSensorValueTable" {
        return &cISCOENTITYSENSORMIB.Entsensorvaluetable
    }
    if childYangName == "entSensorThresholdTable" {
        return &cISCOENTITYSENSORMIB.Entsensorthresholdtable
    }
    return nil
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entSensorGlobalObjects"] = &cISCOENTITYSENSORMIB.Entsensorglobalobjects
    children["entSensorValueTable"] = &cISCOENTITYSENSORMIB.Entsensorvaluetable
    children["entSensorThresholdTable"] = &cISCOENTITYSENSORMIB.Entsensorthresholdtable
    return children
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetYangName() string { return "CISCO-ENTITY-SENSOR-MIB" }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) SetParent(parent types.Entity) { cISCOENTITYSENSORMIB.parent = parent }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetParent() types.Entity { return cISCOENTITYSENSORMIB.parent }

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetParentYangName() string { return "CISCO-ENTITY-SENSOR-MIB" }

// CISCOENTITYSENSORMIB_Entsensorglobalobjects
type CISCOENTITYSENSORMIB_Entsensorglobalobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This variable enables the generation of entSensorThresholdNotification
    // globally on the device. If this object value is 'false', then no
    // entSensorThresholdNotification will be generated on this device. If this
    // object value is 'true', then whether a  entSensorThresholdNotification for
    // a threshold will be generated or not depends on the instance value of
    // entSensorThresholdNotificationEnable for that threshold. The type is bool.
    Entsensorthreshnotifglobalenable interface{}
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetFilter() yfilter.YFilter { return entsensorglobalobjects.YFilter }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) SetFilter(yf yfilter.YFilter) { entsensorglobalobjects.YFilter = yf }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetGoName(yname string) string {
    if yname == "entSensorThreshNotifGlobalEnable" { return "Entsensorthreshnotifglobalenable" }
    return ""
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetSegmentPath() string {
    return "entSensorGlobalObjects"
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entSensorThreshNotifGlobalEnable"] = entsensorglobalobjects.Entsensorthreshnotifglobalenable
    return leafs
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetBundleName() string { return "cisco_ios_xe" }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetYangName() string { return "entSensorGlobalObjects" }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) SetParent(parent types.Entity) { entsensorglobalobjects.parent = parent }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetParent() types.Entity { return entsensorglobalobjects.parent }

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetParentYangName() string { return "CISCO-ENTITY-SENSOR-MIB" }

// CISCOENTITYSENSORMIB_Entsensorvaluetable
// This table lists the type, scale, and present value
// of a sensor listed in the Entity-MIB entPhysicalTable.
type CISCOENTITYSENSORMIB_Entsensorvaluetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entSensorValueTable entry describes the present reading of a sensor, the
    // measurement units and scale, and sensor operational status. The type is
    // slice of CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry.
    Entsensorvalueentry []CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetFilter() yfilter.YFilter { return entsensorvaluetable.YFilter }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) SetFilter(yf yfilter.YFilter) { entsensorvaluetable.YFilter = yf }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetGoName(yname string) string {
    if yname == "entSensorValueEntry" { return "Entsensorvalueentry" }
    return ""
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetSegmentPath() string {
    return "entSensorValueTable"
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entSensorValueEntry" {
        for _, c := range entsensorvaluetable.Entsensorvalueentry {
            if entsensorvaluetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry{}
        entsensorvaluetable.Entsensorvalueentry = append(entsensorvaluetable.Entsensorvalueentry, child)
        return &entsensorvaluetable.Entsensorvalueentry[len(entsensorvaluetable.Entsensorvalueentry)-1]
    }
    return nil
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entsensorvaluetable.Entsensorvalueentry {
        children[entsensorvaluetable.Entsensorvalueentry[i].GetSegmentPath()] = &entsensorvaluetable.Entsensorvalueentry[i]
    }
    return children
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetBundleName() string { return "cisco_ios_xe" }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetYangName() string { return "entSensorValueTable" }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) SetParent(parent types.Entity) { entsensorvaluetable.parent = parent }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetParent() types.Entity { return entsensorvaluetable.parent }

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetParentYangName() string { return "CISCO-ENTITY-SENSOR-MIB" }

// CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry
// An entSensorValueTable entry describes the
// present reading of a sensor, the measurement units
// and scale, and sensor operational status.
type CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This variable indicates the type of data reported by the entSensorValue. 
    // This variable is set by the agent at start-up and the value does not change
    // during operation. The type is SensorDataType.
    Entsensortype interface{}

    // This variable indicates the exponent to apply to sensor values reported by
    // entSensorValue.  This variable is set by the agent at start-up and the
    // value does not change during operation. The type is SensorDataScale.
    Entsensorscale interface{}

    // This variable indicates the number of decimal places of precision in
    // fixed-point sensor values reported by entSensorValue.  This variable is set
    // to 0 when entSensorType is not a fixed-point type: e.g.'percentRH(9)', 
    // 'rpm(10)', 'cmm(11)', or 'truthvalue(12)'.  This variable is set by the
    // agent at start-up and the value does not change during operation. The type
    // is interface{} with range: -8..9.
    Entsensorprecision interface{}

    // This variable reports the most recent measurement seen by the sensor.  To
    // correctly display or interpret this variable's value,  you must also know
    // entSensorType, entSensorScale, and  entSensorPrecision.  However, you can
    // compare entSensorValue with the threshold values given in
    // entSensorThresholdTable without any semantic knowledge. The type is
    // interface{} with range: -1000000000..1073741823.
    Entsensorvalue interface{}

    // This variable indicates the present operational status of the sensor. The
    // type is SensorStatus.
    Entsensorstatus interface{}

    // This variable indicates the age of the value reported by entSensorValue.
    // The type is interface{} with range: 0..4294967295.
    Entsensorvaluetimestamp interface{}

    // This variable indicates the rate that the agent updates entSensorValue. The
    // type is interface{} with range: 0..999999999. Units are seconds.
    Entsensorvalueupdaterate interface{}

    // This object identifies the physical entity for which the sensor is taking
    // measurements.  For example, for a sensor measuring the voltage output of a
    // power-supply, this object would be the entPhysicalIndex of that
    // power-supply; for a sensor measuring the temperature inside one chassis of
    // a multi-chassis system, this object would be the enPhysicalIndex of that
    // chassis.  This object has a value of zero when the physical entity for
    // which the sensor is taking measurements can not be represented by any one
    // row in the entPhysicalTable, or that there is no such physical entity. The
    // type is interface{} with range: 0..2147483647.
    Entsensormeasuredentity interface{}
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetFilter() yfilter.YFilter { return entsensorvalueentry.YFilter }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) SetFilter(yf yfilter.YFilter) { entsensorvalueentry.YFilter = yf }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entSensorType" { return "Entsensortype" }
    if yname == "entSensorScale" { return "Entsensorscale" }
    if yname == "entSensorPrecision" { return "Entsensorprecision" }
    if yname == "entSensorValue" { return "Entsensorvalue" }
    if yname == "entSensorStatus" { return "Entsensorstatus" }
    if yname == "entSensorValueTimeStamp" { return "Entsensorvaluetimestamp" }
    if yname == "entSensorValueUpdateRate" { return "Entsensorvalueupdaterate" }
    if yname == "entSensorMeasuredEntity" { return "Entsensormeasuredentity" }
    return ""
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetSegmentPath() string {
    return "entSensorValueEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entsensorvalueentry.Entphysicalindex) + "']"
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entsensorvalueentry.Entphysicalindex
    leafs["entSensorType"] = entsensorvalueentry.Entsensortype
    leafs["entSensorScale"] = entsensorvalueentry.Entsensorscale
    leafs["entSensorPrecision"] = entsensorvalueentry.Entsensorprecision
    leafs["entSensorValue"] = entsensorvalueentry.Entsensorvalue
    leafs["entSensorStatus"] = entsensorvalueentry.Entsensorstatus
    leafs["entSensorValueTimeStamp"] = entsensorvalueentry.Entsensorvaluetimestamp
    leafs["entSensorValueUpdateRate"] = entsensorvalueentry.Entsensorvalueupdaterate
    leafs["entSensorMeasuredEntity"] = entsensorvalueentry.Entsensormeasuredentity
    return leafs
}

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetBundleName() string { return "cisco_ios_xe" }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetYangName() string { return "entSensorValueEntry" }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) SetParent(parent types.Entity) { entsensorvalueentry.parent = parent }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetParent() types.Entity { return entsensorvalueentry.parent }

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetParentYangName() string { return "entSensorValueTable" }

// CISCOENTITYSENSORMIB_Entsensorthresholdtable
// This table lists the threshold severity, relation, and
// comparison value, for a sensor listed in the Entity-MIB 
// entPhysicalTable.
type CISCOENTITYSENSORMIB_Entsensorthresholdtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entSensorThresholdTable entry describes the thresholds for a sensor: the
    // threshold severity, the threshold value, the relation, and the  evaluation
    // of the threshold.  Only entities of type sensor(8) are listed in this
    // table. Only pre-configured thresholds are listed in this table.  Users can
    // create sensor-value monitoring instruments in different ways, such as RMON
    // alarms, Expression-MIB, etc.  Entries are created by the agent at system
    // startup and FRU insertion.  Entries are deleted by the agent at FRU
    // removal. The type is slice of
    // CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry.
    Entsensorthresholdentry []CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetFilter() yfilter.YFilter { return entsensorthresholdtable.YFilter }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) SetFilter(yf yfilter.YFilter) { entsensorthresholdtable.YFilter = yf }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetGoName(yname string) string {
    if yname == "entSensorThresholdEntry" { return "Entsensorthresholdentry" }
    return ""
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetSegmentPath() string {
    return "entSensorThresholdTable"
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entSensorThresholdEntry" {
        for _, c := range entsensorthresholdtable.Entsensorthresholdentry {
            if entsensorthresholdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry{}
        entsensorthresholdtable.Entsensorthresholdentry = append(entsensorthresholdtable.Entsensorthresholdentry, child)
        return &entsensorthresholdtable.Entsensorthresholdentry[len(entsensorthresholdtable.Entsensorthresholdentry)-1]
    }
    return nil
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entsensorthresholdtable.Entsensorthresholdentry {
        children[entsensorthresholdtable.Entsensorthresholdentry[i].GetSegmentPath()] = &entsensorthresholdtable.Entsensorthresholdentry[i]
    }
    return children
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetBundleName() string { return "cisco_ios_xe" }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetYangName() string { return "entSensorThresholdTable" }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) SetParent(parent types.Entity) { entsensorthresholdtable.parent = parent }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetParent() types.Entity { return entsensorthresholdtable.parent }

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetParentYangName() string { return "CISCO-ENTITY-SENSOR-MIB" }

// CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry
// An entSensorThresholdTable entry describes the
// thresholds for a sensor: the threshold severity,
// the threshold value, the relation, and the 
// evaluation of the threshold.
// 
// Only entities of type sensor(8) are listed in this table.
// Only pre-configured thresholds are listed in this table.
// 
// Users can create sensor-value monitoring instruments
// in different ways, such as RMON alarms, Expression-MIB, etc.
// 
// Entries are created by the agent at system startup and
// FRU insertion.  Entries are deleted by the agent at
// FRU removal.
type CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // entSensorThresholdTable. This index permits the same sensor to have several
    // different thresholds. The type is interface{} with range: 1..99999999.
    Entsensorthresholdindex interface{}

    // This variable indicates the severity of this threshold. The type is
    // SensorThresholdSeverity.
    Entsensorthresholdseverity interface{}

    // This variable indicates the relation between sensor value (entSensorValue)
    // and threshold value (entSensorThresholdValue),  required to trigger the
    // alarm.  when evaluating the relation,  entSensorValue is on the left of
    // entSensorThresholdRelation,  entSensorThresholdValue is on the right.   in
    // pseudo-code, the evaluation-alarm mechanism is:  ... if (entSensorStatus ==
    // ok) then     if (evaluate(entSensorValue, entSensorThresholdRelation,      
    // entSensorThresholdValue))      then         if
    // (entSensorThresholdNotificationEnable == true))          then            
    // raise_alarm(sensor's entPhysicalIndex);         endif     endif endif ...
    // The type is SensorThresholdRelation.
    Entsensorthresholdrelation interface{}

    // This variable indicates the value of the threshold.  To correctly display
    // or interpret this variable's value,  you must also know entSensorType,
    // entSensorScale, and  entSensorPrecision.  However, you can directly compare
    // entSensorValue  with the threshold values given in entSensorThresholdTable 
    // without any semantic knowledge. The type is interface{} with range:
    // -1000000000..1073741823.
    Entsensorthresholdvalue interface{}

    // This variable indicates the result of the most recent evaluation of the
    // threshold.  If the threshold condition is true,
    // entSensorThresholdEvaluation  is true(1).  If the threshold condition is
    // false,  entSensorThresholdEvaluation is false(2).  Thresholds are evaluated
    // at the rate indicated by  entSensorValueUpdateRate. The type is bool.
    Entsensorthresholdevaluation interface{}

    // This variable controls generation of entSensorThresholdNotification for
    // this threshold.  When this variable is 'true', generation of 
    // entSensorThresholdNotification is enabled for this threshold. When this
    // variable is 'false',  generation of entSensorThresholdNotification is
    // disabled for this threshold. The type is bool.
    Entsensorthresholdnotificationenable interface{}
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetFilter() yfilter.YFilter { return entsensorthresholdentry.YFilter }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) SetFilter(yf yfilter.YFilter) { entsensorthresholdentry.YFilter = yf }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entSensorThresholdIndex" { return "Entsensorthresholdindex" }
    if yname == "entSensorThresholdSeverity" { return "Entsensorthresholdseverity" }
    if yname == "entSensorThresholdRelation" { return "Entsensorthresholdrelation" }
    if yname == "entSensorThresholdValue" { return "Entsensorthresholdvalue" }
    if yname == "entSensorThresholdEvaluation" { return "Entsensorthresholdevaluation" }
    if yname == "entSensorThresholdNotificationEnable" { return "Entsensorthresholdnotificationenable" }
    return ""
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetSegmentPath() string {
    return "entSensorThresholdEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entsensorthresholdentry.Entphysicalindex) + "']" + "[entSensorThresholdIndex='" + fmt.Sprintf("%v", entsensorthresholdentry.Entsensorthresholdindex) + "']"
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entsensorthresholdentry.Entphysicalindex
    leafs["entSensorThresholdIndex"] = entsensorthresholdentry.Entsensorthresholdindex
    leafs["entSensorThresholdSeverity"] = entsensorthresholdentry.Entsensorthresholdseverity
    leafs["entSensorThresholdRelation"] = entsensorthresholdentry.Entsensorthresholdrelation
    leafs["entSensorThresholdValue"] = entsensorthresholdentry.Entsensorthresholdvalue
    leafs["entSensorThresholdEvaluation"] = entsensorthresholdentry.Entsensorthresholdevaluation
    leafs["entSensorThresholdNotificationEnable"] = entsensorthresholdentry.Entsensorthresholdnotificationenable
    return leafs
}

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetBundleName() string { return "cisco_ios_xe" }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetYangName() string { return "entSensorThresholdEntry" }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) SetParent(parent types.Entity) { entsensorthresholdentry.parent = parent }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetParent() types.Entity { return entsensorthresholdentry.parent }

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetParentYangName() string { return "entSensorThresholdTable" }

