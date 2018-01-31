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
    parent types.Entity
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

func (dIALCONTROLMIB *DIALCONTROLMIB) GetFilter() yfilter.YFilter { return dIALCONTROLMIB.YFilter }

func (dIALCONTROLMIB *DIALCONTROLMIB) SetFilter(yf yfilter.YFilter) { dIALCONTROLMIB.YFilter = yf }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetGoName(yname string) string {
    if yname == "dialCtlConfiguration" { return "Dialctlconfiguration" }
    if yname == "callHistory" { return "Callhistory" }
    if yname == "dialCtlPeerCfgTable" { return "Dialctlpeercfgtable" }
    if yname == "callActiveTable" { return "Callactivetable" }
    if yname == "callHistoryTable" { return "Callhistorytable" }
    return ""
}

func (dIALCONTROLMIB *DIALCONTROLMIB) GetSegmentPath() string {
    return "DIAL-CONTROL-MIB:DIAL-CONTROL-MIB"
}

func (dIALCONTROLMIB *DIALCONTROLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dialCtlConfiguration" {
        return &dIALCONTROLMIB.Dialctlconfiguration
    }
    if childYangName == "callHistory" {
        return &dIALCONTROLMIB.Callhistory
    }
    if childYangName == "dialCtlPeerCfgTable" {
        return &dIALCONTROLMIB.Dialctlpeercfgtable
    }
    if childYangName == "callActiveTable" {
        return &dIALCONTROLMIB.Callactivetable
    }
    if childYangName == "callHistoryTable" {
        return &dIALCONTROLMIB.Callhistorytable
    }
    return nil
}

func (dIALCONTROLMIB *DIALCONTROLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dialCtlConfiguration"] = &dIALCONTROLMIB.Dialctlconfiguration
    children["callHistory"] = &dIALCONTROLMIB.Callhistory
    children["dialCtlPeerCfgTable"] = &dIALCONTROLMIB.Dialctlpeercfgtable
    children["callActiveTable"] = &dIALCONTROLMIB.Callactivetable
    children["callHistoryTable"] = &dIALCONTROLMIB.Callhistorytable
    return children
}

func (dIALCONTROLMIB *DIALCONTROLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dIALCONTROLMIB *DIALCONTROLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetYangName() string { return "DIAL-CONTROL-MIB" }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dIALCONTROLMIB *DIALCONTROLMIB) SetParent(parent types.Entity) { dIALCONTROLMIB.parent = parent }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetParent() types.Entity { return dIALCONTROLMIB.parent }

func (dIALCONTROLMIB *DIALCONTROLMIB) GetParentYangName() string { return "DIAL-CONTROL-MIB" }

// DIALCONTROLMIB_Dialctlconfiguration
type DIALCONTROLMIB_Dialctlconfiguration struct {
    parent types.Entity
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

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetFilter() yfilter.YFilter { return dialctlconfiguration.YFilter }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) SetFilter(yf yfilter.YFilter) { dialctlconfiguration.YFilter = yf }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetGoName(yname string) string {
    if yname == "dialCtlAcceptMode" { return "Dialctlacceptmode" }
    if yname == "dialCtlTrapEnable" { return "Dialctltrapenable" }
    return ""
}

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetSegmentPath() string {
    return "dialCtlConfiguration"
}

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dialCtlAcceptMode"] = dialctlconfiguration.Dialctlacceptmode
    leafs["dialCtlTrapEnable"] = dialctlconfiguration.Dialctltrapenable
    return leafs
}

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetBundleName() string { return "cisco_ios_xe" }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetYangName() string { return "dialCtlConfiguration" }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) SetParent(parent types.Entity) { dialctlconfiguration.parent = parent }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetParent() types.Entity { return dialctlconfiguration.parent }

func (dialctlconfiguration *DIALCONTROLMIB_Dialctlconfiguration) GetParentYangName() string { return "DIAL-CONTROL-MIB" }

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
    parent types.Entity
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

func (callhistory *DIALCONTROLMIB_Callhistory) GetFilter() yfilter.YFilter { return callhistory.YFilter }

func (callhistory *DIALCONTROLMIB_Callhistory) SetFilter(yf yfilter.YFilter) { callhistory.YFilter = yf }

func (callhistory *DIALCONTROLMIB_Callhistory) GetGoName(yname string) string {
    if yname == "callHistoryTableMaxLength" { return "Callhistorytablemaxlength" }
    if yname == "callHistoryRetainTimer" { return "Callhistoryretaintimer" }
    return ""
}

func (callhistory *DIALCONTROLMIB_Callhistory) GetSegmentPath() string {
    return "callHistory"
}

func (callhistory *DIALCONTROLMIB_Callhistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (callhistory *DIALCONTROLMIB_Callhistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (callhistory *DIALCONTROLMIB_Callhistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["callHistoryTableMaxLength"] = callhistory.Callhistorytablemaxlength
    leafs["callHistoryRetainTimer"] = callhistory.Callhistoryretaintimer
    return leafs
}

func (callhistory *DIALCONTROLMIB_Callhistory) GetBundleName() string { return "cisco_ios_xe" }

func (callhistory *DIALCONTROLMIB_Callhistory) GetYangName() string { return "callHistory" }

func (callhistory *DIALCONTROLMIB_Callhistory) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callhistory *DIALCONTROLMIB_Callhistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callhistory *DIALCONTROLMIB_Callhistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callhistory *DIALCONTROLMIB_Callhistory) SetParent(parent types.Entity) { callhistory.parent = parent }

func (callhistory *DIALCONTROLMIB_Callhistory) GetParent() types.Entity { return callhistory.parent }

func (callhistory *DIALCONTROLMIB_Callhistory) GetParentYangName() string { return "DIAL-CONTROL-MIB" }

// DIALCONTROLMIB_Dialctlpeercfgtable
// The list of peers from which the managed device
// will accept calls or to which it will place them.
type DIALCONTROLMIB_Dialctlpeercfgtable struct {
    parent types.Entity
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

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetFilter() yfilter.YFilter { return dialctlpeercfgtable.YFilter }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) SetFilter(yf yfilter.YFilter) { dialctlpeercfgtable.YFilter = yf }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetGoName(yname string) string {
    if yname == "dialCtlPeerCfgEntry" { return "Dialctlpeercfgentry" }
    return ""
}

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetSegmentPath() string {
    return "dialCtlPeerCfgTable"
}

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dialCtlPeerCfgEntry" {
        for _, c := range dialctlpeercfgtable.Dialctlpeercfgentry {
            if dialctlpeercfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry{}
        dialctlpeercfgtable.Dialctlpeercfgentry = append(dialctlpeercfgtable.Dialctlpeercfgentry, child)
        return &dialctlpeercfgtable.Dialctlpeercfgentry[len(dialctlpeercfgtable.Dialctlpeercfgentry)-1]
    }
    return nil
}

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dialctlpeercfgtable.Dialctlpeercfgentry {
        children[dialctlpeercfgtable.Dialctlpeercfgentry[i].GetSegmentPath()] = &dialctlpeercfgtable.Dialctlpeercfgentry[i]
    }
    return children
}

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetYangName() string { return "dialCtlPeerCfgTable" }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) SetParent(parent types.Entity) { dialctlpeercfgtable.parent = parent }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetParent() types.Entity { return dialctlpeercfgtable.parent }

func (dialctlpeercfgtable *DIALCONTROLMIB_Dialctlpeercfgtable) GetParentYangName() string { return "DIAL-CONTROL-MIB" }

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
    parent types.Entity
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

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetFilter() yfilter.YFilter { return dialctlpeercfgentry.YFilter }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) SetFilter(yf yfilter.YFilter) { dialctlpeercfgentry.YFilter = yf }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetGoName(yname string) string {
    if yname == "dialCtlPeerCfgId" { return "Dialctlpeercfgid" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "dialCtlPeerCfgIfType" { return "Dialctlpeercfgiftype" }
    if yname == "dialCtlPeerCfgLowerIf" { return "Dialctlpeercfglowerif" }
    if yname == "dialCtlPeerCfgOriginateAddress" { return "Dialctlpeercfgoriginateaddress" }
    if yname == "dialCtlPeerCfgAnswerAddress" { return "Dialctlpeercfgansweraddress" }
    if yname == "dialCtlPeerCfgSubAddress" { return "Dialctlpeercfgsubaddress" }
    if yname == "dialCtlPeerCfgClosedUserGroup" { return "Dialctlpeercfgclosedusergroup" }
    if yname == "dialCtlPeerCfgSpeed" { return "Dialctlpeercfgspeed" }
    if yname == "dialCtlPeerCfgInfoType" { return "Dialctlpeercfginfotype" }
    if yname == "dialCtlPeerCfgPermission" { return "Dialctlpeercfgpermission" }
    if yname == "dialCtlPeerCfgInactivityTimer" { return "Dialctlpeercfginactivitytimer" }
    if yname == "dialCtlPeerCfgMinDuration" { return "Dialctlpeercfgminduration" }
    if yname == "dialCtlPeerCfgMaxDuration" { return "Dialctlpeercfgmaxduration" }
    if yname == "dialCtlPeerCfgCarrierDelay" { return "Dialctlpeercfgcarrierdelay" }
    if yname == "dialCtlPeerCfgCallRetries" { return "Dialctlpeercfgcallretries" }
    if yname == "dialCtlPeerCfgRetryDelay" { return "Dialctlpeercfgretrydelay" }
    if yname == "dialCtlPeerCfgFailureDelay" { return "Dialctlpeercfgfailuredelay" }
    if yname == "dialCtlPeerCfgTrapEnable" { return "Dialctlpeercfgtrapenable" }
    if yname == "dialCtlPeerCfgStatus" { return "Dialctlpeercfgstatus" }
    if yname == "dialCtlPeerStatsConnectTime" { return "Dialctlpeerstatsconnecttime" }
    if yname == "dialCtlPeerStatsChargedUnits" { return "Dialctlpeerstatschargedunits" }
    if yname == "dialCtlPeerStatsSuccessCalls" { return "Dialctlpeerstatssuccesscalls" }
    if yname == "dialCtlPeerStatsFailCalls" { return "Dialctlpeerstatsfailcalls" }
    if yname == "dialCtlPeerStatsAcceptCalls" { return "Dialctlpeerstatsacceptcalls" }
    if yname == "dialCtlPeerStatsRefuseCalls" { return "Dialctlpeerstatsrefusecalls" }
    if yname == "dialCtlPeerStatsLastDisconnectCause" { return "Dialctlpeerstatslastdisconnectcause" }
    if yname == "dialCtlPeerStatsLastDisconnectText" { return "Dialctlpeerstatslastdisconnecttext" }
    if yname == "dialCtlPeerStatsLastSetupTime" { return "Dialctlpeerstatslastsetuptime" }
    return ""
}

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetSegmentPath() string {
    return "dialCtlPeerCfgEntry" + "[dialCtlPeerCfgId='" + fmt.Sprintf("%v", dialctlpeercfgentry.Dialctlpeercfgid) + "']" + "[ifIndex='" + fmt.Sprintf("%v", dialctlpeercfgentry.Ifindex) + "']"
}

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dialCtlPeerCfgId"] = dialctlpeercfgentry.Dialctlpeercfgid
    leafs["ifIndex"] = dialctlpeercfgentry.Ifindex
    leafs["dialCtlPeerCfgIfType"] = dialctlpeercfgentry.Dialctlpeercfgiftype
    leafs["dialCtlPeerCfgLowerIf"] = dialctlpeercfgentry.Dialctlpeercfglowerif
    leafs["dialCtlPeerCfgOriginateAddress"] = dialctlpeercfgentry.Dialctlpeercfgoriginateaddress
    leafs["dialCtlPeerCfgAnswerAddress"] = dialctlpeercfgentry.Dialctlpeercfgansweraddress
    leafs["dialCtlPeerCfgSubAddress"] = dialctlpeercfgentry.Dialctlpeercfgsubaddress
    leafs["dialCtlPeerCfgClosedUserGroup"] = dialctlpeercfgentry.Dialctlpeercfgclosedusergroup
    leafs["dialCtlPeerCfgSpeed"] = dialctlpeercfgentry.Dialctlpeercfgspeed
    leafs["dialCtlPeerCfgInfoType"] = dialctlpeercfgentry.Dialctlpeercfginfotype
    leafs["dialCtlPeerCfgPermission"] = dialctlpeercfgentry.Dialctlpeercfgpermission
    leafs["dialCtlPeerCfgInactivityTimer"] = dialctlpeercfgentry.Dialctlpeercfginactivitytimer
    leafs["dialCtlPeerCfgMinDuration"] = dialctlpeercfgentry.Dialctlpeercfgminduration
    leafs["dialCtlPeerCfgMaxDuration"] = dialctlpeercfgentry.Dialctlpeercfgmaxduration
    leafs["dialCtlPeerCfgCarrierDelay"] = dialctlpeercfgentry.Dialctlpeercfgcarrierdelay
    leafs["dialCtlPeerCfgCallRetries"] = dialctlpeercfgentry.Dialctlpeercfgcallretries
    leafs["dialCtlPeerCfgRetryDelay"] = dialctlpeercfgentry.Dialctlpeercfgretrydelay
    leafs["dialCtlPeerCfgFailureDelay"] = dialctlpeercfgentry.Dialctlpeercfgfailuredelay
    leafs["dialCtlPeerCfgTrapEnable"] = dialctlpeercfgentry.Dialctlpeercfgtrapenable
    leafs["dialCtlPeerCfgStatus"] = dialctlpeercfgentry.Dialctlpeercfgstatus
    leafs["dialCtlPeerStatsConnectTime"] = dialctlpeercfgentry.Dialctlpeerstatsconnecttime
    leafs["dialCtlPeerStatsChargedUnits"] = dialctlpeercfgentry.Dialctlpeerstatschargedunits
    leafs["dialCtlPeerStatsSuccessCalls"] = dialctlpeercfgentry.Dialctlpeerstatssuccesscalls
    leafs["dialCtlPeerStatsFailCalls"] = dialctlpeercfgentry.Dialctlpeerstatsfailcalls
    leafs["dialCtlPeerStatsAcceptCalls"] = dialctlpeercfgentry.Dialctlpeerstatsacceptcalls
    leafs["dialCtlPeerStatsRefuseCalls"] = dialctlpeercfgentry.Dialctlpeerstatsrefusecalls
    leafs["dialCtlPeerStatsLastDisconnectCause"] = dialctlpeercfgentry.Dialctlpeerstatslastdisconnectcause
    leafs["dialCtlPeerStatsLastDisconnectText"] = dialctlpeercfgentry.Dialctlpeerstatslastdisconnecttext
    leafs["dialCtlPeerStatsLastSetupTime"] = dialctlpeercfgentry.Dialctlpeerstatslastsetuptime
    return leafs
}

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetYangName() string { return "dialCtlPeerCfgEntry" }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) SetParent(parent types.Entity) { dialctlpeercfgentry.parent = parent }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetParent() types.Entity { return dialctlpeercfgentry.parent }

func (dialctlpeercfgentry *DIALCONTROLMIB_Dialctlpeercfgtable_Dialctlpeercfgentry) GetParentYangName() string { return "dialCtlPeerCfgTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single active Connection. An entry in this
    // table will be created when a call is started. An entry in this table will
    // be deleted when an active call clears. The type is slice of
    // DIALCONTROLMIB_Callactivetable_Callactiveentry.
    Callactiveentry []DIALCONTROLMIB_Callactivetable_Callactiveentry
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetFilter() yfilter.YFilter { return callactivetable.YFilter }

func (callactivetable *DIALCONTROLMIB_Callactivetable) SetFilter(yf yfilter.YFilter) { callactivetable.YFilter = yf }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetGoName(yname string) string {
    if yname == "callActiveEntry" { return "Callactiveentry" }
    return ""
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetSegmentPath() string {
    return "callActiveTable"
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "callActiveEntry" {
        for _, c := range callactivetable.Callactiveentry {
            if callactivetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIALCONTROLMIB_Callactivetable_Callactiveentry{}
        callactivetable.Callactiveentry = append(callactivetable.Callactiveentry, child)
        return &callactivetable.Callactiveentry[len(callactivetable.Callactiveentry)-1]
    }
    return nil
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range callactivetable.Callactiveentry {
        children[callactivetable.Callactiveentry[i].GetSegmentPath()] = &callactivetable.Callactiveentry[i]
    }
    return children
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetBundleName() string { return "cisco_ios_xe" }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetYangName() string { return "callActiveTable" }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callactivetable *DIALCONTROLMIB_Callactivetable) SetParent(parent types.Entity) { callactivetable.parent = parent }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetParent() types.Entity { return callactivetable.parent }

func (callactivetable *DIALCONTROLMIB_Callactivetable) GetParentYangName() string { return "DIAL-CONTROL-MIB" }

// DIALCONTROLMIB_Callactivetable_Callactiveentry
// The information regarding a single active Connection.
// An entry in this table will be created when a call is
// started. An entry in this table will be deleted when
// an active call clears.
type DIALCONTROLMIB_Callactivetable_Callactiveentry struct {
    parent types.Entity
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

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetFilter() yfilter.YFilter { return callactiveentry.YFilter }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) SetFilter(yf yfilter.YFilter) { callactiveentry.YFilter = yf }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetGoName(yname string) string {
    if yname == "callActiveSetupTime" { return "Callactivesetuptime" }
    if yname == "callActiveIndex" { return "Callactiveindex" }
    if yname == "callActivePeerAddress" { return "Callactivepeeraddress" }
    if yname == "callActivePeerSubAddress" { return "Callactivepeersubaddress" }
    if yname == "callActivePeerId" { return "Callactivepeerid" }
    if yname == "callActivePeerIfIndex" { return "Callactivepeerifindex" }
    if yname == "callActiveLogicalIfIndex" { return "Callactivelogicalifindex" }
    if yname == "callActiveConnectTime" { return "Callactiveconnecttime" }
    if yname == "callActiveCallState" { return "Callactivecallstate" }
    if yname == "callActiveCallOrigin" { return "Callactivecallorigin" }
    if yname == "callActiveChargedUnits" { return "Callactivechargedunits" }
    if yname == "callActiveInfoType" { return "Callactiveinfotype" }
    if yname == "callActiveTransmitPackets" { return "Callactivetransmitpackets" }
    if yname == "callActiveTransmitBytes" { return "Callactivetransmitbytes" }
    if yname == "callActiveReceivePackets" { return "Callactivereceivepackets" }
    if yname == "callActiveReceiveBytes" { return "Callactivereceivebytes" }
    return ""
}

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetSegmentPath() string {
    return "callActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", callactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", callactiveentry.Callactiveindex) + "']"
}

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["callActiveSetupTime"] = callactiveentry.Callactivesetuptime
    leafs["callActiveIndex"] = callactiveentry.Callactiveindex
    leafs["callActivePeerAddress"] = callactiveentry.Callactivepeeraddress
    leafs["callActivePeerSubAddress"] = callactiveentry.Callactivepeersubaddress
    leafs["callActivePeerId"] = callactiveentry.Callactivepeerid
    leafs["callActivePeerIfIndex"] = callactiveentry.Callactivepeerifindex
    leafs["callActiveLogicalIfIndex"] = callactiveentry.Callactivelogicalifindex
    leafs["callActiveConnectTime"] = callactiveentry.Callactiveconnecttime
    leafs["callActiveCallState"] = callactiveentry.Callactivecallstate
    leafs["callActiveCallOrigin"] = callactiveentry.Callactivecallorigin
    leafs["callActiveChargedUnits"] = callactiveentry.Callactivechargedunits
    leafs["callActiveInfoType"] = callactiveentry.Callactiveinfotype
    leafs["callActiveTransmitPackets"] = callactiveentry.Callactivetransmitpackets
    leafs["callActiveTransmitBytes"] = callactiveentry.Callactivetransmitbytes
    leafs["callActiveReceivePackets"] = callactiveentry.Callactivereceivepackets
    leafs["callActiveReceiveBytes"] = callactiveentry.Callactivereceivebytes
    return leafs
}

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetBundleName() string { return "cisco_ios_xe" }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetYangName() string { return "callActiveEntry" }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) SetParent(parent types.Entity) { callactiveentry.parent = parent }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetParent() types.Entity { return callactiveentry.parent }

func (callactiveentry *DIALCONTROLMIB_Callactivetable_Callactiveentry) GetParentYangName() string { return "callActiveTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single Connection. The type is slice of
    // DIALCONTROLMIB_Callhistorytable_Callhistoryentry.
    Callhistoryentry []DIALCONTROLMIB_Callhistorytable_Callhistoryentry
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetFilter() yfilter.YFilter { return callhistorytable.YFilter }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) SetFilter(yf yfilter.YFilter) { callhistorytable.YFilter = yf }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetGoName(yname string) string {
    if yname == "callHistoryEntry" { return "Callhistoryentry" }
    return ""
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetSegmentPath() string {
    return "callHistoryTable"
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "callHistoryEntry" {
        for _, c := range callhistorytable.Callhistoryentry {
            if callhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIALCONTROLMIB_Callhistorytable_Callhistoryentry{}
        callhistorytable.Callhistoryentry = append(callhistorytable.Callhistoryentry, child)
        return &callhistorytable.Callhistoryentry[len(callhistorytable.Callhistoryentry)-1]
    }
    return nil
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range callhistorytable.Callhistoryentry {
        children[callhistorytable.Callhistoryentry[i].GetSegmentPath()] = &callhistorytable.Callhistoryentry[i]
    }
    return children
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetYangName() string { return "callHistoryTable" }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) SetParent(parent types.Entity) { callhistorytable.parent = parent }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetParent() types.Entity { return callhistorytable.parent }

func (callhistorytable *DIALCONTROLMIB_Callhistorytable) GetParentYangName() string { return "DIAL-CONTROL-MIB" }

// DIALCONTROLMIB_Callhistorytable_Callhistoryentry
// The information regarding a single Connection.
type DIALCONTROLMIB_Callhistorytable_Callhistoryentry struct {
    parent types.Entity
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

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetFilter() yfilter.YFilter { return callhistoryentry.YFilter }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) SetFilter(yf yfilter.YFilter) { callhistoryentry.YFilter = yf }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetGoName(yname string) string {
    if yname == "callActiveSetupTime" { return "Callactivesetuptime" }
    if yname == "callActiveIndex" { return "Callactiveindex" }
    if yname == "callHistoryPeerAddress" { return "Callhistorypeeraddress" }
    if yname == "callHistoryPeerSubAddress" { return "Callhistorypeersubaddress" }
    if yname == "callHistoryPeerId" { return "Callhistorypeerid" }
    if yname == "callHistoryPeerIfIndex" { return "Callhistorypeerifindex" }
    if yname == "callHistoryLogicalIfIndex" { return "Callhistorylogicalifindex" }
    if yname == "callHistoryDisconnectCause" { return "Callhistorydisconnectcause" }
    if yname == "callHistoryDisconnectText" { return "Callhistorydisconnecttext" }
    if yname == "callHistoryConnectTime" { return "Callhistoryconnecttime" }
    if yname == "callHistoryDisconnectTime" { return "Callhistorydisconnecttime" }
    if yname == "callHistoryCallOrigin" { return "Callhistorycallorigin" }
    if yname == "callHistoryChargedUnits" { return "Callhistorychargedunits" }
    if yname == "callHistoryInfoType" { return "Callhistoryinfotype" }
    if yname == "callHistoryTransmitPackets" { return "Callhistorytransmitpackets" }
    if yname == "callHistoryTransmitBytes" { return "Callhistorytransmitbytes" }
    if yname == "callHistoryReceivePackets" { return "Callhistoryreceivepackets" }
    if yname == "callHistoryReceiveBytes" { return "Callhistoryreceivebytes" }
    return ""
}

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetSegmentPath() string {
    return "callHistoryEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", callhistoryentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", callhistoryentry.Callactiveindex) + "']"
}

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["callActiveSetupTime"] = callhistoryentry.Callactivesetuptime
    leafs["callActiveIndex"] = callhistoryentry.Callactiveindex
    leafs["callHistoryPeerAddress"] = callhistoryentry.Callhistorypeeraddress
    leafs["callHistoryPeerSubAddress"] = callhistoryentry.Callhistorypeersubaddress
    leafs["callHistoryPeerId"] = callhistoryentry.Callhistorypeerid
    leafs["callHistoryPeerIfIndex"] = callhistoryentry.Callhistorypeerifindex
    leafs["callHistoryLogicalIfIndex"] = callhistoryentry.Callhistorylogicalifindex
    leafs["callHistoryDisconnectCause"] = callhistoryentry.Callhistorydisconnectcause
    leafs["callHistoryDisconnectText"] = callhistoryentry.Callhistorydisconnecttext
    leafs["callHistoryConnectTime"] = callhistoryentry.Callhistoryconnecttime
    leafs["callHistoryDisconnectTime"] = callhistoryentry.Callhistorydisconnecttime
    leafs["callHistoryCallOrigin"] = callhistoryentry.Callhistorycallorigin
    leafs["callHistoryChargedUnits"] = callhistoryentry.Callhistorychargedunits
    leafs["callHistoryInfoType"] = callhistoryentry.Callhistoryinfotype
    leafs["callHistoryTransmitPackets"] = callhistoryentry.Callhistorytransmitpackets
    leafs["callHistoryTransmitBytes"] = callhistoryentry.Callhistorytransmitbytes
    leafs["callHistoryReceivePackets"] = callhistoryentry.Callhistoryreceivepackets
    leafs["callHistoryReceiveBytes"] = callhistoryentry.Callhistoryreceivebytes
    return leafs
}

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetYangName() string { return "callHistoryEntry" }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) SetParent(parent types.Entity) { callhistoryentry.parent = parent }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetParent() types.Entity { return callhistoryentry.parent }

func (callhistoryentry *DIALCONTROLMIB_Callhistorytable_Callhistoryentry) GetParentYangName() string { return "callHistoryTable" }

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

