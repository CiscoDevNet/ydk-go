// Remote network monitoring devices, often called
// monitors or probes, are instruments that exist for
// the purpose of managing a network. This MIB defines
// objects for managing remote network monitoring devices.
package rmon_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rmon_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:RMON-MIB RMON-MIB}", reflect.TypeOf(RMONMIB{}))
    ydk.RegisterEntity("RMON-MIB:RMON-MIB", reflect.TypeOf(RMONMIB{}))
}

type RmonEventsV2 struct {
}

func (id RmonEventsV2) String() string {
	return "RMON-MIB:rmonEventsV2"
}

// EntryStatus represents never exercise these additional state transitions.
type EntryStatus string

const (
    EntryStatus_valid EntryStatus = "valid"

    EntryStatus_createRequest EntryStatus = "createRequest"

    EntryStatus_underCreation EntryStatus = "underCreation"

    EntryStatus_invalid EntryStatus = "invalid"
)

// RMONMIB
type RMONMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of Ethernet statistics entries.
    EtherStatsTable RMONMIB_EtherStatsTable

    // A list of history control entries.
    HistoryControlTable RMONMIB_HistoryControlTable

    // A list of Ethernet history entries.
    EtherHistoryTable RMONMIB_EtherHistoryTable

    // A list of alarm entries.
    AlarmTable RMONMIB_AlarmTable

    // A list of host table control entries.
    HostControlTable RMONMIB_HostControlTable

    // A list of host entries.
    HostTable RMONMIB_HostTable

    // A list of time-ordered host table entries.
    HostTimeTable RMONMIB_HostTimeTable

    // A list of top N host control entries.
    HostTopNControlTable RMONMIB_HostTopNControlTable

    // A list of top N host entries.
    HostTopNTable RMONMIB_HostTopNTable

    // A list of information entries for the traffic matrix on each interface.
    MatrixControlTable RMONMIB_MatrixControlTable

    // A list of traffic matrix entries indexed by source and destination MAC
    // address.
    MatrixSDTable RMONMIB_MatrixSDTable

    // A list of traffic matrix entries indexed by destination and source MAC
    // address.
    MatrixDSTable RMONMIB_MatrixDSTable

    // A list of packet filter entries.
    FilterTable RMONMIB_FilterTable

    // A list of packet channel entries.
    ChannelTable RMONMIB_ChannelTable

    // A list of buffers control entries.
    BufferControlTable RMONMIB_BufferControlTable

    // A list of packets captured off of a channel.
    CaptureBufferTable RMONMIB_CaptureBufferTable

    // A list of events to be generated.
    EventTable RMONMIB_EventTable

    // A list of events that have been logged.
    LogTable RMONMIB_LogTable
}

func (rMONMIB *RMONMIB) GetEntityData() *types.CommonEntityData {
    rMONMIB.EntityData.YFilter = rMONMIB.YFilter
    rMONMIB.EntityData.YangName = "RMON-MIB"
    rMONMIB.EntityData.BundleName = "cisco_ios_xe"
    rMONMIB.EntityData.ParentYangName = "RMON-MIB"
    rMONMIB.EntityData.SegmentPath = "RMON-MIB:RMON-MIB"
    rMONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rMONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rMONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rMONMIB.EntityData.Children = types.NewOrderedMap()
    rMONMIB.EntityData.Children.Append("etherStatsTable", types.YChild{"EtherStatsTable", &rMONMIB.EtherStatsTable})
    rMONMIB.EntityData.Children.Append("historyControlTable", types.YChild{"HistoryControlTable", &rMONMIB.HistoryControlTable})
    rMONMIB.EntityData.Children.Append("etherHistoryTable", types.YChild{"EtherHistoryTable", &rMONMIB.EtherHistoryTable})
    rMONMIB.EntityData.Children.Append("alarmTable", types.YChild{"AlarmTable", &rMONMIB.AlarmTable})
    rMONMIB.EntityData.Children.Append("hostControlTable", types.YChild{"HostControlTable", &rMONMIB.HostControlTable})
    rMONMIB.EntityData.Children.Append("hostTable", types.YChild{"HostTable", &rMONMIB.HostTable})
    rMONMIB.EntityData.Children.Append("hostTimeTable", types.YChild{"HostTimeTable", &rMONMIB.HostTimeTable})
    rMONMIB.EntityData.Children.Append("hostTopNControlTable", types.YChild{"HostTopNControlTable", &rMONMIB.HostTopNControlTable})
    rMONMIB.EntityData.Children.Append("hostTopNTable", types.YChild{"HostTopNTable", &rMONMIB.HostTopNTable})
    rMONMIB.EntityData.Children.Append("matrixControlTable", types.YChild{"MatrixControlTable", &rMONMIB.MatrixControlTable})
    rMONMIB.EntityData.Children.Append("matrixSDTable", types.YChild{"MatrixSDTable", &rMONMIB.MatrixSDTable})
    rMONMIB.EntityData.Children.Append("matrixDSTable", types.YChild{"MatrixDSTable", &rMONMIB.MatrixDSTable})
    rMONMIB.EntityData.Children.Append("filterTable", types.YChild{"FilterTable", &rMONMIB.FilterTable})
    rMONMIB.EntityData.Children.Append("channelTable", types.YChild{"ChannelTable", &rMONMIB.ChannelTable})
    rMONMIB.EntityData.Children.Append("bufferControlTable", types.YChild{"BufferControlTable", &rMONMIB.BufferControlTable})
    rMONMIB.EntityData.Children.Append("captureBufferTable", types.YChild{"CaptureBufferTable", &rMONMIB.CaptureBufferTable})
    rMONMIB.EntityData.Children.Append("eventTable", types.YChild{"EventTable", &rMONMIB.EventTable})
    rMONMIB.EntityData.Children.Append("logTable", types.YChild{"LogTable", &rMONMIB.LogTable})
    rMONMIB.EntityData.Leafs = types.NewOrderedMap()

    rMONMIB.EntityData.YListKeys = []string {}

    return &(rMONMIB.EntityData)
}

// RMONMIB_EtherStatsTable
// A list of Ethernet statistics entries.
type RMONMIB_EtherStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics kept for a particular Ethernet interface.  As an
    // example, an instance of the etherStatsPkts object might be named
    // etherStatsPkts.1. The type is slice of
    // RMONMIB_EtherStatsTable_EtherStatsEntry.
    EtherStatsEntry []*RMONMIB_EtherStatsTable_EtherStatsEntry
}

func (etherStatsTable *RMONMIB_EtherStatsTable) GetEntityData() *types.CommonEntityData {
    etherStatsTable.EntityData.YFilter = etherStatsTable.YFilter
    etherStatsTable.EntityData.YangName = "etherStatsTable"
    etherStatsTable.EntityData.BundleName = "cisco_ios_xe"
    etherStatsTable.EntityData.ParentYangName = "RMON-MIB"
    etherStatsTable.EntityData.SegmentPath = "etherStatsTable"
    etherStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherStatsTable.EntityData.Children = types.NewOrderedMap()
    etherStatsTable.EntityData.Children.Append("etherStatsEntry", types.YChild{"EtherStatsEntry", nil})
    for i := range etherStatsTable.EtherStatsEntry {
        etherStatsTable.EntityData.Children.Append(types.GetSegmentPath(etherStatsTable.EtherStatsEntry[i]), types.YChild{"EtherStatsEntry", etherStatsTable.EtherStatsEntry[i]})
    }
    etherStatsTable.EntityData.Leafs = types.NewOrderedMap()

    etherStatsTable.EntityData.YListKeys = []string {}

    return &(etherStatsTable.EntityData)
}

// RMONMIB_EtherStatsTable_EtherStatsEntry
// A collection of statistics kept for a particular
// Ethernet interface.  As an example, an instance of the
// etherStatsPkts object might be named etherStatsPkts.1
type RMONMIB_EtherStatsTable_EtherStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies this
    // etherStats entry. The type is interface{} with range: 1..65535.
    EtherStatsIndex interface{}

    // This object identifies the source of the data that this etherStats entry is
    // configured to analyze.  This source can be any ethernet interface on this
    // device. In order to identify a particular interface, this object shall
    // identify the instance of the ifIndex object, defined in RFC 2233 [17], for
    // the desired interface. For example, if an entry were to receive data from
    // interface #1, this object would be set to ifIndex.1.  The statistics in
    // this group reflect all packets on the local network segment attached to the
    // identified interface.  An agent may or may not be able to tell if
    // fundamental changes to the media of the interface have occurred and
    // necessitate an invalidation of this entry.  For example, a hot-pluggable
    // ethernet card could be pulled out and replaced by a token-ring card.  In
    // such a case, if the agent has such knowledge of the change, it is
    // recommended that it invalidate this entry.  This object may not be modified
    // if the associated etherStatsStatus object is equal to valid(1). The type is
    // string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    EtherStatsDataSource interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources. Note that this number is not necessarily the number
    // of packets dropped; it is just the number of times this condition has been
    // detected. The type is interface{} with range: 0..4294967295.
    EtherStatsDropEvents interface{}

    // The total number of octets of data (including those in bad packets)
    // received on the network (excluding framing bits but including FCS octets).
    // This object can be used as a reasonable estimate of 10-Megabit ethernet
    // utilization.  If greater precision is desired, the etherStatsPkts and
    // etherStatsOctets objects should be sampled before and after a common
    // interval.  The differences in the sampled values are Pkts and Octets,
    // respectively, and the number of seconds in the interval is Interval.  These
    // values are used to calculate the Utilization as follows:                  
    // Pkts * (9.6 + 6.4) + (Octets * .8)  Utilization =
    // -------------------------------------                          Interval *
    // 10,000  The result of this equation is the value Utilization which is the
    // percent utilization of the ethernet segment on a scale of 0 to 100 percent.
    // The type is interface{} with range: 0..4294967295. Units are Octets.
    EtherStatsOctets interface{}

    // The total number of packets (including bad packets, broadcast packets, and
    // multicast packets) received. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    EtherStatsPkts interface{}

    // The total number of good packets received that were directed to the
    // broadcast address.  Note that this does not include multicast packets. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    EtherStatsBroadcastPkts interface{}

    // The total number of good packets received that were directed to a multicast
    // address.  Note that this number does not include packets directed to the
    // broadcast address. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    EtherStatsMulticastPkts interface{}

    // The total number of packets received that had a length (excluding framing
    // bits, but including FCS octets) of between 64 and 1518 octets, inclusive,
    // but had either a bad Frame Check Sequence (FCS) with an integral number of
    // octets (FCS Error) or a bad FCS with a non-integral number of octets
    // (Alignment Error). The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    EtherStatsCRCAlignErrors interface{}

    // The total number of packets received that were less than 64 octets long
    // (excluding framing bits, but including FCS octets) and were otherwise well
    // formed. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    EtherStatsUndersizePkts interface{}

    // The total number of packets received that were longer than 1518 octets
    // (excluding framing bits, but including FCS octets) and were otherwise well
    // formed. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    EtherStatsOversizePkts interface{}

    // The total number of packets received that were less than 64 octets in
    // length (excluding framing bits but including FCS octets) and had either a
    // bad Frame Check Sequence (FCS) with an integral number of octets (FCS
    // Error) or a bad FCS with a non-integral number of octets (Alignment Error).
    // Note that it is entirely normal for etherStatsFragments to increment.  This
    // is because it counts both runts (which are normal occurrences due to
    // collisions) and noise hits. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    EtherStatsFragments interface{}

    // The total number of packets received that were longer than 1518 octets
    // (excluding framing bits, but including FCS octets), and had either a bad
    // Frame Check Sequence (FCS) with an integral number of octets (FCS Error) or
    // a bad FCS with a non-integral number of octets (Alignment Error).  Note
    // that this definition of jabber is different than the definition in
    // IEEE-802.3 section 8.2.1.5 (10BASE5) and section 10.3.1.4 (10BASE2).  These
    // documents define jabber as the condition where any packet exceeds 20 ms. 
    // The allowed range to detect jabber is between 20 ms and 150 ms. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    EtherStatsJabbers interface{}

    // The best estimate of the total number of collisions on this Ethernet
    // segment.  The value returned will depend on the location of the RMON probe.
    // Section 8.2.1.3 (10BASE-5) and section 10.3.1.3 (10BASE-2) of IEEE standard
    // 802.3 states that a station must detect a collision, in the receive mode,
    // if three or more stations are transmitting simultaneously.  A repeater port
    // must detect a collision when two or more stations are transmitting
    // simultaneously.  Thus a probe placed on a repeater port could record more
    // collisions than a probe connected to a station on the same segment would. 
    // Probe location plays a much smaller role when considering 10BASE-T. 
    // 14.2.1.4 (10BASE-T) of IEEE standard 802.3 defines a collision as the
    // simultaneous presence of signals on the DO and RD circuits (transmitting
    // and receiving at the same time).  A 10BASE-T station can only detect
    // collisions when it is transmitting.  Thus probes placed on a station and a
    // repeater, should report the same number of collisions.  Note also that an
    // RMON probe inside a repeater should ideally report collisions between the
    // repeater and one or more other hosts (transmit collisions as defined by
    // IEEE 802.3k) plus receiver collisions observed on any coax segments to
    // which the repeater is connected. The type is interface{} with range:
    // 0..4294967295. Units are Collisions.
    EtherStatsCollisions interface{}

    // The total number of packets (including bad packets) received that were 64
    // octets in length (excluding framing bits but including FCS octets). The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    EtherStatsPkts64Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 65 and 127 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    EtherStatsPkts65to127Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 128 and 255 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    EtherStatsPkts128to255Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 256 and 511 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    EtherStatsPkts256to511Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 512 and 1023 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    EtherStatsPkts512to1023Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 1024 and 1518 octets in length inclusive (excluding framing bits
    // but including FCS octets). The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    EtherStatsPkts1024to1518Octets interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    EtherStatsOwner interface{}

    // The status of this etherStats entry. The type is EntryStatus.
    EtherStatsStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted      because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    EtherStatsDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    EtherStatsCreateTime interface{}
}

func (etherStatsEntry *RMONMIB_EtherStatsTable_EtherStatsEntry) GetEntityData() *types.CommonEntityData {
    etherStatsEntry.EntityData.YFilter = etherStatsEntry.YFilter
    etherStatsEntry.EntityData.YangName = "etherStatsEntry"
    etherStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    etherStatsEntry.EntityData.ParentYangName = "etherStatsTable"
    etherStatsEntry.EntityData.SegmentPath = "etherStatsEntry" + types.AddKeyToken(etherStatsEntry.EtherStatsIndex, "etherStatsIndex")
    etherStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherStatsEntry.EntityData.Children = types.NewOrderedMap()
    etherStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    etherStatsEntry.EntityData.Leafs.Append("etherStatsIndex", types.YLeaf{"EtherStatsIndex", etherStatsEntry.EtherStatsIndex})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsDataSource", types.YLeaf{"EtherStatsDataSource", etherStatsEntry.EtherStatsDataSource})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsDropEvents", types.YLeaf{"EtherStatsDropEvents", etherStatsEntry.EtherStatsDropEvents})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsOctets", types.YLeaf{"EtherStatsOctets", etherStatsEntry.EtherStatsOctets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts", types.YLeaf{"EtherStatsPkts", etherStatsEntry.EtherStatsPkts})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsBroadcastPkts", types.YLeaf{"EtherStatsBroadcastPkts", etherStatsEntry.EtherStatsBroadcastPkts})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsMulticastPkts", types.YLeaf{"EtherStatsMulticastPkts", etherStatsEntry.EtherStatsMulticastPkts})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsCRCAlignErrors", types.YLeaf{"EtherStatsCRCAlignErrors", etherStatsEntry.EtherStatsCRCAlignErrors})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsUndersizePkts", types.YLeaf{"EtherStatsUndersizePkts", etherStatsEntry.EtherStatsUndersizePkts})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsOversizePkts", types.YLeaf{"EtherStatsOversizePkts", etherStatsEntry.EtherStatsOversizePkts})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsFragments", types.YLeaf{"EtherStatsFragments", etherStatsEntry.EtherStatsFragments})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsJabbers", types.YLeaf{"EtherStatsJabbers", etherStatsEntry.EtherStatsJabbers})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsCollisions", types.YLeaf{"EtherStatsCollisions", etherStatsEntry.EtherStatsCollisions})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts64Octets", types.YLeaf{"EtherStatsPkts64Octets", etherStatsEntry.EtherStatsPkts64Octets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts65to127Octets", types.YLeaf{"EtherStatsPkts65to127Octets", etherStatsEntry.EtherStatsPkts65to127Octets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts128to255Octets", types.YLeaf{"EtherStatsPkts128to255Octets", etherStatsEntry.EtherStatsPkts128to255Octets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts256to511Octets", types.YLeaf{"EtherStatsPkts256to511Octets", etherStatsEntry.EtherStatsPkts256to511Octets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts512to1023Octets", types.YLeaf{"EtherStatsPkts512to1023Octets", etherStatsEntry.EtherStatsPkts512to1023Octets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsPkts1024to1518Octets", types.YLeaf{"EtherStatsPkts1024to1518Octets", etherStatsEntry.EtherStatsPkts1024to1518Octets})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsOwner", types.YLeaf{"EtherStatsOwner", etherStatsEntry.EtherStatsOwner})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsStatus", types.YLeaf{"EtherStatsStatus", etherStatsEntry.EtherStatsStatus})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsDroppedFrames", types.YLeaf{"EtherStatsDroppedFrames", etherStatsEntry.EtherStatsDroppedFrames})
    etherStatsEntry.EntityData.Leafs.Append("etherStatsCreateTime", types.YLeaf{"EtherStatsCreateTime", etherStatsEntry.EtherStatsCreateTime})

    etherStatsEntry.EntityData.YListKeys = []string {"EtherStatsIndex"}

    return &(etherStatsEntry.EntityData)
}

// RMONMIB_HistoryControlTable
// A list of history control entries.
type RMONMIB_HistoryControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up a periodic sampling of statistics.  As an
    // example, an instance of the historyControlInterval object might be named
    // historyControlInterval.2. The type is slice of
    // RMONMIB_HistoryControlTable_HistoryControlEntry.
    HistoryControlEntry []*RMONMIB_HistoryControlTable_HistoryControlEntry
}

func (historyControlTable *RMONMIB_HistoryControlTable) GetEntityData() *types.CommonEntityData {
    historyControlTable.EntityData.YFilter = historyControlTable.YFilter
    historyControlTable.EntityData.YangName = "historyControlTable"
    historyControlTable.EntityData.BundleName = "cisco_ios_xe"
    historyControlTable.EntityData.ParentYangName = "RMON-MIB"
    historyControlTable.EntityData.SegmentPath = "historyControlTable"
    historyControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    historyControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    historyControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    historyControlTable.EntityData.Children = types.NewOrderedMap()
    historyControlTable.EntityData.Children.Append("historyControlEntry", types.YChild{"HistoryControlEntry", nil})
    for i := range historyControlTable.HistoryControlEntry {
        historyControlTable.EntityData.Children.Append(types.GetSegmentPath(historyControlTable.HistoryControlEntry[i]), types.YChild{"HistoryControlEntry", historyControlTable.HistoryControlEntry[i]})
    }
    historyControlTable.EntityData.Leafs = types.NewOrderedMap()

    historyControlTable.EntityData.YListKeys = []string {}

    return &(historyControlTable.EntityData)
}

// RMONMIB_HistoryControlTable_HistoryControlEntry
// A list of parameters that set up a periodic sampling of
// statistics.  As an example, an instance of the
// historyControlInterval object might be named
// historyControlInterval.2
type RMONMIB_HistoryControlTable_HistoryControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // historyControl table.  Each such entry defines a set of samples at a
    // particular interval for an interface on the device. The type is interface{}
    // with range: 1..65535.
    HistoryControlIndex interface{}

    // This object identifies the source of the data for which historical data was
    // collected and placed in a media-specific table on behalf of this
    // historyControlEntry.  This source can be any interface on this device.  In
    // order to identify a particular interface, this object shall identify the
    // instance of the ifIndex object, defined in  RFC 2233 [17], for the desired
    // interface. For example, if an entry were to receive data from interface #1,
    // this object would be set to ifIndex.1.  The statistics in this group
    // reflect all packets on the local network segment attached to the identified
    // interface.  An agent may or may not be able to tell if fundamental changes
    // to the media of the interface have occurred and necessitate an invalidation
    // of this entry.  For example, a hot-pluggable ethernet card could be pulled
    // out and replaced by a token-ring card.  In such a case, if the agent has
    // such knowledge of the change, it is recommended that it invalidate this
    // entry.  This object may not be modified if the associated
    // historyControlStatus object is equal to valid(1). The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    HistoryControlDataSource interface{}

    // The requested number of discrete time intervals over which data is to be
    // saved in the part of the media-specific table associated with this
    // historyControlEntry.  When this object is created or modified, the probe
    // should set historyControlBucketsGranted as closely to this object as is
    // possible for the particular probe implementation and available resources.
    // The type is interface{} with range: 1..65535.
    HistoryControlBucketsRequested interface{}

    // The number of discrete sampling intervals over which data shall be saved in
    // the part of the media-specific table associated with this
    // historyControlEntry. When the associated historyControlBucketsRequested
    // object is created or modified, the probe should set this object as closely
    // to the requested value as is possible for the particular probe
    // implementation and available resources.  The probe must not lower this
    // value except as a result of a modification to the associated
    // historyControlBucketsRequested object.  There will be times when the actual
    // number of buckets associated with this entry is less than the value of this
    // object.  In this case, at the end of each sampling interval, a new bucket
    // will be added to the media-specific table.  When the number of buckets
    // reaches the value of this object and a new bucket is to be added to the
    // media-specific table, the oldest bucket associated with this
    // historyControlEntry shall be deleted by the agent so that the new bucket
    // can be added.  When the value of this object changes to a value less than
    // the current value, entries are deleted from the media-specific table
    // associated with this historyControlEntry.  Enough of the oldest of these
    // entries shall be deleted by the agent so that their number remains less
    // than or equal to the new value of this object.  When the value of this
    // object changes to a value greater than the current value, the number of
    // associated media- specific entries may be allowed to grow. The type is
    // interface{} with range: 1..65535.
    HistoryControlBucketsGranted interface{}

    // The interval in seconds over which the data is sampled for each bucket in
    // the part of the media-specific table associated with this
    // historyControlEntry.  This interval can be set to any number of seconds
    // between 1 and 3600 (1 hour).  Because the counters in a bucket may overflow
    // at their maximum value with no indication, a prudent manager will take into
    // account the possibility of overflow in any of the associated counters.  It
    // is important to consider the minimum time in which any counter could
    // overflow on a particular media type and set the historyControlInterval
    // object to a value less than this interval.  This is typically most
    // important for the 'octets' counter in any media-specific table.  For
    // example, on an Ethernet network, the etherHistoryOctets counter could
    // overflow in about one hour at the Ethernet's maximum utilization.  This
    // object may not be modified if the associated historyControlStatus object is
    // equal to valid(1). The type is interface{} with range: 1..3600. Units are
    // Seconds.
    HistoryControlInterval interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    HistoryControlOwner interface{}

    // The status of this historyControl entry.  Each instance of the
    // media-specific table associated with this historyControlEntry will be
    // deleted by the agent if this historyControlEntry is not equal to valid(1).
    // The type is EntryStatus.
    HistoryControlStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this      collection.  This count does not include packets that were
    // not counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    HistoryControlDroppedFrames interface{}
}

func (historyControlEntry *RMONMIB_HistoryControlTable_HistoryControlEntry) GetEntityData() *types.CommonEntityData {
    historyControlEntry.EntityData.YFilter = historyControlEntry.YFilter
    historyControlEntry.EntityData.YangName = "historyControlEntry"
    historyControlEntry.EntityData.BundleName = "cisco_ios_xe"
    historyControlEntry.EntityData.ParentYangName = "historyControlTable"
    historyControlEntry.EntityData.SegmentPath = "historyControlEntry" + types.AddKeyToken(historyControlEntry.HistoryControlIndex, "historyControlIndex")
    historyControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    historyControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    historyControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    historyControlEntry.EntityData.Children = types.NewOrderedMap()
    historyControlEntry.EntityData.Leafs = types.NewOrderedMap()
    historyControlEntry.EntityData.Leafs.Append("historyControlIndex", types.YLeaf{"HistoryControlIndex", historyControlEntry.HistoryControlIndex})
    historyControlEntry.EntityData.Leafs.Append("historyControlDataSource", types.YLeaf{"HistoryControlDataSource", historyControlEntry.HistoryControlDataSource})
    historyControlEntry.EntityData.Leafs.Append("historyControlBucketsRequested", types.YLeaf{"HistoryControlBucketsRequested", historyControlEntry.HistoryControlBucketsRequested})
    historyControlEntry.EntityData.Leafs.Append("historyControlBucketsGranted", types.YLeaf{"HistoryControlBucketsGranted", historyControlEntry.HistoryControlBucketsGranted})
    historyControlEntry.EntityData.Leafs.Append("historyControlInterval", types.YLeaf{"HistoryControlInterval", historyControlEntry.HistoryControlInterval})
    historyControlEntry.EntityData.Leafs.Append("historyControlOwner", types.YLeaf{"HistoryControlOwner", historyControlEntry.HistoryControlOwner})
    historyControlEntry.EntityData.Leafs.Append("historyControlStatus", types.YLeaf{"HistoryControlStatus", historyControlEntry.HistoryControlStatus})
    historyControlEntry.EntityData.Leafs.Append("historyControlDroppedFrames", types.YLeaf{"HistoryControlDroppedFrames", historyControlEntry.HistoryControlDroppedFrames})

    historyControlEntry.EntityData.YListKeys = []string {"HistoryControlIndex"}

    return &(historyControlEntry.EntityData)
}

// RMONMIB_EtherHistoryTable
// A list of Ethernet history entries.
type RMONMIB_EtherHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An historical sample of Ethernet statistics on a particular Ethernet
    // interface.  This sample is associated with the historyControlEntry which
    // set up the parameters for a regular collection of these samples.  As an
    // example, an instance of the etherHistoryPkts object might be named
    // etherHistoryPkts.2.89. The type is slice of
    // RMONMIB_EtherHistoryTable_EtherHistoryEntry.
    EtherHistoryEntry []*RMONMIB_EtherHistoryTable_EtherHistoryEntry
}

func (etherHistoryTable *RMONMIB_EtherHistoryTable) GetEntityData() *types.CommonEntityData {
    etherHistoryTable.EntityData.YFilter = etherHistoryTable.YFilter
    etherHistoryTable.EntityData.YangName = "etherHistoryTable"
    etherHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    etherHistoryTable.EntityData.ParentYangName = "RMON-MIB"
    etherHistoryTable.EntityData.SegmentPath = "etherHistoryTable"
    etherHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherHistoryTable.EntityData.Children = types.NewOrderedMap()
    etherHistoryTable.EntityData.Children.Append("etherHistoryEntry", types.YChild{"EtherHistoryEntry", nil})
    for i := range etherHistoryTable.EtherHistoryEntry {
        etherHistoryTable.EntityData.Children.Append(types.GetSegmentPath(etherHistoryTable.EtherHistoryEntry[i]), types.YChild{"EtherHistoryEntry", etherHistoryTable.EtherHistoryEntry[i]})
    }
    etherHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    etherHistoryTable.EntityData.YListKeys = []string {}

    return &(etherHistoryTable.EntityData)
}

// RMONMIB_EtherHistoryTable_EtherHistoryEntry
// An historical sample of Ethernet statistics on a particular
// Ethernet interface.  This sample is associated with the
// historyControlEntry which set up the parameters for
// a regular collection of these samples.  As an example, an
// instance of the etherHistoryPkts object might be named
// etherHistoryPkts.2.89
type RMONMIB_EtherHistoryTable_EtherHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The history of which this entry is a part.  The
    // history identified by a particular value of this index is the same history
    // as identified by the same value of historyControlIndex. The type is
    // interface{} with range: 1..65535.
    EtherHistoryIndex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // sample this entry represents among all samples associated with the same
    // historyControlEntry. This index starts at 1 and increases by one as each
    // new sample is taken. The type is interface{} with range: 1..2147483647.
    EtherHistorySampleIndex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval.  Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    EtherHistoryIntervalStart interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources during this sampling interval.  Note that this number
    // is not necessarily the number of packets dropped, it is just the number of
    // times this condition has been detected. The type is interface{} with range:
    // 0..4294967295.
    EtherHistoryDropEvents interface{}

    // The total number of octets of data (including those in bad packets)
    // received on the network (excluding framing bits but including FCS octets).
    // The type is interface{} with range: 0..4294967295. Units are Octets.
    EtherHistoryOctets interface{}

    // The number of packets (including bad packets) received during this sampling
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    EtherHistoryPkts interface{}

    // The number of good packets received during this sampling interval that were
    // directed to the broadcast address. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    EtherHistoryBroadcastPkts interface{}

    // The number of good packets received during this sampling interval that were
    // directed to a multicast address.  Note that this number does not include
    // packets addressed to the broadcast address. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    EtherHistoryMulticastPkts interface{}

    // The number of packets received during this sampling interval that had a
    // length (excluding framing bits but including FCS octets) between 64 and
    // 1518 octets, inclusive, but had either a bad Frame Check Sequence (FCS)
    // with an integral number of octets (FCS Error) or a bad FCS with a
    // non-integral number of octets (Alignment Error). The type is interface{}
    // with range: 0..4294967295. Units are Packets.
    EtherHistoryCRCAlignErrors interface{}

    // The number of packets received during this sampling interval that were less
    // than 64 octets long (excluding framing bits but including FCS octets) and
    // were otherwise well formed. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    EtherHistoryUndersizePkts interface{}

    // The number of packets received during this sampling interval that were
    // longer than 1518 octets (excluding framing bits but including FCS octets)
    // but were otherwise well formed. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    EtherHistoryOversizePkts interface{}

    // The total number of packets received during this sampling interval that
    // were less than 64 octets in length (excluding framing bits but including
    // FCS octets) had either a bad Frame Check Sequence (FCS) with an integral
    // number of octets (FCS Error) or a bad FCS with a non-integral number of
    // octets (Alignment Error).  Note that it is entirely normal for
    // etherHistoryFragments to increment.  This is because it counts both runts
    // (which are normal occurrences due to collisions) and noise hits. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    EtherHistoryFragments interface{}

    // The number of packets received during this sampling interval that were
    // longer than 1518 octets (excluding framing bits but including FCS octets),
    // and  had either a bad Frame Check Sequence (FCS) with an integral number of
    // octets (FCS Error) or a bad FCS with a non-integral number of octets
    // (Alignment Error).  Note that this definition of jabber is different than
    // the definition in IEEE-802.3 section 8.2.1.5 (10BASE5) and section 10.3.1.4
    // (10BASE2).  These documents define jabber as the condition where any packet
    // exceeds 20 ms.  The allowed range to detect jabber is between 20 ms and 150
    // ms. The type is interface{} with range: 0..4294967295. Units are Packets.
    EtherHistoryJabbers interface{}

    // The best estimate of the total number of collisions on this Ethernet
    // segment during this sampling interval.  The value returned will depend on
    // the location of the RMON probe. Section 8.2.1.3 (10BASE-5) and section
    // 10.3.1.3 (10BASE-2) of IEEE standard 802.3 states that a station must
    // detect a collision, in the receive mode, if three or more stations are
    // transmitting simultaneously.  A repeater port must detect a collision when
    // two or more stations are transmitting simultaneously.  Thus a probe placed
    // on a repeater port could record more collisions than a probe connected to a
    // station on the same segment would.  Probe location plays a much smaller
    // role when considering 10BASE-T.  14.2.1.4 (10BASE-T) of IEEE standard 802.3
    // defines a collision as the simultaneous presence of signals on the DO and
    // RD circuits (transmitting and receiving at the same time).  A 10BASE-T
    // station can only detect collisions when it is transmitting.  Thus probes
    // placed on a station and a repeater, should report the same number of
    // collisions.  Note also that an RMON probe inside a repeater should ideally
    // report collisions between the repeater and one or more other hosts
    // (transmit collisions as defined by IEEE 802.3k) plus receiver collisions
    // observed on any coax segments to which the repeater is connected. The type
    // is interface{} with range: 0..4294967295. Units are Collisions.
    EtherHistoryCollisions interface{}

    // The best estimate of the mean physical layer network utilization on this
    // interface during this sampling interval, in hundredths of a percent. The
    // type is interface{} with range: 0..10000.
    EtherHistoryUtilization interface{}
}

func (etherHistoryEntry *RMONMIB_EtherHistoryTable_EtherHistoryEntry) GetEntityData() *types.CommonEntityData {
    etherHistoryEntry.EntityData.YFilter = etherHistoryEntry.YFilter
    etherHistoryEntry.EntityData.YangName = "etherHistoryEntry"
    etherHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    etherHistoryEntry.EntityData.ParentYangName = "etherHistoryTable"
    etherHistoryEntry.EntityData.SegmentPath = "etherHistoryEntry" + types.AddKeyToken(etherHistoryEntry.EtherHistoryIndex, "etherHistoryIndex") + types.AddKeyToken(etherHistoryEntry.EtherHistorySampleIndex, "etherHistorySampleIndex")
    etherHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherHistoryEntry.EntityData.Children = types.NewOrderedMap()
    etherHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryIndex", types.YLeaf{"EtherHistoryIndex", etherHistoryEntry.EtherHistoryIndex})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistorySampleIndex", types.YLeaf{"EtherHistorySampleIndex", etherHistoryEntry.EtherHistorySampleIndex})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryIntervalStart", types.YLeaf{"EtherHistoryIntervalStart", etherHistoryEntry.EtherHistoryIntervalStart})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryDropEvents", types.YLeaf{"EtherHistoryDropEvents", etherHistoryEntry.EtherHistoryDropEvents})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryOctets", types.YLeaf{"EtherHistoryOctets", etherHistoryEntry.EtherHistoryOctets})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryPkts", types.YLeaf{"EtherHistoryPkts", etherHistoryEntry.EtherHistoryPkts})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryBroadcastPkts", types.YLeaf{"EtherHistoryBroadcastPkts", etherHistoryEntry.EtherHistoryBroadcastPkts})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryMulticastPkts", types.YLeaf{"EtherHistoryMulticastPkts", etherHistoryEntry.EtherHistoryMulticastPkts})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryCRCAlignErrors", types.YLeaf{"EtherHistoryCRCAlignErrors", etherHistoryEntry.EtherHistoryCRCAlignErrors})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryUndersizePkts", types.YLeaf{"EtherHistoryUndersizePkts", etherHistoryEntry.EtherHistoryUndersizePkts})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryOversizePkts", types.YLeaf{"EtherHistoryOversizePkts", etherHistoryEntry.EtherHistoryOversizePkts})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryFragments", types.YLeaf{"EtherHistoryFragments", etherHistoryEntry.EtherHistoryFragments})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryJabbers", types.YLeaf{"EtherHistoryJabbers", etherHistoryEntry.EtherHistoryJabbers})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryCollisions", types.YLeaf{"EtherHistoryCollisions", etherHistoryEntry.EtherHistoryCollisions})
    etherHistoryEntry.EntityData.Leafs.Append("etherHistoryUtilization", types.YLeaf{"EtherHistoryUtilization", etherHistoryEntry.EtherHistoryUtilization})

    etherHistoryEntry.EntityData.YListKeys = []string {"EtherHistoryIndex", "EtherHistorySampleIndex"}

    return &(etherHistoryEntry.EntityData)
}

// RMONMIB_AlarmTable
// A list of alarm entries.
type RMONMIB_AlarmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up a periodic checking for alarm conditions. 
    // For example, an instance of the alarmValue object might be named
    // alarmValue.8. The type is slice of RMONMIB_AlarmTable_AlarmEntry.
    AlarmEntry []*RMONMIB_AlarmTable_AlarmEntry
}

func (alarmTable *RMONMIB_AlarmTable) GetEntityData() *types.CommonEntityData {
    alarmTable.EntityData.YFilter = alarmTable.YFilter
    alarmTable.EntityData.YangName = "alarmTable"
    alarmTable.EntityData.BundleName = "cisco_ios_xe"
    alarmTable.EntityData.ParentYangName = "RMON-MIB"
    alarmTable.EntityData.SegmentPath = "alarmTable"
    alarmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alarmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alarmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alarmTable.EntityData.Children = types.NewOrderedMap()
    alarmTable.EntityData.Children.Append("alarmEntry", types.YChild{"AlarmEntry", nil})
    for i := range alarmTable.AlarmEntry {
        alarmTable.EntityData.Children.Append(types.GetSegmentPath(alarmTable.AlarmEntry[i]), types.YChild{"AlarmEntry", alarmTable.AlarmEntry[i]})
    }
    alarmTable.EntityData.Leafs = types.NewOrderedMap()

    alarmTable.EntityData.YListKeys = []string {}

    return &(alarmTable.EntityData)
}

// RMONMIB_AlarmTable_AlarmEntry
// A list of parameters that set up a periodic checking
// for alarm conditions.  For example, an instance of the
// alarmValue object might be named alarmValue.8
type RMONMIB_AlarmTable_AlarmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // alarm table.  Each such entry defines a diagnostic sample at a particular
    // interval for an object on the device. The type is interface{} with range:
    // 1..65535.
    AlarmIndex interface{}

    // The interval in seconds over which the data is sampled and compared with
    // the rising and falling thresholds.  When setting this variable, care should
    // be taken in the case of deltaValue sampling - the interval should be set
    // short enough that the sampled variable is very unlikely to increase or
    // decrease by more than 2^31 - 1 during a single sampling interval.  This
    // object may not be modified if the associated alarmStatus object is equal to
    // valid(1). The type is interface{} with range: -2147483648..2147483647.
    // Units are Seconds.
    AlarmInterval interface{}

    // The object identifier of the particular variable to be sampled.  Only
    // variables that resolve to an ASN.1 primitive type of INTEGER (INTEGER,
    // Integer32, Counter32, Counter64, Gauge, or TimeTicks) may be sampled. 
    // Because SNMP access control is articulated entirely in terms of the
    // contents of MIB views, no access control mechanism exists that can restrict
    // the value of this object to identify only those objects that exist in a
    // particular MIB view.  Because there is thus no acceptable means of
    // restricting the read access that could be obtained through the alarm
    // mechanism, the probe must only grant write access to this object in those
    // views that have read access to all objects on the probe.  During a set
    // operation, if the supplied variable name is not available in the selected
    // MIB view, a badValue error must be returned.  If at any time the variable
    // name of an established alarmEntry is no longer available in the selected
    // MIB view, the probe must change the status of this alarmEntry to
    // invalid(4).  This object may not be modified if the associated alarmStatus
    // object is equal to valid(1). The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    AlarmVariable interface{}

    // The method of sampling the selected variable and calculating the value to
    // be compared against the thresholds.  If the value of this object is
    // absoluteValue(1), the value of the selected variable will be compared
    // directly with the thresholds at the end of the sampling interval.  If the
    // value of this object is deltaValue(2), the value of the selected variable
    // at the last sample will be subtracted from the current value, and the
    // difference compared with the thresholds.  This object may not be modified
    // if the associated alarmStatus object is equal to valid(1). The type is
    // AlarmSampleType.
    AlarmSampleType interface{}

    // The value of the statistic during the last sampling period.  For example,
    // if the sample type is deltaValue, this value will be the difference between
    // the samples at the beginning and end of the period.  If the sample type is
    // absoluteValue, this value will be the sampled value at the end of the
    // period. This is the value that is compared with the rising and falling
    // thresholds.  The value during the current sampling period is not made
    // available until the period is completed and will remain available until the
    // next period completes. The type is interface{} with range:
    // -2147483648..2147483647.
    AlarmValue interface{}

    // The alarm that may be sent when this entry is first set to valid.  If the
    // first sample after this entry becomes valid is greater than or equal to the
    // risingThreshold and alarmStartupAlarm is equal to risingAlarm(1) or
    // risingOrFallingAlarm(3), then a single rising alarm will be generated.  If
    // the first sample after this entry becomes valid is less than or equal to
    // the fallingThreshold and alarmStartupAlarm is equal to fallingAlarm(2) or
    // risingOrFallingAlarm(3), then a single falling alarm will be generated. 
    // This object may not be modified if the associated alarmStatus object is
    // equal to valid(1). The type is AlarmStartupAlarm.
    AlarmStartupAlarm interface{}

    // A threshold for the sampled statistic.  When the current sampled value is
    // greater than or equal to this threshold, and the value at the last sampling
    // interval was less than this threshold, a single event will be generated. A
    // single event will also be generated if the first sample after this entry
    // becomes valid is greater than or equal to this threshold and the associated
    // alarmStartupAlarm is equal to risingAlarm(1) or risingOrFallingAlarm(3). 
    // After a rising event is generated, another such event will not be generated
    // until the sampled value falls below this threshold and reaches the
    // alarmFallingThreshold.  This object may not be modified if the associated
    // alarmStatus object is equal to valid(1). The type is interface{} with
    // range: -2147483648..2147483647.
    AlarmRisingThreshold interface{}

    // A threshold for the sampled statistic.  When the current sampled value is
    // less than or equal to this threshold, and the value at the last sampling
    // interval was greater than this threshold, a single event will be generated.
    // A single event will also be generated if the first sample after this entry
    // becomes valid is less than or equal to this threshold and the associated
    // alarmStartupAlarm is equal to fallingAlarm(2) or risingOrFallingAlarm(3). 
    // After a falling event is generated, another such event will not be
    // generated until the sampled value rises above this threshold and reaches
    // the alarmRisingThreshold.  This object may not be modified if the
    // associated alarmStatus object is equal to valid(1). The type is interface{}
    // with range: -2147483648..2147483647.
    AlarmFallingThreshold interface{}

    // The index of the eventEntry that is used when a rising threshold is
    // crossed.  The eventEntry identified by a particular value of this index is
    // the same as identified by the same value of the eventIndex object.  If
    // there is no corresponding entry in the eventTable, then no association
    // exists.  In particular, if this value is zero, no associated event will be
    // generated, as zero is not a valid event index.  This object may not be
    // modified if the associated alarmStatus object is equal to valid(1). The
    // type is interface{} with range: 0..65535.
    AlarmRisingEventIndex interface{}

    // The index of the eventEntry that is used when a falling threshold is
    // crossed.  The eventEntry identified by a particular value of this index is
    // the same as identified by the same value of the eventIndex object.  If
    // there is no corresponding entry in the eventTable, then no association
    // exists.  In particular, if this value is zero, no associated event will be
    // generated, as zero is not a valid event index.  This object may not be
    // modified if the associated alarmStatus object is equal to valid(1). The
    // type is interface{} with range: 0..65535.
    AlarmFallingEventIndex interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    AlarmOwner interface{}

    // The status of this alarm entry. The type is EntryStatus.
    AlarmStatus interface{}
}

func (alarmEntry *RMONMIB_AlarmTable_AlarmEntry) GetEntityData() *types.CommonEntityData {
    alarmEntry.EntityData.YFilter = alarmEntry.YFilter
    alarmEntry.EntityData.YangName = "alarmEntry"
    alarmEntry.EntityData.BundleName = "cisco_ios_xe"
    alarmEntry.EntityData.ParentYangName = "alarmTable"
    alarmEntry.EntityData.SegmentPath = "alarmEntry" + types.AddKeyToken(alarmEntry.AlarmIndex, "alarmIndex")
    alarmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alarmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alarmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alarmEntry.EntityData.Children = types.NewOrderedMap()
    alarmEntry.EntityData.Leafs = types.NewOrderedMap()
    alarmEntry.EntityData.Leafs.Append("alarmIndex", types.YLeaf{"AlarmIndex", alarmEntry.AlarmIndex})
    alarmEntry.EntityData.Leafs.Append("alarmInterval", types.YLeaf{"AlarmInterval", alarmEntry.AlarmInterval})
    alarmEntry.EntityData.Leafs.Append("alarmVariable", types.YLeaf{"AlarmVariable", alarmEntry.AlarmVariable})
    alarmEntry.EntityData.Leafs.Append("alarmSampleType", types.YLeaf{"AlarmSampleType", alarmEntry.AlarmSampleType})
    alarmEntry.EntityData.Leafs.Append("alarmValue", types.YLeaf{"AlarmValue", alarmEntry.AlarmValue})
    alarmEntry.EntityData.Leafs.Append("alarmStartupAlarm", types.YLeaf{"AlarmStartupAlarm", alarmEntry.AlarmStartupAlarm})
    alarmEntry.EntityData.Leafs.Append("alarmRisingThreshold", types.YLeaf{"AlarmRisingThreshold", alarmEntry.AlarmRisingThreshold})
    alarmEntry.EntityData.Leafs.Append("alarmFallingThreshold", types.YLeaf{"AlarmFallingThreshold", alarmEntry.AlarmFallingThreshold})
    alarmEntry.EntityData.Leafs.Append("alarmRisingEventIndex", types.YLeaf{"AlarmRisingEventIndex", alarmEntry.AlarmRisingEventIndex})
    alarmEntry.EntityData.Leafs.Append("alarmFallingEventIndex", types.YLeaf{"AlarmFallingEventIndex", alarmEntry.AlarmFallingEventIndex})
    alarmEntry.EntityData.Leafs.Append("alarmOwner", types.YLeaf{"AlarmOwner", alarmEntry.AlarmOwner})
    alarmEntry.EntityData.Leafs.Append("alarmStatus", types.YLeaf{"AlarmStatus", alarmEntry.AlarmStatus})

    alarmEntry.EntityData.YListKeys = []string {"AlarmIndex"}

    return &(alarmEntry.EntityData)
}

// RMONMIB_AlarmTable_AlarmEntry_AlarmSampleType represents alarmStatus object is equal to valid(1).
type RMONMIB_AlarmTable_AlarmEntry_AlarmSampleType string

const (
    RMONMIB_AlarmTable_AlarmEntry_AlarmSampleType_absoluteValue RMONMIB_AlarmTable_AlarmEntry_AlarmSampleType = "absoluteValue"

    RMONMIB_AlarmTable_AlarmEntry_AlarmSampleType_deltaValue RMONMIB_AlarmTable_AlarmEntry_AlarmSampleType = "deltaValue"
)

// RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm represents alarmStatus object is equal to valid(1).
type RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm string

const (
    RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm_risingAlarm RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm = "risingAlarm"

    RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm_fallingAlarm RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm = "fallingAlarm"

    RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm_risingOrFallingAlarm RMONMIB_AlarmTable_AlarmEntry_AlarmStartupAlarm = "risingOrFallingAlarm"
)

// RMONMIB_HostControlTable
// A list of host table control entries.
type RMONMIB_HostControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up the discovery of hosts on a particular
    // interface and the collection of statistics about these hosts.  For example,
    // an instance of the hostControlTableSize object might be named
    // hostControlTableSize.1. The type is slice of
    // RMONMIB_HostControlTable_HostControlEntry.
    HostControlEntry []*RMONMIB_HostControlTable_HostControlEntry
}

func (hostControlTable *RMONMIB_HostControlTable) GetEntityData() *types.CommonEntityData {
    hostControlTable.EntityData.YFilter = hostControlTable.YFilter
    hostControlTable.EntityData.YangName = "hostControlTable"
    hostControlTable.EntityData.BundleName = "cisco_ios_xe"
    hostControlTable.EntityData.ParentYangName = "RMON-MIB"
    hostControlTable.EntityData.SegmentPath = "hostControlTable"
    hostControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostControlTable.EntityData.Children = types.NewOrderedMap()
    hostControlTable.EntityData.Children.Append("hostControlEntry", types.YChild{"HostControlEntry", nil})
    for i := range hostControlTable.HostControlEntry {
        hostControlTable.EntityData.Children.Append(types.GetSegmentPath(hostControlTable.HostControlEntry[i]), types.YChild{"HostControlEntry", hostControlTable.HostControlEntry[i]})
    }
    hostControlTable.EntityData.Leafs = types.NewOrderedMap()

    hostControlTable.EntityData.YListKeys = []string {}

    return &(hostControlTable.EntityData)
}

// RMONMIB_HostControlTable_HostControlEntry
// A list of parameters that set up the discovery of hosts
// on a particular interface and the collection of statistics
// about these hosts.  For example, an instance of the
// hostControlTableSize object might be named
// hostControlTableSize.1
type RMONMIB_HostControlTable_HostControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostControl table.  Each such entry defines a function that discovers hosts
    // on a particular interface and places statistics about them in the hostTable
    // and the hostTimeTable on behalf of this hostControlEntry. The type is
    // interface{} with range: 1..65535.
    HostControlIndex interface{}

    // This object identifies the source of the data for this instance of the host
    // function.  This source can be any interface on this device.  In order to
    // identify a particular interface, this object shall identify the instance of
    // the ifIndex object, defined in RFC 2233 [17], for the desired interface.
    // For example, if an entry were to receive data from interface #1, this
    // object would be set to ifIndex.1.  The statistics in this group reflect all
    // packets on the local network segment attached to the identified interface. 
    // An agent may or may not be able to tell if fundamental changes to the media
    // of the interface have occurred and necessitate an invalidation of this
    // entry.  For example, a hot-pluggable ethernet card could be pulled out and
    // replaced by a token-ring card.  In such a case, if the agent has such
    // knowledge of the change, it is recommended that it invalidate this entry. 
    // This object may not be modified if the associated hostControlStatus object
    // is equal to valid(1). The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    HostControlDataSource interface{}

    // The number of hostEntries in the hostTable and the hostTimeTable associated
    // with this hostControlEntry. The type is interface{} with range:
    // -2147483648..2147483647.
    HostControlTableSize interface{}

    // The value of sysUpTime when the last entry was deleted from the portion of
    // the hostTable associated with this hostControlEntry.  If no deletions have
    // occurred, this value shall be zero. The type is interface{} with range:
    // 0..4294967295.
    HostControlLastDeleteTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    HostControlOwner interface{}

    // The status of this hostControl entry.  If this object is not equal to
    // valid(1), all associated entries in the hostTable, hostTimeTable, and the
    // hostTopNTable shall be deleted by the agent. The type is EntryStatus.
    HostControlStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    HostControlDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    HostControlCreateTime interface{}
}

func (hostControlEntry *RMONMIB_HostControlTable_HostControlEntry) GetEntityData() *types.CommonEntityData {
    hostControlEntry.EntityData.YFilter = hostControlEntry.YFilter
    hostControlEntry.EntityData.YangName = "hostControlEntry"
    hostControlEntry.EntityData.BundleName = "cisco_ios_xe"
    hostControlEntry.EntityData.ParentYangName = "hostControlTable"
    hostControlEntry.EntityData.SegmentPath = "hostControlEntry" + types.AddKeyToken(hostControlEntry.HostControlIndex, "hostControlIndex")
    hostControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostControlEntry.EntityData.Children = types.NewOrderedMap()
    hostControlEntry.EntityData.Leafs = types.NewOrderedMap()
    hostControlEntry.EntityData.Leafs.Append("hostControlIndex", types.YLeaf{"HostControlIndex", hostControlEntry.HostControlIndex})
    hostControlEntry.EntityData.Leafs.Append("hostControlDataSource", types.YLeaf{"HostControlDataSource", hostControlEntry.HostControlDataSource})
    hostControlEntry.EntityData.Leafs.Append("hostControlTableSize", types.YLeaf{"HostControlTableSize", hostControlEntry.HostControlTableSize})
    hostControlEntry.EntityData.Leafs.Append("hostControlLastDeleteTime", types.YLeaf{"HostControlLastDeleteTime", hostControlEntry.HostControlLastDeleteTime})
    hostControlEntry.EntityData.Leafs.Append("hostControlOwner", types.YLeaf{"HostControlOwner", hostControlEntry.HostControlOwner})
    hostControlEntry.EntityData.Leafs.Append("hostControlStatus", types.YLeaf{"HostControlStatus", hostControlEntry.HostControlStatus})
    hostControlEntry.EntityData.Leafs.Append("hostControlDroppedFrames", types.YLeaf{"HostControlDroppedFrames", hostControlEntry.HostControlDroppedFrames})
    hostControlEntry.EntityData.Leafs.Append("hostControlCreateTime", types.YLeaf{"HostControlCreateTime", hostControlEntry.HostControlCreateTime})

    hostControlEntry.EntityData.YListKeys = []string {"HostControlIndex"}

    return &(hostControlEntry.EntityData)
}

// RMONMIB_HostTable
// A list of host entries.
type RMONMIB_HostTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular host that has been discovered
    // on an interface of this device.  For example, an instance of the
    // hostOutBroadcastPkts object might be named
    // hostOutBroadcastPkts.1.6.8.0.32.27.3.176. The type is slice of
    // RMONMIB_HostTable_HostEntry.
    HostEntry []*RMONMIB_HostTable_HostEntry
}

func (hostTable *RMONMIB_HostTable) GetEntityData() *types.CommonEntityData {
    hostTable.EntityData.YFilter = hostTable.YFilter
    hostTable.EntityData.YangName = "hostTable"
    hostTable.EntityData.BundleName = "cisco_ios_xe"
    hostTable.EntityData.ParentYangName = "RMON-MIB"
    hostTable.EntityData.SegmentPath = "hostTable"
    hostTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTable.EntityData.Children = types.NewOrderedMap()
    hostTable.EntityData.Children.Append("hostEntry", types.YChild{"HostEntry", nil})
    for i := range hostTable.HostEntry {
        hostTable.EntityData.Children.Append(types.GetSegmentPath(hostTable.HostEntry[i]), types.YChild{"HostEntry", hostTable.HostEntry[i]})
    }
    hostTable.EntityData.Leafs = types.NewOrderedMap()

    hostTable.EntityData.YListKeys = []string {}

    return &(hostTable.EntityData)
}

// RMONMIB_HostTable_HostEntry
// A collection of statistics for a particular host that has
// been discovered on an interface of this device.  For example,
// an instance of the hostOutBroadcastPkts object might be
// named hostOutBroadcastPkts.1.6.8.0.32.27.3.176
type RMONMIB_HostTable_HostEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected host statistics of which this
    // entry is a part.  The set of hosts identified by a particular value of this
    // index is associated with the hostControlEntry as identified by the same
    // value of hostControlIndex. The type is interface{} with range: 1..65535.
    HostIndex interface{}

    // This attribute is a key. The physical address of this host. The type is
    // string.
    HostAddress interface{}

    // An index that defines the relative ordering of the creation time of hosts
    // captured for a particular hostControlEntry.  This index shall be between 1
    // and N, where N is the value of the associated hostControlTableSize.  The
    // ordering of the indexes is based on the order of each entry's insertion
    // into the table, in which entries added earlier have a lower index value
    // than entries added later.  It is important to note that the order for a
    // particular entry may change as an (earlier) entry is deleted from the
    // table.  Because this order may change, management stations should make use
    // of the hostControlLastDeleteTime variable in the hostControlEntry
    // associated with the relevant portion of the hostTable.  By observing this
    // variable, the management station may detect the circumstances where a
    // previous association between a value of hostCreationOrder and a hostEntry
    // may no longer hold. The type is interface{} with range: 1..65535.
    HostCreationOrder interface{}

    // The number of good packets transmitted to this address since it was added
    // to the hostTable. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    HostInPkts interface{}

    // The number of packets, including bad packets, transmitted by this address
    // since it was added to the hostTable. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    HostOutPkts interface{}

    // The number of octets transmitted to this address since it was added to the
    // hostTable (excluding framing bits but including FCS octets), except for
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    HostInOctets interface{}

    // The number of octets transmitted by this address since it was added to the
    // hostTable (excluding framing bits but including FCS octets), including
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    HostOutOctets interface{}

    // The number of bad packets transmitted by this address since this host was
    // added to the hostTable. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    HostOutErrors interface{}

    // The number of good packets transmitted by this address that were directed
    // to the broadcast address since this host was added to the hostTable. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    HostOutBroadcastPkts interface{}

    // The number of good packets transmitted by this address that were directed
    // to a multicast address since this host was added to the hostTable. Note
    // that this number does not include packets directed to the broadcast
    // address. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    HostOutMulticastPkts interface{}
}

func (hostEntry *RMONMIB_HostTable_HostEntry) GetEntityData() *types.CommonEntityData {
    hostEntry.EntityData.YFilter = hostEntry.YFilter
    hostEntry.EntityData.YangName = "hostEntry"
    hostEntry.EntityData.BundleName = "cisco_ios_xe"
    hostEntry.EntityData.ParentYangName = "hostTable"
    hostEntry.EntityData.SegmentPath = "hostEntry" + types.AddKeyToken(hostEntry.HostIndex, "hostIndex") + types.AddKeyToken(hostEntry.HostAddress, "hostAddress")
    hostEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostEntry.EntityData.Children = types.NewOrderedMap()
    hostEntry.EntityData.Leafs = types.NewOrderedMap()
    hostEntry.EntityData.Leafs.Append("hostIndex", types.YLeaf{"HostIndex", hostEntry.HostIndex})
    hostEntry.EntityData.Leafs.Append("hostAddress", types.YLeaf{"HostAddress", hostEntry.HostAddress})
    hostEntry.EntityData.Leafs.Append("hostCreationOrder", types.YLeaf{"HostCreationOrder", hostEntry.HostCreationOrder})
    hostEntry.EntityData.Leafs.Append("hostInPkts", types.YLeaf{"HostInPkts", hostEntry.HostInPkts})
    hostEntry.EntityData.Leafs.Append("hostOutPkts", types.YLeaf{"HostOutPkts", hostEntry.HostOutPkts})
    hostEntry.EntityData.Leafs.Append("hostInOctets", types.YLeaf{"HostInOctets", hostEntry.HostInOctets})
    hostEntry.EntityData.Leafs.Append("hostOutOctets", types.YLeaf{"HostOutOctets", hostEntry.HostOutOctets})
    hostEntry.EntityData.Leafs.Append("hostOutErrors", types.YLeaf{"HostOutErrors", hostEntry.HostOutErrors})
    hostEntry.EntityData.Leafs.Append("hostOutBroadcastPkts", types.YLeaf{"HostOutBroadcastPkts", hostEntry.HostOutBroadcastPkts})
    hostEntry.EntityData.Leafs.Append("hostOutMulticastPkts", types.YLeaf{"HostOutMulticastPkts", hostEntry.HostOutMulticastPkts})

    hostEntry.EntityData.YListKeys = []string {"HostIndex", "HostAddress"}

    return &(hostEntry.EntityData)
}

// RMONMIB_HostTimeTable
// A list of time-ordered host table entries.
type RMONMIB_HostTimeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular host that has been discovered
    // on an interface of this device.  This collection includes the relative
    // ordering of the creation time of this object.  For example, an instance of
    // the hostTimeOutBroadcastPkts object might be named
    // hostTimeOutBroadcastPkts.1.687. The type is slice of
    // RMONMIB_HostTimeTable_HostTimeEntry.
    HostTimeEntry []*RMONMIB_HostTimeTable_HostTimeEntry
}

func (hostTimeTable *RMONMIB_HostTimeTable) GetEntityData() *types.CommonEntityData {
    hostTimeTable.EntityData.YFilter = hostTimeTable.YFilter
    hostTimeTable.EntityData.YangName = "hostTimeTable"
    hostTimeTable.EntityData.BundleName = "cisco_ios_xe"
    hostTimeTable.EntityData.ParentYangName = "RMON-MIB"
    hostTimeTable.EntityData.SegmentPath = "hostTimeTable"
    hostTimeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTimeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTimeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTimeTable.EntityData.Children = types.NewOrderedMap()
    hostTimeTable.EntityData.Children.Append("hostTimeEntry", types.YChild{"HostTimeEntry", nil})
    for i := range hostTimeTable.HostTimeEntry {
        hostTimeTable.EntityData.Children.Append(types.GetSegmentPath(hostTimeTable.HostTimeEntry[i]), types.YChild{"HostTimeEntry", hostTimeTable.HostTimeEntry[i]})
    }
    hostTimeTable.EntityData.Leafs = types.NewOrderedMap()

    hostTimeTable.EntityData.YListKeys = []string {}

    return &(hostTimeTable.EntityData)
}

// RMONMIB_HostTimeTable_HostTimeEntry
// A collection of statistics for a particular host that has
// been discovered on an interface of this device.  This
// collection includes the relative ordering of the creation
// time of this object.  For example, an instance of the
// hostTimeOutBroadcastPkts object might be named
// hostTimeOutBroadcastPkts.1.687
type RMONMIB_HostTimeTable_HostTimeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected host statistics of which this
    // entry is a part.  The set of hosts identified by a particular value of this
    // index is associated with the hostControlEntry as identified by the same
    // value of hostControlIndex. The type is interface{} with range: 1..65535.
    HostTimeIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostTime table among those entries associated with the same
    // hostControlEntry.  This index shall be between 1 and N, where N is the
    // value of the associated hostControlTableSize.  The ordering of the indexes
    // is based on the order of each entry's insertion into the table, in which
    // entries added earlier have a lower index value than entries added later.
    // Thus the management station has the ability to learn of new entries added
    // to this table without downloading the entire table.  It is important to
    // note that the index for a particular entry may change as an (earlier) entry
    // is deleted from the table.  Because this order may change, management
    // stations should make use of the hostControlLastDeleteTime variable in the
    // hostControlEntry associated with the relevant portion of the hostTimeTable.
    // By observing this variable, the management station may detect the
    // circumstances where a download of the table may have missed entries, and
    // where a previous association between a value of hostTimeCreationOrder and a
    // hostTimeEntry may no longer hold. The type is interface{} with range:
    // 1..65535.
    HostTimeCreationOrder interface{}

    // The physical address of this host. The type is string.
    HostTimeAddress interface{}

    // The number of good packets transmitted to this address since it was added
    // to the hostTimeTable. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    HostTimeInPkts interface{}

    // The number of packets, including bad packets, transmitted by this address
    // since it was added to the hostTimeTable. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    HostTimeOutPkts interface{}

    // The number of octets transmitted to this address since it was added to the
    // hostTimeTable (excluding framing bits but including FCS octets), except for
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    HostTimeInOctets interface{}

    // The number of octets transmitted by this address since it was added to the
    // hostTimeTable (excluding framing bits but including FCS octets), including
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    HostTimeOutOctets interface{}

    // The number of bad packets transmitted by this address since this host was
    // added to the hostTimeTable. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    HostTimeOutErrors interface{}

    // The number of good packets transmitted by this address that were directed
    // to the broadcast address since this host was added to the hostTimeTable.
    // The type is interface{} with range: 0..4294967295. Units are Packets.
    HostTimeOutBroadcastPkts interface{}

    // The number of good packets transmitted by this address that were directed
    // to a multicast address since this host was added to the hostTimeTable. Note
    // that this number does not include packets directed to the broadcast
    // address. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    HostTimeOutMulticastPkts interface{}
}

func (hostTimeEntry *RMONMIB_HostTimeTable_HostTimeEntry) GetEntityData() *types.CommonEntityData {
    hostTimeEntry.EntityData.YFilter = hostTimeEntry.YFilter
    hostTimeEntry.EntityData.YangName = "hostTimeEntry"
    hostTimeEntry.EntityData.BundleName = "cisco_ios_xe"
    hostTimeEntry.EntityData.ParentYangName = "hostTimeTable"
    hostTimeEntry.EntityData.SegmentPath = "hostTimeEntry" + types.AddKeyToken(hostTimeEntry.HostTimeIndex, "hostTimeIndex") + types.AddKeyToken(hostTimeEntry.HostTimeCreationOrder, "hostTimeCreationOrder")
    hostTimeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTimeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTimeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTimeEntry.EntityData.Children = types.NewOrderedMap()
    hostTimeEntry.EntityData.Leafs = types.NewOrderedMap()
    hostTimeEntry.EntityData.Leafs.Append("hostTimeIndex", types.YLeaf{"HostTimeIndex", hostTimeEntry.HostTimeIndex})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeCreationOrder", types.YLeaf{"HostTimeCreationOrder", hostTimeEntry.HostTimeCreationOrder})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeAddress", types.YLeaf{"HostTimeAddress", hostTimeEntry.HostTimeAddress})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeInPkts", types.YLeaf{"HostTimeInPkts", hostTimeEntry.HostTimeInPkts})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeOutPkts", types.YLeaf{"HostTimeOutPkts", hostTimeEntry.HostTimeOutPkts})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeInOctets", types.YLeaf{"HostTimeInOctets", hostTimeEntry.HostTimeInOctets})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeOutOctets", types.YLeaf{"HostTimeOutOctets", hostTimeEntry.HostTimeOutOctets})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeOutErrors", types.YLeaf{"HostTimeOutErrors", hostTimeEntry.HostTimeOutErrors})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeOutBroadcastPkts", types.YLeaf{"HostTimeOutBroadcastPkts", hostTimeEntry.HostTimeOutBroadcastPkts})
    hostTimeEntry.EntityData.Leafs.Append("hostTimeOutMulticastPkts", types.YLeaf{"HostTimeOutMulticastPkts", hostTimeEntry.HostTimeOutMulticastPkts})

    hostTimeEntry.EntityData.YListKeys = []string {"HostTimeIndex", "HostTimeCreationOrder"}

    return &(hostTimeEntry.EntityData)
}

// RMONMIB_HostTopNControlTable
// A list of top N host control entries.
type RMONMIB_HostTopNControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters that control the creation of a report of the top N
    // hosts according to several metrics.  For example, an instance of the
    // hostTopNDuration object might be named hostTopNDuration.3. The type is
    // slice of RMONMIB_HostTopNControlTable_HostTopNControlEntry.
    HostTopNControlEntry []*RMONMIB_HostTopNControlTable_HostTopNControlEntry
}

func (hostTopNControlTable *RMONMIB_HostTopNControlTable) GetEntityData() *types.CommonEntityData {
    hostTopNControlTable.EntityData.YFilter = hostTopNControlTable.YFilter
    hostTopNControlTable.EntityData.YangName = "hostTopNControlTable"
    hostTopNControlTable.EntityData.BundleName = "cisco_ios_xe"
    hostTopNControlTable.EntityData.ParentYangName = "RMON-MIB"
    hostTopNControlTable.EntityData.SegmentPath = "hostTopNControlTable"
    hostTopNControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTopNControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTopNControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTopNControlTable.EntityData.Children = types.NewOrderedMap()
    hostTopNControlTable.EntityData.Children.Append("hostTopNControlEntry", types.YChild{"HostTopNControlEntry", nil})
    for i := range hostTopNControlTable.HostTopNControlEntry {
        hostTopNControlTable.EntityData.Children.Append(types.GetSegmentPath(hostTopNControlTable.HostTopNControlEntry[i]), types.YChild{"HostTopNControlEntry", hostTopNControlTable.HostTopNControlEntry[i]})
    }
    hostTopNControlTable.EntityData.Leafs = types.NewOrderedMap()

    hostTopNControlTable.EntityData.YListKeys = []string {}

    return &(hostTopNControlTable.EntityData)
}

// RMONMIB_HostTopNControlTable_HostTopNControlEntry
// A set of parameters that control the creation of a report
// of the top N hosts according to several metrics.  For
// example, an instance of the hostTopNDuration object might
// be named hostTopNDuration.3
type RMONMIB_HostTopNControlTable_HostTopNControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostTopNControl table.  Each such entry defines one top N report prepared
    // for one interface. The type is interface{} with range: 1..65535.
    HostTopNControlIndex interface{}

    // The host table for which a top N report will be prepared on behalf of this
    // entry.  The host table identified by a particular value of this index is
    // associated with the same host table as identified by the same value of
    // hostIndex.  This object may not be modified if the associated
    // hostTopNStatus object is equal to valid(1). The type is interface{} with
    // range: 1..65535.
    HostTopNHostIndex interface{}

    // The variable for each host that the hostTopNRate variable is based upon. 
    // This object may not be modified if the associated hostTopNStatus object is
    // equal to valid(1). The type is HostTopNRateBase.
    HostTopNRateBase interface{}

    // The number of seconds left in the report currently being collected.  When
    // this object is modified by the management station, a new collection is
    // started, possibly aborting a currently running report.  The new value is
    // used as the requested duration of this report, which is loaded into the
    // associated hostTopNDuration object.  When this object is set to a non-zero
    // value, any associated hostTopNEntries shall be made inaccessible by the
    // monitor.  While the value of this object is non-zero, it decrements by one
    // per second until it reaches zero.  During this time, all associated
    // hostTopNEntries shall remain inaccessible.  At the time that this object
    // decrements to zero, the report is made accessible in the hostTopNTable. 
    // Thus, the hostTopN table needs to be created only at the end of the
    // collection interval. The type is interface{} with range:
    // -2147483648..2147483647. Units are Seconds.
    HostTopNTimeRemaining interface{}

    // The number of seconds that this report has collected during the last
    // sampling interval, or if this report is currently being collected, the
    // number of seconds that this report is being collected during this sampling
    // interval.  When the associated hostTopNTimeRemaining object is set, this
    // object shall be set by the probe to the same value and shall not be
    // modified until the next time the hostTopNTimeRemaining is set.  This value
    // shall be zero if no reports have been requested for this
    // hostTopNControlEntry. The type is interface{} with range:
    // -2147483648..2147483647. Units are Seconds.
    HostTopNDuration interface{}

    // The maximum number of hosts requested for the top N table.  When this
    // object is created or modified, the probe should set hostTopNGrantedSize as
    // closely to this object as is possible for the particular probe
    // implementation and available resources. The type is interface{} with range:
    // -2147483648..2147483647.
    HostTopNRequestedSize interface{}

    // The maximum number of hosts in the top N table.  When the associated
    // hostTopNRequestedSize object is created or modified, the probe should set
    // this object as closely to the requested value as is possible for the
    // particular implementation and available resources. The probe must not lower
    // this value except as a result of a set to the associated
    // hostTopNRequestedSize object.  Hosts with the highest value of hostTopNRate
    // shall be placed in this table in decreasing order of this rate until there
    // is no more room or until there are no more hosts. The type is interface{}
    // with range: -2147483648..2147483647.
    HostTopNGrantedSize interface{}

    // The value of sysUpTime when this top N report was last started.  In other
    // words, this is the time that the associated hostTopNTimeRemaining object
    // was modified to start the requested report. The type is interface{} with
    // range: 0..4294967295.
    HostTopNStartTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    HostTopNOwner interface{}

    // The status of this hostTopNControl entry.  If this object is not equal to
    // valid(1), all associated hostTopNEntries shall be deleted by the agent. The
    // type is EntryStatus.
    HostTopNStatus interface{}
}

func (hostTopNControlEntry *RMONMIB_HostTopNControlTable_HostTopNControlEntry) GetEntityData() *types.CommonEntityData {
    hostTopNControlEntry.EntityData.YFilter = hostTopNControlEntry.YFilter
    hostTopNControlEntry.EntityData.YangName = "hostTopNControlEntry"
    hostTopNControlEntry.EntityData.BundleName = "cisco_ios_xe"
    hostTopNControlEntry.EntityData.ParentYangName = "hostTopNControlTable"
    hostTopNControlEntry.EntityData.SegmentPath = "hostTopNControlEntry" + types.AddKeyToken(hostTopNControlEntry.HostTopNControlIndex, "hostTopNControlIndex")
    hostTopNControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTopNControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTopNControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTopNControlEntry.EntityData.Children = types.NewOrderedMap()
    hostTopNControlEntry.EntityData.Leafs = types.NewOrderedMap()
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNControlIndex", types.YLeaf{"HostTopNControlIndex", hostTopNControlEntry.HostTopNControlIndex})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNHostIndex", types.YLeaf{"HostTopNHostIndex", hostTopNControlEntry.HostTopNHostIndex})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNRateBase", types.YLeaf{"HostTopNRateBase", hostTopNControlEntry.HostTopNRateBase})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNTimeRemaining", types.YLeaf{"HostTopNTimeRemaining", hostTopNControlEntry.HostTopNTimeRemaining})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNDuration", types.YLeaf{"HostTopNDuration", hostTopNControlEntry.HostTopNDuration})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNRequestedSize", types.YLeaf{"HostTopNRequestedSize", hostTopNControlEntry.HostTopNRequestedSize})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNGrantedSize", types.YLeaf{"HostTopNGrantedSize", hostTopNControlEntry.HostTopNGrantedSize})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNStartTime", types.YLeaf{"HostTopNStartTime", hostTopNControlEntry.HostTopNStartTime})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNOwner", types.YLeaf{"HostTopNOwner", hostTopNControlEntry.HostTopNOwner})
    hostTopNControlEntry.EntityData.Leafs.Append("hostTopNStatus", types.YLeaf{"HostTopNStatus", hostTopNControlEntry.HostTopNStatus})

    hostTopNControlEntry.EntityData.YListKeys = []string {"HostTopNControlIndex"}

    return &(hostTopNControlEntry.EntityData)
}

// RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase represents hostTopNStatus object is equal to valid(1).
type RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase string

const (
    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNInPkts RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNInPkts"

    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNOutPkts RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNOutPkts"

    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNInOctets RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNInOctets"

    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNOutOctets RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNOutOctets"

    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNOutErrors RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNOutErrors"

    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNOutBroadcastPkts RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNOutBroadcastPkts"

    RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase_hostTopNOutMulticastPkts RMONMIB_HostTopNControlTable_HostTopNControlEntry_HostTopNRateBase = "hostTopNOutMulticastPkts"
)

// RMONMIB_HostTopNTable
// A list of top N host entries.
type RMONMIB_HostTopNTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of statistics for a host that is part of a top N report.  For
    // example, an instance of the hostTopNRate object might be named
    // hostTopNRate.3.10. The type is slice of
    // RMONMIB_HostTopNTable_HostTopNEntry.
    HostTopNEntry []*RMONMIB_HostTopNTable_HostTopNEntry
}

func (hostTopNTable *RMONMIB_HostTopNTable) GetEntityData() *types.CommonEntityData {
    hostTopNTable.EntityData.YFilter = hostTopNTable.YFilter
    hostTopNTable.EntityData.YangName = "hostTopNTable"
    hostTopNTable.EntityData.BundleName = "cisco_ios_xe"
    hostTopNTable.EntityData.ParentYangName = "RMON-MIB"
    hostTopNTable.EntityData.SegmentPath = "hostTopNTable"
    hostTopNTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTopNTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTopNTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTopNTable.EntityData.Children = types.NewOrderedMap()
    hostTopNTable.EntityData.Children.Append("hostTopNEntry", types.YChild{"HostTopNEntry", nil})
    for i := range hostTopNTable.HostTopNEntry {
        hostTopNTable.EntityData.Children.Append(types.GetSegmentPath(hostTopNTable.HostTopNEntry[i]), types.YChild{"HostTopNEntry", hostTopNTable.HostTopNEntry[i]})
    }
    hostTopNTable.EntityData.Leafs = types.NewOrderedMap()

    hostTopNTable.EntityData.YListKeys = []string {}

    return &(hostTopNTable.EntityData)
}

// RMONMIB_HostTopNTable_HostTopNEntry
// A set of statistics for a host that is part of a top N
// report.  For example, an instance of the hostTopNRate
// object might be named hostTopNRate.3.10
type RMONMIB_HostTopNTable_HostTopNEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object identifies the top N report of which
    // this entry is a part.  The set of hosts identified by a particular value of
    // this object is part of the same report as identified by the same value of
    // the hostTopNControlIndex object. The type is interface{} with range:
    // 1..65535.
    HostTopNReport interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostTopN table among those in the same report. This index is between 1 and
    // N, where N is the number of entries in this table.  Increasing values of
    // hostTopNIndex shall be assigned to entries with decreasing values of
    // hostTopNRate until index N is assigned to the entry with the lowest value
    // of hostTopNRate or there are no more hostTopNEntries. The type is
    // interface{} with range: 1..65535.
    HostTopNIndex interface{}

    // The physical address of this host. The type is string.
    HostTopNAddress interface{}

    // The amount of change in the selected variable during this sampling
    // interval.  The selected variable is this host's instance of the object
    // selected by hostTopNRateBase. The type is interface{} with range:
    // -2147483648..2147483647.
    HostTopNRate interface{}
}

func (hostTopNEntry *RMONMIB_HostTopNTable_HostTopNEntry) GetEntityData() *types.CommonEntityData {
    hostTopNEntry.EntityData.YFilter = hostTopNEntry.YFilter
    hostTopNEntry.EntityData.YangName = "hostTopNEntry"
    hostTopNEntry.EntityData.BundleName = "cisco_ios_xe"
    hostTopNEntry.EntityData.ParentYangName = "hostTopNTable"
    hostTopNEntry.EntityData.SegmentPath = "hostTopNEntry" + types.AddKeyToken(hostTopNEntry.HostTopNReport, "hostTopNReport") + types.AddKeyToken(hostTopNEntry.HostTopNIndex, "hostTopNIndex")
    hostTopNEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostTopNEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostTopNEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostTopNEntry.EntityData.Children = types.NewOrderedMap()
    hostTopNEntry.EntityData.Leafs = types.NewOrderedMap()
    hostTopNEntry.EntityData.Leafs.Append("hostTopNReport", types.YLeaf{"HostTopNReport", hostTopNEntry.HostTopNReport})
    hostTopNEntry.EntityData.Leafs.Append("hostTopNIndex", types.YLeaf{"HostTopNIndex", hostTopNEntry.HostTopNIndex})
    hostTopNEntry.EntityData.Leafs.Append("hostTopNAddress", types.YLeaf{"HostTopNAddress", hostTopNEntry.HostTopNAddress})
    hostTopNEntry.EntityData.Leafs.Append("hostTopNRate", types.YLeaf{"HostTopNRate", hostTopNEntry.HostTopNRate})

    hostTopNEntry.EntityData.YListKeys = []string {"HostTopNReport", "HostTopNIndex"}

    return &(hostTopNEntry.EntityData)
}

// RMONMIB_MatrixControlTable
// A list of information entries for the
// traffic matrix on each interface.
type RMONMIB_MatrixControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a traffic matrix on a particular interface.  For example,
    // an instance of the matrixControlLastDeleteTime object might be named
    // matrixControlLastDeleteTime.1. The type is slice of
    // RMONMIB_MatrixControlTable_MatrixControlEntry.
    MatrixControlEntry []*RMONMIB_MatrixControlTable_MatrixControlEntry
}

func (matrixControlTable *RMONMIB_MatrixControlTable) GetEntityData() *types.CommonEntityData {
    matrixControlTable.EntityData.YFilter = matrixControlTable.YFilter
    matrixControlTable.EntityData.YangName = "matrixControlTable"
    matrixControlTable.EntityData.BundleName = "cisco_ios_xe"
    matrixControlTable.EntityData.ParentYangName = "RMON-MIB"
    matrixControlTable.EntityData.SegmentPath = "matrixControlTable"
    matrixControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixControlTable.EntityData.Children = types.NewOrderedMap()
    matrixControlTable.EntityData.Children.Append("matrixControlEntry", types.YChild{"MatrixControlEntry", nil})
    for i := range matrixControlTable.MatrixControlEntry {
        matrixControlTable.EntityData.Children.Append(types.GetSegmentPath(matrixControlTable.MatrixControlEntry[i]), types.YChild{"MatrixControlEntry", matrixControlTable.MatrixControlEntry[i]})
    }
    matrixControlTable.EntityData.Leafs = types.NewOrderedMap()

    matrixControlTable.EntityData.YListKeys = []string {}

    return &(matrixControlTable.EntityData)
}

// RMONMIB_MatrixControlTable_MatrixControlEntry
// Information about a traffic matrix on a particular
// interface.  For example, an instance of the
// matrixControlLastDeleteTime object might be named
// matrixControlLastDeleteTime.1
type RMONMIB_MatrixControlTable_MatrixControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // matrixControl table.  Each such entry defines a function that discovers
    // conversations on a particular interface and places statistics about them in
    // the matrixSDTable and the matrixDSTable on behalf of this
    // matrixControlEntry. The type is interface{} with range: 1..65535.
    MatrixControlIndex interface{}

    // This object identifies the source of the data from which this entry creates
    // a traffic matrix. This source can be any interface on this device.  In
    // order to identify a particular interface, this object shall identify the
    // instance of the ifIndex object, defined in RFC 2233 [17], for the desired
    // interface.  For example, if an entry were to receive data from interface
    // #1, this object would be set to ifIndex.1.  The statistics in this group
    // reflect all packets on the local network segment attached to the identified
    // interface.  An agent may or may not be able to tell if fundamental changes
    // to the media of the interface have occurred and necessitate an invalidation
    // of this entry.  For example, a hot-pluggable ethernet card could be pulled
    // out and replaced by a token-ring card.  In such a case, if the agent has
    // such knowledge of the change, it is recommended that it invalidate this
    // entry.  This object may not be modified if the associated
    // matrixControlStatus object is equal to valid(1). The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    MatrixControlDataSource interface{}

    // The number of matrixSDEntries in the matrixSDTable for this interface. 
    // This must also be the value of the number of entries in the matrixDSTable
    // for this interface. The type is interface{} with range:
    // -2147483648..2147483647.
    MatrixControlTableSize interface{}

    // The value of sysUpTime when the last entry was deleted from the portion of
    // the matrixSDTable or matrixDSTable associated with this matrixControlEntry.
    // If no deletions have occurred, this value shall be zero. The type is
    // interface{} with range: 0..4294967295.
    MatrixControlLastDeleteTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    MatrixControlOwner interface{}

    // The status of this matrixControl entry. If this object is not equal to
    // valid(1), all associated entries in the matrixSDTable and the matrixDSTable
    // shall be deleted by the agent. The type is EntryStatus.
    MatrixControlStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted      because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    MatrixControlDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    MatrixControlCreateTime interface{}
}

func (matrixControlEntry *RMONMIB_MatrixControlTable_MatrixControlEntry) GetEntityData() *types.CommonEntityData {
    matrixControlEntry.EntityData.YFilter = matrixControlEntry.YFilter
    matrixControlEntry.EntityData.YangName = "matrixControlEntry"
    matrixControlEntry.EntityData.BundleName = "cisco_ios_xe"
    matrixControlEntry.EntityData.ParentYangName = "matrixControlTable"
    matrixControlEntry.EntityData.SegmentPath = "matrixControlEntry" + types.AddKeyToken(matrixControlEntry.MatrixControlIndex, "matrixControlIndex")
    matrixControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixControlEntry.EntityData.Children = types.NewOrderedMap()
    matrixControlEntry.EntityData.Leafs = types.NewOrderedMap()
    matrixControlEntry.EntityData.Leafs.Append("matrixControlIndex", types.YLeaf{"MatrixControlIndex", matrixControlEntry.MatrixControlIndex})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlDataSource", types.YLeaf{"MatrixControlDataSource", matrixControlEntry.MatrixControlDataSource})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlTableSize", types.YLeaf{"MatrixControlTableSize", matrixControlEntry.MatrixControlTableSize})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlLastDeleteTime", types.YLeaf{"MatrixControlLastDeleteTime", matrixControlEntry.MatrixControlLastDeleteTime})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlOwner", types.YLeaf{"MatrixControlOwner", matrixControlEntry.MatrixControlOwner})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlStatus", types.YLeaf{"MatrixControlStatus", matrixControlEntry.MatrixControlStatus})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlDroppedFrames", types.YLeaf{"MatrixControlDroppedFrames", matrixControlEntry.MatrixControlDroppedFrames})
    matrixControlEntry.EntityData.Leafs.Append("matrixControlCreateTime", types.YLeaf{"MatrixControlCreateTime", matrixControlEntry.MatrixControlCreateTime})

    matrixControlEntry.EntityData.YListKeys = []string {"MatrixControlIndex"}

    return &(matrixControlEntry.EntityData)
}

// RMONMIB_MatrixSDTable
// A list of traffic matrix entries indexed by
// source and destination MAC address.
type RMONMIB_MatrixSDTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for communications between two addresses on a
    // particular interface.  For example, an instance of the matrixSDPkts object
    // might be named matrixSDPkts.1.6.8.0.32.27.3.176.6.8.0.32.10.8.113. The type
    // is slice of RMONMIB_MatrixSDTable_MatrixSDEntry.
    MatrixSDEntry []*RMONMIB_MatrixSDTable_MatrixSDEntry
}

func (matrixSDTable *RMONMIB_MatrixSDTable) GetEntityData() *types.CommonEntityData {
    matrixSDTable.EntityData.YFilter = matrixSDTable.YFilter
    matrixSDTable.EntityData.YangName = "matrixSDTable"
    matrixSDTable.EntityData.BundleName = "cisco_ios_xe"
    matrixSDTable.EntityData.ParentYangName = "RMON-MIB"
    matrixSDTable.EntityData.SegmentPath = "matrixSDTable"
    matrixSDTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixSDTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixSDTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixSDTable.EntityData.Children = types.NewOrderedMap()
    matrixSDTable.EntityData.Children.Append("matrixSDEntry", types.YChild{"MatrixSDEntry", nil})
    for i := range matrixSDTable.MatrixSDEntry {
        matrixSDTable.EntityData.Children.Append(types.GetSegmentPath(matrixSDTable.MatrixSDEntry[i]), types.YChild{"MatrixSDEntry", matrixSDTable.MatrixSDEntry[i]})
    }
    matrixSDTable.EntityData.Leafs = types.NewOrderedMap()

    matrixSDTable.EntityData.YListKeys = []string {}

    return &(matrixSDTable.EntityData)
}

// RMONMIB_MatrixSDTable_MatrixSDEntry
// A collection of statistics for communications between
// two addresses on a particular interface.  For example,
// an instance of the matrixSDPkts object might be named
// matrixSDPkts.1.6.8.0.32.27.3.176.6.8.0.32.10.8.113
type RMONMIB_MatrixSDTable_MatrixSDEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected matrix statistics of which
    // this entry is a part.  The set of matrix statistics identified by a
    // particular value of this index is associated with the same
    // matrixControlEntry as identified by the same value of matrixControlIndex.
    // The type is interface{} with range: 1..65535.
    MatrixSDIndex interface{}

    // This attribute is a key. The source physical address. The type is string.
    MatrixSDSourceAddress interface{}

    // This attribute is a key. The destination physical address. The type is
    // string.
    MatrixSDDestAddress interface{}

    // The number of packets transmitted from the source address to the
    // destination address (this number includes bad packets). The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    MatrixSDPkts interface{}

    // The number of octets (excluding framing bits but including FCS octets)
    // contained in all packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    MatrixSDOctets interface{}

    // The number of bad packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    MatrixSDErrors interface{}
}

func (matrixSDEntry *RMONMIB_MatrixSDTable_MatrixSDEntry) GetEntityData() *types.CommonEntityData {
    matrixSDEntry.EntityData.YFilter = matrixSDEntry.YFilter
    matrixSDEntry.EntityData.YangName = "matrixSDEntry"
    matrixSDEntry.EntityData.BundleName = "cisco_ios_xe"
    matrixSDEntry.EntityData.ParentYangName = "matrixSDTable"
    matrixSDEntry.EntityData.SegmentPath = "matrixSDEntry" + types.AddKeyToken(matrixSDEntry.MatrixSDIndex, "matrixSDIndex") + types.AddKeyToken(matrixSDEntry.MatrixSDSourceAddress, "matrixSDSourceAddress") + types.AddKeyToken(matrixSDEntry.MatrixSDDestAddress, "matrixSDDestAddress")
    matrixSDEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixSDEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixSDEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixSDEntry.EntityData.Children = types.NewOrderedMap()
    matrixSDEntry.EntityData.Leafs = types.NewOrderedMap()
    matrixSDEntry.EntityData.Leafs.Append("matrixSDIndex", types.YLeaf{"MatrixSDIndex", matrixSDEntry.MatrixSDIndex})
    matrixSDEntry.EntityData.Leafs.Append("matrixSDSourceAddress", types.YLeaf{"MatrixSDSourceAddress", matrixSDEntry.MatrixSDSourceAddress})
    matrixSDEntry.EntityData.Leafs.Append("matrixSDDestAddress", types.YLeaf{"MatrixSDDestAddress", matrixSDEntry.MatrixSDDestAddress})
    matrixSDEntry.EntityData.Leafs.Append("matrixSDPkts", types.YLeaf{"MatrixSDPkts", matrixSDEntry.MatrixSDPkts})
    matrixSDEntry.EntityData.Leafs.Append("matrixSDOctets", types.YLeaf{"MatrixSDOctets", matrixSDEntry.MatrixSDOctets})
    matrixSDEntry.EntityData.Leafs.Append("matrixSDErrors", types.YLeaf{"MatrixSDErrors", matrixSDEntry.MatrixSDErrors})

    matrixSDEntry.EntityData.YListKeys = []string {"MatrixSDIndex", "MatrixSDSourceAddress", "MatrixSDDestAddress"}

    return &(matrixSDEntry.EntityData)
}

// RMONMIB_MatrixDSTable
// A list of traffic matrix entries indexed by
// destination and source MAC address.
type RMONMIB_MatrixDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for communications between two addresses on a
    // particular interface.  For example, an instance of the matrixSDPkts object
    // might be named matrixSDPkts.1.6.8.0.32.10.8.113.6.8.0.32.27.3.176. The type
    // is slice of RMONMIB_MatrixDSTable_MatrixDSEntry.
    MatrixDSEntry []*RMONMIB_MatrixDSTable_MatrixDSEntry
}

func (matrixDSTable *RMONMIB_MatrixDSTable) GetEntityData() *types.CommonEntityData {
    matrixDSTable.EntityData.YFilter = matrixDSTable.YFilter
    matrixDSTable.EntityData.YangName = "matrixDSTable"
    matrixDSTable.EntityData.BundleName = "cisco_ios_xe"
    matrixDSTable.EntityData.ParentYangName = "RMON-MIB"
    matrixDSTable.EntityData.SegmentPath = "matrixDSTable"
    matrixDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixDSTable.EntityData.Children = types.NewOrderedMap()
    matrixDSTable.EntityData.Children.Append("matrixDSEntry", types.YChild{"MatrixDSEntry", nil})
    for i := range matrixDSTable.MatrixDSEntry {
        matrixDSTable.EntityData.Children.Append(types.GetSegmentPath(matrixDSTable.MatrixDSEntry[i]), types.YChild{"MatrixDSEntry", matrixDSTable.MatrixDSEntry[i]})
    }
    matrixDSTable.EntityData.Leafs = types.NewOrderedMap()

    matrixDSTable.EntityData.YListKeys = []string {}

    return &(matrixDSTable.EntityData)
}

// RMONMIB_MatrixDSTable_MatrixDSEntry
// A collection of statistics for communications between
// two addresses on a particular interface.  For example,
// an instance of the matrixSDPkts object might be named
// matrixSDPkts.1.6.8.0.32.10.8.113.6.8.0.32.27.3.176
type RMONMIB_MatrixDSTable_MatrixDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected matrix statistics of which
    // this entry is a part.  The set of matrix statistics identified by a
    // particular value of this index is associated with the same
    // matrixControlEntry as identified by the same value of matrixControlIndex.
    // The type is interface{} with range: 1..65535.
    MatrixDSIndex interface{}

    // This attribute is a key. The destination physical address. The type is
    // string.
    MatrixDSDestAddress interface{}

    // This attribute is a key. The source physical address. The type is string.
    MatrixDSSourceAddress interface{}

    // The number of packets transmitted from the source address to the
    // destination address (this number includes bad packets). The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    MatrixDSPkts interface{}

    // The number of octets (excluding framing bits but including FCS octets)
    // contained in all packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    MatrixDSOctets interface{}

    // The number of bad packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    MatrixDSErrors interface{}
}

func (matrixDSEntry *RMONMIB_MatrixDSTable_MatrixDSEntry) GetEntityData() *types.CommonEntityData {
    matrixDSEntry.EntityData.YFilter = matrixDSEntry.YFilter
    matrixDSEntry.EntityData.YangName = "matrixDSEntry"
    matrixDSEntry.EntityData.BundleName = "cisco_ios_xe"
    matrixDSEntry.EntityData.ParentYangName = "matrixDSTable"
    matrixDSEntry.EntityData.SegmentPath = "matrixDSEntry" + types.AddKeyToken(matrixDSEntry.MatrixDSIndex, "matrixDSIndex") + types.AddKeyToken(matrixDSEntry.MatrixDSDestAddress, "matrixDSDestAddress") + types.AddKeyToken(matrixDSEntry.MatrixDSSourceAddress, "matrixDSSourceAddress")
    matrixDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixDSEntry.EntityData.Children = types.NewOrderedMap()
    matrixDSEntry.EntityData.Leafs = types.NewOrderedMap()
    matrixDSEntry.EntityData.Leafs.Append("matrixDSIndex", types.YLeaf{"MatrixDSIndex", matrixDSEntry.MatrixDSIndex})
    matrixDSEntry.EntityData.Leafs.Append("matrixDSDestAddress", types.YLeaf{"MatrixDSDestAddress", matrixDSEntry.MatrixDSDestAddress})
    matrixDSEntry.EntityData.Leafs.Append("matrixDSSourceAddress", types.YLeaf{"MatrixDSSourceAddress", matrixDSEntry.MatrixDSSourceAddress})
    matrixDSEntry.EntityData.Leafs.Append("matrixDSPkts", types.YLeaf{"MatrixDSPkts", matrixDSEntry.MatrixDSPkts})
    matrixDSEntry.EntityData.Leafs.Append("matrixDSOctets", types.YLeaf{"MatrixDSOctets", matrixDSEntry.MatrixDSOctets})
    matrixDSEntry.EntityData.Leafs.Append("matrixDSErrors", types.YLeaf{"MatrixDSErrors", matrixDSEntry.MatrixDSErrors})

    matrixDSEntry.EntityData.YListKeys = []string {"MatrixDSIndex", "MatrixDSDestAddress", "MatrixDSSourceAddress"}

    return &(matrixDSEntry.EntityData)
}

// RMONMIB_FilterTable
// A list of packet filter entries.
type RMONMIB_FilterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters for a packet filter applied on a particular interface. 
    // As an example, an instance of the filterPktData object might be named
    // filterPktData.12. The type is slice of RMONMIB_FilterTable_FilterEntry.
    FilterEntry []*RMONMIB_FilterTable_FilterEntry
}

func (filterTable *RMONMIB_FilterTable) GetEntityData() *types.CommonEntityData {
    filterTable.EntityData.YFilter = filterTable.YFilter
    filterTable.EntityData.YangName = "filterTable"
    filterTable.EntityData.BundleName = "cisco_ios_xe"
    filterTable.EntityData.ParentYangName = "RMON-MIB"
    filterTable.EntityData.SegmentPath = "filterTable"
    filterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    filterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    filterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    filterTable.EntityData.Children = types.NewOrderedMap()
    filterTable.EntityData.Children.Append("filterEntry", types.YChild{"FilterEntry", nil})
    for i := range filterTable.FilterEntry {
        filterTable.EntityData.Children.Append(types.GetSegmentPath(filterTable.FilterEntry[i]), types.YChild{"FilterEntry", filterTable.FilterEntry[i]})
    }
    filterTable.EntityData.Leafs = types.NewOrderedMap()

    filterTable.EntityData.YListKeys = []string {}

    return &(filterTable.EntityData)
}

// RMONMIB_FilterTable_FilterEntry
// A set of parameters for a packet filter applied on a
// particular interface.  As an example, an instance of the
// filterPktData object might be named filterPktData.12
type RMONMIB_FilterTable_FilterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // filter table.  Each such entry defines one filter that is to be applied to
    // every packet received on an interface. The type is interface{} with range:
    // 1..65535.
    FilterIndex interface{}

    // This object identifies the channel of which this filter is a part.  The
    // filters identified by a particular value of this object are associated with
    // the same channel as identified by the same value of the channelIndex
    // object. The type is interface{} with range: 1..65535.
    FilterChannelIndex interface{}

    // The offset from the beginning of each packet where a match of packet data
    // will be attempted.  This offset is measured from the point in the physical
    // layer packet after the framing bits, if any.  For example, in an Ethernet
    // frame, this point is at the beginning of the destination MAC address.  This
    // object may not be modified if the associated filterStatus object is equal
    // to valid(1). The type is interface{} with range: -2147483648..2147483647.
    // Units are Octets.
    FilterPktDataOffset interface{}

    // The data that is to be matched with the input packet. For each packet
    // received, this filter and the accompanying filterPktDataMask and
    // filterPktDataNotMask will be adjusted for the offset.  The only bits
    // relevant to this match algorithm are those that have the corresponding
    // filterPktDataMask bit equal to one.  The following three rules are then
    // applied to every packet:  (1) If the packet is too short and does not have
    // data     corresponding to part of the filterPktData, the packet     will
    // fail this data match.  (2) For each relevant bit from the packet with the  
    // corresponding filterPktDataNotMask bit set to zero, if     the bit from the
    // packet is not equal to the corresponding     bit from the filterPktData,
    // then the packet will fail     this data match.  (3) If for every relevant
    // bit from the packet with the     corresponding filterPktDataNotMask bit set
    // to one, the     bit from the packet is equal to the corresponding bit    
    // from the filterPktData, then the packet will fail this     data match.  Any
    // packets that have not failed any of the three matches above have passed
    // this data match.  In particular, a zero length filter will match any
    // packet.  This object may not be modified if the associated filterStatus
    // object is equal to valid(1). The type is string.
    FilterPktData interface{}

    // The mask that is applied to the match process. After adjusting this mask
    // for the offset, only those bits in the received packet that correspond to
    // bits set in this mask are relevant for further processing by the match
    // algorithm.  The offset is applied to filterPktDataMask in the same way it
    // is applied to the filter.  For the purposes of the matching algorithm, if
    // the associated filterPktData object is longer than this mask, this mask is
    // conceptually extended with '1' bits until it reaches the length of the
    // filterPktData object.  This object may not be modified if the associated
    // filterStatus object is equal to valid(1). The type is string.
    FilterPktDataMask interface{}

    // The inversion mask that is applied to the match process.  After adjusting
    // this mask for the offset, those relevant bits in the received packet that
    // correspond to bits cleared in this mask must all be equal to their
    // corresponding bits in the filterPktData object for the packet to be
    // accepted.  In addition, at least one of those relevant bits in the received
    // packet that correspond to bits set in this mask must be different to its
    // corresponding bit in the filterPktData object.  For the purposes of the
    // matching algorithm, if the associated filterPktData object is longer than
    // this mask, this mask is conceptually extended with '0' bits until it
    // reaches the length of the filterPktData object.  This object may not be
    // modified if the associated filterStatus object is equal to valid(1). The
    // type is string.
    FilterPktDataNotMask interface{}

    // The status that is to be matched with the input packet. The only bits
    // relevant to this match algorithm are those that have the corresponding
    // filterPktStatusMask bit equal to one. The following two rules are then
    // applied to every packet:  (1) For each relevant bit from the packet status
    // with the     corresponding filterPktStatusNotMask bit set to zero, if    
    // the bit from the packet status is not equal to the     corresponding bit
    // from the filterPktStatus, then the     packet will fail this status match. 
    // (2) If for every relevant bit from the packet status with the    
    // corresponding filterPktStatusNotMask bit set to one, the     bit from the
    // packet status is equal to the corresponding     bit from the
    // filterPktStatus, then the packet will fail     this status match.  Any
    // packets that have not failed either of the two matches above have passed
    // this status match.  In particular, a zero length status filter will match
    // any packet's status.  The value of the packet status is a sum.  This sum
    // initially takes the value zero.  Then, for each error, E, that has been
    // discovered in this packet, 2 raised to a value representing E is added to
    // the sum. The errors and the bits that represent them are dependent on the
    // media type of the interface that this channel is receiving packets from. 
    // The errors defined for a packet captured off of an Ethernet interface are
    // as follows:      bit #    Error         0    Packet is longer than 1518
    // octets         1    Packet is shorter than 64 octets         2    Packet
    // experienced a CRC or Alignment error  For example, an Ethernet fragment
    // would have a value of 6 (2^1 + 2^2).  As this MIB is expanded to new media
    // types, this object will have other media-specific errors defined.  For the
    // purposes of this status matching algorithm, if the packet status is longer
    // than this filterPktStatus object, this object is conceptually extended with
    // '0' bits until it reaches the size of the packet status.  This object may
    // not be modified if the associated filterStatus object is equal to valid(1).
    // The type is interface{} with range: -2147483648..2147483647.
    FilterPktStatus interface{}

    // The mask that is applied to the status match process. Only those bits in
    // the received packet that correspond to bits set in this mask are relevant
    // for further processing by the status match algorithm.  For the purposes of
    // the matching algorithm, if the associated filterPktStatus object is longer
    // than this mask, this mask is conceptually extended with '1' bits until it
    // reaches the size of the filterPktStatus.  In addition, if a packet status
    // is longer than this mask, this mask is conceptually extended with '0' bits
    // until it reaches the size of the packet status.  This object may not be
    // modified if the associated filterStatus object is equal to valid(1). The
    // type is interface{} with range: -2147483648..2147483647.
    FilterPktStatusMask interface{}

    // The inversion mask that is applied to the status match process.  Those
    // relevant bits in the received packet status that correspond to bits cleared
    // in this mask must all be equal to their corresponding bits in the
    // filterPktStatus object for the packet to be accepted.  In addition, at
    // least one of those relevant bits in the received packet status that
    // correspond to bits set in this mask must be different to its corresponding
    // bit in the filterPktStatus object for the packet to be accepted.  For the
    // purposes of the matching algorithm, if the associated filterPktStatus
    // object or a packet status is longer than this mask, this mask is
    // conceptually extended with '0' bits until it reaches the longer of the
    // lengths of the filterPktStatus object and the packet status.  This object
    // may not be modified if the associated filterStatus object is equal to
    // valid(1). The type is interface{} with range: -2147483648..2147483647.
    FilterPktStatusNotMask interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    FilterOwner interface{}

    // The status of this filter entry. The type is EntryStatus.
    FilterStatus interface{}

    // When this object is set to a non-zero value, the filter that it is
    // associated with performs the following operations on every packet:  1) - If
    // the packet doesn't match the protocol directory entry      identified by
    // this object, discard the packet and exit      (i.e., discard the packet if
    // it is not of the identified      protocol). 2) - If the associated
    // filterProtocolDirLocalIndex is non-zero      and the packet doesn't match
    // the protocol directory      entry identified by that object, discard the
    // packet and      exit 3) - If the packet matches, perform the regular filter
    // algorithm as if the beginning of this named protocol is      the beginning
    // of the packet, potentially applying the      filterOffset value to move
    // further into the packet. The type is interface{} with range: 0..2147483647.
    FilterProtocolDirDataLocalIndex interface{}

    // When this object is set to a non-zero value, the filter that it is
    // associated with will discard the packet if the packet doesn't match this
    // protocol directory entry. The type is interface{} with range:
    // 0..2147483647.
    FilterProtocolDirLocalIndex interface{}
}

func (filterEntry *RMONMIB_FilterTable_FilterEntry) GetEntityData() *types.CommonEntityData {
    filterEntry.EntityData.YFilter = filterEntry.YFilter
    filterEntry.EntityData.YangName = "filterEntry"
    filterEntry.EntityData.BundleName = "cisco_ios_xe"
    filterEntry.EntityData.ParentYangName = "filterTable"
    filterEntry.EntityData.SegmentPath = "filterEntry" + types.AddKeyToken(filterEntry.FilterIndex, "filterIndex")
    filterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    filterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    filterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    filterEntry.EntityData.Children = types.NewOrderedMap()
    filterEntry.EntityData.Leafs = types.NewOrderedMap()
    filterEntry.EntityData.Leafs.Append("filterIndex", types.YLeaf{"FilterIndex", filterEntry.FilterIndex})
    filterEntry.EntityData.Leafs.Append("filterChannelIndex", types.YLeaf{"FilterChannelIndex", filterEntry.FilterChannelIndex})
    filterEntry.EntityData.Leafs.Append("filterPktDataOffset", types.YLeaf{"FilterPktDataOffset", filterEntry.FilterPktDataOffset})
    filterEntry.EntityData.Leafs.Append("filterPktData", types.YLeaf{"FilterPktData", filterEntry.FilterPktData})
    filterEntry.EntityData.Leafs.Append("filterPktDataMask", types.YLeaf{"FilterPktDataMask", filterEntry.FilterPktDataMask})
    filterEntry.EntityData.Leafs.Append("filterPktDataNotMask", types.YLeaf{"FilterPktDataNotMask", filterEntry.FilterPktDataNotMask})
    filterEntry.EntityData.Leafs.Append("filterPktStatus", types.YLeaf{"FilterPktStatus", filterEntry.FilterPktStatus})
    filterEntry.EntityData.Leafs.Append("filterPktStatusMask", types.YLeaf{"FilterPktStatusMask", filterEntry.FilterPktStatusMask})
    filterEntry.EntityData.Leafs.Append("filterPktStatusNotMask", types.YLeaf{"FilterPktStatusNotMask", filterEntry.FilterPktStatusNotMask})
    filterEntry.EntityData.Leafs.Append("filterOwner", types.YLeaf{"FilterOwner", filterEntry.FilterOwner})
    filterEntry.EntityData.Leafs.Append("filterStatus", types.YLeaf{"FilterStatus", filterEntry.FilterStatus})
    filterEntry.EntityData.Leafs.Append("filterProtocolDirDataLocalIndex", types.YLeaf{"FilterProtocolDirDataLocalIndex", filterEntry.FilterProtocolDirDataLocalIndex})
    filterEntry.EntityData.Leafs.Append("filterProtocolDirLocalIndex", types.YLeaf{"FilterProtocolDirLocalIndex", filterEntry.FilterProtocolDirLocalIndex})

    filterEntry.EntityData.YListKeys = []string {"FilterIndex"}

    return &(filterEntry.EntityData)
}

// RMONMIB_ChannelTable
// A list of packet channel entries.
type RMONMIB_ChannelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters for a packet channel applied on a particular interface.
    // As an example, an instance of the channelMatches object might be named
    // channelMatches.3. The type is slice of RMONMIB_ChannelTable_ChannelEntry.
    ChannelEntry []*RMONMIB_ChannelTable_ChannelEntry
}

func (channelTable *RMONMIB_ChannelTable) GetEntityData() *types.CommonEntityData {
    channelTable.EntityData.YFilter = channelTable.YFilter
    channelTable.EntityData.YangName = "channelTable"
    channelTable.EntityData.BundleName = "cisco_ios_xe"
    channelTable.EntityData.ParentYangName = "RMON-MIB"
    channelTable.EntityData.SegmentPath = "channelTable"
    channelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    channelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    channelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    channelTable.EntityData.Children = types.NewOrderedMap()
    channelTable.EntityData.Children.Append("channelEntry", types.YChild{"ChannelEntry", nil})
    for i := range channelTable.ChannelEntry {
        channelTable.EntityData.Children.Append(types.GetSegmentPath(channelTable.ChannelEntry[i]), types.YChild{"ChannelEntry", channelTable.ChannelEntry[i]})
    }
    channelTable.EntityData.Leafs = types.NewOrderedMap()

    channelTable.EntityData.YListKeys = []string {}

    return &(channelTable.EntityData)
}

// RMONMIB_ChannelTable_ChannelEntry
// A set of parameters for a packet channel applied on a
// particular interface.  As an example, an instance of the
// channelMatches object might be named channelMatches.3
type RMONMIB_ChannelTable_ChannelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // channel table.  Each such entry defines one channel, a logical data and
    // event stream.  It is suggested that before creating a channel, an
    // application should scan all instances of the filterChannelIndex object to
    // make sure that there are no pre-existing filters that would be
    // inadvertently be linked to the channel. The type is interface{} with range:
    // 1..65535.
    ChannelIndex interface{}

    // The value of this object uniquely identifies the interface on this remote
    // network monitoring device to which the associated filters are applied to
    // allow data into this channel.  The interface identified by a particular
    // value of this object is the same interface as identified by the same value
    // of the ifIndex object, defined in RFC 2233 [17].  The filters in this group
    // are applied to all packets on the local network segment attached to the
    // identified interface.  An agent may or may not be able to tell if
    // fundamental changes to the media of the interface have occurred and
    // necessitate an invalidation of this entry.  For example, a hot-pluggable
    // ethernet card could be pulled out and replaced by a token-ring card.  In
    // such a case, if the agent has such knowledge of the change, it is
    // recommended that it invalidate this entry.  This object may not be modified
    // if the associated channelStatus object is equal to valid(1). The type is
    // interface{} with range: 1..65535.
    ChannelIfIndex interface{}

    // This object controls the action of the filters associated with this
    // channel.  If this object is equal to acceptMatched(1), packets will be
    // accepted to this channel if they are accepted by both the packet data and
    // packet status matches of an associated filter.  If this object is equal to
    // acceptFailed(2), packets will be accepted to this channel only if they fail
    // either the packet data match or the packet status match of each of the
    // associated filters.  In particular, a channel with no associated filters
    // will match no packets if set to acceptMatched(1) case and will match all
    // packets in the acceptFailed(2) case.  This object may not be modified if
    // the associated channelStatus object is equal to valid(1). The type is
    // ChannelAcceptType.
    ChannelAcceptType interface{}

    // This object controls the flow of data through this channel. If this object
    // is on(1), data, status and events flow through this channel.  If this
    // object is off(2), data, status and events will not flow through this
    // channel. The type is ChannelDataControl.
    ChannelDataControl interface{}

    // The value of this object identifies the event that is configured to turn
    // the associated channelDataControl from off to on when the event is
    // generated.  The event identified by a particular value of this object is
    // the same event as identified by the same value of the eventIndex object. 
    // If there is no corresponding entry in the eventTable, then no association
    // exists.  In fact, if no event is intended for this channel,
    // channelTurnOnEventIndex must be set to zero, a non-existent event index.
    // This object may not be modified if the associated channelStatus object is
    // equal to valid(1). The type is interface{} with range: 0..65535.
    ChannelTurnOnEventIndex interface{}

    // The value of this object identifies the event that is configured to turn
    // the associated channelDataControl from on to off when the event is
    // generated.  The event identified by a particular value of this object is
    // the same event as identified by the same value of the eventIndex object. 
    // If there is no corresponding entry in the eventTable, then no association
    // exists.  In fact, if no event is intended for this channel,
    // channelTurnOffEventIndex must be set to zero, a non-existent event index. 
    // This object may not be modified if the associated channelStatus object is
    // equal to valid(1). The type is interface{} with range: 0..65535.
    ChannelTurnOffEventIndex interface{}

    // The value of this object identifies the event that is configured to be
    // generated when the associated channelDataControl is on and a packet is
    // matched.  The event identified by a particular value of this object is the
    // same event as identified by the same value of the eventIndex object.  If
    // there is no corresponding entry in the eventTable, then no association
    // exists.  In fact, if no event is intended for this channel,
    // channelEventIndex must be set to zero, a non-existent event index.  This
    // object may not be modified if the associated channelStatus object is equal
    // to valid(1). The type is interface{} with range: 0..65535.
    ChannelEventIndex interface{}

    // The event status of this channel.  If this channel is configured to
    // generate events when packets are matched, a means of controlling the flow
    // of those events is often needed.  When this object is equal to
    // eventReady(1), a single event may be generated, after which this object
    // will be set by the probe to eventFired(2).  While in the eventFired(2)
    // state, no events will be generated until the object is modified to
    // eventReady(1) (or eventAlwaysReady(3)).  The management station can thus
    // easily respond to a notification of an event by re-enabling this object. 
    // If the management station wishes to disable this flow control and allow
    // events to be generated at will, this object may be set to
    // eventAlwaysReady(3).  Disabling the flow control is discouraged as it can
    // result in high network traffic or other performance problems. The type is
    // ChannelEventStatus.
    ChannelEventStatus interface{}

    // The number of times this channel has matched a packet. Note that this
    // object is updated even when channelDataControl is set to off. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    ChannelMatches interface{}

    // A comment describing this channel. The type is string with length: 0..127.
    ChannelDescription interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    ChannelOwner interface{}

    // The status of this channel entry. The type is EntryStatus.
    ChannelStatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe      is out of some resources and decides to shed
    // load from this collection.  This count does not include packets that were
    // not counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    ChannelDroppedFrames interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    ChannelCreateTime interface{}
}

func (channelEntry *RMONMIB_ChannelTable_ChannelEntry) GetEntityData() *types.CommonEntityData {
    channelEntry.EntityData.YFilter = channelEntry.YFilter
    channelEntry.EntityData.YangName = "channelEntry"
    channelEntry.EntityData.BundleName = "cisco_ios_xe"
    channelEntry.EntityData.ParentYangName = "channelTable"
    channelEntry.EntityData.SegmentPath = "channelEntry" + types.AddKeyToken(channelEntry.ChannelIndex, "channelIndex")
    channelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    channelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    channelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    channelEntry.EntityData.Children = types.NewOrderedMap()
    channelEntry.EntityData.Leafs = types.NewOrderedMap()
    channelEntry.EntityData.Leafs.Append("channelIndex", types.YLeaf{"ChannelIndex", channelEntry.ChannelIndex})
    channelEntry.EntityData.Leafs.Append("channelIfIndex", types.YLeaf{"ChannelIfIndex", channelEntry.ChannelIfIndex})
    channelEntry.EntityData.Leafs.Append("channelAcceptType", types.YLeaf{"ChannelAcceptType", channelEntry.ChannelAcceptType})
    channelEntry.EntityData.Leafs.Append("channelDataControl", types.YLeaf{"ChannelDataControl", channelEntry.ChannelDataControl})
    channelEntry.EntityData.Leafs.Append("channelTurnOnEventIndex", types.YLeaf{"ChannelTurnOnEventIndex", channelEntry.ChannelTurnOnEventIndex})
    channelEntry.EntityData.Leafs.Append("channelTurnOffEventIndex", types.YLeaf{"ChannelTurnOffEventIndex", channelEntry.ChannelTurnOffEventIndex})
    channelEntry.EntityData.Leafs.Append("channelEventIndex", types.YLeaf{"ChannelEventIndex", channelEntry.ChannelEventIndex})
    channelEntry.EntityData.Leafs.Append("channelEventStatus", types.YLeaf{"ChannelEventStatus", channelEntry.ChannelEventStatus})
    channelEntry.EntityData.Leafs.Append("channelMatches", types.YLeaf{"ChannelMatches", channelEntry.ChannelMatches})
    channelEntry.EntityData.Leafs.Append("channelDescription", types.YLeaf{"ChannelDescription", channelEntry.ChannelDescription})
    channelEntry.EntityData.Leafs.Append("channelOwner", types.YLeaf{"ChannelOwner", channelEntry.ChannelOwner})
    channelEntry.EntityData.Leafs.Append("channelStatus", types.YLeaf{"ChannelStatus", channelEntry.ChannelStatus})
    channelEntry.EntityData.Leafs.Append("channelDroppedFrames", types.YLeaf{"ChannelDroppedFrames", channelEntry.ChannelDroppedFrames})
    channelEntry.EntityData.Leafs.Append("channelCreateTime", types.YLeaf{"ChannelCreateTime", channelEntry.ChannelCreateTime})

    channelEntry.EntityData.YListKeys = []string {"ChannelIndex"}

    return &(channelEntry.EntityData)
}

// RMONMIB_ChannelTable_ChannelEntry_ChannelAcceptType represents channelStatus object is equal to valid(1).
type RMONMIB_ChannelTable_ChannelEntry_ChannelAcceptType string

const (
    RMONMIB_ChannelTable_ChannelEntry_ChannelAcceptType_acceptMatched RMONMIB_ChannelTable_ChannelEntry_ChannelAcceptType = "acceptMatched"

    RMONMIB_ChannelTable_ChannelEntry_ChannelAcceptType_acceptFailed RMONMIB_ChannelTable_ChannelEntry_ChannelAcceptType = "acceptFailed"
)

// RMONMIB_ChannelTable_ChannelEntry_ChannelDataControl represents status and events will not flow through this channel.
type RMONMIB_ChannelTable_ChannelEntry_ChannelDataControl string

const (
    RMONMIB_ChannelTable_ChannelEntry_ChannelDataControl_on RMONMIB_ChannelTable_ChannelEntry_ChannelDataControl = "on"

    RMONMIB_ChannelTable_ChannelEntry_ChannelDataControl_off RMONMIB_ChannelTable_ChannelEntry_ChannelDataControl = "off"
)

// RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus represents traffic or other performance problems.
type RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus string

const (
    RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus_eventReady RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus = "eventReady"

    RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus_eventFired RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus = "eventFired"

    RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus_eventAlwaysReady RMONMIB_ChannelTable_ChannelEntry_ChannelEventStatus = "eventAlwaysReady"
)

// RMONMIB_BufferControlTable
// A list of buffers control entries.
type RMONMIB_BufferControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters that control the collection of a stream of packets that
    // have matched filters.  As an example, an instance of the
    // bufferControlCaptureSliceSize object might be named
    // bufferControlCaptureSliceSize.3. The type is slice of
    // RMONMIB_BufferControlTable_BufferControlEntry.
    BufferControlEntry []*RMONMIB_BufferControlTable_BufferControlEntry
}

func (bufferControlTable *RMONMIB_BufferControlTable) GetEntityData() *types.CommonEntityData {
    bufferControlTable.EntityData.YFilter = bufferControlTable.YFilter
    bufferControlTable.EntityData.YangName = "bufferControlTable"
    bufferControlTable.EntityData.BundleName = "cisco_ios_xe"
    bufferControlTable.EntityData.ParentYangName = "RMON-MIB"
    bufferControlTable.EntityData.SegmentPath = "bufferControlTable"
    bufferControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bufferControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bufferControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bufferControlTable.EntityData.Children = types.NewOrderedMap()
    bufferControlTable.EntityData.Children.Append("bufferControlEntry", types.YChild{"BufferControlEntry", nil})
    for i := range bufferControlTable.BufferControlEntry {
        bufferControlTable.EntityData.Children.Append(types.GetSegmentPath(bufferControlTable.BufferControlEntry[i]), types.YChild{"BufferControlEntry", bufferControlTable.BufferControlEntry[i]})
    }
    bufferControlTable.EntityData.Leafs = types.NewOrderedMap()

    bufferControlTable.EntityData.YListKeys = []string {}

    return &(bufferControlTable.EntityData)
}

// RMONMIB_BufferControlTable_BufferControlEntry
// A set of parameters that control the collection of a stream
// of packets that have matched filters.  As an example, an
// instance of the bufferControlCaptureSliceSize object might
// be named bufferControlCaptureSliceSize.3
type RMONMIB_BufferControlTable_BufferControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // bufferControl table.  The value of this index shall never be zero.  Each
    // such entry defines one set of packets that is captured and controlled by
    // one or more filters. The type is interface{} with range: 1..65535.
    BufferControlIndex interface{}

    // An index that identifies the channel that is the source of packets for this
    // bufferControl table. The channel identified by a particular value of this
    // index is the same as identified by the same value of the channelIndex
    // object.  This object may not be modified if the associated
    // bufferControlStatus object is equal to valid(1). The type is interface{}
    // with range: 1..65535.
    BufferControlChannelIndex interface{}

    // This object shows whether the buffer has room to accept new packets or if
    // it is full.  If the status is spaceAvailable(1), the buffer is accepting
    // new packets normally.  If the status is full(2) and the associated
    // bufferControlFullAction object is wrapWhenFull, the buffer is accepting new
    // packets by deleting enough of the oldest packets to make room for new ones
    // as they arrive.  Otherwise, if the status is full(2) and the
    // bufferControlFullAction object is lockWhenFull, then the buffer has stopped
    // collecting packets.  When this object is set to full(2) the probe must not
    // later set it to spaceAvailable(1) except in the case of a significant gain
    // in resources such as an increase of bufferControlOctetsGranted.  In
    // particular, the wrap-mode action of deleting old packets to make room for
    // newly arrived packets must not affect the value of this object. The type is
    // BufferControlFullStatus.
    BufferControlFullStatus interface{}

    // Controls the action of the buffer when it reaches the full status.  When in
    // the lockWhenFull(1) state and a packet is added to the buffer that fills
    // the buffer, the bufferControlFullStatus will be set to full(2) and this
    // buffer will stop capturing packets. The type is BufferControlFullAction.
    BufferControlFullAction interface{}

    // The maximum number of octets of each packet that will be saved in this
    // capture buffer. For example, if a 1500 octet packet is received by the
    // probe and this object is set to 500, then only 500 octets of the packet
    // will be stored in the associated capture buffer.  If this variable is set
    // to 0, the capture buffer will save as many octets as is possible.  This
    // object may not be modified if the associated bufferControlStatus object is
    // equal to valid(1). The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    BufferControlCaptureSliceSize interface{}

    // The maximum number of octets of each packet in this capture buffer that
    // will be returned in an SNMP retrieval of that packet.  For example, if 500
    // octets of a packet have been stored in the associated capture buffer, the
    // associated bufferControlDownloadOffset is 0, and this object is set to 100,
    // then the captureBufferPacket object that contains the packet will contain
    // only the first 100 octets of the packet.  A prudent manager will take into
    // account possible interoperability or fragmentation problems that may occur
    // if the download slice size is set too large. In particular, conformant SNMP
    // implementations are not required to accept messages whose length exceeds
    // 484 octets, although they are encouraged to support larger datagrams
    // whenever feasible. The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    BufferControlDownloadSliceSize interface{}

    // The offset of the first octet of each packet in this capture buffer that
    // will be returned in an SNMP retrieval of that packet.  For example, if 500
    // octets of a packet have been stored in the associated capture buffer and
    // this object is set to 100, then the captureBufferPacket object that
    // contains the packet will contain bytes starting 100 octets into the packet.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // Octets.
    BufferControlDownloadOffset interface{}

    // The requested maximum number of octets to be saved in this captureBuffer,
    // including any implementation-specific overhead. If this variable is set to
    // -1, the capture buffer will save as many octets as is possible.  When this
    // object is created or modified, the probe should set
    // bufferControlMaxOctetsGranted as closely to this object as is possible for
    // the particular probe implementation and available resources.  However, if
    // the object has the special value of -1, the probe must set
    // bufferControlMaxOctetsGranted to -1. The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    BufferControlMaxOctetsRequested interface{}

    // The maximum number of octets that can be saved in this captureBuffer,
    // including overhead. If this variable is -1, the capture buffer will save as
    // many octets as possible.  When the bufferControlMaxOctetsRequested object
    // is created or modified, the probe should set this object as closely to the
    // requested value as is possible for the particular probe implementation and
    // available resources. However, if the request object has the special value
    // of -1, the probe must set this object to -1.  The probe must not lower this
    // value except as a result of a modification to the associated
    // bufferControlMaxOctetsRequested object.  When this maximum number of octets
    // is reached and a new packet is to be added to this capture buffer and the
    // corresponding bufferControlFullAction is set to wrapWhenFull(2), enough of
    // the oldest packets associated with this capture buffer shall be deleted by
    // the agent so that the new packet can be added.  If the corresponding
    // bufferControlFullAction is set to lockWhenFull(1), the new packet shall be
    // discarded.  In either case, the probe must set bufferControlFullStatus to
    // full(2).  When the value of this object changes to a value less than the
    // current value, entries are deleted from the captureBufferTable associated
    // with this bufferControlEntry.  Enough of the oldest of these
    // captureBufferEntries shall be deleted by the agent so that the number of
    // octets used remains less than or equal to the new value of this object. 
    // When the value of this object changes to a value greater than the current
    // value, the number of associated captureBufferEntries may be allowed to
    // grow. The type is interface{} with range: -2147483648..2147483647. Units
    // are Octets.
    BufferControlMaxOctetsGranted interface{}

    // The number of packets currently in this captureBuffer. The type is
    // interface{} with range: -2147483648..2147483647. Units are Packets.
    BufferControlCapturedPackets interface{}

    // The value of sysUpTime when this capture buffer was first turned on. The
    // type is interface{} with range: 0..4294967295.
    BufferControlTurnOnTime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    BufferControlOwner interface{}

    // The status of this buffer Control Entry. The type is EntryStatus.
    BufferControlStatus interface{}
}

func (bufferControlEntry *RMONMIB_BufferControlTable_BufferControlEntry) GetEntityData() *types.CommonEntityData {
    bufferControlEntry.EntityData.YFilter = bufferControlEntry.YFilter
    bufferControlEntry.EntityData.YangName = "bufferControlEntry"
    bufferControlEntry.EntityData.BundleName = "cisco_ios_xe"
    bufferControlEntry.EntityData.ParentYangName = "bufferControlTable"
    bufferControlEntry.EntityData.SegmentPath = "bufferControlEntry" + types.AddKeyToken(bufferControlEntry.BufferControlIndex, "bufferControlIndex")
    bufferControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bufferControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bufferControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bufferControlEntry.EntityData.Children = types.NewOrderedMap()
    bufferControlEntry.EntityData.Leafs = types.NewOrderedMap()
    bufferControlEntry.EntityData.Leafs.Append("bufferControlIndex", types.YLeaf{"BufferControlIndex", bufferControlEntry.BufferControlIndex})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlChannelIndex", types.YLeaf{"BufferControlChannelIndex", bufferControlEntry.BufferControlChannelIndex})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlFullStatus", types.YLeaf{"BufferControlFullStatus", bufferControlEntry.BufferControlFullStatus})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlFullAction", types.YLeaf{"BufferControlFullAction", bufferControlEntry.BufferControlFullAction})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlCaptureSliceSize", types.YLeaf{"BufferControlCaptureSliceSize", bufferControlEntry.BufferControlCaptureSliceSize})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlDownloadSliceSize", types.YLeaf{"BufferControlDownloadSliceSize", bufferControlEntry.BufferControlDownloadSliceSize})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlDownloadOffset", types.YLeaf{"BufferControlDownloadOffset", bufferControlEntry.BufferControlDownloadOffset})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlMaxOctetsRequested", types.YLeaf{"BufferControlMaxOctetsRequested", bufferControlEntry.BufferControlMaxOctetsRequested})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlMaxOctetsGranted", types.YLeaf{"BufferControlMaxOctetsGranted", bufferControlEntry.BufferControlMaxOctetsGranted})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlCapturedPackets", types.YLeaf{"BufferControlCapturedPackets", bufferControlEntry.BufferControlCapturedPackets})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlTurnOnTime", types.YLeaf{"BufferControlTurnOnTime", bufferControlEntry.BufferControlTurnOnTime})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlOwner", types.YLeaf{"BufferControlOwner", bufferControlEntry.BufferControlOwner})
    bufferControlEntry.EntityData.Leafs.Append("bufferControlStatus", types.YLeaf{"BufferControlStatus", bufferControlEntry.BufferControlStatus})

    bufferControlEntry.EntityData.YListKeys = []string {"BufferControlIndex"}

    return &(bufferControlEntry.EntityData)
}

// RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullAction represents packets.
type RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullAction string

const (
    RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullAction_lockWhenFull RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullAction = "lockWhenFull"

    RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullAction_wrapWhenFull RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullAction = "wrapWhenFull"
)

// RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullStatus represents must not affect the value of this object.
type RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullStatus string

const (
    RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullStatus_spaceAvailable RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullStatus = "spaceAvailable"

    RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullStatus_full RMONMIB_BufferControlTable_BufferControlEntry_BufferControlFullStatus = "full"
)

// RMONMIB_CaptureBufferTable
// A list of packets captured off of a channel.
type RMONMIB_CaptureBufferTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A packet captured off of an attached network.  As an example, an instance
    // of the captureBufferPacketData object might be named
    // captureBufferPacketData.3.1783. The type is slice of
    // RMONMIB_CaptureBufferTable_CaptureBufferEntry.
    CaptureBufferEntry []*RMONMIB_CaptureBufferTable_CaptureBufferEntry
}

func (captureBufferTable *RMONMIB_CaptureBufferTable) GetEntityData() *types.CommonEntityData {
    captureBufferTable.EntityData.YFilter = captureBufferTable.YFilter
    captureBufferTable.EntityData.YangName = "captureBufferTable"
    captureBufferTable.EntityData.BundleName = "cisco_ios_xe"
    captureBufferTable.EntityData.ParentYangName = "RMON-MIB"
    captureBufferTable.EntityData.SegmentPath = "captureBufferTable"
    captureBufferTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    captureBufferTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    captureBufferTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    captureBufferTable.EntityData.Children = types.NewOrderedMap()
    captureBufferTable.EntityData.Children.Append("captureBufferEntry", types.YChild{"CaptureBufferEntry", nil})
    for i := range captureBufferTable.CaptureBufferEntry {
        captureBufferTable.EntityData.Children.Append(types.GetSegmentPath(captureBufferTable.CaptureBufferEntry[i]), types.YChild{"CaptureBufferEntry", captureBufferTable.CaptureBufferEntry[i]})
    }
    captureBufferTable.EntityData.Leafs = types.NewOrderedMap()

    captureBufferTable.EntityData.YListKeys = []string {}

    return &(captureBufferTable.EntityData)
}

// RMONMIB_CaptureBufferTable_CaptureBufferEntry
// A packet captured off of an attached network.  As an
// example, an instance of the captureBufferPacketData
// object might be named captureBufferPacketData.3.1783
type RMONMIB_CaptureBufferTable_CaptureBufferEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the bufferControlEntry with which
    // this packet is associated. The type is interface{} with range: 1..65535.
    CaptureBufferControlIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // captureBuffer table associated with a particular bufferControlEntry.  This
    // index will start at 1 and increase by one for each new packet added with
    // the same captureBufferControlIndex.  Should this value reach 2147483647,
    // the next packet added with the same captureBufferControlIndex shall cause
    // this value to wrap around to 1. The type is interface{} with range:
    // 1..2147483647.
    CaptureBufferIndex interface{}

    // An index that describes the order of packets that are received on a
    // particular interface. The packetID of a packet captured on an interface is
    // defined to be greater than the packetID's of all packets captured
    // previously on the same interface.  As the captureBufferPacketID object has
    // a maximum positive value of 2^31 - 1, any captureBufferPacketID object
    // shall have the value of the associated packet's packetID mod 2^31. The type
    // is interface{} with range: -2147483648..2147483647.
    CaptureBufferPacketID interface{}

    // The data inside the packet, starting at the beginning of the packet plus
    // any offset specified in the associated bufferControlDownloadOffset,
    // including any link level headers.  The length of the data in this object is
    // the minimum of the length of the captured packet minus the offset, the
    // length of the associated bufferControlCaptureSliceSize minus the offset,
    // and the associated bufferControlDownloadSliceSize.  If this minimum is less
    // than zero, this object shall have a length of zero. The type is string.
    CaptureBufferPacketData interface{}

    // The actual length (off the wire) of the packet stored in this entry,
    // including FCS octets. The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    CaptureBufferPacketLength interface{}

    // The number of milliseconds that had passed since this capture buffer was
    // first turned on when this packet was captured. The type is interface{} with
    // range: -2147483648..2147483647. Units are Milliseconds.
    CaptureBufferPacketTime interface{}

    // A value which indicates the error status of this packet.  The value of this
    // object is defined in the same way as filterPktStatus.  The value is a sum. 
    // This sum initially takes the value zero.  Then, for each error, E, that has
    // been discovered in this packet, 2 raised to a value representing E is added
    // to the sum.  The errors defined for a packet captured off of an Ethernet
    // interface are as follows:      bit #    Error         0    Packet is longer
    // than 1518 octets         1    Packet is shorter than 64 octets         2   
    // Packet experienced a CRC or Alignment error         3    First packet in
    // this capture buffer after              it was detected that some packets
    // were              not processed correctly.         4    Packet's order in
    // buffer is only approximate              (May only be set for packets sent
    // from              the probe)  For example, an Ethernet fragment would have
    // a value of 6 (2^1 + 2^2).  As this MIB is expanded to new media types, this
    // object will have other media-specific errors defined. The type is
    // interface{} with range: -2147483648..2147483647.
    CaptureBufferPacketStatus interface{}
}

func (captureBufferEntry *RMONMIB_CaptureBufferTable_CaptureBufferEntry) GetEntityData() *types.CommonEntityData {
    captureBufferEntry.EntityData.YFilter = captureBufferEntry.YFilter
    captureBufferEntry.EntityData.YangName = "captureBufferEntry"
    captureBufferEntry.EntityData.BundleName = "cisco_ios_xe"
    captureBufferEntry.EntityData.ParentYangName = "captureBufferTable"
    captureBufferEntry.EntityData.SegmentPath = "captureBufferEntry" + types.AddKeyToken(captureBufferEntry.CaptureBufferControlIndex, "captureBufferControlIndex") + types.AddKeyToken(captureBufferEntry.CaptureBufferIndex, "captureBufferIndex")
    captureBufferEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    captureBufferEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    captureBufferEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    captureBufferEntry.EntityData.Children = types.NewOrderedMap()
    captureBufferEntry.EntityData.Leafs = types.NewOrderedMap()
    captureBufferEntry.EntityData.Leafs.Append("captureBufferControlIndex", types.YLeaf{"CaptureBufferControlIndex", captureBufferEntry.CaptureBufferControlIndex})
    captureBufferEntry.EntityData.Leafs.Append("captureBufferIndex", types.YLeaf{"CaptureBufferIndex", captureBufferEntry.CaptureBufferIndex})
    captureBufferEntry.EntityData.Leafs.Append("captureBufferPacketID", types.YLeaf{"CaptureBufferPacketID", captureBufferEntry.CaptureBufferPacketID})
    captureBufferEntry.EntityData.Leafs.Append("captureBufferPacketData", types.YLeaf{"CaptureBufferPacketData", captureBufferEntry.CaptureBufferPacketData})
    captureBufferEntry.EntityData.Leafs.Append("captureBufferPacketLength", types.YLeaf{"CaptureBufferPacketLength", captureBufferEntry.CaptureBufferPacketLength})
    captureBufferEntry.EntityData.Leafs.Append("captureBufferPacketTime", types.YLeaf{"CaptureBufferPacketTime", captureBufferEntry.CaptureBufferPacketTime})
    captureBufferEntry.EntityData.Leafs.Append("captureBufferPacketStatus", types.YLeaf{"CaptureBufferPacketStatus", captureBufferEntry.CaptureBufferPacketStatus})

    captureBufferEntry.EntityData.YListKeys = []string {"CaptureBufferControlIndex", "CaptureBufferIndex"}

    return &(captureBufferEntry.EntityData)
}

// RMONMIB_EventTable
// A list of events to be generated.
type RMONMIB_EventTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters that describe an event to be generated when certain
    // conditions are met.  As an example, an instance of the eventLastTimeSent
    // object might be named eventLastTimeSent.6. The type is slice of
    // RMONMIB_EventTable_EventEntry.
    EventEntry []*RMONMIB_EventTable_EventEntry
}

func (eventTable *RMONMIB_EventTable) GetEntityData() *types.CommonEntityData {
    eventTable.EntityData.YFilter = eventTable.YFilter
    eventTable.EntityData.YangName = "eventTable"
    eventTable.EntityData.BundleName = "cisco_ios_xe"
    eventTable.EntityData.ParentYangName = "RMON-MIB"
    eventTable.EntityData.SegmentPath = "eventTable"
    eventTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eventTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eventTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eventTable.EntityData.Children = types.NewOrderedMap()
    eventTable.EntityData.Children.Append("eventEntry", types.YChild{"EventEntry", nil})
    for i := range eventTable.EventEntry {
        eventTable.EntityData.Children.Append(types.GetSegmentPath(eventTable.EventEntry[i]), types.YChild{"EventEntry", eventTable.EventEntry[i]})
    }
    eventTable.EntityData.Leafs = types.NewOrderedMap()

    eventTable.EntityData.YListKeys = []string {}

    return &(eventTable.EntityData)
}

// RMONMIB_EventTable_EventEntry
// A set of parameters that describe an event to be generated
// when certain conditions are met.  As an example, an instance
// of the eventLastTimeSent object might be named
// eventLastTimeSent.6
type RMONMIB_EventTable_EventEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // event table.  Each such entry defines one event that is to be generated
    // when the appropriate conditions occur. The type is interface{} with range:
    // 1..65535.
    EventIndex interface{}

    // A comment describing this event entry. The type is string with length:
    // 0..127.
    EventDescription interface{}

    // The type of notification that the probe will make about this event.  In the
    // case of log, an entry is made in the log table for each event.  In the case
    // of snmp-trap, an SNMP trap is sent to one or more management stations. The
    // type is EventType.
    EventType interface{}

    // If an SNMP trap is to be sent, it will be sent to the SNMP community
    // specified by this octet string. The type is string with length: 0..127.
    EventCommunity interface{}

    // The value of sysUpTime at the time this event entry last generated an
    // event.  If this entry has not generated any events, this value will be
    // zero. The type is interface{} with range: 0..4294967295.
    EventLastTimeSent interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it.  If this object contains a string starting with 'monitor'
    // and has associated entries in the log table, all connected management
    // stations should retrieve those log entries, as they may have significance
    // to all management stations connected to this device. The type is string
    // with length: 0..127.
    EventOwner interface{}

    // The status of this event entry.  If this object is not equal to valid(1),
    // all associated log entries shall be deleted by the agent. The type is
    // EntryStatus.
    EventStatus interface{}
}

func (eventEntry *RMONMIB_EventTable_EventEntry) GetEntityData() *types.CommonEntityData {
    eventEntry.EntityData.YFilter = eventEntry.YFilter
    eventEntry.EntityData.YangName = "eventEntry"
    eventEntry.EntityData.BundleName = "cisco_ios_xe"
    eventEntry.EntityData.ParentYangName = "eventTable"
    eventEntry.EntityData.SegmentPath = "eventEntry" + types.AddKeyToken(eventEntry.EventIndex, "eventIndex")
    eventEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eventEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eventEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eventEntry.EntityData.Children = types.NewOrderedMap()
    eventEntry.EntityData.Leafs = types.NewOrderedMap()
    eventEntry.EntityData.Leafs.Append("eventIndex", types.YLeaf{"EventIndex", eventEntry.EventIndex})
    eventEntry.EntityData.Leafs.Append("eventDescription", types.YLeaf{"EventDescription", eventEntry.EventDescription})
    eventEntry.EntityData.Leafs.Append("eventType", types.YLeaf{"EventType", eventEntry.EventType})
    eventEntry.EntityData.Leafs.Append("eventCommunity", types.YLeaf{"EventCommunity", eventEntry.EventCommunity})
    eventEntry.EntityData.Leafs.Append("eventLastTimeSent", types.YLeaf{"EventLastTimeSent", eventEntry.EventLastTimeSent})
    eventEntry.EntityData.Leafs.Append("eventOwner", types.YLeaf{"EventOwner", eventEntry.EventOwner})
    eventEntry.EntityData.Leafs.Append("eventStatus", types.YLeaf{"EventStatus", eventEntry.EventStatus})

    eventEntry.EntityData.YListKeys = []string {"EventIndex"}

    return &(eventEntry.EntityData)
}

// RMONMIB_EventTable_EventEntry_EventType represents management stations.
type RMONMIB_EventTable_EventEntry_EventType string

const (
    RMONMIB_EventTable_EventEntry_EventType_none RMONMIB_EventTable_EventEntry_EventType = "none"

    RMONMIB_EventTable_EventEntry_EventType_log RMONMIB_EventTable_EventEntry_EventType = "log"

    RMONMIB_EventTable_EventEntry_EventType_snmptrap RMONMIB_EventTable_EventEntry_EventType = "snmptrap"

    RMONMIB_EventTable_EventEntry_EventType_logandtrap RMONMIB_EventTable_EventEntry_EventType = "logandtrap"
)

// RMONMIB_LogTable
// A list of events that have been logged.
type RMONMIB_LogTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of data describing an event that has been logged.  For example, an
    // instance of the logDescription object might be named logDescription.6.47.
    // The type is slice of RMONMIB_LogTable_LogEntry.
    LogEntry []*RMONMIB_LogTable_LogEntry
}

func (logTable *RMONMIB_LogTable) GetEntityData() *types.CommonEntityData {
    logTable.EntityData.YFilter = logTable.YFilter
    logTable.EntityData.YangName = "logTable"
    logTable.EntityData.BundleName = "cisco_ios_xe"
    logTable.EntityData.ParentYangName = "RMON-MIB"
    logTable.EntityData.SegmentPath = "logTable"
    logTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    logTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    logTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    logTable.EntityData.Children = types.NewOrderedMap()
    logTable.EntityData.Children.Append("logEntry", types.YChild{"LogEntry", nil})
    for i := range logTable.LogEntry {
        logTable.EntityData.Children.Append(types.GetSegmentPath(logTable.LogEntry[i]), types.YChild{"LogEntry", logTable.LogEntry[i]})
    }
    logTable.EntityData.Leafs = types.NewOrderedMap()

    logTable.EntityData.YListKeys = []string {}

    return &(logTable.EntityData)
}

// RMONMIB_LogTable_LogEntry
// A set of data describing an event that has been
// logged.  For example, an instance of the logDescription
// object might be named logDescription.6.47
type RMONMIB_LogTable_LogEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The event entry that generated this log entry. 
    // The log identified by a particular value of this index is associated with
    // the same eventEntry as identified by the same value of eventIndex. The type
    // is interface{} with range: 1..65535.
    LogEventIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // log table amongst those generated by the same eventEntries.  These indexes
    // are assigned beginning with 1 and increase by one with each new log entry. 
    // The association between values of logIndex and logEntries is fixed for the
    // lifetime of each logEntry. The agent may choose to delete the oldest
    // instances of logEntry as required because of lack of memory.  It is an
    // implementation-specific matter as to when this deletion may occur. The type
    // is interface{} with range: 1..2147483647.
    LogIndex interface{}

    // The value of sysUpTime when this log entry was created. The type is
    // interface{} with range: 0..4294967295.
    LogTime interface{}

    // An implementation dependent description of the event that activated this
    // log entry. The type is string with length: 0..255.
    LogDescription interface{}
}

func (logEntry *RMONMIB_LogTable_LogEntry) GetEntityData() *types.CommonEntityData {
    logEntry.EntityData.YFilter = logEntry.YFilter
    logEntry.EntityData.YangName = "logEntry"
    logEntry.EntityData.BundleName = "cisco_ios_xe"
    logEntry.EntityData.ParentYangName = "logTable"
    logEntry.EntityData.SegmentPath = "logEntry" + types.AddKeyToken(logEntry.LogEventIndex, "logEventIndex") + types.AddKeyToken(logEntry.LogIndex, "logIndex")
    logEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    logEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    logEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    logEntry.EntityData.Children = types.NewOrderedMap()
    logEntry.EntityData.Leafs = types.NewOrderedMap()
    logEntry.EntityData.Leafs.Append("logEventIndex", types.YLeaf{"LogEventIndex", logEntry.LogEventIndex})
    logEntry.EntityData.Leafs.Append("logIndex", types.YLeaf{"LogIndex", logEntry.LogIndex})
    logEntry.EntityData.Leafs.Append("logTime", types.YLeaf{"LogTime", logEntry.LogTime})
    logEntry.EntityData.Leafs.Append("logDescription", types.YLeaf{"LogDescription", logEntry.LogDescription})

    logEntry.EntityData.YListKeys = []string {"LogEventIndex", "LogIndex"}

    return &(logEntry.EntityData)
}

