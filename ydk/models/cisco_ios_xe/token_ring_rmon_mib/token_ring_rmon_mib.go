package token_ring_rmon_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package token_ring_rmon_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:TOKEN-RING-RMON-MIB TOKEN-RING-RMON-MIB}", reflect.TypeOf(TOKENRINGRMONMIB{}))
    ydk.RegisterEntity("TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB", reflect.TypeOf(TOKENRINGRMONMIB{}))
}

// EntryStatus
type EntryStatus string

const (
    EntryStatus_valid EntryStatus = "valid"

    EntryStatus_createRequest EntryStatus = "createRequest"

    EntryStatus_underCreation EntryStatus = "underCreation"

    EntryStatus_invalid EntryStatus = "invalid"
)

// TOKENRINGRMONMIB
type TOKENRINGRMONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of Mac-Layer Token Ring statistics      entries.
    Tokenringmlstatstable TOKENRINGRMONMIB_Tokenringmlstatstable

    // A list of promiscuous Token Ring statistics entries.
    Tokenringpstatstable TOKENRINGRMONMIB_Tokenringpstatstable

    // A list of Mac-Layer Token Ring statistics      entries.
    Tokenringmlhistorytable TOKENRINGRMONMIB_Tokenringmlhistorytable

    // A list of promiscuous Token Ring statistics entries.
    Tokenringphistorytable TOKENRINGRMONMIB_Tokenringphistorytable

    // A list of ringStation table control entries.
    Ringstationcontroltable TOKENRINGRMONMIB_Ringstationcontroltable

    // A list of ring station entries.  An entry will exist for each station that
    // is now or has      previously been detected as physically present on this
    // ring.
    Ringstationtable TOKENRINGRMONMIB_Ringstationtable

    // A list of ring station entries for stations in the ring poll, ordered by
    // their ring-order.
    Ringstationordertable TOKENRINGRMONMIB_Ringstationordertable

    // A list of ring station configuration control entries.
    Ringstationconfigcontroltable TOKENRINGRMONMIB_Ringstationconfigcontroltable

    // A list of configuration entries for stations on a ring monitored by this
    // probe.
    Ringstationconfigtable TOKENRINGRMONMIB_Ringstationconfigtable

    // A list of source routing statistics entries.
    Sourceroutingstatstable TOKENRINGRMONMIB_Sourceroutingstatstable
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetFilter() yfilter.YFilter { return tOKENRINGRMONMIB.YFilter }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) SetFilter(yf yfilter.YFilter) { tOKENRINGRMONMIB.YFilter = yf }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetGoName(yname string) string {
    if yname == "tokenRingMLStatsTable" { return "Tokenringmlstatstable" }
    if yname == "tokenRingPStatsTable" { return "Tokenringpstatstable" }
    if yname == "tokenRingMLHistoryTable" { return "Tokenringmlhistorytable" }
    if yname == "tokenRingPHistoryTable" { return "Tokenringphistorytable" }
    if yname == "ringStationControlTable" { return "Ringstationcontroltable" }
    if yname == "ringStationTable" { return "Ringstationtable" }
    if yname == "ringStationOrderTable" { return "Ringstationordertable" }
    if yname == "ringStationConfigControlTable" { return "Ringstationconfigcontroltable" }
    if yname == "ringStationConfigTable" { return "Ringstationconfigtable" }
    if yname == "sourceRoutingStatsTable" { return "Sourceroutingstatstable" }
    return ""
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetSegmentPath() string {
    return "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB"
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tokenRingMLStatsTable" {
        return &tOKENRINGRMONMIB.Tokenringmlstatstable
    }
    if childYangName == "tokenRingPStatsTable" {
        return &tOKENRINGRMONMIB.Tokenringpstatstable
    }
    if childYangName == "tokenRingMLHistoryTable" {
        return &tOKENRINGRMONMIB.Tokenringmlhistorytable
    }
    if childYangName == "tokenRingPHistoryTable" {
        return &tOKENRINGRMONMIB.Tokenringphistorytable
    }
    if childYangName == "ringStationControlTable" {
        return &tOKENRINGRMONMIB.Ringstationcontroltable
    }
    if childYangName == "ringStationTable" {
        return &tOKENRINGRMONMIB.Ringstationtable
    }
    if childYangName == "ringStationOrderTable" {
        return &tOKENRINGRMONMIB.Ringstationordertable
    }
    if childYangName == "ringStationConfigControlTable" {
        return &tOKENRINGRMONMIB.Ringstationconfigcontroltable
    }
    if childYangName == "ringStationConfigTable" {
        return &tOKENRINGRMONMIB.Ringstationconfigtable
    }
    if childYangName == "sourceRoutingStatsTable" {
        return &tOKENRINGRMONMIB.Sourceroutingstatstable
    }
    return nil
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tokenRingMLStatsTable"] = &tOKENRINGRMONMIB.Tokenringmlstatstable
    children["tokenRingPStatsTable"] = &tOKENRINGRMONMIB.Tokenringpstatstable
    children["tokenRingMLHistoryTable"] = &tOKENRINGRMONMIB.Tokenringmlhistorytable
    children["tokenRingPHistoryTable"] = &tOKENRINGRMONMIB.Tokenringphistorytable
    children["ringStationControlTable"] = &tOKENRINGRMONMIB.Ringstationcontroltable
    children["ringStationTable"] = &tOKENRINGRMONMIB.Ringstationtable
    children["ringStationOrderTable"] = &tOKENRINGRMONMIB.Ringstationordertable
    children["ringStationConfigControlTable"] = &tOKENRINGRMONMIB.Ringstationconfigcontroltable
    children["ringStationConfigTable"] = &tOKENRINGRMONMIB.Ringstationconfigtable
    children["sourceRoutingStatsTable"] = &tOKENRINGRMONMIB.Sourceroutingstatstable
    return children
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetYangName() string { return "TOKEN-RING-RMON-MIB" }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) SetParent(parent types.Entity) { tOKENRINGRMONMIB.parent = parent }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetParent() types.Entity { return tOKENRINGRMONMIB.parent }

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Tokenringmlstatstable
// A list of Mac-Layer Token Ring statistics
// 
// 
// 
// 
// 
// entries.
type TOKENRINGRMONMIB_Tokenringmlstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of Mac-Layer statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry.
    Tokenringmlstatsentry []TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetFilter() yfilter.YFilter { return tokenringmlstatstable.YFilter }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) SetFilter(yf yfilter.YFilter) { tokenringmlstatstable.YFilter = yf }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetGoName(yname string) string {
    if yname == "tokenRingMLStatsEntry" { return "Tokenringmlstatsentry" }
    return ""
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetSegmentPath() string {
    return "tokenRingMLStatsTable"
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tokenRingMLStatsEntry" {
        for _, c := range tokenringmlstatstable.Tokenringmlstatsentry {
            if tokenringmlstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry{}
        tokenringmlstatstable.Tokenringmlstatsentry = append(tokenringmlstatstable.Tokenringmlstatsentry, child)
        return &tokenringmlstatstable.Tokenringmlstatsentry[len(tokenringmlstatstable.Tokenringmlstatsentry)-1]
    }
    return nil
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tokenringmlstatstable.Tokenringmlstatsentry {
        children[tokenringmlstatstable.Tokenringmlstatsentry[i].GetSegmentPath()] = &tokenringmlstatstable.Tokenringmlstatsentry[i]
    }
    return children
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetYangName() string { return "tokenRingMLStatsTable" }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) SetParent(parent types.Entity) { tokenringmlstatstable.parent = parent }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetParent() types.Entity { return tokenringmlstatstable.parent }

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry
// A collection of Mac-Layer statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies this
    // tokenRingMLStats entry. The type is interface{} with range: 1..65535.
    Tokenringmlstatsindex interface{}

    // This object identifies the source of the data that this tokenRingMLStats
    // entry is configured to analyze.  This source can be any tokenRing interface
    // on this device.  In order to identify a particular interface, this object
    // shall identify the instance of the ifIndex object, defined in MIB-II [3],
    // for the desired interface.  For example, if an entry were to receive data
    // from interface #1, this object would be set to ifIndex.1.  The statistics
    // in this group reflect all error reports on the local network segment
    // attached to the identified interface.  This object may not be modified if
    // the associated tokenRingMLStatsStatus object is equal to valid(1). The type
    // is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Tokenringmlstatsdatasource interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources. Note that this number is not necessarily the number
    // of packets dropped; it is just the number of times this condition has been
    // detected.  This value is the same as the corresponding
    // tokenRingPStatsDropEvents. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlstatsdropevents interface{}

    // The total number of octets of data in MAC packets (excluding those that
    // were not good frames) received on the network (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    Tokenringmlstatsmacoctets interface{}

    // The total number of MAC packets (excluding packets that were not good
    // frames) received. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatsmacpkts interface{}

    // The total number of times that the ring enters the ring purge state from
    // normal ring state.  The ring purge state that comes in response to the
    // claim token or beacon state is not counted. The type is interface{} with
    // range: 0..4294967295.
    Tokenringmlstatsringpurgeevents interface{}

    // The total number of ring purge MAC packets detected by probe. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlstatsringpurgepkts interface{}

    // The total number of times that the ring enters a beaconing state
    // (beaconFrameStreamingState, beaconBitStreamingState,     
    // beaconSetRecoveryModeState, or beaconRingSignalLossState) from a
    // non-beaconing state.  Note that a change of the source address of the
    // beacon packet does not constitute a new beacon event. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlstatsbeaconevents interface{}

    // The total amount of time that the ring has been in the beaconing state. The
    // type is interface{} with range: -2147483648..2147483647.
    Tokenringmlstatsbeacontime interface{}

    // The total number of beacon MAC packets detected by the probe. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlstatsbeaconpkts interface{}

    // The total number of times that the ring enters the claim token state from
    // normal ring state or ring purge state.  The claim token state that comes in
    // response to a beacon state is not counted. The type is interface{} with
    // range: 0..4294967295.
    Tokenringmlstatsclaimtokenevents interface{}

    // The total number of claim token MAC packets detected by the probe. The type
    // is interface{} with range: 0..4294967295.
    Tokenringmlstatsclaimtokenpkts interface{}

    // The total number of NAUN changes detected by the probe. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlstatsnaunchanges interface{}

    // The total number of line errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatslineerrors interface{}

    // The total number of adapter internal errors reported in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlstatsinternalerrors interface{}

    // The total number of burst errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatsbursterrors interface{}

    // The total number of AC (Address Copied)  errors reported in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlstatsacerrors interface{}

    // The total number of abort delimiters reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatsaborterrors interface{}

    // The total number of lost frame errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatslostframeerrors interface{}

    // The total number of receive congestion errors reported in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlstatscongestionerrors interface{}

    // The total number of frame copied errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatsframecopiederrors interface{}

    // The total number of frequency errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatsfrequencyerrors interface{}

    // The total number of token errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Tokenringmlstatstokenerrors interface{}

    // The total number of soft error report frames detected by the probe. The
    // type is interface{} with range: 0..4294967295.
    Tokenringmlstatssofterrorreports interface{}

    // The total number of ring poll events detected by the probe (i.e. the number
    // of ring polls initiated by the active monitor that were detected). The type
    // is interface{} with range: 0..4294967295.
    Tokenringmlstatsringpollevents interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    Tokenringmlstatsowner interface{}

    // The status of this tokenRingMLStats entry. The type is EntryStatus.
    Tokenringmlstatsstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Tokenringmlstatsdroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlstatscreatetime interface{}
}

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetFilter() yfilter.YFilter { return tokenringmlstatsentry.YFilter }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) SetFilter(yf yfilter.YFilter) { tokenringmlstatsentry.YFilter = yf }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetGoName(yname string) string {
    if yname == "tokenRingMLStatsIndex" { return "Tokenringmlstatsindex" }
    if yname == "tokenRingMLStatsDataSource" { return "Tokenringmlstatsdatasource" }
    if yname == "tokenRingMLStatsDropEvents" { return "Tokenringmlstatsdropevents" }
    if yname == "tokenRingMLStatsMacOctets" { return "Tokenringmlstatsmacoctets" }
    if yname == "tokenRingMLStatsMacPkts" { return "Tokenringmlstatsmacpkts" }
    if yname == "tokenRingMLStatsRingPurgeEvents" { return "Tokenringmlstatsringpurgeevents" }
    if yname == "tokenRingMLStatsRingPurgePkts" { return "Tokenringmlstatsringpurgepkts" }
    if yname == "tokenRingMLStatsBeaconEvents" { return "Tokenringmlstatsbeaconevents" }
    if yname == "tokenRingMLStatsBeaconTime" { return "Tokenringmlstatsbeacontime" }
    if yname == "tokenRingMLStatsBeaconPkts" { return "Tokenringmlstatsbeaconpkts" }
    if yname == "tokenRingMLStatsClaimTokenEvents" { return "Tokenringmlstatsclaimtokenevents" }
    if yname == "tokenRingMLStatsClaimTokenPkts" { return "Tokenringmlstatsclaimtokenpkts" }
    if yname == "tokenRingMLStatsNAUNChanges" { return "Tokenringmlstatsnaunchanges" }
    if yname == "tokenRingMLStatsLineErrors" { return "Tokenringmlstatslineerrors" }
    if yname == "tokenRingMLStatsInternalErrors" { return "Tokenringmlstatsinternalerrors" }
    if yname == "tokenRingMLStatsBurstErrors" { return "Tokenringmlstatsbursterrors" }
    if yname == "tokenRingMLStatsACErrors" { return "Tokenringmlstatsacerrors" }
    if yname == "tokenRingMLStatsAbortErrors" { return "Tokenringmlstatsaborterrors" }
    if yname == "tokenRingMLStatsLostFrameErrors" { return "Tokenringmlstatslostframeerrors" }
    if yname == "tokenRingMLStatsCongestionErrors" { return "Tokenringmlstatscongestionerrors" }
    if yname == "tokenRingMLStatsFrameCopiedErrors" { return "Tokenringmlstatsframecopiederrors" }
    if yname == "tokenRingMLStatsFrequencyErrors" { return "Tokenringmlstatsfrequencyerrors" }
    if yname == "tokenRingMLStatsTokenErrors" { return "Tokenringmlstatstokenerrors" }
    if yname == "tokenRingMLStatsSoftErrorReports" { return "Tokenringmlstatssofterrorreports" }
    if yname == "tokenRingMLStatsRingPollEvents" { return "Tokenringmlstatsringpollevents" }
    if yname == "tokenRingMLStatsOwner" { return "Tokenringmlstatsowner" }
    if yname == "tokenRingMLStatsStatus" { return "Tokenringmlstatsstatus" }
    if yname == "tokenRingMLStatsDroppedFrames" { return "Tokenringmlstatsdroppedframes" }
    if yname == "tokenRingMLStatsCreateTime" { return "Tokenringmlstatscreatetime" }
    return ""
}

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetSegmentPath() string {
    return "tokenRingMLStatsEntry" + "[tokenRingMLStatsIndex='" + fmt.Sprintf("%v", tokenringmlstatsentry.Tokenringmlstatsindex) + "']"
}

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tokenRingMLStatsIndex"] = tokenringmlstatsentry.Tokenringmlstatsindex
    leafs["tokenRingMLStatsDataSource"] = tokenringmlstatsentry.Tokenringmlstatsdatasource
    leafs["tokenRingMLStatsDropEvents"] = tokenringmlstatsentry.Tokenringmlstatsdropevents
    leafs["tokenRingMLStatsMacOctets"] = tokenringmlstatsentry.Tokenringmlstatsmacoctets
    leafs["tokenRingMLStatsMacPkts"] = tokenringmlstatsentry.Tokenringmlstatsmacpkts
    leafs["tokenRingMLStatsRingPurgeEvents"] = tokenringmlstatsentry.Tokenringmlstatsringpurgeevents
    leafs["tokenRingMLStatsRingPurgePkts"] = tokenringmlstatsentry.Tokenringmlstatsringpurgepkts
    leafs["tokenRingMLStatsBeaconEvents"] = tokenringmlstatsentry.Tokenringmlstatsbeaconevents
    leafs["tokenRingMLStatsBeaconTime"] = tokenringmlstatsentry.Tokenringmlstatsbeacontime
    leafs["tokenRingMLStatsBeaconPkts"] = tokenringmlstatsentry.Tokenringmlstatsbeaconpkts
    leafs["tokenRingMLStatsClaimTokenEvents"] = tokenringmlstatsentry.Tokenringmlstatsclaimtokenevents
    leafs["tokenRingMLStatsClaimTokenPkts"] = tokenringmlstatsentry.Tokenringmlstatsclaimtokenpkts
    leafs["tokenRingMLStatsNAUNChanges"] = tokenringmlstatsentry.Tokenringmlstatsnaunchanges
    leafs["tokenRingMLStatsLineErrors"] = tokenringmlstatsentry.Tokenringmlstatslineerrors
    leafs["tokenRingMLStatsInternalErrors"] = tokenringmlstatsentry.Tokenringmlstatsinternalerrors
    leafs["tokenRingMLStatsBurstErrors"] = tokenringmlstatsentry.Tokenringmlstatsbursterrors
    leafs["tokenRingMLStatsACErrors"] = tokenringmlstatsentry.Tokenringmlstatsacerrors
    leafs["tokenRingMLStatsAbortErrors"] = tokenringmlstatsentry.Tokenringmlstatsaborterrors
    leafs["tokenRingMLStatsLostFrameErrors"] = tokenringmlstatsentry.Tokenringmlstatslostframeerrors
    leafs["tokenRingMLStatsCongestionErrors"] = tokenringmlstatsentry.Tokenringmlstatscongestionerrors
    leafs["tokenRingMLStatsFrameCopiedErrors"] = tokenringmlstatsentry.Tokenringmlstatsframecopiederrors
    leafs["tokenRingMLStatsFrequencyErrors"] = tokenringmlstatsentry.Tokenringmlstatsfrequencyerrors
    leafs["tokenRingMLStatsTokenErrors"] = tokenringmlstatsentry.Tokenringmlstatstokenerrors
    leafs["tokenRingMLStatsSoftErrorReports"] = tokenringmlstatsentry.Tokenringmlstatssofterrorreports
    leafs["tokenRingMLStatsRingPollEvents"] = tokenringmlstatsentry.Tokenringmlstatsringpollevents
    leafs["tokenRingMLStatsOwner"] = tokenringmlstatsentry.Tokenringmlstatsowner
    leafs["tokenRingMLStatsStatus"] = tokenringmlstatsentry.Tokenringmlstatsstatus
    leafs["tokenRingMLStatsDroppedFrames"] = tokenringmlstatsentry.Tokenringmlstatsdroppedframes
    leafs["tokenRingMLStatsCreateTime"] = tokenringmlstatsentry.Tokenringmlstatscreatetime
    return leafs
}

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetYangName() string { return "tokenRingMLStatsEntry" }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) SetParent(parent types.Entity) { tokenringmlstatsentry.parent = parent }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetParent() types.Entity { return tokenringmlstatsentry.parent }

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetParentYangName() string { return "tokenRingMLStatsTable" }

// TOKENRINGRMONMIB_Tokenringpstatstable
// A list of promiscuous Token Ring statistics
// entries.
type TOKENRINGRMONMIB_Tokenringpstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of promiscuous statistics kept for non-MAC packets on a
    // particular Token Ring interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry.
    Tokenringpstatsentry []TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetFilter() yfilter.YFilter { return tokenringpstatstable.YFilter }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) SetFilter(yf yfilter.YFilter) { tokenringpstatstable.YFilter = yf }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetGoName(yname string) string {
    if yname == "tokenRingPStatsEntry" { return "Tokenringpstatsentry" }
    return ""
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetSegmentPath() string {
    return "tokenRingPStatsTable"
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tokenRingPStatsEntry" {
        for _, c := range tokenringpstatstable.Tokenringpstatsentry {
            if tokenringpstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry{}
        tokenringpstatstable.Tokenringpstatsentry = append(tokenringpstatstable.Tokenringpstatsentry, child)
        return &tokenringpstatstable.Tokenringpstatsentry[len(tokenringpstatstable.Tokenringpstatsentry)-1]
    }
    return nil
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tokenringpstatstable.Tokenringpstatsentry {
        children[tokenringpstatstable.Tokenringpstatsentry[i].GetSegmentPath()] = &tokenringpstatstable.Tokenringpstatsentry[i]
    }
    return children
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetYangName() string { return "tokenRingPStatsTable" }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) SetParent(parent types.Entity) { tokenringpstatstable.parent = parent }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetParent() types.Entity { return tokenringpstatstable.parent }

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry
// A collection of promiscuous statistics kept for
// non-MAC packets on a particular Token Ring
// interface.
type TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies this
    // tokenRingPStats entry. The type is interface{} with range: 1..65535.
    Tokenringpstatsindex interface{}

    // This object identifies the source of the data that this tokenRingPStats
    // entry is configured to analyze.  This source can be any tokenRing interface
    // on this device.  In order to identify a particular interface, this object
    // shall identify the instance of the ifIndex object, defined in MIB-II [3],
    // for the desired interface.  For example, if an entry were to receive data
    // from interface #1, this object would be set to ifIndex.1.  The statistics
    // in this group reflect all non-MAC packets on the local network segment
    // attached to the identified interface.  This object may not be modified if
    // the associated tokenRingPStatsStatus object is equal to valid(1). The type
    // is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Tokenringpstatsdatasource interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources. Note that this number is not necessarily the number
    // of packets dropped; it is just the number of times this condition has been
    // detected.  This value is the same as the corresponding
    // tokenRingMLStatsDropEvents. The type is interface{} with range:
    // 0..4294967295.
    Tokenringpstatsdropevents interface{}

    // The total number of octets of data in good frames received on the network
    // (excluding framing bits but including FCS octets) in non-MAC packets. The
    // type is interface{} with range: 0..4294967295.
    Tokenringpstatsdataoctets interface{}

    // The total number of non-MAC packets in good frames.  received. The type is
    // interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts interface{}

    // The total number of good non-MAC frames received that were directed to an
    // LLC broadcast address (0xFFFFFFFFFFFF or 0xC000FFFFFFFF). The type is
    // interface{} with range: 0..4294967295.
    Tokenringpstatsdatabroadcastpkts interface{}

    // The total number of good non-MAC frames received that were directed to a
    // local or global multicast or functional address.  Note that this number
    // does not include packets directed to the broadcast address. The type is
    // interface{} with range: 0..4294967295.
    Tokenringpstatsdatamulticastpkts interface{}

    // The total number of good non-MAC frames received that were between 18 and
    // 63 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts18To63Octets interface{}

    // The total number of good non-MAC frames received that were between 64 and
    // 127 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts64To127Octets interface{}

    // The total number of good non-MAC frames received that were between 128 and
    // 255 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts128To255Octets interface{}

    // The total number of good non-MAC frames received that were between 256 and
    // 511 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts256To511Octets interface{}

    // The total number of good non-MAC frames received that were between 512 and
    // 1023 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts512To1023Octets interface{}

    // The total number of good non-MAC frames received that were between 1024 and
    // 2047 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts1024To2047Octets interface{}

    // The total number of good non-MAC frames received that were between 2048 and
    // 4095 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts2048To4095Octets interface{}

    // The total number of good non-MAC frames received that were between 4096 and
    // 8191 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts4096To8191Octets interface{}

    // The total number of good non-MAC frames received that were between 8192 and
    // 18000 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapkts8192To18000Octets interface{}

    // The total number of good non-MAC frames received that were greater than
    // 18000 octets in length, excluding framing bits but including FCS octets.
    // The type is interface{} with range: 0..4294967295.
    Tokenringpstatsdatapktsgreaterthan18000Octets interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    Tokenringpstatsowner interface{}

    // The status of this tokenRingPStats entry. The type is EntryStatus.
    Tokenringpstatsstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Tokenringpstatsdroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Tokenringpstatscreatetime interface{}
}

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetFilter() yfilter.YFilter { return tokenringpstatsentry.YFilter }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) SetFilter(yf yfilter.YFilter) { tokenringpstatsentry.YFilter = yf }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetGoName(yname string) string {
    if yname == "tokenRingPStatsIndex" { return "Tokenringpstatsindex" }
    if yname == "tokenRingPStatsDataSource" { return "Tokenringpstatsdatasource" }
    if yname == "tokenRingPStatsDropEvents" { return "Tokenringpstatsdropevents" }
    if yname == "tokenRingPStatsDataOctets" { return "Tokenringpstatsdataoctets" }
    if yname == "tokenRingPStatsDataPkts" { return "Tokenringpstatsdatapkts" }
    if yname == "tokenRingPStatsDataBroadcastPkts" { return "Tokenringpstatsdatabroadcastpkts" }
    if yname == "tokenRingPStatsDataMulticastPkts" { return "Tokenringpstatsdatamulticastpkts" }
    if yname == "tokenRingPStatsDataPkts18to63Octets" { return "Tokenringpstatsdatapkts18To63Octets" }
    if yname == "tokenRingPStatsDataPkts64to127Octets" { return "Tokenringpstatsdatapkts64To127Octets" }
    if yname == "tokenRingPStatsDataPkts128to255Octets" { return "Tokenringpstatsdatapkts128To255Octets" }
    if yname == "tokenRingPStatsDataPkts256to511Octets" { return "Tokenringpstatsdatapkts256To511Octets" }
    if yname == "tokenRingPStatsDataPkts512to1023Octets" { return "Tokenringpstatsdatapkts512To1023Octets" }
    if yname == "tokenRingPStatsDataPkts1024to2047Octets" { return "Tokenringpstatsdatapkts1024To2047Octets" }
    if yname == "tokenRingPStatsDataPkts2048to4095Octets" { return "Tokenringpstatsdatapkts2048To4095Octets" }
    if yname == "tokenRingPStatsDataPkts4096to8191Octets" { return "Tokenringpstatsdatapkts4096To8191Octets" }
    if yname == "tokenRingPStatsDataPkts8192to18000Octets" { return "Tokenringpstatsdatapkts8192To18000Octets" }
    if yname == "tokenRingPStatsDataPktsGreaterThan18000Octets" { return "Tokenringpstatsdatapktsgreaterthan18000Octets" }
    if yname == "tokenRingPStatsOwner" { return "Tokenringpstatsowner" }
    if yname == "tokenRingPStatsStatus" { return "Tokenringpstatsstatus" }
    if yname == "tokenRingPStatsDroppedFrames" { return "Tokenringpstatsdroppedframes" }
    if yname == "tokenRingPStatsCreateTime" { return "Tokenringpstatscreatetime" }
    return ""
}

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetSegmentPath() string {
    return "tokenRingPStatsEntry" + "[tokenRingPStatsIndex='" + fmt.Sprintf("%v", tokenringpstatsentry.Tokenringpstatsindex) + "']"
}

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tokenRingPStatsIndex"] = tokenringpstatsentry.Tokenringpstatsindex
    leafs["tokenRingPStatsDataSource"] = tokenringpstatsentry.Tokenringpstatsdatasource
    leafs["tokenRingPStatsDropEvents"] = tokenringpstatsentry.Tokenringpstatsdropevents
    leafs["tokenRingPStatsDataOctets"] = tokenringpstatsentry.Tokenringpstatsdataoctets
    leafs["tokenRingPStatsDataPkts"] = tokenringpstatsentry.Tokenringpstatsdatapkts
    leafs["tokenRingPStatsDataBroadcastPkts"] = tokenringpstatsentry.Tokenringpstatsdatabroadcastpkts
    leafs["tokenRingPStatsDataMulticastPkts"] = tokenringpstatsentry.Tokenringpstatsdatamulticastpkts
    leafs["tokenRingPStatsDataPkts18to63Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts18To63Octets
    leafs["tokenRingPStatsDataPkts64to127Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts64To127Octets
    leafs["tokenRingPStatsDataPkts128to255Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts128To255Octets
    leafs["tokenRingPStatsDataPkts256to511Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts256To511Octets
    leafs["tokenRingPStatsDataPkts512to1023Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts512To1023Octets
    leafs["tokenRingPStatsDataPkts1024to2047Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts1024To2047Octets
    leafs["tokenRingPStatsDataPkts2048to4095Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts2048To4095Octets
    leafs["tokenRingPStatsDataPkts4096to8191Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts4096To8191Octets
    leafs["tokenRingPStatsDataPkts8192to18000Octets"] = tokenringpstatsentry.Tokenringpstatsdatapkts8192To18000Octets
    leafs["tokenRingPStatsDataPktsGreaterThan18000Octets"] = tokenringpstatsentry.Tokenringpstatsdatapktsgreaterthan18000Octets
    leafs["tokenRingPStatsOwner"] = tokenringpstatsentry.Tokenringpstatsowner
    leafs["tokenRingPStatsStatus"] = tokenringpstatsentry.Tokenringpstatsstatus
    leafs["tokenRingPStatsDroppedFrames"] = tokenringpstatsentry.Tokenringpstatsdroppedframes
    leafs["tokenRingPStatsCreateTime"] = tokenringpstatsentry.Tokenringpstatscreatetime
    return leafs
}

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetYangName() string { return "tokenRingPStatsEntry" }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) SetParent(parent types.Entity) { tokenringpstatsentry.parent = parent }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetParent() types.Entity { return tokenringpstatsentry.parent }

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetParentYangName() string { return "tokenRingPStatsTable" }

// TOKENRINGRMONMIB_Tokenringmlhistorytable
// A list of Mac-Layer Token Ring statistics
// 
// 
// 
// 
// 
// entries.
type TOKENRINGRMONMIB_Tokenringmlhistorytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of Mac-Layer statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry.
    Tokenringmlhistoryentry []TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetFilter() yfilter.YFilter { return tokenringmlhistorytable.YFilter }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) SetFilter(yf yfilter.YFilter) { tokenringmlhistorytable.YFilter = yf }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetGoName(yname string) string {
    if yname == "tokenRingMLHistoryEntry" { return "Tokenringmlhistoryentry" }
    return ""
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetSegmentPath() string {
    return "tokenRingMLHistoryTable"
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tokenRingMLHistoryEntry" {
        for _, c := range tokenringmlhistorytable.Tokenringmlhistoryentry {
            if tokenringmlhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry{}
        tokenringmlhistorytable.Tokenringmlhistoryentry = append(tokenringmlhistorytable.Tokenringmlhistoryentry, child)
        return &tokenringmlhistorytable.Tokenringmlhistoryentry[len(tokenringmlhistorytable.Tokenringmlhistoryentry)-1]
    }
    return nil
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tokenringmlhistorytable.Tokenringmlhistoryentry {
        children[tokenringmlhistorytable.Tokenringmlhistoryentry[i].GetSegmentPath()] = &tokenringmlhistorytable.Tokenringmlhistoryentry[i]
    }
    return children
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetYangName() string { return "tokenRingMLHistoryTable" }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) SetParent(parent types.Entity) { tokenringmlhistorytable.parent = parent }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetParent() types.Entity { return tokenringmlhistorytable.parent }

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry
// A collection of Mac-Layer statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The history of which this entry is a part.  The
    // history identified by a particular value of this index is the same history
    // as identified by the same value of historyControlIndex. The type is
    // interface{} with range: 1..65535.
    Tokenringmlhistoryindex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // Mac-Layer sample this entry represents among all Mac-Layer samples
    // associated with the same historyControlEntry.  This index starts at 1 and
    // increases by one as each new sample is taken. The type is interface{} with
    // range: -2147483648..2147483647.
    Tokenringmlhistorysampleindex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval.  Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlhistoryintervalstart interface{}

    // The total number of events in which packets were      dropped by the probe
    // due to lack of resources during this sampling interval.  Note that this
    // number is not necessarily the number of packets dropped, it is just the
    // number of times this condition has been detected. The type is interface{}
    // with range: 0..4294967295.
    Tokenringmlhistorydropevents interface{}

    // The total number of octets of data in MAC packets (excluding those that
    // were not good frames) received on the network during this sampling interval
    // (excluding framing bits but including FCS octets). The type is interface{}
    // with range: 0..4294967295.
    Tokenringmlhistorymacoctets interface{}

    // The total number of MAC packets (excluding those that were not good frames)
    // received during this sampling interval. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlhistorymacpkts interface{}

    // The total number of times that the ring entered the ring purge state from
    // normal ring state during this sampling interval.  The ring purge state that
    // comes from the claim token or beacon state is not counted. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistoryringpurgeevents interface{}

    // The total number of Ring Purge MAC packets detected by the probe during
    // this sampling      interval. The type is interface{} with range:
    // 0..4294967295.
    Tokenringmlhistoryringpurgepkts interface{}

    // The total number of times that the ring enters a beaconing state
    // (beaconFrameStreamingState, beaconBitStreamingState,
    // beaconSetRecoveryModeState, or beaconRingSignalLossState) during this
    // sampling interval.  Note that a change of the source address of the beacon
    // packet does not constitute a new beacon event. The type is interface{} with
    // range: 0..4294967295.
    Tokenringmlhistorybeaconevents interface{}

    // The amount of time that the ring has been in the beaconing state during
    // this sampling interval. The type is interface{} with range:
    // -2147483648..2147483647.
    Tokenringmlhistorybeacontime interface{}

    // The total number of beacon MAC packets detected by the probe during this
    // sampling interval. The type is interface{} with range: 0..4294967295.
    Tokenringmlhistorybeaconpkts interface{}

    // The total number of times that the ring enters the claim token state from
    // normal ring state or ring purge state during this sampling interval. The
    // claim token state that comes from the beacon state is not counted. The type
    // is interface{} with range: 0..4294967295.
    Tokenringmlhistoryclaimtokenevents interface{}

    // The total number of claim token MAC packets detected by the probe during
    // this sampling interval. The type is interface{} with range: 0..4294967295.
    Tokenringmlhistoryclaimtokenpkts interface{}

    // The total number of NAUN changes detected by the probe during this sampling
    // interval. The type is interface{} with range: 0..4294967295.
    Tokenringmlhistorynaunchanges interface{}

    // The total number of line errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistorylineerrors interface{}

    // The total number of adapter internal errors reported in error reporting
    // packets detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistoryinternalerrors interface{}

    // The total number of burst errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistorybursterrors interface{}

    // The total number of AC (Address Copied) errors reported in error reporting
    // packets detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistoryacerrors interface{}

    // The total number of abort delimiters reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistoryaborterrors interface{}

    // The total number of lost frame errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistorylostframeerrors interface{}

    // The total number of receive congestion errors reported in error reporting
    // packets detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistorycongestionerrors interface{}

    // The total number of frame copied errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistoryframecopiederrors interface{}

    // The total number of frequency errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistoryfrequencyerrors interface{}

    // The total number of token errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    Tokenringmlhistorytokenerrors interface{}

    // The total number of soft error report frames detected by the probe during
    // this sampling interval. The type is interface{} with range: 0..4294967295.
    Tokenringmlhistorysofterrorreports interface{}

    // The total number of ring poll events detected by the probe during this
    // sampling interval. The type is interface{} with range: 0..4294967295.
    Tokenringmlhistoryringpollevents interface{}

    // The maximum number of active stations on the ring detected by the probe
    // during this sampling      interval. The type is interface{} with range:
    // -2147483648..2147483647.
    Tokenringmlhistoryactivestations interface{}
}

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetFilter() yfilter.YFilter { return tokenringmlhistoryentry.YFilter }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) SetFilter(yf yfilter.YFilter) { tokenringmlhistoryentry.YFilter = yf }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetGoName(yname string) string {
    if yname == "tokenRingMLHistoryIndex" { return "Tokenringmlhistoryindex" }
    if yname == "tokenRingMLHistorySampleIndex" { return "Tokenringmlhistorysampleindex" }
    if yname == "tokenRingMLHistoryIntervalStart" { return "Tokenringmlhistoryintervalstart" }
    if yname == "tokenRingMLHistoryDropEvents" { return "Tokenringmlhistorydropevents" }
    if yname == "tokenRingMLHistoryMacOctets" { return "Tokenringmlhistorymacoctets" }
    if yname == "tokenRingMLHistoryMacPkts" { return "Tokenringmlhistorymacpkts" }
    if yname == "tokenRingMLHistoryRingPurgeEvents" { return "Tokenringmlhistoryringpurgeevents" }
    if yname == "tokenRingMLHistoryRingPurgePkts" { return "Tokenringmlhistoryringpurgepkts" }
    if yname == "tokenRingMLHistoryBeaconEvents" { return "Tokenringmlhistorybeaconevents" }
    if yname == "tokenRingMLHistoryBeaconTime" { return "Tokenringmlhistorybeacontime" }
    if yname == "tokenRingMLHistoryBeaconPkts" { return "Tokenringmlhistorybeaconpkts" }
    if yname == "tokenRingMLHistoryClaimTokenEvents" { return "Tokenringmlhistoryclaimtokenevents" }
    if yname == "tokenRingMLHistoryClaimTokenPkts" { return "Tokenringmlhistoryclaimtokenpkts" }
    if yname == "tokenRingMLHistoryNAUNChanges" { return "Tokenringmlhistorynaunchanges" }
    if yname == "tokenRingMLHistoryLineErrors" { return "Tokenringmlhistorylineerrors" }
    if yname == "tokenRingMLHistoryInternalErrors" { return "Tokenringmlhistoryinternalerrors" }
    if yname == "tokenRingMLHistoryBurstErrors" { return "Tokenringmlhistorybursterrors" }
    if yname == "tokenRingMLHistoryACErrors" { return "Tokenringmlhistoryacerrors" }
    if yname == "tokenRingMLHistoryAbortErrors" { return "Tokenringmlhistoryaborterrors" }
    if yname == "tokenRingMLHistoryLostFrameErrors" { return "Tokenringmlhistorylostframeerrors" }
    if yname == "tokenRingMLHistoryCongestionErrors" { return "Tokenringmlhistorycongestionerrors" }
    if yname == "tokenRingMLHistoryFrameCopiedErrors" { return "Tokenringmlhistoryframecopiederrors" }
    if yname == "tokenRingMLHistoryFrequencyErrors" { return "Tokenringmlhistoryfrequencyerrors" }
    if yname == "tokenRingMLHistoryTokenErrors" { return "Tokenringmlhistorytokenerrors" }
    if yname == "tokenRingMLHistorySoftErrorReports" { return "Tokenringmlhistorysofterrorreports" }
    if yname == "tokenRingMLHistoryRingPollEvents" { return "Tokenringmlhistoryringpollevents" }
    if yname == "tokenRingMLHistoryActiveStations" { return "Tokenringmlhistoryactivestations" }
    return ""
}

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetSegmentPath() string {
    return "tokenRingMLHistoryEntry" + "[tokenRingMLHistoryIndex='" + fmt.Sprintf("%v", tokenringmlhistoryentry.Tokenringmlhistoryindex) + "']" + "[tokenRingMLHistorySampleIndex='" + fmt.Sprintf("%v", tokenringmlhistoryentry.Tokenringmlhistorysampleindex) + "']"
}

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tokenRingMLHistoryIndex"] = tokenringmlhistoryentry.Tokenringmlhistoryindex
    leafs["tokenRingMLHistorySampleIndex"] = tokenringmlhistoryentry.Tokenringmlhistorysampleindex
    leafs["tokenRingMLHistoryIntervalStart"] = tokenringmlhistoryentry.Tokenringmlhistoryintervalstart
    leafs["tokenRingMLHistoryDropEvents"] = tokenringmlhistoryentry.Tokenringmlhistorydropevents
    leafs["tokenRingMLHistoryMacOctets"] = tokenringmlhistoryentry.Tokenringmlhistorymacoctets
    leafs["tokenRingMLHistoryMacPkts"] = tokenringmlhistoryentry.Tokenringmlhistorymacpkts
    leafs["tokenRingMLHistoryRingPurgeEvents"] = tokenringmlhistoryentry.Tokenringmlhistoryringpurgeevents
    leafs["tokenRingMLHistoryRingPurgePkts"] = tokenringmlhistoryentry.Tokenringmlhistoryringpurgepkts
    leafs["tokenRingMLHistoryBeaconEvents"] = tokenringmlhistoryentry.Tokenringmlhistorybeaconevents
    leafs["tokenRingMLHistoryBeaconTime"] = tokenringmlhistoryentry.Tokenringmlhistorybeacontime
    leafs["tokenRingMLHistoryBeaconPkts"] = tokenringmlhistoryentry.Tokenringmlhistorybeaconpkts
    leafs["tokenRingMLHistoryClaimTokenEvents"] = tokenringmlhistoryentry.Tokenringmlhistoryclaimtokenevents
    leafs["tokenRingMLHistoryClaimTokenPkts"] = tokenringmlhistoryentry.Tokenringmlhistoryclaimtokenpkts
    leafs["tokenRingMLHistoryNAUNChanges"] = tokenringmlhistoryentry.Tokenringmlhistorynaunchanges
    leafs["tokenRingMLHistoryLineErrors"] = tokenringmlhistoryentry.Tokenringmlhistorylineerrors
    leafs["tokenRingMLHistoryInternalErrors"] = tokenringmlhistoryentry.Tokenringmlhistoryinternalerrors
    leafs["tokenRingMLHistoryBurstErrors"] = tokenringmlhistoryentry.Tokenringmlhistorybursterrors
    leafs["tokenRingMLHistoryACErrors"] = tokenringmlhistoryentry.Tokenringmlhistoryacerrors
    leafs["tokenRingMLHistoryAbortErrors"] = tokenringmlhistoryentry.Tokenringmlhistoryaborterrors
    leafs["tokenRingMLHistoryLostFrameErrors"] = tokenringmlhistoryentry.Tokenringmlhistorylostframeerrors
    leafs["tokenRingMLHistoryCongestionErrors"] = tokenringmlhistoryentry.Tokenringmlhistorycongestionerrors
    leafs["tokenRingMLHistoryFrameCopiedErrors"] = tokenringmlhistoryentry.Tokenringmlhistoryframecopiederrors
    leafs["tokenRingMLHistoryFrequencyErrors"] = tokenringmlhistoryentry.Tokenringmlhistoryfrequencyerrors
    leafs["tokenRingMLHistoryTokenErrors"] = tokenringmlhistoryentry.Tokenringmlhistorytokenerrors
    leafs["tokenRingMLHistorySoftErrorReports"] = tokenringmlhistoryentry.Tokenringmlhistorysofterrorreports
    leafs["tokenRingMLHistoryRingPollEvents"] = tokenringmlhistoryentry.Tokenringmlhistoryringpollevents
    leafs["tokenRingMLHistoryActiveStations"] = tokenringmlhistoryentry.Tokenringmlhistoryactivestations
    return leafs
}

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetYangName() string { return "tokenRingMLHistoryEntry" }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) SetParent(parent types.Entity) { tokenringmlhistoryentry.parent = parent }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetParent() types.Entity { return tokenringmlhistoryentry.parent }

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetParentYangName() string { return "tokenRingMLHistoryTable" }

// TOKENRINGRMONMIB_Tokenringphistorytable
// A list of promiscuous Token Ring statistics
// entries.
type TOKENRINGRMONMIB_Tokenringphistorytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of promiscuous statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry.
    Tokenringphistoryentry []TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetFilter() yfilter.YFilter { return tokenringphistorytable.YFilter }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) SetFilter(yf yfilter.YFilter) { tokenringphistorytable.YFilter = yf }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetGoName(yname string) string {
    if yname == "tokenRingPHistoryEntry" { return "Tokenringphistoryentry" }
    return ""
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetSegmentPath() string {
    return "tokenRingPHistoryTable"
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tokenRingPHistoryEntry" {
        for _, c := range tokenringphistorytable.Tokenringphistoryentry {
            if tokenringphistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry{}
        tokenringphistorytable.Tokenringphistoryentry = append(tokenringphistorytable.Tokenringphistoryentry, child)
        return &tokenringphistorytable.Tokenringphistoryentry[len(tokenringphistorytable.Tokenringphistoryentry)-1]
    }
    return nil
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tokenringphistorytable.Tokenringphistoryentry {
        children[tokenringphistorytable.Tokenringphistoryentry[i].GetSegmentPath()] = &tokenringphistorytable.Tokenringphistoryentry[i]
    }
    return children
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetYangName() string { return "tokenRingPHistoryTable" }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) SetParent(parent types.Entity) { tokenringphistorytable.parent = parent }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetParent() types.Entity { return tokenringphistorytable.parent }

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry
// A collection of promiscuous statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The history of which this entry is a part.  The
    // history identified by a particular value of this index is the same history
    // as identified by the same value of historyControlIndex. The type is
    // interface{} with range: 1..65535.
    Tokenringphistoryindex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // sample this entry represents among all samples associated with the same
    // historyControlEntry. This index starts at 1 and increases by one as each
    // new sample is taken. The type is interface{} with range:
    // -2147483648..2147483647.
    Tokenringphistorysampleindex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval.  Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    Tokenringphistoryintervalstart interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources during this sampling interval.  Note that this number
    // is not necessarily the number of packets dropped, it is just the number of
    // times this condition has been detected. The type is interface{} with range:
    // 0..4294967295.
    Tokenringphistorydropevents interface{}

    // The total number of octets of data in good frames received on the network
    // (excluding framing bits but including FCS octets) in non-MAC packets during
    // this sampling interval. The type is interface{} with range: 0..4294967295.
    Tokenringphistorydataoctets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval. The type is interface{} with range: 0..4294967295.
    Tokenringphistorydatapkts interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were directed to an LLC broadcast address (0xFFFFFFFFFFFF or
    // 0xC000FFFFFFFF). The type is interface{} with range: 0..4294967295.
    Tokenringphistorydatabroadcastpkts interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were directed to a local or global multicast or functional
    // address.  Note that this number does not include packets directed to the
    // broadcast address. The type is interface{} with range: 0..4294967295.
    Tokenringphistorydatamulticastpkts interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 18 and 63 octets in length inclusive, excluding
    // framing bits but including FCS octets. The type is interface{} with range:
    // 0..4294967295.
    Tokenringphistorydatapkts18To63Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 64 and 127 octets in length inclusive, excluding
    // framing bits but including FCS octets. The type is interface{} with range:
    // 0..4294967295.
    Tokenringphistorydatapkts64To127Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 128 and 255 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts128To255Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between      256 and 511 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts256To511Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 512 and 1023 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts512To1023Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 1024 and 2047 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts1024To2047Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 2048 and 4095 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts2048To4095Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 4096 and 8191 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts4096To8191Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 8192 and 18000 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    Tokenringphistorydatapkts8192To18000Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were greater than 18000 octets in length, excluding framing
    // bits but including FCS octets. The type is interface{} with range:
    // 0..4294967295.
    Tokenringphistorydatapktsgreaterthan18000Octets interface{}
}

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetFilter() yfilter.YFilter { return tokenringphistoryentry.YFilter }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) SetFilter(yf yfilter.YFilter) { tokenringphistoryentry.YFilter = yf }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetGoName(yname string) string {
    if yname == "tokenRingPHistoryIndex" { return "Tokenringphistoryindex" }
    if yname == "tokenRingPHistorySampleIndex" { return "Tokenringphistorysampleindex" }
    if yname == "tokenRingPHistoryIntervalStart" { return "Tokenringphistoryintervalstart" }
    if yname == "tokenRingPHistoryDropEvents" { return "Tokenringphistorydropevents" }
    if yname == "tokenRingPHistoryDataOctets" { return "Tokenringphistorydataoctets" }
    if yname == "tokenRingPHistoryDataPkts" { return "Tokenringphistorydatapkts" }
    if yname == "tokenRingPHistoryDataBroadcastPkts" { return "Tokenringphistorydatabroadcastpkts" }
    if yname == "tokenRingPHistoryDataMulticastPkts" { return "Tokenringphistorydatamulticastpkts" }
    if yname == "tokenRingPHistoryDataPkts18to63Octets" { return "Tokenringphistorydatapkts18To63Octets" }
    if yname == "tokenRingPHistoryDataPkts64to127Octets" { return "Tokenringphistorydatapkts64To127Octets" }
    if yname == "tokenRingPHistoryDataPkts128to255Octets" { return "Tokenringphistorydatapkts128To255Octets" }
    if yname == "tokenRingPHistoryDataPkts256to511Octets" { return "Tokenringphistorydatapkts256To511Octets" }
    if yname == "tokenRingPHistoryDataPkts512to1023Octets" { return "Tokenringphistorydatapkts512To1023Octets" }
    if yname == "tokenRingPHistoryDataPkts1024to2047Octets" { return "Tokenringphistorydatapkts1024To2047Octets" }
    if yname == "tokenRingPHistoryDataPkts2048to4095Octets" { return "Tokenringphistorydatapkts2048To4095Octets" }
    if yname == "tokenRingPHistoryDataPkts4096to8191Octets" { return "Tokenringphistorydatapkts4096To8191Octets" }
    if yname == "tokenRingPHistoryDataPkts8192to18000Octets" { return "Tokenringphistorydatapkts8192To18000Octets" }
    if yname == "tokenRingPHistoryDataPktsGreaterThan18000Octets" { return "Tokenringphistorydatapktsgreaterthan18000Octets" }
    return ""
}

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetSegmentPath() string {
    return "tokenRingPHistoryEntry" + "[tokenRingPHistoryIndex='" + fmt.Sprintf("%v", tokenringphistoryentry.Tokenringphistoryindex) + "']" + "[tokenRingPHistorySampleIndex='" + fmt.Sprintf("%v", tokenringphistoryentry.Tokenringphistorysampleindex) + "']"
}

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tokenRingPHistoryIndex"] = tokenringphistoryentry.Tokenringphistoryindex
    leafs["tokenRingPHistorySampleIndex"] = tokenringphistoryentry.Tokenringphistorysampleindex
    leafs["tokenRingPHistoryIntervalStart"] = tokenringphistoryentry.Tokenringphistoryintervalstart
    leafs["tokenRingPHistoryDropEvents"] = tokenringphistoryentry.Tokenringphistorydropevents
    leafs["tokenRingPHistoryDataOctets"] = tokenringphistoryentry.Tokenringphistorydataoctets
    leafs["tokenRingPHistoryDataPkts"] = tokenringphistoryentry.Tokenringphistorydatapkts
    leafs["tokenRingPHistoryDataBroadcastPkts"] = tokenringphistoryentry.Tokenringphistorydatabroadcastpkts
    leafs["tokenRingPHistoryDataMulticastPkts"] = tokenringphistoryentry.Tokenringphistorydatamulticastpkts
    leafs["tokenRingPHistoryDataPkts18to63Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts18To63Octets
    leafs["tokenRingPHistoryDataPkts64to127Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts64To127Octets
    leafs["tokenRingPHistoryDataPkts128to255Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts128To255Octets
    leafs["tokenRingPHistoryDataPkts256to511Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts256To511Octets
    leafs["tokenRingPHistoryDataPkts512to1023Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts512To1023Octets
    leafs["tokenRingPHistoryDataPkts1024to2047Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts1024To2047Octets
    leafs["tokenRingPHistoryDataPkts2048to4095Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts2048To4095Octets
    leafs["tokenRingPHistoryDataPkts4096to8191Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts4096To8191Octets
    leafs["tokenRingPHistoryDataPkts8192to18000Octets"] = tokenringphistoryentry.Tokenringphistorydatapkts8192To18000Octets
    leafs["tokenRingPHistoryDataPktsGreaterThan18000Octets"] = tokenringphistoryentry.Tokenringphistorydatapktsgreaterthan18000Octets
    return leafs
}

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetYangName() string { return "tokenRingPHistoryEntry" }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) SetParent(parent types.Entity) { tokenringphistoryentry.parent = parent }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetParent() types.Entity { return tokenringphistoryentry.parent }

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetParentYangName() string { return "tokenRingPHistoryTable" }

// TOKENRINGRMONMIB_Ringstationcontroltable
// A list of ringStation table control entries.
type TOKENRINGRMONMIB_Ringstationcontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of parameters that set up the discovery of stations on a particular
    // interface and the collection of statistics about these stations. The type
    // is slice of
    // TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry.
    Ringstationcontrolentry []TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetFilter() yfilter.YFilter { return ringstationcontroltable.YFilter }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) SetFilter(yf yfilter.YFilter) { ringstationcontroltable.YFilter = yf }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetGoName(yname string) string {
    if yname == "ringStationControlEntry" { return "Ringstationcontrolentry" }
    return ""
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetSegmentPath() string {
    return "ringStationControlTable"
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ringStationControlEntry" {
        for _, c := range ringstationcontroltable.Ringstationcontrolentry {
            if ringstationcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry{}
        ringstationcontroltable.Ringstationcontrolentry = append(ringstationcontroltable.Ringstationcontrolentry, child)
        return &ringstationcontroltable.Ringstationcontrolentry[len(ringstationcontroltable.Ringstationcontrolentry)-1]
    }
    return nil
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ringstationcontroltable.Ringstationcontrolentry {
        children[ringstationcontroltable.Ringstationcontrolentry[i].GetSegmentPath()] = &ringstationcontroltable.Ringstationcontrolentry[i]
    }
    return children
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetYangName() string { return "ringStationControlTable" }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) SetParent(parent types.Entity) { ringstationcontroltable.parent = parent }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetParent() types.Entity { return ringstationcontroltable.parent }

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry
// A list of parameters that set up the discovery of
// stations on a particular interface and the
// collection of statistics about these stations.
type TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device from which ringStation
    // data is collected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB- II [3]. The type is interface{} with range:
    // 1..65535.
    Ringstationcontrolifindex interface{}

    // The number of ringStationEntries in the ringStationTable associated with
    // this ringStationControlEntry. The type is interface{} with range:
    // -2147483648..2147483647.
    Ringstationcontroltablesize interface{}

    // The number of active ringStationEntries in the ringStationTable associated
    // with this ringStationControlEntry. The type is interface{} with range:
    // -2147483648..2147483647.
    Ringstationcontrolactivestations interface{}

    // The current status of this ring. The type is Ringstationcontrolringstate.
    Ringstationcontrolringstate interface{}

    // The address of the sender of the last beacon frame received by the probe on
    // this ring.  If no beacon frames have been received, this object shall be
    // equal to six octets of zero. The type is string with length: 6.
    Ringstationcontrolbeaconsender interface{}

    // The address of the NAUN in the last beacon frame received by the probe on
    // this ring.  If no beacon frames have been received, this object shall be
    // equal to six octets of zero. The type is string with length: 6.
    Ringstationcontrolbeaconnaun interface{}

    // The address of the Active Monitor on this segment.  If this address is
    // unknown, this object shall be equal to six octets of zero. The type is
    // string with length: 6.
    Ringstationcontrolactivemonitor interface{}

    // The number of add and delete events in the ringStationOrderTable optionally
    // associated with this ringStationControlEntry. The type is interface{} with
    // range: 0..4294967295.
    Ringstationcontrolorderchanges interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    Ringstationcontrolowner interface{}

    // The status of this ringStationControl entry.  If this object is not equal
    // to valid(1), all associated entries in the ringStationTable shall be
    // deleted by the agent. The type is EntryStatus.
    Ringstationcontrolstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Ringstationcontroldroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Ringstationcontrolcreatetime interface{}
}

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetFilter() yfilter.YFilter { return ringstationcontrolentry.YFilter }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) SetFilter(yf yfilter.YFilter) { ringstationcontrolentry.YFilter = yf }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetGoName(yname string) string {
    if yname == "ringStationControlIfIndex" { return "Ringstationcontrolifindex" }
    if yname == "ringStationControlTableSize" { return "Ringstationcontroltablesize" }
    if yname == "ringStationControlActiveStations" { return "Ringstationcontrolactivestations" }
    if yname == "ringStationControlRingState" { return "Ringstationcontrolringstate" }
    if yname == "ringStationControlBeaconSender" { return "Ringstationcontrolbeaconsender" }
    if yname == "ringStationControlBeaconNAUN" { return "Ringstationcontrolbeaconnaun" }
    if yname == "ringStationControlActiveMonitor" { return "Ringstationcontrolactivemonitor" }
    if yname == "ringStationControlOrderChanges" { return "Ringstationcontrolorderchanges" }
    if yname == "ringStationControlOwner" { return "Ringstationcontrolowner" }
    if yname == "ringStationControlStatus" { return "Ringstationcontrolstatus" }
    if yname == "ringStationControlDroppedFrames" { return "Ringstationcontroldroppedframes" }
    if yname == "ringStationControlCreateTime" { return "Ringstationcontrolcreatetime" }
    return ""
}

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetSegmentPath() string {
    return "ringStationControlEntry" + "[ringStationControlIfIndex='" + fmt.Sprintf("%v", ringstationcontrolentry.Ringstationcontrolifindex) + "']"
}

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ringStationControlIfIndex"] = ringstationcontrolentry.Ringstationcontrolifindex
    leafs["ringStationControlTableSize"] = ringstationcontrolentry.Ringstationcontroltablesize
    leafs["ringStationControlActiveStations"] = ringstationcontrolentry.Ringstationcontrolactivestations
    leafs["ringStationControlRingState"] = ringstationcontrolentry.Ringstationcontrolringstate
    leafs["ringStationControlBeaconSender"] = ringstationcontrolentry.Ringstationcontrolbeaconsender
    leafs["ringStationControlBeaconNAUN"] = ringstationcontrolentry.Ringstationcontrolbeaconnaun
    leafs["ringStationControlActiveMonitor"] = ringstationcontrolentry.Ringstationcontrolactivemonitor
    leafs["ringStationControlOrderChanges"] = ringstationcontrolentry.Ringstationcontrolorderchanges
    leafs["ringStationControlOwner"] = ringstationcontrolentry.Ringstationcontrolowner
    leafs["ringStationControlStatus"] = ringstationcontrolentry.Ringstationcontrolstatus
    leafs["ringStationControlDroppedFrames"] = ringstationcontrolentry.Ringstationcontroldroppedframes
    leafs["ringStationControlCreateTime"] = ringstationcontrolentry.Ringstationcontrolcreatetime
    return leafs
}

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetYangName() string { return "ringStationControlEntry" }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) SetParent(parent types.Entity) { ringstationcontrolentry.parent = parent }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetParent() types.Entity { return ringstationcontrolentry.parent }

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetParentYangName() string { return "ringStationControlTable" }

// TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate represents The current status of this ring.
type TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate string

const (
    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_normalOperation TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "normalOperation"

    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_ringPurgeState TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "ringPurgeState"

    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_claimTokenState TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "claimTokenState"

    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_beaconFrameStreamingState TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "beaconFrameStreamingState"

    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_beaconBitStreamingState TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "beaconBitStreamingState"

    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_beaconRingSignalLossState TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "beaconRingSignalLossState"

    TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate_beaconSetRecoveryModeState TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry_Ringstationcontrolringstate = "beaconSetRecoveryModeState"
)

// TOKENRINGRMONMIB_Ringstationtable
// A list of ring station entries.  An entry will
// exist for each station that is now or has
// 
// 
// 
// 
// 
// previously been detected as physically present on
// this ring.
type TOKENRINGRMONMIB_Ringstationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for a particular station that has been
    // discovered on a ring monitored by this device. The type is slice of
    // TOKENRINGRMONMIB_Ringstationtable_Ringstationentry.
    Ringstationentry []TOKENRINGRMONMIB_Ringstationtable_Ringstationentry
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetFilter() yfilter.YFilter { return ringstationtable.YFilter }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) SetFilter(yf yfilter.YFilter) { ringstationtable.YFilter = yf }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetGoName(yname string) string {
    if yname == "ringStationEntry" { return "Ringstationentry" }
    return ""
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetSegmentPath() string {
    return "ringStationTable"
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ringStationEntry" {
        for _, c := range ringstationtable.Ringstationentry {
            if ringstationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Ringstationtable_Ringstationentry{}
        ringstationtable.Ringstationentry = append(ringstationtable.Ringstationentry, child)
        return &ringstationtable.Ringstationentry[len(ringstationtable.Ringstationentry)-1]
    }
    return nil
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ringstationtable.Ringstationentry {
        children[ringstationtable.Ringstationentry[i].GetSegmentPath()] = &ringstationtable.Ringstationentry[i]
    }
    return children
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetYangName() string { return "ringStationTable" }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) SetParent(parent types.Entity) { ringstationtable.parent = parent }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetParent() types.Entity { return ringstationtable.parent }

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Ringstationtable_Ringstationentry
// A collection of statistics for a particular
// station that has been discovered on a ring
// monitored by this device.
type TOKENRINGRMONMIB_Ringstationtable_Ringstationentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    Ringstationifindex interface{}

    // This attribute is a key. The physical address of this station. The type is
    // string with length: 6.
    Ringstationmacaddress interface{}

    // The physical address of last known NAUN of this station. The type is string
    // with length: 6.
    Ringstationlastnaun interface{}

    // The status of this station on the ring. The type is
    // Ringstationstationstatus.
    Ringstationstationstatus interface{}

    // The value of sysUpTime at the time this station last entered the ring.  If
    // the time is unknown, this value shall be zero. The type is interface{} with
    // range: 0..4294967295.
    Ringstationlastentertime interface{}

    // The value of sysUpTime at the time the probe detected that this station
    // last exited the ring. If the time is unknown, this value shall be zero. The
    // type is interface{} with range: 0..4294967295.
    Ringstationlastexittime interface{}

    // The number of times this station experienced a duplicate address error. The
    // type is interface{} with range: 0..4294967295.
    Ringstationduplicateaddresses interface{}

    // The total number of line errors reported by this station in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    Ringstationinlineerrors interface{}

    // The total number of line errors reported in error reporting packets sent by
    // the nearest active downstream neighbor of this station and detected by the
    // probe. The type is interface{} with range: 0..4294967295.
    Ringstationoutlineerrors interface{}

    // The total number of adapter internal errors reported by this station in
    // error reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationinternalerrors interface{}

    // The total number of burst errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationinbursterrors interface{}

    // The total number of burst errors reported in error reporting packets sent
    // by the nearest active downstream neighbor of this station and detected by
    // the probe. The type is interface{} with range: 0..4294967295.
    Ringstationoutbursterrors interface{}

    // The total number of AC (Address Copied) errors reported in error reporting
    // packets sent by the nearest active downstream neighbor of this station and
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    Ringstationacerrors interface{}

    // The total number of abort delimiters reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationaborterrors interface{}

    // The total number of lost frame errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationlostframeerrors interface{}

    // The total number of receive congestion errors reported by this station in
    // error reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationcongestionerrors interface{}

    // The total number of frame copied errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationframecopiederrors interface{}

    // The total number of frequency errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    Ringstationfrequencyerrors interface{}

    // The total number of token errors reported by this station in error
    // reporting frames detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    Ringstationtokenerrors interface{}

    // The total number of beacon frames sent by this station and detected by the
    // probe. The type is interface{} with range: 0..4294967295.
    Ringstationinbeaconerrors interface{}

    // The total number of beacon frames detected by the probe that name this
    // station as the NAUN. The type is interface{} with range: 0..4294967295.
    Ringstationoutbeaconerrors interface{}

    // The number of times the probe detected this station inserting onto the
    // ring. The type is interface{} with range: 0..4294967295.
    Ringstationinsertions interface{}
}

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetFilter() yfilter.YFilter { return ringstationentry.YFilter }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) SetFilter(yf yfilter.YFilter) { ringstationentry.YFilter = yf }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetGoName(yname string) string {
    if yname == "ringStationIfIndex" { return "Ringstationifindex" }
    if yname == "ringStationMacAddress" { return "Ringstationmacaddress" }
    if yname == "ringStationLastNAUN" { return "Ringstationlastnaun" }
    if yname == "ringStationStationStatus" { return "Ringstationstationstatus" }
    if yname == "ringStationLastEnterTime" { return "Ringstationlastentertime" }
    if yname == "ringStationLastExitTime" { return "Ringstationlastexittime" }
    if yname == "ringStationDuplicateAddresses" { return "Ringstationduplicateaddresses" }
    if yname == "ringStationInLineErrors" { return "Ringstationinlineerrors" }
    if yname == "ringStationOutLineErrors" { return "Ringstationoutlineerrors" }
    if yname == "ringStationInternalErrors" { return "Ringstationinternalerrors" }
    if yname == "ringStationInBurstErrors" { return "Ringstationinbursterrors" }
    if yname == "ringStationOutBurstErrors" { return "Ringstationoutbursterrors" }
    if yname == "ringStationACErrors" { return "Ringstationacerrors" }
    if yname == "ringStationAbortErrors" { return "Ringstationaborterrors" }
    if yname == "ringStationLostFrameErrors" { return "Ringstationlostframeerrors" }
    if yname == "ringStationCongestionErrors" { return "Ringstationcongestionerrors" }
    if yname == "ringStationFrameCopiedErrors" { return "Ringstationframecopiederrors" }
    if yname == "ringStationFrequencyErrors" { return "Ringstationfrequencyerrors" }
    if yname == "ringStationTokenErrors" { return "Ringstationtokenerrors" }
    if yname == "ringStationInBeaconErrors" { return "Ringstationinbeaconerrors" }
    if yname == "ringStationOutBeaconErrors" { return "Ringstationoutbeaconerrors" }
    if yname == "ringStationInsertions" { return "Ringstationinsertions" }
    return ""
}

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetSegmentPath() string {
    return "ringStationEntry" + "[ringStationIfIndex='" + fmt.Sprintf("%v", ringstationentry.Ringstationifindex) + "']" + "[ringStationMacAddress='" + fmt.Sprintf("%v", ringstationentry.Ringstationmacaddress) + "']"
}

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ringStationIfIndex"] = ringstationentry.Ringstationifindex
    leafs["ringStationMacAddress"] = ringstationentry.Ringstationmacaddress
    leafs["ringStationLastNAUN"] = ringstationentry.Ringstationlastnaun
    leafs["ringStationStationStatus"] = ringstationentry.Ringstationstationstatus
    leafs["ringStationLastEnterTime"] = ringstationentry.Ringstationlastentertime
    leafs["ringStationLastExitTime"] = ringstationentry.Ringstationlastexittime
    leafs["ringStationDuplicateAddresses"] = ringstationentry.Ringstationduplicateaddresses
    leafs["ringStationInLineErrors"] = ringstationentry.Ringstationinlineerrors
    leafs["ringStationOutLineErrors"] = ringstationentry.Ringstationoutlineerrors
    leafs["ringStationInternalErrors"] = ringstationentry.Ringstationinternalerrors
    leafs["ringStationInBurstErrors"] = ringstationentry.Ringstationinbursterrors
    leafs["ringStationOutBurstErrors"] = ringstationentry.Ringstationoutbursterrors
    leafs["ringStationACErrors"] = ringstationentry.Ringstationacerrors
    leafs["ringStationAbortErrors"] = ringstationentry.Ringstationaborterrors
    leafs["ringStationLostFrameErrors"] = ringstationentry.Ringstationlostframeerrors
    leafs["ringStationCongestionErrors"] = ringstationentry.Ringstationcongestionerrors
    leafs["ringStationFrameCopiedErrors"] = ringstationentry.Ringstationframecopiederrors
    leafs["ringStationFrequencyErrors"] = ringstationentry.Ringstationfrequencyerrors
    leafs["ringStationTokenErrors"] = ringstationentry.Ringstationtokenerrors
    leafs["ringStationInBeaconErrors"] = ringstationentry.Ringstationinbeaconerrors
    leafs["ringStationOutBeaconErrors"] = ringstationentry.Ringstationoutbeaconerrors
    leafs["ringStationInsertions"] = ringstationentry.Ringstationinsertions
    return leafs
}

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetYangName() string { return "ringStationEntry" }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) SetParent(parent types.Entity) { ringstationentry.parent = parent }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetParent() types.Entity { return ringstationentry.parent }

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetParentYangName() string { return "ringStationTable" }

// TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus represents The status of this station on the ring.
type TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus string

const (
    TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus_active TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus = "active"

    TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus_inactive TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus = "inactive"

    TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus_forcedRemoval TOKENRINGRMONMIB_Ringstationtable_Ringstationentry_Ringstationstationstatus = "forcedRemoval"
)

// TOKENRINGRMONMIB_Ringstationordertable
// A list of ring station entries for stations in
// the ring poll, ordered by their ring-order.
type TOKENRINGRMONMIB_Ringstationordertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for a particular      station that is active on
    // a ring monitored by this device.  This table will contain information for
    // every interface that has a ringStationControlStatus equal to valid. The
    // type is slice of
    // TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry.
    Ringstationorderentry []TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetFilter() yfilter.YFilter { return ringstationordertable.YFilter }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) SetFilter(yf yfilter.YFilter) { ringstationordertable.YFilter = yf }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetGoName(yname string) string {
    if yname == "ringStationOrderEntry" { return "Ringstationorderentry" }
    return ""
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetSegmentPath() string {
    return "ringStationOrderTable"
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ringStationOrderEntry" {
        for _, c := range ringstationordertable.Ringstationorderentry {
            if ringstationordertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry{}
        ringstationordertable.Ringstationorderentry = append(ringstationordertable.Ringstationorderentry, child)
        return &ringstationordertable.Ringstationorderentry[len(ringstationordertable.Ringstationorderentry)-1]
    }
    return nil
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ringstationordertable.Ringstationorderentry {
        children[ringstationordertable.Ringstationorderentry[i].GetSegmentPath()] = &ringstationordertable.Ringstationorderentry[i]
    }
    return children
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetYangName() string { return "ringStationOrderTable" }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) SetParent(parent types.Entity) { ringstationordertable.parent = parent }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetParent() types.Entity { return ringstationordertable.parent }

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry
// A collection of statistics for a particular
// 
// 
// 
// 
// 
// station that is active on a ring monitored by this
// device.  This table will contain information for
// every interface that has a
// ringStationControlStatus equal to valid.
type TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    Ringstationorderifindex interface{}

    // This attribute is a key. This index denotes the location of this station
    // with respect to other stations on the ring.  This index is one more than
    // the number of hops downstream that this station is from the rmon probe. 
    // The rmon probe itself gets the value one. The type is interface{} with
    // range: -2147483648..2147483647.
    Ringstationorderorderindex interface{}

    // The physical address of this station. The type is string with length: 6.
    Ringstationordermacaddress interface{}
}

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetFilter() yfilter.YFilter { return ringstationorderentry.YFilter }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) SetFilter(yf yfilter.YFilter) { ringstationorderentry.YFilter = yf }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetGoName(yname string) string {
    if yname == "ringStationOrderIfIndex" { return "Ringstationorderifindex" }
    if yname == "ringStationOrderOrderIndex" { return "Ringstationorderorderindex" }
    if yname == "ringStationOrderMacAddress" { return "Ringstationordermacaddress" }
    return ""
}

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetSegmentPath() string {
    return "ringStationOrderEntry" + "[ringStationOrderIfIndex='" + fmt.Sprintf("%v", ringstationorderentry.Ringstationorderifindex) + "']" + "[ringStationOrderOrderIndex='" + fmt.Sprintf("%v", ringstationorderentry.Ringstationorderorderindex) + "']"
}

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ringStationOrderIfIndex"] = ringstationorderentry.Ringstationorderifindex
    leafs["ringStationOrderOrderIndex"] = ringstationorderentry.Ringstationorderorderindex
    leafs["ringStationOrderMacAddress"] = ringstationorderentry.Ringstationordermacaddress
    return leafs
}

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetYangName() string { return "ringStationOrderEntry" }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) SetParent(parent types.Entity) { ringstationorderentry.parent = parent }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetParent() types.Entity { return ringstationorderentry.parent }

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetParentYangName() string { return "ringStationOrderTable" }

// TOKENRINGRMONMIB_Ringstationconfigcontroltable
// A list of ring station configuration control
// entries.
type TOKENRINGRMONMIB_Ringstationconfigcontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry controls active management of stations by the probe.  One entry
    // exists in this table for each active station in the ringStationTable. The
    // type is slice of
    // TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry.
    Ringstationconfigcontrolentry []TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetFilter() yfilter.YFilter { return ringstationconfigcontroltable.YFilter }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) SetFilter(yf yfilter.YFilter) { ringstationconfigcontroltable.YFilter = yf }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetGoName(yname string) string {
    if yname == "ringStationConfigControlEntry" { return "Ringstationconfigcontrolentry" }
    return ""
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetSegmentPath() string {
    return "ringStationConfigControlTable"
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ringStationConfigControlEntry" {
        for _, c := range ringstationconfigcontroltable.Ringstationconfigcontrolentry {
            if ringstationconfigcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry{}
        ringstationconfigcontroltable.Ringstationconfigcontrolentry = append(ringstationconfigcontroltable.Ringstationconfigcontrolentry, child)
        return &ringstationconfigcontroltable.Ringstationconfigcontrolentry[len(ringstationconfigcontroltable.Ringstationconfigcontrolentry)-1]
    }
    return nil
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ringstationconfigcontroltable.Ringstationconfigcontrolentry {
        children[ringstationconfigcontroltable.Ringstationconfigcontrolentry[i].GetSegmentPath()] = &ringstationconfigcontroltable.Ringstationconfigcontrolentry[i]
    }
    return children
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetYangName() string { return "ringStationConfigControlTable" }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) SetParent(parent types.Entity) { ringstationconfigcontroltable.parent = parent }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetParent() types.Entity { return ringstationconfigcontroltable.parent }

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry
// This entry controls active management of stations
// by the probe.  One entry exists in this table for
// each active station in the ringStationTable.
type TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the  
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    Ringstationconfigcontrolifindex interface{}

    // This attribute is a key. The physical address of this station. The type is
    // string with length: 6.
    Ringstationconfigcontrolmacaddress interface{}

    // Setting this object to `removing(2)' causes a Remove Station MAC frame to
    // be sent.  The agent will set this object to `stable(1)' after processing
    // the request. The type is Ringstationconfigcontrolremove.
    Ringstationconfigcontrolremove interface{}

    // Setting this object to `updating(2)' causes the configuration information
    // associate with this entry to be updated.  The agent will set this object to
    // `stable(1)' after processing the request. The type is
    // Ringstationconfigcontrolupdatestats.
    Ringstationconfigcontrolupdatestats interface{}
}

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetFilter() yfilter.YFilter { return ringstationconfigcontrolentry.YFilter }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) SetFilter(yf yfilter.YFilter) { ringstationconfigcontrolentry.YFilter = yf }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetGoName(yname string) string {
    if yname == "ringStationConfigControlIfIndex" { return "Ringstationconfigcontrolifindex" }
    if yname == "ringStationConfigControlMacAddress" { return "Ringstationconfigcontrolmacaddress" }
    if yname == "ringStationConfigControlRemove" { return "Ringstationconfigcontrolremove" }
    if yname == "ringStationConfigControlUpdateStats" { return "Ringstationconfigcontrolupdatestats" }
    return ""
}

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetSegmentPath() string {
    return "ringStationConfigControlEntry" + "[ringStationConfigControlIfIndex='" + fmt.Sprintf("%v", ringstationconfigcontrolentry.Ringstationconfigcontrolifindex) + "']" + "[ringStationConfigControlMacAddress='" + fmt.Sprintf("%v", ringstationconfigcontrolentry.Ringstationconfigcontrolmacaddress) + "']"
}

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ringStationConfigControlIfIndex"] = ringstationconfigcontrolentry.Ringstationconfigcontrolifindex
    leafs["ringStationConfigControlMacAddress"] = ringstationconfigcontrolentry.Ringstationconfigcontrolmacaddress
    leafs["ringStationConfigControlRemove"] = ringstationconfigcontrolentry.Ringstationconfigcontrolremove
    leafs["ringStationConfigControlUpdateStats"] = ringstationconfigcontrolentry.Ringstationconfigcontrolupdatestats
    return leafs
}

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetYangName() string { return "ringStationConfigControlEntry" }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) SetParent(parent types.Entity) { ringstationconfigcontrolentry.parent = parent }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetParent() types.Entity { return ringstationconfigcontrolentry.parent }

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetParentYangName() string { return "ringStationConfigControlTable" }

// TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolremove represents processing the request.
type TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolremove string

const (
    TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolremove_stable TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolremove = "stable"

    TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolremove_removing TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolremove = "removing"
)

// TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolupdatestats represents request.
type TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolupdatestats string

const (
    TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolupdatestats_stable TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolupdatestats = "stable"

    TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolupdatestats_updating TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry_Ringstationconfigcontrolupdatestats = "updating"
)

// TOKENRINGRMONMIB_Ringstationconfigtable
// A list of configuration entries for stations on a
// ring monitored by this probe.
type TOKENRINGRMONMIB_Ringstationconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for a particular station that has been
    // discovered on a ring monitored by this probe. The type is slice of
    // TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry.
    Ringstationconfigentry []TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetFilter() yfilter.YFilter { return ringstationconfigtable.YFilter }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) SetFilter(yf yfilter.YFilter) { ringstationconfigtable.YFilter = yf }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetGoName(yname string) string {
    if yname == "ringStationConfigEntry" { return "Ringstationconfigentry" }
    return ""
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetSegmentPath() string {
    return "ringStationConfigTable"
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ringStationConfigEntry" {
        for _, c := range ringstationconfigtable.Ringstationconfigentry {
            if ringstationconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry{}
        ringstationconfigtable.Ringstationconfigentry = append(ringstationconfigtable.Ringstationconfigentry, child)
        return &ringstationconfigtable.Ringstationconfigentry[len(ringstationconfigtable.Ringstationconfigentry)-1]
    }
    return nil
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ringstationconfigtable.Ringstationconfigentry {
        children[ringstationconfigtable.Ringstationconfigentry[i].GetSegmentPath()] = &ringstationconfigtable.Ringstationconfigentry[i]
    }
    return children
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetYangName() string { return "ringStationConfigTable" }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) SetParent(parent types.Entity) { ringstationconfigtable.parent = parent }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetParent() types.Entity { return ringstationconfigtable.parent }

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry
// A collection of statistics for a particular
// station that has been discovered on a ring
// monitored by this probe.
type TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the  
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    Ringstationconfigifindex interface{}

    // This attribute is a key. The physical address of this station. The type is
    // string with length: 6.
    Ringstationconfigmacaddress interface{}

    // The value of sysUpTime at the time this configuration information was last
    // updated (completely). The type is interface{} with range: 0..4294967295.
    Ringstationconfigupdatetime interface{}

    // The assigned physical location of this station. The type is string with
    // length: 4.
    Ringstationconfiglocation interface{}

    // The microcode EC level of this station. The type is string with length: 10.
    Ringstationconfigmicrocode interface{}

    // The low-order 4 octets of the group address recognized by this station. The
    // type is string with length: 4.
    Ringstationconfiggroupaddress interface{}

    // the functional addresses recognized by this station. The type is string
    // with length: 4.
    Ringstationconfigfunctionaladdress interface{}
}

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetFilter() yfilter.YFilter { return ringstationconfigentry.YFilter }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) SetFilter(yf yfilter.YFilter) { ringstationconfigentry.YFilter = yf }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetGoName(yname string) string {
    if yname == "ringStationConfigIfIndex" { return "Ringstationconfigifindex" }
    if yname == "ringStationConfigMacAddress" { return "Ringstationconfigmacaddress" }
    if yname == "ringStationConfigUpdateTime" { return "Ringstationconfigupdatetime" }
    if yname == "ringStationConfigLocation" { return "Ringstationconfiglocation" }
    if yname == "ringStationConfigMicrocode" { return "Ringstationconfigmicrocode" }
    if yname == "ringStationConfigGroupAddress" { return "Ringstationconfiggroupaddress" }
    if yname == "ringStationConfigFunctionalAddress" { return "Ringstationconfigfunctionaladdress" }
    return ""
}

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetSegmentPath() string {
    return "ringStationConfigEntry" + "[ringStationConfigIfIndex='" + fmt.Sprintf("%v", ringstationconfigentry.Ringstationconfigifindex) + "']" + "[ringStationConfigMacAddress='" + fmt.Sprintf("%v", ringstationconfigentry.Ringstationconfigmacaddress) + "']"
}

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ringStationConfigIfIndex"] = ringstationconfigentry.Ringstationconfigifindex
    leafs["ringStationConfigMacAddress"] = ringstationconfigentry.Ringstationconfigmacaddress
    leafs["ringStationConfigUpdateTime"] = ringstationconfigentry.Ringstationconfigupdatetime
    leafs["ringStationConfigLocation"] = ringstationconfigentry.Ringstationconfiglocation
    leafs["ringStationConfigMicrocode"] = ringstationconfigentry.Ringstationconfigmicrocode
    leafs["ringStationConfigGroupAddress"] = ringstationconfigentry.Ringstationconfiggroupaddress
    leafs["ringStationConfigFunctionalAddress"] = ringstationconfigentry.Ringstationconfigfunctionaladdress
    return leafs
}

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetYangName() string { return "ringStationConfigEntry" }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) SetParent(parent types.Entity) { ringstationconfigentry.parent = parent }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetParent() types.Entity { return ringstationconfigentry.parent }

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetParentYangName() string { return "ringStationConfigTable" }

// TOKENRINGRMONMIB_Sourceroutingstatstable
// A list of source routing statistics entries.
type TOKENRINGRMONMIB_Sourceroutingstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of source routing statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry.
    Sourceroutingstatsentry []TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetFilter() yfilter.YFilter { return sourceroutingstatstable.YFilter }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) SetFilter(yf yfilter.YFilter) { sourceroutingstatstable.YFilter = yf }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetGoName(yname string) string {
    if yname == "sourceRoutingStatsEntry" { return "Sourceroutingstatsentry" }
    return ""
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetSegmentPath() string {
    return "sourceRoutingStatsTable"
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sourceRoutingStatsEntry" {
        for _, c := range sourceroutingstatstable.Sourceroutingstatsentry {
            if sourceroutingstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry{}
        sourceroutingstatstable.Sourceroutingstatsentry = append(sourceroutingstatstable.Sourceroutingstatsentry, child)
        return &sourceroutingstatstable.Sourceroutingstatsentry[len(sourceroutingstatstable.Sourceroutingstatsentry)-1]
    }
    return nil
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sourceroutingstatstable.Sourceroutingstatsentry {
        children[sourceroutingstatstable.Sourceroutingstatsentry[i].GetSegmentPath()] = &sourceroutingstatstable.Sourceroutingstatsentry[i]
    }
    return children
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetYangName() string { return "sourceRoutingStatsTable" }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) SetParent(parent types.Entity) { sourceroutingstatstable.parent = parent }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetParent() types.Entity { return sourceroutingstatstable.parent }

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetParentYangName() string { return "TOKEN-RING-RMON-MIB" }

// TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry
// A collection of source routing statistics kept
// for a particular Token Ring interface.
type TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device on which source routing
    // statistics will be detected.  The interface identified by a particular
    // value of this object is the same interface as identified by the same value
    // of the ifIndex object, defined in MIB-II [3]. The type is interface{} with
    // range: -2147483648..2147483647.
    Sourceroutingstatsifindex interface{}

    // The ring number of the ring monitored by this entry.  When any object in
    // this entry is created, the probe will attempt to discover the ring number. 
    // Only after the ring number is discovered will this object be created. 
    // After creating an object in this entry, the management station should poll
    // this object to detect when it is created.  Only after this object is
    // created can the management station set the sourceRoutingStatsStatus entry
    // to valid(1). The type is interface{} with range: -2147483648..2147483647.
    Sourceroutingstatsringnumber interface{}

    // The count of frames sent into this ring from another ring. The type is
    // interface{} with range: 0..4294967295.
    Sourceroutingstatsinframes interface{}

    // The count of frames sent from this ring to another ring. The type is
    // interface{} with range: 0..4294967295.
    Sourceroutingstatsoutframes interface{}

    // The count of frames sent from another ring, through this ring, to another
    // ring. The type is interface{} with range: 0..4294967295.
    Sourceroutingstatsthroughframes interface{}

    // The total number of good frames received that were All Routes Broadcast.
    // The type is interface{} with range: 0..4294967295.
    Sourceroutingstatsallroutesbroadcastframes interface{}

    // The total number of good frames received that were Single Route Broadcast.
    // The type is interface{} with range: 0..4294967295.
    Sourceroutingstatssingleroutebroadcastframes interface{}

    // The count of octets in good frames sent into this ring from another ring.
    // The type is interface{} with range: 0..4294967295.
    Sourceroutingstatsinoctets interface{}

    // The count of octets in good frames sent from this ring to another ring. The
    // type is interface{} with range: 0..4294967295.
    Sourceroutingstatsoutoctets interface{}

    // The count of octets in good frames sent another ring, through this ring, to
    // another ring. The type is interface{} with range: 0..4294967295.
    Sourceroutingstatsthroughoctets interface{}

    // The total number of octets in good frames received that were All Routes
    // Broadcast. The type is interface{} with range: 0..4294967295.
    Sourceroutingstatsallroutesbroadcastoctets interface{}

    // The total number of octets in good frames received that were Single Route
    // Broadcast. The type is interface{} with range: 0..4294967295.
    Sourceroutingstatssingleroutesbroadcastoctets interface{}

    // The total number of frames received who had no RIF field (or had a RIF
    // field that only included the local ring's number) and were not All Route
    // Broadcast Frames. The type is interface{} with range: 0..4294967295.
    Sourceroutingstatslocalllcframes interface{}

    // The total number of frames received whose route had 1 hop, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats1Hopframes interface{}

    // The total number of frames received whose route had 2 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats2Hopsframes interface{}

    // The total number of frames received whose route had 3 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats3Hopsframes interface{}

    // The total number of frames received whose route had 4 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats4Hopsframes interface{}

    // The total number of frames received whose route had 5 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats5Hopsframes interface{}

    // The total number of frames received whose route had 6 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats6Hopsframes interface{}

    // The total number of frames received whose route had 7 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats7Hopsframes interface{}

    // The total number of frames received whose route had 8 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstats8Hopsframes interface{}

    // The total number of frames received whose route had more than 8 hops, were
    // not All Route Broadcast Frames, and whose source or destination were on
    // this ring (i.e. frames that had a RIF field and had this ring number in the
    // first or last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstatsmorethan8Hopsframes interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    Sourceroutingstatsowner interface{}

    // The status of this sourceRoutingStats entry. The type is EntryStatus.
    Sourceroutingstatsstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Sourceroutingstatsdroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Sourceroutingstatscreatetime interface{}
}

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetFilter() yfilter.YFilter { return sourceroutingstatsentry.YFilter }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) SetFilter(yf yfilter.YFilter) { sourceroutingstatsentry.YFilter = yf }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetGoName(yname string) string {
    if yname == "sourceRoutingStatsIfIndex" { return "Sourceroutingstatsifindex" }
    if yname == "sourceRoutingStatsRingNumber" { return "Sourceroutingstatsringnumber" }
    if yname == "sourceRoutingStatsInFrames" { return "Sourceroutingstatsinframes" }
    if yname == "sourceRoutingStatsOutFrames" { return "Sourceroutingstatsoutframes" }
    if yname == "sourceRoutingStatsThroughFrames" { return "Sourceroutingstatsthroughframes" }
    if yname == "sourceRoutingStatsAllRoutesBroadcastFrames" { return "Sourceroutingstatsallroutesbroadcastframes" }
    if yname == "sourceRoutingStatsSingleRouteBroadcastFrames" { return "Sourceroutingstatssingleroutebroadcastframes" }
    if yname == "sourceRoutingStatsInOctets" { return "Sourceroutingstatsinoctets" }
    if yname == "sourceRoutingStatsOutOctets" { return "Sourceroutingstatsoutoctets" }
    if yname == "sourceRoutingStatsThroughOctets" { return "Sourceroutingstatsthroughoctets" }
    if yname == "sourceRoutingStatsAllRoutesBroadcastOctets" { return "Sourceroutingstatsallroutesbroadcastoctets" }
    if yname == "sourceRoutingStatsSingleRoutesBroadcastOctets" { return "Sourceroutingstatssingleroutesbroadcastoctets" }
    if yname == "sourceRoutingStatsLocalLLCFrames" { return "Sourceroutingstatslocalllcframes" }
    if yname == "sourceRoutingStats1HopFrames" { return "Sourceroutingstats1Hopframes" }
    if yname == "sourceRoutingStats2HopsFrames" { return "Sourceroutingstats2Hopsframes" }
    if yname == "sourceRoutingStats3HopsFrames" { return "Sourceroutingstats3Hopsframes" }
    if yname == "sourceRoutingStats4HopsFrames" { return "Sourceroutingstats4Hopsframes" }
    if yname == "sourceRoutingStats5HopsFrames" { return "Sourceroutingstats5Hopsframes" }
    if yname == "sourceRoutingStats6HopsFrames" { return "Sourceroutingstats6Hopsframes" }
    if yname == "sourceRoutingStats7HopsFrames" { return "Sourceroutingstats7Hopsframes" }
    if yname == "sourceRoutingStats8HopsFrames" { return "Sourceroutingstats8Hopsframes" }
    if yname == "sourceRoutingStatsMoreThan8HopsFrames" { return "Sourceroutingstatsmorethan8Hopsframes" }
    if yname == "sourceRoutingStatsOwner" { return "Sourceroutingstatsowner" }
    if yname == "sourceRoutingStatsStatus" { return "Sourceroutingstatsstatus" }
    if yname == "sourceRoutingStatsDroppedFrames" { return "Sourceroutingstatsdroppedframes" }
    if yname == "sourceRoutingStatsCreateTime" { return "Sourceroutingstatscreatetime" }
    return ""
}

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetSegmentPath() string {
    return "sourceRoutingStatsEntry" + "[sourceRoutingStatsIfIndex='" + fmt.Sprintf("%v", sourceroutingstatsentry.Sourceroutingstatsifindex) + "']"
}

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sourceRoutingStatsIfIndex"] = sourceroutingstatsentry.Sourceroutingstatsifindex
    leafs["sourceRoutingStatsRingNumber"] = sourceroutingstatsentry.Sourceroutingstatsringnumber
    leafs["sourceRoutingStatsInFrames"] = sourceroutingstatsentry.Sourceroutingstatsinframes
    leafs["sourceRoutingStatsOutFrames"] = sourceroutingstatsentry.Sourceroutingstatsoutframes
    leafs["sourceRoutingStatsThroughFrames"] = sourceroutingstatsentry.Sourceroutingstatsthroughframes
    leafs["sourceRoutingStatsAllRoutesBroadcastFrames"] = sourceroutingstatsentry.Sourceroutingstatsallroutesbroadcastframes
    leafs["sourceRoutingStatsSingleRouteBroadcastFrames"] = sourceroutingstatsentry.Sourceroutingstatssingleroutebroadcastframes
    leafs["sourceRoutingStatsInOctets"] = sourceroutingstatsentry.Sourceroutingstatsinoctets
    leafs["sourceRoutingStatsOutOctets"] = sourceroutingstatsentry.Sourceroutingstatsoutoctets
    leafs["sourceRoutingStatsThroughOctets"] = sourceroutingstatsentry.Sourceroutingstatsthroughoctets
    leafs["sourceRoutingStatsAllRoutesBroadcastOctets"] = sourceroutingstatsentry.Sourceroutingstatsallroutesbroadcastoctets
    leafs["sourceRoutingStatsSingleRoutesBroadcastOctets"] = sourceroutingstatsentry.Sourceroutingstatssingleroutesbroadcastoctets
    leafs["sourceRoutingStatsLocalLLCFrames"] = sourceroutingstatsentry.Sourceroutingstatslocalllcframes
    leafs["sourceRoutingStats1HopFrames"] = sourceroutingstatsentry.Sourceroutingstats1Hopframes
    leafs["sourceRoutingStats2HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats2Hopsframes
    leafs["sourceRoutingStats3HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats3Hopsframes
    leafs["sourceRoutingStats4HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats4Hopsframes
    leafs["sourceRoutingStats5HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats5Hopsframes
    leafs["sourceRoutingStats6HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats6Hopsframes
    leafs["sourceRoutingStats7HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats7Hopsframes
    leafs["sourceRoutingStats8HopsFrames"] = sourceroutingstatsentry.Sourceroutingstats8Hopsframes
    leafs["sourceRoutingStatsMoreThan8HopsFrames"] = sourceroutingstatsentry.Sourceroutingstatsmorethan8Hopsframes
    leafs["sourceRoutingStatsOwner"] = sourceroutingstatsentry.Sourceroutingstatsowner
    leafs["sourceRoutingStatsStatus"] = sourceroutingstatsentry.Sourceroutingstatsstatus
    leafs["sourceRoutingStatsDroppedFrames"] = sourceroutingstatsentry.Sourceroutingstatsdroppedframes
    leafs["sourceRoutingStatsCreateTime"] = sourceroutingstatsentry.Sourceroutingstatscreatetime
    return leafs
}

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetYangName() string { return "sourceRoutingStatsEntry" }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) SetParent(parent types.Entity) { sourceroutingstatsentry.parent = parent }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetParent() types.Entity { return sourceroutingstatsentry.parent }

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetParentYangName() string { return "sourceRoutingStatsTable" }

