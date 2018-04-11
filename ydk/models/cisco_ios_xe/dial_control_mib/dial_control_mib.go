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

    
    Dialctlconfiguration DIALCONTROLMIB_Dialctlconfiguration

    
    Callhistory DIALCONTROLMIB_Callhistory

    // The list of peers from which the managed device will accept calls or to
    // which it will place them.
    Dialctlpeercfgtable DIALCONTROLMIB_Dialctlpeercfgtable

    // A table containing information about active calls to a specific
    // destination.
    Callactivetable DIALCONTROLMIB_Callactivetable

    // A table containing information about specific calls to a specific
    // destination.
    Callhistorytable DIALCONTROLMIB_Callhistorytable
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

    dIALCONTROLMIB.EntityData.Children = make(map[string]types.YChild)
    dIALCONTROLMIB.EntityData.Children["dialCtlConfiguration"] = types.YChild{"Dialctlconfiguration", &dIALCONTROLMIB.Dialctlconfiguration}
    dIALCONTROLMIB.EntityData.Children["callHistory"] = types.YChild{"Callhistory", &dIALCONTROLMIB.Callhistory}
    dIALCONTROLMIB.EntityData.Children["dialCtlPeerCfgTable"] = types.YChild{"Dialctlpeercfgtable", &dIALCONTROLMIB.Dialctlpeercfgtable}
    dIALCONTROLMIB.EntityData.Children["callActiveTable"] = types.YChild{"Callactivetable", &dIALCONTROLMIB.Callactivetable}
    dIALCONTROLMIB.EntityData.Children["callHistoryTable"] = types.YChild{"Callhistorytable", &dIALCONTROLMIB.Callhistorytable}
    dIALCONTROLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dIALCONTROLMIB.EntityData)
}

// DIALCONTROLMIB_Dialctlconfiguration
type DIALCONTROLMIB_Dialctlconfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The security level for acceptance of incoming calls. acceptNone(1)  -
    // incoming calls will not be accepted acceptAll(2)   - incoming calls will be
    // accepted,                  even if there is no matching entry              
    // in the dialCtlPeerCfgTable acceptKnown(3) - incoming calls will be accepted
    // only                  if there is a matching entry in the                 
    // dialCtlPeerCfgTable. The type is Dialctlacceptmode.
    Dialctlacceptmode interface{}

    // This object indicates whether dialCtlPeerCallInformation and
    // dialCtlPeerCallSetup traps should be generated for all peers. If the value
    // of this object is enabled(1), traps will be generated for all peers. If the
    // value of this object is disabled(2), traps will be generated only for peers
    // having dialCtlPeerCfgTrapEnable set to enabled(1). The type is
    // Dialctltrapenable.
    Dialctltrapenable interface{}
}

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetEntityData() *types.CommonEntityData {
    dialctlconfiguration.EntityData.YFilter = dialctlconfiguration.YFilter
    dialctlconfiguration.EntityData.YangName = "dialCtlConfiguration"
    dialctlconfiguration.EntityData.BundleName = "cisco_ios_xe"
    dialctlconfiguration.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    dialctlconfiguration.EntityData.SegmentPath = "dialCtlConfiguration"
    dialctlconfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dialctlconfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dialctlconfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dialctlconfiguration.EntityData.Children = make(map[string]types.YChild)
    dialctlconfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    dialctlconfiguration.EntityData.Leafs["dialCtlAcceptMode"] = types.YLeaf{"Dialctlacceptmode", dialctlconfiguration.Dialctlacceptmode}
    dialctlconfiguration.EntityData.Leafs["dialCtlTrapEnable"] = types.YLeaf{"Dialctltrapenable", dialctlconfiguration.Dialctltrapenable}
    return &(dialctlconfiguration.EntityData)
}

// DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode represents                  dialCtlPeerCfgTable
type DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode string

const (
    DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode_acceptNone DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode = "acceptNone"

    DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode_acceptAll DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode = "acceptAll"

    DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode_acceptKnown DIALCONTROLMIB_Dialctlconfiguration_Dialctlacceptmode = "acceptKnown"
)

// DIALCONTROLMIB_Dialctlconfiguration_Dialctltrapenable represents to enabled(1).
type DIALCONTROLMIB_Dialctlconfiguration_Dialctltrapenable string

const (
    DIALCONTROLMIB_Dialctlconfiguration_Dialctltrapenable_enabled DIALCONTROLMIB_Dialctlconfiguration_Dialctltrapenable = "enabled"

    DIALCONTROLMIB_Dialctlconfiguration_Dialctltrapenable_disabled DIALCONTROLMIB_Dialctlconfiguration_Dialctltrapenable = "disabled"
)

// DIALCONTROLMIB_Callhistory
type DIALCONTROLMIB_Callhistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The upper limit on the number of entries that the callHistoryTable may
    // contain.  A value of 0 will prevent any history from being retained. When
    // this table is full, the oldest entry will be deleted and the new one will
    // be created. The type is interface{} with range: 0..2147483647.
    Callhistorytablemaxlength interface{}

    // The minimum amount of time that an callHistoryEntry will be maintained
    // before being deleted. A value of 0 will prevent any history from being
    // retained in the callHistoryTable, but will neither prevent callCompletion
    // traps being generated nor affect other tables. The type is interface{} with
    // range: 0..2147483647. Units are minutes.
    Callhistoryretaintimer interface{}
}

func (callhistory *DIALCONTROLMIB_Callhistory) GetEntityData() *types.CommonEntityData {
    callhistory.EntityData.YFilter = callhistory.YFilter
    callhistory.EntityData.YangName = "callHistory"
    callhistory.EntityData.BundleName = "cisco_ios_xe"
    callhistory.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    callhistory.EntityData.SegmentPath = "callHistory"
    callhistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callhistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callhistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callhistory.EntityData.Children = make(map[string]types.YChild)
    callhistory.EntityData.Leafs = make(map[string]types.YLeaf)
    callhistory.EntityData.Leafs["callHistoryTableMaxLength"] = types.YLeaf{"Callhistorytablemaxlength", callhistory.Callhistorytablemaxlength}
    callhistory.EntityData.Leafs["callHistoryRetainTimer"] = types.YLeaf{"Callhistoryretaintimer", callhistory.Callhistoryretaintimer}
    return &(callhistory.EntityData)
}

// DIALCONTROLMIB_Dialctlpeercfgtable
// The list of peers from which the managed device
// will accept calls or to which it will place them.
type DIALCONTROLMIB_Dialctlpeercfgtable struct {
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
    // DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry.
    Dialctlpeercfgentry []DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry
}

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetEntityData() *types.CommonEntityData {
    dialctlpeercfgtable.EntityData.YFilter = dialctlpeercfgtable.YFilter
    dialctlpeercfgtable.EntityData.YangName = "dialCtlPeerCfgTable"
    dialctlpeercfgtable.EntityData.BundleName = "cisco_ios_xe"
    dialctlpeercfgtable.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    dialctlpeercfgtable.EntityData.SegmentPath = "dialCtlPeerCfgTable"
    dialctlpeercfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dialctlpeercfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dialctlpeercfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dialctlpeercfgtable.EntityData.Children = make(map[string]types.YChild)
    dialctlpeercfgtable.EntityData.Children["dialCtlPeerCfgEntry"] = types.YChild{"Dialctlpeercfgentry", nil}
    for i := range dialctlpeercfgtable.Dialctlpeercfgentry {
        dialctlpeercfgtable.EntityData.Children[types.GetSegmentPath(&dialctlpeercfgtable.Dialctlpeercfgentry[i])] = types.YChild{"Dialctlpeercfgentry", &dialctlpeercfgtable.Dialctlpeercfgentry[i]}
    }
    dialctlpeercfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dialctlpeercfgtable.EntityData)
}

// DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry
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
type DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry struct {
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
    Dialctlpeercfgid interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The interface type to be used for calling this peer. In case of ISDN, the
    // value of isdn(63) is to be used. The type is IANAifType.
    Dialctlpeercfgiftype interface{}

    // ifIndex value of an interface the peer will have to be called on. For
    // example, on an ISDN interface, this can be the ifIndex value of a D channel
    // or the ifIndex value of a B channel, whatever is appropriate for a given
    // peer. As an example, for Basic Rate leased lines it will be necessary to
    // specify a B channel ifIndex, while for     semi-permanent connections the D
    // channel ifIndex has to be specified. If the interface can be dynamically
    // assigned, this object has a value of zero. The type is interface{} with
    // range: 0..2147483647.
    Dialctlpeercfglowerif interface{}

    // Call Address at which the peer will be called. Think of this as the set of
    // characters following 'ATDT ' or the 'phone number' included in a D channel
    // call request.  The structure of this information will be switch type
    // specific. If there is no address information required for reaching the
    // peer, i.e., for leased lines, this object will be a zero length string. The
    // type is string.
    Dialctlpeercfgoriginateaddress interface{}

    // Calling Party Number information element, as for example passed in an ISDN
    // SETUP message by a PBX or switch, for incoming calls. This address can be
    // used to identify the peer. If this address is either unknown or identical
    // to dialCtlPeerCfgOriginateAddress, this object will be a zero length
    // string. The type is string.
    Dialctlpeercfgansweraddress interface{}

    // Subaddress at which the peer will be called. If the subaddress is undefined
    // for the given media or unused, this is a zero length string. The type is
    // string.
    Dialctlpeercfgsubaddress interface{}

    // Closed User Group at which the peer will be called. If the Closed User
    // Group is undefined for the given media or unused, this is a zero length
    // string. The type is string.
    Dialctlpeercfgclosedusergroup interface{}

    // The desired information transfer speed in bits/second when calling this
    // peer. The detailed media specific information, e.g. information type and
    // information transfer rate for ISDN circuits, has to be extracted from this
    // object. If the transfer speed to be used is unknown or the default speed
    // for this type of interfaces, the value of this object may be zero. The type
    // is interface{} with range: 0..2147483647.
    Dialctlpeercfgspeed interface{}

    // The Information Transfer Capability to be used when calling this peer. 
    // speech(2) refers to a non-data connection, whereas audio31(6) and audio7(7)
    // refer to data mode connections. The type is Dialctlpeercfginfotype.
    Dialctlpeercfginfotype interface{}

    // Applicable permissions. callback(4) either rejects the call and then calls
    // back, or uses the 'Reverse charging' information element if it is
    // available. Note that callback(4) is supposed to control charging, not
    // security, and applies to callback prior to accepting a call. Callback for
    // security reasons can be handled using PPP callback. The type is
    // Dialctlpeercfgpermission.
    Dialctlpeercfgpermission interface{}

    // The connection will be automatically disconnected if no longer carrying
    // useful data for a time period, in seconds, specified in this object. Useful
    // data in this context refers to forwarding packets, including routing
    // information; it excludes the encapsulator maintenance frames. A value of
    // zero means the connection will not be automatically taken down due to
    // inactivity, which implies that it is a dedicated circuit. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    Dialctlpeercfginactivitytimer interface{}

    // Minimum duration of a call in seconds, starting from the time the call is
    // connected until the call is disconnected. This is to accomplish the fact
    // that in most countries charging applies to units of time, which should be
    // matched as closely as possible. The type is interface{} with range:
    // 0..2147483647.
    Dialctlpeercfgminduration interface{}

    // Maximum call duration in seconds. Zero means 'unlimited'. The type is
    // interface{} with range: 0..2147483647.
    Dialctlpeercfgmaxduration interface{}

    // The call timeout time in seconds. The default value of zero means that the
    // call timeout as specified for the media in question will apply. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    Dialctlpeercfgcarrierdelay interface{}

    // The number of calls to a non-responding address that may be made. A retry
    // count of zero means there is no bound. The intent is to bound the number of
    // successive calls to an address which is inaccessible, or which refuses
    // those calls.  Some countries regulate the number of call retries to a given
    // peer that can be made. The type is interface{} with range: 0..2147483647.
    Dialctlpeercfgcallretries interface{}

    // The time in seconds between call retries if a peer cannot be reached. A
    // value of zero means that call retries may be done without any delay. The
    // type is interface{} with range: 0..2147483647. Units are seconds.
    Dialctlpeercfgretrydelay interface{}

    // The time in seconds after which call attempts are to be placed again after
    // a peer has been noticed to be unreachable, i.e. after
    // dialCtlPeerCfgCallRetries unsuccessful call attempts. A value of zero means
    // that a peer will not be called again after dialCtlPeerCfgCallRetries
    // unsuccessful call attempts. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    Dialctlpeercfgfailuredelay interface{}

    // This object indicates whether dialCtlPeerCallInformation and
    // dialCtlPeerCallSetup traps should be generated for this peer. The type is
    // Dialctlpeercfgtrapenable.
    Dialctlpeercfgtrapenable interface{}

    // Status of one row in this table. The type is RowStatus.
    Dialctlpeercfgstatus interface{}

    // Accumulated connect time to the peer since system startup. This is the
    // total connect time, i.e. the connect time for outgoing calls plus the time
    // for incoming calls. The type is interface{} with range: 0..4294967295.
    // Units are seconds.
    Dialctlpeerstatsconnecttime interface{}

    // The total number of charging units applying to this peer since system
    // startup. Only the charging units applying to the local interface, i.e. for
    // originated calls or for calls with 'Reverse charging' being active, will be
    // counted here. The type is interface{} with range: 0..4294967295.
    Dialctlpeerstatschargedunits interface{}

    // Number of completed calls to this peer. The type is interface{} with range:
    // 0..4294967295.
    Dialctlpeerstatssuccesscalls interface{}

    // Number of failed call attempts to this peer since system startup. The type
    // is interface{} with range: 0..4294967295.
    Dialctlpeerstatsfailcalls interface{}

    // Number of calls from this peer accepted since system startup. The type is
    // interface{} with range: 0..4294967295.
    Dialctlpeerstatsacceptcalls interface{}

    // Number of calls from this peer refused since system startup. The type is
    // interface{} with range: 0..4294967295.
    Dialctlpeerstatsrefusecalls interface{}

    // The encoded network cause value associated with the last call. This object
    // will be updated whenever a call is started or cleared. The value of this
    // object will depend on the interface type as well as on the protocol and
    // protocol version being used on this interface. Some references for possible
    // cause values are given below. The type is string with length: 0..4.
    Dialctlpeerstatslastdisconnectcause interface{}

    // ASCII text describing the reason for the last call termination.  This
    // object exists because it would be impossible for a management station to
    // store all possible cause values for all types of interfaces. It should be
    // used only if a management station is unable to decode the value of
    // dialCtlPeerStatsLastDisconnectCause.  This object will be updated whenever
    // a call is started or cleared. The type is string.
    Dialctlpeerstatslastdisconnecttext interface{}

    // The value of sysUpTime when the last call to this peer was started. For
    // ISDN media, this will be the time when the setup message was received from
    // or sent to the network. This object will be updated whenever a call is
    // started or cleared. The type is interface{} with range: 0..4294967295.
    Dialctlpeerstatslastsetuptime interface{}
}

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetEntityData() *types.CommonEntityData {
    dialctlpeercfgentry.EntityData.YFilter = dialctlpeercfgentry.YFilter
    dialctlpeercfgentry.EntityData.YangName = "dialCtlPeerCfgEntry"
    dialctlpeercfgentry.EntityData.BundleName = "cisco_ios_xe"
    dialctlpeercfgentry.EntityData.ParentYangName = "dialCtlPeerCfgTable"
    dialctlpeercfgentry.EntityData.SegmentPath = "dialCtlPeerCfgEntry" + "[dialCtlPeerCfgId='" + fmt.Sprintf("%v", dialctlpeercfgentry.Dialctlpeercfgid) + "']" + "[ifIndex='" + fmt.Sprintf("%v", dialctlpeercfgentry.Ifindex) + "']"
    dialctlpeercfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dialctlpeercfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dialctlpeercfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dialctlpeercfgentry.EntityData.Children = make(map[string]types.YChild)
    dialctlpeercfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgId"] = types.YLeaf{"Dialctlpeercfgid", dialctlpeercfgentry.Dialctlpeercfgid}
    dialctlpeercfgentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", dialctlpeercfgentry.Ifindex}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgIfType"] = types.YLeaf{"Dialctlpeercfgiftype", dialctlpeercfgentry.Dialctlpeercfgiftype}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgLowerIf"] = types.YLeaf{"Dialctlpeercfglowerif", dialctlpeercfgentry.Dialctlpeercfglowerif}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgOriginateAddress"] = types.YLeaf{"Dialctlpeercfgoriginateaddress", dialctlpeercfgentry.Dialctlpeercfgoriginateaddress}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgAnswerAddress"] = types.YLeaf{"Dialctlpeercfgansweraddress", dialctlpeercfgentry.Dialctlpeercfgansweraddress}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgSubAddress"] = types.YLeaf{"Dialctlpeercfgsubaddress", dialctlpeercfgentry.Dialctlpeercfgsubaddress}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgClosedUserGroup"] = types.YLeaf{"Dialctlpeercfgclosedusergroup", dialctlpeercfgentry.Dialctlpeercfgclosedusergroup}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgSpeed"] = types.YLeaf{"Dialctlpeercfgspeed", dialctlpeercfgentry.Dialctlpeercfgspeed}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgInfoType"] = types.YLeaf{"Dialctlpeercfginfotype", dialctlpeercfgentry.Dialctlpeercfginfotype}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgPermission"] = types.YLeaf{"Dialctlpeercfgpermission", dialctlpeercfgentry.Dialctlpeercfgpermission}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgInactivityTimer"] = types.YLeaf{"Dialctlpeercfginactivitytimer", dialctlpeercfgentry.Dialctlpeercfginactivitytimer}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgMinDuration"] = types.YLeaf{"Dialctlpeercfgminduration", dialctlpeercfgentry.Dialctlpeercfgminduration}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgMaxDuration"] = types.YLeaf{"Dialctlpeercfgmaxduration", dialctlpeercfgentry.Dialctlpeercfgmaxduration}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgCarrierDelay"] = types.YLeaf{"Dialctlpeercfgcarrierdelay", dialctlpeercfgentry.Dialctlpeercfgcarrierdelay}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgCallRetries"] = types.YLeaf{"Dialctlpeercfgcallretries", dialctlpeercfgentry.Dialctlpeercfgcallretries}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgRetryDelay"] = types.YLeaf{"Dialctlpeercfgretrydelay", dialctlpeercfgentry.Dialctlpeercfgretrydelay}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgFailureDelay"] = types.YLeaf{"Dialctlpeercfgfailuredelay", dialctlpeercfgentry.Dialctlpeercfgfailuredelay}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgTrapEnable"] = types.YLeaf{"Dialctlpeercfgtrapenable", dialctlpeercfgentry.Dialctlpeercfgtrapenable}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerCfgStatus"] = types.YLeaf{"Dialctlpeercfgstatus", dialctlpeercfgentry.Dialctlpeercfgstatus}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsConnectTime"] = types.YLeaf{"Dialctlpeerstatsconnecttime", dialctlpeercfgentry.Dialctlpeerstatsconnecttime}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsChargedUnits"] = types.YLeaf{"Dialctlpeerstatschargedunits", dialctlpeercfgentry.Dialctlpeerstatschargedunits}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsSuccessCalls"] = types.YLeaf{"Dialctlpeerstatssuccesscalls", dialctlpeercfgentry.Dialctlpeerstatssuccesscalls}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsFailCalls"] = types.YLeaf{"Dialctlpeerstatsfailcalls", dialctlpeercfgentry.Dialctlpeerstatsfailcalls}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsAcceptCalls"] = types.YLeaf{"Dialctlpeerstatsacceptcalls", dialctlpeercfgentry.Dialctlpeerstatsacceptcalls}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsRefuseCalls"] = types.YLeaf{"Dialctlpeerstatsrefusecalls", dialctlpeercfgentry.Dialctlpeerstatsrefusecalls}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsLastDisconnectCause"] = types.YLeaf{"Dialctlpeerstatslastdisconnectcause", dialctlpeercfgentry.Dialctlpeerstatslastdisconnectcause}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsLastDisconnectText"] = types.YLeaf{"Dialctlpeerstatslastdisconnecttext", dialctlpeercfgentry.Dialctlpeerstatslastdisconnecttext}
    dialctlpeercfgentry.EntityData.Leafs["dialCtlPeerStatsLastSetupTime"] = types.YLeaf{"Dialctlpeerstatslastsetuptime", dialctlpeercfgentry.Dialctlpeerstatslastsetuptime}
    return &(dialctlpeercfgentry.EntityData)
}

// DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype represents connections.
type DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype string

const (
    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_other DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "other"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_speech DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "speech"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_unrestrictedDigital DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "unrestrictedDigital"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_unrestrictedDigital56 DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "unrestrictedDigital56"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_restrictedDigital DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "restrictedDigital"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_audio31 DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "audio31"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_audio7 DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "audio7"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_video DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "video"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_packetSwitched DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "packetSwitched"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype_fax DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfginfotype = "fax"
)

// DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission represents PPP callback.
type DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission string

const (
    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission_originate DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission = "originate"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission_answer DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission = "answer"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission_both DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission = "both"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission_callback DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission = "callback"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission_none DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgpermission = "none"
)

// DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgtrapenable represents this peer.
type DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgtrapenable string

const (
    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgtrapenable_enabled DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgtrapenable = "enabled"

    DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgtrapenable_disabled DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry_Dialctlpeercfgtrapenable = "disabled"
)

// DIALCONTROLMIB_Callactivetable
// A table containing information about active
// calls to a specific destination.
type DIALCONTROLMIB_Callactivetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single active Connection. An entry in this
    // table will be created when a call is started. An entry in this table will
    // be deleted when an active call clears. The type is slice of
    // DIALCONTROLMIB_Callactivetable_Callactiveentry.
    Callactiveentry []DIALCONTROLMIB_Callactivetable_Callactiveentry
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetEntityData() *types.CommonEntityData {
    callactivetable.EntityData.YFilter = callactivetable.YFilter
    callactivetable.EntityData.YangName = "callActiveTable"
    callactivetable.EntityData.BundleName = "cisco_ios_xe"
    callactivetable.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    callactivetable.EntityData.SegmentPath = "callActiveTable"
    callactivetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callactivetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callactivetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callactivetable.EntityData.Children = make(map[string]types.YChild)
    callactivetable.EntityData.Children["callActiveEntry"] = types.YChild{"Callactiveentry", nil}
    for i := range callactivetable.Callactiveentry {
        callactivetable.EntityData.Children[types.GetSegmentPath(&callactivetable.Callactiveentry[i])] = types.YChild{"Callactiveentry", &callactivetable.Callactiveentry[i]}
    }
    callactivetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(callactivetable.EntityData)
}

// DIALCONTROLMIB_Callactivetable_Callactiveentry
// The information regarding a single active Connection.
// An entry in this table will be created when a call is
// started. An entry in this table will be deleted when
// an active call clears.
type DIALCONTROLMIB_Callactivetable_Callactiveentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of sysUpTime when the call associated to
    // this entry was started. This will be useful for an NMS to retrieve all
    // calls after a specific time. Also, this object can be useful in finding
    // large delays between the time the call was started and the time the call
    // was connected. For ISDN media, this will be the time when the setup message
    // was received from or sent to the network. The type is interface{} with
    // range: 0..4294967295.
    Callactivesetuptime interface{}

    // This attribute is a key. Small index variable to distinguish calls that
    // start in the same hundredth of a second. The type is interface{} with
    // range: 1..2147483647.
    Callactiveindex interface{}

    // The number this call is connected to. If the number is not available, then
    // it will have a length of zero. The type is string.
    Callactivepeeraddress interface{}

    // The subaddress this call is connected to. If the subaddress is undefined or
    // not available, this will be a zero length string. The type is string.
    Callactivepeersubaddress interface{}

    // This is the Id value of the peer table entry to which this call was made.
    // If a peer table entry for this call does not exist or is unknown, the value
    // of this object will be zero. The type is interface{} with range:
    // 0..2147483647.
    Callactivepeerid interface{}

    // This is the ifIndex value of the peer table entry to which this call was
    // made. If a peer table entry for this call does not exist or is unknown, the
    // value of this object will be zero. The type is interface{} with range:
    // 0..2147483647.
    Callactivepeerifindex interface{}

    // This is the ifIndex value of the logical interface through which this call
    // was made. For ISDN media, this would be the ifIndex of the B channel which
    // was used for this call. If the ifIndex value is unknown, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    Callactivelogicalifindex interface{}

    // The value of sysUpTime when the call was connected. If the call is not
    // connected, this object will have a value of zero. The type is interface{}
    // with range: 0..4294967295.
    Callactiveconnecttime interface{}

    // The current call state. unknown(1)     - The call state is unknown.
    // connecting(2)  - A connection attempt (outgoing call)                  is
    // being made. connected(3)   - An incoming call is in the process            
    // of validation. active(4)      - The call is active. The type is
    // Callactivecallstate.
    Callactivecallstate interface{}

    // The call origin. The type is Callactivecallorigin.
    Callactivecallorigin interface{}

    // The number of charged units for this connection. For incoming calls or if
    // charging information is not supplied by the switch, the value of this
    // object will be zero. The type is interface{} with range: 0..4294967295.
    Callactivechargedunits interface{}

    // The information type for this call. The type is Callactiveinfotype.
    Callactiveinfotype interface{}

    // The number of packets which were transmitted for this call. The type is
    // interface{} with range: 0..4294967295.
    Callactivetransmitpackets interface{}

    // The number of bytes which were transmitted for this call. The type is
    // interface{} with range: 0..4294967295.
    Callactivetransmitbytes interface{}

    // The number of packets which were received for this call. The type is
    // interface{} with range: 0..4294967295.
    Callactivereceivepackets interface{}

    // The number of bytes which were received for this call. The type is
    // interface{} with range: 0..4294967295.
    Callactivereceivebytes interface{}
}

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetEntityData() *types.CommonEntityData {
    callactiveentry.EntityData.YFilter = callactiveentry.YFilter
    callactiveentry.EntityData.YangName = "callActiveEntry"
    callactiveentry.EntityData.BundleName = "cisco_ios_xe"
    callactiveentry.EntityData.ParentYangName = "callActiveTable"
    callactiveentry.EntityData.SegmentPath = "callActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", callactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", callactiveentry.Callactiveindex) + "']"
    callactiveentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callactiveentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callactiveentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callactiveentry.EntityData.Children = make(map[string]types.YChild)
    callactiveentry.EntityData.Leafs = make(map[string]types.YLeaf)
    callactiveentry.EntityData.Leafs["callActiveSetupTime"] = types.YLeaf{"Callactivesetuptime", callactiveentry.Callactivesetuptime}
    callactiveentry.EntityData.Leafs["callActiveIndex"] = types.YLeaf{"Callactiveindex", callactiveentry.Callactiveindex}
    callactiveentry.EntityData.Leafs["callActivePeerAddress"] = types.YLeaf{"Callactivepeeraddress", callactiveentry.Callactivepeeraddress}
    callactiveentry.EntityData.Leafs["callActivePeerSubAddress"] = types.YLeaf{"Callactivepeersubaddress", callactiveentry.Callactivepeersubaddress}
    callactiveentry.EntityData.Leafs["callActivePeerId"] = types.YLeaf{"Callactivepeerid", callactiveentry.Callactivepeerid}
    callactiveentry.EntityData.Leafs["callActivePeerIfIndex"] = types.YLeaf{"Callactivepeerifindex", callactiveentry.Callactivepeerifindex}
    callactiveentry.EntityData.Leafs["callActiveLogicalIfIndex"] = types.YLeaf{"Callactivelogicalifindex", callactiveentry.Callactivelogicalifindex}
    callactiveentry.EntityData.Leafs["callActiveConnectTime"] = types.YLeaf{"Callactiveconnecttime", callactiveentry.Callactiveconnecttime}
    callactiveentry.EntityData.Leafs["callActiveCallState"] = types.YLeaf{"Callactivecallstate", callactiveentry.Callactivecallstate}
    callactiveentry.EntityData.Leafs["callActiveCallOrigin"] = types.YLeaf{"Callactivecallorigin", callactiveentry.Callactivecallorigin}
    callactiveentry.EntityData.Leafs["callActiveChargedUnits"] = types.YLeaf{"Callactivechargedunits", callactiveentry.Callactivechargedunits}
    callactiveentry.EntityData.Leafs["callActiveInfoType"] = types.YLeaf{"Callactiveinfotype", callactiveentry.Callactiveinfotype}
    callactiveentry.EntityData.Leafs["callActiveTransmitPackets"] = types.YLeaf{"Callactivetransmitpackets", callactiveentry.Callactivetransmitpackets}
    callactiveentry.EntityData.Leafs["callActiveTransmitBytes"] = types.YLeaf{"Callactivetransmitbytes", callactiveentry.Callactivetransmitbytes}
    callactiveentry.EntityData.Leafs["callActiveReceivePackets"] = types.YLeaf{"Callactivereceivepackets", callactiveentry.Callactivereceivepackets}
    callactiveentry.EntityData.Leafs["callActiveReceiveBytes"] = types.YLeaf{"Callactivereceivebytes", callactiveentry.Callactivereceivebytes}
    return &(callactiveentry.EntityData)
}

// DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin represents The call origin.
type DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin string

const (
    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin_originate DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin = "originate"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin_answer DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin = "answer"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin_callback DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallorigin = "callback"
)

// DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate represents active(4)      - The call is active.
type DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate string

const (
    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate_unknown DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate = "unknown"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate_connecting DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate = "connecting"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate_connected DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate = "connected"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate_active DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivecallstate = "active"
)

// DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype represents The information type for this call.
type DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype string

const (
    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_other DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "other"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_speech DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "speech"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_unrestrictedDigital DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "unrestrictedDigital"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_unrestrictedDigital56 DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "unrestrictedDigital56"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_restrictedDigital DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "restrictedDigital"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_audio31 DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "audio31"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_audio7 DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "audio7"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_video DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "video"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_packetSwitched DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "packetSwitched"

    DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype_fax DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveinfotype = "fax"
)

// DIALCONTROLMIB_Callhistorytable
// A table containing information about specific
// calls to a specific destination.
type DIALCONTROLMIB_Callhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single Connection. The type is slice of
    // DIALCONTROLMIB_Callhistorytable_Callhistoryentry.
    Callhistoryentry []DIALCONTROLMIB_Callhistorytable_Callhistoryentry
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetEntityData() *types.CommonEntityData {
    callhistorytable.EntityData.YFilter = callhistorytable.YFilter
    callhistorytable.EntityData.YangName = "callHistoryTable"
    callhistorytable.EntityData.BundleName = "cisco_ios_xe"
    callhistorytable.EntityData.ParentYangName = "DIAL-CONTROL-MIB"
    callhistorytable.EntityData.SegmentPath = "callHistoryTable"
    callhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callhistorytable.EntityData.Children = make(map[string]types.YChild)
    callhistorytable.EntityData.Children["callHistoryEntry"] = types.YChild{"Callhistoryentry", nil}
    for i := range callhistorytable.Callhistoryentry {
        callhistorytable.EntityData.Children[types.GetSegmentPath(&callhistorytable.Callhistoryentry[i])] = types.YChild{"Callhistoryentry", &callhistorytable.Callhistoryentry[i]}
    }
    callhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(callhistorytable.EntityData)
}

// DIALCONTROLMIB_Callhistorytable_Callhistoryentry
// The information regarding a single Connection.
type DIALCONTROLMIB_Callhistorytable_Callhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivesetuptime
    Callactivesetuptime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveindex
    Callactiveindex interface{}

    // The number this call was connected to. If the number is not available, then
    // it will have a length of zero. The type is string.
    Callhistorypeeraddress interface{}

    // The subaddress this call was connected to. If the subaddress is undefined
    // or not available, this will be a zero length string. The type is string.
    Callhistorypeersubaddress interface{}

    // This is the Id value of the peer table entry to which this call was made.
    // If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    Callhistorypeerid interface{}

    // This is the ifIndex value of the peer table entry to which this call was
    // made. If a peer table entry for this call does not exist, the value of this
    // object will be zero. The type is interface{} with range: 0..2147483647.
    Callhistorypeerifindex interface{}

    // This is the ifIndex value of the logical interface through which this call
    // was made. For ISDN media, this would be the ifIndex of the B channel which
    // was used for this call. The type is interface{} with range: 1..2147483647.
    Callhistorylogicalifindex interface{}

    // The encoded network cause value associated with this call.  The value of
    // this object will depend on the interface type as well as on the protocol
    // and protocol version being used on this interface. Some references for
    // possible cause values are given below. The type is string with length:
    // 0..4.
    Callhistorydisconnectcause interface{}

    // ASCII text describing the reason for call termination.  This object exists
    // because it would be impossible for a management station to store all
    // possible cause values for all types of interfaces. It should be used only
    // if a management station is unable to decode the value of
    // dialCtlPeerStatsLastDisconnectCause. The type is string.
    Callhistorydisconnecttext interface{}

    // The value of sysUpTime when the call was connected. The type is interface{}
    // with range: 0..4294967295.
    Callhistoryconnecttime interface{}

    // The value of sysUpTime when the call was disconnected. The type is
    // interface{} with range: 0..4294967295.
    Callhistorydisconnecttime interface{}

    // The call origin. The type is Callhistorycallorigin.
    Callhistorycallorigin interface{}

    // The number of charged units for this connection. For incoming calls or if
    // charging information is not supplied by the switch, the value of this
    // object will be zero. The type is interface{} with range: 0..4294967295.
    Callhistorychargedunits interface{}

    // The information type for this call. The type is Callhistoryinfotype.
    Callhistoryinfotype interface{}

    // The number of packets which were transmitted while this call was active.
    // The type is interface{} with range: 0..4294967295.
    Callhistorytransmitpackets interface{}

    // The number of bytes which were transmitted while this call was active. The
    // type is interface{} with range: 0..4294967295.
    Callhistorytransmitbytes interface{}

    // The number of packets which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    Callhistoryreceivepackets interface{}

    // The number of bytes which were received while this call was active. The
    // type is interface{} with range: 0..4294967295.
    Callhistoryreceivebytes interface{}
}

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetEntityData() *types.CommonEntityData {
    callhistoryentry.EntityData.YFilter = callhistoryentry.YFilter
    callhistoryentry.EntityData.YangName = "callHistoryEntry"
    callhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    callhistoryentry.EntityData.ParentYangName = "callHistoryTable"
    callhistoryentry.EntityData.SegmentPath = "callHistoryEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", callhistoryentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", callhistoryentry.Callactiveindex) + "']"
    callhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callhistoryentry.EntityData.Children = make(map[string]types.YChild)
    callhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    callhistoryentry.EntityData.Leafs["callActiveSetupTime"] = types.YLeaf{"Callactivesetuptime", callhistoryentry.Callactivesetuptime}
    callhistoryentry.EntityData.Leafs["callActiveIndex"] = types.YLeaf{"Callactiveindex", callhistoryentry.Callactiveindex}
    callhistoryentry.EntityData.Leafs["callHistoryPeerAddress"] = types.YLeaf{"Callhistorypeeraddress", callhistoryentry.Callhistorypeeraddress}
    callhistoryentry.EntityData.Leafs["callHistoryPeerSubAddress"] = types.YLeaf{"Callhistorypeersubaddress", callhistoryentry.Callhistorypeersubaddress}
    callhistoryentry.EntityData.Leafs["callHistoryPeerId"] = types.YLeaf{"Callhistorypeerid", callhistoryentry.Callhistorypeerid}
    callhistoryentry.EntityData.Leafs["callHistoryPeerIfIndex"] = types.YLeaf{"Callhistorypeerifindex", callhistoryentry.Callhistorypeerifindex}
    callhistoryentry.EntityData.Leafs["callHistoryLogicalIfIndex"] = types.YLeaf{"Callhistorylogicalifindex", callhistoryentry.Callhistorylogicalifindex}
    callhistoryentry.EntityData.Leafs["callHistoryDisconnectCause"] = types.YLeaf{"Callhistorydisconnectcause", callhistoryentry.Callhistorydisconnectcause}
    callhistoryentry.EntityData.Leafs["callHistoryDisconnectText"] = types.YLeaf{"Callhistorydisconnecttext", callhistoryentry.Callhistorydisconnecttext}
    callhistoryentry.EntityData.Leafs["callHistoryConnectTime"] = types.YLeaf{"Callhistoryconnecttime", callhistoryentry.Callhistoryconnecttime}
    callhistoryentry.EntityData.Leafs["callHistoryDisconnectTime"] = types.YLeaf{"Callhistorydisconnecttime", callhistoryentry.Callhistorydisconnecttime}
    callhistoryentry.EntityData.Leafs["callHistoryCallOrigin"] = types.YLeaf{"Callhistorycallorigin", callhistoryentry.Callhistorycallorigin}
    callhistoryentry.EntityData.Leafs["callHistoryChargedUnits"] = types.YLeaf{"Callhistorychargedunits", callhistoryentry.Callhistorychargedunits}
    callhistoryentry.EntityData.Leafs["callHistoryInfoType"] = types.YLeaf{"Callhistoryinfotype", callhistoryentry.Callhistoryinfotype}
    callhistoryentry.EntityData.Leafs["callHistoryTransmitPackets"] = types.YLeaf{"Callhistorytransmitpackets", callhistoryentry.Callhistorytransmitpackets}
    callhistoryentry.EntityData.Leafs["callHistoryTransmitBytes"] = types.YLeaf{"Callhistorytransmitbytes", callhistoryentry.Callhistorytransmitbytes}
    callhistoryentry.EntityData.Leafs["callHistoryReceivePackets"] = types.YLeaf{"Callhistoryreceivepackets", callhistoryentry.Callhistoryreceivepackets}
    callhistoryentry.EntityData.Leafs["callHistoryReceiveBytes"] = types.YLeaf{"Callhistoryreceivebytes", callhistoryentry.Callhistoryreceivebytes}
    return &(callhistoryentry.EntityData)
}

// DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin represents The call origin.
type DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin string

const (
    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin_originate DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin = "originate"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin_answer DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin = "answer"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin_callback DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistorycallorigin = "callback"
)

// DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype represents The information type for this call.
type DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype string

const (
    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_other DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "other"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_speech DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "speech"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_unrestrictedDigital DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "unrestrictedDigital"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_unrestrictedDigital56 DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "unrestrictedDigital56"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_restrictedDigital DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "restrictedDigital"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_audio31 DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "audio31"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_audio7 DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "audio7"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_video DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "video"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_packetSwitched DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "packetSwitched"

    DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype_fax DIALCONTROLMIB_Callhistorytable_Callhistoryentry_Callhistoryinfotype = "fax"
)

