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

type Dot5TestInsertFunc struct {
}

func (id Dot5TestInsertFunc) String() string {
	return "TOKENRING-MIB:dot5TestInsertFunc"
}

type Dot5TestFullDuplexLoopBack struct {
}

func (id Dot5TestFullDuplexLoopBack) String() string {
	return "TOKENRING-MIB:dot5TestFullDuplexLoopBack"
}

type Dot5ChipSetIBM16 struct {
}

func (id Dot5ChipSetIBM16) String() string {
	return "TOKENRING-MIB:dot5ChipSetIBM16"
}

type Dot5ChipSetTItms380 struct {
}

func (id Dot5ChipSetTItms380) String() string {
	return "TOKENRING-MIB:dot5ChipSetTItms380"
}

type Dot5ChipSetTItms380c16 struct {
}

func (id Dot5ChipSetTItms380c16) String() string {
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
    Dot5StatsTable TOKENRINGMIB_Dot5StatsTable

    // This table contains Token Ring interface timer values, one entry per 802.5
    // interface.
    Dot5TimerTable TOKENRINGMIB_Dot5TimerTable
}

func (tOKENRINGMIB *TOKENRINGMIB) GetEntityData() *types.CommonEntityData {
    tOKENRINGMIB.EntityData.YFilter = tOKENRINGMIB.YFilter
    tOKENRINGMIB.EntityData.YangName = "TOKENRING-MIB"
    tOKENRINGMIB.EntityData.BundleName = "cisco_ios_xe"
    tOKENRINGMIB.EntityData.ParentYangName = "TOKENRING-MIB"
    tOKENRINGMIB.EntityData.SegmentPath = "TOKENRING-MIB:TOKENRING-MIB"
    tOKENRINGMIB.EntityData.AbsolutePath = tOKENRINGMIB.EntityData.SegmentPath
    tOKENRINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tOKENRINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tOKENRINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tOKENRINGMIB.EntityData.Children = types.NewOrderedMap()
    tOKENRINGMIB.EntityData.Children.Append("dot5Table", types.YChild{"Dot5Table", &tOKENRINGMIB.Dot5Table})
    tOKENRINGMIB.EntityData.Children.Append("dot5StatsTable", types.YChild{"Dot5StatsTable", &tOKENRINGMIB.Dot5StatsTable})
    tOKENRINGMIB.EntityData.Children.Append("dot5TimerTable", types.YChild{"Dot5TimerTable", &tOKENRINGMIB.Dot5TimerTable})
    tOKENRINGMIB.EntityData.Leafs = types.NewOrderedMap()

    tOKENRINGMIB.EntityData.YListKeys = []string {}

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
    Dot5Entry []*TOKENRINGMIB_Dot5Table_Dot5Entry
}

func (dot5Table *TOKENRINGMIB_Dot5Table) GetEntityData() *types.CommonEntityData {
    dot5Table.EntityData.YFilter = dot5Table.YFilter
    dot5Table.EntityData.YangName = "dot5Table"
    dot5Table.EntityData.BundleName = "cisco_ios_xe"
    dot5Table.EntityData.ParentYangName = "TOKENRING-MIB"
    dot5Table.EntityData.SegmentPath = "dot5Table"
    dot5Table.EntityData.AbsolutePath = "TOKENRING-MIB:TOKENRING-MIB/" + dot5Table.EntityData.SegmentPath
    dot5Table.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Table.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Table.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Table.EntityData.Children = types.NewOrderedMap()
    dot5Table.EntityData.Children.Append("dot5Entry", types.YChild{"Dot5Entry", nil})
    for i := range dot5Table.Dot5Entry {
        dot5Table.EntityData.Children.Append(types.GetSegmentPath(dot5Table.Dot5Entry[i]), types.YChild{"Dot5Entry", dot5Table.Dot5Entry[i]})
    }
    dot5Table.EntityData.Leafs = types.NewOrderedMap()

    dot5Table.EntityData.YListKeys = []string {}

    return &(dot5Table.EntityData)
}

// TOKENRINGMIB_Dot5Table_Dot5Entry
// A list of Token Ring status and parameter
// values for an 802.5 interface.
type TOKENRINGMIB_Dot5Table_Dot5Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object identifies the 802.5
    // interface for which this entry contains management information.  The value
    // of this object for a particular interface has the same value as the ifIndex
    // object, defined in MIB-II for the same interface. The type is interface{}
    // with range: -2147483648..2147483647.
    Dot5IfIndex interface{}

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
    Dot5RingStatus interface{}

    // The current interface state with respect to entering or leaving the ring.
    // The type is Dot5RingState.
    Dot5RingState interface{}

    // This object indicates the success, or the reason for failure, of the
    // station's most recent attempt to enter the ring. The type is
    // Dot5RingOpenStatus.
    Dot5RingOpenStatus interface{}

    // The ring-speed at the next insertion into the ring.  Note that this may or
    // may not be different to the current ring-speed which is given by MIB-II's
    // ifSpeed.  For interfaces which do not support changing ring-speed,
    // dot5RingSpeed can only be set to its current value.  When dot5RingSpeed has
    // the value unknown(1), the ring's actual ring-speed is to be used. The type
    // is Dot5RingSpeed.
    Dot5RingSpeed interface{}

    // The MAC-address of the up stream neighbor station in the ring. The type is
    // string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot5UpStream interface{}

    // If this object has a value of true(1) then this interface will participate
    // in the active monitor selection process.  If the value is false(2) then it
    // will not. Setting this object does not take effect until the next Active
    // Monitor election, and might not take effect until the next time the
    // interface is opened. The type is Dot5ActMonParticipate.
    Dot5ActMonParticipate interface{}

    // The bit mask of all Token Ring functional addresses for which this
    // interface will accept frames. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot5Functional interface{}

    // The value of MIB-II's sysUpTime object at which the local system last
    // transmitted a Beacon frame on this interface. The type is interface{} with
    // range: 0..4294967295.
    Dot5LastBeaconSent interface{}
}

func (dot5Entry *TOKENRINGMIB_Dot5Table_Dot5Entry) GetEntityData() *types.CommonEntityData {
    dot5Entry.EntityData.YFilter = dot5Entry.YFilter
    dot5Entry.EntityData.YangName = "dot5Entry"
    dot5Entry.EntityData.BundleName = "cisco_ios_xe"
    dot5Entry.EntityData.ParentYangName = "dot5Table"
    dot5Entry.EntityData.SegmentPath = "dot5Entry" + types.AddKeyToken(dot5Entry.Dot5IfIndex, "dot5IfIndex")
    dot5Entry.EntityData.AbsolutePath = "TOKENRING-MIB:TOKENRING-MIB/dot5Table/" + dot5Entry.EntityData.SegmentPath
    dot5Entry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5Entry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5Entry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5Entry.EntityData.Children = types.NewOrderedMap()
    dot5Entry.EntityData.Leafs = types.NewOrderedMap()
    dot5Entry.EntityData.Leafs.Append("dot5IfIndex", types.YLeaf{"Dot5IfIndex", dot5Entry.Dot5IfIndex})
    dot5Entry.EntityData.Leafs.Append("dot5Commands", types.YLeaf{"Dot5Commands", dot5Entry.Dot5Commands})
    dot5Entry.EntityData.Leafs.Append("dot5RingStatus", types.YLeaf{"Dot5RingStatus", dot5Entry.Dot5RingStatus})
    dot5Entry.EntityData.Leafs.Append("dot5RingState", types.YLeaf{"Dot5RingState", dot5Entry.Dot5RingState})
    dot5Entry.EntityData.Leafs.Append("dot5RingOpenStatus", types.YLeaf{"Dot5RingOpenStatus", dot5Entry.Dot5RingOpenStatus})
    dot5Entry.EntityData.Leafs.Append("dot5RingSpeed", types.YLeaf{"Dot5RingSpeed", dot5Entry.Dot5RingSpeed})
    dot5Entry.EntityData.Leafs.Append("dot5UpStream", types.YLeaf{"Dot5UpStream", dot5Entry.Dot5UpStream})
    dot5Entry.EntityData.Leafs.Append("dot5ActMonParticipate", types.YLeaf{"Dot5ActMonParticipate", dot5Entry.Dot5ActMonParticipate})
    dot5Entry.EntityData.Leafs.Append("dot5Functional", types.YLeaf{"Dot5Functional", dot5Entry.Dot5Functional})
    dot5Entry.EntityData.Leafs.Append("dot5LastBeaconSent", types.YLeaf{"Dot5LastBeaconSent", dot5Entry.Dot5LastBeaconSent})

    dot5Entry.EntityData.YListKeys = []string {"Dot5IfIndex"}

    return &(dot5Entry.EntityData)
}

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5ActMonParticipate represents the interface is opened.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5ActMonParticipate string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5ActMonParticipate_true_ TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5ActMonParticipate = "true"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5ActMonParticipate_false_ TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5ActMonParticipate = "false"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands represents dot5Commands and ifOperStatus.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_noop TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "noop"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_open TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "open"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_reset TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "reset"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands_close_ TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5Commands = "close"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus represents recent attempt to enter the ring.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_noOpen TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "noOpen"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_badParam TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "badParam"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_lobeFailed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "lobeFailed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_signalLoss TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "signalLoss"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_insertionTimeout TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "insertionTimeout"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_ringFailed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "ringFailed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_beaconing TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "beaconing"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_duplicateMAC TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "duplicateMAC"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_requestFailed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "requestFailed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_removeReceived TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "removeReceived"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus_open TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingOpenStatus = "open"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed represents to be used.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed_unknown TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed = "unknown"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed_oneMegabit TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed = "oneMegabit"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed_fourMegabit TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed = "fourMegabit"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed_sixteenMegabit TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingSpeed = "sixteenMegabit"
)

// TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState represents to entering or leaving the ring.
type TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState string

const (
    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState_opened TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState = "opened"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState_closed TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState = "closed"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState_opening TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState = "opening"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState_closing TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState = "closing"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState_openFailure TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState = "openFailure"

    TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState_ringFailure TOKENRINGMIB_Dot5Table_Dot5Entry_Dot5RingState = "ringFailure"
)

// TOKENRINGMIB_Dot5StatsTable
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
type TOKENRINGMIB_Dot5StatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry contains the 802.5 statistics for a particular interface. The type
    // is slice of TOKENRINGMIB_Dot5StatsTable_Dot5StatsEntry.
    Dot5StatsEntry []*TOKENRINGMIB_Dot5StatsTable_Dot5StatsEntry
}

func (dot5StatsTable *TOKENRINGMIB_Dot5StatsTable) GetEntityData() *types.CommonEntityData {
    dot5StatsTable.EntityData.YFilter = dot5StatsTable.YFilter
    dot5StatsTable.EntityData.YangName = "dot5StatsTable"
    dot5StatsTable.EntityData.BundleName = "cisco_ios_xe"
    dot5StatsTable.EntityData.ParentYangName = "TOKENRING-MIB"
    dot5StatsTable.EntityData.SegmentPath = "dot5StatsTable"
    dot5StatsTable.EntityData.AbsolutePath = "TOKENRING-MIB:TOKENRING-MIB/" + dot5StatsTable.EntityData.SegmentPath
    dot5StatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5StatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5StatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5StatsTable.EntityData.Children = types.NewOrderedMap()
    dot5StatsTable.EntityData.Children.Append("dot5StatsEntry", types.YChild{"Dot5StatsEntry", nil})
    for i := range dot5StatsTable.Dot5StatsEntry {
        dot5StatsTable.EntityData.Children.Append(types.GetSegmentPath(dot5StatsTable.Dot5StatsEntry[i]), types.YChild{"Dot5StatsEntry", dot5StatsTable.Dot5StatsEntry[i]})
    }
    dot5StatsTable.EntityData.Leafs = types.NewOrderedMap()

    dot5StatsTable.EntityData.YListKeys = []string {}

    return &(dot5StatsTable.EntityData)
}

// TOKENRINGMIB_Dot5StatsTable_Dot5StatsEntry
// An entry contains the 802.5 statistics
// for a particular interface.
type TOKENRINGMIB_Dot5StatsTable_Dot5StatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object identifies the 802.5
    // interface for which this entry contains management information.  The value
    // of this object for a particular interface has the same value as MIB-II's
    // ifIndex object for the same interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5StatsIfIndex interface{}

    // This counter is incremented when a frame or token is copied or repeated by
    // a station, the E bit is zero in the frame or token and one of the following
    // conditions exists: 1) there is a non-data bit (J or K bit) between the SD
    // and the ED of the frame or token, or 2) there is an FCS error in the frame.
    // The type is interface{} with range: 0..4294967295.
    Dot5StatsLineErrors interface{}

    // This counter is incremented when a station detects the absence of
    // transitions for five half-bit timers (burst-five error). The type is
    // interface{} with range: 0..4294967295.
    Dot5StatsBurstErrors interface{}

    // This counter is incremented when a station receives an AMP or SMP frame in
    // which A is equal to C is equal to 0, and then receives another SMP frame
    // with A is equal to C is equal to 0 without first receiving an AMP frame. It
    // denotes a station that cannot set the AC bits properly. The type is
    // interface{} with range: 0..4294967295.
    Dot5StatsACErrors interface{}

    // This counter is incremented when a station transmits an abort delimiter
    // while transmitting. The type is interface{} with range: 0..4294967295.
    Dot5StatsAbortTransErrors interface{}

    // This counter is incremented when a station recognizes an internal error.
    // The type is interface{} with range: 0..4294967295.
    Dot5StatsInternalErrors interface{}

    // This counter is incremented when a station is transmitting and its TRR
    // timer expires. This condition denotes a condition where a transmitting
    // station in strip mode does not receive the trailer of the frame before the
    // TRR timer goes off. The type is interface{} with range: 0..4294967295.
    Dot5StatsLostFrameErrors interface{}

    // This counter is incremented when a station recognizes a frame addressed to
    // its specific address, but has no available buffer space indicating that the
    // station is congested. The type is interface{} with range: 0..4294967295.
    Dot5StatsReceiveCongestions interface{}

    // This counter is incremented when a station recognizes a frame addressed to
    // its specific address and detects that the FS field A bits are set to 1
    // indicating a possible line hit or duplicate address. The type is
    // interface{} with range: 0..4294967295.
    Dot5StatsFrameCopiedErrors interface{}

    // This counter is incremented when a station acting as the active monitor
    // recognizes an error condition that needs a token transmitted. The type is
    // interface{} with range: 0..4294967295.
    Dot5StatsTokenErrors interface{}

    // The number of Soft Errors the interface has detected. It directly
    // corresponds to the number of Report Error MAC frames that this interface
    // has transmitted. Soft Errors are those which are recoverable by the MAC
    // layer protocols. The type is interface{} with range: 0..4294967295.
    Dot5StatsSoftErrors interface{}

    // The number of times this interface has detected an immediately recoverable
    // fatal error.  It denotes the number of times this interface is either
    // transmitting or receiving beacon MAC frames. The type is interface{} with
    // range: 0..4294967295.
    Dot5StatsHardErrors interface{}

    // The number of times this interface has detected the loss of signal
    // condition from the ring. The type is interface{} with range: 0..4294967295.
    Dot5StatsSignalLoss interface{}

    // The number of times this interface has transmitted a beacon frame. The type
    // is interface{} with range: 0..4294967295.
    Dot5StatsTransmitBeacons interface{}

    // The number of Claim Token MAC frames received or transmitted after the
    // interface has received a Ring Purge MAC frame.  This counter signifies the
    // number of times the ring has been purged and is being recovered back into a
    // normal operating state. The type is interface{} with range: 0..4294967295.
    Dot5StatsRecoverys interface{}

    // The number of times the interface has detected an open or short circuit in
    // the lobe data path.  The adapter will be closed and dot5RingState will
    // signify this condition. The type is interface{} with range: 0..4294967295.
    Dot5StatsLobeWires interface{}

    // The number of times the interface has received a Remove Ring Station MAC
    // frame request.  When this frame is received the interface will enter the
    // close state and dot5RingState will signify this condition. The type is
    // interface{} with range: 0..4294967295.
    Dot5StatsRemoves interface{}

    // The number of times the interface has sensed that it is the only station on
    // the ring.  This will happen if the interface is the first one up on a ring,
    // or if there is a hardware problem. The type is interface{} with range:
    // 0..4294967295.
    Dot5StatsSingles interface{}

    // The number of times the interface has detected that the frequency of the
    // incoming signal differs from the expected frequency by more than that
    // specified by the IEEE 802.5 standard. The type is interface{} with range:
    // 0..4294967295.
    Dot5StatsFreqErrors interface{}
}

func (dot5StatsEntry *TOKENRINGMIB_Dot5StatsTable_Dot5StatsEntry) GetEntityData() *types.CommonEntityData {
    dot5StatsEntry.EntityData.YFilter = dot5StatsEntry.YFilter
    dot5StatsEntry.EntityData.YangName = "dot5StatsEntry"
    dot5StatsEntry.EntityData.BundleName = "cisco_ios_xe"
    dot5StatsEntry.EntityData.ParentYangName = "dot5StatsTable"
    dot5StatsEntry.EntityData.SegmentPath = "dot5StatsEntry" + types.AddKeyToken(dot5StatsEntry.Dot5StatsIfIndex, "dot5StatsIfIndex")
    dot5StatsEntry.EntityData.AbsolutePath = "TOKENRING-MIB:TOKENRING-MIB/dot5StatsTable/" + dot5StatsEntry.EntityData.SegmentPath
    dot5StatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5StatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5StatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5StatsEntry.EntityData.Children = types.NewOrderedMap()
    dot5StatsEntry.EntityData.Leafs = types.NewOrderedMap()
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsIfIndex", types.YLeaf{"Dot5StatsIfIndex", dot5StatsEntry.Dot5StatsIfIndex})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsLineErrors", types.YLeaf{"Dot5StatsLineErrors", dot5StatsEntry.Dot5StatsLineErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsBurstErrors", types.YLeaf{"Dot5StatsBurstErrors", dot5StatsEntry.Dot5StatsBurstErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsACErrors", types.YLeaf{"Dot5StatsACErrors", dot5StatsEntry.Dot5StatsACErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsAbortTransErrors", types.YLeaf{"Dot5StatsAbortTransErrors", dot5StatsEntry.Dot5StatsAbortTransErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsInternalErrors", types.YLeaf{"Dot5StatsInternalErrors", dot5StatsEntry.Dot5StatsInternalErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsLostFrameErrors", types.YLeaf{"Dot5StatsLostFrameErrors", dot5StatsEntry.Dot5StatsLostFrameErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsReceiveCongestions", types.YLeaf{"Dot5StatsReceiveCongestions", dot5StatsEntry.Dot5StatsReceiveCongestions})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsFrameCopiedErrors", types.YLeaf{"Dot5StatsFrameCopiedErrors", dot5StatsEntry.Dot5StatsFrameCopiedErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsTokenErrors", types.YLeaf{"Dot5StatsTokenErrors", dot5StatsEntry.Dot5StatsTokenErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsSoftErrors", types.YLeaf{"Dot5StatsSoftErrors", dot5StatsEntry.Dot5StatsSoftErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsHardErrors", types.YLeaf{"Dot5StatsHardErrors", dot5StatsEntry.Dot5StatsHardErrors})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsSignalLoss", types.YLeaf{"Dot5StatsSignalLoss", dot5StatsEntry.Dot5StatsSignalLoss})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsTransmitBeacons", types.YLeaf{"Dot5StatsTransmitBeacons", dot5StatsEntry.Dot5StatsTransmitBeacons})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsRecoverys", types.YLeaf{"Dot5StatsRecoverys", dot5StatsEntry.Dot5StatsRecoverys})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsLobeWires", types.YLeaf{"Dot5StatsLobeWires", dot5StatsEntry.Dot5StatsLobeWires})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsRemoves", types.YLeaf{"Dot5StatsRemoves", dot5StatsEntry.Dot5StatsRemoves})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsSingles", types.YLeaf{"Dot5StatsSingles", dot5StatsEntry.Dot5StatsSingles})
    dot5StatsEntry.EntityData.Leafs.Append("dot5StatsFreqErrors", types.YLeaf{"Dot5StatsFreqErrors", dot5StatsEntry.Dot5StatsFreqErrors})

    dot5StatsEntry.EntityData.YListKeys = []string {"Dot5StatsIfIndex"}

    return &(dot5StatsEntry.EntityData)
}

// TOKENRINGMIB_Dot5TimerTable
// This table contains Token Ring interface
// timer values, one entry per 802.5
// interface.
type TOKENRINGMIB_Dot5TimerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of Token Ring timer values for an 802.5 interface. The type is slice
    // of TOKENRINGMIB_Dot5TimerTable_Dot5TimerEntry.
    Dot5TimerEntry []*TOKENRINGMIB_Dot5TimerTable_Dot5TimerEntry
}

func (dot5TimerTable *TOKENRINGMIB_Dot5TimerTable) GetEntityData() *types.CommonEntityData {
    dot5TimerTable.EntityData.YFilter = dot5TimerTable.YFilter
    dot5TimerTable.EntityData.YangName = "dot5TimerTable"
    dot5TimerTable.EntityData.BundleName = "cisco_ios_xe"
    dot5TimerTable.EntityData.ParentYangName = "TOKENRING-MIB"
    dot5TimerTable.EntityData.SegmentPath = "dot5TimerTable"
    dot5TimerTable.EntityData.AbsolutePath = "TOKENRING-MIB:TOKENRING-MIB/" + dot5TimerTable.EntityData.SegmentPath
    dot5TimerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5TimerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5TimerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5TimerTable.EntityData.Children = types.NewOrderedMap()
    dot5TimerTable.EntityData.Children.Append("dot5TimerEntry", types.YChild{"Dot5TimerEntry", nil})
    for i := range dot5TimerTable.Dot5TimerEntry {
        dot5TimerTable.EntityData.Children.Append(types.GetSegmentPath(dot5TimerTable.Dot5TimerEntry[i]), types.YChild{"Dot5TimerEntry", dot5TimerTable.Dot5TimerEntry[i]})
    }
    dot5TimerTable.EntityData.Leafs = types.NewOrderedMap()

    dot5TimerTable.EntityData.YListKeys = []string {}

    return &(dot5TimerTable.EntityData)
}

// TOKENRINGMIB_Dot5TimerTable_Dot5TimerEntry
// A list of Token Ring timer values for an
// 802.5 interface.
type TOKENRINGMIB_Dot5TimerTable_Dot5TimerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of this object identifies the 802.5
    // interface for which this entry contains timer values.  The value of   this
    // object for a particular interface has the same value as MIB-II's ifIndex
    // object for the same interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5TimerIfIndex interface{}

    // The time-out value used to ensure the interface will return to Repeat
    // State, in units of 100 micro-seconds.  The value should be greater than the
    // maximum ring latency. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5TimerReturnRepeat interface{}

    // Maximum period of time a station is permitted to transmit frames after
    // capturing a token, in units of 100 micro-seconds. The type is interface{}
    // with range: -2147483648..2147483647.
    Dot5TimerHolding interface{}

    // The time-out value for enqueuing of an SMP PDU after reception of an AMP or
    // SMP frame in which the A and C bits were equal to 0, in units of 100
    // micro-seconds. The type is interface{} with range: -2147483648..2147483647.
    Dot5TimerQueuePDU interface{}

    // The time-out value used by the active monitor to detect the absence of
    // valid transmissions, in units of 100 micro-seconds. The type is interface{}
    // with range: -2147483648..2147483647.
    Dot5TimerValidTransmit interface{}

    // The time-out value used to recover from various-related error situations.
    // If N is the maximum number of stations on the ring, the value of this timer
    // is normally: dot5TimerReturnRepeat + N*dot5TimerHolding. The type is
    // interface{} with range: -2147483648..2147483647.
    Dot5TimerNoToken interface{}

    // The time-out value used by the active monitor to stimulate the enqueuing of
    // an AMP PDU for transmission, in units of 100 micro-seconds. The type is
    // interface{} with range: -2147483648..2147483647.
    Dot5TimerActiveMon interface{}

    // The time-out value used by the stand-by monitors to ensure that there is an
    // active monitor on the ring and to detect a continuous stream of tokens, in
    // units of 100 micro-seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5TimerStandbyMon interface{}

    // The time-out value which determines how often a station shall send a Report
    // Error MAC frame to report its error counters, in units of 100
    // micro-seconds. The type is interface{} with range: -2147483648..2147483647.
    Dot5TimerErrorReport interface{}

    // The time-out value which determines how long a station shall remain in the
    // state of transmitting Beacon frames before entering the Bypass state, in
    // units of 100 micro-seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5TimerBeaconTransmit interface{}

    // The time-out value which determines how long a station shall receive Beacon
    // frames from its downstream neighbor before entering the Bypass state, in
    // units of 100 micro-seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot5TimerBeaconReceive interface{}
}

func (dot5TimerEntry *TOKENRINGMIB_Dot5TimerTable_Dot5TimerEntry) GetEntityData() *types.CommonEntityData {
    dot5TimerEntry.EntityData.YFilter = dot5TimerEntry.YFilter
    dot5TimerEntry.EntityData.YangName = "dot5TimerEntry"
    dot5TimerEntry.EntityData.BundleName = "cisco_ios_xe"
    dot5TimerEntry.EntityData.ParentYangName = "dot5TimerTable"
    dot5TimerEntry.EntityData.SegmentPath = "dot5TimerEntry" + types.AddKeyToken(dot5TimerEntry.Dot5TimerIfIndex, "dot5TimerIfIndex")
    dot5TimerEntry.EntityData.AbsolutePath = "TOKENRING-MIB:TOKENRING-MIB/dot5TimerTable/" + dot5TimerEntry.EntityData.SegmentPath
    dot5TimerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot5TimerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot5TimerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot5TimerEntry.EntityData.Children = types.NewOrderedMap()
    dot5TimerEntry.EntityData.Leafs = types.NewOrderedMap()
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerIfIndex", types.YLeaf{"Dot5TimerIfIndex", dot5TimerEntry.Dot5TimerIfIndex})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerReturnRepeat", types.YLeaf{"Dot5TimerReturnRepeat", dot5TimerEntry.Dot5TimerReturnRepeat})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerHolding", types.YLeaf{"Dot5TimerHolding", dot5TimerEntry.Dot5TimerHolding})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerQueuePDU", types.YLeaf{"Dot5TimerQueuePDU", dot5TimerEntry.Dot5TimerQueuePDU})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerValidTransmit", types.YLeaf{"Dot5TimerValidTransmit", dot5TimerEntry.Dot5TimerValidTransmit})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerNoToken", types.YLeaf{"Dot5TimerNoToken", dot5TimerEntry.Dot5TimerNoToken})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerActiveMon", types.YLeaf{"Dot5TimerActiveMon", dot5TimerEntry.Dot5TimerActiveMon})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerStandbyMon", types.YLeaf{"Dot5TimerStandbyMon", dot5TimerEntry.Dot5TimerStandbyMon})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerErrorReport", types.YLeaf{"Dot5TimerErrorReport", dot5TimerEntry.Dot5TimerErrorReport})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerBeaconTransmit", types.YLeaf{"Dot5TimerBeaconTransmit", dot5TimerEntry.Dot5TimerBeaconTransmit})
    dot5TimerEntry.EntityData.Leafs.Append("dot5TimerBeaconReceive", types.YLeaf{"Dot5TimerBeaconReceive", dot5TimerEntry.Dot5TimerBeaconReceive})

    dot5TimerEntry.EntityData.YListKeys = []string {"Dot5TimerIfIndex"}

    return &(dot5TimerEntry.EntityData)
}

