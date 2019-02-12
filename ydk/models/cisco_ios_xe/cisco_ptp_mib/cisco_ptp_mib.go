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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoPtpMIBSystemInfo CISCOPTPMIB_CiscoPtpMIBSystemInfo

    // Table of count information about the PTP system for all domains.
    CPtpSystemTable CISCOPTPMIB_CPtpSystemTable

    // Table of information about the PTP system for all clock modes -- ordinary,
    // boundary or transparent.
    CPtpSystemDomainTable CISCOPTPMIB_CPtpSystemDomainTable

    // Table of information about the PTP system for a given domain.
    CPtpClockNodeTable CISCOPTPMIB_CPtpClockNodeTable

    // Table of information about the PTP clock Current Datasets for all domains.
    CPtpClockCurrentDSTable CISCOPTPMIB_CPtpClockCurrentDSTable

    // Table of information about the PTP clock Parent Datasets for all domains.
    CPtpClockParentDSTable CISCOPTPMIB_CPtpClockParentDSTable

    // Table of information about the PTP clock Default Datasets for all domains.
    CPtpClockDefaultDSTable CISCOPTPMIB_CPtpClockDefaultDSTable

    // Table of information about the PTP clock Running Datasets for all domains.
    CPtpClockRunningTable CISCOPTPMIB_CPtpClockRunningTable

    // Table of information about the PTP clock Timeproperties Datasets for all
    // domains.
    CPtpClockTimePropertiesDSTable CISCOPTPMIB_CPtpClockTimePropertiesDSTable

    // Table of information about the PTP Transparent clock Default Datasets for
    // all domains.
    CPtpClockTransDefaultDSTable CISCOPTPMIB_CPtpClockTransDefaultDSTable

    // Table of information about the clock ports for a particular domain.
    CPtpClockPortTable CISCOPTPMIB_CPtpClockPortTable

    // Table of information about the clock ports dataset for a particular domain.
    CPtpClockPortDSTable CISCOPTPMIB_CPtpClockPortDSTable

    // Table of information about the clock ports running dataset for a particular
    // domain.
    CPtpClockPortRunningTable CISCOPTPMIB_CPtpClockPortRunningTable

    // Table of information about the Transparent clock ports running dataset for
    // a particular domain.
    CPtpClockPortTransDSTable CISCOPTPMIB_CPtpClockPortTransDSTable

    // Table of information about a given port's associated ports.  For a master
    // port - multiple slave ports which have established sessions with the
    // current master port.   For a slave port - the list of masters available for
    // a given slave port.   Session information (pkts, errors) to be displayed
    // based on availability and scenario.
    CPtpClockPortAssociateTable CISCOPTPMIB_CPtpClockPortAssociateTable
}

func (cISCOPTPMIB *CISCOPTPMIB) GetEntityData() *types.CommonEntityData {
    cISCOPTPMIB.EntityData.YFilter = cISCOPTPMIB.YFilter
    cISCOPTPMIB.EntityData.YangName = "CISCO-PTP-MIB"
    cISCOPTPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPTPMIB.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cISCOPTPMIB.EntityData.SegmentPath = "CISCO-PTP-MIB:CISCO-PTP-MIB"
    cISCOPTPMIB.EntityData.AbsolutePath = cISCOPTPMIB.EntityData.SegmentPath
    cISCOPTPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPTPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPTPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPTPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOPTPMIB.EntityData.Children.Append("ciscoPtpMIBSystemInfo", types.YChild{"CiscoPtpMIBSystemInfo", &cISCOPTPMIB.CiscoPtpMIBSystemInfo})
    cISCOPTPMIB.EntityData.Children.Append("cPtpSystemTable", types.YChild{"CPtpSystemTable", &cISCOPTPMIB.CPtpSystemTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpSystemDomainTable", types.YChild{"CPtpSystemDomainTable", &cISCOPTPMIB.CPtpSystemDomainTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockNodeTable", types.YChild{"CPtpClockNodeTable", &cISCOPTPMIB.CPtpClockNodeTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockCurrentDSTable", types.YChild{"CPtpClockCurrentDSTable", &cISCOPTPMIB.CPtpClockCurrentDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockParentDSTable", types.YChild{"CPtpClockParentDSTable", &cISCOPTPMIB.CPtpClockParentDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockDefaultDSTable", types.YChild{"CPtpClockDefaultDSTable", &cISCOPTPMIB.CPtpClockDefaultDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockRunningTable", types.YChild{"CPtpClockRunningTable", &cISCOPTPMIB.CPtpClockRunningTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockTimePropertiesDSTable", types.YChild{"CPtpClockTimePropertiesDSTable", &cISCOPTPMIB.CPtpClockTimePropertiesDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockTransDefaultDSTable", types.YChild{"CPtpClockTransDefaultDSTable", &cISCOPTPMIB.CPtpClockTransDefaultDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockPortTable", types.YChild{"CPtpClockPortTable", &cISCOPTPMIB.CPtpClockPortTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockPortDSTable", types.YChild{"CPtpClockPortDSTable", &cISCOPTPMIB.CPtpClockPortDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockPortRunningTable", types.YChild{"CPtpClockPortRunningTable", &cISCOPTPMIB.CPtpClockPortRunningTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockPortTransDSTable", types.YChild{"CPtpClockPortTransDSTable", &cISCOPTPMIB.CPtpClockPortTransDSTable})
    cISCOPTPMIB.EntityData.Children.Append("cPtpClockPortAssociateTable", types.YChild{"CPtpClockPortAssociateTable", &cISCOPTPMIB.CPtpClockPortAssociateTable})
    cISCOPTPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOPTPMIB.EntityData.YListKeys = []string {}

    return &(cISCOPTPMIB.EntityData)
}

// CISCOPTPMIB_CiscoPtpMIBSystemInfo
type CISCOPTPMIB_CiscoPtpMIBSystemInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the PTP Profile implemented on the system. The type
    // is ClockProfileType.
    CPtpSystemProfile interface{}
}

func (ciscoPtpMIBSystemInfo *CISCOPTPMIB_CiscoPtpMIBSystemInfo) GetEntityData() *types.CommonEntityData {
    ciscoPtpMIBSystemInfo.EntityData.YFilter = ciscoPtpMIBSystemInfo.YFilter
    ciscoPtpMIBSystemInfo.EntityData.YangName = "ciscoPtpMIBSystemInfo"
    ciscoPtpMIBSystemInfo.EntityData.BundleName = "cisco_ios_xe"
    ciscoPtpMIBSystemInfo.EntityData.ParentYangName = "CISCO-PTP-MIB"
    ciscoPtpMIBSystemInfo.EntityData.SegmentPath = "ciscoPtpMIBSystemInfo"
    ciscoPtpMIBSystemInfo.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + ciscoPtpMIBSystemInfo.EntityData.SegmentPath
    ciscoPtpMIBSystemInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoPtpMIBSystemInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoPtpMIBSystemInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoPtpMIBSystemInfo.EntityData.Children = types.NewOrderedMap()
    ciscoPtpMIBSystemInfo.EntityData.Leafs = types.NewOrderedMap()
    ciscoPtpMIBSystemInfo.EntityData.Leafs.Append("cPtpSystemProfile", types.YLeaf{"CPtpSystemProfile", ciscoPtpMIBSystemInfo.CPtpSystemProfile})

    ciscoPtpMIBSystemInfo.EntityData.YListKeys = []string {}

    return &(ciscoPtpMIBSystemInfo.EntityData)
}

// CISCOPTPMIB_CPtpSystemTable
// Table of count information about the PTP system for all
// domains.
type CISCOPTPMIB_CPtpSystemTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing count information about a single domain.
    // New row entries are added when the PTP clock for this domain is configured,
    // while the unconfiguration of the PTP clock removes it. The type is slice of
    // CISCOPTPMIB_CPtpSystemTable_CPtpSystemEntry.
    CPtpSystemEntry []*CISCOPTPMIB_CPtpSystemTable_CPtpSystemEntry
}

func (cPtpSystemTable *CISCOPTPMIB_CPtpSystemTable) GetEntityData() *types.CommonEntityData {
    cPtpSystemTable.EntityData.YFilter = cPtpSystemTable.YFilter
    cPtpSystemTable.EntityData.YangName = "cPtpSystemTable"
    cPtpSystemTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpSystemTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpSystemTable.EntityData.SegmentPath = "cPtpSystemTable"
    cPtpSystemTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpSystemTable.EntityData.SegmentPath
    cPtpSystemTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpSystemTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpSystemTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpSystemTable.EntityData.Children = types.NewOrderedMap()
    cPtpSystemTable.EntityData.Children.Append("cPtpSystemEntry", types.YChild{"CPtpSystemEntry", nil})
    for i := range cPtpSystemTable.CPtpSystemEntry {
        cPtpSystemTable.EntityData.Children.Append(types.GetSegmentPath(cPtpSystemTable.CPtpSystemEntry[i]), types.YChild{"CPtpSystemEntry", cPtpSystemTable.CPtpSystemEntry[i]})
    }
    cPtpSystemTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpSystemTable.EntityData.YListKeys = []string {}

    return &(cPtpSystemTable.EntityData)
}

// CISCOPTPMIB_CPtpSystemTable_CPtpSystemEntry
// An entry in the table, containing count information about a
// single domain. New row entries are added when the PTP clock for
// this domain is configured, while the unconfiguration of the PTP
// clock removes it.
type CISCOPTPMIB_CPtpSystemTable_CPtpSystemEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The Clock Domain is a logical group of
    // clocks and devices that synchronize with each other using the PTP protocol.
    // 0           Default domain 1           Alternate domain 1 2          
    // Alternate domain 2 3           Alternate domain 3 4 - 127     User-defined
    // domains 128 - 255   Reserved. The type is interface{} with range: 0..255.
    CPtpDomainIndex interface{}

    // This attribute is a key. This object specifies the instance of the Clock
    // for this domain. The type is interface{} with range: 0..255.
    CPtpInstanceIndex interface{}

    // This object specifies the total number of clock ports configured within a
    // domain. The type is interface{} with range: 0..4294967295. Units are ptp
    // ports.
    CPtpDomainClockPortsTotal interface{}

    // This object specifies the total number of clock port Physical interfaces
    // configured within a domain instance for PTP communications. The type is
    // interface{} with range: 0..4294967295. Units are physical ports.
    CPtpDomainClockPortPhysicalInterfacesTotal interface{}
}

func (cPtpSystemEntry *CISCOPTPMIB_CPtpSystemTable_CPtpSystemEntry) GetEntityData() *types.CommonEntityData {
    cPtpSystemEntry.EntityData.YFilter = cPtpSystemEntry.YFilter
    cPtpSystemEntry.EntityData.YangName = "cPtpSystemEntry"
    cPtpSystemEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpSystemEntry.EntityData.ParentYangName = "cPtpSystemTable"
    cPtpSystemEntry.EntityData.SegmentPath = "cPtpSystemEntry" + types.AddKeyToken(cPtpSystemEntry.CPtpDomainIndex, "cPtpDomainIndex") + types.AddKeyToken(cPtpSystemEntry.CPtpInstanceIndex, "cPtpInstanceIndex")
    cPtpSystemEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpSystemTable/" + cPtpSystemEntry.EntityData.SegmentPath
    cPtpSystemEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpSystemEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpSystemEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpSystemEntry.EntityData.Children = types.NewOrderedMap()
    cPtpSystemEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpSystemEntry.EntityData.Leafs.Append("cPtpDomainIndex", types.YLeaf{"CPtpDomainIndex", cPtpSystemEntry.CPtpDomainIndex})
    cPtpSystemEntry.EntityData.Leafs.Append("cPtpInstanceIndex", types.YLeaf{"CPtpInstanceIndex", cPtpSystemEntry.CPtpInstanceIndex})
    cPtpSystemEntry.EntityData.Leafs.Append("cPtpDomainClockPortsTotal", types.YLeaf{"CPtpDomainClockPortsTotal", cPtpSystemEntry.CPtpDomainClockPortsTotal})
    cPtpSystemEntry.EntityData.Leafs.Append("cPtpDomainClockPortPhysicalInterfacesTotal", types.YLeaf{"CPtpDomainClockPortPhysicalInterfacesTotal", cPtpSystemEntry.CPtpDomainClockPortPhysicalInterfacesTotal})

    cPtpSystemEntry.EntityData.YListKeys = []string {"CPtpDomainIndex", "CPtpInstanceIndex"}

    return &(cPtpSystemEntry.EntityData)
}

// CISCOPTPMIB_CPtpSystemDomainTable
// Table of information about the PTP system for all clock modes
// -- ordinary, boundary or transparent.
type CISCOPTPMIB_CPtpSystemDomainTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single clock mode for
    // the PTP system. A row entry gets added when PTP clocks are configured on
    // the router. The type is slice of
    // CISCOPTPMIB_CPtpSystemDomainTable_CPtpSystemDomainEntry.
    CPtpSystemDomainEntry []*CISCOPTPMIB_CPtpSystemDomainTable_CPtpSystemDomainEntry
}

func (cPtpSystemDomainTable *CISCOPTPMIB_CPtpSystemDomainTable) GetEntityData() *types.CommonEntityData {
    cPtpSystemDomainTable.EntityData.YFilter = cPtpSystemDomainTable.YFilter
    cPtpSystemDomainTable.EntityData.YangName = "cPtpSystemDomainTable"
    cPtpSystemDomainTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpSystemDomainTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpSystemDomainTable.EntityData.SegmentPath = "cPtpSystemDomainTable"
    cPtpSystemDomainTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpSystemDomainTable.EntityData.SegmentPath
    cPtpSystemDomainTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpSystemDomainTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpSystemDomainTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpSystemDomainTable.EntityData.Children = types.NewOrderedMap()
    cPtpSystemDomainTable.EntityData.Children.Append("cPtpSystemDomainEntry", types.YChild{"CPtpSystemDomainEntry", nil})
    for i := range cPtpSystemDomainTable.CPtpSystemDomainEntry {
        cPtpSystemDomainTable.EntityData.Children.Append(types.GetSegmentPath(cPtpSystemDomainTable.CPtpSystemDomainEntry[i]), types.YChild{"CPtpSystemDomainEntry", cPtpSystemDomainTable.CPtpSystemDomainEntry[i]})
    }
    cPtpSystemDomainTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpSystemDomainTable.EntityData.YListKeys = []string {}

    return &(cPtpSystemDomainTable.EntityData)
}

// CISCOPTPMIB_CPtpSystemDomainTable_CPtpSystemDomainEntry
// An entry in the table, containing information about a single
// clock mode for the PTP system. A row entry gets added when PTP
// clocks are configured on the router.
type CISCOPTPMIB_CPtpSystemDomainTable_CPtpSystemDomainEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpSystemDomainClockTypeIndex interface{}

    // This object specifies the  total number of PTP domains for this particular
    // clock type configured in this node. The type is interface{} with range:
    // 0..4294967295. Units are domains.
    CPtpSystemDomainTotals interface{}
}

func (cPtpSystemDomainEntry *CISCOPTPMIB_CPtpSystemDomainTable_CPtpSystemDomainEntry) GetEntityData() *types.CommonEntityData {
    cPtpSystemDomainEntry.EntityData.YFilter = cPtpSystemDomainEntry.YFilter
    cPtpSystemDomainEntry.EntityData.YangName = "cPtpSystemDomainEntry"
    cPtpSystemDomainEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpSystemDomainEntry.EntityData.ParentYangName = "cPtpSystemDomainTable"
    cPtpSystemDomainEntry.EntityData.SegmentPath = "cPtpSystemDomainEntry" + types.AddKeyToken(cPtpSystemDomainEntry.CPtpSystemDomainClockTypeIndex, "cPtpSystemDomainClockTypeIndex")
    cPtpSystemDomainEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpSystemDomainTable/" + cPtpSystemDomainEntry.EntityData.SegmentPath
    cPtpSystemDomainEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpSystemDomainEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpSystemDomainEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpSystemDomainEntry.EntityData.Children = types.NewOrderedMap()
    cPtpSystemDomainEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpSystemDomainEntry.EntityData.Leafs.Append("cPtpSystemDomainClockTypeIndex", types.YLeaf{"CPtpSystemDomainClockTypeIndex", cPtpSystemDomainEntry.CPtpSystemDomainClockTypeIndex})
    cPtpSystemDomainEntry.EntityData.Leafs.Append("cPtpSystemDomainTotals", types.YLeaf{"CPtpSystemDomainTotals", cPtpSystemDomainEntry.CPtpSystemDomainTotals})

    cPtpSystemDomainEntry.EntityData.YListKeys = []string {"CPtpSystemDomainClockTypeIndex"}

    return &(cPtpSystemDomainEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockNodeTable
// Table of information about the PTP system for a given domain.
type CISCOPTPMIB_CPtpClockNodeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single domain. A
    // entry is added when a new PTP clock domain is configured on the router. The
    // type is slice of CISCOPTPMIB_CPtpClockNodeTable_CPtpClockNodeEntry.
    CPtpClockNodeEntry []*CISCOPTPMIB_CPtpClockNodeTable_CPtpClockNodeEntry
}

func (cPtpClockNodeTable *CISCOPTPMIB_CPtpClockNodeTable) GetEntityData() *types.CommonEntityData {
    cPtpClockNodeTable.EntityData.YFilter = cPtpClockNodeTable.YFilter
    cPtpClockNodeTable.EntityData.YangName = "cPtpClockNodeTable"
    cPtpClockNodeTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockNodeTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockNodeTable.EntityData.SegmentPath = "cPtpClockNodeTable"
    cPtpClockNodeTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockNodeTable.EntityData.SegmentPath
    cPtpClockNodeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockNodeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockNodeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockNodeTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockNodeTable.EntityData.Children.Append("cPtpClockNodeEntry", types.YChild{"CPtpClockNodeEntry", nil})
    for i := range cPtpClockNodeTable.CPtpClockNodeEntry {
        cPtpClockNodeTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockNodeTable.CPtpClockNodeEntry[i]), types.YChild{"CPtpClockNodeEntry", cPtpClockNodeTable.CPtpClockNodeEntry[i]})
    }
    cPtpClockNodeTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockNodeTable.EntityData.YListKeys = []string {}

    return &(cPtpClockNodeTable.EntityData)
}

// CISCOPTPMIB_CPtpClockNodeTable_CPtpClockNodeEntry
// An entry in the table, containing information about a single
// domain. A entry is added when a new PTP clock domain is
// configured on the router.
type CISCOPTPMIB_CPtpClockNodeTable_CPtpClockNodeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the  domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the Clock
    // for this clock type for the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockInstanceIndex interface{}

    // This object specifies whether the node is enabled for PTP input clocking
    // using the 1pps interface. The type is bool.
    CPtpClockInput1ppsEnabled interface{}

    // This object specifies whether enabled for Frequency input using the 1.544
    // Mhz, 2.048 Mhz, or 10Mhz timing interface. The type is bool.
    CPtpClockInputFrequencyEnabled interface{}

    // This object specifies whether the node is enabled for TOD. The type is
    // bool.
    CPtpClockTODEnabled interface{}

    // This object specifies whether the node is enabled for PTP input clocking
    // using the 1pps interface. The type is bool.
    CPtpClockOutput1ppsEnabled interface{}

    // This object specifies whether an offset is configured in order to
    // compensate for a known phase error such as network asymmetry. The type is
    // bool.
    CPtpClockOutput1ppsOffsetEnabled interface{}

    // This object specifies the fixed offset value configured to be added for the
    // 1pps output. The type is interface{} with range: 0..4294967295.
    CPtpClockOutput1ppsOffsetValue interface{}

    // This object specifies whether the added (fixed) offset to the 1pps output
    // is negative or not.  When object returns TRUE the offset is negative and
    // when object returns FALSE the offset is positive. The type is bool.
    CPtpClockOutput1ppsOffsetNegative interface{}

    // This object specifies the 1pps interface used for PTP input clocking. The
    // type is string.
    CPtpClockInput1ppsInterface interface{}

    // This object specifies the 1pps interface used for PTP output clocking. The
    // type is string.
    CPtpClockOutput1ppsInterface interface{}

    // This object specifies the interface used for PTP TOD. The type is string.
    CPtpClockTODInterface interface{}
}

func (cPtpClockNodeEntry *CISCOPTPMIB_CPtpClockNodeTable_CPtpClockNodeEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockNodeEntry.EntityData.YFilter = cPtpClockNodeEntry.YFilter
    cPtpClockNodeEntry.EntityData.YangName = "cPtpClockNodeEntry"
    cPtpClockNodeEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockNodeEntry.EntityData.ParentYangName = "cPtpClockNodeTable"
    cPtpClockNodeEntry.EntityData.SegmentPath = "cPtpClockNodeEntry" + types.AddKeyToken(cPtpClockNodeEntry.CPtpClockDomainIndex, "cPtpClockDomainIndex") + types.AddKeyToken(cPtpClockNodeEntry.CPtpClockTypeIndex, "cPtpClockTypeIndex") + types.AddKeyToken(cPtpClockNodeEntry.CPtpClockInstanceIndex, "cPtpClockInstanceIndex")
    cPtpClockNodeEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockNodeTable/" + cPtpClockNodeEntry.EntityData.SegmentPath
    cPtpClockNodeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockNodeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockNodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockNodeEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockNodeEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockDomainIndex", types.YLeaf{"CPtpClockDomainIndex", cPtpClockNodeEntry.CPtpClockDomainIndex})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockTypeIndex", types.YLeaf{"CPtpClockTypeIndex", cPtpClockNodeEntry.CPtpClockTypeIndex})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockInstanceIndex", types.YLeaf{"CPtpClockInstanceIndex", cPtpClockNodeEntry.CPtpClockInstanceIndex})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockInput1ppsEnabled", types.YLeaf{"CPtpClockInput1ppsEnabled", cPtpClockNodeEntry.CPtpClockInput1ppsEnabled})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockInputFrequencyEnabled", types.YLeaf{"CPtpClockInputFrequencyEnabled", cPtpClockNodeEntry.CPtpClockInputFrequencyEnabled})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockTODEnabled", types.YLeaf{"CPtpClockTODEnabled", cPtpClockNodeEntry.CPtpClockTODEnabled})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockOutput1ppsEnabled", types.YLeaf{"CPtpClockOutput1ppsEnabled", cPtpClockNodeEntry.CPtpClockOutput1ppsEnabled})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockOutput1ppsOffsetEnabled", types.YLeaf{"CPtpClockOutput1ppsOffsetEnabled", cPtpClockNodeEntry.CPtpClockOutput1ppsOffsetEnabled})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockOutput1ppsOffsetValue", types.YLeaf{"CPtpClockOutput1ppsOffsetValue", cPtpClockNodeEntry.CPtpClockOutput1ppsOffsetValue})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockOutput1ppsOffsetNegative", types.YLeaf{"CPtpClockOutput1ppsOffsetNegative", cPtpClockNodeEntry.CPtpClockOutput1ppsOffsetNegative})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockInput1ppsInterface", types.YLeaf{"CPtpClockInput1ppsInterface", cPtpClockNodeEntry.CPtpClockInput1ppsInterface})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockOutput1ppsInterface", types.YLeaf{"CPtpClockOutput1ppsInterface", cPtpClockNodeEntry.CPtpClockOutput1ppsInterface})
    cPtpClockNodeEntry.EntityData.Leafs.Append("cPtpClockTODInterface", types.YLeaf{"CPtpClockTODInterface", cPtpClockNodeEntry.CPtpClockTODInterface})

    cPtpClockNodeEntry.EntityData.YListKeys = []string {"CPtpClockDomainIndex", "CPtpClockTypeIndex", "CPtpClockInstanceIndex"}

    return &(cPtpClockNodeEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockCurrentDSTable
// Table of information about the PTP clock Current Datasets for
// all domains.
type CISCOPTPMIB_CPtpClockCurrentDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Current Datasets for a domain. The type is slice of
    // CISCOPTPMIB_CPtpClockCurrentDSTable_CPtpClockCurrentDSEntry.
    CPtpClockCurrentDSEntry []*CISCOPTPMIB_CPtpClockCurrentDSTable_CPtpClockCurrentDSEntry
}

func (cPtpClockCurrentDSTable *CISCOPTPMIB_CPtpClockCurrentDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockCurrentDSTable.EntityData.YFilter = cPtpClockCurrentDSTable.YFilter
    cPtpClockCurrentDSTable.EntityData.YangName = "cPtpClockCurrentDSTable"
    cPtpClockCurrentDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockCurrentDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockCurrentDSTable.EntityData.SegmentPath = "cPtpClockCurrentDSTable"
    cPtpClockCurrentDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockCurrentDSTable.EntityData.SegmentPath
    cPtpClockCurrentDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockCurrentDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockCurrentDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockCurrentDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockCurrentDSTable.EntityData.Children.Append("cPtpClockCurrentDSEntry", types.YChild{"CPtpClockCurrentDSEntry", nil})
    for i := range cPtpClockCurrentDSTable.CPtpClockCurrentDSEntry {
        cPtpClockCurrentDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockCurrentDSTable.CPtpClockCurrentDSEntry[i]), types.YChild{"CPtpClockCurrentDSEntry", cPtpClockCurrentDSTable.CPtpClockCurrentDSEntry[i]})
    }
    cPtpClockCurrentDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockCurrentDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockCurrentDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockCurrentDSTable_CPtpClockCurrentDSEntry
// An entry in the table, containing information about a single
// PTP clock Current Datasets for a domain.
type CISCOPTPMIB_CPtpClockCurrentDSTable_CPtpClockCurrentDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockCurrentDSDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockCurrentDSClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockCurrentDSInstanceIndex interface{}

    // The current clock dataset StepsRemoved value.  This object specifies the
    // distance measured by the number of Boundary clocks between the local clock
    // and the Foreign master as indicated in the stepsRemoved field of Announce
    // messages. The type is interface{} with range: 0..4294967295. Units are
    // steps.
    CPtpClockCurrentDSStepsRemoved interface{}

    // This object specifies the current clock dataset ClockOffset value. The
    // value of the computation of the offset in time between a slave and a master
    // clock. The type is string with length: 1..255. Units are Time Interval.
    CPtpClockCurrentDSOffsetFromMaster interface{}

    // This object specifies the current clock dataset MeanPathDelay value.  The
    // mean path delay between a pair of ports as measure by the delay
    // request-response mechanism. The type is string with length: 1..255.
    CPtpClockCurrentDSMeanPathDelay interface{}
}

func (cPtpClockCurrentDSEntry *CISCOPTPMIB_CPtpClockCurrentDSTable_CPtpClockCurrentDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockCurrentDSEntry.EntityData.YFilter = cPtpClockCurrentDSEntry.YFilter
    cPtpClockCurrentDSEntry.EntityData.YangName = "cPtpClockCurrentDSEntry"
    cPtpClockCurrentDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockCurrentDSEntry.EntityData.ParentYangName = "cPtpClockCurrentDSTable"
    cPtpClockCurrentDSEntry.EntityData.SegmentPath = "cPtpClockCurrentDSEntry" + types.AddKeyToken(cPtpClockCurrentDSEntry.CPtpClockCurrentDSDomainIndex, "cPtpClockCurrentDSDomainIndex") + types.AddKeyToken(cPtpClockCurrentDSEntry.CPtpClockCurrentDSClockTypeIndex, "cPtpClockCurrentDSClockTypeIndex") + types.AddKeyToken(cPtpClockCurrentDSEntry.CPtpClockCurrentDSInstanceIndex, "cPtpClockCurrentDSInstanceIndex")
    cPtpClockCurrentDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockCurrentDSTable/" + cPtpClockCurrentDSEntry.EntityData.SegmentPath
    cPtpClockCurrentDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockCurrentDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockCurrentDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockCurrentDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockCurrentDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockCurrentDSEntry.EntityData.Leafs.Append("cPtpClockCurrentDSDomainIndex", types.YLeaf{"CPtpClockCurrentDSDomainIndex", cPtpClockCurrentDSEntry.CPtpClockCurrentDSDomainIndex})
    cPtpClockCurrentDSEntry.EntityData.Leafs.Append("cPtpClockCurrentDSClockTypeIndex", types.YLeaf{"CPtpClockCurrentDSClockTypeIndex", cPtpClockCurrentDSEntry.CPtpClockCurrentDSClockTypeIndex})
    cPtpClockCurrentDSEntry.EntityData.Leafs.Append("cPtpClockCurrentDSInstanceIndex", types.YLeaf{"CPtpClockCurrentDSInstanceIndex", cPtpClockCurrentDSEntry.CPtpClockCurrentDSInstanceIndex})
    cPtpClockCurrentDSEntry.EntityData.Leafs.Append("cPtpClockCurrentDSStepsRemoved", types.YLeaf{"CPtpClockCurrentDSStepsRemoved", cPtpClockCurrentDSEntry.CPtpClockCurrentDSStepsRemoved})
    cPtpClockCurrentDSEntry.EntityData.Leafs.Append("cPtpClockCurrentDSOffsetFromMaster", types.YLeaf{"CPtpClockCurrentDSOffsetFromMaster", cPtpClockCurrentDSEntry.CPtpClockCurrentDSOffsetFromMaster})
    cPtpClockCurrentDSEntry.EntityData.Leafs.Append("cPtpClockCurrentDSMeanPathDelay", types.YLeaf{"CPtpClockCurrentDSMeanPathDelay", cPtpClockCurrentDSEntry.CPtpClockCurrentDSMeanPathDelay})

    cPtpClockCurrentDSEntry.EntityData.YListKeys = []string {"CPtpClockCurrentDSDomainIndex", "CPtpClockCurrentDSClockTypeIndex", "CPtpClockCurrentDSInstanceIndex"}

    return &(cPtpClockCurrentDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockParentDSTable
// Table of information about the PTP clock Parent Datasets for
// all domains.
type CISCOPTPMIB_CPtpClockParentDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Parent Datasets for a domain. The type is slice of
    // CISCOPTPMIB_CPtpClockParentDSTable_CPtpClockParentDSEntry.
    CPtpClockParentDSEntry []*CISCOPTPMIB_CPtpClockParentDSTable_CPtpClockParentDSEntry
}

func (cPtpClockParentDSTable *CISCOPTPMIB_CPtpClockParentDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockParentDSTable.EntityData.YFilter = cPtpClockParentDSTable.YFilter
    cPtpClockParentDSTable.EntityData.YangName = "cPtpClockParentDSTable"
    cPtpClockParentDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockParentDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockParentDSTable.EntityData.SegmentPath = "cPtpClockParentDSTable"
    cPtpClockParentDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockParentDSTable.EntityData.SegmentPath
    cPtpClockParentDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockParentDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockParentDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockParentDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockParentDSTable.EntityData.Children.Append("cPtpClockParentDSEntry", types.YChild{"CPtpClockParentDSEntry", nil})
    for i := range cPtpClockParentDSTable.CPtpClockParentDSEntry {
        cPtpClockParentDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockParentDSTable.CPtpClockParentDSEntry[i]), types.YChild{"CPtpClockParentDSEntry", cPtpClockParentDSTable.CPtpClockParentDSEntry[i]})
    }
    cPtpClockParentDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockParentDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockParentDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockParentDSTable_CPtpClockParentDSEntry
// An entry in the table, containing information about a single
// PTP clock Parent Datasets for a domain.
type CISCOPTPMIB_CPtpClockParentDSTable_CPtpClockParentDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockParentDSDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockParentDSClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockParentDSInstanceIndex interface{}

    // This object specifies the value of portIdentity of the port on the master
    // that issues the Sync messages used in synchronizing this clock. The type is
    // string.
    CPtpClockParentDSParentPortIdentity interface{}

    // This object specifies the Parent Dataset ParentStats value.  This value
    // indicates whether the values of ParentDSOffset and ParentDSClockPhChRate
    // have been measured and are valid. A TRUE value shall indicate valid data.
    // The type is bool.
    CPtpClockParentDSParentStats interface{}

    // This object specifies the Parent Dataset ParentOffsetScaledLogVariance
    // value.  This value is the variance of the parent clocks phase as measured
    // by the local clock. The type is interface{} with range: -128..127.
    CPtpClockParentDSOffset interface{}

    // This object specifies the clock's parent dataset ParentClockPhaseChangeRate
    // value.  This value is an estimate of the parent clocks phase change rate as
    // measured by the slave clock. The type is interface{} with range:
    // -2147483648..2147483647.
    CPtpClockParentDSClockPhChRate interface{}

    // This object specifies the parent dataset Grandmaster clock identity. The
    // type is string with length: 1..255.
    CPtpClockParentDSGMClockIdentity interface{}

    // This object specifies the parent dataset Grandmaster clock priority1. The
    // type is interface{} with range: -2147483648..2147483647.
    CPtpClockParentDSGMClockPriority1 interface{}

    // This object specifies the parent dataset grandmaster clock priority2. The
    // type is interface{} with range: -2147483648..2147483647.
    CPtpClockParentDSGMClockPriority2 interface{}

    // This object specifies the parent dataset grandmaster clock quality class.
    // The type is interface{} with range: 0..255.
    CPtpClockParentDSGMClockQualityClass interface{}

    // This object specifies the parent dataset grandmaster clock quality
    // accuracy. The type is ClockQualityAccuracyType.
    CPtpClockParentDSGMClockQualityAccuracy interface{}

    // This object specifies the parent dataset grandmaster clock quality offset.
    // The type is interface{} with range: 0..4294967295.
    CPtpClockParentDSGMClockQualityOffset interface{}
}

func (cPtpClockParentDSEntry *CISCOPTPMIB_CPtpClockParentDSTable_CPtpClockParentDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockParentDSEntry.EntityData.YFilter = cPtpClockParentDSEntry.YFilter
    cPtpClockParentDSEntry.EntityData.YangName = "cPtpClockParentDSEntry"
    cPtpClockParentDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockParentDSEntry.EntityData.ParentYangName = "cPtpClockParentDSTable"
    cPtpClockParentDSEntry.EntityData.SegmentPath = "cPtpClockParentDSEntry" + types.AddKeyToken(cPtpClockParentDSEntry.CPtpClockParentDSDomainIndex, "cPtpClockParentDSDomainIndex") + types.AddKeyToken(cPtpClockParentDSEntry.CPtpClockParentDSClockTypeIndex, "cPtpClockParentDSClockTypeIndex") + types.AddKeyToken(cPtpClockParentDSEntry.CPtpClockParentDSInstanceIndex, "cPtpClockParentDSInstanceIndex")
    cPtpClockParentDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockParentDSTable/" + cPtpClockParentDSEntry.EntityData.SegmentPath
    cPtpClockParentDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockParentDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockParentDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockParentDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockParentDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSDomainIndex", types.YLeaf{"CPtpClockParentDSDomainIndex", cPtpClockParentDSEntry.CPtpClockParentDSDomainIndex})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSClockTypeIndex", types.YLeaf{"CPtpClockParentDSClockTypeIndex", cPtpClockParentDSEntry.CPtpClockParentDSClockTypeIndex})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSInstanceIndex", types.YLeaf{"CPtpClockParentDSInstanceIndex", cPtpClockParentDSEntry.CPtpClockParentDSInstanceIndex})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSParentPortIdentity", types.YLeaf{"CPtpClockParentDSParentPortIdentity", cPtpClockParentDSEntry.CPtpClockParentDSParentPortIdentity})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSParentStats", types.YLeaf{"CPtpClockParentDSParentStats", cPtpClockParentDSEntry.CPtpClockParentDSParentStats})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSOffset", types.YLeaf{"CPtpClockParentDSOffset", cPtpClockParentDSEntry.CPtpClockParentDSOffset})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSClockPhChRate", types.YLeaf{"CPtpClockParentDSClockPhChRate", cPtpClockParentDSEntry.CPtpClockParentDSClockPhChRate})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSGMClockIdentity", types.YLeaf{"CPtpClockParentDSGMClockIdentity", cPtpClockParentDSEntry.CPtpClockParentDSGMClockIdentity})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSGMClockPriority1", types.YLeaf{"CPtpClockParentDSGMClockPriority1", cPtpClockParentDSEntry.CPtpClockParentDSGMClockPriority1})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSGMClockPriority2", types.YLeaf{"CPtpClockParentDSGMClockPriority2", cPtpClockParentDSEntry.CPtpClockParentDSGMClockPriority2})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSGMClockQualityClass", types.YLeaf{"CPtpClockParentDSGMClockQualityClass", cPtpClockParentDSEntry.CPtpClockParentDSGMClockQualityClass})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSGMClockQualityAccuracy", types.YLeaf{"CPtpClockParentDSGMClockQualityAccuracy", cPtpClockParentDSEntry.CPtpClockParentDSGMClockQualityAccuracy})
    cPtpClockParentDSEntry.EntityData.Leafs.Append("cPtpClockParentDSGMClockQualityOffset", types.YLeaf{"CPtpClockParentDSGMClockQualityOffset", cPtpClockParentDSEntry.CPtpClockParentDSGMClockQualityOffset})

    cPtpClockParentDSEntry.EntityData.YListKeys = []string {"CPtpClockParentDSDomainIndex", "CPtpClockParentDSClockTypeIndex", "CPtpClockParentDSInstanceIndex"}

    return &(cPtpClockParentDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockDefaultDSTable
// Table of information about the PTP clock Default Datasets for
// all domains.
type CISCOPTPMIB_CPtpClockDefaultDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // Default Datasets for a domain. The type is slice of
    // CISCOPTPMIB_CPtpClockDefaultDSTable_CPtpClockDefaultDSEntry.
    CPtpClockDefaultDSEntry []*CISCOPTPMIB_CPtpClockDefaultDSTable_CPtpClockDefaultDSEntry
}

func (cPtpClockDefaultDSTable *CISCOPTPMIB_CPtpClockDefaultDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockDefaultDSTable.EntityData.YFilter = cPtpClockDefaultDSTable.YFilter
    cPtpClockDefaultDSTable.EntityData.YangName = "cPtpClockDefaultDSTable"
    cPtpClockDefaultDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockDefaultDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockDefaultDSTable.EntityData.SegmentPath = "cPtpClockDefaultDSTable"
    cPtpClockDefaultDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockDefaultDSTable.EntityData.SegmentPath
    cPtpClockDefaultDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockDefaultDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockDefaultDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockDefaultDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockDefaultDSTable.EntityData.Children.Append("cPtpClockDefaultDSEntry", types.YChild{"CPtpClockDefaultDSEntry", nil})
    for i := range cPtpClockDefaultDSTable.CPtpClockDefaultDSEntry {
        cPtpClockDefaultDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockDefaultDSTable.CPtpClockDefaultDSEntry[i]), types.YChild{"CPtpClockDefaultDSEntry", cPtpClockDefaultDSTable.CPtpClockDefaultDSEntry[i]})
    }
    cPtpClockDefaultDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockDefaultDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockDefaultDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockDefaultDSTable_CPtpClockDefaultDSEntry
// An entry in the table, containing information about a single
// PTP clock Default Datasets for a domain.
type CISCOPTPMIB_CPtpClockDefaultDSTable_CPtpClockDefaultDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockDefaultDSDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockDefaultDSClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockDefaultDSInstanceIndex interface{}

    // This object specifies whether the Two Step process is used. The type is
    // bool.
    CPtpClockDefaultDSTwoStepFlag interface{}

    // This object specifies the default Datasets clock identity. The type is
    // string with length: 1..255.
    CPtpClockDefaultDSClockIdentity interface{}

    // This object specifies the default Datasets clock Priority1. The type is
    // interface{} with range: -2147483648..2147483647.
    CPtpClockDefaultDSPriority1 interface{}

    // This object specifies the default Datasets clock Priority2. The type is
    // interface{} with range: -2147483648..2147483647.
    CPtpClockDefaultDSPriority2 interface{}

    // Whether the SlaveOnly flag is set. The type is bool.
    CPtpClockDefaultDSSlaveOnly interface{}

    // This object specifies the default dataset Quality Class. The type is
    // interface{} with range: 0..255.
    CPtpClockDefaultDSQualityClass interface{}

    // This object specifies the default dataset Quality Accurarcy. The type is
    // ClockQualityAccuracyType.
    CPtpClockDefaultDSQualityAccuracy interface{}

    // This object specifies the default dataset Quality offset. The type is
    // interface{} with range: -2147483648..2147483647.
    CPtpClockDefaultDSQualityOffset interface{}
}

func (cPtpClockDefaultDSEntry *CISCOPTPMIB_CPtpClockDefaultDSTable_CPtpClockDefaultDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockDefaultDSEntry.EntityData.YFilter = cPtpClockDefaultDSEntry.YFilter
    cPtpClockDefaultDSEntry.EntityData.YangName = "cPtpClockDefaultDSEntry"
    cPtpClockDefaultDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockDefaultDSEntry.EntityData.ParentYangName = "cPtpClockDefaultDSTable"
    cPtpClockDefaultDSEntry.EntityData.SegmentPath = "cPtpClockDefaultDSEntry" + types.AddKeyToken(cPtpClockDefaultDSEntry.CPtpClockDefaultDSDomainIndex, "cPtpClockDefaultDSDomainIndex") + types.AddKeyToken(cPtpClockDefaultDSEntry.CPtpClockDefaultDSClockTypeIndex, "cPtpClockDefaultDSClockTypeIndex") + types.AddKeyToken(cPtpClockDefaultDSEntry.CPtpClockDefaultDSInstanceIndex, "cPtpClockDefaultDSInstanceIndex")
    cPtpClockDefaultDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockDefaultDSTable/" + cPtpClockDefaultDSEntry.EntityData.SegmentPath
    cPtpClockDefaultDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockDefaultDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockDefaultDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockDefaultDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockDefaultDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSDomainIndex", types.YLeaf{"CPtpClockDefaultDSDomainIndex", cPtpClockDefaultDSEntry.CPtpClockDefaultDSDomainIndex})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSClockTypeIndex", types.YLeaf{"CPtpClockDefaultDSClockTypeIndex", cPtpClockDefaultDSEntry.CPtpClockDefaultDSClockTypeIndex})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSInstanceIndex", types.YLeaf{"CPtpClockDefaultDSInstanceIndex", cPtpClockDefaultDSEntry.CPtpClockDefaultDSInstanceIndex})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSTwoStepFlag", types.YLeaf{"CPtpClockDefaultDSTwoStepFlag", cPtpClockDefaultDSEntry.CPtpClockDefaultDSTwoStepFlag})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSClockIdentity", types.YLeaf{"CPtpClockDefaultDSClockIdentity", cPtpClockDefaultDSEntry.CPtpClockDefaultDSClockIdentity})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSPriority1", types.YLeaf{"CPtpClockDefaultDSPriority1", cPtpClockDefaultDSEntry.CPtpClockDefaultDSPriority1})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSPriority2", types.YLeaf{"CPtpClockDefaultDSPriority2", cPtpClockDefaultDSEntry.CPtpClockDefaultDSPriority2})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSSlaveOnly", types.YLeaf{"CPtpClockDefaultDSSlaveOnly", cPtpClockDefaultDSEntry.CPtpClockDefaultDSSlaveOnly})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSQualityClass", types.YLeaf{"CPtpClockDefaultDSQualityClass", cPtpClockDefaultDSEntry.CPtpClockDefaultDSQualityClass})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSQualityAccuracy", types.YLeaf{"CPtpClockDefaultDSQualityAccuracy", cPtpClockDefaultDSEntry.CPtpClockDefaultDSQualityAccuracy})
    cPtpClockDefaultDSEntry.EntityData.Leafs.Append("cPtpClockDefaultDSQualityOffset", types.YLeaf{"CPtpClockDefaultDSQualityOffset", cPtpClockDefaultDSEntry.CPtpClockDefaultDSQualityOffset})

    cPtpClockDefaultDSEntry.EntityData.YListKeys = []string {"CPtpClockDefaultDSDomainIndex", "CPtpClockDefaultDSClockTypeIndex", "CPtpClockDefaultDSInstanceIndex"}

    return &(cPtpClockDefaultDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockRunningTable
// Table of information about the PTP clock Running Datasets for
// all domains.
type CISCOPTPMIB_CPtpClockRunningTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // running Datasets for a domain. The type is slice of
    // CISCOPTPMIB_CPtpClockRunningTable_CPtpClockRunningEntry.
    CPtpClockRunningEntry []*CISCOPTPMIB_CPtpClockRunningTable_CPtpClockRunningEntry
}

func (cPtpClockRunningTable *CISCOPTPMIB_CPtpClockRunningTable) GetEntityData() *types.CommonEntityData {
    cPtpClockRunningTable.EntityData.YFilter = cPtpClockRunningTable.YFilter
    cPtpClockRunningTable.EntityData.YangName = "cPtpClockRunningTable"
    cPtpClockRunningTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockRunningTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockRunningTable.EntityData.SegmentPath = "cPtpClockRunningTable"
    cPtpClockRunningTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockRunningTable.EntityData.SegmentPath
    cPtpClockRunningTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockRunningTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockRunningTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockRunningTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockRunningTable.EntityData.Children.Append("cPtpClockRunningEntry", types.YChild{"CPtpClockRunningEntry", nil})
    for i := range cPtpClockRunningTable.CPtpClockRunningEntry {
        cPtpClockRunningTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockRunningTable.CPtpClockRunningEntry[i]), types.YChild{"CPtpClockRunningEntry", cPtpClockRunningTable.CPtpClockRunningEntry[i]})
    }
    cPtpClockRunningTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockRunningTable.EntityData.YListKeys = []string {}

    return &(cPtpClockRunningTable.EntityData)
}

// CISCOPTPMIB_CPtpClockRunningTable_CPtpClockRunningEntry
// An entry in the table, containing information about a single
// PTP clock running Datasets for a domain.
type CISCOPTPMIB_CPtpClockRunningTable_CPtpClockRunningEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockRunningDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockRunningClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockRunningInstanceIndex interface{}

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
    CPtpClockRunningState interface{}

    // This object specifies the total number of all packet Unicast and multicast
    // that have been sent out for this clock in this domain for this type. The
    // type is interface{} with range: 0..18446744073709551615.
    CPtpClockRunningPacketsSent interface{}

    // This object specifies the total number of all packet Unicast and multicast
    // that have been received for this clock in this domain for this type. The
    // type is interface{} with range: 0..18446744073709551615.
    CPtpClockRunningPacketsReceived interface{}
}

func (cPtpClockRunningEntry *CISCOPTPMIB_CPtpClockRunningTable_CPtpClockRunningEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockRunningEntry.EntityData.YFilter = cPtpClockRunningEntry.YFilter
    cPtpClockRunningEntry.EntityData.YangName = "cPtpClockRunningEntry"
    cPtpClockRunningEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockRunningEntry.EntityData.ParentYangName = "cPtpClockRunningTable"
    cPtpClockRunningEntry.EntityData.SegmentPath = "cPtpClockRunningEntry" + types.AddKeyToken(cPtpClockRunningEntry.CPtpClockRunningDomainIndex, "cPtpClockRunningDomainIndex") + types.AddKeyToken(cPtpClockRunningEntry.CPtpClockRunningClockTypeIndex, "cPtpClockRunningClockTypeIndex") + types.AddKeyToken(cPtpClockRunningEntry.CPtpClockRunningInstanceIndex, "cPtpClockRunningInstanceIndex")
    cPtpClockRunningEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockRunningTable/" + cPtpClockRunningEntry.EntityData.SegmentPath
    cPtpClockRunningEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockRunningEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockRunningEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockRunningEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockRunningEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockRunningEntry.EntityData.Leafs.Append("cPtpClockRunningDomainIndex", types.YLeaf{"CPtpClockRunningDomainIndex", cPtpClockRunningEntry.CPtpClockRunningDomainIndex})
    cPtpClockRunningEntry.EntityData.Leafs.Append("cPtpClockRunningClockTypeIndex", types.YLeaf{"CPtpClockRunningClockTypeIndex", cPtpClockRunningEntry.CPtpClockRunningClockTypeIndex})
    cPtpClockRunningEntry.EntityData.Leafs.Append("cPtpClockRunningInstanceIndex", types.YLeaf{"CPtpClockRunningInstanceIndex", cPtpClockRunningEntry.CPtpClockRunningInstanceIndex})
    cPtpClockRunningEntry.EntityData.Leafs.Append("cPtpClockRunningState", types.YLeaf{"CPtpClockRunningState", cPtpClockRunningEntry.CPtpClockRunningState})
    cPtpClockRunningEntry.EntityData.Leafs.Append("cPtpClockRunningPacketsSent", types.YLeaf{"CPtpClockRunningPacketsSent", cPtpClockRunningEntry.CPtpClockRunningPacketsSent})
    cPtpClockRunningEntry.EntityData.Leafs.Append("cPtpClockRunningPacketsReceived", types.YLeaf{"CPtpClockRunningPacketsReceived", cPtpClockRunningEntry.CPtpClockRunningPacketsReceived})

    cPtpClockRunningEntry.EntityData.YListKeys = []string {"CPtpClockRunningDomainIndex", "CPtpClockRunningClockTypeIndex", "CPtpClockRunningInstanceIndex"}

    return &(cPtpClockRunningEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockTimePropertiesDSTable
// Table of information about the PTP clock Timeproperties
// Datasets for all domains.
type CISCOPTPMIB_CPtpClockTimePropertiesDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP clock
    // timeproperties Datasets for a domain. The type is slice of
    // CISCOPTPMIB_CPtpClockTimePropertiesDSTable_CPtpClockTimePropertiesDSEntry.
    CPtpClockTimePropertiesDSEntry []*CISCOPTPMIB_CPtpClockTimePropertiesDSTable_CPtpClockTimePropertiesDSEntry
}

func (cPtpClockTimePropertiesDSTable *CISCOPTPMIB_CPtpClockTimePropertiesDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockTimePropertiesDSTable.EntityData.YFilter = cPtpClockTimePropertiesDSTable.YFilter
    cPtpClockTimePropertiesDSTable.EntityData.YangName = "cPtpClockTimePropertiesDSTable"
    cPtpClockTimePropertiesDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockTimePropertiesDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockTimePropertiesDSTable.EntityData.SegmentPath = "cPtpClockTimePropertiesDSTable"
    cPtpClockTimePropertiesDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockTimePropertiesDSTable.EntityData.SegmentPath
    cPtpClockTimePropertiesDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockTimePropertiesDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockTimePropertiesDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockTimePropertiesDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockTimePropertiesDSTable.EntityData.Children.Append("cPtpClockTimePropertiesDSEntry", types.YChild{"CPtpClockTimePropertiesDSEntry", nil})
    for i := range cPtpClockTimePropertiesDSTable.CPtpClockTimePropertiesDSEntry {
        cPtpClockTimePropertiesDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockTimePropertiesDSTable.CPtpClockTimePropertiesDSEntry[i]), types.YChild{"CPtpClockTimePropertiesDSEntry", cPtpClockTimePropertiesDSTable.CPtpClockTimePropertiesDSEntry[i]})
    }
    cPtpClockTimePropertiesDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockTimePropertiesDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockTimePropertiesDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockTimePropertiesDSTable_CPtpClockTimePropertiesDSEntry
// An entry in the table, containing information about a single
// PTP clock timeproperties Datasets for a domain.
type CISCOPTPMIB_CPtpClockTimePropertiesDSTable_CPtpClockTimePropertiesDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockTimePropertiesDSDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockTimePropertiesDSClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockTimePropertiesDSInstanceIndex interface{}

    // This object specifies the timeproperties dataset value of whether current
    // UTC offset is valid. The type is bool.
    CPtpClockTimePropertiesDSCurrentUTCOffsetValid interface{}

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
    CPtpClockTimePropertiesDSCurrentUTCOffset interface{}

    // This object specifies the Leap59 value in the clock Current Dataset. The
    // type is bool.
    CPtpClockTimePropertiesDSLeap59 interface{}

    // This object specifies the Leap61 value in the clock Current Dataset. The
    // type is bool.
    CPtpClockTimePropertiesDSLeap61 interface{}

    // This object specifies the Timetraceable value in the clock Current Dataset.
    // The type is bool.
    CPtpClockTimePropertiesDSTimeTraceable interface{}

    // This object specifies the Frequency Traceable value in the clock Current
    // Dataset. The type is bool.
    CPtpClockTimePropertiesDSFreqTraceable interface{}

    // This object specifies the PTP Timescale value in the clock Current Dataset.
    // The type is bool.
    CPtpClockTimePropertiesDSPTPTimescale interface{}

    // This object specifies the Timesource value in the clock Current Dataset.
    // The type is ClockTimeSourceType.
    CPtpClockTimePropertiesDSSource interface{}
}

func (cPtpClockTimePropertiesDSEntry *CISCOPTPMIB_CPtpClockTimePropertiesDSTable_CPtpClockTimePropertiesDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockTimePropertiesDSEntry.EntityData.YFilter = cPtpClockTimePropertiesDSEntry.YFilter
    cPtpClockTimePropertiesDSEntry.EntityData.YangName = "cPtpClockTimePropertiesDSEntry"
    cPtpClockTimePropertiesDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockTimePropertiesDSEntry.EntityData.ParentYangName = "cPtpClockTimePropertiesDSTable"
    cPtpClockTimePropertiesDSEntry.EntityData.SegmentPath = "cPtpClockTimePropertiesDSEntry" + types.AddKeyToken(cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSDomainIndex, "cPtpClockTimePropertiesDSDomainIndex") + types.AddKeyToken(cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSClockTypeIndex, "cPtpClockTimePropertiesDSClockTypeIndex") + types.AddKeyToken(cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSInstanceIndex, "cPtpClockTimePropertiesDSInstanceIndex")
    cPtpClockTimePropertiesDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockTimePropertiesDSTable/" + cPtpClockTimePropertiesDSEntry.EntityData.SegmentPath
    cPtpClockTimePropertiesDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockTimePropertiesDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockTimePropertiesDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockTimePropertiesDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSDomainIndex", types.YLeaf{"CPtpClockTimePropertiesDSDomainIndex", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSDomainIndex})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSClockTypeIndex", types.YLeaf{"CPtpClockTimePropertiesDSClockTypeIndex", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSClockTypeIndex})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSInstanceIndex", types.YLeaf{"CPtpClockTimePropertiesDSInstanceIndex", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSInstanceIndex})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSCurrentUTCOffsetValid", types.YLeaf{"CPtpClockTimePropertiesDSCurrentUTCOffsetValid", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSCurrentUTCOffsetValid})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSCurrentUTCOffset", types.YLeaf{"CPtpClockTimePropertiesDSCurrentUTCOffset", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSCurrentUTCOffset})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSLeap59", types.YLeaf{"CPtpClockTimePropertiesDSLeap59", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSLeap59})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSLeap61", types.YLeaf{"CPtpClockTimePropertiesDSLeap61", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSLeap61})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSTimeTraceable", types.YLeaf{"CPtpClockTimePropertiesDSTimeTraceable", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSTimeTraceable})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSFreqTraceable", types.YLeaf{"CPtpClockTimePropertiesDSFreqTraceable", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSFreqTraceable})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSPTPTimescale", types.YLeaf{"CPtpClockTimePropertiesDSPTPTimescale", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSPTPTimescale})
    cPtpClockTimePropertiesDSEntry.EntityData.Leafs.Append("cPtpClockTimePropertiesDSSource", types.YLeaf{"CPtpClockTimePropertiesDSSource", cPtpClockTimePropertiesDSEntry.CPtpClockTimePropertiesDSSource})

    cPtpClockTimePropertiesDSEntry.EntityData.YListKeys = []string {"CPtpClockTimePropertiesDSDomainIndex", "CPtpClockTimePropertiesDSClockTypeIndex", "CPtpClockTimePropertiesDSInstanceIndex"}

    return &(cPtpClockTimePropertiesDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockTransDefaultDSTable
// Table of information about the PTP Transparent clock Default
// Datasets for all domains.
type CISCOPTPMIB_CPtpClockTransDefaultDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single PTP
    // Transparent clock Default Datasets for a domain. The type is slice of
    // CISCOPTPMIB_CPtpClockTransDefaultDSTable_CPtpClockTransDefaultDSEntry.
    CPtpClockTransDefaultDSEntry []*CISCOPTPMIB_CPtpClockTransDefaultDSTable_CPtpClockTransDefaultDSEntry
}

func (cPtpClockTransDefaultDSTable *CISCOPTPMIB_CPtpClockTransDefaultDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockTransDefaultDSTable.EntityData.YFilter = cPtpClockTransDefaultDSTable.YFilter
    cPtpClockTransDefaultDSTable.EntityData.YangName = "cPtpClockTransDefaultDSTable"
    cPtpClockTransDefaultDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockTransDefaultDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockTransDefaultDSTable.EntityData.SegmentPath = "cPtpClockTransDefaultDSTable"
    cPtpClockTransDefaultDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockTransDefaultDSTable.EntityData.SegmentPath
    cPtpClockTransDefaultDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockTransDefaultDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockTransDefaultDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockTransDefaultDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockTransDefaultDSTable.EntityData.Children.Append("cPtpClockTransDefaultDSEntry", types.YChild{"CPtpClockTransDefaultDSEntry", nil})
    for i := range cPtpClockTransDefaultDSTable.CPtpClockTransDefaultDSEntry {
        cPtpClockTransDefaultDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockTransDefaultDSTable.CPtpClockTransDefaultDSEntry[i]), types.YChild{"CPtpClockTransDefaultDSEntry", cPtpClockTransDefaultDSTable.CPtpClockTransDefaultDSEntry[i]})
    }
    cPtpClockTransDefaultDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockTransDefaultDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockTransDefaultDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockTransDefaultDSTable_CPtpClockTransDefaultDSEntry
// An entry in the table, containing information about a single
// PTP Transparent clock Default Datasets for a domain.
type CISCOPTPMIB_CPtpClockTransDefaultDSTable_CPtpClockTransDefaultDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockTransDefaultDSDomainIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockTransDefaultDSInstanceIndex interface{}

    // This object specifies the value of the clockIdentity attribute of the local
    // clock. The type is string with length: 1..255.
    CPtpClockTransDefaultDSClockIdentity interface{}

    // This object specifies the number of PTP ports of the device. The type is
    // interface{} with range: 0..4294967295.
    CPtpClockTransDefaultDSNumOfPorts interface{}

    // This object, if the transparent clock is an end-to-end transparent clock,
    // has the value shall be E2E; If the transparent clock is a peer-to-peer
    // transparent clock, the value shall be P2P. The type is ClockMechanismType.
    CPtpClockTransDefaultDSDelay interface{}

    // This object specifies the value of the primary syntonization domain. The
    // initialization value shall be 0. The type is interface{} with range:
    // -2147483648..2147483647.
    CPtpClockTransDefaultDSPrimaryDomain interface{}
}

func (cPtpClockTransDefaultDSEntry *CISCOPTPMIB_CPtpClockTransDefaultDSTable_CPtpClockTransDefaultDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockTransDefaultDSEntry.EntityData.YFilter = cPtpClockTransDefaultDSEntry.YFilter
    cPtpClockTransDefaultDSEntry.EntityData.YangName = "cPtpClockTransDefaultDSEntry"
    cPtpClockTransDefaultDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockTransDefaultDSEntry.EntityData.ParentYangName = "cPtpClockTransDefaultDSTable"
    cPtpClockTransDefaultDSEntry.EntityData.SegmentPath = "cPtpClockTransDefaultDSEntry" + types.AddKeyToken(cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSDomainIndex, "cPtpClockTransDefaultDSDomainIndex") + types.AddKeyToken(cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSInstanceIndex, "cPtpClockTransDefaultDSInstanceIndex")
    cPtpClockTransDefaultDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockTransDefaultDSTable/" + cPtpClockTransDefaultDSEntry.EntityData.SegmentPath
    cPtpClockTransDefaultDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockTransDefaultDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockTransDefaultDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockTransDefaultDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockTransDefaultDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockTransDefaultDSEntry.EntityData.Leafs.Append("cPtpClockTransDefaultDSDomainIndex", types.YLeaf{"CPtpClockTransDefaultDSDomainIndex", cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSDomainIndex})
    cPtpClockTransDefaultDSEntry.EntityData.Leafs.Append("cPtpClockTransDefaultDSInstanceIndex", types.YLeaf{"CPtpClockTransDefaultDSInstanceIndex", cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSInstanceIndex})
    cPtpClockTransDefaultDSEntry.EntityData.Leafs.Append("cPtpClockTransDefaultDSClockIdentity", types.YLeaf{"CPtpClockTransDefaultDSClockIdentity", cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSClockIdentity})
    cPtpClockTransDefaultDSEntry.EntityData.Leafs.Append("cPtpClockTransDefaultDSNumOfPorts", types.YLeaf{"CPtpClockTransDefaultDSNumOfPorts", cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSNumOfPorts})
    cPtpClockTransDefaultDSEntry.EntityData.Leafs.Append("cPtpClockTransDefaultDSDelay", types.YLeaf{"CPtpClockTransDefaultDSDelay", cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSDelay})
    cPtpClockTransDefaultDSEntry.EntityData.Leafs.Append("cPtpClockTransDefaultDSPrimaryDomain", types.YLeaf{"CPtpClockTransDefaultDSPrimaryDomain", cPtpClockTransDefaultDSEntry.CPtpClockTransDefaultDSPrimaryDomain})

    cPtpClockTransDefaultDSEntry.EntityData.YListKeys = []string {"CPtpClockTransDefaultDSDomainIndex", "CPtpClockTransDefaultDSInstanceIndex"}

    return &(cPtpClockTransDefaultDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockPortTable
// Table of information about the clock ports for a particular
// domain.
type CISCOPTPMIB_CPtpClockPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single clock port.
    // The type is slice of CISCOPTPMIB_CPtpClockPortTable_CPtpClockPortEntry.
    CPtpClockPortEntry []*CISCOPTPMIB_CPtpClockPortTable_CPtpClockPortEntry
}

func (cPtpClockPortTable *CISCOPTPMIB_CPtpClockPortTable) GetEntityData() *types.CommonEntityData {
    cPtpClockPortTable.EntityData.YFilter = cPtpClockPortTable.YFilter
    cPtpClockPortTable.EntityData.YangName = "cPtpClockPortTable"
    cPtpClockPortTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockPortTable.EntityData.SegmentPath = "cPtpClockPortTable"
    cPtpClockPortTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockPortTable.EntityData.SegmentPath
    cPtpClockPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortTable.EntityData.Children.Append("cPtpClockPortEntry", types.YChild{"CPtpClockPortEntry", nil})
    for i := range cPtpClockPortTable.CPtpClockPortEntry {
        cPtpClockPortTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockPortTable.CPtpClockPortEntry[i]), types.YChild{"CPtpClockPortEntry", cPtpClockPortTable.CPtpClockPortEntry[i]})
    }
    cPtpClockPortTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockPortTable.EntityData.YListKeys = []string {}

    return &(cPtpClockPortTable.EntityData)
}

// CISCOPTPMIB_CPtpClockPortTable_CPtpClockPortEntry
// An entry in the table, containing information about a single
// clock port.
type CISCOPTPMIB_CPtpClockPortTable_CPtpClockPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockPortDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockPortClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockPortClockInstanceIndex interface{}

    // This attribute is a key. This object specifies the PTP Portnumber for this
    // port. The type is interface{} with range: 1..65535.
    CPtpClockPortTablePortNumberIndex interface{}

    // This object specifies the PTP clock port name configured on the router. The
    // type is string with length: 1..64.
    CPtpClockPortName interface{}

    // This object describes the current role (slave/master) of the port. The type
    // is ClockRoleType.
    CPtpClockPortRole interface{}

    // This object specifies that one-step clock operation between the PTP master
    // and slave device is enabled. The type is bool.
    CPtpClockPortSyncOneStep interface{}

    // This object specifies the current peer's network address used for PTP
    // communication. Based on the scenario and the setup involved, the values
    // might look like these - Scenario                   Value
    // -------------------   ---------------- Single Master          master port
    // Multiple Masters       selected master port Single Slave           slave
    // port Multiple Slaves        <empty>  (In relevant setups, information on
    // available slaves and available masters will be available through 
    // cPtpClockPortAssociateTable). The type is InetAddressType.
    CPtpClockPortCurrentPeerAddressType interface{}

    // This object specifies the current peer's network address used for PTP
    // communication. Based on the scenario and the setup involved, the values
    // might look like these - Scenario                   Value
    // -------------------   ---------------- Single Master          master port
    // Multiple Masters       selected master port Single Slave           slave
    // port Multiple Slaves        <empty>  (In relevant setups, information on
    // available slaves and available masters will be available through 
    // cPtpClockPortAssociateTable). The type is string with length: 0..255.
    CPtpClockPortCurrentPeerAddress interface{}

    // This object specifies - For a master port - the number of PTP slave
    // sessions (peers) associated with this PTP port. For a slave port - the
    // number of masters available to this slave port (might or might not be
    // peered). The type is interface{} with range: 0..4294967295.
    CPtpClockPortNumOfAssociatedPorts interface{}
}

func (cPtpClockPortEntry *CISCOPTPMIB_CPtpClockPortTable_CPtpClockPortEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockPortEntry.EntityData.YFilter = cPtpClockPortEntry.YFilter
    cPtpClockPortEntry.EntityData.YangName = "cPtpClockPortEntry"
    cPtpClockPortEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortEntry.EntityData.ParentYangName = "cPtpClockPortTable"
    cPtpClockPortEntry.EntityData.SegmentPath = "cPtpClockPortEntry" + types.AddKeyToken(cPtpClockPortEntry.CPtpClockPortDomainIndex, "cPtpClockPortDomainIndex") + types.AddKeyToken(cPtpClockPortEntry.CPtpClockPortClockTypeIndex, "cPtpClockPortClockTypeIndex") + types.AddKeyToken(cPtpClockPortEntry.CPtpClockPortClockInstanceIndex, "cPtpClockPortClockInstanceIndex") + types.AddKeyToken(cPtpClockPortEntry.CPtpClockPortTablePortNumberIndex, "cPtpClockPortTablePortNumberIndex")
    cPtpClockPortEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockPortTable/" + cPtpClockPortEntry.EntityData.SegmentPath
    cPtpClockPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortDomainIndex", types.YLeaf{"CPtpClockPortDomainIndex", cPtpClockPortEntry.CPtpClockPortDomainIndex})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortClockTypeIndex", types.YLeaf{"CPtpClockPortClockTypeIndex", cPtpClockPortEntry.CPtpClockPortClockTypeIndex})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortClockInstanceIndex", types.YLeaf{"CPtpClockPortClockInstanceIndex", cPtpClockPortEntry.CPtpClockPortClockInstanceIndex})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortTablePortNumberIndex", types.YLeaf{"CPtpClockPortTablePortNumberIndex", cPtpClockPortEntry.CPtpClockPortTablePortNumberIndex})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortName", types.YLeaf{"CPtpClockPortName", cPtpClockPortEntry.CPtpClockPortName})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortRole", types.YLeaf{"CPtpClockPortRole", cPtpClockPortEntry.CPtpClockPortRole})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortSyncOneStep", types.YLeaf{"CPtpClockPortSyncOneStep", cPtpClockPortEntry.CPtpClockPortSyncOneStep})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortCurrentPeerAddressType", types.YLeaf{"CPtpClockPortCurrentPeerAddressType", cPtpClockPortEntry.CPtpClockPortCurrentPeerAddressType})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortCurrentPeerAddress", types.YLeaf{"CPtpClockPortCurrentPeerAddress", cPtpClockPortEntry.CPtpClockPortCurrentPeerAddress})
    cPtpClockPortEntry.EntityData.Leafs.Append("cPtpClockPortNumOfAssociatedPorts", types.YLeaf{"CPtpClockPortNumOfAssociatedPorts", cPtpClockPortEntry.CPtpClockPortNumOfAssociatedPorts})

    cPtpClockPortEntry.EntityData.YListKeys = []string {"CPtpClockPortDomainIndex", "CPtpClockPortClockTypeIndex", "CPtpClockPortClockInstanceIndex", "CPtpClockPortTablePortNumberIndex"}

    return &(cPtpClockPortEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockPortDSTable
// Table of information about the clock ports dataset for a
// particular domain.
type CISCOPTPMIB_CPtpClockPortDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing port dataset information for a single
    // clock port. The type is slice of
    // CISCOPTPMIB_CPtpClockPortDSTable_CPtpClockPortDSEntry.
    CPtpClockPortDSEntry []*CISCOPTPMIB_CPtpClockPortDSTable_CPtpClockPortDSEntry
}

func (cPtpClockPortDSTable *CISCOPTPMIB_CPtpClockPortDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockPortDSTable.EntityData.YFilter = cPtpClockPortDSTable.YFilter
    cPtpClockPortDSTable.EntityData.YangName = "cPtpClockPortDSTable"
    cPtpClockPortDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockPortDSTable.EntityData.SegmentPath = "cPtpClockPortDSTable"
    cPtpClockPortDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockPortDSTable.EntityData.SegmentPath
    cPtpClockPortDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortDSTable.EntityData.Children.Append("cPtpClockPortDSEntry", types.YChild{"CPtpClockPortDSEntry", nil})
    for i := range cPtpClockPortDSTable.CPtpClockPortDSEntry {
        cPtpClockPortDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockPortDSTable.CPtpClockPortDSEntry[i]), types.YChild{"CPtpClockPortDSEntry", cPtpClockPortDSTable.CPtpClockPortDSEntry[i]})
    }
    cPtpClockPortDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockPortDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockPortDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockPortDSTable_CPtpClockPortDSEntry
// An entry in the table, containing port dataset information for
// a single clock port.
type CISCOPTPMIB_CPtpClockPortDSTable_CPtpClockPortDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockPortDSDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockPortDSClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockPortDSClockInstanceIndex interface{}

    // This attribute is a key. This object specifies the PTP portnumber
    // associated with this PTP port. The type is interface{} with range:
    // 1..65535.
    CPtpClockPortDSPortNumberIndex interface{}

    // This object specifies the PTP clock port name. The type is string with
    // length: 1..64.
    CPtpClockPortDSName interface{}

    // This object specifies the PTP clock port Identity. The type is string.
    CPtpClockPortDSPortIdentity interface{}

    // This object specifies the Announce message transmission interval associated
    // with this clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    CPtpClockPortDSAnnouncementInterval interface{}

    // This object specifies the Announce receipt timeout associated with this
    // clock port. The type is interface{} with range: -2147483648..2147483647.
    CPtpClockPortDSAnnounceRctTimeout interface{}

    // This object specifies the Sync message transmission interval. The type is
    // interface{} with range: -2147483648..2147483647.
    CPtpClockPortDSSyncInterval interface{}

    // This object specifies the Delay_Req message transmission interval. The type
    // is interface{} with range: -2147483648..2147483647.
    CPtpClockPortDSMinDelayReqInterval interface{}

    // This object specifies the Pdelay_Req message transmission interval. The
    // type is interface{} with range: -2147483648..2147483647.
    CPtpClockPortDSPeerDelayReqInterval interface{}

    // This object specifies the delay mechanism used. If the clock is an
    // end-to-end clock, the value of the is e2e, else if the clock is a peer
    // to-peer clock, the value shall be p2p. The type is ClockMechanismType.
    CPtpClockPortDSDelayMech interface{}

    // This object specifies the peer meanPathDelay. The type is string with
    // length: 1..255.
    CPtpClockPortDSPeerMeanPathDelay interface{}

    // This object specifies the grant duration allocated by the master. The type
    // is interface{} with range: 0..4294967295.
    CPtpClockPortDSGrantDuration interface{}

    // This object specifies the PTP version being used. The type is interface{}
    // with range: -2147483648..2147483647.
    CPtpClockPortDSPTPVersion interface{}
}

func (cPtpClockPortDSEntry *CISCOPTPMIB_CPtpClockPortDSTable_CPtpClockPortDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockPortDSEntry.EntityData.YFilter = cPtpClockPortDSEntry.YFilter
    cPtpClockPortDSEntry.EntityData.YangName = "cPtpClockPortDSEntry"
    cPtpClockPortDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortDSEntry.EntityData.ParentYangName = "cPtpClockPortDSTable"
    cPtpClockPortDSEntry.EntityData.SegmentPath = "cPtpClockPortDSEntry" + types.AddKeyToken(cPtpClockPortDSEntry.CPtpClockPortDSDomainIndex, "cPtpClockPortDSDomainIndex") + types.AddKeyToken(cPtpClockPortDSEntry.CPtpClockPortDSClockTypeIndex, "cPtpClockPortDSClockTypeIndex") + types.AddKeyToken(cPtpClockPortDSEntry.CPtpClockPortDSClockInstanceIndex, "cPtpClockPortDSClockInstanceIndex") + types.AddKeyToken(cPtpClockPortDSEntry.CPtpClockPortDSPortNumberIndex, "cPtpClockPortDSPortNumberIndex")
    cPtpClockPortDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockPortDSTable/" + cPtpClockPortDSEntry.EntityData.SegmentPath
    cPtpClockPortDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSDomainIndex", types.YLeaf{"CPtpClockPortDSDomainIndex", cPtpClockPortDSEntry.CPtpClockPortDSDomainIndex})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSClockTypeIndex", types.YLeaf{"CPtpClockPortDSClockTypeIndex", cPtpClockPortDSEntry.CPtpClockPortDSClockTypeIndex})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSClockInstanceIndex", types.YLeaf{"CPtpClockPortDSClockInstanceIndex", cPtpClockPortDSEntry.CPtpClockPortDSClockInstanceIndex})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSPortNumberIndex", types.YLeaf{"CPtpClockPortDSPortNumberIndex", cPtpClockPortDSEntry.CPtpClockPortDSPortNumberIndex})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSName", types.YLeaf{"CPtpClockPortDSName", cPtpClockPortDSEntry.CPtpClockPortDSName})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSPortIdentity", types.YLeaf{"CPtpClockPortDSPortIdentity", cPtpClockPortDSEntry.CPtpClockPortDSPortIdentity})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSAnnouncementInterval", types.YLeaf{"CPtpClockPortDSAnnouncementInterval", cPtpClockPortDSEntry.CPtpClockPortDSAnnouncementInterval})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSAnnounceRctTimeout", types.YLeaf{"CPtpClockPortDSAnnounceRctTimeout", cPtpClockPortDSEntry.CPtpClockPortDSAnnounceRctTimeout})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSSyncInterval", types.YLeaf{"CPtpClockPortDSSyncInterval", cPtpClockPortDSEntry.CPtpClockPortDSSyncInterval})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSMinDelayReqInterval", types.YLeaf{"CPtpClockPortDSMinDelayReqInterval", cPtpClockPortDSEntry.CPtpClockPortDSMinDelayReqInterval})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSPeerDelayReqInterval", types.YLeaf{"CPtpClockPortDSPeerDelayReqInterval", cPtpClockPortDSEntry.CPtpClockPortDSPeerDelayReqInterval})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSDelayMech", types.YLeaf{"CPtpClockPortDSDelayMech", cPtpClockPortDSEntry.CPtpClockPortDSDelayMech})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSPeerMeanPathDelay", types.YLeaf{"CPtpClockPortDSPeerMeanPathDelay", cPtpClockPortDSEntry.CPtpClockPortDSPeerMeanPathDelay})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSGrantDuration", types.YLeaf{"CPtpClockPortDSGrantDuration", cPtpClockPortDSEntry.CPtpClockPortDSGrantDuration})
    cPtpClockPortDSEntry.EntityData.Leafs.Append("cPtpClockPortDSPTPVersion", types.YLeaf{"CPtpClockPortDSPTPVersion", cPtpClockPortDSEntry.CPtpClockPortDSPTPVersion})

    cPtpClockPortDSEntry.EntityData.YListKeys = []string {"CPtpClockPortDSDomainIndex", "CPtpClockPortDSClockTypeIndex", "CPtpClockPortDSClockInstanceIndex", "CPtpClockPortDSPortNumberIndex"}

    return &(cPtpClockPortDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockPortRunningTable
// Table of information about the clock ports running dataset for
// a particular domain.
type CISCOPTPMIB_CPtpClockPortRunningTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing runing dataset information about a single
    // clock port. The type is slice of
    // CISCOPTPMIB_CPtpClockPortRunningTable_CPtpClockPortRunningEntry.
    CPtpClockPortRunningEntry []*CISCOPTPMIB_CPtpClockPortRunningTable_CPtpClockPortRunningEntry
}

func (cPtpClockPortRunningTable *CISCOPTPMIB_CPtpClockPortRunningTable) GetEntityData() *types.CommonEntityData {
    cPtpClockPortRunningTable.EntityData.YFilter = cPtpClockPortRunningTable.YFilter
    cPtpClockPortRunningTable.EntityData.YangName = "cPtpClockPortRunningTable"
    cPtpClockPortRunningTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortRunningTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockPortRunningTable.EntityData.SegmentPath = "cPtpClockPortRunningTable"
    cPtpClockPortRunningTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockPortRunningTable.EntityData.SegmentPath
    cPtpClockPortRunningTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortRunningTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortRunningTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortRunningTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortRunningTable.EntityData.Children.Append("cPtpClockPortRunningEntry", types.YChild{"CPtpClockPortRunningEntry", nil})
    for i := range cPtpClockPortRunningTable.CPtpClockPortRunningEntry {
        cPtpClockPortRunningTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockPortRunningTable.CPtpClockPortRunningEntry[i]), types.YChild{"CPtpClockPortRunningEntry", cPtpClockPortRunningTable.CPtpClockPortRunningEntry[i]})
    }
    cPtpClockPortRunningTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockPortRunningTable.EntityData.YListKeys = []string {}

    return &(cPtpClockPortRunningTable.EntityData)
}

// CISCOPTPMIB_CPtpClockPortRunningTable_CPtpClockPortRunningEntry
// An entry in the table, containing runing dataset information
// about a single clock port.
type CISCOPTPMIB_CPtpClockPortRunningTable_CPtpClockPortRunningEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockPortRunningDomainIndex interface{}

    // This attribute is a key. This object specifies the clock type as defined in
    // the Textual convention description. The type is ClockType.
    CPtpClockPortRunningClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockPortRunningClockInstanceIndex interface{}

    // This attribute is a key. This object specifies the PTP portnumber
    // associated with this clock port. The type is interface{} with range:
    // 1..65535.
    CPtpClockPortRunningPortNumberIndex interface{}

    // This object specifies the PTP clock port name. The type is string with
    // length: 1..64.
    CPtpClockPortRunningName interface{}

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
    CPtpClockPortRunningState interface{}

    // This object specifies the Clock Role. The type is ClockRoleType.
    CPtpClockPortRunningRole interface{}

    // This object specifies the interface on the router being used by the PTP
    // Clock for PTP communication. The type is interface{} with range:
    // 0..2147483647.
    CPtpClockPortRunningInterfaceIndex interface{}

    // This object specifirst the IP version being used for PTP communication (the
    // mapping used). The type is interface{} with range: -2147483648..2147483647.
    CPtpClockPortRunningIPversion interface{}

    // This object specifies the type of encapsulation if the interface is adding
    // extra  layers (eg. VLAN, Pseudowire encapsulation...) for the PTP messages.
    // The type is interface{} with range: -2147483648..2147483647.
    CPtpClockPortRunningEncapsulationType interface{}

    // This object specifies the clock transmission mode as  unicast:       Using
    // unicast commnuication channel. multicast:     Using Multicast communication
    // channel. multicast-mix: Using multicast-unicast communication channel. The
    // type is ClockTxModeType.
    CPtpClockPortRunningTxMode interface{}

    // This object specifie the clock receive mode as  unicast:       Using
    // unicast commnuication channel. multicast:     Using Multicast communication
    // channel. multicast-mix: Using multicast-unicast communication channel. The
    // type is ClockTxModeType.
    CPtpClockPortRunningRxMode interface{}

    // This object specifies the packets received on the clock port (cummulative).
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CPtpClockPortRunningPacketsReceived interface{}

    // This object specifies the packets sent on the clock port (cummulative). The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    CPtpClockPortRunningPacketsSent interface{}
}

func (cPtpClockPortRunningEntry *CISCOPTPMIB_CPtpClockPortRunningTable_CPtpClockPortRunningEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockPortRunningEntry.EntityData.YFilter = cPtpClockPortRunningEntry.YFilter
    cPtpClockPortRunningEntry.EntityData.YangName = "cPtpClockPortRunningEntry"
    cPtpClockPortRunningEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortRunningEntry.EntityData.ParentYangName = "cPtpClockPortRunningTable"
    cPtpClockPortRunningEntry.EntityData.SegmentPath = "cPtpClockPortRunningEntry" + types.AddKeyToken(cPtpClockPortRunningEntry.CPtpClockPortRunningDomainIndex, "cPtpClockPortRunningDomainIndex") + types.AddKeyToken(cPtpClockPortRunningEntry.CPtpClockPortRunningClockTypeIndex, "cPtpClockPortRunningClockTypeIndex") + types.AddKeyToken(cPtpClockPortRunningEntry.CPtpClockPortRunningClockInstanceIndex, "cPtpClockPortRunningClockInstanceIndex") + types.AddKeyToken(cPtpClockPortRunningEntry.CPtpClockPortRunningPortNumberIndex, "cPtpClockPortRunningPortNumberIndex")
    cPtpClockPortRunningEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockPortRunningTable/" + cPtpClockPortRunningEntry.EntityData.SegmentPath
    cPtpClockPortRunningEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortRunningEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortRunningEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortRunningEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortRunningEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningDomainIndex", types.YLeaf{"CPtpClockPortRunningDomainIndex", cPtpClockPortRunningEntry.CPtpClockPortRunningDomainIndex})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningClockTypeIndex", types.YLeaf{"CPtpClockPortRunningClockTypeIndex", cPtpClockPortRunningEntry.CPtpClockPortRunningClockTypeIndex})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningClockInstanceIndex", types.YLeaf{"CPtpClockPortRunningClockInstanceIndex", cPtpClockPortRunningEntry.CPtpClockPortRunningClockInstanceIndex})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningPortNumberIndex", types.YLeaf{"CPtpClockPortRunningPortNumberIndex", cPtpClockPortRunningEntry.CPtpClockPortRunningPortNumberIndex})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningName", types.YLeaf{"CPtpClockPortRunningName", cPtpClockPortRunningEntry.CPtpClockPortRunningName})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningState", types.YLeaf{"CPtpClockPortRunningState", cPtpClockPortRunningEntry.CPtpClockPortRunningState})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningRole", types.YLeaf{"CPtpClockPortRunningRole", cPtpClockPortRunningEntry.CPtpClockPortRunningRole})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningInterfaceIndex", types.YLeaf{"CPtpClockPortRunningInterfaceIndex", cPtpClockPortRunningEntry.CPtpClockPortRunningInterfaceIndex})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningIPversion", types.YLeaf{"CPtpClockPortRunningIPversion", cPtpClockPortRunningEntry.CPtpClockPortRunningIPversion})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningEncapsulationType", types.YLeaf{"CPtpClockPortRunningEncapsulationType", cPtpClockPortRunningEntry.CPtpClockPortRunningEncapsulationType})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningTxMode", types.YLeaf{"CPtpClockPortRunningTxMode", cPtpClockPortRunningEntry.CPtpClockPortRunningTxMode})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningRxMode", types.YLeaf{"CPtpClockPortRunningRxMode", cPtpClockPortRunningEntry.CPtpClockPortRunningRxMode})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningPacketsReceived", types.YLeaf{"CPtpClockPortRunningPacketsReceived", cPtpClockPortRunningEntry.CPtpClockPortRunningPacketsReceived})
    cPtpClockPortRunningEntry.EntityData.Leafs.Append("cPtpClockPortRunningPacketsSent", types.YLeaf{"CPtpClockPortRunningPacketsSent", cPtpClockPortRunningEntry.CPtpClockPortRunningPacketsSent})

    cPtpClockPortRunningEntry.EntityData.YListKeys = []string {"CPtpClockPortRunningDomainIndex", "CPtpClockPortRunningClockTypeIndex", "CPtpClockPortRunningClockInstanceIndex", "CPtpClockPortRunningPortNumberIndex"}

    return &(cPtpClockPortRunningEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockPortTransDSTable
// Table of information about the Transparent clock ports running
// dataset for a particular domain.
type CISCOPTPMIB_CPtpClockPortTransDSTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing clock port Transparent dataset
    // information about a single clock port. The type is slice of
    // CISCOPTPMIB_CPtpClockPortTransDSTable_CPtpClockPortTransDSEntry.
    CPtpClockPortTransDSEntry []*CISCOPTPMIB_CPtpClockPortTransDSTable_CPtpClockPortTransDSEntry
}

func (cPtpClockPortTransDSTable *CISCOPTPMIB_CPtpClockPortTransDSTable) GetEntityData() *types.CommonEntityData {
    cPtpClockPortTransDSTable.EntityData.YFilter = cPtpClockPortTransDSTable.YFilter
    cPtpClockPortTransDSTable.EntityData.YangName = "cPtpClockPortTransDSTable"
    cPtpClockPortTransDSTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortTransDSTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockPortTransDSTable.EntityData.SegmentPath = "cPtpClockPortTransDSTable"
    cPtpClockPortTransDSTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockPortTransDSTable.EntityData.SegmentPath
    cPtpClockPortTransDSTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortTransDSTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortTransDSTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortTransDSTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortTransDSTable.EntityData.Children.Append("cPtpClockPortTransDSEntry", types.YChild{"CPtpClockPortTransDSEntry", nil})
    for i := range cPtpClockPortTransDSTable.CPtpClockPortTransDSEntry {
        cPtpClockPortTransDSTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockPortTransDSTable.CPtpClockPortTransDSEntry[i]), types.YChild{"CPtpClockPortTransDSEntry", cPtpClockPortTransDSTable.CPtpClockPortTransDSEntry[i]})
    }
    cPtpClockPortTransDSTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockPortTransDSTable.EntityData.YListKeys = []string {}

    return &(cPtpClockPortTransDSTable.EntityData)
}

// CISCOPTPMIB_CPtpClockPortTransDSTable_CPtpClockPortTransDSEntry
// An entry in the table, containing clock port Transparent
// dataset information about a single clock port
type CISCOPTPMIB_CPtpClockPortTransDSTable_CPtpClockPortTransDSEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the domain number used to
    // create logical group of PTP devices. The type is interface{} with range:
    // 0..255.
    CPtpClockPortTransDSDomainIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockPortTransDSInstanceIndex interface{}

    // This attribute is a key. This object specifies the PTP port number
    // associated with this port. The type is interface{} with range: 1..65535.
    CPtpClockPortTransDSPortNumberIndex interface{}

    // This object specifies the value of the PortIdentity attribute of the local
    // port. The type is string with length: 1..255.
    CPtpClockPortTransDSPortIdentity interface{}

    // This object specifies the value of the logarithm to the base 2 of the
    // minPdelayReqInterval. The type is interface{} with range:
    // -2147483648..2147483647.
    CPtpClockPortTransDSlogMinPdelayReqInt interface{}

    // This object specifies the value TRUE if the port is faulty and FALSE if the
    // port is operating normally. The type is bool.
    CPtpClockPortTransDSFaultyFlag interface{}

    // This object specifies, (if the delayMechanism used is P2P) the value is the
    // estimate of the current one-way propagation delay, i.e., <meanPathDelay> on
    // the link attached to this port computed using the peer delay mechanism. If
    // the value of the delayMechanism used is E2E, then the value will be zero.
    // The type is string with length: 1..255.
    CPtpClockPortTransDSPeerMeanPathDelay interface{}
}

func (cPtpClockPortTransDSEntry *CISCOPTPMIB_CPtpClockPortTransDSTable_CPtpClockPortTransDSEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockPortTransDSEntry.EntityData.YFilter = cPtpClockPortTransDSEntry.YFilter
    cPtpClockPortTransDSEntry.EntityData.YangName = "cPtpClockPortTransDSEntry"
    cPtpClockPortTransDSEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortTransDSEntry.EntityData.ParentYangName = "cPtpClockPortTransDSTable"
    cPtpClockPortTransDSEntry.EntityData.SegmentPath = "cPtpClockPortTransDSEntry" + types.AddKeyToken(cPtpClockPortTransDSEntry.CPtpClockPortTransDSDomainIndex, "cPtpClockPortTransDSDomainIndex") + types.AddKeyToken(cPtpClockPortTransDSEntry.CPtpClockPortTransDSInstanceIndex, "cPtpClockPortTransDSInstanceIndex") + types.AddKeyToken(cPtpClockPortTransDSEntry.CPtpClockPortTransDSPortNumberIndex, "cPtpClockPortTransDSPortNumberIndex")
    cPtpClockPortTransDSEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockPortTransDSTable/" + cPtpClockPortTransDSEntry.EntityData.SegmentPath
    cPtpClockPortTransDSEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortTransDSEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortTransDSEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortTransDSEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortTransDSEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSDomainIndex", types.YLeaf{"CPtpClockPortTransDSDomainIndex", cPtpClockPortTransDSEntry.CPtpClockPortTransDSDomainIndex})
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSInstanceIndex", types.YLeaf{"CPtpClockPortTransDSInstanceIndex", cPtpClockPortTransDSEntry.CPtpClockPortTransDSInstanceIndex})
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSPortNumberIndex", types.YLeaf{"CPtpClockPortTransDSPortNumberIndex", cPtpClockPortTransDSEntry.CPtpClockPortTransDSPortNumberIndex})
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSPortIdentity", types.YLeaf{"CPtpClockPortTransDSPortIdentity", cPtpClockPortTransDSEntry.CPtpClockPortTransDSPortIdentity})
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSlogMinPdelayReqInt", types.YLeaf{"CPtpClockPortTransDSlogMinPdelayReqInt", cPtpClockPortTransDSEntry.CPtpClockPortTransDSlogMinPdelayReqInt})
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSFaultyFlag", types.YLeaf{"CPtpClockPortTransDSFaultyFlag", cPtpClockPortTransDSEntry.CPtpClockPortTransDSFaultyFlag})
    cPtpClockPortTransDSEntry.EntityData.Leafs.Append("cPtpClockPortTransDSPeerMeanPathDelay", types.YLeaf{"CPtpClockPortTransDSPeerMeanPathDelay", cPtpClockPortTransDSEntry.CPtpClockPortTransDSPeerMeanPathDelay})

    cPtpClockPortTransDSEntry.EntityData.YListKeys = []string {"CPtpClockPortTransDSDomainIndex", "CPtpClockPortTransDSInstanceIndex", "CPtpClockPortTransDSPortNumberIndex"}

    return &(cPtpClockPortTransDSEntry.EntityData)
}

// CISCOPTPMIB_CPtpClockPortAssociateTable
// Table of information about a given port's associated ports.
// 
// For a master port - multiple slave ports which have established
// sessions with the current master port.  
// For a slave port - the list of masters available for a given
// slave port. 
// 
// Session information (pkts, errors) to be displayed based on
// availability and scenario.
type CISCOPTPMIB_CPtpClockPortAssociateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single associated
    // port for the given clockport. The type is slice of
    // CISCOPTPMIB_CPtpClockPortAssociateTable_CPtpClockPortAssociateEntry.
    CPtpClockPortAssociateEntry []*CISCOPTPMIB_CPtpClockPortAssociateTable_CPtpClockPortAssociateEntry
}

func (cPtpClockPortAssociateTable *CISCOPTPMIB_CPtpClockPortAssociateTable) GetEntityData() *types.CommonEntityData {
    cPtpClockPortAssociateTable.EntityData.YFilter = cPtpClockPortAssociateTable.YFilter
    cPtpClockPortAssociateTable.EntityData.YangName = "cPtpClockPortAssociateTable"
    cPtpClockPortAssociateTable.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortAssociateTable.EntityData.ParentYangName = "CISCO-PTP-MIB"
    cPtpClockPortAssociateTable.EntityData.SegmentPath = "cPtpClockPortAssociateTable"
    cPtpClockPortAssociateTable.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/" + cPtpClockPortAssociateTable.EntityData.SegmentPath
    cPtpClockPortAssociateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortAssociateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortAssociateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortAssociateTable.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortAssociateTable.EntityData.Children.Append("cPtpClockPortAssociateEntry", types.YChild{"CPtpClockPortAssociateEntry", nil})
    for i := range cPtpClockPortAssociateTable.CPtpClockPortAssociateEntry {
        cPtpClockPortAssociateTable.EntityData.Children.Append(types.GetSegmentPath(cPtpClockPortAssociateTable.CPtpClockPortAssociateEntry[i]), types.YChild{"CPtpClockPortAssociateEntry", cPtpClockPortAssociateTable.CPtpClockPortAssociateEntry[i]})
    }
    cPtpClockPortAssociateTable.EntityData.Leafs = types.NewOrderedMap()

    cPtpClockPortAssociateTable.EntityData.YListKeys = []string {}

    return &(cPtpClockPortAssociateTable.EntityData)
}

// CISCOPTPMIB_CPtpClockPortAssociateTable_CPtpClockPortAssociateEntry
// An entry in the table, containing information about a single
// associated port for the given clockport.
type CISCOPTPMIB_CPtpClockPortAssociateTable_CPtpClockPortAssociateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object specifies the given port's domain
    // number. The type is interface{} with range: 0..255.
    CPtpClockPortCurrentDomainIndex interface{}

    // This attribute is a key. This object specifies the given port's clock type.
    // The type is ClockType.
    CPtpClockPortCurrentClockTypeIndex interface{}

    // This attribute is a key. This object specifies the instance of the clock
    // for this clock type in the given domain. The type is interface{} with
    // range: 0..255.
    CPtpClockPortCurrentClockInstanceIndex interface{}

    // This attribute is a key. This object specifies the PTP Port Number for the
    // given port. The type is interface{} with range: 0..65535.
    CPtpClockPortCurrentPortNumberIndex interface{}

    // This attribute is a key. This object specifies the associated port's serial
    // number in the current port's context. The type is interface{} with range:
    // 1..65535.
    CPtpClockPortAssociatePortIndex interface{}

    // This object specifies the peer port's network address type used for PTP
    // communication. The type is InetAddressType.
    CPtpClockPortAssociateAddressType interface{}

    // This object specifies the peer port's network address used for PTP
    // communication. The type is string with length: 0..255.
    CPtpClockPortAssociateAddress interface{}

    // The number of packets sent to this peer port from the current port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    CPtpClockPortAssociatePacketsSent interface{}

    // The number of packets received from this peer port by the current port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    CPtpClockPortAssociatePacketsReceived interface{}

    // This object specifies the input errors associated with the peer port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    CPtpClockPortAssociateInErrors interface{}

    // This object specifies the output errors associated with the peer port. The
    // type is interface{} with range: 0..18446744073709551615. Units are packets.
    CPtpClockPortAssociateOutErrors interface{}
}

func (cPtpClockPortAssociateEntry *CISCOPTPMIB_CPtpClockPortAssociateTable_CPtpClockPortAssociateEntry) GetEntityData() *types.CommonEntityData {
    cPtpClockPortAssociateEntry.EntityData.YFilter = cPtpClockPortAssociateEntry.YFilter
    cPtpClockPortAssociateEntry.EntityData.YangName = "cPtpClockPortAssociateEntry"
    cPtpClockPortAssociateEntry.EntityData.BundleName = "cisco_ios_xe"
    cPtpClockPortAssociateEntry.EntityData.ParentYangName = "cPtpClockPortAssociateTable"
    cPtpClockPortAssociateEntry.EntityData.SegmentPath = "cPtpClockPortAssociateEntry" + types.AddKeyToken(cPtpClockPortAssociateEntry.CPtpClockPortCurrentDomainIndex, "cPtpClockPortCurrentDomainIndex") + types.AddKeyToken(cPtpClockPortAssociateEntry.CPtpClockPortCurrentClockTypeIndex, "cPtpClockPortCurrentClockTypeIndex") + types.AddKeyToken(cPtpClockPortAssociateEntry.CPtpClockPortCurrentClockInstanceIndex, "cPtpClockPortCurrentClockInstanceIndex") + types.AddKeyToken(cPtpClockPortAssociateEntry.CPtpClockPortCurrentPortNumberIndex, "cPtpClockPortCurrentPortNumberIndex") + types.AddKeyToken(cPtpClockPortAssociateEntry.CPtpClockPortAssociatePortIndex, "cPtpClockPortAssociatePortIndex")
    cPtpClockPortAssociateEntry.EntityData.AbsolutePath = "CISCO-PTP-MIB:CISCO-PTP-MIB/cPtpClockPortAssociateTable/" + cPtpClockPortAssociateEntry.EntityData.SegmentPath
    cPtpClockPortAssociateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPtpClockPortAssociateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPtpClockPortAssociateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPtpClockPortAssociateEntry.EntityData.Children = types.NewOrderedMap()
    cPtpClockPortAssociateEntry.EntityData.Leafs = types.NewOrderedMap()
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortCurrentDomainIndex", types.YLeaf{"CPtpClockPortCurrentDomainIndex", cPtpClockPortAssociateEntry.CPtpClockPortCurrentDomainIndex})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortCurrentClockTypeIndex", types.YLeaf{"CPtpClockPortCurrentClockTypeIndex", cPtpClockPortAssociateEntry.CPtpClockPortCurrentClockTypeIndex})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortCurrentClockInstanceIndex", types.YLeaf{"CPtpClockPortCurrentClockInstanceIndex", cPtpClockPortAssociateEntry.CPtpClockPortCurrentClockInstanceIndex})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortCurrentPortNumberIndex", types.YLeaf{"CPtpClockPortCurrentPortNumberIndex", cPtpClockPortAssociateEntry.CPtpClockPortCurrentPortNumberIndex})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociatePortIndex", types.YLeaf{"CPtpClockPortAssociatePortIndex", cPtpClockPortAssociateEntry.CPtpClockPortAssociatePortIndex})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociateAddressType", types.YLeaf{"CPtpClockPortAssociateAddressType", cPtpClockPortAssociateEntry.CPtpClockPortAssociateAddressType})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociateAddress", types.YLeaf{"CPtpClockPortAssociateAddress", cPtpClockPortAssociateEntry.CPtpClockPortAssociateAddress})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociatePacketsSent", types.YLeaf{"CPtpClockPortAssociatePacketsSent", cPtpClockPortAssociateEntry.CPtpClockPortAssociatePacketsSent})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociatePacketsReceived", types.YLeaf{"CPtpClockPortAssociatePacketsReceived", cPtpClockPortAssociateEntry.CPtpClockPortAssociatePacketsReceived})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociateInErrors", types.YLeaf{"CPtpClockPortAssociateInErrors", cPtpClockPortAssociateEntry.CPtpClockPortAssociateInErrors})
    cPtpClockPortAssociateEntry.EntityData.Leafs.Append("cPtpClockPortAssociateOutErrors", types.YLeaf{"CPtpClockPortAssociateOutErrors", cPtpClockPortAssociateEntry.CPtpClockPortAssociateOutErrors})

    cPtpClockPortAssociateEntry.EntityData.YListKeys = []string {"CPtpClockPortCurrentDomainIndex", "CPtpClockPortCurrentClockTypeIndex", "CPtpClockPortCurrentClockInstanceIndex", "CPtpClockPortCurrentPortNumberIndex", "CPtpClockPortAssociatePortIndex"}

    return &(cPtpClockPortAssociateEntry.EntityData)
}

