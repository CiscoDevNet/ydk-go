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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of Mac-Layer Token Ring statistics      entries.
    TokenRingMLStatsTable TOKENRINGRMONMIB_TokenRingMLStatsTable

    // A list of promiscuous Token Ring statistics entries.
    TokenRingPStatsTable TOKENRINGRMONMIB_TokenRingPStatsTable

    // A list of Mac-Layer Token Ring statistics      entries.
    TokenRingMLHistoryTable TOKENRINGRMONMIB_TokenRingMLHistoryTable

    // A list of promiscuous Token Ring statistics entries.
    TokenRingPHistoryTable TOKENRINGRMONMIB_TokenRingPHistoryTable

    // A list of ringStation table control entries.
    RingStationControlTable TOKENRINGRMONMIB_RingStationControlTable

    // A list of ring station entries.  An entry will exist for each station that
    // is now or has      previously been detected as physically present on this
    // ring.
    RingStationTable TOKENRINGRMONMIB_RingStationTable

    // A list of ring station entries for stations in the ring poll, ordered by
    // their ring-order.
    RingStationOrderTable TOKENRINGRMONMIB_RingStationOrderTable

    // A list of ring station configuration control entries.
    RingStationConfigControlTable TOKENRINGRMONMIB_RingStationConfigControlTable

    // A list of configuration entries for stations on a ring monitored by this
    // probe.
    RingStationConfigTable TOKENRINGRMONMIB_RingStationConfigTable

    // A list of source routing statistics entries.
    SourceRoutingStatsTable TOKENRINGRMONMIB_SourceRoutingStatsTable
}

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetEntityData() *types.CommonEntityData {
    tOKENRINGRMONMIB.EntityData.YFilter = tOKENRINGRMONMIB.YFilter
    tOKENRINGRMONMIB.EntityData.YangName = "TOKEN-RING-RMON-MIB"
    tOKENRINGRMONMIB.EntityData.BundleName = "cisco_ios_xe"
    tOKENRINGRMONMIB.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tOKENRINGRMONMIB.EntityData.SegmentPath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB"
    tOKENRINGRMONMIB.EntityData.AbsolutePath = tOKENRINGRMONMIB.EntityData.SegmentPath
    tOKENRINGRMONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tOKENRINGRMONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tOKENRINGRMONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tOKENRINGRMONMIB.EntityData.Children = types.NewOrderedMap()
    tOKENRINGRMONMIB.EntityData.Children.Append("tokenRingMLStatsTable", types.YChild{"TokenRingMLStatsTable", &tOKENRINGRMONMIB.TokenRingMLStatsTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("tokenRingPStatsTable", types.YChild{"TokenRingPStatsTable", &tOKENRINGRMONMIB.TokenRingPStatsTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("tokenRingMLHistoryTable", types.YChild{"TokenRingMLHistoryTable", &tOKENRINGRMONMIB.TokenRingMLHistoryTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("tokenRingPHistoryTable", types.YChild{"TokenRingPHistoryTable", &tOKENRINGRMONMIB.TokenRingPHistoryTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("ringStationControlTable", types.YChild{"RingStationControlTable", &tOKENRINGRMONMIB.RingStationControlTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("ringStationTable", types.YChild{"RingStationTable", &tOKENRINGRMONMIB.RingStationTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("ringStationOrderTable", types.YChild{"RingStationOrderTable", &tOKENRINGRMONMIB.RingStationOrderTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("ringStationConfigControlTable", types.YChild{"RingStationConfigControlTable", &tOKENRINGRMONMIB.RingStationConfigControlTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("ringStationConfigTable", types.YChild{"RingStationConfigTable", &tOKENRINGRMONMIB.RingStationConfigTable})
    tOKENRINGRMONMIB.EntityData.Children.Append("sourceRoutingStatsTable", types.YChild{"SourceRoutingStatsTable", &tOKENRINGRMONMIB.SourceRoutingStatsTable})
    tOKENRINGRMONMIB.EntityData.Leafs = types.NewOrderedMap()

    tOKENRINGRMONMIB.EntityData.YListKeys = []string {}

    return &(tOKENRINGRMONMIB.EntityData)
}

// TOKENRINGRMONMIB_TokenRingMLStatsTable
// A list of Mac-Layer Token Ring statistics
// 
// 
// 
// 
// 
// entries.
type TOKENRINGRMONMIB_TokenRingMLStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of Mac-Layer statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_TokenRingMLStatsTable_TokenRingMLStatsEntry.
    TokenRingMLStatsEntry []*TOKENRINGRMONMIB_TokenRingMLStatsTable_TokenRingMLStatsEntry
}

func (tokenRingMLStatsTable *TOKENRINGRMONMIB_TokenRingMLStatsTable) GetEntityData() *types.CommonEntityData {
    tokenRingMLStatsTable.EntityData.YFilter = tokenRingMLStatsTable.YFilter
    tokenRingMLStatsTable.EntityData.YangName = "tokenRingMLStatsTable"
    tokenRingMLStatsTable.EntityData.BundleName = "cisco_ios_xe"
    tokenRingMLStatsTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenRingMLStatsTable.EntityData.SegmentPath = "tokenRingMLStatsTable"
    tokenRingMLStatsTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + tokenRingMLStatsTable.EntityData.SegmentPath
    tokenRingMLStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingMLStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingMLStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingMLStatsTable.EntityData.Children = types.NewOrderedMap()
    tokenRingMLStatsTable.EntityData.Children.Append("tokenRingMLStatsEntry", types.YChild{"TokenRingMLStatsEntry", nil})
    for i := range tokenRingMLStatsTable.TokenRingMLStatsEntry {
        tokenRingMLStatsTable.EntityData.Children.Append(types.GetSegmentPath(tokenRingMLStatsTable.TokenRingMLStatsEntry[i]), types.YChild{"TokenRingMLStatsEntry", tokenRingMLStatsTable.TokenRingMLStatsEntry[i]})
    }
    tokenRingMLStatsTable.EntityData.Leafs = types.NewOrderedMap()

    tokenRingMLStatsTable.EntityData.YListKeys = []string {}

    return &(tokenRingMLStatsTable.EntityData)
}

// TOKENRINGRMONMIB_TokenRingMLStatsTable_TokenRingMLStatsEntry
// A collection of Mac-Layer statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_TokenRingMLStatsTable_TokenRingMLStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies this
    // tokenRingMLStats entry. The type is interface{} with range: 1..65535.
    TokenRingMLStatsIndex interface{}

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
    TokenRingMLStatsDataSource interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources. Note that this number is not necessarily the number
    // of packets dropped; it is just the number of times this condition has been
    // detected.  This value is the same as the corresponding
    // tokenRingPStatsDropEvents. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLStatsDropEvents interface{}

    // The total number of octets of data in MAC packets (excluding those that
    // were not good frames) received on the network (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsMacOctets interface{}

    // The total number of MAC packets (excluding packets that were not good
    // frames) received. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsMacPkts interface{}

    // The total number of times that the ring enters the ring purge state from
    // normal ring state.  The ring purge state that comes in response to the
    // claim token or beacon state is not counted. The type is interface{} with
    // range: 0..4294967295.
    TokenRingMLStatsRingPurgeEvents interface{}

    // The total number of ring purge MAC packets detected by probe. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLStatsRingPurgePkts interface{}

    // The total number of times that the ring enters a beaconing state
    // (beaconFrameStreamingState, beaconBitStreamingState,     
    // beaconSetRecoveryModeState, or beaconRingSignalLossState) from a
    // non-beaconing state.  Note that a change of the source address of the
    // beacon packet does not constitute a new beacon event. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLStatsBeaconEvents interface{}

    // The total amount of time that the ring has been in the beaconing state. The
    // type is interface{} with range: -2147483648..2147483647.
    TokenRingMLStatsBeaconTime interface{}

    // The total number of beacon MAC packets detected by the probe. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLStatsBeaconPkts interface{}

    // The total number of times that the ring enters the claim token state from
    // normal ring state or ring purge state.  The claim token state that comes in
    // response to a beacon state is not counted. The type is interface{} with
    // range: 0..4294967295.
    TokenRingMLStatsClaimTokenEvents interface{}

    // The total number of claim token MAC packets detected by the probe. The type
    // is interface{} with range: 0..4294967295.
    TokenRingMLStatsClaimTokenPkts interface{}

    // The total number of NAUN changes detected by the probe. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLStatsNAUNChanges interface{}

    // The total number of line errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsLineErrors interface{}

    // The total number of adapter internal errors reported in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLStatsInternalErrors interface{}

    // The total number of burst errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsBurstErrors interface{}

    // The total number of AC (Address Copied)  errors reported in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLStatsACErrors interface{}

    // The total number of abort delimiters reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsAbortErrors interface{}

    // The total number of lost frame errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsLostFrameErrors interface{}

    // The total number of receive congestion errors reported in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLStatsCongestionErrors interface{}

    // The total number of frame copied errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsFrameCopiedErrors interface{}

    // The total number of frequency errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsFrequencyErrors interface{}

    // The total number of token errors reported in error reporting packets
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    TokenRingMLStatsTokenErrors interface{}

    // The total number of soft error report frames detected by the probe. The
    // type is interface{} with range: 0..4294967295.
    TokenRingMLStatsSoftErrorReports interface{}

    // The total number of ring poll events detected by the probe (i.e. the number
    // of ring polls initiated by the active monitor that were detected). The type
    // is interface{} with range: 0..4294967295.
    TokenRingMLStatsRingPollEvents interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    TokenRingMLStatsOwner interface{}

    // The status of this tokenRingMLStats entry. The type is EntryStatus.
    TokenRingMLStatsStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    TokenRingMLStatsDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLStatsCreateTime interface{}
}

func (tokenRingMLStatsEntry *TOKENRINGRMONMIB_TokenRingMLStatsTable_TokenRingMLStatsEntry) GetEntityData() *types.CommonEntityData {
    tokenRingMLStatsEntry.EntityData.YFilter = tokenRingMLStatsEntry.YFilter
    tokenRingMLStatsEntry.EntityData.YangName = "tokenRingMLStatsEntry"
    tokenRingMLStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    tokenRingMLStatsEntry.EntityData.ParentYangName = "tokenRingMLStatsTable"
    tokenRingMLStatsEntry.EntityData.SegmentPath = "tokenRingMLStatsEntry" + types.AddKeyToken(tokenRingMLStatsEntry.TokenRingMLStatsIndex, "tokenRingMLStatsIndex")
    tokenRingMLStatsEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/tokenRingMLStatsTable/" + tokenRingMLStatsEntry.EntityData.SegmentPath
    tokenRingMLStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingMLStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingMLStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingMLStatsEntry.EntityData.Children = types.NewOrderedMap()
    tokenRingMLStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsIndex", types.YLeaf{"TokenRingMLStatsIndex", tokenRingMLStatsEntry.TokenRingMLStatsIndex})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsDataSource", types.YLeaf{"TokenRingMLStatsDataSource", tokenRingMLStatsEntry.TokenRingMLStatsDataSource})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsDropEvents", types.YLeaf{"TokenRingMLStatsDropEvents", tokenRingMLStatsEntry.TokenRingMLStatsDropEvents})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsMacOctets", types.YLeaf{"TokenRingMLStatsMacOctets", tokenRingMLStatsEntry.TokenRingMLStatsMacOctets})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsMacPkts", types.YLeaf{"TokenRingMLStatsMacPkts", tokenRingMLStatsEntry.TokenRingMLStatsMacPkts})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsRingPurgeEvents", types.YLeaf{"TokenRingMLStatsRingPurgeEvents", tokenRingMLStatsEntry.TokenRingMLStatsRingPurgeEvents})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsRingPurgePkts", types.YLeaf{"TokenRingMLStatsRingPurgePkts", tokenRingMLStatsEntry.TokenRingMLStatsRingPurgePkts})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsBeaconEvents", types.YLeaf{"TokenRingMLStatsBeaconEvents", tokenRingMLStatsEntry.TokenRingMLStatsBeaconEvents})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsBeaconTime", types.YLeaf{"TokenRingMLStatsBeaconTime", tokenRingMLStatsEntry.TokenRingMLStatsBeaconTime})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsBeaconPkts", types.YLeaf{"TokenRingMLStatsBeaconPkts", tokenRingMLStatsEntry.TokenRingMLStatsBeaconPkts})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsClaimTokenEvents", types.YLeaf{"TokenRingMLStatsClaimTokenEvents", tokenRingMLStatsEntry.TokenRingMLStatsClaimTokenEvents})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsClaimTokenPkts", types.YLeaf{"TokenRingMLStatsClaimTokenPkts", tokenRingMLStatsEntry.TokenRingMLStatsClaimTokenPkts})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsNAUNChanges", types.YLeaf{"TokenRingMLStatsNAUNChanges", tokenRingMLStatsEntry.TokenRingMLStatsNAUNChanges})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsLineErrors", types.YLeaf{"TokenRingMLStatsLineErrors", tokenRingMLStatsEntry.TokenRingMLStatsLineErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsInternalErrors", types.YLeaf{"TokenRingMLStatsInternalErrors", tokenRingMLStatsEntry.TokenRingMLStatsInternalErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsBurstErrors", types.YLeaf{"TokenRingMLStatsBurstErrors", tokenRingMLStatsEntry.TokenRingMLStatsBurstErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsACErrors", types.YLeaf{"TokenRingMLStatsACErrors", tokenRingMLStatsEntry.TokenRingMLStatsACErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsAbortErrors", types.YLeaf{"TokenRingMLStatsAbortErrors", tokenRingMLStatsEntry.TokenRingMLStatsAbortErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsLostFrameErrors", types.YLeaf{"TokenRingMLStatsLostFrameErrors", tokenRingMLStatsEntry.TokenRingMLStatsLostFrameErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsCongestionErrors", types.YLeaf{"TokenRingMLStatsCongestionErrors", tokenRingMLStatsEntry.TokenRingMLStatsCongestionErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsFrameCopiedErrors", types.YLeaf{"TokenRingMLStatsFrameCopiedErrors", tokenRingMLStatsEntry.TokenRingMLStatsFrameCopiedErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsFrequencyErrors", types.YLeaf{"TokenRingMLStatsFrequencyErrors", tokenRingMLStatsEntry.TokenRingMLStatsFrequencyErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsTokenErrors", types.YLeaf{"TokenRingMLStatsTokenErrors", tokenRingMLStatsEntry.TokenRingMLStatsTokenErrors})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsSoftErrorReports", types.YLeaf{"TokenRingMLStatsSoftErrorReports", tokenRingMLStatsEntry.TokenRingMLStatsSoftErrorReports})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsRingPollEvents", types.YLeaf{"TokenRingMLStatsRingPollEvents", tokenRingMLStatsEntry.TokenRingMLStatsRingPollEvents})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsOwner", types.YLeaf{"TokenRingMLStatsOwner", tokenRingMLStatsEntry.TokenRingMLStatsOwner})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsStatus", types.YLeaf{"TokenRingMLStatsStatus", tokenRingMLStatsEntry.TokenRingMLStatsStatus})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsDroppedFrames", types.YLeaf{"TokenRingMLStatsDroppedFrames", tokenRingMLStatsEntry.TokenRingMLStatsDroppedFrames})
    tokenRingMLStatsEntry.EntityData.Leafs.Append("tokenRingMLStatsCreateTime", types.YLeaf{"TokenRingMLStatsCreateTime", tokenRingMLStatsEntry.TokenRingMLStatsCreateTime})

    tokenRingMLStatsEntry.EntityData.YListKeys = []string {"TokenRingMLStatsIndex"}

    return &(tokenRingMLStatsEntry.EntityData)
}

// TOKENRINGRMONMIB_TokenRingPStatsTable
// A list of promiscuous Token Ring statistics
// entries.
type TOKENRINGRMONMIB_TokenRingPStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of promiscuous statistics kept for non-MAC packets on a
    // particular Token Ring interface. The type is slice of
    // TOKENRINGRMONMIB_TokenRingPStatsTable_TokenRingPStatsEntry.
    TokenRingPStatsEntry []*TOKENRINGRMONMIB_TokenRingPStatsTable_TokenRingPStatsEntry
}

func (tokenRingPStatsTable *TOKENRINGRMONMIB_TokenRingPStatsTable) GetEntityData() *types.CommonEntityData {
    tokenRingPStatsTable.EntityData.YFilter = tokenRingPStatsTable.YFilter
    tokenRingPStatsTable.EntityData.YangName = "tokenRingPStatsTable"
    tokenRingPStatsTable.EntityData.BundleName = "cisco_ios_xe"
    tokenRingPStatsTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenRingPStatsTable.EntityData.SegmentPath = "tokenRingPStatsTable"
    tokenRingPStatsTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + tokenRingPStatsTable.EntityData.SegmentPath
    tokenRingPStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingPStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingPStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingPStatsTable.EntityData.Children = types.NewOrderedMap()
    tokenRingPStatsTable.EntityData.Children.Append("tokenRingPStatsEntry", types.YChild{"TokenRingPStatsEntry", nil})
    for i := range tokenRingPStatsTable.TokenRingPStatsEntry {
        tokenRingPStatsTable.EntityData.Children.Append(types.GetSegmentPath(tokenRingPStatsTable.TokenRingPStatsEntry[i]), types.YChild{"TokenRingPStatsEntry", tokenRingPStatsTable.TokenRingPStatsEntry[i]})
    }
    tokenRingPStatsTable.EntityData.Leafs = types.NewOrderedMap()

    tokenRingPStatsTable.EntityData.YListKeys = []string {}

    return &(tokenRingPStatsTable.EntityData)
}

// TOKENRINGRMONMIB_TokenRingPStatsTable_TokenRingPStatsEntry
// A collection of promiscuous statistics kept for
// non-MAC packets on a particular Token Ring
// interface.
type TOKENRINGRMONMIB_TokenRingPStatsTable_TokenRingPStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies this
    // tokenRingPStats entry. The type is interface{} with range: 1..65535.
    TokenRingPStatsIndex interface{}

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
    TokenRingPStatsDataSource interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources. Note that this number is not necessarily the number
    // of packets dropped; it is just the number of times this condition has been
    // detected.  This value is the same as the corresponding
    // tokenRingMLStatsDropEvents. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPStatsDropEvents interface{}

    // The total number of octets of data in good frames received on the network
    // (excluding framing bits but including FCS octets) in non-MAC packets. The
    // type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataOctets interface{}

    // The total number of non-MAC packets in good frames.  received. The type is
    // interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts interface{}

    // The total number of good non-MAC frames received that were directed to an
    // LLC broadcast address (0xFFFFFFFFFFFF or 0xC000FFFFFFFF). The type is
    // interface{} with range: 0..4294967295.
    TokenRingPStatsDataBroadcastPkts interface{}

    // The total number of good non-MAC frames received that were directed to a
    // local or global multicast or functional address.  Note that this number
    // does not include packets directed to the broadcast address. The type is
    // interface{} with range: 0..4294967295.
    TokenRingPStatsDataMulticastPkts interface{}

    // The total number of good non-MAC frames received that were between 18 and
    // 63 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts18to63Octets interface{}

    // The total number of good non-MAC frames received that were between 64 and
    // 127 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts64to127Octets interface{}

    // The total number of good non-MAC frames received that were between 128 and
    // 255 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts128to255Octets interface{}

    // The total number of good non-MAC frames received that were between 256 and
    // 511 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts256to511Octets interface{}

    // The total number of good non-MAC frames received that were between 512 and
    // 1023 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts512to1023Octets interface{}

    // The total number of good non-MAC frames received that were between 1024 and
    // 2047 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts1024to2047Octets interface{}

    // The total number of good non-MAC frames received that were between 2048 and
    // 4095 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts2048to4095Octets interface{}

    // The total number of good non-MAC frames received that were between 4096 and
    // 8191 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts4096to8191Octets interface{}

    // The total number of good non-MAC frames received that were between 8192 and
    // 18000 octets in length inclusive, excluding framing bits but including FCS
    // octets. The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPkts8192to18000Octets interface{}

    // The total number of good non-MAC frames received that were greater than
    // 18000 octets in length, excluding framing bits but including FCS octets.
    // The type is interface{} with range: 0..4294967295.
    TokenRingPStatsDataPktsGreaterThan18000Octets interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    TokenRingPStatsOwner interface{}

    // The status of this tokenRingPStats entry. The type is EntryStatus.
    TokenRingPStatsStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    TokenRingPStatsDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPStatsCreateTime interface{}
}

func (tokenRingPStatsEntry *TOKENRINGRMONMIB_TokenRingPStatsTable_TokenRingPStatsEntry) GetEntityData() *types.CommonEntityData {
    tokenRingPStatsEntry.EntityData.YFilter = tokenRingPStatsEntry.YFilter
    tokenRingPStatsEntry.EntityData.YangName = "tokenRingPStatsEntry"
    tokenRingPStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    tokenRingPStatsEntry.EntityData.ParentYangName = "tokenRingPStatsTable"
    tokenRingPStatsEntry.EntityData.SegmentPath = "tokenRingPStatsEntry" + types.AddKeyToken(tokenRingPStatsEntry.TokenRingPStatsIndex, "tokenRingPStatsIndex")
    tokenRingPStatsEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/tokenRingPStatsTable/" + tokenRingPStatsEntry.EntityData.SegmentPath
    tokenRingPStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingPStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingPStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingPStatsEntry.EntityData.Children = types.NewOrderedMap()
    tokenRingPStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsIndex", types.YLeaf{"TokenRingPStatsIndex", tokenRingPStatsEntry.TokenRingPStatsIndex})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataSource", types.YLeaf{"TokenRingPStatsDataSource", tokenRingPStatsEntry.TokenRingPStatsDataSource})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDropEvents", types.YLeaf{"TokenRingPStatsDropEvents", tokenRingPStatsEntry.TokenRingPStatsDropEvents})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataOctets", types.YLeaf{"TokenRingPStatsDataOctets", tokenRingPStatsEntry.TokenRingPStatsDataOctets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts", types.YLeaf{"TokenRingPStatsDataPkts", tokenRingPStatsEntry.TokenRingPStatsDataPkts})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataBroadcastPkts", types.YLeaf{"TokenRingPStatsDataBroadcastPkts", tokenRingPStatsEntry.TokenRingPStatsDataBroadcastPkts})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataMulticastPkts", types.YLeaf{"TokenRingPStatsDataMulticastPkts", tokenRingPStatsEntry.TokenRingPStatsDataMulticastPkts})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts18to63Octets", types.YLeaf{"TokenRingPStatsDataPkts18to63Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts18to63Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts64to127Octets", types.YLeaf{"TokenRingPStatsDataPkts64to127Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts64to127Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts128to255Octets", types.YLeaf{"TokenRingPStatsDataPkts128to255Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts128to255Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts256to511Octets", types.YLeaf{"TokenRingPStatsDataPkts256to511Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts256to511Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts512to1023Octets", types.YLeaf{"TokenRingPStatsDataPkts512to1023Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts512to1023Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts1024to2047Octets", types.YLeaf{"TokenRingPStatsDataPkts1024to2047Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts1024to2047Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts2048to4095Octets", types.YLeaf{"TokenRingPStatsDataPkts2048to4095Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts2048to4095Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts4096to8191Octets", types.YLeaf{"TokenRingPStatsDataPkts4096to8191Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts4096to8191Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPkts8192to18000Octets", types.YLeaf{"TokenRingPStatsDataPkts8192to18000Octets", tokenRingPStatsEntry.TokenRingPStatsDataPkts8192to18000Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDataPktsGreaterThan18000Octets", types.YLeaf{"TokenRingPStatsDataPktsGreaterThan18000Octets", tokenRingPStatsEntry.TokenRingPStatsDataPktsGreaterThan18000Octets})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsOwner", types.YLeaf{"TokenRingPStatsOwner", tokenRingPStatsEntry.TokenRingPStatsOwner})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsStatus", types.YLeaf{"TokenRingPStatsStatus", tokenRingPStatsEntry.TokenRingPStatsStatus})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsDroppedFrames", types.YLeaf{"TokenRingPStatsDroppedFrames", tokenRingPStatsEntry.TokenRingPStatsDroppedFrames})
    tokenRingPStatsEntry.EntityData.Leafs.Append("tokenRingPStatsCreateTime", types.YLeaf{"TokenRingPStatsCreateTime", tokenRingPStatsEntry.TokenRingPStatsCreateTime})

    tokenRingPStatsEntry.EntityData.YListKeys = []string {"TokenRingPStatsIndex"}

    return &(tokenRingPStatsEntry.EntityData)
}

// TOKENRINGRMONMIB_TokenRingMLHistoryTable
// A list of Mac-Layer Token Ring statistics
// 
// 
// 
// 
// 
// entries.
type TOKENRINGRMONMIB_TokenRingMLHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of Mac-Layer statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_TokenRingMLHistoryTable_TokenRingMLHistoryEntry.
    TokenRingMLHistoryEntry []*TOKENRINGRMONMIB_TokenRingMLHistoryTable_TokenRingMLHistoryEntry
}

func (tokenRingMLHistoryTable *TOKENRINGRMONMIB_TokenRingMLHistoryTable) GetEntityData() *types.CommonEntityData {
    tokenRingMLHistoryTable.EntityData.YFilter = tokenRingMLHistoryTable.YFilter
    tokenRingMLHistoryTable.EntityData.YangName = "tokenRingMLHistoryTable"
    tokenRingMLHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    tokenRingMLHistoryTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenRingMLHistoryTable.EntityData.SegmentPath = "tokenRingMLHistoryTable"
    tokenRingMLHistoryTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + tokenRingMLHistoryTable.EntityData.SegmentPath
    tokenRingMLHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingMLHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingMLHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingMLHistoryTable.EntityData.Children = types.NewOrderedMap()
    tokenRingMLHistoryTable.EntityData.Children.Append("tokenRingMLHistoryEntry", types.YChild{"TokenRingMLHistoryEntry", nil})
    for i := range tokenRingMLHistoryTable.TokenRingMLHistoryEntry {
        tokenRingMLHistoryTable.EntityData.Children.Append(types.GetSegmentPath(tokenRingMLHistoryTable.TokenRingMLHistoryEntry[i]), types.YChild{"TokenRingMLHistoryEntry", tokenRingMLHistoryTable.TokenRingMLHistoryEntry[i]})
    }
    tokenRingMLHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    tokenRingMLHistoryTable.EntityData.YListKeys = []string {}

    return &(tokenRingMLHistoryTable.EntityData)
}

// TOKENRINGRMONMIB_TokenRingMLHistoryTable_TokenRingMLHistoryEntry
// A collection of Mac-Layer statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_TokenRingMLHistoryTable_TokenRingMLHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The history of which this entry is a part.  The
    // history identified by a particular value of this index is the same history
    // as identified by the same value of historyControlIndex. The type is
    // interface{} with range: 1..65535.
    TokenRingMLHistoryIndex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // Mac-Layer sample this entry represents among all Mac-Layer samples
    // associated with the same historyControlEntry.  This index starts at 1 and
    // increases by one as each new sample is taken. The type is interface{} with
    // range: -2147483648..2147483647.
    TokenRingMLHistorySampleIndex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval.  Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLHistoryIntervalStart interface{}

    // The total number of events in which packets were      dropped by the probe
    // due to lack of resources during this sampling interval.  Note that this
    // number is not necessarily the number of packets dropped, it is just the
    // number of times this condition has been detected. The type is interface{}
    // with range: 0..4294967295.
    TokenRingMLHistoryDropEvents interface{}

    // The total number of octets of data in MAC packets (excluding those that
    // were not good frames) received on the network during this sampling interval
    // (excluding framing bits but including FCS octets). The type is interface{}
    // with range: 0..4294967295.
    TokenRingMLHistoryMacOctets interface{}

    // The total number of MAC packets (excluding those that were not good frames)
    // received during this sampling interval. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLHistoryMacPkts interface{}

    // The total number of times that the ring entered the ring purge state from
    // normal ring state during this sampling interval.  The ring purge state that
    // comes from the claim token or beacon state is not counted. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryRingPurgeEvents interface{}

    // The total number of Ring Purge MAC packets detected by the probe during
    // this sampling      interval. The type is interface{} with range:
    // 0..4294967295.
    TokenRingMLHistoryRingPurgePkts interface{}

    // The total number of times that the ring enters a beaconing state
    // (beaconFrameStreamingState, beaconBitStreamingState,
    // beaconSetRecoveryModeState, or beaconRingSignalLossState) during this
    // sampling interval.  Note that a change of the source address of the beacon
    // packet does not constitute a new beacon event. The type is interface{} with
    // range: 0..4294967295.
    TokenRingMLHistoryBeaconEvents interface{}

    // The amount of time that the ring has been in the beaconing state during
    // this sampling interval. The type is interface{} with range:
    // -2147483648..2147483647.
    TokenRingMLHistoryBeaconTime interface{}

    // The total number of beacon MAC packets detected by the probe during this
    // sampling interval. The type is interface{} with range: 0..4294967295.
    TokenRingMLHistoryBeaconPkts interface{}

    // The total number of times that the ring enters the claim token state from
    // normal ring state or ring purge state during this sampling interval. The
    // claim token state that comes from the beacon state is not counted. The type
    // is interface{} with range: 0..4294967295.
    TokenRingMLHistoryClaimTokenEvents interface{}

    // The total number of claim token MAC packets detected by the probe during
    // this sampling interval. The type is interface{} with range: 0..4294967295.
    TokenRingMLHistoryClaimTokenPkts interface{}

    // The total number of NAUN changes detected by the probe during this sampling
    // interval. The type is interface{} with range: 0..4294967295.
    TokenRingMLHistoryNAUNChanges interface{}

    // The total number of line errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryLineErrors interface{}

    // The total number of adapter internal errors reported in error reporting
    // packets detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryInternalErrors interface{}

    // The total number of burst errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryBurstErrors interface{}

    // The total number of AC (Address Copied) errors reported in error reporting
    // packets detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryACErrors interface{}

    // The total number of abort delimiters reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryAbortErrors interface{}

    // The total number of lost frame errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryLostFrameErrors interface{}

    // The total number of receive congestion errors reported in error reporting
    // packets detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryCongestionErrors interface{}

    // The total number of frame copied errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryFrameCopiedErrors interface{}

    // The total number of frequency errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryFrequencyErrors interface{}

    // The total number of token errors reported in error reporting packets
    // detected by the probe during this sampling interval. The type is
    // interface{} with range: 0..4294967295.
    TokenRingMLHistoryTokenErrors interface{}

    // The total number of soft error report frames detected by the probe during
    // this sampling interval. The type is interface{} with range: 0..4294967295.
    TokenRingMLHistorySoftErrorReports interface{}

    // The total number of ring poll events detected by the probe during this
    // sampling interval. The type is interface{} with range: 0..4294967295.
    TokenRingMLHistoryRingPollEvents interface{}

    // The maximum number of active stations on the ring detected by the probe
    // during this sampling      interval. The type is interface{} with range:
    // -2147483648..2147483647.
    TokenRingMLHistoryActiveStations interface{}
}

func (tokenRingMLHistoryEntry *TOKENRINGRMONMIB_TokenRingMLHistoryTable_TokenRingMLHistoryEntry) GetEntityData() *types.CommonEntityData {
    tokenRingMLHistoryEntry.EntityData.YFilter = tokenRingMLHistoryEntry.YFilter
    tokenRingMLHistoryEntry.EntityData.YangName = "tokenRingMLHistoryEntry"
    tokenRingMLHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    tokenRingMLHistoryEntry.EntityData.ParentYangName = "tokenRingMLHistoryTable"
    tokenRingMLHistoryEntry.EntityData.SegmentPath = "tokenRingMLHistoryEntry" + types.AddKeyToken(tokenRingMLHistoryEntry.TokenRingMLHistoryIndex, "tokenRingMLHistoryIndex") + types.AddKeyToken(tokenRingMLHistoryEntry.TokenRingMLHistorySampleIndex, "tokenRingMLHistorySampleIndex")
    tokenRingMLHistoryEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/tokenRingMLHistoryTable/" + tokenRingMLHistoryEntry.EntityData.SegmentPath
    tokenRingMLHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingMLHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingMLHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingMLHistoryEntry.EntityData.Children = types.NewOrderedMap()
    tokenRingMLHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryIndex", types.YLeaf{"TokenRingMLHistoryIndex", tokenRingMLHistoryEntry.TokenRingMLHistoryIndex})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistorySampleIndex", types.YLeaf{"TokenRingMLHistorySampleIndex", tokenRingMLHistoryEntry.TokenRingMLHistorySampleIndex})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryIntervalStart", types.YLeaf{"TokenRingMLHistoryIntervalStart", tokenRingMLHistoryEntry.TokenRingMLHistoryIntervalStart})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryDropEvents", types.YLeaf{"TokenRingMLHistoryDropEvents", tokenRingMLHistoryEntry.TokenRingMLHistoryDropEvents})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryMacOctets", types.YLeaf{"TokenRingMLHistoryMacOctets", tokenRingMLHistoryEntry.TokenRingMLHistoryMacOctets})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryMacPkts", types.YLeaf{"TokenRingMLHistoryMacPkts", tokenRingMLHistoryEntry.TokenRingMLHistoryMacPkts})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryRingPurgeEvents", types.YLeaf{"TokenRingMLHistoryRingPurgeEvents", tokenRingMLHistoryEntry.TokenRingMLHistoryRingPurgeEvents})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryRingPurgePkts", types.YLeaf{"TokenRingMLHistoryRingPurgePkts", tokenRingMLHistoryEntry.TokenRingMLHistoryRingPurgePkts})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryBeaconEvents", types.YLeaf{"TokenRingMLHistoryBeaconEvents", tokenRingMLHistoryEntry.TokenRingMLHistoryBeaconEvents})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryBeaconTime", types.YLeaf{"TokenRingMLHistoryBeaconTime", tokenRingMLHistoryEntry.TokenRingMLHistoryBeaconTime})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryBeaconPkts", types.YLeaf{"TokenRingMLHistoryBeaconPkts", tokenRingMLHistoryEntry.TokenRingMLHistoryBeaconPkts})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryClaimTokenEvents", types.YLeaf{"TokenRingMLHistoryClaimTokenEvents", tokenRingMLHistoryEntry.TokenRingMLHistoryClaimTokenEvents})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryClaimTokenPkts", types.YLeaf{"TokenRingMLHistoryClaimTokenPkts", tokenRingMLHistoryEntry.TokenRingMLHistoryClaimTokenPkts})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryNAUNChanges", types.YLeaf{"TokenRingMLHistoryNAUNChanges", tokenRingMLHistoryEntry.TokenRingMLHistoryNAUNChanges})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryLineErrors", types.YLeaf{"TokenRingMLHistoryLineErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryLineErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryInternalErrors", types.YLeaf{"TokenRingMLHistoryInternalErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryInternalErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryBurstErrors", types.YLeaf{"TokenRingMLHistoryBurstErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryBurstErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryACErrors", types.YLeaf{"TokenRingMLHistoryACErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryACErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryAbortErrors", types.YLeaf{"TokenRingMLHistoryAbortErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryAbortErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryLostFrameErrors", types.YLeaf{"TokenRingMLHistoryLostFrameErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryLostFrameErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryCongestionErrors", types.YLeaf{"TokenRingMLHistoryCongestionErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryCongestionErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryFrameCopiedErrors", types.YLeaf{"TokenRingMLHistoryFrameCopiedErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryFrameCopiedErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryFrequencyErrors", types.YLeaf{"TokenRingMLHistoryFrequencyErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryFrequencyErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryTokenErrors", types.YLeaf{"TokenRingMLHistoryTokenErrors", tokenRingMLHistoryEntry.TokenRingMLHistoryTokenErrors})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistorySoftErrorReports", types.YLeaf{"TokenRingMLHistorySoftErrorReports", tokenRingMLHistoryEntry.TokenRingMLHistorySoftErrorReports})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryRingPollEvents", types.YLeaf{"TokenRingMLHistoryRingPollEvents", tokenRingMLHistoryEntry.TokenRingMLHistoryRingPollEvents})
    tokenRingMLHistoryEntry.EntityData.Leafs.Append("tokenRingMLHistoryActiveStations", types.YLeaf{"TokenRingMLHistoryActiveStations", tokenRingMLHistoryEntry.TokenRingMLHistoryActiveStations})

    tokenRingMLHistoryEntry.EntityData.YListKeys = []string {"TokenRingMLHistoryIndex", "TokenRingMLHistorySampleIndex"}

    return &(tokenRingMLHistoryEntry.EntityData)
}

// TOKENRINGRMONMIB_TokenRingPHistoryTable
// A list of promiscuous Token Ring statistics
// entries.
type TOKENRINGRMONMIB_TokenRingPHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of promiscuous statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_TokenRingPHistoryTable_TokenRingPHistoryEntry.
    TokenRingPHistoryEntry []*TOKENRINGRMONMIB_TokenRingPHistoryTable_TokenRingPHistoryEntry
}

func (tokenRingPHistoryTable *TOKENRINGRMONMIB_TokenRingPHistoryTable) GetEntityData() *types.CommonEntityData {
    tokenRingPHistoryTable.EntityData.YFilter = tokenRingPHistoryTable.YFilter
    tokenRingPHistoryTable.EntityData.YangName = "tokenRingPHistoryTable"
    tokenRingPHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    tokenRingPHistoryTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenRingPHistoryTable.EntityData.SegmentPath = "tokenRingPHistoryTable"
    tokenRingPHistoryTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + tokenRingPHistoryTable.EntityData.SegmentPath
    tokenRingPHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingPHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingPHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingPHistoryTable.EntityData.Children = types.NewOrderedMap()
    tokenRingPHistoryTable.EntityData.Children.Append("tokenRingPHistoryEntry", types.YChild{"TokenRingPHistoryEntry", nil})
    for i := range tokenRingPHistoryTable.TokenRingPHistoryEntry {
        tokenRingPHistoryTable.EntityData.Children.Append(types.GetSegmentPath(tokenRingPHistoryTable.TokenRingPHistoryEntry[i]), types.YChild{"TokenRingPHistoryEntry", tokenRingPHistoryTable.TokenRingPHistoryEntry[i]})
    }
    tokenRingPHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    tokenRingPHistoryTable.EntityData.YListKeys = []string {}

    return &(tokenRingPHistoryTable.EntityData)
}

// TOKENRINGRMONMIB_TokenRingPHistoryTable_TokenRingPHistoryEntry
// A collection of promiscuous statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_TokenRingPHistoryTable_TokenRingPHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The history of which this entry is a part.  The
    // history identified by a particular value of this index is the same history
    // as identified by the same value of historyControlIndex. The type is
    // interface{} with range: 1..65535.
    TokenRingPHistoryIndex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // sample this entry represents among all samples associated with the same
    // historyControlEntry. This index starts at 1 and increases by one as each
    // new sample is taken. The type is interface{} with range:
    // -2147483648..2147483647.
    TokenRingPHistorySampleIndex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval.  Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPHistoryIntervalStart interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources during this sampling interval.  Note that this number
    // is not necessarily the number of packets dropped, it is just the number of
    // times this condition has been detected. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPHistoryDropEvents interface{}

    // The total number of octets of data in good frames received on the network
    // (excluding framing bits but including FCS octets) in non-MAC packets during
    // this sampling interval. The type is interface{} with range: 0..4294967295.
    TokenRingPHistoryDataOctets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval. The type is interface{} with range: 0..4294967295.
    TokenRingPHistoryDataPkts interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were directed to an LLC broadcast address (0xFFFFFFFFFFFF or
    // 0xC000FFFFFFFF). The type is interface{} with range: 0..4294967295.
    TokenRingPHistoryDataBroadcastPkts interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were directed to a local or global multicast or functional
    // address.  Note that this number does not include packets directed to the
    // broadcast address. The type is interface{} with range: 0..4294967295.
    TokenRingPHistoryDataMulticastPkts interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 18 and 63 octets in length inclusive, excluding
    // framing bits but including FCS octets. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPHistoryDataPkts18to63Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 64 and 127 octets in length inclusive, excluding
    // framing bits but including FCS octets. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPHistoryDataPkts64to127Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 128 and 255 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts128to255Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between      256 and 511 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts256to511Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 512 and 1023 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts512to1023Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 1024 and 2047 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts1024to2047Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 2048 and 4095 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts2048to4095Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 4096 and 8191 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts4096to8191Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were between 8192 and 18000 octets in length inclusive,
    // excluding framing bits but including FCS octets. The type is interface{}
    // with range: 0..4294967295.
    TokenRingPHistoryDataPkts8192to18000Octets interface{}

    // The total number of good non-MAC frames received during this sampling
    // interval that were greater than 18000 octets in length, excluding framing
    // bits but including FCS octets. The type is interface{} with range:
    // 0..4294967295.
    TokenRingPHistoryDataPktsGreaterThan18000Octets interface{}
}

func (tokenRingPHistoryEntry *TOKENRINGRMONMIB_TokenRingPHistoryTable_TokenRingPHistoryEntry) GetEntityData() *types.CommonEntityData {
    tokenRingPHistoryEntry.EntityData.YFilter = tokenRingPHistoryEntry.YFilter
    tokenRingPHistoryEntry.EntityData.YangName = "tokenRingPHistoryEntry"
    tokenRingPHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    tokenRingPHistoryEntry.EntityData.ParentYangName = "tokenRingPHistoryTable"
    tokenRingPHistoryEntry.EntityData.SegmentPath = "tokenRingPHistoryEntry" + types.AddKeyToken(tokenRingPHistoryEntry.TokenRingPHistoryIndex, "tokenRingPHistoryIndex") + types.AddKeyToken(tokenRingPHistoryEntry.TokenRingPHistorySampleIndex, "tokenRingPHistorySampleIndex")
    tokenRingPHistoryEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/tokenRingPHistoryTable/" + tokenRingPHistoryEntry.EntityData.SegmentPath
    tokenRingPHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenRingPHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenRingPHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenRingPHistoryEntry.EntityData.Children = types.NewOrderedMap()
    tokenRingPHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryIndex", types.YLeaf{"TokenRingPHistoryIndex", tokenRingPHistoryEntry.TokenRingPHistoryIndex})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistorySampleIndex", types.YLeaf{"TokenRingPHistorySampleIndex", tokenRingPHistoryEntry.TokenRingPHistorySampleIndex})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryIntervalStart", types.YLeaf{"TokenRingPHistoryIntervalStart", tokenRingPHistoryEntry.TokenRingPHistoryIntervalStart})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDropEvents", types.YLeaf{"TokenRingPHistoryDropEvents", tokenRingPHistoryEntry.TokenRingPHistoryDropEvents})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataOctets", types.YLeaf{"TokenRingPHistoryDataOctets", tokenRingPHistoryEntry.TokenRingPHistoryDataOctets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts", types.YLeaf{"TokenRingPHistoryDataPkts", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataBroadcastPkts", types.YLeaf{"TokenRingPHistoryDataBroadcastPkts", tokenRingPHistoryEntry.TokenRingPHistoryDataBroadcastPkts})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataMulticastPkts", types.YLeaf{"TokenRingPHistoryDataMulticastPkts", tokenRingPHistoryEntry.TokenRingPHistoryDataMulticastPkts})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts18to63Octets", types.YLeaf{"TokenRingPHistoryDataPkts18to63Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts18to63Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts64to127Octets", types.YLeaf{"TokenRingPHistoryDataPkts64to127Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts64to127Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts128to255Octets", types.YLeaf{"TokenRingPHistoryDataPkts128to255Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts128to255Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts256to511Octets", types.YLeaf{"TokenRingPHistoryDataPkts256to511Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts256to511Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts512to1023Octets", types.YLeaf{"TokenRingPHistoryDataPkts512to1023Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts512to1023Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts1024to2047Octets", types.YLeaf{"TokenRingPHistoryDataPkts1024to2047Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts1024to2047Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts2048to4095Octets", types.YLeaf{"TokenRingPHistoryDataPkts2048to4095Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts2048to4095Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts4096to8191Octets", types.YLeaf{"TokenRingPHistoryDataPkts4096to8191Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts4096to8191Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPkts8192to18000Octets", types.YLeaf{"TokenRingPHistoryDataPkts8192to18000Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPkts8192to18000Octets})
    tokenRingPHistoryEntry.EntityData.Leafs.Append("tokenRingPHistoryDataPktsGreaterThan18000Octets", types.YLeaf{"TokenRingPHistoryDataPktsGreaterThan18000Octets", tokenRingPHistoryEntry.TokenRingPHistoryDataPktsGreaterThan18000Octets})

    tokenRingPHistoryEntry.EntityData.YListKeys = []string {"TokenRingPHistoryIndex", "TokenRingPHistorySampleIndex"}

    return &(tokenRingPHistoryEntry.EntityData)
}

// TOKENRINGRMONMIB_RingStationControlTable
// A list of ringStation table control entries.
type TOKENRINGRMONMIB_RingStationControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up the discovery of stations on a particular
    // interface and the collection of statistics about these stations. The type
    // is slice of
    // TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry.
    RingStationControlEntry []*TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry
}

func (ringStationControlTable *TOKENRINGRMONMIB_RingStationControlTable) GetEntityData() *types.CommonEntityData {
    ringStationControlTable.EntityData.YFilter = ringStationControlTable.YFilter
    ringStationControlTable.EntityData.YangName = "ringStationControlTable"
    ringStationControlTable.EntityData.BundleName = "cisco_ios_xe"
    ringStationControlTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringStationControlTable.EntityData.SegmentPath = "ringStationControlTable"
    ringStationControlTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + ringStationControlTable.EntityData.SegmentPath
    ringStationControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationControlTable.EntityData.Children = types.NewOrderedMap()
    ringStationControlTable.EntityData.Children.Append("ringStationControlEntry", types.YChild{"RingStationControlEntry", nil})
    for i := range ringStationControlTable.RingStationControlEntry {
        ringStationControlTable.EntityData.Children.Append(types.GetSegmentPath(ringStationControlTable.RingStationControlEntry[i]), types.YChild{"RingStationControlEntry", ringStationControlTable.RingStationControlEntry[i]})
    }
    ringStationControlTable.EntityData.Leafs = types.NewOrderedMap()

    ringStationControlTable.EntityData.YListKeys = []string {}

    return &(ringStationControlTable.EntityData)
}

// TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry
// A list of parameters that set up the discovery of
// stations on a particular interface and the
// collection of statistics about these stations.
type TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device from which ringStation
    // data is collected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB- II [3]. The type is interface{} with range:
    // 1..65535.
    RingStationControlIfIndex interface{}

    // The number of ringStationEntries in the ringStationTable associated with
    // this ringStationControlEntry. The type is interface{} with range:
    // -2147483648..2147483647.
    RingStationControlTableSize interface{}

    // The number of active ringStationEntries in the ringStationTable associated
    // with this ringStationControlEntry. The type is interface{} with range:
    // -2147483648..2147483647.
    RingStationControlActiveStations interface{}

    // The current status of this ring. The type is RingStationControlRingState.
    RingStationControlRingState interface{}

    // The address of the sender of the last beacon frame received by the probe on
    // this ring.  If no beacon frames have been received, this object shall be
    // equal to six octets of zero. The type is string with length: 6.
    RingStationControlBeaconSender interface{}

    // The address of the NAUN in the last beacon frame received by the probe on
    // this ring.  If no beacon frames have been received, this object shall be
    // equal to six octets of zero. The type is string with length: 6.
    RingStationControlBeaconNAUN interface{}

    // The address of the Active Monitor on this segment.  If this address is
    // unknown, this object shall be equal to six octets of zero. The type is
    // string with length: 6.
    RingStationControlActiveMonitor interface{}

    // The number of add and delete events in the ringStationOrderTable optionally
    // associated with this ringStationControlEntry. The type is interface{} with
    // range: 0..4294967295.
    RingStationControlOrderChanges interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    RingStationControlOwner interface{}

    // The status of this ringStationControl entry.  If this object is not equal
    // to valid(1), all associated entries in the ringStationTable shall be
    // deleted by the agent. The type is EntryStatus.
    RingStationControlStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    RingStationControlDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    RingStationControlCreateTime interface{}
}

func (ringStationControlEntry *TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry) GetEntityData() *types.CommonEntityData {
    ringStationControlEntry.EntityData.YFilter = ringStationControlEntry.YFilter
    ringStationControlEntry.EntityData.YangName = "ringStationControlEntry"
    ringStationControlEntry.EntityData.BundleName = "cisco_ios_xe"
    ringStationControlEntry.EntityData.ParentYangName = "ringStationControlTable"
    ringStationControlEntry.EntityData.SegmentPath = "ringStationControlEntry" + types.AddKeyToken(ringStationControlEntry.RingStationControlIfIndex, "ringStationControlIfIndex")
    ringStationControlEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/ringStationControlTable/" + ringStationControlEntry.EntityData.SegmentPath
    ringStationControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationControlEntry.EntityData.Children = types.NewOrderedMap()
    ringStationControlEntry.EntityData.Leafs = types.NewOrderedMap()
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlIfIndex", types.YLeaf{"RingStationControlIfIndex", ringStationControlEntry.RingStationControlIfIndex})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlTableSize", types.YLeaf{"RingStationControlTableSize", ringStationControlEntry.RingStationControlTableSize})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlActiveStations", types.YLeaf{"RingStationControlActiveStations", ringStationControlEntry.RingStationControlActiveStations})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlRingState", types.YLeaf{"RingStationControlRingState", ringStationControlEntry.RingStationControlRingState})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlBeaconSender", types.YLeaf{"RingStationControlBeaconSender", ringStationControlEntry.RingStationControlBeaconSender})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlBeaconNAUN", types.YLeaf{"RingStationControlBeaconNAUN", ringStationControlEntry.RingStationControlBeaconNAUN})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlActiveMonitor", types.YLeaf{"RingStationControlActiveMonitor", ringStationControlEntry.RingStationControlActiveMonitor})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlOrderChanges", types.YLeaf{"RingStationControlOrderChanges", ringStationControlEntry.RingStationControlOrderChanges})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlOwner", types.YLeaf{"RingStationControlOwner", ringStationControlEntry.RingStationControlOwner})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlStatus", types.YLeaf{"RingStationControlStatus", ringStationControlEntry.RingStationControlStatus})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlDroppedFrames", types.YLeaf{"RingStationControlDroppedFrames", ringStationControlEntry.RingStationControlDroppedFrames})
    ringStationControlEntry.EntityData.Leafs.Append("ringStationControlCreateTime", types.YLeaf{"RingStationControlCreateTime", ringStationControlEntry.RingStationControlCreateTime})

    ringStationControlEntry.EntityData.YListKeys = []string {"RingStationControlIfIndex"}

    return &(ringStationControlEntry.EntityData)
}

// TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState represents The current status of this ring.
type TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState string

const (
    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_normalOperation TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "normalOperation"

    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_ringPurgeState TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "ringPurgeState"

    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_claimTokenState TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "claimTokenState"

    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_beaconFrameStreamingState TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "beaconFrameStreamingState"

    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_beaconBitStreamingState TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "beaconBitStreamingState"

    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_beaconRingSignalLossState TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "beaconRingSignalLossState"

    TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState_beaconSetRecoveryModeState TOKENRINGRMONMIB_RingStationControlTable_RingStationControlEntry_RingStationControlRingState = "beaconSetRecoveryModeState"
)

// TOKENRINGRMONMIB_RingStationTable
// A list of ring station entries.  An entry will
// exist for each station that is now or has
// 
// 
// 
// 
// 
// previously been detected as physically present on
// this ring.
type TOKENRINGRMONMIB_RingStationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular station that has been
    // discovered on a ring monitored by this device. The type is slice of
    // TOKENRINGRMONMIB_RingStationTable_RingStationEntry.
    RingStationEntry []*TOKENRINGRMONMIB_RingStationTable_RingStationEntry
}

func (ringStationTable *TOKENRINGRMONMIB_RingStationTable) GetEntityData() *types.CommonEntityData {
    ringStationTable.EntityData.YFilter = ringStationTable.YFilter
    ringStationTable.EntityData.YangName = "ringStationTable"
    ringStationTable.EntityData.BundleName = "cisco_ios_xe"
    ringStationTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringStationTable.EntityData.SegmentPath = "ringStationTable"
    ringStationTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + ringStationTable.EntityData.SegmentPath
    ringStationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationTable.EntityData.Children = types.NewOrderedMap()
    ringStationTable.EntityData.Children.Append("ringStationEntry", types.YChild{"RingStationEntry", nil})
    for i := range ringStationTable.RingStationEntry {
        ringStationTable.EntityData.Children.Append(types.GetSegmentPath(ringStationTable.RingStationEntry[i]), types.YChild{"RingStationEntry", ringStationTable.RingStationEntry[i]})
    }
    ringStationTable.EntityData.Leafs = types.NewOrderedMap()

    ringStationTable.EntityData.YListKeys = []string {}

    return &(ringStationTable.EntityData)
}

// TOKENRINGRMONMIB_RingStationTable_RingStationEntry
// A collection of statistics for a particular
// station that has been discovered on a ring
// monitored by this device.
type TOKENRINGRMONMIB_RingStationTable_RingStationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    RingStationIfIndex interface{}

    // This attribute is a key. The physical address of this station. The type is
    // string with length: 6.
    RingStationMacAddress interface{}

    // The physical address of last known NAUN of this station. The type is string
    // with length: 6.
    RingStationLastNAUN interface{}

    // The status of this station on the ring. The type is
    // RingStationStationStatus.
    RingStationStationStatus interface{}

    // The value of sysUpTime at the time this station last entered the ring.  If
    // the time is unknown, this value shall be zero. The type is interface{} with
    // range: 0..4294967295.
    RingStationLastEnterTime interface{}

    // The value of sysUpTime at the time the probe detected that this station
    // last exited the ring. If the time is unknown, this value shall be zero. The
    // type is interface{} with range: 0..4294967295.
    RingStationLastExitTime interface{}

    // The number of times this station experienced a duplicate address error. The
    // type is interface{} with range: 0..4294967295.
    RingStationDuplicateAddresses interface{}

    // The total number of line errors reported by this station in error reporting
    // packets detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    RingStationInLineErrors interface{}

    // The total number of line errors reported in error reporting packets sent by
    // the nearest active downstream neighbor of this station and detected by the
    // probe. The type is interface{} with range: 0..4294967295.
    RingStationOutLineErrors interface{}

    // The total number of adapter internal errors reported by this station in
    // error reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationInternalErrors interface{}

    // The total number of burst errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationInBurstErrors interface{}

    // The total number of burst errors reported in error reporting packets sent
    // by the nearest active downstream neighbor of this station and detected by
    // the probe. The type is interface{} with range: 0..4294967295.
    RingStationOutBurstErrors interface{}

    // The total number of AC (Address Copied) errors reported in error reporting
    // packets sent by the nearest active downstream neighbor of this station and
    // detected by the probe. The type is interface{} with range: 0..4294967295.
    RingStationACErrors interface{}

    // The total number of abort delimiters reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationAbortErrors interface{}

    // The total number of lost frame errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationLostFrameErrors interface{}

    // The total number of receive congestion errors reported by this station in
    // error reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationCongestionErrors interface{}

    // The total number of frame copied errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationFrameCopiedErrors interface{}

    // The total number of frequency errors reported by this station in error
    // reporting packets detected by the probe. The type is interface{} with
    // range: 0..4294967295.
    RingStationFrequencyErrors interface{}

    // The total number of token errors reported by this station in error
    // reporting frames detected by the probe. The type is interface{} with range:
    // 0..4294967295.
    RingStationTokenErrors interface{}

    // The total number of beacon frames sent by this station and detected by the
    // probe. The type is interface{} with range: 0..4294967295.
    RingStationInBeaconErrors interface{}

    // The total number of beacon frames detected by the probe that name this
    // station as the NAUN. The type is interface{} with range: 0..4294967295.
    RingStationOutBeaconErrors interface{}

    // The number of times the probe detected this station inserting onto the
    // ring. The type is interface{} with range: 0..4294967295.
    RingStationInsertions interface{}
}

func (ringStationEntry *TOKENRINGRMONMIB_RingStationTable_RingStationEntry) GetEntityData() *types.CommonEntityData {
    ringStationEntry.EntityData.YFilter = ringStationEntry.YFilter
    ringStationEntry.EntityData.YangName = "ringStationEntry"
    ringStationEntry.EntityData.BundleName = "cisco_ios_xe"
    ringStationEntry.EntityData.ParentYangName = "ringStationTable"
    ringStationEntry.EntityData.SegmentPath = "ringStationEntry" + types.AddKeyToken(ringStationEntry.RingStationIfIndex, "ringStationIfIndex") + types.AddKeyToken(ringStationEntry.RingStationMacAddress, "ringStationMacAddress")
    ringStationEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/ringStationTable/" + ringStationEntry.EntityData.SegmentPath
    ringStationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationEntry.EntityData.Children = types.NewOrderedMap()
    ringStationEntry.EntityData.Leafs = types.NewOrderedMap()
    ringStationEntry.EntityData.Leafs.Append("ringStationIfIndex", types.YLeaf{"RingStationIfIndex", ringStationEntry.RingStationIfIndex})
    ringStationEntry.EntityData.Leafs.Append("ringStationMacAddress", types.YLeaf{"RingStationMacAddress", ringStationEntry.RingStationMacAddress})
    ringStationEntry.EntityData.Leafs.Append("ringStationLastNAUN", types.YLeaf{"RingStationLastNAUN", ringStationEntry.RingStationLastNAUN})
    ringStationEntry.EntityData.Leafs.Append("ringStationStationStatus", types.YLeaf{"RingStationStationStatus", ringStationEntry.RingStationStationStatus})
    ringStationEntry.EntityData.Leafs.Append("ringStationLastEnterTime", types.YLeaf{"RingStationLastEnterTime", ringStationEntry.RingStationLastEnterTime})
    ringStationEntry.EntityData.Leafs.Append("ringStationLastExitTime", types.YLeaf{"RingStationLastExitTime", ringStationEntry.RingStationLastExitTime})
    ringStationEntry.EntityData.Leafs.Append("ringStationDuplicateAddresses", types.YLeaf{"RingStationDuplicateAddresses", ringStationEntry.RingStationDuplicateAddresses})
    ringStationEntry.EntityData.Leafs.Append("ringStationInLineErrors", types.YLeaf{"RingStationInLineErrors", ringStationEntry.RingStationInLineErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationOutLineErrors", types.YLeaf{"RingStationOutLineErrors", ringStationEntry.RingStationOutLineErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationInternalErrors", types.YLeaf{"RingStationInternalErrors", ringStationEntry.RingStationInternalErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationInBurstErrors", types.YLeaf{"RingStationInBurstErrors", ringStationEntry.RingStationInBurstErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationOutBurstErrors", types.YLeaf{"RingStationOutBurstErrors", ringStationEntry.RingStationOutBurstErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationACErrors", types.YLeaf{"RingStationACErrors", ringStationEntry.RingStationACErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationAbortErrors", types.YLeaf{"RingStationAbortErrors", ringStationEntry.RingStationAbortErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationLostFrameErrors", types.YLeaf{"RingStationLostFrameErrors", ringStationEntry.RingStationLostFrameErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationCongestionErrors", types.YLeaf{"RingStationCongestionErrors", ringStationEntry.RingStationCongestionErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationFrameCopiedErrors", types.YLeaf{"RingStationFrameCopiedErrors", ringStationEntry.RingStationFrameCopiedErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationFrequencyErrors", types.YLeaf{"RingStationFrequencyErrors", ringStationEntry.RingStationFrequencyErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationTokenErrors", types.YLeaf{"RingStationTokenErrors", ringStationEntry.RingStationTokenErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationInBeaconErrors", types.YLeaf{"RingStationInBeaconErrors", ringStationEntry.RingStationInBeaconErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationOutBeaconErrors", types.YLeaf{"RingStationOutBeaconErrors", ringStationEntry.RingStationOutBeaconErrors})
    ringStationEntry.EntityData.Leafs.Append("ringStationInsertions", types.YLeaf{"RingStationInsertions", ringStationEntry.RingStationInsertions})

    ringStationEntry.EntityData.YListKeys = []string {"RingStationIfIndex", "RingStationMacAddress"}

    return &(ringStationEntry.EntityData)
}

// TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus represents The status of this station on the ring.
type TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus string

const (
    TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus_active TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus = "active"

    TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus_inactive TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus = "inactive"

    TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus_forcedRemoval TOKENRINGRMONMIB_RingStationTable_RingStationEntry_RingStationStationStatus = "forcedRemoval"
)

// TOKENRINGRMONMIB_RingStationOrderTable
// A list of ring station entries for stations in
// the ring poll, ordered by their ring-order.
type TOKENRINGRMONMIB_RingStationOrderTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular      station that is active on
    // a ring monitored by this device.  This table will contain information for
    // every interface that has a ringStationControlStatus equal to valid. The
    // type is slice of
    // TOKENRINGRMONMIB_RingStationOrderTable_RingStationOrderEntry.
    RingStationOrderEntry []*TOKENRINGRMONMIB_RingStationOrderTable_RingStationOrderEntry
}

func (ringStationOrderTable *TOKENRINGRMONMIB_RingStationOrderTable) GetEntityData() *types.CommonEntityData {
    ringStationOrderTable.EntityData.YFilter = ringStationOrderTable.YFilter
    ringStationOrderTable.EntityData.YangName = "ringStationOrderTable"
    ringStationOrderTable.EntityData.BundleName = "cisco_ios_xe"
    ringStationOrderTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringStationOrderTable.EntityData.SegmentPath = "ringStationOrderTable"
    ringStationOrderTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + ringStationOrderTable.EntityData.SegmentPath
    ringStationOrderTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationOrderTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationOrderTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationOrderTable.EntityData.Children = types.NewOrderedMap()
    ringStationOrderTable.EntityData.Children.Append("ringStationOrderEntry", types.YChild{"RingStationOrderEntry", nil})
    for i := range ringStationOrderTable.RingStationOrderEntry {
        ringStationOrderTable.EntityData.Children.Append(types.GetSegmentPath(ringStationOrderTable.RingStationOrderEntry[i]), types.YChild{"RingStationOrderEntry", ringStationOrderTable.RingStationOrderEntry[i]})
    }
    ringStationOrderTable.EntityData.Leafs = types.NewOrderedMap()

    ringStationOrderTable.EntityData.YListKeys = []string {}

    return &(ringStationOrderTable.EntityData)
}

// TOKENRINGRMONMIB_RingStationOrderTable_RingStationOrderEntry
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
type TOKENRINGRMONMIB_RingStationOrderTable_RingStationOrderEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    RingStationOrderIfIndex interface{}

    // This attribute is a key. This index denotes the location of this station
    // with respect to other stations on the ring.  This index is one more than
    // the number of hops downstream that this station is from the rmon probe. 
    // The rmon probe itself gets the value one. The type is interface{} with
    // range: -2147483648..2147483647.
    RingStationOrderOrderIndex interface{}

    // The physical address of this station. The type is string with length: 6.
    RingStationOrderMacAddress interface{}
}

func (ringStationOrderEntry *TOKENRINGRMONMIB_RingStationOrderTable_RingStationOrderEntry) GetEntityData() *types.CommonEntityData {
    ringStationOrderEntry.EntityData.YFilter = ringStationOrderEntry.YFilter
    ringStationOrderEntry.EntityData.YangName = "ringStationOrderEntry"
    ringStationOrderEntry.EntityData.BundleName = "cisco_ios_xe"
    ringStationOrderEntry.EntityData.ParentYangName = "ringStationOrderTable"
    ringStationOrderEntry.EntityData.SegmentPath = "ringStationOrderEntry" + types.AddKeyToken(ringStationOrderEntry.RingStationOrderIfIndex, "ringStationOrderIfIndex") + types.AddKeyToken(ringStationOrderEntry.RingStationOrderOrderIndex, "ringStationOrderOrderIndex")
    ringStationOrderEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/ringStationOrderTable/" + ringStationOrderEntry.EntityData.SegmentPath
    ringStationOrderEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationOrderEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationOrderEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationOrderEntry.EntityData.Children = types.NewOrderedMap()
    ringStationOrderEntry.EntityData.Leafs = types.NewOrderedMap()
    ringStationOrderEntry.EntityData.Leafs.Append("ringStationOrderIfIndex", types.YLeaf{"RingStationOrderIfIndex", ringStationOrderEntry.RingStationOrderIfIndex})
    ringStationOrderEntry.EntityData.Leafs.Append("ringStationOrderOrderIndex", types.YLeaf{"RingStationOrderOrderIndex", ringStationOrderEntry.RingStationOrderOrderIndex})
    ringStationOrderEntry.EntityData.Leafs.Append("ringStationOrderMacAddress", types.YLeaf{"RingStationOrderMacAddress", ringStationOrderEntry.RingStationOrderMacAddress})

    ringStationOrderEntry.EntityData.YListKeys = []string {"RingStationOrderIfIndex", "RingStationOrderOrderIndex"}

    return &(ringStationOrderEntry.EntityData)
}

// TOKENRINGRMONMIB_RingStationConfigControlTable
// A list of ring station configuration control
// entries.
type TOKENRINGRMONMIB_RingStationConfigControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry controls active management of stations by the probe.  One entry
    // exists in this table for each active station in the ringStationTable. The
    // type is slice of
    // TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry.
    RingStationConfigControlEntry []*TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry
}

func (ringStationConfigControlTable *TOKENRINGRMONMIB_RingStationConfigControlTable) GetEntityData() *types.CommonEntityData {
    ringStationConfigControlTable.EntityData.YFilter = ringStationConfigControlTable.YFilter
    ringStationConfigControlTable.EntityData.YangName = "ringStationConfigControlTable"
    ringStationConfigControlTable.EntityData.BundleName = "cisco_ios_xe"
    ringStationConfigControlTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringStationConfigControlTable.EntityData.SegmentPath = "ringStationConfigControlTable"
    ringStationConfigControlTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + ringStationConfigControlTable.EntityData.SegmentPath
    ringStationConfigControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationConfigControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationConfigControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationConfigControlTable.EntityData.Children = types.NewOrderedMap()
    ringStationConfigControlTable.EntityData.Children.Append("ringStationConfigControlEntry", types.YChild{"RingStationConfigControlEntry", nil})
    for i := range ringStationConfigControlTable.RingStationConfigControlEntry {
        ringStationConfigControlTable.EntityData.Children.Append(types.GetSegmentPath(ringStationConfigControlTable.RingStationConfigControlEntry[i]), types.YChild{"RingStationConfigControlEntry", ringStationConfigControlTable.RingStationConfigControlEntry[i]})
    }
    ringStationConfigControlTable.EntityData.Leafs = types.NewOrderedMap()

    ringStationConfigControlTable.EntityData.YListKeys = []string {}

    return &(ringStationConfigControlTable.EntityData)
}

// TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry
// This entry controls active management of stations
// by the probe.  One entry exists in this table for
// each active station in the ringStationTable.
type TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies the  
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    RingStationConfigControlIfIndex interface{}

    // This attribute is a key. The physical address of this station. The type is
    // string with length: 6.
    RingStationConfigControlMacAddress interface{}

    // Setting this object to `removing(2)' causes a Remove Station MAC frame to
    // be sent.  The agent will set this object to `stable(1)' after processing
    // the request. The type is RingStationConfigControlRemove.
    RingStationConfigControlRemove interface{}

    // Setting this object to `updating(2)' causes the configuration information
    // associate with this entry to be updated.  The agent will set this object to
    // `stable(1)' after processing the request. The type is
    // RingStationConfigControlUpdateStats.
    RingStationConfigControlUpdateStats interface{}
}

func (ringStationConfigControlEntry *TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry) GetEntityData() *types.CommonEntityData {
    ringStationConfigControlEntry.EntityData.YFilter = ringStationConfigControlEntry.YFilter
    ringStationConfigControlEntry.EntityData.YangName = "ringStationConfigControlEntry"
    ringStationConfigControlEntry.EntityData.BundleName = "cisco_ios_xe"
    ringStationConfigControlEntry.EntityData.ParentYangName = "ringStationConfigControlTable"
    ringStationConfigControlEntry.EntityData.SegmentPath = "ringStationConfigControlEntry" + types.AddKeyToken(ringStationConfigControlEntry.RingStationConfigControlIfIndex, "ringStationConfigControlIfIndex") + types.AddKeyToken(ringStationConfigControlEntry.RingStationConfigControlMacAddress, "ringStationConfigControlMacAddress")
    ringStationConfigControlEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/ringStationConfigControlTable/" + ringStationConfigControlEntry.EntityData.SegmentPath
    ringStationConfigControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationConfigControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationConfigControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationConfigControlEntry.EntityData.Children = types.NewOrderedMap()
    ringStationConfigControlEntry.EntityData.Leafs = types.NewOrderedMap()
    ringStationConfigControlEntry.EntityData.Leafs.Append("ringStationConfigControlIfIndex", types.YLeaf{"RingStationConfigControlIfIndex", ringStationConfigControlEntry.RingStationConfigControlIfIndex})
    ringStationConfigControlEntry.EntityData.Leafs.Append("ringStationConfigControlMacAddress", types.YLeaf{"RingStationConfigControlMacAddress", ringStationConfigControlEntry.RingStationConfigControlMacAddress})
    ringStationConfigControlEntry.EntityData.Leafs.Append("ringStationConfigControlRemove", types.YLeaf{"RingStationConfigControlRemove", ringStationConfigControlEntry.RingStationConfigControlRemove})
    ringStationConfigControlEntry.EntityData.Leafs.Append("ringStationConfigControlUpdateStats", types.YLeaf{"RingStationConfigControlUpdateStats", ringStationConfigControlEntry.RingStationConfigControlUpdateStats})

    ringStationConfigControlEntry.EntityData.YListKeys = []string {"RingStationConfigControlIfIndex", "RingStationConfigControlMacAddress"}

    return &(ringStationConfigControlEntry.EntityData)
}

// TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlRemove represents processing the request.
type TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlRemove string

const (
    TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlRemove_stable TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlRemove = "stable"

    TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlRemove_removing TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlRemove = "removing"
)

// TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlUpdateStats represents request.
type TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlUpdateStats string

const (
    TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlUpdateStats_stable TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlUpdateStats = "stable"

    TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlUpdateStats_updating TOKENRINGRMONMIB_RingStationConfigControlTable_RingStationConfigControlEntry_RingStationConfigControlUpdateStats = "updating"
)

// TOKENRINGRMONMIB_RingStationConfigTable
// A list of configuration entries for stations on a
// ring monitored by this probe.
type TOKENRINGRMONMIB_RingStationConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular station that has been
    // discovered on a ring monitored by this probe. The type is slice of
    // TOKENRINGRMONMIB_RingStationConfigTable_RingStationConfigEntry.
    RingStationConfigEntry []*TOKENRINGRMONMIB_RingStationConfigTable_RingStationConfigEntry
}

func (ringStationConfigTable *TOKENRINGRMONMIB_RingStationConfigTable) GetEntityData() *types.CommonEntityData {
    ringStationConfigTable.EntityData.YFilter = ringStationConfigTable.YFilter
    ringStationConfigTable.EntityData.YangName = "ringStationConfigTable"
    ringStationConfigTable.EntityData.BundleName = "cisco_ios_xe"
    ringStationConfigTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringStationConfigTable.EntityData.SegmentPath = "ringStationConfigTable"
    ringStationConfigTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + ringStationConfigTable.EntityData.SegmentPath
    ringStationConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationConfigTable.EntityData.Children = types.NewOrderedMap()
    ringStationConfigTable.EntityData.Children.Append("ringStationConfigEntry", types.YChild{"RingStationConfigEntry", nil})
    for i := range ringStationConfigTable.RingStationConfigEntry {
        ringStationConfigTable.EntityData.Children.Append(types.GetSegmentPath(ringStationConfigTable.RingStationConfigEntry[i]), types.YChild{"RingStationConfigEntry", ringStationConfigTable.RingStationConfigEntry[i]})
    }
    ringStationConfigTable.EntityData.Leafs = types.NewOrderedMap()

    ringStationConfigTable.EntityData.YListKeys = []string {}

    return &(ringStationConfigTable.EntityData)
}

// TOKENRINGRMONMIB_RingStationConfigTable_RingStationConfigEntry
// A collection of statistics for a particular
// station that has been discovered on a ring
// monitored by this probe.
type TOKENRINGRMONMIB_RingStationConfigTable_RingStationConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies the  
    // interface on this remote network monitoring device on which this station
    // was detected.  The interface identified by a particular value of this
    // object is the same interface as identified by the same value of the ifIndex
    // object, defined in MIB-II [3]. The type is interface{} with range:
    // -2147483648..2147483647.
    RingStationConfigIfIndex interface{}

    // This attribute is a key. The physical address of this station. The type is
    // string with length: 6.
    RingStationConfigMacAddress interface{}

    // The value of sysUpTime at the time this configuration information was last
    // updated (completely). The type is interface{} with range: 0..4294967295.
    RingStationConfigUpdateTime interface{}

    // The assigned physical location of this station. The type is string with
    // length: 4.
    RingStationConfigLocation interface{}

    // The microcode EC level of this station. The type is string with length: 10.
    RingStationConfigMicrocode interface{}

    // The low-order 4 octets of the group address recognized by this station. The
    // type is string with length: 4.
    RingStationConfigGroupAddress interface{}

    // the functional addresses recognized by this station. The type is string
    // with length: 4.
    RingStationConfigFunctionalAddress interface{}
}

func (ringStationConfigEntry *TOKENRINGRMONMIB_RingStationConfigTable_RingStationConfigEntry) GetEntityData() *types.CommonEntityData {
    ringStationConfigEntry.EntityData.YFilter = ringStationConfigEntry.YFilter
    ringStationConfigEntry.EntityData.YangName = "ringStationConfigEntry"
    ringStationConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    ringStationConfigEntry.EntityData.ParentYangName = "ringStationConfigTable"
    ringStationConfigEntry.EntityData.SegmentPath = "ringStationConfigEntry" + types.AddKeyToken(ringStationConfigEntry.RingStationConfigIfIndex, "ringStationConfigIfIndex") + types.AddKeyToken(ringStationConfigEntry.RingStationConfigMacAddress, "ringStationConfigMacAddress")
    ringStationConfigEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/ringStationConfigTable/" + ringStationConfigEntry.EntityData.SegmentPath
    ringStationConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringStationConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringStationConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringStationConfigEntry.EntityData.Children = types.NewOrderedMap()
    ringStationConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigIfIndex", types.YLeaf{"RingStationConfigIfIndex", ringStationConfigEntry.RingStationConfigIfIndex})
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigMacAddress", types.YLeaf{"RingStationConfigMacAddress", ringStationConfigEntry.RingStationConfigMacAddress})
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigUpdateTime", types.YLeaf{"RingStationConfigUpdateTime", ringStationConfigEntry.RingStationConfigUpdateTime})
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigLocation", types.YLeaf{"RingStationConfigLocation", ringStationConfigEntry.RingStationConfigLocation})
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigMicrocode", types.YLeaf{"RingStationConfigMicrocode", ringStationConfigEntry.RingStationConfigMicrocode})
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigGroupAddress", types.YLeaf{"RingStationConfigGroupAddress", ringStationConfigEntry.RingStationConfigGroupAddress})
    ringStationConfigEntry.EntityData.Leafs.Append("ringStationConfigFunctionalAddress", types.YLeaf{"RingStationConfigFunctionalAddress", ringStationConfigEntry.RingStationConfigFunctionalAddress})

    ringStationConfigEntry.EntityData.YListKeys = []string {"RingStationConfigIfIndex", "RingStationConfigMacAddress"}

    return &(ringStationConfigEntry.EntityData)
}

// TOKENRINGRMONMIB_SourceRoutingStatsTable
// A list of source routing statistics entries.
type TOKENRINGRMONMIB_SourceRoutingStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of source routing statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_SourceRoutingStatsTable_SourceRoutingStatsEntry.
    SourceRoutingStatsEntry []*TOKENRINGRMONMIB_SourceRoutingStatsTable_SourceRoutingStatsEntry
}

func (sourceRoutingStatsTable *TOKENRINGRMONMIB_SourceRoutingStatsTable) GetEntityData() *types.CommonEntityData {
    sourceRoutingStatsTable.EntityData.YFilter = sourceRoutingStatsTable.YFilter
    sourceRoutingStatsTable.EntityData.YangName = "sourceRoutingStatsTable"
    sourceRoutingStatsTable.EntityData.BundleName = "cisco_ios_xe"
    sourceRoutingStatsTable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    sourceRoutingStatsTable.EntityData.SegmentPath = "sourceRoutingStatsTable"
    sourceRoutingStatsTable.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/" + sourceRoutingStatsTable.EntityData.SegmentPath
    sourceRoutingStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sourceRoutingStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sourceRoutingStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sourceRoutingStatsTable.EntityData.Children = types.NewOrderedMap()
    sourceRoutingStatsTable.EntityData.Children.Append("sourceRoutingStatsEntry", types.YChild{"SourceRoutingStatsEntry", nil})
    for i := range sourceRoutingStatsTable.SourceRoutingStatsEntry {
        sourceRoutingStatsTable.EntityData.Children.Append(types.GetSegmentPath(sourceRoutingStatsTable.SourceRoutingStatsEntry[i]), types.YChild{"SourceRoutingStatsEntry", sourceRoutingStatsTable.SourceRoutingStatsEntry[i]})
    }
    sourceRoutingStatsTable.EntityData.Leafs = types.NewOrderedMap()

    sourceRoutingStatsTable.EntityData.YListKeys = []string {}

    return &(sourceRoutingStatsTable.EntityData)
}

// TOKENRINGRMONMIB_SourceRoutingStatsTable_SourceRoutingStatsEntry
// A collection of source routing statistics kept
// for a particular Token Ring interface.
type TOKENRINGRMONMIB_SourceRoutingStatsTable_SourceRoutingStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object uniquely identifies the
    // interface on this remote network monitoring device on which source routing
    // statistics will be detected.  The interface identified by a particular
    // value of this object is the same interface as identified by the same value
    // of the ifIndex object, defined in MIB-II [3]. The type is interface{} with
    // range: -2147483648..2147483647.
    SourceRoutingStatsIfIndex interface{}

    // The ring number of the ring monitored by this entry.  When any object in
    // this entry is created, the probe will attempt to discover the ring number. 
    // Only after the ring number is discovered will this object be created. 
    // After creating an object in this entry, the management station should poll
    // this object to detect when it is created.  Only after this object is
    // created can the management station set the sourceRoutingStatsStatus entry
    // to valid(1). The type is interface{} with range: -2147483648..2147483647.
    SourceRoutingStatsRingNumber interface{}

    // The count of frames sent into this ring from another ring. The type is
    // interface{} with range: 0..4294967295.
    SourceRoutingStatsInFrames interface{}

    // The count of frames sent from this ring to another ring. The type is
    // interface{} with range: 0..4294967295.
    SourceRoutingStatsOutFrames interface{}

    // The count of frames sent from another ring, through this ring, to another
    // ring. The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsThroughFrames interface{}

    // The total number of good frames received that were All Routes Broadcast.
    // The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsAllRoutesBroadcastFrames interface{}

    // The total number of good frames received that were Single Route Broadcast.
    // The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsSingleRouteBroadcastFrames interface{}

    // The count of octets in good frames sent into this ring from another ring.
    // The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsInOctets interface{}

    // The count of octets in good frames sent from this ring to another ring. The
    // type is interface{} with range: 0..4294967295.
    SourceRoutingStatsOutOctets interface{}

    // The count of octets in good frames sent another ring, through this ring, to
    // another ring. The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsThroughOctets interface{}

    // The total number of octets in good frames received that were All Routes
    // Broadcast. The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsAllRoutesBroadcastOctets interface{}

    // The total number of octets in good frames received that were Single Route
    // Broadcast. The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsSingleRoutesBroadcastOctets interface{}

    // The total number of frames received who had no RIF field (or had a RIF
    // field that only included the local ring's number) and were not All Route
    // Broadcast Frames. The type is interface{} with range: 0..4294967295.
    SourceRoutingStatsLocalLLCFrames interface{}

    // The total number of frames received whose route had 1 hop, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats1HopFrames interface{}

    // The total number of frames received whose route had 2 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats2HopsFrames interface{}

    // The total number of frames received whose route had 3 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats3HopsFrames interface{}

    // The total number of frames received whose route had 4 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats4HopsFrames interface{}

    // The total number of frames received whose route had 5 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats5HopsFrames interface{}

    // The total number of frames received whose route had 6 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats6HopsFrames interface{}

    // The total number of frames received whose route had 7 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats7HopsFrames interface{}

    // The total number of frames received whose route had 8 hops, were not All
    // Route Broadcast Frames, and whose source or destination were on this ring
    // (i.e. frames that had a RIF field and had this ring number in the first or
    // last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStats8HopsFrames interface{}

    // The total number of frames received whose route had more than 8 hops, were
    // not All Route Broadcast Frames, and whose source or destination were on
    // this ring (i.e. frames that had a RIF field and had this ring number in the
    // first or last entry of the RIF field). The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStatsMoreThan8HopsFrames interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string.
    SourceRoutingStatsOwner interface{}

    // The status of this sourceRoutingStats entry. The type is EntryStatus.
    SourceRoutingStatsStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    SourceRoutingStatsDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    SourceRoutingStatsCreateTime interface{}
}

func (sourceRoutingStatsEntry *TOKENRINGRMONMIB_SourceRoutingStatsTable_SourceRoutingStatsEntry) GetEntityData() *types.CommonEntityData {
    sourceRoutingStatsEntry.EntityData.YFilter = sourceRoutingStatsEntry.YFilter
    sourceRoutingStatsEntry.EntityData.YangName = "sourceRoutingStatsEntry"
    sourceRoutingStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    sourceRoutingStatsEntry.EntityData.ParentYangName = "sourceRoutingStatsTable"
    sourceRoutingStatsEntry.EntityData.SegmentPath = "sourceRoutingStatsEntry" + types.AddKeyToken(sourceRoutingStatsEntry.SourceRoutingStatsIfIndex, "sourceRoutingStatsIfIndex")
    sourceRoutingStatsEntry.EntityData.AbsolutePath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB/sourceRoutingStatsTable/" + sourceRoutingStatsEntry.EntityData.SegmentPath
    sourceRoutingStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sourceRoutingStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sourceRoutingStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sourceRoutingStatsEntry.EntityData.Children = types.NewOrderedMap()
    sourceRoutingStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsIfIndex", types.YLeaf{"SourceRoutingStatsIfIndex", sourceRoutingStatsEntry.SourceRoutingStatsIfIndex})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsRingNumber", types.YLeaf{"SourceRoutingStatsRingNumber", sourceRoutingStatsEntry.SourceRoutingStatsRingNumber})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsInFrames", types.YLeaf{"SourceRoutingStatsInFrames", sourceRoutingStatsEntry.SourceRoutingStatsInFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsOutFrames", types.YLeaf{"SourceRoutingStatsOutFrames", sourceRoutingStatsEntry.SourceRoutingStatsOutFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsThroughFrames", types.YLeaf{"SourceRoutingStatsThroughFrames", sourceRoutingStatsEntry.SourceRoutingStatsThroughFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsAllRoutesBroadcastFrames", types.YLeaf{"SourceRoutingStatsAllRoutesBroadcastFrames", sourceRoutingStatsEntry.SourceRoutingStatsAllRoutesBroadcastFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsSingleRouteBroadcastFrames", types.YLeaf{"SourceRoutingStatsSingleRouteBroadcastFrames", sourceRoutingStatsEntry.SourceRoutingStatsSingleRouteBroadcastFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsInOctets", types.YLeaf{"SourceRoutingStatsInOctets", sourceRoutingStatsEntry.SourceRoutingStatsInOctets})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsOutOctets", types.YLeaf{"SourceRoutingStatsOutOctets", sourceRoutingStatsEntry.SourceRoutingStatsOutOctets})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsThroughOctets", types.YLeaf{"SourceRoutingStatsThroughOctets", sourceRoutingStatsEntry.SourceRoutingStatsThroughOctets})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsAllRoutesBroadcastOctets", types.YLeaf{"SourceRoutingStatsAllRoutesBroadcastOctets", sourceRoutingStatsEntry.SourceRoutingStatsAllRoutesBroadcastOctets})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsSingleRoutesBroadcastOctets", types.YLeaf{"SourceRoutingStatsSingleRoutesBroadcastOctets", sourceRoutingStatsEntry.SourceRoutingStatsSingleRoutesBroadcastOctets})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsLocalLLCFrames", types.YLeaf{"SourceRoutingStatsLocalLLCFrames", sourceRoutingStatsEntry.SourceRoutingStatsLocalLLCFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats1HopFrames", types.YLeaf{"SourceRoutingStats1HopFrames", sourceRoutingStatsEntry.SourceRoutingStats1HopFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats2HopsFrames", types.YLeaf{"SourceRoutingStats2HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats2HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats3HopsFrames", types.YLeaf{"SourceRoutingStats3HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats3HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats4HopsFrames", types.YLeaf{"SourceRoutingStats4HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats4HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats5HopsFrames", types.YLeaf{"SourceRoutingStats5HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats5HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats6HopsFrames", types.YLeaf{"SourceRoutingStats6HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats6HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats7HopsFrames", types.YLeaf{"SourceRoutingStats7HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats7HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStats8HopsFrames", types.YLeaf{"SourceRoutingStats8HopsFrames", sourceRoutingStatsEntry.SourceRoutingStats8HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsMoreThan8HopsFrames", types.YLeaf{"SourceRoutingStatsMoreThan8HopsFrames", sourceRoutingStatsEntry.SourceRoutingStatsMoreThan8HopsFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsOwner", types.YLeaf{"SourceRoutingStatsOwner", sourceRoutingStatsEntry.SourceRoutingStatsOwner})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsStatus", types.YLeaf{"SourceRoutingStatsStatus", sourceRoutingStatsEntry.SourceRoutingStatsStatus})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsDroppedFrames", types.YLeaf{"SourceRoutingStatsDroppedFrames", sourceRoutingStatsEntry.SourceRoutingStatsDroppedFrames})
    sourceRoutingStatsEntry.EntityData.Leafs.Append("sourceRoutingStatsCreateTime", types.YLeaf{"SourceRoutingStatsCreateTime", sourceRoutingStatsEntry.SourceRoutingStatsCreateTime})

    sourceRoutingStatsEntry.EntityData.YListKeys = []string {"SourceRoutingStatsIfIndex"}

    return &(sourceRoutingStatsEntry.EntityData)
}

