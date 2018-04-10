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

// SensorStatus represents readings.
type SensorStatus string

const (
    SensorStatus_ok SensorStatus = "ok"

    SensorStatus_unavailable SensorStatus = "unavailable"

    SensorStatus_nonoperational SensorStatus = "nonoperational"
)

// SensorThresholdSeverity represents               threshold.
type SensorThresholdSeverity string

const (
    SensorThresholdSeverity_other SensorThresholdSeverity = "other"

    SensorThresholdSeverity_minor SensorThresholdSeverity = "minor"

    SensorThresholdSeverity_major SensorThresholdSeverity = "major"

    SensorThresholdSeverity_critical SensorThresholdSeverity = "critical"
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

// CISCOENTITYSENSORMIB
type CISCOENTITYSENSORMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Entsensorglobalobjects CISCOENTITYSENSORMIB_Entsensorglobalobjects

    // This table lists the type, scale, and present value of a sensor listed in
    // the Entity-MIB entPhysicalTable.
    Entsensorvaluetable CISCOENTITYSENSORMIB_Entsensorvaluetable

    // This table lists the threshold severity, relation, and comparison value,
    // for a sensor listed in the Entity-MIB  entPhysicalTable.
    Entsensorthresholdtable CISCOENTITYSENSORMIB_Entsensorthresholdtable
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYSENSORMIB.EntityData.YFilter = cISCOENTITYSENSORMIB.YFilter
    cISCOENTITYSENSORMIB.EntityData.YangName = "CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENTITYSENSORMIB.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.SegmentPath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENTITYSENSORMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENTITYSENSORMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENTITYSENSORMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENTITYSENSORMIB.EntityData.Children["entSensorGlobalObjects"] = types.YChild{"Entsensorglobalobjects", &cISCOENTITYSENSORMIB.Entsensorglobalobjects}
    cISCOENTITYSENSORMIB.EntityData.Children["entSensorValueTable"] = types.YChild{"Entsensorvaluetable", &cISCOENTITYSENSORMIB.Entsensorvaluetable}
    cISCOENTITYSENSORMIB.EntityData.Children["entSensorThresholdTable"] = types.YChild{"Entsensorthresholdtable", &cISCOENTITYSENSORMIB.Entsensorthresholdtable}
    cISCOENTITYSENSORMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENTITYSENSORMIB.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorglobalobjects
type CISCOENTITYSENSORMIB_Entsensorglobalobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable enables the generation of entSensorThresholdNotification
    // globally on the device. If this object value is 'false', then no
    // entSensorThresholdNotification will be generated on this device. If this
    // object value is 'true', then whether a  entSensorThresholdNotification for
    // a threshold will be generated or not depends on the instance value of
    // entSensorThresholdNotificationEnable for that threshold. The type is bool.
    Entsensorthreshnotifglobalenable interface{}
}

func (entsensorglobalobjects *CISCOENTITYSENSORMIB_Entsensorglobalobjects) GetEntityData() *types.CommonEntityData {
    entsensorglobalobjects.EntityData.YFilter = entsensorglobalobjects.YFilter
    entsensorglobalobjects.EntityData.YangName = "entSensorGlobalObjects"
    entsensorglobalobjects.EntityData.BundleName = "cisco_ios_xe"
    entsensorglobalobjects.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entsensorglobalobjects.EntityData.SegmentPath = "entSensorGlobalObjects"
    entsensorglobalobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entsensorglobalobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entsensorglobalobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entsensorglobalobjects.EntityData.Children = make(map[string]types.YChild)
    entsensorglobalobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    entsensorglobalobjects.EntityData.Leafs["entSensorThreshNotifGlobalEnable"] = types.YLeaf{"Entsensorthreshnotifglobalenable", entsensorglobalobjects.Entsensorthreshnotifglobalenable}
    return &(entsensorglobalobjects.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorvaluetable
// This table lists the type, scale, and present value
// of a sensor listed in the Entity-MIB entPhysicalTable.
type CISCOENTITYSENSORMIB_Entsensorvaluetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entSensorValueTable entry describes the present reading of a sensor, the
    // measurement units and scale, and sensor operational status. The type is
    // slice of CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry.
    Entsensorvalueentry []CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry
}

func (entsensorvaluetable *CISCOENTITYSENSORMIB_Entsensorvaluetable) GetEntityData() *types.CommonEntityData {
    entsensorvaluetable.EntityData.YFilter = entsensorvaluetable.YFilter
    entsensorvaluetable.EntityData.YangName = "entSensorValueTable"
    entsensorvaluetable.EntityData.BundleName = "cisco_ios_xe"
    entsensorvaluetable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entsensorvaluetable.EntityData.SegmentPath = "entSensorValueTable"
    entsensorvaluetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entsensorvaluetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entsensorvaluetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entsensorvaluetable.EntityData.Children = make(map[string]types.YChild)
    entsensorvaluetable.EntityData.Children["entSensorValueEntry"] = types.YChild{"Entsensorvalueentry", nil}
    for i := range entsensorvaluetable.Entsensorvalueentry {
        entsensorvaluetable.EntityData.Children[types.GetSegmentPath(&entsensorvaluetable.Entsensorvalueentry[i])] = types.YChild{"Entsensorvalueentry", &entsensorvaluetable.Entsensorvalueentry[i]}
    }
    entsensorvaluetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entsensorvaluetable.EntityData)
}

// CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry
// An entSensorValueTable entry describes the
// present reading of a sensor, the measurement units
// and scale, and sensor operational status.
type CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry struct {
    EntityData types.CommonEntityData
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

func (entsensorvalueentry *CISCOENTITYSENSORMIB_Entsensorvaluetable_Entsensorvalueentry) GetEntityData() *types.CommonEntityData {
    entsensorvalueentry.EntityData.YFilter = entsensorvalueentry.YFilter
    entsensorvalueentry.EntityData.YangName = "entSensorValueEntry"
    entsensorvalueentry.EntityData.BundleName = "cisco_ios_xe"
    entsensorvalueentry.EntityData.ParentYangName = "entSensorValueTable"
    entsensorvalueentry.EntityData.SegmentPath = "entSensorValueEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entsensorvalueentry.Entphysicalindex) + "']"
    entsensorvalueentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entsensorvalueentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entsensorvalueentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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
// This table lists the threshold severity, relation, and
// comparison value, for a sensor listed in the Entity-MIB 
// entPhysicalTable.
type CISCOENTITYSENSORMIB_Entsensorthresholdtable struct {
    EntityData types.CommonEntityData
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

func (entsensorthresholdtable *CISCOENTITYSENSORMIB_Entsensorthresholdtable) GetEntityData() *types.CommonEntityData {
    entsensorthresholdtable.EntityData.YFilter = entsensorthresholdtable.YFilter
    entsensorthresholdtable.EntityData.YangName = "entSensorThresholdTable"
    entsensorthresholdtable.EntityData.BundleName = "cisco_ios_xe"
    entsensorthresholdtable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entsensorthresholdtable.EntityData.SegmentPath = "entSensorThresholdTable"
    entsensorthresholdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entsensorthresholdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entsensorthresholdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entsensorthresholdtable.EntityData.Children = make(map[string]types.YChild)
    entsensorthresholdtable.EntityData.Children["entSensorThresholdEntry"] = types.YChild{"Entsensorthresholdentry", nil}
    for i := range entsensorthresholdtable.Entsensorthresholdentry {
        entsensorthresholdtable.EntityData.Children[types.GetSegmentPath(&entsensorthresholdtable.Entsensorthresholdentry[i])] = types.YChild{"Entsensorthresholdentry", &entsensorthresholdtable.Entsensorthresholdentry[i]}
    }
    entsensorthresholdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entsensorthresholdtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (entsensorthresholdentry *CISCOENTITYSENSORMIB_Entsensorthresholdtable_Entsensorthresholdentry) GetEntityData() *types.CommonEntityData {
    entsensorthresholdentry.EntityData.YFilter = entsensorthresholdentry.YFilter
    entsensorthresholdentry.EntityData.YangName = "entSensorThresholdEntry"
    entsensorthresholdentry.EntityData.BundleName = "cisco_ios_xe"
    entsensorthresholdentry.EntityData.ParentYangName = "entSensorThresholdTable"
    entsensorthresholdentry.EntityData.SegmentPath = "entSensorThresholdEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entsensorthresholdentry.Entphysicalindex) + "']" + "[entSensorThresholdIndex='" + fmt.Sprintf("%v", entsensorthresholdentry.Entsensorthresholdindex) + "']"
    entsensorthresholdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entsensorthresholdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entsensorthresholdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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

