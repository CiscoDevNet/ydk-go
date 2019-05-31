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

    
    IpTrafficStats IPMIB_IpTrafficStats

    
    Icmp IPMIB_Icmp

    // The table of addressing information relevant to this entity's IPv4
    // addresses.  This table has been deprecated, as a new IP version-neutral
    // table has been added.  It is loosely replaced by the ipAddressTable
    // although several objects that weren't deemed useful weren't carried forward
    // while another (ipAdEntReasmMaxSize) was moved to the ipv4InterfaceTable.
    IpAddrTable IPMIB_IpAddrTable

    // The IPv4 Address Translation table used for mapping from IPv4 addresses to
    // physical addresses.  This table has been deprecated, as a new IP
    // version-neutral table has been added.  It is loosely replaced by the
    // ipNetToPhysicalTable.
    IpNetToMediaTable IPMIB_IpNetToMediaTable

    // The table containing per-interface IPv4-specific information.
    Ipv4InterfaceTable IPMIB_Ipv4InterfaceTable

    // The table containing per-interface IPv6-specific information.
    Ipv6InterfaceTable IPMIB_Ipv6InterfaceTable

    // The table containing system wide, IP version specific traffic statistics. 
    // This table and the ipIfStatsTable contain similar objects whose difference
    // is in their granularity.  Where this table contains system wide traffic
    // statistics, the ipIfStatsTable contains the same statistics but counted on
    // a per-interface basis.
    IpSystemStatsTable IPMIB_IpSystemStatsTable

    // The table containing per-interface traffic statistics.  This table and the
    // ipSystemStatsTable contain similar objects whose difference is in their
    // granularity.  Where this table contains per-interface statistics, the
    // ipSystemStatsTable contains the same statistics, but counted on a system
    // wide basis.
    IpIfStatsTable IPMIB_IpIfStatsTable

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
    IpAddressPrefixTable IPMIB_IpAddressPrefixTable

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
    IpAddressTable IPMIB_IpAddressTable

    // The IP Address Translation table used for mapping from IP addresses to
    // physical addresses.  The Address Translation tables contain the IP address
    // to 'physical' address equivalences.  Some interfaces do not use translation
    // tables for determining address equivalences (e.g., DDN-X.25 has an
    // algorithmic method); if all interfaces are of this type, then the Address
    // Translation table is empty, i.e., has zero entries.  While many protocols
    // may be used to populate this table, ARP and Neighbor Discovery are the most
    // likely options.
    IpNetToPhysicalTable IPMIB_IpNetToPhysicalTable

    // The table used to describe IPv6 unicast and multicast scope zones.  For
    // those objects that have names rather than numbers, the names were chosen to
    // coincide with the names used in the IPv6 address architecture document. .
    Ipv6ScopeZoneIndexTable IPMIB_Ipv6ScopeZoneIndexTable

    // The table used to describe the default routers known to this   entity.
    IpDefaultRouterTable IPMIB_IpDefaultRouterTable

    // The table containing information used to construct router advertisements.
    Ipv6RouterAdvertTable IPMIB_Ipv6RouterAdvertTable

    // The table of generic system-wide ICMP counters.
    IcmpStatsTable IPMIB_IcmpStatsTable

    // The table of system-wide per-version, per-message type ICMP counters.
    IcmpMsgStatsTable IPMIB_IcmpMsgStatsTable
}

func (iPMIB *IPMIB) GetEntityData() *types.CommonEntityData {
    iPMIB.EntityData.YFilter = iPMIB.YFilter
    iPMIB.EntityData.YangName = "IP-MIB"
    iPMIB.EntityData.BundleName = "cisco_ios_xe"
    iPMIB.EntityData.ParentYangName = "IP-MIB"
    iPMIB.EntityData.SegmentPath = "IP-MIB:IP-MIB"
    iPMIB.EntityData.AbsolutePath = iPMIB.EntityData.SegmentPath
    iPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iPMIB.EntityData.Children = types.NewOrderedMap()
    iPMIB.EntityData.Children.Append("ip", types.YChild{"Ip", &iPMIB.Ip})
    iPMIB.EntityData.Children.Append("ipTrafficStats", types.YChild{"IpTrafficStats", &iPMIB.IpTrafficStats})
    iPMIB.EntityData.Children.Append("icmp", types.YChild{"Icmp", &iPMIB.Icmp})
    iPMIB.EntityData.Children.Append("ipAddrTable", types.YChild{"IpAddrTable", &iPMIB.IpAddrTable})
    iPMIB.EntityData.Children.Append("ipNetToMediaTable", types.YChild{"IpNetToMediaTable", &iPMIB.IpNetToMediaTable})
    iPMIB.EntityData.Children.Append("ipv4InterfaceTable", types.YChild{"Ipv4InterfaceTable", &iPMIB.Ipv4InterfaceTable})
    iPMIB.EntityData.Children.Append("ipv6InterfaceTable", types.YChild{"Ipv6InterfaceTable", &iPMIB.Ipv6InterfaceTable})
    iPMIB.EntityData.Children.Append("ipSystemStatsTable", types.YChild{"IpSystemStatsTable", &iPMIB.IpSystemStatsTable})
    iPMIB.EntityData.Children.Append("ipIfStatsTable", types.YChild{"IpIfStatsTable", &iPMIB.IpIfStatsTable})
    iPMIB.EntityData.Children.Append("ipAddressPrefixTable", types.YChild{"IpAddressPrefixTable", &iPMIB.IpAddressPrefixTable})
    iPMIB.EntityData.Children.Append("ipAddressTable", types.YChild{"IpAddressTable", &iPMIB.IpAddressTable})
    iPMIB.EntityData.Children.Append("ipNetToPhysicalTable", types.YChild{"IpNetToPhysicalTable", &iPMIB.IpNetToPhysicalTable})
    iPMIB.EntityData.Children.Append("ipv6ScopeZoneIndexTable", types.YChild{"Ipv6ScopeZoneIndexTable", &iPMIB.Ipv6ScopeZoneIndexTable})
    iPMIB.EntityData.Children.Append("ipDefaultRouterTable", types.YChild{"IpDefaultRouterTable", &iPMIB.IpDefaultRouterTable})
    iPMIB.EntityData.Children.Append("ipv6RouterAdvertTable", types.YChild{"Ipv6RouterAdvertTable", &iPMIB.Ipv6RouterAdvertTable})
    iPMIB.EntityData.Children.Append("icmpStatsTable", types.YChild{"IcmpStatsTable", &iPMIB.IcmpStatsTable})
    iPMIB.EntityData.Children.Append("icmpMsgStatsTable", types.YChild{"IcmpMsgStatsTable", &iPMIB.IcmpMsgStatsTable})
    iPMIB.EntityData.Leafs = types.NewOrderedMap()

    iPMIB.EntityData.YListKeys = []string {}

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
    // type is IpForwarding.
    IpForwarding interface{}

    // The default value inserted into the Time-To-Live field of the IPv4 header
    // of datagrams originated at this entity, whenever a TTL value is not
    // supplied by the transport layer   protocol.  When this object is written,
    // the entity should save the change to non-volatile storage and restore the
    // object from non-volatile storage upon re-initialization of the system.
    // Note: a stronger requirement is not used because this object was previously
    // defined. The type is interface{} with range: 1..255.
    IpDefaultTTL interface{}

    // The total number of input datagrams received from interfaces, including
    // those received in error.  This object has been deprecated, as a new IP
    // version-neutral   table has been added.  It is loosely replaced by
    // ipSystemStatsInRecieves. The type is interface{} with range: 0..4294967295.
    IpInReceives interface{}

    // The number of input datagrams discarded due to errors in their IPv4
    // headers, including bad checksums, version number mismatch, other format
    // errors, time-to-live exceeded, errors discovered in processing their IPv4
    // options, etc.  This object has been deprecated as a new IP version-neutral
    // table has been added.  It is loosely replaced by ipSystemStatsInHdrErrors.
    // The type is interface{} with range: 0..4294967295.
    IpInHdrErrors interface{}

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
    IpInAddrErrors interface{}

    // The number of input datagrams for which this entity was not their final
    // IPv4 destination, as a result of which an attempt was made to find a route
    // to forward them to that final destination.  In entities which do not act as
    // IPv4 routers, this counter will include only those packets which   were
    // Source-Routed via this entity, and the Source-Route option processing was
    // successful.  This object has been deprecated, as a new IP version-neutral
    // table has been added.  It is loosely replaced by
    // ipSystemStatsInForwDatagrams. The type is interface{} with range:
    // 0..4294967295.
    IpForwDatagrams interface{}

    // The number of locally-addressed datagrams received successfully but
    // discarded because of an unknown or unsupported protocol.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by ipSystemStatsInUnknownProtos. The type is interface{}
    // with range: 0..4294967295.
    IpInUnknownProtos interface{}

    // The number of input IPv4 datagrams for which no problems were encountered
    // to prevent their continued processing, but which were discarded (e.g., for
    // lack of buffer space).  Note that this counter does not include any
    // datagrams discarded while awaiting re-assembly.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by ipSystemStatsInDiscards. The type is interface{} with
    // range: 0..4294967295.
    IpInDiscards interface{}

    // The total number of input datagrams successfully delivered to IPv4
    // user-protocols (including ICMP).  This object has been deprecated as a new
    // IP version neutral table has been added.  It is loosely replaced by  
    // ipSystemStatsIndelivers. The type is interface{} with range: 0..4294967295.
    IpInDelivers interface{}

    // The total number of IPv4 datagrams which local IPv4 user protocols
    // (including ICMP) supplied to IPv4 in requests for transmission.  Note that
    // this counter does not include any datagrams counted in ipForwDatagrams. 
    // This object has been deprecated, as a new IP version-neutral table has been
    // added.  It is loosely replaced by ipSystemStatsOutRequests. The type is
    // interface{} with range: 0..4294967295.
    IpOutRequests interface{}

    // The number of output IPv4 datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but which were discarded
    // (e.g., for lack of buffer space).  Note that this counter would include
    // datagrams counted in ipForwDatagrams if any such packets met this
    // (discretionary) discard criterion.  This object has been deprecated, as a
    // new IP version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsOutDiscards. The type is interface{} with range:
    // 0..4294967295.
    IpOutDiscards interface{}

    // The number of IPv4 datagrams discarded because no route could be found to
    // transmit them to their destination.  Note that this counter includes any
    // packets counted in ipForwDatagrams which meet this `no-route' criterion. 
    // Note that this includes any datagrams which a host cannot route because all
    // of its default routers are down.  This object has been deprecated, as a new
    // IP version-neutral   table has been added.  It is loosely replaced by
    // ipSystemStatsOutNoRoutes. The type is interface{} with range:
    // 0..4294967295.
    IpOutNoRoutes interface{}

    // The maximum number of seconds that received fragments are held while they
    // are awaiting reassembly at this entity. The type is interface{} with range:
    // -2147483648..2147483647. Units are seconds.
    IpReasmTimeout interface{}

    // The number of IPv4 fragments received which needed to be reassembled at
    // this entity.  This object has been deprecated, as a new IP version-neutral
    // table has been added.  It is loosely replaced by ipSystemStatsReasmReqds.
    // The type is interface{} with range: 0..4294967295.
    IpReasmReqds interface{}

    // The number of IPv4 datagrams successfully re-assembled.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by ipSystemStatsReasmOKs. The type is interface{} with
    // range: 0..4294967295.
    IpReasmOKs interface{}

    // The number of failures detected by the IPv4 re-assembly algorithm (for
    // whatever reason: timed out, errors, etc). Note that this is not necessarily
    // a count of discarded IPv4 fragments since some algorithms (notably the
    // algorithm in RFC 815) can lose track of the number of fragments by
    // combining them as they are received.  This object has been deprecated, as a
    // new IP version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsReasmFails. The type is interface{} with range: 0..4294967295.
    IpReasmFails interface{}

    // The number of IPv4 datagrams that have been successfully fragmented at this
    // entity.  This object has been deprecated, as a new IP version-neutral table
    // has been added.  It is loosely replaced by ipSystemStatsOutFragOKs. The
    // type is interface{} with range: 0..4294967295.
    IpFragOKs interface{}

    // The number of IPv4 datagrams that have been discarded because they needed
    // to be fragmented at this entity but could not be, e.g., because their Don't
    // Fragment flag was set.  This object has been deprecated, as a new IP
    // version-neutral table has been added.  It is loosely replaced by
    // ipSystemStatsOutFragFails. The type is interface{} with range:
    // 0..4294967295.
    IpFragFails interface{}

    // The number of IPv4 datagram fragments that have been generated as a result
    // of fragmentation at this entity.  This object has been deprecated as a new
    // IP version neutral table has been added.  It is loosely replaced by
    // ipSystemStatsOutFragCreates. The type is interface{} with range:
    // 0..4294967295.
    IpFragCreates interface{}

    // The number of routing entries which were chosen to be discarded even though
    // they are valid.  One possible reason for discarding such an entry could be
    // to free-up buffer space for other routing entries.   This object was
    // defined in pre-IPv6 versions of the IP MIB. It was implicitly IPv4 only,
    // but the original specifications did not indicate this protocol restriction.
    // In order to clarify the specifications, this object has been deprecated and
    // a similar, but more thoroughly clarified, object has been added to the
    // IP-FORWARD-MIB. The type is interface{} with range: 0..4294967295.
    IpRoutingDiscards interface{}

    // The indication of whether this entity is acting as an IPv6 router on any
    // interface in respect to the forwarding of datagrams received by, but not
    // addressed to, this entity. IPv6 routers forward datagrams.  IPv6 hosts do
    // not (except those source-routed via the host).  When this object is
    // written, the entity SHOULD save the change to non-volatile storage and
    // restore the object from non-volatile storage upon re-initialization of the
    // system. The type is Ipv6IpForwarding.
    Ipv6IpForwarding interface{}

    // The default value inserted into the Hop Limit field of the IPv6 header of
    // datagrams originated at this entity whenever a Hop Limit value is not
    // supplied by the transport layer protocol.  When this object is written, the
    // entity SHOULD save the change to non-volatile storage and restore the
    // object from non-volatile storage upon re-initialization of the system. The
    // type is interface{} with range: 0..255.
    Ipv6IpDefaultHopLimit interface{}

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipv4InterfaceTable was added or deleted, or when an
    // ipv4InterfaceReasmMaxSize or an ipv4InterfaceEnableStatus object was
    // modified.  If new objects are added to the ipv4InterfaceTable that require
    // the ipv4InterfaceTableLastChange to be updated when they are modified, they
    // must specify that requirement in their description clause. The type is
    // interface{} with range: 0..4294967295.
    Ipv4InterfaceTableLastChange interface{}

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipv6InterfaceTable was added or deleted or when an
    // ipv6InterfaceReasmMaxSize, ipv6InterfaceIdentifier,
    // ipv6InterfaceEnableStatus, ipv6InterfaceReachableTime,
    // ipv6InterfaceRetransmitTime, or ipv6InterfaceForwarding object was
    // modified.  If new objects are added to the ipv6InterfaceTable that require
    // the ipv6InterfaceTableLastChange to be updated when they are modified, they
    // must specify that requirement in their description clause. The type is
    // interface{} with range: 0..4294967295.
    Ipv6InterfaceTableLastChange interface{}

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
    IpAddressSpinLock interface{}

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
    Ipv6RouterAdvertSpinLock interface{}
}

func (ip *IPMIB_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xe"
    ip.EntityData.ParentYangName = "IP-MIB"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ip.EntityData.SegmentPath
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
    ip.EntityData.Leafs.Append("ipv6IpForwarding", types.YLeaf{"Ipv6IpForwarding", ip.Ipv6IpForwarding})
    ip.EntityData.Leafs.Append("ipv6IpDefaultHopLimit", types.YLeaf{"Ipv6IpDefaultHopLimit", ip.Ipv6IpDefaultHopLimit})
    ip.EntityData.Leafs.Append("ipv4InterfaceTableLastChange", types.YLeaf{"Ipv4InterfaceTableLastChange", ip.Ipv4InterfaceTableLastChange})
    ip.EntityData.Leafs.Append("ipv6InterfaceTableLastChange", types.YLeaf{"Ipv6InterfaceTableLastChange", ip.Ipv6InterfaceTableLastChange})
    ip.EntityData.Leafs.Append("ipAddressSpinLock", types.YLeaf{"IpAddressSpinLock", ip.IpAddressSpinLock})
    ip.EntityData.Leafs.Append("ipv6RouterAdvertSpinLock", types.YLeaf{"Ipv6RouterAdvertSpinLock", ip.Ipv6RouterAdvertSpinLock})

    ip.EntityData.YListKeys = []string {}

    return &(ip.EntityData)
}

// IPMIB_Ip_IpForwarding represents was previously defined.
type IPMIB_Ip_IpForwarding string

const (
    IPMIB_Ip_IpForwarding_forwarding IPMIB_Ip_IpForwarding = "forwarding"

    IPMIB_Ip_IpForwarding_notForwarding IPMIB_Ip_IpForwarding = "notForwarding"
)

// IPMIB_Ip_Ipv6IpForwarding represents non-volatile storage upon re-initialization of the system.
type IPMIB_Ip_Ipv6IpForwarding string

const (
    IPMIB_Ip_Ipv6IpForwarding_forwarding IPMIB_Ip_Ipv6IpForwarding = "forwarding"

    IPMIB_Ip_Ipv6IpForwarding_notForwarding IPMIB_Ip_Ipv6IpForwarding = "notForwarding"
)

// IPMIB_IpTrafficStats
type IPMIB_IpTrafficStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime on the most recent occasion at which a row in the
    // ipIfStatsTable was added or deleted.  If new objects are added to the
    // ipIfStatsTable that require the ipIfStatsTableLastChange to be updated when
    // they are modified, they must specify that requirement in their description
    // clause. The type is interface{} with range: 0..4294967295.
    IpIfStatsTableLastChange interface{}
}

func (ipTrafficStats *IPMIB_IpTrafficStats) GetEntityData() *types.CommonEntityData {
    ipTrafficStats.EntityData.YFilter = ipTrafficStats.YFilter
    ipTrafficStats.EntityData.YangName = "ipTrafficStats"
    ipTrafficStats.EntityData.BundleName = "cisco_ios_xe"
    ipTrafficStats.EntityData.ParentYangName = "IP-MIB"
    ipTrafficStats.EntityData.SegmentPath = "ipTrafficStats"
    ipTrafficStats.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipTrafficStats.EntityData.SegmentPath
    ipTrafficStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipTrafficStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipTrafficStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipTrafficStats.EntityData.Children = types.NewOrderedMap()
    ipTrafficStats.EntityData.Leafs = types.NewOrderedMap()
    ipTrafficStats.EntityData.Leafs.Append("ipIfStatsTableLastChange", types.YLeaf{"IpIfStatsTableLastChange", ipTrafficStats.IpIfStatsTableLastChange})

    ipTrafficStats.EntityData.YListKeys = []string {}

    return &(ipTrafficStats.EntityData)
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
    IcmpInMsgs interface{}

    // The number of ICMP messages which the entity received but determined as
    // having ICMP-specific errors (bad ICMP checksums, bad length, etc.).  This
    // object has been deprecated, as a new IP version-neutral table has been
    // added.  It is loosely replaced by icmpStatsInErrors. The type is
    // interface{} with range: 0..4294967295.
    IcmpInErrors interface{}

    // The number of ICMP Destination Unreachable messages received.  This object
    // has been deprecated, as a new IP version-neutral table has been added.  It
    // is loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInDestUnreachs interface{}

    // The number of ICMP Time Exceeded messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInTimeExcds interface{}

    // The number of ICMP Parameter Problem messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInParmProbs interface{}

    // The number of ICMP Source Quench messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInSrcQuenchs interface{}

    // The number of ICMP Redirect messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInRedirects interface{}

    // The number of ICMP Echo (request) messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInEchos interface{}

    // The number of ICMP Echo Reply messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInEchoReps interface{}

    // The number of ICMP Timestamp (request) messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInTimestamps interface{}

    // The number of ICMP Timestamp Reply messages received.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInTimestampReps interface{}

    // The number of ICMP Address Mask Request messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInAddrMasks interface{}

    // The number of ICMP Address Mask Reply messages received.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpInAddrMaskReps interface{}

    // The total number of ICMP messages which this entity attempted to send. 
    // Note that this counter includes all those counted by icmpOutErrors.  This
    // object has been deprecated, as a new IP version-neutral table has been
    // added.  It is loosely replaced by icmpStatsOutMsgs. The type is interface{}
    // with range: 0..4294967295.
    IcmpOutMsgs interface{}

    // The number of ICMP messages which this entity did not send due to problems
    // discovered within ICMP, such as a lack of buffers.  This value should not
    // include errors discovered outside the ICMP layer, such as the inability of
    // IP to route the resultant datagram.  In some implementations, there may be
    // no types of error which contribute to this counter's value.  This object
    // has been deprecated, as a new IP version-neutral table has been added.  It
    // is loosely replaced by icmpStatsOutErrors. The type is interface{} with
    // range: 0..4294967295.
    IcmpOutErrors interface{}

    // The number of ICMP Destination Unreachable messages sent.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutDestUnreachs interface{}

    // The number of ICMP Time Exceeded messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutTimeExcds interface{}

    // The number of ICMP Parameter Problem messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutParmProbs interface{}

    // The number of ICMP Source Quench messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutSrcQuenchs interface{}

    // The number of ICMP Redirect messages sent.  For a host, this object will
    // always be zero, since hosts do not send redirects.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutRedirects interface{}

    // The number of ICMP Echo (request) messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutEchos interface{}

    // The number of ICMP Echo Reply messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutEchoReps interface{}

    // The number of ICMP Timestamp (request) messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutTimestamps interface{}

    // The number of ICMP Timestamp Reply messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutTimestampReps interface{}

    // The number of ICMP Address Mask Request messages sent.  This object has
    // been deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutAddrMasks interface{}

    // The number of ICMP Address Mask Reply messages sent.  This object has been
    // deprecated, as a new IP version-neutral table has been added.  It is
    // loosely replaced by a column in the icmpMsgStatsTable. The type is
    // interface{} with range: 0..4294967295.
    IcmpOutAddrMaskReps interface{}
}

func (icmp *IPMIB_Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xe"
    icmp.EntityData.ParentYangName = "IP-MIB"
    icmp.EntityData.SegmentPath = "icmp"
    icmp.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + icmp.EntityData.SegmentPath
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

// IPMIB_IpAddrTable
// The table of addressing information relevant to this
// entity's IPv4 addresses.
// 
// This table has been deprecated, as a new IP version-neutral
// table has been added.  It is loosely replaced by the
// ipAddressTable although several objects that weren't deemed
// useful weren't carried forward while another
// (ipAdEntReasmMaxSize) was moved to the ipv4InterfaceTable.
type IPMIB_IpAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The addressing information for one of this entity's IPv4 addresses. The
    // type is slice of IPMIB_IpAddrTable_IpAddrEntry.
    IpAddrEntry []*IPMIB_IpAddrTable_IpAddrEntry
}

func (ipAddrTable *IPMIB_IpAddrTable) GetEntityData() *types.CommonEntityData {
    ipAddrTable.EntityData.YFilter = ipAddrTable.YFilter
    ipAddrTable.EntityData.YangName = "ipAddrTable"
    ipAddrTable.EntityData.BundleName = "cisco_ios_xe"
    ipAddrTable.EntityData.ParentYangName = "IP-MIB"
    ipAddrTable.EntityData.SegmentPath = "ipAddrTable"
    ipAddrTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipAddrTable.EntityData.SegmentPath
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

// IPMIB_IpAddrTable_IpAddrEntry
// The addressing information for one of this entity's IPv4
// addresses.
type IPMIB_IpAddrTable_IpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IPv4 address to which this entry's addressing
    // information pertains. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAdEntAddr interface{}

    // The index value which uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of the IF-MIB's
    // ifIndex. The type is interface{} with range: 1..2147483647.
    IpAdEntIfIndex interface{}

    // The subnet mask associated with the IPv4 address of this entry.  The value
    // of the mask is an IPv4 address with all the network bits set to 1 and all
    // the hosts bits set to 0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAdEntNetMask interface{}

    // The value of the least-significant bit in the IPv4 broadcast address used
    // for sending datagrams on the (logical) interface associated with the IPv4
    // address of this entry. For example, when the Internet standard all-ones
    // broadcast address is used, the value will be 1.  This value applies to both
    // the subnet and network broadcast addresses used by the entity on this
    // (logical) interface. The type is interface{} with range: 0..1.
    IpAdEntBcastAddr interface{}

    // The size of the largest IPv4 datagram which this entity can re-assemble
    // from incoming IPv4 fragmented datagrams received on this interface. The
    // type is interface{} with range: 0..65535.
    IpAdEntReasmMaxSize interface{}
}

func (ipAddrEntry *IPMIB_IpAddrTable_IpAddrEntry) GetEntityData() *types.CommonEntityData {
    ipAddrEntry.EntityData.YFilter = ipAddrEntry.YFilter
    ipAddrEntry.EntityData.YangName = "ipAddrEntry"
    ipAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    ipAddrEntry.EntityData.ParentYangName = "ipAddrTable"
    ipAddrEntry.EntityData.SegmentPath = "ipAddrEntry" + types.AddKeyToken(ipAddrEntry.IpAdEntAddr, "ipAdEntAddr")
    ipAddrEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipAddrTable/" + ipAddrEntry.EntityData.SegmentPath
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

// IPMIB_IpNetToMediaTable
// The IPv4 Address Translation table used for mapping from
// IPv4 addresses to physical addresses.
// 
// This table has been deprecated, as a new IP version-neutral
// table has been added.  It is loosely replaced by the
// ipNetToPhysicalTable.
type IPMIB_IpNetToMediaTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one IpAddress to `physical' address equivalence. The
    // type is slice of IPMIB_IpNetToMediaTable_IpNetToMediaEntry.
    IpNetToMediaEntry []*IPMIB_IpNetToMediaTable_IpNetToMediaEntry
}

func (ipNetToMediaTable *IPMIB_IpNetToMediaTable) GetEntityData() *types.CommonEntityData {
    ipNetToMediaTable.EntityData.YFilter = ipNetToMediaTable.YFilter
    ipNetToMediaTable.EntityData.YangName = "ipNetToMediaTable"
    ipNetToMediaTable.EntityData.BundleName = "cisco_ios_xe"
    ipNetToMediaTable.EntityData.ParentYangName = "IP-MIB"
    ipNetToMediaTable.EntityData.SegmentPath = "ipNetToMediaTable"
    ipNetToMediaTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipNetToMediaTable.EntityData.SegmentPath
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

// IPMIB_IpNetToMediaTable_IpNetToMediaEntry
// Each entry contains one IpAddress to `physical' address
// equivalence.
type IPMIB_IpNetToMediaTable_IpNetToMediaEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The interface on which this entry's equivalence is
    // effective.  The interface identified by a particular value of this index is
    // the same interface as identified by the   same value of the IF-MIB's
    // ifIndex.  This object predates the rule limiting index objects to a max
    // access value of 'not-accessible' and so continues to use a value of
    // 'read-create'. The type is interface{} with range: 1..2147483647.
    IpNetToMediaIfIndex interface{}

    // This attribute is a key. The IpAddress corresponding to the media-dependent
    // `physical' address.  This object predates the rule limiting index objects
    // to a max access value of 'not-accessible' and so continues to use a value
    // of 'read-create'. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpNetToMediaNetAddress interface{}

    // The media-dependent `physical' address.  This object should return 0 when
    // this entry is in the 'incomplete' state.  As the entries in this table are
    // typically not persistent when this object is written the entity should not
    // save the change to non-volatile storage.  Note: a stronger requirement is
    // not used because this object was previously defined. The type is string
    // with length: 0..65535.
    IpNetToMediaPhysAddress interface{}

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
    // defined. The type is IpNetToMediaType.
    IpNetToMediaType interface{}
}

func (ipNetToMediaEntry *IPMIB_IpNetToMediaTable_IpNetToMediaEntry) GetEntityData() *types.CommonEntityData {
    ipNetToMediaEntry.EntityData.YFilter = ipNetToMediaEntry.YFilter
    ipNetToMediaEntry.EntityData.YangName = "ipNetToMediaEntry"
    ipNetToMediaEntry.EntityData.BundleName = "cisco_ios_xe"
    ipNetToMediaEntry.EntityData.ParentYangName = "ipNetToMediaTable"
    ipNetToMediaEntry.EntityData.SegmentPath = "ipNetToMediaEntry" + types.AddKeyToken(ipNetToMediaEntry.IpNetToMediaIfIndex, "ipNetToMediaIfIndex") + types.AddKeyToken(ipNetToMediaEntry.IpNetToMediaNetAddress, "ipNetToMediaNetAddress")
    ipNetToMediaEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipNetToMediaTable/" + ipNetToMediaEntry.EntityData.SegmentPath
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

// IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType represents defined.
type IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType string

const (
    IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_other IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "other"

    IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_invalid IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "invalid"

    IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_dynamic IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "dynamic"

    IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType_static IPMIB_IpNetToMediaTable_IpNetToMediaEntry_IpNetToMediaType = "static"
)

// IPMIB_Ipv4InterfaceTable
// The table containing per-interface IPv4-specific
// information.
type IPMIB_Ipv4InterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing IPv4-specific information for a specific interface. The
    // type is slice of IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry.
    Ipv4InterfaceEntry []*IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry
}

func (ipv4InterfaceTable *IPMIB_Ipv4InterfaceTable) GetEntityData() *types.CommonEntityData {
    ipv4InterfaceTable.EntityData.YFilter = ipv4InterfaceTable.YFilter
    ipv4InterfaceTable.EntityData.YangName = "ipv4InterfaceTable"
    ipv4InterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    ipv4InterfaceTable.EntityData.ParentYangName = "IP-MIB"
    ipv4InterfaceTable.EntityData.SegmentPath = "ipv4InterfaceTable"
    ipv4InterfaceTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipv4InterfaceTable.EntityData.SegmentPath
    ipv4InterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4InterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4InterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4InterfaceTable.EntityData.Children = types.NewOrderedMap()
    ipv4InterfaceTable.EntityData.Children.Append("ipv4InterfaceEntry", types.YChild{"Ipv4InterfaceEntry", nil})
    for i := range ipv4InterfaceTable.Ipv4InterfaceEntry {
        ipv4InterfaceTable.EntityData.Children.Append(types.GetSegmentPath(ipv4InterfaceTable.Ipv4InterfaceEntry[i]), types.YChild{"Ipv4InterfaceEntry", ipv4InterfaceTable.Ipv4InterfaceEntry[i]})
    }
    ipv4InterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    ipv4InterfaceTable.EntityData.YListKeys = []string {}

    return &(ipv4InterfaceTable.EntityData)
}

// IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry
// An entry containing IPv4-specific information for a specific
// interface.
type IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipv4InterfaceIfIndex interface{}

    // The size of the largest IPv4 datagram that this entity can re-assemble from
    // incoming IPv4 fragmented datagrams received on this interface. The type is
    // interface{} with range: 0..65535.
    Ipv4InterfaceReasmMaxSize interface{}

    // The indication of whether IPv4 is enabled (up) or disabled (down) on this
    // interface.  This object does not affect the state of the interface itself,
    // only its connection to an IPv4 stack.  The IF-MIB should be used to control
    // the state of the interface. The type is Ipv4InterfaceEnableStatus.
    Ipv4InterfaceEnableStatus interface{}

    // The time between retransmissions of ARP requests to a neighbor when
    // resolving the address or when probing the reachability of a neighbor. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Ipv4InterfaceRetransmitTime interface{}
}

func (ipv4InterfaceEntry *IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry) GetEntityData() *types.CommonEntityData {
    ipv4InterfaceEntry.EntityData.YFilter = ipv4InterfaceEntry.YFilter
    ipv4InterfaceEntry.EntityData.YangName = "ipv4InterfaceEntry"
    ipv4InterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    ipv4InterfaceEntry.EntityData.ParentYangName = "ipv4InterfaceTable"
    ipv4InterfaceEntry.EntityData.SegmentPath = "ipv4InterfaceEntry" + types.AddKeyToken(ipv4InterfaceEntry.Ipv4InterfaceIfIndex, "ipv4InterfaceIfIndex")
    ipv4InterfaceEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipv4InterfaceTable/" + ipv4InterfaceEntry.EntityData.SegmentPath
    ipv4InterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4InterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4InterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4InterfaceEntry.EntityData.Children = types.NewOrderedMap()
    ipv4InterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    ipv4InterfaceEntry.EntityData.Leafs.Append("ipv4InterfaceIfIndex", types.YLeaf{"Ipv4InterfaceIfIndex", ipv4InterfaceEntry.Ipv4InterfaceIfIndex})
    ipv4InterfaceEntry.EntityData.Leafs.Append("ipv4InterfaceReasmMaxSize", types.YLeaf{"Ipv4InterfaceReasmMaxSize", ipv4InterfaceEntry.Ipv4InterfaceReasmMaxSize})
    ipv4InterfaceEntry.EntityData.Leafs.Append("ipv4InterfaceEnableStatus", types.YLeaf{"Ipv4InterfaceEnableStatus", ipv4InterfaceEntry.Ipv4InterfaceEnableStatus})
    ipv4InterfaceEntry.EntityData.Leafs.Append("ipv4InterfaceRetransmitTime", types.YLeaf{"Ipv4InterfaceRetransmitTime", ipv4InterfaceEntry.Ipv4InterfaceRetransmitTime})

    ipv4InterfaceEntry.EntityData.YListKeys = []string {"Ipv4InterfaceIfIndex"}

    return &(ipv4InterfaceEntry.EntityData)
}

// IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry_Ipv4InterfaceEnableStatus represents of the interface.
type IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry_Ipv4InterfaceEnableStatus string

const (
    IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry_Ipv4InterfaceEnableStatus_up IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry_Ipv4InterfaceEnableStatus = "up"

    IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry_Ipv4InterfaceEnableStatus_down IPMIB_Ipv4InterfaceTable_Ipv4InterfaceEntry_Ipv4InterfaceEnableStatus = "down"
)

// IPMIB_Ipv6InterfaceTable
// The table containing per-interface IPv6-specific
// information.
type IPMIB_Ipv6InterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing IPv6-specific information for a given interface. The
    // type is slice of IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry.
    Ipv6InterfaceEntry []*IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry
}

func (ipv6InterfaceTable *IPMIB_Ipv6InterfaceTable) GetEntityData() *types.CommonEntityData {
    ipv6InterfaceTable.EntityData.YFilter = ipv6InterfaceTable.YFilter
    ipv6InterfaceTable.EntityData.YangName = "ipv6InterfaceTable"
    ipv6InterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    ipv6InterfaceTable.EntityData.ParentYangName = "IP-MIB"
    ipv6InterfaceTable.EntityData.SegmentPath = "ipv6InterfaceTable"
    ipv6InterfaceTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipv6InterfaceTable.EntityData.SegmentPath
    ipv6InterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6InterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6InterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6InterfaceTable.EntityData.Children = types.NewOrderedMap()
    ipv6InterfaceTable.EntityData.Children.Append("ipv6InterfaceEntry", types.YChild{"Ipv6InterfaceEntry", nil})
    for i := range ipv6InterfaceTable.Ipv6InterfaceEntry {
        ipv6InterfaceTable.EntityData.Children.Append(types.GetSegmentPath(ipv6InterfaceTable.Ipv6InterfaceEntry[i]), types.YChild{"Ipv6InterfaceEntry", ipv6InterfaceTable.Ipv6InterfaceEntry[i]})
    }
    ipv6InterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    ipv6InterfaceTable.EntityData.YListKeys = []string {}

    return &(ipv6InterfaceTable.EntityData)
}

// IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry
// An entry containing IPv6-specific information for a given
// interface.
type IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipv6InterfaceIfIndex interface{}

    // The size of the largest IPv6 datagram that this entity can re-assemble from
    // incoming IPv6 fragmented datagrams received on this interface. The type is
    // interface{} with range: 1500..65535. Units are octets.
    Ipv6InterfaceReasmMaxSize interface{}

    // The Interface Identifier for this interface.  The Interface Identifier is
    // combined with an address prefix to form an interface address.  By default,
    // the Interface Identifier is auto-configured according to the rules of the
    // link type to which this interface is attached.   A zero length identifier
    // may be used where appropriate.  One possible example is a loopback
    // interface. The type is string.
    Ipv6InterfaceIdentifier interface{}

    // The indication of whether IPv6 is enabled (up) or disabled (down) on this
    // interface.  This object does not affect the state of the interface itself,
    // only its connection to an IPv6 stack.  The IF-MIB should be used to control
    // the state of the interface.  When this object is written, the entity SHOULD
    // save the change to non-volatile storage and restore the object from
    // non-volatile storage upon re-initialization of the system. The type is
    // Ipv6InterfaceEnableStatus.
    Ipv6InterfaceEnableStatus interface{}

    // The time a neighbor is considered reachable after receiving a reachability
    // confirmation. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Ipv6InterfaceReachableTime interface{}

    // The time between retransmissions of Neighbor Solicitation messages to a
    // neighbor when resolving the address or when probing the reachability of a
    // neighbor. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Ipv6InterfaceRetransmitTime interface{}

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
    // upon re-initialization of the system. The type is Ipv6InterfaceForwarding.
    Ipv6InterfaceForwarding interface{}
}

func (ipv6InterfaceEntry *IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry) GetEntityData() *types.CommonEntityData {
    ipv6InterfaceEntry.EntityData.YFilter = ipv6InterfaceEntry.YFilter
    ipv6InterfaceEntry.EntityData.YangName = "ipv6InterfaceEntry"
    ipv6InterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    ipv6InterfaceEntry.EntityData.ParentYangName = "ipv6InterfaceTable"
    ipv6InterfaceEntry.EntityData.SegmentPath = "ipv6InterfaceEntry" + types.AddKeyToken(ipv6InterfaceEntry.Ipv6InterfaceIfIndex, "ipv6InterfaceIfIndex")
    ipv6InterfaceEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipv6InterfaceTable/" + ipv6InterfaceEntry.EntityData.SegmentPath
    ipv6InterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6InterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6InterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6InterfaceEntry.EntityData.Children = types.NewOrderedMap()
    ipv6InterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceIfIndex", types.YLeaf{"Ipv6InterfaceIfIndex", ipv6InterfaceEntry.Ipv6InterfaceIfIndex})
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceReasmMaxSize", types.YLeaf{"Ipv6InterfaceReasmMaxSize", ipv6InterfaceEntry.Ipv6InterfaceReasmMaxSize})
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceIdentifier", types.YLeaf{"Ipv6InterfaceIdentifier", ipv6InterfaceEntry.Ipv6InterfaceIdentifier})
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceEnableStatus", types.YLeaf{"Ipv6InterfaceEnableStatus", ipv6InterfaceEntry.Ipv6InterfaceEnableStatus})
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceReachableTime", types.YLeaf{"Ipv6InterfaceReachableTime", ipv6InterfaceEntry.Ipv6InterfaceReachableTime})
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceRetransmitTime", types.YLeaf{"Ipv6InterfaceRetransmitTime", ipv6InterfaceEntry.Ipv6InterfaceRetransmitTime})
    ipv6InterfaceEntry.EntityData.Leafs.Append("ipv6InterfaceForwarding", types.YLeaf{"Ipv6InterfaceForwarding", ipv6InterfaceEntry.Ipv6InterfaceForwarding})

    ipv6InterfaceEntry.EntityData.YListKeys = []string {"Ipv6InterfaceIfIndex"}

    return &(ipv6InterfaceEntry.EntityData)
}

// IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceEnableStatus represents non-volatile storage upon re-initialization of the system.
type IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceEnableStatus string

const (
    IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceEnableStatus_up IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceEnableStatus = "up"

    IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceEnableStatus_down IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceEnableStatus = "down"
)

// IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceForwarding represents non-volatile storage upon re-initialization of the system.
type IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceForwarding string

const (
    IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceForwarding_forwarding IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceForwarding = "forwarding"

    IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceForwarding_notForwarding IPMIB_Ipv6InterfaceTable_Ipv6InterfaceEntry_Ipv6InterfaceForwarding = "notForwarding"
)

// IPMIB_IpSystemStatsTable
// The table containing system wide, IP version specific
// traffic statistics.  This table and the ipIfStatsTable
// contain similar objects whose difference is in their
// granularity.  Where this table contains system wide traffic
// statistics, the ipIfStatsTable contains the same statistics
// but counted on a per-interface basis.
type IPMIB_IpSystemStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A statistics entry containing system-wide objects for a particular IP
    // version. The type is slice of IPMIB_IpSystemStatsTable_IpSystemStatsEntry.
    IpSystemStatsEntry []*IPMIB_IpSystemStatsTable_IpSystemStatsEntry
}

func (ipSystemStatsTable *IPMIB_IpSystemStatsTable) GetEntityData() *types.CommonEntityData {
    ipSystemStatsTable.EntityData.YFilter = ipSystemStatsTable.YFilter
    ipSystemStatsTable.EntityData.YangName = "ipSystemStatsTable"
    ipSystemStatsTable.EntityData.BundleName = "cisco_ios_xe"
    ipSystemStatsTable.EntityData.ParentYangName = "IP-MIB"
    ipSystemStatsTable.EntityData.SegmentPath = "ipSystemStatsTable"
    ipSystemStatsTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipSystemStatsTable.EntityData.SegmentPath
    ipSystemStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSystemStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSystemStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSystemStatsTable.EntityData.Children = types.NewOrderedMap()
    ipSystemStatsTable.EntityData.Children.Append("ipSystemStatsEntry", types.YChild{"IpSystemStatsEntry", nil})
    for i := range ipSystemStatsTable.IpSystemStatsEntry {
        ipSystemStatsTable.EntityData.Children.Append(types.GetSegmentPath(ipSystemStatsTable.IpSystemStatsEntry[i]), types.YChild{"IpSystemStatsEntry", ipSystemStatsTable.IpSystemStatsEntry[i]})
    }
    ipSystemStatsTable.EntityData.Leafs = types.NewOrderedMap()

    ipSystemStatsTable.EntityData.YListKeys = []string {}

    return &(ipSystemStatsTable.EntityData)
}

// IPMIB_IpSystemStatsTable_IpSystemStatsEntry
// A statistics entry containing system-wide objects for a
// particular IP version.
type IPMIB_IpSystemStatsTable_IpSystemStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP version of this row. The type is IpVersion.
    IpSystemStatsIPVersion interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    IpSystemStatsInReceives interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  This object counts the same datagrams as
    // ipSystemStatsInReceives, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCInReceives interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  Octets from datagrams counted in
    // ipSystemStatsInReceives MUST be counted here.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsInOctets interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  This object counts the same octets as
    // ipSystemStatsInOctets, but allows for larger   values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCInOctets interface{}

    // The number of input IP datagrams discarded due to errors in their IP
    // headers, including version number mismatch, other format errors, hop count
    // exceeded, errors discovered in processing their IP options, etc. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsInHdrErrors interface{}

    // The number of input IP datagrams discarded because no route could be found
    // to transmit them to their destination.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IpSystemStatsInNoRoutes interface{}

    // The number of input IP datagrams discarded because the IP address in their
    // IP header's destination field was not a valid address to be received at
    // this entity.  This count includes invalid addresses (e.g., ::0).  For
    // entities that are not IP routers and therefore do not forward   datagrams,
    // this counter includes datagrams discarded because the destination address
    // was not a local address.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpSystemStatsInAddrErrors interface{}

    // The number of locally-addressed IP datagrams received successfully but
    // discarded because of an unknown or unsupported protocol.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IpSystemStatsInUnknownProtos interface{}

    // The number of input IP datagrams discarded because the datagram frame
    // didn't carry enough data.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpSystemStatsInTruncatedPkts interface{}

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
    IpSystemStatsInForwDatagrams interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination and for which this entity attempted to find a route to forward
    // them to that final destination.  This object counts the same packets as
    // ipSystemStatsInForwDatagrams, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCInForwDatagrams interface{}

    // The number of IP fragments received that needed to be reassembled at this
    // interface.  When tracking interface statistics, the counter of the
    // interface to which these fragments were addressed is incremented.  This
    // interface might not be the same as the input interface for some of the
    // fragments.  Discontinuities in the value of this counter can occur at  
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    IpSystemStatsReasmReqds interface{}

    // The number of IP datagrams successfully reassembled.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IpSystemStatsReasmOKs interface{}

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
    IpSystemStatsReasmFails interface{}

    // The number of input IP datagrams for which no problems were encountered to
    // prevent their continued processing, but were discarded (e.g., for lack of
    // buffer space).  Note that this counter does not include any datagrams
    // discarded while awaiting re-assembly.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IpSystemStatsInDiscards interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  When tracking interface statistics, the counter of the
    // interface to which these datagrams were addressed is incremented.  This
    // interface might not be the same as the input interface for some of the
    // datagrams.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    IpSystemStatsInDelivers interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  This object counts the same packets as
    // ipSystemStatsInDelivers, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCInDelivers interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  Note that this counter
    // does not include any datagrams counted in ipSystemStatsOutForwDatagrams. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutRequests interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  This object counts the
    // same packets as ipSystemStatsOutRequests, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCOutRequests interface{}

    // The number of locally generated IP datagrams discarded because no route
    // could be found to transmit them to their destination.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutNoRoutes interface{}

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
    IpSystemStatsOutForwDatagrams interface{}

    // The number of datagrams for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.  This object counts the same packets as
    // ipSystemStatsOutForwDatagrams, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCOutForwDatagrams interface{}

    // The number of output IP datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but were discarded (e.g.,
    // for lack of buffer space).  Note that this counter would include  
    // datagrams counted in ipSystemStatsOutForwDatagrams if any such datagrams
    // met this (discretionary) discard criterion.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutDiscards interface{}

    // The number of IP datagrams that would require fragmentation in order to be
    // transmitted.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for a successfully fragmented datagram. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutFragReqds interface{}

    // The number of IP datagrams that have been successfully fragmented.  When
    // tracking interface statistics, the counter of the outgoing interface is
    // incremented for a successfully fragmented datagram.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutFragOKs interface{}

    // The number of IP datagrams that have been discarded because they needed to
    // be fragmented but could not be.  This includes IPv4 packets that have the
    // DF bit set and IPv6 packets that are being forwarded and exceed the
    // outgoing link MTU.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for an unsuccessfully fragmented
    // datagram.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    IpSystemStatsOutFragFails interface{}

    // The number of output datagram fragments that have been generated as a
    // result of IP fragmentation.  When tracking interface statistics, the
    // counter of the outgoing interface is incremented for a successfully
    // fragmented datagram.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpSystemStatsOutFragCreates interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This includes datagrams generated locally and
    // those forwarded by this entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other   times as indicated by the value of ipSystemStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IpSystemStatsOutTransmits interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This object counts the same datagrams as
    // ipSystemStatsOutTransmits, but allows for larger values.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCOutTransmits interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  Octets from datagrams counted in
    // ipSystemStatsOutTransmits MUST be counted here.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutOctets interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  This objects counts the same octets as
    // ipSystemStatsOutOctets, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of  
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCOutOctets interface{}

    // The number of IP multicast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsInMcastPkts interface{}

    // The number of IP multicast datagrams received.  This object counts the same
    // datagrams as ipSystemStatsInMcastPkts but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCInMcastPkts interface{}

    // The total number of octets received in IP multicast datagrams.  Octets from
    // datagrams counted in ipSystemStatsInMcastPkts MUST be counted here. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsInMcastOctets interface{}

    // The total number of octets received in IP multicast datagrams.  This object
    // counts the same octets as ipSystemStatsInMcastOctets, but allows for larger
    // values.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipSystemStatsDiscontinuityTime. The type is interface{}
    // with range: 0..18446744073709551615.
    IpSystemStatsHCInMcastOctets interface{}

    // The number of IP multicast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutMcastPkts interface{}

    // The number of IP multicast datagrams transmitted.  This object counts the
    // same datagrams as ipSystemStatsOutMcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCOutMcastPkts interface{}

    // The total number of octets transmitted in IP multicast datagrams.  Octets
    // from datagrams counted in   ipSystemStatsOutMcastPkts MUST be counted here.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutMcastOctets interface{}

    // The total number of octets transmitted in IP multicast datagrams.  This
    // object counts the same octets as ipSystemStatsOutMcastOctets, but allows
    // for larger values.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ipSystemStatsDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    IpSystemStatsHCOutMcastOctets interface{}

    // The number of IP broadcast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsInBcastPkts interface{}

    // The number of IP broadcast datagrams received.  This object counts the same
    // datagrams as ipSystemStatsInBcastPkts but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of  
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCInBcastPkts interface{}

    // The number of IP broadcast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpSystemStatsOutBcastPkts interface{}

    // The number of IP broadcast datagrams transmitted.  This object counts the
    // same datagrams as ipSystemStatsOutBcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipSystemStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpSystemStatsHCOutBcastPkts interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    IpSystemStatsDiscontinuityTime interface{}

    // The minimum reasonable polling interval for this entry. This object
    // provides an indication of the minimum amount of time required to update the
    // counters in this entry. The type is interface{} with range: 0..4294967295.
    // Units are milli-seconds.
    IpSystemStatsRefreshRate interface{}
}

func (ipSystemStatsEntry *IPMIB_IpSystemStatsTable_IpSystemStatsEntry) GetEntityData() *types.CommonEntityData {
    ipSystemStatsEntry.EntityData.YFilter = ipSystemStatsEntry.YFilter
    ipSystemStatsEntry.EntityData.YangName = "ipSystemStatsEntry"
    ipSystemStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    ipSystemStatsEntry.EntityData.ParentYangName = "ipSystemStatsTable"
    ipSystemStatsEntry.EntityData.SegmentPath = "ipSystemStatsEntry" + types.AddKeyToken(ipSystemStatsEntry.IpSystemStatsIPVersion, "ipSystemStatsIPVersion")
    ipSystemStatsEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipSystemStatsTable/" + ipSystemStatsEntry.EntityData.SegmentPath
    ipSystemStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSystemStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSystemStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSystemStatsEntry.EntityData.Children = types.NewOrderedMap()
    ipSystemStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsIPVersion", types.YLeaf{"IpSystemStatsIPVersion", ipSystemStatsEntry.IpSystemStatsIPVersion})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInReceives", types.YLeaf{"IpSystemStatsInReceives", ipSystemStatsEntry.IpSystemStatsInReceives})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInReceives", types.YLeaf{"IpSystemStatsHCInReceives", ipSystemStatsEntry.IpSystemStatsHCInReceives})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInOctets", types.YLeaf{"IpSystemStatsInOctets", ipSystemStatsEntry.IpSystemStatsInOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInOctets", types.YLeaf{"IpSystemStatsHCInOctets", ipSystemStatsEntry.IpSystemStatsHCInOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInHdrErrors", types.YLeaf{"IpSystemStatsInHdrErrors", ipSystemStatsEntry.IpSystemStatsInHdrErrors})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInNoRoutes", types.YLeaf{"IpSystemStatsInNoRoutes", ipSystemStatsEntry.IpSystemStatsInNoRoutes})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInAddrErrors", types.YLeaf{"IpSystemStatsInAddrErrors", ipSystemStatsEntry.IpSystemStatsInAddrErrors})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInUnknownProtos", types.YLeaf{"IpSystemStatsInUnknownProtos", ipSystemStatsEntry.IpSystemStatsInUnknownProtos})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInTruncatedPkts", types.YLeaf{"IpSystemStatsInTruncatedPkts", ipSystemStatsEntry.IpSystemStatsInTruncatedPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInForwDatagrams", types.YLeaf{"IpSystemStatsInForwDatagrams", ipSystemStatsEntry.IpSystemStatsInForwDatagrams})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInForwDatagrams", types.YLeaf{"IpSystemStatsHCInForwDatagrams", ipSystemStatsEntry.IpSystemStatsHCInForwDatagrams})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsReasmReqds", types.YLeaf{"IpSystemStatsReasmReqds", ipSystemStatsEntry.IpSystemStatsReasmReqds})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsReasmOKs", types.YLeaf{"IpSystemStatsReasmOKs", ipSystemStatsEntry.IpSystemStatsReasmOKs})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsReasmFails", types.YLeaf{"IpSystemStatsReasmFails", ipSystemStatsEntry.IpSystemStatsReasmFails})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInDiscards", types.YLeaf{"IpSystemStatsInDiscards", ipSystemStatsEntry.IpSystemStatsInDiscards})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInDelivers", types.YLeaf{"IpSystemStatsInDelivers", ipSystemStatsEntry.IpSystemStatsInDelivers})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInDelivers", types.YLeaf{"IpSystemStatsHCInDelivers", ipSystemStatsEntry.IpSystemStatsHCInDelivers})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutRequests", types.YLeaf{"IpSystemStatsOutRequests", ipSystemStatsEntry.IpSystemStatsOutRequests})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutRequests", types.YLeaf{"IpSystemStatsHCOutRequests", ipSystemStatsEntry.IpSystemStatsHCOutRequests})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutNoRoutes", types.YLeaf{"IpSystemStatsOutNoRoutes", ipSystemStatsEntry.IpSystemStatsOutNoRoutes})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutForwDatagrams", types.YLeaf{"IpSystemStatsOutForwDatagrams", ipSystemStatsEntry.IpSystemStatsOutForwDatagrams})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutForwDatagrams", types.YLeaf{"IpSystemStatsHCOutForwDatagrams", ipSystemStatsEntry.IpSystemStatsHCOutForwDatagrams})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutDiscards", types.YLeaf{"IpSystemStatsOutDiscards", ipSystemStatsEntry.IpSystemStatsOutDiscards})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutFragReqds", types.YLeaf{"IpSystemStatsOutFragReqds", ipSystemStatsEntry.IpSystemStatsOutFragReqds})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutFragOKs", types.YLeaf{"IpSystemStatsOutFragOKs", ipSystemStatsEntry.IpSystemStatsOutFragOKs})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutFragFails", types.YLeaf{"IpSystemStatsOutFragFails", ipSystemStatsEntry.IpSystemStatsOutFragFails})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutFragCreates", types.YLeaf{"IpSystemStatsOutFragCreates", ipSystemStatsEntry.IpSystemStatsOutFragCreates})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutTransmits", types.YLeaf{"IpSystemStatsOutTransmits", ipSystemStatsEntry.IpSystemStatsOutTransmits})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutTransmits", types.YLeaf{"IpSystemStatsHCOutTransmits", ipSystemStatsEntry.IpSystemStatsHCOutTransmits})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutOctets", types.YLeaf{"IpSystemStatsOutOctets", ipSystemStatsEntry.IpSystemStatsOutOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutOctets", types.YLeaf{"IpSystemStatsHCOutOctets", ipSystemStatsEntry.IpSystemStatsHCOutOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInMcastPkts", types.YLeaf{"IpSystemStatsInMcastPkts", ipSystemStatsEntry.IpSystemStatsInMcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInMcastPkts", types.YLeaf{"IpSystemStatsHCInMcastPkts", ipSystemStatsEntry.IpSystemStatsHCInMcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInMcastOctets", types.YLeaf{"IpSystemStatsInMcastOctets", ipSystemStatsEntry.IpSystemStatsInMcastOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInMcastOctets", types.YLeaf{"IpSystemStatsHCInMcastOctets", ipSystemStatsEntry.IpSystemStatsHCInMcastOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutMcastPkts", types.YLeaf{"IpSystemStatsOutMcastPkts", ipSystemStatsEntry.IpSystemStatsOutMcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutMcastPkts", types.YLeaf{"IpSystemStatsHCOutMcastPkts", ipSystemStatsEntry.IpSystemStatsHCOutMcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutMcastOctets", types.YLeaf{"IpSystemStatsOutMcastOctets", ipSystemStatsEntry.IpSystemStatsOutMcastOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutMcastOctets", types.YLeaf{"IpSystemStatsHCOutMcastOctets", ipSystemStatsEntry.IpSystemStatsHCOutMcastOctets})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsInBcastPkts", types.YLeaf{"IpSystemStatsInBcastPkts", ipSystemStatsEntry.IpSystemStatsInBcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCInBcastPkts", types.YLeaf{"IpSystemStatsHCInBcastPkts", ipSystemStatsEntry.IpSystemStatsHCInBcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsOutBcastPkts", types.YLeaf{"IpSystemStatsOutBcastPkts", ipSystemStatsEntry.IpSystemStatsOutBcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsHCOutBcastPkts", types.YLeaf{"IpSystemStatsHCOutBcastPkts", ipSystemStatsEntry.IpSystemStatsHCOutBcastPkts})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsDiscontinuityTime", types.YLeaf{"IpSystemStatsDiscontinuityTime", ipSystemStatsEntry.IpSystemStatsDiscontinuityTime})
    ipSystemStatsEntry.EntityData.Leafs.Append("ipSystemStatsRefreshRate", types.YLeaf{"IpSystemStatsRefreshRate", ipSystemStatsEntry.IpSystemStatsRefreshRate})

    ipSystemStatsEntry.EntityData.YListKeys = []string {"IpSystemStatsIPVersion"}

    return &(ipSystemStatsEntry.EntityData)
}

// IPMIB_IpIfStatsTable
// The table containing per-interface traffic statistics.  This
// table and the ipSystemStatsTable contain similar objects
// whose difference is in their granularity.  Where this table
// contains per-interface statistics, the ipSystemStatsTable
// contains the same statistics, but counted on a system wide
// basis.
type IPMIB_IpIfStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An interface statistics entry containing objects for a particular interface
    // and version of IP. The type is slice of
    // IPMIB_IpIfStatsTable_IpIfStatsEntry.
    IpIfStatsEntry []*IPMIB_IpIfStatsTable_IpIfStatsEntry
}

func (ipIfStatsTable *IPMIB_IpIfStatsTable) GetEntityData() *types.CommonEntityData {
    ipIfStatsTable.EntityData.YFilter = ipIfStatsTable.YFilter
    ipIfStatsTable.EntityData.YangName = "ipIfStatsTable"
    ipIfStatsTable.EntityData.BundleName = "cisco_ios_xe"
    ipIfStatsTable.EntityData.ParentYangName = "IP-MIB"
    ipIfStatsTable.EntityData.SegmentPath = "ipIfStatsTable"
    ipIfStatsTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipIfStatsTable.EntityData.SegmentPath
    ipIfStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipIfStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipIfStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipIfStatsTable.EntityData.Children = types.NewOrderedMap()
    ipIfStatsTable.EntityData.Children.Append("ipIfStatsEntry", types.YChild{"IpIfStatsEntry", nil})
    for i := range ipIfStatsTable.IpIfStatsEntry {
        ipIfStatsTable.EntityData.Children.Append(types.GetSegmentPath(ipIfStatsTable.IpIfStatsEntry[i]), types.YChild{"IpIfStatsEntry", ipIfStatsTable.IpIfStatsEntry[i]})
    }
    ipIfStatsTable.EntityData.Leafs = types.NewOrderedMap()

    ipIfStatsTable.EntityData.YListKeys = []string {}

    return &(ipIfStatsTable.EntityData)
}

// IPMIB_IpIfStatsTable_IpIfStatsEntry
// An interface statistics entry containing objects for a
// particular interface and version of IP.
type IPMIB_IpIfStatsTable_IpIfStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP version of this row. The type is IpVersion.
    IpIfStatsIPVersion interface{}

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    IpIfStatsIfIndex interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IpIfStatsInReceives interface{}

    // The total number of input IP datagrams received, including those received
    // in error.  This object counts the same datagrams as ipIfStatsInReceives,
    // but allows for larger values.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615.
    IpIfStatsHCInReceives interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  Octets from datagrams counted in ipIfStatsInReceives
    // MUST be counted here.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpIfStatsInOctets interface{}

    // The total number of octets received in input IP datagrams, including those
    // received in error.  This object counts the same octets as
    // ipIfStatsInOctets, but allows for larger values.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCInOctets interface{}

    // The number of input IP datagrams discarded due to errors in their IP
    // headers, including version number mismatch, other format errors, hop count
    // exceeded, errors discovered in processing their IP options, etc. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsInHdrErrors interface{}

    // The number of input IP datagrams discarded because no route could be found
    // to transmit them to their destination.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    IpIfStatsInNoRoutes interface{}

    // The number of input IP datagrams discarded because the IP address in their
    // IP header's destination field was not a valid address to be received at
    // this entity.  This count includes invalid addresses (e.g., ::0).  For
    // entities that are not IP routers and therefore do not forward datagrams,
    // this counter includes datagrams discarded because the destination address
    // was not a local address.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpIfStatsInAddrErrors interface{}

    // The number of locally-addressed IP datagrams received successfully but
    // discarded because of an unknown or unsupported protocol.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of   ipIfStatsDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IpIfStatsInUnknownProtos interface{}

    // The number of input IP datagrams discarded because the datagram frame
    // didn't carry enough data.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpIfStatsInTruncatedPkts interface{}

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
    IpIfStatsInForwDatagrams interface{}

    // The number of input datagrams for which this entity was not their final IP
    // destination and for which this entity attempted to find a route to forward
    // them to that final destination.  This object counts the same packets as  
    // ipIfStatsInForwDatagrams, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCInForwDatagrams interface{}

    // The number of IP fragments received that needed to be reassembled at this
    // interface.  When tracking interface statistics, the counter of the
    // interface to which these fragments were addressed is incremented.  This
    // interface might not be the same as the input interface for some of the
    // fragments.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IpIfStatsReasmReqds interface{}

    // The number of IP datagrams successfully reassembled.  When tracking
    // interface statistics, the counter of the interface to which these datagrams
    // were addressed is incremented.  This interface might not be the same as the
    // input interface for some of the datagrams.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    IpIfStatsReasmOKs interface{}

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
    IpIfStatsReasmFails interface{}

    // The number of input IP datagrams for which no problems were encountered to
    // prevent their continued processing, but were discarded (e.g., for lack of
    // buffer space).  Note that this counter does not include any datagrams
    // discarded while awaiting re-assembly.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    IpIfStatsInDiscards interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  When tracking interface statistics, the counter of the
    // interface to which these datagrams were addressed is incremented.  This
    // interface might not be the same as the   input interface for some of the
    // datagrams.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IpIfStatsInDelivers interface{}

    // The total number of datagrams successfully delivered to IP user-protocols
    // (including ICMP).  This object counts the same packets as
    // ipIfStatsInDelivers, but allows for larger values.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCInDelivers interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  Note that this counter
    // does not include any datagrams counted in ipIfStatsOutForwDatagrams. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsOutRequests interface{}

    // The total number of IP datagrams that local IP user- protocols (including
    // ICMP) supplied to IP in requests for transmission.  This object counts the
    // same packets as   ipIfStatsOutRequests, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCOutRequests interface{}

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
    IpIfStatsOutForwDatagrams interface{}

    // The number of datagrams for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.  This object counts the same packets as
    // ipIfStatsOutForwDatagrams, but allows for larger values.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of  
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCOutForwDatagrams interface{}

    // The number of output IP datagrams for which no problem was encountered to
    // prevent their transmission to their destination, but were discarded (e.g.,
    // for lack of buffer space).  Note that this counter would include datagrams
    // counted in ipIfStatsOutForwDatagrams if any such datagrams met this
    // (discretionary) discard criterion.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    IpIfStatsOutDiscards interface{}

    // The number of IP datagrams that would require fragmentation in order to be
    // transmitted.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for a successfully fragmented datagram. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsOutFragReqds interface{}

    // The number of IP datagrams that have been successfully fragmented.  When
    // tracking interface statistics, the counter of the   outgoing interface is
    // incremented for a successfully fragmented datagram.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsOutFragOKs interface{}

    // The number of IP datagrams that have been discarded because they needed to
    // be fragmented but could not be.  This includes IPv4 packets that have the
    // DF bit set and IPv6 packets that are being forwarded and exceed the
    // outgoing link MTU.  When tracking interface statistics, the counter of the
    // outgoing interface is incremented for an unsuccessfully fragmented
    // datagram.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IpIfStatsOutFragFails interface{}

    // The number of output datagram fragments that have been generated as a
    // result of IP fragmentation.  When tracking interface statistics, the
    // counter of the outgoing interface is incremented for a successfully
    // fragmented datagram.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpIfStatsOutFragCreates interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This includes datagrams generated locally and
    // those forwarded by this entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ipIfStatsDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    IpIfStatsOutTransmits interface{}

    // The total number of IP datagrams that this entity supplied to the lower
    // layers for transmission.  This object counts the same datagrams as
    // ipIfStatsOutTransmits, but allows for larger values.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCOutTransmits interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  Octets from datagrams counted in ipIfStatsOutTransmits
    // MUST be counted here.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ipIfStatsDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IpIfStatsOutOctets interface{}

    // The total number of octets in IP datagrams delivered to the lower layers
    // for transmission.  This objects counts the same octets as
    // ipIfStatsOutOctets, but allows for larger values.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCOutOctets interface{}

    // The number of IP multicast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsInMcastPkts interface{}

    // The number of IP multicast datagrams received.  This object counts the same
    // datagrams as ipIfStatsInMcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCInMcastPkts interface{}

    // The total number of octets received in IP multicast   datagrams.  Octets
    // from datagrams counted in ipIfStatsInMcastPkts MUST be counted here. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsInMcastOctets interface{}

    // The total number of octets received in IP multicast datagrams.  This object
    // counts the same octets as ipIfStatsInMcastOctets, but allows for larger
    // values.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    IpIfStatsHCInMcastOctets interface{}

    // The number of IP multicast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsOutMcastPkts interface{}

    // The number of IP multicast datagrams transmitted.  This object counts the
    // same datagrams as ipIfStatsOutMcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other   times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCOutMcastPkts interface{}

    // The total number of octets transmitted in IP multicast datagrams.  Octets
    // from datagrams counted in ipIfStatsOutMcastPkts MUST be counted here. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsOutMcastOctets interface{}

    // The total number of octets transmitted in IP multicast datagrams.  This
    // object counts the same octets as ipIfStatsOutMcastOctets, but allows for
    // larger values.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ipIfStatsDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615.
    IpIfStatsHCOutMcastOctets interface{}

    // The number of IP broadcast datagrams received.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsInBcastPkts interface{}

    // The number of IP broadcast datagrams received.  This object counts the same
    // datagrams as ipIfStatsInBcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCInBcastPkts interface{}

    // The number of IP broadcast datagrams transmitted.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IpIfStatsOutBcastPkts interface{}

    // The number of IP broadcast datagrams transmitted.  This object counts the
    // same datagrams as ipIfStatsOutBcastPkts, but allows for larger values. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ipIfStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IpIfStatsHCOutBcastPkts interface{}

    // The value of sysUpTime on the most recent occasion at which   any one or
    // more of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    IpIfStatsDiscontinuityTime interface{}

    // The minimum reasonable polling interval for this entry. This object
    // provides an indication of the minimum amount of time required to update the
    // counters in this entry. The type is interface{} with range: 0..4294967295.
    // Units are milli-seconds.
    IpIfStatsRefreshRate interface{}
}

func (ipIfStatsEntry *IPMIB_IpIfStatsTable_IpIfStatsEntry) GetEntityData() *types.CommonEntityData {
    ipIfStatsEntry.EntityData.YFilter = ipIfStatsEntry.YFilter
    ipIfStatsEntry.EntityData.YangName = "ipIfStatsEntry"
    ipIfStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    ipIfStatsEntry.EntityData.ParentYangName = "ipIfStatsTable"
    ipIfStatsEntry.EntityData.SegmentPath = "ipIfStatsEntry" + types.AddKeyToken(ipIfStatsEntry.IpIfStatsIPVersion, "ipIfStatsIPVersion") + types.AddKeyToken(ipIfStatsEntry.IpIfStatsIfIndex, "ipIfStatsIfIndex")
    ipIfStatsEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipIfStatsTable/" + ipIfStatsEntry.EntityData.SegmentPath
    ipIfStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipIfStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipIfStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipIfStatsEntry.EntityData.Children = types.NewOrderedMap()
    ipIfStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsIPVersion", types.YLeaf{"IpIfStatsIPVersion", ipIfStatsEntry.IpIfStatsIPVersion})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsIfIndex", types.YLeaf{"IpIfStatsIfIndex", ipIfStatsEntry.IpIfStatsIfIndex})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInReceives", types.YLeaf{"IpIfStatsInReceives", ipIfStatsEntry.IpIfStatsInReceives})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInReceives", types.YLeaf{"IpIfStatsHCInReceives", ipIfStatsEntry.IpIfStatsHCInReceives})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInOctets", types.YLeaf{"IpIfStatsInOctets", ipIfStatsEntry.IpIfStatsInOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInOctets", types.YLeaf{"IpIfStatsHCInOctets", ipIfStatsEntry.IpIfStatsHCInOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInHdrErrors", types.YLeaf{"IpIfStatsInHdrErrors", ipIfStatsEntry.IpIfStatsInHdrErrors})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInNoRoutes", types.YLeaf{"IpIfStatsInNoRoutes", ipIfStatsEntry.IpIfStatsInNoRoutes})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInAddrErrors", types.YLeaf{"IpIfStatsInAddrErrors", ipIfStatsEntry.IpIfStatsInAddrErrors})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInUnknownProtos", types.YLeaf{"IpIfStatsInUnknownProtos", ipIfStatsEntry.IpIfStatsInUnknownProtos})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInTruncatedPkts", types.YLeaf{"IpIfStatsInTruncatedPkts", ipIfStatsEntry.IpIfStatsInTruncatedPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInForwDatagrams", types.YLeaf{"IpIfStatsInForwDatagrams", ipIfStatsEntry.IpIfStatsInForwDatagrams})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInForwDatagrams", types.YLeaf{"IpIfStatsHCInForwDatagrams", ipIfStatsEntry.IpIfStatsHCInForwDatagrams})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsReasmReqds", types.YLeaf{"IpIfStatsReasmReqds", ipIfStatsEntry.IpIfStatsReasmReqds})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsReasmOKs", types.YLeaf{"IpIfStatsReasmOKs", ipIfStatsEntry.IpIfStatsReasmOKs})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsReasmFails", types.YLeaf{"IpIfStatsReasmFails", ipIfStatsEntry.IpIfStatsReasmFails})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInDiscards", types.YLeaf{"IpIfStatsInDiscards", ipIfStatsEntry.IpIfStatsInDiscards})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInDelivers", types.YLeaf{"IpIfStatsInDelivers", ipIfStatsEntry.IpIfStatsInDelivers})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInDelivers", types.YLeaf{"IpIfStatsHCInDelivers", ipIfStatsEntry.IpIfStatsHCInDelivers})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutRequests", types.YLeaf{"IpIfStatsOutRequests", ipIfStatsEntry.IpIfStatsOutRequests})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutRequests", types.YLeaf{"IpIfStatsHCOutRequests", ipIfStatsEntry.IpIfStatsHCOutRequests})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutForwDatagrams", types.YLeaf{"IpIfStatsOutForwDatagrams", ipIfStatsEntry.IpIfStatsOutForwDatagrams})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutForwDatagrams", types.YLeaf{"IpIfStatsHCOutForwDatagrams", ipIfStatsEntry.IpIfStatsHCOutForwDatagrams})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutDiscards", types.YLeaf{"IpIfStatsOutDiscards", ipIfStatsEntry.IpIfStatsOutDiscards})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutFragReqds", types.YLeaf{"IpIfStatsOutFragReqds", ipIfStatsEntry.IpIfStatsOutFragReqds})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutFragOKs", types.YLeaf{"IpIfStatsOutFragOKs", ipIfStatsEntry.IpIfStatsOutFragOKs})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutFragFails", types.YLeaf{"IpIfStatsOutFragFails", ipIfStatsEntry.IpIfStatsOutFragFails})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutFragCreates", types.YLeaf{"IpIfStatsOutFragCreates", ipIfStatsEntry.IpIfStatsOutFragCreates})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutTransmits", types.YLeaf{"IpIfStatsOutTransmits", ipIfStatsEntry.IpIfStatsOutTransmits})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutTransmits", types.YLeaf{"IpIfStatsHCOutTransmits", ipIfStatsEntry.IpIfStatsHCOutTransmits})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutOctets", types.YLeaf{"IpIfStatsOutOctets", ipIfStatsEntry.IpIfStatsOutOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutOctets", types.YLeaf{"IpIfStatsHCOutOctets", ipIfStatsEntry.IpIfStatsHCOutOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInMcastPkts", types.YLeaf{"IpIfStatsInMcastPkts", ipIfStatsEntry.IpIfStatsInMcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInMcastPkts", types.YLeaf{"IpIfStatsHCInMcastPkts", ipIfStatsEntry.IpIfStatsHCInMcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInMcastOctets", types.YLeaf{"IpIfStatsInMcastOctets", ipIfStatsEntry.IpIfStatsInMcastOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInMcastOctets", types.YLeaf{"IpIfStatsHCInMcastOctets", ipIfStatsEntry.IpIfStatsHCInMcastOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutMcastPkts", types.YLeaf{"IpIfStatsOutMcastPkts", ipIfStatsEntry.IpIfStatsOutMcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutMcastPkts", types.YLeaf{"IpIfStatsHCOutMcastPkts", ipIfStatsEntry.IpIfStatsHCOutMcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutMcastOctets", types.YLeaf{"IpIfStatsOutMcastOctets", ipIfStatsEntry.IpIfStatsOutMcastOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutMcastOctets", types.YLeaf{"IpIfStatsHCOutMcastOctets", ipIfStatsEntry.IpIfStatsHCOutMcastOctets})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsInBcastPkts", types.YLeaf{"IpIfStatsInBcastPkts", ipIfStatsEntry.IpIfStatsInBcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCInBcastPkts", types.YLeaf{"IpIfStatsHCInBcastPkts", ipIfStatsEntry.IpIfStatsHCInBcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsOutBcastPkts", types.YLeaf{"IpIfStatsOutBcastPkts", ipIfStatsEntry.IpIfStatsOutBcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsHCOutBcastPkts", types.YLeaf{"IpIfStatsHCOutBcastPkts", ipIfStatsEntry.IpIfStatsHCOutBcastPkts})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsDiscontinuityTime", types.YLeaf{"IpIfStatsDiscontinuityTime", ipIfStatsEntry.IpIfStatsDiscontinuityTime})
    ipIfStatsEntry.EntityData.Leafs.Append("ipIfStatsRefreshRate", types.YLeaf{"IpIfStatsRefreshRate", ipIfStatsEntry.IpIfStatsRefreshRate})

    ipIfStatsEntry.EntityData.YListKeys = []string {"IpIfStatsIPVersion", "IpIfStatsIfIndex"}

    return &(ipIfStatsEntry.EntityData)
}

// IPMIB_IpAddressPrefixTable
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
type IPMIB_IpAddressPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the ipAddressPrefixTable. The type is slice of
    // IPMIB_IpAddressPrefixTable_IpAddressPrefixEntry.
    IpAddressPrefixEntry []*IPMIB_IpAddressPrefixTable_IpAddressPrefixEntry
}

func (ipAddressPrefixTable *IPMIB_IpAddressPrefixTable) GetEntityData() *types.CommonEntityData {
    ipAddressPrefixTable.EntityData.YFilter = ipAddressPrefixTable.YFilter
    ipAddressPrefixTable.EntityData.YangName = "ipAddressPrefixTable"
    ipAddressPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    ipAddressPrefixTable.EntityData.ParentYangName = "IP-MIB"
    ipAddressPrefixTable.EntityData.SegmentPath = "ipAddressPrefixTable"
    ipAddressPrefixTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipAddressPrefixTable.EntityData.SegmentPath
    ipAddressPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipAddressPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipAddressPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipAddressPrefixTable.EntityData.Children = types.NewOrderedMap()
    ipAddressPrefixTable.EntityData.Children.Append("ipAddressPrefixEntry", types.YChild{"IpAddressPrefixEntry", nil})
    for i := range ipAddressPrefixTable.IpAddressPrefixEntry {
        ipAddressPrefixTable.EntityData.Children.Append(types.GetSegmentPath(ipAddressPrefixTable.IpAddressPrefixEntry[i]), types.YChild{"IpAddressPrefixEntry", ipAddressPrefixTable.IpAddressPrefixEntry[i]})
    }
    ipAddressPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    ipAddressPrefixTable.EntityData.YListKeys = []string {}

    return &(ipAddressPrefixTable.EntityData)
}

// IPMIB_IpAddressPrefixTable_IpAddressPrefixEntry
// An entry in the ipAddressPrefixTable.
type IPMIB_IpAddressPrefixTable_IpAddressPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies the
    // interface on which this prefix is configured.  The interface identified by
    // a particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    IpAddressPrefixIfIndex interface{}

    // This attribute is a key. The address type of ipAddressPrefix. The type is
    // InetAddressType.
    IpAddressPrefixType interface{}

    // This attribute is a key. The address prefix.  The address type of this
    // object is specified in ipAddressPrefixType.  The length of this object is
    // the standard length for objects of that type (4 or 16 bytes).  Any bits
    // after ipAddressPrefixLength must be zero.  Implementors need to be aware
    // that, if the size of ipAddressPrefixPrefix exceeds 114 octets, then OIDS of
    // instances of columns in this row will have more than 128 sub-identifiers
    // and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3. The type is string
    // with length: 0..255.
    IpAddressPrefixPrefix interface{}

    // This attribute is a key. The prefix length associated with this prefix. 
    // The value 0 has no special meaning for this object.  It simply refers to
    // address '::/0'. The type is interface{} with range: 0..2040.
    IpAddressPrefixLength interface{}

    // The origin of this prefix. The type is IpAddressPrefixOriginTC.
    IpAddressPrefixOrigin interface{}

    // This object has the value 'true(1)', if this prefix can be used for on-link
    // determination; otherwise, the value is 'false(2)'.  The default for IPv4
    // prefixes is 'true(1)'. The type is bool.
    IpAddressPrefixOnLinkFlag interface{}

    // Autonomous address configuration flag.  When true(1), indicates that this
    // prefix can be used for autonomous address configuration (i.e., can be used
    // to form a local interface address).  If false(2), it is not used to auto-
    // configure a local interface address.  The default for IPv4 prefixes is
    // 'false(2)'. The type is bool.
    IpAddressPrefixAutonomousFlag interface{}

    // The remaining length of time, in seconds, that this prefix will continue to
    // be preferred, i.e., time until deprecation.  A value of 4,294,967,295
    // represents infinity.  The address generated from a deprecated prefix should
    // no longer be used as a source address in new communications, but packets
    // received on such an interface are processed as expected.  The default for
    // IPv4 prefixes is 4,294,967,295 (infinity). The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    IpAddressPrefixAdvPreferredLifetime interface{}

    // The remaining length of time, in seconds, that this prefix will continue to
    // be valid, i.e., time until invalidation.  A value of 4,294,967,295
    // represents infinity.  The address generated from an invalidated prefix
    // should not appear as the destination or source address of a packet.   The
    // default for IPv4 prefixes is 4,294,967,295 (infinity). The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    IpAddressPrefixAdvValidLifetime interface{}
}

func (ipAddressPrefixEntry *IPMIB_IpAddressPrefixTable_IpAddressPrefixEntry) GetEntityData() *types.CommonEntityData {
    ipAddressPrefixEntry.EntityData.YFilter = ipAddressPrefixEntry.YFilter
    ipAddressPrefixEntry.EntityData.YangName = "ipAddressPrefixEntry"
    ipAddressPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    ipAddressPrefixEntry.EntityData.ParentYangName = "ipAddressPrefixTable"
    ipAddressPrefixEntry.EntityData.SegmentPath = "ipAddressPrefixEntry" + types.AddKeyToken(ipAddressPrefixEntry.IpAddressPrefixIfIndex, "ipAddressPrefixIfIndex") + types.AddKeyToken(ipAddressPrefixEntry.IpAddressPrefixType, "ipAddressPrefixType") + types.AddKeyToken(ipAddressPrefixEntry.IpAddressPrefixPrefix, "ipAddressPrefixPrefix") + types.AddKeyToken(ipAddressPrefixEntry.IpAddressPrefixLength, "ipAddressPrefixLength")
    ipAddressPrefixEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipAddressPrefixTable/" + ipAddressPrefixEntry.EntityData.SegmentPath
    ipAddressPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipAddressPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipAddressPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipAddressPrefixEntry.EntityData.Children = types.NewOrderedMap()
    ipAddressPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixIfIndex", types.YLeaf{"IpAddressPrefixIfIndex", ipAddressPrefixEntry.IpAddressPrefixIfIndex})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixType", types.YLeaf{"IpAddressPrefixType", ipAddressPrefixEntry.IpAddressPrefixType})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixPrefix", types.YLeaf{"IpAddressPrefixPrefix", ipAddressPrefixEntry.IpAddressPrefixPrefix})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixLength", types.YLeaf{"IpAddressPrefixLength", ipAddressPrefixEntry.IpAddressPrefixLength})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixOrigin", types.YLeaf{"IpAddressPrefixOrigin", ipAddressPrefixEntry.IpAddressPrefixOrigin})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixOnLinkFlag", types.YLeaf{"IpAddressPrefixOnLinkFlag", ipAddressPrefixEntry.IpAddressPrefixOnLinkFlag})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixAutonomousFlag", types.YLeaf{"IpAddressPrefixAutonomousFlag", ipAddressPrefixEntry.IpAddressPrefixAutonomousFlag})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixAdvPreferredLifetime", types.YLeaf{"IpAddressPrefixAdvPreferredLifetime", ipAddressPrefixEntry.IpAddressPrefixAdvPreferredLifetime})
    ipAddressPrefixEntry.EntityData.Leafs.Append("ipAddressPrefixAdvValidLifetime", types.YLeaf{"IpAddressPrefixAdvValidLifetime", ipAddressPrefixEntry.IpAddressPrefixAdvValidLifetime})

    ipAddressPrefixEntry.EntityData.YListKeys = []string {"IpAddressPrefixIfIndex", "IpAddressPrefixType", "IpAddressPrefixPrefix", "IpAddressPrefixLength"}

    return &(ipAddressPrefixEntry.EntityData)
}

// IPMIB_IpAddressTable
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
type IPMIB_IpAddressTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An address mapping for a particular interface. The type is slice of
    // IPMIB_IpAddressTable_IpAddressEntry.
    IpAddressEntry []*IPMIB_IpAddressTable_IpAddressEntry
}

func (ipAddressTable *IPMIB_IpAddressTable) GetEntityData() *types.CommonEntityData {
    ipAddressTable.EntityData.YFilter = ipAddressTable.YFilter
    ipAddressTable.EntityData.YangName = "ipAddressTable"
    ipAddressTable.EntityData.BundleName = "cisco_ios_xe"
    ipAddressTable.EntityData.ParentYangName = "IP-MIB"
    ipAddressTable.EntityData.SegmentPath = "ipAddressTable"
    ipAddressTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipAddressTable.EntityData.SegmentPath
    ipAddressTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipAddressTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipAddressTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipAddressTable.EntityData.Children = types.NewOrderedMap()
    ipAddressTable.EntityData.Children.Append("ipAddressEntry", types.YChild{"IpAddressEntry", nil})
    for i := range ipAddressTable.IpAddressEntry {
        ipAddressTable.EntityData.Children.Append(types.GetSegmentPath(ipAddressTable.IpAddressEntry[i]), types.YChild{"IpAddressEntry", ipAddressTable.IpAddressEntry[i]})
    }
    ipAddressTable.EntityData.Leafs = types.NewOrderedMap()

    ipAddressTable.EntityData.YListKeys = []string {}

    return &(ipAddressTable.EntityData)
}

// IPMIB_IpAddressTable_IpAddressEntry
// An address mapping for a particular interface.
type IPMIB_IpAddressTable_IpAddressEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address type of ipAddressAddr. The type is
    // InetAddressType.
    IpAddressAddrType interface{}

    // This attribute is a key. The IP address to which this entry's addressing
    // information   pertains.  The address type of this object is specified in
    // ipAddressAddrType.  Implementors need to be aware that if the size of
    // ipAddressAddr exceeds 116 octets, then OIDS of instances of columns in this
    // row will have more than 128 sub-identifiers and cannot be accessed using
    // SNMPv1, SNMPv2c, or SNMPv3. The type is string with length: 0..255.
    IpAddressAddr interface{}

    // The index value that uniquely identifies the interface to which this entry
    // is applicable.  The interface identified by a particular value of this
    // index is the same interface as identified by the same value of the IF-MIB's
    // ifIndex. The type is interface{} with range: 1..2147483647.
    IpAddressIfIndex interface{}

    // The type of address.  broadcast(3) is not a valid value for IPv6 addresses
    // (RFC 3513). The type is IpAddressType.
    IpAddressType interface{}

    // A pointer to the row in the prefix table to which this address belongs. 
    // May be { 0 0 } if there is no such row. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    IpAddressPrefix interface{}

    // The origin of the address. The type is IpAddressOriginTC.
    IpAddressOrigin interface{}

    // The status of the address, describing if the address can be used for
    // communication.  In the absence of other information, an IPv4 address is
    // always preferred(1). The type is IpAddressStatusTC.
    IpAddressStatus interface{}

    // The value of sysUpTime at the time this entry was created. If this entry
    // was created prior to the last re- initialization of the local network
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    IpAddressCreated interface{}

    // The value of sysUpTime at the time this entry was last updated.  If this
    // entry was updated prior to the last re- initialization of the local network
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    IpAddressLastChanged interface{}

    // The status of this conceptual row.  The RowStatus TC requires that this
    // DESCRIPTION clause states under which circumstances other objects in this
    // row   can be modified.  The value of this object has no effect on whether
    // other objects in this conceptual row can be modified.  A conceptual row can
    // not be made active until the ipAddressIfIndex has been set to a valid
    // index. The type is RowStatus.
    IpAddressRowStatus interface{}

    // The storage type for this conceptual row.  If this object has a value of
    // 'permanent', then no other objects are required to be able to be modified.
    // The type is StorageType.
    IpAddressStorageType interface{}
}

func (ipAddressEntry *IPMIB_IpAddressTable_IpAddressEntry) GetEntityData() *types.CommonEntityData {
    ipAddressEntry.EntityData.YFilter = ipAddressEntry.YFilter
    ipAddressEntry.EntityData.YangName = "ipAddressEntry"
    ipAddressEntry.EntityData.BundleName = "cisco_ios_xe"
    ipAddressEntry.EntityData.ParentYangName = "ipAddressTable"
    ipAddressEntry.EntityData.SegmentPath = "ipAddressEntry" + types.AddKeyToken(ipAddressEntry.IpAddressAddrType, "ipAddressAddrType") + types.AddKeyToken(ipAddressEntry.IpAddressAddr, "ipAddressAddr")
    ipAddressEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipAddressTable/" + ipAddressEntry.EntityData.SegmentPath
    ipAddressEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipAddressEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipAddressEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipAddressEntry.EntityData.Children = types.NewOrderedMap()
    ipAddressEntry.EntityData.Leafs = types.NewOrderedMap()
    ipAddressEntry.EntityData.Leafs.Append("ipAddressAddrType", types.YLeaf{"IpAddressAddrType", ipAddressEntry.IpAddressAddrType})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressAddr", types.YLeaf{"IpAddressAddr", ipAddressEntry.IpAddressAddr})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressIfIndex", types.YLeaf{"IpAddressIfIndex", ipAddressEntry.IpAddressIfIndex})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressType", types.YLeaf{"IpAddressType", ipAddressEntry.IpAddressType})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressPrefix", types.YLeaf{"IpAddressPrefix", ipAddressEntry.IpAddressPrefix})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressOrigin", types.YLeaf{"IpAddressOrigin", ipAddressEntry.IpAddressOrigin})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressStatus", types.YLeaf{"IpAddressStatus", ipAddressEntry.IpAddressStatus})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressCreated", types.YLeaf{"IpAddressCreated", ipAddressEntry.IpAddressCreated})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressLastChanged", types.YLeaf{"IpAddressLastChanged", ipAddressEntry.IpAddressLastChanged})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressRowStatus", types.YLeaf{"IpAddressRowStatus", ipAddressEntry.IpAddressRowStatus})
    ipAddressEntry.EntityData.Leafs.Append("ipAddressStorageType", types.YLeaf{"IpAddressStorageType", ipAddressEntry.IpAddressStorageType})

    ipAddressEntry.EntityData.YListKeys = []string {"IpAddressAddrType", "IpAddressAddr"}

    return &(ipAddressEntry.EntityData)
}

// IPMIB_IpAddressTable_IpAddressEntry_IpAddressType represents IPv6 addresses (RFC 3513).
type IPMIB_IpAddressTable_IpAddressEntry_IpAddressType string

const (
    IPMIB_IpAddressTable_IpAddressEntry_IpAddressType_unicast IPMIB_IpAddressTable_IpAddressEntry_IpAddressType = "unicast"

    IPMIB_IpAddressTable_IpAddressEntry_IpAddressType_anycast IPMIB_IpAddressTable_IpAddressEntry_IpAddressType = "anycast"

    IPMIB_IpAddressTable_IpAddressEntry_IpAddressType_broadcast IPMIB_IpAddressTable_IpAddressEntry_IpAddressType = "broadcast"
)

// IPMIB_IpNetToPhysicalTable
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
type IPMIB_IpNetToPhysicalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains one IP address to `physical' address equivalence. The
    // type is slice of IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry.
    IpNetToPhysicalEntry []*IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry
}

func (ipNetToPhysicalTable *IPMIB_IpNetToPhysicalTable) GetEntityData() *types.CommonEntityData {
    ipNetToPhysicalTable.EntityData.YFilter = ipNetToPhysicalTable.YFilter
    ipNetToPhysicalTable.EntityData.YangName = "ipNetToPhysicalTable"
    ipNetToPhysicalTable.EntityData.BundleName = "cisco_ios_xe"
    ipNetToPhysicalTable.EntityData.ParentYangName = "IP-MIB"
    ipNetToPhysicalTable.EntityData.SegmentPath = "ipNetToPhysicalTable"
    ipNetToPhysicalTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipNetToPhysicalTable.EntityData.SegmentPath
    ipNetToPhysicalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipNetToPhysicalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipNetToPhysicalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipNetToPhysicalTable.EntityData.Children = types.NewOrderedMap()
    ipNetToPhysicalTable.EntityData.Children.Append("ipNetToPhysicalEntry", types.YChild{"IpNetToPhysicalEntry", nil})
    for i := range ipNetToPhysicalTable.IpNetToPhysicalEntry {
        ipNetToPhysicalTable.EntityData.Children.Append(types.GetSegmentPath(ipNetToPhysicalTable.IpNetToPhysicalEntry[i]), types.YChild{"IpNetToPhysicalEntry", ipNetToPhysicalTable.IpNetToPhysicalEntry[i]})
    }
    ipNetToPhysicalTable.EntityData.Leafs = types.NewOrderedMap()

    ipNetToPhysicalTable.EntityData.YListKeys = []string {}

    return &(ipNetToPhysicalTable.EntityData)
}

// IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry
// Each entry contains one IP address to `physical' address
// equivalence.
type IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    IpNetToPhysicalIfIndex interface{}

    // This attribute is a key. The type of ipNetToPhysicalNetAddress. The type is
    // InetAddressType.
    IpNetToPhysicalNetAddressType interface{}

    // This attribute is a key. The IP Address corresponding to the
    // media-dependent `physical' address.  The address type of this object is
    // specified in ipNetToPhysicalAddressType.  Implementors need to be aware
    // that if the size of   ipNetToPhysicalNetAddress exceeds 115 octets, then
    // OIDS of instances of columns in this row will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3.
    // The type is string with length: 0..255.
    IpNetToPhysicalNetAddress interface{}

    // The media-dependent `physical' address.  As the entries in this table are
    // typically not persistent when this object is written the entity SHOULD NOT
    // save the change to non-volatile storage. The type is string with length:
    // 0..65535.
    IpNetToPhysicalPhysAddress interface{}

    // The value of sysUpTime at the time this entry was last updated.  If this
    // entry was updated prior to the last re- initialization of the local network
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    IpNetToPhysicalLastUpdated interface{}

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
    // non-volatile storage. The type is IpNetToPhysicalType.
    IpNetToPhysicalType interface{}

    // The Neighbor Unreachability Detection state for the interface when the
    // address mapping in this entry is used. If Neighbor Unreachability Detection
    // is not in use (e.g. for IPv4), this object is always unknown(6). The type
    // is IpNetToPhysicalState.
    IpNetToPhysicalState interface{}

    // The status of this conceptual row.  The RowStatus TC requires that this
    // DESCRIPTION clause states under which circumstances other objects in this
    // row can be modified.  The value of this object has no effect on whether
    // other objects in this conceptual row can be modified.  A conceptual row can
    // not be made active until the ipNetToPhysicalPhysAddress object has been
    // set.  Note that if the ipNetToPhysicalType is set to 'invalid', the managed
    // node may delete the entry independent of the state of this object. The type
    // is RowStatus.
    IpNetToPhysicalRowStatus interface{}
}

func (ipNetToPhysicalEntry *IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry) GetEntityData() *types.CommonEntityData {
    ipNetToPhysicalEntry.EntityData.YFilter = ipNetToPhysicalEntry.YFilter
    ipNetToPhysicalEntry.EntityData.YangName = "ipNetToPhysicalEntry"
    ipNetToPhysicalEntry.EntityData.BundleName = "cisco_ios_xe"
    ipNetToPhysicalEntry.EntityData.ParentYangName = "ipNetToPhysicalTable"
    ipNetToPhysicalEntry.EntityData.SegmentPath = "ipNetToPhysicalEntry" + types.AddKeyToken(ipNetToPhysicalEntry.IpNetToPhysicalIfIndex, "ipNetToPhysicalIfIndex") + types.AddKeyToken(ipNetToPhysicalEntry.IpNetToPhysicalNetAddressType, "ipNetToPhysicalNetAddressType") + types.AddKeyToken(ipNetToPhysicalEntry.IpNetToPhysicalNetAddress, "ipNetToPhysicalNetAddress")
    ipNetToPhysicalEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipNetToPhysicalTable/" + ipNetToPhysicalEntry.EntityData.SegmentPath
    ipNetToPhysicalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipNetToPhysicalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipNetToPhysicalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipNetToPhysicalEntry.EntityData.Children = types.NewOrderedMap()
    ipNetToPhysicalEntry.EntityData.Leafs = types.NewOrderedMap()
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalIfIndex", types.YLeaf{"IpNetToPhysicalIfIndex", ipNetToPhysicalEntry.IpNetToPhysicalIfIndex})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalNetAddressType", types.YLeaf{"IpNetToPhysicalNetAddressType", ipNetToPhysicalEntry.IpNetToPhysicalNetAddressType})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalNetAddress", types.YLeaf{"IpNetToPhysicalNetAddress", ipNetToPhysicalEntry.IpNetToPhysicalNetAddress})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalPhysAddress", types.YLeaf{"IpNetToPhysicalPhysAddress", ipNetToPhysicalEntry.IpNetToPhysicalPhysAddress})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalLastUpdated", types.YLeaf{"IpNetToPhysicalLastUpdated", ipNetToPhysicalEntry.IpNetToPhysicalLastUpdated})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalType", types.YLeaf{"IpNetToPhysicalType", ipNetToPhysicalEntry.IpNetToPhysicalType})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalState", types.YLeaf{"IpNetToPhysicalState", ipNetToPhysicalEntry.IpNetToPhysicalState})
    ipNetToPhysicalEntry.EntityData.Leafs.Append("ipNetToPhysicalRowStatus", types.YLeaf{"IpNetToPhysicalRowStatus", ipNetToPhysicalEntry.IpNetToPhysicalRowStatus})

    ipNetToPhysicalEntry.EntityData.YListKeys = []string {"IpNetToPhysicalIfIndex", "IpNetToPhysicalNetAddressType", "IpNetToPhysicalNetAddress"}

    return &(ipNetToPhysicalEntry.EntityData)
}

// IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState represents IPv4), this object is always unknown(6).
type IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState string

const (
    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_reachable IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "reachable"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_stale IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "stale"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_delay IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "delay"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_probe IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "probe"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_invalid IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "invalid"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_unknown IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "unknown"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState_incomplete IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalState = "incomplete"
)

// IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType represents change to non-volatile storage.
type IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType string

const (
    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType_other IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType = "other"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType_invalid IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType = "invalid"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType_dynamic IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType = "dynamic"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType_static IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType = "static"

    IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType_local IPMIB_IpNetToPhysicalTable_IpNetToPhysicalEntry_IpNetToPhysicalType = "local"
)

// IPMIB_Ipv6ScopeZoneIndexTable
// The table used to describe IPv6 unicast and multicast scope
// zones.
// 
// For those objects that have names rather than numbers, the
// names were chosen to coincide with the names used in the
// IPv6 address architecture document. 
type IPMIB_Ipv6ScopeZoneIndexTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the list of scope identifiers on a given interface. The
    // type is slice of IPMIB_Ipv6ScopeZoneIndexTable_Ipv6ScopeZoneIndexEntry.
    Ipv6ScopeZoneIndexEntry []*IPMIB_Ipv6ScopeZoneIndexTable_Ipv6ScopeZoneIndexEntry
}

func (ipv6ScopeZoneIndexTable *IPMIB_Ipv6ScopeZoneIndexTable) GetEntityData() *types.CommonEntityData {
    ipv6ScopeZoneIndexTable.EntityData.YFilter = ipv6ScopeZoneIndexTable.YFilter
    ipv6ScopeZoneIndexTable.EntityData.YangName = "ipv6ScopeZoneIndexTable"
    ipv6ScopeZoneIndexTable.EntityData.BundleName = "cisco_ios_xe"
    ipv6ScopeZoneIndexTable.EntityData.ParentYangName = "IP-MIB"
    ipv6ScopeZoneIndexTable.EntityData.SegmentPath = "ipv6ScopeZoneIndexTable"
    ipv6ScopeZoneIndexTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipv6ScopeZoneIndexTable.EntityData.SegmentPath
    ipv6ScopeZoneIndexTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6ScopeZoneIndexTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6ScopeZoneIndexTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6ScopeZoneIndexTable.EntityData.Children = types.NewOrderedMap()
    ipv6ScopeZoneIndexTable.EntityData.Children.Append("ipv6ScopeZoneIndexEntry", types.YChild{"Ipv6ScopeZoneIndexEntry", nil})
    for i := range ipv6ScopeZoneIndexTable.Ipv6ScopeZoneIndexEntry {
        ipv6ScopeZoneIndexTable.EntityData.Children.Append(types.GetSegmentPath(ipv6ScopeZoneIndexTable.Ipv6ScopeZoneIndexEntry[i]), types.YChild{"Ipv6ScopeZoneIndexEntry", ipv6ScopeZoneIndexTable.Ipv6ScopeZoneIndexEntry[i]})
    }
    ipv6ScopeZoneIndexTable.EntityData.Leafs = types.NewOrderedMap()

    ipv6ScopeZoneIndexTable.EntityData.YListKeys = []string {}

    return &(ipv6ScopeZoneIndexTable.EntityData)
}

// IPMIB_Ipv6ScopeZoneIndexTable_Ipv6ScopeZoneIndexEntry
// Each entry contains the list of scope identifiers on a given
// interface.
type IPMIB_Ipv6ScopeZoneIndexTable_Ipv6ScopeZoneIndexEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies the
    // interface to which these scopes belong.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    Ipv6ScopeZoneIndexIfIndex interface{}

    // The zone index for the link-local scope on this interface. The type is
    // interface{} with range: 0..4294967295.
    Ipv6ScopeZoneIndexLinkLocal interface{}

    // The zone index for scope 3 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndex3 interface{}

    // The zone index for the admin-local scope on this interface. The type is
    // interface{} with range: 0..4294967295.
    Ipv6ScopeZoneIndexAdminLocal interface{}

    // The zone index for the site-local scope on this interface. The type is
    // interface{} with range: 0..4294967295.
    Ipv6ScopeZoneIndexSiteLocal interface{}

    // The zone index for scope 6 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndex6 interface{}

    // The zone index for scope 7 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndex7 interface{}

    // The zone index for the organization-local scope on this interface. The type
    // is interface{} with range: 0..4294967295.
    Ipv6ScopeZoneIndexOrganizationLocal interface{}

    // The zone index for scope 9 on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndex9 interface{}

    // The zone index for scope A on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndexA interface{}

    // The zone index for scope B on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndexB interface{}

    // The zone index for scope C on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndexC interface{}

    // The zone index for scope D on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ipv6ScopeZoneIndexD interface{}
}

func (ipv6ScopeZoneIndexEntry *IPMIB_Ipv6ScopeZoneIndexTable_Ipv6ScopeZoneIndexEntry) GetEntityData() *types.CommonEntityData {
    ipv6ScopeZoneIndexEntry.EntityData.YFilter = ipv6ScopeZoneIndexEntry.YFilter
    ipv6ScopeZoneIndexEntry.EntityData.YangName = "ipv6ScopeZoneIndexEntry"
    ipv6ScopeZoneIndexEntry.EntityData.BundleName = "cisco_ios_xe"
    ipv6ScopeZoneIndexEntry.EntityData.ParentYangName = "ipv6ScopeZoneIndexTable"
    ipv6ScopeZoneIndexEntry.EntityData.SegmentPath = "ipv6ScopeZoneIndexEntry" + types.AddKeyToken(ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexIfIndex, "ipv6ScopeZoneIndexIfIndex")
    ipv6ScopeZoneIndexEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipv6ScopeZoneIndexTable/" + ipv6ScopeZoneIndexEntry.EntityData.SegmentPath
    ipv6ScopeZoneIndexEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6ScopeZoneIndexEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6ScopeZoneIndexEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6ScopeZoneIndexEntry.EntityData.Children = types.NewOrderedMap()
    ipv6ScopeZoneIndexEntry.EntityData.Leafs = types.NewOrderedMap()
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexIfIndex", types.YLeaf{"Ipv6ScopeZoneIndexIfIndex", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexIfIndex})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexLinkLocal", types.YLeaf{"Ipv6ScopeZoneIndexLinkLocal", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexLinkLocal})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndex3", types.YLeaf{"Ipv6ScopeZoneIndex3", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndex3})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexAdminLocal", types.YLeaf{"Ipv6ScopeZoneIndexAdminLocal", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexAdminLocal})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexSiteLocal", types.YLeaf{"Ipv6ScopeZoneIndexSiteLocal", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexSiteLocal})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndex6", types.YLeaf{"Ipv6ScopeZoneIndex6", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndex6})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndex7", types.YLeaf{"Ipv6ScopeZoneIndex7", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndex7})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexOrganizationLocal", types.YLeaf{"Ipv6ScopeZoneIndexOrganizationLocal", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexOrganizationLocal})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndex9", types.YLeaf{"Ipv6ScopeZoneIndex9", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndex9})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexA", types.YLeaf{"Ipv6ScopeZoneIndexA", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexA})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexB", types.YLeaf{"Ipv6ScopeZoneIndexB", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexB})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexC", types.YLeaf{"Ipv6ScopeZoneIndexC", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexC})
    ipv6ScopeZoneIndexEntry.EntityData.Leafs.Append("ipv6ScopeZoneIndexD", types.YLeaf{"Ipv6ScopeZoneIndexD", ipv6ScopeZoneIndexEntry.Ipv6ScopeZoneIndexD})

    ipv6ScopeZoneIndexEntry.EntityData.YListKeys = []string {"Ipv6ScopeZoneIndexIfIndex"}

    return &(ipv6ScopeZoneIndexEntry.EntityData)
}

// IPMIB_IpDefaultRouterTable
// The table used to describe the default routers known to this
// 
// 
// entity.
type IPMIB_IpDefaultRouterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains information about a default router known to this
    // entity. The type is slice of
    // IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry.
    IpDefaultRouterEntry []*IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry
}

func (ipDefaultRouterTable *IPMIB_IpDefaultRouterTable) GetEntityData() *types.CommonEntityData {
    ipDefaultRouterTable.EntityData.YFilter = ipDefaultRouterTable.YFilter
    ipDefaultRouterTable.EntityData.YangName = "ipDefaultRouterTable"
    ipDefaultRouterTable.EntityData.BundleName = "cisco_ios_xe"
    ipDefaultRouterTable.EntityData.ParentYangName = "IP-MIB"
    ipDefaultRouterTable.EntityData.SegmentPath = "ipDefaultRouterTable"
    ipDefaultRouterTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipDefaultRouterTable.EntityData.SegmentPath
    ipDefaultRouterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipDefaultRouterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipDefaultRouterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipDefaultRouterTable.EntityData.Children = types.NewOrderedMap()
    ipDefaultRouterTable.EntityData.Children.Append("ipDefaultRouterEntry", types.YChild{"IpDefaultRouterEntry", nil})
    for i := range ipDefaultRouterTable.IpDefaultRouterEntry {
        ipDefaultRouterTable.EntityData.Children.Append(types.GetSegmentPath(ipDefaultRouterTable.IpDefaultRouterEntry[i]), types.YChild{"IpDefaultRouterEntry", ipDefaultRouterTable.IpDefaultRouterEntry[i]})
    }
    ipDefaultRouterTable.EntityData.Leafs = types.NewOrderedMap()

    ipDefaultRouterTable.EntityData.YListKeys = []string {}

    return &(ipDefaultRouterTable.EntityData)
}

// IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry
// Each entry contains information about a default router known
// to this entity.
type IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address type for this row. The type is
    // InetAddressType.
    IpDefaultRouterAddressType interface{}

    // This attribute is a key. The IP address of the default router represented
    // by this row.  The address type of this object is specified in
    // ipDefaultRouterAddressType.  Implementers need to be aware that if the size
    // of ipDefaultRouterAddress exceeds 115 octets, then OIDS of instances of
    // columns in this row will have more than 128 sub-identifiers and cannot be
    // accessed using SNMPv1, SNMPv2c, or SNMPv3. The type is string with length:
    // 0..255.
    IpDefaultRouterAddress interface{}

    // This attribute is a key. The index value that uniquely identifies the
    // interface by which the router can be reached.  The interface identified by
    // a particular value of this index is the same interface as identified by the
    // same value of the IF-MIB's ifIndex. The type is interface{} with range:
    // 1..2147483647.
    IpDefaultRouterIfIndex interface{}

    // The remaining length of time, in seconds, that this router will continue to
    // be useful as a default router.  A value of zero indicates that it is no
    // longer useful as a default router.  It is left to the implementer of the
    // MIB as to whether a router with a lifetime of zero is removed from the
    // list.  For IPv6, this value should be extracted from the router
    // advertisement messages. The type is interface{} with range: 0..65535. Units
    // are seconds.
    IpDefaultRouterLifetime interface{}

    // An indication of preference given to this router as a default router as
    // described in he Default Router Preferences document.  Treating the value as
    // a 2 bit signed integer allows for simple arithmetic comparisons.  For IPv4
    // routers or IPv6 routers that are not using the updated router advertisement
    // format, this object is set to medium (0). The type is
    // IpDefaultRouterPreference.
    IpDefaultRouterPreference interface{}
}

func (ipDefaultRouterEntry *IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry) GetEntityData() *types.CommonEntityData {
    ipDefaultRouterEntry.EntityData.YFilter = ipDefaultRouterEntry.YFilter
    ipDefaultRouterEntry.EntityData.YangName = "ipDefaultRouterEntry"
    ipDefaultRouterEntry.EntityData.BundleName = "cisco_ios_xe"
    ipDefaultRouterEntry.EntityData.ParentYangName = "ipDefaultRouterTable"
    ipDefaultRouterEntry.EntityData.SegmentPath = "ipDefaultRouterEntry" + types.AddKeyToken(ipDefaultRouterEntry.IpDefaultRouterAddressType, "ipDefaultRouterAddressType") + types.AddKeyToken(ipDefaultRouterEntry.IpDefaultRouterAddress, "ipDefaultRouterAddress") + types.AddKeyToken(ipDefaultRouterEntry.IpDefaultRouterIfIndex, "ipDefaultRouterIfIndex")
    ipDefaultRouterEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipDefaultRouterTable/" + ipDefaultRouterEntry.EntityData.SegmentPath
    ipDefaultRouterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipDefaultRouterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipDefaultRouterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipDefaultRouterEntry.EntityData.Children = types.NewOrderedMap()
    ipDefaultRouterEntry.EntityData.Leafs = types.NewOrderedMap()
    ipDefaultRouterEntry.EntityData.Leafs.Append("ipDefaultRouterAddressType", types.YLeaf{"IpDefaultRouterAddressType", ipDefaultRouterEntry.IpDefaultRouterAddressType})
    ipDefaultRouterEntry.EntityData.Leafs.Append("ipDefaultRouterAddress", types.YLeaf{"IpDefaultRouterAddress", ipDefaultRouterEntry.IpDefaultRouterAddress})
    ipDefaultRouterEntry.EntityData.Leafs.Append("ipDefaultRouterIfIndex", types.YLeaf{"IpDefaultRouterIfIndex", ipDefaultRouterEntry.IpDefaultRouterIfIndex})
    ipDefaultRouterEntry.EntityData.Leafs.Append("ipDefaultRouterLifetime", types.YLeaf{"IpDefaultRouterLifetime", ipDefaultRouterEntry.IpDefaultRouterLifetime})
    ipDefaultRouterEntry.EntityData.Leafs.Append("ipDefaultRouterPreference", types.YLeaf{"IpDefaultRouterPreference", ipDefaultRouterEntry.IpDefaultRouterPreference})

    ipDefaultRouterEntry.EntityData.YListKeys = []string {"IpDefaultRouterAddressType", "IpDefaultRouterAddress", "IpDefaultRouterIfIndex"}

    return &(ipDefaultRouterEntry.EntityData)
}

// IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference represents medium (0).
type IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference string

const (
    IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference_reserved IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference = "reserved"

    IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference_low IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference = "low"

    IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference_medium IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference = "medium"

    IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference_high IPMIB_IpDefaultRouterTable_IpDefaultRouterEntry_IpDefaultRouterPreference = "high"
)

// IPMIB_Ipv6RouterAdvertTable
// The table containing information used to construct router
// advertisements.
type IPMIB_Ipv6RouterAdvertTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information used to construct router advertisements. 
    // Information in this table is persistent, and when this object is written,
    // the entity SHOULD save the change to non-volatile storage. The type is
    // slice of IPMIB_Ipv6RouterAdvertTable_Ipv6RouterAdvertEntry.
    Ipv6RouterAdvertEntry []*IPMIB_Ipv6RouterAdvertTable_Ipv6RouterAdvertEntry
}

func (ipv6RouterAdvertTable *IPMIB_Ipv6RouterAdvertTable) GetEntityData() *types.CommonEntityData {
    ipv6RouterAdvertTable.EntityData.YFilter = ipv6RouterAdvertTable.YFilter
    ipv6RouterAdvertTable.EntityData.YangName = "ipv6RouterAdvertTable"
    ipv6RouterAdvertTable.EntityData.BundleName = "cisco_ios_xe"
    ipv6RouterAdvertTable.EntityData.ParentYangName = "IP-MIB"
    ipv6RouterAdvertTable.EntityData.SegmentPath = "ipv6RouterAdvertTable"
    ipv6RouterAdvertTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + ipv6RouterAdvertTable.EntityData.SegmentPath
    ipv6RouterAdvertTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6RouterAdvertTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6RouterAdvertTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6RouterAdvertTable.EntityData.Children = types.NewOrderedMap()
    ipv6RouterAdvertTable.EntityData.Children.Append("ipv6RouterAdvertEntry", types.YChild{"Ipv6RouterAdvertEntry", nil})
    for i := range ipv6RouterAdvertTable.Ipv6RouterAdvertEntry {
        ipv6RouterAdvertTable.EntityData.Children.Append(types.GetSegmentPath(ipv6RouterAdvertTable.Ipv6RouterAdvertEntry[i]), types.YChild{"Ipv6RouterAdvertEntry", ipv6RouterAdvertTable.Ipv6RouterAdvertEntry[i]})
    }
    ipv6RouterAdvertTable.EntityData.Leafs = types.NewOrderedMap()

    ipv6RouterAdvertTable.EntityData.YListKeys = []string {}

    return &(ipv6RouterAdvertTable.EntityData)
}

// IPMIB_Ipv6RouterAdvertTable_Ipv6RouterAdvertEntry
// An entry containing information used to construct router
// advertisements.
// 
// Information in this table is persistent, and when this
// object is written, the entity SHOULD save the change to
// non-volatile storage.
type IPMIB_Ipv6RouterAdvertTable_Ipv6RouterAdvertEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies the
    // interface on which router advertisements constructed with this information
    // will be transmitted.  The interface identified by a particular value of
    // this index is the same interface as identified by the same value of the
    // IF-MIB's ifIndex. The type is interface{} with range: 1..2147483647.
    Ipv6RouterAdvertIfIndex interface{}

    // A flag indicating whether the router sends periodic router advertisements
    // and responds to router solicitations on this interface. The type is bool.
    Ipv6RouterAdvertSendAdverts interface{}

    // The maximum time allowed between sending unsolicited router  
    // advertisements from this interface. The type is interface{} with range:
    // 4..1800. Units are seconds.
    Ipv6RouterAdvertMaxInterval interface{}

    // The minimum time allowed between sending unsolicited router advertisements
    // from this interface.  The default is 0.33 * ipv6RouterAdvertMaxInterval,
    // however, in the case of a low value for ipv6RouterAdvertMaxInterval, the
    // minimum value for this object is restricted to 3. The type is interface{}
    // with range: 3..1350. Units are seconds.
    Ipv6RouterAdvertMinInterval interface{}

    // The true/false value to be placed into the 'managed address configuration'
    // flag field in router advertisements sent from this interface. The type is
    // bool.
    Ipv6RouterAdvertManagedFlag interface{}

    // The true/false value to be placed into the 'other stateful configuration'
    // flag field in router advertisements sent from this interface. The type is
    // bool.
    Ipv6RouterAdvertOtherConfigFlag interface{}

    // The value to be placed in MTU options sent by the router on this interface.
    // A value of zero indicates that no MTU options are sent. The type is
    // interface{} with range: 0..4294967295.
    Ipv6RouterAdvertLinkMTU interface{}

    // The value to be placed in the reachable time field in router advertisement
    // messages sent from this interface.  A value of zero in the router
    // advertisement indicates that the advertisement isn't specifying a value for
    // reachable time. The type is interface{} with range: 0..3600000. Units are
    // milliseconds.
    Ipv6RouterAdvertReachableTime interface{}

    // The value to be placed in the retransmit timer field in router
    // advertisements sent from this interface.  A value of zero in the router
    // advertisement indicates that the advertisement isn't specifying a value for
    // retrans time. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Ipv6RouterAdvertRetransmitTime interface{}

    // The default value to be placed in the current hop limit field in router
    // advertisements sent from this interface.   The value should be set to the
    // current diameter of the Internet.  A value of zero in the router
    // advertisement indicates that the advertisement isn't specifying a value for
    // curHopLimit.  The default should be set to the value specified in the IANA
    // web pages (www.iana.org) at the time of implementation. The type is
    // interface{} with range: 0..255.
    Ipv6RouterAdvertCurHopLimit interface{}

    // The value to be placed in the router lifetime field of router
    // advertisements sent from this interface.  This value MUST be either 0 or
    // between ipv6RouterAdvertMaxInterval and 9000 seconds.  A value of zero
    // indicates that the router is not to be used as a default router.  The
    // default is 3 * ipv6RouterAdvertMaxInterval. The type is interface{} with
    // range: 0..0 | 4..9000. Units are seconds.
    Ipv6RouterAdvertDefaultLifetime interface{}

    // The status of this conceptual row.  As all objects in this conceptual row
    // have default values, a row can be created and made active by setting this
    // object appropriately.  The RowStatus TC requires that this DESCRIPTION
    // clause states under which circumstances other objects in this row can be
    // modified.  The value of this object has no effect on whether other objects
    // in this conceptual row can be modified. The type is RowStatus.
    Ipv6RouterAdvertRowStatus interface{}
}

func (ipv6RouterAdvertEntry *IPMIB_Ipv6RouterAdvertTable_Ipv6RouterAdvertEntry) GetEntityData() *types.CommonEntityData {
    ipv6RouterAdvertEntry.EntityData.YFilter = ipv6RouterAdvertEntry.YFilter
    ipv6RouterAdvertEntry.EntityData.YangName = "ipv6RouterAdvertEntry"
    ipv6RouterAdvertEntry.EntityData.BundleName = "cisco_ios_xe"
    ipv6RouterAdvertEntry.EntityData.ParentYangName = "ipv6RouterAdvertTable"
    ipv6RouterAdvertEntry.EntityData.SegmentPath = "ipv6RouterAdvertEntry" + types.AddKeyToken(ipv6RouterAdvertEntry.Ipv6RouterAdvertIfIndex, "ipv6RouterAdvertIfIndex")
    ipv6RouterAdvertEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/ipv6RouterAdvertTable/" + ipv6RouterAdvertEntry.EntityData.SegmentPath
    ipv6RouterAdvertEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6RouterAdvertEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6RouterAdvertEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6RouterAdvertEntry.EntityData.Children = types.NewOrderedMap()
    ipv6RouterAdvertEntry.EntityData.Leafs = types.NewOrderedMap()
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertIfIndex", types.YLeaf{"Ipv6RouterAdvertIfIndex", ipv6RouterAdvertEntry.Ipv6RouterAdvertIfIndex})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertSendAdverts", types.YLeaf{"Ipv6RouterAdvertSendAdverts", ipv6RouterAdvertEntry.Ipv6RouterAdvertSendAdverts})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertMaxInterval", types.YLeaf{"Ipv6RouterAdvertMaxInterval", ipv6RouterAdvertEntry.Ipv6RouterAdvertMaxInterval})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertMinInterval", types.YLeaf{"Ipv6RouterAdvertMinInterval", ipv6RouterAdvertEntry.Ipv6RouterAdvertMinInterval})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertManagedFlag", types.YLeaf{"Ipv6RouterAdvertManagedFlag", ipv6RouterAdvertEntry.Ipv6RouterAdvertManagedFlag})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertOtherConfigFlag", types.YLeaf{"Ipv6RouterAdvertOtherConfigFlag", ipv6RouterAdvertEntry.Ipv6RouterAdvertOtherConfigFlag})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertLinkMTU", types.YLeaf{"Ipv6RouterAdvertLinkMTU", ipv6RouterAdvertEntry.Ipv6RouterAdvertLinkMTU})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertReachableTime", types.YLeaf{"Ipv6RouterAdvertReachableTime", ipv6RouterAdvertEntry.Ipv6RouterAdvertReachableTime})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertRetransmitTime", types.YLeaf{"Ipv6RouterAdvertRetransmitTime", ipv6RouterAdvertEntry.Ipv6RouterAdvertRetransmitTime})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertCurHopLimit", types.YLeaf{"Ipv6RouterAdvertCurHopLimit", ipv6RouterAdvertEntry.Ipv6RouterAdvertCurHopLimit})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertDefaultLifetime", types.YLeaf{"Ipv6RouterAdvertDefaultLifetime", ipv6RouterAdvertEntry.Ipv6RouterAdvertDefaultLifetime})
    ipv6RouterAdvertEntry.EntityData.Leafs.Append("ipv6RouterAdvertRowStatus", types.YLeaf{"Ipv6RouterAdvertRowStatus", ipv6RouterAdvertEntry.Ipv6RouterAdvertRowStatus})

    ipv6RouterAdvertEntry.EntityData.YListKeys = []string {"Ipv6RouterAdvertIfIndex"}

    return &(ipv6RouterAdvertEntry.EntityData)
}

// IPMIB_IcmpStatsTable
// The table of generic system-wide ICMP counters.
type IPMIB_IcmpStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the icmpStatsTable. The type is slice of
    // IPMIB_IcmpStatsTable_IcmpStatsEntry.
    IcmpStatsEntry []*IPMIB_IcmpStatsTable_IcmpStatsEntry
}

func (icmpStatsTable *IPMIB_IcmpStatsTable) GetEntityData() *types.CommonEntityData {
    icmpStatsTable.EntityData.YFilter = icmpStatsTable.YFilter
    icmpStatsTable.EntityData.YangName = "icmpStatsTable"
    icmpStatsTable.EntityData.BundleName = "cisco_ios_xe"
    icmpStatsTable.EntityData.ParentYangName = "IP-MIB"
    icmpStatsTable.EntityData.SegmentPath = "icmpStatsTable"
    icmpStatsTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + icmpStatsTable.EntityData.SegmentPath
    icmpStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpStatsTable.EntityData.Children = types.NewOrderedMap()
    icmpStatsTable.EntityData.Children.Append("icmpStatsEntry", types.YChild{"IcmpStatsEntry", nil})
    for i := range icmpStatsTable.IcmpStatsEntry {
        icmpStatsTable.EntityData.Children.Append(types.GetSegmentPath(icmpStatsTable.IcmpStatsEntry[i]), types.YChild{"IcmpStatsEntry", icmpStatsTable.IcmpStatsEntry[i]})
    }
    icmpStatsTable.EntityData.Leafs = types.NewOrderedMap()

    icmpStatsTable.EntityData.YListKeys = []string {}

    return &(icmpStatsTable.EntityData)
}

// IPMIB_IcmpStatsTable_IcmpStatsEntry
// A conceptual row in the icmpStatsTable.
type IPMIB_IcmpStatsTable_IcmpStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP version of the statistics. The type is
    // IpVersion.
    IcmpStatsIPVersion interface{}

    // The total number of ICMP messages that the entity received. Note that this
    // counter includes all those counted by icmpStatsInErrors. The type is
    // interface{} with range: 0..4294967295.
    IcmpStatsInMsgs interface{}

    // The number of ICMP messages that the entity received but determined as
    // having ICMP-specific errors (bad ICMP checksums, bad length, etc.). The
    // type is interface{} with range: 0..4294967295.
    IcmpStatsInErrors interface{}

    // The total number of ICMP messages that the entity attempted to send.  Note
    // that this counter includes all those counted by icmpStatsOutErrors. The
    // type is interface{} with range: 0..4294967295.
    IcmpStatsOutMsgs interface{}

    // The number of ICMP messages that this entity did not send due to problems
    // discovered within ICMP, such as a lack of buffers.  This value should not
    // include errors discovered outside the ICMP layer, such as the inability of
    // IP to route the resultant datagram.  In some implementations, there may be
    // no types of error that contribute to this counter's value. The type is
    // interface{} with range: 0..4294967295.
    IcmpStatsOutErrors interface{}
}

func (icmpStatsEntry *IPMIB_IcmpStatsTable_IcmpStatsEntry) GetEntityData() *types.CommonEntityData {
    icmpStatsEntry.EntityData.YFilter = icmpStatsEntry.YFilter
    icmpStatsEntry.EntityData.YangName = "icmpStatsEntry"
    icmpStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    icmpStatsEntry.EntityData.ParentYangName = "icmpStatsTable"
    icmpStatsEntry.EntityData.SegmentPath = "icmpStatsEntry" + types.AddKeyToken(icmpStatsEntry.IcmpStatsIPVersion, "icmpStatsIPVersion")
    icmpStatsEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/icmpStatsTable/" + icmpStatsEntry.EntityData.SegmentPath
    icmpStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpStatsEntry.EntityData.Children = types.NewOrderedMap()
    icmpStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    icmpStatsEntry.EntityData.Leafs.Append("icmpStatsIPVersion", types.YLeaf{"IcmpStatsIPVersion", icmpStatsEntry.IcmpStatsIPVersion})
    icmpStatsEntry.EntityData.Leafs.Append("icmpStatsInMsgs", types.YLeaf{"IcmpStatsInMsgs", icmpStatsEntry.IcmpStatsInMsgs})
    icmpStatsEntry.EntityData.Leafs.Append("icmpStatsInErrors", types.YLeaf{"IcmpStatsInErrors", icmpStatsEntry.IcmpStatsInErrors})
    icmpStatsEntry.EntityData.Leafs.Append("icmpStatsOutMsgs", types.YLeaf{"IcmpStatsOutMsgs", icmpStatsEntry.IcmpStatsOutMsgs})
    icmpStatsEntry.EntityData.Leafs.Append("icmpStatsOutErrors", types.YLeaf{"IcmpStatsOutErrors", icmpStatsEntry.IcmpStatsOutErrors})

    icmpStatsEntry.EntityData.YListKeys = []string {"IcmpStatsIPVersion"}

    return &(icmpStatsEntry.EntityData)
}

// IPMIB_IcmpMsgStatsTable
// The table of system-wide per-version, per-message type ICMP
// counters.
type IPMIB_IcmpMsgStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the icmpMsgStatsTable.  The system should track each
    // ICMP type value, even if that ICMP type is not supported by the system. 
    // However, a given row need not be instantiated unless a message of that type
    // has been processed, i.e., the row for icmpMsgStatsType=X MAY be
    // instantiated before but MUST be instantiated after the first message with
    // Type=X is received or transmitted.  After receiving or transmitting any
    // succeeding messages with Type=X, the relevant counter must be incremented.
    // The type is slice of IPMIB_IcmpMsgStatsTable_IcmpMsgStatsEntry.
    IcmpMsgStatsEntry []*IPMIB_IcmpMsgStatsTable_IcmpMsgStatsEntry
}

func (icmpMsgStatsTable *IPMIB_IcmpMsgStatsTable) GetEntityData() *types.CommonEntityData {
    icmpMsgStatsTable.EntityData.YFilter = icmpMsgStatsTable.YFilter
    icmpMsgStatsTable.EntityData.YangName = "icmpMsgStatsTable"
    icmpMsgStatsTable.EntityData.BundleName = "cisco_ios_xe"
    icmpMsgStatsTable.EntityData.ParentYangName = "IP-MIB"
    icmpMsgStatsTable.EntityData.SegmentPath = "icmpMsgStatsTable"
    icmpMsgStatsTable.EntityData.AbsolutePath = "IP-MIB:IP-MIB/" + icmpMsgStatsTable.EntityData.SegmentPath
    icmpMsgStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpMsgStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpMsgStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpMsgStatsTable.EntityData.Children = types.NewOrderedMap()
    icmpMsgStatsTable.EntityData.Children.Append("icmpMsgStatsEntry", types.YChild{"IcmpMsgStatsEntry", nil})
    for i := range icmpMsgStatsTable.IcmpMsgStatsEntry {
        icmpMsgStatsTable.EntityData.Children.Append(types.GetSegmentPath(icmpMsgStatsTable.IcmpMsgStatsEntry[i]), types.YChild{"IcmpMsgStatsEntry", icmpMsgStatsTable.IcmpMsgStatsEntry[i]})
    }
    icmpMsgStatsTable.EntityData.Leafs = types.NewOrderedMap()

    icmpMsgStatsTable.EntityData.YListKeys = []string {}

    return &(icmpMsgStatsTable.EntityData)
}

// IPMIB_IcmpMsgStatsTable_IcmpMsgStatsEntry
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
type IPMIB_IcmpMsgStatsTable_IcmpMsgStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP version of the statistics. The type is
    // IpVersion.
    IcmpMsgStatsIPVersion interface{}

    // This attribute is a key. The ICMP type field of the message type being
    // counted by this row.  Note that ICMP message types are scoped by the
    // address type in use. The type is interface{} with range: 0..255.
    IcmpMsgStatsType interface{}

    // The number of input packets for this AF and type. The type is interface{}
    // with range: 0..4294967295.
    IcmpMsgStatsInPkts interface{}

    // The number of output packets for this AF and type. The type is interface{}
    // with range: 0..4294967295.
    IcmpMsgStatsOutPkts interface{}
}

func (icmpMsgStatsEntry *IPMIB_IcmpMsgStatsTable_IcmpMsgStatsEntry) GetEntityData() *types.CommonEntityData {
    icmpMsgStatsEntry.EntityData.YFilter = icmpMsgStatsEntry.YFilter
    icmpMsgStatsEntry.EntityData.YangName = "icmpMsgStatsEntry"
    icmpMsgStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    icmpMsgStatsEntry.EntityData.ParentYangName = "icmpMsgStatsTable"
    icmpMsgStatsEntry.EntityData.SegmentPath = "icmpMsgStatsEntry" + types.AddKeyToken(icmpMsgStatsEntry.IcmpMsgStatsIPVersion, "icmpMsgStatsIPVersion") + types.AddKeyToken(icmpMsgStatsEntry.IcmpMsgStatsType, "icmpMsgStatsType")
    icmpMsgStatsEntry.EntityData.AbsolutePath = "IP-MIB:IP-MIB/icmpMsgStatsTable/" + icmpMsgStatsEntry.EntityData.SegmentPath
    icmpMsgStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpMsgStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpMsgStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpMsgStatsEntry.EntityData.Children = types.NewOrderedMap()
    icmpMsgStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    icmpMsgStatsEntry.EntityData.Leafs.Append("icmpMsgStatsIPVersion", types.YLeaf{"IcmpMsgStatsIPVersion", icmpMsgStatsEntry.IcmpMsgStatsIPVersion})
    icmpMsgStatsEntry.EntityData.Leafs.Append("icmpMsgStatsType", types.YLeaf{"IcmpMsgStatsType", icmpMsgStatsEntry.IcmpMsgStatsType})
    icmpMsgStatsEntry.EntityData.Leafs.Append("icmpMsgStatsInPkts", types.YLeaf{"IcmpMsgStatsInPkts", icmpMsgStatsEntry.IcmpMsgStatsInPkts})
    icmpMsgStatsEntry.EntityData.Leafs.Append("icmpMsgStatsOutPkts", types.YLeaf{"IcmpMsgStatsOutPkts", icmpMsgStatsEntry.IcmpMsgStatsOutPkts})

    icmpMsgStatsEntry.EntityData.YListKeys = []string {"IcmpMsgStatsIPVersion", "IcmpMsgStatsType"}

    return &(icmpMsgStatsEntry.EntityData)
}

