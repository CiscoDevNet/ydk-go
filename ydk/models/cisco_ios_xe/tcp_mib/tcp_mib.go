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
    TcpConnTable TCPMIB_TcpConnTable

    // A table containing information about existing TCP connections.  Note that
    // unlike earlier TCP MIBs, there is a separate table for connections in the
    // LISTEN state.
    TcpConnectionTable TCPMIB_TcpConnectionTable

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
    TcpListenerTable TCPMIB_TcpListenerTable
}

func (tCPMIB *TCPMIB) GetEntityData() *types.CommonEntityData {
    tCPMIB.EntityData.YFilter = tCPMIB.YFilter
    tCPMIB.EntityData.YangName = "TCP-MIB"
    tCPMIB.EntityData.BundleName = "cisco_ios_xe"
    tCPMIB.EntityData.ParentYangName = "TCP-MIB"
    tCPMIB.EntityData.SegmentPath = "TCP-MIB:TCP-MIB"
    tCPMIB.EntityData.AbsolutePath = tCPMIB.EntityData.SegmentPath
    tCPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tCPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tCPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tCPMIB.EntityData.Children = types.NewOrderedMap()
    tCPMIB.EntityData.Children.Append("tcp", types.YChild{"Tcp", &tCPMIB.Tcp})
    tCPMIB.EntityData.Children.Append("tcpConnTable", types.YChild{"TcpConnTable", &tCPMIB.TcpConnTable})
    tCPMIB.EntityData.Children.Append("tcpConnectionTable", types.YChild{"TcpConnectionTable", &tCPMIB.TcpConnectionTable})
    tCPMIB.EntityData.Children.Append("tcpListenerTable", types.YChild{"TcpListenerTable", &tCPMIB.TcpListenerTable})
    tCPMIB.EntityData.Leafs = types.NewOrderedMap()

    tCPMIB.EntityData.YListKeys = []string {}

    return &(tCPMIB.EntityData)
}

// TCPMIB_Tcp
type TCPMIB_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The algorithm used to determine the timeout value used for retransmitting
    // unacknowledged octets. The type is TcpRtoAlgorithm.
    TcpRtoAlgorithm interface{}

    // The minimum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds. More refined semantics for objects of
    // this type depend on the algorithm used to determine the retransmission
    // timeout; in particular, the IETF standard algorithm rfc2988(5) provides a
    // minimum value. The type is interface{} with range: 0..2147483647. Units are
    // milliseconds.
    TcpRtoMin interface{}

    // The maximum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds. More refined semantics for objects of
    // this type depend on the algorithm used to determine the retransmission
    // timeout; in particular, the IETF standard algorithm rfc2988(5) provides an
    // upper bound (as part of an adaptive backoff algorithm). The type is
    // interface{} with range: 0..2147483647. Units are milliseconds.
    TcpRtoMax interface{}

    // The limit on the total number of TCP connections the entity can support. 
    // In entities where the maximum number of connections is dynamic, this object
    // should contain the value -1. The type is interface{} with range:
    // -1..2147483647.
    TcpMaxConn interface{}

    // The number of times that TCP connections have made a direct transition to
    // the SYN-SENT state from the CLOSED state.  Discontinuities in the value of
    // this counter are indicated via discontinuities in the value of sysUpTime.
    // The type is interface{} with range: 0..4294967295.
    TcpActiveOpens interface{}

    // The number of times TCP connections have made a direct transition to the
    // SYN-RCVD state from the LISTEN state.  Discontinuities in the value of this
    // counter are indicated via discontinuities in the value of sysUpTime. The
    // type is interface{} with range: 0..4294967295.
    TcpPassiveOpens interface{}

    // The number of times that TCP connections have made a direct transition to
    // the CLOSED state from either the SYN-SENT state or the SYN-RCVD state, plus
    // the number of times that TCP connections have made a direct transition to
    // the LISTEN state from the SYN-RCVD state.  Discontinuities in the value of
    // this counter are indicated via discontinuities in the value of sysUpTime.
    // The type is interface{} with range: 0..4294967295.
    TcpAttemptFails interface{}

    // The number of times that TCP connections have made a direct transition to
    // the CLOSED state from either the ESTABLISHED state or the CLOSE-WAIT state.
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    TcpEstabResets interface{}

    // The number of TCP connections for which the current state is either
    // ESTABLISHED or CLOSE-WAIT. The type is interface{} with range:
    // 0..4294967295.
    TcpCurrEstab interface{}

    // The total number of segments received, including those received in error. 
    // This count includes segments received on currently established connections.
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    TcpInSegs interface{}

    // The total number of segments sent, including those on current connections
    // but excluding those containing only retransmitted octets.  Discontinuities
    // in the value of this counter are indicated via discontinuities in the value
    // of sysUpTime. The type is interface{} with range: 0..4294967295.
    TcpOutSegs interface{}

    // The total number of segments retransmitted; that is, the number of TCP
    // segments transmitted containing one or more previously transmitted octets. 
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    TcpRetransSegs interface{}

    // The total number of segments received in error (e.g., bad TCP checksums). 
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    TcpInErrs interface{}

    // The number of TCP segments sent containing the RST flag.  Discontinuities
    // in the value of this counter are indicated via discontinuities in the value
    // of sysUpTime. The type is interface{} with range: 0..4294967295.
    TcpOutRsts interface{}

    // The total number of segments received, including those received in error. 
    // This count includes segments received  on currently established
    // connections.  This object is the 64-bit equivalent of tcpInSegs. 
    // Discontinuities in the value of this counter are indicated via
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..18446744073709551615.
    TcpHCInSegs interface{}

    // The total number of segments sent, including those on current connections
    // but excluding those containing only retransmitted octets.  This object is
    // the 64-bit equivalent of tcpOutSegs.  Discontinuities in the value of this
    // counter are indicated via discontinuities in the value of sysUpTime. The
    // type is interface{} with range: 0..18446744073709551615.
    TcpHCOutSegs interface{}
}

func (tcp *TCPMIB_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "TCP-MIB"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/" + tcp.EntityData.SegmentPath
    tcp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcp.EntityData.Children = types.NewOrderedMap()
    tcp.EntityData.Leafs = types.NewOrderedMap()
    tcp.EntityData.Leafs.Append("tcpRtoAlgorithm", types.YLeaf{"TcpRtoAlgorithm", tcp.TcpRtoAlgorithm})
    tcp.EntityData.Leafs.Append("tcpRtoMin", types.YLeaf{"TcpRtoMin", tcp.TcpRtoMin})
    tcp.EntityData.Leafs.Append("tcpRtoMax", types.YLeaf{"TcpRtoMax", tcp.TcpRtoMax})
    tcp.EntityData.Leafs.Append("tcpMaxConn", types.YLeaf{"TcpMaxConn", tcp.TcpMaxConn})
    tcp.EntityData.Leafs.Append("tcpActiveOpens", types.YLeaf{"TcpActiveOpens", tcp.TcpActiveOpens})
    tcp.EntityData.Leafs.Append("tcpPassiveOpens", types.YLeaf{"TcpPassiveOpens", tcp.TcpPassiveOpens})
    tcp.EntityData.Leafs.Append("tcpAttemptFails", types.YLeaf{"TcpAttemptFails", tcp.TcpAttemptFails})
    tcp.EntityData.Leafs.Append("tcpEstabResets", types.YLeaf{"TcpEstabResets", tcp.TcpEstabResets})
    tcp.EntityData.Leafs.Append("tcpCurrEstab", types.YLeaf{"TcpCurrEstab", tcp.TcpCurrEstab})
    tcp.EntityData.Leafs.Append("tcpInSegs", types.YLeaf{"TcpInSegs", tcp.TcpInSegs})
    tcp.EntityData.Leafs.Append("tcpOutSegs", types.YLeaf{"TcpOutSegs", tcp.TcpOutSegs})
    tcp.EntityData.Leafs.Append("tcpRetransSegs", types.YLeaf{"TcpRetransSegs", tcp.TcpRetransSegs})
    tcp.EntityData.Leafs.Append("tcpInErrs", types.YLeaf{"TcpInErrs", tcp.TcpInErrs})
    tcp.EntityData.Leafs.Append("tcpOutRsts", types.YLeaf{"TcpOutRsts", tcp.TcpOutRsts})
    tcp.EntityData.Leafs.Append("tcpHCInSegs", types.YLeaf{"TcpHCInSegs", tcp.TcpHCInSegs})
    tcp.EntityData.Leafs.Append("tcpHCOutSegs", types.YLeaf{"TcpHCOutSegs", tcp.TcpHCOutSegs})

    tcp.EntityData.YListKeys = []string {}

    return &(tcp.EntityData)
}

// TCPMIB_Tcp_TcpRtoAlgorithm represents retransmitting unacknowledged octets.
type TCPMIB_Tcp_TcpRtoAlgorithm string

const (
    TCPMIB_Tcp_TcpRtoAlgorithm_other TCPMIB_Tcp_TcpRtoAlgorithm = "other"

    TCPMIB_Tcp_TcpRtoAlgorithm_constant TCPMIB_Tcp_TcpRtoAlgorithm = "constant"

    TCPMIB_Tcp_TcpRtoAlgorithm_rsre TCPMIB_Tcp_TcpRtoAlgorithm = "rsre"

    TCPMIB_Tcp_TcpRtoAlgorithm_vanj TCPMIB_Tcp_TcpRtoAlgorithm = "vanj"

    TCPMIB_Tcp_TcpRtoAlgorithm_rfc2988 TCPMIB_Tcp_TcpRtoAlgorithm = "rfc2988"
)

// TCPMIB_TcpConnTable
// A table containing information about existing IPv4-specific
// TCP connections or listeners.  This table has been
// deprecated in favor of the version neutral
// tcpConnectionTable.
type TCPMIB_TcpConnTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row of the tcpConnTable containing information about a
    // particular current IPv4 TCP connection.  Each row of this table is
    // transient in that it ceases to exist when (or soon after) the connection
    // makes the transition to the CLOSED state. The type is slice of
    // TCPMIB_TcpConnTable_TcpConnEntry.
    TcpConnEntry []*TCPMIB_TcpConnTable_TcpConnEntry
}

func (tcpConnTable *TCPMIB_TcpConnTable) GetEntityData() *types.CommonEntityData {
    tcpConnTable.EntityData.YFilter = tcpConnTable.YFilter
    tcpConnTable.EntityData.YangName = "tcpConnTable"
    tcpConnTable.EntityData.BundleName = "cisco_ios_xe"
    tcpConnTable.EntityData.ParentYangName = "TCP-MIB"
    tcpConnTable.EntityData.SegmentPath = "tcpConnTable"
    tcpConnTable.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/" + tcpConnTable.EntityData.SegmentPath
    tcpConnTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpConnTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpConnTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpConnTable.EntityData.Children = types.NewOrderedMap()
    tcpConnTable.EntityData.Children.Append("tcpConnEntry", types.YChild{"TcpConnEntry", nil})
    for i := range tcpConnTable.TcpConnEntry {
        tcpConnTable.EntityData.Children.Append(types.GetSegmentPath(tcpConnTable.TcpConnEntry[i]), types.YChild{"TcpConnEntry", tcpConnTable.TcpConnEntry[i]})
    }
    tcpConnTable.EntityData.Leafs = types.NewOrderedMap()

    tcpConnTable.EntityData.YListKeys = []string {}

    return &(tcpConnTable.EntityData)
}

// TCPMIB_TcpConnTable_TcpConnEntry
// A conceptual row of the tcpConnTable containing information
// about a particular current IPv4 TCP connection.  Each row
// of this table is transient in that it ceases to exist when
// (or soon after) the connection makes the transition to the
// CLOSED state.
type TCPMIB_TcpConnTable_TcpConnEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The local IP address for this TCP connection.  In
    // the case of a connection in the listen state willing to accept connections
    // for any IP interface associated with the node, the value 0.0.0.0 is used.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TcpConnLocalAddress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    TcpConnLocalPort interface{}

    // This attribute is a key. The remote IP address for this TCP connection. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TcpConnRemAddress interface{}

    // This attribute is a key. The remote port number for this TCP connection.
    // The type is interface{} with range: 0..65535.
    TcpConnRemPort interface{}

    // The state of this TCP connection.  The only value that may be set by a
    // management station is deleteTCB(12).  Accordingly, it is appropriate for an
    // agent to return a `badValue' response if a management station attempts to
    // set this object to any other value.  If a management station sets this
    // object to the value deleteTCB(12), then the TCB (as defined in [RFC793]) of
    // the corresponding connection on the managed node is deleted, resulting in
    // immediate termination of the connection.  As an implementation-specific
    // option, a RST segment may be sent from the managed node to the other TCP
    // endpoint (note, however, that RST segments are not sent reliably). The type
    // is TcpConnState.
    TcpConnState interface{}
}

func (tcpConnEntry *TCPMIB_TcpConnTable_TcpConnEntry) GetEntityData() *types.CommonEntityData {
    tcpConnEntry.EntityData.YFilter = tcpConnEntry.YFilter
    tcpConnEntry.EntityData.YangName = "tcpConnEntry"
    tcpConnEntry.EntityData.BundleName = "cisco_ios_xe"
    tcpConnEntry.EntityData.ParentYangName = "tcpConnTable"
    tcpConnEntry.EntityData.SegmentPath = "tcpConnEntry" + types.AddKeyToken(tcpConnEntry.TcpConnLocalAddress, "tcpConnLocalAddress") + types.AddKeyToken(tcpConnEntry.TcpConnLocalPort, "tcpConnLocalPort") + types.AddKeyToken(tcpConnEntry.TcpConnRemAddress, "tcpConnRemAddress") + types.AddKeyToken(tcpConnEntry.TcpConnRemPort, "tcpConnRemPort")
    tcpConnEntry.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/tcpConnTable/" + tcpConnEntry.EntityData.SegmentPath
    tcpConnEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpConnEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpConnEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpConnEntry.EntityData.Children = types.NewOrderedMap()
    tcpConnEntry.EntityData.Leafs = types.NewOrderedMap()
    tcpConnEntry.EntityData.Leafs.Append("tcpConnLocalAddress", types.YLeaf{"TcpConnLocalAddress", tcpConnEntry.TcpConnLocalAddress})
    tcpConnEntry.EntityData.Leafs.Append("tcpConnLocalPort", types.YLeaf{"TcpConnLocalPort", tcpConnEntry.TcpConnLocalPort})
    tcpConnEntry.EntityData.Leafs.Append("tcpConnRemAddress", types.YLeaf{"TcpConnRemAddress", tcpConnEntry.TcpConnRemAddress})
    tcpConnEntry.EntityData.Leafs.Append("tcpConnRemPort", types.YLeaf{"TcpConnRemPort", tcpConnEntry.TcpConnRemPort})
    tcpConnEntry.EntityData.Leafs.Append("tcpConnState", types.YLeaf{"TcpConnState", tcpConnEntry.TcpConnState})

    tcpConnEntry.EntityData.YListKeys = []string {"TcpConnLocalAddress", "TcpConnLocalPort", "TcpConnRemAddress", "TcpConnRemPort"}

    return &(tcpConnEntry.EntityData)
}

// TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState represents however, that RST segments are not sent reliably).
type TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState string

const (
    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_closed TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "closed"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_listen TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "listen"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_synSent TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "synSent"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_synReceived TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "synReceived"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_established TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "established"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_finWait1 TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "finWait1"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_finWait2 TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "finWait2"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_closeWait TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "closeWait"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_lastAck TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "lastAck"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_closing TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "closing"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_timeWait TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "timeWait"

    TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState_deleteTCB TCPMIB_TcpConnTable_TcpConnEntry_TcpConnState = "deleteTCB"
)

// TCPMIB_TcpConnectionTable
// A table containing information about existing TCP
// connections.  Note that unlike earlier TCP MIBs, there
// is a separate table for connections in the LISTEN state.
type TCPMIB_TcpConnectionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row of the tcpConnectionTable containing information about a
    // particular current TCP connection. Each row of this table is transient in
    // that it ceases to exist when (or soon after) the connection makes the
    // transition to the CLOSED state. The type is slice of
    // TCPMIB_TcpConnectionTable_TcpConnectionEntry.
    TcpConnectionEntry []*TCPMIB_TcpConnectionTable_TcpConnectionEntry
}

func (tcpConnectionTable *TCPMIB_TcpConnectionTable) GetEntityData() *types.CommonEntityData {
    tcpConnectionTable.EntityData.YFilter = tcpConnectionTable.YFilter
    tcpConnectionTable.EntityData.YangName = "tcpConnectionTable"
    tcpConnectionTable.EntityData.BundleName = "cisco_ios_xe"
    tcpConnectionTable.EntityData.ParentYangName = "TCP-MIB"
    tcpConnectionTable.EntityData.SegmentPath = "tcpConnectionTable"
    tcpConnectionTable.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/" + tcpConnectionTable.EntityData.SegmentPath
    tcpConnectionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpConnectionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpConnectionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpConnectionTable.EntityData.Children = types.NewOrderedMap()
    tcpConnectionTable.EntityData.Children.Append("tcpConnectionEntry", types.YChild{"TcpConnectionEntry", nil})
    for i := range tcpConnectionTable.TcpConnectionEntry {
        tcpConnectionTable.EntityData.Children.Append(types.GetSegmentPath(tcpConnectionTable.TcpConnectionEntry[i]), types.YChild{"TcpConnectionEntry", tcpConnectionTable.TcpConnectionEntry[i]})
    }
    tcpConnectionTable.EntityData.Leafs = types.NewOrderedMap()

    tcpConnectionTable.EntityData.YListKeys = []string {}

    return &(tcpConnectionTable.EntityData)
}

// TCPMIB_TcpConnectionTable_TcpConnectionEntry
// A conceptual row of the tcpConnectionTable containing
// information about a particular current TCP connection.
// Each row of this table is transient in that it ceases to
// exist when (or soon after) the connection makes the
// transition to the CLOSED state.
type TCPMIB_TcpConnectionTable_TcpConnectionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address type of tcpConnectionLocalAddress. The
    // type is InetAddressType.
    TcpConnectionLocalAddressType interface{}

    // This attribute is a key. The local IP address for this TCP connection.  The
    // type of this address is determined by the value of
    // tcpConnectionLocalAddressType.  As this object is used in the index for the
    // tcpConnectionTable, implementors should be careful not to create entries
    // that would result in OIDs with more than 128 subidentifiers; otherwise the
    // information cannot be accessed by using SNMPv1, SNMPv2c, or SNMPv3. The
    // type is string with length: 0..255.
    TcpConnectionLocalAddress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    TcpConnectionLocalPort interface{}

    // This attribute is a key. The address type of tcpConnectionRemAddress. The
    // type is InetAddressType.
    TcpConnectionRemAddressType interface{}

    // This attribute is a key. The remote IP address for this TCP connection. 
    // The type of this address is determined by the value of
    // tcpConnectionRemAddressType.  As this object is used in the index for the
    // tcpConnectionTable, implementors should be careful not to create entries
    // that would result in OIDs with more than 128 subidentifiers; otherwise the
    // information cannot be accessed by using SNMPv1, SNMPv2c, or SNMPv3. The
    // type is string with length: 0..255.
    TcpConnectionRemAddress interface{}

    // This attribute is a key. The remote port number for this TCP connection.
    // The type is interface{} with range: 0..65535.
    TcpConnectionRemPort interface{}

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
    // sent reliably). The type is TcpConnectionState.
    TcpConnectionState interface{}

    // The system's process ID for the process associated with this connection, or
    // zero if there is no such process.  This value is expected to be the same as
    // HOST-RESOURCES-MIB:: hrSWRunIndex or SYSAPPL-MIB::sysApplElmtRunIndex for
    // some row in the appropriate tables. The type is interface{} with range:
    // 0..4294967295.
    TcpConnectionProcess interface{}
}

func (tcpConnectionEntry *TCPMIB_TcpConnectionTable_TcpConnectionEntry) GetEntityData() *types.CommonEntityData {
    tcpConnectionEntry.EntityData.YFilter = tcpConnectionEntry.YFilter
    tcpConnectionEntry.EntityData.YangName = "tcpConnectionEntry"
    tcpConnectionEntry.EntityData.BundleName = "cisco_ios_xe"
    tcpConnectionEntry.EntityData.ParentYangName = "tcpConnectionTable"
    tcpConnectionEntry.EntityData.SegmentPath = "tcpConnectionEntry" + types.AddKeyToken(tcpConnectionEntry.TcpConnectionLocalAddressType, "tcpConnectionLocalAddressType") + types.AddKeyToken(tcpConnectionEntry.TcpConnectionLocalAddress, "tcpConnectionLocalAddress") + types.AddKeyToken(tcpConnectionEntry.TcpConnectionLocalPort, "tcpConnectionLocalPort") + types.AddKeyToken(tcpConnectionEntry.TcpConnectionRemAddressType, "tcpConnectionRemAddressType") + types.AddKeyToken(tcpConnectionEntry.TcpConnectionRemAddress, "tcpConnectionRemAddress") + types.AddKeyToken(tcpConnectionEntry.TcpConnectionRemPort, "tcpConnectionRemPort")
    tcpConnectionEntry.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/tcpConnectionTable/" + tcpConnectionEntry.EntityData.SegmentPath
    tcpConnectionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpConnectionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpConnectionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpConnectionEntry.EntityData.Children = types.NewOrderedMap()
    tcpConnectionEntry.EntityData.Leafs = types.NewOrderedMap()
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionLocalAddressType", types.YLeaf{"TcpConnectionLocalAddressType", tcpConnectionEntry.TcpConnectionLocalAddressType})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionLocalAddress", types.YLeaf{"TcpConnectionLocalAddress", tcpConnectionEntry.TcpConnectionLocalAddress})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionLocalPort", types.YLeaf{"TcpConnectionLocalPort", tcpConnectionEntry.TcpConnectionLocalPort})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionRemAddressType", types.YLeaf{"TcpConnectionRemAddressType", tcpConnectionEntry.TcpConnectionRemAddressType})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionRemAddress", types.YLeaf{"TcpConnectionRemAddress", tcpConnectionEntry.TcpConnectionRemAddress})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionRemPort", types.YLeaf{"TcpConnectionRemPort", tcpConnectionEntry.TcpConnectionRemPort})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionState", types.YLeaf{"TcpConnectionState", tcpConnectionEntry.TcpConnectionState})
    tcpConnectionEntry.EntityData.Leafs.Append("tcpConnectionProcess", types.YLeaf{"TcpConnectionProcess", tcpConnectionEntry.TcpConnectionProcess})

    tcpConnectionEntry.EntityData.YListKeys = []string {"TcpConnectionLocalAddressType", "TcpConnectionLocalAddress", "TcpConnectionLocalPort", "TcpConnectionRemAddressType", "TcpConnectionRemAddress", "TcpConnectionRemPort"}

    return &(tcpConnectionEntry.EntityData)
}

// TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState represents however, that RST segments are not sent reliably).
type TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState string

const (
    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_closed TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "closed"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_listen TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "listen"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_synSent TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "synSent"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_synReceived TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "synReceived"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_established TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "established"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_finWait1 TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "finWait1"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_finWait2 TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "finWait2"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_closeWait TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "closeWait"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_lastAck TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "lastAck"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_closing TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "closing"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_timeWait TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "timeWait"

    TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState_deleteTCB TCPMIB_TcpConnectionTable_TcpConnectionEntry_TcpConnectionState = "deleteTCB"
)

// TCPMIB_TcpListenerTable
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
type TCPMIB_TcpListenerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row of the tcpListenerTable containing information about a
    // particular TCP listener. The type is slice of
    // TCPMIB_TcpListenerTable_TcpListenerEntry.
    TcpListenerEntry []*TCPMIB_TcpListenerTable_TcpListenerEntry
}

func (tcpListenerTable *TCPMIB_TcpListenerTable) GetEntityData() *types.CommonEntityData {
    tcpListenerTable.EntityData.YFilter = tcpListenerTable.YFilter
    tcpListenerTable.EntityData.YangName = "tcpListenerTable"
    tcpListenerTable.EntityData.BundleName = "cisco_ios_xe"
    tcpListenerTable.EntityData.ParentYangName = "TCP-MIB"
    tcpListenerTable.EntityData.SegmentPath = "tcpListenerTable"
    tcpListenerTable.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/" + tcpListenerTable.EntityData.SegmentPath
    tcpListenerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpListenerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpListenerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpListenerTable.EntityData.Children = types.NewOrderedMap()
    tcpListenerTable.EntityData.Children.Append("tcpListenerEntry", types.YChild{"TcpListenerEntry", nil})
    for i := range tcpListenerTable.TcpListenerEntry {
        tcpListenerTable.EntityData.Children.Append(types.GetSegmentPath(tcpListenerTable.TcpListenerEntry[i]), types.YChild{"TcpListenerEntry", tcpListenerTable.TcpListenerEntry[i]})
    }
    tcpListenerTable.EntityData.Leafs = types.NewOrderedMap()

    tcpListenerTable.EntityData.YListKeys = []string {}

    return &(tcpListenerTable.EntityData)
}

// TCPMIB_TcpListenerTable_TcpListenerEntry
// A conceptual row of the tcpListenerTable containing
// information about a particular TCP listener.
type TCPMIB_TcpListenerTable_TcpListenerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address type of tcpListenerLocalAddress.  The
    // value should be unknown (0) if connection initiations to all local IP
    // addresses are accepted. The type is InetAddressType.
    TcpListenerLocalAddressType interface{}

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
    TcpListenerLocalAddress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    TcpListenerLocalPort interface{}

    // The system's process ID for the process associated with this listener, or
    // zero if there is no such process.  This value is expected to be the same as
    // HOST-RESOURCES-MIB:: hrSWRunIndex or SYSAPPL-MIB::sysApplElmtRunIndex for
    // some row in the appropriate tables. The type is interface{} with range:
    // 0..4294967295.
    TcpListenerProcess interface{}
}

func (tcpListenerEntry *TCPMIB_TcpListenerTable_TcpListenerEntry) GetEntityData() *types.CommonEntityData {
    tcpListenerEntry.EntityData.YFilter = tcpListenerEntry.YFilter
    tcpListenerEntry.EntityData.YangName = "tcpListenerEntry"
    tcpListenerEntry.EntityData.BundleName = "cisco_ios_xe"
    tcpListenerEntry.EntityData.ParentYangName = "tcpListenerTable"
    tcpListenerEntry.EntityData.SegmentPath = "tcpListenerEntry" + types.AddKeyToken(tcpListenerEntry.TcpListenerLocalAddressType, "tcpListenerLocalAddressType") + types.AddKeyToken(tcpListenerEntry.TcpListenerLocalAddress, "tcpListenerLocalAddress") + types.AddKeyToken(tcpListenerEntry.TcpListenerLocalPort, "tcpListenerLocalPort")
    tcpListenerEntry.EntityData.AbsolutePath = "TCP-MIB:TCP-MIB/tcpListenerTable/" + tcpListenerEntry.EntityData.SegmentPath
    tcpListenerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpListenerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpListenerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpListenerEntry.EntityData.Children = types.NewOrderedMap()
    tcpListenerEntry.EntityData.Leafs = types.NewOrderedMap()
    tcpListenerEntry.EntityData.Leafs.Append("tcpListenerLocalAddressType", types.YLeaf{"TcpListenerLocalAddressType", tcpListenerEntry.TcpListenerLocalAddressType})
    tcpListenerEntry.EntityData.Leafs.Append("tcpListenerLocalAddress", types.YLeaf{"TcpListenerLocalAddress", tcpListenerEntry.TcpListenerLocalAddress})
    tcpListenerEntry.EntityData.Leafs.Append("tcpListenerLocalPort", types.YLeaf{"TcpListenerLocalPort", tcpListenerEntry.TcpListenerLocalPort})
    tcpListenerEntry.EntityData.Leafs.Append("tcpListenerProcess", types.YLeaf{"TcpListenerProcess", tcpListenerEntry.TcpListenerProcess})

    tcpListenerEntry.EntityData.YListKeys = []string {"TcpListenerLocalAddressType", "TcpListenerLocalAddress", "TcpListenerLocalPort"}

    return &(tcpListenerEntry.EntityData)
}

