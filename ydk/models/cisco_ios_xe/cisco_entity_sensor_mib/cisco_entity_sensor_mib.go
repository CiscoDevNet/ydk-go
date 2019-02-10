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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    EntSensorGlobalObjects CISCOENTITYSENSORMIB_EntSensorGlobalObjects

    // This table lists the type, scale, and present value of a sensor listed in
    // the Entity-MIB entPhysicalTable.
    EntSensorValueTable CISCOENTITYSENSORMIB_EntSensorValueTable

    // This table lists the threshold severity, relation, and comparison value,
    // for a sensor listed in the Entity-MIB  entPhysicalTable.
    EntSensorThresholdTable CISCOENTITYSENSORMIB_EntSensorThresholdTable
}

func (cISCOENTITYSENSORMIB *CISCOENTITYSENSORMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYSENSORMIB.EntityData.YFilter = cISCOENTITYSENSORMIB.YFilter
    cISCOENTITYSENSORMIB.EntityData.YangName = "CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENTITYSENSORMIB.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.SegmentPath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB"
    cISCOENTITYSENSORMIB.EntityData.AbsolutePath = cISCOENTITYSENSORMIB.EntityData.SegmentPath
    cISCOENTITYSENSORMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENTITYSENSORMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENTITYSENSORMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENTITYSENSORMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENTITYSENSORMIB.EntityData.Children.Append("entSensorGlobalObjects", types.YChild{"EntSensorGlobalObjects", &cISCOENTITYSENSORMIB.EntSensorGlobalObjects})
    cISCOENTITYSENSORMIB.EntityData.Children.Append("entSensorValueTable", types.YChild{"EntSensorValueTable", &cISCOENTITYSENSORMIB.EntSensorValueTable})
    cISCOENTITYSENSORMIB.EntityData.Children.Append("entSensorThresholdTable", types.YChild{"EntSensorThresholdTable", &cISCOENTITYSENSORMIB.EntSensorThresholdTable})
    cISCOENTITYSENSORMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENTITYSENSORMIB.EntityData.YListKeys = []string {}

    return &(cISCOENTITYSENSORMIB.EntityData)
}

// CISCOENTITYSENSORMIB_EntSensorGlobalObjects
type CISCOENTITYSENSORMIB_EntSensorGlobalObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable enables the generation of entSensorThresholdNotification
    // globally on the device. If this object value is 'false', then no
    // entSensorThresholdNotification will be generated on this device. If this
    // object value is 'true', then whether a  entSensorThresholdNotification for
    // a threshold will be generated or not depends on the instance value of
    // entSensorThresholdNotificationEnable for that threshold. The type is bool.
    EntSensorThreshNotifGlobalEnable interface{}
}

func (entSensorGlobalObjects *CISCOENTITYSENSORMIB_EntSensorGlobalObjects) GetEntityData() *types.CommonEntityData {
    entSensorGlobalObjects.EntityData.YFilter = entSensorGlobalObjects.YFilter
    entSensorGlobalObjects.EntityData.YangName = "entSensorGlobalObjects"
    entSensorGlobalObjects.EntityData.BundleName = "cisco_ios_xe"
    entSensorGlobalObjects.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entSensorGlobalObjects.EntityData.SegmentPath = "entSensorGlobalObjects"
    entSensorGlobalObjects.EntityData.AbsolutePath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB/" + entSensorGlobalObjects.EntityData.SegmentPath
    entSensorGlobalObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entSensorGlobalObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entSensorGlobalObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entSensorGlobalObjects.EntityData.Children = types.NewOrderedMap()
    entSensorGlobalObjects.EntityData.Leafs = types.NewOrderedMap()
    entSensorGlobalObjects.EntityData.Leafs.Append("entSensorThreshNotifGlobalEnable", types.YLeaf{"EntSensorThreshNotifGlobalEnable", entSensorGlobalObjects.EntSensorThreshNotifGlobalEnable})

    entSensorGlobalObjects.EntityData.YListKeys = []string {}

    return &(entSensorGlobalObjects.EntityData)
}

// CISCOENTITYSENSORMIB_EntSensorValueTable
// This table lists the type, scale, and present value
// of a sensor listed in the Entity-MIB entPhysicalTable.
type CISCOENTITYSENSORMIB_EntSensorValueTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entSensorValueTable entry describes the present reading of a sensor, the
    // measurement units and scale, and sensor operational status. The type is
    // slice of CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry.
    EntSensorValueEntry []*CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry
}

func (entSensorValueTable *CISCOENTITYSENSORMIB_EntSensorValueTable) GetEntityData() *types.CommonEntityData {
    entSensorValueTable.EntityData.YFilter = entSensorValueTable.YFilter
    entSensorValueTable.EntityData.YangName = "entSensorValueTable"
    entSensorValueTable.EntityData.BundleName = "cisco_ios_xe"
    entSensorValueTable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entSensorValueTable.EntityData.SegmentPath = "entSensorValueTable"
    entSensorValueTable.EntityData.AbsolutePath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB/" + entSensorValueTable.EntityData.SegmentPath
    entSensorValueTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entSensorValueTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entSensorValueTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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
// An entSensorValueTable entry describes the
// present reading of a sensor, the measurement units
// and scale, and sensor operational status.
type CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This variable indicates the type of data reported by the entSensorValue. 
    // This variable is set by the agent at start-up and the value does not change
    // during operation. The type is SensorDataType.
    EntSensorType interface{}

    // This variable indicates the exponent to apply to sensor values reported by
    // entSensorValue.  This variable is set by the agent at start-up and the
    // value does not change during operation. The type is SensorDataScale.
    EntSensorScale interface{}

    // This variable indicates the number of decimal places of precision in
    // fixed-point sensor values reported by entSensorValue.  This variable is set
    // to 0 when entSensorType is not a fixed-point type: e.g.'percentRH(9)', 
    // 'rpm(10)', 'cmm(11)', or 'truthvalue(12)'.  This variable is set by the
    // agent at start-up and the value does not change during operation. The type
    // is interface{} with range: -8..9.
    EntSensorPrecision interface{}

    // This variable reports the most recent measurement seen by the sensor.  To
    // correctly display or interpret this variable's value,  you must also know
    // entSensorType, entSensorScale, and  entSensorPrecision.  However, you can
    // compare entSensorValue with the threshold values given in
    // entSensorThresholdTable without any semantic knowledge. The type is
    // interface{} with range: -1000000000..1073741823.
    EntSensorValue interface{}

    // This variable indicates the present operational status of the sensor. The
    // type is SensorStatus.
    EntSensorStatus interface{}

    // This variable indicates the age of the value reported by entSensorValue.
    // The type is interface{} with range: 0..4294967295.
    EntSensorValueTimeStamp interface{}

    // This variable indicates the rate that the agent updates entSensorValue. The
    // type is interface{} with range: 0..999999999. Units are seconds.
    EntSensorValueUpdateRate interface{}

    // This object identifies the physical entity for which the sensor is taking
    // measurements.  For example, for a sensor measuring the voltage output of a
    // power-supply, this object would be the entPhysicalIndex of that
    // power-supply; for a sensor measuring the temperature inside one chassis of
    // a multi-chassis system, this object would be the enPhysicalIndex of that
    // chassis.  This object has a value of zero when the physical entity for
    // which the sensor is taking measurements can not be represented by any one
    // row in the entPhysicalTable, or that there is no such physical entity. The
    // type is interface{} with range: 0..2147483647.
    EntSensorMeasuredEntity interface{}
}

func (entSensorValueEntry *CISCOENTITYSENSORMIB_EntSensorValueTable_EntSensorValueEntry) GetEntityData() *types.CommonEntityData {
    entSensorValueEntry.EntityData.YFilter = entSensorValueEntry.YFilter
    entSensorValueEntry.EntityData.YangName = "entSensorValueEntry"
    entSensorValueEntry.EntityData.BundleName = "cisco_ios_xe"
    entSensorValueEntry.EntityData.ParentYangName = "entSensorValueTable"
    entSensorValueEntry.EntityData.SegmentPath = "entSensorValueEntry" + types.AddKeyToken(entSensorValueEntry.EntPhysicalIndex, "entPhysicalIndex")
    entSensorValueEntry.EntityData.AbsolutePath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB/entSensorValueTable/" + entSensorValueEntry.EntityData.SegmentPath
    entSensorValueEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entSensorValueEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entSensorValueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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
// This table lists the threshold severity, relation, and
// comparison value, for a sensor listed in the Entity-MIB 
// entPhysicalTable.
type CISCOENTITYSENSORMIB_EntSensorThresholdTable struct {
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
    // CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry.
    EntSensorThresholdEntry []*CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry
}

func (entSensorThresholdTable *CISCOENTITYSENSORMIB_EntSensorThresholdTable) GetEntityData() *types.CommonEntityData {
    entSensorThresholdTable.EntityData.YFilter = entSensorThresholdTable.YFilter
    entSensorThresholdTable.EntityData.YangName = "entSensorThresholdTable"
    entSensorThresholdTable.EntityData.BundleName = "cisco_ios_xe"
    entSensorThresholdTable.EntityData.ParentYangName = "CISCO-ENTITY-SENSOR-MIB"
    entSensorThresholdTable.EntityData.SegmentPath = "entSensorThresholdTable"
    entSensorThresholdTable.EntityData.AbsolutePath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB/" + entSensorThresholdTable.EntityData.SegmentPath
    entSensorThresholdTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entSensorThresholdTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entSensorThresholdTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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
type CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // entSensorThresholdTable. This index permits the same sensor to have several
    // different thresholds. The type is interface{} with range: 1..99999999.
    EntSensorThresholdIndex interface{}

    // This variable indicates the severity of this threshold. The type is
    // SensorThresholdSeverity.
    EntSensorThresholdSeverity interface{}

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
    EntSensorThresholdRelation interface{}

    // This variable indicates the value of the threshold.  To correctly display
    // or interpret this variable's value,  you must also know entSensorType,
    // entSensorScale, and  entSensorPrecision.  However, you can directly compare
    // entSensorValue  with the threshold values given in entSensorThresholdTable 
    // without any semantic knowledge. The type is interface{} with range:
    // -1000000000..1073741823.
    EntSensorThresholdValue interface{}

    // This variable indicates the result of the most recent evaluation of the
    // threshold.  If the threshold condition is true,
    // entSensorThresholdEvaluation  is true(1).  If the threshold condition is
    // false,  entSensorThresholdEvaluation is false(2).  Thresholds are evaluated
    // at the rate indicated by  entSensorValueUpdateRate. The type is bool.
    EntSensorThresholdEvaluation interface{}

    // This variable controls generation of entSensorThresholdNotification for
    // this threshold.  When this variable is 'true', generation of 
    // entSensorThresholdNotification is enabled for this threshold. When this
    // variable is 'false',  generation of entSensorThresholdNotification is
    // disabled for this threshold. The type is bool.
    EntSensorThresholdNotificationEnable interface{}
}

func (entSensorThresholdEntry *CISCOENTITYSENSORMIB_EntSensorThresholdTable_EntSensorThresholdEntry) GetEntityData() *types.CommonEntityData {
    entSensorThresholdEntry.EntityData.YFilter = entSensorThresholdEntry.YFilter
    entSensorThresholdEntry.EntityData.YangName = "entSensorThresholdEntry"
    entSensorThresholdEntry.EntityData.BundleName = "cisco_ios_xe"
    entSensorThresholdEntry.EntityData.ParentYangName = "entSensorThresholdTable"
    entSensorThresholdEntry.EntityData.SegmentPath = "entSensorThresholdEntry" + types.AddKeyToken(entSensorThresholdEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(entSensorThresholdEntry.EntSensorThresholdIndex, "entSensorThresholdIndex")
    entSensorThresholdEntry.EntityData.AbsolutePath = "CISCO-ENTITY-SENSOR-MIB:CISCO-ENTITY-SENSOR-MIB/entSensorThresholdTable/" + entSensorThresholdEntry.EntityData.SegmentPath
    entSensorThresholdEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entSensorThresholdEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entSensorThresholdEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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

