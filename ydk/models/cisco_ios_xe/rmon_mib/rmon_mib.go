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
    parent types.Entity
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

func (rMONMIB *RMONMIB) GetFilter() yfilter.YFilter { return rMONMIB.YFilter }

func (rMONMIB *RMONMIB) SetFilter(yf yfilter.YFilter) { rMONMIB.YFilter = yf }

func (rMONMIB *RMONMIB) GetGoName(yname string) string {
    if yname == "etherStatsTable" { return "Etherstatstable" }
    if yname == "historyControlTable" { return "Historycontroltable" }
    if yname == "etherHistoryTable" { return "Etherhistorytable" }
    if yname == "alarmTable" { return "Alarmtable" }
    if yname == "hostControlTable" { return "Hostcontroltable" }
    if yname == "hostTable" { return "Hosttable" }
    if yname == "hostTimeTable" { return "Hosttimetable" }
    if yname == "hostTopNControlTable" { return "Hosttopncontroltable" }
    if yname == "hostTopNTable" { return "Hosttopntable" }
    if yname == "matrixControlTable" { return "Matrixcontroltable" }
    if yname == "matrixSDTable" { return "Matrixsdtable" }
    if yname == "matrixDSTable" { return "Matrixdstable" }
    if yname == "filterTable" { return "Filtertable" }
    if yname == "channelTable" { return "Channeltable" }
    if yname == "bufferControlTable" { return "Buffercontroltable" }
    if yname == "captureBufferTable" { return "Capturebuffertable" }
    if yname == "eventTable" { return "Eventtable" }
    if yname == "logTable" { return "Logtable" }
    return ""
}

func (rMONMIB *RMONMIB) GetSegmentPath() string {
    return "RMON-MIB:RMON-MIB"
}

func (rMONMIB *RMONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherStatsTable" {
        return &rMONMIB.Etherstatstable
    }
    if childYangName == "historyControlTable" {
        return &rMONMIB.Historycontroltable
    }
    if childYangName == "etherHistoryTable" {
        return &rMONMIB.Etherhistorytable
    }
    if childYangName == "alarmTable" {
        return &rMONMIB.Alarmtable
    }
    if childYangName == "hostControlTable" {
        return &rMONMIB.Hostcontroltable
    }
    if childYangName == "hostTable" {
        return &rMONMIB.Hosttable
    }
    if childYangName == "hostTimeTable" {
        return &rMONMIB.Hosttimetable
    }
    if childYangName == "hostTopNControlTable" {
        return &rMONMIB.Hosttopncontroltable
    }
    if childYangName == "hostTopNTable" {
        return &rMONMIB.Hosttopntable
    }
    if childYangName == "matrixControlTable" {
        return &rMONMIB.Matrixcontroltable
    }
    if childYangName == "matrixSDTable" {
        return &rMONMIB.Matrixsdtable
    }
    if childYangName == "matrixDSTable" {
        return &rMONMIB.Matrixdstable
    }
    if childYangName == "filterTable" {
        return &rMONMIB.Filtertable
    }
    if childYangName == "channelTable" {
        return &rMONMIB.Channeltable
    }
    if childYangName == "bufferControlTable" {
        return &rMONMIB.Buffercontroltable
    }
    if childYangName == "captureBufferTable" {
        return &rMONMIB.Capturebuffertable
    }
    if childYangName == "eventTable" {
        return &rMONMIB.Eventtable
    }
    if childYangName == "logTable" {
        return &rMONMIB.Logtable
    }
    return nil
}

func (rMONMIB *RMONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["etherStatsTable"] = &rMONMIB.Etherstatstable
    children["historyControlTable"] = &rMONMIB.Historycontroltable
    children["etherHistoryTable"] = &rMONMIB.Etherhistorytable
    children["alarmTable"] = &rMONMIB.Alarmtable
    children["hostControlTable"] = &rMONMIB.Hostcontroltable
    children["hostTable"] = &rMONMIB.Hosttable
    children["hostTimeTable"] = &rMONMIB.Hosttimetable
    children["hostTopNControlTable"] = &rMONMIB.Hosttopncontroltable
    children["hostTopNTable"] = &rMONMIB.Hosttopntable
    children["matrixControlTable"] = &rMONMIB.Matrixcontroltable
    children["matrixSDTable"] = &rMONMIB.Matrixsdtable
    children["matrixDSTable"] = &rMONMIB.Matrixdstable
    children["filterTable"] = &rMONMIB.Filtertable
    children["channelTable"] = &rMONMIB.Channeltable
    children["bufferControlTable"] = &rMONMIB.Buffercontroltable
    children["captureBufferTable"] = &rMONMIB.Capturebuffertable
    children["eventTable"] = &rMONMIB.Eventtable
    children["logTable"] = &rMONMIB.Logtable
    return children
}

func (rMONMIB *RMONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rMONMIB *RMONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (rMONMIB *RMONMIB) GetYangName() string { return "RMON-MIB" }

func (rMONMIB *RMONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rMONMIB *RMONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rMONMIB *RMONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rMONMIB *RMONMIB) SetParent(parent types.Entity) { rMONMIB.parent = parent }

func (rMONMIB *RMONMIB) GetParent() types.Entity { return rMONMIB.parent }

func (rMONMIB *RMONMIB) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Etherstatstable
// A list of Ethernet statistics entries.
type RMONMIB_Etherstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics kept for a particular Ethernet interface.  As an
    // example, an instance of the etherStatsPkts object might be named
    // etherStatsPkts.1. The type is slice of
    // RMONMIB_Etherstatstable_Etherstatsentry.
    Etherstatsentry []RMONMIB_Etherstatstable_Etherstatsentry
}

func (etherstatstable *RMONMIB_Etherstatstable) GetFilter() yfilter.YFilter { return etherstatstable.YFilter }

func (etherstatstable *RMONMIB_Etherstatstable) SetFilter(yf yfilter.YFilter) { etherstatstable.YFilter = yf }

func (etherstatstable *RMONMIB_Etherstatstable) GetGoName(yname string) string {
    if yname == "etherStatsEntry" { return "Etherstatsentry" }
    return ""
}

func (etherstatstable *RMONMIB_Etherstatstable) GetSegmentPath() string {
    return "etherStatsTable"
}

func (etherstatstable *RMONMIB_Etherstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherStatsEntry" {
        for _, c := range etherstatstable.Etherstatsentry {
            if etherstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Etherstatstable_Etherstatsentry{}
        etherstatstable.Etherstatsentry = append(etherstatstable.Etherstatsentry, child)
        return &etherstatstable.Etherstatsentry[len(etherstatstable.Etherstatsentry)-1]
    }
    return nil
}

func (etherstatstable *RMONMIB_Etherstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etherstatstable.Etherstatsentry {
        children[etherstatstable.Etherstatsentry[i].GetSegmentPath()] = &etherstatstable.Etherstatsentry[i]
    }
    return children
}

func (etherstatstable *RMONMIB_Etherstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherstatstable *RMONMIB_Etherstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (etherstatstable *RMONMIB_Etherstatstable) GetYangName() string { return "etherStatsTable" }

func (etherstatstable *RMONMIB_Etherstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherstatstable *RMONMIB_Etherstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherstatstable *RMONMIB_Etherstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherstatstable *RMONMIB_Etherstatstable) SetParent(parent types.Entity) { etherstatstable.parent = parent }

func (etherstatstable *RMONMIB_Etherstatstable) GetParent() types.Entity { return etherstatstable.parent }

func (etherstatstable *RMONMIB_Etherstatstable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Etherstatstable_Etherstatsentry
// A collection of statistics kept for a particular
// Ethernet interface.  As an example, an instance of the
// etherStatsPkts object might be named etherStatsPkts.1
type RMONMIB_Etherstatstable_Etherstatsentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetFilter() yfilter.YFilter { return etherstatsentry.YFilter }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) SetFilter(yf yfilter.YFilter) { etherstatsentry.YFilter = yf }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetGoName(yname string) string {
    if yname == "etherStatsIndex" { return "Etherstatsindex" }
    if yname == "etherStatsDataSource" { return "Etherstatsdatasource" }
    if yname == "etherStatsDropEvents" { return "Etherstatsdropevents" }
    if yname == "etherStatsOctets" { return "Etherstatsoctets" }
    if yname == "etherStatsPkts" { return "Etherstatspkts" }
    if yname == "etherStatsBroadcastPkts" { return "Etherstatsbroadcastpkts" }
    if yname == "etherStatsMulticastPkts" { return "Etherstatsmulticastpkts" }
    if yname == "etherStatsCRCAlignErrors" { return "Etherstatscrcalignerrors" }
    if yname == "etherStatsUndersizePkts" { return "Etherstatsundersizepkts" }
    if yname == "etherStatsOversizePkts" { return "Etherstatsoversizepkts" }
    if yname == "etherStatsFragments" { return "Etherstatsfragments" }
    if yname == "etherStatsJabbers" { return "Etherstatsjabbers" }
    if yname == "etherStatsCollisions" { return "Etherstatscollisions" }
    if yname == "etherStatsPkts64Octets" { return "Etherstatspkts64Octets" }
    if yname == "etherStatsPkts65to127Octets" { return "Etherstatspkts65To127Octets" }
    if yname == "etherStatsPkts128to255Octets" { return "Etherstatspkts128To255Octets" }
    if yname == "etherStatsPkts256to511Octets" { return "Etherstatspkts256To511Octets" }
    if yname == "etherStatsPkts512to1023Octets" { return "Etherstatspkts512To1023Octets" }
    if yname == "etherStatsPkts1024to1518Octets" { return "Etherstatspkts1024To1518Octets" }
    if yname == "etherStatsOwner" { return "Etherstatsowner" }
    if yname == "etherStatsStatus" { return "Etherstatsstatus" }
    if yname == "etherStatsDroppedFrames" { return "Etherstatsdroppedframes" }
    if yname == "etherStatsCreateTime" { return "Etherstatscreatetime" }
    return ""
}

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetSegmentPath() string {
    return "etherStatsEntry" + "[etherStatsIndex='" + fmt.Sprintf("%v", etherstatsentry.Etherstatsindex) + "']"
}

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["etherStatsIndex"] = etherstatsentry.Etherstatsindex
    leafs["etherStatsDataSource"] = etherstatsentry.Etherstatsdatasource
    leafs["etherStatsDropEvents"] = etherstatsentry.Etherstatsdropevents
    leafs["etherStatsOctets"] = etherstatsentry.Etherstatsoctets
    leafs["etherStatsPkts"] = etherstatsentry.Etherstatspkts
    leafs["etherStatsBroadcastPkts"] = etherstatsentry.Etherstatsbroadcastpkts
    leafs["etherStatsMulticastPkts"] = etherstatsentry.Etherstatsmulticastpkts
    leafs["etherStatsCRCAlignErrors"] = etherstatsentry.Etherstatscrcalignerrors
    leafs["etherStatsUndersizePkts"] = etherstatsentry.Etherstatsundersizepkts
    leafs["etherStatsOversizePkts"] = etherstatsentry.Etherstatsoversizepkts
    leafs["etherStatsFragments"] = etherstatsentry.Etherstatsfragments
    leafs["etherStatsJabbers"] = etherstatsentry.Etherstatsjabbers
    leafs["etherStatsCollisions"] = etherstatsentry.Etherstatscollisions
    leafs["etherStatsPkts64Octets"] = etherstatsentry.Etherstatspkts64Octets
    leafs["etherStatsPkts65to127Octets"] = etherstatsentry.Etherstatspkts65To127Octets
    leafs["etherStatsPkts128to255Octets"] = etherstatsentry.Etherstatspkts128To255Octets
    leafs["etherStatsPkts256to511Octets"] = etherstatsentry.Etherstatspkts256To511Octets
    leafs["etherStatsPkts512to1023Octets"] = etherstatsentry.Etherstatspkts512To1023Octets
    leafs["etherStatsPkts1024to1518Octets"] = etherstatsentry.Etherstatspkts1024To1518Octets
    leafs["etherStatsOwner"] = etherstatsentry.Etherstatsowner
    leafs["etherStatsStatus"] = etherstatsentry.Etherstatsstatus
    leafs["etherStatsDroppedFrames"] = etherstatsentry.Etherstatsdroppedframes
    leafs["etherStatsCreateTime"] = etherstatsentry.Etherstatscreatetime
    return leafs
}

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetYangName() string { return "etherStatsEntry" }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) SetParent(parent types.Entity) { etherstatsentry.parent = parent }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetParent() types.Entity { return etherstatsentry.parent }

func (etherstatsentry *RMONMIB_Etherstatstable_Etherstatsentry) GetParentYangName() string { return "etherStatsTable" }

// RMONMIB_Historycontroltable
// A list of history control entries.
type RMONMIB_Historycontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of parameters that set up a periodic sampling of statistics.  As an
    // example, an instance of the historyControlInterval object might be named
    // historyControlInterval.2. The type is slice of
    // RMONMIB_Historycontroltable_Historycontrolentry.
    Historycontrolentry []RMONMIB_Historycontroltable_Historycontrolentry
}

func (historycontroltable *RMONMIB_Historycontroltable) GetFilter() yfilter.YFilter { return historycontroltable.YFilter }

func (historycontroltable *RMONMIB_Historycontroltable) SetFilter(yf yfilter.YFilter) { historycontroltable.YFilter = yf }

func (historycontroltable *RMONMIB_Historycontroltable) GetGoName(yname string) string {
    if yname == "historyControlEntry" { return "Historycontrolentry" }
    return ""
}

func (historycontroltable *RMONMIB_Historycontroltable) GetSegmentPath() string {
    return "historyControlTable"
}

func (historycontroltable *RMONMIB_Historycontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "historyControlEntry" {
        for _, c := range historycontroltable.Historycontrolentry {
            if historycontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Historycontroltable_Historycontrolentry{}
        historycontroltable.Historycontrolentry = append(historycontroltable.Historycontrolentry, child)
        return &historycontroltable.Historycontrolentry[len(historycontroltable.Historycontrolentry)-1]
    }
    return nil
}

func (historycontroltable *RMONMIB_Historycontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range historycontroltable.Historycontrolentry {
        children[historycontroltable.Historycontrolentry[i].GetSegmentPath()] = &historycontroltable.Historycontrolentry[i]
    }
    return children
}

func (historycontroltable *RMONMIB_Historycontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (historycontroltable *RMONMIB_Historycontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (historycontroltable *RMONMIB_Historycontroltable) GetYangName() string { return "historyControlTable" }

func (historycontroltable *RMONMIB_Historycontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (historycontroltable *RMONMIB_Historycontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (historycontroltable *RMONMIB_Historycontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (historycontroltable *RMONMIB_Historycontroltable) SetParent(parent types.Entity) { historycontroltable.parent = parent }

func (historycontroltable *RMONMIB_Historycontroltable) GetParent() types.Entity { return historycontroltable.parent }

func (historycontroltable *RMONMIB_Historycontroltable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Historycontroltable_Historycontrolentry
// A list of parameters that set up a periodic sampling of
// statistics.  As an example, an instance of the
// historyControlInterval object might be named
// historyControlInterval.2
type RMONMIB_Historycontroltable_Historycontrolentry struct {
    parent types.Entity
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetFilter() yfilter.YFilter { return historycontrolentry.YFilter }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) SetFilter(yf yfilter.YFilter) { historycontrolentry.YFilter = yf }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetGoName(yname string) string {
    if yname == "historyControlIndex" { return "Historycontrolindex" }
    if yname == "historyControlDataSource" { return "Historycontroldatasource" }
    if yname == "historyControlBucketsRequested" { return "Historycontrolbucketsrequested" }
    if yname == "historyControlBucketsGranted" { return "Historycontrolbucketsgranted" }
    if yname == "historyControlInterval" { return "Historycontrolinterval" }
    if yname == "historyControlOwner" { return "Historycontrolowner" }
    if yname == "historyControlStatus" { return "Historycontrolstatus" }
    if yname == "historyControlDroppedFrames" { return "Historycontroldroppedframes" }
    return ""
}

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetSegmentPath() string {
    return "historyControlEntry" + "[historyControlIndex='" + fmt.Sprintf("%v", historycontrolentry.Historycontrolindex) + "']"
}

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["historyControlIndex"] = historycontrolentry.Historycontrolindex
    leafs["historyControlDataSource"] = historycontrolentry.Historycontroldatasource
    leafs["historyControlBucketsRequested"] = historycontrolentry.Historycontrolbucketsrequested
    leafs["historyControlBucketsGranted"] = historycontrolentry.Historycontrolbucketsgranted
    leafs["historyControlInterval"] = historycontrolentry.Historycontrolinterval
    leafs["historyControlOwner"] = historycontrolentry.Historycontrolowner
    leafs["historyControlStatus"] = historycontrolentry.Historycontrolstatus
    leafs["historyControlDroppedFrames"] = historycontrolentry.Historycontroldroppedframes
    return leafs
}

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetYangName() string { return "historyControlEntry" }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) SetParent(parent types.Entity) { historycontrolentry.parent = parent }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetParent() types.Entity { return historycontrolentry.parent }

func (historycontrolentry *RMONMIB_Historycontroltable_Historycontrolentry) GetParentYangName() string { return "historyControlTable" }

// RMONMIB_Etherhistorytable
// A list of Ethernet history entries.
type RMONMIB_Etherhistorytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An historical sample of Ethernet statistics on a particular Ethernet
    // interface.  This sample is associated with the historyControlEntry which
    // set up the parameters for a regular collection of these samples.  As an
    // example, an instance of the etherHistoryPkts object might be named
    // etherHistoryPkts.2.89. The type is slice of
    // RMONMIB_Etherhistorytable_Etherhistoryentry.
    Etherhistoryentry []RMONMIB_Etherhistorytable_Etherhistoryentry
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetFilter() yfilter.YFilter { return etherhistorytable.YFilter }

func (etherhistorytable *RMONMIB_Etherhistorytable) SetFilter(yf yfilter.YFilter) { etherhistorytable.YFilter = yf }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetGoName(yname string) string {
    if yname == "etherHistoryEntry" { return "Etherhistoryentry" }
    return ""
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetSegmentPath() string {
    return "etherHistoryTable"
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etherHistoryEntry" {
        for _, c := range etherhistorytable.Etherhistoryentry {
            if etherhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Etherhistorytable_Etherhistoryentry{}
        etherhistorytable.Etherhistoryentry = append(etherhistorytable.Etherhistoryentry, child)
        return &etherhistorytable.Etherhistoryentry[len(etherhistorytable.Etherhistoryentry)-1]
    }
    return nil
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etherhistorytable.Etherhistoryentry {
        children[etherhistorytable.Etherhistoryentry[i].GetSegmentPath()] = &etherhistorytable.Etherhistoryentry[i]
    }
    return children
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherhistorytable *RMONMIB_Etherhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetYangName() string { return "etherHistoryTable" }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherhistorytable *RMONMIB_Etherhistorytable) SetParent(parent types.Entity) { etherhistorytable.parent = parent }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetParent() types.Entity { return etherhistorytable.parent }

func (etherhistorytable *RMONMIB_Etherhistorytable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Etherhistorytable_Etherhistoryentry
// An historical sample of Ethernet statistics on a particular
// Ethernet interface.  This sample is associated with the
// historyControlEntry which set up the parameters for
// a regular collection of these samples.  As an example, an
// instance of the etherHistoryPkts object might be named
// etherHistoryPkts.2.89
type RMONMIB_Etherhistorytable_Etherhistoryentry struct {
    parent types.Entity
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

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetFilter() yfilter.YFilter { return etherhistoryentry.YFilter }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) SetFilter(yf yfilter.YFilter) { etherhistoryentry.YFilter = yf }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetGoName(yname string) string {
    if yname == "etherHistoryIndex" { return "Etherhistoryindex" }
    if yname == "etherHistorySampleIndex" { return "Etherhistorysampleindex" }
    if yname == "etherHistoryIntervalStart" { return "Etherhistoryintervalstart" }
    if yname == "etherHistoryDropEvents" { return "Etherhistorydropevents" }
    if yname == "etherHistoryOctets" { return "Etherhistoryoctets" }
    if yname == "etherHistoryPkts" { return "Etherhistorypkts" }
    if yname == "etherHistoryBroadcastPkts" { return "Etherhistorybroadcastpkts" }
    if yname == "etherHistoryMulticastPkts" { return "Etherhistorymulticastpkts" }
    if yname == "etherHistoryCRCAlignErrors" { return "Etherhistorycrcalignerrors" }
    if yname == "etherHistoryUndersizePkts" { return "Etherhistoryundersizepkts" }
    if yname == "etherHistoryOversizePkts" { return "Etherhistoryoversizepkts" }
    if yname == "etherHistoryFragments" { return "Etherhistoryfragments" }
    if yname == "etherHistoryJabbers" { return "Etherhistoryjabbers" }
    if yname == "etherHistoryCollisions" { return "Etherhistorycollisions" }
    if yname == "etherHistoryUtilization" { return "Etherhistoryutilization" }
    return ""
}

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetSegmentPath() string {
    return "etherHistoryEntry" + "[etherHistoryIndex='" + fmt.Sprintf("%v", etherhistoryentry.Etherhistoryindex) + "']" + "[etherHistorySampleIndex='" + fmt.Sprintf("%v", etherhistoryentry.Etherhistorysampleindex) + "']"
}

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["etherHistoryIndex"] = etherhistoryentry.Etherhistoryindex
    leafs["etherHistorySampleIndex"] = etherhistoryentry.Etherhistorysampleindex
    leafs["etherHistoryIntervalStart"] = etherhistoryentry.Etherhistoryintervalstart
    leafs["etherHistoryDropEvents"] = etherhistoryentry.Etherhistorydropevents
    leafs["etherHistoryOctets"] = etherhistoryentry.Etherhistoryoctets
    leafs["etherHistoryPkts"] = etherhistoryentry.Etherhistorypkts
    leafs["etherHistoryBroadcastPkts"] = etherhistoryentry.Etherhistorybroadcastpkts
    leafs["etherHistoryMulticastPkts"] = etherhistoryentry.Etherhistorymulticastpkts
    leafs["etherHistoryCRCAlignErrors"] = etherhistoryentry.Etherhistorycrcalignerrors
    leafs["etherHistoryUndersizePkts"] = etherhistoryentry.Etherhistoryundersizepkts
    leafs["etherHistoryOversizePkts"] = etherhistoryentry.Etherhistoryoversizepkts
    leafs["etherHistoryFragments"] = etherhistoryentry.Etherhistoryfragments
    leafs["etherHistoryJabbers"] = etherhistoryentry.Etherhistoryjabbers
    leafs["etherHistoryCollisions"] = etherhistoryentry.Etherhistorycollisions
    leafs["etherHistoryUtilization"] = etherhistoryentry.Etherhistoryutilization
    return leafs
}

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetYangName() string { return "etherHistoryEntry" }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) SetParent(parent types.Entity) { etherhistoryentry.parent = parent }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetParent() types.Entity { return etherhistoryentry.parent }

func (etherhistoryentry *RMONMIB_Etherhistorytable_Etherhistoryentry) GetParentYangName() string { return "etherHistoryTable" }

// RMONMIB_Alarmtable
// A list of alarm entries.
type RMONMIB_Alarmtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of parameters that set up a periodic checking for alarm conditions. 
    // For example, an instance of the alarmValue object might be named
    // alarmValue.8. The type is slice of RMONMIB_Alarmtable_Alarmentry.
    Alarmentry []RMONMIB_Alarmtable_Alarmentry
}

func (alarmtable *RMONMIB_Alarmtable) GetFilter() yfilter.YFilter { return alarmtable.YFilter }

func (alarmtable *RMONMIB_Alarmtable) SetFilter(yf yfilter.YFilter) { alarmtable.YFilter = yf }

func (alarmtable *RMONMIB_Alarmtable) GetGoName(yname string) string {
    if yname == "alarmEntry" { return "Alarmentry" }
    return ""
}

func (alarmtable *RMONMIB_Alarmtable) GetSegmentPath() string {
    return "alarmTable"
}

func (alarmtable *RMONMIB_Alarmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarmEntry" {
        for _, c := range alarmtable.Alarmentry {
            if alarmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Alarmtable_Alarmentry{}
        alarmtable.Alarmentry = append(alarmtable.Alarmentry, child)
        return &alarmtable.Alarmentry[len(alarmtable.Alarmentry)-1]
    }
    return nil
}

func (alarmtable *RMONMIB_Alarmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alarmtable.Alarmentry {
        children[alarmtable.Alarmentry[i].GetSegmentPath()] = &alarmtable.Alarmentry[i]
    }
    return children
}

func (alarmtable *RMONMIB_Alarmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarmtable *RMONMIB_Alarmtable) GetBundleName() string { return "cisco_ios_xe" }

func (alarmtable *RMONMIB_Alarmtable) GetYangName() string { return "alarmTable" }

func (alarmtable *RMONMIB_Alarmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (alarmtable *RMONMIB_Alarmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (alarmtable *RMONMIB_Alarmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (alarmtable *RMONMIB_Alarmtable) SetParent(parent types.Entity) { alarmtable.parent = parent }

func (alarmtable *RMONMIB_Alarmtable) GetParent() types.Entity { return alarmtable.parent }

func (alarmtable *RMONMIB_Alarmtable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Alarmtable_Alarmentry
// A list of parameters that set up a periodic checking
// for alarm conditions.  For example, an instance of the
// alarmValue object might be named alarmValue.8
type RMONMIB_Alarmtable_Alarmentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetFilter() yfilter.YFilter { return alarmentry.YFilter }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) SetFilter(yf yfilter.YFilter) { alarmentry.YFilter = yf }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetGoName(yname string) string {
    if yname == "alarmIndex" { return "Alarmindex" }
    if yname == "alarmInterval" { return "Alarminterval" }
    if yname == "alarmVariable" { return "Alarmvariable" }
    if yname == "alarmSampleType" { return "Alarmsampletype" }
    if yname == "alarmValue" { return "Alarmvalue" }
    if yname == "alarmStartupAlarm" { return "Alarmstartupalarm" }
    if yname == "alarmRisingThreshold" { return "Alarmrisingthreshold" }
    if yname == "alarmFallingThreshold" { return "Alarmfallingthreshold" }
    if yname == "alarmRisingEventIndex" { return "Alarmrisingeventindex" }
    if yname == "alarmFallingEventIndex" { return "Alarmfallingeventindex" }
    if yname == "alarmOwner" { return "Alarmowner" }
    if yname == "alarmStatus" { return "Alarmstatus" }
    return ""
}

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetSegmentPath() string {
    return "alarmEntry" + "[alarmIndex='" + fmt.Sprintf("%v", alarmentry.Alarmindex) + "']"
}

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["alarmIndex"] = alarmentry.Alarmindex
    leafs["alarmInterval"] = alarmentry.Alarminterval
    leafs["alarmVariable"] = alarmentry.Alarmvariable
    leafs["alarmSampleType"] = alarmentry.Alarmsampletype
    leafs["alarmValue"] = alarmentry.Alarmvalue
    leafs["alarmStartupAlarm"] = alarmentry.Alarmstartupalarm
    leafs["alarmRisingThreshold"] = alarmentry.Alarmrisingthreshold
    leafs["alarmFallingThreshold"] = alarmentry.Alarmfallingthreshold
    leafs["alarmRisingEventIndex"] = alarmentry.Alarmrisingeventindex
    leafs["alarmFallingEventIndex"] = alarmentry.Alarmfallingeventindex
    leafs["alarmOwner"] = alarmentry.Alarmowner
    leafs["alarmStatus"] = alarmentry.Alarmstatus
    return leafs
}

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetBundleName() string { return "cisco_ios_xe" }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetYangName() string { return "alarmEntry" }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) SetParent(parent types.Entity) { alarmentry.parent = parent }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetParent() types.Entity { return alarmentry.parent }

func (alarmentry *RMONMIB_Alarmtable_Alarmentry) GetParentYangName() string { return "alarmTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of parameters that set up the discovery of hosts on a particular
    // interface and the collection of statistics about these hosts.  For example,
    // an instance of the hostControlTableSize object might be named
    // hostControlTableSize.1. The type is slice of
    // RMONMIB_Hostcontroltable_Hostcontrolentry.
    Hostcontrolentry []RMONMIB_Hostcontroltable_Hostcontrolentry
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetFilter() yfilter.YFilter { return hostcontroltable.YFilter }

func (hostcontroltable *RMONMIB_Hostcontroltable) SetFilter(yf yfilter.YFilter) { hostcontroltable.YFilter = yf }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetGoName(yname string) string {
    if yname == "hostControlEntry" { return "Hostcontrolentry" }
    return ""
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetSegmentPath() string {
    return "hostControlTable"
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hostControlEntry" {
        for _, c := range hostcontroltable.Hostcontrolentry {
            if hostcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Hostcontroltable_Hostcontrolentry{}
        hostcontroltable.Hostcontrolentry = append(hostcontroltable.Hostcontrolentry, child)
        return &hostcontroltable.Hostcontrolentry[len(hostcontroltable.Hostcontrolentry)-1]
    }
    return nil
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hostcontroltable.Hostcontrolentry {
        children[hostcontroltable.Hostcontrolentry[i].GetSegmentPath()] = &hostcontroltable.Hostcontrolentry[i]
    }
    return children
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hostcontroltable *RMONMIB_Hostcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetYangName() string { return "hostControlTable" }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hostcontroltable *RMONMIB_Hostcontroltable) SetParent(parent types.Entity) { hostcontroltable.parent = parent }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetParent() types.Entity { return hostcontroltable.parent }

func (hostcontroltable *RMONMIB_Hostcontroltable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Hostcontroltable_Hostcontrolentry
// A list of parameters that set up the discovery of hosts
// on a particular interface and the collection of statistics
// about these hosts.  For example, an instance of the
// hostControlTableSize object might be named
// hostControlTableSize.1
type RMONMIB_Hostcontroltable_Hostcontrolentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetFilter() yfilter.YFilter { return hostcontrolentry.YFilter }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) SetFilter(yf yfilter.YFilter) { hostcontrolentry.YFilter = yf }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetGoName(yname string) string {
    if yname == "hostControlIndex" { return "Hostcontrolindex" }
    if yname == "hostControlDataSource" { return "Hostcontroldatasource" }
    if yname == "hostControlTableSize" { return "Hostcontroltablesize" }
    if yname == "hostControlLastDeleteTime" { return "Hostcontrollastdeletetime" }
    if yname == "hostControlOwner" { return "Hostcontrolowner" }
    if yname == "hostControlStatus" { return "Hostcontrolstatus" }
    if yname == "hostControlDroppedFrames" { return "Hostcontroldroppedframes" }
    if yname == "hostControlCreateTime" { return "Hostcontrolcreatetime" }
    return ""
}

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetSegmentPath() string {
    return "hostControlEntry" + "[hostControlIndex='" + fmt.Sprintf("%v", hostcontrolentry.Hostcontrolindex) + "']"
}

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hostControlIndex"] = hostcontrolentry.Hostcontrolindex
    leafs["hostControlDataSource"] = hostcontrolentry.Hostcontroldatasource
    leafs["hostControlTableSize"] = hostcontrolentry.Hostcontroltablesize
    leafs["hostControlLastDeleteTime"] = hostcontrolentry.Hostcontrollastdeletetime
    leafs["hostControlOwner"] = hostcontrolentry.Hostcontrolowner
    leafs["hostControlStatus"] = hostcontrolentry.Hostcontrolstatus
    leafs["hostControlDroppedFrames"] = hostcontrolentry.Hostcontroldroppedframes
    leafs["hostControlCreateTime"] = hostcontrolentry.Hostcontrolcreatetime
    return leafs
}

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetYangName() string { return "hostControlEntry" }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) SetParent(parent types.Entity) { hostcontrolentry.parent = parent }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetParent() types.Entity { return hostcontrolentry.parent }

func (hostcontrolentry *RMONMIB_Hostcontroltable_Hostcontrolentry) GetParentYangName() string { return "hostControlTable" }

// RMONMIB_Hosttable
// A list of host entries.
type RMONMIB_Hosttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for a particular host that has been discovered
    // on an interface of this device.  For example, an instance of the
    // hostOutBroadcastPkts object might be named
    // hostOutBroadcastPkts.1.6.8.0.32.27.3.176. The type is slice of
    // RMONMIB_Hosttable_Hostentry.
    Hostentry []RMONMIB_Hosttable_Hostentry
}

func (hosttable *RMONMIB_Hosttable) GetFilter() yfilter.YFilter { return hosttable.YFilter }

func (hosttable *RMONMIB_Hosttable) SetFilter(yf yfilter.YFilter) { hosttable.YFilter = yf }

func (hosttable *RMONMIB_Hosttable) GetGoName(yname string) string {
    if yname == "hostEntry" { return "Hostentry" }
    return ""
}

func (hosttable *RMONMIB_Hosttable) GetSegmentPath() string {
    return "hostTable"
}

func (hosttable *RMONMIB_Hosttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hostEntry" {
        for _, c := range hosttable.Hostentry {
            if hosttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Hosttable_Hostentry{}
        hosttable.Hostentry = append(hosttable.Hostentry, child)
        return &hosttable.Hostentry[len(hosttable.Hostentry)-1]
    }
    return nil
}

func (hosttable *RMONMIB_Hosttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosttable.Hostentry {
        children[hosttable.Hostentry[i].GetSegmentPath()] = &hosttable.Hostentry[i]
    }
    return children
}

func (hosttable *RMONMIB_Hosttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosttable *RMONMIB_Hosttable) GetBundleName() string { return "cisco_ios_xe" }

func (hosttable *RMONMIB_Hosttable) GetYangName() string { return "hostTable" }

func (hosttable *RMONMIB_Hosttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttable *RMONMIB_Hosttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttable *RMONMIB_Hosttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttable *RMONMIB_Hosttable) SetParent(parent types.Entity) { hosttable.parent = parent }

func (hosttable *RMONMIB_Hosttable) GetParent() types.Entity { return hosttable.parent }

func (hosttable *RMONMIB_Hosttable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Hosttable_Hostentry
// A collection of statistics for a particular host that has
// been discovered on an interface of this device.  For example,
// an instance of the hostOutBroadcastPkts object might be
// named hostOutBroadcastPkts.1.6.8.0.32.27.3.176
type RMONMIB_Hosttable_Hostentry struct {
    parent types.Entity
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

func (hostentry *RMONMIB_Hosttable_Hostentry) GetFilter() yfilter.YFilter { return hostentry.YFilter }

func (hostentry *RMONMIB_Hosttable_Hostentry) SetFilter(yf yfilter.YFilter) { hostentry.YFilter = yf }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetGoName(yname string) string {
    if yname == "hostIndex" { return "Hostindex" }
    if yname == "hostAddress" { return "Hostaddress" }
    if yname == "hostCreationOrder" { return "Hostcreationorder" }
    if yname == "hostInPkts" { return "Hostinpkts" }
    if yname == "hostOutPkts" { return "Hostoutpkts" }
    if yname == "hostInOctets" { return "Hostinoctets" }
    if yname == "hostOutOctets" { return "Hostoutoctets" }
    if yname == "hostOutErrors" { return "Hostouterrors" }
    if yname == "hostOutBroadcastPkts" { return "Hostoutbroadcastpkts" }
    if yname == "hostOutMulticastPkts" { return "Hostoutmulticastpkts" }
    return ""
}

func (hostentry *RMONMIB_Hosttable_Hostentry) GetSegmentPath() string {
    return "hostEntry" + "[hostIndex='" + fmt.Sprintf("%v", hostentry.Hostindex) + "']" + "[hostAddress='" + fmt.Sprintf("%v", hostentry.Hostaddress) + "']"
}

func (hostentry *RMONMIB_Hosttable_Hostentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostentry *RMONMIB_Hosttable_Hostentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostentry *RMONMIB_Hosttable_Hostentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hostIndex"] = hostentry.Hostindex
    leafs["hostAddress"] = hostentry.Hostaddress
    leafs["hostCreationOrder"] = hostentry.Hostcreationorder
    leafs["hostInPkts"] = hostentry.Hostinpkts
    leafs["hostOutPkts"] = hostentry.Hostoutpkts
    leafs["hostInOctets"] = hostentry.Hostinoctets
    leafs["hostOutOctets"] = hostentry.Hostoutoctets
    leafs["hostOutErrors"] = hostentry.Hostouterrors
    leafs["hostOutBroadcastPkts"] = hostentry.Hostoutbroadcastpkts
    leafs["hostOutMulticastPkts"] = hostentry.Hostoutmulticastpkts
    return leafs
}

func (hostentry *RMONMIB_Hosttable_Hostentry) GetBundleName() string { return "cisco_ios_xe" }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetYangName() string { return "hostEntry" }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hostentry *RMONMIB_Hosttable_Hostentry) SetParent(parent types.Entity) { hostentry.parent = parent }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetParent() types.Entity { return hostentry.parent }

func (hostentry *RMONMIB_Hosttable_Hostentry) GetParentYangName() string { return "hostTable" }

// RMONMIB_Hosttimetable
// A list of time-ordered host table entries.
type RMONMIB_Hosttimetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for a particular host that has been discovered
    // on an interface of this device.  This collection includes the relative
    // ordering of the creation time of this object.  For example, an instance of
    // the hostTimeOutBroadcastPkts object might be named
    // hostTimeOutBroadcastPkts.1.687. The type is slice of
    // RMONMIB_Hosttimetable_Hosttimeentry.
    Hosttimeentry []RMONMIB_Hosttimetable_Hosttimeentry
}

func (hosttimetable *RMONMIB_Hosttimetable) GetFilter() yfilter.YFilter { return hosttimetable.YFilter }

func (hosttimetable *RMONMIB_Hosttimetable) SetFilter(yf yfilter.YFilter) { hosttimetable.YFilter = yf }

func (hosttimetable *RMONMIB_Hosttimetable) GetGoName(yname string) string {
    if yname == "hostTimeEntry" { return "Hosttimeentry" }
    return ""
}

func (hosttimetable *RMONMIB_Hosttimetable) GetSegmentPath() string {
    return "hostTimeTable"
}

func (hosttimetable *RMONMIB_Hosttimetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hostTimeEntry" {
        for _, c := range hosttimetable.Hosttimeentry {
            if hosttimetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Hosttimetable_Hosttimeentry{}
        hosttimetable.Hosttimeentry = append(hosttimetable.Hosttimeentry, child)
        return &hosttimetable.Hosttimeentry[len(hosttimetable.Hosttimeentry)-1]
    }
    return nil
}

func (hosttimetable *RMONMIB_Hosttimetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosttimetable.Hosttimeentry {
        children[hosttimetable.Hosttimeentry[i].GetSegmentPath()] = &hosttimetable.Hosttimeentry[i]
    }
    return children
}

func (hosttimetable *RMONMIB_Hosttimetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosttimetable *RMONMIB_Hosttimetable) GetBundleName() string { return "cisco_ios_xe" }

func (hosttimetable *RMONMIB_Hosttimetable) GetYangName() string { return "hostTimeTable" }

func (hosttimetable *RMONMIB_Hosttimetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttimetable *RMONMIB_Hosttimetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttimetable *RMONMIB_Hosttimetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttimetable *RMONMIB_Hosttimetable) SetParent(parent types.Entity) { hosttimetable.parent = parent }

func (hosttimetable *RMONMIB_Hosttimetable) GetParent() types.Entity { return hosttimetable.parent }

func (hosttimetable *RMONMIB_Hosttimetable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Hosttimetable_Hosttimeentry
// A collection of statistics for a particular host that has
// been discovered on an interface of this device.  This
// collection includes the relative ordering of the creation
// time of this object.  For example, an instance of the
// hostTimeOutBroadcastPkts object might be named
// hostTimeOutBroadcastPkts.1.687
type RMONMIB_Hosttimetable_Hosttimeentry struct {
    parent types.Entity
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

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetFilter() yfilter.YFilter { return hosttimeentry.YFilter }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) SetFilter(yf yfilter.YFilter) { hosttimeentry.YFilter = yf }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetGoName(yname string) string {
    if yname == "hostTimeIndex" { return "Hosttimeindex" }
    if yname == "hostTimeCreationOrder" { return "Hosttimecreationorder" }
    if yname == "hostTimeAddress" { return "Hosttimeaddress" }
    if yname == "hostTimeInPkts" { return "Hosttimeinpkts" }
    if yname == "hostTimeOutPkts" { return "Hosttimeoutpkts" }
    if yname == "hostTimeInOctets" { return "Hosttimeinoctets" }
    if yname == "hostTimeOutOctets" { return "Hosttimeoutoctets" }
    if yname == "hostTimeOutErrors" { return "Hosttimeouterrors" }
    if yname == "hostTimeOutBroadcastPkts" { return "Hosttimeoutbroadcastpkts" }
    if yname == "hostTimeOutMulticastPkts" { return "Hosttimeoutmulticastpkts" }
    return ""
}

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetSegmentPath() string {
    return "hostTimeEntry" + "[hostTimeIndex='" + fmt.Sprintf("%v", hosttimeentry.Hosttimeindex) + "']" + "[hostTimeCreationOrder='" + fmt.Sprintf("%v", hosttimeentry.Hosttimecreationorder) + "']"
}

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hostTimeIndex"] = hosttimeentry.Hosttimeindex
    leafs["hostTimeCreationOrder"] = hosttimeentry.Hosttimecreationorder
    leafs["hostTimeAddress"] = hosttimeentry.Hosttimeaddress
    leafs["hostTimeInPkts"] = hosttimeentry.Hosttimeinpkts
    leafs["hostTimeOutPkts"] = hosttimeentry.Hosttimeoutpkts
    leafs["hostTimeInOctets"] = hosttimeentry.Hosttimeinoctets
    leafs["hostTimeOutOctets"] = hosttimeentry.Hosttimeoutoctets
    leafs["hostTimeOutErrors"] = hosttimeentry.Hosttimeouterrors
    leafs["hostTimeOutBroadcastPkts"] = hosttimeentry.Hosttimeoutbroadcastpkts
    leafs["hostTimeOutMulticastPkts"] = hosttimeentry.Hosttimeoutmulticastpkts
    return leafs
}

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetBundleName() string { return "cisco_ios_xe" }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetYangName() string { return "hostTimeEntry" }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) SetParent(parent types.Entity) { hosttimeentry.parent = parent }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetParent() types.Entity { return hosttimeentry.parent }

func (hosttimeentry *RMONMIB_Hosttimetable_Hosttimeentry) GetParentYangName() string { return "hostTimeTable" }

// RMONMIB_Hosttopncontroltable
// A list of top N host control entries.
type RMONMIB_Hosttopncontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of parameters that control the creation of a report of the top N
    // hosts according to several metrics.  For example, an instance of the
    // hostTopNDuration object might be named hostTopNDuration.3. The type is
    // slice of RMONMIB_Hosttopncontroltable_Hosttopncontrolentry.
    Hosttopncontrolentry []RMONMIB_Hosttopncontroltable_Hosttopncontrolentry
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetFilter() yfilter.YFilter { return hosttopncontroltable.YFilter }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) SetFilter(yf yfilter.YFilter) { hosttopncontroltable.YFilter = yf }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetGoName(yname string) string {
    if yname == "hostTopNControlEntry" { return "Hosttopncontrolentry" }
    return ""
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetSegmentPath() string {
    return "hostTopNControlTable"
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hostTopNControlEntry" {
        for _, c := range hosttopncontroltable.Hosttopncontrolentry {
            if hosttopncontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Hosttopncontroltable_Hosttopncontrolentry{}
        hosttopncontroltable.Hosttopncontrolentry = append(hosttopncontroltable.Hosttopncontrolentry, child)
        return &hosttopncontroltable.Hosttopncontrolentry[len(hosttopncontroltable.Hosttopncontrolentry)-1]
    }
    return nil
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosttopncontroltable.Hosttopncontrolentry {
        children[hosttopncontroltable.Hosttopncontrolentry[i].GetSegmentPath()] = &hosttopncontroltable.Hosttopncontrolentry[i]
    }
    return children
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetYangName() string { return "hostTopNControlTable" }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) SetParent(parent types.Entity) { hosttopncontroltable.parent = parent }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetParent() types.Entity { return hosttopncontroltable.parent }

func (hosttopncontroltable *RMONMIB_Hosttopncontroltable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Hosttopncontroltable_Hosttopncontrolentry
// A set of parameters that control the creation of a report
// of the top N hosts according to several metrics.  For
// example, an instance of the hostTopNDuration object might
// be named hostTopNDuration.3
type RMONMIB_Hosttopncontroltable_Hosttopncontrolentry struct {
    parent types.Entity
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

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetFilter() yfilter.YFilter { return hosttopncontrolentry.YFilter }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) SetFilter(yf yfilter.YFilter) { hosttopncontrolentry.YFilter = yf }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetGoName(yname string) string {
    if yname == "hostTopNControlIndex" { return "Hosttopncontrolindex" }
    if yname == "hostTopNHostIndex" { return "Hosttopnhostindex" }
    if yname == "hostTopNRateBase" { return "Hosttopnratebase" }
    if yname == "hostTopNTimeRemaining" { return "Hosttopntimeremaining" }
    if yname == "hostTopNDuration" { return "Hosttopnduration" }
    if yname == "hostTopNRequestedSize" { return "Hosttopnrequestedsize" }
    if yname == "hostTopNGrantedSize" { return "Hosttopngrantedsize" }
    if yname == "hostTopNStartTime" { return "Hosttopnstarttime" }
    if yname == "hostTopNOwner" { return "Hosttopnowner" }
    if yname == "hostTopNStatus" { return "Hosttopnstatus" }
    return ""
}

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetSegmentPath() string {
    return "hostTopNControlEntry" + "[hostTopNControlIndex='" + fmt.Sprintf("%v", hosttopncontrolentry.Hosttopncontrolindex) + "']"
}

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hostTopNControlIndex"] = hosttopncontrolentry.Hosttopncontrolindex
    leafs["hostTopNHostIndex"] = hosttopncontrolentry.Hosttopnhostindex
    leafs["hostTopNRateBase"] = hosttopncontrolentry.Hosttopnratebase
    leafs["hostTopNTimeRemaining"] = hosttopncontrolentry.Hosttopntimeremaining
    leafs["hostTopNDuration"] = hosttopncontrolentry.Hosttopnduration
    leafs["hostTopNRequestedSize"] = hosttopncontrolentry.Hosttopnrequestedsize
    leafs["hostTopNGrantedSize"] = hosttopncontrolentry.Hosttopngrantedsize
    leafs["hostTopNStartTime"] = hosttopncontrolentry.Hosttopnstarttime
    leafs["hostTopNOwner"] = hosttopncontrolentry.Hosttopnowner
    leafs["hostTopNStatus"] = hosttopncontrolentry.Hosttopnstatus
    return leafs
}

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetYangName() string { return "hostTopNControlEntry" }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) SetParent(parent types.Entity) { hosttopncontrolentry.parent = parent }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetParent() types.Entity { return hosttopncontrolentry.parent }

func (hosttopncontrolentry *RMONMIB_Hosttopncontroltable_Hosttopncontrolentry) GetParentYangName() string { return "hostTopNControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of statistics for a host that is part of a top N report.  For
    // example, an instance of the hostTopNRate object might be named
    // hostTopNRate.3.10. The type is slice of
    // RMONMIB_Hosttopntable_Hosttopnentry.
    Hosttopnentry []RMONMIB_Hosttopntable_Hosttopnentry
}

func (hosttopntable *RMONMIB_Hosttopntable) GetFilter() yfilter.YFilter { return hosttopntable.YFilter }

func (hosttopntable *RMONMIB_Hosttopntable) SetFilter(yf yfilter.YFilter) { hosttopntable.YFilter = yf }

func (hosttopntable *RMONMIB_Hosttopntable) GetGoName(yname string) string {
    if yname == "hostTopNEntry" { return "Hosttopnentry" }
    return ""
}

func (hosttopntable *RMONMIB_Hosttopntable) GetSegmentPath() string {
    return "hostTopNTable"
}

func (hosttopntable *RMONMIB_Hosttopntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hostTopNEntry" {
        for _, c := range hosttopntable.Hosttopnentry {
            if hosttopntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Hosttopntable_Hosttopnentry{}
        hosttopntable.Hosttopnentry = append(hosttopntable.Hosttopnentry, child)
        return &hosttopntable.Hosttopnentry[len(hosttopntable.Hosttopnentry)-1]
    }
    return nil
}

func (hosttopntable *RMONMIB_Hosttopntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosttopntable.Hosttopnentry {
        children[hosttopntable.Hosttopnentry[i].GetSegmentPath()] = &hosttopntable.Hosttopnentry[i]
    }
    return children
}

func (hosttopntable *RMONMIB_Hosttopntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosttopntable *RMONMIB_Hosttopntable) GetBundleName() string { return "cisco_ios_xe" }

func (hosttopntable *RMONMIB_Hosttopntable) GetYangName() string { return "hostTopNTable" }

func (hosttopntable *RMONMIB_Hosttopntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttopntable *RMONMIB_Hosttopntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttopntable *RMONMIB_Hosttopntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttopntable *RMONMIB_Hosttopntable) SetParent(parent types.Entity) { hosttopntable.parent = parent }

func (hosttopntable *RMONMIB_Hosttopntable) GetParent() types.Entity { return hosttopntable.parent }

func (hosttopntable *RMONMIB_Hosttopntable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Hosttopntable_Hosttopnentry
// A set of statistics for a host that is part of a top N
// report.  For example, an instance of the hostTopNRate
// object might be named hostTopNRate.3.10
type RMONMIB_Hosttopntable_Hosttopnentry struct {
    parent types.Entity
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

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetFilter() yfilter.YFilter { return hosttopnentry.YFilter }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) SetFilter(yf yfilter.YFilter) { hosttopnentry.YFilter = yf }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetGoName(yname string) string {
    if yname == "hostTopNReport" { return "Hosttopnreport" }
    if yname == "hostTopNIndex" { return "Hosttopnindex" }
    if yname == "hostTopNAddress" { return "Hosttopnaddress" }
    if yname == "hostTopNRate" { return "Hosttopnrate" }
    return ""
}

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetSegmentPath() string {
    return "hostTopNEntry" + "[hostTopNReport='" + fmt.Sprintf("%v", hosttopnentry.Hosttopnreport) + "']" + "[hostTopNIndex='" + fmt.Sprintf("%v", hosttopnentry.Hosttopnindex) + "']"
}

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hostTopNReport"] = hosttopnentry.Hosttopnreport
    leafs["hostTopNIndex"] = hosttopnentry.Hosttopnindex
    leafs["hostTopNAddress"] = hosttopnentry.Hosttopnaddress
    leafs["hostTopNRate"] = hosttopnentry.Hosttopnrate
    return leafs
}

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetBundleName() string { return "cisco_ios_xe" }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetYangName() string { return "hostTopNEntry" }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) SetParent(parent types.Entity) { hosttopnentry.parent = parent }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetParent() types.Entity { return hosttopnentry.parent }

func (hosttopnentry *RMONMIB_Hosttopntable_Hosttopnentry) GetParentYangName() string { return "hostTopNTable" }

// RMONMIB_Matrixcontroltable
// A list of information entries for the
// traffic matrix on each interface.
type RMONMIB_Matrixcontroltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a traffic matrix on a particular interface.  For example,
    // an instance of the matrixControlLastDeleteTime object might be named
    // matrixControlLastDeleteTime.1. The type is slice of
    // RMONMIB_Matrixcontroltable_Matrixcontrolentry.
    Matrixcontrolentry []RMONMIB_Matrixcontroltable_Matrixcontrolentry
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetFilter() yfilter.YFilter { return matrixcontroltable.YFilter }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) SetFilter(yf yfilter.YFilter) { matrixcontroltable.YFilter = yf }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetGoName(yname string) string {
    if yname == "matrixControlEntry" { return "Matrixcontrolentry" }
    return ""
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetSegmentPath() string {
    return "matrixControlTable"
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "matrixControlEntry" {
        for _, c := range matrixcontroltable.Matrixcontrolentry {
            if matrixcontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Matrixcontroltable_Matrixcontrolentry{}
        matrixcontroltable.Matrixcontrolentry = append(matrixcontroltable.Matrixcontrolentry, child)
        return &matrixcontroltable.Matrixcontrolentry[len(matrixcontroltable.Matrixcontrolentry)-1]
    }
    return nil
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range matrixcontroltable.Matrixcontrolentry {
        children[matrixcontroltable.Matrixcontrolentry[i].GetSegmentPath()] = &matrixcontroltable.Matrixcontrolentry[i]
    }
    return children
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetYangName() string { return "matrixControlTable" }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) SetParent(parent types.Entity) { matrixcontroltable.parent = parent }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetParent() types.Entity { return matrixcontroltable.parent }

func (matrixcontroltable *RMONMIB_Matrixcontroltable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Matrixcontroltable_Matrixcontrolentry
// Information about a traffic matrix on a particular
// interface.  For example, an instance of the
// matrixControlLastDeleteTime object might be named
// matrixControlLastDeleteTime.1
type RMONMIB_Matrixcontroltable_Matrixcontrolentry struct {
    parent types.Entity
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetFilter() yfilter.YFilter { return matrixcontrolentry.YFilter }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) SetFilter(yf yfilter.YFilter) { matrixcontrolentry.YFilter = yf }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetGoName(yname string) string {
    if yname == "matrixControlIndex" { return "Matrixcontrolindex" }
    if yname == "matrixControlDataSource" { return "Matrixcontroldatasource" }
    if yname == "matrixControlTableSize" { return "Matrixcontroltablesize" }
    if yname == "matrixControlLastDeleteTime" { return "Matrixcontrollastdeletetime" }
    if yname == "matrixControlOwner" { return "Matrixcontrolowner" }
    if yname == "matrixControlStatus" { return "Matrixcontrolstatus" }
    if yname == "matrixControlDroppedFrames" { return "Matrixcontroldroppedframes" }
    if yname == "matrixControlCreateTime" { return "Matrixcontrolcreatetime" }
    return ""
}

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetSegmentPath() string {
    return "matrixControlEntry" + "[matrixControlIndex='" + fmt.Sprintf("%v", matrixcontrolentry.Matrixcontrolindex) + "']"
}

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matrixControlIndex"] = matrixcontrolentry.Matrixcontrolindex
    leafs["matrixControlDataSource"] = matrixcontrolentry.Matrixcontroldatasource
    leafs["matrixControlTableSize"] = matrixcontrolentry.Matrixcontroltablesize
    leafs["matrixControlLastDeleteTime"] = matrixcontrolentry.Matrixcontrollastdeletetime
    leafs["matrixControlOwner"] = matrixcontrolentry.Matrixcontrolowner
    leafs["matrixControlStatus"] = matrixcontrolentry.Matrixcontrolstatus
    leafs["matrixControlDroppedFrames"] = matrixcontrolentry.Matrixcontroldroppedframes
    leafs["matrixControlCreateTime"] = matrixcontrolentry.Matrixcontrolcreatetime
    return leafs
}

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetYangName() string { return "matrixControlEntry" }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) SetParent(parent types.Entity) { matrixcontrolentry.parent = parent }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetParent() types.Entity { return matrixcontrolentry.parent }

func (matrixcontrolentry *RMONMIB_Matrixcontroltable_Matrixcontrolentry) GetParentYangName() string { return "matrixControlTable" }

// RMONMIB_Matrixsdtable
// A list of traffic matrix entries indexed by
// source and destination MAC address.
type RMONMIB_Matrixsdtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for communications between two addresses on a
    // particular interface.  For example, an instance of the matrixSDPkts object
    // might be named matrixSDPkts.1.6.8.0.32.27.3.176.6.8.0.32.10.8.113. The type
    // is slice of RMONMIB_Matrixsdtable_Matrixsdentry.
    Matrixsdentry []RMONMIB_Matrixsdtable_Matrixsdentry
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetFilter() yfilter.YFilter { return matrixsdtable.YFilter }

func (matrixsdtable *RMONMIB_Matrixsdtable) SetFilter(yf yfilter.YFilter) { matrixsdtable.YFilter = yf }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetGoName(yname string) string {
    if yname == "matrixSDEntry" { return "Matrixsdentry" }
    return ""
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetSegmentPath() string {
    return "matrixSDTable"
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "matrixSDEntry" {
        for _, c := range matrixsdtable.Matrixsdentry {
            if matrixsdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Matrixsdtable_Matrixsdentry{}
        matrixsdtable.Matrixsdentry = append(matrixsdtable.Matrixsdentry, child)
        return &matrixsdtable.Matrixsdentry[len(matrixsdtable.Matrixsdentry)-1]
    }
    return nil
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range matrixsdtable.Matrixsdentry {
        children[matrixsdtable.Matrixsdentry[i].GetSegmentPath()] = &matrixsdtable.Matrixsdentry[i]
    }
    return children
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matrixsdtable *RMONMIB_Matrixsdtable) GetBundleName() string { return "cisco_ios_xe" }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetYangName() string { return "matrixSDTable" }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (matrixsdtable *RMONMIB_Matrixsdtable) SetParent(parent types.Entity) { matrixsdtable.parent = parent }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetParent() types.Entity { return matrixsdtable.parent }

func (matrixsdtable *RMONMIB_Matrixsdtable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Matrixsdtable_Matrixsdentry
// A collection of statistics for communications between
// two addresses on a particular interface.  For example,
// an instance of the matrixSDPkts object might be named
// matrixSDPkts.1.6.8.0.32.27.3.176.6.8.0.32.10.8.113
type RMONMIB_Matrixsdtable_Matrixsdentry struct {
    parent types.Entity
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

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetFilter() yfilter.YFilter { return matrixsdentry.YFilter }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) SetFilter(yf yfilter.YFilter) { matrixsdentry.YFilter = yf }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetGoName(yname string) string {
    if yname == "matrixSDIndex" { return "Matrixsdindex" }
    if yname == "matrixSDSourceAddress" { return "Matrixsdsourceaddress" }
    if yname == "matrixSDDestAddress" { return "Matrixsddestaddress" }
    if yname == "matrixSDPkts" { return "Matrixsdpkts" }
    if yname == "matrixSDOctets" { return "Matrixsdoctets" }
    if yname == "matrixSDErrors" { return "Matrixsderrors" }
    return ""
}

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetSegmentPath() string {
    return "matrixSDEntry" + "[matrixSDIndex='" + fmt.Sprintf("%v", matrixsdentry.Matrixsdindex) + "']" + "[matrixSDSourceAddress='" + fmt.Sprintf("%v", matrixsdentry.Matrixsdsourceaddress) + "']" + "[matrixSDDestAddress='" + fmt.Sprintf("%v", matrixsdentry.Matrixsddestaddress) + "']"
}

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matrixSDIndex"] = matrixsdentry.Matrixsdindex
    leafs["matrixSDSourceAddress"] = matrixsdentry.Matrixsdsourceaddress
    leafs["matrixSDDestAddress"] = matrixsdentry.Matrixsddestaddress
    leafs["matrixSDPkts"] = matrixsdentry.Matrixsdpkts
    leafs["matrixSDOctets"] = matrixsdentry.Matrixsdoctets
    leafs["matrixSDErrors"] = matrixsdentry.Matrixsderrors
    return leafs
}

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetBundleName() string { return "cisco_ios_xe" }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetYangName() string { return "matrixSDEntry" }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) SetParent(parent types.Entity) { matrixsdentry.parent = parent }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetParent() types.Entity { return matrixsdentry.parent }

func (matrixsdentry *RMONMIB_Matrixsdtable_Matrixsdentry) GetParentYangName() string { return "matrixSDTable" }

// RMONMIB_Matrixdstable
// A list of traffic matrix entries indexed by
// destination and source MAC address.
type RMONMIB_Matrixdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of statistics for communications between two addresses on a
    // particular interface.  For example, an instance of the matrixSDPkts object
    // might be named matrixSDPkts.1.6.8.0.32.10.8.113.6.8.0.32.27.3.176. The type
    // is slice of RMONMIB_Matrixdstable_Matrixdsentry.
    Matrixdsentry []RMONMIB_Matrixdstable_Matrixdsentry
}

func (matrixdstable *RMONMIB_Matrixdstable) GetFilter() yfilter.YFilter { return matrixdstable.YFilter }

func (matrixdstable *RMONMIB_Matrixdstable) SetFilter(yf yfilter.YFilter) { matrixdstable.YFilter = yf }

func (matrixdstable *RMONMIB_Matrixdstable) GetGoName(yname string) string {
    if yname == "matrixDSEntry" { return "Matrixdsentry" }
    return ""
}

func (matrixdstable *RMONMIB_Matrixdstable) GetSegmentPath() string {
    return "matrixDSTable"
}

func (matrixdstable *RMONMIB_Matrixdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "matrixDSEntry" {
        for _, c := range matrixdstable.Matrixdsentry {
            if matrixdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Matrixdstable_Matrixdsentry{}
        matrixdstable.Matrixdsentry = append(matrixdstable.Matrixdsentry, child)
        return &matrixdstable.Matrixdsentry[len(matrixdstable.Matrixdsentry)-1]
    }
    return nil
}

func (matrixdstable *RMONMIB_Matrixdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range matrixdstable.Matrixdsentry {
        children[matrixdstable.Matrixdsentry[i].GetSegmentPath()] = &matrixdstable.Matrixdsentry[i]
    }
    return children
}

func (matrixdstable *RMONMIB_Matrixdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matrixdstable *RMONMIB_Matrixdstable) GetBundleName() string { return "cisco_ios_xe" }

func (matrixdstable *RMONMIB_Matrixdstable) GetYangName() string { return "matrixDSTable" }

func (matrixdstable *RMONMIB_Matrixdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (matrixdstable *RMONMIB_Matrixdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (matrixdstable *RMONMIB_Matrixdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (matrixdstable *RMONMIB_Matrixdstable) SetParent(parent types.Entity) { matrixdstable.parent = parent }

func (matrixdstable *RMONMIB_Matrixdstable) GetParent() types.Entity { return matrixdstable.parent }

func (matrixdstable *RMONMIB_Matrixdstable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Matrixdstable_Matrixdsentry
// A collection of statistics for communications between
// two addresses on a particular interface.  For example,
// an instance of the matrixSDPkts object might be named
// matrixSDPkts.1.6.8.0.32.10.8.113.6.8.0.32.27.3.176
type RMONMIB_Matrixdstable_Matrixdsentry struct {
    parent types.Entity
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

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetFilter() yfilter.YFilter { return matrixdsentry.YFilter }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) SetFilter(yf yfilter.YFilter) { matrixdsentry.YFilter = yf }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetGoName(yname string) string {
    if yname == "matrixDSIndex" { return "Matrixdsindex" }
    if yname == "matrixDSDestAddress" { return "Matrixdsdestaddress" }
    if yname == "matrixDSSourceAddress" { return "Matrixdssourceaddress" }
    if yname == "matrixDSPkts" { return "Matrixdspkts" }
    if yname == "matrixDSOctets" { return "Matrixdsoctets" }
    if yname == "matrixDSErrors" { return "Matrixdserrors" }
    return ""
}

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetSegmentPath() string {
    return "matrixDSEntry" + "[matrixDSIndex='" + fmt.Sprintf("%v", matrixdsentry.Matrixdsindex) + "']" + "[matrixDSDestAddress='" + fmt.Sprintf("%v", matrixdsentry.Matrixdsdestaddress) + "']" + "[matrixDSSourceAddress='" + fmt.Sprintf("%v", matrixdsentry.Matrixdssourceaddress) + "']"
}

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matrixDSIndex"] = matrixdsentry.Matrixdsindex
    leafs["matrixDSDestAddress"] = matrixdsentry.Matrixdsdestaddress
    leafs["matrixDSSourceAddress"] = matrixdsentry.Matrixdssourceaddress
    leafs["matrixDSPkts"] = matrixdsentry.Matrixdspkts
    leafs["matrixDSOctets"] = matrixdsentry.Matrixdsoctets
    leafs["matrixDSErrors"] = matrixdsentry.Matrixdserrors
    return leafs
}

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetYangName() string { return "matrixDSEntry" }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) SetParent(parent types.Entity) { matrixdsentry.parent = parent }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetParent() types.Entity { return matrixdsentry.parent }

func (matrixdsentry *RMONMIB_Matrixdstable_Matrixdsentry) GetParentYangName() string { return "matrixDSTable" }

// RMONMIB_Filtertable
// A list of packet filter entries.
type RMONMIB_Filtertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of parameters for a packet filter applied on a particular interface. 
    // As an example, an instance of the filterPktData object might be named
    // filterPktData.12. The type is slice of RMONMIB_Filtertable_Filterentry.
    Filterentry []RMONMIB_Filtertable_Filterentry
}

func (filtertable *RMONMIB_Filtertable) GetFilter() yfilter.YFilter { return filtertable.YFilter }

func (filtertable *RMONMIB_Filtertable) SetFilter(yf yfilter.YFilter) { filtertable.YFilter = yf }

func (filtertable *RMONMIB_Filtertable) GetGoName(yname string) string {
    if yname == "filterEntry" { return "Filterentry" }
    return ""
}

func (filtertable *RMONMIB_Filtertable) GetSegmentPath() string {
    return "filterTable"
}

func (filtertable *RMONMIB_Filtertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "filterEntry" {
        for _, c := range filtertable.Filterentry {
            if filtertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Filtertable_Filterentry{}
        filtertable.Filterentry = append(filtertable.Filterentry, child)
        return &filtertable.Filterentry[len(filtertable.Filterentry)-1]
    }
    return nil
}

func (filtertable *RMONMIB_Filtertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range filtertable.Filterentry {
        children[filtertable.Filterentry[i].GetSegmentPath()] = &filtertable.Filterentry[i]
    }
    return children
}

func (filtertable *RMONMIB_Filtertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (filtertable *RMONMIB_Filtertable) GetBundleName() string { return "cisco_ios_xe" }

func (filtertable *RMONMIB_Filtertable) GetYangName() string { return "filterTable" }

func (filtertable *RMONMIB_Filtertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (filtertable *RMONMIB_Filtertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (filtertable *RMONMIB_Filtertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (filtertable *RMONMIB_Filtertable) SetParent(parent types.Entity) { filtertable.parent = parent }

func (filtertable *RMONMIB_Filtertable) GetParent() types.Entity { return filtertable.parent }

func (filtertable *RMONMIB_Filtertable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Filtertable_Filterentry
// A set of parameters for a packet filter applied on a
// particular interface.  As an example, an instance of the
// filterPktData object might be named filterPktData.12
type RMONMIB_Filtertable_Filterentry struct {
    parent types.Entity
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

func (filterentry *RMONMIB_Filtertable_Filterentry) GetFilter() yfilter.YFilter { return filterentry.YFilter }

func (filterentry *RMONMIB_Filtertable_Filterentry) SetFilter(yf yfilter.YFilter) { filterentry.YFilter = yf }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetGoName(yname string) string {
    if yname == "filterIndex" { return "Filterindex" }
    if yname == "filterChannelIndex" { return "Filterchannelindex" }
    if yname == "filterPktDataOffset" { return "Filterpktdataoffset" }
    if yname == "filterPktData" { return "Filterpktdata" }
    if yname == "filterPktDataMask" { return "Filterpktdatamask" }
    if yname == "filterPktDataNotMask" { return "Filterpktdatanotmask" }
    if yname == "filterPktStatus" { return "Filterpktstatus" }
    if yname == "filterPktStatusMask" { return "Filterpktstatusmask" }
    if yname == "filterPktStatusNotMask" { return "Filterpktstatusnotmask" }
    if yname == "filterOwner" { return "Filterowner" }
    if yname == "filterStatus" { return "Filterstatus" }
    if yname == "filterProtocolDirDataLocalIndex" { return "Filterprotocoldirdatalocalindex" }
    if yname == "filterProtocolDirLocalIndex" { return "Filterprotocoldirlocalindex" }
    return ""
}

func (filterentry *RMONMIB_Filtertable_Filterentry) GetSegmentPath() string {
    return "filterEntry" + "[filterIndex='" + fmt.Sprintf("%v", filterentry.Filterindex) + "']"
}

func (filterentry *RMONMIB_Filtertable_Filterentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (filterentry *RMONMIB_Filtertable_Filterentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (filterentry *RMONMIB_Filtertable_Filterentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filterIndex"] = filterentry.Filterindex
    leafs["filterChannelIndex"] = filterentry.Filterchannelindex
    leafs["filterPktDataOffset"] = filterentry.Filterpktdataoffset
    leafs["filterPktData"] = filterentry.Filterpktdata
    leafs["filterPktDataMask"] = filterentry.Filterpktdatamask
    leafs["filterPktDataNotMask"] = filterentry.Filterpktdatanotmask
    leafs["filterPktStatus"] = filterentry.Filterpktstatus
    leafs["filterPktStatusMask"] = filterentry.Filterpktstatusmask
    leafs["filterPktStatusNotMask"] = filterentry.Filterpktstatusnotmask
    leafs["filterOwner"] = filterentry.Filterowner
    leafs["filterStatus"] = filterentry.Filterstatus
    leafs["filterProtocolDirDataLocalIndex"] = filterentry.Filterprotocoldirdatalocalindex
    leafs["filterProtocolDirLocalIndex"] = filterentry.Filterprotocoldirlocalindex
    return leafs
}

func (filterentry *RMONMIB_Filtertable_Filterentry) GetBundleName() string { return "cisco_ios_xe" }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetYangName() string { return "filterEntry" }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (filterentry *RMONMIB_Filtertable_Filterentry) SetParent(parent types.Entity) { filterentry.parent = parent }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetParent() types.Entity { return filterentry.parent }

func (filterentry *RMONMIB_Filtertable_Filterentry) GetParentYangName() string { return "filterTable" }

// RMONMIB_Channeltable
// A list of packet channel entries.
type RMONMIB_Channeltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of parameters for a packet channel applied on a particular interface.
    // As an example, an instance of the channelMatches object might be named
    // channelMatches.3. The type is slice of RMONMIB_Channeltable_Channelentry.
    Channelentry []RMONMIB_Channeltable_Channelentry
}

func (channeltable *RMONMIB_Channeltable) GetFilter() yfilter.YFilter { return channeltable.YFilter }

func (channeltable *RMONMIB_Channeltable) SetFilter(yf yfilter.YFilter) { channeltable.YFilter = yf }

func (channeltable *RMONMIB_Channeltable) GetGoName(yname string) string {
    if yname == "channelEntry" { return "Channelentry" }
    return ""
}

func (channeltable *RMONMIB_Channeltable) GetSegmentPath() string {
    return "channelTable"
}

func (channeltable *RMONMIB_Channeltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channelEntry" {
        for _, c := range channeltable.Channelentry {
            if channeltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Channeltable_Channelentry{}
        channeltable.Channelentry = append(channeltable.Channelentry, child)
        return &channeltable.Channelentry[len(channeltable.Channelentry)-1]
    }
    return nil
}

func (channeltable *RMONMIB_Channeltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range channeltable.Channelentry {
        children[channeltable.Channelentry[i].GetSegmentPath()] = &channeltable.Channelentry[i]
    }
    return children
}

func (channeltable *RMONMIB_Channeltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (channeltable *RMONMIB_Channeltable) GetBundleName() string { return "cisco_ios_xe" }

func (channeltable *RMONMIB_Channeltable) GetYangName() string { return "channelTable" }

func (channeltable *RMONMIB_Channeltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (channeltable *RMONMIB_Channeltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (channeltable *RMONMIB_Channeltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (channeltable *RMONMIB_Channeltable) SetParent(parent types.Entity) { channeltable.parent = parent }

func (channeltable *RMONMIB_Channeltable) GetParent() types.Entity { return channeltable.parent }

func (channeltable *RMONMIB_Channeltable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Channeltable_Channelentry
// A set of parameters for a packet channel applied on a
// particular interface.  As an example, an instance of the
// channelMatches object might be named channelMatches.3
type RMONMIB_Channeltable_Channelentry struct {
    parent types.Entity
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

func (channelentry *RMONMIB_Channeltable_Channelentry) GetFilter() yfilter.YFilter { return channelentry.YFilter }

func (channelentry *RMONMIB_Channeltable_Channelentry) SetFilter(yf yfilter.YFilter) { channelentry.YFilter = yf }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetGoName(yname string) string {
    if yname == "channelIndex" { return "Channelindex" }
    if yname == "channelIfIndex" { return "Channelifindex" }
    if yname == "channelAcceptType" { return "Channelaccepttype" }
    if yname == "channelDataControl" { return "Channeldatacontrol" }
    if yname == "channelTurnOnEventIndex" { return "Channelturnoneventindex" }
    if yname == "channelTurnOffEventIndex" { return "Channelturnoffeventindex" }
    if yname == "channelEventIndex" { return "Channeleventindex" }
    if yname == "channelEventStatus" { return "Channeleventstatus" }
    if yname == "channelMatches" { return "Channelmatches" }
    if yname == "channelDescription" { return "Channeldescription" }
    if yname == "channelOwner" { return "Channelowner" }
    if yname == "channelStatus" { return "Channelstatus" }
    if yname == "channelDroppedFrames" { return "Channeldroppedframes" }
    if yname == "channelCreateTime" { return "Channelcreatetime" }
    return ""
}

func (channelentry *RMONMIB_Channeltable_Channelentry) GetSegmentPath() string {
    return "channelEntry" + "[channelIndex='" + fmt.Sprintf("%v", channelentry.Channelindex) + "']"
}

func (channelentry *RMONMIB_Channeltable_Channelentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (channelentry *RMONMIB_Channeltable_Channelentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (channelentry *RMONMIB_Channeltable_Channelentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["channelIndex"] = channelentry.Channelindex
    leafs["channelIfIndex"] = channelentry.Channelifindex
    leafs["channelAcceptType"] = channelentry.Channelaccepttype
    leafs["channelDataControl"] = channelentry.Channeldatacontrol
    leafs["channelTurnOnEventIndex"] = channelentry.Channelturnoneventindex
    leafs["channelTurnOffEventIndex"] = channelentry.Channelturnoffeventindex
    leafs["channelEventIndex"] = channelentry.Channeleventindex
    leafs["channelEventStatus"] = channelentry.Channeleventstatus
    leafs["channelMatches"] = channelentry.Channelmatches
    leafs["channelDescription"] = channelentry.Channeldescription
    leafs["channelOwner"] = channelentry.Channelowner
    leafs["channelStatus"] = channelentry.Channelstatus
    leafs["channelDroppedFrames"] = channelentry.Channeldroppedframes
    leafs["channelCreateTime"] = channelentry.Channelcreatetime
    return leafs
}

func (channelentry *RMONMIB_Channeltable_Channelentry) GetBundleName() string { return "cisco_ios_xe" }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetYangName() string { return "channelEntry" }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (channelentry *RMONMIB_Channeltable_Channelentry) SetParent(parent types.Entity) { channelentry.parent = parent }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetParent() types.Entity { return channelentry.parent }

func (channelentry *RMONMIB_Channeltable_Channelentry) GetParentYangName() string { return "channelTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of parameters that control the collection of a stream of packets that
    // have matched filters.  As an example, an instance of the
    // bufferControlCaptureSliceSize object might be named
    // bufferControlCaptureSliceSize.3. The type is slice of
    // RMONMIB_Buffercontroltable_Buffercontrolentry.
    Buffercontrolentry []RMONMIB_Buffercontroltable_Buffercontrolentry
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetFilter() yfilter.YFilter { return buffercontroltable.YFilter }

func (buffercontroltable *RMONMIB_Buffercontroltable) SetFilter(yf yfilter.YFilter) { buffercontroltable.YFilter = yf }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetGoName(yname string) string {
    if yname == "bufferControlEntry" { return "Buffercontrolentry" }
    return ""
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetSegmentPath() string {
    return "bufferControlTable"
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bufferControlEntry" {
        for _, c := range buffercontroltable.Buffercontrolentry {
            if buffercontroltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Buffercontroltable_Buffercontrolentry{}
        buffercontroltable.Buffercontrolentry = append(buffercontroltable.Buffercontrolentry, child)
        return &buffercontroltable.Buffercontrolentry[len(buffercontroltable.Buffercontrolentry)-1]
    }
    return nil
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range buffercontroltable.Buffercontrolentry {
        children[buffercontroltable.Buffercontrolentry[i].GetSegmentPath()] = &buffercontroltable.Buffercontrolentry[i]
    }
    return children
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (buffercontroltable *RMONMIB_Buffercontroltable) GetBundleName() string { return "cisco_ios_xe" }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetYangName() string { return "bufferControlTable" }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (buffercontroltable *RMONMIB_Buffercontroltable) SetParent(parent types.Entity) { buffercontroltable.parent = parent }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetParent() types.Entity { return buffercontroltable.parent }

func (buffercontroltable *RMONMIB_Buffercontroltable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Buffercontroltable_Buffercontrolentry
// A set of parameters that control the collection of a stream
// of packets that have matched filters.  As an example, an
// instance of the bufferControlCaptureSliceSize object might
// be named bufferControlCaptureSliceSize.3
type RMONMIB_Buffercontroltable_Buffercontrolentry struct {
    parent types.Entity
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

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetFilter() yfilter.YFilter { return buffercontrolentry.YFilter }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) SetFilter(yf yfilter.YFilter) { buffercontrolentry.YFilter = yf }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetGoName(yname string) string {
    if yname == "bufferControlIndex" { return "Buffercontrolindex" }
    if yname == "bufferControlChannelIndex" { return "Buffercontrolchannelindex" }
    if yname == "bufferControlFullStatus" { return "Buffercontrolfullstatus" }
    if yname == "bufferControlFullAction" { return "Buffercontrolfullaction" }
    if yname == "bufferControlCaptureSliceSize" { return "Buffercontrolcaptureslicesize" }
    if yname == "bufferControlDownloadSliceSize" { return "Buffercontroldownloadslicesize" }
    if yname == "bufferControlDownloadOffset" { return "Buffercontroldownloadoffset" }
    if yname == "bufferControlMaxOctetsRequested" { return "Buffercontrolmaxoctetsrequested" }
    if yname == "bufferControlMaxOctetsGranted" { return "Buffercontrolmaxoctetsgranted" }
    if yname == "bufferControlCapturedPackets" { return "Buffercontrolcapturedpackets" }
    if yname == "bufferControlTurnOnTime" { return "Buffercontrolturnontime" }
    if yname == "bufferControlOwner" { return "Buffercontrolowner" }
    if yname == "bufferControlStatus" { return "Buffercontrolstatus" }
    return ""
}

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetSegmentPath() string {
    return "bufferControlEntry" + "[bufferControlIndex='" + fmt.Sprintf("%v", buffercontrolentry.Buffercontrolindex) + "']"
}

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bufferControlIndex"] = buffercontrolentry.Buffercontrolindex
    leafs["bufferControlChannelIndex"] = buffercontrolentry.Buffercontrolchannelindex
    leafs["bufferControlFullStatus"] = buffercontrolentry.Buffercontrolfullstatus
    leafs["bufferControlFullAction"] = buffercontrolentry.Buffercontrolfullaction
    leafs["bufferControlCaptureSliceSize"] = buffercontrolentry.Buffercontrolcaptureslicesize
    leafs["bufferControlDownloadSliceSize"] = buffercontrolentry.Buffercontroldownloadslicesize
    leafs["bufferControlDownloadOffset"] = buffercontrolentry.Buffercontroldownloadoffset
    leafs["bufferControlMaxOctetsRequested"] = buffercontrolentry.Buffercontrolmaxoctetsrequested
    leafs["bufferControlMaxOctetsGranted"] = buffercontrolentry.Buffercontrolmaxoctetsgranted
    leafs["bufferControlCapturedPackets"] = buffercontrolentry.Buffercontrolcapturedpackets
    leafs["bufferControlTurnOnTime"] = buffercontrolentry.Buffercontrolturnontime
    leafs["bufferControlOwner"] = buffercontrolentry.Buffercontrolowner
    leafs["bufferControlStatus"] = buffercontrolentry.Buffercontrolstatus
    return leafs
}

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetBundleName() string { return "cisco_ios_xe" }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetYangName() string { return "bufferControlEntry" }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) SetParent(parent types.Entity) { buffercontrolentry.parent = parent }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetParent() types.Entity { return buffercontrolentry.parent }

func (buffercontrolentry *RMONMIB_Buffercontroltable_Buffercontrolentry) GetParentYangName() string { return "bufferControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A packet captured off of an attached network.  As an example, an instance
    // of the captureBufferPacketData object might be named
    // captureBufferPacketData.3.1783. The type is slice of
    // RMONMIB_Capturebuffertable_Capturebufferentry.
    Capturebufferentry []RMONMIB_Capturebuffertable_Capturebufferentry
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetFilter() yfilter.YFilter { return capturebuffertable.YFilter }

func (capturebuffertable *RMONMIB_Capturebuffertable) SetFilter(yf yfilter.YFilter) { capturebuffertable.YFilter = yf }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetGoName(yname string) string {
    if yname == "captureBufferEntry" { return "Capturebufferentry" }
    return ""
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetSegmentPath() string {
    return "captureBufferTable"
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "captureBufferEntry" {
        for _, c := range capturebuffertable.Capturebufferentry {
            if capturebuffertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Capturebuffertable_Capturebufferentry{}
        capturebuffertable.Capturebufferentry = append(capturebuffertable.Capturebufferentry, child)
        return &capturebuffertable.Capturebufferentry[len(capturebuffertable.Capturebufferentry)-1]
    }
    return nil
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range capturebuffertable.Capturebufferentry {
        children[capturebuffertable.Capturebufferentry[i].GetSegmentPath()] = &capturebuffertable.Capturebufferentry[i]
    }
    return children
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capturebuffertable *RMONMIB_Capturebuffertable) GetBundleName() string { return "cisco_ios_xe" }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetYangName() string { return "captureBufferTable" }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (capturebuffertable *RMONMIB_Capturebuffertable) SetParent(parent types.Entity) { capturebuffertable.parent = parent }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetParent() types.Entity { return capturebuffertable.parent }

func (capturebuffertable *RMONMIB_Capturebuffertable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Capturebuffertable_Capturebufferentry
// A packet captured off of an attached network.  As an
// example, an instance of the captureBufferPacketData
// object might be named captureBufferPacketData.3.1783
type RMONMIB_Capturebuffertable_Capturebufferentry struct {
    parent types.Entity
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

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetFilter() yfilter.YFilter { return capturebufferentry.YFilter }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) SetFilter(yf yfilter.YFilter) { capturebufferentry.YFilter = yf }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetGoName(yname string) string {
    if yname == "captureBufferControlIndex" { return "Capturebuffercontrolindex" }
    if yname == "captureBufferIndex" { return "Capturebufferindex" }
    if yname == "captureBufferPacketID" { return "Capturebufferpacketid" }
    if yname == "captureBufferPacketData" { return "Capturebufferpacketdata" }
    if yname == "captureBufferPacketLength" { return "Capturebufferpacketlength" }
    if yname == "captureBufferPacketTime" { return "Capturebufferpackettime" }
    if yname == "captureBufferPacketStatus" { return "Capturebufferpacketstatus" }
    return ""
}

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetSegmentPath() string {
    return "captureBufferEntry" + "[captureBufferControlIndex='" + fmt.Sprintf("%v", capturebufferentry.Capturebuffercontrolindex) + "']" + "[captureBufferIndex='" + fmt.Sprintf("%v", capturebufferentry.Capturebufferindex) + "']"
}

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["captureBufferControlIndex"] = capturebufferentry.Capturebuffercontrolindex
    leafs["captureBufferIndex"] = capturebufferentry.Capturebufferindex
    leafs["captureBufferPacketID"] = capturebufferentry.Capturebufferpacketid
    leafs["captureBufferPacketData"] = capturebufferentry.Capturebufferpacketdata
    leafs["captureBufferPacketLength"] = capturebufferentry.Capturebufferpacketlength
    leafs["captureBufferPacketTime"] = capturebufferentry.Capturebufferpackettime
    leafs["captureBufferPacketStatus"] = capturebufferentry.Capturebufferpacketstatus
    return leafs
}

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetBundleName() string { return "cisco_ios_xe" }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetYangName() string { return "captureBufferEntry" }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) SetParent(parent types.Entity) { capturebufferentry.parent = parent }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetParent() types.Entity { return capturebufferentry.parent }

func (capturebufferentry *RMONMIB_Capturebuffertable_Capturebufferentry) GetParentYangName() string { return "captureBufferTable" }

// RMONMIB_Eventtable
// A list of events to be generated.
type RMONMIB_Eventtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of parameters that describe an event to be generated when certain
    // conditions are met.  As an example, an instance of the eventLastTimeSent
    // object might be named eventLastTimeSent.6. The type is slice of
    // RMONMIB_Eventtable_Evententry.
    Evententry []RMONMIB_Eventtable_Evententry
}

func (eventtable *RMONMIB_Eventtable) GetFilter() yfilter.YFilter { return eventtable.YFilter }

func (eventtable *RMONMIB_Eventtable) SetFilter(yf yfilter.YFilter) { eventtable.YFilter = yf }

func (eventtable *RMONMIB_Eventtable) GetGoName(yname string) string {
    if yname == "eventEntry" { return "Evententry" }
    return ""
}

func (eventtable *RMONMIB_Eventtable) GetSegmentPath() string {
    return "eventTable"
}

func (eventtable *RMONMIB_Eventtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "eventEntry" {
        for _, c := range eventtable.Evententry {
            if eventtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Eventtable_Evententry{}
        eventtable.Evententry = append(eventtable.Evententry, child)
        return &eventtable.Evententry[len(eventtable.Evententry)-1]
    }
    return nil
}

func (eventtable *RMONMIB_Eventtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eventtable.Evententry {
        children[eventtable.Evententry[i].GetSegmentPath()] = &eventtable.Evententry[i]
    }
    return children
}

func (eventtable *RMONMIB_Eventtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eventtable *RMONMIB_Eventtable) GetBundleName() string { return "cisco_ios_xe" }

func (eventtable *RMONMIB_Eventtable) GetYangName() string { return "eventTable" }

func (eventtable *RMONMIB_Eventtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (eventtable *RMONMIB_Eventtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (eventtable *RMONMIB_Eventtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (eventtable *RMONMIB_Eventtable) SetParent(parent types.Entity) { eventtable.parent = parent }

func (eventtable *RMONMIB_Eventtable) GetParent() types.Entity { return eventtable.parent }

func (eventtable *RMONMIB_Eventtable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Eventtable_Evententry
// A set of parameters that describe an event to be generated
// when certain conditions are met.  As an example, an instance
// of the eventLastTimeSent object might be named
// eventLastTimeSent.6
type RMONMIB_Eventtable_Evententry struct {
    parent types.Entity
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

func (evententry *RMONMIB_Eventtable_Evententry) GetFilter() yfilter.YFilter { return evententry.YFilter }

func (evententry *RMONMIB_Eventtable_Evententry) SetFilter(yf yfilter.YFilter) { evententry.YFilter = yf }

func (evententry *RMONMIB_Eventtable_Evententry) GetGoName(yname string) string {
    if yname == "eventIndex" { return "Eventindex" }
    if yname == "eventDescription" { return "Eventdescription" }
    if yname == "eventType" { return "Eventtype" }
    if yname == "eventCommunity" { return "Eventcommunity" }
    if yname == "eventLastTimeSent" { return "Eventlasttimesent" }
    if yname == "eventOwner" { return "Eventowner" }
    if yname == "eventStatus" { return "Eventstatus" }
    return ""
}

func (evententry *RMONMIB_Eventtable_Evententry) GetSegmentPath() string {
    return "eventEntry" + "[eventIndex='" + fmt.Sprintf("%v", evententry.Eventindex) + "']"
}

func (evententry *RMONMIB_Eventtable_Evententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (evententry *RMONMIB_Eventtable_Evententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (evententry *RMONMIB_Eventtable_Evententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["eventIndex"] = evententry.Eventindex
    leafs["eventDescription"] = evententry.Eventdescription
    leafs["eventType"] = evententry.Eventtype
    leafs["eventCommunity"] = evententry.Eventcommunity
    leafs["eventLastTimeSent"] = evententry.Eventlasttimesent
    leafs["eventOwner"] = evententry.Eventowner
    leafs["eventStatus"] = evententry.Eventstatus
    return leafs
}

func (evententry *RMONMIB_Eventtable_Evententry) GetBundleName() string { return "cisco_ios_xe" }

func (evententry *RMONMIB_Eventtable_Evententry) GetYangName() string { return "eventEntry" }

func (evententry *RMONMIB_Eventtable_Evententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (evententry *RMONMIB_Eventtable_Evententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (evententry *RMONMIB_Eventtable_Evententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (evententry *RMONMIB_Eventtable_Evententry) SetParent(parent types.Entity) { evententry.parent = parent }

func (evententry *RMONMIB_Eventtable_Evententry) GetParent() types.Entity { return evententry.parent }

func (evententry *RMONMIB_Eventtable_Evententry) GetParentYangName() string { return "eventTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of data describing an event that has been logged.  For example, an
    // instance of the logDescription object might be named logDescription.6.47.
    // The type is slice of RMONMIB_Logtable_Logentry.
    Logentry []RMONMIB_Logtable_Logentry
}

func (logtable *RMONMIB_Logtable) GetFilter() yfilter.YFilter { return logtable.YFilter }

func (logtable *RMONMIB_Logtable) SetFilter(yf yfilter.YFilter) { logtable.YFilter = yf }

func (logtable *RMONMIB_Logtable) GetGoName(yname string) string {
    if yname == "logEntry" { return "Logentry" }
    return ""
}

func (logtable *RMONMIB_Logtable) GetSegmentPath() string {
    return "logTable"
}

func (logtable *RMONMIB_Logtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logEntry" {
        for _, c := range logtable.Logentry {
            if logtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RMONMIB_Logtable_Logentry{}
        logtable.Logentry = append(logtable.Logentry, child)
        return &logtable.Logentry[len(logtable.Logentry)-1]
    }
    return nil
}

func (logtable *RMONMIB_Logtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logtable.Logentry {
        children[logtable.Logentry[i].GetSegmentPath()] = &logtable.Logentry[i]
    }
    return children
}

func (logtable *RMONMIB_Logtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logtable *RMONMIB_Logtable) GetBundleName() string { return "cisco_ios_xe" }

func (logtable *RMONMIB_Logtable) GetYangName() string { return "logTable" }

func (logtable *RMONMIB_Logtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (logtable *RMONMIB_Logtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (logtable *RMONMIB_Logtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (logtable *RMONMIB_Logtable) SetParent(parent types.Entity) { logtable.parent = parent }

func (logtable *RMONMIB_Logtable) GetParent() types.Entity { return logtable.parent }

func (logtable *RMONMIB_Logtable) GetParentYangName() string { return "RMON-MIB" }

// RMONMIB_Logtable_Logentry
// A set of data describing an event that has been
// logged.  For example, an instance of the logDescription
// object might be named logDescription.6.47
type RMONMIB_Logtable_Logentry struct {
    parent types.Entity
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

func (logentry *RMONMIB_Logtable_Logentry) GetFilter() yfilter.YFilter { return logentry.YFilter }

func (logentry *RMONMIB_Logtable_Logentry) SetFilter(yf yfilter.YFilter) { logentry.YFilter = yf }

func (logentry *RMONMIB_Logtable_Logentry) GetGoName(yname string) string {
    if yname == "logEventIndex" { return "Logeventindex" }
    if yname == "logIndex" { return "Logindex" }
    if yname == "logTime" { return "Logtime" }
    if yname == "logDescription" { return "Logdescription" }
    return ""
}

func (logentry *RMONMIB_Logtable_Logentry) GetSegmentPath() string {
    return "logEntry" + "[logEventIndex='" + fmt.Sprintf("%v", logentry.Logeventindex) + "']" + "[logIndex='" + fmt.Sprintf("%v", logentry.Logindex) + "']"
}

func (logentry *RMONMIB_Logtable_Logentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logentry *RMONMIB_Logtable_Logentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logentry *RMONMIB_Logtable_Logentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logEventIndex"] = logentry.Logeventindex
    leafs["logIndex"] = logentry.Logindex
    leafs["logTime"] = logentry.Logtime
    leafs["logDescription"] = logentry.Logdescription
    return leafs
}

func (logentry *RMONMIB_Logtable_Logentry) GetBundleName() string { return "cisco_ios_xe" }

func (logentry *RMONMIB_Logtable_Logentry) GetYangName() string { return "logEntry" }

func (logentry *RMONMIB_Logtable_Logentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (logentry *RMONMIB_Logtable_Logentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (logentry *RMONMIB_Logtable_Logentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (logentry *RMONMIB_Logtable_Logentry) SetParent(parent types.Entity) { logentry.parent = parent }

func (logentry *RMONMIB_Logtable_Logentry) GetParent() types.Entity { return logentry.parent }

func (logentry *RMONMIB_Logtable_Logentry) GetParentYangName() string { return "logTable" }

