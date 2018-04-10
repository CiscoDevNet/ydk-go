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

// CsApsLineFailureCode represents     csApsModeMismatch:        APS architecture mode mismatch.
type CsApsLineFailureCode string

const (
    CsApsLineFailureCode_csApsChannelMismatch CsApsLineFailureCode = "csApsChannelMismatch"

    CsApsLineFailureCode_csApsProtectionByteFail CsApsLineFailureCode = "csApsProtectionByteFail"

    CsApsLineFailureCode_csApsFEProtectionFailure CsApsLineFailureCode = "csApsFEProtectionFailure"

    CsApsLineFailureCode_csApsModeMismatch CsApsLineFailureCode = "csApsModeMismatch"
)

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

// CISCOSONETMIB
type CISCOSONETMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Csapsconfig CISCOSONETMIB_Csapsconfig

    
    Csnotifications CISCOSONETMIB_Csnotifications

    // The SONET/SDH configuration table. This table has objects  for configuring
    // sonet lines.
    Csconfigtable CISCOSONETMIB_Csconfigtable

    // This table contains objects to configure APS  (Automatic Protection
    // Switching) feature in a SONET  Line. APS is the ability to configure a pair
    // of SONET  lines for redundancy so that the hardware will  automatically
    // switch the active line from working line to the protection line or vice
    // versa, within 60ms,  when the active line fails.
    Csapsconfigtable CISCOSONETMIB_Csapsconfigtable

    // The SONET/SDH Section Total table. It contains the  cumulative sum of the
    // various statistics for the 24 hour period preceding the current interval.
    // The object  'sonetMediumValidIntervals' from RFC2558 contains the number of
    // 15 minute intervals that have elapsed since the line is enabled. .
    Csstotaltable CISCOSONETMIB_Csstotaltable

    // The SONET/SDH Section Trace table. This table contains  objects for tracing
    // the sonet section.
    Csstracetable CISCOSONETMIB_Csstracetable

    // The SONET/SDH Line Total table. It contains the  cumulative sum of the
    // various statistics for the 24  hour period preceding the current interval.
    // The object  'sonetMediumValidIntervals' from RFC2558 contains the  number
    // of 15 minute intervals that have elapsed since the line is enabled.
    Csltotaltable CISCOSONETMIB_Csltotaltable

    // The SONET/SDH Far End Line Total table. It contains the  cumulative sum of
    // the various statistics for the 24 hour  period preceding the current
    // interval. The object  'sonetMediumValidIntervals' from RFC2558 contains the
    // number of 15 minute intervals that have elapsed since the  line is enabled.
    Cslfarendtotaltable CISCOSONETMIB_Cslfarendtotaltable

    // The SONET/SDH Path Total table. It contains the cumulative  sum of the
    // various statistics for the 24 hour period  preceding the current
    // interval.The object  'sonetMediumValidIntervals' from RFC2558 contains the
    // number of 15 minute intervals that have elapsed since the line is  enabled.
    Csptotaltable CISCOSONETMIB_Csptotaltable

    // The SONET/SDH Far End Path Total table. Far End is the  remote end of the
    // line. The table contains the cumulative sum of the various statistics for
    // the 24 hour period  preceding the current interval. The object 
    // 'sonetMediumValidIntervals' from RFC2558 contains the number of 15 minute
    // intervals that have elapsed since the line is enabled. .
    Cspfarendtotaltable CISCOSONETMIB_Cspfarendtotaltable

    // The SONET/SDH Path Trace table. This table contains objects  for tracing
    // the sonet path.
    Csptracetable CISCOSONETMIB_Csptracetable

    // The SONET/SDH Section statistics table. This table  maintains the number of
    // times the line encountered Loss of Signal(LOS), Loss of frame(LOF), Alarm
    // Indication  signals(AISs), Remote failure indications(RFIs).
    Csstatstable CISCOSONETMIB_Csstatstable

    // This table contains objects to configure the VC( Virtual Container) related
    // properties of a TUG-3 within a AU-4  paths.  This table allows creation of
    // following multiplexing structure: STM-1/AU-4/TUG-3/TU-3/DS3
    // STM-1/AU-4/TUG-3/TU-3/E3 STM-1/AU-4/TUG-3/TUG-2/TU-11/DS1
    // STM-1/AU-4/TUG-3/TUG-2/TU-12/E1  Three entries are created in this table
    // for a given AU-4  path when cspSonetPathPayload object is set to one of the
    // following:     vt15vc11(4),     vt2vc12(5),     ds3(3),     e3(8),    
    // vtStructured(9).
    Csau4Tug3Configtable CISCOSONETMIB_Csau4Tug3Configtable
}

func (cISCOSONETMIB *CISCOSONETMIB) GetEntityData() *types.CommonEntityData {
    cISCOSONETMIB.EntityData.YFilter = cISCOSONETMIB.YFilter
    cISCOSONETMIB.EntityData.YangName = "CISCO-SONET-MIB"
    cISCOSONETMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSONETMIB.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cISCOSONETMIB.EntityData.SegmentPath = "CISCO-SONET-MIB:CISCO-SONET-MIB"
    cISCOSONETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSONETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSONETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSONETMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOSONETMIB.EntityData.Children["csApsConfig"] = types.YChild{"Csapsconfig", &cISCOSONETMIB.Csapsconfig}
    cISCOSONETMIB.EntityData.Children["csNotifications"] = types.YChild{"Csnotifications", &cISCOSONETMIB.Csnotifications}
    cISCOSONETMIB.EntityData.Children["csConfigTable"] = types.YChild{"Csconfigtable", &cISCOSONETMIB.Csconfigtable}
    cISCOSONETMIB.EntityData.Children["csApsConfigTable"] = types.YChild{"Csapsconfigtable", &cISCOSONETMIB.Csapsconfigtable}
    cISCOSONETMIB.EntityData.Children["cssTotalTable"] = types.YChild{"Csstotaltable", &cISCOSONETMIB.Csstotaltable}
    cISCOSONETMIB.EntityData.Children["cssTraceTable"] = types.YChild{"Csstracetable", &cISCOSONETMIB.Csstracetable}
    cISCOSONETMIB.EntityData.Children["cslTotalTable"] = types.YChild{"Csltotaltable", &cISCOSONETMIB.Csltotaltable}
    cISCOSONETMIB.EntityData.Children["cslFarEndTotalTable"] = types.YChild{"Cslfarendtotaltable", &cISCOSONETMIB.Cslfarendtotaltable}
    cISCOSONETMIB.EntityData.Children["cspTotalTable"] = types.YChild{"Csptotaltable", &cISCOSONETMIB.Csptotaltable}
    cISCOSONETMIB.EntityData.Children["cspFarEndTotalTable"] = types.YChild{"Cspfarendtotaltable", &cISCOSONETMIB.Cspfarendtotaltable}
    cISCOSONETMIB.EntityData.Children["cspTraceTable"] = types.YChild{"Csptracetable", &cISCOSONETMIB.Csptracetable}
    cISCOSONETMIB.EntityData.Children["csStatsTable"] = types.YChild{"Csstatstable", &cISCOSONETMIB.Csstatstable}
    cISCOSONETMIB.EntityData.Children["csAu4Tug3ConfigTable"] = types.YChild{"Csau4Tug3Configtable", &cISCOSONETMIB.Csau4Tug3Configtable}
    cISCOSONETMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOSONETMIB.EntityData)
}

// CISCOSONETMIB_Csapsconfig
type CISCOSONETMIB_Csapsconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The object indicates the APS line failure code. This is the failure
    // encountered by the APS line. Refer to CsApsLineFailureCode TC for failure
    // code definitions. The object is used for notifications. The type is
    // CsApsLineFailureCode.
    Csapslinefailurecode interface{}

    // This object indicates the APS line switch reason. When the working line on
    // one end fails, its other end is told to do an APS switch.  Refer to
    // CsApsLineSwitchReason TC for more information. The object is used for
    // notifications. The type is CsApsLineSwitchReason.
    Csapslineswitchreason interface{}
}

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetEntityData() *types.CommonEntityData {
    csapsconfig.EntityData.YFilter = csapsconfig.YFilter
    csapsconfig.EntityData.YangName = "csApsConfig"
    csapsconfig.EntityData.BundleName = "cisco_ios_xe"
    csapsconfig.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csapsconfig.EntityData.SegmentPath = "csApsConfig"
    csapsconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csapsconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csapsconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csapsconfig.EntityData.Children = make(map[string]types.YChild)
    csapsconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    csapsconfig.EntityData.Leafs["csApsLineFailureCode"] = types.YLeaf{"Csapslinefailurecode", csapsconfig.Csapslinefailurecode}
    csapsconfig.EntityData.Leafs["csApsLineSwitchReason"] = types.YLeaf{"Csapslineswitchreason", csapsconfig.Csapslineswitchreason}
    return &(csapsconfig.EntityData)
}

// CISCOSONETMIB_Csnotifications
type CISCOSONETMIB_Csnotifications struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object controls if the generation of  ciscoSonetSectionStatusChange,
    // ciscoSonetLineStatusChange,  ciscoSonetPathStatusChange and
    // ciscoSonetVTStatusChange notifications is enabled. If the value of this
    // object is 'true(1)', then all notifications in this MIB are enabled;
    // otherwise they are disabled. The type is bool.
    Csnotificationsenabled interface{}
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetEntityData() *types.CommonEntityData {
    csnotifications.EntityData.YFilter = csnotifications.YFilter
    csnotifications.EntityData.YangName = "csNotifications"
    csnotifications.EntityData.BundleName = "cisco_ios_xe"
    csnotifications.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csnotifications.EntityData.SegmentPath = "csNotifications"
    csnotifications.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csnotifications.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csnotifications.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csnotifications.EntityData.Children = make(map[string]types.YChild)
    csnotifications.EntityData.Leafs = make(map[string]types.YLeaf)
    csnotifications.EntityData.Leafs["csNotificationsEnabled"] = types.YLeaf{"Csnotificationsenabled", csnotifications.Csnotificationsenabled}
    return &(csnotifications.EntityData)
}

// CISCOSONETMIB_Csconfigtable
// The SONET/SDH configuration table. This table has objects 
// for configuring sonet lines.
type CISCOSONETMIB_Csconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table. There is an entry for each SONET line  in the table.
    // Entries are automatically created for an  ifType value of sonet(39).
    // 'ifAdminStatus' from the ifTable  must be used to enable or disable a line.
    // A line is in  disabled(down) state unless provisioned 'up' using 
    // 'ifAdminStatus'. The type is slice of
    // CISCOSONETMIB_Csconfigtable_Csconfigentry.
    Csconfigentry []CISCOSONETMIB_Csconfigtable_Csconfigentry
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetEntityData() *types.CommonEntityData {
    csconfigtable.EntityData.YFilter = csconfigtable.YFilter
    csconfigtable.EntityData.YangName = "csConfigTable"
    csconfigtable.EntityData.BundleName = "cisco_ios_xe"
    csconfigtable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csconfigtable.EntityData.SegmentPath = "csConfigTable"
    csconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csconfigtable.EntityData.Children = make(map[string]types.YChild)
    csconfigtable.EntityData.Children["csConfigEntry"] = types.YChild{"Csconfigentry", nil}
    for i := range csconfigtable.Csconfigentry {
        csconfigtable.EntityData.Children[types.GetSegmentPath(&csconfigtable.Csconfigentry[i])] = types.YChild{"Csconfigentry", &csconfigtable.Csconfigentry[i]}
    }
    csconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csconfigtable.EntityData)
}

// CISCOSONETMIB_Csconfigtable_Csconfigentry
// An entry in the table. There is an entry for each SONET line 
// in the table. Entries are automatically created for an 
// ifType value of sonet(39). 'ifAdminStatus' from the ifTable 
// must be used to enable or disable a line. A line is in 
// disabled(down) state unless provisioned 'up' using 
// 'ifAdminStatus'.
type CISCOSONETMIB_Csconfigtable_Csconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

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
    // medium. The type is Csconfigloopbacktype.
    Csconfigloopbacktype interface{}

    // Specifies the source of the transmit clock.  loopTiming: indicates that the
    // recovered receive              clock is used as the transmit clock. 
    // localTiming: indicates that a local clock source is              used or
    // that an external clock is               attached to the box containing the 
    // interface. . The type is Csconfigxmtclocksource.
    Csconfigxmtclocksource interface{}

    // This object is used to disable or enable the Scrambling option in SONET
    // line. The type is Csconfigframescramble.
    Csconfigframescramble interface{}

    // This object represents the configured line type. Sts is SONET format. Stm
    // is SDH format.      sonetSts3c   : OC3 concatenated     sonetStm1    :
    // European standard OC3     sonetSts12c  : OC12 concatenated     sonetStm4   
    // : European standard OC12     sonetSts48c  : OC48 concatenated    
    // sonetStm16   : European standard OC48      sonetSts192c : OC-192
    // concatenated     sonetStm64   : European standard OC-192     sonetSts3    :
    // OC3 (unconcatenated)     . The type is Csconfigtype.
    Csconfigtype interface{}

    // This object specifies the type of RDI-V (Remote Defect Indication - Virtual
    // Tributary/Container) sent by this  Network Element (NE) to the remote
    // Network Element.        onebit     : use 1 bit RDI-V       threebit   : use
    // 3 bit enhanced RDI-V.  Default is onebit. The type is Csconfigrdivtype.
    Csconfigrdivtype interface{}

    // This object represents the type of RDI-P (Remote Defect Indication - Path)
    // sent by this Network Element (NE) to remote Network Element.        onebit 
    // : use 1 bit RDI-P       threebit   : use 3 bit enhanced RDI-P.  Default is
    // onebit. The type is Csconfigrdiptype.
    Csconfigrdiptype interface{}

    // Type of the tributary carried within the SONET/SDH signal.  vt15vc11    :
    // carries T1 signal vt2vc12     : carries E1 signal. The type is
    // Cstributarytype.
    Cstributarytype interface{}

    // This object represents the VT/VC mapping type.  asynchronous:    In this
    // mode, the channel structure of                   DS1/E1 is neither visible
    // nor preserved.  byteSynchronous: In this mode, the DS0 signals inside the  
    // VT/VC can be found and extracted from the                  frame.  Default
    // is asynchronous(1). The type is Cstributarymappingtype.
    Cstributarymappingtype interface{}

    // This object represents the framing type to be assigned to the virtual
    // tributaries in byte sync mapping mode.        notApplicable  :  If VT
    // mapping is not byteSynchronous(2).       dsx1ESF        :  Extended
    // Superframe Format       dsx1D4         :  Superframe Format  Default is
    // dsx1ESF(3) if csTributaryMappingType is  byteSynchronous(2). For
    // asynchronous(1) mapping, the default is  notApplicable(1).  The value
    // notApplicable(1) can not be set. The type is Cstributaryframingtype.
    Cstributaryframingtype interface{}

    // This object represents the mode used to transport DS0  signalling
    // information for T1 byteSynchronous mapping (GR253). In
    // signallingTransferMode(2), the robbed-bit signalling is  transferred to the
    // VT header. In clearMode(3), only the  framing bit is transferred to the VT
    // header.       Default is signallingTransferMode(2) if
    // csTributaryMappingType  is byteSynchronous. For asynchronous mapping, it is
    // notApplicable(1).  The value notApplicable(1) can not be set. The type is
    // Cssignallingtransportmode.
    Cssignallingtransportmode interface{}

    // This object represents the method used to group VCs into an STM-1 signal.
    // Applicable only to SDH.    au3Grouping: STM1<-AU-3<-TUG-2<-TU-12<-VC12 or  
    // STM1<-AU-3<-TUG-2<-TU-11<-VC11.    au4Grouping:
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-12<-VC12 or               
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-11<-VC11.  Default is au3Grouping(2) for SDH
    // and notApplicable(1) for SONET. The type is Cstributarygroupingtype.
    Cstributarygroupingtype interface{}
}

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetEntityData() *types.CommonEntityData {
    csconfigentry.EntityData.YFilter = csconfigentry.YFilter
    csconfigentry.EntityData.YangName = "csConfigEntry"
    csconfigentry.EntityData.BundleName = "cisco_ios_xe"
    csconfigentry.EntityData.ParentYangName = "csConfigTable"
    csconfigentry.EntityData.SegmentPath = "csConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", csconfigentry.Ifindex) + "']"
    csconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csconfigentry.EntityData.Children = make(map[string]types.YChild)
    csconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csconfigentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csconfigentry.Ifindex}
    csconfigentry.EntityData.Leafs["csConfigLoopbackType"] = types.YLeaf{"Csconfigloopbacktype", csconfigentry.Csconfigloopbacktype}
    csconfigentry.EntityData.Leafs["csConfigXmtClockSource"] = types.YLeaf{"Csconfigxmtclocksource", csconfigentry.Csconfigxmtclocksource}
    csconfigentry.EntityData.Leafs["csConfigFrameScramble"] = types.YLeaf{"Csconfigframescramble", csconfigentry.Csconfigframescramble}
    csconfigentry.EntityData.Leafs["csConfigType"] = types.YLeaf{"Csconfigtype", csconfigentry.Csconfigtype}
    csconfigentry.EntityData.Leafs["csConfigRDIVType"] = types.YLeaf{"Csconfigrdivtype", csconfigentry.Csconfigrdivtype}
    csconfigentry.EntityData.Leafs["csConfigRDIPType"] = types.YLeaf{"Csconfigrdiptype", csconfigentry.Csconfigrdiptype}
    csconfigentry.EntityData.Leafs["csTributaryType"] = types.YLeaf{"Cstributarytype", csconfigentry.Cstributarytype}
    csconfigentry.EntityData.Leafs["csTributaryMappingType"] = types.YLeaf{"Cstributarymappingtype", csconfigentry.Cstributarymappingtype}
    csconfigentry.EntityData.Leafs["csTributaryFramingType"] = types.YLeaf{"Cstributaryframingtype", csconfigentry.Cstributaryframingtype}
    csconfigentry.EntityData.Leafs["csSignallingTransportMode"] = types.YLeaf{"Cssignallingtransportmode", csconfigentry.Cssignallingtransportmode}
    csconfigentry.EntityData.Leafs["csTributaryGroupingType"] = types.YLeaf{"Cstributarygroupingtype", csconfigentry.Cstributarygroupingtype}
    return &(csconfigentry.EntityData)
}

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigframescramble represents option in SONET line.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigframescramble string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigframescramble_disabled CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigframescramble = "disabled"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigframescramble_enabled CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigframescramble = "enabled"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype represents         medium.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype_noLoopback CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype = "noLoopback"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype_lineLocal CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype = "lineLocal"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype_lineRemote CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigloopbacktype = "lineRemote"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdiptype represents Default is onebit.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdiptype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdiptype_onebit CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdiptype = "onebit"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdiptype_threebit CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdiptype = "threebit"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdivtype represents Default is onebit.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdivtype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdivtype_onebit CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdivtype = "onebit"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdivtype_threebit CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigrdivtype = "threebit"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype represents     
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetSts3c CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetSts3c"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetStm1 CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetStm1"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetSts12c CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetSts12c"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetStm4 CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetStm4"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetSts48c CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetSts48c"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetStm16 CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetStm16"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetSts192c CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetSts192c"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetStm64 CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetStm64"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype_sonetSts3 CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigtype = "sonetSts3"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigxmtclocksource represents              interface. 
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigxmtclocksource string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigxmtclocksource_loopTiming CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigxmtclocksource = "loopTiming"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigxmtclocksource_localTiming CISCOSONETMIB_Csconfigtable_Csconfigentry_Csconfigxmtclocksource = "localTiming"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode represents The value notApplicable(1) can not be set.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode_notApplicable CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode = "notApplicable"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode_signallingTransferMode CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode = "signallingTransferMode"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode_clearMode CISCOSONETMIB_Csconfigtable_Csconfigentry_Cssignallingtransportmode = "clearMode"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype represents The value notApplicable(1) can not be set.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype_notApplicable CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype = "notApplicable"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype_dsx1D4 CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype = "dsx1D4"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype_dsx1ESF CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributaryframingtype = "dsx1ESF"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype represents Default is au3Grouping(2) for SDH and notApplicable(1) for SONET.
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype_notApplicable CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype = "notApplicable"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype_au3Grouping CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype = "au3Grouping"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype_au4Grouping CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarygroupingtype = "au4Grouping"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarymappingtype represents Default is asynchronous(1).
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarymappingtype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarymappingtype_asynchronous CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarymappingtype = "asynchronous"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarymappingtype_byteSynchronous CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarymappingtype = "byteSynchronous"
)

// CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarytype represents vt2vc12     : carries E1 signal
type CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarytype string

const (
    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarytype_vt15vc11 CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarytype = "vt15vc11"

    CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarytype_vt2vc12 CISCOSONETMIB_Csconfigtable_Csconfigentry_Cstributarytype = "vt2vc12"
)

// CISCOSONETMIB_Csapsconfigtable
// This table contains objects to configure APS 
// (Automatic Protection Switching) feature in a SONET 
// Line. APS is the ability to configure a pair of SONET 
// lines for redundancy so that the hardware will 
// automatically switch the active line from working line
// to the protection line or vice versa, within 60ms, 
// when the active line fails.
type CISCOSONETMIB_Csapsconfigtable struct {
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
    // CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry.
    Csapsconfigentry []CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry
}

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetEntityData() *types.CommonEntityData {
    csapsconfigtable.EntityData.YFilter = csapsconfigtable.YFilter
    csapsconfigtable.EntityData.YangName = "csApsConfigTable"
    csapsconfigtable.EntityData.BundleName = "cisco_ios_xe"
    csapsconfigtable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csapsconfigtable.EntityData.SegmentPath = "csApsConfigTable"
    csapsconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csapsconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csapsconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csapsconfigtable.EntityData.Children = make(map[string]types.YChild)
    csapsconfigtable.EntityData.Children["csApsConfigEntry"] = types.YChild{"Csapsconfigentry", nil}
    for i := range csapsconfigtable.Csapsconfigentry {
        csapsconfigtable.EntityData.Children[types.GetSegmentPath(&csapsconfigtable.Csapsconfigentry[i])] = types.YChild{"Csapsconfigentry", &csapsconfigtable.Csapsconfigentry[i]}
    }
    csapsconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csapsconfigtable.EntityData)
}

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry
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
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. When a pair of APS lines is configured, one line
    // has to be the working line, which is the primary line, and the other has to
    // be the protection line, which is the backup line. This object refers to the
    // working line in the APS pair. For G.783 AnnexB, this index refers to
    // Working Section 1. The type is interface{} with range: 1..2147483647.
    Csapsworkingindex interface{}

    // The protection line indicates that it will become the active line when an
    // APS switch occurs (APS switch could occur because of a failure on the
    // working line). For G.783 AnnexB, This index refers to Working Section 2.
    // The type is interface{} with range: 1..2147483647.
    Csapsprotectionindex interface{}

    // This object is used to enable or disable the APS feature on the
    // working/protection line pairs. When enabled, the hardware will
    // automatically switch the active line  from the working line to the
    // protection line within 60ms, or vice versa. The type is Csapsenable.
    Csapsenable interface{}

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
    // Csapsarchmode.
    Csapsarchmode interface{}

    // This object indicates which line is currently active.  It could be the
    // working line(Section 1 for Annex B), the protection line(Section 2 for
    // Annex B) or none if neither working nor protection line is active.  This
    // object reflects the status of receive direction. The type is
    // Csapsactiveline.
    Csapsactiveline interface{}

    // This object contains the Bit Error Rate threshold for Signal Fault
    // detection on the working line. Once this threshold is exceeded, an APS
    // switch will occur. This value is 10 to the -n, where n is between 3 and 5.
    // The type is interface{} with range: 3..5.
    Csapssigfaultber interface{}

    // This object contains the Bit Error Rate threshold for Signal Degrade
    // detection on the working line. Once this threshold is exceeded, an APS
    // switch will occur. This value is 10 to -n where n is between 5 and 9. The
    // type is interface{} with range: 5..9.
    Csapssigdegradeber interface{}

    // This object contains interval in minutes to wait  before attempting to
    // switch back to working line.  Not applicable if the line is configured in 
    // non-revertive mode, i.e. protection line will  continue to be active, even
    // if failures on the  working line are cleared. The framer clears the 
    // signal-fault and signal-degrade when APS switch  occurs. Please refer to
    // 'csApsRevertive' for  description of non-revertive. The type is interface{}
    // with range: 1..12. Units are minutes.
    Csapswaittorestore interface{}

    // This object is used to configure the switching  direction which this APS
    // line supports.   Unidirectional : APS switch only in one direction.
    // Bidirectional  : APS switch in both ends of the line. The type is
    // Csapsdirection.
    Csapsdirection interface{}

    // This object is used to configure the APS revertive or nonrevertive option. 
    // revertive :    Will switch the working line back to active state after  
    // the Wait-To-restore interval has expired and the    working line
    // Signal-Fault/Signal-Degrade has been    cleared. Please refer to
    // 'csApsWaitToRestore' for    description of Wait-To-Restore interval.
    // nonrevertive :    The  protection line continues to be the active line,  
    // The active line does not switch to the working line. The type is
    // Csapsrevertive.
    Csapsrevertive interface{}

    // This object shows the actual APS direction that is  implemented on the Near
    // End terminal. APS direction  configured through csApsDirection is
    // negotiated with the Far End and APS direction setting acceptable to  both
    // ends is operational at the Near End. The type is Csapsdirectionoperational.
    Csapsdirectionoperational interface{}

    // This object shows the actual APS architecture mode that is implemented on
    // the Near End terminal. APS architecture mode configured through
    // csApsArchMode object is  negotiated with the Far End through APS channel. 
    // Architecture mode acceptable to both the Near End and  the Far End
    // terminals is then operational at the Near  End. This value can be different
    // than the APS  Architecture mode configured. The type is
    // Csapsarchmodeoperational.
    Csapsarchmodeoperational interface{}

    // This object allows to configure APS channel protocol to  be implemented at
    // Near End terminal.  K1 and K2 overhead bytes in a SONET signal are used as
    // an APS channel. This channel is used to carry APS protocol.  Possible
    // values: bellcore(1) : Implements APS channel protocol as defined           
    // in bellcore document GR-253-CORE. itu(2) : Implements APS channel protocol
    // as defined in           ITU document G.783. The type is
    // Csapschannelprotocol.
    Csapschannelprotocol interface{}

    // This object indicates APS line failure status. The type is map[string]bool.
    Csapsfailurestatus interface{}

    // This object indicates APS line switch reason. The type is
    // CsApsLineSwitchReason.
    Csapsswitchreason interface{}

    // This object indicates which working section is the APS primary section. In
    // G.783 AnnexB, the K1/K2 Bytes are received on the secondary Section. All
    // the Switch Requests are for a switch from the primary section to the
    // secondary section. The object csApsActiveline will indicate which section
    // is currently carrying the traffic.  Once the switch request clears
    // normally, traffic is maintained on the section to which it was switched by
    // making that section the primary section.   Possible values: 
    // workingSection1(1): Working Section 1 is Primary Section
    // workingSection2(2): Working Section 2 is Primary Section none(3)          
    // : none. The type is Csapsprimarysection.
    Csapsprimarysection interface{}
}

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetEntityData() *types.CommonEntityData {
    csapsconfigentry.EntityData.YFilter = csapsconfigentry.YFilter
    csapsconfigentry.EntityData.YangName = "csApsConfigEntry"
    csapsconfigentry.EntityData.BundleName = "cisco_ios_xe"
    csapsconfigentry.EntityData.ParentYangName = "csApsConfigTable"
    csapsconfigentry.EntityData.SegmentPath = "csApsConfigEntry" + "[csApsWorkingIndex='" + fmt.Sprintf("%v", csapsconfigentry.Csapsworkingindex) + "']"
    csapsconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csapsconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csapsconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csapsconfigentry.EntityData.Children = make(map[string]types.YChild)
    csapsconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csapsconfigentry.EntityData.Leafs["csApsWorkingIndex"] = types.YLeaf{"Csapsworkingindex", csapsconfigentry.Csapsworkingindex}
    csapsconfigentry.EntityData.Leafs["csApsProtectionIndex"] = types.YLeaf{"Csapsprotectionindex", csapsconfigentry.Csapsprotectionindex}
    csapsconfigentry.EntityData.Leafs["csApsEnable"] = types.YLeaf{"Csapsenable", csapsconfigentry.Csapsenable}
    csapsconfigentry.EntityData.Leafs["csApsArchMode"] = types.YLeaf{"Csapsarchmode", csapsconfigentry.Csapsarchmode}
    csapsconfigentry.EntityData.Leafs["csApsActiveLine"] = types.YLeaf{"Csapsactiveline", csapsconfigentry.Csapsactiveline}
    csapsconfigentry.EntityData.Leafs["csApsSigFaultBER"] = types.YLeaf{"Csapssigfaultber", csapsconfigentry.Csapssigfaultber}
    csapsconfigentry.EntityData.Leafs["csApsSigDegradeBER"] = types.YLeaf{"Csapssigdegradeber", csapsconfigentry.Csapssigdegradeber}
    csapsconfigentry.EntityData.Leafs["csApsWaitToRestore"] = types.YLeaf{"Csapswaittorestore", csapsconfigentry.Csapswaittorestore}
    csapsconfigentry.EntityData.Leafs["csApsDirection"] = types.YLeaf{"Csapsdirection", csapsconfigentry.Csapsdirection}
    csapsconfigentry.EntityData.Leafs["csApsRevertive"] = types.YLeaf{"Csapsrevertive", csapsconfigentry.Csapsrevertive}
    csapsconfigentry.EntityData.Leafs["csApsDirectionOperational"] = types.YLeaf{"Csapsdirectionoperational", csapsconfigentry.Csapsdirectionoperational}
    csapsconfigentry.EntityData.Leafs["csApsArchModeOperational"] = types.YLeaf{"Csapsarchmodeoperational", csapsconfigentry.Csapsarchmodeoperational}
    csapsconfigentry.EntityData.Leafs["csApsChannelProtocol"] = types.YLeaf{"Csapschannelprotocol", csapsconfigentry.Csapschannelprotocol}
    csapsconfigentry.EntityData.Leafs["csApsFailureStatus"] = types.YLeaf{"Csapsfailurestatus", csapsconfigentry.Csapsfailurestatus}
    csapsconfigentry.EntityData.Leafs["csApsSwitchReason"] = types.YLeaf{"Csapsswitchreason", csapsconfigentry.Csapsswitchreason}
    csapsconfigentry.EntityData.Leafs["csApsPrimarySection"] = types.YLeaf{"Csapsprimarysection", csapsconfigentry.Csapsprimarysection}
    return &(csapsconfigentry.EntityData)
}

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline represents This object reflects the status of receive direction.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline_csApsWorkingLine CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline = "csApsWorkingLine"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline_csApsProtectionLine CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline = "csApsProtectionLine"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline_csApsNone CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsactiveline = "csApsNone"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode represents                             bytes ignored.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode_onePlusOne CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode = "onePlusOne"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode_oneToOne CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode = "oneToOne"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode_anexBOnePlusOne CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode = "anexBOnePlusOne"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode_ycableOnePlusOneNok1k2 CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode = "ycableOnePlusOneNok1k2"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode_straightOnePlusOneNok1k2 CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmode = "straightOnePlusOneNok1k2"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational represents Architecture mode configured.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational_onePlusOne CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational = "onePlusOne"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational_oneToOne CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational = "oneToOne"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational_anexBOnePlusOne CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational = "anexBOnePlusOne"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational_ycableOnePlusOneNok1k2 CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational = "ycableOnePlusOneNok1k2"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational_straightOnePlusOneNok1k2 CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsarchmodeoperational = "straightOnePlusOneNok1k2"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapschannelprotocol represents          ITU document G.783.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapschannelprotocol string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapschannelprotocol_bellcore CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapschannelprotocol = "bellcore"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapschannelprotocol_itu CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapschannelprotocol = "itu"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirection represents Bidirectional  : APS switch in both ends of the line.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirection string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirection_uniDirectional CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirection = "uniDirectional"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirection_biDirectional CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirection = "biDirectional"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirectionoperational represents both ends is operational at the Near End.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirectionoperational string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirectionoperational_uniDirectional CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirectionoperational = "uniDirectional"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirectionoperational_biDirectional CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsdirectionoperational = "biDirectional"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsenable represents or vice versa.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsenable string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsenable_csApsDisabled CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsenable = "csApsDisabled"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsenable_csApsEnabled CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsenable = "csApsEnabled"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection represents none(3)           : none.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection_workingSection1 CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection = "workingSection1"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection_workingSection2 CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection = "workingSection2"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection_none CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsprimarysection = "none"
)

// CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsrevertive represents   The active line does not switch to the working line.
type CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsrevertive string

const (
    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsrevertive_nonrevertive CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsrevertive = "nonrevertive"

    CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsrevertive_revertive CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry_Csapsrevertive = "revertive"
)

// CISCOSONETMIB_Csstotaltable
// The SONET/SDH Section Total table. It contains the 
// cumulative sum of the various statistics for the 24 hour
// period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the
// number of 15 minute intervals that have elapsed since
// the line is enabled. 
type CISCOSONETMIB_Csstotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Csstotaltable_Csstotalentry.
    Csstotalentry []CISCOSONETMIB_Csstotaltable_Csstotalentry
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetEntityData() *types.CommonEntityData {
    csstotaltable.EntityData.YFilter = csstotaltable.YFilter
    csstotaltable.EntityData.YangName = "cssTotalTable"
    csstotaltable.EntityData.BundleName = "cisco_ios_xe"
    csstotaltable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csstotaltable.EntityData.SegmentPath = "cssTotalTable"
    csstotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csstotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csstotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csstotaltable.EntityData.Children = make(map[string]types.YChild)
    csstotaltable.EntityData.Children["cssTotalEntry"] = types.YChild{"Csstotalentry", nil}
    for i := range csstotaltable.Csstotalentry {
        csstotaltable.EntityData.Children[types.GetSegmentPath(&csstotaltable.Csstotalentry[i])] = types.YChild{"Csstotalentry", &csstotaltable.Csstotalentry[i]}
    }
    csstotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csstotaltable.EntityData)
}

// CISCOSONETMIB_Csstotaltable_Csstotalentry
// An entry in the SONET/SDH Section Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Csstotaltable_Csstotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Section in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // errored seconds.
    Csstotaless interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Section
    // in the last 24 hours. The type is interface{} with range: 0..4294967295.
    // Units are severely errored seconds.
    Csstotalsess interface{}

    // The number of Severely Errored Framing Seconds  encountered by a SONET/SDH
    // Section in the last 24 hours. The type is interface{} with range:
    // 0..4294967295. Units are severely errored framing seconds.
    Csstotalsefss interface{}

    // The number of Coding Violations encountered by a  SONET/SDH Section in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // coding violations.
    Csstotalcvs interface{}
}

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetEntityData() *types.CommonEntityData {
    csstotalentry.EntityData.YFilter = csstotalentry.YFilter
    csstotalentry.EntityData.YangName = "cssTotalEntry"
    csstotalentry.EntityData.BundleName = "cisco_ios_xe"
    csstotalentry.EntityData.ParentYangName = "cssTotalTable"
    csstotalentry.EntityData.SegmentPath = "cssTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", csstotalentry.Ifindex) + "']"
    csstotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csstotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csstotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csstotalentry.EntityData.Children = make(map[string]types.YChild)
    csstotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csstotalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csstotalentry.Ifindex}
    csstotalentry.EntityData.Leafs["cssTotalESs"] = types.YLeaf{"Csstotaless", csstotalentry.Csstotaless}
    csstotalentry.EntityData.Leafs["cssTotalSESs"] = types.YLeaf{"Csstotalsess", csstotalentry.Csstotalsess}
    csstotalentry.EntityData.Leafs["cssTotalSEFSs"] = types.YLeaf{"Csstotalsefss", csstotalentry.Csstotalsefss}
    csstotalentry.EntityData.Leafs["cssTotalCVs"] = types.YLeaf{"Csstotalcvs", csstotalentry.Csstotalcvs}
    return &(csstotalentry.EntityData)
}

// CISCOSONETMIB_Csstracetable
// The SONET/SDH Section Trace table. This table contains 
// objects for tracing the sonet section.
type CISCOSONETMIB_Csstracetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the trace table. Entries exist for active sonet lines. The
    // objects in this table are used to verify  continued connection between the
    // two ends of the line. The type is slice of
    // CISCOSONETMIB_Csstracetable_Csstraceentry.
    Csstraceentry []CISCOSONETMIB_Csstracetable_Csstraceentry
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetEntityData() *types.CommonEntityData {
    csstracetable.EntityData.YFilter = csstracetable.YFilter
    csstracetable.EntityData.YangName = "cssTraceTable"
    csstracetable.EntityData.BundleName = "cisco_ios_xe"
    csstracetable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csstracetable.EntityData.SegmentPath = "cssTraceTable"
    csstracetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csstracetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csstracetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csstracetable.EntityData.Children = make(map[string]types.YChild)
    csstracetable.EntityData.Children["cssTraceEntry"] = types.YChild{"Csstraceentry", nil}
    for i := range csstracetable.Csstraceentry {
        csstracetable.EntityData.Children[types.GetSegmentPath(&csstracetable.Csstraceentry[i])] = types.YChild{"Csstraceentry", &csstracetable.Csstraceentry[i]}
    }
    csstracetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csstracetable.EntityData)
}

// CISCOSONETMIB_Csstracetable_Csstraceentry
// An entry in the trace table. Entries exist for active sonet
// lines. The objects in this table are used to verify 
// continued connection between the two ends of the line.
type CISCOSONETMIB_Csstracetable_Csstraceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // Sonet Section Trace To Transmit. This is string that is transmitted to
    // perform Sonet section trace  diagnostics. The trace string is  repetitively
    // transmited so that a trace receiving terminal can  verify its continued
    // connection to the intended  transmitter. The default value is a zero-length
    // string. Unless this object is set to a non-zero length string,  tracing
    // will not be performed. The type is string with length: 0 | 16 | 64.
    Csstracetotransmit interface{}

    // Sonet Section Trace To Expect. The receiving terminal  verifies if the
    // incoming string matches this string.  The value of 'cssTraceFailure'
    // indicates whether a  trace mismatch occurred. The default value is a 
    // zero-length string. The type is string with length: 0 | 16 | 64.
    Csstracetoexpect interface{}

    // The value of this object is set to 'true' when Sonet  Section received
    // trace does not match the  'cssTraceToExpect'. The type is bool.
    Csstracefailure interface{}

    // This object is used to view the Sonet Section Trace that  is received by
    // the receiving terminal. The type is string with length: 0 | 16 | 64.
    Csstracereceived interface{}
}

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetEntityData() *types.CommonEntityData {
    csstraceentry.EntityData.YFilter = csstraceentry.YFilter
    csstraceentry.EntityData.YangName = "cssTraceEntry"
    csstraceentry.EntityData.BundleName = "cisco_ios_xe"
    csstraceentry.EntityData.ParentYangName = "cssTraceTable"
    csstraceentry.EntityData.SegmentPath = "cssTraceEntry" + "[ifIndex='" + fmt.Sprintf("%v", csstraceentry.Ifindex) + "']"
    csstraceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csstraceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csstraceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csstraceentry.EntityData.Children = make(map[string]types.YChild)
    csstraceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csstraceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csstraceentry.Ifindex}
    csstraceentry.EntityData.Leafs["cssTraceToTransmit"] = types.YLeaf{"Csstracetotransmit", csstraceentry.Csstracetotransmit}
    csstraceentry.EntityData.Leafs["cssTraceToExpect"] = types.YLeaf{"Csstracetoexpect", csstraceentry.Csstracetoexpect}
    csstraceentry.EntityData.Leafs["cssTraceFailure"] = types.YLeaf{"Csstracefailure", csstraceentry.Csstracefailure}
    csstraceentry.EntityData.Leafs["cssTraceReceived"] = types.YLeaf{"Csstracereceived", csstraceentry.Csstracereceived}
    return &(csstraceentry.EntityData)
}

// CISCOSONETMIB_Csltotaltable
// The SONET/SDH Line Total table. It contains the 
// cumulative sum of the various statistics for the 24 
// hour period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the 
// number of 15 minute intervals that have elapsed since
// the line is enabled.
type CISCOSONETMIB_Csltotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Csltotaltable_Csltotalentry.
    Csltotalentry []CISCOSONETMIB_Csltotaltable_Csltotalentry
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetEntityData() *types.CommonEntityData {
    csltotaltable.EntityData.YFilter = csltotaltable.YFilter
    csltotaltable.EntityData.YangName = "cslTotalTable"
    csltotaltable.EntityData.BundleName = "cisco_ios_xe"
    csltotaltable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csltotaltable.EntityData.SegmentPath = "cslTotalTable"
    csltotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csltotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csltotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csltotaltable.EntityData.Children = make(map[string]types.YChild)
    csltotaltable.EntityData.Children["cslTotalEntry"] = types.YChild{"Csltotalentry", nil}
    for i := range csltotaltable.Csltotalentry {
        csltotaltable.EntityData.Children[types.GetSegmentPath(&csltotaltable.Csltotalentry[i])] = types.YChild{"Csltotalentry", &csltotaltable.Csltotalentry[i]}
    }
    csltotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csltotaltable.EntityData)
}

// CISCOSONETMIB_Csltotaltable_Csltotalentry
// An entry in the SONET/SDH Line Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Csltotaltable_Csltotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Line in the last
    // 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // errored seconds.
    Csltotaless interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Line in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are severely errored seconds.
    Csltotalsess interface{}

    // The number of Coding Violations encountered by a  SONET/SDH Line in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // coding violations.
    Csltotalcvs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH Line in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // unavailable seconds.
    Csltotaluass interface{}
}

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetEntityData() *types.CommonEntityData {
    csltotalentry.EntityData.YFilter = csltotalentry.YFilter
    csltotalentry.EntityData.YangName = "cslTotalEntry"
    csltotalentry.EntityData.BundleName = "cisco_ios_xe"
    csltotalentry.EntityData.ParentYangName = "cslTotalTable"
    csltotalentry.EntityData.SegmentPath = "cslTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", csltotalentry.Ifindex) + "']"
    csltotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csltotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csltotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csltotalentry.EntityData.Children = make(map[string]types.YChild)
    csltotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csltotalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csltotalentry.Ifindex}
    csltotalentry.EntityData.Leafs["cslTotalESs"] = types.YLeaf{"Csltotaless", csltotalentry.Csltotaless}
    csltotalentry.EntityData.Leafs["cslTotalSESs"] = types.YLeaf{"Csltotalsess", csltotalentry.Csltotalsess}
    csltotalentry.EntityData.Leafs["cslTotalCVs"] = types.YLeaf{"Csltotalcvs", csltotalentry.Csltotalcvs}
    csltotalentry.EntityData.Leafs["cslTotalUASs"] = types.YLeaf{"Csltotaluass", csltotalentry.Csltotaluass}
    return &(csltotalentry.EntityData)
}

// CISCOSONETMIB_Cslfarendtotaltable
// The SONET/SDH Far End Line Total table. It contains the 
// cumulative sum of the various statistics for the 24 hour 
// period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the 
// number of 15 minute intervals that have elapsed since the 
// line is enabled.
type CISCOSONETMIB_Cslfarendtotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry.
    Cslfarendtotalentry []CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetEntityData() *types.CommonEntityData {
    cslfarendtotaltable.EntityData.YFilter = cslfarendtotaltable.YFilter
    cslfarendtotaltable.EntityData.YangName = "cslFarEndTotalTable"
    cslfarendtotaltable.EntityData.BundleName = "cisco_ios_xe"
    cslfarendtotaltable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cslfarendtotaltable.EntityData.SegmentPath = "cslFarEndTotalTable"
    cslfarendtotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cslfarendtotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cslfarendtotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cslfarendtotaltable.EntityData.Children = make(map[string]types.YChild)
    cslfarendtotaltable.EntityData.Children["cslFarEndTotalEntry"] = types.YChild{"Cslfarendtotalentry", nil}
    for i := range cslfarendtotaltable.Cslfarendtotalentry {
        cslfarendtotaltable.EntityData.Children[types.GetSegmentPath(&cslfarendtotaltable.Cslfarendtotalentry[i])] = types.YChild{"Cslfarendtotalentry", &cslfarendtotaltable.Cslfarendtotalentry[i]}
    }
    cslfarendtotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cslfarendtotaltable.EntityData)
}

// CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry
// An entry in the SONET/SDH Far End Line Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Far End Line in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are errored seconds.
    Cslfarendtotaless interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Far End
    // Line in the last 24 hours. The type is interface{} with range:
    // 0..4294967295. Units are severely errored seconds.
    Cslfarendtotalsess interface{}

    // The number of Coding Violations encountered by a SONET/SDH  Far End Line in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are coding violations.
    Cslfarendtotalcvs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH Far End Line
    // in the last 24 hours. The type is interface{} with range: 0..4294967295.
    // Units are unavailable seconds.
    Cslfarendtotaluass interface{}
}

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetEntityData() *types.CommonEntityData {
    cslfarendtotalentry.EntityData.YFilter = cslfarendtotalentry.YFilter
    cslfarendtotalentry.EntityData.YangName = "cslFarEndTotalEntry"
    cslfarendtotalentry.EntityData.BundleName = "cisco_ios_xe"
    cslfarendtotalentry.EntityData.ParentYangName = "cslFarEndTotalTable"
    cslfarendtotalentry.EntityData.SegmentPath = "cslFarEndTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", cslfarendtotalentry.Ifindex) + "']"
    cslfarendtotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cslfarendtotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cslfarendtotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cslfarendtotalentry.EntityData.Children = make(map[string]types.YChild)
    cslfarendtotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cslfarendtotalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cslfarendtotalentry.Ifindex}
    cslfarendtotalentry.EntityData.Leafs["cslFarEndTotalESs"] = types.YLeaf{"Cslfarendtotaless", cslfarendtotalentry.Cslfarendtotaless}
    cslfarendtotalentry.EntityData.Leafs["cslFarEndTotalSESs"] = types.YLeaf{"Cslfarendtotalsess", cslfarendtotalentry.Cslfarendtotalsess}
    cslfarendtotalentry.EntityData.Leafs["cslFarEndTotalCVs"] = types.YLeaf{"Cslfarendtotalcvs", cslfarendtotalentry.Cslfarendtotalcvs}
    cslfarendtotalentry.EntityData.Leafs["cslFarEndTotalUASs"] = types.YLeaf{"Cslfarendtotaluass", cslfarendtotalentry.Cslfarendtotaluass}
    return &(cslfarendtotalentry.EntityData)
}

// CISCOSONETMIB_Csptotaltable
// The SONET/SDH Path Total table. It contains the cumulative 
// sum of the various statistics for the 24 hour period 
// preceding the current interval.The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the number
// of 15 minute intervals that have elapsed since the line is 
// enabled.
type CISCOSONETMIB_Csptotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Csptotaltable_Csptotalentry.
    Csptotalentry []CISCOSONETMIB_Csptotaltable_Csptotalentry
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetEntityData() *types.CommonEntityData {
    csptotaltable.EntityData.YFilter = csptotaltable.YFilter
    csptotaltable.EntityData.YangName = "cspTotalTable"
    csptotaltable.EntityData.BundleName = "cisco_ios_xe"
    csptotaltable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csptotaltable.EntityData.SegmentPath = "cspTotalTable"
    csptotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csptotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csptotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csptotaltable.EntityData.Children = make(map[string]types.YChild)
    csptotaltable.EntityData.Children["cspTotalEntry"] = types.YChild{"Csptotalentry", nil}
    for i := range csptotaltable.Csptotalentry {
        csptotaltable.EntityData.Children[types.GetSegmentPath(&csptotaltable.Csptotalentry[i])] = types.YChild{"Csptotalentry", &csptotaltable.Csptotalentry[i]}
    }
    csptotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csptotaltable.EntityData)
}

// CISCOSONETMIB_Csptotaltable_Csptotalentry
// An entry in the SONET/SDH Path Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Csptotaltable_Csptotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH Path in the last
    // 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // errored seconds.
    Csptotaless interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH Path in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are severely errored seconds.
    Csptotalsess interface{}

    // The number of Coding Violations encountered by a  SONET/SDH Path in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // coding violations.
    Csptotalcvs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH Path in the
    // last 24 hours. The type is interface{} with range: 0..4294967295. Units are
    // unavailable seconds.
    Csptotaluass interface{}
}

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetEntityData() *types.CommonEntityData {
    csptotalentry.EntityData.YFilter = csptotalentry.YFilter
    csptotalentry.EntityData.YangName = "cspTotalEntry"
    csptotalentry.EntityData.BundleName = "cisco_ios_xe"
    csptotalentry.EntityData.ParentYangName = "cspTotalTable"
    csptotalentry.EntityData.SegmentPath = "cspTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", csptotalentry.Ifindex) + "']"
    csptotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csptotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csptotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csptotalentry.EntityData.Children = make(map[string]types.YChild)
    csptotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csptotalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csptotalentry.Ifindex}
    csptotalentry.EntityData.Leafs["cspTotalESs"] = types.YLeaf{"Csptotaless", csptotalentry.Csptotaless}
    csptotalentry.EntityData.Leafs["cspTotalSESs"] = types.YLeaf{"Csptotalsess", csptotalentry.Csptotalsess}
    csptotalentry.EntityData.Leafs["cspTotalCVs"] = types.YLeaf{"Csptotalcvs", csptotalentry.Csptotalcvs}
    csptotalentry.EntityData.Leafs["cspTotalUASs"] = types.YLeaf{"Csptotaluass", csptotalentry.Csptotaluass}
    return &(csptotalentry.EntityData)
}

// CISCOSONETMIB_Cspfarendtotaltable
// The SONET/SDH Far End Path Total table. Far End is the 
// remote end of the line. The table contains the cumulative
// sum of the various statistics for the 24 hour period 
// preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the
// number of 15 minute intervals that have elapsed since
// the line is enabled. 
type CISCOSONETMIB_Cspfarendtotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Total table.  Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry.
    Cspfarendtotalentry []CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetEntityData() *types.CommonEntityData {
    cspfarendtotaltable.EntityData.YFilter = cspfarendtotaltable.YFilter
    cspfarendtotaltable.EntityData.YangName = "cspFarEndTotalTable"
    cspfarendtotaltable.EntityData.BundleName = "cisco_ios_xe"
    cspfarendtotaltable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    cspfarendtotaltable.EntityData.SegmentPath = "cspFarEndTotalTable"
    cspfarendtotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspfarendtotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspfarendtotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspfarendtotaltable.EntityData.Children = make(map[string]types.YChild)
    cspfarendtotaltable.EntityData.Children["cspFarEndTotalEntry"] = types.YChild{"Cspfarendtotalentry", nil}
    for i := range cspfarendtotaltable.Cspfarendtotalentry {
        cspfarendtotaltable.EntityData.Children[types.GetSegmentPath(&cspfarendtotaltable.Cspfarendtotalentry[i])] = types.YChild{"Cspfarendtotalentry", &cspfarendtotaltable.Cspfarendtotalentry[i]}
    }
    cspfarendtotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cspfarendtotaltable.EntityData)
}

// CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry
// An entry in the SONET/SDH Far End Path Total table. 
// Entries are created automatically for sonet lines.
type CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of Errored Seconds encountered by a SONET/SDH far end path in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are errored seconds.
    Cspfarendtotaless interface{}

    // The number of Severely Errored Seconds encountered by a  SONET/SDH far end
    // path in the last 24 hours. The type is interface{} with range:
    // 0..4294967295. Units are severely errored seconds.
    Cspfarendtotalsess interface{}

    // The number of Coding Violations encountered by a  SONET/SDH far end path in
    // the last 24 hours. The type is interface{} with range: 0..4294967295. Units
    // are coding violations.
    Cspfarendtotalcvs interface{}

    // The number of Unavailable Seconds encountered by a  SONET/SDH far end path
    // in the last 24 hours. The type is interface{} with range: 0..4294967295.
    // Units are unavailable seconds.
    Cspfarendtotaluass interface{}
}

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetEntityData() *types.CommonEntityData {
    cspfarendtotalentry.EntityData.YFilter = cspfarendtotalentry.YFilter
    cspfarendtotalentry.EntityData.YangName = "cspFarEndTotalEntry"
    cspfarendtotalentry.EntityData.BundleName = "cisco_ios_xe"
    cspfarendtotalentry.EntityData.ParentYangName = "cspFarEndTotalTable"
    cspfarendtotalentry.EntityData.SegmentPath = "cspFarEndTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", cspfarendtotalentry.Ifindex) + "']"
    cspfarendtotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cspfarendtotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cspfarendtotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cspfarendtotalentry.EntityData.Children = make(map[string]types.YChild)
    cspfarendtotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cspfarendtotalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cspfarendtotalentry.Ifindex}
    cspfarendtotalentry.EntityData.Leafs["cspFarEndTotalESs"] = types.YLeaf{"Cspfarendtotaless", cspfarendtotalentry.Cspfarendtotaless}
    cspfarendtotalentry.EntityData.Leafs["cspFarEndTotalSESs"] = types.YLeaf{"Cspfarendtotalsess", cspfarendtotalentry.Cspfarendtotalsess}
    cspfarendtotalentry.EntityData.Leafs["cspFarEndTotalCVs"] = types.YLeaf{"Cspfarendtotalcvs", cspfarendtotalentry.Cspfarendtotalcvs}
    cspfarendtotalentry.EntityData.Leafs["cspFarEndTotalUASs"] = types.YLeaf{"Cspfarendtotaluass", cspfarendtotalentry.Cspfarendtotaluass}
    return &(cspfarendtotalentry.EntityData)
}

// CISCOSONETMIB_Csptracetable
// The SONET/SDH Path Trace table. This table contains objects 
// for tracing the sonet path.
type CISCOSONETMIB_Csptracetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Trace table. The entries  exist for active
    // sonet lines. The objects in this table are  used to verify continued
    // connection between the two ends of the line. The type is slice of
    // CISCOSONETMIB_Csptracetable_Csptraceentry.
    Csptraceentry []CISCOSONETMIB_Csptracetable_Csptraceentry
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetEntityData() *types.CommonEntityData {
    csptracetable.EntityData.YFilter = csptracetable.YFilter
    csptracetable.EntityData.YangName = "cspTraceTable"
    csptracetable.EntityData.BundleName = "cisco_ios_xe"
    csptracetable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csptracetable.EntityData.SegmentPath = "cspTraceTable"
    csptracetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csptracetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csptracetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csptracetable.EntityData.Children = make(map[string]types.YChild)
    csptracetable.EntityData.Children["cspTraceEntry"] = types.YChild{"Csptraceentry", nil}
    for i := range csptracetable.Csptraceentry {
        csptracetable.EntityData.Children[types.GetSegmentPath(&csptracetable.Csptraceentry[i])] = types.YChild{"Csptraceentry", &csptracetable.Csptraceentry[i]}
    }
    csptracetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csptracetable.EntityData)
}

// CISCOSONETMIB_Csptracetable_Csptraceentry
// An entry in the SONET/SDH Path Trace table. The entries 
// exist for active sonet lines. The objects in this table are 
// used to verify continued connection between the two ends of
// the line.
type CISCOSONETMIB_Csptracetable_Csptraceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // Sonet Path Trace To Transmit. The trace string is  repetitively transmited
    // so that a trace receiving terminal  can verify its continued receiving
    // terminal can verify its  continued connection to the intended transmitter.
    // The  default value is a zero-length string. Unless this object  is set to a
    // non-zero length string, tracing will not be  performed. The type is string
    // with length: 0 | 16 | 64.
    Csptracetotransmit interface{}

    // Sonet Path Trace To Expect.  The receiving terminal verifies if the
    // incoming string matches this string. The value of  'cspTraceFailure'
    // indicates whether a trace mismatch  occured. The default value is a
    // zero-length string. The type is string with length: 0 | 16 | 64.
    Csptracetoexpect interface{}

    // The value of this object is set to 'true' when Sonet Path  received trace
    // does not match the 'cspTraceToExpect'. The type is bool.
    Csptracefailure interface{}

    // This object is used to view the Sonet Path Trace that is received by the
    // receiving terminal. The type is string with length: 0 | 16 | 64.
    Csptracereceived interface{}
}

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetEntityData() *types.CommonEntityData {
    csptraceentry.EntityData.YFilter = csptraceentry.YFilter
    csptraceentry.EntityData.YangName = "cspTraceEntry"
    csptraceentry.EntityData.BundleName = "cisco_ios_xe"
    csptraceentry.EntityData.ParentYangName = "cspTraceTable"
    csptraceentry.EntityData.SegmentPath = "cspTraceEntry" + "[ifIndex='" + fmt.Sprintf("%v", csptraceentry.Ifindex) + "']"
    csptraceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csptraceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csptraceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csptraceentry.EntityData.Children = make(map[string]types.YChild)
    csptraceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csptraceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csptraceentry.Ifindex}
    csptraceentry.EntityData.Leafs["cspTraceToTransmit"] = types.YLeaf{"Csptracetotransmit", csptraceentry.Csptracetotransmit}
    csptraceentry.EntityData.Leafs["cspTraceToExpect"] = types.YLeaf{"Csptracetoexpect", csptraceentry.Csptracetoexpect}
    csptraceentry.EntityData.Leafs["cspTraceFailure"] = types.YLeaf{"Csptracefailure", csptraceentry.Csptracefailure}
    csptraceentry.EntityData.Leafs["cspTraceReceived"] = types.YLeaf{"Csptracereceived", csptraceentry.Csptracereceived}
    return &(csptraceentry.EntityData)
}

// CISCOSONETMIB_Csstatstable
// The SONET/SDH Section statistics table. This table 
// maintains the number of times the line encountered Loss of
// Signal(LOS), Loss of frame(LOF), Alarm Indication 
// signals(AISs), Remote failure indications(RFIs).
type CISCOSONETMIB_Csstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH statistics table. These are  realtime statistics
    // for the Sonet section, line and path layers. The statistics are gathered
    // for each sonet line.  An entry is created automatically and is indexed by 
    // ifIndex. The type is slice of CISCOSONETMIB_Csstatstable_Csstatsentry.
    Csstatsentry []CISCOSONETMIB_Csstatstable_Csstatsentry
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetEntityData() *types.CommonEntityData {
    csstatstable.EntityData.YFilter = csstatstable.YFilter
    csstatstable.EntityData.YangName = "csStatsTable"
    csstatstable.EntityData.BundleName = "cisco_ios_xe"
    csstatstable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csstatstable.EntityData.SegmentPath = "csStatsTable"
    csstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csstatstable.EntityData.Children = make(map[string]types.YChild)
    csstatstable.EntityData.Children["csStatsEntry"] = types.YChild{"Csstatsentry", nil}
    for i := range csstatstable.Csstatsentry {
        csstatstable.EntityData.Children[types.GetSegmentPath(&csstatstable.Csstatsentry[i])] = types.YChild{"Csstatsentry", &csstatstable.Csstatsentry[i]}
    }
    csstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csstatstable.EntityData)
}

// CISCOSONETMIB_Csstatstable_Csstatsentry
// An entry in the SONET/SDH statistics table. These are 
// realtime statistics for the Sonet section, line and path
// layers. The statistics are gathered for each sonet line. 
// An entry is created automatically and is indexed by 
// ifIndex.
type CISCOSONETMIB_Csstatstable_Csstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of Loss of signals(LOS) encountered by a  SONET/SDH Section. A
    // high value for this object may  indicate a problem with the Sonet Section
    // layer. The type is interface{} with range: 0..4294967295. Units are loss of
    // signals.
    Cssloss interface{}

    // The number of Loss of Frames (LOF) encountered by a  SONET/SDH Section. A
    // high value for this object may  indicate a problem with the Sonet Section
    // layer. The type is interface{} with range: 0..4294967295. Units are loss of
    // frames.
    Csslofs interface{}

    // The number of alarm indication signals(AIS)  encountered by  a SONET/SDH
    // Line. A high value for this object may indicate a problem with the Sonet
    // Line layer. The type is interface{} with range: 0..4294967295. Units are
    // alarm indication signals.
    Cslaiss interface{}

    // The number of remote failure indications (RFI) encountered  by a SONET/SDH
    // Line. A high value for this object may  indicate a problem with the Sonet
    // Line layer. The type is interface{} with range: 0..4294967295. Units are
    // remote failure indications.
    Cslrfis interface{}

    // The  number of alarm indication signals (AIS) encountered by a SONET/SDH
    // Path. A high value for this object may  indicate a problem with the Sonet
    // Path layer. The type is interface{} with range: 0..4294967295. Units are
    // alarm indication signals.
    Cspaiss interface{}

    // The number of  remote failure indications (RFI)  encountered by a SONET/SDH
    // Path. A high value for this  object may indicate a problem with the Sonet
    // Path layer. The type is interface{} with range: 0..4294967295. Units are
    // remote failure indications.
    Csprfis interface{}
}

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetEntityData() *types.CommonEntityData {
    csstatsentry.EntityData.YFilter = csstatsentry.YFilter
    csstatsentry.EntityData.YangName = "csStatsEntry"
    csstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csstatsentry.EntityData.ParentYangName = "csStatsTable"
    csstatsentry.EntityData.SegmentPath = "csStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", csstatsentry.Ifindex) + "']"
    csstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csstatsentry.EntityData.Children = make(map[string]types.YChild)
    csstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csstatsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csstatsentry.Ifindex}
    csstatsentry.EntityData.Leafs["cssLOSs"] = types.YLeaf{"Cssloss", csstatsentry.Cssloss}
    csstatsentry.EntityData.Leafs["cssLOFs"] = types.YLeaf{"Csslofs", csstatsentry.Csslofs}
    csstatsentry.EntityData.Leafs["cslAISs"] = types.YLeaf{"Cslaiss", csstatsentry.Cslaiss}
    csstatsentry.EntityData.Leafs["cslRFIs"] = types.YLeaf{"Cslrfis", csstatsentry.Cslrfis}
    csstatsentry.EntityData.Leafs["cspAISs"] = types.YLeaf{"Cspaiss", csstatsentry.Cspaiss}
    csstatsentry.EntityData.Leafs["cspRFIs"] = types.YLeaf{"Csprfis", csstatsentry.Csprfis}
    return &(csstatsentry.EntityData)
}

// CISCOSONETMIB_Csau4Tug3Configtable
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
type CISCOSONETMIB_Csau4Tug3Configtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // There is an entry in this table for each TUG-3 within a  AU-4 SDH path that
    // supports SDH virtual container VC-4. The ifIndex value represents an entry
    // in ifTable with ifType = sonetPath(50).The ifTable entry applicable for
    // this entry belongs to AU-4 path. The type is slice of
    // CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry.
    Csau4Tug3Configentry []CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetEntityData() *types.CommonEntityData {
    csau4Tug3Configtable.EntityData.YFilter = csau4Tug3Configtable.YFilter
    csau4Tug3Configtable.EntityData.YangName = "csAu4Tug3ConfigTable"
    csau4Tug3Configtable.EntityData.BundleName = "cisco_ios_xe"
    csau4Tug3Configtable.EntityData.ParentYangName = "CISCO-SONET-MIB"
    csau4Tug3Configtable.EntityData.SegmentPath = "csAu4Tug3ConfigTable"
    csau4Tug3Configtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csau4Tug3Configtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csau4Tug3Configtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csau4Tug3Configtable.EntityData.Children = make(map[string]types.YChild)
    csau4Tug3Configtable.EntityData.Children["csAu4Tug3ConfigEntry"] = types.YChild{"Csau4Tug3Configentry", nil}
    for i := range csau4Tug3Configtable.Csau4Tug3Configentry {
        csau4Tug3Configtable.EntityData.Children[types.GetSegmentPath(&csau4Tug3Configtable.Csau4Tug3Configentry[i])] = types.YChild{"Csau4Tug3Configentry", &csau4Tug3Configtable.Csau4Tug3Configentry[i]}
    }
    csau4Tug3Configtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csau4Tug3Configtable.EntityData)
}

// CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry
// There is an entry in this table for each TUG-3 within a 
// AU-4 SDH path that supports SDH virtual container VC-4.
// The ifIndex value represents an entry in ifTable with
// ifType = sonetPath(50).The ifTable entry applicable for
// this entry belongs to AU-4 path.
type CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. This object represents the TUG-3 number. The type
    // is interface{} with range: 1..3.
    Csau4Tug3 interface{}

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
    // Csau4Tug3Payload.
    Csau4Tug3Payload interface{}
}

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetEntityData() *types.CommonEntityData {
    csau4Tug3Configentry.EntityData.YFilter = csau4Tug3Configentry.YFilter
    csau4Tug3Configentry.EntityData.YangName = "csAu4Tug3ConfigEntry"
    csau4Tug3Configentry.EntityData.BundleName = "cisco_ios_xe"
    csau4Tug3Configentry.EntityData.ParentYangName = "csAu4Tug3ConfigTable"
    csau4Tug3Configentry.EntityData.SegmentPath = "csAu4Tug3ConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", csau4Tug3Configentry.Ifindex) + "']" + "[csAu4Tug3='" + fmt.Sprintf("%v", csau4Tug3Configentry.Csau4Tug3) + "']"
    csau4Tug3Configentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csau4Tug3Configentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csau4Tug3Configentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csau4Tug3Configentry.EntityData.Children = make(map[string]types.YChild)
    csau4Tug3Configentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csau4Tug3Configentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csau4Tug3Configentry.Ifindex}
    csau4Tug3Configentry.EntityData.Leafs["csAu4Tug3"] = types.YLeaf{"Csau4Tug3", csau4Tug3Configentry.Csau4Tug3}
    csau4Tug3Configentry.EntityData.Leafs["csAu4Tug3Payload"] = types.YLeaf{"Csau4Tug3Payload", csau4Tug3Configentry.Csau4Tug3Payload}
    return &(csau4Tug3Configentry.EntityData)
}

// CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload represents The value 'other' can not be set.
type CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload string

const (
    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_other CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "other"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_vc11 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "vc11"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_vc12 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "vc12"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_tu3ds3 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "tu3ds3"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_tu3e3 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "tu3e3"
)

