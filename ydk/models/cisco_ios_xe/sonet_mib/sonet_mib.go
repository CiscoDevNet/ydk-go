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

    
    SonetMedium SONETMIB_SonetMedium

    // The SONET/SDH Medium table.
    SonetMediumTable SONETMIB_SonetMediumTable

    // The SONET/SDH Section Current table.
    SonetSectionCurrentTable SONETMIB_SonetSectionCurrentTable

    // The SONET/SDH Section Interval table.
    SonetSectionIntervalTable SONETMIB_SonetSectionIntervalTable

    // The SONET/SDH Line Current table.
    SonetLineCurrentTable SONETMIB_SonetLineCurrentTable

    // The SONET/SDH Line Interval table.
    SonetLineIntervalTable SONETMIB_SonetLineIntervalTable

    // The SONET/SDH Far End Line Current table.
    SonetFarEndLineCurrentTable SONETMIB_SonetFarEndLineCurrentTable

    // The SONET/SDH Far End Line Interval table.
    SonetFarEndLineIntervalTable SONETMIB_SonetFarEndLineIntervalTable

    // The SONET/SDH Path Current table.
    SonetPathCurrentTable SONETMIB_SonetPathCurrentTable

    // The SONET/SDH Path Interval table.
    SonetPathIntervalTable SONETMIB_SonetPathIntervalTable

    // The SONET/SDH Far End Path Current table.
    SonetFarEndPathCurrentTable SONETMIB_SonetFarEndPathCurrentTable

    // The SONET/SDH Far End Path Interval table.
    SonetFarEndPathIntervalTable SONETMIB_SonetFarEndPathIntervalTable

    // The SONET/SDH VT Current table.
    SonetVTCurrentTable SONETMIB_SonetVTCurrentTable

    // The SONET/SDH VT Interval table.
    SonetVTIntervalTable SONETMIB_SonetVTIntervalTable

    // The SONET/SDH Far End VT Current table.
    SonetFarEndVTCurrentTable SONETMIB_SonetFarEndVTCurrentTable

    // The SONET/SDH Far End VT Interval table.
    SonetFarEndVTIntervalTable SONETMIB_SonetFarEndVTIntervalTable
}

func (sONETMIB *SONETMIB) GetEntityData() *types.CommonEntityData {
    sONETMIB.EntityData.YFilter = sONETMIB.YFilter
    sONETMIB.EntityData.YangName = "SONET-MIB"
    sONETMIB.EntityData.BundleName = "cisco_ios_xe"
    sONETMIB.EntityData.ParentYangName = "SONET-MIB"
    sONETMIB.EntityData.SegmentPath = "SONET-MIB:SONET-MIB"
    sONETMIB.EntityData.AbsolutePath = sONETMIB.EntityData.SegmentPath
    sONETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sONETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sONETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sONETMIB.EntityData.Children = types.NewOrderedMap()
    sONETMIB.EntityData.Children.Append("sonetMedium", types.YChild{"SonetMedium", &sONETMIB.SonetMedium})
    sONETMIB.EntityData.Children.Append("sonetMediumTable", types.YChild{"SonetMediumTable", &sONETMIB.SonetMediumTable})
    sONETMIB.EntityData.Children.Append("sonetSectionCurrentTable", types.YChild{"SonetSectionCurrentTable", &sONETMIB.SonetSectionCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetSectionIntervalTable", types.YChild{"SonetSectionIntervalTable", &sONETMIB.SonetSectionIntervalTable})
    sONETMIB.EntityData.Children.Append("sonetLineCurrentTable", types.YChild{"SonetLineCurrentTable", &sONETMIB.SonetLineCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetLineIntervalTable", types.YChild{"SonetLineIntervalTable", &sONETMIB.SonetLineIntervalTable})
    sONETMIB.EntityData.Children.Append("sonetFarEndLineCurrentTable", types.YChild{"SonetFarEndLineCurrentTable", &sONETMIB.SonetFarEndLineCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetFarEndLineIntervalTable", types.YChild{"SonetFarEndLineIntervalTable", &sONETMIB.SonetFarEndLineIntervalTable})
    sONETMIB.EntityData.Children.Append("sonetPathCurrentTable", types.YChild{"SonetPathCurrentTable", &sONETMIB.SonetPathCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetPathIntervalTable", types.YChild{"SonetPathIntervalTable", &sONETMIB.SonetPathIntervalTable})
    sONETMIB.EntityData.Children.Append("sonetFarEndPathCurrentTable", types.YChild{"SonetFarEndPathCurrentTable", &sONETMIB.SonetFarEndPathCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetFarEndPathIntervalTable", types.YChild{"SonetFarEndPathIntervalTable", &sONETMIB.SonetFarEndPathIntervalTable})
    sONETMIB.EntityData.Children.Append("sonetVTCurrentTable", types.YChild{"SonetVTCurrentTable", &sONETMIB.SonetVTCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetVTIntervalTable", types.YChild{"SonetVTIntervalTable", &sONETMIB.SonetVTIntervalTable})
    sONETMIB.EntityData.Children.Append("sonetFarEndVTCurrentTable", types.YChild{"SonetFarEndVTCurrentTable", &sONETMIB.SonetFarEndVTCurrentTable})
    sONETMIB.EntityData.Children.Append("sonetFarEndVTIntervalTable", types.YChild{"SonetFarEndVTIntervalTable", &sONETMIB.SonetFarEndVTIntervalTable})
    sONETMIB.EntityData.Leafs = types.NewOrderedMap()

    sONETMIB.EntityData.YListKeys = []string {}

    return &(sONETMIB.EntityData)
}

// SONETMIB_SonetMedium
type SONETMIB_SonetMedium struct {
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
    // SonetSESthresholdSet.
    SonetSESthresholdSet interface{}
}

func (sonetMedium *SONETMIB_SonetMedium) GetEntityData() *types.CommonEntityData {
    sonetMedium.EntityData.YFilter = sonetMedium.YFilter
    sonetMedium.EntityData.YangName = "sonetMedium"
    sonetMedium.EntityData.BundleName = "cisco_ios_xe"
    sonetMedium.EntityData.ParentYangName = "SONET-MIB"
    sonetMedium.EntityData.SegmentPath = "sonetMedium"
    sonetMedium.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetMedium.EntityData.SegmentPath
    sonetMedium.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetMedium.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetMedium.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetMedium.EntityData.Children = types.NewOrderedMap()
    sonetMedium.EntityData.Leafs = types.NewOrderedMap()
    sonetMedium.EntityData.Leafs.Append("sonetSESthresholdSet", types.YLeaf{"SonetSESthresholdSet", sonetMedium.SonetSESthresholdSet})

    sonetMedium.EntityData.YListKeys = []string {}

    return &(sonetMedium.EntityData)
}

// SONETMIB_SonetMedium_SonetSESthresholdSet represents prior to this change must be invalidated.
type SONETMIB_SonetMedium_SonetSESthresholdSet string

const (
    SONETMIB_SonetMedium_SonetSESthresholdSet_other SONETMIB_SonetMedium_SonetSESthresholdSet = "other"

    SONETMIB_SonetMedium_SonetSESthresholdSet_bellcore1991 SONETMIB_SonetMedium_SonetSESthresholdSet = "bellcore1991"

    SONETMIB_SonetMedium_SonetSESthresholdSet_ansi1993 SONETMIB_SonetMedium_SonetSESthresholdSet = "ansi1993"

    SONETMIB_SonetMedium_SonetSESthresholdSet_itu1995 SONETMIB_SonetMedium_SonetSESthresholdSet = "itu1995"

    SONETMIB_SonetMedium_SonetSESthresholdSet_ansi1997 SONETMIB_SonetMedium_SonetSESthresholdSet = "ansi1997"
)

// SONETMIB_SonetMediumTable
// The SONET/SDH Medium table.
type SONETMIB_SonetMediumTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Medium table. The type is slice of
    // SONETMIB_SonetMediumTable_SonetMediumEntry.
    SonetMediumEntry []*SONETMIB_SonetMediumTable_SonetMediumEntry
}

func (sonetMediumTable *SONETMIB_SonetMediumTable) GetEntityData() *types.CommonEntityData {
    sonetMediumTable.EntityData.YFilter = sonetMediumTable.YFilter
    sonetMediumTable.EntityData.YangName = "sonetMediumTable"
    sonetMediumTable.EntityData.BundleName = "cisco_ios_xe"
    sonetMediumTable.EntityData.ParentYangName = "SONET-MIB"
    sonetMediumTable.EntityData.SegmentPath = "sonetMediumTable"
    sonetMediumTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetMediumTable.EntityData.SegmentPath
    sonetMediumTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetMediumTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetMediumTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetMediumTable.EntityData.Children = types.NewOrderedMap()
    sonetMediumTable.EntityData.Children.Append("sonetMediumEntry", types.YChild{"SonetMediumEntry", nil})
    for i := range sonetMediumTable.SonetMediumEntry {
        sonetMediumTable.EntityData.Children.Append(types.GetSegmentPath(sonetMediumTable.SonetMediumEntry[i]), types.YChild{"SonetMediumEntry", sonetMediumTable.SonetMediumEntry[i]})
    }
    sonetMediumTable.EntityData.Leafs = types.NewOrderedMap()

    sonetMediumTable.EntityData.YListKeys = []string {}

    return &(sonetMediumTable.EntityData)
}

// SONETMIB_SonetMediumTable_SonetMediumEntry
// An entry in the SONET/SDH Medium table.
type SONETMIB_SonetMediumTable_SonetMediumEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This variable identifies whether a SONET or a SDH signal is used across
    // this interface. The type is SonetMediumType.
    SonetMediumType interface{}

    // The number of seconds, including partial seconds, that have elapsed since
    // the beginning of the current measurement period. If, for some reason, such
    // as an adjustment in the system's time-of-day clock, the current interval
    // exceeds the maximum value, the agent will return the maximum value. The
    // type is interface{} with range: 1..900.
    SonetMediumTimeElapsed interface{}

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
    SonetMediumValidIntervals interface{}

    // This variable describes the line coding for this interface. The B3ZS and
    // CMI are used for electrical SONET/SDH signals (STS-1 and STS-3). The
    // Non-Return to Zero (NRZ) and the Return to Zero are used for optical
    // SONET/SDH signals. The type is SonetMediumLineCoding.
    SonetMediumLineCoding interface{}

    // This variable describes the line type for this interface. The line types
    // are Short and Long Range Single Mode fiber or Multi-Mode fiber interfaces,
    // and coax and UTP for electrical interfaces.  The value sonetOther should be
    // used when the Line Type is not one of the listed values. The type is
    // SonetMediumLineType.
    SonetMediumLineType interface{}

    // This variable contains the transmission vendor's circuit identifier, for
    // the purpose of facilitating troubleshooting. Note that the circuit
    // identifier, if available, is also represented by ifPhysAddress. The type is
    // string with length: 0..255.
    SonetMediumCircuitIdentifier interface{}

    // The number of intervals in the range from 0 to sonetMediumValidIntervals
    // for which no data is available. This object will typically be zero except
    // in cases where the data for some intervals are not available (e.g., in
    // proxy situations). The type is interface{} with range: 0..96.
    SonetMediumInvalidIntervals interface{}

    // The current loopback state of the SONET/SDH interface.  The values mean:   
    // sonetNoLoop      Not in the loopback state. A device that is not     
    // capable of performing a loopback on this interface      shall always return
    // this value.    sonetFacilityLoop      The received signal at this interface
    // is looped back      out through the corresponding transmitter in the return
    // direction.    sonetTerminalLoop      The signal that is about to be
    // transmitted is connected      to the associated incoming receiver.   
    // sonetOtherLoop      Loopbacks that are not defined here. The type is
    // map[string]bool.
    SonetMediumLoopbackConfig interface{}
}

func (sonetMediumEntry *SONETMIB_SonetMediumTable_SonetMediumEntry) GetEntityData() *types.CommonEntityData {
    sonetMediumEntry.EntityData.YFilter = sonetMediumEntry.YFilter
    sonetMediumEntry.EntityData.YangName = "sonetMediumEntry"
    sonetMediumEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetMediumEntry.EntityData.ParentYangName = "sonetMediumTable"
    sonetMediumEntry.EntityData.SegmentPath = "sonetMediumEntry" + types.AddKeyToken(sonetMediumEntry.IfIndex, "ifIndex")
    sonetMediumEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetMediumTable/" + sonetMediumEntry.EntityData.SegmentPath
    sonetMediumEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetMediumEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetMediumEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetMediumEntry.EntityData.Children = types.NewOrderedMap()
    sonetMediumEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetMediumEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetMediumEntry.IfIndex})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumType", types.YLeaf{"SonetMediumType", sonetMediumEntry.SonetMediumType})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumTimeElapsed", types.YLeaf{"SonetMediumTimeElapsed", sonetMediumEntry.SonetMediumTimeElapsed})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumValidIntervals", types.YLeaf{"SonetMediumValidIntervals", sonetMediumEntry.SonetMediumValidIntervals})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumLineCoding", types.YLeaf{"SonetMediumLineCoding", sonetMediumEntry.SonetMediumLineCoding})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumLineType", types.YLeaf{"SonetMediumLineType", sonetMediumEntry.SonetMediumLineType})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumCircuitIdentifier", types.YLeaf{"SonetMediumCircuitIdentifier", sonetMediumEntry.SonetMediumCircuitIdentifier})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumInvalidIntervals", types.YLeaf{"SonetMediumInvalidIntervals", sonetMediumEntry.SonetMediumInvalidIntervals})
    sonetMediumEntry.EntityData.Leafs.Append("sonetMediumLoopbackConfig", types.YLeaf{"SonetMediumLoopbackConfig", sonetMediumEntry.SonetMediumLoopbackConfig})

    sonetMediumEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetMediumEntry.EntityData)
}

// SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding represents to Zero are used for optical SONET/SDH signals.
type SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding string

const (
    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding_sonetMediumOther SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding = "sonetMediumOther"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding_sonetMediumB3ZS SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding = "sonetMediumB3ZS"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding_sonetMediumCMI SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding = "sonetMediumCMI"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding_sonetMediumNRZ SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding = "sonetMediumNRZ"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding_sonetMediumRZ SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineCoding = "sonetMediumRZ"
)

// SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType represents not one of the listed values.
type SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType string

const (
    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType_sonetOther SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType = "sonetOther"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType_sonetShortSingleMode SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType = "sonetShortSingleMode"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType_sonetLongSingleMode SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType = "sonetLongSingleMode"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType_sonetMultiMode SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType = "sonetMultiMode"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType_sonetCoax SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType = "sonetCoax"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType_sonetUTP SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumLineType = "sonetUTP"
)

// SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumType represents or a SDH signal is used across this interface.
type SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumType string

const (
    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumType_sonet SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumType = "sonet"

    SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumType_sdh SONETMIB_SonetMediumTable_SonetMediumEntry_SonetMediumType = "sdh"
)

// SONETMIB_SonetSectionCurrentTable
// The SONET/SDH Section Current table.
type SONETMIB_SonetSectionCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Current table. The type is slice of
    // SONETMIB_SonetSectionCurrentTable_SonetSectionCurrentEntry.
    SonetSectionCurrentEntry []*SONETMIB_SonetSectionCurrentTable_SonetSectionCurrentEntry
}

func (sonetSectionCurrentTable *SONETMIB_SonetSectionCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetSectionCurrentTable.EntityData.YFilter = sonetSectionCurrentTable.YFilter
    sonetSectionCurrentTable.EntityData.YangName = "sonetSectionCurrentTable"
    sonetSectionCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetSectionCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetSectionCurrentTable.EntityData.SegmentPath = "sonetSectionCurrentTable"
    sonetSectionCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetSectionCurrentTable.EntityData.SegmentPath
    sonetSectionCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetSectionCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetSectionCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetSectionCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetSectionCurrentTable.EntityData.Children.Append("sonetSectionCurrentEntry", types.YChild{"SonetSectionCurrentEntry", nil})
    for i := range sonetSectionCurrentTable.SonetSectionCurrentEntry {
        sonetSectionCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetSectionCurrentTable.SonetSectionCurrentEntry[i]), types.YChild{"SonetSectionCurrentEntry", sonetSectionCurrentTable.SonetSectionCurrentEntry[i]})
    }
    sonetSectionCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetSectionCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetSectionCurrentTable.EntityData)
}

// SONETMIB_SonetSectionCurrentTable_SonetSectionCurrentEntry
// An entry in the SONET/SDH Section Current table.
type SONETMIB_SonetSectionCurrentTable_SonetSectionCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This variable indicates the status of the interface. The
    // sonetSectionCurrentStatus is a bit map represented as a sum, therefore, it
    // can represent multiple defects simultaneously. The sonetSectionNoDefect
    // should be set if and only if no other flag is set.  The various bit
    // positions are:       1   sonetSectionNoDefect       2   sonetSectionLOS    
    // 4   sonetSectionLOF. The type is interface{} with range: 1..6.
    SonetSectionCurrentStatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Section in the current 15 minute interval. The type is
    // interface{} with range: 0..4294967295.
    SonetSectionCurrentESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Section in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    SonetSectionCurrentSESs interface{}

    // The counter associated with the number of Severely Errored Framing Seconds
    // encountered by a SONET/SDH Section in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    SonetSectionCurrentSEFSs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Section in the current 15 minute interval. The type is
    // interface{} with range: 0..4294967295.
    SonetSectionCurrentCVs interface{}
}

func (sonetSectionCurrentEntry *SONETMIB_SonetSectionCurrentTable_SonetSectionCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetSectionCurrentEntry.EntityData.YFilter = sonetSectionCurrentEntry.YFilter
    sonetSectionCurrentEntry.EntityData.YangName = "sonetSectionCurrentEntry"
    sonetSectionCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetSectionCurrentEntry.EntityData.ParentYangName = "sonetSectionCurrentTable"
    sonetSectionCurrentEntry.EntityData.SegmentPath = "sonetSectionCurrentEntry" + types.AddKeyToken(sonetSectionCurrentEntry.IfIndex, "ifIndex")
    sonetSectionCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetSectionCurrentTable/" + sonetSectionCurrentEntry.EntityData.SegmentPath
    sonetSectionCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetSectionCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetSectionCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetSectionCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetSectionCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetSectionCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetSectionCurrentEntry.IfIndex})
    sonetSectionCurrentEntry.EntityData.Leafs.Append("sonetSectionCurrentStatus", types.YLeaf{"SonetSectionCurrentStatus", sonetSectionCurrentEntry.SonetSectionCurrentStatus})
    sonetSectionCurrentEntry.EntityData.Leafs.Append("sonetSectionCurrentESs", types.YLeaf{"SonetSectionCurrentESs", sonetSectionCurrentEntry.SonetSectionCurrentESs})
    sonetSectionCurrentEntry.EntityData.Leafs.Append("sonetSectionCurrentSESs", types.YLeaf{"SonetSectionCurrentSESs", sonetSectionCurrentEntry.SonetSectionCurrentSESs})
    sonetSectionCurrentEntry.EntityData.Leafs.Append("sonetSectionCurrentSEFSs", types.YLeaf{"SonetSectionCurrentSEFSs", sonetSectionCurrentEntry.SonetSectionCurrentSEFSs})
    sonetSectionCurrentEntry.EntityData.Leafs.Append("sonetSectionCurrentCVs", types.YLeaf{"SonetSectionCurrentCVs", sonetSectionCurrentEntry.SonetSectionCurrentCVs})

    sonetSectionCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetSectionCurrentEntry.EntityData)
}

// SONETMIB_SonetSectionIntervalTable
// The SONET/SDH Section Interval table.
type SONETMIB_SonetSectionIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Section Interval table. The type is slice of
    // SONETMIB_SonetSectionIntervalTable_SonetSectionIntervalEntry.
    SonetSectionIntervalEntry []*SONETMIB_SonetSectionIntervalTable_SonetSectionIntervalEntry
}

func (sonetSectionIntervalTable *SONETMIB_SonetSectionIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetSectionIntervalTable.EntityData.YFilter = sonetSectionIntervalTable.YFilter
    sonetSectionIntervalTable.EntityData.YangName = "sonetSectionIntervalTable"
    sonetSectionIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetSectionIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetSectionIntervalTable.EntityData.SegmentPath = "sonetSectionIntervalTable"
    sonetSectionIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetSectionIntervalTable.EntityData.SegmentPath
    sonetSectionIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetSectionIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetSectionIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetSectionIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetSectionIntervalTable.EntityData.Children.Append("sonetSectionIntervalEntry", types.YChild{"SonetSectionIntervalEntry", nil})
    for i := range sonetSectionIntervalTable.SonetSectionIntervalEntry {
        sonetSectionIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetSectionIntervalTable.SonetSectionIntervalEntry[i]), types.YChild{"SonetSectionIntervalEntry", sonetSectionIntervalTable.SonetSectionIntervalEntry[i]})
    }
    sonetSectionIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetSectionIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetSectionIntervalTable.EntityData)
}

// SONETMIB_SonetSectionIntervalTable_SonetSectionIntervalEntry
// An entry in the SONET/SDH Section Interval table.
type SONETMIB_SonetSectionIntervalTable_SonetSectionIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetSectionIntervalNumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Section in a particular 15-minute interval in the past 24 hours.
    // The type is interface{} with range: 0..4294967295.
    SonetSectionIntervalESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Section in a particular 15-minute interval in
    // the past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetSectionIntervalSESs interface{}

    // The counter associated with the number of Severely Errored Framing Seconds
    // encountered by a SONET/SDH Section in a particular 15-minute interval in
    // the past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetSectionIntervalSEFSs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Section in a particular 15-minute interval in the past 24
    // hours. The type is interface{} with range: 0..4294967295.
    SonetSectionIntervalCVs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetSectionIntervalValidData interface{}
}

func (sonetSectionIntervalEntry *SONETMIB_SonetSectionIntervalTable_SonetSectionIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetSectionIntervalEntry.EntityData.YFilter = sonetSectionIntervalEntry.YFilter
    sonetSectionIntervalEntry.EntityData.YangName = "sonetSectionIntervalEntry"
    sonetSectionIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetSectionIntervalEntry.EntityData.ParentYangName = "sonetSectionIntervalTable"
    sonetSectionIntervalEntry.EntityData.SegmentPath = "sonetSectionIntervalEntry" + types.AddKeyToken(sonetSectionIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetSectionIntervalEntry.SonetSectionIntervalNumber, "sonetSectionIntervalNumber")
    sonetSectionIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetSectionIntervalTable/" + sonetSectionIntervalEntry.EntityData.SegmentPath
    sonetSectionIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetSectionIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetSectionIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetSectionIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetSectionIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetSectionIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetSectionIntervalEntry.IfIndex})
    sonetSectionIntervalEntry.EntityData.Leafs.Append("sonetSectionIntervalNumber", types.YLeaf{"SonetSectionIntervalNumber", sonetSectionIntervalEntry.SonetSectionIntervalNumber})
    sonetSectionIntervalEntry.EntityData.Leafs.Append("sonetSectionIntervalESs", types.YLeaf{"SonetSectionIntervalESs", sonetSectionIntervalEntry.SonetSectionIntervalESs})
    sonetSectionIntervalEntry.EntityData.Leafs.Append("sonetSectionIntervalSESs", types.YLeaf{"SonetSectionIntervalSESs", sonetSectionIntervalEntry.SonetSectionIntervalSESs})
    sonetSectionIntervalEntry.EntityData.Leafs.Append("sonetSectionIntervalSEFSs", types.YLeaf{"SonetSectionIntervalSEFSs", sonetSectionIntervalEntry.SonetSectionIntervalSEFSs})
    sonetSectionIntervalEntry.EntityData.Leafs.Append("sonetSectionIntervalCVs", types.YLeaf{"SonetSectionIntervalCVs", sonetSectionIntervalEntry.SonetSectionIntervalCVs})
    sonetSectionIntervalEntry.EntityData.Leafs.Append("sonetSectionIntervalValidData", types.YLeaf{"SonetSectionIntervalValidData", sonetSectionIntervalEntry.SonetSectionIntervalValidData})

    sonetSectionIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetSectionIntervalNumber"}

    return &(sonetSectionIntervalEntry.EntityData)
}

// SONETMIB_SonetLineCurrentTable
// The SONET/SDH Line Current table.
type SONETMIB_SonetLineCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Current table. The type is slice of
    // SONETMIB_SonetLineCurrentTable_SonetLineCurrentEntry.
    SonetLineCurrentEntry []*SONETMIB_SonetLineCurrentTable_SonetLineCurrentEntry
}

func (sonetLineCurrentTable *SONETMIB_SonetLineCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetLineCurrentTable.EntityData.YFilter = sonetLineCurrentTable.YFilter
    sonetLineCurrentTable.EntityData.YangName = "sonetLineCurrentTable"
    sonetLineCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetLineCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetLineCurrentTable.EntityData.SegmentPath = "sonetLineCurrentTable"
    sonetLineCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetLineCurrentTable.EntityData.SegmentPath
    sonetLineCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetLineCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetLineCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetLineCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetLineCurrentTable.EntityData.Children.Append("sonetLineCurrentEntry", types.YChild{"SonetLineCurrentEntry", nil})
    for i := range sonetLineCurrentTable.SonetLineCurrentEntry {
        sonetLineCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetLineCurrentTable.SonetLineCurrentEntry[i]), types.YChild{"SonetLineCurrentEntry", sonetLineCurrentTable.SonetLineCurrentEntry[i]})
    }
    sonetLineCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetLineCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetLineCurrentTable.EntityData)
}

// SONETMIB_SonetLineCurrentTable_SonetLineCurrentEntry
// An entry in the SONET/SDH Line Current table.
type SONETMIB_SonetLineCurrentTable_SonetLineCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This variable indicates the status of the interface. The
    // sonetLineCurrentStatus is a bit map represented as a sum, therefore, it can
    // represent multiple defects simultaneously. The sonetLineNoDefect should be
    // set if and only if no other flag is set.  The various bit positions are:  1
    // sonetLineNoDefect  2   sonetLineAIS  4   sonetLineRDI. The type is
    // interface{} with range: 1..6.
    SonetLineCurrentStatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Line in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    SonetLineCurrentESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Line in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    SonetLineCurrentSESs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Line in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    SonetLineCurrentCVs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a SONET/SDH Line in the current 15 minute interval. The type is
    // interface{} with range: 0..4294967295.
    SonetLineCurrentUASs interface{}
}

func (sonetLineCurrentEntry *SONETMIB_SonetLineCurrentTable_SonetLineCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetLineCurrentEntry.EntityData.YFilter = sonetLineCurrentEntry.YFilter
    sonetLineCurrentEntry.EntityData.YangName = "sonetLineCurrentEntry"
    sonetLineCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetLineCurrentEntry.EntityData.ParentYangName = "sonetLineCurrentTable"
    sonetLineCurrentEntry.EntityData.SegmentPath = "sonetLineCurrentEntry" + types.AddKeyToken(sonetLineCurrentEntry.IfIndex, "ifIndex")
    sonetLineCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetLineCurrentTable/" + sonetLineCurrentEntry.EntityData.SegmentPath
    sonetLineCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetLineCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetLineCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetLineCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetLineCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetLineCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetLineCurrentEntry.IfIndex})
    sonetLineCurrentEntry.EntityData.Leafs.Append("sonetLineCurrentStatus", types.YLeaf{"SonetLineCurrentStatus", sonetLineCurrentEntry.SonetLineCurrentStatus})
    sonetLineCurrentEntry.EntityData.Leafs.Append("sonetLineCurrentESs", types.YLeaf{"SonetLineCurrentESs", sonetLineCurrentEntry.SonetLineCurrentESs})
    sonetLineCurrentEntry.EntityData.Leafs.Append("sonetLineCurrentSESs", types.YLeaf{"SonetLineCurrentSESs", sonetLineCurrentEntry.SonetLineCurrentSESs})
    sonetLineCurrentEntry.EntityData.Leafs.Append("sonetLineCurrentCVs", types.YLeaf{"SonetLineCurrentCVs", sonetLineCurrentEntry.SonetLineCurrentCVs})
    sonetLineCurrentEntry.EntityData.Leafs.Append("sonetLineCurrentUASs", types.YLeaf{"SonetLineCurrentUASs", sonetLineCurrentEntry.SonetLineCurrentUASs})

    sonetLineCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetLineCurrentEntry.EntityData)
}

// SONETMIB_SonetLineIntervalTable
// The SONET/SDH Line Interval table.
type SONETMIB_SonetLineIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Line Interval table. The type is slice of
    // SONETMIB_SonetLineIntervalTable_SonetLineIntervalEntry.
    SonetLineIntervalEntry []*SONETMIB_SonetLineIntervalTable_SonetLineIntervalEntry
}

func (sonetLineIntervalTable *SONETMIB_SonetLineIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetLineIntervalTable.EntityData.YFilter = sonetLineIntervalTable.YFilter
    sonetLineIntervalTable.EntityData.YangName = "sonetLineIntervalTable"
    sonetLineIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetLineIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetLineIntervalTable.EntityData.SegmentPath = "sonetLineIntervalTable"
    sonetLineIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetLineIntervalTable.EntityData.SegmentPath
    sonetLineIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetLineIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetLineIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetLineIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetLineIntervalTable.EntityData.Children.Append("sonetLineIntervalEntry", types.YChild{"SonetLineIntervalEntry", nil})
    for i := range sonetLineIntervalTable.SonetLineIntervalEntry {
        sonetLineIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetLineIntervalTable.SonetLineIntervalEntry[i]), types.YChild{"SonetLineIntervalEntry", sonetLineIntervalTable.SonetLineIntervalEntry[i]})
    }
    sonetLineIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetLineIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetLineIntervalTable.EntityData)
}

// SONETMIB_SonetLineIntervalTable_SonetLineIntervalEntry
// An entry in the SONET/SDH Line Interval table.
type SONETMIB_SonetLineIntervalTable_SonetLineIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetLineIntervalNumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Line in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    SonetLineIntervalESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Line in a particular 15-minute interval in the
    // past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetLineIntervalSESs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Line in a particular 15-minute interval in the past 24 hours.
    // The type is interface{} with range: 0..4294967295.
    SonetLineIntervalCVs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a SONET/SDH Line in a particular 15-minute interval in the past 24
    // hours. The type is interface{} with range: 0..4294967295.
    SonetLineIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetLineIntervalValidData interface{}
}

func (sonetLineIntervalEntry *SONETMIB_SonetLineIntervalTable_SonetLineIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetLineIntervalEntry.EntityData.YFilter = sonetLineIntervalEntry.YFilter
    sonetLineIntervalEntry.EntityData.YangName = "sonetLineIntervalEntry"
    sonetLineIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetLineIntervalEntry.EntityData.ParentYangName = "sonetLineIntervalTable"
    sonetLineIntervalEntry.EntityData.SegmentPath = "sonetLineIntervalEntry" + types.AddKeyToken(sonetLineIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetLineIntervalEntry.SonetLineIntervalNumber, "sonetLineIntervalNumber")
    sonetLineIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetLineIntervalTable/" + sonetLineIntervalEntry.EntityData.SegmentPath
    sonetLineIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetLineIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetLineIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetLineIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetLineIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetLineIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetLineIntervalEntry.IfIndex})
    sonetLineIntervalEntry.EntityData.Leafs.Append("sonetLineIntervalNumber", types.YLeaf{"SonetLineIntervalNumber", sonetLineIntervalEntry.SonetLineIntervalNumber})
    sonetLineIntervalEntry.EntityData.Leafs.Append("sonetLineIntervalESs", types.YLeaf{"SonetLineIntervalESs", sonetLineIntervalEntry.SonetLineIntervalESs})
    sonetLineIntervalEntry.EntityData.Leafs.Append("sonetLineIntervalSESs", types.YLeaf{"SonetLineIntervalSESs", sonetLineIntervalEntry.SonetLineIntervalSESs})
    sonetLineIntervalEntry.EntityData.Leafs.Append("sonetLineIntervalCVs", types.YLeaf{"SonetLineIntervalCVs", sonetLineIntervalEntry.SonetLineIntervalCVs})
    sonetLineIntervalEntry.EntityData.Leafs.Append("sonetLineIntervalUASs", types.YLeaf{"SonetLineIntervalUASs", sonetLineIntervalEntry.SonetLineIntervalUASs})
    sonetLineIntervalEntry.EntityData.Leafs.Append("sonetLineIntervalValidData", types.YLeaf{"SonetLineIntervalValidData", sonetLineIntervalEntry.SonetLineIntervalValidData})

    sonetLineIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetLineIntervalNumber"}

    return &(sonetLineIntervalEntry.EntityData)
}

// SONETMIB_SonetFarEndLineCurrentTable
// The SONET/SDH Far End Line Current table.
type SONETMIB_SonetFarEndLineCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Current table. The type is slice of
    // SONETMIB_SonetFarEndLineCurrentTable_SonetFarEndLineCurrentEntry.
    SonetFarEndLineCurrentEntry []*SONETMIB_SonetFarEndLineCurrentTable_SonetFarEndLineCurrentEntry
}

func (sonetFarEndLineCurrentTable *SONETMIB_SonetFarEndLineCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetFarEndLineCurrentTable.EntityData.YFilter = sonetFarEndLineCurrentTable.YFilter
    sonetFarEndLineCurrentTable.EntityData.YangName = "sonetFarEndLineCurrentTable"
    sonetFarEndLineCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndLineCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetFarEndLineCurrentTable.EntityData.SegmentPath = "sonetFarEndLineCurrentTable"
    sonetFarEndLineCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetFarEndLineCurrentTable.EntityData.SegmentPath
    sonetFarEndLineCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndLineCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndLineCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndLineCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetFarEndLineCurrentTable.EntityData.Children.Append("sonetFarEndLineCurrentEntry", types.YChild{"SonetFarEndLineCurrentEntry", nil})
    for i := range sonetFarEndLineCurrentTable.SonetFarEndLineCurrentEntry {
        sonetFarEndLineCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetFarEndLineCurrentTable.SonetFarEndLineCurrentEntry[i]), types.YChild{"SonetFarEndLineCurrentEntry", sonetFarEndLineCurrentTable.SonetFarEndLineCurrentEntry[i]})
    }
    sonetFarEndLineCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetFarEndLineCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetFarEndLineCurrentTable.EntityData)
}

// SONETMIB_SonetFarEndLineCurrentTable_SonetFarEndLineCurrentEntry
// An entry in the SONET/SDH Far End Line Current table.
type SONETMIB_SonetFarEndLineCurrentTable_SonetFarEndLineCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH interface in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    SonetFarEndLineCurrentESs interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Medium/Section/Line interface in the current 15
    // minute interval. The type is interface{} with range: 0..4294967295.
    SonetFarEndLineCurrentSESs interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH
    // Medium/Section/Line interface in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    SonetFarEndLineCurrentCVs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Medium/Section/Line interface in the current 15
    // minute interval. The type is interface{} with range: 0..4294967295.
    SonetFarEndLineCurrentUASs interface{}
}

func (sonetFarEndLineCurrentEntry *SONETMIB_SonetFarEndLineCurrentTable_SonetFarEndLineCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetFarEndLineCurrentEntry.EntityData.YFilter = sonetFarEndLineCurrentEntry.YFilter
    sonetFarEndLineCurrentEntry.EntityData.YangName = "sonetFarEndLineCurrentEntry"
    sonetFarEndLineCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndLineCurrentEntry.EntityData.ParentYangName = "sonetFarEndLineCurrentTable"
    sonetFarEndLineCurrentEntry.EntityData.SegmentPath = "sonetFarEndLineCurrentEntry" + types.AddKeyToken(sonetFarEndLineCurrentEntry.IfIndex, "ifIndex")
    sonetFarEndLineCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetFarEndLineCurrentTable/" + sonetFarEndLineCurrentEntry.EntityData.SegmentPath
    sonetFarEndLineCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndLineCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndLineCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndLineCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetFarEndLineCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetFarEndLineCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetFarEndLineCurrentEntry.IfIndex})
    sonetFarEndLineCurrentEntry.EntityData.Leafs.Append("sonetFarEndLineCurrentESs", types.YLeaf{"SonetFarEndLineCurrentESs", sonetFarEndLineCurrentEntry.SonetFarEndLineCurrentESs})
    sonetFarEndLineCurrentEntry.EntityData.Leafs.Append("sonetFarEndLineCurrentSESs", types.YLeaf{"SonetFarEndLineCurrentSESs", sonetFarEndLineCurrentEntry.SonetFarEndLineCurrentSESs})
    sonetFarEndLineCurrentEntry.EntityData.Leafs.Append("sonetFarEndLineCurrentCVs", types.YLeaf{"SonetFarEndLineCurrentCVs", sonetFarEndLineCurrentEntry.SonetFarEndLineCurrentCVs})
    sonetFarEndLineCurrentEntry.EntityData.Leafs.Append("sonetFarEndLineCurrentUASs", types.YLeaf{"SonetFarEndLineCurrentUASs", sonetFarEndLineCurrentEntry.SonetFarEndLineCurrentUASs})

    sonetFarEndLineCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetFarEndLineCurrentEntry.EntityData)
}

// SONETMIB_SonetFarEndLineIntervalTable
// The SONET/SDH Far End Line Interval table.
type SONETMIB_SonetFarEndLineIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Line Interval table. The type is slice of
    // SONETMIB_SonetFarEndLineIntervalTable_SonetFarEndLineIntervalEntry.
    SonetFarEndLineIntervalEntry []*SONETMIB_SonetFarEndLineIntervalTable_SonetFarEndLineIntervalEntry
}

func (sonetFarEndLineIntervalTable *SONETMIB_SonetFarEndLineIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetFarEndLineIntervalTable.EntityData.YFilter = sonetFarEndLineIntervalTable.YFilter
    sonetFarEndLineIntervalTable.EntityData.YangName = "sonetFarEndLineIntervalTable"
    sonetFarEndLineIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndLineIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetFarEndLineIntervalTable.EntityData.SegmentPath = "sonetFarEndLineIntervalTable"
    sonetFarEndLineIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetFarEndLineIntervalTable.EntityData.SegmentPath
    sonetFarEndLineIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndLineIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndLineIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndLineIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetFarEndLineIntervalTable.EntityData.Children.Append("sonetFarEndLineIntervalEntry", types.YChild{"SonetFarEndLineIntervalEntry", nil})
    for i := range sonetFarEndLineIntervalTable.SonetFarEndLineIntervalEntry {
        sonetFarEndLineIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetFarEndLineIntervalTable.SonetFarEndLineIntervalEntry[i]), types.YChild{"SonetFarEndLineIntervalEntry", sonetFarEndLineIntervalTable.SonetFarEndLineIntervalEntry[i]})
    }
    sonetFarEndLineIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetFarEndLineIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetFarEndLineIntervalTable.EntityData)
}

// SONETMIB_SonetFarEndLineIntervalTable_SonetFarEndLineIntervalEntry
// An entry in the SONET/SDH Far
// End Line Interval table.
type SONETMIB_SonetFarEndLineIntervalTable_SonetFarEndLineIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetFarEndLineIntervalNumber interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH Line interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    SonetFarEndLineIntervalESs interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Line interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    SonetFarEndLineIntervalSESs interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH Line
    // interface in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    SonetFarEndLineIntervalCVs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Line interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    SonetFarEndLineIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetFarEndLineIntervalValidData interface{}
}

func (sonetFarEndLineIntervalEntry *SONETMIB_SonetFarEndLineIntervalTable_SonetFarEndLineIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetFarEndLineIntervalEntry.EntityData.YFilter = sonetFarEndLineIntervalEntry.YFilter
    sonetFarEndLineIntervalEntry.EntityData.YangName = "sonetFarEndLineIntervalEntry"
    sonetFarEndLineIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndLineIntervalEntry.EntityData.ParentYangName = "sonetFarEndLineIntervalTable"
    sonetFarEndLineIntervalEntry.EntityData.SegmentPath = "sonetFarEndLineIntervalEntry" + types.AddKeyToken(sonetFarEndLineIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalNumber, "sonetFarEndLineIntervalNumber")
    sonetFarEndLineIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetFarEndLineIntervalTable/" + sonetFarEndLineIntervalEntry.EntityData.SegmentPath
    sonetFarEndLineIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndLineIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndLineIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndLineIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetFarEndLineIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetFarEndLineIntervalEntry.IfIndex})
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("sonetFarEndLineIntervalNumber", types.YLeaf{"SonetFarEndLineIntervalNumber", sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalNumber})
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("sonetFarEndLineIntervalESs", types.YLeaf{"SonetFarEndLineIntervalESs", sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalESs})
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("sonetFarEndLineIntervalSESs", types.YLeaf{"SonetFarEndLineIntervalSESs", sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalSESs})
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("sonetFarEndLineIntervalCVs", types.YLeaf{"SonetFarEndLineIntervalCVs", sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalCVs})
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("sonetFarEndLineIntervalUASs", types.YLeaf{"SonetFarEndLineIntervalUASs", sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalUASs})
    sonetFarEndLineIntervalEntry.EntityData.Leafs.Append("sonetFarEndLineIntervalValidData", types.YLeaf{"SonetFarEndLineIntervalValidData", sonetFarEndLineIntervalEntry.SonetFarEndLineIntervalValidData})

    sonetFarEndLineIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetFarEndLineIntervalNumber"}

    return &(sonetFarEndLineIntervalEntry.EntityData)
}

// SONETMIB_SonetPathCurrentTable
// The SONET/SDH Path Current table.
type SONETMIB_SonetPathCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Current table. The type is slice of
    // SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry.
    SonetPathCurrentEntry []*SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry
}

func (sonetPathCurrentTable *SONETMIB_SonetPathCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetPathCurrentTable.EntityData.YFilter = sonetPathCurrentTable.YFilter
    sonetPathCurrentTable.EntityData.YangName = "sonetPathCurrentTable"
    sonetPathCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetPathCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetPathCurrentTable.EntityData.SegmentPath = "sonetPathCurrentTable"
    sonetPathCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetPathCurrentTable.EntityData.SegmentPath
    sonetPathCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetPathCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetPathCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetPathCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetPathCurrentTable.EntityData.Children.Append("sonetPathCurrentEntry", types.YChild{"SonetPathCurrentEntry", nil})
    for i := range sonetPathCurrentTable.SonetPathCurrentEntry {
        sonetPathCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetPathCurrentTable.SonetPathCurrentEntry[i]), types.YChild{"SonetPathCurrentEntry", sonetPathCurrentTable.SonetPathCurrentEntry[i]})
    }
    sonetPathCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetPathCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetPathCurrentTable.EntityData)
}

// SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry
// An entry in the SONET/SDH Path Current table.
type SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // A value that indicates the type of the SONET/SDH Path.  For SONET, the
    // assigned types are the STS-Nc SPEs, where N = 1, 3, 12, 24, 48, 192 and
    // 768. STS-1 is equal to 51.84 Mbps.  For SDH, the assigned types are the
    // STM-Nc VCs, where N = 1, 4, 16, 64 and 256. The type is
    // SonetPathCurrentWidth.
    SonetPathCurrentWidth interface{}

    // This variable indicates the status of the interface. The
    // sonetPathCurrentStatus is a bit map represented as a sum, therefore, it can
    // represent multiple defects simultaneously. The sonetPathNoDefect should be
    // set if and only if no other flag is set.  The various bit positions are:   
    // 1   sonetPathNoDefect    2   sonetPathSTSLOP    4   sonetPathSTSAIS    8  
    // sonetPathSTSRDI   16   sonetPathUnequipped   32  
    // sonetPathSignalLabelMismatch. The type is interface{} with range: 1..62.
    SonetPathCurrentStatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Path in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    SonetPathCurrentESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Path in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    SonetPathCurrentSESs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Path in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    SonetPathCurrentCVs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a Path in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    SonetPathCurrentUASs interface{}

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
    // CspSonetPathPayload.
    CspSonetPathPayload interface{}

    // This object represents the VT/VC mapping type. asynchronous: In this mode,
    // the channel structure of                DS1/E1 is neither visible nor
    // preserved.  byteSynchronous: In this mode, the DS0 signals inside          
    // the VT/VC can be found and extracted                   from the frame. The
    // initial value is asynchronous(1). The type is CspTributaryMappingType.
    CspTributaryMappingType interface{}

    // This object represents the mode used to transport DS0  Signalling
    // information for T1 byteSynchronous mapping (GR253). In
    // signallingTransferMode(2), the robbed-bit signalling  is transferred to the
    // VT header. In clearMode(3), only  the framing bit is transferred to the VT
    // header.           The initial value is signallingTransferMode(2)  if
    // csTributaryMappingType is byteSynchronous.  For asynchronous mapping, it is
    // notApplicable(1).  The value notApplicable(1) can not be set. The type is
    // CspSignallingTransportMode.
    CspSignallingTransportMode interface{}

    // This object represents the method used to group VCs into an STM-1 signal.
    // Applicable only to SDH.  au3Grouping: STM1<-AU-3<-TUG-2<-TU-12<-VC12 or    
    // STM1<-AU-3<-TUG-2<-TU-11<-VC11.  au4Grouping:
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-12<-VC12 or             
    // STM1<-AU-4<-TUG-3<-TUG-2<-TU-11<-VC11.  The initial value is au3Grouping(2)
    // for SDH and  notApplicable(1) for SONET. The type is
    // CspTributaryGroupingType.
    CspTributaryGroupingType interface{}
}

func (sonetPathCurrentEntry *SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetPathCurrentEntry.EntityData.YFilter = sonetPathCurrentEntry.YFilter
    sonetPathCurrentEntry.EntityData.YangName = "sonetPathCurrentEntry"
    sonetPathCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetPathCurrentEntry.EntityData.ParentYangName = "sonetPathCurrentTable"
    sonetPathCurrentEntry.EntityData.SegmentPath = "sonetPathCurrentEntry" + types.AddKeyToken(sonetPathCurrentEntry.IfIndex, "ifIndex")
    sonetPathCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetPathCurrentTable/" + sonetPathCurrentEntry.EntityData.SegmentPath
    sonetPathCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetPathCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetPathCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetPathCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetPathCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetPathCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetPathCurrentEntry.IfIndex})
    sonetPathCurrentEntry.EntityData.Leafs.Append("sonetPathCurrentWidth", types.YLeaf{"SonetPathCurrentWidth", sonetPathCurrentEntry.SonetPathCurrentWidth})
    sonetPathCurrentEntry.EntityData.Leafs.Append("sonetPathCurrentStatus", types.YLeaf{"SonetPathCurrentStatus", sonetPathCurrentEntry.SonetPathCurrentStatus})
    sonetPathCurrentEntry.EntityData.Leafs.Append("sonetPathCurrentESs", types.YLeaf{"SonetPathCurrentESs", sonetPathCurrentEntry.SonetPathCurrentESs})
    sonetPathCurrentEntry.EntityData.Leafs.Append("sonetPathCurrentSESs", types.YLeaf{"SonetPathCurrentSESs", sonetPathCurrentEntry.SonetPathCurrentSESs})
    sonetPathCurrentEntry.EntityData.Leafs.Append("sonetPathCurrentCVs", types.YLeaf{"SonetPathCurrentCVs", sonetPathCurrentEntry.SonetPathCurrentCVs})
    sonetPathCurrentEntry.EntityData.Leafs.Append("sonetPathCurrentUASs", types.YLeaf{"SonetPathCurrentUASs", sonetPathCurrentEntry.SonetPathCurrentUASs})
    sonetPathCurrentEntry.EntityData.Leafs.Append("cspSonetPathPayload", types.YLeaf{"CspSonetPathPayload", sonetPathCurrentEntry.CspSonetPathPayload})
    sonetPathCurrentEntry.EntityData.Leafs.Append("cspTributaryMappingType", types.YLeaf{"CspTributaryMappingType", sonetPathCurrentEntry.CspTributaryMappingType})
    sonetPathCurrentEntry.EntityData.Leafs.Append("cspSignallingTransportMode", types.YLeaf{"CspSignallingTransportMode", sonetPathCurrentEntry.CspSignallingTransportMode})
    sonetPathCurrentEntry.EntityData.Leafs.Append("cspTributaryGroupingType", types.YLeaf{"CspTributaryGroupingType", sonetPathCurrentEntry.CspTributaryGroupingType})

    sonetPathCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetPathCurrentEntry.EntityData)
}

// SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode represents The value notApplicable(1) can not be set.
type SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode string

const (
    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode_notApplicable SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode = "notApplicable"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode_signallingTransferMode SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode = "signallingTransferMode"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode_clearMode SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSignallingTransportMode = "clearMode"
)

// SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload represents C2 overhead byte in the Path Overhead.
type SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload string

const (
    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_unequipped SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "unequipped"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_unspecified SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "unspecified"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_ds3 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "ds3"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_vt15vc11 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "vt15vc11"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_vt2vc12 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "vt2vc12"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_atmCell SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "atmCell"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_hdlcFr SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "hdlcFr"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_e3 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "e3"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload_vtStructured SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspSonetPathPayload = "vtStructured"
)

// SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType represents notApplicable(1) for SONET.
type SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType string

const (
    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType_notApplicable SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType = "notApplicable"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType_au3Grouping SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType = "au3Grouping"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType_au4Grouping SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryGroupingType = "au4Grouping"
)

// SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryMappingType represents The initial value is asynchronous(1).
type SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryMappingType string

const (
    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryMappingType_asynchronous SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryMappingType = "asynchronous"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryMappingType_byteSynchronous SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_CspTributaryMappingType = "byteSynchronous"
)

// SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth represents types are the STM-Nc VCs, where N = 1, 4, 16, 64 and 256.
type SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth string

const (
    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts1 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts1"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts3cSTM1 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts3cSTM1"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts12cSTM4 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts12cSTM4"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts24c SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts24c"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts48cSTM16 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts48cSTM16"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts192cSTM64 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts192cSTM64"

    SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth_sts768cSTM256 SONETMIB_SonetPathCurrentTable_SonetPathCurrentEntry_SonetPathCurrentWidth = "sts768cSTM256"
)

// SONETMIB_SonetPathIntervalTable
// The SONET/SDH Path Interval table.
type SONETMIB_SonetPathIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Path Interval table. The type is slice of
    // SONETMIB_SonetPathIntervalTable_SonetPathIntervalEntry.
    SonetPathIntervalEntry []*SONETMIB_SonetPathIntervalTable_SonetPathIntervalEntry
}

func (sonetPathIntervalTable *SONETMIB_SonetPathIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetPathIntervalTable.EntityData.YFilter = sonetPathIntervalTable.YFilter
    sonetPathIntervalTable.EntityData.YangName = "sonetPathIntervalTable"
    sonetPathIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetPathIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetPathIntervalTable.EntityData.SegmentPath = "sonetPathIntervalTable"
    sonetPathIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetPathIntervalTable.EntityData.SegmentPath
    sonetPathIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetPathIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetPathIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetPathIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetPathIntervalTable.EntityData.Children.Append("sonetPathIntervalEntry", types.YChild{"SonetPathIntervalEntry", nil})
    for i := range sonetPathIntervalTable.SonetPathIntervalEntry {
        sonetPathIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetPathIntervalTable.SonetPathIntervalEntry[i]), types.YChild{"SonetPathIntervalEntry", sonetPathIntervalTable.SonetPathIntervalEntry[i]})
    }
    sonetPathIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetPathIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetPathIntervalTable.EntityData)
}

// SONETMIB_SonetPathIntervalTable_SonetPathIntervalEntry
// An entry in the SONET/SDH Path Interval table.
type SONETMIB_SonetPathIntervalTable_SonetPathIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetPathIntervalNumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH Path in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    SonetPathIntervalESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH Path in a particular 15-minute interval in the
    // past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetPathIntervalSESs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH Path in a particular 15-minute interval in the past 24 hours.
    // The type is interface{} with range: 0..4294967295.
    SonetPathIntervalCVs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a Path in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    SonetPathIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetPathIntervalValidData interface{}
}

func (sonetPathIntervalEntry *SONETMIB_SonetPathIntervalTable_SonetPathIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetPathIntervalEntry.EntityData.YFilter = sonetPathIntervalEntry.YFilter
    sonetPathIntervalEntry.EntityData.YangName = "sonetPathIntervalEntry"
    sonetPathIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetPathIntervalEntry.EntityData.ParentYangName = "sonetPathIntervalTable"
    sonetPathIntervalEntry.EntityData.SegmentPath = "sonetPathIntervalEntry" + types.AddKeyToken(sonetPathIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetPathIntervalEntry.SonetPathIntervalNumber, "sonetPathIntervalNumber")
    sonetPathIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetPathIntervalTable/" + sonetPathIntervalEntry.EntityData.SegmentPath
    sonetPathIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetPathIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetPathIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetPathIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetPathIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetPathIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetPathIntervalEntry.IfIndex})
    sonetPathIntervalEntry.EntityData.Leafs.Append("sonetPathIntervalNumber", types.YLeaf{"SonetPathIntervalNumber", sonetPathIntervalEntry.SonetPathIntervalNumber})
    sonetPathIntervalEntry.EntityData.Leafs.Append("sonetPathIntervalESs", types.YLeaf{"SonetPathIntervalESs", sonetPathIntervalEntry.SonetPathIntervalESs})
    sonetPathIntervalEntry.EntityData.Leafs.Append("sonetPathIntervalSESs", types.YLeaf{"SonetPathIntervalSESs", sonetPathIntervalEntry.SonetPathIntervalSESs})
    sonetPathIntervalEntry.EntityData.Leafs.Append("sonetPathIntervalCVs", types.YLeaf{"SonetPathIntervalCVs", sonetPathIntervalEntry.SonetPathIntervalCVs})
    sonetPathIntervalEntry.EntityData.Leafs.Append("sonetPathIntervalUASs", types.YLeaf{"SonetPathIntervalUASs", sonetPathIntervalEntry.SonetPathIntervalUASs})
    sonetPathIntervalEntry.EntityData.Leafs.Append("sonetPathIntervalValidData", types.YLeaf{"SonetPathIntervalValidData", sonetPathIntervalEntry.SonetPathIntervalValidData})

    sonetPathIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetPathIntervalNumber"}

    return &(sonetPathIntervalEntry.EntityData)
}

// SONETMIB_SonetFarEndPathCurrentTable
// The SONET/SDH Far End Path Current table.
type SONETMIB_SonetFarEndPathCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Current table. The type is slice of
    // SONETMIB_SonetFarEndPathCurrentTable_SonetFarEndPathCurrentEntry.
    SonetFarEndPathCurrentEntry []*SONETMIB_SonetFarEndPathCurrentTable_SonetFarEndPathCurrentEntry
}

func (sonetFarEndPathCurrentTable *SONETMIB_SonetFarEndPathCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetFarEndPathCurrentTable.EntityData.YFilter = sonetFarEndPathCurrentTable.YFilter
    sonetFarEndPathCurrentTable.EntityData.YangName = "sonetFarEndPathCurrentTable"
    sonetFarEndPathCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndPathCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetFarEndPathCurrentTable.EntityData.SegmentPath = "sonetFarEndPathCurrentTable"
    sonetFarEndPathCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetFarEndPathCurrentTable.EntityData.SegmentPath
    sonetFarEndPathCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndPathCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndPathCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndPathCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetFarEndPathCurrentTable.EntityData.Children.Append("sonetFarEndPathCurrentEntry", types.YChild{"SonetFarEndPathCurrentEntry", nil})
    for i := range sonetFarEndPathCurrentTable.SonetFarEndPathCurrentEntry {
        sonetFarEndPathCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetFarEndPathCurrentTable.SonetFarEndPathCurrentEntry[i]), types.YChild{"SonetFarEndPathCurrentEntry", sonetFarEndPathCurrentTable.SonetFarEndPathCurrentEntry[i]})
    }
    sonetFarEndPathCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetFarEndPathCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetFarEndPathCurrentTable.EntityData)
}

// SONETMIB_SonetFarEndPathCurrentTable_SonetFarEndPathCurrentEntry
// An entry in the SONET/SDH Far End Path Current table.
type SONETMIB_SonetFarEndPathCurrentTable_SonetFarEndPathCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH interface in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    SonetFarEndPathCurrentESs interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Path interface in the current 15 minute
    // interval. The type is interface{} with range: 0..4294967295.
    SonetFarEndPathCurrentSESs interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH Path
    // interface in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    SonetFarEndPathCurrentCVs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Path interface in the current 15 minute
    // interval. The type is interface{} with range: 0..4294967295.
    SonetFarEndPathCurrentUASs interface{}
}

func (sonetFarEndPathCurrentEntry *SONETMIB_SonetFarEndPathCurrentTable_SonetFarEndPathCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetFarEndPathCurrentEntry.EntityData.YFilter = sonetFarEndPathCurrentEntry.YFilter
    sonetFarEndPathCurrentEntry.EntityData.YangName = "sonetFarEndPathCurrentEntry"
    sonetFarEndPathCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndPathCurrentEntry.EntityData.ParentYangName = "sonetFarEndPathCurrentTable"
    sonetFarEndPathCurrentEntry.EntityData.SegmentPath = "sonetFarEndPathCurrentEntry" + types.AddKeyToken(sonetFarEndPathCurrentEntry.IfIndex, "ifIndex")
    sonetFarEndPathCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetFarEndPathCurrentTable/" + sonetFarEndPathCurrentEntry.EntityData.SegmentPath
    sonetFarEndPathCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndPathCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndPathCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndPathCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetFarEndPathCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetFarEndPathCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetFarEndPathCurrentEntry.IfIndex})
    sonetFarEndPathCurrentEntry.EntityData.Leafs.Append("sonetFarEndPathCurrentESs", types.YLeaf{"SonetFarEndPathCurrentESs", sonetFarEndPathCurrentEntry.SonetFarEndPathCurrentESs})
    sonetFarEndPathCurrentEntry.EntityData.Leafs.Append("sonetFarEndPathCurrentSESs", types.YLeaf{"SonetFarEndPathCurrentSESs", sonetFarEndPathCurrentEntry.SonetFarEndPathCurrentSESs})
    sonetFarEndPathCurrentEntry.EntityData.Leafs.Append("sonetFarEndPathCurrentCVs", types.YLeaf{"SonetFarEndPathCurrentCVs", sonetFarEndPathCurrentEntry.SonetFarEndPathCurrentCVs})
    sonetFarEndPathCurrentEntry.EntityData.Leafs.Append("sonetFarEndPathCurrentUASs", types.YLeaf{"SonetFarEndPathCurrentUASs", sonetFarEndPathCurrentEntry.SonetFarEndPathCurrentUASs})

    sonetFarEndPathCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetFarEndPathCurrentEntry.EntityData)
}

// SONETMIB_SonetFarEndPathIntervalTable
// The SONET/SDH Far End Path Interval table.
type SONETMIB_SonetFarEndPathIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End Path Interval table. The type is slice of
    // SONETMIB_SonetFarEndPathIntervalTable_SonetFarEndPathIntervalEntry.
    SonetFarEndPathIntervalEntry []*SONETMIB_SonetFarEndPathIntervalTable_SonetFarEndPathIntervalEntry
}

func (sonetFarEndPathIntervalTable *SONETMIB_SonetFarEndPathIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetFarEndPathIntervalTable.EntityData.YFilter = sonetFarEndPathIntervalTable.YFilter
    sonetFarEndPathIntervalTable.EntityData.YangName = "sonetFarEndPathIntervalTable"
    sonetFarEndPathIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndPathIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetFarEndPathIntervalTable.EntityData.SegmentPath = "sonetFarEndPathIntervalTable"
    sonetFarEndPathIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetFarEndPathIntervalTable.EntityData.SegmentPath
    sonetFarEndPathIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndPathIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndPathIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndPathIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetFarEndPathIntervalTable.EntityData.Children.Append("sonetFarEndPathIntervalEntry", types.YChild{"SonetFarEndPathIntervalEntry", nil})
    for i := range sonetFarEndPathIntervalTable.SonetFarEndPathIntervalEntry {
        sonetFarEndPathIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetFarEndPathIntervalTable.SonetFarEndPathIntervalEntry[i]), types.YChild{"SonetFarEndPathIntervalEntry", sonetFarEndPathIntervalTable.SonetFarEndPathIntervalEntry[i]})
    }
    sonetFarEndPathIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetFarEndPathIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetFarEndPathIntervalTable.EntityData)
}

// SONETMIB_SonetFarEndPathIntervalTable_SonetFarEndPathIntervalEntry
// An entry in the SONET/SDH Far
// End Path Interval table.
type SONETMIB_SonetFarEndPathIntervalTable_SonetFarEndPathIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetFarEndPathIntervalNumber interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH Path interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    SonetFarEndPathIntervalESs interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH Path interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    SonetFarEndPathIntervalSESs interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH Path
    // interface in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    SonetFarEndPathIntervalCVs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH Path interface in a particular 15-minute
    // interval in the past 24 hours. The type is interface{} with range:
    // 0..4294967295.
    SonetFarEndPathIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetFarEndPathIntervalValidData interface{}
}

func (sonetFarEndPathIntervalEntry *SONETMIB_SonetFarEndPathIntervalTable_SonetFarEndPathIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetFarEndPathIntervalEntry.EntityData.YFilter = sonetFarEndPathIntervalEntry.YFilter
    sonetFarEndPathIntervalEntry.EntityData.YangName = "sonetFarEndPathIntervalEntry"
    sonetFarEndPathIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndPathIntervalEntry.EntityData.ParentYangName = "sonetFarEndPathIntervalTable"
    sonetFarEndPathIntervalEntry.EntityData.SegmentPath = "sonetFarEndPathIntervalEntry" + types.AddKeyToken(sonetFarEndPathIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalNumber, "sonetFarEndPathIntervalNumber")
    sonetFarEndPathIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetFarEndPathIntervalTable/" + sonetFarEndPathIntervalEntry.EntityData.SegmentPath
    sonetFarEndPathIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndPathIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndPathIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndPathIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetFarEndPathIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetFarEndPathIntervalEntry.IfIndex})
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("sonetFarEndPathIntervalNumber", types.YLeaf{"SonetFarEndPathIntervalNumber", sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalNumber})
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("sonetFarEndPathIntervalESs", types.YLeaf{"SonetFarEndPathIntervalESs", sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalESs})
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("sonetFarEndPathIntervalSESs", types.YLeaf{"SonetFarEndPathIntervalSESs", sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalSESs})
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("sonetFarEndPathIntervalCVs", types.YLeaf{"SonetFarEndPathIntervalCVs", sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalCVs})
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("sonetFarEndPathIntervalUASs", types.YLeaf{"SonetFarEndPathIntervalUASs", sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalUASs})
    sonetFarEndPathIntervalEntry.EntityData.Leafs.Append("sonetFarEndPathIntervalValidData", types.YLeaf{"SonetFarEndPathIntervalValidData", sonetFarEndPathIntervalEntry.SonetFarEndPathIntervalValidData})

    sonetFarEndPathIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetFarEndPathIntervalNumber"}

    return &(sonetFarEndPathIntervalEntry.EntityData)
}

// SONETMIB_SonetVTCurrentTable
// The SONET/SDH VT Current table.
type SONETMIB_SonetVTCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH VT Current table. The type is slice of
    // SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry.
    SonetVTCurrentEntry []*SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry
}

func (sonetVTCurrentTable *SONETMIB_SonetVTCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetVTCurrentTable.EntityData.YFilter = sonetVTCurrentTable.YFilter
    sonetVTCurrentTable.EntityData.YangName = "sonetVTCurrentTable"
    sonetVTCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetVTCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetVTCurrentTable.EntityData.SegmentPath = "sonetVTCurrentTable"
    sonetVTCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetVTCurrentTable.EntityData.SegmentPath
    sonetVTCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetVTCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetVTCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetVTCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetVTCurrentTable.EntityData.Children.Append("sonetVTCurrentEntry", types.YChild{"SonetVTCurrentEntry", nil})
    for i := range sonetVTCurrentTable.SonetVTCurrentEntry {
        sonetVTCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetVTCurrentTable.SonetVTCurrentEntry[i]), types.YChild{"SonetVTCurrentEntry", sonetVTCurrentTable.SonetVTCurrentEntry[i]})
    }
    sonetVTCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetVTCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetVTCurrentTable.EntityData)
}

// SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry
// An entry in the SONET/SDH VT Current table.
type SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // A value that indicates the type of the SONET VT and SDH VC.  Assigned
    // widths are VT1.5/VC11, VT2/VC12, VT3, VT6/VC2, and VT6c. The type is
    // SonetVTCurrentWidth.
    SonetVTCurrentWidth interface{}

    // This variable indicates the status of the interface. The
    // sonetVTCurrentStatus is a bit map represented as a sum, therefore, it can
    // represent multiple defects and failures simultaneously. The sonetVTNoDefect
    // should be set if and only if no other flag is set.  The various bit
    // positions are:    1   sonetVTNoDefect    2   sonetVTLOP    4  
    // sonetVTPathAIS    8   sonetVTPathRDI   16   sonetVTPathRFI   32  
    // sonetVTUnequipped   64   sonetVTSignalLabelMismatch. The type is
    // interface{} with range: 1..126.
    SonetVTCurrentStatus interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH VT in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    SonetVTCurrentESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH VT in the current 15 minute interval. The type
    // is interface{} with range: 0..4294967295.
    SonetVTCurrentSESs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH VT in the current 15 minute interval. The type is interface{}
    // with range: 0..4294967295.
    SonetVTCurrentCVs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a VT in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    SonetVTCurrentUASs interface{}
}

func (sonetVTCurrentEntry *SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetVTCurrentEntry.EntityData.YFilter = sonetVTCurrentEntry.YFilter
    sonetVTCurrentEntry.EntityData.YangName = "sonetVTCurrentEntry"
    sonetVTCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetVTCurrentEntry.EntityData.ParentYangName = "sonetVTCurrentTable"
    sonetVTCurrentEntry.EntityData.SegmentPath = "sonetVTCurrentEntry" + types.AddKeyToken(sonetVTCurrentEntry.IfIndex, "ifIndex")
    sonetVTCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetVTCurrentTable/" + sonetVTCurrentEntry.EntityData.SegmentPath
    sonetVTCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetVTCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetVTCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetVTCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetVTCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetVTCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetVTCurrentEntry.IfIndex})
    sonetVTCurrentEntry.EntityData.Leafs.Append("sonetVTCurrentWidth", types.YLeaf{"SonetVTCurrentWidth", sonetVTCurrentEntry.SonetVTCurrentWidth})
    sonetVTCurrentEntry.EntityData.Leafs.Append("sonetVTCurrentStatus", types.YLeaf{"SonetVTCurrentStatus", sonetVTCurrentEntry.SonetVTCurrentStatus})
    sonetVTCurrentEntry.EntityData.Leafs.Append("sonetVTCurrentESs", types.YLeaf{"SonetVTCurrentESs", sonetVTCurrentEntry.SonetVTCurrentESs})
    sonetVTCurrentEntry.EntityData.Leafs.Append("sonetVTCurrentSESs", types.YLeaf{"SonetVTCurrentSESs", sonetVTCurrentEntry.SonetVTCurrentSESs})
    sonetVTCurrentEntry.EntityData.Leafs.Append("sonetVTCurrentCVs", types.YLeaf{"SonetVTCurrentCVs", sonetVTCurrentEntry.SonetVTCurrentCVs})
    sonetVTCurrentEntry.EntityData.Leafs.Append("sonetVTCurrentUASs", types.YLeaf{"SonetVTCurrentUASs", sonetVTCurrentEntry.SonetVTCurrentUASs})

    sonetVTCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetVTCurrentEntry.EntityData)
}

// SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth represents VT1.5/VC11, VT2/VC12, VT3, VT6/VC2, and VT6c.
type SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth string

const (
    SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth_vtWidth15VC11 SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth = "vtWidth15VC11"

    SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth_vtWidth2VC12 SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth = "vtWidth2VC12"

    SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth_vtWidth3 SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth = "vtWidth3"

    SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth_vtWidth6VC2 SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth = "vtWidth6VC2"

    SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth_vtWidth6c SONETMIB_SonetVTCurrentTable_SonetVTCurrentEntry_SonetVTCurrentWidth = "vtWidth6c"
)

// SONETMIB_SonetVTIntervalTable
// The SONET/SDH VT Interval table.
type SONETMIB_SonetVTIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH VT Interval table. The type is slice of
    // SONETMIB_SonetVTIntervalTable_SonetVTIntervalEntry.
    SonetVTIntervalEntry []*SONETMIB_SonetVTIntervalTable_SonetVTIntervalEntry
}

func (sonetVTIntervalTable *SONETMIB_SonetVTIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetVTIntervalTable.EntityData.YFilter = sonetVTIntervalTable.YFilter
    sonetVTIntervalTable.EntityData.YangName = "sonetVTIntervalTable"
    sonetVTIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetVTIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetVTIntervalTable.EntityData.SegmentPath = "sonetVTIntervalTable"
    sonetVTIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetVTIntervalTable.EntityData.SegmentPath
    sonetVTIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetVTIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetVTIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetVTIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetVTIntervalTable.EntityData.Children.Append("sonetVTIntervalEntry", types.YChild{"SonetVTIntervalEntry", nil})
    for i := range sonetVTIntervalTable.SonetVTIntervalEntry {
        sonetVTIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetVTIntervalTable.SonetVTIntervalEntry[i]), types.YChild{"SonetVTIntervalEntry", sonetVTIntervalTable.SonetVTIntervalEntry[i]})
    }
    sonetVTIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetVTIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetVTIntervalTable.EntityData)
}

// SONETMIB_SonetVTIntervalTable_SonetVTIntervalEntry
// An entry in the SONET/SDH VT Interval table.
type SONETMIB_SonetVTIntervalTable_SonetVTIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetVTIntervalNumber interface{}

    // The counter associated with the number of Errored Seconds encountered by a
    // SONET/SDH VT in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    SonetVTIntervalESs interface{}

    // The counter associated with the number of Severely Errored Seconds
    // encountered by a SONET/SDH VT in a particular 15-minute interval in the
    // past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetVTIntervalSESs interface{}

    // The counter associated with the number of Coding Violations encountered by
    // a SONET/SDH VT in a particular 15-minute interval in the past 24 hours. The
    // type is interface{} with range: 0..4294967295.
    SonetVTIntervalCVs interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by a VT in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    SonetVTIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetVTIntervalValidData interface{}
}

func (sonetVTIntervalEntry *SONETMIB_SonetVTIntervalTable_SonetVTIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetVTIntervalEntry.EntityData.YFilter = sonetVTIntervalEntry.YFilter
    sonetVTIntervalEntry.EntityData.YangName = "sonetVTIntervalEntry"
    sonetVTIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetVTIntervalEntry.EntityData.ParentYangName = "sonetVTIntervalTable"
    sonetVTIntervalEntry.EntityData.SegmentPath = "sonetVTIntervalEntry" + types.AddKeyToken(sonetVTIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetVTIntervalEntry.SonetVTIntervalNumber, "sonetVTIntervalNumber")
    sonetVTIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetVTIntervalTable/" + sonetVTIntervalEntry.EntityData.SegmentPath
    sonetVTIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetVTIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetVTIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetVTIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetVTIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetVTIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetVTIntervalEntry.IfIndex})
    sonetVTIntervalEntry.EntityData.Leafs.Append("sonetVTIntervalNumber", types.YLeaf{"SonetVTIntervalNumber", sonetVTIntervalEntry.SonetVTIntervalNumber})
    sonetVTIntervalEntry.EntityData.Leafs.Append("sonetVTIntervalESs", types.YLeaf{"SonetVTIntervalESs", sonetVTIntervalEntry.SonetVTIntervalESs})
    sonetVTIntervalEntry.EntityData.Leafs.Append("sonetVTIntervalSESs", types.YLeaf{"SonetVTIntervalSESs", sonetVTIntervalEntry.SonetVTIntervalSESs})
    sonetVTIntervalEntry.EntityData.Leafs.Append("sonetVTIntervalCVs", types.YLeaf{"SonetVTIntervalCVs", sonetVTIntervalEntry.SonetVTIntervalCVs})
    sonetVTIntervalEntry.EntityData.Leafs.Append("sonetVTIntervalUASs", types.YLeaf{"SonetVTIntervalUASs", sonetVTIntervalEntry.SonetVTIntervalUASs})
    sonetVTIntervalEntry.EntityData.Leafs.Append("sonetVTIntervalValidData", types.YLeaf{"SonetVTIntervalValidData", sonetVTIntervalEntry.SonetVTIntervalValidData})

    sonetVTIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetVTIntervalNumber"}

    return &(sonetVTIntervalEntry.EntityData)
}

// SONETMIB_SonetFarEndVTCurrentTable
// The SONET/SDH Far End VT Current table.
type SONETMIB_SonetFarEndVTCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End VT Current table. The type is slice of
    // SONETMIB_SonetFarEndVTCurrentTable_SonetFarEndVTCurrentEntry.
    SonetFarEndVTCurrentEntry []*SONETMIB_SonetFarEndVTCurrentTable_SonetFarEndVTCurrentEntry
}

func (sonetFarEndVTCurrentTable *SONETMIB_SonetFarEndVTCurrentTable) GetEntityData() *types.CommonEntityData {
    sonetFarEndVTCurrentTable.EntityData.YFilter = sonetFarEndVTCurrentTable.YFilter
    sonetFarEndVTCurrentTable.EntityData.YangName = "sonetFarEndVTCurrentTable"
    sonetFarEndVTCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndVTCurrentTable.EntityData.ParentYangName = "SONET-MIB"
    sonetFarEndVTCurrentTable.EntityData.SegmentPath = "sonetFarEndVTCurrentTable"
    sonetFarEndVTCurrentTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetFarEndVTCurrentTable.EntityData.SegmentPath
    sonetFarEndVTCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndVTCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndVTCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndVTCurrentTable.EntityData.Children = types.NewOrderedMap()
    sonetFarEndVTCurrentTable.EntityData.Children.Append("sonetFarEndVTCurrentEntry", types.YChild{"SonetFarEndVTCurrentEntry", nil})
    for i := range sonetFarEndVTCurrentTable.SonetFarEndVTCurrentEntry {
        sonetFarEndVTCurrentTable.EntityData.Children.Append(types.GetSegmentPath(sonetFarEndVTCurrentTable.SonetFarEndVTCurrentEntry[i]), types.YChild{"SonetFarEndVTCurrentEntry", sonetFarEndVTCurrentTable.SonetFarEndVTCurrentEntry[i]})
    }
    sonetFarEndVTCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    sonetFarEndVTCurrentTable.EntityData.YListKeys = []string {}

    return &(sonetFarEndVTCurrentTable.EntityData)
}

// SONETMIB_SonetFarEndVTCurrentTable_SonetFarEndVTCurrentEntry
// An entry in the SONET/SDH Far End VT Current table.
type SONETMIB_SonetFarEndVTCurrentTable_SonetFarEndVTCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH interface in the current 15 minute interval. The
    // type is interface{} with range: 0..4294967295.
    SonetFarEndVTCurrentESs interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH VT interface in the current 15 minute interval.
    // The type is interface{} with range: 0..4294967295.
    SonetFarEndVTCurrentSESs interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH VT
    // interface in the current 15 minute interval. The type is interface{} with
    // range: 0..4294967295.
    SonetFarEndVTCurrentCVs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH VT interface in the current 15 minute interval.
    // The type is interface{} with range: 0..4294967295.
    SonetFarEndVTCurrentUASs interface{}
}

func (sonetFarEndVTCurrentEntry *SONETMIB_SonetFarEndVTCurrentTable_SonetFarEndVTCurrentEntry) GetEntityData() *types.CommonEntityData {
    sonetFarEndVTCurrentEntry.EntityData.YFilter = sonetFarEndVTCurrentEntry.YFilter
    sonetFarEndVTCurrentEntry.EntityData.YangName = "sonetFarEndVTCurrentEntry"
    sonetFarEndVTCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndVTCurrentEntry.EntityData.ParentYangName = "sonetFarEndVTCurrentTable"
    sonetFarEndVTCurrentEntry.EntityData.SegmentPath = "sonetFarEndVTCurrentEntry" + types.AddKeyToken(sonetFarEndVTCurrentEntry.IfIndex, "ifIndex")
    sonetFarEndVTCurrentEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetFarEndVTCurrentTable/" + sonetFarEndVTCurrentEntry.EntityData.SegmentPath
    sonetFarEndVTCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndVTCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndVTCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndVTCurrentEntry.EntityData.Children = types.NewOrderedMap()
    sonetFarEndVTCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetFarEndVTCurrentEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetFarEndVTCurrentEntry.IfIndex})
    sonetFarEndVTCurrentEntry.EntityData.Leafs.Append("sonetFarEndVTCurrentESs", types.YLeaf{"SonetFarEndVTCurrentESs", sonetFarEndVTCurrentEntry.SonetFarEndVTCurrentESs})
    sonetFarEndVTCurrentEntry.EntityData.Leafs.Append("sonetFarEndVTCurrentSESs", types.YLeaf{"SonetFarEndVTCurrentSESs", sonetFarEndVTCurrentEntry.SonetFarEndVTCurrentSESs})
    sonetFarEndVTCurrentEntry.EntityData.Leafs.Append("sonetFarEndVTCurrentCVs", types.YLeaf{"SonetFarEndVTCurrentCVs", sonetFarEndVTCurrentEntry.SonetFarEndVTCurrentCVs})
    sonetFarEndVTCurrentEntry.EntityData.Leafs.Append("sonetFarEndVTCurrentUASs", types.YLeaf{"SonetFarEndVTCurrentUASs", sonetFarEndVTCurrentEntry.SonetFarEndVTCurrentUASs})

    sonetFarEndVTCurrentEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(sonetFarEndVTCurrentEntry.EntityData)
}

// SONETMIB_SonetFarEndVTIntervalTable
// The SONET/SDH Far End VT Interval table.
type SONETMIB_SonetFarEndVTIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the SONET/SDH Far End VT Interval table. The type is slice of
    // SONETMIB_SonetFarEndVTIntervalTable_SonetFarEndVTIntervalEntry.
    SonetFarEndVTIntervalEntry []*SONETMIB_SonetFarEndVTIntervalTable_SonetFarEndVTIntervalEntry
}

func (sonetFarEndVTIntervalTable *SONETMIB_SonetFarEndVTIntervalTable) GetEntityData() *types.CommonEntityData {
    sonetFarEndVTIntervalTable.EntityData.YFilter = sonetFarEndVTIntervalTable.YFilter
    sonetFarEndVTIntervalTable.EntityData.YangName = "sonetFarEndVTIntervalTable"
    sonetFarEndVTIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndVTIntervalTable.EntityData.ParentYangName = "SONET-MIB"
    sonetFarEndVTIntervalTable.EntityData.SegmentPath = "sonetFarEndVTIntervalTable"
    sonetFarEndVTIntervalTable.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/" + sonetFarEndVTIntervalTable.EntityData.SegmentPath
    sonetFarEndVTIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndVTIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndVTIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndVTIntervalTable.EntityData.Children = types.NewOrderedMap()
    sonetFarEndVTIntervalTable.EntityData.Children.Append("sonetFarEndVTIntervalEntry", types.YChild{"SonetFarEndVTIntervalEntry", nil})
    for i := range sonetFarEndVTIntervalTable.SonetFarEndVTIntervalEntry {
        sonetFarEndVTIntervalTable.EntityData.Children.Append(types.GetSegmentPath(sonetFarEndVTIntervalTable.SonetFarEndVTIntervalEntry[i]), types.YChild{"SonetFarEndVTIntervalEntry", sonetFarEndVTIntervalTable.SonetFarEndVTIntervalEntry[i]})
    }
    sonetFarEndVTIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    sonetFarEndVTIntervalTable.EntityData.YListKeys = []string {}

    return &(sonetFarEndVTIntervalTable.EntityData)
}

// SONETMIB_SonetFarEndVTIntervalTable_SonetFarEndVTIntervalEntry
// An entry in the SONET/SDH Far
// End VT Interval table.
type SONETMIB_SonetFarEndVTIntervalTable_SonetFarEndVTIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. A number between 1 and 96, which identifies the
    // interval for which the set of statistics is available. The interval
    // identified by 1 is the most recently completed 15 minute interval, and the
    // interval identified by N is the interval immediately preceding the one
    // identified by N-1. The type is interface{} with range: 1..96.
    SonetFarEndVTIntervalNumber interface{}

    // The counter associated with the number of Far End Errored Seconds
    // encountered by a SONET/SDH VT interface in a particular 15-minute interval
    // in the past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetFarEndVTIntervalESs interface{}

    // The counter associated with the number of Far End Severely Errored Seconds
    // encountered by a SONET/SDH VT interface in a particular 15-minute interval
    // in the past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetFarEndVTIntervalSESs interface{}

    // The counter associated with the number of Far End Coding Violations
    // reported via the far end block error count encountered by a SONET/SDH VT
    // interface in a particular 15-minute interval in the past 24 hours. The type
    // is interface{} with range: 0..4294967295.
    SonetFarEndVTIntervalCVs interface{}

    // The counter associated with the number of Far End Unavailable Seconds
    // encountered by a SONET/SDH VT interface in a particular 15-minute interval
    // in the past 24 hours. The type is interface{} with range: 0..4294967295.
    SonetFarEndVTIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    SonetFarEndVTIntervalValidData interface{}
}

func (sonetFarEndVTIntervalEntry *SONETMIB_SonetFarEndVTIntervalTable_SonetFarEndVTIntervalEntry) GetEntityData() *types.CommonEntityData {
    sonetFarEndVTIntervalEntry.EntityData.YFilter = sonetFarEndVTIntervalEntry.YFilter
    sonetFarEndVTIntervalEntry.EntityData.YangName = "sonetFarEndVTIntervalEntry"
    sonetFarEndVTIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    sonetFarEndVTIntervalEntry.EntityData.ParentYangName = "sonetFarEndVTIntervalTable"
    sonetFarEndVTIntervalEntry.EntityData.SegmentPath = "sonetFarEndVTIntervalEntry" + types.AddKeyToken(sonetFarEndVTIntervalEntry.IfIndex, "ifIndex") + types.AddKeyToken(sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalNumber, "sonetFarEndVTIntervalNumber")
    sonetFarEndVTIntervalEntry.EntityData.AbsolutePath = "SONET-MIB:SONET-MIB/sonetFarEndVTIntervalTable/" + sonetFarEndVTIntervalEntry.EntityData.SegmentPath
    sonetFarEndVTIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sonetFarEndVTIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sonetFarEndVTIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sonetFarEndVTIntervalEntry.EntityData.Children = types.NewOrderedMap()
    sonetFarEndVTIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", sonetFarEndVTIntervalEntry.IfIndex})
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("sonetFarEndVTIntervalNumber", types.YLeaf{"SonetFarEndVTIntervalNumber", sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalNumber})
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("sonetFarEndVTIntervalESs", types.YLeaf{"SonetFarEndVTIntervalESs", sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalESs})
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("sonetFarEndVTIntervalSESs", types.YLeaf{"SonetFarEndVTIntervalSESs", sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalSESs})
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("sonetFarEndVTIntervalCVs", types.YLeaf{"SonetFarEndVTIntervalCVs", sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalCVs})
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("sonetFarEndVTIntervalUASs", types.YLeaf{"SonetFarEndVTIntervalUASs", sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalUASs})
    sonetFarEndVTIntervalEntry.EntityData.Leafs.Append("sonetFarEndVTIntervalValidData", types.YLeaf{"SonetFarEndVTIntervalValidData", sonetFarEndVTIntervalEntry.SonetFarEndVTIntervalValidData})

    sonetFarEndVTIntervalEntry.EntityData.YListKeys = []string {"IfIndex", "SonetFarEndVTIntervalNumber"}

    return &(sonetFarEndVTIntervalEntry.EntityData)
}

