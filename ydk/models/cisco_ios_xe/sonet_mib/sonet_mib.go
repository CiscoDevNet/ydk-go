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
    EntityData types.CommonEntityData
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

func (sONETMIB *SONETMIB) GetEntityData() *types.CommonEntityData {
    sONETMIB.EntityData.YFilter = sONETMIB.YFilter
    sONETMIB.EntityData.YangName = "SONET-MIB"
    sONETMIB.EntityData.BundleName = "cisco_ios_xe"
    sONETMIB.EntityData.ParentYangName = "SONET-MIB"
    sONETMIB.EntityData.SegmentPath = "SONET-MIB:SONET-MIB"
    sONETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sONETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sONETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sONETMIB.EntityData.Children = make(map[string]types.YChild)
    sONETMIB.EntityData.Children["sonetMedium"] = types.YChild{"Sonetmedium", &sONETMIB.Sonetmedium}
    sONETMIB.EntityData.Children["sonetMediumTable"] = types.YChild{"Sonetmediumtable", &sONETMIB.Sonetmediumtable}
    sONETMIB.EntityData.Children["sonetSectionCurrentTable"] = types.YChild{"Sonetsectioncurrenttable", &sONETMIB.Sonetsectioncurrenttable}
    sONETMIB.EntityData.Children["sonetSectionIntervalTable"] = types.YChild{"Sonetsectionintervaltable", &sONETMIB.Sonetsectionintervaltable}
    sONETMIB.EntityData.Children["sonetLineCurrentTable"] = types.YChild{"Sonetlinecurrenttable", &sONETMIB.Sonetlinecurrenttable}
    sONETMIB.EntityData.Children["sonetLineIntervalTable"] = types.YChild{"Sonetlineintervaltable", &sONETMIB.Sonetlineintervaltable}
    sONETMIB.EntityData.Children["sonetFarEndLineCurrentTable"] = types.YChild{"Sonetfarendlinecurrenttable", &sONETMIB.Sonetfarendlinecurrenttable}
    sONETMIB.EntityData.Children["sonetFarEndLineIntervalTable"] = types.YChild{"Sonetfarendlineintervaltable", &sONETMIB.Sonetfarendlineintervaltable}
    sONETMIB.EntityData.Children["sonetPathCurrentTable"] = types.YChild{"Sonetpathcurrenttable", &sONETMIB.Sonetpathcurrenttable}
    sONETMIB.EntityData.Children["sonetPathIntervalTable"] = types.YChild{"Sonetpathintervaltable", &sONETMIB.Sonetpathintervaltable}
    sONETMIB.EntityData.Children["sonetFarEndPathCurrentTable"] = types.YChild{"Sonetfarendpathcurrenttable", &sONETMIB.Sonetfarendpathcurrenttable}
    sONETMIB.EntityData.Children["sonetFarEndPathIntervalTable"] = types.YChild{"Sonetfarendpathintervaltable", &sONETMIB.Sonetfarendpathintervaltable}
    sONETMIB.EntityData.Children["sonetVTCurrentTable"] = types.YChild{"Sonetvtcurrenttable", &sONETMIB.Sonetvtcurrenttable}
    sONETMIB.EntityData.Children["sonetVTIntervalTable"] = types.YChild{"Sonetvtintervaltable", &sONETMIB.Sonetvtintervaltable}
    sONETMIB.EntityData.Children["sonetFarEndVTCurrentTable"] = types.YChild{"Sonetfarendvtcurrenttable", &sONETMIB.Sonetfarendvtcurrenttable}
    sONETMIB.EntityData.Children["sonetFarEndVTIntervalTable"] = types.YChild{"Sonetfarendvtintervaltable", &sONETMIB.Sonetfarendvtintervaltable}
    sONETMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sONETMIB.EntityData)
}

// SONETMIB_Sonetmedium
type SONETMIB_Sonetmedium struct {
    EntityData types.CommonEntityData
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

func (sonetmedium *SONETMIB_Sonetmedium) GetEntityData() *types.CommonEntityData {
    sonetmedium.EntityData.YFilter = sonetmedium.YFilter
    sonetmedium.EntityData.YangName = "sonetMedium"
    sonetmedium.EntityData.BundleName = "cisco_ios_xe"
    sonetmedium.EntityData.ParentYangName = "SONET-MIB"
    sonetmedium.EntityData.SegmentPath = "sonetMedium"
    sonetmedium.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetmedium.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetmedium.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetmedium.EntityData.Children = make(map[string]types.YChild)
    sonetmedium.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetmedium.EntityData.Leafs["sonetSESthresholdSet"] = types.YLeaf{"Sonetsesthresholdset", sonetmedium.Sonetsesthresholdset}
    return &(sonetmedium.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Medium table. The type is slice of
    // SONETMIB_Sonetmediumtable_Sonetmediumentry.
    Sonetmediumentry []SONETMIB_Sonetmediumtable_Sonetmediumentry
}

func (sonetmediumtable *SONETMIB_Sonetmediumtable) GetEntityData() *types.CommonEntityData {
    sonetmediumtable.EntityData.YFilter = sonetmediumtable.YFilter
    sonetmediumtable.EntityData.YangName = "sonetMediumTable"
    sonetmediumtable.EntityData.BundleName = "cisco_ios_xe"
    sonetmediumtable.EntityData.ParentYangName = "SONET-MIB"
    sonetmediumtable.EntityData.SegmentPath = "sonetMediumTable"
    sonetmediumtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetmediumtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetmediumtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetmediumtable.EntityData.Children = make(map[string]types.YChild)
    sonetmediumtable.EntityData.Children["sonetMediumEntry"] = types.YChild{"Sonetmediumentry", nil}
    for i := range sonetmediumtable.Sonetmediumentry {
        sonetmediumtable.EntityData.Children[types.GetSegmentPath(&sonetmediumtable.Sonetmediumentry[i])] = types.YChild{"Sonetmediumentry", &sonetmediumtable.Sonetmediumentry[i]}
    }
    sonetmediumtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetmediumtable.EntityData)
}

// SONETMIB_Sonetmediumtable_Sonetmediumentry
// An entry in the SONET/SDH Medium table.
type SONETMIB_Sonetmediumtable_Sonetmediumentry struct {
    EntityData types.CommonEntityData
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

func (sonetmediumentry *SONETMIB_Sonetmediumtable_Sonetmediumentry) GetEntityData() *types.CommonEntityData {
    sonetmediumentry.EntityData.YFilter = sonetmediumentry.YFilter
    sonetmediumentry.EntityData.YangName = "sonetMediumEntry"
    sonetmediumentry.EntityData.BundleName = "cisco_ios_xe"
    sonetmediumentry.EntityData.ParentYangName = "sonetMediumTable"
    sonetmediumentry.EntityData.SegmentPath = "sonetMediumEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetmediumentry.Ifindex) + "']"
    sonetmediumentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetmediumentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetmediumentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetmediumentry.EntityData.Children = make(map[string]types.YChild)
    sonetmediumentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetmediumentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetmediumentry.Ifindex}
    sonetmediumentry.EntityData.Leafs["sonetMediumType"] = types.YLeaf{"Sonetmediumtype", sonetmediumentry.Sonetmediumtype}
    sonetmediumentry.EntityData.Leafs["sonetMediumTimeElapsed"] = types.YLeaf{"Sonetmediumtimeelapsed", sonetmediumentry.Sonetmediumtimeelapsed}
    sonetmediumentry.EntityData.Leafs["sonetMediumValidIntervals"] = types.YLeaf{"Sonetmediumvalidintervals", sonetmediumentry.Sonetmediumvalidintervals}
    sonetmediumentry.EntityData.Leafs["sonetMediumLineCoding"] = types.YLeaf{"Sonetmediumlinecoding", sonetmediumentry.Sonetmediumlinecoding}
    sonetmediumentry.EntityData.Leafs["sonetMediumLineType"] = types.YLeaf{"Sonetmediumlinetype", sonetmediumentry.Sonetmediumlinetype}
    sonetmediumentry.EntityData.Leafs["sonetMediumCircuitIdentifier"] = types.YLeaf{"Sonetmediumcircuitidentifier", sonetmediumentry.Sonetmediumcircuitidentifier}
    sonetmediumentry.EntityData.Leafs["sonetMediumInvalidIntervals"] = types.YLeaf{"Sonetmediuminvalidintervals", sonetmediumentry.Sonetmediuminvalidintervals}
    sonetmediumentry.EntityData.Leafs["sonetMediumLoopbackConfig"] = types.YLeaf{"Sonetmediumloopbackconfig", sonetmediumentry.Sonetmediumloopbackconfig}
    return &(sonetmediumentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Current table. The type is slice of
    // SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry.
    Sonetsectioncurrententry []SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry
}

func (sonetsectioncurrenttable *SONETMIB_Sonetsectioncurrenttable) GetEntityData() *types.CommonEntityData {
    sonetsectioncurrenttable.EntityData.YFilter = sonetsectioncurrenttable.YFilter
    sonetsectioncurrenttable.EntityData.YangName = "sonetSectionCurrentTable"
    sonetsectioncurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetsectioncurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetsectioncurrenttable.EntityData.SegmentPath = "sonetSectionCurrentTable"
    sonetsectioncurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetsectioncurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetsectioncurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetsectioncurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetsectioncurrenttable.EntityData.Children["sonetSectionCurrentEntry"] = types.YChild{"Sonetsectioncurrententry", nil}
    for i := range sonetsectioncurrenttable.Sonetsectioncurrententry {
        sonetsectioncurrenttable.EntityData.Children[types.GetSegmentPath(&sonetsectioncurrenttable.Sonetsectioncurrententry[i])] = types.YChild{"Sonetsectioncurrententry", &sonetsectioncurrenttable.Sonetsectioncurrententry[i]}
    }
    sonetsectioncurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetsectioncurrenttable.EntityData)
}

// SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry
// An entry in the SONET/SDH Section Current table.
type SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetsectioncurrententry *SONETMIB_Sonetsectioncurrenttable_Sonetsectioncurrententry) GetEntityData() *types.CommonEntityData {
    sonetsectioncurrententry.EntityData.YFilter = sonetsectioncurrententry.YFilter
    sonetsectioncurrententry.EntityData.YangName = "sonetSectionCurrentEntry"
    sonetsectioncurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetsectioncurrententry.EntityData.ParentYangName = "sonetSectionCurrentTable"
    sonetsectioncurrententry.EntityData.SegmentPath = "sonetSectionCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetsectioncurrententry.Ifindex) + "']"
    sonetsectioncurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetsectioncurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetsectioncurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetsectioncurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetsectioncurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetsectioncurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetsectioncurrententry.Ifindex}
    sonetsectioncurrententry.EntityData.Leafs["sonetSectionCurrentStatus"] = types.YLeaf{"Sonetsectioncurrentstatus", sonetsectioncurrententry.Sonetsectioncurrentstatus}
    sonetsectioncurrententry.EntityData.Leafs["sonetSectionCurrentESs"] = types.YLeaf{"Sonetsectioncurrentess", sonetsectioncurrententry.Sonetsectioncurrentess}
    sonetsectioncurrententry.EntityData.Leafs["sonetSectionCurrentSESs"] = types.YLeaf{"Sonetsectioncurrentsess", sonetsectioncurrententry.Sonetsectioncurrentsess}
    sonetsectioncurrententry.EntityData.Leafs["sonetSectionCurrentSEFSs"] = types.YLeaf{"Sonetsectioncurrentsefss", sonetsectioncurrententry.Sonetsectioncurrentsefss}
    sonetsectioncurrententry.EntityData.Leafs["sonetSectionCurrentCVs"] = types.YLeaf{"Sonetsectioncurrentcvs", sonetsectioncurrententry.Sonetsectioncurrentcvs}
    return &(sonetsectioncurrententry.EntityData)
}

// SONETMIB_Sonetsectionintervaltable
// The SONET/SDH Section Interval table.
type SONETMIB_Sonetsectionintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Interval table. The type is slice of
    // SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry.
    Sonetsectionintervalentry []SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry
}

func (sonetsectionintervaltable *SONETMIB_Sonetsectionintervaltable) GetEntityData() *types.CommonEntityData {
    sonetsectionintervaltable.EntityData.YFilter = sonetsectionintervaltable.YFilter
    sonetsectionintervaltable.EntityData.YangName = "sonetSectionIntervalTable"
    sonetsectionintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetsectionintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetsectionintervaltable.EntityData.SegmentPath = "sonetSectionIntervalTable"
    sonetsectionintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetsectionintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetsectionintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetsectionintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetsectionintervaltable.EntityData.Children["sonetSectionIntervalEntry"] = types.YChild{"Sonetsectionintervalentry", nil}
    for i := range sonetsectionintervaltable.Sonetsectionintervalentry {
        sonetsectionintervaltable.EntityData.Children[types.GetSegmentPath(&sonetsectionintervaltable.Sonetsectionintervalentry[i])] = types.YChild{"Sonetsectionintervalentry", &sonetsectionintervaltable.Sonetsectionintervalentry[i]}
    }
    sonetsectionintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetsectionintervaltable.EntityData)
}

// SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry
// An entry in the SONET/SDH Section Interval table.
type SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetsectionintervalentry *SONETMIB_Sonetsectionintervaltable_Sonetsectionintervalentry) GetEntityData() *types.CommonEntityData {
    sonetsectionintervalentry.EntityData.YFilter = sonetsectionintervalentry.YFilter
    sonetsectionintervalentry.EntityData.YangName = "sonetSectionIntervalEntry"
    sonetsectionintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetsectionintervalentry.EntityData.ParentYangName = "sonetSectionIntervalTable"
    sonetsectionintervalentry.EntityData.SegmentPath = "sonetSectionIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetsectionintervalentry.Ifindex) + "']" + "[sonetSectionIntervalNumber='" + fmt.Sprintf("%v", sonetsectionintervalentry.Sonetsectionintervalnumber) + "']"
    sonetsectionintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetsectionintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetsectionintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetsectionintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetsectionintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetsectionintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetsectionintervalentry.Ifindex}
    sonetsectionintervalentry.EntityData.Leafs["sonetSectionIntervalNumber"] = types.YLeaf{"Sonetsectionintervalnumber", sonetsectionintervalentry.Sonetsectionintervalnumber}
    sonetsectionintervalentry.EntityData.Leafs["sonetSectionIntervalESs"] = types.YLeaf{"Sonetsectionintervaless", sonetsectionintervalentry.Sonetsectionintervaless}
    sonetsectionintervalentry.EntityData.Leafs["sonetSectionIntervalSESs"] = types.YLeaf{"Sonetsectionintervalsess", sonetsectionintervalentry.Sonetsectionintervalsess}
    sonetsectionintervalentry.EntityData.Leafs["sonetSectionIntervalSEFSs"] = types.YLeaf{"Sonetsectionintervalsefss", sonetsectionintervalentry.Sonetsectionintervalsefss}
    sonetsectionintervalentry.EntityData.Leafs["sonetSectionIntervalCVs"] = types.YLeaf{"Sonetsectionintervalcvs", sonetsectionintervalentry.Sonetsectionintervalcvs}
    sonetsectionintervalentry.EntityData.Leafs["sonetSectionIntervalValidData"] = types.YLeaf{"Sonetsectionintervalvaliddata", sonetsectionintervalentry.Sonetsectionintervalvaliddata}
    return &(sonetsectionintervalentry.EntityData)
}

// SONETMIB_Sonetlinecurrenttable
// The SONET/SDH Line Current table.
type SONETMIB_Sonetlinecurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Current table. The type is slice of
    // SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry.
    Sonetlinecurrententry []SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry
}

func (sonetlinecurrenttable *SONETMIB_Sonetlinecurrenttable) GetEntityData() *types.CommonEntityData {
    sonetlinecurrenttable.EntityData.YFilter = sonetlinecurrenttable.YFilter
    sonetlinecurrenttable.EntityData.YangName = "sonetLineCurrentTable"
    sonetlinecurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetlinecurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetlinecurrenttable.EntityData.SegmentPath = "sonetLineCurrentTable"
    sonetlinecurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetlinecurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetlinecurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetlinecurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetlinecurrenttable.EntityData.Children["sonetLineCurrentEntry"] = types.YChild{"Sonetlinecurrententry", nil}
    for i := range sonetlinecurrenttable.Sonetlinecurrententry {
        sonetlinecurrenttable.EntityData.Children[types.GetSegmentPath(&sonetlinecurrenttable.Sonetlinecurrententry[i])] = types.YChild{"Sonetlinecurrententry", &sonetlinecurrenttable.Sonetlinecurrententry[i]}
    }
    sonetlinecurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetlinecurrenttable.EntityData)
}

// SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry
// An entry in the SONET/SDH Line Current table.
type SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetlinecurrententry *SONETMIB_Sonetlinecurrenttable_Sonetlinecurrententry) GetEntityData() *types.CommonEntityData {
    sonetlinecurrententry.EntityData.YFilter = sonetlinecurrententry.YFilter
    sonetlinecurrententry.EntityData.YangName = "sonetLineCurrentEntry"
    sonetlinecurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetlinecurrententry.EntityData.ParentYangName = "sonetLineCurrentTable"
    sonetlinecurrententry.EntityData.SegmentPath = "sonetLineCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetlinecurrententry.Ifindex) + "']"
    sonetlinecurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetlinecurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetlinecurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetlinecurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetlinecurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetlinecurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetlinecurrententry.Ifindex}
    sonetlinecurrententry.EntityData.Leafs["sonetLineCurrentStatus"] = types.YLeaf{"Sonetlinecurrentstatus", sonetlinecurrententry.Sonetlinecurrentstatus}
    sonetlinecurrententry.EntityData.Leafs["sonetLineCurrentESs"] = types.YLeaf{"Sonetlinecurrentess", sonetlinecurrententry.Sonetlinecurrentess}
    sonetlinecurrententry.EntityData.Leafs["sonetLineCurrentSESs"] = types.YLeaf{"Sonetlinecurrentsess", sonetlinecurrententry.Sonetlinecurrentsess}
    sonetlinecurrententry.EntityData.Leafs["sonetLineCurrentCVs"] = types.YLeaf{"Sonetlinecurrentcvs", sonetlinecurrententry.Sonetlinecurrentcvs}
    sonetlinecurrententry.EntityData.Leafs["sonetLineCurrentUASs"] = types.YLeaf{"Sonetlinecurrentuass", sonetlinecurrententry.Sonetlinecurrentuass}
    return &(sonetlinecurrententry.EntityData)
}

// SONETMIB_Sonetlineintervaltable
// The SONET/SDH Line Interval table.
type SONETMIB_Sonetlineintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Interval table. The type is slice of
    // SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry.
    Sonetlineintervalentry []SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry
}

func (sonetlineintervaltable *SONETMIB_Sonetlineintervaltable) GetEntityData() *types.CommonEntityData {
    sonetlineintervaltable.EntityData.YFilter = sonetlineintervaltable.YFilter
    sonetlineintervaltable.EntityData.YangName = "sonetLineIntervalTable"
    sonetlineintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetlineintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetlineintervaltable.EntityData.SegmentPath = "sonetLineIntervalTable"
    sonetlineintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetlineintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetlineintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetlineintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetlineintervaltable.EntityData.Children["sonetLineIntervalEntry"] = types.YChild{"Sonetlineintervalentry", nil}
    for i := range sonetlineintervaltable.Sonetlineintervalentry {
        sonetlineintervaltable.EntityData.Children[types.GetSegmentPath(&sonetlineintervaltable.Sonetlineintervalentry[i])] = types.YChild{"Sonetlineintervalentry", &sonetlineintervaltable.Sonetlineintervalentry[i]}
    }
    sonetlineintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetlineintervaltable.EntityData)
}

// SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry
// An entry in the SONET/SDH Line Interval table.
type SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetlineintervalentry *SONETMIB_Sonetlineintervaltable_Sonetlineintervalentry) GetEntityData() *types.CommonEntityData {
    sonetlineintervalentry.EntityData.YFilter = sonetlineintervalentry.YFilter
    sonetlineintervalentry.EntityData.YangName = "sonetLineIntervalEntry"
    sonetlineintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetlineintervalentry.EntityData.ParentYangName = "sonetLineIntervalTable"
    sonetlineintervalentry.EntityData.SegmentPath = "sonetLineIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetlineintervalentry.Ifindex) + "']" + "[sonetLineIntervalNumber='" + fmt.Sprintf("%v", sonetlineintervalentry.Sonetlineintervalnumber) + "']"
    sonetlineintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetlineintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetlineintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetlineintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetlineintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetlineintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetlineintervalentry.Ifindex}
    sonetlineintervalentry.EntityData.Leafs["sonetLineIntervalNumber"] = types.YLeaf{"Sonetlineintervalnumber", sonetlineintervalentry.Sonetlineintervalnumber}
    sonetlineintervalentry.EntityData.Leafs["sonetLineIntervalESs"] = types.YLeaf{"Sonetlineintervaless", sonetlineintervalentry.Sonetlineintervaless}
    sonetlineintervalentry.EntityData.Leafs["sonetLineIntervalSESs"] = types.YLeaf{"Sonetlineintervalsess", sonetlineintervalentry.Sonetlineintervalsess}
    sonetlineintervalentry.EntityData.Leafs["sonetLineIntervalCVs"] = types.YLeaf{"Sonetlineintervalcvs", sonetlineintervalentry.Sonetlineintervalcvs}
    sonetlineintervalentry.EntityData.Leafs["sonetLineIntervalUASs"] = types.YLeaf{"Sonetlineintervaluass", sonetlineintervalentry.Sonetlineintervaluass}
    sonetlineintervalentry.EntityData.Leafs["sonetLineIntervalValidData"] = types.YLeaf{"Sonetlineintervalvaliddata", sonetlineintervalentry.Sonetlineintervalvaliddata}
    return &(sonetlineintervalentry.EntityData)
}

// SONETMIB_Sonetfarendlinecurrenttable
// The SONET/SDH Far End Line Current table.
type SONETMIB_Sonetfarendlinecurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Current table. The type is slice of
    // SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry.
    Sonetfarendlinecurrententry []SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry
}

func (sonetfarendlinecurrenttable *SONETMIB_Sonetfarendlinecurrenttable) GetEntityData() *types.CommonEntityData {
    sonetfarendlinecurrenttable.EntityData.YFilter = sonetfarendlinecurrenttable.YFilter
    sonetfarendlinecurrenttable.EntityData.YangName = "sonetFarEndLineCurrentTable"
    sonetfarendlinecurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendlinecurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetfarendlinecurrenttable.EntityData.SegmentPath = "sonetFarEndLineCurrentTable"
    sonetfarendlinecurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendlinecurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendlinecurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendlinecurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetfarendlinecurrenttable.EntityData.Children["sonetFarEndLineCurrentEntry"] = types.YChild{"Sonetfarendlinecurrententry", nil}
    for i := range sonetfarendlinecurrenttable.Sonetfarendlinecurrententry {
        sonetfarendlinecurrenttable.EntityData.Children[types.GetSegmentPath(&sonetfarendlinecurrenttable.Sonetfarendlinecurrententry[i])] = types.YChild{"Sonetfarendlinecurrententry", &sonetfarendlinecurrenttable.Sonetfarendlinecurrententry[i]}
    }
    sonetfarendlinecurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetfarendlinecurrenttable.EntityData)
}

// SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry
// An entry in the SONET/SDH Far End Line Current table.
type SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetfarendlinecurrententry *SONETMIB_Sonetfarendlinecurrenttable_Sonetfarendlinecurrententry) GetEntityData() *types.CommonEntityData {
    sonetfarendlinecurrententry.EntityData.YFilter = sonetfarendlinecurrententry.YFilter
    sonetfarendlinecurrententry.EntityData.YangName = "sonetFarEndLineCurrentEntry"
    sonetfarendlinecurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendlinecurrententry.EntityData.ParentYangName = "sonetFarEndLineCurrentTable"
    sonetfarendlinecurrententry.EntityData.SegmentPath = "sonetFarEndLineCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendlinecurrententry.Ifindex) + "']"
    sonetfarendlinecurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendlinecurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendlinecurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendlinecurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetfarendlinecurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetfarendlinecurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetfarendlinecurrententry.Ifindex}
    sonetfarendlinecurrententry.EntityData.Leafs["sonetFarEndLineCurrentESs"] = types.YLeaf{"Sonetfarendlinecurrentess", sonetfarendlinecurrententry.Sonetfarendlinecurrentess}
    sonetfarendlinecurrententry.EntityData.Leafs["sonetFarEndLineCurrentSESs"] = types.YLeaf{"Sonetfarendlinecurrentsess", sonetfarendlinecurrententry.Sonetfarendlinecurrentsess}
    sonetfarendlinecurrententry.EntityData.Leafs["sonetFarEndLineCurrentCVs"] = types.YLeaf{"Sonetfarendlinecurrentcvs", sonetfarendlinecurrententry.Sonetfarendlinecurrentcvs}
    sonetfarendlinecurrententry.EntityData.Leafs["sonetFarEndLineCurrentUASs"] = types.YLeaf{"Sonetfarendlinecurrentuass", sonetfarendlinecurrententry.Sonetfarendlinecurrentuass}
    return &(sonetfarendlinecurrententry.EntityData)
}

// SONETMIB_Sonetfarendlineintervaltable
// The SONET/SDH Far End Line Interval table.
type SONETMIB_Sonetfarendlineintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Interval table. The type is slice of
    // SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry.
    Sonetfarendlineintervalentry []SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry
}

func (sonetfarendlineintervaltable *SONETMIB_Sonetfarendlineintervaltable) GetEntityData() *types.CommonEntityData {
    sonetfarendlineintervaltable.EntityData.YFilter = sonetfarendlineintervaltable.YFilter
    sonetfarendlineintervaltable.EntityData.YangName = "sonetFarEndLineIntervalTable"
    sonetfarendlineintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendlineintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetfarendlineintervaltable.EntityData.SegmentPath = "sonetFarEndLineIntervalTable"
    sonetfarendlineintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendlineintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendlineintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendlineintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetfarendlineintervaltable.EntityData.Children["sonetFarEndLineIntervalEntry"] = types.YChild{"Sonetfarendlineintervalentry", nil}
    for i := range sonetfarendlineintervaltable.Sonetfarendlineintervalentry {
        sonetfarendlineintervaltable.EntityData.Children[types.GetSegmentPath(&sonetfarendlineintervaltable.Sonetfarendlineintervalentry[i])] = types.YChild{"Sonetfarendlineintervalentry", &sonetfarendlineintervaltable.Sonetfarendlineintervalentry[i]}
    }
    sonetfarendlineintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetfarendlineintervaltable.EntityData)
}

// SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry
// An entry in the SONET/SDH Far
// End Line Interval table.
type SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetfarendlineintervalentry *SONETMIB_Sonetfarendlineintervaltable_Sonetfarendlineintervalentry) GetEntityData() *types.CommonEntityData {
    sonetfarendlineintervalentry.EntityData.YFilter = sonetfarendlineintervalentry.YFilter
    sonetfarendlineintervalentry.EntityData.YangName = "sonetFarEndLineIntervalEntry"
    sonetfarendlineintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendlineintervalentry.EntityData.ParentYangName = "sonetFarEndLineIntervalTable"
    sonetfarendlineintervalentry.EntityData.SegmentPath = "sonetFarEndLineIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendlineintervalentry.Ifindex) + "']" + "[sonetFarEndLineIntervalNumber='" + fmt.Sprintf("%v", sonetfarendlineintervalentry.Sonetfarendlineintervalnumber) + "']"
    sonetfarendlineintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendlineintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendlineintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendlineintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetfarendlineintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetfarendlineintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetfarendlineintervalentry.Ifindex}
    sonetfarendlineintervalentry.EntityData.Leafs["sonetFarEndLineIntervalNumber"] = types.YLeaf{"Sonetfarendlineintervalnumber", sonetfarendlineintervalentry.Sonetfarendlineintervalnumber}
    sonetfarendlineintervalentry.EntityData.Leafs["sonetFarEndLineIntervalESs"] = types.YLeaf{"Sonetfarendlineintervaless", sonetfarendlineintervalentry.Sonetfarendlineintervaless}
    sonetfarendlineintervalentry.EntityData.Leafs["sonetFarEndLineIntervalSESs"] = types.YLeaf{"Sonetfarendlineintervalsess", sonetfarendlineintervalentry.Sonetfarendlineintervalsess}
    sonetfarendlineintervalentry.EntityData.Leafs["sonetFarEndLineIntervalCVs"] = types.YLeaf{"Sonetfarendlineintervalcvs", sonetfarendlineintervalentry.Sonetfarendlineintervalcvs}
    sonetfarendlineintervalentry.EntityData.Leafs["sonetFarEndLineIntervalUASs"] = types.YLeaf{"Sonetfarendlineintervaluass", sonetfarendlineintervalentry.Sonetfarendlineintervaluass}
    sonetfarendlineintervalentry.EntityData.Leafs["sonetFarEndLineIntervalValidData"] = types.YLeaf{"Sonetfarendlineintervalvaliddata", sonetfarendlineintervalentry.Sonetfarendlineintervalvaliddata}
    return &(sonetfarendlineintervalentry.EntityData)
}

// SONETMIB_Sonetpathcurrenttable
// The SONET/SDH Path Current table.
type SONETMIB_Sonetpathcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Current table. The type is slice of
    // SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry.
    Sonetpathcurrententry []SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry
}

func (sonetpathcurrenttable *SONETMIB_Sonetpathcurrenttable) GetEntityData() *types.CommonEntityData {
    sonetpathcurrenttable.EntityData.YFilter = sonetpathcurrenttable.YFilter
    sonetpathcurrenttable.EntityData.YangName = "sonetPathCurrentTable"
    sonetpathcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetpathcurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetpathcurrenttable.EntityData.SegmentPath = "sonetPathCurrentTable"
    sonetpathcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetpathcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetpathcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetpathcurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetpathcurrenttable.EntityData.Children["sonetPathCurrentEntry"] = types.YChild{"Sonetpathcurrententry", nil}
    for i := range sonetpathcurrenttable.Sonetpathcurrententry {
        sonetpathcurrenttable.EntityData.Children[types.GetSegmentPath(&sonetpathcurrenttable.Sonetpathcurrententry[i])] = types.YChild{"Sonetpathcurrententry", &sonetpathcurrenttable.Sonetpathcurrententry[i]}
    }
    sonetpathcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetpathcurrenttable.EntityData)
}

// SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry
// An entry in the SONET/SDH Path Current table.
type SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetpathcurrententry *SONETMIB_Sonetpathcurrenttable_Sonetpathcurrententry) GetEntityData() *types.CommonEntityData {
    sonetpathcurrententry.EntityData.YFilter = sonetpathcurrententry.YFilter
    sonetpathcurrententry.EntityData.YangName = "sonetPathCurrentEntry"
    sonetpathcurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetpathcurrententry.EntityData.ParentYangName = "sonetPathCurrentTable"
    sonetpathcurrententry.EntityData.SegmentPath = "sonetPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetpathcurrententry.Ifindex) + "']"
    sonetpathcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetpathcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetpathcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetpathcurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetpathcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetpathcurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetpathcurrententry.Ifindex}
    sonetpathcurrententry.EntityData.Leafs["sonetPathCurrentWidth"] = types.YLeaf{"Sonetpathcurrentwidth", sonetpathcurrententry.Sonetpathcurrentwidth}
    sonetpathcurrententry.EntityData.Leafs["sonetPathCurrentStatus"] = types.YLeaf{"Sonetpathcurrentstatus", sonetpathcurrententry.Sonetpathcurrentstatus}
    sonetpathcurrententry.EntityData.Leafs["sonetPathCurrentESs"] = types.YLeaf{"Sonetpathcurrentess", sonetpathcurrententry.Sonetpathcurrentess}
    sonetpathcurrententry.EntityData.Leafs["sonetPathCurrentSESs"] = types.YLeaf{"Sonetpathcurrentsess", sonetpathcurrententry.Sonetpathcurrentsess}
    sonetpathcurrententry.EntityData.Leafs["sonetPathCurrentCVs"] = types.YLeaf{"Sonetpathcurrentcvs", sonetpathcurrententry.Sonetpathcurrentcvs}
    sonetpathcurrententry.EntityData.Leafs["sonetPathCurrentUASs"] = types.YLeaf{"Sonetpathcurrentuass", sonetpathcurrententry.Sonetpathcurrentuass}
    sonetpathcurrententry.EntityData.Leafs["cspSonetPathPayload"] = types.YLeaf{"Cspsonetpathpayload", sonetpathcurrententry.Cspsonetpathpayload}
    sonetpathcurrententry.EntityData.Leafs["cspTributaryMappingType"] = types.YLeaf{"Csptributarymappingtype", sonetpathcurrententry.Csptributarymappingtype}
    sonetpathcurrententry.EntityData.Leafs["cspSignallingTransportMode"] = types.YLeaf{"Cspsignallingtransportmode", sonetpathcurrententry.Cspsignallingtransportmode}
    sonetpathcurrententry.EntityData.Leafs["cspTributaryGroupingType"] = types.YLeaf{"Csptributarygroupingtype", sonetpathcurrententry.Csptributarygroupingtype}
    return &(sonetpathcurrententry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Interval table. The type is slice of
    // SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry.
    Sonetpathintervalentry []SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry
}

func (sonetpathintervaltable *SONETMIB_Sonetpathintervaltable) GetEntityData() *types.CommonEntityData {
    sonetpathintervaltable.EntityData.YFilter = sonetpathintervaltable.YFilter
    sonetpathintervaltable.EntityData.YangName = "sonetPathIntervalTable"
    sonetpathintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetpathintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetpathintervaltable.EntityData.SegmentPath = "sonetPathIntervalTable"
    sonetpathintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetpathintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetpathintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetpathintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetpathintervaltable.EntityData.Children["sonetPathIntervalEntry"] = types.YChild{"Sonetpathintervalentry", nil}
    for i := range sonetpathintervaltable.Sonetpathintervalentry {
        sonetpathintervaltable.EntityData.Children[types.GetSegmentPath(&sonetpathintervaltable.Sonetpathintervalentry[i])] = types.YChild{"Sonetpathintervalentry", &sonetpathintervaltable.Sonetpathintervalentry[i]}
    }
    sonetpathintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetpathintervaltable.EntityData)
}

// SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry
// An entry in the SONET/SDH Path Interval table.
type SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetpathintervalentry *SONETMIB_Sonetpathintervaltable_Sonetpathintervalentry) GetEntityData() *types.CommonEntityData {
    sonetpathintervalentry.EntityData.YFilter = sonetpathintervalentry.YFilter
    sonetpathintervalentry.EntityData.YangName = "sonetPathIntervalEntry"
    sonetpathintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetpathintervalentry.EntityData.ParentYangName = "sonetPathIntervalTable"
    sonetpathintervalentry.EntityData.SegmentPath = "sonetPathIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetpathintervalentry.Ifindex) + "']" + "[sonetPathIntervalNumber='" + fmt.Sprintf("%v", sonetpathintervalentry.Sonetpathintervalnumber) + "']"
    sonetpathintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetpathintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetpathintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetpathintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetpathintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetpathintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetpathintervalentry.Ifindex}
    sonetpathintervalentry.EntityData.Leafs["sonetPathIntervalNumber"] = types.YLeaf{"Sonetpathintervalnumber", sonetpathintervalentry.Sonetpathintervalnumber}
    sonetpathintervalentry.EntityData.Leafs["sonetPathIntervalESs"] = types.YLeaf{"Sonetpathintervaless", sonetpathintervalentry.Sonetpathintervaless}
    sonetpathintervalentry.EntityData.Leafs["sonetPathIntervalSESs"] = types.YLeaf{"Sonetpathintervalsess", sonetpathintervalentry.Sonetpathintervalsess}
    sonetpathintervalentry.EntityData.Leafs["sonetPathIntervalCVs"] = types.YLeaf{"Sonetpathintervalcvs", sonetpathintervalentry.Sonetpathintervalcvs}
    sonetpathintervalentry.EntityData.Leafs["sonetPathIntervalUASs"] = types.YLeaf{"Sonetpathintervaluass", sonetpathintervalentry.Sonetpathintervaluass}
    sonetpathintervalentry.EntityData.Leafs["sonetPathIntervalValidData"] = types.YLeaf{"Sonetpathintervalvaliddata", sonetpathintervalentry.Sonetpathintervalvaliddata}
    return &(sonetpathintervalentry.EntityData)
}

// SONETMIB_Sonetfarendpathcurrenttable
// The SONET/SDH Far End Path Current table.
type SONETMIB_Sonetfarendpathcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Current table. The type is slice of
    // SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry.
    Sonetfarendpathcurrententry []SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry
}

func (sonetfarendpathcurrenttable *SONETMIB_Sonetfarendpathcurrenttable) GetEntityData() *types.CommonEntityData {
    sonetfarendpathcurrenttable.EntityData.YFilter = sonetfarendpathcurrenttable.YFilter
    sonetfarendpathcurrenttable.EntityData.YangName = "sonetFarEndPathCurrentTable"
    sonetfarendpathcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendpathcurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetfarendpathcurrenttable.EntityData.SegmentPath = "sonetFarEndPathCurrentTable"
    sonetfarendpathcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendpathcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendpathcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendpathcurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetfarendpathcurrenttable.EntityData.Children["sonetFarEndPathCurrentEntry"] = types.YChild{"Sonetfarendpathcurrententry", nil}
    for i := range sonetfarendpathcurrenttable.Sonetfarendpathcurrententry {
        sonetfarendpathcurrenttable.EntityData.Children[types.GetSegmentPath(&sonetfarendpathcurrenttable.Sonetfarendpathcurrententry[i])] = types.YChild{"Sonetfarendpathcurrententry", &sonetfarendpathcurrenttable.Sonetfarendpathcurrententry[i]}
    }
    sonetfarendpathcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetfarendpathcurrenttable.EntityData)
}

// SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry
// An entry in the SONET/SDH Far End Path Current table.
type SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetfarendpathcurrententry *SONETMIB_Sonetfarendpathcurrenttable_Sonetfarendpathcurrententry) GetEntityData() *types.CommonEntityData {
    sonetfarendpathcurrententry.EntityData.YFilter = sonetfarendpathcurrententry.YFilter
    sonetfarendpathcurrententry.EntityData.YangName = "sonetFarEndPathCurrentEntry"
    sonetfarendpathcurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendpathcurrententry.EntityData.ParentYangName = "sonetFarEndPathCurrentTable"
    sonetfarendpathcurrententry.EntityData.SegmentPath = "sonetFarEndPathCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendpathcurrententry.Ifindex) + "']"
    sonetfarendpathcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendpathcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendpathcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendpathcurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetfarendpathcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetfarendpathcurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetfarendpathcurrententry.Ifindex}
    sonetfarendpathcurrententry.EntityData.Leafs["sonetFarEndPathCurrentESs"] = types.YLeaf{"Sonetfarendpathcurrentess", sonetfarendpathcurrententry.Sonetfarendpathcurrentess}
    sonetfarendpathcurrententry.EntityData.Leafs["sonetFarEndPathCurrentSESs"] = types.YLeaf{"Sonetfarendpathcurrentsess", sonetfarendpathcurrententry.Sonetfarendpathcurrentsess}
    sonetfarendpathcurrententry.EntityData.Leafs["sonetFarEndPathCurrentCVs"] = types.YLeaf{"Sonetfarendpathcurrentcvs", sonetfarendpathcurrententry.Sonetfarendpathcurrentcvs}
    sonetfarendpathcurrententry.EntityData.Leafs["sonetFarEndPathCurrentUASs"] = types.YLeaf{"Sonetfarendpathcurrentuass", sonetfarendpathcurrententry.Sonetfarendpathcurrentuass}
    return &(sonetfarendpathcurrententry.EntityData)
}

// SONETMIB_Sonetfarendpathintervaltable
// The SONET/SDH Far End Path Interval table.
type SONETMIB_Sonetfarendpathintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Interval table. The type is slice of
    // SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry.
    Sonetfarendpathintervalentry []SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry
}

func (sonetfarendpathintervaltable *SONETMIB_Sonetfarendpathintervaltable) GetEntityData() *types.CommonEntityData {
    sonetfarendpathintervaltable.EntityData.YFilter = sonetfarendpathintervaltable.YFilter
    sonetfarendpathintervaltable.EntityData.YangName = "sonetFarEndPathIntervalTable"
    sonetfarendpathintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendpathintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetfarendpathintervaltable.EntityData.SegmentPath = "sonetFarEndPathIntervalTable"
    sonetfarendpathintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendpathintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendpathintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendpathintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetfarendpathintervaltable.EntityData.Children["sonetFarEndPathIntervalEntry"] = types.YChild{"Sonetfarendpathintervalentry", nil}
    for i := range sonetfarendpathintervaltable.Sonetfarendpathintervalentry {
        sonetfarendpathintervaltable.EntityData.Children[types.GetSegmentPath(&sonetfarendpathintervaltable.Sonetfarendpathintervalentry[i])] = types.YChild{"Sonetfarendpathintervalentry", &sonetfarendpathintervaltable.Sonetfarendpathintervalentry[i]}
    }
    sonetfarendpathintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetfarendpathintervaltable.EntityData)
}

// SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry
// An entry in the SONET/SDH Far
// End Path Interval table.
type SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetfarendpathintervalentry *SONETMIB_Sonetfarendpathintervaltable_Sonetfarendpathintervalentry) GetEntityData() *types.CommonEntityData {
    sonetfarendpathintervalentry.EntityData.YFilter = sonetfarendpathintervalentry.YFilter
    sonetfarendpathintervalentry.EntityData.YangName = "sonetFarEndPathIntervalEntry"
    sonetfarendpathintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendpathintervalentry.EntityData.ParentYangName = "sonetFarEndPathIntervalTable"
    sonetfarendpathintervalentry.EntityData.SegmentPath = "sonetFarEndPathIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendpathintervalentry.Ifindex) + "']" + "[sonetFarEndPathIntervalNumber='" + fmt.Sprintf("%v", sonetfarendpathintervalentry.Sonetfarendpathintervalnumber) + "']"
    sonetfarendpathintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendpathintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendpathintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendpathintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetfarendpathintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetfarendpathintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetfarendpathintervalentry.Ifindex}
    sonetfarendpathintervalentry.EntityData.Leafs["sonetFarEndPathIntervalNumber"] = types.YLeaf{"Sonetfarendpathintervalnumber", sonetfarendpathintervalentry.Sonetfarendpathintervalnumber}
    sonetfarendpathintervalentry.EntityData.Leafs["sonetFarEndPathIntervalESs"] = types.YLeaf{"Sonetfarendpathintervaless", sonetfarendpathintervalentry.Sonetfarendpathintervaless}
    sonetfarendpathintervalentry.EntityData.Leafs["sonetFarEndPathIntervalSESs"] = types.YLeaf{"Sonetfarendpathintervalsess", sonetfarendpathintervalentry.Sonetfarendpathintervalsess}
    sonetfarendpathintervalentry.EntityData.Leafs["sonetFarEndPathIntervalCVs"] = types.YLeaf{"Sonetfarendpathintervalcvs", sonetfarendpathintervalentry.Sonetfarendpathintervalcvs}
    sonetfarendpathintervalentry.EntityData.Leafs["sonetFarEndPathIntervalUASs"] = types.YLeaf{"Sonetfarendpathintervaluass", sonetfarendpathintervalentry.Sonetfarendpathintervaluass}
    sonetfarendpathintervalentry.EntityData.Leafs["sonetFarEndPathIntervalValidData"] = types.YLeaf{"Sonetfarendpathintervalvaliddata", sonetfarendpathintervalentry.Sonetfarendpathintervalvaliddata}
    return &(sonetfarendpathintervalentry.EntityData)
}

// SONETMIB_Sonetvtcurrenttable
// The SONET/SDH VT Current table.
type SONETMIB_Sonetvtcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH VT Current table. The type is slice of
    // SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry.
    Sonetvtcurrententry []SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry
}

func (sonetvtcurrenttable *SONETMIB_Sonetvtcurrenttable) GetEntityData() *types.CommonEntityData {
    sonetvtcurrenttable.EntityData.YFilter = sonetvtcurrenttable.YFilter
    sonetvtcurrenttable.EntityData.YangName = "sonetVTCurrentTable"
    sonetvtcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetvtcurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetvtcurrenttable.EntityData.SegmentPath = "sonetVTCurrentTable"
    sonetvtcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetvtcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetvtcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetvtcurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetvtcurrenttable.EntityData.Children["sonetVTCurrentEntry"] = types.YChild{"Sonetvtcurrententry", nil}
    for i := range sonetvtcurrenttable.Sonetvtcurrententry {
        sonetvtcurrenttable.EntityData.Children[types.GetSegmentPath(&sonetvtcurrenttable.Sonetvtcurrententry[i])] = types.YChild{"Sonetvtcurrententry", &sonetvtcurrenttable.Sonetvtcurrententry[i]}
    }
    sonetvtcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetvtcurrenttable.EntityData)
}

// SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry
// An entry in the SONET/SDH VT Current table.
type SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetvtcurrententry *SONETMIB_Sonetvtcurrenttable_Sonetvtcurrententry) GetEntityData() *types.CommonEntityData {
    sonetvtcurrententry.EntityData.YFilter = sonetvtcurrententry.YFilter
    sonetvtcurrententry.EntityData.YangName = "sonetVTCurrentEntry"
    sonetvtcurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetvtcurrententry.EntityData.ParentYangName = "sonetVTCurrentTable"
    sonetvtcurrententry.EntityData.SegmentPath = "sonetVTCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetvtcurrententry.Ifindex) + "']"
    sonetvtcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetvtcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetvtcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetvtcurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetvtcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetvtcurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetvtcurrententry.Ifindex}
    sonetvtcurrententry.EntityData.Leafs["sonetVTCurrentWidth"] = types.YLeaf{"Sonetvtcurrentwidth", sonetvtcurrententry.Sonetvtcurrentwidth}
    sonetvtcurrententry.EntityData.Leafs["sonetVTCurrentStatus"] = types.YLeaf{"Sonetvtcurrentstatus", sonetvtcurrententry.Sonetvtcurrentstatus}
    sonetvtcurrententry.EntityData.Leafs["sonetVTCurrentESs"] = types.YLeaf{"Sonetvtcurrentess", sonetvtcurrententry.Sonetvtcurrentess}
    sonetvtcurrententry.EntityData.Leafs["sonetVTCurrentSESs"] = types.YLeaf{"Sonetvtcurrentsess", sonetvtcurrententry.Sonetvtcurrentsess}
    sonetvtcurrententry.EntityData.Leafs["sonetVTCurrentCVs"] = types.YLeaf{"Sonetvtcurrentcvs", sonetvtcurrententry.Sonetvtcurrentcvs}
    sonetvtcurrententry.EntityData.Leafs["sonetVTCurrentUASs"] = types.YLeaf{"Sonetvtcurrentuass", sonetvtcurrententry.Sonetvtcurrentuass}
    return &(sonetvtcurrententry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH VT Interval table. The type is slice of
    // SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry.
    Sonetvtintervalentry []SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry
}

func (sonetvtintervaltable *SONETMIB_Sonetvtintervaltable) GetEntityData() *types.CommonEntityData {
    sonetvtintervaltable.EntityData.YFilter = sonetvtintervaltable.YFilter
    sonetvtintervaltable.EntityData.YangName = "sonetVTIntervalTable"
    sonetvtintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetvtintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetvtintervaltable.EntityData.SegmentPath = "sonetVTIntervalTable"
    sonetvtintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetvtintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetvtintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetvtintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetvtintervaltable.EntityData.Children["sonetVTIntervalEntry"] = types.YChild{"Sonetvtintervalentry", nil}
    for i := range sonetvtintervaltable.Sonetvtintervalentry {
        sonetvtintervaltable.EntityData.Children[types.GetSegmentPath(&sonetvtintervaltable.Sonetvtintervalentry[i])] = types.YChild{"Sonetvtintervalentry", &sonetvtintervaltable.Sonetvtintervalentry[i]}
    }
    sonetvtintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetvtintervaltable.EntityData)
}

// SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry
// An entry in the SONET/SDH VT Interval table.
type SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetvtintervalentry *SONETMIB_Sonetvtintervaltable_Sonetvtintervalentry) GetEntityData() *types.CommonEntityData {
    sonetvtintervalentry.EntityData.YFilter = sonetvtintervalentry.YFilter
    sonetvtintervalentry.EntityData.YangName = "sonetVTIntervalEntry"
    sonetvtintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetvtintervalentry.EntityData.ParentYangName = "sonetVTIntervalTable"
    sonetvtintervalentry.EntityData.SegmentPath = "sonetVTIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetvtintervalentry.Ifindex) + "']" + "[sonetVTIntervalNumber='" + fmt.Sprintf("%v", sonetvtintervalentry.Sonetvtintervalnumber) + "']"
    sonetvtintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetvtintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetvtintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetvtintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetvtintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetvtintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetvtintervalentry.Ifindex}
    sonetvtintervalentry.EntityData.Leafs["sonetVTIntervalNumber"] = types.YLeaf{"Sonetvtintervalnumber", sonetvtintervalentry.Sonetvtintervalnumber}
    sonetvtintervalentry.EntityData.Leafs["sonetVTIntervalESs"] = types.YLeaf{"Sonetvtintervaless", sonetvtintervalentry.Sonetvtintervaless}
    sonetvtintervalentry.EntityData.Leafs["sonetVTIntervalSESs"] = types.YLeaf{"Sonetvtintervalsess", sonetvtintervalentry.Sonetvtintervalsess}
    sonetvtintervalentry.EntityData.Leafs["sonetVTIntervalCVs"] = types.YLeaf{"Sonetvtintervalcvs", sonetvtintervalentry.Sonetvtintervalcvs}
    sonetvtintervalentry.EntityData.Leafs["sonetVTIntervalUASs"] = types.YLeaf{"Sonetvtintervaluass", sonetvtintervalentry.Sonetvtintervaluass}
    sonetvtintervalentry.EntityData.Leafs["sonetVTIntervalValidData"] = types.YLeaf{"Sonetvtintervalvaliddata", sonetvtintervalentry.Sonetvtintervalvaliddata}
    return &(sonetvtintervalentry.EntityData)
}

// SONETMIB_Sonetfarendvtcurrenttable
// The SONET/SDH Far End VT Current table.
type SONETMIB_Sonetfarendvtcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End VT Current table. The type is slice of
    // SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry.
    Sonetfarendvtcurrententry []SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry
}

func (sonetfarendvtcurrenttable *SONETMIB_Sonetfarendvtcurrenttable) GetEntityData() *types.CommonEntityData {
    sonetfarendvtcurrenttable.EntityData.YFilter = sonetfarendvtcurrenttable.YFilter
    sonetfarendvtcurrenttable.EntityData.YangName = "sonetFarEndVTCurrentTable"
    sonetfarendvtcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendvtcurrenttable.EntityData.ParentYangName = "SONET-MIB"
    sonetfarendvtcurrenttable.EntityData.SegmentPath = "sonetFarEndVTCurrentTable"
    sonetfarendvtcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendvtcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendvtcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendvtcurrenttable.EntityData.Children = make(map[string]types.YChild)
    sonetfarendvtcurrenttable.EntityData.Children["sonetFarEndVTCurrentEntry"] = types.YChild{"Sonetfarendvtcurrententry", nil}
    for i := range sonetfarendvtcurrenttable.Sonetfarendvtcurrententry {
        sonetfarendvtcurrenttable.EntityData.Children[types.GetSegmentPath(&sonetfarendvtcurrenttable.Sonetfarendvtcurrententry[i])] = types.YChild{"Sonetfarendvtcurrententry", &sonetfarendvtcurrenttable.Sonetfarendvtcurrententry[i]}
    }
    sonetfarendvtcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetfarendvtcurrenttable.EntityData)
}

// SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry
// An entry in the SONET/SDH Far End VT Current table.
type SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry struct {
    EntityData types.CommonEntityData
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

func (sonetfarendvtcurrententry *SONETMIB_Sonetfarendvtcurrenttable_Sonetfarendvtcurrententry) GetEntityData() *types.CommonEntityData {
    sonetfarendvtcurrententry.EntityData.YFilter = sonetfarendvtcurrententry.YFilter
    sonetfarendvtcurrententry.EntityData.YangName = "sonetFarEndVTCurrentEntry"
    sonetfarendvtcurrententry.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendvtcurrententry.EntityData.ParentYangName = "sonetFarEndVTCurrentTable"
    sonetfarendvtcurrententry.EntityData.SegmentPath = "sonetFarEndVTCurrentEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendvtcurrententry.Ifindex) + "']"
    sonetfarendvtcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendvtcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendvtcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendvtcurrententry.EntityData.Children = make(map[string]types.YChild)
    sonetfarendvtcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetfarendvtcurrententry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetfarendvtcurrententry.Ifindex}
    sonetfarendvtcurrententry.EntityData.Leafs["sonetFarEndVTCurrentESs"] = types.YLeaf{"Sonetfarendvtcurrentess", sonetfarendvtcurrententry.Sonetfarendvtcurrentess}
    sonetfarendvtcurrententry.EntityData.Leafs["sonetFarEndVTCurrentSESs"] = types.YLeaf{"Sonetfarendvtcurrentsess", sonetfarendvtcurrententry.Sonetfarendvtcurrentsess}
    sonetfarendvtcurrententry.EntityData.Leafs["sonetFarEndVTCurrentCVs"] = types.YLeaf{"Sonetfarendvtcurrentcvs", sonetfarendvtcurrententry.Sonetfarendvtcurrentcvs}
    sonetfarendvtcurrententry.EntityData.Leafs["sonetFarEndVTCurrentUASs"] = types.YLeaf{"Sonetfarendvtcurrentuass", sonetfarendvtcurrententry.Sonetfarendvtcurrentuass}
    return &(sonetfarendvtcurrententry.EntityData)
}

// SONETMIB_Sonetfarendvtintervaltable
// The SONET/SDH Far End VT Interval table.
type SONETMIB_Sonetfarendvtintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End VT Interval table. The type is slice of
    // SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry.
    Sonetfarendvtintervalentry []SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry
}

func (sonetfarendvtintervaltable *SONETMIB_Sonetfarendvtintervaltable) GetEntityData() *types.CommonEntityData {
    sonetfarendvtintervaltable.EntityData.YFilter = sonetfarendvtintervaltable.YFilter
    sonetfarendvtintervaltable.EntityData.YangName = "sonetFarEndVTIntervalTable"
    sonetfarendvtintervaltable.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendvtintervaltable.EntityData.ParentYangName = "SONET-MIB"
    sonetfarendvtintervaltable.EntityData.SegmentPath = "sonetFarEndVTIntervalTable"
    sonetfarendvtintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendvtintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendvtintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendvtintervaltable.EntityData.Children = make(map[string]types.YChild)
    sonetfarendvtintervaltable.EntityData.Children["sonetFarEndVTIntervalEntry"] = types.YChild{"Sonetfarendvtintervalentry", nil}
    for i := range sonetfarendvtintervaltable.Sonetfarendvtintervalentry {
        sonetfarendvtintervaltable.EntityData.Children[types.GetSegmentPath(&sonetfarendvtintervaltable.Sonetfarendvtintervalentry[i])] = types.YChild{"Sonetfarendvtintervalentry", &sonetfarendvtintervaltable.Sonetfarendvtintervalentry[i]}
    }
    sonetfarendvtintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sonetfarendvtintervaltable.EntityData)
}

// SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry
// An entry in the SONET/SDH Far
// End VT Interval table.
type SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry struct {
    EntityData types.CommonEntityData
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

func (sonetfarendvtintervalentry *SONETMIB_Sonetfarendvtintervaltable_Sonetfarendvtintervalentry) GetEntityData() *types.CommonEntityData {
    sonetfarendvtintervalentry.EntityData.YFilter = sonetfarendvtintervalentry.YFilter
    sonetfarendvtintervalentry.EntityData.YangName = "sonetFarEndVTIntervalEntry"
    sonetfarendvtintervalentry.EntityData.BundleName = "cisco_ios_xe"
    sonetfarendvtintervalentry.EntityData.ParentYangName = "sonetFarEndVTIntervalTable"
    sonetfarendvtintervalentry.EntityData.SegmentPath = "sonetFarEndVTIntervalEntry" + "[ifIndex='" + fmt.Sprintf("%v", sonetfarendvtintervalentry.Ifindex) + "']" + "[sonetFarEndVTIntervalNumber='" + fmt.Sprintf("%v", sonetfarendvtintervalentry.Sonetfarendvtintervalnumber) + "']"
    sonetfarendvtintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetfarendvtintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetfarendvtintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetfarendvtintervalentry.EntityData.Children = make(map[string]types.YChild)
    sonetfarendvtintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sonetfarendvtintervalentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", sonetfarendvtintervalentry.Ifindex}
    sonetfarendvtintervalentry.EntityData.Leafs["sonetFarEndVTIntervalNumber"] = types.YLeaf{"Sonetfarendvtintervalnumber", sonetfarendvtintervalentry.Sonetfarendvtintervalnumber}
    sonetfarendvtintervalentry.EntityData.Leafs["sonetFarEndVTIntervalESs"] = types.YLeaf{"Sonetfarendvtintervaless", sonetfarendvtintervalentry.Sonetfarendvtintervaless}
    sonetfarendvtintervalentry.EntityData.Leafs["sonetFarEndVTIntervalSESs"] = types.YLeaf{"Sonetfarendvtintervalsess", sonetfarendvtintervalentry.Sonetfarendvtintervalsess}
    sonetfarendvtintervalentry.EntityData.Leafs["sonetFarEndVTIntervalCVs"] = types.YLeaf{"Sonetfarendvtintervalcvs", sonetfarendvtintervalentry.Sonetfarendvtintervalcvs}
    sonetfarendvtintervalentry.EntityData.Leafs["sonetFarEndVTIntervalUASs"] = types.YLeaf{"Sonetfarendvtintervaluass", sonetfarendvtintervalentry.Sonetfarendvtintervaluass}
    sonetfarendvtintervalentry.EntityData.Leafs["sonetFarEndVTIntervalValidData"] = types.YLeaf{"Sonetfarendvtintervalvaliddata", sonetfarendvtintervalentry.Sonetfarendvtintervalvaliddata}
    return &(sonetfarendvtintervalentry.EntityData)
}

