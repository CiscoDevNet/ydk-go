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
    Iftable RFC1213MIB_Iftable

    // The Address Translation tables contain the NetworkAddress to `physical'
    // address equivalences. Some interfaces do not use translation tables for
    // determining address equivalences (e.g., DDN-X.25 has an algorithmic
    // method); if all interfaces are of this type, then the Address Translation
    // table is empty, i.e., has zero entries.
    Attable RFC1213MIB_Attable

    // The table of addressing information relevant to this entity's IP addresses.
    Ipaddrtable RFC1213MIB_Ipaddrtable

    // This entity's IP Routing table.
    Iproutetable RFC1213MIB_Iproutetable

    // The IP Address Translation table used for mapping from IP addresses to
    // physical addresses.
    Ipnettomediatable RFC1213MIB_Ipnettomediatable

    // A table containing TCP connection-specific information.
    Tcpconntable RFC1213MIB_Tcpconntable

    // A table containing UDP listener information.
    Udptable RFC1213MIB_Udptable

    // The EGP neighbor table.
    Egpneightable RFC1213MIB_Egpneightable
}

func (rFC1213MIB *RFC1213MIB) GetEntityData() *types.CommonEntityData {
    rFC1213MIB.EntityData.YFilter = rFC1213MIB.YFilter
    rFC1213MIB.EntityData.YangName = "RFC1213-MIB"
    rFC1213MIB.EntityData.BundleName = "cisco_ios_xe"
    rFC1213MIB.EntityData.ParentYangName = "RFC1213-MIB"
    rFC1213MIB.EntityData.SegmentPath = "RFC1213-MIB:RFC1213-MIB"
    rFC1213MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rFC1213MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rFC1213MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rFC1213MIB.EntityData.Children = make(map[string]types.YChild)
    rFC1213MIB.EntityData.Children["system"] = types.YChild{"System", &rFC1213MIB.System}
    rFC1213MIB.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &rFC1213MIB.Interfaces}
    rFC1213MIB.EntityData.Children["ip"] = types.YChild{"Ip", &rFC1213MIB.Ip}
    rFC1213MIB.EntityData.Children["icmp"] = types.YChild{"Icmp", &rFC1213MIB.Icmp}
    rFC1213MIB.EntityData.Children["tcp"] = types.YChild{"Tcp", &rFC1213MIB.Tcp}
    rFC1213MIB.EntityData.Children["udp"] = types.YChild{"Udp", &rFC1213MIB.Udp}
    rFC1213MIB.EntityData.Children["egp"] = types.YChild{"Egp", &rFC1213MIB.Egp}
    rFC1213MIB.EntityData.Children["snmp"] = types.YChild{"Snmp", &rFC1213MIB.Snmp}
    rFC1213MIB.EntityData.Children["ifTable"] = types.YChild{"Iftable", &rFC1213MIB.Iftable}
    rFC1213MIB.EntityData.Children["atTable"] = types.YChild{"Attable", &rFC1213MIB.Attable}
    rFC1213MIB.EntityData.Children["ipAddrTable"] = types.YChild{"Ipaddrtable", &rFC1213MIB.Ipaddrtable}
    rFC1213MIB.EntityData.Children["ipRouteTable"] = types.YChild{"Iproutetable", &rFC1213MIB.Iproutetable}
    rFC1213MIB.EntityData.Children["ipNetToMediaTable"] = types.YChild{"Ipnettomediatable", &rFC1213MIB.Ipnettomediatable}
    rFC1213MIB.EntityData.Children["tcpConnTable"] = types.YChild{"Tcpconntable", &rFC1213MIB.Tcpconntable}
    rFC1213MIB.EntityData.Children["udpTable"] = types.YChild{"Udptable", &rFC1213MIB.Udptable}
    rFC1213MIB.EntityData.Children["egpNeighTable"] = types.YChild{"Egpneightable", &rFC1213MIB.Egpneightable}
    rFC1213MIB.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Sysdescr interface{}

    // The vendor's authoritative identification of the network management
    // subsystem contained in the entity.  This value is allocated within the SMI
    // enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous
    // means for determining `what kind of box' is being managed.  For example, if
    // vendor `Flintstones, Inc.' was assigned the subtree 1.3.6.1.4.1.4242, it
    // could assign the identifier 1.3.6.1.4.1.4242.1.1 to its `Fred Router'. The
    // type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Sysobjectid interface{}

    // The time (in hundredths of a second) since the network management portion
    // of the system was last re-initialized. The type is interface{} with range:
    // 0..4294967295.
    Sysuptime interface{}

    // The textual identification of the contact person for this managed node,
    // together with information on how to contact this person. The type is string
    // with length: 0..255.
    Syscontact interface{}

    // An administratively-assigned name for this managed node.  By convention,
    // this is the node's fully-qualified domain name. The type is string with
    // length: 0..255.
    Sysname interface{}

    // The physical location of this node (e.g., `telephone closet, 3rd floor').
    // The type is string with length: 0..255.
    Syslocation interface{}

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
    Sysservices interface{}
}

func (system *RFC1213MIB_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xe"
    system.EntityData.ParentYangName = "RFC1213-MIB"
    system.EntityData.SegmentPath = "system"
    system.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    system.EntityData.Children = make(map[string]types.YChild)
    system.EntityData.Leafs = make(map[string]types.YLeaf)
    system.EntityData.Leafs["sysDescr"] = types.YLeaf{"Sysdescr", system.Sysdescr}
    system.EntityData.Leafs["sysObjectID"] = types.YLeaf{"Sysobjectid", system.Sysobjectid}
    system.EntityData.Leafs["sysUpTime"] = types.YLeaf{"Sysuptime", system.Sysuptime}
    system.EntityData.Leafs["sysContact"] = types.YLeaf{"Syscontact", system.Syscontact}
    system.EntityData.Leafs["sysName"] = types.YLeaf{"Sysname", system.Sysname}
    system.EntityData.Leafs["sysLocation"] = types.YLeaf{"Syslocation", system.Syslocation}
    system.EntityData.Leafs["sysServices"] = types.YLeaf{"Sysservices", system.Sysservices}
    return &(system.EntityData)
}

// RFC1213MIB_Interfaces
type RFC1213MIB_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of network interfaces (regardless of their current state)
    // present on this system. The type is interface{} with range:
    // -2147483648..2147483647.
    Ifnumber interface{}
}

func (interfaces *RFC1213MIB_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "RFC1213-MIB"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["ifNumber"] = types.YLeaf{"Ifnumber", interfaces.Ifnumber}
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
    // is Ipforwarding.
    Ipforwarding interface{}

    // The default value inserted into the Time-To-Live field of the IP header of
    // datagrams originated at this entity, whenever a TTL value is not supplied
    // by the transport layer protocol. The type is interface{} with range:
    // -2147483648..2147483647.
    Ipdefaultttl interface{}

    // The total number of input datagrams received from interfaces, including
    // those received in error. The type is interface{} with range: 0..4294967295.
    Ipinreceives interface{}

    // The number of input datagrams discarded due to errors in their IP headers,
    // including bad checksums, version number mismatch, other format errors,
    // time-to-live exceeded, errors discovered in processing their IP options,
    // etc. The type is interface{} with range: 0..4294967295.
    Ipinhdrerrors interface{}

    // The number of input datagrams discarded because the IP address in their IP
    // header's destination field was not a valid address to be received at this
    // entity.  This count includes invalid addresses (e.g., 0.0.0.0) and
    // addresses of unsupported Classes (e.g., Class E).  For entities which are
    // not IP Gateways and therefore do not forward datagrams, this counter
    // includes datagrams discarded because the destination address was not a
    // local address. The type is interface{} with range: 0..4294967295.
    Ipinaddrerrors interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination, as a result of which an attempt was made to find a route to
    // forward them to that final destination. In entities which do not act as IP
    // Gateways, this counter will include only those packets which were
    // Source-Routed via this entity, and the Source- Route option processing was
    // successful. The type is interface{} with range: 0..4294967295.
    Ipforwdatagrams interface{}

    // The number of locally-addressed datagrams received successfully but
    // discarded because of an unknown or unsupported protocol. The type is
    // interface{} with range: 0..4294967295.
    Ipinunknownprotos interface{}

    // The number of input IP datagrams for which no problems were encountered to
    // prevent their continued processing, but which were discarded (e.g., for
    // lack of buffer space).  Note that this counter does not include any
    // datagrams discarded while awaiting re-assembly. The type is interface{}
    // with range: 0..4294967295.
    Ipindiscards interface{}

    // The total number of input datagrams successfully delivered to IP
    // user-protocols (including ICMP). The type is interface{} with range:
    // 0..4294967295.
    Ipindelivers interface{}

    // The total number of IP datagrams which local IP user-protocols (including
    // ICMP) supplied to IP in requests for transmission.  Note that this counter
    // does not include any datagrams counted in ipForwDatagrams. The type is
    // interface{} with range: 0..4294967295.
    Ipoutrequests interface{}

    // The number of output IP datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but which were discarded
    // (e.g., for lack of buffer space).  Note that this counter would include
    // datagrams counted in ipForwDatagrams if any such packets met this
    // (discretionary) discard criterion. The type is interface{} with range:
    // 0..4294967295.
    Ipoutdiscards interface{}

    // The number of IP datagrams discarded because no route could be found to
    // transmit them to their destination.  Note that this counter includes any
    // packets counted in ipForwDatagrams which meet this `no-route' criterion. 
    // Note that this includes any datagrams which a host cannot route because all
    // of its default gateways are down. The type is interface{} with range:
    // 0..4294967295.
    Ipoutnoroutes interface{}

    // The maximum number of seconds which received fragments are held while they
    // are awaiting reassembly at this entity. The type is interface{} with range:
    // -2147483648..2147483647.
    Ipreasmtimeout interface{}

    // The number of IP fragments received which needed to be reassembled at this
    // entity. The type is interface{} with range: 0..4294967295.
    Ipreasmreqds interface{}

    // The number of IP datagrams successfully re- assembled. The type is
    // interface{} with range: 0..4294967295.
    Ipreasmoks interface{}

    // The number of failures detected by the IP re- assembly algorithm (for
    // whatever reason: timed out, errors, etc).  Note that this is not
    // necessarily a count of discarded IP fragments since some algorithms
    // (notably the algorithm in RFC 815) can lose track of the number of
    // fragments by combining them as they are received. The type is interface{}
    // with range: 0..4294967295.
    Ipreasmfails interface{}

    // The number of IP datagrams that have been successfully fragmented at this
    // entity. The type is interface{} with range: 0..4294967295.
    Ipfragoks interface{}

    // The number of IP datagrams that have been discarded because they needed to
    // be fragmented at this entity but could not be, e.g., because their Don't
    // Fragment flag was set. The type is interface{} with range: 0..4294967295.
    Ipfragfails interface{}

    // The number of IP datagram fragments that have been generated as a result of
    // fragmentation at this entity. The type is interface{} with range:
    // 0..4294967295.
    Ipfragcreates interface{}

    // The number of routing entries which were chosen to be discarded even though
    // they are valid.  One possible reason for discarding such an entry could be
    // to free-up buffer space for other routing entries. The type is interface{}
    // with range: 0..4294967295.
    Iproutingdiscards interface{}
}

func (ip *RFC1213MIB_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xe"
    ip.EntityData.ParentYangName = "RFC1213-MIB"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    ip.EntityData.Leafs["ipForwarding"] = types.YLeaf{"Ipforwarding", ip.Ipforwarding}
    ip.EntityData.Leafs["ipDefaultTTL"] = types.YLeaf{"Ipdefaultttl", ip.Ipdefaultttl}
    ip.EntityData.Leafs["ipInReceives"] = types.YLeaf{"Ipinreceives", ip.Ipinreceives}
    ip.EntityData.Leafs["ipInHdrErrors"] = types.YLeaf{"Ipinhdrerrors", ip.Ipinhdrerrors}
    ip.EntityData.Leafs["ipInAddrErrors"] = types.YLeaf{"Ipinaddrerrors", ip.Ipinaddrerrors}
    ip.EntityData.Leafs["ipForwDatagrams"] = types.YLeaf{"Ipforwdatagrams", ip.Ipforwdatagrams}
    ip.EntityData.Leafs["ipInUnknownProtos"] = types.YLeaf{"Ipinunknownprotos", ip.Ipinunknownprotos}
    ip.EntityData.Leafs["ipInDiscards"] = types.YLeaf{"Ipindiscards", ip.Ipindiscards}
    ip.EntityData.Leafs["ipInDelivers"] = types.YLeaf{"Ipindelivers", ip.Ipindelivers}
    ip.EntityData.Leafs["ipOutRequests"] = types.YLeaf{"Ipoutrequests", ip.Ipoutrequests}
    ip.EntityData.Leafs["ipOutDiscards"] = types.YLeaf{"Ipoutdiscards", ip.Ipoutdiscards}
    ip.EntityData.Leafs["ipOutNoRoutes"] = types.YLeaf{"Ipoutnoroutes", ip.Ipoutnoroutes}
    ip.EntityData.Leafs["ipReasmTimeout"] = types.YLeaf{"Ipreasmtimeout", ip.Ipreasmtimeout}
    ip.EntityData.Leafs["ipReasmReqds"] = types.YLeaf{"Ipreasmreqds", ip.Ipreasmreqds}
    ip.EntityData.Leafs["ipReasmOKs"] = types.YLeaf{"Ipreasmoks", ip.Ipreasmoks}
    ip.EntityData.Leafs["ipReasmFails"] = types.YLeaf{"Ipreasmfails", ip.Ipreasmfails}
    ip.EntityData.Leafs["ipFragOKs"] = types.YLeaf{"Ipfragoks", ip.Ipfragoks}
    ip.EntityData.Leafs["ipFragFails"] = types.YLeaf{"Ipfragfails", ip.Ipfragfails}
    ip.EntityData.Leafs["ipFragCreates"] = types.YLeaf{"Ipfragcreates", ip.Ipfragcreates}
    ip.EntityData.Leafs["ipRoutingDiscards"] = types.YLeaf{"Iproutingdiscards", ip.Iproutingdiscards}
    return &(ip.EntityData)
}

// RFC1213MIB_Ip_Ipforwarding represents inappropriate value.
type RFC1213MIB_Ip_Ipforwarding string

const (
    RFC1213MIB_Ip_Ipforwarding_forwarding RFC1213MIB_Ip_Ipforwarding = "forwarding"

    RFC1213MIB_Ip_Ipforwarding_not_forwarding RFC1213MIB_Ip_Ipforwarding = "not-forwarding"
)

// RFC1213MIB_Icmp
type RFC1213MIB_Icmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of ICMP messages which the entity received.  Note that
    // this counter includes all those counted by icmpInErrors. The type is
    // interface{} with range: 0..4294967295.
    Icmpinmsgs interface{}

    // The number of ICMP messages which the entity received but determined as
    // having ICMP-specific errors (bad ICMP checksums, bad length, etc.). The
    // type is interface{} with range: 0..4294967295.
    Icmpinerrors interface{}

    // The number of ICMP Destination Unreachable messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpindestunreachs interface{}

    // The number of ICMP Time Exceeded messages received. The type is interface{}
    // with range: 0..4294967295.
    Icmpintimeexcds interface{}

    // The number of ICMP Parameter Problem messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpinparmprobs interface{}

    // The number of ICMP Source Quench messages received. The type is interface{}
    // with range: 0..4294967295.
    Icmpinsrcquenchs interface{}

    // The number of ICMP Redirect messages received. The type is interface{} with
    // range: 0..4294967295.
    Icmpinredirects interface{}

    // The number of ICMP Echo (request) messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpinechos interface{}

    // The number of ICMP Echo Reply messages received. The type is interface{}
    // with range: 0..4294967295.
    Icmpinechoreps interface{}

    // The number of ICMP Timestamp (request) messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpintimestamps interface{}

    // The number of ICMP Timestamp Reply messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpintimestampreps interface{}

    // The number of ICMP Address Mask Request messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpinaddrmasks interface{}

    // The number of ICMP Address Mask Reply messages received. The type is
    // interface{} with range: 0..4294967295.
    Icmpinaddrmaskreps interface{}

    // The total number of ICMP messages which this entity attempted to send. 
    // Note that this counter includes all those counted by icmpOutErrors. The
    // type is interface{} with range: 0..4294967295.
    Icmpoutmsgs interface{}

    // The number of ICMP messages which this entity did not send due to problems
    // discovered within ICMP such as a lack of buffers.  This value should not
    // include errors discovered outside the ICMP layer such as the inability of
    // IP to route the resultant datagram.  In some implementations there may be
    // no types of error which contribute to this counter's value. The type is
    // interface{} with range: 0..4294967295.
    Icmpouterrors interface{}

    // The number of ICMP Destination Unreachable messages sent. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutdestunreachs interface{}

    // The number of ICMP Time Exceeded messages sent. The type is interface{}
    // with range: 0..4294967295.
    Icmpouttimeexcds interface{}

    // The number of ICMP Parameter Problem messages sent. The type is interface{}
    // with range: 0..4294967295.
    Icmpoutparmprobs interface{}

    // The number of ICMP Source Quench messages sent. The type is interface{}
    // with range: 0..4294967295.
    Icmpoutsrcquenchs interface{}

    // The number of ICMP Redirect messages sent.  For a host, this object will
    // always be zero, since hosts do not send redirects. The type is interface{}
    // with range: 0..4294967295.
    Icmpoutredirects interface{}

    // The number of ICMP Echo (request) messages sent. The type is interface{}
    // with range: 0..4294967295.
    Icmpoutechos interface{}

    // The number of ICMP Echo Reply messages sent. The type is interface{} with
    // range: 0..4294967295.
    Icmpoutechoreps interface{}

    // The number of ICMP Timestamp (request) messages sent. The type is
    // interface{} with range: 0..4294967295.
    Icmpouttimestamps interface{}

    // The number of ICMP Timestamp Reply messages sent. The type is interface{}
    // with range: 0..4294967295.
    Icmpouttimestampreps interface{}

    // The number of ICMP Address Mask Request messages sent. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutaddrmasks interface{}

    // The number of ICMP Address Mask Reply messages sent. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutaddrmaskreps interface{}
}

func (icmp *RFC1213MIB_Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xe"
    icmp.EntityData.ParentYangName = "RFC1213-MIB"
    icmp.EntityData.SegmentPath = "icmp"
    icmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmp.EntityData.Children = make(map[string]types.YChild)
    icmp.EntityData.Leafs = make(map[string]types.YLeaf)
    icmp.EntityData.Leafs["icmpInMsgs"] = types.YLeaf{"Icmpinmsgs", icmp.Icmpinmsgs}
    icmp.EntityData.Leafs["icmpInErrors"] = types.YLeaf{"Icmpinerrors", icmp.Icmpinerrors}
    icmp.EntityData.Leafs["icmpInDestUnreachs"] = types.YLeaf{"Icmpindestunreachs", icmp.Icmpindestunreachs}
    icmp.EntityData.Leafs["icmpInTimeExcds"] = types.YLeaf{"Icmpintimeexcds", icmp.Icmpintimeexcds}
    icmp.EntityData.Leafs["icmpInParmProbs"] = types.YLeaf{"Icmpinparmprobs", icmp.Icmpinparmprobs}
    icmp.EntityData.Leafs["icmpInSrcQuenchs"] = types.YLeaf{"Icmpinsrcquenchs", icmp.Icmpinsrcquenchs}
    icmp.EntityData.Leafs["icmpInRedirects"] = types.YLeaf{"Icmpinredirects", icmp.Icmpinredirects}
    icmp.EntityData.Leafs["icmpInEchos"] = types.YLeaf{"Icmpinechos", icmp.Icmpinechos}
    icmp.EntityData.Leafs["icmpInEchoReps"] = types.YLeaf{"Icmpinechoreps", icmp.Icmpinechoreps}
    icmp.EntityData.Leafs["icmpInTimestamps"] = types.YLeaf{"Icmpintimestamps", icmp.Icmpintimestamps}
    icmp.EntityData.Leafs["icmpInTimestampReps"] = types.YLeaf{"Icmpintimestampreps", icmp.Icmpintimestampreps}
    icmp.EntityData.Leafs["icmpInAddrMasks"] = types.YLeaf{"Icmpinaddrmasks", icmp.Icmpinaddrmasks}
    icmp.EntityData.Leafs["icmpInAddrMaskReps"] = types.YLeaf{"Icmpinaddrmaskreps", icmp.Icmpinaddrmaskreps}
    icmp.EntityData.Leafs["icmpOutMsgs"] = types.YLeaf{"Icmpoutmsgs", icmp.Icmpoutmsgs}
    icmp.EntityData.Leafs["icmpOutErrors"] = types.YLeaf{"Icmpouterrors", icmp.Icmpouterrors}
    icmp.EntityData.Leafs["icmpOutDestUnreachs"] = types.YLeaf{"Icmpoutdestunreachs", icmp.Icmpoutdestunreachs}
    icmp.EntityData.Leafs["icmpOutTimeExcds"] = types.YLeaf{"Icmpouttimeexcds", icmp.Icmpouttimeexcds}
    icmp.EntityData.Leafs["icmpOutParmProbs"] = types.YLeaf{"Icmpoutparmprobs", icmp.Icmpoutparmprobs}
    icmp.EntityData.Leafs["icmpOutSrcQuenchs"] = types.YLeaf{"Icmpoutsrcquenchs", icmp.Icmpoutsrcquenchs}
    icmp.EntityData.Leafs["icmpOutRedirects"] = types.YLeaf{"Icmpoutredirects", icmp.Icmpoutredirects}
    icmp.EntityData.Leafs["icmpOutEchos"] = types.YLeaf{"Icmpoutechos", icmp.Icmpoutechos}
    icmp.EntityData.Leafs["icmpOutEchoReps"] = types.YLeaf{"Icmpoutechoreps", icmp.Icmpoutechoreps}
    icmp.EntityData.Leafs["icmpOutTimestamps"] = types.YLeaf{"Icmpouttimestamps", icmp.Icmpouttimestamps}
    icmp.EntityData.Leafs["icmpOutTimestampReps"] = types.YLeaf{"Icmpouttimestampreps", icmp.Icmpouttimestampreps}
    icmp.EntityData.Leafs["icmpOutAddrMasks"] = types.YLeaf{"Icmpoutaddrmasks", icmp.Icmpoutaddrmasks}
    icmp.EntityData.Leafs["icmpOutAddrMaskReps"] = types.YLeaf{"Icmpoutaddrmaskreps", icmp.Icmpoutaddrmaskreps}
    return &(icmp.EntityData)
}

// RFC1213MIB_Tcp
type RFC1213MIB_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The algorithm used to determine the timeout value used for retransmitting
    // unacknowledged octets. The type is Tcprtoalgorithm.
    Tcprtoalgorithm interface{}

    // The minimum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds.  More refined semantics for objects of
    // this type depend upon the algorithm used to determine the retransmission
    // timeout.  In particular, when the timeout algorithm is rsre(3), an object
    // of this type has the semantics of the LBOUND quantity described in RFC 793.
    // The type is interface{} with range: -2147483648..2147483647.
    Tcprtomin interface{}

    // The maximum value permitted by a TCP implementation for the retransmission
    // timeout, measured in milliseconds.  More refined semantics for objects of
    // this type depend upon the algorithm used to determine the retransmission
    // timeout.  In particular, when the timeout algorithm is rsre(3), an object
    // of this type has the semantics of the UBOUND quantity described in RFC 793.
    // The type is interface{} with range: -2147483648..2147483647.
    Tcprtomax interface{}

    // The limit on the total number of TCP connections the entity can support. 
    // In entities where the maximum number of connections is dynamic, this object
    // should contain the value -1. The type is interface{} with range:
    // -2147483648..2147483647.
    Tcpmaxconn interface{}

    // The number of times TCP connections have made a direct transition to the
    // SYN-SENT state from the CLOSED state. The type is interface{} with range:
    // 0..4294967295.
    Tcpactiveopens interface{}

    // The number of times TCP connections have made a direct transition to the
    // SYN-RCVD state from the LISTEN state. The type is interface{} with range:
    // 0..4294967295.
    Tcppassiveopens interface{}

    // The number of times TCP connections have made a direct transition to the
    // CLOSED state from either the SYN-SENT state or the SYN-RCVD state, plus the
    // number of times TCP connections have made a direct transition to the LISTEN
    // state from the SYN-RCVD state. The type is interface{} with range:
    // 0..4294967295.
    Tcpattemptfails interface{}

    // The number of times TCP connections have made a direct transition to the
    // CLOSED state from either the ESTABLISHED state or the CLOSE-WAIT state. The
    // type is interface{} with range: 0..4294967295.
    Tcpestabresets interface{}

    // The number of TCP connections for which the current state is either
    // ESTABLISHED or CLOSE- WAIT. The type is interface{} with range:
    // 0..4294967295.
    Tcpcurrestab interface{}

    // The total number of segments received, including those received in error. 
    // This count includes segments received on currently established connections.
    // The type is interface{} with range: 0..4294967295.
    Tcpinsegs interface{}

    // The total number of segments sent, including those on current connections
    // but excluding those containing only retransmitted octets. The type is
    // interface{} with range: 0..4294967295.
    Tcpoutsegs interface{}

    // The total number of segments retransmitted - that is, the number of TCP
    // segments transmitted containing one or more previously transmitted octets.
    // The type is interface{} with range: 0..4294967295.
    Tcpretranssegs interface{}

    // The total number of segments received in error (e.g., bad TCP checksums).
    // The type is interface{} with range: 0..4294967295.
    Tcpinerrs interface{}

    // The number of TCP segments sent containing the RST flag. The type is
    // interface{} with range: 0..4294967295.
    Tcpoutrsts interface{}
}

func (tcp *RFC1213MIB_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "RFC1213-MIB"
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
    return &(tcp.EntityData)
}

// RFC1213MIB_Tcp_Tcprtoalgorithm represents used for retransmitting unacknowledged octets.
type RFC1213MIB_Tcp_Tcprtoalgorithm string

const (
    RFC1213MIB_Tcp_Tcprtoalgorithm_other RFC1213MIB_Tcp_Tcprtoalgorithm = "other"

    RFC1213MIB_Tcp_Tcprtoalgorithm_constant RFC1213MIB_Tcp_Tcprtoalgorithm = "constant"

    RFC1213MIB_Tcp_Tcprtoalgorithm_rsre RFC1213MIB_Tcp_Tcprtoalgorithm = "rsre"

    RFC1213MIB_Tcp_Tcprtoalgorithm_vanj RFC1213MIB_Tcp_Tcprtoalgorithm = "vanj"

    RFC1213MIB_Tcp_Tcprtoalgorithm_rfc2988 RFC1213MIB_Tcp_Tcprtoalgorithm = "rfc2988"
)

// RFC1213MIB_Udp
type RFC1213MIB_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of UDP datagrams delivered to UDP users. The type is
    // interface{} with range: 0..4294967295.
    Udpindatagrams interface{}

    // The total number of received UDP datagrams for which there was no
    // application at the destination port. The type is interface{} with range:
    // 0..4294967295.
    Udpnoports interface{}

    // The number of received UDP datagrams that could not be delivered for
    // reasons other than the lack of an application at the destination port. The
    // type is interface{} with range: 0..4294967295.
    Udpinerrors interface{}

    // The total number of UDP datagrams sent from this entity. The type is
    // interface{} with range: 0..4294967295.
    Udpoutdatagrams interface{}
}

func (udp *RFC1213MIB_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xe"
    udp.EntityData.ParentYangName = "RFC1213-MIB"
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
    return &(udp.EntityData)
}

// RFC1213MIB_Egp
type RFC1213MIB_Egp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of EGP messages received without error. The type is interface{}
    // with range: 0..4294967295.
    Egpinmsgs interface{}

    // The number of EGP messages received that proved to be in error. The type is
    // interface{} with range: 0..4294967295.
    Egpinerrors interface{}

    // The total number of locally generated EGP messages. The type is interface{}
    // with range: 0..4294967295.
    Egpoutmsgs interface{}

    // The number of locally generated EGP messages not sent due to resource
    // limitations within an EGP entity. The type is interface{} with range:
    // 0..4294967295.
    Egpouterrors interface{}

    // The autonomous system number of this EGP entity. The type is interface{}
    // with range: -2147483648..2147483647.
    Egpas interface{}
}

func (egp *RFC1213MIB_Egp) GetEntityData() *types.CommonEntityData {
    egp.EntityData.YFilter = egp.YFilter
    egp.EntityData.YangName = "egp"
    egp.EntityData.BundleName = "cisco_ios_xe"
    egp.EntityData.ParentYangName = "RFC1213-MIB"
    egp.EntityData.SegmentPath = "egp"
    egp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egp.EntityData.Children = make(map[string]types.YChild)
    egp.EntityData.Leafs = make(map[string]types.YLeaf)
    egp.EntityData.Leafs["egpInMsgs"] = types.YLeaf{"Egpinmsgs", egp.Egpinmsgs}
    egp.EntityData.Leafs["egpInErrors"] = types.YLeaf{"Egpinerrors", egp.Egpinerrors}
    egp.EntityData.Leafs["egpOutMsgs"] = types.YLeaf{"Egpoutmsgs", egp.Egpoutmsgs}
    egp.EntityData.Leafs["egpOutErrors"] = types.YLeaf{"Egpouterrors", egp.Egpouterrors}
    egp.EntityData.Leafs["egpAs"] = types.YLeaf{"Egpas", egp.Egpas}
    return &(egp.EntityData)
}

// RFC1213MIB_Snmp
type RFC1213MIB_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of Messages delivered to the SNMP entity from the
    // transport service. The type is interface{} with range: 0..4294967295.
    Snmpinpkts interface{}

    // The total number of SNMP Messages which were passed from the SNMP protocol
    // entity to the transport service. The type is interface{} with range:
    // 0..4294967295.
    Snmpoutpkts interface{}

    // The total number of SNMP Messages which were delivered to the SNMP protocol
    // entity and were for an unsupported SNMP version. The type is interface{}
    // with range: 0..4294967295.
    Snmpinbadversions interface{}

    // The total number of SNMP Messages delivered to the SNMP protocol entity
    // which used a SNMP community name not known to said entity. The type is
    // interface{} with range: 0..4294967295.
    Snmpinbadcommunitynames interface{}

    // The total number of SNMP Messages delivered to the SNMP protocol entity
    // which represented an SNMP operation which was not allowed by the SNMP
    // community named in the Message. The type is interface{} with range:
    // 0..4294967295.
    Snmpinbadcommunityuses interface{}

    // The total number of ASN.1 or BER errors encountered by the SNMP protocol
    // entity when decoding received SNMP Messages. The type is interface{} with
    // range: 0..4294967295.
    Snmpinasnparseerrs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `tooBig'. The
    // type is interface{} with range: 0..4294967295.
    Snmpintoobigs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `noSuchName'.
    // The type is interface{} with range: 0..4294967295.
    Snmpinnosuchnames interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `badValue'. The
    // type is interface{} with range: 0..4294967295.
    Snmpinbadvalues interface{}

    // The total number valid SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `readOnly'.  It
    // should be noted that it is a protocol error to generate an SNMP PDU which
    // contains the value `readOnly' in the error-status field, as such this
    // object is provided as a means of detecting incorrect implementations of the
    // SNMP. The type is interface{} with range: 0..4294967295.
    Snmpinreadonlys interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field is `genErr'. The
    // type is interface{} with range: 0..4294967295.
    Snmpingenerrs interface{}

    // The total number of MIB objects which have been retrieved successfully by
    // the SNMP protocol entity as the result of receiving valid SNMP Get-Request
    // and Get-Next PDUs. The type is interface{} with range: 0..4294967295.
    Snmpintotalreqvars interface{}

    // The total number of MIB objects which have been altered successfully by the
    // SNMP protocol entity as the result of receiving valid SNMP Set-Request
    // PDUs. The type is interface{} with range: 0..4294967295.
    Snmpintotalsetvars interface{}

    // The total number of SNMP Get-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpingetrequests interface{}

    // The total number of SNMP Get-Next PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpingetnexts interface{}

    // The total number of SNMP Set-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpinsetrequests interface{}

    // The total number of SNMP Get-Response PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpingetresponses interface{}

    // The total number of SNMP Trap PDUs which have been accepted and processed
    // by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpintraps interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field is `tooBig.'. The
    // type is interface{} with range: 0..4294967295.
    Snmpouttoobigs interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status is `noSuchName'. The
    // type is interface{} with range: 0..4294967295.
    Snmpoutnosuchnames interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field is `badValue'. The
    // type is interface{} with range: 0..4294967295.
    Snmpoutbadvalues interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field is `genErr'. The
    // type is interface{} with range: 0..4294967295.
    Snmpoutgenerrs interface{}

    // The total number of SNMP Get-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutgetrequests interface{}

    // The total number of SNMP Get-Next PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutgetnexts interface{}

    // The total number of SNMP Set-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutsetrequests interface{}

    // The total number of SNMP Get-Response PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutgetresponses interface{}

    // The total number of SNMP Trap PDUs which have been generated by the SNMP
    // protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpouttraps interface{}

    // Indicates whether the SNMP agent process is permitted to generate
    // authentication-failure traps.  The value of this object overrides any
    // configuration information; as such, it provides a means whereby all
    // authentication-failure traps may be disabled.  Note that it is strongly
    // recommended that this object be stored in non-volatile memory so that it
    // remains constant between re-initializations of the network management
    // system. The type is Snmpenableauthentraps.
    Snmpenableauthentraps interface{}
}

func (snmp *RFC1213MIB_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xe"
    snmp.EntityData.ParentYangName = "RFC1213-MIB"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmp.EntityData.Children = make(map[string]types.YChild)
    snmp.EntityData.Leafs = make(map[string]types.YLeaf)
    snmp.EntityData.Leafs["snmpInPkts"] = types.YLeaf{"Snmpinpkts", snmp.Snmpinpkts}
    snmp.EntityData.Leafs["snmpOutPkts"] = types.YLeaf{"Snmpoutpkts", snmp.Snmpoutpkts}
    snmp.EntityData.Leafs["snmpInBadVersions"] = types.YLeaf{"Snmpinbadversions", snmp.Snmpinbadversions}
    snmp.EntityData.Leafs["snmpInBadCommunityNames"] = types.YLeaf{"Snmpinbadcommunitynames", snmp.Snmpinbadcommunitynames}
    snmp.EntityData.Leafs["snmpInBadCommunityUses"] = types.YLeaf{"Snmpinbadcommunityuses", snmp.Snmpinbadcommunityuses}
    snmp.EntityData.Leafs["snmpInASNParseErrs"] = types.YLeaf{"Snmpinasnparseerrs", snmp.Snmpinasnparseerrs}
    snmp.EntityData.Leafs["snmpInTooBigs"] = types.YLeaf{"Snmpintoobigs", snmp.Snmpintoobigs}
    snmp.EntityData.Leafs["snmpInNoSuchNames"] = types.YLeaf{"Snmpinnosuchnames", snmp.Snmpinnosuchnames}
    snmp.EntityData.Leafs["snmpInBadValues"] = types.YLeaf{"Snmpinbadvalues", snmp.Snmpinbadvalues}
    snmp.EntityData.Leafs["snmpInReadOnlys"] = types.YLeaf{"Snmpinreadonlys", snmp.Snmpinreadonlys}
    snmp.EntityData.Leafs["snmpInGenErrs"] = types.YLeaf{"Snmpingenerrs", snmp.Snmpingenerrs}
    snmp.EntityData.Leafs["snmpInTotalReqVars"] = types.YLeaf{"Snmpintotalreqvars", snmp.Snmpintotalreqvars}
    snmp.EntityData.Leafs["snmpInTotalSetVars"] = types.YLeaf{"Snmpintotalsetvars", snmp.Snmpintotalsetvars}
    snmp.EntityData.Leafs["snmpInGetRequests"] = types.YLeaf{"Snmpingetrequests", snmp.Snmpingetrequests}
    snmp.EntityData.Leafs["snmpInGetNexts"] = types.YLeaf{"Snmpingetnexts", snmp.Snmpingetnexts}
    snmp.EntityData.Leafs["snmpInSetRequests"] = types.YLeaf{"Snmpinsetrequests", snmp.Snmpinsetrequests}
    snmp.EntityData.Leafs["snmpInGetResponses"] = types.YLeaf{"Snmpingetresponses", snmp.Snmpingetresponses}
    snmp.EntityData.Leafs["snmpInTraps"] = types.YLeaf{"Snmpintraps", snmp.Snmpintraps}
    snmp.EntityData.Leafs["snmpOutTooBigs"] = types.YLeaf{"Snmpouttoobigs", snmp.Snmpouttoobigs}
    snmp.EntityData.Leafs["snmpOutNoSuchNames"] = types.YLeaf{"Snmpoutnosuchnames", snmp.Snmpoutnosuchnames}
    snmp.EntityData.Leafs["snmpOutBadValues"] = types.YLeaf{"Snmpoutbadvalues", snmp.Snmpoutbadvalues}
    snmp.EntityData.Leafs["snmpOutGenErrs"] = types.YLeaf{"Snmpoutgenerrs", snmp.Snmpoutgenerrs}
    snmp.EntityData.Leafs["snmpOutGetRequests"] = types.YLeaf{"Snmpoutgetrequests", snmp.Snmpoutgetrequests}
    snmp.EntityData.Leafs["snmpOutGetNexts"] = types.YLeaf{"Snmpoutgetnexts", snmp.Snmpoutgetnexts}
    snmp.EntityData.Leafs["snmpOutSetRequests"] = types.YLeaf{"Snmpoutsetrequests", snmp.Snmpoutsetrequests}
    snmp.EntityData.Leafs["snmpOutGetResponses"] = types.YLeaf{"Snmpoutgetresponses", snmp.Snmpoutgetresponses}
    snmp.EntityData.Leafs["snmpOutTraps"] = types.YLeaf{"Snmpouttraps", snmp.Snmpouttraps}
    snmp.EntityData.Leafs["snmpEnableAuthenTraps"] = types.YLeaf{"Snmpenableauthentraps", snmp.Snmpenableauthentraps}
    return &(snmp.EntityData)
}

// RFC1213MIB_Snmp_Snmpenableauthentraps represents network management system.
type RFC1213MIB_Snmp_Snmpenableauthentraps string

const (
    RFC1213MIB_Snmp_Snmpenableauthentraps_enabled RFC1213MIB_Snmp_Snmpenableauthentraps = "enabled"

    RFC1213MIB_Snmp_Snmpenableauthentraps_disabled RFC1213MIB_Snmp_Snmpenableauthentraps = "disabled"
)

// RFC1213MIB_Iftable
// A list of interface entries.  The number of
// entries is given by the value of ifNumber.
type RFC1213MIB_Iftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An interface entry containing objects at the subnetwork layer and below for
    // a particular interface. The type is slice of RFC1213MIB_Iftable_Ifentry.
    Ifentry []RFC1213MIB_Iftable_Ifentry
}

func (iftable *RFC1213MIB_Iftable) GetEntityData() *types.CommonEntityData {
    iftable.EntityData.YFilter = iftable.YFilter
    iftable.EntityData.YangName = "ifTable"
    iftable.EntityData.BundleName = "cisco_ios_xe"
    iftable.EntityData.ParentYangName = "RFC1213-MIB"
    iftable.EntityData.SegmentPath = "ifTable"
    iftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iftable.EntityData.Children = make(map[string]types.YChild)
    iftable.EntityData.Children["ifEntry"] = types.YChild{"Ifentry", nil}
    for i := range iftable.Ifentry {
        iftable.EntityData.Children[types.GetSegmentPath(&iftable.Ifentry[i])] = types.YChild{"Ifentry", &iftable.Ifentry[i]}
    }
    iftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iftable.EntityData)
}

// RFC1213MIB_Iftable_Ifentry
// An interface entry containing objects at the
// subnetwork layer and below for a particular
// interface.
type RFC1213MIB_Iftable_Ifentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique value for each interface.  Its value
    // ranges between 1 and the value of ifNumber.  The value for each interface
    // must remain constant at least from one re-initialization of the entity's
    // network management system to the next re- initialization. The type is
    // interface{} with range: -2147483648..2147483647.
    Ifindex interface{}

    // A textual string containing information about the interface.  This string
    // should include the name of the manufacturer, the product name and the
    // version of the hardware interface. The type is string with length: 0..255.
    Ifdescr interface{}

    // The type of interface.  Additional values for ifType are assigned by the
    // Internet Assigned Numbers Authority (IANA), through updating the syntax of
    // the IANAifType textual convention. The type is IANAifType.
    Iftype interface{}

    // The size of the largest datagram which can be sent/received on the
    // interface, specified in octets.  For interfaces that are used for
    // transmitting network datagrams, this is the size of the largest network
    // datagram that can be sent on the interface. The type is interface{} with
    // range: -2147483648..2147483647.
    Ifmtu interface{}

    // An estimate of the interface's current bandwidth in bits per second.  For
    // interfaces which do not vary in bandwidth or for those where no accurate
    // estimation can be made, this object should contain the nominal bandwidth.
    // The type is interface{} with range: 0..4294967295.
    Ifspeed interface{}

    // The interface's address at the protocol layer immediately `below' the
    // network layer in the protocol stack.  For interfaces which do not have such
    // an address (e.g., a serial line), this object should contain an octet
    // string of zero length. The type is string.
    Ifphysaddress interface{}

    // The desired state of the interface.  The testing(3) state indicates that no
    // operational packets can be passed. The type is Ifadminstatus.
    Ifadminstatus interface{}

    // The current operational state of the interface. The testing(3) state
    // indicates that no operational packets can be passed. The type is
    // Ifoperstatus.
    Ifoperstatus interface{}

    // The value of sysUpTime at the time the interface entered its current
    // operational state.  If the current state was entered prior to the last re-
    // initialization of the local network management subsystem, then this object
    // contains a zero value. The type is interface{} with range: 0..4294967295.
    Iflastchange interface{}

    // The total number of octets received on the interface, including framing
    // characters. The type is interface{} with range: 0..4294967295.
    Ifinoctets interface{}

    // The number of subnetwork-unicast packets delivered to a higher-layer
    // protocol. The type is interface{} with range: 0..4294967295.
    Ifinucastpkts interface{}

    // The number of non-unicast (i.e., subnetwork- broadcast or
    // subnetwork-multicast) packets delivered to a higher-layer protocol. The
    // type is interface{} with range: 0..4294967295.
    Ifinnucastpkts interface{}

    // The number of inbound packets which were chosen to be discarded even though
    // no errors had been detected to prevent their being deliverable to a
    // higher-layer protocol.  One possible reason for discarding such a packet
    // could be to free up buffer space. The type is interface{} with range:
    // 0..4294967295.
    Ifindiscards interface{}

    // The number of inbound packets that contained errors preventing them from
    // being deliverable to a higher-layer protocol. The type is interface{} with
    // range: 0..4294967295.
    Ifinerrors interface{}

    // The number of packets received via the interface which were discarded
    // because of an unknown or unsupported protocol. The type is interface{} with
    // range: 0..4294967295.
    Ifinunknownprotos interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters. The type is interface{} with range: 0..4294967295.
    Ifoutoctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted to a subnetwork-unicast address, including those that were
    // discarded or not sent. The type is interface{} with range: 0..4294967295.
    Ifoutucastpkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted to a non- unicast (i.e., a subnetwork-broadcast or
    // subnetwork-multicast) address, including those that were discarded or not
    // sent. The type is interface{} with range: 0..4294967295.
    Ifoutnucastpkts interface{}

    // The number of outbound packets which were chosen to be discarded even
    // though no errors had been detected to prevent their being transmitted.  One
    // possible reason for discarding such a packet could be to free up buffer
    // space. The type is interface{} with range: 0..4294967295.
    Ifoutdiscards interface{}

    // The number of outbound packets that could not be transmitted because of
    // errors. The type is interface{} with range: 0..4294967295.
    Ifouterrors interface{}

    // The length of the output packet queue (in packets). The type is interface{}
    // with range: 0..4294967295.
    Ifoutqlen interface{}

    // A reference to MIB definitions specific to the particular media being used
    // to realize the interface.  For example, if the interface is realized by an
    // ethernet, then the value of this object refers to a document defining
    // objects specific to ethernet.  If this information is not present, its
    // value should be set to the OBJECT IDENTIFIER { 0 0 }, which is a
    // syntactically valid object identifier, and any conformant implementation of
    // ASN.1 and BER must be able to generate and recognize this value. The type
    // is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Ifspecific interface{}
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetEntityData() *types.CommonEntityData {
    ifentry.EntityData.YFilter = ifentry.YFilter
    ifentry.EntityData.YangName = "ifEntry"
    ifentry.EntityData.BundleName = "cisco_ios_xe"
    ifentry.EntityData.ParentYangName = "ifTable"
    ifentry.EntityData.SegmentPath = "ifEntry" + "[ifIndex='" + fmt.Sprintf("%v", ifentry.Ifindex) + "']"
    ifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifentry.EntityData.Children = make(map[string]types.YChild)
    ifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", ifentry.Ifindex}
    ifentry.EntityData.Leafs["ifDescr"] = types.YLeaf{"Ifdescr", ifentry.Ifdescr}
    ifentry.EntityData.Leafs["ifType"] = types.YLeaf{"Iftype", ifentry.Iftype}
    ifentry.EntityData.Leafs["ifMtu"] = types.YLeaf{"Ifmtu", ifentry.Ifmtu}
    ifentry.EntityData.Leafs["ifSpeed"] = types.YLeaf{"Ifspeed", ifentry.Ifspeed}
    ifentry.EntityData.Leafs["ifPhysAddress"] = types.YLeaf{"Ifphysaddress", ifentry.Ifphysaddress}
    ifentry.EntityData.Leafs["ifAdminStatus"] = types.YLeaf{"Ifadminstatus", ifentry.Ifadminstatus}
    ifentry.EntityData.Leafs["ifOperStatus"] = types.YLeaf{"Ifoperstatus", ifentry.Ifoperstatus}
    ifentry.EntityData.Leafs["ifLastChange"] = types.YLeaf{"Iflastchange", ifentry.Iflastchange}
    ifentry.EntityData.Leafs["ifInOctets"] = types.YLeaf{"Ifinoctets", ifentry.Ifinoctets}
    ifentry.EntityData.Leafs["ifInUcastPkts"] = types.YLeaf{"Ifinucastpkts", ifentry.Ifinucastpkts}
    ifentry.EntityData.Leafs["ifInNUcastPkts"] = types.YLeaf{"Ifinnucastpkts", ifentry.Ifinnucastpkts}
    ifentry.EntityData.Leafs["ifInDiscards"] = types.YLeaf{"Ifindiscards", ifentry.Ifindiscards}
    ifentry.EntityData.Leafs["ifInErrors"] = types.YLeaf{"Ifinerrors", ifentry.Ifinerrors}
    ifentry.EntityData.Leafs["ifInUnknownProtos"] = types.YLeaf{"Ifinunknownprotos", ifentry.Ifinunknownprotos}
    ifentry.EntityData.Leafs["ifOutOctets"] = types.YLeaf{"Ifoutoctets", ifentry.Ifoutoctets}
    ifentry.EntityData.Leafs["ifOutUcastPkts"] = types.YLeaf{"Ifoutucastpkts", ifentry.Ifoutucastpkts}
    ifentry.EntityData.Leafs["ifOutNUcastPkts"] = types.YLeaf{"Ifoutnucastpkts", ifentry.Ifoutnucastpkts}
    ifentry.EntityData.Leafs["ifOutDiscards"] = types.YLeaf{"Ifoutdiscards", ifentry.Ifoutdiscards}
    ifentry.EntityData.Leafs["ifOutErrors"] = types.YLeaf{"Ifouterrors", ifentry.Ifouterrors}
    ifentry.EntityData.Leafs["ifOutQLen"] = types.YLeaf{"Ifoutqlen", ifentry.Ifoutqlen}
    ifentry.EntityData.Leafs["ifSpecific"] = types.YLeaf{"Ifspecific", ifentry.Ifspecific}
    return &(ifentry.EntityData)
}

// RFC1213MIB_Iftable_Ifentry_Ifadminstatus represents packets can be passed.
type RFC1213MIB_Iftable_Ifentry_Ifadminstatus string

const (
    RFC1213MIB_Iftable_Ifentry_Ifadminstatus_up RFC1213MIB_Iftable_Ifentry_Ifadminstatus = "up"

    RFC1213MIB_Iftable_Ifentry_Ifadminstatus_down RFC1213MIB_Iftable_Ifentry_Ifadminstatus = "down"

    RFC1213MIB_Iftable_Ifentry_Ifadminstatus_testing RFC1213MIB_Iftable_Ifentry_Ifadminstatus = "testing"
)

// RFC1213MIB_Iftable_Ifentry_Ifoperstatus represents packets can be passed.
type RFC1213MIB_Iftable_Ifentry_Ifoperstatus string

const (
    RFC1213MIB_Iftable_Ifentry_Ifoperstatus_up RFC1213MIB_Iftable_Ifentry_Ifoperstatus = "up"

    RFC1213MIB_Iftable_Ifentry_Ifoperstatus_down RFC1213MIB_Iftable_Ifentry_Ifoperstatus = "down"

    RFC1213MIB_Iftable_Ifentry_Ifoperstatus_testing RFC1213MIB_Iftable_Ifentry_Ifoperstatus = "testing"

    RFC1213MIB_Iftable_Ifentry_Ifoperstatus_unknown RFC1213MIB_Iftable_Ifentry_Ifoperstatus = "unknown"

    RFC1213MIB_Iftable_Ifentry_Ifoperstatus_dormant RFC1213MIB_Iftable_Ifentry_Ifoperstatus = "dormant"
)

// RFC1213MIB_Attable
// The Address Translation tables contain the
// NetworkAddress to `physical' address equivalences.
// Some interfaces do not use translation tables for
// determining address equivalences (e.g., DDN-X.25
// has an algorithmic method); if all interfaces are
// of this type, then the Address Translation table
// is empty, i.e., has zero entries.
type RFC1213MIB_Attable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one NetworkAddress to `physical' address equivalence.
    // The type is slice of RFC1213MIB_Attable_Atentry.
    Atentry []RFC1213MIB_Attable_Atentry
}

func (attable *RFC1213MIB_Attable) GetEntityData() *types.CommonEntityData {
    attable.EntityData.YFilter = attable.YFilter
    attable.EntityData.YangName = "atTable"
    attable.EntityData.BundleName = "cisco_ios_xe"
    attable.EntityData.ParentYangName = "RFC1213-MIB"
    attable.EntityData.SegmentPath = "atTable"
    attable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    attable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    attable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    attable.EntityData.Children = make(map[string]types.YChild)
    attable.EntityData.Children["atEntry"] = types.YChild{"Atentry", nil}
    for i := range attable.Atentry {
        attable.EntityData.Children[types.GetSegmentPath(&attable.Atentry[i])] = types.YChild{"Atentry", &attable.Atentry[i]}
    }
    attable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attable.EntityData)
}

// RFC1213MIB_Attable_Atentry
// Each entry contains one NetworkAddress to
// `physical' address equivalence.
type RFC1213MIB_Attable_Atentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    Atifindex interface{}

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    Atifindex2 interface{}

    // This attribute is a key. The NetworkAddress (e.g., the IP address)
    // corresponding to the media-dependent `physical' address. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Atnetaddress interface{}

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
    Atphysaddress interface{}
}

func (atentry *RFC1213MIB_Attable_Atentry) GetEntityData() *types.CommonEntityData {
    atentry.EntityData.YFilter = atentry.YFilter
    atentry.EntityData.YangName = "atEntry"
    atentry.EntityData.BundleName = "cisco_ios_xe"
    atentry.EntityData.ParentYangName = "atTable"
    atentry.EntityData.SegmentPath = "atEntry" + "[atIfIndex='" + fmt.Sprintf("%v", atentry.Atifindex) + "']" + "[atIfIndex_2='" + fmt.Sprintf("%v", atentry.Atifindex2) + "']" + "[atNetAddress='" + fmt.Sprintf("%v", atentry.Atnetaddress) + "']"
    atentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atentry.EntityData.Children = make(map[string]types.YChild)
    atentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atentry.EntityData.Leafs["atIfIndex"] = types.YLeaf{"Atifindex", atentry.Atifindex}
    atentry.EntityData.Leafs["atIfIndex_2"] = types.YLeaf{"Atifindex2", atentry.Atifindex2}
    atentry.EntityData.Leafs["atNetAddress"] = types.YLeaf{"Atnetaddress", atentry.Atnetaddress}
    atentry.EntityData.Leafs["atPhysAddress"] = types.YLeaf{"Atphysaddress", atentry.Atphysaddress}
    return &(atentry.EntityData)
}

// RFC1213MIB_Ipaddrtable
// The table of addressing information relevant to
// this entity's IP addresses.
type RFC1213MIB_Ipaddrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The addressing information for one of this entity's IP addresses. The type
    // is slice of RFC1213MIB_Ipaddrtable_Ipaddrentry.
    Ipaddrentry []RFC1213MIB_Ipaddrtable_Ipaddrentry
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetEntityData() *types.CommonEntityData {
    ipaddrtable.EntityData.YFilter = ipaddrtable.YFilter
    ipaddrtable.EntityData.YangName = "ipAddrTable"
    ipaddrtable.EntityData.BundleName = "cisco_ios_xe"
    ipaddrtable.EntityData.ParentYangName = "RFC1213-MIB"
    ipaddrtable.EntityData.SegmentPath = "ipAddrTable"
    ipaddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipaddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipaddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipaddrtable.EntityData.Children = make(map[string]types.YChild)
    ipaddrtable.EntityData.Children["ipAddrEntry"] = types.YChild{"Ipaddrentry", nil}
    for i := range ipaddrtable.Ipaddrentry {
        ipaddrtable.EntityData.Children[types.GetSegmentPath(&ipaddrtable.Ipaddrentry[i])] = types.YChild{"Ipaddrentry", &ipaddrtable.Ipaddrentry[i]}
    }
    ipaddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipaddrtable.EntityData)
}

// RFC1213MIB_Ipaddrtable_Ipaddrentry
// The addressing information for one of this
// entity's IP addresses.
type RFC1213MIB_Ipaddrtable_Ipaddrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address to which this entry's addressing
    // information pertains. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipadentaddr interface{}

    // The index value which uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of ifIndex. The
    // type is interface{} with range: -2147483648..2147483647.
    Ipadentifindex interface{}

    // The subnet mask associated with the IP address of this entry.  The value of
    // the mask is an IP address with all the network bits set to 1 and all the
    // hosts bits set to 0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipadentnetmask interface{}

    // The value of the least-significant bit in the IP broadcast address used for
    // sending datagrams on the (logical) interface associated with the IP address
    // of this entry.  For example, when the Internet standard all-ones broadcast
    // address is used, the value will be 1.  This value applies to both the
    // subnet and network broadcasts addresses used by the entity on this
    // (logical) interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Ipadentbcastaddr interface{}

    // The size of the largest IP datagram which this entity can re-assemble from
    // incoming IP fragmented datagrams received on this interface. The type is
    // interface{} with range: 0..65535.
    Ipadentreasmmaxsize interface{}
}

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetEntityData() *types.CommonEntityData {
    ipaddrentry.EntityData.YFilter = ipaddrentry.YFilter
    ipaddrentry.EntityData.YangName = "ipAddrEntry"
    ipaddrentry.EntityData.BundleName = "cisco_ios_xe"
    ipaddrentry.EntityData.ParentYangName = "ipAddrTable"
    ipaddrentry.EntityData.SegmentPath = "ipAddrEntry" + "[ipAdEntAddr='" + fmt.Sprintf("%v", ipaddrentry.Ipadentaddr) + "']"
    ipaddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipaddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipaddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipaddrentry.EntityData.Children = make(map[string]types.YChild)
    ipaddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipaddrentry.EntityData.Leafs["ipAdEntAddr"] = types.YLeaf{"Ipadentaddr", ipaddrentry.Ipadentaddr}
    ipaddrentry.EntityData.Leafs["ipAdEntIfIndex"] = types.YLeaf{"Ipadentifindex", ipaddrentry.Ipadentifindex}
    ipaddrentry.EntityData.Leafs["ipAdEntNetMask"] = types.YLeaf{"Ipadentnetmask", ipaddrentry.Ipadentnetmask}
    ipaddrentry.EntityData.Leafs["ipAdEntBcastAddr"] = types.YLeaf{"Ipadentbcastaddr", ipaddrentry.Ipadentbcastaddr}
    ipaddrentry.EntityData.Leafs["ipAdEntReasmMaxSize"] = types.YLeaf{"Ipadentreasmmaxsize", ipaddrentry.Ipadentreasmmaxsize}
    return &(ipaddrentry.EntityData)
}

// RFC1213MIB_Iproutetable
// This entity's IP Routing table.
type RFC1213MIB_Iproutetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A route to a particular destination. The type is slice of
    // RFC1213MIB_Iproutetable_Iprouteentry.
    Iprouteentry []RFC1213MIB_Iproutetable_Iprouteentry
}

func (iproutetable *RFC1213MIB_Iproutetable) GetEntityData() *types.CommonEntityData {
    iproutetable.EntityData.YFilter = iproutetable.YFilter
    iproutetable.EntityData.YangName = "ipRouteTable"
    iproutetable.EntityData.BundleName = "cisco_ios_xe"
    iproutetable.EntityData.ParentYangName = "RFC1213-MIB"
    iproutetable.EntityData.SegmentPath = "ipRouteTable"
    iproutetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iproutetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iproutetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iproutetable.EntityData.Children = make(map[string]types.YChild)
    iproutetable.EntityData.Children["ipRouteEntry"] = types.YChild{"Iprouteentry", nil}
    for i := range iproutetable.Iprouteentry {
        iproutetable.EntityData.Children[types.GetSegmentPath(&iproutetable.Iprouteentry[i])] = types.YChild{"Iprouteentry", &iproutetable.Iprouteentry[i]}
    }
    iproutetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iproutetable.EntityData)
}

// RFC1213MIB_Iproutetable_Iprouteentry
// A route to a particular destination.
type RFC1213MIB_Iproutetable_Iprouteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The destination IP address of this route.  An
    // entry with a value of 0.0.0.0 is considered a default route.  Multiple
    // routes to a single destination can appear in the table, but access to such
    // multiple entries is dependent on the table- access mechanisms defined by
    // the network management protocol in use. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Iproutedest interface{}

    // The index value which uniquely identifies the local interface through which
    // the next hop of this route should be reached.  The interface identified by
    // a particular value of this index is the same interface as identified by the
    // same value of ifIndex. The type is interface{} with range:
    // -2147483648..2147483647.
    Iprouteifindex interface{}

    // The primary routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    Iproutemetric1 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    Iproutemetric2 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    Iproutemetric3 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    Iproutemetric4 interface{}

    // The IP address of the next hop of this route. (In the case of a route bound
    // to an interface which is realized via a broadcast media, the value of this
    // field is the agent's IP address on that interface.). The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Iproutenexthop interface{}

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
    // ipRouteType object. The type is Iproutetype.
    Iproutetype interface{}

    // The routing mechanism via which this route was learned.  Inclusion of
    // values for gateway routing protocols is not intended to imply that hosts
    // should support those protocols. The type is Iprouteproto.
    Iprouteproto interface{}

    // The number of seconds since this route was last updated or otherwise
    // determined to be correct. Note that no semantics of `too old' can be
    // implied except through knowledge of the routing protocol by which the route
    // was learned. The type is interface{} with range: -2147483648..2147483647.
    Iprouteage interface{}

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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Iproutemask interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing-protocol specified in the route's
    // ipRouteProto value.  If this metric is not used, its value should be set to
    // -1. The type is interface{} with range: -2147483648..2147483647.
    Iproutemetric5 interface{}

    // A reference to MIB definitions specific to the particular routing protocol
    // which is responsible for this route, as determined by the value specified
    // in the route's ipRouteProto value.  If this information is not present, its
    // value should be set to the OBJECT IDENTIFIER { 0 0 }, which is a
    // syntactically valid object identifier, and any conformant implementation of
    // ASN.1 and BER must be able to generate and recognize this value. The type
    // is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Iprouteinfo interface{}
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetEntityData() *types.CommonEntityData {
    iprouteentry.EntityData.YFilter = iprouteentry.YFilter
    iprouteentry.EntityData.YangName = "ipRouteEntry"
    iprouteentry.EntityData.BundleName = "cisco_ios_xe"
    iprouteentry.EntityData.ParentYangName = "ipRouteTable"
    iprouteentry.EntityData.SegmentPath = "ipRouteEntry" + "[ipRouteDest='" + fmt.Sprintf("%v", iprouteentry.Iproutedest) + "']"
    iprouteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iprouteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iprouteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iprouteentry.EntityData.Children = make(map[string]types.YChild)
    iprouteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    iprouteentry.EntityData.Leafs["ipRouteDest"] = types.YLeaf{"Iproutedest", iprouteentry.Iproutedest}
    iprouteentry.EntityData.Leafs["ipRouteIfIndex"] = types.YLeaf{"Iprouteifindex", iprouteentry.Iprouteifindex}
    iprouteentry.EntityData.Leafs["ipRouteMetric1"] = types.YLeaf{"Iproutemetric1", iprouteentry.Iproutemetric1}
    iprouteentry.EntityData.Leafs["ipRouteMetric2"] = types.YLeaf{"Iproutemetric2", iprouteentry.Iproutemetric2}
    iprouteentry.EntityData.Leafs["ipRouteMetric3"] = types.YLeaf{"Iproutemetric3", iprouteentry.Iproutemetric3}
    iprouteentry.EntityData.Leafs["ipRouteMetric4"] = types.YLeaf{"Iproutemetric4", iprouteentry.Iproutemetric4}
    iprouteentry.EntityData.Leafs["ipRouteNextHop"] = types.YLeaf{"Iproutenexthop", iprouteentry.Iproutenexthop}
    iprouteentry.EntityData.Leafs["ipRouteType"] = types.YLeaf{"Iproutetype", iprouteentry.Iproutetype}
    iprouteentry.EntityData.Leafs["ipRouteProto"] = types.YLeaf{"Iprouteproto", iprouteentry.Iprouteproto}
    iprouteentry.EntityData.Leafs["ipRouteAge"] = types.YLeaf{"Iprouteage", iprouteentry.Iprouteage}
    iprouteentry.EntityData.Leafs["ipRouteMask"] = types.YLeaf{"Iproutemask", iprouteentry.Iproutemask}
    iprouteentry.EntityData.Leafs["ipRouteMetric5"] = types.YLeaf{"Iproutemetric5", iprouteentry.Iproutemetric5}
    iprouteentry.EntityData.Leafs["ipRouteInfo"] = types.YLeaf{"Iprouteinfo", iprouteentry.Iprouteinfo}
    return &(iprouteentry.EntityData)
}

// RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto represents should support those protocols.
type RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto string

const (
    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_other RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "other"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_local RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "local"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_netmgmt RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "netmgmt"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_icmp RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "icmp"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_egp RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "egp"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_ggp RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "ggp"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_hello RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "hello"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_rip RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "rip"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_is_is RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "is-is"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_es_is RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "es-is"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_ciscoIgrp RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "ciscoIgrp"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_bbnSpfIgp RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "bbnSpfIgp"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_ospf RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "ospf"

    RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto_bgp RFC1213MIB_Iproutetable_Iprouteentry_Iprouteproto = "bgp"
)

// RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype represents examination of the relevant ipRouteType object.
type RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype string

const (
    RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype_other RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype = "other"

    RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype_invalid RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype = "invalid"

    RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype_direct RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype = "direct"

    RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype_indirect RFC1213MIB_Iproutetable_Iprouteentry_Iproutetype = "indirect"
)

// RFC1213MIB_Ipnettomediatable
// The IP Address Translation table used for mapping
// from IP addresses to physical addresses.
type RFC1213MIB_Ipnettomediatable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one IpAddress to `physical' address equivalence. The
    // type is slice of RFC1213MIB_Ipnettomediatable_Ipnettomediaentry.
    Ipnettomediaentry []RFC1213MIB_Ipnettomediatable_Ipnettomediaentry
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetEntityData() *types.CommonEntityData {
    ipnettomediatable.EntityData.YFilter = ipnettomediatable.YFilter
    ipnettomediatable.EntityData.YangName = "ipNetToMediaTable"
    ipnettomediatable.EntityData.BundleName = "cisco_ios_xe"
    ipnettomediatable.EntityData.ParentYangName = "RFC1213-MIB"
    ipnettomediatable.EntityData.SegmentPath = "ipNetToMediaTable"
    ipnettomediatable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipnettomediatable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipnettomediatable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipnettomediatable.EntityData.Children = make(map[string]types.YChild)
    ipnettomediatable.EntityData.Children["ipNetToMediaEntry"] = types.YChild{"Ipnettomediaentry", nil}
    for i := range ipnettomediatable.Ipnettomediaentry {
        ipnettomediatable.EntityData.Children[types.GetSegmentPath(&ipnettomediatable.Ipnettomediaentry[i])] = types.YChild{"Ipnettomediaentry", &ipnettomediatable.Ipnettomediaentry[i]}
    }
    ipnettomediatable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipnettomediatable.EntityData)
}

// RFC1213MIB_Ipnettomediatable_Ipnettomediaentry
// Each entry contains one IpAddress to `physical'
// address equivalence.
type RFC1213MIB_Ipnettomediatable_Ipnettomediaentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    Ipnettomediaifindex interface{}

    // This attribute is a key. The IpAddress corresponding to the media-
    // dependent `physical' address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipnettomedianetaddress interface{}

    // The media-dependent `physical' address. The type is string.
    Ipnettomediaphysaddress interface{}

    // The type of mapping.  Setting this object to the value invalid(2) has the
    // effect of invalidating the corresponding entry in the ipNetToMediaTable. 
    // That is, it effectively disassociates the interface identified with said
    // entry from the mapping identified with said entry. It is an
    // implementation-specific matter as to whether the agent removes an
    // invalidated entry from the table.  Accordingly, management stations must be
    // prepared to receive tabular information from agents that corresponds to
    // entries not currently in use.  Proper interpretation of such entries
    // requires examination of the relevant ipNetToMediaType object. The type is
    // Ipnettomediatype.
    Ipnettomediatype interface{}
}

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetEntityData() *types.CommonEntityData {
    ipnettomediaentry.EntityData.YFilter = ipnettomediaentry.YFilter
    ipnettomediaentry.EntityData.YangName = "ipNetToMediaEntry"
    ipnettomediaentry.EntityData.BundleName = "cisco_ios_xe"
    ipnettomediaentry.EntityData.ParentYangName = "ipNetToMediaTable"
    ipnettomediaentry.EntityData.SegmentPath = "ipNetToMediaEntry" + "[ipNetToMediaIfIndex='" + fmt.Sprintf("%v", ipnettomediaentry.Ipnettomediaifindex) + "']" + "[ipNetToMediaNetAddress='" + fmt.Sprintf("%v", ipnettomediaentry.Ipnettomedianetaddress) + "']"
    ipnettomediaentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipnettomediaentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipnettomediaentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipnettomediaentry.EntityData.Children = make(map[string]types.YChild)
    ipnettomediaentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipnettomediaentry.EntityData.Leafs["ipNetToMediaIfIndex"] = types.YLeaf{"Ipnettomediaifindex", ipnettomediaentry.Ipnettomediaifindex}
    ipnettomediaentry.EntityData.Leafs["ipNetToMediaNetAddress"] = types.YLeaf{"Ipnettomedianetaddress", ipnettomediaentry.Ipnettomedianetaddress}
    ipnettomediaentry.EntityData.Leafs["ipNetToMediaPhysAddress"] = types.YLeaf{"Ipnettomediaphysaddress", ipnettomediaentry.Ipnettomediaphysaddress}
    ipnettomediaentry.EntityData.Leafs["ipNetToMediaType"] = types.YLeaf{"Ipnettomediatype", ipnettomediaentry.Ipnettomediatype}
    return &(ipnettomediaentry.EntityData)
}

// RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype represents ipNetToMediaType object.
type RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype string

const (
    RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_other RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "other"

    RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_invalid RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "invalid"

    RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_dynamic RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "dynamic"

    RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_static RFC1213MIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "static"
)

// RFC1213MIB_Tcpconntable
// A table containing TCP connection-specific
// information.
type RFC1213MIB_Tcpconntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current TCP connection.  An object of this
    // type is transient, in that it ceases to exist when (or soon after) the
    // connection makes the transition to the CLOSED state. The type is slice of
    // RFC1213MIB_Tcpconntable_Tcpconnentry.
    Tcpconnentry []RFC1213MIB_Tcpconntable_Tcpconnentry
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetEntityData() *types.CommonEntityData {
    tcpconntable.EntityData.YFilter = tcpconntable.YFilter
    tcpconntable.EntityData.YangName = "tcpConnTable"
    tcpconntable.EntityData.BundleName = "cisco_ios_xe"
    tcpconntable.EntityData.ParentYangName = "RFC1213-MIB"
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

// RFC1213MIB_Tcpconntable_Tcpconnentry
// Information about a particular current TCP
// connection.  An object of this type is transient,
// in that it ceases to exist when (or soon after)
// the connection makes the transition to the CLOSED
// state.
type RFC1213MIB_Tcpconntable_Tcpconnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this TCP connection.  In
    // the case of a connection in the listen state which is willing to accept
    // connections for any IP interface associated with the node, the value
    // 0.0.0.0 is used. The type is string with pattern:
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

    // The state of this TCP connection.  The only value which may be set by a
    // management station is deleteTCB(12).  Accordingly, it is appropriate for an
    // agent to return a `badValue' response if a management station attempts to
    // set this object to any other value.  If a management station sets this
    // object to the value deleteTCB(12), then this has the effect of deleting the
    // TCB (as defined in RFC 793) of the corresponding connection on the managed
    // node, resulting in immediate termination of the connection.  As an
    // implementation-specific option, a RST segment may be sent from the managed
    // node to the other TCP endpoint (note however that RST segments are not sent
    // reliably). The type is Tcpconnstate.
    Tcpconnstate interface{}
}

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetEntityData() *types.CommonEntityData {
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

// RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate represents are not sent reliably).
type RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate string

const (
    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_closed RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "closed"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_listen RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "listen"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_synSent RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "synSent"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_synReceived RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "synReceived"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_established RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "established"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_finWait1 RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "finWait1"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_finWait2 RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "finWait2"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_closeWait RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "closeWait"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_lastAck RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "lastAck"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_closing RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "closing"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_timeWait RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "timeWait"

    RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate_deleteTCB RFC1213MIB_Tcpconntable_Tcpconnentry_Tcpconnstate = "deleteTCB"
)

// RFC1213MIB_Udptable
// A table containing UDP listener information.
type RFC1213MIB_Udptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular current UDP listener. The type is slice of
    // RFC1213MIB_Udptable_Udpentry.
    Udpentry []RFC1213MIB_Udptable_Udpentry
}

func (udptable *RFC1213MIB_Udptable) GetEntityData() *types.CommonEntityData {
    udptable.EntityData.YFilter = udptable.YFilter
    udptable.EntityData.YangName = "udpTable"
    udptable.EntityData.BundleName = "cisco_ios_xe"
    udptable.EntityData.ParentYangName = "RFC1213-MIB"
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

// RFC1213MIB_Udptable_Udpentry
// Information about a particular current UDP
// listener.
type RFC1213MIB_Udptable_Udpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this UDP listener.  In
    // the case of a UDP listener which is willing to accept datagrams for any IP
    // interface associated with the node, the value 0.0.0.0 is used. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Udplocaladdress interface{}

    // This attribute is a key. The local port number for this UDP listener. The
    // type is interface{} with range: 0..65535.
    Udplocalport interface{}
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetEntityData() *types.CommonEntityData {
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

// RFC1213MIB_Egpneightable
// The EGP neighbor table.
type RFC1213MIB_Egpneightable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about this entity's relationship with a particular EGP
    // neighbor. The type is slice of RFC1213MIB_Egpneightable_Egpneighentry.
    Egpneighentry []RFC1213MIB_Egpneightable_Egpneighentry
}

func (egpneightable *RFC1213MIB_Egpneightable) GetEntityData() *types.CommonEntityData {
    egpneightable.EntityData.YFilter = egpneightable.YFilter
    egpneightable.EntityData.YangName = "egpNeighTable"
    egpneightable.EntityData.BundleName = "cisco_ios_xe"
    egpneightable.EntityData.ParentYangName = "RFC1213-MIB"
    egpneightable.EntityData.SegmentPath = "egpNeighTable"
    egpneightable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egpneightable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egpneightable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egpneightable.EntityData.Children = make(map[string]types.YChild)
    egpneightable.EntityData.Children["egpNeighEntry"] = types.YChild{"Egpneighentry", nil}
    for i := range egpneightable.Egpneighentry {
        egpneightable.EntityData.Children[types.GetSegmentPath(&egpneightable.Egpneighentry[i])] = types.YChild{"Egpneighentry", &egpneightable.Egpneighentry[i]}
    }
    egpneightable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(egpneightable.EntityData)
}

// RFC1213MIB_Egpneightable_Egpneighentry
// Information about this entity's relationship with
// a particular EGP neighbor.
type RFC1213MIB_Egpneightable_Egpneighentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this entry's EGP neighbor. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Egpneighaddr interface{}

    // The EGP state of the local system with respect to this entry's EGP
    // neighbor.  Each EGP state is represented by a value that is one greater
    // than the numerical value associated with said state in RFC 904. The type is
    // Egpneighstate.
    Egpneighstate interface{}

    // The autonomous system of this EGP peer.  Zero should be specified if the
    // autonomous system number of the neighbor is not yet known. The type is
    // interface{} with range: -2147483648..2147483647.
    Egpneighas interface{}

    // The number of EGP messages received without error from this EGP peer. The
    // type is interface{} with range: 0..4294967295.
    Egpneighinmsgs interface{}

    // The number of EGP messages received from this EGP peer that proved to be in
    // error (e.g., bad EGP checksum). The type is interface{} with range:
    // 0..4294967295.
    Egpneighinerrs interface{}

    // The number of locally generated EGP messages to this EGP peer. The type is
    // interface{} with range: 0..4294967295.
    Egpneighoutmsgs interface{}

    // The number of locally generated EGP messages not sent to this EGP peer due
    // to resource limitations within an EGP entity. The type is interface{} with
    // range: 0..4294967295.
    Egpneighouterrs interface{}

    // The number of EGP-defined error messages received from this EGP peer. The
    // type is interface{} with range: 0..4294967295.
    Egpneighinerrmsgs interface{}

    // The number of EGP-defined error messages sent to this EGP peer. The type is
    // interface{} with range: 0..4294967295.
    Egpneighouterrmsgs interface{}

    // The number of EGP state transitions to the UP state with this EGP peer. The
    // type is interface{} with range: 0..4294967295.
    Egpneighstateups interface{}

    // The number of EGP state transitions from the UP state to any other state
    // with this EGP peer. The type is interface{} with range: 0..4294967295.
    Egpneighstatedowns interface{}

    // The interval between EGP Hello command retransmissions (in hundredths of a
    // second).  This represents the t1 timer as defined in RFC 904. The type is
    // interface{} with range: -2147483648..2147483647.
    Egpneighintervalhello interface{}

    // The interval between EGP poll command retransmissions (in hundredths of a
    // second).  This represents the t3 timer as defined in RFC 904. The type is
    // interface{} with range: -2147483648..2147483647.
    Egpneighintervalpoll interface{}

    // The polling mode of this EGP entity, either passive or active. The type is
    // Egpneighmode.
    Egpneighmode interface{}

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
    // Egpneigheventtrigger.
    Egpneigheventtrigger interface{}
}

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetEntityData() *types.CommonEntityData {
    egpneighentry.EntityData.YFilter = egpneighentry.YFilter
    egpneighentry.EntityData.YangName = "egpNeighEntry"
    egpneighentry.EntityData.BundleName = "cisco_ios_xe"
    egpneighentry.EntityData.ParentYangName = "egpNeighTable"
    egpneighentry.EntityData.SegmentPath = "egpNeighEntry" + "[egpNeighAddr='" + fmt.Sprintf("%v", egpneighentry.Egpneighaddr) + "']"
    egpneighentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egpneighentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egpneighentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egpneighentry.EntityData.Children = make(map[string]types.YChild)
    egpneighentry.EntityData.Leafs = make(map[string]types.YLeaf)
    egpneighentry.EntityData.Leafs["egpNeighAddr"] = types.YLeaf{"Egpneighaddr", egpneighentry.Egpneighaddr}
    egpneighentry.EntityData.Leafs["egpNeighState"] = types.YLeaf{"Egpneighstate", egpneighentry.Egpneighstate}
    egpneighentry.EntityData.Leafs["egpNeighAs"] = types.YLeaf{"Egpneighas", egpneighentry.Egpneighas}
    egpneighentry.EntityData.Leafs["egpNeighInMsgs"] = types.YLeaf{"Egpneighinmsgs", egpneighentry.Egpneighinmsgs}
    egpneighentry.EntityData.Leafs["egpNeighInErrs"] = types.YLeaf{"Egpneighinerrs", egpneighentry.Egpneighinerrs}
    egpneighentry.EntityData.Leafs["egpNeighOutMsgs"] = types.YLeaf{"Egpneighoutmsgs", egpneighentry.Egpneighoutmsgs}
    egpneighentry.EntityData.Leafs["egpNeighOutErrs"] = types.YLeaf{"Egpneighouterrs", egpneighentry.Egpneighouterrs}
    egpneighentry.EntityData.Leafs["egpNeighInErrMsgs"] = types.YLeaf{"Egpneighinerrmsgs", egpneighentry.Egpneighinerrmsgs}
    egpneighentry.EntityData.Leafs["egpNeighOutErrMsgs"] = types.YLeaf{"Egpneighouterrmsgs", egpneighentry.Egpneighouterrmsgs}
    egpneighentry.EntityData.Leafs["egpNeighStateUps"] = types.YLeaf{"Egpneighstateups", egpneighentry.Egpneighstateups}
    egpneighentry.EntityData.Leafs["egpNeighStateDowns"] = types.YLeaf{"Egpneighstatedowns", egpneighentry.Egpneighstatedowns}
    egpneighentry.EntityData.Leafs["egpNeighIntervalHello"] = types.YLeaf{"Egpneighintervalhello", egpneighentry.Egpneighintervalhello}
    egpneighentry.EntityData.Leafs["egpNeighIntervalPoll"] = types.YLeaf{"Egpneighintervalpoll", egpneighentry.Egpneighintervalpoll}
    egpneighentry.EntityData.Leafs["egpNeighMode"] = types.YLeaf{"Egpneighmode", egpneighentry.Egpneighmode}
    egpneighentry.EntityData.Leafs["egpNeighEventTrigger"] = types.YLeaf{"Egpneigheventtrigger", egpneighentry.Egpneigheventtrigger}
    return &(egpneighentry.EntityData)
}

// RFC1213MIB_Egpneightable_Egpneighentry_Egpneigheventtrigger represents otherwise.
type RFC1213MIB_Egpneightable_Egpneighentry_Egpneigheventtrigger string

const (
    RFC1213MIB_Egpneightable_Egpneighentry_Egpneigheventtrigger_start RFC1213MIB_Egpneightable_Egpneighentry_Egpneigheventtrigger = "start"

    RFC1213MIB_Egpneightable_Egpneighentry_Egpneigheventtrigger_stop RFC1213MIB_Egpneightable_Egpneighentry_Egpneigheventtrigger = "stop"
)

// RFC1213MIB_Egpneightable_Egpneighentry_Egpneighmode represents passive or active.
type RFC1213MIB_Egpneightable_Egpneighentry_Egpneighmode string

const (
    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighmode_active RFC1213MIB_Egpneightable_Egpneighentry_Egpneighmode = "active"

    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighmode_passive RFC1213MIB_Egpneightable_Egpneighentry_Egpneighmode = "passive"
)

// RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate represents RFC 904.
type RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate string

const (
    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate_idle RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate = "idle"

    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate_acquisition RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate = "acquisition"

    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate_down RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate = "down"

    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate_up RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate = "up"

    RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate_cease RFC1213MIB_Egpneightable_Egpneighentry_Egpneighstate = "cease"
)

