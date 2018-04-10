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

// ClockProfileType represents a device.
type ClockProfileType string

const (
    ClockProfileType_default_ ClockProfileType = "default"

    ClockProfileType_telecom ClockProfileType = "telecom"

    ClockProfileType_vendorspecific ClockProfileType = "vendorspecific"
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

// ClockRoleType represents                           another clock (master).
type ClockRoleType string

const (
    ClockRoleType_master ClockRoleType = "master"

    ClockRoleType_slave ClockRoleType = "slave"
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

// ClockTxModeType represents multicast-mix. Using multicast-unicast communication channel
type ClockTxModeType string

const (
    ClockTxModeType_unicast ClockTxModeType = "unicast"

    ClockTxModeType_multicast ClockTxModeType = "multicast"

    ClockTxModeType_multicastmix ClockTxModeType = "multicastmix"
)

// ClockType represents The clock types as defined in the MIB module description.
type ClockType string

const (
    ClockType_ordinaryClock ClockType = "ordinaryClock"

    ClockType_boundaryClock ClockType = "boundaryClock"

    ClockType_transparentClock ClockType = "transparentClock"

    ClockType_boundaryNode ClockType = "boundaryNode"
)

// CISCOPTPMIB
type CISCOPTPMIB struct {
    EntityData types.CommonEntityData
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

func (cISCOPTPMIB *CISCOPTPMIB) GetEntityData() *types.CommonEntityData {
    cISCOPTPMIB.EntityData.YFilter = cISCOPTPMIB.YFilter
    cISCOPTPMIB.EntityData.YangName = "CISCO-PTP-MIB"
    cISCOPTPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPTPMIB.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cISCOPTPMIB.EntityData.SegmentPath = "CISCO-PTP-MIB:CISCO-PTP-MIB"
    cISCOPTPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPTPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPTPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPTPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOPTPMIB.EntityData.Children["ciscoPtpMIBSystemInfo"] = types.YChild{"Ciscoptpmibsysteminfo", &cISCOPTPMIB.Ciscoptpmibsysteminfo}
    cISCOPTPMIB.EntityData.Children["cPtpSystemTable"] = types.YChild{"Cptpsystemtable", &cISCOPTPMIB.Cptpsystemtable}
    cISCOPTPMIB.EntityData.Children["cPtpSystemDomainTable"] = types.YChild{"Cptpsystemdomaintable", &cISCOPTPMIB.Cptpsystemdomaintable}
    cISCOPTPMIB.EntityData.Children["cPtpClockNodeTable"] = types.YChild{"Cptpclocknodetable", &cISCOPTPMIB.Cptpclocknodetable}
    cISCOPTPMIB.EntityData.Children["cPtpClockCurrentDSTable"] = types.YChild{"Cptpclockcurrentdstable", &cISCOPTPMIB.Cptpclockcurrentdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockParentDSTable"] = types.YChild{"Cptpclockparentdstable", &cISCOPTPMIB.Cptpclockparentdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockDefaultDSTable"] = types.YChild{"Cptpclockdefaultdstable", &cISCOPTPMIB.Cptpclockdefaultdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockRunningTable"] = types.YChild{"Cptpclockrunningtable", &cISCOPTPMIB.Cptpclockrunningtable}
    cISCOPTPMIB.EntityData.Children["cPtpClockTimePropertiesDSTable"] = types.YChild{"Cptpclocktimepropertiesdstable", &cISCOPTPMIB.Cptpclocktimepropertiesdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockTransDefaultDSTable"] = types.YChild{"Cptpclocktransdefaultdstable", &cISCOPTPMIB.Cptpclocktransdefaultdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockPortTable"] = types.YChild{"Cptpclockporttable", &cISCOPTPMIB.Cptpclockporttable}
    cISCOPTPMIB.EntityData.Children["cPtpClockPortDSTable"] = types.YChild{"Cptpclockportdstable", &cISCOPTPMIB.Cptpclockportdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockPortRunningTable"] = types.YChild{"Cptpclockportrunningtable", &cISCOPTPMIB.Cptpclockportrunningtable}
    cISCOPTPMIB.EntityData.Children["cPtpClockPortTransDSTable"] = types.YChild{"Cptpclockporttransdstable", &cISCOPTPMIB.Cptpclockporttransdstable}
    cISCOPTPMIB.EntityData.Children["cPtpClockPortAssociateTable"] = types.YChild{"Cptpclockportassociatetable", &cISCOPTPMIB.Cptpclockportassociatetable}
    cISCOPTPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOPTPMIB.EntityData)
}

// CISCOPTPMIB_Ciscoptpmibsysteminfo
type CISCOPTPMIB_Ciscoptpmibsysteminfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the PTP Profile implemented on the system. The type
    // is ClockProfileType.
    Cptpsystemprofile interface{}
}

func (ciscoptpmibsysteminfo *CISCOPTPMIB_Ciscoptpmibsysteminfo) GetEntityData() *types.CommonEntityData {
    ciscoptpmibsysteminfo.EntityData.YFilter = ciscoptpmibsysteminfo.YFilter
    ciscoptpmibsysteminfo.EntityData.YangName = "ciscoPtpMIBSystemInfo"
    ciscoptpmibsysteminfo.EntityData.BundleName = "cisco_ios_xe"
    ciscoptpmibsysteminfo.EntityData.ParentYangName = "CISCO-PTP-MIB"
    ciscoptpmibsysteminfo.EntityData.SegmentPath = "ciscoPtpMIBSystemInfo"
    ciscoptpmibsysteminfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoptpmibsysteminfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoptpmibsysteminfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoptpmibsysteminfo.EntityData.Children = make(map[string]types.YChild)
    ciscoptpmibsysteminfo.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoptpmibsysteminfo.EntityData.Leafs["cPtpSystemProfile"] = types.YLeaf{"Cptpsystemprofile", ciscoptpmibsysteminfo.Cptpsystemprofile}
    return &(ciscoptpmibsysteminfo.EntityData)
}

// CISCOPTPMIB_Cptpsystemtable
// Table of count information about the PTP system for all
// domains.
type CISCOPTPMIB_Cptpsystemtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing count information about a single domain.
    // New row entries are added when the PTP clock for this domain is configured,
    // while the unconfiguration of the PTP clock removes it. The type is slice of
    // CISCOPTPMIB_Cptpsystemtable_Cptpsystementry.
    Cptpsystementry []CISCOPTPMIB_Cptpsystemtable_Cptpsystementry
}

func (cptpsystemtable *CISCOPTPMIB_Cptpsystemtable) GetEntityData() *types.CommonEntityData {
    cptpsystemtable.EntityData.YFilter = cptpsystemtable.YFilter
    cptpsystemtable.EntityData.YangName = "cPtpSystemTable"
    cptpsystemtable.EntityData.BundleName = "cisco_ios_xe"
    cptpsystemtable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpsystemtable.EntityData.SegmentPath = "cPtpSystemTable"
    cptpsystemtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpsystemtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpsystemtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpsystemtable.EntityData.Children = make(map[string]types.YChild)
    cptpsystemtable.EntityData.Children["cPtpSystemEntry"] = types.YChild{"Cptpsystementry", nil}
    for i := range cptpsystemtable.Cptpsystementry {
        cptpsystemtable.EntityData.Children[types.GetSegmentPath(&cptpsystemtable.Cptpsystementry[i])] = types.YChild{"Cptpsystementry", &cptpsystemtable.Cptpsystementry[i]}
    }
    cptpsystemtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpsystemtable.EntityData)
}

// CISCOPTPMIB_Cptpsystemtable_Cptpsystementry
// An entry in the table, containing count information about a
// single domain. New row entries are added when the PTP clock for
// this domain is configured, while the unconfiguration of the PTP
// clock removes it.
type CISCOPTPMIB_Cptpsystemtable_Cptpsystementry struct {
    EntityData types.CommonEntityData
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

func (cptpsystementry *CISCOPTPMIB_Cptpsystemtable_Cptpsystementry) GetEntityData() *types.CommonEntityData {
    cptpsystementry.EntityData.YFilter = cptpsystementry.YFilter
    cptpsystementry.EntityData.YangName = "cPtpSystemEntry"
    cptpsystementry.EntityData.BundleName = "cisco_ios_xe"
    cptpsystementry.EntityData.ParentYangName = "cPtpSystemTable"
    cptpsystementry.EntityData.SegmentPath = "cPtpSystemEntry" + "[cPtpDomainIndex='" + fmt.Sprintf("%v", cptpsystementry.Cptpdomainindex) + "']" + "[cPtpInstanceIndex='" + fmt.Sprintf("%v", cptpsystementry.Cptpinstanceindex) + "']"
    cptpsystementry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpsystementry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpsystementry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpsystementry.EntityData.Children = make(map[string]types.YChild)
    cptpsystementry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpsystementry.EntityData.Leafs["cPtpDomainIndex"] = types.YLeaf{"Cptpdomainindex", cptpsystementry.Cptpdomainindex}
    cptpsystementry.EntityData.Leafs["cPtpInstanceIndex"] = types.YLeaf{"Cptpinstanceindex", cptpsystementry.Cptpinstanceindex}
    cptpsystementry.EntityData.Leafs["cPtpDomainClockPortsTotal"] = types.YLeaf{"Cptpdomainclockportstotal", cptpsystementry.Cptpdomainclockportstotal}
    cptpsystementry.EntityData.Leafs["cPtpDomainClockPortPhysicalInterfacesTotal"] = types.YLeaf{"Cptpdomainclockportphysicalinterfacestotal", cptpsystementry.Cptpdomainclockportphysicalinterfacestotal}
    return &(cptpsystementry.EntityData)
}

// CISCOPTPMIB_Cptpsystemdomaintable
// Table of information about the PTP system for all clock modes
// -- ordinary, boundary or transparent.
type CISCOPTPMIB_Cptpsystemdomaintable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single clock mode for
    // the PTP system. A row entry gets added when PTP clocks are configured on
    // the router. The type is slice of
    // CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry.
    Cptpsystemdomainentry []CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry
}

func (cptpsystemdomaintable *CISCOPTPMIB_Cptpsystemdomaintable) GetEntityData() *types.CommonEntityData {
    cptpsystemdomaintable.EntityData.YFilter = cptpsystemdomaintable.YFilter
    cptpsystemdomaintable.EntityData.YangName = "cPtpSystemDomainTable"
    cptpsystemdomaintable.EntityData.BundleName = "cisco_ios_xe"
    cptpsystemdomaintable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpsystemdomaintable.EntityData.SegmentPath = "cPtpSystemDomainTable"
    cptpsystemdomaintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpsystemdomaintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpsystemdomaintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpsystemdomaintable.EntityData.Children = make(map[string]types.YChild)
    cptpsystemdomaintable.EntityData.Children["cPtpSystemDomainEntry"] = types.YChild{"Cptpsystemdomainentry", nil}
    for i := range cptpsystemdomaintable.Cptpsystemdomainentry {
        cptpsystemdomaintable.EntityData.Children[types.GetSegmentPath(&cptpsystemdomaintable.Cptpsystemdomainentry[i])] = types.YChild{"Cptpsystemdomainentry", &cptpsystemdomaintable.Cptpsystemdomainentry[i]}
    }
    cptpsystemdomaintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpsystemdomaintable.EntityData)
}

// CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry
// An entry in the table, containing information about a single
// clock mode for the PTP system. A row entry gets added when PTP
// clocks are configured on the router.
type CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    Cptpsystemdomainclocktypeindex interface{}

    // This object specifies the  total number of PTP domains for this particular
    // clock type configured in this node. The type is interface{} with range:
    // 0..4294967295. Units are domains.
    Cptpsystemdomaintotals interface{}
}

func (cptpsystemdomainentry *CISCOPTPMIB_Cptpsystemdomaintable_Cptpsystemdomainentry) GetEntityData() *types.CommonEntityData {
    cptpsystemdomainentry.EntityData.YFilter = cptpsystemdomainentry.YFilter
    cptpsystemdomainentry.EntityData.YangName = "cPtpSystemDomainEntry"
    cptpsystemdomainentry.EntityData.BundleName = "cisco_ios_xe"
    cptpsystemdomainentry.EntityData.ParentYangName = "cPtpSystemDomainTable"
    cptpsystemdomainentry.EntityData.SegmentPath = "cPtpSystemDomainEntry" + "[cPtpSystemDomainClockTypeIndex='" + fmt.Sprintf("%v", cptpsystemdomainentry.Cptpsystemdomainclocktypeindex) + "']"
    cptpsystemdomainentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpsystemdomainentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpsystemdomainentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpsystemdomainentry.EntityData.Children = make(map[string]types.YChild)
    cptpsystemdomainentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpsystemdomainentry.EntityData.Leafs["cPtpSystemDomainClockTypeIndex"] = types.YLeaf{"Cptpsystemdomainclocktypeindex", cptpsystemdomainentry.Cptpsystemdomainclocktypeindex}
    cptpsystemdomainentry.EntityData.Leafs["cPtpSystemDomainTotals"] = types.YLeaf{"Cptpsystemdomaintotals", cptpsystemdomainentry.Cptpsystemdomaintotals}
    return &(cptpsystemdomainentry.EntityData)
}

// CISCOPTPMIB_Cptpclocknodetable
// Table of information about the PTP system for a given domain.
type CISCOPTPMIB_Cptpclocknodetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single domain. A
    // entry is added when a new PTP clock domain is configured on the router. The
    // type is slice of CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry.
    Cptpclocknodeentry []CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry
}

func (cptpclocknodetable *CISCOPTPMIB_Cptpclocknodetable) GetEntityData() *types.CommonEntityData {
    cptpclocknodetable.EntityData.YFilter = cptpclocknodetable.YFilter
    cptpclocknodetable.EntityData.YangName = "cPtpClockNodeTable"
    cptpclocknodetable.EntityData.BundleName = "cisco_ios_xe"
    cptpclocknodetable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclocknodetable.EntityData.SegmentPath = "cPtpClockNodeTable"
    cptpclocknodetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclocknodetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclocknodetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclocknodetable.EntityData.Children = make(map[string]types.YChild)
    cptpclocknodetable.EntityData.Children["cPtpClockNodeEntry"] = types.YChild{"Cptpclocknodeentry", nil}
    for i := range cptpclocknodetable.Cptpclocknodeentry {
        cptpclocknodetable.EntityData.Children[types.GetSegmentPath(&cptpclocknodetable.Cptpclocknodeentry[i])] = types.YChild{"Cptpclocknodeentry", &cptpclocknodetable.Cptpclocknodeentry[i]}
    }
    cptpclocknodetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclocknodetable.EntityData)
}

// CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry
// An entry in the table, containing information about a single
// domain. A entry is added when a new PTP clock domain is
// configured on the router.
type CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry struct {
    EntityData types.CommonEntityData
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

func (cptpclocknodeentry *CISCOPTPMIB_Cptpclocknodetable_Cptpclocknodeentry) GetEntityData() *types.CommonEntityData {
    cptpclocknodeentry.EntityData.YFilter = cptpclocknodeentry.YFilter
    cptpclocknodeentry.EntityData.YangName = "cPtpClockNodeEntry"
    cptpclocknodeentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclocknodeentry.EntityData.ParentYangName = "cPtpClockNodeTable"
    cptpclocknodeentry.EntityData.SegmentPath = "cPtpClockNodeEntry" + "[cPtpClockDomainIndex='" + fmt.Sprintf("%v", cptpclocknodeentry.Cptpclockdomainindex) + "']" + "[cPtpClockTypeIndex='" + fmt.Sprintf("%v", cptpclocknodeentry.Cptpclocktypeindex) + "']" + "[cPtpClockInstanceIndex='" + fmt.Sprintf("%v", cptpclocknodeentry.Cptpclockinstanceindex) + "']"
    cptpclocknodeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclocknodeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclocknodeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclocknodeentry.EntityData.Children = make(map[string]types.YChild)
    cptpclocknodeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclocknodeentry.EntityData.Leafs["cPtpClockDomainIndex"] = types.YLeaf{"Cptpclockdomainindex", cptpclocknodeentry.Cptpclockdomainindex}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockTypeIndex"] = types.YLeaf{"Cptpclocktypeindex", cptpclocknodeentry.Cptpclocktypeindex}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockInstanceIndex"] = types.YLeaf{"Cptpclockinstanceindex", cptpclocknodeentry.Cptpclockinstanceindex}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockInput1ppsEnabled"] = types.YLeaf{"Cptpclockinput1Ppsenabled", cptpclocknodeentry.Cptpclockinput1Ppsenabled}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockInputFrequencyEnabled"] = types.YLeaf{"Cptpclockinputfrequencyenabled", cptpclocknodeentry.Cptpclockinputfrequencyenabled}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockTODEnabled"] = types.YLeaf{"Cptpclocktodenabled", cptpclocknodeentry.Cptpclocktodenabled}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockOutput1ppsEnabled"] = types.YLeaf{"Cptpclockoutput1Ppsenabled", cptpclocknodeentry.Cptpclockoutput1Ppsenabled}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockOutput1ppsOffsetEnabled"] = types.YLeaf{"Cptpclockoutput1Ppsoffsetenabled", cptpclocknodeentry.Cptpclockoutput1Ppsoffsetenabled}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockOutput1ppsOffsetValue"] = types.YLeaf{"Cptpclockoutput1Ppsoffsetvalue", cptpclocknodeentry.Cptpclockoutput1Ppsoffsetvalue}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockOutput1ppsOffsetNegative"] = types.YLeaf{"Cptpclockoutput1Ppsoffsetnegative", cptpclocknodeentry.Cptpclockoutput1Ppsoffsetnegative}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockInput1ppsInterface"] = types.YLeaf{"Cptpclockinput1Ppsinterface", cptpclocknodeentry.Cptpclockinput1Ppsinterface}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockOutput1ppsInterface"] = types.YLeaf{"Cptpclockoutput1Ppsinterface", cptpclocknodeentry.Cptpclockoutput1Ppsinterface}
    cptpclocknodeentry.EntityData.Leafs["cPtpClockTODInterface"] = types.YLeaf{"Cptpclocktodinterface", cptpclocknodeentry.Cptpclocktodinterface}
    return &(cptpclocknodeentry.EntityData)
}

// CISCOPTPMIB_Cptpclockcurrentdstable
// Table of information about the PTP clock Current Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockcurrentdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Current Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry.
    Cptpclockcurrentdsentry []CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry
}

func (cptpclockcurrentdstable *CISCOPTPMIB_Cptpclockcurrentdstable) GetEntityData() *types.CommonEntityData {
    cptpclockcurrentdstable.EntityData.YFilter = cptpclockcurrentdstable.YFilter
    cptpclockcurrentdstable.EntityData.YangName = "cPtpClockCurrentDSTable"
    cptpclockcurrentdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockcurrentdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockcurrentdstable.EntityData.SegmentPath = "cPtpClockCurrentDSTable"
    cptpclockcurrentdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockcurrentdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockcurrentdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockcurrentdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclockcurrentdstable.EntityData.Children["cPtpClockCurrentDSEntry"] = types.YChild{"Cptpclockcurrentdsentry", nil}
    for i := range cptpclockcurrentdstable.Cptpclockcurrentdsentry {
        cptpclockcurrentdstable.EntityData.Children[types.GetSegmentPath(&cptpclockcurrentdstable.Cptpclockcurrentdsentry[i])] = types.YChild{"Cptpclockcurrentdsentry", &cptpclockcurrentdstable.Cptpclockcurrentdsentry[i]}
    }
    cptpclockcurrentdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockcurrentdstable.EntityData)
}

// CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry
// An entry in the table, containing information about a single
// PTP clock Current Datasets for a domain.
type CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockcurrentdsentry *CISCOPTPMIB_Cptpclockcurrentdstable_Cptpclockcurrentdsentry) GetEntityData() *types.CommonEntityData {
    cptpclockcurrentdsentry.EntityData.YFilter = cptpclockcurrentdsentry.YFilter
    cptpclockcurrentdsentry.EntityData.YangName = "cPtpClockCurrentDSEntry"
    cptpclockcurrentdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockcurrentdsentry.EntityData.ParentYangName = "cPtpClockCurrentDSTable"
    cptpclockcurrentdsentry.EntityData.SegmentPath = "cPtpClockCurrentDSEntry" + "[cPtpClockCurrentDSDomainIndex='" + fmt.Sprintf("%v", cptpclockcurrentdsentry.Cptpclockcurrentdsdomainindex) + "']" + "[cPtpClockCurrentDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockcurrentdsentry.Cptpclockcurrentdsclocktypeindex) + "']" + "[cPtpClockCurrentDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockcurrentdsentry.Cptpclockcurrentdsinstanceindex) + "']"
    cptpclockcurrentdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockcurrentdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockcurrentdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockcurrentdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockcurrentdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockcurrentdsentry.EntityData.Leafs["cPtpClockCurrentDSDomainIndex"] = types.YLeaf{"Cptpclockcurrentdsdomainindex", cptpclockcurrentdsentry.Cptpclockcurrentdsdomainindex}
    cptpclockcurrentdsentry.EntityData.Leafs["cPtpClockCurrentDSClockTypeIndex"] = types.YLeaf{"Cptpclockcurrentdsclocktypeindex", cptpclockcurrentdsentry.Cptpclockcurrentdsclocktypeindex}
    cptpclockcurrentdsentry.EntityData.Leafs["cPtpClockCurrentDSInstanceIndex"] = types.YLeaf{"Cptpclockcurrentdsinstanceindex", cptpclockcurrentdsentry.Cptpclockcurrentdsinstanceindex}
    cptpclockcurrentdsentry.EntityData.Leafs["cPtpClockCurrentDSStepsRemoved"] = types.YLeaf{"Cptpclockcurrentdsstepsremoved", cptpclockcurrentdsentry.Cptpclockcurrentdsstepsremoved}
    cptpclockcurrentdsentry.EntityData.Leafs["cPtpClockCurrentDSOffsetFromMaster"] = types.YLeaf{"Cptpclockcurrentdsoffsetfrommaster", cptpclockcurrentdsentry.Cptpclockcurrentdsoffsetfrommaster}
    cptpclockcurrentdsentry.EntityData.Leafs["cPtpClockCurrentDSMeanPathDelay"] = types.YLeaf{"Cptpclockcurrentdsmeanpathdelay", cptpclockcurrentdsentry.Cptpclockcurrentdsmeanpathdelay}
    return &(cptpclockcurrentdsentry.EntityData)
}

// CISCOPTPMIB_Cptpclockparentdstable
// Table of information about the PTP clock Parent Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockparentdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Parent Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry.
    Cptpclockparentdsentry []CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry
}

func (cptpclockparentdstable *CISCOPTPMIB_Cptpclockparentdstable) GetEntityData() *types.CommonEntityData {
    cptpclockparentdstable.EntityData.YFilter = cptpclockparentdstable.YFilter
    cptpclockparentdstable.EntityData.YangName = "cPtpClockParentDSTable"
    cptpclockparentdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockparentdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockparentdstable.EntityData.SegmentPath = "cPtpClockParentDSTable"
    cptpclockparentdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockparentdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockparentdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockparentdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclockparentdstable.EntityData.Children["cPtpClockParentDSEntry"] = types.YChild{"Cptpclockparentdsentry", nil}
    for i := range cptpclockparentdstable.Cptpclockparentdsentry {
        cptpclockparentdstable.EntityData.Children[types.GetSegmentPath(&cptpclockparentdstable.Cptpclockparentdsentry[i])] = types.YChild{"Cptpclockparentdsentry", &cptpclockparentdstable.Cptpclockparentdsentry[i]}
    }
    cptpclockparentdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockparentdstable.EntityData)
}

// CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry
// An entry in the table, containing information about a single
// PTP clock Parent Datasets for a domain.
type CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockparentdsentry *CISCOPTPMIB_Cptpclockparentdstable_Cptpclockparentdsentry) GetEntityData() *types.CommonEntityData {
    cptpclockparentdsentry.EntityData.YFilter = cptpclockparentdsentry.YFilter
    cptpclockparentdsentry.EntityData.YangName = "cPtpClockParentDSEntry"
    cptpclockparentdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockparentdsentry.EntityData.ParentYangName = "cPtpClockParentDSTable"
    cptpclockparentdsentry.EntityData.SegmentPath = "cPtpClockParentDSEntry" + "[cPtpClockParentDSDomainIndex='" + fmt.Sprintf("%v", cptpclockparentdsentry.Cptpclockparentdsdomainindex) + "']" + "[cPtpClockParentDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockparentdsentry.Cptpclockparentdsclocktypeindex) + "']" + "[cPtpClockParentDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockparentdsentry.Cptpclockparentdsinstanceindex) + "']"
    cptpclockparentdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockparentdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockparentdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockparentdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockparentdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSDomainIndex"] = types.YLeaf{"Cptpclockparentdsdomainindex", cptpclockparentdsentry.Cptpclockparentdsdomainindex}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSClockTypeIndex"] = types.YLeaf{"Cptpclockparentdsclocktypeindex", cptpclockparentdsentry.Cptpclockparentdsclocktypeindex}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSInstanceIndex"] = types.YLeaf{"Cptpclockparentdsinstanceindex", cptpclockparentdsentry.Cptpclockparentdsinstanceindex}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSParentPortIdentity"] = types.YLeaf{"Cptpclockparentdsparentportidentity", cptpclockparentdsentry.Cptpclockparentdsparentportidentity}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSParentStats"] = types.YLeaf{"Cptpclockparentdsparentstats", cptpclockparentdsentry.Cptpclockparentdsparentstats}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSOffset"] = types.YLeaf{"Cptpclockparentdsoffset", cptpclockparentdsentry.Cptpclockparentdsoffset}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSClockPhChRate"] = types.YLeaf{"Cptpclockparentdsclockphchrate", cptpclockparentdsentry.Cptpclockparentdsclockphchrate}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSGMClockIdentity"] = types.YLeaf{"Cptpclockparentdsgmclockidentity", cptpclockparentdsentry.Cptpclockparentdsgmclockidentity}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSGMClockPriority1"] = types.YLeaf{"Cptpclockparentdsgmclockpriority1", cptpclockparentdsentry.Cptpclockparentdsgmclockpriority1}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSGMClockPriority2"] = types.YLeaf{"Cptpclockparentdsgmclockpriority2", cptpclockparentdsentry.Cptpclockparentdsgmclockpriority2}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSGMClockQualityClass"] = types.YLeaf{"Cptpclockparentdsgmclockqualityclass", cptpclockparentdsentry.Cptpclockparentdsgmclockqualityclass}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSGMClockQualityAccuracy"] = types.YLeaf{"Cptpclockparentdsgmclockqualityaccuracy", cptpclockparentdsentry.Cptpclockparentdsgmclockqualityaccuracy}
    cptpclockparentdsentry.EntityData.Leafs["cPtpClockParentDSGMClockQualityOffset"] = types.YLeaf{"Cptpclockparentdsgmclockqualityoffset", cptpclockparentdsentry.Cptpclockparentdsgmclockqualityoffset}
    return &(cptpclockparentdsentry.EntityData)
}

// CISCOPTPMIB_Cptpclockdefaultdstable
// Table of information about the PTP clock Default Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockdefaultdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Default Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry.
    Cptpclockdefaultdsentry []CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry
}

func (cptpclockdefaultdstable *CISCOPTPMIB_Cptpclockdefaultdstable) GetEntityData() *types.CommonEntityData {
    cptpclockdefaultdstable.EntityData.YFilter = cptpclockdefaultdstable.YFilter
    cptpclockdefaultdstable.EntityData.YangName = "cPtpClockDefaultDSTable"
    cptpclockdefaultdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockdefaultdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockdefaultdstable.EntityData.SegmentPath = "cPtpClockDefaultDSTable"
    cptpclockdefaultdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockdefaultdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockdefaultdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockdefaultdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclockdefaultdstable.EntityData.Children["cPtpClockDefaultDSEntry"] = types.YChild{"Cptpclockdefaultdsentry", nil}
    for i := range cptpclockdefaultdstable.Cptpclockdefaultdsentry {
        cptpclockdefaultdstable.EntityData.Children[types.GetSegmentPath(&cptpclockdefaultdstable.Cptpclockdefaultdsentry[i])] = types.YChild{"Cptpclockdefaultdsentry", &cptpclockdefaultdstable.Cptpclockdefaultdsentry[i]}
    }
    cptpclockdefaultdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockdefaultdstable.EntityData)
}

// CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry
// An entry in the table, containing information about a single
// PTP clock Default Datasets for a domain.
type CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockdefaultdsentry *CISCOPTPMIB_Cptpclockdefaultdstable_Cptpclockdefaultdsentry) GetEntityData() *types.CommonEntityData {
    cptpclockdefaultdsentry.EntityData.YFilter = cptpclockdefaultdsentry.YFilter
    cptpclockdefaultdsentry.EntityData.YangName = "cPtpClockDefaultDSEntry"
    cptpclockdefaultdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockdefaultdsentry.EntityData.ParentYangName = "cPtpClockDefaultDSTable"
    cptpclockdefaultdsentry.EntityData.SegmentPath = "cPtpClockDefaultDSEntry" + "[cPtpClockDefaultDSDomainIndex='" + fmt.Sprintf("%v", cptpclockdefaultdsentry.Cptpclockdefaultdsdomainindex) + "']" + "[cPtpClockDefaultDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockdefaultdsentry.Cptpclockdefaultdsclocktypeindex) + "']" + "[cPtpClockDefaultDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockdefaultdsentry.Cptpclockdefaultdsinstanceindex) + "']"
    cptpclockdefaultdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockdefaultdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockdefaultdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockdefaultdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockdefaultdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSDomainIndex"] = types.YLeaf{"Cptpclockdefaultdsdomainindex", cptpclockdefaultdsentry.Cptpclockdefaultdsdomainindex}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSClockTypeIndex"] = types.YLeaf{"Cptpclockdefaultdsclocktypeindex", cptpclockdefaultdsentry.Cptpclockdefaultdsclocktypeindex}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSInstanceIndex"] = types.YLeaf{"Cptpclockdefaultdsinstanceindex", cptpclockdefaultdsentry.Cptpclockdefaultdsinstanceindex}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSTwoStepFlag"] = types.YLeaf{"Cptpclockdefaultdstwostepflag", cptpclockdefaultdsentry.Cptpclockdefaultdstwostepflag}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSClockIdentity"] = types.YLeaf{"Cptpclockdefaultdsclockidentity", cptpclockdefaultdsentry.Cptpclockdefaultdsclockidentity}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSPriority1"] = types.YLeaf{"Cptpclockdefaultdspriority1", cptpclockdefaultdsentry.Cptpclockdefaultdspriority1}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSPriority2"] = types.YLeaf{"Cptpclockdefaultdspriority2", cptpclockdefaultdsentry.Cptpclockdefaultdspriority2}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSSlaveOnly"] = types.YLeaf{"Cptpclockdefaultdsslaveonly", cptpclockdefaultdsentry.Cptpclockdefaultdsslaveonly}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSQualityClass"] = types.YLeaf{"Cptpclockdefaultdsqualityclass", cptpclockdefaultdsentry.Cptpclockdefaultdsqualityclass}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSQualityAccuracy"] = types.YLeaf{"Cptpclockdefaultdsqualityaccuracy", cptpclockdefaultdsentry.Cptpclockdefaultdsqualityaccuracy}
    cptpclockdefaultdsentry.EntityData.Leafs["cPtpClockDefaultDSQualityOffset"] = types.YLeaf{"Cptpclockdefaultdsqualityoffset", cptpclockdefaultdsentry.Cptpclockdefaultdsqualityoffset}
    return &(cptpclockdefaultdsentry.EntityData)
}

// CISCOPTPMIB_Cptpclockrunningtable
// Table of information about the PTP clock Running Datasets for
// all domains.
type CISCOPTPMIB_Cptpclockrunningtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // running Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry.
    Cptpclockrunningentry []CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry
}

func (cptpclockrunningtable *CISCOPTPMIB_Cptpclockrunningtable) GetEntityData() *types.CommonEntityData {
    cptpclockrunningtable.EntityData.YFilter = cptpclockrunningtable.YFilter
    cptpclockrunningtable.EntityData.YangName = "cPtpClockRunningTable"
    cptpclockrunningtable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockrunningtable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockrunningtable.EntityData.SegmentPath = "cPtpClockRunningTable"
    cptpclockrunningtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockrunningtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockrunningtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockrunningtable.EntityData.Children = make(map[string]types.YChild)
    cptpclockrunningtable.EntityData.Children["cPtpClockRunningEntry"] = types.YChild{"Cptpclockrunningentry", nil}
    for i := range cptpclockrunningtable.Cptpclockrunningentry {
        cptpclockrunningtable.EntityData.Children[types.GetSegmentPath(&cptpclockrunningtable.Cptpclockrunningentry[i])] = types.YChild{"Cptpclockrunningentry", &cptpclockrunningtable.Cptpclockrunningentry[i]}
    }
    cptpclockrunningtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockrunningtable.EntityData)
}

// CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry
// An entry in the table, containing information about a single
// PTP clock running Datasets for a domain.
type CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockrunningentry *CISCOPTPMIB_Cptpclockrunningtable_Cptpclockrunningentry) GetEntityData() *types.CommonEntityData {
    cptpclockrunningentry.EntityData.YFilter = cptpclockrunningentry.YFilter
    cptpclockrunningentry.EntityData.YangName = "cPtpClockRunningEntry"
    cptpclockrunningentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockrunningentry.EntityData.ParentYangName = "cPtpClockRunningTable"
    cptpclockrunningentry.EntityData.SegmentPath = "cPtpClockRunningEntry" + "[cPtpClockRunningDomainIndex='" + fmt.Sprintf("%v", cptpclockrunningentry.Cptpclockrunningdomainindex) + "']" + "[cPtpClockRunningClockTypeIndex='" + fmt.Sprintf("%v", cptpclockrunningentry.Cptpclockrunningclocktypeindex) + "']" + "[cPtpClockRunningInstanceIndex='" + fmt.Sprintf("%v", cptpclockrunningentry.Cptpclockrunninginstanceindex) + "']"
    cptpclockrunningentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockrunningentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockrunningentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockrunningentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockrunningentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockrunningentry.EntityData.Leafs["cPtpClockRunningDomainIndex"] = types.YLeaf{"Cptpclockrunningdomainindex", cptpclockrunningentry.Cptpclockrunningdomainindex}
    cptpclockrunningentry.EntityData.Leafs["cPtpClockRunningClockTypeIndex"] = types.YLeaf{"Cptpclockrunningclocktypeindex", cptpclockrunningentry.Cptpclockrunningclocktypeindex}
    cptpclockrunningentry.EntityData.Leafs["cPtpClockRunningInstanceIndex"] = types.YLeaf{"Cptpclockrunninginstanceindex", cptpclockrunningentry.Cptpclockrunninginstanceindex}
    cptpclockrunningentry.EntityData.Leafs["cPtpClockRunningState"] = types.YLeaf{"Cptpclockrunningstate", cptpclockrunningentry.Cptpclockrunningstate}
    cptpclockrunningentry.EntityData.Leafs["cPtpClockRunningPacketsSent"] = types.YLeaf{"Cptpclockrunningpacketssent", cptpclockrunningentry.Cptpclockrunningpacketssent}
    cptpclockrunningentry.EntityData.Leafs["cPtpClockRunningPacketsReceived"] = types.YLeaf{"Cptpclockrunningpacketsreceived", cptpclockrunningentry.Cptpclockrunningpacketsreceived}
    return &(cptpclockrunningentry.EntityData)
}

// CISCOPTPMIB_Cptpclocktimepropertiesdstable
// Table of information about the PTP clock Timeproperties
// Datasets for all domains.
type CISCOPTPMIB_Cptpclocktimepropertiesdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // timeproperties Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry.
    Cptpclocktimepropertiesdsentry []CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry
}

func (cptpclocktimepropertiesdstable *CISCOPTPMIB_Cptpclocktimepropertiesdstable) GetEntityData() *types.CommonEntityData {
    cptpclocktimepropertiesdstable.EntityData.YFilter = cptpclocktimepropertiesdstable.YFilter
    cptpclocktimepropertiesdstable.EntityData.YangName = "cPtpClockTimePropertiesDSTable"
    cptpclocktimepropertiesdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclocktimepropertiesdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclocktimepropertiesdstable.EntityData.SegmentPath = "cPtpClockTimePropertiesDSTable"
    cptpclocktimepropertiesdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclocktimepropertiesdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclocktimepropertiesdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclocktimepropertiesdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclocktimepropertiesdstable.EntityData.Children["cPtpClockTimePropertiesDSEntry"] = types.YChild{"Cptpclocktimepropertiesdsentry", nil}
    for i := range cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry {
        cptpclocktimepropertiesdstable.EntityData.Children[types.GetSegmentPath(&cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry[i])] = types.YChild{"Cptpclocktimepropertiesdsentry", &cptpclocktimepropertiesdstable.Cptpclocktimepropertiesdsentry[i]}
    }
    cptpclocktimepropertiesdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclocktimepropertiesdstable.EntityData)
}

// CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry
// An entry in the table, containing information about a single
// PTP clock timeproperties Datasets for a domain.
type CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclocktimepropertiesdsentry *CISCOPTPMIB_Cptpclocktimepropertiesdstable_Cptpclocktimepropertiesdsentry) GetEntityData() *types.CommonEntityData {
    cptpclocktimepropertiesdsentry.EntityData.YFilter = cptpclocktimepropertiesdsentry.YFilter
    cptpclocktimepropertiesdsentry.EntityData.YangName = "cPtpClockTimePropertiesDSEntry"
    cptpclocktimepropertiesdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclocktimepropertiesdsentry.EntityData.ParentYangName = "cPtpClockTimePropertiesDSTable"
    cptpclocktimepropertiesdsentry.EntityData.SegmentPath = "cPtpClockTimePropertiesDSEntry" + "[cPtpClockTimePropertiesDSDomainIndex='" + fmt.Sprintf("%v", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsdomainindex) + "']" + "[cPtpClockTimePropertiesDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsclocktypeindex) + "']" + "[cPtpClockTimePropertiesDSInstanceIndex='" + fmt.Sprintf("%v", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsinstanceindex) + "']"
    cptpclocktimepropertiesdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclocktimepropertiesdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclocktimepropertiesdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclocktimepropertiesdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclocktimepropertiesdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSDomainIndex"] = types.YLeaf{"Cptpclocktimepropertiesdsdomainindex", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsdomainindex}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSClockTypeIndex"] = types.YLeaf{"Cptpclocktimepropertiesdsclocktypeindex", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsclocktypeindex}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSInstanceIndex"] = types.YLeaf{"Cptpclocktimepropertiesdsinstanceindex", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsinstanceindex}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSCurrentUTCOffsetValid"] = types.YLeaf{"Cptpclocktimepropertiesdscurrentutcoffsetvalid", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdscurrentutcoffsetvalid}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSCurrentUTCOffset"] = types.YLeaf{"Cptpclocktimepropertiesdscurrentutcoffset", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdscurrentutcoffset}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSLeap59"] = types.YLeaf{"Cptpclocktimepropertiesdsleap59", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsleap59}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSLeap61"] = types.YLeaf{"Cptpclocktimepropertiesdsleap61", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsleap61}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSTimeTraceable"] = types.YLeaf{"Cptpclocktimepropertiesdstimetraceable", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdstimetraceable}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSFreqTraceable"] = types.YLeaf{"Cptpclocktimepropertiesdsfreqtraceable", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsfreqtraceable}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSPTPTimescale"] = types.YLeaf{"Cptpclocktimepropertiesdsptptimescale", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdsptptimescale}
    cptpclocktimepropertiesdsentry.EntityData.Leafs["cPtpClockTimePropertiesDSSource"] = types.YLeaf{"Cptpclocktimepropertiesdssource", cptpclocktimepropertiesdsentry.Cptpclocktimepropertiesdssource}
    return &(cptpclocktimepropertiesdsentry.EntityData)
}

// CISCOPTPMIB_Cptpclocktransdefaultdstable
// Table of information about the PTP Transparent clock Default
// Datasets for all domains.
type CISCOPTPMIB_Cptpclocktransdefaultdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP
    // Transparent clock Default Datasets for a domain. The type is slice of
    // CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry.
    Cptpclocktransdefaultdsentry []CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry
}

func (cptpclocktransdefaultdstable *CISCOPTPMIB_Cptpclocktransdefaultdstable) GetEntityData() *types.CommonEntityData {
    cptpclocktransdefaultdstable.EntityData.YFilter = cptpclocktransdefaultdstable.YFilter
    cptpclocktransdefaultdstable.EntityData.YangName = "cPtpClockTransDefaultDSTable"
    cptpclocktransdefaultdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclocktransdefaultdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclocktransdefaultdstable.EntityData.SegmentPath = "cPtpClockTransDefaultDSTable"
    cptpclocktransdefaultdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclocktransdefaultdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclocktransdefaultdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclocktransdefaultdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclocktransdefaultdstable.EntityData.Children["cPtpClockTransDefaultDSEntry"] = types.YChild{"Cptpclocktransdefaultdsentry", nil}
    for i := range cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry {
        cptpclocktransdefaultdstable.EntityData.Children[types.GetSegmentPath(&cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry[i])] = types.YChild{"Cptpclocktransdefaultdsentry", &cptpclocktransdefaultdstable.Cptpclocktransdefaultdsentry[i]}
    }
    cptpclocktransdefaultdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclocktransdefaultdstable.EntityData)
}

// CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry
// An entry in the table, containing information about a single
// PTP Transparent clock Default Datasets for a domain.
type CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclocktransdefaultdsentry *CISCOPTPMIB_Cptpclocktransdefaultdstable_Cptpclocktransdefaultdsentry) GetEntityData() *types.CommonEntityData {
    cptpclocktransdefaultdsentry.EntityData.YFilter = cptpclocktransdefaultdsentry.YFilter
    cptpclocktransdefaultdsentry.EntityData.YangName = "cPtpClockTransDefaultDSEntry"
    cptpclocktransdefaultdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclocktransdefaultdsentry.EntityData.ParentYangName = "cPtpClockTransDefaultDSTable"
    cptpclocktransdefaultdsentry.EntityData.SegmentPath = "cPtpClockTransDefaultDSEntry" + "[cPtpClockTransDefaultDSDomainIndex='" + fmt.Sprintf("%v", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsdomainindex) + "']" + "[cPtpClockTransDefaultDSInstanceIndex='" + fmt.Sprintf("%v", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsinstanceindex) + "']"
    cptpclocktransdefaultdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclocktransdefaultdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclocktransdefaultdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclocktransdefaultdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclocktransdefaultdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclocktransdefaultdsentry.EntityData.Leafs["cPtpClockTransDefaultDSDomainIndex"] = types.YLeaf{"Cptpclocktransdefaultdsdomainindex", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsdomainindex}
    cptpclocktransdefaultdsentry.EntityData.Leafs["cPtpClockTransDefaultDSInstanceIndex"] = types.YLeaf{"Cptpclocktransdefaultdsinstanceindex", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsinstanceindex}
    cptpclocktransdefaultdsentry.EntityData.Leafs["cPtpClockTransDefaultDSClockIdentity"] = types.YLeaf{"Cptpclocktransdefaultdsclockidentity", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsclockidentity}
    cptpclocktransdefaultdsentry.EntityData.Leafs["cPtpClockTransDefaultDSNumOfPorts"] = types.YLeaf{"Cptpclocktransdefaultdsnumofports", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsnumofports}
    cptpclocktransdefaultdsentry.EntityData.Leafs["cPtpClockTransDefaultDSDelay"] = types.YLeaf{"Cptpclocktransdefaultdsdelay", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsdelay}
    cptpclocktransdefaultdsentry.EntityData.Leafs["cPtpClockTransDefaultDSPrimaryDomain"] = types.YLeaf{"Cptpclocktransdefaultdsprimarydomain", cptpclocktransdefaultdsentry.Cptpclocktransdefaultdsprimarydomain}
    return &(cptpclocktransdefaultdsentry.EntityData)
}

// CISCOPTPMIB_Cptpclockporttable
// Table of information about the clock ports for a particular
// domain.
type CISCOPTPMIB_Cptpclockporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single clock port.
    // The type is slice of CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry.
    Cptpclockportentry []CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry
}

func (cptpclockporttable *CISCOPTPMIB_Cptpclockporttable) GetEntityData() *types.CommonEntityData {
    cptpclockporttable.EntityData.YFilter = cptpclockporttable.YFilter
    cptpclockporttable.EntityData.YangName = "cPtpClockPortTable"
    cptpclockporttable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockporttable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockporttable.EntityData.SegmentPath = "cPtpClockPortTable"
    cptpclockporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockporttable.EntityData.Children = make(map[string]types.YChild)
    cptpclockporttable.EntityData.Children["cPtpClockPortEntry"] = types.YChild{"Cptpclockportentry", nil}
    for i := range cptpclockporttable.Cptpclockportentry {
        cptpclockporttable.EntityData.Children[types.GetSegmentPath(&cptpclockporttable.Cptpclockportentry[i])] = types.YChild{"Cptpclockportentry", &cptpclockporttable.Cptpclockportentry[i]}
    }
    cptpclockporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockporttable.EntityData)
}

// CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry
// An entry in the table, containing information about a single
// clock port.
type CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockportentry *CISCOPTPMIB_Cptpclockporttable_Cptpclockportentry) GetEntityData() *types.CommonEntityData {
    cptpclockportentry.EntityData.YFilter = cptpclockportentry.YFilter
    cptpclockportentry.EntityData.YangName = "cPtpClockPortEntry"
    cptpclockportentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportentry.EntityData.ParentYangName = "cPtpClockPortTable"
    cptpclockportentry.EntityData.SegmentPath = "cPtpClockPortEntry" + "[cPtpClockPortDomainIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockportdomainindex) + "']" + "[cPtpClockPortClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockportclocktypeindex) + "']" + "[cPtpClockPortClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockportclockinstanceindex) + "']" + "[cPtpClockPortTablePortNumberIndex='" + fmt.Sprintf("%v", cptpclockportentry.Cptpclockporttableportnumberindex) + "']"
    cptpclockportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockportentry.EntityData.Leafs["cPtpClockPortDomainIndex"] = types.YLeaf{"Cptpclockportdomainindex", cptpclockportentry.Cptpclockportdomainindex}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortClockTypeIndex"] = types.YLeaf{"Cptpclockportclocktypeindex", cptpclockportentry.Cptpclockportclocktypeindex}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortClockInstanceIndex"] = types.YLeaf{"Cptpclockportclockinstanceindex", cptpclockportentry.Cptpclockportclockinstanceindex}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortTablePortNumberIndex"] = types.YLeaf{"Cptpclockporttableportnumberindex", cptpclockportentry.Cptpclockporttableportnumberindex}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortName"] = types.YLeaf{"Cptpclockportname", cptpclockportentry.Cptpclockportname}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortRole"] = types.YLeaf{"Cptpclockportrole", cptpclockportentry.Cptpclockportrole}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortSyncOneStep"] = types.YLeaf{"Cptpclockportsynconestep", cptpclockportentry.Cptpclockportsynconestep}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortCurrentPeerAddressType"] = types.YLeaf{"Cptpclockportcurrentpeeraddresstype", cptpclockportentry.Cptpclockportcurrentpeeraddresstype}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortCurrentPeerAddress"] = types.YLeaf{"Cptpclockportcurrentpeeraddress", cptpclockportentry.Cptpclockportcurrentpeeraddress}
    cptpclockportentry.EntityData.Leafs["cPtpClockPortNumOfAssociatedPorts"] = types.YLeaf{"Cptpclockportnumofassociatedports", cptpclockportentry.Cptpclockportnumofassociatedports}
    return &(cptpclockportentry.EntityData)
}

// CISCOPTPMIB_Cptpclockportdstable
// Table of information about the clock ports dataset for a
// particular domain.
type CISCOPTPMIB_Cptpclockportdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing port dataset information for a single
    // clock port. The type is slice of
    // CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry.
    Cptpclockportdsentry []CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry
}

func (cptpclockportdstable *CISCOPTPMIB_Cptpclockportdstable) GetEntityData() *types.CommonEntityData {
    cptpclockportdstable.EntityData.YFilter = cptpclockportdstable.YFilter
    cptpclockportdstable.EntityData.YangName = "cPtpClockPortDSTable"
    cptpclockportdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockportdstable.EntityData.SegmentPath = "cPtpClockPortDSTable"
    cptpclockportdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclockportdstable.EntityData.Children["cPtpClockPortDSEntry"] = types.YChild{"Cptpclockportdsentry", nil}
    for i := range cptpclockportdstable.Cptpclockportdsentry {
        cptpclockportdstable.EntityData.Children[types.GetSegmentPath(&cptpclockportdstable.Cptpclockportdsentry[i])] = types.YChild{"Cptpclockportdsentry", &cptpclockportdstable.Cptpclockportdsentry[i]}
    }
    cptpclockportdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockportdstable.EntityData)
}

// CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry
// An entry in the table, containing port dataset information for
// a single clock port.
type CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockportdsentry *CISCOPTPMIB_Cptpclockportdstable_Cptpclockportdsentry) GetEntityData() *types.CommonEntityData {
    cptpclockportdsentry.EntityData.YFilter = cptpclockportdsentry.YFilter
    cptpclockportdsentry.EntityData.YangName = "cPtpClockPortDSEntry"
    cptpclockportdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportdsentry.EntityData.ParentYangName = "cPtpClockPortDSTable"
    cptpclockportdsentry.EntityData.SegmentPath = "cPtpClockPortDSEntry" + "[cPtpClockPortDSDomainIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsdomainindex) + "']" + "[cPtpClockPortDSClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsclocktypeindex) + "']" + "[cPtpClockPortDSClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsclockinstanceindex) + "']" + "[cPtpClockPortDSPortNumberIndex='" + fmt.Sprintf("%v", cptpclockportdsentry.Cptpclockportdsportnumberindex) + "']"
    cptpclockportdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockportdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSDomainIndex"] = types.YLeaf{"Cptpclockportdsdomainindex", cptpclockportdsentry.Cptpclockportdsdomainindex}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSClockTypeIndex"] = types.YLeaf{"Cptpclockportdsclocktypeindex", cptpclockportdsentry.Cptpclockportdsclocktypeindex}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSClockInstanceIndex"] = types.YLeaf{"Cptpclockportdsclockinstanceindex", cptpclockportdsentry.Cptpclockportdsclockinstanceindex}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSPortNumberIndex"] = types.YLeaf{"Cptpclockportdsportnumberindex", cptpclockportdsentry.Cptpclockportdsportnumberindex}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSName"] = types.YLeaf{"Cptpclockportdsname", cptpclockportdsentry.Cptpclockportdsname}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSPortIdentity"] = types.YLeaf{"Cptpclockportdsportidentity", cptpclockportdsentry.Cptpclockportdsportidentity}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSAnnouncementInterval"] = types.YLeaf{"Cptpclockportdsannouncementinterval", cptpclockportdsentry.Cptpclockportdsannouncementinterval}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSAnnounceRctTimeout"] = types.YLeaf{"Cptpclockportdsannouncercttimeout", cptpclockportdsentry.Cptpclockportdsannouncercttimeout}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSSyncInterval"] = types.YLeaf{"Cptpclockportdssyncinterval", cptpclockportdsentry.Cptpclockportdssyncinterval}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSMinDelayReqInterval"] = types.YLeaf{"Cptpclockportdsmindelayreqinterval", cptpclockportdsentry.Cptpclockportdsmindelayreqinterval}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSPeerDelayReqInterval"] = types.YLeaf{"Cptpclockportdspeerdelayreqinterval", cptpclockportdsentry.Cptpclockportdspeerdelayreqinterval}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSDelayMech"] = types.YLeaf{"Cptpclockportdsdelaymech", cptpclockportdsentry.Cptpclockportdsdelaymech}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSPeerMeanPathDelay"] = types.YLeaf{"Cptpclockportdspeermeanpathdelay", cptpclockportdsentry.Cptpclockportdspeermeanpathdelay}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSGrantDuration"] = types.YLeaf{"Cptpclockportdsgrantduration", cptpclockportdsentry.Cptpclockportdsgrantduration}
    cptpclockportdsentry.EntityData.Leafs["cPtpClockPortDSPTPVersion"] = types.YLeaf{"Cptpclockportdsptpversion", cptpclockportdsentry.Cptpclockportdsptpversion}
    return &(cptpclockportdsentry.EntityData)
}

// CISCOPTPMIB_Cptpclockportrunningtable
// Table of information about the clock ports running dataset for
// a particular domain.
type CISCOPTPMIB_Cptpclockportrunningtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing runing dataset information about a single
    // clock port. The type is slice of
    // CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry.
    Cptpclockportrunningentry []CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry
}

func (cptpclockportrunningtable *CISCOPTPMIB_Cptpclockportrunningtable) GetEntityData() *types.CommonEntityData {
    cptpclockportrunningtable.EntityData.YFilter = cptpclockportrunningtable.YFilter
    cptpclockportrunningtable.EntityData.YangName = "cPtpClockPortRunningTable"
    cptpclockportrunningtable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportrunningtable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockportrunningtable.EntityData.SegmentPath = "cPtpClockPortRunningTable"
    cptpclockportrunningtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportrunningtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportrunningtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportrunningtable.EntityData.Children = make(map[string]types.YChild)
    cptpclockportrunningtable.EntityData.Children["cPtpClockPortRunningEntry"] = types.YChild{"Cptpclockportrunningentry", nil}
    for i := range cptpclockportrunningtable.Cptpclockportrunningentry {
        cptpclockportrunningtable.EntityData.Children[types.GetSegmentPath(&cptpclockportrunningtable.Cptpclockportrunningentry[i])] = types.YChild{"Cptpclockportrunningentry", &cptpclockportrunningtable.Cptpclockportrunningentry[i]}
    }
    cptpclockportrunningtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockportrunningtable.EntityData)
}

// CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry
// An entry in the table, containing runing dataset information
// about a single clock port.
type CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockportrunningentry *CISCOPTPMIB_Cptpclockportrunningtable_Cptpclockportrunningentry) GetEntityData() *types.CommonEntityData {
    cptpclockportrunningentry.EntityData.YFilter = cptpclockportrunningentry.YFilter
    cptpclockportrunningentry.EntityData.YangName = "cPtpClockPortRunningEntry"
    cptpclockportrunningentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportrunningentry.EntityData.ParentYangName = "cPtpClockPortRunningTable"
    cptpclockportrunningentry.EntityData.SegmentPath = "cPtpClockPortRunningEntry" + "[cPtpClockPortRunningDomainIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningdomainindex) + "']" + "[cPtpClockPortRunningClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningclocktypeindex) + "']" + "[cPtpClockPortRunningClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningclockinstanceindex) + "']" + "[cPtpClockPortRunningPortNumberIndex='" + fmt.Sprintf("%v", cptpclockportrunningentry.Cptpclockportrunningportnumberindex) + "']"
    cptpclockportrunningentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportrunningentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportrunningentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportrunningentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockportrunningentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningDomainIndex"] = types.YLeaf{"Cptpclockportrunningdomainindex", cptpclockportrunningentry.Cptpclockportrunningdomainindex}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningClockTypeIndex"] = types.YLeaf{"Cptpclockportrunningclocktypeindex", cptpclockportrunningentry.Cptpclockportrunningclocktypeindex}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningClockInstanceIndex"] = types.YLeaf{"Cptpclockportrunningclockinstanceindex", cptpclockportrunningentry.Cptpclockportrunningclockinstanceindex}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningPortNumberIndex"] = types.YLeaf{"Cptpclockportrunningportnumberindex", cptpclockportrunningentry.Cptpclockportrunningportnumberindex}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningName"] = types.YLeaf{"Cptpclockportrunningname", cptpclockportrunningentry.Cptpclockportrunningname}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningState"] = types.YLeaf{"Cptpclockportrunningstate", cptpclockportrunningentry.Cptpclockportrunningstate}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningRole"] = types.YLeaf{"Cptpclockportrunningrole", cptpclockportrunningentry.Cptpclockportrunningrole}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningInterfaceIndex"] = types.YLeaf{"Cptpclockportrunninginterfaceindex", cptpclockportrunningentry.Cptpclockportrunninginterfaceindex}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningIPversion"] = types.YLeaf{"Cptpclockportrunningipversion", cptpclockportrunningentry.Cptpclockportrunningipversion}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningEncapsulationType"] = types.YLeaf{"Cptpclockportrunningencapsulationtype", cptpclockportrunningentry.Cptpclockportrunningencapsulationtype}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningTxMode"] = types.YLeaf{"Cptpclockportrunningtxmode", cptpclockportrunningentry.Cptpclockportrunningtxmode}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningRxMode"] = types.YLeaf{"Cptpclockportrunningrxmode", cptpclockportrunningentry.Cptpclockportrunningrxmode}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningPacketsReceived"] = types.YLeaf{"Cptpclockportrunningpacketsreceived", cptpclockportrunningentry.Cptpclockportrunningpacketsreceived}
    cptpclockportrunningentry.EntityData.Leafs["cPtpClockPortRunningPacketsSent"] = types.YLeaf{"Cptpclockportrunningpacketssent", cptpclockportrunningentry.Cptpclockportrunningpacketssent}
    return &(cptpclockportrunningentry.EntityData)
}

// CISCOPTPMIB_Cptpclockporttransdstable
// Table of information about the Transparent clock ports running
// dataset for a particular domain.
type CISCOPTPMIB_Cptpclockporttransdstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing clock port Transparent dataset
    // information about a single clock port. The type is slice of
    // CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry.
    Cptpclockporttransdsentry []CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry
}

func (cptpclockporttransdstable *CISCOPTPMIB_Cptpclockporttransdstable) GetEntityData() *types.CommonEntityData {
    cptpclockporttransdstable.EntityData.YFilter = cptpclockporttransdstable.YFilter
    cptpclockporttransdstable.EntityData.YangName = "cPtpClockPortTransDSTable"
    cptpclockporttransdstable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockporttransdstable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockporttransdstable.EntityData.SegmentPath = "cPtpClockPortTransDSTable"
    cptpclockporttransdstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockporttransdstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockporttransdstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockporttransdstable.EntityData.Children = make(map[string]types.YChild)
    cptpclockporttransdstable.EntityData.Children["cPtpClockPortTransDSEntry"] = types.YChild{"Cptpclockporttransdsentry", nil}
    for i := range cptpclockporttransdstable.Cptpclockporttransdsentry {
        cptpclockporttransdstable.EntityData.Children[types.GetSegmentPath(&cptpclockporttransdstable.Cptpclockporttransdsentry[i])] = types.YChild{"Cptpclockporttransdsentry", &cptpclockporttransdstable.Cptpclockporttransdsentry[i]}
    }
    cptpclockporttransdstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockporttransdstable.EntityData)
}

// CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry
// An entry in the table, containing clock port Transparent
// dataset information about a single clock port
type CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockporttransdsentry *CISCOPTPMIB_Cptpclockporttransdstable_Cptpclockporttransdsentry) GetEntityData() *types.CommonEntityData {
    cptpclockporttransdsentry.EntityData.YFilter = cptpclockporttransdsentry.YFilter
    cptpclockporttransdsentry.EntityData.YangName = "cPtpClockPortTransDSEntry"
    cptpclockporttransdsentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockporttransdsentry.EntityData.ParentYangName = "cPtpClockPortTransDSTable"
    cptpclockporttransdsentry.EntityData.SegmentPath = "cPtpClockPortTransDSEntry" + "[cPtpClockPortTransDSDomainIndex='" + fmt.Sprintf("%v", cptpclockporttransdsentry.Cptpclockporttransdsdomainindex) + "']" + "[cPtpClockPortTransDSInstanceIndex='" + fmt.Sprintf("%v", cptpclockporttransdsentry.Cptpclockporttransdsinstanceindex) + "']" + "[cPtpClockPortTransDSPortNumberIndex='" + fmt.Sprintf("%v", cptpclockporttransdsentry.Cptpclockporttransdsportnumberindex) + "']"
    cptpclockporttransdsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockporttransdsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockporttransdsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockporttransdsentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockporttransdsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSDomainIndex"] = types.YLeaf{"Cptpclockporttransdsdomainindex", cptpclockporttransdsentry.Cptpclockporttransdsdomainindex}
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSInstanceIndex"] = types.YLeaf{"Cptpclockporttransdsinstanceindex", cptpclockporttransdsentry.Cptpclockporttransdsinstanceindex}
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSPortNumberIndex"] = types.YLeaf{"Cptpclockporttransdsportnumberindex", cptpclockporttransdsentry.Cptpclockporttransdsportnumberindex}
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSPortIdentity"] = types.YLeaf{"Cptpclockporttransdsportidentity", cptpclockporttransdsentry.Cptpclockporttransdsportidentity}
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSlogMinPdelayReqInt"] = types.YLeaf{"Cptpclockporttransdslogminpdelayreqint", cptpclockporttransdsentry.Cptpclockporttransdslogminpdelayreqint}
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSFaultyFlag"] = types.YLeaf{"Cptpclockporttransdsfaultyflag", cptpclockporttransdsentry.Cptpclockporttransdsfaultyflag}
    cptpclockporttransdsentry.EntityData.Leafs["cPtpClockPortTransDSPeerMeanPathDelay"] = types.YLeaf{"Cptpclockporttransdspeermeanpathdelay", cptpclockporttransdsentry.Cptpclockporttransdspeermeanpathdelay}
    return &(cptpclockporttransdsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single associated
    // port for the given clockport. The type is slice of
    // CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry.
    Cptpclockportassociateentry []CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry
}

func (cptpclockportassociatetable *CISCOPTPMIB_Cptpclockportassociatetable) GetEntityData() *types.CommonEntityData {
    cptpclockportassociatetable.EntityData.YFilter = cptpclockportassociatetable.YFilter
    cptpclockportassociatetable.EntityData.YangName = "cPtpClockPortAssociateTable"
    cptpclockportassociatetable.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportassociatetable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cptpclockportassociatetable.EntityData.SegmentPath = "cPtpClockPortAssociateTable"
    cptpclockportassociatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportassociatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportassociatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportassociatetable.EntityData.Children = make(map[string]types.YChild)
    cptpclockportassociatetable.EntityData.Children["cPtpClockPortAssociateEntry"] = types.YChild{"Cptpclockportassociateentry", nil}
    for i := range cptpclockportassociatetable.Cptpclockportassociateentry {
        cptpclockportassociatetable.EntityData.Children[types.GetSegmentPath(&cptpclockportassociatetable.Cptpclockportassociateentry[i])] = types.YChild{"Cptpclockportassociateentry", &cptpclockportassociatetable.Cptpclockportassociateentry[i]}
    }
    cptpclockportassociatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cptpclockportassociatetable.EntityData)
}

// CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry
// An entry in the table, containing information about a single
// associated port for the given clockport.
type CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry struct {
    EntityData types.CommonEntityData
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

func (cptpclockportassociateentry *CISCOPTPMIB_Cptpclockportassociatetable_Cptpclockportassociateentry) GetEntityData() *types.CommonEntityData {
    cptpclockportassociateentry.EntityData.YFilter = cptpclockportassociateentry.YFilter
    cptpclockportassociateentry.EntityData.YangName = "cPtpClockPortAssociateEntry"
    cptpclockportassociateentry.EntityData.BundleName = "cisco_ios_xe"
    cptpclockportassociateentry.EntityData.ParentYangName = "cPtpClockPortAssociateTable"
    cptpclockportassociateentry.EntityData.SegmentPath = "cPtpClockPortAssociateEntry" + "[cPtpClockPortCurrentDomainIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentdomainindex) + "']" + "[cPtpClockPortCurrentClockTypeIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentclocktypeindex) + "']" + "[cPtpClockPortCurrentClockInstanceIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentclockinstanceindex) + "']" + "[cPtpClockPortCurrentPortNumberIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportcurrentportnumberindex) + "']" + "[cPtpClockPortAssociatePortIndex='" + fmt.Sprintf("%v", cptpclockportassociateentry.Cptpclockportassociateportindex) + "']"
    cptpclockportassociateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cptpclockportassociateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cptpclockportassociateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cptpclockportassociateentry.EntityData.Children = make(map[string]types.YChild)
    cptpclockportassociateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortCurrentDomainIndex"] = types.YLeaf{"Cptpclockportcurrentdomainindex", cptpclockportassociateentry.Cptpclockportcurrentdomainindex}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortCurrentClockTypeIndex"] = types.YLeaf{"Cptpclockportcurrentclocktypeindex", cptpclockportassociateentry.Cptpclockportcurrentclocktypeindex}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortCurrentClockInstanceIndex"] = types.YLeaf{"Cptpclockportcurrentclockinstanceindex", cptpclockportassociateentry.Cptpclockportcurrentclockinstanceindex}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortCurrentPortNumberIndex"] = types.YLeaf{"Cptpclockportcurrentportnumberindex", cptpclockportassociateentry.Cptpclockportcurrentportnumberindex}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociatePortIndex"] = types.YLeaf{"Cptpclockportassociateportindex", cptpclockportassociateentry.Cptpclockportassociateportindex}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociateAddressType"] = types.YLeaf{"Cptpclockportassociateaddresstype", cptpclockportassociateentry.Cptpclockportassociateaddresstype}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociateAddress"] = types.YLeaf{"Cptpclockportassociateaddress", cptpclockportassociateentry.Cptpclockportassociateaddress}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociatePacketsSent"] = types.YLeaf{"Cptpclockportassociatepacketssent", cptpclockportassociateentry.Cptpclockportassociatepacketssent}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociatePacketsReceived"] = types.YLeaf{"Cptpclockportassociatepacketsreceived", cptpclockportassociateentry.Cptpclockportassociatepacketsreceived}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociateInErrors"] = types.YLeaf{"Cptpclockportassociateinerrors", cptpclockportassociateentry.Cptpclockportassociateinerrors}
    cptpclockportassociateentry.EntityData.Leafs["cPtpClockPortAssociateOutErrors"] = types.YLeaf{"Cptpclockportassociateouterrors", cptpclockportassociateentry.Cptpclockportassociateouterrors}
    return &(cptpclockportassociateentry.EntityData)
}

