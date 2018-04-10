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

type Dot5Chipsetibm16 struct {
}

func (id Dot5Chipsetibm16) String() string {
	return "TOKENRING-MIB:dot5ChipSetIBM16"
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

// TOKENRINGMIB
type TOKENRINGMIB struct {
    EntityData types.CommonEntityData
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

func (tOKENRINGMIB *TOKENRINGMIB) GetEntityData() *types.CommonEntityData {
    tOKENRINGMIB.EntityData.YFilter = tOKENRINGMIB.YFilter
    tOKENRINGMIB.EntityData.YangName = "TOKENRING-MIB"
    tOKENRINGMIB.EntityData.BundleName = "cisco_ios_xe"
    tOKENRINGMIB.EntityData.ParentYangName = "TOKENRING-MIB"
    tOKENRINGMIB.EntityData.SegmentPath = "TOKENRING-MIB:TOKENRING-MIB"
    tOKENRINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tOKENRINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tOKENRINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tOKENRINGMIB.EntityData.Children = make(map[string]types.YChild)
    tOKENRINGMIB.EntityData.Children["dot5Table"] = types.YChild{"Dot5Table", &tOKENRINGMIB.Dot5Table}
    tOKENRINGMIB.EntityData.Children["dot5StatsTable"] = types.YChild{"Dot5Statstable", &tOKENRINGMIB.Dot5Statstable}
    tOKENRINGMIB.EntityData.Children["dot5TimerTable"] = types.YChild{"Dot5Timertable", &tOKENRINGMIB.Dot5Timertable}
    tOKENRINGMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tOKENRINGMIB.EntityData)
}

// TOKENRINGMIB_Dot5Table
// This table contains Token Ring interface
// parameters and state variables, one entry
// per 802.5 interface.
type TOKENRINGMIB_Dot5Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of Token Ring status and parameter values for an 802.5 interface.
    // The type is slice of TOKENRINGMIB_Dot5Table_Dot5Entry.
    Dot5Entry []TOKENRINGMIB_Dot5Table_Dot5Entry
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetEntityData() *types.CommonEntityData {
    dot5Table.EntityData.YFilter = dot5Table.YFilter
    dot5Table.EntityData.YangName = "dot5Table"
    dot5Table.EntityData.BundleName = "cisco_ios_xe"
    dot5Table.EntityData.ParentYangName = "TOKENRING-MIB"
    dot5Table.EntityData.SegmentPath = "dot5Table"
    dot5Table.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Table.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Table.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Table.EntityData.Children = make(map[string]types.YChild)
    dot5Table.EntityData.Children["dot5Entry"] = types.YChild{"Dot5Entry", nil}
    for i := range dot5Table.Dot5Entry {
        dot5Table.EntityData.Children[types.GetSegmentPath(&dot5Table.Dot5Entry[i])] = types.YChild{"Dot5Entry", &dot5Table.Dot5Entry[i]}
    }
    dot5Table.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot5Table.EntityData)
}

// TOKENRINGMIB_Dot5Table_Dot5Entry
// A list of Token Ring status and parameter
// values for an 802.5 interface.
type TOKENRINGMIB_Dot5Table_Dot5Entry struct {
    EntityData types.CommonEntityData
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
    // string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot5Upstream interface{}

    // If this object has a value of true(1) then this interface will participate
    // in the active monitor selection process.  If the value is false(2) then it
    // will not. Setting this object does not take effect until the next Active
    // Monitor election, and might not take effect until the next time the
    // interface is opened. The type is Dot5Actmonparticipate.
    Dot5Actmonparticipate interface{}

    // The bit mask of all Token Ring functional addresses for which this
    // interface will accept frames. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot5Functional interface{}

    // The value of MIB-II's sysUpTime object at which the local system last
    // transmitted a Beacon frame on this interface. The type is interface{} with
    // range: 0..4294967295.
    Dot5Lastbeaconsent interface{}
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetEntityData() *types.CommonEntityData {
    dot5Entry.EntityData.YFilter = dot5Entry.YFilter
    dot5Entry.EntityData.YangName = "dot5Entry"
    dot5Entry.EntityData.BundleName = "cisco_ios_xe"
    dot5Entry.EntityData.ParentYangName = "dot5Table"
    dot5Entry.EntityData.SegmentPath = "dot5Entry" + "[dot5IfIndex='" + fmt.Sprintf("%v", dot5Entry.Dot5Ifindex) + "']"
    dot5Entry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Entry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Entry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Entry.EntityData.Children = make(map[string]types.YChild)
    dot5Entry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot5Entry.EntityData.Leafs["dot5IfIndex"] = types.YLeaf{"Dot5Ifindex", dot5Entry.Dot5Ifindex}
    dot5Entry.EntityData.Leafs["dot5Commands"] = types.YLeaf{"Dot5Commands", dot5Entry.Dot5Commands}
    dot5Entry.EntityData.Leafs["dot5RingStatus"] = types.YLeaf{"Dot5Ringstatus", dot5Entry.Dot5Ringstatus}
    dot5Entry.EntityData.Leafs["dot5RingState"] = types.YLeaf{"Dot5Ringstate", dot5Entry.Dot5Ringstate}
    dot5Entry.EntityData.Leafs["dot5RingOpenStatus"] = types.YLeaf{"Dot5Ringopenstatus", dot5Entry.Dot5Ringopenstatus}
    dot5Entry.EntityData.Leafs["dot5RingSpeed"] = types.YLeaf{"Dot5Ringspeed", dot5Entry.Dot5Ringspeed}
    dot5Entry.EntityData.Leafs["dot5UpStream"] = types.YLeaf{"Dot5Upstream", dot5Entry.Dot5Upstream}
    dot5Entry.EntityData.Leafs["dot5ActMonParticipate"] = types.YLeaf{"Dot5Actmonparticipate", dot5Entry.Dot5Actmonparticipate}
    dot5Entry.EntityData.Leafs["dot5Functional"] = types.YLeaf{"Dot5Functional", dot5Entry.Dot5Functional}
    dot5Entry.EntityData.Leafs["dot5LastBeaconSent"] = types.YLeaf{"Dot5Lastbeaconsent", dot5Entry.Dot5Lastbeaconsent}
    return &(dot5Entry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry contains the 802.5 statistics for a particular interface. The type
    // is slice of TOKENRINGMIB_Dot5Statstable_Dot5Statsentry.
    Dot5Statsentry []TOKENRINGMIB_Dot5Statstable_Dot5Statsentry
}

func (dot5Statstable *TOKENRINGMIB_Dot5Statstable) GetEntityData() *types.CommonEntityData {
    dot5Statstable.EntityData.YFilter = dot5Statstable.YFilter
    dot5Statstable.EntityData.YangName = "dot5StatsTable"
    dot5Statstable.EntityData.BundleName = "cisco_ios_xe"
    dot5Statstable.EntityData.ParentYangName = "TOKENRING-MIB"
    dot5Statstable.EntityData.SegmentPath = "dot5StatsTable"
    dot5Statstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Statstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Statstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Statstable.EntityData.Children = make(map[string]types.YChild)
    dot5Statstable.EntityData.Children["dot5StatsEntry"] = types.YChild{"Dot5Statsentry", nil}
    for i := range dot5Statstable.Dot5Statsentry {
        dot5Statstable.EntityData.Children[types.GetSegmentPath(&dot5Statstable.Dot5Statsentry[i])] = types.YChild{"Dot5Statsentry", &dot5Statstable.Dot5Statsentry[i]}
    }
    dot5Statstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot5Statstable.EntityData)
}

// TOKENRINGMIB_Dot5Statstable_Dot5Statsentry
// An entry contains the 802.5 statistics
// for a particular interface.
type TOKENRINGMIB_Dot5Statstable_Dot5Statsentry struct {
    EntityData types.CommonEntityData
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

func (dot5Statsentry *TOKENRINGMIB_Dot5Statstable_Dot5Statsentry) GetEntityData() *types.CommonEntityData {
    dot5Statsentry.EntityData.YFilter = dot5Statsentry.YFilter
    dot5Statsentry.EntityData.YangName = "dot5StatsEntry"
    dot5Statsentry.EntityData.BundleName = "cisco_ios_xe"
    dot5Statsentry.EntityData.ParentYangName = "dot5StatsTable"
    dot5Statsentry.EntityData.SegmentPath = "dot5StatsEntry" + "[dot5StatsIfIndex='" + fmt.Sprintf("%v", dot5Statsentry.Dot5Statsifindex) + "']"
    dot5Statsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Statsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Statsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Statsentry.EntityData.Children = make(map[string]types.YChild)
    dot5Statsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot5Statsentry.EntityData.Leafs["dot5StatsIfIndex"] = types.YLeaf{"Dot5Statsifindex", dot5Statsentry.Dot5Statsifindex}
    dot5Statsentry.EntityData.Leafs["dot5StatsLineErrors"] = types.YLeaf{"Dot5Statslineerrors", dot5Statsentry.Dot5Statslineerrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsBurstErrors"] = types.YLeaf{"Dot5Statsbursterrors", dot5Statsentry.Dot5Statsbursterrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsACErrors"] = types.YLeaf{"Dot5Statsacerrors", dot5Statsentry.Dot5Statsacerrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsAbortTransErrors"] = types.YLeaf{"Dot5Statsaborttranserrors", dot5Statsentry.Dot5Statsaborttranserrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsInternalErrors"] = types.YLeaf{"Dot5Statsinternalerrors", dot5Statsentry.Dot5Statsinternalerrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsLostFrameErrors"] = types.YLeaf{"Dot5Statslostframeerrors", dot5Statsentry.Dot5Statslostframeerrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsReceiveCongestions"] = types.YLeaf{"Dot5Statsreceivecongestions", dot5Statsentry.Dot5Statsreceivecongestions}
    dot5Statsentry.EntityData.Leafs["dot5StatsFrameCopiedErrors"] = types.YLeaf{"Dot5Statsframecopiederrors", dot5Statsentry.Dot5Statsframecopiederrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsTokenErrors"] = types.YLeaf{"Dot5Statstokenerrors", dot5Statsentry.Dot5Statstokenerrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsSoftErrors"] = types.YLeaf{"Dot5Statssofterrors", dot5Statsentry.Dot5Statssofterrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsHardErrors"] = types.YLeaf{"Dot5Statsharderrors", dot5Statsentry.Dot5Statsharderrors}
    dot5Statsentry.EntityData.Leafs["dot5StatsSignalLoss"] = types.YLeaf{"Dot5Statssignalloss", dot5Statsentry.Dot5Statssignalloss}
    dot5Statsentry.EntityData.Leafs["dot5StatsTransmitBeacons"] = types.YLeaf{"Dot5Statstransmitbeacons", dot5Statsentry.Dot5Statstransmitbeacons}
    dot5Statsentry.EntityData.Leafs["dot5StatsRecoverys"] = types.YLeaf{"Dot5Statsrecoverys", dot5Statsentry.Dot5Statsrecoverys}
    dot5Statsentry.EntityData.Leafs["dot5StatsLobeWires"] = types.YLeaf{"Dot5Statslobewires", dot5Statsentry.Dot5Statslobewires}
    dot5Statsentry.EntityData.Leafs["dot5StatsRemoves"] = types.YLeaf{"Dot5Statsremoves", dot5Statsentry.Dot5Statsremoves}
    dot5Statsentry.EntityData.Leafs["dot5StatsSingles"] = types.YLeaf{"Dot5Statssingles", dot5Statsentry.Dot5Statssingles}
    dot5Statsentry.EntityData.Leafs["dot5StatsFreqErrors"] = types.YLeaf{"Dot5Statsfreqerrors", dot5Statsentry.Dot5Statsfreqerrors}
    return &(dot5Statsentry.EntityData)
}

// TOKENRINGMIB_Dot5Timertable
// This table contains Token Ring interface
// timer values, one entry per 802.5
// interface.
type TOKENRINGMIB_Dot5Timertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of Token Ring timer values for an 802.5 interface. The type is slice
    // of TOKENRINGMIB_Dot5Timertable_Dot5Timerentry.
    Dot5Timerentry []TOKENRINGMIB_Dot5Timertable_Dot5Timerentry
}

func (dot5Timertable *TOKENRINGMIB_Dot5Timertable) GetEntityData() *types.CommonEntityData {
    dot5Timertable.EntityData.YFilter = dot5Timertable.YFilter
    dot5Timertable.EntityData.YangName = "dot5TimerTable"
    dot5Timertable.EntityData.BundleName = "cisco_ios_xe"
    dot5Timertable.EntityData.ParentYangName = "TOKENRING-MIB"
    dot5Timertable.EntityData.SegmentPath = "dot5TimerTable"
    dot5Timertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Timertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Timertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Timertable.EntityData.Children = make(map[string]types.YChild)
    dot5Timertable.EntityData.Children["dot5TimerEntry"] = types.YChild{"Dot5Timerentry", nil}
    for i := range dot5Timertable.Dot5Timerentry {
        dot5Timertable.EntityData.Children[types.GetSegmentPath(&dot5Timertable.Dot5Timerentry[i])] = types.YChild{"Dot5Timerentry", &dot5Timertable.Dot5Timerentry[i]}
    }
    dot5Timertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot5Timertable.EntityData)
}

// TOKENRINGMIB_Dot5Timertable_Dot5Timerentry
// A list of Token Ring timer values for an
// 802.5 interface.
type TOKENRINGMIB_Dot5Timertable_Dot5Timerentry struct {
    EntityData types.CommonEntityData
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

func (dot5Timerentry *TOKENRINGMIB_Dot5Timertable_Dot5Timerentry) GetEntityData() *types.CommonEntityData {
    dot5Timerentry.EntityData.YFilter = dot5Timerentry.YFilter
    dot5Timerentry.EntityData.YangName = "dot5TimerEntry"
    dot5Timerentry.EntityData.BundleName = "cisco_ios_xe"
    dot5Timerentry.EntityData.ParentYangName = "dot5TimerTable"
    dot5Timerentry.EntityData.SegmentPath = "dot5TimerEntry" + "[dot5TimerIfIndex='" + fmt.Sprintf("%v", dot5Timerentry.Dot5Timerifindex) + "']"
    dot5Timerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Timerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Timerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Timerentry.EntityData.Children = make(map[string]types.YChild)
    dot5Timerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot5Timerentry.EntityData.Leafs["dot5TimerIfIndex"] = types.YLeaf{"Dot5Timerifindex", dot5Timerentry.Dot5Timerifindex}
    dot5Timerentry.EntityData.Leafs["dot5TimerReturnRepeat"] = types.YLeaf{"Dot5Timerreturnrepeat", dot5Timerentry.Dot5Timerreturnrepeat}
    dot5Timerentry.EntityData.Leafs["dot5TimerHolding"] = types.YLeaf{"Dot5Timerholding", dot5Timerentry.Dot5Timerholding}
    dot5Timerentry.EntityData.Leafs["dot5TimerQueuePDU"] = types.YLeaf{"Dot5Timerqueuepdu", dot5Timerentry.Dot5Timerqueuepdu}
    dot5Timerentry.EntityData.Leafs["dot5TimerValidTransmit"] = types.YLeaf{"Dot5Timervalidtransmit", dot5Timerentry.Dot5Timervalidtransmit}
    dot5Timerentry.EntityData.Leafs["dot5TimerNoToken"] = types.YLeaf{"Dot5Timernotoken", dot5Timerentry.Dot5Timernotoken}
    dot5Timerentry.EntityData.Leafs["dot5TimerActiveMon"] = types.YLeaf{"Dot5Timeractivemon", dot5Timerentry.Dot5Timeractivemon}
    dot5Timerentry.EntityData.Leafs["dot5TimerStandbyMon"] = types.YLeaf{"Dot5Timerstandbymon", dot5Timerentry.Dot5Timerstandbymon}
    dot5Timerentry.EntityData.Leafs["dot5TimerErrorReport"] = types.YLeaf{"Dot5Timererrorreport", dot5Timerentry.Dot5Timererrorreport}
    dot5Timerentry.EntityData.Leafs["dot5TimerBeaconTransmit"] = types.YLeaf{"Dot5Timerbeacontransmit", dot5Timerentry.Dot5Timerbeacontransmit}
    dot5Timerentry.EntityData.Leafs["dot5TimerBeaconReceive"] = types.YLeaf{"Dot5Timerbeaconreceive", dot5Timerentry.Dot5Timerbeaconreceive}
    return &(dot5Timerentry.EntityData)
}

