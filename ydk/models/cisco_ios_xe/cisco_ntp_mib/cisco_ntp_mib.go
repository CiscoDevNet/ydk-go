// This MIB module defines a MIB which provides
// mechanisms to monitor an NTP server.
// 
// The MIB is derived from the Technical Report
// #Management of the NTP with SNMP# TR No. 98-09
// authored by A.S. Sethi and Dave Mills in the
// University of Delaware.
// 
// Below is a brief overview of NTP system architecture
// and implementation model. This will help understand
// the objects defined below and their relationships.
// 
// NTP Intro:
// The Network Time Protocol (NTP) Version 3, is used to
// synchronize timekeeping among a set of distributed
// time servers and clients.  The service model is based
// on a returnable-time design which depends only on
// measured clock offsets, but does not require reliable
// message delivery.  The synchronization subnet uses a
// self-organizing, hierarchical master-slave
// configuration, with synchronization paths determined
// by a minimum-weight spanning tree.  While multiple
// masters (primary servers) may exist, there is no
// requirement for an election protocol.
// 
// System Archiecture:
// In the NTP model a number of primary reference
// sources, synchronized by wire or radio to national
// standards, are connected to widely accessible
// resources, such as backbone gateways, and operated as
// primary time servers.  The purpose of NTP is to convey
// timekeeping information from these servers to other
// time servers via the Internet and also to cross-check
// clocks and mitigate errors due to equipment or
// propagation failures.  Some number of local-net hosts
// or gateways, acting as secondary time servers, run NTP
// with one or more of the primary servers.  In order to
// reduce the protocol overhead, the secondary servers
// distribute time via NTP to the remaining local-net
// hosts.  In the interest of reliability, selected hosts
// can be equipped with less accurate but less expensive
// radio clocks and used for backup in case of failure of
// the primary and/or secondary servers or communication
// paths between them.
// 
// NTP is designed to produce three products: clock
// offset, round-trip delay and dispersion, all of which
// are relative to a selected reference clock.  Clock
// offset represents the amount to adjust the local clock
// to bring it into correspondence with the reference
// clock.  Roundtrip delay provides the capability to
// launch a message to arrive at the reference clock at a
// specified time.  Dispersion represents the maximum
// error of the local clock relative to the reference
// clock.  Since most host time servers will synchronize
// via another peer time server, there are two components
// in each of these three products, those determined by
// the peer relative to the primary reference source of
// standard time and those measured by the host relative
// to the peer.  Each of these components are maintained
// separately in the protocol in order to facilitate
// error control and management of the subnet itself.  
// They provide not only precision measurements of offset
// and delay, but also definitive maximum error bounds,
// so that the user interface can determine not only the
// time, but the quality of the time as well.
// 
// Implementation Model:
// In what may be the most common client/server model a
// client sends an NTP message to one or more servers and
// processes the replies as received.  The server
// interchanges addresses and ports, overwrites certain
// fields in the message, recalculates the checksum and
// returns the message immediately.  Information included
// in the NTP message allows the client to determine the
// server time with respect to local time and adjust the
// local clock accordingly.  In addition, the message
// includes information to calculate the expected
// timekeeping accuracy and reliability, as well as
// select the best from possibly several servers.
// 
// While the client/server model may suffice for use on
// local nets involving a public server and perhaps many
// workstation clients, the full generality of NTP
// requires distributed participation of a number of
// client/servers or peers arranged in a dynamically
// reconfigurable, hierarchically distributed
// configuration.  It also requires sophisticated
// algorithms for association management, data
// manipulation and local-clock control.
// 
// Glossary:
// 1. Host: Refers to an instantiation of the NTP
//         protocol on a local processor.
// 2. Peer: Refers to an instantiation of the NTP
//         protocol on a remote processor connected by
//         a network path from the local host.
package cisco_ntp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ntp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-NTP-MIB CISCO-NTP-MIB}", reflect.TypeOf(CISCONTPMIB{}))
    ydk.RegisterEntity("CISCO-NTP-MIB:CISCO-NTP-MIB", reflect.TypeOf(CISCONTPMIB{}))
}

// NTPLeapIndicator represents 11, alarm condition (clock not synchronized)
type NTPLeapIndicator string

const (
    NTPLeapIndicator_noWarning NTPLeapIndicator = "noWarning"

    NTPLeapIndicator_addSecond NTPLeapIndicator = "addSecond"

    NTPLeapIndicator_subtractSecond NTPLeapIndicator = "subtractSecond"

    NTPLeapIndicator_alarm NTPLeapIndicator = "alarm"
)

// CISCONTPMIB
type CISCONTPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cntpsystem CISCONTPMIB_Cntpsystem

    // This table provides information on the peers with which the local NTP
    // server has associations.  The peers are also NTP servers but running on
    // different hosts.
    Cntppeersvartable CISCONTPMIB_Cntppeersvartable

    // The following table contains NTP state variables used by the NTP clock
    // filter and selection algorithms. This table depicts a shift register.  Each
    // stage in the shift register is a 3-tuple consisting of the measured clock
    // offset, measured clock delay and measured clock dispersion associated with
    // a single observation.  An important factor affecting the accuracy and
    // reliability of time distribution is the complex of algorithms used to
    // reduce the effect of statistical errors and falsetickers due to failure of
    // various subnet components, reference sources or propagation media.  The NTP
    // clock-filter and selection algorithms are designed to do exactly this.  The
    // objects in the filter register table below are used by these algorthims to
    // minimize the error in the calculated time.
    Cntpfilterregistertable CISCONTPMIB_Cntpfilterregistertable
}

func (cISCONTPMIB *CISCONTPMIB) GetEntityData() *types.CommonEntityData {
    cISCONTPMIB.EntityData.YFilter = cISCONTPMIB.YFilter
    cISCONTPMIB.EntityData.YangName = "CISCO-NTP-MIB"
    cISCONTPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCONTPMIB.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cISCONTPMIB.EntityData.SegmentPath = "CISCO-NTP-MIB:CISCO-NTP-MIB"
    cISCONTPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCONTPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCONTPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCONTPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCONTPMIB.EntityData.Children["cntpSystem"] = types.YChild{"Cntpsystem", &cISCONTPMIB.Cntpsystem}
    cISCONTPMIB.EntityData.Children["cntpPeersVarTable"] = types.YChild{"Cntppeersvartable", &cISCONTPMIB.Cntppeersvartable}
    cISCONTPMIB.EntityData.Children["cntpFilterRegisterTable"] = types.YChild{"Cntpfilterregistertable", &cISCONTPMIB.Cntpfilterregistertable}
    cISCONTPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCONTPMIB.EntityData)
}

// CISCONTPMIB_Cntpsystem
type CISCONTPMIB_Cntpsystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Two-bit code warning of an impending leap second to be inserted in the NTP
    // timescale. This object can be set only when the cntpSysStratum has a value
    // of 1. The type is NTPLeapIndicator.
    Cntpsysleap interface{}

    // The stratum of the local clock. If the value is set to 1, i.e., this is a
    // primary reference, then the Primary-Clock procedure described in Section
    // 3.4.6, in RFC-1305 is invoked. The type is interface{} with range: 0..255.
    Cntpsysstratum interface{}

    // Signed integer indicating the precision of the system clock, in seconds to
    // the nearest power of two.  The value must be rounded to the next larger
    // power of two; for instance, a 50-Hz (20 ms) or 60-Hz (16.67 ms)
    // power-frequency clock would be assigned the value -5 (31.25 ms), while a
    // 1000-Hz (1 ms) crystal-controlled clock would be assigned the value -9
    // (1.95 ms). The type is interface{} with range: -20..20.
    Cntpsysprecision interface{}

    // A signed fixed-point number indicating the total round-trip delay in
    // seconds, to the primary reference source at the root of the synchronization
    // subnet. The type is string with length: 4. Units are seconds.
    Cntpsysrootdelay interface{}

    // The maximum error in seconds, relative to the primary reference source at
    // the root of the synchronization subnet.  Only positive values greater than
    // zero are possible. The type is string with length: 4. Units are seconds.
    Cntpsysrootdispersion interface{}

    // The reference identifier of the local clock. The type is string with
    // length: 4.
    Cntpsysrefid interface{}

    // The local time when the local clock was last updated.  If the local clock
    // has never been synchronized, the value is zero. The type is string with
    // length: 8.
    Cntpsysreftime interface{}

    // The interval at which the NTP server polls other NTP servers to synchronize
    // its clock. The type is interface{} with range: -20..20.
    Cntpsyspoll interface{}

    // The current synchronization source.  This will contain the unique
    // association identifier cntpPeersAssocId of the corresponding peer entry in
    // the cntpPeersVarTable of the peer acting as the synchronization source.  If
    // there is no peer, the value will be 0. The type is interface{} with range:
    // 0..2147483647.
    Cntpsyspeer interface{}

    // The current local time.  Local time is derived from the hardware clock of
    // the particular machine and increments at intervals depending on the design
    // used. The type is string with length: 8.
    Cntpsysclock interface{}

    // Current state of the NTP server with values coded as follows: 1: server
    // status is unknown 2: server is not running 3: server is not synchronized to
    // any time source 4: server is synchronized to its own local clock 5: server
    // is synchronized to a local hardware refclock (e.g. GPS) 6: server is
    // synchronized to a remote NTP server. The type is Cntpsyssrvstatus.
    Cntpsyssrvstatus interface{}
}

func (cntpsystem *CISCONTPMIB_Cntpsystem) GetEntityData() *types.CommonEntityData {
    cntpsystem.EntityData.YFilter = cntpsystem.YFilter
    cntpsystem.EntityData.YangName = "cntpSystem"
    cntpsystem.EntityData.BundleName = "cisco_ios_xe"
    cntpsystem.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cntpsystem.EntityData.SegmentPath = "cntpSystem"
    cntpsystem.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpsystem.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpsystem.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpsystem.EntityData.Children = make(map[string]types.YChild)
    cntpsystem.EntityData.Leafs = make(map[string]types.YLeaf)
    cntpsystem.EntityData.Leafs["cntpSysLeap"] = types.YLeaf{"Cntpsysleap", cntpsystem.Cntpsysleap}
    cntpsystem.EntityData.Leafs["cntpSysStratum"] = types.YLeaf{"Cntpsysstratum", cntpsystem.Cntpsysstratum}
    cntpsystem.EntityData.Leafs["cntpSysPrecision"] = types.YLeaf{"Cntpsysprecision", cntpsystem.Cntpsysprecision}
    cntpsystem.EntityData.Leafs["cntpSysRootDelay"] = types.YLeaf{"Cntpsysrootdelay", cntpsystem.Cntpsysrootdelay}
    cntpsystem.EntityData.Leafs["cntpSysRootDispersion"] = types.YLeaf{"Cntpsysrootdispersion", cntpsystem.Cntpsysrootdispersion}
    cntpsystem.EntityData.Leafs["cntpSysRefId"] = types.YLeaf{"Cntpsysrefid", cntpsystem.Cntpsysrefid}
    cntpsystem.EntityData.Leafs["cntpSysRefTime"] = types.YLeaf{"Cntpsysreftime", cntpsystem.Cntpsysreftime}
    cntpsystem.EntityData.Leafs["cntpSysPoll"] = types.YLeaf{"Cntpsyspoll", cntpsystem.Cntpsyspoll}
    cntpsystem.EntityData.Leafs["cntpSysPeer"] = types.YLeaf{"Cntpsyspeer", cntpsystem.Cntpsyspeer}
    cntpsystem.EntityData.Leafs["cntpSysClock"] = types.YLeaf{"Cntpsysclock", cntpsystem.Cntpsysclock}
    cntpsystem.EntityData.Leafs["cntpSysSrvStatus"] = types.YLeaf{"Cntpsyssrvstatus", cntpsystem.Cntpsyssrvstatus}
    return &(cntpsystem.EntityData)
}

// CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus represents 6: server is synchronized to a remote NTP server
type CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus string

const (
    CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus_unknown CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus = "unknown"

    CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus_notRunning CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus = "notRunning"

    CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus_notSynchronized CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus = "notSynchronized"

    CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus_syncToLocal CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus = "syncToLocal"

    CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus_syncToRefclock CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus = "syncToRefclock"

    CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus_syncToRemoteServer CISCONTPMIB_Cntpsystem_Cntpsyssrvstatus = "syncToRemoteServer"
)

// CISCONTPMIB_Cntppeersvartable
// This table provides information on the peers with
// which the local NTP server has associations.  The
// peers are also NTP servers but running on different
// hosts.
type CISCONTPMIB_Cntppeersvartable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each peers' entry provides NTP information retrieved from a particular peer
    // NTP server.  Each peer is identified by a unique association identifier. 
    // Entries are automatically created when the user configures the NTP server
    // to be associated with remote peers.  Similarly entries are deleted when the
    // user removes the peer association from the NTP server.  Entries can also be
    // created by the management station by setting values for the following
    // objects: cntpPeersPeerAddress or cntpPeersPeerName,  cntpPeersHostAddress
    // and cntpPeersMode and making the cntpPeersEntryStatus as active(1).  At the
    // least, the management station has to set a value for cntpPeersPeerAddress
    // or cntpPeersPeerName to make the row active. The type is slice of
    // CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry.
    Cntppeersvarentry []CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry
}

func (cntppeersvartable *CISCONTPMIB_Cntppeersvartable) GetEntityData() *types.CommonEntityData {
    cntppeersvartable.EntityData.YFilter = cntppeersvartable.YFilter
    cntppeersvartable.EntityData.YangName = "cntpPeersVarTable"
    cntppeersvartable.EntityData.BundleName = "cisco_ios_xe"
    cntppeersvartable.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cntppeersvartable.EntityData.SegmentPath = "cntpPeersVarTable"
    cntppeersvartable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntppeersvartable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntppeersvartable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntppeersvartable.EntityData.Children = make(map[string]types.YChild)
    cntppeersvartable.EntityData.Children["cntpPeersVarEntry"] = types.YChild{"Cntppeersvarentry", nil}
    for i := range cntppeersvartable.Cntppeersvarentry {
        cntppeersvartable.EntityData.Children[types.GetSegmentPath(&cntppeersvartable.Cntppeersvarentry[i])] = types.YChild{"Cntppeersvarentry", &cntppeersvartable.Cntppeersvarentry[i]}
    }
    cntppeersvartable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cntppeersvartable.EntityData)
}

// CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry
// Each peers' entry provides NTP information retrieved
// from a particular peer NTP server.  Each peer is
// identified by a unique association identifier.
// 
// Entries are automatically created when the user
// configures the NTP server to be associated with remote
// peers.  Similarly entries are deleted when the user
// removes the peer association from the NTP server.
// 
// Entries can also be created by the management station
// by setting values for the following objects:
// cntpPeersPeerAddress or cntpPeersPeerName, 
// cntpPeersHostAddress and
// cntpPeersMode and making the cntpPeersEntryStatus as
// active(1).  At the least, the management station has
// to set a value for cntpPeersPeerAddress or
// cntpPeersPeerName to make the row active.
type CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer value greater than 0 that uniquely
    // identifies a peer with which the local NTP server is associated. The type
    // is interface{} with range: 0..2147483647.
    Cntppeersassocid interface{}

    // This is a bit indicating that the association was created from
    // configuration information and should not be de-associated even if the peer
    // becomes unreachable. The type is bool.
    Cntppeersconfigured interface{}

    // The IP address of the peer.  When creating a new association, a value
    // should be set either for this object or the corresponding instance of 
    // cntpPeersPeerName, before the row is made active. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cntppeerspeeraddress interface{}

    // The UDP port number on which the peer receives NTP messages. The type is
    // interface{} with range: 1..65535.
    Cntppeerspeerport interface{}

    // The IP address of the local host.  Multi-homing can be supported using this
    // object. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cntppeershostaddress interface{}

    // The UDP port number on which the local host receives NTP messages. The type
    // is interface{} with range: 1..65535.
    Cntppeershostport interface{}

    // Two-bit code warning of an impending leap second to be inserted in the NTP
    // timescale of the peer. The type is NTPLeapIndicator.
    Cntppeersleap interface{}

    // The association mode of the NTP server, with values coded as follows, 0,
    // unspecified 1, symmetric active - A host operating in this mode        
    // sends periodic messages regardless of the         reachability state or
    // stratum of its peer.  By         operating in this mode the host announces
    // its         willingness to synchronize and be synchronized         by the
    // peer 2, symmetric passive - This type of association is         ordinarily
    // created upon arrival of a message         from a peer operating in the
    // symmetric active         mode and persists only as long as the peer is     
    // reachable and operating at a stratum level         less than or equal to
    // the host; otherwise, the         association is dissolved.  However, the   
    // association will always persist until at least         one message has been
    // sent in reply.  By         operating in this mode the host announces its   
    // willingness to synchronize and be synchronized         by the peer 3,
    // client -  A host operating in this mode sends         periodic messages
    // regardless of the         reachability state or stratum of its peer.  By   
    // operating in this mode the host, usually a LAN         workstation,
    // announces its willingness to be         synchronized by, but not to
    // synchronize the peer 4, server - This type of association is ordinarily    
    // created upon arrival of a client request message         and exists only in
    // order to reply to that         request, after which the association is     
    // dissolved.  By operating in this mode the host,         usually a LAN time
    // server, announces its         willingness to synchronize, but not to be    
    // synchronized by the peer 5, broadcast - A host operating in this mode sends
    // periodic messages regardless of the         reachability state or stratum
    // of the peers.         By operating in this mode the host, usually a        
    // LAN time server operating on a high-speed         broadcast medium,
    // announces its willingness to         synchronize all of the peers, but not
    // to be         synchronized by any of them 6, reserved for NTP control
    // messages 7, reserved for private use.  When creating a new peer
    // association, if no value is specified for this object, it defaults to
    // symmetricActive(1). The type is Cntppeersmode.
    Cntppeersmode interface{}

    // The stratum of the peer clock. The type is interface{} with range: 0..255.
    Cntppeersstratum interface{}

    // The interval at which the peer polls the local host. The type is
    // interface{} with range: -20..20.
    Cntppeerspeerpoll interface{}

    // The interval at which the local host polls the peer. The type is
    // interface{} with range: -20..20.
    Cntppeershostpoll interface{}

    // Signed integer indicating the precision of the peer clock, in seconds to
    // the nearest power of two.  The value must be rounded to the next larger
    // power of two; for instance, a 50-Hz (20 ms) or 60-Hz (16.67 ms)
    // power-frequency clock would be assigned the value -5 (31.25 ms), while a
    // 1000-Hz (1 ms) crystal-controlled clock would be assigned the value -9
    // (1.95 ms). The type is interface{} with range: -20..20.
    Cntppeersprecision interface{}

    // A signed fixed-point number indicating the total round-trip delay in
    // seconds, from the peer to the primary reference source at the root of the
    // synchronization subnet. The type is string with length: 4. Units are
    // seconds.
    Cntppeersrootdelay interface{}

    // The maximum error in seconds, of the peer clock relative to the primary
    // reference source at the root of the synchronization subnet.  Only positive
    // values greater than zero are possible. The type is string with length: 4.
    // Units are seconds.
    Cntppeersrootdispersion interface{}

    // The reference identifier of the peer. The type is string with length: 4.
    Cntppeersrefid interface{}

    // The local time at the peer when its clock was last updated.  If the peer
    // clock has never been synchronized, the value is zero. The type is string
    // with length: 8.
    Cntppeersreftime interface{}

    // The local time at the peer, when its latest NTP message was sent.  If the
    // peer becomes unreachable the value is set to zero. The type is string with
    // length: 8.
    Cntppeersorgtime interface{}

    // The local time, when the latest NTP message from the peer arrived.  If the
    // peer becomes unreachable the value is set to zero. The type is string with
    // length: 8.
    Cntppeersreceivetime interface{}

    // The local time at which the NTP message departed the sender. The type is
    // string with length: 8.
    Cntppeerstransmittime interface{}

    // The local time, when the most recent NTP message was received from the peer
    // that was used to calculate the skew dispersion.  This represents only the
    // 32-bit integer part of the NTPTimestamp. The type is interface{} with
    // range: 0..2147483647.
    Cntppeersupdatetime interface{}

    // A shift register of used to determine the reachability status of the peer,
    // with bits entering from the least significant (rightmost) end.  A peer is
    // considered reachable if at least one bit in this register is set to one
    // i.e, if the value of this object is non-zero. The data in the shift
    // register would be populated by the NTP protocol procedures. The type is
    // interface{} with range: 0..255.
    Cntppeersreach interface{}

    // The interval in seconds, between transmitted NTP messages from the local
    // host to the peer. The type is interface{} with range: 0..2147483647. Units
    // are seconds.
    Cntppeerstimer interface{}

    // The estimated offset of the peer clock relative to the local clock, in
    // seconds.  The host determines the value of this object using the NTP
    // clock-filter algorithm. The type is string with length: 4. Units are
    // seconds.
    Cntppeersoffset interface{}

    // The estimated round-trip delay of the peer clock relative to the local
    // clock over the network path between them, in seconds.  The host determines
    // the value of this object using the NTP clock-filter algorithm. The type is
    // string with length: 4. Units are seconds.
    Cntppeersdelay interface{}

    // The estimated maximum error of the peer clock relative to the local clock
    // over the network path between them, in seconds.  The host determines the
    // value of this object using the NTP clock-filter algorithm. The type is
    // string with length: 4. Units are seconds.
    Cntppeersdispersion interface{}

    // The number of valid entries for a peer in the Filter Register Table. Since,
    // the Filter Register Table is optional, this object will have a value 0 if
    // the Filter Register Table is not implemented. The type is interface{} with
    // range: 0..4294967295.
    Cntppeersfiltervalidentries interface{}

    // The status object for this row. When a management station is creating a new
    // row, it should set the value for cntpPeersPeerAddress at least, before the
    // row can be made active(1). The type is RowStatus.
    Cntppeersentrystatus interface{}

    // The local time, when the most recent NTP message was received from the peer
    // that was used to calculate the skew dispersion.  This represents only the
    // 32-bit integer part of the NTPTimestamp. The type is interface{} with
    // range: 0..4294967295.
    Cntppeersupdatetimerev1 interface{}

    // This object specifies whether this peer is the preferred one over the
    // others. By default, when the value of this object is 'false', NTP chooses 
    // the peer with which to synchronize the time on  the local system. If this
    // object is set to 'true', NTP will choose the corresponding peer to
    // synchronize the time with. If multiple entries have this object set to
    // 'true', NTP will choose the first one to be set. This object is a means to
    // override the selection of the peer by NTP. The type is bool.
    Cntppeersprefpeer interface{}

    // Represents the type of the corresponding instance of cntpPeersPeerName
    // object. The type is InetAddressType.
    Cntppeerspeertype interface{}

    // The address of the peer. When creating a new association, a value must be
    // set for either this object or the corresponding instance of
    // cntpPeersPeerAddress object, before the row is made active. The type is
    // string with length: 0..255.
    Cntppeerspeername interface{}
}

func (cntppeersvarentry *CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry) GetEntityData() *types.CommonEntityData {
    cntppeersvarentry.EntityData.YFilter = cntppeersvarentry.YFilter
    cntppeersvarentry.EntityData.YangName = "cntpPeersVarEntry"
    cntppeersvarentry.EntityData.BundleName = "cisco_ios_xe"
    cntppeersvarentry.EntityData.ParentYangName = "cntpPeersVarTable"
    cntppeersvarentry.EntityData.SegmentPath = "cntpPeersVarEntry" + "[cntpPeersAssocId='" + fmt.Sprintf("%v", cntppeersvarentry.Cntppeersassocid) + "']"
    cntppeersvarentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntppeersvarentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntppeersvarentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntppeersvarentry.EntityData.Children = make(map[string]types.YChild)
    cntppeersvarentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cntppeersvarentry.EntityData.Leafs["cntpPeersAssocId"] = types.YLeaf{"Cntppeersassocid", cntppeersvarentry.Cntppeersassocid}
    cntppeersvarentry.EntityData.Leafs["cntpPeersConfigured"] = types.YLeaf{"Cntppeersconfigured", cntppeersvarentry.Cntppeersconfigured}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPeerAddress"] = types.YLeaf{"Cntppeerspeeraddress", cntppeersvarentry.Cntppeerspeeraddress}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPeerPort"] = types.YLeaf{"Cntppeerspeerport", cntppeersvarentry.Cntppeerspeerport}
    cntppeersvarentry.EntityData.Leafs["cntpPeersHostAddress"] = types.YLeaf{"Cntppeershostaddress", cntppeersvarentry.Cntppeershostaddress}
    cntppeersvarentry.EntityData.Leafs["cntpPeersHostPort"] = types.YLeaf{"Cntppeershostport", cntppeersvarentry.Cntppeershostport}
    cntppeersvarentry.EntityData.Leafs["cntpPeersLeap"] = types.YLeaf{"Cntppeersleap", cntppeersvarentry.Cntppeersleap}
    cntppeersvarentry.EntityData.Leafs["cntpPeersMode"] = types.YLeaf{"Cntppeersmode", cntppeersvarentry.Cntppeersmode}
    cntppeersvarentry.EntityData.Leafs["cntpPeersStratum"] = types.YLeaf{"Cntppeersstratum", cntppeersvarentry.Cntppeersstratum}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPeerPoll"] = types.YLeaf{"Cntppeerspeerpoll", cntppeersvarentry.Cntppeerspeerpoll}
    cntppeersvarentry.EntityData.Leafs["cntpPeersHostPoll"] = types.YLeaf{"Cntppeershostpoll", cntppeersvarentry.Cntppeershostpoll}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPrecision"] = types.YLeaf{"Cntppeersprecision", cntppeersvarentry.Cntppeersprecision}
    cntppeersvarentry.EntityData.Leafs["cntpPeersRootDelay"] = types.YLeaf{"Cntppeersrootdelay", cntppeersvarentry.Cntppeersrootdelay}
    cntppeersvarentry.EntityData.Leafs["cntpPeersRootDispersion"] = types.YLeaf{"Cntppeersrootdispersion", cntppeersvarentry.Cntppeersrootdispersion}
    cntppeersvarentry.EntityData.Leafs["cntpPeersRefId"] = types.YLeaf{"Cntppeersrefid", cntppeersvarentry.Cntppeersrefid}
    cntppeersvarentry.EntityData.Leafs["cntpPeersRefTime"] = types.YLeaf{"Cntppeersreftime", cntppeersvarentry.Cntppeersreftime}
    cntppeersvarentry.EntityData.Leafs["cntpPeersOrgTime"] = types.YLeaf{"Cntppeersorgtime", cntppeersvarentry.Cntppeersorgtime}
    cntppeersvarentry.EntityData.Leafs["cntpPeersReceiveTime"] = types.YLeaf{"Cntppeersreceivetime", cntppeersvarentry.Cntppeersreceivetime}
    cntppeersvarentry.EntityData.Leafs["cntpPeersTransmitTime"] = types.YLeaf{"Cntppeerstransmittime", cntppeersvarentry.Cntppeerstransmittime}
    cntppeersvarentry.EntityData.Leafs["cntpPeersUpdateTime"] = types.YLeaf{"Cntppeersupdatetime", cntppeersvarentry.Cntppeersupdatetime}
    cntppeersvarentry.EntityData.Leafs["cntpPeersReach"] = types.YLeaf{"Cntppeersreach", cntppeersvarentry.Cntppeersreach}
    cntppeersvarentry.EntityData.Leafs["cntpPeersTimer"] = types.YLeaf{"Cntppeerstimer", cntppeersvarentry.Cntppeerstimer}
    cntppeersvarentry.EntityData.Leafs["cntpPeersOffset"] = types.YLeaf{"Cntppeersoffset", cntppeersvarentry.Cntppeersoffset}
    cntppeersvarentry.EntityData.Leafs["cntpPeersDelay"] = types.YLeaf{"Cntppeersdelay", cntppeersvarentry.Cntppeersdelay}
    cntppeersvarentry.EntityData.Leafs["cntpPeersDispersion"] = types.YLeaf{"Cntppeersdispersion", cntppeersvarentry.Cntppeersdispersion}
    cntppeersvarentry.EntityData.Leafs["cntpPeersFilterValidEntries"] = types.YLeaf{"Cntppeersfiltervalidentries", cntppeersvarentry.Cntppeersfiltervalidentries}
    cntppeersvarentry.EntityData.Leafs["cntpPeersEntryStatus"] = types.YLeaf{"Cntppeersentrystatus", cntppeersvarentry.Cntppeersentrystatus}
    cntppeersvarentry.EntityData.Leafs["cntpPeersUpdateTimeRev1"] = types.YLeaf{"Cntppeersupdatetimerev1", cntppeersvarentry.Cntppeersupdatetimerev1}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPrefPeer"] = types.YLeaf{"Cntppeersprefpeer", cntppeersvarentry.Cntppeersprefpeer}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPeerType"] = types.YLeaf{"Cntppeerspeertype", cntppeersvarentry.Cntppeerspeertype}
    cntppeersvarentry.EntityData.Leafs["cntpPeersPeerName"] = types.YLeaf{"Cntppeerspeername", cntppeersvarentry.Cntppeerspeername}
    return &(cntppeersvarentry.EntityData)
}

// CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode represents symmetricActive(1).
type CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode string

const (
    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_unspecified CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "unspecified"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_symmetricActive CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "symmetricActive"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_symmetricPassive CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "symmetricPassive"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_client CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "client"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_server CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "server"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_broadcast CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "broadcast"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_reservedControl CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "reservedControl"

    CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode_reservedPrivate CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersmode = "reservedPrivate"
)

// CISCONTPMIB_Cntpfilterregistertable
// The following table contains NTP state variables
// used by the NTP clock filter and selection algorithms.
// This table depicts a shift register.  Each stage in
// the shift register is a 3-tuple consisting of the
// measured clock offset, measured clock delay and
// measured clock dispersion associated with a single
// observation.
// 
// An important factor affecting the accuracy and
// reliability of time distribution is the complex of
// algorithms used to reduce the effect of statistical
// errors and falsetickers due to failure of various
// subnet components, reference sources or propagation
// media.  The NTP clock-filter and selection algorithms
// are designed to do exactly this.  The objects in the
// filter register table below are used by these
// algorthims to minimize the error in the calculated
// time.
type CISCONTPMIB_Cntpfilterregistertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry corresponds to one stage of the shift register, i.e., one
    // reading of the variables clock delay, clock offset and clock dispersion. 
    // Entries are automatically created whenever a peer is configured and deleted
    // when the peer is removed. The type is slice of
    // CISCONTPMIB_Cntpfilterregistertable_Cntpfilterregisterentry.
    Cntpfilterregisterentry []CISCONTPMIB_Cntpfilterregistertable_Cntpfilterregisterentry
}

func (cntpfilterregistertable *CISCONTPMIB_Cntpfilterregistertable) GetEntityData() *types.CommonEntityData {
    cntpfilterregistertable.EntityData.YFilter = cntpfilterregistertable.YFilter
    cntpfilterregistertable.EntityData.YangName = "cntpFilterRegisterTable"
    cntpfilterregistertable.EntityData.BundleName = "cisco_ios_xe"
    cntpfilterregistertable.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cntpfilterregistertable.EntityData.SegmentPath = "cntpFilterRegisterTable"
    cntpfilterregistertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpfilterregistertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpfilterregistertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpfilterregistertable.EntityData.Children = make(map[string]types.YChild)
    cntpfilterregistertable.EntityData.Children["cntpFilterRegisterEntry"] = types.YChild{"Cntpfilterregisterentry", nil}
    for i := range cntpfilterregistertable.Cntpfilterregisterentry {
        cntpfilterregistertable.EntityData.Children[types.GetSegmentPath(&cntpfilterregistertable.Cntpfilterregisterentry[i])] = types.YChild{"Cntpfilterregisterentry", &cntpfilterregistertable.Cntpfilterregisterentry[i]}
    }
    cntpfilterregistertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cntpfilterregistertable.EntityData)
}

// CISCONTPMIB_Cntpfilterregistertable_Cntpfilterregisterentry
// Each entry corresponds to one stage of the shift
// register, i.e., one reading of the variables clock
// delay, clock offset and clock dispersion.
// 
// Entries are automatically created whenever a peer is
// configured and deleted when the peer is removed.
type CISCONTPMIB_Cntpfilterregistertable_Cntpfilterregisterentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // cisco_ntp_mib.CISCONTPMIB_Cntppeersvartable_Cntppeersvarentry_Cntppeersassocid
    Cntppeersassocid interface{}

    // This attribute is a key. An integer value in the specified range that is
    // used to index into the table.  The size of the table is fixed at 8.  Each
    // entry identifies a particular reading of the clock filter variables in the
    // shift register.  Entries are added starting at index 1.  The index wraps
    // back to 1 when it reaches 8.  When the index wraps back, the new entries
    // will overwrite the old entries effectively deleting the old entry. The type
    // is interface{} with range: 1..8.
    Cntpfilterindex interface{}

    // The offset of the peer clock relative to the local clock in seconds. The
    // type is string with length: 4. Units are seconds.
    Cntpfilterpeersoffset interface{}

    // Round-trip delay of the peer clock relative to the local clock over the
    // network path between them, in seconds.  This variable can take on both
    // positive and negative values, depending on clock precision and skew-error
    // accumulation. The type is string with length: 4. Units are seconds.
    Cntpfilterpeersdelay interface{}

    // The maximum error of the peer clock relative to the local clock over the
    // network path between them, in seconds.  Only positive values greater than
    // zero are possible. The type is string with length: 4. Units are seconds.
    Cntpfilterpeersdispersion interface{}
}

func (cntpfilterregisterentry *CISCONTPMIB_Cntpfilterregistertable_Cntpfilterregisterentry) GetEntityData() *types.CommonEntityData {
    cntpfilterregisterentry.EntityData.YFilter = cntpfilterregisterentry.YFilter
    cntpfilterregisterentry.EntityData.YangName = "cntpFilterRegisterEntry"
    cntpfilterregisterentry.EntityData.BundleName = "cisco_ios_xe"
    cntpfilterregisterentry.EntityData.ParentYangName = "cntpFilterRegisterTable"
    cntpfilterregisterentry.EntityData.SegmentPath = "cntpFilterRegisterEntry" + "[cntpPeersAssocId='" + fmt.Sprintf("%v", cntpfilterregisterentry.Cntppeersassocid) + "']" + "[cntpFilterIndex='" + fmt.Sprintf("%v", cntpfilterregisterentry.Cntpfilterindex) + "']"
    cntpfilterregisterentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpfilterregisterentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpfilterregisterentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpfilterregisterentry.EntityData.Children = make(map[string]types.YChild)
    cntpfilterregisterentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cntpfilterregisterentry.EntityData.Leafs["cntpPeersAssocId"] = types.YLeaf{"Cntppeersassocid", cntpfilterregisterentry.Cntppeersassocid}
    cntpfilterregisterentry.EntityData.Leafs["cntpFilterIndex"] = types.YLeaf{"Cntpfilterindex", cntpfilterregisterentry.Cntpfilterindex}
    cntpfilterregisterentry.EntityData.Leafs["cntpFilterPeersOffset"] = types.YLeaf{"Cntpfilterpeersoffset", cntpfilterregisterentry.Cntpfilterpeersoffset}
    cntpfilterregisterentry.EntityData.Leafs["cntpFilterPeersDelay"] = types.YLeaf{"Cntpfilterpeersdelay", cntpfilterregisterentry.Cntpfilterpeersdelay}
    cntpfilterregisterentry.EntityData.Leafs["cntpFilterPeersDispersion"] = types.YLeaf{"Cntpfilterpeersdispersion", cntpfilterregisterentry.Cntpfilterpeersdispersion}
    return &(cntpfilterregisterentry.EntityData)
}

