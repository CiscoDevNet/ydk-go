// The MIB module to describe peer information for
// demand access and possibly other kinds of interfaces.
package dial_control_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dial_control_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DIAL-CONTROL-MIB DIAL-CONTROL-MIB}", reflect.TypeOf(DIALCONTROLMIB{}))
    ydk.RegisterEntity("DIAL-CONTROL-MIB:DIAL-CONTROL-MIB", reflect.TypeOf(DIALCONTROLMIB{}))
}

// DIALCONTROLMIB
type DIALCONTROLMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DialCtlConfiguration DIALCONTROLMIB_DialCtlConfiguration

    
    CallHistory DIALCONTROLMIB_CallHistory

    // The list of peers from which the managed device will accept calls or to
    // which it will place them.
    DialCtlPeerCfgTable DIALCONTROLMIB_DialCtlPeerCfgTable

    // A table containing information about active calls to a specific
    // destination.
    CallActiveTable DIALCONTROLMIB_CallActiveTable

    // A table containing information about specific calls to a specific
    // destination.
    CallHistoryTable DIALCONTROLMIB_CallHistoryTable
}

func (dIALCONTROLMIB *DIALCONTROLMIB) GetEntityData() *types.CommonEntityData {
    dIALCONTROLMIB.EntityData.YFilter = dIALCONTROLMIB.YFilter
    dIALCONTROLMIB.EntityData.YangName = "DIAL-CONTROL-MIB"
    dIALCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    dIALCONTROLMIB.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    dIALCONTROLMIB.EntityData.SegmentPath = "DIAL-CONTROL-MIB:DIAL-CONTROL-MIB"
    dIALCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dIALCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dIALCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dIALCONTROLMIB.EntityData.Children = types.NewOrderedMap()
    dIALCONTROLMIB.EntityData.Children.Append("dialCtlConfiguration", types.YChild{"DialCtlConfiguration", &dIALCONTROLMIB.DialCtlConfiguration})
    dIALCONTROLMIB.EntityData.Children.Append("callHistory", types.YChild{"CallHistory", &dIALCONTROLMIB.CallHistory})
    dIALCONTROLMIB.EntityData.Children.Append("dialCtlPeerCfgTable", types.YChild{"DialCtlPeerCfgTable", &dIALCONTROLMIB.DialCtlPeerCfgTable})
    dIALCONTROLMIB.EntityData.Children.Append("callActiveTable", types.YChild{"CallActiveTable", &dIALCONTROLMIB.CallActiveTable})
    dIALCONTROLMIB.EntityData.Children.Append("callHistoryTable", types.YChild{"CallHistoryTable", &dIALCONTROLMIB.CallHistoryTable})
    dIALCONTROLMIB.EntityData.Leafs = types.NewOrderedMap()

    dIALCONTROLMIB.EntityData.YListKeys = []string {}

    return &(dIALCONTROLMIB.EntityData)
}

// DIALCONTROLMIB_DialCtlConfiguration
type DIALCONTROLMIB_DialCtlConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The security level for acceptance of incoming calls. acceptNone(1)  -
    // incoming calls will not be accepted acceptAll(2)   - incoming calls will be
    // accepted,                  even if there is no matching entry              
    // in the dialCtlPeerCfgTable acceptKnown(3) - incoming calls will be accepted
    // only                  if there is a matching entry in the                 
    // dialCtlPeerCfgTable. The type is DialCtlAcceptMode.
    DialCtlAcceptMode interface{}

    // This object indicates whether dialCtlPeerCallInformation and
    // dialCtlPeerCallSetup traps should be generated for all peers. If the value
    // of this object is enabled(1), traps will be generated for all peers. If the
    // value of this object is disabled(2), traps will be generated only for peers
    // having dialCtlPeerCfgTrapEnable set to enabled(1). The type is
    // DialCtlTrapEnable.
    DialCtlTrapEnable interface{}
}

func (dialCtlConfiguration *DIALCONTROLMIB_DialCtlConfiguration) GetEntityData() *types.CommonEntityData {
    dialCtlConfiguration.EntityData.YFilter = dialCtlConfiguration.YFilter
    dialCtlConfiguration.EntityData.YangName = "dialCtlConfiguration"
    dialCtlConfiguration.EntityData.BundleName = "cisco_ios_xe"
    dialCtlConfiguration.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    dialCtlConfiguration.EntityData.SegmentPath = "dialCtlConfiguration"
    dialCtlConfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dialCtlConfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dialCtlConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dialCtlConfiguration.EntityData.Children = types.NewOrderedMap()
    dialCtlConfiguration.EntityData.Leafs = types.NewOrderedMap()
    dialCtlConfiguration.EntityData.Leafs.Append("dialCtlAcceptMode", types.YLeaf{"DialCtlAcceptMode", dialCtlConfiguration.DialCtlAcceptMode})
    dialCtlConfiguration.EntityData.Leafs.Append("dialCtlTrapEnable", types.YLeaf{"DialCtlTrapEnable", dialCtlConfiguration.DialCtlTrapEnable})

    dialCtlConfiguration.EntityData.YListKeys = []string {}

    return &(dialCtlConfiguration.EntityData)
}

// DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode represents                  dialCtlPeerCfgTable
type DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode string

const (
    DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode_acceptNone DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode = "acceptNone"

    DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode_acceptAll DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode = "acceptAll"

    DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode_acceptKnown DIALCONTROLMIB_DialCtlConfiguration_DialCtlAcceptMode = "acceptKnown"
)

// DIALCONTROLMIB_DialCtlConfiguration_DialCtlTrapEnable represents to enabled(1).
type DIALCONTROLMIB_DialCtlConfiguration_DialCtlTrapEnable string

const (
    DIALCONTROLMIB_DialCtlConfiguration_DialCtlTrapEnable_enabled DIALCONTROLMIB_DialCtlConfiguration_DialCtlTrapEnable = "enabled"

    DIALCONTROLMIB_DialCtlConfiguration_DialCtlTrapEnable_disabled DIALCONTROLMIB_DialCtlConfiguration_DialCtlTrapEnable = "disabled"
)

// DIALCONTROLMIB_CallHistory
type DIALCONTROLMIB_CallHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The upper limit on the number of entries that the callHistoryTable may
    // contain.  A value of 0 will prevent any history from being retained. When
    // this table is full, the oldest entry will be deleted and the new one will
    // be created. The type is interface{} with range: 0..2147483647.
    CallHistoryTableMaxLength interface{}

    // The minimum amount of time that an callHistoryEntry will be maintained
    // before being deleted. A value of 0 will prevent any history from being
    // retained in the callHistoryTable, but will neither prevent callCompletion
    // traps being generated nor affect other tables. The type is interface{} with
    // range: 0..2147483647. Units are minutes.
    CallHistoryRetainTimer interface{}
}

func (callHistory *DIALCONTROLMIB_CallHistory) GetEntityData() *types.CommonEntityData {
    callHistory.EntityData.YFilter = callHistory.YFilter
    callHistory.EntityData.YangName = "callHistory"
    callHistory.EntityData.BundleName = "cisco_ios_xe"
    callHistory.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    callHistory.EntityData.SegmentPath = "callHistory"
    callHistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callHistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callHistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callHistory.EntityData.Children = types.NewOrderedMap()
    callHistory.EntityData.Leafs = types.NewOrderedMap()
    callHistory.EntityData.Leafs.Append("callHistoryTableMaxLength", types.YLeaf{"CallHistoryTableMaxLength", callHistory.CallHistoryTableMaxLength})
    callHistory.EntityData.Leafs.Append("callHistoryRetainTimer", types.YLeaf{"CallHistoryRetainTimer", callHistory.CallHistoryRetainTimer})

    callHistory.EntityData.YListKeys = []string {}

    return &(callHistory.EntityData)
}

// DIALCONTROLMIB_DialCtlPeerCfgTable
// The list of peers from which the managed device
// will accept calls or to which it will place them.
type DIALCONTROLMIB_DialCtlPeerCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for a single Peer. This entry is effectively permanent,
    // and contains information to identify the peer, how to connect to the peer,
    // how to identify the peer and its permissions. The value of
    // dialCtlPeerCfgOriginateAddress must be specified before a new row in this
    // table can become active(1). Any writeable parameters in an existing entry
    // can be modified while the entry is active. The modification will take
    // effect when the peer in question will be called the next time. An entry in
    // this table can only be created if the associated ifEntry already exists.
    // The type is slice of
    // DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry.
    DialCtlPeerCfgEntry []*DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry
}

func (dialCtlPeerCfgTable *DIALCONTROLMIB_DialCtlPeerCfgTable) GetEntityData() *types.CommonEntityData {
    dialCtlPeerCfgTable.EntityData.YFilter = dialCtlPeerCfgTable.YFilter
    dialCtlPeerCfgTable.EntityData.YangName = "dialCtlPeerCfgTable"
    dialCtlPeerCfgTable.EntityData.BundleName = "cisco_ios_xe"
    dialCtlPeerCfgTable.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    dialCtlPeerCfgTable.EntityData.SegmentPath = "dialCtlPeerCfgTable"
    dialCtlPeerCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dialCtlPeerCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dialCtlPeerCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dialCtlPeerCfgTable.EntityData.Children = types.NewOrderedMap()
    dialCtlPeerCfgTable.EntityData.Children.Append("dialCtlPeerCfgEntry", types.YChild{"DialCtlPeerCfgEntry", nil})
    for i := range dialCtlPeerCfgTable.DialCtlPeerCfgEntry {
        dialCtlPeerCfgTable.EntityData.Children.Append(types.GetSegmentPath(dialCtlPeerCfgTable.DialCtlPeerCfgEntry[i]), types.YChild{"DialCtlPeerCfgEntry", dialCtlPeerCfgTable.DialCtlPeerCfgEntry[i]})
    }
    dialCtlPeerCfgTable.EntityData.Leafs = types.NewOrderedMap()

    dialCtlPeerCfgTable.EntityData.YListKeys = []string {}

    return &(dialCtlPeerCfgTable.EntityData)
}

// DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry
// Configuration data for a single Peer. This entry is
// effectively permanent, and contains information
// to identify the peer, how to connect to the peer,
// how to identify the peer and its permissions.
// The value of dialCtlPeerCfgOriginateAddress must be
// specified before a new row in this table can become
// active(1). Any writeable parameters in an existing entry
// can be modified while the entry is active. The modification
// will take effect when the peer in question will be
// called the next time.
// An entry in this table can only be created if the
// associated ifEntry already exists.
type DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object identifies a single peer. There may be
    // several entries in this table for one peer, defining different ways of
    // reaching this peer. Thus, there may be several entries in this table with
    // the same value of dialCtlPeerCfgId. Multiple entries for one peer may be
    // used to support multilink as well as backup lines. A single peer will be
    // identified by a unique value of this object. Several entries for one peer
    // MUST have the same value of dialCtlPeerCfgId, but different ifEntries and
    // thus different values of ifIndex. The type is interface{} with range:
    // 1..2147483647.
    DialCtlPeerCfgId interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The interface type to be used for calling this peer. In case of ISDN, the
    // value of isdn(63) is to be used. The type is IANAifType.
    DialCtlPeerCfgIfType interface{}

    // ifIndex value of an interface the peer will have to be called on. For
    // example, on an ISDN interface, this can be the ifIndex value of a D channel
    // or the ifIndex value of a B channel, whatever is appropriate for a given
    // peer. As an example, for Basic Rate leased lines it will be necessary to
    // specify a B channel ifIndex, while for     semi-permanent connections the D
    // channel ifIndex has to be specified. If the interface can be dynamically
    // assigned, this object has a value of zero. The type is interface{} with
    // range: 0..2147483647.
    DialCtlPeerCfgLowerIf interface{}

    // Call Address at which the peer will be called. Think of this as the set of
    // characters following 'ATDT ' or the 'phone number' included in a D channel
    // call request.  The structure of this information will be switch type
    // specific. If there is no address information required for reaching the
    // peer, i.e., for leased lines, this object will be a zero length string. The
    // type is string.
    DialCtlPeerCfgOriginateAddress interface{}

    // Calling Party Number information element, as for example passed in an ISDN
    // SETUP message by a PBX or switch, for incoming calls. This address can be
    // used to identify the peer. If this address is either unknown or identical
    // to dialCtlPeerCfgOriginateAddress, this object will be a zero length
    // string. The type is string.
    DialCtlPeerCfgAnswerAddress interface{}

    // Subaddress at which the peer will be called. If the subaddress is undefined
    // for the given media or unused, this is a zero length string. The type is
    // string.
    DialCtlPeerCfgSubAddress interface{}

    // Closed User Group at which the peer will be called. If the Closed User
    // Group is undefined for the given media or unused, this is a zero length
    // string. The type is string.
    DialCtlPeerCfgClosedUserGroup interface{}

    // The desired information transfer speed in bits/second when calling this
    // peer. The detailed media specific information, e.g. information type and
    // information transfer rate for ISDN circuits, has to be extracted from this
    // object. If the transfer speed to be used is unknown or the default speed
    // for this type of interfaces, the value of this object may be zero. The type
    // is interface{} with range: 0..2147483647.
    DialCtlPeerCfgSpeed interface{}

    // The Information Transfer Capability to be used when calling this peer. 
    // speech(2) refers to a non-data connection, whereas audio31(6) and audio7(7)
    // refer to data mode connections. The type is DialCtlPeerCfgInfoType.
    DialCtlPeerCfgInfoType interface{}

    // Applicable permissions. callback(4) either rejects the call and then calls
    // back, or uses the 'Reverse charging' information element if it is
    // available. Note that callback(4) is supposed to control charging, not
    // security, and applies to callback prior to accepting a call. Callback for
    // security reasons can be handled using PPP callback. The type is
    // DialCtlPeerCfgPermission.
    DialCtlPeerCfgPermission interface{}

    // The connection will be automatically disconnected if no longer carrying
    // useful data for a time period, in seconds, specified in this object. Useful
    // data in this context refers to forwarding packets, including routing
    // information; it excludes the encapsulator maintenance frames. A value of
    // zero means the connection will not be automatically taken down due to
    // inactivity, which implies that it is a dedicated circuit. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    DialCtlPeerCfgInactivityTimer interface{}

    // Minimum duration of a call in seconds, starting from the time the call is
    // connected until the call is disconnected. This is to accomplish the fact
    // that in most countries charging applies to units of time, which should be
    // matched as closely as possible. The type is interface{} with range:
    // 0..2147483647.
    DialCtlPeerCfgMinDuration interface{}

    // Maximum call duration in seconds. Zero means 'unlimited'. The type is
    // interface{} with range: 0..2147483647.
    DialCtlPeerCfgMaxDuration interface{}

    // The call timeout time in seconds. The default value of zero means that the
    // call timeout as specified for the media in question will apply. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    DialCtlPeerCfgCarrierDelay interface{}

    // The number of calls to a non-responding address that may be made. A retry
    // count of zero means there is no bound. The intent is to bound the number of
    // successive calls to an address which is inaccessible, or which refuses
    // those calls.  Some countries regulate the number of call retries to a given
    // peer that can be made. The type is interface{} with range: 0..2147483647.
    DialCtlPeerCfgCallRetries interface{}

    // The time in seconds between call retries if a peer cannot be reached. A
    // value of zero means that call retries may be done without any delay. The
    // type is interface{} with range: 0..2147483647. Units are seconds.
    DialCtlPeerCfgRetryDelay interface{}

    // The time in seconds after which call attempts are to be placed again after
    // a peer has been noticed to be unreachable, i.e. after
    // dialCtlPeerCfgCallRetries unsuccessful call attempts. A value of zero means
    // that a peer will not be called again after dialCtlPeerCfgCallRetries
    // unsuccessful call attempts. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    DialCtlPeerCfgFailureDelay interface{}

    // This object indicates whether dialCtlPeerCallInformation and
    // dialCtlPeerCallSetup traps should be generated for this peer. The type is
    // DialCtlPeerCfgTrapEnable.
    DialCtlPeerCfgTrapEnable interface{}

    // Status of one row in this table. The type is RowStatus.
    DialCtlPeerCfgStatus interface{}

    // Accumulated connect time to the peer since system startup. This is the
    // total connect time, i.e. the connect time for outgoing calls plus the time
    // for incoming calls. The type is interface{} with range: 0..4294967295.
    // Units are seconds.
    DialCtlPeerStatsConnectTime interface{}

    // The total number of charging units applying to this peer since system
    // startup. Only the charging units applying to the local interface, i.e. for
    // originated calls or for calls with 'Reverse charging' being active, will be
    // counted here. The type is interface{} with range: 0..4294967295.
    DialCtlPeerStatsChargedUnits interface{}

    // Number of completed calls to this peer. The type is interface{} with range:
    // 0..4294967295.
    DialCtlPeerStatsSuccessCalls interface{}

    // Number of failed call attempts to this peer since system startup. The type
    // is interface{} with range: 0..4294967295.
    DialCtlPeerStatsFailCalls interface{}

    // Number of calls from this peer accepted since system startup. The type is
    // interface{} with range: 0..4294967295.
    DialCtlPeerStatsAcceptCalls interface{}

    // Number of calls from this peer refused since system startup. The type is
    // interface{} with range: 0..4294967295.
    DialCtlPeerStatsRefuseCalls interface{}

    // The encoded network cause value associated with the last call. This object
    // will be updated whenever a call is started or cleared. The value of this
    // object will depend on the interface type as well as on the protocol and
    // protocol version being used on this interface. Some references for possible
    // cause values are given below. The type is string with length: 0..4.
    DialCtlPeerStatsLastDisconnectCause interface{}

    // ASCII text describing the reason for the last call termination.  This
    // object exists because it would be impossible for a management station to
    // store all possible cause values for all types of interfaces. It should be
    // used only if a management station is unable to decode the value of
    // dialCtlPeerStatsLastDisconnectCause.  This object will be updated whenever
    // a call is started or cleared. The type is string.
    DialCtlPeerStatsLastDisconnectText interface{}

    // The value of sysUpTime when the last call to this peer was started. For
    // ISDN media, this will be the time when the setup message was received from
    // or sent to the network. This object will be updated whenever a call is
    // started or cleared. The type is interface{} with range: 0..4294967295.
    DialCtlPeerStatsLastSetupTime interface{}
}

func (dialCtlPeerCfgEntry *DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry) GetEntityData() *types.CommonEntityData {
    dialCtlPeerCfgEntry.EntityData.YFilter = dialCtlPeerCfgEntry.YFilter
    dialCtlPeerCfgEntry.EntityData.YangName = "dialCtlPeerCfgEntry"
    dialCtlPeerCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    dialCtlPeerCfgEntry.EntityData.ParentYangName = "dialCtlPeerCfgTable"
    dialCtlPeerCfgEntry.EntityData.SegmentPath = "dialCtlPeerCfgEntry" + types.AddKeyToken(dialCtlPeerCfgEntry.DialCtlPeerCfgId, "dialCtlPeerCfgId") + types.AddKeyToken(dialCtlPeerCfgEntry.IfIndex, "ifIndex")
    dialCtlPeerCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dialCtlPeerCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dialCtlPeerCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dialCtlPeerCfgEntry.EntityData.Children = types.NewOrderedMap()
    dialCtlPeerCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgId", types.YLeaf{"DialCtlPeerCfgId", dialCtlPeerCfgEntry.DialCtlPeerCfgId})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", dialCtlPeerCfgEntry.IfIndex})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgIfType", types.YLeaf{"DialCtlPeerCfgIfType", dialCtlPeerCfgEntry.DialCtlPeerCfgIfType})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgLowerIf", types.YLeaf{"DialCtlPeerCfgLowerIf", dialCtlPeerCfgEntry.DialCtlPeerCfgLowerIf})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgOriginateAddress", types.YLeaf{"DialCtlPeerCfgOriginateAddress", dialCtlPeerCfgEntry.DialCtlPeerCfgOriginateAddress})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgAnswerAddress", types.YLeaf{"DialCtlPeerCfgAnswerAddress", dialCtlPeerCfgEntry.DialCtlPeerCfgAnswerAddress})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgSubAddress", types.YLeaf{"DialCtlPeerCfgSubAddress", dialCtlPeerCfgEntry.DialCtlPeerCfgSubAddress})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgClosedUserGroup", types.YLeaf{"DialCtlPeerCfgClosedUserGroup", dialCtlPeerCfgEntry.DialCtlPeerCfgClosedUserGroup})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgSpeed", types.YLeaf{"DialCtlPeerCfgSpeed", dialCtlPeerCfgEntry.DialCtlPeerCfgSpeed})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgInfoType", types.YLeaf{"DialCtlPeerCfgInfoType", dialCtlPeerCfgEntry.DialCtlPeerCfgInfoType})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgPermission", types.YLeaf{"DialCtlPeerCfgPermission", dialCtlPeerCfgEntry.DialCtlPeerCfgPermission})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgInactivityTimer", types.YLeaf{"DialCtlPeerCfgInactivityTimer", dialCtlPeerCfgEntry.DialCtlPeerCfgInactivityTimer})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgMinDuration", types.YLeaf{"DialCtlPeerCfgMinDuration", dialCtlPeerCfgEntry.DialCtlPeerCfgMinDuration})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgMaxDuration", types.YLeaf{"DialCtlPeerCfgMaxDuration", dialCtlPeerCfgEntry.DialCtlPeerCfgMaxDuration})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgCarrierDelay", types.YLeaf{"DialCtlPeerCfgCarrierDelay", dialCtlPeerCfgEntry.DialCtlPeerCfgCarrierDelay})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgCallRetries", types.YLeaf{"DialCtlPeerCfgCallRetries", dialCtlPeerCfgEntry.DialCtlPeerCfgCallRetries})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgRetryDelay", types.YLeaf{"DialCtlPeerCfgRetryDelay", dialCtlPeerCfgEntry.DialCtlPeerCfgRetryDelay})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgFailureDelay", types.YLeaf{"DialCtlPeerCfgFailureDelay", dialCtlPeerCfgEntry.DialCtlPeerCfgFailureDelay})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgTrapEnable", types.YLeaf{"DialCtlPeerCfgTrapEnable", dialCtlPeerCfgEntry.DialCtlPeerCfgTrapEnable})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerCfgStatus", types.YLeaf{"DialCtlPeerCfgStatus", dialCtlPeerCfgEntry.DialCtlPeerCfgStatus})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsConnectTime", types.YLeaf{"DialCtlPeerStatsConnectTime", dialCtlPeerCfgEntry.DialCtlPeerStatsConnectTime})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsChargedUnits", types.YLeaf{"DialCtlPeerStatsChargedUnits", dialCtlPeerCfgEntry.DialCtlPeerStatsChargedUnits})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsSuccessCalls", types.YLeaf{"DialCtlPeerStatsSuccessCalls", dialCtlPeerCfgEntry.DialCtlPeerStatsSuccessCalls})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsFailCalls", types.YLeaf{"DialCtlPeerStatsFailCalls", dialCtlPeerCfgEntry.DialCtlPeerStatsFailCalls})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsAcceptCalls", types.YLeaf{"DialCtlPeerStatsAcceptCalls", dialCtlPeerCfgEntry.DialCtlPeerStatsAcceptCalls})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsRefuseCalls", types.YLeaf{"DialCtlPeerStatsRefuseCalls", dialCtlPeerCfgEntry.DialCtlPeerStatsRefuseCalls})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsLastDisconnectCause", types.YLeaf{"DialCtlPeerStatsLastDisconnectCause", dialCtlPeerCfgEntry.DialCtlPeerStatsLastDisconnectCause})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsLastDisconnectText", types.YLeaf{"DialCtlPeerStatsLastDisconnectText", dialCtlPeerCfgEntry.DialCtlPeerStatsLastDisconnectText})
    dialCtlPeerCfgEntry.EntityData.Leafs.Append("dialCtlPeerStatsLastSetupTime", types.YLeaf{"DialCtlPeerStatsLastSetupTime", dialCtlPeerCfgEntry.DialCtlPeerStatsLastSetupTime})

    dialCtlPeerCfgEntry.EntityData.YListKeys = []string {"DialCtlPeerCfgId", "IfIndex"}

    return &(dialCtlPeerCfgEntry.EntityData)
}

// DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType represents connections.
type DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType string

const (
    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_other DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "other"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_speech DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "speech"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_unrestrictedDigital DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "unrestrictedDigital"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_unrestrictedDigital56 DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "unrestrictedDigital56"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_restrictedDigital DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "restrictedDigital"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_audio31 DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "audio31"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_audio7 DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "audio7"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_video DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "video"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_packetSwitched DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "packetSwitched"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType_fax DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgInfoType = "fax"
)

// DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission represents PPP callback.
type DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission string

const (
    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission_originate DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission = "originate"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission_answer DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission = "answer"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission_both DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission = "both"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission_callback DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission = "callback"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission_none DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgPermission = "none"
)

// DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgTrapEnable represents this peer.
type DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgTrapEnable string

const (
    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgTrapEnable_enabled DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgTrapEnable = "enabled"

    DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgTrapEnable_disabled DIALCONTROLMIB_DialCtlPeerCfgTable_DialCtlPeerCfgEntry_DialCtlPeerCfgTrapEnable = "disabled"
)

// DIALCONTROLMIB_CallActiveTable
// A table containing information about active
// calls to a specific destination.
type DIALCONTROLMIB_CallActiveTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single active Connection. An entry in this
    // table will be created when a call is started. An entry in this table will
    // be deleted when an active call clears. The type is slice of
    // DIALCONTROLMIB_CallActiveTable_CallActiveEntry.
    CallActiveEntry []*DIALCONTROLMIB_CallActiveTable_CallActiveEntry
}

func (callActiveTable *DIALCONTROLMIB_CallActiveTable) GetEntityData() *types.CommonEntityData {
    callActiveTable.EntityData.YFilter = callActiveTable.YFilter
    callActiveTable.EntityData.YangName = "callActiveTable"
    callActiveTable.EntityData.BundleName = "cisco_ios_xe"
    callActiveTable.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    callActiveTable.EntityData.SegmentPath = "callActiveTable"
    callActiveTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callActiveTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callActiveTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callActiveTable.EntityData.Children = types.NewOrderedMap()
    callActiveTable.EntityData.Children.Append("callActiveEntry", types.YChild{"CallActiveEntry", nil})
    for i := range callActiveTable.CallActiveEntry {
        callActiveTable.EntityData.Children.Append(types.GetSegmentPath(callActiveTable.CallActiveEntry[i]), types.YChild{"CallActiveEntry", callActiveTable.CallActiveEntry[i]})
    }
    callActiveTable.EntityData.Leafs = types.NewOrderedMap()

    callActiveTable.EntityData.YListKeys = []string {}

    return &(callActiveTable.EntityData)
}

// DIALCONTROLMIB_CallActiveTable_CallActiveEntry
// The information regarding a single active Connection.
// An entry in this table will be created when a call is
// started. An entry in this table will be deleted when
// an active call clears.
type DIALCONTROLMIB_CallActiveTable_CallActiveEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of sysUpTime when the call associated to
    // this entry was started. This will be useful for an NMS to retrieve all
    // calls after a specific time. Also, this object can be useful in finding
    // large delays between the time the call was started and the time the call
    // was connected. For ISDN media, this will be the time when the setup message
    // was received from or sent to the network. The type is interface{} with
    // range: 0..4294967295.
    CallActiveSetupTime interface{}

    // This attribute is a key. Small index variable to distinguish calls that
    // start in the same hundredth of a second. The type is interface{} with
    // range: 1..2147483647.
    CallActiveIndex interface{}

    // The number this call is connected to. If the number is not available, then
    // it will have a length of zero. The type is string.
    CallActivePeerAddress interface{}

    // The subaddress this call is connected to. If the subaddress is undefined or
    // not available, this will be a zero length string. The type is string.
    CallActivePeerSubAddress interface{}

    // This is the Id value of the peer table entry to which this call was made.
    // If a peer table entry for this call does not exist or is unknown, the value
    // of this object will be zero. The type is interface{} with range:
    // 0..2147483647.
    CallActivePeerId interface{}

    // This is the ifIndex value of the peer table entry to which this call was
    // made. If a peer table entry for this call does not exist or is unknown, the
    // value of this object will be zero. The type is interface{} with range:
    // 0..2147483647.
    CallActivePeerIfIndex interface{}

    // This is the ifIndex value of the logical interface through which this call
    // was made. For ISDN media, this would be the ifIndex of the B channel which
    // was used for this call. If the ifIndex value is unknown, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    CallActiveLogicalIfIndex interface{}

    // The value of sysUpTime when the call was connected. If the call is not
    // connected, this object will have a value of zero. The type is interface{}
    // with range: 0..4294967295.
    CallActiveConnectTime interface{}

    // The current call state. unknown(1)     - The call state is unknown.
    // connecting(2)  - A connection attempt (outgoing call)                  is
    // being made. connected(3)   - An incoming call is in the process            
    // of validation. active(4)      - The call is active. The type is
    // CallActiveCallState.
    CallActiveCallState interface{}

    // The call origin. The type is CallActiveCallOrigin.
    CallActiveCallOrigin interface{}

    // The number of charged units for this connection. For incoming calls or if
    // charging information is not supplied by the switch, the value of this
    // object will be zero. The type is interface{} with range: 0..4294967295.
    CallActiveChargedUnits interface{}

    // The information type for this call. The type is CallActiveInfoType.
    CallActiveInfoType interface{}

    // The number of packets which were transmitted for this call. The type is
    // interface{} with range: 0..4294967295.
    CallActiveTransmitPackets interface{}

    // The number of bytes which were transmitted for this call. The type is
    // interface{} with range: 0..4294967295.
    CallActiveTransmitBytes interface{}

    // The number of packets which were received for this call. The type is
    // interface{} with range: 0..4294967295.
    CallActiveReceivePackets interface{}

    // The number of bytes which were received for this call. The type is
    // interface{} with range: 0..4294967295.
    CallActiveReceiveBytes interface{}
}

func (callActiveEntry *DIALCONTROLMIB_CallActiveTable_CallActiveEntry) GetEntityData() *types.CommonEntityData {
    callActiveEntry.EntityData.YFilter = callActiveEntry.YFilter
    callActiveEntry.EntityData.YangName = "callActiveEntry"
    callActiveEntry.EntityData.BundleName = "cisco_ios_xe"
    callActiveEntry.EntityData.ParentYangName = "callActiveTable"
    callActiveEntry.EntityData.SegmentPath = "callActiveEntry" + types.AddKeyToken(callActiveEntry.CallActiveSetupTime, "callActiveSetupTime") + types.AddKeyToken(callActiveEntry.CallActiveIndex, "callActiveIndex")
    callActiveEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callActiveEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callActiveEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callActiveEntry.EntityData.Children = types.NewOrderedMap()
    callActiveEntry.EntityData.Leafs = types.NewOrderedMap()
    callActiveEntry.EntityData.Leafs.Append("callActiveSetupTime", types.YLeaf{"CallActiveSetupTime", callActiveEntry.CallActiveSetupTime})
    callActiveEntry.EntityData.Leafs.Append("callActiveIndex", types.YLeaf{"CallActiveIndex", callActiveEntry.CallActiveIndex})
    callActiveEntry.EntityData.Leafs.Append("callActivePeerAddress", types.YLeaf{"CallActivePeerAddress", callActiveEntry.CallActivePeerAddress})
    callActiveEntry.EntityData.Leafs.Append("callActivePeerSubAddress", types.YLeaf{"CallActivePeerSubAddress", callActiveEntry.CallActivePeerSubAddress})
    callActiveEntry.EntityData.Leafs.Append("callActivePeerId", types.YLeaf{"CallActivePeerId", callActiveEntry.CallActivePeerId})
    callActiveEntry.EntityData.Leafs.Append("callActivePeerIfIndex", types.YLeaf{"CallActivePeerIfIndex", callActiveEntry.CallActivePeerIfIndex})
    callActiveEntry.EntityData.Leafs.Append("callActiveLogicalIfIndex", types.YLeaf{"CallActiveLogicalIfIndex", callActiveEntry.CallActiveLogicalIfIndex})
    callActiveEntry.EntityData.Leafs.Append("callActiveConnectTime", types.YLeaf{"CallActiveConnectTime", callActiveEntry.CallActiveConnectTime})
    callActiveEntry.EntityData.Leafs.Append("callActiveCallState", types.YLeaf{"CallActiveCallState", callActiveEntry.CallActiveCallState})
    callActiveEntry.EntityData.Leafs.Append("callActiveCallOrigin", types.YLeaf{"CallActiveCallOrigin", callActiveEntry.CallActiveCallOrigin})
    callActiveEntry.EntityData.Leafs.Append("callActiveChargedUnits", types.YLeaf{"CallActiveChargedUnits", callActiveEntry.CallActiveChargedUnits})
    callActiveEntry.EntityData.Leafs.Append("callActiveInfoType", types.YLeaf{"CallActiveInfoType", callActiveEntry.CallActiveInfoType})
    callActiveEntry.EntityData.Leafs.Append("callActiveTransmitPackets", types.YLeaf{"CallActiveTransmitPackets", callActiveEntry.CallActiveTransmitPackets})
    callActiveEntry.EntityData.Leafs.Append("callActiveTransmitBytes", types.YLeaf{"CallActiveTransmitBytes", callActiveEntry.CallActiveTransmitBytes})
    callActiveEntry.EntityData.Leafs.Append("callActiveReceivePackets", types.YLeaf{"CallActiveReceivePackets", callActiveEntry.CallActiveReceivePackets})
    callActiveEntry.EntityData.Leafs.Append("callActiveReceiveBytes", types.YLeaf{"CallActiveReceiveBytes", callActiveEntry.CallActiveReceiveBytes})

    callActiveEntry.EntityData.YListKeys = []string {"CallActiveSetupTime", "CallActiveIndex"}

    return &(callActiveEntry.EntityData)
}

// DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin represents The call origin.
type DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin string

const (
    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin_originate DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin = "originate"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin_answer DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin = "answer"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin_callback DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallOrigin = "callback"
)

// DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState represents active(4)      - The call is active.
type DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState string

const (
    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState_unknown DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState = "unknown"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState_connecting DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState = "connecting"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState_connected DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState = "connected"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState_active DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveCallState = "active"
)

// DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType represents The information type for this call.
type DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType string

const (
    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_other DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "other"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_speech DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "speech"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_unrestrictedDigital DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "unrestrictedDigital"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_unrestrictedDigital56 DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "unrestrictedDigital56"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_restrictedDigital DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "restrictedDigital"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_audio31 DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "audio31"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_audio7 DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "audio7"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_video DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "video"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_packetSwitched DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "packetSwitched"

    DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType_fax DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveInfoType = "fax"
)

// DIALCONTROLMIB_CallHistoryTable
// A table containing information about specific
// calls to a specific destination.
type DIALCONTROLMIB_CallHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single Connection. The type is slice of
    // DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry.
    CallHistoryEntry []*DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry
}

func (callHistoryTable *DIALCONTROLMIB_CallHistoryTable) GetEntityData() *types.CommonEntityData {
    callHistoryTable.EntityData.YFilter = callHistoryTable.YFilter
    callHistoryTable.EntityData.YangName = "callHistoryTable"
    callHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    callHistoryTable.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    callHistoryTable.EntityData.SegmentPath = "callHistoryTable"
    callHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callHistoryTable.EntityData.Children = types.NewOrderedMap()
    callHistoryTable.EntityData.Children.Append("callHistoryEntry", types.YChild{"CallHistoryEntry", nil})
    for i := range callHistoryTable.CallHistoryEntry {
        callHistoryTable.EntityData.Children.Append(types.GetSegmentPath(callHistoryTable.CallHistoryEntry[i]), types.YChild{"CallHistoryEntry", callHistoryTable.CallHistoryEntry[i]})
    }
    callHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    callHistoryTable.EntityData.YListKeys = []string {}

    return &(callHistoryTable.EntityData)
}

// DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry
// The information regarding a single Connection.
type DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveSetupTime
    CallActiveSetupTime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveIndex
    CallActiveIndex interface{}

    // The number this call was connected to. If the number is not available, then
    // it will have a length of zero. The type is string.
    CallHistoryPeerAddress interface{}

    // The subaddress this call was connected to. If the subaddress is undefined
    // or not available, this will be a zero length string. The type is string.
    CallHistoryPeerSubAddress interface{}

    // This is the Id value of the peer table entry to which this call was made.
    // If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    CallHistoryPeerId interface{}

    // This is the ifIndex value of the peer table entry to which this call was
    // made. If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    CallHistoryPeerIfIndex interface{}

    // This is the ifIndex value of the logical interface through which this call
    // was made. For ISDN media, this would be the ifIndex of the B channel which
    // was used for this call. The type is interface{} with range: 1..2147483647.
    CallHistoryLogicalIfIndex interface{}

    // The encoded network cause value associated with this call.  The value of
    // this object will depend on the interface type as well as on the protocol
    // and protocol version being used on this interface. Some references for
    // possible cause values are given below. The type is string with length:
    // 0..4.
    CallHistoryDisconnectCause interface{}

    // ASCII text describing the reason for call termination.  This object exists
    // because it would be impossible for a management station to store all
    // possible cause values for all types of interfaces. It should be used only
    // if a management station is unable to decode the value of
    // dialCtlPeerStatsLastDisconnectCause. The type is string.
    CallHistoryDisconnectText interface{}

    // The value of sysUpTime when the call was connected. The type is interface{}
    // with range: 0..4294967295.
    CallHistoryConnectTime interface{}

    // The value of sysUpTime when the call was disconnected. The type is
    // interface{} with range: 0..4294967295.
    CallHistoryDisconnectTime interface{}

    // The call origin. The type is CallHistoryCallOrigin.
    CallHistoryCallOrigin interface{}

    // The number of charged units for this connection. For incoming calls or if
    // charging information is not supplied by the switch, the value of this
    // object will be zero. The type is interface{} with range: 0..4294967295.
    CallHistoryChargedUnits interface{}

    // The information type for this call. The type is CallHistoryInfoType.
    CallHistoryInfoType interface{}

    // The number of packets which were transmitted while this call was active.
    // The type is interface{} with range: 0..4294967295.
    CallHistoryTransmitPackets interface{}

    // The number of bytes which were transmitted while this call was active. The
    // type is interface{} with range: 0..4294967295.
    CallHistoryTransmitBytes interface{}

    // The number of packets which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    CallHistoryReceivePackets interface{}

    // The number of bytes which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    CallHistoryReceiveBytes interface{}
}

func (callHistoryEntry *DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry) GetEntityData() *types.CommonEntityData {
    callHistoryEntry.EntityData.YFilter = callHistoryEntry.YFilter
    callHistoryEntry.EntityData.YangName = "callHistoryEntry"
    callHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    callHistoryEntry.EntityData.ParentYangName = "callHistoryTable"
    callHistoryEntry.EntityData.SegmentPath = "callHistoryEntry" + types.AddKeyToken(callHistoryEntry.CallActiveSetupTime, "callActiveSetupTime") + types.AddKeyToken(callHistoryEntry.CallActiveIndex, "callActiveIndex")
    callHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callHistoryEntry.EntityData.Children = types.NewOrderedMap()
    callHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    callHistoryEntry.EntityData.Leafs.Append("callActiveSetupTime", types.YLeaf{"CallActiveSetupTime", callHistoryEntry.CallActiveSetupTime})
    callHistoryEntry.EntityData.Leafs.Append("callActiveIndex", types.YLeaf{"CallActiveIndex", callHistoryEntry.CallActiveIndex})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryPeerAddress", types.YLeaf{"CallHistoryPeerAddress", callHistoryEntry.CallHistoryPeerAddress})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryPeerSubAddress", types.YLeaf{"CallHistoryPeerSubAddress", callHistoryEntry.CallHistoryPeerSubAddress})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryPeerId", types.YLeaf{"CallHistoryPeerId", callHistoryEntry.CallHistoryPeerId})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryPeerIfIndex", types.YLeaf{"CallHistoryPeerIfIndex", callHistoryEntry.CallHistoryPeerIfIndex})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryLogicalIfIndex", types.YLeaf{"CallHistoryLogicalIfIndex", callHistoryEntry.CallHistoryLogicalIfIndex})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryDisconnectCause", types.YLeaf{"CallHistoryDisconnectCause", callHistoryEntry.CallHistoryDisconnectCause})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryDisconnectText", types.YLeaf{"CallHistoryDisconnectText", callHistoryEntry.CallHistoryDisconnectText})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryConnectTime", types.YLeaf{"CallHistoryConnectTime", callHistoryEntry.CallHistoryConnectTime})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryDisconnectTime", types.YLeaf{"CallHistoryDisconnectTime", callHistoryEntry.CallHistoryDisconnectTime})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryCallOrigin", types.YLeaf{"CallHistoryCallOrigin", callHistoryEntry.CallHistoryCallOrigin})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryChargedUnits", types.YLeaf{"CallHistoryChargedUnits", callHistoryEntry.CallHistoryChargedUnits})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryInfoType", types.YLeaf{"CallHistoryInfoType", callHistoryEntry.CallHistoryInfoType})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryTransmitPackets", types.YLeaf{"CallHistoryTransmitPackets", callHistoryEntry.CallHistoryTransmitPackets})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryTransmitBytes", types.YLeaf{"CallHistoryTransmitBytes", callHistoryEntry.CallHistoryTransmitBytes})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryReceivePackets", types.YLeaf{"CallHistoryReceivePackets", callHistoryEntry.CallHistoryReceivePackets})
    callHistoryEntry.EntityData.Leafs.Append("callHistoryReceiveBytes", types.YLeaf{"CallHistoryReceiveBytes", callHistoryEntry.CallHistoryReceiveBytes})

    callHistoryEntry.EntityData.YListKeys = []string {"CallActiveSetupTime", "CallActiveIndex"}

    return &(callHistoryEntry.EntityData)
}

// DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin represents The call origin.
type DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin string

const (
    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin_originate DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin = "originate"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin_answer DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin = "answer"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin_callback DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryCallOrigin = "callback"
)

// DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType represents The information type for this call.
type DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType string

const (
    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_other DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "other"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_speech DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "speech"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_unrestrictedDigital DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "unrestrictedDigital"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_unrestrictedDigital56 DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "unrestrictedDigital56"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_restrictedDigital DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "restrictedDigital"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_audio31 DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "audio31"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_audio7 DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "audio7"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_video DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "video"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_packetSwitched DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "packetSwitched"

    DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType_fax DIALCONTROLMIB_CallHistoryTable_CallHistoryEntry_CallHistoryInfoType = "fax"
)

