package rfc1213_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rfc1213_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:RFC1213-MIB RFC1213-MIB}", reflect.TypeOf(RFC1213MIB{}))
    ydk.RegisterEntity("RFC1213-MIB:RFC1213-MIB", reflect.TypeOf(RFC1213MIB{}))
}

// RFC1213MIB
type RFC1213MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    System RFC1213MIB_System

    
    Interfaces RFC1213MIB_Interfaces

    
    Ip RFC1213MIB_Ip

    
    Icmp RFC1213MIB_Icmp

    
    Tcp RFC1213MIB_Tcp

    
    Udp RFC1213MIB_Udp

    
    Egp RFC1213MIB_Egp

    
    Snmp RFC1213MIB_Snmp

    // A list of interface entries.  The number of entries is given by the value
    // of ifNumber.
    IfTable RFC1213MIB_IfTable

    // The Address Translation tables contain the NetworkAddress to `physical'
    // address equivalences. Some interfaces do not use translation tables for
    // determining address equivalences (e.g., DDN-X.25 has an algorithmic
    // method); if all interfaces are of this type, then the Address Translation
    // table is empty, i.e., has zero entries.
    AtTable RFC1213MIB_AtTable

    // The table of addressing information relevant to this entity's IP addresses.
    IpAddrTable RFC1213MIB_IpAddrTable

    // This entity's IP Routing table.
    IpRouteTable RFC1213MIB_IpRouteTable

    // The IP Address Translation table used for mapping from IP addresses to
    // physical addresses.
    IpNetToMediaTable RFC1213MIB_IpNetToMediaTable

    // A table containing TCP connection-specific information.
    TcpConnTable RFC1213MIB_TcpConnTable

    // A table containing UDP listener information.
    UdpTable RFC1213MIB_UdpTable

    // The EGP neighbor table.
    EgpNeighTable RFC1213MIB_EgpNeighTable
}

func (rFC1213MIB *RFC1213MIB) GetEntityData() *types.CommonEntityData {
    rFC1213MIB.EntityData.YFilter = rFC1213MIB.YFilter
    rFC1213MIB.EntityData.YangName = "RFC1213-MIB"
    rFC1213MIB.EntityData.BundleName = "cisco_ios_xe"
    rFC1213MIB.EntityData.ParentYangName = "RFC1213-MIB"
    rFC1213MIB.EntityData.SegmentPath = "RFC1213-MIB:RFC1213-MIB"
    rFC1213MIB.EntityData.AbsolutePath = rFC1213MIB.EntityData.SegmentPath
    rFC1213MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rFC1213MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rFC1213MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rFC1213MIB.EntityData.Children = types.NewOrderedMap()
    rFC1213MIB.EntityData.Children.Append("system", types.YChild{"System", &rFC1213MIB.System})
    rFC1213MIB.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &rFC1213MIB.Interfaces})
    rFC1213MIB.EntityData.Children.Append("ip", types.YChild{"Ip", &rFC1213MIB.Ip})
    rFC1213MIB.EntityData.Children.Append("icmp", types.YChild{"Icmp", &rFC1213MIB.Icmp})
    rFC1213MIB.EntityData.Children.Append("tcp", types.YChild{"Tcp", &rFC1213MIB.Tcp})
    rFC1213MIB.EntityData.Children.Append("udp", types.YChild{"Udp", &rFC1213MIB.Udp})
    rFC1213MIB.EntityData.Children.Append("egp", types.YChild{"Egp", &rFC1213MIB.Egp})
    rFC1213MIB.EntityData.Children.Append("snmp", types.YChild{"Snmp", &rFC1213MIB.Snmp})
    rFC1213MIB.EntityData.Children.Append("ifTable", types.YChild{"IfTable", &rFC1213MIB.IfTable})
    rFC1213MIB.EntityData.Children.Append("atTable", types.YChild{"AtTable", &rFC1213MIB.AtTable})
    rFC1213MIB.EntityData.Children.Append("ipAddrTable", types.YChild{"IpAddrTable", &rFC1213MIB.IpAddrTable})
    rFC1213MIB.EntityData.Children.Append("ipRouteTable", types.YChild{"IpRouteTable", &rFC1213MIB.IpRouteTable})
    rFC1213MIB.EntityData.Children.Append("ipNetToMediaTable", types.YChild{"IpNetToMediaTable", &rFC1213MIB.IpNetToMediaTable})
    rFC1213MIB.EntityData.Children.Append("tcpConnTable", types.YChild{"TcpConnTable", &rFC1213MIB.TcpConnTable})
    rFC1213MIB.EntityData.Children.Append("udpTable", types.YChild{"UdpTable", &rFC1213MIB.UdpTable})
    rFC1213MIB.EntityData.Children.Append("egpNeighTable", types.YChild{"EgpNeighTable", &rFC1213MIB.EgpNeighTable})
    rFC1213MIB.EntityData.Leafs = types.NewOrderedMap()

    rFC1213MIB.EntityData.YListKeys = []string {}

    return &(rFC1213MIB.EntityData)
}

// RFC1213MIB_System
type RFC1213MIB_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A textual description of the entity.  This value should include the full
    // name and version identification of the system's hardware type, software
    // operating-system, and networking software.  It is mandatory that this only
    // contain printable ASCII characters. The type is string with length: 0..255.
    SysDescr interface{}

    // The vendor's authoritative identification of the network management
    // subsystem contained in the entity.  This value is allocated within the SMI
    // enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous
    // means for determining `what kind of box' is being managed.  For example, if
    // vendor `Flintstones, Inc.' was assigned the subtree 1.3.6.1.4.1.4242, it
    // could assign the identifier 1.3.6.1.4.1.4242.1.1 to its `Fred Router'. The
    // type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    SysObjectID interface{}

    // The time (in hundredths of a second) since the network management portion
    // of the system was last re-initialized. The type is interface{} with range:
    // 0..4294967295.
    SysUpTime interface{}

    // The textual identification of the contact person for this managed node,
    // together with information on how to contact this person. The type is string
    // with length: 0..255.
    SysContact interface{}

    // An administratively-assigned name for this managed node.  By convention,
    // this is the node's fully-qualified domain name. The type is string with
    // length: 0..255.
    SysName interface{}

    // The physical location of this node (e.g., `telephone closet, 3rd floor').
    // The type is string with length: 0..255.
    SysLocation interface{}

    // A value which indicates the set of services that this entity primarily
    // offers.  The value is a sum.  This sum initially takes the value zero,
    // Then, for each layer, L, in the range 1 through 7, that this node performs
    // transactions for, 2 raised to (L - 1) is added to the sum.  For example, a
    // node which performs primarily routing functions would have a value of 4
    // (2^(3-1)).  In contrast, a node which is a host offering application
    // services would have a value of 72 (2^(4-1) + 2^(7-1)).  Note that in the
    // context of the Internet suite of protocols, values should be calculated
    // accordingly:       layer  functionality          1  physical (e.g.,
    // repeaters)          2  datalink/subnetwork (e.g., bridges)          3 
    // internet (e.g., IP gateways)          4  end-to-end  (e.g., IP hosts)      
    // 7  applications (e.g., mail relays)  For systems including OSI protocols,
    // layers 5 and 6 may also be counted. The type is interface{} with range:
    // 0..127.
    SysServices interface{}
}

func (system *RFC1213MIB_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xe"
    system.EntityData.ParentYangName = "RFC1213-MIB"
    system.EntityData.SegmentPath = "system"
    system.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + system.EntityData.SegmentPath
    system.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Leafs = types.NewOrderedMap()
    system.EntityData.Leafs.Append("sysDescr", types.YLeaf{"SysDescr", system.SysDescr})
    system.EntityData.Leafs.Append("sysObjectID", types.YLeaf{"SysObjectID", system.SysObjectID})
    system.EntityData.Leafs.Append("sysUpTime", types.YLeaf{"SysUpTime", system.SysUpTime})
    system.EntityData.Leafs.Append("sysContact", types.YLeaf{"SysContact", system.SysContact})
    system.EntityData.Leafs.Append("sysName", types.YLeaf{"SysName", system.SysName})
    system.EntityData.Leafs.Append("sysLocation", types.YLeaf{"SysLocation", system.SysLocation})
    system.EntityData.Leafs.Append("sysServices", types.YLeaf{"SysServices", system.SysServices})

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// RFC1213MIB_Interfaces
type RFC1213MIB_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of network interfaces (regardless of their current state)
    // present on this system. The type is interface{} with range:
    // -2147483648..2147483647.
    IfNumber interface{}
}

func (interfaces *RFC1213MIB_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "RFC1213-MIB"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Leafs = types.NewOrderedMap()
    interfaces.EntityData.Leafs.Append("ifNumber", types.YLeaf{"IfNumber", interfaces.IfNumber})

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// RFC1213MIB_Ip
type RFC1213MIB_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The indication of whether this entity is acting as an IP gateway in respect
    // to the forwarding of datagrams received by, but not addressed to, this
    // entity.  IP gateways forward datagrams.  IP hosts do not (except those
    // source-routed via the host).  Note that for some managed nodes, this object
    // may take on only a subset of the values possible. Accordingly, it is
    // appropriate for an agent to return a `badValue' response if a management
    // station attempts to change this object to an inappropriate value. The type
    // is IpForwarding.
    IpForwarding interface{}

    // The default value inserted into the Time-To-Live field of the IP header of
    // datagrams originated at this entity, whenever a TTL value is not supplied
    // by the transport layer protocol. The type is interface{} with range:
    // -2147483648..2147483647.
    IpDefaultTTL interface{}

    // The total number of input datagrams received from interfaces, including
    // those received in error. The type is interface{} with range: 0..4294967295.
    IpInReceives interface{}

    // The number of input datagrams discarded due to errors in their IP headers,
    // including bad checksums, version number mismatch, other format errors,
    // time-to-live exceeded, errors discovered in processing their IP options,
    // etc. The type is interface{} with range: 0..4294967295.
    IpInHdrErrors interface{}

    // The number of input datagrams discarded because the IP address in their IP
    // header's destination field was not a valid address to be received at this
    // entity.  This count includes invalid addresses (e.g., 0.0.0.0) and
    // addresses of unsupported Classes (e.g., Class E).  For entities which are
    // not IP Gateways and therefore do not forward datagrams, this counter
    // includes datagrams discarded because the destination address was not a
    // local address. The type is interface{} with range: 0..4294967295.
    IpInAddrErrors interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination, as a result of which an attempt was made to find a route to
    // forward them to that final destination. In entities which do not act as IP
    // Gateways, this counter will include only those packets which were
    // Source-Routed via this entity, and the Source- Route option processing was
    // successful. The type is interface{} with range: 0..4294967295.
    IpForwDatagrams interface{}

    // The number of locally-addressed datagrams received successfully but
    // discarded because of an unknown or unsupported protocol. The type is
    // interface{} with range: 0..4294967295.
    IpInUnknownProtos interface{}

    // The number of input IP datagrams for which no problems were encountered to
    // prevent their continued processing, but which were discarded (e.g., for
    // lack of buffer space).  Note that this counter does not include any
    // datagrams discarded while awaiting re-assembly. The type is interface{}
    // with range: 0..4294967295.
    IpInDiscards interface{}

    // The total number of input datagrams successfully delivered to IP
    // user-protocols (including ICMP). The type is interface{} with range:
    // 0..4294967295.
    IpInDelivers interface{}

    // The total number of IP datagrams which local IP user-protocols (including
    // ICMP) supplied to IP in requests for transmission.  Note that this counter
    // does not include any datagrams counted in ipForwDatagrams. The type is
    // interface{} with range: 0..4294967295.
    IpOutRequests interface{}

    // The number of output IP datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but which were discarded
    // (e.g., for lack of buffer space).  Note that this counter would include
    // datagrams counted in ipForwDatagrams if any such packets met this
    // (discretionary) discard criterion. The type is interface{} with range:
    // 0..4294967295.
    IpOutDiscards interface{}

    // The number of IP datagrams discarded because no route could be found to
    // transmit them to their destination.  Note that this counter includes any
    // packets counted in ipForwDatagrams which meet this `no-route' criterion. 
    // Note that this includes any datagrams which a host cannot route because all
    // of its default gateways are down. The type is interface{} with range:
    // 0..4294967295.
    IpOutNoRoutes interface{}

    // The maximum number of seconds which received fragments are held while they
    // are awaiting reassembly at this entity. The type is interface{} with range:
    // -2147483648..2147483647.
    IpReasmTimeout interface{}

    // The number of IP fragments received which needed to be reassembled at this
    // entity. The type is interface{} with range: 0..4294967295.
    IpReasmReqds interface{}

    // The number of IP datagrams successfully re- assembled. The type is
    // interface{} with range: 0..4294967295.
    IpReasmOKs interface{}

    // The number of failures detected by the IP re- assembly algorithm (for
    // whatever reason: timed out, errors, etc).  Note that this is not
    // necessarily a count of discarded IP fragments since some algorithms
    // (notably the algorithm in RFC 815) can lose track of the number of
    // fragments by combining them as they are received. The type is interface{}
    // with range: 0..4294967295.
    IpReasmFails interface{}

    // The number of IP datagrams that have been successfully fragmented at this
    // entity. The type is interface{} with range: 0..4294967295.
    IpFragOKs interface{}

    // The number of IP datagrams that have been discarded because they needed to
    // be fragmented at this entity but could not be, e.g., because their Don't
    // Fragment flag was set. The type is interface{} with range: 0..4294967295.
    IpFragFails interface{}

    // The number of IP datagram fragments that have been generated as a result of
    // fragmentation at this entity. The type is interface{} with range:
    // 0..4294967295.
    IpFragCreates interface{}

    // The number of routing entries which were chosen to be discarded even though
    // they are valid.  One possible reason for discarding such an entry could be
    // to free-up buffer space for other routing entries. The type is interface{}
    // with range: 0..4294967295.
    IpRoutingDiscards interface{}
}

func (ip *RFC1213MIB_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xe"
    ip.EntityData.ParentYangName = "RFC1213-MIB"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + ip.EntityData.SegmentPath
    ip.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("ipForwarding", types.YLeaf{"IpForwarding", ip.IpForwarding})
    ip.EntityData.Leafs.Append("ipDefaultTTL", types.YLeaf{"IpDefaultTTL", ip.IpDefaultTTL})
    ip.EntityData.Leafs.Append("ipInReceives", types.YLeaf{"IpInReceives", ip.IpInReceives})
    ip.EntityData.Leafs.Append("ipInHdrErrors", types.YLeaf{"IpInHdrErrors", ip.IpInHdrErrors})
    ip.EntityData.Leafs.Append("ipInAddrErrors", types.YLeaf{"IpInAddrErrors", ip.IpInAddrErrors})
    ip.EntityData.Leafs.Append("ipForwDatagrams", types.YLeaf{"IpForwDatagrams", ip.IpForwDatagrams})
    ip.EntityData.Leafs.Append("ipInUnknownProtos", types.YLeaf{"IpInUnknownProtos", ip.IpInUnknownProtos})
    ip.EntityData.Leafs.Append("ipInDiscards", types.YLeaf{"IpInDiscards", ip.IpInDiscards})
    ip.EntityData.Leafs.Append("ipInDelivers", types.YLeaf{"IpInDelivers", ip.IpInDelivers})
    ip.EntityData.Leafs.Append("ipOutRequests", types.YLeaf{"IpOutRequests", ip.IpOutRequests})
    ip.EntityData.Leafs.Append("ipOutDiscards", types.YLeaf{"IpOutDiscards", ip.IpOutDiscards})
    ip.EntityData.Leafs.Append("ipOutNoRoutes", types.YLeaf{"IpOutNoRoutes", ip.IpOutNoRoutes})
    ip.EntityData.Leafs.Append("ipReasmTimeout", types.YLeaf{"IpReasmTimeout", ip.IpReasmTimeout})
    ip.EntityData.Leafs.Append("ipReasmReqds", types.YLeaf{"IpReasmReqds", ip.IpReasmReqds})
    ip.EntityData.Leafs.Append("ipReasmOKs", types.YLeaf{"IpReasmOKs", ip.IpReasmOKs})
    ip.EntityData.Leafs.Append("ipReasmFails", types.YLeaf{"IpReasmFails", ip.IpReasmFails})
    ip.EntityData.Leafs.Append("ipFragOKs", types.YLeaf{"IpFragOKs", ip.IpFragOKs})
    ip.EntityData.Leafs.Append("ipFragFails", types.YLeaf{"IpFragFails", ip.IpFragFails})
    ip.EntityData.Leafs.Append("ipFragCreates", types.YLeaf{"IpFragCreates", ip.IpFragCreates})
    ip.EntityData.Leafs.Append("ipRoutingDiscards", types.YLeaf{"IpRoutingDiscards", ip.IpRoutingDiscards})

    ip.EntityData.YListKeys = []string {}

    return &(ip.EntityData)
}

// RFC1213MIB_Ip_IpForwarding represents inappropriate value.
type RFC1213MIB_Ip_IpForwarding string

const (
    RFC1213MIB_Ip_IpForwarding_forwarding RFC1213MIB_Ip_IpForwarding = "forwarding"

    RFC1213MIB_Ip_IpForwarding_not_forwarding RFC1213MIB_Ip_IpForwarding = "not-forwarding"
)

// RFC1213MIB_Icmp
type RFC1213MIB_Icmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of ICMP messages which the entity received.  Note that
    // this counter includes all those counted by icmpInErrors. The type is
    // interface{} with range: 0..4294967295.
    IcmpInMsgs interface{}

    // The number of ICMP messages which the entity received but determined as
    // having ICMP-specific errors (bad ICMP checksums, bad length, etc.). The
    // type is interface{} with range: 0..4294967295.
    IcmpInErrors interface{}

    // The number of ICMP Destination Unreachable messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInDestUnreachs interface{}

    // The number of ICMP Time Exceeded messages received. The type is interface{}
    // with range: 0..4294967295.
    IcmpInTimeExcds interface{}

    // The number of ICMP Parameter Problem messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInParmProbs interface{}

    // The number of ICMP Source Quench messages received. The type is interface{}
    // with range: 0..4294967295.
    IcmpInSrcQuenchs interface{}

    // The number of ICMP Redirect messages received. The type is interface{} with
    // range: 0..4294967295.
    IcmpInRedirects interface{}

    // The number of ICMP Echo (request) messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInEchos interface{}

    // The number of ICMP Echo Reply messages received. The type is interface{}
    // with range: 0..4294967295.
    IcmpInEchoReps interface{}

    // The number of ICMP Timestamp (request) messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInTimestamps interface{}

    // The number of ICMP Timestamp Reply messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInTimestampReps interface{}

    // The number of ICMP Address Mask Request messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInAddrMasks interface{}

    // The number of ICMP Address Mask Reply messages received. The type is
    // interface{} with range: 0..4294967295.
    IcmpInAddrMaskReps interface{}

    // The total number of ICMP messages which this entity attempted to send. 
    // Note that this counter includes all those counted by icmpOutErrors. The
    // type is interface{} with range: 0..4294967295.
    IcmpOutMsgs interface{}

    // The number of ICMP messages which this entity did not send due to problems
    // discovered within ICMP such as a lack of buffers.  This value should not
    // include errors discovered outside the ICMP layer such as the inability of
    // IP to route the resultant datagram.  In some implementations there may be
    // no types of error which contribute to this counter's value. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutErrors interface{}

    // The number of ICMP Destination Unreachable messages sent. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutDestUnreachs interface{}

    // The number of ICMP Time Exceeded messages sent. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutTimeExcds interface{}

    // The number of ICMP Parameter Problem messages sent. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutParmProbs interface{}

    // The number of ICMP Source Quench messages sent. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutSrcQuenchs interface{}

    // The number of ICMP Redirect messages sent.  For a host, this object will
    // always be zero, since hosts do not send redirects. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutRedirects interface{}

    // The number of ICMP Echo (request) messages sent. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutEchos interface{}

    // The number of ICMP Echo Reply messages sent. The type is interface{} with
    // range: 0..4294967295.
    IcmpOutEchoReps interface{}

    // The number of ICMP Timestamp (request) messages sent. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutTimestamps interface{}

    // The number of ICMP Timestamp Reply messages sent. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutTimestampReps interface{}

    // The number of ICMP Address Mask Request messages sent. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutAddrMasks interface{}

    // The number of ICMP Address Mask Reply messages sent. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutAddrMaskReps interface{}
}

func (icmp *RFC1213MIB_Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xe"
    icmp.EntityData.ParentYangName = "RFC1213-MIB"
    icmp.EntityData.SegmentPath = "icmp"
    icmp.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + icmp.EntityData.SegmentPath
    icmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmp.EntityData.Children = types.NewOrderedMap()
    icmp.EntityData.Leafs = types.NewOrderedMap()
    icmp.EntityData.Leafs.Append("icmpInMsgs", types.YLeaf{"IcmpInMsgs", icmp.IcmpInMsgs})
    icmp.EntityData.Leafs.Append("icmpInErrors", types.YLeaf{"IcmpInErrors", icmp.IcmpInErrors})
    icmp.EntityData.Leafs.Append("icmpInDestUnreachs", types.YLeaf{"IcmpInDestUnreachs", icmp.IcmpInDestUnreachs})
    icmp.EntityData.Leafs.Append("icmpInTimeExcds", types.YLeaf{"IcmpInTimeExcds", icmp.IcmpInTimeExcds})
    icmp.EntityData.Leafs.Append("icmpInParmProbs", types.YLeaf{"IcmpInParmProbs", icmp.IcmpInParmProbs})
    icmp.EntityData.Leafs.Append("icmpInSrcQuenchs", types.YLeaf{"IcmpInSrcQuenchs", icmp.IcmpInSrcQuenchs})
    icmp.EntityData.Leafs.Append("icmpInRedirects", types.YLeaf{"IcmpInRedirects", icmp.IcmpInRedirects})
    icmp.EntityData.Leafs.Append("icmpInEchos", types.YLeaf{"IcmpInEchos", icmp.IcmpInEchos})
    icmp.EntityData.Leafs.Append("icmpInEchoReps", types.YLeaf{"IcmpInEchoReps", icmp.IcmpInEchoReps})
    icmp.EntityData.Leafs.Append("icmpInTimestamps", types.YLeaf{"IcmpInTimestamps", icmp.IcmpInTimestamps})
    icmp.EntityData.Leafs.Append("icmpInTimestampReps", types.YLeaf{"IcmpInTimestampReps", icmp.IcmpInTimestampReps})
    icmp.EntityData.Leafs.Append("icmpInAddrMasks", types.YLeaf{"IcmpInAddrMasks", icmp.IcmpInAddrMasks})
    icmp.EntityData.Leafs.Append("icmpInAddrMaskReps", types.YLeaf{"IcmpInAddrMaskReps", icmp.IcmpInAddrMaskReps})
    icmp.EntityData.Leafs.Append("icmpOutMsgs", types.YLeaf{"IcmpOutMsgs", icmp.IcmpOutMsgs})
    icmp.EntityData.Leafs.Append("icmpOutErrors", types.YLeaf{"IcmpOutErrors", icmp.IcmpOutErrors})
    icmp.EntityData.Leafs.Append("icmpOutDestUnreachs", types.YLeaf{"IcmpOutDestUnreachs", icmp.IcmpOutDestUnreachs})
    icmp.EntityData.Leafs.Append("icmpOutTimeExcds", types.YLeaf{"IcmpOutTimeExcds", icmp.IcmpOutTimeExcds})
    icmp.EntityData.Leafs.Append("icmpOutParmProbs", types.YLeaf{"IcmpOutParmProbs", icmp.IcmpOutParmProbs})
    icmp.EntityData.Leafs.Append("icmpOutSrcQuenchs", types.YLeaf{"IcmpOutSrcQuenchs", icmp.IcmpOutSrcQuenchs})
    icmp.EntityData.Leafs.Append("icmpOutRedirects", types.YLeaf{"IcmpOutRedirects", icmp.IcmpOutRedirects})
    icmp.EntityData.Leafs.Append("icmpOutEchos", types.YLeaf{"IcmpOutEchos", icmp.IcmpOutEchos})
    icmp.EntityData.Leafs.Append("icmpOutEchoReps", types.YLeaf{"IcmpOutEchoReps", icmp.IcmpOutEchoReps})
    icmp.EntityData.Leafs.Append("icmpOutTimestamps", types.YLeaf{"IcmpOutTimestamps", icmp.IcmpOutTimestamps})
    icmp.EntityData.Leafs.Append("icmpOutTimestampReps", types.YLeaf{"IcmpOutTimestampReps", icmp.IcmpOutTimestampReps})
    icmp.EntityData.Leafs.Append("icmpOutAddrMasks", types.YLeaf{"IcmpOutAddrMasks", icmp.IcmpOutAddrMasks})
    icmp.EntityData.Leafs.Append("icmpOutAddrMaskReps", types.YLeaf{"IcmpOutAddrMaskReps", icmp.IcmpOutAddrMaskReps})

    icmp.EntityData.YListKeys = []string {}

    return &(icmp.EntityData)
}

// RFC1213MIB_Tcp
type RFC1213MIB_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The algorithm used to determine the timeout value used for retransmitting
    // unacknowledged octets. The type is TcpRtoAlgorithm.
    TcpRtoAlgorithm interface{}

    // The minimum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds.  More refined semantics for objects of
    // this type depend upon the algorithm used to determine the retransmission
    // timeout.  In particular, when the timeout algorithm is rsre(3), an object
    // of this type has the semantics of the LBOUND quantity described in RFC 793.
    // The type is interface{} with range: -2147483648..2147483647.
    TcpRtoMin interface{}

    // The maximum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds.  More refined semantics for objects of
    // this type depend upon the algorithm used to determine the retransmission
    // timeout.  In particular, when the timeout algorithm is rsre(3), an object
    // of this type has the semantics of the UBOUND quantity described in RFC 793.
    // The type is interface{} with range: -2147483648..2147483647.
    TcpRtoMax interface{}

    // The limit on the total number of TCP connections the entity can support. 
    // In entities where the maximum number of connections is dynamic, this object
    // should contain the value -1. The type is interface{} with range:
    // -2147483648..2147483647.
    TcpMaxConn interface{}

    // The number of times TCP connections have made a direct transition to the
    // SYN-SENT state from the CLOSED state. The type is interface{} with range:
    // 0..4294967295.
    TcpActiveOpens interface{}

    // The number of times TCP connections have made a direct transition to the
    // SYN-RCVD state from the LISTEN state. The type is interface{} with range:
    // 0..4294967295.
    TcpPassiveOpens interface{}

    // The number of times TCP connections have made a direct transition to the
    // CLOSED state from either the SYN-SENT state or the SYN-RCVD state, plus the
    // number of times TCP connections have made a direct transition to the LISTEN
    // state from the SYN-RCVD state. The type is interface{} with range:
    // 0..4294967295.
    TcpAttemptFails interface{}

    // The number of times TCP connections have made a direct transition to the
    // CLOSED state from either the ESTABLISHED state or the CLOSE-WAIT state. The
    // type is interface{} with range: 0..4294967295.
    TcpEstabResets interface{}

    // The number of TCP connections for which the current state is either
    // ESTABLISHED or CLOSE- WAIT. The type is interface{} with range:
    // 0..4294967295.
    TcpCurrEstab interface{}

    // The total number of segments received, including those received in error. 
    // This count includes segments received on currently established connections.
    // The type is interface{} with range: 0..4294967295.
    TcpInSegs interface{}

    // The total number of segments sent, including those on current connections
    // but excluding those containing only retransmitted octets. The type is
    // interface{} with range: 0..4294967295.
    TcpOutSegs interface{}

    // The total number of segments retransmitted - that is, the number of TCP
    // segments transmitted containing one or more previously transmitted octets.
    // The type is interface{} with range: 0..4294967295.
    TcpRetransSegs interface{}

    // The total number of segments received in error (e.g., bad TCP checksums).
    // The type is interface{} with range: 0..4294967295.
    TcpInErrs interface{}

    // The number of TCP segments sent containing the RST flag. The type is
    // interface{} with range: 0..4294967295.
    TcpOutRsts interface{}
}

func (tcp *RFC1213MIB_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "RFC1213-MIB"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + tcp.EntityData.SegmentPath
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

    tcp.EntityData.YListKeys = []string {}

    return &(tcp.EntityData)
}

// RFC1213MIB_Tcp_TcpRtoAlgorithm represents used for retransmitting unacknowledged octets.
type RFC1213MIB_Tcp_TcpRtoAlgorithm string

const (
    RFC1213MIB_Tcp_TcpRtoAlgorithm_other RFC1213MIB_Tcp_TcpRtoAlgorithm = "other"

    RFC1213MIB_Tcp_TcpRtoAlgorithm_constant RFC1213MIB_Tcp_TcpRtoAlgorithm = "constant"

    RFC1213MIB_Tcp_TcpRtoAlgorithm_rsre RFC1213MIB_Tcp_TcpRtoAlgorithm = "rsre"

    RFC1213MIB_Tcp_TcpRtoAlgorithm_vanj RFC1213MIB_Tcp_TcpRtoAlgorithm = "vanj"

    RFC1213MIB_Tcp_TcpRtoAlgorithm_rfc2988 RFC1213MIB_Tcp_TcpRtoAlgorithm = "rfc2988"
)

// RFC1213MIB_Udp
type RFC1213MIB_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of UDP datagrams delivered to UDP users. The type is
    // interface{} with range: 0..4294967295.
    UdpInDatagrams interface{}

    // The total number of received UDP datagrams for which there was no
    // application at the destination port. The type is interface{} with range:
    // 0..4294967295.
    UdpNoPorts interface{}

    // The number of received UDP datagrams that could not be delivered for
    // reasons other than the lack of an application at the destination port. The
    // type is interface{} with range: 0..4294967295.
    UdpInErrors interface{}

    // The total number of UDP datagrams sent from this entity. The type is
    // interface{} with range: 0..4294967295.
    UdpOutDatagrams interface{}
}

func (udp *RFC1213MIB_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xe"
    udp.EntityData.ParentYangName = "RFC1213-MIB"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + udp.EntityData.SegmentPath
    udp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udp.EntityData.Children = types.NewOrderedMap()
    udp.EntityData.Leafs = types.NewOrderedMap()
    udp.EntityData.Leafs.Append("udpInDatagrams", types.YLeaf{"UdpInDatagrams", udp.UdpInDatagrams})
    udp.EntityData.Leafs.Append("udpNoPorts", types.YLeaf{"UdpNoPorts", udp.UdpNoPorts})
    udp.EntityData.Leafs.Append("udpInErrors", types.YLeaf{"UdpInErrors", udp.UdpInErrors})
    udp.EntityData.Leafs.Append("udpOutDatagrams", types.YLeaf{"UdpOutDatagrams", udp.UdpOutDatagrams})

    udp.EntityData.YListKeys = []string {}

    return &(udp.EntityData)
}

// RFC1213MIB_Egp
type RFC1213MIB_Egp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of EGP messages received without error. The type is interface{}
    // with range: 0..4294967295.
    EgpInMsgs interface{}

    // The number of EGP messages received that proved to be in error. The type is
    // interface{} with range: 0..4294967295.
    EgpInErrors interface{}

    // The total number of locally generated EGP messages. The type is interface{}
    // with range: 0..4294967295.
    EgpOutMsgs interface{}

    // The number of locally generated EGP messages not sent due to resource
    // limitations within an EGP entity. The type is interface{} with range:
    // 0..4294967295.
    EgpOutErrors interface{}

    // The autonomous system number of this EGP entity. The type is interface{}
    // with range: -2147483648..2147483647.
    EgpAs interface{}
}

func (egp *RFC1213MIB_Egp) GetEntityData() *types.CommonEntityData {
    egp.EntityData.YFilter = egp.YFilter
    egp.EntityData.YangName = "egp"
    egp.EntityData.BundleName = "cisco_ios_xe"
    egp.EntityData.ParentYangName = "RFC1213-MIB"
    egp.EntityData.SegmentPath = "egp"
    egp.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + egp.EntityData.SegmentPath
    egp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egp.EntityData.Children = types.NewOrderedMap()
    egp.EntityData.Leafs = types.NewOrderedMap()
    egp.EntityData.Leafs.Append("egpInMsgs", types.YLeaf{"EgpInMsgs", egp.EgpInMsgs})
    egp.EntityData.Leafs.Append("egpInErrors", types.YLeaf{"EgpInErrors", egp.EgpInErrors})
    egp.EntityData.Leafs.Append("egpOutMsgs", types.YLeaf{"EgpOutMsgs", egp.EgpOutMsgs})
    egp.EntityData.Leafs.Append("egpOutErrors", types.YLeaf{"EgpOutErrors", egp.EgpOutErrors})
    egp.EntityData.Leafs.Append("egpAs", types.YLeaf{"EgpAs", egp.EgpAs})

    egp.EntityData.YListKeys = []string {}

    return &(egp.EntityData)
}

// RFC1213MIB_Snmp
type RFC1213MIB_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of Messages delivered to the SNMP entity from the
    // transport service. The type is interface{} with range: 0..4294967295.
    SnmpInPkts interface{}

    // The total number of SNMP Messages which were passed from the SNMP protocol
    // entity to the transport service. The type is interface{} with range:
    // 0..4294967295.
    SnmpOutPkts interface{}

    // The total number of SNMP Messages which were delivered to the SNMP protocol
    // entity and were for an unsupported SNMP version. The type is interface{}
    // with range: 0..4294967295.
    SnmpInBadVersions interface{}

    // The total number of SNMP Messages delivered to the SNMP protocol entity
    // which used a SNMP community name not known to said entity. The type is
    // interface{} with range: 0..4294967295.
    SnmpInBadCommunityNames interface{}

    // The total number of SNMP Messages delivered to the SNMP protocol entity
    // which represented an SNMP operation which was not allowed by the SNMP
    // community named in the Message. The type is interface{} with range:
    // 0..4294967295.
    SnmpInBadCommunityUses interface{}

    // The total number of ASN.1 or BER errors encountered by the SNMP protocol
    // entity when decoding received SNMP Messages. The type is interface{} with
    // range: 0..4294967295.
    SnmpInASNParseErrs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `tooBig'. The
    // type is interface{} with range: 0..4294967295.
    SnmpInTooBigs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `noSuchName'.
    // The type is interface{} with range: 0..4294967295.
    SnmpInNoSuchNames interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `badValue'. The
    // type is interface{} with range: 0..4294967295.
    SnmpInBadValues interface{}

    // The total number valid SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `readOnly'.  It
    // should be noted that it is a protocol error to generate an SNMP PDU which
    // contains the value `readOnly' in the error-status field, as such this
    // object is provided as a means of detecting incorrect implementations of the
    // SNMP. The type is interface{} with range: 0..4294967295.
    SnmpInReadOnlys interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `genErr'. The
    // type is interface{} with range: 0..4294967295.
    SnmpInGenErrs interface{}

    // The total number of MIB objects which have been retrieved successfully by
    // the SNMP protocol entity as the result of receiving valid SNMP Get-Request
    // and Get-Next PDUs. The type is interface{} with range: 0..4294967295.
    SnmpInTotalReqVars interface{}

    // The total number of MIB objects which have been altered successfully by the
    // SNMP protocol entity as the result of receiving valid SNMP Set-Request
    // PDUs. The type is interface{} with range: 0..4294967295.
    SnmpInTotalSetVars interface{}

    // The total number of SNMP Get-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInGetRequests interface{}

    // The total number of SNMP Get-Next PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInGetNexts interface{}

    // The total number of SNMP Set-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInSetRequests interface{}

    // The total number of SNMP Get-Response PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInGetResponses interface{}

    // The total number of SNMP Trap PDUs which have been accepted and processed
    // by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInTraps interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field is `tooBig.'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutTooBigs interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status is `noSuchName'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutNoSuchNames interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field is `badValue'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutBadValues interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field is `genErr'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutGenErrs interface{}

    // The total number of SNMP Get-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutGetRequests interface{}

    // The total number of SNMP Get-Next PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutGetNexts interface{}

    // The total number of SNMP Set-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutSetRequests interface{}

    // The total number of SNMP Get-Response PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutGetResponses interface{}

    // The total number of SNMP Trap PDUs which have been generated by the SNMP
    // protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutTraps interface{}

    // Indicates whether the SNMP agent process is permitted to generate
    // authentication-failure traps.  The value of this object overrides any
    // configuration information; as such, it provides a means whereby all
    // authentication-failure traps may be disabled.  Note that it is strongly
    // recommended that this object be stored in non-volatile memory so that it
    // remains constant between re-initializations of the network management
    // system. The type is SnmpEnableAuthenTraps.
    SnmpEnableAuthenTraps interface{}
}

func (snmp *RFC1213MIB_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xe"
    snmp.EntityData.ParentYangName = "RFC1213-MIB"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + snmp.EntityData.SegmentPath
    snmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("snmpInPkts", types.YLeaf{"SnmpInPkts", snmp.SnmpInPkts})
    snmp.EntityData.Leafs.Append("snmpOutPkts", types.YLeaf{"SnmpOutPkts", snmp.SnmpOutPkts})
    snmp.EntityData.Leafs.Append("snmpInBadVersions", types.YLeaf{"SnmpInBadVersions", snmp.SnmpInBadVersions})
    snmp.EntityData.Leafs.Append("snmpInBadCommunityNames", types.YLeaf{"SnmpInBadCommunityNames", snmp.SnmpInBadCommunityNames})
    snmp.EntityData.Leafs.Append("snmpInBadCommunityUses", types.YLeaf{"SnmpInBadCommunityUses", snmp.SnmpInBadCommunityUses})
    snmp.EntityData.Leafs.Append("snmpInASNParseErrs", types.YLeaf{"SnmpInASNParseErrs", snmp.SnmpInASNParseErrs})
    snmp.EntityData.Leafs.Append("snmpInTooBigs", types.YLeaf{"SnmpInTooBigs", snmp.SnmpInTooBigs})
    snmp.EntityData.Leafs.Append("snmpInNoSuchNames", types.YLeaf{"SnmpInNoSuchNames", snmp.SnmpInNoSuchNames})
    snmp.EntityData.Leafs.Append("snmpInBadValues", types.YLeaf{"SnmpInBadValues", snmp.SnmpInBadValues})
    snmp.EntityData.Leafs.Append("snmpInReadOnlys", types.YLeaf{"SnmpInReadOnlys", snmp.SnmpInReadOnlys})
    snmp.EntityData.Leafs.Append("snmpInGenErrs", types.YLeaf{"SnmpInGenErrs", snmp.SnmpInGenErrs})
    snmp.EntityData.Leafs.Append("snmpInTotalReqVars", types.YLeaf{"SnmpInTotalReqVars", snmp.SnmpInTotalReqVars})
    snmp.EntityData.Leafs.Append("snmpInTotalSetVars", types.YLeaf{"SnmpInTotalSetVars", snmp.SnmpInTotalSetVars})
    snmp.EntityData.Leafs.Append("snmpInGetRequests", types.YLeaf{"SnmpInGetRequests", snmp.SnmpInGetRequests})
    snmp.EntityData.Leafs.Append("snmpInGetNexts", types.YLeaf{"SnmpInGetNexts", snmp.SnmpInGetNexts})
    snmp.EntityData.Leafs.Append("snmpInSetRequests", types.YLeaf{"SnmpInSetRequests", snmp.SnmpInSetRequests})
    snmp.EntityData.Leafs.Append("snmpInGetResponses", types.YLeaf{"SnmpInGetResponses", snmp.SnmpInGetResponses})
    snmp.EntityData.Leafs.Append("snmpInTraps", types.YLeaf{"SnmpInTraps", snmp.SnmpInTraps})
    snmp.EntityData.Leafs.Append("snmpOutTooBigs", types.YLeaf{"SnmpOutTooBigs", snmp.SnmpOutTooBigs})
    snmp.EntityData.Leafs.Append("snmpOutNoSuchNames", types.YLeaf{"SnmpOutNoSuchNames", snmp.SnmpOutNoSuchNames})
    snmp.EntityData.Leafs.Append("snmpOutBadValues", types.YLeaf{"SnmpOutBadValues", snmp.SnmpOutBadValues})
    snmp.EntityData.Leafs.Append("snmpOutGenErrs", types.YLeaf{"SnmpOutGenErrs", snmp.SnmpOutGenErrs})
    snmp.EntityData.Leafs.Append("snmpOutGetRequests", types.YLeaf{"SnmpOutGetRequests", snmp.SnmpOutGetRequests})
    snmp.EntityData.Leafs.Append("snmpOutGetNexts", types.YLeaf{"SnmpOutGetNexts", snmp.SnmpOutGetNexts})
    snmp.EntityData.Leafs.Append("snmpOutSetRequests", types.YLeaf{"SnmpOutSetRequests", snmp.SnmpOutSetRequests})
    snmp.EntityData.Leafs.Append("snmpOutGetResponses", types.YLeaf{"SnmpOutGetResponses", snmp.SnmpOutGetResponses})
    snmp.EntityData.Leafs.Append("snmpOutTraps", types.YLeaf{"SnmpOutTraps", snmp.SnmpOutTraps})
    snmp.EntityData.Leafs.Append("snmpEnableAuthenTraps", types.YLeaf{"SnmpEnableAuthenTraps", snmp.SnmpEnableAuthenTraps})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// RFC1213MIB_Snmp_SnmpEnableAuthenTraps represents network management system.
type RFC1213MIB_Snmp_SnmpEnableAuthenTraps string

const (
    RFC1213MIB_Snmp_SnmpEnableAuthenTraps_enabled RFC1213MIB_Snmp_SnmpEnableAuthenTraps = "enabled"

    RFC1213MIB_Snmp_SnmpEnableAuthenTraps_disabled RFC1213MIB_Snmp_SnmpEnableAuthenTraps = "disabled"
)

// RFC1213MIB_IfTable
// A list of interface entries.  The number of
// entries is given by the value of ifNumber.
type RFC1213MIB_IfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An interface entry containing objects at the subnetwork layer and below for
    // a particular interface. The type is slice of RFC1213MIB_IfTable_IfEntry.
    IfEntry []*RFC1213MIB_IfTable_IfEntry
}

func (ifTable *RFC1213MIB_IfTable) GetEntityData() *types.CommonEntityData {
    ifTable.EntityData.YFilter = ifTable.YFilter
    ifTable.EntityData.YangName = "ifTable"
    ifTable.EntityData.BundleName = "cisco_ios_xe"
    ifTable.EntityData.ParentYangName = "RFC1213-MIB"
    ifTable.EntityData.SegmentPath = "ifTable"
    ifTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + ifTable.EntityData.SegmentPath
    ifTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifTable.EntityData.Children = types.NewOrderedMap()
    ifTable.EntityData.Children.Append("ifEntry", types.YChild{"IfEntry", nil})
    for i := range ifTable.IfEntry {
        ifTable.EntityData.Children.Append(types.GetSegmentPath(ifTable.IfEntry[i]), types.YChild{"IfEntry", ifTable.IfEntry[i]})
    }
    ifTable.EntityData.Leafs = types.NewOrderedMap()

    ifTable.EntityData.YListKeys = []string {}

    return &(ifTable.EntityData)
}

// RFC1213MIB_IfTable_IfEntry
// An interface entry containing objects at the
// subnetwork layer and below for a particular
// interface.
type RFC1213MIB_IfTable_IfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique value for each interface.  Its value
    // ranges between 1 and the value of ifNumber.  The value for each interface
    // must remain constant at least from one re-initialization of the entity's
    // network management system to the next re- initialization. The type is
    // interface{} with range: -2147483648..2147483647.
    IfIndex interface{}

    // A textual string containing information about the interface.  This string
    // should include the name of the manufacturer, the product name and the
    // version of the hardware interface. The type is string with length: 0..255.
    IfDescr interface{}

    // The type of interface.  Additional values for ifType are assigned by the
    // Internet Assigned Numbers Authority (IANA), through updating the syntax of
    // the IANAifType textual convention. The type is IANAifType.
    IfType interface{}

    // The size of the largest datagram which can be sent/received on the
    // interface, specified in octets.  For interfaces that are used for
    // transmitting network datagrams, this is the size of the largest network
    // datagram that can be sent on the interface. The type is interface{} with
    // range: -2147483648..2147483647.
    IfMtu interface{}

    // An estimate of the interface's current bandwidth in bits per second.  For
    // interfaces which do not vary in bandwidth or for those where no accurate
    // estimation can be made, this object should contain the nominal bandwidth.
    // The type is interface{} with range: 0..4294967295.
    IfSpeed interface{}

    // The interface's address at the protocol layer immediately `below' the
    // network layer in the protocol stack.  For interfaces which do not have such
    // an address (e.g., a serial line), this object should contain an octet
    // string of zero length. The type is string.
    IfPhysAddress interface{}

    // The desired state of the interface.  The testing(3) state indicates that no
    // operational packets can be passed. The type is IfAdminStatus.
    IfAdminStatus interface{}

    // The current operational state of the interface. The testing(3) state
    // indicates that no operational packets can be passed. The type is
    // IfOperStatus.
    IfOperStatus interface{}

    // The value of sysUpTime at the time the interface entered its current
    // operational state.  If the current state was entered prior to the last re-
    // initialization of the local network management subsystem, then this object
    // contains a zero value. The type is interface{} with range: 0..4294967295.
    IfLastChange interface{}

    // The total number of octets received on the interface, including framing
    // characters. The type is interface{} with range: 0..4294967295.
    IfInOctets interface{}

    // The number of subnetwork-unicast packets delivered to a higher-layer
    // protocol. The type is interface{} with range: 0..4294967295.
    IfInUcastPkts interface{}

    // The number of non-unicast (i.e., subnetwork- broadcast or
    // subnetwork-multicast) packets delivered to a higher-layer protocol. The
    // type is interface{} with range: 0..4294967295.
    IfInNUcastPkts interface{}

    // The number of inbound packets which were chosen to be discarded even though
    // no errors had been detected to prevent their being deliverable to a
    // higher-layer protocol.  One possible reason for discarding such a packet
    // could be to free up buffer space. The type is interface{} with range:
    // 0..4294967295.
    IfInDiscards interface{}

    // The number of inbound packets that contained errors preventing them from
    // being deliverable to a higher-layer protocol. The type is interface{} with
    // range: 0..4294967295.
    IfInErrors interface{}

    // The number of packets received via the interface which were discarded
    // because of an unknown or unsupported protocol. The type is interface{} with
    // range: 0..4294967295.
    IfInUnknownProtos interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters. The type is interface{} with range: 0..4294967295.
    IfOutOctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted to a subnetwork-unicast address, including those that were
    // discarded or not sent. The type is interface{} with range: 0..4294967295.
    IfOutUcastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted to a non- unicast (i.e., a subnetwork-broadcast or
    // subnetwork-multicast) address, including those that were discarded or not
    // sent. The type is interface{} with range: 0..4294967295.
    IfOutNUcastPkts interface{}

    // The number of outbound packets which were chosen to be discarded even
    // though no errors had been detected to prevent their being transmitted.  One
    // possible reason for discarding such a packet could be to free up buffer
    // space. The type is interface{} with range: 0..4294967295.
    IfOutDiscards interface{}

    // The number of outbound packets that could not be transmitted because of
    // errors. The type is interface{} with range: 0..4294967295.
    IfOutErrors interface{}

    // The length of the output packet queue (in packets). The type is interface{}
    // with range: 0..4294967295.
    IfOutQLen interface{}

    // A reference to MIB definitions specific to the particular media being used
    // to realize the interface.  For example, if the interface is realized by an
    // ethernet, then the value of this object refers to a document defining
    // objects specific to ethernet.  If this information is not present, its
    // value should be set to the OBJECT IDENTIFIER { 0 0 }, which is a
    // syntactically valid object identifier, and any conformant implementation of
    // ASN.1 and BER must be able to generate and recognize this value. The type
    // is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    IfSpecific interface{}
}

func (ifEntry *RFC1213MIB_IfTable_IfEntry) GetEntityData() *types.CommonEntityData {
    ifEntry.EntityData.YFilter = ifEntry.YFilter
    ifEntry.EntityData.YangName = "ifEntry"
    ifEntry.EntityData.BundleName = "cisco_ios_xe"
    ifEntry.EntityData.ParentYangName = "ifTable"
    ifEntry.EntityData.SegmentPath = "ifEntry" + types.AddKeyToken(ifEntry.IfIndex, "ifIndex")
    ifEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/ifTable/" + ifEntry.EntityData.SegmentPath
    ifEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifEntry.EntityData.Children = types.NewOrderedMap()
    ifEntry.EntityData.Leafs = types.NewOrderedMap()
    ifEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ifEntry.IfIndex})
    ifEntry.EntityData.Leafs.Append("ifDescr", types.YLeaf{"IfDescr", ifEntry.IfDescr})
    ifEntry.EntityData.Leafs.Append("ifType", types.YLeaf{"IfType", ifEntry.IfType})
    ifEntry.EntityData.Leafs.Append("ifMtu", types.YLeaf{"IfMtu", ifEntry.IfMtu})
    ifEntry.EntityData.Leafs.Append("ifSpeed", types.YLeaf{"IfSpeed", ifEntry.IfSpeed})
    ifEntry.EntityData.Leafs.Append("ifPhysAddress", types.YLeaf{"IfPhysAddress", ifEntry.IfPhysAddress})
    ifEntry.EntityData.Leafs.Append("ifAdminStatus", types.YLeaf{"IfAdminStatus", ifEntry.IfAdminStatus})
    ifEntry.EntityData.Leafs.Append("ifOperStatus", types.YLeaf{"IfOperStatus", ifEntry.IfOperStatus})
    ifEntry.EntityData.Leafs.Append("ifLastChange", types.YLeaf{"IfLastChange", ifEntry.IfLastChange})
    ifEntry.EntityData.Leafs.Append("ifInOctets", types.YLeaf{"IfInOctets", ifEntry.IfInOctets})
    ifEntry.EntityData.Leafs.Append("ifInUcastPkts", types.YLeaf{"IfInUcastPkts", ifEntry.IfInUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifInNUcastPkts", types.YLeaf{"IfInNUcastPkts", ifEntry.IfInNUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifInDiscards", types.YLeaf{"IfInDiscards", ifEntry.IfInDiscards})
    ifEntry.EntityData.Leafs.Append("ifInErrors", types.YLeaf{"IfInErrors", ifEntry.IfInErrors})
    ifEntry.EntityData.Leafs.Append("ifInUnknownProtos", types.YLeaf{"IfInUnknownProtos", ifEntry.IfInUnknownProtos})
    ifEntry.EntityData.Leafs.Append("ifOutOctets", types.YLeaf{"IfOutOctets", ifEntry.IfOutOctets})
    ifEntry.EntityData.Leafs.Append("ifOutUcastPkts", types.YLeaf{"IfOutUcastPkts", ifEntry.IfOutUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifOutNUcastPkts", types.YLeaf{"IfOutNUcastPkts", ifEntry.IfOutNUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifOutDiscards", types.YLeaf{"IfOutDiscards", ifEntry.IfOutDiscards})
    ifEntry.EntityData.Leafs.Append("ifOutErrors", types.YLeaf{"IfOutErrors", ifEntry.IfOutErrors})
    ifEntry.EntityData.Leafs.Append("ifOutQLen", types.YLeaf{"IfOutQLen", ifEntry.IfOutQLen})
    ifEntry.EntityData.Leafs.Append("ifSpecific", types.YLeaf{"IfSpecific", ifEntry.IfSpecific})

    ifEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(ifEntry.EntityData)
}

// RFC1213MIB_IfTable_IfEntry_IfAdminStatus represents packets can be passed.
type RFC1213MIB_IfTable_IfEntry_IfAdminStatus string

const (
    RFC1213MIB_IfTable_IfEntry_IfAdminStatus_up RFC1213MIB_IfTable_IfEntry_IfAdminStatus = "up"

    RFC1213MIB_IfTable_IfEntry_IfAdminStatus_down RFC1213MIB_IfTable_IfEntry_IfAdminStatus = "down"

    RFC1213MIB_IfTable_IfEntry_IfAdminStatus_testing RFC1213MIB_IfTable_IfEntry_IfAdminStatus = "testing"
)

// RFC1213MIB_IfTable_IfEntry_IfOperStatus represents packets can be passed.
type RFC1213MIB_IfTable_IfEntry_IfOperStatus string

const (
    RFC1213MIB_IfTable_IfEntry_IfOperStatus_up RFC1213MIB_IfTable_IfEntry_IfOperStatus = "up"

    RFC1213MIB_IfTable_IfEntry_IfOperStatus_down RFC1213MIB_IfTable_IfEntry_IfOperStatus = "down"

    RFC1213MIB_IfTable_IfEntry_IfOperStatus_testing RFC1213MIB_IfTable_IfEntry_IfOperStatus = "testing"

    RFC1213MIB_IfTable_IfEntry_IfOperStatus_unknown RFC1213MIB_IfTable_IfEntry_IfOperStatus = "unknown"

    RFC1213MIB_IfTable_IfEntry_IfOperStatus_dormant RFC1213MIB_IfTable_IfEntry_IfOperStatus = "dormant"
)

// RFC1213MIB_AtTable
// The Address Translation tables contain the
// NetworkAddress to `physical' address equivalences.
// Some interfaces do not use translation tables for
// determining address equivalences (e.g., DDN-X.25
// has an algorithmic method); if all interfaces are
// of this type, then the Address Translation table
// is empty, i.e., has zero entries.
type RFC1213MIB_AtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one NetworkAddress to `physical' address equivalence.
    // The type is slice of RFC1213MIB_AtTable_AtEntry.
    AtEntry []*RFC1213MIB_AtTable_AtEntry
}

func (atTable *RFC1213MIB_AtTable) GetEntityData() *types.CommonEntityData {
    atTable.EntityData.YFilter = atTable.YFilter
    atTable.EntityData.YangName = "atTable"
    atTable.EntityData.BundleName = "cisco_ios_xe"
    atTable.EntityData.ParentYangName = "RFC1213-MIB"
    atTable.EntityData.SegmentPath = "atTable"
    atTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + atTable.EntityData.SegmentPath
    atTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atTable.EntityData.Children = types.NewOrderedMap()
    atTable.EntityData.Children.Append("atEntry", types.YChild{"AtEntry", nil})
    for i := range atTable.AtEntry {
        atTable.EntityData.Children.Append(types.GetSegmentPath(atTable.AtEntry[i]), types.YChild{"AtEntry", atTable.AtEntry[i]})
    }
    atTable.EntityData.Leafs = types.NewOrderedMap()

    atTable.EntityData.YListKeys = []string {}

    return &(atTable.EntityData)
}

// RFC1213MIB_AtTable_AtEntry
// Each entry contains one NetworkAddress to
// `physical' address equivalence.
type RFC1213MIB_AtTable_AtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    AtIfIndex interface{}

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    AtIfIndex2 interface{}

    // This attribute is a key. The NetworkAddress (e.g., the IP address)
    // corresponding to the media-dependent `physical' address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AtNetAddress interface{}

    // The media-dependent `physical' address.  Setting this object to a null
    // string (one of zero length) has the effect of invaliding the corresponding
    // entry in the atTable object.  That is, it effectively disassociates the
    // interface identified with said entry from the mapping identified with said
    // entry.  It is an implementation-specific matter as to whether the agent
    // removes an invalidated entry from the table. Accordingly, management
    // stations must be prepared to receive tabular information from agents that
    // corresponds to entries not currently in use. Proper interpretation of such
    // entries requires examination of the relevant atPhysAddress object. The type
    // is string.
    AtPhysAddress interface{}
}

func (atEntry *RFC1213MIB_AtTable_AtEntry) GetEntityData() *types.CommonEntityData {
    atEntry.EntityData.YFilter = atEntry.YFilter
    atEntry.EntityData.YangName = "atEntry"
    atEntry.EntityData.BundleName = "cisco_ios_xe"
    atEntry.EntityData.ParentYangName = "atTable"
    atEntry.EntityData.SegmentPath = "atEntry" + types.AddKeyToken(atEntry.AtIfIndex, "atIfIndex") + types.AddKeyToken(atEntry.AtIfIndex2, "atIfIndex_2") + types.AddKeyToken(atEntry.AtNetAddress, "atNetAddress")
    atEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/atTable/" + atEntry.EntityData.SegmentPath
    atEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atEntry.EntityData.Children = types.NewOrderedMap()
    atEntry.EntityData.Leafs = types.NewOrderedMap()
    atEntry.EntityData.Leafs.Append("atIfIndex", types.YLeaf{"AtIfIndex", atEntry.AtIfIndex})
    atEntry.EntityData.Leafs.Append("atIfIndex_2", types.YLeaf{"AtIfIndex2", atEntry.AtIfIndex2})
    atEntry.EntityData.Leafs.Append("atNetAddress", types.YLeaf{"AtNetAddress", atEntry.AtNetAddress})
    atEntry.EntityData.Leafs.Append("atPhysAddress", types.YLeaf{"AtPhysAddress", atEntry.AtPhysAddress})

    atEntry.EntityData.YListKeys = []string {"AtIfIndex", "AtIfIndex2", "AtNetAddress"}

    return &(atEntry.EntityData)
}

// RFC1213MIB_IpAddrTable
// The table of addressing information relevant to
// this entity's IP addresses.
type RFC1213MIB_IpAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The addressing information for one of this entity's IP addresses. The type
    // is slice of RFC1213MIB_IpAddrTable_IpAddrEntry.
    IpAddrEntry []*RFC1213MIB_IpAddrTable_IpAddrEntry
}

func (ipAddrTable *RFC1213MIB_IpAddrTable) GetEntityData() *types.CommonEntityData {
    ipAddrTable.EntityData.YFilter = ipAddrTable.YFilter
    ipAddrTable.EntityData.YangName = "ipAddrTable"
    ipAddrTable.EntityData.BundleName = "cisco_ios_xe"
    ipAddrTable.EntityData.ParentYangName = "RFC1213-MIB"
    ipAddrTable.EntityData.SegmentPath = "ipAddrTable"
    ipAddrTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + ipAddrTable.EntityData.SegmentPath
    ipAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipAddrTable.EntityData.Children = types.NewOrderedMap()
    ipAddrTable.EntityData.Children.Append("ipAddrEntry", types.YChild{"IpAddrEntry", nil})
    for i := range ipAddrTable.IpAddrEntry {
        ipAddrTable.EntityData.Children.Append(types.GetSegmentPath(ipAddrTable.IpAddrEntry[i]), types.YChild{"IpAddrEntry", ipAddrTable.IpAddrEntry[i]})
    }
    ipAddrTable.EntityData.Leafs = types.NewOrderedMap()

    ipAddrTable.EntityData.YListKeys = []string {}

    return &(ipAddrTable.EntityData)
}

// RFC1213MIB_IpAddrTable_IpAddrEntry
// The addressing information for one of this
// entity's IP addresses.
type RFC1213MIB_IpAddrTable_IpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP address to which this entry's addressing
    // information pertains. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAdEntAddr interface{}

    // The index value which uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of ifIndex. The
    // type is interface{} with range: -2147483648..2147483647.
    IpAdEntIfIndex interface{}

    // The subnet mask associated with the IP address of this entry.  The value of
    // the mask is an IP address with all the network bits set to 1 and all the
    // hosts bits set to 0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAdEntNetMask interface{}

    // The value of the least-significant bit in the IP broadcast address used for
    // sending datagrams on the (logical) interface associated with the IP address
    // of this entry.  For example, when the Internet standard all-ones broadcast
    // address is used, the value will be 1.  This value applies to both the
    // subnet and network broadcasts addresses used by the entity on this
    // (logical) interface. The type is interface{} with range:
    // -2147483648..2147483647.
    IpAdEntBcastAddr interface{}

    // The size of the largest IP datagram which this entity can re-assemble from
    // incoming IP fragmented datagrams received on this interface. The type is
    // interface{} with range: 0..65535.
    IpAdEntReasmMaxSize interface{}
}

func (ipAddrEntry *RFC1213MIB_IpAddrTable_IpAddrEntry) GetEntityData() *types.CommonEntityData {
    ipAddrEntry.EntityData.YFilter = ipAddrEntry.YFilter
    ipAddrEntry.EntityData.YangName = "ipAddrEntry"
    ipAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    ipAddrEntry.EntityData.ParentYangName = "ipAddrTable"
    ipAddrEntry.EntityData.SegmentPath = "ipAddrEntry" + types.AddKeyToken(ipAddrEntry.IpAdEntAddr, "ipAdEntAddr")
    ipAddrEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/ipAddrTable/" + ipAddrEntry.EntityData.SegmentPath
    ipAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipAddrEntry.EntityData.Children = types.NewOrderedMap()
    ipAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    ipAddrEntry.EntityData.Leafs.Append("ipAdEntAddr", types.YLeaf{"IpAdEntAddr", ipAddrEntry.IpAdEntAddr})
    ipAddrEntry.EntityData.Leafs.Append("ipAdEntIfIndex", types.YLeaf{"IpAdEntIfIndex", ipAddrEntry.IpAdEntIfIndex})
    ipAddrEntry.EntityData.Leafs.Append("ipAdEntNetMask", types.YLeaf{"IpAdEntNetMask", ipAddrEntry.IpAdEntNetMask})
    ipAddrEntry.EntityData.Leafs.Append("ipAdEntBcastAddr", types.YLeaf{"IpAdEntBcastAddr", ipAddrEntry.IpAdEntBcastAddr})
    ipAddrEntry.EntityData.Leafs.Append("ipAdEntReasmMaxSize", types.YLeaf{"IpAdEntReasmMaxSize", ipAddrEntry.IpAdEntReasmMaxSize})

    ipAddrEntry.EntityData.YListKeys = []string {"IpAdEntAddr"}

    return &(ipAddrEntry.EntityData)
}

// RFC1213MIB_IpRouteTable
// This entity's IP Routing table.
type RFC1213MIB_IpRouteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A route to a particular destination. The type is slice of
    // RFC1213MIB_IpRouteTable_IpRouteEntry.
    IpRouteEntry []*RFC1213MIB_IpRouteTable_IpRouteEntry
}

func (ipRouteTable *RFC1213MIB_IpRouteTable) GetEntityData() *types.CommonEntityData {
    ipRouteTable.EntityData.YFilter = ipRouteTable.YFilter
    ipRouteTable.EntityData.YangName = "ipRouteTable"
    ipRouteTable.EntityData.BundleName = "cisco_ios_xe"
    ipRouteTable.EntityData.ParentYangName = "RFC1213-MIB"
    ipRouteTable.EntityData.SegmentPath = "ipRouteTable"
    ipRouteTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + ipRouteTable.EntityData.SegmentPath
    ipRouteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipRouteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipRouteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipRouteTable.EntityData.Children = types.NewOrderedMap()
    ipRouteTable.EntityData.Children.Append("ipRouteEntry", types.YChild{"IpRouteEntry", nil})
    for i := range ipRouteTable.IpRouteEntry {
        ipRouteTable.EntityData.Children.Append(types.GetSegmentPath(ipRouteTable.IpRouteEntry[i]), types.YChild{"IpRouteEntry", ipRouteTable.IpRouteEntry[i]})
    }
    ipRouteTable.EntityData.Leafs = types.NewOrderedMap()

    ipRouteTable.EntityData.YListKeys = []string {}

    return &(ipRouteTable.EntityData)
}

// RFC1213MIB_IpRouteTable_IpRouteEntry
// A route to a particular destination.
type RFC1213MIB_IpRouteTable_IpRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The destination IP address of this route.  An
    // entry with a value of 0.0.0.0 is considered a default route.  Multiple
    // routes to a single destination can appear in the table, but access to such
    // multiple entries is dependent on the table- access mechanisms defined by
    // the network management protocol in use. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpRouteDest interface{}

    // The index value which uniquely identifies the local interface through which
    // the next hop of this route should be reached.  The interface identified by
    // a particular value of this index is the same interface as identified by the
    // same value of ifIndex. The type is interface{} with range:
    // -2147483648..2147483647.
    IpRouteIfIndex interface{}

    // The primary routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    IpRouteMetric1 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    IpRouteMetric2 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    IpRouteMetric3 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    IpRouteMetric4 interface{}

    // The IP address of the next hop of this route. (In the case of a route bound
    // to an interface which is realized via a broadcast media, the value of this
    // field is the agent's IP address on that interface.). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpRouteNextHop interface{}

    // The type of route.  Note that the values direct(3) and indirect(4) refer to
    // the notion of direct and indirect routing in the IP architecture.  Setting
    // this object to the value invalid(2) has the effect of invalidating the
    // corresponding entry in the ipRouteTable object.  That is, it effectively
    // disassociates the destination identified with said entry from the route
    // identified with said entry.  It is an implementation-specific matter as to
    // whether the agent removes an invalidated entry from the table. Accordingly,
    // management stations must be prepared to receive tabular information from
    // agents that corresponds to entries not currently in use. Proper
    // interpretation of such entries requires examination of the relevant
    // ipRouteType object. The type is IpRouteType.
    IpRouteType interface{}

    // The routing mechanism via which this route was learned.  Inclusion of
    // values for gateway routing protocols is not intended to imply that hosts
    // should support those protocols. The type is IpRouteProto.
    IpRouteProto interface{}

    // The number of seconds since this route was last updated or otherwise
    // determined to be correct. Note that no semantics of `too old' can be
    // implied except through knowledge of the routing protocol by which the route
    // was learned. The type is interface{} with range: -2147483648..2147483647.
    IpRouteAge interface{}

    // Indicate the mask to be logical-ANDed with the destination address before
    // being compared to the value in the ipRouteDest field.  For those systems
    // that do not support arbitrary subnet masks, an agent constructs the value
    // of the ipRouteMask by determining whether the value of the correspondent
    // ipRouteDest field belong to a class-A, B, or C network, and then using one
    // of:       mask           network      255.0.0.0      class-A     
    // 255.255.0.0    class-B      255.255.255.0  class-C  If the value of the
    // ipRouteDest is 0.0.0.0 (a default route), then the mask value is also
    // 0.0.0.0.  It should be noted that all IP routing subsystems implicitly use
    // this mechanism. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpRouteMask interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    IpRouteMetric5 interface{}

    // A reference to MIB definitions specific to the particular routing protocol
    // which is responsible for this route, as determined by the value specified
    // in the route's ipRouteProto value.  If this information is not present, its
    // value should be set to the OBJECT IDENTIFIER { 0 0 }, which is a
    // syntactically valid object identifier, and any conformant implementation of
    // ASN.1 and BER must be able to generate and recognize this value. The type
    // is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    IpRouteInfo interface{}
}

func (ipRouteEntry *RFC1213MIB_IpRouteTable_IpRouteEntry) GetEntityData() *types.CommonEntityData {
    ipRouteEntry.EntityData.YFilter = ipRouteEntry.YFilter
    ipRouteEntry.EntityData.YangName = "ipRouteEntry"
    ipRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    ipRouteEntry.EntityData.ParentYangName = "ipRouteTable"
    ipRouteEntry.EntityData.SegmentPath = "ipRouteEntry" + types.AddKeyToken(ipRouteEntry.IpRouteDest, "ipRouteDest")
    ipRouteEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/ipRouteTable/" + ipRouteEntry.EntityData.SegmentPath
    ipRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipRouteEntry.EntityData.Children = types.NewOrderedMap()
    ipRouteEntry.EntityData.Leafs = types.NewOrderedMap()
    ipRouteEntry.EntityData.Leafs.Append("ipRouteDest", types.YLeaf{"IpRouteDest", ipRouteEntry.IpRouteDest})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteIfIndex", types.YLeaf{"IpRouteIfIndex", ipRouteEntry.IpRouteIfIndex})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteMetric1", types.YLeaf{"IpRouteMetric1", ipRouteEntry.IpRouteMetric1})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteMetric2", types.YLeaf{"IpRouteMetric2", ipRouteEntry.IpRouteMetric2})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteMetric3", types.YLeaf{"IpRouteMetric3", ipRouteEntry.IpRouteMetric3})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteMetric4", types.YLeaf{"IpRouteMetric4", ipRouteEntry.IpRouteMetric4})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteNextHop", types.YLeaf{"IpRouteNextHop", ipRouteEntry.IpRouteNextHop})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteType", types.YLeaf{"IpRouteType", ipRouteEntry.IpRouteType})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteProto", types.YLeaf{"IpRouteProto", ipRouteEntry.IpRouteProto})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteAge", types.YLeaf{"IpRouteAge", ipRouteEntry.IpRouteAge})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteMask", types.YLeaf{"IpRouteMask", ipRouteEntry.IpRouteMask})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteMetric5", types.YLeaf{"IpRouteMetric5", ipRouteEntry.IpRouteMetric5})
    ipRouteEntry.EntityData.Leafs.Append("ipRouteInfo", types.YLeaf{"IpRouteInfo", ipRouteEntry.IpRouteInfo})

    ipRouteEntry.EntityData.YListKeys = []string {"IpRouteDest"}

    return &(ipRouteEntry.EntityData)
}

// RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto represents should support those protocols.
type RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto string

const (
    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_other RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "other"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_local RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "local"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_netmgmt RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "netmgmt"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_icmp RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "icmp"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_egp RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "egp"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_ggp RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "ggp"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_hello RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "hello"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_rip RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "rip"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_is_is RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "is-is"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_es_is RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "es-is"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_ciscoIgrp RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "ciscoIgrp"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_bbnSpfIgp RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "bbnSpfIgp"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_ospf RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "ospf"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto_bgp RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteProto = "bgp"
)

// RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType represents examination of the relevant ipRouteType object.
type RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType string

const (
    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType_other RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType = "other"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType_invalid RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType = "invalid"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType_direct RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType = "direct"

    RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType_indirect RFC1213MIB_IpRouteTable_IpRouteEntry_IpRouteType = "indirect"
)

// RFC1213MIB_IpNetToMediaTable
// The IP Address Translation table used for mapping
// from IP addresses to physical addresses.
type RFC1213MIB_IpNetToMediaTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one IpAddress to `physical' address equivalence. The
    // type is slice of RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry.
    IpNetToMediaEntry []*RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry
}

func (ipNetToMediaTable *RFC1213MIB_IpNetToMediaTable) GetEntityData() *types.CommonEntityData {
    ipNetToMediaTable.EntityData.YFilter = ipNetToMediaTable.YFilter
    ipNetToMediaTable.EntityData.YangName = "ipNetToMediaTable"
    ipNetToMediaTable.EntityData.BundleName = "cisco_ios_xe"
    ipNetToMediaTable.EntityData.ParentYangName = "RFC1213-MIB"
    ipNetToMediaTable.EntityData.SegmentPath = "ipNetToMediaTable"
    ipNetToMediaTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + ipNetToMediaTable.EntityData.SegmentPath
    ipNetToMediaTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipNetToMediaTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipNetToMediaTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipNetToMediaTable.EntityData.Children = types.NewOrderedMap()
    ipNetToMediaTable.EntityData.Children.Append("ipNetToMediaEntry", types.YChild{"IpNetToMediaEntry", nil})
    for i := range ipNetToMediaTable.IpNetToMediaEntry {
        ipNetToMediaTable.EntityData.Children.Append(types.GetSegmentPath(ipNetToMediaTable.IpNetToMediaEntry[i]), types.YChild{"IpNetToMediaEntry", ipNetToMediaTable.IpNetToMediaEntry[i]})
    }
    ipNetToMediaTable.EntityData.Leafs = types.NewOrderedMap()

    ipNetToMediaTable.EntityData.YListKeys = []string {}

    return &(ipNetToMediaTable.EntityData)
}

// RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry
// Each entry contains one IpAddress to `physical'
// address equivalence.
type RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    IpNetToMediaIfIndex interface{}

    // This attribute is a key. The IpAddress corresponding to the media-
    // dependent `physical' address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpNetToMediaNetAddress interface{}

    // The media-dependent `physical' address. The type is string.
    IpNetToMediaPhysAddress interface{}

    // The type of mapping.  Setting this object to the value invalid(2) has the
    // effect of invalidating the corresponding entry in the ipNetToMediaTable. 
    // That is, it effectively disassociates the interface identified with said
    // entry from the mapping identified with said entry. It is an
    // implementation-specific matter as to whether the agent removes an
    // invalidated entry from the table.  Accordingly, management stations must be
    // prepared to receive tabular information from agents that corresponds to
    // entries not currently in use.  Proper interpretation of such entries
    // requires examination of the relevant ipNetToMediaType object. The type is
    // IpNetToMediaType.
    IpNetToMediaType interface{}
}

func (ipNetToMediaEntry *RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry) GetEntityData() *types.CommonEntityData {
    ipNetToMediaEntry.EntityData.YFilter = ipNetToMediaEntry.YFilter
    ipNetToMediaEntry.EntityData.YangName = "ipNetToMediaEntry"
    ipNetToMediaEntry.EntityData.BundleName = "cisco_ios_xe"
    ipNetToMediaEntry.EntityData.ParentYangName = "ipNetToMediaTable"
    ipNetToMediaEntry.EntityData.SegmentPath = "ipNetToMediaEntry" + types.AddKeyToken(ipNetToMediaEntry.IpNetToMediaIfIndex, "ipNetToMediaIfIndex") + types.AddKeyToken(ipNetToMediaEntry.IpNetToMediaNetAddress, "ipNetToMediaNetAddress")
    ipNetToMediaEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/ipNetToMediaTable/" + ipNetToMediaEntry.EntityData.SegmentPath
    ipNetToMediaEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipNetToMediaEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipNetToMediaEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipNetToMediaEntry.EntityData.Children = types.NewOrderedMap()
    ipNetToMediaEntry.EntityData.Leafs = types.NewOrderedMap()
    ipNetToMediaEntry.EntityData.Leafs.Append("ipNetToMediaIfIndex", types.YLeaf{"IpNetToMediaIfIndex", ipNetToMediaEntry.IpNetToMediaIfIndex})
    ipNetToMediaEntry.EntityData.Leafs.Append("ipNetToMediaNetAddress", types.YLeaf{"IpNetToMediaNetAddress", ipNetToMediaEntry.IpNetToMediaNetAddress})
    ipNetToMediaEntry.EntityData.Leafs.Append("ipNetToMediaPhysAddress", types.YLeaf{"IpNetToMediaPhysAddress", ipNetToMediaEntry.IpNetToMediaPhysAddress})
    ipNetToMediaEntry.EntityData.Leafs.Append("ipNetToMediaType", types.YLeaf{"IpNetToMediaType", ipNetToMediaEntry.IpNetToMediaType})

    ipNetToMediaEntry.EntityData.YListKeys = []string {"IpNetToMediaIfIndex", "IpNetToMediaNetAddress"}

    return &(ipNetToMediaEntry.EntityData)
}

// RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType represents ipNetToMediaType object.
type RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType string

const (
    RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_other RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "other"

    RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_invalid RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "invalid"

    RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_dynamic RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "dynamic"

    RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_static RFC1213MIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "static"
)

// RFC1213MIB_TcpConnTable
// A table containing TCP connection-specific
// information.
type RFC1213MIB_TcpConnTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current TCP connection.  An object of this
    // type is transient, in that it ceases to exist when (or soon after) the
    // connection makes the transition to the CLOSED state. The type is slice of
    // RFC1213MIB_TcpConnTable_TcpConnEntry.
    TcpConnEntry []*RFC1213MIB_TcpConnTable_TcpConnEntry
}

func (tcpConnTable *RFC1213MIB_TcpConnTable) GetEntityData() *types.CommonEntityData {
    tcpConnTable.EntityData.YFilter = tcpConnTable.YFilter
    tcpConnTable.EntityData.YangName = "tcpConnTable"
    tcpConnTable.EntityData.BundleName = "cisco_ios_xe"
    tcpConnTable.EntityData.ParentYangName = "RFC1213-MIB"
    tcpConnTable.EntityData.SegmentPath = "tcpConnTable"
    tcpConnTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + tcpConnTable.EntityData.SegmentPath
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

// RFC1213MIB_TcpConnTable_TcpConnEntry
// Information about a particular current TCP
// connection.  An object of this type is transient,
// in that it ceases to exist when (or soon after)
// the connection makes the transition to the CLOSED
// state.
type RFC1213MIB_TcpConnTable_TcpConnEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The local IP address for this TCP connection.  In
    // the case of a connection in the listen state which is willing to accept
    // connections for any IP interface associated with the node, the value
    // 0.0.0.0 is used. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    TcpConnLocalAddress interface{}

    // This attribute is a key. The local port number for this TCP connection. The
    // type is interface{} with range: 0..65535.
    TcpConnLocalPort interface{}

    // This attribute is a key. The remote IP address for this TCP connection. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    TcpConnRemAddress interface{}

    // This attribute is a key. The remote port number for this TCP connection.
    // The type is interface{} with range: 0..65535.
    TcpConnRemPort interface{}

    // The state of this TCP connection.  The only value which may be set by a
    // management station is deleteTCB(12).  Accordingly, it is appropriate for an
    // agent to return a `badValue' response if a management station attempts to
    // set this object to any other value.  If a management station sets this
    // object to the value deleteTCB(12), then this has the effect of deleting the
    // TCB (as defined in RFC 793) of the corresponding connection on the managed
    // node, resulting in immediate termination of the connection.  As an
    // implementation-specific option, a RST segment may be sent from the managed
    // node to the other TCP endpoint (note however that RST segments are not sent
    // reliably). The type is TcpConnState.
    TcpConnState interface{}
}

func (tcpConnEntry *RFC1213MIB_TcpConnTable_TcpConnEntry) GetEntityData() *types.CommonEntityData {
    tcpConnEntry.EntityData.YFilter = tcpConnEntry.YFilter
    tcpConnEntry.EntityData.YangName = "tcpConnEntry"
    tcpConnEntry.EntityData.BundleName = "cisco_ios_xe"
    tcpConnEntry.EntityData.ParentYangName = "tcpConnTable"
    tcpConnEntry.EntityData.SegmentPath = "tcpConnEntry" + types.AddKeyToken(tcpConnEntry.TcpConnLocalAddress, "tcpConnLocalAddress") + types.AddKeyToken(tcpConnEntry.TcpConnLocalPort, "tcpConnLocalPort") + types.AddKeyToken(tcpConnEntry.TcpConnRemAddress, "tcpConnRemAddress") + types.AddKeyToken(tcpConnEntry.TcpConnRemPort, "tcpConnRemPort")
    tcpConnEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/tcpConnTable/" + tcpConnEntry.EntityData.SegmentPath
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

// RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState represents are not sent reliably).
type RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState string

const (
    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_closed RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "closed"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_listen RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "listen"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_synSent RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "synSent"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_synReceived RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "synReceived"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_established RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "established"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_finWait1 RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "finWait1"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_finWait2 RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "finWait2"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_closeWait RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "closeWait"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_lastAck RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "lastAck"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_closing RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "closing"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_timeWait RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "timeWait"

    RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState_deleteTCB RFC1213MIB_TcpConnTable_TcpConnEntry_TcpConnState = "deleteTCB"
)

// RFC1213MIB_UdpTable
// A table containing UDP listener information.
type RFC1213MIB_UdpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current UDP listener. The type is slice of
    // RFC1213MIB_UdpTable_UdpEntry.
    UdpEntry []*RFC1213MIB_UdpTable_UdpEntry
}

func (udpTable *RFC1213MIB_UdpTable) GetEntityData() *types.CommonEntityData {
    udpTable.EntityData.YFilter = udpTable.YFilter
    udpTable.EntityData.YangName = "udpTable"
    udpTable.EntityData.BundleName = "cisco_ios_xe"
    udpTable.EntityData.ParentYangName = "RFC1213-MIB"
    udpTable.EntityData.SegmentPath = "udpTable"
    udpTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + udpTable.EntityData.SegmentPath
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

// RFC1213MIB_UdpTable_UdpEntry
// Information about a particular current UDP
// listener.
type RFC1213MIB_UdpTable_UdpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The local IP address for this UDP listener.  In
    // the case of a UDP listener which is willing to accept datagrams for any IP
    // interface associated with the node, the value 0.0.0.0 is used. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    UdpLocalAddress interface{}

    // This attribute is a key. The local port number for this UDP listener. The
    // type is interface{} with range: 0..65535.
    UdpLocalPort interface{}
}

func (udpEntry *RFC1213MIB_UdpTable_UdpEntry) GetEntityData() *types.CommonEntityData {
    udpEntry.EntityData.YFilter = udpEntry.YFilter
    udpEntry.EntityData.YangName = "udpEntry"
    udpEntry.EntityData.BundleName = "cisco_ios_xe"
    udpEntry.EntityData.ParentYangName = "udpTable"
    udpEntry.EntityData.SegmentPath = "udpEntry" + types.AddKeyToken(udpEntry.UdpLocalAddress, "udpLocalAddress") + types.AddKeyToken(udpEntry.UdpLocalPort, "udpLocalPort")
    udpEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/udpTable/" + udpEntry.EntityData.SegmentPath
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

// RFC1213MIB_EgpNeighTable
// The EGP neighbor table.
type RFC1213MIB_EgpNeighTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about this entity's relationship with a particular EGP
    // neighbor. The type is slice of RFC1213MIB_EgpNeighTable_EgpNeighEntry.
    EgpNeighEntry []*RFC1213MIB_EgpNeighTable_EgpNeighEntry
}

func (egpNeighTable *RFC1213MIB_EgpNeighTable) GetEntityData() *types.CommonEntityData {
    egpNeighTable.EntityData.YFilter = egpNeighTable.YFilter
    egpNeighTable.EntityData.YangName = "egpNeighTable"
    egpNeighTable.EntityData.BundleName = "cisco_ios_xe"
    egpNeighTable.EntityData.ParentYangName = "RFC1213-MIB"
    egpNeighTable.EntityData.SegmentPath = "egpNeighTable"
    egpNeighTable.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/" + egpNeighTable.EntityData.SegmentPath
    egpNeighTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egpNeighTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egpNeighTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egpNeighTable.EntityData.Children = types.NewOrderedMap()
    egpNeighTable.EntityData.Children.Append("egpNeighEntry", types.YChild{"EgpNeighEntry", nil})
    for i := range egpNeighTable.EgpNeighEntry {
        egpNeighTable.EntityData.Children.Append(types.GetSegmentPath(egpNeighTable.EgpNeighEntry[i]), types.YChild{"EgpNeighEntry", egpNeighTable.EgpNeighEntry[i]})
    }
    egpNeighTable.EntityData.Leafs = types.NewOrderedMap()

    egpNeighTable.EntityData.YListKeys = []string {}

    return &(egpNeighTable.EntityData)
}

// RFC1213MIB_EgpNeighTable_EgpNeighEntry
// Information about this entity's relationship with
// a particular EGP neighbor.
type RFC1213MIB_EgpNeighTable_EgpNeighEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP address of this entry's EGP neighbor. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    EgpNeighAddr interface{}

    // The EGP state of the local system with respect to this entry's EGP
    // neighbor.  Each EGP state is represented by a value that is one greater
    // than the numerical value associated with said state in RFC 904. The type is
    // EgpNeighState.
    EgpNeighState interface{}

    // The autonomous system of this EGP peer.  Zero should be specified if the
    // autonomous system number of the neighbor is not yet known. The type is
    // interface{} with range: -2147483648..2147483647.
    EgpNeighAs interface{}

    // The number of EGP messages received without error from this EGP peer. The
    // type is interface{} with range: 0..4294967295.
    EgpNeighInMsgs interface{}

    // The number of EGP messages received from this EGP peer that proved to be in
    // error (e.g., bad EGP checksum). The type is interface{} with range:
    // 0..4294967295.
    EgpNeighInErrs interface{}

    // The number of locally generated EGP messages to this EGP peer. The type is
    // interface{} with range: 0..4294967295.
    EgpNeighOutMsgs interface{}

    // The number of locally generated EGP messages not sent to this EGP peer due
    // to resource limitations within an EGP entity. The type is interface{} with
    // range: 0..4294967295.
    EgpNeighOutErrs interface{}

    // The number of EGP-defined error messages received from this EGP peer. The
    // type is interface{} with range: 0..4294967295.
    EgpNeighInErrMsgs interface{}

    // The number of EGP-defined error messages sent to this EGP peer. The type is
    // interface{} with range: 0..4294967295.
    EgpNeighOutErrMsgs interface{}

    // The number of EGP state transitions to the UP state with this EGP peer. The
    // type is interface{} with range: 0..4294967295.
    EgpNeighStateUps interface{}

    // The number of EGP state transitions from the UP state to any other state
    // with this EGP peer. The type is interface{} with range: 0..4294967295.
    EgpNeighStateDowns interface{}

    // The interval between EGP Hello command retransmissions (in hundredths of a
    // second).  This represents the t1 timer as defined in RFC 904. The type is
    // interface{} with range: -2147483648..2147483647.
    EgpNeighIntervalHello interface{}

    // The interval between EGP poll command retransmissions (in hundredths of a
    // second).  This represents the t3 timer as defined in RFC 904. The type is
    // interface{} with range: -2147483648..2147483647.
    EgpNeighIntervalPoll interface{}

    // The polling mode of this EGP entity, either passive or active. The type is
    // EgpNeighMode.
    EgpNeighMode interface{}

    // A control variable used to trigger operator- initiated Start and Stop
    // events.  When read, this variable always returns the most recent value that
    // egpNeighEventTrigger was set to.  If it has not been set since the last
    // initialization of the network management subsystem on the node, it returns
    // a value of `stop'.  When set, this variable causes a Start or Stop event on
    // the specified neighbor, as specified on pages 8-10 of RFC 904.  Briefly, a
    // Start event causes an Idle peer to begin neighbor acquisition and a
    // non-Idle peer to reinitiate neighbor acquisition.  A stop event causes a
    // non-Idle peer to return to the Idle state until a Start event occurs,
    // either via egpNeighEventTrigger or otherwise. The type is
    // EgpNeighEventTrigger.
    EgpNeighEventTrigger interface{}
}

func (egpNeighEntry *RFC1213MIB_EgpNeighTable_EgpNeighEntry) GetEntityData() *types.CommonEntityData {
    egpNeighEntry.EntityData.YFilter = egpNeighEntry.YFilter
    egpNeighEntry.EntityData.YangName = "egpNeighEntry"
    egpNeighEntry.EntityData.BundleName = "cisco_ios_xe"
    egpNeighEntry.EntityData.ParentYangName = "egpNeighTable"
    egpNeighEntry.EntityData.SegmentPath = "egpNeighEntry" + types.AddKeyToken(egpNeighEntry.EgpNeighAddr, "egpNeighAddr")
    egpNeighEntry.EntityData.AbsolutePath = "RFC1213-MIB:RFC1213-MIB/egpNeighTable/" + egpNeighEntry.EntityData.SegmentPath
    egpNeighEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egpNeighEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egpNeighEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egpNeighEntry.EntityData.Children = types.NewOrderedMap()
    egpNeighEntry.EntityData.Leafs = types.NewOrderedMap()
    egpNeighEntry.EntityData.Leafs.Append("egpNeighAddr", types.YLeaf{"EgpNeighAddr", egpNeighEntry.EgpNeighAddr})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighState", types.YLeaf{"EgpNeighState", egpNeighEntry.EgpNeighState})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighAs", types.YLeaf{"EgpNeighAs", egpNeighEntry.EgpNeighAs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighInMsgs", types.YLeaf{"EgpNeighInMsgs", egpNeighEntry.EgpNeighInMsgs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighInErrs", types.YLeaf{"EgpNeighInErrs", egpNeighEntry.EgpNeighInErrs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighOutMsgs", types.YLeaf{"EgpNeighOutMsgs", egpNeighEntry.EgpNeighOutMsgs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighOutErrs", types.YLeaf{"EgpNeighOutErrs", egpNeighEntry.EgpNeighOutErrs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighInErrMsgs", types.YLeaf{"EgpNeighInErrMsgs", egpNeighEntry.EgpNeighInErrMsgs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighOutErrMsgs", types.YLeaf{"EgpNeighOutErrMsgs", egpNeighEntry.EgpNeighOutErrMsgs})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighStateUps", types.YLeaf{"EgpNeighStateUps", egpNeighEntry.EgpNeighStateUps})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighStateDowns", types.YLeaf{"EgpNeighStateDowns", egpNeighEntry.EgpNeighStateDowns})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighIntervalHello", types.YLeaf{"EgpNeighIntervalHello", egpNeighEntry.EgpNeighIntervalHello})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighIntervalPoll", types.YLeaf{"EgpNeighIntervalPoll", egpNeighEntry.EgpNeighIntervalPoll})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighMode", types.YLeaf{"EgpNeighMode", egpNeighEntry.EgpNeighMode})
    egpNeighEntry.EntityData.Leafs.Append("egpNeighEventTrigger", types.YLeaf{"EgpNeighEventTrigger", egpNeighEntry.EgpNeighEventTrigger})

    egpNeighEntry.EntityData.YListKeys = []string {"EgpNeighAddr"}

    return &(egpNeighEntry.EntityData)
}

// RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighEventTrigger represents otherwise.
type RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighEventTrigger string

const (
    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighEventTrigger_start RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighEventTrigger = "start"

    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighEventTrigger_stop RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighEventTrigger = "stop"
)

// RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighMode represents passive or active.
type RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighMode string

const (
    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighMode_active RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighMode = "active"

    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighMode_passive RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighMode = "passive"
)

// RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState represents RFC 904.
type RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState string

const (
    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState_idle RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState = "idle"

    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState_acquisition RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState = "acquisition"

    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState_down RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState = "down"

    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState_up RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState = "up"

    RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState_cease RFC1213MIB_EgpNeighTable_EgpNeighEntry_EgpNeighState = "cease"
)

