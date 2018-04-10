// The MIB module to describe call history information for
// demand access and possibly other kinds of interfaces.
package cisco_dial_control_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_dial_control_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-DIAL-CONTROL-MIB CISCO-DIAL-CONTROL-MIB}", reflect.TypeOf(CISCODIALCONTROLMIB{}))
    ydk.RegisterEntity("CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB", reflect.TypeOf(CISCODIALCONTROLMIB{}))
}

// CISCODIALCONTROLMIB
type CISCODIALCONTROLMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cpeerglobalconfiguration CISCODIALCONTROLMIB_Cpeerglobalconfiguration

    // A table containing information about specific calls to a specific
    // destination.
    Ccallhistorytable CISCODIALCONTROLMIB_Ccallhistorytable

    // This table contains information about Internal Error Code(s) (IEC) which
    // caused the call to fail.
    Ccallhistoryiectable CISCODIALCONTROLMIB_Ccallhistoryiectable
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCODIALCONTROLMIB.EntityData.YFilter = cISCODIALCONTROLMIB.YFilter
    cISCODIALCONTROLMIB.EntityData.YangName = "CISCO-DIAL-CONTROL-MIB"
    cISCODIALCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODIALCONTROLMIB.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    cISCODIALCONTROLMIB.EntityData.SegmentPath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB"
    cISCODIALCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODIALCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODIALCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODIALCONTROLMIB.EntityData.Children = make(map[string]types.YChild)
    cISCODIALCONTROLMIB.EntityData.Children["cPeerGlobalConfiguration"] = types.YChild{"Cpeerglobalconfiguration", &cISCODIALCONTROLMIB.Cpeerglobalconfiguration}
    cISCODIALCONTROLMIB.EntityData.Children["cCallHistoryTable"] = types.YChild{"Ccallhistorytable", &cISCODIALCONTROLMIB.Ccallhistorytable}
    cISCODIALCONTROLMIB.EntityData.Children["cCallHistoryIecTable"] = types.YChild{"Ccallhistoryiectable", &cISCODIALCONTROLMIB.Ccallhistoryiectable}
    cISCODIALCONTROLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCODIALCONTROLMIB.EntityData)
}

// CISCODIALCONTROLMIB_Cpeerglobalconfiguration
type CISCODIALCONTROLMIB_Cpeerglobalconfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies the peer search preference based on the type of peers in an
    // universal/integrated port platform.  none      - both voice and data peers
    // are searched            in same preference. datavoice - search data peers
    // first. If no data peers            are found, the voice peers are searched.
    // voicedata - search voice peers first. If no voice peers            are
    // found, the data peers are searched. The type is Cpeersearchtype.
    Cpeersearchtype interface{}
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetEntityData() *types.CommonEntityData {
    cpeerglobalconfiguration.EntityData.YFilter = cpeerglobalconfiguration.YFilter
    cpeerglobalconfiguration.EntityData.YangName = "cPeerGlobalConfiguration"
    cpeerglobalconfiguration.EntityData.BundleName = "cisco_ios_xe"
    cpeerglobalconfiguration.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    cpeerglobalconfiguration.EntityData.SegmentPath = "cPeerGlobalConfiguration"
    cpeerglobalconfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeerglobalconfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeerglobalconfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeerglobalconfiguration.EntityData.Children = make(map[string]types.YChild)
    cpeerglobalconfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    cpeerglobalconfiguration.EntityData.Leafs["cPeerSearchType"] = types.YLeaf{"Cpeersearchtype", cpeerglobalconfiguration.Cpeersearchtype}
    return &(cpeerglobalconfiguration.EntityData)
}

// CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype represents            are found, the data peers are searched.
type CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype string

const (
    CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype_none CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype = "none"

    CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype_datavoice CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype = "datavoice"

    CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype_voicedata CISCODIALCONTROLMIB_Cpeerglobalconfiguration_Cpeersearchtype = "voicedata"
)

// CISCODIALCONTROLMIB_Ccallhistorytable
// A table containing information about specific
// calls to a specific destination.
type CISCODIALCONTROLMIB_Ccallhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single Connection. The type is slice of
    // CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry.
    Ccallhistoryentry []CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetEntityData() *types.CommonEntityData {
    ccallhistorytable.EntityData.YFilter = ccallhistorytable.YFilter
    ccallhistorytable.EntityData.YangName = "cCallHistoryTable"
    ccallhistorytable.EntityData.BundleName = "cisco_ios_xe"
    ccallhistorytable.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    ccallhistorytable.EntityData.SegmentPath = "cCallHistoryTable"
    ccallhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccallhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccallhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccallhistorytable.EntityData.Children = make(map[string]types.YChild)
    ccallhistorytable.EntityData.Children["cCallHistoryEntry"] = types.YChild{"Ccallhistoryentry", nil}
    for i := range ccallhistorytable.Ccallhistoryentry {
        ccallhistorytable.EntityData.Children[types.GetSegmentPath(&ccallhistorytable.Ccallhistoryentry[i])] = types.YChild{"Ccallhistoryentry", &ccallhistorytable.Ccallhistoryentry[i]}
    }
    ccallhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccallhistorytable.EntityData)
}

// CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry
// The information regarding a single Connection.
type CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing call disconnection events.  When it reaches the 
    // maximum value, an extremely unlikely event, the agent wraps  the value back
    // to 1 and may flush existing entries. The type is interface{} with range:
    // 1..4294967295.
    Ccallhistoryindex interface{}

    // The value of sysUpTime when the call setup was started. This will be useful
    // for an NMS to sort the call history entry with call setup time. Also, this
    // object can be useful in finding large delays between the time the call was
    // started and the time the call was connected. For ISDN media, this will be
    // the time when the setup message was received from or sent to the network.
    // The value of this object is the same as callActiveSetupTime in the
    // callActiveTable. The type is interface{} with range: 0..4294967295.
    Ccallhistorysetuptime interface{}

    // The number this call was connected to. If the number is not available, then
    // it will have a length of zero. The type is string.
    Ccallhistorypeeraddress interface{}

    // The subaddress this call was connected to. If the subaddress is undefined
    // or not available, this will be a zero length string. The type is string.
    Ccallhistorypeersubaddress interface{}

    // This is the Id value of the peer table entry to which this call was made.
    // If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    Ccallhistorypeerid interface{}

    // This is the ifIndex value of the peer table entry to which this call was
    // made. If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    Ccallhistorypeerifindex interface{}

    // This is the ifIndex value of the logical interface through which this call
    // was made. For ISDN media, this would be the ifIndex of the B channel which
    // was used for this call. If the ifIndex value is unknown, the value of this
    // object  will be zero. For an IP call, the value will be zero. The type is
    // interface{} with range: 0..2147483647.
    Ccallhistorylogicalifindex interface{}

    // The encoded network cause value associated with this call.  The value of
    // this object will depend on the interface type as well as on the protocol
    // and protocol version being used on this interface. Some references for
    // possible cause values are given below. The type is string with length:
    // 0..4.
    Ccallhistorydisconnectcause interface{}

    // ASCII text describing the reason for call termination.  This object exists
    // because it would be impossible for a management station to store all
    // possible cause values for all types of interfaces. It should be used only
    // if a management station is unable to decode the value of
    // dialCtlPeerStatsLastDisconnectCause. The type is string.
    Ccallhistorydisconnecttext interface{}

    // The value of sysUpTime when the call was connected. The type is interface{}
    // with range: 0..4294967295.
    Ccallhistoryconnecttime interface{}

    // The value of sysUpTime when the call was disconnected. The type is
    // interface{} with range: 0..4294967295.
    Ccallhistorydisconnecttime interface{}

    // The call origin. The type is Ccallhistorycallorigin.
    Ccallhistorycallorigin interface{}

    // The number of charged units for this connection. For incoming calls or if
    // charging information is not supplied by the switch, the value of this
    // object will be zero. The type is interface{} with range: 0..4294967295.
    Ccallhistorychargedunits interface{}

    // The information type for this call. The type is Ccallhistoryinfotype.
    Ccallhistoryinfotype interface{}

    // The number of packets which were transmitted while this call was active.
    // The type is interface{} with range: 0..4294967295.
    Ccallhistorytransmitpackets interface{}

    // The number of bytes which were transmitted while this call was active. The
    // type is interface{} with range: 0..4294967295.
    Ccallhistorytransmitbytes interface{}

    // The number of packets which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    Ccallhistoryreceivepackets interface{}

    // The number of bytes which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    Ccallhistoryreceivebytes interface{}

    // Originator of the call release. The type is Ccallhistoryreleasesource.
    Ccallhistoryreleasesource interface{}

    // Originator of the call release. Indicates the source of  the call release
    // as follows :  1) callingPartyInPstn : Calling party in PSTN. 2)
    // callingPartyInVoip : Calling party in VoIP. 3) calledPartyInPstn : Called
    // party in PSTN. 4) calledPartyInVoip : Called party in VoIP. 5)
    // internalReleaseInPotsLeg : Due to an internal error     in Telephony call
    // leg. 6) internalReleaseInVoipLeg : Due to an internal error    in VoIP call
    // leg. 7) internalCallControlApp : Due to an internal error    in Session
    // Application or Tcl or VXML script originated    release.  8)
    // internalReleaseInVoipAAA : Due to an internal error    in VoIP AAA module.
    // 9) consoleCommand : Due to CLI or MML. 10) externalRadiusServer : Call
    // failed during authorization     , authentication or due to receipt of POD
    // from the      RADIUS server. 11) externalNmsApp : Due to SNMP request to
    // clear      the call. 12) externalCallControlAgent : External Call Control
    // Agent     initiated clear. 13) gatekeeper : Gatekeeper initiated clear due
    // to receipt     of Admission Reject, Disengage Request message. 14)
    // externalGKTMPServer : External GKTMP server initiated     clear due to
    // receipt of Admission Reject message from     the gatekeeper, triggered by
    // RESPONSE.ARJ message from     the GKTMP server. The type is
    // Ccallhistoryreleasesrc.
    Ccallhistoryreleasesrc interface{}
}

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetEntityData() *types.CommonEntityData {
    ccallhistoryentry.EntityData.YFilter = ccallhistoryentry.YFilter
    ccallhistoryentry.EntityData.YangName = "cCallHistoryEntry"
    ccallhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    ccallhistoryentry.EntityData.ParentYangName = "cCallHistoryTable"
    ccallhistoryentry.EntityData.SegmentPath = "cCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", ccallhistoryentry.Ccallhistoryindex) + "']"
    ccallhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccallhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccallhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccallhistoryentry.EntityData.Children = make(map[string]types.YChild)
    ccallhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccallhistoryentry.EntityData.Leafs["cCallHistoryIndex"] = types.YLeaf{"Ccallhistoryindex", ccallhistoryentry.Ccallhistoryindex}
    ccallhistoryentry.EntityData.Leafs["cCallHistorySetupTime"] = types.YLeaf{"Ccallhistorysetuptime", ccallhistoryentry.Ccallhistorysetuptime}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryPeerAddress"] = types.YLeaf{"Ccallhistorypeeraddress", ccallhistoryentry.Ccallhistorypeeraddress}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryPeerSubAddress"] = types.YLeaf{"Ccallhistorypeersubaddress", ccallhistoryentry.Ccallhistorypeersubaddress}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryPeerId"] = types.YLeaf{"Ccallhistorypeerid", ccallhistoryentry.Ccallhistorypeerid}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryPeerIfIndex"] = types.YLeaf{"Ccallhistorypeerifindex", ccallhistoryentry.Ccallhistorypeerifindex}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryLogicalIfIndex"] = types.YLeaf{"Ccallhistorylogicalifindex", ccallhistoryentry.Ccallhistorylogicalifindex}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryDisconnectCause"] = types.YLeaf{"Ccallhistorydisconnectcause", ccallhistoryentry.Ccallhistorydisconnectcause}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryDisconnectText"] = types.YLeaf{"Ccallhistorydisconnecttext", ccallhistoryentry.Ccallhistorydisconnecttext}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryConnectTime"] = types.YLeaf{"Ccallhistoryconnecttime", ccallhistoryentry.Ccallhistoryconnecttime}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryDisconnectTime"] = types.YLeaf{"Ccallhistorydisconnecttime", ccallhistoryentry.Ccallhistorydisconnecttime}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryCallOrigin"] = types.YLeaf{"Ccallhistorycallorigin", ccallhistoryentry.Ccallhistorycallorigin}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryChargedUnits"] = types.YLeaf{"Ccallhistorychargedunits", ccallhistoryentry.Ccallhistorychargedunits}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryInfoType"] = types.YLeaf{"Ccallhistoryinfotype", ccallhistoryentry.Ccallhistoryinfotype}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryTransmitPackets"] = types.YLeaf{"Ccallhistorytransmitpackets", ccallhistoryentry.Ccallhistorytransmitpackets}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryTransmitBytes"] = types.YLeaf{"Ccallhistorytransmitbytes", ccallhistoryentry.Ccallhistorytransmitbytes}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryReceivePackets"] = types.YLeaf{"Ccallhistoryreceivepackets", ccallhistoryentry.Ccallhistoryreceivepackets}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryReceiveBytes"] = types.YLeaf{"Ccallhistoryreceivebytes", ccallhistoryentry.Ccallhistoryreceivebytes}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryReleaseSource"] = types.YLeaf{"Ccallhistoryreleasesource", ccallhistoryentry.Ccallhistoryreleasesource}
    ccallhistoryentry.EntityData.Leafs["cCallHistoryReleaseSrc"] = types.YLeaf{"Ccallhistoryreleasesrc", ccallhistoryentry.Ccallhistoryreleasesrc}
    return &(ccallhistoryentry.EntityData)
}

// CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin represents The call origin.
type CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin string

const (
    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin_originate CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin = "originate"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin_answer CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin = "answer"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin_callback CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistorycallorigin = "callback"
)

// CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype represents The information type for this call.
type CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype string

const (
    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_other CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "other"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_speech CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "speech"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_unrestrictedDigital CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "unrestrictedDigital"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_unrestrictedDigital56 CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "unrestrictedDigital56"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_restrictedDigital CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "restrictedDigital"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_audio31 CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "audio31"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_audio7 CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "audio7"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_video CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "video"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_packetSwitched CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "packetSwitched"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype_fax CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryinfotype = "fax"
)

// CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource represents Originator of the call release.
type CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource string

const (
    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_callingPartyInPstn CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "callingPartyInPstn"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_callingPartyInVoip CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "callingPartyInVoip"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_calledPartyInPstn CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "calledPartyInPstn"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_calledPartyInVoip CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "calledPartyInVoip"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_internalRelease CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "internalRelease"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_internalCallControlApp CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "internalCallControlApp"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_consoleCommand CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "consoleCommand"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_externalRadiusServer CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "externalRadiusServer"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_externalNmsApp CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "externalNmsApp"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource_externalCallControlAgent CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesource = "externalCallControlAgent"
)

// CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc represents     the GKTMP server.
type CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc string

const (
    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_callingPartyInPstn CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "callingPartyInPstn"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_callingPartyInVoip CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "callingPartyInVoip"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_calledPartyInPstn CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "calledPartyInPstn"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_calledPartyInVoip CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "calledPartyInVoip"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_internalReleaseInPotsLeg CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "internalReleaseInPotsLeg"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_internalReleaseInVoipLeg CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "internalReleaseInVoipLeg"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_internalCallControlApp CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "internalCallControlApp"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_internalReleaseInVoipAAA CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "internalReleaseInVoipAAA"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_consoleCommand CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "consoleCommand"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_externalRadiusServer CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "externalRadiusServer"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_externalNmsApp CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "externalNmsApp"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_externalCallControlAgent CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "externalCallControlAgent"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_gatekeeper CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "gatekeeper"

    CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc_externalGKTMPServer CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryreleasesrc = "externalGKTMPServer"
)

// CISCODIALCONTROLMIB_Ccallhistoryiectable
// This table contains information about Internal Error
// Code(s) (IEC) which caused the call to fail.
type CISCODIALCONTROLMIB_Ccallhistoryiectable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IEC information regarding a single call. The type is slice of
    // CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry.
    Ccallhistoryiecentry []CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetEntityData() *types.CommonEntityData {
    ccallhistoryiectable.EntityData.YFilter = ccallhistoryiectable.YFilter
    ccallhistoryiectable.EntityData.YangName = "cCallHistoryIecTable"
    ccallhistoryiectable.EntityData.BundleName = "cisco_ios_xe"
    ccallhistoryiectable.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    ccallhistoryiectable.EntityData.SegmentPath = "cCallHistoryIecTable"
    ccallhistoryiectable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccallhistoryiectable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccallhistoryiectable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccallhistoryiectable.EntityData.Children = make(map[string]types.YChild)
    ccallhistoryiectable.EntityData.Children["cCallHistoryIecEntry"] = types.YChild{"Ccallhistoryiecentry", nil}
    for i := range ccallhistoryiectable.Ccallhistoryiecentry {
        ccallhistoryiectable.EntityData.Children[types.GetSegmentPath(&ccallhistoryiectable.Ccallhistoryiecentry[i])] = types.YChild{"Ccallhistoryiecentry", &ccallhistoryiectable.Ccallhistoryiecentry[i]}
    }
    ccallhistoryiectable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccallhistoryiectable.EntityData)
}

// CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry
// The IEC information regarding a single call.
type CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryindex
    Ccallhistoryindex interface{}

    // This attribute is a key. This object is used to index one or more IECs in
    // the context of a single call.  In most cases there will only be one IEC
    // when a call fails.  However, it is possible for the software processing the
    // call to  generate multiple IECs before the call ultimately fails. In that
    // scenario, there will be multiple entries in this table related to a single
    // call (cCallHistoryIndex) and this object will serve to uniquely identify
    // each IEC. The type is interface{} with range: 1..1024.
    Ccallhistoryiecindex interface{}

    // This object reflects the Internal Error Code. The format is a string of
    // dotted decimal numbers composed of the following components: 
    // Version.Entity.Category.Subsystem.Errorcode.Diagnostic  Each component is
    // defined as follows: Version     : The version of IEC software. Entity     
    // : The network entity that originated               the error. Category    :
    // The category of the error (eg, software               error, hardware
    // resource unavailable, ...) Subsystem   : The subsystem in which the error
    // occurred. Errorcode   : A subsytem-specific error code. Diagnostic  : An
    // implementation-specific diagnostic code. The type is string.
    Ccallhistoryiec interface{}
}

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetEntityData() *types.CommonEntityData {
    ccallhistoryiecentry.EntityData.YFilter = ccallhistoryiecentry.YFilter
    ccallhistoryiecentry.EntityData.YangName = "cCallHistoryIecEntry"
    ccallhistoryiecentry.EntityData.BundleName = "cisco_ios_xe"
    ccallhistoryiecentry.EntityData.ParentYangName = "cCallHistoryIecTable"
    ccallhistoryiecentry.EntityData.SegmentPath = "cCallHistoryIecEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", ccallhistoryiecentry.Ccallhistoryindex) + "']" + "[cCallHistoryIecIndex='" + fmt.Sprintf("%v", ccallhistoryiecentry.Ccallhistoryiecindex) + "']"
    ccallhistoryiecentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccallhistoryiecentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccallhistoryiecentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccallhistoryiecentry.EntityData.Children = make(map[string]types.YChild)
    ccallhistoryiecentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccallhistoryiecentry.EntityData.Leafs["cCallHistoryIndex"] = types.YLeaf{"Ccallhistoryindex", ccallhistoryiecentry.Ccallhistoryindex}
    ccallhistoryiecentry.EntityData.Leafs["cCallHistoryIecIndex"] = types.YLeaf{"Ccallhistoryiecindex", ccallhistoryiecentry.Ccallhistoryiecindex}
    ccallhistoryiecentry.EntityData.Leafs["cCallHistoryIec"] = types.YLeaf{"Ccallhistoryiec", ccallhistoryiecentry.Ccallhistoryiec}
    return &(ccallhistoryiecentry.EntityData)
}

