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
    parent types.Entity
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

func (uDPMIB *UDPMIB) GetFilter() yfilter.YFilter { return uDPMIB.YFilter }

func (uDPMIB *UDPMIB) SetFilter(yf yfilter.YFilter) { uDPMIB.YFilter = yf }

func (uDPMIB *UDPMIB) GetGoName(yname string) string {
    if yname == "udp" { return "Udp" }
    if yname == "udpTable" { return "Udptable" }
    if yname == "udpEndpointTable" { return "Udpendpointtable" }
    return ""
}

func (uDPMIB *UDPMIB) GetSegmentPath() string {
    return "UDP-MIB:UDP-MIB"
}

func (uDPMIB *UDPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udp" {
        return &uDPMIB.Udp
    }
    if childYangName == "udpTable" {
        return &uDPMIB.Udptable
    }
    if childYangName == "udpEndpointTable" {
        return &uDPMIB.Udpendpointtable
    }
    return nil
}

func (uDPMIB *UDPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udp"] = &uDPMIB.Udp
    children["udpTable"] = &uDPMIB.Udptable
    children["udpEndpointTable"] = &uDPMIB.Udpendpointtable
    return children
}

func (uDPMIB *UDPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (uDPMIB *UDPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (uDPMIB *UDPMIB) GetYangName() string { return "UDP-MIB" }

func (uDPMIB *UDPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (uDPMIB *UDPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (uDPMIB *UDPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (uDPMIB *UDPMIB) SetParent(parent types.Entity) { uDPMIB.parent = parent }

func (uDPMIB *UDPMIB) GetParent() types.Entity { return uDPMIB.parent }

func (uDPMIB *UDPMIB) GetParentYangName() string { return "UDP-MIB" }

// UDPMIB_Udp
type UDPMIB_Udp struct {
    parent types.Entity
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

func (udp *UDPMIB_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *UDPMIB_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *UDPMIB_Udp) GetGoName(yname string) string {
    if yname == "udpInDatagrams" { return "Udpindatagrams" }
    if yname == "udpNoPorts" { return "Udpnoports" }
    if yname == "udpInErrors" { return "Udpinerrors" }
    if yname == "udpOutDatagrams" { return "Udpoutdatagrams" }
    if yname == "udpHCInDatagrams" { return "Udphcindatagrams" }
    if yname == "udpHCOutDatagrams" { return "Udphcoutdatagrams" }
    return ""
}

func (udp *UDPMIB_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *UDPMIB_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udp *UDPMIB_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udp *UDPMIB_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udpInDatagrams"] = udp.Udpindatagrams
    leafs["udpNoPorts"] = udp.Udpnoports
    leafs["udpInErrors"] = udp.Udpinerrors
    leafs["udpOutDatagrams"] = udp.Udpoutdatagrams
    leafs["udpHCInDatagrams"] = udp.Udphcindatagrams
    leafs["udpHCOutDatagrams"] = udp.Udphcoutdatagrams
    return leafs
}

func (udp *UDPMIB_Udp) GetBundleName() string { return "cisco_ios_xe" }

func (udp *UDPMIB_Udp) GetYangName() string { return "udp" }

func (udp *UDPMIB_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udp *UDPMIB_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udp *UDPMIB_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udp *UDPMIB_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *UDPMIB_Udp) GetParent() types.Entity { return udp.parent }

func (udp *UDPMIB_Udp) GetParentYangName() string { return "UDP-MIB" }

// UDPMIB_Udptable
// A table containing IPv4-specific UDP listener
// information.  It contains information about all local
// IPv4 UDP end-points on which an application is
// currently accepting datagrams.  This table has been
// deprecated in favor of the version neutral
// udpEndpointTable.
type UDPMIB_Udptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular current UDP listener. The type is slice of
    // UDPMIB_Udptable_Udpentry.
    Udpentry []UDPMIB_Udptable_Udpentry
}

func (udptable *UDPMIB_Udptable) GetFilter() yfilter.YFilter { return udptable.YFilter }

func (udptable *UDPMIB_Udptable) SetFilter(yf yfilter.YFilter) { udptable.YFilter = yf }

func (udptable *UDPMIB_Udptable) GetGoName(yname string) string {
    if yname == "udpEntry" { return "Udpentry" }
    return ""
}

func (udptable *UDPMIB_Udptable) GetSegmentPath() string {
    return "udpTable"
}

func (udptable *UDPMIB_Udptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udpEntry" {
        for _, c := range udptable.Udpentry {
            if udptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UDPMIB_Udptable_Udpentry{}
        udptable.Udpentry = append(udptable.Udpentry, child)
        return &udptable.Udpentry[len(udptable.Udpentry)-1]
    }
    return nil
}

func (udptable *UDPMIB_Udptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range udptable.Udpentry {
        children[udptable.Udpentry[i].GetSegmentPath()] = &udptable.Udpentry[i]
    }
    return children
}

func (udptable *UDPMIB_Udptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udptable *UDPMIB_Udptable) GetBundleName() string { return "cisco_ios_xe" }

func (udptable *UDPMIB_Udptable) GetYangName() string { return "udpTable" }

func (udptable *UDPMIB_Udptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udptable *UDPMIB_Udptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udptable *UDPMIB_Udptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udptable *UDPMIB_Udptable) SetParent(parent types.Entity) { udptable.parent = parent }

func (udptable *UDPMIB_Udptable) GetParent() types.Entity { return udptable.parent }

func (udptable *UDPMIB_Udptable) GetParentYangName() string { return "UDP-MIB" }

// UDPMIB_Udptable_Udpentry
// Information about a particular current UDP listener.
type UDPMIB_Udptable_Udpentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this UDP listener.  In
    // the case of a UDP listener that is willing to accept datagrams for any IP
    // interface associated with the node, the value 0.0.0.0 is used. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Udplocaladdress interface{}

    // This attribute is a key. The local port number for this UDP listener. The
    // type is interface{} with range: 0..65535.
    Udplocalport interface{}
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetFilter() yfilter.YFilter { return udpentry.YFilter }

func (udpentry *UDPMIB_Udptable_Udpentry) SetFilter(yf yfilter.YFilter) { udpentry.YFilter = yf }

func (udpentry *UDPMIB_Udptable_Udpentry) GetGoName(yname string) string {
    if yname == "udpLocalAddress" { return "Udplocaladdress" }
    if yname == "udpLocalPort" { return "Udplocalport" }
    return ""
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetSegmentPath() string {
    return "udpEntry" + "[udpLocalAddress='" + fmt.Sprintf("%v", udpentry.Udplocaladdress) + "']" + "[udpLocalPort='" + fmt.Sprintf("%v", udpentry.Udplocalport) + "']"
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udpLocalAddress"] = udpentry.Udplocaladdress
    leafs["udpLocalPort"] = udpentry.Udplocalport
    return leafs
}

func (udpentry *UDPMIB_Udptable_Udpentry) GetBundleName() string { return "cisco_ios_xe" }

func (udpentry *UDPMIB_Udptable_Udpentry) GetYangName() string { return "udpEntry" }

func (udpentry *UDPMIB_Udptable_Udpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udpentry *UDPMIB_Udptable_Udpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udpentry *UDPMIB_Udptable_Udpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udpentry *UDPMIB_Udptable_Udpentry) SetParent(parent types.Entity) { udpentry.parent = parent }

func (udpentry *UDPMIB_Udptable_Udpentry) GetParent() types.Entity { return udpentry.parent }

func (udpentry *UDPMIB_Udptable_Udpentry) GetParentYangName() string { return "udpTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular current UDP endpoint.  Implementers need to
    // be aware that if the total number of elements (octets or sub-identifiers)
    // in udpEndpointLocalAddress and udpEndpointRemoteAddress exceeds 111, then
    // OIDs of column instances in this table will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3.
    // The type is slice of UDPMIB_Udpendpointtable_Udpendpointentry.
    Udpendpointentry []UDPMIB_Udpendpointtable_Udpendpointentry
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetFilter() yfilter.YFilter { return udpendpointtable.YFilter }

func (udpendpointtable *UDPMIB_Udpendpointtable) SetFilter(yf yfilter.YFilter) { udpendpointtable.YFilter = yf }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetGoName(yname string) string {
    if yname == "udpEndpointEntry" { return "Udpendpointentry" }
    return ""
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetSegmentPath() string {
    return "udpEndpointTable"
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udpEndpointEntry" {
        for _, c := range udpendpointtable.Udpendpointentry {
            if udpendpointtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UDPMIB_Udpendpointtable_Udpendpointentry{}
        udpendpointtable.Udpendpointentry = append(udpendpointtable.Udpendpointentry, child)
        return &udpendpointtable.Udpendpointentry[len(udpendpointtable.Udpendpointentry)-1]
    }
    return nil
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range udpendpointtable.Udpendpointentry {
        children[udpendpointtable.Udpendpointentry[i].GetSegmentPath()] = &udpendpointtable.Udpendpointentry[i]
    }
    return children
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udpendpointtable *UDPMIB_Udpendpointtable) GetBundleName() string { return "cisco_ios_xe" }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetYangName() string { return "udpEndpointTable" }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udpendpointtable *UDPMIB_Udpendpointtable) SetParent(parent types.Entity) { udpendpointtable.parent = parent }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetParent() types.Entity { return udpendpointtable.parent }

func (udpendpointtable *UDPMIB_Udpendpointtable) GetParentYangName() string { return "UDP-MIB" }

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
    parent types.Entity
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

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetFilter() yfilter.YFilter { return udpendpointentry.YFilter }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) SetFilter(yf yfilter.YFilter) { udpendpointentry.YFilter = yf }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetGoName(yname string) string {
    if yname == "udpEndpointLocalAddressType" { return "Udpendpointlocaladdresstype" }
    if yname == "udpEndpointLocalAddress" { return "Udpendpointlocaladdress" }
    if yname == "udpEndpointLocalPort" { return "Udpendpointlocalport" }
    if yname == "udpEndpointRemoteAddressType" { return "Udpendpointremoteaddresstype" }
    if yname == "udpEndpointRemoteAddress" { return "Udpendpointremoteaddress" }
    if yname == "udpEndpointRemotePort" { return "Udpendpointremoteport" }
    if yname == "udpEndpointInstance" { return "Udpendpointinstance" }
    if yname == "udpEndpointProcess" { return "Udpendpointprocess" }
    return ""
}

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetSegmentPath() string {
    return "udpEndpointEntry" + "[udpEndpointLocalAddressType='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointlocaladdresstype) + "']" + "[udpEndpointLocalAddress='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointlocaladdress) + "']" + "[udpEndpointLocalPort='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointlocalport) + "']" + "[udpEndpointRemoteAddressType='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointremoteaddresstype) + "']" + "[udpEndpointRemoteAddress='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointremoteaddress) + "']" + "[udpEndpointRemotePort='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointremoteport) + "']" + "[udpEndpointInstance='" + fmt.Sprintf("%v", udpendpointentry.Udpendpointinstance) + "']"
}

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udpEndpointLocalAddressType"] = udpendpointentry.Udpendpointlocaladdresstype
    leafs["udpEndpointLocalAddress"] = udpendpointentry.Udpendpointlocaladdress
    leafs["udpEndpointLocalPort"] = udpendpointentry.Udpendpointlocalport
    leafs["udpEndpointRemoteAddressType"] = udpendpointentry.Udpendpointremoteaddresstype
    leafs["udpEndpointRemoteAddress"] = udpendpointentry.Udpendpointremoteaddress
    leafs["udpEndpointRemotePort"] = udpendpointentry.Udpendpointremoteport
    leafs["udpEndpointInstance"] = udpendpointentry.Udpendpointinstance
    leafs["udpEndpointProcess"] = udpendpointentry.Udpendpointprocess
    return leafs
}

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetBundleName() string { return "cisco_ios_xe" }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetYangName() string { return "udpEndpointEntry" }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) SetParent(parent types.Entity) { udpendpointentry.parent = parent }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetParent() types.Entity { return udpendpointentry.parent }

func (udpendpointentry *UDPMIB_Udpendpointtable_Udpendpointentry) GetParentYangName() string { return "udpEndpointTable" }

