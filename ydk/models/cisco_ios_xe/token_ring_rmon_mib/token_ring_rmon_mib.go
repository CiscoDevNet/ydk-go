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

func (tOKENRINGRMONMIB *TOKENRINGRMONMIB) GetEntityData() *types.CommonEntityData {
    tOKENRINGRMONMIB.EntityData.YFilter = tOKENRINGRMONMIB.YFilter
    tOKENRINGRMONMIB.EntityData.YangName = "TOKEN-RING-RMON-MIB"
    tOKENRINGRMONMIB.EntityData.BundleName = "cisco_ios_xe"
    tOKENRINGRMONMIB.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tOKENRINGRMONMIB.EntityData.SegmentPath = "TOKEN-RING-RMON-MIB:TOKEN-RING-RMON-MIB"
    tOKENRINGRMONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tOKENRINGRMONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tOKENRINGRMONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tOKENRINGRMONMIB.EntityData.Children = make(map[string]types.YChild)
    tOKENRINGRMONMIB.EntityData.Children["tokenRingMLStatsTable"] = types.YChild{"Tokenringmlstatstable", &tOKENRINGRMONMIB.Tokenringmlstatstable}
    tOKENRINGRMONMIB.EntityData.Children["tokenRingPStatsTable"] = types.YChild{"Tokenringpstatstable", &tOKENRINGRMONMIB.Tokenringpstatstable}
    tOKENRINGRMONMIB.EntityData.Children["tokenRingMLHistoryTable"] = types.YChild{"Tokenringmlhistorytable", &tOKENRINGRMONMIB.Tokenringmlhistorytable}
    tOKENRINGRMONMIB.EntityData.Children["tokenRingPHistoryTable"] = types.YChild{"Tokenringphistorytable", &tOKENRINGRMONMIB.Tokenringphistorytable}
    tOKENRINGRMONMIB.EntityData.Children["ringStationControlTable"] = types.YChild{"Ringstationcontroltable", &tOKENRINGRMONMIB.Ringstationcontroltable}
    tOKENRINGRMONMIB.EntityData.Children["ringStationTable"] = types.YChild{"Ringstationtable", &tOKENRINGRMONMIB.Ringstationtable}
    tOKENRINGRMONMIB.EntityData.Children["ringStationOrderTable"] = types.YChild{"Ringstationordertable", &tOKENRINGRMONMIB.Ringstationordertable}
    tOKENRINGRMONMIB.EntityData.Children["ringStationConfigControlTable"] = types.YChild{"Ringstationconfigcontroltable", &tOKENRINGRMONMIB.Ringstationconfigcontroltable}
    tOKENRINGRMONMIB.EntityData.Children["ringStationConfigTable"] = types.YChild{"Ringstationconfigtable", &tOKENRINGRMONMIB.Ringstationconfigtable}
    tOKENRINGRMONMIB.EntityData.Children["sourceRoutingStatsTable"] = types.YChild{"Sourceroutingstatstable", &tOKENRINGRMONMIB.Sourceroutingstatstable}
    tOKENRINGRMONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tOKENRINGRMONMIB.EntityData)
}

// TOKENRINGRMONMIB_Tokenringmlstatstable
// A list of Mac-Layer Token Ring statistics
// 
// 
// 
// 
// 
// entries.
type TOKENRINGRMONMIB_Tokenringmlstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of Mac-Layer statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry.
    Tokenringmlstatsentry []TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry
}

func (tokenringmlstatstable *TOKENRINGRMONMIB_Tokenringmlstatstable) GetEntityData() *types.CommonEntityData {
    tokenringmlstatstable.EntityData.YFilter = tokenringmlstatstable.YFilter
    tokenringmlstatstable.EntityData.YangName = "tokenRingMLStatsTable"
    tokenringmlstatstable.EntityData.BundleName = "cisco_ios_xe"
    tokenringmlstatstable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenringmlstatstable.EntityData.SegmentPath = "tokenRingMLStatsTable"
    tokenringmlstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringmlstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringmlstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringmlstatstable.EntityData.Children = make(map[string]types.YChild)
    tokenringmlstatstable.EntityData.Children["tokenRingMLStatsEntry"] = types.YChild{"Tokenringmlstatsentry", nil}
    for i := range tokenringmlstatstable.Tokenringmlstatsentry {
        tokenringmlstatstable.EntityData.Children[types.GetSegmentPath(&tokenringmlstatstable.Tokenringmlstatsentry[i])] = types.YChild{"Tokenringmlstatsentry", &tokenringmlstatstable.Tokenringmlstatsentry[i]}
    }
    tokenringmlstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tokenringmlstatstable.EntityData)
}

// TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry
// A collection of Mac-Layer statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (tokenringmlstatsentry *TOKENRINGRMONMIB_Tokenringmlstatstable_Tokenringmlstatsentry) GetEntityData() *types.CommonEntityData {
    tokenringmlstatsentry.EntityData.YFilter = tokenringmlstatsentry.YFilter
    tokenringmlstatsentry.EntityData.YangName = "tokenRingMLStatsEntry"
    tokenringmlstatsentry.EntityData.BundleName = "cisco_ios_xe"
    tokenringmlstatsentry.EntityData.ParentYangName = "tokenRingMLStatsTable"
    tokenringmlstatsentry.EntityData.SegmentPath = "tokenRingMLStatsEntry" + "[tokenRingMLStatsIndex='" + fmt.Sprintf("%v", tokenringmlstatsentry.Tokenringmlstatsindex) + "']"
    tokenringmlstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringmlstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringmlstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringmlstatsentry.EntityData.Children = make(map[string]types.YChild)
    tokenringmlstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsIndex"] = types.YLeaf{"Tokenringmlstatsindex", tokenringmlstatsentry.Tokenringmlstatsindex}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsDataSource"] = types.YLeaf{"Tokenringmlstatsdatasource", tokenringmlstatsentry.Tokenringmlstatsdatasource}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsDropEvents"] = types.YLeaf{"Tokenringmlstatsdropevents", tokenringmlstatsentry.Tokenringmlstatsdropevents}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsMacOctets"] = types.YLeaf{"Tokenringmlstatsmacoctets", tokenringmlstatsentry.Tokenringmlstatsmacoctets}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsMacPkts"] = types.YLeaf{"Tokenringmlstatsmacpkts", tokenringmlstatsentry.Tokenringmlstatsmacpkts}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsRingPurgeEvents"] = types.YLeaf{"Tokenringmlstatsringpurgeevents", tokenringmlstatsentry.Tokenringmlstatsringpurgeevents}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsRingPurgePkts"] = types.YLeaf{"Tokenringmlstatsringpurgepkts", tokenringmlstatsentry.Tokenringmlstatsringpurgepkts}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsBeaconEvents"] = types.YLeaf{"Tokenringmlstatsbeaconevents", tokenringmlstatsentry.Tokenringmlstatsbeaconevents}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsBeaconTime"] = types.YLeaf{"Tokenringmlstatsbeacontime", tokenringmlstatsentry.Tokenringmlstatsbeacontime}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsBeaconPkts"] = types.YLeaf{"Tokenringmlstatsbeaconpkts", tokenringmlstatsentry.Tokenringmlstatsbeaconpkts}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsClaimTokenEvents"] = types.YLeaf{"Tokenringmlstatsclaimtokenevents", tokenringmlstatsentry.Tokenringmlstatsclaimtokenevents}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsClaimTokenPkts"] = types.YLeaf{"Tokenringmlstatsclaimtokenpkts", tokenringmlstatsentry.Tokenringmlstatsclaimtokenpkts}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsNAUNChanges"] = types.YLeaf{"Tokenringmlstatsnaunchanges", tokenringmlstatsentry.Tokenringmlstatsnaunchanges}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsLineErrors"] = types.YLeaf{"Tokenringmlstatslineerrors", tokenringmlstatsentry.Tokenringmlstatslineerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsInternalErrors"] = types.YLeaf{"Tokenringmlstatsinternalerrors", tokenringmlstatsentry.Tokenringmlstatsinternalerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsBurstErrors"] = types.YLeaf{"Tokenringmlstatsbursterrors", tokenringmlstatsentry.Tokenringmlstatsbursterrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsACErrors"] = types.YLeaf{"Tokenringmlstatsacerrors", tokenringmlstatsentry.Tokenringmlstatsacerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsAbortErrors"] = types.YLeaf{"Tokenringmlstatsaborterrors", tokenringmlstatsentry.Tokenringmlstatsaborterrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsLostFrameErrors"] = types.YLeaf{"Tokenringmlstatslostframeerrors", tokenringmlstatsentry.Tokenringmlstatslostframeerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsCongestionErrors"] = types.YLeaf{"Tokenringmlstatscongestionerrors", tokenringmlstatsentry.Tokenringmlstatscongestionerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsFrameCopiedErrors"] = types.YLeaf{"Tokenringmlstatsframecopiederrors", tokenringmlstatsentry.Tokenringmlstatsframecopiederrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsFrequencyErrors"] = types.YLeaf{"Tokenringmlstatsfrequencyerrors", tokenringmlstatsentry.Tokenringmlstatsfrequencyerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsTokenErrors"] = types.YLeaf{"Tokenringmlstatstokenerrors", tokenringmlstatsentry.Tokenringmlstatstokenerrors}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsSoftErrorReports"] = types.YLeaf{"Tokenringmlstatssofterrorreports", tokenringmlstatsentry.Tokenringmlstatssofterrorreports}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsRingPollEvents"] = types.YLeaf{"Tokenringmlstatsringpollevents", tokenringmlstatsentry.Tokenringmlstatsringpollevents}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsOwner"] = types.YLeaf{"Tokenringmlstatsowner", tokenringmlstatsentry.Tokenringmlstatsowner}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsStatus"] = types.YLeaf{"Tokenringmlstatsstatus", tokenringmlstatsentry.Tokenringmlstatsstatus}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsDroppedFrames"] = types.YLeaf{"Tokenringmlstatsdroppedframes", tokenringmlstatsentry.Tokenringmlstatsdroppedframes}
    tokenringmlstatsentry.EntityData.Leafs["tokenRingMLStatsCreateTime"] = types.YLeaf{"Tokenringmlstatscreatetime", tokenringmlstatsentry.Tokenringmlstatscreatetime}
    return &(tokenringmlstatsentry.EntityData)
}

// TOKENRINGRMONMIB_Tokenringpstatstable
// A list of promiscuous Token Ring statistics
// entries.
type TOKENRINGRMONMIB_Tokenringpstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of promiscuous statistics kept for non-MAC packets on a
    // particular Token Ring interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry.
    Tokenringpstatsentry []TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry
}

func (tokenringpstatstable *TOKENRINGRMONMIB_Tokenringpstatstable) GetEntityData() *types.CommonEntityData {
    tokenringpstatstable.EntityData.YFilter = tokenringpstatstable.YFilter
    tokenringpstatstable.EntityData.YangName = "tokenRingPStatsTable"
    tokenringpstatstable.EntityData.BundleName = "cisco_ios_xe"
    tokenringpstatstable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenringpstatstable.EntityData.SegmentPath = "tokenRingPStatsTable"
    tokenringpstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringpstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringpstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringpstatstable.EntityData.Children = make(map[string]types.YChild)
    tokenringpstatstable.EntityData.Children["tokenRingPStatsEntry"] = types.YChild{"Tokenringpstatsentry", nil}
    for i := range tokenringpstatstable.Tokenringpstatsentry {
        tokenringpstatstable.EntityData.Children[types.GetSegmentPath(&tokenringpstatstable.Tokenringpstatsentry[i])] = types.YChild{"Tokenringpstatsentry", &tokenringpstatstable.Tokenringpstatsentry[i]}
    }
    tokenringpstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tokenringpstatstable.EntityData)
}

// TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry
// A collection of promiscuous statistics kept for
// non-MAC packets on a particular Token Ring
// interface.
type TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (tokenringpstatsentry *TOKENRINGRMONMIB_Tokenringpstatstable_Tokenringpstatsentry) GetEntityData() *types.CommonEntityData {
    tokenringpstatsentry.EntityData.YFilter = tokenringpstatsentry.YFilter
    tokenringpstatsentry.EntityData.YangName = "tokenRingPStatsEntry"
    tokenringpstatsentry.EntityData.BundleName = "cisco_ios_xe"
    tokenringpstatsentry.EntityData.ParentYangName = "tokenRingPStatsTable"
    tokenringpstatsentry.EntityData.SegmentPath = "tokenRingPStatsEntry" + "[tokenRingPStatsIndex='" + fmt.Sprintf("%v", tokenringpstatsentry.Tokenringpstatsindex) + "']"
    tokenringpstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringpstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringpstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringpstatsentry.EntityData.Children = make(map[string]types.YChild)
    tokenringpstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsIndex"] = types.YLeaf{"Tokenringpstatsindex", tokenringpstatsentry.Tokenringpstatsindex}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataSource"] = types.YLeaf{"Tokenringpstatsdatasource", tokenringpstatsentry.Tokenringpstatsdatasource}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDropEvents"] = types.YLeaf{"Tokenringpstatsdropevents", tokenringpstatsentry.Tokenringpstatsdropevents}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataOctets"] = types.YLeaf{"Tokenringpstatsdataoctets", tokenringpstatsentry.Tokenringpstatsdataoctets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts"] = types.YLeaf{"Tokenringpstatsdatapkts", tokenringpstatsentry.Tokenringpstatsdatapkts}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataBroadcastPkts"] = types.YLeaf{"Tokenringpstatsdatabroadcastpkts", tokenringpstatsentry.Tokenringpstatsdatabroadcastpkts}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataMulticastPkts"] = types.YLeaf{"Tokenringpstatsdatamulticastpkts", tokenringpstatsentry.Tokenringpstatsdatamulticastpkts}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts18to63Octets"] = types.YLeaf{"Tokenringpstatsdatapkts18To63Octets", tokenringpstatsentry.Tokenringpstatsdatapkts18To63Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts64to127Octets"] = types.YLeaf{"Tokenringpstatsdatapkts64To127Octets", tokenringpstatsentry.Tokenringpstatsdatapkts64To127Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts128to255Octets"] = types.YLeaf{"Tokenringpstatsdatapkts128To255Octets", tokenringpstatsentry.Tokenringpstatsdatapkts128To255Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts256to511Octets"] = types.YLeaf{"Tokenringpstatsdatapkts256To511Octets", tokenringpstatsentry.Tokenringpstatsdatapkts256To511Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts512to1023Octets"] = types.YLeaf{"Tokenringpstatsdatapkts512To1023Octets", tokenringpstatsentry.Tokenringpstatsdatapkts512To1023Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts1024to2047Octets"] = types.YLeaf{"Tokenringpstatsdatapkts1024To2047Octets", tokenringpstatsentry.Tokenringpstatsdatapkts1024To2047Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts2048to4095Octets"] = types.YLeaf{"Tokenringpstatsdatapkts2048To4095Octets", tokenringpstatsentry.Tokenringpstatsdatapkts2048To4095Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts4096to8191Octets"] = types.YLeaf{"Tokenringpstatsdatapkts4096To8191Octets", tokenringpstatsentry.Tokenringpstatsdatapkts4096To8191Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPkts8192to18000Octets"] = types.YLeaf{"Tokenringpstatsdatapkts8192To18000Octets", tokenringpstatsentry.Tokenringpstatsdatapkts8192To18000Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDataPktsGreaterThan18000Octets"] = types.YLeaf{"Tokenringpstatsdatapktsgreaterthan18000Octets", tokenringpstatsentry.Tokenringpstatsdatapktsgreaterthan18000Octets}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsOwner"] = types.YLeaf{"Tokenringpstatsowner", tokenringpstatsentry.Tokenringpstatsowner}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsStatus"] = types.YLeaf{"Tokenringpstatsstatus", tokenringpstatsentry.Tokenringpstatsstatus}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsDroppedFrames"] = types.YLeaf{"Tokenringpstatsdroppedframes", tokenringpstatsentry.Tokenringpstatsdroppedframes}
    tokenringpstatsentry.EntityData.Leafs["tokenRingPStatsCreateTime"] = types.YLeaf{"Tokenringpstatscreatetime", tokenringpstatsentry.Tokenringpstatscreatetime}
    return &(tokenringpstatsentry.EntityData)
}

// TOKENRINGRMONMIB_Tokenringmlhistorytable
// A list of Mac-Layer Token Ring statistics
// 
// 
// 
// 
// 
// entries.
type TOKENRINGRMONMIB_Tokenringmlhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of Mac-Layer statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry.
    Tokenringmlhistoryentry []TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry
}

func (tokenringmlhistorytable *TOKENRINGRMONMIB_Tokenringmlhistorytable) GetEntityData() *types.CommonEntityData {
    tokenringmlhistorytable.EntityData.YFilter = tokenringmlhistorytable.YFilter
    tokenringmlhistorytable.EntityData.YangName = "tokenRingMLHistoryTable"
    tokenringmlhistorytable.EntityData.BundleName = "cisco_ios_xe"
    tokenringmlhistorytable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenringmlhistorytable.EntityData.SegmentPath = "tokenRingMLHistoryTable"
    tokenringmlhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringmlhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringmlhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringmlhistorytable.EntityData.Children = make(map[string]types.YChild)
    tokenringmlhistorytable.EntityData.Children["tokenRingMLHistoryEntry"] = types.YChild{"Tokenringmlhistoryentry", nil}
    for i := range tokenringmlhistorytable.Tokenringmlhistoryentry {
        tokenringmlhistorytable.EntityData.Children[types.GetSegmentPath(&tokenringmlhistorytable.Tokenringmlhistoryentry[i])] = types.YChild{"Tokenringmlhistoryentry", &tokenringmlhistorytable.Tokenringmlhistoryentry[i]}
    }
    tokenringmlhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tokenringmlhistorytable.EntityData)
}

// TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry
// A collection of Mac-Layer statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry struct {
    EntityData types.CommonEntityData
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

func (tokenringmlhistoryentry *TOKENRINGRMONMIB_Tokenringmlhistorytable_Tokenringmlhistoryentry) GetEntityData() *types.CommonEntityData {
    tokenringmlhistoryentry.EntityData.YFilter = tokenringmlhistoryentry.YFilter
    tokenringmlhistoryentry.EntityData.YangName = "tokenRingMLHistoryEntry"
    tokenringmlhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    tokenringmlhistoryentry.EntityData.ParentYangName = "tokenRingMLHistoryTable"
    tokenringmlhistoryentry.EntityData.SegmentPath = "tokenRingMLHistoryEntry" + "[tokenRingMLHistoryIndex='" + fmt.Sprintf("%v", tokenringmlhistoryentry.Tokenringmlhistoryindex) + "']" + "[tokenRingMLHistorySampleIndex='" + fmt.Sprintf("%v", tokenringmlhistoryentry.Tokenringmlhistorysampleindex) + "']"
    tokenringmlhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringmlhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringmlhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringmlhistoryentry.EntityData.Children = make(map[string]types.YChild)
    tokenringmlhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryIndex"] = types.YLeaf{"Tokenringmlhistoryindex", tokenringmlhistoryentry.Tokenringmlhistoryindex}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistorySampleIndex"] = types.YLeaf{"Tokenringmlhistorysampleindex", tokenringmlhistoryentry.Tokenringmlhistorysampleindex}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryIntervalStart"] = types.YLeaf{"Tokenringmlhistoryintervalstart", tokenringmlhistoryentry.Tokenringmlhistoryintervalstart}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryDropEvents"] = types.YLeaf{"Tokenringmlhistorydropevents", tokenringmlhistoryentry.Tokenringmlhistorydropevents}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryMacOctets"] = types.YLeaf{"Tokenringmlhistorymacoctets", tokenringmlhistoryentry.Tokenringmlhistorymacoctets}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryMacPkts"] = types.YLeaf{"Tokenringmlhistorymacpkts", tokenringmlhistoryentry.Tokenringmlhistorymacpkts}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryRingPurgeEvents"] = types.YLeaf{"Tokenringmlhistoryringpurgeevents", tokenringmlhistoryentry.Tokenringmlhistoryringpurgeevents}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryRingPurgePkts"] = types.YLeaf{"Tokenringmlhistoryringpurgepkts", tokenringmlhistoryentry.Tokenringmlhistoryringpurgepkts}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryBeaconEvents"] = types.YLeaf{"Tokenringmlhistorybeaconevents", tokenringmlhistoryentry.Tokenringmlhistorybeaconevents}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryBeaconTime"] = types.YLeaf{"Tokenringmlhistorybeacontime", tokenringmlhistoryentry.Tokenringmlhistorybeacontime}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryBeaconPkts"] = types.YLeaf{"Tokenringmlhistorybeaconpkts", tokenringmlhistoryentry.Tokenringmlhistorybeaconpkts}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryClaimTokenEvents"] = types.YLeaf{"Tokenringmlhistoryclaimtokenevents", tokenringmlhistoryentry.Tokenringmlhistoryclaimtokenevents}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryClaimTokenPkts"] = types.YLeaf{"Tokenringmlhistoryclaimtokenpkts", tokenringmlhistoryentry.Tokenringmlhistoryclaimtokenpkts}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryNAUNChanges"] = types.YLeaf{"Tokenringmlhistorynaunchanges", tokenringmlhistoryentry.Tokenringmlhistorynaunchanges}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryLineErrors"] = types.YLeaf{"Tokenringmlhistorylineerrors", tokenringmlhistoryentry.Tokenringmlhistorylineerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryInternalErrors"] = types.YLeaf{"Tokenringmlhistoryinternalerrors", tokenringmlhistoryentry.Tokenringmlhistoryinternalerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryBurstErrors"] = types.YLeaf{"Tokenringmlhistorybursterrors", tokenringmlhistoryentry.Tokenringmlhistorybursterrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryACErrors"] = types.YLeaf{"Tokenringmlhistoryacerrors", tokenringmlhistoryentry.Tokenringmlhistoryacerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryAbortErrors"] = types.YLeaf{"Tokenringmlhistoryaborterrors", tokenringmlhistoryentry.Tokenringmlhistoryaborterrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryLostFrameErrors"] = types.YLeaf{"Tokenringmlhistorylostframeerrors", tokenringmlhistoryentry.Tokenringmlhistorylostframeerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryCongestionErrors"] = types.YLeaf{"Tokenringmlhistorycongestionerrors", tokenringmlhistoryentry.Tokenringmlhistorycongestionerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryFrameCopiedErrors"] = types.YLeaf{"Tokenringmlhistoryframecopiederrors", tokenringmlhistoryentry.Tokenringmlhistoryframecopiederrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryFrequencyErrors"] = types.YLeaf{"Tokenringmlhistoryfrequencyerrors", tokenringmlhistoryentry.Tokenringmlhistoryfrequencyerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryTokenErrors"] = types.YLeaf{"Tokenringmlhistorytokenerrors", tokenringmlhistoryentry.Tokenringmlhistorytokenerrors}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistorySoftErrorReports"] = types.YLeaf{"Tokenringmlhistorysofterrorreports", tokenringmlhistoryentry.Tokenringmlhistorysofterrorreports}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryRingPollEvents"] = types.YLeaf{"Tokenringmlhistoryringpollevents", tokenringmlhistoryentry.Tokenringmlhistoryringpollevents}
    tokenringmlhistoryentry.EntityData.Leafs["tokenRingMLHistoryActiveStations"] = types.YLeaf{"Tokenringmlhistoryactivestations", tokenringmlhistoryentry.Tokenringmlhistoryactivestations}
    return &(tokenringmlhistoryentry.EntityData)
}

// TOKENRINGRMONMIB_Tokenringphistorytable
// A list of promiscuous Token Ring statistics
// entries.
type TOKENRINGRMONMIB_Tokenringphistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of promiscuous statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry.
    Tokenringphistoryentry []TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry
}

func (tokenringphistorytable *TOKENRINGRMONMIB_Tokenringphistorytable) GetEntityData() *types.CommonEntityData {
    tokenringphistorytable.EntityData.YFilter = tokenringphistorytable.YFilter
    tokenringphistorytable.EntityData.YangName = "tokenRingPHistoryTable"
    tokenringphistorytable.EntityData.BundleName = "cisco_ios_xe"
    tokenringphistorytable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    tokenringphistorytable.EntityData.SegmentPath = "tokenRingPHistoryTable"
    tokenringphistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringphistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringphistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringphistorytable.EntityData.Children = make(map[string]types.YChild)
    tokenringphistorytable.EntityData.Children["tokenRingPHistoryEntry"] = types.YChild{"Tokenringphistoryentry", nil}
    for i := range tokenringphistorytable.Tokenringphistoryentry {
        tokenringphistorytable.EntityData.Children[types.GetSegmentPath(&tokenringphistorytable.Tokenringphistoryentry[i])] = types.YChild{"Tokenringphistoryentry", &tokenringphistorytable.Tokenringphistoryentry[i]}
    }
    tokenringphistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tokenringphistorytable.EntityData)
}

// TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry
// A collection of promiscuous statistics kept for a
// particular Token Ring interface.
type TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry struct {
    EntityData types.CommonEntityData
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

func (tokenringphistoryentry *TOKENRINGRMONMIB_Tokenringphistorytable_Tokenringphistoryentry) GetEntityData() *types.CommonEntityData {
    tokenringphistoryentry.EntityData.YFilter = tokenringphistoryentry.YFilter
    tokenringphistoryentry.EntityData.YangName = "tokenRingPHistoryEntry"
    tokenringphistoryentry.EntityData.BundleName = "cisco_ios_xe"
    tokenringphistoryentry.EntityData.ParentYangName = "tokenRingPHistoryTable"
    tokenringphistoryentry.EntityData.SegmentPath = "tokenRingPHistoryEntry" + "[tokenRingPHistoryIndex='" + fmt.Sprintf("%v", tokenringphistoryentry.Tokenringphistoryindex) + "']" + "[tokenRingPHistorySampleIndex='" + fmt.Sprintf("%v", tokenringphistoryentry.Tokenringphistorysampleindex) + "']"
    tokenringphistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tokenringphistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tokenringphistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tokenringphistoryentry.EntityData.Children = make(map[string]types.YChild)
    tokenringphistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryIndex"] = types.YLeaf{"Tokenringphistoryindex", tokenringphistoryentry.Tokenringphistoryindex}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistorySampleIndex"] = types.YLeaf{"Tokenringphistorysampleindex", tokenringphistoryentry.Tokenringphistorysampleindex}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryIntervalStart"] = types.YLeaf{"Tokenringphistoryintervalstart", tokenringphistoryentry.Tokenringphistoryintervalstart}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDropEvents"] = types.YLeaf{"Tokenringphistorydropevents", tokenringphistoryentry.Tokenringphistorydropevents}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataOctets"] = types.YLeaf{"Tokenringphistorydataoctets", tokenringphistoryentry.Tokenringphistorydataoctets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts"] = types.YLeaf{"Tokenringphistorydatapkts", tokenringphistoryentry.Tokenringphistorydatapkts}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataBroadcastPkts"] = types.YLeaf{"Tokenringphistorydatabroadcastpkts", tokenringphistoryentry.Tokenringphistorydatabroadcastpkts}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataMulticastPkts"] = types.YLeaf{"Tokenringphistorydatamulticastpkts", tokenringphistoryentry.Tokenringphistorydatamulticastpkts}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts18to63Octets"] = types.YLeaf{"Tokenringphistorydatapkts18To63Octets", tokenringphistoryentry.Tokenringphistorydatapkts18To63Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts64to127Octets"] = types.YLeaf{"Tokenringphistorydatapkts64To127Octets", tokenringphistoryentry.Tokenringphistorydatapkts64To127Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts128to255Octets"] = types.YLeaf{"Tokenringphistorydatapkts128To255Octets", tokenringphistoryentry.Tokenringphistorydatapkts128To255Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts256to511Octets"] = types.YLeaf{"Tokenringphistorydatapkts256To511Octets", tokenringphistoryentry.Tokenringphistorydatapkts256To511Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts512to1023Octets"] = types.YLeaf{"Tokenringphistorydatapkts512To1023Octets", tokenringphistoryentry.Tokenringphistorydatapkts512To1023Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts1024to2047Octets"] = types.YLeaf{"Tokenringphistorydatapkts1024To2047Octets", tokenringphistoryentry.Tokenringphistorydatapkts1024To2047Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts2048to4095Octets"] = types.YLeaf{"Tokenringphistorydatapkts2048To4095Octets", tokenringphistoryentry.Tokenringphistorydatapkts2048To4095Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts4096to8191Octets"] = types.YLeaf{"Tokenringphistorydatapkts4096To8191Octets", tokenringphistoryentry.Tokenringphistorydatapkts4096To8191Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPkts8192to18000Octets"] = types.YLeaf{"Tokenringphistorydatapkts8192To18000Octets", tokenringphistoryentry.Tokenringphistorydatapkts8192To18000Octets}
    tokenringphistoryentry.EntityData.Leafs["tokenRingPHistoryDataPktsGreaterThan18000Octets"] = types.YLeaf{"Tokenringphistorydatapktsgreaterthan18000Octets", tokenringphistoryentry.Tokenringphistorydatapktsgreaterthan18000Octets}
    return &(tokenringphistoryentry.EntityData)
}

// TOKENRINGRMONMIB_Ringstationcontroltable
// A list of ringStation table control entries.
type TOKENRINGRMONMIB_Ringstationcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up the discovery of stations on a particular
    // interface and the collection of statistics about these stations. The type
    // is slice of
    // TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry.
    Ringstationcontrolentry []TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry
}

func (ringstationcontroltable *TOKENRINGRMONMIB_Ringstationcontroltable) GetEntityData() *types.CommonEntityData {
    ringstationcontroltable.EntityData.YFilter = ringstationcontroltable.YFilter
    ringstationcontroltable.EntityData.YangName = "ringStationControlTable"
    ringstationcontroltable.EntityData.BundleName = "cisco_ios_xe"
    ringstationcontroltable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringstationcontroltable.EntityData.SegmentPath = "ringStationControlTable"
    ringstationcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationcontroltable.EntityData.Children = make(map[string]types.YChild)
    ringstationcontroltable.EntityData.Children["ringStationControlEntry"] = types.YChild{"Ringstationcontrolentry", nil}
    for i := range ringstationcontroltable.Ringstationcontrolentry {
        ringstationcontroltable.EntityData.Children[types.GetSegmentPath(&ringstationcontroltable.Ringstationcontrolentry[i])] = types.YChild{"Ringstationcontrolentry", &ringstationcontroltable.Ringstationcontrolentry[i]}
    }
    ringstationcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ringstationcontroltable.EntityData)
}

// TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry
// A list of parameters that set up the discovery of
// stations on a particular interface and the
// collection of statistics about these stations.
type TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry struct {
    EntityData types.CommonEntityData
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

func (ringstationcontrolentry *TOKENRINGRMONMIB_Ringstationcontroltable_Ringstationcontrolentry) GetEntityData() *types.CommonEntityData {
    ringstationcontrolentry.EntityData.YFilter = ringstationcontrolentry.YFilter
    ringstationcontrolentry.EntityData.YangName = "ringStationControlEntry"
    ringstationcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    ringstationcontrolentry.EntityData.ParentYangName = "ringStationControlTable"
    ringstationcontrolentry.EntityData.SegmentPath = "ringStationControlEntry" + "[ringStationControlIfIndex='" + fmt.Sprintf("%v", ringstationcontrolentry.Ringstationcontrolifindex) + "']"
    ringstationcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationcontrolentry.EntityData.Children = make(map[string]types.YChild)
    ringstationcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ringstationcontrolentry.EntityData.Leafs["ringStationControlIfIndex"] = types.YLeaf{"Ringstationcontrolifindex", ringstationcontrolentry.Ringstationcontrolifindex}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlTableSize"] = types.YLeaf{"Ringstationcontroltablesize", ringstationcontrolentry.Ringstationcontroltablesize}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlActiveStations"] = types.YLeaf{"Ringstationcontrolactivestations", ringstationcontrolentry.Ringstationcontrolactivestations}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlRingState"] = types.YLeaf{"Ringstationcontrolringstate", ringstationcontrolentry.Ringstationcontrolringstate}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlBeaconSender"] = types.YLeaf{"Ringstationcontrolbeaconsender", ringstationcontrolentry.Ringstationcontrolbeaconsender}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlBeaconNAUN"] = types.YLeaf{"Ringstationcontrolbeaconnaun", ringstationcontrolentry.Ringstationcontrolbeaconnaun}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlActiveMonitor"] = types.YLeaf{"Ringstationcontrolactivemonitor", ringstationcontrolentry.Ringstationcontrolactivemonitor}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlOrderChanges"] = types.YLeaf{"Ringstationcontrolorderchanges", ringstationcontrolentry.Ringstationcontrolorderchanges}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlOwner"] = types.YLeaf{"Ringstationcontrolowner", ringstationcontrolentry.Ringstationcontrolowner}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlStatus"] = types.YLeaf{"Ringstationcontrolstatus", ringstationcontrolentry.Ringstationcontrolstatus}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlDroppedFrames"] = types.YLeaf{"Ringstationcontroldroppedframes", ringstationcontrolentry.Ringstationcontroldroppedframes}
    ringstationcontrolentry.EntityData.Leafs["ringStationControlCreateTime"] = types.YLeaf{"Ringstationcontrolcreatetime", ringstationcontrolentry.Ringstationcontrolcreatetime}
    return &(ringstationcontrolentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular station that has been
    // discovered on a ring monitored by this device. The type is slice of
    // TOKENRINGRMONMIB_Ringstationtable_Ringstationentry.
    Ringstationentry []TOKENRINGRMONMIB_Ringstationtable_Ringstationentry
}

func (ringstationtable *TOKENRINGRMONMIB_Ringstationtable) GetEntityData() *types.CommonEntityData {
    ringstationtable.EntityData.YFilter = ringstationtable.YFilter
    ringstationtable.EntityData.YangName = "ringStationTable"
    ringstationtable.EntityData.BundleName = "cisco_ios_xe"
    ringstationtable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringstationtable.EntityData.SegmentPath = "ringStationTable"
    ringstationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationtable.EntityData.Children = make(map[string]types.YChild)
    ringstationtable.EntityData.Children["ringStationEntry"] = types.YChild{"Ringstationentry", nil}
    for i := range ringstationtable.Ringstationentry {
        ringstationtable.EntityData.Children[types.GetSegmentPath(&ringstationtable.Ringstationentry[i])] = types.YChild{"Ringstationentry", &ringstationtable.Ringstationentry[i]}
    }
    ringstationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ringstationtable.EntityData)
}

// TOKENRINGRMONMIB_Ringstationtable_Ringstationentry
// A collection of statistics for a particular
// station that has been discovered on a ring
// monitored by this device.
type TOKENRINGRMONMIB_Ringstationtable_Ringstationentry struct {
    EntityData types.CommonEntityData
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

func (ringstationentry *TOKENRINGRMONMIB_Ringstationtable_Ringstationentry) GetEntityData() *types.CommonEntityData {
    ringstationentry.EntityData.YFilter = ringstationentry.YFilter
    ringstationentry.EntityData.YangName = "ringStationEntry"
    ringstationentry.EntityData.BundleName = "cisco_ios_xe"
    ringstationentry.EntityData.ParentYangName = "ringStationTable"
    ringstationentry.EntityData.SegmentPath = "ringStationEntry" + "[ringStationIfIndex='" + fmt.Sprintf("%v", ringstationentry.Ringstationifindex) + "']" + "[ringStationMacAddress='" + fmt.Sprintf("%v", ringstationentry.Ringstationmacaddress) + "']"
    ringstationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationentry.EntityData.Children = make(map[string]types.YChild)
    ringstationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ringstationentry.EntityData.Leafs["ringStationIfIndex"] = types.YLeaf{"Ringstationifindex", ringstationentry.Ringstationifindex}
    ringstationentry.EntityData.Leafs["ringStationMacAddress"] = types.YLeaf{"Ringstationmacaddress", ringstationentry.Ringstationmacaddress}
    ringstationentry.EntityData.Leafs["ringStationLastNAUN"] = types.YLeaf{"Ringstationlastnaun", ringstationentry.Ringstationlastnaun}
    ringstationentry.EntityData.Leafs["ringStationStationStatus"] = types.YLeaf{"Ringstationstationstatus", ringstationentry.Ringstationstationstatus}
    ringstationentry.EntityData.Leafs["ringStationLastEnterTime"] = types.YLeaf{"Ringstationlastentertime", ringstationentry.Ringstationlastentertime}
    ringstationentry.EntityData.Leafs["ringStationLastExitTime"] = types.YLeaf{"Ringstationlastexittime", ringstationentry.Ringstationlastexittime}
    ringstationentry.EntityData.Leafs["ringStationDuplicateAddresses"] = types.YLeaf{"Ringstationduplicateaddresses", ringstationentry.Ringstationduplicateaddresses}
    ringstationentry.EntityData.Leafs["ringStationInLineErrors"] = types.YLeaf{"Ringstationinlineerrors", ringstationentry.Ringstationinlineerrors}
    ringstationentry.EntityData.Leafs["ringStationOutLineErrors"] = types.YLeaf{"Ringstationoutlineerrors", ringstationentry.Ringstationoutlineerrors}
    ringstationentry.EntityData.Leafs["ringStationInternalErrors"] = types.YLeaf{"Ringstationinternalerrors", ringstationentry.Ringstationinternalerrors}
    ringstationentry.EntityData.Leafs["ringStationInBurstErrors"] = types.YLeaf{"Ringstationinbursterrors", ringstationentry.Ringstationinbursterrors}
    ringstationentry.EntityData.Leafs["ringStationOutBurstErrors"] = types.YLeaf{"Ringstationoutbursterrors", ringstationentry.Ringstationoutbursterrors}
    ringstationentry.EntityData.Leafs["ringStationACErrors"] = types.YLeaf{"Ringstationacerrors", ringstationentry.Ringstationacerrors}
    ringstationentry.EntityData.Leafs["ringStationAbortErrors"] = types.YLeaf{"Ringstationaborterrors", ringstationentry.Ringstationaborterrors}
    ringstationentry.EntityData.Leafs["ringStationLostFrameErrors"] = types.YLeaf{"Ringstationlostframeerrors", ringstationentry.Ringstationlostframeerrors}
    ringstationentry.EntityData.Leafs["ringStationCongestionErrors"] = types.YLeaf{"Ringstationcongestionerrors", ringstationentry.Ringstationcongestionerrors}
    ringstationentry.EntityData.Leafs["ringStationFrameCopiedErrors"] = types.YLeaf{"Ringstationframecopiederrors", ringstationentry.Ringstationframecopiederrors}
    ringstationentry.EntityData.Leafs["ringStationFrequencyErrors"] = types.YLeaf{"Ringstationfrequencyerrors", ringstationentry.Ringstationfrequencyerrors}
    ringstationentry.EntityData.Leafs["ringStationTokenErrors"] = types.YLeaf{"Ringstationtokenerrors", ringstationentry.Ringstationtokenerrors}
    ringstationentry.EntityData.Leafs["ringStationInBeaconErrors"] = types.YLeaf{"Ringstationinbeaconerrors", ringstationentry.Ringstationinbeaconerrors}
    ringstationentry.EntityData.Leafs["ringStationOutBeaconErrors"] = types.YLeaf{"Ringstationoutbeaconerrors", ringstationentry.Ringstationoutbeaconerrors}
    ringstationentry.EntityData.Leafs["ringStationInsertions"] = types.YLeaf{"Ringstationinsertions", ringstationentry.Ringstationinsertions}
    return &(ringstationentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular      station that is active on
    // a ring monitored by this device.  This table will contain information for
    // every interface that has a ringStationControlStatus equal to valid. The
    // type is slice of
    // TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry.
    Ringstationorderentry []TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry
}

func (ringstationordertable *TOKENRINGRMONMIB_Ringstationordertable) GetEntityData() *types.CommonEntityData {
    ringstationordertable.EntityData.YFilter = ringstationordertable.YFilter
    ringstationordertable.EntityData.YangName = "ringStationOrderTable"
    ringstationordertable.EntityData.BundleName = "cisco_ios_xe"
    ringstationordertable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringstationordertable.EntityData.SegmentPath = "ringStationOrderTable"
    ringstationordertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationordertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationordertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationordertable.EntityData.Children = make(map[string]types.YChild)
    ringstationordertable.EntityData.Children["ringStationOrderEntry"] = types.YChild{"Ringstationorderentry", nil}
    for i := range ringstationordertable.Ringstationorderentry {
        ringstationordertable.EntityData.Children[types.GetSegmentPath(&ringstationordertable.Ringstationorderentry[i])] = types.YChild{"Ringstationorderentry", &ringstationordertable.Ringstationorderentry[i]}
    }
    ringstationordertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ringstationordertable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ringstationorderentry *TOKENRINGRMONMIB_Ringstationordertable_Ringstationorderentry) GetEntityData() *types.CommonEntityData {
    ringstationorderentry.EntityData.YFilter = ringstationorderentry.YFilter
    ringstationorderentry.EntityData.YangName = "ringStationOrderEntry"
    ringstationorderentry.EntityData.BundleName = "cisco_ios_xe"
    ringstationorderentry.EntityData.ParentYangName = "ringStationOrderTable"
    ringstationorderentry.EntityData.SegmentPath = "ringStationOrderEntry" + "[ringStationOrderIfIndex='" + fmt.Sprintf("%v", ringstationorderentry.Ringstationorderifindex) + "']" + "[ringStationOrderOrderIndex='" + fmt.Sprintf("%v", ringstationorderentry.Ringstationorderorderindex) + "']"
    ringstationorderentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationorderentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationorderentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationorderentry.EntityData.Children = make(map[string]types.YChild)
    ringstationorderentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ringstationorderentry.EntityData.Leafs["ringStationOrderIfIndex"] = types.YLeaf{"Ringstationorderifindex", ringstationorderentry.Ringstationorderifindex}
    ringstationorderentry.EntityData.Leafs["ringStationOrderOrderIndex"] = types.YLeaf{"Ringstationorderorderindex", ringstationorderentry.Ringstationorderorderindex}
    ringstationorderentry.EntityData.Leafs["ringStationOrderMacAddress"] = types.YLeaf{"Ringstationordermacaddress", ringstationorderentry.Ringstationordermacaddress}
    return &(ringstationorderentry.EntityData)
}

// TOKENRINGRMONMIB_Ringstationconfigcontroltable
// A list of ring station configuration control
// entries.
type TOKENRINGRMONMIB_Ringstationconfigcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry controls active management of stations by the probe.  One entry
    // exists in this table for each active station in the ringStationTable. The
    // type is slice of
    // TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry.
    Ringstationconfigcontrolentry []TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry
}

func (ringstationconfigcontroltable *TOKENRINGRMONMIB_Ringstationconfigcontroltable) GetEntityData() *types.CommonEntityData {
    ringstationconfigcontroltable.EntityData.YFilter = ringstationconfigcontroltable.YFilter
    ringstationconfigcontroltable.EntityData.YangName = "ringStationConfigControlTable"
    ringstationconfigcontroltable.EntityData.BundleName = "cisco_ios_xe"
    ringstationconfigcontroltable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringstationconfigcontroltable.EntityData.SegmentPath = "ringStationConfigControlTable"
    ringstationconfigcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationconfigcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationconfigcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationconfigcontroltable.EntityData.Children = make(map[string]types.YChild)
    ringstationconfigcontroltable.EntityData.Children["ringStationConfigControlEntry"] = types.YChild{"Ringstationconfigcontrolentry", nil}
    for i := range ringstationconfigcontroltable.Ringstationconfigcontrolentry {
        ringstationconfigcontroltable.EntityData.Children[types.GetSegmentPath(&ringstationconfigcontroltable.Ringstationconfigcontrolentry[i])] = types.YChild{"Ringstationconfigcontrolentry", &ringstationconfigcontroltable.Ringstationconfigcontrolentry[i]}
    }
    ringstationconfigcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ringstationconfigcontroltable.EntityData)
}

// TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry
// This entry controls active management of stations
// by the probe.  One entry exists in this table for
// each active station in the ringStationTable.
type TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry struct {
    EntityData types.CommonEntityData
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

func (ringstationconfigcontrolentry *TOKENRINGRMONMIB_Ringstationconfigcontroltable_Ringstationconfigcontrolentry) GetEntityData() *types.CommonEntityData {
    ringstationconfigcontrolentry.EntityData.YFilter = ringstationconfigcontrolentry.YFilter
    ringstationconfigcontrolentry.EntityData.YangName = "ringStationConfigControlEntry"
    ringstationconfigcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    ringstationconfigcontrolentry.EntityData.ParentYangName = "ringStationConfigControlTable"
    ringstationconfigcontrolentry.EntityData.SegmentPath = "ringStationConfigControlEntry" + "[ringStationConfigControlIfIndex='" + fmt.Sprintf("%v", ringstationconfigcontrolentry.Ringstationconfigcontrolifindex) + "']" + "[ringStationConfigControlMacAddress='" + fmt.Sprintf("%v", ringstationconfigcontrolentry.Ringstationconfigcontrolmacaddress) + "']"
    ringstationconfigcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationconfigcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationconfigcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationconfigcontrolentry.EntityData.Children = make(map[string]types.YChild)
    ringstationconfigcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ringstationconfigcontrolentry.EntityData.Leafs["ringStationConfigControlIfIndex"] = types.YLeaf{"Ringstationconfigcontrolifindex", ringstationconfigcontrolentry.Ringstationconfigcontrolifindex}
    ringstationconfigcontrolentry.EntityData.Leafs["ringStationConfigControlMacAddress"] = types.YLeaf{"Ringstationconfigcontrolmacaddress", ringstationconfigcontrolentry.Ringstationconfigcontrolmacaddress}
    ringstationconfigcontrolentry.EntityData.Leafs["ringStationConfigControlRemove"] = types.YLeaf{"Ringstationconfigcontrolremove", ringstationconfigcontrolentry.Ringstationconfigcontrolremove}
    ringstationconfigcontrolentry.EntityData.Leafs["ringStationConfigControlUpdateStats"] = types.YLeaf{"Ringstationconfigcontrolupdatestats", ringstationconfigcontrolentry.Ringstationconfigcontrolupdatestats}
    return &(ringstationconfigcontrolentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular station that has been
    // discovered on a ring monitored by this probe. The type is slice of
    // TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry.
    Ringstationconfigentry []TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry
}

func (ringstationconfigtable *TOKENRINGRMONMIB_Ringstationconfigtable) GetEntityData() *types.CommonEntityData {
    ringstationconfigtable.EntityData.YFilter = ringstationconfigtable.YFilter
    ringstationconfigtable.EntityData.YangName = "ringStationConfigTable"
    ringstationconfigtable.EntityData.BundleName = "cisco_ios_xe"
    ringstationconfigtable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    ringstationconfigtable.EntityData.SegmentPath = "ringStationConfigTable"
    ringstationconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationconfigtable.EntityData.Children = make(map[string]types.YChild)
    ringstationconfigtable.EntityData.Children["ringStationConfigEntry"] = types.YChild{"Ringstationconfigentry", nil}
    for i := range ringstationconfigtable.Ringstationconfigentry {
        ringstationconfigtable.EntityData.Children[types.GetSegmentPath(&ringstationconfigtable.Ringstationconfigentry[i])] = types.YChild{"Ringstationconfigentry", &ringstationconfigtable.Ringstationconfigentry[i]}
    }
    ringstationconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ringstationconfigtable.EntityData)
}

// TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry
// A collection of statistics for a particular
// station that has been discovered on a ring
// monitored by this probe.
type TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry struct {
    EntityData types.CommonEntityData
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

func (ringstationconfigentry *TOKENRINGRMONMIB_Ringstationconfigtable_Ringstationconfigentry) GetEntityData() *types.CommonEntityData {
    ringstationconfigentry.EntityData.YFilter = ringstationconfigentry.YFilter
    ringstationconfigentry.EntityData.YangName = "ringStationConfigEntry"
    ringstationconfigentry.EntityData.BundleName = "cisco_ios_xe"
    ringstationconfigentry.EntityData.ParentYangName = "ringStationConfigTable"
    ringstationconfigentry.EntityData.SegmentPath = "ringStationConfigEntry" + "[ringStationConfigIfIndex='" + fmt.Sprintf("%v", ringstationconfigentry.Ringstationconfigifindex) + "']" + "[ringStationConfigMacAddress='" + fmt.Sprintf("%v", ringstationconfigentry.Ringstationconfigmacaddress) + "']"
    ringstationconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ringstationconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ringstationconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ringstationconfigentry.EntityData.Children = make(map[string]types.YChild)
    ringstationconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ringstationconfigentry.EntityData.Leafs["ringStationConfigIfIndex"] = types.YLeaf{"Ringstationconfigifindex", ringstationconfigentry.Ringstationconfigifindex}
    ringstationconfigentry.EntityData.Leafs["ringStationConfigMacAddress"] = types.YLeaf{"Ringstationconfigmacaddress", ringstationconfigentry.Ringstationconfigmacaddress}
    ringstationconfigentry.EntityData.Leafs["ringStationConfigUpdateTime"] = types.YLeaf{"Ringstationconfigupdatetime", ringstationconfigentry.Ringstationconfigupdatetime}
    ringstationconfigentry.EntityData.Leafs["ringStationConfigLocation"] = types.YLeaf{"Ringstationconfiglocation", ringstationconfigentry.Ringstationconfiglocation}
    ringstationconfigentry.EntityData.Leafs["ringStationConfigMicrocode"] = types.YLeaf{"Ringstationconfigmicrocode", ringstationconfigentry.Ringstationconfigmicrocode}
    ringstationconfigentry.EntityData.Leafs["ringStationConfigGroupAddress"] = types.YLeaf{"Ringstationconfiggroupaddress", ringstationconfigentry.Ringstationconfiggroupaddress}
    ringstationconfigentry.EntityData.Leafs["ringStationConfigFunctionalAddress"] = types.YLeaf{"Ringstationconfigfunctionaladdress", ringstationconfigentry.Ringstationconfigfunctionaladdress}
    return &(ringstationconfigentry.EntityData)
}

// TOKENRINGRMONMIB_Sourceroutingstatstable
// A list of source routing statistics entries.
type TOKENRINGRMONMIB_Sourceroutingstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of source routing statistics kept for a particular Token Ring
    // interface. The type is slice of
    // TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry.
    Sourceroutingstatsentry []TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry
}

func (sourceroutingstatstable *TOKENRINGRMONMIB_Sourceroutingstatstable) GetEntityData() *types.CommonEntityData {
    sourceroutingstatstable.EntityData.YFilter = sourceroutingstatstable.YFilter
    sourceroutingstatstable.EntityData.YangName = "sourceRoutingStatsTable"
    sourceroutingstatstable.EntityData.BundleName = "cisco_ios_xe"
    sourceroutingstatstable.EntityData.ParentYangName = "TOKEN-RING-RMON-MIB"
    sourceroutingstatstable.EntityData.SegmentPath = "sourceRoutingStatsTable"
    sourceroutingstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sourceroutingstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sourceroutingstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sourceroutingstatstable.EntityData.Children = make(map[string]types.YChild)
    sourceroutingstatstable.EntityData.Children["sourceRoutingStatsEntry"] = types.YChild{"Sourceroutingstatsentry", nil}
    for i := range sourceroutingstatstable.Sourceroutingstatsentry {
        sourceroutingstatstable.EntityData.Children[types.GetSegmentPath(&sourceroutingstatstable.Sourceroutingstatsentry[i])] = types.YChild{"Sourceroutingstatsentry", &sourceroutingstatstable.Sourceroutingstatsentry[i]}
    }
    sourceroutingstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sourceroutingstatstable.EntityData)
}

// TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry
// A collection of source routing statistics kept
// for a particular Token Ring interface.
type TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry struct {
    EntityData types.CommonEntityData
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

func (sourceroutingstatsentry *TOKENRINGRMONMIB_Sourceroutingstatstable_Sourceroutingstatsentry) GetEntityData() *types.CommonEntityData {
    sourceroutingstatsentry.EntityData.YFilter = sourceroutingstatsentry.YFilter
    sourceroutingstatsentry.EntityData.YangName = "sourceRoutingStatsEntry"
    sourceroutingstatsentry.EntityData.BundleName = "cisco_ios_xe"
    sourceroutingstatsentry.EntityData.ParentYangName = "sourceRoutingStatsTable"
    sourceroutingstatsentry.EntityData.SegmentPath = "sourceRoutingStatsEntry" + "[sourceRoutingStatsIfIndex='" + fmt.Sprintf("%v", sourceroutingstatsentry.Sourceroutingstatsifindex) + "']"
    sourceroutingstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sourceroutingstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sourceroutingstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sourceroutingstatsentry.EntityData.Children = make(map[string]types.YChild)
    sourceroutingstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsIfIndex"] = types.YLeaf{"Sourceroutingstatsifindex", sourceroutingstatsentry.Sourceroutingstatsifindex}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsRingNumber"] = types.YLeaf{"Sourceroutingstatsringnumber", sourceroutingstatsentry.Sourceroutingstatsringnumber}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsInFrames"] = types.YLeaf{"Sourceroutingstatsinframes", sourceroutingstatsentry.Sourceroutingstatsinframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsOutFrames"] = types.YLeaf{"Sourceroutingstatsoutframes", sourceroutingstatsentry.Sourceroutingstatsoutframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsThroughFrames"] = types.YLeaf{"Sourceroutingstatsthroughframes", sourceroutingstatsentry.Sourceroutingstatsthroughframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsAllRoutesBroadcastFrames"] = types.YLeaf{"Sourceroutingstatsallroutesbroadcastframes", sourceroutingstatsentry.Sourceroutingstatsallroutesbroadcastframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsSingleRouteBroadcastFrames"] = types.YLeaf{"Sourceroutingstatssingleroutebroadcastframes", sourceroutingstatsentry.Sourceroutingstatssingleroutebroadcastframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsInOctets"] = types.YLeaf{"Sourceroutingstatsinoctets", sourceroutingstatsentry.Sourceroutingstatsinoctets}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsOutOctets"] = types.YLeaf{"Sourceroutingstatsoutoctets", sourceroutingstatsentry.Sourceroutingstatsoutoctets}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsThroughOctets"] = types.YLeaf{"Sourceroutingstatsthroughoctets", sourceroutingstatsentry.Sourceroutingstatsthroughoctets}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsAllRoutesBroadcastOctets"] = types.YLeaf{"Sourceroutingstatsallroutesbroadcastoctets", sourceroutingstatsentry.Sourceroutingstatsallroutesbroadcastoctets}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsSingleRoutesBroadcastOctets"] = types.YLeaf{"Sourceroutingstatssingleroutesbroadcastoctets", sourceroutingstatsentry.Sourceroutingstatssingleroutesbroadcastoctets}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsLocalLLCFrames"] = types.YLeaf{"Sourceroutingstatslocalllcframes", sourceroutingstatsentry.Sourceroutingstatslocalllcframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats1HopFrames"] = types.YLeaf{"Sourceroutingstats1Hopframes", sourceroutingstatsentry.Sourceroutingstats1Hopframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats2HopsFrames"] = types.YLeaf{"Sourceroutingstats2Hopsframes", sourceroutingstatsentry.Sourceroutingstats2Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats3HopsFrames"] = types.YLeaf{"Sourceroutingstats3Hopsframes", sourceroutingstatsentry.Sourceroutingstats3Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats4HopsFrames"] = types.YLeaf{"Sourceroutingstats4Hopsframes", sourceroutingstatsentry.Sourceroutingstats4Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats5HopsFrames"] = types.YLeaf{"Sourceroutingstats5Hopsframes", sourceroutingstatsentry.Sourceroutingstats5Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats6HopsFrames"] = types.YLeaf{"Sourceroutingstats6Hopsframes", sourceroutingstatsentry.Sourceroutingstats6Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats7HopsFrames"] = types.YLeaf{"Sourceroutingstats7Hopsframes", sourceroutingstatsentry.Sourceroutingstats7Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStats8HopsFrames"] = types.YLeaf{"Sourceroutingstats8Hopsframes", sourceroutingstatsentry.Sourceroutingstats8Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsMoreThan8HopsFrames"] = types.YLeaf{"Sourceroutingstatsmorethan8Hopsframes", sourceroutingstatsentry.Sourceroutingstatsmorethan8Hopsframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsOwner"] = types.YLeaf{"Sourceroutingstatsowner", sourceroutingstatsentry.Sourceroutingstatsowner}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsStatus"] = types.YLeaf{"Sourceroutingstatsstatus", sourceroutingstatsentry.Sourceroutingstatsstatus}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsDroppedFrames"] = types.YLeaf{"Sourceroutingstatsdroppedframes", sourceroutingstatsentry.Sourceroutingstatsdroppedframes}
    sourceroutingstatsentry.EntityData.Leafs["sourceRoutingStatsCreateTime"] = types.YLeaf{"Sourceroutingstatscreatetime", sourceroutingstatsentry.Sourceroutingstatscreatetime}
    return &(sourceroutingstatsentry.EntityData)
}

