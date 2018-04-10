// The MIB module for managing TCP implementations.
// 
// Copyright (C) The Internet Society (2005). This version
// of this MIB module is a part of RFC 4022; see the RFC
// itself for full legal notices.
package tcp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tcp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:TCP-MIB TCP-MIB}", reflect.TypeOf(TCPMIB{}))
    ydk.RegisterEntity("TCP-MIB:TCP-MIB", reflect.TypeOf(TCPMIB{}))
}

// TCPMIB
type TCPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Tcp TCPMIB_Tcp

    // A table containing information about existing IPv4-specific TCP connections
    // or listeners.  This table has been deprecated in favor of the version
    // neutral tcpConnectionTable.
    Tcpconntable TCPMIB_Tcpconntable

    // A table containing information about existing TCP connections.  Note that
    // unlike earlier TCP MIBs, there is a separate table for connections in the
    // LISTEN state.
    Tcpconnectiontable TCPMIB_Tcpconnectiontable

    // A table containing information about TCP listeners.  A listening
    // application can be represented in three possible ways:  1. An application
    // that is willing to accept both IPv4 and    IPv6 datagrams is represented by
    // a tcpListenerLocalAddressType of unknown (0) and    a
    // tcpListenerLocalAddress of ''h (a zero-length    octet-string).  2. An
    // application that is willing to accept only IPv4 or    IPv6 datagrams is
    // represented by a    tcpListenerLocalAddressType of the appropriate address 
    // type and a tcpListenerLocalAddress of '0.0.0.0' or '::'    respectively. 
    // 3. An application that is listening for data destined    only to a specific
    // IP address, but from any remote    system, is represented by a
    // tcpListenerLocalAddressType    of an appropriate address type, with   
    // tcpListenerLocalAddress as the specific local address.  NOTE: The address
    // type in this table represents the address type used for the communication,
    // irrespective of the higher-layer abstraction.  For example, an application
    // using IPv6 'sockets' to communicate via IPv4 between ::ffff:10.0.0.1 and
    // ::ffff:10.0.0.2 would use InetAddressType ipv4(1)).
    Tcplistenertable TCPMIB_Tcplistenertable
}

func (tCPMIB *TCPMIB) GetEntityData() *types.CommonEntityData {
    tCPMIB.EntityData.YFilter = tCPMIB.YFilter
    tCPMIB.EntityData.YangName = "TCP-MIB"
    tCPMIB.EntityData.BundleName = "cisco_ios_xe"
    tCPMIB.EntityData.ParentYangName = "TCP-MIB"
    tCPMIB.EntityData.SegmentPath = "TCP-MIB:TCP-MIB"
    tCPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tCPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tCPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tCPMIB.EntityData.Children = make(map[string]types.YChild)
    tCPMIB.EntityData.Children["tcp"] = types.YChild{"Tcp", &tCPMIB.Tcp}
    tCPMIB.EntityData.Children["tcpConnTable"] = types.YChild{"Tcpconntable", &tCPMIB.Tcpconntable}
    tCPMIB.EntityData.Children["tcpConnectionTable"] = types.YChild{"Tcpconnectiontable", &tCPMIB.Tcpconnectiontable}
    tCPMIB.EntityData.Children["tcpListenerTable"] = types.YChild{"Tcplistenertable", &tCPMIB.Tcplistenertable}
    tCPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tCPMIB.EntityData)
}

// TCPMIB_Tcp
type TCPMIB_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The algorithm used to determine the timeout value used for retransmitting
    // unacknowledged octets. The type is Tcprtoalgorithm.
    Tcprtoalgorithm interface{}

    // The minimum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds. More refined semantics for objects of
    // this type depend on the algorithm used to determine the retransmission
    // timeout; in particular, the IETF standard algorithm rfc2988(5) provides a
    // minimum value. The type is interface{} with range: 0..2147483647. Units are
    // milliseconds.
    Tcprtomin interface{}

    // The maximum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds. More refined semantics for objects of
    // this type depend on the algorithm used to determine the retransmission
    // timeout; in particular, the IETF standard algorithm rfc2988(5) provides an
    // upper bound (as part of an adaptive backoff algorithm). The type is
    // interface{} with range: 0..2147483647. Units are milliseconds.
    Tcprtomax interface{}

    // The limit on the total number of TCP connections the entity can support. 
    // In entities where the maximum number of connections is dynamic, this object
    // should contain the value -1. The type is interface{} with range:
    // -1..2147483647.
    Tcpmaxconn interface{}

    // The number of times that TCP connections have made a direct transition to
    // the SYN-SENT state from the CLOSED state.  Discontinuities in the value of
    // this counter are indicated via discontinuities in the value of sysUpTime.
    // The type is interface{} with range: 0..4294967295.
    Tcpactiveopens interface{}

    // The number of times TCP connections have made a direct transition to the
    // SYN-RCVD state from the LISTEN state.  Discontinuities in the value of this
    // counter are indicated via discontinuities in the value of sysUpTime. The
    // type is interface{} with range: 0..4294967295.
    Tcppassiveopens interface{}

    // The number of times that TCP connections have made a direct transition to
    // the CLOSED state from either the SYN-SENT state or the SYN-RCVD state, plus
    // the number of times that TCP connections have made a direct transition to
    // the LISTEN state from the SYN-RCVD state.  Discontinuities in the value of
    // this counter are indicated via discontinuities in the value of sysUpTime.
    // The type is interface{} with range: 0..4294967295.
    Tcpattemptfails interface{}

    // The number of times that TCP connections have made a direct transition to
    // the CLOSED state from either the ESTABLISHED state or the CLOSE-WAIT state.
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    Tcpestabresets interface{}

    // The number of TCP connections for which the current state is either
    // ESTABLISHED or CLOSE-WAIT. The type is interface{} with range:
    // 0..4294967295.
    Tcpcurrestab interface{}

    // The total number of segments received, including those received in error. 
    // This count includes segments received on currently established connections.
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    Tcpinsegs interface{}

    // The total number of segments sent, including those on current connections
    // but excluding those containing only retransmitted octets.  Discontinuities
    // in the value of this counter are indicated via discontinuities in the value
    // of sysUpTime. The type is interface{} with range: 0..4294967295.
    Tcpoutsegs interface{}

    // The total number of segments retransmitted; that is, the number of TCP
    // segments transmitted containing one or more previously transmitted octets. 
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    Tcpretranssegs interface{}

    // The total number of segments received in error (e.g., bad TCP checksums). 
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    Tcpinerrs interface{}

    // The number of TCP segments sent containing the RST flag.  Discontinuities
    // in the value of this counter are indicated via discontinuities in the value
    // of sysUpTime. The type is interface{} with range: 0..4294967295.
    Tcpoutrsts interface{}

    // The total number of segments received, including those received in error. 
    // This count includes segments received  on currently established
    // connections.  This object is the 64-bit equivalent of tcpInSegs. 
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..18446744073709551615.
    Tcphcinsegs interface{}

    // The total number of segments sent, including those on current connections
    // but excluding those containing only retransmitted octets.  This object is
    // the 64-bit equivalent of tcpOutSegs.  Discontinuities in the value of this
    // counter are indicated via discontinuities in the value of sysUpTime. The
    // type is interface{} with range: 0..18446744073709551615.
    Tcphcoutsegs interface{}
}

func (tcp *TCPMIB_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "TCP-MIB"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcp.EntityData.Children = make(map[string]types.YChild)
    tcp.EntityData.Leafs = make(map[string]types.YLeaf)
    tcp.EntityData.Leafs["tcpRtoAlgorithm"] = types.YLeaf{"Tcprtoalgorithm", tcp.Tcprtoalgorithm}
    tcp.EntityData.Leafs["tcpRtoMin"] = types.YLeaf{"Tcprtomin", tcp.Tcprtomin}
    tcp.EntityData.Leafs["tcpRtoMax"] = types.YLeaf{"Tcprtomax", tcp.Tcprtomax}
    tcp.EntityData.Leafs["tcpMaxConn"] = types.YLeaf{"Tcpmaxconn", tcp.Tcpmaxconn}
    tcp.EntityData.Leafs["tcpActiveOpens"] = types.YLeaf{"Tcpactiveopens", tcp.Tcpactiveopens}
    tcp.EntityData.Leafs["tcpPassiveOpens"] = types.YLeaf{"Tcppassiveopens", tcp.Tcppassiveopens}
    tcp.EntityData.Leafs["tcpAttemptFails"] = types.YLeaf{"Tcpattemptfails", tcp.Tcpattemptfails}
    tcp.EntityData.Leafs["tcpEstabResets"] = types.YLeaf{"Tcpestabresets", tcp.Tcpestabresets}
    tcp.EntityData.Leafs["tcpCurrEstab"] = types.YLeaf{"Tcpcurrestab", tcp.Tcpcurrestab}
    tcp.EntityData.Leafs["tcpInSegs"] = types.YLeaf{"Tcpinsegs", tcp.Tcpinsegs}
    tcp.EntityData.Leafs["tcpOutSegs"] = types.YLeaf{"Tcpoutsegs", tcp.Tcpoutsegs}
    tcp.EntityData.Leafs["tcpRetransSegs"] = types.YLeaf{"Tcpretranssegs", tcp.Tcpretranssegs}
    tcp.EntityData.Leafs["tcpInErrs"] = types.YLeaf{"Tcpinerrs", tcp.Tcpinerrs}
    tcp.EntityData.Leafs["tcpOutRsts"] = types.YLeaf{"Tcpoutrsts", tcp.Tcpoutrsts}
    tcp.EntityData.Leafs["tcpHCInSegs"] = types.YLeaf{"Tcphcinsegs", tcp.Tcphcinsegs}
    tcp.EntityData.Leafs["tcpHCOutSegs"] = types.YLeaf{"Tcphcoutsegs", tcp.Tcphcoutsegs}
    return &(tcp.EntityData)
}

// TCPMIB_Tcp_Tcprtoalgorithm represents retransmitting unacknowledged octets.
type TCPMIB_Tcp_Tcprtoalgorithm string

const (
    TCPMIB_Tcp_Tcprtoalgorithm_other TCPMIB_Tcp_Tcprtoalgorithm = "other"

    TCPMIB_Tcp_Tcprtoalgorithm_constant TCPMIB_Tcp_Tcprtoalgorithm = "constant"

    TCPMIB_Tcp_Tcprtoalgorithm_rsre TCPMIB_Tcp_Tcprtoalgorithm = "rsre"

    TCPMIB_Tcp_Tcprtoalgorithm_vanj TCPMIB_Tcp_Tcprtoalgorithm = "vanj"

    TCPMIB_Tcp_Tcprtoalgorithm_rfc2988 TCPMIB_Tcp_Tcprtoalgorithm = "rfc2988"
)

// TCPMIB_Tcpconntable
// A table containing information about existing IPv4-specific
// TCP connections or listeners.  This table has been
// deprecated in favor of the version neutral
// tcpConnectionTable.
type TCPMIB_Tcpconntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row of the tcpConnTable containing information about a
    // particular current IPv4 TCP connection.  Each row of this table is
    // transient in that it ceases to exist when (or soon after) the connection
    // makes the transition to the CLOSED state. The type is slice of
    // TCPMIB_Tcpconntable_Tcpconnentry.
    Tcpconnentry []TCPMIB_Tcpconntable_Tcpconnentry
}

func (tcpconntable *TCPMIB_Tcpconntable) GetEntityData() *types.CommonEntityData {
    tcpconntable.EntityData.YFilter = tcpconntable.YFilter
    tcpconntable.EntityData.YangName = "tcpConnTable"
    tcpconntable.EntityData.BundleName = "cisco_ios_xe"
    tcpconntable.EntityData.ParentYangName = "TCP-MIB"
    tcpconntable.EntityData.SegmentPath = "tcpConnTable"
    tcpconntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpconntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpconntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpconntable.EntityData.Children = make(map[string]types.YChild)
    tcpconntable.EntityData.Children["tcpConnEntry"] = types.YChild{"Tcpconnentry", nil}
    for i := range tcpconntable.Tcpconnentry {
        tcpconntable.EntityData.Children[types.GetSegmentPath(&tcpconntable.Tcpconnentry[i])] = types.YChild{"Tcpconnentry", &tcpconntable.Tcpconnentry[i]}
    }
    tcpconntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tcpconntable.EntityData)
}

// TCPMIB_Tcpconntable_Tcpconnentry
// A conceptual row of the tcpConnTable containing information
// about a particular current IPv4 TCP connection.  Each row
// of this table is transient in that it ceases to exist when
// (or soon after) the connection makes the transition to the
// CLOSED state.
type TCPMIB_Tcpconntable_Tcpconnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this TCP connection.  In
    // the case of a connection in the listen state willing to accept connections
    // for any IP interface associated with the node, the value 0.0.0.0 is used.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Tcpconnlocaladdress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    Tcpconnlocalport interface{}

    // This attribute is a key. The remote IP address for this TCP connection. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Tcpconnremaddress interface{}

    // This attribute is a key. The remote port number for this TCP connection.
    // The type is interface{} with range: 0..65535.
    Tcpconnremport interface{}

    // The state of this TCP connection.  The only value that may be set by a
    // management station is deleteTCB(12).  Accordingly, it is appropriate for an
    // agent to return a `badValue' response if a management station attempts to
    // set this object to any other value.  If a management station sets this
    // object to the value deleteTCB(12), then the TCB (as defined in [RFC793]) of
    // the corresponding connection on the managed node is deleted, resulting in
    // immediate termination of the connection.  As an implementation-specific
    // option, a RST segment may be sent from the managed node to the other TCP
    // endpoint (note, however, that RST segments are not sent reliably). The type
    // is Tcpconnstate.
    Tcpconnstate interface{}
}

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetEntityData() *types.CommonEntityData {
    tcpconnentry.EntityData.YFilter = tcpconnentry.YFilter
    tcpconnentry.EntityData.YangName = "tcpConnEntry"
    tcpconnentry.EntityData.BundleName = "cisco_ios_xe"
    tcpconnentry.EntityData.ParentYangName = "tcpConnTable"
    tcpconnentry.EntityData.SegmentPath = "tcpConnEntry" + "[tcpConnLocalAddress='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnlocaladdress) + "']" + "[tcpConnLocalPort='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnlocalport) + "']" + "[tcpConnRemAddress='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnremaddress) + "']" + "[tcpConnRemPort='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnremport) + "']"
    tcpconnentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpconnentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpconnentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpconnentry.EntityData.Children = make(map[string]types.YChild)
    tcpconnentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tcpconnentry.EntityData.Leafs["tcpConnLocalAddress"] = types.YLeaf{"Tcpconnlocaladdress", tcpconnentry.Tcpconnlocaladdress}
    tcpconnentry.EntityData.Leafs["tcpConnLocalPort"] = types.YLeaf{"Tcpconnlocalport", tcpconnentry.Tcpconnlocalport}
    tcpconnentry.EntityData.Leafs["tcpConnRemAddress"] = types.YLeaf{"Tcpconnremaddress", tcpconnentry.Tcpconnremaddress}
    tcpconnentry.EntityData.Leafs["tcpConnRemPort"] = types.YLeaf{"Tcpconnremport", tcpconnentry.Tcpconnremport}
    tcpconnentry.EntityData.Leafs["tcpConnState"] = types.YLeaf{"Tcpconnstate", tcpconnentry.Tcpconnstate}
    return &(tcpconnentry.EntityData)
}

// TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate represents however, that RST segments are not sent reliably).
type TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate string

const (
    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_closed TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "closed"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_listen TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "listen"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_synSent TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "synSent"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_synReceived TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "synReceived"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_established TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "established"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_finWait1 TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "finWait1"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_finWait2 TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "finWait2"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_closeWait TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "closeWait"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_lastAck TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "lastAck"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_closing TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "closing"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_timeWait TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "timeWait"

    TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate_deleteTCB TCPMIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "deleteTCB"
)

// TCPMIB_Tcpconnectiontable
// A table containing information about existing TCP
// connections.  Note that unlike earlier TCP MIBs, there
// is a separate table for connections in the LISTEN state.
type TCPMIB_Tcpconnectiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row of the tcpConnectionTable containing information about a
    // particular current TCP connection. Each row of this table is transient in
    // that it ceases to exist when (or soon after) the connection makes the
    // transition to the CLOSED state. The type is slice of
    // TCPMIB_Tcpconnectiontable_Tcpconnectionentry.
    Tcpconnectionentry []TCPMIB_Tcpconnectiontable_Tcpconnectionentry
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetEntityData() *types.CommonEntityData {
    tcpconnectiontable.EntityData.YFilter = tcpconnectiontable.YFilter
    tcpconnectiontable.EntityData.YangName = "tcpConnectionTable"
    tcpconnectiontable.EntityData.BundleName = "cisco_ios_xe"
    tcpconnectiontable.EntityData.ParentYangName = "TCP-MIB"
    tcpconnectiontable.EntityData.SegmentPath = "tcpConnectionTable"
    tcpconnectiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpconnectiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpconnectiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpconnectiontable.EntityData.Children = make(map[string]types.YChild)
    tcpconnectiontable.EntityData.Children["tcpConnectionEntry"] = types.YChild{"Tcpconnectionentry", nil}
    for i := range tcpconnectiontable.Tcpconnectionentry {
        tcpconnectiontable.EntityData.Children[types.GetSegmentPath(&tcpconnectiontable.Tcpconnectionentry[i])] = types.YChild{"Tcpconnectionentry", &tcpconnectiontable.Tcpconnectionentry[i]}
    }
    tcpconnectiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tcpconnectiontable.EntityData)
}

// TCPMIB_Tcpconnectiontable_Tcpconnectionentry
// A conceptual row of the tcpConnectionTable containing
// information about a particular current TCP connection.
// Each row of this table is transient in that it ceases to
// exist when (or soon after) the connection makes the
// transition to the CLOSED state.
type TCPMIB_Tcpconnectiontable_Tcpconnectionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The address type of tcpConnectionLocalAddress. The
    // type is InetAddressType.
    Tcpconnectionlocaladdresstype interface{}

    // This attribute is a key. The local IP address for this TCP connection.  The
    // type of this address is determined by the value of
    // tcpConnectionLocalAddressType.  As this object is used in the index for the
    // tcpConnectionTable, implementors should be careful not to create entries
    // that would result in OIDs with more than 128 subidentifiers; otherwise the
    // information cannot be accessed by using SNMPv1, SNMPv2c, or SNMPv3. The
    // type is string with length: 0..255.
    Tcpconnectionlocaladdress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    Tcpconnectionlocalport interface{}

    // This attribute is a key. The address type of tcpConnectionRemAddress. The
    // type is InetAddressType.
    Tcpconnectionremaddresstype interface{}

    // This attribute is a key. The remote IP address for this TCP connection. 
    // The type of this address is determined by the value of
    // tcpConnectionRemAddressType.  As this object is used in the index for the
    // tcpConnectionTable, implementors should be careful not to create entries
    // that would result in OIDs with more than 128 subidentifiers; otherwise the
    // information cannot be accessed by using SNMPv1, SNMPv2c, or SNMPv3. The
    // type is string with length: 0..255.
    Tcpconnectionremaddress interface{}

    // This attribute is a key. The remote port number for this TCP connection.
    // The type is interface{} with range: 0..65535.
    Tcpconnectionremport interface{}

    // The state of this TCP connection.  The value listen(2) is included only for
    // parallelism to the old tcpConnTable and should not be used.  A connection
    // in LISTEN state should be present in the tcpListenerTable.  The only value
    // that may be set by a management station is deleteTCB(12).  Accordingly, it
    // is appropriate for an agent to return a `badValue' response if a management
    // station attempts to set this object to any other value.  If a management
    // station sets this object to the value deleteTCB(12), then the TCB (as
    // defined in [RFC793]) of the corresponding connection on the managed node is
    // deleted, resulting in immediate termination of the connection.  As an
    // implementation-specific option, a RST segment may be sent from the managed
    // node to the other TCP endpoint (note, however, that RST segments are not
    // sent reliably). The type is Tcpconnectionstate.
    Tcpconnectionstate interface{}

    // The system's process ID for the process associated with this connection, or
    // zero if there is no such process.  This value is expected to be the same as
    // HOST-RESOURCES-MIB:: hrSWRunIndex or SYSAPPL-MIB::sysApplElmtRunIndex for
    // some row in the appropriate tables. The type is interface{} with range:
    // 0..4294967295.
    Tcpconnectionprocess interface{}
}

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetEntityData() *types.CommonEntityData {
    tcpconnectionentry.EntityData.YFilter = tcpconnectionentry.YFilter
    tcpconnectionentry.EntityData.YangName = "tcpConnectionEntry"
    tcpconnectionentry.EntityData.BundleName = "cisco_ios_xe"
    tcpconnectionentry.EntityData.ParentYangName = "tcpConnectionTable"
    tcpconnectionentry.EntityData.SegmentPath = "tcpConnectionEntry" + "[tcpConnectionLocalAddressType='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionlocaladdresstype) + "']" + "[tcpConnectionLocalAddress='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionlocaladdress) + "']" + "[tcpConnectionLocalPort='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionlocalport) + "']" + "[tcpConnectionRemAddressType='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionremaddresstype) + "']" + "[tcpConnectionRemAddress='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionremaddress) + "']" + "[tcpConnectionRemPort='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionremport) + "']"
    tcpconnectionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpconnectionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpconnectionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpconnectionentry.EntityData.Children = make(map[string]types.YChild)
    tcpconnectionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tcpconnectionentry.EntityData.Leafs["tcpConnectionLocalAddressType"] = types.YLeaf{"Tcpconnectionlocaladdresstype", tcpconnectionentry.Tcpconnectionlocaladdresstype}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionLocalAddress"] = types.YLeaf{"Tcpconnectionlocaladdress", tcpconnectionentry.Tcpconnectionlocaladdress}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionLocalPort"] = types.YLeaf{"Tcpconnectionlocalport", tcpconnectionentry.Tcpconnectionlocalport}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionRemAddressType"] = types.YLeaf{"Tcpconnectionremaddresstype", tcpconnectionentry.Tcpconnectionremaddresstype}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionRemAddress"] = types.YLeaf{"Tcpconnectionremaddress", tcpconnectionentry.Tcpconnectionremaddress}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionRemPort"] = types.YLeaf{"Tcpconnectionremport", tcpconnectionentry.Tcpconnectionremport}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionState"] = types.YLeaf{"Tcpconnectionstate", tcpconnectionentry.Tcpconnectionstate}
    tcpconnectionentry.EntityData.Leafs["tcpConnectionProcess"] = types.YLeaf{"Tcpconnectionprocess", tcpconnectionentry.Tcpconnectionprocess}
    return &(tcpconnectionentry.EntityData)
}

// TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate represents however, that RST segments are not sent reliably).
type TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate string

const (
    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_closed TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "closed"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_listen TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "listen"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_synSent TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "synSent"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_synReceived TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "synReceived"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_established TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "established"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_finWait1 TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "finWait1"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_finWait2 TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "finWait2"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_closeWait TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "closeWait"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_lastAck TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "lastAck"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_closing TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "closing"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_timeWait TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "timeWait"

    TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate_deleteTCB TCPMIB_Tcpconnectiontable_Tcpconnectionentry_Tcpconnectionstate = "deleteTCB"
)

// TCPMIB_Tcplistenertable
// A table containing information about TCP listeners.  A
// listening application can be represented in three
// possible ways:
// 
// 1. An application that is willing to accept both IPv4 and
//    IPv6 datagrams is represented by
// 
//    a tcpListenerLocalAddressType of unknown (0) and
//    a tcpListenerLocalAddress of ''h (a zero-length
//    octet-string).
// 
// 2. An application that is willing to accept only IPv4 or
//    IPv6 datagrams is represented by a
//    tcpListenerLocalAddressType of the appropriate address
//    type and a tcpListenerLocalAddress of '0.0.0.0' or '::'
//    respectively.
// 
// 3. An application that is listening for data destined
//    only to a specific IP address, but from any remote
//    system, is represented by a tcpListenerLocalAddressType
//    of an appropriate address type, with
//    tcpListenerLocalAddress as the specific local address.
// 
// NOTE: The address type in this table represents the
// address type used for the communication, irrespective
// of the higher-layer abstraction.  For example, an
// application using IPv6 'sockets' to communicate via
// IPv4 between ::ffff:10.0.0.1 and ::ffff:10.0.0.2 would
// use InetAddressType ipv4(1)).
type TCPMIB_Tcplistenertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row of the tcpListenerTable containing information about a
    // particular TCP listener. The type is slice of
    // TCPMIB_Tcplistenertable_Tcplistenerentry.
    Tcplistenerentry []TCPMIB_Tcplistenertable_Tcplistenerentry
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetEntityData() *types.CommonEntityData {
    tcplistenertable.EntityData.YFilter = tcplistenertable.YFilter
    tcplistenertable.EntityData.YangName = "tcpListenerTable"
    tcplistenertable.EntityData.BundleName = "cisco_ios_xe"
    tcplistenertable.EntityData.ParentYangName = "TCP-MIB"
    tcplistenertable.EntityData.SegmentPath = "tcpListenerTable"
    tcplistenertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcplistenertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcplistenertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcplistenertable.EntityData.Children = make(map[string]types.YChild)
    tcplistenertable.EntityData.Children["tcpListenerEntry"] = types.YChild{"Tcplistenerentry", nil}
    for i := range tcplistenertable.Tcplistenerentry {
        tcplistenertable.EntityData.Children[types.GetSegmentPath(&tcplistenertable.Tcplistenerentry[i])] = types.YChild{"Tcplistenerentry", &tcplistenertable.Tcplistenerentry[i]}
    }
    tcplistenertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tcplistenertable.EntityData)
}

// TCPMIB_Tcplistenertable_Tcplistenerentry
// A conceptual row of the tcpListenerTable containing
// information about a particular TCP listener.
type TCPMIB_Tcplistenertable_Tcplistenerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The address type of tcpListenerLocalAddress.  The
    // value should be unknown (0) if connection initiations to all local IP
    // addresses are accepted. The type is InetAddressType.
    Tcplistenerlocaladdresstype interface{}

    // This attribute is a key. The local IP address for this TCP connection.  The
    // value of this object can be represented in three possible ways, depending
    // on the characteristics of the listening application:  1. For an application
    // willing to accept both IPv4 and    IPv6 datagrams, the value of this object
    // must be    ''h (a zero-length octet-string), with the value    of the
    // corresponding tcpListenerLocalAddressType    object being unknown (0).  2.
    // For an application willing to accept only IPv4 or    IPv6 datagrams, the
    // value of this object must be    '0.0.0.0' or '::' respectively, with   
    // tcpListenerLocalAddressType representing the    appropriate address type. 
    // 3. For an application which is listening for data    destined only to a
    // specific IP address, the value    of this object is the specific local
    // address, with    tcpListenerLocalAddressType representing the   
    // appropriate address type.  As this object is used in the index for the
    // tcpListenerTable, implementors should be careful not to create entries that
    // would result in OIDs with more than 128 subidentifiers; otherwise the
    // information cannot be accessed, using SNMPv1, SNMPv2c, or SNMPv3. The type
    // is string with length: 0..255.
    Tcplistenerlocaladdress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    Tcplistenerlocalport interface{}

    // The system's process ID for the process associated with this listener, or
    // zero if there is no such process.  This value is expected to be the same as
    // HOST-RESOURCES-MIB:: hrSWRunIndex or SYSAPPL-MIB::sysApplElmtRunIndex for
    // some row in the appropriate tables. The type is interface{} with range:
    // 0..4294967295.
    Tcplistenerprocess interface{}
}

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetEntityData() *types.CommonEntityData {
    tcplistenerentry.EntityData.YFilter = tcplistenerentry.YFilter
    tcplistenerentry.EntityData.YangName = "tcpListenerEntry"
    tcplistenerentry.EntityData.BundleName = "cisco_ios_xe"
    tcplistenerentry.EntityData.ParentYangName = "tcpListenerTable"
    tcplistenerentry.EntityData.SegmentPath = "tcpListenerEntry" + "[tcpListenerLocalAddressType='" + fmt.Sprintf("%v", tcplistenerentry.Tcplistenerlocaladdresstype) + "']" + "[tcpListenerLocalAddress='" + fmt.Sprintf("%v", tcplistenerentry.Tcplistenerlocaladdress) + "']" + "[tcpListenerLocalPort='" + fmt.Sprintf("%v", tcplistenerentry.Tcplistenerlocalport) + "']"
    tcplistenerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcplistenerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcplistenerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcplistenerentry.EntityData.Children = make(map[string]types.YChild)
    tcplistenerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tcplistenerentry.EntityData.Leafs["tcpListenerLocalAddressType"] = types.YLeaf{"Tcplistenerlocaladdresstype", tcplistenerentry.Tcplistenerlocaladdresstype}
    tcplistenerentry.EntityData.Leafs["tcpListenerLocalAddress"] = types.YLeaf{"Tcplistenerlocaladdress", tcplistenerentry.Tcplistenerlocaladdress}
    tcplistenerentry.EntityData.Leafs["tcpListenerLocalPort"] = types.YLeaf{"Tcplistenerlocalport", tcplistenerentry.Tcplistenerlocalport}
    tcplistenerentry.EntityData.Leafs["tcpListenerProcess"] = types.YLeaf{"Tcplistenerprocess", tcplistenerentry.Tcplistenerprocess}
    return &(tcplistenerentry.EntityData)
}

