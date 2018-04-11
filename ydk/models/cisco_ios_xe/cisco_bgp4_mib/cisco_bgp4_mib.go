// An extension to the IETF BGP4 MIB module defined in
// RFC 1657.
// 
// Following is the terminology associated with Border
// Gateway Protocol(BGP).
// 
// UPDATE message
//     UPDATE messages are used to transfer routing 
//     information between BGP peers. An UPDATE message 
//     is used to advertise a single feasible route to a
//     peer, or to withdraw multiple unfeasible routes 
//     from service.                 
// 
// Adj-RIBs-In 
//    The Adj-RIBs-In store routing information that has
//    been learned from inbound UPDATE messages. Their 
//    contents represent routes that are available as an 
//    input to the Decision Process.
// 
// Loc-RIB(BGP table) 
//    The Loc-RIB contains the local routing information
//    that the BGP speaker has selected by applying its 
//    local policies to the routing information contained 
//    in its Adj-RIBs-In.
// 
// Adj-RIBs-Out 
//    The Adj-RIBs-Out store the information that the
//    local BGP speaker has selected for advertisement to 
//    its peers. The routing information stored in the 
//    Adj-RIBs-Out will be carried in the local BGP 
//    speaker's UPDATE messages and advertised to its
//    peers.
// 
// Path Attributes
//    A variable length sequence of path attributes is 
//    present in every UPDATE. Each path attribute is a 
//    triple <attribute type, attribute length, 
//    attribute value> of variable length. 
// 
// Network Layer Reachability Information(NLRI)
//    A variable length field present in UPDATE messages
//    which contains a list of Network Layer address 
//    prefixes. 
// 
// Address Family Identifier(AFI) 
//    Primary identifier to indicate the type of the 
//    Network Layer Reachability Information(NLRI) being 
//    carried.
// 
// Subsequent Address Family Identifier(SAFI) 
//    Secondary identifier to indicate the type of the 
//    Network Layer Reachability Information(NLRI) being 
//    carried.
package cisco_bgp4_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_bgp4_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-BGP4-MIB CISCO-BGP4-MIB}", reflect.TypeOf(CISCOBGP4MIB{}))
    ydk.RegisterEntity("CISCO-BGP4-MIB:CISCO-BGP4-MIB", reflect.TypeOf(CISCOBGP4MIB{}))
}

// CbgpSafi represents SAFI values 128 through 255 are for private use.
type CbgpSafi string

const (
    CbgpSafi_unicast CbgpSafi = "unicast"

    CbgpSafi_multicast CbgpSafi = "multicast"

    CbgpSafi_unicastAndMulticast CbgpSafi = "unicastAndMulticast"

    CbgpSafi_vpn CbgpSafi = "vpn"
)

// CISCOBGP4MIB
type CISCOBGP4MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cbgpglobal CISCOBGP4MIB_Cbgpglobal

    // This table contains information about routes to destination networks from
    // all BGP4 peers.  Since  BGP4 can carry routes for multiple Network Layer 
    // protocols, this table has the Address Family  Identifier(AFI) of the
    // Network Layer protocol as the  first index. Further for a given AFI, routes
    // carried by BGP4 are distinguished based on Subsequent Address  Family
    // Identifiers(SAFI).  Hence that is used as the second index.  Conceptually
    // there is a separate Loc-RIB maintained by the BGP speaker for each
    // combination of  AFI and SAFI supported by it.
    Cbgproutetable CISCOBGP4MIB_Cbgproutetable

    // This table contains the capabilities that are supported by a peer.
    // Capabilities of a peer are  received during BGP connection establishment.
    // Values corresponding to each received capability are stored in this table.
    // When a new capability  is received, this table is updated with a new 
    // entry. When an existing capability is not received  during the latest
    // connection establishment, the  corresponding entry is deleted from the
    // table.
    Cbgppeercapstable CISCOBGP4MIB_Cbgppeercapstable

    // This table contains information related to address families supported by a
    // peer. Supported address families of a peer are known during BGP  connection
    // establishment. When a new supported  address family is known, this table is
    // updated  with a new entry. When an address family is not  supported any
    // more, corresponding entry is deleted  from the table.
    Cbgppeeraddrfamilytable CISCOBGP4MIB_Cbgppeeraddrfamilytable

    // This table contains prefix related information related to address families
    // supported by a peer.  Supported address families of a peer are known 
    // during BGP connection establishment. When a new  supported address family
    // is known, this table  is updated with a new entry. When an address  family
    // is not supported any more, corresponding  entry is deleted from the table.
    Cbgppeeraddrfamilyprefixtable CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable

    // BGP peer table.  This table contains, one entry per BGP peer, information
    // about the connections with BGP peers.
    Cbgppeer2Table CISCOBGP4MIB_Cbgppeer2Table

    // This table contains the capabilities that are supported by a peer.
    // Capabilities of a peer are received during BGP connection establishment.
    // Values corresponding to each received capability are stored in this table.
    // When a new capability is received, this table is updated with a new entry.
    // When an existing capability is not received during the latest connection
    // establishment, the corresponding entry is deleted from the table.
    Cbgppeer2Capstable CISCOBGP4MIB_Cbgppeer2Capstable

    // This table contains information related to address families supported by a
    // peer. Supported address families of a peer are known during BGP connection
    // establishment. When a new supported address family is known, this table is
    // updated with a new entry. When an address family is not supported any more,
    // corresponding entry is deleted from the table.
    Cbgppeer2Addrfamilytable CISCOBGP4MIB_Cbgppeer2Addrfamilytable

    // This table contains prefix related information related to address families
    // supported by a peer. Supported address families of a peer are known during
    // BGP connection establishment. When a new supported address family is known,
    // this table is updated with a new entry. When an address family is not
    // supported any more, corresponding entry is deleted from the table.
    Cbgppeer2Addrfamilyprefixtable CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable
}

func (cISCOBGP4MIB *CISCOBGP4MIB) GetEntityData() *types.CommonEntityData {
    cISCOBGP4MIB.EntityData.YFilter = cISCOBGP4MIB.YFilter
    cISCOBGP4MIB.EntityData.YangName = "CISCO-BGP4-MIB"
    cISCOBGP4MIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOBGP4MIB.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cISCOBGP4MIB.EntityData.SegmentPath = "CISCO-BGP4-MIB:CISCO-BGP4-MIB"
    cISCOBGP4MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOBGP4MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOBGP4MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOBGP4MIB.EntityData.Children = make(map[string]types.YChild)
    cISCOBGP4MIB.EntityData.Children["cbgpGlobal"] = types.YChild{"Cbgpglobal", &cISCOBGP4MIB.Cbgpglobal}
    cISCOBGP4MIB.EntityData.Children["cbgpRouteTable"] = types.YChild{"Cbgproutetable", &cISCOBGP4MIB.Cbgproutetable}
    cISCOBGP4MIB.EntityData.Children["cbgpPeerCapsTable"] = types.YChild{"Cbgppeercapstable", &cISCOBGP4MIB.Cbgppeercapstable}
    cISCOBGP4MIB.EntityData.Children["cbgpPeerAddrFamilyTable"] = types.YChild{"Cbgppeeraddrfamilytable", &cISCOBGP4MIB.Cbgppeeraddrfamilytable}
    cISCOBGP4MIB.EntityData.Children["cbgpPeerAddrFamilyPrefixTable"] = types.YChild{"Cbgppeeraddrfamilyprefixtable", &cISCOBGP4MIB.Cbgppeeraddrfamilyprefixtable}
    cISCOBGP4MIB.EntityData.Children["cbgpPeer2Table"] = types.YChild{"Cbgppeer2Table", &cISCOBGP4MIB.Cbgppeer2Table}
    cISCOBGP4MIB.EntityData.Children["cbgpPeer2CapsTable"] = types.YChild{"Cbgppeer2Capstable", &cISCOBGP4MIB.Cbgppeer2Capstable}
    cISCOBGP4MIB.EntityData.Children["cbgpPeer2AddrFamilyTable"] = types.YChild{"Cbgppeer2Addrfamilytable", &cISCOBGP4MIB.Cbgppeer2Addrfamilytable}
    cISCOBGP4MIB.EntityData.Children["cbgpPeer2AddrFamilyPrefixTable"] = types.YChild{"Cbgppeer2Addrfamilyprefixtable", &cISCOBGP4MIB.Cbgppeer2Addrfamilyprefixtable}
    cISCOBGP4MIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOBGP4MIB.EntityData)
}

// CISCOBGP4MIB_Cbgpglobal
type CISCOBGP4MIB_Cbgpglobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether the specific notifications are enabled.  If
    // notifsEnable(0) bit is set to 1, then the notifications defined in
    // ciscoBgp4NotificationsGroup1 are enabled;  If notifsPeer2Enable(1) bit is
    // set to 1, then the notifications defined in
    // ciscoBgp4Peer2NotificationsGroup are enabled. The type is map[string]bool.
    Cbgpnotifsenable interface{}

    // The local autonomous system (AS) number. The type is interface{} with
    // range: 0..4294967295.
    Cbgplocalas interface{}
}

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetEntityData() *types.CommonEntityData {
    cbgpglobal.EntityData.YFilter = cbgpglobal.YFilter
    cbgpglobal.EntityData.YangName = "cbgpGlobal"
    cbgpglobal.EntityData.BundleName = "cisco_ios_xe"
    cbgpglobal.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpglobal.EntityData.SegmentPath = "cbgpGlobal"
    cbgpglobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpglobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpglobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpglobal.EntityData.Children = make(map[string]types.YChild)
    cbgpglobal.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgpglobal.EntityData.Leafs["cbgpNotifsEnable"] = types.YLeaf{"Cbgpnotifsenable", cbgpglobal.Cbgpnotifsenable}
    cbgpglobal.EntityData.Leafs["cbgpLocalAs"] = types.YLeaf{"Cbgplocalas", cbgpglobal.Cbgplocalas}
    return &(cbgpglobal.EntityData)
}

// CISCOBGP4MIB_Cbgproutetable
// This table contains information about routes to
// destination networks from all BGP4 peers.  Since 
// BGP4 can carry routes for multiple Network Layer 
// protocols, this table has the Address Family 
// Identifier(AFI) of the Network Layer protocol as the 
// first index. Further for a given AFI, routes carried
// by BGP4 are distinguished based on Subsequent Address 
// Family Identifiers(SAFI).  Hence that is used as the
// second index.  Conceptually there is a separate Loc-RIB
// maintained by the BGP speaker for each combination of 
// AFI and SAFI supported by it.
type CISCOBGP4MIB_Cbgproutetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a path to a network received from a peer. The type is
    // slice of CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry.
    Cbgprouteentry []CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetEntityData() *types.CommonEntityData {
    cbgproutetable.EntityData.YFilter = cbgproutetable.YFilter
    cbgproutetable.EntityData.YangName = "cbgpRouteTable"
    cbgproutetable.EntityData.BundleName = "cisco_ios_xe"
    cbgproutetable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgproutetable.EntityData.SegmentPath = "cbgpRouteTable"
    cbgproutetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgproutetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgproutetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgproutetable.EntityData.Children = make(map[string]types.YChild)
    cbgproutetable.EntityData.Children["cbgpRouteEntry"] = types.YChild{"Cbgprouteentry", nil}
    for i := range cbgproutetable.Cbgprouteentry {
        cbgproutetable.EntityData.Children[types.GetSegmentPath(&cbgproutetable.Cbgprouteentry[i])] = types.YChild{"Cbgprouteentry", &cbgproutetable.Cbgprouteentry[i]}
    }
    cbgproutetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgproutetable.EntityData)
}

// CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry
// Information about a path to a network received from
// a peer.
type CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Represents Address Family Identifier(AFI) of the
    // Network Layer protocol associated with the route. An implementation is only
    // required to support IPv4 unicast and VPNv4 (Value - 1) address families.
    // The type is InetAddressType.
    Cbgprouteafi interface{}

    // This attribute is a key. Represents Subsequent Address Family
    // Identifier(SAFI) of the route. It gives additional information about the
    // type of the route. An implementation is only  required to support IPv4
    // unicast(Value - 1) and VPNv4( Value - 128) address families. The type is
    // CbgpSafi.
    Cbgproutesafi interface{}

    // This attribute is a key. Represents the type of Network Layer address
    // stored in cbgpRoutePeer. An implementation is only required to support IPv4
    // address type(Value - 1). The type is InetAddressType.
    Cbgproutepeertype interface{}

    // This attribute is a key. The Network Layer address of the peer where the
    // route information was learned. An implementation is only  required to
    // support an IPv4 peer. The type is string with length: 0..255.
    Cbgproutepeer interface{}

    // This attribute is a key. A Network Address prefix in the Network Layer
    // Reachability Information field of BGP UPDATE message. This object is a
    // Network Address containing the prefix with length specified by
    // cbgpRouteAddrPrefixLen. Any bits beyond the length specified by
    // cbgpRouteAddrPrefixLen are zeroed. The type is string with length: 0..255.
    Cbgprouteaddrprefix interface{}

    // This attribute is a key. Length in bits of the Network Address prefix in
    // the Network Layer Reachability Information field. The type is interface{}
    // with range: 0..2040.
    Cbgprouteaddrprefixlen interface{}

    // The ultimate origin of the route information. The type is Cbgprouteorigin.
    Cbgprouteorigin interface{}

    // The sequence of AS path segments.  Each AS path segment is represented by a
    // triple <type, length, value>.  The type is a 1-octet field which has two
    // possible values: 1  AS_SET: unordered set of ASs a route in the           
    // UPDATE message has traversed 2  AS_SEQUENCE: ordered set of ASs a route in
    // the                UPDATE message has traversed.  The length is a 1-octet
    // field containing the number of ASs in the value field.  The value field
    // contains one or more AS numbers, each AS is represented in the octet string
    // as a pair of octets according to the following algorithm: 
    // first-byte-of-pair = ASNumber / 256; second-byte-of-pair = ASNumber & 255;.
    // The type is string with length: 0..255.
    Cbgprouteaspathsegment interface{}

    // The Network Layer address of the border router that should be used for the
    // destination network. The type is string with length: 0..255.
    Cbgproutenexthop interface{}

    // Indicates the presence/absence of MULTI_EXIT_DISC attribute for the route.
    // The type is bool.
    Cbgproutemedpresent interface{}

    // This metric is used to discriminate between multiple exit points to an
    // adjacent autonomous system.  The value of this object is irrelevant if the
    // value of of cbgpRouteMedPresent is false(2). The type is interface{} with
    // range: 0..4294967295.
    Cbgproutemultiexitdisc interface{}

    // Indicates the presence/absence of LOCAL_PREF attribute for the route. The
    // type is bool.
    Cbgproutelocalprefpresent interface{}

    // The degree of preference calculated by the local BGP4 speaker for the
    // route. The value of this object is  irrelevant if the value of
    // cbgpRouteLocalPrefPresent  is false(2). The type is interface{} with range:
    // 0..4294967295.
    Cbgproutelocalpref interface{}

    // Whether or not the local system has selected a less specific route without
    // selecting a more specific route. The type is Cbgprouteatomicaggregate.
    Cbgprouteatomicaggregate interface{}

    // The AS number of the last BGP4 speaker that performed route aggregation.  A
    // value of zero (0) indicates the  absence of this attribute. The type is
    // interface{} with range: 0..65535.
    Cbgprouteaggregatoras interface{}

    // Represents the type of Network Layer address stored in
    // cbgpRouteAggregatorAddr. The type is InetAddressType.
    Cbgprouteaggregatoraddrtype interface{}

    // The Network Layer address of the last BGP4 speaker that performed route
    // aggregation.  A value of all zeros indicates the absence of this attribute.
    // The type is string with length: 0..255.
    Cbgprouteaggregatoraddr interface{}

    // An indication of whether or not this route was chosen as the best BGP4
    // route. The type is bool.
    Cbgproutebest interface{}

    // One or more path attributes not understood by this BGP4 speaker.  Size zero
    // (0) indicates the absence of such attribute(s).  Octets beyond the maximum
    // size, if any, are not recorded by this object.    Each path attribute is a
    // triple <attribute type, attribute length, attribute value> of variable
    // length. Attribute Type is a two-octet field that consists of the Attribute
    // Flags octet followed by the Attribute Type Code octet.  If the Extended
    // Length bit of the  Attribute Flags octet is set to 0, the third octet of 
    // the Path Attribute contains the length of the attribute data in octets.  If
    // the Extended Length bit  of the Attribute Flags octet is set to 1, then the
    // third and the fourth octets of the path attribute  contain the length of
    // the attribute data in octets. The remaining octets of the Path Attribute
    // represent  the attribute value and are interpreted according to  the
    // Attribute Flags and the Attribute Type Code. The type is string with
    // length: 0..255.
    Cbgprouteunknownattr interface{}
}

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetEntityData() *types.CommonEntityData {
    cbgprouteentry.EntityData.YFilter = cbgprouteentry.YFilter
    cbgprouteentry.EntityData.YangName = "cbgpRouteEntry"
    cbgprouteentry.EntityData.BundleName = "cisco_ios_xe"
    cbgprouteentry.EntityData.ParentYangName = "cbgpRouteTable"
    cbgprouteentry.EntityData.SegmentPath = "cbgpRouteEntry" + "[cbgpRouteAfi='" + fmt.Sprintf("%v", cbgprouteentry.Cbgprouteafi) + "']" + "[cbgpRouteSafi='" + fmt.Sprintf("%v", cbgprouteentry.Cbgproutesafi) + "']" + "[cbgpRoutePeerType='" + fmt.Sprintf("%v", cbgprouteentry.Cbgproutepeertype) + "']" + "[cbgpRoutePeer='" + fmt.Sprintf("%v", cbgprouteentry.Cbgproutepeer) + "']" + "[cbgpRouteAddrPrefix='" + fmt.Sprintf("%v", cbgprouteentry.Cbgprouteaddrprefix) + "']" + "[cbgpRouteAddrPrefixLen='" + fmt.Sprintf("%v", cbgprouteentry.Cbgprouteaddrprefixlen) + "']"
    cbgprouteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgprouteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgprouteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgprouteentry.EntityData.Children = make(map[string]types.YChild)
    cbgprouteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgprouteentry.EntityData.Leafs["cbgpRouteAfi"] = types.YLeaf{"Cbgprouteafi", cbgprouteentry.Cbgprouteafi}
    cbgprouteentry.EntityData.Leafs["cbgpRouteSafi"] = types.YLeaf{"Cbgproutesafi", cbgprouteentry.Cbgproutesafi}
    cbgprouteentry.EntityData.Leafs["cbgpRoutePeerType"] = types.YLeaf{"Cbgproutepeertype", cbgprouteentry.Cbgproutepeertype}
    cbgprouteentry.EntityData.Leafs["cbgpRoutePeer"] = types.YLeaf{"Cbgproutepeer", cbgprouteentry.Cbgproutepeer}
    cbgprouteentry.EntityData.Leafs["cbgpRouteAddrPrefix"] = types.YLeaf{"Cbgprouteaddrprefix", cbgprouteentry.Cbgprouteaddrprefix}
    cbgprouteentry.EntityData.Leafs["cbgpRouteAddrPrefixLen"] = types.YLeaf{"Cbgprouteaddrprefixlen", cbgprouteentry.Cbgprouteaddrprefixlen}
    cbgprouteentry.EntityData.Leafs["cbgpRouteOrigin"] = types.YLeaf{"Cbgprouteorigin", cbgprouteentry.Cbgprouteorigin}
    cbgprouteentry.EntityData.Leafs["cbgpRouteASPathSegment"] = types.YLeaf{"Cbgprouteaspathsegment", cbgprouteentry.Cbgprouteaspathsegment}
    cbgprouteentry.EntityData.Leafs["cbgpRouteNextHop"] = types.YLeaf{"Cbgproutenexthop", cbgprouteentry.Cbgproutenexthop}
    cbgprouteentry.EntityData.Leafs["cbgpRouteMedPresent"] = types.YLeaf{"Cbgproutemedpresent", cbgprouteentry.Cbgproutemedpresent}
    cbgprouteentry.EntityData.Leafs["cbgpRouteMultiExitDisc"] = types.YLeaf{"Cbgproutemultiexitdisc", cbgprouteentry.Cbgproutemultiexitdisc}
    cbgprouteentry.EntityData.Leafs["cbgpRouteLocalPrefPresent"] = types.YLeaf{"Cbgproutelocalprefpresent", cbgprouteentry.Cbgproutelocalprefpresent}
    cbgprouteentry.EntityData.Leafs["cbgpRouteLocalPref"] = types.YLeaf{"Cbgproutelocalpref", cbgprouteentry.Cbgproutelocalpref}
    cbgprouteentry.EntityData.Leafs["cbgpRouteAtomicAggregate"] = types.YLeaf{"Cbgprouteatomicaggregate", cbgprouteentry.Cbgprouteatomicaggregate}
    cbgprouteentry.EntityData.Leafs["cbgpRouteAggregatorAS"] = types.YLeaf{"Cbgprouteaggregatoras", cbgprouteentry.Cbgprouteaggregatoras}
    cbgprouteentry.EntityData.Leafs["cbgpRouteAggregatorAddrType"] = types.YLeaf{"Cbgprouteaggregatoraddrtype", cbgprouteentry.Cbgprouteaggregatoraddrtype}
    cbgprouteentry.EntityData.Leafs["cbgpRouteAggregatorAddr"] = types.YLeaf{"Cbgprouteaggregatoraddr", cbgprouteentry.Cbgprouteaggregatoraddr}
    cbgprouteentry.EntityData.Leafs["cbgpRouteBest"] = types.YLeaf{"Cbgproutebest", cbgprouteentry.Cbgproutebest}
    cbgprouteentry.EntityData.Leafs["cbgpRouteUnknownAttr"] = types.YLeaf{"Cbgprouteunknownattr", cbgprouteentry.Cbgprouteunknownattr}
    return &(cbgprouteentry.EntityData)
}

// CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteatomicaggregate represents route.
type CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteatomicaggregate string

const (
    CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteatomicaggregate_lessSpecificRouteNotSelected CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteatomicaggregate = "lessSpecificRouteNotSelected"

    CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteatomicaggregate_lessSpecificRouteSelected CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteatomicaggregate = "lessSpecificRouteSelected"
)

// CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin represents The ultimate origin of the route information.
type CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin string

const (
    CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin_igp CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin = "igp"

    CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin_egp CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin = "egp"

    CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin_incomplete CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry_Cbgprouteorigin = "incomplete"
)

// CISCOBGP4MIB_Cbgppeercapstable
// This table contains the capabilities that are
// supported by a peer. Capabilities of a peer are 
// received during BGP connection establishment.
// Values corresponding to each received capability
// are stored in this table. When a new capability 
// is received, this table is updated with a new 
// entry. When an existing capability is not received 
// during the latest connection establishment, the 
// corresponding entry is deleted from the table.
type CISCOBGP4MIB_Cbgppeercapstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a capability received from a peer with a particular
    // code and an index. When a  capability is received multiple times with
    // different values during a BGP connection establishment,  corresponding
    // entries are differentiated with indices. The type is slice of
    // CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry.
    Cbgppeercapsentry []CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetEntityData() *types.CommonEntityData {
    cbgppeercapstable.EntityData.YFilter = cbgppeercapstable.YFilter
    cbgppeercapstable.EntityData.YangName = "cbgpPeerCapsTable"
    cbgppeercapstable.EntityData.BundleName = "cisco_ios_xe"
    cbgppeercapstable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeercapstable.EntityData.SegmentPath = "cbgpPeerCapsTable"
    cbgppeercapstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeercapstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeercapstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeercapstable.EntityData.Children = make(map[string]types.YChild)
    cbgppeercapstable.EntityData.Children["cbgpPeerCapsEntry"] = types.YChild{"Cbgppeercapsentry", nil}
    for i := range cbgppeercapstable.Cbgppeercapsentry {
        cbgppeercapstable.EntityData.Children[types.GetSegmentPath(&cbgppeercapstable.Cbgppeercapsentry[i])] = types.YChild{"Cbgppeercapsentry", &cbgppeercapstable.Cbgppeercapsentry[i]}
    }
    cbgppeercapstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeercapstable.EntityData)
}

// CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry
// Each entry represents a capability received from a
// peer with a particular code and an index. When a 
// capability is received multiple times with different
// values during a BGP connection establishment, 
// corresponding entries are differentiated with indices.
type CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to bgp4_mib.BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerremoteaddr
    Bgppeerremoteaddr interface{}

    // This attribute is a key. The BGP Capability Advertisement Capability Code.
    // The type is Cbgppeercapcode.
    Cbgppeercapcode interface{}

    // This attribute is a key. Multiple instances of a given capability may be
    // sent by a BGP speaker.  This variable is used to index them. The type is
    // interface{} with range: 1..128.
    Cbgppeercapindex interface{}

    // The value of the announced capability. This MIB object value is organized
    // as given below,     Capability : Route Refresh Capability                 
    // Null string     Capability : Multiprotocol Extensions      
    // +----------------------------------+       | AFI(16 bits)                  
    // |       +----------------------------------+       | SAFI (8 bits)         
    // |       +----------------------------------+     Capability : Graceful
    // Restart       +----------------------------------+       | Restart Flags (4
    // bits)           |       +----------------------------------+       |
    // Restart Time in seconds (12 bits)|      
    // +----------------------------------+       | AFI(16 bits)                  
    // |       +----------------------------------+       | SAFI (8 bits)         
    // |       +----------------------------------+       | Flags for Address
    // Family (8 bits)|       +----------------------------------+       | ...    
    // |       +----------------------------------+       | AFI(16 bits)          
    // |       +----------------------------------+       | SAFI (8 bits)         
    // |       +----------------------------------+       | Flags for Address
    // Family (8 bits)|       +----------------------------------+. The type is
    // string with length: 0..255.
    Cbgppeercapvalue interface{}
}

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetEntityData() *types.CommonEntityData {
    cbgppeercapsentry.EntityData.YFilter = cbgppeercapsentry.YFilter
    cbgppeercapsentry.EntityData.YangName = "cbgpPeerCapsEntry"
    cbgppeercapsentry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeercapsentry.EntityData.ParentYangName = "cbgpPeerCapsTable"
    cbgppeercapsentry.EntityData.SegmentPath = "cbgpPeerCapsEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", cbgppeercapsentry.Bgppeerremoteaddr) + "']" + "[cbgpPeerCapCode='" + fmt.Sprintf("%v", cbgppeercapsentry.Cbgppeercapcode) + "']" + "[cbgpPeerCapIndex='" + fmt.Sprintf("%v", cbgppeercapsentry.Cbgppeercapindex) + "']"
    cbgppeercapsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeercapsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeercapsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeercapsentry.EntityData.Children = make(map[string]types.YChild)
    cbgppeercapsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeercapsentry.EntityData.Leafs["bgpPeerRemoteAddr"] = types.YLeaf{"Bgppeerremoteaddr", cbgppeercapsentry.Bgppeerremoteaddr}
    cbgppeercapsentry.EntityData.Leafs["cbgpPeerCapCode"] = types.YLeaf{"Cbgppeercapcode", cbgppeercapsentry.Cbgppeercapcode}
    cbgppeercapsentry.EntityData.Leafs["cbgpPeerCapIndex"] = types.YLeaf{"Cbgppeercapindex", cbgppeercapsentry.Cbgppeercapindex}
    cbgppeercapsentry.EntityData.Leafs["cbgpPeerCapValue"] = types.YLeaf{"Cbgppeercapvalue", cbgppeercapsentry.Cbgppeercapvalue}
    return &(cbgppeercapsentry.EntityData)
}

// CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode represents The BGP Capability Advertisement Capability Code.
type CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode string

const (
    CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode_multiProtocol CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode = "multiProtocol"

    CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode_routeRefresh CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode = "routeRefresh"

    CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode_gracefulRestart CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode = "gracefulRestart"

    CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode_routeRefreshOld CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry_Cbgppeercapcode = "routeRefreshOld"
)

// CISCOBGP4MIB_Cbgppeeraddrfamilytable
// This table contains information related to
// address families supported by a peer. Supported
// address families of a peer are known during BGP 
// connection establishment. When a new supported 
// address family is known, this table is updated 
// with a new entry. When an address family is not 
// supported any more, corresponding entry is deleted 
// from the table.
type CISCOBGP4MIB_Cbgppeeraddrfamilytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // names associated with an address family. The type is slice of
    // CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry.
    Cbgppeeraddrfamilyentry []CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetEntityData() *types.CommonEntityData {
    cbgppeeraddrfamilytable.EntityData.YFilter = cbgppeeraddrfamilytable.YFilter
    cbgppeeraddrfamilytable.EntityData.YangName = "cbgpPeerAddrFamilyTable"
    cbgppeeraddrfamilytable.EntityData.BundleName = "cisco_ios_xe"
    cbgppeeraddrfamilytable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeeraddrfamilytable.EntityData.SegmentPath = "cbgpPeerAddrFamilyTable"
    cbgppeeraddrfamilytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeeraddrfamilytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeeraddrfamilytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeeraddrfamilytable.EntityData.Children = make(map[string]types.YChild)
    cbgppeeraddrfamilytable.EntityData.Children["cbgpPeerAddrFamilyEntry"] = types.YChild{"Cbgppeeraddrfamilyentry", nil}
    for i := range cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry {
        cbgppeeraddrfamilytable.EntityData.Children[types.GetSegmentPath(&cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry[i])] = types.YChild{"Cbgppeeraddrfamilyentry", &cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry[i]}
    }
    cbgppeeraddrfamilytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeeraddrfamilytable.EntityData)
}

// CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains names associated with
// an address family.
type CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to bgp4_mib.BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerremoteaddr
    Bgppeerremoteaddr interface{}

    // This attribute is a key. The AFI index of the entry. An implementation is
    // only required to support IPv4 unicast and  VPNv4 (Value - 1) address
    // families. The type is InetAddressType.
    Cbgppeeraddrfamilyafi interface{}

    // This attribute is a key. The SAFI index of the entry. An implementation is
    // only required to support IPv4 unicast(Value  - 1) and VPNv4( Value - 128)
    // address families. The type is CbgpSafi.
    Cbgppeeraddrfamilysafi interface{}

    // Implementation specific Address Family name. The type is string.
    Cbgppeeraddrfamilyname interface{}
}

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetEntityData() *types.CommonEntityData {
    cbgppeeraddrfamilyentry.EntityData.YFilter = cbgppeeraddrfamilyentry.YFilter
    cbgppeeraddrfamilyentry.EntityData.YangName = "cbgpPeerAddrFamilyEntry"
    cbgppeeraddrfamilyentry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeeraddrfamilyentry.EntityData.ParentYangName = "cbgpPeerAddrFamilyTable"
    cbgppeeraddrfamilyentry.EntityData.SegmentPath = "cbgpPeerAddrFamilyEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", cbgppeeraddrfamilyentry.Bgppeerremoteaddr) + "']" + "[cbgpPeerAddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilyafi) + "']" + "[cbgpPeerAddrFamilySafi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilysafi) + "']"
    cbgppeeraddrfamilyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeeraddrfamilyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeeraddrfamilyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeeraddrfamilyentry.EntityData.Children = make(map[string]types.YChild)
    cbgppeeraddrfamilyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeeraddrfamilyentry.EntityData.Leafs["bgpPeerRemoteAddr"] = types.YLeaf{"Bgppeerremoteaddr", cbgppeeraddrfamilyentry.Bgppeerremoteaddr}
    cbgppeeraddrfamilyentry.EntityData.Leafs["cbgpPeerAddrFamilyAfi"] = types.YLeaf{"Cbgppeeraddrfamilyafi", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilyafi}
    cbgppeeraddrfamilyentry.EntityData.Leafs["cbgpPeerAddrFamilySafi"] = types.YLeaf{"Cbgppeeraddrfamilysafi", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilysafi}
    cbgppeeraddrfamilyentry.EntityData.Leafs["cbgpPeerAddrFamilyName"] = types.YLeaf{"Cbgppeeraddrfamilyname", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilyname}
    return &(cbgppeeraddrfamilyentry.EntityData)
}

// CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable
// This table contains prefix related information
// related to address families supported by a peer. 
// Supported address families of a peer are known 
// during BGP connection establishment. When a new 
// supported address family is known, this table 
// is updated with a new entry. When an address 
// family is not supported any more, corresponding 
// entry is deleted from the table.
type CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // information associated  with route prefixes belonging to an address family.
    // The type is slice of
    // CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry.
    Cbgppeeraddrfamilyprefixentry []CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetEntityData() *types.CommonEntityData {
    cbgppeeraddrfamilyprefixtable.EntityData.YFilter = cbgppeeraddrfamilyprefixtable.YFilter
    cbgppeeraddrfamilyprefixtable.EntityData.YangName = "cbgpPeerAddrFamilyPrefixTable"
    cbgppeeraddrfamilyprefixtable.EntityData.BundleName = "cisco_ios_xe"
    cbgppeeraddrfamilyprefixtable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeeraddrfamilyprefixtable.EntityData.SegmentPath = "cbgpPeerAddrFamilyPrefixTable"
    cbgppeeraddrfamilyprefixtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeeraddrfamilyprefixtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeeraddrfamilyprefixtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeeraddrfamilyprefixtable.EntityData.Children = make(map[string]types.YChild)
    cbgppeeraddrfamilyprefixtable.EntityData.Children["cbgpPeerAddrFamilyPrefixEntry"] = types.YChild{"Cbgppeeraddrfamilyprefixentry", nil}
    for i := range cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry {
        cbgppeeraddrfamilyprefixtable.EntityData.Children[types.GetSegmentPath(&cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry[i])] = types.YChild{"Cbgppeeraddrfamilyprefixentry", &cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry[i]}
    }
    cbgppeeraddrfamilyprefixtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeeraddrfamilyprefixtable.EntityData)
}

// CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains information associated 
// with route prefixes belonging to an address family.
type CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to bgp4_mib.BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerremoteaddr
    Bgppeerremoteaddr interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry_Cbgppeeraddrfamilyafi
    Cbgppeeraddrfamilyafi interface{}

    // This attribute is a key. The type is CbgpSafi. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry_Cbgppeeraddrfamilysafi
    Cbgppeeraddrfamilysafi interface{}

    // Number of accepted route prefixes on this connection, which belong to an
    // address family. The type is interface{} with range: 0..4294967295.
    Cbgppeeracceptedprefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, received on this  connection is denied. It is initialized
    // to zero when  the connection is undergone a hard reset. The type is
    // interface{} with range: 0..4294967295.
    Cbgppeerdeniedprefixes interface{}

    // Max number of route prefixes accepted for an address family on this
    // connection. The type is interface{} with range: 1..4294967295.
    Cbgppeerprefixadminlimit interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which warning message stating the prefix count is crossed the threshold or 
    // corresponding SNMP notification is generated. The type is interface{} with
    // range: 1..100.
    Cbgppeerprefixthreshold interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which SNMP clear notification is generated if prefix threshold notification
    // is already generated. The type is interface{} with range: 1..100.
    Cbgppeerprefixclearthreshold interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is advertised on this connection. It is initialized to zero
    // when  the connection is undergone a hard reset. The type is interface{}
    // with range: 0..4294967295.
    Cbgppeeradvertisedprefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is suppressed from being sent on this connection. It is 
    // initialized to zero when the connection is undergone a hard reset. The type
    // is interface{} with range: 0..4294967295.
    Cbgppeersuppressedprefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, is withdrawn on this connection. It is initialized to zero
    // when the connection is undergone a hard reset. The type is interface{} with
    // range: 0..4294967295.
    Cbgppeerwithdrawnprefixes interface{}
}

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetEntityData() *types.CommonEntityData {
    cbgppeeraddrfamilyprefixentry.EntityData.YFilter = cbgppeeraddrfamilyprefixentry.YFilter
    cbgppeeraddrfamilyprefixentry.EntityData.YangName = "cbgpPeerAddrFamilyPrefixEntry"
    cbgppeeraddrfamilyprefixentry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeeraddrfamilyprefixentry.EntityData.ParentYangName = "cbgpPeerAddrFamilyPrefixTable"
    cbgppeeraddrfamilyprefixentry.EntityData.SegmentPath = "cbgpPeerAddrFamilyPrefixEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", cbgppeeraddrfamilyprefixentry.Bgppeerremoteaddr) + "']" + "[cbgpPeerAddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilyafi) + "']" + "[cbgpPeerAddrFamilySafi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilysafi) + "']"
    cbgppeeraddrfamilyprefixentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeeraddrfamilyprefixentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeeraddrfamilyprefixentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeeraddrfamilyprefixentry.EntityData.Children = make(map[string]types.YChild)
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["bgpPeerRemoteAddr"] = types.YLeaf{"Bgppeerremoteaddr", cbgppeeraddrfamilyprefixentry.Bgppeerremoteaddr}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerAddrFamilyAfi"] = types.YLeaf{"Cbgppeeraddrfamilyafi", cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilyafi}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerAddrFamilySafi"] = types.YLeaf{"Cbgppeeraddrfamilysafi", cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilysafi}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerAcceptedPrefixes"] = types.YLeaf{"Cbgppeeracceptedprefixes", cbgppeeraddrfamilyprefixentry.Cbgppeeracceptedprefixes}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerDeniedPrefixes"] = types.YLeaf{"Cbgppeerdeniedprefixes", cbgppeeraddrfamilyprefixentry.Cbgppeerdeniedprefixes}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerPrefixAdminLimit"] = types.YLeaf{"Cbgppeerprefixadminlimit", cbgppeeraddrfamilyprefixentry.Cbgppeerprefixadminlimit}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerPrefixThreshold"] = types.YLeaf{"Cbgppeerprefixthreshold", cbgppeeraddrfamilyprefixentry.Cbgppeerprefixthreshold}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerPrefixClearThreshold"] = types.YLeaf{"Cbgppeerprefixclearthreshold", cbgppeeraddrfamilyprefixentry.Cbgppeerprefixclearthreshold}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerAdvertisedPrefixes"] = types.YLeaf{"Cbgppeeradvertisedprefixes", cbgppeeraddrfamilyprefixentry.Cbgppeeradvertisedprefixes}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerSuppressedPrefixes"] = types.YLeaf{"Cbgppeersuppressedprefixes", cbgppeeraddrfamilyprefixentry.Cbgppeersuppressedprefixes}
    cbgppeeraddrfamilyprefixentry.EntityData.Leafs["cbgpPeerWithdrawnPrefixes"] = types.YLeaf{"Cbgppeerwithdrawnprefixes", cbgppeeraddrfamilyprefixentry.Cbgppeerwithdrawnprefixes}
    return &(cbgppeeraddrfamilyprefixentry.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Table
// BGP peer table.  This table contains,
// one entry per BGP peer, information about
// the connections with BGP peers.
type CISCOBGP4MIB_Cbgppeer2Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry containing information about the connection with a BGP peer. The type
    // is slice of CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry.
    Cbgppeer2Entry []CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetEntityData() *types.CommonEntityData {
    cbgppeer2Table.EntityData.YFilter = cbgppeer2Table.YFilter
    cbgppeer2Table.EntityData.YangName = "cbgpPeer2Table"
    cbgppeer2Table.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Table.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeer2Table.EntityData.SegmentPath = "cbgpPeer2Table"
    cbgppeer2Table.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Table.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Table.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Table.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Table.EntityData.Children["cbgpPeer2Entry"] = types.YChild{"Cbgppeer2Entry", nil}
    for i := range cbgppeer2Table.Cbgppeer2Entry {
        cbgppeer2Table.EntityData.Children[types.GetSegmentPath(&cbgppeer2Table.Cbgppeer2Entry[i])] = types.YChild{"Cbgppeer2Entry", &cbgppeer2Table.Cbgppeer2Entry[i]}
    }
    cbgppeer2Table.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeer2Table.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry
// Entry containing information about the
// connection with a BGP peer.
type CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Represents the type of Peer address stored in
    // cbgpPeer2Entry. The type is InetAddressType.
    Cbgppeer2Type interface{}

    // This attribute is a key. The remote IP address of this entry's BGP peer.
    // The type is string with length: 0..255.
    Cbgppeer2Remoteaddr interface{}

    // The BGP peer connection state. The type is Cbgppeer2State.
    Cbgppeer2State interface{}

    // The desired state of the BGP connection. A transition from 'stop' to
    // 'start' will cause the BGP Manual Start Event to be generated. A transition
    // from 'start' to 'stop' will cause the BGP Manual Stop Event to be
    // generated. This parameter can be used to restart BGP peer connections. 
    // Care should be used in providing write access to this object without
    // adequate authentication. The type is Cbgppeer2Adminstatus.
    Cbgppeer2Adminstatus interface{}

    // The negotiated version of BGP running between the two peers.  This entry
    // MUST be zero (0) unless the cbgpPeer2State is in the openconfirm or the
    // established state.  Note that legal values for this object are between 0
    // and 255. The type is interface{} with range: -2147483648..2147483647.
    Cbgppeer2Negotiatedversion interface{}

    // The local IP address of this entry's BGP connection. The type is string
    // with length: 0..255.
    Cbgppeer2Localaddr interface{}

    // The local port for the TCP connection between the BGP peers. The type is
    // interface{} with range: 0..65535.
    Cbgppeer2Localport interface{}

    // The local AS number for this session. The type is interface{} with range:
    // 0..4294967295.
    Cbgppeer2Localas interface{}

    // The BGP Identifier of this entry's BGP peer. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cbgppeer2Localidentifier interface{}

    // The remote port for the TCP connection between the BGP peers.  Note that
    // the objects cbgpPeer2LocalAddr, cbgpPeer2LocalPort, cbgpPeer2RemoteAddr,
    // and cbgpPeer2RemotePort provide the appropriate reference to the standard
    // MIB TCP connection table. The type is interface{} with range: 0..65535.
    Cbgppeer2Remoteport interface{}

    // The remote autonomous system number received in the BGP OPEN message. The
    // type is interface{} with range: 0..4294967295.
    Cbgppeer2Remoteas interface{}

    // The BGP Identifier of this entry's BGP peer. This entry MUST be 0.0.0.0
    // unless the cbgpPeer2State is in the openconfirm or the established state.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cbgppeer2Remoteidentifier interface{}

    // The number of BGP UPDATE messages received on this connection. The type is
    // interface{} with range: 0..4294967295.
    Cbgppeer2Inupdates interface{}

    // The number of BGP UPDATE messages transmitted on this connection. The type
    // is interface{} with range: 0..4294967295.
    Cbgppeer2Outupdates interface{}

    // The total number of messages received from the remote peer on this
    // connection. The type is interface{} with range: 0..4294967295.
    Cbgppeer2Intotalmessages interface{}

    // The total number of messages transmitted to the remote peer on this
    // connection. The type is interface{} with range: 0..4294967295.
    Cbgppeer2Outtotalmessages interface{}

    // The last error code and subcode seen by this peer on this connection.  If
    // no error has occurred, this field is zero.  Otherwise, the first byte of
    // this two byte OCTET STRING contains the error code, and the second byte
    // contains the subcode. The type is string with length: 2.
    Cbgppeer2Lasterror interface{}

    // The total number of times the BGP FSM transitioned into the established
    // state for this peer. The type is interface{} with range: 0..4294967295.
    Cbgppeer2Fsmestablishedtransitions interface{}

    // This timer indicates how long (in seconds) this peer has been in the
    // established state or how long since this peer was last in the established
    // state.  It is set to zero when a new peer is configured or when the router
    // is booted. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Cbgppeer2Fsmestablishedtime interface{}

    // Time interval (in seconds) for the ConnectRetry timer.  The suggested value
    // for this timer is 120 seconds. The type is interface{} with range:
    // 1..65535. Units are seconds.
    Cbgppeer2Connectretryinterval interface{}

    // Time interval (in seconds) for the Hold Timer established with the peer. 
    // The value of this object is calculated by this BGP speaker, using the
    // smaller of the values in cbgpPeer2HoldTimeConfigured and the Hold Time
    // received in the OPEN message.  This value must be at least three seconds if
    // it is not zero (0).  If the Hold Timer has not been established with the
    // peer this object MUST have a value of zero (0).  If the
    // cbgpPeer2HoldTimeConfigured object has a value of (0), then this object
    // MUST have a value of (0). The type is interface{} with range: 0..None |
    // 3..65535. Units are seconds.
    Cbgppeer2Holdtime interface{}

    // Time interval (in seconds) for the KeepAlive timer established with the
    // peer.  The value of this object is calculated by this BGP speaker such
    // that, when compared with cbgpPeer2HoldTime, it has the same proportion that
    // cbgpPeer2KeepAliveConfigured has, compared with
    // cbgpPeer2HoldTimeConfigured.  If the KeepAlive timer has not been
    // established with the peer, this object MUST have a value of zero (0).  If
    // the of cbgpPeer2KeepAliveConfigured object has a value of (0), then this
    // object MUST have a value of (0). The type is interface{} with range:
    // 0..21845. Units are seconds.
    Cbgppeer2Keepalive interface{}

    // Time interval (in seconds) for the Hold Time configured for this BGP
    // speaker with this peer.  This value is placed in an OPEN message sent to
    // this peer by this BGP speaker, and is compared with the Hold Time field in
    // an OPEN message received from the peer when determining the Hold Time
    // (cbgpPeer2HoldTime) with the peer. This value must not be less than three
    // seconds if it is not zero (0).  If it is zero (0), the Hold Time is NOT to
    // be established with the peer.  The suggested value for this timer is 90
    // seconds. The type is interface{} with range: 0..None | 3..65535. Units are
    // seconds.
    Cbgppeer2Holdtimeconfigured interface{}

    // Time interval (in seconds) for the KeepAlive timer configured for this BGP
    // speaker with this peer.  The value of this object will only determine the
    // KEEPALIVE messages' frequency relative to the value specified in
    // cbgpPeer2HoldTimeConfigured; the actual time interval for the KEEPALIVE
    // messages is indicated by cbgpPeer2KeepAlive.  A reasonable maximum value
    // for this timer would be one third of that of cbgpPeer2HoldTimeConfigured.
    // If the value of this object is zero (0), no periodical KEEPALIVE messages
    // are sent to the peer after the BGP connection has been established.  The
    // suggested value for this timer is 30 seconds. The type is interface{} with
    // range: 0..21845. Units are seconds.
    Cbgppeer2Keepaliveconfigured interface{}

    // Time interval (in seconds) for the MinASOriginationInterval timer. The
    // suggested value for this timer is 15 seconds. The type is interface{} with
    // range: 1..65535. Units are seconds.
    Cbgppeer2Minasoriginationinterval interface{}

    // Time interval (in seconds) for the MinRouteAdvertisementInterval timer. The
    // suggested value for this timer is 30 seconds for EBGP connections and 5
    // seconds for IBGP connections. The type is interface{} with range: 1..65535.
    // Units are seconds.
    Cbgppeer2Minrouteadvertisementinterval interface{}

    // Elapsed time (in seconds) since the last BGP UPDATE message was received
    // from the peer. Each time cbgpPeer2InUpdates is incremented, the value of
    // this object is set to zero (0). The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    Cbgppeer2Inupdateelapsedtime interface{}

    // Implementation specific error description for bgpPeerLastErrorReceived. The
    // type is string.
    Cbgppeer2Lasterrortxt interface{}

    // The BGP peer connection previous state. The type is Cbgppeer2Prevstate.
    Cbgppeer2Prevstate interface{}
}

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetEntityData() *types.CommonEntityData {
    cbgppeer2Entry.EntityData.YFilter = cbgppeer2Entry.YFilter
    cbgppeer2Entry.EntityData.YangName = "cbgpPeer2Entry"
    cbgppeer2Entry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Entry.EntityData.ParentYangName = "cbgpPeer2Table"
    cbgppeer2Entry.EntityData.SegmentPath = "cbgpPeer2Entry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Entry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Entry.Cbgppeer2Remoteaddr) + "']"
    cbgppeer2Entry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Entry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Entry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Entry.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Entry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2Type"] = types.YLeaf{"Cbgppeer2Type", cbgppeer2Entry.Cbgppeer2Type}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2RemoteAddr"] = types.YLeaf{"Cbgppeer2Remoteaddr", cbgppeer2Entry.Cbgppeer2Remoteaddr}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2State"] = types.YLeaf{"Cbgppeer2State", cbgppeer2Entry.Cbgppeer2State}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2AdminStatus"] = types.YLeaf{"Cbgppeer2Adminstatus", cbgppeer2Entry.Cbgppeer2Adminstatus}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2NegotiatedVersion"] = types.YLeaf{"Cbgppeer2Negotiatedversion", cbgppeer2Entry.Cbgppeer2Negotiatedversion}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2LocalAddr"] = types.YLeaf{"Cbgppeer2Localaddr", cbgppeer2Entry.Cbgppeer2Localaddr}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2LocalPort"] = types.YLeaf{"Cbgppeer2Localport", cbgppeer2Entry.Cbgppeer2Localport}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2LocalAs"] = types.YLeaf{"Cbgppeer2Localas", cbgppeer2Entry.Cbgppeer2Localas}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2LocalIdentifier"] = types.YLeaf{"Cbgppeer2Localidentifier", cbgppeer2Entry.Cbgppeer2Localidentifier}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2RemotePort"] = types.YLeaf{"Cbgppeer2Remoteport", cbgppeer2Entry.Cbgppeer2Remoteport}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2RemoteAs"] = types.YLeaf{"Cbgppeer2Remoteas", cbgppeer2Entry.Cbgppeer2Remoteas}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2RemoteIdentifier"] = types.YLeaf{"Cbgppeer2Remoteidentifier", cbgppeer2Entry.Cbgppeer2Remoteidentifier}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2InUpdates"] = types.YLeaf{"Cbgppeer2Inupdates", cbgppeer2Entry.Cbgppeer2Inupdates}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2OutUpdates"] = types.YLeaf{"Cbgppeer2Outupdates", cbgppeer2Entry.Cbgppeer2Outupdates}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2InTotalMessages"] = types.YLeaf{"Cbgppeer2Intotalmessages", cbgppeer2Entry.Cbgppeer2Intotalmessages}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2OutTotalMessages"] = types.YLeaf{"Cbgppeer2Outtotalmessages", cbgppeer2Entry.Cbgppeer2Outtotalmessages}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2LastError"] = types.YLeaf{"Cbgppeer2Lasterror", cbgppeer2Entry.Cbgppeer2Lasterror}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2FsmEstablishedTransitions"] = types.YLeaf{"Cbgppeer2Fsmestablishedtransitions", cbgppeer2Entry.Cbgppeer2Fsmestablishedtransitions}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2FsmEstablishedTime"] = types.YLeaf{"Cbgppeer2Fsmestablishedtime", cbgppeer2Entry.Cbgppeer2Fsmestablishedtime}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2ConnectRetryInterval"] = types.YLeaf{"Cbgppeer2Connectretryinterval", cbgppeer2Entry.Cbgppeer2Connectretryinterval}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2HoldTime"] = types.YLeaf{"Cbgppeer2Holdtime", cbgppeer2Entry.Cbgppeer2Holdtime}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2KeepAlive"] = types.YLeaf{"Cbgppeer2Keepalive", cbgppeer2Entry.Cbgppeer2Keepalive}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2HoldTimeConfigured"] = types.YLeaf{"Cbgppeer2Holdtimeconfigured", cbgppeer2Entry.Cbgppeer2Holdtimeconfigured}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2KeepAliveConfigured"] = types.YLeaf{"Cbgppeer2Keepaliveconfigured", cbgppeer2Entry.Cbgppeer2Keepaliveconfigured}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2MinASOriginationInterval"] = types.YLeaf{"Cbgppeer2Minasoriginationinterval", cbgppeer2Entry.Cbgppeer2Minasoriginationinterval}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2MinRouteAdvertisementInterval"] = types.YLeaf{"Cbgppeer2Minrouteadvertisementinterval", cbgppeer2Entry.Cbgppeer2Minrouteadvertisementinterval}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2InUpdateElapsedTime"] = types.YLeaf{"Cbgppeer2Inupdateelapsedtime", cbgppeer2Entry.Cbgppeer2Inupdateelapsedtime}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2LastErrorTxt"] = types.YLeaf{"Cbgppeer2Lasterrortxt", cbgppeer2Entry.Cbgppeer2Lasterrortxt}
    cbgppeer2Entry.EntityData.Leafs["cbgpPeer2PrevState"] = types.YLeaf{"Cbgppeer2Prevstate", cbgppeer2Entry.Cbgppeer2Prevstate}
    return &(cbgppeer2Entry.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Adminstatus represents authentication.
type CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Adminstatus string

const (
    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Adminstatus_stop CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Adminstatus = "stop"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Adminstatus_start CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Adminstatus = "start"
)

// CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate represents The BGP peer connection previous state.
type CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate string

const (
    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_none CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "none"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_idle CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "idle"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_connect CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "connect"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_active CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "active"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_opensent CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "opensent"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_openconfirm CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "openconfirm"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate_established CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Prevstate = "established"
)

// CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State represents The BGP peer connection state.
type CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State string

const (
    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State_idle CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State = "idle"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State_connect CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State = "connect"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State_active CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State = "active"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State_opensent CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State = "opensent"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State_openconfirm CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State = "openconfirm"

    CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State_established CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2State = "established"
)

// CISCOBGP4MIB_Cbgppeer2Capstable
// This table contains the capabilities that are
// supported by a peer. Capabilities of a peer are
// received during BGP connection establishment.
// Values corresponding to each received capability
// are stored in this table. When a new capability
// is received, this table is updated with a new
// entry. When an existing capability is not received
// during the latest connection establishment, the
// corresponding entry is deleted from the table.
type CISCOBGP4MIB_Cbgppeer2Capstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a capability received from a peer with a particular
    // code and an index. When a capability is received multiple times with
    // different values during a BGP connection establishment, corresponding
    // entries are differentiated with indices. The type is slice of
    // CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry.
    Cbgppeer2Capsentry []CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetEntityData() *types.CommonEntityData {
    cbgppeer2Capstable.EntityData.YFilter = cbgppeer2Capstable.YFilter
    cbgppeer2Capstable.EntityData.YangName = "cbgpPeer2CapsTable"
    cbgppeer2Capstable.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Capstable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeer2Capstable.EntityData.SegmentPath = "cbgpPeer2CapsTable"
    cbgppeer2Capstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Capstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Capstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Capstable.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Capstable.EntityData.Children["cbgpPeer2CapsEntry"] = types.YChild{"Cbgppeer2Capsentry", nil}
    for i := range cbgppeer2Capstable.Cbgppeer2Capsentry {
        cbgppeer2Capstable.EntityData.Children[types.GetSegmentPath(&cbgppeer2Capstable.Cbgppeer2Capsentry[i])] = types.YChild{"Cbgppeer2Capsentry", &cbgppeer2Capstable.Cbgppeer2Capsentry[i]}
    }
    cbgppeer2Capstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeer2Capstable.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry
// Each entry represents a capability received from a
// peer with a particular code and an index. When a
// capability is received multiple times with different
// values during a BGP connection establishment,
// corresponding entries are differentiated with indices.
type CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Type
    Cbgppeer2Type interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Remoteaddr
    Cbgppeer2Remoteaddr interface{}

    // This attribute is a key. The BGP Capability Advertisement Capability Code.
    // The type is Cbgppeer2Capcode.
    Cbgppeer2Capcode interface{}

    // This attribute is a key. Multiple instances of a given capability may be
    // sent by a BGP speaker.  This variable is used to index them. The type is
    // interface{} with range: 1..128.
    Cbgppeer2Capindex interface{}

    // The value of the announced capability. This MIB object value is organized
    // as given below,     Capability : Route Refresh Capability                 
    // 4-Byte AS Capability                  Null string     Capability :
    // Multiprotocol Extensions       +----------------------------------+       |
    // AFI(16 bits)                     |      
    // +----------------------------------+       | SAFI (8 bits)                 
    // |       +----------------------------------+     Capability : Graceful
    // Restart       +----------------------------------+       | Restart Flags (4
    // bits)           |       +----------------------------------+       |
    // Restart Time in seconds (12 bits)|      
    // +----------------------------------+       | AFI(16 bits)                  
    // |       +----------------------------------+       | SAFI (8 bits)         
    // |       +----------------------------------+       | Flags for Address
    // Family (8 bits)|       +----------------------------------+       | ...    
    // |       +----------------------------------+       | AFI(16 bits)          
    // |       +----------------------------------+       | SAFI (8 bits)         
    // |       +----------------------------------+       | Flags for Address
    // Family (8 bits)|       +----------------------------------+     Capability
    // : Additional Paths       +----------------------------------+       |
    // AFI(16 bits)                     |      
    // +----------------------------------+       | SAFI (8 bits)                 
    // |       +----------------------------------+       | Send/Receive (8 bits) 
    // |       +----------------------------------+. The type is string with
    // length: 0..255.
    Cbgppeer2Capvalue interface{}
}

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetEntityData() *types.CommonEntityData {
    cbgppeer2Capsentry.EntityData.YFilter = cbgppeer2Capsentry.YFilter
    cbgppeer2Capsentry.EntityData.YangName = "cbgpPeer2CapsEntry"
    cbgppeer2Capsentry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Capsentry.EntityData.ParentYangName = "cbgpPeer2CapsTable"
    cbgppeer2Capsentry.EntityData.SegmentPath = "cbgpPeer2CapsEntry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Remoteaddr) + "']" + "[cbgpPeer2CapCode='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Capcode) + "']" + "[cbgpPeer2CapIndex='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Capindex) + "']"
    cbgppeer2Capsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Capsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Capsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Capsentry.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Capsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeer2Capsentry.EntityData.Leafs["cbgpPeer2Type"] = types.YLeaf{"Cbgppeer2Type", cbgppeer2Capsentry.Cbgppeer2Type}
    cbgppeer2Capsentry.EntityData.Leafs["cbgpPeer2RemoteAddr"] = types.YLeaf{"Cbgppeer2Remoteaddr", cbgppeer2Capsentry.Cbgppeer2Remoteaddr}
    cbgppeer2Capsentry.EntityData.Leafs["cbgpPeer2CapCode"] = types.YLeaf{"Cbgppeer2Capcode", cbgppeer2Capsentry.Cbgppeer2Capcode}
    cbgppeer2Capsentry.EntityData.Leafs["cbgpPeer2CapIndex"] = types.YLeaf{"Cbgppeer2Capindex", cbgppeer2Capsentry.Cbgppeer2Capindex}
    cbgppeer2Capsentry.EntityData.Leafs["cbgpPeer2CapValue"] = types.YLeaf{"Cbgppeer2Capvalue", cbgppeer2Capsentry.Cbgppeer2Capvalue}
    return &(cbgppeer2Capsentry.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode represents The BGP Capability Advertisement Capability Code.
type CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode string

const (
    CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode_multiProtocol CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode = "multiProtocol"

    CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode_routeRefresh CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode = "routeRefresh"

    CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode_gracefulRestart CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode = "gracefulRestart"

    CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode_fourByteAs CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode = "fourByteAs"

    CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode_addPath CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode = "addPath"

    CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode_routeRefreshOld CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry_Cbgppeer2Capcode = "routeRefreshOld"
)

// CISCOBGP4MIB_Cbgppeer2Addrfamilytable
// This table contains information related to
// address families supported by a peer. Supported
// address families of a peer are known during BGP
// connection establishment. When a new supported
// address family is known, this table is updated
// with a new entry. When an address family is not
// supported any more, corresponding entry is deleted
// from the table.
type CISCOBGP4MIB_Cbgppeer2Addrfamilytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // names associated with an address family. The type is slice of
    // CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry.
    Cbgppeer2Addrfamilyentry []CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetEntityData() *types.CommonEntityData {
    cbgppeer2Addrfamilytable.EntityData.YFilter = cbgppeer2Addrfamilytable.YFilter
    cbgppeer2Addrfamilytable.EntityData.YangName = "cbgpPeer2AddrFamilyTable"
    cbgppeer2Addrfamilytable.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Addrfamilytable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeer2Addrfamilytable.EntityData.SegmentPath = "cbgpPeer2AddrFamilyTable"
    cbgppeer2Addrfamilytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Addrfamilytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Addrfamilytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Addrfamilytable.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Addrfamilytable.EntityData.Children["cbgpPeer2AddrFamilyEntry"] = types.YChild{"Cbgppeer2Addrfamilyentry", nil}
    for i := range cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry {
        cbgppeer2Addrfamilytable.EntityData.Children[types.GetSegmentPath(&cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry[i])] = types.YChild{"Cbgppeer2Addrfamilyentry", &cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry[i]}
    }
    cbgppeer2Addrfamilytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeer2Addrfamilytable.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains names associated with
// an address family.
type CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Type
    Cbgppeer2Type interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Remoteaddr
    Cbgppeer2Remoteaddr interface{}

    // This attribute is a key. The AFI index of the entry. An implementation is
    // only required to support IPv4 unicast and VPNv4 (Value - 1) address
    // families. The type is InetAddressType.
    Cbgppeer2Addrfamilyafi interface{}

    // This attribute is a key. The SAFI index of the entry. An implementation is
    // only required to support IPv4 unicast(Value - 1) and VPNv4( Value - 128)
    // address families. The type is CbgpSafi.
    Cbgppeer2Addrfamilysafi interface{}

    // Implementation specific Address Family name. The type is string.
    Cbgppeer2Addrfamilyname interface{}
}

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetEntityData() *types.CommonEntityData {
    cbgppeer2Addrfamilyentry.EntityData.YFilter = cbgppeer2Addrfamilyentry.YFilter
    cbgppeer2Addrfamilyentry.EntityData.YangName = "cbgpPeer2AddrFamilyEntry"
    cbgppeer2Addrfamilyentry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Addrfamilyentry.EntityData.ParentYangName = "cbgpPeer2AddrFamilyTable"
    cbgppeer2Addrfamilyentry.EntityData.SegmentPath = "cbgpPeer2AddrFamilyEntry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Remoteaddr) + "']" + "[cbgpPeer2AddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilyafi) + "']" + "[cbgpPeer2AddrFamilySafi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilysafi) + "']"
    cbgppeer2Addrfamilyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Addrfamilyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Addrfamilyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Addrfamilyentry.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Addrfamilyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeer2Addrfamilyentry.EntityData.Leafs["cbgpPeer2Type"] = types.YLeaf{"Cbgppeer2Type", cbgppeer2Addrfamilyentry.Cbgppeer2Type}
    cbgppeer2Addrfamilyentry.EntityData.Leafs["cbgpPeer2RemoteAddr"] = types.YLeaf{"Cbgppeer2Remoteaddr", cbgppeer2Addrfamilyentry.Cbgppeer2Remoteaddr}
    cbgppeer2Addrfamilyentry.EntityData.Leafs["cbgpPeer2AddrFamilyAfi"] = types.YLeaf{"Cbgppeer2Addrfamilyafi", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilyafi}
    cbgppeer2Addrfamilyentry.EntityData.Leafs["cbgpPeer2AddrFamilySafi"] = types.YLeaf{"Cbgppeer2Addrfamilysafi", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilysafi}
    cbgppeer2Addrfamilyentry.EntityData.Leafs["cbgpPeer2AddrFamilyName"] = types.YLeaf{"Cbgppeer2Addrfamilyname", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilyname}
    return &(cbgppeer2Addrfamilyentry.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable
// This table contains prefix related information
// related to address families supported by a peer.
// Supported address families of a peer are known
// during BGP connection establishment. When a new
// supported address family is known, this table
// is updated with a new entry. When an address
// family is not supported any more, corresponding
// entry is deleted from the table.
type CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // information associated with route prefixes belonging to an address family.
    // The type is slice of
    // CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry.
    Cbgppeer2Addrfamilyprefixentry []CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetEntityData() *types.CommonEntityData {
    cbgppeer2Addrfamilyprefixtable.EntityData.YFilter = cbgppeer2Addrfamilyprefixtable.YFilter
    cbgppeer2Addrfamilyprefixtable.EntityData.YangName = "cbgpPeer2AddrFamilyPrefixTable"
    cbgppeer2Addrfamilyprefixtable.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Addrfamilyprefixtable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgppeer2Addrfamilyprefixtable.EntityData.SegmentPath = "cbgpPeer2AddrFamilyPrefixTable"
    cbgppeer2Addrfamilyprefixtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Addrfamilyprefixtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Addrfamilyprefixtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Addrfamilyprefixtable.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Addrfamilyprefixtable.EntityData.Children["cbgpPeer2AddrFamilyPrefixEntry"] = types.YChild{"Cbgppeer2Addrfamilyprefixentry", nil}
    for i := range cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry {
        cbgppeer2Addrfamilyprefixtable.EntityData.Children[types.GetSegmentPath(&cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry[i])] = types.YChild{"Cbgppeer2Addrfamilyprefixentry", &cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry[i]}
    }
    cbgppeer2Addrfamilyprefixtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbgppeer2Addrfamilyprefixtable.EntityData)
}

// CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains information associated
// with route prefixes belonging to an address family.
type CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Type
    Cbgppeer2Type interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry_Cbgppeer2Remoteaddr
    Cbgppeer2Remoteaddr interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry_Cbgppeer2Addrfamilyafi
    Cbgppeer2Addrfamilyafi interface{}

    // This attribute is a key. The type is CbgpSafi. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry_Cbgppeer2Addrfamilysafi
    Cbgppeer2Addrfamilysafi interface{}

    // Number of accepted route prefixes on this connection, which belong to an
    // address family. The type is interface{} with range: 0..4294967295.
    Cbgppeer2Acceptedprefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, received on this connection is denied. It is initialized to
    // zero when the connection is undergone a hard reset. The type is interface{}
    // with range: 0..4294967295.
    Cbgppeer2Deniedprefixes interface{}

    // Max number of route prefixes accepted for an address family on this
    // connection. The type is interface{} with range: 1..4294967295.
    Cbgppeer2Prefixadminlimit interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which warning message stating the prefix count is crossed the threshold or
    // corresponding SNMP notification is generated. The type is interface{} with
    // range: 1..100. Units are percent.
    Cbgppeer2Prefixthreshold interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which SNMP clear notification is generated if prefix threshold notification
    // is already generated. The type is interface{} with range: 1..100. Units are
    // percent.
    Cbgppeer2Prefixclearthreshold interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is advertised on this connection. It is initialized to zero
    // when the connection is undergone a hard reset. The type is interface{} with
    // range: 0..4294967295.
    Cbgppeer2Advertisedprefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is suppressed from being sent on this connection. It is
    // initialized to zero when the connection is undergone a hard reset. The type
    // is interface{} with range: 0..4294967295.
    Cbgppeer2Suppressedprefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, is withdrawn on this connection. It is initialized to zero
    // when the connection is undergone a hard reset. The type is interface{} with
    // range: 0..4294967295.
    Cbgppeer2Withdrawnprefixes interface{}
}

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetEntityData() *types.CommonEntityData {
    cbgppeer2Addrfamilyprefixentry.EntityData.YFilter = cbgppeer2Addrfamilyprefixentry.YFilter
    cbgppeer2Addrfamilyprefixentry.EntityData.YangName = "cbgpPeer2AddrFamilyPrefixEntry"
    cbgppeer2Addrfamilyprefixentry.EntityData.BundleName = "cisco_ios_xe"
    cbgppeer2Addrfamilyprefixentry.EntityData.ParentYangName = "cbgpPeer2AddrFamilyPrefixTable"
    cbgppeer2Addrfamilyprefixentry.EntityData.SegmentPath = "cbgpPeer2AddrFamilyPrefixEntry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Remoteaddr) + "']" + "[cbgpPeer2AddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilyafi) + "']" + "[cbgpPeer2AddrFamilySafi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilysafi) + "']"
    cbgppeer2Addrfamilyprefixentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgppeer2Addrfamilyprefixentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgppeer2Addrfamilyprefixentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgppeer2Addrfamilyprefixentry.EntityData.Children = make(map[string]types.YChild)
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2Type"] = types.YLeaf{"Cbgppeer2Type", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Type}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2RemoteAddr"] = types.YLeaf{"Cbgppeer2Remoteaddr", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Remoteaddr}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2AddrFamilyAfi"] = types.YLeaf{"Cbgppeer2Addrfamilyafi", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilyafi}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2AddrFamilySafi"] = types.YLeaf{"Cbgppeer2Addrfamilysafi", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilysafi}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2AcceptedPrefixes"] = types.YLeaf{"Cbgppeer2Acceptedprefixes", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Acceptedprefixes}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2DeniedPrefixes"] = types.YLeaf{"Cbgppeer2Deniedprefixes", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Deniedprefixes}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2PrefixAdminLimit"] = types.YLeaf{"Cbgppeer2Prefixadminlimit", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Prefixadminlimit}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2PrefixThreshold"] = types.YLeaf{"Cbgppeer2Prefixthreshold", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Prefixthreshold}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2PrefixClearThreshold"] = types.YLeaf{"Cbgppeer2Prefixclearthreshold", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Prefixclearthreshold}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2AdvertisedPrefixes"] = types.YLeaf{"Cbgppeer2Advertisedprefixes", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Advertisedprefixes}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2SuppressedPrefixes"] = types.YLeaf{"Cbgppeer2Suppressedprefixes", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Suppressedprefixes}
    cbgppeer2Addrfamilyprefixentry.EntityData.Leafs["cbgpPeer2WithdrawnPrefixes"] = types.YLeaf{"Cbgppeer2Withdrawnprefixes", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Withdrawnprefixes}
    return &(cbgppeer2Addrfamilyprefixentry.EntityData)
}

