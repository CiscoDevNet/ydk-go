// The MIB module to describe the status of the Environmental
// Monitor on those devices which support one.
package cisco_envmon_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_envmon_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ENVMON-MIB CISCO-ENVMON-MIB}", reflect.TypeOf(CISCOENVMONMIB{}))
    ydk.RegisterEntity("CISCO-ENVMON-MIB:CISCO-ENVMON-MIB", reflect.TypeOf(CISCOENVMONMIB{}))
}

// CiscoEnvMonState represents                    1000 C.
type CiscoEnvMonState string

const (
    CiscoEnvMonState_normal CiscoEnvMonState = "normal"

    CiscoEnvMonState_warning CiscoEnvMonState = "warning"

    CiscoEnvMonState_critical CiscoEnvMonState = "critical"

    CiscoEnvMonState_shutdown CiscoEnvMonState = "shutdown"

    CiscoEnvMonState_notPresent CiscoEnvMonState = "notPresent"

    CiscoEnvMonState_notFunctioning CiscoEnvMonState = "notFunctioning"
)

// CISCOENVMONMIB
type CISCOENVMONMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ciscoenvmonobjects CISCOENVMONMIB_Ciscoenvmonobjects

    
    Ciscoenvmonmibnotificationenables CISCOENVMONMIB_Ciscoenvmonmibnotificationenables

    // The table of voltage status maintained by the environmental monitor.
    Ciscoenvmonvoltagestatustable CISCOENVMONMIB_Ciscoenvmonvoltagestatustable

    // The table of ambient temperature status maintained by the environmental
    // monitor.
    Ciscoenvmontemperaturestatustable CISCOENVMONMIB_Ciscoenvmontemperaturestatustable

    // The table of fan status maintained by the environmental monitor.
    Ciscoenvmonfanstatustable CISCOENVMONMIB_Ciscoenvmonfanstatustable

    // The table of power supply status maintained by the environmental monitor
    // card.
    Ciscoenvmonsupplystatustable CISCOENVMONMIB_Ciscoenvmonsupplystatustable
}

func (cISCOENVMONMIB *CISCOENVMONMIB) GetEntityData() *types.CommonEntityData {
    cISCOENVMONMIB.EntityData.YFilter = cISCOENVMONMIB.YFilter
    cISCOENVMONMIB.EntityData.YangName = "CISCO-ENVMON-MIB"
    cISCOENVMONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENVMONMIB.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    cISCOENVMONMIB.EntityData.SegmentPath = "CISCO-ENVMON-MIB:CISCO-ENVMON-MIB"
    cISCOENVMONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENVMONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENVMONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENVMONMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENVMONMIB.EntityData.Children["ciscoEnvMonObjects"] = types.YChild{"Ciscoenvmonobjects", &cISCOENVMONMIB.Ciscoenvmonobjects}
    cISCOENVMONMIB.EntityData.Children["ciscoEnvMonMIBNotificationEnables"] = types.YChild{"Ciscoenvmonmibnotificationenables", &cISCOENVMONMIB.Ciscoenvmonmibnotificationenables}
    cISCOENVMONMIB.EntityData.Children["ciscoEnvMonVoltageStatusTable"] = types.YChild{"Ciscoenvmonvoltagestatustable", &cISCOENVMONMIB.Ciscoenvmonvoltagestatustable}
    cISCOENVMONMIB.EntityData.Children["ciscoEnvMonTemperatureStatusTable"] = types.YChild{"Ciscoenvmontemperaturestatustable", &cISCOENVMONMIB.Ciscoenvmontemperaturestatustable}
    cISCOENVMONMIB.EntityData.Children["ciscoEnvMonFanStatusTable"] = types.YChild{"Ciscoenvmonfanstatustable", &cISCOENVMONMIB.Ciscoenvmonfanstatustable}
    cISCOENVMONMIB.EntityData.Children["ciscoEnvMonSupplyStatusTable"] = types.YChild{"Ciscoenvmonsupplystatustable", &cISCOENVMONMIB.Ciscoenvmonsupplystatustable}
    cISCOENVMONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENVMONMIB.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonobjects
type CISCOENVMONMIB_Ciscoenvmonobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of environmental monitor located in the chassis. An oldAgs
    // environmental monitor card is identical to an ags environmental card except
    // that it is not capable of supplying data, and hence no instance of the
    // remaining objects in this MIB will be returned in response to an SNMP
    // query.  Note that only a firmware upgrade is required to convert an oldAgs
    // into an ags card. The type is Ciscoenvmonpresent.
    Ciscoenvmonpresent interface{}

    // Each bit is set to reflect the respective alarm being set.  The bit will be
    // cleared when the respective alarm is cleared. The type is map[string]bool.
    Ciscoenvmonalarmcontacts interface{}
}

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetEntityData() *types.CommonEntityData {
    ciscoenvmonobjects.EntityData.YFilter = ciscoenvmonobjects.YFilter
    ciscoenvmonobjects.EntityData.YangName = "ciscoEnvMonObjects"
    ciscoenvmonobjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonobjects.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoenvmonobjects.EntityData.SegmentPath = "ciscoEnvMonObjects"
    ciscoenvmonobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonobjects.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoenvmonobjects.EntityData.Leafs["ciscoEnvMonPresent"] = types.YLeaf{"Ciscoenvmonpresent", ciscoenvmonobjects.Ciscoenvmonpresent}
    ciscoenvmonobjects.EntityData.Leafs["ciscoEnvMonAlarmContacts"] = types.YLeaf{"Ciscoenvmonalarmcontacts", ciscoenvmonobjects.Ciscoenvmonalarmcontacts}
    return &(ciscoenvmonobjects.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent represents an ags card.
type CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent string

const (
    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_oldAgs CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "oldAgs"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_ags CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "ags"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_c7000 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "c7000"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_ci CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "ci"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_cAccessMon CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "cAccessMon"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_cat6000 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "cat6000"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_ubr7200 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "ubr7200"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_cat4000 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "cat4000"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_c10000 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "c10000"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_osr7600 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "osr7600"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_c7600 CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "c7600"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_c37xx CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "c37xx"

    CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent_other CISCOENVMONMIB_Ciscoenvmonobjects_Ciscoenvmonpresent = "other"
)

// CISCOENVMONMIB_Ciscoenvmonmibnotificationenables
type CISCOENVMONMIB_Ciscoenvmonmibnotificationenables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonShutdownNotification.  A false  value will prevent shutdown
    // notifications  from being generated by this system. The type is bool.
    Ciscoenvmonenableshutdownnotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonVoltageNotification. A false  value will prevent voltage
    // notifications from being  generated by this system. This object is
    // deprecated in favour of ciscoEnvMonEnableStatChangeNotif. The type is bool.
    Ciscoenvmonenablevoltagenotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonTemperatureNotification. A false value prevents temperature
    // notifications  from being sent by  this entity. This object is  deprecated
    // in favour of  ciscoEnvMonEnableStatChangeNotif. The type is bool.
    Ciscoenvmonenabletemperaturenotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonFanNotification. A false value prevents fan notifications  from
    // being sent by  this entity. This object is  deprecated in favour of 
    // ciscoEnvMonEnableStatChangeNotif. The type is bool.
    Ciscoenvmonenablefannotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonRedundantSupplyNotification.  A false value prevents redundant
    // supply notifications from being generated by this system. This object is
    // deprecated in favour of  ciscoEnvMonEnableStatChangeNotif. The type is
    // bool.
    Ciscoenvmonenableredundantsupplynotification interface{}

    // This variable indicates whether the system produces the
    // ciscoEnvMonVoltStatusChangeNotif, ciscoEnvMonTempStatusChangeNotif, 
    // ciscoEnvMonFanStatusChangeNotif and   ciscoEnvMonSuppStatusChangeNotif. A
    // false value will  prevent these notifications from being generated by  this
    // system. The type is bool.
    Ciscoenvmonenablestatchangenotif interface{}
}

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetEntityData() *types.CommonEntityData {
    ciscoenvmonmibnotificationenables.EntityData.YFilter = ciscoenvmonmibnotificationenables.YFilter
    ciscoenvmonmibnotificationenables.EntityData.YangName = "ciscoEnvMonMIBNotificationEnables"
    ciscoenvmonmibnotificationenables.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonmibnotificationenables.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoenvmonmibnotificationenables.EntityData.SegmentPath = "ciscoEnvMonMIBNotificationEnables"
    ciscoenvmonmibnotificationenables.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonmibnotificationenables.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonmibnotificationenables.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonmibnotificationenables.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonmibnotificationenables.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoenvmonmibnotificationenables.EntityData.Leafs["ciscoEnvMonEnableShutdownNotification"] = types.YLeaf{"Ciscoenvmonenableshutdownnotification", ciscoenvmonmibnotificationenables.Ciscoenvmonenableshutdownnotification}
    ciscoenvmonmibnotificationenables.EntityData.Leafs["ciscoEnvMonEnableVoltageNotification"] = types.YLeaf{"Ciscoenvmonenablevoltagenotification", ciscoenvmonmibnotificationenables.Ciscoenvmonenablevoltagenotification}
    ciscoenvmonmibnotificationenables.EntityData.Leafs["ciscoEnvMonEnableTemperatureNotification"] = types.YLeaf{"Ciscoenvmonenabletemperaturenotification", ciscoenvmonmibnotificationenables.Ciscoenvmonenabletemperaturenotification}
    ciscoenvmonmibnotificationenables.EntityData.Leafs["ciscoEnvMonEnableFanNotification"] = types.YLeaf{"Ciscoenvmonenablefannotification", ciscoenvmonmibnotificationenables.Ciscoenvmonenablefannotification}
    ciscoenvmonmibnotificationenables.EntityData.Leafs["ciscoEnvMonEnableRedundantSupplyNotification"] = types.YLeaf{"Ciscoenvmonenableredundantsupplynotification", ciscoenvmonmibnotificationenables.Ciscoenvmonenableredundantsupplynotification}
    ciscoenvmonmibnotificationenables.EntityData.Leafs["ciscoEnvMonEnableStatChangeNotif"] = types.YLeaf{"Ciscoenvmonenablestatchangenotif", ciscoenvmonmibnotificationenables.Ciscoenvmonenablestatchangenotif}
    return &(ciscoenvmonmibnotificationenables.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonvoltagestatustable
// The table of voltage status maintained by the environmental
// monitor.
type CISCOENVMONMIB_Ciscoenvmonvoltagestatustable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the voltage status table, representing the status of the
    // associated testpoint maintained by the environmental monitor. The type is
    // slice of
    // CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry.
    Ciscoenvmonvoltagestatusentry []CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetEntityData() *types.CommonEntityData {
    ciscoenvmonvoltagestatustable.EntityData.YFilter = ciscoenvmonvoltagestatustable.YFilter
    ciscoenvmonvoltagestatustable.EntityData.YangName = "ciscoEnvMonVoltageStatusTable"
    ciscoenvmonvoltagestatustable.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonvoltagestatustable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoenvmonvoltagestatustable.EntityData.SegmentPath = "ciscoEnvMonVoltageStatusTable"
    ciscoenvmonvoltagestatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonvoltagestatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonvoltagestatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonvoltagestatustable.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonvoltagestatustable.EntityData.Children["ciscoEnvMonVoltageStatusEntry"] = types.YChild{"Ciscoenvmonvoltagestatusentry", nil}
    for i := range ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry {
        ciscoenvmonvoltagestatustable.EntityData.Children[types.GetSegmentPath(&ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry[i])] = types.YChild{"Ciscoenvmonvoltagestatusentry", &ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry[i]}
    }
    ciscoenvmonvoltagestatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoenvmonvoltagestatustable.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry
// An entry in the voltage status table, representing the status
// of the associated testpoint maintained by the environmental
// monitor.
type CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the testpoint being instrumented.
    // This index is for SNMP purposes only, and has no intrinsic meaning. The
    // type is interface{} with range: 0..2147483647.
    Ciscoenvmonvoltagestatusindex interface{}

    // Textual description of the testpoint being instrumented. This description
    // is a short textual label, suitable as a human-sensible identification for
    // the rest of the information in the entry. The type is string with length:
    // 0..32.
    Ciscoenvmonvoltagestatusdescr interface{}

    // The current measurement of the testpoint being instrumented. The type is
    // interface{} with range: -2147483648..2147483647. Units are millivolts.
    Ciscoenvmonvoltagestatusvalue interface{}

    // The lowest value that the associated instance of the object
    // ciscoEnvMonVoltageStatusValue may obtain before an emergency shutdown of
    // the managed device is initiated. The type is interface{} with range:
    // -2147483648..2147483647. Units are millivolts.
    Ciscoenvmonvoltagethresholdlow interface{}

    // The highest value that the associated instance of the object
    // ciscoEnvMonVoltageStatusValue may obtain before an emergency shutdown of
    // the managed device is initiated. The type is interface{} with range:
    // -2147483648..2147483647. Units are millivolts.
    Ciscoenvmonvoltagethresholdhigh interface{}

    // The value of the associated instance of the object
    // ciscoEnvMonVoltageStatusValue at the time an emergency shutdown of the
    // managed device was last initiated.  This value is stored in non-volatile
    // RAM and hence is able to survive the shutdown. The type is interface{} with
    // range: -2147483648..2147483647. Units are millivolts.
    Ciscoenvmonvoltagelastshutdown interface{}

    // The current state of the testpoint being instrumented. The type is
    // CiscoEnvMonState.
    Ciscoenvmonvoltagestate interface{}
}

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetEntityData() *types.CommonEntityData {
    ciscoenvmonvoltagestatusentry.EntityData.YFilter = ciscoenvmonvoltagestatusentry.YFilter
    ciscoenvmonvoltagestatusentry.EntityData.YangName = "ciscoEnvMonVoltageStatusEntry"
    ciscoenvmonvoltagestatusentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonvoltagestatusentry.EntityData.ParentYangName = "ciscoEnvMonVoltageStatusTable"
    ciscoenvmonvoltagestatusentry.EntityData.SegmentPath = "ciscoEnvMonVoltageStatusEntry" + "[ciscoEnvMonVoltageStatusIndex='" + fmt.Sprintf("%v", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusindex) + "']"
    ciscoenvmonvoltagestatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonvoltagestatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonvoltagestatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonvoltagestatusentry.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonvoltagestatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageStatusIndex"] = types.YLeaf{"Ciscoenvmonvoltagestatusindex", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusindex}
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageStatusDescr"] = types.YLeaf{"Ciscoenvmonvoltagestatusdescr", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusdescr}
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageStatusValue"] = types.YLeaf{"Ciscoenvmonvoltagestatusvalue", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusvalue}
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageThresholdLow"] = types.YLeaf{"Ciscoenvmonvoltagethresholdlow", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagethresholdlow}
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageThresholdHigh"] = types.YLeaf{"Ciscoenvmonvoltagethresholdhigh", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagethresholdhigh}
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageLastShutdown"] = types.YLeaf{"Ciscoenvmonvoltagelastshutdown", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagelastshutdown}
    ciscoenvmonvoltagestatusentry.EntityData.Leafs["ciscoEnvMonVoltageState"] = types.YLeaf{"Ciscoenvmonvoltagestate", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestate}
    return &(ciscoenvmonvoltagestatusentry.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmontemperaturestatustable
// The table of ambient temperature status maintained by the
// environmental monitor.
type CISCOENVMONMIB_Ciscoenvmontemperaturestatustable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the ambient temperature status table, representing the status
    // of the associated testpoint maintained by the environmental monitor. The
    // type is slice of
    // CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry.
    Ciscoenvmontemperaturestatusentry []CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetEntityData() *types.CommonEntityData {
    ciscoenvmontemperaturestatustable.EntityData.YFilter = ciscoenvmontemperaturestatustable.YFilter
    ciscoenvmontemperaturestatustable.EntityData.YangName = "ciscoEnvMonTemperatureStatusTable"
    ciscoenvmontemperaturestatustable.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmontemperaturestatustable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoenvmontemperaturestatustable.EntityData.SegmentPath = "ciscoEnvMonTemperatureStatusTable"
    ciscoenvmontemperaturestatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmontemperaturestatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmontemperaturestatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmontemperaturestatustable.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmontemperaturestatustable.EntityData.Children["ciscoEnvMonTemperatureStatusEntry"] = types.YChild{"Ciscoenvmontemperaturestatusentry", nil}
    for i := range ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry {
        ciscoenvmontemperaturestatustable.EntityData.Children[types.GetSegmentPath(&ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry[i])] = types.YChild{"Ciscoenvmontemperaturestatusentry", &ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry[i]}
    }
    ciscoenvmontemperaturestatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoenvmontemperaturestatustable.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry
// An entry in the ambient temperature status table, representing
// the status of the associated testpoint maintained by the
// environmental monitor.
type CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the testpoint being instrumented.
    // This index is for SNMP purposes only, and has no intrinsic meaning. The
    // type is interface{} with range: 0..2147483647.
    Ciscoenvmontemperaturestatusindex interface{}

    // Textual description of the testpoint being instrumented. This description
    // is a short textual label, suitable as a human-sensible identification for
    // the rest of the information in the entry. The type is string with length:
    // 0..32.
    Ciscoenvmontemperaturestatusdescr interface{}

    // The current measurement of the testpoint being instrumented. The type is
    // interface{} with range: 0..4294967295. Units are degrees Celsius.
    Ciscoenvmontemperaturestatusvalue interface{}

    // The highest value that the associated instance of the object
    // ciscoEnvMonTemperatureStatusValue may obtain before an emergency shutdown
    // of the managed device is initiated. The type is interface{} with range:
    // -2147483648..2147483647. Units are degrees Celsius.
    Ciscoenvmontemperaturethreshold interface{}

    // The value of the associated instance of the object
    // ciscoEnvMonTemperatureStatusValue at the time an emergency shutdown of the
    // managed device was last initiated.  This value is stored in non-volatile
    // RAM and hence is able to survive the shutdown. The type is interface{} with
    // range: -2147483648..2147483647. Units are degrees Celsius.
    Ciscoenvmontemperaturelastshutdown interface{}

    // The current state of the testpoint being instrumented. The type is
    // CiscoEnvMonState.
    Ciscoenvmontemperaturestate interface{}
}

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetEntityData() *types.CommonEntityData {
    ciscoenvmontemperaturestatusentry.EntityData.YFilter = ciscoenvmontemperaturestatusentry.YFilter
    ciscoenvmontemperaturestatusentry.EntityData.YangName = "ciscoEnvMonTemperatureStatusEntry"
    ciscoenvmontemperaturestatusentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmontemperaturestatusentry.EntityData.ParentYangName = "ciscoEnvMonTemperatureStatusTable"
    ciscoenvmontemperaturestatusentry.EntityData.SegmentPath = "ciscoEnvMonTemperatureStatusEntry" + "[ciscoEnvMonTemperatureStatusIndex='" + fmt.Sprintf("%v", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusindex) + "']"
    ciscoenvmontemperaturestatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmontemperaturestatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmontemperaturestatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmontemperaturestatusentry.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmontemperaturestatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoenvmontemperaturestatusentry.EntityData.Leafs["ciscoEnvMonTemperatureStatusIndex"] = types.YLeaf{"Ciscoenvmontemperaturestatusindex", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusindex}
    ciscoenvmontemperaturestatusentry.EntityData.Leafs["ciscoEnvMonTemperatureStatusDescr"] = types.YLeaf{"Ciscoenvmontemperaturestatusdescr", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusdescr}
    ciscoenvmontemperaturestatusentry.EntityData.Leafs["ciscoEnvMonTemperatureStatusValue"] = types.YLeaf{"Ciscoenvmontemperaturestatusvalue", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusvalue}
    ciscoenvmontemperaturestatusentry.EntityData.Leafs["ciscoEnvMonTemperatureThreshold"] = types.YLeaf{"Ciscoenvmontemperaturethreshold", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturethreshold}
    ciscoenvmontemperaturestatusentry.EntityData.Leafs["ciscoEnvMonTemperatureLastShutdown"] = types.YLeaf{"Ciscoenvmontemperaturelastshutdown", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturelastshutdown}
    ciscoenvmontemperaturestatusentry.EntityData.Leafs["ciscoEnvMonTemperatureState"] = types.YLeaf{"Ciscoenvmontemperaturestate", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestate}
    return &(ciscoenvmontemperaturestatusentry.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonfanstatustable
// The table of fan status maintained by the environmental
// monitor.
type CISCOENVMONMIB_Ciscoenvmonfanstatustable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the fan status table, representing the status of the associated
    // fan maintained by the environmental monitor. The type is slice of
    // CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry.
    Ciscoenvmonfanstatusentry []CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetEntityData() *types.CommonEntityData {
    ciscoenvmonfanstatustable.EntityData.YFilter = ciscoenvmonfanstatustable.YFilter
    ciscoenvmonfanstatustable.EntityData.YangName = "ciscoEnvMonFanStatusTable"
    ciscoenvmonfanstatustable.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonfanstatustable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoenvmonfanstatustable.EntityData.SegmentPath = "ciscoEnvMonFanStatusTable"
    ciscoenvmonfanstatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonfanstatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonfanstatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonfanstatustable.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonfanstatustable.EntityData.Children["ciscoEnvMonFanStatusEntry"] = types.YChild{"Ciscoenvmonfanstatusentry", nil}
    for i := range ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry {
        ciscoenvmonfanstatustable.EntityData.Children[types.GetSegmentPath(&ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry[i])] = types.YChild{"Ciscoenvmonfanstatusentry", &ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry[i]}
    }
    ciscoenvmonfanstatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoenvmonfanstatustable.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry
// An entry in the fan status table, representing the status of
// the associated fan maintained by the environmental monitor.
type CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the fan being instrumented. This
    // index is for SNMP purposes only, and has no intrinsic meaning. The type is
    // interface{} with range: 0..2147483647.
    Ciscoenvmonfanstatusindex interface{}

    // Textual description of the fan being instrumented. This description is a
    // short textual label, suitable as a human-sensible identification for the
    // rest of the information in the entry. The type is string with length:
    // 0..32.
    Ciscoenvmonfanstatusdescr interface{}

    // The current state of the fan being instrumented. The type is
    // CiscoEnvMonState.
    Ciscoenvmonfanstate interface{}
}

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetEntityData() *types.CommonEntityData {
    ciscoenvmonfanstatusentry.EntityData.YFilter = ciscoenvmonfanstatusentry.YFilter
    ciscoenvmonfanstatusentry.EntityData.YangName = "ciscoEnvMonFanStatusEntry"
    ciscoenvmonfanstatusentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonfanstatusentry.EntityData.ParentYangName = "ciscoEnvMonFanStatusTable"
    ciscoenvmonfanstatusentry.EntityData.SegmentPath = "ciscoEnvMonFanStatusEntry" + "[ciscoEnvMonFanStatusIndex='" + fmt.Sprintf("%v", ciscoenvmonfanstatusentry.Ciscoenvmonfanstatusindex) + "']"
    ciscoenvmonfanstatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonfanstatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonfanstatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonfanstatusentry.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonfanstatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoenvmonfanstatusentry.EntityData.Leafs["ciscoEnvMonFanStatusIndex"] = types.YLeaf{"Ciscoenvmonfanstatusindex", ciscoenvmonfanstatusentry.Ciscoenvmonfanstatusindex}
    ciscoenvmonfanstatusentry.EntityData.Leafs["ciscoEnvMonFanStatusDescr"] = types.YLeaf{"Ciscoenvmonfanstatusdescr", ciscoenvmonfanstatusentry.Ciscoenvmonfanstatusdescr}
    ciscoenvmonfanstatusentry.EntityData.Leafs["ciscoEnvMonFanState"] = types.YLeaf{"Ciscoenvmonfanstate", ciscoenvmonfanstatusentry.Ciscoenvmonfanstate}
    return &(ciscoenvmonfanstatusentry.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonsupplystatustable
// The table of power supply status maintained by the
// environmental monitor card.
type CISCOENVMONMIB_Ciscoenvmonsupplystatustable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the power supply status table, representing the status of the
    // associated power supply maintained by the environmental monitor card. The
    // type is slice of
    // CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry.
    Ciscoenvmonsupplystatusentry []CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetEntityData() *types.CommonEntityData {
    ciscoenvmonsupplystatustable.EntityData.YFilter = ciscoenvmonsupplystatustable.YFilter
    ciscoenvmonsupplystatustable.EntityData.YangName = "ciscoEnvMonSupplyStatusTable"
    ciscoenvmonsupplystatustable.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonsupplystatustable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoenvmonsupplystatustable.EntityData.SegmentPath = "ciscoEnvMonSupplyStatusTable"
    ciscoenvmonsupplystatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonsupplystatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonsupplystatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonsupplystatustable.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonsupplystatustable.EntityData.Children["ciscoEnvMonSupplyStatusEntry"] = types.YChild{"Ciscoenvmonsupplystatusentry", nil}
    for i := range ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry {
        ciscoenvmonsupplystatustable.EntityData.Children[types.GetSegmentPath(&ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry[i])] = types.YChild{"Ciscoenvmonsupplystatusentry", &ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry[i]}
    }
    ciscoenvmonsupplystatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoenvmonsupplystatustable.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry
// An entry in the power supply status table, representing the
// status of the associated power supply maintained by the
// environmental monitor card.
type CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the power supply being
    // instrumented. This index is for SNMP purposes only, and has no intrinsic
    // meaning. The type is interface{} with range: 0..2147483647.
    Ciscoenvmonsupplystatusindex interface{}

    // Textual description of the power supply being instrumented. This
    // description is a short textual label, suitable as a human-sensible
    // identification for the rest of the information in the entry. The type is
    // string with length: 0..64.
    Ciscoenvmonsupplystatusdescr interface{}

    // The current state of the power supply being instrumented. The type is
    // CiscoEnvMonState.
    Ciscoenvmonsupplystate interface{}

    // The power supply source. unknown - Power supply source unknown ac      - AC
    // power supply dc      - DC power supply externalPowerSupply - External power
    // supply internalRedundant - Internal redundant power supply . The type is
    // Ciscoenvmonsupplysource.
    Ciscoenvmonsupplysource interface{}
}

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetEntityData() *types.CommonEntityData {
    ciscoenvmonsupplystatusentry.EntityData.YFilter = ciscoenvmonsupplystatusentry.YFilter
    ciscoenvmonsupplystatusentry.EntityData.YangName = "ciscoEnvMonSupplyStatusEntry"
    ciscoenvmonsupplystatusentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoenvmonsupplystatusentry.EntityData.ParentYangName = "ciscoEnvMonSupplyStatusTable"
    ciscoenvmonsupplystatusentry.EntityData.SegmentPath = "ciscoEnvMonSupplyStatusEntry" + "[ciscoEnvMonSupplyStatusIndex='" + fmt.Sprintf("%v", ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystatusindex) + "']"
    ciscoenvmonsupplystatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoenvmonsupplystatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoenvmonsupplystatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoenvmonsupplystatusentry.EntityData.Children = make(map[string]types.YChild)
    ciscoenvmonsupplystatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoenvmonsupplystatusentry.EntityData.Leafs["ciscoEnvMonSupplyStatusIndex"] = types.YLeaf{"Ciscoenvmonsupplystatusindex", ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystatusindex}
    ciscoenvmonsupplystatusentry.EntityData.Leafs["ciscoEnvMonSupplyStatusDescr"] = types.YLeaf{"Ciscoenvmonsupplystatusdescr", ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystatusdescr}
    ciscoenvmonsupplystatusentry.EntityData.Leafs["ciscoEnvMonSupplyState"] = types.YLeaf{"Ciscoenvmonsupplystate", ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystate}
    ciscoenvmonsupplystatusentry.EntityData.Leafs["ciscoEnvMonSupplySource"] = types.YLeaf{"Ciscoenvmonsupplysource", ciscoenvmonsupplystatusentry.Ciscoenvmonsupplysource}
    return &(ciscoenvmonsupplystatusentry.EntityData)
}

// CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource represents internalRedundant - Internal redundant power supply 
type CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource string

const (
    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_unknown CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "unknown"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_ac CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "ac"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_dc CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "dc"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_externalPowerSupply CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "externalPowerSupply"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_internalRedundant CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "internalRedundant"
)

