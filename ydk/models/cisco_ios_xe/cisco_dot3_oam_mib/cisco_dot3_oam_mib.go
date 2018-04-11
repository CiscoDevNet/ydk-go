// The MIB module for managing the new Ethernet OAM features
// introduced by the Ethernet in the First Mile task force (IEEE
// 802.3ah).  The functionality presented here is based on IEEE
// 802.3ah [802.3ah], released in October, 2004.  [802.3ah] was
// prepared as an addendum to the standing version of IEEE 802.3
// [802.3-2002] at the time.  Since then, [802.3ah] has been
// merged into the base IEEE 802.3 specification in [802.3-2005].
// 
// In particular, this MIB focuses on the new OAM functions
// introduced in Clause 57 of [802.3ah].  The OAM functionality
// of Clause 57 is controlled by new management attributes
// introduced in Clause 30 of [802.3ah].  The OAM functions are
// not specific to any particular Ethernet physical layer, and
// can be generically applied to any Ethernet interface of
// [802.3-2002].  
// 
// An Ethernet OAM protocol data unit is a valid Ethernet frame
// with a destination MAC address equal to the reserved MAC
// address for Slow Protocols (See 43B of [802.3ah]), a
// lengthOrType field equal to the reserved type for Slow
// Protocols, and a Slow Protocols subtype equal to that of the
// subtype reserved for Ethernet OAM.  OAMPDU is used throughout
// this document as an abbreviation for Ethernet OAM protocol
// data unit.  
// 
// The following reference is used throughout this MIB module:  
// 
//   [802.3ah] refers to:
//     IEEE Std 802.3ah-2004: 'Draft amendment to -
//     Information technology - Telecommunications and
//     information exchange between systems - Local and
//     metropolitan are networks - Specific requirements - Part
//     3: Carrier sense multiple access with collision detection
//     (CSMA/CD) access method and physical layer specifications
//     - Media Access Control Parameters, Physical Layers and
//     Management Parameters for subscriber access networks',
//     October 2004.
// 
//   [802.3-2002] refers to:
//     IEEE Std 802.3-2002: 
//     'Information technology - Telecommunications and
//     information exchange between systems - Local and
//     metropolitan are networks - Specific requirements - Part
//     3: Carrier sense multiple access with collision detection
//     (CSMA/CD) access method and physical layer specifications
//     - Media Access Control Parameters, Physical Layers and
//     Management Parameters for subscriber access networks',
//     March 2002.
// 
//   [802.3-2005] refers to:
//     IEEE Std 802.3-2002: 
//     'Information technology - Telecommunications and
//     information exchange between systems - Local and
//     metropolitan are networks - Specific requirements - Part
//     3: Carrier sense multiple access with collision detection
//     (CSMA/CD) access method and physical layer specifications
//     - Media Access Control Parameters, Physical Layers and
//     Management Parameters for subscriber access networks',
//     December 2005.
// 
//   [802-2001] refers to:
//     'IEEE Standard for LAN/MAN (Local Area
//     Network/Metropolitan Area Network): Overview and
//     Architecture', IEEE 802, June 2001.
package cisco_dot3_oam_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_dot3_oam_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-DOT3-OAM-MIB CISCO-DOT3-OAM-MIB}", reflect.TypeOf(CISCODOT3OAMMIB{}))
    ydk.RegisterEntity("CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB", reflect.TypeOf(CISCODOT3OAMMIB{}))
}

// CISCODOT3OAMMIB
type CISCODOT3OAMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains the primary controls and status for the OAM
    // capabilities of an Ethernet like interface.  There will be one row in this
    // table for each Ethernet like interface in the system that supports the OAM
    // functions defined in [802.3ah].
    Cdot3Oamtable CISCODOT3OAMMIB_Cdot3Oamtable

    // This table contains information about the OAM peer for a particular
    // Ethernet like interface. OAM entities communicate with a single OAM peer
    // entity on Ethernet links on which OAM is enabled and operating properly. 
    // There is one entry in this table for each entry in the cdot3OamTable for
    // which information on the peer OAM entity is available.  .
    Cdot3Oampeertable CISCODOT3OAMMIB_Cdot3Oampeertable

    // This table contains controls for the loopback state of the local link as
    // well as indicating the status of the loopback function.  There is one entry
    // in this table for each entry in cdot3OamTable that supports loopback
    // functionality (where cdot3OamFunctionsSupported includes the
    // loopbackSupport bit set).  Loopback can be used to place the remote OAM
    // entity in a state where every received frame (except OAMPDUs) is echoed
    // back over the same interface on which they were received.   In this state,
    // at the remote entity, 'normal' traffic is disabled as only the looped back
    // frames are transmitted on the interface. Loopback is thus an intrusive
    // operation that prohibits normal data flow and should be used accordingly. 
    // .
    Cdot3Oamloopbacktable CISCODOT3OAMMIB_Cdot3Oamloopbacktable

    // This table contains statistics for the OAM function on a particular
    // Ethernet like interface. There is an entry in the table for every entry in
    // the cdot3OamTable.   The counters in this table are defined as 32-bit
    // entries to match the counter size as defined in [802.3ah].  Given the OAM
    // protocol is a slow protocol, the counters increment at a slow rate. .
    Cdot3Oamstatstable CISCODOT3OAMMIB_Cdot3Oamstatstable

    // Ethernet OAM includes the ability to generate and receive Event
    // Notification OAMPDUs to indicate various link problems. This table contains
    // the mechanisms to enable Event Notifications and configure the thresholds
    // to generate the standard Ethernet OAM events.  There is one entry in the
    // table for every entry in cdot3OamTable that supports OAM events (where
    // cdot3OamFunctionsSupported includes the eventSupport bit set). The values
    // in the table are maintained across changes to cdot3OamOperStatus.    The
    // standard threshold crossing events are:   - Errored Symbol Period Event. 
    // Generated when the number of     symbol errors exceeds a threshold within a
    // given window      defined by a number of symbols (for example, 1,000
    // symbols     out of 1,000,000 had errors).     - Errored Frame Period Event.
    // Generated when the number of      frame errors exceeds a threshold within a
    // given window      defined by a number of frames (for example, 10 frames out
    // of 1000 had errors).     - Errored Frame Event.  Generated when the number
    // of frame      errors exceeds a threshold within a given window defined     
    // by a period of time (for example, 10 frames in 1 second      had errors).  
    // - Errored Frame Seconds Summary Event.  Generated when the      number of
    // errored frame seconds exceeds a threshold within     a given time period
    // (for example, 10 errored frame seconds     within the last 100 seconds). 
    // An errored frame second is      defined as a 1 second interval which had >0
    // frame errors.   There are other events (dying gasp, critical events) that
    // are not threshold crossing events but which can be enabled/disabled via
    // this table.  .
    Cdot3Oameventconfigtable CISCODOT3OAMMIB_Cdot3Oameventconfigtable

    // This table records a history of the events that have occurred at the
    // Ethernet OAM level.  These events can include locally detected events,
    // which may result in locally generated OAMPDUs, and remotely detected
    // events, which are detected by the OAM peer entity and signaled to the local
    // entity via Ethernet OAM.  Ethernet OAM events can be signaled by Event
    // Notification OAMPDUs or by the flags field in any OAMPDU.    This table
    // contains both threshold crossing events and non-threshold crossing events. 
    // The parameters for the threshold window, threshold value, and actual value
    // (cdot3OamEventLogWindowXX, cdot3OamEventLogThresholdXX,
    // cdot3OamEventLogValue) are only applicable to threshold crossing events,
    // and are returned as all F's (2^32 - 1) for non-threshold crossing events.  
    // Entries in the table are automatically created when such events are
    // detected.  The size of the table is implementation dependent.  When the
    // table reaches its maximum size, older entries are automatically deleted to
    // make room for newer entries. .
    Cdot3Oameventlogtable CISCODOT3OAMMIB_Cdot3Oameventlogtable
}

func (cISCODOT3OAMMIB *CISCODOT3OAMMIB) GetEntityData() *types.CommonEntityData {
    cISCODOT3OAMMIB.EntityData.YFilter = cISCODOT3OAMMIB.YFilter
    cISCODOT3OAMMIB.EntityData.YangName = "CISCO-DOT3-OAM-MIB"
    cISCODOT3OAMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODOT3OAMMIB.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cISCODOT3OAMMIB.EntityData.SegmentPath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB"
    cISCODOT3OAMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODOT3OAMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODOT3OAMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODOT3OAMMIB.EntityData.Children = make(map[string]types.YChild)
    cISCODOT3OAMMIB.EntityData.Children["cdot3OamTable"] = types.YChild{"Cdot3Oamtable", &cISCODOT3OAMMIB.Cdot3Oamtable}
    cISCODOT3OAMMIB.EntityData.Children["cdot3OamPeerTable"] = types.YChild{"Cdot3Oampeertable", &cISCODOT3OAMMIB.Cdot3Oampeertable}
    cISCODOT3OAMMIB.EntityData.Children["cdot3OamLoopbackTable"] = types.YChild{"Cdot3Oamloopbacktable", &cISCODOT3OAMMIB.Cdot3Oamloopbacktable}
    cISCODOT3OAMMIB.EntityData.Children["cdot3OamStatsTable"] = types.YChild{"Cdot3Oamstatstable", &cISCODOT3OAMMIB.Cdot3Oamstatstable}
    cISCODOT3OAMMIB.EntityData.Children["cdot3OamEventConfigTable"] = types.YChild{"Cdot3Oameventconfigtable", &cISCODOT3OAMMIB.Cdot3Oameventconfigtable}
    cISCODOT3OAMMIB.EntityData.Children["cdot3OamEventLogTable"] = types.YChild{"Cdot3Oameventlogtable", &cISCODOT3OAMMIB.Cdot3Oameventlogtable}
    cISCODOT3OAMMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCODOT3OAMMIB.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oamtable
// This table contains the primary controls and status for the
// OAM capabilities of an Ethernet like interface.  There will be
// one row in this table for each Ethernet like interface in the
// system that supports the OAM functions defined in [802.3ah].
type CISCODOT3OAMMIB_Cdot3Oamtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information on the Ethernet OAM function
    // for a single Ethernet like interface. Entries in the table are created
    // automatically for each interface supporting Ethernet OAM. The status of the
    // row entry can be determined from cdot3OamOperStatus.    A cdot3OamEntry is
    // indexed in the cdot3OamTable by the ifIndex object of the Interfaces MIB. 
    // . The type is slice of CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry.
    Cdot3Oamentry []CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry
}

func (cdot3Oamtable *CISCODOT3OAMMIB_Cdot3Oamtable) GetEntityData() *types.CommonEntityData {
    cdot3Oamtable.EntityData.YFilter = cdot3Oamtable.YFilter
    cdot3Oamtable.EntityData.YangName = "cdot3OamTable"
    cdot3Oamtable.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oamtable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3Oamtable.EntityData.SegmentPath = "cdot3OamTable"
    cdot3Oamtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oamtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oamtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oamtable.EntityData.Children = make(map[string]types.YChild)
    cdot3Oamtable.EntityData.Children["cdot3OamEntry"] = types.YChild{"Cdot3Oamentry", nil}
    for i := range cdot3Oamtable.Cdot3Oamentry {
        cdot3Oamtable.EntityData.Children[types.GetSegmentPath(&cdot3Oamtable.Cdot3Oamentry[i])] = types.YChild{"Cdot3Oamentry", &cdot3Oamtable.Cdot3Oamentry[i]}
    }
    cdot3Oamtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdot3Oamtable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry
// An entry in the table, containing information on the Ethernet
// OAM function for a single Ethernet like interface. Entries in
// the table are created automatically for each interface
// supporting Ethernet OAM. The status of the row entry can be
// determined from cdot3OamOperStatus.  
// 
// A cdot3OamEntry is indexed in the cdot3OamTable by the ifIndex
// object of the Interfaces MIB.  
type CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object is used to provision the default administrative OAM mode for
    // this interface.  This object represents the desired state of OAM for this
    // interface.    The cdot3OamAdminState always starts in the disabled(1) state
    // until an explicit management action or configuration information retained
    // by the system causes a transition to the enabled(2) state.   When
    // enabled(2), Ethernet OAM will attempt to operate over this interface.  .
    // The type is Cdot3Oamadminstate.
    Cdot3Oamadminstate interface{}

    // At initialization and failure conditions, two OAM entities on the same
    // full-duplex Ethernet link begin a discovery phase to determine what OAM
    // capabilities may be used on that link.  The progress of this initialization
    // is controlled by the OAM sublayer.               This value is always
    // disabled(1) if OAM is disabled on this interface via the
    // cdot3OamAdminState.    If the link has detected a fault and is transmitting
    // OAMPDUs with a link fault indication, the value is linkFault(2). Also, if
    // the interface is not operational (ifOperStatus is not  up(1)), linkFault(2)
    // is returned.  Note that the object  ifOperStatus may not be up(1) as a
    // result of link failure or administrative action (ifAdminState being down(2)
    // or testing(3)).                     The passiveWait(3) state is returned
    // only by OAM entities in passive mode (cdot3OamMode) and reflects the state
    // in which the OAM entity is waiting to see if the peer device is OAM
    // capable.  The activeSendLocal(4) value is used by active mode devices
    // (cdot3OamMode) and reflects the OAM entity actively trying to discover
    // whether the peer has OAM capability but has not yet made that
    // determination.                     The state sendLocalAndRemote(5) reflects
    // that the local OAM entity has discovered the peer but has not yet accepted
    // or rejected the configuration of the peer.  The local device can, for
    // whatever reason, decide that the peer device is unacceptable and decline
    // OAM peering.  If the local OAM entity rejects the peer OAM entity, the
    // state becomes oamPeeringLocallyRejected(7).  If the OAM peering is allowed
    // by the local device, the state moves to sendLocalAndRemoteOk(6).  Note that
    // both the sendLocalAndRemote(5) and oamPeeringLocallyRejected(7) states fall
    // within the state SEND_LOCAL_REMOTE of the Discovery state diagram [802.3ah,
    // Figure 57-5], with the difference being whether the local OAM client has
    // actively rejected the peering or has just not indicated any decision yet. 
    // Whether a peering decision has been made is indicated via the local flags
    // field in the OAMPDU (reflected in the aOAMLocalFlagsField of 30.3.6.1.10). 
    // If the remote OAM entity rejects the peering, the state becomes
    // oamPeeringRemotelyRejected(8).  Note that both the sendLocalAndRemoteOk(6)
    // and oamPeeringRemotelyRejected(8) states fall within the state
    // SEND_LOCAL_REMOTE_OK of the Discovery state diagram [802.3ah, Figure 57-5],
    // with the difference being whether the remote OAM client has rejected the
    // peering or has just not yet decided.  This is indicated via the remote
    // flags field in the OAM PDU (reflected in the aOAMRemoteFlagsField of
    // 30.3.6.1.11).                     When the local OAM entity learns that
    // both it and the remote OAM entity have accepted the peering, the state
    // moves to operational(9) corresponding to the SEND_ANY state of the
    // Discovery state diagram [802.3ah, Figure 57-5].    Since Ethernet OAM
    // functions are not designed to work completely over half-duplex interfaces,
    // the value nonOperHalfDuplex(10) is returned whenever Ethernet OAM is
    // enabled (cdot3OamAdminState is enabled(1)) but the interface is in
    // half-duplex operation.  . The type is Cdot3Oamoperstatus.
    Cdot3Oamoperstatus interface{}

    // This object configures the mode of OAM operation for this Ethernet like
    // interface.  OAM on Ethernet interfaces may be in 'active' mode or 'passive'
    // mode.  These two modes differ in that active mode provides additional
    // capabilities to initiate monitoring activities with the remote OAM peer
    // entity, while passive mode generally waits for the peer to initiate OAM
    // actions with it.  As an example, an active OAM entity can put the remote
    // OAM entity in a loopback state, where a passive OAM entity cannot.    The
    // default value of cdot3OamMode is dependent on the type of system on which
    // this Ethernet like interface resides.  The default value should be
    // 'active(1)' unless it is known that this system should take on a
    // subservient role to the other device connected over this interface.   
    // Changing this value results in incrementing the configuration revision
    // field of locally generated OAMPDUs (30.3.6.1.12) and potentially re-doing
    // the OAM discovery process if the cdot3OamOperStatus was already
    // operational(9).  . The type is Cdot3Oammode.
    Cdot3Oammode interface{}

    // The largest OAMPDU that the OAM entity supports.  OAM entities exchange
    // maximum OAMPDU sizes and negotiate to use the smaller of the two maximum
    // OAMPDU sizes between the peers. This value is determined by the local
    // implementation.  . The type is interface{} with range: 64..1518. Units are
    // octets.
    Cdot3Oammaxoampdusize interface{}

    // The configuration revision of the OAM entity as reflected in the latest
    // OAMPDU sent by the OAM entity.  The config revision is used by OAM entities
    // to indicate configuration changes have occurred which might require the
    // peer OAM entity to re-evaluate whether OAM peering is allowed. . The type
    // is interface{} with range: 0..65535.
    Cdot3Oamconfigrevision interface{}

    // The OAM functions supported on this Ethernet like interface. OAM consists
    // of separate functional sets beyond the basic discovery process which is
    // always required.  These functional groups can be supported independently by
    // any implementation. These values are communicated to the peer via the local
    // configuration field of Information OAMPDUs.    Setting
    // 'unidirectionalSupport(0)' indicates that the OAM entity supports the
    // transmission of OAMPDUs on links that are operating in unidirectional mode
    // (traffic flowing in one direction only).  Setting 'loopbackSupport(1)'
    // indicates the OAM entity can initiate and respond to loopback commands.
    // Setting 'eventSupport(2)' indicates the OAM entity can send and receive
    // Event Notification OAMPDUs. Setting 'variableSupport(3)' indicates the OAM
    // entity can send and receive Variable Request and Response OAMPDUs.  . The
    // type is map[string]bool.
    Cdot3Oamfunctionssupported interface{}
}

func (cdot3Oamentry *CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry) GetEntityData() *types.CommonEntityData {
    cdot3Oamentry.EntityData.YFilter = cdot3Oamentry.YFilter
    cdot3Oamentry.EntityData.YangName = "cdot3OamEntry"
    cdot3Oamentry.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oamentry.EntityData.ParentYangName = "cdot3OamTable"
    cdot3Oamentry.EntityData.SegmentPath = "cdot3OamEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdot3Oamentry.Ifindex) + "']"
    cdot3Oamentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oamentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oamentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oamentry.EntityData.Children = make(map[string]types.YChild)
    cdot3Oamentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdot3Oamentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdot3Oamentry.Ifindex}
    cdot3Oamentry.EntityData.Leafs["cdot3OamAdminState"] = types.YLeaf{"Cdot3Oamadminstate", cdot3Oamentry.Cdot3Oamadminstate}
    cdot3Oamentry.EntityData.Leafs["cdot3OamOperStatus"] = types.YLeaf{"Cdot3Oamoperstatus", cdot3Oamentry.Cdot3Oamoperstatus}
    cdot3Oamentry.EntityData.Leafs["cdot3OamMode"] = types.YLeaf{"Cdot3Oammode", cdot3Oamentry.Cdot3Oammode}
    cdot3Oamentry.EntityData.Leafs["cdot3OamMaxOamPduSize"] = types.YLeaf{"Cdot3Oammaxoampdusize", cdot3Oamentry.Cdot3Oammaxoampdusize}
    cdot3Oamentry.EntityData.Leafs["cdot3OamConfigRevision"] = types.YLeaf{"Cdot3Oamconfigrevision", cdot3Oamentry.Cdot3Oamconfigrevision}
    cdot3Oamentry.EntityData.Leafs["cdot3OamFunctionsSupported"] = types.YLeaf{"Cdot3Oamfunctionssupported", cdot3Oamentry.Cdot3Oamfunctionssupported}
    return &(cdot3Oamentry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamadminstate represents to operate over this interface.  
type CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamadminstate string

const (
    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamadminstate_disabled CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamadminstate = "disabled"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamadminstate_enabled CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamadminstate = "enabled"
)

// CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oammode represents cdot3OamOperStatus was already operational(9).  
type CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oammode string

const (
    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oammode_active CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oammode = "active"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oammode_passive CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oammode = "passive"
)

// CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus represents in half-duplex operation.  
type CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus string

const (
    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_disabled CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "disabled"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_linkFault CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "linkFault"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_passiveWait CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "passiveWait"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_activeSendLocal CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "activeSendLocal"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_sendLocalAndRemote CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "sendLocalAndRemote"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_sendLocalAndRemoteOk CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "sendLocalAndRemoteOk"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_oamPeeringLocallyRejected CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "oamPeeringLocallyRejected"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_oamPeeringRemotelyRejected CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "oamPeeringRemotelyRejected"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_operational CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "operational"

    CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus_nonOperHalfDuplex CISCODOT3OAMMIB_Cdot3Oamtable_Cdot3Oamentry_Cdot3Oamoperstatus = "nonOperHalfDuplex"
)

// CISCODOT3OAMMIB_Cdot3Oampeertable
// This table contains information about the OAM peer for a
// particular Ethernet like interface. OAM entities communicate
// with a single OAM peer entity on Ethernet links on which OAM
// is enabled and operating properly.  There is one entry in this
// table for each entry in the cdot3OamTable for which information
// on the peer OAM entity is available.  
type CISCODOT3OAMMIB_Cdot3Oampeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information on the peer OAM entity for a
    // single Ethernet like interface.    Note that there is at most one OAM peer
    // for each Ethernet like interface.  Entries are automatically created when
    // information about the OAM peer entity becomes available, and automatically
    // deleted when the OAM peer entity is no longer in communication.  Peer
    // information is not available when cdot3OamOperStatus is disabled(1),
    // linkFault(2), passiveWait(3), activeSendLocal(4). or
    // nonOperHalfDuplex(10)). . The type is slice of
    // CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry.
    Cdot3Oampeerentry []CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry
}

func (cdot3Oampeertable *CISCODOT3OAMMIB_Cdot3Oampeertable) GetEntityData() *types.CommonEntityData {
    cdot3Oampeertable.EntityData.YFilter = cdot3Oampeertable.YFilter
    cdot3Oampeertable.EntityData.YangName = "cdot3OamPeerTable"
    cdot3Oampeertable.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oampeertable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3Oampeertable.EntityData.SegmentPath = "cdot3OamPeerTable"
    cdot3Oampeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oampeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oampeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oampeertable.EntityData.Children = make(map[string]types.YChild)
    cdot3Oampeertable.EntityData.Children["cdot3OamPeerEntry"] = types.YChild{"Cdot3Oampeerentry", nil}
    for i := range cdot3Oampeertable.Cdot3Oampeerentry {
        cdot3Oampeertable.EntityData.Children[types.GetSegmentPath(&cdot3Oampeertable.Cdot3Oampeerentry[i])] = types.YChild{"Cdot3Oampeerentry", &cdot3Oampeertable.Cdot3Oampeerentry[i]}
    }
    cdot3Oampeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdot3Oampeertable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry
// An entry in the table, containing information on the peer OAM
// entity for a single Ethernet like interface.  
// 
// Note that there is at most one OAM peer for each Ethernet like
// interface.  Entries are automatically created when information
// about the OAM peer entity becomes available, and automatically
// deleted when the OAM peer entity is no longer in
// communication.  Peer information is not available when
// cdot3OamOperStatus is disabled(1), linkFault(2),
// passiveWait(3), activeSendLocal(4). or nonOperHalfDuplex(10)). 
type CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The MAC address of the peer OAM entity.  The MAC address is derived from
    // the most recently received OAMPDU. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Cdot3Oampeermacaddress interface{}

    // The OUI of the OAM peer as reflected in the latest Information OAMPDU
    // received with a Local Information TLV.  The OUI can be used to identify the
    // vendor of the remote OAM entity.  This value is initialized to zero before
    // any Local Information TLV is received.  . The type is string with length:
    // 3.
    Cdot3Oampeervendoroui interface{}

    // The Vendor Info of the OAM peer as reflected in the latest Information
    // OAMPDU received with a Local Information TLV.  The vendor information field
    // is within the Local Information TLV, and can be used to determine
    // additional information about the peer entity.  The format of the vendor
    // information is unspecified within the 32-bit field.  This value is
    // initialized to zero before any Local Information TLV is received.  . The
    // type is interface{} with range: 0..4294967295.
    Cdot3Oampeervendorinfo interface{}

    // The mode of the OAM peer as reflected in the latest Information OAMPDU
    // received with a Local Information TLV.  The mode of the peer can be
    // determined from the Configuration field in the Local Information TLV of the
    // last Information OAMPDU received from the peer.  The value is unknown(3)
    // whenever no Local Information TLV has been received.  The values of
    // active(1) and passive(2) are returned when a Local Information TLV has been
    // received indicating the peer is in active or passive mode, respectively. .
    // The type is Cdot3Oampeermode.
    Cdot3Oampeermode interface{}

    // The maximum size of OAMPDU supported by the peer as reflected in the latest
    // Information OAMPDU received with a Local Information TLV.   Ethernet OAM on
    // this interface must not use OAMPDUs that exceed this size.  The maximum
    // OAMPDU size can be determined from the PDU Configuration field of the Local
    // Information TLV of the last Information OAMPDU received from the peer.  A
    // value of zero is returned if no Local Information TLV has been received. 
    // Otherwise, the value of the OAM peer's maximum OAMPDU size is returned in
    // this value.   Note that the values 1..63 are invalid sizes for Ethernet
    // frames and should never appear. . The type is interface{} with range:
    // 0..1518. Units are octets.
    Cdot3Oampeermaxoampdusize interface{}

    // The configuration revision of the OAM peer as reflected in the latest
    // OAMPDU.  This attribute is changed by the peer whenever it has a local
    // configuration change for Ethernet OAM this interface.  The configuration
    // revision can be determined from the Revision field of the Local Information
    // TLV of the most recently received Information OAMPDU with a Local
    // Information TLV. A value of zero is returned if no Local Information TLV
    // has been received.  . The type is interface{} with range: 0..65535.
    Cdot3Oampeerconfigrevision interface{}

    // The OAM functions supported on this Ethernet like interface. OAM consists
    // of separate functionality sets above the basic discovery process.  This
    // value indicates the capabilities of the peer OAM entity with respect to
    // these functions.  This value is initialized so all bits are clear.   If
    // unidirectionalSupport(0) is set, then the peer OAM entity supports sending
    // OAM frames on Ethernet interfaces when the receive path is known to be
    // inoperable.   If loopbackSupport(1) is set, then the peer OAM entity can
    // send and receive OAM loopback commands.  If eventSupport(2) is set, then
    // the peer OAM entity can send and receive event OAMPDUs to signal various
    // error conditions. If variableSupport(3) is set, then the peer OAM entity
    // can send and receive variable requests to monitor attribute value as
    // described in Clause 57 of [802.3ah].     The capabilities of the OAM peer
    // can be determined from the configuration field of the Local Information TLV
    // of the most recently received Information OAMPDU with a Local Information
    // TLV.  All zeros are returned if no Local Information TLV has yet been
    // received. . The type is map[string]bool.
    Cdot3Oampeerfunctionssupported interface{}
}

func (cdot3Oampeerentry *CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry) GetEntityData() *types.CommonEntityData {
    cdot3Oampeerentry.EntityData.YFilter = cdot3Oampeerentry.YFilter
    cdot3Oampeerentry.EntityData.YangName = "cdot3OamPeerEntry"
    cdot3Oampeerentry.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oampeerentry.EntityData.ParentYangName = "cdot3OamPeerTable"
    cdot3Oampeerentry.EntityData.SegmentPath = "cdot3OamPeerEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdot3Oampeerentry.Ifindex) + "']"
    cdot3Oampeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oampeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oampeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oampeerentry.EntityData.Children = make(map[string]types.YChild)
    cdot3Oampeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdot3Oampeerentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdot3Oampeerentry.Ifindex}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerMacAddress"] = types.YLeaf{"Cdot3Oampeermacaddress", cdot3Oampeerentry.Cdot3Oampeermacaddress}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerVendorOui"] = types.YLeaf{"Cdot3Oampeervendoroui", cdot3Oampeerentry.Cdot3Oampeervendoroui}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerVendorInfo"] = types.YLeaf{"Cdot3Oampeervendorinfo", cdot3Oampeerentry.Cdot3Oampeervendorinfo}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerMode"] = types.YLeaf{"Cdot3Oampeermode", cdot3Oampeerentry.Cdot3Oampeermode}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerMaxOamPduSize"] = types.YLeaf{"Cdot3Oampeermaxoampdusize", cdot3Oampeerentry.Cdot3Oampeermaxoampdusize}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerConfigRevision"] = types.YLeaf{"Cdot3Oampeerconfigrevision", cdot3Oampeerentry.Cdot3Oampeerconfigrevision}
    cdot3Oampeerentry.EntityData.Leafs["cdot3OamPeerFunctionsSupported"] = types.YLeaf{"Cdot3Oampeerfunctionssupported", cdot3Oampeerentry.Cdot3Oampeerfunctionssupported}
    return &(cdot3Oampeerentry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode represents active or passive mode, respectively. 
type CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode string

const (
    CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode_active CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode = "active"

    CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode_passive CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode = "passive"

    CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode_unknown CISCODOT3OAMMIB_Cdot3Oampeertable_Cdot3Oampeerentry_Cdot3Oampeermode = "unknown"
)

// CISCODOT3OAMMIB_Cdot3Oamloopbacktable
// This table contains controls for the loopback state of the
// local link as well as indicating the status of the loopback
// function.  There is one entry in this table for each entry in
// cdot3OamTable that supports loopback functionality (where
// cdot3OamFunctionsSupported includes the loopbackSupport bit
// set).
// 
// Loopback can be used to place the remote OAM entity in a state
// where every received frame (except OAMPDUs) is echoed back
// over the same interface on which they were received.   In this
// state, at the remote entity, 'normal' traffic is disabled as
// only the looped back frames are transmitted on the interface.
// Loopback is thus an intrusive operation that prohibits normal
// data flow and should be used accordingly.  
type CISCODOT3OAMMIB_Cdot3Oamloopbacktable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information on the loopback status for a
    // single Ethernet like interface.  Entries in the table are automatically
    // created whenever the local OAM entity supports loopback capabilities.  The
    // loopback status on the interface can be determined from the
    // cdot3OamLoopbackStatus object.  . The type is slice of
    // CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry.
    Cdot3Oamloopbackentry []CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry
}

func (cdot3Oamloopbacktable *CISCODOT3OAMMIB_Cdot3Oamloopbacktable) GetEntityData() *types.CommonEntityData {
    cdot3Oamloopbacktable.EntityData.YFilter = cdot3Oamloopbacktable.YFilter
    cdot3Oamloopbacktable.EntityData.YangName = "cdot3OamLoopbackTable"
    cdot3Oamloopbacktable.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oamloopbacktable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3Oamloopbacktable.EntityData.SegmentPath = "cdot3OamLoopbackTable"
    cdot3Oamloopbacktable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oamloopbacktable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oamloopbacktable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oamloopbacktable.EntityData.Children = make(map[string]types.YChild)
    cdot3Oamloopbacktable.EntityData.Children["cdot3OamLoopbackEntry"] = types.YChild{"Cdot3Oamloopbackentry", nil}
    for i := range cdot3Oamloopbacktable.Cdot3Oamloopbackentry {
        cdot3Oamloopbacktable.EntityData.Children[types.GetSegmentPath(&cdot3Oamloopbacktable.Cdot3Oamloopbackentry[i])] = types.YChild{"Cdot3Oamloopbackentry", &cdot3Oamloopbacktable.Cdot3Oamloopbackentry[i]}
    }
    cdot3Oamloopbacktable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdot3Oamloopbacktable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry
// An entry in the table, containing information on the loopback
// status for a single Ethernet like interface.  Entries in the
// table are automatically created whenever the local OAM entity
// supports loopback capabilities.  The loopback status on the
// interface can be determined from the cdot3OamLoopbackStatus
// object.  
type CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The loopback status of the OAM entity.  This status is determined by a
    // combination of the local parser and multiplexer states, the remote parser
    // and multiplexer states, as well as by the actions of the local OAM client. 
    // When operating in normal mode with no loopback in progress, the status
    // reads noLoopback(1).    The values initiatingLooopback(2) and
    // terminatingLoopback(4) can be read or written.  The other values can only
    // be read - they can never be written.  Writing initiatingLoopback causes the
    // local OAM entity to start the loopback process with its peer.  This value
    // can only be written when the status is noLoopback(1).  Writing the value
    // initiatingLoopback(2) in any other state has no effect.  When in
    // remoteLoopback(3), writing terminatingLoopback(4) causes the local OAM
    // entity to initiate the termination of the loopback state.  Writing
    // terminatingLoopack(4) in any other state has no effect.                    
    // If the OAM client initiates a looopback and has sent an Loopback OAMPDU and
    // is waiting for a response, where the local parser and multiplexer states
    // are DISCARD (see [802.3ah, 57.2.11.1]), the status is 'initiatingLoopback'.
    // In this case, the local OAM entity has yet to receive any acknowledgement
    // that the remote OAM entity has received its loopback command request.      
    // If the local OAM client knows that the remote OAM entity is in loopback
    // mode (via the remote state information as described in [802.3ah, 57.2.11.1,
    // 30.3.6.1.15]), the status is remoteLoopback(3).  If the local OAM client is
    // in the process of terminating the remote loopback [802.3ah, 57.2.11.3,
    // 30.3.6.1.14], with its local multiplexer and parser states in DISCARD, the
    // status is terminatingLoopback(4).  If the remote OAM client has put the
    // local OAM entity in loopback mode as indicated by its local parser state,
    // the status is localLoopback(5).    The unknown(6) status indicates the
    // parser and multiplexer combination is unexpected.  This status may be
    // returned if the OAM loopback is in a transition state but should not
    // persist.   The values of this attribute correspond to the following values
    // of the local and remote parser and multiplexer states.     value           
    // LclPrsr   LclMux    RmtPrsr   RmtMux   noLoopback         FWD       FWD    
    // FWD       FWD     initLoopback     DISCARD   DISCARD     FWD       FWD   
    // rmtLoopback      DISCARD     FWD      LPBK    DISCARD   tmtngLoopback   
    // DISCARD   DISCARD    LPBK    DISCARD   lclLoopback        LPBK    DISCARD  
    // DISCARD     FWD   unknown            ***   any other combination   ***. The
    // type is Cdot3Oamloopbackstatus.
    Cdot3Oamloopbackstatus interface{}

    // Since OAM loopback is a disruptive operation (user traffic does not pass),
    // this attribute provides a mechanism to provide controls over whether
    // received OAM loopback commands are processed or ignored.  When the value is
    // ignore(1), received loopback commands are ignored.  When the value is
    // process(2), OAM loopback commands are processed.  The default value is to
    // ignore loopback commands (ignore(1)).  . The type is
    // Cdot3Oamloopbackignorerx.
    Cdot3Oamloopbackignorerx interface{}
}

func (cdot3Oamloopbackentry *CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry) GetEntityData() *types.CommonEntityData {
    cdot3Oamloopbackentry.EntityData.YFilter = cdot3Oamloopbackentry.YFilter
    cdot3Oamloopbackentry.EntityData.YangName = "cdot3OamLoopbackEntry"
    cdot3Oamloopbackentry.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oamloopbackentry.EntityData.ParentYangName = "cdot3OamLoopbackTable"
    cdot3Oamloopbackentry.EntityData.SegmentPath = "cdot3OamLoopbackEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdot3Oamloopbackentry.Ifindex) + "']"
    cdot3Oamloopbackentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oamloopbackentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oamloopbackentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oamloopbackentry.EntityData.Children = make(map[string]types.YChild)
    cdot3Oamloopbackentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdot3Oamloopbackentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdot3Oamloopbackentry.Ifindex}
    cdot3Oamloopbackentry.EntityData.Leafs["cdot3OamLoopbackStatus"] = types.YLeaf{"Cdot3Oamloopbackstatus", cdot3Oamloopbackentry.Cdot3Oamloopbackstatus}
    cdot3Oamloopbackentry.EntityData.Leafs["cdot3OamLoopbackIgnoreRx"] = types.YLeaf{"Cdot3Oamloopbackignorerx", cdot3Oamloopbackentry.Cdot3Oamloopbackignorerx}
    return &(cdot3Oamloopbackentry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackignorerx represents ignore loopback commands (ignore(1)).  
type CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackignorerx string

const (
    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackignorerx_ignore CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackignorerx = "ignore"

    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackignorerx_process CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackignorerx = "process"
)

// CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus represents   unknown            ***   any other combination   ***
type CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus string

const (
    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus_noLoopback CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus = "noLoopback"

    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus_initiatingLoopback CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus = "initiatingLoopback"

    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus_remoteLoopback CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus = "remoteLoopback"

    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus_terminatingLoopback CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus = "terminatingLoopback"

    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus_localLoopback CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus = "localLoopback"

    CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus_unknown CISCODOT3OAMMIB_Cdot3Oamloopbacktable_Cdot3Oamloopbackentry_Cdot3Oamloopbackstatus = "unknown"
)

// CISCODOT3OAMMIB_Cdot3Oamstatstable
// This table contains statistics for the OAM function on a
// particular Ethernet like interface. There is an entry in the
// table for every entry in the cdot3OamTable. 
// 
// The counters in this table are defined as 32-bit entries to
// match the counter size as defined in [802.3ah].  Given the OAM
// protocol is a slow protocol, the counters increment at a slow
// rate. 
type CISCODOT3OAMMIB_Cdot3Oamstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing statistics information on the Ethernet
    // OAM function for a single Ethernet like interface.  Entries are
    // automatically created for every entry in the cdot3OamTable.  Counters are
    // maintained across transitions in cdot3OamOperStatus.  . The type is slice
    // of CISCODOT3OAMMIB_Cdot3Oamstatstable_Cdot3Oamstatsentry.
    Cdot3Oamstatsentry []CISCODOT3OAMMIB_Cdot3Oamstatstable_Cdot3Oamstatsentry
}

func (cdot3Oamstatstable *CISCODOT3OAMMIB_Cdot3Oamstatstable) GetEntityData() *types.CommonEntityData {
    cdot3Oamstatstable.EntityData.YFilter = cdot3Oamstatstable.YFilter
    cdot3Oamstatstable.EntityData.YangName = "cdot3OamStatsTable"
    cdot3Oamstatstable.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oamstatstable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3Oamstatstable.EntityData.SegmentPath = "cdot3OamStatsTable"
    cdot3Oamstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oamstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oamstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oamstatstable.EntityData.Children = make(map[string]types.YChild)
    cdot3Oamstatstable.EntityData.Children["cdot3OamStatsEntry"] = types.YChild{"Cdot3Oamstatsentry", nil}
    for i := range cdot3Oamstatstable.Cdot3Oamstatsentry {
        cdot3Oamstatstable.EntityData.Children[types.GetSegmentPath(&cdot3Oamstatstable.Cdot3Oamstatsentry[i])] = types.YChild{"Cdot3Oamstatsentry", &cdot3Oamstatstable.Cdot3Oamstatsentry[i]}
    }
    cdot3Oamstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdot3Oamstatstable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oamstatstable_Cdot3Oamstatsentry
// An entry in the table, containing statistics information on
// the Ethernet OAM function for a single Ethernet like
// interface.  Entries are automatically created for every entry
// in the cdot3OamTable.  Counters are maintained across
// transitions in cdot3OamOperStatus.  
type CISCODOT3OAMMIB_Cdot3Oamstatstable_Cdot3Oamstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // A count of the number of Information OAMPDUs transmitted on this interface.
    // Discontinuities of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of the
    // ifCounterDiscontinuityTime.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3Oaminformationtx interface{}

    // A count of the number of Information OAMPDUs received on this interface. 
    // Discontinuities of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of the
    // ifCounterDiscontinuityTime.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3Oaminformationrx interface{}

    // A count of the number of unique Event OAMPDUs transmitted on this
    // interface.  Event notifications may be sent in duplicate to increase the
    // probability of successfully being received, given the possibility that a
    // frame may be lost in transit. Duplicate Event Notification transmissions
    // are counted by cdot3OamDuplicateEventNotificationTx.    A unique Event
    // Notification OAMPDU is indicated as an Event Notification OAMPDU with a
    // Sequence Number field that is distinct from the previously transmitted
    // Event Notification OAMPDU Sequence Number.    Discontinuities of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of the ifCounterDiscontinuityTime.  .
    // The type is interface{} with range: 0..4294967295. Units are frames.
    Cdot3Oamuniqueeventnotificationtx interface{}

    // A count of the number of unique Event OAMPDUs received on this interface. 
    // Event notification OAMPDUs may be sent in duplicate to increase the
    // probability of successfully being received, given the possibility that a
    // frame may be lost in transit.  Duplicate Event Notification receptions are
    // counted by cdot3OamDuplicateEventNotificationRx.    A unique Event
    // Notification OAMPDU is indicated as an Event Notification OAMPDU with a
    // Sequence Number field that is distinct from the previously received Event
    // Notification OAMPDU Sequence Number.    Discontinuities of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of the ifCounterDiscontinuityTime.  . The type is
    // interface{} with range: 0..4294967295. Units are frames.
    Cdot3Oamuniqueeventnotificationrx interface{}

    // A count of the number of duplicate Event OAMPDUs transmitted on this
    // interface.  Event notification OAMPDUs may be sent in duplicate to increase
    // the probability of successfully being received, given the possibility that
    // a frame may be lost in transit.    A duplicate Event Notification OAMPDU is
    // indicated as an Event Notification OAMPDU with a Sequence Number field that
    // is identical to the previously transmitted Event Notification OAMPDU
    // Sequence Number.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamduplicateeventnotificationtx interface{}

    // A count of the number of duplicate Event OAMPDUs received on this
    // interface.  Event notification OAMPDUs may be sent in duplicate to increase
    // the probability of successfully being received, given the possibility that
    // a frame may be lost in transit.    A duplicate Event Notification OAMPDU is
    // indicated as an Event Notification OAMPDU with a Sequence Number field that
    // is identical to the previously received Event Notification OAMPDU Sequence
    // Number.    Discontinuities of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // the ifCounterDiscontinuityTime.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3Oamduplicateeventnotificationrx interface{}

    // A count of the number of Loopback Control OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamloopbackcontroltx interface{}

    // A count of the number of Loopback Control OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamloopbackcontrolrx interface{}

    // A count of the number of Variable Request OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamvariablerequesttx interface{}

    // A count of the number of Variable Request OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamvariablerequestrx interface{}

    // A count of the number of Variable Response OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamvariableresponsetx interface{}

    // A count of the number of Variable Response OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamvariableresponserx interface{}

    // A count of the number of Organization Specific OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamorgspecifictx interface{}

    // A count of the number of Organization Specific OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamorgspecificrx interface{}

    // A count of the number of OAMPDUs transmitted on this interface with an
    // unsupported op-code.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamunsupportedcodestx interface{}

    // A count of the number of OAMPDUs received on this interface with an
    // unsupported op-code.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3Oamunsupportedcodesrx interface{}

    // A count of the number of frames that were dropped by the OAM multiplexer. 
    // Since the OAM multiplexer has multiple inputs and a single output, there
    // may be cases where frames are dropped due to transmit resource contention. 
    // This counter is incremented whenever a frame is dropped by the OAM layer.
    // Note that any Ethernet frame, not just OAMPDUs, may be dropped by the OAM
    // layer.  This can occur when an OAMPDU takes precedence over a 'normal'
    // frame resulting in the 'normal' frame being dropped.    When this counter
    // is incremented, no other counters in this MIB are incremented.             
    // Discontinuities of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of the
    // ifCounterDiscontinuityTime.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3Oamframeslostduetooam interface{}
}

func (cdot3Oamstatsentry *CISCODOT3OAMMIB_Cdot3Oamstatstable_Cdot3Oamstatsentry) GetEntityData() *types.CommonEntityData {
    cdot3Oamstatsentry.EntityData.YFilter = cdot3Oamstatsentry.YFilter
    cdot3Oamstatsentry.EntityData.YangName = "cdot3OamStatsEntry"
    cdot3Oamstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oamstatsentry.EntityData.ParentYangName = "cdot3OamStatsTable"
    cdot3Oamstatsentry.EntityData.SegmentPath = "cdot3OamStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdot3Oamstatsentry.Ifindex) + "']"
    cdot3Oamstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oamstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oamstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oamstatsentry.EntityData.Children = make(map[string]types.YChild)
    cdot3Oamstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdot3Oamstatsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdot3Oamstatsentry.Ifindex}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamInformationTx"] = types.YLeaf{"Cdot3Oaminformationtx", cdot3Oamstatsentry.Cdot3Oaminformationtx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamInformationRx"] = types.YLeaf{"Cdot3Oaminformationrx", cdot3Oamstatsentry.Cdot3Oaminformationrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamUniqueEventNotificationTx"] = types.YLeaf{"Cdot3Oamuniqueeventnotificationtx", cdot3Oamstatsentry.Cdot3Oamuniqueeventnotificationtx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamUniqueEventNotificationRx"] = types.YLeaf{"Cdot3Oamuniqueeventnotificationrx", cdot3Oamstatsentry.Cdot3Oamuniqueeventnotificationrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamDuplicateEventNotificationTx"] = types.YLeaf{"Cdot3Oamduplicateeventnotificationtx", cdot3Oamstatsentry.Cdot3Oamduplicateeventnotificationtx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamDuplicateEventNotificationRx"] = types.YLeaf{"Cdot3Oamduplicateeventnotificationrx", cdot3Oamstatsentry.Cdot3Oamduplicateeventnotificationrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamLoopbackControlTx"] = types.YLeaf{"Cdot3Oamloopbackcontroltx", cdot3Oamstatsentry.Cdot3Oamloopbackcontroltx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamLoopbackControlRx"] = types.YLeaf{"Cdot3Oamloopbackcontrolrx", cdot3Oamstatsentry.Cdot3Oamloopbackcontrolrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamVariableRequestTx"] = types.YLeaf{"Cdot3Oamvariablerequesttx", cdot3Oamstatsentry.Cdot3Oamvariablerequesttx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamVariableRequestRx"] = types.YLeaf{"Cdot3Oamvariablerequestrx", cdot3Oamstatsentry.Cdot3Oamvariablerequestrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamVariableResponseTx"] = types.YLeaf{"Cdot3Oamvariableresponsetx", cdot3Oamstatsentry.Cdot3Oamvariableresponsetx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamVariableResponseRx"] = types.YLeaf{"Cdot3Oamvariableresponserx", cdot3Oamstatsentry.Cdot3Oamvariableresponserx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamOrgSpecificTx"] = types.YLeaf{"Cdot3Oamorgspecifictx", cdot3Oamstatsentry.Cdot3Oamorgspecifictx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamOrgSpecificRx"] = types.YLeaf{"Cdot3Oamorgspecificrx", cdot3Oamstatsentry.Cdot3Oamorgspecificrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamUnsupportedCodesTx"] = types.YLeaf{"Cdot3Oamunsupportedcodestx", cdot3Oamstatsentry.Cdot3Oamunsupportedcodestx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamUnsupportedCodesRx"] = types.YLeaf{"Cdot3Oamunsupportedcodesrx", cdot3Oamstatsentry.Cdot3Oamunsupportedcodesrx}
    cdot3Oamstatsentry.EntityData.Leafs["cdot3OamFramesLostDueToOam"] = types.YLeaf{"Cdot3Oamframeslostduetooam", cdot3Oamstatsentry.Cdot3Oamframeslostduetooam}
    return &(cdot3Oamstatsentry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oameventconfigtable
// Ethernet OAM includes the ability to generate and receive
// Event Notification OAMPDUs to indicate various link problems.
// This table contains the mechanisms to enable Event
// Notifications and configure the thresholds to generate the
// standard Ethernet OAM events.  There is one entry in the table
// for every entry in cdot3OamTable that supports OAM events
// (where cdot3OamFunctionsSupported includes the eventSupport
// bit set). The values in the table are maintained across
// changes to cdot3OamOperStatus.  
// 
// The standard threshold crossing events are:
//   - Errored Symbol Period Event.  Generated when the number of
//     symbol errors exceeds a threshold within a given window 
//     defined by a number of symbols (for example, 1,000 symbols
//     out of 1,000,000 had errors).  
//   - Errored Frame Period Event.  Generated when the number of 
//     frame errors exceeds a threshold within a given window 
//     defined by a number of frames (for example, 10 frames out
//     of 1000 had errors).  
//   - Errored Frame Event.  Generated when the number of frame 
//     errors exceeds a threshold within a given window defined 
//     by a period of time (for example, 10 frames in 1 second 
//     had errors).
//   - Errored Frame Seconds Summary Event.  Generated when the 
//     number of errored frame seconds exceeds a threshold within
//     a given time period (for example, 10 errored frame seconds
//     within the last 100 seconds).  An errored frame second is 
//     defined as a 1 second interval which had >0 frame errors.  
// There are other events (dying gasp, critical events) that are
// not threshold crossing events but which can be
// enabled/disabled via this table.  
type CISCODOT3OAMMIB_Cdot3Oameventconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entries are automatically created and deleted from this table, and exist
    // whenever the OAM entity supports Ethernet OAM events (as indicated by the
    // eventSupport bit in cdot3OamFunctionsSuppported).  Values in the table are
    // maintained across changes to the value of cdot3OamOperStatus.  Event
    // configuration controls when the local management entity sends Event
    // Notification OAMPDUs to its OAM peer, and when certain event flags are set
    // or cleared in OAMPDUs. . The type is slice of
    // CISCODOT3OAMMIB_Cdot3Oameventconfigtable_Cdot3Oameventconfigentry.
    Cdot3Oameventconfigentry []CISCODOT3OAMMIB_Cdot3Oameventconfigtable_Cdot3Oameventconfigentry
}

func (cdot3Oameventconfigtable *CISCODOT3OAMMIB_Cdot3Oameventconfigtable) GetEntityData() *types.CommonEntityData {
    cdot3Oameventconfigtable.EntityData.YFilter = cdot3Oameventconfigtable.YFilter
    cdot3Oameventconfigtable.EntityData.YangName = "cdot3OamEventConfigTable"
    cdot3Oameventconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oameventconfigtable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3Oameventconfigtable.EntityData.SegmentPath = "cdot3OamEventConfigTable"
    cdot3Oameventconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oameventconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oameventconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oameventconfigtable.EntityData.Children = make(map[string]types.YChild)
    cdot3Oameventconfigtable.EntityData.Children["cdot3OamEventConfigEntry"] = types.YChild{"Cdot3Oameventconfigentry", nil}
    for i := range cdot3Oameventconfigtable.Cdot3Oameventconfigentry {
        cdot3Oameventconfigtable.EntityData.Children[types.GetSegmentPath(&cdot3Oameventconfigtable.Cdot3Oameventconfigentry[i])] = types.YChild{"Cdot3Oameventconfigentry", &cdot3Oameventconfigtable.Cdot3Oameventconfigentry[i]}
    }
    cdot3Oameventconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdot3Oameventconfigtable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oameventconfigtable_Cdot3Oameventconfigentry
// Entries are automatically created and deleted from this
// table, and exist whenever the OAM entity supports Ethernet OAM
// events (as indicated by the eventSupport bit in
// cdot3OamFunctionsSuppported).  Values in the table are
// maintained across changes to the value of cdot3OamOperStatus.
// 
// Event configuration controls when the local management entity
// sends Event Notification OAMPDUs to its OAM peer, and when
// certain event flags are set or cleared in OAMPDUs. 
type CISCODOT3OAMMIB_Cdot3Oameventconfigtable_Cdot3Oameventconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The two objects cdot3OamErrSymPeriodWindowHi and cdot3OamErrSymPeriodLo
    // together form an unsigned 64-bit integer representing the number of symbols
    // over which this threshold event is defined.  This is defined as 
    // cdot3OamErrSymPeriodWindow = ((2^32)*cdot3OamErrSymPeriodWindowHi)         
    // + cdot3OamErrSymPeriodWindowLo  If cdot3OamErrSymPeriodThreshold symbol
    // errors occur within a window of cdot3OamErrSymPeriodWindow symbols, an
    // Event Notification OAMPDU should be generated with an Errored Symbol Period
    // Event TLV indicating the threshold has been crossed in this window.    The
    // default value for cdot3OamErrSymPeriodWindow is the number of symbols in
    // one second for the underlying physical layer. The type is interface{} with
    // range: 0..4294967295. Units are 2^32 symbols.
    Cdot3Oamerrsymperiodwindowhi interface{}

    // The two objects cdot3OamErrSymPeriodWindowHi and
    // cdot3OamErrSymPeriodWindowLo together form an unsigned 64-bit integer
    // representing the number of symbols over which this threshold event is
    // defined.  This is defined as  cdot3OamErrSymPeriodWindow =
    // ((2^32)*cdot3OamErrSymPeriodWindowHi)                                 +
    // cdot3OamErrSymPeriodWindowLo  If cdot3OamErrSymPeriodThreshold symbol
    // errors occur within a window of cdot3OamErrSymPeriodWindow symbols, an
    // Event Notification OAMPDU should be generated with an Errored Symbol Period
    // Event TLV indicating the threshold has been crossed in this window.    The
    // default value for cdot3OamErrSymPeriodWindow is the number of symbols in
    // one second for the underlying physical layer. . The type is interface{}
    // with range: 0..4294967295. Units are symbols.
    Cdot3Oamerrsymperiodwindowlo interface{}

    // The two objects cdot3OamErrSymPeriodThresholdHi and
    // cdot3OamErrSymPeriodThresholdLo together form an unsigned 64-bit integer
    // representing the number of symbol errors that must occur within a given
    // window to cause this event.    This is defined as   
    // cdot3OamErrSymPeriodThreshold =                     ((2^32) *
    // cdot3OamErrSymPeriodThresholdHi)                             +
    // cdot3OamErrSymPeriodThresholdLo                     If
    // cdot3OamErrSymPeriodThreshold symbol errors occur within a window of
    // cdot3OamErrSymPeriodWindow symbols, an Event Notification OAMPDU should be
    // generated with an Errored Symbol Period Event TLV indicating the threshold
    // has been crossed in this window.    The default value for
    // cdot3OamErrSymPeriodThreshold is one symbol errors.  If the threshold value
    // is zero, then an Event Notification OAMPDU is sent periodically (at the end
    // of every window).  This can be used as an asynchronous notification to the
    // peer OAM entity of the statistics related to this threshold crossing alarm.
    // . The type is interface{} with range: 0..4294967295. Units are 2^32
    // symbols.
    Cdot3Oamerrsymperiodthresholdhi interface{}

    // The two objects cdot3OamErrSymPeriodThresholdHi and
    // cdot3OamErrSymPeriodThresholdLo together form an unsigned 64-bit integer
    // representing the number of symbol errors that must occur within a given
    // window to cause this event.    This is defined as   
    // cdot3OamErrSymPeriodThreshold =                     ((2^32) *
    // cdot3OamErrSymPeriodThresholdHi)                             +
    // cdot3OamErrSymPeriodThresholdLo                     If
    // cdot3OamErrSymPeriodThreshold symbol errors occur within a window of
    // cdot3OamErrSymPeriodWindow symbols, an Event Notification OAMPDU should be
    // generated with an Errored Symbol Period Event TLV indicating the threshold
    // has been crossed in this window.    The default value for
    // cdot3OamErrSymPeriodThreshold is one symbol error. If the threshold value
    // is zero, then an Event Notification OAMPDU is sent periodically (at the end
    // of every window).  This can be used as an asynchronous notification to the
    // peer OAM entity of the statistics related to this threshold crossing alarm.
    // . The type is interface{} with range: 0..4294967295. Units are symbols.
    Cdot3Oamerrsymperiodthresholdlo interface{}

    // If true, the OAM entity should send an Event Notification OAMPDU when an
    // Errored Symbol Period Event occurs.  By default, this object should have
    // the value true for Ethernet like interfaces that support OAM.  If the OAM
    // layer does not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. The type is
    // bool.
    Cdot3Oamerrsymperiodevnotifenable interface{}

    // The number of frames over which the threshold is defined. The default value
    // of the window is the number of minimum size Ethernet frames that can be
    // received over the physical layer in one second.                    If
    // cdot3OamErrFramePeriodThreshold frame errors occur within a window of
    // cdot3OamErrFramePeriodWindow frames, an Event Notification OAMPDU should be
    // generated with an Errored Frame Period Event TLV indicating the threshold
    // has been crossed in this window.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3Oamerrframeperiodwindow interface{}

    // The number of frame errors that must occur for this event to be triggered. 
    // The default value is one frame error.  If the threshold value is zero, then
    // an Event Notification OAMPDU is sent periodically (at the end of every
    // window).  This can be used as an asynchronous notification to the peer OAM
    // entity of the statistics related to this threshold crossing alarm.         
    // If cdot3OamErrFramePeriodThreshold frame errors occur within a window of
    // cdot3OamErrFramePeriodWindow frames, an Event Notification OAMPDU should be
    // generated with an Errored Frame Period Event TLV indicating the threshold
    // has been crossed in this window.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3Oamerrframeperiodthreshold interface{}

    // If true, the OAM entity should send an Event Notification OAMPDU when an
    // Errored Frame Period Event occurs.   By default, this object should have
    // the value true for Ethernet like interfaces that support OAM.  If the OAM
    // layer does not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. . The type is
    // bool.
    Cdot3Oamerrframeperiodevnotifenable interface{}

    // The amount of time (in 100ms increments) over which the threshold is
    // defined.  The default value is 10 (1 second).                    If
    // cdot3OamErrFrameThreshold frame errors occur within a window of
    // cdot3OamErrFrameWindow seconds (measured in tenths of seconds), an Event
    // Notification OAMPDU should be generated with an Errored Frame Event TLV
    // indicating the threshold has been crossed in this window.  . The type is
    // interface{} with range: 0..4294967295. Units are tenths of a second.
    Cdot3Oamerrframewindow interface{}

    // The number of frame errors that must occur for this event to be triggered. 
    // The default value is one frame error. If the threshold value is zero, then
    // an Event Notification OAMPDU is sent periodically (at the end of every
    // window).  This can be used as an asynchronous notification to the peer OAM
    // entity of the statistics related to this threshold crossing alarm.         
    // If cdot3OamErrFrameThreshold frame errors occur within a window of
    // cdot3OamErrFrameWindow (in tenths of seconds), an Event Notification OAMPDU
    // should be generated with an Errored Frame Event TLV indicating the
    // threshold has been crossed in this window.  . The type is interface{} with
    // range: 0..4294967295. Units are frames.
    Cdot3Oamerrframethreshold interface{}

    // If true, the OAM entity should send an Event Notification OAMPDU when an
    // Errored Frame Event occurs.                   By default, this object
    // should have the value true for Ethernet like interfaces that support OAM. 
    // If the OAM layer does not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. . The type is
    // bool.
    Cdot3Oamerrframeevnotifenable interface{}

    // The amount of time (in 100ms intervals) over which the threshold is
    // defined.  The default value is 100 (10 seconds).                    If
    // cdot3OamErrFrameSecsSummaryThreshold frame errors occur within a window of
    // cdot3OamErrFrameSecsSummaryWindow (in tenths of seconds), an Event
    // Notification OAMPDU should be generated with an Errored Frame Seconds
    // Summary Event TLV indicating the threshold has been crossed in this window.
    // . The type is interface{} with range: 100..9000. Units are tenths of a
    // second.
    Cdot3Oamerrframesecssummarywindow interface{}

    // The number of errored frame seconds that must occur for this event to be
    // triggered.  The default value is one errored frame second. If the threshold
    // value is zero, then an Event Notification OAMPDU is sent periodically (at
    // the end of every window).  This can be used as an asynchronous notification
    // to the peer OAM entity of the statistics related to this threshold crossing
    // alarm.                   If cdot3OamErrFrameSecsSummaryThreshold frame
    // errors occur within a window of cdot3OamErrFrameSecsSummaryWindow (in
    // tenths of seconds), an Event Notification OAMPDU should be generated with
    // an Errored Frame Seconds Summary Event TLV indicating the threshold has
    // been crossed in this window.  . The type is interface{} with range: 1..900.
    // Units are errored frame seconds.
    Cdot3Oamerrframesecssummarythreshold interface{}

    // If true, the local OAM entity should send an Event Notification OAMPDU when
    // an Errored Frame Seconds Event occurs.                   By default, this
    // object should have the value true for Ethernet like interfaces that support
    // OAM.  If the OAM layer does not support event notifications (as indicated
    // via the cdot3OamFunctionsSupported attribute), this value is ignored. The
    // type is bool.
    Cdot3Oamerrframesecsevnotifenable interface{}

    // If true, the local OAM entity should attempt to indicate a dying gasp via
    // the OAMPDU flags field to its peer OAM entity when a dying gasp event
    // occurs.  The exact definition of a dying gasp event is implementation
    // dependent.  If the system does not support dying gasp capability, setting
    // this object has no effect, and reading the object should always result in
    // 'false'.                    By default, this object should have the value
    // true for Ethernet like interfaces that support OAM.  If the OAM layer does
    // not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. The type is
    // bool.
    Cdot3Oamdyinggaspenable interface{}

    // If true, the local OAM entity should attempt to indicate a critical event
    // via the OAMPDU flags to its peer OAM entity when a critical event occurs.  
    // The exact definition of a critical event is implementation dependent.  If
    // the system does not support critical event capability, setting this object
    // has no effect, and reading the object should always result in 'false'.     
    // By default, this object should have the value true for Ethernet like
    // interfaces that support OAM.  If the OAM layer does not support event
    // notifications (as indicated via the cdot3OamFunctionsSupported attribute),
    // this value is ignored. The type is bool.
    Cdot3Oamcriticaleventenable interface{}
}

func (cdot3Oameventconfigentry *CISCODOT3OAMMIB_Cdot3Oameventconfigtable_Cdot3Oameventconfigentry) GetEntityData() *types.CommonEntityData {
    cdot3Oameventconfigentry.EntityData.YFilter = cdot3Oameventconfigentry.YFilter
    cdot3Oameventconfigentry.EntityData.YangName = "cdot3OamEventConfigEntry"
    cdot3Oameventconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oameventconfigentry.EntityData.ParentYangName = "cdot3OamEventConfigTable"
    cdot3Oameventconfigentry.EntityData.SegmentPath = "cdot3OamEventConfigEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdot3Oameventconfigentry.Ifindex) + "']"
    cdot3Oameventconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oameventconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oameventconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oameventconfigentry.EntityData.Children = make(map[string]types.YChild)
    cdot3Oameventconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdot3Oameventconfigentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdot3Oameventconfigentry.Ifindex}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrSymPeriodWindowHi"] = types.YLeaf{"Cdot3Oamerrsymperiodwindowhi", cdot3Oameventconfigentry.Cdot3Oamerrsymperiodwindowhi}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrSymPeriodWindowLo"] = types.YLeaf{"Cdot3Oamerrsymperiodwindowlo", cdot3Oameventconfigentry.Cdot3Oamerrsymperiodwindowlo}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrSymPeriodThresholdHi"] = types.YLeaf{"Cdot3Oamerrsymperiodthresholdhi", cdot3Oameventconfigentry.Cdot3Oamerrsymperiodthresholdhi}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrSymPeriodThresholdLo"] = types.YLeaf{"Cdot3Oamerrsymperiodthresholdlo", cdot3Oameventconfigentry.Cdot3Oamerrsymperiodthresholdlo}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrSymPeriodEvNotifEnable"] = types.YLeaf{"Cdot3Oamerrsymperiodevnotifenable", cdot3Oameventconfigentry.Cdot3Oamerrsymperiodevnotifenable}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFramePeriodWindow"] = types.YLeaf{"Cdot3Oamerrframeperiodwindow", cdot3Oameventconfigentry.Cdot3Oamerrframeperiodwindow}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFramePeriodThreshold"] = types.YLeaf{"Cdot3Oamerrframeperiodthreshold", cdot3Oameventconfigentry.Cdot3Oamerrframeperiodthreshold}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFramePeriodEvNotifEnable"] = types.YLeaf{"Cdot3Oamerrframeperiodevnotifenable", cdot3Oameventconfigentry.Cdot3Oamerrframeperiodevnotifenable}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFrameWindow"] = types.YLeaf{"Cdot3Oamerrframewindow", cdot3Oameventconfigentry.Cdot3Oamerrframewindow}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFrameThreshold"] = types.YLeaf{"Cdot3Oamerrframethreshold", cdot3Oameventconfigentry.Cdot3Oamerrframethreshold}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFrameEvNotifEnable"] = types.YLeaf{"Cdot3Oamerrframeevnotifenable", cdot3Oameventconfigentry.Cdot3Oamerrframeevnotifenable}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFrameSecsSummaryWindow"] = types.YLeaf{"Cdot3Oamerrframesecssummarywindow", cdot3Oameventconfigentry.Cdot3Oamerrframesecssummarywindow}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFrameSecsSummaryThreshold"] = types.YLeaf{"Cdot3Oamerrframesecssummarythreshold", cdot3Oameventconfigentry.Cdot3Oamerrframesecssummarythreshold}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamErrFrameSecsEvNotifEnable"] = types.YLeaf{"Cdot3Oamerrframesecsevnotifenable", cdot3Oameventconfigentry.Cdot3Oamerrframesecsevnotifenable}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamDyingGaspEnable"] = types.YLeaf{"Cdot3Oamdyinggaspenable", cdot3Oameventconfigentry.Cdot3Oamdyinggaspenable}
    cdot3Oameventconfigentry.EntityData.Leafs["cdot3OamCriticalEventEnable"] = types.YLeaf{"Cdot3Oamcriticaleventenable", cdot3Oameventconfigentry.Cdot3Oamcriticaleventenable}
    return &(cdot3Oameventconfigentry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oameventlogtable
// This table records a history of the events that have occurred
// at the Ethernet OAM level.  These events can include locally
// detected events, which may result in locally generated
// OAMPDUs, and remotely detected events, which are detected by
// the OAM peer entity and signaled to the local entity via
// Ethernet OAM.  Ethernet OAM events can be signaled by Event
// Notification OAMPDUs or by the flags field in any OAMPDU.  
// 
// This table contains both threshold crossing events and
// non-threshold crossing events.  The parameters for the
// threshold window, threshold value, and actual value
// (cdot3OamEventLogWindowXX, cdot3OamEventLogThresholdXX,
// cdot3OamEventLogValue) are only applicable to threshold
// crossing events, and are returned as all F's (2^32 - 1) for
// non-threshold crossing events.  
// Entries in the table are automatically created when such
// events are detected.  The size of the table is implementation
// dependent.  When the table reaches its maximum size, older
// entries are automatically deleted to make room for newer
// entries. 
type CISCODOT3OAMMIB_Cdot3Oameventlogtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cdot3OamEventLogTable.  Entries are automatically created
    // whenever Ethernet OAM events occur at the local OAM entity, and when Event
    // Notification OAMPDUs are received at the local OAM entity (indicating
    // events have occurred at the peer OAM entity).  The size of the table is
    // implementation dependent, but when the table becomes full, older events are
    // automatically deleted to make room for newer events.  The table index
    // cdot3OamEventLogIndex increments for each new entry, and when the maximum
    // value is reached the value restarts at zero.  . The type is slice of
    // CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry.
    Cdot3Oameventlogentry []CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry
}

func (cdot3Oameventlogtable *CISCODOT3OAMMIB_Cdot3Oameventlogtable) GetEntityData() *types.CommonEntityData {
    cdot3Oameventlogtable.EntityData.YFilter = cdot3Oameventlogtable.YFilter
    cdot3Oameventlogtable.EntityData.YangName = "cdot3OamEventLogTable"
    cdot3Oameventlogtable.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oameventlogtable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3Oameventlogtable.EntityData.SegmentPath = "cdot3OamEventLogTable"
    cdot3Oameventlogtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oameventlogtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oameventlogtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oameventlogtable.EntityData.Children = make(map[string]types.YChild)
    cdot3Oameventlogtable.EntityData.Children["cdot3OamEventLogEntry"] = types.YChild{"Cdot3Oameventlogentry", nil}
    for i := range cdot3Oameventlogtable.Cdot3Oameventlogentry {
        cdot3Oameventlogtable.EntityData.Children[types.GetSegmentPath(&cdot3Oameventlogtable.Cdot3Oameventlogentry[i])] = types.YChild{"Cdot3Oameventlogentry", &cdot3Oameventlogtable.Cdot3Oameventlogentry[i]}
    }
    cdot3Oameventlogtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdot3Oameventlogtable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry
// An entry in the cdot3OamEventLogTable.  Entries are
// automatically created whenever Ethernet OAM events occur at
// the local OAM entity, and when Event Notification OAMPDUs are
// received at the local OAM entity (indicating events have
// occurred at the peer OAM entity).  The size of the table is
// implementation dependent, but when the table becomes full,
// older events are automatically deleted to make room for newer
// events.  The table index cdot3OamEventLogIndex increments for
// each new entry, and when the maximum value is reached the
// value restarts at zero.  
type CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. An arbitrary integer for identifying individual
    // events within the event log.  . The type is interface{} with range:
    // 0..4294967295.
    Cdot3Oameventlogindex interface{}

    // The value of sysUpTime at the time of the logged event.  For locally
    // generated events, the time of the event can be accurately retrieved from
    // sysUpTime.  For remotely generated events, the time of the event is
    // indicated by the reception of the Event Notification OAMPDU indicating the
    // event occurred on the peer.  A system may attempt to adjust the timestamp
    // value to more accurately reflect the time of the event at the peer OAM
    // entity by using other information, such as that found in the timestamp
    // found of the Event Notification TLVs, which provides an indication of the
    // relative time between events at the peer entity.  . The type is interface{}
    // with range: 0..4294967295.
    Cdot3Oameventlogtimestamp interface{}

    // The OUI of the entity defining the object type.  All IEEE 802.3 defined
    // events (as appearing in [802.3ah] except for the Organizationally Unique
    // Event TLVs) use the IEEE 802.3 OUI of 0x0180C2.  Organizations defining
    // their own Event Notification TLVs include their OUI in the Event
    // Notification TLV which gets reflected here.  . The type is string with
    // length: 3.
    Cdot3Oameventlogoui interface{}

    // The type of event that generated this entry in the event log.  When the OUI
    // is the IEEE 802.3 OUI of 0x0180C2, the following event types are defined:  
    // erroredSymbolEvent(1),      erroredFramePeriodEvent (2),     
    // erroredFrameEvent(3),     erroredFrameSecondsEvent(4),      linkFault(256),
    // dyingGaspEvent(257),     criticalLinkEvent(258) The first four are
    // considered threshold crossing events as they are generated when a metric
    // exceeds a given value within a specified window.  The other three are not
    // threshold crossing events.    When the OUI is not 71874 (0x0180C2 in hex),
    // then some other organization has defined the event space.  If event
    // subtyping is known to the implementation, it may be reflected here.
    // Otherwise, this value should return all Fs (2^32 - 1).  . The type is
    // interface{} with range: 0..4294967295.
    Cdot3Oameventlogtype interface{}

    // Whether this event occurred locally (local(1)), or was  received from the
    // OAM peer via Ethernet OAM (remote(2)). The type is
    // Cdot3Oameventloglocation.
    Cdot3Oameventloglocation interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventWindowHi and cdot3OamEventWindowLo form an unsigned 64-bit
    // integer yielding the window over which the value was measured for the
    // threshold crossing event (for example, 5, when 11 occurrences happened in 5
    // seconds while the threshold was 10).   The two objects are combined as: 
    // cdot3OamEventLogWindow = ((2^32) * cdot3OamEventLogWindowHi)               
    // + cdot3OamEventLogWindowLo   Otherwise, this value is returned as all F's
    // (2^32 - 1) and  adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3Oameventlogwindowhi interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventWindowHi and cdot3OamEventWindowLo form an unsigned 64-bit
    // integer yielding the window over which the value was measured for the
    // threshold crossing event (for example, 5, when 11 occurrences happened in 5
    // seconds while the threshold was 10).   The two objects are combined as: 
    // cdot3OamEventLogWindow = ((2^32) * cdot3OamEventLogWindowHi)               
    // + cdot3OamEventLogWindowLo  Otherwise, this value is returned as all F's
    // (2^32 - 1) and  adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3Oameventlogwindowlo interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventThresholdHi and cdot3OamEventThresholdLo form an unsigned
    // 64-bit integer yielding the value that was crossed for the threshold
    // crossing event (for example, 10, when 11 occurrences happened in 5 seconds
    // while the threshold was 10).  The two objects are combined as: 
    // cdot3OamEventLogThreshold = ((2^32) * cdot3OamEventLogThresholdHi)         
    // + cdot3OamEventLogThresholdLo  Otherwise, this value is returned as all F's
    // (2^32 -1) and  adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3Oameventlogthresholdhi interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventThresholdHi and cdot3OamEventThresholdLo form an unsigned
    // 64-bit integer yielding the value that was crossed for the threshold
    // crossing event (for example, 10, when 11 occurrences happened in 5 seconds
    // while the threshold was 10).  The two objects are combined as: 
    // cdot3OamEventLogThreshold = ((2^32) * cdot3OamEventLogThresholdHi)         
    // + cdot3OamEventLogThresholdLo  Otherwise, this value is returned as all F's
    // (2^32 - 1) and adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3Oameventlogthresholdlo interface{}

    // If the event represents a threshold crossing event, this value indicates
    // the value of the parameter within the given window that generated this
    // event (for example, 11, when 11 occurrences happened in 5 seconds while the
    // threshold was 10).    Otherwise, this value is returned as all F's  (2^64 -
    // 1) and adds no useful information.  . The type is interface{} with range:
    // 0..18446744073709551615.
    Cdot3Oameventlogvalue interface{}

    // Each Event Notification TLV contains a running total of the number of times
    // an event has occurred, as well as the number of times an Event Notification
    // for the event has been transmitted.  For non-threshold crossing events, the
    // number of events (cdot3OamLogRunningTotal) and the number of resultant
    // Event Notifications (cdot3OamLogEventTotal) should be identical.   For
    // threshold crossing events, since multiple occurrences may be required to
    // cross the threshold, these values are likely different.  This value
    // represents the total number of times this event has happened since the last
    // reset (for example, 3253, when 3253 symbol errors have occurred since the
    // last reset, which has resulted in 51 symbol error threshold crossing events
    // since the last reset).  . The type is interface{} with range:
    // 0..18446744073709551615.
    Cdot3Oameventlogrunningtotal interface{}

    // Each Event Notification TLV contains a running total of the number of times
    // an event has occurred, as well as the number of times an Event Notification
    // for the event has been transmitted.  For non-threshold crossing events, the
    // number of events (cdot3OamLogRunningTotal) and the number of resultant
    // Event Notifications (cdot3OamLogEventTotal) should be identical.   For
    // threshold crossing events, since multiple occurrences may be required to
    // cross the threshold, these values are likely different.  This value
    // represents the total number of times one or more of these occurrences have
    // resulted in an Event Notification (for example, 51 when 3253 symbol errors
    // have occurred since the last reset, which has resulted in 51 symbol error
    // threshold crossing events since the last reset).  . The type is interface{}
    // with range: 0..4294967295.
    Cdot3Oameventlogeventtotal interface{}
}

func (cdot3Oameventlogentry *CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry) GetEntityData() *types.CommonEntityData {
    cdot3Oameventlogentry.EntityData.YFilter = cdot3Oameventlogentry.YFilter
    cdot3Oameventlogentry.EntityData.YangName = "cdot3OamEventLogEntry"
    cdot3Oameventlogentry.EntityData.BundleName = "cisco_ios_xe"
    cdot3Oameventlogentry.EntityData.ParentYangName = "cdot3OamEventLogTable"
    cdot3Oameventlogentry.EntityData.SegmentPath = "cdot3OamEventLogEntry" + "[ifIndex='" + fmt.Sprintf("%v", cdot3Oameventlogentry.Ifindex) + "']" + "[cdot3OamEventLogIndex='" + fmt.Sprintf("%v", cdot3Oameventlogentry.Cdot3Oameventlogindex) + "']"
    cdot3Oameventlogentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3Oameventlogentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3Oameventlogentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3Oameventlogentry.EntityData.Children = make(map[string]types.YChild)
    cdot3Oameventlogentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdot3Oameventlogentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cdot3Oameventlogentry.Ifindex}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogIndex"] = types.YLeaf{"Cdot3Oameventlogindex", cdot3Oameventlogentry.Cdot3Oameventlogindex}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogTimestamp"] = types.YLeaf{"Cdot3Oameventlogtimestamp", cdot3Oameventlogentry.Cdot3Oameventlogtimestamp}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogOui"] = types.YLeaf{"Cdot3Oameventlogoui", cdot3Oameventlogentry.Cdot3Oameventlogoui}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogType"] = types.YLeaf{"Cdot3Oameventlogtype", cdot3Oameventlogentry.Cdot3Oameventlogtype}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogLocation"] = types.YLeaf{"Cdot3Oameventloglocation", cdot3Oameventlogentry.Cdot3Oameventloglocation}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogWindowHi"] = types.YLeaf{"Cdot3Oameventlogwindowhi", cdot3Oameventlogentry.Cdot3Oameventlogwindowhi}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogWindowLo"] = types.YLeaf{"Cdot3Oameventlogwindowlo", cdot3Oameventlogentry.Cdot3Oameventlogwindowlo}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogThresholdHi"] = types.YLeaf{"Cdot3Oameventlogthresholdhi", cdot3Oameventlogentry.Cdot3Oameventlogthresholdhi}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogThresholdLo"] = types.YLeaf{"Cdot3Oameventlogthresholdlo", cdot3Oameventlogentry.Cdot3Oameventlogthresholdlo}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogValue"] = types.YLeaf{"Cdot3Oameventlogvalue", cdot3Oameventlogentry.Cdot3Oameventlogvalue}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogRunningTotal"] = types.YLeaf{"Cdot3Oameventlogrunningtotal", cdot3Oameventlogentry.Cdot3Oameventlogrunningtotal}
    cdot3Oameventlogentry.EntityData.Leafs["cdot3OamEventLogEventTotal"] = types.YLeaf{"Cdot3Oameventlogeventtotal", cdot3Oameventlogentry.Cdot3Oameventlogeventtotal}
    return &(cdot3Oameventlogentry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry_Cdot3Oameventloglocation represents received from the OAM peer via Ethernet OAM (remote(2)).
type CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry_Cdot3Oameventloglocation string

const (
    CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry_Cdot3Oameventloglocation_local CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry_Cdot3Oameventloglocation = "local"

    CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry_Cdot3Oameventloglocation_remote CISCODOT3OAMMIB_Cdot3Oameventlogtable_Cdot3Oameventlogentry_Cdot3Oameventloglocation = "remote"
)

