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
    parent types.Entity
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

func (cISCOENVMONMIB *CISCOENVMONMIB) GetFilter() yfilter.YFilter { return cISCOENVMONMIB.YFilter }

func (cISCOENVMONMIB *CISCOENVMONMIB) SetFilter(yf yfilter.YFilter) { cISCOENVMONMIB.YFilter = yf }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetGoName(yname string) string {
    if yname == "ciscoEnvMonObjects" { return "Ciscoenvmonobjects" }
    if yname == "ciscoEnvMonMIBNotificationEnables" { return "Ciscoenvmonmibnotificationenables" }
    if yname == "ciscoEnvMonVoltageStatusTable" { return "Ciscoenvmonvoltagestatustable" }
    if yname == "ciscoEnvMonTemperatureStatusTable" { return "Ciscoenvmontemperaturestatustable" }
    if yname == "ciscoEnvMonFanStatusTable" { return "Ciscoenvmonfanstatustable" }
    if yname == "ciscoEnvMonSupplyStatusTable" { return "Ciscoenvmonsupplystatustable" }
    return ""
}

func (cISCOENVMONMIB *CISCOENVMONMIB) GetSegmentPath() string {
    return "CISCO-ENVMON-MIB:CISCO-ENVMON-MIB"
}

func (cISCOENVMONMIB *CISCOENVMONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoEnvMonObjects" {
        return &cISCOENVMONMIB.Ciscoenvmonobjects
    }
    if childYangName == "ciscoEnvMonMIBNotificationEnables" {
        return &cISCOENVMONMIB.Ciscoenvmonmibnotificationenables
    }
    if childYangName == "ciscoEnvMonVoltageStatusTable" {
        return &cISCOENVMONMIB.Ciscoenvmonvoltagestatustable
    }
    if childYangName == "ciscoEnvMonTemperatureStatusTable" {
        return &cISCOENVMONMIB.Ciscoenvmontemperaturestatustable
    }
    if childYangName == "ciscoEnvMonFanStatusTable" {
        return &cISCOENVMONMIB.Ciscoenvmonfanstatustable
    }
    if childYangName == "ciscoEnvMonSupplyStatusTable" {
        return &cISCOENVMONMIB.Ciscoenvmonsupplystatustable
    }
    return nil
}

func (cISCOENVMONMIB *CISCOENVMONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoEnvMonObjects"] = &cISCOENVMONMIB.Ciscoenvmonobjects
    children["ciscoEnvMonMIBNotificationEnables"] = &cISCOENVMONMIB.Ciscoenvmonmibnotificationenables
    children["ciscoEnvMonVoltageStatusTable"] = &cISCOENVMONMIB.Ciscoenvmonvoltagestatustable
    children["ciscoEnvMonTemperatureStatusTable"] = &cISCOENVMONMIB.Ciscoenvmontemperaturestatustable
    children["ciscoEnvMonFanStatusTable"] = &cISCOENVMONMIB.Ciscoenvmonfanstatustable
    children["ciscoEnvMonSupplyStatusTable"] = &cISCOENVMONMIB.Ciscoenvmonsupplystatustable
    return children
}

func (cISCOENVMONMIB *CISCOENVMONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOENVMONMIB *CISCOENVMONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetYangName() string { return "CISCO-ENVMON-MIB" }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOENVMONMIB *CISCOENVMONMIB) SetParent(parent types.Entity) { cISCOENVMONMIB.parent = parent }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetParent() types.Entity { return cISCOENVMONMIB.parent }

func (cISCOENVMONMIB *CISCOENVMONMIB) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

// CISCOENVMONMIB_Ciscoenvmonobjects
type CISCOENVMONMIB_Ciscoenvmonobjects struct {
    parent types.Entity
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

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetFilter() yfilter.YFilter { return ciscoenvmonobjects.YFilter }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) SetFilter(yf yfilter.YFilter) { ciscoenvmonobjects.YFilter = yf }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetGoName(yname string) string {
    if yname == "ciscoEnvMonPresent" { return "Ciscoenvmonpresent" }
    if yname == "ciscoEnvMonAlarmContacts" { return "Ciscoenvmonalarmcontacts" }
    return ""
}

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetSegmentPath() string {
    return "ciscoEnvMonObjects"
}

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoEnvMonPresent"] = ciscoenvmonobjects.Ciscoenvmonpresent
    leafs["ciscoEnvMonAlarmContacts"] = ciscoenvmonobjects.Ciscoenvmonalarmcontacts
    return leafs
}

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetYangName() string { return "ciscoEnvMonObjects" }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) SetParent(parent types.Entity) { ciscoenvmonobjects.parent = parent }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetParent() types.Entity { return ciscoenvmonobjects.parent }

func (ciscoenvmonobjects *CISCOENVMONMIB_Ciscoenvmonobjects) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

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
    parent types.Entity
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

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetFilter() yfilter.YFilter { return ciscoenvmonmibnotificationenables.YFilter }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) SetFilter(yf yfilter.YFilter) { ciscoenvmonmibnotificationenables.YFilter = yf }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetGoName(yname string) string {
    if yname == "ciscoEnvMonEnableShutdownNotification" { return "Ciscoenvmonenableshutdownnotification" }
    if yname == "ciscoEnvMonEnableVoltageNotification" { return "Ciscoenvmonenablevoltagenotification" }
    if yname == "ciscoEnvMonEnableTemperatureNotification" { return "Ciscoenvmonenabletemperaturenotification" }
    if yname == "ciscoEnvMonEnableFanNotification" { return "Ciscoenvmonenablefannotification" }
    if yname == "ciscoEnvMonEnableRedundantSupplyNotification" { return "Ciscoenvmonenableredundantsupplynotification" }
    if yname == "ciscoEnvMonEnableStatChangeNotif" { return "Ciscoenvmonenablestatchangenotif" }
    return ""
}

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetSegmentPath() string {
    return "ciscoEnvMonMIBNotificationEnables"
}

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoEnvMonEnableShutdownNotification"] = ciscoenvmonmibnotificationenables.Ciscoenvmonenableshutdownnotification
    leafs["ciscoEnvMonEnableVoltageNotification"] = ciscoenvmonmibnotificationenables.Ciscoenvmonenablevoltagenotification
    leafs["ciscoEnvMonEnableTemperatureNotification"] = ciscoenvmonmibnotificationenables.Ciscoenvmonenabletemperaturenotification
    leafs["ciscoEnvMonEnableFanNotification"] = ciscoenvmonmibnotificationenables.Ciscoenvmonenablefannotification
    leafs["ciscoEnvMonEnableRedundantSupplyNotification"] = ciscoenvmonmibnotificationenables.Ciscoenvmonenableredundantsupplynotification
    leafs["ciscoEnvMonEnableStatChangeNotif"] = ciscoenvmonmibnotificationenables.Ciscoenvmonenablestatchangenotif
    return leafs
}

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetYangName() string { return "ciscoEnvMonMIBNotificationEnables" }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) SetParent(parent types.Entity) { ciscoenvmonmibnotificationenables.parent = parent }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetParent() types.Entity { return ciscoenvmonmibnotificationenables.parent }

func (ciscoenvmonmibnotificationenables *CISCOENVMONMIB_Ciscoenvmonmibnotificationenables) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

// CISCOENVMONMIB_Ciscoenvmonvoltagestatustable
// The table of voltage status maintained by the environmental
// monitor.
type CISCOENVMONMIB_Ciscoenvmonvoltagestatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the voltage status table, representing the status of the
    // associated testpoint maintained by the environmental monitor. The type is
    // slice of
    // CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry.
    Ciscoenvmonvoltagestatusentry []CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetFilter() yfilter.YFilter { return ciscoenvmonvoltagestatustable.YFilter }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) SetFilter(yf yfilter.YFilter) { ciscoenvmonvoltagestatustable.YFilter = yf }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetGoName(yname string) string {
    if yname == "ciscoEnvMonVoltageStatusEntry" { return "Ciscoenvmonvoltagestatusentry" }
    return ""
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetSegmentPath() string {
    return "ciscoEnvMonVoltageStatusTable"
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoEnvMonVoltageStatusEntry" {
        for _, c := range ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry {
            if ciscoenvmonvoltagestatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry{}
        ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry = append(ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry, child)
        return &ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry[len(ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry)-1]
    }
    return nil
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry {
        children[ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry[i].GetSegmentPath()] = &ciscoenvmonvoltagestatustable.Ciscoenvmonvoltagestatusentry[i]
    }
    return children
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetYangName() string { return "ciscoEnvMonVoltageStatusTable" }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) SetParent(parent types.Entity) { ciscoenvmonvoltagestatustable.parent = parent }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetParent() types.Entity { return ciscoenvmonvoltagestatustable.parent }

func (ciscoenvmonvoltagestatustable *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

// CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry
// An entry in the voltage status table, representing the status
// of the associated testpoint maintained by the environmental
// monitor.
type CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry struct {
    parent types.Entity
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

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetFilter() yfilter.YFilter { return ciscoenvmonvoltagestatusentry.YFilter }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) SetFilter(yf yfilter.YFilter) { ciscoenvmonvoltagestatusentry.YFilter = yf }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetGoName(yname string) string {
    if yname == "ciscoEnvMonVoltageStatusIndex" { return "Ciscoenvmonvoltagestatusindex" }
    if yname == "ciscoEnvMonVoltageStatusDescr" { return "Ciscoenvmonvoltagestatusdescr" }
    if yname == "ciscoEnvMonVoltageStatusValue" { return "Ciscoenvmonvoltagestatusvalue" }
    if yname == "ciscoEnvMonVoltageThresholdLow" { return "Ciscoenvmonvoltagethresholdlow" }
    if yname == "ciscoEnvMonVoltageThresholdHigh" { return "Ciscoenvmonvoltagethresholdhigh" }
    if yname == "ciscoEnvMonVoltageLastShutdown" { return "Ciscoenvmonvoltagelastshutdown" }
    if yname == "ciscoEnvMonVoltageState" { return "Ciscoenvmonvoltagestate" }
    return ""
}

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetSegmentPath() string {
    return "ciscoEnvMonVoltageStatusEntry" + "[ciscoEnvMonVoltageStatusIndex='" + fmt.Sprintf("%v", ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusindex) + "']"
}

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoEnvMonVoltageStatusIndex"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusindex
    leafs["ciscoEnvMonVoltageStatusDescr"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusdescr
    leafs["ciscoEnvMonVoltageStatusValue"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestatusvalue
    leafs["ciscoEnvMonVoltageThresholdLow"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagethresholdlow
    leafs["ciscoEnvMonVoltageThresholdHigh"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagethresholdhigh
    leafs["ciscoEnvMonVoltageLastShutdown"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagelastshutdown
    leafs["ciscoEnvMonVoltageState"] = ciscoenvmonvoltagestatusentry.Ciscoenvmonvoltagestate
    return leafs
}

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetYangName() string { return "ciscoEnvMonVoltageStatusEntry" }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) SetParent(parent types.Entity) { ciscoenvmonvoltagestatusentry.parent = parent }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetParent() types.Entity { return ciscoenvmonvoltagestatusentry.parent }

func (ciscoenvmonvoltagestatusentry *CISCOENVMONMIB_Ciscoenvmonvoltagestatustable_Ciscoenvmonvoltagestatusentry) GetParentYangName() string { return "ciscoEnvMonVoltageStatusTable" }

// CISCOENVMONMIB_Ciscoenvmontemperaturestatustable
// The table of ambient temperature status maintained by the
// environmental monitor.
type CISCOENVMONMIB_Ciscoenvmontemperaturestatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the ambient temperature status table, representing the status
    // of the associated testpoint maintained by the environmental monitor. The
    // type is slice of
    // CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry.
    Ciscoenvmontemperaturestatusentry []CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetFilter() yfilter.YFilter { return ciscoenvmontemperaturestatustable.YFilter }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) SetFilter(yf yfilter.YFilter) { ciscoenvmontemperaturestatustable.YFilter = yf }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetGoName(yname string) string {
    if yname == "ciscoEnvMonTemperatureStatusEntry" { return "Ciscoenvmontemperaturestatusentry" }
    return ""
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetSegmentPath() string {
    return "ciscoEnvMonTemperatureStatusTable"
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoEnvMonTemperatureStatusEntry" {
        for _, c := range ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry {
            if ciscoenvmontemperaturestatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry{}
        ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry = append(ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry, child)
        return &ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry[len(ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry)-1]
    }
    return nil
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry {
        children[ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry[i].GetSegmentPath()] = &ciscoenvmontemperaturestatustable.Ciscoenvmontemperaturestatusentry[i]
    }
    return children
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetYangName() string { return "ciscoEnvMonTemperatureStatusTable" }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) SetParent(parent types.Entity) { ciscoenvmontemperaturestatustable.parent = parent }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetParent() types.Entity { return ciscoenvmontemperaturestatustable.parent }

func (ciscoenvmontemperaturestatustable *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

// CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry
// An entry in the ambient temperature status table, representing
// the status of the associated testpoint maintained by the
// environmental monitor.
type CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry struct {
    parent types.Entity
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

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetFilter() yfilter.YFilter { return ciscoenvmontemperaturestatusentry.YFilter }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) SetFilter(yf yfilter.YFilter) { ciscoenvmontemperaturestatusentry.YFilter = yf }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetGoName(yname string) string {
    if yname == "ciscoEnvMonTemperatureStatusIndex" { return "Ciscoenvmontemperaturestatusindex" }
    if yname == "ciscoEnvMonTemperatureStatusDescr" { return "Ciscoenvmontemperaturestatusdescr" }
    if yname == "ciscoEnvMonTemperatureStatusValue" { return "Ciscoenvmontemperaturestatusvalue" }
    if yname == "ciscoEnvMonTemperatureThreshold" { return "Ciscoenvmontemperaturethreshold" }
    if yname == "ciscoEnvMonTemperatureLastShutdown" { return "Ciscoenvmontemperaturelastshutdown" }
    if yname == "ciscoEnvMonTemperatureState" { return "Ciscoenvmontemperaturestate" }
    return ""
}

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetSegmentPath() string {
    return "ciscoEnvMonTemperatureStatusEntry" + "[ciscoEnvMonTemperatureStatusIndex='" + fmt.Sprintf("%v", ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusindex) + "']"
}

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoEnvMonTemperatureStatusIndex"] = ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusindex
    leafs["ciscoEnvMonTemperatureStatusDescr"] = ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusdescr
    leafs["ciscoEnvMonTemperatureStatusValue"] = ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestatusvalue
    leafs["ciscoEnvMonTemperatureThreshold"] = ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturethreshold
    leafs["ciscoEnvMonTemperatureLastShutdown"] = ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturelastshutdown
    leafs["ciscoEnvMonTemperatureState"] = ciscoenvmontemperaturestatusentry.Ciscoenvmontemperaturestate
    return leafs
}

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetYangName() string { return "ciscoEnvMonTemperatureStatusEntry" }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) SetParent(parent types.Entity) { ciscoenvmontemperaturestatusentry.parent = parent }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetParent() types.Entity { return ciscoenvmontemperaturestatusentry.parent }

func (ciscoenvmontemperaturestatusentry *CISCOENVMONMIB_Ciscoenvmontemperaturestatustable_Ciscoenvmontemperaturestatusentry) GetParentYangName() string { return "ciscoEnvMonTemperatureStatusTable" }

// CISCOENVMONMIB_Ciscoenvmonfanstatustable
// The table of fan status maintained by the environmental
// monitor.
type CISCOENVMONMIB_Ciscoenvmonfanstatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the fan status table, representing the status of the associated
    // fan maintained by the environmental monitor. The type is slice of
    // CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry.
    Ciscoenvmonfanstatusentry []CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetFilter() yfilter.YFilter { return ciscoenvmonfanstatustable.YFilter }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) SetFilter(yf yfilter.YFilter) { ciscoenvmonfanstatustable.YFilter = yf }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetGoName(yname string) string {
    if yname == "ciscoEnvMonFanStatusEntry" { return "Ciscoenvmonfanstatusentry" }
    return ""
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetSegmentPath() string {
    return "ciscoEnvMonFanStatusTable"
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoEnvMonFanStatusEntry" {
        for _, c := range ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry {
            if ciscoenvmonfanstatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry{}
        ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry = append(ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry, child)
        return &ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry[len(ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry)-1]
    }
    return nil
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry {
        children[ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry[i].GetSegmentPath()] = &ciscoenvmonfanstatustable.Ciscoenvmonfanstatusentry[i]
    }
    return children
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetYangName() string { return "ciscoEnvMonFanStatusTable" }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) SetParent(parent types.Entity) { ciscoenvmonfanstatustable.parent = parent }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetParent() types.Entity { return ciscoenvmonfanstatustable.parent }

func (ciscoenvmonfanstatustable *CISCOENVMONMIB_Ciscoenvmonfanstatustable) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

// CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry
// An entry in the fan status table, representing the status of
// the associated fan maintained by the environmental monitor.
type CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry struct {
    parent types.Entity
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

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetFilter() yfilter.YFilter { return ciscoenvmonfanstatusentry.YFilter }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) SetFilter(yf yfilter.YFilter) { ciscoenvmonfanstatusentry.YFilter = yf }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetGoName(yname string) string {
    if yname == "ciscoEnvMonFanStatusIndex" { return "Ciscoenvmonfanstatusindex" }
    if yname == "ciscoEnvMonFanStatusDescr" { return "Ciscoenvmonfanstatusdescr" }
    if yname == "ciscoEnvMonFanState" { return "Ciscoenvmonfanstate" }
    return ""
}

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetSegmentPath() string {
    return "ciscoEnvMonFanStatusEntry" + "[ciscoEnvMonFanStatusIndex='" + fmt.Sprintf("%v", ciscoenvmonfanstatusentry.Ciscoenvmonfanstatusindex) + "']"
}

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoEnvMonFanStatusIndex"] = ciscoenvmonfanstatusentry.Ciscoenvmonfanstatusindex
    leafs["ciscoEnvMonFanStatusDescr"] = ciscoenvmonfanstatusentry.Ciscoenvmonfanstatusdescr
    leafs["ciscoEnvMonFanState"] = ciscoenvmonfanstatusentry.Ciscoenvmonfanstate
    return leafs
}

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetYangName() string { return "ciscoEnvMonFanStatusEntry" }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) SetParent(parent types.Entity) { ciscoenvmonfanstatusentry.parent = parent }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetParent() types.Entity { return ciscoenvmonfanstatusentry.parent }

func (ciscoenvmonfanstatusentry *CISCOENVMONMIB_Ciscoenvmonfanstatustable_Ciscoenvmonfanstatusentry) GetParentYangName() string { return "ciscoEnvMonFanStatusTable" }

// CISCOENVMONMIB_Ciscoenvmonsupplystatustable
// The table of power supply status maintained by the
// environmental monitor card.
type CISCOENVMONMIB_Ciscoenvmonsupplystatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the power supply status table, representing the status of the
    // associated power supply maintained by the environmental monitor card. The
    // type is slice of
    // CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry.
    Ciscoenvmonsupplystatusentry []CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetFilter() yfilter.YFilter { return ciscoenvmonsupplystatustable.YFilter }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) SetFilter(yf yfilter.YFilter) { ciscoenvmonsupplystatustable.YFilter = yf }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetGoName(yname string) string {
    if yname == "ciscoEnvMonSupplyStatusEntry" { return "Ciscoenvmonsupplystatusentry" }
    return ""
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetSegmentPath() string {
    return "ciscoEnvMonSupplyStatusTable"
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoEnvMonSupplyStatusEntry" {
        for _, c := range ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry {
            if ciscoenvmonsupplystatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry{}
        ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry = append(ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry, child)
        return &ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry[len(ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry)-1]
    }
    return nil
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry {
        children[ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry[i].GetSegmentPath()] = &ciscoenvmonsupplystatustable.Ciscoenvmonsupplystatusentry[i]
    }
    return children
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetYangName() string { return "ciscoEnvMonSupplyStatusTable" }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) SetParent(parent types.Entity) { ciscoenvmonsupplystatustable.parent = parent }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetParent() types.Entity { return ciscoenvmonsupplystatustable.parent }

func (ciscoenvmonsupplystatustable *CISCOENVMONMIB_Ciscoenvmonsupplystatustable) GetParentYangName() string { return "CISCO-ENVMON-MIB" }

// CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry
// An entry in the power supply status table, representing the
// status of the associated power supply maintained by the
// environmental monitor card.
type CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry struct {
    parent types.Entity
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

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetFilter() yfilter.YFilter { return ciscoenvmonsupplystatusentry.YFilter }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) SetFilter(yf yfilter.YFilter) { ciscoenvmonsupplystatusentry.YFilter = yf }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetGoName(yname string) string {
    if yname == "ciscoEnvMonSupplyStatusIndex" { return "Ciscoenvmonsupplystatusindex" }
    if yname == "ciscoEnvMonSupplyStatusDescr" { return "Ciscoenvmonsupplystatusdescr" }
    if yname == "ciscoEnvMonSupplyState" { return "Ciscoenvmonsupplystate" }
    if yname == "ciscoEnvMonSupplySource" { return "Ciscoenvmonsupplysource" }
    return ""
}

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetSegmentPath() string {
    return "ciscoEnvMonSupplyStatusEntry" + "[ciscoEnvMonSupplyStatusIndex='" + fmt.Sprintf("%v", ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystatusindex) + "']"
}

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoEnvMonSupplyStatusIndex"] = ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystatusindex
    leafs["ciscoEnvMonSupplyStatusDescr"] = ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystatusdescr
    leafs["ciscoEnvMonSupplyState"] = ciscoenvmonsupplystatusentry.Ciscoenvmonsupplystate
    leafs["ciscoEnvMonSupplySource"] = ciscoenvmonsupplystatusentry.Ciscoenvmonsupplysource
    return leafs
}

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetYangName() string { return "ciscoEnvMonSupplyStatusEntry" }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) SetParent(parent types.Entity) { ciscoenvmonsupplystatusentry.parent = parent }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetParent() types.Entity { return ciscoenvmonsupplystatusentry.parent }

func (ciscoenvmonsupplystatusentry *CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry) GetParentYangName() string { return "ciscoEnvMonSupplyStatusTable" }

// CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource represents internalRedundant - Internal redundant power supply 
type CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource string

const (
    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_unknown CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "unknown"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_ac CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "ac"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_dc CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "dc"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_externalPowerSupply CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "externalPowerSupply"

    CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource_internalRedundant CISCOENVMONMIB_Ciscoenvmonsupplystatustable_Ciscoenvmonsupplystatusentry_Ciscoenvmonsupplysource = "internalRedundant"
)

