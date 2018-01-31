// The MIB module for management of IP Multicast routing, but
// independent of the specific multicast routing protocol in
// use.
package ipmroute_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipmroute_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:IPMROUTE-STD-MIB IPMROUTE-STD-MIB}", reflect.TypeOf(IPMROUTESTDMIB{}))
    ydk.RegisterEntity("IPMROUTE-STD-MIB:IPMROUTE-STD-MIB", reflect.TypeOf(IPMROUTESTDMIB{}))
}

// IPMROUTESTDMIB
type IPMROUTESTDMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ipmroute IPMROUTESTDMIB_Ipmroute

    // The (conceptual) table containing multicast routing information for IP
    // datagrams sent by particular sources to the IP multicast groups known to
    // this router.
    Ipmroutetable IPMROUTESTDMIB_Ipmroutetable

    // The (conceptual) table containing information on the next- hops on outgoing
    // interfaces for routing IP multicast datagrams.  Each entry is one of a list
    // of next-hops on outgoing interfaces for particular sources sending to a
    // particular multicast group address.
    Ipmroutenexthoptable IPMROUTESTDMIB_Ipmroutenexthoptable

    // The (conceptual) table containing multicast routing information specific to
    // interfaces.
    Ipmrouteinterfacetable IPMROUTESTDMIB_Ipmrouteinterfacetable

    // The (conceptual) table listing the router's scoped multicast address
    // boundaries.
    Ipmrouteboundarytable IPMROUTESTDMIB_Ipmrouteboundarytable

    // The (conceptual) table listing the multicast scope names.
    Ipmroutescopenametable IPMROUTESTDMIB_Ipmroutescopenametable
}

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetFilter() yfilter.YFilter { return iPMROUTESTDMIB.YFilter }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) SetFilter(yf yfilter.YFilter) { iPMROUTESTDMIB.YFilter = yf }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetGoName(yname string) string {
    if yname == "ipMRoute" { return "Ipmroute" }
    if yname == "ipMRouteTable" { return "Ipmroutetable" }
    if yname == "ipMRouteNextHopTable" { return "Ipmroutenexthoptable" }
    if yname == "ipMRouteInterfaceTable" { return "Ipmrouteinterfacetable" }
    if yname == "ipMRouteBoundaryTable" { return "Ipmrouteboundarytable" }
    if yname == "ipMRouteScopeNameTable" { return "Ipmroutescopenametable" }
    return ""
}

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetSegmentPath() string {
    return "IPMROUTE-STD-MIB:IPMROUTE-STD-MIB"
}

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipMRoute" {
        return &iPMROUTESTDMIB.Ipmroute
    }
    if childYangName == "ipMRouteTable" {
        return &iPMROUTESTDMIB.Ipmroutetable
    }
    if childYangName == "ipMRouteNextHopTable" {
        return &iPMROUTESTDMIB.Ipmroutenexthoptable
    }
    if childYangName == "ipMRouteInterfaceTable" {
        return &iPMROUTESTDMIB.Ipmrouteinterfacetable
    }
    if childYangName == "ipMRouteBoundaryTable" {
        return &iPMROUTESTDMIB.Ipmrouteboundarytable
    }
    if childYangName == "ipMRouteScopeNameTable" {
        return &iPMROUTESTDMIB.Ipmroutescopenametable
    }
    return nil
}

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipMRoute"] = &iPMROUTESTDMIB.Ipmroute
    children["ipMRouteTable"] = &iPMROUTESTDMIB.Ipmroutetable
    children["ipMRouteNextHopTable"] = &iPMROUTESTDMIB.Ipmroutenexthoptable
    children["ipMRouteInterfaceTable"] = &iPMROUTESTDMIB.Ipmrouteinterfacetable
    children["ipMRouteBoundaryTable"] = &iPMROUTESTDMIB.Ipmrouteboundarytable
    children["ipMRouteScopeNameTable"] = &iPMROUTESTDMIB.Ipmroutescopenametable
    return children
}

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetYangName() string { return "IPMROUTE-STD-MIB" }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) SetParent(parent types.Entity) { iPMROUTESTDMIB.parent = parent }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetParent() types.Entity { return iPMROUTESTDMIB.parent }

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmroute
type IPMROUTESTDMIB_Ipmroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The enabled status of IP Multicast routing on this router. The type is
    // Ipmrouteenable.
    Ipmrouteenable interface{}

    // The number of rows in the ipMRouteTable.  This can be used to monitor the
    // multicast routing table size. The type is interface{} with range:
    // 0..4294967295.
    Ipmrouteentrycount interface{}
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetFilter() yfilter.YFilter { return ipmroute.YFilter }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) SetFilter(yf yfilter.YFilter) { ipmroute.YFilter = yf }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetGoName(yname string) string {
    if yname == "ipMRouteEnable" { return "Ipmrouteenable" }
    if yname == "ipMRouteEntryCount" { return "Ipmrouteentrycount" }
    return ""
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetSegmentPath() string {
    return "ipMRoute"
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteEnable"] = ipmroute.Ipmrouteenable
    leafs["ipMRouteEntryCount"] = ipmroute.Ipmrouteentrycount
    return leafs
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetBundleName() string { return "cisco_ios_xe" }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetYangName() string { return "ipMRoute" }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) SetParent(parent types.Entity) { ipmroute.parent = parent }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetParent() types.Entity { return ipmroute.parent }

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmroute_Ipmrouteenable represents The enabled status of IP Multicast routing on this router.
type IPMROUTESTDMIB_Ipmroute_Ipmrouteenable string

const (
    IPMROUTESTDMIB_Ipmroute_Ipmrouteenable_enabled IPMROUTESTDMIB_Ipmroute_Ipmrouteenable = "enabled"

    IPMROUTESTDMIB_Ipmroute_Ipmrouteenable_disabled IPMROUTESTDMIB_Ipmroute_Ipmrouteenable = "disabled"
)

// IPMROUTESTDMIB_Ipmroutetable
// The (conceptual) table containing multicast routing
// information for IP datagrams sent by particular sources to
// the IP multicast groups known to this router.
type IPMROUTESTDMIB_Ipmroutetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the multicast routing information for
    // IP datagrams from a particular source and addressed to a particular IP
    // multicast group address. Discontinuities in counters in this entry can be
    // detected by observing the value of ipMRouteUpTime. The type is slice of
    // IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry.
    Ipmrouteentry []IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetFilter() yfilter.YFilter { return ipmroutetable.YFilter }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) SetFilter(yf yfilter.YFilter) { ipmroutetable.YFilter = yf }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetGoName(yname string) string {
    if yname == "ipMRouteEntry" { return "Ipmrouteentry" }
    return ""
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetSegmentPath() string {
    return "ipMRouteTable"
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipMRouteEntry" {
        for _, c := range ipmroutetable.Ipmrouteentry {
            if ipmroutetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry{}
        ipmroutetable.Ipmrouteentry = append(ipmroutetable.Ipmrouteentry, child)
        return &ipmroutetable.Ipmrouteentry[len(ipmroutetable.Ipmrouteentry)-1]
    }
    return nil
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipmroutetable.Ipmrouteentry {
        children[ipmroutetable.Ipmrouteentry[i].GetSegmentPath()] = &ipmroutetable.Ipmrouteentry[i]
    }
    return children
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetBundleName() string { return "cisco_ios_xe" }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetYangName() string { return "ipMRouteTable" }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) SetParent(parent types.Entity) { ipmroutetable.parent = parent }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetParent() types.Entity { return ipmroutetable.parent }

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry
// An entry (conceptual row) containing the multicast routing
// information for IP datagrams from a particular source and
// addressed to a particular IP multicast group address.
// Discontinuities in counters in this entry can be detected by
// observing the value of ipMRouteUpTime.
type IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains multicast routing information. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutegroup interface{}

    // This attribute is a key. The network address which when combined with the
    // corresponding value of ipMRouteSourceMask identifies the sources for which
    // this entry contains multicast routing information. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutesource interface{}

    // This attribute is a key. The network mask which when combined with the
    // corresponding value of ipMRouteSource identifies the sources for which this
    // entry contains multicast routing information. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutesourcemask interface{}

    // The address of the upstream neighbor (e.g., RPF neighbor) from which IP
    // datagrams from these sources to this multicast address are received, or
    // 0.0.0.0 if the upstream neighbor is unknown (e.g., in CBT). The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmrouteupstreamneighbor interface{}

    // The value of ifIndex for the interface on which IP datagrams sent by these
    // sources to this multicast address are received.  A value of 0 indicates
    // that datagrams are not subject to an incoming interface check, but may be
    // accepted on multiple interfaces (e.g., in CBT). The type is interface{}
    // with range: 0..2147483647.
    Ipmrouteinifindex interface{}

    // The time since the multicast routing information represented by this entry
    // was learned by the router. The type is interface{} with range:
    // 0..4294967295.
    Ipmrouteuptime interface{}

    // The minimum amount of time remaining before this entry will be aged out. 
    // The value 0 indicates that the entry is not subject to aging. The type is
    // interface{} with range: 0..4294967295.
    Ipmrouteexpirytime interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address. The type is interface{} with
    // range: 0..4294967295.
    Ipmroutepkts interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address, which were dropped because they
    // were not received on the interface indicated by ipMRouteInIfIndex.  Packets
    // which are not subject to an incoming interface check (e.g., using CBT) are
    // not counted. The type is interface{} with range: 0..4294967295.
    Ipmroutedifferentinifpackets interface{}

    // The number of octets contained in IP datagrams which were received from
    // these sources and addressed to this multicast group address, and which were
    // forwarded by this router. The type is interface{} with range:
    // 0..4294967295.
    Ipmrouteoctets interface{}

    // The multicast routing protocol via which this multicast forwarding entry
    // was learned. The type is IANAipMRouteProtocol.
    Ipmrouteprotocol interface{}

    // The routing mechanism via which the route used to find the upstream or
    // parent interface for this multicast forwarding entry was learned. 
    // Inclusion of values for routing protocols is not intended to imply that
    // those protocols need be supported. The type is IANAipRouteProtocol.
    Ipmroutertproto interface{}

    // The address portion of the route used to find the upstream or parent
    // interface for this multicast forwarding entry. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutertaddress interface{}

    // The mask associated with the route used to find the upstream or parent
    // interface for this multicast forwarding entry. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutertmask interface{}

    // The reason the given route was placed in the (logical) multicast Routing
    // Information Base (RIB).  A value of unicast means that the route would
    // normally be placed only in the unicast RIB, but was placed in the multicast
    // RIB (instead or in addition) due to local configuration, such as when
    // running PIM over RIP.  A value of multicast means that      the route was
    // explicitly added to the multicast RIB by the routing protocol, such as
    // DVMRP or Multiprotocol BGP. The type is Ipmrouterttype.
    Ipmrouterttype interface{}

    // The number of octets contained in IP datagrams which were received from
    // these sources and addressed to this multicast group address, and which were
    // forwarded by this router. This object is a 64-bit version of
    // ipMRouteOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipmroutehcoctets interface{}

    // Boolean, indicates whether this route is pruned. A pruned route is one that
    // has an empty outgoing interface list or all interfaces are in Pruned state.
    // A multicast packet that matches a pruned route doesn't get forwarded. The
    // type is bool.
    Ciscoipmroutepruneflag interface{}

    // Boolean, indicating PIM multicast routing protocol sparse-mode (versus
    // dense-mode).  In sparse-mode, packets are forwarded only out interfaces
    // that have been joined. In dense-mode, they are forwarded out all interfaces
    // that have not been pruned. The type is bool.
    Ciscoipmroutesparseflag interface{}

    // Boolean, indicating whether there is a directly connected member for a
    // group attached to the router. The type is bool.
    Ciscoipmrouteconnectedflag interface{}

    // Boolean, indicating whether local system is a member of a group on any
    // interface. The type is bool.
    Ciscoipmroutelocalflag interface{}

    // Boolean, indicates whether to send registers for the entry. A first hop
    // router directly connected to a multicast source host, as well as a border
    // router on the boundary of two domains running different multicast routing
    // protocols, encapsulates packets to be sent on the shared tree. This is done
    // until the RP sends Joins back to this router. The type is bool.
    Ciscoipmrouteregisterflag interface{}

    // Boolean, indicating whether there is a Prune state for this source along
    // the shared tree. The type is bool.
    Ciscoipmrouterpflag interface{}

    // Boolean, indicating whether data is being received on the SPT tree, ie the
    // Shortest Path Tree. The type is bool.
    Ciscoipmroutesptflag interface{}

    // Bits per second forwarded by this router.  This is the sum of all forwarded
    // bits during a 1 second interval.  At the end of each second the field is
    // cleared. This object has been superseded by ciscoIpMRouteBps2 (which is the
    // 64-bit version of this object). The type is interface{} with range:
    // 0..4294967295.
    Ciscoipmroutebps interface{}

    // Metric - The best metric heard from Assert messages. This object has been
    // replaced by ciscoIpMRouteMetric2 in order to correctly represent a 32-bit
    // unsigned metric value. The type is interface{} with range: 0..2147483647.
    Ciscoipmroutemetric interface{}

    // Metric Preference - The best metric preference heard from Assert messages.
    // The type is interface{} with range: 0..2147483647.
    Ciscoipmroutemetricpreference interface{}

    // Incoming interface's limit for rate limiting data traffic, in Kbps.
    // Replaced by ciscoIpMRouteInLimit2. The type is interface{} with range:
    // 0..2147483647. Units are Kbits/second.
    Ciscoipmrouteinlimit interface{}

    // How long has it been since the last multicast packet was fastswitched. The
    // type is interface{} with range: 0..4294967295.
    Ciscoipmroutelastused interface{}

    // Incoming interface's limit for rate limiting data traffic, in Kbps. The
    // type is interface{} with range: 0..4294967295. Units are Kbits/second.
    Ciscoipmrouteinlimit2 interface{}

    // Boolean, indicates whether this route is created due to SPT threshold. The
    // type is bool.
    Ciscoipmroutejoinflag interface{}

    // Boolean, indicates whether this route is learned via MSDP. The type is
    // bool.
    Ciscoipmroutemsdpflag interface{}

    // Boolean, indicates whether to send join for this entry. The type is bool.
    Ciscoipmrouteproxyjoinflag interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address. This object is a 64-bit version
    // of ipMRoutePkts. The type is interface{} with range:
    // 0..18446744073709551615.
    Ciscoipmroutepkts interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address, which were not received from the
    // interface indicated by ipMRouteInIfIndex. This object is a 64-bit version
    // of ipMRouteDifferentInIfPackets. The type is interface{} with range:
    // 0..18446744073709551615.
    Ciscoipmroutedifferentinifpkts interface{}

    // The number of octets contained in IP datagrams which were received from
    // these sources and addressed to this multicast group address, and which were
    // forwarded by this router. This object is a 64-bit version of
    // ipMRouteOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Ciscoipmrouteoctets interface{}

    // Bits per second forwarded by this router. This is the sum of all forwarded
    // bits during a 1 second interval. At the end of each second the field is
    // cleared. The type is interface{} with range: 0..18446744073709551615.
    Ciscoipmroutebps2 interface{}

    // Metric - The best metric heard from Assert messages. The type is
    // interface{} with range: 0..4294967295.
    Ciscoipmroutemetric2 interface{}
}

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetFilter() yfilter.YFilter { return ipmrouteentry.YFilter }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) SetFilter(yf yfilter.YFilter) { ipmrouteentry.YFilter = yf }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetGoName(yname string) string {
    if yname == "ipMRouteGroup" { return "Ipmroutegroup" }
    if yname == "ipMRouteSource" { return "Ipmroutesource" }
    if yname == "ipMRouteSourceMask" { return "Ipmroutesourcemask" }
    if yname == "ipMRouteUpstreamNeighbor" { return "Ipmrouteupstreamneighbor" }
    if yname == "ipMRouteInIfIndex" { return "Ipmrouteinifindex" }
    if yname == "ipMRouteUpTime" { return "Ipmrouteuptime" }
    if yname == "ipMRouteExpiryTime" { return "Ipmrouteexpirytime" }
    if yname == "ipMRoutePkts" { return "Ipmroutepkts" }
    if yname == "ipMRouteDifferentInIfPackets" { return "Ipmroutedifferentinifpackets" }
    if yname == "ipMRouteOctets" { return "Ipmrouteoctets" }
    if yname == "ipMRouteProtocol" { return "Ipmrouteprotocol" }
    if yname == "ipMRouteRtProto" { return "Ipmroutertproto" }
    if yname == "ipMRouteRtAddress" { return "Ipmroutertaddress" }
    if yname == "ipMRouteRtMask" { return "Ipmroutertmask" }
    if yname == "ipMRouteRtType" { return "Ipmrouterttype" }
    if yname == "ipMRouteHCOctets" { return "Ipmroutehcoctets" }
    if yname == "ciscoIpMRoutePruneFlag" { return "Ciscoipmroutepruneflag" }
    if yname == "ciscoIpMRouteSparseFlag" { return "Ciscoipmroutesparseflag" }
    if yname == "ciscoIpMRouteConnectedFlag" { return "Ciscoipmrouteconnectedflag" }
    if yname == "ciscoIpMRouteLocalFlag" { return "Ciscoipmroutelocalflag" }
    if yname == "ciscoIpMRouteRegisterFlag" { return "Ciscoipmrouteregisterflag" }
    if yname == "ciscoIpMRouteRpFlag" { return "Ciscoipmrouterpflag" }
    if yname == "ciscoIpMRouteSptFlag" { return "Ciscoipmroutesptflag" }
    if yname == "ciscoIpMRouteBps" { return "Ciscoipmroutebps" }
    if yname == "ciscoIpMRouteMetric" { return "Ciscoipmroutemetric" }
    if yname == "ciscoIpMRouteMetricPreference" { return "Ciscoipmroutemetricpreference" }
    if yname == "ciscoIpMRouteInLimit" { return "Ciscoipmrouteinlimit" }
    if yname == "ciscoIpMRouteLastUsed" { return "Ciscoipmroutelastused" }
    if yname == "ciscoIpMRouteInLimit2" { return "Ciscoipmrouteinlimit2" }
    if yname == "ciscoIpMRouteJoinFlag" { return "Ciscoipmroutejoinflag" }
    if yname == "ciscoIpMRouteMsdpFlag" { return "Ciscoipmroutemsdpflag" }
    if yname == "ciscoIpMRouteProxyJoinFlag" { return "Ciscoipmrouteproxyjoinflag" }
    if yname == "ciscoIpMRoutePkts" { return "Ciscoipmroutepkts" }
    if yname == "ciscoIpMRouteDifferentInIfPkts" { return "Ciscoipmroutedifferentinifpkts" }
    if yname == "ciscoIpMRouteOctets" { return "Ciscoipmrouteoctets" }
    if yname == "ciscoIpMRouteBps2" { return "Ciscoipmroutebps2" }
    if yname == "ciscoIpMRouteMetric2" { return "Ciscoipmroutemetric2" }
    return ""
}

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetSegmentPath() string {
    return "ipMRouteEntry" + "[ipMRouteGroup='" + fmt.Sprintf("%v", ipmrouteentry.Ipmroutegroup) + "']" + "[ipMRouteSource='" + fmt.Sprintf("%v", ipmrouteentry.Ipmroutesource) + "']" + "[ipMRouteSourceMask='" + fmt.Sprintf("%v", ipmrouteentry.Ipmroutesourcemask) + "']"
}

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteGroup"] = ipmrouteentry.Ipmroutegroup
    leafs["ipMRouteSource"] = ipmrouteentry.Ipmroutesource
    leafs["ipMRouteSourceMask"] = ipmrouteentry.Ipmroutesourcemask
    leafs["ipMRouteUpstreamNeighbor"] = ipmrouteentry.Ipmrouteupstreamneighbor
    leafs["ipMRouteInIfIndex"] = ipmrouteentry.Ipmrouteinifindex
    leafs["ipMRouteUpTime"] = ipmrouteentry.Ipmrouteuptime
    leafs["ipMRouteExpiryTime"] = ipmrouteentry.Ipmrouteexpirytime
    leafs["ipMRoutePkts"] = ipmrouteentry.Ipmroutepkts
    leafs["ipMRouteDifferentInIfPackets"] = ipmrouteentry.Ipmroutedifferentinifpackets
    leafs["ipMRouteOctets"] = ipmrouteentry.Ipmrouteoctets
    leafs["ipMRouteProtocol"] = ipmrouteentry.Ipmrouteprotocol
    leafs["ipMRouteRtProto"] = ipmrouteentry.Ipmroutertproto
    leafs["ipMRouteRtAddress"] = ipmrouteentry.Ipmroutertaddress
    leafs["ipMRouteRtMask"] = ipmrouteentry.Ipmroutertmask
    leafs["ipMRouteRtType"] = ipmrouteentry.Ipmrouterttype
    leafs["ipMRouteHCOctets"] = ipmrouteentry.Ipmroutehcoctets
    leafs["ciscoIpMRoutePruneFlag"] = ipmrouteentry.Ciscoipmroutepruneflag
    leafs["ciscoIpMRouteSparseFlag"] = ipmrouteentry.Ciscoipmroutesparseflag
    leafs["ciscoIpMRouteConnectedFlag"] = ipmrouteentry.Ciscoipmrouteconnectedflag
    leafs["ciscoIpMRouteLocalFlag"] = ipmrouteentry.Ciscoipmroutelocalflag
    leafs["ciscoIpMRouteRegisterFlag"] = ipmrouteentry.Ciscoipmrouteregisterflag
    leafs["ciscoIpMRouteRpFlag"] = ipmrouteentry.Ciscoipmrouterpflag
    leafs["ciscoIpMRouteSptFlag"] = ipmrouteentry.Ciscoipmroutesptflag
    leafs["ciscoIpMRouteBps"] = ipmrouteentry.Ciscoipmroutebps
    leafs["ciscoIpMRouteMetric"] = ipmrouteentry.Ciscoipmroutemetric
    leafs["ciscoIpMRouteMetricPreference"] = ipmrouteentry.Ciscoipmroutemetricpreference
    leafs["ciscoIpMRouteInLimit"] = ipmrouteentry.Ciscoipmrouteinlimit
    leafs["ciscoIpMRouteLastUsed"] = ipmrouteentry.Ciscoipmroutelastused
    leafs["ciscoIpMRouteInLimit2"] = ipmrouteentry.Ciscoipmrouteinlimit2
    leafs["ciscoIpMRouteJoinFlag"] = ipmrouteentry.Ciscoipmroutejoinflag
    leafs["ciscoIpMRouteMsdpFlag"] = ipmrouteentry.Ciscoipmroutemsdpflag
    leafs["ciscoIpMRouteProxyJoinFlag"] = ipmrouteentry.Ciscoipmrouteproxyjoinflag
    leafs["ciscoIpMRoutePkts"] = ipmrouteentry.Ciscoipmroutepkts
    leafs["ciscoIpMRouteDifferentInIfPkts"] = ipmrouteentry.Ciscoipmroutedifferentinifpkts
    leafs["ciscoIpMRouteOctets"] = ipmrouteentry.Ciscoipmrouteoctets
    leafs["ciscoIpMRouteBps2"] = ipmrouteentry.Ciscoipmroutebps2
    leafs["ciscoIpMRouteMetric2"] = ipmrouteentry.Ciscoipmroutemetric2
    return leafs
}

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetYangName() string { return "ipMRouteEntry" }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) SetParent(parent types.Entity) { ipmrouteentry.parent = parent }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetParent() types.Entity { return ipmrouteentry.parent }

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetParentYangName() string { return "ipMRouteTable" }

// IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmrouterttype represents routing protocol, such as DVMRP or Multiprotocol BGP.
type IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmrouterttype string

const (
    IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmrouterttype_unicast IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmrouterttype = "unicast"

    IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmrouterttype_multicast IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmrouterttype = "multicast"
)

// IPMROUTESTDMIB_Ipmroutenexthoptable
// The (conceptual) table containing information on the next-
// hops on outgoing interfaces for routing IP multicast
// datagrams.  Each entry is one of a list of next-hops on
// outgoing interfaces for particular sources sending to a
// particular multicast group address.
type IPMROUTESTDMIB_Ipmroutenexthoptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the list of next-hops on outgoing interfaces
    // to which IP multicast datagrams from particular sources to a IP multicast
    // group address are routed.  Discontinuities in counters in this entry can be
    // detected by observing the value of ipMRouteUpTime. The type is slice of
    // IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry.
    Ipmroutenexthopentry []IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetFilter() yfilter.YFilter { return ipmroutenexthoptable.YFilter }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) SetFilter(yf yfilter.YFilter) { ipmroutenexthoptable.YFilter = yf }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetGoName(yname string) string {
    if yname == "ipMRouteNextHopEntry" { return "Ipmroutenexthopentry" }
    return ""
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetSegmentPath() string {
    return "ipMRouteNextHopTable"
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipMRouteNextHopEntry" {
        for _, c := range ipmroutenexthoptable.Ipmroutenexthopentry {
            if ipmroutenexthoptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry{}
        ipmroutenexthoptable.Ipmroutenexthopentry = append(ipmroutenexthoptable.Ipmroutenexthopentry, child)
        return &ipmroutenexthoptable.Ipmroutenexthopentry[len(ipmroutenexthoptable.Ipmroutenexthopentry)-1]
    }
    return nil
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipmroutenexthoptable.Ipmroutenexthopentry {
        children[ipmroutenexthoptable.Ipmroutenexthopentry[i].GetSegmentPath()] = &ipmroutenexthoptable.Ipmroutenexthopentry[i]
    }
    return children
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetBundleName() string { return "cisco_ios_xe" }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetYangName() string { return "ipMRouteNextHopTable" }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) SetParent(parent types.Entity) { ipmroutenexthoptable.parent = parent }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetParent() types.Entity { return ipmroutenexthoptable.parent }

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry
// An entry (conceptual row) in the list of next-hops on
// outgoing interfaces to which IP multicast datagrams from
// particular sources to a IP multicast group address are
// routed.  Discontinuities in counters in this entry can be
// detected by observing the value of ipMRouteUpTime.
type IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group for which this entry
    // specifies a next-hop on an outgoing interface. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutenexthopgroup interface{}

    // This attribute is a key. The network address which when combined with the
    // corresponding value of ipMRouteNextHopSourceMask identifies the sources for
    // which this entry specifies a next-hop on an outgoing interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutenexthopsource interface{}

    // This attribute is a key. The network mask which when combined with the
    // corresponding value of ipMRouteNextHopSource identifies the sources for
    // which this entry specifies a next-hop on an outgoing interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutenexthopsourcemask interface{}

    // This attribute is a key. The ifIndex value of the interface for the
    // outgoing interface for this next-hop. The type is interface{} with range:
    // 1..2147483647.
    Ipmroutenexthopifindex interface{}

    // This attribute is a key. The address of the next-hop specific to this
    // entry.  For most interfaces, this is identical to ipMRouteNextHopGroup.
    // NBMA interfaces, however, may have multiple next-hop addresses out a single
    // outgoing interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutenexthopaddress interface{}

    // An indication of whether the outgoing interface and next- hop represented
    // by this entry is currently being used to forward IP datagrams.  The value
    // 'forwarding' indicates it is currently being used; the value 'pruned'
    // indicates it is not. The type is Ipmroutenexthopstate.
    Ipmroutenexthopstate interface{}

    // The time since the multicast routing information represented by this entry
    // was learned by the router. The type is interface{} with range:
    // 0..4294967295.
    Ipmroutenexthopuptime interface{}

    // The minimum amount of time remaining before this entry will be aged out. 
    // If ipMRouteNextHopState is pruned(1), the remaining time until the prune
    // expires and the state reverts to forwarding(2).  Otherwise, the remaining
    // time until this entry is removed from the table.  The time remaining may be
    // copied from ipMRouteExpiryTime if the protocol in use for this entry does
    // not specify next-hop timers.  The value 0      indicates that the entry is
    // not subject to aging. The type is interface{} with range: 0..4294967295.
    Ipmroutenexthopexpirytime interface{}

    // The minimum number of hops between this router and any member of this IP
    // multicast group reached via this next-hop on this outgoing interface.  Any
    // IP multicast datagrams for the group which have a TTL less than this number
    // of hops will not be forwarded to this next-hop. The type is interface{}
    // with range: -2147483648..2147483647.
    Ipmroutenexthopclosestmemberhops interface{}

    // The routing mechanism via which this next-hop was learned. The type is
    // IANAipMRouteProtocol.
    Ipmroutenexthopprotocol interface{}

    // The number of packets which have been forwarded using this route. The type
    // is interface{} with range: 0..4294967295.
    Ipmroutenexthoppkts interface{}

    // An outgoing interface's limit for rate limiting data traffic, in Kbps. The
    // type is interface{} with range: 0..4294967295. Units are Kbits/second.
    Ciscoipmroutenexthopoutlimit interface{}

    // The data link mac address header for a multicast datagram. Used by IP
    // multicast fastswitching. The type is string.
    Ciscoipmroutenexthopmachdr interface{}

    // The number of packets which have been forwarded using this route. This
    // object is a 64-bit version of ipMRouteNextHopPkts. The type is interface{}
    // with range: 0..18446744073709551615.
    Ciscoipmroutenexthoppkts interface{}
}

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetFilter() yfilter.YFilter { return ipmroutenexthopentry.YFilter }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) SetFilter(yf yfilter.YFilter) { ipmroutenexthopentry.YFilter = yf }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetGoName(yname string) string {
    if yname == "ipMRouteNextHopGroup" { return "Ipmroutenexthopgroup" }
    if yname == "ipMRouteNextHopSource" { return "Ipmroutenexthopsource" }
    if yname == "ipMRouteNextHopSourceMask" { return "Ipmroutenexthopsourcemask" }
    if yname == "ipMRouteNextHopIfIndex" { return "Ipmroutenexthopifindex" }
    if yname == "ipMRouteNextHopAddress" { return "Ipmroutenexthopaddress" }
    if yname == "ipMRouteNextHopState" { return "Ipmroutenexthopstate" }
    if yname == "ipMRouteNextHopUpTime" { return "Ipmroutenexthopuptime" }
    if yname == "ipMRouteNextHopExpiryTime" { return "Ipmroutenexthopexpirytime" }
    if yname == "ipMRouteNextHopClosestMemberHops" { return "Ipmroutenexthopclosestmemberhops" }
    if yname == "ipMRouteNextHopProtocol" { return "Ipmroutenexthopprotocol" }
    if yname == "ipMRouteNextHopPkts" { return "Ipmroutenexthoppkts" }
    if yname == "ciscoIpMRouteNextHopOutLimit" { return "Ciscoipmroutenexthopoutlimit" }
    if yname == "ciscoIpMRouteNextHopMacHdr" { return "Ciscoipmroutenexthopmachdr" }
    if yname == "ciscoIpMRouteNextHopPkts" { return "Ciscoipmroutenexthoppkts" }
    return ""
}

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetSegmentPath() string {
    return "ipMRouteNextHopEntry" + "[ipMRouteNextHopGroup='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopgroup) + "']" + "[ipMRouteNextHopSource='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopsource) + "']" + "[ipMRouteNextHopSourceMask='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopsourcemask) + "']" + "[ipMRouteNextHopIfIndex='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopifindex) + "']" + "[ipMRouteNextHopAddress='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopaddress) + "']"
}

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteNextHopGroup"] = ipmroutenexthopentry.Ipmroutenexthopgroup
    leafs["ipMRouteNextHopSource"] = ipmroutenexthopentry.Ipmroutenexthopsource
    leafs["ipMRouteNextHopSourceMask"] = ipmroutenexthopentry.Ipmroutenexthopsourcemask
    leafs["ipMRouteNextHopIfIndex"] = ipmroutenexthopentry.Ipmroutenexthopifindex
    leafs["ipMRouteNextHopAddress"] = ipmroutenexthopentry.Ipmroutenexthopaddress
    leafs["ipMRouteNextHopState"] = ipmroutenexthopentry.Ipmroutenexthopstate
    leafs["ipMRouteNextHopUpTime"] = ipmroutenexthopentry.Ipmroutenexthopuptime
    leafs["ipMRouteNextHopExpiryTime"] = ipmroutenexthopentry.Ipmroutenexthopexpirytime
    leafs["ipMRouteNextHopClosestMemberHops"] = ipmroutenexthopentry.Ipmroutenexthopclosestmemberhops
    leafs["ipMRouteNextHopProtocol"] = ipmroutenexthopentry.Ipmroutenexthopprotocol
    leafs["ipMRouteNextHopPkts"] = ipmroutenexthopentry.Ipmroutenexthoppkts
    leafs["ciscoIpMRouteNextHopOutLimit"] = ipmroutenexthopentry.Ciscoipmroutenexthopoutlimit
    leafs["ciscoIpMRouteNextHopMacHdr"] = ipmroutenexthopentry.Ciscoipmroutenexthopmachdr
    leafs["ciscoIpMRouteNextHopPkts"] = ipmroutenexthopentry.Ciscoipmroutenexthoppkts
    return leafs
}

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetYangName() string { return "ipMRouteNextHopEntry" }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) SetParent(parent types.Entity) { ipmroutenexthopentry.parent = parent }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetParent() types.Entity { return ipmroutenexthopentry.parent }

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetParentYangName() string { return "ipMRouteNextHopTable" }

// IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopstate represents not.
type IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopstate string

const (
    IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopstate_pruned IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopstate = "pruned"

    IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopstate_forwarding IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopstate = "forwarding"
)

// IPMROUTESTDMIB_Ipmrouteinterfacetable
// The (conceptual) table containing multicast routing
// information specific to interfaces.
type IPMROUTESTDMIB_Ipmrouteinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the multicast routing information for
    // a particular interface. The type is slice of
    // IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry.
    Ipmrouteinterfaceentry []IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetFilter() yfilter.YFilter { return ipmrouteinterfacetable.YFilter }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) SetFilter(yf yfilter.YFilter) { ipmrouteinterfacetable.YFilter = yf }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetGoName(yname string) string {
    if yname == "ipMRouteInterfaceEntry" { return "Ipmrouteinterfaceentry" }
    return ""
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetSegmentPath() string {
    return "ipMRouteInterfaceTable"
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipMRouteInterfaceEntry" {
        for _, c := range ipmrouteinterfacetable.Ipmrouteinterfaceentry {
            if ipmrouteinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry{}
        ipmrouteinterfacetable.Ipmrouteinterfaceentry = append(ipmrouteinterfacetable.Ipmrouteinterfaceentry, child)
        return &ipmrouteinterfacetable.Ipmrouteinterfaceentry[len(ipmrouteinterfacetable.Ipmrouteinterfaceentry)-1]
    }
    return nil
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipmrouteinterfacetable.Ipmrouteinterfaceentry {
        children[ipmrouteinterfacetable.Ipmrouteinterfaceentry[i].GetSegmentPath()] = &ipmrouteinterfacetable.Ipmrouteinterfaceentry[i]
    }
    return children
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetYangName() string { return "ipMRouteInterfaceTable" }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) SetParent(parent types.Entity) { ipmrouteinterfacetable.parent = parent }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetParent() types.Entity { return ipmrouteinterfacetable.parent }

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry
// An entry (conceptual row) containing the multicast routing
// information for a particular interface.
type IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the interface for which this
    // entry contains information. The type is interface{} with range:
    // 1..2147483647.
    Ipmrouteinterfaceifindex interface{}

    // The datagram TTL threshold for the interface. Any IP multicast datagrams
    // with a TTL less than this threshold will not be forwarded out the
    // interface. The default value of 0 means all multicast packets are forwarded
    // out the interface. The type is interface{} with range: 0..255.
    Ipmrouteinterfacettl interface{}

    // The routing protocol running on this interface. The type is
    // IANAipMRouteProtocol.
    Ipmrouteinterfaceprotocol interface{}

    // The rate-limit, in kilobits per second, of forwarded multicast traffic on
    // the interface.  A rate-limit of 0 indicates that no rate limiting is done.
    // The type is interface{} with range: -2147483648..2147483647.
    Ipmrouteinterfaceratelimit interface{}

    // The number of octets of multicast packets that have arrived on the
    // interface, including framing characters.  This object is similar to
    // ifInOctets in the Interfaces MIB, except that only multicast packets are
    // counted. The type is interface{} with range: 0..4294967295.
    Ipmrouteinterfaceinmcastoctets interface{}

    // The number of octets of multicast packets that have been sent on the
    // interface. The type is interface{} with range: 0..4294967295.
    Ipmrouteinterfaceoutmcastoctets interface{}

    // The number of octets of multicast packets that have arrived on the
    // interface, including framing characters.  This object is a 64-bit version
    // of ipMRouteInterfaceInMcastOctets.  It is similar to ifHCInOctets in the
    // Interfaces MIB, except that only multicast packets are counted. The type is
    // interface{} with range: 0..18446744073709551615.
    Ipmrouteinterfacehcinmcastoctets interface{}

    // The number of octets of multicast packets that have been      sent on the
    // interface.  This object is a 64-bit version of
    // ipMRouteInterfaceOutMcastOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipmrouteinterfacehcoutmcastoctets interface{}

    // The number of octets of multicast packets that have arrived on the
    // interface. This object is a 64-bit version of
    // ipMRouteInterfaceInMcastOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Ciscoipmrouteifinmcastoctets interface{}

    // The number of octets of multicast packets that have been sent on the
    // interface. This object is a 64-bit version of
    // ipMRouteInterfaceOutMcastOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Ciscoipmrouteifoutmcastoctets interface{}

    // The number of multicast packets that have arrived on the interface. The
    // type is interface{} with range: 0..4294967295.
    Ciscoipmrouteifinmcastpkts interface{}

    // The number of multicast packets that have arrived on the interface. This
    // object is a 64-bit version of ciscoIpMRouteIfInMcastPkts. The type is
    // interface{} with range: 0..18446744073709551615.
    Ciscoipmrouteifhcinmcastpkts interface{}

    // The number of multicast packets that have been sent on the interface. The
    // type is interface{} with range: 0..4294967295.
    Ciscoipmrouteifoutmcastpkts interface{}

    // The number of multicast packets that have been sent on the interface. This
    // object is a 64-bit version of ciscoIpMRouteIfOutMcastPkts. The type is
    // interface{} with range: 0..18446744073709551615.
    Ciscoipmrouteifhcoutmcastpkts interface{}
}

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetFilter() yfilter.YFilter { return ipmrouteinterfaceentry.YFilter }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) SetFilter(yf yfilter.YFilter) { ipmrouteinterfaceentry.YFilter = yf }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetGoName(yname string) string {
    if yname == "ipMRouteInterfaceIfIndex" { return "Ipmrouteinterfaceifindex" }
    if yname == "ipMRouteInterfaceTtl" { return "Ipmrouteinterfacettl" }
    if yname == "ipMRouteInterfaceProtocol" { return "Ipmrouteinterfaceprotocol" }
    if yname == "ipMRouteInterfaceRateLimit" { return "Ipmrouteinterfaceratelimit" }
    if yname == "ipMRouteInterfaceInMcastOctets" { return "Ipmrouteinterfaceinmcastoctets" }
    if yname == "ipMRouteInterfaceOutMcastOctets" { return "Ipmrouteinterfaceoutmcastoctets" }
    if yname == "ipMRouteInterfaceHCInMcastOctets" { return "Ipmrouteinterfacehcinmcastoctets" }
    if yname == "ipMRouteInterfaceHCOutMcastOctets" { return "Ipmrouteinterfacehcoutmcastoctets" }
    if yname == "ciscoIpMRouteIfInMcastOctets" { return "Ciscoipmrouteifinmcastoctets" }
    if yname == "ciscoIpMRouteIfOutMcastOctets" { return "Ciscoipmrouteifoutmcastoctets" }
    if yname == "ciscoIpMRouteIfInMcastPkts" { return "Ciscoipmrouteifinmcastpkts" }
    if yname == "ciscoIpMRouteIfHCInMcastPkts" { return "Ciscoipmrouteifhcinmcastpkts" }
    if yname == "ciscoIpMRouteIfOutMcastPkts" { return "Ciscoipmrouteifoutmcastpkts" }
    if yname == "ciscoIpMRouteIfHCOutMcastPkts" { return "Ciscoipmrouteifhcoutmcastpkts" }
    return ""
}

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetSegmentPath() string {
    return "ipMRouteInterfaceEntry" + "[ipMRouteInterfaceIfIndex='" + fmt.Sprintf("%v", ipmrouteinterfaceentry.Ipmrouteinterfaceifindex) + "']"
}

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteInterfaceIfIndex"] = ipmrouteinterfaceentry.Ipmrouteinterfaceifindex
    leafs["ipMRouteInterfaceTtl"] = ipmrouteinterfaceentry.Ipmrouteinterfacettl
    leafs["ipMRouteInterfaceProtocol"] = ipmrouteinterfaceentry.Ipmrouteinterfaceprotocol
    leafs["ipMRouteInterfaceRateLimit"] = ipmrouteinterfaceentry.Ipmrouteinterfaceratelimit
    leafs["ipMRouteInterfaceInMcastOctets"] = ipmrouteinterfaceentry.Ipmrouteinterfaceinmcastoctets
    leafs["ipMRouteInterfaceOutMcastOctets"] = ipmrouteinterfaceentry.Ipmrouteinterfaceoutmcastoctets
    leafs["ipMRouteInterfaceHCInMcastOctets"] = ipmrouteinterfaceentry.Ipmrouteinterfacehcinmcastoctets
    leafs["ipMRouteInterfaceHCOutMcastOctets"] = ipmrouteinterfaceentry.Ipmrouteinterfacehcoutmcastoctets
    leafs["ciscoIpMRouteIfInMcastOctets"] = ipmrouteinterfaceentry.Ciscoipmrouteifinmcastoctets
    leafs["ciscoIpMRouteIfOutMcastOctets"] = ipmrouteinterfaceentry.Ciscoipmrouteifoutmcastoctets
    leafs["ciscoIpMRouteIfInMcastPkts"] = ipmrouteinterfaceentry.Ciscoipmrouteifinmcastpkts
    leafs["ciscoIpMRouteIfHCInMcastPkts"] = ipmrouteinterfaceentry.Ciscoipmrouteifhcinmcastpkts
    leafs["ciscoIpMRouteIfOutMcastPkts"] = ipmrouteinterfaceentry.Ciscoipmrouteifoutmcastpkts
    leafs["ciscoIpMRouteIfHCOutMcastPkts"] = ipmrouteinterfaceentry.Ciscoipmrouteifhcoutmcastpkts
    return leafs
}

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetYangName() string { return "ipMRouteInterfaceEntry" }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) SetParent(parent types.Entity) { ipmrouteinterfaceentry.parent = parent }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetParent() types.Entity { return ipmrouteinterfaceentry.parent }

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetParentYangName() string { return "ipMRouteInterfaceTable" }

// IPMROUTESTDMIB_Ipmrouteboundarytable
// The (conceptual) table listing the router's scoped
// multicast address boundaries.
type IPMROUTESTDMIB_Ipmrouteboundarytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ipMRouteBoundaryTable representing a
    // scoped boundary. The type is slice of
    // IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry.
    Ipmrouteboundaryentry []IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetFilter() yfilter.YFilter { return ipmrouteboundarytable.YFilter }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) SetFilter(yf yfilter.YFilter) { ipmrouteboundarytable.YFilter = yf }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetGoName(yname string) string {
    if yname == "ipMRouteBoundaryEntry" { return "Ipmrouteboundaryentry" }
    return ""
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetSegmentPath() string {
    return "ipMRouteBoundaryTable"
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipMRouteBoundaryEntry" {
        for _, c := range ipmrouteboundarytable.Ipmrouteboundaryentry {
            if ipmrouteboundarytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry{}
        ipmrouteboundarytable.Ipmrouteboundaryentry = append(ipmrouteboundarytable.Ipmrouteboundaryentry, child)
        return &ipmrouteboundarytable.Ipmrouteboundaryentry[len(ipmrouteboundarytable.Ipmrouteboundaryentry)-1]
    }
    return nil
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipmrouteboundarytable.Ipmrouteboundaryentry {
        children[ipmrouteboundarytable.Ipmrouteboundaryentry[i].GetSegmentPath()] = &ipmrouteboundarytable.Ipmrouteboundaryentry[i]
    }
    return children
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetBundleName() string { return "cisco_ios_xe" }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetYangName() string { return "ipMRouteBoundaryTable" }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) SetParent(parent types.Entity) { ipmrouteboundarytable.parent = parent }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetParent() types.Entity { return ipmrouteboundarytable.parent }

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry
// An entry (conceptual row) in the ipMRouteBoundaryTable
// representing a scoped boundary.
type IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IfIndex value for the interface to which this
    // boundary applies.  Packets with a destination address in the associated
    // address/mask range will not be forwarded out this interface. The type is
    // interface{} with range: 1..2147483647.
    Ipmrouteboundaryifindex interface{}

    // This attribute is a key. The group address which when combined with the
    // corresponding value of ipMRouteBoundaryAddressMask identifies the group
    // range for which the scoped boundary exists.  Scoped addresses must come
    // from the range 239.x.x.x as specified in RFC 2365. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmrouteboundaryaddress interface{}

    // This attribute is a key. The group address mask which when combined with
    // the corresponding value of ipMRouteBoundaryAddress identifies the group
    // range for which the scoped boundary exists. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmrouteboundaryaddressmask interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    Ipmrouteboundarystatus interface{}
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetFilter() yfilter.YFilter { return ipmrouteboundaryentry.YFilter }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) SetFilter(yf yfilter.YFilter) { ipmrouteboundaryentry.YFilter = yf }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetGoName(yname string) string {
    if yname == "ipMRouteBoundaryIfIndex" { return "Ipmrouteboundaryifindex" }
    if yname == "ipMRouteBoundaryAddress" { return "Ipmrouteboundaryaddress" }
    if yname == "ipMRouteBoundaryAddressMask" { return "Ipmrouteboundaryaddressmask" }
    if yname == "ipMRouteBoundaryStatus" { return "Ipmrouteboundarystatus" }
    return ""
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetSegmentPath() string {
    return "ipMRouteBoundaryEntry" + "[ipMRouteBoundaryIfIndex='" + fmt.Sprintf("%v", ipmrouteboundaryentry.Ipmrouteboundaryifindex) + "']" + "[ipMRouteBoundaryAddress='" + fmt.Sprintf("%v", ipmrouteboundaryentry.Ipmrouteboundaryaddress) + "']" + "[ipMRouteBoundaryAddressMask='" + fmt.Sprintf("%v", ipmrouteboundaryentry.Ipmrouteboundaryaddressmask) + "']"
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteBoundaryIfIndex"] = ipmrouteboundaryentry.Ipmrouteboundaryifindex
    leafs["ipMRouteBoundaryAddress"] = ipmrouteboundaryentry.Ipmrouteboundaryaddress
    leafs["ipMRouteBoundaryAddressMask"] = ipmrouteboundaryentry.Ipmrouteboundaryaddressmask
    leafs["ipMRouteBoundaryStatus"] = ipmrouteboundaryentry.Ipmrouteboundarystatus
    return leafs
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetYangName() string { return "ipMRouteBoundaryEntry" }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) SetParent(parent types.Entity) { ipmrouteboundaryentry.parent = parent }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetParent() types.Entity { return ipmrouteboundaryentry.parent }

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetParentYangName() string { return "ipMRouteBoundaryTable" }

// IPMROUTESTDMIB_Ipmroutescopenametable
// The (conceptual) table listing the multicast scope names.
type IPMROUTESTDMIB_Ipmroutescopenametable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ipMRouteScopeNameTable representing a
    // multicast scope name. The type is slice of
    // IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry.
    Ipmroutescopenameentry []IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetFilter() yfilter.YFilter { return ipmroutescopenametable.YFilter }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) SetFilter(yf yfilter.YFilter) { ipmroutescopenametable.YFilter = yf }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetGoName(yname string) string {
    if yname == "ipMRouteScopeNameEntry" { return "Ipmroutescopenameentry" }
    return ""
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetSegmentPath() string {
    return "ipMRouteScopeNameTable"
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipMRouteScopeNameEntry" {
        for _, c := range ipmroutescopenametable.Ipmroutescopenameentry {
            if ipmroutescopenametable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry{}
        ipmroutescopenametable.Ipmroutescopenameentry = append(ipmroutescopenametable.Ipmroutescopenameentry, child)
        return &ipmroutescopenametable.Ipmroutescopenameentry[len(ipmroutescopenametable.Ipmroutescopenameentry)-1]
    }
    return nil
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipmroutescopenametable.Ipmroutescopenameentry {
        children[ipmroutescopenametable.Ipmroutescopenameentry[i].GetSegmentPath()] = &ipmroutescopenametable.Ipmroutescopenameentry[i]
    }
    return children
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetBundleName() string { return "cisco_ios_xe" }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetYangName() string { return "ipMRouteScopeNameTable" }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) SetParent(parent types.Entity) { ipmroutescopenametable.parent = parent }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetParent() types.Entity { return ipmroutescopenametable.parent }

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetParentYangName() string { return "IPMROUTE-STD-MIB" }

// IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry
// An entry (conceptual row) in the ipMRouteScopeNameTable
// representing a multicast scope name.
type IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The group address which when combined with the
    // corresponding value of ipMRouteScopeNameAddressMask identifies the group
    // range associated with the multicast scope.  Scoped addresses must come from
    // the range 239.x.x.x. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutescopenameaddress interface{}

    // This attribute is a key. The group address mask which when combined with
    // the corresponding value of ipMRouteScopeNameAddress identifies the group
    // range associated with the multicast scope. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipmroutescopenameaddressmask interface{}

    // This attribute is a key. The RFC 1766-style language tag associated with
    // the scope name. The type is string.
    Ipmroutescopenamelanguage interface{}

    // The textual name associated with the multicast scope.  The value of this
    // object should be suitable for displaying to end-users, such as when
    // allocating a multicast address in this scope.  When no name is specified,
    // the default value of this object should be the string 239.x.x.x/y with x
    // and y replaced appropriately to describe the address and mask length
    // associated with the scope. The type is string.
    Ipmroutescopenamestring interface{}

    // If true, indicates a preference that the name in the following language
    // should be used by applications if no name is available in a desired
    // language. The type is bool.
    Ipmroutescopenamedefault interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    Ipmroutescopenamestatus interface{}
}

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetFilter() yfilter.YFilter { return ipmroutescopenameentry.YFilter }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) SetFilter(yf yfilter.YFilter) { ipmroutescopenameentry.YFilter = yf }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetGoName(yname string) string {
    if yname == "ipMRouteScopeNameAddress" { return "Ipmroutescopenameaddress" }
    if yname == "ipMRouteScopeNameAddressMask" { return "Ipmroutescopenameaddressmask" }
    if yname == "ipMRouteScopeNameLanguage" { return "Ipmroutescopenamelanguage" }
    if yname == "ipMRouteScopeNameString" { return "Ipmroutescopenamestring" }
    if yname == "ipMRouteScopeNameDefault" { return "Ipmroutescopenamedefault" }
    if yname == "ipMRouteScopeNameStatus" { return "Ipmroutescopenamestatus" }
    return ""
}

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetSegmentPath() string {
    return "ipMRouteScopeNameEntry" + "[ipMRouteScopeNameAddress='" + fmt.Sprintf("%v", ipmroutescopenameentry.Ipmroutescopenameaddress) + "']" + "[ipMRouteScopeNameAddressMask='" + fmt.Sprintf("%v", ipmroutescopenameentry.Ipmroutescopenameaddressmask) + "']" + "[ipMRouteScopeNameLanguage='" + fmt.Sprintf("%v", ipmroutescopenameentry.Ipmroutescopenamelanguage) + "']"
}

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteScopeNameAddress"] = ipmroutescopenameentry.Ipmroutescopenameaddress
    leafs["ipMRouteScopeNameAddressMask"] = ipmroutescopenameentry.Ipmroutescopenameaddressmask
    leafs["ipMRouteScopeNameLanguage"] = ipmroutescopenameentry.Ipmroutescopenamelanguage
    leafs["ipMRouteScopeNameString"] = ipmroutescopenameentry.Ipmroutescopenamestring
    leafs["ipMRouteScopeNameDefault"] = ipmroutescopenameentry.Ipmroutescopenamedefault
    leafs["ipMRouteScopeNameStatus"] = ipmroutescopenameentry.Ipmroutescopenamestatus
    return leafs
}

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetYangName() string { return "ipMRouteScopeNameEntry" }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) SetParent(parent types.Entity) { ipmroutescopenameentry.parent = parent }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetParent() types.Entity { return ipmroutescopenameentry.parent }

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetParentYangName() string { return "ipMRouteScopeNameTable" }

