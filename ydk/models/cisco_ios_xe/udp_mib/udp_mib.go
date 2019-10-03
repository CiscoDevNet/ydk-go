// The MIB module for managing UDP implementations.
// Copyright (C) The Internet Society (2005).  This
// version of this MIB module is part of RFC 4113;
// see the RFC itself for full legal notices.
package udp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package udp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:UDP-MIB UDP-MIB}", reflect.TypeOf(UDPMIB{}))
    ydk.RegisterEntity("UDP-MIB:UDP-MIB", reflect.TypeOf(UDPMIB{}))
}

// UDPMIB
type UDPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Udp UDPMIB_Udp

    // A table containing IPv4-specific UDP listener information.  It contains
    // information about all local IPv4 UDP end-points on which an application is
    // currently accepting datagrams.  This table has been deprecated in favor of
    // the version neutral udpEndpointTable.
    UdpTable UDPMIB_UdpTable

    // A table containing information about this entity's UDP endpoints on which a
    // local application is currently accepting or sending datagrams.  The address
    // type in this table represents the address type used for the communication,
    // irrespective of the higher-layer abstraction.  For example, an application
    // using IPv6 'sockets' to communicate via IPv4 between ::ffff:10.0.0.1 and
    // ::ffff:10.0.0.2 would use InetAddressType ipv4(1).  Unlike the udpTable in
    // RFC 2013, this table also allows the representation of an application that
    // completely specifies both local and remote addresses and ports.  A
    // listening application is represented in three possible ways:  1) An
    // application that is willing to accept both IPv4    and IPv6 datagrams is
    // represented by a    udpEndpointLocalAddressType of unknown(0) and a   
    // udpEndpointLocalAddress of ''h (a zero-length    octet-string).  2) An
    // application that is willing to accept only IPv4    or only IPv6 datagrams
    // is represented by a    udpEndpointLocalAddressType of the appropriate   
    // address type and a udpEndpointLocalAddress of    '0.0.0.0' or '::'
    // respectively.  3) An application that is listening for datagrams only   
    // for a specific IP address but from any remote    system is represented by a
    // udpEndpointLocalAddressType of the appropriate    address type, with
    // udpEndpointLocalAddress    specifying the local address.  In all cases
    // where the remote is a wildcard, the udpEndpointRemoteAddressType is
    // unknown(0), the udpEndpointRemoteAddress is ''h (a zero-length
    // octet-string), and the udpEndpointRemotePort is 0.  If the operating system
    // is demultiplexing UDP packets by remote address and port, or if the
    // application has 'connected' the socket specifying a default remote address
    // and port, the udpEndpointRemote* values should be used to reflect this.
    UdpEndpointTable UDPMIB_UdpEndpointTable
}

func (uDPMIB *UDPMIB) GetEntityData() *types.CommonEntityData {
    uDPMIB.EntityData.YFilter = uDPMIB.YFilter
    uDPMIB.EntityData.YangName = "UDP-MIB"
    uDPMIB.EntityData.BundleName = "cisco_ios_xe"
    uDPMIB.EntityData.ParentYangName = "UDP-MIB"
    uDPMIB.EntityData.SegmentPath = "UDP-MIB:UDP-MIB"
    uDPMIB.EntityData.AbsolutePath = uDPMIB.EntityData.SegmentPath
    uDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    uDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    uDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    uDPMIB.EntityData.Children = types.NewOrderedMap()
    uDPMIB.EntityData.Children.Append("udp", types.YChild{"Udp", &uDPMIB.Udp})
    uDPMIB.EntityData.Children.Append("udpTable", types.YChild{"UdpTable", &uDPMIB.UdpTable})
    uDPMIB.EntityData.Children.Append("udpEndpointTable", types.YChild{"UdpEndpointTable", &uDPMIB.UdpEndpointTable})
    uDPMIB.EntityData.Leafs = types.NewOrderedMap()

    uDPMIB.EntityData.YListKeys = []string {}

    return &(uDPMIB.EntityData)
}

// UDPMIB_Udp
type UDPMIB_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of UDP datagrams delivered to UDP users.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range: 0..4294967295.
    UdpInDatagrams interface{}

    // The total number of received UDP datagrams for which there was no
    // application at the destination port.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by discontinuities in the value of sysUpTime. The
    // type is interface{} with range: 0..4294967295.
    UdpNoPorts interface{}

    // The number of received UDP datagrams that could not be delivered for
    // reasons other than the lack of an application at the destination port. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    UdpInErrors interface{}

    // The total number of UDP datagrams sent from this entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range: 0..4294967295.
    UdpOutDatagrams interface{}

    // The total number of UDP datagrams delivered to UDP users, for devices that
    // can receive more than 1 million UDP datagrams per second.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range:
    // 0..18446744073709551615.
    UdpHCInDatagrams interface{}

    // The total number of UDP datagrams sent from this entity, for devices that
    // can transmit more than 1 million UDP datagrams per second.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range:
    // 0..18446744073709551615.
    UdpHCOutDatagrams interface{}
}

func (udp *UDPMIB_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xe"
    udp.EntityData.ParentYangName = "UDP-MIB"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.AbsolutePath = "UDP-MIB:UDP-MIB/" + udp.EntityData.SegmentPath
    udp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udp.EntityData.Children = types.NewOrderedMap()
    udp.EntityData.Leafs = types.NewOrderedMap()
    udp.EntityData.Leafs.Append("udpInDatagrams", types.YLeaf{"UdpInDatagrams", udp.UdpInDatagrams})
    udp.EntityData.Leafs.Append("udpNoPorts", types.YLeaf{"UdpNoPorts", udp.UdpNoPorts})
    udp.EntityData.Leafs.Append("udpInErrors", types.YLeaf{"UdpInErrors", udp.UdpInErrors})
    udp.EntityData.Leafs.Append("udpOutDatagrams", types.YLeaf{"UdpOutDatagrams", udp.UdpOutDatagrams})
    udp.EntityData.Leafs.Append("udpHCInDatagrams", types.YLeaf{"UdpHCInDatagrams", udp.UdpHCInDatagrams})
    udp.EntityData.Leafs.Append("udpHCOutDatagrams", types.YLeaf{"UdpHCOutDatagrams", udp.UdpHCOutDatagrams})

    udp.EntityData.YListKeys = []string {}

    return &(udp.EntityData)
}

// UDPMIB_UdpTable
// A table containing IPv4-specific UDP listener
// information.  It contains information about all local
// IPv4 UDP end-points on which an application is
// currently accepting datagrams.  This table has been
// deprecated in favor of the version neutral
// udpEndpointTable.
type UDPMIB_UdpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current UDP listener. The type is slice of
    // UDPMIB_UdpTable_UdpEntry.
    UdpEntry []*UDPMIB_UdpTable_UdpEntry
}

func (udpTable *UDPMIB_UdpTable) GetEntityData() *types.CommonEntityData {
    udpTable.EntityData.YFilter = udpTable.YFilter
    udpTable.EntityData.YangName = "udpTable"
    udpTable.EntityData.BundleName = "cisco_ios_xe"
    udpTable.EntityData.ParentYangName = "UDP-MIB"
    udpTable.EntityData.SegmentPath = "udpTable"
    udpTable.EntityData.AbsolutePath = "UDP-MIB:UDP-MIB/" + udpTable.EntityData.SegmentPath
    udpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpTable.EntityData.Children = types.NewOrderedMap()
    udpTable.EntityData.Children.Append("udpEntry", types.YChild{"UdpEntry", nil})
    for i := range udpTable.UdpEntry {
        udpTable.EntityData.Children.Append(types.GetSegmentPath(udpTable.UdpEntry[i]), types.YChild{"UdpEntry", udpTable.UdpEntry[i]})
    }
    udpTable.EntityData.Leafs = types.NewOrderedMap()

    udpTable.EntityData.YListKeys = []string {}

    return &(udpTable.EntityData)
}

// UDPMIB_UdpTable_UdpEntry
// Information about a particular current UDP listener.
type UDPMIB_UdpTable_UdpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The local IP address for this UDP listener.  In
    // the case of a UDP listener that is willing to accept datagrams for any IP
    // interface associated with the node, the value 0.0.0.0 is used. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    UdpLocalAddress interface{}

    // This attribute is a key. The local port number for this UDP listener. The
    // type is interface{} with range: 0..65535.
    UdpLocalPort interface{}
}

func (udpEntry *UDPMIB_UdpTable_UdpEntry) GetEntityData() *types.CommonEntityData {
    udpEntry.EntityData.YFilter = udpEntry.YFilter
    udpEntry.EntityData.YangName = "udpEntry"
    udpEntry.EntityData.BundleName = "cisco_ios_xe"
    udpEntry.EntityData.ParentYangName = "udpTable"
    udpEntry.EntityData.SegmentPath = "udpEntry" + types.AddKeyToken(udpEntry.UdpLocalAddress, "udpLocalAddress") + types.AddKeyToken(udpEntry.UdpLocalPort, "udpLocalPort")
    udpEntry.EntityData.AbsolutePath = "UDP-MIB:UDP-MIB/udpTable/" + udpEntry.EntityData.SegmentPath
    udpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpEntry.EntityData.Children = types.NewOrderedMap()
    udpEntry.EntityData.Leafs = types.NewOrderedMap()
    udpEntry.EntityData.Leafs.Append("udpLocalAddress", types.YLeaf{"UdpLocalAddress", udpEntry.UdpLocalAddress})
    udpEntry.EntityData.Leafs.Append("udpLocalPort", types.YLeaf{"UdpLocalPort", udpEntry.UdpLocalPort})

    udpEntry.EntityData.YListKeys = []string {"UdpLocalAddress", "UdpLocalPort"}

    return &(udpEntry.EntityData)
}

// UDPMIB_UdpEndpointTable
// A table containing information about this entity's UDP
// endpoints on which a local application is currently
// accepting or sending datagrams.
// 
// The address type in this table represents the address
// type used for the communication, irrespective of the
// higher-layer abstraction.  For example, an application
// using IPv6 'sockets' to communicate via IPv4 between
// ::ffff:10.0.0.1 and ::ffff:10.0.0.2 would use
// InetAddressType ipv4(1).
// 
// Unlike the udpTable in RFC 2013, this table also allows
// the representation of an application that completely
// specifies both local and remote addresses and ports.  A
// listening application is represented in three possible
// ways:
// 
// 1) An application that is willing to accept both IPv4
//    and IPv6 datagrams is represented by a
//    udpEndpointLocalAddressType of unknown(0) and a
//    udpEndpointLocalAddress of ''h (a zero-length
//    octet-string).
// 
// 2) An application that is willing to accept only IPv4
//    or only IPv6 datagrams is represented by a
//    udpEndpointLocalAddressType of the appropriate
//    address type and a udpEndpointLocalAddress of
//    '0.0.0.0' or '::' respectively.
// 
// 3) An application that is listening for datagrams only
//    for a specific IP address but from any remote
//    system is represented by a
//    udpEndpointLocalAddressType of the appropriate
//    address type, with udpEndpointLocalAddress
//    specifying the local address.
// 
// In all cases where the remote is a wildcard, the
// udpEndpointRemoteAddressType is unknown(0), the
// udpEndpointRemoteAddress is ''h (a zero-length
// octet-string), and the udpEndpointRemotePort is 0.
// 
// If the operating system is demultiplexing UDP packets
// by remote address and port, or if the application has
// 'connected' the socket specifying a default remote
// address and port, the udpEndpointRemote* values should
// be used to reflect this.
type UDPMIB_UdpEndpointTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current UDP endpoint.  Implementers need to
    // be aware that if the total number of elements (octets or sub-identifiers)
    // in udpEndpointLocalAddress and udpEndpointRemoteAddress exceeds 111, then
    // OIDs of column instances in this table will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3.
    // The type is slice of UDPMIB_UdpEndpointTable_UdpEndpointEntry.
    UdpEndpointEntry []*UDPMIB_UdpEndpointTable_UdpEndpointEntry
}

func (udpEndpointTable *UDPMIB_UdpEndpointTable) GetEntityData() *types.CommonEntityData {
    udpEndpointTable.EntityData.YFilter = udpEndpointTable.YFilter
    udpEndpointTable.EntityData.YangName = "udpEndpointTable"
    udpEndpointTable.EntityData.BundleName = "cisco_ios_xe"
    udpEndpointTable.EntityData.ParentYangName = "UDP-MIB"
    udpEndpointTable.EntityData.SegmentPath = "udpEndpointTable"
    udpEndpointTable.EntityData.AbsolutePath = "UDP-MIB:UDP-MIB/" + udpEndpointTable.EntityData.SegmentPath
    udpEndpointTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpEndpointTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpEndpointTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpEndpointTable.EntityData.Children = types.NewOrderedMap()
    udpEndpointTable.EntityData.Children.Append("udpEndpointEntry", types.YChild{"UdpEndpointEntry", nil})
    for i := range udpEndpointTable.UdpEndpointEntry {
        udpEndpointTable.EntityData.Children.Append(types.GetSegmentPath(udpEndpointTable.UdpEndpointEntry[i]), types.YChild{"UdpEndpointEntry", udpEndpointTable.UdpEndpointEntry[i]})
    }
    udpEndpointTable.EntityData.Leafs = types.NewOrderedMap()

    udpEndpointTable.EntityData.YListKeys = []string {}

    return &(udpEndpointTable.EntityData)
}

// UDPMIB_UdpEndpointTable_UdpEndpointEntry
// Information about a particular current UDP endpoint.
// 
// Implementers need to be aware that if the total number
// of elements (octets or sub-identifiers) in
// udpEndpointLocalAddress and udpEndpointRemoteAddress
// exceeds 111, then OIDs of column instances in this table
// will have more than 128 sub-identifiers and cannot be
// accessed using SNMPv1, SNMPv2c, or SNMPv3.
type UDPMIB_UdpEndpointTable_UdpEndpointEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address type of udpEndpointLocalAddress.  Only
    // IPv4, IPv4z, IPv6, and IPv6z addresses are expected, or unknown(0) if
    // datagrams for all local IP addresses are accepted. The type is
    // InetAddressType.
    UdpEndpointLocalAddressType interface{}

    // This attribute is a key. The local IP address for this UDP endpoint.  The
    // value of this object can be represented in three  possible ways, depending
    // on the characteristics of the listening application:  1. For an application
    // that is willing to accept both    IPv4 and IPv6 datagrams, the value of
    // this object    must be ''h (a zero-length octet-string), with    the value
    // of the corresponding instance of the    udpEndpointLocalAddressType object
    // being unknown(0).  2. For an application that is willing to accept only
    // IPv4    or only IPv6 datagrams, the value of this object    must be
    // '0.0.0.0' or '::', respectively, while the    corresponding instance of the
    // udpEndpointLocalAddressType object represents the    appropriate address
    // type.  3. For an application that is listening for data    destined only to
    // a specific IP address, the value    of this object is the specific IP
    // address for which    this node is receiving packets, with the   
    // corresponding instance of the    udpEndpointLocalAddressType object
    // representing the    appropriate address type.  As this object is used in
    // the index for the udpEndpointTable, implementors of this table should be
    // careful not to create entries that would result in OIDs with more than 128
    // subidentifiers; else the information cannot be accessed using SNMPv1,
    // SNMPv2c, or SNMPv3. The type is string with length: 0..255.
    UdpEndpointLocalAddress interface{}

    // This attribute is a key. The local port number for this UDP endpoint. The
    // type is interface{} with range: 0..65535.
    UdpEndpointLocalPort interface{}

    // This attribute is a key. The address type of udpEndpointRemoteAddress. 
    // Only IPv4, IPv4z, IPv6, and IPv6z addresses are expected, or unknown(0) if
    // datagrams for all remote IP addresses are accepted.  Also, note that some
    // combinations of  udpEndpointLocalAdressType and
    // udpEndpointRemoteAddressType are not supported.  In particular, if the
    // value of this object is not unknown(0), it is expected to always refer to
    // the same IP version as udpEndpointLocalAddressType. The type is
    // InetAddressType.
    UdpEndpointRemoteAddressType interface{}

    // This attribute is a key. The remote IP address for this UDP endpoint.  If
    // datagrams from any remote system are to be accepted, this value is ''h (a
    // zero-length octet-string). Otherwise, it has the type described by
    // udpEndpointRemoteAddressType and is the address of the remote system from
    // which datagrams are to be accepted (or to which all datagrams will be
    // sent).  As this object is used in the index for the udpEndpointTable,
    // implementors of this table should be careful not to create entries that
    // would result in OIDs with more than 128 subidentifiers; else the
    // information cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3. The type
    // is string with length: 0..255.
    UdpEndpointRemoteAddress interface{}

    // This attribute is a key. The remote port number for this UDP endpoint.  If
    // datagrams from any remote system are to be accepted, this value is zero.
    // The type is interface{} with range: 0..65535.
    UdpEndpointRemotePort interface{}

    // This attribute is a key. The instance of this tuple.  This object is used
    // to distinguish among multiple processes 'connected' to the same UDP
    // endpoint.  For example, on a system implementing the BSD sockets interface,
    // this would be used to support the SO_REUSEADDR and SO_REUSEPORT socket
    // options. The type is interface{} with range: 1..4294967295.
    UdpEndpointInstance interface{}

    // The system's process ID for the process associated with this endpoint, or
    // zero if there is no such process. This value is expected to be the same as
    // HOST-RESOURCES-MIB::hrSWRunIndex or SYSAPPL-MIB:: sysApplElmtRunIndex for
    // some row in the appropriate tables. The type is interface{} with range:
    // 0..4294967295.
    UdpEndpointProcess interface{}
}

func (udpEndpointEntry *UDPMIB_UdpEndpointTable_UdpEndpointEntry) GetEntityData() *types.CommonEntityData {
    udpEndpointEntry.EntityData.YFilter = udpEndpointEntry.YFilter
    udpEndpointEntry.EntityData.YangName = "udpEndpointEntry"
    udpEndpointEntry.EntityData.BundleName = "cisco_ios_xe"
    udpEndpointEntry.EntityData.ParentYangName = "udpEndpointTable"
    udpEndpointEntry.EntityData.SegmentPath = "udpEndpointEntry" + types.AddKeyToken(udpEndpointEntry.UdpEndpointLocalAddressType, "udpEndpointLocalAddressType") + types.AddKeyToken(udpEndpointEntry.UdpEndpointLocalAddress, "udpEndpointLocalAddress") + types.AddKeyToken(udpEndpointEntry.UdpEndpointLocalPort, "udpEndpointLocalPort") + types.AddKeyToken(udpEndpointEntry.UdpEndpointRemoteAddressType, "udpEndpointRemoteAddressType") + types.AddKeyToken(udpEndpointEntry.UdpEndpointRemoteAddress, "udpEndpointRemoteAddress") + types.AddKeyToken(udpEndpointEntry.UdpEndpointRemotePort, "udpEndpointRemotePort") + types.AddKeyToken(udpEndpointEntry.UdpEndpointInstance, "udpEndpointInstance")
    udpEndpointEntry.EntityData.AbsolutePath = "UDP-MIB:UDP-MIB/udpEndpointTable/" + udpEndpointEntry.EntityData.SegmentPath
    udpEndpointEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpEndpointEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpEndpointEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpEndpointEntry.EntityData.Children = types.NewOrderedMap()
    udpEndpointEntry.EntityData.Leafs = types.NewOrderedMap()
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointLocalAddressType", types.YLeaf{"UdpEndpointLocalAddressType", udpEndpointEntry.UdpEndpointLocalAddressType})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointLocalAddress", types.YLeaf{"UdpEndpointLocalAddress", udpEndpointEntry.UdpEndpointLocalAddress})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointLocalPort", types.YLeaf{"UdpEndpointLocalPort", udpEndpointEntry.UdpEndpointLocalPort})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointRemoteAddressType", types.YLeaf{"UdpEndpointRemoteAddressType", udpEndpointEntry.UdpEndpointRemoteAddressType})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointRemoteAddress", types.YLeaf{"UdpEndpointRemoteAddress", udpEndpointEntry.UdpEndpointRemoteAddress})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointRemotePort", types.YLeaf{"UdpEndpointRemotePort", udpEndpointEntry.UdpEndpointRemotePort})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointInstance", types.YLeaf{"UdpEndpointInstance", udpEndpointEntry.UdpEndpointInstance})
    udpEndpointEntry.EntityData.Leafs.Append("udpEndpointProcess", types.YLeaf{"UdpEndpointProcess", udpEndpointEntry.UdpEndpointProcess})

    udpEndpointEntry.EntityData.YListKeys = []string {"UdpEndpointLocalAddressType", "UdpEndpointLocalAddress", "UdpEndpointLocalPort", "UdpEndpointRemoteAddressType", "UdpEndpointRemoteAddress", "UdpEndpointRemotePort", "UdpEndpointInstance"}

    return &(udpEndpointEntry.EntityData)
}

