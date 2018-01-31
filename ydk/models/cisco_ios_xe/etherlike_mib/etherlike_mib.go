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

type Dot3Erroriniterror struct {
}

func (id Dot3Erroriniterror) String() string {
	return "EtherLike-MIB:dot3ErrorInitError"
}

type Dot3Testtdr struct {
}

func (id Dot3Testtdr) String() string {
	return "EtherLike-MIB:dot3TestTdr"
}

type Dot3Errorloopbackerror struct {
}

func (id Dot3Errorloopbackerror) String() string {
	return "EtherLike-MIB:dot3ErrorLoopbackError"
}

type Dot3Testloopback struct {
}

func (id Dot3Testloopback) String() string {
	return "EtherLike-MIB:dot3TestLoopBack"
}

// EtherLikeMIB
type EtherLikeMIB struct {
    parent types.Entity
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

func (etherLikeMIB *EtherLikeMIB) GetFilter() yfilter.YFilter { return etherLikeMIB.YFilter }

func (etherLikeMIB *EtherLikeMIB) SetFilter(yf yfilter.YFilter) { etherLikeMIB.YFilter = yf }

func (etherLikeMIB *EtherLikeMIB) GetGoName(yname string) string {
    if yname == "dot3StatsTable" { return "Dot3Statstable" }
    if yname == "dot3CollTable" { return "Dot3Colltable" }
    if yname == "dot3ControlTable" { return "Dot3Controltable" }
    if yname == "dot3PauseTable" { return "Dot3Pausetable" }
    if yname == "dot3HCStatsTable" { return "Dot3Hcstatstable" }
    return ""
}

func (etherLikeMIB *EtherLikeMIB) GetSegmentPath() string {
    return "EtherLike-MIB:EtherLike-MIB"
}

func (etherLikeMIB *EtherLikeMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot3StatsTable" {
        return &etherLikeMIB.Dot3Statstable
    }
    if childYangName == "dot3CollTable" {
        return &etherLikeMIB.Dot3Colltable
    }
    if childYangName == "dot3ControlTable" {
        return &etherLikeMIB.Dot3Controltable
    }
    if childYangName == "dot3PauseTable" {
        return &etherLikeMIB.Dot3Pausetable
    }
    if childYangName == "dot3HCStatsTable" {
        return &etherLikeMIB.Dot3Hcstatstable
    }
    return nil
}

func (etherLikeMIB *EtherLikeMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dot3StatsTable"] = &etherLikeMIB.Dot3Statstable
    children["dot3CollTable"] = &etherLikeMIB.Dot3Colltable
    children["dot3ControlTable"] = &etherLikeMIB.Dot3Controltable
    children["dot3PauseTable"] = &etherLikeMIB.Dot3Pausetable
    children["dot3HCStatsTable"] = &etherLikeMIB.Dot3Hcstatstable
    return children
}

func (etherLikeMIB *EtherLikeMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherLikeMIB *EtherLikeMIB) GetBundleName() string { return "cisco_ios_xe" }

func (etherLikeMIB *EtherLikeMIB) GetYangName() string { return "EtherLike-MIB" }

func (etherLikeMIB *EtherLikeMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (etherLikeMIB *EtherLikeMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (etherLikeMIB *EtherLikeMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (etherLikeMIB *EtherLikeMIB) SetParent(parent types.Entity) { etherLikeMIB.parent = parent }

func (etherLikeMIB *EtherLikeMIB) GetParent() types.Entity { return etherLikeMIB.parent }

func (etherLikeMIB *EtherLikeMIB) GetParentYangName() string { return "EtherLike-MIB" }

// EtherLikeMIB_Dot3Statstable
// Statistics for a collection of ethernet-like
// interfaces attached to a particular system.
// There will be one row in this table for each
// ethernet-like interface in the system.
type EtherLikeMIB_Dot3Statstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics for a particular interface to an ethernet-like medium. The type
    // is slice of EtherLikeMIB_Dot3Statstable_Dot3Statsentry.
    Dot3Statsentry []EtherLikeMIB_Dot3Statstable_Dot3Statsentry
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetFilter() yfilter.YFilter { return dot3Statstable.YFilter }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) SetFilter(yf yfilter.YFilter) { dot3Statstable.YFilter = yf }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetGoName(yname string) string {
    if yname == "dot3StatsEntry" { return "Dot3Statsentry" }
    return ""
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetSegmentPath() string {
    return "dot3StatsTable"
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot3StatsEntry" {
        for _, c := range dot3Statstable.Dot3Statsentry {
            if dot3Statstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLikeMIB_Dot3Statstable_Dot3Statsentry{}
        dot3Statstable.Dot3Statsentry = append(dot3Statstable.Dot3Statsentry, child)
        return &dot3Statstable.Dot3Statsentry[len(dot3Statstable.Dot3Statsentry)-1]
    }
    return nil
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot3Statstable.Dot3Statsentry {
        children[dot3Statstable.Dot3Statsentry[i].GetSegmentPath()] = &dot3Statstable.Dot3Statsentry[i]
    }
    return children
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetYangName() string { return "dot3StatsTable" }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) SetParent(parent types.Entity) { dot3Statstable.parent = parent }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetParent() types.Entity { return dot3Statstable.parent }

func (dot3Statstable *EtherLikeMIB_Dot3Statstable) GetParentYangName() string { return "EtherLike-MIB" }

// EtherLikeMIB_Dot3Statstable_Dot3Statsentry
// Statistics for a particular interface to an
// ethernet-like medium.
type EtherLikeMIB_Dot3Statstable_Dot3Statsentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetFilter() yfilter.YFilter { return dot3Statsentry.YFilter }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) SetFilter(yf yfilter.YFilter) { dot3Statsentry.YFilter = yf }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetGoName(yname string) string {
    if yname == "dot3StatsIndex" { return "Dot3Statsindex" }
    if yname == "dot3StatsAlignmentErrors" { return "Dot3Statsalignmenterrors" }
    if yname == "dot3StatsFCSErrors" { return "Dot3Statsfcserrors" }
    if yname == "dot3StatsSingleCollisionFrames" { return "Dot3Statssinglecollisionframes" }
    if yname == "dot3StatsMultipleCollisionFrames" { return "Dot3Statsmultiplecollisionframes" }
    if yname == "dot3StatsSQETestErrors" { return "Dot3Statssqetesterrors" }
    if yname == "dot3StatsDeferredTransmissions" { return "Dot3Statsdeferredtransmissions" }
    if yname == "dot3StatsLateCollisions" { return "Dot3Statslatecollisions" }
    if yname == "dot3StatsExcessiveCollisions" { return "Dot3Statsexcessivecollisions" }
    if yname == "dot3StatsInternalMacTransmitErrors" { return "Dot3Statsinternalmactransmiterrors" }
    if yname == "dot3StatsCarrierSenseErrors" { return "Dot3Statscarriersenseerrors" }
    if yname == "dot3StatsFrameTooLongs" { return "Dot3Statsframetoolongs" }
    if yname == "dot3StatsInternalMacReceiveErrors" { return "Dot3Statsinternalmacreceiveerrors" }
    if yname == "dot3StatsEtherChipSet" { return "Dot3Statsetherchipset" }
    if yname == "dot3StatsSymbolErrors" { return "Dot3Statssymbolerrors" }
    if yname == "dot3StatsDuplexStatus" { return "Dot3Statsduplexstatus" }
    if yname == "dot3StatsRateControlAbility" { return "Dot3Statsratecontrolability" }
    if yname == "dot3StatsRateControlStatus" { return "Dot3Statsratecontrolstatus" }
    return ""
}

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetSegmentPath() string {
    return "dot3StatsEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Statsentry.Dot3Statsindex) + "']"
}

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot3StatsIndex"] = dot3Statsentry.Dot3Statsindex
    leafs["dot3StatsAlignmentErrors"] = dot3Statsentry.Dot3Statsalignmenterrors
    leafs["dot3StatsFCSErrors"] = dot3Statsentry.Dot3Statsfcserrors
    leafs["dot3StatsSingleCollisionFrames"] = dot3Statsentry.Dot3Statssinglecollisionframes
    leafs["dot3StatsMultipleCollisionFrames"] = dot3Statsentry.Dot3Statsmultiplecollisionframes
    leafs["dot3StatsSQETestErrors"] = dot3Statsentry.Dot3Statssqetesterrors
    leafs["dot3StatsDeferredTransmissions"] = dot3Statsentry.Dot3Statsdeferredtransmissions
    leafs["dot3StatsLateCollisions"] = dot3Statsentry.Dot3Statslatecollisions
    leafs["dot3StatsExcessiveCollisions"] = dot3Statsentry.Dot3Statsexcessivecollisions
    leafs["dot3StatsInternalMacTransmitErrors"] = dot3Statsentry.Dot3Statsinternalmactransmiterrors
    leafs["dot3StatsCarrierSenseErrors"] = dot3Statsentry.Dot3Statscarriersenseerrors
    leafs["dot3StatsFrameTooLongs"] = dot3Statsentry.Dot3Statsframetoolongs
    leafs["dot3StatsInternalMacReceiveErrors"] = dot3Statsentry.Dot3Statsinternalmacreceiveerrors
    leafs["dot3StatsEtherChipSet"] = dot3Statsentry.Dot3Statsetherchipset
    leafs["dot3StatsSymbolErrors"] = dot3Statsentry.Dot3Statssymbolerrors
    leafs["dot3StatsDuplexStatus"] = dot3Statsentry.Dot3Statsduplexstatus
    leafs["dot3StatsRateControlAbility"] = dot3Statsentry.Dot3Statsratecontrolability
    leafs["dot3StatsRateControlStatus"] = dot3Statsentry.Dot3Statsratecontrolstatus
    return leafs
}

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetYangName() string { return "dot3StatsEntry" }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) SetParent(parent types.Entity) { dot3Statsentry.parent = parent }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetParent() types.Entity { return dot3Statsentry.parent }

func (dot3Statsentry *EtherLikeMIB_Dot3Statstable_Dot3Statsentry) GetParentYangName() string { return "dot3StatsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A cell in the histogram of per-frame collisions for a particular interface.
    // An instance of this object represents the frequency of individual MAC
    // frames for which the transmission (successful or otherwise) on a particular
    // interface is accompanied by a particular number of media collisions. The
    // type is slice of EtherLikeMIB_Dot3Colltable_Dot3Collentry.
    Dot3Collentry []EtherLikeMIB_Dot3Colltable_Dot3Collentry
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetFilter() yfilter.YFilter { return dot3Colltable.YFilter }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) SetFilter(yf yfilter.YFilter) { dot3Colltable.YFilter = yf }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetGoName(yname string) string {
    if yname == "dot3CollEntry" { return "Dot3Collentry" }
    return ""
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetSegmentPath() string {
    return "dot3CollTable"
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot3CollEntry" {
        for _, c := range dot3Colltable.Dot3Collentry {
            if dot3Colltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLikeMIB_Dot3Colltable_Dot3Collentry{}
        dot3Colltable.Dot3Collentry = append(dot3Colltable.Dot3Collentry, child)
        return &dot3Colltable.Dot3Collentry[len(dot3Colltable.Dot3Collentry)-1]
    }
    return nil
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot3Colltable.Dot3Collentry {
        children[dot3Colltable.Dot3Collentry[i].GetSegmentPath()] = &dot3Colltable.Dot3Collentry[i]
    }
    return children
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetYangName() string { return "dot3CollTable" }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) SetParent(parent types.Entity) { dot3Colltable.parent = parent }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetParent() types.Entity { return dot3Colltable.parent }

func (dot3Colltable *EtherLikeMIB_Dot3Colltable) GetParentYangName() string { return "EtherLike-MIB" }

// EtherLikeMIB_Dot3Colltable_Dot3Collentry
// A cell in the histogram of per-frame
// collisions for a particular interface.  An
// instance of this object represents the
// frequency of individual MAC frames for which
// the transmission (successful or otherwise) on a
// particular interface is accompanied by a
// particular number of media collisions.
type EtherLikeMIB_Dot3Colltable_Dot3Collentry struct {
    parent types.Entity
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

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetFilter() yfilter.YFilter { return dot3Collentry.YFilter }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) SetFilter(yf yfilter.YFilter) { dot3Collentry.YFilter = yf }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "dot3CollCount" { return "Dot3Collcount" }
    if yname == "dot3CollFrequencies" { return "Dot3Collfrequencies" }
    return ""
}

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetSegmentPath() string {
    return "dot3CollEntry" + "[ifIndex='" + fmt.Sprintf("%v", dot3Collentry.Ifindex) + "']" + "[dot3CollCount='" + fmt.Sprintf("%v", dot3Collentry.Dot3Collcount) + "']"
}

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = dot3Collentry.Ifindex
    leafs["dot3CollCount"] = dot3Collentry.Dot3Collcount
    leafs["dot3CollFrequencies"] = dot3Collentry.Dot3Collfrequencies
    return leafs
}

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetYangName() string { return "dot3CollEntry" }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) SetParent(parent types.Entity) { dot3Collentry.parent = parent }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetParent() types.Entity { return dot3Collentry.parent }

func (dot3Collentry *EtherLikeMIB_Dot3Colltable_Dot3Collentry) GetParentYangName() string { return "dot3CollTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about the MAC Control
    // sublayer on a single ethernet-like interface. The type is slice of
    // EtherLikeMIB_Dot3Controltable_Dot3Controlentry.
    Dot3Controlentry []EtherLikeMIB_Dot3Controltable_Dot3Controlentry
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetFilter() yfilter.YFilter { return dot3Controltable.YFilter }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) SetFilter(yf yfilter.YFilter) { dot3Controltable.YFilter = yf }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetGoName(yname string) string {
    if yname == "dot3ControlEntry" { return "Dot3Controlentry" }
    return ""
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetSegmentPath() string {
    return "dot3ControlTable"
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot3ControlEntry" {
        for _, c := range dot3Controltable.Dot3Controlentry {
            if dot3Controltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLikeMIB_Dot3Controltable_Dot3Controlentry{}
        dot3Controltable.Dot3Controlentry = append(dot3Controltable.Dot3Controlentry, child)
        return &dot3Controltable.Dot3Controlentry[len(dot3Controltable.Dot3Controlentry)-1]
    }
    return nil
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot3Controltable.Dot3Controlentry {
        children[dot3Controltable.Dot3Controlentry[i].GetSegmentPath()] = &dot3Controltable.Dot3Controlentry[i]
    }
    return children
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetYangName() string { return "dot3ControlTable" }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) SetParent(parent types.Entity) { dot3Controltable.parent = parent }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetParent() types.Entity { return dot3Controltable.parent }

func (dot3Controltable *EtherLikeMIB_Dot3Controltable) GetParentYangName() string { return "EtherLike-MIB" }

// EtherLikeMIB_Dot3Controltable_Dot3Controlentry
// An entry in the table, containing information
// about the MAC Control sublayer on a single
// ethernet-like interface.
type EtherLikeMIB_Dot3Controltable_Dot3Controlentry struct {
    parent types.Entity
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

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetFilter() yfilter.YFilter { return dot3Controlentry.YFilter }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) SetFilter(yf yfilter.YFilter) { dot3Controlentry.YFilter = yf }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetGoName(yname string) string {
    if yname == "dot3StatsIndex" { return "Dot3Statsindex" }
    if yname == "dot3ControlFunctionsSupported" { return "Dot3Controlfunctionssupported" }
    if yname == "dot3ControlInUnknownOpcodes" { return "Dot3Controlinunknownopcodes" }
    if yname == "dot3HCControlInUnknownOpcodes" { return "Dot3Hccontrolinunknownopcodes" }
    return ""
}

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetSegmentPath() string {
    return "dot3ControlEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Controlentry.Dot3Statsindex) + "']"
}

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot3StatsIndex"] = dot3Controlentry.Dot3Statsindex
    leafs["dot3ControlFunctionsSupported"] = dot3Controlentry.Dot3Controlfunctionssupported
    leafs["dot3ControlInUnknownOpcodes"] = dot3Controlentry.Dot3Controlinunknownopcodes
    leafs["dot3HCControlInUnknownOpcodes"] = dot3Controlentry.Dot3Hccontrolinunknownopcodes
    return leafs
}

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetYangName() string { return "dot3ControlEntry" }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) SetParent(parent types.Entity) { dot3Controlentry.parent = parent }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetParent() types.Entity { return dot3Controlentry.parent }

func (dot3Controlentry *EtherLikeMIB_Dot3Controltable_Dot3Controlentry) GetParentYangName() string { return "dot3ControlTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about the MAC Control PAUSE
    // function on a single ethernet-like interface. The type is slice of
    // EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry.
    Dot3Pauseentry []EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetFilter() yfilter.YFilter { return dot3Pausetable.YFilter }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) SetFilter(yf yfilter.YFilter) { dot3Pausetable.YFilter = yf }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetGoName(yname string) string {
    if yname == "dot3PauseEntry" { return "Dot3Pauseentry" }
    return ""
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetSegmentPath() string {
    return "dot3PauseTable"
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot3PauseEntry" {
        for _, c := range dot3Pausetable.Dot3Pauseentry {
            if dot3Pausetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry{}
        dot3Pausetable.Dot3Pauseentry = append(dot3Pausetable.Dot3Pauseentry, child)
        return &dot3Pausetable.Dot3Pauseentry[len(dot3Pausetable.Dot3Pauseentry)-1]
    }
    return nil
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot3Pausetable.Dot3Pauseentry {
        children[dot3Pausetable.Dot3Pauseentry[i].GetSegmentPath()] = &dot3Pausetable.Dot3Pauseentry[i]
    }
    return children
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetYangName() string { return "dot3PauseTable" }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) SetParent(parent types.Entity) { dot3Pausetable.parent = parent }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetParent() types.Entity { return dot3Pausetable.parent }

func (dot3Pausetable *EtherLikeMIB_Dot3Pausetable) GetParentYangName() string { return "EtherLike-MIB" }

// EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry
// An entry in the table, containing information
// about the MAC Control PAUSE function on a single
// ethernet-like interface.
type EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry struct {
    parent types.Entity
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

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetFilter() yfilter.YFilter { return dot3Pauseentry.YFilter }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) SetFilter(yf yfilter.YFilter) { dot3Pauseentry.YFilter = yf }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetGoName(yname string) string {
    if yname == "dot3StatsIndex" { return "Dot3Statsindex" }
    if yname == "dot3PauseAdminMode" { return "Dot3Pauseadminmode" }
    if yname == "dot3PauseOperMode" { return "Dot3Pauseopermode" }
    if yname == "dot3InPauseFrames" { return "Dot3Inpauseframes" }
    if yname == "dot3OutPauseFrames" { return "Dot3Outpauseframes" }
    if yname == "dot3HCInPauseFrames" { return "Dot3Hcinpauseframes" }
    if yname == "dot3HCOutPauseFrames" { return "Dot3Hcoutpauseframes" }
    return ""
}

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetSegmentPath() string {
    return "dot3PauseEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Pauseentry.Dot3Statsindex) + "']"
}

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot3StatsIndex"] = dot3Pauseentry.Dot3Statsindex
    leafs["dot3PauseAdminMode"] = dot3Pauseentry.Dot3Pauseadminmode
    leafs["dot3PauseOperMode"] = dot3Pauseentry.Dot3Pauseopermode
    leafs["dot3InPauseFrames"] = dot3Pauseentry.Dot3Inpauseframes
    leafs["dot3OutPauseFrames"] = dot3Pauseentry.Dot3Outpauseframes
    leafs["dot3HCInPauseFrames"] = dot3Pauseentry.Dot3Hcinpauseframes
    leafs["dot3HCOutPauseFrames"] = dot3Pauseentry.Dot3Hcoutpauseframes
    return leafs
}

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetYangName() string { return "dot3PauseEntry" }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) SetParent(parent types.Entity) { dot3Pauseentry.parent = parent }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetParent() types.Entity { return dot3Pauseentry.parent }

func (dot3Pauseentry *EtherLikeMIB_Dot3Pausetable_Dot3Pauseentry) GetParentYangName() string { return "dot3PauseTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing 64-bit statistics for a single ethernet-like interface.
    // The type is slice of EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry.
    Dot3Hcstatsentry []EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetFilter() yfilter.YFilter { return dot3Hcstatstable.YFilter }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) SetFilter(yf yfilter.YFilter) { dot3Hcstatstable.YFilter = yf }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetGoName(yname string) string {
    if yname == "dot3HCStatsEntry" { return "Dot3Hcstatsentry" }
    return ""
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetSegmentPath() string {
    return "dot3HCStatsTable"
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot3HCStatsEntry" {
        for _, c := range dot3Hcstatstable.Dot3Hcstatsentry {
            if dot3Hcstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry{}
        dot3Hcstatstable.Dot3Hcstatsentry = append(dot3Hcstatstable.Dot3Hcstatsentry, child)
        return &dot3Hcstatstable.Dot3Hcstatsentry[len(dot3Hcstatstable.Dot3Hcstatsentry)-1]
    }
    return nil
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot3Hcstatstable.Dot3Hcstatsentry {
        children[dot3Hcstatstable.Dot3Hcstatsentry[i].GetSegmentPath()] = &dot3Hcstatstable.Dot3Hcstatsentry[i]
    }
    return children
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetYangName() string { return "dot3HCStatsTable" }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) SetParent(parent types.Entity) { dot3Hcstatstable.parent = parent }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetParent() types.Entity { return dot3Hcstatstable.parent }

func (dot3Hcstatstable *EtherLikeMIB_Dot3Hcstatstable) GetParentYangName() string { return "EtherLike-MIB" }

// EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry
// An entry containing 64-bit statistics for a
// single ethernet-like interface.
type EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry struct {
    parent types.Entity
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

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetFilter() yfilter.YFilter { return dot3Hcstatsentry.YFilter }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) SetFilter(yf yfilter.YFilter) { dot3Hcstatsentry.YFilter = yf }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetGoName(yname string) string {
    if yname == "dot3StatsIndex" { return "Dot3Statsindex" }
    if yname == "dot3HCStatsAlignmentErrors" { return "Dot3Hcstatsalignmenterrors" }
    if yname == "dot3HCStatsFCSErrors" { return "Dot3Hcstatsfcserrors" }
    if yname == "dot3HCStatsInternalMacTransmitErrors" { return "Dot3Hcstatsinternalmactransmiterrors" }
    if yname == "dot3HCStatsFrameTooLongs" { return "Dot3Hcstatsframetoolongs" }
    if yname == "dot3HCStatsInternalMacReceiveErrors" { return "Dot3Hcstatsinternalmacreceiveerrors" }
    if yname == "dot3HCStatsSymbolErrors" { return "Dot3Hcstatssymbolerrors" }
    return ""
}

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetSegmentPath() string {
    return "dot3HCStatsEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", dot3Hcstatsentry.Dot3Statsindex) + "']"
}

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot3StatsIndex"] = dot3Hcstatsentry.Dot3Statsindex
    leafs["dot3HCStatsAlignmentErrors"] = dot3Hcstatsentry.Dot3Hcstatsalignmenterrors
    leafs["dot3HCStatsFCSErrors"] = dot3Hcstatsentry.Dot3Hcstatsfcserrors
    leafs["dot3HCStatsInternalMacTransmitErrors"] = dot3Hcstatsentry.Dot3Hcstatsinternalmactransmiterrors
    leafs["dot3HCStatsFrameTooLongs"] = dot3Hcstatsentry.Dot3Hcstatsframetoolongs
    leafs["dot3HCStatsInternalMacReceiveErrors"] = dot3Hcstatsentry.Dot3Hcstatsinternalmacreceiveerrors
    leafs["dot3HCStatsSymbolErrors"] = dot3Hcstatsentry.Dot3Hcstatssymbolerrors
    return leafs
}

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetYangName() string { return "dot3HCStatsEntry" }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) SetParent(parent types.Entity) { dot3Hcstatsentry.parent = parent }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetParent() types.Entity { return dot3Hcstatsentry.parent }

func (dot3Hcstatsentry *EtherLikeMIB_Dot3Hcstatstable_Dot3Hcstatsentry) GetParentYangName() string { return "dot3HCStatsTable" }

