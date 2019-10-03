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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoNetsyncMIBNotifControl CISCONETSYNCMIB_CiscoNetsyncMIBNotifControl

    // G.781 clock selection process table. This table contains the global
    // parameters for the G.781 clock selection process.
    CnsClkSelGlobalTable CISCONETSYNCMIB_CnsClkSelGlobalTable

    // T0 selected clock source table. This table contains the selected clock
    // source for the input T0 clock.
    CnsSelectedInputSourceTable CISCONETSYNCMIB_CnsSelectedInputSourceTable

    // T0 clock source table. This table contains a list of input sources for
    // input T0 clock selection.
    CnsInputSourceTable CISCONETSYNCMIB_CnsInputSourceTable

    // T4 external output table. This table contains a list of T4 external
    // outputs.  Each T4 external output is associated with clock source(s) to be
    // found in cnsT4ClockSourceTable. The clock selection process considers all
    // the available clock sources and select the T4 clock source based on the
    // G.781 clock selection algorithm.
    CnsExtOutputTable CISCONETSYNCMIB_CnsExtOutputTable

    // T4 clock source table. This table contains a list of input sources for a
    // specific T4 external output. An entry shall be added to cnsExtOutputTable
    // first. Then clock sources shall be added in this table for the selection
    // process to select the appropriate T4 clock source.
    CnsT4ClockSourceTable CISCONETSYNCMIB_CnsT4ClockSourceTable
}

func (cISCONETSYNCMIB *CISCONETSYNCMIB) GetEntityData() *types.CommonEntityData {
    cISCONETSYNCMIB.EntityData.YFilter = cISCONETSYNCMIB.YFilter
    cISCONETSYNCMIB.EntityData.YangName = "CISCO-NETSYNC-MIB"
    cISCONETSYNCMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCONETSYNCMIB.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cISCONETSYNCMIB.EntityData.SegmentPath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB"
    cISCONETSYNCMIB.EntityData.AbsolutePath = cISCONETSYNCMIB.EntityData.SegmentPath
    cISCONETSYNCMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCONETSYNCMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCONETSYNCMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCONETSYNCMIB.EntityData.Children = types.NewOrderedMap()
    cISCONETSYNCMIB.EntityData.Children.Append("ciscoNetsyncMIBNotifControl", types.YChild{"CiscoNetsyncMIBNotifControl", &cISCONETSYNCMIB.CiscoNetsyncMIBNotifControl})
    cISCONETSYNCMIB.EntityData.Children.Append("cnsClkSelGlobalTable", types.YChild{"CnsClkSelGlobalTable", &cISCONETSYNCMIB.CnsClkSelGlobalTable})
    cISCONETSYNCMIB.EntityData.Children.Append("cnsSelectedInputSourceTable", types.YChild{"CnsSelectedInputSourceTable", &cISCONETSYNCMIB.CnsSelectedInputSourceTable})
    cISCONETSYNCMIB.EntityData.Children.Append("cnsInputSourceTable", types.YChild{"CnsInputSourceTable", &cISCONETSYNCMIB.CnsInputSourceTable})
    cISCONETSYNCMIB.EntityData.Children.Append("cnsExtOutputTable", types.YChild{"CnsExtOutputTable", &cISCONETSYNCMIB.CnsExtOutputTable})
    cISCONETSYNCMIB.EntityData.Children.Append("cnsT4ClockSourceTable", types.YChild{"CnsT4ClockSourceTable", &cISCONETSYNCMIB.CnsT4ClockSourceTable})
    cISCONETSYNCMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCONETSYNCMIB.EntityData.YListKeys = []string {}

    return &(cISCONETSYNCMIB.EntityData)
}

// CISCONETSYNCMIB_CiscoNetsyncMIBNotifControl
type CISCONETSYNCMIB_CiscoNetsyncMIBNotifControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A control object to enable/disable ciscoNetsyncSelectedT0Clock,
    // ciscoNetsyncSelectedT4Clock, ciscoNetsyncInputSignalFailureStatus,
    // ciscoNetsyncInputAlarmStatus notifications at the system level. This object
    // should hold any of the below values.     true - The notif is enabled
    // globally for the system     false- The notif is disabled globally for the
    // system. The type is bool.
    CnsMIBEnableStatusNotification interface{}
}

func (ciscoNetsyncMIBNotifControl *CISCONETSYNCMIB_CiscoNetsyncMIBNotifControl) GetEntityData() *types.CommonEntityData {
    ciscoNetsyncMIBNotifControl.EntityData.YFilter = ciscoNetsyncMIBNotifControl.YFilter
    ciscoNetsyncMIBNotifControl.EntityData.YangName = "ciscoNetsyncMIBNotifControl"
    ciscoNetsyncMIBNotifControl.EntityData.BundleName = "cisco_ios_xe"
    ciscoNetsyncMIBNotifControl.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    ciscoNetsyncMIBNotifControl.EntityData.SegmentPath = "ciscoNetsyncMIBNotifControl"
    ciscoNetsyncMIBNotifControl.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/" + ciscoNetsyncMIBNotifControl.EntityData.SegmentPath
    ciscoNetsyncMIBNotifControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoNetsyncMIBNotifControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoNetsyncMIBNotifControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoNetsyncMIBNotifControl.EntityData.Children = types.NewOrderedMap()
    ciscoNetsyncMIBNotifControl.EntityData.Leafs = types.NewOrderedMap()
    ciscoNetsyncMIBNotifControl.EntityData.Leafs.Append("cnsMIBEnableStatusNotification", types.YLeaf{"CnsMIBEnableStatusNotification", ciscoNetsyncMIBNotifControl.CnsMIBEnableStatusNotification})

    ciscoNetsyncMIBNotifControl.EntityData.YListKeys = []string {}

    return &(ciscoNetsyncMIBNotifControl.EntityData)
}

// CISCONETSYNCMIB_CnsClkSelGlobalTable
// G.781 clock selection process table.
// This table contains the global parameters for the G.781 clock
// selection process.
type CISCONETSYNCMIB_CnsClkSelGlobalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is added to cnsClkSelGlobalTable when G.781 clock selection is
    // enabled in the device configuration.  The entry is removed when G.781 clock
    // selection is removed from the configuration. The type is slice of
    // CISCONETSYNCMIB_CnsClkSelGlobalTable_CnsClkSelGlobalEntry.
    CnsClkSelGlobalEntry []*CISCONETSYNCMIB_CnsClkSelGlobalTable_CnsClkSelGlobalEntry
}

func (cnsClkSelGlobalTable *CISCONETSYNCMIB_CnsClkSelGlobalTable) GetEntityData() *types.CommonEntityData {
    cnsClkSelGlobalTable.EntityData.YFilter = cnsClkSelGlobalTable.YFilter
    cnsClkSelGlobalTable.EntityData.YangName = "cnsClkSelGlobalTable"
    cnsClkSelGlobalTable.EntityData.BundleName = "cisco_ios_xe"
    cnsClkSelGlobalTable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsClkSelGlobalTable.EntityData.SegmentPath = "cnsClkSelGlobalTable"
    cnsClkSelGlobalTable.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/" + cnsClkSelGlobalTable.EntityData.SegmentPath
    cnsClkSelGlobalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsClkSelGlobalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsClkSelGlobalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsClkSelGlobalTable.EntityData.Children = types.NewOrderedMap()
    cnsClkSelGlobalTable.EntityData.Children.Append("cnsClkSelGlobalEntry", types.YChild{"CnsClkSelGlobalEntry", nil})
    for i := range cnsClkSelGlobalTable.CnsClkSelGlobalEntry {
        cnsClkSelGlobalTable.EntityData.Children.Append(types.GetSegmentPath(cnsClkSelGlobalTable.CnsClkSelGlobalEntry[i]), types.YChild{"CnsClkSelGlobalEntry", cnsClkSelGlobalTable.CnsClkSelGlobalEntry[i]})
    }
    cnsClkSelGlobalTable.EntityData.Leafs = types.NewOrderedMap()

    cnsClkSelGlobalTable.EntityData.YListKeys = []string {}

    return &(cnsClkSelGlobalTable.EntityData)
}

// CISCONETSYNCMIB_CnsClkSelGlobalTable_CnsClkSelGlobalEntry
// An entry is added to cnsClkSelGlobalTable when G.781 clock
// selection is enabled in the device configuration.  The entry
// is removed when G.781 clock selection is removed from the
// configuration.
type CISCONETSYNCMIB_CnsClkSelGlobalTable_CnsClkSelGlobalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely represents a clock
    // selection process.  This index is assigned arbitrarily by the system and
    // may not be persistent across reboots. The type is interface{} with range:
    // 0..4294967295.
    CnsClkSelGloProcIndex interface{}

    // This object indicates the QL mode of the network synchronization clock
    // selection process as described in ITU-T standard G.781 section 5.12. The
    // type is CiscoNetsyncQLMode.
    CnsClkSelGlobProcessMode interface{}

    // This object indicates the operating mode of the system clock. The type is
    // CiscoNetsyncClockMode.
    CnsClkSelGlobClockMode interface{}

    // This object indicates whether the G.781 clock selection is enabled or not. 
    // 'true'  - G.781 clock selection is enabled 'false' - G.781 clock selection
    // is disabled. The type is bool.
    CnsClkSelGlobNetsyncEnable interface{}

    // This object indicates the revertive mode setting in the G.781 clock
    // selection process.  The switching of clock sources can be made revertive or
    // non-revertive. In non-revertive mode, an alternate clock source is
    // maintained even after the original clock source has recovered from the
    // failure that caused the switch. In revertive mode, the clock selection
    // process switches back to the original clock source after recovering from
    // the failure. The type is bool.
    CnsClkSelGlobRevertiveMode interface{}

    // This object indicates if global ESMC is enabled. With ESMC enabled
    // globally, the system is capable of handling ESMC messages. The type is
    // bool.
    CnsClkSelGlobESMCMode interface{}

    // This object indicates the network synchronization EEC (Ethernet Equipment
    // Clock) option. The type is CiscoNetsyncEECOption.
    CnsClkSelGlobEECOption interface{}

    // This object indicates the synchronization network option. The type is
    // CiscoNetsyncNetworkOption.
    CnsClkSelGlobNetworkOption interface{}

    // This object indicates the global holdoff time in the G.781 clock selection
    // process. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CnsClkSelGlobHoldoffTime interface{}

    // This object indicates the global wait-to-restore time in the G.781 clock
    // selection process. The type is interface{} with range: 0..4294967295. Units
    // are seconds.
    CnsClkSelGlobWtrTime interface{}

    // This object indicates the number of synchronization sources currently
    // configured for the G.781 clock selection process. The type is interface{}
    // with range: 0..255. Units are clock sources.
    CnsClkSelGlobNofSources interface{}

    // This object indicates the duration of the last holdover period in seconds.
    // If the holdover duration is less than a second, the object will carry the
    // value zero. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CnsClkSelGlobLastHoldoverSeconds interface{}

    // This object indicates the duration of the current holdover period. If a
    // system clock is in holdover mode, the object carries the current holdover
    // duration in seconds. If the system clock is not in holdover, the object
    // carries the value 0. The type is interface{} with range: 0..4294967295.
    // Units are seconds.
    CnsClkSelGlobCurrHoldoverSeconds interface{}
}

func (cnsClkSelGlobalEntry *CISCONETSYNCMIB_CnsClkSelGlobalTable_CnsClkSelGlobalEntry) GetEntityData() *types.CommonEntityData {
    cnsClkSelGlobalEntry.EntityData.YFilter = cnsClkSelGlobalEntry.YFilter
    cnsClkSelGlobalEntry.EntityData.YangName = "cnsClkSelGlobalEntry"
    cnsClkSelGlobalEntry.EntityData.BundleName = "cisco_ios_xe"
    cnsClkSelGlobalEntry.EntityData.ParentYangName = "cnsClkSelGlobalTable"
    cnsClkSelGlobalEntry.EntityData.SegmentPath = "cnsClkSelGlobalEntry" + types.AddKeyToken(cnsClkSelGlobalEntry.CnsClkSelGloProcIndex, "cnsClkSelGloProcIndex")
    cnsClkSelGlobalEntry.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/cnsClkSelGlobalTable/" + cnsClkSelGlobalEntry.EntityData.SegmentPath
    cnsClkSelGlobalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsClkSelGlobalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsClkSelGlobalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsClkSelGlobalEntry.EntityData.Children = types.NewOrderedMap()
    cnsClkSelGlobalEntry.EntityData.Leafs = types.NewOrderedMap()
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGloProcIndex", types.YLeaf{"CnsClkSelGloProcIndex", cnsClkSelGlobalEntry.CnsClkSelGloProcIndex})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobProcessMode", types.YLeaf{"CnsClkSelGlobProcessMode", cnsClkSelGlobalEntry.CnsClkSelGlobProcessMode})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobClockMode", types.YLeaf{"CnsClkSelGlobClockMode", cnsClkSelGlobalEntry.CnsClkSelGlobClockMode})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobNetsyncEnable", types.YLeaf{"CnsClkSelGlobNetsyncEnable", cnsClkSelGlobalEntry.CnsClkSelGlobNetsyncEnable})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobRevertiveMode", types.YLeaf{"CnsClkSelGlobRevertiveMode", cnsClkSelGlobalEntry.CnsClkSelGlobRevertiveMode})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobESMCMode", types.YLeaf{"CnsClkSelGlobESMCMode", cnsClkSelGlobalEntry.CnsClkSelGlobESMCMode})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobEECOption", types.YLeaf{"CnsClkSelGlobEECOption", cnsClkSelGlobalEntry.CnsClkSelGlobEECOption})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobNetworkOption", types.YLeaf{"CnsClkSelGlobNetworkOption", cnsClkSelGlobalEntry.CnsClkSelGlobNetworkOption})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobHoldoffTime", types.YLeaf{"CnsClkSelGlobHoldoffTime", cnsClkSelGlobalEntry.CnsClkSelGlobHoldoffTime})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobWtrTime", types.YLeaf{"CnsClkSelGlobWtrTime", cnsClkSelGlobalEntry.CnsClkSelGlobWtrTime})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobNofSources", types.YLeaf{"CnsClkSelGlobNofSources", cnsClkSelGlobalEntry.CnsClkSelGlobNofSources})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobLastHoldoverSeconds", types.YLeaf{"CnsClkSelGlobLastHoldoverSeconds", cnsClkSelGlobalEntry.CnsClkSelGlobLastHoldoverSeconds})
    cnsClkSelGlobalEntry.EntityData.Leafs.Append("cnsClkSelGlobCurrHoldoverSeconds", types.YLeaf{"CnsClkSelGlobCurrHoldoverSeconds", cnsClkSelGlobalEntry.CnsClkSelGlobCurrHoldoverSeconds})

    cnsClkSelGlobalEntry.EntityData.YListKeys = []string {"CnsClkSelGloProcIndex"}

    return &(cnsClkSelGlobalEntry.EntityData)
}

// CISCONETSYNCMIB_CnsSelectedInputSourceTable
// T0 selected clock source table.
// This table contains the selected clock source for the input T0
// clock.
type CISCONETSYNCMIB_CnsSelectedInputSourceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when the G.781 clock selection process has
    // successfully selected a T0 clock source.  The entry shall remain during the
    // time the G.781 clock selection process remains enabled. The type is slice
    // of CISCONETSYNCMIB_CnsSelectedInputSourceTable_CnsSelectedInputSourceEntry.
    CnsSelectedInputSourceEntry []*CISCONETSYNCMIB_CnsSelectedInputSourceTable_CnsSelectedInputSourceEntry
}

func (cnsSelectedInputSourceTable *CISCONETSYNCMIB_CnsSelectedInputSourceTable) GetEntityData() *types.CommonEntityData {
    cnsSelectedInputSourceTable.EntityData.YFilter = cnsSelectedInputSourceTable.YFilter
    cnsSelectedInputSourceTable.EntityData.YangName = "cnsSelectedInputSourceTable"
    cnsSelectedInputSourceTable.EntityData.BundleName = "cisco_ios_xe"
    cnsSelectedInputSourceTable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsSelectedInputSourceTable.EntityData.SegmentPath = "cnsSelectedInputSourceTable"
    cnsSelectedInputSourceTable.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/" + cnsSelectedInputSourceTable.EntityData.SegmentPath
    cnsSelectedInputSourceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsSelectedInputSourceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsSelectedInputSourceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsSelectedInputSourceTable.EntityData.Children = types.NewOrderedMap()
    cnsSelectedInputSourceTable.EntityData.Children.Append("cnsSelectedInputSourceEntry", types.YChild{"CnsSelectedInputSourceEntry", nil})
    for i := range cnsSelectedInputSourceTable.CnsSelectedInputSourceEntry {
        cnsSelectedInputSourceTable.EntityData.Children.Append(types.GetSegmentPath(cnsSelectedInputSourceTable.CnsSelectedInputSourceEntry[i]), types.YChild{"CnsSelectedInputSourceEntry", cnsSelectedInputSourceTable.CnsSelectedInputSourceEntry[i]})
    }
    cnsSelectedInputSourceTable.EntityData.Leafs = types.NewOrderedMap()

    cnsSelectedInputSourceTable.EntityData.YListKeys = []string {}

    return &(cnsSelectedInputSourceTable.EntityData)
}

// CISCONETSYNCMIB_CnsSelectedInputSourceTable_CnsSelectedInputSourceEntry
// An entry is created in the table when the G.781 clock
// selection process has successfully selected a T0 clock
// source.  The entry shall remain during the time
// the G.781 clock selection process remains enabled.
type CISCONETSYNCMIB_CnsSelectedInputSourceTable_CnsSelectedInputSourceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely represents an entry in this
    // table. This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    CnsSelInpSrcNetsyncIndex interface{}

    // This object indicates the name of the selected T0 clock. The type is string
    // with length: 1..255.
    CnsSelInpSrcName interface{}

    // This object indicates the type of the selected T0 clock. The type is
    // CiscoNetsyncIfType.
    CnsSelInpSrcIntfType interface{}

    // This object indicates the selected T0 clock source's effective quality
    // level, which is the derived clock quality based on the three factors:  (a)
    // Received quality level.  (b) Configured Rx quality level.  This factor
    // supersedes (a).  (c) System overridden quality level as a result of
    // exceptional events such as signal failure or ESMC failure.  This factor
    // supersedes (a) and (b). The type is CiscoNetsyncQualityLevel.
    CnsSelInpSrcQualityLevel interface{}

    // This object indicates the configured priority of the selected T0 clock. A
    // smaller value represents a higher priority. The type is interface{} with
    // range: 1..1024.
    CnsSelInpSrcPriority interface{}

    // This object indicates the timestamp of the T0 clock source being selected
    // by the G.781 clock selection process. The type is interface{} with range:
    // 0..4294967295.
    CnsSelInpSrcTimestamp interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source.  The 'true' value indicates the currently selected
    // clock source is a result of the forced switching. The 'false' value
    // indicates the currently selected clock source is not a result of forced
    // switching. The type is bool.
    CnsSelInpSrcFSW interface{}

    // This object indicates the manual switching flag. The 'true' value indicates
    // the currently selected clock source is a result of the manual switch
    // command. The command allows a user to select a synchronization source
    // assuming it is enabled, not locked out, not in signal fail condition, and
    // has a QL better than DNU in QL-enabled mode. Furthermore, in QL-enabled
    // mode, a manual switch can be performed only to a source which has the
    // highest available QL. The type is bool.
    CnsSelInpSrcMSW interface{}
}

func (cnsSelectedInputSourceEntry *CISCONETSYNCMIB_CnsSelectedInputSourceTable_CnsSelectedInputSourceEntry) GetEntityData() *types.CommonEntityData {
    cnsSelectedInputSourceEntry.EntityData.YFilter = cnsSelectedInputSourceEntry.YFilter
    cnsSelectedInputSourceEntry.EntityData.YangName = "cnsSelectedInputSourceEntry"
    cnsSelectedInputSourceEntry.EntityData.BundleName = "cisco_ios_xe"
    cnsSelectedInputSourceEntry.EntityData.ParentYangName = "cnsSelectedInputSourceTable"
    cnsSelectedInputSourceEntry.EntityData.SegmentPath = "cnsSelectedInputSourceEntry" + types.AddKeyToken(cnsSelectedInputSourceEntry.CnsSelInpSrcNetsyncIndex, "cnsSelInpSrcNetsyncIndex")
    cnsSelectedInputSourceEntry.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/cnsSelectedInputSourceTable/" + cnsSelectedInputSourceEntry.EntityData.SegmentPath
    cnsSelectedInputSourceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsSelectedInputSourceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsSelectedInputSourceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsSelectedInputSourceEntry.EntityData.Children = types.NewOrderedMap()
    cnsSelectedInputSourceEntry.EntityData.Leafs = types.NewOrderedMap()
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcNetsyncIndex", types.YLeaf{"CnsSelInpSrcNetsyncIndex", cnsSelectedInputSourceEntry.CnsSelInpSrcNetsyncIndex})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcName", types.YLeaf{"CnsSelInpSrcName", cnsSelectedInputSourceEntry.CnsSelInpSrcName})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcIntfType", types.YLeaf{"CnsSelInpSrcIntfType", cnsSelectedInputSourceEntry.CnsSelInpSrcIntfType})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcQualityLevel", types.YLeaf{"CnsSelInpSrcQualityLevel", cnsSelectedInputSourceEntry.CnsSelInpSrcQualityLevel})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcPriority", types.YLeaf{"CnsSelInpSrcPriority", cnsSelectedInputSourceEntry.CnsSelInpSrcPriority})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcTimestamp", types.YLeaf{"CnsSelInpSrcTimestamp", cnsSelectedInputSourceEntry.CnsSelInpSrcTimestamp})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcFSW", types.YLeaf{"CnsSelInpSrcFSW", cnsSelectedInputSourceEntry.CnsSelInpSrcFSW})
    cnsSelectedInputSourceEntry.EntityData.Leafs.Append("cnsSelInpSrcMSW", types.YLeaf{"CnsSelInpSrcMSW", cnsSelectedInputSourceEntry.CnsSelInpSrcMSW})

    cnsSelectedInputSourceEntry.EntityData.YListKeys = []string {"CnsSelInpSrcNetsyncIndex"}

    return &(cnsSelectedInputSourceEntry.EntityData)
}

// CISCONETSYNCMIB_CnsInputSourceTable
// T0 clock source table.
// This table contains a list of input sources for input T0 clock
// selection.
type CISCONETSYNCMIB_CnsInputSourceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a T0 clock source in the
    // configuration. An entry is removed  in the table when a user removes a T0
    // clock source from the configuration. The type is slice of
    // CISCONETSYNCMIB_CnsInputSourceTable_CnsInputSourceEntry.
    CnsInputSourceEntry []*CISCONETSYNCMIB_CnsInputSourceTable_CnsInputSourceEntry
}

func (cnsInputSourceTable *CISCONETSYNCMIB_CnsInputSourceTable) GetEntityData() *types.CommonEntityData {
    cnsInputSourceTable.EntityData.YFilter = cnsInputSourceTable.YFilter
    cnsInputSourceTable.EntityData.YangName = "cnsInputSourceTable"
    cnsInputSourceTable.EntityData.BundleName = "cisco_ios_xe"
    cnsInputSourceTable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsInputSourceTable.EntityData.SegmentPath = "cnsInputSourceTable"
    cnsInputSourceTable.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/" + cnsInputSourceTable.EntityData.SegmentPath
    cnsInputSourceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsInputSourceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsInputSourceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsInputSourceTable.EntityData.Children = types.NewOrderedMap()
    cnsInputSourceTable.EntityData.Children.Append("cnsInputSourceEntry", types.YChild{"CnsInputSourceEntry", nil})
    for i := range cnsInputSourceTable.CnsInputSourceEntry {
        cnsInputSourceTable.EntityData.Children.Append(types.GetSegmentPath(cnsInputSourceTable.CnsInputSourceEntry[i]), types.YChild{"CnsInputSourceEntry", cnsInputSourceTable.CnsInputSourceEntry[i]})
    }
    cnsInputSourceTable.EntityData.Leafs = types.NewOrderedMap()

    cnsInputSourceTable.EntityData.YListKeys = []string {}

    return &(cnsInputSourceTable.EntityData)
}

// CISCONETSYNCMIB_CnsInputSourceTable_CnsInputSourceEntry
// An entry is created in the table when a user adds a T0
// clock source in the configuration. An entry is removed 
// in the table when a user removes a T0 clock source from
// the configuration.
type CISCONETSYNCMIB_CnsInputSourceTable_CnsInputSourceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely represents an entry in this
    // table. This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    CnsInpSrcNetsyncIndex interface{}

    // This object indicates the name of an input clock source configured for the
    // T0 clock selection. The type is string with length: 1..255.
    CnsInpSrcName interface{}

    // This object indicates the type of an input clock source configured for the
    // T0 clock selection. The type is CiscoNetsyncIfType.
    CnsInpSrcIntfType interface{}

    // This object indicates the priority of an input clock source configured for
    // the T0 clock selection.  A smaller value represents a higher priority. The
    // type is interface{} with range: 1..1024.
    CnsInpSrcPriority interface{}

    // This object indicates the ESMC capability of an input clock source
    // configured for the T0 clock selection.  This is applicable only to
    // Synchronous Ethernet input clock source identified by cnsInpSrcIntfType
    // 'netsyncIfTypeEthernet'. The type is CiscoNetsyncESMCCap.
    CnsInpSrcESMCCap interface{}

    // This object indicates the SSM capability of an input clock source
    // configured for the T0 clock selection. This is applicable only to any
    // synchronous interface clock source except SyncE interface, which is
    // identified by cnsInpSrcIntfType 'netsyncIfTypeEthernet'. The type is
    // CiscoNetsyncSSMCap.
    CnsInpSrcSSMCap interface{}

    // This object indicates the configured transmit clock quality level of an
    // input clock source. The type is CiscoNetsyncQualityLevel.
    CnsInpSrcQualityLevelTxCfg interface{}

    // This object indicates the configured receive clock quality level of an
    // input clock source. The type is CiscoNetsyncQualityLevel.
    CnsInpSrcQualityLevelRxCfg interface{}

    // This object indicates the most recent clock quality level transmitted on
    // the input clock source. The type is CiscoNetsyncQualityLevel.
    CnsInpSrcQualityLevelTx interface{}

    // This object indicates the last clock quality level received on the input
    // clock source. The type is CiscoNetsyncQualityLevel.
    CnsInpSrcQualityLevelRx interface{}

    // This object indicates the current clock quality level of the input clock
    // source.  This is the effective quality which is derived from three values: 
    // 1) most recent clock quality level received, 2) forced clock quality level
    // (entered via configuration) 3) overridden clock quality level as a result
    // of line protocol down, signal failure, or alarms. The type is
    // CiscoNetsyncQualityLevel.
    CnsInpSrcQualityLevel interface{}

    // This object indicates the hold-off time value of an input clock source. 
    // The hold-off time prevents short activation of signal failure is passed to
    // the selection process.  When a signal failure event is reported on a clock
    // source, it waits the duration of the hold-off time before declaring signal
    // failure on the clock source. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CnsInpSrcHoldoffTime interface{}

    // This object indicates the wait-to-restore time value of an input clock
    // source.  The wait-to-restore time ensures that a previous failed
    // synchronization source is only again considered as available by the
    // selection process if it is fault-free for a certain time. When a signal
    // failure condition is cleared on a clock source, it waits the duration of
    // the wait-to-restore time before clearing the signal failure status on the
    // clock source. The type is interface{} with range: 0..4294967295. Units are
    // Seconds.
    CnsInpSrcWtrTime interface{}

    // This object indicates whether or not the lockout command has been applied
    // to a clock source.  The 'true' value means the clock source is not
    // considered by the selection process. The type is bool.
    CnsInpSrcLockout interface{}

    // This object indicates whether or not a signal failure event is currently
    // being reported on the input clock source. The type is bool.
    CnsInpSrcSignalFailure interface{}

    // This object indicates whether or not an alarm event is currently being
    // reported on the input clock source. The type is bool.
    CnsInpSrcAlarm interface{}

    // This object indicates the alarm reasons of an input clock source if an
    // alarm event is being reported on it. The type is map[string]bool.
    CnsInpSrcAlarmInfo interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source.  The 'true' value indicates the currently selected
    // clock source is a result of the forced switching. The 'false' value
    // indicates the currently selected clock source is not a result of forced
    // switching. The type is bool.
    CnsInpSrcFSW interface{}

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
    CnsInpSrcMSW interface{}
}

func (cnsInputSourceEntry *CISCONETSYNCMIB_CnsInputSourceTable_CnsInputSourceEntry) GetEntityData() *types.CommonEntityData {
    cnsInputSourceEntry.EntityData.YFilter = cnsInputSourceEntry.YFilter
    cnsInputSourceEntry.EntityData.YangName = "cnsInputSourceEntry"
    cnsInputSourceEntry.EntityData.BundleName = "cisco_ios_xe"
    cnsInputSourceEntry.EntityData.ParentYangName = "cnsInputSourceTable"
    cnsInputSourceEntry.EntityData.SegmentPath = "cnsInputSourceEntry" + types.AddKeyToken(cnsInputSourceEntry.CnsInpSrcNetsyncIndex, "cnsInpSrcNetsyncIndex")
    cnsInputSourceEntry.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/cnsInputSourceTable/" + cnsInputSourceEntry.EntityData.SegmentPath
    cnsInputSourceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsInputSourceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsInputSourceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsInputSourceEntry.EntityData.Children = types.NewOrderedMap()
    cnsInputSourceEntry.EntityData.Leafs = types.NewOrderedMap()
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcNetsyncIndex", types.YLeaf{"CnsInpSrcNetsyncIndex", cnsInputSourceEntry.CnsInpSrcNetsyncIndex})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcName", types.YLeaf{"CnsInpSrcName", cnsInputSourceEntry.CnsInpSrcName})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcIntfType", types.YLeaf{"CnsInpSrcIntfType", cnsInputSourceEntry.CnsInpSrcIntfType})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcPriority", types.YLeaf{"CnsInpSrcPriority", cnsInputSourceEntry.CnsInpSrcPriority})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcESMCCap", types.YLeaf{"CnsInpSrcESMCCap", cnsInputSourceEntry.CnsInpSrcESMCCap})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcSSMCap", types.YLeaf{"CnsInpSrcSSMCap", cnsInputSourceEntry.CnsInpSrcSSMCap})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcQualityLevelTxCfg", types.YLeaf{"CnsInpSrcQualityLevelTxCfg", cnsInputSourceEntry.CnsInpSrcQualityLevelTxCfg})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcQualityLevelRxCfg", types.YLeaf{"CnsInpSrcQualityLevelRxCfg", cnsInputSourceEntry.CnsInpSrcQualityLevelRxCfg})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcQualityLevelTx", types.YLeaf{"CnsInpSrcQualityLevelTx", cnsInputSourceEntry.CnsInpSrcQualityLevelTx})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcQualityLevelRx", types.YLeaf{"CnsInpSrcQualityLevelRx", cnsInputSourceEntry.CnsInpSrcQualityLevelRx})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcQualityLevel", types.YLeaf{"CnsInpSrcQualityLevel", cnsInputSourceEntry.CnsInpSrcQualityLevel})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcHoldoffTime", types.YLeaf{"CnsInpSrcHoldoffTime", cnsInputSourceEntry.CnsInpSrcHoldoffTime})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcWtrTime", types.YLeaf{"CnsInpSrcWtrTime", cnsInputSourceEntry.CnsInpSrcWtrTime})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcLockout", types.YLeaf{"CnsInpSrcLockout", cnsInputSourceEntry.CnsInpSrcLockout})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcSignalFailure", types.YLeaf{"CnsInpSrcSignalFailure", cnsInputSourceEntry.CnsInpSrcSignalFailure})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcAlarm", types.YLeaf{"CnsInpSrcAlarm", cnsInputSourceEntry.CnsInpSrcAlarm})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcAlarmInfo", types.YLeaf{"CnsInpSrcAlarmInfo", cnsInputSourceEntry.CnsInpSrcAlarmInfo})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcFSW", types.YLeaf{"CnsInpSrcFSW", cnsInputSourceEntry.CnsInpSrcFSW})
    cnsInputSourceEntry.EntityData.Leafs.Append("cnsInpSrcMSW", types.YLeaf{"CnsInpSrcMSW", cnsInputSourceEntry.CnsInpSrcMSW})

    cnsInputSourceEntry.EntityData.YListKeys = []string {"CnsInpSrcNetsyncIndex"}

    return &(cnsInputSourceEntry.EntityData)
}

// CISCONETSYNCMIB_CnsExtOutputTable
// T4 external output table.
// This table contains a list of T4 external outputs.
// 
// Each T4 external output is associated with clock
// source(s) to be found in cnsT4ClockSourceTable.
// The clock selection process considers all the
// available clock sources and select the T4 clock
// source based on the G.781 clock selection algorithm.
type CISCONETSYNCMIB_CnsExtOutputTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a T4 external output in
    // the configuration.  A T4 external output configured input clock sources are
    // defined in cnsT4ClockSourceTable.  An entry is removed from the table when
    // a user removes a T4 external output from the configuration. The type is
    // slice of CISCONETSYNCMIB_CnsExtOutputTable_CnsExtOutputEntry.
    CnsExtOutputEntry []*CISCONETSYNCMIB_CnsExtOutputTable_CnsExtOutputEntry
}

func (cnsExtOutputTable *CISCONETSYNCMIB_CnsExtOutputTable) GetEntityData() *types.CommonEntityData {
    cnsExtOutputTable.EntityData.YFilter = cnsExtOutputTable.YFilter
    cnsExtOutputTable.EntityData.YangName = "cnsExtOutputTable"
    cnsExtOutputTable.EntityData.BundleName = "cisco_ios_xe"
    cnsExtOutputTable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsExtOutputTable.EntityData.SegmentPath = "cnsExtOutputTable"
    cnsExtOutputTable.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/" + cnsExtOutputTable.EntityData.SegmentPath
    cnsExtOutputTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsExtOutputTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsExtOutputTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsExtOutputTable.EntityData.Children = types.NewOrderedMap()
    cnsExtOutputTable.EntityData.Children.Append("cnsExtOutputEntry", types.YChild{"CnsExtOutputEntry", nil})
    for i := range cnsExtOutputTable.CnsExtOutputEntry {
        cnsExtOutputTable.EntityData.Children.Append(types.GetSegmentPath(cnsExtOutputTable.CnsExtOutputEntry[i]), types.YChild{"CnsExtOutputEntry", cnsExtOutputTable.CnsExtOutputEntry[i]})
    }
    cnsExtOutputTable.EntityData.Leafs = types.NewOrderedMap()

    cnsExtOutputTable.EntityData.YListKeys = []string {}

    return &(cnsExtOutputTable.EntityData)
}

// CISCONETSYNCMIB_CnsExtOutputTable_CnsExtOutputEntry
// An entry is created in the table when a user adds
// a T4 external output in the configuration.  A T4 external
// output configured input clock sources are defined in
// cnsT4ClockSourceTable.
// 
// An entry is removed from the table when a user removes
// a T4 external output from the configuration.
type CISCONETSYNCMIB_CnsExtOutputTable_CnsExtOutputEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that uniquely represents an entry in this
    // table.  This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    CnsExtOutListIndex interface{}

    // An index that uniquely represents the selected input clock source whose
    // information is reported by a row in cnsT4ClockSourceTable. The index lists
    // the value of cnsT4ClkSrcNetsyncIndex, which is the input clock source of
    // the T4 external output selected by the G.781 clock selection process. The
    // type is interface{} with range: 1..4294967295.
    CnsExtOutSelNetsyncIndex interface{}

    // This object indicates the name of a T4 external output. The type is string
    // with length: 1..255.
    CnsExtOutName interface{}

    // This object indicates the interface type of the T4 external output. The
    // type is CiscoNetsyncIfType.
    CnsExtOutIntfType interface{}

    // This object indicates the clock quality of the T4 external output. The type
    // is CiscoNetsyncQualityLevel.
    CnsExtOutQualityLevel interface{}

    // This object indicates the priority of the selected clock source for a T4
    // external output.  A smaller value represents a higher priority. The type is
    // interface{} with range: 1..1024.
    CnsExtOutPriority interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source, The T4 selected synchronization source is
    // identified by cnsExtOutSelNetsyncIndex, which contains the index to the
    // clock source in cnsT4ClockSourceTable.  The 'true' value indicates the
    // currently selected T4 clock source is a result of the forced switching. The
    // 'false' value indicates the currently selected T4 clock source is not a
    // result of forced switching. The type is bool.
    CnsExtOutFSW interface{}

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
    CnsExtOutMSW interface{}

    // This object indicates whether or not a T4 external output is squelched. 
    // Squelching is a sychronization function defined to prevent transmission of
    // a timing signal with a quality that is lower than the quality of the clock
    // in the receiving networks element or SASE. It is also used for the
    // prevention of timing loops. The type is bool.
    CnsExtOutSquelch interface{}
}

func (cnsExtOutputEntry *CISCONETSYNCMIB_CnsExtOutputTable_CnsExtOutputEntry) GetEntityData() *types.CommonEntityData {
    cnsExtOutputEntry.EntityData.YFilter = cnsExtOutputEntry.YFilter
    cnsExtOutputEntry.EntityData.YangName = "cnsExtOutputEntry"
    cnsExtOutputEntry.EntityData.BundleName = "cisco_ios_xe"
    cnsExtOutputEntry.EntityData.ParentYangName = "cnsExtOutputTable"
    cnsExtOutputEntry.EntityData.SegmentPath = "cnsExtOutputEntry" + types.AddKeyToken(cnsExtOutputEntry.CnsExtOutListIndex, "cnsExtOutListIndex")
    cnsExtOutputEntry.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/cnsExtOutputTable/" + cnsExtOutputEntry.EntityData.SegmentPath
    cnsExtOutputEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsExtOutputEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsExtOutputEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsExtOutputEntry.EntityData.Children = types.NewOrderedMap()
    cnsExtOutputEntry.EntityData.Leafs = types.NewOrderedMap()
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutListIndex", types.YLeaf{"CnsExtOutListIndex", cnsExtOutputEntry.CnsExtOutListIndex})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutSelNetsyncIndex", types.YLeaf{"CnsExtOutSelNetsyncIndex", cnsExtOutputEntry.CnsExtOutSelNetsyncIndex})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutName", types.YLeaf{"CnsExtOutName", cnsExtOutputEntry.CnsExtOutName})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutIntfType", types.YLeaf{"CnsExtOutIntfType", cnsExtOutputEntry.CnsExtOutIntfType})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutQualityLevel", types.YLeaf{"CnsExtOutQualityLevel", cnsExtOutputEntry.CnsExtOutQualityLevel})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutPriority", types.YLeaf{"CnsExtOutPriority", cnsExtOutputEntry.CnsExtOutPriority})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutFSW", types.YLeaf{"CnsExtOutFSW", cnsExtOutputEntry.CnsExtOutFSW})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutMSW", types.YLeaf{"CnsExtOutMSW", cnsExtOutputEntry.CnsExtOutMSW})
    cnsExtOutputEntry.EntityData.Leafs.Append("cnsExtOutSquelch", types.YLeaf{"CnsExtOutSquelch", cnsExtOutputEntry.CnsExtOutSquelch})

    cnsExtOutputEntry.EntityData.YListKeys = []string {"CnsExtOutListIndex"}

    return &(cnsExtOutputEntry.EntityData)
}

// CISCONETSYNCMIB_CnsT4ClockSourceTable
// T4 clock source table.
// This table contains a list of input sources for a specific
// T4 external output. An entry shall be added to
// cnsExtOutputTable first. Then clock sources shall be
// added in this table for the selection process to select
// the appropriate T4 clock source.
type CISCONETSYNCMIB_CnsT4ClockSourceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in the table when a user adds a clock source to a T4
    // external output in the configuration. The T4 external output is defined in
    // the T4 external output table. An entry is removed in the table when a user
    // removes a T4 clock source from the configuration. The type is slice of
    // CISCONETSYNCMIB_CnsT4ClockSourceTable_CnsT4ClockSourceEntry.
    CnsT4ClockSourceEntry []*CISCONETSYNCMIB_CnsT4ClockSourceTable_CnsT4ClockSourceEntry
}

func (cnsT4ClockSourceTable *CISCONETSYNCMIB_CnsT4ClockSourceTable) GetEntityData() *types.CommonEntityData {
    cnsT4ClockSourceTable.EntityData.YFilter = cnsT4ClockSourceTable.YFilter
    cnsT4ClockSourceTable.EntityData.YangName = "cnsT4ClockSourceTable"
    cnsT4ClockSourceTable.EntityData.BundleName = "cisco_ios_xe"
    cnsT4ClockSourceTable.EntityData.ParentYangName = "CISCO-NETSYNC-MIB"
    cnsT4ClockSourceTable.EntityData.SegmentPath = "cnsT4ClockSourceTable"
    cnsT4ClockSourceTable.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/" + cnsT4ClockSourceTable.EntityData.SegmentPath
    cnsT4ClockSourceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsT4ClockSourceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsT4ClockSourceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsT4ClockSourceTable.EntityData.Children = types.NewOrderedMap()
    cnsT4ClockSourceTable.EntityData.Children.Append("cnsT4ClockSourceEntry", types.YChild{"CnsT4ClockSourceEntry", nil})
    for i := range cnsT4ClockSourceTable.CnsT4ClockSourceEntry {
        cnsT4ClockSourceTable.EntityData.Children.Append(types.GetSegmentPath(cnsT4ClockSourceTable.CnsT4ClockSourceEntry[i]), types.YChild{"CnsT4ClockSourceEntry", cnsT4ClockSourceTable.CnsT4ClockSourceEntry[i]})
    }
    cnsT4ClockSourceTable.EntityData.Leafs = types.NewOrderedMap()

    cnsT4ClockSourceTable.EntityData.YListKeys = []string {}

    return &(cnsT4ClockSourceTable.EntityData)
}

// CISCONETSYNCMIB_CnsT4ClockSourceTable_CnsT4ClockSourceEntry
// An entry is created in the table when a user adds a
// clock source to a T4 external output in the configuration.
// The T4 external output is defined in the T4 external
// output table. An entry is removed in the table when a user
// removes a T4 clock source from the configuration.
type CISCONETSYNCMIB_CnsT4ClockSourceTable_CnsT4ClockSourceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_netsync_mib.CISCONETSYNCMIB_CnsExtOutputTable_CnsExtOutputEntry_CnsExtOutListIndex
    CnsExtOutListIndex interface{}

    // This attribute is a key. An index that uniquely represents an entry in this
    // table.  This index is assigned arbitrarily by the clock selection process
    // and may not be persistent across reboots. The type is interface{} with
    // range: 1..4294967295.
    CnsT4ClkSrcNetsyncIndex interface{}

    // This object indicates the name of a input clock source configured for the
    // T4 clock selection. The type is string with length: 1..255.
    CnsT4ClkSrcName interface{}

    // This object indicates the type of an input clock source configured for the
    // T4 clock selection. The type is CiscoNetsyncIfType.
    CnsT4ClkSrcIntfType interface{}

    // This object indicates the priority of an input clock source configured for
    // the T4 clock selection.  A smaller value represents a higher priority. The
    // type is interface{} with range: 1..1024.
    CnsT4ClkSrcPriority interface{}

    // This object indicates the ESMC capability of an input clock source
    // configured for the T4 clock selection.  This is applicable only to
    // Synchronous Ethernet input clock source identified by cnsT4ClkSrcIntfType
    // 'netsyncIfTypeEthernet'. The type is CiscoNetsyncESMCCap.
    CnsT4ClkSrcESMCCap interface{}

    // This object indicates the SSM capability of an input clock source
    // configured for the T4 clock selection. This is applicable only to any
    // synchronous interface clock source except SyncE interface, which is
    // identified by cnsT4ClkSrcIntfType 'netsyncIfTypeEthernet'. The type is
    // CiscoNetsyncSSMCap.
    CnsT4ClkSrcSSMCap interface{}

    // This object indicates the configured transmit clock quality level of a T4
    // input clock source. The type is CiscoNetsyncQualityLevel.
    CnsT4ClkSrcQualityLevelTxCfg interface{}

    // This object indicates the configured receive clock quality level of a T4
    // input clock source. The type is CiscoNetsyncQualityLevel.
    CnsT4ClkSrcQualityLevelRxCfg interface{}

    // This object indicates the most recent clock quality level transmitted on
    // the T4 input clock source. The type is CiscoNetsyncQualityLevel.
    CnsT4ClkSrcQualityLevelTx interface{}

    // This object indicates the last clock quality level received on the T4 input
    // clock source. The type is CiscoNetsyncQualityLevel.
    CnsT4ClkSrcQualityLevelRx interface{}

    // This object indicates the current clock quality level of the T4 input clock
    // source.  This is the effective quality which is derived from three values: 
    // 1) most recent clock quality level received, 2) forced clock quality level
    // (entered via configuration) 3) overridden clock quality level as a result
    // of line protocol down, signal failure, or alarms. The type is
    // CiscoNetsyncQualityLevel.
    CnsT4ClkSrcQualityLevel interface{}

    // This object indicates the hold-off time value of a T4 input clock source. 
    // The hold-off time prevents short activation of signal failure is passed to
    // the selection process.  When a signal failure event is reported on a clock
    // source, it waits the duration of the hold-off time before declaring signal
    // failure on the clock source. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CnsT4ClkSrcHoldoffTime interface{}

    // This object indicates the wait-to-restore time value of a T4 input clock
    // source.  The wait-to-restore time ensures that a previous failed
    // synchronization source is only again considered as available by the
    // selection process if it is fault-free for a certain time. When a signal
    // failure condition is cleared on a clock source, it waits the duration of
    // the wait-to-restore time before clearing the signal failure status on the
    // clock source. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CnsT4ClkSrcWtrTime interface{}

    // This object indicates whether or not the lockout command has been applied
    // on a T4 clock source.  The 'true' value means the clock source is not
    // considered by the selection process. The type is bool.
    CnsT4ClkSrcLockout interface{}

    // This object indicates whether or not a signal failure event is currently
    // being reported on the T4 input clock source. The type is bool.
    CnsT4ClkSrcSignalFailure interface{}

    // This object indicates whether or not an alarm event is currently being
    // reported on the T4 input clock source. The type is bool.
    CnsT4ClkSrcAlarm interface{}

    // This object indicates the alarm reasons of a T4 input clock source if an
    // alarm event is being reported on the clock source. The type is
    // map[string]bool.
    CnsT4ClkSrcAlarmInfo interface{}

    // This object indicates the forced switching flag. Forced switching, as
    // described in G.781, is used to override the currently selected
    // synchronization source.  The 'true' value indicates the currently selected
    // T4 clock source is a result of the forced switching. The 'false' value
    // indicates the currently selected T4 clock source is not a result of forced
    // switching. The type is bool.
    CnsT4ClkSrcFSW interface{}

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
    CnsT4ClkSrcMSW interface{}
}

func (cnsT4ClockSourceEntry *CISCONETSYNCMIB_CnsT4ClockSourceTable_CnsT4ClockSourceEntry) GetEntityData() *types.CommonEntityData {
    cnsT4ClockSourceEntry.EntityData.YFilter = cnsT4ClockSourceEntry.YFilter
    cnsT4ClockSourceEntry.EntityData.YangName = "cnsT4ClockSourceEntry"
    cnsT4ClockSourceEntry.EntityData.BundleName = "cisco_ios_xe"
    cnsT4ClockSourceEntry.EntityData.ParentYangName = "cnsT4ClockSourceTable"
    cnsT4ClockSourceEntry.EntityData.SegmentPath = "cnsT4ClockSourceEntry" + types.AddKeyToken(cnsT4ClockSourceEntry.CnsExtOutListIndex, "cnsExtOutListIndex") + types.AddKeyToken(cnsT4ClockSourceEntry.CnsT4ClkSrcNetsyncIndex, "cnsT4ClkSrcNetsyncIndex")
    cnsT4ClockSourceEntry.EntityData.AbsolutePath = "CISCO-NETSYNC-MIB:CISCO-NETSYNC-MIB/cnsT4ClockSourceTable/" + cnsT4ClockSourceEntry.EntityData.SegmentPath
    cnsT4ClockSourceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnsT4ClockSourceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnsT4ClockSourceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnsT4ClockSourceEntry.EntityData.Children = types.NewOrderedMap()
    cnsT4ClockSourceEntry.EntityData.Leafs = types.NewOrderedMap()
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsExtOutListIndex", types.YLeaf{"CnsExtOutListIndex", cnsT4ClockSourceEntry.CnsExtOutListIndex})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcNetsyncIndex", types.YLeaf{"CnsT4ClkSrcNetsyncIndex", cnsT4ClockSourceEntry.CnsT4ClkSrcNetsyncIndex})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcName", types.YLeaf{"CnsT4ClkSrcName", cnsT4ClockSourceEntry.CnsT4ClkSrcName})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcIntfType", types.YLeaf{"CnsT4ClkSrcIntfType", cnsT4ClockSourceEntry.CnsT4ClkSrcIntfType})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcPriority", types.YLeaf{"CnsT4ClkSrcPriority", cnsT4ClockSourceEntry.CnsT4ClkSrcPriority})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcESMCCap", types.YLeaf{"CnsT4ClkSrcESMCCap", cnsT4ClockSourceEntry.CnsT4ClkSrcESMCCap})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcSSMCap", types.YLeaf{"CnsT4ClkSrcSSMCap", cnsT4ClockSourceEntry.CnsT4ClkSrcSSMCap})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcQualityLevelTxCfg", types.YLeaf{"CnsT4ClkSrcQualityLevelTxCfg", cnsT4ClockSourceEntry.CnsT4ClkSrcQualityLevelTxCfg})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcQualityLevelRxCfg", types.YLeaf{"CnsT4ClkSrcQualityLevelRxCfg", cnsT4ClockSourceEntry.CnsT4ClkSrcQualityLevelRxCfg})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcQualityLevelTx", types.YLeaf{"CnsT4ClkSrcQualityLevelTx", cnsT4ClockSourceEntry.CnsT4ClkSrcQualityLevelTx})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcQualityLevelRx", types.YLeaf{"CnsT4ClkSrcQualityLevelRx", cnsT4ClockSourceEntry.CnsT4ClkSrcQualityLevelRx})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcQualityLevel", types.YLeaf{"CnsT4ClkSrcQualityLevel", cnsT4ClockSourceEntry.CnsT4ClkSrcQualityLevel})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcHoldoffTime", types.YLeaf{"CnsT4ClkSrcHoldoffTime", cnsT4ClockSourceEntry.CnsT4ClkSrcHoldoffTime})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcWtrTime", types.YLeaf{"CnsT4ClkSrcWtrTime", cnsT4ClockSourceEntry.CnsT4ClkSrcWtrTime})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcLockout", types.YLeaf{"CnsT4ClkSrcLockout", cnsT4ClockSourceEntry.CnsT4ClkSrcLockout})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcSignalFailure", types.YLeaf{"CnsT4ClkSrcSignalFailure", cnsT4ClockSourceEntry.CnsT4ClkSrcSignalFailure})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcAlarm", types.YLeaf{"CnsT4ClkSrcAlarm", cnsT4ClockSourceEntry.CnsT4ClkSrcAlarm})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcAlarmInfo", types.YLeaf{"CnsT4ClkSrcAlarmInfo", cnsT4ClockSourceEntry.CnsT4ClkSrcAlarmInfo})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcFSW", types.YLeaf{"CnsT4ClkSrcFSW", cnsT4ClockSourceEntry.CnsT4ClkSrcFSW})
    cnsT4ClockSourceEntry.EntityData.Leafs.Append("cnsT4ClkSrcMSW", types.YLeaf{"CnsT4ClkSrcMSW", cnsT4ClockSourceEntry.CnsT4ClkSrcMSW})

    cnsT4ClockSourceEntry.EntityData.YListKeys = []string {"CnsExtOutListIndex", "CnsT4ClkSrcNetsyncIndex"}

    return &(cnsT4ClockSourceEntry.EntityData)
}

