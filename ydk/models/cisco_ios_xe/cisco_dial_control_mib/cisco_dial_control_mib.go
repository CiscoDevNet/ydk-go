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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cpeerglobalconfiguration CISCODIALCONTROLMIB_Cpeerglobalconfiguration

    // A table containing information about specific calls to a specific
    // destination.
    Ccallhistorytable CISCODIALCONTROLMIB_Ccallhistorytable

    // This table contains information about Internal Error Code(s) (IEC) which
    // caused the call to fail.
    Ccallhistoryiectable CISCODIALCONTROLMIB_Ccallhistoryiectable
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetFilter() yfilter.YFilter { return cISCODIALCONTROLMIB.YFilter }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) SetFilter(yf yfilter.YFilter) { cISCODIALCONTROLMIB.YFilter = yf }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetGoName(yname string) string {
    if yname == "cPeerGlobalConfiguration" { return "Cpeerglobalconfiguration" }
    if yname == "cCallHistoryTable" { return "Ccallhistorytable" }
    if yname == "cCallHistoryIecTable" { return "Ccallhistoryiectable" }
    return ""
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetSegmentPath() string {
    return "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB"
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cPeerGlobalConfiguration" {
        return &cISCODIALCONTROLMIB.Cpeerglobalconfiguration
    }
    if childYangName == "cCallHistoryTable" {
        return &cISCODIALCONTROLMIB.Ccallhistorytable
    }
    if childYangName == "cCallHistoryIecTable" {
        return &cISCODIALCONTROLMIB.Ccallhistoryiectable
    }
    return nil
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cPeerGlobalConfiguration"] = &cISCODIALCONTROLMIB.Cpeerglobalconfiguration
    children["cCallHistoryTable"] = &cISCODIALCONTROLMIB.Ccallhistorytable
    children["cCallHistoryIecTable"] = &cISCODIALCONTROLMIB.Ccallhistoryiectable
    return children
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetYangName() string { return "CISCO-DIAL-CONTROL-MIB" }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) SetParent(parent types.Entity) { cISCODIALCONTROLMIB.parent = parent }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetParent() types.Entity { return cISCODIALCONTROLMIB.parent }

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetParentYangName() string { return "CISCO-DIAL-CONTROL-MIB" }

// CISCODIALCONTROLMIB_Cpeerglobalconfiguration
type CISCODIALCONTROLMIB_Cpeerglobalconfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies the peer search preference based on the type of peers in an
    // universal/integrated port platform.  none      - both voice and data peers
    // are searched            in same preference. datavoice - search data peers
    // first. If no data peers            are found, the voice peers are searched.
    // voicedata - search voice peers first. If no voice peers            are
    // found, the data peers are searched. The type is Cpeersearchtype.
    Cpeersearchtype interface{}
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetFilter() yfilter.YFilter { return cpeerglobalconfiguration.YFilter }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) SetFilter(yf yfilter.YFilter) { cpeerglobalconfiguration.YFilter = yf }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetGoName(yname string) string {
    if yname == "cPeerSearchType" { return "Cpeersearchtype" }
    return ""
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetSegmentPath() string {
    return "cPeerGlobalConfiguration"
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cPeerSearchType"] = cpeerglobalconfiguration.Cpeersearchtype
    return leafs
}

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetBundleName() string { return "cisco_ios_xe" }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetYangName() string { return "cPeerGlobalConfiguration" }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) SetParent(parent types.Entity) { cpeerglobalconfiguration.parent = parent }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetParent() types.Entity { return cpeerglobalconfiguration.parent }

func (cpeerglobalconfiguration *CISCODIALCONTROLMIB_Cpeerglobalconfiguration) GetParentYangName() string { return "CISCO-DIAL-CONTROL-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single Connection. The type is slice of
    // CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry.
    Ccallhistoryentry []CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetFilter() yfilter.YFilter { return ccallhistorytable.YFilter }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) SetFilter(yf yfilter.YFilter) { ccallhistorytable.YFilter = yf }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetGoName(yname string) string {
    if yname == "cCallHistoryEntry" { return "Ccallhistoryentry" }
    return ""
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetSegmentPath() string {
    return "cCallHistoryTable"
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cCallHistoryEntry" {
        for _, c := range ccallhistorytable.Ccallhistoryentry {
            if ccallhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry{}
        ccallhistorytable.Ccallhistoryentry = append(ccallhistorytable.Ccallhistoryentry, child)
        return &ccallhistorytable.Ccallhistoryentry[len(ccallhistorytable.Ccallhistoryentry)-1]
    }
    return nil
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccallhistorytable.Ccallhistoryentry {
        children[ccallhistorytable.Ccallhistoryentry[i].GetSegmentPath()] = &ccallhistorytable.Ccallhistoryentry[i]
    }
    return children
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetYangName() string { return "cCallHistoryTable" }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) SetParent(parent types.Entity) { ccallhistorytable.parent = parent }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetParent() types.Entity { return ccallhistorytable.parent }

func (ccallhistorytable *CISCODIALCONTROLMIB_Ccallhistorytable) GetParentYangName() string { return "CISCO-DIAL-CONTROL-MIB" }

// CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry
// The information regarding a single Connection.
type CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry struct {
    parent types.Entity
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

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetFilter() yfilter.YFilter { return ccallhistoryentry.YFilter }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) SetFilter(yf yfilter.YFilter) { ccallhistoryentry.YFilter = yf }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetGoName(yname string) string {
    if yname == "cCallHistoryIndex" { return "Ccallhistoryindex" }
    if yname == "cCallHistorySetupTime" { return "Ccallhistorysetuptime" }
    if yname == "cCallHistoryPeerAddress" { return "Ccallhistorypeeraddress" }
    if yname == "cCallHistoryPeerSubAddress" { return "Ccallhistorypeersubaddress" }
    if yname == "cCallHistoryPeerId" { return "Ccallhistorypeerid" }
    if yname == "cCallHistoryPeerIfIndex" { return "Ccallhistorypeerifindex" }
    if yname == "cCallHistoryLogicalIfIndex" { return "Ccallhistorylogicalifindex" }
    if yname == "cCallHistoryDisconnectCause" { return "Ccallhistorydisconnectcause" }
    if yname == "cCallHistoryDisconnectText" { return "Ccallhistorydisconnecttext" }
    if yname == "cCallHistoryConnectTime" { return "Ccallhistoryconnecttime" }
    if yname == "cCallHistoryDisconnectTime" { return "Ccallhistorydisconnecttime" }
    if yname == "cCallHistoryCallOrigin" { return "Ccallhistorycallorigin" }
    if yname == "cCallHistoryChargedUnits" { return "Ccallhistorychargedunits" }
    if yname == "cCallHistoryInfoType" { return "Ccallhistoryinfotype" }
    if yname == "cCallHistoryTransmitPackets" { return "Ccallhistorytransmitpackets" }
    if yname == "cCallHistoryTransmitBytes" { return "Ccallhistorytransmitbytes" }
    if yname == "cCallHistoryReceivePackets" { return "Ccallhistoryreceivepackets" }
    if yname == "cCallHistoryReceiveBytes" { return "Ccallhistoryreceivebytes" }
    if yname == "cCallHistoryReleaseSource" { return "Ccallhistoryreleasesource" }
    if yname == "cCallHistoryReleaseSrc" { return "Ccallhistoryreleasesrc" }
    return ""
}

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetSegmentPath() string {
    return "cCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", ccallhistoryentry.Ccallhistoryindex) + "']"
}

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cCallHistoryIndex"] = ccallhistoryentry.Ccallhistoryindex
    leafs["cCallHistorySetupTime"] = ccallhistoryentry.Ccallhistorysetuptime
    leafs["cCallHistoryPeerAddress"] = ccallhistoryentry.Ccallhistorypeeraddress
    leafs["cCallHistoryPeerSubAddress"] = ccallhistoryentry.Ccallhistorypeersubaddress
    leafs["cCallHistoryPeerId"] = ccallhistoryentry.Ccallhistorypeerid
    leafs["cCallHistoryPeerIfIndex"] = ccallhistoryentry.Ccallhistorypeerifindex
    leafs["cCallHistoryLogicalIfIndex"] = ccallhistoryentry.Ccallhistorylogicalifindex
    leafs["cCallHistoryDisconnectCause"] = ccallhistoryentry.Ccallhistorydisconnectcause
    leafs["cCallHistoryDisconnectText"] = ccallhistoryentry.Ccallhistorydisconnecttext
    leafs["cCallHistoryConnectTime"] = ccallhistoryentry.Ccallhistoryconnecttime
    leafs["cCallHistoryDisconnectTime"] = ccallhistoryentry.Ccallhistorydisconnecttime
    leafs["cCallHistoryCallOrigin"] = ccallhistoryentry.Ccallhistorycallorigin
    leafs["cCallHistoryChargedUnits"] = ccallhistoryentry.Ccallhistorychargedunits
    leafs["cCallHistoryInfoType"] = ccallhistoryentry.Ccallhistoryinfotype
    leafs["cCallHistoryTransmitPackets"] = ccallhistoryentry.Ccallhistorytransmitpackets
    leafs["cCallHistoryTransmitBytes"] = ccallhistoryentry.Ccallhistorytransmitbytes
    leafs["cCallHistoryReceivePackets"] = ccallhistoryentry.Ccallhistoryreceivepackets
    leafs["cCallHistoryReceiveBytes"] = ccallhistoryentry.Ccallhistoryreceivebytes
    leafs["cCallHistoryReleaseSource"] = ccallhistoryentry.Ccallhistoryreleasesource
    leafs["cCallHistoryReleaseSrc"] = ccallhistoryentry.Ccallhistoryreleasesrc
    return leafs
}

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetYangName() string { return "cCallHistoryEntry" }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) SetParent(parent types.Entity) { ccallhistoryentry.parent = parent }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetParent() types.Entity { return ccallhistoryentry.parent }

func (ccallhistoryentry *CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry) GetParentYangName() string { return "cCallHistoryTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // The IEC information regarding a single call. The type is slice of
    // CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry.
    Ccallhistoryiecentry []CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetFilter() yfilter.YFilter { return ccallhistoryiectable.YFilter }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) SetFilter(yf yfilter.YFilter) { ccallhistoryiectable.YFilter = yf }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetGoName(yname string) string {
    if yname == "cCallHistoryIecEntry" { return "Ccallhistoryiecentry" }
    return ""
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetSegmentPath() string {
    return "cCallHistoryIecTable"
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cCallHistoryIecEntry" {
        for _, c := range ccallhistoryiectable.Ccallhistoryiecentry {
            if ccallhistoryiectable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry{}
        ccallhistoryiectable.Ccallhistoryiecentry = append(ccallhistoryiectable.Ccallhistoryiecentry, child)
        return &ccallhistoryiectable.Ccallhistoryiecentry[len(ccallhistoryiectable.Ccallhistoryiecentry)-1]
    }
    return nil
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccallhistoryiectable.Ccallhistoryiecentry {
        children[ccallhistoryiectable.Ccallhistoryiecentry[i].GetSegmentPath()] = &ccallhistoryiectable.Ccallhistoryiecentry[i]
    }
    return children
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetBundleName() string { return "cisco_ios_xe" }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetYangName() string { return "cCallHistoryIecTable" }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) SetParent(parent types.Entity) { ccallhistoryiectable.parent = parent }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetParent() types.Entity { return ccallhistoryiectable.parent }

func (ccallhistoryiectable *CISCODIALCONTROLMIB_Ccallhistoryiectable) GetParentYangName() string { return "CISCO-DIAL-CONTROL-MIB" }

// CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry
// The IEC information regarding a single call.
type CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry struct {
    parent types.Entity
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

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetFilter() yfilter.YFilter { return ccallhistoryiecentry.YFilter }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) SetFilter(yf yfilter.YFilter) { ccallhistoryiecentry.YFilter = yf }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetGoName(yname string) string {
    if yname == "cCallHistoryIndex" { return "Ccallhistoryindex" }
    if yname == "cCallHistoryIecIndex" { return "Ccallhistoryiecindex" }
    if yname == "cCallHistoryIec" { return "Ccallhistoryiec" }
    return ""
}

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetSegmentPath() string {
    return "cCallHistoryIecEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", ccallhistoryiecentry.Ccallhistoryindex) + "']" + "[cCallHistoryIecIndex='" + fmt.Sprintf("%v", ccallhistoryiecentry.Ccallhistoryiecindex) + "']"
}

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cCallHistoryIndex"] = ccallhistoryiecentry.Ccallhistoryindex
    leafs["cCallHistoryIecIndex"] = ccallhistoryiecentry.Ccallhistoryiecindex
    leafs["cCallHistoryIec"] = ccallhistoryiecentry.Ccallhistoryiec
    return leafs
}

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetYangName() string { return "cCallHistoryIecEntry" }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) SetParent(parent types.Entity) { ccallhistoryiecentry.parent = parent }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetParent() types.Entity { return ccallhistoryiecentry.parent }

func (ccallhistoryiecentry *CISCODIALCONTROLMIB_Ccallhistoryiectable_Ccallhistoryiecentry) GetParentYangName() string { return "cCallHistoryIecTable" }

