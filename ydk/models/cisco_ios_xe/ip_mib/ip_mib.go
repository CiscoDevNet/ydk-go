// The MIB module for managing IP and ICMP implementations, but
// excluding their management of IP routes.
// 
// Copyright (C) The Internet Society (2006).  This version of
// this MIB module is part of RFC 4293; see the RFC itself for
// full legal notices.
package ip_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:IP-MIB IP-MIB}", reflect.TypeOf(IPMIB{}))
    ydk.RegisterEntity("IP-MIB:IP-MIB", reflect.TypeOf(IPMIB{}))
}

// IpAddressPrefixOriginTC represents prefix was found.
type IpAddressPrefixOriginTC string

const (
    IpAddressPrefixOriginTC_other IpAddressPrefixOriginTC = "other"

    IpAddressPrefixOriginTC_manual IpAddressPrefixOriginTC = "manual"

    IpAddressPrefixOriginTC_wellknown IpAddressPrefixOriginTC = "wellknown"

    IpAddressPrefixOriginTC_dhcp IpAddressPrefixOriginTC = "dhcp"

    IpAddressPrefixOriginTC_routeradv IpAddressPrefixOriginTC = "routeradv"
)

// IpAddressOriginTC represents 3041 privacy address.
type IpAddressOriginTC string

const (
    IpAddressOriginTC_other IpAddressOriginTC = "other"

    IpAddressOriginTC_manual IpAddressOriginTC = "manual"

    IpAddressOriginTC_dhcp IpAddressOriginTC = "dhcp"

    IpAddressOriginTC_linklayer IpAddressOriginTC = "linklayer"

    IpAddressOriginTC_random IpAddressOriginTC = "random"
)

// IpAddressStatusTC represents always preferred(1).
type IpAddressStatusTC string

const (
    IpAddressStatusTC_preferred IpAddressStatusTC = "preferred"

    IpAddressStatusTC_deprecated IpAddressStatusTC = "deprecated"

    IpAddressStatusTC_invalid IpAddressStatusTC = "invalid"

    IpAddressStatusTC_inaccessible IpAddressStatusTC = "inaccessible"

    IpAddressStatusTC_unknown IpAddressStatusTC = "unknown"

    IpAddressStatusTC_tentative IpAddressStatusTC = "tentative"

    IpAddressStatusTC_duplicate IpAddressStatusTC = "duplicate"

    IpAddressStatusTC_optimistic IpAddressStatusTC = "optimistic"
)

// IPMIB
type IPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ip IPMIB_Ip

    
    Iptrafficstats IPMIB_Iptrafficstats

    
    Icmp IPMIB_Icmp

    // The table of addressing information relevant to this entity's IPv4
    // addresses.  This table has been deprecated, as a new IP version-neutral
    // table has been added.  It is loosely replaced by the ipAddressTable
    // although several objects that weren't deemed useful weren't carried forward
    // while another (ipAdEntReasmMaxSize) was moved to the ipv4InterfaceTable.
    Ipaddrtable IPMIB_Ipaddrtable

    // The IPv4 Address Translation table used for mapping from IPv4 addresses to
    // physical addresses.  This table has been deprecated, as a new IP
    // version-neutral table has been added.  It is loosely replaced by the
    // ipNetToPhysicalTable.
    Ipnettomediatable IPMIB_Ipnettomediatable

    // The table containing per-interface IPv4-specific information.
    Ipv4Interfacetable IPMIB_Ipv4Interfacetable

    // The table containing per-interface IPv6-specific information.
    Ipv6Interfacetable IPMIB_Ipv6Interfacetable

    // The table containing system wide, IP version specific traffic statistics. 
    // This table and the ipIfStatsTable contain similar objects whose difference
    // is in their granularity.  Where this table contains system wide traffic
    // statistics, the ipIfStatsTable contains the same statistics but counted on
    // a per-interface basis.
    Ipsystemstatstable IPMIB_Ipsystemstatstable

    // The table containing per-interface traffic statistics.  This table and the
    // ipSystemStatsTable contain similar objects whose difference is in their
    // granularity.  Where this table contains per-interface statistics, the
    // ipSystemStatsTable contains the same statistics, but counted on a system
    // wide basis.
    Ipifstatstable IPMIB_Ipifstatstable

    // This table allows the user to determine the source of an IP address or set
    // of IP addresses, and allows other tables to share the information via
    // pointer rather than by copying.  For example, when the node configures both
    // a unicast and anycast address for a prefix, the ipAddressPrefix objects for
    // those addresses will point to a single row in this table.  This table
    // primarily provides support for IPv6 prefixes, and several of the objects
    // are less meaningful for IPv4.  The table continues to allow IPv4 addresses
    // to allow future flexibility.  In order to promote a common configuration,
    // this document includes suggestions for default values for IPv4 prefixes. 
    // Each of these values may be overridden if an object is meaningful to the
    // node.  All prefixes used by this entity should be included in this table
    // independent of how the entity learned the prefix. (This table isn't limited
    // to prefixes learned from router   advertisements.).
    Ipaddressprefixtable IPMIB_Ipaddressprefixtable

    // This table contains addressing information relevant to the entity's
    // interfaces.  This table does not contain multicast address information.
    // Tables for such information should be contained in multicast specific MIBs,
    // such as RFC 3019.  While this table is writable, the user will note that
    // several objects, such as ipAddressOrigin, are not.  The intention in
    // allowing a user to write to this table is to allow them to add or remove
    // any entry that isn't   permanent.  The user should be allowed to modify
    // objects and entries when that would not cause inconsistencies within the
    // table.  Allowing write access to objects, such as ipAddressOrigin, could
    // allow a user to insert an entry and then label it incorrectly.  Note well:
    // When including IPv6 link-local addresses in this table, the entry must use
    // an InetAddressType of 'ipv6z' in order to differentiate between the
    // possible interfaces.
    Ipaddresstable IPMIB_Ipaddresstable

    // The IP Address Translation table used for mapping from IP addresses to
    // physical addresses.  The Address Translation tables contain the IP address
    // to 'physical' address equivalences.  Some interfaces do not use translation
    // tables for determining address equivalences (e.g., DDN-X.25 has an
    // algorithmic method); if all interfaces are of this type, then the Address
    // Translation table is empty, i.e., has zero entries.  While many protocols
    // may be used to populate this table, ARP and Neighbor Discovery are the most
    // likely options.
    Ipnettophysicaltable IPMIB_Ipnettophysicaltable

    // The table used to describe IPv6 unicast and multicast scope zones.  For
    // those objects that have names rather than numbers, the names were chosen to
    // coincide with the names used in the IPv6 address architecture document. .
    Ipv6Scopezoneindextable IPMIB_Ipv6Scopezoneindextable

    // The table used to describe the default routers known to this   entity.
    Ipdefaultroutertable IPMIB_Ipdefaultroutertable

    // The table containing information used to construct router advertisements.
    Ipv6Routeradverttable IPMIB_Ipv6Routeradverttable

    // The table of generic system-wide ICMP counters.
    Icmpstatstable IPMIB_Icmpstatstable

    // The table of system-wide per-version, per-message type ICMP counters.
    Icmpmsgstatstable IPMIB_Icmpmsgstatstable
}

func (iPMIB *IPMIB) GetFilter() yfilter.YFilter { return iPMIB.YFilter }

func (iPMIB *IPMIB) SetFilter(yf yfilter.YFilter) { iPMIB.YFilter = yf }

func (iPMIB *IPMIB) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "ipTrafficStats" { return "Iptrafficstats" }
    if yname == "icmp" { return "Icmp" }
    if yname == "ipAddrTable" { return "Ipaddrtable" }
    if yname == "ipNetToMediaTable" { return "Ipnettomediatable" }
    if yname == "ipv4InterfaceTable" { return "Ipv4Interfacetable" }
    if yname == "ipv6InterfaceTable" { return "Ipv6Interfacetable" }
    if yname == "ipSystemStatsTable" { return "Ipsystemstatstable" }
    if yname == "ipIfStatsTable" { return "Ipifstatstable" }
    if yname == "ipAddressPrefixTable" { return "Ipaddressprefixtable" }
    if yname == "ipAddressTable" { return "Ipaddresstable" }
    if yname == "ipNetToPhysicalTable" { return "Ipnettophysicaltable" }
    if yname == "ipv6ScopeZoneIndexTable" { return "Ipv6Scopezoneindextable" }
    if yname == "ipDefaultRouterTable" { return "Ipdefaultroutertable" }
    if yname == "ipv6RouterAdvertTable" { return "Ipv6Routeradverttable" }
    if yname == "icmpStatsTable" { return "Icmpstatstable" }
    if yname == "icmpMsgStatsTable" { return "Icmpmsgstatstable" }
    return ""
}

func (iPMIB *IPMIB) GetSegmentPath() string {
    return "IP-MIB:IP-MIB"
}

func (iPMIB *IPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip" {
        return &iPMIB.Ip
    }
    if childYangName == "ipTrafficStats" {
        return &iPMIB.Iptrafficstats
    }
    if childYangName == "icmp" {
        return &iPMIB.Icmp
    }
    if childYangName == "ipAddrTable" {
        return &iPMIB.Ipaddrtable
    }
    if childYangName == "ipNetToMediaTable" {
        return &iPMIB.Ipnettomediatable
    }
    if childYangName == "ipv4InterfaceTable" {
        return &iPMIB.Ipv4Interfacetable
    }
    if childYangName == "ipv6InterfaceTable" {
        return &iPMIB.Ipv6Interfacetable
    }
    if childYangName == "ipSystemStatsTable" {
        return &iPMIB.Ipsystemstatstable
    }
    if childYangName == "ipIfStatsTable" {
        return &iPMIB.Ipifstatstable
    }
    if childYangName == "ipAddressPrefixTable" {
        return &iPMIB.Ipaddressprefixtable
    }
    if childYangName == "ipAddressTable" {
        return &iPMIB.Ipaddresstable
    }
    if childYangName == "ipNetToPhysicalTable" {
        return &iPMIB.Ipnettophysicaltable
    }
    if childYangName == "ipv6ScopeZoneIndexTable" {
        return &iPMIB.Ipv6Scopezoneindextable
    }
    if childYangName == "ipDefaultRouterTable" {
        return &iPMIB.Ipdefaultroutertable
    }
    if childYangName == "ipv6RouterAdvertTable" {
        return &iPMIB.Ipv6Routeradverttable
    }
    if childYangName == "icmpStatsTable" {
        return &iPMIB.Icmpstatstable
    }
    if childYangName == "icmpMsgStatsTable" {
        return &iPMIB.Icmpmsgstatstable
    }
    return nil
}

func (iPMIB *IPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ip"] = &iPMIB.Ip
    children["ipTrafficStats"] = &iPMIB.Iptrafficstats
    children["icmp"] = &iPMIB.Icmp
    children["ipAddrTable"] = &iPMIB.Ipaddrtable
    children["ipNetToMediaTable"] = &iPMIB.Ipnettomediatable
    children["ipv4InterfaceTable"] = &iPMIB.Ipv4Interfacetable
    children["ipv6InterfaceTable"] = &iPMIB.Ipv6Interfacetable
    children["ipSystemStatsTable"] = &iPMIB.Ipsystemstatstable
    children["ipIfStatsTable"] = &iPMIB.Ipifstatstable
    children["ipAddressPrefixTable"] = &iPMIB.Ipaddressprefixtable
    children["ipAddressTable"] = &iPMIB.Ipaddresstable
    children["ipNetToPhysicalTable"] = &iPMIB.Ipnettophysicaltable
    children["ipv6ScopeZoneIndexTable"] = &iPMIB.Ipv6Scopezoneindextable
    children["ipDefaultRouterTable"] = &iPMIB.Ipdefaultroutertable
    children["ipv6RouterAdvertTable"] = &iPMIB.Ipv6Routeradverttable
    children["icmpStatsTable"] = &iPMIB.Icmpstatstable
    children["icmpMsgStatsTable"] = &iPMIB.Icmpmsgstatstable
    return children
}

func (iPMIB *IPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iPMIB *IPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iPMIB *IPMIB) GetYangName() string { return "IP-MIB" }

func (iPMIB *IPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iPMIB *IPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iPMIB *IPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iPMIB *IPMIB) SetParent(parent types.Entity) { iPMIB.parent = parent }

func (iPMIB *IPMIB) GetParent() types.Entity { return iPMIB.parent }

func (iPMIB *IPMIB) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ip
type IPMIB_Ip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The indication of whether this entity is acting as an IPv4 router in
    // respect to the forwarding of datagrams received by, but not addressed to,
    // this entity.  IPv4 routers forward datagrams.  IPv4 hosts do not (except
    // those source-routed via the host).  When this object is written, the entity
    // should save the change to non-volatile storage and restore the object from
    // non-volatile storage upon re-initialization of the system. Note: a stronger
    // requirement is not used because this object was previously defined. The
    // type is Ipforwarding.
    Ipforwarding interface{}

    // The default value inserted into the Time-To-Live field of the IPv4 header
    // of datagrams originated at this entity, whenever a TTL value is not
    // supplied by the transport layer   protocol.  When this object is written,
    // the entity should save the change to non-volatile storage and restore the
    // object from non-volatile storage upon re-initialization of the system.
    // Note: a stronger requirement is not used because this object was previously
    // defined. The type is interface{} with range: 1..255.
    Ipdefaultttl interface{}

    // The total number of input datagrams received from interfaces, including
    // those received in error.  This object has been deprecated, as a new IP
    // version-neutral   table has been added.  It is loosely replaced by
    // ipSystemStatsInRecieves. The type is interface{} with range: 0..4294967295.
    Ipinreceives interface{}

    // The number of input datagrams discarded due to errors in their IPv4
    // headers, including bad checksums, version number mismatch, other format
    // errors, time-to-live exceeded, errors discovered in processing their IPv4
    // options, etc.  This object has been deprecated as a new IP version-neutral
    // table has been added.  It is loosely replaced by ipSystemStatsInHdrErrors.
    // The type is interface{} with range: 0..4294967295.
    Ipinhdrerrors interface{}

    // The number of input datagrams discarded because the IPv4 address in their
    // IPv4 header's destination field was not a valid address to be received at
    // this entity.  This count includes invalid addresses (e.g., 0.0.0.0) and
    // addresses of unsupported Classes (e.g., Class E).  For entities which are
    // not IPv4 routers, and therefore do not forward datagrams, this counter
    // includes datagrams discarded because the destination address was not a
    // local address.  This object has been deprecated, as a new IP
    // version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsInAddrErrors. The type is interface{} with range:
    // 0..4294967295.
    Ipinaddrerrors interface{}

    // The number of input datagrams for which this entity was not their final
    // IPv4 destination, as a result of which an attempt was made to find a route
    // to forward them to that final destination.  In entities which do not act as
    // IPv4 routers, this counter will include only those packets which   were
    // Source-Routed via this entity, and the Source-Route option processing was
    // successful.  This object has been deprecated, as a new IP version-neutral
    // table has been added.  It is loosely replaced by
    // ipSystemStatsInForwDatagrams. The type is interface{} with range:
    // 0..4294967295.
    Ipforwdatagrams interface{}

    // The number of locally-addressed datagrams received successfully but
    // discarded because of an unknown or unsupported protocol.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by ipSystemStatsInUnknownProtos. The type is interface{}
    // with range: 0..4294967295.
    Ipinunknownprotos interface{}

    // The number of input IPv4 datagrams for which no problems were encountered
    // to prevent their continued processing, but which were discarded (e.g., for
    // lack of buffer space).  Note that this counter does not include any
    // datagrams discarded while awaiting re-assembly.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by ipSystemStatsInDiscards. The type is interface{} with
    // range: 0..4294967295.
    Ipindiscards interface{}

    // The total number of input datagrams successfully delivered to IPv4
    // user-protocols (including ICMP).  This object has been deprecated as a new
    // IP version neutral table has been added.  It is loosely replaced by  
    // ipSystemStatsIndelivers. The type is interface{} with range: 0..4294967295.
    Ipindelivers interface{}

    // The total number of IPv4 datagrams which local IPv4 user protocols
    // (including ICMP) supplied to IPv4 in requests for transmission.  Note that
    // this counter does not include any datagrams counted in ipForwDatagrams. 
    // This object has been deprecated, as a new IP version-neutral table has been
    // added.  It is loosely replaced by ipSystemStatsOutRequests. The type is
    // interface{} with range: 0..4294967295.
    Ipoutrequests interface{}

    // The number of output IPv4 datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but which were discarded
    // (e.g., for lack of buffer space).  Note that this counter would include
    // datagrams counted in ipForwDatagrams if any such packets met this
    // (discretionary) discard criterion.  This object has been deprecated, as a
    // new IP version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsOutDiscards. The type is interface{} with range:
    // 0..4294967295.
    Ipoutdiscards interface{}

    // The number of IPv4 datagrams discarded because no route could be found to
    // transmit them to their destination.  Note that this counter includes any
    // packets counted in ipForwDatagrams which meet this `no-route' criterion. 
    // Note that this includes any datagrams which a host cannot route because all
    // of its default routers are down.  This object has been deprecated, as a new
    // IP version-neutral   table has been added.  It is loosely replaced by
    // ipSystemStatsOutNoRoutes. The type is interface{} with range:
    // 0..4294967295.
    Ipoutnoroutes interface{}

    // The maximum number of seconds that received fragments are held while they
    // are awaiting reassembly at this entity. The type is interface{} with range:
    // -2147483648..2147483647. Units are seconds.
    Ipreasmtimeout interface{}

    // The number of IPv4 fragments received which needed to be reassembled at
    // this entity.  This object has been deprecated, as a new IP version-neutral
    // table has been added.  It is loosely replaced by ipSystemStatsReasmReqds.
    // The type is interface{} with range: 0..4294967295.
    Ipreasmreqds interface{}

    // The number of IPv4 datagrams successfully re-assembled.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by ipSystemStatsReasmOKs. The type is interface{} with
    // range: 0..4294967295.
    Ipreasmoks interface{}

    // The number of failures detected by the IPv4 re-assembly algorithm (for
    // whatever reason: timed out, errors, etc). Note that this is not necessarily
    // a count of discarded IPv4 fragments since some algorithms (notably the
    // algorithm in RFC 815) can lose track of the number of fragments by
    // combining them as they are received.  This object has been deprecated, as a
    // new IP version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsReasmFails. The type is interface{} with range: 0..4294967295.
    Ipreasmfails interface{}

    // The number of IPv4 datagrams that have been successfully fragmented at this
    // entity.  This object has been deprecated, as a new IP version-neutral table
    // has been added.  It is loosely replaced by ipSystemStatsOutFragOKs. The
    // type is interface{} with range: 0..4294967295.
    Ipfragoks interface{}

    // The number of IPv4 datagrams that have been discarded because they needed
    // to be fragmented at this entity but could not be, e.g., because their Don't
    // Fragment flag was set.  This object has been deprecated, as a new IP
    // version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsOutFragFails. The type is interface{} with range:
    // 0..4294967295.
    Ipfragfails interface{}

    // The number of IPv4 datagram fragments that have been generated as a result
    // of fragmentation at this entity.  This object has been deprecated as a new
    // IP version neutral table has been added.  It is loosely replaced by
    // ipSystemStatsOutFragCreates. The type is interface{} with range:
    // 0..4294967295.
    Ipfragcreates interface{}

    // The number of routing entries which were chosen to be discarded even though
    // they are valid.  One possible reason for discarding such an entry could be
    // to free-up buffer space for other routing entries.   This object was
    // defined in pre-IPv6 versions of the IP MIB. It was implicitly IPv4 only,
    // but the original specifications did not indicate this protocol restriction.
    // In order to clarify the specifications, this object has been deprecated and
    // a similar, but more thoroughly clarified, object has been added to the
    // IP-FORWARD-MIB. The type is interface{} with range: 0..4294967295.
    Iproutingdiscards interface{}

    // The indication of whether this entity is acting as an IPv6 router on any
    // interface in respect to the forwarding of datagrams received by, but not
    // addressed to, this entity. IPv6 routers forward datagrams.  IPv6 hosts do
    // not (except those source-routed via the host).  When this object is
    // written, the entity SHOULD save the change to non-volatile storage and
    // restore the object from non-volatile storage upon re-initialization of the
    // system. The type is Ipv6Ipforwarding.
    Ipv6Ipforwarding interface{}

    // The default value inserted into the Hop Limit field of the IPv6 header of
    // datagrams originated at this entity whenever a Hop Limit value is not
    // supplied by the transport layer protocol.  When this object is written, the
    // entity SHOULD save the change to non-volatile storage and restore the
    // object from non-volatile storage upon re-initialization of the system. The
    // type is interface{} with range: 0..255.
    Ipv6Ipdefaulthoplimit interface{}

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipv4InterfaceTable was added or deleted, or when an
    // ipv4InterfaceReasmMaxSize or an ipv4InterfaceEnableStatus object was
    // modified.  If new objects are added to the ipv4InterfaceTable that require
    // the ipv4InterfaceTableLastChange to be updated when they are modified, they
    // must specify that requirement in their description clause. The type is
    // interface{} with range: 0..4294967295.
    Ipv4Interfacetablelastchange interface{}

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipv6InterfaceTable was added or deleted or when an
    // ipv6InterfaceReasmMaxSize, ipv6InterfaceIdentifier,
    // ipv6InterfaceEnableStatus, ipv6InterfaceReachableTime,
    // ipv6InterfaceRetransmitTime, or ipv6InterfaceForwarding object was
    // modified.  If new objects are added to the ipv6InterfaceTable that require
    // the ipv6InterfaceTableLastChange to be updated when they are modified, they
    // must specify that requirement in their description clause. The type is
    // interface{} with range: 0..4294967295.
    Ipv6Interfacetablelastchange interface{}

    // An advisory lock used to allow cooperating SNMP managers to coordinate
    // their use of the set operation in creating or modifying rows within this
    // table.  In order to use this lock to coordinate the use of set operations,
    // managers should first retrieve ipAddressTableSpinLock.  They should then
    // determine the appropriate row to create or modify.  Finally, they should
    // issue the appropriate set command, including the retrieved value of
    // ipAddressSpinLock.  If another manager has altered the table in the
    // meantime, then the value of ipAddressSpinLock will have changed, and the
    // creation will fail as it will be specifying an incorrect value for
    // ipAddressSpinLock.  It is suggested, but not required, that the
    // ipAddressSpinLock be the first var bind for each set of objects
    // representing a 'row' in a PDU. The type is interface{} with range:
    // 0..2147483647.
    Ipaddressspinlock interface{}

    // An advisory lock used to allow cooperating SNMP managers to coordinate
    // their use of the set operation in creating or modifying rows within this
    // table.  In order to use this lock to coordinate the use of set operations,
    // managers should first retrieve ipv6RouterAdvertSpinLock.  They should then
    // determine the appropriate row to create or modify.  Finally, they should
    // issue the appropriate set command including the retrieved value of
    // ipv6RouterAdvertSpinLock.  If another manager has altered the table in the
    // meantime, then the value of ipv6RouterAdvertSpinLock will have changed and
    // the creation will fail as it will be specifying an incorrect value for
    // ipv6RouterAdvertSpinLock.  It is suggested, but not required, that the
    // ipv6RouterAdvertSpinLock be the first var bind for each set of objects
    // representing a 'row' in a PDU. The type is interface{} with range:
    // 0..2147483647.
    Ipv6Routeradvertspinlock interface{}
}

func (ip *IPMIB_Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *IPMIB_Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *IPMIB_Ip) GetGoName(yname string) string {
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
    if yname == "ipv6IpForwarding" { return "Ipv6Ipforwarding" }
    if yname == "ipv6IpDefaultHopLimit" { return "Ipv6Ipdefaulthoplimit" }
    if yname == "ipv4InterfaceTableLastChange" { return "Ipv4Interfacetablelastchange" }
    if yname == "ipv6InterfaceTableLastChange" { return "Ipv6Interfacetablelastchange" }
    if yname == "ipAddressSpinLock" { return "Ipaddressspinlock" }
    if yname == "ipv6RouterAdvertSpinLock" { return "Ipv6Routeradvertspinlock" }
    return ""
}

func (ip *IPMIB_Ip) GetSegmentPath() string {
    return "ip"
}

func (ip *IPMIB_Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ip *IPMIB_Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ip *IPMIB_Ip) GetLeafs() map[string]interface{} {
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
    leafs["ipv6IpForwarding"] = ip.Ipv6Ipforwarding
    leafs["ipv6IpDefaultHopLimit"] = ip.Ipv6Ipdefaulthoplimit
    leafs["ipv4InterfaceTableLastChange"] = ip.Ipv4Interfacetablelastchange
    leafs["ipv6InterfaceTableLastChange"] = ip.Ipv6Interfacetablelastchange
    leafs["ipAddressSpinLock"] = ip.Ipaddressspinlock
    leafs["ipv6RouterAdvertSpinLock"] = ip.Ipv6Routeradvertspinlock
    return leafs
}

func (ip *IPMIB_Ip) GetBundleName() string { return "cisco_ios_xe" }

func (ip *IPMIB_Ip) GetYangName() string { return "ip" }

func (ip *IPMIB_Ip) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ip *IPMIB_Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ip *IPMIB_Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ip *IPMIB_Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *IPMIB_Ip) GetParent() types.Entity { return ip.parent }

func (ip *IPMIB_Ip) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ip_Ipforwarding represents was previously defined.
type IPMIB_Ip_Ipforwarding string

const (
    IPMIB_Ip_Ipforwarding_forwarding IPMIB_Ip_Ipforwarding = "forwarding"

    IPMIB_Ip_Ipforwarding_notForwarding IPMIB_Ip_Ipforwarding = "notForwarding"
)

// IPMIB_Ip_Ipv6Ipforwarding represents non-volatile storage upon re-initialization of the system.
type IPMIB_Ip_Ipv6Ipforwarding string

const (
    IPMIB_Ip_Ipv6Ipforwarding_forwarding IPMIB_Ip_Ipv6Ipforwarding = "forwarding"

    IPMIB_Ip_Ipv6Ipforwarding_notForwarding IPMIB_Ip_Ipv6Ipforwarding = "notForwarding"
)

// IPMIB_Iptrafficstats
type IPMIB_Iptrafficstats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipIfStatsTable was added or deleted.  If new objects are added to the
    // ipIfStatsTable that require the ipIfStatsTableLastChange to be updated when
    // they are modified, they must specify that requirement in their description
    // clause. The type is interface{} with range: 0..4294967295.
    Ipifstatstablelastchange interface{}
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetFilter() yfilter.YFilter { return iptrafficstats.YFilter }

func (iptrafficstats *IPMIB_Iptrafficstats) SetFilter(yf yfilter.YFilter) { iptrafficstats.YFilter = yf }

func (iptrafficstats *IPMIB_Iptrafficstats) GetGoName(yname string) string {
    if yname == "ipIfStatsTableLastChange" { return "Ipifstatstablelastchange" }
    return ""
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetSegmentPath() string {
    return "ipTrafficStats"
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipIfStatsTableLastChange"] = iptrafficstats.Ipifstatstablelastchange
    return leafs
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetBundleName() string { return "cisco_ios_xe" }

func (iptrafficstats *IPMIB_Iptrafficstats) GetYangName() string { return "ipTrafficStats" }

func (iptrafficstats *IPMIB_Iptrafficstats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iptrafficstats *IPMIB_Iptrafficstats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iptrafficstats *IPMIB_Iptrafficstats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iptrafficstats *IPMIB_Iptrafficstats) SetParent(parent types.Entity) { iptrafficstats.parent = parent }

func (iptrafficstats *IPMIB_Iptrafficstats) GetParent() types.Entity { return iptrafficstats.parent }

func (iptrafficstats *IPMIB_Iptrafficstats) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Icmp
type IPMIB_Icmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total number of ICMP messages which the entity received. Note that this
    // counter includes all those counted by icmpInErrors.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by icmpStatsInMsgs. The type is interface{} with range:
    // 0..4294967295.
    Icmpinmsgs interface{}

    // The number of ICMP messages which the entity received but determined as
    // having ICMP-specific errors (bad ICMP checksums, bad length, etc.).  This
    // object has been deprecated, as a new IP version-neutral table has been
    // added.  It is loosely replaced by icmpStatsInErrors. The type is
    // interface{} with range: 0..4294967295.
    Icmpinerrors interface{}

    // The number of ICMP Destination Unreachable messages received.  This object
    // has been deprecated, as a new IP version-neutral table has been added.  It
    // is loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpindestunreachs interface{}

    // The number of ICMP Time Exceeded messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpintimeexcds interface{}

    // The number of ICMP Parameter Problem messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinparmprobs interface{}

    // The number of ICMP Source Quench messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinsrcquenchs interface{}

    // The number of ICMP Redirect messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinredirects interface{}

    // The number of ICMP Echo (request) messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinechos interface{}

    // The number of ICMP Echo Reply messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinechoreps interface{}

    // The number of ICMP Timestamp (request) messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpintimestamps interface{}

    // The number of ICMP Timestamp Reply messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpintimestampreps interface{}

    // The number of ICMP Address Mask Request messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinaddrmasks interface{}

    // The number of ICMP Address Mask Reply messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpinaddrmaskreps interface{}

    // The total number of ICMP messages which this entity attempted to send. 
    // Note that this counter includes all those counted by icmpOutErrors.  This
    // object has been deprecated, as a new IP version-neutral table has been
    // added.  It is loosely replaced by icmpStatsOutMsgs. The type is interface{}
    // with range: 0..4294967295.
    Icmpoutmsgs interface{}

    // The number of ICMP messages which this entity did not send due to problems
    // discovered within ICMP, such as a lack of buffers.  This value should not
    // include errors discovered outside the ICMP layer, such as the inability of
    // IP to route the resultant datagram.  In some implementations, there may be
    // no types of error which contribute to this counter's value.  This object
    // has been deprecated, as a new IP version-neutral table has been added.  It
    // is loosely replaced by icmpStatsOutErrors. The type is interface{} with
    // range: 0..4294967295.
    Icmpouterrors interface{}

    // The number of ICMP Destination Unreachable messages sent.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutdestunreachs interface{}

    // The number of ICMP Time Exceeded messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpouttimeexcds interface{}

    // The number of ICMP Parameter Problem messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutparmprobs interface{}

    // The number of ICMP Source Quench messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutsrcquenchs interface{}

    // The number of ICMP Redirect messages sent.  For a host, this object will
    // always be zero, since hosts do not send redirects.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutredirects interface{}

    // The number of ICMP Echo (request) messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutechos interface{}

    // The number of ICMP Echo Reply messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutechoreps interface{}

    // The number of ICMP Timestamp (request) messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpouttimestamps interface{}

    // The number of ICMP Timestamp Reply messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpouttimestampreps interface{}

    // The number of ICMP Address Mask Request messages sent.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutaddrmasks interface{}

    // The number of ICMP Address Mask Reply messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    Icmpoutaddrmaskreps interface{}
}

func (icmp *IPMIB_Icmp) GetFilter() yfilter.YFilter { return icmp.YFilter }

func (icmp *IPMIB_Icmp) SetFilter(yf yfilter.YFilter) { icmp.YFilter = yf }

func (icmp *IPMIB_Icmp) GetGoName(yname string) string {
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

func (icmp *IPMIB_Icmp) GetSegmentPath() string {
    return "icmp"
}

func (icmp *IPMIB_Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmp *IPMIB_Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmp *IPMIB_Icmp) GetLeafs() map[string]interface{} {
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

func (icmp *IPMIB_Icmp) GetBundleName() string { return "cisco_ios_xe" }

func (icmp *IPMIB_Icmp) GetYangName() string { return "icmp" }

func (icmp *IPMIB_Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmp *IPMIB_Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmp *IPMIB_Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmp *IPMIB_Icmp) SetParent(parent types.Entity) { icmp.parent = parent }

func (icmp *IPMIB_Icmp) GetParent() types.Entity { return icmp.parent }

func (icmp *IPMIB_Icmp) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipaddrtable
// The table of addressing information relevant to this
// entity's IPv4 addresses.
// 
// This table has been deprecated, as a new IP version-neutral
// table has been added.  It is loosely replaced by the
// ipAddressTable although several objects that weren't deemed
// useful weren't carried forward while another
// (ipAdEntReasmMaxSize) was moved to the ipv4InterfaceTable.
type IPMIB_Ipaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The addressing information for one of this entity's IPv4 addresses. The
    // type is slice of IPMIB_Ipaddrtable_Ipaddrentry.
    Ipaddrentry []IPMIB_Ipaddrtable_Ipaddrentry
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetFilter() yfilter.YFilter { return ipaddrtable.YFilter }

func (ipaddrtable *IPMIB_Ipaddrtable) SetFilter(yf yfilter.YFilter) { ipaddrtable.YFilter = yf }

func (ipaddrtable *IPMIB_Ipaddrtable) GetGoName(yname string) string {
    if yname == "ipAddrEntry" { return "Ipaddrentry" }
    return ""
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetSegmentPath() string {
    return "ipAddrTable"
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipAddrEntry" {
        for _, c := range ipaddrtable.Ipaddrentry {
            if ipaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipaddrtable_Ipaddrentry{}
        ipaddrtable.Ipaddrentry = append(ipaddrtable.Ipaddrentry, child)
        return &ipaddrtable.Ipaddrentry[len(ipaddrtable.Ipaddrentry)-1]
    }
    return nil
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipaddrtable.Ipaddrentry {
        children[ipaddrtable.Ipaddrentry[i].GetSegmentPath()] = &ipaddrtable.Ipaddrentry[i]
    }
    return children
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddrtable *IPMIB_Ipaddrtable) GetYangName() string { return "ipAddrTable" }

func (ipaddrtable *IPMIB_Ipaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddrtable *IPMIB_Ipaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddrtable *IPMIB_Ipaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddrtable *IPMIB_Ipaddrtable) SetParent(parent types.Entity) { ipaddrtable.parent = parent }

func (ipaddrtable *IPMIB_Ipaddrtable) GetParent() types.Entity { return ipaddrtable.parent }

func (ipaddrtable *IPMIB_Ipaddrtable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipaddrtable_Ipaddrentry
// The addressing information for one of this entity's IPv4
// addresses.
type IPMIB_Ipaddrtable_Ipaddrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address to which this entry's addressing
    // information pertains. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipadentaddr interface{}

    // The index value which uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of the IF-MIB's
    // ifIndex. The type is interface{} with range: 1..2147483647.
    Ipadentifindex interface{}

    // The subnet mask associated with the IPv4 address of this entry.  The value
    // of the mask is an IPv4 address with all the network bits set to 1 and all
    // the hosts bits set to 0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipadentnetmask interface{}

    // The value of the least-significant bit in the IPv4 broadcast address used
    // for sending datagrams on the (logical) interface associated with the IPv4
    // address of this entry. For example, when the Internet standard all-ones
    // broadcast address is used, the value will be 1.  This value applies to both
    // the subnet and network broadcast addresses used by the entity on this
    // (logical) interface. The type is interface{} with range: 0..1.
    Ipadentbcastaddr interface{}

    // The size of the largest IPv4 datagram which this entity can re-assemble
    // from incoming IPv4 fragmented datagrams received on this interface. The
    // type is interface{} with range: 0..65535.
    Ipadentreasmmaxsize interface{}
}

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetFilter() yfilter.YFilter { return ipaddrentry.YFilter }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) SetFilter(yf yfilter.YFilter) { ipaddrentry.YFilter = yf }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetGoName(yname string) string {
    if yname == "ipAdEntAddr" { return "Ipadentaddr" }
    if yname == "ipAdEntIfIndex" { return "Ipadentifindex" }
    if yname == "ipAdEntNetMask" { return "Ipadentnetmask" }
    if yname == "ipAdEntBcastAddr" { return "Ipadentbcastaddr" }
    if yname == "ipAdEntReasmMaxSize" { return "Ipadentreasmmaxsize" }
    return ""
}

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetSegmentPath() string {
    return "ipAddrEntry" + "[ipAdEntAddr='" + fmt.Sprintf("%v", ipaddrentry.Ipadentaddr) + "']"
}

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipAdEntAddr"] = ipaddrentry.Ipadentaddr
    leafs["ipAdEntIfIndex"] = ipaddrentry.Ipadentifindex
    leafs["ipAdEntNetMask"] = ipaddrentry.Ipadentnetmask
    leafs["ipAdEntBcastAddr"] = ipaddrentry.Ipadentbcastaddr
    leafs["ipAdEntReasmMaxSize"] = ipaddrentry.Ipadentreasmmaxsize
    return leafs
}

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetYangName() string { return "ipAddrEntry" }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) SetParent(parent types.Entity) { ipaddrentry.parent = parent }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetParent() types.Entity { return ipaddrentry.parent }

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetParentYangName() string { return "ipAddrTable" }

// IPMIB_Ipnettomediatable
// The IPv4 Address Translation table used for mapping from
// IPv4 addresses to physical addresses.
// 
// This table has been deprecated, as a new IP version-neutral
// table has been added.  It is loosely replaced by the
// ipNetToPhysicalTable.
type IPMIB_Ipnettomediatable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains one IpAddress to `physical' address equivalence. The
    // type is slice of IPMIB_Ipnettomediatable_Ipnettomediaentry.
    Ipnettomediaentry []IPMIB_Ipnettomediatable_Ipnettomediaentry
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetFilter() yfilter.YFilter { return ipnettomediatable.YFilter }

func (ipnettomediatable *IPMIB_Ipnettomediatable) SetFilter(yf yfilter.YFilter) { ipnettomediatable.YFilter = yf }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetGoName(yname string) string {
    if yname == "ipNetToMediaEntry" { return "Ipnettomediaentry" }
    return ""
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetSegmentPath() string {
    return "ipNetToMediaTable"
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipNetToMediaEntry" {
        for _, c := range ipnettomediatable.Ipnettomediaentry {
            if ipnettomediatable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipnettomediatable_Ipnettomediaentry{}
        ipnettomediatable.Ipnettomediaentry = append(ipnettomediatable.Ipnettomediaentry, child)
        return &ipnettomediatable.Ipnettomediaentry[len(ipnettomediatable.Ipnettomediaentry)-1]
    }
    return nil
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipnettomediatable.Ipnettomediaentry {
        children[ipnettomediatable.Ipnettomediaentry[i].GetSegmentPath()] = &ipnettomediatable.Ipnettomediaentry[i]
    }
    return children
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetBundleName() string { return "cisco_ios_xe" }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetYangName() string { return "ipNetToMediaTable" }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipnettomediatable *IPMIB_Ipnettomediatable) SetParent(parent types.Entity) { ipnettomediatable.parent = parent }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetParent() types.Entity { return ipnettomediatable.parent }

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipnettomediatable_Ipnettomediaentry
// Each entry contains one IpAddress to `physical' address
// equivalence.
type IPMIB_Ipnettomediatable_Ipnettomediaentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the   same value of the IF-MIB's
    // ifIndex.  This object predates the rule limiting index objects to a max
    // access value of 'not-accessible' and so continues to use a value of
    // 'read-create'. The type is interface{} with range: 1..2147483647.
    Ipnettomediaifindex interface{}

    // This attribute is a key. The IpAddress corresponding to the media-dependent
    // `physical' address.  This object predates the rule limiting index objects
    // to a max access value of 'not-accessible' and so continues to use a value
    // of 'read-create'. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipnettomedianetaddress interface{}

    // The media-dependent `physical' address.  This object should return 0 when
    // this entry is in the 'incomplete' state.  As the entries in this table are
    // typically not persistent when this object is written the entity should not
    // save the change to non-volatile storage.  Note: a stronger requirement is
    // not used because this object was previously defined. The type is string
    // with length: 0..65535.
    Ipnettomediaphysaddress interface{}

    // The type of mapping.  Setting this object to the value invalid(2) has the
    // effect   of invalidating the corresponding entry in the ipNetToMediaTable. 
    // That is, it effectively dis-associates the interface identified with said
    // entry from the mapping identified with said entry.  It is an
    // implementation- specific matter as to whether the agent removes an
    // invalidated entry from the table.  Accordingly, management stations must be
    // prepared to receive tabular information from agents that corresponds to
    // entries not currently in use.  Proper interpretation of such entries
    // requires examination of the relevant ipNetToMediaType object.  As the
    // entries in this table are typically not persistent when this object is
    // written the entity should not save the change to non-volatile storage. 
    // Note: a stronger requirement is not used because this object was previously
    // defined. The type is Ipnettomediatype.
    Ipnettomediatype interface{}
}

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetFilter() yfilter.YFilter { return ipnettomediaentry.YFilter }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) SetFilter(yf yfilter.YFilter) { ipnettomediaentry.YFilter = yf }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetGoName(yname string) string {
    if yname == "ipNetToMediaIfIndex" { return "Ipnettomediaifindex" }
    if yname == "ipNetToMediaNetAddress" { return "Ipnettomedianetaddress" }
    if yname == "ipNetToMediaPhysAddress" { return "Ipnettomediaphysaddress" }
    if yname == "ipNetToMediaType" { return "Ipnettomediatype" }
    return ""
}

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetSegmentPath() string {
    return "ipNetToMediaEntry" + "[ipNetToMediaIfIndex='" + fmt.Sprintf("%v", ipnettomediaentry.Ipnettomediaifindex) + "']" + "[ipNetToMediaNetAddress='" + fmt.Sprintf("%v", ipnettomediaentry.Ipnettomedianetaddress) + "']"
}

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipNetToMediaIfIndex"] = ipnettomediaentry.Ipnettomediaifindex
    leafs["ipNetToMediaNetAddress"] = ipnettomediaentry.Ipnettomedianetaddress
    leafs["ipNetToMediaPhysAddress"] = ipnettomediaentry.Ipnettomediaphysaddress
    leafs["ipNetToMediaType"] = ipnettomediaentry.Ipnettomediatype
    return leafs
}

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetYangName() string { return "ipNetToMediaEntry" }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) SetParent(parent types.Entity) { ipnettomediaentry.parent = parent }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetParent() types.Entity { return ipnettomediaentry.parent }

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetParentYangName() string { return "ipNetToMediaTable" }

// IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype represents defined.
type IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype string

const (
    IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_other IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "other"

    IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_invalid IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "invalid"

    IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_dynamic IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "dynamic"

    IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype_static IPMIB_Ipnettomediatable_Ipnettomediaentry_Ipnettomediatype = "static"
)

// IPMIB_Ipv4Interfacetable
// The table containing per-interface IPv4-specific
// information.
type IPMIB_Ipv4Interfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing IPv4-specific information for a specific interface. The
    // type is slice of IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry.
    Ipv4Interfaceentry []IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetFilter() yfilter.YFilter { return ipv4Interfacetable.YFilter }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) SetFilter(yf yfilter.YFilter) { ipv4Interfacetable.YFilter = yf }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetGoName(yname string) string {
    if yname == "ipv4InterfaceEntry" { return "Ipv4Interfaceentry" }
    return ""
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetSegmentPath() string {
    return "ipv4InterfaceTable"
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4InterfaceEntry" {
        for _, c := range ipv4Interfacetable.Ipv4Interfaceentry {
            if ipv4Interfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry{}
        ipv4Interfacetable.Ipv4Interfaceentry = append(ipv4Interfacetable.Ipv4Interfaceentry, child)
        return &ipv4Interfacetable.Ipv4Interfaceentry[len(ipv4Interfacetable.Ipv4Interfaceentry)-1]
    }
    return nil
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Interfacetable.Ipv4Interfaceentry {
        children[ipv4Interfacetable.Ipv4Interfaceentry[i].GetSegmentPath()] = &ipv4Interfacetable.Ipv4Interfaceentry[i]
    }
    return children
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetYangName() string { return "ipv4InterfaceTable" }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) SetParent(parent types.Entity) { ipv4Interfacetable.parent = parent }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetParent() types.Entity { return ipv4Interfacetable.parent }

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry
// An entry containing IPv4-specific information for a specific
// interface.
type IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipv4Interfaceifindex interface{}

    // The size of the largest IPv4 datagram that this entity can re-assemble from
    // incoming IPv4 fragmented datagrams received on this interface. The type is
    // interface{} with range: 0..65535.
    Ipv4Interfacereasmmaxsize interface{}

    // The indication of whether IPv4 is enabled (up) or disabled (down) on this
    // interface.  This object does not affect the state of the interface itself,
    // only its connection to an IPv4 stack.  The IF-MIB should be used to control
    // the state of the interface. The type is Ipv4Interfaceenablestatus.
    Ipv4Interfaceenablestatus interface{}

    // The time between retransmissions of ARP requests to a neighbor when
    // resolving the address or when probing the reachability of a neighbor. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Ipv4Interfaceretransmittime interface{}
}

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetFilter() yfilter.YFilter { return ipv4Interfaceentry.YFilter }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) SetFilter(yf yfilter.YFilter) { ipv4Interfaceentry.YFilter = yf }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetGoName(yname string) string {
    if yname == "ipv4InterfaceIfIndex" { return "Ipv4Interfaceifindex" }
    if yname == "ipv4InterfaceReasmMaxSize" { return "Ipv4Interfacereasmmaxsize" }
    if yname == "ipv4InterfaceEnableStatus" { return "Ipv4Interfaceenablestatus" }
    if yname == "ipv4InterfaceRetransmitTime" { return "Ipv4Interfaceretransmittime" }
    return ""
}

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetSegmentPath() string {
    return "ipv4InterfaceEntry" + "[ipv4InterfaceIfIndex='" + fmt.Sprintf("%v", ipv4Interfaceentry.Ipv4Interfaceifindex) + "']"
}

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4InterfaceIfIndex"] = ipv4Interfaceentry.Ipv4Interfaceifindex
    leafs["ipv4InterfaceReasmMaxSize"] = ipv4Interfaceentry.Ipv4Interfacereasmmaxsize
    leafs["ipv4InterfaceEnableStatus"] = ipv4Interfaceentry.Ipv4Interfaceenablestatus
    leafs["ipv4InterfaceRetransmitTime"] = ipv4Interfaceentry.Ipv4Interfaceretransmittime
    return leafs
}

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetYangName() string { return "ipv4InterfaceEntry" }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) SetParent(parent types.Entity) { ipv4Interfaceentry.parent = parent }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetParent() types.Entity { return ipv4Interfaceentry.parent }

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetParentYangName() string { return "ipv4InterfaceTable" }

// IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry_Ipv4Interfaceenablestatus represents of the interface.
type IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry_Ipv4Interfaceenablestatus string

const (
    IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry_Ipv4Interfaceenablestatus_up IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry_Ipv4Interfaceenablestatus = "up"

    IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry_Ipv4Interfaceenablestatus_down IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry_Ipv4Interfaceenablestatus = "down"
)

// IPMIB_Ipv6Interfacetable
// The table containing per-interface IPv6-specific
// information.
type IPMIB_Ipv6Interfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing IPv6-specific information for a given interface. The
    // type is slice of IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry.
    Ipv6Interfaceentry []IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetFilter() yfilter.YFilter { return ipv6Interfacetable.YFilter }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) SetFilter(yf yfilter.YFilter) { ipv6Interfacetable.YFilter = yf }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetGoName(yname string) string {
    if yname == "ipv6InterfaceEntry" { return "Ipv6Interfaceentry" }
    return ""
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetSegmentPath() string {
    return "ipv6InterfaceTable"
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6InterfaceEntry" {
        for _, c := range ipv6Interfacetable.Ipv6Interfaceentry {
            if ipv6Interfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry{}
        ipv6Interfacetable.Ipv6Interfaceentry = append(ipv6Interfacetable.Ipv6Interfaceentry, child)
        return &ipv6Interfacetable.Ipv6Interfaceentry[len(ipv6Interfacetable.Ipv6Interfaceentry)-1]
    }
    return nil
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Interfacetable.Ipv6Interfaceentry {
        children[ipv6Interfacetable.Ipv6Interfaceentry[i].GetSegmentPath()] = &ipv6Interfacetable.Ipv6Interfaceentry[i]
    }
    return children
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetYangName() string { return "ipv6InterfaceTable" }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) SetParent(parent types.Entity) { ipv6Interfacetable.parent = parent }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetParent() types.Entity { return ipv6Interfacetable.parent }

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry
// An entry containing IPv6-specific information for a given
// interface.
type IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipv6Interfaceifindex interface{}

    // The size of the largest IPv6 datagram that this entity can re-assemble from
    // incoming IPv6 fragmented datagrams received on this interface. The type is
    // interface{} with range: 1500..65535. Units are octets.
    Ipv6Interfacereasmmaxsize interface{}

    // The Interface Identifier for this interface.  The Interface Identifier is
    // combined with an address prefix to form an interface address.  By default,
    // the Interface Identifier is auto-configured according to the rules of the
    // link type to which this interface is attached.   A zero length identifier
    // may be used where appropriate.  One possible example is a loopback
    // interface. The type is string.
    Ipv6Interfaceidentifier interface{}

    // The indication of whether IPv6 is enabled (up) or disabled (down) on this
    // interface.  This object does not affect the state of the interface itself,
    // only its connection to an IPv6 stack.  The IF-MIB should be used to control
    // the state of the interface.  When this object is written, the entity SHOULD
    // save the change to non-volatile storage and restore the object from
    // non-volatile storage upon re-initialization of the system. The type is
    // Ipv6Interfaceenablestatus.
    Ipv6Interfaceenablestatus interface{}

    // The time a neighbor is considered reachable after receiving a reachability
    // confirmation. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Ipv6Interfacereachabletime interface{}

    // The time between retransmissions of Neighbor Solicitation messages to a
    // neighbor when resolving the address or when probing the reachability of a
    // neighbor. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Ipv6Interfaceretransmittime interface{}

    // The indication of whether this entity is acting as an IPv6 router on this
    // interface with respect to the forwarding of datagrams received by, but not
    // addressed to, this entity. IPv6 routers forward datagrams.  IPv6 hosts do
    // not (except those source-routed via the host).  This object is constrained
    // by ipv6IpForwarding and is ignored if ipv6IpForwarding is set to
    // notForwarding.  Those systems that do not provide per-interface control of
    // the forwarding function should set this object to forwarding for all
    // interfaces and allow the ipv6IpForwarding object to control the forwarding
    // capability.  When this object is written, the entity SHOULD save the change
    // to non-volatile storage and restore the object from non-volatile storage
    // upon re-initialization of the system. The type is Ipv6Interfaceforwarding.
    Ipv6Interfaceforwarding interface{}
}

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetFilter() yfilter.YFilter { return ipv6Interfaceentry.YFilter }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) SetFilter(yf yfilter.YFilter) { ipv6Interfaceentry.YFilter = yf }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetGoName(yname string) string {
    if yname == "ipv6InterfaceIfIndex" { return "Ipv6Interfaceifindex" }
    if yname == "ipv6InterfaceReasmMaxSize" { return "Ipv6Interfacereasmmaxsize" }
    if yname == "ipv6InterfaceIdentifier" { return "Ipv6Interfaceidentifier" }
    if yname == "ipv6InterfaceEnableStatus" { return "Ipv6Interfaceenablestatus" }
    if yname == "ipv6InterfaceReachableTime" { return "Ipv6Interfacereachabletime" }
    if yname == "ipv6InterfaceRetransmitTime" { return "Ipv6Interfaceretransmittime" }
    if yname == "ipv6InterfaceForwarding" { return "Ipv6Interfaceforwarding" }
    return ""
}

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetSegmentPath() string {
    return "ipv6InterfaceEntry" + "[ipv6InterfaceIfIndex='" + fmt.Sprintf("%v", ipv6Interfaceentry.Ipv6Interfaceifindex) + "']"
}

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6InterfaceIfIndex"] = ipv6Interfaceentry.Ipv6Interfaceifindex
    leafs["ipv6InterfaceReasmMaxSize"] = ipv6Interfaceentry.Ipv6Interfacereasmmaxsize
    leafs["ipv6InterfaceIdentifier"] = ipv6Interfaceentry.Ipv6Interfaceidentifier
    leafs["ipv6InterfaceEnableStatus"] = ipv6Interfaceentry.Ipv6Interfaceenablestatus
    leafs["ipv6InterfaceReachableTime"] = ipv6Interfaceentry.Ipv6Interfacereachabletime
    leafs["ipv6InterfaceRetransmitTime"] = ipv6Interfaceentry.Ipv6Interfaceretransmittime
    leafs["ipv6InterfaceForwarding"] = ipv6Interfaceentry.Ipv6Interfaceforwarding
    return leafs
}

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetYangName() string { return "ipv6InterfaceEntry" }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) SetParent(parent types.Entity) { ipv6Interfaceentry.parent = parent }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetParent() types.Entity { return ipv6Interfaceentry.parent }

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetParentYangName() string { return "ipv6InterfaceTable" }

// IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceenablestatus represents non-volatile storage upon re-initialization of the system.
type IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceenablestatus string

const (
    IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceenablestatus_up IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceenablestatus = "up"

    IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceenablestatus_down IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceenablestatus = "down"
)

// IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceforwarding represents non-volatile storage upon re-initialization of the system.
type IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceforwarding string

const (
    IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceforwarding_forwarding IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceforwarding = "forwarding"

    IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceforwarding_notForwarding IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry_Ipv6Interfaceforwarding = "notForwarding"
)

// IPMIB_Ipsystemstatstable
// The table containing system wide, IP version specific
// traffic statistics.  This table and the ipIfStatsTable
// contain similar objects whose difference is in their
// granularity.  Where this table contains system wide traffic
// statistics, the ipIfStatsTable contains the same statistics
// but counted on a per-interface basis.
type IPMIB_Ipsystemstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A statistics entry containing system-wide objects for a particular IP
    // version. The type is slice of IPMIB_Ipsystemstatstable_Ipsystemstatsentry.
    Ipsystemstatsentry []IPMIB_Ipsystemstatstable_Ipsystemstatsentry
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetFilter() yfilter.YFilter { return ipsystemstatstable.YFilter }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) SetFilter(yf yfilter.YFilter) { ipsystemstatstable.YFilter = yf }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetGoName(yname string) string {
    if yname == "ipSystemStatsEntry" { return "Ipsystemstatsentry" }
    return ""
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetSegmentPath() string {
    return "ipSystemStatsTable"
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipSystemStatsEntry" {
        for _, c := range ipsystemstatstable.Ipsystemstatsentry {
            if ipsystemstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipsystemstatstable_Ipsystemstatsentry{}
        ipsystemstatstable.Ipsystemstatsentry = append(ipsystemstatstable.Ipsystemstatsentry, child)
        return &ipsystemstatstable.Ipsystemstatsentry[len(ipsystemstatstable.Ipsystemstatsentry)-1]
    }
    return nil
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipsystemstatstable.Ipsystemstatsentry {
        children[ipsystemstatstable.Ipsystemstatsentry[i].GetSegmentPath()] = &ipsystemstatstable.Ipsystemstatsentry[i]
    }
    return children
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetYangName() string { return "ipSystemStatsTable" }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) SetParent(parent types.Entity) { ipsystemstatstable.parent = parent }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetParent() types.Entity { return ipsystemstatstable.parent }

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipsystemstatstable_Ipsystemstatsentry
// A statistics entry containing system-wide objects for a
// particular IP version.
type IPMIB_Ipsystemstatstable_Ipsystemstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP version of this row. The type is IpVersion.
    Ipsystemstatsipversion interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ipsystemstatsinreceives interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  This object counts the same datagrams as
    // ipSystemStatsInReceives, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcinreceives interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  Octets from datagrams counted in
    // ipSystemStatsInReceives MUST be counted here.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsinoctets interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  This object counts the same octets as
    // ipSystemStatsInOctets, but allows for larger   values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcinoctets interface{}

    // The number of input IP datagrams discarded due to errors in their IP
    // headers, including version number mismatch, other format errors, hop count
    // exceeded, errors discovered in processing their IP options, etc. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsinhdrerrors interface{}

    // The number of input IP datagrams discarded because no route could be found
    // to transmit them to their destination.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsinnoroutes interface{}

    // The number of input IP datagrams discarded because the IP address in their
    // IP header's destination field was not a valid address to be received at
    // this entity.  This count includes invalid addresses (e.g., ::0).  For
    // entities that are not IP routers and therefore do not forward   datagrams,
    // this counter includes datagrams discarded because the destination address
    // was not a local address.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipsystemstatsinaddrerrors interface{}

    // The number of locally-addressed IP datagrams received successfully but
    // discarded because of an unknown or unsupported protocol.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsinunknownprotos interface{}

    // The number of input IP datagrams discarded because the datagram frame
    // didn't carry enough data.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipsystemstatsintruncatedpkts interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination and for which this entity attempted to find a route to forward
    // them to that final destination.  In entities that do not act as IP routers,
    // this counter will include only those datagrams that were Source-Routed via
    // this entity, and the Source-Route processing was successful.  When tracking
    // interface statistics, the counter of the incoming interface is incremented
    // for each datagram.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipsystemstatsinforwdatagrams interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination and for which this entity attempted to find a route to forward
    // them to that final destination.  This object counts the same packets as
    // ipSystemStatsInForwDatagrams, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcinforwdatagrams interface{}

    // The number of IP fragments received that needed to be reassembled at this
    // interface.  When tracking interface statistics, the counter of the
    // interface to which these fragments were addressed is incremented.  This
    // interface might not be the same as the input interface for some of the
    // fragments.  Discontinuities in the value of this counter can occur at  
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ipsystemstatsreasmreqds interface{}

    // The number of IP datagrams successfully reassembled.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsreasmoks interface{}

    // The number of failures detected by the IP re-assembly algorithm (for
    // whatever reason: timed out, errors, etc.). Note that this is not
    // necessarily a count of discarded IP fragments since some algorithms
    // (notably the algorithm in RFC 815) can lose track of the number of
    // fragments by combining them as they are received.  When tracking interface
    // statistics, the counter of the interface to which these fragments were
    // addressed is incremented.  This interface might not be the same as the
    // input interface for some of the fragments.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsreasmfails interface{}

    // The number of input IP datagrams for which no problems were encountered to
    // prevent their continued processing, but were discarded (e.g., for lack of
    // buffer space).  Note that this counter does not include any datagrams
    // discarded while awaiting re-assembly.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsindiscards interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  When tracking interface statistics, the counter of the
    // interface to which these datagrams were addressed is incremented.  This
    // interface might not be the same as the input interface for some of the
    // datagrams.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ipsystemstatsindelivers interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  This object counts the same packets as
    // ipSystemStatsInDelivers, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcindelivers interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  Note that this counter
    // does not include any datagrams counted in ipSystemStatsOutForwDatagrams. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutrequests interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  This object counts the
    // same packets as ipSystemStatsOutRequests, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcoutrequests interface{}

    // The number of locally generated IP datagrams discarded because no route
    // could be found to transmit them to their destination.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutnoroutes interface{}

    // The number of datagrams for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.  In entities that do not act as IP routers, this counter
    // will include only those datagrams that were Source-Routed via this entity,
    // and the Source-Route processing was successful.  When tracking interface
    // statistics, the counter of the outgoing interface is incremented for a
    // successfully forwarded datagram.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsoutforwdatagrams interface{}

    // The number of datagrams for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.  This object counts the same packets as
    // ipSystemStatsOutForwDatagrams, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcoutforwdatagrams interface{}

    // The number of output IP datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but were discarded (e.g.,
    // for lack of buffer space).  Note that this counter would include  
    // datagrams counted in ipSystemStatsOutForwDatagrams if any such datagrams
    // met this (discretionary) discard criterion.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutdiscards interface{}

    // The number of IP datagrams that would require fragmentation in order to be
    // transmitted.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for a successfully fragmented datagram. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutfragreqds interface{}

    // The number of IP datagrams that have been successfully fragmented.  When
    // tracking interface statistics, the counter of the outgoing interface is
    // incremented for a successfully fragmented datagram.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutfragoks interface{}

    // The number of IP datagrams that have been discarded because they needed to
    // be fragmented but could not be.  This includes IPv4 packets that have the
    // DF bit set and IPv6 packets that are being forwarded and exceed the
    // outgoing link MTU.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for an unsuccessfully fragmented
    // datagram.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ipsystemstatsoutfragfails interface{}

    // The number of output datagram fragments that have been generated as a
    // result of IP fragmentation.  When tracking interface statistics, the
    // counter of the outgoing interface is incremented for a successfully
    // fragmented datagram.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipsystemstatsoutfragcreates interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This includes datagrams generated locally and
    // those forwarded by this entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other   times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipsystemstatsouttransmits interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This object counts the same datagrams as
    // ipSystemStatsOutTransmits, but allows for larger values.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcouttransmits interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  Octets from datagrams counted in
    // ipSystemStatsOutTransmits MUST be counted here.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutoctets interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  This objects counts the same octets as
    // ipSystemStatsOutOctets, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of  
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcoutoctets interface{}

    // The number of IP multicast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsinmcastpkts interface{}

    // The number of IP multicast datagrams received.  This object counts the same
    // datagrams as ipSystemStatsInMcastPkts but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcinmcastpkts interface{}

    // The total number of octets received in IP multicast datagrams.  Octets from
    // datagrams counted in ipSystemStatsInMcastPkts MUST be counted here. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsinmcastoctets interface{}

    // The total number of octets received in IP multicast datagrams.  This object
    // counts the same octets as ipSystemStatsInMcastOctets, but allows for larger
    // values.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..18446744073709551615.
    Ipsystemstatshcinmcastoctets interface{}

    // The number of IP multicast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutmcastpkts interface{}

    // The number of IP multicast datagrams transmitted.  This object counts the
    // same datagrams as ipSystemStatsOutMcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcoutmcastpkts interface{}

    // The total number of octets transmitted in IP multicast datagrams.  Octets
    // from datagrams counted in   ipSystemStatsOutMcastPkts MUST be counted here.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutmcastoctets interface{}

    // The total number of octets transmitted in IP multicast datagrams.  This
    // object counts the same octets as ipSystemStatsOutMcastOctets, but allows
    // for larger values.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    Ipsystemstatshcoutmcastoctets interface{}

    // The number of IP broadcast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsinbcastpkts interface{}

    // The number of IP broadcast datagrams received.  This object counts the same
    // datagrams as ipSystemStatsInBcastPkts but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of  
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcinbcastpkts interface{}

    // The number of IP broadcast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipsystemstatsoutbcastpkts interface{}

    // The number of IP broadcast datagrams transmitted.  This object counts the
    // same datagrams as ipSystemStatsOutBcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipsystemstatshcoutbcastpkts interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Ipsystemstatsdiscontinuitytime interface{}

    // The minimum reasonable polling interval for this entry. This object
    // provides an indication of the minimum amount of time required to update the
    // counters in this entry. The type is interface{} with range: 0..4294967295.
    // Units are milli-seconds.
    Ipsystemstatsrefreshrate interface{}
}

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetFilter() yfilter.YFilter { return ipsystemstatsentry.YFilter }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) SetFilter(yf yfilter.YFilter) { ipsystemstatsentry.YFilter = yf }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetGoName(yname string) string {
    if yname == "ipSystemStatsIPVersion" { return "Ipsystemstatsipversion" }
    if yname == "ipSystemStatsInReceives" { return "Ipsystemstatsinreceives" }
    if yname == "ipSystemStatsHCInReceives" { return "Ipsystemstatshcinreceives" }
    if yname == "ipSystemStatsInOctets" { return "Ipsystemstatsinoctets" }
    if yname == "ipSystemStatsHCInOctets" { return "Ipsystemstatshcinoctets" }
    if yname == "ipSystemStatsInHdrErrors" { return "Ipsystemstatsinhdrerrors" }
    if yname == "ipSystemStatsInNoRoutes" { return "Ipsystemstatsinnoroutes" }
    if yname == "ipSystemStatsInAddrErrors" { return "Ipsystemstatsinaddrerrors" }
    if yname == "ipSystemStatsInUnknownProtos" { return "Ipsystemstatsinunknownprotos" }
    if yname == "ipSystemStatsInTruncatedPkts" { return "Ipsystemstatsintruncatedpkts" }
    if yname == "ipSystemStatsInForwDatagrams" { return "Ipsystemstatsinforwdatagrams" }
    if yname == "ipSystemStatsHCInForwDatagrams" { return "Ipsystemstatshcinforwdatagrams" }
    if yname == "ipSystemStatsReasmReqds" { return "Ipsystemstatsreasmreqds" }
    if yname == "ipSystemStatsReasmOKs" { return "Ipsystemstatsreasmoks" }
    if yname == "ipSystemStatsReasmFails" { return "Ipsystemstatsreasmfails" }
    if yname == "ipSystemStatsInDiscards" { return "Ipsystemstatsindiscards" }
    if yname == "ipSystemStatsInDelivers" { return "Ipsystemstatsindelivers" }
    if yname == "ipSystemStatsHCInDelivers" { return "Ipsystemstatshcindelivers" }
    if yname == "ipSystemStatsOutRequests" { return "Ipsystemstatsoutrequests" }
    if yname == "ipSystemStatsHCOutRequests" { return "Ipsystemstatshcoutrequests" }
    if yname == "ipSystemStatsOutNoRoutes" { return "Ipsystemstatsoutnoroutes" }
    if yname == "ipSystemStatsOutForwDatagrams" { return "Ipsystemstatsoutforwdatagrams" }
    if yname == "ipSystemStatsHCOutForwDatagrams" { return "Ipsystemstatshcoutforwdatagrams" }
    if yname == "ipSystemStatsOutDiscards" { return "Ipsystemstatsoutdiscards" }
    if yname == "ipSystemStatsOutFragReqds" { return "Ipsystemstatsoutfragreqds" }
    if yname == "ipSystemStatsOutFragOKs" { return "Ipsystemstatsoutfragoks" }
    if yname == "ipSystemStatsOutFragFails" { return "Ipsystemstatsoutfragfails" }
    if yname == "ipSystemStatsOutFragCreates" { return "Ipsystemstatsoutfragcreates" }
    if yname == "ipSystemStatsOutTransmits" { return "Ipsystemstatsouttransmits" }
    if yname == "ipSystemStatsHCOutTransmits" { return "Ipsystemstatshcouttransmits" }
    if yname == "ipSystemStatsOutOctets" { return "Ipsystemstatsoutoctets" }
    if yname == "ipSystemStatsHCOutOctets" { return "Ipsystemstatshcoutoctets" }
    if yname == "ipSystemStatsInMcastPkts" { return "Ipsystemstatsinmcastpkts" }
    if yname == "ipSystemStatsHCInMcastPkts" { return "Ipsystemstatshcinmcastpkts" }
    if yname == "ipSystemStatsInMcastOctets" { return "Ipsystemstatsinmcastoctets" }
    if yname == "ipSystemStatsHCInMcastOctets" { return "Ipsystemstatshcinmcastoctets" }
    if yname == "ipSystemStatsOutMcastPkts" { return "Ipsystemstatsoutmcastpkts" }
    if yname == "ipSystemStatsHCOutMcastPkts" { return "Ipsystemstatshcoutmcastpkts" }
    if yname == "ipSystemStatsOutMcastOctets" { return "Ipsystemstatsoutmcastoctets" }
    if yname == "ipSystemStatsHCOutMcastOctets" { return "Ipsystemstatshcoutmcastoctets" }
    if yname == "ipSystemStatsInBcastPkts" { return "Ipsystemstatsinbcastpkts" }
    if yname == "ipSystemStatsHCInBcastPkts" { return "Ipsystemstatshcinbcastpkts" }
    if yname == "ipSystemStatsOutBcastPkts" { return "Ipsystemstatsoutbcastpkts" }
    if yname == "ipSystemStatsHCOutBcastPkts" { return "Ipsystemstatshcoutbcastpkts" }
    if yname == "ipSystemStatsDiscontinuityTime" { return "Ipsystemstatsdiscontinuitytime" }
    if yname == "ipSystemStatsRefreshRate" { return "Ipsystemstatsrefreshrate" }
    return ""
}

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetSegmentPath() string {
    return "ipSystemStatsEntry" + "[ipSystemStatsIPVersion='" + fmt.Sprintf("%v", ipsystemstatsentry.Ipsystemstatsipversion) + "']"
}

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipSystemStatsIPVersion"] = ipsystemstatsentry.Ipsystemstatsipversion
    leafs["ipSystemStatsInReceives"] = ipsystemstatsentry.Ipsystemstatsinreceives
    leafs["ipSystemStatsHCInReceives"] = ipsystemstatsentry.Ipsystemstatshcinreceives
    leafs["ipSystemStatsInOctets"] = ipsystemstatsentry.Ipsystemstatsinoctets
    leafs["ipSystemStatsHCInOctets"] = ipsystemstatsentry.Ipsystemstatshcinoctets
    leafs["ipSystemStatsInHdrErrors"] = ipsystemstatsentry.Ipsystemstatsinhdrerrors
    leafs["ipSystemStatsInNoRoutes"] = ipsystemstatsentry.Ipsystemstatsinnoroutes
    leafs["ipSystemStatsInAddrErrors"] = ipsystemstatsentry.Ipsystemstatsinaddrerrors
    leafs["ipSystemStatsInUnknownProtos"] = ipsystemstatsentry.Ipsystemstatsinunknownprotos
    leafs["ipSystemStatsInTruncatedPkts"] = ipsystemstatsentry.Ipsystemstatsintruncatedpkts
    leafs["ipSystemStatsInForwDatagrams"] = ipsystemstatsentry.Ipsystemstatsinforwdatagrams
    leafs["ipSystemStatsHCInForwDatagrams"] = ipsystemstatsentry.Ipsystemstatshcinforwdatagrams
    leafs["ipSystemStatsReasmReqds"] = ipsystemstatsentry.Ipsystemstatsreasmreqds
    leafs["ipSystemStatsReasmOKs"] = ipsystemstatsentry.Ipsystemstatsreasmoks
    leafs["ipSystemStatsReasmFails"] = ipsystemstatsentry.Ipsystemstatsreasmfails
    leafs["ipSystemStatsInDiscards"] = ipsystemstatsentry.Ipsystemstatsindiscards
    leafs["ipSystemStatsInDelivers"] = ipsystemstatsentry.Ipsystemstatsindelivers
    leafs["ipSystemStatsHCInDelivers"] = ipsystemstatsentry.Ipsystemstatshcindelivers
    leafs["ipSystemStatsOutRequests"] = ipsystemstatsentry.Ipsystemstatsoutrequests
    leafs["ipSystemStatsHCOutRequests"] = ipsystemstatsentry.Ipsystemstatshcoutrequests
    leafs["ipSystemStatsOutNoRoutes"] = ipsystemstatsentry.Ipsystemstatsoutnoroutes
    leafs["ipSystemStatsOutForwDatagrams"] = ipsystemstatsentry.Ipsystemstatsoutforwdatagrams
    leafs["ipSystemStatsHCOutForwDatagrams"] = ipsystemstatsentry.Ipsystemstatshcoutforwdatagrams
    leafs["ipSystemStatsOutDiscards"] = ipsystemstatsentry.Ipsystemstatsoutdiscards
    leafs["ipSystemStatsOutFragReqds"] = ipsystemstatsentry.Ipsystemstatsoutfragreqds
    leafs["ipSystemStatsOutFragOKs"] = ipsystemstatsentry.Ipsystemstatsoutfragoks
    leafs["ipSystemStatsOutFragFails"] = ipsystemstatsentry.Ipsystemstatsoutfragfails
    leafs["ipSystemStatsOutFragCreates"] = ipsystemstatsentry.Ipsystemstatsoutfragcreates
    leafs["ipSystemStatsOutTransmits"] = ipsystemstatsentry.Ipsystemstatsouttransmits
    leafs["ipSystemStatsHCOutTransmits"] = ipsystemstatsentry.Ipsystemstatshcouttransmits
    leafs["ipSystemStatsOutOctets"] = ipsystemstatsentry.Ipsystemstatsoutoctets
    leafs["ipSystemStatsHCOutOctets"] = ipsystemstatsentry.Ipsystemstatshcoutoctets
    leafs["ipSystemStatsInMcastPkts"] = ipsystemstatsentry.Ipsystemstatsinmcastpkts
    leafs["ipSystemStatsHCInMcastPkts"] = ipsystemstatsentry.Ipsystemstatshcinmcastpkts
    leafs["ipSystemStatsInMcastOctets"] = ipsystemstatsentry.Ipsystemstatsinmcastoctets
    leafs["ipSystemStatsHCInMcastOctets"] = ipsystemstatsentry.Ipsystemstatshcinmcastoctets
    leafs["ipSystemStatsOutMcastPkts"] = ipsystemstatsentry.Ipsystemstatsoutmcastpkts
    leafs["ipSystemStatsHCOutMcastPkts"] = ipsystemstatsentry.Ipsystemstatshcoutmcastpkts
    leafs["ipSystemStatsOutMcastOctets"] = ipsystemstatsentry.Ipsystemstatsoutmcastoctets
    leafs["ipSystemStatsHCOutMcastOctets"] = ipsystemstatsentry.Ipsystemstatshcoutmcastoctets
    leafs["ipSystemStatsInBcastPkts"] = ipsystemstatsentry.Ipsystemstatsinbcastpkts
    leafs["ipSystemStatsHCInBcastPkts"] = ipsystemstatsentry.Ipsystemstatshcinbcastpkts
    leafs["ipSystemStatsOutBcastPkts"] = ipsystemstatsentry.Ipsystemstatsoutbcastpkts
    leafs["ipSystemStatsHCOutBcastPkts"] = ipsystemstatsentry.Ipsystemstatshcoutbcastpkts
    leafs["ipSystemStatsDiscontinuityTime"] = ipsystemstatsentry.Ipsystemstatsdiscontinuitytime
    leafs["ipSystemStatsRefreshRate"] = ipsystemstatsentry.Ipsystemstatsrefreshrate
    return leafs
}

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetYangName() string { return "ipSystemStatsEntry" }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) SetParent(parent types.Entity) { ipsystemstatsentry.parent = parent }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetParent() types.Entity { return ipsystemstatsentry.parent }

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetParentYangName() string { return "ipSystemStatsTable" }

// IPMIB_Ipifstatstable
// The table containing per-interface traffic statistics.  This
// table and the ipSystemStatsTable contain similar objects
// whose difference is in their granularity.  Where this table
// contains per-interface statistics, the ipSystemStatsTable
// contains the same statistics, but counted on a system wide
// basis.
type IPMIB_Ipifstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An interface statistics entry containing objects for a particular interface
    // and version of IP. The type is slice of
    // IPMIB_Ipifstatstable_Ipifstatsentry.
    Ipifstatsentry []IPMIB_Ipifstatstable_Ipifstatsentry
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetFilter() yfilter.YFilter { return ipifstatstable.YFilter }

func (ipifstatstable *IPMIB_Ipifstatstable) SetFilter(yf yfilter.YFilter) { ipifstatstable.YFilter = yf }

func (ipifstatstable *IPMIB_Ipifstatstable) GetGoName(yname string) string {
    if yname == "ipIfStatsEntry" { return "Ipifstatsentry" }
    return ""
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetSegmentPath() string {
    return "ipIfStatsTable"
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipIfStatsEntry" {
        for _, c := range ipifstatstable.Ipifstatsentry {
            if ipifstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipifstatstable_Ipifstatsentry{}
        ipifstatstable.Ipifstatsentry = append(ipifstatstable.Ipifstatsentry, child)
        return &ipifstatstable.Ipifstatsentry[len(ipifstatstable.Ipifstatsentry)-1]
    }
    return nil
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipifstatstable.Ipifstatsentry {
        children[ipifstatstable.Ipifstatsentry[i].GetSegmentPath()] = &ipifstatstable.Ipifstatsentry[i]
    }
    return children
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (ipifstatstable *IPMIB_Ipifstatstable) GetYangName() string { return "ipIfStatsTable" }

func (ipifstatstable *IPMIB_Ipifstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipifstatstable *IPMIB_Ipifstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipifstatstable *IPMIB_Ipifstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipifstatstable *IPMIB_Ipifstatstable) SetParent(parent types.Entity) { ipifstatstable.parent = parent }

func (ipifstatstable *IPMIB_Ipifstatstable) GetParent() types.Entity { return ipifstatstable.parent }

func (ipifstatstable *IPMIB_Ipifstatstable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipifstatstable_Ipifstatsentry
// An interface statistics entry containing objects for a
// particular interface and version of IP.
type IPMIB_Ipifstatstable_Ipifstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP version of this row. The type is IpVersion.
    Ipifstatsipversion interface{}

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipifstatsifindex interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ipifstatsinreceives interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  This object counts the same datagrams as ipIfStatsInReceives,
    // but allows for larger values.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    Ipifstatshcinreceives interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  Octets from datagrams counted in ipIfStatsInReceives
    // MUST be counted here.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipifstatsinoctets interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  This object counts the same octets as
    // ipIfStatsInOctets, but allows for larger values.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcinoctets interface{}

    // The number of input IP datagrams discarded due to errors in their IP
    // headers, including version number mismatch, other format errors, hop count
    // exceeded, errors discovered in processing their IP options, etc. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsinhdrerrors interface{}

    // The number of input IP datagrams discarded because no route could be found
    // to transmit them to their destination.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsinnoroutes interface{}

    // The number of input IP datagrams discarded because the IP address in their
    // IP header's destination field was not a valid address to be received at
    // this entity.  This count includes invalid addresses (e.g., ::0).  For
    // entities that are not IP routers and therefore do not forward datagrams,
    // this counter includes datagrams discarded because the destination address
    // was not a local address.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipifstatsinaddrerrors interface{}

    // The number of locally-addressed IP datagrams received successfully but
    // discarded because of an unknown or unsupported protocol.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of   ipIfStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ipifstatsinunknownprotos interface{}

    // The number of input IP datagrams discarded because the datagram frame
    // didn't carry enough data.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipifstatsintruncatedpkts interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination and for which this entity attempted to find a route to forward
    // them to that final destination.  In entities that do not act as IP routers,
    // this counter will include only those datagrams that were Source-Routed via
    // this entity, and the Source-Route processing was successful.  When tracking
    // interface statistics, the counter of the incoming interface is incremented
    // for each datagram.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipifstatsinforwdatagrams interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination and for which this entity attempted to find a route to forward
    // them to that final destination.  This object counts the same packets as  
    // ipIfStatsInForwDatagrams, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcinforwdatagrams interface{}

    // The number of IP fragments received that needed to be reassembled at this
    // interface.  When tracking interface statistics, the counter of the
    // interface to which these fragments were addressed is incremented.  This
    // interface might not be the same as the input interface for some of the
    // fragments.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ipifstatsreasmreqds interface{}

    // The number of IP datagrams successfully reassembled.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsreasmoks interface{}

    // The number of failures detected by the IP re-assembly algorithm (for
    // whatever reason: timed out, errors, etc.). Note that this is not
    // necessarily a count of discarded IP fragments since some algorithms
    // (notably the algorithm in RFC 815) can lose track of the number of
    // fragments by combining them as they are received.  When tracking interface
    // statistics, the counter of the interface to which these fragments were
    // addressed is incremented.  This interface might not be the same as the
    // input interface for some of the fragments.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsreasmfails interface{}

    // The number of input IP datagrams for which no problems were encountered to
    // prevent their continued processing, but were discarded (e.g., for lack of
    // buffer space).  Note that this counter does not include any datagrams
    // discarded while awaiting re-assembly.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsindiscards interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  When tracking interface statistics, the counter of the
    // interface to which these datagrams were addressed is incremented.  This
    // interface might not be the same as the   input interface for some of the
    // datagrams.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ipifstatsindelivers interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  This object counts the same packets as
    // ipIfStatsInDelivers, but allows for larger values.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcindelivers interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  Note that this counter
    // does not include any datagrams counted in ipIfStatsOutForwDatagrams. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsoutrequests interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  This object counts the
    // same packets as   ipIfStatsOutRequests, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcoutrequests interface{}

    // The number of datagrams for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.  In entities that do not act as IP routers, this counter
    // will include only those datagrams that were Source-Routed via this entity,
    // and the Source-Route processing was successful.  When tracking interface
    // statistics, the counter of the outgoing interface is incremented for a
    // successfully forwarded datagram.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsoutforwdatagrams interface{}

    // The number of datagrams for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.  This object counts the same packets as
    // ipIfStatsOutForwDatagrams, but allows for larger values.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of  
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcoutforwdatagrams interface{}

    // The number of output IP datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but were discarded (e.g.,
    // for lack of buffer space).  Note that this counter would include datagrams
    // counted in ipIfStatsOutForwDatagrams if any such datagrams met this
    // (discretionary) discard criterion.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsoutdiscards interface{}

    // The number of IP datagrams that would require fragmentation in order to be
    // transmitted.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for a successfully fragmented datagram. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsoutfragreqds interface{}

    // The number of IP datagrams that have been successfully fragmented.  When
    // tracking interface statistics, the counter of the   outgoing interface is
    // incremented for a successfully fragmented datagram.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsoutfragoks interface{}

    // The number of IP datagrams that have been discarded because they needed to
    // be fragmented but could not be.  This includes IPv4 packets that have the
    // DF bit set and IPv6 packets that are being forwarded and exceed the
    // outgoing link MTU.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for an unsuccessfully fragmented
    // datagram.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ipifstatsoutfragfails interface{}

    // The number of output datagram fragments that have been generated as a
    // result of IP fragmentation.  When tracking interface statistics, the
    // counter of the outgoing interface is incremented for a successfully
    // fragmented datagram.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipifstatsoutfragcreates interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This includes datagrams generated locally and
    // those forwarded by this entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsouttransmits interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This object counts the same datagrams as
    // ipIfStatsOutTransmits, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcouttransmits interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  Octets from datagrams counted in ipIfStatsOutTransmits
    // MUST be counted here.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ipifstatsoutoctets interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  This objects counts the same octets as
    // ipIfStatsOutOctets, but allows for larger values.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcoutoctets interface{}

    // The number of IP multicast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsinmcastpkts interface{}

    // The number of IP multicast datagrams received.  This object counts the same
    // datagrams as ipIfStatsInMcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcinmcastpkts interface{}

    // The total number of octets received in IP multicast   datagrams.  Octets
    // from datagrams counted in ipIfStatsInMcastPkts MUST be counted here. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsinmcastoctets interface{}

    // The total number of octets received in IP multicast datagrams.  This object
    // counts the same octets as ipIfStatsInMcastOctets, but allows for larger
    // values.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    Ipifstatshcinmcastoctets interface{}

    // The number of IP multicast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsoutmcastpkts interface{}

    // The number of IP multicast datagrams transmitted.  This object counts the
    // same datagrams as ipIfStatsOutMcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other   times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcoutmcastpkts interface{}

    // The total number of octets transmitted in IP multicast datagrams.  Octets
    // from datagrams counted in ipIfStatsOutMcastPkts MUST be counted here. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsoutmcastoctets interface{}

    // The total number of octets transmitted in IP multicast datagrams.  This
    // object counts the same octets as ipIfStatsOutMcastOctets, but allows for
    // larger values.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    Ipifstatshcoutmcastoctets interface{}

    // The number of IP broadcast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsinbcastpkts interface{}

    // The number of IP broadcast datagrams received.  This object counts the same
    // datagrams as ipIfStatsInBcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcinbcastpkts interface{}

    // The number of IP broadcast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ipifstatsoutbcastpkts interface{}

    // The number of IP broadcast datagrams transmitted.  This object counts the
    // same datagrams as ipIfStatsOutBcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipifstatshcoutbcastpkts interface{}

    // The value of sysUpTime on the most recent occasion at which   any one or
    // more of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Ipifstatsdiscontinuitytime interface{}

    // The minimum reasonable polling interval for this entry. This object
    // provides an indication of the minimum amount of time required to update the
    // counters in this entry. The type is interface{} with range: 0..4294967295.
    // Units are milli-seconds.
    Ipifstatsrefreshrate interface{}
}

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetFilter() yfilter.YFilter { return ipifstatsentry.YFilter }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) SetFilter(yf yfilter.YFilter) { ipifstatsentry.YFilter = yf }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetGoName(yname string) string {
    if yname == "ipIfStatsIPVersion" { return "Ipifstatsipversion" }
    if yname == "ipIfStatsIfIndex" { return "Ipifstatsifindex" }
    if yname == "ipIfStatsInReceives" { return "Ipifstatsinreceives" }
    if yname == "ipIfStatsHCInReceives" { return "Ipifstatshcinreceives" }
    if yname == "ipIfStatsInOctets" { return "Ipifstatsinoctets" }
    if yname == "ipIfStatsHCInOctets" { return "Ipifstatshcinoctets" }
    if yname == "ipIfStatsInHdrErrors" { return "Ipifstatsinhdrerrors" }
    if yname == "ipIfStatsInNoRoutes" { return "Ipifstatsinnoroutes" }
    if yname == "ipIfStatsInAddrErrors" { return "Ipifstatsinaddrerrors" }
    if yname == "ipIfStatsInUnknownProtos" { return "Ipifstatsinunknownprotos" }
    if yname == "ipIfStatsInTruncatedPkts" { return "Ipifstatsintruncatedpkts" }
    if yname == "ipIfStatsInForwDatagrams" { return "Ipifstatsinforwdatagrams" }
    if yname == "ipIfStatsHCInForwDatagrams" { return "Ipifstatshcinforwdatagrams" }
    if yname == "ipIfStatsReasmReqds" { return "Ipifstatsreasmreqds" }
    if yname == "ipIfStatsReasmOKs" { return "Ipifstatsreasmoks" }
    if yname == "ipIfStatsReasmFails" { return "Ipifstatsreasmfails" }
    if yname == "ipIfStatsInDiscards" { return "Ipifstatsindiscards" }
    if yname == "ipIfStatsInDelivers" { return "Ipifstatsindelivers" }
    if yname == "ipIfStatsHCInDelivers" { return "Ipifstatshcindelivers" }
    if yname == "ipIfStatsOutRequests" { return "Ipifstatsoutrequests" }
    if yname == "ipIfStatsHCOutRequests" { return "Ipifstatshcoutrequests" }
    if yname == "ipIfStatsOutForwDatagrams" { return "Ipifstatsoutforwdatagrams" }
    if yname == "ipIfStatsHCOutForwDatagrams" { return "Ipifstatshcoutforwdatagrams" }
    if yname == "ipIfStatsOutDiscards" { return "Ipifstatsoutdiscards" }
    if yname == "ipIfStatsOutFragReqds" { return "Ipifstatsoutfragreqds" }
    if yname == "ipIfStatsOutFragOKs" { return "Ipifstatsoutfragoks" }
    if yname == "ipIfStatsOutFragFails" { return "Ipifstatsoutfragfails" }
    if yname == "ipIfStatsOutFragCreates" { return "Ipifstatsoutfragcreates" }
    if yname == "ipIfStatsOutTransmits" { return "Ipifstatsouttransmits" }
    if yname == "ipIfStatsHCOutTransmits" { return "Ipifstatshcouttransmits" }
    if yname == "ipIfStatsOutOctets" { return "Ipifstatsoutoctets" }
    if yname == "ipIfStatsHCOutOctets" { return "Ipifstatshcoutoctets" }
    if yname == "ipIfStatsInMcastPkts" { return "Ipifstatsinmcastpkts" }
    if yname == "ipIfStatsHCInMcastPkts" { return "Ipifstatshcinmcastpkts" }
    if yname == "ipIfStatsInMcastOctets" { return "Ipifstatsinmcastoctets" }
    if yname == "ipIfStatsHCInMcastOctets" { return "Ipifstatshcinmcastoctets" }
    if yname == "ipIfStatsOutMcastPkts" { return "Ipifstatsoutmcastpkts" }
    if yname == "ipIfStatsHCOutMcastPkts" { return "Ipifstatshcoutmcastpkts" }
    if yname == "ipIfStatsOutMcastOctets" { return "Ipifstatsoutmcastoctets" }
    if yname == "ipIfStatsHCOutMcastOctets" { return "Ipifstatshcoutmcastoctets" }
    if yname == "ipIfStatsInBcastPkts" { return "Ipifstatsinbcastpkts" }
    if yname == "ipIfStatsHCInBcastPkts" { return "Ipifstatshcinbcastpkts" }
    if yname == "ipIfStatsOutBcastPkts" { return "Ipifstatsoutbcastpkts" }
    if yname == "ipIfStatsHCOutBcastPkts" { return "Ipifstatshcoutbcastpkts" }
    if yname == "ipIfStatsDiscontinuityTime" { return "Ipifstatsdiscontinuitytime" }
    if yname == "ipIfStatsRefreshRate" { return "Ipifstatsrefreshrate" }
    return ""
}

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetSegmentPath() string {
    return "ipIfStatsEntry" + "[ipIfStatsIPVersion='" + fmt.Sprintf("%v", ipifstatsentry.Ipifstatsipversion) + "']" + "[ipIfStatsIfIndex='" + fmt.Sprintf("%v", ipifstatsentry.Ipifstatsifindex) + "']"
}

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipIfStatsIPVersion"] = ipifstatsentry.Ipifstatsipversion
    leafs["ipIfStatsIfIndex"] = ipifstatsentry.Ipifstatsifindex
    leafs["ipIfStatsInReceives"] = ipifstatsentry.Ipifstatsinreceives
    leafs["ipIfStatsHCInReceives"] = ipifstatsentry.Ipifstatshcinreceives
    leafs["ipIfStatsInOctets"] = ipifstatsentry.Ipifstatsinoctets
    leafs["ipIfStatsHCInOctets"] = ipifstatsentry.Ipifstatshcinoctets
    leafs["ipIfStatsInHdrErrors"] = ipifstatsentry.Ipifstatsinhdrerrors
    leafs["ipIfStatsInNoRoutes"] = ipifstatsentry.Ipifstatsinnoroutes
    leafs["ipIfStatsInAddrErrors"] = ipifstatsentry.Ipifstatsinaddrerrors
    leafs["ipIfStatsInUnknownProtos"] = ipifstatsentry.Ipifstatsinunknownprotos
    leafs["ipIfStatsInTruncatedPkts"] = ipifstatsentry.Ipifstatsintruncatedpkts
    leafs["ipIfStatsInForwDatagrams"] = ipifstatsentry.Ipifstatsinforwdatagrams
    leafs["ipIfStatsHCInForwDatagrams"] = ipifstatsentry.Ipifstatshcinforwdatagrams
    leafs["ipIfStatsReasmReqds"] = ipifstatsentry.Ipifstatsreasmreqds
    leafs["ipIfStatsReasmOKs"] = ipifstatsentry.Ipifstatsreasmoks
    leafs["ipIfStatsReasmFails"] = ipifstatsentry.Ipifstatsreasmfails
    leafs["ipIfStatsInDiscards"] = ipifstatsentry.Ipifstatsindiscards
    leafs["ipIfStatsInDelivers"] = ipifstatsentry.Ipifstatsindelivers
    leafs["ipIfStatsHCInDelivers"] = ipifstatsentry.Ipifstatshcindelivers
    leafs["ipIfStatsOutRequests"] = ipifstatsentry.Ipifstatsoutrequests
    leafs["ipIfStatsHCOutRequests"] = ipifstatsentry.Ipifstatshcoutrequests
    leafs["ipIfStatsOutForwDatagrams"] = ipifstatsentry.Ipifstatsoutforwdatagrams
    leafs["ipIfStatsHCOutForwDatagrams"] = ipifstatsentry.Ipifstatshcoutforwdatagrams
    leafs["ipIfStatsOutDiscards"] = ipifstatsentry.Ipifstatsoutdiscards
    leafs["ipIfStatsOutFragReqds"] = ipifstatsentry.Ipifstatsoutfragreqds
    leafs["ipIfStatsOutFragOKs"] = ipifstatsentry.Ipifstatsoutfragoks
    leafs["ipIfStatsOutFragFails"] = ipifstatsentry.Ipifstatsoutfragfails
    leafs["ipIfStatsOutFragCreates"] = ipifstatsentry.Ipifstatsoutfragcreates
    leafs["ipIfStatsOutTransmits"] = ipifstatsentry.Ipifstatsouttransmits
    leafs["ipIfStatsHCOutTransmits"] = ipifstatsentry.Ipifstatshcouttransmits
    leafs["ipIfStatsOutOctets"] = ipifstatsentry.Ipifstatsoutoctets
    leafs["ipIfStatsHCOutOctets"] = ipifstatsentry.Ipifstatshcoutoctets
    leafs["ipIfStatsInMcastPkts"] = ipifstatsentry.Ipifstatsinmcastpkts
    leafs["ipIfStatsHCInMcastPkts"] = ipifstatsentry.Ipifstatshcinmcastpkts
    leafs["ipIfStatsInMcastOctets"] = ipifstatsentry.Ipifstatsinmcastoctets
    leafs["ipIfStatsHCInMcastOctets"] = ipifstatsentry.Ipifstatshcinmcastoctets
    leafs["ipIfStatsOutMcastPkts"] = ipifstatsentry.Ipifstatsoutmcastpkts
    leafs["ipIfStatsHCOutMcastPkts"] = ipifstatsentry.Ipifstatshcoutmcastpkts
    leafs["ipIfStatsOutMcastOctets"] = ipifstatsentry.Ipifstatsoutmcastoctets
    leafs["ipIfStatsHCOutMcastOctets"] = ipifstatsentry.Ipifstatshcoutmcastoctets
    leafs["ipIfStatsInBcastPkts"] = ipifstatsentry.Ipifstatsinbcastpkts
    leafs["ipIfStatsHCInBcastPkts"] = ipifstatsentry.Ipifstatshcinbcastpkts
    leafs["ipIfStatsOutBcastPkts"] = ipifstatsentry.Ipifstatsoutbcastpkts
    leafs["ipIfStatsHCOutBcastPkts"] = ipifstatsentry.Ipifstatshcoutbcastpkts
    leafs["ipIfStatsDiscontinuityTime"] = ipifstatsentry.Ipifstatsdiscontinuitytime
    leafs["ipIfStatsRefreshRate"] = ipifstatsentry.Ipifstatsrefreshrate
    return leafs
}

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetYangName() string { return "ipIfStatsEntry" }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) SetParent(parent types.Entity) { ipifstatsentry.parent = parent }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetParent() types.Entity { return ipifstatsentry.parent }

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetParentYangName() string { return "ipIfStatsTable" }

// IPMIB_Ipaddressprefixtable
// This table allows the user to determine the source of an IP
// address or set of IP addresses, and allows other tables to
// share the information via pointer rather than by copying.
// 
// For example, when the node configures both a unicast and
// anycast address for a prefix, the ipAddressPrefix objects
// for those addresses will point to a single row in this
// table.
// 
// This table primarily provides support for IPv6 prefixes, and
// several of the objects are less meaningful for IPv4.  The
// table continues to allow IPv4 addresses to allow future
// flexibility.  In order to promote a common configuration,
// this document includes suggestions for default values for
// IPv4 prefixes.  Each of these values may be overridden if an
// object is meaningful to the node.
// 
// All prefixes used by this entity should be included in this
// table independent of how the entity learned the prefix.
// (This table isn't limited to prefixes learned from router
// 
// 
// advertisements.)
type IPMIB_Ipaddressprefixtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the ipAddressPrefixTable. The type is slice of
    // IPMIB_Ipaddressprefixtable_Ipaddressprefixentry.
    Ipaddressprefixentry []IPMIB_Ipaddressprefixtable_Ipaddressprefixentry
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetFilter() yfilter.YFilter { return ipaddressprefixtable.YFilter }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) SetFilter(yf yfilter.YFilter) { ipaddressprefixtable.YFilter = yf }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetGoName(yname string) string {
    if yname == "ipAddressPrefixEntry" { return "Ipaddressprefixentry" }
    return ""
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetSegmentPath() string {
    return "ipAddressPrefixTable"
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipAddressPrefixEntry" {
        for _, c := range ipaddressprefixtable.Ipaddressprefixentry {
            if ipaddressprefixtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipaddressprefixtable_Ipaddressprefixentry{}
        ipaddressprefixtable.Ipaddressprefixentry = append(ipaddressprefixtable.Ipaddressprefixentry, child)
        return &ipaddressprefixtable.Ipaddressprefixentry[len(ipaddressprefixtable.Ipaddressprefixentry)-1]
    }
    return nil
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipaddressprefixtable.Ipaddressprefixentry {
        children[ipaddressprefixtable.Ipaddressprefixentry[i].GetSegmentPath()] = &ipaddressprefixtable.Ipaddressprefixentry[i]
    }
    return children
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetYangName() string { return "ipAddressPrefixTable" }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) SetParent(parent types.Entity) { ipaddressprefixtable.parent = parent }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetParent() types.Entity { return ipaddressprefixtable.parent }

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipaddressprefixtable_Ipaddressprefixentry
// An entry in the ipAddressPrefixTable.
type IPMIB_Ipaddressprefixtable_Ipaddressprefixentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value that uniquely identifies the
    // interface on which this prefix is configured.  The interface identified by
    // a particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipaddressprefixifindex interface{}

    // This attribute is a key. The address type of ipAddressPrefix. The type is
    // InetAddressType.
    Ipaddressprefixtype interface{}

    // This attribute is a key. The address prefix.  The address type of this
    // object is specified in ipAddressPrefixType.  The length of this object is
    // the standard length for objects of that type (4 or 16 bytes).  Any bits
    // after ipAddressPrefixLength must be zero.  Implementors need to be aware
    // that, if the size of ipAddressPrefixPrefix exceeds 114 octets, then OIDS of
    // instances of columns in this row will have more than 128 sub-identifiers
    // and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3. The type is string
    // with length: 0..255.
    Ipaddressprefixprefix interface{}

    // This attribute is a key. The prefix length associated with this prefix. 
    // The value 0 has no special meaning for this object.  It simply refers to
    // address '::/0'. The type is interface{} with range: 0..2040.
    Ipaddressprefixlength interface{}

    // The origin of this prefix. The type is IpAddressPrefixOriginTC.
    Ipaddressprefixorigin interface{}

    // This object has the value 'true(1)', if this prefix can be used for on-link
    // determination; otherwise, the value is 'false(2)'.  The default for IPv4
    // prefixes is 'true(1)'. The type is bool.
    Ipaddressprefixonlinkflag interface{}

    // Autonomous address configuration flag.  When true(1), indicates that this
    // prefix can be used for autonomous address configuration (i.e., can be used
    // to form a local interface address).  If false(2), it is not used to auto-
    // configure a local interface address.  The default for IPv4 prefixes is
    // 'false(2)'. The type is bool.
    Ipaddressprefixautonomousflag interface{}

    // The remaining length of time, in seconds, that this prefix will continue to
    // be preferred, i.e., time until deprecation.  A value of 4,294,967,295
    // represents infinity.  The address generated from a deprecated prefix should
    // no longer be used as a source address in new communications, but packets
    // received on such an interface are processed as expected.  The default for
    // IPv4 prefixes is 4,294,967,295 (infinity). The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Ipaddressprefixadvpreferredlifetime interface{}

    // The remaining length of time, in seconds, that this prefix will continue to
    // be valid, i.e., time until invalidation.  A value of 4,294,967,295
    // represents infinity.  The address generated from an invalidated prefix
    // should not appear as the destination or source address of a packet.   The
    // default for IPv4 prefixes is 4,294,967,295 (infinity). The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Ipaddressprefixadvvalidlifetime interface{}
}

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetFilter() yfilter.YFilter { return ipaddressprefixentry.YFilter }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) SetFilter(yf yfilter.YFilter) { ipaddressprefixentry.YFilter = yf }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetGoName(yname string) string {
    if yname == "ipAddressPrefixIfIndex" { return "Ipaddressprefixifindex" }
    if yname == "ipAddressPrefixType" { return "Ipaddressprefixtype" }
    if yname == "ipAddressPrefixPrefix" { return "Ipaddressprefixprefix" }
    if yname == "ipAddressPrefixLength" { return "Ipaddressprefixlength" }
    if yname == "ipAddressPrefixOrigin" { return "Ipaddressprefixorigin" }
    if yname == "ipAddressPrefixOnLinkFlag" { return "Ipaddressprefixonlinkflag" }
    if yname == "ipAddressPrefixAutonomousFlag" { return "Ipaddressprefixautonomousflag" }
    if yname == "ipAddressPrefixAdvPreferredLifetime" { return "Ipaddressprefixadvpreferredlifetime" }
    if yname == "ipAddressPrefixAdvValidLifetime" { return "Ipaddressprefixadvvalidlifetime" }
    return ""
}

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetSegmentPath() string {
    return "ipAddressPrefixEntry" + "[ipAddressPrefixIfIndex='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixifindex) + "']" + "[ipAddressPrefixType='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixtype) + "']" + "[ipAddressPrefixPrefix='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixprefix) + "']" + "[ipAddressPrefixLength='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixlength) + "']"
}

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipAddressPrefixIfIndex"] = ipaddressprefixentry.Ipaddressprefixifindex
    leafs["ipAddressPrefixType"] = ipaddressprefixentry.Ipaddressprefixtype
    leafs["ipAddressPrefixPrefix"] = ipaddressprefixentry.Ipaddressprefixprefix
    leafs["ipAddressPrefixLength"] = ipaddressprefixentry.Ipaddressprefixlength
    leafs["ipAddressPrefixOrigin"] = ipaddressprefixentry.Ipaddressprefixorigin
    leafs["ipAddressPrefixOnLinkFlag"] = ipaddressprefixentry.Ipaddressprefixonlinkflag
    leafs["ipAddressPrefixAutonomousFlag"] = ipaddressprefixentry.Ipaddressprefixautonomousflag
    leafs["ipAddressPrefixAdvPreferredLifetime"] = ipaddressprefixentry.Ipaddressprefixadvpreferredlifetime
    leafs["ipAddressPrefixAdvValidLifetime"] = ipaddressprefixentry.Ipaddressprefixadvvalidlifetime
    return leafs
}

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetYangName() string { return "ipAddressPrefixEntry" }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) SetParent(parent types.Entity) { ipaddressprefixentry.parent = parent }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetParent() types.Entity { return ipaddressprefixentry.parent }

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetParentYangName() string { return "ipAddressPrefixTable" }

// IPMIB_Ipaddresstable
// This table contains addressing information relevant to the
// entity's interfaces.
// 
// This table does not contain multicast address information.
// Tables for such information should be contained in multicast
// specific MIBs, such as RFC 3019.
// 
// While this table is writable, the user will note that
// several objects, such as ipAddressOrigin, are not.  The
// intention in allowing a user to write to this table is to
// allow them to add or remove any entry that isn't
// 
// 
// permanent.  The user should be allowed to modify objects
// and entries when that would not cause inconsistencies
// within the table.  Allowing write access to objects, such
// as ipAddressOrigin, could allow a user to insert an entry
// and then label it incorrectly.
// 
// Note well: When including IPv6 link-local addresses in this
// table, the entry must use an InetAddressType of 'ipv6z' in
// order to differentiate between the possible interfaces.
type IPMIB_Ipaddresstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An address mapping for a particular interface. The type is slice of
    // IPMIB_Ipaddresstable_Ipaddressentry.
    Ipaddressentry []IPMIB_Ipaddresstable_Ipaddressentry
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetFilter() yfilter.YFilter { return ipaddresstable.YFilter }

func (ipaddresstable *IPMIB_Ipaddresstable) SetFilter(yf yfilter.YFilter) { ipaddresstable.YFilter = yf }

func (ipaddresstable *IPMIB_Ipaddresstable) GetGoName(yname string) string {
    if yname == "ipAddressEntry" { return "Ipaddressentry" }
    return ""
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetSegmentPath() string {
    return "ipAddressTable"
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipAddressEntry" {
        for _, c := range ipaddresstable.Ipaddressentry {
            if ipaddresstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipaddresstable_Ipaddressentry{}
        ipaddresstable.Ipaddressentry = append(ipaddresstable.Ipaddressentry, child)
        return &ipaddresstable.Ipaddressentry[len(ipaddresstable.Ipaddressentry)-1]
    }
    return nil
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipaddresstable.Ipaddressentry {
        children[ipaddresstable.Ipaddressentry[i].GetSegmentPath()] = &ipaddresstable.Ipaddressentry[i]
    }
    return children
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddresstable *IPMIB_Ipaddresstable) GetYangName() string { return "ipAddressTable" }

func (ipaddresstable *IPMIB_Ipaddresstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddresstable *IPMIB_Ipaddresstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddresstable *IPMIB_Ipaddresstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddresstable *IPMIB_Ipaddresstable) SetParent(parent types.Entity) { ipaddresstable.parent = parent }

func (ipaddresstable *IPMIB_Ipaddresstable) GetParent() types.Entity { return ipaddresstable.parent }

func (ipaddresstable *IPMIB_Ipaddresstable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipaddresstable_Ipaddressentry
// An address mapping for a particular interface.
type IPMIB_Ipaddresstable_Ipaddressentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The address type of ipAddressAddr. The type is
    // InetAddressType.
    Ipaddressaddrtype interface{}

    // This attribute is a key. The IP address to which this entry's addressing
    // information   pertains.  The address type of this object is specified in
    // ipAddressAddrType.  Implementors need to be aware that if the size of
    // ipAddressAddr exceeds 116 octets, then OIDS of instances of columns in this
    // row will have more than 128 sub-identifiers and cannot be accessed using
    // SNMPv1, SNMPv2c, or SNMPv3. The type is string with length: 0..255.
    Ipaddressaddr interface{}

    // The index value that uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of the IF-MIB's
    // ifIndex. The type is interface{} with range: 1..2147483647.
    Ipaddressifindex interface{}

    // The type of address.  broadcast(3) is not a valid value for IPv6 addresses
    // (RFC 3513). The type is Ipaddresstype.
    Ipaddresstype interface{}

    // A pointer to the row in the prefix table to which this address belongs. 
    // May be { 0 0 } if there is no such row. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Ipaddressprefix interface{}

    // The origin of the address. The type is IpAddressOriginTC.
    Ipaddressorigin interface{}

    // The status of the address, describing if the address can be used for
    // communication.  In the absence of other information, an IPv4 address is
    // always preferred(1). The type is IpAddressStatusTC.
    Ipaddressstatus interface{}

    // The value of sysUpTime at the time this entry was created. If this entry
    // was created prior to the last re- initialization of the local network
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Ipaddresscreated interface{}

    // The value of sysUpTime at the time this entry was last updated.  If this
    // entry was updated prior to the last re- initialization of the local network
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Ipaddresslastchanged interface{}

    // The status of this conceptual row.  The RowStatus TC requires that this
    // DESCRIPTION clause states under which circumstances other objects in this
    // row   can be modified.  The value of this object has no effect on whether
    // other objects in this conceptual row can be modified.  A conceptual row can
    // not be made active until the ipAddressIfIndex has been set to a valid
    // index. The type is RowStatus.
    Ipaddressrowstatus interface{}

    // The storage type for this conceptual row.  If this object has a value of
    // 'permanent', then no other objects are required to be able to be modified.
    // The type is StorageType.
    Ipaddressstoragetype interface{}
}

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetFilter() yfilter.YFilter { return ipaddressentry.YFilter }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) SetFilter(yf yfilter.YFilter) { ipaddressentry.YFilter = yf }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetGoName(yname string) string {
    if yname == "ipAddressAddrType" { return "Ipaddressaddrtype" }
    if yname == "ipAddressAddr" { return "Ipaddressaddr" }
    if yname == "ipAddressIfIndex" { return "Ipaddressifindex" }
    if yname == "ipAddressType" { return "Ipaddresstype" }
    if yname == "ipAddressPrefix" { return "Ipaddressprefix" }
    if yname == "ipAddressOrigin" { return "Ipaddressorigin" }
    if yname == "ipAddressStatus" { return "Ipaddressstatus" }
    if yname == "ipAddressCreated" { return "Ipaddresscreated" }
    if yname == "ipAddressLastChanged" { return "Ipaddresslastchanged" }
    if yname == "ipAddressRowStatus" { return "Ipaddressrowstatus" }
    if yname == "ipAddressStorageType" { return "Ipaddressstoragetype" }
    return ""
}

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetSegmentPath() string {
    return "ipAddressEntry" + "[ipAddressAddrType='" + fmt.Sprintf("%v", ipaddressentry.Ipaddressaddrtype) + "']" + "[ipAddressAddr='" + fmt.Sprintf("%v", ipaddressentry.Ipaddressaddr) + "']"
}

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipAddressAddrType"] = ipaddressentry.Ipaddressaddrtype
    leafs["ipAddressAddr"] = ipaddressentry.Ipaddressaddr
    leafs["ipAddressIfIndex"] = ipaddressentry.Ipaddressifindex
    leafs["ipAddressType"] = ipaddressentry.Ipaddresstype
    leafs["ipAddressPrefix"] = ipaddressentry.Ipaddressprefix
    leafs["ipAddressOrigin"] = ipaddressentry.Ipaddressorigin
    leafs["ipAddressStatus"] = ipaddressentry.Ipaddressstatus
    leafs["ipAddressCreated"] = ipaddressentry.Ipaddresscreated
    leafs["ipAddressLastChanged"] = ipaddressentry.Ipaddresslastchanged
    leafs["ipAddressRowStatus"] = ipaddressentry.Ipaddressrowstatus
    leafs["ipAddressStorageType"] = ipaddressentry.Ipaddressstoragetype
    return leafs
}

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetYangName() string { return "ipAddressEntry" }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) SetParent(parent types.Entity) { ipaddressentry.parent = parent }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetParent() types.Entity { return ipaddressentry.parent }

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetParentYangName() string { return "ipAddressTable" }

// IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype represents IPv6 addresses (RFC 3513).
type IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype string

const (
    IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype_unicast IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype = "unicast"

    IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype_anycast IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype = "anycast"

    IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype_broadcast IPMIB_Ipaddresstable_Ipaddressentry_Ipaddresstype = "broadcast"
)

// IPMIB_Ipnettophysicaltable
// The IP Address Translation table used for mapping from IP
// addresses to physical addresses.
// 
// The Address Translation tables contain the IP address to
// 'physical' address equivalences.  Some interfaces do not use
// translation tables for determining address equivalences
// (e.g., DDN-X.25 has an algorithmic method); if all
// interfaces are of this type, then the Address Translation
// table is empty, i.e., has zero entries.
// 
// While many protocols may be used to populate this table, ARP
// and Neighbor Discovery are the most likely
// options.
type IPMIB_Ipnettophysicaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains one IP address to `physical' address equivalence. The
    // type is slice of IPMIB_Ipnettophysicaltable_Ipnettophysicalentry.
    Ipnettophysicalentry []IPMIB_Ipnettophysicaltable_Ipnettophysicalentry
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetFilter() yfilter.YFilter { return ipnettophysicaltable.YFilter }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) SetFilter(yf yfilter.YFilter) { ipnettophysicaltable.YFilter = yf }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetGoName(yname string) string {
    if yname == "ipNetToPhysicalEntry" { return "Ipnettophysicalentry" }
    return ""
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetSegmentPath() string {
    return "ipNetToPhysicalTable"
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipNetToPhysicalEntry" {
        for _, c := range ipnettophysicaltable.Ipnettophysicalentry {
            if ipnettophysicaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipnettophysicaltable_Ipnettophysicalentry{}
        ipnettophysicaltable.Ipnettophysicalentry = append(ipnettophysicaltable.Ipnettophysicalentry, child)
        return &ipnettophysicaltable.Ipnettophysicalentry[len(ipnettophysicaltable.Ipnettophysicalentry)-1]
    }
    return nil
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipnettophysicaltable.Ipnettophysicalentry {
        children[ipnettophysicaltable.Ipnettophysicalentry[i].GetSegmentPath()] = &ipnettophysicaltable.Ipnettophysicalentry[i]
    }
    return children
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetBundleName() string { return "cisco_ios_xe" }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetYangName() string { return "ipNetToPhysicalTable" }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) SetParent(parent types.Entity) { ipnettophysicaltable.parent = parent }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetParent() types.Entity { return ipnettophysicaltable.parent }

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipnettophysicaltable_Ipnettophysicalentry
// Each entry contains one IP address to `physical' address
// equivalence.
type IPMIB_Ipnettophysicaltable_Ipnettophysicalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipnettophysicalifindex interface{}

    // This attribute is a key. The type of ipNetToPhysicalNetAddress. The type is
    // InetAddressType.
    Ipnettophysicalnetaddresstype interface{}

    // This attribute is a key. The IP Address corresponding to the
    // media-dependent `physical' address.  The address type of this object is
    // specified in ipNetToPhysicalAddressType.  Implementors need to be aware
    // that if the size of   ipNetToPhysicalNetAddress exceeds 115 octets, then
    // OIDS of instances of columns in this row will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3.
    // The type is string with length: 0..255.
    Ipnettophysicalnetaddress interface{}

    // The media-dependent `physical' address.  As the entries in this table are
    // typically not persistent when this object is written the entity SHOULD NOT
    // save the change to non-volatile storage. The type is string with length:
    // 0..65535.
    Ipnettophysicalphysaddress interface{}

    // The value of sysUpTime at the time this entry was last updated.  If this
    // entry was updated prior to the last re- initialization of the local network
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Ipnettophysicallastupdated interface{}

    // The type of mapping.  Setting this object to the value invalid(2) has the
    // effect of invalidating the corresponding entry in the ipNetToPhysicalTable.
    // That is, it effectively dis- associates the interface identified with said
    // entry from the mapping identified with said entry.  It is an
    // implementation-specific matter as to whether the agent   removes an
    // invalidated entry from the table.  Accordingly, management stations must be
    // prepared to receive tabular information from agents that corresponds to
    // entries not currently in use.  Proper interpretation of such entries
    // requires examination of the relevant ipNetToPhysicalType object.  The
    // 'dynamic(3)' type indicates that the IP address to physical addresses
    // mapping has been dynamically resolved using e.g., IPv4 ARP or the IPv6
    // Neighbor Discovery protocol.  The 'static(4)' type indicates that the
    // mapping has been statically configured.  Both of these refer to entries
    // that provide mappings for other entities addresses.  The 'local(5)' type
    // indicates that the mapping is provided for an entity's own interface
    // address.  As the entries in this table are typically not persistent when
    // this object is written the entity SHOULD NOT save the change to
    // non-volatile storage. The type is Ipnettophysicaltype.
    Ipnettophysicaltype interface{}

    // The Neighbor Unreachability Detection state for the interface when the
    // address mapping in this entry is used. If Neighbor Unreachability Detection
    // is not in use (e.g. for IPv4), this object is always unknown(6). The type
    // is Ipnettophysicalstate.
    Ipnettophysicalstate interface{}

    // The status of this conceptual row.  The RowStatus TC requires that this
    // DESCRIPTION clause states under which circumstances other objects in this
    // row can be modified.  The value of this object has no effect on whether
    // other objects in this conceptual row can be modified.  A conceptual row can
    // not be made active until the ipNetToPhysicalPhysAddress object has been
    // set.  Note that if the ipNetToPhysicalType is set to 'invalid', the managed
    // node may delete the entry independent of the state of this object. The type
    // is RowStatus.
    Ipnettophysicalrowstatus interface{}
}

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetFilter() yfilter.YFilter { return ipnettophysicalentry.YFilter }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) SetFilter(yf yfilter.YFilter) { ipnettophysicalentry.YFilter = yf }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetGoName(yname string) string {
    if yname == "ipNetToPhysicalIfIndex" { return "Ipnettophysicalifindex" }
    if yname == "ipNetToPhysicalNetAddressType" { return "Ipnettophysicalnetaddresstype" }
    if yname == "ipNetToPhysicalNetAddress" { return "Ipnettophysicalnetaddress" }
    if yname == "ipNetToPhysicalPhysAddress" { return "Ipnettophysicalphysaddress" }
    if yname == "ipNetToPhysicalLastUpdated" { return "Ipnettophysicallastupdated" }
    if yname == "ipNetToPhysicalType" { return "Ipnettophysicaltype" }
    if yname == "ipNetToPhysicalState" { return "Ipnettophysicalstate" }
    if yname == "ipNetToPhysicalRowStatus" { return "Ipnettophysicalrowstatus" }
    return ""
}

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetSegmentPath() string {
    return "ipNetToPhysicalEntry" + "[ipNetToPhysicalIfIndex='" + fmt.Sprintf("%v", ipnettophysicalentry.Ipnettophysicalifindex) + "']" + "[ipNetToPhysicalNetAddressType='" + fmt.Sprintf("%v", ipnettophysicalentry.Ipnettophysicalnetaddresstype) + "']" + "[ipNetToPhysicalNetAddress='" + fmt.Sprintf("%v", ipnettophysicalentry.Ipnettophysicalnetaddress) + "']"
}

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipNetToPhysicalIfIndex"] = ipnettophysicalentry.Ipnettophysicalifindex
    leafs["ipNetToPhysicalNetAddressType"] = ipnettophysicalentry.Ipnettophysicalnetaddresstype
    leafs["ipNetToPhysicalNetAddress"] = ipnettophysicalentry.Ipnettophysicalnetaddress
    leafs["ipNetToPhysicalPhysAddress"] = ipnettophysicalentry.Ipnettophysicalphysaddress
    leafs["ipNetToPhysicalLastUpdated"] = ipnettophysicalentry.Ipnettophysicallastupdated
    leafs["ipNetToPhysicalType"] = ipnettophysicalentry.Ipnettophysicaltype
    leafs["ipNetToPhysicalState"] = ipnettophysicalentry.Ipnettophysicalstate
    leafs["ipNetToPhysicalRowStatus"] = ipnettophysicalentry.Ipnettophysicalrowstatus
    return leafs
}

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetYangName() string { return "ipNetToPhysicalEntry" }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) SetParent(parent types.Entity) { ipnettophysicalentry.parent = parent }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetParent() types.Entity { return ipnettophysicalentry.parent }

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetParentYangName() string { return "ipNetToPhysicalTable" }

// IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate represents IPv4), this object is always unknown(6).
type IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate string

const (
    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_reachable IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "reachable"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_stale IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "stale"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_delay IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "delay"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_probe IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "probe"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_invalid IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "invalid"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_unknown IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "unknown"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate_incomplete IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicalstate = "incomplete"
)

// IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype represents change to non-volatile storage.
type IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype string

const (
    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype_other IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype = "other"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype_invalid IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype = "invalid"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype_dynamic IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype = "dynamic"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype_static IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype = "static"

    IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype_local IPMIB_Ipnettophysicaltable_Ipnettophysicalentry_Ipnettophysicaltype = "local"
)

// IPMIB_Ipv6Scopezoneindextable
// The table used to describe IPv6 unicast and multicast scope
// zones.
// 
// For those objects that have names rather than numbers, the
// names were chosen to coincide with the names used in the
// IPv6 address architecture document. 
type IPMIB_Ipv6Scopezoneindextable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the list of scope identifiers on a given interface. The
    // type is slice of IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry.
    Ipv6Scopezoneindexentry []IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetFilter() yfilter.YFilter { return ipv6Scopezoneindextable.YFilter }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) SetFilter(yf yfilter.YFilter) { ipv6Scopezoneindextable.YFilter = yf }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetGoName(yname string) string {
    if yname == "ipv6ScopeZoneIndexEntry" { return "Ipv6Scopezoneindexentry" }
    return ""
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetSegmentPath() string {
    return "ipv6ScopeZoneIndexTable"
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6ScopeZoneIndexEntry" {
        for _, c := range ipv6Scopezoneindextable.Ipv6Scopezoneindexentry {
            if ipv6Scopezoneindextable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry{}
        ipv6Scopezoneindextable.Ipv6Scopezoneindexentry = append(ipv6Scopezoneindextable.Ipv6Scopezoneindexentry, child)
        return &ipv6Scopezoneindextable.Ipv6Scopezoneindexentry[len(ipv6Scopezoneindextable.Ipv6Scopezoneindexentry)-1]
    }
    return nil
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Scopezoneindextable.Ipv6Scopezoneindexentry {
        children[ipv6Scopezoneindextable.Ipv6Scopezoneindexentry[i].GetSegmentPath()] = &ipv6Scopezoneindextable.Ipv6Scopezoneindexentry[i]
    }
    return children
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetYangName() string { return "ipv6ScopeZoneIndexTable" }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) SetParent(parent types.Entity) { ipv6Scopezoneindextable.parent = parent }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetParent() types.Entity { return ipv6Scopezoneindextable.parent }

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry
// Each entry contains the list of scope identifiers on a given
// interface.
type IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which these scopes belong.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipv6Scopezoneindexifindex interface{}

    // The zone index for the link-local scope on this interface. The type is
    // interface{} with range: 0..4294967295.
    Ipv6Scopezoneindexlinklocal interface{}

    // The zone index for scope 3 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindex3 interface{}

    // The zone index for the admin-local scope on this interface. The type is
    // interface{} with range: 0..4294967295.
    Ipv6Scopezoneindexadminlocal interface{}

    // The zone index for the site-local scope on this interface. The type is
    // interface{} with range: 0..4294967295.
    Ipv6Scopezoneindexsitelocal interface{}

    // The zone index for scope 6 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindex6 interface{}

    // The zone index for scope 7 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindex7 interface{}

    // The zone index for the organization-local scope on this interface. The type
    // is interface{} with range: 0..4294967295.
    Ipv6Scopezoneindexorganizationlocal interface{}

    // The zone index for scope 9 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindex9 interface{}

    // The zone index for scope A on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindexa interface{}

    // The zone index for scope B on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindexb interface{}

    // The zone index for scope C on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindexc interface{}

    // The zone index for scope D on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6Scopezoneindexd interface{}
}

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetFilter() yfilter.YFilter { return ipv6Scopezoneindexentry.YFilter }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) SetFilter(yf yfilter.YFilter) { ipv6Scopezoneindexentry.YFilter = yf }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetGoName(yname string) string {
    if yname == "ipv6ScopeZoneIndexIfIndex" { return "Ipv6Scopezoneindexifindex" }
    if yname == "ipv6ScopeZoneIndexLinkLocal" { return "Ipv6Scopezoneindexlinklocal" }
    if yname == "ipv6ScopeZoneIndex3" { return "Ipv6Scopezoneindex3" }
    if yname == "ipv6ScopeZoneIndexAdminLocal" { return "Ipv6Scopezoneindexadminlocal" }
    if yname == "ipv6ScopeZoneIndexSiteLocal" { return "Ipv6Scopezoneindexsitelocal" }
    if yname == "ipv6ScopeZoneIndex6" { return "Ipv6Scopezoneindex6" }
    if yname == "ipv6ScopeZoneIndex7" { return "Ipv6Scopezoneindex7" }
    if yname == "ipv6ScopeZoneIndexOrganizationLocal" { return "Ipv6Scopezoneindexorganizationlocal" }
    if yname == "ipv6ScopeZoneIndex9" { return "Ipv6Scopezoneindex9" }
    if yname == "ipv6ScopeZoneIndexA" { return "Ipv6Scopezoneindexa" }
    if yname == "ipv6ScopeZoneIndexB" { return "Ipv6Scopezoneindexb" }
    if yname == "ipv6ScopeZoneIndexC" { return "Ipv6Scopezoneindexc" }
    if yname == "ipv6ScopeZoneIndexD" { return "Ipv6Scopezoneindexd" }
    return ""
}

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetSegmentPath() string {
    return "ipv6ScopeZoneIndexEntry" + "[ipv6ScopeZoneIndexIfIndex='" + fmt.Sprintf("%v", ipv6Scopezoneindexentry.Ipv6Scopezoneindexifindex) + "']"
}

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6ScopeZoneIndexIfIndex"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexifindex
    leafs["ipv6ScopeZoneIndexLinkLocal"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexlinklocal
    leafs["ipv6ScopeZoneIndex3"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindex3
    leafs["ipv6ScopeZoneIndexAdminLocal"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexadminlocal
    leafs["ipv6ScopeZoneIndexSiteLocal"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexsitelocal
    leafs["ipv6ScopeZoneIndex6"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindex6
    leafs["ipv6ScopeZoneIndex7"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindex7
    leafs["ipv6ScopeZoneIndexOrganizationLocal"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexorganizationlocal
    leafs["ipv6ScopeZoneIndex9"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindex9
    leafs["ipv6ScopeZoneIndexA"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexa
    leafs["ipv6ScopeZoneIndexB"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexb
    leafs["ipv6ScopeZoneIndexC"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexc
    leafs["ipv6ScopeZoneIndexD"] = ipv6Scopezoneindexentry.Ipv6Scopezoneindexd
    return leafs
}

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetYangName() string { return "ipv6ScopeZoneIndexEntry" }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) SetParent(parent types.Entity) { ipv6Scopezoneindexentry.parent = parent }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetParent() types.Entity { return ipv6Scopezoneindexentry.parent }

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetParentYangName() string { return "ipv6ScopeZoneIndexTable" }

// IPMIB_Ipdefaultroutertable
// The table used to describe the default routers known to this
// 
// 
// entity.
type IPMIB_Ipdefaultroutertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains information about a default router known to this
    // entity. The type is slice of
    // IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry.
    Ipdefaultrouterentry []IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetFilter() yfilter.YFilter { return ipdefaultroutertable.YFilter }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) SetFilter(yf yfilter.YFilter) { ipdefaultroutertable.YFilter = yf }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetGoName(yname string) string {
    if yname == "ipDefaultRouterEntry" { return "Ipdefaultrouterentry" }
    return ""
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetSegmentPath() string {
    return "ipDefaultRouterTable"
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipDefaultRouterEntry" {
        for _, c := range ipdefaultroutertable.Ipdefaultrouterentry {
            if ipdefaultroutertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry{}
        ipdefaultroutertable.Ipdefaultrouterentry = append(ipdefaultroutertable.Ipdefaultrouterentry, child)
        return &ipdefaultroutertable.Ipdefaultrouterentry[len(ipdefaultroutertable.Ipdefaultrouterentry)-1]
    }
    return nil
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipdefaultroutertable.Ipdefaultrouterentry {
        children[ipdefaultroutertable.Ipdefaultrouterentry[i].GetSegmentPath()] = &ipdefaultroutertable.Ipdefaultrouterentry[i]
    }
    return children
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetBundleName() string { return "cisco_ios_xe" }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetYangName() string { return "ipDefaultRouterTable" }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) SetParent(parent types.Entity) { ipdefaultroutertable.parent = parent }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetParent() types.Entity { return ipdefaultroutertable.parent }

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry
// Each entry contains information about a default router known
// to this entity.
type IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The address type for this row. The type is
    // InetAddressType.
    Ipdefaultrouteraddresstype interface{}

    // This attribute is a key. The IP address of the default router represented
    // by this row.  The address type of this object is specified in
    // ipDefaultRouterAddressType.  Implementers need to be aware that if the size
    // of ipDefaultRouterAddress exceeds 115 octets, then OIDS of instances of
    // columns in this row will have more than 128 sub-identifiers and cannot be
    // accessed using SNMPv1, SNMPv2c, or SNMPv3. The type is string with length:
    // 0..255.
    Ipdefaultrouteraddress interface{}

    // This attribute is a key. The index value that uniquely identifies the
    // interface by which the router can be reached.  The interface identified by
    // a particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipdefaultrouterifindex interface{}

    // The remaining length of time, in seconds, that this router will continue to
    // be useful as a default router.  A value of zero indicates that it is no
    // longer useful as a default router.  It is left to the implementer of the
    // MIB as to whether a router with a lifetime of zero is removed from the
    // list.  For IPv6, this value should be extracted from the router
    // advertisement messages. The type is interface{} with range: 0..65535. Units
    // are seconds.
    Ipdefaultrouterlifetime interface{}

    // An indication of preference given to this router as a default router as
    // described in he Default Router Preferences document.  Treating the value as
    // a 2 bit signed integer allows for simple arithmetic comparisons.  For IPv4
    // routers or IPv6 routers that are not using the updated router advertisement
    // format, this object is set to medium (0). The type is
    // Ipdefaultrouterpreference.
    Ipdefaultrouterpreference interface{}
}

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetFilter() yfilter.YFilter { return ipdefaultrouterentry.YFilter }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) SetFilter(yf yfilter.YFilter) { ipdefaultrouterentry.YFilter = yf }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetGoName(yname string) string {
    if yname == "ipDefaultRouterAddressType" { return "Ipdefaultrouteraddresstype" }
    if yname == "ipDefaultRouterAddress" { return "Ipdefaultrouteraddress" }
    if yname == "ipDefaultRouterIfIndex" { return "Ipdefaultrouterifindex" }
    if yname == "ipDefaultRouterLifetime" { return "Ipdefaultrouterlifetime" }
    if yname == "ipDefaultRouterPreference" { return "Ipdefaultrouterpreference" }
    return ""
}

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetSegmentPath() string {
    return "ipDefaultRouterEntry" + "[ipDefaultRouterAddressType='" + fmt.Sprintf("%v", ipdefaultrouterentry.Ipdefaultrouteraddresstype) + "']" + "[ipDefaultRouterAddress='" + fmt.Sprintf("%v", ipdefaultrouterentry.Ipdefaultrouteraddress) + "']" + "[ipDefaultRouterIfIndex='" + fmt.Sprintf("%v", ipdefaultrouterentry.Ipdefaultrouterifindex) + "']"
}

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipDefaultRouterAddressType"] = ipdefaultrouterentry.Ipdefaultrouteraddresstype
    leafs["ipDefaultRouterAddress"] = ipdefaultrouterentry.Ipdefaultrouteraddress
    leafs["ipDefaultRouterIfIndex"] = ipdefaultrouterentry.Ipdefaultrouterifindex
    leafs["ipDefaultRouterLifetime"] = ipdefaultrouterentry.Ipdefaultrouterlifetime
    leafs["ipDefaultRouterPreference"] = ipdefaultrouterentry.Ipdefaultrouterpreference
    return leafs
}

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetYangName() string { return "ipDefaultRouterEntry" }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) SetParent(parent types.Entity) { ipdefaultrouterentry.parent = parent }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetParent() types.Entity { return ipdefaultrouterentry.parent }

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetParentYangName() string { return "ipDefaultRouterTable" }

// IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference represents medium (0).
type IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference string

const (
    IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference_reserved IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference = "reserved"

    IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference_low IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference = "low"

    IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference_medium IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference = "medium"

    IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference_high IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry_Ipdefaultrouterpreference = "high"
)

// IPMIB_Ipv6Routeradverttable
// The table containing information used to construct router
// advertisements.
type IPMIB_Ipv6Routeradverttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing information used to construct router advertisements. 
    // Information in this table is persistent, and when this object is written,
    // the entity SHOULD save the change to non-volatile storage. The type is
    // slice of IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry.
    Ipv6Routeradvertentry []IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetFilter() yfilter.YFilter { return ipv6Routeradverttable.YFilter }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) SetFilter(yf yfilter.YFilter) { ipv6Routeradverttable.YFilter = yf }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetGoName(yname string) string {
    if yname == "ipv6RouterAdvertEntry" { return "Ipv6Routeradvertentry" }
    return ""
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetSegmentPath() string {
    return "ipv6RouterAdvertTable"
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6RouterAdvertEntry" {
        for _, c := range ipv6Routeradverttable.Ipv6Routeradvertentry {
            if ipv6Routeradverttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry{}
        ipv6Routeradverttable.Ipv6Routeradvertentry = append(ipv6Routeradverttable.Ipv6Routeradvertentry, child)
        return &ipv6Routeradverttable.Ipv6Routeradvertentry[len(ipv6Routeradverttable.Ipv6Routeradvertentry)-1]
    }
    return nil
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Routeradverttable.Ipv6Routeradvertentry {
        children[ipv6Routeradverttable.Ipv6Routeradvertentry[i].GetSegmentPath()] = &ipv6Routeradverttable.Ipv6Routeradvertentry[i]
    }
    return children
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetYangName() string { return "ipv6RouterAdvertTable" }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) SetParent(parent types.Entity) { ipv6Routeradverttable.parent = parent }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetParent() types.Entity { return ipv6Routeradverttable.parent }

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry
// An entry containing information used to construct router
// advertisements.
// 
// Information in this table is persistent, and when this
// object is written, the entity SHOULD save the change to
// non-volatile storage.
type IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value that uniquely identifies the
    // interface on which router advertisements constructed with this information
    // will be transmitted.  The interface identified by a particular value of
    // this index is the same interface as identified by the same value of the
    // IF-MIB's ifIndex. The type is interface{} with range: 1..2147483647.
    Ipv6Routeradvertifindex interface{}

    // A flag indicating whether the router sends periodic router advertisements
    // and responds to router solicitations on this interface. The type is bool.
    Ipv6Routeradvertsendadverts interface{}

    // The maximum time allowed between sending unsolicited router  
    // advertisements from this interface. The type is interface{} with range:
    // 4..1800. Units are seconds.
    Ipv6Routeradvertmaxinterval interface{}

    // The minimum time allowed between sending unsolicited router advertisements
    // from this interface.  The default is 0.33 * ipv6RouterAdvertMaxInterval,
    // however, in the case of a low value for ipv6RouterAdvertMaxInterval, the
    // minimum value for this object is restricted to 3. The type is interface{}
    // with range: 3..1350. Units are seconds.
    Ipv6Routeradvertmininterval interface{}

    // The true/false value to be placed into the 'managed address configuration'
    // flag field in router advertisements sent from this interface. The type is
    // bool.
    Ipv6Routeradvertmanagedflag interface{}

    // The true/false value to be placed into the 'other stateful configuration'
    // flag field in router advertisements sent from this interface. The type is
    // bool.
    Ipv6Routeradvertotherconfigflag interface{}

    // The value to be placed in MTU options sent by the router on this interface.
    // A value of zero indicates that no MTU options are sent. The type is
    // interface{} with range: 0..4294967295.
    Ipv6Routeradvertlinkmtu interface{}

    // The value to be placed in the reachable time field in router advertisement
    // messages sent from this interface.  A value of zero in the router
    // advertisement indicates that the advertisement isn't specifying a value for
    // reachable time. The type is interface{} with range: 0..3600000. Units are
    // milliseconds.
    Ipv6Routeradvertreachabletime interface{}

    // The value to be placed in the retransmit timer field in router
    // advertisements sent from this interface.  A value of zero in the router
    // advertisement indicates that the advertisement isn't specifying a value for
    // retrans time. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Ipv6Routeradvertretransmittime interface{}

    // The default value to be placed in the current hop limit field in router
    // advertisements sent from this interface.   The value should be set to the
    // current diameter of the Internet.  A value of zero in the router
    // advertisement indicates that the advertisement isn't specifying a value for
    // curHopLimit.  The default should be set to the value specified in the IANA
    // web pages (www.iana.org) at the time of implementation. The type is
    // interface{} with range: 0..255.
    Ipv6Routeradvertcurhoplimit interface{}

    // The value to be placed in the router lifetime field of router
    // advertisements sent from this interface.  This value MUST be either 0 or
    // between ipv6RouterAdvertMaxInterval and 9000 seconds.  A value of zero
    // indicates that the router is not to be used as a default router.  The
    // default is 3 * ipv6RouterAdvertMaxInterval. The type is interface{} with
    // range: 0..None | 4..9000. Units are seconds.
    Ipv6Routeradvertdefaultlifetime interface{}

    // The status of this conceptual row.  As all objects in this conceptual row
    // have default values, a row can be created and made active by setting this
    // object appropriately.  The RowStatus TC requires that this DESCRIPTION
    // clause states under which circumstances other objects in this row can be
    // modified.  The value of this object has no effect on whether other objects
    // in this conceptual row can be modified. The type is RowStatus.
    Ipv6Routeradvertrowstatus interface{}
}

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetFilter() yfilter.YFilter { return ipv6Routeradvertentry.YFilter }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) SetFilter(yf yfilter.YFilter) { ipv6Routeradvertentry.YFilter = yf }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetGoName(yname string) string {
    if yname == "ipv6RouterAdvertIfIndex" { return "Ipv6Routeradvertifindex" }
    if yname == "ipv6RouterAdvertSendAdverts" { return "Ipv6Routeradvertsendadverts" }
    if yname == "ipv6RouterAdvertMaxInterval" { return "Ipv6Routeradvertmaxinterval" }
    if yname == "ipv6RouterAdvertMinInterval" { return "Ipv6Routeradvertmininterval" }
    if yname == "ipv6RouterAdvertManagedFlag" { return "Ipv6Routeradvertmanagedflag" }
    if yname == "ipv6RouterAdvertOtherConfigFlag" { return "Ipv6Routeradvertotherconfigflag" }
    if yname == "ipv6RouterAdvertLinkMTU" { return "Ipv6Routeradvertlinkmtu" }
    if yname == "ipv6RouterAdvertReachableTime" { return "Ipv6Routeradvertreachabletime" }
    if yname == "ipv6RouterAdvertRetransmitTime" { return "Ipv6Routeradvertretransmittime" }
    if yname == "ipv6RouterAdvertCurHopLimit" { return "Ipv6Routeradvertcurhoplimit" }
    if yname == "ipv6RouterAdvertDefaultLifetime" { return "Ipv6Routeradvertdefaultlifetime" }
    if yname == "ipv6RouterAdvertRowStatus" { return "Ipv6Routeradvertrowstatus" }
    return ""
}

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetSegmentPath() string {
    return "ipv6RouterAdvertEntry" + "[ipv6RouterAdvertIfIndex='" + fmt.Sprintf("%v", ipv6Routeradvertentry.Ipv6Routeradvertifindex) + "']"
}

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6RouterAdvertIfIndex"] = ipv6Routeradvertentry.Ipv6Routeradvertifindex
    leafs["ipv6RouterAdvertSendAdverts"] = ipv6Routeradvertentry.Ipv6Routeradvertsendadverts
    leafs["ipv6RouterAdvertMaxInterval"] = ipv6Routeradvertentry.Ipv6Routeradvertmaxinterval
    leafs["ipv6RouterAdvertMinInterval"] = ipv6Routeradvertentry.Ipv6Routeradvertmininterval
    leafs["ipv6RouterAdvertManagedFlag"] = ipv6Routeradvertentry.Ipv6Routeradvertmanagedflag
    leafs["ipv6RouterAdvertOtherConfigFlag"] = ipv6Routeradvertentry.Ipv6Routeradvertotherconfigflag
    leafs["ipv6RouterAdvertLinkMTU"] = ipv6Routeradvertentry.Ipv6Routeradvertlinkmtu
    leafs["ipv6RouterAdvertReachableTime"] = ipv6Routeradvertentry.Ipv6Routeradvertreachabletime
    leafs["ipv6RouterAdvertRetransmitTime"] = ipv6Routeradvertentry.Ipv6Routeradvertretransmittime
    leafs["ipv6RouterAdvertCurHopLimit"] = ipv6Routeradvertentry.Ipv6Routeradvertcurhoplimit
    leafs["ipv6RouterAdvertDefaultLifetime"] = ipv6Routeradvertentry.Ipv6Routeradvertdefaultlifetime
    leafs["ipv6RouterAdvertRowStatus"] = ipv6Routeradvertentry.Ipv6Routeradvertrowstatus
    return leafs
}

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetYangName() string { return "ipv6RouterAdvertEntry" }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) SetParent(parent types.Entity) { ipv6Routeradvertentry.parent = parent }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetParent() types.Entity { return ipv6Routeradvertentry.parent }

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetParentYangName() string { return "ipv6RouterAdvertTable" }

// IPMIB_Icmpstatstable
// The table of generic system-wide ICMP counters.
type IPMIB_Icmpstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the icmpStatsTable. The type is slice of
    // IPMIB_Icmpstatstable_Icmpstatsentry.
    Icmpstatsentry []IPMIB_Icmpstatstable_Icmpstatsentry
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetFilter() yfilter.YFilter { return icmpstatstable.YFilter }

func (icmpstatstable *IPMIB_Icmpstatstable) SetFilter(yf yfilter.YFilter) { icmpstatstable.YFilter = yf }

func (icmpstatstable *IPMIB_Icmpstatstable) GetGoName(yname string) string {
    if yname == "icmpStatsEntry" { return "Icmpstatsentry" }
    return ""
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetSegmentPath() string {
    return "icmpStatsTable"
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "icmpStatsEntry" {
        for _, c := range icmpstatstable.Icmpstatsentry {
            if icmpstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Icmpstatstable_Icmpstatsentry{}
        icmpstatstable.Icmpstatsentry = append(icmpstatstable.Icmpstatsentry, child)
        return &icmpstatstable.Icmpstatsentry[len(icmpstatstable.Icmpstatsentry)-1]
    }
    return nil
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icmpstatstable.Icmpstatsentry {
        children[icmpstatstable.Icmpstatsentry[i].GetSegmentPath()] = &icmpstatstable.Icmpstatsentry[i]
    }
    return children
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (icmpstatstable *IPMIB_Icmpstatstable) GetYangName() string { return "icmpStatsTable" }

func (icmpstatstable *IPMIB_Icmpstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmpstatstable *IPMIB_Icmpstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmpstatstable *IPMIB_Icmpstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmpstatstable *IPMIB_Icmpstatstable) SetParent(parent types.Entity) { icmpstatstable.parent = parent }

func (icmpstatstable *IPMIB_Icmpstatstable) GetParent() types.Entity { return icmpstatstable.parent }

func (icmpstatstable *IPMIB_Icmpstatstable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Icmpstatstable_Icmpstatsentry
// A conceptual row in the icmpStatsTable.
type IPMIB_Icmpstatstable_Icmpstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP version of the statistics. The type is
    // IpVersion.
    Icmpstatsipversion interface{}

    // The total number of ICMP messages that the entity received. Note that this
    // counter includes all those counted by icmpStatsInErrors. The type is
    // interface{} with range: 0..4294967295.
    Icmpstatsinmsgs interface{}

    // The number of ICMP messages that the entity received but determined as
    // having ICMP-specific errors (bad ICMP checksums, bad length, etc.). The
    // type is interface{} with range: 0..4294967295.
    Icmpstatsinerrors interface{}

    // The total number of ICMP messages that the entity attempted to send.  Note
    // that this counter includes all those counted by icmpStatsOutErrors. The
    // type is interface{} with range: 0..4294967295.
    Icmpstatsoutmsgs interface{}

    // The number of ICMP messages that this entity did not send due to problems
    // discovered within ICMP, such as a lack of buffers.  This value should not
    // include errors discovered outside the ICMP layer, such as the inability of
    // IP to route the resultant datagram.  In some implementations, there may be
    // no types of error that contribute to this counter's value. The type is
    // interface{} with range: 0..4294967295.
    Icmpstatsouterrors interface{}
}

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetFilter() yfilter.YFilter { return icmpstatsentry.YFilter }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) SetFilter(yf yfilter.YFilter) { icmpstatsentry.YFilter = yf }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetGoName(yname string) string {
    if yname == "icmpStatsIPVersion" { return "Icmpstatsipversion" }
    if yname == "icmpStatsInMsgs" { return "Icmpstatsinmsgs" }
    if yname == "icmpStatsInErrors" { return "Icmpstatsinerrors" }
    if yname == "icmpStatsOutMsgs" { return "Icmpstatsoutmsgs" }
    if yname == "icmpStatsOutErrors" { return "Icmpstatsouterrors" }
    return ""
}

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetSegmentPath() string {
    return "icmpStatsEntry" + "[icmpStatsIPVersion='" + fmt.Sprintf("%v", icmpstatsentry.Icmpstatsipversion) + "']"
}

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icmpStatsIPVersion"] = icmpstatsentry.Icmpstatsipversion
    leafs["icmpStatsInMsgs"] = icmpstatsentry.Icmpstatsinmsgs
    leafs["icmpStatsInErrors"] = icmpstatsentry.Icmpstatsinerrors
    leafs["icmpStatsOutMsgs"] = icmpstatsentry.Icmpstatsoutmsgs
    leafs["icmpStatsOutErrors"] = icmpstatsentry.Icmpstatsouterrors
    return leafs
}

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetYangName() string { return "icmpStatsEntry" }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) SetParent(parent types.Entity) { icmpstatsentry.parent = parent }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetParent() types.Entity { return icmpstatsentry.parent }

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetParentYangName() string { return "icmpStatsTable" }

// IPMIB_Icmpmsgstatstable
// The table of system-wide per-version, per-message type ICMP
// counters.
type IPMIB_Icmpmsgstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the icmpMsgStatsTable.  The system should track each
    // ICMP type value, even if that ICMP type is not supported by the system. 
    // However, a given row need not be instantiated unless a message of that type
    // has been processed, i.e., the row for icmpMsgStatsType=X MAY be
    // instantiated before but MUST be instantiated after the first message with
    // Type=X is received or transmitted.  After receiving or transmitting any
    // succeeding messages with Type=X, the relevant counter must be incremented.
    // The type is slice of IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry.
    Icmpmsgstatsentry []IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry
}

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetFilter() yfilter.YFilter { return icmpmsgstatstable.YFilter }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) SetFilter(yf yfilter.YFilter) { icmpmsgstatstable.YFilter = yf }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetGoName(yname string) string {
    if yname == "icmpMsgStatsEntry" { return "Icmpmsgstatsentry" }
    return ""
}

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetSegmentPath() string {
    return "icmpMsgStatsTable"
}

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "icmpMsgStatsEntry" {
        for _, c := range icmpmsgstatstable.Icmpmsgstatsentry {
            if icmpmsgstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry{}
        icmpmsgstatstable.Icmpmsgstatsentry = append(icmpmsgstatstable.Icmpmsgstatsentry, child)
        return &icmpmsgstatstable.Icmpmsgstatsentry[len(icmpmsgstatstable.Icmpmsgstatsentry)-1]
    }
    return nil
}

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icmpmsgstatstable.Icmpmsgstatsentry {
        children[icmpmsgstatstable.Icmpmsgstatsentry[i].GetSegmentPath()] = &icmpmsgstatstable.Icmpmsgstatsentry[i]
    }
    return children
}

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetYangName() string { return "icmpMsgStatsTable" }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) SetParent(parent types.Entity) { icmpmsgstatstable.parent = parent }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetParent() types.Entity { return icmpmsgstatstable.parent }

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetParentYangName() string { return "IP-MIB" }

// IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry
// A conceptual row in the icmpMsgStatsTable.
// 
// The system should track each ICMP type value, even if that
// ICMP type is not supported by the system.  However, a
// given row need not be instantiated unless a message of that
// type has been processed, i.e., the row for
// icmpMsgStatsType=X MAY be instantiated before but MUST be
// instantiated after the first message with Type=X is
// received or transmitted.  After receiving or transmitting
// any succeeding messages with Type=X, the relevant counter
// must be incremented.
type IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP version of the statistics. The type is
    // IpVersion.
    Icmpmsgstatsipversion interface{}

    // This attribute is a key. The ICMP type field of the message type being
    // counted by this row.  Note that ICMP message types are scoped by the
    // address type in use. The type is interface{} with range: 0..255.
    Icmpmsgstatstype interface{}

    // The number of input packets for this AF and type. The type is interface{}
    // with range: 0..4294967295.
    Icmpmsgstatsinpkts interface{}

    // The number of output packets for this AF and type. The type is interface{}
    // with range: 0..4294967295.
    Icmpmsgstatsoutpkts interface{}
}

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetFilter() yfilter.YFilter { return icmpmsgstatsentry.YFilter }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) SetFilter(yf yfilter.YFilter) { icmpmsgstatsentry.YFilter = yf }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetGoName(yname string) string {
    if yname == "icmpMsgStatsIPVersion" { return "Icmpmsgstatsipversion" }
    if yname == "icmpMsgStatsType" { return "Icmpmsgstatstype" }
    if yname == "icmpMsgStatsInPkts" { return "Icmpmsgstatsinpkts" }
    if yname == "icmpMsgStatsOutPkts" { return "Icmpmsgstatsoutpkts" }
    return ""
}

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetSegmentPath() string {
    return "icmpMsgStatsEntry" + "[icmpMsgStatsIPVersion='" + fmt.Sprintf("%v", icmpmsgstatsentry.Icmpmsgstatsipversion) + "']" + "[icmpMsgStatsType='" + fmt.Sprintf("%v", icmpmsgstatsentry.Icmpmsgstatstype) + "']"
}

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icmpMsgStatsIPVersion"] = icmpmsgstatsentry.Icmpmsgstatsipversion
    leafs["icmpMsgStatsType"] = icmpmsgstatsentry.Icmpmsgstatstype
    leafs["icmpMsgStatsInPkts"] = icmpmsgstatsentry.Icmpmsgstatsinpkts
    leafs["icmpMsgStatsOutPkts"] = icmpmsgstatsentry.Icmpmsgstatsoutpkts
    return leafs
}

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetYangName() string { return "icmpMsgStatsEntry" }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) SetParent(parent types.Entity) { icmpmsgstatsentry.parent = parent }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetParent() types.Entity { return icmpmsgstatsentry.parent }

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetParentYangName() string { return "icmpMsgStatsTable" }

