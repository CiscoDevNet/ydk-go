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
    parent types.Entity
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

func (tCPMIB *TCPMIB) GetFilter() yfilter.YFilter { return tCPMIB.YFilter }

func (tCPMIB *TCPMIB) SetFilter(yf yfilter.YFilter) { tCPMIB.YFilter = yf }

func (tCPMIB *TCPMIB) GetGoName(yname string) string {
    if yname == "tcp" { return "Tcp" }
    if yname == "tcpConnTable" { return "Tcpconntable" }
    if yname == "tcpConnectionTable" { return "Tcpconnectiontable" }
    if yname == "tcpListenerTable" { return "Tcplistenertable" }
    return ""
}

func (tCPMIB *TCPMIB) GetSegmentPath() string {
    return "TCP-MIB:TCP-MIB"
}

func (tCPMIB *TCPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp" {
        return &tCPMIB.Tcp
    }
    if childYangName == "tcpConnTable" {
        return &tCPMIB.Tcpconntable
    }
    if childYangName == "tcpConnectionTable" {
        return &tCPMIB.Tcpconnectiontable
    }
    if childYangName == "tcpListenerTable" {
        return &tCPMIB.Tcplistenertable
    }
    return nil
}

func (tCPMIB *TCPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcp"] = &tCPMIB.Tcp
    children["tcpConnTable"] = &tCPMIB.Tcpconntable
    children["tcpConnectionTable"] = &tCPMIB.Tcpconnectiontable
    children["tcpListenerTable"] = &tCPMIB.Tcplistenertable
    return children
}

func (tCPMIB *TCPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tCPMIB *TCPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (tCPMIB *TCPMIB) GetYangName() string { return "TCP-MIB" }

func (tCPMIB *TCPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tCPMIB *TCPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tCPMIB *TCPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tCPMIB *TCPMIB) SetParent(parent types.Entity) { tCPMIB.parent = parent }

func (tCPMIB *TCPMIB) GetParent() types.Entity { return tCPMIB.parent }

func (tCPMIB *TCPMIB) GetParentYangName() string { return "TCP-MIB" }

// TCPMIB_Tcp
type TCPMIB_Tcp struct {
    parent types.Entity
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

func (tcp *TCPMIB_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *TCPMIB_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *TCPMIB_Tcp) GetGoName(yname string) string {
    if yname == "tcpRtoAlgorithm" { return "Tcprtoalgorithm" }
    if yname == "tcpRtoMin" { return "Tcprtomin" }
    if yname == "tcpRtoMax" { return "Tcprtomax" }
    if yname == "tcpMaxConn" { return "Tcpmaxconn" }
    if yname == "tcpActiveOpens" { return "Tcpactiveopens" }
    if yname == "tcpPassiveOpens" { return "Tcppassiveopens" }
    if yname == "tcpAttemptFails" { return "Tcpattemptfails" }
    if yname == "tcpEstabResets" { return "Tcpestabresets" }
    if yname == "tcpCurrEstab" { return "Tcpcurrestab" }
    if yname == "tcpInSegs" { return "Tcpinsegs" }
    if yname == "tcpOutSegs" { return "Tcpoutsegs" }
    if yname == "tcpRetransSegs" { return "Tcpretranssegs" }
    if yname == "tcpInErrs" { return "Tcpinerrs" }
    if yname == "tcpOutRsts" { return "Tcpoutrsts" }
    if yname == "tcpHCInSegs" { return "Tcphcinsegs" }
    if yname == "tcpHCOutSegs" { return "Tcphcoutsegs" }
    return ""
}

func (tcp *TCPMIB_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *TCPMIB_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *TCPMIB_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *TCPMIB_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcpRtoAlgorithm"] = tcp.Tcprtoalgorithm
    leafs["tcpRtoMin"] = tcp.Tcprtomin
    leafs["tcpRtoMax"] = tcp.Tcprtomax
    leafs["tcpMaxConn"] = tcp.Tcpmaxconn
    leafs["tcpActiveOpens"] = tcp.Tcpactiveopens
    leafs["tcpPassiveOpens"] = tcp.Tcppassiveopens
    leafs["tcpAttemptFails"] = tcp.Tcpattemptfails
    leafs["tcpEstabResets"] = tcp.Tcpestabresets
    leafs["tcpCurrEstab"] = tcp.Tcpcurrestab
    leafs["tcpInSegs"] = tcp.Tcpinsegs
    leafs["tcpOutSegs"] = tcp.Tcpoutsegs
    leafs["tcpRetransSegs"] = tcp.Tcpretranssegs
    leafs["tcpInErrs"] = tcp.Tcpinerrs
    leafs["tcpOutRsts"] = tcp.Tcpoutrsts
    leafs["tcpHCInSegs"] = tcp.Tcphcinsegs
    leafs["tcpHCOutSegs"] = tcp.Tcphcoutsegs
    return leafs
}

func (tcp *TCPMIB_Tcp) GetBundleName() string { return "cisco_ios_xe" }

func (tcp *TCPMIB_Tcp) GetYangName() string { return "tcp" }

func (tcp *TCPMIB_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcp *TCPMIB_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcp *TCPMIB_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcp *TCPMIB_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *TCPMIB_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *TCPMIB_Tcp) GetParentYangName() string { return "TCP-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row of the tcpConnTable containing information about a
    // particular current IPv4 TCP connection.  Each row of this table is
    // transient in that it ceases to exist when (or soon after) the connection
    // makes the transition to the CLOSED state. The type is slice of
    // TCPMIB_Tcpconntable_Tcpconnentry.
    Tcpconnentry []TCPMIB_Tcpconntable_Tcpconnentry
}

func (tcpconntable *TCPMIB_Tcpconntable) GetFilter() yfilter.YFilter { return tcpconntable.YFilter }

func (tcpconntable *TCPMIB_Tcpconntable) SetFilter(yf yfilter.YFilter) { tcpconntable.YFilter = yf }

func (tcpconntable *TCPMIB_Tcpconntable) GetGoName(yname string) string {
    if yname == "tcpConnEntry" { return "Tcpconnentry" }
    return ""
}

func (tcpconntable *TCPMIB_Tcpconntable) GetSegmentPath() string {
    return "tcpConnTable"
}

func (tcpconntable *TCPMIB_Tcpconntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcpConnEntry" {
        for _, c := range tcpconntable.Tcpconnentry {
            if tcpconntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TCPMIB_Tcpconntable_Tcpconnentry{}
        tcpconntable.Tcpconnentry = append(tcpconntable.Tcpconnentry, child)
        return &tcpconntable.Tcpconnentry[len(tcpconntable.Tcpconnentry)-1]
    }
    return nil
}

func (tcpconntable *TCPMIB_Tcpconntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tcpconntable.Tcpconnentry {
        children[tcpconntable.Tcpconnentry[i].GetSegmentPath()] = &tcpconntable.Tcpconnentry[i]
    }
    return children
}

func (tcpconntable *TCPMIB_Tcpconntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcpconntable *TCPMIB_Tcpconntable) GetBundleName() string { return "cisco_ios_xe" }

func (tcpconntable *TCPMIB_Tcpconntable) GetYangName() string { return "tcpConnTable" }

func (tcpconntable *TCPMIB_Tcpconntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpconntable *TCPMIB_Tcpconntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpconntable *TCPMIB_Tcpconntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpconntable *TCPMIB_Tcpconntable) SetParent(parent types.Entity) { tcpconntable.parent = parent }

func (tcpconntable *TCPMIB_Tcpconntable) GetParent() types.Entity { return tcpconntable.parent }

func (tcpconntable *TCPMIB_Tcpconntable) GetParentYangName() string { return "TCP-MIB" }

// TCPMIB_Tcpconntable_Tcpconnentry
// A conceptual row of the tcpConnTable containing information
// about a particular current IPv4 TCP connection.  Each row
// of this table is transient in that it ceases to exist when
// (or soon after) the connection makes the transition to the
// CLOSED state.
type TCPMIB_Tcpconntable_Tcpconnentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this TCP connection.  In
    // the case of a connection in the listen state willing to accept connections
    // for any IP interface associated with the node, the value 0.0.0.0 is used.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Tcpconnlocaladdress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    Tcpconnlocalport interface{}

    // This attribute is a key. The remote IP address for this TCP connection. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetFilter() yfilter.YFilter { return tcpconnentry.YFilter }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) SetFilter(yf yfilter.YFilter) { tcpconnentry.YFilter = yf }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetGoName(yname string) string {
    if yname == "tcpConnLocalAddress" { return "Tcpconnlocaladdress" }
    if yname == "tcpConnLocalPort" { return "Tcpconnlocalport" }
    if yname == "tcpConnRemAddress" { return "Tcpconnremaddress" }
    if yname == "tcpConnRemPort" { return "Tcpconnremport" }
    if yname == "tcpConnState" { return "Tcpconnstate" }
    return ""
}

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetSegmentPath() string {
    return "tcpConnEntry" + "[tcpConnLocalAddress='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnlocaladdress) + "']" + "[tcpConnLocalPort='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnlocalport) + "']" + "[tcpConnRemAddress='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnremaddress) + "']" + "[tcpConnRemPort='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnremport) + "']"
}

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcpConnLocalAddress"] = tcpconnentry.Tcpconnlocaladdress
    leafs["tcpConnLocalPort"] = tcpconnentry.Tcpconnlocalport
    leafs["tcpConnRemAddress"] = tcpconnentry.Tcpconnremaddress
    leafs["tcpConnRemPort"] = tcpconnentry.Tcpconnremport
    leafs["tcpConnState"] = tcpconnentry.Tcpconnstate
    return leafs
}

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetBundleName() string { return "cisco_ios_xe" }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetYangName() string { return "tcpConnEntry" }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) SetParent(parent types.Entity) { tcpconnentry.parent = parent }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetParent() types.Entity { return tcpconnentry.parent }

func (tcpconnentry *TCPMIB_Tcpconntable_Tcpconnentry) GetParentYangName() string { return "tcpConnTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row of the tcpConnectionTable containing information about a
    // particular current TCP connection. Each row of this table is transient in
    // that it ceases to exist when (or soon after) the connection makes the
    // transition to the CLOSED state. The type is slice of
    // TCPMIB_Tcpconnectiontable_Tcpconnectionentry.
    Tcpconnectionentry []TCPMIB_Tcpconnectiontable_Tcpconnectionentry
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetFilter() yfilter.YFilter { return tcpconnectiontable.YFilter }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) SetFilter(yf yfilter.YFilter) { tcpconnectiontable.YFilter = yf }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetGoName(yname string) string {
    if yname == "tcpConnectionEntry" { return "Tcpconnectionentry" }
    return ""
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetSegmentPath() string {
    return "tcpConnectionTable"
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcpConnectionEntry" {
        for _, c := range tcpconnectiontable.Tcpconnectionentry {
            if tcpconnectiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TCPMIB_Tcpconnectiontable_Tcpconnectionentry{}
        tcpconnectiontable.Tcpconnectionentry = append(tcpconnectiontable.Tcpconnectionentry, child)
        return &tcpconnectiontable.Tcpconnectionentry[len(tcpconnectiontable.Tcpconnectionentry)-1]
    }
    return nil
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tcpconnectiontable.Tcpconnectionentry {
        children[tcpconnectiontable.Tcpconnectionentry[i].GetSegmentPath()] = &tcpconnectiontable.Tcpconnectionentry[i]
    }
    return children
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetBundleName() string { return "cisco_ios_xe" }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetYangName() string { return "tcpConnectionTable" }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) SetParent(parent types.Entity) { tcpconnectiontable.parent = parent }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetParent() types.Entity { return tcpconnectiontable.parent }

func (tcpconnectiontable *TCPMIB_Tcpconnectiontable) GetParentYangName() string { return "TCP-MIB" }

// TCPMIB_Tcpconnectiontable_Tcpconnectionentry
// A conceptual row of the tcpConnectionTable containing
// information about a particular current TCP connection.
// Each row of this table is transient in that it ceases to
// exist when (or soon after) the connection makes the
// transition to the CLOSED state.
type TCPMIB_Tcpconnectiontable_Tcpconnectionentry struct {
    parent types.Entity
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

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetFilter() yfilter.YFilter { return tcpconnectionentry.YFilter }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) SetFilter(yf yfilter.YFilter) { tcpconnectionentry.YFilter = yf }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetGoName(yname string) string {
    if yname == "tcpConnectionLocalAddressType" { return "Tcpconnectionlocaladdresstype" }
    if yname == "tcpConnectionLocalAddress" { return "Tcpconnectionlocaladdress" }
    if yname == "tcpConnectionLocalPort" { return "Tcpconnectionlocalport" }
    if yname == "tcpConnectionRemAddressType" { return "Tcpconnectionremaddresstype" }
    if yname == "tcpConnectionRemAddress" { return "Tcpconnectionremaddress" }
    if yname == "tcpConnectionRemPort" { return "Tcpconnectionremport" }
    if yname == "tcpConnectionState" { return "Tcpconnectionstate" }
    if yname == "tcpConnectionProcess" { return "Tcpconnectionprocess" }
    return ""
}

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetSegmentPath() string {
    return "tcpConnectionEntry" + "[tcpConnectionLocalAddressType='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionlocaladdresstype) + "']" + "[tcpConnectionLocalAddress='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionlocaladdress) + "']" + "[tcpConnectionLocalPort='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionlocalport) + "']" + "[tcpConnectionRemAddressType='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionremaddresstype) + "']" + "[tcpConnectionRemAddress='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionremaddress) + "']" + "[tcpConnectionRemPort='" + fmt.Sprintf("%v", tcpconnectionentry.Tcpconnectionremport) + "']"
}

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcpConnectionLocalAddressType"] = tcpconnectionentry.Tcpconnectionlocaladdresstype
    leafs["tcpConnectionLocalAddress"] = tcpconnectionentry.Tcpconnectionlocaladdress
    leafs["tcpConnectionLocalPort"] = tcpconnectionentry.Tcpconnectionlocalport
    leafs["tcpConnectionRemAddressType"] = tcpconnectionentry.Tcpconnectionremaddresstype
    leafs["tcpConnectionRemAddress"] = tcpconnectionentry.Tcpconnectionremaddress
    leafs["tcpConnectionRemPort"] = tcpconnectionentry.Tcpconnectionremport
    leafs["tcpConnectionState"] = tcpconnectionentry.Tcpconnectionstate
    leafs["tcpConnectionProcess"] = tcpconnectionentry.Tcpconnectionprocess
    return leafs
}

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetBundleName() string { return "cisco_ios_xe" }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetYangName() string { return "tcpConnectionEntry" }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) SetParent(parent types.Entity) { tcpconnectionentry.parent = parent }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetParent() types.Entity { return tcpconnectionentry.parent }

func (tcpconnectionentry *TCPMIB_Tcpconnectiontable_Tcpconnectionentry) GetParentYangName() string { return "tcpConnectionTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row of the tcpListenerTable containing information about a
    // particular TCP listener. The type is slice of
    // TCPMIB_Tcplistenertable_Tcplistenerentry.
    Tcplistenerentry []TCPMIB_Tcplistenertable_Tcplistenerentry
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetFilter() yfilter.YFilter { return tcplistenertable.YFilter }

func (tcplistenertable *TCPMIB_Tcplistenertable) SetFilter(yf yfilter.YFilter) { tcplistenertable.YFilter = yf }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetGoName(yname string) string {
    if yname == "tcpListenerEntry" { return "Tcplistenerentry" }
    return ""
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetSegmentPath() string {
    return "tcpListenerTable"
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcpListenerEntry" {
        for _, c := range tcplistenertable.Tcplistenerentry {
            if tcplistenertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TCPMIB_Tcplistenertable_Tcplistenerentry{}
        tcplistenertable.Tcplistenerentry = append(tcplistenertable.Tcplistenerentry, child)
        return &tcplistenertable.Tcplistenerentry[len(tcplistenertable.Tcplistenerentry)-1]
    }
    return nil
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tcplistenertable.Tcplistenerentry {
        children[tcplistenertable.Tcplistenerentry[i].GetSegmentPath()] = &tcplistenertable.Tcplistenerentry[i]
    }
    return children
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcplistenertable *TCPMIB_Tcplistenertable) GetBundleName() string { return "cisco_ios_xe" }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetYangName() string { return "tcpListenerTable" }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcplistenertable *TCPMIB_Tcplistenertable) SetParent(parent types.Entity) { tcplistenertable.parent = parent }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetParent() types.Entity { return tcplistenertable.parent }

func (tcplistenertable *TCPMIB_Tcplistenertable) GetParentYangName() string { return "TCP-MIB" }

// TCPMIB_Tcplistenertable_Tcplistenerentry
// A conceptual row of the tcpListenerTable containing
// information about a particular TCP listener.
type TCPMIB_Tcplistenertable_Tcplistenerentry struct {
    parent types.Entity
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

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetFilter() yfilter.YFilter { return tcplistenerentry.YFilter }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) SetFilter(yf yfilter.YFilter) { tcplistenerentry.YFilter = yf }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetGoName(yname string) string {
    if yname == "tcpListenerLocalAddressType" { return "Tcplistenerlocaladdresstype" }
    if yname == "tcpListenerLocalAddress" { return "Tcplistenerlocaladdress" }
    if yname == "tcpListenerLocalPort" { return "Tcplistenerlocalport" }
    if yname == "tcpListenerProcess" { return "Tcplistenerprocess" }
    return ""
}

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetSegmentPath() string {
    return "tcpListenerEntry" + "[tcpListenerLocalAddressType='" + fmt.Sprintf("%v", tcplistenerentry.Tcplistenerlocaladdresstype) + "']" + "[tcpListenerLocalAddress='" + fmt.Sprintf("%v", tcplistenerentry.Tcplistenerlocaladdress) + "']" + "[tcpListenerLocalPort='" + fmt.Sprintf("%v", tcplistenerentry.Tcplistenerlocalport) + "']"
}

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcpListenerLocalAddressType"] = tcplistenerentry.Tcplistenerlocaladdresstype
    leafs["tcpListenerLocalAddress"] = tcplistenerentry.Tcplistenerlocaladdress
    leafs["tcpListenerLocalPort"] = tcplistenerentry.Tcplistenerlocalport
    leafs["tcpListenerProcess"] = tcplistenerentry.Tcplistenerprocess
    return leafs
}

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetBundleName() string { return "cisco_ios_xe" }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetYangName() string { return "tcpListenerEntry" }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) SetParent(parent types.Entity) { tcplistenerentry.parent = parent }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetParent() types.Entity { return tcplistenerentry.parent }

func (tcplistenerentry *TCPMIB_Tcplistenertable_Tcplistenerentry) GetParentYangName() string { return "tcpListenerTable" }

