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

// IpAddressPrefixOriginTC represents prefix was found.
type IpAddressPrefixOriginTC string

const (
    IpAddressPrefixOriginTC_other IpAddressPrefixOriginTC = "other"

    IpAddressPrefixOriginTC_manual IpAddressPrefixOriginTC = "manual"

    IpAddressPrefixOriginTC_wellknown IpAddressPrefixOriginTC = "wellknown"

    IpAddressPrefixOriginTC_dhcp IpAddressPrefixOriginTC = "dhcp"

    IpAddressPrefixOriginTC_routeradv IpAddressPrefixOriginTC = "routeradv"
)

// IPMIB
type IPMIB struct {
    EntityData types.CommonEntityData
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

func (iPMIB *IPMIB) GetEntityData() *types.CommonEntityData {
    iPMIB.EntityData.YFilter = iPMIB.YFilter
    iPMIB.EntityData.YangName = "IP-MIB"
    iPMIB.EntityData.BundleName = "cisco_ios_xe"
    iPMIB.EntityData.ParentYangName = "IP-MIB"
    iPMIB.EntityData.SegmentPath = "IP-MIB:IP-MIB"
    iPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iPMIB.EntityData.Children = make(map[string]types.YChild)
    iPMIB.EntityData.Children["ip"] = types.YChild{"Ip", &iPMIB.Ip}
    iPMIB.EntityData.Children["ipTrafficStats"] = types.YChild{"Iptrafficstats", &iPMIB.Iptrafficstats}
    iPMIB.EntityData.Children["icmp"] = types.YChild{"Icmp", &iPMIB.Icmp}
    iPMIB.EntityData.Children["ipAddrTable"] = types.YChild{"Ipaddrtable", &iPMIB.Ipaddrtable}
    iPMIB.EntityData.Children["ipNetToMediaTable"] = types.YChild{"Ipnettomediatable", &iPMIB.Ipnettomediatable}
    iPMIB.EntityData.Children["ipv4InterfaceTable"] = types.YChild{"Ipv4Interfacetable", &iPMIB.Ipv4Interfacetable}
    iPMIB.EntityData.Children["ipv6InterfaceTable"] = types.YChild{"Ipv6Interfacetable", &iPMIB.Ipv6Interfacetable}
    iPMIB.EntityData.Children["ipSystemStatsTable"] = types.YChild{"Ipsystemstatstable", &iPMIB.Ipsystemstatstable}
    iPMIB.EntityData.Children["ipIfStatsTable"] = types.YChild{"Ipifstatstable", &iPMIB.Ipifstatstable}
    iPMIB.EntityData.Children["ipAddressPrefixTable"] = types.YChild{"Ipaddressprefixtable", &iPMIB.Ipaddressprefixtable}
    iPMIB.EntityData.Children["ipAddressTable"] = types.YChild{"Ipaddresstable", &iPMIB.Ipaddresstable}
    iPMIB.EntityData.Children["ipNetToPhysicalTable"] = types.YChild{"Ipnettophysicaltable", &iPMIB.Ipnettophysicaltable}
    iPMIB.EntityData.Children["ipv6ScopeZoneIndexTable"] = types.YChild{"Ipv6Scopezoneindextable", &iPMIB.Ipv6Scopezoneindextable}
    iPMIB.EntityData.Children["ipDefaultRouterTable"] = types.YChild{"Ipdefaultroutertable", &iPMIB.Ipdefaultroutertable}
    iPMIB.EntityData.Children["ipv6RouterAdvertTable"] = types.YChild{"Ipv6Routeradverttable", &iPMIB.Ipv6Routeradverttable}
    iPMIB.EntityData.Children["icmpStatsTable"] = types.YChild{"Icmpstatstable", &iPMIB.Icmpstatstable}
    iPMIB.EntityData.Children["icmpMsgStatsTable"] = types.YChild{"Icmpmsgstatstable", &iPMIB.Icmpmsgstatstable}
    iPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iPMIB.EntityData)
}

// IPMIB_Ip
type IPMIB_Ip struct {
    EntityData types.CommonEntityData
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

func (ip *IPMIB_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xe"
    ip.EntityData.ParentYangName = "IP-MIB"
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
    ip.EntityData.Leafs["ipv6IpForwarding"] = types.YLeaf{"Ipv6Ipforwarding", ip.Ipv6Ipforwarding}
    ip.EntityData.Leafs["ipv6IpDefaultHopLimit"] = types.YLeaf{"Ipv6Ipdefaulthoplimit", ip.Ipv6Ipdefaulthoplimit}
    ip.EntityData.Leafs["ipv4InterfaceTableLastChange"] = types.YLeaf{"Ipv4Interfacetablelastchange", ip.Ipv4Interfacetablelastchange}
    ip.EntityData.Leafs["ipv6InterfaceTableLastChange"] = types.YLeaf{"Ipv6Interfacetablelastchange", ip.Ipv6Interfacetablelastchange}
    ip.EntityData.Leafs["ipAddressSpinLock"] = types.YLeaf{"Ipaddressspinlock", ip.Ipaddressspinlock}
    ip.EntityData.Leafs["ipv6RouterAdvertSpinLock"] = types.YLeaf{"Ipv6Routeradvertspinlock", ip.Ipv6Routeradvertspinlock}
    return &(ip.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipIfStatsTable was added or deleted.  If new objects are added to the
    // ipIfStatsTable that require the ipIfStatsTableLastChange to be updated when
    // they are modified, they must specify that requirement in their description
    // clause. The type is interface{} with range: 0..4294967295.
    Ipifstatstablelastchange interface{}
}

func (iptrafficstats *IPMIB_Iptrafficstats) GetEntityData() *types.CommonEntityData {
    iptrafficstats.EntityData.YFilter = iptrafficstats.YFilter
    iptrafficstats.EntityData.YangName = "ipTrafficStats"
    iptrafficstats.EntityData.BundleName = "cisco_ios_xe"
    iptrafficstats.EntityData.ParentYangName = "IP-MIB"
    iptrafficstats.EntityData.SegmentPath = "ipTrafficStats"
    iptrafficstats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iptrafficstats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iptrafficstats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iptrafficstats.EntityData.Children = make(map[string]types.YChild)
    iptrafficstats.EntityData.Leafs = make(map[string]types.YLeaf)
    iptrafficstats.EntityData.Leafs["ipIfStatsTableLastChange"] = types.YLeaf{"Ipifstatstablelastchange", iptrafficstats.Ipifstatstablelastchange}
    return &(iptrafficstats.EntityData)
}

// IPMIB_Icmp
type IPMIB_Icmp struct {
    EntityData types.CommonEntityData
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

func (icmp *IPMIB_Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xe"
    icmp.EntityData.ParentYangName = "IP-MIB"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The addressing information for one of this entity's IPv4 addresses. The
    // type is slice of IPMIB_Ipaddrtable_Ipaddrentry.
    Ipaddrentry []IPMIB_Ipaddrtable_Ipaddrentry
}

func (ipaddrtable *IPMIB_Ipaddrtable) GetEntityData() *types.CommonEntityData {
    ipaddrtable.EntityData.YFilter = ipaddrtable.YFilter
    ipaddrtable.EntityData.YangName = "ipAddrTable"
    ipaddrtable.EntityData.BundleName = "cisco_ios_xe"
    ipaddrtable.EntityData.ParentYangName = "IP-MIB"
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

// IPMIB_Ipaddrtable_Ipaddrentry
// The addressing information for one of this entity's IPv4
// addresses.
type IPMIB_Ipaddrtable_Ipaddrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address to which this entry's addressing
    // information pertains. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipadentaddr interface{}

    // The index value which uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of the IF-MIB's
    // ifIndex. The type is interface{} with range: 1..2147483647.
    Ipadentifindex interface{}

    // The subnet mask associated with the IPv4 address of this entry.  The value
    // of the mask is an IPv4 address with all the network bits set to 1 and all
    // the hosts bits set to 0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipaddrentry *IPMIB_Ipaddrtable_Ipaddrentry) GetEntityData() *types.CommonEntityData {
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

// IPMIB_Ipnettomediatable
// The IPv4 Address Translation table used for mapping from
// IPv4 addresses to physical addresses.
// 
// This table has been deprecated, as a new IP version-neutral
// table has been added.  It is loosely replaced by the
// ipNetToPhysicalTable.
type IPMIB_Ipnettomediatable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one IpAddress to `physical' address equivalence. The
    // type is slice of IPMIB_Ipnettomediatable_Ipnettomediaentry.
    Ipnettomediaentry []IPMIB_Ipnettomediatable_Ipnettomediaentry
}

func (ipnettomediatable *IPMIB_Ipnettomediatable) GetEntityData() *types.CommonEntityData {
    ipnettomediatable.EntityData.YFilter = ipnettomediatable.YFilter
    ipnettomediatable.EntityData.YangName = "ipNetToMediaTable"
    ipnettomediatable.EntityData.BundleName = "cisco_ios_xe"
    ipnettomediatable.EntityData.ParentYangName = "IP-MIB"
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

// IPMIB_Ipnettomediatable_Ipnettomediaentry
// Each entry contains one IpAddress to `physical' address
// equivalence.
type IPMIB_Ipnettomediatable_Ipnettomediaentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipnettomediaentry *IPMIB_Ipnettomediatable_Ipnettomediaentry) GetEntityData() *types.CommonEntityData {
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing IPv4-specific information for a specific interface. The
    // type is slice of IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry.
    Ipv4Interfaceentry []IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry
}

func (ipv4Interfacetable *IPMIB_Ipv4Interfacetable) GetEntityData() *types.CommonEntityData {
    ipv4Interfacetable.EntityData.YFilter = ipv4Interfacetable.YFilter
    ipv4Interfacetable.EntityData.YangName = "ipv4InterfaceTable"
    ipv4Interfacetable.EntityData.BundleName = "cisco_ios_xe"
    ipv4Interfacetable.EntityData.ParentYangName = "IP-MIB"
    ipv4Interfacetable.EntityData.SegmentPath = "ipv4InterfaceTable"
    ipv4Interfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4Interfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4Interfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4Interfacetable.EntityData.Children = make(map[string]types.YChild)
    ipv4Interfacetable.EntityData.Children["ipv4InterfaceEntry"] = types.YChild{"Ipv4Interfaceentry", nil}
    for i := range ipv4Interfacetable.Ipv4Interfaceentry {
        ipv4Interfacetable.EntityData.Children[types.GetSegmentPath(&ipv4Interfacetable.Ipv4Interfaceentry[i])] = types.YChild{"Ipv4Interfaceentry", &ipv4Interfacetable.Ipv4Interfaceentry[i]}
    }
    ipv4Interfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Interfacetable.EntityData)
}

// IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry
// An entry containing IPv4-specific information for a specific
// interface.
type IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry struct {
    EntityData types.CommonEntityData
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

func (ipv4Interfaceentry *IPMIB_Ipv4Interfacetable_Ipv4Interfaceentry) GetEntityData() *types.CommonEntityData {
    ipv4Interfaceentry.EntityData.YFilter = ipv4Interfaceentry.YFilter
    ipv4Interfaceentry.EntityData.YangName = "ipv4InterfaceEntry"
    ipv4Interfaceentry.EntityData.BundleName = "cisco_ios_xe"
    ipv4Interfaceentry.EntityData.ParentYangName = "ipv4InterfaceTable"
    ipv4Interfaceentry.EntityData.SegmentPath = "ipv4InterfaceEntry" + "[ipv4InterfaceIfIndex='" + fmt.Sprintf("%v", ipv4Interfaceentry.Ipv4Interfaceifindex) + "']"
    ipv4Interfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4Interfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4Interfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4Interfaceentry.EntityData.Children = make(map[string]types.YChild)
    ipv4Interfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Interfaceentry.EntityData.Leafs["ipv4InterfaceIfIndex"] = types.YLeaf{"Ipv4Interfaceifindex", ipv4Interfaceentry.Ipv4Interfaceifindex}
    ipv4Interfaceentry.EntityData.Leafs["ipv4InterfaceReasmMaxSize"] = types.YLeaf{"Ipv4Interfacereasmmaxsize", ipv4Interfaceentry.Ipv4Interfacereasmmaxsize}
    ipv4Interfaceentry.EntityData.Leafs["ipv4InterfaceEnableStatus"] = types.YLeaf{"Ipv4Interfaceenablestatus", ipv4Interfaceentry.Ipv4Interfaceenablestatus}
    ipv4Interfaceentry.EntityData.Leafs["ipv4InterfaceRetransmitTime"] = types.YLeaf{"Ipv4Interfaceretransmittime", ipv4Interfaceentry.Ipv4Interfaceretransmittime}
    return &(ipv4Interfaceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing IPv6-specific information for a given interface. The
    // type is slice of IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry.
    Ipv6Interfaceentry []IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry
}

func (ipv6Interfacetable *IPMIB_Ipv6Interfacetable) GetEntityData() *types.CommonEntityData {
    ipv6Interfacetable.EntityData.YFilter = ipv6Interfacetable.YFilter
    ipv6Interfacetable.EntityData.YangName = "ipv6InterfaceTable"
    ipv6Interfacetable.EntityData.BundleName = "cisco_ios_xe"
    ipv6Interfacetable.EntityData.ParentYangName = "IP-MIB"
    ipv6Interfacetable.EntityData.SegmentPath = "ipv6InterfaceTable"
    ipv6Interfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Interfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Interfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Interfacetable.EntityData.Children = make(map[string]types.YChild)
    ipv6Interfacetable.EntityData.Children["ipv6InterfaceEntry"] = types.YChild{"Ipv6Interfaceentry", nil}
    for i := range ipv6Interfacetable.Ipv6Interfaceentry {
        ipv6Interfacetable.EntityData.Children[types.GetSegmentPath(&ipv6Interfacetable.Ipv6Interfaceentry[i])] = types.YChild{"Ipv6Interfaceentry", &ipv6Interfacetable.Ipv6Interfaceentry[i]}
    }
    ipv6Interfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Interfacetable.EntityData)
}

// IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry
// An entry containing IPv6-specific information for a given
// interface.
type IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry struct {
    EntityData types.CommonEntityData
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

func (ipv6Interfaceentry *IPMIB_Ipv6Interfacetable_Ipv6Interfaceentry) GetEntityData() *types.CommonEntityData {
    ipv6Interfaceentry.EntityData.YFilter = ipv6Interfaceentry.YFilter
    ipv6Interfaceentry.EntityData.YangName = "ipv6InterfaceEntry"
    ipv6Interfaceentry.EntityData.BundleName = "cisco_ios_xe"
    ipv6Interfaceentry.EntityData.ParentYangName = "ipv6InterfaceTable"
    ipv6Interfaceentry.EntityData.SegmentPath = "ipv6InterfaceEntry" + "[ipv6InterfaceIfIndex='" + fmt.Sprintf("%v", ipv6Interfaceentry.Ipv6Interfaceifindex) + "']"
    ipv6Interfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Interfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Interfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Interfaceentry.EntityData.Children = make(map[string]types.YChild)
    ipv6Interfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceIfIndex"] = types.YLeaf{"Ipv6Interfaceifindex", ipv6Interfaceentry.Ipv6Interfaceifindex}
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceReasmMaxSize"] = types.YLeaf{"Ipv6Interfacereasmmaxsize", ipv6Interfaceentry.Ipv6Interfacereasmmaxsize}
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceIdentifier"] = types.YLeaf{"Ipv6Interfaceidentifier", ipv6Interfaceentry.Ipv6Interfaceidentifier}
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceEnableStatus"] = types.YLeaf{"Ipv6Interfaceenablestatus", ipv6Interfaceentry.Ipv6Interfaceenablestatus}
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceReachableTime"] = types.YLeaf{"Ipv6Interfacereachabletime", ipv6Interfaceentry.Ipv6Interfacereachabletime}
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceRetransmitTime"] = types.YLeaf{"Ipv6Interfaceretransmittime", ipv6Interfaceentry.Ipv6Interfaceretransmittime}
    ipv6Interfaceentry.EntityData.Leafs["ipv6InterfaceForwarding"] = types.YLeaf{"Ipv6Interfaceforwarding", ipv6Interfaceentry.Ipv6Interfaceforwarding}
    return &(ipv6Interfaceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A statistics entry containing system-wide objects for a particular IP
    // version. The type is slice of IPMIB_Ipsystemstatstable_Ipsystemstatsentry.
    Ipsystemstatsentry []IPMIB_Ipsystemstatstable_Ipsystemstatsentry
}

func (ipsystemstatstable *IPMIB_Ipsystemstatstable) GetEntityData() *types.CommonEntityData {
    ipsystemstatstable.EntityData.YFilter = ipsystemstatstable.YFilter
    ipsystemstatstable.EntityData.YangName = "ipSystemStatsTable"
    ipsystemstatstable.EntityData.BundleName = "cisco_ios_xe"
    ipsystemstatstable.EntityData.ParentYangName = "IP-MIB"
    ipsystemstatstable.EntityData.SegmentPath = "ipSystemStatsTable"
    ipsystemstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsystemstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsystemstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsystemstatstable.EntityData.Children = make(map[string]types.YChild)
    ipsystemstatstable.EntityData.Children["ipSystemStatsEntry"] = types.YChild{"Ipsystemstatsentry", nil}
    for i := range ipsystemstatstable.Ipsystemstatsentry {
        ipsystemstatstable.EntityData.Children[types.GetSegmentPath(&ipsystemstatstable.Ipsystemstatsentry[i])] = types.YChild{"Ipsystemstatsentry", &ipsystemstatstable.Ipsystemstatsentry[i]}
    }
    ipsystemstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipsystemstatstable.EntityData)
}

// IPMIB_Ipsystemstatstable_Ipsystemstatsentry
// A statistics entry containing system-wide objects for a
// particular IP version.
type IPMIB_Ipsystemstatstable_Ipsystemstatsentry struct {
    EntityData types.CommonEntityData
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

func (ipsystemstatsentry *IPMIB_Ipsystemstatstable_Ipsystemstatsentry) GetEntityData() *types.CommonEntityData {
    ipsystemstatsentry.EntityData.YFilter = ipsystemstatsentry.YFilter
    ipsystemstatsentry.EntityData.YangName = "ipSystemStatsEntry"
    ipsystemstatsentry.EntityData.BundleName = "cisco_ios_xe"
    ipsystemstatsentry.EntityData.ParentYangName = "ipSystemStatsTable"
    ipsystemstatsentry.EntityData.SegmentPath = "ipSystemStatsEntry" + "[ipSystemStatsIPVersion='" + fmt.Sprintf("%v", ipsystemstatsentry.Ipsystemstatsipversion) + "']"
    ipsystemstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsystemstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsystemstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsystemstatsentry.EntityData.Children = make(map[string]types.YChild)
    ipsystemstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsIPVersion"] = types.YLeaf{"Ipsystemstatsipversion", ipsystemstatsentry.Ipsystemstatsipversion}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInReceives"] = types.YLeaf{"Ipsystemstatsinreceives", ipsystemstatsentry.Ipsystemstatsinreceives}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInReceives"] = types.YLeaf{"Ipsystemstatshcinreceives", ipsystemstatsentry.Ipsystemstatshcinreceives}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInOctets"] = types.YLeaf{"Ipsystemstatsinoctets", ipsystemstatsentry.Ipsystemstatsinoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInOctets"] = types.YLeaf{"Ipsystemstatshcinoctets", ipsystemstatsentry.Ipsystemstatshcinoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInHdrErrors"] = types.YLeaf{"Ipsystemstatsinhdrerrors", ipsystemstatsentry.Ipsystemstatsinhdrerrors}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInNoRoutes"] = types.YLeaf{"Ipsystemstatsinnoroutes", ipsystemstatsentry.Ipsystemstatsinnoroutes}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInAddrErrors"] = types.YLeaf{"Ipsystemstatsinaddrerrors", ipsystemstatsentry.Ipsystemstatsinaddrerrors}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInUnknownProtos"] = types.YLeaf{"Ipsystemstatsinunknownprotos", ipsystemstatsentry.Ipsystemstatsinunknownprotos}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInTruncatedPkts"] = types.YLeaf{"Ipsystemstatsintruncatedpkts", ipsystemstatsentry.Ipsystemstatsintruncatedpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInForwDatagrams"] = types.YLeaf{"Ipsystemstatsinforwdatagrams", ipsystemstatsentry.Ipsystemstatsinforwdatagrams}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInForwDatagrams"] = types.YLeaf{"Ipsystemstatshcinforwdatagrams", ipsystemstatsentry.Ipsystemstatshcinforwdatagrams}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsReasmReqds"] = types.YLeaf{"Ipsystemstatsreasmreqds", ipsystemstatsentry.Ipsystemstatsreasmreqds}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsReasmOKs"] = types.YLeaf{"Ipsystemstatsreasmoks", ipsystemstatsentry.Ipsystemstatsreasmoks}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsReasmFails"] = types.YLeaf{"Ipsystemstatsreasmfails", ipsystemstatsentry.Ipsystemstatsreasmfails}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInDiscards"] = types.YLeaf{"Ipsystemstatsindiscards", ipsystemstatsentry.Ipsystemstatsindiscards}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInDelivers"] = types.YLeaf{"Ipsystemstatsindelivers", ipsystemstatsentry.Ipsystemstatsindelivers}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInDelivers"] = types.YLeaf{"Ipsystemstatshcindelivers", ipsystemstatsentry.Ipsystemstatshcindelivers}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutRequests"] = types.YLeaf{"Ipsystemstatsoutrequests", ipsystemstatsentry.Ipsystemstatsoutrequests}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutRequests"] = types.YLeaf{"Ipsystemstatshcoutrequests", ipsystemstatsentry.Ipsystemstatshcoutrequests}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutNoRoutes"] = types.YLeaf{"Ipsystemstatsoutnoroutes", ipsystemstatsentry.Ipsystemstatsoutnoroutes}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutForwDatagrams"] = types.YLeaf{"Ipsystemstatsoutforwdatagrams", ipsystemstatsentry.Ipsystemstatsoutforwdatagrams}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutForwDatagrams"] = types.YLeaf{"Ipsystemstatshcoutforwdatagrams", ipsystemstatsentry.Ipsystemstatshcoutforwdatagrams}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutDiscards"] = types.YLeaf{"Ipsystemstatsoutdiscards", ipsystemstatsentry.Ipsystemstatsoutdiscards}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutFragReqds"] = types.YLeaf{"Ipsystemstatsoutfragreqds", ipsystemstatsentry.Ipsystemstatsoutfragreqds}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutFragOKs"] = types.YLeaf{"Ipsystemstatsoutfragoks", ipsystemstatsentry.Ipsystemstatsoutfragoks}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutFragFails"] = types.YLeaf{"Ipsystemstatsoutfragfails", ipsystemstatsentry.Ipsystemstatsoutfragfails}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutFragCreates"] = types.YLeaf{"Ipsystemstatsoutfragcreates", ipsystemstatsentry.Ipsystemstatsoutfragcreates}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutTransmits"] = types.YLeaf{"Ipsystemstatsouttransmits", ipsystemstatsentry.Ipsystemstatsouttransmits}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutTransmits"] = types.YLeaf{"Ipsystemstatshcouttransmits", ipsystemstatsentry.Ipsystemstatshcouttransmits}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutOctets"] = types.YLeaf{"Ipsystemstatsoutoctets", ipsystemstatsentry.Ipsystemstatsoutoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutOctets"] = types.YLeaf{"Ipsystemstatshcoutoctets", ipsystemstatsentry.Ipsystemstatshcoutoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInMcastPkts"] = types.YLeaf{"Ipsystemstatsinmcastpkts", ipsystemstatsentry.Ipsystemstatsinmcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInMcastPkts"] = types.YLeaf{"Ipsystemstatshcinmcastpkts", ipsystemstatsentry.Ipsystemstatshcinmcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInMcastOctets"] = types.YLeaf{"Ipsystemstatsinmcastoctets", ipsystemstatsentry.Ipsystemstatsinmcastoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInMcastOctets"] = types.YLeaf{"Ipsystemstatshcinmcastoctets", ipsystemstatsentry.Ipsystemstatshcinmcastoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutMcastPkts"] = types.YLeaf{"Ipsystemstatsoutmcastpkts", ipsystemstatsentry.Ipsystemstatsoutmcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutMcastPkts"] = types.YLeaf{"Ipsystemstatshcoutmcastpkts", ipsystemstatsentry.Ipsystemstatshcoutmcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutMcastOctets"] = types.YLeaf{"Ipsystemstatsoutmcastoctets", ipsystemstatsentry.Ipsystemstatsoutmcastoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutMcastOctets"] = types.YLeaf{"Ipsystemstatshcoutmcastoctets", ipsystemstatsentry.Ipsystemstatshcoutmcastoctets}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsInBcastPkts"] = types.YLeaf{"Ipsystemstatsinbcastpkts", ipsystemstatsentry.Ipsystemstatsinbcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCInBcastPkts"] = types.YLeaf{"Ipsystemstatshcinbcastpkts", ipsystemstatsentry.Ipsystemstatshcinbcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsOutBcastPkts"] = types.YLeaf{"Ipsystemstatsoutbcastpkts", ipsystemstatsentry.Ipsystemstatsoutbcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsHCOutBcastPkts"] = types.YLeaf{"Ipsystemstatshcoutbcastpkts", ipsystemstatsentry.Ipsystemstatshcoutbcastpkts}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsDiscontinuityTime"] = types.YLeaf{"Ipsystemstatsdiscontinuitytime", ipsystemstatsentry.Ipsystemstatsdiscontinuitytime}
    ipsystemstatsentry.EntityData.Leafs["ipSystemStatsRefreshRate"] = types.YLeaf{"Ipsystemstatsrefreshrate", ipsystemstatsentry.Ipsystemstatsrefreshrate}
    return &(ipsystemstatsentry.EntityData)
}

// IPMIB_Ipifstatstable
// The table containing per-interface traffic statistics.  This
// table and the ipSystemStatsTable contain similar objects
// whose difference is in their granularity.  Where this table
// contains per-interface statistics, the ipSystemStatsTable
// contains the same statistics, but counted on a system wide
// basis.
type IPMIB_Ipifstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An interface statistics entry containing objects for a particular interface
    // and version of IP. The type is slice of
    // IPMIB_Ipifstatstable_Ipifstatsentry.
    Ipifstatsentry []IPMIB_Ipifstatstable_Ipifstatsentry
}

func (ipifstatstable *IPMIB_Ipifstatstable) GetEntityData() *types.CommonEntityData {
    ipifstatstable.EntityData.YFilter = ipifstatstable.YFilter
    ipifstatstable.EntityData.YangName = "ipIfStatsTable"
    ipifstatstable.EntityData.BundleName = "cisco_ios_xe"
    ipifstatstable.EntityData.ParentYangName = "IP-MIB"
    ipifstatstable.EntityData.SegmentPath = "ipIfStatsTable"
    ipifstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipifstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipifstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipifstatstable.EntityData.Children = make(map[string]types.YChild)
    ipifstatstable.EntityData.Children["ipIfStatsEntry"] = types.YChild{"Ipifstatsentry", nil}
    for i := range ipifstatstable.Ipifstatsentry {
        ipifstatstable.EntityData.Children[types.GetSegmentPath(&ipifstatstable.Ipifstatsentry[i])] = types.YChild{"Ipifstatsentry", &ipifstatstable.Ipifstatsentry[i]}
    }
    ipifstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipifstatstable.EntityData)
}

// IPMIB_Ipifstatstable_Ipifstatsentry
// An interface statistics entry containing objects for a
// particular interface and version of IP.
type IPMIB_Ipifstatstable_Ipifstatsentry struct {
    EntityData types.CommonEntityData
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

func (ipifstatsentry *IPMIB_Ipifstatstable_Ipifstatsentry) GetEntityData() *types.CommonEntityData {
    ipifstatsentry.EntityData.YFilter = ipifstatsentry.YFilter
    ipifstatsentry.EntityData.YangName = "ipIfStatsEntry"
    ipifstatsentry.EntityData.BundleName = "cisco_ios_xe"
    ipifstatsentry.EntityData.ParentYangName = "ipIfStatsTable"
    ipifstatsentry.EntityData.SegmentPath = "ipIfStatsEntry" + "[ipIfStatsIPVersion='" + fmt.Sprintf("%v", ipifstatsentry.Ipifstatsipversion) + "']" + "[ipIfStatsIfIndex='" + fmt.Sprintf("%v", ipifstatsentry.Ipifstatsifindex) + "']"
    ipifstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipifstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipifstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipifstatsentry.EntityData.Children = make(map[string]types.YChild)
    ipifstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipifstatsentry.EntityData.Leafs["ipIfStatsIPVersion"] = types.YLeaf{"Ipifstatsipversion", ipifstatsentry.Ipifstatsipversion}
    ipifstatsentry.EntityData.Leafs["ipIfStatsIfIndex"] = types.YLeaf{"Ipifstatsifindex", ipifstatsentry.Ipifstatsifindex}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInReceives"] = types.YLeaf{"Ipifstatsinreceives", ipifstatsentry.Ipifstatsinreceives}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInReceives"] = types.YLeaf{"Ipifstatshcinreceives", ipifstatsentry.Ipifstatshcinreceives}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInOctets"] = types.YLeaf{"Ipifstatsinoctets", ipifstatsentry.Ipifstatsinoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInOctets"] = types.YLeaf{"Ipifstatshcinoctets", ipifstatsentry.Ipifstatshcinoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInHdrErrors"] = types.YLeaf{"Ipifstatsinhdrerrors", ipifstatsentry.Ipifstatsinhdrerrors}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInNoRoutes"] = types.YLeaf{"Ipifstatsinnoroutes", ipifstatsentry.Ipifstatsinnoroutes}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInAddrErrors"] = types.YLeaf{"Ipifstatsinaddrerrors", ipifstatsentry.Ipifstatsinaddrerrors}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInUnknownProtos"] = types.YLeaf{"Ipifstatsinunknownprotos", ipifstatsentry.Ipifstatsinunknownprotos}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInTruncatedPkts"] = types.YLeaf{"Ipifstatsintruncatedpkts", ipifstatsentry.Ipifstatsintruncatedpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInForwDatagrams"] = types.YLeaf{"Ipifstatsinforwdatagrams", ipifstatsentry.Ipifstatsinforwdatagrams}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInForwDatagrams"] = types.YLeaf{"Ipifstatshcinforwdatagrams", ipifstatsentry.Ipifstatshcinforwdatagrams}
    ipifstatsentry.EntityData.Leafs["ipIfStatsReasmReqds"] = types.YLeaf{"Ipifstatsreasmreqds", ipifstatsentry.Ipifstatsreasmreqds}
    ipifstatsentry.EntityData.Leafs["ipIfStatsReasmOKs"] = types.YLeaf{"Ipifstatsreasmoks", ipifstatsentry.Ipifstatsreasmoks}
    ipifstatsentry.EntityData.Leafs["ipIfStatsReasmFails"] = types.YLeaf{"Ipifstatsreasmfails", ipifstatsentry.Ipifstatsreasmfails}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInDiscards"] = types.YLeaf{"Ipifstatsindiscards", ipifstatsentry.Ipifstatsindiscards}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInDelivers"] = types.YLeaf{"Ipifstatsindelivers", ipifstatsentry.Ipifstatsindelivers}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInDelivers"] = types.YLeaf{"Ipifstatshcindelivers", ipifstatsentry.Ipifstatshcindelivers}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutRequests"] = types.YLeaf{"Ipifstatsoutrequests", ipifstatsentry.Ipifstatsoutrequests}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutRequests"] = types.YLeaf{"Ipifstatshcoutrequests", ipifstatsentry.Ipifstatshcoutrequests}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutForwDatagrams"] = types.YLeaf{"Ipifstatsoutforwdatagrams", ipifstatsentry.Ipifstatsoutforwdatagrams}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutForwDatagrams"] = types.YLeaf{"Ipifstatshcoutforwdatagrams", ipifstatsentry.Ipifstatshcoutforwdatagrams}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutDiscards"] = types.YLeaf{"Ipifstatsoutdiscards", ipifstatsentry.Ipifstatsoutdiscards}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutFragReqds"] = types.YLeaf{"Ipifstatsoutfragreqds", ipifstatsentry.Ipifstatsoutfragreqds}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutFragOKs"] = types.YLeaf{"Ipifstatsoutfragoks", ipifstatsentry.Ipifstatsoutfragoks}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutFragFails"] = types.YLeaf{"Ipifstatsoutfragfails", ipifstatsentry.Ipifstatsoutfragfails}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutFragCreates"] = types.YLeaf{"Ipifstatsoutfragcreates", ipifstatsentry.Ipifstatsoutfragcreates}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutTransmits"] = types.YLeaf{"Ipifstatsouttransmits", ipifstatsentry.Ipifstatsouttransmits}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutTransmits"] = types.YLeaf{"Ipifstatshcouttransmits", ipifstatsentry.Ipifstatshcouttransmits}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutOctets"] = types.YLeaf{"Ipifstatsoutoctets", ipifstatsentry.Ipifstatsoutoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutOctets"] = types.YLeaf{"Ipifstatshcoutoctets", ipifstatsentry.Ipifstatshcoutoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInMcastPkts"] = types.YLeaf{"Ipifstatsinmcastpkts", ipifstatsentry.Ipifstatsinmcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInMcastPkts"] = types.YLeaf{"Ipifstatshcinmcastpkts", ipifstatsentry.Ipifstatshcinmcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInMcastOctets"] = types.YLeaf{"Ipifstatsinmcastoctets", ipifstatsentry.Ipifstatsinmcastoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInMcastOctets"] = types.YLeaf{"Ipifstatshcinmcastoctets", ipifstatsentry.Ipifstatshcinmcastoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutMcastPkts"] = types.YLeaf{"Ipifstatsoutmcastpkts", ipifstatsentry.Ipifstatsoutmcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutMcastPkts"] = types.YLeaf{"Ipifstatshcoutmcastpkts", ipifstatsentry.Ipifstatshcoutmcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutMcastOctets"] = types.YLeaf{"Ipifstatsoutmcastoctets", ipifstatsentry.Ipifstatsoutmcastoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutMcastOctets"] = types.YLeaf{"Ipifstatshcoutmcastoctets", ipifstatsentry.Ipifstatshcoutmcastoctets}
    ipifstatsentry.EntityData.Leafs["ipIfStatsInBcastPkts"] = types.YLeaf{"Ipifstatsinbcastpkts", ipifstatsentry.Ipifstatsinbcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCInBcastPkts"] = types.YLeaf{"Ipifstatshcinbcastpkts", ipifstatsentry.Ipifstatshcinbcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsOutBcastPkts"] = types.YLeaf{"Ipifstatsoutbcastpkts", ipifstatsentry.Ipifstatsoutbcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsHCOutBcastPkts"] = types.YLeaf{"Ipifstatshcoutbcastpkts", ipifstatsentry.Ipifstatshcoutbcastpkts}
    ipifstatsentry.EntityData.Leafs["ipIfStatsDiscontinuityTime"] = types.YLeaf{"Ipifstatsdiscontinuitytime", ipifstatsentry.Ipifstatsdiscontinuitytime}
    ipifstatsentry.EntityData.Leafs["ipIfStatsRefreshRate"] = types.YLeaf{"Ipifstatsrefreshrate", ipifstatsentry.Ipifstatsrefreshrate}
    return &(ipifstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the ipAddressPrefixTable. The type is slice of
    // IPMIB_Ipaddressprefixtable_Ipaddressprefixentry.
    Ipaddressprefixentry []IPMIB_Ipaddressprefixtable_Ipaddressprefixentry
}

func (ipaddressprefixtable *IPMIB_Ipaddressprefixtable) GetEntityData() *types.CommonEntityData {
    ipaddressprefixtable.EntityData.YFilter = ipaddressprefixtable.YFilter
    ipaddressprefixtable.EntityData.YangName = "ipAddressPrefixTable"
    ipaddressprefixtable.EntityData.BundleName = "cisco_ios_xe"
    ipaddressprefixtable.EntityData.ParentYangName = "IP-MIB"
    ipaddressprefixtable.EntityData.SegmentPath = "ipAddressPrefixTable"
    ipaddressprefixtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipaddressprefixtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipaddressprefixtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipaddressprefixtable.EntityData.Children = make(map[string]types.YChild)
    ipaddressprefixtable.EntityData.Children["ipAddressPrefixEntry"] = types.YChild{"Ipaddressprefixentry", nil}
    for i := range ipaddressprefixtable.Ipaddressprefixentry {
        ipaddressprefixtable.EntityData.Children[types.GetSegmentPath(&ipaddressprefixtable.Ipaddressprefixentry[i])] = types.YChild{"Ipaddressprefixentry", &ipaddressprefixtable.Ipaddressprefixentry[i]}
    }
    ipaddressprefixtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipaddressprefixtable.EntityData)
}

// IPMIB_Ipaddressprefixtable_Ipaddressprefixentry
// An entry in the ipAddressPrefixTable.
type IPMIB_Ipaddressprefixtable_Ipaddressprefixentry struct {
    EntityData types.CommonEntityData
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

func (ipaddressprefixentry *IPMIB_Ipaddressprefixtable_Ipaddressprefixentry) GetEntityData() *types.CommonEntityData {
    ipaddressprefixentry.EntityData.YFilter = ipaddressprefixentry.YFilter
    ipaddressprefixentry.EntityData.YangName = "ipAddressPrefixEntry"
    ipaddressprefixentry.EntityData.BundleName = "cisco_ios_xe"
    ipaddressprefixentry.EntityData.ParentYangName = "ipAddressPrefixTable"
    ipaddressprefixentry.EntityData.SegmentPath = "ipAddressPrefixEntry" + "[ipAddressPrefixIfIndex='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixifindex) + "']" + "[ipAddressPrefixType='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixtype) + "']" + "[ipAddressPrefixPrefix='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixprefix) + "']" + "[ipAddressPrefixLength='" + fmt.Sprintf("%v", ipaddressprefixentry.Ipaddressprefixlength) + "']"
    ipaddressprefixentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipaddressprefixentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipaddressprefixentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipaddressprefixentry.EntityData.Children = make(map[string]types.YChild)
    ipaddressprefixentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixIfIndex"] = types.YLeaf{"Ipaddressprefixifindex", ipaddressprefixentry.Ipaddressprefixifindex}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixType"] = types.YLeaf{"Ipaddressprefixtype", ipaddressprefixentry.Ipaddressprefixtype}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixPrefix"] = types.YLeaf{"Ipaddressprefixprefix", ipaddressprefixentry.Ipaddressprefixprefix}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixLength"] = types.YLeaf{"Ipaddressprefixlength", ipaddressprefixentry.Ipaddressprefixlength}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixOrigin"] = types.YLeaf{"Ipaddressprefixorigin", ipaddressprefixentry.Ipaddressprefixorigin}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixOnLinkFlag"] = types.YLeaf{"Ipaddressprefixonlinkflag", ipaddressprefixentry.Ipaddressprefixonlinkflag}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixAutonomousFlag"] = types.YLeaf{"Ipaddressprefixautonomousflag", ipaddressprefixentry.Ipaddressprefixautonomousflag}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixAdvPreferredLifetime"] = types.YLeaf{"Ipaddressprefixadvpreferredlifetime", ipaddressprefixentry.Ipaddressprefixadvpreferredlifetime}
    ipaddressprefixentry.EntityData.Leafs["ipAddressPrefixAdvValidLifetime"] = types.YLeaf{"Ipaddressprefixadvvalidlifetime", ipaddressprefixentry.Ipaddressprefixadvvalidlifetime}
    return &(ipaddressprefixentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An address mapping for a particular interface. The type is slice of
    // IPMIB_Ipaddresstable_Ipaddressentry.
    Ipaddressentry []IPMIB_Ipaddresstable_Ipaddressentry
}

func (ipaddresstable *IPMIB_Ipaddresstable) GetEntityData() *types.CommonEntityData {
    ipaddresstable.EntityData.YFilter = ipaddresstable.YFilter
    ipaddresstable.EntityData.YangName = "ipAddressTable"
    ipaddresstable.EntityData.BundleName = "cisco_ios_xe"
    ipaddresstable.EntityData.ParentYangName = "IP-MIB"
    ipaddresstable.EntityData.SegmentPath = "ipAddressTable"
    ipaddresstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipaddresstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipaddresstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipaddresstable.EntityData.Children = make(map[string]types.YChild)
    ipaddresstable.EntityData.Children["ipAddressEntry"] = types.YChild{"Ipaddressentry", nil}
    for i := range ipaddresstable.Ipaddressentry {
        ipaddresstable.EntityData.Children[types.GetSegmentPath(&ipaddresstable.Ipaddressentry[i])] = types.YChild{"Ipaddressentry", &ipaddresstable.Ipaddressentry[i]}
    }
    ipaddresstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipaddresstable.EntityData)
}

// IPMIB_Ipaddresstable_Ipaddressentry
// An address mapping for a particular interface.
type IPMIB_Ipaddresstable_Ipaddressentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (ipaddressentry *IPMIB_Ipaddresstable_Ipaddressentry) GetEntityData() *types.CommonEntityData {
    ipaddressentry.EntityData.YFilter = ipaddressentry.YFilter
    ipaddressentry.EntityData.YangName = "ipAddressEntry"
    ipaddressentry.EntityData.BundleName = "cisco_ios_xe"
    ipaddressentry.EntityData.ParentYangName = "ipAddressTable"
    ipaddressentry.EntityData.SegmentPath = "ipAddressEntry" + "[ipAddressAddrType='" + fmt.Sprintf("%v", ipaddressentry.Ipaddressaddrtype) + "']" + "[ipAddressAddr='" + fmt.Sprintf("%v", ipaddressentry.Ipaddressaddr) + "']"
    ipaddressentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipaddressentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipaddressentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipaddressentry.EntityData.Children = make(map[string]types.YChild)
    ipaddressentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipaddressentry.EntityData.Leafs["ipAddressAddrType"] = types.YLeaf{"Ipaddressaddrtype", ipaddressentry.Ipaddressaddrtype}
    ipaddressentry.EntityData.Leafs["ipAddressAddr"] = types.YLeaf{"Ipaddressaddr", ipaddressentry.Ipaddressaddr}
    ipaddressentry.EntityData.Leafs["ipAddressIfIndex"] = types.YLeaf{"Ipaddressifindex", ipaddressentry.Ipaddressifindex}
    ipaddressentry.EntityData.Leafs["ipAddressType"] = types.YLeaf{"Ipaddresstype", ipaddressentry.Ipaddresstype}
    ipaddressentry.EntityData.Leafs["ipAddressPrefix"] = types.YLeaf{"Ipaddressprefix", ipaddressentry.Ipaddressprefix}
    ipaddressentry.EntityData.Leafs["ipAddressOrigin"] = types.YLeaf{"Ipaddressorigin", ipaddressentry.Ipaddressorigin}
    ipaddressentry.EntityData.Leafs["ipAddressStatus"] = types.YLeaf{"Ipaddressstatus", ipaddressentry.Ipaddressstatus}
    ipaddressentry.EntityData.Leafs["ipAddressCreated"] = types.YLeaf{"Ipaddresscreated", ipaddressentry.Ipaddresscreated}
    ipaddressentry.EntityData.Leafs["ipAddressLastChanged"] = types.YLeaf{"Ipaddresslastchanged", ipaddressentry.Ipaddresslastchanged}
    ipaddressentry.EntityData.Leafs["ipAddressRowStatus"] = types.YLeaf{"Ipaddressrowstatus", ipaddressentry.Ipaddressrowstatus}
    ipaddressentry.EntityData.Leafs["ipAddressStorageType"] = types.YLeaf{"Ipaddressstoragetype", ipaddressentry.Ipaddressstoragetype}
    return &(ipaddressentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one IP address to `physical' address equivalence. The
    // type is slice of IPMIB_Ipnettophysicaltable_Ipnettophysicalentry.
    Ipnettophysicalentry []IPMIB_Ipnettophysicaltable_Ipnettophysicalentry
}

func (ipnettophysicaltable *IPMIB_Ipnettophysicaltable) GetEntityData() *types.CommonEntityData {
    ipnettophysicaltable.EntityData.YFilter = ipnettophysicaltable.YFilter
    ipnettophysicaltable.EntityData.YangName = "ipNetToPhysicalTable"
    ipnettophysicaltable.EntityData.BundleName = "cisco_ios_xe"
    ipnettophysicaltable.EntityData.ParentYangName = "IP-MIB"
    ipnettophysicaltable.EntityData.SegmentPath = "ipNetToPhysicalTable"
    ipnettophysicaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipnettophysicaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipnettophysicaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipnettophysicaltable.EntityData.Children = make(map[string]types.YChild)
    ipnettophysicaltable.EntityData.Children["ipNetToPhysicalEntry"] = types.YChild{"Ipnettophysicalentry", nil}
    for i := range ipnettophysicaltable.Ipnettophysicalentry {
        ipnettophysicaltable.EntityData.Children[types.GetSegmentPath(&ipnettophysicaltable.Ipnettophysicalentry[i])] = types.YChild{"Ipnettophysicalentry", &ipnettophysicaltable.Ipnettophysicalentry[i]}
    }
    ipnettophysicaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipnettophysicaltable.EntityData)
}

// IPMIB_Ipnettophysicaltable_Ipnettophysicalentry
// Each entry contains one IP address to `physical' address
// equivalence.
type IPMIB_Ipnettophysicaltable_Ipnettophysicalentry struct {
    EntityData types.CommonEntityData
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

func (ipnettophysicalentry *IPMIB_Ipnettophysicaltable_Ipnettophysicalentry) GetEntityData() *types.CommonEntityData {
    ipnettophysicalentry.EntityData.YFilter = ipnettophysicalentry.YFilter
    ipnettophysicalentry.EntityData.YangName = "ipNetToPhysicalEntry"
    ipnettophysicalentry.EntityData.BundleName = "cisco_ios_xe"
    ipnettophysicalentry.EntityData.ParentYangName = "ipNetToPhysicalTable"
    ipnettophysicalentry.EntityData.SegmentPath = "ipNetToPhysicalEntry" + "[ipNetToPhysicalIfIndex='" + fmt.Sprintf("%v", ipnettophysicalentry.Ipnettophysicalifindex) + "']" + "[ipNetToPhysicalNetAddressType='" + fmt.Sprintf("%v", ipnettophysicalentry.Ipnettophysicalnetaddresstype) + "']" + "[ipNetToPhysicalNetAddress='" + fmt.Sprintf("%v", ipnettophysicalentry.Ipnettophysicalnetaddress) + "']"
    ipnettophysicalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipnettophysicalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipnettophysicalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipnettophysicalentry.EntityData.Children = make(map[string]types.YChild)
    ipnettophysicalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalIfIndex"] = types.YLeaf{"Ipnettophysicalifindex", ipnettophysicalentry.Ipnettophysicalifindex}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalNetAddressType"] = types.YLeaf{"Ipnettophysicalnetaddresstype", ipnettophysicalentry.Ipnettophysicalnetaddresstype}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalNetAddress"] = types.YLeaf{"Ipnettophysicalnetaddress", ipnettophysicalentry.Ipnettophysicalnetaddress}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalPhysAddress"] = types.YLeaf{"Ipnettophysicalphysaddress", ipnettophysicalentry.Ipnettophysicalphysaddress}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalLastUpdated"] = types.YLeaf{"Ipnettophysicallastupdated", ipnettophysicalentry.Ipnettophysicallastupdated}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalType"] = types.YLeaf{"Ipnettophysicaltype", ipnettophysicalentry.Ipnettophysicaltype}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalState"] = types.YLeaf{"Ipnettophysicalstate", ipnettophysicalentry.Ipnettophysicalstate}
    ipnettophysicalentry.EntityData.Leafs["ipNetToPhysicalRowStatus"] = types.YLeaf{"Ipnettophysicalrowstatus", ipnettophysicalentry.Ipnettophysicalrowstatus}
    return &(ipnettophysicalentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the list of scope identifiers on a given interface. The
    // type is slice of IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry.
    Ipv6Scopezoneindexentry []IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry
}

func (ipv6Scopezoneindextable *IPMIB_Ipv6Scopezoneindextable) GetEntityData() *types.CommonEntityData {
    ipv6Scopezoneindextable.EntityData.YFilter = ipv6Scopezoneindextable.YFilter
    ipv6Scopezoneindextable.EntityData.YangName = "ipv6ScopeZoneIndexTable"
    ipv6Scopezoneindextable.EntityData.BundleName = "cisco_ios_xe"
    ipv6Scopezoneindextable.EntityData.ParentYangName = "IP-MIB"
    ipv6Scopezoneindextable.EntityData.SegmentPath = "ipv6ScopeZoneIndexTable"
    ipv6Scopezoneindextable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Scopezoneindextable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Scopezoneindextable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Scopezoneindextable.EntityData.Children = make(map[string]types.YChild)
    ipv6Scopezoneindextable.EntityData.Children["ipv6ScopeZoneIndexEntry"] = types.YChild{"Ipv6Scopezoneindexentry", nil}
    for i := range ipv6Scopezoneindextable.Ipv6Scopezoneindexentry {
        ipv6Scopezoneindextable.EntityData.Children[types.GetSegmentPath(&ipv6Scopezoneindextable.Ipv6Scopezoneindexentry[i])] = types.YChild{"Ipv6Scopezoneindexentry", &ipv6Scopezoneindextable.Ipv6Scopezoneindexentry[i]}
    }
    ipv6Scopezoneindextable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Scopezoneindextable.EntityData)
}

// IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry
// Each entry contains the list of scope identifiers on a given
// interface.
type IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry struct {
    EntityData types.CommonEntityData
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

func (ipv6Scopezoneindexentry *IPMIB_Ipv6Scopezoneindextable_Ipv6Scopezoneindexentry) GetEntityData() *types.CommonEntityData {
    ipv6Scopezoneindexentry.EntityData.YFilter = ipv6Scopezoneindexentry.YFilter
    ipv6Scopezoneindexentry.EntityData.YangName = "ipv6ScopeZoneIndexEntry"
    ipv6Scopezoneindexentry.EntityData.BundleName = "cisco_ios_xe"
    ipv6Scopezoneindexentry.EntityData.ParentYangName = "ipv6ScopeZoneIndexTable"
    ipv6Scopezoneindexentry.EntityData.SegmentPath = "ipv6ScopeZoneIndexEntry" + "[ipv6ScopeZoneIndexIfIndex='" + fmt.Sprintf("%v", ipv6Scopezoneindexentry.Ipv6Scopezoneindexifindex) + "']"
    ipv6Scopezoneindexentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Scopezoneindexentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Scopezoneindexentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Scopezoneindexentry.EntityData.Children = make(map[string]types.YChild)
    ipv6Scopezoneindexentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexIfIndex"] = types.YLeaf{"Ipv6Scopezoneindexifindex", ipv6Scopezoneindexentry.Ipv6Scopezoneindexifindex}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexLinkLocal"] = types.YLeaf{"Ipv6Scopezoneindexlinklocal", ipv6Scopezoneindexentry.Ipv6Scopezoneindexlinklocal}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndex3"] = types.YLeaf{"Ipv6Scopezoneindex3", ipv6Scopezoneindexentry.Ipv6Scopezoneindex3}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexAdminLocal"] = types.YLeaf{"Ipv6Scopezoneindexadminlocal", ipv6Scopezoneindexentry.Ipv6Scopezoneindexadminlocal}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexSiteLocal"] = types.YLeaf{"Ipv6Scopezoneindexsitelocal", ipv6Scopezoneindexentry.Ipv6Scopezoneindexsitelocal}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndex6"] = types.YLeaf{"Ipv6Scopezoneindex6", ipv6Scopezoneindexentry.Ipv6Scopezoneindex6}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndex7"] = types.YLeaf{"Ipv6Scopezoneindex7", ipv6Scopezoneindexentry.Ipv6Scopezoneindex7}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexOrganizationLocal"] = types.YLeaf{"Ipv6Scopezoneindexorganizationlocal", ipv6Scopezoneindexentry.Ipv6Scopezoneindexorganizationlocal}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndex9"] = types.YLeaf{"Ipv6Scopezoneindex9", ipv6Scopezoneindexentry.Ipv6Scopezoneindex9}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexA"] = types.YLeaf{"Ipv6Scopezoneindexa", ipv6Scopezoneindexentry.Ipv6Scopezoneindexa}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexB"] = types.YLeaf{"Ipv6Scopezoneindexb", ipv6Scopezoneindexentry.Ipv6Scopezoneindexb}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexC"] = types.YLeaf{"Ipv6Scopezoneindexc", ipv6Scopezoneindexentry.Ipv6Scopezoneindexc}
    ipv6Scopezoneindexentry.EntityData.Leafs["ipv6ScopeZoneIndexD"] = types.YLeaf{"Ipv6Scopezoneindexd", ipv6Scopezoneindexentry.Ipv6Scopezoneindexd}
    return &(ipv6Scopezoneindexentry.EntityData)
}

// IPMIB_Ipdefaultroutertable
// The table used to describe the default routers known to this
// 
// 
// entity.
type IPMIB_Ipdefaultroutertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains information about a default router known to this
    // entity. The type is slice of
    // IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry.
    Ipdefaultrouterentry []IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry
}

func (ipdefaultroutertable *IPMIB_Ipdefaultroutertable) GetEntityData() *types.CommonEntityData {
    ipdefaultroutertable.EntityData.YFilter = ipdefaultroutertable.YFilter
    ipdefaultroutertable.EntityData.YangName = "ipDefaultRouterTable"
    ipdefaultroutertable.EntityData.BundleName = "cisco_ios_xe"
    ipdefaultroutertable.EntityData.ParentYangName = "IP-MIB"
    ipdefaultroutertable.EntityData.SegmentPath = "ipDefaultRouterTable"
    ipdefaultroutertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipdefaultroutertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipdefaultroutertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipdefaultroutertable.EntityData.Children = make(map[string]types.YChild)
    ipdefaultroutertable.EntityData.Children["ipDefaultRouterEntry"] = types.YChild{"Ipdefaultrouterentry", nil}
    for i := range ipdefaultroutertable.Ipdefaultrouterentry {
        ipdefaultroutertable.EntityData.Children[types.GetSegmentPath(&ipdefaultroutertable.Ipdefaultrouterentry[i])] = types.YChild{"Ipdefaultrouterentry", &ipdefaultroutertable.Ipdefaultrouterentry[i]}
    }
    ipdefaultroutertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipdefaultroutertable.EntityData)
}

// IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry
// Each entry contains information about a default router known
// to this entity.
type IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry struct {
    EntityData types.CommonEntityData
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

func (ipdefaultrouterentry *IPMIB_Ipdefaultroutertable_Ipdefaultrouterentry) GetEntityData() *types.CommonEntityData {
    ipdefaultrouterentry.EntityData.YFilter = ipdefaultrouterentry.YFilter
    ipdefaultrouterentry.EntityData.YangName = "ipDefaultRouterEntry"
    ipdefaultrouterentry.EntityData.BundleName = "cisco_ios_xe"
    ipdefaultrouterentry.EntityData.ParentYangName = "ipDefaultRouterTable"
    ipdefaultrouterentry.EntityData.SegmentPath = "ipDefaultRouterEntry" + "[ipDefaultRouterAddressType='" + fmt.Sprintf("%v", ipdefaultrouterentry.Ipdefaultrouteraddresstype) + "']" + "[ipDefaultRouterAddress='" + fmt.Sprintf("%v", ipdefaultrouterentry.Ipdefaultrouteraddress) + "']" + "[ipDefaultRouterIfIndex='" + fmt.Sprintf("%v", ipdefaultrouterentry.Ipdefaultrouterifindex) + "']"
    ipdefaultrouterentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipdefaultrouterentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipdefaultrouterentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipdefaultrouterentry.EntityData.Children = make(map[string]types.YChild)
    ipdefaultrouterentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipdefaultrouterentry.EntityData.Leafs["ipDefaultRouterAddressType"] = types.YLeaf{"Ipdefaultrouteraddresstype", ipdefaultrouterentry.Ipdefaultrouteraddresstype}
    ipdefaultrouterentry.EntityData.Leafs["ipDefaultRouterAddress"] = types.YLeaf{"Ipdefaultrouteraddress", ipdefaultrouterentry.Ipdefaultrouteraddress}
    ipdefaultrouterentry.EntityData.Leafs["ipDefaultRouterIfIndex"] = types.YLeaf{"Ipdefaultrouterifindex", ipdefaultrouterentry.Ipdefaultrouterifindex}
    ipdefaultrouterentry.EntityData.Leafs["ipDefaultRouterLifetime"] = types.YLeaf{"Ipdefaultrouterlifetime", ipdefaultrouterentry.Ipdefaultrouterlifetime}
    ipdefaultrouterentry.EntityData.Leafs["ipDefaultRouterPreference"] = types.YLeaf{"Ipdefaultrouterpreference", ipdefaultrouterentry.Ipdefaultrouterpreference}
    return &(ipdefaultrouterentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information used to construct router advertisements. 
    // Information in this table is persistent, and when this object is written,
    // the entity SHOULD save the change to non-volatile storage. The type is
    // slice of IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry.
    Ipv6Routeradvertentry []IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry
}

func (ipv6Routeradverttable *IPMIB_Ipv6Routeradverttable) GetEntityData() *types.CommonEntityData {
    ipv6Routeradverttable.EntityData.YFilter = ipv6Routeradverttable.YFilter
    ipv6Routeradverttable.EntityData.YangName = "ipv6RouterAdvertTable"
    ipv6Routeradverttable.EntityData.BundleName = "cisco_ios_xe"
    ipv6Routeradverttable.EntityData.ParentYangName = "IP-MIB"
    ipv6Routeradverttable.EntityData.SegmentPath = "ipv6RouterAdvertTable"
    ipv6Routeradverttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Routeradverttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Routeradverttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Routeradverttable.EntityData.Children = make(map[string]types.YChild)
    ipv6Routeradverttable.EntityData.Children["ipv6RouterAdvertEntry"] = types.YChild{"Ipv6Routeradvertentry", nil}
    for i := range ipv6Routeradverttable.Ipv6Routeradvertentry {
        ipv6Routeradverttable.EntityData.Children[types.GetSegmentPath(&ipv6Routeradverttable.Ipv6Routeradvertentry[i])] = types.YChild{"Ipv6Routeradvertentry", &ipv6Routeradverttable.Ipv6Routeradvertentry[i]}
    }
    ipv6Routeradverttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Routeradverttable.EntityData)
}

// IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry
// An entry containing information used to construct router
// advertisements.
// 
// Information in this table is persistent, and when this
// object is written, the entity SHOULD save the change to
// non-volatile storage.
type IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry struct {
    EntityData types.CommonEntityData
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

func (ipv6Routeradvertentry *IPMIB_Ipv6Routeradverttable_Ipv6Routeradvertentry) GetEntityData() *types.CommonEntityData {
    ipv6Routeradvertentry.EntityData.YFilter = ipv6Routeradvertentry.YFilter
    ipv6Routeradvertentry.EntityData.YangName = "ipv6RouterAdvertEntry"
    ipv6Routeradvertentry.EntityData.BundleName = "cisco_ios_xe"
    ipv6Routeradvertentry.EntityData.ParentYangName = "ipv6RouterAdvertTable"
    ipv6Routeradvertentry.EntityData.SegmentPath = "ipv6RouterAdvertEntry" + "[ipv6RouterAdvertIfIndex='" + fmt.Sprintf("%v", ipv6Routeradvertentry.Ipv6Routeradvertifindex) + "']"
    ipv6Routeradvertentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Routeradvertentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Routeradvertentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Routeradvertentry.EntityData.Children = make(map[string]types.YChild)
    ipv6Routeradvertentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertIfIndex"] = types.YLeaf{"Ipv6Routeradvertifindex", ipv6Routeradvertentry.Ipv6Routeradvertifindex}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertSendAdverts"] = types.YLeaf{"Ipv6Routeradvertsendadverts", ipv6Routeradvertentry.Ipv6Routeradvertsendadverts}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertMaxInterval"] = types.YLeaf{"Ipv6Routeradvertmaxinterval", ipv6Routeradvertentry.Ipv6Routeradvertmaxinterval}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertMinInterval"] = types.YLeaf{"Ipv6Routeradvertmininterval", ipv6Routeradvertentry.Ipv6Routeradvertmininterval}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertManagedFlag"] = types.YLeaf{"Ipv6Routeradvertmanagedflag", ipv6Routeradvertentry.Ipv6Routeradvertmanagedflag}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertOtherConfigFlag"] = types.YLeaf{"Ipv6Routeradvertotherconfigflag", ipv6Routeradvertentry.Ipv6Routeradvertotherconfigflag}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertLinkMTU"] = types.YLeaf{"Ipv6Routeradvertlinkmtu", ipv6Routeradvertentry.Ipv6Routeradvertlinkmtu}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertReachableTime"] = types.YLeaf{"Ipv6Routeradvertreachabletime", ipv6Routeradvertentry.Ipv6Routeradvertreachabletime}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertRetransmitTime"] = types.YLeaf{"Ipv6Routeradvertretransmittime", ipv6Routeradvertentry.Ipv6Routeradvertretransmittime}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertCurHopLimit"] = types.YLeaf{"Ipv6Routeradvertcurhoplimit", ipv6Routeradvertentry.Ipv6Routeradvertcurhoplimit}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertDefaultLifetime"] = types.YLeaf{"Ipv6Routeradvertdefaultlifetime", ipv6Routeradvertentry.Ipv6Routeradvertdefaultlifetime}
    ipv6Routeradvertentry.EntityData.Leafs["ipv6RouterAdvertRowStatus"] = types.YLeaf{"Ipv6Routeradvertrowstatus", ipv6Routeradvertentry.Ipv6Routeradvertrowstatus}
    return &(ipv6Routeradvertentry.EntityData)
}

// IPMIB_Icmpstatstable
// The table of generic system-wide ICMP counters.
type IPMIB_Icmpstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the icmpStatsTable. The type is slice of
    // IPMIB_Icmpstatstable_Icmpstatsentry.
    Icmpstatsentry []IPMIB_Icmpstatstable_Icmpstatsentry
}

func (icmpstatstable *IPMIB_Icmpstatstable) GetEntityData() *types.CommonEntityData {
    icmpstatstable.EntityData.YFilter = icmpstatstable.YFilter
    icmpstatstable.EntityData.YangName = "icmpStatsTable"
    icmpstatstable.EntityData.BundleName = "cisco_ios_xe"
    icmpstatstable.EntityData.ParentYangName = "IP-MIB"
    icmpstatstable.EntityData.SegmentPath = "icmpStatsTable"
    icmpstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpstatstable.EntityData.Children = make(map[string]types.YChild)
    icmpstatstable.EntityData.Children["icmpStatsEntry"] = types.YChild{"Icmpstatsentry", nil}
    for i := range icmpstatstable.Icmpstatsentry {
        icmpstatstable.EntityData.Children[types.GetSegmentPath(&icmpstatstable.Icmpstatsentry[i])] = types.YChild{"Icmpstatsentry", &icmpstatstable.Icmpstatsentry[i]}
    }
    icmpstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(icmpstatstable.EntityData)
}

// IPMIB_Icmpstatstable_Icmpstatsentry
// A conceptual row in the icmpStatsTable.
type IPMIB_Icmpstatstable_Icmpstatsentry struct {
    EntityData types.CommonEntityData
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

func (icmpstatsentry *IPMIB_Icmpstatstable_Icmpstatsentry) GetEntityData() *types.CommonEntityData {
    icmpstatsentry.EntityData.YFilter = icmpstatsentry.YFilter
    icmpstatsentry.EntityData.YangName = "icmpStatsEntry"
    icmpstatsentry.EntityData.BundleName = "cisco_ios_xe"
    icmpstatsentry.EntityData.ParentYangName = "icmpStatsTable"
    icmpstatsentry.EntityData.SegmentPath = "icmpStatsEntry" + "[icmpStatsIPVersion='" + fmt.Sprintf("%v", icmpstatsentry.Icmpstatsipversion) + "']"
    icmpstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpstatsentry.EntityData.Children = make(map[string]types.YChild)
    icmpstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpstatsentry.EntityData.Leafs["icmpStatsIPVersion"] = types.YLeaf{"Icmpstatsipversion", icmpstatsentry.Icmpstatsipversion}
    icmpstatsentry.EntityData.Leafs["icmpStatsInMsgs"] = types.YLeaf{"Icmpstatsinmsgs", icmpstatsentry.Icmpstatsinmsgs}
    icmpstatsentry.EntityData.Leafs["icmpStatsInErrors"] = types.YLeaf{"Icmpstatsinerrors", icmpstatsentry.Icmpstatsinerrors}
    icmpstatsentry.EntityData.Leafs["icmpStatsOutMsgs"] = types.YLeaf{"Icmpstatsoutmsgs", icmpstatsentry.Icmpstatsoutmsgs}
    icmpstatsentry.EntityData.Leafs["icmpStatsOutErrors"] = types.YLeaf{"Icmpstatsouterrors", icmpstatsentry.Icmpstatsouterrors}
    return &(icmpstatsentry.EntityData)
}

// IPMIB_Icmpmsgstatstable
// The table of system-wide per-version, per-message type ICMP
// counters.
type IPMIB_Icmpmsgstatstable struct {
    EntityData types.CommonEntityData
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

func (icmpmsgstatstable *IPMIB_Icmpmsgstatstable) GetEntityData() *types.CommonEntityData {
    icmpmsgstatstable.EntityData.YFilter = icmpmsgstatstable.YFilter
    icmpmsgstatstable.EntityData.YangName = "icmpMsgStatsTable"
    icmpmsgstatstable.EntityData.BundleName = "cisco_ios_xe"
    icmpmsgstatstable.EntityData.ParentYangName = "IP-MIB"
    icmpmsgstatstable.EntityData.SegmentPath = "icmpMsgStatsTable"
    icmpmsgstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpmsgstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpmsgstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpmsgstatstable.EntityData.Children = make(map[string]types.YChild)
    icmpmsgstatstable.EntityData.Children["icmpMsgStatsEntry"] = types.YChild{"Icmpmsgstatsentry", nil}
    for i := range icmpmsgstatstable.Icmpmsgstatsentry {
        icmpmsgstatstable.EntityData.Children[types.GetSegmentPath(&icmpmsgstatstable.Icmpmsgstatsentry[i])] = types.YChild{"Icmpmsgstatsentry", &icmpmsgstatstable.Icmpmsgstatsentry[i]}
    }
    icmpmsgstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(icmpmsgstatstable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (icmpmsgstatsentry *IPMIB_Icmpmsgstatstable_Icmpmsgstatsentry) GetEntityData() *types.CommonEntityData {
    icmpmsgstatsentry.EntityData.YFilter = icmpmsgstatsentry.YFilter
    icmpmsgstatsentry.EntityData.YangName = "icmpMsgStatsEntry"
    icmpmsgstatsentry.EntityData.BundleName = "cisco_ios_xe"
    icmpmsgstatsentry.EntityData.ParentYangName = "icmpMsgStatsTable"
    icmpmsgstatsentry.EntityData.SegmentPath = "icmpMsgStatsEntry" + "[icmpMsgStatsIPVersion='" + fmt.Sprintf("%v", icmpmsgstatsentry.Icmpmsgstatsipversion) + "']" + "[icmpMsgStatsType='" + fmt.Sprintf("%v", icmpmsgstatsentry.Icmpmsgstatstype) + "']"
    icmpmsgstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpmsgstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpmsgstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpmsgstatsentry.EntityData.Children = make(map[string]types.YChild)
    icmpmsgstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpmsgstatsentry.EntityData.Leafs["icmpMsgStatsIPVersion"] = types.YLeaf{"Icmpmsgstatsipversion", icmpmsgstatsentry.Icmpmsgstatsipversion}
    icmpmsgstatsentry.EntityData.Leafs["icmpMsgStatsType"] = types.YLeaf{"Icmpmsgstatstype", icmpmsgstatsentry.Icmpmsgstatstype}
    icmpmsgstatsentry.EntityData.Leafs["icmpMsgStatsInPkts"] = types.YLeaf{"Icmpmsgstatsinpkts", icmpmsgstatsentry.Icmpmsgstatsinpkts}
    icmpmsgstatsentry.EntityData.Leafs["icmpMsgStatsOutPkts"] = types.YLeaf{"Icmpmsgstatsoutpkts", icmpmsgstatsentry.Icmpmsgstatsoutpkts}
    return &(icmpmsgstatsentry.EntityData)
}

