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

    
    CntpSystem CISCONTPMIB_CntpSystem

    // This table provides information on the peers with which the local NTP
    // server has associations.  The peers are also NTP servers but running on
    // different hosts.
    CntpPeersVarTable CISCONTPMIB_CntpPeersVarTable

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
    CntpFilterRegisterTable CISCONTPMIB_CntpFilterRegisterTable
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

    cISCONTPMIB.EntityData.Children = types.NewOrderedMap()
    cISCONTPMIB.EntityData.Children.Append("cntpSystem", types.YChild{"CntpSystem", &cISCONTPMIB.CntpSystem})
    cISCONTPMIB.EntityData.Children.Append("cntpPeersVarTable", types.YChild{"CntpPeersVarTable", &cISCONTPMIB.CntpPeersVarTable})
    cISCONTPMIB.EntityData.Children.Append("cntpFilterRegisterTable", types.YChild{"CntpFilterRegisterTable", &cISCONTPMIB.CntpFilterRegisterTable})
    cISCONTPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCONTPMIB.EntityData.YListKeys = []string {}

    return &(cISCONTPMIB.EntityData)
}

// CISCONTPMIB_CntpSystem
type CISCONTPMIB_CntpSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Two-bit code warning of an impending leap second to be inserted in the NTP
    // timescale. This object can be set only when the cntpSysStratum has a value
    // of 1. The type is NTPLeapIndicator.
    CntpSysLeap interface{}

    // The stratum of the local clock. If the value is set to 1, i.e., this is a
    // primary reference, then the Primary-Clock procedure described in Section
    // 3.4.6, in RFC-1305 is invoked. The type is interface{} with range: 0..255.
    CntpSysStratum interface{}

    // Signed integer indicating the precision of the system clock, in seconds to
    // the nearest power of two.  The value must be rounded to the next larger
    // power of two; for instance, a 50-Hz (20 ms) or 60-Hz (16.67 ms)
    // power-frequency clock would be assigned the value -5 (31.25 ms), while a
    // 1000-Hz (1 ms) crystal-controlled clock would be assigned the value -9
    // (1.95 ms). The type is interface{} with range: -20..20.
    CntpSysPrecision interface{}

    // A signed fixed-point number indicating the total round-trip delay in
    // seconds, to the primary reference source at the root of the synchronization
    // subnet. The type is string with length: 4. Units are seconds.
    CntpSysRootDelay interface{}

    // The maximum error in seconds, relative to the primary reference source at
    // the root of the synchronization subnet.  Only positive values greater than
    // zero are possible. The type is string with length: 4. Units are seconds.
    CntpSysRootDispersion interface{}

    // The reference identifier of the local clock. The type is string with
    // length: 4.
    CntpSysRefId interface{}

    // The local time when the local clock was last updated.  If the local clock
    // has never been synchronized, the value is zero. The type is string with
    // length: 8.
    CntpSysRefTime interface{}

    // The interval at which the NTP server polls other NTP servers to synchronize
    // its clock. The type is interface{} with range: -20..20.
    CntpSysPoll interface{}

    // The current synchronization source.  This will contain the unique
    // association identifier cntpPeersAssocId of the corresponding peer entry in
    // the cntpPeersVarTable of the peer acting as the synchronization source.  If
    // there is no peer, the value will be 0. The type is interface{} with range:
    // 0..2147483647.
    CntpSysPeer interface{}

    // The current local time.  Local time is derived from the hardware clock of
    // the particular machine and increments at intervals depending on the design
    // used. The type is string with length: 8.
    CntpSysClock interface{}

    // Current state of the NTP server with values coded as follows: 1: server
    // status is unknown 2: server is not running 3: server is not synchronized to
    // any time source 4: server is synchronized to its own local clock 5: server
    // is synchronized to a local hardware refclock (e.g. GPS) 6: server is
    // synchronized to a remote NTP server. The type is CntpSysSrvStatus.
    CntpSysSrvStatus interface{}
}

func (cntpSystem *CISCONTPMIB_CntpSystem) GetEntityData() *types.CommonEntityData {
    cntpSystem.EntityData.YFilter = cntpSystem.YFilter
    cntpSystem.EntityData.YangName = "cntpSystem"
    cntpSystem.EntityData.BundleName = "cisco_ios_xe"
    cntpSystem.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cntpSystem.EntityData.SegmentPath = "cntpSystem"
    cntpSystem.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpSystem.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpSystem.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpSystem.EntityData.Children = types.NewOrderedMap()
    cntpSystem.EntityData.Leafs = types.NewOrderedMap()
    cntpSystem.EntityData.Leafs.Append("cntpSysLeap", types.YLeaf{"CntpSysLeap", cntpSystem.CntpSysLeap})
    cntpSystem.EntityData.Leafs.Append("cntpSysStratum", types.YLeaf{"CntpSysStratum", cntpSystem.CntpSysStratum})
    cntpSystem.EntityData.Leafs.Append("cntpSysPrecision", types.YLeaf{"CntpSysPrecision", cntpSystem.CntpSysPrecision})
    cntpSystem.EntityData.Leafs.Append("cntpSysRootDelay", types.YLeaf{"CntpSysRootDelay", cntpSystem.CntpSysRootDelay})
    cntpSystem.EntityData.Leafs.Append("cntpSysRootDispersion", types.YLeaf{"CntpSysRootDispersion", cntpSystem.CntpSysRootDispersion})
    cntpSystem.EntityData.Leafs.Append("cntpSysRefId", types.YLeaf{"CntpSysRefId", cntpSystem.CntpSysRefId})
    cntpSystem.EntityData.Leafs.Append("cntpSysRefTime", types.YLeaf{"CntpSysRefTime", cntpSystem.CntpSysRefTime})
    cntpSystem.EntityData.Leafs.Append("cntpSysPoll", types.YLeaf{"CntpSysPoll", cntpSystem.CntpSysPoll})
    cntpSystem.EntityData.Leafs.Append("cntpSysPeer", types.YLeaf{"CntpSysPeer", cntpSystem.CntpSysPeer})
    cntpSystem.EntityData.Leafs.Append("cntpSysClock", types.YLeaf{"CntpSysClock", cntpSystem.CntpSysClock})
    cntpSystem.EntityData.Leafs.Append("cntpSysSrvStatus", types.YLeaf{"CntpSysSrvStatus", cntpSystem.CntpSysSrvStatus})

    cntpSystem.EntityData.YListKeys = []string {}

    return &(cntpSystem.EntityData)
}

// CISCONTPMIB_CntpSystem_CntpSysSrvStatus represents 6: server is synchronized to a remote NTP server
type CISCONTPMIB_CntpSystem_CntpSysSrvStatus string

const (
    CISCONTPMIB_CntpSystem_CntpSysSrvStatus_unknown CISCONTPMIB_CntpSystem_CntpSysSrvStatus = "unknown"

    CISCONTPMIB_CntpSystem_CntpSysSrvStatus_notRunning CISCONTPMIB_CntpSystem_CntpSysSrvStatus = "notRunning"

    CISCONTPMIB_CntpSystem_CntpSysSrvStatus_notSynchronized CISCONTPMIB_CntpSystem_CntpSysSrvStatus = "notSynchronized"

    CISCONTPMIB_CntpSystem_CntpSysSrvStatus_syncToLocal CISCONTPMIB_CntpSystem_CntpSysSrvStatus = "syncToLocal"

    CISCONTPMIB_CntpSystem_CntpSysSrvStatus_syncToRefclock CISCONTPMIB_CntpSystem_CntpSysSrvStatus = "syncToRefclock"

    CISCONTPMIB_CntpSystem_CntpSysSrvStatus_syncToRemoteServer CISCONTPMIB_CntpSystem_CntpSysSrvStatus = "syncToRemoteServer"
)

// CISCONTPMIB_CntpPeersVarTable
// This table provides information on the peers with
// which the local NTP server has associations.  The
// peers are also NTP servers but running on different
// hosts.
type CISCONTPMIB_CntpPeersVarTable struct {
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
    // CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry.
    CntpPeersVarEntry []*CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry
}

func (cntpPeersVarTable *CISCONTPMIB_CntpPeersVarTable) GetEntityData() *types.CommonEntityData {
    cntpPeersVarTable.EntityData.YFilter = cntpPeersVarTable.YFilter
    cntpPeersVarTable.EntityData.YangName = "cntpPeersVarTable"
    cntpPeersVarTable.EntityData.BundleName = "cisco_ios_xe"
    cntpPeersVarTable.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cntpPeersVarTable.EntityData.SegmentPath = "cntpPeersVarTable"
    cntpPeersVarTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpPeersVarTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpPeersVarTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpPeersVarTable.EntityData.Children = types.NewOrderedMap()
    cntpPeersVarTable.EntityData.Children.Append("cntpPeersVarEntry", types.YChild{"CntpPeersVarEntry", nil})
    for i := range cntpPeersVarTable.CntpPeersVarEntry {
        cntpPeersVarTable.EntityData.Children.Append(types.GetSegmentPath(cntpPeersVarTable.CntpPeersVarEntry[i]), types.YChild{"CntpPeersVarEntry", cntpPeersVarTable.CntpPeersVarEntry[i]})
    }
    cntpPeersVarTable.EntityData.Leafs = types.NewOrderedMap()

    cntpPeersVarTable.EntityData.YListKeys = []string {}

    return &(cntpPeersVarTable.EntityData)
}

// CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry
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
type CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer value greater than 0 that uniquely
    // identifies a peer with which the local NTP server is associated. The type
    // is interface{} with range: 0..2147483647.
    CntpPeersAssocId interface{}

    // This is a bit indicating that the association was created from
    // configuration information and should not be de-associated even if the peer
    // becomes unreachable. The type is bool.
    CntpPeersConfigured interface{}

    // The IP address of the peer.  When creating a new association, a value
    // should be set either for this object or the corresponding instance of 
    // cntpPeersPeerName, before the row is made active. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CntpPeersPeerAddress interface{}

    // The UDP port number on which the peer receives NTP messages. The type is
    // interface{} with range: 1..65535.
    CntpPeersPeerPort interface{}

    // The IP address of the local host.  Multi-homing can be supported using this
    // object. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CntpPeersHostAddress interface{}

    // The UDP port number on which the local host receives NTP messages. The type
    // is interface{} with range: 1..65535.
    CntpPeersHostPort interface{}

    // Two-bit code warning of an impending leap second to be inserted in the NTP
    // timescale of the peer. The type is NTPLeapIndicator.
    CntpPeersLeap interface{}

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
    // symmetricActive(1). The type is CntpPeersMode.
    CntpPeersMode interface{}

    // The stratum of the peer clock. The type is interface{} with range: 0..255.
    CntpPeersStratum interface{}

    // The interval at which the peer polls the local host. The type is
    // interface{} with range: -20..20.
    CntpPeersPeerPoll interface{}

    // The interval at which the local host polls the peer. The type is
    // interface{} with range: -20..20.
    CntpPeersHostPoll interface{}

    // Signed integer indicating the precision of the peer clock, in seconds to
    // the nearest power of two.  The value must be rounded to the next larger
    // power of two; for instance, a 50-Hz (20 ms) or 60-Hz (16.67 ms)
    // power-frequency clock would be assigned the value -5 (31.25 ms), while a
    // 1000-Hz (1 ms) crystal-controlled clock would be assigned the value -9
    // (1.95 ms). The type is interface{} with range: -20..20.
    CntpPeersPrecision interface{}

    // A signed fixed-point number indicating the total round-trip delay in
    // seconds, from the peer to the primary reference source at the root of the
    // synchronization subnet. The type is string with length: 4. Units are
    // seconds.
    CntpPeersRootDelay interface{}

    // The maximum error in seconds, of the peer clock relative to the primary
    // reference source at the root of the synchronization subnet.  Only positive
    // values greater than zero are possible. The type is string with length: 4.
    // Units are seconds.
    CntpPeersRootDispersion interface{}

    // The reference identifier of the peer. The type is string with length: 4.
    CntpPeersRefId interface{}

    // The local time at the peer when its clock was last updated.  If the peer
    // clock has never been synchronized, the value is zero. The type is string
    // with length: 8.
    CntpPeersRefTime interface{}

    // The local time at the peer, when its latest NTP message was sent.  If the
    // peer becomes unreachable the value is set to zero. The type is string with
    // length: 8.
    CntpPeersOrgTime interface{}

    // The local time, when the latest NTP message from the peer arrived.  If the
    // peer becomes unreachable the value is set to zero. The type is string with
    // length: 8.
    CntpPeersReceiveTime interface{}

    // The local time at which the NTP message departed the sender. The type is
    // string with length: 8.
    CntpPeersTransmitTime interface{}

    // The local time, when the most recent NTP message was received from the peer
    // that was used to calculate the skew dispersion.  This represents only the
    // 32-bit integer part of the NTPTimestamp. The type is interface{} with
    // range: 0..2147483647.
    CntpPeersUpdateTime interface{}

    // A shift register of used to determine the reachability status of the peer,
    // with bits entering from the least significant (rightmost) end.  A peer is
    // considered reachable if at least one bit in this register is set to one
    // i.e, if the value of this object is non-zero. The data in the shift
    // register would be populated by the NTP protocol procedures. The type is
    // interface{} with range: 0..255.
    CntpPeersReach interface{}

    // The interval in seconds, between transmitted NTP messages from the local
    // host to the peer. The type is interface{} with range: 0..2147483647. Units
    // are seconds.
    CntpPeersTimer interface{}

    // The estimated offset of the peer clock relative to the local clock, in
    // seconds.  The host determines the value of this object using the NTP
    // clock-filter algorithm. The type is string with length: 4. Units are
    // seconds.
    CntpPeersOffset interface{}

    // The estimated round-trip delay of the peer clock relative to the local
    // clock over the network path between them, in seconds.  The host determines
    // the value of this object using the NTP clock-filter algorithm. The type is
    // string with length: 4. Units are seconds.
    CntpPeersDelay interface{}

    // The estimated maximum error of the peer clock relative to the local clock
    // over the network path between them, in seconds.  The host determines the
    // value of this object using the NTP clock-filter algorithm. The type is
    // string with length: 4. Units are seconds.
    CntpPeersDispersion interface{}

    // The number of valid entries for a peer in the Filter Register Table. Since,
    // the Filter Register Table is optional, this object will have a value 0 if
    // the Filter Register Table is not implemented. The type is interface{} with
    // range: 0..4294967295.
    CntpPeersFilterValidEntries interface{}

    // The status object for this row. When a management station is creating a new
    // row, it should set the value for cntpPeersPeerAddress at least, before the
    // row can be made active(1). The type is RowStatus.
    CntpPeersEntryStatus interface{}

    // The local time, when the most recent NTP message was received from the peer
    // that was used to calculate the skew dispersion.  This represents only the
    // 32-bit integer part of the NTPTimestamp. The type is interface{} with
    // range: 0..4294967295.
    CntpPeersUpdateTimeRev1 interface{}

    // This object specifies whether this peer is the preferred one over the
    // others. By default, when the value of this object is 'false', NTP chooses 
    // the peer with which to synchronize the time on  the local system. If this
    // object is set to 'true', NTP will choose the corresponding peer to
    // synchronize the time with. If multiple entries have this object set to
    // 'true', NTP will choose the first one to be set. This object is a means to
    // override the selection of the peer by NTP. The type is bool.
    CntpPeersPrefPeer interface{}

    // Represents the type of the corresponding instance of cntpPeersPeerName
    // object. The type is InetAddressType.
    CntpPeersPeerType interface{}

    // The address of the peer. When creating a new association, a value must be
    // set for either this object or the corresponding instance of
    // cntpPeersPeerAddress object, before the row is made active. The type is
    // string with length: 0..255.
    CntpPeersPeerName interface{}
}

func (cntpPeersVarEntry *CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry) GetEntityData() *types.CommonEntityData {
    cntpPeersVarEntry.EntityData.YFilter = cntpPeersVarEntry.YFilter
    cntpPeersVarEntry.EntityData.YangName = "cntpPeersVarEntry"
    cntpPeersVarEntry.EntityData.BundleName = "cisco_ios_xe"
    cntpPeersVarEntry.EntityData.ParentYangName = "cntpPeersVarTable"
    cntpPeersVarEntry.EntityData.SegmentPath = "cntpPeersVarEntry" + types.AddKeyToken(cntpPeersVarEntry.CntpPeersAssocId, "cntpPeersAssocId")
    cntpPeersVarEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpPeersVarEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpPeersVarEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpPeersVarEntry.EntityData.Children = types.NewOrderedMap()
    cntpPeersVarEntry.EntityData.Leafs = types.NewOrderedMap()
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersAssocId", types.YLeaf{"CntpPeersAssocId", cntpPeersVarEntry.CntpPeersAssocId})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersConfigured", types.YLeaf{"CntpPeersConfigured", cntpPeersVarEntry.CntpPeersConfigured})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPeerAddress", types.YLeaf{"CntpPeersPeerAddress", cntpPeersVarEntry.CntpPeersPeerAddress})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPeerPort", types.YLeaf{"CntpPeersPeerPort", cntpPeersVarEntry.CntpPeersPeerPort})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersHostAddress", types.YLeaf{"CntpPeersHostAddress", cntpPeersVarEntry.CntpPeersHostAddress})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersHostPort", types.YLeaf{"CntpPeersHostPort", cntpPeersVarEntry.CntpPeersHostPort})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersLeap", types.YLeaf{"CntpPeersLeap", cntpPeersVarEntry.CntpPeersLeap})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersMode", types.YLeaf{"CntpPeersMode", cntpPeersVarEntry.CntpPeersMode})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersStratum", types.YLeaf{"CntpPeersStratum", cntpPeersVarEntry.CntpPeersStratum})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPeerPoll", types.YLeaf{"CntpPeersPeerPoll", cntpPeersVarEntry.CntpPeersPeerPoll})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersHostPoll", types.YLeaf{"CntpPeersHostPoll", cntpPeersVarEntry.CntpPeersHostPoll})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPrecision", types.YLeaf{"CntpPeersPrecision", cntpPeersVarEntry.CntpPeersPrecision})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersRootDelay", types.YLeaf{"CntpPeersRootDelay", cntpPeersVarEntry.CntpPeersRootDelay})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersRootDispersion", types.YLeaf{"CntpPeersRootDispersion", cntpPeersVarEntry.CntpPeersRootDispersion})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersRefId", types.YLeaf{"CntpPeersRefId", cntpPeersVarEntry.CntpPeersRefId})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersRefTime", types.YLeaf{"CntpPeersRefTime", cntpPeersVarEntry.CntpPeersRefTime})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersOrgTime", types.YLeaf{"CntpPeersOrgTime", cntpPeersVarEntry.CntpPeersOrgTime})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersReceiveTime", types.YLeaf{"CntpPeersReceiveTime", cntpPeersVarEntry.CntpPeersReceiveTime})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersTransmitTime", types.YLeaf{"CntpPeersTransmitTime", cntpPeersVarEntry.CntpPeersTransmitTime})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersUpdateTime", types.YLeaf{"CntpPeersUpdateTime", cntpPeersVarEntry.CntpPeersUpdateTime})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersReach", types.YLeaf{"CntpPeersReach", cntpPeersVarEntry.CntpPeersReach})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersTimer", types.YLeaf{"CntpPeersTimer", cntpPeersVarEntry.CntpPeersTimer})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersOffset", types.YLeaf{"CntpPeersOffset", cntpPeersVarEntry.CntpPeersOffset})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersDelay", types.YLeaf{"CntpPeersDelay", cntpPeersVarEntry.CntpPeersDelay})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersDispersion", types.YLeaf{"CntpPeersDispersion", cntpPeersVarEntry.CntpPeersDispersion})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersFilterValidEntries", types.YLeaf{"CntpPeersFilterValidEntries", cntpPeersVarEntry.CntpPeersFilterValidEntries})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersEntryStatus", types.YLeaf{"CntpPeersEntryStatus", cntpPeersVarEntry.CntpPeersEntryStatus})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersUpdateTimeRev1", types.YLeaf{"CntpPeersUpdateTimeRev1", cntpPeersVarEntry.CntpPeersUpdateTimeRev1})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPrefPeer", types.YLeaf{"CntpPeersPrefPeer", cntpPeersVarEntry.CntpPeersPrefPeer})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPeerType", types.YLeaf{"CntpPeersPeerType", cntpPeersVarEntry.CntpPeersPeerType})
    cntpPeersVarEntry.EntityData.Leafs.Append("cntpPeersPeerName", types.YLeaf{"CntpPeersPeerName", cntpPeersVarEntry.CntpPeersPeerName})

    cntpPeersVarEntry.EntityData.YListKeys = []string {"CntpPeersAssocId"}

    return &(cntpPeersVarEntry.EntityData)
}

// CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode represents symmetricActive(1).
type CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode string

const (
    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_unspecified CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "unspecified"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_symmetricActive CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "symmetricActive"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_symmetricPassive CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "symmetricPassive"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_client CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "client"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_server CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "server"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_broadcast CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "broadcast"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_reservedControl CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "reservedControl"

    CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode_reservedPrivate CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersMode = "reservedPrivate"
)

// CISCONTPMIB_CntpFilterRegisterTable
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
type CISCONTPMIB_CntpFilterRegisterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry corresponds to one stage of the shift register, i.e., one
    // reading of the variables clock delay, clock offset and clock dispersion. 
    // Entries are automatically created whenever a peer is configured and deleted
    // when the peer is removed. The type is slice of
    // CISCONTPMIB_CntpFilterRegisterTable_CntpFilterRegisterEntry.
    CntpFilterRegisterEntry []*CISCONTPMIB_CntpFilterRegisterTable_CntpFilterRegisterEntry
}

func (cntpFilterRegisterTable *CISCONTPMIB_CntpFilterRegisterTable) GetEntityData() *types.CommonEntityData {
    cntpFilterRegisterTable.EntityData.YFilter = cntpFilterRegisterTable.YFilter
    cntpFilterRegisterTable.EntityData.YangName = "cntpFilterRegisterTable"
    cntpFilterRegisterTable.EntityData.BundleName = "cisco_ios_xe"
    cntpFilterRegisterTable.EntityData.ParentYangName = "CISCO-NTP-MIB"
    cntpFilterRegisterTable.EntityData.SegmentPath = "cntpFilterRegisterTable"
    cntpFilterRegisterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpFilterRegisterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpFilterRegisterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpFilterRegisterTable.EntityData.Children = types.NewOrderedMap()
    cntpFilterRegisterTable.EntityData.Children.Append("cntpFilterRegisterEntry", types.YChild{"CntpFilterRegisterEntry", nil})
    for i := range cntpFilterRegisterTable.CntpFilterRegisterEntry {
        cntpFilterRegisterTable.EntityData.Children.Append(types.GetSegmentPath(cntpFilterRegisterTable.CntpFilterRegisterEntry[i]), types.YChild{"CntpFilterRegisterEntry", cntpFilterRegisterTable.CntpFilterRegisterEntry[i]})
    }
    cntpFilterRegisterTable.EntityData.Leafs = types.NewOrderedMap()

    cntpFilterRegisterTable.EntityData.YListKeys = []string {}

    return &(cntpFilterRegisterTable.EntityData)
}

// CISCONTPMIB_CntpFilterRegisterTable_CntpFilterRegisterEntry
// Each entry corresponds to one stage of the shift
// register, i.e., one reading of the variables clock
// delay, clock offset and clock dispersion.
// 
// Entries are automatically created whenever a peer is
// configured and deleted when the peer is removed.
type CISCONTPMIB_CntpFilterRegisterTable_CntpFilterRegisterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..2147483647.
    // Refers to
    // cisco_ntp_mib.CISCONTPMIB_CntpPeersVarTable_CntpPeersVarEntry_CntpPeersAssocId
    CntpPeersAssocId interface{}

    // This attribute is a key. An integer value in the specified range that is
    // used to index into the table.  The size of the table is fixed at 8.  Each
    // entry identifies a particular reading of the clock filter variables in the
    // shift register.  Entries are added starting at index 1.  The index wraps
    // back to 1 when it reaches 8.  When the index wraps back, the new entries
    // will overwrite the old entries effectively deleting the old entry. The type
    // is interface{} with range: 1..8.
    CntpFilterIndex interface{}

    // The offset of the peer clock relative to the local clock in seconds. The
    // type is string with length: 4. Units are seconds.
    CntpFilterPeersOffset interface{}

    // Round-trip delay of the peer clock relative to the local clock over the
    // network path between them, in seconds.  This variable can take on both
    // positive and negative values, depending on clock precision and skew-error
    // accumulation. The type is string with length: 4. Units are seconds.
    CntpFilterPeersDelay interface{}

    // The maximum error of the peer clock relative to the local clock over the
    // network path between them, in seconds.  Only positive values greater than
    // zero are possible. The type is string with length: 4. Units are seconds.
    CntpFilterPeersDispersion interface{}
}

func (cntpFilterRegisterEntry *CISCONTPMIB_CntpFilterRegisterTable_CntpFilterRegisterEntry) GetEntityData() *types.CommonEntityData {
    cntpFilterRegisterEntry.EntityData.YFilter = cntpFilterRegisterEntry.YFilter
    cntpFilterRegisterEntry.EntityData.YangName = "cntpFilterRegisterEntry"
    cntpFilterRegisterEntry.EntityData.BundleName = "cisco_ios_xe"
    cntpFilterRegisterEntry.EntityData.ParentYangName = "cntpFilterRegisterTable"
    cntpFilterRegisterEntry.EntityData.SegmentPath = "cntpFilterRegisterEntry" + types.AddKeyToken(cntpFilterRegisterEntry.CntpPeersAssocId, "cntpPeersAssocId") + types.AddKeyToken(cntpFilterRegisterEntry.CntpFilterIndex, "cntpFilterIndex")
    cntpFilterRegisterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cntpFilterRegisterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cntpFilterRegisterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cntpFilterRegisterEntry.EntityData.Children = types.NewOrderedMap()
    cntpFilterRegisterEntry.EntityData.Leafs = types.NewOrderedMap()
    cntpFilterRegisterEntry.EntityData.Leafs.Append("cntpPeersAssocId", types.YLeaf{"CntpPeersAssocId", cntpFilterRegisterEntry.CntpPeersAssocId})
    cntpFilterRegisterEntry.EntityData.Leafs.Append("cntpFilterIndex", types.YLeaf{"CntpFilterIndex", cntpFilterRegisterEntry.CntpFilterIndex})
    cntpFilterRegisterEntry.EntityData.Leafs.Append("cntpFilterPeersOffset", types.YLeaf{"CntpFilterPeersOffset", cntpFilterRegisterEntry.CntpFilterPeersOffset})
    cntpFilterRegisterEntry.EntityData.Leafs.Append("cntpFilterPeersDelay", types.YLeaf{"CntpFilterPeersDelay", cntpFilterRegisterEntry.CntpFilterPeersDelay})
    cntpFilterRegisterEntry.EntityData.Leafs.Append("cntpFilterPeersDispersion", types.YLeaf{"CntpFilterPeersDispersion", cntpFilterRegisterEntry.CntpFilterPeersDispersion})

    cntpFilterRegisterEntry.EntityData.YListKeys = []string {"CntpPeersAssocId", "CntpFilterIndex"}

    return &(cntpFilterRegisterEntry.EntityData)
}

