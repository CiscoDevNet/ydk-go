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
    EntityData types.CommonEntityData
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

func (iPMROUTESTDMIB *IPMROUTESTDMIB) GetEntityData() *types.CommonEntityData {
    iPMROUTESTDMIB.EntityData.YFilter = iPMROUTESTDMIB.YFilter
    iPMROUTESTDMIB.EntityData.YangName = "IPMROUTE-STD-MIB"
    iPMROUTESTDMIB.EntityData.BundleName = "cisco_ios_xe"
    iPMROUTESTDMIB.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    iPMROUTESTDMIB.EntityData.SegmentPath = "IPMROUTE-STD-MIB:IPMROUTE-STD-MIB"
    iPMROUTESTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iPMROUTESTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iPMROUTESTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iPMROUTESTDMIB.EntityData.Children = make(map[string]types.YChild)
    iPMROUTESTDMIB.EntityData.Children["ipMRoute"] = types.YChild{"Ipmroute", &iPMROUTESTDMIB.Ipmroute}
    iPMROUTESTDMIB.EntityData.Children["ipMRouteTable"] = types.YChild{"Ipmroutetable", &iPMROUTESTDMIB.Ipmroutetable}
    iPMROUTESTDMIB.EntityData.Children["ipMRouteNextHopTable"] = types.YChild{"Ipmroutenexthoptable", &iPMROUTESTDMIB.Ipmroutenexthoptable}
    iPMROUTESTDMIB.EntityData.Children["ipMRouteInterfaceTable"] = types.YChild{"Ipmrouteinterfacetable", &iPMROUTESTDMIB.Ipmrouteinterfacetable}
    iPMROUTESTDMIB.EntityData.Children["ipMRouteBoundaryTable"] = types.YChild{"Ipmrouteboundarytable", &iPMROUTESTDMIB.Ipmrouteboundarytable}
    iPMROUTESTDMIB.EntityData.Children["ipMRouteScopeNameTable"] = types.YChild{"Ipmroutescopenametable", &iPMROUTESTDMIB.Ipmroutescopenametable}
    iPMROUTESTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iPMROUTESTDMIB.EntityData)
}

// IPMROUTESTDMIB_Ipmroute
type IPMROUTESTDMIB_Ipmroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The enabled status of IP Multicast routing on this router. The type is
    // Ipmrouteenable.
    Ipmrouteenable interface{}

    // The number of rows in the ipMRouteTable.  This can be used to monitor the
    // multicast routing table size. The type is interface{} with range:
    // 0..4294967295.
    Ipmrouteentrycount interface{}
}

func (ipmroute *IPMROUTESTDMIB_Ipmroute) GetEntityData() *types.CommonEntityData {
    ipmroute.EntityData.YFilter = ipmroute.YFilter
    ipmroute.EntityData.YangName = "ipMRoute"
    ipmroute.EntityData.BundleName = "cisco_ios_xe"
    ipmroute.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipmroute.EntityData.SegmentPath = "ipMRoute"
    ipmroute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmroute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmroute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmroute.EntityData.Children = make(map[string]types.YChild)
    ipmroute.EntityData.Leafs = make(map[string]types.YLeaf)
    ipmroute.EntityData.Leafs["ipMRouteEnable"] = types.YLeaf{"Ipmrouteenable", ipmroute.Ipmrouteenable}
    ipmroute.EntityData.Leafs["ipMRouteEntryCount"] = types.YLeaf{"Ipmrouteentrycount", ipmroute.Ipmrouteentrycount}
    return &(ipmroute.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the multicast routing information for
    // IP datagrams from a particular source and addressed to a particular IP
    // multicast group address. Discontinuities in counters in this entry can be
    // detected by observing the value of ipMRouteUpTime. The type is slice of
    // IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry.
    Ipmrouteentry []IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry
}

func (ipmroutetable *IPMROUTESTDMIB_Ipmroutetable) GetEntityData() *types.CommonEntityData {
    ipmroutetable.EntityData.YFilter = ipmroutetable.YFilter
    ipmroutetable.EntityData.YangName = "ipMRouteTable"
    ipmroutetable.EntityData.BundleName = "cisco_ios_xe"
    ipmroutetable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipmroutetable.EntityData.SegmentPath = "ipMRouteTable"
    ipmroutetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmroutetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmroutetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmroutetable.EntityData.Children = make(map[string]types.YChild)
    ipmroutetable.EntityData.Children["ipMRouteEntry"] = types.YChild{"Ipmrouteentry", nil}
    for i := range ipmroutetable.Ipmrouteentry {
        ipmroutetable.EntityData.Children[types.GetSegmentPath(&ipmroutetable.Ipmrouteentry[i])] = types.YChild{"Ipmrouteentry", &ipmroutetable.Ipmrouteentry[i]}
    }
    ipmroutetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipmroutetable.EntityData)
}

// IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry
// An entry (conceptual row) containing the multicast routing
// information for IP datagrams from a particular source and
// addressed to a particular IP multicast group address.
// Discontinuities in counters in this entry can be detected by
// observing the value of ipMRouteUpTime.
type IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains multicast routing information. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutegroup interface{}

    // This attribute is a key. The network address which when combined with the
    // corresponding value of ipMRouteSourceMask identifies the sources for which
    // this entry contains multicast routing information. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutesource interface{}

    // This attribute is a key. The network mask which when combined with the
    // corresponding value of ipMRouteSource identifies the sources for which this
    // entry contains multicast routing information. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutesourcemask interface{}

    // The address of the upstream neighbor (e.g., RPF neighbor) from which IP
    // datagrams from these sources to this multicast address are received, or
    // 0.0.0.0 if the upstream neighbor is unknown (e.g., in CBT). The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutertaddress interface{}

    // The mask associated with the route used to find the upstream or parent
    // interface for this multicast forwarding entry. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipmrouteentry *IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry) GetEntityData() *types.CommonEntityData {
    ipmrouteentry.EntityData.YFilter = ipmrouteentry.YFilter
    ipmrouteentry.EntityData.YangName = "ipMRouteEntry"
    ipmrouteentry.EntityData.BundleName = "cisco_ios_xe"
    ipmrouteentry.EntityData.ParentYangName = "ipMRouteTable"
    ipmrouteentry.EntityData.SegmentPath = "ipMRouteEntry" + "[ipMRouteGroup='" + fmt.Sprintf("%v", ipmrouteentry.Ipmroutegroup) + "']" + "[ipMRouteSource='" + fmt.Sprintf("%v", ipmrouteentry.Ipmroutesource) + "']" + "[ipMRouteSourceMask='" + fmt.Sprintf("%v", ipmrouteentry.Ipmroutesourcemask) + "']"
    ipmrouteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmrouteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmrouteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmrouteentry.EntityData.Children = make(map[string]types.YChild)
    ipmrouteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipmrouteentry.EntityData.Leafs["ipMRouteGroup"] = types.YLeaf{"Ipmroutegroup", ipmrouteentry.Ipmroutegroup}
    ipmrouteentry.EntityData.Leafs["ipMRouteSource"] = types.YLeaf{"Ipmroutesource", ipmrouteentry.Ipmroutesource}
    ipmrouteentry.EntityData.Leafs["ipMRouteSourceMask"] = types.YLeaf{"Ipmroutesourcemask", ipmrouteentry.Ipmroutesourcemask}
    ipmrouteentry.EntityData.Leafs["ipMRouteUpstreamNeighbor"] = types.YLeaf{"Ipmrouteupstreamneighbor", ipmrouteentry.Ipmrouteupstreamneighbor}
    ipmrouteentry.EntityData.Leafs["ipMRouteInIfIndex"] = types.YLeaf{"Ipmrouteinifindex", ipmrouteentry.Ipmrouteinifindex}
    ipmrouteentry.EntityData.Leafs["ipMRouteUpTime"] = types.YLeaf{"Ipmrouteuptime", ipmrouteentry.Ipmrouteuptime}
    ipmrouteentry.EntityData.Leafs["ipMRouteExpiryTime"] = types.YLeaf{"Ipmrouteexpirytime", ipmrouteentry.Ipmrouteexpirytime}
    ipmrouteentry.EntityData.Leafs["ipMRoutePkts"] = types.YLeaf{"Ipmroutepkts", ipmrouteentry.Ipmroutepkts}
    ipmrouteentry.EntityData.Leafs["ipMRouteDifferentInIfPackets"] = types.YLeaf{"Ipmroutedifferentinifpackets", ipmrouteentry.Ipmroutedifferentinifpackets}
    ipmrouteentry.EntityData.Leafs["ipMRouteOctets"] = types.YLeaf{"Ipmrouteoctets", ipmrouteentry.Ipmrouteoctets}
    ipmrouteentry.EntityData.Leafs["ipMRouteProtocol"] = types.YLeaf{"Ipmrouteprotocol", ipmrouteentry.Ipmrouteprotocol}
    ipmrouteentry.EntityData.Leafs["ipMRouteRtProto"] = types.YLeaf{"Ipmroutertproto", ipmrouteentry.Ipmroutertproto}
    ipmrouteentry.EntityData.Leafs["ipMRouteRtAddress"] = types.YLeaf{"Ipmroutertaddress", ipmrouteentry.Ipmroutertaddress}
    ipmrouteentry.EntityData.Leafs["ipMRouteRtMask"] = types.YLeaf{"Ipmroutertmask", ipmrouteentry.Ipmroutertmask}
    ipmrouteentry.EntityData.Leafs["ipMRouteRtType"] = types.YLeaf{"Ipmrouterttype", ipmrouteentry.Ipmrouterttype}
    ipmrouteentry.EntityData.Leafs["ipMRouteHCOctets"] = types.YLeaf{"Ipmroutehcoctets", ipmrouteentry.Ipmroutehcoctets}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRoutePruneFlag"] = types.YLeaf{"Ciscoipmroutepruneflag", ipmrouteentry.Ciscoipmroutepruneflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteSparseFlag"] = types.YLeaf{"Ciscoipmroutesparseflag", ipmrouteentry.Ciscoipmroutesparseflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteConnectedFlag"] = types.YLeaf{"Ciscoipmrouteconnectedflag", ipmrouteentry.Ciscoipmrouteconnectedflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteLocalFlag"] = types.YLeaf{"Ciscoipmroutelocalflag", ipmrouteentry.Ciscoipmroutelocalflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteRegisterFlag"] = types.YLeaf{"Ciscoipmrouteregisterflag", ipmrouteentry.Ciscoipmrouteregisterflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteRpFlag"] = types.YLeaf{"Ciscoipmrouterpflag", ipmrouteentry.Ciscoipmrouterpflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteSptFlag"] = types.YLeaf{"Ciscoipmroutesptflag", ipmrouteentry.Ciscoipmroutesptflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteBps"] = types.YLeaf{"Ciscoipmroutebps", ipmrouteentry.Ciscoipmroutebps}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteMetric"] = types.YLeaf{"Ciscoipmroutemetric", ipmrouteentry.Ciscoipmroutemetric}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteMetricPreference"] = types.YLeaf{"Ciscoipmroutemetricpreference", ipmrouteentry.Ciscoipmroutemetricpreference}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteInLimit"] = types.YLeaf{"Ciscoipmrouteinlimit", ipmrouteentry.Ciscoipmrouteinlimit}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteLastUsed"] = types.YLeaf{"Ciscoipmroutelastused", ipmrouteentry.Ciscoipmroutelastused}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteInLimit2"] = types.YLeaf{"Ciscoipmrouteinlimit2", ipmrouteentry.Ciscoipmrouteinlimit2}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteJoinFlag"] = types.YLeaf{"Ciscoipmroutejoinflag", ipmrouteentry.Ciscoipmroutejoinflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteMsdpFlag"] = types.YLeaf{"Ciscoipmroutemsdpflag", ipmrouteentry.Ciscoipmroutemsdpflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteProxyJoinFlag"] = types.YLeaf{"Ciscoipmrouteproxyjoinflag", ipmrouteentry.Ciscoipmrouteproxyjoinflag}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRoutePkts"] = types.YLeaf{"Ciscoipmroutepkts", ipmrouteentry.Ciscoipmroutepkts}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteDifferentInIfPkts"] = types.YLeaf{"Ciscoipmroutedifferentinifpkts", ipmrouteentry.Ciscoipmroutedifferentinifpkts}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteOctets"] = types.YLeaf{"Ciscoipmrouteoctets", ipmrouteentry.Ciscoipmrouteoctets}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteBps2"] = types.YLeaf{"Ciscoipmroutebps2", ipmrouteentry.Ciscoipmroutebps2}
    ipmrouteentry.EntityData.Leafs["ciscoIpMRouteMetric2"] = types.YLeaf{"Ciscoipmroutemetric2", ipmrouteentry.Ciscoipmroutemetric2}
    return &(ipmrouteentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the list of next-hops on outgoing interfaces
    // to which IP multicast datagrams from particular sources to a IP multicast
    // group address are routed.  Discontinuities in counters in this entry can be
    // detected by observing the value of ipMRouteUpTime. The type is slice of
    // IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry.
    Ipmroutenexthopentry []IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry
}

func (ipmroutenexthoptable *IPMROUTESTDMIB_Ipmroutenexthoptable) GetEntityData() *types.CommonEntityData {
    ipmroutenexthoptable.EntityData.YFilter = ipmroutenexthoptable.YFilter
    ipmroutenexthoptable.EntityData.YangName = "ipMRouteNextHopTable"
    ipmroutenexthoptable.EntityData.BundleName = "cisco_ios_xe"
    ipmroutenexthoptable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipmroutenexthoptable.EntityData.SegmentPath = "ipMRouteNextHopTable"
    ipmroutenexthoptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmroutenexthoptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmroutenexthoptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmroutenexthoptable.EntityData.Children = make(map[string]types.YChild)
    ipmroutenexthoptable.EntityData.Children["ipMRouteNextHopEntry"] = types.YChild{"Ipmroutenexthopentry", nil}
    for i := range ipmroutenexthoptable.Ipmroutenexthopentry {
        ipmroutenexthoptable.EntityData.Children[types.GetSegmentPath(&ipmroutenexthoptable.Ipmroutenexthopentry[i])] = types.YChild{"Ipmroutenexthopentry", &ipmroutenexthoptable.Ipmroutenexthopentry[i]}
    }
    ipmroutenexthoptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipmroutenexthoptable.EntityData)
}

// IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry
// An entry (conceptual row) in the list of next-hops on
// outgoing interfaces to which IP multicast datagrams from
// particular sources to a IP multicast group address are
// routed.  Discontinuities in counters in this entry can be
// detected by observing the value of ipMRouteUpTime.
type IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group for which this entry
    // specifies a next-hop on an outgoing interface. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutenexthopgroup interface{}

    // This attribute is a key. The network address which when combined with the
    // corresponding value of ipMRouteNextHopSourceMask identifies the sources for
    // which this entry specifies a next-hop on an outgoing interface. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutenexthopsource interface{}

    // This attribute is a key. The network mask which when combined with the
    // corresponding value of ipMRouteNextHopSource identifies the sources for
    // which this entry specifies a next-hop on an outgoing interface. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutenexthopsourcemask interface{}

    // This attribute is a key. The ifIndex value of the interface for the
    // outgoing interface for this next-hop. The type is interface{} with range:
    // 1..2147483647.
    Ipmroutenexthopifindex interface{}

    // This attribute is a key. The address of the next-hop specific to this
    // entry.  For most interfaces, this is identical to ipMRouteNextHopGroup.
    // NBMA interfaces, however, may have multiple next-hop addresses out a single
    // outgoing interface. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipmroutenexthopentry *IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry) GetEntityData() *types.CommonEntityData {
    ipmroutenexthopentry.EntityData.YFilter = ipmroutenexthopentry.YFilter
    ipmroutenexthopentry.EntityData.YangName = "ipMRouteNextHopEntry"
    ipmroutenexthopentry.EntityData.BundleName = "cisco_ios_xe"
    ipmroutenexthopentry.EntityData.ParentYangName = "ipMRouteNextHopTable"
    ipmroutenexthopentry.EntityData.SegmentPath = "ipMRouteNextHopEntry" + "[ipMRouteNextHopGroup='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopgroup) + "']" + "[ipMRouteNextHopSource='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopsource) + "']" + "[ipMRouteNextHopSourceMask='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopsourcemask) + "']" + "[ipMRouteNextHopIfIndex='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopifindex) + "']" + "[ipMRouteNextHopAddress='" + fmt.Sprintf("%v", ipmroutenexthopentry.Ipmroutenexthopaddress) + "']"
    ipmroutenexthopentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmroutenexthopentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmroutenexthopentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmroutenexthopentry.EntityData.Children = make(map[string]types.YChild)
    ipmroutenexthopentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopGroup"] = types.YLeaf{"Ipmroutenexthopgroup", ipmroutenexthopentry.Ipmroutenexthopgroup}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopSource"] = types.YLeaf{"Ipmroutenexthopsource", ipmroutenexthopentry.Ipmroutenexthopsource}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopSourceMask"] = types.YLeaf{"Ipmroutenexthopsourcemask", ipmroutenexthopentry.Ipmroutenexthopsourcemask}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopIfIndex"] = types.YLeaf{"Ipmroutenexthopifindex", ipmroutenexthopentry.Ipmroutenexthopifindex}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopAddress"] = types.YLeaf{"Ipmroutenexthopaddress", ipmroutenexthopentry.Ipmroutenexthopaddress}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopState"] = types.YLeaf{"Ipmroutenexthopstate", ipmroutenexthopentry.Ipmroutenexthopstate}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopUpTime"] = types.YLeaf{"Ipmroutenexthopuptime", ipmroutenexthopentry.Ipmroutenexthopuptime}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopExpiryTime"] = types.YLeaf{"Ipmroutenexthopexpirytime", ipmroutenexthopentry.Ipmroutenexthopexpirytime}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopClosestMemberHops"] = types.YLeaf{"Ipmroutenexthopclosestmemberhops", ipmroutenexthopentry.Ipmroutenexthopclosestmemberhops}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopProtocol"] = types.YLeaf{"Ipmroutenexthopprotocol", ipmroutenexthopentry.Ipmroutenexthopprotocol}
    ipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopPkts"] = types.YLeaf{"Ipmroutenexthoppkts", ipmroutenexthopentry.Ipmroutenexthoppkts}
    ipmroutenexthopentry.EntityData.Leafs["ciscoIpMRouteNextHopOutLimit"] = types.YLeaf{"Ciscoipmroutenexthopoutlimit", ipmroutenexthopentry.Ciscoipmroutenexthopoutlimit}
    ipmroutenexthopentry.EntityData.Leafs["ciscoIpMRouteNextHopMacHdr"] = types.YLeaf{"Ciscoipmroutenexthopmachdr", ipmroutenexthopentry.Ciscoipmroutenexthopmachdr}
    ipmroutenexthopentry.EntityData.Leafs["ciscoIpMRouteNextHopPkts"] = types.YLeaf{"Ciscoipmroutenexthoppkts", ipmroutenexthopentry.Ciscoipmroutenexthoppkts}
    return &(ipmroutenexthopentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the multicast routing information for
    // a particular interface. The type is slice of
    // IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry.
    Ipmrouteinterfaceentry []IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry
}

func (ipmrouteinterfacetable *IPMROUTESTDMIB_Ipmrouteinterfacetable) GetEntityData() *types.CommonEntityData {
    ipmrouteinterfacetable.EntityData.YFilter = ipmrouteinterfacetable.YFilter
    ipmrouteinterfacetable.EntityData.YangName = "ipMRouteInterfaceTable"
    ipmrouteinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    ipmrouteinterfacetable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipmrouteinterfacetable.EntityData.SegmentPath = "ipMRouteInterfaceTable"
    ipmrouteinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmrouteinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmrouteinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmrouteinterfacetable.EntityData.Children = make(map[string]types.YChild)
    ipmrouteinterfacetable.EntityData.Children["ipMRouteInterfaceEntry"] = types.YChild{"Ipmrouteinterfaceentry", nil}
    for i := range ipmrouteinterfacetable.Ipmrouteinterfaceentry {
        ipmrouteinterfacetable.EntityData.Children[types.GetSegmentPath(&ipmrouteinterfacetable.Ipmrouteinterfaceentry[i])] = types.YChild{"Ipmrouteinterfaceentry", &ipmrouteinterfacetable.Ipmrouteinterfaceentry[i]}
    }
    ipmrouteinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipmrouteinterfacetable.EntityData)
}

// IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry
// An entry (conceptual row) containing the multicast routing
// information for a particular interface.
type IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry struct {
    EntityData types.CommonEntityData
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

func (ipmrouteinterfaceentry *IPMROUTESTDMIB_Ipmrouteinterfacetable_Ipmrouteinterfaceentry) GetEntityData() *types.CommonEntityData {
    ipmrouteinterfaceentry.EntityData.YFilter = ipmrouteinterfaceentry.YFilter
    ipmrouteinterfaceentry.EntityData.YangName = "ipMRouteInterfaceEntry"
    ipmrouteinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    ipmrouteinterfaceentry.EntityData.ParentYangName = "ipMRouteInterfaceTable"
    ipmrouteinterfaceentry.EntityData.SegmentPath = "ipMRouteInterfaceEntry" + "[ipMRouteInterfaceIfIndex='" + fmt.Sprintf("%v", ipmrouteinterfaceentry.Ipmrouteinterfaceifindex) + "']"
    ipmrouteinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmrouteinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmrouteinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmrouteinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    ipmrouteinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceIfIndex"] = types.YLeaf{"Ipmrouteinterfaceifindex", ipmrouteinterfaceentry.Ipmrouteinterfaceifindex}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceTtl"] = types.YLeaf{"Ipmrouteinterfacettl", ipmrouteinterfaceentry.Ipmrouteinterfacettl}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceProtocol"] = types.YLeaf{"Ipmrouteinterfaceprotocol", ipmrouteinterfaceentry.Ipmrouteinterfaceprotocol}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceRateLimit"] = types.YLeaf{"Ipmrouteinterfaceratelimit", ipmrouteinterfaceentry.Ipmrouteinterfaceratelimit}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceInMcastOctets"] = types.YLeaf{"Ipmrouteinterfaceinmcastoctets", ipmrouteinterfaceentry.Ipmrouteinterfaceinmcastoctets}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceOutMcastOctets"] = types.YLeaf{"Ipmrouteinterfaceoutmcastoctets", ipmrouteinterfaceentry.Ipmrouteinterfaceoutmcastoctets}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceHCInMcastOctets"] = types.YLeaf{"Ipmrouteinterfacehcinmcastoctets", ipmrouteinterfaceentry.Ipmrouteinterfacehcinmcastoctets}
    ipmrouteinterfaceentry.EntityData.Leafs["ipMRouteInterfaceHCOutMcastOctets"] = types.YLeaf{"Ipmrouteinterfacehcoutmcastoctets", ipmrouteinterfaceentry.Ipmrouteinterfacehcoutmcastoctets}
    ipmrouteinterfaceentry.EntityData.Leafs["ciscoIpMRouteIfInMcastOctets"] = types.YLeaf{"Ciscoipmrouteifinmcastoctets", ipmrouteinterfaceentry.Ciscoipmrouteifinmcastoctets}
    ipmrouteinterfaceentry.EntityData.Leafs["ciscoIpMRouteIfOutMcastOctets"] = types.YLeaf{"Ciscoipmrouteifoutmcastoctets", ipmrouteinterfaceentry.Ciscoipmrouteifoutmcastoctets}
    ipmrouteinterfaceentry.EntityData.Leafs["ciscoIpMRouteIfInMcastPkts"] = types.YLeaf{"Ciscoipmrouteifinmcastpkts", ipmrouteinterfaceentry.Ciscoipmrouteifinmcastpkts}
    ipmrouteinterfaceentry.EntityData.Leafs["ciscoIpMRouteIfHCInMcastPkts"] = types.YLeaf{"Ciscoipmrouteifhcinmcastpkts", ipmrouteinterfaceentry.Ciscoipmrouteifhcinmcastpkts}
    ipmrouteinterfaceentry.EntityData.Leafs["ciscoIpMRouteIfOutMcastPkts"] = types.YLeaf{"Ciscoipmrouteifoutmcastpkts", ipmrouteinterfaceentry.Ciscoipmrouteifoutmcastpkts}
    ipmrouteinterfaceentry.EntityData.Leafs["ciscoIpMRouteIfHCOutMcastPkts"] = types.YLeaf{"Ciscoipmrouteifhcoutmcastpkts", ipmrouteinterfaceentry.Ciscoipmrouteifhcoutmcastpkts}
    return &(ipmrouteinterfaceentry.EntityData)
}

// IPMROUTESTDMIB_Ipmrouteboundarytable
// The (conceptual) table listing the router's scoped
// multicast address boundaries.
type IPMROUTESTDMIB_Ipmrouteboundarytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ipMRouteBoundaryTable representing a
    // scoped boundary. The type is slice of
    // IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry.
    Ipmrouteboundaryentry []IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry
}

func (ipmrouteboundarytable *IPMROUTESTDMIB_Ipmrouteboundarytable) GetEntityData() *types.CommonEntityData {
    ipmrouteboundarytable.EntityData.YFilter = ipmrouteboundarytable.YFilter
    ipmrouteboundarytable.EntityData.YangName = "ipMRouteBoundaryTable"
    ipmrouteboundarytable.EntityData.BundleName = "cisco_ios_xe"
    ipmrouteboundarytable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipmrouteboundarytable.EntityData.SegmentPath = "ipMRouteBoundaryTable"
    ipmrouteboundarytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmrouteboundarytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmrouteboundarytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmrouteboundarytable.EntityData.Children = make(map[string]types.YChild)
    ipmrouteboundarytable.EntityData.Children["ipMRouteBoundaryEntry"] = types.YChild{"Ipmrouteboundaryentry", nil}
    for i := range ipmrouteboundarytable.Ipmrouteboundaryentry {
        ipmrouteboundarytable.EntityData.Children[types.GetSegmentPath(&ipmrouteboundarytable.Ipmrouteboundaryentry[i])] = types.YChild{"Ipmrouteboundaryentry", &ipmrouteboundarytable.Ipmrouteboundaryentry[i]}
    }
    ipmrouteboundarytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipmrouteboundarytable.EntityData)
}

// IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry
// An entry (conceptual row) in the ipMRouteBoundaryTable
// representing a scoped boundary.
type IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmrouteboundaryaddress interface{}

    // This attribute is a key. The group address mask which when combined with
    // the corresponding value of ipMRouteBoundaryAddress identifies the group
    // range for which the scoped boundary exists. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmrouteboundaryaddressmask interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    Ipmrouteboundarystatus interface{}
}

func (ipmrouteboundaryentry *IPMROUTESTDMIB_Ipmrouteboundarytable_Ipmrouteboundaryentry) GetEntityData() *types.CommonEntityData {
    ipmrouteboundaryentry.EntityData.YFilter = ipmrouteboundaryentry.YFilter
    ipmrouteboundaryentry.EntityData.YangName = "ipMRouteBoundaryEntry"
    ipmrouteboundaryentry.EntityData.BundleName = "cisco_ios_xe"
    ipmrouteboundaryentry.EntityData.ParentYangName = "ipMRouteBoundaryTable"
    ipmrouteboundaryentry.EntityData.SegmentPath = "ipMRouteBoundaryEntry" + "[ipMRouteBoundaryIfIndex='" + fmt.Sprintf("%v", ipmrouteboundaryentry.Ipmrouteboundaryifindex) + "']" + "[ipMRouteBoundaryAddress='" + fmt.Sprintf("%v", ipmrouteboundaryentry.Ipmrouteboundaryaddress) + "']" + "[ipMRouteBoundaryAddressMask='" + fmt.Sprintf("%v", ipmrouteboundaryentry.Ipmrouteboundaryaddressmask) + "']"
    ipmrouteboundaryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmrouteboundaryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmrouteboundaryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmrouteboundaryentry.EntityData.Children = make(map[string]types.YChild)
    ipmrouteboundaryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipmrouteboundaryentry.EntityData.Leafs["ipMRouteBoundaryIfIndex"] = types.YLeaf{"Ipmrouteboundaryifindex", ipmrouteboundaryentry.Ipmrouteboundaryifindex}
    ipmrouteboundaryentry.EntityData.Leafs["ipMRouteBoundaryAddress"] = types.YLeaf{"Ipmrouteboundaryaddress", ipmrouteboundaryentry.Ipmrouteboundaryaddress}
    ipmrouteboundaryentry.EntityData.Leafs["ipMRouteBoundaryAddressMask"] = types.YLeaf{"Ipmrouteboundaryaddressmask", ipmrouteboundaryentry.Ipmrouteboundaryaddressmask}
    ipmrouteboundaryentry.EntityData.Leafs["ipMRouteBoundaryStatus"] = types.YLeaf{"Ipmrouteboundarystatus", ipmrouteboundaryentry.Ipmrouteboundarystatus}
    return &(ipmrouteboundaryentry.EntityData)
}

// IPMROUTESTDMIB_Ipmroutescopenametable
// The (conceptual) table listing the multicast scope names.
type IPMROUTESTDMIB_Ipmroutescopenametable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ipMRouteScopeNameTable representing a
    // multicast scope name. The type is slice of
    // IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry.
    Ipmroutescopenameentry []IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry
}

func (ipmroutescopenametable *IPMROUTESTDMIB_Ipmroutescopenametable) GetEntityData() *types.CommonEntityData {
    ipmroutescopenametable.EntityData.YFilter = ipmroutescopenametable.YFilter
    ipmroutescopenametable.EntityData.YangName = "ipMRouteScopeNameTable"
    ipmroutescopenametable.EntityData.BundleName = "cisco_ios_xe"
    ipmroutescopenametable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipmroutescopenametable.EntityData.SegmentPath = "ipMRouteScopeNameTable"
    ipmroutescopenametable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmroutescopenametable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmroutescopenametable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmroutescopenametable.EntityData.Children = make(map[string]types.YChild)
    ipmroutescopenametable.EntityData.Children["ipMRouteScopeNameEntry"] = types.YChild{"Ipmroutescopenameentry", nil}
    for i := range ipmroutescopenametable.Ipmroutescopenameentry {
        ipmroutescopenametable.EntityData.Children[types.GetSegmentPath(&ipmroutescopenametable.Ipmroutescopenameentry[i])] = types.YChild{"Ipmroutescopenameentry", &ipmroutescopenametable.Ipmroutescopenameentry[i]}
    }
    ipmroutescopenametable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipmroutescopenametable.EntityData)
}

// IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry
// An entry (conceptual row) in the ipMRouteScopeNameTable
// representing a multicast scope name.
type IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The group address which when combined with the
    // corresponding value of ipMRouteScopeNameAddressMask identifies the group
    // range associated with the multicast scope.  Scoped addresses must come from
    // the range 239.x.x.x. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipmroutescopenameaddress interface{}

    // This attribute is a key. The group address mask which when combined with
    // the corresponding value of ipMRouteScopeNameAddress identifies the group
    // range associated with the multicast scope. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipmroutescopenameentry *IPMROUTESTDMIB_Ipmroutescopenametable_Ipmroutescopenameentry) GetEntityData() *types.CommonEntityData {
    ipmroutescopenameentry.EntityData.YFilter = ipmroutescopenameentry.YFilter
    ipmroutescopenameentry.EntityData.YangName = "ipMRouteScopeNameEntry"
    ipmroutescopenameentry.EntityData.BundleName = "cisco_ios_xe"
    ipmroutescopenameentry.EntityData.ParentYangName = "ipMRouteScopeNameTable"
    ipmroutescopenameentry.EntityData.SegmentPath = "ipMRouteScopeNameEntry" + "[ipMRouteScopeNameAddress='" + fmt.Sprintf("%v", ipmroutescopenameentry.Ipmroutescopenameaddress) + "']" + "[ipMRouteScopeNameAddressMask='" + fmt.Sprintf("%v", ipmroutescopenameentry.Ipmroutescopenameaddressmask) + "']" + "[ipMRouteScopeNameLanguage='" + fmt.Sprintf("%v", ipmroutescopenameentry.Ipmroutescopenamelanguage) + "']"
    ipmroutescopenameentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipmroutescopenameentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipmroutescopenameentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipmroutescopenameentry.EntityData.Children = make(map[string]types.YChild)
    ipmroutescopenameentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipmroutescopenameentry.EntityData.Leafs["ipMRouteScopeNameAddress"] = types.YLeaf{"Ipmroutescopenameaddress", ipmroutescopenameentry.Ipmroutescopenameaddress}
    ipmroutescopenameentry.EntityData.Leafs["ipMRouteScopeNameAddressMask"] = types.YLeaf{"Ipmroutescopenameaddressmask", ipmroutescopenameentry.Ipmroutescopenameaddressmask}
    ipmroutescopenameentry.EntityData.Leafs["ipMRouteScopeNameLanguage"] = types.YLeaf{"Ipmroutescopenamelanguage", ipmroutescopenameentry.Ipmroutescopenamelanguage}
    ipmroutescopenameentry.EntityData.Leafs["ipMRouteScopeNameString"] = types.YLeaf{"Ipmroutescopenamestring", ipmroutescopenameentry.Ipmroutescopenamestring}
    ipmroutescopenameentry.EntityData.Leafs["ipMRouteScopeNameDefault"] = types.YLeaf{"Ipmroutescopenamedefault", ipmroutescopenameentry.Ipmroutescopenamedefault}
    ipmroutescopenameentry.EntityData.Leafs["ipMRouteScopeNameStatus"] = types.YLeaf{"Ipmroutescopenamestatus", ipmroutescopenameentry.Ipmroutescopenamestatus}
    return &(ipmroutescopenameentry.EntityData)
}

