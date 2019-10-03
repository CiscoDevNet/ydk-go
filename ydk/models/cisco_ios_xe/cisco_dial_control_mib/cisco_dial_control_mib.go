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

    
    CPeerGlobalConfiguration CISCODIALCONTROLMIB_CPeerGlobalConfiguration

    // A table containing information about specific calls to a specific
    // destination.
    CCallHistoryTable CISCODIALCONTROLMIB_CCallHistoryTable

    // This table contains information about Internal Error Code(s) (IEC) which
    // caused the call to fail.
    CCallHistoryIecTable CISCODIALCONTROLMIB_CCallHistoryIecTable
}

func (cISCODIALCONTROLMIB *CISCODIALCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCODIALCONTROLMIB.EntityData.YFilter = cISCODIALCONTROLMIB.YFilter
    cISCODIALCONTROLMIB.EntityData.YangName = "CISCO-DIAL-CONTROL-MIB"
    cISCODIALCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODIALCONTROLMIB.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    cISCODIALCONTROLMIB.EntityData.SegmentPath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB"
    cISCODIALCONTROLMIB.EntityData.AbsolutePath = cISCODIALCONTROLMIB.EntityData.SegmentPath
    cISCODIALCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODIALCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODIALCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODIALCONTROLMIB.EntityData.Children = types.NewOrderedMap()
    cISCODIALCONTROLMIB.EntityData.Children.Append("cPeerGlobalConfiguration", types.YChild{"CPeerGlobalConfiguration", &cISCODIALCONTROLMIB.CPeerGlobalConfiguration})
    cISCODIALCONTROLMIB.EntityData.Children.Append("cCallHistoryTable", types.YChild{"CCallHistoryTable", &cISCODIALCONTROLMIB.CCallHistoryTable})
    cISCODIALCONTROLMIB.EntityData.Children.Append("cCallHistoryIecTable", types.YChild{"CCallHistoryIecTable", &cISCODIALCONTROLMIB.CCallHistoryIecTable})
    cISCODIALCONTROLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCODIALCONTROLMIB.EntityData.YListKeys = []string {}

    return &(cISCODIALCONTROLMIB.EntityData)
}

// CISCODIALCONTROLMIB_CPeerGlobalConfiguration
type CISCODIALCONTROLMIB_CPeerGlobalConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies the peer search preference based on the type of peers in an
    // universal/integrated port platform.  none      - both voice and data peers
    // are searched            in same preference. datavoice - search data peers
    // first. If no data peers            are found, the voice peers are searched.
    // voicedata - search voice peers first. If no voice peers            are
    // found, the data peers are searched. The type is CPeerSearchType.
    CPeerSearchType interface{}
}

func (cPeerGlobalConfiguration *CISCODIALCONTROLMIB_CPeerGlobalConfiguration) GetEntityData() *types.CommonEntityData {
    cPeerGlobalConfiguration.EntityData.YFilter = cPeerGlobalConfiguration.YFilter
    cPeerGlobalConfiguration.EntityData.YangName = "cPeerGlobalConfiguration"
    cPeerGlobalConfiguration.EntityData.BundleName = "cisco_ios_xe"
    cPeerGlobalConfiguration.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    cPeerGlobalConfiguration.EntityData.SegmentPath = "cPeerGlobalConfiguration"
    cPeerGlobalConfiguration.EntityData.AbsolutePath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB/" + cPeerGlobalConfiguration.EntityData.SegmentPath
    cPeerGlobalConfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cPeerGlobalConfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cPeerGlobalConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cPeerGlobalConfiguration.EntityData.Children = types.NewOrderedMap()
    cPeerGlobalConfiguration.EntityData.Leafs = types.NewOrderedMap()
    cPeerGlobalConfiguration.EntityData.Leafs.Append("cPeerSearchType", types.YLeaf{"CPeerSearchType", cPeerGlobalConfiguration.CPeerSearchType})

    cPeerGlobalConfiguration.EntityData.YListKeys = []string {}

    return &(cPeerGlobalConfiguration.EntityData)
}

// CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType represents            are found, the data peers are searched.
type CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType string

const (
    CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType_none CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType = "none"

    CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType_datavoice CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType = "datavoice"

    CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType_voicedata CISCODIALCONTROLMIB_CPeerGlobalConfiguration_CPeerSearchType = "voicedata"
)

// CISCODIALCONTROLMIB_CCallHistoryTable
// A table containing information about specific
// calls to a specific destination.
type CISCODIALCONTROLMIB_CCallHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single Connection. The type is slice of
    // CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry.
    CCallHistoryEntry []*CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry
}

func (cCallHistoryTable *CISCODIALCONTROLMIB_CCallHistoryTable) GetEntityData() *types.CommonEntityData {
    cCallHistoryTable.EntityData.YFilter = cCallHistoryTable.YFilter
    cCallHistoryTable.EntityData.YangName = "cCallHistoryTable"
    cCallHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cCallHistoryTable.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    cCallHistoryTable.EntityData.SegmentPath = "cCallHistoryTable"
    cCallHistoryTable.EntityData.AbsolutePath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB/" + cCallHistoryTable.EntityData.SegmentPath
    cCallHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cCallHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cCallHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cCallHistoryTable.EntityData.Children = types.NewOrderedMap()
    cCallHistoryTable.EntityData.Children.Append("cCallHistoryEntry", types.YChild{"CCallHistoryEntry", nil})
    for i := range cCallHistoryTable.CCallHistoryEntry {
        cCallHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cCallHistoryTable.CCallHistoryEntry[i]), types.YChild{"CCallHistoryEntry", cCallHistoryTable.CCallHistoryEntry[i]})
    }
    cCallHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cCallHistoryTable.EntityData.YListKeys = []string {}

    return &(cCallHistoryTable.EntityData)
}

// CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry
// The information regarding a single Connection.
type CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing call disconnection events.  When it reaches the 
    // maximum value, an extremely unlikely event, the agent wraps  the value back
    // to 1 and may flush existing entries. The type is interface{} with range:
    // 1..4294967295.
    CCallHistoryIndex interface{}

    // The value of sysUpTime when the call setup was started. This will be useful
    // for an NMS to sort the call history entry with call setup time. Also, this
    // object can be useful in finding large delays between the time the call was
    // started and the time the call was connected. For ISDN media, this will be
    // the time when the setup message was received from or sent to the network.
    // The value of this object is the same as callActiveSetupTime in the
    // callActiveTable. The type is interface{} with range: 0..4294967295.
    CCallHistorySetupTime interface{}

    // The number this call was connected to. If the number is not available, then
    // it will have a length of zero. The type is string.
    CCallHistoryPeerAddress interface{}

    // The subaddress this call was connected to. If the subaddress is undefined
    // or not available, this will be a zero length string. The type is string.
    CCallHistoryPeerSubAddress interface{}

    // This is the Id value of the peer table entry to which this call was made.
    // If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    CCallHistoryPeerId interface{}

    // This is the ifIndex value of the peer table entry to which this call was
    // made. If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    CCallHistoryPeerIfIndex interface{}

    // This is the ifIndex value of the logical interface through which this call
    // was made. For ISDN media, this would be the ifIndex of the B channel which
    // was used for this call. If the ifIndex value is unknown, the value of this
    // object  will be zero. For an IP call, the value will be zero. The type is
    // interface{} with range: 0..2147483647.
    CCallHistoryLogicalIfIndex interface{}

    // The encoded network cause value associated with this call.  The value of
    // this object will depend on the interface type as well as on the protocol
    // and protocol version being used on this interface. Some references for
    // possible cause values are given below. The type is string with length:
    // 0..4.
    CCallHistoryDisconnectCause interface{}

    // ASCII text describing the reason for call termination.  This object exists
    // because it would be impossible for a management station to store all
    // possible cause values for all types of interfaces. It should be used only
    // if a management station is unable to decode the value of
    // dialCtlPeerStatsLastDisconnectCause. The type is string.
    CCallHistoryDisconnectText interface{}

    // The value of sysUpTime when the call was connected. The type is interface{}
    // with range: 0..4294967295.
    CCallHistoryConnectTime interface{}

    // The value of sysUpTime when the call was disconnected. The type is
    // interface{} with range: 0..4294967295.
    CCallHistoryDisconnectTime interface{}

    // The call origin. The type is CCallHistoryCallOrigin.
    CCallHistoryCallOrigin interface{}

    // The number of charged units for this connection. For incoming calls or if
    // charging information is not supplied by the switch, the value of this
    // object will be zero. The type is interface{} with range: 0..4294967295.
    CCallHistoryChargedUnits interface{}

    // The information type for this call. The type is CCallHistoryInfoType.
    CCallHistoryInfoType interface{}

    // The number of packets which were transmitted while this call was active.
    // The type is interface{} with range: 0..4294967295.
    CCallHistoryTransmitPackets interface{}

    // The number of bytes which were transmitted while this call was active. The
    // type is interface{} with range: 0..4294967295.
    CCallHistoryTransmitBytes interface{}

    // The number of packets which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    CCallHistoryReceivePackets interface{}

    // The number of bytes which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    CCallHistoryReceiveBytes interface{}

    // Originator of the call release. The type is CCallHistoryReleaseSource.
    CCallHistoryReleaseSource interface{}

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
    // CCallHistoryReleaseSrc.
    CCallHistoryReleaseSrc interface{}
}

func (cCallHistoryEntry *CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry) GetEntityData() *types.CommonEntityData {
    cCallHistoryEntry.EntityData.YFilter = cCallHistoryEntry.YFilter
    cCallHistoryEntry.EntityData.YangName = "cCallHistoryEntry"
    cCallHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cCallHistoryEntry.EntityData.ParentYangName = "cCallHistoryTable"
    cCallHistoryEntry.EntityData.SegmentPath = "cCallHistoryEntry" + types.AddKeyToken(cCallHistoryEntry.CCallHistoryIndex, "cCallHistoryIndex")
    cCallHistoryEntry.EntityData.AbsolutePath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB/cCallHistoryTable/" + cCallHistoryEntry.EntityData.SegmentPath
    cCallHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cCallHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cCallHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cCallHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cCallHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryIndex", types.YLeaf{"CCallHistoryIndex", cCallHistoryEntry.CCallHistoryIndex})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistorySetupTime", types.YLeaf{"CCallHistorySetupTime", cCallHistoryEntry.CCallHistorySetupTime})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryPeerAddress", types.YLeaf{"CCallHistoryPeerAddress", cCallHistoryEntry.CCallHistoryPeerAddress})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryPeerSubAddress", types.YLeaf{"CCallHistoryPeerSubAddress", cCallHistoryEntry.CCallHistoryPeerSubAddress})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryPeerId", types.YLeaf{"CCallHistoryPeerId", cCallHistoryEntry.CCallHistoryPeerId})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryPeerIfIndex", types.YLeaf{"CCallHistoryPeerIfIndex", cCallHistoryEntry.CCallHistoryPeerIfIndex})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryLogicalIfIndex", types.YLeaf{"CCallHistoryLogicalIfIndex", cCallHistoryEntry.CCallHistoryLogicalIfIndex})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryDisconnectCause", types.YLeaf{"CCallHistoryDisconnectCause", cCallHistoryEntry.CCallHistoryDisconnectCause})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryDisconnectText", types.YLeaf{"CCallHistoryDisconnectText", cCallHistoryEntry.CCallHistoryDisconnectText})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryConnectTime", types.YLeaf{"CCallHistoryConnectTime", cCallHistoryEntry.CCallHistoryConnectTime})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryDisconnectTime", types.YLeaf{"CCallHistoryDisconnectTime", cCallHistoryEntry.CCallHistoryDisconnectTime})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryCallOrigin", types.YLeaf{"CCallHistoryCallOrigin", cCallHistoryEntry.CCallHistoryCallOrigin})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryChargedUnits", types.YLeaf{"CCallHistoryChargedUnits", cCallHistoryEntry.CCallHistoryChargedUnits})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryInfoType", types.YLeaf{"CCallHistoryInfoType", cCallHistoryEntry.CCallHistoryInfoType})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryTransmitPackets", types.YLeaf{"CCallHistoryTransmitPackets", cCallHistoryEntry.CCallHistoryTransmitPackets})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryTransmitBytes", types.YLeaf{"CCallHistoryTransmitBytes", cCallHistoryEntry.CCallHistoryTransmitBytes})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryReceivePackets", types.YLeaf{"CCallHistoryReceivePackets", cCallHistoryEntry.CCallHistoryReceivePackets})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryReceiveBytes", types.YLeaf{"CCallHistoryReceiveBytes", cCallHistoryEntry.CCallHistoryReceiveBytes})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryReleaseSource", types.YLeaf{"CCallHistoryReleaseSource", cCallHistoryEntry.CCallHistoryReleaseSource})
    cCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryReleaseSrc", types.YLeaf{"CCallHistoryReleaseSrc", cCallHistoryEntry.CCallHistoryReleaseSrc})

    cCallHistoryEntry.EntityData.YListKeys = []string {"CCallHistoryIndex"}

    return &(cCallHistoryEntry.EntityData)
}

// CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin represents The call origin.
type CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin string

const (
    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin_originate CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin = "originate"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin_answer CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin = "answer"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin_callback CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryCallOrigin = "callback"
)

// CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType represents The information type for this call.
type CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType string

const (
    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_other CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "other"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_speech CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "speech"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_unrestrictedDigital CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "unrestrictedDigital"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_unrestrictedDigital56 CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "unrestrictedDigital56"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_restrictedDigital CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "restrictedDigital"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_audio31 CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "audio31"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_audio7 CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "audio7"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_video CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "video"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_packetSwitched CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "packetSwitched"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType_fax CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryInfoType = "fax"
)

// CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource represents Originator of the call release.
type CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource string

const (
    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_callingPartyInPstn CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "callingPartyInPstn"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_callingPartyInVoip CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "callingPartyInVoip"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_calledPartyInPstn CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "calledPartyInPstn"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_calledPartyInVoip CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "calledPartyInVoip"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_internalRelease CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "internalRelease"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_internalCallControlApp CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "internalCallControlApp"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_consoleCommand CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "consoleCommand"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_externalRadiusServer CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "externalRadiusServer"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_externalNmsApp CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "externalNmsApp"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource_externalCallControlAgent CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSource = "externalCallControlAgent"
)

// CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc represents     the GKTMP server.
type CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc string

const (
    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_callingPartyInPstn CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "callingPartyInPstn"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_callingPartyInVoip CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "callingPartyInVoip"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_calledPartyInPstn CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "calledPartyInPstn"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_calledPartyInVoip CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "calledPartyInVoip"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_internalReleaseInPotsLeg CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "internalReleaseInPotsLeg"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_internalReleaseInVoipLeg CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "internalReleaseInVoipLeg"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_internalCallControlApp CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "internalCallControlApp"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_internalReleaseInVoipAAA CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "internalReleaseInVoipAAA"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_consoleCommand CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "consoleCommand"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_externalRadiusServer CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "externalRadiusServer"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_externalNmsApp CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "externalNmsApp"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_externalCallControlAgent CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "externalCallControlAgent"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_gatekeeper CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "gatekeeper"

    CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc_externalGKTMPServer CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryReleaseSrc = "externalGKTMPServer"
)

// CISCODIALCONTROLMIB_CCallHistoryIecTable
// This table contains information about Internal Error
// Code(s) (IEC) which caused the call to fail.
type CISCODIALCONTROLMIB_CCallHistoryIecTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IEC information regarding a single call. The type is slice of
    // CISCODIALCONTROLMIB_CCallHistoryIecTable_CCallHistoryIecEntry.
    CCallHistoryIecEntry []*CISCODIALCONTROLMIB_CCallHistoryIecTable_CCallHistoryIecEntry
}

func (cCallHistoryIecTable *CISCODIALCONTROLMIB_CCallHistoryIecTable) GetEntityData() *types.CommonEntityData {
    cCallHistoryIecTable.EntityData.YFilter = cCallHistoryIecTable.YFilter
    cCallHistoryIecTable.EntityData.YangName = "cCallHistoryIecTable"
    cCallHistoryIecTable.EntityData.BundleName = "cisco_ios_xe"
    cCallHistoryIecTable.EntityData.ParentYangName = "CISCO-DIAL-CONTROL-MIB"
    cCallHistoryIecTable.EntityData.SegmentPath = "cCallHistoryIecTable"
    cCallHistoryIecTable.EntityData.AbsolutePath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB/" + cCallHistoryIecTable.EntityData.SegmentPath
    cCallHistoryIecTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cCallHistoryIecTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cCallHistoryIecTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cCallHistoryIecTable.EntityData.Children = types.NewOrderedMap()
    cCallHistoryIecTable.EntityData.Children.Append("cCallHistoryIecEntry", types.YChild{"CCallHistoryIecEntry", nil})
    for i := range cCallHistoryIecTable.CCallHistoryIecEntry {
        cCallHistoryIecTable.EntityData.Children.Append(types.GetSegmentPath(cCallHistoryIecTable.CCallHistoryIecEntry[i]), types.YChild{"CCallHistoryIecEntry", cCallHistoryIecTable.CCallHistoryIecEntry[i]})
    }
    cCallHistoryIecTable.EntityData.Leafs = types.NewOrderedMap()

    cCallHistoryIecTable.EntityData.YListKeys = []string {}

    return &(cCallHistoryIecTable.EntityData)
}

// CISCODIALCONTROLMIB_CCallHistoryIecTable_CCallHistoryIecEntry
// The IEC information regarding a single call.
type CISCODIALCONTROLMIB_CCallHistoryIecTable_CCallHistoryIecEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryIndex
    CCallHistoryIndex interface{}

    // This attribute is a key. This object is used to index one or more IECs in
    // the context of a single call.  In most cases there will only be one IEC
    // when a call fails.  However, it is possible for the software processing the
    // call to  generate multiple IECs before the call ultimately fails. In that
    // scenario, there will be multiple entries in this table related to a single
    // call (cCallHistoryIndex) and this object will serve to uniquely identify
    // each IEC. The type is interface{} with range: 1..1024.
    CCallHistoryIecIndex interface{}

    // This object reflects the Internal Error Code. The format is a string of
    // dotted decimal numbers composed of the following components: 
    // Version.Entity.Category.Subsystem.Errorcode.Diagnostic  Each component is
    // defined as follows: Version     : The version of IEC software. Entity     
    // : The network entity that originated               the error. Category    :
    // The category of the error (eg, software               error, hardware
    // resource unavailable, ...) Subsystem   : The subsystem in which the error
    // occurred. Errorcode   : A subsytem-specific error code. Diagnostic  : An
    // implementation-specific diagnostic code. The type is string.
    CCallHistoryIec interface{}
}

func (cCallHistoryIecEntry *CISCODIALCONTROLMIB_CCallHistoryIecTable_CCallHistoryIecEntry) GetEntityData() *types.CommonEntityData {
    cCallHistoryIecEntry.EntityData.YFilter = cCallHistoryIecEntry.YFilter
    cCallHistoryIecEntry.EntityData.YangName = "cCallHistoryIecEntry"
    cCallHistoryIecEntry.EntityData.BundleName = "cisco_ios_xe"
    cCallHistoryIecEntry.EntityData.ParentYangName = "cCallHistoryIecTable"
    cCallHistoryIecEntry.EntityData.SegmentPath = "cCallHistoryIecEntry" + types.AddKeyToken(cCallHistoryIecEntry.CCallHistoryIndex, "cCallHistoryIndex") + types.AddKeyToken(cCallHistoryIecEntry.CCallHistoryIecIndex, "cCallHistoryIecIndex")
    cCallHistoryIecEntry.EntityData.AbsolutePath = "CISCO-DIAL-CONTROL-MIB:CISCO-DIAL-CONTROL-MIB/cCallHistoryIecTable/" + cCallHistoryIecEntry.EntityData.SegmentPath
    cCallHistoryIecEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cCallHistoryIecEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cCallHistoryIecEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cCallHistoryIecEntry.EntityData.Children = types.NewOrderedMap()
    cCallHistoryIecEntry.EntityData.Leafs = types.NewOrderedMap()
    cCallHistoryIecEntry.EntityData.Leafs.Append("cCallHistoryIndex", types.YLeaf{"CCallHistoryIndex", cCallHistoryIecEntry.CCallHistoryIndex})
    cCallHistoryIecEntry.EntityData.Leafs.Append("cCallHistoryIecIndex", types.YLeaf{"CCallHistoryIecIndex", cCallHistoryIecEntry.CCallHistoryIecIndex})
    cCallHistoryIecEntry.EntityData.Leafs.Append("cCallHistoryIec", types.YLeaf{"CCallHistoryIec", cCallHistoryIecEntry.CCallHistoryIec})

    cCallHistoryIecEntry.EntityData.YListKeys = []string {"CCallHistoryIndex", "CCallHistoryIecIndex"}

    return &(cCallHistoryIecEntry.EntityData)
}

