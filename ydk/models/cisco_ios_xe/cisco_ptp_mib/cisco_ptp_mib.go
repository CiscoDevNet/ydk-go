// The MIB module for PTPv2 (IEEE1588 - 2008)
// 
// Overview of PTPv2 (IEEE 1588-2008)
// 
// This IEEE standard defines a protocol enabling precise
// synchronization of clocks in measurement and control systems
// implemented with packet-based networks, the IEEE Standard PTPv2
// 1588 (2008). This MIB does not address the standard IEEE
// 1588 (2002). The protocol is applicable to network elements
// communicating using IP. The protocol enables heterogeneous
// systems that include clocks of various inherent precision,
// resolution, and stability to synchronize to a grandmaster
// clock.
// The protocol supports system-wide synchronization accuracy in
// the sub-microsecond range with minimal network and local clock
// computing resources. The standard uses UDP/IP. It includes
// formal mechanisms for message extensions, higher sampling
// rates, correction for asymmetry, a clock type to reduce error
// accumulation in large topologies, and specifications on how to
// incorporate the resulting additional data into the
// synchronization protocol. The standard defines conformance and
// management capability also.
// 
// MIB description
// 
// This MIB is to support the Precision Timing Protocol (PTP)
// feature of Cisco System devices.
// 
// Acronyms:
//          ARB       arbitrary
//          BMC       best master clock
//          CAN       Controller Area Network
//          CP        Communication Profile
//                    [according to IEC 61784-1:200710]
//          CPF       Communication Profile Family
//                    [according to IEC 61784-1:2007]
//          DS        Differentiated Service
//          E2E       End-to-End
//          E2ETC     End-to-End Transparent Clock
//          EUI       Extended Unique Identifier.
//          FFO       Fractional Frequency Offset
//          GPS       Global Positioning System
//          IANA      Internet Assigned Numbers Authority
//          ICV       Integrity Check Value
//          ID        Identification
//          IPv4      Internet Protocol version 4
//          IPv6      Internet Protocol version 6
//          JD        Julian Date
//          JDN       Julian Day Number
//          MAC       Media Access Control
//                    [according to IEEE Std 802.3-2005]
//          MJD       Modified Julian Day
//          NIST      National Institute of Standards and
// Technology                    (see www.nist.gov)
//          NTP       Network Time Protocol (see IETF RFC 1305
// [B7])
//          OUI       Organizational Unique Identifier(allocated
// by
// the IEEE)
//          P2P       Peer-to-Peer
//          P2PTC     Peer-To-Peer Transparent Clock
//          PHY       physical layer [according to IEEE Std
// 802.3-2005]
//          POSIX     Portable Operating System Interface
//                    (see ISO/IEC 9945:2003)
//          PPS       Pulse per Second
//          PTP       Precision Time Protocol
//          SA        Security Associations
//          SNTP      Simple Network Time Protocol
//          SOF       Start of Frame
//          TAI       International Atomic Time
//          TC        Traffic Class
//          TC        Transparent Clock
//          TLV       Type, Length, Value [according to IEEE Std
// 802.1AB]
//          ToD       Time of Day Synchronization
//          ToS       Type of Service
//          UCMM      UnConnect Message Manager
//          UDP/IP    User Datagram Protocol
//          UTC       Coordinated Universal Time
// 
// References:
//    [1]  Precision clock synchronization protocol for networked
// measurement and control systems - IEC 61588 IEEE 1588(tm)
// Edition 2.0 2009-02
// 
// 
// Definitions from [1] section 3.1
// 
// Accuracy:
// The mean of the time or frequency error between the clock under
// test and a perfect reference clock, over an ensemble of
// measurements.  Stability is a measure of how the mean varies
// with respect to variables such as time, temperature, and so on.
// 
// The precision is a measure of the deviation of the error from
// the mean.
// 
// Atomic process:
// A process is atomic if the values of all inputs to the process
// are not permitted to change until all of the results of the
// process are instantiated, and the outputs of the process are
// not visible to other processes until the processing of each
// output is complete.
// 
// Boundary clock:
// A clock that has multiple Precision Time Protocol(PTP) ports in
// a domain and maintains the timescale used in the domain.  It
// may serve as the source of time, i.e., be a master clock, and
// may synchronize to another clock, i.e., be a slave clock.
// 
// Boundary node clock:
// A clock that has multiple Precision Time Protocol(PTP) ports in
// a domain and maintains the timescale used in the domain. It
// differs from the boundary clock in that the clock roles can
// change.
// 
// Clock:
// A node participating in the Precision Time Protocol (PTP) that
// is capable of providing a measurement of the passage of time
// since a defined epoch.
// 
// Domain:
// A logical grouping of clocks that synchronize to each other
// using the protocol, but that are not necessarily synchronized
// to clocks in another domain.
// 
// End-to-end transparent clock:
// A transparent clock that supports the use of the end-to-end
// delay measurement mechanism between slave clocks and the master
// clock.  Each node must measure the residence time of PTP event
// messages and accumulate it in Correction Field.
// 
// Epoch:
// The origin of a timescale.
// 
// Event:
// An abstraction of the mechanism by which signals or conditions
// are generated and represented.
// 
// Foreign master:
// An ordinary or boundary clock sending Announce messages to
// another clock that is not the current master recognized by the
// other clock.
// 
// Grandmaster clock:
// Within a domain, a clock that is the ultimate source of time
// for clock synchronization using the protocol.
// 
// Holdover:
// A clock previously synchronized/syntonized to another clock
// (normally a primary reference or a master clock) but now
// free-running based on its own internal oscillator, whose
// frequency is being adjusted using data acquired while it had
// been synchronized/syntonized to the other clock.  It is said to
// be in holdover or in the holdover mode, as long as it is within
// its accuracy requirements.
// 
// Link:
// A network segment between two Precision Time Protocol ports
// supporting the peer delay mechanism of this standard.  The peer
// delay mechanism is designed to measure the propagation time
// over such a link.
// 
// Management node:
// A device that configures and monitors clocks.
// 
// Master clock:
// In the context of a single Precision Time Protocol
// communication path, a clock that is the source of time to which
// all other clocks on that path synchronize.
// 
// Message timestamp point:
// A point within a Precision Time Protocol event message serving
// as a reference point in the message.  A timestamp is defined by
// the instant a message timestamp point passes the reference
// plane of a clock.
// 
// Multicast communication:
// A communication model in which each Precision Time Protocol
// message sent from any PTP port is capable of being received and
// processed by all PTP ports on the same PTP communication path.
// 
// Node:
// A device that can issue or receive Precision Time Protocol
// communications on a network.
// 
// One-step clock:
// A clock that provides time information using a single event
// message.
// 
// On-pass support:
// Indicates that each node in the synchronization chain from
// master to slave can support IEEE-1588.
// 
// Ordinary clock:
// A clock that has a single Precision Time Protocol port in a
// domain and maintains the timescale used in the domain.  It may
// serve as a source of time, i.e., be a master clock, or may
// synchronize to another clock, i.e., be a slave clock.
// 
// Parent clock:
// The master clock to which a clock is synchronized.
// 
// 
// Peer-to-peer transparent clock:
// A transparent clock that, in addition to providing Precision
// Time Protocol event transit time information, also provides
// corrections for the propagation delay of the link connected to
// the port receiving the PTP event message.  In the presence of
// peer-to-peer transparent clocks, delay measurements between
// slave clocks and the master clock are performed using the
// peer-to-peer delay measurement mechanism.
// 
// 
// Phase change rate:
// The observed rate of change in the measured time with respect
// to the reference time.  The phase change rate is equal to the
// fractional frequency offset between the measured frequency and
// the reference frequency.
// 
// PortNumber:
// An index identifying a specific Precision Time Protocol port on
// a PTP node.
// 
// Primary reference:
// A source of time and or frequency that is traceable to
// international standards.
// 
// Profile:
// The set of allowed Precision Time Protocol features applicable
// to a device.
// 
// Precision Time Protocol communication:
// Information used in the operation of the protocol, transmitted
// in a PTP message over a PTP communication path.
// 
// Precision Time Protocol communication path: The signaling path
// portion of a particular network enabling direct communication
// among ordinary and boundary clocks.
// 
// Precision Time Protocol node:
// PTP ordinary, boundary, or transparent clock or a device that
// generates or parses PTP messages.
// 
// Precision Time Protocol port:
// A logical access point of a clock for PTP communications to the
// communications network.
// 
// Recognized standard time source:
// A recognized standard time source is a source external to
// Precision Time Protocol that provides time and/or frequency as
// appropriate that is traceable to the international standards
// laboratories maintaining clocks that form the basis for the
// International Atomic Time and Universal Coordinated Time
// timescales.  Examples of these are Global Positioning System,
// NTP, and National Institute of Standards and Technology (NIST)
// timeservers.
// 
// Requestor:
// The port implementing the peer-to-peer delay mechanism that
// initiates the mechanism by sending a Pdelay_Req message.
// 
// Responder:
// The port responding to the receipt of a Pdelay_Req message as
// part of the operation of the peer-to-peer delay mechanism.
// 
// Synchronized clocks:
// Two clocks are synchronized to a specified uncertainty if they
// have the same  epoch and their measurements of the time of a
// single event at an arbitrary time differ by no more than that
// uncertainty.
// 
// Syntonized clocks:
// Two clocks are syntonized if the duration of the second is the
// same on both, which means the time as measured by each advances
// at the same rate. They may or may not share the same epoch.
// 
// Time of Day:
// 
// 
// Timeout:
// A mechanism for terminating requested activity that, at least
// from the requester's perspective, does not complete within the
// specified time.
// 
// Timescale:
// A linear measure of time from an epoch.
// 
// Traceability:
// A property of the result of a measurement or the value of a
// standard whereby it can be related to stated references,
// usually national or international standards, through an unbroken
// chain of comparisons all having stated uncertainties.
// 
// Translation device:
// A boundary clock or, in some cases, a transparent clock that
// translates the protocol messages between regions implementing
// different transport and messaging protocols, between different
// versions of IEEE Std 1588-2008/IEC 61588:2009, or different
// Precision Time Protocol profiles.
// 
// transparent clock:
// A device that measures the time taken for a Precision Time
// Protocol event message to transit the device and provides this
// information to clocks receiving this PTP event message.
// 
// Two-step clock:
// A clock that provides time information using the combination of
// an event message and a subsequent general message.
// 
// The below table specifies the object formats of the various
// textual conventions used.
// 
// Data type mapping     Textual Convention  SYNTAX
// --------------------  ------------------  ---------------------
// 5.3.2 TimeInterval    ClockTimeInterval   OCTET
// STRING(SIZE(1..255))
// 5.3.3 Timestamp       ClockTimestamp      OCTET STRING(SIZE(6))
// 5.3.4 ClockIdentity   ClockIdentity       OCTET
// STRING(SIZE(1..255))
// 5.3.5 PortIdentity    ClockPortNumber     INTEGER(1..65535)
// 5.3.7 ClockQuality    ClockQualityClassType
// 
// Simple master-slave hierarchy [1] section 6.6.2.4
// 
//   ---------------
//   - Ordinary    -
//   - Clock(1)    -
//   - GrandMaster -
//   -------M-------
//          |
//          1
//          |
//   -------S-------------------------------
//   - Boundary                            -
//   - Clock(1)                            -
//   --------------M------------------M-----
//                 |                  |
//                 2                  3
//                 |                  |
//            -----S------     -------S------------------
//            - Ordinary -     - Boundary               -
//            - Clock(2) -     - Clock(2)               -
//            ------------     -----M-------------M------
//                                  |             |
//                                  4             5
//                                  |             |
//                             -----S------  -----S------
//                             - Ordinary -  - Ordinary -
//                             - Clock(3) -  - Clock(4) -
//                             ------------  ------------
// 
//   Grandmaster
// 
//   Boundary Clock(0-N)   Ordinary Clocks(0-N)
//   Ordinary Clocks(0-N)
// 
// 
//  Relationship cardinality
//     PTP system 1 : N PTP Clock
//     PTP Clock 1 : 1 Domain
//     PTP Clock 1 : N PTP Ports
//     PTP Port N : N Physical Port (interface in IF-MIB)
// 
// Transparent clock diagram from section 6.7.1.3 of [1]
// 
// 
//           +----------------------------+
//           |     Boundary clock - 1     |
//           +----------------------------+
//           |                       |
//           |                       |
//   +-- A --+                       B
//   |                               |
// +---------------------+           |
// |   Ordinary clock - 1|           |
// +---------------------+           |
//                                   +----------------------+
// +--------------+                  |      End-to-end      |
// |  Ordinary    |------------------|  transparent clock-  |
// |  clock 1-1   |                  |       1 - 1          |
// +--------------+                  +----------------------+
//                                    |
//                                    |
//                                    C
//                                    |
//                                    |
//                                   +----------------------+
// +--------------+                  |      End-to-end      |
// |  Ordinary    |------------------|  transparent clock-  |
// |  clock 1-2   |                  |       1 - 2          |
// +--------------+                  +----------------------+
// 
// 
// The MIB refers to the sections of the IEEE 1588 standard for
// reference. Throughout the MIB various secions from the standard
// are referenced
package cisco_ptp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ptp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-PTP-MIB CISCO-PTP-MIB}", reflect.TypeOf(CISCOPTPMIB{}))
    ydk.RegisterEntity("CISCO-PTP-MIB:CISCO-PTP-MIB", reflect.TypeOf(CISCOPTPMIB{}))
}

// ClockMechanismType represents                                  the delay mechanism.
type ClockMechanismType string

const (
    ClockMechanismType_e2e ClockMechanismType = "e2e"

    ClockMechanismType_p2p ClockMechanismType = "p2p"

    ClockMechanismType_disabled ClockMechanismType = "disabled"
)

// ClockPortState represents                           selected master port.
type ClockPortState string

const (
    ClockPortState_initializing ClockPortState = "initializing"

    ClockPortState_faulty ClockPortState = "faulty"

    ClockPortState_disabled ClockPortState = "disabled"

    ClockPortState_listening ClockPortState = "listening"

    ClockPortState_preMaster ClockPortState = "preMaster"

    ClockPortState_master ClockPortState = "master"

    ClockPortState_passive ClockPortState = "passive"

    ClockPortState_uncalibrated ClockPortState = "uncalibrated"

    ClockPortState_slave ClockPortState = "slave"
)

// ClockTimeSourceType represents the protocol.
type ClockTimeSourceType string

const (
    ClockTimeSourceType_atomicClock ClockTimeSourceType = "atomicClock"

    ClockTimeSourceType_gps ClockTimeSourceType = "gps"

    ClockTimeSourceType_terrestrialRadio ClockTimeSourceType = "terrestrialRadio"

    ClockTimeSourceType_ptp ClockTimeSourceType = "ptp"

    ClockTimeSourceType_ntp ClockTimeSourceType = "ntp"

    ClockTimeSourceType_handSet ClockTimeSourceType = "handSet"

    ClockTimeSourceType_other ClockTimeSourceType = "other"

    ClockTimeSourceType_internalOsillator ClockTimeSourceType = "internalOsillator"
)

// ClockProfileType represents a device.
type ClockProfileType string

const (
    ClockProfileType_default_ ClockProfileType = "default"

    ClockProfileType_telecom ClockProfileType = "telecom"

    ClockProfileType_vendorspecific ClockProfileType = "vendorspecific"
)

// ClockRoleType represents                           another clock (master).
type ClockRoleType string

const (
    ClockRoleType_master ClockRoleType = "master"

    ClockRoleType_slave ClockRoleType = "slave"
)

// ClockType represents The clock types as defined in the MIB module description.
type ClockType string

const (
    ClockType_ordinaryClock ClockType = "ordinaryClock"

    ClockType_boundaryClock ClockType = "boundaryClock"

    ClockType_transparentClock ClockType = "transparentClock"

    ClockType_boundaryNode ClockType = "boundaryNode"
)

// ClockTxModeType represents multicast-mix. Using multicast-unicast communication channel
type ClockTxModeType string

const (
    ClockTxModeType_unicast ClockTxModeType = "unicast"

    ClockTxModeType_multicast ClockTxModeType = "multicast"

    ClockTxModeType_multicastmix ClockTxModeType = "multicastmix"
)

// ClockQualityAccuracyType represents the protocol.
type ClockQualityAccuracyType string

const (
    ClockQualityAccuracyType_reserved00 ClockQualityAccuracyType = "reserved00"

    ClockQualityAccuracyType_nanoSecond25 ClockQualityAccuracyType = "nanoSecond25"

    ClockQualityAccuracyType_nanoSecond100 ClockQualityAccuracyType = "nanoSecond100"

    ClockQualityAccuracyType_nanoSecond250 ClockQualityAccuracyType = "nanoSecond250"

    ClockQualityAccuracyType_microSec1 ClockQualityAccuracyType = "microSec1"

    ClockQualityAccuracyType_microSec2dot5 ClockQualityAccuracyType = "microSec2dot5"

    ClockQualityAccuracyType_microSec10 ClockQualityAccuracyType = "microSec10"

    ClockQualityAccuracyType_microSec25 ClockQualityAccuracyType = "microSec25"

    ClockQualityAccuracyType_microSec100 ClockQualityAccuracyType = "microSec100"

    ClockQualityAccuracyType_microSec250 ClockQualityAccuracyType = "microSec250"

    ClockQualityAccuracyType_milliSec1 ClockQualityAccuracyType = "milliSec1"

    ClockQualityAccuracyType_milliSec2dot5 ClockQualityAccuracyType = "milliSec2dot5"

    ClockQualityAccuracyType_milliSec10 ClockQualityAccuracyType = "milliSec10"

    ClockQualityAccuracyType_milliSec25 ClockQualityAccuracyType = "milliSec25"

    ClockQualityAccuracyType_milliSec100 ClockQualityAccuracyType = "milliSec100"

    ClockQualityAccuracyType_milliSec250 ClockQualityAccuracyType = "milliSec250"

    ClockQualityAccuracyType_second1 ClockQualityAccuracyType = "second1"

    ClockQualityAccuracyType_second10 ClockQualityAccuracyType = "second10"

    ClockQualityAccuracyType_secondGreater10 ClockQualityAccuracyType = "secondGreater10"

    ClockQualityAccuracyType_unknown ClockQualityAccuracyType = "unknown"

    ClockQualityAccuracyType_reserved255 ClockQualityAccuracyType = "reserved255"
)

// ClockStateType represents                         frequency and phase.
type ClockStateType string

const (
    ClockStateType_freerun ClockStateType = "freerun"

    ClockStateType_holdover ClockStateType = "holdover"

    ClockStateType_acquiring ClockStateType = "acquiring"

    ClockStateType_frequencyLocked ClockStateType = "frequencyLocked"

    ClockStateType_phaseAligned ClockStateType = "phaseAligned"
)

// CISCOPTPMIB
type CISCOPTPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciscoptpmibsysteminfo CISCOPTPMIB_Ciscoptpmibsysteminfo

    // Table of count information about the PTP system for all domains.
    Cptpsystemtable CISCOPTPMIB_Cptpsystemtable

    // Table of information about the PTP system for all clock modes -- ordinary,
    // boundary or transparent.
    Cptpsystemdomaintable CISCOPTPMIB_Cptpsystemdomaintable

    // Table of information about the PTP system for a given domain.
    Cptpclocknodetable CISCOPTPMIB_Cptpclocknodetable

    // Table of information about the PTP clock Current Datasets for all domains.
    Cptpclockcurrentdstable CISCOPTPMIB_Cptpclockcurrentdstable

    // Table of information about the PTP clock Parent Datasets for all domains.
    Cptpclockparentdstable CISCOPTPMIB_Cptpclockparentdstable

    // Table of information about the PTP clock Default Datasets for all domains.
    Cptpclockdefaultdstable CISCOPTPMIB_Cptpclockdefaultdstable

    // Table of information about the PTP clock Running Datasets for all domains.
    Cptpclockrunningtable CISCOPTPMIB_Cptpclockrunningtable

    // Table of information about the PTP clock Timeproperties Datasets for all
    // domains.
    Cptpclocktimepropertiesdstable CISCOPTPMIB_Cptpclocktimepropertiesdstable

    // Table of information about the PTP Transparent clock Default Datasets for
    // all domains.
    Cptpclocktransdefaultdstable CISCOPTPMIB_Cptpclocktransdefaultdstable

    // Table of information about the clock ports for a particular domain.
    Cptpclockporttable CISCOPTPMIB_Cptpclockporttable

    // Table of information about the clock ports dataset for a particular domain.
    Cptpclockportdstable CISCOPTPMIB_Cptpclockportdstable

    // Table of information about the clock ports running dataset for a particular
    // domain.
    Cptpclockportrunningtable CISCOPTPMIB_Cptpclockportrunningtable

    // Table of information about the Transparent clock ports running dataset for
    // a particular domain.
    Cptpclockporttransdstable CISCOPTPMIB_Cptpclockporttransdstable

    // Table of information about a given port's associated ports.  For a master
    // port - multiple slave ports which have established sessions with the
    // current master port.   For a slave port - the list of masters available for
    // a given slave port.   Session information (pkts, errors) to be displayed
    // based on availability and scenario.
    Cptpclockportassociatetable CISCOPTPMIB_Cptpclockportassociatetable
}

func (cISCOPTPMIB *CISCOPTPMIB) GetFilter() yfilter.YFilter { return cISCOPTPMIB.YFilter }

func (cISCOPTPMIB *CISCOPTPMIB) SetFilter(yf yfilter.YFilter) { cISCOPTPMIB.YFilter = yf }

func (cISCOPTPMIB *CISCOPTPMIB) GetGoName(yname string) string {
    if yname == "ciscoPtpMIBSystemInfo" { return "Ciscoptpmibsysteminfo" }
    if yname == "cPtpSystemTable" { return "Cptpsystemtable" }
    if yname == "cPtpSystemDomainTable" { return "Cptpsystemdomaintable" }
    if yname == "cPtpClockNodeTable" { return "Cptpclocknodetable" }
    if yname == "cPtpClockCurrentDSTable" { return "Cptpclockcurrentdstable" }
    if yname == "cPtpClockParentDSTable" { return "Cptpclockparentdstable" }
    if yname == "cPtpClockDefaultDSTable" { return "Cptpclockdefaultdstable" }
    if yname == "cPtpClockRunningTable" { return "Cptpclockrunningtable" }
    if yname == "cPtpClockTimePropertiesDSTable" { return "Cptpclocktimepropertiesdstable" }
    if yname == "cPtpClockTransDefaultDSTable" { return "Cptpclocktransdefaultdstable" }
    if yname == "cPtpClockPortTable" { return "Cptpclockporttable" }
    if yname == "cPtpClockPortDSTable" { return "Cptpclockportdstable" }
    if yname == "cPtpClockPortRunningTable" { return "Cptpclockportrunningtable" }
    if yname == "cPtpClockPortTransDSTable" { return "Cptpclockporttransdstable" }
    if yname == "cPtpClockPortAssociateTable" { return "Cptpclockportassociatetable" }
    return ""
}

func (cISCOPTPMIB *CISCOPTPMIB) GetSegmentPath() string {
    return "CISCO-PTP-MIB:CISCO-PTP-MIB"
}

func (cISCOPTPMIB *CISCOPTPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoPtpMIBSystemInfo" {
        return &cISCOPTPMIB.Ciscoptpmibsysteminfo
    }
    if childYangName == "cPtpSystemTable" {
        return &cISCOPTPMIB.Cptpsystemtable
    }
    if childYangName == "cPtpSystemDomainTable" {
        return &cISCOPTPMIB.Cptpsystemdomaintable
    }
    if childYangName == "cPtpClockNodeTable" {
        return &cISCOPTPMIB.Cptpclocknodetable
    }
    if childYangName == "cPtpClockCurrentDSTable" {
        return &cISCOPTPMIB.Cptpclockcurrentdstable
    }
    if childYangName == "cPtpClockParentDSTable" {
        return &cISCOPTPMIB.Cptpclockparentdstable
    }
    if childYangName == "cPtpClockDefaultDSTable" {
        return &cISCOPTPMIB.Cptpclockdefaultdstable
    }
    if childYangName == "cPtpClockRunningTable" {
        return &cISCOPTPMIB.Cptpclockrunningtable
    }
    if childYangName == "cPtpClockTimePropertiesDSTable" {
        return &cISCOPTPMIB.Cptpclocktimepropertiesdstable
    }
    if childYangName == "cPtpClockTransDefaultDSTable" {
        return &cISCOPTPMIB.Cptpclocktransdefaultdstable
    }
    if childYangName == "cPtpClockPortTable" {
        return &cISCOPTPMIB.Cptpclockporttable
    }
    if childYangName == "cPtpClockPortDSTable" {
        return &cISCOPTPMIB.Cptpclockportdstable
    }
    if childYangName == "cPtpClockPortRunningTable" {
        return &cISCOPTPMIB.Cptpclockportrunningtable
    }
    if childYangName == "cPtpClockPortTransDSTable" {
        return &cISCOPTPMIB.Cptpclockporttransdstable
    }
    if childYangName == "cPtpClockPortAssociateTable" {
        return &cISCOPTPMIB.Cptpclockportassociatetable
    }
    return nil
}

func (cISCOPTPMIB *CISCOPTPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoPtpMIBSystemInfo"] = &cISCOPTPMIB.Ciscoptpmibsysteminfo
    children["cPtpSystemTable"] = &cISCOPTPMIB.Cptpsystemtable
    children["cPtpSystemDomainTable"] = &cISCOPTPMIB.Cptpsystemdomaintable
    children["cPtpClockNodeTable"] = &cISCOPTPMIB.Cptpclocknodetable
    children["cPtpClockCurrentDSTable"] = &cISCOPTPMIB.Cptpclockcurrentdstable
    children["cPtpClockParentDSTable"] = &cISCOPTPMIB.Cptpclockparentdstable
    children["cPtpClockDefaultDSTable"] = &cISCOPTPMIB.Cptpclockdefaultdstable
    children["cPtpClockRunningTable"] = &cISCOPTPMIB.Cptpclockrunningtable
    children["cPtpClockTimePropertiesDSTable"] = &cISCOPTPMIB.Cptpclocktimepropertiesdstable
    children["cPtpClockTransDefaultDSTable"] = &cISCOPTPMIB.Cptpclocktransdefaultdstable
    children["cPtpClockPortTable"] = &cISCOPTPMIB.Cptpclockporttable
    children["cPtpClockPortDSTable"] = &cISCOPTPMIB.Cptpclockportdstable
    children["cPtpClockPortRunningTable"] = &cISCOPTPMIB.Cptpclockportrunningtable
    children["cPtpClockPortTransDSTable"] = &cISCOPTPMIB.Cptpclockporttransdstable
    children["cPtpClockPortAssociateTable"] = &cISCOPTPMIB.Cptpclockportassociatetable
    return children
}

func (cISCOPTPMIB *CISCOPTPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOPTPMIB *CISCOPTPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOPTPMIB *CISCOPTPMIB) GetYangName() string { return "CISCO-PTP-MIB" }

func (cISCOPTPMIB *CISCOPTPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOPTPMIB *CISCOPTPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOPTPMIB *CISCOPTPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOPTPMIB *CISCOPTPMIB) SetParent(parent types.Entity) { cISCOPTPMIB.parent = parent }

func (cISCOPTPMIB *CISCOPTPMIB) GetParent() types.Entity { return cISCOPTPMIB.parent }

func (cISCOPTPMIB *CISCOPTPMIB) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Ciscoptpmibsysteminfo
type CISCOPTPMIB_Ciscoptpmibsysteminfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object specifies the PTP Profile implemented on the system. The type
    // is ClockProfileType.
    Cptpsystemprofile interface{}
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetFilter() yfilter.YFilter { return ciscoptpmibsysteminfo.YFilter }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) SetFilter(yf yfilter.YFilter) { ciscoptpmibsysteminfo.YFilter = yf }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetGoName(yname string) string {
    if yname == "cPtpSystemProfile" { return "Cptpsystemprofile" }
    return ""
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetSegmentPath() string {
    return "ciscoPtpMIBSystemInfo"
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpSystemProfile"] = ciscoptpmibsysteminfo.Cptpsystemprofile
    return leafs
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetYangName() string { return "ciscoPtpMIBSystemInfo" }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) SetParent(parent types.Entity) { ciscoptpmibsysteminfo.parent = parent }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetParent() types.Entity { return ciscoptpmibsysteminfo.parent }

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpsystemtable
// Table of count information about the PTP system for all
// domains.
type CISCOPTPMIB_Cptpsystemtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing count information about a single domain.
    // New row entries are added when the PTP clock for this domain is configured,
    // while the unconfiguration of the PTP clock removes it. The type is slice of
    // CISCOPTPMIB_Cptpsystemtable_Cptpsystementry.
    Cptpsystementry []CISCOPTPMIB_Cptpsystemtable_Cptpsystementry
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetFilter() yfilter.YFilter { return cptpsystemtable.YFilter }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) SetFilter(yf yfilter.YFilter) { cptpsystemtable.YFilter = yf }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetGoName(yname string) string {
    if yname == "cPtpSystemEntry" { return "Cptpsystementry" }
    return ""
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetSegmentPath() string {
    return "cPtpSystemTable"
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpSystemEntry" {
        for _, c := range cptpsystemtable.Cptpsystementry {
            if cptpsystemtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpsystemtable_Cptpsystementry{}
        cptpsystemtable.Cptpsystementry = append(cptpsystemtable.Cptpsystementry, child)
        return &cptpsystemtable.Cptpsystementry[len(cptpsystemtable.Cptpsystementry)-1]
    }
    return nil
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpsystemtable.Cptpsystementry {
        children[cptpsystemtable.Cptpsystementry[i].GetSegmentPath()] = &cptpsystemtable.Cptpsystementry[i]
    }
    return children
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetYangName() string { return "cPtpSystemTable" }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) SetParent(parent types.Entity) { cptpsystemtable.parent = parent }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetParent() types.Entity { return cptpsystemtable.parent }

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpsystemtable_Cptpsystementry
// An entry in the table, containing count information about a
// single domain. New row entries are added when the PTP clock for
// this domain is configured, while the unconfiguration of the PTP
// clock removes it.
type CISCOPTPMIB_Cptpsystemtable_Cptpsystementry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The Clock Domain is a logical group of
    // clocks and devices that synchronize with each other using the PTP protocol.
    // 0           Default domain 1           Alternate domain 1 2          
    // Alternate domain 2 3           Alternate domain 3 4 - 127     User-defined
    // domains 128 - 255   Reserved. The type is interface{} with range: 0..255.
    Cptpdomainindex interface{}

    // This attribute is a key. This object specifies the instance of the Clock
    // for this domain. The type is interface{} with range: 0..255.
    Cptpinstanceindex interface{}

    // This object specifies the total number of clock ports configured within a
    // domain. The type is interface{} with range: 0..4294967295. Units are ptp
    // ports.
    Cptpdomainclockportstotal interface{}

    // This object specifies the total number of clock port Physical interfaces
    // configured within a domain instance for PTP communications. The type is
    // interface{} with range: 0..4294967295. Units are physical ports.
    Cptpdomainclockportphysicalinterfacestotal interface{}
}

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetFilter() yfilter.YFilter { return cptpsystementry.YFilter }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) SetFilter(yf yfilter.YFilter) { cptpsystementry.YFilter = yf }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetGoName(yname string) string {
    if yname == "cPtpDomainIndex" { return "Cptpdomainindex" }
    if yname == "cPtpInstanceIndex" { return "Cptpinstanceindex" }
    if yname == "cPtpDomainClockPortsTotal" { return "Cptpdomainclockportstotal" }
    if yname == "cPtpDomainClockPortPhysicalInterfacesTotal" { return "Cptpdomainclockportphysicalinterfacestotal" }
    return ""
}

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetSegmentPath() string {
    return "cPtpSystemEntry" + "[cPtpDomainIndex='" + fmt.Sprintf("%v", cptpsystementry.Cptpdomainindex) + "']" + "[cPtpInstanceIndex='" + fmt.Sprintf("%v", cptpsystementry.Cptpinstanceindex) + "']"
}

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpDomainIndex"] = cptpsystementry.Cptpdomainindex
    leafs["cPtpInstanceIndex"] = cptpsystementry.Cptpinstanceindex
    leafs["cPtpDomainClockPortsTotal"] = cptpsystementry.Cptpdomainclockportstotal
    leafs["cPtpDomainClockPortPhysicalInterfacesTotal"] = cptpsystementry.Cptpdomainclockportphysicalinterfacestotal
    return leafs
}

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetYangName() string { return "cPtpSystemEntry" }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) SetParent(parent types.Entity) { cptpsystementry.parent = parent }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetParent() types.Entity { return cptpsystementry.parent }

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetParentYangName() string { return "cPtpSystemTable" }

// CISCOPTPMIB_Cptpsystemdomaintable
// Table of information about the PTP system for all clock modes
// -- ordinary, boundary or transparent.
type CISCOPTPMIB_Cptpsystemdomaintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single clock mode for
    // the PTP system. A row entry gets added when PTP clocks are configured on
    // the router. The type is slice of
    // CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry.
    Cptpsystemdomainentry []CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetFilter() yfilter.YFilter { return cptpsystemdomaintable.YFilter }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) SetFilter(yf yfilter.YFilter) { cptpsystemdomaintable.YFilter = yf }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetGoName(yname string) string {
    if yname == "cPtpSystemDomainEntry" { return "Cptpsystemdomainentry" }
    return ""
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetSegmentPath() string {
    return "cPtpSystemDomainTable"
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpSystemDomainEntry" {
        for _, c := range cptpsystemdomaintable.Cptpsystemdomainentry {
            if cptpsystemdomaintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry{}
        cptpsystemdomaintable.Cptpsystemdomainentry = append(cptpsystemdomaintable.Cptpsystemdomainentry, child)
        return &cptpsystemdomaintable.Cptpsystemdomainentry[len(cptpsystemdomaintable.Cptpsystemdomainentry)-1]
    }
    return nil
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpsystemdomaintable.Cptpsystemdomainentry {
        children[cptpsystemdomaintable.Cptpsystemdomainentry[i].GetSegmentPath()] = &cptpsystemdomaintable.Cptpsystemdomainentry[i]
    }
    return children
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetYangName() string { return "cPtpSystemDomainTable" }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) SetParent(parent types.Entity) { cptpsystemdomaintable.parent = parent }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetParent() types.Entity { return cptpsystemdomaintable.parent }

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry
// An entry in the table, containing information about a single
// clock mode for the PTP system. A row entry gets added when PTP
// clocks are configured on the router.
type CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpsystemdomainclocktypeindex interface{}

    // This object specifies the  total number of PTP domains for this particular
    // clock type configured in this node. The type is interface{} with range:
    // 0..4294967295. Units are domains.
    Cptpsystemdomaintotals interface{}
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetFilter() yfilter.YFilter { return cptpsystemdomainentry.YFilter }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) SetFilter(yf yfilter.YFilter) { cptpsystemdomainentry.YFilter = yf }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetGoName(yname string) string {
    if yname == "cPtpSystemDomainClockTypeIndex" { return "Cptpsystemdomainclocktypeindex" }
    if yname == "cPtpSystemDomainTotals" { return "Cptpsystemdomaintotals" }
    return ""
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetSegmentPath() string {
    return "cPtpSystemDomainEntry" + "[cPtpSystemDomainClockTypeIndex='" + fmt.Sprintf("%v", cptpsystemdomainentry.Cptpsystemdomainclocktypeindex) + "']"
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpSystemDomainClockTypeIndex"] = cptpsystemdomainentry.Cptpsystemdomainclocktypeindex
    leafs["cPtpSystemDomainTotals"] = cptpsystemdomainentry.Cptpsystemdomaintotals
    return leafs
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetYangName() string { return "cPtpSystemDomainEntry" }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) SetParent(parent types.Entity) { cptpsystemdomainentry.parent = parent }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetParent() types.Entity { return cptpsystemdomainentry.parent }

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetParentYangName() string { return "cPtpSystemDomainTable" }

// CISCOPTPMIB_Cptpclocknodetable
// Table of information about the PTP system for a given domain.
type CISCOPTPMIB_Cptpclocknodetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single domain. A
    // entry is added when a new PTP clock domain is configured on the router. The
    // type is slice of CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry.
    Cptpclocknodeentry []CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetFilter() yfilter.YFilter { return cptpclocknodetable.YFilter }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) SetFilter(yf yfilter.YFilter) { cptpclocknodetable.YFilter = yf }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetGoName(yname string) string {
    if yname == "cPtpClockNodeEntry" { return "Cptpclocknodeentry" }
    return ""
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetSegmentPath() string {
    return "cPtpClockNodeTable"
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockNodeEntry" {
        for _, c := range cptpclocknodetable.Cptpclocknodeentry {
            if cptpclocknodetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry{}
        cptpclocknodetable.Cptpclocknodeentry = append(cptpclocknodetable.Cptpclocknodeentry, child)
        return &cptpclocknodetable.Cptpclocknodeentry[len(cptpclocknodetable.Cptpclocknodeentry)-1]
    }
    return nil
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclocknodetable.Cptpclocknodeentry {
        children[cptpclocknodetable.Cptpclocknodeentry[i].GetSegmentPath()] = &cptpclocknodetable.Cptpclocknodeentry[i]
    }
    return children
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetYangName() string { return "cPtpClockNodeTable" }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) SetParent(parent types.Entity) { cptpclocknodetable.parent = parent }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetParent() types.Entity { return cptpclocknodetable.parent }

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry
// An entry in the table, containing information about a single
// domain. A entry is added when a new PTP clock domain is
// configured on the router.
type CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the  domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the Clock
    // for this clock type for the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockinstanceindex interface{}

    // This object specifies whether the node is enabled for PTP input clocking
    // using the 1pps interface. The type is bool.
    Cptpclockinput1Ppsenabled interface{}

    // This object specifies whether enabled for Frequency input using the 1.544
    // Mhz, 2.048 Mhz, or 10Mhz timing interface. The type is bool.
    Cptpclockinputfrequencyenabled interface{}

    // This object specifies whether the node is enabled for TOD. The type is
    // bool.
    Cptpclocktodenabled interface{}

    // This object specifies whether the node is enabled for PTP input clocking
    // using the 1pps interface. The type is bool.
    Cptpclockoutput1Ppsenabled interface{}

    // This object specifies whether an offset is configured in order to
    // compensate for a known phase error such as network asymmetry. The type is
    // bool.
    Cptpclockoutput1Ppsoffsetenabled interface{}

    // This object specifies the fixed offset value configured to be added for the
    // 1pps output. The type is interface{} with range: 0..4294967295.
    Cptpclockoutput1Ppsoffsetvalue interface{}

    // This object specifies whether the added (fixed) offset to the 1pps output
    // is negative or not.  When object returns TRUE the offset is negative and
    // when object returns FALSE the offset is positive. The type is bool.
    Cptpclockoutput1Ppsoffsetnegative interface{}

    // This object specifies the 1pps interface used for PTP input clocking. The
    // type is string.
    Cptpclockinput1Ppsinterface interface{}

    // This object specifies the 1pps interface used for PTP output clocking. The
    // type is string.
    Cptpclockoutput1Ppsinterface interface{}

    // This object specifies the interface used for PTP TOD. The type is string.
    Cptpclocktodinterface interface{}
}

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetFilter() yfilter.YFilter { return cptpclocknodeentry.YFilter }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) SetFilter(yf yfilter.YFilter) { cptpclocknodeentry.YFilter = yf }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetGoName(yname string) string {
    if yname == "cPtpClockDomainIndex" { return "Cptpclockdomainindex" }
    if yname == "cPtpClockTypeIndex" { return "Cptpclocktypeindex" }
    if yname == "cPtpClockInstanceIndex" { return "Cptpclockinstanceindex" }
    if yname == "cPtpClockInput1ppsEnabled" { return "Cptpclockinput1Ppsenabled" }
    if yname == "cPtpClockInputFrequencyEnabled" { return "Cptpclockinputfrequencyenabled" }
    if yname == "cPtpClockTODEnabled" { return "Cptpclocktodenabled" }
    if yname == "cPtpClockOutput1ppsEnabled" { return "Cptpclockoutput1Ppsenabled" }
    if yname == "cPtpClockOutput1ppsOffsetEnabled" { return "Cptpclockoutput1Ppsoffsetenabled" }
    if yname == "cPtpClockOutput1ppsOffsetValue" { return "Cptpclockoutput1Ppsoffsetvalue" }
    if yname == "cPtpClockOutput1ppsOffsetNegative" { return "Cptpclockoutput1Ppsoffsetnegative" }
    if yname == "cPtpClockInput1ppsInterface" { return "Cptpclockinput1Ppsinterface" }
    if yname == "cPtpClockOutput1ppsInterface" { return "Cptpclockoutput1Ppsinterface" }
    if yname == "cPtpClockTODInterface" { return "Cptpclocktodinterface" }
    return ""
}

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetSegmentPath() string {
    return "cPtpClockNodeEntry" + "[cPtpClockDomainIndex='" + fmt.Sprintf("%v", cptpclocknodeentry.Cptpclockdomainindex) + "']" + "[cPtpClockTypeIndex='" + fmt.Sprintf("%v", cptpclocknodeentry.Cptpclocktypeindex) + "']" + "[cPtpClockInstanceIndex='" + fmt.Sprintf("%v", cptpclocknodeentry.Cptpclockinstanceindex) + "']"
}

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockDomainIndex"] = cptpclocknodeentry.Cptpclockdomainindex
    leafs["cPtpClockTypeIndex"] = cptpclocknodeentry.Cptpclocktypeindex
    leafs["cPtpClockInstanceIndex"] = cptpclocknodeentry.Cptpclockinstanceindex
    leafs["cPtpClockInput1ppsEnabled"] = cptpclocknodeentry.Cptpclockinput1Ppsenabled
    leafs["cPtpClockInputFrequencyEnabled"] = cptpclocknodeentry.Cptpclockinputfrequencyenabled
    leafs["cPtpClockTODEnabled"] = cptpclocknodeentry.Cptpclocktodenabled
    leafs["cPtpClockOutput1ppsEnabled"] = cptpclocknodeentry.Cptpclockoutput1Ppsenabled
    leafs["cPtpClockOutput1ppsOffsetEnabled"] = cptpclocknodeentry.Cptpclockoutput1Ppsoffsetenabled
    leafs["cPtpClockOutput1ppsOffsetValue"] = cptpclocknodeentry.Cptpclockoutput1Ppsoffsetvalue
    leafs["cPtpClockOutput1ppsOffsetNegative"] = cptpclocknodeentry.Cptpclockoutput1Ppsoffsetnegative
    leafs["cPtpClockInput1ppsInterface"] = cptpclocknodeentry.Cptpclockinput1Ppsinterface
    leafs["cPtpClockOutput1ppsInterface"] = cptpclocknodeentry.Cptpclockoutput1Ppsinterface
    leafs["cPtpClockTODInterface"] = cptpclocknodeentry.Cptpclocktodinterface
    return leafs
}

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetYangName() string { return "cPtpClockNodeEntry" }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) SetParent(parent types.Entity) { cptpclocknodeentry.parent = parent }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetParent() types.Entity { return cptpclocknodeentry.parent }

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetParentYangName() string { return "cPtpClockNodeTable" }

// CISCOPTPMIB_Cptpclockcurrentdstable
// Table of information about the PTP clock Current Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockcurrentdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Current Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry.
    Cptpclockcurrentdsentry []CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetFilter() yfilter.YFilter { return cptpclockcurrentdstable.YFilter }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) SetFilter(yf yfilter.YFilter) { cptpclockcurrentdstable.YFilter = yf }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetGoName(yname string) string {
    if yname == "cPtpClockCurrentDSEntry" { return "Cptpclockcurrentdsentry" }
    return ""
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetSegmentPath() string {
    return "cPtpClockCurrentDSTable"
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockCurrentDSEntry" {
        for _, c := range cptpclockcurrentdstable.Cptpclockcurrentdsentry {
            if cptpclockcurrentdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry{}
        cptpclockcurrentdstable.Cptpclockcurrentdsentry = append(cptpclockcurrentdstable.Cptpclockcurrentdsentry, child)
        return &cptpclockcurrentdstable.Cptpclockcurrentdsentry[len(cptpclockcurrentdstable.Cptpclockcurrentdsentry)-1]
    }
    return nil
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockcurrentdstable.Cptpclockcurrentdsentry {
        children[cptpclockcurrentdstable.Cptpclockcurrentdsentry[i].GetSegmentPath()] = &cptpclockcurrentdstable.Cptpclockcurrentdsentry[i]
    }
    return children
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetYangName() string { return "cPtpClockCurrentDSTable" }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) SetParent(parent types.Entity) { cptpclockcurrentdstable.parent = parent }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetParent() types.Entity { return cptpclockcurrentdstable.parent }

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry
// An entry in the table, containing information about a single
// PTP clock Current Datasets for a domain.
type CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockcurrentdsdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockcurrentdsclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockcurrentdsinstanceindex interface{}

    // The current clock dataset StepsRemoved value.  This object specifies the
    // distance measured by the number of Boundary clocks between the local clock
    // and the Foreign master as indicated in the stepsRemoved field of Announce
    // messages. The type is interface{} with range: 0..4294967295. Units are
    // steps.
    Cptpclockcurrentdsstepsremoved interface{}

    // This object specifies the current clock dataset ClockOffset value. The
    // value of the computation of the offset in time between a slave and a master
    // clock. The type is string with length: 1..255. Units are Time Interval.
    Cptpclockcurrentdsoffsetfrommaster interface{}

    // This object specifies the current clock dataset MeanPathDelay value.  The
    // mean path delay between a pair of ports as measure by the delay
    // request-response mechanism. The type is string with length: 1..255.
    Cptpclockcurrentdsmeanpathdelay interface{}
}

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetFilter() yfilter.YFilter { return cptpclockcurrentdsentry.YFilter }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) SetFilter(yf yfilter.YFilter) { cptpclockcurrentdsentry.YFilter = yf }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockCurrentDSDomainIndex" { return "Cptpclockcurrentdsdomainindex" }
    if yname == "cPtpClockCurrentDSClockTypeIndex" { return "Cptpclockcurrentdsclocktypeindex" }
    if yname == "cPtpClockCurrentDSInstanceIndex" { return "Cptpclockcurrentdsinstanceindex" }
    if yname == "cPtpClockCurrentDSStepsRemoved" { return "Cptpclockcurrentdsstepsremoved" }
    if yname == "cPtpClockCurrentDSOffsetFromMaster" { return "Cptpclockcurrentdsoffsetfrommaster" }
    if yname == "cPtpClockCurrentDSMeanPathDelay" { return "Cptpclockcurrentdsmeanpathdelay" }
    return ""
}

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetSegmentPath() string {
    return "cPtpClockCurrentDSEntry" + "[cPtpClockCurrentDSDomainIndex='" + fmt.Sprintf("%v", cptpclockcurrentdsentry.Cptpclockcurrentdsdomainindex) + "']" + "[cPtpClockCurrentDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockcurrentdsentry.Cptpclockcurrentdsclocktypeindex) + "']" + "[cPtpClockCurrentDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockcurrentdsentry.Cptpclockcurrentdsinstanceindex) + "']"
}

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockCurrentDSDomainIndex"] = cptpclockcurrentdsentry.Cptpclockcurrentdsdomainindex
    leafs["cPtpClockCurrentDSClockTypeIndex"] = cptpclockcurrentdsentry.Cptpclockcurrentdsclocktypeindex
    leafs["cPtpClockCurrentDSInstanceIndex"] = cptpclockcurrentdsentry.Cptpclockcurrentdsinstanceindex
    leafs["cPtpClockCurrentDSStepsRemoved"] = cptpclockcurrentdsentry.Cptpclockcurrentdsstepsremoved
    leafs["cPtpClockCurrentDSOffsetFromMaster"] = cptpclockcurrentdsentry.Cptpclockcurrentdsoffsetfrommaster
    leafs["cPtpClockCurrentDSMeanPathDelay"] = cptpclockcurrentdsentry.Cptpclockcurrentdsmeanpathdelay
    return leafs
}

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetYangName() string { return "cPtpClockCurrentDSEntry" }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) SetParent(parent types.Entity) { cptpclockcurrentdsentry.parent = parent }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetParent() types.Entity { return cptpclockcurrentdsentry.parent }

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetParentYangName() string { return "cPtpClockCurrentDSTable" }

// CISCOPTPMIB_Cptpclockparentdstable
// Table of information about the PTP clock Parent Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockparentdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Parent Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry.
    Cptpclockparentdsentry []CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetFilter() yfilter.YFilter { return cptpclockparentdstable.YFilter }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) SetFilter(yf yfilter.YFilter) { cptpclockparentdstable.YFilter = yf }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetGoName(yname string) string {
    if yname == "cPtpClockParentDSEntry" { return "Cptpclockparentdsentry" }
    return ""
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetSegmentPath() string {
    return "cPtpClockParentDSTable"
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockParentDSEntry" {
        for _, c := range cptpclockparentdstable.Cptpclockparentdsentry {
            if cptpclockparentdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry{}
        cptpclockparentdstable.Cptpclockparentdsentry = append(cptpclockparentdstable.Cptpclockparentdsentry, child)
        return &cptpclockparentdstable.Cptpclockparentdsentry[len(cptpclockparentdstable.Cptpclockparentdsentry)-1]
    }
    return nil
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockparentdstable.Cptpclockparentdsentry {
        children[cptpclockparentdstable.Cptpclockparentdsentry[i].GetSegmentPath()] = &cptpclockparentdstable.Cptpclockparentdsentry[i]
    }
    return children
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetYangName() string { return "cPtpClockParentDSTable" }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) SetParent(parent types.Entity) { cptpclockparentdstable.parent = parent }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetParent() types.Entity { return cptpclockparentdstable.parent }

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry
// An entry in the table, containing information about a single
// PTP clock Parent Datasets for a domain.
type CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockparentdsdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockparentdsclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockparentdsinstanceindex interface{}

    // This object specifies the value of portIdentity of the port on the master
    // that issues the Sync messages used in synchronizing this clock. The type is
    // string.
    Cptpclockparentdsparentportidentity interface{}

    // This object specifies the Parent Dataset ParentStats value.  This value
    // indicates whether the values of ParentDSOffset and ParentDSClockPhChRate
    // have been measured and are valid. A TRUE value shall indicate valid data.
    // The type is bool.
    Cptpclockparentdsparentstats interface{}

    // This object specifies the Parent Dataset ParentOffsetScaledLogVariance
    // value.  This value is the variance of the parent clocks phase as measured
    // by the local clock. The type is interface{} with range: -128..127.
    Cptpclockparentdsoffset interface{}

    // This object specifies the clock's parent dataset ParentClockPhaseChangeRate
    // value.  This value is an estimate of the parent clocks phase change rate as
    // measured by the slave clock. The type is interface{} with range:
    // -2147483648..2147483647.
    Cptpclockparentdsclockphchrate interface{}

    // This object specifies the parent dataset Grandmaster clock identity. The
    // type is string with length: 1..255.
    Cptpclockparentdsgmclockidentity interface{}

    // This object specifies the parent dataset Grandmaster clock priority1. The
    // type is interface{} with range: -2147483648..2147483647.
    Cptpclockparentdsgmclockpriority1 interface{}

    // This object specifies the parent dataset grandmaster clock priority2. The
    // type is interface{} with range: -2147483648..2147483647.
    Cptpclockparentdsgmclockpriority2 interface{}

    // This object specifies the parent dataset grandmaster clock quality class.
    // The type is interface{} with range: 0..255.
    Cptpclockparentdsgmclockqualityclass interface{}

    // This object specifies the parent dataset grandmaster clock quality
    // accuracy. The type is ClockQualityAccuracyType.
    Cptpclockparentdsgmclockqualityaccuracy interface{}

    // This object specifies the parent dataset grandmaster clock quality offset.
    // The type is interface{} with range: 0..4294967295.
    Cptpclockparentdsgmclockqualityoffset interface{}
}

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetFilter() yfilter.YFilter { return cptpclockparentdsentry.YFilter }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) SetFilter(yf yfilter.YFilter) { cptpclockparentdsentry.YFilter = yf }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockParentDSDomainIndex" { return "Cptpclockparentdsdomainindex" }
    if yname == "cPtpClockParentDSClockTypeIndex" { return "Cptpclockparentdsclocktypeindex" }
    if yname == "cPtpClockParentDSInstanceIndex" { return "Cptpclockparentdsinstanceindex" }
    if yname == "cPtpClockParentDSParentPortIdentity" { return "Cptpclockparentdsparentportidentity" }
    if yname == "cPtpClockParentDSParentStats" { return "Cptpclockparentdsparentstats" }
    if yname == "cPtpClockParentDSOffset" { return "Cptpclockparentdsoffset" }
    if yname == "cPtpClockParentDSClockPhChRate" { return "Cptpclockparentdsclockphchrate" }
    if yname == "cPtpClockParentDSGMClockIdentity" { return "Cptpclockparentdsgmclockidentity" }
    if yname == "cPtpClockParentDSGMClockPriority1" { return "Cptpclockparentdsgmclockpriority1" }
    if yname == "cPtpClockParentDSGMClockPriority2" { return "Cptpclockparentdsgmclockpriority2" }
    if yname == "cPtpClockParentDSGMClockQualityClass" { return "Cptpclockparentdsgmclockqualityclass" }
    if yname == "cPtpClockParentDSGMClockQualityAccuracy" { return "Cptpclockparentdsgmclockqualityaccuracy" }
    if yname == "cPtpClockParentDSGMClockQualityOffset" { return "Cptpclockparentdsgmclockqualityoffset" }
    return ""
}

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetSegmentPath() string {
    return "cPtpClockParentDSEntry" + "[cPtpClockParentDSDomainIndex='" + fmt.Sprintf("%v", cptpclockparentdsentry.Cptpclockparentdsdomainindex) + "']" + "[cPtpClockParentDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockparentdsentry.Cptpclockparentdsclocktypeindex) + "']" + "[cPtpClockParentDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockparentdsentry.Cptpclockparentdsinstanceindex) + "']"
}

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockParentDSDomainIndex"] = cptpclockparentdsentry.Cptpclockparentdsdomainindex
    leafs["cPtpClockParentDSClockTypeIndex"] = cptpclockparentdsentry.Cptpclockparentdsclocktypeindex
    leafs["cPtpClockParentDSInstanceIndex"] = cptpclockparentdsentry.Cptpclockparentdsinstanceindex
    leafs["cPtpClockParentDSParentPortIdentity"] = cptpclockparentdsentry.Cptpclockparentdsparentportidentity
    leafs["cPtpClockParentDSParentStats"] = cptpclockparentdsentry.Cptpclockparentdsparentstats
    leafs["cPtpClockParentDSOffset"] = cptpclockparentdsentry.Cptpclockparentdsoffset
    leafs["cPtpClockParentDSClockPhChRate"] = cptpclockparentdsentry.Cptpclockparentdsclockphchrate
    leafs["cPtpClockParentDSGMClockIdentity"] = cptpclockparentdsentry.Cptpclockparentdsgmclockidentity
    leafs["cPtpClockParentDSGMClockPriority1"] = cptpclockparentdsentry.Cptpclockparentdsgmclockpriority1
    leafs["cPtpClockParentDSGMClockPriority2"] = cptpclockparentdsentry.Cptpclockparentdsgmclockpriority2
    leafs["cPtpClockParentDSGMClockQualityClass"] = cptpclockparentdsentry.Cptpclockparentdsgmclockqualityclass
    leafs["cPtpClockParentDSGMClockQualityAccuracy"] = cptpclockparentdsentry.Cptpclockparentdsgmclockqualityaccuracy
    leafs["cPtpClockParentDSGMClockQualityOffset"] = cptpclockparentdsentry.Cptpclockparentdsgmclockqualityoffset
    return leafs
}

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetYangName() string { return "cPtpClockParentDSEntry" }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) SetParent(parent types.Entity) { cptpclockparentdsentry.parent = parent }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetParent() types.Entity { return cptpclockparentdsentry.parent }

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetParentYangName() string { return "cPtpClockParentDSTable" }

// CISCOPTPMIB_Cptpclockdefaultdstable
// Table of information about the PTP clock Default Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockdefaultdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Default Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry.
    Cptpclockdefaultdsentry []CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetFilter() yfilter.YFilter { return cptpclockdefaultdstable.YFilter }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) SetFilter(yf yfilter.YFilter) { cptpclockdefaultdstable.YFilter = yf }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetGoName(yname string) string {
    if yname == "cPtpClockDefaultDSEntry" { return "Cptpclockdefaultdsentry" }
    return ""
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetSegmentPath() string {
    return "cPtpClockDefaultDSTable"
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockDefaultDSEntry" {
        for _, c := range cptpclockdefaultdstable.Cptpclockdefaultdsentry {
            if cptpclockdefaultdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry{}
        cptpclockdefaultdstable.Cptpclockdefaultdsentry = append(cptpclockdefaultdstable.Cptpclockdefaultdsentry, child)
        return &cptpclockdefaultdstable.Cptpclockdefaultdsentry[len(cptpclockdefaultdstable.Cptpclockdefaultdsentry)-1]
    }
    return nil
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockdefaultdstable.Cptpclockdefaultdsentry {
        children[cptpclockdefaultdstable.Cptpclockdefaultdsentry[i].GetSegmentPath()] = &cptpclockdefaultdstable.Cptpclockdefaultdsentry[i]
    }
    return children
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetYangName() string { return "cPtpClockDefaultDSTable" }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) SetParent(parent types.Entity) { cptpclockdefaultdstable.parent = parent }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetParent() types.Entity { return cptpclockdefaultdstable.parent }

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry
// An entry in the table, containing information about a single
// PTP clock Default Datasets for a domain.
type CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockdefaultdsdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockdefaultdsclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockdefaultdsinstanceindex interface{}

    // This object specifies whether the Two Step process is used. The type is
    // bool.
    Cptpclockdefaultdstwostepflag interface{}

    // This object specifies the default Datasets clock identity. The type is
    // string with length: 1..255.
    Cptpclockdefaultdsclockidentity interface{}

    // This object specifies the default Datasets clock Priority1. The type is
    // interface{} with range: -2147483648..2147483647.
    Cptpclockdefaultdspriority1 interface{}

    // This object specifies the default Datasets clock Priority2. The type is
    // interface{} with range: -2147483648..2147483647.
    Cptpclockdefaultdspriority2 interface{}

    // Whether the SlaveOnly flag is set. The type is bool.
    Cptpclockdefaultdsslaveonly interface{}

    // This object specifies the default dataset Quality Class. The type is
    // interface{} with range: 0..255.
    Cptpclockdefaultdsqualityclass interface{}

    // This object specifies the default dataset Quality Accurarcy. The type is
    // ClockQualityAccuracyType.
    Cptpclockdefaultdsqualityaccuracy interface{}

    // This object specifies the default dataset Quality offset. The type is
    // interface{} with range: -2147483648..2147483647.
    Cptpclockdefaultdsqualityoffset interface{}
}

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetFilter() yfilter.YFilter { return cptpclockdefaultdsentry.YFilter }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) SetFilter(yf yfilter.YFilter) { cptpclockdefaultdsentry.YFilter = yf }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockDefaultDSDomainIndex" { return "Cptpclockdefaultdsdomainindex" }
    if yname == "cPtpClockDefaultDSClockTypeIndex" { return "Cptpclockdefaultdsclocktypeindex" }
    if yname == "cPtpClockDefaultDSInstanceIndex" { return "Cptpclockdefaultdsinstanceindex" }
    if yname == "cPtpClockDefaultDSTwoStepFlag" { return "Cptpclockdefaultdstwostepflag" }
    if yname == "cPtpClockDefaultDSClockIdentity" { return "Cptpclockdefaultdsclockidentity" }
    if yname == "cPtpClockDefaultDSPriority1" { return "Cptpclockdefaultdspriority1" }
    if yname == "cPtpClockDefaultDSPriority2" { return "Cptpclockdefaultdspriority2" }
    if yname == "cPtpClockDefaultDSSlaveOnly" { return "Cptpclockdefaultdsslaveonly" }
    if yname == "cPtpClockDefaultDSQualityClass" { return "Cptpclockdefaultdsqualityclass" }
    if yname == "cPtpClockDefaultDSQualityAccuracy" { return "Cptpclockdefaultdsqualityaccuracy" }
    if yname == "cPtpClockDefaultDSQualityOffset" { return "Cptpclockdefaultdsqualityoffset" }
    return ""
}

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetSegmentPath() string {
    return "cPtpClockDefaultDSEntry" + "[cPtpClockDefaultDSDomainIndex='" + fmt.Sprintf("%v", cptpclockdefaultdsentry.Cptpclockdefaultdsdomainindex) + "']" + "[cPtpClockDefaultDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockdefaultdsentry.Cptpclockdefaultdsclocktypeindex) + "']" + "[cPtpClockDefaultDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockdefaultdsentry.Cptpclockdefaultdsinstanceindex) + "']"
}

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockDefaultDSDomainIndex"] = cptpclockdefaultdsentry.Cptpclockdefaultdsdomainindex
    leafs["cPtpClockDefaultDSClockTypeIndex"] = cptpclockdefaultdsentry.Cptpclockdefaultdsclocktypeindex
    leafs["cPtpClockDefaultDSInstanceIndex"] = cptpclockdefaultdsentry.Cptpclockdefaultdsinstanceindex
    leafs["cPtpClockDefaultDSTwoStepFlag"] = cptpclockdefaultdsentry.Cptpclockdefaultdstwostepflag
    leafs["cPtpClockDefaultDSClockIdentity"] = cptpclockdefaultdsentry.Cptpclockdefaultdsclockidentity
    leafs["cPtpClockDefaultDSPriority1"] = cptpclockdefaultdsentry.Cptpclockdefaultdspriority1
    leafs["cPtpClockDefaultDSPriority2"] = cptpclockdefaultdsentry.Cptpclockdefaultdspriority2
    leafs["cPtpClockDefaultDSSlaveOnly"] = cptpclockdefaultdsentry.Cptpclockdefaultdsslaveonly
    leafs["cPtpClockDefaultDSQualityClass"] = cptpclockdefaultdsentry.Cptpclockdefaultdsqualityclass
    leafs["cPtpClockDefaultDSQualityAccuracy"] = cptpclockdefaultdsentry.Cptpclockdefaultdsqualityaccuracy
    leafs["cPtpClockDefaultDSQualityOffset"] = cptpclockdefaultdsentry.Cptpclockdefaultdsqualityoffset
    return leafs
}

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetYangName() string { return "cPtpClockDefaultDSEntry" }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) SetParent(parent types.Entity) { cptpclockdefaultdsentry.parent = parent }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetParent() types.Entity { return cptpclockdefaultdsentry.parent }

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetParentYangName() string { return "cPtpClockDefaultDSTable" }

// CISCOPTPMIB_Cptpclockrunningtable
// Table of information about the PTP clock Running Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockrunningtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // running Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry.
    Cptpclockrunningentry []CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetFilter() yfilter.YFilter { return cptpclockrunningtable.YFilter }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) SetFilter(yf yfilter.YFilter) { cptpclockrunningtable.YFilter = yf }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetGoName(yname string) string {
    if yname == "cPtpClockRunningEntry" { return "Cptpclockrunningentry" }
    return ""
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetSegmentPath() string {
    return "cPtpClockRunningTable"
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockRunningEntry" {
        for _, c := range cptpclockrunningtable.Cptpclockrunningentry {
            if cptpclockrunningtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry{}
        cptpclockrunningtable.Cptpclockrunningentry = append(cptpclockrunningtable.Cptpclockrunningentry, child)
        return &cptpclockrunningtable.Cptpclockrunningentry[len(cptpclockrunningtable.Cptpclockrunningentry)-1]
    }
    return nil
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockrunningtable.Cptpclockrunningentry {
        children[cptpclockrunningtable.Cptpclockrunningentry[i].GetSegmentPath()] = &cptpclockrunningtable.Cptpclockrunningentry[i]
    }
    return children
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetYangName() string { return "cPtpClockRunningTable" }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) SetParent(parent types.Entity) { cptpclockrunningtable.parent = parent }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetParent() types.Entity { return cptpclockrunningtable.parent }

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry
// An entry in the table, containing information about a single
// PTP clock running Datasets for a domain.
type CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockrunningdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockrunningclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockrunninginstanceindex interface{}

    // This object specifies the Clock state returned by PTP engine which was
    // described earlier.  Freerun state. Applies to a slave device that is not
    // locked to a master. This is the initial state a slave starts out with when
    // it is not getting any PTP packets from the master or because of some other
    // input error (erroneous packets, etc).  Holdover state. In this state the
    // slave device is locked to a master but communication with the master is
    // lost or the timestamps in the ptp packets are incorrect. But since the
    // slave was locked to the master, it can run with the same accuracy for
    // sometime. The slave can continue to operate in this state for some time. If
    // communication with the master is not restored for a while, the device is
    // moved to the FREERUN state.  Acquiring state. The slave device is receiving
    // packets from a master and is trying to acquire a lock.  Freq_locked state.
    // Slave device is locked to the Master with respect to frequency, but not
    // phase aligned  Phase_aligned state. Locked to the master with respect to
    // frequency and phase. The type is ClockStateType.
    Cptpclockrunningstate interface{}

    // This object specifies the total number of all packet Unicast and multicast
    // that have been sent out for this clock in this domain for this type. The
    // type is interface{} with range: 0..18446744073709551615.
    Cptpclockrunningpacketssent interface{}

    // This object specifies the total number of all packet Unicast and multicast
    // that have been received for this clock in this domain for this type. The
    // type is interface{} with range: 0..18446744073709551615.
    Cptpclockrunningpacketsreceived interface{}
}

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetFilter() yfilter.YFilter { return cptpclockrunningentry.YFilter }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) SetFilter(yf yfilter.YFilter) { cptpclockrunningentry.YFilter = yf }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetGoName(yname string) string {
    if yname == "cPtpClockRunningDomainIndex" { return "Cptpclockrunningdomainindex" }
    if yname == "cPtpClockRunningClockTypeIndex" { return "Cptpclockrunningclocktypeindex" }
    if yname == "cPtpClockRunningInstanceIndex" { return "Cptpclockrunninginstanceindex" }
    if yname == "cPtpClockRunningState" { return "Cptpclockrunningstate" }
    if yname == "cPtpClockRunningPacketsSent" { return "Cptpclockrunningpacketssent" }
    if yname == "cPtpClockRunningPacketsReceived" { return "Cptpclockrunningpacketsreceived" }
    return ""
}

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetSegmentPath() string {
    return "cPtpClockRunningEntry" + "[cPtpClockRunningDomainIndex='" + fmt.Sprintf("%v", cptpclockrunningentry.Cptpclockrunningdomainindex) + "']" + "[cPtpClockRunningClockTypeIndex='" + fmt.Sprintf("%v", cptpclockrunningentry.Cptpclockrunningclocktypeindex) + "']" + "[cPtpClockRunningInstanceIndex='" + fmt.Sprintf("%v", cptpclockrunningentry.Cptpclockrunninginstanceindex) + "']"
}

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockRunningDomainIndex"] = cptpclockrunningentry.Cptpclockrunningdomainindex
    leafs["cPtpClockRunningClockTypeIndex"] = cptpclockrunningentry.Cptpclockrunningclocktypeindex
    leafs["cPtpClockRunningInstanceIndex"] = cptpclockrunningentry.Cptpclockrunninginstanceindex
    leafs["cPtpClockRunningState"] = cptpclockrunningentry.Cptpclockrunningstate
    leafs["cPtpClockRunningPacketsSent"] = cptpclockrunningentry.Cptpclockrunningpacketssent
    leafs["cPtpClockRunningPacketsReceived"] = cptpclockrunningentry.Cptpclockrunningpacketsreceived
    return leafs
}

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetYangName() string { return "cPtpClockRunningEntry" }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) SetParent(parent types.Entity) { cptpclockrunningentry.parent = parent }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetParent() types.Entity { return cptpclockrunningentry.parent }

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetParentYangName() string { return "cPtpClockRunningTable" }

// CISCOPTPMIB_Cptpclocktimepropertiesdstable
// Table of information about the PTP clock Timeproperties
// Datasets for all domains.
type CISCOPTPMIB_Cptpclocktimepropertiesdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // timeproperties Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry.
    Cptpclocktimepropertiesdsentry []CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetFilter() yfilter.YFilter { return cptpclocktimepropertiesdstable.YFilter }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) SetFilter(yf yfilter.YFilter) { cptpclocktimepropertiesdstable.YFilter = yf }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetGoName(yname string) string {
    if yname == "cPtpClockTimePropertiesDSEntry" { return "Cptpclocktimepropertiesdsentry" }
    return ""
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetSegmentPath() string {
    return "cPtpClockTimePropertiesDSTable"
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockTimePropertiesDSEntry" {
        for _, c := range cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry {
            if cptpclocktimepropertiesdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry{}
        cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry = append(cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry, child)
        return &cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry[len(cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry)-1]
    }
    return nil
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry {
        children[cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry[i].GetSegmentPath()] = &cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry[i]
    }
    return children
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetYangName() string { return "cPtpClockTimePropertiesDSTable" }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) SetParent(parent types.Entity) { cptpclocktimepropertiesdstable.parent = parent }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetParent() types.Entity { return cptpclocktimepropertiesdstable.parent }

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry
// An entry in the table, containing information about a single
// PTP clock timeproperties Datasets for a domain.
type CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclocktimepropertiesdsdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclocktimepropertiesdsclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclocktimepropertiesdsinstanceindex interface{}

    // This object specifies the timeproperties dataset value of whether current
    // UTC offset is valid. The type is bool.
    Cptpclocktimepropertiesdscurrentutcoffsetvalid interface{}

    // This object specifies the timeproperties dataset value of current UTC
    // offset.  In PTP systems whose epoch is the PTP epoch, the value of
    // timePropertiesDS.currentUtcOffset is the offset between TAI and UTC;
    // otherwise the value has no meaning. The value shall be in units of seconds.
    // The initialization value shall be selected as follows: a) If the
    // timePropertiesDS.ptpTimescale (see 8.2.4.8) is TRUE, the value is the value
    // obtained from a primary reference if the value is known at the time of
    // initialization, else. b) The value shall be the current number of leap
    // seconds (7.2.3) when the node is designed. The type is interface{} with
    // range: -2147483648..2147483647.
    Cptpclocktimepropertiesdscurrentutcoffset interface{}

    // This object specifies the Leap59 value in the clock Current Dataset. The
    // type is bool.
    Cptpclocktimepropertiesdsleap59 interface{}

    // This object specifies the Leap61 value in the clock Current Dataset. The
    // type is bool.
    Cptpclocktimepropertiesdsleap61 interface{}

    // This object specifies the Timetraceable value in the clock Current Dataset.
    // The type is bool.
    Cptpclocktimepropertiesdstimetraceable interface{}

    // This object specifies the Frequency Traceable value in the clock Current
    // Dataset. The type is bool.
    Cptpclocktimepropertiesdsfreqtraceable interface{}

    // This object specifies the PTP Timescale value in the clock Current Dataset.
    // The type is bool.
    Cptpclocktimepropertiesdsptptimescale interface{}

    // This object specifies the Timesource value in the clock Current Dataset.
    // The type is ClockTimeSourceType.
    Cptpclocktimepropertiesdssource interface{}
}

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetFilter() yfilter.YFilter { return cptpclocktimepropertiesdsentry.YFilter }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) SetFilter(yf yfilter.YFilter) { cptpclocktimepropertiesdsentry.YFilter = yf }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockTimePropertiesDSDomainIndex" { return "Cptpclocktimepropertiesdsdomainindex" }
    if yname == "cPtpClockTimePropertiesDSClockTypeIndex" { return "Cptpclocktimepropertiesdsclocktypeindex" }
    if yname == "cPtpClockTimePropertiesDSInstanceIndex" { return "Cptpclocktimepropertiesdsinstanceindex" }
    if yname == "cPtpClockTimePropertiesDSCurrentUTCOffsetValid" { return "Cptpclocktimepropertiesdscurrentutcoffsetvalid" }
    if yname == "cPtpClockTimePropertiesDSCurrentUTCOffset" { return "Cptpclocktimepropertiesdscurrentutcoffset" }
    if yname == "cPtpClockTimePropertiesDSLeap59" { return "Cptpclocktimepropertiesdsleap59" }
    if yname == "cPtpClockTimePropertiesDSLeap61" { return "Cptpclocktimepropertiesdsleap61" }
    if yname == "cPtpClockTimePropertiesDSTimeTraceable" { return "Cptpclocktimepropertiesdstimetraceable" }
    if yname == "cPtpClockTimePropertiesDSFreqTraceable" { return "Cptpclocktimepropertiesdsfreqtraceable" }
    if yname == "cPtpClockTimePropertiesDSPTPTimescale" { return "Cptpclocktimepropertiesdsptptimescale" }
    if yname == "cPtpClockTimePropertiesDSSource" { return "Cptpclocktimepropertiesdssource" }
    return ""
}

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetSegmentPath() string {
    return "cPtpClockTimePropertiesDSEntry" + "[cPtpClockTimePropertiesDSDomainIndex='" + fmt.Sprintf("%v", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsdomainindex) + "']" + "[cPtpClockTimePropertiesDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsclocktypeindex) + "']" + "[cPtpClockTimePropertiesDSInstanceIndex='" + fmt.Sprintf("%v", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsinstanceindex) + "']"
}

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockTimePropertiesDSDomainIndex"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsdomainindex
    leafs["cPtpClockTimePropertiesDSClockTypeIndex"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsclocktypeindex
    leafs["cPtpClockTimePropertiesDSInstanceIndex"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsinstanceindex
    leafs["cPtpClockTimePropertiesDSCurrentUTCOffsetValid"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdscurrentutcoffsetvalid
    leafs["cPtpClockTimePropertiesDSCurrentUTCOffset"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdscurrentutcoffset
    leafs["cPtpClockTimePropertiesDSLeap59"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsleap59
    leafs["cPtpClockTimePropertiesDSLeap61"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsleap61
    leafs["cPtpClockTimePropertiesDSTimeTraceable"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdstimetraceable
    leafs["cPtpClockTimePropertiesDSFreqTraceable"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsfreqtraceable
    leafs["cPtpClockTimePropertiesDSPTPTimescale"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsptptimescale
    leafs["cPtpClockTimePropertiesDSSource"] = cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdssource
    return leafs
}

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetYangName() string { return "cPtpClockTimePropertiesDSEntry" }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) SetParent(parent types.Entity) { cptpclocktimepropertiesdsentry.parent = parent }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetParent() types.Entity { return cptpclocktimepropertiesdsentry.parent }

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetParentYangName() string { return "cPtpClockTimePropertiesDSTable" }

// CISCOPTPMIB_Cptpclocktransdefaultdstable
// Table of information about the PTP Transparent clock Default
// Datasets for all domains.
type CISCOPTPMIB_Cptpclocktransdefaultdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP
    // Transparent clock Default Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry.
    Cptpclocktransdefaultdsentry []CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetFilter() yfilter.YFilter { return cptpclocktransdefaultdstable.YFilter }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) SetFilter(yf yfilter.YFilter) { cptpclocktransdefaultdstable.YFilter = yf }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetGoName(yname string) string {
    if yname == "cPtpClockTransDefaultDSEntry" { return "Cptpclocktransdefaultdsentry" }
    return ""
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetSegmentPath() string {
    return "cPtpClockTransDefaultDSTable"
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockTransDefaultDSEntry" {
        for _, c := range cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry {
            if cptpclocktransdefaultdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry{}
        cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry = append(cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry, child)
        return &cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry[len(cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry)-1]
    }
    return nil
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry {
        children[cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry[i].GetSegmentPath()] = &cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry[i]
    }
    return children
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetYangName() string { return "cPtpClockTransDefaultDSTable" }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) SetParent(parent types.Entity) { cptpclocktransdefaultdstable.parent = parent }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetParent() types.Entity { return cptpclocktransdefaultdstable.parent }

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry
// An entry in the table, containing information about a single
// PTP Transparent clock Default Datasets for a domain.
type CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclocktransdefaultdsdomainindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclocktransdefaultdsinstanceindex interface{}

    // This object specifies the value of the clockIdentity attribute of the local
    // clock. The type is string with length: 1..255.
    Cptpclocktransdefaultdsclockidentity interface{}

    // This object specifies the number of PTP ports of the device. The type is
    // interface{} with range: 0..4294967295.
    Cptpclocktransdefaultdsnumofports interface{}

    // This object, if the transparent clock is an end-to-end transparent clock,
    // has the value shall be E2E; If the transparent clock is a peer-to-peer
    // transparent clock, the value shall be P2P. The type is ClockMechanismType.
    Cptpclocktransdefaultdsdelay interface{}

    // This object specifies the value of the primary syntonization domain. The
    // initialization value shall be 0. The type is interface{} with range:
    // -2147483648..2147483647.
    Cptpclocktransdefaultdsprimarydomain interface{}
}

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetFilter() yfilter.YFilter { return cptpclocktransdefaultdsentry.YFilter }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) SetFilter(yf yfilter.YFilter) { cptpclocktransdefaultdsentry.YFilter = yf }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockTransDefaultDSDomainIndex" { return "Cptpclocktransdefaultdsdomainindex" }
    if yname == "cPtpClockTransDefaultDSInstanceIndex" { return "Cptpclocktransdefaultdsinstanceindex" }
    if yname == "cPtpClockTransDefaultDSClockIdentity" { return "Cptpclocktransdefaultdsclockidentity" }
    if yname == "cPtpClockTransDefaultDSNumOfPorts" { return "Cptpclocktransdefaultdsnumofports" }
    if yname == "cPtpClockTransDefaultDSDelay" { return "Cptpclocktransdefaultdsdelay" }
    if yname == "cPtpClockTransDefaultDSPrimaryDomain" { return "Cptpclocktransdefaultdsprimarydomain" }
    return ""
}

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetSegmentPath() string {
    return "cPtpClockTransDefaultDSEntry" + "[cPtpClockTransDefaultDSDomainIndex='" + fmt.Sprintf("%v", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsdomainindex) + "']" + "[cPtpClockTransDefaultDSInstanceIndex='" + fmt.Sprintf("%v", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsinstanceindex) + "']"
}

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockTransDefaultDSDomainIndex"] = cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsdomainindex
    leafs["cPtpClockTransDefaultDSInstanceIndex"] = cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsinstanceindex
    leafs["cPtpClockTransDefaultDSClockIdentity"] = cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsclockidentity
    leafs["cPtpClockTransDefaultDSNumOfPorts"] = cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsnumofports
    leafs["cPtpClockTransDefaultDSDelay"] = cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsdelay
    leafs["cPtpClockTransDefaultDSPrimaryDomain"] = cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsprimarydomain
    return leafs
}

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetYangName() string { return "cPtpClockTransDefaultDSEntry" }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) SetParent(parent types.Entity) { cptpclocktransdefaultdsentry.parent = parent }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetParent() types.Entity { return cptpclocktransdefaultdsentry.parent }

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetParentYangName() string { return "cPtpClockTransDefaultDSTable" }

// CISCOPTPMIB_Cptpclockporttable
// Table of information about the clock ports for a particular
// domain.
type CISCOPTPMIB_Cptpclockporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single clock port.
    // The type is slice of CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry.
    Cptpclockportentry []CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetFilter() yfilter.YFilter { return cptpclockporttable.YFilter }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) SetFilter(yf yfilter.YFilter) { cptpclockporttable.YFilter = yf }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetGoName(yname string) string {
    if yname == "cPtpClockPortEntry" { return "Cptpclockportentry" }
    return ""
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetSegmentPath() string {
    return "cPtpClockPortTable"
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockPortEntry" {
        for _, c := range cptpclockporttable.Cptpclockportentry {
            if cptpclockporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry{}
        cptpclockporttable.Cptpclockportentry = append(cptpclockporttable.Cptpclockportentry, child)
        return &cptpclockporttable.Cptpclockportentry[len(cptpclockporttable.Cptpclockportentry)-1]
    }
    return nil
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockporttable.Cptpclockportentry {
        children[cptpclockporttable.Cptpclockportentry[i].GetSegmentPath()] = &cptpclockporttable.Cptpclockportentry[i]
    }
    return children
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetYangName() string { return "cPtpClockPortTable" }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) SetParent(parent types.Entity) { cptpclockporttable.parent = parent }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetParent() types.Entity { return cptpclockporttable.parent }

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry
// An entry in the table, containing information about a single
// clock port.
type CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockportdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockportclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockportclockinstanceindex interface{}

    // This attribute is a key. This object specifies the PTP Portnumber for this
    // port. The type is interface{} with range: 1..65535.
    Cptpclockporttableportnumberindex interface{}

    // This object specifies the PTP clock port name configured on the router. The
    // type is string with length: 1..64.
    Cptpclockportname interface{}

    // This object describes the current role (slave/master) of the port. The type
    // is ClockRoleType.
    Cptpclockportrole interface{}

    // This object specifies that one-step clock operation between the PTP master
    // and slave device is enabled. The type is bool.
    Cptpclockportsynconestep interface{}

    // This object specifies the current peer's network address used for PTP
    // communication. Based on the scenario and the setup involved, the values
    // might look like these - Scenario                   Value
    // -------------------   ---------------- Single Master          master port
    // Multiple Masters       selected master port Single Slave           slave
    // port Multiple Slaves        <empty>  (In relevant setups, information on
    // available slaves and available masters will be available through 
    // cPtpClockPortAssociateTable). The type is InetAddressType.
    Cptpclockportcurrentpeeraddresstype interface{}

    // This object specifies the current peer's network address used for PTP
    // communication. Based on the scenario and the setup involved, the values
    // might look like these - Scenario                   Value
    // -------------------   ---------------- Single Master          master port
    // Multiple Masters       selected master port Single Slave           slave
    // port Multiple Slaves        <empty>  (In relevant setups, information on
    // available slaves and available masters will be available through 
    // cPtpClockPortAssociateTable). The type is string with length: 0..255.
    Cptpclockportcurrentpeeraddress interface{}

    // This object specifies - For a master port - the number of PTP slave
    // sessions (peers) associated with this PTP port. For a slave port - the
    // number of masters available to this slave port (might or might not be
    // peered). The type is interface{} with range: 0..4294967295.
    Cptpclockportnumofassociatedports interface{}
}

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetFilter() yfilter.YFilter { return cptpclockportentry.YFilter }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) SetFilter(yf yfilter.YFilter) { cptpclockportentry.YFilter = yf }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetGoName(yname string) string {
    if yname == "cPtpClockPortDomainIndex" { return "Cptpclockportdomainindex" }
    if yname == "cPtpClockPortClockTypeIndex" { return "Cptpclockportclocktypeindex" }
    if yname == "cPtpClockPortClockInstanceIndex" { return "Cptpclockportclockinstanceindex" }
    if yname == "cPtpClockPortTablePortNumberIndex" { return "Cptpclockporttableportnumberindex" }
    if yname == "cPtpClockPortName" { return "Cptpclockportname" }
    if yname == "cPtpClockPortRole" { return "Cptpclockportrole" }
    if yname == "cPtpClockPortSyncOneStep" { return "Cptpclockportsynconestep" }
    if yname == "cPtpClockPortCurrentPeerAddressType" { return "Cptpclockportcurrentpeeraddresstype" }
    if yname == "cPtpClockPortCurrentPeerAddress" { return "Cptpclockportcurrentpeeraddress" }
    if yname == "cPtpClockPortNumOfAssociatedPorts" { return "Cptpclockportnumofassociatedports" }
    return ""
}

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetSegmentPath() string {
    return "cPtpClockPortEntry" + "[cPtpClockPortDomainIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockportdomainindex) + "']" + "[cPtpClockPortClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockportclocktypeindex) + "']" + "[cPtpClockPortClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockportclockinstanceindex) + "']" + "[cPtpClockPortTablePortNumberIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockporttableportnumberindex) + "']"
}

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockPortDomainIndex"] = cptpclockportentry.Cptpclockportdomainindex
    leafs["cPtpClockPortClockTypeIndex"] = cptpclockportentry.Cptpclockportclocktypeindex
    leafs["cPtpClockPortClockInstanceIndex"] = cptpclockportentry.Cptpclockportclockinstanceindex
    leafs["cPtpClockPortTablePortNumberIndex"] = cptpclockportentry.Cptpclockporttableportnumberindex
    leafs["cPtpClockPortName"] = cptpclockportentry.Cptpclockportname
    leafs["cPtpClockPortRole"] = cptpclockportentry.Cptpclockportrole
    leafs["cPtpClockPortSyncOneStep"] = cptpclockportentry.Cptpclockportsynconestep
    leafs["cPtpClockPortCurrentPeerAddressType"] = cptpclockportentry.Cptpclockportcurrentpeeraddresstype
    leafs["cPtpClockPortCurrentPeerAddress"] = cptpclockportentry.Cptpclockportcurrentpeeraddress
    leafs["cPtpClockPortNumOfAssociatedPorts"] = cptpclockportentry.Cptpclockportnumofassociatedports
    return leafs
}

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetYangName() string { return "cPtpClockPortEntry" }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) SetParent(parent types.Entity) { cptpclockportentry.parent = parent }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetParent() types.Entity { return cptpclockportentry.parent }

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetParentYangName() string { return "cPtpClockPortTable" }

// CISCOPTPMIB_Cptpclockportdstable
// Table of information about the clock ports dataset for a
// particular domain.
type CISCOPTPMIB_Cptpclockportdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing port dataset information for a single
    // clock port. The type is slice of
    // CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry.
    Cptpclockportdsentry []CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetFilter() yfilter.YFilter { return cptpclockportdstable.YFilter }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) SetFilter(yf yfilter.YFilter) { cptpclockportdstable.YFilter = yf }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetGoName(yname string) string {
    if yname == "cPtpClockPortDSEntry" { return "Cptpclockportdsentry" }
    return ""
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetSegmentPath() string {
    return "cPtpClockPortDSTable"
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockPortDSEntry" {
        for _, c := range cptpclockportdstable.Cptpclockportdsentry {
            if cptpclockportdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry{}
        cptpclockportdstable.Cptpclockportdsentry = append(cptpclockportdstable.Cptpclockportdsentry, child)
        return &cptpclockportdstable.Cptpclockportdsentry[len(cptpclockportdstable.Cptpclockportdsentry)-1]
    }
    return nil
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockportdstable.Cptpclockportdsentry {
        children[cptpclockportdstable.Cptpclockportdsentry[i].GetSegmentPath()] = &cptpclockportdstable.Cptpclockportdsentry[i]
    }
    return children
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetYangName() string { return "cPtpClockPortDSTable" }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) SetParent(parent types.Entity) { cptpclockportdstable.parent = parent }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetParent() types.Entity { return cptpclockportdstable.parent }

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry
// An entry in the table, containing port dataset information for
// a single clock port.
type CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockportdsdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockportdsclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockportdsclockinstanceindex interface{}

    // This attribute is a key. This object specifies the PTP portnumber
    // associated with this PTP port. The type is interface{} with range:
    // 1..65535.
    Cptpclockportdsportnumberindex interface{}

    // This object specifies the PTP clock port name. The type is string with
    // length: 1..64.
    Cptpclockportdsname interface{}

    // This object specifies the PTP clock port Identity. The type is string.
    Cptpclockportdsportidentity interface{}

    // This object specifies the Announce message transmission interval associated
    // with this clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    Cptpclockportdsannouncementinterval interface{}

    // This object specifies the Announce receipt timeout associated with this
    // clock port. The type is interface{} with range: -2147483648..2147483647.
    Cptpclockportdsannouncercttimeout interface{}

    // This object specifies the Sync message transmission interval. The type is
    // interface{} with range: -2147483648..2147483647.
    Cptpclockportdssyncinterval interface{}

    // This object specifies the Delay_Req message transmission interval. The type
    // is interface{} with range: -2147483648..2147483647.
    Cptpclockportdsmindelayreqinterval interface{}

    // This object specifies the Pdelay_Req message transmission interval. The
    // type is interface{} with range: -2147483648..2147483647.
    Cptpclockportdspeerdelayreqinterval interface{}

    // This object specifies the delay mechanism used. If the clock is an
    // end-to-end clock, the value of the is e2e, else if the clock is a peer
    // to-peer clock, the value shall be p2p. The type is ClockMechanismType.
    Cptpclockportdsdelaymech interface{}

    // This object specifies the peer meanPathDelay. The type is string with
    // length: 1..255.
    Cptpclockportdspeermeanpathdelay interface{}

    // This object specifies the grant duration allocated by the master. The type
    // is interface{} with range: 0..4294967295.
    Cptpclockportdsgrantduration interface{}

    // This object specifies the PTP version being used. The type is interface{}
    // with range: -2147483648..2147483647.
    Cptpclockportdsptpversion interface{}
}

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetFilter() yfilter.YFilter { return cptpclockportdsentry.YFilter }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) SetFilter(yf yfilter.YFilter) { cptpclockportdsentry.YFilter = yf }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockPortDSDomainIndex" { return "Cptpclockportdsdomainindex" }
    if yname == "cPtpClockPortDSClockTypeIndex" { return "Cptpclockportdsclocktypeindex" }
    if yname == "cPtpClockPortDSClockInstanceIndex" { return "Cptpclockportdsclockinstanceindex" }
    if yname == "cPtpClockPortDSPortNumberIndex" { return "Cptpclockportdsportnumberindex" }
    if yname == "cPtpClockPortDSName" { return "Cptpclockportdsname" }
    if yname == "cPtpClockPortDSPortIdentity" { return "Cptpclockportdsportidentity" }
    if yname == "cPtpClockPortDSAnnouncementInterval" { return "Cptpclockportdsannouncementinterval" }
    if yname == "cPtpClockPortDSAnnounceRctTimeout" { return "Cptpclockportdsannouncercttimeout" }
    if yname == "cPtpClockPortDSSyncInterval" { return "Cptpclockportdssyncinterval" }
    if yname == "cPtpClockPortDSMinDelayReqInterval" { return "Cptpclockportdsmindelayreqinterval" }
    if yname == "cPtpClockPortDSPeerDelayReqInterval" { return "Cptpclockportdspeerdelayreqinterval" }
    if yname == "cPtpClockPortDSDelayMech" { return "Cptpclockportdsdelaymech" }
    if yname == "cPtpClockPortDSPeerMeanPathDelay" { return "Cptpclockportdspeermeanpathdelay" }
    if yname == "cPtpClockPortDSGrantDuration" { return "Cptpclockportdsgrantduration" }
    if yname == "cPtpClockPortDSPTPVersion" { return "Cptpclockportdsptpversion" }
    return ""
}

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetSegmentPath() string {
    return "cPtpClockPortDSEntry" + "[cPtpClockPortDSDomainIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsdomainindex) + "']" + "[cPtpClockPortDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsclocktypeindex) + "']" + "[cPtpClockPortDSClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsclockinstanceindex) + "']" + "[cPtpClockPortDSPortNumberIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsportnumberindex) + "']"
}

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockPortDSDomainIndex"] = cptpclockportdsentry.Cptpclockportdsdomainindex
    leafs["cPtpClockPortDSClockTypeIndex"] = cptpclockportdsentry.Cptpclockportdsclocktypeindex
    leafs["cPtpClockPortDSClockInstanceIndex"] = cptpclockportdsentry.Cptpclockportdsclockinstanceindex
    leafs["cPtpClockPortDSPortNumberIndex"] = cptpclockportdsentry.Cptpclockportdsportnumberindex
    leafs["cPtpClockPortDSName"] = cptpclockportdsentry.Cptpclockportdsname
    leafs["cPtpClockPortDSPortIdentity"] = cptpclockportdsentry.Cptpclockportdsportidentity
    leafs["cPtpClockPortDSAnnouncementInterval"] = cptpclockportdsentry.Cptpclockportdsannouncementinterval
    leafs["cPtpClockPortDSAnnounceRctTimeout"] = cptpclockportdsentry.Cptpclockportdsannouncercttimeout
    leafs["cPtpClockPortDSSyncInterval"] = cptpclockportdsentry.Cptpclockportdssyncinterval
    leafs["cPtpClockPortDSMinDelayReqInterval"] = cptpclockportdsentry.Cptpclockportdsmindelayreqinterval
    leafs["cPtpClockPortDSPeerDelayReqInterval"] = cptpclockportdsentry.Cptpclockportdspeerdelayreqinterval
    leafs["cPtpClockPortDSDelayMech"] = cptpclockportdsentry.Cptpclockportdsdelaymech
    leafs["cPtpClockPortDSPeerMeanPathDelay"] = cptpclockportdsentry.Cptpclockportdspeermeanpathdelay
    leafs["cPtpClockPortDSGrantDuration"] = cptpclockportdsentry.Cptpclockportdsgrantduration
    leafs["cPtpClockPortDSPTPVersion"] = cptpclockportdsentry.Cptpclockportdsptpversion
    return leafs
}

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetYangName() string { return "cPtpClockPortDSEntry" }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) SetParent(parent types.Entity) { cptpclockportdsentry.parent = parent }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetParent() types.Entity { return cptpclockportdsentry.parent }

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetParentYangName() string { return "cPtpClockPortDSTable" }

// CISCOPTPMIB_Cptpclockportrunningtable
// Table of information about the clock ports running dataset for
// a particular domain.
type CISCOPTPMIB_Cptpclockportrunningtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing runing dataset information about a single
    // clock port. The type is slice of
    // CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry.
    Cptpclockportrunningentry []CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetFilter() yfilter.YFilter { return cptpclockportrunningtable.YFilter }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) SetFilter(yf yfilter.YFilter) { cptpclockportrunningtable.YFilter = yf }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetGoName(yname string) string {
    if yname == "cPtpClockPortRunningEntry" { return "Cptpclockportrunningentry" }
    return ""
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetSegmentPath() string {
    return "cPtpClockPortRunningTable"
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockPortRunningEntry" {
        for _, c := range cptpclockportrunningtable.Cptpclockportrunningentry {
            if cptpclockportrunningtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry{}
        cptpclockportrunningtable.Cptpclockportrunningentry = append(cptpclockportrunningtable.Cptpclockportrunningentry, child)
        return &cptpclockportrunningtable.Cptpclockportrunningentry[len(cptpclockportrunningtable.Cptpclockportrunningentry)-1]
    }
    return nil
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockportrunningtable.Cptpclockportrunningentry {
        children[cptpclockportrunningtable.Cptpclockportrunningentry[i].GetSegmentPath()] = &cptpclockportrunningtable.Cptpclockportrunningentry[i]
    }
    return children
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetYangName() string { return "cPtpClockPortRunningTable" }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) SetParent(parent types.Entity) { cptpclockportrunningtable.parent = parent }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetParent() types.Entity { return cptpclockportrunningtable.parent }

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry
// An entry in the table, containing runing dataset information
// about a single clock port.
type CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockportrunningdomainindex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpclockportrunningclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockportrunningclockinstanceindex interface{}

    // This attribute is a key. This object specifies the PTP portnumber
    // associated with this clock port. The type is interface{} with range:
    // 1..65535.
    Cptpclockportrunningportnumberindex interface{}

    // This object specifies the PTP clock port name. The type is string with
    // length: 1..64.
    Cptpclockportrunningname interface{}

    // This object specifies the port state returned by PTP engine.  initializing
    // - In this state a port initializes                its data sets, hardware,
    // and                communication facilities. faulty       - The fault state
    // of the protocol. disabled     - The port shall not place any               
    // messages on its communication path. listening    - The port is waiting for
    // the                announceReceiptTimeout to expire or                to
    // receive an Announce message from                a master. preMaster    -
    // The port shall behave in all respects                as though it were in
    // the MASTER state                except that it shall not place any         
    // messages on its communication path                except for Pdelay_Req,
    // Pdelay_Resp,                Pdelay_Resp_Follow_Up, signaling, or           
    // management messages. master       - The port is behaving as a master port. 
    // passive      - The port shall not place any                messages on its
    // communication path                except for Pdelay_Req, Pdelay_Resp,      
    // Pdelay_Resp_Follow_Up, or signaling                messages, or management
    // messages                that are a required response to               
    // another management message uncalibrated - The local port is preparing to   
    // synchronize to the master port. slave        - The port is synchronizing to
    // the                selected master port. The type is ClockPortState.
    Cptpclockportrunningstate interface{}

    // This object specifies the Clock Role. The type is ClockRoleType.
    Cptpclockportrunningrole interface{}

    // This object specifies the interface on the router being used by the PTP
    // Clock for PTP communication. The type is interface{} with range:
    // 0..2147483647.
    Cptpclockportrunninginterfaceindex interface{}

    // This object specifirst the IP version being used for PTP communication (the
    // mapping used). The type is interface{} with range: -2147483648..2147483647.
    Cptpclockportrunningipversion interface{}

    // This object specifies the type of encapsulation if the interface is adding
    // extra  layers (eg. VLAN, Pseudowire encapsulation...) for the PTP messages.
    // The type is interface{} with range: -2147483648..2147483647.
    Cptpclockportrunningencapsulationtype interface{}

    // This object specifies the clock transmission mode as  unicast:       Using
    // unicast commnuication channel. multicast:     Using Multicast communication
    // channel. multicast-mix: Using multicast-unicast communication channel. The
    // type is ClockTxModeType.
    Cptpclockportrunningtxmode interface{}

    // This object specifie the clock receive mode as  unicast:       Using
    // unicast commnuication channel. multicast:     Using Multicast communication
    // channel. multicast-mix: Using multicast-unicast communication channel. The
    // type is ClockTxModeType.
    Cptpclockportrunningrxmode interface{}

    // This object specifies the packets received on the clock port (cummulative).
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Cptpclockportrunningpacketsreceived interface{}

    // This object specifies the packets sent on the clock port (cummulative). The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    Cptpclockportrunningpacketssent interface{}
}

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetFilter() yfilter.YFilter { return cptpclockportrunningentry.YFilter }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) SetFilter(yf yfilter.YFilter) { cptpclockportrunningentry.YFilter = yf }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetGoName(yname string) string {
    if yname == "cPtpClockPortRunningDomainIndex" { return "Cptpclockportrunningdomainindex" }
    if yname == "cPtpClockPortRunningClockTypeIndex" { return "Cptpclockportrunningclocktypeindex" }
    if yname == "cPtpClockPortRunningClockInstanceIndex" { return "Cptpclockportrunningclockinstanceindex" }
    if yname == "cPtpClockPortRunningPortNumberIndex" { return "Cptpclockportrunningportnumberindex" }
    if yname == "cPtpClockPortRunningName" { return "Cptpclockportrunningname" }
    if yname == "cPtpClockPortRunningState" { return "Cptpclockportrunningstate" }
    if yname == "cPtpClockPortRunningRole" { return "Cptpclockportrunningrole" }
    if yname == "cPtpClockPortRunningInterfaceIndex" { return "Cptpclockportrunninginterfaceindex" }
    if yname == "cPtpClockPortRunningIPversion" { return "Cptpclockportrunningipversion" }
    if yname == "cPtpClockPortRunningEncapsulationType" { return "Cptpclockportrunningencapsulationtype" }
    if yname == "cPtpClockPortRunningTxMode" { return "Cptpclockportrunningtxmode" }
    if yname == "cPtpClockPortRunningRxMode" { return "Cptpclockportrunningrxmode" }
    if yname == "cPtpClockPortRunningPacketsReceived" { return "Cptpclockportrunningpacketsreceived" }
    if yname == "cPtpClockPortRunningPacketsSent" { return "Cptpclockportrunningpacketssent" }
    return ""
}

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetSegmentPath() string {
    return "cPtpClockPortRunningEntry" + "[cPtpClockPortRunningDomainIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningdomainindex) + "']" + "[cPtpClockPortRunningClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningclocktypeindex) + "']" + "[cPtpClockPortRunningClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningclockinstanceindex) + "']" + "[cPtpClockPortRunningPortNumberIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningportnumberindex) + "']"
}

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockPortRunningDomainIndex"] = cptpclockportrunningentry.Cptpclockportrunningdomainindex
    leafs["cPtpClockPortRunningClockTypeIndex"] = cptpclockportrunningentry.Cptpclockportrunningclocktypeindex
    leafs["cPtpClockPortRunningClockInstanceIndex"] = cptpclockportrunningentry.Cptpclockportrunningclockinstanceindex
    leafs["cPtpClockPortRunningPortNumberIndex"] = cptpclockportrunningentry.Cptpclockportrunningportnumberindex
    leafs["cPtpClockPortRunningName"] = cptpclockportrunningentry.Cptpclockportrunningname
    leafs["cPtpClockPortRunningState"] = cptpclockportrunningentry.Cptpclockportrunningstate
    leafs["cPtpClockPortRunningRole"] = cptpclockportrunningentry.Cptpclockportrunningrole
    leafs["cPtpClockPortRunningInterfaceIndex"] = cptpclockportrunningentry.Cptpclockportrunninginterfaceindex
    leafs["cPtpClockPortRunningIPversion"] = cptpclockportrunningentry.Cptpclockportrunningipversion
    leafs["cPtpClockPortRunningEncapsulationType"] = cptpclockportrunningentry.Cptpclockportrunningencapsulationtype
    leafs["cPtpClockPortRunningTxMode"] = cptpclockportrunningentry.Cptpclockportrunningtxmode
    leafs["cPtpClockPortRunningRxMode"] = cptpclockportrunningentry.Cptpclockportrunningrxmode
    leafs["cPtpClockPortRunningPacketsReceived"] = cptpclockportrunningentry.Cptpclockportrunningpacketsreceived
    leafs["cPtpClockPortRunningPacketsSent"] = cptpclockportrunningentry.Cptpclockportrunningpacketssent
    return leafs
}

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetYangName() string { return "cPtpClockPortRunningEntry" }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) SetParent(parent types.Entity) { cptpclockportrunningentry.parent = parent }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetParent() types.Entity { return cptpclockportrunningentry.parent }

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetParentYangName() string { return "cPtpClockPortRunningTable" }

// CISCOPTPMIB_Cptpclockporttransdstable
// Table of information about the Transparent clock ports running
// dataset for a particular domain.
type CISCOPTPMIB_Cptpclockporttransdstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing clock port Transparent dataset
    // information about a single clock port. The type is slice of
    // CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry.
    Cptpclockporttransdsentry []CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetFilter() yfilter.YFilter { return cptpclockporttransdstable.YFilter }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) SetFilter(yf yfilter.YFilter) { cptpclockporttransdstable.YFilter = yf }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetGoName(yname string) string {
    if yname == "cPtpClockPortTransDSEntry" { return "Cptpclockporttransdsentry" }
    return ""
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetSegmentPath() string {
    return "cPtpClockPortTransDSTable"
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockPortTransDSEntry" {
        for _, c := range cptpclockporttransdstable.Cptpclockporttransdsentry {
            if cptpclockporttransdstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry{}
        cptpclockporttransdstable.Cptpclockporttransdsentry = append(cptpclockporttransdstable.Cptpclockporttransdsentry, child)
        return &cptpclockporttransdstable.Cptpclockporttransdsentry[len(cptpclockporttransdstable.Cptpclockporttransdsentry)-1]
    }
    return nil
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockporttransdstable.Cptpclockporttransdsentry {
        children[cptpclockporttransdstable.Cptpclockporttransdsentry[i].GetSegmentPath()] = &cptpclockporttransdstable.Cptpclockporttransdsentry[i]
    }
    return children
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetYangName() string { return "cPtpClockPortTransDSTable" }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) SetParent(parent types.Entity) { cptpclockporttransdstable.parent = parent }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetParent() types.Entity { return cptpclockporttransdstable.parent }

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry
// An entry in the table, containing clock port Transparent
// dataset information about a single clock port
type CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    Cptpclockporttransdsdomainindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockporttransdsinstanceindex interface{}

    // This attribute is a key. This object specifies the PTP port number
    // associated with this port. The type is interface{} with range: 1..65535.
    Cptpclockporttransdsportnumberindex interface{}

    // This object specifies the value of the PortIdentity attribute of the local
    // port. The type is string with length: 1..255.
    Cptpclockporttransdsportidentity interface{}

    // This object specifies the value of the logarithm to the base 2 of the
    // minPdelayReqInterval. The type is interface{} with range:
    // -2147483648..2147483647.
    Cptpclockporttransdslogminpdelayreqint interface{}

    // This object specifies the value TRUE if the port is faulty and FALSE if the
    // port is operating normally. The type is bool.
    Cptpclockporttransdsfaultyflag interface{}

    // This object specifies, (if the delayMechanism used is P2P) the value is the
    // estimate of the current one-way propagation delay, i.e., <meanPathDelay> on
    // the link attached to this port computed using the peer delay mechanism. If
    // the value of the delayMechanism used is E2E, then the value will be zero.
    // The type is string with length: 1..255.
    Cptpclockporttransdspeermeanpathdelay interface{}
}

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetFilter() yfilter.YFilter { return cptpclockporttransdsentry.YFilter }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) SetFilter(yf yfilter.YFilter) { cptpclockporttransdsentry.YFilter = yf }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetGoName(yname string) string {
    if yname == "cPtpClockPortTransDSDomainIndex" { return "Cptpclockporttransdsdomainindex" }
    if yname == "cPtpClockPortTransDSInstanceIndex" { return "Cptpclockporttransdsinstanceindex" }
    if yname == "cPtpClockPortTransDSPortNumberIndex" { return "Cptpclockporttransdsportnumberindex" }
    if yname == "cPtpClockPortTransDSPortIdentity" { return "Cptpclockporttransdsportidentity" }
    if yname == "cPtpClockPortTransDSlogMinPdelayReqInt" { return "Cptpclockporttransdslogminpdelayreqint" }
    if yname == "cPtpClockPortTransDSFaultyFlag" { return "Cptpclockporttransdsfaultyflag" }
    if yname == "cPtpClockPortTransDSPeerMeanPathDelay" { return "Cptpclockporttransdspeermeanpathdelay" }
    return ""
}

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetSegmentPath() string {
    return "cPtpClockPortTransDSEntry" + "[cPtpClockPortTransDSDomainIndex='" + fmt.Sprintf("%v", cptpclockporttransdsentry.Cptpclockporttransdsdomainindex) + "']" + "[cPtpClockPortTransDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockporttransdsentry.Cptpclockporttransdsinstanceindex) + "']" + "[cPtpClockPortTransDSPortNumberIndex='" + fmt.Sprintf("%v", cptpclockporttransdsentry.Cptpclockporttransdsportnumberindex) + "']"
}

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockPortTransDSDomainIndex"] = cptpclockporttransdsentry.Cptpclockporttransdsdomainindex
    leafs["cPtpClockPortTransDSInstanceIndex"] = cptpclockporttransdsentry.Cptpclockporttransdsinstanceindex
    leafs["cPtpClockPortTransDSPortNumberIndex"] = cptpclockporttransdsentry.Cptpclockporttransdsportnumberindex
    leafs["cPtpClockPortTransDSPortIdentity"] = cptpclockporttransdsentry.Cptpclockporttransdsportidentity
    leafs["cPtpClockPortTransDSlogMinPdelayReqInt"] = cptpclockporttransdsentry.Cptpclockporttransdslogminpdelayreqint
    leafs["cPtpClockPortTransDSFaultyFlag"] = cptpclockporttransdsentry.Cptpclockporttransdsfaultyflag
    leafs["cPtpClockPortTransDSPeerMeanPathDelay"] = cptpclockporttransdsentry.Cptpclockporttransdspeermeanpathdelay
    return leafs
}

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetYangName() string { return "cPtpClockPortTransDSEntry" }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) SetParent(parent types.Entity) { cptpclockporttransdsentry.parent = parent }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetParent() types.Entity { return cptpclockporttransdsentry.parent }

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetParentYangName() string { return "cPtpClockPortTransDSTable" }

// CISCOPTPMIB_Cptpclockportassociatetable
// Table of information about a given port's associated ports.
// 
// For a master port - multiple slave ports which have established
// sessions with the current master port.  
// For a slave port - the list of masters available for a given
// slave port. 
// 
// Session information (pkts, errors) to be displayed based on
// availability and scenario.
type CISCOPTPMIB_Cptpclockportassociatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single associated
    // port for the given clockport. The type is slice of
    // CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry.
    Cptpclockportassociateentry []CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetFilter() yfilter.YFilter { return cptpclockportassociatetable.YFilter }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) SetFilter(yf yfilter.YFilter) { cptpclockportassociatetable.YFilter = yf }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetGoName(yname string) string {
    if yname == "cPtpClockPortAssociateEntry" { return "Cptpclockportassociateentry" }
    return ""
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetSegmentPath() string {
    return "cPtpClockPortAssociateTable"
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPtpClockPortAssociateEntry" {
        for _, c := range cptpclockportassociatetable.Cptpclockportassociateentry {
            if cptpclockportassociatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry{}
        cptpclockportassociatetable.Cptpclockportassociateentry = append(cptpclockportassociatetable.Cptpclockportassociateentry, child)
        return &cptpclockportassociatetable.Cptpclockportassociateentry[len(cptpclockportassociatetable.Cptpclockportassociateentry)-1]
    }
    return nil
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cptpclockportassociatetable.Cptpclockportassociateentry {
        children[cptpclockportassociatetable.Cptpclockportassociateentry[i].GetSegmentPath()] = &cptpclockportassociatetable.Cptpclockportassociateentry[i]
    }
    return children
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetYangName() string { return "cPtpClockPortAssociateTable" }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) SetParent(parent types.Entity) { cptpclockportassociatetable.parent = parent }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetParent() types.Entity { return cptpclockportassociatetable.parent }

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetParentYangName() string { return "CISCO-PTP-MIB" }

// CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry
// An entry in the table, containing information about a single
// associated port for the given clockport.
type CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the given port's domain
    // number. The type is interface{} with range: 0..255.
    Cptpclockportcurrentdomainindex interface{}

    // This attribute is a key. This object specifies the given port's clock type.
    // The type is ClockType.
    Cptpclockportcurrentclocktypeindex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    Cptpclockportcurrentclockinstanceindex interface{}

    // This attribute is a key. This object specifies the PTP Port Number for the
    // given port. The type is interface{} with range: 0..65535.
    Cptpclockportcurrentportnumberindex interface{}

    // This attribute is a key. This object specifies the associated port's serial
    // number in the current port's context. The type is interface{} with range:
    // 1..65535.
    Cptpclockportassociateportindex interface{}

    // This object specifies the peer port's network address type used for PTP
    // communication. The type is InetAddressType.
    Cptpclockportassociateaddresstype interface{}

    // This object specifies the peer port's network address used for PTP
    // communication. The type is string with length: 0..255.
    Cptpclockportassociateaddress interface{}

    // The number of packets sent to this peer port from the current port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    Cptpclockportassociatepacketssent interface{}

    // The number of packets received from this peer port by the current port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    Cptpclockportassociatepacketsreceived interface{}

    // This object specifies the input errors associated with the peer port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    Cptpclockportassociateinerrors interface{}

    // This object specifies the output errors associated with the peer port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    Cptpclockportassociateouterrors interface{}
}

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetFilter() yfilter.YFilter { return cptpclockportassociateentry.YFilter }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) SetFilter(yf yfilter.YFilter) { cptpclockportassociateentry.YFilter = yf }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetGoName(yname string) string {
    if yname == "cPtpClockPortCurrentDomainIndex" { return "Cptpclockportcurrentdomainindex" }
    if yname == "cPtpClockPortCurrentClockTypeIndex" { return "Cptpclockportcurrentclocktypeindex" }
    if yname == "cPtpClockPortCurrentClockInstanceIndex" { return "Cptpclockportcurrentclockinstanceindex" }
    if yname == "cPtpClockPortCurrentPortNumberIndex" { return "Cptpclockportcurrentportnumberindex" }
    if yname == "cPtpClockPortAssociatePortIndex" { return "Cptpclockportassociateportindex" }
    if yname == "cPtpClockPortAssociateAddressType" { return "Cptpclockportassociateaddresstype" }
    if yname == "cPtpClockPortAssociateAddress" { return "Cptpclockportassociateaddress" }
    if yname == "cPtpClockPortAssociatePacketsSent" { return "Cptpclockportassociatepacketssent" }
    if yname == "cPtpClockPortAssociatePacketsReceived" { return "Cptpclockportassociatepacketsreceived" }
    if yname == "cPtpClockPortAssociateInErrors" { return "Cptpclockportassociateinerrors" }
    if yname == "cPtpClockPortAssociateOutErrors" { return "Cptpclockportassociateouterrors" }
    return ""
}

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetSegmentPath() string {
    return "cPtpClockPortAssociateEntry" + "[cPtpClockPortCurrentDomainIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentdomainindex) + "']" + "[cPtpClockPortCurrentClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentclocktypeindex) + "']" + "[cPtpClockPortCurrentClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentclockinstanceindex) + "']" + "[cPtpClockPortCurrentPortNumberIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentportnumberindex) + "']" + "[cPtpClockPortAssociatePortIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportassociateportindex) + "']"
}

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPtpClockPortCurrentDomainIndex"] = cptpclockportassociateentry.Cptpclockportcurrentdomainindex
    leafs["cPtpClockPortCurrentClockTypeIndex"] = cptpclockportassociateentry.Cptpclockportcurrentclocktypeindex
    leafs["cPtpClockPortCurrentClockInstanceIndex"] = cptpclockportassociateentry.Cptpclockportcurrentclockinstanceindex
    leafs["cPtpClockPortCurrentPortNumberIndex"] = cptpclockportassociateentry.Cptpclockportcurrentportnumberindex
    leafs["cPtpClockPortAssociatePortIndex"] = cptpclockportassociateentry.Cptpclockportassociateportindex
    leafs["cPtpClockPortAssociateAddressType"] = cptpclockportassociateentry.Cptpclockportassociateaddresstype
    leafs["cPtpClockPortAssociateAddress"] = cptpclockportassociateentry.Cptpclockportassociateaddress
    leafs["cPtpClockPortAssociatePacketsSent"] = cptpclockportassociateentry.Cptpclockportassociatepacketssent
    leafs["cPtpClockPortAssociatePacketsReceived"] = cptpclockportassociateentry.Cptpclockportassociatepacketsreceived
    leafs["cPtpClockPortAssociateInErrors"] = cptpclockportassociateentry.Cptpclockportassociateinerrors
    leafs["cPtpClockPortAssociateOutErrors"] = cptpclockportassociateentry.Cptpclockportassociateouterrors
    return leafs
}

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetYangName() string { return "cPtpClockPortAssociateEntry" }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) SetParent(parent types.Entity) { cptpclockportassociateentry.parent = parent }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetParent() types.Entity { return cptpclockportassociateentry.parent }

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetParentYangName() string { return "cPtpClockPortAssociateTable" }

