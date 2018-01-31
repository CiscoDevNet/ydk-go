// The MIB module to describe SONET/SDH interface objects.
// 
// Copyright (C) The Internet Society (2003).  This version
// of this MIB module is part of RFC 3592;  see the RFC
// itself for full legal notices.
package sonet_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sonet_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SONET-MIB SONET-MIB}", reflect.TypeOf(SONETMIB{}))
    ydk.RegisterEntity("SONET-MIB:SONET-MIB", reflect.TypeOf(SONETMIB{}))
}

// SONETMIB
type SONETMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Sonetmedium SONETMIB_Sonetmedium

    // The SONET/SDH Medium table.
    Sonetmediumtable SONETMIB_Sonetmediumtable

    // The SONET/SDH Section Current table.
    Sonetsectioncurrenttable SONETMIB_Sonetsectioncurrenttable

    // The SONET/SDH Section Interval table.
    Sonetsectionintervaltable SONETMIB_Sonetsectionintervaltable

    // The SONET/SDH Line Current table.
    Sonetlinecurrenttable SONETMIB_Sonetlinecurrenttable

    // The SONET/SDH Line Interval table.
    Sonetlineintervaltable SONETMIB_Sonetlineintervaltable

    // The SONET/SDH Far End Line Current table.
    Sonetfarendlinecurrenttable SONETMIB_Sonetfarendlinecurrenttable

    // The SONET/SDH Far End Line Interval table.
    Sonetfarendlineintervaltable SONETMIB_Sonetfarendlineintervaltable

    // The SONET/SDH Path Current table.
    Sonetpathcurrenttable SONETMIB_Sonetpathcurrenttable

    // The SONET/SDH Path Interval table.
    Sonetpathintervaltable SONETMIB_Sonetpathintervaltable

    // The SONET/SDH Far End Path Current table.
    Sonetfarendpathcurrenttable SONETMIB_Sonetfarendpathcurrenttable

    // The SONET/SDH Far End Path Interval table.
    Sonetfarendpathintervaltable SONETMIB_Sonetfarendpathintervaltable

    // The SONET/SDH VT Current table.
    Sonetvtcurrenttable SONETMIB_Sonetvtcurrenttable

    // The SONET/SDH VT Interval table.
    Sonetvtintervaltable SONETMIB_Sonetvtintervaltable

    // The SONET/SDH Far End VT Current table.
    Sonetfarendvtcurrenttable SONETMIB_Sonetfarendvtcurrenttable

    // The SONET/SDH Far End VT Interval table.
    Sonetfarendvtintervaltable SONETMIB_Sonetfarendvtintervaltable
}

func (sONETMIB *SONETMIB) GetFilter() yfilter.YFilter { return sONETMIB.YFilter }

func (sONETMIB *SONETMIB) SetFilter(yf yfilter.YFilter) { sONETMIB.YFilter = yf }

func (sONETMIB *SONETMIB) GetGoName(yname string) string {
    if yname == "sonetMedium" { return "Sonetmedium" }
    if yname == "sonetMediumTable" { return "Sonetmediumtable" }
    if yname == "sonetSectionCurrentTable" { return "Sonetsectioncurrenttable" }
    if yname == "sonetSectionIntervalTable" { return "Sonetsectionintervaltable" }
    if yname == "sonetLineCurrentTable" { return "Sonetlinecurrenttable" }
    if yname == "sonetLineIntervalTable" { return "Sonetlineintervaltable" }
    if yname == "sonetFarEndLineCurrentTable" { return "Sonetfarendlinecurrenttable" }
    if yname == "sonetFarEndLineIntervalTable" { return "Sonetfarendlineintervaltable" }
    if yname == "sonetPathCurrentTable" { return "Sonetpathcurrenttable" }
    if yname == "sonetPathIntervalTable" { return "Sonetpathintervaltable" }
    if yname == "sonetFarEndPathCurrentTable" { return "Sonetfarendpathcurrenttable" }
    if yname == "sonetFarEndPathIntervalTable" { return "Sonetfarendpathintervaltable" }
    if yname == "sonetVTCurrentTable" { return "Sonetvtcurrenttable" }
    if yname == "sonetVTIntervalTable" { return "Sonetvtintervaltable" }
    if yname == "sonetFarEndVTCurrentTable" { return "Sonetfarendvtcurrenttable" }
    if yname == "sonetFarEndVTIntervalTable" { return "Sonetfarendvtintervaltable" }
    return ""
}

func (sONETMIB *SONETMIB) GetSegmentPath() string {
    return "SONET-MIB:SONET-MIB"
}

func (sONETMIB *SONETMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetMedium" {
        return &sONETMIB.Sonetmedium
    }
    if childYangName == "sonetMediumTable" {
        return &sONETMIB.Sonetmediumtable
    }
    if childYangName == "sonetSectionCurrentTable" {
        return &sONETMIB.Sonetsectioncurrenttable
    }
    if childYangName == "sonetSectionIntervalTable" {
        return &sONETMIB.Sonetsectionintervaltable
    }
    if childYangName == "sonetLineCurrentTable" {
        return &sONETMIB.Sonetlinecurrenttable
    }
    if childYangName == "sonetLineIntervalTable" {
        return &sONETMIB.Sonetlineintervaltable
    }
    if childYangName == "sonetFarEndLineCurrentTable" {
        return &sONETMIB.Sonetfarendlinecurrenttable
    }
    if childYangName == "sonetFarEndLineIntervalTable" {
        return &sONETMIB.Sonetfarendlineintervaltable
    }
    if childYangName == "sonetPathCurrentTable" {
        return &sONETMIB.Sonetpathcurrenttable
    }
    if childYangName == "sonetPathIntervalTable" {
        return &sONETMIB.Sonetpathintervaltable
    }
    if childYangName == "sonetFarEndPathCurrentTable" {
        return &sONETMIB.Sonetfarendpathcurrenttable
    }
    if childYangName == "sonetFarEndPathIntervalTable" {
        return &sONETMIB.Sonetfarendpathintervaltable
    }
    if childYangName == "sonetVTCurrentTable" {
        return &sONETMIB.Sonetvtcurrenttable
    }
    if childYangName == "sonetVTIntervalTable" {
        return &sONETMIB.Sonetvtintervaltable
    }
    if childYangName == "sonetFarEndVTCurrentTable" {
        return &sONETMIB.Sonetfarendvtcurrenttable
    }
    if childYangName == "sonetFarEndVTIntervalTable" {
        return &sONETMIB.Sonetfarendvtintervaltable
    }
    return nil
}

func (sONETMIB *SONETMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sonetMedium"] = &sONETMIB.Sonetmedium
    children["sonetMediumTable"] = &sONETMIB.Sonetmediumtable
    children["sonetSectionCurrentTable"] = &sONETMIB.Sonetsectioncurrenttable
    children["sonetSectionIntervalTable"] = &sONETMIB.Sonetsectionintervaltable
    children["sonetLineCurrentTable"] = &sONETMIB.Sonetlinecurrenttable
    children["sonetLineIntervalTable"] = &sONETMIB.Sonetlineintervaltable
    children["sonetFarEndLineCurrentTable"] = &sONETMIB.Sonetfarendlinecurrenttable
    children["sonetFarEndLineIntervalTable"] = &sONETMIB.Sonetfarendlineintervaltable
    children["sonetPathCurrentTable"] = &sONETMIB.Sonetpathcurrenttable
    children["sonetPathIntervalTable"] = &sONETMIB.Sonetpathintervaltable
    children["sonetFarEndPathCurrentTable"] = &sONETMIB.Sonetfarendpathcurrenttable
    children["sonetFarEndPathIntervalTable"] = &sONETMIB.Sonetfarendpathintervaltable
    children["sonetVTCurrentTable"] = &sONETMIB.Sonetvtcurrenttable
    children["sonetVTIntervalTable"] = &sONETMIB.Sonetvtintervaltable
    children["sonetFarEndVTCurrentTable"] = &sONETMIB.Sonetfarendvtcurrenttable
    children["sonetFarEndVTIntervalTable"] = &sONETMIB.Sonetfarendvtintervaltable
    return children
}

func (sONETMIB *SONETMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sONETMIB *SONETMIB) GetBundleName() string { return "cisco_ios_xe" }

func (sONETMIB *SONETMIB) GetYangName() string { return "SONET-MIB" }

func (sONETMIB *SONETMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sONETMIB *SONETMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sONETMIB *SONETMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sONETMIB *SONETMIB) SetParent(parent types.Entity) { sONETMIB.parent = parent }

func (sONETMIB *SONETMIB) GetParent() types.Entity { return sONETMIB.parent }

func (sONETMIB *SONETMIB) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetmedium
type SONETMIB_Sonetmedium struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An enumerated integer indicating which recognized set of SES thresholds
    // that the agent uses for determining severely errored seconds and
    // unavailable time.  other(1)   None of the following.  bellcore1991(2)  
    // Bellcore TR-NWT-000253, 1991 [TR253], or   ANSI T1M1.3/93-005R2, 1993
    // [T1M1.3].   See also Appendix B.  ansi1993(3)   ANSI T1.231, 1993
    // [T1.231a], or   Bellcore GR-253-CORE, Issue 2, 1995 [GR253]  itu1995(4)  
    // ITU Recommendation G.826, 1995 [G.826]  ansi1997(5)   ANSI T1.231, 1997
    // [T1.231b]  If a manager changes the value of this object then the SES
    // statistics collected prior to this change must be invalidated. The type is
    // Sonetsesthresholdset.
    Sonetsesthresholdset interface{}
}

func (sonetmedium *SONETMIB_Sonetmedium) GetFilter() yfilter.YFilter { return sonetmedium.YFilter }

func (sonetmedium *SONETMIB_Sonetmedium) SetFilter(yf yfilter.YFilter) { sonetmedium.YFilter = yf }

func (sonetmedium *SONETMIB_Sonetmedium) GetGoName(yname string) string {
    if yname == "sonetSESthresholdSet" { return "Sonetsesthresholdset" }
    return ""
}

func (sonetmedium *SONETMIB_Sonetmedium) GetSegmentPath() string {
    return "sonetMedium"
}

func (sonetmedium *SONETMIB_Sonetmedium) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetmedium *SONETMIB_Sonetmedium) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetmedium *SONETMIB_Sonetmedium) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sonetSESthresholdSet"] = sonetmedium.Sonetsesthresholdset
    return leafs
}

func (sonetmedium *SONETMIB_Sonetmedium) GetBundleName() string { return "cisco_ios_xe" }

func (sonetmedium *SONETMIB_Sonetmedium) GetYangName() string { return "sonetMedium" }

func (sonetmedium *SONETMIB_Sonetmedium) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetmedium *SONETMIB_Sonetmedium) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetmedium *SONETMIB_Sonetmedium) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetmedium *SONETMIB_Sonetmedium) SetParent(parent types.Entity) { sonetmedium.parent = parent }

func (sonetmedium *SONETMIB_Sonetmedium) GetParent() types.Entity { return sonetmedium.parent }

func (sonetmedium *SONETMIB_Sonetmedium) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetmedium_Sonetsesthresholdset represents prior to this change must be invalidated.
type SONETMIB_Sonetmedium_Sonetsesthresholdset string

const (
    SONETMIB_Sonetmedium_Sonetsesthresholdset_other SONETMIB_Sonetmedium_Sonetsesthresholdset = "other"

    SONETMIB_Sonetmedium_Sonetsesthresholdset_bellcore1991 SONETMIB_Sonetmedium_Sonetsesthresholdset = "bellcore1991"

    SONETMIB_Sonetmedium_Sonetsesthresholdset_ansi1993 SONETMIB_Sonetmedium_Sonetsesthresholdset = "ansi1993"

    SONETMIB_Sonetmedium_Sonetsesthresholdset_itu1995 SONETMIB_Sonetmedium_Sonetsesthresholdset = "itu1995"

    SONETMIB_Sonetmedium_Sonetsesthresholdset_ansi1997 SONETMIB_Sonetmedium_Sonetsesthresholdset = "ansi1997"
)

// SONETMIB_Sonetmediumtable
// The SONET/SDH Medium table.
type SONETMIB_Sonetmediumtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Medium table. The type is slice of
    // SONETMIB_Sonetmediumtable_Sonetmediumentry.
    Sonetmediumentry []SONETMIB_Sonetmediumtable_Sonetmediumentry
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetFilter() yfilter.YFilter { return sonetmediumtable.YFilter }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) SetFilter(yf yfilter.YFilter) { sonetmediumtable.YFilter = yf }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetGoName(yname string) string {
    if yname == "sonetMediumEntry" { return "Sonetmediumentry" }
    return ""
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetSegmentPath() string {
    return "sonetMediumTable"
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetMediumEntry" {
        for _, c := range sonetmediumtable.Sonetmediumentry {
            if sonetmediumtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetmediumtable_Sonetmediumentry{}
        sonetmediumtable.Sonetmediumentry = append(sonetmediumtable.Sonetmediumentry, child)
        return &sonetmediumtable.Sonetmediumentry[len(sonetmediumtable.Sonetmediumentry)-1]
    }
    return nil
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetmediumtable.Sonetmediumentry {
        children[sonetmediumtable.Sonetmediumentry[i].GetSegmentPath()] = &sonetmediumtable.Sonetmediumentry[i]
    }
    return children
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetYangName() string { return "sonetMediumTable" }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) SetParent(parent types.Entity) { sonetmediumtable.parent = parent }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetParent() types.Entity { return sonetmediumtable.parent }

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetmediumtable_Sonetmediumentry
// An entry in the SONET/SDH Medium table.
type SONETMIB_Sonetmediumtable_Sonetmediumentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This variable identifies whether a SONET or a SDH signal is used across
    // this interface. The type is Sonetmediumtype.
    Sonetmediumtype interface{}

    // The number of seconds, including partial seconds, that have elapsed since
    // the beginning of the current measurement period. If, for some reason, such
    // as an adjustment in the system's time-of-day clock, the current interval
    // exceeds the maximum value, the agent will return the maximum value. The
    // type is interface{} with range: 1..900.
    Sonetmediumtimeelapsed interface{}

    // The number of previous 15-minute intervals for which data was collected. A
    // SONET/SDH interface must be capable of supporting at least n intervals. The
    // minimum value of n is 4. The default of n is 32. The maximum value of n is
    // 96. The value will be <n> unless the measurement was (re-)started within
    // the last (<n>*15) minutes, in which case the value will be the number of
    // complete 15 minute intervals for which the agent has at least some data. In
    // certain cases (e.g., in the case where the agent is a proxy) it is possible
    // that some intervals are unavailable.  In this case, this interval is the
    // maximum interval number for which data is available. . The type is
    // interface{} with range: 0..96.
    Sonetmediumvalidintervals interface{}

    // This variable describes the line coding for this interface. The B3ZS and
    // CMI are used for electrical SONET/SDH signals (STS-1 and STS-3). The
    // Non-Return to Zero (NRZ) and the Return to Zero are used for optical
    // SONET/SDH signals. The type is Sonetmediumlinecoding.
    Sonetmediumlinecoding interface{}

    // This variable describes the line type for this interface. The line types
    // are Short and Long Range Single Mode fiber or Multi-Mode fiber interfaces,
    // and coax and UTP for electrical interfaces.  The value sonetOther should be
    // used when the Line Type is not one of the listed values. The type is
    // Sonetmediumlinetype.
    Sonetmediumlinetype interface{}

    // This variable contains the transmission vendor's circuit identifier, for
    // the purpose of facilitating troubleshooting. Note that the circuit
    // identifier, if available, is also represented by ifPhysAddress. The type is
    // string with length: 0..255.
    Sonetmediumcircuitidentifier interface{}

    // The number of intervals in the range from 0 to sonetMediumValidIntervals
    // for which no data is available. This object will typically be zero except
    // in cases where the data for some intervals are not available (e.g., in
    // proxy situations). The type is interface{} with range: 0..96.
    Sonetmediuminvalidintervals interface{}

    // The current loopback state of the SONET/SDH interface.  The values mean:   
    // sonetNoLoop      Not in the loopback state. A device that is not     
    // capable of performing a loopback on this interface      shall always return
    // this value.    sonetFacilityLoop      The received signal at this interface
    // is looped back      out through the corresponding transmitter in the return
    // direction.    sonetTerminalLoop      The signal that is about to be
    // transmitted is connected      to the associated incoming receiver.   
    // sonetOtherLoop      Loopbacks that are not defined here. The type is
    // map[string]bool.
    Sonetmediumloopbackconfig interface{}
}

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetFilter() yfilter.YFilter { return sonetmediumentry.YFilter }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) SetFilter(yf yfilter.YFilter) { sonetmediumentry.YFilter = yf }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetMediumType" { return "Sonetmediumtype" }
    if yname == "sonetMediumTimeElapsed" { return "Sonetmediumtimeelapsed" }
    if yname == "sonetMediumValidIntervals" { return "Sonetmediumvalidintervals" }
    if yname == "sonetMediumLineCoding" { return "Sonetmediumlinecoding" }
    if yname == "sonetMediumLineType" { return "Sonetmediumlinetype" }
    if yname == "sonetMediumCircuitIdentifier" { return "Sonetmediumcircuitidentifier" }
    if yname == "sonetMediumInvalidIntervals" { return "Sonetmediuminvalidintervals" }
    if yname == "sonetMediumLoopbackConfig" { return "Sonetmediumloopbackconfig" }
    return ""
}

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetSegmentPath() string {
    return "sonetMediumEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetmediumentry.Ifindex) + "']"
}

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetmediumentry.Ifindex
    leafs["sonetMediumType"] = sonetmediumentry.Sonetmediumtype
    leafs["sonetMediumTimeElapsed"] = sonetmediumentry.Sonetmediumtimeelapsed
    leafs["sonetMediumValidIntervals"] = sonetmediumentry.Sonetmediumvalidintervals
    leafs["sonetMediumLineCoding"] = sonetmediumentry.Sonetmediumlinecoding
    leafs["sonetMediumLineType"] = sonetmediumentry.Sonetmediumlinetype
    leafs["sonetMediumCircuitIdentifier"] = sonetmediumentry.Sonetmediumcircuitidentifier
    leafs["sonetMediumInvalidIntervals"] = sonetmediumentry.Sonetmediuminvalidintervals
    leafs["sonetMediumLoopbackConfig"] = sonetmediumentry.Sonetmediumloopbackconfig
    return leafs
}

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetYangName() string { return "sonetMediumEntry" }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) SetParent(parent types.Entity) { sonetmediumentry.parent = parent }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetParent() types.Entity { return sonetmediumentry.parent }

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetParentYangName() string { return "sonetMediumTable" }

// SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding represents to Zero are used for optical SONET/SDH signals.
type SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding string

const (
    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding_sonetMediumOther SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding = "sonetMediumOther"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding_sonetMediumB3ZS SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding = "sonetMediumB3ZS"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding_sonetMediumCMI SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding = "sonetMediumCMI"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding_sonetMediumNRZ SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding = "sonetMediumNRZ"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding_sonetMediumRZ SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinecoding = "sonetMediumRZ"
)

// SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype represents not one of the listed values.
type SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype string

const (
    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype_sonetOther SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype = "sonetOther"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype_sonetShortSingleMode SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype = "sonetShortSingleMode"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype_sonetLongSingleMode SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype = "sonetLongSingleMode"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype_sonetMultiMode SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype = "sonetMultiMode"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype_sonetCoax SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype = "sonetCoax"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype_sonetUTP SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumlinetype = "sonetUTP"
)

// SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumtype represents or a SDH signal is used across this interface.
type SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumtype string

const (
    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumtype_sonet SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumtype = "sonet"

    SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumtype_sdh SONETMIB_Sonetmediumtable_Sonetmediumentry_Sonetmediumtype = "sdh"
)

// SONETMIB_Sonetsectioncurrenttable
// The SONET/SDH Section Current table.
type SONETMIB_Sonetsectioncurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Current table. The type is slice of
    // SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry.
    Sonetsectioncurrententry []SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetFilter() yfilter.YFilter { return sonetsectioncurrenttable.YFilter }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) SetFilter(yf yfilter.YFilter) { sonetsectioncurrenttable.YFilter = yf }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetGoName(yname string) string {
    if yname == "sonetSectionCurrentEntry" { return "Sonetsectioncurrententry" }
    return ""
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetSegmentPath() string {
    return "sonetSectionCurrentTable"
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetSectionCurrentEntry" {
        for _, c := range sonetsectioncurrenttable.Sonetsectioncurrententry {
            if sonetsectioncurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry{}
        sonetsectioncurrenttable.Sonetsectioncurrententry = append(sonetsectioncurrenttable.Sonetsectioncurrententry, child)
        return &sonetsectioncurrenttable.Sonetsectioncurrententry[len(sonetsectioncurrenttable.Sonetsectioncurrententry)-1]
    }
    return nil
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetsectioncurrenttable.Sonetsectioncurrententry {
        children[sonetsectioncurrenttable.Sonetsectioncurrententry[i].GetSegmentPath()] = &sonetsectioncurrenttable.Sonetsectioncurrententry[i]
    }
    return children
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetYangName() string { return "sonetSectionCurrentTable" }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) SetParent(parent types.Entity) { sonetsectioncurrenttable.parent = parent }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetParent() types.Entity { return sonetsectioncurrenttable.parent }

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry
// An entry in the SONET/SDH Section Current table.
type SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This variable indicates the status of the interface. The
    // sonetSectionCurrentStatus is a bit map represented as a sum, therefore, it
    // can represent multiple defects simultaneously. The sonetSectionNoDefect
    // should be set if and only if no other flag is set.  The various bit
    // positions are:       1   sonetSectionNoDefect       2   sonetSectionLOS    
    // 4   sonetSectionLOF. The type is interface{} with range: 1..6.
    Sonetsectioncurrentstatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Section in the current 15 minute interval. The type is
    // interface{} with range: 0..4294967295.
    Sonetsectioncurrentess interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Section in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    Sonetsectioncurrentsess interface{}

    // The counter associated with the number of Severely Errored Framing Seconds
    // encountered by a SONET/SDH Section in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    Sonetsectioncurrentsefss interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Section in the current 15 minute interval. The type is
    // interface{} with range: 0..4294967295.
    Sonetsectioncurrentcvs interface{}
}

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetFilter() yfilter.YFilter { return sonetsectioncurrententry.YFilter }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) SetFilter(yf yfilter.YFilter) { sonetsectioncurrententry.YFilter = yf }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetSectionCurrentStatus" { return "Sonetsectioncurrentstatus" }
    if yname == "sonetSectionCurrentESs" { return "Sonetsectioncurrentess" }
    if yname == "sonetSectionCurrentSESs" { return "Sonetsectioncurrentsess" }
    if yname == "sonetSectionCurrentSEFSs" { return "Sonetsectioncurrentsefss" }
    if yname == "sonetSectionCurrentCVs" { return "Sonetsectioncurrentcvs" }
    return ""
}

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetSegmentPath() string {
    return "sonetSectionCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetsectioncurrententry.Ifindex) + "']"
}

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetsectioncurrententry.Ifindex
    leafs["sonetSectionCurrentStatus"] = sonetsectioncurrententry.Sonetsectioncurrentstatus
    leafs["sonetSectionCurrentESs"] = sonetsectioncurrententry.Sonetsectioncurrentess
    leafs["sonetSectionCurrentSESs"] = sonetsectioncurrententry.Sonetsectioncurrentsess
    leafs["sonetSectionCurrentSEFSs"] = sonetsectioncurrententry.Sonetsectioncurrentsefss
    leafs["sonetSectionCurrentCVs"] = sonetsectioncurrententry.Sonetsectioncurrentcvs
    return leafs
}

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetYangName() string { return "sonetSectionCurrentEntry" }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) SetParent(parent types.Entity) { sonetsectioncurrententry.parent = parent }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetParent() types.Entity { return sonetsectioncurrententry.parent }

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetParentYangName() string { return "sonetSectionCurrentTable" }

// SONETMIB_Sonetsectionintervaltable
// The SONET/SDH Section Interval table.
type SONETMIB_Sonetsectionintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Interval table. The type is slice of
    // SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry.
    Sonetsectionintervalentry []SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetFilter() yfilter.YFilter { return sonetsectionintervaltable.YFilter }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) SetFilter(yf yfilter.YFilter) { sonetsectionintervaltable.YFilter = yf }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetGoName(yname string) string {
    if yname == "sonetSectionIntervalEntry" { return "Sonetsectionintervalentry" }
    return ""
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetSegmentPath() string {
    return "sonetSectionIntervalTable"
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetSectionIntervalEntry" {
        for _, c := range sonetsectionintervaltable.Sonetsectionintervalentry {
            if sonetsectionintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry{}
        sonetsectionintervaltable.Sonetsectionintervalentry = append(sonetsectionintervaltable.Sonetsectionintervalentry, child)
        return &sonetsectionintervaltable.Sonetsectionintervalentry[len(sonetsectionintervaltable.Sonetsectionintervalentry)-1]
    }
    return nil
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetsectionintervaltable.Sonetsectionintervalentry {
        children[sonetsectionintervaltable.Sonetsectionintervalentry[i].GetSegmentPath()] = &sonetsectionintervaltable.Sonetsectionintervalentry[i]
    }
    return children
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetYangName() string { return "sonetSectionIntervalTable" }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) SetParent(parent types.Entity) { sonetsectionintervaltable.parent = parent }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetParent() types.Entity { return sonetsectionintervaltable.parent }

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry
// An entry in the SONET/SDH Section Interval table.
type SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetsectionintervalnumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Section in a particular 15-minute interval in the past 24 hours.
    // The type is interface{} with range: 0..4294967295.
    Sonetsectionintervaless interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Section in a particular 15-minute interval in
    // the past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetsectionintervalsess interface{}

    // The counter associated with the number of Severely Errored Framing Seconds
    // encountered by a SONET/SDH Section in a particular 15-minute interval in
    // the past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetsectionintervalsefss interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Section in a particular 15-minute interval in the past 24
    // hours. The type is interface{} with range: 0..4294967295.
    Sonetsectionintervalcvs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetsectionintervalvaliddata interface{}
}

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetFilter() yfilter.YFilter { return sonetsectionintervalentry.YFilter }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) SetFilter(yf yfilter.YFilter) { sonetsectionintervalentry.YFilter = yf }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetSectionIntervalNumber" { return "Sonetsectionintervalnumber" }
    if yname == "sonetSectionIntervalESs" { return "Sonetsectionintervaless" }
    if yname == "sonetSectionIntervalSESs" { return "Sonetsectionintervalsess" }
    if yname == "sonetSectionIntervalSEFSs" { return "Sonetsectionintervalsefss" }
    if yname == "sonetSectionIntervalCVs" { return "Sonetsectionintervalcvs" }
    if yname == "sonetSectionIntervalValidData" { return "Sonetsectionintervalvaliddata" }
    return ""
}

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetSegmentPath() string {
    return "sonetSectionIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetsectionintervalentry.Ifindex) + "']" + "[sonetSectionIntervalNumber='" + fmt.Sprintf("%v", sonetsectionintervalentry.Sonetsectionintervalnumber) + "']"
}

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetsectionintervalentry.Ifindex
    leafs["sonetSectionIntervalNumber"] = sonetsectionintervalentry.Sonetsectionintervalnumber
    leafs["sonetSectionIntervalESs"] = sonetsectionintervalentry.Sonetsectionintervaless
    leafs["sonetSectionIntervalSESs"] = sonetsectionintervalentry.Sonetsectionintervalsess
    leafs["sonetSectionIntervalSEFSs"] = sonetsectionintervalentry.Sonetsectionintervalsefss
    leafs["sonetSectionIntervalCVs"] = sonetsectionintervalentry.Sonetsectionintervalcvs
    leafs["sonetSectionIntervalValidData"] = sonetsectionintervalentry.Sonetsectionintervalvaliddata
    return leafs
}

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetYangName() string { return "sonetSectionIntervalEntry" }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) SetParent(parent types.Entity) { sonetsectionintervalentry.parent = parent }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetParent() types.Entity { return sonetsectionintervalentry.parent }

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetParentYangName() string { return "sonetSectionIntervalTable" }

// SONETMIB_Sonetlinecurrenttable
// The SONET/SDH Line Current table.
type SONETMIB_Sonetlinecurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Current table. The type is slice of
    // SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry.
    Sonetlinecurrententry []SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetFilter() yfilter.YFilter { return sonetlinecurrenttable.YFilter }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) SetFilter(yf yfilter.YFilter) { sonetlinecurrenttable.YFilter = yf }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetGoName(yname string) string {
    if yname == "sonetLineCurrentEntry" { return "Sonetlinecurrententry" }
    return ""
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetSegmentPath() string {
    return "sonetLineCurrentTable"
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetLineCurrentEntry" {
        for _, c := range sonetlinecurrenttable.Sonetlinecurrententry {
            if sonetlinecurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry{}
        sonetlinecurrenttable.Sonetlinecurrententry = append(sonetlinecurrenttable.Sonetlinecurrententry, child)
        return &sonetlinecurrenttable.Sonetlinecurrententry[len(sonetlinecurrenttable.Sonetlinecurrententry)-1]
    }
    return nil
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetlinecurrenttable.Sonetlinecurrententry {
        children[sonetlinecurrenttable.Sonetlinecurrententry[i].GetSegmentPath()] = &sonetlinecurrenttable.Sonetlinecurrententry[i]
    }
    return children
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetYangName() string { return "sonetLineCurrentTable" }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) SetParent(parent types.Entity) { sonetlinecurrenttable.parent = parent }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetParent() types.Entity { return sonetlinecurrenttable.parent }

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry
// An entry in the SONET/SDH Line Current table.
type SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This variable indicates the status of the interface. The
    // sonetLineCurrentStatus is a bit map represented as a sum, therefore, it can
    // represent multiple defects simultaneously. The sonetLineNoDefect should be
    // set if and only if no other flag is set.  The various bit positions are:  1
    // sonetLineNoDefect  2   sonetLineAIS  4   sonetLineRDI. The type is
    // interface{} with range: 1..6.
    Sonetlinecurrentstatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Line in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    Sonetlinecurrentess interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Line in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    Sonetlinecurrentsess interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Line in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    Sonetlinecurrentcvs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a SONET/SDH Line in the current 15 minute interval. The type is
    // interface{} with range: 0..4294967295.
    Sonetlinecurrentuass interface{}
}

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetFilter() yfilter.YFilter { return sonetlinecurrententry.YFilter }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) SetFilter(yf yfilter.YFilter) { sonetlinecurrententry.YFilter = yf }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetLineCurrentStatus" { return "Sonetlinecurrentstatus" }
    if yname == "sonetLineCurrentESs" { return "Sonetlinecurrentess" }
    if yname == "sonetLineCurrentSESs" { return "Sonetlinecurrentsess" }
    if yname == "sonetLineCurrentCVs" { return "Sonetlinecurrentcvs" }
    if yname == "sonetLineCurrentUASs" { return "Sonetlinecurrentuass" }
    return ""
}

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetSegmentPath() string {
    return "sonetLineCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetlinecurrententry.Ifindex) + "']"
}

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetlinecurrententry.Ifindex
    leafs["sonetLineCurrentStatus"] = sonetlinecurrententry.Sonetlinecurrentstatus
    leafs["sonetLineCurrentESs"] = sonetlinecurrententry.Sonetlinecurrentess
    leafs["sonetLineCurrentSESs"] = sonetlinecurrententry.Sonetlinecurrentsess
    leafs["sonetLineCurrentCVs"] = sonetlinecurrententry.Sonetlinecurrentcvs
    leafs["sonetLineCurrentUASs"] = sonetlinecurrententry.Sonetlinecurrentuass
    return leafs
}

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetYangName() string { return "sonetLineCurrentEntry" }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) SetParent(parent types.Entity) { sonetlinecurrententry.parent = parent }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetParent() types.Entity { return sonetlinecurrententry.parent }

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetParentYangName() string { return "sonetLineCurrentTable" }

// SONETMIB_Sonetlineintervaltable
// The SONET/SDH Line Interval table.
type SONETMIB_Sonetlineintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Interval table. The type is slice of
    // SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry.
    Sonetlineintervalentry []SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetFilter() yfilter.YFilter { return sonetlineintervaltable.YFilter }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) SetFilter(yf yfilter.YFilter) { sonetlineintervaltable.YFilter = yf }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetGoName(yname string) string {
    if yname == "sonetLineIntervalEntry" { return "Sonetlineintervalentry" }
    return ""
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetSegmentPath() string {
    return "sonetLineIntervalTable"
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetLineIntervalEntry" {
        for _, c := range sonetlineintervaltable.Sonetlineintervalentry {
            if sonetlineintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry{}
        sonetlineintervaltable.Sonetlineintervalentry = append(sonetlineintervaltable.Sonetlineintervalentry, child)
        return &sonetlineintervaltable.Sonetlineintervalentry[len(sonetlineintervaltable.Sonetlineintervalentry)-1]
    }
    return nil
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetlineintervaltable.Sonetlineintervalentry {
        children[sonetlineintervaltable.Sonetlineintervalentry[i].GetSegmentPath()] = &sonetlineintervaltable.Sonetlineintervalentry[i]
    }
    return children
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetYangName() string { return "sonetLineIntervalTable" }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) SetParent(parent types.Entity) { sonetlineintervaltable.parent = parent }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetParent() types.Entity { return sonetlineintervaltable.parent }

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry
// An entry in the SONET/SDH Line Interval table.
type SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetlineintervalnumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Line in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    Sonetlineintervaless interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Line in a particular 15-minute interval in the
    // past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetlineintervalsess interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Line in a particular 15-minute interval in the past 24 hours.
    // The type is interface{} with range: 0..4294967295.
    Sonetlineintervalcvs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a SONET/SDH Line in a particular 15-minute interval in the past 24
    // hours. The type is interface{} with range: 0..4294967295.
    Sonetlineintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetlineintervalvaliddata interface{}
}

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetFilter() yfilter.YFilter { return sonetlineintervalentry.YFilter }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) SetFilter(yf yfilter.YFilter) { sonetlineintervalentry.YFilter = yf }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetLineIntervalNumber" { return "Sonetlineintervalnumber" }
    if yname == "sonetLineIntervalESs" { return "Sonetlineintervaless" }
    if yname == "sonetLineIntervalSESs" { return "Sonetlineintervalsess" }
    if yname == "sonetLineIntervalCVs" { return "Sonetlineintervalcvs" }
    if yname == "sonetLineIntervalUASs" { return "Sonetlineintervaluass" }
    if yname == "sonetLineIntervalValidData" { return "Sonetlineintervalvaliddata" }
    return ""
}

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetSegmentPath() string {
    return "sonetLineIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetlineintervalentry.Ifindex) + "']" + "[sonetLineIntervalNumber='" + fmt.Sprintf("%v", sonetlineintervalentry.Sonetlineintervalnumber) + "']"
}

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetlineintervalentry.Ifindex
    leafs["sonetLineIntervalNumber"] = sonetlineintervalentry.Sonetlineintervalnumber
    leafs["sonetLineIntervalESs"] = sonetlineintervalentry.Sonetlineintervaless
    leafs["sonetLineIntervalSESs"] = sonetlineintervalentry.Sonetlineintervalsess
    leafs["sonetLineIntervalCVs"] = sonetlineintervalentry.Sonetlineintervalcvs
    leafs["sonetLineIntervalUASs"] = sonetlineintervalentry.Sonetlineintervaluass
    leafs["sonetLineIntervalValidData"] = sonetlineintervalentry.Sonetlineintervalvaliddata
    return leafs
}

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetYangName() string { return "sonetLineIntervalEntry" }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) SetParent(parent types.Entity) { sonetlineintervalentry.parent = parent }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetParent() types.Entity { return sonetlineintervalentry.parent }

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetParentYangName() string { return "sonetLineIntervalTable" }

// SONETMIB_Sonetfarendlinecurrenttable
// The SONET/SDH Far End Line Current table.
type SONETMIB_Sonetfarendlinecurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Current table. The type is slice of
    // SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry.
    Sonetfarendlinecurrententry []SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetFilter() yfilter.YFilter { return sonetfarendlinecurrenttable.YFilter }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) SetFilter(yf yfilter.YFilter) { sonetfarendlinecurrenttable.YFilter = yf }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetGoName(yname string) string {
    if yname == "sonetFarEndLineCurrentEntry" { return "Sonetfarendlinecurrententry" }
    return ""
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetSegmentPath() string {
    return "sonetFarEndLineCurrentTable"
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetFarEndLineCurrentEntry" {
        for _, c := range sonetfarendlinecurrenttable.Sonetfarendlinecurrententry {
            if sonetfarendlinecurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry{}
        sonetfarendlinecurrenttable.Sonetfarendlinecurrententry = append(sonetfarendlinecurrenttable.Sonetfarendlinecurrententry, child)
        return &sonetfarendlinecurrenttable.Sonetfarendlinecurrententry[len(sonetfarendlinecurrenttable.Sonetfarendlinecurrententry)-1]
    }
    return nil
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetfarendlinecurrenttable.Sonetfarendlinecurrententry {
        children[sonetfarendlinecurrenttable.Sonetfarendlinecurrententry[i].GetSegmentPath()] = &sonetfarendlinecurrenttable.Sonetfarendlinecurrententry[i]
    }
    return children
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetYangName() string { return "sonetFarEndLineCurrentTable" }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) SetParent(parent types.Entity) { sonetfarendlinecurrenttable.parent = parent }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetParent() types.Entity { return sonetfarendlinecurrenttable.parent }

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry
// An entry in the SONET/SDH Far End Line Current table.
type SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH interface in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    Sonetfarendlinecurrentess interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Medium/Section/Line interface in the current 15
    // minute interval. The type is interface{} with range: 0..4294967295.
    Sonetfarendlinecurrentsess interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH
    // Medium/Section/Line interface in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    Sonetfarendlinecurrentcvs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Medium/Section/Line interface in the current 15
    // minute interval. The type is interface{} with range: 0..4294967295.
    Sonetfarendlinecurrentuass interface{}
}

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetFilter() yfilter.YFilter { return sonetfarendlinecurrententry.YFilter }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) SetFilter(yf yfilter.YFilter) { sonetfarendlinecurrententry.YFilter = yf }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetFarEndLineCurrentESs" { return "Sonetfarendlinecurrentess" }
    if yname == "sonetFarEndLineCurrentSESs" { return "Sonetfarendlinecurrentsess" }
    if yname == "sonetFarEndLineCurrentCVs" { return "Sonetfarendlinecurrentcvs" }
    if yname == "sonetFarEndLineCurrentUASs" { return "Sonetfarendlinecurrentuass" }
    return ""
}

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetSegmentPath() string {
    return "sonetFarEndLineCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendlinecurrententry.Ifindex) + "']"
}

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetfarendlinecurrententry.Ifindex
    leafs["sonetFarEndLineCurrentESs"] = sonetfarendlinecurrententry.Sonetfarendlinecurrentess
    leafs["sonetFarEndLineCurrentSESs"] = sonetfarendlinecurrententry.Sonetfarendlinecurrentsess
    leafs["sonetFarEndLineCurrentCVs"] = sonetfarendlinecurrententry.Sonetfarendlinecurrentcvs
    leafs["sonetFarEndLineCurrentUASs"] = sonetfarendlinecurrententry.Sonetfarendlinecurrentuass
    return leafs
}

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetYangName() string { return "sonetFarEndLineCurrentEntry" }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) SetParent(parent types.Entity) { sonetfarendlinecurrententry.parent = parent }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetParent() types.Entity { return sonetfarendlinecurrententry.parent }

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetParentYangName() string { return "sonetFarEndLineCurrentTable" }

// SONETMIB_Sonetfarendlineintervaltable
// The SONET/SDH Far End Line Interval table.
type SONETMIB_Sonetfarendlineintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Interval table. The type is slice of
    // SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry.
    Sonetfarendlineintervalentry []SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetFilter() yfilter.YFilter { return sonetfarendlineintervaltable.YFilter }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) SetFilter(yf yfilter.YFilter) { sonetfarendlineintervaltable.YFilter = yf }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetGoName(yname string) string {
    if yname == "sonetFarEndLineIntervalEntry" { return "Sonetfarendlineintervalentry" }
    return ""
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetSegmentPath() string {
    return "sonetFarEndLineIntervalTable"
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetFarEndLineIntervalEntry" {
        for _, c := range sonetfarendlineintervaltable.Sonetfarendlineintervalentry {
            if sonetfarendlineintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry{}
        sonetfarendlineintervaltable.Sonetfarendlineintervalentry = append(sonetfarendlineintervaltable.Sonetfarendlineintervalentry, child)
        return &sonetfarendlineintervaltable.Sonetfarendlineintervalentry[len(sonetfarendlineintervaltable.Sonetfarendlineintervalentry)-1]
    }
    return nil
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetfarendlineintervaltable.Sonetfarendlineintervalentry {
        children[sonetfarendlineintervaltable.Sonetfarendlineintervalentry[i].GetSegmentPath()] = &sonetfarendlineintervaltable.Sonetfarendlineintervalentry[i]
    }
    return children
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetYangName() string { return "sonetFarEndLineIntervalTable" }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) SetParent(parent types.Entity) { sonetfarendlineintervaltable.parent = parent }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetParent() types.Entity { return sonetfarendlineintervaltable.parent }

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry
// An entry in the SONET/SDH Far
// End Line Interval table.
type SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetfarendlineintervalnumber interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH Line interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    Sonetfarendlineintervaless interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Line interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    Sonetfarendlineintervalsess interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH Line
    // interface in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    Sonetfarendlineintervalcvs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Line interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    Sonetfarendlineintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetfarendlineintervalvaliddata interface{}
}

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetFilter() yfilter.YFilter { return sonetfarendlineintervalentry.YFilter }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) SetFilter(yf yfilter.YFilter) { sonetfarendlineintervalentry.YFilter = yf }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetFarEndLineIntervalNumber" { return "Sonetfarendlineintervalnumber" }
    if yname == "sonetFarEndLineIntervalESs" { return "Sonetfarendlineintervaless" }
    if yname == "sonetFarEndLineIntervalSESs" { return "Sonetfarendlineintervalsess" }
    if yname == "sonetFarEndLineIntervalCVs" { return "Sonetfarendlineintervalcvs" }
    if yname == "sonetFarEndLineIntervalUASs" { return "Sonetfarendlineintervaluass" }
    if yname == "sonetFarEndLineIntervalValidData" { return "Sonetfarendlineintervalvaliddata" }
    return ""
}

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetSegmentPath() string {
    return "sonetFarEndLineIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendlineintervalentry.Ifindex) + "']" + "[sonetFarEndLineIntervalNumber='" + fmt.Sprintf("%v", sonetfarendlineintervalentry.Sonetfarendlineintervalnumber) + "']"
}

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetfarendlineintervalentry.Ifindex
    leafs["sonetFarEndLineIntervalNumber"] = sonetfarendlineintervalentry.Sonetfarendlineintervalnumber
    leafs["sonetFarEndLineIntervalESs"] = sonetfarendlineintervalentry.Sonetfarendlineintervaless
    leafs["sonetFarEndLineIntervalSESs"] = sonetfarendlineintervalentry.Sonetfarendlineintervalsess
    leafs["sonetFarEndLineIntervalCVs"] = sonetfarendlineintervalentry.Sonetfarendlineintervalcvs
    leafs["sonetFarEndLineIntervalUASs"] = sonetfarendlineintervalentry.Sonetfarendlineintervaluass
    leafs["sonetFarEndLineIntervalValidData"] = sonetfarendlineintervalentry.Sonetfarendlineintervalvaliddata
    return leafs
}

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetYangName() string { return "sonetFarEndLineIntervalEntry" }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) SetParent(parent types.Entity) { sonetfarendlineintervalentry.parent = parent }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetParent() types.Entity { return sonetfarendlineintervalentry.parent }

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetParentYangName() string { return "sonetFarEndLineIntervalTable" }

// SONETMIB_Sonetpathcurrenttable
// The SONET/SDH Path Current table.
type SONETMIB_Sonetpathcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Current table. The type is slice of
    // SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry.
    Sonetpathcurrententry []SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetFilter() yfilter.YFilter { return sonetpathcurrenttable.YFilter }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) SetFilter(yf yfilter.YFilter) { sonetpathcurrenttable.YFilter = yf }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetGoName(yname string) string {
    if yname == "sonetPathCurrentEntry" { return "Sonetpathcurrententry" }
    return ""
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetSegmentPath() string {
    return "sonetPathCurrentTable"
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetPathCurrentEntry" {
        for _, c := range sonetpathcurrenttable.Sonetpathcurrententry {
            if sonetpathcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry{}
        sonetpathcurrenttable.Sonetpathcurrententry = append(sonetpathcurrenttable.Sonetpathcurrententry, child)
        return &sonetpathcurrenttable.Sonetpathcurrententry[len(sonetpathcurrenttable.Sonetpathcurrententry)-1]
    }
    return nil
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetpathcurrenttable.Sonetpathcurrententry {
        children[sonetpathcurrenttable.Sonetpathcurrententry[i].GetSegmentPath()] = &sonetpathcurrenttable.Sonetpathcurrententry[i]
    }
    return children
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetYangName() string { return "sonetPathCurrentTable" }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) SetParent(parent types.Entity) { sonetpathcurrenttable.parent = parent }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetParent() types.Entity { return sonetpathcurrenttable.parent }

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry
// An entry in the SONET/SDH Path Current table.
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // A value that indicates the type of the SONET/SDH Path.  For SONET, the
    // assigned types are the STS-Nc SPEs, where N = 1, 3, 12, 24, 48, 192 and
    // 768. STS-1 is equal to 51.84 Mbps.  For SDH, the assigned types are the
    // STM-Nc VCs, where N = 1, 4, 16, 64 and 256. The type is
    // Sonetpathcurrentwidth.
    Sonetpathcurrentwidth interface{}

    // This variable indicates the status of the interface. The
    // sonetPathCurrentStatus is a bit map represented as a sum, therefore, it can
    // represent multiple defects simultaneously. The sonetPathNoDefect should be
    // set if and only if no other flag is set.  The various bit positions are:   
    // 1   sonetPathNoDefect    2   sonetPathSTSLOP    4   sonetPathSTSAIS    8  
    // sonetPathSTSRDI   16   sonetPathUnequipped   32  
    // sonetPathSignalLabelMismatch. The type is interface{} with range: 1..62.
    Sonetpathcurrentstatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Path in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    Sonetpathcurrentess interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Path in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    Sonetpathcurrentsess interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Path in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    Sonetpathcurrentcvs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a Path in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    Sonetpathcurrentuass interface{}

    // Specifies the payload carried by the SONET/SDH Path. The payload
    // specification corresponds to C2 (Signal Label) overhead byte in SONET/SDH
    // Path Overhead: unequipped(1)    : Path is not provisioned to carry any
    // payload. unspecified(2)   : Path is carrying an unspecifed payload. ds3(3) 
    // : Path is carrying a DS3 path as payload. vt15vc11(4)      : Path is
    // carrying SONET-VT1.5/SDH-VC11 payload. vt2vc12(5)       : Path is carrying
    // SONET-VT2/SDH-VC12 as payload. atmCell(6)       : Path is carrying ATM
    // Cells as payload. hdlcFr(7)        : Path is carrying Frame Relay (HDLC)
    // payload. e3(8)            : Path is carrying an E3 path as payload.
    // vtStructured(9)  : Path is carrying VTGs/TUG3s/TUG2s which may             
    // each carry a different payload.  A write operation on this object will
    // result in update to C2 overhead byte in the Path Overhead. The type is
    // Cspsonetpathpayload.
    Cspsonetpathpayload interface{}

    // This object represents the VT/VC mapping type. asynchronous: In this mode,
    // the channel structure of                DS1/E1 is neither visible nor
    // preserved.  byteSynchronous: In this mode, the DS0 signals inside          
    // the VT/VC can be found and extracted                   from the frame. The
    // initial value is asynchronous(1). The type is Csptributarymappingtype.
    Csptributarymappingtype interface{}

    // This object represents the mode used to transport DS0  Signalling
    // information for T1 byteSynchronous mapping (GR253). In
    // signallingTransferMode(2), the robbed-bit signalling  is transferred to the
    // VT header. In clearMode(3), only  the framing bit is transferred to the VT
    // header.           The initial value is signallingTransferMode(2)  if
    // csTributaryMappingType is byteSynchronous.  For asynchronous mapping, it is
    // notApplicable(1).  The value notApplicable(1) can not be set. The type is
    // Cspsignallingtransportmode.
    Cspsignallingtransportmode interface{}

    // This object represents the method used to group VCs into an STM-1 signal.
    // Applicable only to SDH.  au3Grouping: STM1<-AU-3<-TUG-2<-TU-12<-VC12 or    
    // STM1<-AU-3<-TUG-2<-TU-11<-VC11.  au4Grouping:
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-12<-VC12 or             
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-11<-VC11.  The initial value is au3Grouping(2)
    // for SDH and  notApplicable(1) for SONET. The type is
    // Csptributarygroupingtype.
    Csptributarygroupingtype interface{}
}

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetFilter() yfilter.YFilter { return sonetpathcurrententry.YFilter }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) SetFilter(yf yfilter.YFilter) { sonetpathcurrententry.YFilter = yf }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetPathCurrentWidth" { return "Sonetpathcurrentwidth" }
    if yname == "sonetPathCurrentStatus" { return "Sonetpathcurrentstatus" }
    if yname == "sonetPathCurrentESs" { return "Sonetpathcurrentess" }
    if yname == "sonetPathCurrentSESs" { return "Sonetpathcurrentsess" }
    if yname == "sonetPathCurrentCVs" { return "Sonetpathcurrentcvs" }
    if yname == "sonetPathCurrentUASs" { return "Sonetpathcurrentuass" }
    if yname == "cspSonetPathPayload" { return "Cspsonetpathpayload" }
    if yname == "cspTributaryMappingType" { return "Csptributarymappingtype" }
    if yname == "cspSignallingTransportMode" { return "Cspsignallingtransportmode" }
    if yname == "cspTributaryGroupingType" { return "Csptributarygroupingtype" }
    return ""
}

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetSegmentPath() string {
    return "sonetPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetpathcurrententry.Ifindex) + "']"
}

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetpathcurrententry.Ifindex
    leafs["sonetPathCurrentWidth"] = sonetpathcurrententry.Sonetpathcurrentwidth
    leafs["sonetPathCurrentStatus"] = sonetpathcurrententry.Sonetpathcurrentstatus
    leafs["sonetPathCurrentESs"] = sonetpathcurrententry.Sonetpathcurrentess
    leafs["sonetPathCurrentSESs"] = sonetpathcurrententry.Sonetpathcurrentsess
    leafs["sonetPathCurrentCVs"] = sonetpathcurrententry.Sonetpathcurrentcvs
    leafs["sonetPathCurrentUASs"] = sonetpathcurrententry.Sonetpathcurrentuass
    leafs["cspSonetPathPayload"] = sonetpathcurrententry.Cspsonetpathpayload
    leafs["cspTributaryMappingType"] = sonetpathcurrententry.Csptributarymappingtype
    leafs["cspSignallingTransportMode"] = sonetpathcurrententry.Cspsignallingtransportmode
    leafs["cspTributaryGroupingType"] = sonetpathcurrententry.Csptributarygroupingtype
    return leafs
}

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetYangName() string { return "sonetPathCurrentEntry" }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) SetParent(parent types.Entity) { sonetpathcurrententry.parent = parent }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetParent() types.Entity { return sonetpathcurrententry.parent }

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetParentYangName() string { return "sonetPathCurrentTable" }

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode represents The value notApplicable(1) can not be set.
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode string

const (
    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode_notApplicable SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode = "notApplicable"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode_signallingTransferMode SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode = "signallingTransferMode"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode_clearMode SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsignallingtransportmode = "clearMode"
)

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload represents C2 overhead byte in the Path Overhead.
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload string

const (
    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_unequipped SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "unequipped"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_unspecified SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "unspecified"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_ds3 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "ds3"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_vt15vc11 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "vt15vc11"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_vt2vc12 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "vt2vc12"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_atmCell SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "atmCell"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_hdlcFr SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "hdlcFr"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_e3 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "e3"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload_vtStructured SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Cspsonetpathpayload = "vtStructured"
)

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype represents notApplicable(1) for SONET.
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype string

const (
    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype_notApplicable SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype = "notApplicable"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype_au3Grouping SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype = "au3Grouping"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype_au4Grouping SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarygroupingtype = "au4Grouping"
)

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarymappingtype represents The initial value is asynchronous(1).
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarymappingtype string

const (
    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarymappingtype_asynchronous SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarymappingtype = "asynchronous"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarymappingtype_byteSynchronous SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Csptributarymappingtype = "byteSynchronous"
)

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth represents types are the STM-Nc VCs, where N = 1, 4, 16, 64 and 256.
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth string

const (
    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts1 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts1"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts3cSTM1 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts3cSTM1"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts12cSTM4 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts12cSTM4"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts24c SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts24c"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts48cSTM16 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts48cSTM16"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts192cSTM64 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts192cSTM64"

    SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth_sts768cSTM256 SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry_Sonetpathcurrentwidth = "sts768cSTM256"
)

// SONETMIB_Sonetpathintervaltable
// The SONET/SDH Path Interval table.
type SONETMIB_Sonetpathintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Interval table. The type is slice of
    // SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry.
    Sonetpathintervalentry []SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetFilter() yfilter.YFilter { return sonetpathintervaltable.YFilter }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) SetFilter(yf yfilter.YFilter) { sonetpathintervaltable.YFilter = yf }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetGoName(yname string) string {
    if yname == "sonetPathIntervalEntry" { return "Sonetpathintervalentry" }
    return ""
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetSegmentPath() string {
    return "sonetPathIntervalTable"
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetPathIntervalEntry" {
        for _, c := range sonetpathintervaltable.Sonetpathintervalentry {
            if sonetpathintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry{}
        sonetpathintervaltable.Sonetpathintervalentry = append(sonetpathintervaltable.Sonetpathintervalentry, child)
        return &sonetpathintervaltable.Sonetpathintervalentry[len(sonetpathintervaltable.Sonetpathintervalentry)-1]
    }
    return nil
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetpathintervaltable.Sonetpathintervalentry {
        children[sonetpathintervaltable.Sonetpathintervalentry[i].GetSegmentPath()] = &sonetpathintervaltable.Sonetpathintervalentry[i]
    }
    return children
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetYangName() string { return "sonetPathIntervalTable" }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) SetParent(parent types.Entity) { sonetpathintervaltable.parent = parent }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetParent() types.Entity { return sonetpathintervaltable.parent }

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry
// An entry in the SONET/SDH Path Interval table.
type SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetpathintervalnumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Path in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    Sonetpathintervaless interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Path in a particular 15-minute interval in the
    // past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetpathintervalsess interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Path in a particular 15-minute interval in the past 24 hours.
    // The type is interface{} with range: 0..4294967295.
    Sonetpathintervalcvs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a Path in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    Sonetpathintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetpathintervalvaliddata interface{}
}

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetFilter() yfilter.YFilter { return sonetpathintervalentry.YFilter }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) SetFilter(yf yfilter.YFilter) { sonetpathintervalentry.YFilter = yf }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetPathIntervalNumber" { return "Sonetpathintervalnumber" }
    if yname == "sonetPathIntervalESs" { return "Sonetpathintervaless" }
    if yname == "sonetPathIntervalSESs" { return "Sonetpathintervalsess" }
    if yname == "sonetPathIntervalCVs" { return "Sonetpathintervalcvs" }
    if yname == "sonetPathIntervalUASs" { return "Sonetpathintervaluass" }
    if yname == "sonetPathIntervalValidData" { return "Sonetpathintervalvaliddata" }
    return ""
}

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetSegmentPath() string {
    return "sonetPathIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetpathintervalentry.Ifindex) + "']" + "[sonetPathIntervalNumber='" + fmt.Sprintf("%v", sonetpathintervalentry.Sonetpathintervalnumber) + "']"
}

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetpathintervalentry.Ifindex
    leafs["sonetPathIntervalNumber"] = sonetpathintervalentry.Sonetpathintervalnumber
    leafs["sonetPathIntervalESs"] = sonetpathintervalentry.Sonetpathintervaless
    leafs["sonetPathIntervalSESs"] = sonetpathintervalentry.Sonetpathintervalsess
    leafs["sonetPathIntervalCVs"] = sonetpathintervalentry.Sonetpathintervalcvs
    leafs["sonetPathIntervalUASs"] = sonetpathintervalentry.Sonetpathintervaluass
    leafs["sonetPathIntervalValidData"] = sonetpathintervalentry.Sonetpathintervalvaliddata
    return leafs
}

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetYangName() string { return "sonetPathIntervalEntry" }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) SetParent(parent types.Entity) { sonetpathintervalentry.parent = parent }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetParent() types.Entity { return sonetpathintervalentry.parent }

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetParentYangName() string { return "sonetPathIntervalTable" }

// SONETMIB_Sonetfarendpathcurrenttable
// The SONET/SDH Far End Path Current table.
type SONETMIB_Sonetfarendpathcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Current table. The type is slice of
    // SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry.
    Sonetfarendpathcurrententry []SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetFilter() yfilter.YFilter { return sonetfarendpathcurrenttable.YFilter }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) SetFilter(yf yfilter.YFilter) { sonetfarendpathcurrenttable.YFilter = yf }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetGoName(yname string) string {
    if yname == "sonetFarEndPathCurrentEntry" { return "Sonetfarendpathcurrententry" }
    return ""
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetSegmentPath() string {
    return "sonetFarEndPathCurrentTable"
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetFarEndPathCurrentEntry" {
        for _, c := range sonetfarendpathcurrenttable.Sonetfarendpathcurrententry {
            if sonetfarendpathcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry{}
        sonetfarendpathcurrenttable.Sonetfarendpathcurrententry = append(sonetfarendpathcurrenttable.Sonetfarendpathcurrententry, child)
        return &sonetfarendpathcurrenttable.Sonetfarendpathcurrententry[len(sonetfarendpathcurrenttable.Sonetfarendpathcurrententry)-1]
    }
    return nil
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetfarendpathcurrenttable.Sonetfarendpathcurrententry {
        children[sonetfarendpathcurrenttable.Sonetfarendpathcurrententry[i].GetSegmentPath()] = &sonetfarendpathcurrenttable.Sonetfarendpathcurrententry[i]
    }
    return children
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetYangName() string { return "sonetFarEndPathCurrentTable" }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) SetParent(parent types.Entity) { sonetfarendpathcurrenttable.parent = parent }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetParent() types.Entity { return sonetfarendpathcurrenttable.parent }

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry
// An entry in the SONET/SDH Far End Path Current table.
type SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH interface in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    Sonetfarendpathcurrentess interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Path interface in the current 15 minute
    // interval. The type is interface{} with range: 0..4294967295.
    Sonetfarendpathcurrentsess interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH Path
    // interface in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    Sonetfarendpathcurrentcvs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Path interface in the current 15 minute
    // interval. The type is interface{} with range: 0..4294967295.
    Sonetfarendpathcurrentuass interface{}
}

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetFilter() yfilter.YFilter { return sonetfarendpathcurrententry.YFilter }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) SetFilter(yf yfilter.YFilter) { sonetfarendpathcurrententry.YFilter = yf }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetFarEndPathCurrentESs" { return "Sonetfarendpathcurrentess" }
    if yname == "sonetFarEndPathCurrentSESs" { return "Sonetfarendpathcurrentsess" }
    if yname == "sonetFarEndPathCurrentCVs" { return "Sonetfarendpathcurrentcvs" }
    if yname == "sonetFarEndPathCurrentUASs" { return "Sonetfarendpathcurrentuass" }
    return ""
}

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetSegmentPath() string {
    return "sonetFarEndPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendpathcurrententry.Ifindex) + "']"
}

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetfarendpathcurrententry.Ifindex
    leafs["sonetFarEndPathCurrentESs"] = sonetfarendpathcurrententry.Sonetfarendpathcurrentess
    leafs["sonetFarEndPathCurrentSESs"] = sonetfarendpathcurrententry.Sonetfarendpathcurrentsess
    leafs["sonetFarEndPathCurrentCVs"] = sonetfarendpathcurrententry.Sonetfarendpathcurrentcvs
    leafs["sonetFarEndPathCurrentUASs"] = sonetfarendpathcurrententry.Sonetfarendpathcurrentuass
    return leafs
}

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetYangName() string { return "sonetFarEndPathCurrentEntry" }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) SetParent(parent types.Entity) { sonetfarendpathcurrententry.parent = parent }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetParent() types.Entity { return sonetfarendpathcurrententry.parent }

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetParentYangName() string { return "sonetFarEndPathCurrentTable" }

// SONETMIB_Sonetfarendpathintervaltable
// The SONET/SDH Far End Path Interval table.
type SONETMIB_Sonetfarendpathintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Interval table. The type is slice of
    // SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry.
    Sonetfarendpathintervalentry []SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetFilter() yfilter.YFilter { return sonetfarendpathintervaltable.YFilter }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) SetFilter(yf yfilter.YFilter) { sonetfarendpathintervaltable.YFilter = yf }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetGoName(yname string) string {
    if yname == "sonetFarEndPathIntervalEntry" { return "Sonetfarendpathintervalentry" }
    return ""
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetSegmentPath() string {
    return "sonetFarEndPathIntervalTable"
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetFarEndPathIntervalEntry" {
        for _, c := range sonetfarendpathintervaltable.Sonetfarendpathintervalentry {
            if sonetfarendpathintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry{}
        sonetfarendpathintervaltable.Sonetfarendpathintervalentry = append(sonetfarendpathintervaltable.Sonetfarendpathintervalentry, child)
        return &sonetfarendpathintervaltable.Sonetfarendpathintervalentry[len(sonetfarendpathintervaltable.Sonetfarendpathintervalentry)-1]
    }
    return nil
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetfarendpathintervaltable.Sonetfarendpathintervalentry {
        children[sonetfarendpathintervaltable.Sonetfarendpathintervalentry[i].GetSegmentPath()] = &sonetfarendpathintervaltable.Sonetfarendpathintervalentry[i]
    }
    return children
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetYangName() string { return "sonetFarEndPathIntervalTable" }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) SetParent(parent types.Entity) { sonetfarendpathintervaltable.parent = parent }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetParent() types.Entity { return sonetfarendpathintervaltable.parent }

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry
// An entry in the SONET/SDH Far
// End Path Interval table.
type SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetfarendpathintervalnumber interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH Path interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    Sonetfarendpathintervaless interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Path interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    Sonetfarendpathintervalsess interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH Path
    // interface in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    Sonetfarendpathintervalcvs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Path interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    Sonetfarendpathintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetfarendpathintervalvaliddata interface{}
}

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetFilter() yfilter.YFilter { return sonetfarendpathintervalentry.YFilter }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) SetFilter(yf yfilter.YFilter) { sonetfarendpathintervalentry.YFilter = yf }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetFarEndPathIntervalNumber" { return "Sonetfarendpathintervalnumber" }
    if yname == "sonetFarEndPathIntervalESs" { return "Sonetfarendpathintervaless" }
    if yname == "sonetFarEndPathIntervalSESs" { return "Sonetfarendpathintervalsess" }
    if yname == "sonetFarEndPathIntervalCVs" { return "Sonetfarendpathintervalcvs" }
    if yname == "sonetFarEndPathIntervalUASs" { return "Sonetfarendpathintervaluass" }
    if yname == "sonetFarEndPathIntervalValidData" { return "Sonetfarendpathintervalvaliddata" }
    return ""
}

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetSegmentPath() string {
    return "sonetFarEndPathIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendpathintervalentry.Ifindex) + "']" + "[sonetFarEndPathIntervalNumber='" + fmt.Sprintf("%v", sonetfarendpathintervalentry.Sonetfarendpathintervalnumber) + "']"
}

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetfarendpathintervalentry.Ifindex
    leafs["sonetFarEndPathIntervalNumber"] = sonetfarendpathintervalentry.Sonetfarendpathintervalnumber
    leafs["sonetFarEndPathIntervalESs"] = sonetfarendpathintervalentry.Sonetfarendpathintervaless
    leafs["sonetFarEndPathIntervalSESs"] = sonetfarendpathintervalentry.Sonetfarendpathintervalsess
    leafs["sonetFarEndPathIntervalCVs"] = sonetfarendpathintervalentry.Sonetfarendpathintervalcvs
    leafs["sonetFarEndPathIntervalUASs"] = sonetfarendpathintervalentry.Sonetfarendpathintervaluass
    leafs["sonetFarEndPathIntervalValidData"] = sonetfarendpathintervalentry.Sonetfarendpathintervalvaliddata
    return leafs
}

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetYangName() string { return "sonetFarEndPathIntervalEntry" }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) SetParent(parent types.Entity) { sonetfarendpathintervalentry.parent = parent }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetParent() types.Entity { return sonetfarendpathintervalentry.parent }

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetParentYangName() string { return "sonetFarEndPathIntervalTable" }

// SONETMIB_Sonetvtcurrenttable
// The SONET/SDH VT Current table.
type SONETMIB_Sonetvtcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH VT Current table. The type is slice of
    // SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry.
    Sonetvtcurrententry []SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetFilter() yfilter.YFilter { return sonetvtcurrenttable.YFilter }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) SetFilter(yf yfilter.YFilter) { sonetvtcurrenttable.YFilter = yf }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetGoName(yname string) string {
    if yname == "sonetVTCurrentEntry" { return "Sonetvtcurrententry" }
    return ""
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetSegmentPath() string {
    return "sonetVTCurrentTable"
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetVTCurrentEntry" {
        for _, c := range sonetvtcurrenttable.Sonetvtcurrententry {
            if sonetvtcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry{}
        sonetvtcurrenttable.Sonetvtcurrententry = append(sonetvtcurrenttable.Sonetvtcurrententry, child)
        return &sonetvtcurrenttable.Sonetvtcurrententry[len(sonetvtcurrenttable.Sonetvtcurrententry)-1]
    }
    return nil
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetvtcurrenttable.Sonetvtcurrententry {
        children[sonetvtcurrenttable.Sonetvtcurrententry[i].GetSegmentPath()] = &sonetvtcurrenttable.Sonetvtcurrententry[i]
    }
    return children
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetYangName() string { return "sonetVTCurrentTable" }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) SetParent(parent types.Entity) { sonetvtcurrenttable.parent = parent }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetParent() types.Entity { return sonetvtcurrenttable.parent }

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry
// An entry in the SONET/SDH VT Current table.
type SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // A value that indicates the type of the SONET VT and SDH VC.  Assigned
    // widths are VT1.5/VC11, VT2/VC12, VT3, VT6/VC2, and VT6c. The type is
    // Sonetvtcurrentwidth.
    Sonetvtcurrentwidth interface{}

    // This variable indicates the status of the interface. The
    // sonetVTCurrentStatus is a bit map represented as a sum, therefore, it can
    // represent multiple defects and failures simultaneously. The sonetVTNoDefect
    // should be set if and only if no other flag is set.  The various bit
    // positions are:    1   sonetVTNoDefect    2   sonetVTLOP    4  
    // sonetVTPathAIS    8   sonetVTPathRDI   16   sonetVTPathRFI   32  
    // sonetVTUnequipped   64   sonetVTSignalLabelMismatch. The type is
    // interface{} with range: 1..126.
    Sonetvtcurrentstatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH VT in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    Sonetvtcurrentess interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH VT in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    Sonetvtcurrentsess interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH VT in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    Sonetvtcurrentcvs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a VT in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    Sonetvtcurrentuass interface{}
}

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetFilter() yfilter.YFilter { return sonetvtcurrententry.YFilter }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) SetFilter(yf yfilter.YFilter) { sonetvtcurrententry.YFilter = yf }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetVTCurrentWidth" { return "Sonetvtcurrentwidth" }
    if yname == "sonetVTCurrentStatus" { return "Sonetvtcurrentstatus" }
    if yname == "sonetVTCurrentESs" { return "Sonetvtcurrentess" }
    if yname == "sonetVTCurrentSESs" { return "Sonetvtcurrentsess" }
    if yname == "sonetVTCurrentCVs" { return "Sonetvtcurrentcvs" }
    if yname == "sonetVTCurrentUASs" { return "Sonetvtcurrentuass" }
    return ""
}

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetSegmentPath() string {
    return "sonetVTCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetvtcurrententry.Ifindex) + "']"
}

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetvtcurrententry.Ifindex
    leafs["sonetVTCurrentWidth"] = sonetvtcurrententry.Sonetvtcurrentwidth
    leafs["sonetVTCurrentStatus"] = sonetvtcurrententry.Sonetvtcurrentstatus
    leafs["sonetVTCurrentESs"] = sonetvtcurrententry.Sonetvtcurrentess
    leafs["sonetVTCurrentSESs"] = sonetvtcurrententry.Sonetvtcurrentsess
    leafs["sonetVTCurrentCVs"] = sonetvtcurrententry.Sonetvtcurrentcvs
    leafs["sonetVTCurrentUASs"] = sonetvtcurrententry.Sonetvtcurrentuass
    return leafs
}

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetYangName() string { return "sonetVTCurrentEntry" }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) SetParent(parent types.Entity) { sonetvtcurrententry.parent = parent }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetParent() types.Entity { return sonetvtcurrententry.parent }

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetParentYangName() string { return "sonetVTCurrentTable" }

// SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth represents VT1.5/VC11, VT2/VC12, VT3, VT6/VC2, and VT6c.
type SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth string

const (
    SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth_vtWidth15VC11 SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth = "vtWidth15VC11"

    SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth_vtWidth2VC12 SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth = "vtWidth2VC12"

    SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth_vtWidth3 SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth = "vtWidth3"

    SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth_vtWidth6VC2 SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth = "vtWidth6VC2"

    SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth_vtWidth6c SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry_Sonetvtcurrentwidth = "vtWidth6c"
)

// SONETMIB_Sonetvtintervaltable
// The SONET/SDH VT Interval table.
type SONETMIB_Sonetvtintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH VT Interval table. The type is slice of
    // SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry.
    Sonetvtintervalentry []SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetFilter() yfilter.YFilter { return sonetvtintervaltable.YFilter }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) SetFilter(yf yfilter.YFilter) { sonetvtintervaltable.YFilter = yf }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetGoName(yname string) string {
    if yname == "sonetVTIntervalEntry" { return "Sonetvtintervalentry" }
    return ""
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetSegmentPath() string {
    return "sonetVTIntervalTable"
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetVTIntervalEntry" {
        for _, c := range sonetvtintervaltable.Sonetvtintervalentry {
            if sonetvtintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry{}
        sonetvtintervaltable.Sonetvtintervalentry = append(sonetvtintervaltable.Sonetvtintervalentry, child)
        return &sonetvtintervaltable.Sonetvtintervalentry[len(sonetvtintervaltable.Sonetvtintervalentry)-1]
    }
    return nil
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetvtintervaltable.Sonetvtintervalentry {
        children[sonetvtintervaltable.Sonetvtintervalentry[i].GetSegmentPath()] = &sonetvtintervaltable.Sonetvtintervalentry[i]
    }
    return children
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetYangName() string { return "sonetVTIntervalTable" }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) SetParent(parent types.Entity) { sonetvtintervaltable.parent = parent }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetParent() types.Entity { return sonetvtintervaltable.parent }

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry
// An entry in the SONET/SDH VT Interval table.
type SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetvtintervalnumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH VT in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    Sonetvtintervaless interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH VT in a particular 15-minute interval in the
    // past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetvtintervalsess interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH VT in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    Sonetvtintervalcvs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a VT in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    Sonetvtintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetvtintervalvaliddata interface{}
}

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetFilter() yfilter.YFilter { return sonetvtintervalentry.YFilter }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) SetFilter(yf yfilter.YFilter) { sonetvtintervalentry.YFilter = yf }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetVTIntervalNumber" { return "Sonetvtintervalnumber" }
    if yname == "sonetVTIntervalESs" { return "Sonetvtintervaless" }
    if yname == "sonetVTIntervalSESs" { return "Sonetvtintervalsess" }
    if yname == "sonetVTIntervalCVs" { return "Sonetvtintervalcvs" }
    if yname == "sonetVTIntervalUASs" { return "Sonetvtintervaluass" }
    if yname == "sonetVTIntervalValidData" { return "Sonetvtintervalvaliddata" }
    return ""
}

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetSegmentPath() string {
    return "sonetVTIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetvtintervalentry.Ifindex) + "']" + "[sonetVTIntervalNumber='" + fmt.Sprintf("%v", sonetvtintervalentry.Sonetvtintervalnumber) + "']"
}

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetvtintervalentry.Ifindex
    leafs["sonetVTIntervalNumber"] = sonetvtintervalentry.Sonetvtintervalnumber
    leafs["sonetVTIntervalESs"] = sonetvtintervalentry.Sonetvtintervaless
    leafs["sonetVTIntervalSESs"] = sonetvtintervalentry.Sonetvtintervalsess
    leafs["sonetVTIntervalCVs"] = sonetvtintervalentry.Sonetvtintervalcvs
    leafs["sonetVTIntervalUASs"] = sonetvtintervalentry.Sonetvtintervaluass
    leafs["sonetVTIntervalValidData"] = sonetvtintervalentry.Sonetvtintervalvaliddata
    return leafs
}

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetYangName() string { return "sonetVTIntervalEntry" }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) SetParent(parent types.Entity) { sonetvtintervalentry.parent = parent }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetParent() types.Entity { return sonetvtintervalentry.parent }

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetParentYangName() string { return "sonetVTIntervalTable" }

// SONETMIB_Sonetfarendvtcurrenttable
// The SONET/SDH Far End VT Current table.
type SONETMIB_Sonetfarendvtcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End VT Current table. The type is slice of
    // SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry.
    Sonetfarendvtcurrententry []SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetFilter() yfilter.YFilter { return sonetfarendvtcurrenttable.YFilter }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) SetFilter(yf yfilter.YFilter) { sonetfarendvtcurrenttable.YFilter = yf }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetGoName(yname string) string {
    if yname == "sonetFarEndVTCurrentEntry" { return "Sonetfarendvtcurrententry" }
    return ""
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetSegmentPath() string {
    return "sonetFarEndVTCurrentTable"
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetFarEndVTCurrentEntry" {
        for _, c := range sonetfarendvtcurrenttable.Sonetfarendvtcurrententry {
            if sonetfarendvtcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry{}
        sonetfarendvtcurrenttable.Sonetfarendvtcurrententry = append(sonetfarendvtcurrenttable.Sonetfarendvtcurrententry, child)
        return &sonetfarendvtcurrenttable.Sonetfarendvtcurrententry[len(sonetfarendvtcurrenttable.Sonetfarendvtcurrententry)-1]
    }
    return nil
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetfarendvtcurrenttable.Sonetfarendvtcurrententry {
        children[sonetfarendvtcurrenttable.Sonetfarendvtcurrententry[i].GetSegmentPath()] = &sonetfarendvtcurrenttable.Sonetfarendvtcurrententry[i]
    }
    return children
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetYangName() string { return "sonetFarEndVTCurrentTable" }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) SetParent(parent types.Entity) { sonetfarendvtcurrenttable.parent = parent }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetParent() types.Entity { return sonetfarendvtcurrenttable.parent }

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry
// An entry in the SONET/SDH Far End VT Current table.
type SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH interface in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    Sonetfarendvtcurrentess interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH VT interface in the current 15 minute interval.
    // The type is interface{} with range: 0..4294967295.
    Sonetfarendvtcurrentsess interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH VT
    // interface in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    Sonetfarendvtcurrentcvs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH VT interface in the current 15 minute interval.
    // The type is interface{} with range: 0..4294967295.
    Sonetfarendvtcurrentuass interface{}
}

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetFilter() yfilter.YFilter { return sonetfarendvtcurrententry.YFilter }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) SetFilter(yf yfilter.YFilter) { sonetfarendvtcurrententry.YFilter = yf }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetFarEndVTCurrentESs" { return "Sonetfarendvtcurrentess" }
    if yname == "sonetFarEndVTCurrentSESs" { return "Sonetfarendvtcurrentsess" }
    if yname == "sonetFarEndVTCurrentCVs" { return "Sonetfarendvtcurrentcvs" }
    if yname == "sonetFarEndVTCurrentUASs" { return "Sonetfarendvtcurrentuass" }
    return ""
}

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetSegmentPath() string {
    return "sonetFarEndVTCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendvtcurrententry.Ifindex) + "']"
}

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetfarendvtcurrententry.Ifindex
    leafs["sonetFarEndVTCurrentESs"] = sonetfarendvtcurrententry.Sonetfarendvtcurrentess
    leafs["sonetFarEndVTCurrentSESs"] = sonetfarendvtcurrententry.Sonetfarendvtcurrentsess
    leafs["sonetFarEndVTCurrentCVs"] = sonetfarendvtcurrententry.Sonetfarendvtcurrentcvs
    leafs["sonetFarEndVTCurrentUASs"] = sonetfarendvtcurrententry.Sonetfarendvtcurrentuass
    return leafs
}

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetYangName() string { return "sonetFarEndVTCurrentEntry" }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) SetParent(parent types.Entity) { sonetfarendvtcurrententry.parent = parent }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetParent() types.Entity { return sonetfarendvtcurrententry.parent }

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetParentYangName() string { return "sonetFarEndVTCurrentTable" }

// SONETMIB_Sonetfarendvtintervaltable
// The SONET/SDH Far End VT Interval table.
type SONETMIB_Sonetfarendvtintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End VT Interval table. The type is slice of
    // SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry.
    Sonetfarendvtintervalentry []SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetFilter() yfilter.YFilter { return sonetfarendvtintervaltable.YFilter }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) SetFilter(yf yfilter.YFilter) { sonetfarendvtintervaltable.YFilter = yf }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetGoName(yname string) string {
    if yname == "sonetFarEndVTIntervalEntry" { return "Sonetfarendvtintervalentry" }
    return ""
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetSegmentPath() string {
    return "sonetFarEndVTIntervalTable"
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sonetFarEndVTIntervalEntry" {
        for _, c := range sonetfarendvtintervaltable.Sonetfarendvtintervalentry {
            if sonetfarendvtintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry{}
        sonetfarendvtintervaltable.Sonetfarendvtintervalentry = append(sonetfarendvtintervaltable.Sonetfarendvtintervalentry, child)
        return &sonetfarendvtintervaltable.Sonetfarendvtintervalentry[len(sonetfarendvtintervaltable.Sonetfarendvtintervalentry)-1]
    }
    return nil
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sonetfarendvtintervaltable.Sonetfarendvtintervalentry {
        children[sonetfarendvtintervaltable.Sonetfarendvtintervalentry[i].GetSegmentPath()] = &sonetfarendvtintervaltable.Sonetfarendvtintervalentry[i]
    }
    return children
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetYangName() string { return "sonetFarEndVTIntervalTable" }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) SetParent(parent types.Entity) { sonetfarendvtintervaltable.parent = parent }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetParent() types.Entity { return sonetfarendvtintervaltable.parent }

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetParentYangName() string { return "SONET-MIB" }

// SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry
// An entry in the SONET/SDH Far
// End VT Interval table.
type SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    Sonetfarendvtintervalnumber interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH VT interface in a particular 15-minute interval
    // in the past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetfarendvtintervaless interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH VT interface in a particular 15-minute interval
    // in the past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetfarendvtintervalsess interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH VT
    // interface in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    Sonetfarendvtintervalcvs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH VT interface in a particular 15-minute interval
    // in the past 24 hours. The type is interface{} with range: 0..4294967295.
    Sonetfarendvtintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Sonetfarendvtintervalvaliddata interface{}
}

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetFilter() yfilter.YFilter { return sonetfarendvtintervalentry.YFilter }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) SetFilter(yf yfilter.YFilter) { sonetfarendvtintervalentry.YFilter = yf }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "sonetFarEndVTIntervalNumber" { return "Sonetfarendvtintervalnumber" }
    if yname == "sonetFarEndVTIntervalESs" { return "Sonetfarendvtintervaless" }
    if yname == "sonetFarEndVTIntervalSESs" { return "Sonetfarendvtintervalsess" }
    if yname == "sonetFarEndVTIntervalCVs" { return "Sonetfarendvtintervalcvs" }
    if yname == "sonetFarEndVTIntervalUASs" { return "Sonetfarendvtintervaluass" }
    if yname == "sonetFarEndVTIntervalValidData" { return "Sonetfarendvtintervalvaliddata" }
    return ""
}

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetSegmentPath() string {
    return "sonetFarEndVTIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendvtintervalentry.Ifindex) + "']" + "[sonetFarEndVTIntervalNumber='" + fmt.Sprintf("%v", sonetfarendvtintervalentry.Sonetfarendvtintervalnumber) + "']"
}

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = sonetfarendvtintervalentry.Ifindex
    leafs["sonetFarEndVTIntervalNumber"] = sonetfarendvtintervalentry.Sonetfarendvtintervalnumber
    leafs["sonetFarEndVTIntervalESs"] = sonetfarendvtintervalentry.Sonetfarendvtintervaless
    leafs["sonetFarEndVTIntervalSESs"] = sonetfarendvtintervalentry.Sonetfarendvtintervalsess
    leafs["sonetFarEndVTIntervalCVs"] = sonetfarendvtintervalentry.Sonetfarendvtintervalcvs
    leafs["sonetFarEndVTIntervalUASs"] = sonetfarendvtintervalentry.Sonetfarendvtintervaluass
    leafs["sonetFarEndVTIntervalValidData"] = sonetfarendvtintervalentry.Sonetfarendvtintervalvaliddata
    return leafs
}

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetYangName() string { return "sonetFarEndVTIntervalEntry" }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) SetParent(parent types.Entity) { sonetfarendvtintervalentry.parent = parent }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetParent() types.Entity { return sonetfarendvtintervalentry.parent }

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetParentYangName() string { return "sonetFarEndVTIntervalTable" }

