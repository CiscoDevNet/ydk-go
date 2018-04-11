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

type Dot3Testtdr struct {
}

func (id Dot3Testtdr) String() string {
	return "EtherLike-MIB:dot3TestTdr"
}

type Dot3Testloopback struct {
}

func (id Dot3Testloopback) String() string {
	return "EtherLike-MIB:dot3TestLoopBack"
}

type Dot3Erroriniterror struct {
}

func (id Dot3Erroriniterror) String() string {
	return "EtherLike-MIB:dot3ErrorInitError"
}

type Dot3Errorloopbackerror struct {
}

func (id Dot3Errorloopbackerror) String() string {
	return "EtherLike-MIB:dot3ErrorLoopbackError"
}

// EtherLikeMIB
type EtherLikeMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a collection of ethernet-like interfaces attached to a
    // particular system. There will be one row in this table for each
    // ethernet-like interface in the system.
    Dot3Statstable EtherLikeMIB_Dot3Statstable

    // A collection of collision histograms for a particular set of interfaces.
    Dot3Colltable EtherLikeMIB_Dot3Colltable

    // A table of descriptive and status information about the MAC Control
    // sublayer on the ethernet-like interfaces attached to a particular system. 
    // There will be one row in this table for each ethernet-like interface in the
    // system which implements the MAC Control sublayer.  If some, but not all, of
    // the ethernet-like interfaces in the system implement the MAC Control
    // sublayer, there will be fewer rows in this table than in the
    // dot3StatsTable.
    Dot3Controltable EtherLikeMIB_Dot3Controltable

    // A table of descriptive and status information about the MAC Control PAUSE
    // function on the ethernet-like interfaces attached to a particular system.
    // There will be one row in this table for each ethernet-like interface in the
    // system which supports the MAC Control PAUSE function (i.e., the 'pause' bit
    // in the corresponding instance of dot3ControlFunctionsSupported is set).  If
    // some, but not all, of the ethernet-like interfaces in the system implement
    // the MAC Control PAUSE function (for example, if some interfaces only
    // support half-duplex), there will be fewer rows in this table than in the
    // dot3StatsTable.
    Dot3Pausetable EtherLikeMIB_Dot3Pausetable

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
    Dot3Hcstatstable EtherLikeMIB_Dot3Hcstatstable
}

func (etherLikeMIB *EtherLikeMIB) GetEntityData() *types.CommonEntityData {
    etherLikeMIB.EntityData.YFilter = etherLikeMIB.YFilter
    etherLikeMIB.EntityData.YangName = "EtherLike-MIB"
    etherLikeMIB.EntityData.BundleName = "cisco_ios_xe"
    etherLikeMIB.EntityData.ParentYangName = "EtherLike-MIB"
    etherLikeMIB.EntityData.SegmentPath = "EtherLike-MIB:EtherLike-MIB"
    etherLikeMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherLikeMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherLikeMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherLikeMIB.EntityData.Children = make(map[string]types.YChild)
    etherLikeMIB.EntityData.Children["dot3StatsTable"] = types.YChild{"Dot3Statstable", &etherLikeMIB.Dot3Statstable}
    etherLikeMIB.EntityData.Children["dot3CollTable"] = types.YChild{"Dot3Colltable", &etherLikeMIB.Dot3Colltable}
    etherLikeMIB.EntityData.Children["dot3ControlTable"] = types.YChild{"Dot3Controltable", &etherLikeMIB.Dot3Controltable}
    etherLikeMIB.EntityData.Children["dot3PauseTable"] = types.YChild{"Dot3Pausetable", &etherLikeMIB.Dot3Pausetable}
    etherLikeMIB.EntityData.Children["dot3HCStatsTable"] = types.YChild{"Dot3Hcstatstable", &etherLikeMIB.Dot3Hcstatstable}
    etherLikeMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherLikeMIB.EntityData)
}

// EtherLikeMIB_Dot3Statstable
// Statistics for a collection of ethernet-like
// interfaces attached to a particular system.
// There will be one row in this table for each
// ethernet-like interface in the system.
type EtherLikeMIB_Dot3Statstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a particular interface to an ethernet-like medium. The type
    // is slice of EtherLikeMIB_Dot3Statstable_Dot3Statsentry.
    Dot3Statsentry []EtherLikeMIB_Dot3Statstable_Dot3Statsentry
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetEntityData() *types.CommonEntityData {
    dot3Statstable.EntityData.YFilter = dot3Statstable.YFilter
    dot3Statstable.EntityData.YangName = "dot3StatsTable"
    dot3Statstable.EntityData.BundleName = "cisco_ios_xe"
    dot3Statstable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3Statstable.EntityData.SegmentPath = "dot3StatsTable"
    dot3Statstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Statstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Statstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Statstable.EntityData.Children = make(map[string]types.YChild)
    dot3Statstable.EntityData.Children["dot3StatsEntry"] = types.YChild{"Dot3Statsentry", nil}
    for i := range dot3Statstable.Dot3Statsentry {
        dot3Statstable.EntityData.Children[types.GetSegmentPath(&dot3Statstable.Dot3Statsentry[i])] = types.YChild{"Dot3Statsentry", &dot3Statstable.Dot3Statsentry[i]}
    }
    dot3Statstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot3Statstable.EntityData)
}

// EtherLikeMIB_Dot3Statstable_Dot3Statsentry
// Statistics for a particular interface to an
// ethernet-like medium.
type EtherLikeMIB_Dot3Statstable_Dot3Statsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index value that uniquely identifies an
    // interface to an ethernet-like medium.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of ifIndex. The type is interface{} with range: 1..2147483647.
    Dot3Statsindex interface{}

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
    Dot3Statsalignmenterrors interface{}

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
    Dot3Statsfcserrors interface{}

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
    Dot3Statssinglecollisionframes interface{}

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
    Dot3Statsmultiplecollisionframes interface{}

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
    Dot3Statssqetesterrors interface{}

    // A count of frames for which the first transmission attempt on a particular
    // interface is delayed because the medium is busy.  The count represented by
    // an instance of this object does not include frames involved in collisions. 
    // This counter does not increment when the interface is operating in
    // full-duplex mode.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Dot3Statsdeferredtransmissions interface{}

    // The number of times that a collision is detected on a particular interface
    // later than one slotTime into the transmission of a packet.  A (late)
    // collision included in a count represented by an instance of this object is
    // also considered as a (generic) collision for purposes of other
    // collision-related statistics.  This counter does not increment when the
    // interface is operating in full-duplex mode.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Dot3Statslatecollisions interface{}

    // A count of frames for which transmission on a particular interface fails
    // due to excessive collisions.  This counter does not increment when the
    // interface is operating in full-duplex mode.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Dot3Statsexcessivecollisions interface{}

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
    Dot3Statsinternalmactransmiterrors interface{}

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
    Dot3Statscarriersenseerrors interface{}

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
    Dot3Statsframetoolongs interface{}

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
    Dot3Statsinternalmacreceiveerrors interface{}

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
    Dot3Statsetherchipset interface{}

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
    Dot3Statssymbolerrors interface{}

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
    // Dot3Statsduplexstatus.
    Dot3Statsduplexstatus interface{}

    // 'true' for interfaces operating at speeds above 1000 Mb/s that support Rate
    // Control through lowering the average data rate of the MAC sublayer, with
    // frame granularity, and 'false' otherwise. The type is bool.
    Dot3Statsratecontrolability interface{}

    // The current Rate Control mode of operation of the MAC sublayer of this
    // interface. The type is Dot3Statsratecontrolstatus.
    Dot3Statsratecontrolstatus interface{}
}

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetEntityData() *types.CommonEntityData {
    dot3Statsentry.EntityData.YFilter = dot3Statsentry.YFilter
    dot3Statsentry.EntityData.YangName = "dot3StatsEntry"
    dot3Statsentry.EntityData.BundleName = "cisco_ios_xe"
    dot3Statsentry.EntityData.ParentYangName = "dot3StatsTable"
    dot3Statsentry.EntityData.SegmentPath = "dot3StatsEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Statsentry.Dot3Statsindex) + "']"
    dot3Statsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Statsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Statsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Statsentry.EntityData.Children = make(map[string]types.YChild)
    dot3Statsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot3Statsentry.EntityData.Leafs["dot3StatsIndex"] = types.YLeaf{"Dot3Statsindex", dot3Statsentry.Dot3Statsindex}
    dot3Statsentry.EntityData.Leafs["dot3StatsAlignmentErrors"] = types.YLeaf{"Dot3Statsalignmenterrors", dot3Statsentry.Dot3Statsalignmenterrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsFCSErrors"] = types.YLeaf{"Dot3Statsfcserrors", dot3Statsentry.Dot3Statsfcserrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsSingleCollisionFrames"] = types.YLeaf{"Dot3Statssinglecollisionframes", dot3Statsentry.Dot3Statssinglecollisionframes}
    dot3Statsentry.EntityData.Leafs["dot3StatsMultipleCollisionFrames"] = types.YLeaf{"Dot3Statsmultiplecollisionframes", dot3Statsentry.Dot3Statsmultiplecollisionframes}
    dot3Statsentry.EntityData.Leafs["dot3StatsSQETestErrors"] = types.YLeaf{"Dot3Statssqetesterrors", dot3Statsentry.Dot3Statssqetesterrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsDeferredTransmissions"] = types.YLeaf{"Dot3Statsdeferredtransmissions", dot3Statsentry.Dot3Statsdeferredtransmissions}
    dot3Statsentry.EntityData.Leafs["dot3StatsLateCollisions"] = types.YLeaf{"Dot3Statslatecollisions", dot3Statsentry.Dot3Statslatecollisions}
    dot3Statsentry.EntityData.Leafs["dot3StatsExcessiveCollisions"] = types.YLeaf{"Dot3Statsexcessivecollisions", dot3Statsentry.Dot3Statsexcessivecollisions}
    dot3Statsentry.EntityData.Leafs["dot3StatsInternalMacTransmitErrors"] = types.YLeaf{"Dot3Statsinternalmactransmiterrors", dot3Statsentry.Dot3Statsinternalmactransmiterrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsCarrierSenseErrors"] = types.YLeaf{"Dot3Statscarriersenseerrors", dot3Statsentry.Dot3Statscarriersenseerrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsFrameTooLongs"] = types.YLeaf{"Dot3Statsframetoolongs", dot3Statsentry.Dot3Statsframetoolongs}
    dot3Statsentry.EntityData.Leafs["dot3StatsInternalMacReceiveErrors"] = types.YLeaf{"Dot3Statsinternalmacreceiveerrors", dot3Statsentry.Dot3Statsinternalmacreceiveerrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsEtherChipSet"] = types.YLeaf{"Dot3Statsetherchipset", dot3Statsentry.Dot3Statsetherchipset}
    dot3Statsentry.EntityData.Leafs["dot3StatsSymbolErrors"] = types.YLeaf{"Dot3Statssymbolerrors", dot3Statsentry.Dot3Statssymbolerrors}
    dot3Statsentry.EntityData.Leafs["dot3StatsDuplexStatus"] = types.YLeaf{"Dot3Statsduplexstatus", dot3Statsentry.Dot3Statsduplexstatus}
    dot3Statsentry.EntityData.Leafs["dot3StatsRateControlAbility"] = types.YLeaf{"Dot3Statsratecontrolability", dot3Statsentry.Dot3Statsratecontrolability}
    dot3Statsentry.EntityData.Leafs["dot3StatsRateControlStatus"] = types.YLeaf{"Dot3Statsratecontrolstatus", dot3Statsentry.Dot3Statsratecontrolstatus}
    return &(dot3Statsentry.EntityData)
}

// EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus represents valuable to justify the redundancy.
type EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus string

const (
    EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus_unknown EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus = "unknown"

    EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus_halfDuplex EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus = "halfDuplex"

    EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus_fullDuplex EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsduplexstatus = "fullDuplex"
)

// EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus represents the MAC sublayer of this interface.
type EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus string

const (
    EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus_rateControlOff EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus = "rateControlOff"

    EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus_rateControlOn EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus = "rateControlOn"

    EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus_unknown EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsratecontrolstatus = "unknown"
)

// EtherLikeMIB_Dot3Colltable
// A collection of collision histograms for a
// particular set of interfaces.
type EtherLikeMIB_Dot3Colltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cell in the histogram of per-frame collisions for a particular interface.
    // An instance of this object represents the frequency of individual MAC
    // frames for which the transmission (successful or otherwise) on a particular
    // interface is accompanied by a particular number of media collisions. The
    // type is slice of EtherLikeMIB_Dot3Colltable_Dot3Collentry.
    Dot3Collentry []EtherLikeMIB_Dot3Colltable_Dot3Collentry
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetEntityData() *types.CommonEntityData {
    dot3Colltable.EntityData.YFilter = dot3Colltable.YFilter
    dot3Colltable.EntityData.YangName = "dot3CollTable"
    dot3Colltable.EntityData.BundleName = "cisco_ios_xe"
    dot3Colltable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3Colltable.EntityData.SegmentPath = "dot3CollTable"
    dot3Colltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Colltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Colltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Colltable.EntityData.Children = make(map[string]types.YChild)
    dot3Colltable.EntityData.Children["dot3CollEntry"] = types.YChild{"Dot3Collentry", nil}
    for i := range dot3Colltable.Dot3Collentry {
        dot3Colltable.EntityData.Children[types.GetSegmentPath(&dot3Colltable.Dot3Collentry[i])] = types.YChild{"Dot3Collentry", &dot3Colltable.Dot3Collentry[i]}
    }
    dot3Colltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot3Colltable.EntityData)
}

// EtherLikeMIB_Dot3Colltable_Dot3Collentry
// A cell in the histogram of per-frame
// collisions for a particular interface.  An
// instance of this object represents the
// frequency of individual MAC frames for which
// the transmission (successful or otherwise) on a
// particular interface is accompanied by a
// particular number of media collisions.
type EtherLikeMIB_Dot3Colltable_Dot3Collentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The number of per-frame media collisions for which
    // a particular collision histogram cell represents the frequency on a
    // particular interface. The type is interface{} with range: 1..16.
    Dot3Collcount interface{}

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
    Dot3Collfrequencies interface{}
}

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetEntityData() *types.CommonEntityData {
    dot3Collentry.EntityData.YFilter = dot3Collentry.YFilter
    dot3Collentry.EntityData.YangName = "dot3CollEntry"
    dot3Collentry.EntityData.BundleName = "cisco_ios_xe"
    dot3Collentry.EntityData.ParentYangName = "dot3CollTable"
    dot3Collentry.EntityData.SegmentPath = "dot3CollEntry" + "[ifIndex='" + fmt.Sprintf("%v", dot3Collentry.Ifindex) + "']" + "[dot3CollCount='" + fmt.Sprintf("%v", dot3Collentry.Dot3Collcount) + "']"
    dot3Collentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Collentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Collentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Collentry.EntityData.Children = make(map[string]types.YChild)
    dot3Collentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot3Collentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", dot3Collentry.Ifindex}
    dot3Collentry.EntityData.Leafs["dot3CollCount"] = types.YLeaf{"Dot3Collcount", dot3Collentry.Dot3Collcount}
    dot3Collentry.EntityData.Leafs["dot3CollFrequencies"] = types.YLeaf{"Dot3Collfrequencies", dot3Collentry.Dot3Collfrequencies}
    return &(dot3Collentry.EntityData)
}

// EtherLikeMIB_Dot3Controltable
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
type EtherLikeMIB_Dot3Controltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about the MAC Control
    // sublayer on a single ethernet-like interface. The type is slice of
    // EtherLikeMIB_Dot3Controltable_Dot3Controlentry.
    Dot3Controlentry []EtherLikeMIB_Dot3Controltable_Dot3Controlentry
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetEntityData() *types.CommonEntityData {
    dot3Controltable.EntityData.YFilter = dot3Controltable.YFilter
    dot3Controltable.EntityData.YangName = "dot3ControlTable"
    dot3Controltable.EntityData.BundleName = "cisco_ios_xe"
    dot3Controltable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3Controltable.EntityData.SegmentPath = "dot3ControlTable"
    dot3Controltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Controltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Controltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Controltable.EntityData.Children = make(map[string]types.YChild)
    dot3Controltable.EntityData.Children["dot3ControlEntry"] = types.YChild{"Dot3Controlentry", nil}
    for i := range dot3Controltable.Dot3Controlentry {
        dot3Controltable.EntityData.Children[types.GetSegmentPath(&dot3Controltable.Dot3Controlentry[i])] = types.YChild{"Dot3Controlentry", &dot3Controltable.Dot3Controlentry[i]}
    }
    dot3Controltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot3Controltable.EntityData)
}

// EtherLikeMIB_Dot3Controltable_Dot3Controlentry
// An entry in the table, containing information
// about the MAC Control sublayer on a single
// ethernet-like interface.
type EtherLikeMIB_Dot3Controltable_Dot3Controlentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsindex
    Dot3Statsindex interface{}

    // A list of the possible MAC Control functions implemented for this
    // interface. The type is map[string]bool.
    Dot3Controlfunctionssupported interface{}

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
    Dot3Controlinunknownopcodes interface{}

    // A count of MAC Control frames received on this interface that contain an
    // opcode that is not supported by this device.  This counter is a 64 bit
    // version of dot3ControlInUnknownOpcodes.  It should be used on interfaces
    // operating at 10 Gb/s or faster.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    Dot3Hccontrolinunknownopcodes interface{}
}

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetEntityData() *types.CommonEntityData {
    dot3Controlentry.EntityData.YFilter = dot3Controlentry.YFilter
    dot3Controlentry.EntityData.YangName = "dot3ControlEntry"
    dot3Controlentry.EntityData.BundleName = "cisco_ios_xe"
    dot3Controlentry.EntityData.ParentYangName = "dot3ControlTable"
    dot3Controlentry.EntityData.SegmentPath = "dot3ControlEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Controlentry.Dot3Statsindex) + "']"
    dot3Controlentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Controlentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Controlentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Controlentry.EntityData.Children = make(map[string]types.YChild)
    dot3Controlentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot3Controlentry.EntityData.Leafs["dot3StatsIndex"] = types.YLeaf{"Dot3Statsindex", dot3Controlentry.Dot3Statsindex}
    dot3Controlentry.EntityData.Leafs["dot3ControlFunctionsSupported"] = types.YLeaf{"Dot3Controlfunctionssupported", dot3Controlentry.Dot3Controlfunctionssupported}
    dot3Controlentry.EntityData.Leafs["dot3ControlInUnknownOpcodes"] = types.YLeaf{"Dot3Controlinunknownopcodes", dot3Controlentry.Dot3Controlinunknownopcodes}
    dot3Controlentry.EntityData.Leafs["dot3HCControlInUnknownOpcodes"] = types.YLeaf{"Dot3Hccontrolinunknownopcodes", dot3Controlentry.Dot3Hccontrolinunknownopcodes}
    return &(dot3Controlentry.EntityData)
}

// EtherLikeMIB_Dot3Pausetable
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
type EtherLikeMIB_Dot3Pausetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about the MAC Control PAUSE
    // function on a single ethernet-like interface. The type is slice of
    // EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry.
    Dot3Pauseentry []EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetEntityData() *types.CommonEntityData {
    dot3Pausetable.EntityData.YFilter = dot3Pausetable.YFilter
    dot3Pausetable.EntityData.YangName = "dot3PauseTable"
    dot3Pausetable.EntityData.BundleName = "cisco_ios_xe"
    dot3Pausetable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3Pausetable.EntityData.SegmentPath = "dot3PauseTable"
    dot3Pausetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Pausetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Pausetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Pausetable.EntityData.Children = make(map[string]types.YChild)
    dot3Pausetable.EntityData.Children["dot3PauseEntry"] = types.YChild{"Dot3Pauseentry", nil}
    for i := range dot3Pausetable.Dot3Pauseentry {
        dot3Pausetable.EntityData.Children[types.GetSegmentPath(&dot3Pausetable.Dot3Pauseentry[i])] = types.YChild{"Dot3Pauseentry", &dot3Pausetable.Dot3Pauseentry[i]}
    }
    dot3Pausetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot3Pausetable.EntityData)
}

// EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry
// An entry in the table, containing information
// about the MAC Control PAUSE function on a single
// ethernet-like interface.
type EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsindex
    Dot3Statsindex interface{}

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
    // support operation at greater than 100 Mb/s. The type is Dot3Pauseadminmode.
    Dot3Pauseadminmode interface{}

    // This object reflects the PAUSE mode currently in use on this interface, as
    // determined by either (1) the result of the auto-negotiation function or (2)
    // if auto-negotiation is not enabled or is not implemented for the active MAU
    // attached to this interface, by the value of dot3PauseAdminMode.  Interfaces
    // operating at 100 Mb/s or less will never return 'enabledXmit(2)' or
    // 'enabledRcv(3)'.  Interfaces operating in half-duplex mode will always
    // return 'disabled(1)'.  Interfaces on which auto-negotiation is enabled but
    // not yet completed should return the value 'disabled(1)'. The type is
    // Dot3Pauseopermode.
    Dot3Pauseopermode interface{}

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
    Dot3Inpauseframes interface{}

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
    Dot3Outpauseframes interface{}

    // A count of MAC Control frames received on this interface with an opcode
    // indicating the PAUSE operation.  This counter does not increment when the
    // interface is operating in half-duplex mode.  This counter is a 64 bit
    // version of dot3InPauseFrames.  It should be used on interfaces operating at
    // 10 Gb/s or faster.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    Dot3Hcinpauseframes interface{}

    // A count of MAC Control frames transmitted on this interface with an opcode
    // indicating the PAUSE operation.  This counter does not increment when the
    // interface is operating in half-duplex mode.  This counter is a 64 bit
    // version of dot3OutPauseFrames.  It should be used on interfaces operating
    // at 10 Gb/s or faster.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    Dot3Hcoutpauseframes interface{}
}

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetEntityData() *types.CommonEntityData {
    dot3Pauseentry.EntityData.YFilter = dot3Pauseentry.YFilter
    dot3Pauseentry.EntityData.YangName = "dot3PauseEntry"
    dot3Pauseentry.EntityData.BundleName = "cisco_ios_xe"
    dot3Pauseentry.EntityData.ParentYangName = "dot3PauseTable"
    dot3Pauseentry.EntityData.SegmentPath = "dot3PauseEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Pauseentry.Dot3Statsindex) + "']"
    dot3Pauseentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Pauseentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Pauseentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Pauseentry.EntityData.Children = make(map[string]types.YChild)
    dot3Pauseentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot3Pauseentry.EntityData.Leafs["dot3StatsIndex"] = types.YLeaf{"Dot3Statsindex", dot3Pauseentry.Dot3Statsindex}
    dot3Pauseentry.EntityData.Leafs["dot3PauseAdminMode"] = types.YLeaf{"Dot3Pauseadminmode", dot3Pauseentry.Dot3Pauseadminmode}
    dot3Pauseentry.EntityData.Leafs["dot3PauseOperMode"] = types.YLeaf{"Dot3Pauseopermode", dot3Pauseentry.Dot3Pauseopermode}
    dot3Pauseentry.EntityData.Leafs["dot3InPauseFrames"] = types.YLeaf{"Dot3Inpauseframes", dot3Pauseentry.Dot3Inpauseframes}
    dot3Pauseentry.EntityData.Leafs["dot3OutPauseFrames"] = types.YLeaf{"Dot3Outpauseframes", dot3Pauseentry.Dot3Outpauseframes}
    dot3Pauseentry.EntityData.Leafs["dot3HCInPauseFrames"] = types.YLeaf{"Dot3Hcinpauseframes", dot3Pauseentry.Dot3Hcinpauseframes}
    dot3Pauseentry.EntityData.Leafs["dot3HCOutPauseFrames"] = types.YLeaf{"Dot3Hcoutpauseframes", dot3Pauseentry.Dot3Hcoutpauseframes}
    return &(dot3Pauseentry.EntityData)
}

// EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode represents at greater than 100 Mb/s.
type EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode string

const (
    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode_disabled EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode = "disabled"

    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode_enabledXmit EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode = "enabledXmit"

    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode_enabledRcv EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode = "enabledRcv"

    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode_enabledXmitAndRcv EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseadminmode = "enabledXmitAndRcv"
)

// EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode represents 'disabled(1)'.
type EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode string

const (
    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode_disabled EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode = "disabled"

    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode_enabledXmit EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode = "enabledXmit"

    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode_enabledRcv EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode = "enabledRcv"

    EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode_enabledXmitAndRcv EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry_Dot3Pauseopermode = "enabledXmitAndRcv"
)

// EtherLikeMIB_Dot3Hcstatstable
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
type EtherLikeMIB_Dot3Hcstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing 64-bit statistics for a single ethernet-like interface.
    // The type is slice of EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry.
    Dot3Hcstatsentry []EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetEntityData() *types.CommonEntityData {
    dot3Hcstatstable.EntityData.YFilter = dot3Hcstatstable.YFilter
    dot3Hcstatstable.EntityData.YangName = "dot3HCStatsTable"
    dot3Hcstatstable.EntityData.BundleName = "cisco_ios_xe"
    dot3Hcstatstable.EntityData.ParentYangName = "EtherLike-MIB"
    dot3Hcstatstable.EntityData.SegmentPath = "dot3HCStatsTable"
    dot3Hcstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Hcstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Hcstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Hcstatstable.EntityData.Children = make(map[string]types.YChild)
    dot3Hcstatstable.EntityData.Children["dot3HCStatsEntry"] = types.YChild{"Dot3Hcstatsentry", nil}
    for i := range dot3Hcstatstable.Dot3Hcstatsentry {
        dot3Hcstatstable.EntityData.Children[types.GetSegmentPath(&dot3Hcstatstable.Dot3Hcstatsentry[i])] = types.YChild{"Dot3Hcstatsentry", &dot3Hcstatstable.Dot3Hcstatsentry[i]}
    }
    dot3Hcstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot3Hcstatstable.EntityData)
}

// EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry
// An entry containing 64-bit statistics for a
// single ethernet-like interface.
type EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsindex
    Dot3Statsindex interface{}

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
    Dot3Hcstatsalignmenterrors interface{}

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
    Dot3Hcstatsfcserrors interface{}

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
    Dot3Hcstatsinternalmactransmiterrors interface{}

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
    Dot3Hcstatsframetoolongs interface{}

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
    Dot3Hcstatsinternalmacreceiveerrors interface{}

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
    Dot3Hcstatssymbolerrors interface{}
}

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetEntityData() *types.CommonEntityData {
    dot3Hcstatsentry.EntityData.YFilter = dot3Hcstatsentry.YFilter
    dot3Hcstatsentry.EntityData.YangName = "dot3HCStatsEntry"
    dot3Hcstatsentry.EntityData.BundleName = "cisco_ios_xe"
    dot3Hcstatsentry.EntityData.ParentYangName = "dot3HCStatsTable"
    dot3Hcstatsentry.EntityData.SegmentPath = "dot3HCStatsEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Hcstatsentry.Dot3Statsindex) + "']"
    dot3Hcstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot3Hcstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot3Hcstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot3Hcstatsentry.EntityData.Children = make(map[string]types.YChild)
    dot3Hcstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot3Hcstatsentry.EntityData.Leafs["dot3StatsIndex"] = types.YLeaf{"Dot3Statsindex", dot3Hcstatsentry.Dot3Statsindex}
    dot3Hcstatsentry.EntityData.Leafs["dot3HCStatsAlignmentErrors"] = types.YLeaf{"Dot3Hcstatsalignmenterrors", dot3Hcstatsentry.Dot3Hcstatsalignmenterrors}
    dot3Hcstatsentry.EntityData.Leafs["dot3HCStatsFCSErrors"] = types.YLeaf{"Dot3Hcstatsfcserrors", dot3Hcstatsentry.Dot3Hcstatsfcserrors}
    dot3Hcstatsentry.EntityData.Leafs["dot3HCStatsInternalMacTransmitErrors"] = types.YLeaf{"Dot3Hcstatsinternalmactransmiterrors", dot3Hcstatsentry.Dot3Hcstatsinternalmactransmiterrors}
    dot3Hcstatsentry.EntityData.Leafs["dot3HCStatsFrameTooLongs"] = types.YLeaf{"Dot3Hcstatsframetoolongs", dot3Hcstatsentry.Dot3Hcstatsframetoolongs}
    dot3Hcstatsentry.EntityData.Leafs["dot3HCStatsInternalMacReceiveErrors"] = types.YLeaf{"Dot3Hcstatsinternalmacreceiveerrors", dot3Hcstatsentry.Dot3Hcstatsinternalmacreceiveerrors}
    dot3Hcstatsentry.EntityData.Leafs["dot3HCStatsSymbolErrors"] = types.YLeaf{"Dot3Hcstatssymbolerrors", dot3Hcstatsentry.Dot3Hcstatssymbolerrors}
    return &(dot3Hcstatsentry.EntityData)
}

