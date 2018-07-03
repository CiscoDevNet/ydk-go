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

    
    CiscoEnvMonObjects CISCOENVMONMIB_CiscoEnvMonObjects

    
    CiscoEnvMonMIBNotificationEnables CISCOENVMONMIB_CiscoEnvMonMIBNotificationEnables

    // The table of voltage status maintained by the environmental monitor.
    CiscoEnvMonVoltageStatusTable CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable

    // The table of ambient temperature status maintained by the environmental
    // monitor.
    CiscoEnvMonTemperatureStatusTable CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable

    // The table of fan status maintained by the environmental monitor.
    CiscoEnvMonFanStatusTable CISCOENVMONMIB_CiscoEnvMonFanStatusTable

    // The table of power supply status maintained by the environmental monitor
    // card.
    CiscoEnvMonSupplyStatusTable CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable
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

    cISCOENVMONMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENVMONMIB.EntityData.Children.Append("ciscoEnvMonObjects", types.YChild{"CiscoEnvMonObjects", &cISCOENVMONMIB.CiscoEnvMonObjects})
    cISCOENVMONMIB.EntityData.Children.Append("ciscoEnvMonMIBNotificationEnables", types.YChild{"CiscoEnvMonMIBNotificationEnables", &cISCOENVMONMIB.CiscoEnvMonMIBNotificationEnables})
    cISCOENVMONMIB.EntityData.Children.Append("ciscoEnvMonVoltageStatusTable", types.YChild{"CiscoEnvMonVoltageStatusTable", &cISCOENVMONMIB.CiscoEnvMonVoltageStatusTable})
    cISCOENVMONMIB.EntityData.Children.Append("ciscoEnvMonTemperatureStatusTable", types.YChild{"CiscoEnvMonTemperatureStatusTable", &cISCOENVMONMIB.CiscoEnvMonTemperatureStatusTable})
    cISCOENVMONMIB.EntityData.Children.Append("ciscoEnvMonFanStatusTable", types.YChild{"CiscoEnvMonFanStatusTable", &cISCOENVMONMIB.CiscoEnvMonFanStatusTable})
    cISCOENVMONMIB.EntityData.Children.Append("ciscoEnvMonSupplyStatusTable", types.YChild{"CiscoEnvMonSupplyStatusTable", &cISCOENVMONMIB.CiscoEnvMonSupplyStatusTable})
    cISCOENVMONMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENVMONMIB.EntityData.YListKeys = []string {}

    return &(cISCOENVMONMIB.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonObjects
type CISCOENVMONMIB_CiscoEnvMonObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of environmental monitor located in the chassis. An oldAgs
    // environmental monitor card is identical to an ags environmental card except
    // that it is not capable of supplying data, and hence no instance of the
    // remaining objects in this MIB will be returned in response to an SNMP
    // query.  Note that only a firmware upgrade is required to convert an oldAgs
    // into an ags card. The type is CiscoEnvMonPresent.
    CiscoEnvMonPresent interface{}

    // Each bit is set to reflect the respective alarm being set.  The bit will be
    // cleared when the respective alarm is cleared. The type is map[string]bool.
    CiscoEnvMonAlarmContacts interface{}
}

func (ciscoEnvMonObjects *CISCOENVMONMIB_CiscoEnvMonObjects) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonObjects.EntityData.YFilter = ciscoEnvMonObjects.YFilter
    ciscoEnvMonObjects.EntityData.YangName = "ciscoEnvMonObjects"
    ciscoEnvMonObjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonObjects.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoEnvMonObjects.EntityData.SegmentPath = "ciscoEnvMonObjects"
    ciscoEnvMonObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonObjects.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonObjects.EntityData.Leafs = types.NewOrderedMap()
    ciscoEnvMonObjects.EntityData.Leafs.Append("ciscoEnvMonPresent", types.YLeaf{"CiscoEnvMonPresent", ciscoEnvMonObjects.CiscoEnvMonPresent})
    ciscoEnvMonObjects.EntityData.Leafs.Append("ciscoEnvMonAlarmContacts", types.YLeaf{"CiscoEnvMonAlarmContacts", ciscoEnvMonObjects.CiscoEnvMonAlarmContacts})

    ciscoEnvMonObjects.EntityData.YListKeys = []string {}

    return &(ciscoEnvMonObjects.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent represents an ags card.
type CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent string

const (
    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_oldAgs CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "oldAgs"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_ags CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "ags"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_c7000 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "c7000"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_ci CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "ci"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_cAccessMon CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "cAccessMon"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_cat6000 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "cat6000"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_ubr7200 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "ubr7200"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_cat4000 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "cat4000"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_c10000 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "c10000"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_osr7600 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "osr7600"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_c7600 CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "c7600"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_c37xx CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "c37xx"

    CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent_other CISCOENVMONMIB_CiscoEnvMonObjects_CiscoEnvMonPresent = "other"
)

// CISCOENVMONMIB_CiscoEnvMonMIBNotificationEnables
type CISCOENVMONMIB_CiscoEnvMonMIBNotificationEnables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonShutdownNotification.  A false  value will prevent shutdown
    // notifications  from being generated by this system. The type is bool.
    CiscoEnvMonEnableShutdownNotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonVoltageNotification. A false  value will prevent voltage
    // notifications from being  generated by this system. This object is
    // deprecated in favour of ciscoEnvMonEnableStatChangeNotif. The type is bool.
    CiscoEnvMonEnableVoltageNotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonTemperatureNotification. A false value prevents temperature
    // notifications  from being sent by  this entity. This object is  deprecated
    // in favour of  ciscoEnvMonEnableStatChangeNotif. The type is bool.
    CiscoEnvMonEnableTemperatureNotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonFanNotification. A false value prevents fan notifications  from
    // being sent by  this entity. This object is  deprecated in favour of 
    // ciscoEnvMonEnableStatChangeNotif. The type is bool.
    CiscoEnvMonEnableFanNotification interface{}

    // This variable  indicates  whether  the  system produces the
    // ciscoEnvMonRedundantSupplyNotification.  A false value prevents redundant
    // supply notifications from being generated by this system. This object is
    // deprecated in favour of  ciscoEnvMonEnableStatChangeNotif. The type is
    // bool.
    CiscoEnvMonEnableRedundantSupplyNotification interface{}

    // This variable indicates whether the system produces the
    // ciscoEnvMonVoltStatusChangeNotif, ciscoEnvMonTempStatusChangeNotif, 
    // ciscoEnvMonFanStatusChangeNotif and   ciscoEnvMonSuppStatusChangeNotif. A
    // false value will  prevent these notifications from being generated by  this
    // system. The type is bool.
    CiscoEnvMonEnableStatChangeNotif interface{}
}

func (ciscoEnvMonMIBNotificationEnables *CISCOENVMONMIB_CiscoEnvMonMIBNotificationEnables) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonMIBNotificationEnables.EntityData.YFilter = ciscoEnvMonMIBNotificationEnables.YFilter
    ciscoEnvMonMIBNotificationEnables.EntityData.YangName = "ciscoEnvMonMIBNotificationEnables"
    ciscoEnvMonMIBNotificationEnables.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonMIBNotificationEnables.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoEnvMonMIBNotificationEnables.EntityData.SegmentPath = "ciscoEnvMonMIBNotificationEnables"
    ciscoEnvMonMIBNotificationEnables.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonMIBNotificationEnables.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonMIBNotificationEnables.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonMIBNotificationEnables.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs = types.NewOrderedMap()
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs.Append("ciscoEnvMonEnableShutdownNotification", types.YLeaf{"CiscoEnvMonEnableShutdownNotification", ciscoEnvMonMIBNotificationEnables.CiscoEnvMonEnableShutdownNotification})
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs.Append("ciscoEnvMonEnableVoltageNotification", types.YLeaf{"CiscoEnvMonEnableVoltageNotification", ciscoEnvMonMIBNotificationEnables.CiscoEnvMonEnableVoltageNotification})
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs.Append("ciscoEnvMonEnableTemperatureNotification", types.YLeaf{"CiscoEnvMonEnableTemperatureNotification", ciscoEnvMonMIBNotificationEnables.CiscoEnvMonEnableTemperatureNotification})
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs.Append("ciscoEnvMonEnableFanNotification", types.YLeaf{"CiscoEnvMonEnableFanNotification", ciscoEnvMonMIBNotificationEnables.CiscoEnvMonEnableFanNotification})
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs.Append("ciscoEnvMonEnableRedundantSupplyNotification", types.YLeaf{"CiscoEnvMonEnableRedundantSupplyNotification", ciscoEnvMonMIBNotificationEnables.CiscoEnvMonEnableRedundantSupplyNotification})
    ciscoEnvMonMIBNotificationEnables.EntityData.Leafs.Append("ciscoEnvMonEnableStatChangeNotif", types.YLeaf{"CiscoEnvMonEnableStatChangeNotif", ciscoEnvMonMIBNotificationEnables.CiscoEnvMonEnableStatChangeNotif})

    ciscoEnvMonMIBNotificationEnables.EntityData.YListKeys = []string {}

    return &(ciscoEnvMonMIBNotificationEnables.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable
// The table of voltage status maintained by the environmental
// monitor.
type CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the voltage status table, representing the status of the
    // associated testpoint maintained by the environmental monitor. The type is
    // slice of
    // CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable_CiscoEnvMonVoltageStatusEntry.
    CiscoEnvMonVoltageStatusEntry []*CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable_CiscoEnvMonVoltageStatusEntry
}

func (ciscoEnvMonVoltageStatusTable *CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonVoltageStatusTable.EntityData.YFilter = ciscoEnvMonVoltageStatusTable.YFilter
    ciscoEnvMonVoltageStatusTable.EntityData.YangName = "ciscoEnvMonVoltageStatusTable"
    ciscoEnvMonVoltageStatusTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonVoltageStatusTable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoEnvMonVoltageStatusTable.EntityData.SegmentPath = "ciscoEnvMonVoltageStatusTable"
    ciscoEnvMonVoltageStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonVoltageStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonVoltageStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonVoltageStatusTable.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonVoltageStatusTable.EntityData.Children.Append("ciscoEnvMonVoltageStatusEntry", types.YChild{"CiscoEnvMonVoltageStatusEntry", nil})
    for i := range ciscoEnvMonVoltageStatusTable.CiscoEnvMonVoltageStatusEntry {
        ciscoEnvMonVoltageStatusTable.EntityData.Children.Append(types.GetSegmentPath(ciscoEnvMonVoltageStatusTable.CiscoEnvMonVoltageStatusEntry[i]), types.YChild{"CiscoEnvMonVoltageStatusEntry", ciscoEnvMonVoltageStatusTable.CiscoEnvMonVoltageStatusEntry[i]})
    }
    ciscoEnvMonVoltageStatusTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoEnvMonVoltageStatusTable.EntityData.YListKeys = []string {}

    return &(ciscoEnvMonVoltageStatusTable.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable_CiscoEnvMonVoltageStatusEntry
// An entry in the voltage status table, representing the status
// of the associated testpoint maintained by the environmental
// monitor.
type CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable_CiscoEnvMonVoltageStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the testpoint being instrumented.
    // This index is for SNMP purposes only, and has no intrinsic meaning. The
    // type is interface{} with range: 0..2147483647.
    CiscoEnvMonVoltageStatusIndex interface{}

    // Textual description of the testpoint being instrumented. This description
    // is a short textual label, suitable as a human-sensible identification for
    // the rest of the information in the entry. The type is string with length:
    // 0..32.
    CiscoEnvMonVoltageStatusDescr interface{}

    // The current measurement of the testpoint being instrumented. The type is
    // interface{} with range: -2147483648..2147483647. Units are millivolts.
    CiscoEnvMonVoltageStatusValue interface{}

    // The lowest value that the associated instance of the object
    // ciscoEnvMonVoltageStatusValue may obtain before an emergency shutdown of
    // the managed device is initiated. The type is interface{} with range:
    // -2147483648..2147483647. Units are millivolts.
    CiscoEnvMonVoltageThresholdLow interface{}

    // The highest value that the associated instance of the object
    // ciscoEnvMonVoltageStatusValue may obtain before an emergency shutdown of
    // the managed device is initiated. The type is interface{} with range:
    // -2147483648..2147483647. Units are millivolts.
    CiscoEnvMonVoltageThresholdHigh interface{}

    // The value of the associated instance of the object
    // ciscoEnvMonVoltageStatusValue at the time an emergency shutdown of the
    // managed device was last initiated.  This value is stored in non-volatile
    // RAM and hence is able to survive the shutdown. The type is interface{} with
    // range: -2147483648..2147483647. Units are millivolts.
    CiscoEnvMonVoltageLastShutdown interface{}

    // The current state of the testpoint being instrumented. The type is
    // CiscoEnvMonState.
    CiscoEnvMonVoltageState interface{}
}

func (ciscoEnvMonVoltageStatusEntry *CISCOENVMONMIB_CiscoEnvMonVoltageStatusTable_CiscoEnvMonVoltageStatusEntry) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonVoltageStatusEntry.EntityData.YFilter = ciscoEnvMonVoltageStatusEntry.YFilter
    ciscoEnvMonVoltageStatusEntry.EntityData.YangName = "ciscoEnvMonVoltageStatusEntry"
    ciscoEnvMonVoltageStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonVoltageStatusEntry.EntityData.ParentYangName = "ciscoEnvMonVoltageStatusTable"
    ciscoEnvMonVoltageStatusEntry.EntityData.SegmentPath = "ciscoEnvMonVoltageStatusEntry" + types.AddKeyToken(ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageStatusIndex, "ciscoEnvMonVoltageStatusIndex")
    ciscoEnvMonVoltageStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonVoltageStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonVoltageStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonVoltageStatusEntry.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageStatusIndex", types.YLeaf{"CiscoEnvMonVoltageStatusIndex", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageStatusIndex})
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageStatusDescr", types.YLeaf{"CiscoEnvMonVoltageStatusDescr", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageStatusDescr})
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageStatusValue", types.YLeaf{"CiscoEnvMonVoltageStatusValue", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageStatusValue})
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageThresholdLow", types.YLeaf{"CiscoEnvMonVoltageThresholdLow", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageThresholdLow})
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageThresholdHigh", types.YLeaf{"CiscoEnvMonVoltageThresholdHigh", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageThresholdHigh})
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageLastShutdown", types.YLeaf{"CiscoEnvMonVoltageLastShutdown", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageLastShutdown})
    ciscoEnvMonVoltageStatusEntry.EntityData.Leafs.Append("ciscoEnvMonVoltageState", types.YLeaf{"CiscoEnvMonVoltageState", ciscoEnvMonVoltageStatusEntry.CiscoEnvMonVoltageState})

    ciscoEnvMonVoltageStatusEntry.EntityData.YListKeys = []string {"CiscoEnvMonVoltageStatusIndex"}

    return &(ciscoEnvMonVoltageStatusEntry.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable
// The table of ambient temperature status maintained by the
// environmental monitor.
type CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the ambient temperature status table, representing the status
    // of the associated testpoint maintained by the environmental monitor. The
    // type is slice of
    // CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable_CiscoEnvMonTemperatureStatusEntry.
    CiscoEnvMonTemperatureStatusEntry []*CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable_CiscoEnvMonTemperatureStatusEntry
}

func (ciscoEnvMonTemperatureStatusTable *CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonTemperatureStatusTable.EntityData.YFilter = ciscoEnvMonTemperatureStatusTable.YFilter
    ciscoEnvMonTemperatureStatusTable.EntityData.YangName = "ciscoEnvMonTemperatureStatusTable"
    ciscoEnvMonTemperatureStatusTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonTemperatureStatusTable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoEnvMonTemperatureStatusTable.EntityData.SegmentPath = "ciscoEnvMonTemperatureStatusTable"
    ciscoEnvMonTemperatureStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonTemperatureStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonTemperatureStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonTemperatureStatusTable.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonTemperatureStatusTable.EntityData.Children.Append("ciscoEnvMonTemperatureStatusEntry", types.YChild{"CiscoEnvMonTemperatureStatusEntry", nil})
    for i := range ciscoEnvMonTemperatureStatusTable.CiscoEnvMonTemperatureStatusEntry {
        ciscoEnvMonTemperatureStatusTable.EntityData.Children.Append(types.GetSegmentPath(ciscoEnvMonTemperatureStatusTable.CiscoEnvMonTemperatureStatusEntry[i]), types.YChild{"CiscoEnvMonTemperatureStatusEntry", ciscoEnvMonTemperatureStatusTable.CiscoEnvMonTemperatureStatusEntry[i]})
    }
    ciscoEnvMonTemperatureStatusTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoEnvMonTemperatureStatusTable.EntityData.YListKeys = []string {}

    return &(ciscoEnvMonTemperatureStatusTable.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable_CiscoEnvMonTemperatureStatusEntry
// An entry in the ambient temperature status table, representing
// the status of the associated testpoint maintained by the
// environmental monitor.
type CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable_CiscoEnvMonTemperatureStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the testpoint being instrumented.
    // This index is for SNMP purposes only, and has no intrinsic meaning. The
    // type is interface{} with range: 0..2147483647.
    CiscoEnvMonTemperatureStatusIndex interface{}

    // Textual description of the testpoint being instrumented. This description
    // is a short textual label, suitable as a human-sensible identification for
    // the rest of the information in the entry. The type is string with length:
    // 0..32.
    CiscoEnvMonTemperatureStatusDescr interface{}

    // The current measurement of the testpoint being instrumented. The type is
    // interface{} with range: 0..4294967295. Units are degrees Celsius.
    CiscoEnvMonTemperatureStatusValue interface{}

    // The highest value that the associated instance of the object
    // ciscoEnvMonTemperatureStatusValue may obtain before an emergency shutdown
    // of the managed device is initiated. The type is interface{} with range:
    // -2147483648..2147483647. Units are degrees Celsius.
    CiscoEnvMonTemperatureThreshold interface{}

    // The value of the associated instance of the object
    // ciscoEnvMonTemperatureStatusValue at the time an emergency shutdown of the
    // managed device was last initiated.  This value is stored in non-volatile
    // RAM and hence is able to survive the shutdown. The type is interface{} with
    // range: -2147483648..2147483647. Units are degrees Celsius.
    CiscoEnvMonTemperatureLastShutdown interface{}

    // The current state of the testpoint being instrumented. The type is
    // CiscoEnvMonState.
    CiscoEnvMonTemperatureState interface{}
}

func (ciscoEnvMonTemperatureStatusEntry *CISCOENVMONMIB_CiscoEnvMonTemperatureStatusTable_CiscoEnvMonTemperatureStatusEntry) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonTemperatureStatusEntry.EntityData.YFilter = ciscoEnvMonTemperatureStatusEntry.YFilter
    ciscoEnvMonTemperatureStatusEntry.EntityData.YangName = "ciscoEnvMonTemperatureStatusEntry"
    ciscoEnvMonTemperatureStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonTemperatureStatusEntry.EntityData.ParentYangName = "ciscoEnvMonTemperatureStatusTable"
    ciscoEnvMonTemperatureStatusEntry.EntityData.SegmentPath = "ciscoEnvMonTemperatureStatusEntry" + types.AddKeyToken(ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureStatusIndex, "ciscoEnvMonTemperatureStatusIndex")
    ciscoEnvMonTemperatureStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonTemperatureStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonTemperatureStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonTemperatureStatusEntry.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs.Append("ciscoEnvMonTemperatureStatusIndex", types.YLeaf{"CiscoEnvMonTemperatureStatusIndex", ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureStatusIndex})
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs.Append("ciscoEnvMonTemperatureStatusDescr", types.YLeaf{"CiscoEnvMonTemperatureStatusDescr", ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureStatusDescr})
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs.Append("ciscoEnvMonTemperatureStatusValue", types.YLeaf{"CiscoEnvMonTemperatureStatusValue", ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureStatusValue})
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs.Append("ciscoEnvMonTemperatureThreshold", types.YLeaf{"CiscoEnvMonTemperatureThreshold", ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureThreshold})
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs.Append("ciscoEnvMonTemperatureLastShutdown", types.YLeaf{"CiscoEnvMonTemperatureLastShutdown", ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureLastShutdown})
    ciscoEnvMonTemperatureStatusEntry.EntityData.Leafs.Append("ciscoEnvMonTemperatureState", types.YLeaf{"CiscoEnvMonTemperatureState", ciscoEnvMonTemperatureStatusEntry.CiscoEnvMonTemperatureState})

    ciscoEnvMonTemperatureStatusEntry.EntityData.YListKeys = []string {"CiscoEnvMonTemperatureStatusIndex"}

    return &(ciscoEnvMonTemperatureStatusEntry.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonFanStatusTable
// The table of fan status maintained by the environmental
// monitor.
type CISCOENVMONMIB_CiscoEnvMonFanStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the fan status table, representing the status of the associated
    // fan maintained by the environmental monitor. The type is slice of
    // CISCOENVMONMIB_CiscoEnvMonFanStatusTable_CiscoEnvMonFanStatusEntry.
    CiscoEnvMonFanStatusEntry []*CISCOENVMONMIB_CiscoEnvMonFanStatusTable_CiscoEnvMonFanStatusEntry
}

func (ciscoEnvMonFanStatusTable *CISCOENVMONMIB_CiscoEnvMonFanStatusTable) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonFanStatusTable.EntityData.YFilter = ciscoEnvMonFanStatusTable.YFilter
    ciscoEnvMonFanStatusTable.EntityData.YangName = "ciscoEnvMonFanStatusTable"
    ciscoEnvMonFanStatusTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonFanStatusTable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoEnvMonFanStatusTable.EntityData.SegmentPath = "ciscoEnvMonFanStatusTable"
    ciscoEnvMonFanStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonFanStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonFanStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonFanStatusTable.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonFanStatusTable.EntityData.Children.Append("ciscoEnvMonFanStatusEntry", types.YChild{"CiscoEnvMonFanStatusEntry", nil})
    for i := range ciscoEnvMonFanStatusTable.CiscoEnvMonFanStatusEntry {
        ciscoEnvMonFanStatusTable.EntityData.Children.Append(types.GetSegmentPath(ciscoEnvMonFanStatusTable.CiscoEnvMonFanStatusEntry[i]), types.YChild{"CiscoEnvMonFanStatusEntry", ciscoEnvMonFanStatusTable.CiscoEnvMonFanStatusEntry[i]})
    }
    ciscoEnvMonFanStatusTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoEnvMonFanStatusTable.EntityData.YListKeys = []string {}

    return &(ciscoEnvMonFanStatusTable.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonFanStatusTable_CiscoEnvMonFanStatusEntry
// An entry in the fan status table, representing the status of
// the associated fan maintained by the environmental monitor.
type CISCOENVMONMIB_CiscoEnvMonFanStatusTable_CiscoEnvMonFanStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the fan being instrumented. This
    // index is for SNMP purposes only, and has no intrinsic meaning. The type is
    // interface{} with range: 0..2147483647.
    CiscoEnvMonFanStatusIndex interface{}

    // Textual description of the fan being instrumented. This description is a
    // short textual label, suitable as a human-sensible identification for the
    // rest of the information in the entry. The type is string with length:
    // 0..32.
    CiscoEnvMonFanStatusDescr interface{}

    // The current state of the fan being instrumented. The type is
    // CiscoEnvMonState.
    CiscoEnvMonFanState interface{}
}

func (ciscoEnvMonFanStatusEntry *CISCOENVMONMIB_CiscoEnvMonFanStatusTable_CiscoEnvMonFanStatusEntry) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonFanStatusEntry.EntityData.YFilter = ciscoEnvMonFanStatusEntry.YFilter
    ciscoEnvMonFanStatusEntry.EntityData.YangName = "ciscoEnvMonFanStatusEntry"
    ciscoEnvMonFanStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonFanStatusEntry.EntityData.ParentYangName = "ciscoEnvMonFanStatusTable"
    ciscoEnvMonFanStatusEntry.EntityData.SegmentPath = "ciscoEnvMonFanStatusEntry" + types.AddKeyToken(ciscoEnvMonFanStatusEntry.CiscoEnvMonFanStatusIndex, "ciscoEnvMonFanStatusIndex")
    ciscoEnvMonFanStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonFanStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonFanStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonFanStatusEntry.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonFanStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoEnvMonFanStatusEntry.EntityData.Leafs.Append("ciscoEnvMonFanStatusIndex", types.YLeaf{"CiscoEnvMonFanStatusIndex", ciscoEnvMonFanStatusEntry.CiscoEnvMonFanStatusIndex})
    ciscoEnvMonFanStatusEntry.EntityData.Leafs.Append("ciscoEnvMonFanStatusDescr", types.YLeaf{"CiscoEnvMonFanStatusDescr", ciscoEnvMonFanStatusEntry.CiscoEnvMonFanStatusDescr})
    ciscoEnvMonFanStatusEntry.EntityData.Leafs.Append("ciscoEnvMonFanState", types.YLeaf{"CiscoEnvMonFanState", ciscoEnvMonFanStatusEntry.CiscoEnvMonFanState})

    ciscoEnvMonFanStatusEntry.EntityData.YListKeys = []string {"CiscoEnvMonFanStatusIndex"}

    return &(ciscoEnvMonFanStatusEntry.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable
// The table of power supply status maintained by the
// environmental monitor card.
type CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the power supply status table, representing the status of the
    // associated power supply maintained by the environmental monitor card. The
    // type is slice of
    // CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry.
    CiscoEnvMonSupplyStatusEntry []*CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry
}

func (ciscoEnvMonSupplyStatusTable *CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonSupplyStatusTable.EntityData.YFilter = ciscoEnvMonSupplyStatusTable.YFilter
    ciscoEnvMonSupplyStatusTable.EntityData.YangName = "ciscoEnvMonSupplyStatusTable"
    ciscoEnvMonSupplyStatusTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonSupplyStatusTable.EntityData.ParentYangName = "CISCO-ENVMON-MIB"
    ciscoEnvMonSupplyStatusTable.EntityData.SegmentPath = "ciscoEnvMonSupplyStatusTable"
    ciscoEnvMonSupplyStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonSupplyStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonSupplyStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonSupplyStatusTable.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonSupplyStatusTable.EntityData.Children.Append("ciscoEnvMonSupplyStatusEntry", types.YChild{"CiscoEnvMonSupplyStatusEntry", nil})
    for i := range ciscoEnvMonSupplyStatusTable.CiscoEnvMonSupplyStatusEntry {
        ciscoEnvMonSupplyStatusTable.EntityData.Children.Append(types.GetSegmentPath(ciscoEnvMonSupplyStatusTable.CiscoEnvMonSupplyStatusEntry[i]), types.YChild{"CiscoEnvMonSupplyStatusEntry", ciscoEnvMonSupplyStatusTable.CiscoEnvMonSupplyStatusEntry[i]})
    }
    ciscoEnvMonSupplyStatusTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoEnvMonSupplyStatusTable.EntityData.YListKeys = []string {}

    return &(ciscoEnvMonSupplyStatusTable.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry
// An entry in the power supply status table, representing the
// status of the associated power supply maintained by the
// environmental monitor card.
type CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique index for the power supply being
    // instrumented. This index is for SNMP purposes only, and has no intrinsic
    // meaning. The type is interface{} with range: 0..2147483647.
    CiscoEnvMonSupplyStatusIndex interface{}

    // Textual description of the power supply being instrumented. This
    // description is a short textual label, suitable as a human-sensible
    // identification for the rest of the information in the entry. The type is
    // string with length: 0..64.
    CiscoEnvMonSupplyStatusDescr interface{}

    // The current state of the power supply being instrumented. The type is
    // CiscoEnvMonState.
    CiscoEnvMonSupplyState interface{}

    // The power supply source. unknown - Power supply source unknown ac      - AC
    // power supply dc      - DC power supply externalPowerSupply - External power
    // supply internalRedundant - Internal redundant power supply . The type is
    // CiscoEnvMonSupplySource.
    CiscoEnvMonSupplySource interface{}
}

func (ciscoEnvMonSupplyStatusEntry *CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry) GetEntityData() *types.CommonEntityData {
    ciscoEnvMonSupplyStatusEntry.EntityData.YFilter = ciscoEnvMonSupplyStatusEntry.YFilter
    ciscoEnvMonSupplyStatusEntry.EntityData.YangName = "ciscoEnvMonSupplyStatusEntry"
    ciscoEnvMonSupplyStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoEnvMonSupplyStatusEntry.EntityData.ParentYangName = "ciscoEnvMonSupplyStatusTable"
    ciscoEnvMonSupplyStatusEntry.EntityData.SegmentPath = "ciscoEnvMonSupplyStatusEntry" + types.AddKeyToken(ciscoEnvMonSupplyStatusEntry.CiscoEnvMonSupplyStatusIndex, "ciscoEnvMonSupplyStatusIndex")
    ciscoEnvMonSupplyStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEnvMonSupplyStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEnvMonSupplyStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEnvMonSupplyStatusEntry.EntityData.Children = types.NewOrderedMap()
    ciscoEnvMonSupplyStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoEnvMonSupplyStatusEntry.EntityData.Leafs.Append("ciscoEnvMonSupplyStatusIndex", types.YLeaf{"CiscoEnvMonSupplyStatusIndex", ciscoEnvMonSupplyStatusEntry.CiscoEnvMonSupplyStatusIndex})
    ciscoEnvMonSupplyStatusEntry.EntityData.Leafs.Append("ciscoEnvMonSupplyStatusDescr", types.YLeaf{"CiscoEnvMonSupplyStatusDescr", ciscoEnvMonSupplyStatusEntry.CiscoEnvMonSupplyStatusDescr})
    ciscoEnvMonSupplyStatusEntry.EntityData.Leafs.Append("ciscoEnvMonSupplyState", types.YLeaf{"CiscoEnvMonSupplyState", ciscoEnvMonSupplyStatusEntry.CiscoEnvMonSupplyState})
    ciscoEnvMonSupplyStatusEntry.EntityData.Leafs.Append("ciscoEnvMonSupplySource", types.YLeaf{"CiscoEnvMonSupplySource", ciscoEnvMonSupplyStatusEntry.CiscoEnvMonSupplySource})

    ciscoEnvMonSupplyStatusEntry.EntityData.YListKeys = []string {"CiscoEnvMonSupplyStatusIndex"}

    return &(ciscoEnvMonSupplyStatusEntry.EntityData)
}

// CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource represents internalRedundant - Internal redundant power supply 
type CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource string

const (
    CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource_unknown CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource = "unknown"

    CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource_ac CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource = "ac"

    CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource_dc CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource = "dc"

    CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource_externalPowerSupply CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource = "externalPowerSupply"

    CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource_internalRedundant CISCOENVMONMIB_CiscoEnvMonSupplyStatusTable_CiscoEnvMonSupplyStatusEntry_CiscoEnvMonSupplySource = "internalRedundant"
)

