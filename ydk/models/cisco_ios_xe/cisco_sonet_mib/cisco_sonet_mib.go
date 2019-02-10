// The MIB module to describe SONET/SDH interfaces
// objects. This is an extension to the standard SONET 
// MIB(RFC 2558).
package cisco_sonet_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_sonet_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-SONET-MIB CISCO-SONET-MIB}", reflect.TypeOf(CISCOSONETMIB{}))
    ydk.RegisterEntity("CISCO-SONET-MIB:CISCO-SONET-MIB", reflect.TypeOf(CISCOSONETMIB{}))
}

// CsApsLineSwitchReason represents   csApsNoSwitch : This is a state when no switch happens.
type CsApsLineSwitchReason string

const (
    CsApsLineSwitchReason_csApsOther CsApsLineSwitchReason = "csApsOther"

    CsApsLineSwitchReason_csApsRevertive CsApsLineSwitchReason = "csApsRevertive"

    CsApsLineSwitchReason_csApsManual CsApsLineSwitchReason = "csApsManual"

    CsApsLineSwitchReason_csApsSignalDefectLow CsApsLineSwitchReason = "csApsSignalDefectLow"

    CsApsLineSwitchReason_csApsSignalDefectHigh CsApsLineSwitchReason = "csApsSignalDefectHigh"

    CsApsLineSwitchReason_csApsSignalFailureLow CsApsLineSwitchReason = "csApsSignalFailureLow"

    CsApsLineSwitchReason_csApsSignalFailureHigh CsApsLineSwitchReason = "csApsSignalFailureHigh"

    CsApsLineSwitchReason_csApsForceSwitch CsApsLineSwitchReason = "csApsForceSwitch"

    CsApsLineSwitchReason_csApsLockOut CsApsLineSwitchReason = "csApsLockOut"

    CsApsLineSwitchReason_csApsNoSwitch CsApsLineSwitchReason = "csApsNoSwitch"
)

// CsApsLineFailureCode represents     csApsModeMismatch:        APS architecture mode mismatch.
type CsApsLineFailureCode string

const (
    CsApsLineFailureCode_csApsChannelMismatch CsApsLineFailureCode = "csApsChannelMismatch"

    CsApsLineFailureCode_csApsProtectionByteFail CsApsLineFailureCode = "csApsProtectionByteFail"

    CsApsLineFailureCode_csApsFEProtectionFailure CsApsLineFailureCode = "csApsFEProtectionFailure"

    CsApsLineFailureCode_csApsModeMismatch CsApsLineFailureCode = "csApsModeMismatch"
)

// CISCOSONETMIB
type CISCOSONETMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CsApsConfig CISCOSONETMIB_CsApsConfig

    
    CsNotifications CISCOSONETMIB_CsNotifications

    // The SONET/SDH configuration table. This table has objects  for configuring
    // sonet lines.
    CsConfigTable CISCOSONETMIB_CsConfigTable

    // This table contains objects to configure APS  (Automatic Protection
    // Switching) feature in a SONET  Line. APS is the ability to configure a pair
    // of SONET  lines for redundancy so that the hardware will  automatically
    // switch the active line from working line to the protection line or vice
    // versa, within 60ms,  when the active line fails.
    CsApsConfigTable CISCOSONETMIB_CsApsConfigTable

    // The SONET/SDH Section Total table. It contains the  cumulative sum of the
    // various statistics for the 24 hour period preceding the current interval.
    // The object  'sonetMediumValidIntervals' from RFC2558 contains the number of
    // 15 minute intervals that have elapsed since the line is enabled. .
    CssTotalTable CISCOSONETMIB_CssTotalTable

    // The SONET/SDH Section Trace table. This table contains  objects for tracing
    // the sonet section.
    CssTraceTable CISCOSONETMIB_CssTraceTable

    // The SONET/SDH Line Total table. It contains the  cumulative sum of the
    // various statistics for the 24  hour period preceding the current interval.
    // The object  'sonetMediumValidIntervals' from RFC2558 contains the  number
    // of 15 minute intervals that have elapsed since the line is enabled.
    CslTotalTable CISCOSONETMIB_CslTotalTable

    // The SONET/SDH Far End Line Total table. It contains the  cumulative sum of
    // the various statistics for the 24 hour  period preceding the current
    // interval. The object  'sonetMediumValidIntervals' from RFC2558 contains the
    // number of 15 minute intervals that have elapsed since the  line is enabled.
    CslFarEndTotalTable CISCOSONETMIB_CslFarEndTotalTable

    // The SONET/SDH Path Total table. It contains the cumulative  sum of the
    // various statistics for the 24 hour period  preceding the current
    // interval.The object  'sonetMediumValidIntervals' from RFC2558 contains the
    // number of 15 minute intervals that have elapsed since the line is  enabled.
    CspTotalTable CISCOSONETMIB_CspTotalTable

    // The SONET/SDH Far End Path Total table. Far End is the  remote end of the
    // line. The table contains the cumulative sum of the various statistics for
    // the 24 hour period  preceding the current interval. The object 
    // 'sonetMediumValidIntervals' from RFC2558 contains the number of 15 minute
    // intervals that have elapsed since the line is enabled. .
    CspFarEndTotalTable CISCOSONETMIB_CspFarEndTotalTable

    // The SONET/SDH Path Trace table. This table contains objects  for tracing
    // the sonet path.
    CspTraceTable CISCOSONETMIB_CspTraceTable

    // The SONET/SDH Section statistics table. This table  maintains the number of
    // times the line encountered Loss of Signal(LOS), Loss of frame(LOF), Alarm
    // Indication  signals(AISs), Remote failure indications(RFIs).
    CsStatsTable CISCOSONETMIB_CsStatsTable

    // This table contains objects to configure the VC( Virtual Container) related
    // properties of a TUG-3 within a AU-4  paths.  This table allows creation of
    // following multiplexing structure: STM-1/AU-4/TUG-3/TU-3/DS3
    // STM-1/AU-4/TUG-3/TU-3/E3 STM-1/AU-4/TUG-3/TUG-2/TU-11/DS1
    // STM-1/AU-4/TUG-3/TUG-2/TU-12/E1  Three entries are created in this table
    // for a given AU-4  path when cspSonetPathPayload object is set to one of the
    // following:     vt15vc11(4),     vt2vc12(5),     ds3(3),     e3(8),    
    // vtStructured(9).
    CsAu4Tug3ConfigTable CISCOSONETMIB_CsAu4Tug3ConfigTable
}

func (cISCOSONETMIB *CISCOSONETMIB) GetEntityData() *types.CommonEntityData {
    cISCOSONETMIB.EntityData.YFilter = cISCOSONETMIB.YFilter
    cISCOSONETMIB.EntityData.YangName = "CISCO-SONET-MIB"
    cISCOSONETMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSONETMIB.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cISCOSONETMIB.EntityData.SegmentPath = "CISCO-SONET-MIB:CISCO-SONET-MIB"
    cISCOSONETMIB.EntityData.AbsolutePath = cISCOSONETMIB.EntityData.SegmentPath
    cISCOSONETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSONETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSONETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSONETMIB.EntityData.Children = types.NewOrderedMap()
    cISCOSONETMIB.EntityData.Children.Append("csApsConfig", types.YChild{"CsApsConfig", &cISCOSONETMIB.CsApsConfig})
    cISCOSONETMIB.EntityData.Children.Append("csNotifications", types.YChild{"CsNotifications", &cISCOSONETMIB.CsNotifications})
    cISCOSONETMIB.EntityData.Children.Append("csConfigTable", types.YChild{"CsConfigTable", &cISCOSONETMIB.CsConfigTable})
    cISCOSONETMIB.EntityData.Children.Append("csApsConfigTable", types.YChild{"CsApsConfigTable", &cISCOSONETMIB.CsApsConfigTable})
    cISCOSONETMIB.EntityData.Children.Append("cssTotalTable", types.YChild{"CssTotalTable", &cISCOSONETMIB.CssTotalTable})
    cISCOSONETMIB.EntityData.Children.Append("cssTraceTable", types.YChild{"CssTraceTable", &cISCOSONETMIB.CssTraceTable})
    cISCOSONETMIB.EntityData.Children.Append("cslTotalTable", types.YChild{"CslTotalTable", &cISCOSONETMIB.CslTotalTable})
    cISCOSONETMIB.EntityData.Children.Append("cslFarEndTotalTable", types.YChild{"CslFarEndTotalTable", &cISCOSONETMIB.CslFarEndTotalTable})
    cISCOSONETMIB.EntityData.Children.Append("cspTotalTable", types.YChild{"CspTotalTable", &cISCOSONETMIB.CspTotalTable})
    cISCOSONETMIB.EntityData.Children.Append("cspFarEndTotalTable", types.YChild{"CspFarEndTotalTable", &cISCOSONETMIB.CspFarEndTotalTable})
    cISCOSONETMIB.EntityData.Children.Append("cspTraceTable", types.YChild{"CspTraceTable", &cISCOSONETMIB.CspTraceTable})
    cISCOSONETMIB.EntityData.Children.Append("csStatsTable", types.YChild{"CsStatsTable", &cISCOSONETMIB.CsStatsTable})
    cISCOSONETMIB.EntityData.Children.Append("csAu4Tug3ConfigTable", types.YChild{"CsAu4Tug3ConfigTable", &cISCOSONETMIB.CsAu4Tug3ConfigTable})
    cISCOSONETMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOSONETMIB.EntityData.YListKeys = []string {}

    return &(cISCOSONETMIB.EntityData)
}

// CISCOSONETMIB_CsApsConfig
type CISCOSONETMIB_CsApsConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The object indicates the APS line failure code. This is the failure
    // encountered by the APS line. Refer to CsApsLineFailureCode TC for failure
    // code definitions. The object is used for notifications. The type is
    // CsApsLineFailureCode.
    CsApsLineFailureCode interface{}

    // This object indicates the APS line switch reason. When the working line on
    // one end fails, its other end is told to do an APS switch.  Refer to
    // CsApsLineSwitchReason TC for more information. The object is used for
    // notifications. The type is CsApsLineSwitchReason.
    CsApsLineSwitchReason interface{}
}

func (csApsConfig *CISCOSONETMIB_CsApsConfig) GetEntityData() *types.CommonEntityData {
    csApsConfig.EntityData.YFilter = csApsConfig.YFilter
    csApsConfig.EntityData.YangName = "csApsConfig"
    csApsConfig.EntityData.BundleName = "cisco_ios_xe"
    csApsConfig.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csApsConfig.EntityData.SegmentPath = "csApsConfig"
    csApsConfig.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + csApsConfig.EntityData.SegmentPath
    csApsConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csApsConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csApsConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csApsConfig.EntityData.Children = types.NewOrderedMap()
    csApsConfig.EntityData.Leafs = types.NewOrderedMap()
    csApsConfig.EntityData.Leafs.Append("csApsLineFailureCode", types.YLeaf{"CsApsLineFailureCode", csApsConfig.CsApsLineFailureCode})
    csApsConfig.EntityData.Leafs.Append("csApsLineSwitchReason", types.YLeaf{"CsApsLineSwitchReason", csApsConfig.CsApsLineSwitchReason})

    csApsConfig.EntityData.YListKeys = []string {}

    return &(csApsConfig.EntityData)
}

// CISCOSONETMIB_CsNotifications
type CISCOSONETMIB_CsNotifications struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object controls if the generation of  ciscoSonetSectionStatusChange,
    // ciscoSonetLineStatusChange,  ciscoSonetPathStatusChange and
    // ciscoSonetVTStatusChange notifications is enabled. If the value of this
    // object is 'true(1)', then all notifications in this MIB are enabled;
    // otherwise they are disabled. The type is bool.
    CsNotificationsEnabled interface{}
}

func (csNotifications *CISCOSONETMIB_CsNotifications) GetEntityData() *types.CommonEntityData {
    csNotifications.EntityData.YFilter = csNotifications.YFilter
    csNotifications.EntityData.YangName = "csNotifications"
    csNotifications.EntityData.BundleName = "cisco_ios_xe"
    csNotifications.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csNotifications.EntityData.SegmentPath = "csNotifications"
    csNotifications.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + csNotifications.EntityData.SegmentPath
    csNotifications.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csNotifications.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csNotifications.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csNotifications.EntityData.Children = types.NewOrderedMap()
    csNotifications.EntityData.Leafs = types.NewOrderedMap()
    csNotifications.EntityData.Leafs.Append("csNotificationsEnabled", types.YLeaf{"CsNotificationsEnabled", csNotifications.CsNotificationsEnabled})

    csNotifications.EntityData.YListKeys = []string {}

    return &(csNotifications.EntityData)
}

// CISCOSONETMIB_CsConfigTable
// The SONET/SDH configuration table. This table has objects 
// for configuring sonet lines.
type CISCOSONETMIB_CsConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table. There is an entry for each SONET line  in the table.
    // Entries are automatically created for an  ifType value of sonet(39).
    // 'ifAdminStatus' from the ifTable  must be used to enable or disable a line.
    // A line is in  disabled(down) state unless provisioned 'up' using 
    // 'ifAdminStatus'. The type is slice of
    // CISCOSONETMIB_CsConfigTable_CsConfigEntry.
    CsConfigEntry []*CISCOSONETMIB_CsConfigTable_CsConfigEntry
}

func (csConfigTable *CISCOSONETMIB_CsConfigTable) GetEntityData() *types.CommonEntityData {
    csConfigTable.EntityData.YFilter = csConfigTable.YFilter
    csConfigTable.EntityData.YangName = "csConfigTable"
    csConfigTable.EntityData.BundleName = "cisco_ios_xe"
    csConfigTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csConfigTable.EntityData.SegmentPath = "csConfigTable"
    csConfigTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + csConfigTable.EntityData.SegmentPath
    csConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csConfigTable.EntityData.Children = types.NewOrderedMap()
    csConfigTable.EntityData.Children.Append("csConfigEntry", types.YChild{"CsConfigEntry", nil})
    for i := range csConfigTable.CsConfigEntry {
        csConfigTable.EntityData.Children.Append(types.GetSegmentPath(csConfigTable.CsConfigEntry[i]), types.YChild{"CsConfigEntry", csConfigTable.CsConfigEntry[i]})
    }
    csConfigTable.EntityData.Leafs = types.NewOrderedMap()

    csConfigTable.EntityData.YListKeys = []string {}

    return &(csConfigTable.EntityData)
}

// CISCOSONETMIB_CsConfigTable_CsConfigEntry
// An entry in the table. There is an entry for each SONET line 
// in the table. Entries are automatically created for an 
// ifType value of sonet(39). 'ifAdminStatus' from the ifTable 
// must be used to enable or disable a line. A line is in 
// disabled(down) state unless provisioned 'up' using 
// 'ifAdminStatus'.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object specifies the desired loopback mode configuration of the SONET
    // line. The possible values of this objects are follows:  noLoopback :       
    // Not in the loopback state.   lineLocal :          The signal transmitted
    // from this interface         is connected to the associated incoming        
    // receiver. This ensures that the SONET frame         transmitted from the
    // interface is received back         at the interface. lineRemote :        
    // The signal received at the interface is looped         back out to the
    // associated transmitter.         This ensures that the remote equipment that
    // originated the signal receives it back. The          signal may undergo
    // degradation as a result of         the characteristics of the transmission 
    // medium. The type is CsConfigLoopbackType.
    CsConfigLoopbackType interface{}

    // Specifies the source of the transmit clock.  loopTiming: indicates that the
    // recovered receive              clock is used as the transmit clock. 
    // localTiming: indicates that a local clock source is              used or
    // that an external clock is               attached to the box containing the 
    // interface. . The type is CsConfigXmtClockSource.
    CsConfigXmtClockSource interface{}

    // This object is used to disable or enable the Scrambling option in SONET
    // line. The type is CsConfigFrameScramble.
    CsConfigFrameScramble interface{}

    // This object represents the configured line type. Sts is SONET format. Stm
    // is SDH format.      sonetSts3c   : OC3 concatenated     sonetStm1    :
    // European standard OC3     sonetSts12c  : OC12 concatenated     sonetStm4   
    // : European standard OC12     sonetSts48c  : OC48 concatenated    
    // sonetStm16   : European standard OC48      sonetSts192c : OC-192
    // concatenated     sonetStm64   : European standard OC-192     sonetSts3    :
    // OC3 (unconcatenated)     . The type is CsConfigType.
    CsConfigType interface{}

    // This object specifies the type of RDI-V (Remote Defect Indication - Virtual
    // Tributary/Container) sent by this  Network Element (NE) to the remote
    // Network Element.        onebit     : use 1 bit RDI-V       threebit   : use
    // 3 bit enhanced RDI-V.  Default is onebit. The type is CsConfigRDIVType.
    CsConfigRDIVType interface{}

    // This object represents the type of RDI-P (Remote Defect Indication - Path)
    // sent by this Network Element (NE) to remote Network Element.        onebit 
    // : use 1 bit RDI-P       threebit   : use 3 bit enhanced RDI-P.  Default is
    // onebit. The type is CsConfigRDIPType.
    CsConfigRDIPType interface{}

    // Type of the tributary carried within the SONET/SDH signal.  vt15vc11    :
    // carries T1 signal vt2vc12     : carries E1 signal. The type is
    // CsTributaryType.
    CsTributaryType interface{}

    // This object represents the VT/VC mapping type.  asynchronous:    In this
    // mode, the channel structure of                   DS1/E1 is neither visible
    // nor preserved.  byteSynchronous: In this mode, the DS0 signals inside the  
    // VT/VC can be found and extracted from the                  frame.  Default
    // is asynchronous(1). The type is CsTributaryMappingType.
    CsTributaryMappingType interface{}

    // This object represents the framing type to be assigned to the virtual
    // tributaries in byte sync mapping mode.        notApplicable  :  If VT
    // mapping is not byteSynchronous(2).       dsx1ESF        :  Extended
    // Superframe Format       dsx1D4         :  Superframe Format  Default is
    // dsx1ESF(3) if csTributaryMappingType is  byteSynchronous(2). For
    // asynchronous(1) mapping, the default is  notApplicable(1).  The value
    // notApplicable(1) can not be set. The type is CsTributaryFramingType.
    CsTributaryFramingType interface{}

    // This object represents the mode used to transport DS0  signalling
    // information for T1 byteSynchronous mapping (GR253). In
    // signallingTransferMode(2), the robbed-bit signalling is  transferred to the
    // VT header. In clearMode(3), only the  framing bit is transferred to the VT
    // header.       Default is signallingTransferMode(2) if
    // csTributaryMappingType  is byteSynchronous. For asynchronous mapping, it is
    // notApplicable(1).  The value notApplicable(1) can not be set. The type is
    // CsSignallingTransportMode.
    CsSignallingTransportMode interface{}

    // This object represents the method used to group VCs into an STM-1 signal.
    // Applicable only to SDH.    au3Grouping: STM1<-AU-3<-TUG-2<-TU-12<-VC12 or  
    // STM1<-AU-3<-TUG-2<-TU-11<-VC11.    au4Grouping:
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-12<-VC12 or               
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-11<-VC11.  Default is au3Grouping(2) for SDH
    // and notApplicable(1) for SONET. The type is CsTributaryGroupingType.
    CsTributaryGroupingType interface{}
}

func (csConfigEntry *CISCOSONETMIB_CsConfigTable_CsConfigEntry) GetEntityData() *types.CommonEntityData {
    csConfigEntry.EntityData.YFilter = csConfigEntry.YFilter
    csConfigEntry.EntityData.YangName = "csConfigEntry"
    csConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    csConfigEntry.EntityData.ParentYangName = "csConfigTable"
    csConfigEntry.EntityData.SegmentPath = "csConfigEntry" + types.AddKeyToken(csConfigEntry.IfIndex, "ifIndex")
    csConfigEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/csConfigTable/" + csConfigEntry.EntityData.SegmentPath
    csConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csConfigEntry.EntityData.Children = types.NewOrderedMap()
    csConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    csConfigEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", csConfigEntry.IfIndex})
    csConfigEntry.EntityData.Leafs.Append("csConfigLoopbackType", types.YLeaf{"CsConfigLoopbackType", csConfigEntry.CsConfigLoopbackType})
    csConfigEntry.EntityData.Leafs.Append("csConfigXmtClockSource", types.YLeaf{"CsConfigXmtClockSource", csConfigEntry.CsConfigXmtClockSource})
    csConfigEntry.EntityData.Leafs.Append("csConfigFrameScramble", types.YLeaf{"CsConfigFrameScramble", csConfigEntry.CsConfigFrameScramble})
    csConfigEntry.EntityData.Leafs.Append("csConfigType", types.YLeaf{"CsConfigType", csConfigEntry.CsConfigType})
    csConfigEntry.EntityData.Leafs.Append("csConfigRDIVType", types.YLeaf{"CsConfigRDIVType", csConfigEntry.CsConfigRDIVType})
    csConfigEntry.EntityData.Leafs.Append("csConfigRDIPType", types.YLeaf{"CsConfigRDIPType", csConfigEntry.CsConfigRDIPType})
    csConfigEntry.EntityData.Leafs.Append("csTributaryType", types.YLeaf{"CsTributaryType", csConfigEntry.CsTributaryType})
    csConfigEntry.EntityData.Leafs.Append("csTributaryMappingType", types.YLeaf{"CsTributaryMappingType", csConfigEntry.CsTributaryMappingType})
    csConfigEntry.EntityData.Leafs.Append("csTributaryFramingType", types.YLeaf{"CsTributaryFramingType", csConfigEntry.CsTributaryFramingType})
    csConfigEntry.EntityData.Leafs.Append("csSignallingTransportMode", types.YLeaf{"CsSignallingTransportMode", csConfigEntry.CsSignallingTransportMode})
    csConfigEntry.EntityData.Leafs.Append("csTributaryGroupingType", types.YLeaf{"CsTributaryGroupingType", csConfigEntry.CsTributaryGroupingType})

    csConfigEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(csConfigEntry.EntityData)
}

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigFrameScramble represents option in SONET line.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigFrameScramble string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigFrameScramble_disabled CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigFrameScramble = "disabled"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigFrameScramble_enabled CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigFrameScramble = "enabled"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType represents         medium.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType_noLoopback CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType = "noLoopback"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType_lineLocal CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType = "lineLocal"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType_lineRemote CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigLoopbackType = "lineRemote"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIPType represents Default is onebit.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIPType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIPType_onebit CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIPType = "onebit"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIPType_threebit CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIPType = "threebit"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIVType represents Default is onebit.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIVType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIVType_onebit CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIVType = "onebit"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIVType_threebit CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigRDIVType = "threebit"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType represents     
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetSts3c CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetSts3c"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetStm1 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetStm1"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetSts12c CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetSts12c"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetStm4 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetStm4"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetSts48c CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetSts48c"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetStm16 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetStm16"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetSts192c CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetSts192c"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetStm64 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetStm64"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType_sonetSts3 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigType = "sonetSts3"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigXmtClockSource represents              interface. 
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigXmtClockSource string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigXmtClockSource_loopTiming CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigXmtClockSource = "loopTiming"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigXmtClockSource_localTiming CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsConfigXmtClockSource = "localTiming"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode represents The value notApplicable(1) can not be set.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode_notApplicable CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode = "notApplicable"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode_signallingTransferMode CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode = "signallingTransferMode"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode_clearMode CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsSignallingTransportMode = "clearMode"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType represents The value notApplicable(1) can not be set.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType_notApplicable CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType = "notApplicable"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType_dsx1D4 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType = "dsx1D4"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType_dsx1ESF CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryFramingType = "dsx1ESF"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType represents Default is au3Grouping(2) for SDH and notApplicable(1) for SONET.
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType_notApplicable CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType = "notApplicable"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType_au3Grouping CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType = "au3Grouping"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType_au4Grouping CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryGroupingType = "au4Grouping"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryMappingType represents Default is asynchronous(1).
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryMappingType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryMappingType_asynchronous CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryMappingType = "asynchronous"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryMappingType_byteSynchronous CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryMappingType = "byteSynchronous"
)

// CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryType represents vt2vc12     : carries E1 signal
type CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryType string

const (
    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryType_vt15vc11 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryType = "vt15vc11"

    CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryType_vt2vc12 CISCOSONETMIB_CsConfigTable_CsConfigEntry_CsTributaryType = "vt2vc12"
)

// CISCOSONETMIB_CsApsConfigTable
// This table contains objects to configure APS 
// (Automatic Protection Switching) feature in a SONET 
// Line. APS is the ability to configure a pair of SONET 
// lines for redundancy so that the hardware will 
// automatically switch the active line from working line
// to the protection line or vice versa, within 60ms, 
// when the active line fails.
type CISCOSONETMIB_CsApsConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created when an APS pair is configured. To create an entry, the
    // following objects must be  specified: csApsWorkingIndex,
    // csApsProtectionIndex, csApsEnable,   csApsArchMode. The protection line
    // must not be active, i.e, ifAdminStatus must be 'down',  while configuring 
    // APS. An entry is created by setting the value of  'csApsEnable' to
    // csApsEnabled (2) and deleted by  setting it to csApsDisabled (1). Once a
    // line is  configured as working line or protection line, it  remains in that
    // role until APS is disabled on that  sonet line pair. It remains in the 
    // working/protection  role even after the card is reset. The type is slice of
    // CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry.
    CsApsConfigEntry []*CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry
}

func (csApsConfigTable *CISCOSONETMIB_CsApsConfigTable) GetEntityData() *types.CommonEntityData {
    csApsConfigTable.EntityData.YFilter = csApsConfigTable.YFilter
    csApsConfigTable.EntityData.YangName = "csApsConfigTable"
    csApsConfigTable.EntityData.BundleName = "cisco_ios_xe"
    csApsConfigTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csApsConfigTable.EntityData.SegmentPath = "csApsConfigTable"
    csApsConfigTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + csApsConfigTable.EntityData.SegmentPath
    csApsConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csApsConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csApsConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csApsConfigTable.EntityData.Children = types.NewOrderedMap()
    csApsConfigTable.EntityData.Children.Append("csApsConfigEntry", types.YChild{"CsApsConfigEntry", nil})
    for i := range csApsConfigTable.CsApsConfigEntry {
        csApsConfigTable.EntityData.Children.Append(types.GetSegmentPath(csApsConfigTable.CsApsConfigEntry[i]), types.YChild{"CsApsConfigEntry", csApsConfigTable.CsApsConfigEntry[i]})
    }
    csApsConfigTable.EntityData.Leafs = types.NewOrderedMap()

    csApsConfigTable.EntityData.YListKeys = []string {}

    return &(csApsConfigTable.EntityData)
}

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry
// An entry is created when an APS pair is configured.
// To create an entry, the following objects must be 
// specified:
// csApsWorkingIndex, csApsProtectionIndex, csApsEnable,  
// csApsArchMode. The protection line must not be active,
// i.e, ifAdminStatus must be 'down',  while configuring 
// APS. An entry is created by setting the value of 
// 'csApsEnable' to csApsEnabled (2) and deleted by 
// setting it to csApsDisabled (1). Once a line is 
// configured as working line or protection line, it 
// remains in that role until APS is disabled on that 
// sonet line pair. It remains in the  working/protection 
// role even after the card is reset.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. When a pair of APS lines is configured, one line
    // has to be the working line, which is the primary line, and the other has to
    // be the protection line, which is the backup line. This object refers to the
    // working line in the APS pair. For G.783 AnnexB, this index refers to
    // Working Section 1. The type is interface{} with range: 1..2147483647.
    CsApsWorkingIndex interface{}

    // The protection line indicates that it will become the active line when an
    // APS switch occurs (APS switch could occur because of a failure on the
    // working line). For G.783 AnnexB, This index refers to Working Section 2.
    // The type is interface{} with range: 1..2147483647.
    CsApsProtectionIndex interface{}

    // This object is used to enable or disable the APS feature on the
    // working/protection line pairs. When enabled, the hardware will
    // automatically switch the active line  from the working line to the
    // protection line within 60ms, or vice versa. The type is CsApsEnable.
    CsApsEnable interface{}

    // This object is used to configure APS architecture mode on the
    // working/protection line pairs.   All of the following are supported on
    // single slot.  oneToOne(2) is not  supported across 2 slots,i.e. the  
    // working and protection slot numbers must be the same in   oneToOne(2).  
    // onePlusOne : This can be supported on the same card  and across 2 cards. 
    // This mode means that the transmit and receive signals  go only over the
    // active line(which could be working or   protection line). (straight cable
    // implied)   oneToOne : This is supported only on the same card  This mode
    // means that the transmit and receive signals  go over the working and
    // protection lines.  (straight cable implied)   anexBOnePlusOne : This can be
    // supported on the same card  and across 2 cards.  This mode is like the
    // onePlusOne mode, except that the  'csApsDirection' can only be
    // bi-directional.  (straight cable implied)   ycableOnePlusOneNok1k2: With
    // Y-cable ignore K1K2 bytes.  This mode is the Y-cable redundancy mode.  
    // straightOnePlusOneNok1k2 : With straight cable, ignore                     
    // K1K2 bytes. This mode is like                             onePlusOne, but
    // with K1, K2                              bytes ignored. The type is
    // CsApsArchMode.
    CsApsArchMode interface{}

    // This object indicates which line is currently active.  It could be the
    // working line(Section 1 for Annex B), the protection line(Section 2 for
    // Annex B) or none if neither working nor protection line is active.  This
    // object reflects the status of receive direction. The type is
    // CsApsActiveLine.
    CsApsActiveLine interface{}

    // This object contains the Bit Error Rate threshold for Signal Fault
    // detection on the working line. Once this threshold is exceeded, an APS
    // switch will occur. This value is 10 to the -n, where n is between 3 and 5.
    // The type is interface{} with range: 3..5.
    CsApsSigFaultBER interface{}

    // This object contains the Bit Error Rate threshold for Signal Degrade
    // detection on the working line. Once this threshold is exceeded, an APS
    // switch will occur. This value is 10 to -n where n is between 5 and 9. The
    // type is interface{} with range: 5..9.
    CsApsSigDegradeBER interface{}

    // This object contains interval in minutes to wait  before attempting to
    // switch back to working line.  Not applicable if the line is configured in 
    // non-revertive mode, i.e. protection line will  continue to be active, even
    // if failures on the  working line are cleared. The framer clears the 
    // signal-fault and signal-degrade when APS switch  occurs. Please refer to
    // 'csApsRevertive' for  description of non-revertive. The type is interface{}
    // with range: 1..12. Units are minutes.
    CsApsWaitToRestore interface{}

    // This object is used to configure the switching  direction which this APS
    // line supports.   Unidirectional : APS switch only in one direction.
    // Bidirectional  : APS switch in both ends of the line. The type is
    // CsApsDirection.
    CsApsDirection interface{}

    // This object is used to configure the APS revertive or nonrevertive option. 
    // revertive :    Will switch the working line back to active state after  
    // the Wait-To-restore interval has expired and the    working line
    // Signal-Fault/Signal-Degrade has been    cleared. Please refer to
    // 'csApsWaitToRestore' for    description of Wait-To-Restore interval.
    // nonrevertive :    The  protection line continues to be the active line,  
    // The active line does not switch to the working line. The type is
    // CsApsRevertive.
    CsApsRevertive interface{}

    // This object shows the actual APS direction that is  implemented on the Near
    // End terminal. APS direction  configured through csApsDirection is
    // negotiated with the Far End and APS direction setting acceptable to  both
    // ends is operational at the Near End. The type is CsApsDirectionOperational.
    CsApsDirectionOperational interface{}

    // This object shows the actual APS architecture mode that is implemented on
    // the Near End terminal. APS architecture mode configured through
    // csApsArchMode object is  negotiated with the Far End through APS channel. 
    // Architecture mode acceptable to both the Near End and  the Far End
    // terminals is then operational at the Near  End. This value can be different
    // than the APS  Architecture mode configured. The type is
    // CsApsArchModeOperational.
    CsApsArchModeOperational interface{}

    // This object allows to configure APS channel protocol to  be implemented at
    // Near End terminal.  K1 and K2 overhead bytes in a SONET signal are used as
    // an APS channel. This channel is used to carry APS protocol.  Possible
    // values: bellcore(1) : Implements APS channel protocol as defined           
    // in bellcore document GR-253-CORE. itu(2) : Implements APS channel protocol
    // as defined in           ITU document G.783. The type is
    // CsApsChannelProtocol.
    CsApsChannelProtocol interface{}

    // This object indicates APS line failure status. The type is map[string]bool.
    CsApsFailureStatus interface{}

    // This object indicates APS line switch reason. The type is
    // CsApsLineSwitchReason.
    CsApsSwitchReason interface{}

    // This object indicates which working section is the APS primary section. In
    // G.783 AnnexB, the K1/K2 Bytes are received on the secondary Section. All
    // the Switch Requests are for a switch from the primary section to the
    // secondary section. The object csApsActiveline will indicate which section
    // is currently carrying the traffic.  Once the switch request clears
    // normally, traffic is maintained on the section to which it was switched by
    // making that section the primary section.   Possible values: 
    // workingSection1(1): Working Section 1 is Primary Section
    // workingSection2(2): Working Section 2 is Primary Section none(3)          
    // : none. The type is CsApsPrimarySection.
    CsApsPrimarySection interface{}
}

func (csApsConfigEntry *CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry) GetEntityData() *types.CommonEntityData {
    csApsConfigEntry.EntityData.YFilter = csApsConfigEntry.YFilter
    csApsConfigEntry.EntityData.YangName = "csApsConfigEntry"
    csApsConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    csApsConfigEntry.EntityData.ParentYangName = "csApsConfigTable"
    csApsConfigEntry.EntityData.SegmentPath = "csApsConfigEntry" + types.AddKeyToken(csApsConfigEntry.CsApsWorkingIndex, "csApsWorkingIndex")
    csApsConfigEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/csApsConfigTable/" + csApsConfigEntry.EntityData.SegmentPath
    csApsConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csApsConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csApsConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csApsConfigEntry.EntityData.Children = types.NewOrderedMap()
    csApsConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    csApsConfigEntry.EntityData.Leafs.Append("csApsWorkingIndex", types.YLeaf{"CsApsWorkingIndex", csApsConfigEntry.CsApsWorkingIndex})
    csApsConfigEntry.EntityData.Leafs.Append("csApsProtectionIndex", types.YLeaf{"CsApsProtectionIndex", csApsConfigEntry.CsApsProtectionIndex})
    csApsConfigEntry.EntityData.Leafs.Append("csApsEnable", types.YLeaf{"CsApsEnable", csApsConfigEntry.CsApsEnable})
    csApsConfigEntry.EntityData.Leafs.Append("csApsArchMode", types.YLeaf{"CsApsArchMode", csApsConfigEntry.CsApsArchMode})
    csApsConfigEntry.EntityData.Leafs.Append("csApsActiveLine", types.YLeaf{"CsApsActiveLine", csApsConfigEntry.CsApsActiveLine})
    csApsConfigEntry.EntityData.Leafs.Append("csApsSigFaultBER", types.YLeaf{"CsApsSigFaultBER", csApsConfigEntry.CsApsSigFaultBER})
    csApsConfigEntry.EntityData.Leafs.Append("csApsSigDegradeBER", types.YLeaf{"CsApsSigDegradeBER", csApsConfigEntry.CsApsSigDegradeBER})
    csApsConfigEntry.EntityData.Leafs.Append("csApsWaitToRestore", types.YLeaf{"CsApsWaitToRestore", csApsConfigEntry.CsApsWaitToRestore})
    csApsConfigEntry.EntityData.Leafs.Append("csApsDirection", types.YLeaf{"CsApsDirection", csApsConfigEntry.CsApsDirection})
    csApsConfigEntry.EntityData.Leafs.Append("csApsRevertive", types.YLeaf{"CsApsRevertive", csApsConfigEntry.CsApsRevertive})
    csApsConfigEntry.EntityData.Leafs.Append("csApsDirectionOperational", types.YLeaf{"CsApsDirectionOperational", csApsConfigEntry.CsApsDirectionOperational})
    csApsConfigEntry.EntityData.Leafs.Append("csApsArchModeOperational", types.YLeaf{"CsApsArchModeOperational", csApsConfigEntry.CsApsArchModeOperational})
    csApsConfigEntry.EntityData.Leafs.Append("csApsChannelProtocol", types.YLeaf{"CsApsChannelProtocol", csApsConfigEntry.CsApsChannelProtocol})
    csApsConfigEntry.EntityData.Leafs.Append("csApsFailureStatus", types.YLeaf{"CsApsFailureStatus", csApsConfigEntry.CsApsFailureStatus})
    csApsConfigEntry.EntityData.Leafs.Append("csApsSwitchReason", types.YLeaf{"CsApsSwitchReason", csApsConfigEntry.CsApsSwitchReason})
    csApsConfigEntry.EntityData.Leafs.Append("csApsPrimarySection", types.YLeaf{"CsApsPrimarySection", csApsConfigEntry.CsApsPrimarySection})

    csApsConfigEntry.EntityData.YListKeys = []string {"CsApsWorkingIndex"}

    return &(csApsConfigEntry.EntityData)
}

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine represents This object reflects the status of receive direction.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine_csApsWorkingLine CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine = "csApsWorkingLine"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine_csApsProtectionLine CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine = "csApsProtectionLine"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine_csApsNone CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsActiveLine = "csApsNone"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode represents                             bytes ignored.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode_onePlusOne CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode = "onePlusOne"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode_oneToOne CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode = "oneToOne"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode_anexBOnePlusOne CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode = "anexBOnePlusOne"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode_ycableOnePlusOneNok1k2 CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode = "ycableOnePlusOneNok1k2"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode_straightOnePlusOneNok1k2 CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchMode = "straightOnePlusOneNok1k2"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational represents Architecture mode configured.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational_onePlusOne CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational = "onePlusOne"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational_oneToOne CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational = "oneToOne"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational_anexBOnePlusOne CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational = "anexBOnePlusOne"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational_ycableOnePlusOneNok1k2 CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational = "ycableOnePlusOneNok1k2"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational_straightOnePlusOneNok1k2 CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsArchModeOperational = "straightOnePlusOneNok1k2"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsChannelProtocol represents          ITU document G.783.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsChannelProtocol string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsChannelProtocol_bellcore CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsChannelProtocol = "bellcore"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsChannelProtocol_itu CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsChannelProtocol = "itu"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirection represents Bidirectional  : APS switch in both ends of the line.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirection string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirection_uniDirectional CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirection = "uniDirectional"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirection_biDirectional CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirection = "biDirectional"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirectionOperational represents both ends is operational at the Near End.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirectionOperational string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirectionOperational_uniDirectional CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirectionOperational = "uniDirectional"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirectionOperational_biDirectional CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsDirectionOperational = "biDirectional"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsEnable represents or vice versa.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsEnable string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsEnable_csApsDisabled CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsEnable = "csApsDisabled"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsEnable_csApsEnabled CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsEnable = "csApsEnabled"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection represents none(3)           : none.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection_workingSection1 CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection = "workingSection1"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection_workingSection2 CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection = "workingSection2"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection_none CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsPrimarySection = "none"
)

// CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsRevertive represents   The active line does not switch to the working line.
type CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsRevertive string

const (
    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsRevertive_nonrevertive CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsRevertive = "nonrevertive"

    CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsRevertive_revertive CISCOSONETMIB_CsApsConfigTable_CsApsConfigEntry_CsApsRevertive = "revertive"
)

// CISCOSONETMIB_CssTotalTable
// The SONET/SDH Section Total table. It contains the 
// cumulative sum of the various statistics for the 24 hour
// period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the
// number of 15 minute intervals that have elapsed since
// the line is enabled. 
type CISCOSONETMIB_CssTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_CssTotalTable_CssTotalEntry.
    CssTotalEntry []*CISCOSONETMIB_CssTotalTable_CssTotalEntry
}

func (cssTotalTable *CISCOSONETMIB_CssTotalTable) GetEntityData() *types.CommonEntityData {
    cssTotalTable.EntityData.YFilter = cssTotalTable.YFilter
    cssTotalTable.EntityData.YangName = "cssTotalTable"
    cssTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cssTotalTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cssTotalTable.EntityData.SegmentPath = "cssTotalTable"
    cssTotalTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cssTotalTable.EntityData.SegmentPath
    cssTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cssTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cssTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cssTotalTable.EntityData.Children = types.NewOrderedMap()
    cssTotalTable.EntityData.Children.Append("cssTotalEntry", types.YChild{"CssTotalEntry", nil})
    for i := range cssTotalTable.CssTotalEntry {
        cssTotalTable.EntityData.Children.Append(types.GetSegmentPath(cssTotalTable.CssTotalEntry[i]), types.YChild{"CssTotalEntry", cssTotalTable.CssTotalEntry[i]})
    }
    cssTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cssTotalTable.EntityData.YListKeys = []string {}

    return &(cssTotalTable.EntityData)
}

// CISCOSONETMIB_CssTotalTable_CssTotalEntry
// An entry in the SONET/SDH Section Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_CssTotalTable_CssTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Section in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // errored seconds.
    CssTotalESs interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Section
    // in the last 24 hours. The type is interface{} with range: 0..4294967295.
    // Units are severely errored seconds.
    CssTotalSESs interface{}

    // The number of Severely Errored Framing Seconds  encountered by a SONET/SDH
    // Section in the last 24 hours. The type is interface{} with range:
    // 0..4294967295. Units are severely errored framing seconds.
    CssTotalSEFSs interface{}

    // The number of Coding Violations encountered by a  SONET/SDH Section in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // coding violations.
    CssTotalCVs interface{}
}

func (cssTotalEntry *CISCOSONETMIB_CssTotalTable_CssTotalEntry) GetEntityData() *types.CommonEntityData {
    cssTotalEntry.EntityData.YFilter = cssTotalEntry.YFilter
    cssTotalEntry.EntityData.YangName = "cssTotalEntry"
    cssTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cssTotalEntry.EntityData.ParentYangName = "cssTotalTable"
    cssTotalEntry.EntityData.SegmentPath = "cssTotalEntry" + types.AddKeyToken(cssTotalEntry.IfIndex, "ifIndex")
    cssTotalEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cssTotalTable/" + cssTotalEntry.EntityData.SegmentPath
    cssTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cssTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cssTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cssTotalEntry.EntityData.Children = types.NewOrderedMap()
    cssTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cssTotalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cssTotalEntry.IfIndex})
    cssTotalEntry.EntityData.Leafs.Append("cssTotalESs", types.YLeaf{"CssTotalESs", cssTotalEntry.CssTotalESs})
    cssTotalEntry.EntityData.Leafs.Append("cssTotalSESs", types.YLeaf{"CssTotalSESs", cssTotalEntry.CssTotalSESs})
    cssTotalEntry.EntityData.Leafs.Append("cssTotalSEFSs", types.YLeaf{"CssTotalSEFSs", cssTotalEntry.CssTotalSEFSs})
    cssTotalEntry.EntityData.Leafs.Append("cssTotalCVs", types.YLeaf{"CssTotalCVs", cssTotalEntry.CssTotalCVs})

    cssTotalEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cssTotalEntry.EntityData)
}

// CISCOSONETMIB_CssTraceTable
// The SONET/SDH Section Trace table. This table contains 
// objects for tracing the sonet section.
type CISCOSONETMIB_CssTraceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the trace table. Entries exist for active sonet lines. The
    // objects in this table are used to verify  continued connection between the
    // two ends of the line. The type is slice of
    // CISCOSONETMIB_CssTraceTable_CssTraceEntry.
    CssTraceEntry []*CISCOSONETMIB_CssTraceTable_CssTraceEntry
}

func (cssTraceTable *CISCOSONETMIB_CssTraceTable) GetEntityData() *types.CommonEntityData {
    cssTraceTable.EntityData.YFilter = cssTraceTable.YFilter
    cssTraceTable.EntityData.YangName = "cssTraceTable"
    cssTraceTable.EntityData.BundleName = "cisco_ios_xe"
    cssTraceTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cssTraceTable.EntityData.SegmentPath = "cssTraceTable"
    cssTraceTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cssTraceTable.EntityData.SegmentPath
    cssTraceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cssTraceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cssTraceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cssTraceTable.EntityData.Children = types.NewOrderedMap()
    cssTraceTable.EntityData.Children.Append("cssTraceEntry", types.YChild{"CssTraceEntry", nil})
    for i := range cssTraceTable.CssTraceEntry {
        cssTraceTable.EntityData.Children.Append(types.GetSegmentPath(cssTraceTable.CssTraceEntry[i]), types.YChild{"CssTraceEntry", cssTraceTable.CssTraceEntry[i]})
    }
    cssTraceTable.EntityData.Leafs = types.NewOrderedMap()

    cssTraceTable.EntityData.YListKeys = []string {}

    return &(cssTraceTable.EntityData)
}

// CISCOSONETMIB_CssTraceTable_CssTraceEntry
// An entry in the trace table. Entries exist for active sonet
// lines. The objects in this table are used to verify 
// continued connection between the two ends of the line.
type CISCOSONETMIB_CssTraceTable_CssTraceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Sonet Section Trace To Transmit. This is string that is transmitted to
    // perform Sonet section trace  diagnostics. The trace string is  repetitively
    // transmited so that a trace receiving terminal can  verify its continued
    // connection to the intended  transmitter. The default value is a zero-length
    // string. Unless this object is set to a non-zero length string,  tracing
    // will not be performed. The type is string with length: 0 | 16 | 64.
    CssTraceToTransmit interface{}

    // Sonet Section Trace To Expect. The receiving terminal  verifies if the
    // incoming string matches this string.  The value of 'cssTraceFailure'
    // indicates whether a  trace mismatch occurred. The default value is a 
    // zero-length string. The type is string with length: 0 | 16 | 64.
    CssTraceToExpect interface{}

    // The value of this object is set to 'true' when Sonet  Section received
    // trace does not match the  'cssTraceToExpect'. The type is bool.
    CssTraceFailure interface{}

    // This object is used to view the Sonet Section Trace that  is received by
    // the receiving terminal. The type is string with length: 0 | 16 | 64.
    CssTraceReceived interface{}
}

func (cssTraceEntry *CISCOSONETMIB_CssTraceTable_CssTraceEntry) GetEntityData() *types.CommonEntityData {
    cssTraceEntry.EntityData.YFilter = cssTraceEntry.YFilter
    cssTraceEntry.EntityData.YangName = "cssTraceEntry"
    cssTraceEntry.EntityData.BundleName = "cisco_ios_xe"
    cssTraceEntry.EntityData.ParentYangName = "cssTraceTable"
    cssTraceEntry.EntityData.SegmentPath = "cssTraceEntry" + types.AddKeyToken(cssTraceEntry.IfIndex, "ifIndex")
    cssTraceEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cssTraceTable/" + cssTraceEntry.EntityData.SegmentPath
    cssTraceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cssTraceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cssTraceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cssTraceEntry.EntityData.Children = types.NewOrderedMap()
    cssTraceEntry.EntityData.Leafs = types.NewOrderedMap()
    cssTraceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cssTraceEntry.IfIndex})
    cssTraceEntry.EntityData.Leafs.Append("cssTraceToTransmit", types.YLeaf{"CssTraceToTransmit", cssTraceEntry.CssTraceToTransmit})
    cssTraceEntry.EntityData.Leafs.Append("cssTraceToExpect", types.YLeaf{"CssTraceToExpect", cssTraceEntry.CssTraceToExpect})
    cssTraceEntry.EntityData.Leafs.Append("cssTraceFailure", types.YLeaf{"CssTraceFailure", cssTraceEntry.CssTraceFailure})
    cssTraceEntry.EntityData.Leafs.Append("cssTraceReceived", types.YLeaf{"CssTraceReceived", cssTraceEntry.CssTraceReceived})

    cssTraceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cssTraceEntry.EntityData)
}

// CISCOSONETMIB_CslTotalTable
// The SONET/SDH Line Total table. It contains the 
// cumulative sum of the various statistics for the 24 
// hour period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the 
// number of 15 minute intervals that have elapsed since
// the line is enabled.
type CISCOSONETMIB_CslTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_CslTotalTable_CslTotalEntry.
    CslTotalEntry []*CISCOSONETMIB_CslTotalTable_CslTotalEntry
}

func (cslTotalTable *CISCOSONETMIB_CslTotalTable) GetEntityData() *types.CommonEntityData {
    cslTotalTable.EntityData.YFilter = cslTotalTable.YFilter
    cslTotalTable.EntityData.YangName = "cslTotalTable"
    cslTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cslTotalTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cslTotalTable.EntityData.SegmentPath = "cslTotalTable"
    cslTotalTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cslTotalTable.EntityData.SegmentPath
    cslTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cslTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cslTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cslTotalTable.EntityData.Children = types.NewOrderedMap()
    cslTotalTable.EntityData.Children.Append("cslTotalEntry", types.YChild{"CslTotalEntry", nil})
    for i := range cslTotalTable.CslTotalEntry {
        cslTotalTable.EntityData.Children.Append(types.GetSegmentPath(cslTotalTable.CslTotalEntry[i]), types.YChild{"CslTotalEntry", cslTotalTable.CslTotalEntry[i]})
    }
    cslTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cslTotalTable.EntityData.YListKeys = []string {}

    return &(cslTotalTable.EntityData)
}

// CISCOSONETMIB_CslTotalTable_CslTotalEntry
// An entry in the SONET/SDH Line Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_CslTotalTable_CslTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Line in the last
    // 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // errored seconds.
    CslTotalESs interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Line in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are severely errored seconds.
    CslTotalSESs interface{}

    // The number of Coding Violations encountered by a  SONET/SDH Line in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // coding violations.
    CslTotalCVs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH Line in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // unavailable seconds.
    CslTotalUASs interface{}
}

func (cslTotalEntry *CISCOSONETMIB_CslTotalTable_CslTotalEntry) GetEntityData() *types.CommonEntityData {
    cslTotalEntry.EntityData.YFilter = cslTotalEntry.YFilter
    cslTotalEntry.EntityData.YangName = "cslTotalEntry"
    cslTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cslTotalEntry.EntityData.ParentYangName = "cslTotalTable"
    cslTotalEntry.EntityData.SegmentPath = "cslTotalEntry" + types.AddKeyToken(cslTotalEntry.IfIndex, "ifIndex")
    cslTotalEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cslTotalTable/" + cslTotalEntry.EntityData.SegmentPath
    cslTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cslTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cslTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cslTotalEntry.EntityData.Children = types.NewOrderedMap()
    cslTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cslTotalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cslTotalEntry.IfIndex})
    cslTotalEntry.EntityData.Leafs.Append("cslTotalESs", types.YLeaf{"CslTotalESs", cslTotalEntry.CslTotalESs})
    cslTotalEntry.EntityData.Leafs.Append("cslTotalSESs", types.YLeaf{"CslTotalSESs", cslTotalEntry.CslTotalSESs})
    cslTotalEntry.EntityData.Leafs.Append("cslTotalCVs", types.YLeaf{"CslTotalCVs", cslTotalEntry.CslTotalCVs})
    cslTotalEntry.EntityData.Leafs.Append("cslTotalUASs", types.YLeaf{"CslTotalUASs", cslTotalEntry.CslTotalUASs})

    cslTotalEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cslTotalEntry.EntityData)
}

// CISCOSONETMIB_CslFarEndTotalTable
// The SONET/SDH Far End Line Total table. It contains the 
// cumulative sum of the various statistics for the 24 hour 
// period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the 
// number of 15 minute intervals that have elapsed since the 
// line is enabled.
type CISCOSONETMIB_CslFarEndTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_CslFarEndTotalTable_CslFarEndTotalEntry.
    CslFarEndTotalEntry []*CISCOSONETMIB_CslFarEndTotalTable_CslFarEndTotalEntry
}

func (cslFarEndTotalTable *CISCOSONETMIB_CslFarEndTotalTable) GetEntityData() *types.CommonEntityData {
    cslFarEndTotalTable.EntityData.YFilter = cslFarEndTotalTable.YFilter
    cslFarEndTotalTable.EntityData.YangName = "cslFarEndTotalTable"
    cslFarEndTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cslFarEndTotalTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cslFarEndTotalTable.EntityData.SegmentPath = "cslFarEndTotalTable"
    cslFarEndTotalTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cslFarEndTotalTable.EntityData.SegmentPath
    cslFarEndTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cslFarEndTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cslFarEndTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cslFarEndTotalTable.EntityData.Children = types.NewOrderedMap()
    cslFarEndTotalTable.EntityData.Children.Append("cslFarEndTotalEntry", types.YChild{"CslFarEndTotalEntry", nil})
    for i := range cslFarEndTotalTable.CslFarEndTotalEntry {
        cslFarEndTotalTable.EntityData.Children.Append(types.GetSegmentPath(cslFarEndTotalTable.CslFarEndTotalEntry[i]), types.YChild{"CslFarEndTotalEntry", cslFarEndTotalTable.CslFarEndTotalEntry[i]})
    }
    cslFarEndTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cslFarEndTotalTable.EntityData.YListKeys = []string {}

    return &(cslFarEndTotalTable.EntityData)
}

// CISCOSONETMIB_CslFarEndTotalTable_CslFarEndTotalEntry
// An entry in the SONET/SDH Far End Line Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_CslFarEndTotalTable_CslFarEndTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Far End Line in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are errored seconds.
    CslFarEndTotalESs interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Far End
    // Line in the last 24 hours. The type is interface{} with range:
    // 0..4294967295. Units are severely errored seconds.
    CslFarEndTotalSESs interface{}

    // The number of Coding Violations encountered by a SONET/SDH  Far End Line in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are coding violations.
    CslFarEndTotalCVs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH Far End Line
    // in the last 24 hours. The type is interface{} with range: 0..4294967295.
    // Units are unavailable seconds.
    CslFarEndTotalUASs interface{}
}

func (cslFarEndTotalEntry *CISCOSONETMIB_CslFarEndTotalTable_CslFarEndTotalEntry) GetEntityData() *types.CommonEntityData {
    cslFarEndTotalEntry.EntityData.YFilter = cslFarEndTotalEntry.YFilter
    cslFarEndTotalEntry.EntityData.YangName = "cslFarEndTotalEntry"
    cslFarEndTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cslFarEndTotalEntry.EntityData.ParentYangName = "cslFarEndTotalTable"
    cslFarEndTotalEntry.EntityData.SegmentPath = "cslFarEndTotalEntry" + types.AddKeyToken(cslFarEndTotalEntry.IfIndex, "ifIndex")
    cslFarEndTotalEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cslFarEndTotalTable/" + cslFarEndTotalEntry.EntityData.SegmentPath
    cslFarEndTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cslFarEndTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cslFarEndTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cslFarEndTotalEntry.EntityData.Children = types.NewOrderedMap()
    cslFarEndTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cslFarEndTotalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cslFarEndTotalEntry.IfIndex})
    cslFarEndTotalEntry.EntityData.Leafs.Append("cslFarEndTotalESs", types.YLeaf{"CslFarEndTotalESs", cslFarEndTotalEntry.CslFarEndTotalESs})
    cslFarEndTotalEntry.EntityData.Leafs.Append("cslFarEndTotalSESs", types.YLeaf{"CslFarEndTotalSESs", cslFarEndTotalEntry.CslFarEndTotalSESs})
    cslFarEndTotalEntry.EntityData.Leafs.Append("cslFarEndTotalCVs", types.YLeaf{"CslFarEndTotalCVs", cslFarEndTotalEntry.CslFarEndTotalCVs})
    cslFarEndTotalEntry.EntityData.Leafs.Append("cslFarEndTotalUASs", types.YLeaf{"CslFarEndTotalUASs", cslFarEndTotalEntry.CslFarEndTotalUASs})

    cslFarEndTotalEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cslFarEndTotalEntry.EntityData)
}

// CISCOSONETMIB_CspTotalTable
// The SONET/SDH Path Total table. It contains the cumulative 
// sum of the various statistics for the 24 hour period 
// preceding the current interval.The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the number
// of 15 minute intervals that have elapsed since the line is 
// enabled.
type CISCOSONETMIB_CspTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_CspTotalTable_CspTotalEntry.
    CspTotalEntry []*CISCOSONETMIB_CspTotalTable_CspTotalEntry
}

func (cspTotalTable *CISCOSONETMIB_CspTotalTable) GetEntityData() *types.CommonEntityData {
    cspTotalTable.EntityData.YFilter = cspTotalTable.YFilter
    cspTotalTable.EntityData.YangName = "cspTotalTable"
    cspTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cspTotalTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cspTotalTable.EntityData.SegmentPath = "cspTotalTable"
    cspTotalTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cspTotalTable.EntityData.SegmentPath
    cspTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspTotalTable.EntityData.Children = types.NewOrderedMap()
    cspTotalTable.EntityData.Children.Append("cspTotalEntry", types.YChild{"CspTotalEntry", nil})
    for i := range cspTotalTable.CspTotalEntry {
        cspTotalTable.EntityData.Children.Append(types.GetSegmentPath(cspTotalTable.CspTotalEntry[i]), types.YChild{"CspTotalEntry", cspTotalTable.CspTotalEntry[i]})
    }
    cspTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cspTotalTable.EntityData.YListKeys = []string {}

    return &(cspTotalTable.EntityData)
}

// CISCOSONETMIB_CspTotalTable_CspTotalEntry
// An entry in the SONET/SDH Path Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_CspTotalTable_CspTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Path in the last
    // 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // errored seconds.
    CspTotalESs interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Path in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are severely errored seconds.
    CspTotalSESs interface{}

    // The number of Coding Violations encountered by a  SONET/SDH Path in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // coding violations.
    CspTotalCVs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH Path in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // unavailable seconds.
    CspTotalUASs interface{}
}

func (cspTotalEntry *CISCOSONETMIB_CspTotalTable_CspTotalEntry) GetEntityData() *types.CommonEntityData {
    cspTotalEntry.EntityData.YFilter = cspTotalEntry.YFilter
    cspTotalEntry.EntityData.YangName = "cspTotalEntry"
    cspTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cspTotalEntry.EntityData.ParentYangName = "cspTotalTable"
    cspTotalEntry.EntityData.SegmentPath = "cspTotalEntry" + types.AddKeyToken(cspTotalEntry.IfIndex, "ifIndex")
    cspTotalEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cspTotalTable/" + cspTotalEntry.EntityData.SegmentPath
    cspTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspTotalEntry.EntityData.Children = types.NewOrderedMap()
    cspTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cspTotalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cspTotalEntry.IfIndex})
    cspTotalEntry.EntityData.Leafs.Append("cspTotalESs", types.YLeaf{"CspTotalESs", cspTotalEntry.CspTotalESs})
    cspTotalEntry.EntityData.Leafs.Append("cspTotalSESs", types.YLeaf{"CspTotalSESs", cspTotalEntry.CspTotalSESs})
    cspTotalEntry.EntityData.Leafs.Append("cspTotalCVs", types.YLeaf{"CspTotalCVs", cspTotalEntry.CspTotalCVs})
    cspTotalEntry.EntityData.Leafs.Append("cspTotalUASs", types.YLeaf{"CspTotalUASs", cspTotalEntry.CspTotalUASs})

    cspTotalEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cspTotalEntry.EntityData)
}

// CISCOSONETMIB_CspFarEndTotalTable
// The SONET/SDH Far End Path Total table. Far End is the 
// remote end of the line. The table contains the cumulative
// sum of the various statistics for the 24 hour period 
// preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the
// number of 15 minute intervals that have elapsed since
// the line is enabled. 
type CISCOSONETMIB_CspFarEndTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Total table.  Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_CspFarEndTotalTable_CspFarEndTotalEntry.
    CspFarEndTotalEntry []*CISCOSONETMIB_CspFarEndTotalTable_CspFarEndTotalEntry
}

func (cspFarEndTotalTable *CISCOSONETMIB_CspFarEndTotalTable) GetEntityData() *types.CommonEntityData {
    cspFarEndTotalTable.EntityData.YFilter = cspFarEndTotalTable.YFilter
    cspFarEndTotalTable.EntityData.YangName = "cspFarEndTotalTable"
    cspFarEndTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cspFarEndTotalTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cspFarEndTotalTable.EntityData.SegmentPath = "cspFarEndTotalTable"
    cspFarEndTotalTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cspFarEndTotalTable.EntityData.SegmentPath
    cspFarEndTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspFarEndTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspFarEndTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspFarEndTotalTable.EntityData.Children = types.NewOrderedMap()
    cspFarEndTotalTable.EntityData.Children.Append("cspFarEndTotalEntry", types.YChild{"CspFarEndTotalEntry", nil})
    for i := range cspFarEndTotalTable.CspFarEndTotalEntry {
        cspFarEndTotalTable.EntityData.Children.Append(types.GetSegmentPath(cspFarEndTotalTable.CspFarEndTotalEntry[i]), types.YChild{"CspFarEndTotalEntry", cspFarEndTotalTable.CspFarEndTotalEntry[i]})
    }
    cspFarEndTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cspFarEndTotalTable.EntityData.YListKeys = []string {}

    return &(cspFarEndTotalTable.EntityData)
}

// CISCOSONETMIB_CspFarEndTotalTable_CspFarEndTotalEntry
// An entry in the SONET/SDH Far End Path Total table. 
// Entries are created automatically for sonet lines.
type CISCOSONETMIB_CspFarEndTotalTable_CspFarEndTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH far end path in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are errored seconds.
    CspFarEndTotalESs interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH far end
    // path in the last 24 hours. The type is interface{} with range:
    // 0..4294967295. Units are severely errored seconds.
    CspFarEndTotalSESs interface{}

    // The number of Coding Violations encountered by a  SONET/SDH far end path in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are coding violations.
    CspFarEndTotalCVs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH far end path
    // in the last 24 hours. The type is interface{} with range: 0..4294967295.
    // Units are unavailable seconds.
    CspFarEndTotalUASs interface{}
}

func (cspFarEndTotalEntry *CISCOSONETMIB_CspFarEndTotalTable_CspFarEndTotalEntry) GetEntityData() *types.CommonEntityData {
    cspFarEndTotalEntry.EntityData.YFilter = cspFarEndTotalEntry.YFilter
    cspFarEndTotalEntry.EntityData.YangName = "cspFarEndTotalEntry"
    cspFarEndTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cspFarEndTotalEntry.EntityData.ParentYangName = "cspFarEndTotalTable"
    cspFarEndTotalEntry.EntityData.SegmentPath = "cspFarEndTotalEntry" + types.AddKeyToken(cspFarEndTotalEntry.IfIndex, "ifIndex")
    cspFarEndTotalEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cspFarEndTotalTable/" + cspFarEndTotalEntry.EntityData.SegmentPath
    cspFarEndTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspFarEndTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspFarEndTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspFarEndTotalEntry.EntityData.Children = types.NewOrderedMap()
    cspFarEndTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cspFarEndTotalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cspFarEndTotalEntry.IfIndex})
    cspFarEndTotalEntry.EntityData.Leafs.Append("cspFarEndTotalESs", types.YLeaf{"CspFarEndTotalESs", cspFarEndTotalEntry.CspFarEndTotalESs})
    cspFarEndTotalEntry.EntityData.Leafs.Append("cspFarEndTotalSESs", types.YLeaf{"CspFarEndTotalSESs", cspFarEndTotalEntry.CspFarEndTotalSESs})
    cspFarEndTotalEntry.EntityData.Leafs.Append("cspFarEndTotalCVs", types.YLeaf{"CspFarEndTotalCVs", cspFarEndTotalEntry.CspFarEndTotalCVs})
    cspFarEndTotalEntry.EntityData.Leafs.Append("cspFarEndTotalUASs", types.YLeaf{"CspFarEndTotalUASs", cspFarEndTotalEntry.CspFarEndTotalUASs})

    cspFarEndTotalEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cspFarEndTotalEntry.EntityData)
}

// CISCOSONETMIB_CspTraceTable
// The SONET/SDH Path Trace table. This table contains objects 
// for tracing the sonet path.
type CISCOSONETMIB_CspTraceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Trace table. The entries  exist for active
    // sonet lines. The objects in this table are  used to verify continued
    // connection between the two ends of the line. The type is slice of
    // CISCOSONETMIB_CspTraceTable_CspTraceEntry.
    CspTraceEntry []*CISCOSONETMIB_CspTraceTable_CspTraceEntry
}

func (cspTraceTable *CISCOSONETMIB_CspTraceTable) GetEntityData() *types.CommonEntityData {
    cspTraceTable.EntityData.YFilter = cspTraceTable.YFilter
    cspTraceTable.EntityData.YangName = "cspTraceTable"
    cspTraceTable.EntityData.BundleName = "cisco_ios_xe"
    cspTraceTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cspTraceTable.EntityData.SegmentPath = "cspTraceTable"
    cspTraceTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + cspTraceTable.EntityData.SegmentPath
    cspTraceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspTraceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspTraceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspTraceTable.EntityData.Children = types.NewOrderedMap()
    cspTraceTable.EntityData.Children.Append("cspTraceEntry", types.YChild{"CspTraceEntry", nil})
    for i := range cspTraceTable.CspTraceEntry {
        cspTraceTable.EntityData.Children.Append(types.GetSegmentPath(cspTraceTable.CspTraceEntry[i]), types.YChild{"CspTraceEntry", cspTraceTable.CspTraceEntry[i]})
    }
    cspTraceTable.EntityData.Leafs = types.NewOrderedMap()

    cspTraceTable.EntityData.YListKeys = []string {}

    return &(cspTraceTable.EntityData)
}

// CISCOSONETMIB_CspTraceTable_CspTraceEntry
// An entry in the SONET/SDH Path Trace table. The entries 
// exist for active sonet lines. The objects in this table are 
// used to verify continued connection between the two ends of
// the line.
type CISCOSONETMIB_CspTraceTable_CspTraceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Sonet Path Trace To Transmit. The trace string is  repetitively transmited
    // so that a trace receiving terminal  can verify its continued receiving
    // terminal can verify its  continued connection to the intended transmitter.
    // The  default value is a zero-length string. Unless this object  is set to a
    // non-zero length string, tracing will not be  performed. The type is string
    // with length: 0 | 16 | 64.
    CspTraceToTransmit interface{}

    // Sonet Path Trace To Expect.  The receiving terminal verifies if the
    // incoming string matches this string. The value of  'cspTraceFailure'
    // indicates whether a trace mismatch  occured. The default value is a
    // zero-length string. The type is string with length: 0 | 16 | 64.
    CspTraceToExpect interface{}

    // The value of this object is set to 'true' when Sonet Path  received trace
    // does not match the 'cspTraceToExpect'. The type is bool.
    CspTraceFailure interface{}

    // This object is used to view the Sonet Path Trace that is received by the
    // receiving terminal. The type is string with length: 0 | 16 | 64.
    CspTraceReceived interface{}
}

func (cspTraceEntry *CISCOSONETMIB_CspTraceTable_CspTraceEntry) GetEntityData() *types.CommonEntityData {
    cspTraceEntry.EntityData.YFilter = cspTraceEntry.YFilter
    cspTraceEntry.EntityData.YangName = "cspTraceEntry"
    cspTraceEntry.EntityData.BundleName = "cisco_ios_xe"
    cspTraceEntry.EntityData.ParentYangName = "cspTraceTable"
    cspTraceEntry.EntityData.SegmentPath = "cspTraceEntry" + types.AddKeyToken(cspTraceEntry.IfIndex, "ifIndex")
    cspTraceEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/cspTraceTable/" + cspTraceEntry.EntityData.SegmentPath
    cspTraceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspTraceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspTraceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspTraceEntry.EntityData.Children = types.NewOrderedMap()
    cspTraceEntry.EntityData.Leafs = types.NewOrderedMap()
    cspTraceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cspTraceEntry.IfIndex})
    cspTraceEntry.EntityData.Leafs.Append("cspTraceToTransmit", types.YLeaf{"CspTraceToTransmit", cspTraceEntry.CspTraceToTransmit})
    cspTraceEntry.EntityData.Leafs.Append("cspTraceToExpect", types.YLeaf{"CspTraceToExpect", cspTraceEntry.CspTraceToExpect})
    cspTraceEntry.EntityData.Leafs.Append("cspTraceFailure", types.YLeaf{"CspTraceFailure", cspTraceEntry.CspTraceFailure})
    cspTraceEntry.EntityData.Leafs.Append("cspTraceReceived", types.YLeaf{"CspTraceReceived", cspTraceEntry.CspTraceReceived})

    cspTraceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cspTraceEntry.EntityData)
}

// CISCOSONETMIB_CsStatsTable
// The SONET/SDH Section statistics table. This table 
// maintains the number of times the line encountered Loss of
// Signal(LOS), Loss of frame(LOF), Alarm Indication 
// signals(AISs), Remote failure indications(RFIs).
type CISCOSONETMIB_CsStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH statistics table. These are  realtime statistics
    // for the Sonet section, line and path layers. The statistics are gathered
    // for each sonet line.  An entry is created automatically and is indexed by 
    // ifIndex. The type is slice of CISCOSONETMIB_CsStatsTable_CsStatsEntry.
    CsStatsEntry []*CISCOSONETMIB_CsStatsTable_CsStatsEntry
}

func (csStatsTable *CISCOSONETMIB_CsStatsTable) GetEntityData() *types.CommonEntityData {
    csStatsTable.EntityData.YFilter = csStatsTable.YFilter
    csStatsTable.EntityData.YangName = "csStatsTable"
    csStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csStatsTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csStatsTable.EntityData.SegmentPath = "csStatsTable"
    csStatsTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + csStatsTable.EntityData.SegmentPath
    csStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csStatsTable.EntityData.Children = types.NewOrderedMap()
    csStatsTable.EntityData.Children.Append("csStatsEntry", types.YChild{"CsStatsEntry", nil})
    for i := range csStatsTable.CsStatsEntry {
        csStatsTable.EntityData.Children.Append(types.GetSegmentPath(csStatsTable.CsStatsEntry[i]), types.YChild{"CsStatsEntry", csStatsTable.CsStatsEntry[i]})
    }
    csStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csStatsTable.EntityData.YListKeys = []string {}

    return &(csStatsTable.EntityData)
}

// CISCOSONETMIB_CsStatsTable_CsStatsEntry
// An entry in the SONET/SDH statistics table. These are 
// realtime statistics for the Sonet section, line and path
// layers. The statistics are gathered for each sonet line. 
// An entry is created automatically and is indexed by 
// ifIndex.
type CISCOSONETMIB_CsStatsTable_CsStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of Loss of signals(LOS) encountered by a  SONET/SDH Section. A
    // high value for this object may  indicate a problem with the Sonet Section
    // layer. The type is interface{} with range: 0..4294967295. Units are loss of
    // signals.
    CssLOSs interface{}

    // The number of Loss of Frames (LOF) encountered by a  SONET/SDH Section. A
    // high value for this object may  indicate a problem with the Sonet Section
    // layer. The type is interface{} with range: 0..4294967295. Units are loss of
    // frames.
    CssLOFs interface{}

    // The number of alarm indication signals(AIS)  encountered by  a SONET/SDH
    // Line. A high value for this object may indicate a problem with the Sonet
    // Line layer. The type is interface{} with range: 0..4294967295. Units are
    // alarm indication signals.
    CslAISs interface{}

    // The number of remote failure indications (RFI) encountered  by a SONET/SDH
    // Line. A high value for this object may  indicate a problem with the Sonet
    // Line layer. The type is interface{} with range: 0..4294967295. Units are
    // remote failure indications.
    CslRFIs interface{}

    // The  number of alarm indication signals (AIS) encountered by a SONET/SDH
    // Path. A high value for this object may  indicate a problem with the Sonet
    // Path layer. The type is interface{} with range: 0..4294967295. Units are
    // alarm indication signals.
    CspAISs interface{}

    // The number of  remote failure indications (RFI)  encountered by a SONET/SDH
    // Path. A high value for this  object may indicate a problem with the Sonet
    // Path layer. The type is interface{} with range: 0..4294967295. Units are
    // remote failure indications.
    CspRFIs interface{}
}

func (csStatsEntry *CISCOSONETMIB_CsStatsTable_CsStatsEntry) GetEntityData() *types.CommonEntityData {
    csStatsEntry.EntityData.YFilter = csStatsEntry.YFilter
    csStatsEntry.EntityData.YangName = "csStatsEntry"
    csStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csStatsEntry.EntityData.ParentYangName = "csStatsTable"
    csStatsEntry.EntityData.SegmentPath = "csStatsEntry" + types.AddKeyToken(csStatsEntry.IfIndex, "ifIndex")
    csStatsEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/csStatsTable/" + csStatsEntry.EntityData.SegmentPath
    csStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csStatsEntry.EntityData.Children = types.NewOrderedMap()
    csStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", csStatsEntry.IfIndex})
    csStatsEntry.EntityData.Leafs.Append("cssLOSs", types.YLeaf{"CssLOSs", csStatsEntry.CssLOSs})
    csStatsEntry.EntityData.Leafs.Append("cssLOFs", types.YLeaf{"CssLOFs", csStatsEntry.CssLOFs})
    csStatsEntry.EntityData.Leafs.Append("cslAISs", types.YLeaf{"CslAISs", csStatsEntry.CslAISs})
    csStatsEntry.EntityData.Leafs.Append("cslRFIs", types.YLeaf{"CslRFIs", csStatsEntry.CslRFIs})
    csStatsEntry.EntityData.Leafs.Append("cspAISs", types.YLeaf{"CspAISs", csStatsEntry.CspAISs})
    csStatsEntry.EntityData.Leafs.Append("cspRFIs", types.YLeaf{"CspRFIs", csStatsEntry.CspRFIs})

    csStatsEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(csStatsEntry.EntityData)
}

// CISCOSONETMIB_CsAu4Tug3ConfigTable
// This table contains objects to configure the VC( Virtual
// Container) related properties of a TUG-3 within a AU-4 
// paths.
// 
// This table allows creation of
// following multiplexing structure:
// STM-1/AU-4/TUG-3/TU-3/DS3
// STM-1/AU-4/TUG-3/TU-3/E3
// STM-1/AU-4/TUG-3/TUG-2/TU-11/DS1
// STM-1/AU-4/TUG-3/TUG-2/TU-12/E1
// 
// Three entries are created in this table for a given AU-4 
// path when cspSonetPathPayload object is set to one of the 
// following:
//     vt15vc11(4),
//     vt2vc12(5),
//     ds3(3),
//     e3(8),
//     vtStructured(9)
type CISCOSONETMIB_CsAu4Tug3ConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // There is an entry in this table for each TUG-3 within a  AU-4 SDH path that
    // supports SDH virtual container VC-4. The ifIndex value represents an entry
    // in ifTable with ifType = sonetPath(50).The ifTable entry applicable for
    // this entry belongs to AU-4 path. The type is slice of
    // CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry.
    CsAu4Tug3ConfigEntry []*CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry
}

func (csAu4Tug3ConfigTable *CISCOSONETMIB_CsAu4Tug3ConfigTable) GetEntityData() *types.CommonEntityData {
    csAu4Tug3ConfigTable.EntityData.YFilter = csAu4Tug3ConfigTable.YFilter
    csAu4Tug3ConfigTable.EntityData.YangName = "csAu4Tug3ConfigTable"
    csAu4Tug3ConfigTable.EntityData.BundleName = "cisco_ios_xe"
    csAu4Tug3ConfigTable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csAu4Tug3ConfigTable.EntityData.SegmentPath = "csAu4Tug3ConfigTable"
    csAu4Tug3ConfigTable.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/" + csAu4Tug3ConfigTable.EntityData.SegmentPath
    csAu4Tug3ConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csAu4Tug3ConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csAu4Tug3ConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csAu4Tug3ConfigTable.EntityData.Children = types.NewOrderedMap()
    csAu4Tug3ConfigTable.EntityData.Children.Append("csAu4Tug3ConfigEntry", types.YChild{"CsAu4Tug3ConfigEntry", nil})
    for i := range csAu4Tug3ConfigTable.CsAu4Tug3ConfigEntry {
        csAu4Tug3ConfigTable.EntityData.Children.Append(types.GetSegmentPath(csAu4Tug3ConfigTable.CsAu4Tug3ConfigEntry[i]), types.YChild{"CsAu4Tug3ConfigEntry", csAu4Tug3ConfigTable.CsAu4Tug3ConfigEntry[i]})
    }
    csAu4Tug3ConfigTable.EntityData.Leafs = types.NewOrderedMap()

    csAu4Tug3ConfigTable.EntityData.YListKeys = []string {}

    return &(csAu4Tug3ConfigTable.EntityData)
}

// CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry
// There is an entry in this table for each TUG-3 within a 
// AU-4 SDH path that supports SDH virtual container VC-4.
// The ifIndex value represents an entry in ifTable with
// ifType = sonetPath(50).The ifTable entry applicable for
// this entry belongs to AU-4 path.
type CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This object represents the TUG-3 number. The type
    // is interface{} with range: 1..3.
    CsAu4Tug3 interface{}

    // This object is used for configuring the payload for the tributary group. 
    // The possible values are :  vc11   : When set to 'vc11' following things are
    // done:        - 28 entries created in ifTable for TU-11 with          
    // ifType = sonetVT(51)        - 28 entries created in ifTable for DS1 with   
    // ifType = ds1(18)           STM1<-AU-4<-TUG-3<-TUG-2<-TU-11<-VC11  vc12   :
    // When set to 'vc12' following things are done:        - 21 entries created
    // in ifTable for TU-12 with           ifType = sonetVT(51)        - 21
    // entries created in ifTable for E1 with           ifType = ds1(18)          
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-12<-VC12  tu3ds3 : When set to 'tu3ds3'
    // following things are done:        - 1 entry created in ifTable for TU-3
    // with           ifType = sonetVT(51)        - 1 entry created in ifTable for
    // DS3 with           ifType = ds3(30)           STM1<-AU-4<-TUG-3<-TU-3<-VC3 
    // tu3e3  : When set to 'tu3e3' following things are done:        - 1 entry
    // created in ifTable for TU-3 with           ifType = sonetVT(51)        - 1
    // entry created in ifTable for E3 with           ifType = ds3(30)          
    // STM1<-AU-4<-TUG-3<-TU-3<-VC3  The value 'other' can not be set. The type is
    // CsAu4Tug3Payload.
    CsAu4Tug3Payload interface{}
}

func (csAu4Tug3ConfigEntry *CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry) GetEntityData() *types.CommonEntityData {
    csAu4Tug3ConfigEntry.EntityData.YFilter = csAu4Tug3ConfigEntry.YFilter
    csAu4Tug3ConfigEntry.EntityData.YangName = "csAu4Tug3ConfigEntry"
    csAu4Tug3ConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    csAu4Tug3ConfigEntry.EntityData.ParentYangName = "csAu4Tug3ConfigTable"
    csAu4Tug3ConfigEntry.EntityData.SegmentPath = "csAu4Tug3ConfigEntry" + types.AddKeyToken(csAu4Tug3ConfigEntry.IfIndex, "ifIndex") + types.AddKeyToken(csAu4Tug3ConfigEntry.CsAu4Tug3, "csAu4Tug3")
    csAu4Tug3ConfigEntry.EntityData.AbsolutePath = "CISCO-SONET-MIB:CISCO-SONET-MIB/csAu4Tug3ConfigTable/" + csAu4Tug3ConfigEntry.EntityData.SegmentPath
    csAu4Tug3ConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csAu4Tug3ConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csAu4Tug3ConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csAu4Tug3ConfigEntry.EntityData.Children = types.NewOrderedMap()
    csAu4Tug3ConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    csAu4Tug3ConfigEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", csAu4Tug3ConfigEntry.IfIndex})
    csAu4Tug3ConfigEntry.EntityData.Leafs.Append("csAu4Tug3", types.YLeaf{"CsAu4Tug3", csAu4Tug3ConfigEntry.CsAu4Tug3})
    csAu4Tug3ConfigEntry.EntityData.Leafs.Append("csAu4Tug3Payload", types.YLeaf{"CsAu4Tug3Payload", csAu4Tug3ConfigEntry.CsAu4Tug3Payload})

    csAu4Tug3ConfigEntry.EntityData.YListKeys = []string {"IfIndex", "CsAu4Tug3"}

    return &(csAu4Tug3ConfigEntry.EntityData)
}

// CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload represents The value 'other' can not be set.
type CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload string

const (
    CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload_other CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload = "other"

    CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload_vc11 CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload = "vc11"

    CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload_vc12 CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload = "vc12"

    CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload_tu3ds3 CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload = "tu3ds3"

    CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload_tu3e3 CISCOSONETMIB_CsAu4Tug3ConfigTable_CsAu4Tug3ConfigEntry_CsAu4Tug3Payload = "tu3e3"
)

