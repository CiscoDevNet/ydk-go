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
    Cdot3OamTable CISCODOT3OAMMIB_Cdot3OamTable

    // This table contains information about the OAM peer for a particular
    // Ethernet like interface. OAM entities communicate with a single OAM peer
    // entity on Ethernet links on which OAM is enabled and operating properly. 
    // There is one entry in this table for each entry in the cdot3OamTable for
    // which information on the peer OAM entity is available.  .
    Cdot3OamPeerTable CISCODOT3OAMMIB_Cdot3OamPeerTable

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
    Cdot3OamLoopbackTable CISCODOT3OAMMIB_Cdot3OamLoopbackTable

    // This table contains statistics for the OAM function on a particular
    // Ethernet like interface. There is an entry in the table for every entry in
    // the cdot3OamTable.   The counters in this table are defined as 32-bit
    // entries to match the counter size as defined in [802.3ah].  Given the OAM
    // protocol is a slow protocol, the counters increment at a slow rate. .
    Cdot3OamStatsTable CISCODOT3OAMMIB_Cdot3OamStatsTable

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
    Cdot3OamEventConfigTable CISCODOT3OAMMIB_Cdot3OamEventConfigTable

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
    Cdot3OamEventLogTable CISCODOT3OAMMIB_Cdot3OamEventLogTable
}

func (cISCODOT3OAMMIB *CISCODOT3OAMMIB) GetEntityData() *types.CommonEntityData {
    cISCODOT3OAMMIB.EntityData.YFilter = cISCODOT3OAMMIB.YFilter
    cISCODOT3OAMMIB.EntityData.YangName = "CISCO-DOT3-OAM-MIB"
    cISCODOT3OAMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODOT3OAMMIB.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cISCODOT3OAMMIB.EntityData.SegmentPath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB"
    cISCODOT3OAMMIB.EntityData.AbsolutePath = cISCODOT3OAMMIB.EntityData.SegmentPath
    cISCODOT3OAMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODOT3OAMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODOT3OAMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODOT3OAMMIB.EntityData.Children = types.NewOrderedMap()
    cISCODOT3OAMMIB.EntityData.Children.Append("cdot3OamTable", types.YChild{"Cdot3OamTable", &cISCODOT3OAMMIB.Cdot3OamTable})
    cISCODOT3OAMMIB.EntityData.Children.Append("cdot3OamPeerTable", types.YChild{"Cdot3OamPeerTable", &cISCODOT3OAMMIB.Cdot3OamPeerTable})
    cISCODOT3OAMMIB.EntityData.Children.Append("cdot3OamLoopbackTable", types.YChild{"Cdot3OamLoopbackTable", &cISCODOT3OAMMIB.Cdot3OamLoopbackTable})
    cISCODOT3OAMMIB.EntityData.Children.Append("cdot3OamStatsTable", types.YChild{"Cdot3OamStatsTable", &cISCODOT3OAMMIB.Cdot3OamStatsTable})
    cISCODOT3OAMMIB.EntityData.Children.Append("cdot3OamEventConfigTable", types.YChild{"Cdot3OamEventConfigTable", &cISCODOT3OAMMIB.Cdot3OamEventConfigTable})
    cISCODOT3OAMMIB.EntityData.Children.Append("cdot3OamEventLogTable", types.YChild{"Cdot3OamEventLogTable", &cISCODOT3OAMMIB.Cdot3OamEventLogTable})
    cISCODOT3OAMMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCODOT3OAMMIB.EntityData.YListKeys = []string {}

    return &(cISCODOT3OAMMIB.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamTable
// This table contains the primary controls and status for the
// OAM capabilities of an Ethernet like interface.  There will be
// one row in this table for each Ethernet like interface in the
// system that supports the OAM functions defined in [802.3ah].
type CISCODOT3OAMMIB_Cdot3OamTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information on the Ethernet OAM function
    // for a single Ethernet like interface. Entries in the table are created
    // automatically for each interface supporting Ethernet OAM. The status of the
    // row entry can be determined from cdot3OamOperStatus.    A cdot3OamEntry is
    // indexed in the cdot3OamTable by the ifIndex object of the Interfaces MIB. 
    // . The type is slice of CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry.
    Cdot3OamEntry []*CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry
}

func (cdot3OamTable *CISCODOT3OAMMIB_Cdot3OamTable) GetEntityData() *types.CommonEntityData {
    cdot3OamTable.EntityData.YFilter = cdot3OamTable.YFilter
    cdot3OamTable.EntityData.YangName = "cdot3OamTable"
    cdot3OamTable.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamTable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3OamTable.EntityData.SegmentPath = "cdot3OamTable"
    cdot3OamTable.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/" + cdot3OamTable.EntityData.SegmentPath
    cdot3OamTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamTable.EntityData.Children = types.NewOrderedMap()
    cdot3OamTable.EntityData.Children.Append("cdot3OamEntry", types.YChild{"Cdot3OamEntry", nil})
    for i := range cdot3OamTable.Cdot3OamEntry {
        cdot3OamTable.EntityData.Children.Append(types.GetSegmentPath(cdot3OamTable.Cdot3OamEntry[i]), types.YChild{"Cdot3OamEntry", cdot3OamTable.Cdot3OamEntry[i]})
    }
    cdot3OamTable.EntityData.Leafs = types.NewOrderedMap()

    cdot3OamTable.EntityData.YListKeys = []string {}

    return &(cdot3OamTable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry
// An entry in the table, containing information on the Ethernet
// OAM function for a single Ethernet like interface. Entries in
// the table are created automatically for each interface
// supporting Ethernet OAM. The status of the row entry can be
// determined from cdot3OamOperStatus.  
// 
// A cdot3OamEntry is indexed in the cdot3OamTable by the ifIndex
// object of the Interfaces MIB.  
type CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object is used to provision the default administrative OAM mode for
    // this interface.  This object represents the desired state of OAM for this
    // interface.    The cdot3OamAdminState always starts in the disabled(1) state
    // until an explicit management action or configuration information retained
    // by the system causes a transition to the enabled(2) state.   When
    // enabled(2), Ethernet OAM will attempt to operate over this interface.  .
    // The type is Cdot3OamAdminState.
    Cdot3OamAdminState interface{}

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
    // half-duplex operation.  . The type is Cdot3OamOperStatus.
    Cdot3OamOperStatus interface{}

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
    // operational(9).  . The type is Cdot3OamMode.
    Cdot3OamMode interface{}

    // The largest OAMPDU that the OAM entity supports.  OAM entities exchange
    // maximum OAMPDU sizes and negotiate to use the smaller of the two maximum
    // OAMPDU sizes between the peers. This value is determined by the local
    // implementation.  . The type is interface{} with range: 64..1518. Units are
    // octets.
    Cdot3OamMaxOamPduSize interface{}

    // The configuration revision of the OAM entity as reflected in the latest
    // OAMPDU sent by the OAM entity.  The config revision is used by OAM entities
    // to indicate configuration changes have occurred which might require the
    // peer OAM entity to re-evaluate whether OAM peering is allowed. . The type
    // is interface{} with range: 0..65535.
    Cdot3OamConfigRevision interface{}

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
    Cdot3OamFunctionsSupported interface{}
}

func (cdot3OamEntry *CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry) GetEntityData() *types.CommonEntityData {
    cdot3OamEntry.EntityData.YFilter = cdot3OamEntry.YFilter
    cdot3OamEntry.EntityData.YangName = "cdot3OamEntry"
    cdot3OamEntry.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamEntry.EntityData.ParentYangName = "cdot3OamTable"
    cdot3OamEntry.EntityData.SegmentPath = "cdot3OamEntry" + types.AddKeyToken(cdot3OamEntry.IfIndex, "ifIndex")
    cdot3OamEntry.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/cdot3OamTable/" + cdot3OamEntry.EntityData.SegmentPath
    cdot3OamEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamEntry.EntityData.Children = types.NewOrderedMap()
    cdot3OamEntry.EntityData.Leafs = types.NewOrderedMap()
    cdot3OamEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdot3OamEntry.IfIndex})
    cdot3OamEntry.EntityData.Leafs.Append("cdot3OamAdminState", types.YLeaf{"Cdot3OamAdminState", cdot3OamEntry.Cdot3OamAdminState})
    cdot3OamEntry.EntityData.Leafs.Append("cdot3OamOperStatus", types.YLeaf{"Cdot3OamOperStatus", cdot3OamEntry.Cdot3OamOperStatus})
    cdot3OamEntry.EntityData.Leafs.Append("cdot3OamMode", types.YLeaf{"Cdot3OamMode", cdot3OamEntry.Cdot3OamMode})
    cdot3OamEntry.EntityData.Leafs.Append("cdot3OamMaxOamPduSize", types.YLeaf{"Cdot3OamMaxOamPduSize", cdot3OamEntry.Cdot3OamMaxOamPduSize})
    cdot3OamEntry.EntityData.Leafs.Append("cdot3OamConfigRevision", types.YLeaf{"Cdot3OamConfigRevision", cdot3OamEntry.Cdot3OamConfigRevision})
    cdot3OamEntry.EntityData.Leafs.Append("cdot3OamFunctionsSupported", types.YLeaf{"Cdot3OamFunctionsSupported", cdot3OamEntry.Cdot3OamFunctionsSupported})

    cdot3OamEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdot3OamEntry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamAdminState represents to operate over this interface.  
type CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamAdminState string

const (
    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamAdminState_disabled CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamAdminState = "disabled"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamAdminState_enabled CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamAdminState = "enabled"
)

// CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamMode represents cdot3OamOperStatus was already operational(9).  
type CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamMode string

const (
    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamMode_active CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamMode = "active"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamMode_passive CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamMode = "passive"
)

// CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus represents in half-duplex operation.  
type CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus string

const (
    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_disabled CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "disabled"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_linkFault CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "linkFault"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_passiveWait CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "passiveWait"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_activeSendLocal CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "activeSendLocal"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_sendLocalAndRemote CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "sendLocalAndRemote"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_sendLocalAndRemoteOk CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "sendLocalAndRemoteOk"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_oamPeeringLocallyRejected CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "oamPeeringLocallyRejected"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_oamPeeringRemotelyRejected CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "oamPeeringRemotelyRejected"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_operational CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "operational"

    CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus_nonOperHalfDuplex CISCODOT3OAMMIB_Cdot3OamTable_Cdot3OamEntry_Cdot3OamOperStatus = "nonOperHalfDuplex"
)

// CISCODOT3OAMMIB_Cdot3OamPeerTable
// This table contains information about the OAM peer for a
// particular Ethernet like interface. OAM entities communicate
// with a single OAM peer entity on Ethernet links on which OAM
// is enabled and operating properly.  There is one entry in this
// table for each entry in the cdot3OamTable for which information
// on the peer OAM entity is available.  
type CISCODOT3OAMMIB_Cdot3OamPeerTable struct {
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
    // CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry.
    Cdot3OamPeerEntry []*CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry
}

func (cdot3OamPeerTable *CISCODOT3OAMMIB_Cdot3OamPeerTable) GetEntityData() *types.CommonEntityData {
    cdot3OamPeerTable.EntityData.YFilter = cdot3OamPeerTable.YFilter
    cdot3OamPeerTable.EntityData.YangName = "cdot3OamPeerTable"
    cdot3OamPeerTable.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamPeerTable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3OamPeerTable.EntityData.SegmentPath = "cdot3OamPeerTable"
    cdot3OamPeerTable.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/" + cdot3OamPeerTable.EntityData.SegmentPath
    cdot3OamPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamPeerTable.EntityData.Children = types.NewOrderedMap()
    cdot3OamPeerTable.EntityData.Children.Append("cdot3OamPeerEntry", types.YChild{"Cdot3OamPeerEntry", nil})
    for i := range cdot3OamPeerTable.Cdot3OamPeerEntry {
        cdot3OamPeerTable.EntityData.Children.Append(types.GetSegmentPath(cdot3OamPeerTable.Cdot3OamPeerEntry[i]), types.YChild{"Cdot3OamPeerEntry", cdot3OamPeerTable.Cdot3OamPeerEntry[i]})
    }
    cdot3OamPeerTable.EntityData.Leafs = types.NewOrderedMap()

    cdot3OamPeerTable.EntityData.YListKeys = []string {}

    return &(cdot3OamPeerTable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry
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
type CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The MAC address of the peer OAM entity.  The MAC address is derived from
    // the most recently received OAMPDU. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Cdot3OamPeerMacAddress interface{}

    // The OUI of the OAM peer as reflected in the latest Information OAMPDU
    // received with a Local Information TLV.  The OUI can be used to identify the
    // vendor of the remote OAM entity.  This value is initialized to zero before
    // any Local Information TLV is received.  . The type is string with length:
    // 3..3.
    Cdot3OamPeerVendorOui interface{}

    // The Vendor Info of the OAM peer as reflected in the latest Information
    // OAMPDU received with a Local Information TLV.  The vendor information field
    // is within the Local Information TLV, and can be used to determine
    // additional information about the peer entity.  The format of the vendor
    // information is unspecified within the 32-bit field.  This value is
    // initialized to zero before any Local Information TLV is received.  . The
    // type is interface{} with range: 0..4294967295.
    Cdot3OamPeerVendorInfo interface{}

    // The mode of the OAM peer as reflected in the latest Information OAMPDU
    // received with a Local Information TLV.  The mode of the peer can be
    // determined from the Configuration field in the Local Information TLV of the
    // last Information OAMPDU received from the peer.  The value is unknown(3)
    // whenever no Local Information TLV has been received.  The values of
    // active(1) and passive(2) are returned when a Local Information TLV has been
    // received indicating the peer is in active or passive mode, respectively. .
    // The type is Cdot3OamPeerMode.
    Cdot3OamPeerMode interface{}

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
    Cdot3OamPeerMaxOamPduSize interface{}

    // The configuration revision of the OAM peer as reflected in the latest
    // OAMPDU.  This attribute is changed by the peer whenever it has a local
    // configuration change for Ethernet OAM this interface.  The configuration
    // revision can be determined from the Revision field of the Local Information
    // TLV of the most recently received Information OAMPDU with a Local
    // Information TLV. A value of zero is returned if no Local Information TLV
    // has been received.  . The type is interface{} with range: 0..65535.
    Cdot3OamPeerConfigRevision interface{}

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
    Cdot3OamPeerFunctionsSupported interface{}
}

func (cdot3OamPeerEntry *CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry) GetEntityData() *types.CommonEntityData {
    cdot3OamPeerEntry.EntityData.YFilter = cdot3OamPeerEntry.YFilter
    cdot3OamPeerEntry.EntityData.YangName = "cdot3OamPeerEntry"
    cdot3OamPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamPeerEntry.EntityData.ParentYangName = "cdot3OamPeerTable"
    cdot3OamPeerEntry.EntityData.SegmentPath = "cdot3OamPeerEntry" + types.AddKeyToken(cdot3OamPeerEntry.IfIndex, "ifIndex")
    cdot3OamPeerEntry.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/cdot3OamPeerTable/" + cdot3OamPeerEntry.EntityData.SegmentPath
    cdot3OamPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamPeerEntry.EntityData.Children = types.NewOrderedMap()
    cdot3OamPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    cdot3OamPeerEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdot3OamPeerEntry.IfIndex})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerMacAddress", types.YLeaf{"Cdot3OamPeerMacAddress", cdot3OamPeerEntry.Cdot3OamPeerMacAddress})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerVendorOui", types.YLeaf{"Cdot3OamPeerVendorOui", cdot3OamPeerEntry.Cdot3OamPeerVendorOui})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerVendorInfo", types.YLeaf{"Cdot3OamPeerVendorInfo", cdot3OamPeerEntry.Cdot3OamPeerVendorInfo})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerMode", types.YLeaf{"Cdot3OamPeerMode", cdot3OamPeerEntry.Cdot3OamPeerMode})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerMaxOamPduSize", types.YLeaf{"Cdot3OamPeerMaxOamPduSize", cdot3OamPeerEntry.Cdot3OamPeerMaxOamPduSize})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerConfigRevision", types.YLeaf{"Cdot3OamPeerConfigRevision", cdot3OamPeerEntry.Cdot3OamPeerConfigRevision})
    cdot3OamPeerEntry.EntityData.Leafs.Append("cdot3OamPeerFunctionsSupported", types.YLeaf{"Cdot3OamPeerFunctionsSupported", cdot3OamPeerEntry.Cdot3OamPeerFunctionsSupported})

    cdot3OamPeerEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdot3OamPeerEntry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode represents active or passive mode, respectively. 
type CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode string

const (
    CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode_active CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode = "active"

    CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode_passive CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode = "passive"

    CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode_unknown CISCODOT3OAMMIB_Cdot3OamPeerTable_Cdot3OamPeerEntry_Cdot3OamPeerMode = "unknown"
)

// CISCODOT3OAMMIB_Cdot3OamLoopbackTable
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
type CISCODOT3OAMMIB_Cdot3OamLoopbackTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information on the loopback status for a
    // single Ethernet like interface.  Entries in the table are automatically
    // created whenever the local OAM entity supports loopback capabilities.  The
    // loopback status on the interface can be determined from the
    // cdot3OamLoopbackStatus object.  . The type is slice of
    // CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry.
    Cdot3OamLoopbackEntry []*CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry
}

func (cdot3OamLoopbackTable *CISCODOT3OAMMIB_Cdot3OamLoopbackTable) GetEntityData() *types.CommonEntityData {
    cdot3OamLoopbackTable.EntityData.YFilter = cdot3OamLoopbackTable.YFilter
    cdot3OamLoopbackTable.EntityData.YangName = "cdot3OamLoopbackTable"
    cdot3OamLoopbackTable.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamLoopbackTable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3OamLoopbackTable.EntityData.SegmentPath = "cdot3OamLoopbackTable"
    cdot3OamLoopbackTable.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/" + cdot3OamLoopbackTable.EntityData.SegmentPath
    cdot3OamLoopbackTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamLoopbackTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamLoopbackTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamLoopbackTable.EntityData.Children = types.NewOrderedMap()
    cdot3OamLoopbackTable.EntityData.Children.Append("cdot3OamLoopbackEntry", types.YChild{"Cdot3OamLoopbackEntry", nil})
    for i := range cdot3OamLoopbackTable.Cdot3OamLoopbackEntry {
        cdot3OamLoopbackTable.EntityData.Children.Append(types.GetSegmentPath(cdot3OamLoopbackTable.Cdot3OamLoopbackEntry[i]), types.YChild{"Cdot3OamLoopbackEntry", cdot3OamLoopbackTable.Cdot3OamLoopbackEntry[i]})
    }
    cdot3OamLoopbackTable.EntityData.Leafs = types.NewOrderedMap()

    cdot3OamLoopbackTable.EntityData.YListKeys = []string {}

    return &(cdot3OamLoopbackTable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry
// An entry in the table, containing information on the loopback
// status for a single Ethernet like interface.  Entries in the
// table are automatically created whenever the local OAM entity
// supports loopback capabilities.  The loopback status on the
// interface can be determined from the cdot3OamLoopbackStatus
// object.  
type CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

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
    // type is Cdot3OamLoopbackStatus.
    Cdot3OamLoopbackStatus interface{}

    // Since OAM loopback is a disruptive operation (user traffic does not pass),
    // this attribute provides a mechanism to provide controls over whether
    // received OAM loopback commands are processed or ignored.  When the value is
    // ignore(1), received loopback commands are ignored.  When the value is
    // process(2), OAM loopback commands are processed.  The default value is to
    // ignore loopback commands (ignore(1)).  . The type is
    // Cdot3OamLoopbackIgnoreRx.
    Cdot3OamLoopbackIgnoreRx interface{}
}

func (cdot3OamLoopbackEntry *CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry) GetEntityData() *types.CommonEntityData {
    cdot3OamLoopbackEntry.EntityData.YFilter = cdot3OamLoopbackEntry.YFilter
    cdot3OamLoopbackEntry.EntityData.YangName = "cdot3OamLoopbackEntry"
    cdot3OamLoopbackEntry.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamLoopbackEntry.EntityData.ParentYangName = "cdot3OamLoopbackTable"
    cdot3OamLoopbackEntry.EntityData.SegmentPath = "cdot3OamLoopbackEntry" + types.AddKeyToken(cdot3OamLoopbackEntry.IfIndex, "ifIndex")
    cdot3OamLoopbackEntry.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/cdot3OamLoopbackTable/" + cdot3OamLoopbackEntry.EntityData.SegmentPath
    cdot3OamLoopbackEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamLoopbackEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamLoopbackEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamLoopbackEntry.EntityData.Children = types.NewOrderedMap()
    cdot3OamLoopbackEntry.EntityData.Leafs = types.NewOrderedMap()
    cdot3OamLoopbackEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdot3OamLoopbackEntry.IfIndex})
    cdot3OamLoopbackEntry.EntityData.Leafs.Append("cdot3OamLoopbackStatus", types.YLeaf{"Cdot3OamLoopbackStatus", cdot3OamLoopbackEntry.Cdot3OamLoopbackStatus})
    cdot3OamLoopbackEntry.EntityData.Leafs.Append("cdot3OamLoopbackIgnoreRx", types.YLeaf{"Cdot3OamLoopbackIgnoreRx", cdot3OamLoopbackEntry.Cdot3OamLoopbackIgnoreRx})

    cdot3OamLoopbackEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdot3OamLoopbackEntry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackIgnoreRx represents ignore loopback commands (ignore(1)).  
type CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackIgnoreRx string

const (
    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackIgnoreRx_ignore CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackIgnoreRx = "ignore"

    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackIgnoreRx_process CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackIgnoreRx = "process"
)

// CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus represents   unknown            ***   any other combination   ***
type CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus string

const (
    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus_noLoopback CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus = "noLoopback"

    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus_initiatingLoopback CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus = "initiatingLoopback"

    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus_remoteLoopback CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus = "remoteLoopback"

    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus_terminatingLoopback CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus = "terminatingLoopback"

    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus_localLoopback CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus = "localLoopback"

    CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus_unknown CISCODOT3OAMMIB_Cdot3OamLoopbackTable_Cdot3OamLoopbackEntry_Cdot3OamLoopbackStatus = "unknown"
)

// CISCODOT3OAMMIB_Cdot3OamStatsTable
// This table contains statistics for the OAM function on a
// particular Ethernet like interface. There is an entry in the
// table for every entry in the cdot3OamTable. 
// 
// The counters in this table are defined as 32-bit entries to
// match the counter size as defined in [802.3ah].  Given the OAM
// protocol is a slow protocol, the counters increment at a slow
// rate. 
type CISCODOT3OAMMIB_Cdot3OamStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing statistics information on the Ethernet
    // OAM function for a single Ethernet like interface.  Entries are
    // automatically created for every entry in the cdot3OamTable.  Counters are
    // maintained across transitions in cdot3OamOperStatus.  . The type is slice
    // of CISCODOT3OAMMIB_Cdot3OamStatsTable_Cdot3OamStatsEntry.
    Cdot3OamStatsEntry []*CISCODOT3OAMMIB_Cdot3OamStatsTable_Cdot3OamStatsEntry
}

func (cdot3OamStatsTable *CISCODOT3OAMMIB_Cdot3OamStatsTable) GetEntityData() *types.CommonEntityData {
    cdot3OamStatsTable.EntityData.YFilter = cdot3OamStatsTable.YFilter
    cdot3OamStatsTable.EntityData.YangName = "cdot3OamStatsTable"
    cdot3OamStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamStatsTable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3OamStatsTable.EntityData.SegmentPath = "cdot3OamStatsTable"
    cdot3OamStatsTable.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/" + cdot3OamStatsTable.EntityData.SegmentPath
    cdot3OamStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamStatsTable.EntityData.Children = types.NewOrderedMap()
    cdot3OamStatsTable.EntityData.Children.Append("cdot3OamStatsEntry", types.YChild{"Cdot3OamStatsEntry", nil})
    for i := range cdot3OamStatsTable.Cdot3OamStatsEntry {
        cdot3OamStatsTable.EntityData.Children.Append(types.GetSegmentPath(cdot3OamStatsTable.Cdot3OamStatsEntry[i]), types.YChild{"Cdot3OamStatsEntry", cdot3OamStatsTable.Cdot3OamStatsEntry[i]})
    }
    cdot3OamStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cdot3OamStatsTable.EntityData.YListKeys = []string {}

    return &(cdot3OamStatsTable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamStatsTable_Cdot3OamStatsEntry
// An entry in the table, containing statistics information on
// the Ethernet OAM function for a single Ethernet like
// interface.  Entries are automatically created for every entry
// in the cdot3OamTable.  Counters are maintained across
// transitions in cdot3OamOperStatus.  
type CISCODOT3OAMMIB_Cdot3OamStatsTable_Cdot3OamStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // A count of the number of Information OAMPDUs transmitted on this interface.
    // Discontinuities of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of the
    // ifCounterDiscontinuityTime.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3OamInformationTx interface{}

    // A count of the number of Information OAMPDUs received on this interface. 
    // Discontinuities of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of the
    // ifCounterDiscontinuityTime.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3OamInformationRx interface{}

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
    Cdot3OamUniqueEventNotificationTx interface{}

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
    Cdot3OamUniqueEventNotificationRx interface{}

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
    Cdot3OamDuplicateEventNotificationTx interface{}

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
    Cdot3OamDuplicateEventNotificationRx interface{}

    // A count of the number of Loopback Control OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamLoopbackControlTx interface{}

    // A count of the number of Loopback Control OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamLoopbackControlRx interface{}

    // A count of the number of Variable Request OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamVariableRequestTx interface{}

    // A count of the number of Variable Request OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamVariableRequestRx interface{}

    // A count of the number of Variable Response OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamVariableResponseTx interface{}

    // A count of the number of Variable Response OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamVariableResponseRx interface{}

    // A count of the number of Organization Specific OAMPDUs transmitted on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamOrgSpecificTx interface{}

    // A count of the number of Organization Specific OAMPDUs received on this
    // interface.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamOrgSpecificRx interface{}

    // A count of the number of OAMPDUs transmitted on this interface with an
    // unsupported op-code.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamUnsupportedCodesTx interface{}

    // A count of the number of OAMPDUs received on this interface with an
    // unsupported op-code.    Discontinuities of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of the ifCounterDiscontinuityTime.  . The type is interface{}
    // with range: 0..4294967295. Units are frames.
    Cdot3OamUnsupportedCodesRx interface{}

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
    Cdot3OamFramesLostDueToOam interface{}
}

func (cdot3OamStatsEntry *CISCODOT3OAMMIB_Cdot3OamStatsTable_Cdot3OamStatsEntry) GetEntityData() *types.CommonEntityData {
    cdot3OamStatsEntry.EntityData.YFilter = cdot3OamStatsEntry.YFilter
    cdot3OamStatsEntry.EntityData.YangName = "cdot3OamStatsEntry"
    cdot3OamStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamStatsEntry.EntityData.ParentYangName = "cdot3OamStatsTable"
    cdot3OamStatsEntry.EntityData.SegmentPath = "cdot3OamStatsEntry" + types.AddKeyToken(cdot3OamStatsEntry.IfIndex, "ifIndex")
    cdot3OamStatsEntry.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/cdot3OamStatsTable/" + cdot3OamStatsEntry.EntityData.SegmentPath
    cdot3OamStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamStatsEntry.EntityData.Children = types.NewOrderedMap()
    cdot3OamStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cdot3OamStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdot3OamStatsEntry.IfIndex})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamInformationTx", types.YLeaf{"Cdot3OamInformationTx", cdot3OamStatsEntry.Cdot3OamInformationTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamInformationRx", types.YLeaf{"Cdot3OamInformationRx", cdot3OamStatsEntry.Cdot3OamInformationRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamUniqueEventNotificationTx", types.YLeaf{"Cdot3OamUniqueEventNotificationTx", cdot3OamStatsEntry.Cdot3OamUniqueEventNotificationTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamUniqueEventNotificationRx", types.YLeaf{"Cdot3OamUniqueEventNotificationRx", cdot3OamStatsEntry.Cdot3OamUniqueEventNotificationRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamDuplicateEventNotificationTx", types.YLeaf{"Cdot3OamDuplicateEventNotificationTx", cdot3OamStatsEntry.Cdot3OamDuplicateEventNotificationTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamDuplicateEventNotificationRx", types.YLeaf{"Cdot3OamDuplicateEventNotificationRx", cdot3OamStatsEntry.Cdot3OamDuplicateEventNotificationRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamLoopbackControlTx", types.YLeaf{"Cdot3OamLoopbackControlTx", cdot3OamStatsEntry.Cdot3OamLoopbackControlTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamLoopbackControlRx", types.YLeaf{"Cdot3OamLoopbackControlRx", cdot3OamStatsEntry.Cdot3OamLoopbackControlRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamVariableRequestTx", types.YLeaf{"Cdot3OamVariableRequestTx", cdot3OamStatsEntry.Cdot3OamVariableRequestTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamVariableRequestRx", types.YLeaf{"Cdot3OamVariableRequestRx", cdot3OamStatsEntry.Cdot3OamVariableRequestRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamVariableResponseTx", types.YLeaf{"Cdot3OamVariableResponseTx", cdot3OamStatsEntry.Cdot3OamVariableResponseTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamVariableResponseRx", types.YLeaf{"Cdot3OamVariableResponseRx", cdot3OamStatsEntry.Cdot3OamVariableResponseRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamOrgSpecificTx", types.YLeaf{"Cdot3OamOrgSpecificTx", cdot3OamStatsEntry.Cdot3OamOrgSpecificTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamOrgSpecificRx", types.YLeaf{"Cdot3OamOrgSpecificRx", cdot3OamStatsEntry.Cdot3OamOrgSpecificRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamUnsupportedCodesTx", types.YLeaf{"Cdot3OamUnsupportedCodesTx", cdot3OamStatsEntry.Cdot3OamUnsupportedCodesTx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamUnsupportedCodesRx", types.YLeaf{"Cdot3OamUnsupportedCodesRx", cdot3OamStatsEntry.Cdot3OamUnsupportedCodesRx})
    cdot3OamStatsEntry.EntityData.Leafs.Append("cdot3OamFramesLostDueToOam", types.YLeaf{"Cdot3OamFramesLostDueToOam", cdot3OamStatsEntry.Cdot3OamFramesLostDueToOam})

    cdot3OamStatsEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdot3OamStatsEntry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamEventConfigTable
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
type CISCODOT3OAMMIB_Cdot3OamEventConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entries are automatically created and deleted from this table, and exist
    // whenever the OAM entity supports Ethernet OAM events (as indicated by the
    // eventSupport bit in cdot3OamFunctionsSuppported).  Values in the table are
    // maintained across changes to the value of cdot3OamOperStatus.  Event
    // configuration controls when the local management entity sends Event
    // Notification OAMPDUs to its OAM peer, and when certain event flags are set
    // or cleared in OAMPDUs. . The type is slice of
    // CISCODOT3OAMMIB_Cdot3OamEventConfigTable_Cdot3OamEventConfigEntry.
    Cdot3OamEventConfigEntry []*CISCODOT3OAMMIB_Cdot3OamEventConfigTable_Cdot3OamEventConfigEntry
}

func (cdot3OamEventConfigTable *CISCODOT3OAMMIB_Cdot3OamEventConfigTable) GetEntityData() *types.CommonEntityData {
    cdot3OamEventConfigTable.EntityData.YFilter = cdot3OamEventConfigTable.YFilter
    cdot3OamEventConfigTable.EntityData.YangName = "cdot3OamEventConfigTable"
    cdot3OamEventConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamEventConfigTable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3OamEventConfigTable.EntityData.SegmentPath = "cdot3OamEventConfigTable"
    cdot3OamEventConfigTable.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/" + cdot3OamEventConfigTable.EntityData.SegmentPath
    cdot3OamEventConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamEventConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamEventConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamEventConfigTable.EntityData.Children = types.NewOrderedMap()
    cdot3OamEventConfigTable.EntityData.Children.Append("cdot3OamEventConfigEntry", types.YChild{"Cdot3OamEventConfigEntry", nil})
    for i := range cdot3OamEventConfigTable.Cdot3OamEventConfigEntry {
        cdot3OamEventConfigTable.EntityData.Children.Append(types.GetSegmentPath(cdot3OamEventConfigTable.Cdot3OamEventConfigEntry[i]), types.YChild{"Cdot3OamEventConfigEntry", cdot3OamEventConfigTable.Cdot3OamEventConfigEntry[i]})
    }
    cdot3OamEventConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cdot3OamEventConfigTable.EntityData.YListKeys = []string {}

    return &(cdot3OamEventConfigTable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamEventConfigTable_Cdot3OamEventConfigEntry
// Entries are automatically created and deleted from this
// table, and exist whenever the OAM entity supports Ethernet OAM
// events (as indicated by the eventSupport bit in
// cdot3OamFunctionsSuppported).  Values in the table are
// maintained across changes to the value of cdot3OamOperStatus.
// 
// Event configuration controls when the local management entity
// sends Event Notification OAMPDUs to its OAM peer, and when
// certain event flags are set or cleared in OAMPDUs. 
type CISCODOT3OAMMIB_Cdot3OamEventConfigTable_Cdot3OamEventConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

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
    Cdot3OamErrSymPeriodWindowHi interface{}

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
    Cdot3OamErrSymPeriodWindowLo interface{}

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
    Cdot3OamErrSymPeriodThresholdHi interface{}

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
    Cdot3OamErrSymPeriodThresholdLo interface{}

    // If true, the OAM entity should send an Event Notification OAMPDU when an
    // Errored Symbol Period Event occurs.  By default, this object should have
    // the value true for Ethernet like interfaces that support OAM.  If the OAM
    // layer does not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. The type is
    // bool.
    Cdot3OamErrSymPeriodEvNotifEnable interface{}

    // The number of frames over which the threshold is defined. The default value
    // of the window is the number of minimum size Ethernet frames that can be
    // received over the physical layer in one second.                    If
    // cdot3OamErrFramePeriodThreshold frame errors occur within a window of
    // cdot3OamErrFramePeriodWindow frames, an Event Notification OAMPDU should be
    // generated with an Errored Frame Period Event TLV indicating the threshold
    // has been crossed in this window.  . The type is interface{} with range:
    // 0..4294967295. Units are frames.
    Cdot3OamErrFramePeriodWindow interface{}

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
    Cdot3OamErrFramePeriodThreshold interface{}

    // If true, the OAM entity should send an Event Notification OAMPDU when an
    // Errored Frame Period Event occurs.   By default, this object should have
    // the value true for Ethernet like interfaces that support OAM.  If the OAM
    // layer does not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. . The type is
    // bool.
    Cdot3OamErrFramePeriodEvNotifEnable interface{}

    // The amount of time (in 100ms increments) over which the threshold is
    // defined.  The default value is 10 (1 second).                    If
    // cdot3OamErrFrameThreshold frame errors occur within a window of
    // cdot3OamErrFrameWindow seconds (measured in tenths of seconds), an Event
    // Notification OAMPDU should be generated with an Errored Frame Event TLV
    // indicating the threshold has been crossed in this window.  . The type is
    // interface{} with range: 0..4294967295. Units are tenths of a second.
    Cdot3OamErrFrameWindow interface{}

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
    Cdot3OamErrFrameThreshold interface{}

    // If true, the OAM entity should send an Event Notification OAMPDU when an
    // Errored Frame Event occurs.                   By default, this object
    // should have the value true for Ethernet like interfaces that support OAM. 
    // If the OAM layer does not support event notifications (as indicated via the
    // cdot3OamFunctionsSupported attribute), this value is ignored. . The type is
    // bool.
    Cdot3OamErrFrameEvNotifEnable interface{}

    // The amount of time (in 100ms intervals) over which the threshold is
    // defined.  The default value is 100 (10 seconds).                    If
    // cdot3OamErrFrameSecsSummaryThreshold frame errors occur within a window of
    // cdot3OamErrFrameSecsSummaryWindow (in tenths of seconds), an Event
    // Notification OAMPDU should be generated with an Errored Frame Seconds
    // Summary Event TLV indicating the threshold has been crossed in this window.
    // . The type is interface{} with range: 100..9000. Units are tenths of a
    // second.
    Cdot3OamErrFrameSecsSummaryWindow interface{}

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
    Cdot3OamErrFrameSecsSummaryThreshold interface{}

    // If true, the local OAM entity should send an Event Notification OAMPDU when
    // an Errored Frame Seconds Event occurs.                   By default, this
    // object should have the value true for Ethernet like interfaces that support
    // OAM.  If the OAM layer does not support event notifications (as indicated
    // via the cdot3OamFunctionsSupported attribute), this value is ignored. The
    // type is bool.
    Cdot3OamErrFrameSecsEvNotifEnable interface{}

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
    Cdot3OamDyingGaspEnable interface{}

    // If true, the local OAM entity should attempt to indicate a critical event
    // via the OAMPDU flags to its peer OAM entity when a critical event occurs.  
    // The exact definition of a critical event is implementation dependent.  If
    // the system does not support critical event capability, setting this object
    // has no effect, and reading the object should always result in 'false'.     
    // By default, this object should have the value true for Ethernet like
    // interfaces that support OAM.  If the OAM layer does not support event
    // notifications (as indicated via the cdot3OamFunctionsSupported attribute),
    // this value is ignored. The type is bool.
    Cdot3OamCriticalEventEnable interface{}
}

func (cdot3OamEventConfigEntry *CISCODOT3OAMMIB_Cdot3OamEventConfigTable_Cdot3OamEventConfigEntry) GetEntityData() *types.CommonEntityData {
    cdot3OamEventConfigEntry.EntityData.YFilter = cdot3OamEventConfigEntry.YFilter
    cdot3OamEventConfigEntry.EntityData.YangName = "cdot3OamEventConfigEntry"
    cdot3OamEventConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamEventConfigEntry.EntityData.ParentYangName = "cdot3OamEventConfigTable"
    cdot3OamEventConfigEntry.EntityData.SegmentPath = "cdot3OamEventConfigEntry" + types.AddKeyToken(cdot3OamEventConfigEntry.IfIndex, "ifIndex")
    cdot3OamEventConfigEntry.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/cdot3OamEventConfigTable/" + cdot3OamEventConfigEntry.EntityData.SegmentPath
    cdot3OamEventConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamEventConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamEventConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamEventConfigEntry.EntityData.Children = types.NewOrderedMap()
    cdot3OamEventConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdot3OamEventConfigEntry.IfIndex})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrSymPeriodWindowHi", types.YLeaf{"Cdot3OamErrSymPeriodWindowHi", cdot3OamEventConfigEntry.Cdot3OamErrSymPeriodWindowHi})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrSymPeriodWindowLo", types.YLeaf{"Cdot3OamErrSymPeriodWindowLo", cdot3OamEventConfigEntry.Cdot3OamErrSymPeriodWindowLo})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrSymPeriodThresholdHi", types.YLeaf{"Cdot3OamErrSymPeriodThresholdHi", cdot3OamEventConfigEntry.Cdot3OamErrSymPeriodThresholdHi})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrSymPeriodThresholdLo", types.YLeaf{"Cdot3OamErrSymPeriodThresholdLo", cdot3OamEventConfigEntry.Cdot3OamErrSymPeriodThresholdLo})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrSymPeriodEvNotifEnable", types.YLeaf{"Cdot3OamErrSymPeriodEvNotifEnable", cdot3OamEventConfigEntry.Cdot3OamErrSymPeriodEvNotifEnable})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFramePeriodWindow", types.YLeaf{"Cdot3OamErrFramePeriodWindow", cdot3OamEventConfigEntry.Cdot3OamErrFramePeriodWindow})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFramePeriodThreshold", types.YLeaf{"Cdot3OamErrFramePeriodThreshold", cdot3OamEventConfigEntry.Cdot3OamErrFramePeriodThreshold})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFramePeriodEvNotifEnable", types.YLeaf{"Cdot3OamErrFramePeriodEvNotifEnable", cdot3OamEventConfigEntry.Cdot3OamErrFramePeriodEvNotifEnable})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFrameWindow", types.YLeaf{"Cdot3OamErrFrameWindow", cdot3OamEventConfigEntry.Cdot3OamErrFrameWindow})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFrameThreshold", types.YLeaf{"Cdot3OamErrFrameThreshold", cdot3OamEventConfigEntry.Cdot3OamErrFrameThreshold})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFrameEvNotifEnable", types.YLeaf{"Cdot3OamErrFrameEvNotifEnable", cdot3OamEventConfigEntry.Cdot3OamErrFrameEvNotifEnable})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFrameSecsSummaryWindow", types.YLeaf{"Cdot3OamErrFrameSecsSummaryWindow", cdot3OamEventConfigEntry.Cdot3OamErrFrameSecsSummaryWindow})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFrameSecsSummaryThreshold", types.YLeaf{"Cdot3OamErrFrameSecsSummaryThreshold", cdot3OamEventConfigEntry.Cdot3OamErrFrameSecsSummaryThreshold})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamErrFrameSecsEvNotifEnable", types.YLeaf{"Cdot3OamErrFrameSecsEvNotifEnable", cdot3OamEventConfigEntry.Cdot3OamErrFrameSecsEvNotifEnable})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamDyingGaspEnable", types.YLeaf{"Cdot3OamDyingGaspEnable", cdot3OamEventConfigEntry.Cdot3OamDyingGaspEnable})
    cdot3OamEventConfigEntry.EntityData.Leafs.Append("cdot3OamCriticalEventEnable", types.YLeaf{"Cdot3OamCriticalEventEnable", cdot3OamEventConfigEntry.Cdot3OamCriticalEventEnable})

    cdot3OamEventConfigEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cdot3OamEventConfigEntry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamEventLogTable
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
type CISCODOT3OAMMIB_Cdot3OamEventLogTable struct {
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
    // CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry.
    Cdot3OamEventLogEntry []*CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry
}

func (cdot3OamEventLogTable *CISCODOT3OAMMIB_Cdot3OamEventLogTable) GetEntityData() *types.CommonEntityData {
    cdot3OamEventLogTable.EntityData.YFilter = cdot3OamEventLogTable.YFilter
    cdot3OamEventLogTable.EntityData.YangName = "cdot3OamEventLogTable"
    cdot3OamEventLogTable.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamEventLogTable.EntityData.ParentYangName = "CISCO-DOT3-OAM-MIB"
    cdot3OamEventLogTable.EntityData.SegmentPath = "cdot3OamEventLogTable"
    cdot3OamEventLogTable.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/" + cdot3OamEventLogTable.EntityData.SegmentPath
    cdot3OamEventLogTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamEventLogTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamEventLogTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamEventLogTable.EntityData.Children = types.NewOrderedMap()
    cdot3OamEventLogTable.EntityData.Children.Append("cdot3OamEventLogEntry", types.YChild{"Cdot3OamEventLogEntry", nil})
    for i := range cdot3OamEventLogTable.Cdot3OamEventLogEntry {
        cdot3OamEventLogTable.EntityData.Children.Append(types.GetSegmentPath(cdot3OamEventLogTable.Cdot3OamEventLogEntry[i]), types.YChild{"Cdot3OamEventLogEntry", cdot3OamEventLogTable.Cdot3OamEventLogEntry[i]})
    }
    cdot3OamEventLogTable.EntityData.Leafs = types.NewOrderedMap()

    cdot3OamEventLogTable.EntityData.YListKeys = []string {}

    return &(cdot3OamEventLogTable.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry
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
type CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. An arbitrary integer for identifying individual
    // events within the event log.  . The type is interface{} with range:
    // 0..4294967295.
    Cdot3OamEventLogIndex interface{}

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
    Cdot3OamEventLogTimestamp interface{}

    // The OUI of the entity defining the object type.  All IEEE 802.3 defined
    // events (as appearing in [802.3ah] except for the Organizationally Unique
    // Event TLVs) use the IEEE 802.3 OUI of 0x0180C2.  Organizations defining
    // their own Event Notification TLVs include their OUI in the Event
    // Notification TLV which gets reflected here.  . The type is string with
    // length: 3..3.
    Cdot3OamEventLogOui interface{}

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
    Cdot3OamEventLogType interface{}

    // Whether this event occurred locally (local(1)), or was  received from the
    // OAM peer via Ethernet OAM (remote(2)). The type is
    // Cdot3OamEventLogLocation.
    Cdot3OamEventLogLocation interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventWindowHi and cdot3OamEventWindowLo form an unsigned 64-bit
    // integer yielding the window over which the value was measured for the
    // threshold crossing event (for example, 5, when 11 occurrences happened in 5
    // seconds while the threshold was 10).   The two objects are combined as: 
    // cdot3OamEventLogWindow = ((2^32) * cdot3OamEventLogWindowHi)               
    // + cdot3OamEventLogWindowLo   Otherwise, this value is returned as all F's
    // (2^32 - 1) and  adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3OamEventLogWindowHi interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventWindowHi and cdot3OamEventWindowLo form an unsigned 64-bit
    // integer yielding the window over which the value was measured for the
    // threshold crossing event (for example, 5, when 11 occurrences happened in 5
    // seconds while the threshold was 10).   The two objects are combined as: 
    // cdot3OamEventLogWindow = ((2^32) * cdot3OamEventLogWindowHi)               
    // + cdot3OamEventLogWindowLo  Otherwise, this value is returned as all F's
    // (2^32 - 1) and  adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3OamEventLogWindowLo interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventThresholdHi and cdot3OamEventThresholdLo form an unsigned
    // 64-bit integer yielding the value that was crossed for the threshold
    // crossing event (for example, 10, when 11 occurrences happened in 5 seconds
    // while the threshold was 10).  The two objects are combined as: 
    // cdot3OamEventLogThreshold = ((2^32) * cdot3OamEventLogThresholdHi)         
    // + cdot3OamEventLogThresholdLo  Otherwise, this value is returned as all F's
    // (2^32 -1) and  adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3OamEventLogThresholdHi interface{}

    // If the event represents a threshold crossing event, the two objects
    // cdot3OamEventThresholdHi and cdot3OamEventThresholdLo form an unsigned
    // 64-bit integer yielding the value that was crossed for the threshold
    // crossing event (for example, 10, when 11 occurrences happened in 5 seconds
    // while the threshold was 10).  The two objects are combined as: 
    // cdot3OamEventLogThreshold = ((2^32) * cdot3OamEventLogThresholdHi)         
    // + cdot3OamEventLogThresholdLo  Otherwise, this value is returned as all F's
    // (2^32 - 1) and adds no useful information.  . The type is interface{} with
    // range: 0..4294967295.
    Cdot3OamEventLogThresholdLo interface{}

    // If the event represents a threshold crossing event, this value indicates
    // the value of the parameter within the given window that generated this
    // event (for example, 11, when 11 occurrences happened in 5 seconds while the
    // threshold was 10).    Otherwise, this value is returned as all F's  (2^64 -
    // 1) and adds no useful information.  . The type is interface{} with range:
    // 0..18446744073709551615.
    Cdot3OamEventLogValue interface{}

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
    Cdot3OamEventLogRunningTotal interface{}

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
    Cdot3OamEventLogEventTotal interface{}
}

func (cdot3OamEventLogEntry *CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry) GetEntityData() *types.CommonEntityData {
    cdot3OamEventLogEntry.EntityData.YFilter = cdot3OamEventLogEntry.YFilter
    cdot3OamEventLogEntry.EntityData.YangName = "cdot3OamEventLogEntry"
    cdot3OamEventLogEntry.EntityData.BundleName = "cisco_ios_xe"
    cdot3OamEventLogEntry.EntityData.ParentYangName = "cdot3OamEventLogTable"
    cdot3OamEventLogEntry.EntityData.SegmentPath = "cdot3OamEventLogEntry" + types.AddKeyToken(cdot3OamEventLogEntry.IfIndex, "ifIndex") + types.AddKeyToken(cdot3OamEventLogEntry.Cdot3OamEventLogIndex, "cdot3OamEventLogIndex")
    cdot3OamEventLogEntry.EntityData.AbsolutePath = "CISCO-DOT3-OAM-MIB:CISCO-DOT3-OAM-MIB/cdot3OamEventLogTable/" + cdot3OamEventLogEntry.EntityData.SegmentPath
    cdot3OamEventLogEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdot3OamEventLogEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdot3OamEventLogEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdot3OamEventLogEntry.EntityData.Children = types.NewOrderedMap()
    cdot3OamEventLogEntry.EntityData.Leafs = types.NewOrderedMap()
    cdot3OamEventLogEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cdot3OamEventLogEntry.IfIndex})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogIndex", types.YLeaf{"Cdot3OamEventLogIndex", cdot3OamEventLogEntry.Cdot3OamEventLogIndex})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogTimestamp", types.YLeaf{"Cdot3OamEventLogTimestamp", cdot3OamEventLogEntry.Cdot3OamEventLogTimestamp})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogOui", types.YLeaf{"Cdot3OamEventLogOui", cdot3OamEventLogEntry.Cdot3OamEventLogOui})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogType", types.YLeaf{"Cdot3OamEventLogType", cdot3OamEventLogEntry.Cdot3OamEventLogType})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogLocation", types.YLeaf{"Cdot3OamEventLogLocation", cdot3OamEventLogEntry.Cdot3OamEventLogLocation})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogWindowHi", types.YLeaf{"Cdot3OamEventLogWindowHi", cdot3OamEventLogEntry.Cdot3OamEventLogWindowHi})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogWindowLo", types.YLeaf{"Cdot3OamEventLogWindowLo", cdot3OamEventLogEntry.Cdot3OamEventLogWindowLo})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogThresholdHi", types.YLeaf{"Cdot3OamEventLogThresholdHi", cdot3OamEventLogEntry.Cdot3OamEventLogThresholdHi})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogThresholdLo", types.YLeaf{"Cdot3OamEventLogThresholdLo", cdot3OamEventLogEntry.Cdot3OamEventLogThresholdLo})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogValue", types.YLeaf{"Cdot3OamEventLogValue", cdot3OamEventLogEntry.Cdot3OamEventLogValue})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogRunningTotal", types.YLeaf{"Cdot3OamEventLogRunningTotal", cdot3OamEventLogEntry.Cdot3OamEventLogRunningTotal})
    cdot3OamEventLogEntry.EntityData.Leafs.Append("cdot3OamEventLogEventTotal", types.YLeaf{"Cdot3OamEventLogEventTotal", cdot3OamEventLogEntry.Cdot3OamEventLogEventTotal})

    cdot3OamEventLogEntry.EntityData.YListKeys = []string {"IfIndex", "Cdot3OamEventLogIndex"}

    return &(cdot3OamEventLogEntry.EntityData)
}

// CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry_Cdot3OamEventLogLocation represents received from the OAM peer via Ethernet OAM (remote(2)).
type CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry_Cdot3OamEventLogLocation string

const (
    CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry_Cdot3OamEventLogLocation_local CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry_Cdot3OamEventLogLocation = "local"

    CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry_Cdot3OamEventLogLocation_remote CISCODOT3OAMMIB_Cdot3OamEventLogTable_Cdot3OamEventLogEntry_Cdot3OamEventLogLocation = "remote"
)

