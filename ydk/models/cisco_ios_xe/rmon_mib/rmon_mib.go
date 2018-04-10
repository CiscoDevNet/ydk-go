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

type Rmoneventsv2 struct {
}

func (id Rmoneventsv2) String() string {
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
    Etherstatstable RMONMIB_Etherstatstable

    // A list of history control entries.
    Historycontroltable RMONMIB_Historycontroltable

    // A list of Ethernet history entries.
    Etherhistorytable RMONMIB_Etherhistorytable

    // A list of alarm entries.
    Alarmtable RMONMIB_Alarmtable

    // A list of host table control entries.
    Hostcontroltable RMONMIB_Hostcontroltable

    // A list of host entries.
    Hosttable RMONMIB_Hosttable

    // A list of time-ordered host table entries.
    Hosttimetable RMONMIB_Hosttimetable

    // A list of top N host control entries.
    Hosttopncontroltable RMONMIB_Hosttopncontroltable

    // A list of top N host entries.
    Hosttopntable RMONMIB_Hosttopntable

    // A list of information entries for the traffic matrix on each interface.
    Matrixcontroltable RMONMIB_Matrixcontroltable

    // A list of traffic matrix entries indexed by source and destination MAC
    // address.
    Matrixsdtable RMONMIB_Matrixsdtable

    // A list of traffic matrix entries indexed by destination and source MAC
    // address.
    Matrixdstable RMONMIB_Matrixdstable

    // A list of packet filter entries.
    Filtertable RMONMIB_Filtertable

    // A list of packet channel entries.
    Channeltable RMONMIB_Channeltable

    // A list of buffers control entries.
    Buffercontroltable RMONMIB_Buffercontroltable

    // A list of packets captured off of a channel.
    Capturebuffertable RMONMIB_Capturebuffertable

    // A list of events to be generated.
    Eventtable RMONMIB_Eventtable

    // A list of events that have been logged.
    Logtable RMONMIB_Logtable
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

    rMONMIB.EntityData.Children = make(map[string]types.YChild)
    rMONMIB.EntityData.Children["etherStatsTable"] = types.YChild{"Etherstatstable", &rMONMIB.Etherstatstable}
    rMONMIB.EntityData.Children["historyControlTable"] = types.YChild{"Historycontroltable", &rMONMIB.Historycontroltable}
    rMONMIB.EntityData.Children["etherHistoryTable"] = types.YChild{"Etherhistorytable", &rMONMIB.Etherhistorytable}
    rMONMIB.EntityData.Children["alarmTable"] = types.YChild{"Alarmtable", &rMONMIB.Alarmtable}
    rMONMIB.EntityData.Children["hostControlTable"] = types.YChild{"Hostcontroltable", &rMONMIB.Hostcontroltable}
    rMONMIB.EntityData.Children["hostTable"] = types.YChild{"Hosttable", &rMONMIB.Hosttable}
    rMONMIB.EntityData.Children["hostTimeTable"] = types.YChild{"Hosttimetable", &rMONMIB.Hosttimetable}
    rMONMIB.EntityData.Children["hostTopNControlTable"] = types.YChild{"Hosttopncontroltable", &rMONMIB.Hosttopncontroltable}
    rMONMIB.EntityData.Children["hostTopNTable"] = types.YChild{"Hosttopntable", &rMONMIB.Hosttopntable}
    rMONMIB.EntityData.Children["matrixControlTable"] = types.YChild{"Matrixcontroltable", &rMONMIB.Matrixcontroltable}
    rMONMIB.EntityData.Children["matrixSDTable"] = types.YChild{"Matrixsdtable", &rMONMIB.Matrixsdtable}
    rMONMIB.EntityData.Children["matrixDSTable"] = types.YChild{"Matrixdstable", &rMONMIB.Matrixdstable}
    rMONMIB.EntityData.Children["filterTable"] = types.YChild{"Filtertable", &rMONMIB.Filtertable}
    rMONMIB.EntityData.Children["channelTable"] = types.YChild{"Channeltable", &rMONMIB.Channeltable}
    rMONMIB.EntityData.Children["bufferControlTable"] = types.YChild{"Buffercontroltable", &rMONMIB.Buffercontroltable}
    rMONMIB.EntityData.Children["captureBufferTable"] = types.YChild{"Capturebuffertable", &rMONMIB.Capturebuffertable}
    rMONMIB.EntityData.Children["eventTable"] = types.YChild{"Eventtable", &rMONMIB.Eventtable}
    rMONMIB.EntityData.Children["logTable"] = types.YChild{"Logtable", &rMONMIB.Logtable}
    rMONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rMONMIB.EntityData)
}

// RMONMIB_Etherstatstable
// A list of Ethernet statistics entries.
type RMONMIB_Etherstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics kept for a particular Ethernet interface.  As an
    // example, an instance of the etherStatsPkts object might be named
    // etherStatsPkts.1. The type is slice of
    // RMONMIB_Etherstatstable_Etherstatsentry.
    Etherstatsentry []RMONMIB_Etherstatstable_Etherstatsentry
}

func (etherstatstable *RMONMIB_Etherstatstable) GetEntityData() *types.CommonEntityData {
    etherstatstable.EntityData.YFilter = etherstatstable.YFilter
    etherstatstable.EntityData.YangName = "etherStatsTable"
    etherstatstable.EntityData.BundleName = "cisco_ios_xe"
    etherstatstable.EntityData.ParentYangName = "RMON-MIB"
    etherstatstable.EntityData.SegmentPath = "etherStatsTable"
    etherstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherstatstable.EntityData.Children = make(map[string]types.YChild)
    etherstatstable.EntityData.Children["etherStatsEntry"] = types.YChild{"Etherstatsentry", nil}
    for i := range etherstatstable.Etherstatsentry {
        etherstatstable.EntityData.Children[types.GetSegmentPath(&etherstatstable.Etherstatsentry[i])] = types.YChild{"Etherstatsentry", &etherstatstable.Etherstatsentry[i]}
    }
    etherstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherstatstable.EntityData)
}

// RMONMIB_Etherstatstable_Etherstatsentry
// A collection of statistics kept for a particular
// Ethernet interface.  As an example, an instance of the
// etherStatsPkts object might be named etherStatsPkts.1
type RMONMIB_Etherstatstable_Etherstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies this
    // etherStats entry. The type is interface{} with range: 1..65535.
    Etherstatsindex interface{}

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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Etherstatsdatasource interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources. Note that this number is not necessarily the number
    // of packets dropped; it is just the number of times this condition has been
    // detected. The type is interface{} with range: 0..4294967295.
    Etherstatsdropevents interface{}

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
    Etherstatsoctets interface{}

    // The total number of packets (including bad packets, broadcast packets, and
    // multicast packets) received. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Etherstatspkts interface{}

    // The total number of good packets received that were directed to the
    // broadcast address.  Note that this does not include multicast packets. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    Etherstatsbroadcastpkts interface{}

    // The total number of good packets received that were directed to a multicast
    // address.  Note that this number does not include packets directed to the
    // broadcast address. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Etherstatsmulticastpkts interface{}

    // The total number of packets received that had a length (excluding framing
    // bits, but including FCS octets) of between 64 and 1518 octets, inclusive,
    // but had either a bad Frame Check Sequence (FCS) with an integral number of
    // octets (FCS Error) or a bad FCS with a non-integral number of octets
    // (Alignment Error). The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Etherstatscrcalignerrors interface{}

    // The total number of packets received that were less than 64 octets long
    // (excluding framing bits, but including FCS octets) and were otherwise well
    // formed. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Etherstatsundersizepkts interface{}

    // The total number of packets received that were longer than 1518 octets
    // (excluding framing bits, but including FCS octets) and were otherwise well
    // formed. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Etherstatsoversizepkts interface{}

    // The total number of packets received that were less than 64 octets in
    // length (excluding framing bits but including FCS octets) and had either a
    // bad Frame Check Sequence (FCS) with an integral number of octets (FCS
    // Error) or a bad FCS with a non-integral number of octets (Alignment Error).
    // Note that it is entirely normal for etherStatsFragments to increment.  This
    // is because it counts both runts (which are normal occurrences due to
    // collisions) and noise hits. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Etherstatsfragments interface{}

    // The total number of packets received that were longer than 1518 octets
    // (excluding framing bits, but including FCS octets), and had either a bad
    // Frame Check Sequence (FCS) with an integral number of octets (FCS Error) or
    // a bad FCS with a non-integral number of octets (Alignment Error).  Note
    // that this definition of jabber is different than the definition in
    // IEEE-802.3 section 8.2.1.5 (10BASE5) and section 10.3.1.4 (10BASE2).  These
    // documents define jabber as the condition where any packet exceeds 20 ms. 
    // The allowed range to detect jabber is between 20 ms and 150 ms. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Etherstatsjabbers interface{}

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
    Etherstatscollisions interface{}

    // The total number of packets (including bad packets) received that were 64
    // octets in length (excluding framing bits but including FCS octets). The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    Etherstatspkts64Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 65 and 127 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Etherstatspkts65To127Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 128 and 255 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Etherstatspkts128To255Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 256 and 511 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Etherstatspkts256To511Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 512 and 1023 octets in length inclusive (excluding framing bits but
    // including FCS octets). The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Etherstatspkts512To1023Octets interface{}

    // The total number of packets (including bad packets) received that were
    // between 1024 and 1518 octets in length inclusive (excluding framing bits
    // but including FCS octets). The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Etherstatspkts1024To1518Octets interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Etherstatsowner interface{}

    // The status of this etherStats entry. The type is EntryStatus.
    Etherstatsstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted      because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Etherstatsdroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Etherstatscreatetime interface{}
}

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetEntityData() *types.CommonEntityData {
    etherstatsentry.EntityData.YFilter = etherstatsentry.YFilter
    etherstatsentry.EntityData.YangName = "etherStatsEntry"
    etherstatsentry.EntityData.BundleName = "cisco_ios_xe"
    etherstatsentry.EntityData.ParentYangName = "etherStatsTable"
    etherstatsentry.EntityData.SegmentPath = "etherStatsEntry" + "[etherStatsIndex='" + fmt.Sprintf("%v", etherstatsentry.Etherstatsindex) + "']"
    etherstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherstatsentry.EntityData.Children = make(map[string]types.YChild)
    etherstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    etherstatsentry.EntityData.Leafs["etherStatsIndex"] = types.YLeaf{"Etherstatsindex", etherstatsentry.Etherstatsindex}
    etherstatsentry.EntityData.Leafs["etherStatsDataSource"] = types.YLeaf{"Etherstatsdatasource", etherstatsentry.Etherstatsdatasource}
    etherstatsentry.EntityData.Leafs["etherStatsDropEvents"] = types.YLeaf{"Etherstatsdropevents", etherstatsentry.Etherstatsdropevents}
    etherstatsentry.EntityData.Leafs["etherStatsOctets"] = types.YLeaf{"Etherstatsoctets", etherstatsentry.Etherstatsoctets}
    etherstatsentry.EntityData.Leafs["etherStatsPkts"] = types.YLeaf{"Etherstatspkts", etherstatsentry.Etherstatspkts}
    etherstatsentry.EntityData.Leafs["etherStatsBroadcastPkts"] = types.YLeaf{"Etherstatsbroadcastpkts", etherstatsentry.Etherstatsbroadcastpkts}
    etherstatsentry.EntityData.Leafs["etherStatsMulticastPkts"] = types.YLeaf{"Etherstatsmulticastpkts", etherstatsentry.Etherstatsmulticastpkts}
    etherstatsentry.EntityData.Leafs["etherStatsCRCAlignErrors"] = types.YLeaf{"Etherstatscrcalignerrors", etherstatsentry.Etherstatscrcalignerrors}
    etherstatsentry.EntityData.Leafs["etherStatsUndersizePkts"] = types.YLeaf{"Etherstatsundersizepkts", etherstatsentry.Etherstatsundersizepkts}
    etherstatsentry.EntityData.Leafs["etherStatsOversizePkts"] = types.YLeaf{"Etherstatsoversizepkts", etherstatsentry.Etherstatsoversizepkts}
    etherstatsentry.EntityData.Leafs["etherStatsFragments"] = types.YLeaf{"Etherstatsfragments", etherstatsentry.Etherstatsfragments}
    etherstatsentry.EntityData.Leafs["etherStatsJabbers"] = types.YLeaf{"Etherstatsjabbers", etherstatsentry.Etherstatsjabbers}
    etherstatsentry.EntityData.Leafs["etherStatsCollisions"] = types.YLeaf{"Etherstatscollisions", etherstatsentry.Etherstatscollisions}
    etherstatsentry.EntityData.Leafs["etherStatsPkts64Octets"] = types.YLeaf{"Etherstatspkts64Octets", etherstatsentry.Etherstatspkts64Octets}
    etherstatsentry.EntityData.Leafs["etherStatsPkts65to127Octets"] = types.YLeaf{"Etherstatspkts65To127Octets", etherstatsentry.Etherstatspkts65To127Octets}
    etherstatsentry.EntityData.Leafs["etherStatsPkts128to255Octets"] = types.YLeaf{"Etherstatspkts128To255Octets", etherstatsentry.Etherstatspkts128To255Octets}
    etherstatsentry.EntityData.Leafs["etherStatsPkts256to511Octets"] = types.YLeaf{"Etherstatspkts256To511Octets", etherstatsentry.Etherstatspkts256To511Octets}
    etherstatsentry.EntityData.Leafs["etherStatsPkts512to1023Octets"] = types.YLeaf{"Etherstatspkts512To1023Octets", etherstatsentry.Etherstatspkts512To1023Octets}
    etherstatsentry.EntityData.Leafs["etherStatsPkts1024to1518Octets"] = types.YLeaf{"Etherstatspkts1024To1518Octets", etherstatsentry.Etherstatspkts1024To1518Octets}
    etherstatsentry.EntityData.Leafs["etherStatsOwner"] = types.YLeaf{"Etherstatsowner", etherstatsentry.Etherstatsowner}
    etherstatsentry.EntityData.Leafs["etherStatsStatus"] = types.YLeaf{"Etherstatsstatus", etherstatsentry.Etherstatsstatus}
    etherstatsentry.EntityData.Leafs["etherStatsDroppedFrames"] = types.YLeaf{"Etherstatsdroppedframes", etherstatsentry.Etherstatsdroppedframes}
    etherstatsentry.EntityData.Leafs["etherStatsCreateTime"] = types.YLeaf{"Etherstatscreatetime", etherstatsentry.Etherstatscreatetime}
    return &(etherstatsentry.EntityData)
}

// RMONMIB_Historycontroltable
// A list of history control entries.
type RMONMIB_Historycontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up a periodic sampling of statistics.  As an
    // example, an instance of the historyControlInterval object might be named
    // historyControlInterval.2. The type is slice of
    // RMONMIB_Historycontroltable_Historycontrolentry.
    Historycontrolentry []RMONMIB_Historycontroltable_Historycontrolentry
}

func (historycontroltable *RMONMIB_Historycontroltable) GetEntityData() *types.CommonEntityData {
    historycontroltable.EntityData.YFilter = historycontroltable.YFilter
    historycontroltable.EntityData.YangName = "historyControlTable"
    historycontroltable.EntityData.BundleName = "cisco_ios_xe"
    historycontroltable.EntityData.ParentYangName = "RMON-MIB"
    historycontroltable.EntityData.SegmentPath = "historyControlTable"
    historycontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    historycontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    historycontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    historycontroltable.EntityData.Children = make(map[string]types.YChild)
    historycontroltable.EntityData.Children["historyControlEntry"] = types.YChild{"Historycontrolentry", nil}
    for i := range historycontroltable.Historycontrolentry {
        historycontroltable.EntityData.Children[types.GetSegmentPath(&historycontroltable.Historycontrolentry[i])] = types.YChild{"Historycontrolentry", &historycontroltable.Historycontrolentry[i]}
    }
    historycontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(historycontroltable.EntityData)
}

// RMONMIB_Historycontroltable_Historycontrolentry
// A list of parameters that set up a periodic sampling of
// statistics.  As an example, an instance of the
// historyControlInterval object might be named
// historyControlInterval.2
type RMONMIB_Historycontroltable_Historycontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // historyControl table.  Each such entry defines a set of samples at a
    // particular interval for an interface on the device. The type is interface{}
    // with range: 1..65535.
    Historycontrolindex interface{}

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
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Historycontroldatasource interface{}

    // The requested number of discrete time intervals over which data is to be
    // saved in the part of the media-specific table associated with this
    // historyControlEntry.  When this object is created or modified, the probe
    // should set historyControlBucketsGranted as closely to this object as is
    // possible for the particular probe implementation and available resources.
    // The type is interface{} with range: 1..65535.
    Historycontrolbucketsrequested interface{}

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
    Historycontrolbucketsgranted interface{}

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
    Historycontrolinterval interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Historycontrolowner interface{}

    // The status of this historyControl entry.  Each instance of the
    // media-specific table associated with this historyControlEntry will be
    // deleted by the agent if this historyControlEntry is not equal to valid(1).
    // The type is EntryStatus.
    Historycontrolstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this      collection.  This count does not include packets that were
    // not counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Historycontroldroppedframes interface{}
}

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetEntityData() *types.CommonEntityData {
    historycontrolentry.EntityData.YFilter = historycontrolentry.YFilter
    historycontrolentry.EntityData.YangName = "historyControlEntry"
    historycontrolentry.EntityData.BundleName = "cisco_ios_xe"
    historycontrolentry.EntityData.ParentYangName = "historyControlTable"
    historycontrolentry.EntityData.SegmentPath = "historyControlEntry" + "[historyControlIndex='" + fmt.Sprintf("%v", historycontrolentry.Historycontrolindex) + "']"
    historycontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    historycontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    historycontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    historycontrolentry.EntityData.Children = make(map[string]types.YChild)
    historycontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    historycontrolentry.EntityData.Leafs["historyControlIndex"] = types.YLeaf{"Historycontrolindex", historycontrolentry.Historycontrolindex}
    historycontrolentry.EntityData.Leafs["historyControlDataSource"] = types.YLeaf{"Historycontroldatasource", historycontrolentry.Historycontroldatasource}
    historycontrolentry.EntityData.Leafs["historyControlBucketsRequested"] = types.YLeaf{"Historycontrolbucketsrequested", historycontrolentry.Historycontrolbucketsrequested}
    historycontrolentry.EntityData.Leafs["historyControlBucketsGranted"] = types.YLeaf{"Historycontrolbucketsgranted", historycontrolentry.Historycontrolbucketsgranted}
    historycontrolentry.EntityData.Leafs["historyControlInterval"] = types.YLeaf{"Historycontrolinterval", historycontrolentry.Historycontrolinterval}
    historycontrolentry.EntityData.Leafs["historyControlOwner"] = types.YLeaf{"Historycontrolowner", historycontrolentry.Historycontrolowner}
    historycontrolentry.EntityData.Leafs["historyControlStatus"] = types.YLeaf{"Historycontrolstatus", historycontrolentry.Historycontrolstatus}
    historycontrolentry.EntityData.Leafs["historyControlDroppedFrames"] = types.YLeaf{"Historycontroldroppedframes", historycontrolentry.Historycontroldroppedframes}
    return &(historycontrolentry.EntityData)
}

// RMONMIB_Etherhistorytable
// A list of Ethernet history entries.
type RMONMIB_Etherhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An historical sample of Ethernet statistics on a particular Ethernet
    // interface.  This sample is associated with the historyControlEntry which
    // set up the parameters for a regular collection of these samples.  As an
    // example, an instance of the etherHistoryPkts object might be named
    // etherHistoryPkts.2.89. The type is slice of
    // RMONMIB_Etherhistorytable_Etherhistoryentry.
    Etherhistoryentry []RMONMIB_Etherhistorytable_Etherhistoryentry
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetEntityData() *types.CommonEntityData {
    etherhistorytable.EntityData.YFilter = etherhistorytable.YFilter
    etherhistorytable.EntityData.YangName = "etherHistoryTable"
    etherhistorytable.EntityData.BundleName = "cisco_ios_xe"
    etherhistorytable.EntityData.ParentYangName = "RMON-MIB"
    etherhistorytable.EntityData.SegmentPath = "etherHistoryTable"
    etherhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherhistorytable.EntityData.Children = make(map[string]types.YChild)
    etherhistorytable.EntityData.Children["etherHistoryEntry"] = types.YChild{"Etherhistoryentry", nil}
    for i := range etherhistorytable.Etherhistoryentry {
        etherhistorytable.EntityData.Children[types.GetSegmentPath(&etherhistorytable.Etherhistoryentry[i])] = types.YChild{"Etherhistoryentry", &etherhistorytable.Etherhistoryentry[i]}
    }
    etherhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherhistorytable.EntityData)
}

// RMONMIB_Etherhistorytable_Etherhistoryentry
// An historical sample of Ethernet statistics on a particular
// Ethernet interface.  This sample is associated with the
// historyControlEntry which set up the parameters for
// a regular collection of these samples.  As an example, an
// instance of the etherHistoryPkts object might be named
// etherHistoryPkts.2.89
type RMONMIB_Etherhistorytable_Etherhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The history of which this entry is a part.  The
    // history identified by a particular value of this index is the same history
    // as identified by the same value of historyControlIndex. The type is
    // interface{} with range: 1..65535.
    Etherhistoryindex interface{}

    // This attribute is a key. An index that uniquely identifies the particular
    // sample this entry represents among all samples associated with the same
    // historyControlEntry. This index starts at 1 and increases by one as each
    // new sample is taken. The type is interface{} with range: 1..2147483647.
    Etherhistorysampleindex interface{}

    // The value of sysUpTime at the start of the interval over which this sample
    // was measured.  If the probe keeps track of the time of day, it should start
    // the first sample of the history at a time such that when the next hour of
    // the day begins, a sample is started at that instant.  Note that following
    // this rule may require the probe to delay collecting the first sample of the
    // history, as each sample must be of the same interval.  Also note that the
    // sample which is currently being collected is not accessible in this table
    // until the end of its interval. The type is interface{} with range:
    // 0..4294967295.
    Etherhistoryintervalstart interface{}

    // The total number of events in which packets were dropped by the probe due
    // to lack of resources during this sampling interval.  Note that this number
    // is not necessarily the number of packets dropped, it is just the number of
    // times this condition has been detected. The type is interface{} with range:
    // 0..4294967295.
    Etherhistorydropevents interface{}

    // The total number of octets of data (including those in bad packets)
    // received on the network (excluding framing bits but including FCS octets).
    // The type is interface{} with range: 0..4294967295. Units are Octets.
    Etherhistoryoctets interface{}

    // The number of packets (including bad packets) received during this sampling
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Etherhistorypkts interface{}

    // The number of good packets received during this sampling interval that were
    // directed to the broadcast address. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Etherhistorybroadcastpkts interface{}

    // The number of good packets received during this sampling interval that were
    // directed to a multicast address.  Note that this number does not include
    // packets addressed to the broadcast address. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    Etherhistorymulticastpkts interface{}

    // The number of packets received during this sampling interval that had a
    // length (excluding framing bits but including FCS octets) between 64 and
    // 1518 octets, inclusive, but had either a bad Frame Check Sequence (FCS)
    // with an integral number of octets (FCS Error) or a bad FCS with a
    // non-integral number of octets (Alignment Error). The type is interface{}
    // with range: 0..4294967295. Units are Packets.
    Etherhistorycrcalignerrors interface{}

    // The number of packets received during this sampling interval that were less
    // than 64 octets long (excluding framing bits but including FCS octets) and
    // were otherwise well formed. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Etherhistoryundersizepkts interface{}

    // The number of packets received during this sampling interval that were
    // longer than 1518 octets (excluding framing bits but including FCS octets)
    // but were otherwise well formed. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Etherhistoryoversizepkts interface{}

    // The total number of packets received during this sampling interval that
    // were less than 64 octets in length (excluding framing bits but including
    // FCS octets) had either a bad Frame Check Sequence (FCS) with an integral
    // number of octets (FCS Error) or a bad FCS with a non-integral number of
    // octets (Alignment Error).  Note that it is entirely normal for
    // etherHistoryFragments to increment.  This is because it counts both runts
    // (which are normal occurrences due to collisions) and noise hits. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Etherhistoryfragments interface{}

    // The number of packets received during this sampling interval that were
    // longer than 1518 octets (excluding framing bits but including FCS octets),
    // and  had either a bad Frame Check Sequence (FCS) with an integral number of
    // octets (FCS Error) or a bad FCS with a non-integral number of octets
    // (Alignment Error).  Note that this definition of jabber is different than
    // the definition in IEEE-802.3 section 8.2.1.5 (10BASE5) and section 10.3.1.4
    // (10BASE2).  These documents define jabber as the condition where any packet
    // exceeds 20 ms.  The allowed range to detect jabber is between 20 ms and 150
    // ms. The type is interface{} with range: 0..4294967295. Units are Packets.
    Etherhistoryjabbers interface{}

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
    Etherhistorycollisions interface{}

    // The best estimate of the mean physical layer network utilization on this
    // interface during this sampling interval, in hundredths of a percent. The
    // type is interface{} with range: 0..10000.
    Etherhistoryutilization interface{}
}

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetEntityData() *types.CommonEntityData {
    etherhistoryentry.EntityData.YFilter = etherhistoryentry.YFilter
    etherhistoryentry.EntityData.YangName = "etherHistoryEntry"
    etherhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    etherhistoryentry.EntityData.ParentYangName = "etherHistoryTable"
    etherhistoryentry.EntityData.SegmentPath = "etherHistoryEntry" + "[etherHistoryIndex='" + fmt.Sprintf("%v", etherhistoryentry.Etherhistoryindex) + "']" + "[etherHistorySampleIndex='" + fmt.Sprintf("%v", etherhistoryentry.Etherhistorysampleindex) + "']"
    etherhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherhistoryentry.EntityData.Children = make(map[string]types.YChild)
    etherhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    etherhistoryentry.EntityData.Leafs["etherHistoryIndex"] = types.YLeaf{"Etherhistoryindex", etherhistoryentry.Etherhistoryindex}
    etherhistoryentry.EntityData.Leafs["etherHistorySampleIndex"] = types.YLeaf{"Etherhistorysampleindex", etherhistoryentry.Etherhistorysampleindex}
    etherhistoryentry.EntityData.Leafs["etherHistoryIntervalStart"] = types.YLeaf{"Etherhistoryintervalstart", etherhistoryentry.Etherhistoryintervalstart}
    etherhistoryentry.EntityData.Leafs["etherHistoryDropEvents"] = types.YLeaf{"Etherhistorydropevents", etherhistoryentry.Etherhistorydropevents}
    etherhistoryentry.EntityData.Leafs["etherHistoryOctets"] = types.YLeaf{"Etherhistoryoctets", etherhistoryentry.Etherhistoryoctets}
    etherhistoryentry.EntityData.Leafs["etherHistoryPkts"] = types.YLeaf{"Etherhistorypkts", etherhistoryentry.Etherhistorypkts}
    etherhistoryentry.EntityData.Leafs["etherHistoryBroadcastPkts"] = types.YLeaf{"Etherhistorybroadcastpkts", etherhistoryentry.Etherhistorybroadcastpkts}
    etherhistoryentry.EntityData.Leafs["etherHistoryMulticastPkts"] = types.YLeaf{"Etherhistorymulticastpkts", etherhistoryentry.Etherhistorymulticastpkts}
    etherhistoryentry.EntityData.Leafs["etherHistoryCRCAlignErrors"] = types.YLeaf{"Etherhistorycrcalignerrors", etherhistoryentry.Etherhistorycrcalignerrors}
    etherhistoryentry.EntityData.Leafs["etherHistoryUndersizePkts"] = types.YLeaf{"Etherhistoryundersizepkts", etherhistoryentry.Etherhistoryundersizepkts}
    etherhistoryentry.EntityData.Leafs["etherHistoryOversizePkts"] = types.YLeaf{"Etherhistoryoversizepkts", etherhistoryentry.Etherhistoryoversizepkts}
    etherhistoryentry.EntityData.Leafs["etherHistoryFragments"] = types.YLeaf{"Etherhistoryfragments", etherhistoryentry.Etherhistoryfragments}
    etherhistoryentry.EntityData.Leafs["etherHistoryJabbers"] = types.YLeaf{"Etherhistoryjabbers", etherhistoryentry.Etherhistoryjabbers}
    etherhistoryentry.EntityData.Leafs["etherHistoryCollisions"] = types.YLeaf{"Etherhistorycollisions", etherhistoryentry.Etherhistorycollisions}
    etherhistoryentry.EntityData.Leafs["etherHistoryUtilization"] = types.YLeaf{"Etherhistoryutilization", etherhistoryentry.Etherhistoryutilization}
    return &(etherhistoryentry.EntityData)
}

// RMONMIB_Alarmtable
// A list of alarm entries.
type RMONMIB_Alarmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up a periodic checking for alarm conditions. 
    // For example, an instance of the alarmValue object might be named
    // alarmValue.8. The type is slice of RMONMIB_Alarmtable_Alarmentry.
    Alarmentry []RMONMIB_Alarmtable_Alarmentry
}

func (alarmtable *RMONMIB_Alarmtable) GetEntityData() *types.CommonEntityData {
    alarmtable.EntityData.YFilter = alarmtable.YFilter
    alarmtable.EntityData.YangName = "alarmTable"
    alarmtable.EntityData.BundleName = "cisco_ios_xe"
    alarmtable.EntityData.ParentYangName = "RMON-MIB"
    alarmtable.EntityData.SegmentPath = "alarmTable"
    alarmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alarmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alarmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alarmtable.EntityData.Children = make(map[string]types.YChild)
    alarmtable.EntityData.Children["alarmEntry"] = types.YChild{"Alarmentry", nil}
    for i := range alarmtable.Alarmentry {
        alarmtable.EntityData.Children[types.GetSegmentPath(&alarmtable.Alarmentry[i])] = types.YChild{"Alarmentry", &alarmtable.Alarmentry[i]}
    }
    alarmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alarmtable.EntityData)
}

// RMONMIB_Alarmtable_Alarmentry
// A list of parameters that set up a periodic checking
// for alarm conditions.  For example, an instance of the
// alarmValue object might be named alarmValue.8
type RMONMIB_Alarmtable_Alarmentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // alarm table.  Each such entry defines a diagnostic sample at a particular
    // interval for an object on the device. The type is interface{} with range:
    // 1..65535.
    Alarmindex interface{}

    // The interval in seconds over which the data is sampled and compared with
    // the rising and falling thresholds.  When setting this variable, care should
    // be taken in the case of deltaValue sampling - the interval should be set
    // short enough that the sampled variable is very unlikely to increase or
    // decrease by more than 2^31 - 1 during a single sampling interval.  This
    // object may not be modified if the associated alarmStatus object is equal to
    // valid(1). The type is interface{} with range: -2147483648..2147483647.
    // Units are Seconds.
    Alarminterval interface{}

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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Alarmvariable interface{}

    // The method of sampling the selected variable and calculating the value to
    // be compared against the thresholds.  If the value of this object is
    // absoluteValue(1), the value of the selected variable will be compared
    // directly with the thresholds at the end of the sampling interval.  If the
    // value of this object is deltaValue(2), the value of the selected variable
    // at the last sample will be subtracted from the current value, and the
    // difference compared with the thresholds.  This object may not be modified
    // if the associated alarmStatus object is equal to valid(1). The type is
    // Alarmsampletype.
    Alarmsampletype interface{}

    // The value of the statistic during the last sampling period.  For example,
    // if the sample type is deltaValue, this value will be the difference between
    // the samples at the beginning and end of the period.  If the sample type is
    // absoluteValue, this value will be the sampled value at the end of the
    // period. This is the value that is compared with the rising and falling
    // thresholds.  The value during the current sampling period is not made
    // available until the period is completed and will remain available until the
    // next period completes. The type is interface{} with range:
    // -2147483648..2147483647.
    Alarmvalue interface{}

    // The alarm that may be sent when this entry is first set to valid.  If the
    // first sample after this entry becomes valid is greater than or equal to the
    // risingThreshold and alarmStartupAlarm is equal to risingAlarm(1) or
    // risingOrFallingAlarm(3), then a single rising alarm will be generated.  If
    // the first sample after this entry becomes valid is less than or equal to
    // the fallingThreshold and alarmStartupAlarm is equal to fallingAlarm(2) or
    // risingOrFallingAlarm(3), then a single falling alarm will be generated. 
    // This object may not be modified if the associated alarmStatus object is
    // equal to valid(1). The type is Alarmstartupalarm.
    Alarmstartupalarm interface{}

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
    Alarmrisingthreshold interface{}

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
    Alarmfallingthreshold interface{}

    // The index of the eventEntry that is used when a rising threshold is
    // crossed.  The eventEntry identified by a particular value of this index is
    // the same as identified by the same value of the eventIndex object.  If
    // there is no corresponding entry in the eventTable, then no association
    // exists.  In particular, if this value is zero, no associated event will be
    // generated, as zero is not a valid event index.  This object may not be
    // modified if the associated alarmStatus object is equal to valid(1). The
    // type is interface{} with range: 0..65535.
    Alarmrisingeventindex interface{}

    // The index of the eventEntry that is used when a falling threshold is
    // crossed.  The eventEntry identified by a particular value of this index is
    // the same as identified by the same value of the eventIndex object.  If
    // there is no corresponding entry in the eventTable, then no association
    // exists.  In particular, if this value is zero, no associated event will be
    // generated, as zero is not a valid event index.  This object may not be
    // modified if the associated alarmStatus object is equal to valid(1). The
    // type is interface{} with range: 0..65535.
    Alarmfallingeventindex interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Alarmowner interface{}

    // The status of this alarm entry. The type is EntryStatus.
    Alarmstatus interface{}
}

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetEntityData() *types.CommonEntityData {
    alarmentry.EntityData.YFilter = alarmentry.YFilter
    alarmentry.EntityData.YangName = "alarmEntry"
    alarmentry.EntityData.BundleName = "cisco_ios_xe"
    alarmentry.EntityData.ParentYangName = "alarmTable"
    alarmentry.EntityData.SegmentPath = "alarmEntry" + "[alarmIndex='" + fmt.Sprintf("%v", alarmentry.Alarmindex) + "']"
    alarmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    alarmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    alarmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    alarmentry.EntityData.Children = make(map[string]types.YChild)
    alarmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmentry.EntityData.Leafs["alarmIndex"] = types.YLeaf{"Alarmindex", alarmentry.Alarmindex}
    alarmentry.EntityData.Leafs["alarmInterval"] = types.YLeaf{"Alarminterval", alarmentry.Alarminterval}
    alarmentry.EntityData.Leafs["alarmVariable"] = types.YLeaf{"Alarmvariable", alarmentry.Alarmvariable}
    alarmentry.EntityData.Leafs["alarmSampleType"] = types.YLeaf{"Alarmsampletype", alarmentry.Alarmsampletype}
    alarmentry.EntityData.Leafs["alarmValue"] = types.YLeaf{"Alarmvalue", alarmentry.Alarmvalue}
    alarmentry.EntityData.Leafs["alarmStartupAlarm"] = types.YLeaf{"Alarmstartupalarm", alarmentry.Alarmstartupalarm}
    alarmentry.EntityData.Leafs["alarmRisingThreshold"] = types.YLeaf{"Alarmrisingthreshold", alarmentry.Alarmrisingthreshold}
    alarmentry.EntityData.Leafs["alarmFallingThreshold"] = types.YLeaf{"Alarmfallingthreshold", alarmentry.Alarmfallingthreshold}
    alarmentry.EntityData.Leafs["alarmRisingEventIndex"] = types.YLeaf{"Alarmrisingeventindex", alarmentry.Alarmrisingeventindex}
    alarmentry.EntityData.Leafs["alarmFallingEventIndex"] = types.YLeaf{"Alarmfallingeventindex", alarmentry.Alarmfallingeventindex}
    alarmentry.EntityData.Leafs["alarmOwner"] = types.YLeaf{"Alarmowner", alarmentry.Alarmowner}
    alarmentry.EntityData.Leafs["alarmStatus"] = types.YLeaf{"Alarmstatus", alarmentry.Alarmstatus}
    return &(alarmentry.EntityData)
}

// RMONMIB_Alarmtable_Alarmentry_Alarmsampletype represents alarmStatus object is equal to valid(1).
type RMONMIB_Alarmtable_Alarmentry_Alarmsampletype string

const (
    RMONMIB_Alarmtable_Alarmentry_Alarmsampletype_absoluteValue RMONMIB_Alarmtable_Alarmentry_Alarmsampletype = "absoluteValue"

    RMONMIB_Alarmtable_Alarmentry_Alarmsampletype_deltaValue RMONMIB_Alarmtable_Alarmentry_Alarmsampletype = "deltaValue"
)

// RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm represents alarmStatus object is equal to valid(1).
type RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm string

const (
    RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm_risingAlarm RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm = "risingAlarm"

    RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm_fallingAlarm RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm = "fallingAlarm"

    RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm_risingOrFallingAlarm RMONMIB_Alarmtable_Alarmentry_Alarmstartupalarm = "risingOrFallingAlarm"
)

// RMONMIB_Hostcontroltable
// A list of host table control entries.
type RMONMIB_Hostcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of parameters that set up the discovery of hosts on a particular
    // interface and the collection of statistics about these hosts.  For example,
    // an instance of the hostControlTableSize object might be named
    // hostControlTableSize.1. The type is slice of
    // RMONMIB_Hostcontroltable_Hostcontrolentry.
    Hostcontrolentry []RMONMIB_Hostcontroltable_Hostcontrolentry
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetEntityData() *types.CommonEntityData {
    hostcontroltable.EntityData.YFilter = hostcontroltable.YFilter
    hostcontroltable.EntityData.YangName = "hostControlTable"
    hostcontroltable.EntityData.BundleName = "cisco_ios_xe"
    hostcontroltable.EntityData.ParentYangName = "RMON-MIB"
    hostcontroltable.EntityData.SegmentPath = "hostControlTable"
    hostcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostcontroltable.EntityData.Children = make(map[string]types.YChild)
    hostcontroltable.EntityData.Children["hostControlEntry"] = types.YChild{"Hostcontrolentry", nil}
    for i := range hostcontroltable.Hostcontrolentry {
        hostcontroltable.EntityData.Children[types.GetSegmentPath(&hostcontroltable.Hostcontrolentry[i])] = types.YChild{"Hostcontrolentry", &hostcontroltable.Hostcontrolentry[i]}
    }
    hostcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hostcontroltable.EntityData)
}

// RMONMIB_Hostcontroltable_Hostcontrolentry
// A list of parameters that set up the discovery of hosts
// on a particular interface and the collection of statistics
// about these hosts.  For example, an instance of the
// hostControlTableSize object might be named
// hostControlTableSize.1
type RMONMIB_Hostcontroltable_Hostcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostControl table.  Each such entry defines a function that discovers hosts
    // on a particular interface and places statistics about them in the hostTable
    // and the hostTimeTable on behalf of this hostControlEntry. The type is
    // interface{} with range: 1..65535.
    Hostcontrolindex interface{}

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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Hostcontroldatasource interface{}

    // The number of hostEntries in the hostTable and the hostTimeTable associated
    // with this hostControlEntry. The type is interface{} with range:
    // -2147483648..2147483647.
    Hostcontroltablesize interface{}

    // The value of sysUpTime when the last entry was deleted from the portion of
    // the hostTable associated with this hostControlEntry.  If no deletions have
    // occurred, this value shall be zero. The type is interface{} with range:
    // 0..4294967295.
    Hostcontrollastdeletetime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Hostcontrolowner interface{}

    // The status of this hostControl entry.  If this object is not equal to
    // valid(1), all associated entries in the hostTable, hostTimeTable, and the
    // hostTopNTable shall be deleted by the agent. The type is EntryStatus.
    Hostcontrolstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Hostcontroldroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Hostcontrolcreatetime interface{}
}

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetEntityData() *types.CommonEntityData {
    hostcontrolentry.EntityData.YFilter = hostcontrolentry.YFilter
    hostcontrolentry.EntityData.YangName = "hostControlEntry"
    hostcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    hostcontrolentry.EntityData.ParentYangName = "hostControlTable"
    hostcontrolentry.EntityData.SegmentPath = "hostControlEntry" + "[hostControlIndex='" + fmt.Sprintf("%v", hostcontrolentry.Hostcontrolindex) + "']"
    hostcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostcontrolentry.EntityData.Children = make(map[string]types.YChild)
    hostcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hostcontrolentry.EntityData.Leafs["hostControlIndex"] = types.YLeaf{"Hostcontrolindex", hostcontrolentry.Hostcontrolindex}
    hostcontrolentry.EntityData.Leafs["hostControlDataSource"] = types.YLeaf{"Hostcontroldatasource", hostcontrolentry.Hostcontroldatasource}
    hostcontrolentry.EntityData.Leafs["hostControlTableSize"] = types.YLeaf{"Hostcontroltablesize", hostcontrolentry.Hostcontroltablesize}
    hostcontrolentry.EntityData.Leafs["hostControlLastDeleteTime"] = types.YLeaf{"Hostcontrollastdeletetime", hostcontrolentry.Hostcontrollastdeletetime}
    hostcontrolentry.EntityData.Leafs["hostControlOwner"] = types.YLeaf{"Hostcontrolowner", hostcontrolentry.Hostcontrolowner}
    hostcontrolentry.EntityData.Leafs["hostControlStatus"] = types.YLeaf{"Hostcontrolstatus", hostcontrolentry.Hostcontrolstatus}
    hostcontrolentry.EntityData.Leafs["hostControlDroppedFrames"] = types.YLeaf{"Hostcontroldroppedframes", hostcontrolentry.Hostcontroldroppedframes}
    hostcontrolentry.EntityData.Leafs["hostControlCreateTime"] = types.YLeaf{"Hostcontrolcreatetime", hostcontrolentry.Hostcontrolcreatetime}
    return &(hostcontrolentry.EntityData)
}

// RMONMIB_Hosttable
// A list of host entries.
type RMONMIB_Hosttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular host that has been discovered
    // on an interface of this device.  For example, an instance of the
    // hostOutBroadcastPkts object might be named
    // hostOutBroadcastPkts.1.6.8.0.32.27.3.176. The type is slice of
    // RMONMIB_Hosttable_Hostentry.
    Hostentry []RMONMIB_Hosttable_Hostentry
}

func (hosttable *RMONMIB_Hosttable) GetEntityData() *types.CommonEntityData {
    hosttable.EntityData.YFilter = hosttable.YFilter
    hosttable.EntityData.YangName = "hostTable"
    hosttable.EntityData.BundleName = "cisco_ios_xe"
    hosttable.EntityData.ParentYangName = "RMON-MIB"
    hosttable.EntityData.SegmentPath = "hostTable"
    hosttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttable.EntityData.Children = make(map[string]types.YChild)
    hosttable.EntityData.Children["hostEntry"] = types.YChild{"Hostentry", nil}
    for i := range hosttable.Hostentry {
        hosttable.EntityData.Children[types.GetSegmentPath(&hosttable.Hostentry[i])] = types.YChild{"Hostentry", &hosttable.Hostentry[i]}
    }
    hosttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosttable.EntityData)
}

// RMONMIB_Hosttable_Hostentry
// A collection of statistics for a particular host that has
// been discovered on an interface of this device.  For example,
// an instance of the hostOutBroadcastPkts object might be
// named hostOutBroadcastPkts.1.6.8.0.32.27.3.176
type RMONMIB_Hosttable_Hostentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected host statistics of which this
    // entry is a part.  The set of hosts identified by a particular value of this
    // index is associated with the hostControlEntry as identified by the same
    // value of hostControlIndex. The type is interface{} with range: 1..65535.
    Hostindex interface{}

    // This attribute is a key. The physical address of this host. The type is
    // string.
    Hostaddress interface{}

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
    Hostcreationorder interface{}

    // The number of good packets transmitted to this address since it was added
    // to the hostTable. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Hostinpkts interface{}

    // The number of packets, including bad packets, transmitted by this address
    // since it was added to the hostTable. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Hostoutpkts interface{}

    // The number of octets transmitted to this address since it was added to the
    // hostTable (excluding framing bits but including FCS octets), except for
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Hostinoctets interface{}

    // The number of octets transmitted by this address since it was added to the
    // hostTable (excluding framing bits but including FCS octets), including
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Hostoutoctets interface{}

    // The number of bad packets transmitted by this address since this host was
    // added to the hostTable. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Hostouterrors interface{}

    // The number of good packets transmitted by this address that were directed
    // to the broadcast address since this host was added to the hostTable. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    Hostoutbroadcastpkts interface{}

    // The number of good packets transmitted by this address that were directed
    // to a multicast address since this host was added to the hostTable. Note
    // that this number does not include packets directed to the broadcast
    // address. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Hostoutmulticastpkts interface{}
}

func (hostentry *RMONMIB_Hosttable_Hostentry) GetEntityData() *types.CommonEntityData {
    hostentry.EntityData.YFilter = hostentry.YFilter
    hostentry.EntityData.YangName = "hostEntry"
    hostentry.EntityData.BundleName = "cisco_ios_xe"
    hostentry.EntityData.ParentYangName = "hostTable"
    hostentry.EntityData.SegmentPath = "hostEntry" + "[hostIndex='" + fmt.Sprintf("%v", hostentry.Hostindex) + "']" + "[hostAddress='" + fmt.Sprintf("%v", hostentry.Hostaddress) + "']"
    hostentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hostentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hostentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hostentry.EntityData.Children = make(map[string]types.YChild)
    hostentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hostentry.EntityData.Leafs["hostIndex"] = types.YLeaf{"Hostindex", hostentry.Hostindex}
    hostentry.EntityData.Leafs["hostAddress"] = types.YLeaf{"Hostaddress", hostentry.Hostaddress}
    hostentry.EntityData.Leafs["hostCreationOrder"] = types.YLeaf{"Hostcreationorder", hostentry.Hostcreationorder}
    hostentry.EntityData.Leafs["hostInPkts"] = types.YLeaf{"Hostinpkts", hostentry.Hostinpkts}
    hostentry.EntityData.Leafs["hostOutPkts"] = types.YLeaf{"Hostoutpkts", hostentry.Hostoutpkts}
    hostentry.EntityData.Leafs["hostInOctets"] = types.YLeaf{"Hostinoctets", hostentry.Hostinoctets}
    hostentry.EntityData.Leafs["hostOutOctets"] = types.YLeaf{"Hostoutoctets", hostentry.Hostoutoctets}
    hostentry.EntityData.Leafs["hostOutErrors"] = types.YLeaf{"Hostouterrors", hostentry.Hostouterrors}
    hostentry.EntityData.Leafs["hostOutBroadcastPkts"] = types.YLeaf{"Hostoutbroadcastpkts", hostentry.Hostoutbroadcastpkts}
    hostentry.EntityData.Leafs["hostOutMulticastPkts"] = types.YLeaf{"Hostoutmulticastpkts", hostentry.Hostoutmulticastpkts}
    return &(hostentry.EntityData)
}

// RMONMIB_Hosttimetable
// A list of time-ordered host table entries.
type RMONMIB_Hosttimetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for a particular host that has been discovered
    // on an interface of this device.  This collection includes the relative
    // ordering of the creation time of this object.  For example, an instance of
    // the hostTimeOutBroadcastPkts object might be named
    // hostTimeOutBroadcastPkts.1.687. The type is slice of
    // RMONMIB_Hosttimetable_Hosttimeentry.
    Hosttimeentry []RMONMIB_Hosttimetable_Hosttimeentry
}

func (hosttimetable *RMONMIB_Hosttimetable) GetEntityData() *types.CommonEntityData {
    hosttimetable.EntityData.YFilter = hosttimetable.YFilter
    hosttimetable.EntityData.YangName = "hostTimeTable"
    hosttimetable.EntityData.BundleName = "cisco_ios_xe"
    hosttimetable.EntityData.ParentYangName = "RMON-MIB"
    hosttimetable.EntityData.SegmentPath = "hostTimeTable"
    hosttimetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttimetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttimetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttimetable.EntityData.Children = make(map[string]types.YChild)
    hosttimetable.EntityData.Children["hostTimeEntry"] = types.YChild{"Hosttimeentry", nil}
    for i := range hosttimetable.Hosttimeentry {
        hosttimetable.EntityData.Children[types.GetSegmentPath(&hosttimetable.Hosttimeentry[i])] = types.YChild{"Hosttimeentry", &hosttimetable.Hosttimeentry[i]}
    }
    hosttimetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosttimetable.EntityData)
}

// RMONMIB_Hosttimetable_Hosttimeentry
// A collection of statistics for a particular host that has
// been discovered on an interface of this device.  This
// collection includes the relative ordering of the creation
// time of this object.  For example, an instance of the
// hostTimeOutBroadcastPkts object might be named
// hostTimeOutBroadcastPkts.1.687
type RMONMIB_Hosttimetable_Hosttimeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected host statistics of which this
    // entry is a part.  The set of hosts identified by a particular value of this
    // index is associated with the hostControlEntry as identified by the same
    // value of hostControlIndex. The type is interface{} with range: 1..65535.
    Hosttimeindex interface{}

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
    Hosttimecreationorder interface{}

    // The physical address of this host. The type is string.
    Hosttimeaddress interface{}

    // The number of good packets transmitted to this address since it was added
    // to the hostTimeTable. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Hosttimeinpkts interface{}

    // The number of packets, including bad packets, transmitted by this address
    // since it was added to the hostTimeTable. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    Hosttimeoutpkts interface{}

    // The number of octets transmitted to this address since it was added to the
    // hostTimeTable (excluding framing bits but including FCS octets), except for
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Hosttimeinoctets interface{}

    // The number of octets transmitted by this address since it was added to the
    // hostTimeTable (excluding framing bits but including FCS octets), including
    // those octets in bad packets. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Hosttimeoutoctets interface{}

    // The number of bad packets transmitted by this address since this host was
    // added to the hostTimeTable. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Hosttimeouterrors interface{}

    // The number of good packets transmitted by this address that were directed
    // to the broadcast address since this host was added to the hostTimeTable.
    // The type is interface{} with range: 0..4294967295. Units are Packets.
    Hosttimeoutbroadcastpkts interface{}

    // The number of good packets transmitted by this address that were directed
    // to a multicast address since this host was added to the hostTimeTable. Note
    // that this number does not include packets directed to the broadcast
    // address. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Hosttimeoutmulticastpkts interface{}
}

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetEntityData() *types.CommonEntityData {
    hosttimeentry.EntityData.YFilter = hosttimeentry.YFilter
    hosttimeentry.EntityData.YangName = "hostTimeEntry"
    hosttimeentry.EntityData.BundleName = "cisco_ios_xe"
    hosttimeentry.EntityData.ParentYangName = "hostTimeTable"
    hosttimeentry.EntityData.SegmentPath = "hostTimeEntry" + "[hostTimeIndex='" + fmt.Sprintf("%v", hosttimeentry.Hosttimeindex) + "']" + "[hostTimeCreationOrder='" + fmt.Sprintf("%v", hosttimeentry.Hosttimecreationorder) + "']"
    hosttimeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttimeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttimeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttimeentry.EntityData.Children = make(map[string]types.YChild)
    hosttimeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hosttimeentry.EntityData.Leafs["hostTimeIndex"] = types.YLeaf{"Hosttimeindex", hosttimeentry.Hosttimeindex}
    hosttimeentry.EntityData.Leafs["hostTimeCreationOrder"] = types.YLeaf{"Hosttimecreationorder", hosttimeentry.Hosttimecreationorder}
    hosttimeentry.EntityData.Leafs["hostTimeAddress"] = types.YLeaf{"Hosttimeaddress", hosttimeentry.Hosttimeaddress}
    hosttimeentry.EntityData.Leafs["hostTimeInPkts"] = types.YLeaf{"Hosttimeinpkts", hosttimeentry.Hosttimeinpkts}
    hosttimeentry.EntityData.Leafs["hostTimeOutPkts"] = types.YLeaf{"Hosttimeoutpkts", hosttimeentry.Hosttimeoutpkts}
    hosttimeentry.EntityData.Leafs["hostTimeInOctets"] = types.YLeaf{"Hosttimeinoctets", hosttimeentry.Hosttimeinoctets}
    hosttimeentry.EntityData.Leafs["hostTimeOutOctets"] = types.YLeaf{"Hosttimeoutoctets", hosttimeentry.Hosttimeoutoctets}
    hosttimeentry.EntityData.Leafs["hostTimeOutErrors"] = types.YLeaf{"Hosttimeouterrors", hosttimeentry.Hosttimeouterrors}
    hosttimeentry.EntityData.Leafs["hostTimeOutBroadcastPkts"] = types.YLeaf{"Hosttimeoutbroadcastpkts", hosttimeentry.Hosttimeoutbroadcastpkts}
    hosttimeentry.EntityData.Leafs["hostTimeOutMulticastPkts"] = types.YLeaf{"Hosttimeoutmulticastpkts", hosttimeentry.Hosttimeoutmulticastpkts}
    return &(hosttimeentry.EntityData)
}

// RMONMIB_Hosttopncontroltable
// A list of top N host control entries.
type RMONMIB_Hosttopncontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters that control the creation of a report of the top N
    // hosts according to several metrics.  For example, an instance of the
    // hostTopNDuration object might be named hostTopNDuration.3. The type is
    // slice of RMONMIB_Hosttopncontroltable_Hosttopncontrolentry.
    Hosttopncontrolentry []RMONMIB_Hosttopncontroltable_Hosttopncontrolentry
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetEntityData() *types.CommonEntityData {
    hosttopncontroltable.EntityData.YFilter = hosttopncontroltable.YFilter
    hosttopncontroltable.EntityData.YangName = "hostTopNControlTable"
    hosttopncontroltable.EntityData.BundleName = "cisco_ios_xe"
    hosttopncontroltable.EntityData.ParentYangName = "RMON-MIB"
    hosttopncontroltable.EntityData.SegmentPath = "hostTopNControlTable"
    hosttopncontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttopncontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttopncontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttopncontroltable.EntityData.Children = make(map[string]types.YChild)
    hosttopncontroltable.EntityData.Children["hostTopNControlEntry"] = types.YChild{"Hosttopncontrolentry", nil}
    for i := range hosttopncontroltable.Hosttopncontrolentry {
        hosttopncontroltable.EntityData.Children[types.GetSegmentPath(&hosttopncontroltable.Hosttopncontrolentry[i])] = types.YChild{"Hosttopncontrolentry", &hosttopncontroltable.Hosttopncontrolentry[i]}
    }
    hosttopncontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosttopncontroltable.EntityData)
}

// RMONMIB_Hosttopncontroltable_Hosttopncontrolentry
// A set of parameters that control the creation of a report
// of the top N hosts according to several metrics.  For
// example, an instance of the hostTopNDuration object might
// be named hostTopNDuration.3
type RMONMIB_Hosttopncontroltable_Hosttopncontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostTopNControl table.  Each such entry defines one top N report prepared
    // for one interface. The type is interface{} with range: 1..65535.
    Hosttopncontrolindex interface{}

    // The host table for which a top N report will be prepared on behalf of this
    // entry.  The host table identified by a particular value of this index is
    // associated with the same host table as identified by the same value of
    // hostIndex.  This object may not be modified if the associated
    // hostTopNStatus object is equal to valid(1). The type is interface{} with
    // range: 1..65535.
    Hosttopnhostindex interface{}

    // The variable for each host that the hostTopNRate variable is based upon. 
    // This object may not be modified if the associated hostTopNStatus object is
    // equal to valid(1). The type is Hosttopnratebase.
    Hosttopnratebase interface{}

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
    Hosttopntimeremaining interface{}

    // The number of seconds that this report has collected during the last
    // sampling interval, or if this report is currently being collected, the
    // number of seconds that this report is being collected during this sampling
    // interval.  When the associated hostTopNTimeRemaining object is set, this
    // object shall be set by the probe to the same value and shall not be
    // modified until the next time the hostTopNTimeRemaining is set.  This value
    // shall be zero if no reports have been requested for this
    // hostTopNControlEntry. The type is interface{} with range:
    // -2147483648..2147483647. Units are Seconds.
    Hosttopnduration interface{}

    // The maximum number of hosts requested for the top N table.  When this
    // object is created or modified, the probe should set hostTopNGrantedSize as
    // closely to this object as is possible for the particular probe
    // implementation and available resources. The type is interface{} with range:
    // -2147483648..2147483647.
    Hosttopnrequestedsize interface{}

    // The maximum number of hosts in the top N table.  When the associated
    // hostTopNRequestedSize object is created or modified, the probe should set
    // this object as closely to the requested value as is possible for the
    // particular implementation and available resources. The probe must not lower
    // this value except as a result of a set to the associated
    // hostTopNRequestedSize object.  Hosts with the highest value of hostTopNRate
    // shall be placed in this table in decreasing order of this rate until there
    // is no more room or until there are no more hosts. The type is interface{}
    // with range: -2147483648..2147483647.
    Hosttopngrantedsize interface{}

    // The value of sysUpTime when this top N report was last started.  In other
    // words, this is the time that the associated hostTopNTimeRemaining object
    // was modified to start the requested report. The type is interface{} with
    // range: 0..4294967295.
    Hosttopnstarttime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Hosttopnowner interface{}

    // The status of this hostTopNControl entry.  If this object is not equal to
    // valid(1), all associated hostTopNEntries shall be deleted by the agent. The
    // type is EntryStatus.
    Hosttopnstatus interface{}
}

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetEntityData() *types.CommonEntityData {
    hosttopncontrolentry.EntityData.YFilter = hosttopncontrolentry.YFilter
    hosttopncontrolentry.EntityData.YangName = "hostTopNControlEntry"
    hosttopncontrolentry.EntityData.BundleName = "cisco_ios_xe"
    hosttopncontrolentry.EntityData.ParentYangName = "hostTopNControlTable"
    hosttopncontrolentry.EntityData.SegmentPath = "hostTopNControlEntry" + "[hostTopNControlIndex='" + fmt.Sprintf("%v", hosttopncontrolentry.Hosttopncontrolindex) + "']"
    hosttopncontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttopncontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttopncontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttopncontrolentry.EntityData.Children = make(map[string]types.YChild)
    hosttopncontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hosttopncontrolentry.EntityData.Leafs["hostTopNControlIndex"] = types.YLeaf{"Hosttopncontrolindex", hosttopncontrolentry.Hosttopncontrolindex}
    hosttopncontrolentry.EntityData.Leafs["hostTopNHostIndex"] = types.YLeaf{"Hosttopnhostindex", hosttopncontrolentry.Hosttopnhostindex}
    hosttopncontrolentry.EntityData.Leafs["hostTopNRateBase"] = types.YLeaf{"Hosttopnratebase", hosttopncontrolentry.Hosttopnratebase}
    hosttopncontrolentry.EntityData.Leafs["hostTopNTimeRemaining"] = types.YLeaf{"Hosttopntimeremaining", hosttopncontrolentry.Hosttopntimeremaining}
    hosttopncontrolentry.EntityData.Leafs["hostTopNDuration"] = types.YLeaf{"Hosttopnduration", hosttopncontrolentry.Hosttopnduration}
    hosttopncontrolentry.EntityData.Leafs["hostTopNRequestedSize"] = types.YLeaf{"Hosttopnrequestedsize", hosttopncontrolentry.Hosttopnrequestedsize}
    hosttopncontrolentry.EntityData.Leafs["hostTopNGrantedSize"] = types.YLeaf{"Hosttopngrantedsize", hosttopncontrolentry.Hosttopngrantedsize}
    hosttopncontrolentry.EntityData.Leafs["hostTopNStartTime"] = types.YLeaf{"Hosttopnstarttime", hosttopncontrolentry.Hosttopnstarttime}
    hosttopncontrolentry.EntityData.Leafs["hostTopNOwner"] = types.YLeaf{"Hosttopnowner", hosttopncontrolentry.Hosttopnowner}
    hosttopncontrolentry.EntityData.Leafs["hostTopNStatus"] = types.YLeaf{"Hosttopnstatus", hosttopncontrolentry.Hosttopnstatus}
    return &(hosttopncontrolentry.EntityData)
}

// RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase represents hostTopNStatus object is equal to valid(1).
type RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase string

const (
    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNInPkts RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNInPkts"

    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNOutPkts RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNOutPkts"

    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNInOctets RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNInOctets"

    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNOutOctets RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNOutOctets"

    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNOutErrors RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNOutErrors"

    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNOutBroadcastPkts RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNOutBroadcastPkts"

    RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase_hostTopNOutMulticastPkts RMONMIB_Hosttopncontroltable_Hosttopncontrolentry_Hosttopnratebase = "hostTopNOutMulticastPkts"
)

// RMONMIB_Hosttopntable
// A list of top N host entries.
type RMONMIB_Hosttopntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of statistics for a host that is part of a top N report.  For
    // example, an instance of the hostTopNRate object might be named
    // hostTopNRate.3.10. The type is slice of
    // RMONMIB_Hosttopntable_Hosttopnentry.
    Hosttopnentry []RMONMIB_Hosttopntable_Hosttopnentry
}

func (hosttopntable *RMONMIB_Hosttopntable) GetEntityData() *types.CommonEntityData {
    hosttopntable.EntityData.YFilter = hosttopntable.YFilter
    hosttopntable.EntityData.YangName = "hostTopNTable"
    hosttopntable.EntityData.BundleName = "cisco_ios_xe"
    hosttopntable.EntityData.ParentYangName = "RMON-MIB"
    hosttopntable.EntityData.SegmentPath = "hostTopNTable"
    hosttopntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttopntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttopntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttopntable.EntityData.Children = make(map[string]types.YChild)
    hosttopntable.EntityData.Children["hostTopNEntry"] = types.YChild{"Hosttopnentry", nil}
    for i := range hosttopntable.Hosttopnentry {
        hosttopntable.EntityData.Children[types.GetSegmentPath(&hosttopntable.Hosttopnentry[i])] = types.YChild{"Hosttopnentry", &hosttopntable.Hosttopnentry[i]}
    }
    hosttopntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosttopntable.EntityData)
}

// RMONMIB_Hosttopntable_Hosttopnentry
// A set of statistics for a host that is part of a top N
// report.  For example, an instance of the hostTopNRate
// object might be named hostTopNRate.3.10
type RMONMIB_Hosttopntable_Hosttopnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object identifies the top N report of which
    // this entry is a part.  The set of hosts identified by a particular value of
    // this object is part of the same report as identified by the same value of
    // the hostTopNControlIndex object. The type is interface{} with range:
    // 1..65535.
    Hosttopnreport interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // hostTopN table among those in the same report. This index is between 1 and
    // N, where N is the number of entries in this table.  Increasing values of
    // hostTopNIndex shall be assigned to entries with decreasing values of
    // hostTopNRate until index N is assigned to the entry with the lowest value
    // of hostTopNRate or there are no more hostTopNEntries. The type is
    // interface{} with range: 1..65535.
    Hosttopnindex interface{}

    // The physical address of this host. The type is string.
    Hosttopnaddress interface{}

    // The amount of change in the selected variable during this sampling
    // interval.  The selected variable is this host's instance of the object
    // selected by hostTopNRateBase. The type is interface{} with range:
    // -2147483648..2147483647.
    Hosttopnrate interface{}
}

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetEntityData() *types.CommonEntityData {
    hosttopnentry.EntityData.YFilter = hosttopnentry.YFilter
    hosttopnentry.EntityData.YangName = "hostTopNEntry"
    hosttopnentry.EntityData.BundleName = "cisco_ios_xe"
    hosttopnentry.EntityData.ParentYangName = "hostTopNTable"
    hosttopnentry.EntityData.SegmentPath = "hostTopNEntry" + "[hostTopNReport='" + fmt.Sprintf("%v", hosttopnentry.Hosttopnreport) + "']" + "[hostTopNIndex='" + fmt.Sprintf("%v", hosttopnentry.Hosttopnindex) + "']"
    hosttopnentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    hosttopnentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    hosttopnentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    hosttopnentry.EntityData.Children = make(map[string]types.YChild)
    hosttopnentry.EntityData.Leafs = make(map[string]types.YLeaf)
    hosttopnentry.EntityData.Leafs["hostTopNReport"] = types.YLeaf{"Hosttopnreport", hosttopnentry.Hosttopnreport}
    hosttopnentry.EntityData.Leafs["hostTopNIndex"] = types.YLeaf{"Hosttopnindex", hosttopnentry.Hosttopnindex}
    hosttopnentry.EntityData.Leafs["hostTopNAddress"] = types.YLeaf{"Hosttopnaddress", hosttopnentry.Hosttopnaddress}
    hosttopnentry.EntityData.Leafs["hostTopNRate"] = types.YLeaf{"Hosttopnrate", hosttopnentry.Hosttopnrate}
    return &(hosttopnentry.EntityData)
}

// RMONMIB_Matrixcontroltable
// A list of information entries for the
// traffic matrix on each interface.
type RMONMIB_Matrixcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a traffic matrix on a particular interface.  For example,
    // an instance of the matrixControlLastDeleteTime object might be named
    // matrixControlLastDeleteTime.1. The type is slice of
    // RMONMIB_Matrixcontroltable_Matrixcontrolentry.
    Matrixcontrolentry []RMONMIB_Matrixcontroltable_Matrixcontrolentry
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetEntityData() *types.CommonEntityData {
    matrixcontroltable.EntityData.YFilter = matrixcontroltable.YFilter
    matrixcontroltable.EntityData.YangName = "matrixControlTable"
    matrixcontroltable.EntityData.BundleName = "cisco_ios_xe"
    matrixcontroltable.EntityData.ParentYangName = "RMON-MIB"
    matrixcontroltable.EntityData.SegmentPath = "matrixControlTable"
    matrixcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixcontroltable.EntityData.Children = make(map[string]types.YChild)
    matrixcontroltable.EntityData.Children["matrixControlEntry"] = types.YChild{"Matrixcontrolentry", nil}
    for i := range matrixcontroltable.Matrixcontrolentry {
        matrixcontroltable.EntityData.Children[types.GetSegmentPath(&matrixcontroltable.Matrixcontrolentry[i])] = types.YChild{"Matrixcontrolentry", &matrixcontroltable.Matrixcontrolentry[i]}
    }
    matrixcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matrixcontroltable.EntityData)
}

// RMONMIB_Matrixcontroltable_Matrixcontrolentry
// Information about a traffic matrix on a particular
// interface.  For example, an instance of the
// matrixControlLastDeleteTime object might be named
// matrixControlLastDeleteTime.1
type RMONMIB_Matrixcontroltable_Matrixcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // matrixControl table.  Each such entry defines a function that discovers
    // conversations on a particular interface and places statistics about them in
    // the matrixSDTable and the matrixDSTable on behalf of this
    // matrixControlEntry. The type is interface{} with range: 1..65535.
    Matrixcontrolindex interface{}

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
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Matrixcontroldatasource interface{}

    // The number of matrixSDEntries in the matrixSDTable for this interface. 
    // This must also be the value of the number of entries in the matrixDSTable
    // for this interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Matrixcontroltablesize interface{}

    // The value of sysUpTime when the last entry was deleted from the portion of
    // the matrixSDTable or matrixDSTable associated with this matrixControlEntry.
    // If no deletions have occurred, this value shall be zero. The type is
    // interface{} with range: 0..4294967295.
    Matrixcontrollastdeletetime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Matrixcontrolowner interface{}

    // The status of this matrixControl entry. If this object is not equal to
    // valid(1), all associated entries in the matrixSDTable and the matrixDSTable
    // shall be deleted by the agent. The type is EntryStatus.
    Matrixcontrolstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe is out of some resources and decides to shed load
    // from this collection.  This count does not include packets that were not
    // counted      because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Matrixcontroldroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Matrixcontrolcreatetime interface{}
}

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetEntityData() *types.CommonEntityData {
    matrixcontrolentry.EntityData.YFilter = matrixcontrolentry.YFilter
    matrixcontrolentry.EntityData.YangName = "matrixControlEntry"
    matrixcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    matrixcontrolentry.EntityData.ParentYangName = "matrixControlTable"
    matrixcontrolentry.EntityData.SegmentPath = "matrixControlEntry" + "[matrixControlIndex='" + fmt.Sprintf("%v", matrixcontrolentry.Matrixcontrolindex) + "']"
    matrixcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixcontrolentry.EntityData.Children = make(map[string]types.YChild)
    matrixcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    matrixcontrolentry.EntityData.Leafs["matrixControlIndex"] = types.YLeaf{"Matrixcontrolindex", matrixcontrolentry.Matrixcontrolindex}
    matrixcontrolentry.EntityData.Leafs["matrixControlDataSource"] = types.YLeaf{"Matrixcontroldatasource", matrixcontrolentry.Matrixcontroldatasource}
    matrixcontrolentry.EntityData.Leafs["matrixControlTableSize"] = types.YLeaf{"Matrixcontroltablesize", matrixcontrolentry.Matrixcontroltablesize}
    matrixcontrolentry.EntityData.Leafs["matrixControlLastDeleteTime"] = types.YLeaf{"Matrixcontrollastdeletetime", matrixcontrolentry.Matrixcontrollastdeletetime}
    matrixcontrolentry.EntityData.Leafs["matrixControlOwner"] = types.YLeaf{"Matrixcontrolowner", matrixcontrolentry.Matrixcontrolowner}
    matrixcontrolentry.EntityData.Leafs["matrixControlStatus"] = types.YLeaf{"Matrixcontrolstatus", matrixcontrolentry.Matrixcontrolstatus}
    matrixcontrolentry.EntityData.Leafs["matrixControlDroppedFrames"] = types.YLeaf{"Matrixcontroldroppedframes", matrixcontrolentry.Matrixcontroldroppedframes}
    matrixcontrolentry.EntityData.Leafs["matrixControlCreateTime"] = types.YLeaf{"Matrixcontrolcreatetime", matrixcontrolentry.Matrixcontrolcreatetime}
    return &(matrixcontrolentry.EntityData)
}

// RMONMIB_Matrixsdtable
// A list of traffic matrix entries indexed by
// source and destination MAC address.
type RMONMIB_Matrixsdtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for communications between two addresses on a
    // particular interface.  For example, an instance of the matrixSDPkts object
    // might be named matrixSDPkts.1.6.8.0.32.27.3.176.6.8.0.32.10.8.113. The type
    // is slice of RMONMIB_Matrixsdtable_Matrixsdentry.
    Matrixsdentry []RMONMIB_Matrixsdtable_Matrixsdentry
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetEntityData() *types.CommonEntityData {
    matrixsdtable.EntityData.YFilter = matrixsdtable.YFilter
    matrixsdtable.EntityData.YangName = "matrixSDTable"
    matrixsdtable.EntityData.BundleName = "cisco_ios_xe"
    matrixsdtable.EntityData.ParentYangName = "RMON-MIB"
    matrixsdtable.EntityData.SegmentPath = "matrixSDTable"
    matrixsdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixsdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixsdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixsdtable.EntityData.Children = make(map[string]types.YChild)
    matrixsdtable.EntityData.Children["matrixSDEntry"] = types.YChild{"Matrixsdentry", nil}
    for i := range matrixsdtable.Matrixsdentry {
        matrixsdtable.EntityData.Children[types.GetSegmentPath(&matrixsdtable.Matrixsdentry[i])] = types.YChild{"Matrixsdentry", &matrixsdtable.Matrixsdentry[i]}
    }
    matrixsdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matrixsdtable.EntityData)
}

// RMONMIB_Matrixsdtable_Matrixsdentry
// A collection of statistics for communications between
// two addresses on a particular interface.  For example,
// an instance of the matrixSDPkts object might be named
// matrixSDPkts.1.6.8.0.32.27.3.176.6.8.0.32.10.8.113
type RMONMIB_Matrixsdtable_Matrixsdentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected matrix statistics of which
    // this entry is a part.  The set of matrix statistics identified by a
    // particular value of this index is associated with the same
    // matrixControlEntry as identified by the same value of matrixControlIndex.
    // The type is interface{} with range: 1..65535.
    Matrixsdindex interface{}

    // This attribute is a key. The source physical address. The type is string.
    Matrixsdsourceaddress interface{}

    // This attribute is a key. The destination physical address. The type is
    // string.
    Matrixsddestaddress interface{}

    // The number of packets transmitted from the source address to the
    // destination address (this number includes bad packets). The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Matrixsdpkts interface{}

    // The number of octets (excluding framing bits but including FCS octets)
    // contained in all packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    Matrixsdoctets interface{}

    // The number of bad packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Matrixsderrors interface{}
}

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetEntityData() *types.CommonEntityData {
    matrixsdentry.EntityData.YFilter = matrixsdentry.YFilter
    matrixsdentry.EntityData.YangName = "matrixSDEntry"
    matrixsdentry.EntityData.BundleName = "cisco_ios_xe"
    matrixsdentry.EntityData.ParentYangName = "matrixSDTable"
    matrixsdentry.EntityData.SegmentPath = "matrixSDEntry" + "[matrixSDIndex='" + fmt.Sprintf("%v", matrixsdentry.Matrixsdindex) + "']" + "[matrixSDSourceAddress='" + fmt.Sprintf("%v", matrixsdentry.Matrixsdsourceaddress) + "']" + "[matrixSDDestAddress='" + fmt.Sprintf("%v", matrixsdentry.Matrixsddestaddress) + "']"
    matrixsdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixsdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixsdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixsdentry.EntityData.Children = make(map[string]types.YChild)
    matrixsdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    matrixsdentry.EntityData.Leafs["matrixSDIndex"] = types.YLeaf{"Matrixsdindex", matrixsdentry.Matrixsdindex}
    matrixsdentry.EntityData.Leafs["matrixSDSourceAddress"] = types.YLeaf{"Matrixsdsourceaddress", matrixsdentry.Matrixsdsourceaddress}
    matrixsdentry.EntityData.Leafs["matrixSDDestAddress"] = types.YLeaf{"Matrixsddestaddress", matrixsdentry.Matrixsddestaddress}
    matrixsdentry.EntityData.Leafs["matrixSDPkts"] = types.YLeaf{"Matrixsdpkts", matrixsdentry.Matrixsdpkts}
    matrixsdentry.EntityData.Leafs["matrixSDOctets"] = types.YLeaf{"Matrixsdoctets", matrixsdentry.Matrixsdoctets}
    matrixsdentry.EntityData.Leafs["matrixSDErrors"] = types.YLeaf{"Matrixsderrors", matrixsdentry.Matrixsderrors}
    return &(matrixsdentry.EntityData)
}

// RMONMIB_Matrixdstable
// A list of traffic matrix entries indexed by
// destination and source MAC address.
type RMONMIB_Matrixdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of statistics for communications between two addresses on a
    // particular interface.  For example, an instance of the matrixSDPkts object
    // might be named matrixSDPkts.1.6.8.0.32.10.8.113.6.8.0.32.27.3.176. The type
    // is slice of RMONMIB_Matrixdstable_Matrixdsentry.
    Matrixdsentry []RMONMIB_Matrixdstable_Matrixdsentry
}

func (matrixdstable *RMONMIB_Matrixdstable) GetEntityData() *types.CommonEntityData {
    matrixdstable.EntityData.YFilter = matrixdstable.YFilter
    matrixdstable.EntityData.YangName = "matrixDSTable"
    matrixdstable.EntityData.BundleName = "cisco_ios_xe"
    matrixdstable.EntityData.ParentYangName = "RMON-MIB"
    matrixdstable.EntityData.SegmentPath = "matrixDSTable"
    matrixdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixdstable.EntityData.Children = make(map[string]types.YChild)
    matrixdstable.EntityData.Children["matrixDSEntry"] = types.YChild{"Matrixdsentry", nil}
    for i := range matrixdstable.Matrixdsentry {
        matrixdstable.EntityData.Children[types.GetSegmentPath(&matrixdstable.Matrixdsentry[i])] = types.YChild{"Matrixdsentry", &matrixdstable.Matrixdsentry[i]}
    }
    matrixdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matrixdstable.EntityData)
}

// RMONMIB_Matrixdstable_Matrixdsentry
// A collection of statistics for communications between
// two addresses on a particular interface.  For example,
// an instance of the matrixSDPkts object might be named
// matrixSDPkts.1.6.8.0.32.10.8.113.6.8.0.32.27.3.176
type RMONMIB_Matrixdstable_Matrixdsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The set of collected matrix statistics of which
    // this entry is a part.  The set of matrix statistics identified by a
    // particular value of this index is associated with the same
    // matrixControlEntry as identified by the same value of matrixControlIndex.
    // The type is interface{} with range: 1..65535.
    Matrixdsindex interface{}

    // This attribute is a key. The destination physical address. The type is
    // string.
    Matrixdsdestaddress interface{}

    // This attribute is a key. The source physical address. The type is string.
    Matrixdssourceaddress interface{}

    // The number of packets transmitted from the source address to the
    // destination address (this number includes bad packets). The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Matrixdspkts interface{}

    // The number of octets (excluding framing bits but including FCS octets)
    // contained in all packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    Matrixdsoctets interface{}

    // The number of bad packets transmitted from the source address to the
    // destination address. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Matrixdserrors interface{}
}

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetEntityData() *types.CommonEntityData {
    matrixdsentry.EntityData.YFilter = matrixdsentry.YFilter
    matrixdsentry.EntityData.YangName = "matrixDSEntry"
    matrixdsentry.EntityData.BundleName = "cisco_ios_xe"
    matrixdsentry.EntityData.ParentYangName = "matrixDSTable"
    matrixdsentry.EntityData.SegmentPath = "matrixDSEntry" + "[matrixDSIndex='" + fmt.Sprintf("%v", matrixdsentry.Matrixdsindex) + "']" + "[matrixDSDestAddress='" + fmt.Sprintf("%v", matrixdsentry.Matrixdsdestaddress) + "']" + "[matrixDSSourceAddress='" + fmt.Sprintf("%v", matrixdsentry.Matrixdssourceaddress) + "']"
    matrixdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    matrixdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    matrixdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    matrixdsentry.EntityData.Children = make(map[string]types.YChild)
    matrixdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    matrixdsentry.EntityData.Leafs["matrixDSIndex"] = types.YLeaf{"Matrixdsindex", matrixdsentry.Matrixdsindex}
    matrixdsentry.EntityData.Leafs["matrixDSDestAddress"] = types.YLeaf{"Matrixdsdestaddress", matrixdsentry.Matrixdsdestaddress}
    matrixdsentry.EntityData.Leafs["matrixDSSourceAddress"] = types.YLeaf{"Matrixdssourceaddress", matrixdsentry.Matrixdssourceaddress}
    matrixdsentry.EntityData.Leafs["matrixDSPkts"] = types.YLeaf{"Matrixdspkts", matrixdsentry.Matrixdspkts}
    matrixdsentry.EntityData.Leafs["matrixDSOctets"] = types.YLeaf{"Matrixdsoctets", matrixdsentry.Matrixdsoctets}
    matrixdsentry.EntityData.Leafs["matrixDSErrors"] = types.YLeaf{"Matrixdserrors", matrixdsentry.Matrixdserrors}
    return &(matrixdsentry.EntityData)
}

// RMONMIB_Filtertable
// A list of packet filter entries.
type RMONMIB_Filtertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters for a packet filter applied on a particular interface. 
    // As an example, an instance of the filterPktData object might be named
    // filterPktData.12. The type is slice of RMONMIB_Filtertable_Filterentry.
    Filterentry []RMONMIB_Filtertable_Filterentry
}

func (filtertable *RMONMIB_Filtertable) GetEntityData() *types.CommonEntityData {
    filtertable.EntityData.YFilter = filtertable.YFilter
    filtertable.EntityData.YangName = "filterTable"
    filtertable.EntityData.BundleName = "cisco_ios_xe"
    filtertable.EntityData.ParentYangName = "RMON-MIB"
    filtertable.EntityData.SegmentPath = "filterTable"
    filtertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    filtertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    filtertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    filtertable.EntityData.Children = make(map[string]types.YChild)
    filtertable.EntityData.Children["filterEntry"] = types.YChild{"Filterentry", nil}
    for i := range filtertable.Filterentry {
        filtertable.EntityData.Children[types.GetSegmentPath(&filtertable.Filterentry[i])] = types.YChild{"Filterentry", &filtertable.Filterentry[i]}
    }
    filtertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(filtertable.EntityData)
}

// RMONMIB_Filtertable_Filterentry
// A set of parameters for a packet filter applied on a
// particular interface.  As an example, an instance of the
// filterPktData object might be named filterPktData.12
type RMONMIB_Filtertable_Filterentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // filter table.  Each such entry defines one filter that is to be applied to
    // every packet received on an interface. The type is interface{} with range:
    // 1..65535.
    Filterindex interface{}

    // This object identifies the channel of which this filter is a part.  The
    // filters identified by a particular value of this object are associated with
    // the same channel as identified by the same value of the channelIndex
    // object. The type is interface{} with range: 1..65535.
    Filterchannelindex interface{}

    // The offset from the beginning of each packet where a match of packet data
    // will be attempted.  This offset is measured from the point in the physical
    // layer packet after the framing bits, if any.  For example, in an Ethernet
    // frame, this point is at the beginning of the destination MAC address.  This
    // object may not be modified if the associated filterStatus object is equal
    // to valid(1). The type is interface{} with range: -2147483648..2147483647.
    // Units are Octets.
    Filterpktdataoffset interface{}

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
    Filterpktdata interface{}

    // The mask that is applied to the match process. After adjusting this mask
    // for the offset, only those bits in the received packet that correspond to
    // bits set in this mask are relevant for further processing by the match
    // algorithm.  The offset is applied to filterPktDataMask in the same way it
    // is applied to the filter.  For the purposes of the matching algorithm, if
    // the associated filterPktData object is longer than this mask, this mask is
    // conceptually extended with '1' bits until it reaches the length of the
    // filterPktData object.  This object may not be modified if the associated
    // filterStatus object is equal to valid(1). The type is string.
    Filterpktdatamask interface{}

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
    Filterpktdatanotmask interface{}

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
    Filterpktstatus interface{}

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
    Filterpktstatusmask interface{}

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
    Filterpktstatusnotmask interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Filterowner interface{}

    // The status of this filter entry. The type is EntryStatus.
    Filterstatus interface{}

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
    Filterprotocoldirdatalocalindex interface{}

    // When this object is set to a non-zero value, the filter that it is
    // associated with will discard the packet if the packet doesn't match this
    // protocol directory entry. The type is interface{} with range:
    // 0..2147483647.
    Filterprotocoldirlocalindex interface{}
}

func (filterentry *RMONMIB_Filtertable_Filterentry) GetEntityData() *types.CommonEntityData {
    filterentry.EntityData.YFilter = filterentry.YFilter
    filterentry.EntityData.YangName = "filterEntry"
    filterentry.EntityData.BundleName = "cisco_ios_xe"
    filterentry.EntityData.ParentYangName = "filterTable"
    filterentry.EntityData.SegmentPath = "filterEntry" + "[filterIndex='" + fmt.Sprintf("%v", filterentry.Filterindex) + "']"
    filterentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    filterentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    filterentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    filterentry.EntityData.Children = make(map[string]types.YChild)
    filterentry.EntityData.Leafs = make(map[string]types.YLeaf)
    filterentry.EntityData.Leafs["filterIndex"] = types.YLeaf{"Filterindex", filterentry.Filterindex}
    filterentry.EntityData.Leafs["filterChannelIndex"] = types.YLeaf{"Filterchannelindex", filterentry.Filterchannelindex}
    filterentry.EntityData.Leafs["filterPktDataOffset"] = types.YLeaf{"Filterpktdataoffset", filterentry.Filterpktdataoffset}
    filterentry.EntityData.Leafs["filterPktData"] = types.YLeaf{"Filterpktdata", filterentry.Filterpktdata}
    filterentry.EntityData.Leafs["filterPktDataMask"] = types.YLeaf{"Filterpktdatamask", filterentry.Filterpktdatamask}
    filterentry.EntityData.Leafs["filterPktDataNotMask"] = types.YLeaf{"Filterpktdatanotmask", filterentry.Filterpktdatanotmask}
    filterentry.EntityData.Leafs["filterPktStatus"] = types.YLeaf{"Filterpktstatus", filterentry.Filterpktstatus}
    filterentry.EntityData.Leafs["filterPktStatusMask"] = types.YLeaf{"Filterpktstatusmask", filterentry.Filterpktstatusmask}
    filterentry.EntityData.Leafs["filterPktStatusNotMask"] = types.YLeaf{"Filterpktstatusnotmask", filterentry.Filterpktstatusnotmask}
    filterentry.EntityData.Leafs["filterOwner"] = types.YLeaf{"Filterowner", filterentry.Filterowner}
    filterentry.EntityData.Leafs["filterStatus"] = types.YLeaf{"Filterstatus", filterentry.Filterstatus}
    filterentry.EntityData.Leafs["filterProtocolDirDataLocalIndex"] = types.YLeaf{"Filterprotocoldirdatalocalindex", filterentry.Filterprotocoldirdatalocalindex}
    filterentry.EntityData.Leafs["filterProtocolDirLocalIndex"] = types.YLeaf{"Filterprotocoldirlocalindex", filterentry.Filterprotocoldirlocalindex}
    return &(filterentry.EntityData)
}

// RMONMIB_Channeltable
// A list of packet channel entries.
type RMONMIB_Channeltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters for a packet channel applied on a particular interface.
    // As an example, an instance of the channelMatches object might be named
    // channelMatches.3. The type is slice of RMONMIB_Channeltable_Channelentry.
    Channelentry []RMONMIB_Channeltable_Channelentry
}

func (channeltable *RMONMIB_Channeltable) GetEntityData() *types.CommonEntityData {
    channeltable.EntityData.YFilter = channeltable.YFilter
    channeltable.EntityData.YangName = "channelTable"
    channeltable.EntityData.BundleName = "cisco_ios_xe"
    channeltable.EntityData.ParentYangName = "RMON-MIB"
    channeltable.EntityData.SegmentPath = "channelTable"
    channeltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    channeltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    channeltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    channeltable.EntityData.Children = make(map[string]types.YChild)
    channeltable.EntityData.Children["channelEntry"] = types.YChild{"Channelentry", nil}
    for i := range channeltable.Channelentry {
        channeltable.EntityData.Children[types.GetSegmentPath(&channeltable.Channelentry[i])] = types.YChild{"Channelentry", &channeltable.Channelentry[i]}
    }
    channeltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(channeltable.EntityData)
}

// RMONMIB_Channeltable_Channelentry
// A set of parameters for a packet channel applied on a
// particular interface.  As an example, an instance of the
// channelMatches object might be named channelMatches.3
type RMONMIB_Channeltable_Channelentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // channel table.  Each such entry defines one channel, a logical data and
    // event stream.  It is suggested that before creating a channel, an
    // application should scan all instances of the filterChannelIndex object to
    // make sure that there are no pre-existing filters that would be
    // inadvertently be linked to the channel. The type is interface{} with range:
    // 1..65535.
    Channelindex interface{}

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
    Channelifindex interface{}

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
    // Channelaccepttype.
    Channelaccepttype interface{}

    // This object controls the flow of data through this channel. If this object
    // is on(1), data, status and events flow through this channel.  If this
    // object is off(2), data, status and events will not flow through this
    // channel. The type is Channeldatacontrol.
    Channeldatacontrol interface{}

    // The value of this object identifies the event that is configured to turn
    // the associated channelDataControl from off to on when the event is
    // generated.  The event identified by a particular value of this object is
    // the same event as identified by the same value of the eventIndex object. 
    // If there is no corresponding entry in the eventTable, then no association
    // exists.  In fact, if no event is intended for this channel,
    // channelTurnOnEventIndex must be set to zero, a non-existent event index.
    // This object may not be modified if the associated channelStatus object is
    // equal to valid(1). The type is interface{} with range: 0..65535.
    Channelturnoneventindex interface{}

    // The value of this object identifies the event that is configured to turn
    // the associated channelDataControl from on to off when the event is
    // generated.  The event identified by a particular value of this object is
    // the same event as identified by the same value of the eventIndex object. 
    // If there is no corresponding entry in the eventTable, then no association
    // exists.  In fact, if no event is intended for this channel,
    // channelTurnOffEventIndex must be set to zero, a non-existent event index. 
    // This object may not be modified if the associated channelStatus object is
    // equal to valid(1). The type is interface{} with range: 0..65535.
    Channelturnoffeventindex interface{}

    // The value of this object identifies the event that is configured to be
    // generated when the associated channelDataControl is on and a packet is
    // matched.  The event identified by a particular value of this object is the
    // same event as identified by the same value of the eventIndex object.  If
    // there is no corresponding entry in the eventTable, then no association
    // exists.  In fact, if no event is intended for this channel,
    // channelEventIndex must be set to zero, a non-existent event index.  This
    // object may not be modified if the associated channelStatus object is equal
    // to valid(1). The type is interface{} with range: 0..65535.
    Channeleventindex interface{}

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
    // Channeleventstatus.
    Channeleventstatus interface{}

    // The number of times this channel has matched a packet. Note that this
    // object is updated even when channelDataControl is set to off. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Channelmatches interface{}

    // A comment describing this channel. The type is string with length: 0..127.
    Channeldescription interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Channelowner interface{}

    // The status of this channel entry. The type is EntryStatus.
    Channelstatus interface{}

    // The total number of frames which were received by the probe and therefore
    // not accounted for in the *StatsDropEvents, but for which the probe chose
    // not to count for this entry for whatever reason.  Most often, this event
    // occurs when the probe      is out of some resources and decides to shed
    // load from this collection.  This count does not include packets that were
    // not counted because they had MAC-layer errors.  Note that, unlike the
    // dropEvents counter, this number is the exact number of frames dropped. The
    // type is interface{} with range: 0..4294967295.
    Channeldroppedframes interface{}

    // The value of sysUpTime when this control entry was last activated. This can
    // be used by the management station to ensure that the table has not been
    // deleted and recreated between polls. The type is interface{} with range:
    // 0..4294967295.
    Channelcreatetime interface{}
}

func (channelentry *RMONMIB_Channeltable_Channelentry) GetEntityData() *types.CommonEntityData {
    channelentry.EntityData.YFilter = channelentry.YFilter
    channelentry.EntityData.YangName = "channelEntry"
    channelentry.EntityData.BundleName = "cisco_ios_xe"
    channelentry.EntityData.ParentYangName = "channelTable"
    channelentry.EntityData.SegmentPath = "channelEntry" + "[channelIndex='" + fmt.Sprintf("%v", channelentry.Channelindex) + "']"
    channelentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    channelentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    channelentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    channelentry.EntityData.Children = make(map[string]types.YChild)
    channelentry.EntityData.Leafs = make(map[string]types.YLeaf)
    channelentry.EntityData.Leafs["channelIndex"] = types.YLeaf{"Channelindex", channelentry.Channelindex}
    channelentry.EntityData.Leafs["channelIfIndex"] = types.YLeaf{"Channelifindex", channelentry.Channelifindex}
    channelentry.EntityData.Leafs["channelAcceptType"] = types.YLeaf{"Channelaccepttype", channelentry.Channelaccepttype}
    channelentry.EntityData.Leafs["channelDataControl"] = types.YLeaf{"Channeldatacontrol", channelentry.Channeldatacontrol}
    channelentry.EntityData.Leafs["channelTurnOnEventIndex"] = types.YLeaf{"Channelturnoneventindex", channelentry.Channelturnoneventindex}
    channelentry.EntityData.Leafs["channelTurnOffEventIndex"] = types.YLeaf{"Channelturnoffeventindex", channelentry.Channelturnoffeventindex}
    channelentry.EntityData.Leafs["channelEventIndex"] = types.YLeaf{"Channeleventindex", channelentry.Channeleventindex}
    channelentry.EntityData.Leafs["channelEventStatus"] = types.YLeaf{"Channeleventstatus", channelentry.Channeleventstatus}
    channelentry.EntityData.Leafs["channelMatches"] = types.YLeaf{"Channelmatches", channelentry.Channelmatches}
    channelentry.EntityData.Leafs["channelDescription"] = types.YLeaf{"Channeldescription", channelentry.Channeldescription}
    channelentry.EntityData.Leafs["channelOwner"] = types.YLeaf{"Channelowner", channelentry.Channelowner}
    channelentry.EntityData.Leafs["channelStatus"] = types.YLeaf{"Channelstatus", channelentry.Channelstatus}
    channelentry.EntityData.Leafs["channelDroppedFrames"] = types.YLeaf{"Channeldroppedframes", channelentry.Channeldroppedframes}
    channelentry.EntityData.Leafs["channelCreateTime"] = types.YLeaf{"Channelcreatetime", channelentry.Channelcreatetime}
    return &(channelentry.EntityData)
}

// RMONMIB_Channeltable_Channelentry_Channelaccepttype represents channelStatus object is equal to valid(1).
type RMONMIB_Channeltable_Channelentry_Channelaccepttype string

const (
    RMONMIB_Channeltable_Channelentry_Channelaccepttype_acceptMatched RMONMIB_Channeltable_Channelentry_Channelaccepttype = "acceptMatched"

    RMONMIB_Channeltable_Channelentry_Channelaccepttype_acceptFailed RMONMIB_Channeltable_Channelentry_Channelaccepttype = "acceptFailed"
)

// RMONMIB_Channeltable_Channelentry_Channeldatacontrol represents status and events will not flow through this channel.
type RMONMIB_Channeltable_Channelentry_Channeldatacontrol string

const (
    RMONMIB_Channeltable_Channelentry_Channeldatacontrol_on RMONMIB_Channeltable_Channelentry_Channeldatacontrol = "on"

    RMONMIB_Channeltable_Channelentry_Channeldatacontrol_off RMONMIB_Channeltable_Channelentry_Channeldatacontrol = "off"
)

// RMONMIB_Channeltable_Channelentry_Channeleventstatus represents traffic or other performance problems.
type RMONMIB_Channeltable_Channelentry_Channeleventstatus string

const (
    RMONMIB_Channeltable_Channelentry_Channeleventstatus_eventReady RMONMIB_Channeltable_Channelentry_Channeleventstatus = "eventReady"

    RMONMIB_Channeltable_Channelentry_Channeleventstatus_eventFired RMONMIB_Channeltable_Channelentry_Channeleventstatus = "eventFired"

    RMONMIB_Channeltable_Channelentry_Channeleventstatus_eventAlwaysReady RMONMIB_Channeltable_Channelentry_Channeleventstatus = "eventAlwaysReady"
)

// RMONMIB_Buffercontroltable
// A list of buffers control entries.
type RMONMIB_Buffercontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters that control the collection of a stream of packets that
    // have matched filters.  As an example, an instance of the
    // bufferControlCaptureSliceSize object might be named
    // bufferControlCaptureSliceSize.3. The type is slice of
    // RMONMIB_Buffercontroltable_Buffercontrolentry.
    Buffercontrolentry []RMONMIB_Buffercontroltable_Buffercontrolentry
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetEntityData() *types.CommonEntityData {
    buffercontroltable.EntityData.YFilter = buffercontroltable.YFilter
    buffercontroltable.EntityData.YangName = "bufferControlTable"
    buffercontroltable.EntityData.BundleName = "cisco_ios_xe"
    buffercontroltable.EntityData.ParentYangName = "RMON-MIB"
    buffercontroltable.EntityData.SegmentPath = "bufferControlTable"
    buffercontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    buffercontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    buffercontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    buffercontroltable.EntityData.Children = make(map[string]types.YChild)
    buffercontroltable.EntityData.Children["bufferControlEntry"] = types.YChild{"Buffercontrolentry", nil}
    for i := range buffercontroltable.Buffercontrolentry {
        buffercontroltable.EntityData.Children[types.GetSegmentPath(&buffercontroltable.Buffercontrolentry[i])] = types.YChild{"Buffercontrolentry", &buffercontroltable.Buffercontrolentry[i]}
    }
    buffercontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(buffercontroltable.EntityData)
}

// RMONMIB_Buffercontroltable_Buffercontrolentry
// A set of parameters that control the collection of a stream
// of packets that have matched filters.  As an example, an
// instance of the bufferControlCaptureSliceSize object might
// be named bufferControlCaptureSliceSize.3
type RMONMIB_Buffercontroltable_Buffercontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // bufferControl table.  The value of this index shall never be zero.  Each
    // such entry defines one set of packets that is captured and controlled by
    // one or more filters. The type is interface{} with range: 1..65535.
    Buffercontrolindex interface{}

    // An index that identifies the channel that is the source of packets for this
    // bufferControl table. The channel identified by a particular value of this
    // index is the same as identified by the same value of the channelIndex
    // object.  This object may not be modified if the associated
    // bufferControlStatus object is equal to valid(1). The type is interface{}
    // with range: 1..65535.
    Buffercontrolchannelindex interface{}

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
    // Buffercontrolfullstatus.
    Buffercontrolfullstatus interface{}

    // Controls the action of the buffer when it reaches the full status.  When in
    // the lockWhenFull(1) state and a packet is added to the buffer that fills
    // the buffer, the bufferControlFullStatus will be set to full(2) and this
    // buffer will stop capturing packets. The type is Buffercontrolfullaction.
    Buffercontrolfullaction interface{}

    // The maximum number of octets of each packet that will be saved in this
    // capture buffer. For example, if a 1500 octet packet is received by the
    // probe and this object is set to 500, then only 500 octets of the packet
    // will be stored in the associated capture buffer.  If this variable is set
    // to 0, the capture buffer will save as many octets as is possible.  This
    // object may not be modified if the associated bufferControlStatus object is
    // equal to valid(1). The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    Buffercontrolcaptureslicesize interface{}

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
    Buffercontroldownloadslicesize interface{}

    // The offset of the first octet of each packet in this capture buffer that
    // will be returned in an SNMP retrieval of that packet.  For example, if 500
    // octets of a packet have been stored in the associated capture buffer and
    // this object is set to 100, then the captureBufferPacket object that
    // contains the packet will contain bytes starting 100 octets into the packet.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // Octets.
    Buffercontroldownloadoffset interface{}

    // The requested maximum number of octets to be saved in this captureBuffer,
    // including any implementation-specific overhead. If this variable is set to
    // -1, the capture buffer will save as many octets as is possible.  When this
    // object is created or modified, the probe should set
    // bufferControlMaxOctetsGranted as closely to this object as is possible for
    // the particular probe implementation and available resources.  However, if
    // the object has the special value of -1, the probe must set
    // bufferControlMaxOctetsGranted to -1. The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    Buffercontrolmaxoctetsrequested interface{}

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
    Buffercontrolmaxoctetsgranted interface{}

    // The number of packets currently in this captureBuffer. The type is
    // interface{} with range: -2147483648..2147483647. Units are Packets.
    Buffercontrolcapturedpackets interface{}

    // The value of sysUpTime when this capture buffer was first turned on. The
    // type is interface{} with range: 0..4294967295.
    Buffercontrolturnontime interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..127.
    Buffercontrolowner interface{}

    // The status of this buffer Control Entry. The type is EntryStatus.
    Buffercontrolstatus interface{}
}

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetEntityData() *types.CommonEntityData {
    buffercontrolentry.EntityData.YFilter = buffercontrolentry.YFilter
    buffercontrolentry.EntityData.YangName = "bufferControlEntry"
    buffercontrolentry.EntityData.BundleName = "cisco_ios_xe"
    buffercontrolentry.EntityData.ParentYangName = "bufferControlTable"
    buffercontrolentry.EntityData.SegmentPath = "bufferControlEntry" + "[bufferControlIndex='" + fmt.Sprintf("%v", buffercontrolentry.Buffercontrolindex) + "']"
    buffercontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    buffercontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    buffercontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    buffercontrolentry.EntityData.Children = make(map[string]types.YChild)
    buffercontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    buffercontrolentry.EntityData.Leafs["bufferControlIndex"] = types.YLeaf{"Buffercontrolindex", buffercontrolentry.Buffercontrolindex}
    buffercontrolentry.EntityData.Leafs["bufferControlChannelIndex"] = types.YLeaf{"Buffercontrolchannelindex", buffercontrolentry.Buffercontrolchannelindex}
    buffercontrolentry.EntityData.Leafs["bufferControlFullStatus"] = types.YLeaf{"Buffercontrolfullstatus", buffercontrolentry.Buffercontrolfullstatus}
    buffercontrolentry.EntityData.Leafs["bufferControlFullAction"] = types.YLeaf{"Buffercontrolfullaction", buffercontrolentry.Buffercontrolfullaction}
    buffercontrolentry.EntityData.Leafs["bufferControlCaptureSliceSize"] = types.YLeaf{"Buffercontrolcaptureslicesize", buffercontrolentry.Buffercontrolcaptureslicesize}
    buffercontrolentry.EntityData.Leafs["bufferControlDownloadSliceSize"] = types.YLeaf{"Buffercontroldownloadslicesize", buffercontrolentry.Buffercontroldownloadslicesize}
    buffercontrolentry.EntityData.Leafs["bufferControlDownloadOffset"] = types.YLeaf{"Buffercontroldownloadoffset", buffercontrolentry.Buffercontroldownloadoffset}
    buffercontrolentry.EntityData.Leafs["bufferControlMaxOctetsRequested"] = types.YLeaf{"Buffercontrolmaxoctetsrequested", buffercontrolentry.Buffercontrolmaxoctetsrequested}
    buffercontrolentry.EntityData.Leafs["bufferControlMaxOctetsGranted"] = types.YLeaf{"Buffercontrolmaxoctetsgranted", buffercontrolentry.Buffercontrolmaxoctetsgranted}
    buffercontrolentry.EntityData.Leafs["bufferControlCapturedPackets"] = types.YLeaf{"Buffercontrolcapturedpackets", buffercontrolentry.Buffercontrolcapturedpackets}
    buffercontrolentry.EntityData.Leafs["bufferControlTurnOnTime"] = types.YLeaf{"Buffercontrolturnontime", buffercontrolentry.Buffercontrolturnontime}
    buffercontrolentry.EntityData.Leafs["bufferControlOwner"] = types.YLeaf{"Buffercontrolowner", buffercontrolentry.Buffercontrolowner}
    buffercontrolentry.EntityData.Leafs["bufferControlStatus"] = types.YLeaf{"Buffercontrolstatus", buffercontrolentry.Buffercontrolstatus}
    return &(buffercontrolentry.EntityData)
}

// RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullaction represents packets.
type RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullaction string

const (
    RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullaction_lockWhenFull RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullaction = "lockWhenFull"

    RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullaction_wrapWhenFull RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullaction = "wrapWhenFull"
)

// RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullstatus represents must not affect the value of this object.
type RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullstatus string

const (
    RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullstatus_spaceAvailable RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullstatus = "spaceAvailable"

    RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullstatus_full RMONMIB_Buffercontroltable_Buffercontrolentry_Buffercontrolfullstatus = "full"
)

// RMONMIB_Capturebuffertable
// A list of packets captured off of a channel.
type RMONMIB_Capturebuffertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A packet captured off of an attached network.  As an example, an instance
    // of the captureBufferPacketData object might be named
    // captureBufferPacketData.3.1783. The type is slice of
    // RMONMIB_Capturebuffertable_Capturebufferentry.
    Capturebufferentry []RMONMIB_Capturebuffertable_Capturebufferentry
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetEntityData() *types.CommonEntityData {
    capturebuffertable.EntityData.YFilter = capturebuffertable.YFilter
    capturebuffertable.EntityData.YangName = "captureBufferTable"
    capturebuffertable.EntityData.BundleName = "cisco_ios_xe"
    capturebuffertable.EntityData.ParentYangName = "RMON-MIB"
    capturebuffertable.EntityData.SegmentPath = "captureBufferTable"
    capturebuffertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capturebuffertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capturebuffertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capturebuffertable.EntityData.Children = make(map[string]types.YChild)
    capturebuffertable.EntityData.Children["captureBufferEntry"] = types.YChild{"Capturebufferentry", nil}
    for i := range capturebuffertable.Capturebufferentry {
        capturebuffertable.EntityData.Children[types.GetSegmentPath(&capturebuffertable.Capturebufferentry[i])] = types.YChild{"Capturebufferentry", &capturebuffertable.Capturebufferentry[i]}
    }
    capturebuffertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(capturebuffertable.EntityData)
}

// RMONMIB_Capturebuffertable_Capturebufferentry
// A packet captured off of an attached network.  As an
// example, an instance of the captureBufferPacketData
// object might be named captureBufferPacketData.3.1783
type RMONMIB_Capturebuffertable_Capturebufferentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the bufferControlEntry with which
    // this packet is associated. The type is interface{} with range: 1..65535.
    Capturebuffercontrolindex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // captureBuffer table associated with a particular bufferControlEntry.  This
    // index will start at 1 and increase by one for each new packet added with
    // the same captureBufferControlIndex.  Should this value reach 2147483647,
    // the next packet added with the same captureBufferControlIndex shall cause
    // this value to wrap around to 1. The type is interface{} with range:
    // 1..2147483647.
    Capturebufferindex interface{}

    // An index that describes the order of packets that are received on a
    // particular interface. The packetID of a packet captured on an interface is
    // defined to be greater than the packetID's of all packets captured
    // previously on the same interface.  As the captureBufferPacketID object has
    // a maximum positive value of 2^31 - 1, any captureBufferPacketID object
    // shall have the value of the associated packet's packetID mod 2^31. The type
    // is interface{} with range: -2147483648..2147483647.
    Capturebufferpacketid interface{}

    // The data inside the packet, starting at the beginning of the packet plus
    // any offset specified in the associated bufferControlDownloadOffset,
    // including any link level headers.  The length of the data in this object is
    // the minimum of the length of the captured packet minus the offset, the
    // length of the associated bufferControlCaptureSliceSize minus the offset,
    // and the associated bufferControlDownloadSliceSize.  If this minimum is less
    // than zero, this object shall have a length of zero. The type is string.
    Capturebufferpacketdata interface{}

    // The actual length (off the wire) of the packet stored in this entry,
    // including FCS octets. The type is interface{} with range:
    // -2147483648..2147483647. Units are Octets.
    Capturebufferpacketlength interface{}

    // The number of milliseconds that had passed since this capture buffer was
    // first turned on when this packet was captured. The type is interface{} with
    // range: -2147483648..2147483647. Units are Milliseconds.
    Capturebufferpackettime interface{}

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
    Capturebufferpacketstatus interface{}
}

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetEntityData() *types.CommonEntityData {
    capturebufferentry.EntityData.YFilter = capturebufferentry.YFilter
    capturebufferentry.EntityData.YangName = "captureBufferEntry"
    capturebufferentry.EntityData.BundleName = "cisco_ios_xe"
    capturebufferentry.EntityData.ParentYangName = "captureBufferTable"
    capturebufferentry.EntityData.SegmentPath = "captureBufferEntry" + "[captureBufferControlIndex='" + fmt.Sprintf("%v", capturebufferentry.Capturebuffercontrolindex) + "']" + "[captureBufferIndex='" + fmt.Sprintf("%v", capturebufferentry.Capturebufferindex) + "']"
    capturebufferentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capturebufferentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capturebufferentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capturebufferentry.EntityData.Children = make(map[string]types.YChild)
    capturebufferentry.EntityData.Leafs = make(map[string]types.YLeaf)
    capturebufferentry.EntityData.Leafs["captureBufferControlIndex"] = types.YLeaf{"Capturebuffercontrolindex", capturebufferentry.Capturebuffercontrolindex}
    capturebufferentry.EntityData.Leafs["captureBufferIndex"] = types.YLeaf{"Capturebufferindex", capturebufferentry.Capturebufferindex}
    capturebufferentry.EntityData.Leafs["captureBufferPacketID"] = types.YLeaf{"Capturebufferpacketid", capturebufferentry.Capturebufferpacketid}
    capturebufferentry.EntityData.Leafs["captureBufferPacketData"] = types.YLeaf{"Capturebufferpacketdata", capturebufferentry.Capturebufferpacketdata}
    capturebufferentry.EntityData.Leafs["captureBufferPacketLength"] = types.YLeaf{"Capturebufferpacketlength", capturebufferentry.Capturebufferpacketlength}
    capturebufferentry.EntityData.Leafs["captureBufferPacketTime"] = types.YLeaf{"Capturebufferpackettime", capturebufferentry.Capturebufferpackettime}
    capturebufferentry.EntityData.Leafs["captureBufferPacketStatus"] = types.YLeaf{"Capturebufferpacketstatus", capturebufferentry.Capturebufferpacketstatus}
    return &(capturebufferentry.EntityData)
}

// RMONMIB_Eventtable
// A list of events to be generated.
type RMONMIB_Eventtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of parameters that describe an event to be generated when certain
    // conditions are met.  As an example, an instance of the eventLastTimeSent
    // object might be named eventLastTimeSent.6. The type is slice of
    // RMONMIB_Eventtable_Evententry.
    Evententry []RMONMIB_Eventtable_Evententry
}

func (eventtable *RMONMIB_Eventtable) GetEntityData() *types.CommonEntityData {
    eventtable.EntityData.YFilter = eventtable.YFilter
    eventtable.EntityData.YangName = "eventTable"
    eventtable.EntityData.BundleName = "cisco_ios_xe"
    eventtable.EntityData.ParentYangName = "RMON-MIB"
    eventtable.EntityData.SegmentPath = "eventTable"
    eventtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eventtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eventtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eventtable.EntityData.Children = make(map[string]types.YChild)
    eventtable.EntityData.Children["eventEntry"] = types.YChild{"Evententry", nil}
    for i := range eventtable.Evententry {
        eventtable.EntityData.Children[types.GetSegmentPath(&eventtable.Evententry[i])] = types.YChild{"Evententry", &eventtable.Evententry[i]}
    }
    eventtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eventtable.EntityData)
}

// RMONMIB_Eventtable_Evententry
// A set of parameters that describe an event to be generated
// when certain conditions are met.  As an example, an instance
// of the eventLastTimeSent object might be named
// eventLastTimeSent.6
type RMONMIB_Eventtable_Evententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the
    // event table.  Each such entry defines one event that is to be generated
    // when the appropriate conditions occur. The type is interface{} with range:
    // 1..65535.
    Eventindex interface{}

    // A comment describing this event entry. The type is string with length:
    // 0..127.
    Eventdescription interface{}

    // The type of notification that the probe will make about this event.  In the
    // case of log, an entry is made in the log table for each event.  In the case
    // of snmp-trap, an SNMP trap is sent to one or more management stations. The
    // type is Eventtype.
    Eventtype interface{}

    // If an SNMP trap is to be sent, it will be sent to the SNMP community
    // specified by this octet string. The type is string with length: 0..127.
    Eventcommunity interface{}

    // The value of sysUpTime at the time this event entry last generated an
    // event.  If this entry has not generated any events, this value will be
    // zero. The type is interface{} with range: 0..4294967295.
    Eventlasttimesent interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it.  If this object contains a string starting with 'monitor'
    // and has associated entries in the log table, all connected management
    // stations should retrieve those log entries, as they may have significance
    // to all management stations connected to this device. The type is string
    // with length: 0..127.
    Eventowner interface{}

    // The status of this event entry.  If this object is not equal to valid(1),
    // all associated log entries shall be deleted by the agent. The type is
    // EntryStatus.
    Eventstatus interface{}
}

func (evententry *RMONMIB_Eventtable_Evententry) GetEntityData() *types.CommonEntityData {
    evententry.EntityData.YFilter = evententry.YFilter
    evententry.EntityData.YangName = "eventEntry"
    evententry.EntityData.BundleName = "cisco_ios_xe"
    evententry.EntityData.ParentYangName = "eventTable"
    evententry.EntityData.SegmentPath = "eventEntry" + "[eventIndex='" + fmt.Sprintf("%v", evententry.Eventindex) + "']"
    evententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    evententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    evententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    evententry.EntityData.Children = make(map[string]types.YChild)
    evententry.EntityData.Leafs = make(map[string]types.YLeaf)
    evententry.EntityData.Leafs["eventIndex"] = types.YLeaf{"Eventindex", evententry.Eventindex}
    evententry.EntityData.Leafs["eventDescription"] = types.YLeaf{"Eventdescription", evententry.Eventdescription}
    evententry.EntityData.Leafs["eventType"] = types.YLeaf{"Eventtype", evententry.Eventtype}
    evententry.EntityData.Leafs["eventCommunity"] = types.YLeaf{"Eventcommunity", evententry.Eventcommunity}
    evententry.EntityData.Leafs["eventLastTimeSent"] = types.YLeaf{"Eventlasttimesent", evententry.Eventlasttimesent}
    evententry.EntityData.Leafs["eventOwner"] = types.YLeaf{"Eventowner", evententry.Eventowner}
    evententry.EntityData.Leafs["eventStatus"] = types.YLeaf{"Eventstatus", evententry.Eventstatus}
    return &(evententry.EntityData)
}

// RMONMIB_Eventtable_Evententry_Eventtype represents management stations.
type RMONMIB_Eventtable_Evententry_Eventtype string

const (
    RMONMIB_Eventtable_Evententry_Eventtype_none RMONMIB_Eventtable_Evententry_Eventtype = "none"

    RMONMIB_Eventtable_Evententry_Eventtype_log RMONMIB_Eventtable_Evententry_Eventtype = "log"

    RMONMIB_Eventtable_Evententry_Eventtype_snmptrap RMONMIB_Eventtable_Evententry_Eventtype = "snmptrap"

    RMONMIB_Eventtable_Evententry_Eventtype_logandtrap RMONMIB_Eventtable_Evententry_Eventtype = "logandtrap"
)

// RMONMIB_Logtable
// A list of events that have been logged.
type RMONMIB_Logtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of data describing an event that has been logged.  For example, an
    // instance of the logDescription object might be named logDescription.6.47.
    // The type is slice of RMONMIB_Logtable_Logentry.
    Logentry []RMONMIB_Logtable_Logentry
}

func (logtable *RMONMIB_Logtable) GetEntityData() *types.CommonEntityData {
    logtable.EntityData.YFilter = logtable.YFilter
    logtable.EntityData.YangName = "logTable"
    logtable.EntityData.BundleName = "cisco_ios_xe"
    logtable.EntityData.ParentYangName = "RMON-MIB"
    logtable.EntityData.SegmentPath = "logTable"
    logtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    logtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    logtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    logtable.EntityData.Children = make(map[string]types.YChild)
    logtable.EntityData.Children["logEntry"] = types.YChild{"Logentry", nil}
    for i := range logtable.Logentry {
        logtable.EntityData.Children[types.GetSegmentPath(&logtable.Logentry[i])] = types.YChild{"Logentry", &logtable.Logentry[i]}
    }
    logtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(logtable.EntityData)
}

// RMONMIB_Logtable_Logentry
// A set of data describing an event that has been
// logged.  For example, an instance of the logDescription
// object might be named logDescription.6.47
type RMONMIB_Logtable_Logentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The event entry that generated this log entry. 
    // The log identified by a particular value of this index is associated with
    // the same eventEntry as identified by the same value of eventIndex. The type
    // is interface{} with range: 1..65535.
    Logeventindex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // log table amongst those generated by the same eventEntries.  These indexes
    // are assigned beginning with 1 and increase by one with each new log entry. 
    // The association between values of logIndex and logEntries is fixed for the
    // lifetime of each logEntry. The agent may choose to delete the oldest
    // instances of logEntry as required because of lack of memory.  It is an
    // implementation-specific matter as to when this deletion may occur. The type
    // is interface{} with range: 1..2147483647.
    Logindex interface{}

    // The value of sysUpTime when this log entry was created. The type is
    // interface{} with range: 0..4294967295.
    Logtime interface{}

    // An implementation dependent description of the event that activated this
    // log entry. The type is string with length: 0..255.
    Logdescription interface{}
}

func (logentry *RMONMIB_Logtable_Logentry) GetEntityData() *types.CommonEntityData {
    logentry.EntityData.YFilter = logentry.YFilter
    logentry.EntityData.YangName = "logEntry"
    logentry.EntityData.BundleName = "cisco_ios_xe"
    logentry.EntityData.ParentYangName = "logTable"
    logentry.EntityData.SegmentPath = "logEntry" + "[logEventIndex='" + fmt.Sprintf("%v", logentry.Logeventindex) + "']" + "[logIndex='" + fmt.Sprintf("%v", logentry.Logindex) + "']"
    logentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    logentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    logentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    logentry.EntityData.Children = make(map[string]types.YChild)
    logentry.EntityData.Leafs = make(map[string]types.YLeaf)
    logentry.EntityData.Leafs["logEventIndex"] = types.YLeaf{"Logeventindex", logentry.Logeventindex}
    logentry.EntityData.Leafs["logIndex"] = types.YLeaf{"Logindex", logentry.Logindex}
    logentry.EntityData.Leafs["logTime"] = types.YLeaf{"Logtime", logentry.Logtime}
    logentry.EntityData.Leafs["logDescription"] = types.YLeaf{"Logdescription", logentry.Logdescription}
    return &(logentry.EntityData)
}

