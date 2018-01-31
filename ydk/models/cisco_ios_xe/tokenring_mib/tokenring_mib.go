// The MIB module for IEEE Token Ring entities.
package tokenring_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tokenring_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:TOKENRING-MIB TOKENRING-MIB}", reflect.TypeOf(TOKENRINGMIB{}))
    ydk.RegisterEntity("TOKENRING-MIB:TOKENRING-MIB", reflect.TypeOf(TOKENRINGMIB{}))
}

type Dot5Chipsettitms380 struct {
}

func (id Dot5Chipsettitms380) String() string {
	return "TOKENRING-MIB:dot5ChipSetTItms380"
}

type Dot5Chipsettitms380C16 struct {
}

func (id Dot5Chipsettitms380C16) String() string {
	return "TOKENRING-MIB:dot5ChipSetTItms380c16"
}

type Dot5Chipsetibm16 struct {
}

func (id Dot5Chipsetibm16) String() string {
	return "TOKENRING-MIB:dot5ChipSetIBM16"
}

type Dot5Testinsertfunc struct {
}

func (id Dot5Testinsertfunc) String() string {
	return "TOKENRING-MIB:dot5TestInsertFunc"
}

type Dot5Testfullduplexloopback struct {
}

func (id Dot5Testfullduplexloopback) String() string {
	return "TOKENRING-MIB:dot5TestFullDuplexLoopBack"
}

// TOKENRINGMIB
type TOKENRINGMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table contains Token Ring interface parameters and state variables,
    // one entry per 802.5 interface.
    Dot5Table TOKENRINGMIB_Dot5Table

    // A table containing Token Ring statistics, one entry per 802.5 interface.   
    // All the statistics are defined using the syntax Counter32 as 32-bit wrap
    // around counters.  Thus, if an interface's hardware maintains these
    // statistics in 16-bit counters, then the agent must read the hardware's
    // counters frequently enough to prevent loss of significance, in order to
    // maintain 32-bit counters in software.
    Dot5Statstable TOKENRINGMIB_Dot5Statstable

    // This table contains Token Ring interface timer values, one entry per 802.5
    // interface.
    Dot5Timertable TOKENRINGMIB_Dot5Timertable
}

func (tOKENRINGMIB *TOKENRINGMIB) GetFilter() yfilter.YFilter { return tOKENRINGMIB.YFilter }

func (tOKENRINGMIB *TOKENRINGMIB) SetFilter(yf yfilter.YFilter) { tOKENRINGMIB.YFilter = yf }

func (tOKENRINGMIB *TOKENRINGMIB) GetGoName(yname string) string {
    if yname == "dot5Table" { return "Dot5Table" }
    if yname == "dot5StatsTable" { return "Dot5Statstable" }
    if yname == "dot5TimerTable" { return "Dot5Timertable" }
    return ""
}

func (tOKENRINGMIB *TOKENRINGMIB) GetSegmentPath() string {
    return "TOKENRING-MIB:TOKENRING-MIB"
}

func (tOKENRINGMIB *TOKENRINGMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot5Table" {
        return &tOKENRINGMIB.Dot5Table
    }
    if childYangName == "dot5StatsTable" {
        return &tOKENRINGMIB.Dot5Statstable
    }
    if childYangName == "dot5TimerTable" {
        return &tOKENRINGMIB.Dot5Timertable
    }
    return nil
}

func (tOKENRINGMIB *TOKENRINGMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dot5Table"] = &tOKENRINGMIB.Dot5Table
    children["dot5StatsTable"] = &tOKENRINGMIB.Dot5Statstable
    children["dot5TimerTable"] = &tOKENRINGMIB.Dot5Timertable
    return children
}

func (tOKENRINGMIB *TOKENRINGMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tOKENRINGMIB *TOKENRINGMIB) GetBundleName() string { return "cisco_ios_xe" }

func (tOKENRINGMIB *TOKENRINGMIB) GetYangName() string { return "TOKENRING-MIB" }

func (tOKENRINGMIB *TOKENRINGMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tOKENRINGMIB *TOKENRINGMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tOKENRINGMIB *TOKENRINGMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tOKENRINGMIB *TOKENRINGMIB) SetParent(parent types.Entity) { tOKENRINGMIB.parent = parent }

func (tOKENRINGMIB *TOKENRINGMIB) GetParent() types.Entity { return tOKENRINGMIB.parent }

func (tOKENRINGMIB *TOKENRINGMIB) GetParentYangName() string { return "TOKENRING-MIB" }

// TOKENRINGMIB_Dot5Table
// This table contains Token Ring interface
// parameters and state variables, one entry
// per 802.5 interface.
type TOKENRINGMIB_Dot5Table struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of Token Ring status and parameter values for an 802.5 interface.
    // The type is slice of TOKENRINGMIB_Dot5Table_Dot5Entry.
    Dot5Entry []TOKENRINGMIB_Dot5Table_Dot5Entry
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetFilter() yfilter.YFilter { return dot5Table.YFilter }

func (dot5Table *TOKENRINGMIB_Dot5Table) SetFilter(yf yfilter.YFilter) { dot5Table.YFilter = yf }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetGoName(yname string) string {
    if yname == "dot5Entry" { return "Dot5Entry" }
    return ""
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetSegmentPath() string {
    return "dot5Table"
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot5Entry" {
        for _, c := range dot5Table.Dot5Entry {
            if dot5Table.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGMIB_Dot5Table_Dot5Entry{}
        dot5Table.Dot5Entry = append(dot5Table.Dot5Entry, child)
        return &dot5Table.Dot5Entry[len(dot5Table.Dot5Entry)-1]
    }
    return nil
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot5Table.Dot5Entry {
        children[dot5Table.Dot5Entry[i].GetSegmentPath()] = &dot5Table.Dot5Entry[i]
    }
    return children
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetBundleName() string { return "cisco_ios_xe" }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetYangName() string { return "dot5Table" }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot5Table *TOKENRINGMIB_Dot5Table) SetParent(parent types.Entity) { dot5Table.parent = parent }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetParent() types.Entity { return dot5Table.parent }

func (dot5Table *TOKENRINGMIB_Dot5Table) GetParentYangName() string { return "TOKENRING-MIB" }

// TOKENRINGMIB_Dot5Table_Dot5Entry
// A list of Token Ring status and parameter
// values for an 802.5 interface.
type TOKENRINGMIB_Dot5Table_Dot5Entry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object identifies the 802.5
    // interface for which this entry contains management information.  The value
    // of this object for a particular interface has the same value as the ifIndex
    // object, defined in MIB-II for the same interface. The type is interface{}
    // with range: -2147483648..2147483647.
    Dot5Ifindex interface{}

    // When this object is set to the value of open(2), the station should go into
    // the open state.  The progress and success of the open is given by the
    // values of the objects dot5RingState and dot5RingOpenStatus.     When this
    // object is set to the value of reset(3), then the station should do a reset.
    // On a reset, all MIB counters should retain their values, if possible. Other
    // side affects are dependent on the hardware chip set.     When this object
    // is set to the value of close(4), the station should go into the stopped
    // state by removing itself from the ring.     Setting this object to a value
    // of noop(1) has no effect.     When read, this object always has a value of
    // noop(1).     The open(2) and close(4) values correspond to the up(1) and
    // down(2) values of MIB-II's ifAdminStatus and ifOperStatus, i.e., the
    // setting of ifAdminStatus and   dot5Commands affects the values of both
    // dot5Commands and ifOperStatus. The type is Dot5Commands.
    Dot5Commands interface{}

    // The current interface status which can be used to diagnose fluctuating
    // problems that can occur on token rings, after a station has successfully
    // been added to the ring.    Before an open is completed, this object has the
    // value for the 'no status' condition.  The dot5RingState and
    // dot5RingOpenStatus objects provide for debugging problems when the station
    // can not even enter the ring.     The object's value is a sum of values, one
    // for each currently applicable condition.  The following values are defined
    // for various conditions:          0 = No Problems detected        32 = Ring
    // Recovery        64 = Single Station       256 = Remove Received       512 =
    // reserved      1024 = Auto-Removal Error      2048 = Lobe Wire Fault     
    // 4096 = Transmit Beacon      8192 = Soft Error     16384 = Hard Error    
    // 32768 = Signal Loss    131072 = no status, open not completed. The type is
    // interface{} with range: 0..262143.
    Dot5Ringstatus interface{}

    // The current interface state with respect to entering or leaving the ring.
    // The type is Dot5Ringstate.
    Dot5Ringstate interface{}

    // This object indicates the success, or the reason for failure, of the
    // station's most recent attempt to enter the ring. The type is
    // Dot5Ringopenstatus.
    Dot5Ringopenstatus interface{}

    // The ring-speed at the next insertion into the ring.  Note that this may or
    // may not be different to the current ring-speed which is given by MIB-II's
    // ifSpeed.  For interfaces which do not support changing ring-speed,
    // dot5RingSpeed can only be set to its current value.  When dot5RingSpeed has
    // the value unknown(1), the ring's actual ring-speed is to be used. The type
    // is Dot5Ringspeed.
    Dot5Ringspeed interface{}

    // The MAC-address of the up stream neighbor station in the ring. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot5Upstream interface{}

    // If this object has a value of true(1) then this interface will participate
    // in the active monitor selection process.  If the value is false(2) then it
    // will not. Setting this object does not take effect until the next Active
    // Monitor election, and might not take effect until the next time the
    // interface is opened. The type is Dot5Actmonparticipate.
    Dot5Actmonparticipate interface{}

    // The bit mask of all Token Ring functional addresses for which this
    // interface will accept frames. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot5Functional interface{}

    // The value of MIB-II's sysUpTime object at which the local system last
    // transmitted a Beacon frame on this interface. The type is interface{} with
    // range: 0..4294967295.
    Dot5Lastbeaconsent interface{}
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetFilter() yfilter.YFilter { return dot5Entry.YFilter }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) SetFilter(yf yfilter.YFilter) { dot5Entry.YFilter = yf }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetGoName(yname string) string {
    if yname == "dot5IfIndex" { return "Dot5Ifindex" }
    if yname == "dot5Commands" { return "Dot5Commands" }
    if yname == "dot5RingStatus" { return "Dot5Ringstatus" }
    if yname == "dot5RingState" { return "Dot5Ringstate" }
    if yname == "dot5RingOpenStatus" { return "Dot5Ringopenstatus" }
    if yname == "dot5RingSpeed" { return "Dot5Ringspeed" }
    if yname == "dot5UpStream" { return "Dot5Upstream" }
    if yname == "dot5ActMonParticipate" { return "Dot5Actmonparticipate" }
    if yname == "dot5Functional" { return "Dot5Functional" }
    if yname == "dot5LastBeaconSent" { return "Dot5Lastbeaconsent" }
    return ""
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetSegmentPath() string {
    return "dot5Entry" + "[dot5IfIndex='" + fmt.Sprintf("%v", dot5Entry.Dot5Ifindex) + "']"
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot5IfIndex"] = dot5Entry.Dot5Ifindex
    leafs["dot5Commands"] = dot5Entry.Dot5Commands
    leafs["dot5RingStatus"] = dot5Entry.Dot5Ringstatus
    leafs["dot5RingState"] = dot5Entry.Dot5Ringstate
    leafs["dot5RingOpenStatus"] = dot5Entry.Dot5Ringopenstatus
    leafs["dot5RingSpeed"] = dot5Entry.Dot5Ringspeed
    leafs["dot5UpStream"] = dot5Entry.Dot5Upstream
    leafs["dot5ActMonParticipate"] = dot5Entry.Dot5Actmonparticipate
    leafs["dot5Functional"] = dot5Entry.Dot5Functional
    leafs["dot5LastBeaconSent"] = dot5Entry.Dot5Lastbeaconsent
    return leafs
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetBundleName() string { return "cisco_ios_xe" }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetYangName() string { return "dot5Entry" }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) SetParent(parent types.Entity) { dot5Entry.parent = parent }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetParent() types.Entity { return dot5Entry.parent }

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetParentYangName() string { return "dot5Table" }

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Actmonparticipate represents the interface is opened.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Actmonparticipate string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Actmonparticipate_true TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Actmonparticipate = "true"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Actmonparticipate_false TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Actmonparticipate = "false"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands represents dot5Commands and ifOperStatus.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_noop TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "noop"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_open TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "open"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_reset TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "reset"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_close TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "close"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus represents recent attempt to enter the ring.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_noOpen TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "noOpen"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_badParam TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "badParam"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_lobeFailed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "lobeFailed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_signalLoss TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "signalLoss"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_insertionTimeout TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "insertionTimeout"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_ringFailed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "ringFailed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_beaconing TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "beaconing"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_duplicateMAC TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "duplicateMAC"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_requestFailed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "requestFailed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_removeReceived TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "removeReceived"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus_open TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringopenstatus = "open"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed represents to be used.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed_unknown TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed = "unknown"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed_oneMegabit TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed = "oneMegabit"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed_fourMegabit TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed = "fourMegabit"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed_sixteenMegabit TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringspeed = "sixteenMegabit"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate represents to entering or leaving the ring.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate_opened TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate = "opened"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate_closed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate = "closed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate_opening TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate = "opening"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate_closing TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate = "closing"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate_openFailure TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate = "openFailure"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate_ringFailure TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Ringstate = "ringFailure"
)

// TOKENRINGMIB_Dot5Statstable
// A table containing Token Ring statistics,
// one entry per 802.5 interface.
//     All the statistics are defined using
// the syntax Counter32 as 32-bit wrap around
// counters.  Thus, if an interface's
// hardware maintains these statistics in
// 16-bit counters, then the agent must read
// the hardware's counters frequently enough
// to prevent loss of significance, in order
// to maintain 32-bit counters in software.
type TOKENRINGMIB_Dot5Statstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry contains the 802.5 statistics for a particular interface. The type
    // is slice of TOKENRINGMIB_Dot5Statstable_Dot5Statsentry.
    Dot5Statsentry []TOKENRINGMIB_Dot5Statstable_Dot5Statsentry
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetFilter() yfilter.YFilter { return dot5Statstable.YFilter }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) SetFilter(yf yfilter.YFilter) { dot5Statstable.YFilter = yf }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetGoName(yname string) string {
    if yname == "dot5StatsEntry" { return "Dot5Statsentry" }
    return ""
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetSegmentPath() string {
    return "dot5StatsTable"
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot5StatsEntry" {
        for _, c := range dot5Statstable.Dot5Statsentry {
            if dot5Statstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGMIB_Dot5Statstable_Dot5Statsentry{}
        dot5Statstable.Dot5Statsentry = append(dot5Statstable.Dot5Statsentry, child)
        return &dot5Statstable.Dot5Statsentry[len(dot5Statstable.Dot5Statsentry)-1]
    }
    return nil
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot5Statstable.Dot5Statsentry {
        children[dot5Statstable.Dot5Statsentry[i].GetSegmentPath()] = &dot5Statstable.Dot5Statsentry[i]
    }
    return children
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetYangName() string { return "dot5StatsTable" }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) SetParent(parent types.Entity) { dot5Statstable.parent = parent }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetParent() types.Entity { return dot5Statstable.parent }

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetParentYangName() string { return "TOKENRING-MIB" }

// TOKENRINGMIB_Dot5Statstable_Dot5Statsentry
// An entry contains the 802.5 statistics
// for a particular interface.
type TOKENRINGMIB_Dot5Statstable_Dot5Statsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object identifies the 802.5
    // interface for which this entry contains management information.  The value
    // of this object for a particular interface has the same value as MIB-II's
    // ifIndex object for the same interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5Statsifindex interface{}

    // This counter is incremented when a frame or token is copied or repeated by
    // a station, the E bit is zero in the frame or token and one of the following
    // conditions exists: 1) there is a non-data bit (J or K bit) between the SD
    // and the ED of the frame or token, or 2) there is an FCS error in the frame.
    // The type is interface{} with range: 0..4294967295.
    Dot5Statslineerrors interface{}

    // This counter is incremented when a station detects the absence of
    // transitions for five half-bit timers (burst-five error). The type is
    // interface{} with range: 0..4294967295.
    Dot5Statsbursterrors interface{}

    // This counter is incremented when a station receives an AMP or SMP frame in
    // which A is equal to C is equal to 0, and then receives another SMP frame
    // with A is equal to C is equal to 0 without first receiving an AMP frame. It
    // denotes a station that cannot set the AC bits properly. The type is
    // interface{} with range: 0..4294967295.
    Dot5Statsacerrors interface{}

    // This counter is incremented when a station transmits an abort delimiter
    // while transmitting. The type is interface{} with range: 0..4294967295.
    Dot5Statsaborttranserrors interface{}

    // This counter is incremented when a station recognizes an internal error.
    // The type is interface{} with range: 0..4294967295.
    Dot5Statsinternalerrors interface{}

    // This counter is incremented when a station is transmitting and its TRR
    // timer expires. This condition denotes a condition where a transmitting
    // station in strip mode does not receive the trailer of the frame before the
    // TRR timer goes off. The type is interface{} with range: 0..4294967295.
    Dot5Statslostframeerrors interface{}

    // This counter is incremented when a station recognizes a frame addressed to
    // its specific address, but has no available buffer space indicating that the
    // station is congested. The type is interface{} with range: 0..4294967295.
    Dot5Statsreceivecongestions interface{}

    // This counter is incremented when a station recognizes a frame addressed to
    // its specific address and detects that the FS field A bits are set to 1
    // indicating a possible line hit or duplicate address. The type is
    // interface{} with range: 0..4294967295.
    Dot5Statsframecopiederrors interface{}

    // This counter is incremented when a station acting as the active monitor
    // recognizes an error condition that needs a token transmitted. The type is
    // interface{} with range: 0..4294967295.
    Dot5Statstokenerrors interface{}

    // The number of Soft Errors the interface has detected. It directly
    // corresponds to the number of Report Error MAC frames that this interface
    // has transmitted. Soft Errors are those which are recoverable by the MAC
    // layer protocols. The type is interface{} with range: 0..4294967295.
    Dot5Statssofterrors interface{}

    // The number of times this interface has detected an immediately recoverable
    // fatal error.  It denotes the number of times this interface is either
    // transmitting or receiving beacon MAC frames. The type is interface{} with
    // range: 0..4294967295.
    Dot5Statsharderrors interface{}

    // The number of times this interface has detected the loss of signal
    // condition from the ring. The type is interface{} with range: 0..4294967295.
    Dot5Statssignalloss interface{}

    // The number of times this interface has transmitted a beacon frame. The type
    // is interface{} with range: 0..4294967295.
    Dot5Statstransmitbeacons interface{}

    // The number of Claim Token MAC frames received or transmitted after the
    // interface has received a Ring Purge MAC frame.  This counter signifies the
    // number of times the ring has been purged and is being recovered back into a
    // normal operating state. The type is interface{} with range: 0..4294967295.
    Dot5Statsrecoverys interface{}

    // The number of times the interface has detected an open or short circuit in
    // the lobe data path.  The adapter will be closed and dot5RingState will
    // signify this condition. The type is interface{} with range: 0..4294967295.
    Dot5Statslobewires interface{}

    // The number of times the interface has received a Remove Ring Station MAC
    // frame request.  When this frame is received the interface will enter the
    // close state and dot5RingState will signify this condition. The type is
    // interface{} with range: 0..4294967295.
    Dot5Statsremoves interface{}

    // The number of times the interface has sensed that it is the only station on
    // the ring.  This will happen if the interface is the first one up on a ring,
    // or if there is a hardware problem. The type is interface{} with range:
    // 0..4294967295.
    Dot5Statssingles interface{}

    // The number of times the interface has detected that the frequency of the
    // incoming signal differs from the expected frequency by more than that
    // specified by the IEEE 802.5 standard. The type is interface{} with range:
    // 0..4294967295.
    Dot5Statsfreqerrors interface{}
}

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetFilter() yfilter.YFilter { return dot5Statsentry.YFilter }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) SetFilter(yf yfilter.YFilter) { dot5Statsentry.YFilter = yf }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetGoName(yname string) string {
    if yname == "dot5StatsIfIndex" { return "Dot5Statsifindex" }
    if yname == "dot5StatsLineErrors" { return "Dot5Statslineerrors" }
    if yname == "dot5StatsBurstErrors" { return "Dot5Statsbursterrors" }
    if yname == "dot5StatsACErrors" { return "Dot5Statsacerrors" }
    if yname == "dot5StatsAbortTransErrors" { return "Dot5Statsaborttranserrors" }
    if yname == "dot5StatsInternalErrors" { return "Dot5Statsinternalerrors" }
    if yname == "dot5StatsLostFrameErrors" { return "Dot5Statslostframeerrors" }
    if yname == "dot5StatsReceiveCongestions" { return "Dot5Statsreceivecongestions" }
    if yname == "dot5StatsFrameCopiedErrors" { return "Dot5Statsframecopiederrors" }
    if yname == "dot5StatsTokenErrors" { return "Dot5Statstokenerrors" }
    if yname == "dot5StatsSoftErrors" { return "Dot5Statssofterrors" }
    if yname == "dot5StatsHardErrors" { return "Dot5Statsharderrors" }
    if yname == "dot5StatsSignalLoss" { return "Dot5Statssignalloss" }
    if yname == "dot5StatsTransmitBeacons" { return "Dot5Statstransmitbeacons" }
    if yname == "dot5StatsRecoverys" { return "Dot5Statsrecoverys" }
    if yname == "dot5StatsLobeWires" { return "Dot5Statslobewires" }
    if yname == "dot5StatsRemoves" { return "Dot5Statsremoves" }
    if yname == "dot5StatsSingles" { return "Dot5Statssingles" }
    if yname == "dot5StatsFreqErrors" { return "Dot5Statsfreqerrors" }
    return ""
}

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetSegmentPath() string {
    return "dot5StatsEntry" + "[dot5StatsIfIndex='" + fmt.Sprintf("%v", dot5Statsentry.Dot5Statsifindex) + "']"
}

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot5StatsIfIndex"] = dot5Statsentry.Dot5Statsifindex
    leafs["dot5StatsLineErrors"] = dot5Statsentry.Dot5Statslineerrors
    leafs["dot5StatsBurstErrors"] = dot5Statsentry.Dot5Statsbursterrors
    leafs["dot5StatsACErrors"] = dot5Statsentry.Dot5Statsacerrors
    leafs["dot5StatsAbortTransErrors"] = dot5Statsentry.Dot5Statsaborttranserrors
    leafs["dot5StatsInternalErrors"] = dot5Statsentry.Dot5Statsinternalerrors
    leafs["dot5StatsLostFrameErrors"] = dot5Statsentry.Dot5Statslostframeerrors
    leafs["dot5StatsReceiveCongestions"] = dot5Statsentry.Dot5Statsreceivecongestions
    leafs["dot5StatsFrameCopiedErrors"] = dot5Statsentry.Dot5Statsframecopiederrors
    leafs["dot5StatsTokenErrors"] = dot5Statsentry.Dot5Statstokenerrors
    leafs["dot5StatsSoftErrors"] = dot5Statsentry.Dot5Statssofterrors
    leafs["dot5StatsHardErrors"] = dot5Statsentry.Dot5Statsharderrors
    leafs["dot5StatsSignalLoss"] = dot5Statsentry.Dot5Statssignalloss
    leafs["dot5StatsTransmitBeacons"] = dot5Statsentry.Dot5Statstransmitbeacons
    leafs["dot5StatsRecoverys"] = dot5Statsentry.Dot5Statsrecoverys
    leafs["dot5StatsLobeWires"] = dot5Statsentry.Dot5Statslobewires
    leafs["dot5StatsRemoves"] = dot5Statsentry.Dot5Statsremoves
    leafs["dot5StatsSingles"] = dot5Statsentry.Dot5Statssingles
    leafs["dot5StatsFreqErrors"] = dot5Statsentry.Dot5Statsfreqerrors
    return leafs
}

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetYangName() string { return "dot5StatsEntry" }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) SetParent(parent types.Entity) { dot5Statsentry.parent = parent }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetParent() types.Entity { return dot5Statsentry.parent }

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetParentYangName() string { return "dot5StatsTable" }

// TOKENRINGMIB_Dot5Timertable
// This table contains Token Ring interface
// timer values, one entry per 802.5
// interface.
type TOKENRINGMIB_Dot5Timertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of Token Ring timer values for an 802.5 interface. The type is slice
    // of TOKENRINGMIB_Dot5Timertable_Dot5Timerentry.
    Dot5Timerentry []TOKENRINGMIB_Dot5Timertable_Dot5Timerentry
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetFilter() yfilter.YFilter { return dot5Timertable.YFilter }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) SetFilter(yf yfilter.YFilter) { dot5Timertable.YFilter = yf }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetGoName(yname string) string {
    if yname == "dot5TimerEntry" { return "Dot5Timerentry" }
    return ""
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetSegmentPath() string {
    return "dot5TimerTable"
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot5TimerEntry" {
        for _, c := range dot5Timertable.Dot5Timerentry {
            if dot5Timertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TOKENRINGMIB_Dot5Timertable_Dot5Timerentry{}
        dot5Timertable.Dot5Timerentry = append(dot5Timertable.Dot5Timerentry, child)
        return &dot5Timertable.Dot5Timerentry[len(dot5Timertable.Dot5Timerentry)-1]
    }
    return nil
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot5Timertable.Dot5Timerentry {
        children[dot5Timertable.Dot5Timerentry[i].GetSegmentPath()] = &dot5Timertable.Dot5Timerentry[i]
    }
    return children
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetBundleName() string { return "cisco_ios_xe" }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetYangName() string { return "dot5TimerTable" }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) SetParent(parent types.Entity) { dot5Timertable.parent = parent }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetParent() types.Entity { return dot5Timertable.parent }

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetParentYangName() string { return "TOKENRING-MIB" }

// TOKENRINGMIB_Dot5Timertable_Dot5Timerentry
// A list of Token Ring timer values for an
// 802.5 interface.
type TOKENRINGMIB_Dot5Timertable_Dot5Timerentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object identifies the 802.5
    // interface for which this entry contains timer values.  The value of   this
    // object for a particular interface has the same value as MIB-II's ifIndex
    // object for the same interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5Timerifindex interface{}

    // The time-out value used to ensure the interface will return to Repeat
    // State, in units of 100 micro-seconds.  The value should be greater than the
    // maximum ring latency. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5Timerreturnrepeat interface{}

    // Maximum period of time a station is permitted to transmit frames after
    // capturing a token, in units of 100 micro-seconds. The type is interface{}
    // with range: -2147483648..2147483647.
    Dot5Timerholding interface{}

    // The time-out value for enqueuing of an SMP PDU after reception of an AMP or
    // SMP frame in which the A and C bits were equal to 0, in units of 100
    // micro-seconds. The type is interface{} with range: -2147483648..2147483647.
    Dot5Timerqueuepdu interface{}

    // The time-out value used by the active monitor to detect the absence of
    // valid transmissions, in units of 100 micro-seconds. The type is interface{}
    // with range: -2147483648..2147483647.
    Dot5Timervalidtransmit interface{}

    // The time-out value used to recover from various-related error situations.
    // If N is the maximum number of stations on the ring, the value of this timer
    // is normally: dot5TimerReturnRepeat + N*dot5TimerHolding. The type is
    // interface{} with range: -2147483648..2147483647.
    Dot5Timernotoken interface{}

    // The time-out value used by the active monitor to stimulate the enqueuing of
    // an AMP PDU for transmission, in units of 100 micro-seconds. The type is
    // interface{} with range: -2147483648..2147483647.
    Dot5Timeractivemon interface{}

    // The time-out value used by the stand-by monitors to ensure that there is an
    // active monitor on the ring and to detect a continuous stream of tokens, in
    // units of 100 micro-seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5Timerstandbymon interface{}

    // The time-out value which determines how often a station shall send a Report
    // Error MAC frame to report its error counters, in units of 100
    // micro-seconds. The type is interface{} with range: -2147483648..2147483647.
    Dot5Timererrorreport interface{}

    // The time-out value which determines how long a station shall remain in the
    // state of transmitting Beacon frames before entering the Bypass state, in
    // units of 100 micro-seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5Timerbeacontransmit interface{}

    // The time-out value which determines how long a station shall receive Beacon
    // frames from its downstream neighbor before entering the Bypass state, in
    // units of 100 micro-seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5Timerbeaconreceive interface{}
}

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetFilter() yfilter.YFilter { return dot5Timerentry.YFilter }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) SetFilter(yf yfilter.YFilter) { dot5Timerentry.YFilter = yf }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetGoName(yname string) string {
    if yname == "dot5TimerIfIndex" { return "Dot5Timerifindex" }
    if yname == "dot5TimerReturnRepeat" { return "Dot5Timerreturnrepeat" }
    if yname == "dot5TimerHolding" { return "Dot5Timerholding" }
    if yname == "dot5TimerQueuePDU" { return "Dot5Timerqueuepdu" }
    if yname == "dot5TimerValidTransmit" { return "Dot5Timervalidtransmit" }
    if yname == "dot5TimerNoToken" { return "Dot5Timernotoken" }
    if yname == "dot5TimerActiveMon" { return "Dot5Timeractivemon" }
    if yname == "dot5TimerStandbyMon" { return "Dot5Timerstandbymon" }
    if yname == "dot5TimerErrorReport" { return "Dot5Timererrorreport" }
    if yname == "dot5TimerBeaconTransmit" { return "Dot5Timerbeacontransmit" }
    if yname == "dot5TimerBeaconReceive" { return "Dot5Timerbeaconreceive" }
    return ""
}

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetSegmentPath() string {
    return "dot5TimerEntry" + "[dot5TimerIfIndex='" + fmt.Sprintf("%v", dot5Timerentry.Dot5Timerifindex) + "']"
}

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot5TimerIfIndex"] = dot5Timerentry.Dot5Timerifindex
    leafs["dot5TimerReturnRepeat"] = dot5Timerentry.Dot5Timerreturnrepeat
    leafs["dot5TimerHolding"] = dot5Timerentry.Dot5Timerholding
    leafs["dot5TimerQueuePDU"] = dot5Timerentry.Dot5Timerqueuepdu
    leafs["dot5TimerValidTransmit"] = dot5Timerentry.Dot5Timervalidtransmit
    leafs["dot5TimerNoToken"] = dot5Timerentry.Dot5Timernotoken
    leafs["dot5TimerActiveMon"] = dot5Timerentry.Dot5Timeractivemon
    leafs["dot5TimerStandbyMon"] = dot5Timerentry.Dot5Timerstandbymon
    leafs["dot5TimerErrorReport"] = dot5Timerentry.Dot5Timererrorreport
    leafs["dot5TimerBeaconTransmit"] = dot5Timerentry.Dot5Timerbeacontransmit
    leafs["dot5TimerBeaconReceive"] = dot5Timerentry.Dot5Timerbeaconreceive
    return leafs
}

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetYangName() string { return "dot5TimerEntry" }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) SetParent(parent types.Entity) { dot5Timerentry.parent = parent }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetParent() types.Entity { return dot5Timerentry.parent }

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetParentYangName() string { return "dot5TimerTable" }

