// The Synchronous Ethernet (SyncE) MIB is defined
// for monitoring network synchronization based on
// ITU-T G.781 clock selection.
// 
// Synchronous Ethernet (SyncE) is a standard defined
// for delivering timing to the remote NEs through a 
// Packet Network.   SyncE is well defined by ITU-T
// which included G.8261, G.8262, G.8264 and G.781.
// It leverages the PHY layer of Ethernet to transmit
// frequency to the remote sites.  Its functionality
// and accuracy mimics that of the SONET/SDH network
// because of its physical layer characteristic.
// In order to allow best clock source traceabiliy,
// correctly define the timing source and helps
// preventing timing loop, Synchronization Status
// Message is required for SyncE.  This is similar
// to SONET/SDH.  However, since SONET/SDH use 4 bits
// from the two S bytes in the SONET/SDH overhead
// frame for such message, Ethernet relies on a
// different channel called ESMC (Ethernet
// Synchronization Messaging Channel) which is based
// on IEEE 802.3 Organization Specific Slow Protocol.
// 
// Glossary:
// AIS: Alarm Indication Signal
// ATM: Asynchronous Transfer Mode
// EEC: Ethernet Equipment Clock
// ESMC: Ethernet Synchronization Messaging Channel
// QL: Quality Level
// SASE: Stand Alone Synchronization Equipment
// SSM: Synchronization Status Messaging
package cisco_netsync_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_netsync_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-NETSYNC-MIB CISCO-NETSYNC-MIB}", reflect.TypeOf(CISCONETSYNCMIB{}))
    ydk.RegisterEntity("CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB", reflect.TypeOf(CISCONETSYNCMIB{}))
}

// CiscoNetsyncESMCCap represents netsyncESMCCapInvalid   Capability invalid or unsupported
type CiscoNetsyncESMCCap string

const (
    CiscoNetsyncESMCCap_netsyncESMCCapNone CiscoNetsyncESMCCap = "netsyncESMCCapNone"

    CiscoNetsyncESMCCap_netsyncESMCCapTxRx CiscoNetsyncESMCCap = "netsyncESMCCapTxRx"

    CiscoNetsyncESMCCap_netsyncESMCCapTx CiscoNetsyncESMCCap = "netsyncESMCCapTx"

    CiscoNetsyncESMCCap_netsyncESMCCapRx CiscoNetsyncESMCCap = "netsyncESMCCapRx"

    CiscoNetsyncESMCCap_netsyncESMCCapInvalid CiscoNetsyncESMCCap = "netsyncESMCCapInvalid"
)

// CiscoNetsyncQLMode represents                           clock selection criteria.
type CiscoNetsyncQLMode string

const (
    CiscoNetsyncQLMode_netsyncQLModeUnknown CiscoNetsyncQLMode = "netsyncQLModeUnknown"

    CiscoNetsyncQLMode_netsyncQLModeQlDisabled CiscoNetsyncQLMode = "netsyncQLModeQlDisabled"

    CiscoNetsyncQLMode_netsyncQLModeQlEnabled CiscoNetsyncQLMode = "netsyncQLModeQlEnabled"
)

// CiscoNetsyncQualityLevel represents QL-NSUPP   Not supporting the SSM processing        
type CiscoNetsyncQualityLevel string

const (
    CiscoNetsyncQualityLevel_netsyncQualityLevelNULL CiscoNetsyncQualityLevel = "netsyncQualityLevelNULL"

    CiscoNetsyncQualityLevel_netsyncQualityLevelDNU CiscoNetsyncQualityLevel = "netsyncQualityLevelDNU"

    CiscoNetsyncQualityLevel_netsyncQualityLevelDUS CiscoNetsyncQualityLevel = "netsyncQualityLevelDUS"

    CiscoNetsyncQualityLevel_netsyncQualityLevelFAILED CiscoNetsyncQualityLevel = "netsyncQualityLevelFAILED"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV0 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV0"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV1 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV1"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV2 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV2"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV3 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV3"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV4 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV4"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV5 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV5"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV6 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV6"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV7 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV7"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV8 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV8"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV9 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV9"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV10 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV10"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV11 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV11"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV12 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV12"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV13 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV13"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV14 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV14"

    CiscoNetsyncQualityLevel_netsyncQualityLevelINV15 CiscoNetsyncQualityLevel = "netsyncQualityLevelINV15"

    CiscoNetsyncQualityLevel_netsyncQualityLevelNSUPP CiscoNetsyncQualityLevel = "netsyncQualityLevelNSUPP"

    CiscoNetsyncQualityLevel_netsyncQualityLevelPRC CiscoNetsyncQualityLevel = "netsyncQualityLevelPRC"

    CiscoNetsyncQualityLevel_netsyncQualityLevelPROV CiscoNetsyncQualityLevel = "netsyncQualityLevelPROV"

    CiscoNetsyncQualityLevel_netsyncQualityLevelPRS CiscoNetsyncQualityLevel = "netsyncQualityLevelPRS"

    CiscoNetsyncQualityLevel_netsyncQualityLevelSEC CiscoNetsyncQualityLevel = "netsyncQualityLevelSEC"

    CiscoNetsyncQualityLevel_netsyncQualityLevelSMC CiscoNetsyncQualityLevel = "netsyncQualityLevelSMC"

    CiscoNetsyncQualityLevel_netsyncQualityLevelSSUA CiscoNetsyncQualityLevel = "netsyncQualityLevelSSUA"

    CiscoNetsyncQualityLevel_netsyncQualityLevelSSUB CiscoNetsyncQualityLevel = "netsyncQualityLevelSSUB"

    CiscoNetsyncQualityLevel_netsyncQualityLevelST2 CiscoNetsyncQualityLevel = "netsyncQualityLevelST2"

    CiscoNetsyncQualityLevel_netsyncQualityLevelST3 CiscoNetsyncQualityLevel = "netsyncQualityLevelST3"

    CiscoNetsyncQualityLevel_netsyncQualityLevelST3E CiscoNetsyncQualityLevel = "netsyncQualityLevelST3E"

    CiscoNetsyncQualityLevel_netsyncQualityLevelST4 CiscoNetsyncQualityLevel = "netsyncQualityLevelST4"

    CiscoNetsyncQualityLevel_netsyncQualityLevelSTU CiscoNetsyncQualityLevel = "netsyncQualityLevelSTU"

    CiscoNetsyncQualityLevel_netsyncQualityLevelTNC CiscoNetsyncQualityLevel = "netsyncQualityLevelTNC"

    CiscoNetsyncQualityLevel_netsyncQualityLevelUNC CiscoNetsyncQualityLevel = "netsyncQualityLevelUNC"

    CiscoNetsyncQualityLevel_netsyncQualityLevelUNK CiscoNetsyncQualityLevel = "netsyncQualityLevelUNK"
)

// CiscoNetsyncClockMode represents netsyncClockModeLocked   - a valid clock source is locked.
type CiscoNetsyncClockMode string

const (
    CiscoNetsyncClockMode_netsyncClockModeUnknown CiscoNetsyncClockMode = "netsyncClockModeUnknown"

    CiscoNetsyncClockMode_netsyncClockModeFreerun CiscoNetsyncClockMode = "netsyncClockModeFreerun"

    CiscoNetsyncClockMode_netsyncClockModeHoldover CiscoNetsyncClockMode = "netsyncClockModeHoldover"

    CiscoNetsyncClockMode_netsyncClockModeLocked CiscoNetsyncClockMode = "netsyncClockModeLocked"
)

// CiscoNetsyncSSMCap represents netsyncSSMCapInvalid   Capability invalid or unsupported
type CiscoNetsyncSSMCap string

const (
    CiscoNetsyncSSMCap_netsyncSSMCapNone CiscoNetsyncSSMCap = "netsyncSSMCapNone"

    CiscoNetsyncSSMCap_netsyncSSMCapTxRx CiscoNetsyncSSMCap = "netsyncSSMCapTxRx"

    CiscoNetsyncSSMCap_netsyncSSMCapTx CiscoNetsyncSSMCap = "netsyncSSMCapTx"

    CiscoNetsyncSSMCap_netsyncSSMCapRx CiscoNetsyncSSMCap = "netsyncSSMCapRx"

    CiscoNetsyncSSMCap_netsyncSSMCapInvalid CiscoNetsyncSSMCap = "netsyncSSMCapInvalid"
)

// CiscoNetsyncIfType represents              type are T1 SF and T1 ESF.
type CiscoNetsyncIfType string

const (
    CiscoNetsyncIfType_netsyncIfTypeUnknown CiscoNetsyncIfType = "netsyncIfTypeUnknown"

    CiscoNetsyncIfType_netsyncIfTypeInternal CiscoNetsyncIfType = "netsyncIfTypeInternal"

    CiscoNetsyncIfType_netsyncIfTypeEthernet CiscoNetsyncIfType = "netsyncIfTypeEthernet"

    CiscoNetsyncIfType_netsyncIfTypeSonet CiscoNetsyncIfType = "netsyncIfTypeSonet"

    CiscoNetsyncIfType_netsyncIfTypeTop CiscoNetsyncIfType = "netsyncIfTypeTop"

    CiscoNetsyncIfType_netsyncIfTypeExt CiscoNetsyncIfType = "netsyncIfTypeExt"

    CiscoNetsyncIfType_netsyncIfTypeController CiscoNetsyncIfType = "netsyncIfTypeController"

    CiscoNetsyncIfType_netsyncIfTypeGps CiscoNetsyncIfType = "netsyncIfTypeGps"

    CiscoNetsyncIfType_netsyncIfTypeAtm CiscoNetsyncIfType = "netsyncIfTypeAtm"
)

// CiscoNetsyncNetworkOption represents designed for Japan.
type CiscoNetsyncNetworkOption string

const (
    CiscoNetsyncNetworkOption_netsyncNetworkOptionUnknown CiscoNetsyncNetworkOption = "netsyncNetworkOptionUnknown"

    CiscoNetsyncNetworkOption_netsyncNetworkOption1 CiscoNetsyncNetworkOption = "netsyncNetworkOption1"

    CiscoNetsyncNetworkOption_netsyncNetworkOption2Gen1 CiscoNetsyncNetworkOption = "netsyncNetworkOption2Gen1"

    CiscoNetsyncNetworkOption_netsyncNetworkOption2Gen2 CiscoNetsyncNetworkOption = "netsyncNetworkOption2Gen2"

    CiscoNetsyncNetworkOption_netsyncNetworkOption3 CiscoNetsyncNetworkOption = "netsyncNetworkOption3"

    CiscoNetsyncNetworkOption_netsyncNetworkOptionInvalid CiscoNetsyncNetworkOption = "netsyncNetworkOptionInvalid"
)

// CiscoNetsyncEECOption represents netsyncEECOptionInvalid   Invalid EEC option
type CiscoNetsyncEECOption string

const (
    CiscoNetsyncEECOption_netsyncEECOptionUnknown CiscoNetsyncEECOption = "netsyncEECOptionUnknown"

    CiscoNetsyncEECOption_netsyncEECOption1 CiscoNetsyncEECOption = "netsyncEECOption1"

    CiscoNetsyncEECOption_netsyncEECOption2 CiscoNetsyncEECOption = "netsyncEECOption2"

    CiscoNetsyncEECOption_netsyncEECOptionInvalid CiscoNetsyncEECOption = "netsyncEECOptionInvalid"
)

// CISCONETSYNCMIB
type CISCONETSYNCMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cisconetsyncmibnotifcontrol CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol

    // G.781 clock selection process table. This table contains the global
    // parameters for the G.781 clock selection process.
    Cnsclkselglobaltable CISCONETSYNCMIB_Cnsclkselglobaltable

    // T0 selected clock source table. This table contains the selected clock
    // source for the input T0 clock.
    Cnsselectedinputsourcetable CISCONETSYNCMIB_Cnsselectedinputsourcetable

    // T0 clock source table. This table contains a list of input sources for
    // input T0 clock selection.
    Cnsinputsourcetable CISCONETSYNCMIB_Cnsinputsourcetable

    // T4 external output table. This table contains a list of T4 external
    // outputs.  Each T4 external output is associated with clock source(s) to be
    // found in cnsT4ClockSourceTable. The clock selection process considers all
    // the available clock sources and select the T4 clock source based on the
    // G.781 clock selection algorithm.
    Cnsextoutputtable CISCONETSYNCMIB_Cnsextoutputtable

    // T4 clock source table. This table contains a list of input sources for a
    // specific T4 external output. An entry shall be added to cnsExtOutputTable
    // first. Then clock sources shall be added in this table for the selection
    // process to select the appropriate T4 clock source.
    Cnst4Clocksourcetable CISCONETSYNCMIB_Cnst4Clocksourcetable
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetFilter() yfilter.YFilter { return cISCONETSYNCMIB.YFilter }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) SetFilter(yf yfilter.YFilter) { cISCONETSYNCMIB.YFilter = yf }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetGoName(yname string) string {
    if yname == "ciscoNetsyncMIBNotifControl" { return "Cisconetsyncmibnotifcontrol" }
    if yname == "cnsClkSelGlobalTable" { return "Cnsclkselglobaltable" }
    if yname == "cnsSelectedInputSourceTable" { return "Cnsselectedinputsourcetable" }
    if yname == "cnsInputSourceTable" { return "Cnsinputsourcetable" }
    if yname == "cnsExtOutputTable" { return "Cnsextoutputtable" }
    if yname == "cnsT4ClockSourceTable" { return "Cnst4Clocksourcetable" }
    return ""
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetSegmentPath() string {
    return "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB"
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoNetsyncMIBNotifControl" {
        return &cISCONETSYNCMIB.Cisconetsyncmibnotifcontrol
    }
    if childYangName == "cnsClkSelGlobalTable" {
        return &cISCONETSYNCMIB.Cnsclkselglobaltable
    }
    if childYangName == "cnsSelectedInputSourceTable" {
        return &cISCONETSYNCMIB.Cnsselectedinputsourcetable
    }
    if childYangName == "cnsInputSourceTable" {
        return &cISCONETSYNCMIB.Cnsinputsourcetable
    }
    if childYangName == "cnsExtOutputTable" {
        return &cISCONETSYNCMIB.Cnsextoutputtable
    }
    if childYangName == "cnsT4ClockSourceTable" {
        return &cISCONETSYNCMIB.Cnst4Clocksourcetable
    }
    return nil
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoNetsyncMIBNotifControl"] = &cISCONETSYNCMIB.Cisconetsyncmibnotifcontrol
    children["cnsClkSelGlobalTable"] = &cISCONETSYNCMIB.Cnsclkselglobaltable
    children["cnsSelectedInputSourceTable"] = &cISCONETSYNCMIB.Cnsselectedinputsourcetable
    children["cnsInputSourceTable"] = &cISCONETSYNCMIB.Cnsinputsourcetable
    children["cnsExtOutputTable"] = &cISCONETSYNCMIB.Cnsextoutputtable
    children["cnsT4ClockSourceTable"] = &cISCONETSYNCMIB.Cnst4Clocksourcetable
    return children
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetYangName() string { return "CISCO-NETSYNC-MIB" }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) SetParent(parent types.Entity) { cISCONETSYNCMIB.parent = parent }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetParent() types.Entity { return cISCONETSYNCMIB.parent }

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol
type CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A control object to enable/disable ciscoNetsyncSelectedT0Clock,
    // ciscoNetsyncSelectedT4Clock, ciscoNetsyncInputSignalFailureStatus,
    // ciscoNetsyncInputAlarmStatus notifications at the system level. This object
    // should hold any of the below values.     true - The notif is enabled
    // globally for the system     false- The notif is disabled globally for the
    // system. The type is bool.
    Cnsmibenablestatusnotification interface{}
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetFilter() yfilter.YFilter { return cisconetsyncmibnotifcontrol.YFilter }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) SetFilter(yf yfilter.YFilter) { cisconetsyncmibnotifcontrol.YFilter = yf }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetGoName(yname string) string {
    if yname == "cnsMIBEnableStatusNotification" { return "Cnsmibenablestatusnotification" }
    return ""
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetSegmentPath() string {
    return "ciscoNetsyncMIBNotifControl"
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnsMIBEnableStatusNotification"] = cisconetsyncmibnotifcontrol.Cnsmibenablestatusnotification
    return leafs
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetBundleName() string { return "cisco_ios_xe" }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetYangName() string { return "ciscoNetsyncMIBNotifControl" }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) SetParent(parent types.Entity) { cisconetsyncmibnotifcontrol.parent = parent }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetParent() types.Entity { return cisconetsyncmibnotifcontrol.parent }

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cnsclkselglobaltable
// G.781 clock selection process table.
// This table contains the global parameters for the G.781 clock
// selection process.
type CISCONETSYNCMIB_Cnsclkselglobaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is added to cnsClkSelGlobalTable when G.781 clock selection is
    // enabled in the device configuration.  The entry is removed when G.781 clock
    // selection is removed from the configuration. The type is slice of
    // CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry.
    Cnsclkselglobalentry []CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetFilter() yfilter.YFilter { return cnsclkselglobaltable.YFilter }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) SetFilter(yf yfilter.YFilter) { cnsclkselglobaltable.YFilter = yf }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetGoName(yname string) string {
    if yname == "cnsClkSelGlobalEntry" { return "Cnsclkselglobalentry" }
    return ""
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetSegmentPath() string {
    return "cnsClkSelGlobalTable"
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnsClkSelGlobalEntry" {
        for _, c := range cnsclkselglobaltable.Cnsclkselglobalentry {
            if cnsclkselglobaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry{}
        cnsclkselglobaltable.Cnsclkselglobalentry = append(cnsclkselglobaltable.Cnsclkselglobalentry, child)
        return &cnsclkselglobaltable.Cnsclkselglobalentry[len(cnsclkselglobaltable.Cnsclkselglobalentry)-1]
    }
    return nil
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnsclkselglobaltable.Cnsclkselglobalentry {
        children[cnsclkselglobaltable.Cnsclkselglobalentry[i].GetSegmentPath()] = &cnsclkselglobaltable.Cnsclkselglobalentry[i]
    }
    return children
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetYangName() string { return "cnsClkSelGlobalTable" }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) SetParent(parent types.Entity) { cnsclkselglobaltable.parent = parent }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetParent() types.Entity { return cnsclkselglobaltable.parent }

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry
// An entry is added to cnsClkSelGlobalTable when G.781 clock
// selection is enabled in the device configuration.  The entry
// is removed when G.781 clock selection is removed from the
// configuration.
type CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely represents a clock
    // selection process.  This index is assigned arbitrarily by the system and
    // may not be persistent across reboots. The type is interface{} with range:
    // 0..4294967295.
    Cnsclkselgloprocindex interface{}

    // This object indicates the QL mode of the network synchronization clock
    // selection process as described in ITU-T standard G.781 section 5.12. The
    // type is CiscoNetsyncQLMode.
    Cnsclkselglobprocessmode interface{}

    // This object indicates the operating mode of the system clock. The type is
    // CiscoNetsyncClockMode.
    Cnsclkselglobclockmode interface{}

    // This object indicates whether the G.781 clock selection is enabled or not. 
    // 'true'  - G.781 clock selection is enabled 'false' - G.781 clock selection
    // is disabled. The type is bool.
    Cnsclkselglobnetsyncenable interface{}

    // This object indicates the revertive mode setting in the G.781 clock
    // selection process.  The switching of clock sources can be made revertive or
    // non-revertive. In non-revertive mode, an alternate clock source is
    // maintained even after the original clock source has recovered from the
    // failure that caused the switch. In revertive mode, the clock selection
    // process switches back to the original clock source after recovering from
    // the failure. The type is bool.
    Cnsclkselglobrevertivemode interface{}

    // This object indicates if global ESMC is enabled. With ESMC enabled
    // globally, the system is capable of handling ESMC messages. The type is
    // bool.
    Cnsclkselglobesmcmode interface{}

    // This object indicates the network synchronization EEC (Ethernet Equipment
    // Clock) option. The type is CiscoNetsyncEECOption.
    Cnsclkselglobeecoption interface{}

    // This object indicates the synchronization network option. The type is
    // CiscoNetsyncNetworkOption.
    Cnsclkselglobnetworkoption interface{}

    // This object indicates the global holdoff time in the G.781 clock selection
    // process. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cnsclkselglobholdofftime interface{}

    // This object indicates the global wait-to-restore time in the G.781 clock
    // selection process. The type is interface{} with range: 0..4294967295. Units
    // are seconds.
    Cnsclkselglobwtrtime interface{}

    // This object indicates the number of synchronization sources currently
    // configured for the G.781 clock selection process. The type is interface{}
    // with range: 0..255. Units are clock sources.
    Cnsclkselglobnofsources interface{}

    // This object indicates the duration of the last holdover period in seconds.
    // If the holdover duration is less than a second, the object will carry the
    // value zero. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Cnsclkselgloblastholdoverseconds interface{}

    // This object indicates the duration of the current holdover period. If a
    // system clock is in holdover mode, the object carries the current holdover
    // duration in seconds. If the system clock is not in holdover, the object
    // carries the value 0. The type is interface{} with range: 0..4294967295.
    // Units are seconds.
    Cnsclkselglobcurrholdoverseconds interface{}
}

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetFilter() yfilter.YFilter { return cnsclkselglobalentry.YFilter }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) SetFilter(yf yfilter.YFilter) { cnsclkselglobalentry.YFilter = yf }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetGoName(yname string) string {
    if yname == "cnsClkSelGloProcIndex" { return "Cnsclkselgloprocindex" }
    if yname == "cnsClkSelGlobProcessMode" { return "Cnsclkselglobprocessmode" }
    if yname == "cnsClkSelGlobClockMode" { return "Cnsclkselglobclockmode" }
    if yname == "cnsClkSelGlobNetsyncEnable" { return "Cnsclkselglobnetsyncenable" }
    if yname == "cnsClkSelGlobRevertiveMode" { return "Cnsclkselglobrevertivemode" }
    if yname == "cnsClkSelGlobESMCMode" { return "Cnsclkselglobesmcmode" }
    if yname == "cnsClkSelGlobEECOption" { return "Cnsclkselglobeecoption" }
    if yname == "cnsClkSelGlobNetworkOption" { return "Cnsclkselglobnetworkoption" }
    if yname == "cnsClkSelGlobHoldoffTime" { return "Cnsclkselglobholdofftime" }
    if yname == "cnsClkSelGlobWtrTime" { return "Cnsclkselglobwtrtime" }
    if yname == "cnsClkSelGlobNofSources" { return "Cnsclkselglobnofsources" }
    if yname == "cnsClkSelGlobLastHoldoverSeconds" { return "Cnsclkselgloblastholdoverseconds" }
    if yname == "cnsClkSelGlobCurrHoldoverSeconds" { return "Cnsclkselglobcurrholdoverseconds" }
    return ""
}

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetSegmentPath() string {
    return "cnsClkSelGlobalEntry" + "[cnsClkSelGloProcIndex='" + fmt.Sprintf("%v", cnsclkselglobalentry.Cnsclkselgloprocindex) + "']"
}

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnsClkSelGloProcIndex"] = cnsclkselglobalentry.Cnsclkselgloprocindex
    leafs["cnsClkSelGlobProcessMode"] = cnsclkselglobalentry.Cnsclkselglobprocessmode
    leafs["cnsClkSelGlobClockMode"] = cnsclkselglobalentry.Cnsclkselglobclockmode
    leafs["cnsClkSelGlobNetsyncEnable"] = cnsclkselglobalentry.Cnsclkselglobnetsyncenable
    leafs["cnsClkSelGlobRevertiveMode"] = cnsclkselglobalentry.Cnsclkselglobrevertivemode
    leafs["cnsClkSelGlobESMCMode"] = cnsclkselglobalentry.Cnsclkselglobesmcmode
    leafs["cnsClkSelGlobEECOption"] = cnsclkselglobalentry.Cnsclkselglobeecoption
    leafs["cnsClkSelGlobNetworkOption"] = cnsclkselglobalentry.Cnsclkselglobnetworkoption
    leafs["cnsClkSelGlobHoldoffTime"] = cnsclkselglobalentry.Cnsclkselglobholdofftime
    leafs["cnsClkSelGlobWtrTime"] = cnsclkselglobalentry.Cnsclkselglobwtrtime
    leafs["cnsClkSelGlobNofSources"] = cnsclkselglobalentry.Cnsclkselglobnofsources
    leafs["cnsClkSelGlobLastHoldoverSeconds"] = cnsclkselglobalentry.Cnsclkselgloblastholdoverseconds
    leafs["cnsClkSelGlobCurrHoldoverSeconds"] = cnsclkselglobalentry.Cnsclkselglobcurrholdoverseconds
    return leafs
}

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetYangName() string { return "cnsClkSelGlobalEntry" }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) SetParent(parent types.Entity) { cnsclkselglobalentry.parent = parent }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetParent() types.Entity { return cnsclkselglobalentry.parent }

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetParentYangName() string { return "cnsClkSelGlobalTable" }

// CISCONETSYNCMIB_Cnsselectedinputsourcetable
// T0 selected clock source table.
// This table contains the selected clock source for the input T0
// clock.
type CISCONETSYNCMIB_Cnsselectedinputsourcetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is created in the table when the G.781 clock selection process has
    // successfully selected a T0 clock source.  The entry shall remain during the
    // time the G.781 clock selection process remains enabled. The type is slice
    // of CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry.
    Cnsselectedinputsourceentry []CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetFilter() yfilter.YFilter { return cnsselectedinputsourcetable.YFilter }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) SetFilter(yf yfilter.YFilter) { cnsselectedinputsourcetable.YFilter = yf }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetGoName(yname string) string {
    if yname == "cnsSelectedInputSourceEntry" { return "Cnsselectedinputsourceentry" }
    return ""
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetSegmentPath() string {
    return "cnsSelectedInputSourceTable"
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnsSelectedInputSourceEntry" {
        for _, c := range cnsselectedinputsourcetable.Cnsselectedinputsourceentry {
            if cnsselectedinputsourcetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry{}
        cnsselectedinputsourcetable.Cnsselectedinputsourceentry = append(cnsselectedinputsourcetable.Cnsselectedinputsourceentry, child)
        return &cnsselectedinputsourcetable.Cnsselectedinputsourceentry[len(cnsselectedinputsourcetable.Cnsselectedinputsourceentry)-1]
    }
    return nil
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnsselectedinputsourcetable.Cnsselectedinputsourceentry {
        children[cnsselectedinputsourcetable.Cnsselectedinputsourceentry[i].GetSegmentPath()] = &cnsselectedinputsourcetable.Cnsselectedinputsourceentry[i]
    }
    return children
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetBundleName() string { return "cisco_ios_xe" }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetYangName() string { return "cnsSelectedInputSourceTable" }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) SetParent(parent types.Entity) { cnsselectedinputsourcetable.parent = parent }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetParent() types.Entity { return cnsselectedinputsourcetable.parent }

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry
// An entry is created in the table when the G.781 clock
// selection process has successfully selected a T0 clock
// source.  The entry shall remain during the time
// the G.781 clock selection process remains enabled.
type CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely represents an entry in this
    // table. This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    Cnsselinpsrcnetsyncindex interface{}

    // This object indicates the name of the selected T0 clock. The type is string
    // with length: 1..255.
    Cnsselinpsrcname interface{}

    // This object indicates the type of the selected T0 clock. The type is
    // CiscoNetsyncIfType.
    Cnsselinpsrcintftype interface{}

    // This object indicates the selected T0 clock source's effective quality
    // level, which is the derived clock quality based on the three factors:  (a)
    // Received quality level.  (b) Configured Rx quality level.  This factor
    // supersedes (a).  (c) System overridden quality level as a result of
    // exceptional events such as signal failure or ESMC failure.  This factor
    // supersedes (a) and (b). The type is CiscoNetsyncQualityLevel.
    Cnsselinpsrcqualitylevel interface{}

    // This object indicates the configured priority of the selected T0 clock. A
    // smaller value represents a higher priority. The type is interface{} with
    // range: 1..1024.
    Cnsselinpsrcpriority interface{}

    // This object indicates the timestamp of the T0 clock source being selected
    // by the G.781 clock selection process. The type is interface{} with range:
    // 0..4294967295.
    Cnsselinpsrctimestamp interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source.  The 'true' value indicates the currently selected
    // clock source is a result of the forced switching. The 'false' value
    // indicates the currently selected clock source is not a result of forced
    // switching. The type is bool.
    Cnsselinpsrcfsw interface{}

    // This object indicates the manual switching flag. The 'true' value indicates
    // the currently selected clock source is a result of the manual switch
    // command. The command allows a user to select a synchronization source
    // assuming it is enabled, not locked out, not in signal fail condition, and
    // has a QL better than DNU in QL-enabled mode. Furthermore, in QL-enabled
    // mode, a manual switch can be performed only to a source which has the
    // highest available QL. The type is bool.
    Cnsselinpsrcmsw interface{}
}

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetFilter() yfilter.YFilter { return cnsselectedinputsourceentry.YFilter }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) SetFilter(yf yfilter.YFilter) { cnsselectedinputsourceentry.YFilter = yf }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetGoName(yname string) string {
    if yname == "cnsSelInpSrcNetsyncIndex" { return "Cnsselinpsrcnetsyncindex" }
    if yname == "cnsSelInpSrcName" { return "Cnsselinpsrcname" }
    if yname == "cnsSelInpSrcIntfType" { return "Cnsselinpsrcintftype" }
    if yname == "cnsSelInpSrcQualityLevel" { return "Cnsselinpsrcqualitylevel" }
    if yname == "cnsSelInpSrcPriority" { return "Cnsselinpsrcpriority" }
    if yname == "cnsSelInpSrcTimestamp" { return "Cnsselinpsrctimestamp" }
    if yname == "cnsSelInpSrcFSW" { return "Cnsselinpsrcfsw" }
    if yname == "cnsSelInpSrcMSW" { return "Cnsselinpsrcmsw" }
    return ""
}

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetSegmentPath() string {
    return "cnsSelectedInputSourceEntry" + "[cnsSelInpSrcNetsyncIndex='" + fmt.Sprintf("%v", cnsselectedinputsourceentry.Cnsselinpsrcnetsyncindex) + "']"
}

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnsSelInpSrcNetsyncIndex"] = cnsselectedinputsourceentry.Cnsselinpsrcnetsyncindex
    leafs["cnsSelInpSrcName"] = cnsselectedinputsourceentry.Cnsselinpsrcname
    leafs["cnsSelInpSrcIntfType"] = cnsselectedinputsourceentry.Cnsselinpsrcintftype
    leafs["cnsSelInpSrcQualityLevel"] = cnsselectedinputsourceentry.Cnsselinpsrcqualitylevel
    leafs["cnsSelInpSrcPriority"] = cnsselectedinputsourceentry.Cnsselinpsrcpriority
    leafs["cnsSelInpSrcTimestamp"] = cnsselectedinputsourceentry.Cnsselinpsrctimestamp
    leafs["cnsSelInpSrcFSW"] = cnsselectedinputsourceentry.Cnsselinpsrcfsw
    leafs["cnsSelInpSrcMSW"] = cnsselectedinputsourceentry.Cnsselinpsrcmsw
    return leafs
}

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetYangName() string { return "cnsSelectedInputSourceEntry" }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) SetParent(parent types.Entity) { cnsselectedinputsourceentry.parent = parent }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetParent() types.Entity { return cnsselectedinputsourceentry.parent }

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetParentYangName() string { return "cnsSelectedInputSourceTable" }

// CISCONETSYNCMIB_Cnsinputsourcetable
// T0 clock source table.
// This table contains a list of input sources for input T0 clock
// selection.
type CISCONETSYNCMIB_Cnsinputsourcetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a T0 clock source in the
    // configuration. An entry is removed  in the table when a user removes a T0
    // clock source from the configuration. The type is slice of
    // CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry.
    Cnsinputsourceentry []CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetFilter() yfilter.YFilter { return cnsinputsourcetable.YFilter }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) SetFilter(yf yfilter.YFilter) { cnsinputsourcetable.YFilter = yf }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetGoName(yname string) string {
    if yname == "cnsInputSourceEntry" { return "Cnsinputsourceentry" }
    return ""
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetSegmentPath() string {
    return "cnsInputSourceTable"
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnsInputSourceEntry" {
        for _, c := range cnsinputsourcetable.Cnsinputsourceentry {
            if cnsinputsourcetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry{}
        cnsinputsourcetable.Cnsinputsourceentry = append(cnsinputsourcetable.Cnsinputsourceentry, child)
        return &cnsinputsourcetable.Cnsinputsourceentry[len(cnsinputsourcetable.Cnsinputsourceentry)-1]
    }
    return nil
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnsinputsourcetable.Cnsinputsourceentry {
        children[cnsinputsourcetable.Cnsinputsourceentry[i].GetSegmentPath()] = &cnsinputsourcetable.Cnsinputsourceentry[i]
    }
    return children
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetBundleName() string { return "cisco_ios_xe" }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetYangName() string { return "cnsInputSourceTable" }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) SetParent(parent types.Entity) { cnsinputsourcetable.parent = parent }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetParent() types.Entity { return cnsinputsourcetable.parent }

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry
// An entry is created in the table when a user adds a T0
// clock source in the configuration. An entry is removed 
// in the table when a user removes a T0 clock source from
// the configuration.
type CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely represents an entry in this
    // table. This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    Cnsinpsrcnetsyncindex interface{}

    // This object indicates the name of an input clock source configured for the
    // T0 clock selection. The type is string with length: 1..255.
    Cnsinpsrcname interface{}

    // This object indicates the type of an input clock source configured for the
    // T0 clock selection. The type is CiscoNetsyncIfType.
    Cnsinpsrcintftype interface{}

    // This object indicates the priority of an input clock source configured for
    // the T0 clock selection.  A smaller value represents a higher priority. The
    // type is interface{} with range: 1..1024.
    Cnsinpsrcpriority interface{}

    // This object indicates the ESMC capability of an input clock source
    // configured for the T0 clock selection.  This is applicable only to
    // Synchronous Ethernet input clock source identified by cnsInpSrcIntfType
    // 'netsyncIfTypeEthernet'. The type is CiscoNetsyncESMCCap.
    Cnsinpsrcesmccap interface{}

    // This object indicates the SSM capability of an input clock source
    // configured for the T0 clock selection. This is applicable only to any
    // synchronous interface clock source except SyncE interface, which is
    // identified by cnsInpSrcIntfType 'netsyncIfTypeEthernet'. The type is
    // CiscoNetsyncSSMCap.
    Cnsinpsrcssmcap interface{}

    // This object indicates the configured transmit clock quality level of an
    // input clock source. The type is CiscoNetsyncQualityLevel.
    Cnsinpsrcqualityleveltxcfg interface{}

    // This object indicates the configured receive clock quality level of an
    // input clock source. The type is CiscoNetsyncQualityLevel.
    Cnsinpsrcqualitylevelrxcfg interface{}

    // This object indicates the most recent clock quality level transmitted on
    // the input clock source. The type is CiscoNetsyncQualityLevel.
    Cnsinpsrcqualityleveltx interface{}

    // This object indicates the last clock quality level received on the input
    // clock source. The type is CiscoNetsyncQualityLevel.
    Cnsinpsrcqualitylevelrx interface{}

    // This object indicates the current clock quality level of the input clock
    // source.  This is the effective quality which is derived from three values: 
    // 1) most recent clock quality level received, 2) forced clock quality level
    // (entered via configuration) 3) overridden clock quality level as a result
    // of line protocol down, signal failure, or alarms. The type is
    // CiscoNetsyncQualityLevel.
    Cnsinpsrcqualitylevel interface{}

    // This object indicates the hold-off time value of an input clock source. 
    // The hold-off time prevents short activation of signal failure is passed to
    // the selection process.  When a signal failure event is reported on a clock
    // source, it waits the duration of the hold-off time before declaring signal
    // failure on the clock source. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cnsinpsrcholdofftime interface{}

    // This object indicates the wait-to-restore time value of an input clock
    // source.  The wait-to-restore time ensures that a previous failed
    // synchronization source is only again considered as available by the
    // selection process if it is fault-free for a certain time. When a signal
    // failure condition is cleared on a clock source, it waits the duration of
    // the wait-to-restore time before clearing the signal failure status on the
    // clock source. The type is interface{} with range: 0..4294967295. Units are
    // Seconds.
    Cnsinpsrcwtrtime interface{}

    // This object indicates whether or not the lockout command has been applied
    // to a clock source.  The 'true' value means the clock source is not
    // considered by the selection process. The type is bool.
    Cnsinpsrclockout interface{}

    // This object indicates whether or not a signal failure event is currently
    // being reported on the input clock source. The type is bool.
    Cnsinpsrcsignalfailure interface{}

    // This object indicates whether or not an alarm event is currently being
    // reported on the input clock source. The type is bool.
    Cnsinpsrcalarm interface{}

    // This object indicates the alarm reasons of an input clock source if an
    // alarm event is being reported on it. The type is map[string]bool.
    Cnsinpsrcalarminfo interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source.  The 'true' value indicates the currently selected
    // clock source is a result of the forced switching. The 'false' value
    // indicates the currently selected clock source is not a result of forced
    // switching. The type is bool.
    Cnsinpsrcfsw interface{}

    // This object indicates the manual switching flag.  The 'true' value
    // indicates the currently selected clock source is a result of the manual
    // switching. The switch allows a user to select a synchronization source
    // assuming it is enabled, not locked out, not in signal fail condition, and
    // has a QL better than DNU in QL-enabled mode.  A clock source is enabled
    // when it occupies a row in cnsInputSourceTable.  A clock source is not
    // locked out when cnsInpSrcLockout contains the value 'false'. A clock source
    // is not in signal failure condition when cnsInpSrcSignalFailure contains the
    // value 'false'.  The QL is identified in cnsInpSrcQualityLevel.  In
    // QL-enabled mode, a manual switch can be performed only to a source which
    // has the highest available QL. The type is bool.
    Cnsinpsrcmsw interface{}
}

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetFilter() yfilter.YFilter { return cnsinputsourceentry.YFilter }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) SetFilter(yf yfilter.YFilter) { cnsinputsourceentry.YFilter = yf }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetGoName(yname string) string {
    if yname == "cnsInpSrcNetsyncIndex" { return "Cnsinpsrcnetsyncindex" }
    if yname == "cnsInpSrcName" { return "Cnsinpsrcname" }
    if yname == "cnsInpSrcIntfType" { return "Cnsinpsrcintftype" }
    if yname == "cnsInpSrcPriority" { return "Cnsinpsrcpriority" }
    if yname == "cnsInpSrcESMCCap" { return "Cnsinpsrcesmccap" }
    if yname == "cnsInpSrcSSMCap" { return "Cnsinpsrcssmcap" }
    if yname == "cnsInpSrcQualityLevelTxCfg" { return "Cnsinpsrcqualityleveltxcfg" }
    if yname == "cnsInpSrcQualityLevelRxCfg" { return "Cnsinpsrcqualitylevelrxcfg" }
    if yname == "cnsInpSrcQualityLevelTx" { return "Cnsinpsrcqualityleveltx" }
    if yname == "cnsInpSrcQualityLevelRx" { return "Cnsinpsrcqualitylevelrx" }
    if yname == "cnsInpSrcQualityLevel" { return "Cnsinpsrcqualitylevel" }
    if yname == "cnsInpSrcHoldoffTime" { return "Cnsinpsrcholdofftime" }
    if yname == "cnsInpSrcWtrTime" { return "Cnsinpsrcwtrtime" }
    if yname == "cnsInpSrcLockout" { return "Cnsinpsrclockout" }
    if yname == "cnsInpSrcSignalFailure" { return "Cnsinpsrcsignalfailure" }
    if yname == "cnsInpSrcAlarm" { return "Cnsinpsrcalarm" }
    if yname == "cnsInpSrcAlarmInfo" { return "Cnsinpsrcalarminfo" }
    if yname == "cnsInpSrcFSW" { return "Cnsinpsrcfsw" }
    if yname == "cnsInpSrcMSW" { return "Cnsinpsrcmsw" }
    return ""
}

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetSegmentPath() string {
    return "cnsInputSourceEntry" + "[cnsInpSrcNetsyncIndex='" + fmt.Sprintf("%v", cnsinputsourceentry.Cnsinpsrcnetsyncindex) + "']"
}

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnsInpSrcNetsyncIndex"] = cnsinputsourceentry.Cnsinpsrcnetsyncindex
    leafs["cnsInpSrcName"] = cnsinputsourceentry.Cnsinpsrcname
    leafs["cnsInpSrcIntfType"] = cnsinputsourceentry.Cnsinpsrcintftype
    leafs["cnsInpSrcPriority"] = cnsinputsourceentry.Cnsinpsrcpriority
    leafs["cnsInpSrcESMCCap"] = cnsinputsourceentry.Cnsinpsrcesmccap
    leafs["cnsInpSrcSSMCap"] = cnsinputsourceentry.Cnsinpsrcssmcap
    leafs["cnsInpSrcQualityLevelTxCfg"] = cnsinputsourceentry.Cnsinpsrcqualityleveltxcfg
    leafs["cnsInpSrcQualityLevelRxCfg"] = cnsinputsourceentry.Cnsinpsrcqualitylevelrxcfg
    leafs["cnsInpSrcQualityLevelTx"] = cnsinputsourceentry.Cnsinpsrcqualityleveltx
    leafs["cnsInpSrcQualityLevelRx"] = cnsinputsourceentry.Cnsinpsrcqualitylevelrx
    leafs["cnsInpSrcQualityLevel"] = cnsinputsourceentry.Cnsinpsrcqualitylevel
    leafs["cnsInpSrcHoldoffTime"] = cnsinputsourceentry.Cnsinpsrcholdofftime
    leafs["cnsInpSrcWtrTime"] = cnsinputsourceentry.Cnsinpsrcwtrtime
    leafs["cnsInpSrcLockout"] = cnsinputsourceentry.Cnsinpsrclockout
    leafs["cnsInpSrcSignalFailure"] = cnsinputsourceentry.Cnsinpsrcsignalfailure
    leafs["cnsInpSrcAlarm"] = cnsinputsourceentry.Cnsinpsrcalarm
    leafs["cnsInpSrcAlarmInfo"] = cnsinputsourceentry.Cnsinpsrcalarminfo
    leafs["cnsInpSrcFSW"] = cnsinputsourceentry.Cnsinpsrcfsw
    leafs["cnsInpSrcMSW"] = cnsinputsourceentry.Cnsinpsrcmsw
    return leafs
}

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetYangName() string { return "cnsInputSourceEntry" }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) SetParent(parent types.Entity) { cnsinputsourceentry.parent = parent }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetParent() types.Entity { return cnsinputsourceentry.parent }

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetParentYangName() string { return "cnsInputSourceTable" }

// CISCONETSYNCMIB_Cnsextoutputtable
// T4 external output table.
// This table contains a list of T4 external outputs.
// 
// Each T4 external output is associated with clock
// source(s) to be found in cnsT4ClockSourceTable.
// The clock selection process considers all the
// available clock sources and select the T4 clock
// source based on the G.781 clock selection algorithm.
type CISCONETSYNCMIB_Cnsextoutputtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a T4 external output in
    // the configuration.  A T4 external output configured input clock sources are
    // defined in cnsT4ClockSourceTable.  An entry is removed from the table when
    // a user removes a T4 external output from the configuration. The type is
    // slice of CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry.
    Cnsextoutputentry []CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetFilter() yfilter.YFilter { return cnsextoutputtable.YFilter }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) SetFilter(yf yfilter.YFilter) { cnsextoutputtable.YFilter = yf }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetGoName(yname string) string {
    if yname == "cnsExtOutputEntry" { return "Cnsextoutputentry" }
    return ""
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetSegmentPath() string {
    return "cnsExtOutputTable"
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnsExtOutputEntry" {
        for _, c := range cnsextoutputtable.Cnsextoutputentry {
            if cnsextoutputtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry{}
        cnsextoutputtable.Cnsextoutputentry = append(cnsextoutputtable.Cnsextoutputentry, child)
        return &cnsextoutputtable.Cnsextoutputentry[len(cnsextoutputtable.Cnsextoutputentry)-1]
    }
    return nil
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnsextoutputtable.Cnsextoutputentry {
        children[cnsextoutputtable.Cnsextoutputentry[i].GetSegmentPath()] = &cnsextoutputtable.Cnsextoutputentry[i]
    }
    return children
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetBundleName() string { return "cisco_ios_xe" }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetYangName() string { return "cnsExtOutputTable" }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) SetParent(parent types.Entity) { cnsextoutputtable.parent = parent }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetParent() types.Entity { return cnsextoutputtable.parent }

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry
// An entry is created in the table when a user adds
// a T4 external output in the configuration.  A T4 external
// output configured input clock sources are defined in
// cnsT4ClockSourceTable.
// 
// An entry is removed from the table when a user removes
// a T4 external output from the configuration.
type CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely represents an entry in this
    // table.  This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    Cnsextoutlistindex interface{}

    // An index that uniquely represents the selected input clock source whose
    // information is reported by a row in cnsT4ClockSourceTable. The index lists
    // the value of cnsT4ClkSrcNetsyncIndex, which is the input clock source of
    // the T4 external output selected by the G.781 clock selection process. The
    // type is interface{} with range: 1..4294967295.
    Cnsextoutselnetsyncindex interface{}

    // This object indicates the name of a T4 external output. The type is string
    // with length: 1..255.
    Cnsextoutname interface{}

    // This object indicates the interface type of the T4 external output. The
    // type is CiscoNetsyncIfType.
    Cnsextoutintftype interface{}

    // This object indicates the clock quality of the T4 external output. The type
    // is CiscoNetsyncQualityLevel.
    Cnsextoutqualitylevel interface{}

    // This object indicates the priority of the selected clock source for a T4
    // external output.  A smaller value represents a higher priority. The type is
    // interface{} with range: 1..1024.
    Cnsextoutpriority interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source, The T4 selected synchronization source is
    // identified by cnsExtOutSelNetsyncIndex, which contains the index to the
    // clock source in cnsT4ClockSourceTable.  The 'true' value indicates the
    // currently selected T4 clock source is a result of the forced switching. The
    // 'false' value indicates the currently selected T4 clock source is not a
    // result of forced switching. The type is bool.
    Cnsextoutfsw interface{}

    // This object indicates the manual switching flag.  The 'true' value
    // indicates the currently selected T4 clock source is a result of the manual
    // switch command. The command allows a user to select a synchronization
    // source assuming it is enabled, not locked out, not in signal fail
    // condition, and has a QL better than DNU in QL-enabled mode.  A clock source
    // is enabled when it occupies in row in cnsT4ClockSourceTable.  A clock
    // source is not locked out when cnsT4ClkSrcLockout contains the value
    // 'false'. A clock source is not in signal failure condition when
    // cnsT4ClkSrcSignalFailure contains the value 'false'.  The QL is identified
    // in cnsT4ClkSrcQualityLevel.  In QL-enabled mode, a manual switch can be 
    // performed only to a source which has the highest available QL. The type is
    // bool.
    Cnsextoutmsw interface{}

    // This object indicates whether or not a T4 external output is squelched. 
    // Squelching is a sychronization function defined to prevent transmission of
    // a timing signal with a quality that is lower than the quality of the clock
    // in the receiving networks element or SASE. It is also used for the
    // prevention of timing loops. The type is bool.
    Cnsextoutsquelch interface{}
}

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetFilter() yfilter.YFilter { return cnsextoutputentry.YFilter }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) SetFilter(yf yfilter.YFilter) { cnsextoutputentry.YFilter = yf }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetGoName(yname string) string {
    if yname == "cnsExtOutListIndex" { return "Cnsextoutlistindex" }
    if yname == "cnsExtOutSelNetsyncIndex" { return "Cnsextoutselnetsyncindex" }
    if yname == "cnsExtOutName" { return "Cnsextoutname" }
    if yname == "cnsExtOutIntfType" { return "Cnsextoutintftype" }
    if yname == "cnsExtOutQualityLevel" { return "Cnsextoutqualitylevel" }
    if yname == "cnsExtOutPriority" { return "Cnsextoutpriority" }
    if yname == "cnsExtOutFSW" { return "Cnsextoutfsw" }
    if yname == "cnsExtOutMSW" { return "Cnsextoutmsw" }
    if yname == "cnsExtOutSquelch" { return "Cnsextoutsquelch" }
    return ""
}

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetSegmentPath() string {
    return "cnsExtOutputEntry" + "[cnsExtOutListIndex='" + fmt.Sprintf("%v", cnsextoutputentry.Cnsextoutlistindex) + "']"
}

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnsExtOutListIndex"] = cnsextoutputentry.Cnsextoutlistindex
    leafs["cnsExtOutSelNetsyncIndex"] = cnsextoutputentry.Cnsextoutselnetsyncindex
    leafs["cnsExtOutName"] = cnsextoutputentry.Cnsextoutname
    leafs["cnsExtOutIntfType"] = cnsextoutputentry.Cnsextoutintftype
    leafs["cnsExtOutQualityLevel"] = cnsextoutputentry.Cnsextoutqualitylevel
    leafs["cnsExtOutPriority"] = cnsextoutputentry.Cnsextoutpriority
    leafs["cnsExtOutFSW"] = cnsextoutputentry.Cnsextoutfsw
    leafs["cnsExtOutMSW"] = cnsextoutputentry.Cnsextoutmsw
    leafs["cnsExtOutSquelch"] = cnsextoutputentry.Cnsextoutsquelch
    return leafs
}

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetYangName() string { return "cnsExtOutputEntry" }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) SetParent(parent types.Entity) { cnsextoutputentry.parent = parent }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetParent() types.Entity { return cnsextoutputentry.parent }

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetParentYangName() string { return "cnsExtOutputTable" }

// CISCONETSYNCMIB_Cnst4Clocksourcetable
// T4 clock source table.
// This table contains a list of input sources for a specific
// T4 external output. An entry shall be added to
// cnsExtOutputTable first. Then clock sources shall be
// added in this table for the selection process to select
// the appropriate T4 clock source.
type CISCONETSYNCMIB_Cnst4Clocksourcetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a clock source to a T4
    // external output in the configuration. The T4 external output is defined in
    // the T4 external output table. An entry is removed in the table when a user
    // removes a T4 clock source from the configuration. The type is slice of
    // CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry.
    Cnst4Clocksourceentry []CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetFilter() yfilter.YFilter { return cnst4Clocksourcetable.YFilter }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) SetFilter(yf yfilter.YFilter) { cnst4Clocksourcetable.YFilter = yf }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetGoName(yname string) string {
    if yname == "cnsT4ClockSourceEntry" { return "Cnst4Clocksourceentry" }
    return ""
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetSegmentPath() string {
    return "cnsT4ClockSourceTable"
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnsT4ClockSourceEntry" {
        for _, c := range cnst4Clocksourcetable.Cnst4Clocksourceentry {
            if cnst4Clocksourcetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry{}
        cnst4Clocksourcetable.Cnst4Clocksourceentry = append(cnst4Clocksourcetable.Cnst4Clocksourceentry, child)
        return &cnst4Clocksourcetable.Cnst4Clocksourceentry[len(cnst4Clocksourcetable.Cnst4Clocksourceentry)-1]
    }
    return nil
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnst4Clocksourcetable.Cnst4Clocksourceentry {
        children[cnst4Clocksourcetable.Cnst4Clocksourceentry[i].GetSegmentPath()] = &cnst4Clocksourcetable.Cnst4Clocksourceentry[i]
    }
    return children
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetBundleName() string { return "cisco_ios_xe" }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetYangName() string { return "cnsT4ClockSourceTable" }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) SetParent(parent types.Entity) { cnst4Clocksourcetable.parent = parent }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetParent() types.Entity { return cnst4Clocksourcetable.parent }

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetParentYangName() string { return "CISCO-NETSYNC-MIB" }

// CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry
// An entry is created in the table when a user adds a
// clock source to a T4 external output in the configuration.
// The T4 external output is defined in the T4 external
// output table. An entry is removed in the table when a user
// removes a T4 clock source from the configuration.
type CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_netsync_mib.CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry_Cnsextoutlistindex
    Cnsextoutlistindex interface{}

    // This attribute is a key. An index that uniquely represents an entry in this
    // table.  This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    Cnst4Clksrcnetsyncindex interface{}

    // This object indicates the name of a input clock source configured for the
    // T4 clock selection. The type is string with length: 1..255.
    Cnst4Clksrcname interface{}

    // This object indicates the type of an input clock source configured for the
    // T4 clock selection. The type is CiscoNetsyncIfType.
    Cnst4Clksrcintftype interface{}

    // This object indicates the priority of an input clock source configured for
    // the T4 clock selection.  A smaller value represents a higher priority. The
    // type is interface{} with range: 1..1024.
    Cnst4Clksrcpriority interface{}

    // This object indicates the ESMC capability of an input clock source
    // configured for the T4 clock selection.  This is applicable only to
    // Synchronous Ethernet input clock source identified by cnsT4ClkSrcIntfType
    // 'netsyncIfTypeEthernet'. The type is CiscoNetsyncESMCCap.
    Cnst4Clksrcesmccap interface{}

    // This object indicates the SSM capability of an input clock source
    // configured for the T4 clock selection. This is applicable only to any
    // synchronous interface clock source except SyncE interface, which is
    // identified by cnsT4ClkSrcIntfType 'netsyncIfTypeEthernet'. The type is
    // CiscoNetsyncSSMCap.
    Cnst4Clksrcssmcap interface{}

    // This object indicates the configured transmit clock quality level of a T4
    // input clock source. The type is CiscoNetsyncQualityLevel.
    Cnst4Clksrcqualityleveltxcfg interface{}

    // This object indicates the configured receive clock quality level of a T4
    // input clock source. The type is CiscoNetsyncQualityLevel.
    Cnst4Clksrcqualitylevelrxcfg interface{}

    // This object indicates the most recent clock quality level transmitted on
    // the T4 input clock source. The type is CiscoNetsyncQualityLevel.
    Cnst4Clksrcqualityleveltx interface{}

    // This object indicates the last clock quality level received on the T4 input
    // clock source. The type is CiscoNetsyncQualityLevel.
    Cnst4Clksrcqualitylevelrx interface{}

    // This object indicates the current clock quality level of the T4 input clock
    // source.  This is the effective quality which is derived from three values: 
    // 1) most recent clock quality level received, 2) forced clock quality level
    // (entered via configuration) 3) overridden clock quality level as a result
    // of line protocol down, signal failure, or alarms. The type is
    // CiscoNetsyncQualityLevel.
    Cnst4Clksrcqualitylevel interface{}

    // This object indicates the hold-off time value of a T4 input clock source. 
    // The hold-off time prevents short activation of signal failure is passed to
    // the selection process.  When a signal failure event is reported on a clock
    // source, it waits the duration of the hold-off time before declaring signal
    // failure on the clock source. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cnst4Clksrcholdofftime interface{}

    // This object indicates the wait-to-restore time value of a T4 input clock
    // source.  The wait-to-restore time ensures that a previous failed
    // synchronization source is only again considered as available by the
    // selection process if it is fault-free for a certain time. When a signal
    // failure condition is cleared on a clock source, it waits the duration of
    // the wait-to-restore time before clearing the signal failure status on the
    // clock source. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Cnst4Clksrcwtrtime interface{}

    // This object indicates whether or not the lockout command has been applied
    // on a T4 clock source.  The 'true' value means the clock source is not
    // considered by the selection process. The type is bool.
    Cnst4Clksrclockout interface{}

    // This object indicates whether or not a signal failure event is currently
    // being reported on the T4 input clock source. The type is bool.
    Cnst4Clksrcsignalfailure interface{}

    // This object indicates whether or not an alarm event is currently being
    // reported on the T4 input clock source. The type is bool.
    Cnst4Clksrcalarm interface{}

    // This object indicates the alarm reasons of a T4 input clock source if an
    // alarm event is being reported on the clock source. The type is
    // map[string]bool.
    Cnst4Clksrcalarminfo interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source.  The 'true' value indicates the currently selected
    // T4 clock source is a result of the forced switching. The 'false' value
    // indicates the currently selected T4 clock source is not a result of forced
    // switching. The type is bool.
    Cnst4Clksrcfsw interface{}

    // This object indicates the manual switching flag.  The 'true' value
    // indicates the currently selected T4 clock source is a result of the manual
    // switching. The switch allows a user to select a  synchronization source
    // assuming it is enabled, not locked out, not in signal fail condition, and
    // has a QL better than DNU in QL-enabled mode.  A clock source is enabled
    // when it occupies a row in  cnsT4ClockSourceTable.  A clock source is not
    // locked out when cnsT4ClkSrcLockout contains the value 'false'. A clock
    // source is not in signal failure condition when cnsT4ClkSrcSignalFailure
    // contains the value 'false'. The QL is identified in
    // cnsT4ClkSrcQualityLevel.  In QL-enabled mode, a manual switch can be
    // performed only to a source which has the highest available QL. The type is
    // bool.
    Cnst4Clksrcmsw interface{}
}

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetFilter() yfilter.YFilter { return cnst4Clocksourceentry.YFilter }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) SetFilter(yf yfilter.YFilter) { cnst4Clocksourceentry.YFilter = yf }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetGoName(yname string) string {
    if yname == "cnsExtOutListIndex" { return "Cnsextoutlistindex" }
    if yname == "cnsT4ClkSrcNetsyncIndex" { return "Cnst4Clksrcnetsyncindex" }
    if yname == "cnsT4ClkSrcName" { return "Cnst4Clksrcname" }
    if yname == "cnsT4ClkSrcIntfType" { return "Cnst4Clksrcintftype" }
    if yname == "cnsT4ClkSrcPriority" { return "Cnst4Clksrcpriority" }
    if yname == "cnsT4ClkSrcESMCCap" { return "Cnst4Clksrcesmccap" }
    if yname == "cnsT4ClkSrcSSMCap" { return "Cnst4Clksrcssmcap" }
    if yname == "cnsT4ClkSrcQualityLevelTxCfg" { return "Cnst4Clksrcqualityleveltxcfg" }
    if yname == "cnsT4ClkSrcQualityLevelRxCfg" { return "Cnst4Clksrcqualitylevelrxcfg" }
    if yname == "cnsT4ClkSrcQualityLevelTx" { return "Cnst4Clksrcqualityleveltx" }
    if yname == "cnsT4ClkSrcQualityLevelRx" { return "Cnst4Clksrcqualitylevelrx" }
    if yname == "cnsT4ClkSrcQualityLevel" { return "Cnst4Clksrcqualitylevel" }
    if yname == "cnsT4ClkSrcHoldoffTime" { return "Cnst4Clksrcholdofftime" }
    if yname == "cnsT4ClkSrcWtrTime" { return "Cnst4Clksrcwtrtime" }
    if yname == "cnsT4ClkSrcLockout" { return "Cnst4Clksrclockout" }
    if yname == "cnsT4ClkSrcSignalFailure" { return "Cnst4Clksrcsignalfailure" }
    if yname == "cnsT4ClkSrcAlarm" { return "Cnst4Clksrcalarm" }
    if yname == "cnsT4ClkSrcAlarmInfo" { return "Cnst4Clksrcalarminfo" }
    if yname == "cnsT4ClkSrcFSW" { return "Cnst4Clksrcfsw" }
    if yname == "cnsT4ClkSrcMSW" { return "Cnst4Clksrcmsw" }
    return ""
}

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetSegmentPath() string {
    return "cnsT4ClockSourceEntry" + "[cnsExtOutListIndex='" + fmt.Sprintf("%v", cnst4Clocksourceentry.Cnsextoutlistindex) + "']" + "[cnsT4ClkSrcNetsyncIndex='" + fmt.Sprintf("%v", cnst4Clocksourceentry.Cnst4Clksrcnetsyncindex) + "']"
}

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnsExtOutListIndex"] = cnst4Clocksourceentry.Cnsextoutlistindex
    leafs["cnsT4ClkSrcNetsyncIndex"] = cnst4Clocksourceentry.Cnst4Clksrcnetsyncindex
    leafs["cnsT4ClkSrcName"] = cnst4Clocksourceentry.Cnst4Clksrcname
    leafs["cnsT4ClkSrcIntfType"] = cnst4Clocksourceentry.Cnst4Clksrcintftype
    leafs["cnsT4ClkSrcPriority"] = cnst4Clocksourceentry.Cnst4Clksrcpriority
    leafs["cnsT4ClkSrcESMCCap"] = cnst4Clocksourceentry.Cnst4Clksrcesmccap
    leafs["cnsT4ClkSrcSSMCap"] = cnst4Clocksourceentry.Cnst4Clksrcssmcap
    leafs["cnsT4ClkSrcQualityLevelTxCfg"] = cnst4Clocksourceentry.Cnst4Clksrcqualityleveltxcfg
    leafs["cnsT4ClkSrcQualityLevelRxCfg"] = cnst4Clocksourceentry.Cnst4Clksrcqualitylevelrxcfg
    leafs["cnsT4ClkSrcQualityLevelTx"] = cnst4Clocksourceentry.Cnst4Clksrcqualityleveltx
    leafs["cnsT4ClkSrcQualityLevelRx"] = cnst4Clocksourceentry.Cnst4Clksrcqualitylevelrx
    leafs["cnsT4ClkSrcQualityLevel"] = cnst4Clocksourceentry.Cnst4Clksrcqualitylevel
    leafs["cnsT4ClkSrcHoldoffTime"] = cnst4Clocksourceentry.Cnst4Clksrcholdofftime
    leafs["cnsT4ClkSrcWtrTime"] = cnst4Clocksourceentry.Cnst4Clksrcwtrtime
    leafs["cnsT4ClkSrcLockout"] = cnst4Clocksourceentry.Cnst4Clksrclockout
    leafs["cnsT4ClkSrcSignalFailure"] = cnst4Clocksourceentry.Cnst4Clksrcsignalfailure
    leafs["cnsT4ClkSrcAlarm"] = cnst4Clocksourceentry.Cnst4Clksrcalarm
    leafs["cnsT4ClkSrcAlarmInfo"] = cnst4Clocksourceentry.Cnst4Clksrcalarminfo
    leafs["cnsT4ClkSrcFSW"] = cnst4Clocksourceentry.Cnst4Clksrcfsw
    leafs["cnsT4ClkSrcMSW"] = cnst4Clocksourceentry.Cnst4Clksrcmsw
    return leafs
}

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetYangName() string { return "cnsT4ClockSourceEntry" }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) SetParent(parent types.Entity) { cnst4Clocksourceentry.parent = parent }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetParent() types.Entity { return cnst4Clocksourceentry.parent }

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetParentYangName() string { return "cnsT4ClockSourceTable" }

