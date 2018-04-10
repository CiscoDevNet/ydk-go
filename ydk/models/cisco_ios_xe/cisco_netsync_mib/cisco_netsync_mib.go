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

// CiscoNetsyncQLMode represents                           clock selection criteria.
type CiscoNetsyncQLMode string

const (
    CiscoNetsyncQLMode_netsyncQLModeUnknown CiscoNetsyncQLMode = "netsyncQLModeUnknown"

    CiscoNetsyncQLMode_netsyncQLModeQlDisabled CiscoNetsyncQLMode = "netsyncQLModeQlDisabled"

    CiscoNetsyncQLMode_netsyncQLModeQlEnabled CiscoNetsyncQLMode = "netsyncQLModeQlEnabled"
)

// CiscoNetsyncClockMode represents netsyncClockModeLocked   - a valid clock source is locked.
type CiscoNetsyncClockMode string

const (
    CiscoNetsyncClockMode_netsyncClockModeUnknown CiscoNetsyncClockMode = "netsyncClockModeUnknown"

    CiscoNetsyncClockMode_netsyncClockModeFreerun CiscoNetsyncClockMode = "netsyncClockModeFreerun"

    CiscoNetsyncClockMode_netsyncClockModeHoldover CiscoNetsyncClockMode = "netsyncClockModeHoldover"

    CiscoNetsyncClockMode_netsyncClockModeLocked CiscoNetsyncClockMode = "netsyncClockModeLocked"
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

// CiscoNetsyncSSMCap represents netsyncSSMCapInvalid   Capability invalid or unsupported
type CiscoNetsyncSSMCap string

const (
    CiscoNetsyncSSMCap_netsyncSSMCapNone CiscoNetsyncSSMCap = "netsyncSSMCapNone"

    CiscoNetsyncSSMCap_netsyncSSMCapTxRx CiscoNetsyncSSMCap = "netsyncSSMCapTxRx"

    CiscoNetsyncSSMCap_netsyncSSMCapTx CiscoNetsyncSSMCap = "netsyncSSMCapTx"

    CiscoNetsyncSSMCap_netsyncSSMCapRx CiscoNetsyncSSMCap = "netsyncSSMCapRx"

    CiscoNetsyncSSMCap_netsyncSSMCapInvalid CiscoNetsyncSSMCap = "netsyncSSMCapInvalid"
)

// CiscoNetsyncESMCCap represents netsyncESMCCapInvalid   Capability invalid or unsupported
type CiscoNetsyncESMCCap string

const (
    CiscoNetsyncESMCCap_netsyncESMCCapNone CiscoNetsyncESMCCap = "netsyncESMCCapNone"

    CiscoNetsyncESMCCap_netsyncESMCCapTxRx CiscoNetsyncESMCCap = "netsyncESMCCapTxRx"

    CiscoNetsyncESMCCap_netsyncESMCCapTx CiscoNetsyncESMCCap = "netsyncESMCCapTx"

    CiscoNetsyncESMCCap_netsyncESMCCapRx CiscoNetsyncESMCCap = "netsyncESMCCapRx"

    CiscoNetsyncESMCCap_netsyncESMCCapInvalid CiscoNetsyncESMCCap = "netsyncESMCCapInvalid"
)

// CISCONETSYNCMIB
type CISCONETSYNCMIB struct {
    EntityData types.CommonEntityData
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

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetEntityData() *types.CommonEntityData {
    cISCONETSYNCMIB.EntityData.YFilter = cISCONETSYNCMIB.YFilter
    cISCONETSYNCMIB.EntityData.YangName = "CISCO-NETSYNC-MIB"
    cISCONETSYNCMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCONETSYNCMIB.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cISCONETSYNCMIB.EntityData.SegmentPath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB"
    cISCONETSYNCMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCONETSYNCMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCONETSYNCMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCONETSYNCMIB.EntityData.Children = make(map[string]types.YChild)
    cISCONETSYNCMIB.EntityData.Children["ciscoNetsyncMIBNotifControl"] = types.YChild{"Cisconetsyncmibnotifcontrol", &cISCONETSYNCMIB.Cisconetsyncmibnotifcontrol}
    cISCONETSYNCMIB.EntityData.Children["cnsClkSelGlobalTable"] = types.YChild{"Cnsclkselglobaltable", &cISCONETSYNCMIB.Cnsclkselglobaltable}
    cISCONETSYNCMIB.EntityData.Children["cnsSelectedInputSourceTable"] = types.YChild{"Cnsselectedinputsourcetable", &cISCONETSYNCMIB.Cnsselectedinputsourcetable}
    cISCONETSYNCMIB.EntityData.Children["cnsInputSourceTable"] = types.YChild{"Cnsinputsourcetable", &cISCONETSYNCMIB.Cnsinputsourcetable}
    cISCONETSYNCMIB.EntityData.Children["cnsExtOutputTable"] = types.YChild{"Cnsextoutputtable", &cISCONETSYNCMIB.Cnsextoutputtable}
    cISCONETSYNCMIB.EntityData.Children["cnsT4ClockSourceTable"] = types.YChild{"Cnst4Clocksourcetable", &cISCONETSYNCMIB.Cnst4Clocksourcetable}
    cISCONETSYNCMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCONETSYNCMIB.EntityData)
}

// CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol
type CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A control object to enable/disable ciscoNetsyncSelectedT0Clock,
    // ciscoNetsyncSelectedT4Clock, ciscoNetsyncInputSignalFailureStatus,
    // ciscoNetsyncInputAlarmStatus notifications at the system level. This object
    // should hold any of the below values.     true - The notif is enabled
    // globally for the system     false- The notif is disabled globally for the
    // system. The type is bool.
    Cnsmibenablestatusnotification interface{}
}

func (cisconetsyncmibnotifcontrol *CISCONETSYNCMIB_Cisconetsyncmibnotifcontrol) GetEntityData() *types.CommonEntityData {
    cisconetsyncmibnotifcontrol.EntityData.YFilter = cisconetsyncmibnotifcontrol.YFilter
    cisconetsyncmibnotifcontrol.EntityData.YangName = "ciscoNetsyncMIBNotifControl"
    cisconetsyncmibnotifcontrol.EntityData.BundleName = "cisco_ios_xe"
    cisconetsyncmibnotifcontrol.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cisconetsyncmibnotifcontrol.EntityData.SegmentPath = "ciscoNetsyncMIBNotifControl"
    cisconetsyncmibnotifcontrol.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cisconetsyncmibnotifcontrol.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cisconetsyncmibnotifcontrol.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cisconetsyncmibnotifcontrol.EntityData.Children = make(map[string]types.YChild)
    cisconetsyncmibnotifcontrol.EntityData.Leafs = make(map[string]types.YLeaf)
    cisconetsyncmibnotifcontrol.EntityData.Leafs["cnsMIBEnableStatusNotification"] = types.YLeaf{"Cnsmibenablestatusnotification", cisconetsyncmibnotifcontrol.Cnsmibenablestatusnotification}
    return &(cisconetsyncmibnotifcontrol.EntityData)
}

// CISCONETSYNCMIB_Cnsclkselglobaltable
// G.781 clock selection process table.
// This table contains the global parameters for the G.781 clock
// selection process.
type CISCONETSYNCMIB_Cnsclkselglobaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is added to cnsClkSelGlobalTable when G.781 clock selection is
    // enabled in the device configuration.  The entry is removed when G.781 clock
    // selection is removed from the configuration. The type is slice of
    // CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry.
    Cnsclkselglobalentry []CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry
}

func (cnsclkselglobaltable *CISCONETSYNCMIB_Cnsclkselglobaltable) GetEntityData() *types.CommonEntityData {
    cnsclkselglobaltable.EntityData.YFilter = cnsclkselglobaltable.YFilter
    cnsclkselglobaltable.EntityData.YangName = "cnsClkSelGlobalTable"
    cnsclkselglobaltable.EntityData.BundleName = "cisco_ios_xe"
    cnsclkselglobaltable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsclkselglobaltable.EntityData.SegmentPath = "cnsClkSelGlobalTable"
    cnsclkselglobaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsclkselglobaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsclkselglobaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsclkselglobaltable.EntityData.Children = make(map[string]types.YChild)
    cnsclkselglobaltable.EntityData.Children["cnsClkSelGlobalEntry"] = types.YChild{"Cnsclkselglobalentry", nil}
    for i := range cnsclkselglobaltable.Cnsclkselglobalentry {
        cnsclkselglobaltable.EntityData.Children[types.GetSegmentPath(&cnsclkselglobaltable.Cnsclkselglobalentry[i])] = types.YChild{"Cnsclkselglobalentry", &cnsclkselglobaltable.Cnsclkselglobalentry[i]}
    }
    cnsclkselglobaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnsclkselglobaltable.EntityData)
}

// CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry
// An entry is added to cnsClkSelGlobalTable when G.781 clock
// selection is enabled in the device configuration.  The entry
// is removed when G.781 clock selection is removed from the
// configuration.
type CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry struct {
    EntityData types.CommonEntityData
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

func (cnsclkselglobalentry *CISCONETSYNCMIB_Cnsclkselglobaltable_Cnsclkselglobalentry) GetEntityData() *types.CommonEntityData {
    cnsclkselglobalentry.EntityData.YFilter = cnsclkselglobalentry.YFilter
    cnsclkselglobalentry.EntityData.YangName = "cnsClkSelGlobalEntry"
    cnsclkselglobalentry.EntityData.BundleName = "cisco_ios_xe"
    cnsclkselglobalentry.EntityData.ParentYangName = "cnsClkSelGlobalTable"
    cnsclkselglobalentry.EntityData.SegmentPath = "cnsClkSelGlobalEntry" + "[cnsClkSelGloProcIndex='" + fmt.Sprintf("%v", cnsclkselglobalentry.Cnsclkselgloprocindex) + "']"
    cnsclkselglobalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsclkselglobalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsclkselglobalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsclkselglobalentry.EntityData.Children = make(map[string]types.YChild)
    cnsclkselglobalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGloProcIndex"] = types.YLeaf{"Cnsclkselgloprocindex", cnsclkselglobalentry.Cnsclkselgloprocindex}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobProcessMode"] = types.YLeaf{"Cnsclkselglobprocessmode", cnsclkselglobalentry.Cnsclkselglobprocessmode}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobClockMode"] = types.YLeaf{"Cnsclkselglobclockmode", cnsclkselglobalentry.Cnsclkselglobclockmode}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobNetsyncEnable"] = types.YLeaf{"Cnsclkselglobnetsyncenable", cnsclkselglobalentry.Cnsclkselglobnetsyncenable}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobRevertiveMode"] = types.YLeaf{"Cnsclkselglobrevertivemode", cnsclkselglobalentry.Cnsclkselglobrevertivemode}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobESMCMode"] = types.YLeaf{"Cnsclkselglobesmcmode", cnsclkselglobalentry.Cnsclkselglobesmcmode}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobEECOption"] = types.YLeaf{"Cnsclkselglobeecoption", cnsclkselglobalentry.Cnsclkselglobeecoption}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobNetworkOption"] = types.YLeaf{"Cnsclkselglobnetworkoption", cnsclkselglobalentry.Cnsclkselglobnetworkoption}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobHoldoffTime"] = types.YLeaf{"Cnsclkselglobholdofftime", cnsclkselglobalentry.Cnsclkselglobholdofftime}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobWtrTime"] = types.YLeaf{"Cnsclkselglobwtrtime", cnsclkselglobalentry.Cnsclkselglobwtrtime}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobNofSources"] = types.YLeaf{"Cnsclkselglobnofsources", cnsclkselglobalentry.Cnsclkselglobnofsources}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobLastHoldoverSeconds"] = types.YLeaf{"Cnsclkselgloblastholdoverseconds", cnsclkselglobalentry.Cnsclkselgloblastholdoverseconds}
    cnsclkselglobalentry.EntityData.Leafs["cnsClkSelGlobCurrHoldoverSeconds"] = types.YLeaf{"Cnsclkselglobcurrholdoverseconds", cnsclkselglobalentry.Cnsclkselglobcurrholdoverseconds}
    return &(cnsclkselglobalentry.EntityData)
}

// CISCONETSYNCMIB_Cnsselectedinputsourcetable
// T0 selected clock source table.
// This table contains the selected clock source for the input T0
// clock.
type CISCONETSYNCMIB_Cnsselectedinputsourcetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when the G.781 clock selection process has
    // successfully selected a T0 clock source.  The entry shall remain during the
    // time the G.781 clock selection process remains enabled. The type is slice
    // of CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry.
    Cnsselectedinputsourceentry []CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry
}

func (cnsselectedinputsourcetable *CISCONETSYNCMIB_Cnsselectedinputsourcetable) GetEntityData() *types.CommonEntityData {
    cnsselectedinputsourcetable.EntityData.YFilter = cnsselectedinputsourcetable.YFilter
    cnsselectedinputsourcetable.EntityData.YangName = "cnsSelectedInputSourceTable"
    cnsselectedinputsourcetable.EntityData.BundleName = "cisco_ios_xe"
    cnsselectedinputsourcetable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsselectedinputsourcetable.EntityData.SegmentPath = "cnsSelectedInputSourceTable"
    cnsselectedinputsourcetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsselectedinputsourcetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsselectedinputsourcetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsselectedinputsourcetable.EntityData.Children = make(map[string]types.YChild)
    cnsselectedinputsourcetable.EntityData.Children["cnsSelectedInputSourceEntry"] = types.YChild{"Cnsselectedinputsourceentry", nil}
    for i := range cnsselectedinputsourcetable.Cnsselectedinputsourceentry {
        cnsselectedinputsourcetable.EntityData.Children[types.GetSegmentPath(&cnsselectedinputsourcetable.Cnsselectedinputsourceentry[i])] = types.YChild{"Cnsselectedinputsourceentry", &cnsselectedinputsourcetable.Cnsselectedinputsourceentry[i]}
    }
    cnsselectedinputsourcetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnsselectedinputsourcetable.EntityData)
}

// CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry
// An entry is created in the table when the G.781 clock
// selection process has successfully selected a T0 clock
// source.  The entry shall remain during the time
// the G.781 clock selection process remains enabled.
type CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry struct {
    EntityData types.CommonEntityData
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

func (cnsselectedinputsourceentry *CISCONETSYNCMIB_Cnsselectedinputsourcetable_Cnsselectedinputsourceentry) GetEntityData() *types.CommonEntityData {
    cnsselectedinputsourceentry.EntityData.YFilter = cnsselectedinputsourceentry.YFilter
    cnsselectedinputsourceentry.EntityData.YangName = "cnsSelectedInputSourceEntry"
    cnsselectedinputsourceentry.EntityData.BundleName = "cisco_ios_xe"
    cnsselectedinputsourceentry.EntityData.ParentYangName = "cnsSelectedInputSourceTable"
    cnsselectedinputsourceentry.EntityData.SegmentPath = "cnsSelectedInputSourceEntry" + "[cnsSelInpSrcNetsyncIndex='" + fmt.Sprintf("%v", cnsselectedinputsourceentry.Cnsselinpsrcnetsyncindex) + "']"
    cnsselectedinputsourceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsselectedinputsourceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsselectedinputsourceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsselectedinputsourceentry.EntityData.Children = make(map[string]types.YChild)
    cnsselectedinputsourceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcNetsyncIndex"] = types.YLeaf{"Cnsselinpsrcnetsyncindex", cnsselectedinputsourceentry.Cnsselinpsrcnetsyncindex}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcName"] = types.YLeaf{"Cnsselinpsrcname", cnsselectedinputsourceentry.Cnsselinpsrcname}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcIntfType"] = types.YLeaf{"Cnsselinpsrcintftype", cnsselectedinputsourceentry.Cnsselinpsrcintftype}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcQualityLevel"] = types.YLeaf{"Cnsselinpsrcqualitylevel", cnsselectedinputsourceentry.Cnsselinpsrcqualitylevel}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcPriority"] = types.YLeaf{"Cnsselinpsrcpriority", cnsselectedinputsourceentry.Cnsselinpsrcpriority}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcTimestamp"] = types.YLeaf{"Cnsselinpsrctimestamp", cnsselectedinputsourceentry.Cnsselinpsrctimestamp}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcFSW"] = types.YLeaf{"Cnsselinpsrcfsw", cnsselectedinputsourceentry.Cnsselinpsrcfsw}
    cnsselectedinputsourceentry.EntityData.Leafs["cnsSelInpSrcMSW"] = types.YLeaf{"Cnsselinpsrcmsw", cnsselectedinputsourceentry.Cnsselinpsrcmsw}
    return &(cnsselectedinputsourceentry.EntityData)
}

// CISCONETSYNCMIB_Cnsinputsourcetable
// T0 clock source table.
// This table contains a list of input sources for input T0 clock
// selection.
type CISCONETSYNCMIB_Cnsinputsourcetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a T0 clock source in the
    // configuration. An entry is removed  in the table when a user removes a T0
    // clock source from the configuration. The type is slice of
    // CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry.
    Cnsinputsourceentry []CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry
}

func (cnsinputsourcetable *CISCONETSYNCMIB_Cnsinputsourcetable) GetEntityData() *types.CommonEntityData {
    cnsinputsourcetable.EntityData.YFilter = cnsinputsourcetable.YFilter
    cnsinputsourcetable.EntityData.YangName = "cnsInputSourceTable"
    cnsinputsourcetable.EntityData.BundleName = "cisco_ios_xe"
    cnsinputsourcetable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsinputsourcetable.EntityData.SegmentPath = "cnsInputSourceTable"
    cnsinputsourcetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsinputsourcetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsinputsourcetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsinputsourcetable.EntityData.Children = make(map[string]types.YChild)
    cnsinputsourcetable.EntityData.Children["cnsInputSourceEntry"] = types.YChild{"Cnsinputsourceentry", nil}
    for i := range cnsinputsourcetable.Cnsinputsourceentry {
        cnsinputsourcetable.EntityData.Children[types.GetSegmentPath(&cnsinputsourcetable.Cnsinputsourceentry[i])] = types.YChild{"Cnsinputsourceentry", &cnsinputsourcetable.Cnsinputsourceentry[i]}
    }
    cnsinputsourcetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnsinputsourcetable.EntityData)
}

// CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry
// An entry is created in the table when a user adds a T0
// clock source in the configuration. An entry is removed 
// in the table when a user removes a T0 clock source from
// the configuration.
type CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry struct {
    EntityData types.CommonEntityData
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

func (cnsinputsourceentry *CISCONETSYNCMIB_Cnsinputsourcetable_Cnsinputsourceentry) GetEntityData() *types.CommonEntityData {
    cnsinputsourceentry.EntityData.YFilter = cnsinputsourceentry.YFilter
    cnsinputsourceentry.EntityData.YangName = "cnsInputSourceEntry"
    cnsinputsourceentry.EntityData.BundleName = "cisco_ios_xe"
    cnsinputsourceentry.EntityData.ParentYangName = "cnsInputSourceTable"
    cnsinputsourceentry.EntityData.SegmentPath = "cnsInputSourceEntry" + "[cnsInpSrcNetsyncIndex='" + fmt.Sprintf("%v", cnsinputsourceentry.Cnsinpsrcnetsyncindex) + "']"
    cnsinputsourceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsinputsourceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsinputsourceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsinputsourceentry.EntityData.Children = make(map[string]types.YChild)
    cnsinputsourceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcNetsyncIndex"] = types.YLeaf{"Cnsinpsrcnetsyncindex", cnsinputsourceentry.Cnsinpsrcnetsyncindex}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcName"] = types.YLeaf{"Cnsinpsrcname", cnsinputsourceentry.Cnsinpsrcname}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcIntfType"] = types.YLeaf{"Cnsinpsrcintftype", cnsinputsourceentry.Cnsinpsrcintftype}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcPriority"] = types.YLeaf{"Cnsinpsrcpriority", cnsinputsourceentry.Cnsinpsrcpriority}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcESMCCap"] = types.YLeaf{"Cnsinpsrcesmccap", cnsinputsourceentry.Cnsinpsrcesmccap}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcSSMCap"] = types.YLeaf{"Cnsinpsrcssmcap", cnsinputsourceentry.Cnsinpsrcssmcap}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcQualityLevelTxCfg"] = types.YLeaf{"Cnsinpsrcqualityleveltxcfg", cnsinputsourceentry.Cnsinpsrcqualityleveltxcfg}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcQualityLevelRxCfg"] = types.YLeaf{"Cnsinpsrcqualitylevelrxcfg", cnsinputsourceentry.Cnsinpsrcqualitylevelrxcfg}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcQualityLevelTx"] = types.YLeaf{"Cnsinpsrcqualityleveltx", cnsinputsourceentry.Cnsinpsrcqualityleveltx}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcQualityLevelRx"] = types.YLeaf{"Cnsinpsrcqualitylevelrx", cnsinputsourceentry.Cnsinpsrcqualitylevelrx}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcQualityLevel"] = types.YLeaf{"Cnsinpsrcqualitylevel", cnsinputsourceentry.Cnsinpsrcqualitylevel}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcHoldoffTime"] = types.YLeaf{"Cnsinpsrcholdofftime", cnsinputsourceentry.Cnsinpsrcholdofftime}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcWtrTime"] = types.YLeaf{"Cnsinpsrcwtrtime", cnsinputsourceentry.Cnsinpsrcwtrtime}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcLockout"] = types.YLeaf{"Cnsinpsrclockout", cnsinputsourceentry.Cnsinpsrclockout}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcSignalFailure"] = types.YLeaf{"Cnsinpsrcsignalfailure", cnsinputsourceentry.Cnsinpsrcsignalfailure}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcAlarm"] = types.YLeaf{"Cnsinpsrcalarm", cnsinputsourceentry.Cnsinpsrcalarm}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcAlarmInfo"] = types.YLeaf{"Cnsinpsrcalarminfo", cnsinputsourceentry.Cnsinpsrcalarminfo}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcFSW"] = types.YLeaf{"Cnsinpsrcfsw", cnsinputsourceentry.Cnsinpsrcfsw}
    cnsinputsourceentry.EntityData.Leafs["cnsInpSrcMSW"] = types.YLeaf{"Cnsinpsrcmsw", cnsinputsourceentry.Cnsinpsrcmsw}
    return &(cnsinputsourceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a T4 external output in
    // the configuration.  A T4 external output configured input clock sources are
    // defined in cnsT4ClockSourceTable.  An entry is removed from the table when
    // a user removes a T4 external output from the configuration. The type is
    // slice of CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry.
    Cnsextoutputentry []CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry
}

func (cnsextoutputtable *CISCONETSYNCMIB_Cnsextoutputtable) GetEntityData() *types.CommonEntityData {
    cnsextoutputtable.EntityData.YFilter = cnsextoutputtable.YFilter
    cnsextoutputtable.EntityData.YangName = "cnsExtOutputTable"
    cnsextoutputtable.EntityData.BundleName = "cisco_ios_xe"
    cnsextoutputtable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsextoutputtable.EntityData.SegmentPath = "cnsExtOutputTable"
    cnsextoutputtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsextoutputtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsextoutputtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsextoutputtable.EntityData.Children = make(map[string]types.YChild)
    cnsextoutputtable.EntityData.Children["cnsExtOutputEntry"] = types.YChild{"Cnsextoutputentry", nil}
    for i := range cnsextoutputtable.Cnsextoutputentry {
        cnsextoutputtable.EntityData.Children[types.GetSegmentPath(&cnsextoutputtable.Cnsextoutputentry[i])] = types.YChild{"Cnsextoutputentry", &cnsextoutputtable.Cnsextoutputentry[i]}
    }
    cnsextoutputtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnsextoutputtable.EntityData)
}

// CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry
// An entry is created in the table when a user adds
// a T4 external output in the configuration.  A T4 external
// output configured input clock sources are defined in
// cnsT4ClockSourceTable.
// 
// An entry is removed from the table when a user removes
// a T4 external output from the configuration.
type CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry struct {
    EntityData types.CommonEntityData
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

func (cnsextoutputentry *CISCONETSYNCMIB_Cnsextoutputtable_Cnsextoutputentry) GetEntityData() *types.CommonEntityData {
    cnsextoutputentry.EntityData.YFilter = cnsextoutputentry.YFilter
    cnsextoutputentry.EntityData.YangName = "cnsExtOutputEntry"
    cnsextoutputentry.EntityData.BundleName = "cisco_ios_xe"
    cnsextoutputentry.EntityData.ParentYangName = "cnsExtOutputTable"
    cnsextoutputentry.EntityData.SegmentPath = "cnsExtOutputEntry" + "[cnsExtOutListIndex='" + fmt.Sprintf("%v", cnsextoutputentry.Cnsextoutlistindex) + "']"
    cnsextoutputentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsextoutputentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsextoutputentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsextoutputentry.EntityData.Children = make(map[string]types.YChild)
    cnsextoutputentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnsextoutputentry.EntityData.Leafs["cnsExtOutListIndex"] = types.YLeaf{"Cnsextoutlistindex", cnsextoutputentry.Cnsextoutlistindex}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutSelNetsyncIndex"] = types.YLeaf{"Cnsextoutselnetsyncindex", cnsextoutputentry.Cnsextoutselnetsyncindex}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutName"] = types.YLeaf{"Cnsextoutname", cnsextoutputentry.Cnsextoutname}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutIntfType"] = types.YLeaf{"Cnsextoutintftype", cnsextoutputentry.Cnsextoutintftype}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutQualityLevel"] = types.YLeaf{"Cnsextoutqualitylevel", cnsextoutputentry.Cnsextoutqualitylevel}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutPriority"] = types.YLeaf{"Cnsextoutpriority", cnsextoutputentry.Cnsextoutpriority}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutFSW"] = types.YLeaf{"Cnsextoutfsw", cnsextoutputentry.Cnsextoutfsw}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutMSW"] = types.YLeaf{"Cnsextoutmsw", cnsextoutputentry.Cnsextoutmsw}
    cnsextoutputentry.EntityData.Leafs["cnsExtOutSquelch"] = types.YLeaf{"Cnsextoutsquelch", cnsextoutputentry.Cnsextoutsquelch}
    return &(cnsextoutputentry.EntityData)
}

// CISCONETSYNCMIB_Cnst4Clocksourcetable
// T4 clock source table.
// This table contains a list of input sources for a specific
// T4 external output. An entry shall be added to
// cnsExtOutputTable first. Then clock sources shall be
// added in this table for the selection process to select
// the appropriate T4 clock source.
type CISCONETSYNCMIB_Cnst4Clocksourcetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a clock source to a T4
    // external output in the configuration. The T4 external output is defined in
    // the T4 external output table. An entry is removed in the table when a user
    // removes a T4 clock source from the configuration. The type is slice of
    // CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry.
    Cnst4Clocksourceentry []CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry
}

func (cnst4Clocksourcetable *CISCONETSYNCMIB_Cnst4Clocksourcetable) GetEntityData() *types.CommonEntityData {
    cnst4Clocksourcetable.EntityData.YFilter = cnst4Clocksourcetable.YFilter
    cnst4Clocksourcetable.EntityData.YangName = "cnsT4ClockSourceTable"
    cnst4Clocksourcetable.EntityData.BundleName = "cisco_ios_xe"
    cnst4Clocksourcetable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnst4Clocksourcetable.EntityData.SegmentPath = "cnsT4ClockSourceTable"
    cnst4Clocksourcetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnst4Clocksourcetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnst4Clocksourcetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnst4Clocksourcetable.EntityData.Children = make(map[string]types.YChild)
    cnst4Clocksourcetable.EntityData.Children["cnsT4ClockSourceEntry"] = types.YChild{"Cnst4Clocksourceentry", nil}
    for i := range cnst4Clocksourcetable.Cnst4Clocksourceentry {
        cnst4Clocksourcetable.EntityData.Children[types.GetSegmentPath(&cnst4Clocksourcetable.Cnst4Clocksourceentry[i])] = types.YChild{"Cnst4Clocksourceentry", &cnst4Clocksourcetable.Cnst4Clocksourceentry[i]}
    }
    cnst4Clocksourcetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnst4Clocksourcetable.EntityData)
}

// CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry
// An entry is created in the table when a user adds a
// clock source to a T4 external output in the configuration.
// The T4 external output is defined in the T4 external
// output table. An entry is removed in the table when a user
// removes a T4 clock source from the configuration.
type CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry struct {
    EntityData types.CommonEntityData
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

func (cnst4Clocksourceentry *CISCONETSYNCMIB_Cnst4Clocksourcetable_Cnst4Clocksourceentry) GetEntityData() *types.CommonEntityData {
    cnst4Clocksourceentry.EntityData.YFilter = cnst4Clocksourceentry.YFilter
    cnst4Clocksourceentry.EntityData.YangName = "cnsT4ClockSourceEntry"
    cnst4Clocksourceentry.EntityData.BundleName = "cisco_ios_xe"
    cnst4Clocksourceentry.EntityData.ParentYangName = "cnsT4ClockSourceTable"
    cnst4Clocksourceentry.EntityData.SegmentPath = "cnsT4ClockSourceEntry" + "[cnsExtOutListIndex='" + fmt.Sprintf("%v", cnst4Clocksourceentry.Cnsextoutlistindex) + "']" + "[cnsT4ClkSrcNetsyncIndex='" + fmt.Sprintf("%v", cnst4Clocksourceentry.Cnst4Clksrcnetsyncindex) + "']"
    cnst4Clocksourceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnst4Clocksourceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnst4Clocksourceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnst4Clocksourceentry.EntityData.Children = make(map[string]types.YChild)
    cnst4Clocksourceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnst4Clocksourceentry.EntityData.Leafs["cnsExtOutListIndex"] = types.YLeaf{"Cnsextoutlistindex", cnst4Clocksourceentry.Cnsextoutlistindex}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcNetsyncIndex"] = types.YLeaf{"Cnst4Clksrcnetsyncindex", cnst4Clocksourceentry.Cnst4Clksrcnetsyncindex}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcName"] = types.YLeaf{"Cnst4Clksrcname", cnst4Clocksourceentry.Cnst4Clksrcname}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcIntfType"] = types.YLeaf{"Cnst4Clksrcintftype", cnst4Clocksourceentry.Cnst4Clksrcintftype}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcPriority"] = types.YLeaf{"Cnst4Clksrcpriority", cnst4Clocksourceentry.Cnst4Clksrcpriority}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcESMCCap"] = types.YLeaf{"Cnst4Clksrcesmccap", cnst4Clocksourceentry.Cnst4Clksrcesmccap}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcSSMCap"] = types.YLeaf{"Cnst4Clksrcssmcap", cnst4Clocksourceentry.Cnst4Clksrcssmcap}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcQualityLevelTxCfg"] = types.YLeaf{"Cnst4Clksrcqualityleveltxcfg", cnst4Clocksourceentry.Cnst4Clksrcqualityleveltxcfg}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcQualityLevelRxCfg"] = types.YLeaf{"Cnst4Clksrcqualitylevelrxcfg", cnst4Clocksourceentry.Cnst4Clksrcqualitylevelrxcfg}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcQualityLevelTx"] = types.YLeaf{"Cnst4Clksrcqualityleveltx", cnst4Clocksourceentry.Cnst4Clksrcqualityleveltx}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcQualityLevelRx"] = types.YLeaf{"Cnst4Clksrcqualitylevelrx", cnst4Clocksourceentry.Cnst4Clksrcqualitylevelrx}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcQualityLevel"] = types.YLeaf{"Cnst4Clksrcqualitylevel", cnst4Clocksourceentry.Cnst4Clksrcqualitylevel}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcHoldoffTime"] = types.YLeaf{"Cnst4Clksrcholdofftime", cnst4Clocksourceentry.Cnst4Clksrcholdofftime}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcWtrTime"] = types.YLeaf{"Cnst4Clksrcwtrtime", cnst4Clocksourceentry.Cnst4Clksrcwtrtime}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcLockout"] = types.YLeaf{"Cnst4Clksrclockout", cnst4Clocksourceentry.Cnst4Clksrclockout}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcSignalFailure"] = types.YLeaf{"Cnst4Clksrcsignalfailure", cnst4Clocksourceentry.Cnst4Clksrcsignalfailure}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcAlarm"] = types.YLeaf{"Cnst4Clksrcalarm", cnst4Clocksourceentry.Cnst4Clksrcalarm}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcAlarmInfo"] = types.YLeaf{"Cnst4Clksrcalarminfo", cnst4Clocksourceentry.Cnst4Clksrcalarminfo}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcFSW"] = types.YLeaf{"Cnst4Clksrcfsw", cnst4Clocksourceentry.Cnst4Clksrcfsw}
    cnst4Clocksourceentry.EntityData.Leafs["cnsT4ClkSrcMSW"] = types.YLeaf{"Cnst4Clksrcmsw", cnst4Clocksourceentry.Cnst4Clksrcmsw}
    return &(cnst4Clocksourceentry.EntityData)
}

