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

    
    IpMRoute IPMROUTESTDMIB_IpMRoute

    // The (conceptual) table containing multicast routing information for IP
    // datagrams sent by particular sources to the IP multicast groups known to
    // this router.
    IpMRouteTable IPMROUTESTDMIB_IpMRouteTable

    // The (conceptual) table containing information on the next- hops on outgoing
    // interfaces for routing IP multicast datagrams.  Each entry is one of a list
    // of next-hops on outgoing interfaces for particular sources sending to a
    // particular multicast group address.
    IpMRouteNextHopTable IPMROUTESTDMIB_IpMRouteNextHopTable

    // The (conceptual) table containing multicast routing information specific to
    // interfaces.
    IpMRouteInterfaceTable IPMROUTESTDMIB_IpMRouteInterfaceTable

    // The (conceptual) table listing the router's scoped multicast address
    // boundaries.
    IpMRouteBoundaryTable IPMROUTESTDMIB_IpMRouteBoundaryTable

    // The (conceptual) table listing the multicast scope names.
    IpMRouteScopeNameTable IPMROUTESTDMIB_IpMRouteScopeNameTable
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

    iPMROUTESTDMIB.EntityData.Children = types.NewOrderedMap()
    iPMROUTESTDMIB.EntityData.Children.Append("ipMRoute", types.YChild{"IpMRoute", &iPMROUTESTDMIB.IpMRoute})
    iPMROUTESTDMIB.EntityData.Children.Append("ipMRouteTable", types.YChild{"IpMRouteTable", &iPMROUTESTDMIB.IpMRouteTable})
    iPMROUTESTDMIB.EntityData.Children.Append("ipMRouteNextHopTable", types.YChild{"IpMRouteNextHopTable", &iPMROUTESTDMIB.IpMRouteNextHopTable})
    iPMROUTESTDMIB.EntityData.Children.Append("ipMRouteInterfaceTable", types.YChild{"IpMRouteInterfaceTable", &iPMROUTESTDMIB.IpMRouteInterfaceTable})
    iPMROUTESTDMIB.EntityData.Children.Append("ipMRouteBoundaryTable", types.YChild{"IpMRouteBoundaryTable", &iPMROUTESTDMIB.IpMRouteBoundaryTable})
    iPMROUTESTDMIB.EntityData.Children.Append("ipMRouteScopeNameTable", types.YChild{"IpMRouteScopeNameTable", &iPMROUTESTDMIB.IpMRouteScopeNameTable})
    iPMROUTESTDMIB.EntityData.Leafs = types.NewOrderedMap()

    iPMROUTESTDMIB.EntityData.YListKeys = []string {}

    return &(iPMROUTESTDMIB.EntityData)
}

// IPMROUTESTDMIB_IpMRoute
type IPMROUTESTDMIB_IpMRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The enabled status of IP Multicast routing on this router. The type is
    // IpMRouteEnable.
    IpMRouteEnable interface{}

    // The number of rows in the ipMRouteTable.  This can be used to monitor the
    // multicast routing table size. The type is interface{} with range:
    // 0..4294967295.
    IpMRouteEntryCount interface{}
}

func (ipMRoute *IPMROUTESTDMIB_IpMRoute) GetEntityData() *types.CommonEntityData {
    ipMRoute.EntityData.YFilter = ipMRoute.YFilter
    ipMRoute.EntityData.YangName = "ipMRoute"
    ipMRoute.EntityData.BundleName = "cisco_ios_xe"
    ipMRoute.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipMRoute.EntityData.SegmentPath = "ipMRoute"
    ipMRoute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRoute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRoute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRoute.EntityData.Children = types.NewOrderedMap()
    ipMRoute.EntityData.Leafs = types.NewOrderedMap()
    ipMRoute.EntityData.Leafs.Append("ipMRouteEnable", types.YLeaf{"IpMRouteEnable", ipMRoute.IpMRouteEnable})
    ipMRoute.EntityData.Leafs.Append("ipMRouteEntryCount", types.YLeaf{"IpMRouteEntryCount", ipMRoute.IpMRouteEntryCount})

    ipMRoute.EntityData.YListKeys = []string {}

    return &(ipMRoute.EntityData)
}

// IPMROUTESTDMIB_IpMRoute_IpMRouteEnable represents The enabled status of IP Multicast routing on this router.
type IPMROUTESTDMIB_IpMRoute_IpMRouteEnable string

const (
    IPMROUTESTDMIB_IpMRoute_IpMRouteEnable_enabled IPMROUTESTDMIB_IpMRoute_IpMRouteEnable = "enabled"

    IPMROUTESTDMIB_IpMRoute_IpMRouteEnable_disabled IPMROUTESTDMIB_IpMRoute_IpMRouteEnable = "disabled"
)

// IPMROUTESTDMIB_IpMRouteTable
// The (conceptual) table containing multicast routing
// information for IP datagrams sent by particular sources to
// the IP multicast groups known to this router.
type IPMROUTESTDMIB_IpMRouteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the multicast routing information for
    // IP datagrams from a particular source and addressed to a particular IP
    // multicast group address. Discontinuities in counters in this entry can be
    // detected by observing the value of ipMRouteUpTime. The type is slice of
    // IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry.
    IpMRouteEntry []*IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry
}

func (ipMRouteTable *IPMROUTESTDMIB_IpMRouteTable) GetEntityData() *types.CommonEntityData {
    ipMRouteTable.EntityData.YFilter = ipMRouteTable.YFilter
    ipMRouteTable.EntityData.YangName = "ipMRouteTable"
    ipMRouteTable.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteTable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipMRouteTable.EntityData.SegmentPath = "ipMRouteTable"
    ipMRouteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteTable.EntityData.Children = types.NewOrderedMap()
    ipMRouteTable.EntityData.Children.Append("ipMRouteEntry", types.YChild{"IpMRouteEntry", nil})
    for i := range ipMRouteTable.IpMRouteEntry {
        ipMRouteTable.EntityData.Children.Append(types.GetSegmentPath(ipMRouteTable.IpMRouteEntry[i]), types.YChild{"IpMRouteEntry", ipMRouteTable.IpMRouteEntry[i]})
    }
    ipMRouteTable.EntityData.Leafs = types.NewOrderedMap()

    ipMRouteTable.EntityData.YListKeys = []string {}

    return &(ipMRouteTable.EntityData)
}

// IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry
// An entry (conceptual row) containing the multicast routing
// information for IP datagrams from a particular source and
// addressed to a particular IP multicast group address.
// Discontinuities in counters in this entry can be detected by
// observing the value of ipMRouteUpTime.
type IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains multicast routing information. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteGroup interface{}

    // This attribute is a key. The network address which when combined with the
    // corresponding value of ipMRouteSourceMask identifies the sources for which
    // this entry contains multicast routing information. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteSource interface{}

    // This attribute is a key. The network mask which when combined with the
    // corresponding value of ipMRouteSource identifies the sources for which this
    // entry contains multicast routing information. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteSourceMask interface{}

    // The address of the upstream neighbor (e.g., RPF neighbor) from which IP
    // datagrams from these sources to this multicast address are received, or
    // 0.0.0.0 if the upstream neighbor is unknown (e.g., in CBT). The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteUpstreamNeighbor interface{}

    // The value of ifIndex for the interface on which IP datagrams sent by these
    // sources to this multicast address are received.  A value of 0 indicates
    // that datagrams are not subject to an incoming interface check, but may be
    // accepted on multiple interfaces (e.g., in CBT). The type is interface{}
    // with range: 0..2147483647.
    IpMRouteInIfIndex interface{}

    // The time since the multicast routing information represented by this entry
    // was learned by the router. The type is interface{} with range:
    // 0..4294967295.
    IpMRouteUpTime interface{}

    // The minimum amount of time remaining before this entry will be aged out. 
    // The value 0 indicates that the entry is not subject to aging. The type is
    // interface{} with range: 0..4294967295.
    IpMRouteExpiryTime interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address. The type is interface{} with
    // range: 0..4294967295.
    IpMRoutePkts interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address, which were dropped because they
    // were not received on the interface indicated by ipMRouteInIfIndex.  Packets
    // which are not subject to an incoming interface check (e.g., using CBT) are
    // not counted. The type is interface{} with range: 0..4294967295.
    IpMRouteDifferentInIfPackets interface{}

    // The number of octets contained in IP datagrams which were received from
    // these sources and addressed to this multicast group address, and which were
    // forwarded by this router. The type is interface{} with range:
    // 0..4294967295.
    IpMRouteOctets interface{}

    // The multicast routing protocol via which this multicast forwarding entry
    // was learned. The type is IANAipMRouteProtocol.
    IpMRouteProtocol interface{}

    // The routing mechanism via which the route used to find the upstream or
    // parent interface for this multicast forwarding entry was learned. 
    // Inclusion of values for routing protocols is not intended to imply that
    // those protocols need be supported. The type is IANAipRouteProtocol.
    IpMRouteRtProto interface{}

    // The address portion of the route used to find the upstream or parent
    // interface for this multicast forwarding entry. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteRtAddress interface{}

    // The mask associated with the route used to find the upstream or parent
    // interface for this multicast forwarding entry. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteRtMask interface{}

    // The reason the given route was placed in the (logical) multicast Routing
    // Information Base (RIB).  A value of unicast means that the route would
    // normally be placed only in the unicast RIB, but was placed in the multicast
    // RIB (instead or in addition) due to local configuration, such as when
    // running PIM over RIP.  A value of multicast means that      the route was
    // explicitly added to the multicast RIB by the routing protocol, such as
    // DVMRP or Multiprotocol BGP. The type is IpMRouteRtType.
    IpMRouteRtType interface{}

    // The number of octets contained in IP datagrams which were received from
    // these sources and addressed to this multicast group address, and which were
    // forwarded by this router. This object is a 64-bit version of
    // ipMRouteOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    IpMRouteHCOctets interface{}

    // Boolean, indicates whether this route is pruned. A pruned route is one that
    // has an empty outgoing interface list or all interfaces are in Pruned state.
    // A multicast packet that matches a pruned route doesn't get forwarded. The
    // type is bool.
    CiscoIpMRoutePruneFlag interface{}

    // Boolean, indicating PIM multicast routing protocol sparse-mode (versus
    // dense-mode).  In sparse-mode, packets are forwarded only out interfaces
    // that have been joined. In dense-mode, they are forwarded out all interfaces
    // that have not been pruned. The type is bool.
    CiscoIpMRouteSparseFlag interface{}

    // Boolean, indicating whether there is a directly connected member for a
    // group attached to the router. The type is bool.
    CiscoIpMRouteConnectedFlag interface{}

    // Boolean, indicating whether local system is a member of a group on any
    // interface. The type is bool.
    CiscoIpMRouteLocalFlag interface{}

    // Boolean, indicates whether to send registers for the entry. A first hop
    // router directly connected to a multicast source host, as well as a border
    // router on the boundary of two domains running different multicast routing
    // protocols, encapsulates packets to be sent on the shared tree. This is done
    // until the RP sends Joins back to this router. The type is bool.
    CiscoIpMRouteRegisterFlag interface{}

    // Boolean, indicating whether there is a Prune state for this source along
    // the shared tree. The type is bool.
    CiscoIpMRouteRpFlag interface{}

    // Boolean, indicating whether data is being received on the SPT tree, ie the
    // Shortest Path Tree. The type is bool.
    CiscoIpMRouteSptFlag interface{}

    // Bits per second forwarded by this router.  This is the sum of all forwarded
    // bits during a 1 second interval.  At the end of each second the field is
    // cleared. This object has been superseded by ciscoIpMRouteBps2 (which is the
    // 64-bit version of this object). The type is interface{} with range:
    // 0..4294967295.
    CiscoIpMRouteBps interface{}

    // Metric - The best metric heard from Assert messages. This object has been
    // replaced by ciscoIpMRouteMetric2 in order to correctly represent a 32-bit
    // unsigned metric value. The type is interface{} with range: 0..2147483647.
    CiscoIpMRouteMetric interface{}

    // Metric Preference - The best metric preference heard from Assert messages.
    // The type is interface{} with range: 0..2147483647.
    CiscoIpMRouteMetricPreference interface{}

    // Incoming interface's limit for rate limiting data traffic, in Kbps.
    // Replaced by ciscoIpMRouteInLimit2. The type is interface{} with range:
    // 0..2147483647. Units are Kbits/second.
    CiscoIpMRouteInLimit interface{}

    // How long has it been since the last multicast packet was fastswitched. The
    // type is interface{} with range: 0..4294967295.
    CiscoIpMRouteLastUsed interface{}

    // Incoming interface's limit for rate limiting data traffic, in Kbps. The
    // type is interface{} with range: 0..4294967295. Units are Kbits/second.
    CiscoIpMRouteInLimit2 interface{}

    // Boolean, indicates whether this route is created due to SPT threshold. The
    // type is bool.
    CiscoIpMRouteJoinFlag interface{}

    // Boolean, indicates whether this route is learned via MSDP. The type is
    // bool.
    CiscoIpMRouteMsdpFlag interface{}

    // Boolean, indicates whether to send join for this entry. The type is bool.
    CiscoIpMRouteProxyJoinFlag interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address. This object is a 64-bit version
    // of ipMRoutePkts. The type is interface{} with range:
    // 0..18446744073709551615.
    CiscoIpMRoutePkts interface{}

    // The number of packets which this router has received from these sources and
    // addressed to this multicast group address, which were not received from the
    // interface indicated by ipMRouteInIfIndex. This object is a 64-bit version
    // of ipMRouteDifferentInIfPackets. The type is interface{} with range:
    // 0..18446744073709551615.
    CiscoIpMRouteDifferentInIfPkts interface{}

    // The number of octets contained in IP datagrams which were received from
    // these sources and addressed to this multicast group address, and which were
    // forwarded by this router. This object is a 64-bit version of
    // ipMRouteOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    CiscoIpMRouteOctets interface{}

    // Bits per second forwarded by this router. This is the sum of all forwarded
    // bits during a 1 second interval. At the end of each second the field is
    // cleared. The type is interface{} with range: 0..18446744073709551615.
    CiscoIpMRouteBps2 interface{}

    // Metric - The best metric heard from Assert messages. The type is
    // interface{} with range: 0..4294967295.
    CiscoIpMRouteMetric2 interface{}
}

func (ipMRouteEntry *IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry) GetEntityData() *types.CommonEntityData {
    ipMRouteEntry.EntityData.YFilter = ipMRouteEntry.YFilter
    ipMRouteEntry.EntityData.YangName = "ipMRouteEntry"
    ipMRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteEntry.EntityData.ParentYangName = "ipMRouteTable"
    ipMRouteEntry.EntityData.SegmentPath = "ipMRouteEntry" + types.AddKeyToken(ipMRouteEntry.IpMRouteGroup, "ipMRouteGroup") + types.AddKeyToken(ipMRouteEntry.IpMRouteSource, "ipMRouteSource") + types.AddKeyToken(ipMRouteEntry.IpMRouteSourceMask, "ipMRouteSourceMask")
    ipMRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteEntry.EntityData.Children = types.NewOrderedMap()
    ipMRouteEntry.EntityData.Leafs = types.NewOrderedMap()
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteGroup", types.YLeaf{"IpMRouteGroup", ipMRouteEntry.IpMRouteGroup})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteSource", types.YLeaf{"IpMRouteSource", ipMRouteEntry.IpMRouteSource})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteSourceMask", types.YLeaf{"IpMRouteSourceMask", ipMRouteEntry.IpMRouteSourceMask})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteUpstreamNeighbor", types.YLeaf{"IpMRouteUpstreamNeighbor", ipMRouteEntry.IpMRouteUpstreamNeighbor})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteInIfIndex", types.YLeaf{"IpMRouteInIfIndex", ipMRouteEntry.IpMRouteInIfIndex})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteUpTime", types.YLeaf{"IpMRouteUpTime", ipMRouteEntry.IpMRouteUpTime})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteExpiryTime", types.YLeaf{"IpMRouteExpiryTime", ipMRouteEntry.IpMRouteExpiryTime})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRoutePkts", types.YLeaf{"IpMRoutePkts", ipMRouteEntry.IpMRoutePkts})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteDifferentInIfPackets", types.YLeaf{"IpMRouteDifferentInIfPackets", ipMRouteEntry.IpMRouteDifferentInIfPackets})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteOctets", types.YLeaf{"IpMRouteOctets", ipMRouteEntry.IpMRouteOctets})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteProtocol", types.YLeaf{"IpMRouteProtocol", ipMRouteEntry.IpMRouteProtocol})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteRtProto", types.YLeaf{"IpMRouteRtProto", ipMRouteEntry.IpMRouteRtProto})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteRtAddress", types.YLeaf{"IpMRouteRtAddress", ipMRouteEntry.IpMRouteRtAddress})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteRtMask", types.YLeaf{"IpMRouteRtMask", ipMRouteEntry.IpMRouteRtMask})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteRtType", types.YLeaf{"IpMRouteRtType", ipMRouteEntry.IpMRouteRtType})
    ipMRouteEntry.EntityData.Leafs.Append("ipMRouteHCOctets", types.YLeaf{"IpMRouteHCOctets", ipMRouteEntry.IpMRouteHCOctets})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRoutePruneFlag", types.YLeaf{"CiscoIpMRoutePruneFlag", ipMRouteEntry.CiscoIpMRoutePruneFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteSparseFlag", types.YLeaf{"CiscoIpMRouteSparseFlag", ipMRouteEntry.CiscoIpMRouteSparseFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteConnectedFlag", types.YLeaf{"CiscoIpMRouteConnectedFlag", ipMRouteEntry.CiscoIpMRouteConnectedFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteLocalFlag", types.YLeaf{"CiscoIpMRouteLocalFlag", ipMRouteEntry.CiscoIpMRouteLocalFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteRegisterFlag", types.YLeaf{"CiscoIpMRouteRegisterFlag", ipMRouteEntry.CiscoIpMRouteRegisterFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteRpFlag", types.YLeaf{"CiscoIpMRouteRpFlag", ipMRouteEntry.CiscoIpMRouteRpFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteSptFlag", types.YLeaf{"CiscoIpMRouteSptFlag", ipMRouteEntry.CiscoIpMRouteSptFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteBps", types.YLeaf{"CiscoIpMRouteBps", ipMRouteEntry.CiscoIpMRouteBps})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteMetric", types.YLeaf{"CiscoIpMRouteMetric", ipMRouteEntry.CiscoIpMRouteMetric})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteMetricPreference", types.YLeaf{"CiscoIpMRouteMetricPreference", ipMRouteEntry.CiscoIpMRouteMetricPreference})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteInLimit", types.YLeaf{"CiscoIpMRouteInLimit", ipMRouteEntry.CiscoIpMRouteInLimit})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteLastUsed", types.YLeaf{"CiscoIpMRouteLastUsed", ipMRouteEntry.CiscoIpMRouteLastUsed})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteInLimit2", types.YLeaf{"CiscoIpMRouteInLimit2", ipMRouteEntry.CiscoIpMRouteInLimit2})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteJoinFlag", types.YLeaf{"CiscoIpMRouteJoinFlag", ipMRouteEntry.CiscoIpMRouteJoinFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteMsdpFlag", types.YLeaf{"CiscoIpMRouteMsdpFlag", ipMRouteEntry.CiscoIpMRouteMsdpFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteProxyJoinFlag", types.YLeaf{"CiscoIpMRouteProxyJoinFlag", ipMRouteEntry.CiscoIpMRouteProxyJoinFlag})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRoutePkts", types.YLeaf{"CiscoIpMRoutePkts", ipMRouteEntry.CiscoIpMRoutePkts})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteDifferentInIfPkts", types.YLeaf{"CiscoIpMRouteDifferentInIfPkts", ipMRouteEntry.CiscoIpMRouteDifferentInIfPkts})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteOctets", types.YLeaf{"CiscoIpMRouteOctets", ipMRouteEntry.CiscoIpMRouteOctets})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteBps2", types.YLeaf{"CiscoIpMRouteBps2", ipMRouteEntry.CiscoIpMRouteBps2})
    ipMRouteEntry.EntityData.Leafs.Append("ciscoIpMRouteMetric2", types.YLeaf{"CiscoIpMRouteMetric2", ipMRouteEntry.CiscoIpMRouteMetric2})

    ipMRouteEntry.EntityData.YListKeys = []string {"IpMRouteGroup", "IpMRouteSource", "IpMRouteSourceMask"}

    return &(ipMRouteEntry.EntityData)
}

// IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteRtType represents routing protocol, such as DVMRP or Multiprotocol BGP.
type IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteRtType string

const (
    IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteRtType_unicast IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteRtType = "unicast"

    IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteRtType_multicast IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteRtType = "multicast"
)

// IPMROUTESTDMIB_IpMRouteNextHopTable
// The (conceptual) table containing information on the next-
// hops on outgoing interfaces for routing IP multicast
// datagrams.  Each entry is one of a list of next-hops on
// outgoing interfaces for particular sources sending to a
// particular multicast group address.
type IPMROUTESTDMIB_IpMRouteNextHopTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the list of next-hops on outgoing interfaces
    // to which IP multicast datagrams from particular sources to a IP multicast
    // group address are routed.  Discontinuities in counters in this entry can be
    // detected by observing the value of ipMRouteUpTime. The type is slice of
    // IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry.
    IpMRouteNextHopEntry []*IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry
}

func (ipMRouteNextHopTable *IPMROUTESTDMIB_IpMRouteNextHopTable) GetEntityData() *types.CommonEntityData {
    ipMRouteNextHopTable.EntityData.YFilter = ipMRouteNextHopTable.YFilter
    ipMRouteNextHopTable.EntityData.YangName = "ipMRouteNextHopTable"
    ipMRouteNextHopTable.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteNextHopTable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipMRouteNextHopTable.EntityData.SegmentPath = "ipMRouteNextHopTable"
    ipMRouteNextHopTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteNextHopTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteNextHopTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteNextHopTable.EntityData.Children = types.NewOrderedMap()
    ipMRouteNextHopTable.EntityData.Children.Append("ipMRouteNextHopEntry", types.YChild{"IpMRouteNextHopEntry", nil})
    for i := range ipMRouteNextHopTable.IpMRouteNextHopEntry {
        ipMRouteNextHopTable.EntityData.Children.Append(types.GetSegmentPath(ipMRouteNextHopTable.IpMRouteNextHopEntry[i]), types.YChild{"IpMRouteNextHopEntry", ipMRouteNextHopTable.IpMRouteNextHopEntry[i]})
    }
    ipMRouteNextHopTable.EntityData.Leafs = types.NewOrderedMap()

    ipMRouteNextHopTable.EntityData.YListKeys = []string {}

    return &(ipMRouteNextHopTable.EntityData)
}

// IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry
// An entry (conceptual row) in the list of next-hops on
// outgoing interfaces to which IP multicast datagrams from
// particular sources to a IP multicast group address are
// routed.  Discontinuities in counters in this entry can be
// detected by observing the value of ipMRouteUpTime.
type IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group for which this entry
    // specifies a next-hop on an outgoing interface. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteNextHopGroup interface{}

    // This attribute is a key. The network address which when combined with the
    // corresponding value of ipMRouteNextHopSourceMask identifies the sources for
    // which this entry specifies a next-hop on an outgoing interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteNextHopSource interface{}

    // This attribute is a key. The network mask which when combined with the
    // corresponding value of ipMRouteNextHopSource identifies the sources for
    // which this entry specifies a next-hop on an outgoing interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteNextHopSourceMask interface{}

    // This attribute is a key. The ifIndex value of the interface for the
    // outgoing interface for this next-hop. The type is interface{} with range:
    // 1..2147483647.
    IpMRouteNextHopIfIndex interface{}

    // This attribute is a key. The address of the next-hop specific to this
    // entry.  For most interfaces, this is identical to ipMRouteNextHopGroup.
    // NBMA interfaces, however, may have multiple next-hop addresses out a single
    // outgoing interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteNextHopAddress interface{}

    // An indication of whether the outgoing interface and next- hop represented
    // by this entry is currently being used to forward IP datagrams.  The value
    // 'forwarding' indicates it is currently being used; the value 'pruned'
    // indicates it is not. The type is IpMRouteNextHopState.
    IpMRouteNextHopState interface{}

    // The time since the multicast routing information represented by this entry
    // was learned by the router. The type is interface{} with range:
    // 0..4294967295.
    IpMRouteNextHopUpTime interface{}

    // The minimum amount of time remaining before this entry will be aged out. 
    // If ipMRouteNextHopState is pruned(1), the remaining time until the prune
    // expires and the state reverts to forwarding(2).  Otherwise, the remaining
    // time until this entry is removed from the table.  The time remaining may be
    // copied from ipMRouteExpiryTime if the protocol in use for this entry does
    // not specify next-hop timers.  The value 0      indicates that the entry is
    // not subject to aging. The type is interface{} with range: 0..4294967295.
    IpMRouteNextHopExpiryTime interface{}

    // The minimum number of hops between this router and any member of this IP
    // multicast group reached via this next-hop on this outgoing interface.  Any
    // IP multicast datagrams for the group which have a TTL less than this number
    // of hops will not be forwarded to this next-hop. The type is interface{}
    // with range: -2147483648..2147483647.
    IpMRouteNextHopClosestMemberHops interface{}

    // The routing mechanism via which this next-hop was learned. The type is
    // IANAipMRouteProtocol.
    IpMRouteNextHopProtocol interface{}

    // The number of packets which have been forwarded using this route. The type
    // is interface{} with range: 0..4294967295.
    IpMRouteNextHopPkts interface{}

    // An outgoing interface's limit for rate limiting data traffic, in Kbps. The
    // type is interface{} with range: 0..4294967295. Units are Kbits/second.
    CiscoIpMRouteNextHopOutLimit interface{}

    // The data link mac address header for a multicast datagram. Used by IP
    // multicast fastswitching. The type is string.
    CiscoIpMRouteNextHopMacHdr interface{}

    // The number of packets which have been forwarded using this route. This
    // object is a 64-bit version of ipMRouteNextHopPkts. The type is interface{}
    // with range: 0..18446744073709551615.
    CiscoIpMRouteNextHopPkts interface{}
}

func (ipMRouteNextHopEntry *IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry) GetEntityData() *types.CommonEntityData {
    ipMRouteNextHopEntry.EntityData.YFilter = ipMRouteNextHopEntry.YFilter
    ipMRouteNextHopEntry.EntityData.YangName = "ipMRouteNextHopEntry"
    ipMRouteNextHopEntry.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteNextHopEntry.EntityData.ParentYangName = "ipMRouteNextHopTable"
    ipMRouteNextHopEntry.EntityData.SegmentPath = "ipMRouteNextHopEntry" + types.AddKeyToken(ipMRouteNextHopEntry.IpMRouteNextHopGroup, "ipMRouteNextHopGroup") + types.AddKeyToken(ipMRouteNextHopEntry.IpMRouteNextHopSource, "ipMRouteNextHopSource") + types.AddKeyToken(ipMRouteNextHopEntry.IpMRouteNextHopSourceMask, "ipMRouteNextHopSourceMask") + types.AddKeyToken(ipMRouteNextHopEntry.IpMRouteNextHopIfIndex, "ipMRouteNextHopIfIndex") + types.AddKeyToken(ipMRouteNextHopEntry.IpMRouteNextHopAddress, "ipMRouteNextHopAddress")
    ipMRouteNextHopEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteNextHopEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteNextHopEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteNextHopEntry.EntityData.Children = types.NewOrderedMap()
    ipMRouteNextHopEntry.EntityData.Leafs = types.NewOrderedMap()
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopGroup", types.YLeaf{"IpMRouteNextHopGroup", ipMRouteNextHopEntry.IpMRouteNextHopGroup})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopSource", types.YLeaf{"IpMRouteNextHopSource", ipMRouteNextHopEntry.IpMRouteNextHopSource})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopSourceMask", types.YLeaf{"IpMRouteNextHopSourceMask", ipMRouteNextHopEntry.IpMRouteNextHopSourceMask})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopIfIndex", types.YLeaf{"IpMRouteNextHopIfIndex", ipMRouteNextHopEntry.IpMRouteNextHopIfIndex})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopAddress", types.YLeaf{"IpMRouteNextHopAddress", ipMRouteNextHopEntry.IpMRouteNextHopAddress})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopState", types.YLeaf{"IpMRouteNextHopState", ipMRouteNextHopEntry.IpMRouteNextHopState})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopUpTime", types.YLeaf{"IpMRouteNextHopUpTime", ipMRouteNextHopEntry.IpMRouteNextHopUpTime})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopExpiryTime", types.YLeaf{"IpMRouteNextHopExpiryTime", ipMRouteNextHopEntry.IpMRouteNextHopExpiryTime})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopClosestMemberHops", types.YLeaf{"IpMRouteNextHopClosestMemberHops", ipMRouteNextHopEntry.IpMRouteNextHopClosestMemberHops})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopProtocol", types.YLeaf{"IpMRouteNextHopProtocol", ipMRouteNextHopEntry.IpMRouteNextHopProtocol})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopPkts", types.YLeaf{"IpMRouteNextHopPkts", ipMRouteNextHopEntry.IpMRouteNextHopPkts})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ciscoIpMRouteNextHopOutLimit", types.YLeaf{"CiscoIpMRouteNextHopOutLimit", ipMRouteNextHopEntry.CiscoIpMRouteNextHopOutLimit})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ciscoIpMRouteNextHopMacHdr", types.YLeaf{"CiscoIpMRouteNextHopMacHdr", ipMRouteNextHopEntry.CiscoIpMRouteNextHopMacHdr})
    ipMRouteNextHopEntry.EntityData.Leafs.Append("ciscoIpMRouteNextHopPkts", types.YLeaf{"CiscoIpMRouteNextHopPkts", ipMRouteNextHopEntry.CiscoIpMRouteNextHopPkts})

    ipMRouteNextHopEntry.EntityData.YListKeys = []string {"IpMRouteNextHopGroup", "IpMRouteNextHopSource", "IpMRouteNextHopSourceMask", "IpMRouteNextHopIfIndex", "IpMRouteNextHopAddress"}

    return &(ipMRouteNextHopEntry.EntityData)
}

// IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopState represents not.
type IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopState string

const (
    IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopState_pruned IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopState = "pruned"

    IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopState_forwarding IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopState = "forwarding"
)

// IPMROUTESTDMIB_IpMRouteInterfaceTable
// The (conceptual) table containing multicast routing
// information specific to interfaces.
type IPMROUTESTDMIB_IpMRouteInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the multicast routing information for
    // a particular interface. The type is slice of
    // IPMROUTESTDMIB_IpMRouteInterfaceTable_IpMRouteInterfaceEntry.
    IpMRouteInterfaceEntry []*IPMROUTESTDMIB_IpMRouteInterfaceTable_IpMRouteInterfaceEntry
}

func (ipMRouteInterfaceTable *IPMROUTESTDMIB_IpMRouteInterfaceTable) GetEntityData() *types.CommonEntityData {
    ipMRouteInterfaceTable.EntityData.YFilter = ipMRouteInterfaceTable.YFilter
    ipMRouteInterfaceTable.EntityData.YangName = "ipMRouteInterfaceTable"
    ipMRouteInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteInterfaceTable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipMRouteInterfaceTable.EntityData.SegmentPath = "ipMRouteInterfaceTable"
    ipMRouteInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteInterfaceTable.EntityData.Children = types.NewOrderedMap()
    ipMRouteInterfaceTable.EntityData.Children.Append("ipMRouteInterfaceEntry", types.YChild{"IpMRouteInterfaceEntry", nil})
    for i := range ipMRouteInterfaceTable.IpMRouteInterfaceEntry {
        ipMRouteInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(ipMRouteInterfaceTable.IpMRouteInterfaceEntry[i]), types.YChild{"IpMRouteInterfaceEntry", ipMRouteInterfaceTable.IpMRouteInterfaceEntry[i]})
    }
    ipMRouteInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    ipMRouteInterfaceTable.EntityData.YListKeys = []string {}

    return &(ipMRouteInterfaceTable.EntityData)
}

// IPMROUTESTDMIB_IpMRouteInterfaceTable_IpMRouteInterfaceEntry
// An entry (conceptual row) containing the multicast routing
// information for a particular interface.
type IPMROUTESTDMIB_IpMRouteInterfaceTable_IpMRouteInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the interface for which this
    // entry contains information. The type is interface{} with range:
    // 1..2147483647.
    IpMRouteInterfaceIfIndex interface{}

    // The datagram TTL threshold for the interface. Any IP multicast datagrams
    // with a TTL less than this threshold will not be forwarded out the
    // interface. The default value of 0 means all multicast packets are forwarded
    // out the interface. The type is interface{} with range: 0..255.
    IpMRouteInterfaceTtl interface{}

    // The routing protocol running on this interface. The type is
    // IANAipMRouteProtocol.
    IpMRouteInterfaceProtocol interface{}

    // The rate-limit, in kilobits per second, of forwarded multicast traffic on
    // the interface.  A rate-limit of 0 indicates that no rate limiting is done.
    // The type is interface{} with range: -2147483648..2147483647.
    IpMRouteInterfaceRateLimit interface{}

    // The number of octets of multicast packets that have arrived on the
    // interface, including framing characters.  This object is similar to
    // ifInOctets in the Interfaces MIB, except that only multicast packets are
    // counted. The type is interface{} with range: 0..4294967295.
    IpMRouteInterfaceInMcastOctets interface{}

    // The number of octets of multicast packets that have been sent on the
    // interface. The type is interface{} with range: 0..4294967295.
    IpMRouteInterfaceOutMcastOctets interface{}

    // The number of octets of multicast packets that have arrived on the
    // interface, including framing characters.  This object is a 64-bit version
    // of ipMRouteInterfaceInMcastOctets.  It is similar to ifHCInOctets in the
    // Interfaces MIB, except that only multicast packets are counted. The type is
    // interface{} with range: 0..18446744073709551615.
    IpMRouteInterfaceHCInMcastOctets interface{}

    // The number of octets of multicast packets that have been      sent on the
    // interface.  This object is a 64-bit version of
    // ipMRouteInterfaceOutMcastOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    IpMRouteInterfaceHCOutMcastOctets interface{}

    // The number of octets of multicast packets that have arrived on the
    // interface. This object is a 64-bit version of
    // ipMRouteInterfaceInMcastOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    CiscoIpMRouteIfInMcastOctets interface{}

    // The number of octets of multicast packets that have been sent on the
    // interface. This object is a 64-bit version of
    // ipMRouteInterfaceOutMcastOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    CiscoIpMRouteIfOutMcastOctets interface{}

    // The number of multicast packets that have arrived on the interface. The
    // type is interface{} with range: 0..4294967295.
    CiscoIpMRouteIfInMcastPkts interface{}

    // The number of multicast packets that have arrived on the interface. This
    // object is a 64-bit version of ciscoIpMRouteIfInMcastPkts. The type is
    // interface{} with range: 0..18446744073709551615.
    CiscoIpMRouteIfHCInMcastPkts interface{}

    // The number of multicast packets that have been sent on the interface. The
    // type is interface{} with range: 0..4294967295.
    CiscoIpMRouteIfOutMcastPkts interface{}

    // The number of multicast packets that have been sent on the interface. This
    // object is a 64-bit version of ciscoIpMRouteIfOutMcastPkts. The type is
    // interface{} with range: 0..18446744073709551615.
    CiscoIpMRouteIfHCOutMcastPkts interface{}
}

func (ipMRouteInterfaceEntry *IPMROUTESTDMIB_IpMRouteInterfaceTable_IpMRouteInterfaceEntry) GetEntityData() *types.CommonEntityData {
    ipMRouteInterfaceEntry.EntityData.YFilter = ipMRouteInterfaceEntry.YFilter
    ipMRouteInterfaceEntry.EntityData.YangName = "ipMRouteInterfaceEntry"
    ipMRouteInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteInterfaceEntry.EntityData.ParentYangName = "ipMRouteInterfaceTable"
    ipMRouteInterfaceEntry.EntityData.SegmentPath = "ipMRouteInterfaceEntry" + types.AddKeyToken(ipMRouteInterfaceEntry.IpMRouteInterfaceIfIndex, "ipMRouteInterfaceIfIndex")
    ipMRouteInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    ipMRouteInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceIfIndex", types.YLeaf{"IpMRouteInterfaceIfIndex", ipMRouteInterfaceEntry.IpMRouteInterfaceIfIndex})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceTtl", types.YLeaf{"IpMRouteInterfaceTtl", ipMRouteInterfaceEntry.IpMRouteInterfaceTtl})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceProtocol", types.YLeaf{"IpMRouteInterfaceProtocol", ipMRouteInterfaceEntry.IpMRouteInterfaceProtocol})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceRateLimit", types.YLeaf{"IpMRouteInterfaceRateLimit", ipMRouteInterfaceEntry.IpMRouteInterfaceRateLimit})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceInMcastOctets", types.YLeaf{"IpMRouteInterfaceInMcastOctets", ipMRouteInterfaceEntry.IpMRouteInterfaceInMcastOctets})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceOutMcastOctets", types.YLeaf{"IpMRouteInterfaceOutMcastOctets", ipMRouteInterfaceEntry.IpMRouteInterfaceOutMcastOctets})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceHCInMcastOctets", types.YLeaf{"IpMRouteInterfaceHCInMcastOctets", ipMRouteInterfaceEntry.IpMRouteInterfaceHCInMcastOctets})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ipMRouteInterfaceHCOutMcastOctets", types.YLeaf{"IpMRouteInterfaceHCOutMcastOctets", ipMRouteInterfaceEntry.IpMRouteInterfaceHCOutMcastOctets})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ciscoIpMRouteIfInMcastOctets", types.YLeaf{"CiscoIpMRouteIfInMcastOctets", ipMRouteInterfaceEntry.CiscoIpMRouteIfInMcastOctets})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ciscoIpMRouteIfOutMcastOctets", types.YLeaf{"CiscoIpMRouteIfOutMcastOctets", ipMRouteInterfaceEntry.CiscoIpMRouteIfOutMcastOctets})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ciscoIpMRouteIfInMcastPkts", types.YLeaf{"CiscoIpMRouteIfInMcastPkts", ipMRouteInterfaceEntry.CiscoIpMRouteIfInMcastPkts})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ciscoIpMRouteIfHCInMcastPkts", types.YLeaf{"CiscoIpMRouteIfHCInMcastPkts", ipMRouteInterfaceEntry.CiscoIpMRouteIfHCInMcastPkts})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ciscoIpMRouteIfOutMcastPkts", types.YLeaf{"CiscoIpMRouteIfOutMcastPkts", ipMRouteInterfaceEntry.CiscoIpMRouteIfOutMcastPkts})
    ipMRouteInterfaceEntry.EntityData.Leafs.Append("ciscoIpMRouteIfHCOutMcastPkts", types.YLeaf{"CiscoIpMRouteIfHCOutMcastPkts", ipMRouteInterfaceEntry.CiscoIpMRouteIfHCOutMcastPkts})

    ipMRouteInterfaceEntry.EntityData.YListKeys = []string {"IpMRouteInterfaceIfIndex"}

    return &(ipMRouteInterfaceEntry.EntityData)
}

// IPMROUTESTDMIB_IpMRouteBoundaryTable
// The (conceptual) table listing the router's scoped
// multicast address boundaries.
type IPMROUTESTDMIB_IpMRouteBoundaryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ipMRouteBoundaryTable representing a
    // scoped boundary. The type is slice of
    // IPMROUTESTDMIB_IpMRouteBoundaryTable_IpMRouteBoundaryEntry.
    IpMRouteBoundaryEntry []*IPMROUTESTDMIB_IpMRouteBoundaryTable_IpMRouteBoundaryEntry
}

func (ipMRouteBoundaryTable *IPMROUTESTDMIB_IpMRouteBoundaryTable) GetEntityData() *types.CommonEntityData {
    ipMRouteBoundaryTable.EntityData.YFilter = ipMRouteBoundaryTable.YFilter
    ipMRouteBoundaryTable.EntityData.YangName = "ipMRouteBoundaryTable"
    ipMRouteBoundaryTable.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteBoundaryTable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipMRouteBoundaryTable.EntityData.SegmentPath = "ipMRouteBoundaryTable"
    ipMRouteBoundaryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteBoundaryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteBoundaryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteBoundaryTable.EntityData.Children = types.NewOrderedMap()
    ipMRouteBoundaryTable.EntityData.Children.Append("ipMRouteBoundaryEntry", types.YChild{"IpMRouteBoundaryEntry", nil})
    for i := range ipMRouteBoundaryTable.IpMRouteBoundaryEntry {
        ipMRouteBoundaryTable.EntityData.Children.Append(types.GetSegmentPath(ipMRouteBoundaryTable.IpMRouteBoundaryEntry[i]), types.YChild{"IpMRouteBoundaryEntry", ipMRouteBoundaryTable.IpMRouteBoundaryEntry[i]})
    }
    ipMRouteBoundaryTable.EntityData.Leafs = types.NewOrderedMap()

    ipMRouteBoundaryTable.EntityData.YListKeys = []string {}

    return &(ipMRouteBoundaryTable.EntityData)
}

// IPMROUTESTDMIB_IpMRouteBoundaryTable_IpMRouteBoundaryEntry
// An entry (conceptual row) in the ipMRouteBoundaryTable
// representing a scoped boundary.
type IPMROUTESTDMIB_IpMRouteBoundaryTable_IpMRouteBoundaryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IfIndex value for the interface to which this
    // boundary applies.  Packets with a destination address in the associated
    // address/mask range will not be forwarded out this interface. The type is
    // interface{} with range: 1..2147483647.
    IpMRouteBoundaryIfIndex interface{}

    // This attribute is a key. The group address which when combined with the
    // corresponding value of ipMRouteBoundaryAddressMask identifies the group
    // range for which the scoped boundary exists.  Scoped addresses must come
    // from the range 239.x.x.x as specified in RFC 2365. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteBoundaryAddress interface{}

    // This attribute is a key. The group address mask which when combined with
    // the corresponding value of ipMRouteBoundaryAddress identifies the group
    // range for which the scoped boundary exists. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteBoundaryAddressMask interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    IpMRouteBoundaryStatus interface{}
}

func (ipMRouteBoundaryEntry *IPMROUTESTDMIB_IpMRouteBoundaryTable_IpMRouteBoundaryEntry) GetEntityData() *types.CommonEntityData {
    ipMRouteBoundaryEntry.EntityData.YFilter = ipMRouteBoundaryEntry.YFilter
    ipMRouteBoundaryEntry.EntityData.YangName = "ipMRouteBoundaryEntry"
    ipMRouteBoundaryEntry.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteBoundaryEntry.EntityData.ParentYangName = "ipMRouteBoundaryTable"
    ipMRouteBoundaryEntry.EntityData.SegmentPath = "ipMRouteBoundaryEntry" + types.AddKeyToken(ipMRouteBoundaryEntry.IpMRouteBoundaryIfIndex, "ipMRouteBoundaryIfIndex") + types.AddKeyToken(ipMRouteBoundaryEntry.IpMRouteBoundaryAddress, "ipMRouteBoundaryAddress") + types.AddKeyToken(ipMRouteBoundaryEntry.IpMRouteBoundaryAddressMask, "ipMRouteBoundaryAddressMask")
    ipMRouteBoundaryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteBoundaryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteBoundaryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteBoundaryEntry.EntityData.Children = types.NewOrderedMap()
    ipMRouteBoundaryEntry.EntityData.Leafs = types.NewOrderedMap()
    ipMRouteBoundaryEntry.EntityData.Leafs.Append("ipMRouteBoundaryIfIndex", types.YLeaf{"IpMRouteBoundaryIfIndex", ipMRouteBoundaryEntry.IpMRouteBoundaryIfIndex})
    ipMRouteBoundaryEntry.EntityData.Leafs.Append("ipMRouteBoundaryAddress", types.YLeaf{"IpMRouteBoundaryAddress", ipMRouteBoundaryEntry.IpMRouteBoundaryAddress})
    ipMRouteBoundaryEntry.EntityData.Leafs.Append("ipMRouteBoundaryAddressMask", types.YLeaf{"IpMRouteBoundaryAddressMask", ipMRouteBoundaryEntry.IpMRouteBoundaryAddressMask})
    ipMRouteBoundaryEntry.EntityData.Leafs.Append("ipMRouteBoundaryStatus", types.YLeaf{"IpMRouteBoundaryStatus", ipMRouteBoundaryEntry.IpMRouteBoundaryStatus})

    ipMRouteBoundaryEntry.EntityData.YListKeys = []string {"IpMRouteBoundaryIfIndex", "IpMRouteBoundaryAddress", "IpMRouteBoundaryAddressMask"}

    return &(ipMRouteBoundaryEntry.EntityData)
}

// IPMROUTESTDMIB_IpMRouteScopeNameTable
// The (conceptual) table listing the multicast scope names.
type IPMROUTESTDMIB_IpMRouteScopeNameTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ipMRouteScopeNameTable representing a
    // multicast scope name. The type is slice of
    // IPMROUTESTDMIB_IpMRouteScopeNameTable_IpMRouteScopeNameEntry.
    IpMRouteScopeNameEntry []*IPMROUTESTDMIB_IpMRouteScopeNameTable_IpMRouteScopeNameEntry
}

func (ipMRouteScopeNameTable *IPMROUTESTDMIB_IpMRouteScopeNameTable) GetEntityData() *types.CommonEntityData {
    ipMRouteScopeNameTable.EntityData.YFilter = ipMRouteScopeNameTable.YFilter
    ipMRouteScopeNameTable.EntityData.YangName = "ipMRouteScopeNameTable"
    ipMRouteScopeNameTable.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteScopeNameTable.EntityData.ParentYangName = "IPMROUTE-STD-MIB"
    ipMRouteScopeNameTable.EntityData.SegmentPath = "ipMRouteScopeNameTable"
    ipMRouteScopeNameTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteScopeNameTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteScopeNameTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteScopeNameTable.EntityData.Children = types.NewOrderedMap()
    ipMRouteScopeNameTable.EntityData.Children.Append("ipMRouteScopeNameEntry", types.YChild{"IpMRouteScopeNameEntry", nil})
    for i := range ipMRouteScopeNameTable.IpMRouteScopeNameEntry {
        ipMRouteScopeNameTable.EntityData.Children.Append(types.GetSegmentPath(ipMRouteScopeNameTable.IpMRouteScopeNameEntry[i]), types.YChild{"IpMRouteScopeNameEntry", ipMRouteScopeNameTable.IpMRouteScopeNameEntry[i]})
    }
    ipMRouteScopeNameTable.EntityData.Leafs = types.NewOrderedMap()

    ipMRouteScopeNameTable.EntityData.YListKeys = []string {}

    return &(ipMRouteScopeNameTable.EntityData)
}

// IPMROUTESTDMIB_IpMRouteScopeNameTable_IpMRouteScopeNameEntry
// An entry (conceptual row) in the ipMRouteScopeNameTable
// representing a multicast scope name.
type IPMROUTESTDMIB_IpMRouteScopeNameTable_IpMRouteScopeNameEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The group address which when combined with the
    // corresponding value of ipMRouteScopeNameAddressMask identifies the group
    // range associated with the multicast scope.  Scoped addresses must come from
    // the range 239.x.x.x. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteScopeNameAddress interface{}

    // This attribute is a key. The group address mask which when combined with
    // the corresponding value of ipMRouteScopeNameAddress identifies the group
    // range associated with the multicast scope. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpMRouteScopeNameAddressMask interface{}

    // This attribute is a key. The RFC 1766-style language tag associated with
    // the scope name. The type is string.
    IpMRouteScopeNameLanguage interface{}

    // The textual name associated with the multicast scope.  The value of this
    // object should be suitable for displaying to end-users, such as when
    // allocating a multicast address in this scope.  When no name is specified,
    // the default value of this object should be the string 239.x.x.x/y with x
    // and y replaced appropriately to describe the address and mask length
    // associated with the scope. The type is string.
    IpMRouteScopeNameString interface{}

    // If true, indicates a preference that the name in the following language
    // should be used by applications if no name is available in a desired
    // language. The type is bool.
    IpMRouteScopeNameDefault interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    IpMRouteScopeNameStatus interface{}
}

func (ipMRouteScopeNameEntry *IPMROUTESTDMIB_IpMRouteScopeNameTable_IpMRouteScopeNameEntry) GetEntityData() *types.CommonEntityData {
    ipMRouteScopeNameEntry.EntityData.YFilter = ipMRouteScopeNameEntry.YFilter
    ipMRouteScopeNameEntry.EntityData.YangName = "ipMRouteScopeNameEntry"
    ipMRouteScopeNameEntry.EntityData.BundleName = "cisco_ios_xe"
    ipMRouteScopeNameEntry.EntityData.ParentYangName = "ipMRouteScopeNameTable"
    ipMRouteScopeNameEntry.EntityData.SegmentPath = "ipMRouteScopeNameEntry" + types.AddKeyToken(ipMRouteScopeNameEntry.IpMRouteScopeNameAddress, "ipMRouteScopeNameAddress") + types.AddKeyToken(ipMRouteScopeNameEntry.IpMRouteScopeNameAddressMask, "ipMRouteScopeNameAddressMask") + types.AddKeyToken(ipMRouteScopeNameEntry.IpMRouteScopeNameLanguage, "ipMRouteScopeNameLanguage")
    ipMRouteScopeNameEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipMRouteScopeNameEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipMRouteScopeNameEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipMRouteScopeNameEntry.EntityData.Children = types.NewOrderedMap()
    ipMRouteScopeNameEntry.EntityData.Leafs = types.NewOrderedMap()
    ipMRouteScopeNameEntry.EntityData.Leafs.Append("ipMRouteScopeNameAddress", types.YLeaf{"IpMRouteScopeNameAddress", ipMRouteScopeNameEntry.IpMRouteScopeNameAddress})
    ipMRouteScopeNameEntry.EntityData.Leafs.Append("ipMRouteScopeNameAddressMask", types.YLeaf{"IpMRouteScopeNameAddressMask", ipMRouteScopeNameEntry.IpMRouteScopeNameAddressMask})
    ipMRouteScopeNameEntry.EntityData.Leafs.Append("ipMRouteScopeNameLanguage", types.YLeaf{"IpMRouteScopeNameLanguage", ipMRouteScopeNameEntry.IpMRouteScopeNameLanguage})
    ipMRouteScopeNameEntry.EntityData.Leafs.Append("ipMRouteScopeNameString", types.YLeaf{"IpMRouteScopeNameString", ipMRouteScopeNameEntry.IpMRouteScopeNameString})
    ipMRouteScopeNameEntry.EntityData.Leafs.Append("ipMRouteScopeNameDefault", types.YLeaf{"IpMRouteScopeNameDefault", ipMRouteScopeNameEntry.IpMRouteScopeNameDefault})
    ipMRouteScopeNameEntry.EntityData.Leafs.Append("ipMRouteScopeNameStatus", types.YLeaf{"IpMRouteScopeNameStatus", ipMRouteScopeNameEntry.IpMRouteScopeNameStatus})

    ipMRouteScopeNameEntry.EntityData.YListKeys = []string {"IpMRouteScopeNameAddress", "IpMRouteScopeNameAddressMask", "IpMRouteScopeNameLanguage"}

    return &(ipMRouteScopeNameEntry.EntityData)
}

