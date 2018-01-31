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
    parent types.Entity
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

func (rFC1213MIB *RFC1213MIB) GetFilter() yfilter.YFilter { return rFC1213MIB.YFilter }

func (rFC1213MIB *RFC1213MIB) SetFilter(yf yfilter.YFilter) { rFC1213MIB.YFilter = yf }

func (rFC1213MIB *RFC1213MIB) GetGoName(yname string) string {
    if yname == "system" { return "System" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "ip" { return "Ip" }
    if yname == "icmp" { return "Icmp" }
    if yname == "tcp" { return "Tcp" }
    if yname == "udp" { return "Udp" }
    if yname == "egp" { return "Egp" }
    if yname == "snmp" { return "Snmp" }
    if yname == "ifTable" { return "Iftable" }
    if yname == "atTable" { return "Attable" }
    if yname == "ipAddrTable" { return "Ipaddrtable" }
    if yname == "ipRouteTable" { return "Iproutetable" }
    if yname == "ipNetToMediaTable" { return "Ipnettomediatable" }
    if yname == "tcpConnTable" { return "Tcpconntable" }
    if yname == "udpTable" { return "Udptable" }
    if yname == "egpNeighTable" { return "Egpneightable" }
    return ""
}

func (rFC1213MIB *RFC1213MIB) GetSegmentPath() string {
    return "RFC1213-MIB:RFC1213-MIB"
}

func (rFC1213MIB *RFC1213MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "system" {
        return &rFC1213MIB.System
    }
    if childYangName == "interfaces" {
        return &rFC1213MIB.Interfaces
    }
    if childYangName == "ip" {
        return &rFC1213MIB.Ip
    }
    if childYangName == "icmp" {
        return &rFC1213MIB.Icmp
    }
    if childYangName == "tcp" {
        return &rFC1213MIB.Tcp
    }
    if childYangName == "udp" {
        return &rFC1213MIB.Udp
    }
    if childYangName == "egp" {
        return &rFC1213MIB.Egp
    }
    if childYangName == "snmp" {
        return &rFC1213MIB.Snmp
    }
    if childYangName == "ifTable" {
        return &rFC1213MIB.Iftable
    }
    if childYangName == "atTable" {
        return &rFC1213MIB.Attable
    }
    if childYangName == "ipAddrTable" {
        return &rFC1213MIB.Ipaddrtable
    }
    if childYangName == "ipRouteTable" {
        return &rFC1213MIB.Iproutetable
    }
    if childYangName == "ipNetToMediaTable" {
        return &rFC1213MIB.Ipnettomediatable
    }
    if childYangName == "tcpConnTable" {
        return &rFC1213MIB.Tcpconntable
    }
    if childYangName == "udpTable" {
        return &rFC1213MIB.Udptable
    }
    if childYangName == "egpNeighTable" {
        return &rFC1213MIB.Egpneightable
    }
    return nil
}

func (rFC1213MIB *RFC1213MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["system"] = &rFC1213MIB.System
    children["interfaces"] = &rFC1213MIB.Interfaces
    children["ip"] = &rFC1213MIB.Ip
    children["icmp"] = &rFC1213MIB.Icmp
    children["tcp"] = &rFC1213MIB.Tcp
    children["udp"] = &rFC1213MIB.Udp
    children["egp"] = &rFC1213MIB.Egp
    children["snmp"] = &rFC1213MIB.Snmp
    children["ifTable"] = &rFC1213MIB.Iftable
    children["atTable"] = &rFC1213MIB.Attable
    children["ipAddrTable"] = &rFC1213MIB.Ipaddrtable
    children["ipRouteTable"] = &rFC1213MIB.Iproutetable
    children["ipNetToMediaTable"] = &rFC1213MIB.Ipnettomediatable
    children["tcpConnTable"] = &rFC1213MIB.Tcpconntable
    children["udpTable"] = &rFC1213MIB.Udptable
    children["egpNeighTable"] = &rFC1213MIB.Egpneightable
    return children
}

func (rFC1213MIB *RFC1213MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rFC1213MIB *RFC1213MIB) GetBundleName() string { return "cisco_ios_xe" }

func (rFC1213MIB *RFC1213MIB) GetYangName() string { return "RFC1213-MIB" }

func (rFC1213MIB *RFC1213MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rFC1213MIB *RFC1213MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rFC1213MIB *RFC1213MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rFC1213MIB *RFC1213MIB) SetParent(parent types.Entity) { rFC1213MIB.parent = parent }

func (rFC1213MIB *RFC1213MIB) GetParent() types.Entity { return rFC1213MIB.parent }

func (rFC1213MIB *RFC1213MIB) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_System
type RFC1213MIB_System struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (system *RFC1213MIB_System) GetFilter() yfilter.YFilter { return system.YFilter }

func (system *RFC1213MIB_System) SetFilter(yf yfilter.YFilter) { system.YFilter = yf }

func (system *RFC1213MIB_System) GetGoName(yname string) string {
    if yname == "sysDescr" { return "Sysdescr" }
    if yname == "sysObjectID" { return "Sysobjectid" }
    if yname == "sysUpTime" { return "Sysuptime" }
    if yname == "sysContact" { return "Syscontact" }
    if yname == "sysName" { return "Sysname" }
    if yname == "sysLocation" { return "Syslocation" }
    if yname == "sysServices" { return "Sysservices" }
    return ""
}

func (system *RFC1213MIB_System) GetSegmentPath() string {
    return "system"
}

func (system *RFC1213MIB_System) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (system *RFC1213MIB_System) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (system *RFC1213MIB_System) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sysDescr"] = system.Sysdescr
    leafs["sysObjectID"] = system.Sysobjectid
    leafs["sysUpTime"] = system.Sysuptime
    leafs["sysContact"] = system.Syscontact
    leafs["sysName"] = system.Sysname
    leafs["sysLocation"] = system.Syslocation
    leafs["sysServices"] = system.Sysservices
    return leafs
}

func (system *RFC1213MIB_System) GetBundleName() string { return "cisco_ios_xe" }

func (system *RFC1213MIB_System) GetYangName() string { return "system" }

func (system *RFC1213MIB_System) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (system *RFC1213MIB_System) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (system *RFC1213MIB_System) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (system *RFC1213MIB_System) SetParent(parent types.Entity) { system.parent = parent }

func (system *RFC1213MIB_System) GetParent() types.Entity { return system.parent }

func (system *RFC1213MIB_System) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Interfaces
type RFC1213MIB_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of network interfaces (regardless of their current state)
    // present on this system. The type is interface{} with range:
    // -2147483648..2147483647.
    Ifnumber interface{}
}

func (interfaces *RFC1213MIB_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *RFC1213MIB_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *RFC1213MIB_Interfaces) GetGoName(yname string) string {
    if yname == "ifNumber" { return "Ifnumber" }
    return ""
}

func (interfaces *RFC1213MIB_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *RFC1213MIB_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaces *RFC1213MIB_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaces *RFC1213MIB_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifNumber"] = interfaces.Ifnumber
    return leafs
}

func (interfaces *RFC1213MIB_Interfaces) GetBundleName() string { return "cisco_ios_xe" }

func (interfaces *RFC1213MIB_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *RFC1213MIB_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (interfaces *RFC1213MIB_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (interfaces *RFC1213MIB_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (interfaces *RFC1213MIB_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *RFC1213MIB_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *RFC1213MIB_Interfaces) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Ip
type RFC1213MIB_Ip struct {
    parent types.Entity
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

func (ip *RFC1213MIB_Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *RFC1213MIB_Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *RFC1213MIB_Ip) GetGoName(yname string) string {
    if yname == "ipForwarding" { return "Ipforwarding" }
    if yname == "ipDefaultTTL" { return "Ipdefaultttl" }
    if yname == "ipInReceives" { return "Ipinreceives" }
    if yname == "ipInHdrErrors" { return "Ipinhdrerrors" }
    if yname == "ipInAddrErrors" { return "Ipinaddrerrors" }
    if yname == "ipForwDatagrams" { return "Ipforwdatagrams" }
    if yname == "ipInUnknownProtos" { return "Ipinunknownprotos" }
    if yname == "ipInDiscards" { return "Ipindiscards" }
    if yname == "ipInDelivers" { return "Ipindelivers" }
    if yname == "ipOutRequests" { return "Ipoutrequests" }
    if yname == "ipOutDiscards" { return "Ipoutdiscards" }
    if yname == "ipOutNoRoutes" { return "Ipoutnoroutes" }
    if yname == "ipReasmTimeout" { return "Ipreasmtimeout" }
    if yname == "ipReasmReqds" { return "Ipreasmreqds" }
    if yname == "ipReasmOKs" { return "Ipreasmoks" }
    if yname == "ipReasmFails" { return "Ipreasmfails" }
    if yname == "ipFragOKs" { return "Ipfragoks" }
    if yname == "ipFragFails" { return "Ipfragfails" }
    if yname == "ipFragCreates" { return "Ipfragcreates" }
    if yname == "ipRoutingDiscards" { return "Iproutingdiscards" }
    return ""
}

func (ip *RFC1213MIB_Ip) GetSegmentPath() string {
    return "ip"
}

func (ip *RFC1213MIB_Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ip *RFC1213MIB_Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ip *RFC1213MIB_Ip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipForwarding"] = ip.Ipforwarding
    leafs["ipDefaultTTL"] = ip.Ipdefaultttl
    leafs["ipInReceives"] = ip.Ipinreceives
    leafs["ipInHdrErrors"] = ip.Ipinhdrerrors
    leafs["ipInAddrErrors"] = ip.Ipinaddrerrors
    leafs["ipForwDatagrams"] = ip.Ipforwdatagrams
    leafs["ipInUnknownProtos"] = ip.Ipinunknownprotos
    leafs["ipInDiscards"] = ip.Ipindiscards
    leafs["ipInDelivers"] = ip.Ipindelivers
    leafs["ipOutRequests"] = ip.Ipoutrequests
    leafs["ipOutDiscards"] = ip.Ipoutdiscards
    leafs["ipOutNoRoutes"] = ip.Ipoutnoroutes
    leafs["ipReasmTimeout"] = ip.Ipreasmtimeout
    leafs["ipReasmReqds"] = ip.Ipreasmreqds
    leafs["ipReasmOKs"] = ip.Ipreasmoks
    leafs["ipReasmFails"] = ip.Ipreasmfails
    leafs["ipFragOKs"] = ip.Ipfragoks
    leafs["ipFragFails"] = ip.Ipfragfails
    leafs["ipFragCreates"] = ip.Ipfragcreates
    leafs["ipRoutingDiscards"] = ip.Iproutingdiscards
    return leafs
}

func (ip *RFC1213MIB_Ip) GetBundleName() string { return "cisco_ios_xe" }

func (ip *RFC1213MIB_Ip) GetYangName() string { return "ip" }

func (ip *RFC1213MIB_Ip) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ip *RFC1213MIB_Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ip *RFC1213MIB_Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ip *RFC1213MIB_Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *RFC1213MIB_Ip) GetParent() types.Entity { return ip.parent }

func (ip *RFC1213MIB_Ip) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Ip_Ipforwarding represents inappropriate value.
type RFC1213MIB_Ip_Ipforwarding string

const (
    RFC1213MIB_Ip_Ipforwarding_forwarding RFC1213MIB_Ip_Ipforwarding = "forwarding"

    RFC1213MIB_Ip_Ipforwarding_not_forwarding RFC1213MIB_Ip_Ipforwarding = "not-forwarding"
)

// RFC1213MIB_Icmp
type RFC1213MIB_Icmp struct {
    parent types.Entity
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

func (icmp *RFC1213MIB_Icmp) GetFilter() yfilter.YFilter { return icmp.YFilter }

func (icmp *RFC1213MIB_Icmp) SetFilter(yf yfilter.YFilter) { icmp.YFilter = yf }

func (icmp *RFC1213MIB_Icmp) GetGoName(yname string) string {
    if yname == "icmpInMsgs" { return "Icmpinmsgs" }
    if yname == "icmpInErrors" { return "Icmpinerrors" }
    if yname == "icmpInDestUnreachs" { return "Icmpindestunreachs" }
    if yname == "icmpInTimeExcds" { return "Icmpintimeexcds" }
    if yname == "icmpInParmProbs" { return "Icmpinparmprobs" }
    if yname == "icmpInSrcQuenchs" { return "Icmpinsrcquenchs" }
    if yname == "icmpInRedirects" { return "Icmpinredirects" }
    if yname == "icmpInEchos" { return "Icmpinechos" }
    if yname == "icmpInEchoReps" { return "Icmpinechoreps" }
    if yname == "icmpInTimestamps" { return "Icmpintimestamps" }
    if yname == "icmpInTimestampReps" { return "Icmpintimestampreps" }
    if yname == "icmpInAddrMasks" { return "Icmpinaddrmasks" }
    if yname == "icmpInAddrMaskReps" { return "Icmpinaddrmaskreps" }
    if yname == "icmpOutMsgs" { return "Icmpoutmsgs" }
    if yname == "icmpOutErrors" { return "Icmpouterrors" }
    if yname == "icmpOutDestUnreachs" { return "Icmpoutdestunreachs" }
    if yname == "icmpOutTimeExcds" { return "Icmpouttimeexcds" }
    if yname == "icmpOutParmProbs" { return "Icmpoutparmprobs" }
    if yname == "icmpOutSrcQuenchs" { return "Icmpoutsrcquenchs" }
    if yname == "icmpOutRedirects" { return "Icmpoutredirects" }
    if yname == "icmpOutEchos" { return "Icmpoutechos" }
    if yname == "icmpOutEchoReps" { return "Icmpoutechoreps" }
    if yname == "icmpOutTimestamps" { return "Icmpouttimestamps" }
    if yname == "icmpOutTimestampReps" { return "Icmpouttimestampreps" }
    if yname == "icmpOutAddrMasks" { return "Icmpoutaddrmasks" }
    if yname == "icmpOutAddrMaskReps" { return "Icmpoutaddrmaskreps" }
    return ""
}

func (icmp *RFC1213MIB_Icmp) GetSegmentPath() string {
    return "icmp"
}

func (icmp *RFC1213MIB_Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmp *RFC1213MIB_Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmp *RFC1213MIB_Icmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icmpInMsgs"] = icmp.Icmpinmsgs
    leafs["icmpInErrors"] = icmp.Icmpinerrors
    leafs["icmpInDestUnreachs"] = icmp.Icmpindestunreachs
    leafs["icmpInTimeExcds"] = icmp.Icmpintimeexcds
    leafs["icmpInParmProbs"] = icmp.Icmpinparmprobs
    leafs["icmpInSrcQuenchs"] = icmp.Icmpinsrcquenchs
    leafs["icmpInRedirects"] = icmp.Icmpinredirects
    leafs["icmpInEchos"] = icmp.Icmpinechos
    leafs["icmpInEchoReps"] = icmp.Icmpinechoreps
    leafs["icmpInTimestamps"] = icmp.Icmpintimestamps
    leafs["icmpInTimestampReps"] = icmp.Icmpintimestampreps
    leafs["icmpInAddrMasks"] = icmp.Icmpinaddrmasks
    leafs["icmpInAddrMaskReps"] = icmp.Icmpinaddrmaskreps
    leafs["icmpOutMsgs"] = icmp.Icmpoutmsgs
    leafs["icmpOutErrors"] = icmp.Icmpouterrors
    leafs["icmpOutDestUnreachs"] = icmp.Icmpoutdestunreachs
    leafs["icmpOutTimeExcds"] = icmp.Icmpouttimeexcds
    leafs["icmpOutParmProbs"] = icmp.Icmpoutparmprobs
    leafs["icmpOutSrcQuenchs"] = icmp.Icmpoutsrcquenchs
    leafs["icmpOutRedirects"] = icmp.Icmpoutredirects
    leafs["icmpOutEchos"] = icmp.Icmpoutechos
    leafs["icmpOutEchoReps"] = icmp.Icmpoutechoreps
    leafs["icmpOutTimestamps"] = icmp.Icmpouttimestamps
    leafs["icmpOutTimestampReps"] = icmp.Icmpouttimestampreps
    leafs["icmpOutAddrMasks"] = icmp.Icmpoutaddrmasks
    leafs["icmpOutAddrMaskReps"] = icmp.Icmpoutaddrmaskreps
    return leafs
}

func (icmp *RFC1213MIB_Icmp) GetBundleName() string { return "cisco_ios_xe" }

func (icmp *RFC1213MIB_Icmp) GetYangName() string { return "icmp" }

func (icmp *RFC1213MIB_Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmp *RFC1213MIB_Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmp *RFC1213MIB_Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmp *RFC1213MIB_Icmp) SetParent(parent types.Entity) { icmp.parent = parent }

func (icmp *RFC1213MIB_Icmp) GetParent() types.Entity { return icmp.parent }

func (icmp *RFC1213MIB_Icmp) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Tcp
type RFC1213MIB_Tcp struct {
    parent types.Entity
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

func (tcp *RFC1213MIB_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *RFC1213MIB_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *RFC1213MIB_Tcp) GetGoName(yname string) string {
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
    return ""
}

func (tcp *RFC1213MIB_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *RFC1213MIB_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *RFC1213MIB_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *RFC1213MIB_Tcp) GetLeafs() map[string]interface{} {
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
    return leafs
}

func (tcp *RFC1213MIB_Tcp) GetBundleName() string { return "cisco_ios_xe" }

func (tcp *RFC1213MIB_Tcp) GetYangName() string { return "tcp" }

func (tcp *RFC1213MIB_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcp *RFC1213MIB_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcp *RFC1213MIB_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcp *RFC1213MIB_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *RFC1213MIB_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *RFC1213MIB_Tcp) GetParentYangName() string { return "RFC1213-MIB" }

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
    parent types.Entity
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

func (udp *RFC1213MIB_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *RFC1213MIB_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *RFC1213MIB_Udp) GetGoName(yname string) string {
    if yname == "udpInDatagrams" { return "Udpindatagrams" }
    if yname == "udpNoPorts" { return "Udpnoports" }
    if yname == "udpInErrors" { return "Udpinerrors" }
    if yname == "udpOutDatagrams" { return "Udpoutdatagrams" }
    return ""
}

func (udp *RFC1213MIB_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *RFC1213MIB_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udp *RFC1213MIB_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udp *RFC1213MIB_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udpInDatagrams"] = udp.Udpindatagrams
    leafs["udpNoPorts"] = udp.Udpnoports
    leafs["udpInErrors"] = udp.Udpinerrors
    leafs["udpOutDatagrams"] = udp.Udpoutdatagrams
    return leafs
}

func (udp *RFC1213MIB_Udp) GetBundleName() string { return "cisco_ios_xe" }

func (udp *RFC1213MIB_Udp) GetYangName() string { return "udp" }

func (udp *RFC1213MIB_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udp *RFC1213MIB_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udp *RFC1213MIB_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udp *RFC1213MIB_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *RFC1213MIB_Udp) GetParent() types.Entity { return udp.parent }

func (udp *RFC1213MIB_Udp) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Egp
type RFC1213MIB_Egp struct {
    parent types.Entity
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

func (egp *RFC1213MIB_Egp) GetFilter() yfilter.YFilter { return egp.YFilter }

func (egp *RFC1213MIB_Egp) SetFilter(yf yfilter.YFilter) { egp.YFilter = yf }

func (egp *RFC1213MIB_Egp) GetGoName(yname string) string {
    if yname == "egpInMsgs" { return "Egpinmsgs" }
    if yname == "egpInErrors" { return "Egpinerrors" }
    if yname == "egpOutMsgs" { return "Egpoutmsgs" }
    if yname == "egpOutErrors" { return "Egpouterrors" }
    if yname == "egpAs" { return "Egpas" }
    return ""
}

func (egp *RFC1213MIB_Egp) GetSegmentPath() string {
    return "egp"
}

func (egp *RFC1213MIB_Egp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egp *RFC1213MIB_Egp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egp *RFC1213MIB_Egp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["egpInMsgs"] = egp.Egpinmsgs
    leafs["egpInErrors"] = egp.Egpinerrors
    leafs["egpOutMsgs"] = egp.Egpoutmsgs
    leafs["egpOutErrors"] = egp.Egpouterrors
    leafs["egpAs"] = egp.Egpas
    return leafs
}

func (egp *RFC1213MIB_Egp) GetBundleName() string { return "cisco_ios_xe" }

func (egp *RFC1213MIB_Egp) GetYangName() string { return "egp" }

func (egp *RFC1213MIB_Egp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (egp *RFC1213MIB_Egp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (egp *RFC1213MIB_Egp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (egp *RFC1213MIB_Egp) SetParent(parent types.Entity) { egp.parent = parent }

func (egp *RFC1213MIB_Egp) GetParent() types.Entity { return egp.parent }

func (egp *RFC1213MIB_Egp) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Snmp
type RFC1213MIB_Snmp struct {
    parent types.Entity
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

func (snmp *RFC1213MIB_Snmp) GetFilter() yfilter.YFilter { return snmp.YFilter }

func (snmp *RFC1213MIB_Snmp) SetFilter(yf yfilter.YFilter) { snmp.YFilter = yf }

func (snmp *RFC1213MIB_Snmp) GetGoName(yname string) string {
    if yname == "snmpInPkts" { return "Snmpinpkts" }
    if yname == "snmpOutPkts" { return "Snmpoutpkts" }
    if yname == "snmpInBadVersions" { return "Snmpinbadversions" }
    if yname == "snmpInBadCommunityNames" { return "Snmpinbadcommunitynames" }
    if yname == "snmpInBadCommunityUses" { return "Snmpinbadcommunityuses" }
    if yname == "snmpInASNParseErrs" { return "Snmpinasnparseerrs" }
    if yname == "snmpInTooBigs" { return "Snmpintoobigs" }
    if yname == "snmpInNoSuchNames" { return "Snmpinnosuchnames" }
    if yname == "snmpInBadValues" { return "Snmpinbadvalues" }
    if yname == "snmpInReadOnlys" { return "Snmpinreadonlys" }
    if yname == "snmpInGenErrs" { return "Snmpingenerrs" }
    if yname == "snmpInTotalReqVars" { return "Snmpintotalreqvars" }
    if yname == "snmpInTotalSetVars" { return "Snmpintotalsetvars" }
    if yname == "snmpInGetRequests" { return "Snmpingetrequests" }
    if yname == "snmpInGetNexts" { return "Snmpingetnexts" }
    if yname == "snmpInSetRequests" { return "Snmpinsetrequests" }
    if yname == "snmpInGetResponses" { return "Snmpingetresponses" }
    if yname == "snmpInTraps" { return "Snmpintraps" }
    if yname == "snmpOutTooBigs" { return "Snmpouttoobigs" }
    if yname == "snmpOutNoSuchNames" { return "Snmpoutnosuchnames" }
    if yname == "snmpOutBadValues" { return "Snmpoutbadvalues" }
    if yname == "snmpOutGenErrs" { return "Snmpoutgenerrs" }
    if yname == "snmpOutGetRequests" { return "Snmpoutgetrequests" }
    if yname == "snmpOutGetNexts" { return "Snmpoutgetnexts" }
    if yname == "snmpOutSetRequests" { return "Snmpoutsetrequests" }
    if yname == "snmpOutGetResponses" { return "Snmpoutgetresponses" }
    if yname == "snmpOutTraps" { return "Snmpouttraps" }
    if yname == "snmpEnableAuthenTraps" { return "Snmpenableauthentraps" }
    return ""
}

func (snmp *RFC1213MIB_Snmp) GetSegmentPath() string {
    return "snmp"
}

func (snmp *RFC1213MIB_Snmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmp *RFC1213MIB_Snmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmp *RFC1213MIB_Snmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpInPkts"] = snmp.Snmpinpkts
    leafs["snmpOutPkts"] = snmp.Snmpoutpkts
    leafs["snmpInBadVersions"] = snmp.Snmpinbadversions
    leafs["snmpInBadCommunityNames"] = snmp.Snmpinbadcommunitynames
    leafs["snmpInBadCommunityUses"] = snmp.Snmpinbadcommunityuses
    leafs["snmpInASNParseErrs"] = snmp.Snmpinasnparseerrs
    leafs["snmpInTooBigs"] = snmp.Snmpintoobigs
    leafs["snmpInNoSuchNames"] = snmp.Snmpinnosuchnames
    leafs["snmpInBadValues"] = snmp.Snmpinbadvalues
    leafs["snmpInReadOnlys"] = snmp.Snmpinreadonlys
    leafs["snmpInGenErrs"] = snmp.Snmpingenerrs
    leafs["snmpInTotalReqVars"] = snmp.Snmpintotalreqvars
    leafs["snmpInTotalSetVars"] = snmp.Snmpintotalsetvars
    leafs["snmpInGetRequests"] = snmp.Snmpingetrequests
    leafs["snmpInGetNexts"] = snmp.Snmpingetnexts
    leafs["snmpInSetRequests"] = snmp.Snmpinsetrequests
    leafs["snmpInGetResponses"] = snmp.Snmpingetresponses
    leafs["snmpInTraps"] = snmp.Snmpintraps
    leafs["snmpOutTooBigs"] = snmp.Snmpouttoobigs
    leafs["snmpOutNoSuchNames"] = snmp.Snmpoutnosuchnames
    leafs["snmpOutBadValues"] = snmp.Snmpoutbadvalues
    leafs["snmpOutGenErrs"] = snmp.Snmpoutgenerrs
    leafs["snmpOutGetRequests"] = snmp.Snmpoutgetrequests
    leafs["snmpOutGetNexts"] = snmp.Snmpoutgetnexts
    leafs["snmpOutSetRequests"] = snmp.Snmpoutsetrequests
    leafs["snmpOutGetResponses"] = snmp.Snmpoutgetresponses
    leafs["snmpOutTraps"] = snmp.Snmpouttraps
    leafs["snmpEnableAuthenTraps"] = snmp.Snmpenableauthentraps
    return leafs
}

func (snmp *RFC1213MIB_Snmp) GetBundleName() string { return "cisco_ios_xe" }

func (snmp *RFC1213MIB_Snmp) GetYangName() string { return "snmp" }

func (snmp *RFC1213MIB_Snmp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmp *RFC1213MIB_Snmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmp *RFC1213MIB_Snmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmp *RFC1213MIB_Snmp) SetParent(parent types.Entity) { snmp.parent = parent }

func (snmp *RFC1213MIB_Snmp) GetParent() types.Entity { return snmp.parent }

func (snmp *RFC1213MIB_Snmp) GetParentYangName() string { return "RFC1213-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An interface entry containing objects at the subnetwork layer and below for
    // a particular interface. The type is slice of RFC1213MIB_Iftable_Ifentry.
    Ifentry []RFC1213MIB_Iftable_Ifentry
}

func (iftable *RFC1213MIB_Iftable) GetFilter() yfilter.YFilter { return iftable.YFilter }

func (iftable *RFC1213MIB_Iftable) SetFilter(yf yfilter.YFilter) { iftable.YFilter = yf }

func (iftable *RFC1213MIB_Iftable) GetGoName(yname string) string {
    if yname == "ifEntry" { return "Ifentry" }
    return ""
}

func (iftable *RFC1213MIB_Iftable) GetSegmentPath() string {
    return "ifTable"
}

func (iftable *RFC1213MIB_Iftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ifEntry" {
        for _, c := range iftable.Ifentry {
            if iftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Iftable_Ifentry{}
        iftable.Ifentry = append(iftable.Ifentry, child)
        return &iftable.Ifentry[len(iftable.Ifentry)-1]
    }
    return nil
}

func (iftable *RFC1213MIB_Iftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iftable.Ifentry {
        children[iftable.Ifentry[i].GetSegmentPath()] = &iftable.Ifentry[i]
    }
    return children
}

func (iftable *RFC1213MIB_Iftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iftable *RFC1213MIB_Iftable) GetBundleName() string { return "cisco_ios_xe" }

func (iftable *RFC1213MIB_Iftable) GetYangName() string { return "ifTable" }

func (iftable *RFC1213MIB_Iftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iftable *RFC1213MIB_Iftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iftable *RFC1213MIB_Iftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iftable *RFC1213MIB_Iftable) SetParent(parent types.Entity) { iftable.parent = parent }

func (iftable *RFC1213MIB_Iftable) GetParent() types.Entity { return iftable.parent }

func (iftable *RFC1213MIB_Iftable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Iftable_Ifentry
// An interface entry containing objects at the
// subnetwork layer and below for a particular
// interface.
type RFC1213MIB_Iftable_Ifentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Ifspecific interface{}
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetFilter() yfilter.YFilter { return ifentry.YFilter }

func (ifentry *RFC1213MIB_Iftable_Ifentry) SetFilter(yf yfilter.YFilter) { ifentry.YFilter = yf }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "ifDescr" { return "Ifdescr" }
    if yname == "ifType" { return "Iftype" }
    if yname == "ifMtu" { return "Ifmtu" }
    if yname == "ifSpeed" { return "Ifspeed" }
    if yname == "ifPhysAddress" { return "Ifphysaddress" }
    if yname == "ifAdminStatus" { return "Ifadminstatus" }
    if yname == "ifOperStatus" { return "Ifoperstatus" }
    if yname == "ifLastChange" { return "Iflastchange" }
    if yname == "ifInOctets" { return "Ifinoctets" }
    if yname == "ifInUcastPkts" { return "Ifinucastpkts" }
    if yname == "ifInNUcastPkts" { return "Ifinnucastpkts" }
    if yname == "ifInDiscards" { return "Ifindiscards" }
    if yname == "ifInErrors" { return "Ifinerrors" }
    if yname == "ifInUnknownProtos" { return "Ifinunknownprotos" }
    if yname == "ifOutOctets" { return "Ifoutoctets" }
    if yname == "ifOutUcastPkts" { return "Ifoutucastpkts" }
    if yname == "ifOutNUcastPkts" { return "Ifoutnucastpkts" }
    if yname == "ifOutDiscards" { return "Ifoutdiscards" }
    if yname == "ifOutErrors" { return "Ifouterrors" }
    if yname == "ifOutQLen" { return "Ifoutqlen" }
    if yname == "ifSpecific" { return "Ifspecific" }
    return ""
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetSegmentPath() string {
    return "ifEntry" + "[ifIndex='" + fmt.Sprintf("%v", ifentry.Ifindex) + "']"
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = ifentry.Ifindex
    leafs["ifDescr"] = ifentry.Ifdescr
    leafs["ifType"] = ifentry.Iftype
    leafs["ifMtu"] = ifentry.Ifmtu
    leafs["ifSpeed"] = ifentry.Ifspeed
    leafs["ifPhysAddress"] = ifentry.Ifphysaddress
    leafs["ifAdminStatus"] = ifentry.Ifadminstatus
    leafs["ifOperStatus"] = ifentry.Ifoperstatus
    leafs["ifLastChange"] = ifentry.Iflastchange
    leafs["ifInOctets"] = ifentry.Ifinoctets
    leafs["ifInUcastPkts"] = ifentry.Ifinucastpkts
    leafs["ifInNUcastPkts"] = ifentry.Ifinnucastpkts
    leafs["ifInDiscards"] = ifentry.Ifindiscards
    leafs["ifInErrors"] = ifentry.Ifinerrors
    leafs["ifInUnknownProtos"] = ifentry.Ifinunknownprotos
    leafs["ifOutOctets"] = ifentry.Ifoutoctets
    leafs["ifOutUcastPkts"] = ifentry.Ifoutucastpkts
    leafs["ifOutNUcastPkts"] = ifentry.Ifoutnucastpkts
    leafs["ifOutDiscards"] = ifentry.Ifoutdiscards
    leafs["ifOutErrors"] = ifentry.Ifouterrors
    leafs["ifOutQLen"] = ifentry.Ifoutqlen
    leafs["ifSpecific"] = ifentry.Ifspecific
    return leafs
}

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetBundleName() string { return "cisco_ios_xe" }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetYangName() string { return "ifEntry" }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifentry *RFC1213MIB_Iftable_Ifentry) SetParent(parent types.Entity) { ifentry.parent = parent }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetParent() types.Entity { return ifentry.parent }

func (ifentry *RFC1213MIB_Iftable_Ifentry) GetParentYangName() string { return "ifTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains one NetworkAddress to `physical' address equivalence.
    // The type is slice of RFC1213MIB_Attable_Atentry.
    Atentry []RFC1213MIB_Attable_Atentry
}

func (attable *RFC1213MIB_Attable) GetFilter() yfilter.YFilter { return attable.YFilter }

func (attable *RFC1213MIB_Attable) SetFilter(yf yfilter.YFilter) { attable.YFilter = yf }

func (attable *RFC1213MIB_Attable) GetGoName(yname string) string {
    if yname == "atEntry" { return "Atentry" }
    return ""
}

func (attable *RFC1213MIB_Attable) GetSegmentPath() string {
    return "atTable"
}

func (attable *RFC1213MIB_Attable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atEntry" {
        for _, c := range attable.Atentry {
            if attable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Attable_Atentry{}
        attable.Atentry = append(attable.Atentry, child)
        return &attable.Atentry[len(attable.Atentry)-1]
    }
    return nil
}

func (attable *RFC1213MIB_Attable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attable.Atentry {
        children[attable.Atentry[i].GetSegmentPath()] = &attable.Atentry[i]
    }
    return children
}

func (attable *RFC1213MIB_Attable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attable *RFC1213MIB_Attable) GetBundleName() string { return "cisco_ios_xe" }

func (attable *RFC1213MIB_Attable) GetYangName() string { return "atTable" }

func (attable *RFC1213MIB_Attable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (attable *RFC1213MIB_Attable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (attable *RFC1213MIB_Attable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (attable *RFC1213MIB_Attable) SetParent(parent types.Entity) { attable.parent = parent }

func (attable *RFC1213MIB_Attable) GetParent() types.Entity { return attable.parent }

func (attable *RFC1213MIB_Attable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Attable_Atentry
// Each entry contains one NetworkAddress to
// `physical' address equivalence.
type RFC1213MIB_Attable_Atentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (atentry *RFC1213MIB_Attable_Atentry) GetFilter() yfilter.YFilter { return atentry.YFilter }

func (atentry *RFC1213MIB_Attable_Atentry) SetFilter(yf yfilter.YFilter) { atentry.YFilter = yf }

func (atentry *RFC1213MIB_Attable_Atentry) GetGoName(yname string) string {
    if yname == "atIfIndex" { return "Atifindex" }
    if yname == "atIfIndex_2" { return "Atifindex2" }
    if yname == "atNetAddress" { return "Atnetaddress" }
    if yname == "atPhysAddress" { return "Atphysaddress" }
    return ""
}

func (atentry *RFC1213MIB_Attable_Atentry) GetSegmentPath() string {
    return "atEntry" + "[atIfIndex='" + fmt.Sprintf("%v", atentry.Atifindex) + "']" + "[atIfIndex_2='" + fmt.Sprintf("%v", atentry.Atifindex2) + "']" + "[atNetAddress='" + fmt.Sprintf("%v", atentry.Atnetaddress) + "']"
}

func (atentry *RFC1213MIB_Attable_Atentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atentry *RFC1213MIB_Attable_Atentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atentry *RFC1213MIB_Attable_Atentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["atIfIndex"] = atentry.Atifindex
    leafs["atIfIndex_2"] = atentry.Atifindex2
    leafs["atNetAddress"] = atentry.Atnetaddress
    leafs["atPhysAddress"] = atentry.Atphysaddress
    return leafs
}

func (atentry *RFC1213MIB_Attable_Atentry) GetBundleName() string { return "cisco_ios_xe" }

func (atentry *RFC1213MIB_Attable_Atentry) GetYangName() string { return "atEntry" }

func (atentry *RFC1213MIB_Attable_Atentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atentry *RFC1213MIB_Attable_Atentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atentry *RFC1213MIB_Attable_Atentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atentry *RFC1213MIB_Attable_Atentry) SetParent(parent types.Entity) { atentry.parent = parent }

func (atentry *RFC1213MIB_Attable_Atentry) GetParent() types.Entity { return atentry.parent }

func (atentry *RFC1213MIB_Attable_Atentry) GetParentYangName() string { return "atTable" }

// RFC1213MIB_Ipaddrtable
// The table of addressing information relevant to
// this entity's IP addresses.
type RFC1213MIB_Ipaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The addressing information for one of this entity's IP addresses. The type
    // is slice of RFC1213MIB_Ipaddrtable_Ipaddrentry.
    Ipaddrentry []RFC1213MIB_Ipaddrtable_Ipaddrentry
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetFilter() yfilter.YFilter { return ipaddrtable.YFilter }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) SetFilter(yf yfilter.YFilter) { ipaddrtable.YFilter = yf }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetGoName(yname string) string {
    if yname == "ipAddrEntry" { return "Ipaddrentry" }
    return ""
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetSegmentPath() string {
    return "ipAddrTable"
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipAddrEntry" {
        for _, c := range ipaddrtable.Ipaddrentry {
            if ipaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Ipaddrtable_Ipaddrentry{}
        ipaddrtable.Ipaddrentry = append(ipaddrtable.Ipaddrentry, child)
        return &ipaddrtable.Ipaddrentry[len(ipaddrtable.Ipaddrentry)-1]
    }
    return nil
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipaddrtable.Ipaddrentry {
        children[ipaddrtable.Ipaddrentry[i].GetSegmentPath()] = &ipaddrtable.Ipaddrentry[i]
    }
    return children
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetYangName() string { return "ipAddrTable" }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) SetParent(parent types.Entity) { ipaddrtable.parent = parent }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetParent() types.Entity { return ipaddrtable.parent }

func (ipaddrtable *RFC1213MIB_Ipaddrtable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Ipaddrtable_Ipaddrentry
// The addressing information for one of this
// entity's IP addresses.
type RFC1213MIB_Ipaddrtable_Ipaddrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address to which this entry's addressing
    // information pertains. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipadentaddr interface{}

    // The index value which uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of ifIndex. The
    // type is interface{} with range: -2147483648..2147483647.
    Ipadentifindex interface{}

    // The subnet mask associated with the IP address of this entry.  The value of
    // the mask is an IP address with all the network bits set to 1 and all the
    // hosts bits set to 0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetFilter() yfilter.YFilter { return ipaddrentry.YFilter }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) SetFilter(yf yfilter.YFilter) { ipaddrentry.YFilter = yf }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetGoName(yname string) string {
    if yname == "ipAdEntAddr" { return "Ipadentaddr" }
    if yname == "ipAdEntIfIndex" { return "Ipadentifindex" }
    if yname == "ipAdEntNetMask" { return "Ipadentnetmask" }
    if yname == "ipAdEntBcastAddr" { return "Ipadentbcastaddr" }
    if yname == "ipAdEntReasmMaxSize" { return "Ipadentreasmmaxsize" }
    return ""
}

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetSegmentPath() string {
    return "ipAddrEntry" + "[ipAdEntAddr='" + fmt.Sprintf("%v", ipaddrentry.Ipadentaddr) + "']"
}

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipAdEntAddr"] = ipaddrentry.Ipadentaddr
    leafs["ipAdEntIfIndex"] = ipaddrentry.Ipadentifindex
    leafs["ipAdEntNetMask"] = ipaddrentry.Ipadentnetmask
    leafs["ipAdEntBcastAddr"] = ipaddrentry.Ipadentbcastaddr
    leafs["ipAdEntReasmMaxSize"] = ipaddrentry.Ipadentreasmmaxsize
    return leafs
}

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetYangName() string { return "ipAddrEntry" }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) SetParent(parent types.Entity) { ipaddrentry.parent = parent }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetParent() types.Entity { return ipaddrentry.parent }

func (ipaddrentry *RFC1213MIB_Ipaddrtable_Ipaddrentry) GetParentYangName() string { return "ipAddrTable" }

// RFC1213MIB_Iproutetable
// This entity's IP Routing table.
type RFC1213MIB_Iproutetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A route to a particular destination. The type is slice of
    // RFC1213MIB_Iproutetable_Iprouteentry.
    Iprouteentry []RFC1213MIB_Iproutetable_Iprouteentry
}

func (iproutetable *RFC1213MIB_Iproutetable) GetFilter() yfilter.YFilter { return iproutetable.YFilter }

func (iproutetable *RFC1213MIB_Iproutetable) SetFilter(yf yfilter.YFilter) { iproutetable.YFilter = yf }

func (iproutetable *RFC1213MIB_Iproutetable) GetGoName(yname string) string {
    if yname == "ipRouteEntry" { return "Iprouteentry" }
    return ""
}

func (iproutetable *RFC1213MIB_Iproutetable) GetSegmentPath() string {
    return "ipRouteTable"
}

func (iproutetable *RFC1213MIB_Iproutetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipRouteEntry" {
        for _, c := range iproutetable.Iprouteentry {
            if iproutetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Iproutetable_Iprouteentry{}
        iproutetable.Iprouteentry = append(iproutetable.Iprouteentry, child)
        return &iproutetable.Iprouteentry[len(iproutetable.Iprouteentry)-1]
    }
    return nil
}

func (iproutetable *RFC1213MIB_Iproutetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iproutetable.Iprouteentry {
        children[iproutetable.Iprouteentry[i].GetSegmentPath()] = &iproutetable.Iprouteentry[i]
    }
    return children
}

func (iproutetable *RFC1213MIB_Iproutetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iproutetable *RFC1213MIB_Iproutetable) GetBundleName() string { return "cisco_ios_xe" }

func (iproutetable *RFC1213MIB_Iproutetable) GetYangName() string { return "ipRouteTable" }

func (iproutetable *RFC1213MIB_Iproutetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iproutetable *RFC1213MIB_Iproutetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iproutetable *RFC1213MIB_Iproutetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iproutetable *RFC1213MIB_Iproutetable) SetParent(parent types.Entity) { iproutetable.parent = parent }

func (iproutetable *RFC1213MIB_Iproutetable) GetParent() types.Entity { return iproutetable.parent }

func (iproutetable *RFC1213MIB_Iproutetable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Iproutetable_Iprouteentry
// A route to a particular destination.
type RFC1213MIB_Iproutetable_Iprouteentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The destination IP address of this route.  An
    // entry with a value of 0.0.0.0 is considered a default route.  Multiple
    // routes to a single destination can appear in the table, but access to such
    // multiple entries is dependent on the table- access mechanisms defined by
    // the network management protocol in use. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Iprouteinfo interface{}
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetFilter() yfilter.YFilter { return iprouteentry.YFilter }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) SetFilter(yf yfilter.YFilter) { iprouteentry.YFilter = yf }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetGoName(yname string) string {
    if yname == "ipRouteDest" { return "Iproutedest" }
    if yname == "ipRouteIfIndex" { return "Iprouteifindex" }
    if yname == "ipRouteMetric1" { return "Iproutemetric1" }
    if yname == "ipRouteMetric2" { return "Iproutemetric2" }
    if yname == "ipRouteMetric3" { return "Iproutemetric3" }
    if yname == "ipRouteMetric4" { return "Iproutemetric4" }
    if yname == "ipRouteNextHop" { return "Iproutenexthop" }
    if yname == "ipRouteType" { return "Iproutetype" }
    if yname == "ipRouteProto" { return "Iprouteproto" }
    if yname == "ipRouteAge" { return "Iprouteage" }
    if yname == "ipRouteMask" { return "Iproutemask" }
    if yname == "ipRouteMetric5" { return "Iproutemetric5" }
    if yname == "ipRouteInfo" { return "Iprouteinfo" }
    return ""
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetSegmentPath() string {
    return "ipRouteEntry" + "[ipRouteDest='" + fmt.Sprintf("%v", iprouteentry.Iproutedest) + "']"
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipRouteDest"] = iprouteentry.Iproutedest
    leafs["ipRouteIfIndex"] = iprouteentry.Iprouteifindex
    leafs["ipRouteMetric1"] = iprouteentry.Iproutemetric1
    leafs["ipRouteMetric2"] = iprouteentry.Iproutemetric2
    leafs["ipRouteMetric3"] = iprouteentry.Iproutemetric3
    leafs["ipRouteMetric4"] = iprouteentry.Iproutemetric4
    leafs["ipRouteNextHop"] = iprouteentry.Iproutenexthop
    leafs["ipRouteType"] = iprouteentry.Iproutetype
    leafs["ipRouteProto"] = iprouteentry.Iprouteproto
    leafs["ipRouteAge"] = iprouteentry.Iprouteage
    leafs["ipRouteMask"] = iprouteentry.Iproutemask
    leafs["ipRouteMetric5"] = iprouteentry.Iproutemetric5
    leafs["ipRouteInfo"] = iprouteentry.Iprouteinfo
    return leafs
}

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetBundleName() string { return "cisco_ios_xe" }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetYangName() string { return "ipRouteEntry" }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) SetParent(parent types.Entity) { iprouteentry.parent = parent }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetParent() types.Entity { return iprouteentry.parent }

func (iprouteentry *RFC1213MIB_Iproutetable_Iprouteentry) GetParentYangName() string { return "ipRouteTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains one IpAddress to `physical' address equivalence. The
    // type is slice of RFC1213MIB_Ipnettomediatable_Ipnettomediaentry.
    Ipnettomediaentry []RFC1213MIB_Ipnettomediatable_Ipnettomediaentry
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetFilter() yfilter.YFilter { return ipnettomediatable.YFilter }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) SetFilter(yf yfilter.YFilter) { ipnettomediatable.YFilter = yf }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetGoName(yname string) string {
    if yname == "ipNetToMediaEntry" { return "Ipnettomediaentry" }
    return ""
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetSegmentPath() string {
    return "ipNetToMediaTable"
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipNetToMediaEntry" {
        for _, c := range ipnettomediatable.Ipnettomediaentry {
            if ipnettomediatable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Ipnettomediatable_Ipnettomediaentry{}
        ipnettomediatable.Ipnettomediaentry = append(ipnettomediatable.Ipnettomediaentry, child)
        return &ipnettomediatable.Ipnettomediaentry[len(ipnettomediatable.Ipnettomediaentry)-1]
    }
    return nil
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipnettomediatable.Ipnettomediaentry {
        children[ipnettomediatable.Ipnettomediaentry[i].GetSegmentPath()] = &ipnettomediatable.Ipnettomediaentry[i]
    }
    return children
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetBundleName() string { return "cisco_ios_xe" }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetYangName() string { return "ipNetToMediaTable" }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) SetParent(parent types.Entity) { ipnettomediatable.parent = parent }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetParent() types.Entity { return ipnettomediatable.parent }

func (ipnettomediatable *RFC1213MIB_Ipnettomediatable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Ipnettomediatable_Ipnettomediaentry
// Each entry contains one IpAddress to `physical'
// address equivalence.
type RFC1213MIB_Ipnettomediatable_Ipnettomediaentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the same value of ifIndex. The type is
    // interface{} with range: -2147483648..2147483647.
    Ipnettomediaifindex interface{}

    // This attribute is a key. The IpAddress corresponding to the media-
    // dependent `physical' address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetFilter() yfilter.YFilter { return ipnettomediaentry.YFilter }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) SetFilter(yf yfilter.YFilter) { ipnettomediaentry.YFilter = yf }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetGoName(yname string) string {
    if yname == "ipNetToMediaIfIndex" { return "Ipnettomediaifindex" }
    if yname == "ipNetToMediaNetAddress" { return "Ipnettomedianetaddress" }
    if yname == "ipNetToMediaPhysAddress" { return "Ipnettomediaphysaddress" }
    if yname == "ipNetToMediaType" { return "Ipnettomediatype" }
    return ""
}

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetSegmentPath() string {
    return "ipNetToMediaEntry" + "[ipNetToMediaIfIndex='" + fmt.Sprintf("%v", ipnettomediaentry.Ipnettomediaifindex) + "']" + "[ipNetToMediaNetAddress='" + fmt.Sprintf("%v", ipnettomediaentry.Ipnettomedianetaddress) + "']"
}

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipNetToMediaIfIndex"] = ipnettomediaentry.Ipnettomediaifindex
    leafs["ipNetToMediaNetAddress"] = ipnettomediaentry.Ipnettomedianetaddress
    leafs["ipNetToMediaPhysAddress"] = ipnettomediaentry.Ipnettomediaphysaddress
    leafs["ipNetToMediaType"] = ipnettomediaentry.Ipnettomediatype
    return leafs
}

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetYangName() string { return "ipNetToMediaEntry" }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) SetParent(parent types.Entity) { ipnettomediaentry.parent = parent }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetParent() types.Entity { return ipnettomediaentry.parent }

func (ipnettomediaentry *RFC1213MIB_Ipnettomediatable_Ipnettomediaentry) GetParentYangName() string { return "ipNetToMediaTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular current TCP connection.  An object of this
    // type is transient, in that it ceases to exist when (or soon after) the
    // connection makes the transition to the CLOSED state. The type is slice of
    // RFC1213MIB_Tcpconntable_Tcpconnentry.
    Tcpconnentry []RFC1213MIB_Tcpconntable_Tcpconnentry
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetFilter() yfilter.YFilter { return tcpconntable.YFilter }

func (tcpconntable *RFC1213MIB_Tcpconntable) SetFilter(yf yfilter.YFilter) { tcpconntable.YFilter = yf }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetGoName(yname string) string {
    if yname == "tcpConnEntry" { return "Tcpconnentry" }
    return ""
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetSegmentPath() string {
    return "tcpConnTable"
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcpConnEntry" {
        for _, c := range tcpconntable.Tcpconnentry {
            if tcpconntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Tcpconntable_Tcpconnentry{}
        tcpconntable.Tcpconnentry = append(tcpconntable.Tcpconnentry, child)
        return &tcpconntable.Tcpconnentry[len(tcpconntable.Tcpconnentry)-1]
    }
    return nil
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tcpconntable.Tcpconnentry {
        children[tcpconntable.Tcpconnentry[i].GetSegmentPath()] = &tcpconntable.Tcpconnentry[i]
    }
    return children
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcpconntable *RFC1213MIB_Tcpconntable) GetBundleName() string { return "cisco_ios_xe" }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetYangName() string { return "tcpConnTable" }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpconntable *RFC1213MIB_Tcpconntable) SetParent(parent types.Entity) { tcpconntable.parent = parent }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetParent() types.Entity { return tcpconntable.parent }

func (tcpconntable *RFC1213MIB_Tcpconntable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Tcpconntable_Tcpconnentry
// Information about a particular current TCP
// connection.  An object of this type is transient,
// in that it ceases to exist when (or soon after)
// the connection makes the transition to the CLOSED
// state.
type RFC1213MIB_Tcpconntable_Tcpconnentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this TCP connection.  In
    // the case of a connection in the listen state which is willing to accept
    // connections for any IP interface associated with the node, the value
    // 0.0.0.0 is used. The type is string with pattern:
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

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetFilter() yfilter.YFilter { return tcpconnentry.YFilter }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) SetFilter(yf yfilter.YFilter) { tcpconnentry.YFilter = yf }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetGoName(yname string) string {
    if yname == "tcpConnLocalAddress" { return "Tcpconnlocaladdress" }
    if yname == "tcpConnLocalPort" { return "Tcpconnlocalport" }
    if yname == "tcpConnRemAddress" { return "Tcpconnremaddress" }
    if yname == "tcpConnRemPort" { return "Tcpconnremport" }
    if yname == "tcpConnState" { return "Tcpconnstate" }
    return ""
}

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetSegmentPath() string {
    return "tcpConnEntry" + "[tcpConnLocalAddress='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnlocaladdress) + "']" + "[tcpConnLocalPort='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnlocalport) + "']" + "[tcpConnRemAddress='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnremaddress) + "']" + "[tcpConnRemPort='" + fmt.Sprintf("%v", tcpconnentry.Tcpconnremport) + "']"
}

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcpConnLocalAddress"] = tcpconnentry.Tcpconnlocaladdress
    leafs["tcpConnLocalPort"] = tcpconnentry.Tcpconnlocalport
    leafs["tcpConnRemAddress"] = tcpconnentry.Tcpconnremaddress
    leafs["tcpConnRemPort"] = tcpconnentry.Tcpconnremport
    leafs["tcpConnState"] = tcpconnentry.Tcpconnstate
    return leafs
}

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetBundleName() string { return "cisco_ios_xe" }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetYangName() string { return "tcpConnEntry" }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) SetParent(parent types.Entity) { tcpconnentry.parent = parent }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetParent() types.Entity { return tcpconnentry.parent }

func (tcpconnentry *RFC1213MIB_Tcpconntable_Tcpconnentry) GetParentYangName() string { return "tcpConnTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular current UDP listener. The type is slice of
    // RFC1213MIB_Udptable_Udpentry.
    Udpentry []RFC1213MIB_Udptable_Udpentry
}

func (udptable *RFC1213MIB_Udptable) GetFilter() yfilter.YFilter { return udptable.YFilter }

func (udptable *RFC1213MIB_Udptable) SetFilter(yf yfilter.YFilter) { udptable.YFilter = yf }

func (udptable *RFC1213MIB_Udptable) GetGoName(yname string) string {
    if yname == "udpEntry" { return "Udpentry" }
    return ""
}

func (udptable *RFC1213MIB_Udptable) GetSegmentPath() string {
    return "udpTable"
}

func (udptable *RFC1213MIB_Udptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udpEntry" {
        for _, c := range udptable.Udpentry {
            if udptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Udptable_Udpentry{}
        udptable.Udpentry = append(udptable.Udpentry, child)
        return &udptable.Udpentry[len(udptable.Udpentry)-1]
    }
    return nil
}

func (udptable *RFC1213MIB_Udptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range udptable.Udpentry {
        children[udptable.Udpentry[i].GetSegmentPath()] = &udptable.Udpentry[i]
    }
    return children
}

func (udptable *RFC1213MIB_Udptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udptable *RFC1213MIB_Udptable) GetBundleName() string { return "cisco_ios_xe" }

func (udptable *RFC1213MIB_Udptable) GetYangName() string { return "udpTable" }

func (udptable *RFC1213MIB_Udptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udptable *RFC1213MIB_Udptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udptable *RFC1213MIB_Udptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udptable *RFC1213MIB_Udptable) SetParent(parent types.Entity) { udptable.parent = parent }

func (udptable *RFC1213MIB_Udptable) GetParent() types.Entity { return udptable.parent }

func (udptable *RFC1213MIB_Udptable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Udptable_Udpentry
// Information about a particular current UDP
// listener.
type RFC1213MIB_Udptable_Udpentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The local IP address for this UDP listener.  In
    // the case of a UDP listener which is willing to accept datagrams for any IP
    // interface associated with the node, the value 0.0.0.0 is used. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Udplocaladdress interface{}

    // This attribute is a key. The local port number for this UDP listener. The
    // type is interface{} with range: 0..65535.
    Udplocalport interface{}
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetFilter() yfilter.YFilter { return udpentry.YFilter }

func (udpentry *RFC1213MIB_Udptable_Udpentry) SetFilter(yf yfilter.YFilter) { udpentry.YFilter = yf }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetGoName(yname string) string {
    if yname == "udpLocalAddress" { return "Udplocaladdress" }
    if yname == "udpLocalPort" { return "Udplocalport" }
    return ""
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetSegmentPath() string {
    return "udpEntry" + "[udpLocalAddress='" + fmt.Sprintf("%v", udpentry.Udplocaladdress) + "']" + "[udpLocalPort='" + fmt.Sprintf("%v", udpentry.Udplocalport) + "']"
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udpLocalAddress"] = udpentry.Udplocaladdress
    leafs["udpLocalPort"] = udpentry.Udplocalport
    return leafs
}

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetBundleName() string { return "cisco_ios_xe" }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetYangName() string { return "udpEntry" }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udpentry *RFC1213MIB_Udptable_Udpentry) SetParent(parent types.Entity) { udpentry.parent = parent }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetParent() types.Entity { return udpentry.parent }

func (udpentry *RFC1213MIB_Udptable_Udpentry) GetParentYangName() string { return "udpTable" }

// RFC1213MIB_Egpneightable
// The EGP neighbor table.
type RFC1213MIB_Egpneightable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about this entity's relationship with a particular EGP
    // neighbor. The type is slice of RFC1213MIB_Egpneightable_Egpneighentry.
    Egpneighentry []RFC1213MIB_Egpneightable_Egpneighentry
}

func (egpneightable *RFC1213MIB_Egpneightable) GetFilter() yfilter.YFilter { return egpneightable.YFilter }

func (egpneightable *RFC1213MIB_Egpneightable) SetFilter(yf yfilter.YFilter) { egpneightable.YFilter = yf }

func (egpneightable *RFC1213MIB_Egpneightable) GetGoName(yname string) string {
    if yname == "egpNeighEntry" { return "Egpneighentry" }
    return ""
}

func (egpneightable *RFC1213MIB_Egpneightable) GetSegmentPath() string {
    return "egpNeighTable"
}

func (egpneightable *RFC1213MIB_Egpneightable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "egpNeighEntry" {
        for _, c := range egpneightable.Egpneighentry {
            if egpneightable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1213MIB_Egpneightable_Egpneighentry{}
        egpneightable.Egpneighentry = append(egpneightable.Egpneighentry, child)
        return &egpneightable.Egpneighentry[len(egpneightable.Egpneighentry)-1]
    }
    return nil
}

func (egpneightable *RFC1213MIB_Egpneightable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range egpneightable.Egpneighentry {
        children[egpneightable.Egpneighentry[i].GetSegmentPath()] = &egpneightable.Egpneighentry[i]
    }
    return children
}

func (egpneightable *RFC1213MIB_Egpneightable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (egpneightable *RFC1213MIB_Egpneightable) GetBundleName() string { return "cisco_ios_xe" }

func (egpneightable *RFC1213MIB_Egpneightable) GetYangName() string { return "egpNeighTable" }

func (egpneightable *RFC1213MIB_Egpneightable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (egpneightable *RFC1213MIB_Egpneightable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (egpneightable *RFC1213MIB_Egpneightable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (egpneightable *RFC1213MIB_Egpneightable) SetParent(parent types.Entity) { egpneightable.parent = parent }

func (egpneightable *RFC1213MIB_Egpneightable) GetParent() types.Entity { return egpneightable.parent }

func (egpneightable *RFC1213MIB_Egpneightable) GetParentYangName() string { return "RFC1213-MIB" }

// RFC1213MIB_Egpneightable_Egpneighentry
// Information about this entity's relationship with
// a particular EGP neighbor.
type RFC1213MIB_Egpneightable_Egpneighentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this entry's EGP neighbor. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetFilter() yfilter.YFilter { return egpneighentry.YFilter }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) SetFilter(yf yfilter.YFilter) { egpneighentry.YFilter = yf }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetGoName(yname string) string {
    if yname == "egpNeighAddr" { return "Egpneighaddr" }
    if yname == "egpNeighState" { return "Egpneighstate" }
    if yname == "egpNeighAs" { return "Egpneighas" }
    if yname == "egpNeighInMsgs" { return "Egpneighinmsgs" }
    if yname == "egpNeighInErrs" { return "Egpneighinerrs" }
    if yname == "egpNeighOutMsgs" { return "Egpneighoutmsgs" }
    if yname == "egpNeighOutErrs" { return "Egpneighouterrs" }
    if yname == "egpNeighInErrMsgs" { return "Egpneighinerrmsgs" }
    if yname == "egpNeighOutErrMsgs" { return "Egpneighouterrmsgs" }
    if yname == "egpNeighStateUps" { return "Egpneighstateups" }
    if yname == "egpNeighStateDowns" { return "Egpneighstatedowns" }
    if yname == "egpNeighIntervalHello" { return "Egpneighintervalhello" }
    if yname == "egpNeighIntervalPoll" { return "Egpneighintervalpoll" }
    if yname == "egpNeighMode" { return "Egpneighmode" }
    if yname == "egpNeighEventTrigger" { return "Egpneigheventtrigger" }
    return ""
}

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetSegmentPath() string {
    return "egpNeighEntry" + "[egpNeighAddr='" + fmt.Sprintf("%v", egpneighentry.Egpneighaddr) + "']"
}

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["egpNeighAddr"] = egpneighentry.Egpneighaddr
    leafs["egpNeighState"] = egpneighentry.Egpneighstate
    leafs["egpNeighAs"] = egpneighentry.Egpneighas
    leafs["egpNeighInMsgs"] = egpneighentry.Egpneighinmsgs
    leafs["egpNeighInErrs"] = egpneighentry.Egpneighinerrs
    leafs["egpNeighOutMsgs"] = egpneighentry.Egpneighoutmsgs
    leafs["egpNeighOutErrs"] = egpneighentry.Egpneighouterrs
    leafs["egpNeighInErrMsgs"] = egpneighentry.Egpneighinerrmsgs
    leafs["egpNeighOutErrMsgs"] = egpneighentry.Egpneighouterrmsgs
    leafs["egpNeighStateUps"] = egpneighentry.Egpneighstateups
    leafs["egpNeighStateDowns"] = egpneighentry.Egpneighstatedowns
    leafs["egpNeighIntervalHello"] = egpneighentry.Egpneighintervalhello
    leafs["egpNeighIntervalPoll"] = egpneighentry.Egpneighintervalpoll
    leafs["egpNeighMode"] = egpneighentry.Egpneighmode
    leafs["egpNeighEventTrigger"] = egpneighentry.Egpneigheventtrigger
    return leafs
}

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetBundleName() string { return "cisco_ios_xe" }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetYangName() string { return "egpNeighEntry" }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) SetParent(parent types.Entity) { egpneighentry.parent = parent }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetParent() types.Entity { return egpneighentry.parent }

func (egpneighentry *RFC1213MIB_Egpneightable_Egpneighentry) GetParentYangName() string { return "egpNeighTable" }

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

