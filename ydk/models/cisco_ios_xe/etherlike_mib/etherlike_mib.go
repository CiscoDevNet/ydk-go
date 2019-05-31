// The MIB module to describe generic objects for
// ethernet-like network interfaces.
// 
// The following reference is used throughout this
// MIB module:
// 
// [IEEE 802.3 Std] refers to:
//    IEEE Std 802.3, 2002 Edition: 'IEEE Standard
//    for Information technology -
//    Telecommunications and information exchange
//    between systems - Local and metropolitan
//    area networks - Specific requirements -
//    Part 3: Carrier sense multiple access with
//    collision detection (CSMA/CD) access method
//    and physical layer specifications', as
//    amended by IEEE Std 802.3ae-2002:
//    'Amendment: Media Access Control (MAC)
//    Parameters, Physical Layer, and Management
//    Parameters for 10 Gb/s Operation', August,
//    2002.
// 
// Of particular interest is Clause 30, '10 Mb/s,
// 100 Mb/s, 1000 Mb/s, and 10 Gb/s Management'.
// 
// Copyright (C) The Internet Society (2003).  This
// version of this MIB module is part of RFC 3635;
// see the RFC itself for full legal notices.
package etherlike_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package etherlike_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:EtherLike-MIB EtherLike-MIB}", reflect.TypeOf(EtherLikeMIB{}))
    ydk.RegisterEntity("EtherLike-MIB:EtherLike-MIB", reflect.TypeOf(EtherLikeMIB{}))
}

type Dot3TestTdr struct {
}

func (id Dot3TestTdr) String() string {
	return "EtherLike-MIB:dot3TestTdr"
}

type Dot3TestLoopBack struct {
}

func (id Dot3TestLoopBack) String() string {
	return "EtherLike-MIB:dot3TestLoopBack"
}

type Dot3ErrorInitError struct {
}

func (id Dot3ErrorInitError) String() string {
	return "EtherLike-MIB:dot3ErrorInitError"
}

type Dot3ErrorLoopbackError struct {
}

func (id Dot3ErrorLoopbackError) String() string {
	return "EtherLike-MIB:dot3ErrorLoopbackError"
}

// EtherLikeMIB
type EtherLikeMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a collection of ethernet-like interfaces attached to a
    // particular system. There will be one row in this table for each
    // ethernet-like interface in the system.
    Dot3StatsTable EtherLikeMIB_Dot3StatsTable

    // A collection of collision histograms for a particular set of interfaces.
    Dot3CollTable EtherLikeMIB_Dot3CollTable

    // A table of descriptive and status information about the MAC Control
    // sublayer on the ethernet-like interfaces attached to a particular system. 
    // There will be one row in this table for each ethernet-like interface in the
    // system which implements the MAC Control sublayer.  If some, but not all, of
    // the ethernet-like interfaces in the system implement the MAC Control
    // sublayer, there will be fewer rows in this table than in the
    // dot3StatsTable.
    Dot3ControlTable EtherLikeMIB_Dot3ControlTable

    // A table of descriptive and status information about the MAC Control PAUSE
    // function on the ethernet-like interfaces attached to a particular system.
    // There will be one row in this table for each ethernet-like interface in the
    // system which supports the MAC Control PAUSE function (i.e., the 'pause' bit
    // in the corresponding instance of dot3ControlFunctionsSupported is set).  If
    // some, but not all, of the ethernet-like interfaces in the system implement
    // the MAC Control PAUSE function (for example, if some interfaces only
    // support half-duplex), there will be fewer rows in this table than in the
    // dot3StatsTable.
    Dot3PauseTable EtherLikeMIB_Dot3PauseTable

    // A table containing 64-bit versions of error counters from the
    // dot3StatsTable.  The 32-bit versions of these counters may roll over quite
    // quickly on higher speed ethernet interfaces. The counters that have 64-bit
    // versions in this table are the counters that apply to full-duplex
    // interfaces, since 10 Gb/s and faster ethernet-like interfaces do not
    // support half-duplex, and very few 1000 Mb/s ethernet-like interfaces
    // support half-duplex.  Entries in this table are recommended for interfaces
    // capable of operating at 1000 Mb/s or faster, and are required for
    // interfaces capable of operating at 10 Gb/s or faster.  Lower speed
    // ethernet-like interfaces do not need entries in this table, in which case
    // there may be fewer entries in this table than in the dot3StatsTable.
    // However, implementations containing interfaces with a mix of speeds may
    // choose to implement entries in this table for all ethernet-like interfaces.
    Dot3HCStatsTable EtherLikeMIB_Dot3HCStatsTable
}

func (etherLikeMIB *EtherLikeMIB) GetEntityData() *types.CommonEntityData {
    etherLikeMIB.EntityData.YFilter = etherLikeMIB.YFilter
    etherLikeMIB.EntityData.YangName = "EtherLike-MIB"
    etherLikeMIB.EntityData.BundleName = "cisco_ios_xe"
    etherLikeMIB.EntityData.ParentYangName = "EtherLike-MIB"
    etherLikeMIB.EntityData.SegmentPath = "EtherLike-MIB:EtherLike-MIB"
    etherLikeMIB.EntityData.AbsolutePath = etherLikeMIB.EntityData.SegmentPath
    etherLikeMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherLikeMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherLikeMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherLikeMIB.EntityData.Children = types.NewOrderedMap()
    etherLikeMIB.EntityData.Children.Append("dot3StatsTable", types.YChild{"Dot3StatsTable", &etherLikeMIB.Dot3StatsTable})
    etherLikeMIB.EntityData.Children.Append("dot3CollTable", types.YChild{"Dot3CollTable", &etherLikeMIB.Dot3CollTable})
    etherLikeMIB.EntityData.Children.Append("dot3ControlTable", types.YChild{"Dot3ControlTable", &etherLikeMIB.Dot3ControlTable})
    etherLikeMIB.EntityData.Children.Append("dot3PauseTable", types.YChild{"Dot3PauseTable", &etherLikeMIB.Dot3PauseTable})
    etherLikeMIB.EntityData.Children.Append("dot3HCStatsTable", types.YChild{"Dot3HCStatsTable", &etherLikeMIB.Dot3HCStatsTable})
    etherLikeMIB.EntityData.Leafs = types.NewOrderedMap()

    etherLikeMIB.EntityData.YListKeys = []string {}

    return &(etherLikeMIB.EntityData)
}

// EtherLikeMIB_Dot3StatsTable
// Statistics for a collection of ethernet-like
// interfaces attached to a particular system.
// There will be one row in this table for each
// ethernet-like interface in the system.
type EtherLikeMIB_Dot3StatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a particular interface to an ethernet-like medium. The type
    // is slice of EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry.
    Dot3StatsEntry []*EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry
}

func (dot3StatsTable *EtherLikeMIB_Dot3StatsTable) GetEntityData() *types.CommonEntityData {
    dot3StatsTable.EntityData.YFilter = dot3StatsTable.YFilter
    dot3StatsTable.EntityData.YangName = "dot3StatsTable"
    dot3StatsTable.EntityData.BundleName = "cisco_ios_xe"
    dot3StatsTable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3StatsTable.EntityData.SegmentPath = "dot3StatsTable"
    dot3StatsTable.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/" + dot3StatsTable.EntityData.SegmentPath
    dot3StatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3StatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3StatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3StatsTable.EntityData.Children = types.NewOrderedMap()
    dot3StatsTable.EntityData.Children.Append("dot3StatsEntry", types.YChild{"Dot3StatsEntry", nil})
    for i := range dot3StatsTable.Dot3StatsEntry {
        dot3StatsTable.EntityData.Children.Append(types.GetSegmentPath(dot3StatsTable.Dot3StatsEntry[i]), types.YChild{"Dot3StatsEntry", dot3StatsTable.Dot3StatsEntry[i]})
    }
    dot3StatsTable.EntityData.Leafs = types.NewOrderedMap()

    dot3StatsTable.EntityData.YListKeys = []string {}

    return &(dot3StatsTable.EntityData)
}

// EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry
// Statistics for a particular interface to an
// ethernet-like medium.
type EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index value that uniquely identifies an
    // interface to an ethernet-like medium.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of ifIndex. The type is interface{} with range: 1..2147483647.
    Dot3StatsIndex interface{}

    // A count of frames received on a particular interface that are not an
    // integral number of octets in length and do not pass the FCS check.  The
    // count represented by an instance of this object is incremented when the
    // alignmentError status is returned by the MAC service to the LLC (or other
    // MAC user). Received frames for which multiple error conditions pertain are,
    // according to the conventions of IEEE 802.3 Layer Management, counted
    // exclusively according to the error status presented to the LLC.  This
    // counter does not increment for group encoding schemes greater than 4 bits
    // per group.  For interfaces operating at 10 Gb/s, this counter can roll over
    // in less than 5 minutes if it is incrementing at its maximum rate.  Since
    // that amount of time could be less than a management station's poll cycle
    // time, in order to avoid a loss of information, a management station is
    // advised to poll the dot3HCStatsAlignmentErrors object for 10 Gb/s or faster
    // interfaces.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Dot3StatsAlignmentErrors interface{}

    // A count of frames received on a particular interface that are an integral
    // number of octets in length but do not pass the FCS check.  This count does
    // not include frames received with frame-too-long or frame-too-short error. 
    // The count represented by an instance of this object is incremented when the
    // frameCheckError status is returned by the MAC service to the LLC (or other
    // MAC user). Received frames for which multiple error conditions pertain are,
    // according to the conventions of IEEE 802.3 Layer Management, counted
    // exclusively according to the error status presented to the LLC.  Note: 
    // Coding errors detected by the physical layer for speeds above 10 Mb/s will
    // cause the frame to fail the FCS check.  For interfaces operating at 10
    // Gb/s, this counter can roll over in less than 5 minutes if it is
    // incrementing at its maximum rate.  Since that amount of time could be less
    // than a management station's poll cycle time, in order to avoid a loss of
    // information, a management station is advised to poll the
    // dot3HCStatsFCSErrors object for 10 Gb/s or faster interfaces. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3StatsFCSErrors interface{}

    // A count of frames that are involved in a single collision, and are
    // subsequently transmitted successfully.  A frame that is counted by an
    // instance of this object is also counted by the corresponding instance of
    // either the ifOutUcastPkts, ifOutMulticastPkts, or ifOutBroadcastPkts, and
    // is not counted by the corresponding instance of the
    // dot3StatsMultipleCollisionFrames object.  This counter does not increment
    // when the interface is operating in full-duplex mode.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3StatsSingleCollisionFrames interface{}

    // A count of frames that are involved in more than one collision and are
    // subsequently transmitted successfully.  A frame that is counted by an
    // instance of this object is also counted by the corresponding instance of
    // either the ifOutUcastPkts, ifOutMulticastPkts, or ifOutBroadcastPkts, and
    // is not counted by the corresponding instance of the
    // dot3StatsSingleCollisionFrames object.  This counter does not increment
    // when the interface is operating in full-duplex mode.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3StatsMultipleCollisionFrames interface{}

    // A count of times that the SQE TEST ERROR is received on a particular
    // interface. The SQE TEST ERROR is set in accordance with the rules for
    // verification of the SQE detection mechanism in the PLS Carrier Sense
    // Function as described in IEEE Std. 802.3, 2000 Edition, section 7.2.4.6. 
    // This counter does not increment on interfaces operating at speeds greater
    // than 10 Mb/s, or on interfaces operating in full-duplex mode. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3StatsSQETestErrors interface{}

    // A count of frames for which the first transmission attempt on a particular
    // interface is delayed because the medium is busy.  The count represented by
    // an instance of this object does not include frames involved in collisions. 
    // This counter does not increment when the interface is operating in
    // full-duplex mode.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Dot3StatsDeferredTransmissions interface{}

    // The number of times that a collision is detected on a particular interface
    // later than one slotTime into the transmission of a packet.  A (late)
    // collision included in a count represented by an instance of this object is
    // also considered as a (generic) collision for purposes of other
    // collision-related statistics.  This counter does not increment when the
    // interface is operating in full-duplex mode.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Dot3StatsLateCollisions interface{}

    // A count of frames for which transmission on a particular interface fails
    // due to excessive collisions.  This counter does not increment when the
    // interface is operating in full-duplex mode.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Dot3StatsExcessiveCollisions interface{}

    // A count of frames for which transmission on a particular interface fails
    // due to an internal MAC sublayer transmit error. A frame is only counted by
    // an instance of this object if it is not counted by the corresponding
    // instance of either the dot3StatsLateCollisions object, the
    // dot3StatsExcessiveCollisions object, or the dot3StatsCarrierSenseErrors
    // object.  The precise meaning of the count represented by an instance of
    // this object is implementation- specific.  In particular, an instance of
    // this object may represent a count of transmission errors on a particular
    // interface that are not otherwise counted.  For interfaces operating at 10
    // Gb/s, this counter can roll over in less than 5 minutes if it is
    // incrementing at its maximum rate.  Since that amount of time could be less
    // than a management station's poll cycle time, in order to avoid a loss of
    // information, a management station is advised to poll the
    // dot3HCStatsInternalMacTransmitErrors object for 10 Gb/s or faster
    // interfaces.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Dot3StatsInternalMacTransmitErrors interface{}

    // The number of times that the carrier sense condition was lost or never
    // asserted when attempting to transmit a frame on a particular interface. 
    // The count represented by an instance of this object is incremented at most
    // once per transmission attempt, even if the carrier sense condition
    // fluctuates during a transmission attempt.  This counter does not increment
    // when the interface is operating in full-duplex mode.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3StatsCarrierSenseErrors interface{}

    // A count of frames received on a particular interface that exceed the
    // maximum permitted frame size.  The count represented by an instance of this
    // object is incremented when the frameTooLong status is returned by the MAC
    // service to the LLC (or other MAC user). Received frames for which multiple
    // error conditions pertain are, according to the conventions of IEEE 802.3
    // Layer Management, counted exclusively according to the error status
    // presented to the LLC.  For interfaces operating at 10 Gb/s, this counter
    // can roll over in less than 80 minutes if it is incrementing at its maximum
    // rate.  Since that amount of time could be less than a management station's
    // poll cycle time, in order to avoid a loss of information, a management
    // station is advised to poll the dot3HCStatsFrameTooLongs object for 10 Gb/s
    // or faster interfaces.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Dot3StatsFrameTooLongs interface{}

    // A count of frames for which reception on a particular interface fails due
    // to an internal MAC sublayer receive error. A frame is only counted by an
    // instance of this object if it is not counted by the corresponding instance
    // of either the dot3StatsFrameTooLongs object, the dot3StatsAlignmentErrors
    // object, or the dot3StatsFCSErrors object.  The precise meaning of the count
    // represented by an instance of this object is implementation- specific.  In
    // particular, an instance of this object may represent a count of receive
    // errors on a particular interface that are not otherwise counted.  For
    // interfaces operating at 10 Gb/s, this counter can roll over in less than 5
    // minutes if it is incrementing at its maximum rate.  Since that amount of
    // time could be less than a management station's poll cycle time, in order to
    // avoid a loss of information, a management station is advised to poll the
    // dot3HCStatsInternalMacReceiveErrors object for 10 Gb/s or faster
    // interfaces.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Dot3StatsInternalMacReceiveErrors interface{}

    // ******** THIS OBJECT IS DEPRECATED ********  This object contains an OBJECT
    // IDENTIFIER which identifies the chipset used to realize the interface.
    // Ethernet-like interfaces are typically built out of several different
    // chips. The MIB implementor is presented with a decision of which chip to
    // identify via this object. The implementor should identify the chip which is
    // usually called the Medium Access Control chip. If no such chip is easily
    // identifiable, the implementor should identify the chip which actually
    // gathers the transmit and receive statistics and error indications. This
    // would allow a manager station to correlate the statistics and the chip
    // generating them, giving it the ability to take into account any known
    // anomalies in the chip.  This object has been deprecated.  Implementation
    // feedback indicates that it is of limited use for debugging network problems
    // in the field, and the administrative overhead involved in maintaining a
    // registry of chipset OIDs is not justified. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Dot3StatsEtherChipSet interface{}

    // For an interface operating at 100 Mb/s, the number of times there was an
    // invalid data symbol when a valid carrier was present.  For an interface
    // operating in half-duplex mode at 1000 Mb/s, the number of times the
    // receiving media is non-idle (a carrier event) for a period of time equal to
    // or greater than slotTime, and during which there was at least one
    // occurrence of an event that causes the PHY to indicate 'Data reception
    // error' or 'carrier extend error' on the GMII.  For an interface operating
    // in full-duplex mode at 1000 Mb/s, the number of times the receiving media
    // is non-idle (a carrier event) for a period of time equal to or greater than
    // minFrameSize, and during which there was at least one occurrence of an
    // event that causes the PHY to indicate 'Data reception error' on the GMII. 
    // For an interface operating at 10 Gb/s, the number of times the receiving
    // media is non-idle (a carrier event) for a period of time equal to or
    // greater than minFrameSize, and during which there was at least one
    // occurrence of an event that causes the PHY to indicate 'Receive Error' on
    // the XGMII.  The count represented by an instance of this object is
    // incremented at most once per carrier event, even if multiple symbol errors
    // occur during the carrier event.  This count does not increment if a
    // collision is present.  This counter does not increment when the interface
    // is operating at 10 Mb/s.  For interfaces operating at 10 Gb/s, this counter
    // can roll over in less than 5 minutes if it is incrementing at its maximum
    // rate.  Since that amount of time could be less than a management station's
    // poll cycle time, in order to avoid a loss of information, a management
    // station is advised to poll the dot3HCStatsSymbolErrors object for 10 Gb/s
    // or faster interfaces.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Dot3StatsSymbolErrors interface{}

    // The current mode of operation of the MAC entity.  'unknown' indicates that
    // the current duplex mode could not be determined.  Management control of the
    // duplex mode is accomplished through the MAU MIB.  When an interface does
    // not support autonegotiation, or when autonegotiation is not enabled, the
    // duplex mode is controlled using ifMauDefaultType.  When autonegotiation is
    // supported and enabled, duplex mode is controlled using
    // ifMauAutoNegAdvertisedBits.  In either case, the currently operating duplex
    // mode is reflected both in this object and in ifMauType.  Note that this
    // object provides redundant information with ifMauType.  Normally, redundant
    // objects are discouraged.  However, in this instance, it allows a management
    // application to determine the duplex status of an interface without having
    // to know every possible value of ifMauType.  This was felt to be
    // sufficiently valuable to justify the redundancy. The type is
    // Dot3StatsDuplexStatus.
    Dot3StatsDuplexStatus interface{}

    // 'true' for interfaces operating at speeds above 1000 Mb/s that support Rate
    // Control through lowering the average data rate of the MAC sublayer, with
    // frame granularity, and 'false' otherwise. The type is bool.
    Dot3StatsRateControlAbility interface{}

    // The current Rate Control mode of operation of the MAC sublayer of this
    // interface. The type is Dot3StatsRateControlStatus.
    Dot3StatsRateControlStatus interface{}
}

func (dot3StatsEntry *EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry) GetEntityData() *types.CommonEntityData {
    dot3StatsEntry.EntityData.YFilter = dot3StatsEntry.YFilter
    dot3StatsEntry.EntityData.YangName = "dot3StatsEntry"
    dot3StatsEntry.EntityData.BundleName = "cisco_ios_xe"
    dot3StatsEntry.EntityData.ParentYangName = "dot3StatsTable"
    dot3StatsEntry.EntityData.SegmentPath = "dot3StatsEntry" + types.AddKeyToken(dot3StatsEntry.Dot3StatsIndex, "dot3StatsIndex")
    dot3StatsEntry.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/dot3StatsTable/" + dot3StatsEntry.EntityData.SegmentPath
    dot3StatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3StatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3StatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3StatsEntry.EntityData.Children = types.NewOrderedMap()
    dot3StatsEntry.EntityData.Leafs = types.NewOrderedMap()
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsIndex", types.YLeaf{"Dot3StatsIndex", dot3StatsEntry.Dot3StatsIndex})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsAlignmentErrors", types.YLeaf{"Dot3StatsAlignmentErrors", dot3StatsEntry.Dot3StatsAlignmentErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsFCSErrors", types.YLeaf{"Dot3StatsFCSErrors", dot3StatsEntry.Dot3StatsFCSErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsSingleCollisionFrames", types.YLeaf{"Dot3StatsSingleCollisionFrames", dot3StatsEntry.Dot3StatsSingleCollisionFrames})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsMultipleCollisionFrames", types.YLeaf{"Dot3StatsMultipleCollisionFrames", dot3StatsEntry.Dot3StatsMultipleCollisionFrames})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsSQETestErrors", types.YLeaf{"Dot3StatsSQETestErrors", dot3StatsEntry.Dot3StatsSQETestErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsDeferredTransmissions", types.YLeaf{"Dot3StatsDeferredTransmissions", dot3StatsEntry.Dot3StatsDeferredTransmissions})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsLateCollisions", types.YLeaf{"Dot3StatsLateCollisions", dot3StatsEntry.Dot3StatsLateCollisions})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsExcessiveCollisions", types.YLeaf{"Dot3StatsExcessiveCollisions", dot3StatsEntry.Dot3StatsExcessiveCollisions})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsInternalMacTransmitErrors", types.YLeaf{"Dot3StatsInternalMacTransmitErrors", dot3StatsEntry.Dot3StatsInternalMacTransmitErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsCarrierSenseErrors", types.YLeaf{"Dot3StatsCarrierSenseErrors", dot3StatsEntry.Dot3StatsCarrierSenseErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsFrameTooLongs", types.YLeaf{"Dot3StatsFrameTooLongs", dot3StatsEntry.Dot3StatsFrameTooLongs})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsInternalMacReceiveErrors", types.YLeaf{"Dot3StatsInternalMacReceiveErrors", dot3StatsEntry.Dot3StatsInternalMacReceiveErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsEtherChipSet", types.YLeaf{"Dot3StatsEtherChipSet", dot3StatsEntry.Dot3StatsEtherChipSet})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsSymbolErrors", types.YLeaf{"Dot3StatsSymbolErrors", dot3StatsEntry.Dot3StatsSymbolErrors})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsDuplexStatus", types.YLeaf{"Dot3StatsDuplexStatus", dot3StatsEntry.Dot3StatsDuplexStatus})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsRateControlAbility", types.YLeaf{"Dot3StatsRateControlAbility", dot3StatsEntry.Dot3StatsRateControlAbility})
    dot3StatsEntry.EntityData.Leafs.Append("dot3StatsRateControlStatus", types.YLeaf{"Dot3StatsRateControlStatus", dot3StatsEntry.Dot3StatsRateControlStatus})

    dot3StatsEntry.EntityData.YListKeys = []string {"Dot3StatsIndex"}

    return &(dot3StatsEntry.EntityData)
}

// EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus represents valuable to justify the redundancy.
type EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus string

const (
    EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus_unknown EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus = "unknown"

    EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus_halfDuplex EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus = "halfDuplex"

    EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus_fullDuplex EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsDuplexStatus = "fullDuplex"
)

// EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus represents the MAC sublayer of this interface.
type EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus string

const (
    EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus_rateControlOff EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus = "rateControlOff"

    EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus_rateControlOn EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus = "rateControlOn"

    EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus_unknown EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsRateControlStatus = "unknown"
)

// EtherLikeMIB_Dot3CollTable
// A collection of collision histograms for a
// particular set of interfaces.
type EtherLikeMIB_Dot3CollTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cell in the histogram of per-frame collisions for a particular interface.
    // An instance of this object represents the frequency of individual MAC
    // frames for which the transmission (successful or otherwise) on a particular
    // interface is accompanied by a particular number of media collisions. The
    // type is slice of EtherLikeMIB_Dot3CollTable_Dot3CollEntry.
    Dot3CollEntry []*EtherLikeMIB_Dot3CollTable_Dot3CollEntry
}

func (dot3CollTable *EtherLikeMIB_Dot3CollTable) GetEntityData() *types.CommonEntityData {
    dot3CollTable.EntityData.YFilter = dot3CollTable.YFilter
    dot3CollTable.EntityData.YangName = "dot3CollTable"
    dot3CollTable.EntityData.BundleName = "cisco_ios_xe"
    dot3CollTable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3CollTable.EntityData.SegmentPath = "dot3CollTable"
    dot3CollTable.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/" + dot3CollTable.EntityData.SegmentPath
    dot3CollTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3CollTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3CollTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3CollTable.EntityData.Children = types.NewOrderedMap()
    dot3CollTable.EntityData.Children.Append("dot3CollEntry", types.YChild{"Dot3CollEntry", nil})
    for i := range dot3CollTable.Dot3CollEntry {
        dot3CollTable.EntityData.Children.Append(types.GetSegmentPath(dot3CollTable.Dot3CollEntry[i]), types.YChild{"Dot3CollEntry", dot3CollTable.Dot3CollEntry[i]})
    }
    dot3CollTable.EntityData.Leafs = types.NewOrderedMap()

    dot3CollTable.EntityData.YListKeys = []string {}

    return &(dot3CollTable.EntityData)
}

// EtherLikeMIB_Dot3CollTable_Dot3CollEntry
// A cell in the histogram of per-frame
// collisions for a particular interface.  An
// instance of this object represents the
// frequency of individual MAC frames for which
// the transmission (successful or otherwise) on a
// particular interface is accompanied by a
// particular number of media collisions.
type EtherLikeMIB_Dot3CollTable_Dot3CollEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The number of per-frame media collisions for which
    // a particular collision histogram cell represents the frequency on a
    // particular interface. The type is interface{} with range: 1..16.
    Dot3CollCount interface{}

    // A count of individual MAC frames for which the transmission (successful or
    // otherwise) on a particular interface occurs after the frame has experienced
    // exactly the number of collisions in the associated dot3CollCount object. 
    // For example, a frame which is transmitted on interface 77 after
    // experiencing exactly 4 collisions would be indicated by incrementing only
    // dot3CollFrequencies.77.4. No other instance of dot3CollFrequencies would be
    // incremented in this example.  This counter does not increment when the
    // interface is operating in full-duplex mode.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Dot3CollFrequencies interface{}
}

func (dot3CollEntry *EtherLikeMIB_Dot3CollTable_Dot3CollEntry) GetEntityData() *types.CommonEntityData {
    dot3CollEntry.EntityData.YFilter = dot3CollEntry.YFilter
    dot3CollEntry.EntityData.YangName = "dot3CollEntry"
    dot3CollEntry.EntityData.BundleName = "cisco_ios_xe"
    dot3CollEntry.EntityData.ParentYangName = "dot3CollTable"
    dot3CollEntry.EntityData.SegmentPath = "dot3CollEntry" + types.AddKeyToken(dot3CollEntry.IfIndex, "ifIndex") + types.AddKeyToken(dot3CollEntry.Dot3CollCount, "dot3CollCount")
    dot3CollEntry.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/dot3CollTable/" + dot3CollEntry.EntityData.SegmentPath
    dot3CollEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3CollEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3CollEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3CollEntry.EntityData.Children = types.NewOrderedMap()
    dot3CollEntry.EntityData.Leafs = types.NewOrderedMap()
    dot3CollEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", dot3CollEntry.IfIndex})
    dot3CollEntry.EntityData.Leafs.Append("dot3CollCount", types.YLeaf{"Dot3CollCount", dot3CollEntry.Dot3CollCount})
    dot3CollEntry.EntityData.Leafs.Append("dot3CollFrequencies", types.YLeaf{"Dot3CollFrequencies", dot3CollEntry.Dot3CollFrequencies})

    dot3CollEntry.EntityData.YListKeys = []string {"IfIndex", "Dot3CollCount"}

    return &(dot3CollEntry.EntityData)
}

// EtherLikeMIB_Dot3ControlTable
// A table of descriptive and status information
// about the MAC Control sublayer on the
// ethernet-like interfaces attached to a
// particular system.  There will be one row in
// this table for each ethernet-like interface in
// the system which implements the MAC Control
// sublayer.  If some, but not all, of the
// ethernet-like interfaces in the system implement
// the MAC Control sublayer, there will be fewer
// rows in this table than in the dot3StatsTable.
type EtherLikeMIB_Dot3ControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about the MAC Control
    // sublayer on a single ethernet-like interface. The type is slice of
    // EtherLikeMIB_Dot3ControlTable_Dot3ControlEntry.
    Dot3ControlEntry []*EtherLikeMIB_Dot3ControlTable_Dot3ControlEntry
}

func (dot3ControlTable *EtherLikeMIB_Dot3ControlTable) GetEntityData() *types.CommonEntityData {
    dot3ControlTable.EntityData.YFilter = dot3ControlTable.YFilter
    dot3ControlTable.EntityData.YangName = "dot3ControlTable"
    dot3ControlTable.EntityData.BundleName = "cisco_ios_xe"
    dot3ControlTable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3ControlTable.EntityData.SegmentPath = "dot3ControlTable"
    dot3ControlTable.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/" + dot3ControlTable.EntityData.SegmentPath
    dot3ControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3ControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3ControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3ControlTable.EntityData.Children = types.NewOrderedMap()
    dot3ControlTable.EntityData.Children.Append("dot3ControlEntry", types.YChild{"Dot3ControlEntry", nil})
    for i := range dot3ControlTable.Dot3ControlEntry {
        dot3ControlTable.EntityData.Children.Append(types.GetSegmentPath(dot3ControlTable.Dot3ControlEntry[i]), types.YChild{"Dot3ControlEntry", dot3ControlTable.Dot3ControlEntry[i]})
    }
    dot3ControlTable.EntityData.Leafs = types.NewOrderedMap()

    dot3ControlTable.EntityData.YListKeys = []string {}

    return &(dot3ControlTable.EntityData)
}

// EtherLikeMIB_Dot3ControlTable_Dot3ControlEntry
// An entry in the table, containing information
// about the MAC Control sublayer on a single
// ethernet-like interface.
type EtherLikeMIB_Dot3ControlTable_Dot3ControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsIndex
    Dot3StatsIndex interface{}

    // A list of the possible MAC Control functions implemented for this
    // interface. The type is map[string]bool.
    Dot3ControlFunctionsSupported interface{}

    // A count of MAC Control frames received on this interface that contain an
    // opcode that is not supported by this device.  For interfaces operating at
    // 10 Gb/s, this counter can roll over in less than 5 minutes if it is
    // incrementing at its maximum rate.  Since that amount of time could be less
    // than a management station's poll cycle time, in order to avoid a loss of
    // information, a management station is advised to poll the
    // dot3HCControlInUnknownOpcodes object for 10 Gb/s or faster interfaces. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3ControlInUnknownOpcodes interface{}

    // A count of MAC Control frames received on this interface that contain an
    // opcode that is not supported by this device.  This counter is a 64 bit
    // version of dot3ControlInUnknownOpcodes.  It should be used on interfaces
    // operating at 10 Gb/s or faster.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    Dot3HCControlInUnknownOpcodes interface{}
}

func (dot3ControlEntry *EtherLikeMIB_Dot3ControlTable_Dot3ControlEntry) GetEntityData() *types.CommonEntityData {
    dot3ControlEntry.EntityData.YFilter = dot3ControlEntry.YFilter
    dot3ControlEntry.EntityData.YangName = "dot3ControlEntry"
    dot3ControlEntry.EntityData.BundleName = "cisco_ios_xe"
    dot3ControlEntry.EntityData.ParentYangName = "dot3ControlTable"
    dot3ControlEntry.EntityData.SegmentPath = "dot3ControlEntry" + types.AddKeyToken(dot3ControlEntry.Dot3StatsIndex, "dot3StatsIndex")
    dot3ControlEntry.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/dot3ControlTable/" + dot3ControlEntry.EntityData.SegmentPath
    dot3ControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3ControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3ControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3ControlEntry.EntityData.Children = types.NewOrderedMap()
    dot3ControlEntry.EntityData.Leafs = types.NewOrderedMap()
    dot3ControlEntry.EntityData.Leafs.Append("dot3StatsIndex", types.YLeaf{"Dot3StatsIndex", dot3ControlEntry.Dot3StatsIndex})
    dot3ControlEntry.EntityData.Leafs.Append("dot3ControlFunctionsSupported", types.YLeaf{"Dot3ControlFunctionsSupported", dot3ControlEntry.Dot3ControlFunctionsSupported})
    dot3ControlEntry.EntityData.Leafs.Append("dot3ControlInUnknownOpcodes", types.YLeaf{"Dot3ControlInUnknownOpcodes", dot3ControlEntry.Dot3ControlInUnknownOpcodes})
    dot3ControlEntry.EntityData.Leafs.Append("dot3HCControlInUnknownOpcodes", types.YLeaf{"Dot3HCControlInUnknownOpcodes", dot3ControlEntry.Dot3HCControlInUnknownOpcodes})

    dot3ControlEntry.EntityData.YListKeys = []string {"Dot3StatsIndex"}

    return &(dot3ControlEntry.EntityData)
}

// EtherLikeMIB_Dot3PauseTable
// A table of descriptive and status information
// about the MAC Control PAUSE function on the
// ethernet-like interfaces attached to a
// particular system. There will be one row in
// this table for each ethernet-like interface in
// the system which supports the MAC Control PAUSE
// function (i.e., the 'pause' bit in the
// corresponding instance of
// dot3ControlFunctionsSupported is set).  If some,
// but not all, of the ethernet-like interfaces in
// the system implement the MAC Control PAUSE
// function (for example, if some interfaces only
// support half-duplex), there will be fewer rows
// in this table than in the dot3StatsTable.
type EtherLikeMIB_Dot3PauseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about the MAC Control PAUSE
    // function on a single ethernet-like interface. The type is slice of
    // EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry.
    Dot3PauseEntry []*EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry
}

func (dot3PauseTable *EtherLikeMIB_Dot3PauseTable) GetEntityData() *types.CommonEntityData {
    dot3PauseTable.EntityData.YFilter = dot3PauseTable.YFilter
    dot3PauseTable.EntityData.YangName = "dot3PauseTable"
    dot3PauseTable.EntityData.BundleName = "cisco_ios_xe"
    dot3PauseTable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3PauseTable.EntityData.SegmentPath = "dot3PauseTable"
    dot3PauseTable.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/" + dot3PauseTable.EntityData.SegmentPath
    dot3PauseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3PauseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3PauseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3PauseTable.EntityData.Children = types.NewOrderedMap()
    dot3PauseTable.EntityData.Children.Append("dot3PauseEntry", types.YChild{"Dot3PauseEntry", nil})
    for i := range dot3PauseTable.Dot3PauseEntry {
        dot3PauseTable.EntityData.Children.Append(types.GetSegmentPath(dot3PauseTable.Dot3PauseEntry[i]), types.YChild{"Dot3PauseEntry", dot3PauseTable.Dot3PauseEntry[i]})
    }
    dot3PauseTable.EntityData.Leafs = types.NewOrderedMap()

    dot3PauseTable.EntityData.YListKeys = []string {}

    return &(dot3PauseTable.EntityData)
}

// EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry
// An entry in the table, containing information
// about the MAC Control PAUSE function on a single
// ethernet-like interface.
type EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsIndex
    Dot3StatsIndex interface{}

    // This object is used to configure the default administrative PAUSE mode for
    // this interface.  This object represents the administratively-configured
    // PAUSE mode for this interface.  If auto-negotiation is not enabled or is
    // not implemented for the active MAU attached to this interface, the value of
    // this object determines the operational PAUSE mode of the interface whenever
    // it is operating in full-duplex mode.  In this case, a set to this object
    // will force the interface into the specified mode.  If auto-negotiation is
    // implemented and enabled for the MAU attached to this interface, the PAUSE
    // mode for this interface is determined by auto-negotiation, and the value of
    // this object denotes the mode to which the interface will automatically
    // revert if/when auto-negotiation is later disabled.  Note that when
    // auto-negotiation is running, administrative control of the PAUSE mode may
    // be accomplished using the ifMauAutoNegCapAdvertisedBits object in the
    // MAU-MIB.  Note that the value of this object is ignored when the interface
    // is not operating in full-duplex mode.  An attempt to set this object to
    // 'enabledXmit(2)' or 'enabledRcv(3)' will fail on interfaces that do not
    // support operation at greater than 100 Mb/s. The type is Dot3PauseAdminMode.
    Dot3PauseAdminMode interface{}

    // This object reflects the PAUSE mode currently in use on this interface, as
    // determined by either (1) the result of the auto-negotiation function or (2)
    // if auto-negotiation is not enabled or is not implemented for the active MAU
    // attached to this interface, by the value of dot3PauseAdminMode.  Interfaces
    // operating at 100 Mb/s or less will never return 'enabledXmit(2)' or
    // 'enabledRcv(3)'.  Interfaces operating in half-duplex mode will always
    // return 'disabled(1)'.  Interfaces on which auto-negotiation is enabled but
    // not yet completed should return the value 'disabled(1)'. The type is
    // Dot3PauseOperMode.
    Dot3PauseOperMode interface{}

    // A count of MAC Control frames received on this interface with an opcode
    // indicating the PAUSE operation.  This counter does not increment when the
    // interface is operating in half-duplex mode.  For interfaces operating at 10
    // Gb/s, this counter can roll over in less than 5 minutes if it is
    // incrementing at its maximum rate.  Since that amount of time could be less
    // than a management station's poll cycle time, in order to avoid a loss of
    // information, a management station is advised to poll the
    // dot3HCInPauseFrames object for 10 Gb/s or faster interfaces. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3InPauseFrames interface{}

    // A count of MAC Control frames transmitted on this interface with an opcode
    // indicating the PAUSE operation.  This counter does not increment when the
    // interface is operating in half-duplex mode.  For interfaces operating at 10
    // Gb/s, this counter can roll over in less than 5 minutes if it is
    // incrementing at its maximum rate.  Since that amount of time could be less
    // than a management station's poll cycle time, in order to avoid a loss of
    // information, a management station is advised to poll the
    // dot3HCOutPauseFrames object for 10 Gb/s or faster interfaces. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Dot3OutPauseFrames interface{}

    // A count of MAC Control frames received on this interface with an opcode
    // indicating the PAUSE operation.  This counter does not increment when the
    // interface is operating in half-duplex mode.  This counter is a 64 bit
    // version of dot3InPauseFrames.  It should be used on interfaces operating at
    // 10 Gb/s or faster.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    Dot3HCInPauseFrames interface{}

    // A count of MAC Control frames transmitted on this interface with an opcode
    // indicating the PAUSE operation.  This counter does not increment when the
    // interface is operating in half-duplex mode.  This counter is a 64 bit
    // version of dot3OutPauseFrames.  It should be used on interfaces operating
    // at 10 Gb/s or faster.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    Dot3HCOutPauseFrames interface{}
}

func (dot3PauseEntry *EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry) GetEntityData() *types.CommonEntityData {
    dot3PauseEntry.EntityData.YFilter = dot3PauseEntry.YFilter
    dot3PauseEntry.EntityData.YangName = "dot3PauseEntry"
    dot3PauseEntry.EntityData.BundleName = "cisco_ios_xe"
    dot3PauseEntry.EntityData.ParentYangName = "dot3PauseTable"
    dot3PauseEntry.EntityData.SegmentPath = "dot3PauseEntry" + types.AddKeyToken(dot3PauseEntry.Dot3StatsIndex, "dot3StatsIndex")
    dot3PauseEntry.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/dot3PauseTable/" + dot3PauseEntry.EntityData.SegmentPath
    dot3PauseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3PauseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3PauseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3PauseEntry.EntityData.Children = types.NewOrderedMap()
    dot3PauseEntry.EntityData.Leafs = types.NewOrderedMap()
    dot3PauseEntry.EntityData.Leafs.Append("dot3StatsIndex", types.YLeaf{"Dot3StatsIndex", dot3PauseEntry.Dot3StatsIndex})
    dot3PauseEntry.EntityData.Leafs.Append("dot3PauseAdminMode", types.YLeaf{"Dot3PauseAdminMode", dot3PauseEntry.Dot3PauseAdminMode})
    dot3PauseEntry.EntityData.Leafs.Append("dot3PauseOperMode", types.YLeaf{"Dot3PauseOperMode", dot3PauseEntry.Dot3PauseOperMode})
    dot3PauseEntry.EntityData.Leafs.Append("dot3InPauseFrames", types.YLeaf{"Dot3InPauseFrames", dot3PauseEntry.Dot3InPauseFrames})
    dot3PauseEntry.EntityData.Leafs.Append("dot3OutPauseFrames", types.YLeaf{"Dot3OutPauseFrames", dot3PauseEntry.Dot3OutPauseFrames})
    dot3PauseEntry.EntityData.Leafs.Append("dot3HCInPauseFrames", types.YLeaf{"Dot3HCInPauseFrames", dot3PauseEntry.Dot3HCInPauseFrames})
    dot3PauseEntry.EntityData.Leafs.Append("dot3HCOutPauseFrames", types.YLeaf{"Dot3HCOutPauseFrames", dot3PauseEntry.Dot3HCOutPauseFrames})

    dot3PauseEntry.EntityData.YListKeys = []string {"Dot3StatsIndex"}

    return &(dot3PauseEntry.EntityData)
}

// EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode represents at greater than 100 Mb/s.
type EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode string

const (
    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode_disabled EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode = "disabled"

    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode_enabledXmit EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode = "enabledXmit"

    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode_enabledRcv EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode = "enabledRcv"

    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode_enabledXmitAndRcv EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseAdminMode = "enabledXmitAndRcv"
)

// EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode represents 'disabled(1)'.
type EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode string

const (
    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode_disabled EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode = "disabled"

    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode_enabledXmit EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode = "enabledXmit"

    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode_enabledRcv EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode = "enabledRcv"

    EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode_enabledXmitAndRcv EtherLikeMIB_Dot3PauseTable_Dot3PauseEntry_Dot3PauseOperMode = "enabledXmitAndRcv"
)

// EtherLikeMIB_Dot3HCStatsTable
// A table containing 64-bit versions of error
// counters from the dot3StatsTable.  The 32-bit
// versions of these counters may roll over quite
// quickly on higher speed ethernet interfaces.
// The counters that have 64-bit versions in this
// table are the counters that apply to full-duplex
// interfaces, since 10 Gb/s and faster
// ethernet-like interfaces do not support
// half-duplex, and very few 1000 Mb/s
// ethernet-like interfaces support half-duplex.
// 
// Entries in this table are recommended for
// interfaces capable of operating at 1000 Mb/s or
// faster, and are required for interfaces capable
// of operating at 10 Gb/s or faster.  Lower speed
// ethernet-like interfaces do not need entries in
// this table, in which case there may be fewer
// entries in this table than in the
// dot3StatsTable. However, implementations
// containing interfaces with a mix of speeds may
// choose to implement entries in this table for
// all ethernet-like interfaces.
type EtherLikeMIB_Dot3HCStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing 64-bit statistics for a single ethernet-like interface.
    // The type is slice of EtherLikeMIB_Dot3HCStatsTable_Dot3HCStatsEntry.
    Dot3HCStatsEntry []*EtherLikeMIB_Dot3HCStatsTable_Dot3HCStatsEntry
}

func (dot3HCStatsTable *EtherLikeMIB_Dot3HCStatsTable) GetEntityData() *types.CommonEntityData {
    dot3HCStatsTable.EntityData.YFilter = dot3HCStatsTable.YFilter
    dot3HCStatsTable.EntityData.YangName = "dot3HCStatsTable"
    dot3HCStatsTable.EntityData.BundleName = "cisco_ios_xe"
    dot3HCStatsTable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3HCStatsTable.EntityData.SegmentPath = "dot3HCStatsTable"
    dot3HCStatsTable.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/" + dot3HCStatsTable.EntityData.SegmentPath
    dot3HCStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3HCStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3HCStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3HCStatsTable.EntityData.Children = types.NewOrderedMap()
    dot3HCStatsTable.EntityData.Children.Append("dot3HCStatsEntry", types.YChild{"Dot3HCStatsEntry", nil})
    for i := range dot3HCStatsTable.Dot3HCStatsEntry {
        dot3HCStatsTable.EntityData.Children.Append(types.GetSegmentPath(dot3HCStatsTable.Dot3HCStatsEntry[i]), types.YChild{"Dot3HCStatsEntry", dot3HCStatsTable.Dot3HCStatsEntry[i]})
    }
    dot3HCStatsTable.EntityData.Leafs = types.NewOrderedMap()

    dot3HCStatsTable.EntityData.YListKeys = []string {}

    return &(dot3HCStatsTable.EntityData)
}

// EtherLikeMIB_Dot3HCStatsTable_Dot3HCStatsEntry
// An entry containing 64-bit statistics for a
// single ethernet-like interface.
type EtherLikeMIB_Dot3HCStatsTable_Dot3HCStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsIndex
    Dot3StatsIndex interface{}

    // A count of frames received on a particular interface that are not an
    // integral number of octets in length and do not pass the FCS check.  The
    // count represented by an instance of this object is incremented when the
    // alignmentError status is returned by the MAC service to the LLC (or other
    // MAC user). Received frames for which multiple error conditions pertain are,
    // according to the conventions of IEEE 802.3 Layer Management, counted
    // exclusively according to the error status presented to the LLC.  This
    // counter does not increment for group encoding schemes greater than 4 bits
    // per group.  This counter is a 64 bit version of dot3StatsAlignmentErrors. 
    // It should be used on interfaces operating at 10 Gb/s or faster. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Dot3HCStatsAlignmentErrors interface{}

    // A count of frames received on a particular interface that are an integral
    // number of octets in length but do not pass the FCS check.  This count does
    // not include frames received with frame-too-long or frame-too-short error. 
    // The count represented by an instance of this object is incremented when the
    // frameCheckError status is returned by the MAC service to the LLC (or other
    // MAC user). Received frames for which multiple error conditions pertain are,
    // according to the conventions of IEEE 802.3 Layer Management, counted
    // exclusively according to the error status presented to the LLC.  Note: 
    // Coding errors detected by the physical layer for speeds above 10 Mb/s will
    // cause the frame to fail the FCS check.  This counter is a 64 bit version of
    // dot3StatsFCSErrors.  It should be used on interfaces operating at 10 Gb/s
    // or faster.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    Dot3HCStatsFCSErrors interface{}

    // A count of frames for which transmission on a particular interface fails
    // due to an internal MAC sublayer transmit error. A frame is only counted by
    // an instance of this object if it is not counted by the corresponding
    // instance of either the dot3StatsLateCollisions object, the
    // dot3StatsExcessiveCollisions object, or the dot3StatsCarrierSenseErrors
    // object.  The precise meaning of the count represented by an instance of
    // this object is implementation- specific.  In particular, an instance of
    // this object may represent a count of transmission errors on a particular
    // interface that are not otherwise counted.  This counter is a 64 bit version
    // of dot3StatsInternalMacTransmitErrors.  It should be used on interfaces
    // operating at 10 Gb/s or faster.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    Dot3HCStatsInternalMacTransmitErrors interface{}

    // A count of frames received on a particular interface that exceed the
    // maximum permitted frame size.  The count represented by an instance of this
    // object is incremented when the frameTooLong status is returned by the MAC
    // service to the LLC (or other MAC user). Received frames for which multiple
    // error conditions pertain are, according to the conventions of IEEE 802.3
    // Layer Management, counted exclusively according to the error status
    // presented to the LLC.  This counter is a 64 bit version of
    // dot3StatsFrameTooLongs.  It should be used on interfaces operating at 10
    // Gb/s or faster.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    Dot3HCStatsFrameTooLongs interface{}

    // A count of frames for which reception on a particular interface fails due
    // to an internal MAC sublayer receive error. A frame is only counted by an
    // instance of this object if it is not counted by the corresponding instance
    // of either the dot3StatsFrameTooLongs object, the dot3StatsAlignmentErrors
    // object, or the dot3StatsFCSErrors object.  The precise meaning of the count
    // represented by an instance of this object is implementation- specific.  In
    // particular, an instance of this object may represent a count of receive
    // errors on a particular interface that are not otherwise counted.  This
    // counter is a 64 bit version of dot3StatsInternalMacReceiveErrors.  It
    // should be used on interfaces operating at 10 Gb/s or faster. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Dot3HCStatsInternalMacReceiveErrors interface{}

    // For an interface operating at 100 Mb/s, the number of times there was an
    // invalid data symbol when a valid carrier was present.  For an interface
    // operating in half-duplex mode at 1000 Mb/s, the number of times the
    // receiving media is non-idle (a carrier event) for a period of time equal to
    // or greater than slotTime, and during which there was at least one
    // occurrence of an event that causes the PHY to indicate 'Data reception
    // error' or 'carrier extend error' on the GMII.  For an interface operating
    // in full-duplex mode at 1000 Mb/s, the number of times the receiving media
    // is non-idle (a carrier event) for a period of time equal to or greater than
    // minFrameSize, and during which there was at least one occurrence of an
    // event that causes the PHY to indicate 'Data reception error' on the GMII. 
    // For an interface operating at 10 Gb/s, the number of times the receiving
    // media is non-idle (a carrier event) for a period of time equal to or
    // greater than minFrameSize, and during which there was at least one
    // occurrence of an event that causes the PHY to indicate 'Receive Error' on
    // the XGMII.  The count represented by an instance of this object is
    // incremented at most once per carrier event, even if multiple symbol errors
    // occur during the carrier event.  This count does not increment if a
    // collision is present.  This counter is a 64 bit version of
    // dot3StatsSymbolErrors.  It should be used on interfaces operating at 10
    // Gb/s or faster.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    Dot3HCStatsSymbolErrors interface{}
}

func (dot3HCStatsEntry *EtherLikeMIB_Dot3HCStatsTable_Dot3HCStatsEntry) GetEntityData() *types.CommonEntityData {
    dot3HCStatsEntry.EntityData.YFilter = dot3HCStatsEntry.YFilter
    dot3HCStatsEntry.EntityData.YangName = "dot3HCStatsEntry"
    dot3HCStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    dot3HCStatsEntry.EntityData.ParentYangName = "dot3HCStatsTable"
    dot3HCStatsEntry.EntityData.SegmentPath = "dot3HCStatsEntry" + types.AddKeyToken(dot3HCStatsEntry.Dot3StatsIndex, "dot3StatsIndex")
    dot3HCStatsEntry.EntityData.AbsolutePath = "EtherLike-MIB:EtherLike-MIB/dot3HCStatsTable/" + dot3HCStatsEntry.EntityData.SegmentPath
    dot3HCStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3HCStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3HCStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3HCStatsEntry.EntityData.Children = types.NewOrderedMap()
    dot3HCStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3StatsIndex", types.YLeaf{"Dot3StatsIndex", dot3HCStatsEntry.Dot3StatsIndex})
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3HCStatsAlignmentErrors", types.YLeaf{"Dot3HCStatsAlignmentErrors", dot3HCStatsEntry.Dot3HCStatsAlignmentErrors})
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3HCStatsFCSErrors", types.YLeaf{"Dot3HCStatsFCSErrors", dot3HCStatsEntry.Dot3HCStatsFCSErrors})
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3HCStatsInternalMacTransmitErrors", types.YLeaf{"Dot3HCStatsInternalMacTransmitErrors", dot3HCStatsEntry.Dot3HCStatsInternalMacTransmitErrors})
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3HCStatsFrameTooLongs", types.YLeaf{"Dot3HCStatsFrameTooLongs", dot3HCStatsEntry.Dot3HCStatsFrameTooLongs})
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3HCStatsInternalMacReceiveErrors", types.YLeaf{"Dot3HCStatsInternalMacReceiveErrors", dot3HCStatsEntry.Dot3HCStatsInternalMacReceiveErrors})
    dot3HCStatsEntry.EntityData.Leafs.Append("dot3HCStatsSymbolErrors", types.YLeaf{"Dot3HCStatsSymbolErrors", dot3HCStatsEntry.Dot3HCStatsSymbolErrors})

    dot3HCStatsEntry.EntityData.YListKeys = []string {"Dot3StatsIndex"}

    return &(dot3HCStatsEntry.EntityData)
}

