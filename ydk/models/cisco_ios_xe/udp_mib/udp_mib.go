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
    Udptable UDPMIB_Udptable

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
    Udpendpointtable UDPMIB_Udpendpointtable
}

func (uDPMIB *UDPMIB) GetEntityData() *types.CommonEntityData {
    uDPMIB.EntityData.YFilter = uDPMIB.YFilter
    uDPMIB.EntityData.YangName = "UDP-MIB"
    uDPMIB.EntityData.BundleName = "cisco_ios_xe"
    uDPMIB.EntityData.ParentYangName = "UDP-MIB"
    uDPMIB.EntityData.SegmentPath = "UDP-MIB:UDP-MIB"
    uDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    uDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    uDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    uDPMIB.EntityData.Children = make(map[string]types.YChild)
    uDPMIB.EntityData.Children["udp"] = types.YChild{"Udp", &uDPMIB.Udp}
    uDPMIB.EntityData.Children["udpTable"] = types.YChild{"Udptable", &uDPMIB.Udptable}
    uDPMIB.EntityData.Children["udpEndpointTable"] = types.YChild{"Udpendpointtable", &uDPMIB.Udpendpointtable}
    uDPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Udpindatagrams interface{}

    // The total number of received UDP datagrams for which there was no
    // application at the destination port.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by discontinuities in the value of sysUpTime. The
    // type is interface{} with range: 0..4294967295.
    Udpnoports interface{}

    // The number of received UDP datagrams that could not be delivered for
    // reasons other than the lack of an application at the destination port. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by
    // discontinuities in the value of sysUpTime. The type is interface{} with
    // range: 0..4294967295.
    Udpinerrors interface{}

    // The total number of UDP datagrams sent from this entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range: 0..4294967295.
    Udpoutdatagrams interface{}

    // The total number of UDP datagrams delivered to UDP users, for devices that
    // can receive more than 1 million UDP datagrams per second.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Udphcindatagrams interface{}

    // The total number of UDP datagrams sent from this entity, for devices that
    // can transmit more than 1 million UDP datagrams per second.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by discontinuities in
    // the value of sysUpTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Udphcoutdatagrams interface{}
}

func (udp *UDPMIB_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xe"
    udp.EntityData.ParentYangName = "UDP-MIB"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    udp.EntityData.Leafs["udpInDatagrams"] = types.YLeaf{"Udpindatagrams", udp.Udpindatagrams}
    udp.EntityData.Leafs["udpNoPorts"] = types.YLeaf{"Udpnoports", udp.Udpnoports}
    udp.EntityData.Leafs["udpInErrors"] = types.YLeaf{"Udpinerrors", udp.Udpinerrors}
    udp.EntityData.Leafs["udpOutDatagrams"] = types.YLeaf{"Udpoutdatagrams", udp.Udpoutdatagrams}
    udp.EntityData.Leafs["udpHCInDatagrams"] = types.YLeaf{"Udphcindatagrams", udp.Udphcindatagrams}
    udp.EntityData.Leafs["udpHCOutDatagrams"] = types.YLeaf{"Udphcoutdatagrams", udp.Udphcoutdatagrams}
    return &(udp.EntityData)
}

// UDPMIB_Udptable
// A table containing IPv4-specific UDP listener
// information.  It contains information about all local
// IPv4 UDP end-points on which an application is
// currently accepting datagrams.  This table has been
// deprecated in favor of the version neutral
// udpEndpointTable.
type UDPMIB_Udptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current UDP listener. The type is slice of
    // UDPMIB_Udptable_Udpentry.
    Udpentry []UDPMIB_Udptable_Udpentry
}

func (udptable *UDPMIB_Udptable) GetEntityData() *types.CommonEntityData {
    udptable.EntityData.YFilter = udptable.YFilter
    udptable.EntityData.YangName = "udpTable"
    udptable.EntityData.BundleName = "cisco_ios_xe"
    udptable.EntityData.ParentYangName = "UDP-MIB"
    udptable.EntityData.SegmentPath = "udpTable"
    udptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udptable.EntityData.Children = make(map[string]types.YChild)
    udptable.EntityData.Children["udpEntry"] = types.YChild{"Udpentry", nil}
    for i := range udptable.Udpentry {
        udptable.EntityData.Children[types.GetSegmentPath(&udptable.Udpentry[i])] = types.YChild{"Udpentry", &udptable.Udpentry[i]}
    }
    udptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(udptable.EntityData)
}

// UDPMIB_Udptable_Udpentry
// Information about a particular current UDP listener.
type UDPMIB_Udptable_Udpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this UDP listener.  In
    // the case of a UDP listener that is willing to accept datagrams for any IP
    // interface associated with the node, the value 0.0.0.0 is used. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Udplocaladdress interface{}

    // This attribute is a key. The local port number for this UDP listener. The
    // type is interface{} with range: 0..65535.
    Udplocalport interface{}
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetEntityData() *types.CommonEntityData {
    udpentry.EntityData.YFilter = udpentry.YFilter
    udpentry.EntityData.YangName = "udpEntry"
    udpentry.EntityData.BundleName = "cisco_ios_xe"
    udpentry.EntityData.ParentYangName = "udpTable"
    udpentry.EntityData.SegmentPath = "udpEntry" + "[udpLocalAddress='" + fmt.Sprintf("%v", udpentry.Udplocaladdress) + "']" + "[udpLocalPort='" + fmt.Sprintf("%v", udpentry.Udplocalport) + "']"
    udpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpentry.EntityData.Children = make(map[string]types.YChild)
    udpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    udpentry.EntityData.Leafs["udpLocalAddress"] = types.YLeaf{"Udplocaladdress", udpentry.Udplocaladdress}
    udpentry.EntityData.Leafs["udpLocalPort"] = types.YLeaf{"Udplocalport", udpentry.Udplocalport}
    return &(udpentry.EntityData)
}

// UDPMIB_Udpendpointtable
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
type UDPMIB_Udpendpointtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current UDP endpoint.  Implementers need to
    // be aware that if the total number of elements (octets or sub-identifiers)
    // in udpEndpointLocalAddress and udpEndpointRemoteAddress exceeds 111, then
    // OIDs of column instances in this table will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3.
    // The type is slice of UDPMIB_Udpendpointtable_Udpendpointentry.
    Udpendpointentry []UDPMIB_Udpendpointtable_Udpendpointentry
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetEntityData() *types.CommonEntityData {
    udpendpointtable.EntityData.YFilter = udpendpointtable.YFilter
    udpendpointtable.EntityData.YangName = "udpEndpointTable"
    udpendpointtable.EntityData.BundleName = "cisco_ios_xe"
    udpendpointtable.EntityData.ParentYangName = "UDP-MIB"
    udpendpointtable.EntityData.SegmentPath = "udpEndpointTable"
    udpendpointtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpendpointtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpendpointtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpendpointtable.EntityData.Children = make(map[string]types.YChild)
    udpendpointtable.EntityData.Children["udpEndpointEntry"] = types.YChild{"Udpendpointentry", nil}
    for i := range udpendpointtable.Udpendpointentry {
        udpendpointtable.EntityData.Children[types.GetSegmentPath(&udpendpointtable.Udpendpointentry[i])] = types.YChild{"Udpendpointentry", &udpendpointtable.Udpendpointentry[i]}
    }
    udpendpointtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(udpendpointtable.EntityData)
}

// UDPMIB_Udpendpointtable_Udpendpointentry
// Information about a particular current UDP endpoint.
// 
// Implementers need to be aware that if the total number
// of elements (octets or sub-identifiers) in
// udpEndpointLocalAddress and udpEndpointRemoteAddress
// exceeds 111, then OIDs of column instances in this table
// will have more than 128 sub-identifiers and cannot be
// accessed using SNMPv1, SNMPv2c, or SNMPv3.
type UDPMIB_Udpendpointtable_Udpendpointentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The address type of udpEndpointLocalAddress.  Only
    // IPv4, IPv4z, IPv6, and IPv6z addresses are expected, or unknown(0) if
    // datagrams for all local IP addresses are accepted. The type is
    // InetAddressType.
    Udpendpointlocaladdresstype interface{}

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
    Udpendpointlocaladdress interface{}

    // This attribute is a key. The local port number for this UDP endpoint. The
    // type is interface{} with range: 0..65535.
    Udpendpointlocalport interface{}

    // This attribute is a key. The address type of udpEndpointRemoteAddress. 
    // Only IPv4, IPv4z, IPv6, and IPv6z addresses are expected, or unknown(0) if
    // datagrams for all remote IP addresses are accepted.  Also, note that some
    // combinations of  udpEndpointLocalAdressType and
    // udpEndpointRemoteAddressType are not supported.  In particular, if the
    // value of this object is not unknown(0), it is expected to always refer to
    // the same IP version as udpEndpointLocalAddressType. The type is
    // InetAddressType.
    Udpendpointremoteaddresstype interface{}

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
    Udpendpointremoteaddress interface{}

    // This attribute is a key. The remote port number for this UDP endpoint.  If
    // datagrams from any remote system are to be accepted, this value is zero.
    // The type is interface{} with range: 0..65535.
    Udpendpointremoteport interface{}

    // This attribute is a key. The instance of this tuple.  This object is used
    // to distinguish among multiple processes 'connected' to the same UDP
    // endpoint.  For example, on a system implementing the BSD sockets interface,
    // this would be used to support the SO_REUSEADDR and SO_REUSEPORT socket
    // options. The type is interface{} with range: 1..4294967295.
    Udpendpointinstance interface{}

    // The system's process ID for the process associated with this endpoint, or
    // zero if there is no such process. This value is expected to be the same as
    // HOST-RESOURCES-MIB::hrSWRunIndex or SYSAPPL-MIB:: sysApplElmtRunIndex for
    // some row in the appropriate tables. The type is interface{} with range:
    // 0..4294967295.
    Udpendpointprocess interface{}
}

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetEntityData() *types.CommonEntityData {
    udpendpointentry.EntityData.YFilter = udpendpointentry.YFilter
    udpendpointentry.EntityData.YangName = "udpEndpointEntry"
    udpendpointentry.EntityData.BundleName = "cisco_ios_xe"
    udpendpointentry.EntityData.ParentYangName = "udpEndpointTable"
    udpendpointentry.EntityData.SegmentPath = "udpEndpointEntry" + "[udpEndpointLocalAddressType='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointlocaladdresstype) + "']" + "[udpEndpointLocalAddress='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointlocaladdress) + "']" + "[udpEndpointLocalPort='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointlocalport) + "']" + "[udpEndpointRemoteAddressType='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointremoteaddresstype) + "']" + "[udpEndpointRemoteAddress='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointremoteaddress) + "']" + "[udpEndpointRemotePort='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointremoteport) + "']" + "[udpEndpointInstance='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointinstance) + "']"
    udpendpointentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udpendpointentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udpendpointentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udpendpointentry.EntityData.Children = make(map[string]types.YChild)
    udpendpointentry.EntityData.Leafs = make(map[string]types.YLeaf)
    udpendpointentry.EntityData.Leafs["udpEndpointLocalAddressType"] = types.YLeaf{"Udpendpointlocaladdresstype", udpendpointentry.Udpendpointlocaladdresstype}
    udpendpointentry.EntityData.Leafs["udpEndpointLocalAddress"] = types.YLeaf{"Udpendpointlocaladdress", udpendpointentry.Udpendpointlocaladdress}
    udpendpointentry.EntityData.Leafs["udpEndpointLocalPort"] = types.YLeaf{"Udpendpointlocalport", udpendpointentry.Udpendpointlocalport}
    udpendpointentry.EntityData.Leafs["udpEndpointRemoteAddressType"] = types.YLeaf{"Udpendpointremoteaddresstype", udpendpointentry.Udpendpointremoteaddresstype}
    udpendpointentry.EntityData.Leafs["udpEndpointRemoteAddress"] = types.YLeaf{"Udpendpointremoteaddress", udpendpointentry.Udpendpointremoteaddress}
    udpendpointentry.EntityData.Leafs["udpEndpointRemotePort"] = types.YLeaf{"Udpendpointremoteport", udpendpointentry.Udpendpointremoteport}
    udpendpointentry.EntityData.Leafs["udpEndpointInstance"] = types.YLeaf{"Udpendpointinstance", udpendpointentry.Udpendpointinstance}
    udpendpointentry.EntityData.Leafs["udpEndpointProcess"] = types.YLeaf{"Udpendpointprocess", udpendpointentry.Udpendpointprocess}
    return &(udpendpointentry.EntityData)
}

