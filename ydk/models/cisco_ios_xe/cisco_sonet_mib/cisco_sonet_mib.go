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
    parent types.Entity
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

func (cISCOSONETMIB *CISCOSONETMIB) GetFilter() yfilter.YFilter { return cISCOSONETMIB.YFilter }

func (cISCOSONETMIB *CISCOSONETMIB) SetFilter(yf yfilter.YFilter) { cISCOSONETMIB.YFilter = yf }

func (cISCOSONETMIB *CISCOSONETMIB) GetGoName(yname string) string {
    if yname == "csApsConfig" { return "Csapsconfig" }
    if yname == "csNotifications" { return "Csnotifications" }
    if yname == "csConfigTable" { return "Csconfigtable" }
    if yname == "csApsConfigTable" { return "Csapsconfigtable" }
    if yname == "cssTotalTable" { return "Csstotaltable" }
    if yname == "cssTraceTable" { return "Csstracetable" }
    if yname == "cslTotalTable" { return "Csltotaltable" }
    if yname == "cslFarEndTotalTable" { return "Cslfarendtotaltable" }
    if yname == "cspTotalTable" { return "Csptotaltable" }
    if yname == "cspFarEndTotalTable" { return "Cspfarendtotaltable" }
    if yname == "cspTraceTable" { return "Csptracetable" }
    if yname == "csStatsTable" { return "Csstatstable" }
    if yname == "csAu4Tug3ConfigTable" { return "Csau4Tug3Configtable" }
    return ""
}

func (cISCOSONETMIB *CISCOSONETMIB) GetSegmentPath() string {
    return "CISCO-SONET-MIB:CISCO-SONET-MIB"
}

func (cISCOSONETMIB *CISCOSONETMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csApsConfig" {
        return &cISCOSONETMIB.Csapsconfig
    }
    if childYangName == "csNotifications" {
        return &cISCOSONETMIB.Csnotifications
    }
    if childYangName == "csConfigTable" {
        return &cISCOSONETMIB.Csconfigtable
    }
    if childYangName == "csApsConfigTable" {
        return &cISCOSONETMIB.Csapsconfigtable
    }
    if childYangName == "cssTotalTable" {
        return &cISCOSONETMIB.Csstotaltable
    }
    if childYangName == "cssTraceTable" {
        return &cISCOSONETMIB.Csstracetable
    }
    if childYangName == "cslTotalTable" {
        return &cISCOSONETMIB.Csltotaltable
    }
    if childYangName == "cslFarEndTotalTable" {
        return &cISCOSONETMIB.Cslfarendtotaltable
    }
    if childYangName == "cspTotalTable" {
        return &cISCOSONETMIB.Csptotaltable
    }
    if childYangName == "cspFarEndTotalTable" {
        return &cISCOSONETMIB.Cspfarendtotaltable
    }
    if childYangName == "cspTraceTable" {
        return &cISCOSONETMIB.Csptracetable
    }
    if childYangName == "csStatsTable" {
        return &cISCOSONETMIB.Csstatstable
    }
    if childYangName == "csAu4Tug3ConfigTable" {
        return &cISCOSONETMIB.Csau4Tug3Configtable
    }
    return nil
}

func (cISCOSONETMIB *CISCOSONETMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["csApsConfig"] = &cISCOSONETMIB.Csapsconfig
    children["csNotifications"] = &cISCOSONETMIB.Csnotifications
    children["csConfigTable"] = &cISCOSONETMIB.Csconfigtable
    children["csApsConfigTable"] = &cISCOSONETMIB.Csapsconfigtable
    children["cssTotalTable"] = &cISCOSONETMIB.Csstotaltable
    children["cssTraceTable"] = &cISCOSONETMIB.Csstracetable
    children["cslTotalTable"] = &cISCOSONETMIB.Csltotaltable
    children["cslFarEndTotalTable"] = &cISCOSONETMIB.Cslfarendtotaltable
    children["cspTotalTable"] = &cISCOSONETMIB.Csptotaltable
    children["cspFarEndTotalTable"] = &cISCOSONETMIB.Cspfarendtotaltable
    children["cspTraceTable"] = &cISCOSONETMIB.Csptracetable
    children["csStatsTable"] = &cISCOSONETMIB.Csstatstable
    children["csAu4Tug3ConfigTable"] = &cISCOSONETMIB.Csau4Tug3Configtable
    return children
}

func (cISCOSONETMIB *CISCOSONETMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOSONETMIB *CISCOSONETMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOSONETMIB *CISCOSONETMIB) GetYangName() string { return "CISCO-SONET-MIB" }

func (cISCOSONETMIB *CISCOSONETMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOSONETMIB *CISCOSONETMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOSONETMIB *CISCOSONETMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOSONETMIB *CISCOSONETMIB) SetParent(parent types.Entity) { cISCOSONETMIB.parent = parent }

func (cISCOSONETMIB *CISCOSONETMIB) GetParent() types.Entity { return cISCOSONETMIB.parent }

func (cISCOSONETMIB *CISCOSONETMIB) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csapsconfig
type CISCOSONETMIB_Csapsconfig struct {
    parent types.Entity
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

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetFilter() yfilter.YFilter { return csapsconfig.YFilter }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) SetFilter(yf yfilter.YFilter) { csapsconfig.YFilter = yf }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetGoName(yname string) string {
    if yname == "csApsLineFailureCode" { return "Csapslinefailurecode" }
    if yname == "csApsLineSwitchReason" { return "Csapslineswitchreason" }
    return ""
}

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetSegmentPath() string {
    return "csApsConfig"
}

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csApsLineFailureCode"] = csapsconfig.Csapslinefailurecode
    leafs["csApsLineSwitchReason"] = csapsconfig.Csapslineswitchreason
    return leafs
}

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetBundleName() string { return "cisco_ios_xe" }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetYangName() string { return "csApsConfig" }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) SetParent(parent types.Entity) { csapsconfig.parent = parent }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetParent() types.Entity { return csapsconfig.parent }

func (csapsconfig *CISCOSONETMIB_Csapsconfig) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csnotifications
type CISCOSONETMIB_Csnotifications struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object controls if the generation of  ciscoSonetSectionStatusChange,
    // ciscoSonetLineStatusChange,  ciscoSonetPathStatusChange and
    // ciscoSonetVTStatusChange notifications is enabled. If the value of this
    // object is 'true(1)', then all notifications in this MIB are enabled;
    // otherwise they are disabled. The type is bool.
    Csnotificationsenabled interface{}
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetFilter() yfilter.YFilter { return csnotifications.YFilter }

func (csnotifications *CISCOSONETMIB_Csnotifications) SetFilter(yf yfilter.YFilter) { csnotifications.YFilter = yf }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetGoName(yname string) string {
    if yname == "csNotificationsEnabled" { return "Csnotificationsenabled" }
    return ""
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetSegmentPath() string {
    return "csNotifications"
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csNotificationsEnabled"] = csnotifications.Csnotificationsenabled
    return leafs
}

func (csnotifications *CISCOSONETMIB_Csnotifications) GetBundleName() string { return "cisco_ios_xe" }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetYangName() string { return "csNotifications" }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csnotifications *CISCOSONETMIB_Csnotifications) SetParent(parent types.Entity) { csnotifications.parent = parent }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetParent() types.Entity { return csnotifications.parent }

func (csnotifications *CISCOSONETMIB_Csnotifications) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csconfigtable
// The SONET/SDH configuration table. This table has objects 
// for configuring sonet lines.
type CISCOSONETMIB_Csconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table. There is an entry for each SONET line  in the table.
    // Entries are automatically created for an  ifType value of sonet(39).
    // 'ifAdminStatus' from the ifTable  must be used to enable or disable a line.
    // A line is in  disabled(down) state unless provisioned 'up' using 
    // 'ifAdminStatus'. The type is slice of
    // CISCOSONETMIB_Csconfigtable_Csconfigentry.
    Csconfigentry []CISCOSONETMIB_Csconfigtable_Csconfigentry
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetFilter() yfilter.YFilter { return csconfigtable.YFilter }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) SetFilter(yf yfilter.YFilter) { csconfigtable.YFilter = yf }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetGoName(yname string) string {
    if yname == "csConfigEntry" { return "Csconfigentry" }
    return ""
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetSegmentPath() string {
    return "csConfigTable"
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csConfigEntry" {
        for _, c := range csconfigtable.Csconfigentry {
            if csconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csconfigtable_Csconfigentry{}
        csconfigtable.Csconfigentry = append(csconfigtable.Csconfigentry, child)
        return &csconfigtable.Csconfigentry[len(csconfigtable.Csconfigentry)-1]
    }
    return nil
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csconfigtable.Csconfigentry {
        children[csconfigtable.Csconfigentry[i].GetSegmentPath()] = &csconfigtable.Csconfigentry[i]
    }
    return children
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetYangName() string { return "csConfigTable" }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) SetParent(parent types.Entity) { csconfigtable.parent = parent }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetParent() types.Entity { return csconfigtable.parent }

func (csconfigtable *CISCOSONETMIB_Csconfigtable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csconfigtable_Csconfigentry
// An entry in the table. There is an entry for each SONET line 
// in the table. Entries are automatically created for an 
// ifType value of sonet(39). 'ifAdminStatus' from the ifTable 
// must be used to enable or disable a line. A line is in 
// disabled(down) state unless provisioned 'up' using 
// 'ifAdminStatus'.
type CISCOSONETMIB_Csconfigtable_Csconfigentry struct {
    parent types.Entity
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

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetFilter() yfilter.YFilter { return csconfigentry.YFilter }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) SetFilter(yf yfilter.YFilter) { csconfigentry.YFilter = yf }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "csConfigLoopbackType" { return "Csconfigloopbacktype" }
    if yname == "csConfigXmtClockSource" { return "Csconfigxmtclocksource" }
    if yname == "csConfigFrameScramble" { return "Csconfigframescramble" }
    if yname == "csConfigType" { return "Csconfigtype" }
    if yname == "csConfigRDIVType" { return "Csconfigrdivtype" }
    if yname == "csConfigRDIPType" { return "Csconfigrdiptype" }
    if yname == "csTributaryType" { return "Cstributarytype" }
    if yname == "csTributaryMappingType" { return "Cstributarymappingtype" }
    if yname == "csTributaryFramingType" { return "Cstributaryframingtype" }
    if yname == "csSignallingTransportMode" { return "Cssignallingtransportmode" }
    if yname == "csTributaryGroupingType" { return "Cstributarygroupingtype" }
    return ""
}

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetSegmentPath() string {
    return "csConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", csconfigentry.Ifindex) + "']"
}

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csconfigentry.Ifindex
    leafs["csConfigLoopbackType"] = csconfigentry.Csconfigloopbacktype
    leafs["csConfigXmtClockSource"] = csconfigentry.Csconfigxmtclocksource
    leafs["csConfigFrameScramble"] = csconfigentry.Csconfigframescramble
    leafs["csConfigType"] = csconfigentry.Csconfigtype
    leafs["csConfigRDIVType"] = csconfigentry.Csconfigrdivtype
    leafs["csConfigRDIPType"] = csconfigentry.Csconfigrdiptype
    leafs["csTributaryType"] = csconfigentry.Cstributarytype
    leafs["csTributaryMappingType"] = csconfigentry.Cstributarymappingtype
    leafs["csTributaryFramingType"] = csconfigentry.Cstributaryframingtype
    leafs["csSignallingTransportMode"] = csconfigentry.Cssignallingtransportmode
    leafs["csTributaryGroupingType"] = csconfigentry.Cstributarygroupingtype
    return leafs
}

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetYangName() string { return "csConfigEntry" }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) SetParent(parent types.Entity) { csconfigentry.parent = parent }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetParent() types.Entity { return csconfigentry.parent }

func (csconfigentry *CISCOSONETMIB_Csconfigtable_Csconfigentry) GetParentYangName() string { return "csConfigTable" }

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
    parent types.Entity
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

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetFilter() yfilter.YFilter { return csapsconfigtable.YFilter }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) SetFilter(yf yfilter.YFilter) { csapsconfigtable.YFilter = yf }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetGoName(yname string) string {
    if yname == "csApsConfigEntry" { return "Csapsconfigentry" }
    return ""
}

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetSegmentPath() string {
    return "csApsConfigTable"
}

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csApsConfigEntry" {
        for _, c := range csapsconfigtable.Csapsconfigentry {
            if csapsconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry{}
        csapsconfigtable.Csapsconfigentry = append(csapsconfigtable.Csapsconfigentry, child)
        return &csapsconfigtable.Csapsconfigentry[len(csapsconfigtable.Csapsconfigentry)-1]
    }
    return nil
}

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csapsconfigtable.Csapsconfigentry {
        children[csapsconfigtable.Csapsconfigentry[i].GetSegmentPath()] = &csapsconfigtable.Csapsconfigentry[i]
    }
    return children
}

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetYangName() string { return "csApsConfigTable" }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) SetParent(parent types.Entity) { csapsconfigtable.parent = parent }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetParent() types.Entity { return csapsconfigtable.parent }

func (csapsconfigtable *CISCOSONETMIB_Csapsconfigtable) GetParentYangName() string { return "CISCO-SONET-MIB" }

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
    parent types.Entity
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

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetFilter() yfilter.YFilter { return csapsconfigentry.YFilter }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) SetFilter(yf yfilter.YFilter) { csapsconfigentry.YFilter = yf }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetGoName(yname string) string {
    if yname == "csApsWorkingIndex" { return "Csapsworkingindex" }
    if yname == "csApsProtectionIndex" { return "Csapsprotectionindex" }
    if yname == "csApsEnable" { return "Csapsenable" }
    if yname == "csApsArchMode" { return "Csapsarchmode" }
    if yname == "csApsActiveLine" { return "Csapsactiveline" }
    if yname == "csApsSigFaultBER" { return "Csapssigfaultber" }
    if yname == "csApsSigDegradeBER" { return "Csapssigdegradeber" }
    if yname == "csApsWaitToRestore" { return "Csapswaittorestore" }
    if yname == "csApsDirection" { return "Csapsdirection" }
    if yname == "csApsRevertive" { return "Csapsrevertive" }
    if yname == "csApsDirectionOperational" { return "Csapsdirectionoperational" }
    if yname == "csApsArchModeOperational" { return "Csapsarchmodeoperational" }
    if yname == "csApsChannelProtocol" { return "Csapschannelprotocol" }
    if yname == "csApsFailureStatus" { return "Csapsfailurestatus" }
    if yname == "csApsSwitchReason" { return "Csapsswitchreason" }
    if yname == "csApsPrimarySection" { return "Csapsprimarysection" }
    return ""
}

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetSegmentPath() string {
    return "csApsConfigEntry" + "[csApsWorkingIndex='" + fmt.Sprintf("%v", csapsconfigentry.Csapsworkingindex) + "']"
}

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csApsWorkingIndex"] = csapsconfigentry.Csapsworkingindex
    leafs["csApsProtectionIndex"] = csapsconfigentry.Csapsprotectionindex
    leafs["csApsEnable"] = csapsconfigentry.Csapsenable
    leafs["csApsArchMode"] = csapsconfigentry.Csapsarchmode
    leafs["csApsActiveLine"] = csapsconfigentry.Csapsactiveline
    leafs["csApsSigFaultBER"] = csapsconfigentry.Csapssigfaultber
    leafs["csApsSigDegradeBER"] = csapsconfigentry.Csapssigdegradeber
    leafs["csApsWaitToRestore"] = csapsconfigentry.Csapswaittorestore
    leafs["csApsDirection"] = csapsconfigentry.Csapsdirection
    leafs["csApsRevertive"] = csapsconfigentry.Csapsrevertive
    leafs["csApsDirectionOperational"] = csapsconfigentry.Csapsdirectionoperational
    leafs["csApsArchModeOperational"] = csapsconfigentry.Csapsarchmodeoperational
    leafs["csApsChannelProtocol"] = csapsconfigentry.Csapschannelprotocol
    leafs["csApsFailureStatus"] = csapsconfigentry.Csapsfailurestatus
    leafs["csApsSwitchReason"] = csapsconfigentry.Csapsswitchreason
    leafs["csApsPrimarySection"] = csapsconfigentry.Csapsprimarysection
    return leafs
}

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetYangName() string { return "csApsConfigEntry" }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) SetParent(parent types.Entity) { csapsconfigentry.parent = parent }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetParent() types.Entity { return csapsconfigentry.parent }

func (csapsconfigentry *CISCOSONETMIB_Csapsconfigtable_Csapsconfigentry) GetParentYangName() string { return "csApsConfigTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Csstotaltable_Csstotalentry.
    Csstotalentry []CISCOSONETMIB_Csstotaltable_Csstotalentry
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetFilter() yfilter.YFilter { return csstotaltable.YFilter }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) SetFilter(yf yfilter.YFilter) { csstotaltable.YFilter = yf }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetGoName(yname string) string {
    if yname == "cssTotalEntry" { return "Csstotalentry" }
    return ""
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetSegmentPath() string {
    return "cssTotalTable"
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cssTotalEntry" {
        for _, c := range csstotaltable.Csstotalentry {
            if csstotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csstotaltable_Csstotalentry{}
        csstotaltable.Csstotalentry = append(csstotaltable.Csstotalentry, child)
        return &csstotaltable.Csstotalentry[len(csstotaltable.Csstotalentry)-1]
    }
    return nil
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csstotaltable.Csstotalentry {
        children[csstotaltable.Csstotalentry[i].GetSegmentPath()] = &csstotaltable.Csstotalentry[i]
    }
    return children
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetYangName() string { return "cssTotalTable" }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) SetParent(parent types.Entity) { csstotaltable.parent = parent }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetParent() types.Entity { return csstotaltable.parent }

func (csstotaltable *CISCOSONETMIB_Csstotaltable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csstotaltable_Csstotalentry
// An entry in the SONET/SDH Section Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Csstotaltable_Csstotalentry struct {
    parent types.Entity
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

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetFilter() yfilter.YFilter { return csstotalentry.YFilter }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) SetFilter(yf yfilter.YFilter) { csstotalentry.YFilter = yf }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cssTotalESs" { return "Csstotaless" }
    if yname == "cssTotalSESs" { return "Csstotalsess" }
    if yname == "cssTotalSEFSs" { return "Csstotalsefss" }
    if yname == "cssTotalCVs" { return "Csstotalcvs" }
    return ""
}

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetSegmentPath() string {
    return "cssTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", csstotalentry.Ifindex) + "']"
}

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csstotalentry.Ifindex
    leafs["cssTotalESs"] = csstotalentry.Csstotaless
    leafs["cssTotalSESs"] = csstotalentry.Csstotalsess
    leafs["cssTotalSEFSs"] = csstotalentry.Csstotalsefss
    leafs["cssTotalCVs"] = csstotalentry.Csstotalcvs
    return leafs
}

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetYangName() string { return "cssTotalEntry" }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) SetParent(parent types.Entity) { csstotalentry.parent = parent }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetParent() types.Entity { return csstotalentry.parent }

func (csstotalentry *CISCOSONETMIB_Csstotaltable_Csstotalentry) GetParentYangName() string { return "cssTotalTable" }

// CISCOSONETMIB_Csstracetable
// The SONET/SDH Section Trace table. This table contains 
// objects for tracing the sonet section.
type CISCOSONETMIB_Csstracetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the trace table. Entries exist for active sonet lines. The
    // objects in this table are used to verify  continued connection between the
    // two ends of the line. The type is slice of
    // CISCOSONETMIB_Csstracetable_Csstraceentry.
    Csstraceentry []CISCOSONETMIB_Csstracetable_Csstraceentry
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetFilter() yfilter.YFilter { return csstracetable.YFilter }

func (csstracetable *CISCOSONETMIB_Csstracetable) SetFilter(yf yfilter.YFilter) { csstracetable.YFilter = yf }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetGoName(yname string) string {
    if yname == "cssTraceEntry" { return "Csstraceentry" }
    return ""
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetSegmentPath() string {
    return "cssTraceTable"
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cssTraceEntry" {
        for _, c := range csstracetable.Csstraceentry {
            if csstracetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csstracetable_Csstraceentry{}
        csstracetable.Csstraceentry = append(csstracetable.Csstraceentry, child)
        return &csstracetable.Csstraceentry[len(csstracetable.Csstraceentry)-1]
    }
    return nil
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csstracetable.Csstraceentry {
        children[csstracetable.Csstraceentry[i].GetSegmentPath()] = &csstracetable.Csstraceentry[i]
    }
    return children
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csstracetable *CISCOSONETMIB_Csstracetable) GetBundleName() string { return "cisco_ios_xe" }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetYangName() string { return "cssTraceTable" }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csstracetable *CISCOSONETMIB_Csstracetable) SetParent(parent types.Entity) { csstracetable.parent = parent }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetParent() types.Entity { return csstracetable.parent }

func (csstracetable *CISCOSONETMIB_Csstracetable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csstracetable_Csstraceentry
// An entry in the trace table. Entries exist for active sonet
// lines. The objects in this table are used to verify 
// continued connection between the two ends of the line.
type CISCOSONETMIB_Csstracetable_Csstraceentry struct {
    parent types.Entity
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

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetFilter() yfilter.YFilter { return csstraceentry.YFilter }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) SetFilter(yf yfilter.YFilter) { csstraceentry.YFilter = yf }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cssTraceToTransmit" { return "Csstracetotransmit" }
    if yname == "cssTraceToExpect" { return "Csstracetoexpect" }
    if yname == "cssTraceFailure" { return "Csstracefailure" }
    if yname == "cssTraceReceived" { return "Csstracereceived" }
    return ""
}

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetSegmentPath() string {
    return "cssTraceEntry" + "[ifIndex='" + fmt.Sprintf("%v", csstraceentry.Ifindex) + "']"
}

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csstraceentry.Ifindex
    leafs["cssTraceToTransmit"] = csstraceentry.Csstracetotransmit
    leafs["cssTraceToExpect"] = csstraceentry.Csstracetoexpect
    leafs["cssTraceFailure"] = csstraceentry.Csstracefailure
    leafs["cssTraceReceived"] = csstraceentry.Csstracereceived
    return leafs
}

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetBundleName() string { return "cisco_ios_xe" }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetYangName() string { return "cssTraceEntry" }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) SetParent(parent types.Entity) { csstraceentry.parent = parent }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetParent() types.Entity { return csstraceentry.parent }

func (csstraceentry *CISCOSONETMIB_Csstracetable_Csstraceentry) GetParentYangName() string { return "cssTraceTable" }

// CISCOSONETMIB_Csltotaltable
// The SONET/SDH Line Total table. It contains the 
// cumulative sum of the various statistics for the 24 
// hour period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the 
// number of 15 minute intervals that have elapsed since
// the line is enabled.
type CISCOSONETMIB_Csltotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Csltotaltable_Csltotalentry.
    Csltotalentry []CISCOSONETMIB_Csltotaltable_Csltotalentry
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetFilter() yfilter.YFilter { return csltotaltable.YFilter }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) SetFilter(yf yfilter.YFilter) { csltotaltable.YFilter = yf }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetGoName(yname string) string {
    if yname == "cslTotalEntry" { return "Csltotalentry" }
    return ""
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetSegmentPath() string {
    return "cslTotalTable"
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cslTotalEntry" {
        for _, c := range csltotaltable.Csltotalentry {
            if csltotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csltotaltable_Csltotalentry{}
        csltotaltable.Csltotalentry = append(csltotaltable.Csltotalentry, child)
        return &csltotaltable.Csltotalentry[len(csltotaltable.Csltotalentry)-1]
    }
    return nil
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csltotaltable.Csltotalentry {
        children[csltotaltable.Csltotalentry[i].GetSegmentPath()] = &csltotaltable.Csltotalentry[i]
    }
    return children
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetYangName() string { return "cslTotalTable" }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) SetParent(parent types.Entity) { csltotaltable.parent = parent }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetParent() types.Entity { return csltotaltable.parent }

func (csltotaltable *CISCOSONETMIB_Csltotaltable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csltotaltable_Csltotalentry
// An entry in the SONET/SDH Line Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Csltotaltable_Csltotalentry struct {
    parent types.Entity
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

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetFilter() yfilter.YFilter { return csltotalentry.YFilter }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) SetFilter(yf yfilter.YFilter) { csltotalentry.YFilter = yf }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cslTotalESs" { return "Csltotaless" }
    if yname == "cslTotalSESs" { return "Csltotalsess" }
    if yname == "cslTotalCVs" { return "Csltotalcvs" }
    if yname == "cslTotalUASs" { return "Csltotaluass" }
    return ""
}

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetSegmentPath() string {
    return "cslTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", csltotalentry.Ifindex) + "']"
}

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csltotalentry.Ifindex
    leafs["cslTotalESs"] = csltotalentry.Csltotaless
    leafs["cslTotalSESs"] = csltotalentry.Csltotalsess
    leafs["cslTotalCVs"] = csltotalentry.Csltotalcvs
    leafs["cslTotalUASs"] = csltotalentry.Csltotaluass
    return leafs
}

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetYangName() string { return "cslTotalEntry" }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) SetParent(parent types.Entity) { csltotalentry.parent = parent }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetParent() types.Entity { return csltotalentry.parent }

func (csltotalentry *CISCOSONETMIB_Csltotaltable_Csltotalentry) GetParentYangName() string { return "cslTotalTable" }

// CISCOSONETMIB_Cslfarendtotaltable
// The SONET/SDH Far End Line Total table. It contains the 
// cumulative sum of the various statistics for the 24 hour 
// period preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the 
// number of 15 minute intervals that have elapsed since the 
// line is enabled.
type CISCOSONETMIB_Cslfarendtotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry.
    Cslfarendtotalentry []CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetFilter() yfilter.YFilter { return cslfarendtotaltable.YFilter }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) SetFilter(yf yfilter.YFilter) { cslfarendtotaltable.YFilter = yf }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetGoName(yname string) string {
    if yname == "cslFarEndTotalEntry" { return "Cslfarendtotalentry" }
    return ""
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetSegmentPath() string {
    return "cslFarEndTotalTable"
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cslFarEndTotalEntry" {
        for _, c := range cslfarendtotaltable.Cslfarendtotalentry {
            if cslfarendtotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry{}
        cslfarendtotaltable.Cslfarendtotalentry = append(cslfarendtotaltable.Cslfarendtotalentry, child)
        return &cslfarendtotaltable.Cslfarendtotalentry[len(cslfarendtotaltable.Cslfarendtotalentry)-1]
    }
    return nil
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cslfarendtotaltable.Cslfarendtotalentry {
        children[cslfarendtotaltable.Cslfarendtotalentry[i].GetSegmentPath()] = &cslfarendtotaltable.Cslfarendtotalentry[i]
    }
    return children
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetYangName() string { return "cslFarEndTotalTable" }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) SetParent(parent types.Entity) { cslfarendtotaltable.parent = parent }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetParent() types.Entity { return cslfarendtotaltable.parent }

func (cslfarendtotaltable *CISCOSONETMIB_Cslfarendtotaltable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry
// An entry in the SONET/SDH Far End Line Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry struct {
    parent types.Entity
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

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetFilter() yfilter.YFilter { return cslfarendtotalentry.YFilter }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) SetFilter(yf yfilter.YFilter) { cslfarendtotalentry.YFilter = yf }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cslFarEndTotalESs" { return "Cslfarendtotaless" }
    if yname == "cslFarEndTotalSESs" { return "Cslfarendtotalsess" }
    if yname == "cslFarEndTotalCVs" { return "Cslfarendtotalcvs" }
    if yname == "cslFarEndTotalUASs" { return "Cslfarendtotaluass" }
    return ""
}

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetSegmentPath() string {
    return "cslFarEndTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", cslfarendtotalentry.Ifindex) + "']"
}

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cslfarendtotalentry.Ifindex
    leafs["cslFarEndTotalESs"] = cslfarendtotalentry.Cslfarendtotaless
    leafs["cslFarEndTotalSESs"] = cslfarendtotalentry.Cslfarendtotalsess
    leafs["cslFarEndTotalCVs"] = cslfarendtotalentry.Cslfarendtotalcvs
    leafs["cslFarEndTotalUASs"] = cslfarendtotalentry.Cslfarendtotaluass
    return leafs
}

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetYangName() string { return "cslFarEndTotalEntry" }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) SetParent(parent types.Entity) { cslfarendtotalentry.parent = parent }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetParent() types.Entity { return cslfarendtotalentry.parent }

func (cslfarendtotalentry *CISCOSONETMIB_Cslfarendtotaltable_Cslfarendtotalentry) GetParentYangName() string { return "cslFarEndTotalTable" }

// CISCOSONETMIB_Csptotaltable
// The SONET/SDH Path Total table. It contains the cumulative 
// sum of the various statistics for the 24 hour period 
// preceding the current interval.The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the number
// of 15 minute intervals that have elapsed since the line is 
// enabled.
type CISCOSONETMIB_Csptotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Total table. Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Csptotaltable_Csptotalentry.
    Csptotalentry []CISCOSONETMIB_Csptotaltable_Csptotalentry
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetFilter() yfilter.YFilter { return csptotaltable.YFilter }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) SetFilter(yf yfilter.YFilter) { csptotaltable.YFilter = yf }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetGoName(yname string) string {
    if yname == "cspTotalEntry" { return "Csptotalentry" }
    return ""
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetSegmentPath() string {
    return "cspTotalTable"
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cspTotalEntry" {
        for _, c := range csptotaltable.Csptotalentry {
            if csptotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csptotaltable_Csptotalentry{}
        csptotaltable.Csptotalentry = append(csptotaltable.Csptotalentry, child)
        return &csptotaltable.Csptotalentry[len(csptotaltable.Csptotalentry)-1]
    }
    return nil
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csptotaltable.Csptotalentry {
        children[csptotaltable.Csptotalentry[i].GetSegmentPath()] = &csptotaltable.Csptotalentry[i]
    }
    return children
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetYangName() string { return "cspTotalTable" }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) SetParent(parent types.Entity) { csptotaltable.parent = parent }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetParent() types.Entity { return csptotaltable.parent }

func (csptotaltable *CISCOSONETMIB_Csptotaltable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csptotaltable_Csptotalentry
// An entry in the SONET/SDH Path Total table. Entries
// are created automatically for sonet lines.
type CISCOSONETMIB_Csptotaltable_Csptotalentry struct {
    parent types.Entity
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

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetFilter() yfilter.YFilter { return csptotalentry.YFilter }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) SetFilter(yf yfilter.YFilter) { csptotalentry.YFilter = yf }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cspTotalESs" { return "Csptotaless" }
    if yname == "cspTotalSESs" { return "Csptotalsess" }
    if yname == "cspTotalCVs" { return "Csptotalcvs" }
    if yname == "cspTotalUASs" { return "Csptotaluass" }
    return ""
}

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetSegmentPath() string {
    return "cspTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", csptotalentry.Ifindex) + "']"
}

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csptotalentry.Ifindex
    leafs["cspTotalESs"] = csptotalentry.Csptotaless
    leafs["cspTotalSESs"] = csptotalentry.Csptotalsess
    leafs["cspTotalCVs"] = csptotalentry.Csptotalcvs
    leafs["cspTotalUASs"] = csptotalentry.Csptotaluass
    return leafs
}

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetYangName() string { return "cspTotalEntry" }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) SetParent(parent types.Entity) { csptotalentry.parent = parent }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetParent() types.Entity { return csptotalentry.parent }

func (csptotalentry *CISCOSONETMIB_Csptotaltable_Csptotalentry) GetParentYangName() string { return "cspTotalTable" }

// CISCOSONETMIB_Cspfarendtotaltable
// The SONET/SDH Far End Path Total table. Far End is the 
// remote end of the line. The table contains the cumulative
// sum of the various statistics for the 24 hour period 
// preceding the current interval. The object 
// 'sonetMediumValidIntervals' from RFC2558 contains the
// number of 15 minute intervals that have elapsed since
// the line is enabled. 
type CISCOSONETMIB_Cspfarendtotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Total table.  Entries are created
    // automatically for sonet lines. The type is slice of
    // CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry.
    Cspfarendtotalentry []CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetFilter() yfilter.YFilter { return cspfarendtotaltable.YFilter }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) SetFilter(yf yfilter.YFilter) { cspfarendtotaltable.YFilter = yf }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetGoName(yname string) string {
    if yname == "cspFarEndTotalEntry" { return "Cspfarendtotalentry" }
    return ""
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetSegmentPath() string {
    return "cspFarEndTotalTable"
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cspFarEndTotalEntry" {
        for _, c := range cspfarendtotaltable.Cspfarendtotalentry {
            if cspfarendtotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry{}
        cspfarendtotaltable.Cspfarendtotalentry = append(cspfarendtotaltable.Cspfarendtotalentry, child)
        return &cspfarendtotaltable.Cspfarendtotalentry[len(cspfarendtotaltable.Cspfarendtotalentry)-1]
    }
    return nil
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cspfarendtotaltable.Cspfarendtotalentry {
        children[cspfarendtotaltable.Cspfarendtotalentry[i].GetSegmentPath()] = &cspfarendtotaltable.Cspfarendtotalentry[i]
    }
    return children
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetYangName() string { return "cspFarEndTotalTable" }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) SetParent(parent types.Entity) { cspfarendtotaltable.parent = parent }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetParent() types.Entity { return cspfarendtotaltable.parent }

func (cspfarendtotaltable *CISCOSONETMIB_Cspfarendtotaltable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry
// An entry in the SONET/SDH Far End Path Total table. 
// Entries are created automatically for sonet lines.
type CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry struct {
    parent types.Entity
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

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetFilter() yfilter.YFilter { return cspfarendtotalentry.YFilter }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) SetFilter(yf yfilter.YFilter) { cspfarendtotalentry.YFilter = yf }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cspFarEndTotalESs" { return "Cspfarendtotaless" }
    if yname == "cspFarEndTotalSESs" { return "Cspfarendtotalsess" }
    if yname == "cspFarEndTotalCVs" { return "Cspfarendtotalcvs" }
    if yname == "cspFarEndTotalUASs" { return "Cspfarendtotaluass" }
    return ""
}

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetSegmentPath() string {
    return "cspFarEndTotalEntry" + "[ifIndex='" + fmt.Sprintf("%v", cspfarendtotalentry.Ifindex) + "']"
}

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cspfarendtotalentry.Ifindex
    leafs["cspFarEndTotalESs"] = cspfarendtotalentry.Cspfarendtotaless
    leafs["cspFarEndTotalSESs"] = cspfarendtotalentry.Cspfarendtotalsess
    leafs["cspFarEndTotalCVs"] = cspfarendtotalentry.Cspfarendtotalcvs
    leafs["cspFarEndTotalUASs"] = cspfarendtotalentry.Cspfarendtotaluass
    return leafs
}

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetYangName() string { return "cspFarEndTotalEntry" }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) SetParent(parent types.Entity) { cspfarendtotalentry.parent = parent }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetParent() types.Entity { return cspfarendtotalentry.parent }

func (cspfarendtotalentry *CISCOSONETMIB_Cspfarendtotaltable_Cspfarendtotalentry) GetParentYangName() string { return "cspFarEndTotalTable" }

// CISCOSONETMIB_Csptracetable
// The SONET/SDH Path Trace table. This table contains objects 
// for tracing the sonet path.
type CISCOSONETMIB_Csptracetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Trace table. The entries  exist for active
    // sonet lines. The objects in this table are  used to verify continued
    // connection between the two ends of the line. The type is slice of
    // CISCOSONETMIB_Csptracetable_Csptraceentry.
    Csptraceentry []CISCOSONETMIB_Csptracetable_Csptraceentry
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetFilter() yfilter.YFilter { return csptracetable.YFilter }

func (csptracetable *CISCOSONETMIB_Csptracetable) SetFilter(yf yfilter.YFilter) { csptracetable.YFilter = yf }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetGoName(yname string) string {
    if yname == "cspTraceEntry" { return "Csptraceentry" }
    return ""
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetSegmentPath() string {
    return "cspTraceTable"
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cspTraceEntry" {
        for _, c := range csptracetable.Csptraceentry {
            if csptracetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csptracetable_Csptraceentry{}
        csptracetable.Csptraceentry = append(csptracetable.Csptraceentry, child)
        return &csptracetable.Csptraceentry[len(csptracetable.Csptraceentry)-1]
    }
    return nil
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csptracetable.Csptraceentry {
        children[csptracetable.Csptraceentry[i].GetSegmentPath()] = &csptracetable.Csptraceentry[i]
    }
    return children
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csptracetable *CISCOSONETMIB_Csptracetable) GetBundleName() string { return "cisco_ios_xe" }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetYangName() string { return "cspTraceTable" }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csptracetable *CISCOSONETMIB_Csptracetable) SetParent(parent types.Entity) { csptracetable.parent = parent }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetParent() types.Entity { return csptracetable.parent }

func (csptracetable *CISCOSONETMIB_Csptracetable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csptracetable_Csptraceentry
// An entry in the SONET/SDH Path Trace table. The entries 
// exist for active sonet lines. The objects in this table are 
// used to verify continued connection between the two ends of
// the line.
type CISCOSONETMIB_Csptracetable_Csptraceentry struct {
    parent types.Entity
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

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetFilter() yfilter.YFilter { return csptraceentry.YFilter }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) SetFilter(yf yfilter.YFilter) { csptraceentry.YFilter = yf }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cspTraceToTransmit" { return "Csptracetotransmit" }
    if yname == "cspTraceToExpect" { return "Csptracetoexpect" }
    if yname == "cspTraceFailure" { return "Csptracefailure" }
    if yname == "cspTraceReceived" { return "Csptracereceived" }
    return ""
}

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetSegmentPath() string {
    return "cspTraceEntry" + "[ifIndex='" + fmt.Sprintf("%v", csptraceentry.Ifindex) + "']"
}

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csptraceentry.Ifindex
    leafs["cspTraceToTransmit"] = csptraceentry.Csptracetotransmit
    leafs["cspTraceToExpect"] = csptraceentry.Csptracetoexpect
    leafs["cspTraceFailure"] = csptraceentry.Csptracefailure
    leafs["cspTraceReceived"] = csptraceentry.Csptracereceived
    return leafs
}

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetBundleName() string { return "cisco_ios_xe" }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetYangName() string { return "cspTraceEntry" }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) SetParent(parent types.Entity) { csptraceentry.parent = parent }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetParent() types.Entity { return csptraceentry.parent }

func (csptraceentry *CISCOSONETMIB_Csptracetable_Csptraceentry) GetParentYangName() string { return "cspTraceTable" }

// CISCOSONETMIB_Csstatstable
// The SONET/SDH Section statistics table. This table 
// maintains the number of times the line encountered Loss of
// Signal(LOS), Loss of frame(LOF), Alarm Indication 
// signals(AISs), Remote failure indications(RFIs).
type CISCOSONETMIB_Csstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH statistics table. These are  realtime statistics
    // for the Sonet section, line and path layers. The statistics are gathered
    // for each sonet line.  An entry is created automatically and is indexed by 
    // ifIndex. The type is slice of CISCOSONETMIB_Csstatstable_Csstatsentry.
    Csstatsentry []CISCOSONETMIB_Csstatstable_Csstatsentry
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetFilter() yfilter.YFilter { return csstatstable.YFilter }

func (csstatstable *CISCOSONETMIB_Csstatstable) SetFilter(yf yfilter.YFilter) { csstatstable.YFilter = yf }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetGoName(yname string) string {
    if yname == "csStatsEntry" { return "Csstatsentry" }
    return ""
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetSegmentPath() string {
    return "csStatsTable"
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csStatsEntry" {
        for _, c := range csstatstable.Csstatsentry {
            if csstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csstatstable_Csstatsentry{}
        csstatstable.Csstatsentry = append(csstatstable.Csstatsentry, child)
        return &csstatstable.Csstatsentry[len(csstatstable.Csstatsentry)-1]
    }
    return nil
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csstatstable.Csstatsentry {
        children[csstatstable.Csstatsentry[i].GetSegmentPath()] = &csstatstable.Csstatsentry[i]
    }
    return children
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csstatstable *CISCOSONETMIB_Csstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetYangName() string { return "csStatsTable" }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csstatstable *CISCOSONETMIB_Csstatstable) SetParent(parent types.Entity) { csstatstable.parent = parent }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetParent() types.Entity { return csstatstable.parent }

func (csstatstable *CISCOSONETMIB_Csstatstable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csstatstable_Csstatsentry
// An entry in the SONET/SDH statistics table. These are 
// realtime statistics for the Sonet section, line and path
// layers. The statistics are gathered for each sonet line. 
// An entry is created automatically and is indexed by 
// ifIndex.
type CISCOSONETMIB_Csstatstable_Csstatsentry struct {
    parent types.Entity
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

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetFilter() yfilter.YFilter { return csstatsentry.YFilter }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) SetFilter(yf yfilter.YFilter) { csstatsentry.YFilter = yf }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cssLOSs" { return "Cssloss" }
    if yname == "cssLOFs" { return "Csslofs" }
    if yname == "cslAISs" { return "Cslaiss" }
    if yname == "cslRFIs" { return "Cslrfis" }
    if yname == "cspAISs" { return "Cspaiss" }
    if yname == "cspRFIs" { return "Csprfis" }
    return ""
}

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetSegmentPath() string {
    return "csStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", csstatsentry.Ifindex) + "']"
}

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csstatsentry.Ifindex
    leafs["cssLOSs"] = csstatsentry.Cssloss
    leafs["cssLOFs"] = csstatsentry.Csslofs
    leafs["cslAISs"] = csstatsentry.Cslaiss
    leafs["cslRFIs"] = csstatsentry.Cslrfis
    leafs["cspAISs"] = csstatsentry.Cspaiss
    leafs["cspRFIs"] = csstatsentry.Csprfis
    return leafs
}

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetYangName() string { return "csStatsEntry" }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) SetParent(parent types.Entity) { csstatsentry.parent = parent }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetParent() types.Entity { return csstatsentry.parent }

func (csstatsentry *CISCOSONETMIB_Csstatstable_Csstatsentry) GetParentYangName() string { return "csStatsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // There is an entry in this table for each TUG-3 within a  AU-4 SDH path that
    // supports SDH virtual container VC-4. The ifIndex value represents an entry
    // in ifTable with ifType = sonetPath(50).The ifTable entry applicable for
    // this entry belongs to AU-4 path. The type is slice of
    // CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry.
    Csau4Tug3Configentry []CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetFilter() yfilter.YFilter { return csau4Tug3Configtable.YFilter }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) SetFilter(yf yfilter.YFilter) { csau4Tug3Configtable.YFilter = yf }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetGoName(yname string) string {
    if yname == "csAu4Tug3ConfigEntry" { return "Csau4Tug3Configentry" }
    return ""
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetSegmentPath() string {
    return "csAu4Tug3ConfigTable"
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csAu4Tug3ConfigEntry" {
        for _, c := range csau4Tug3Configtable.Csau4Tug3Configentry {
            if csau4Tug3Configtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry{}
        csau4Tug3Configtable.Csau4Tug3Configentry = append(csau4Tug3Configtable.Csau4Tug3Configentry, child)
        return &csau4Tug3Configtable.Csau4Tug3Configentry[len(csau4Tug3Configtable.Csau4Tug3Configentry)-1]
    }
    return nil
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csau4Tug3Configtable.Csau4Tug3Configentry {
        children[csau4Tug3Configtable.Csau4Tug3Configentry[i].GetSegmentPath()] = &csau4Tug3Configtable.Csau4Tug3Configentry[i]
    }
    return children
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetBundleName() string { return "cisco_ios_xe" }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetYangName() string { return "csAu4Tug3ConfigTable" }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) SetParent(parent types.Entity) { csau4Tug3Configtable.parent = parent }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetParent() types.Entity { return csau4Tug3Configtable.parent }

func (csau4Tug3Configtable *CISCOSONETMIB_Csau4Tug3Configtable) GetParentYangName() string { return "CISCO-SONET-MIB" }

// CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry
// There is an entry in this table for each TUG-3 within a 
// AU-4 SDH path that supports SDH virtual container VC-4.
// The ifIndex value represents an entry in ifTable with
// ifType = sonetPath(50).The ifTable entry applicable for
// this entry belongs to AU-4 path.
type CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry struct {
    parent types.Entity
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

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetFilter() yfilter.YFilter { return csau4Tug3Configentry.YFilter }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) SetFilter(yf yfilter.YFilter) { csau4Tug3Configentry.YFilter = yf }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "csAu4Tug3" { return "Csau4Tug3" }
    if yname == "csAu4Tug3Payload" { return "Csau4Tug3Payload" }
    return ""
}

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetSegmentPath() string {
    return "csAu4Tug3ConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", csau4Tug3Configentry.Ifindex) + "']" + "[csAu4Tug3='" + fmt.Sprintf("%v", csau4Tug3Configentry.Csau4Tug3) + "']"
}

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csau4Tug3Configentry.Ifindex
    leafs["csAu4Tug3"] = csau4Tug3Configentry.Csau4Tug3
    leafs["csAu4Tug3Payload"] = csau4Tug3Configentry.Csau4Tug3Payload
    return leafs
}

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetBundleName() string { return "cisco_ios_xe" }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetYangName() string { return "csAu4Tug3ConfigEntry" }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) SetParent(parent types.Entity) { csau4Tug3Configentry.parent = parent }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetParent() types.Entity { return csau4Tug3Configentry.parent }

func (csau4Tug3Configentry *CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry) GetParentYangName() string { return "csAu4Tug3ConfigTable" }

// CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload represents The value 'other' can not be set.
type CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload string

const (
    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_other CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "other"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_vc11 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "vc11"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_vc12 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "vc12"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_tu3ds3 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "tu3ds3"

    CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload_tu3e3 CISCOSONETMIB_Csau4Tug3Configtable_Csau4Tug3Configentry_Csau4Tug3Payload = "tu3e3"
)

