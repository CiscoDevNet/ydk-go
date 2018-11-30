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

    
    CbgpGlobal CISCOBGP4MIB_CbgpGlobal

    // This table contains information about routes to destination networks from
    // all BGP4 peers.  Since  BGP4 can carry routes for multiple Network Layer 
    // protocols, this table has the Address Family  Identifier(AFI) of the
    // Network Layer protocol as the  first index. Further for a given AFI, routes
    // carried by BGP4 are distinguished based on Subsequent Address  Family
    // Identifiers(SAFI).  Hence that is used as the second index.  Conceptually
    // there is a separate Loc-RIB maintained by the BGP speaker for each
    // combination of  AFI and SAFI supported by it.
    CbgpRouteTable CISCOBGP4MIB_CbgpRouteTable

    // This table contains the capabilities that are supported by a peer.
    // Capabilities of a peer are  received during BGP connection establishment.
    // Values corresponding to each received capability are stored in this table.
    // When a new capability  is received, this table is updated with a new 
    // entry. When an existing capability is not received  during the latest
    // connection establishment, the  corresponding entry is deleted from the
    // table.
    CbgpPeerCapsTable CISCOBGP4MIB_CbgpPeerCapsTable

    // This table contains information related to address families supported by a
    // peer. Supported address families of a peer are known during BGP  connection
    // establishment. When a new supported  address family is known, this table is
    // updated  with a new entry. When an address family is not  supported any
    // more, corresponding entry is deleted  from the table.
    CbgpPeerAddrFamilyTable CISCOBGP4MIB_CbgpPeerAddrFamilyTable

    // This table contains prefix related information related to address families
    // supported by a peer.  Supported address families of a peer are known 
    // during BGP connection establishment. When a new  supported address family
    // is known, this table  is updated with a new entry. When an address  family
    // is not supported any more, corresponding  entry is deleted from the table.
    CbgpPeerAddrFamilyPrefixTable CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable

    // BGP peer table.  This table contains, one entry per BGP peer, information
    // about the connections with BGP peers.
    CbgpPeer2Table CISCOBGP4MIB_CbgpPeer2Table

    // This table contains the capabilities that are supported by a peer.
    // Capabilities of a peer are received during BGP connection establishment.
    // Values corresponding to each received capability are stored in this table.
    // When a new capability is received, this table is updated with a new entry.
    // When an existing capability is not received during the latest connection
    // establishment, the corresponding entry is deleted from the table.
    CbgpPeer2CapsTable CISCOBGP4MIB_CbgpPeer2CapsTable

    // This table contains information related to address families supported by a
    // peer. Supported address families of a peer are known during BGP connection
    // establishment. When a new supported address family is known, this table is
    // updated with a new entry. When an address family is not supported any more,
    // corresponding entry is deleted from the table.
    CbgpPeer2AddrFamilyTable CISCOBGP4MIB_CbgpPeer2AddrFamilyTable

    // This table contains prefix related information related to address families
    // supported by a peer. Supported address families of a peer are known during
    // BGP connection establishment. When a new supported address family is known,
    // this table is updated with a new entry. When an address family is not
    // supported any more, corresponding entry is deleted from the table.
    CbgpPeer2AddrFamilyPrefixTable CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable
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

    cISCOBGP4MIB.EntityData.Children = types.NewOrderedMap()
    cISCOBGP4MIB.EntityData.Children.Append("cbgpGlobal", types.YChild{"CbgpGlobal", &cISCOBGP4MIB.CbgpGlobal})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpRouteTable", types.YChild{"CbgpRouteTable", &cISCOBGP4MIB.CbgpRouteTable})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeerCapsTable", types.YChild{"CbgpPeerCapsTable", &cISCOBGP4MIB.CbgpPeerCapsTable})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeerAddrFamilyTable", types.YChild{"CbgpPeerAddrFamilyTable", &cISCOBGP4MIB.CbgpPeerAddrFamilyTable})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeerAddrFamilyPrefixTable", types.YChild{"CbgpPeerAddrFamilyPrefixTable", &cISCOBGP4MIB.CbgpPeerAddrFamilyPrefixTable})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeer2Table", types.YChild{"CbgpPeer2Table", &cISCOBGP4MIB.CbgpPeer2Table})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeer2CapsTable", types.YChild{"CbgpPeer2CapsTable", &cISCOBGP4MIB.CbgpPeer2CapsTable})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeer2AddrFamilyTable", types.YChild{"CbgpPeer2AddrFamilyTable", &cISCOBGP4MIB.CbgpPeer2AddrFamilyTable})
    cISCOBGP4MIB.EntityData.Children.Append("cbgpPeer2AddrFamilyPrefixTable", types.YChild{"CbgpPeer2AddrFamilyPrefixTable", &cISCOBGP4MIB.CbgpPeer2AddrFamilyPrefixTable})
    cISCOBGP4MIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOBGP4MIB.EntityData.YListKeys = []string {}

    return &(cISCOBGP4MIB.EntityData)
}

// CISCOBGP4MIB_CbgpGlobal
type CISCOBGP4MIB_CbgpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether the specific notifications are enabled.  If
    // notifsEnable(0) bit is set to 1, then the notifications defined in
    // ciscoBgp4NotificationsGroup1 are enabled;  If notifsPeer2Enable(1) bit is
    // set to 1, then the notifications defined in
    // ciscoBgp4Peer2NotificationsGroup are enabled. The type is map[string]bool.
    CbgpNotifsEnable interface{}

    // The local autonomous system (AS) number. The type is interface{} with
    // range: 0..4294967295.
    CbgpLocalAs interface{}
}

func (cbgpGlobal *CISCOBGP4MIB_CbgpGlobal) GetEntityData() *types.CommonEntityData {
    cbgpGlobal.EntityData.YFilter = cbgpGlobal.YFilter
    cbgpGlobal.EntityData.YangName = "cbgpGlobal"
    cbgpGlobal.EntityData.BundleName = "cisco_ios_xe"
    cbgpGlobal.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpGlobal.EntityData.SegmentPath = "cbgpGlobal"
    cbgpGlobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpGlobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpGlobal.EntityData.Children = types.NewOrderedMap()
    cbgpGlobal.EntityData.Leafs = types.NewOrderedMap()
    cbgpGlobal.EntityData.Leafs.Append("cbgpNotifsEnable", types.YLeaf{"CbgpNotifsEnable", cbgpGlobal.CbgpNotifsEnable})
    cbgpGlobal.EntityData.Leafs.Append("cbgpLocalAs", types.YLeaf{"CbgpLocalAs", cbgpGlobal.CbgpLocalAs})

    cbgpGlobal.EntityData.YListKeys = []string {}

    return &(cbgpGlobal.EntityData)
}

// CISCOBGP4MIB_CbgpRouteTable
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
type CISCOBGP4MIB_CbgpRouteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a path to a network received from a peer. The type is
    // slice of CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry.
    CbgpRouteEntry []*CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry
}

func (cbgpRouteTable *CISCOBGP4MIB_CbgpRouteTable) GetEntityData() *types.CommonEntityData {
    cbgpRouteTable.EntityData.YFilter = cbgpRouteTable.YFilter
    cbgpRouteTable.EntityData.YangName = "cbgpRouteTable"
    cbgpRouteTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpRouteTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpRouteTable.EntityData.SegmentPath = "cbgpRouteTable"
    cbgpRouteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpRouteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpRouteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpRouteTable.EntityData.Children = types.NewOrderedMap()
    cbgpRouteTable.EntityData.Children.Append("cbgpRouteEntry", types.YChild{"CbgpRouteEntry", nil})
    for i := range cbgpRouteTable.CbgpRouteEntry {
        cbgpRouteTable.EntityData.Children.Append(types.GetSegmentPath(cbgpRouteTable.CbgpRouteEntry[i]), types.YChild{"CbgpRouteEntry", cbgpRouteTable.CbgpRouteEntry[i]})
    }
    cbgpRouteTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpRouteTable.EntityData.YListKeys = []string {}

    return &(cbgpRouteTable.EntityData)
}

// CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry
// Information about a path to a network received from
// a peer.
type CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Represents Address Family Identifier(AFI) of the
    // Network Layer protocol associated with the route. An implementation is only
    // required to support IPv4 unicast and VPNv4 (Value - 1) address families.
    // The type is InetAddressType.
    CbgpRouteAfi interface{}

    // This attribute is a key. Represents Subsequent Address Family
    // Identifier(SAFI) of the route. It gives additional information about the
    // type of the route. An implementation is only  required to support IPv4
    // unicast(Value - 1) and VPNv4( Value - 128) address families. The type is
    // CbgpSafi.
    CbgpRouteSafi interface{}

    // This attribute is a key. Represents the type of Network Layer address
    // stored in cbgpRoutePeer. An implementation is only required to support IPv4
    // address type(Value - 1). The type is InetAddressType.
    CbgpRoutePeerType interface{}

    // This attribute is a key. The Network Layer address of the peer where the
    // route information was learned. An implementation is only  required to
    // support an IPv4 peer. The type is string with length: 0..255.
    CbgpRoutePeer interface{}

    // This attribute is a key. A Network Address prefix in the Network Layer
    // Reachability Information field of BGP UPDATE message. This object is a
    // Network Address containing the prefix with length specified by
    // cbgpRouteAddrPrefixLen. Any bits beyond the length specified by
    // cbgpRouteAddrPrefixLen are zeroed. The type is string with length: 0..255.
    CbgpRouteAddrPrefix interface{}

    // This attribute is a key. Length in bits of the Network Address prefix in
    // the Network Layer Reachability Information field. The type is interface{}
    // with range: 0..2040.
    CbgpRouteAddrPrefixLen interface{}

    // The ultimate origin of the route information. The type is CbgpRouteOrigin.
    CbgpRouteOrigin interface{}

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
    CbgpRouteASPathSegment interface{}

    // The Network Layer address of the border router that should be used for the
    // destination network. The type is string with length: 0..255.
    CbgpRouteNextHop interface{}

    // Indicates the presence/absence of MULTI_EXIT_DISC attribute for the route.
    // The type is bool.
    CbgpRouteMedPresent interface{}

    // This metric is used to discriminate between multiple exit points to an
    // adjacent autonomous system.  The value of this object is irrelevant if the
    // value of of cbgpRouteMedPresent is false(2). The type is interface{} with
    // range: 0..4294967295.
    CbgpRouteMultiExitDisc interface{}

    // Indicates the presence/absence of LOCAL_PREF attribute for the route. The
    // type is bool.
    CbgpRouteLocalPrefPresent interface{}

    // The degree of preference calculated by the local BGP4 speaker for the
    // route. The value of this object is  irrelevant if the value of
    // cbgpRouteLocalPrefPresent  is false(2). The type is interface{} with range:
    // 0..4294967295.
    CbgpRouteLocalPref interface{}

    // Whether or not the local system has selected a less specific route without
    // selecting a more specific route. The type is CbgpRouteAtomicAggregate.
    CbgpRouteAtomicAggregate interface{}

    // The AS number of the last BGP4 speaker that performed route aggregation.  A
    // value of zero (0) indicates the  absence of this attribute. The type is
    // interface{} with range: 0..65535.
    CbgpRouteAggregatorAS interface{}

    // Represents the type of Network Layer address stored in
    // cbgpRouteAggregatorAddr. The type is InetAddressType.
    CbgpRouteAggregatorAddrType interface{}

    // The Network Layer address of the last BGP4 speaker that performed route
    // aggregation.  A value of all zeros indicates the absence of this attribute.
    // The type is string with length: 0..255.
    CbgpRouteAggregatorAddr interface{}

    // An indication of whether or not this route was chosen as the best BGP4
    // route. The type is bool.
    CbgpRouteBest interface{}

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
    CbgpRouteUnknownAttr interface{}
}

func (cbgpRouteEntry *CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry) GetEntityData() *types.CommonEntityData {
    cbgpRouteEntry.EntityData.YFilter = cbgpRouteEntry.YFilter
    cbgpRouteEntry.EntityData.YangName = "cbgpRouteEntry"
    cbgpRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpRouteEntry.EntityData.ParentYangName = "cbgpRouteTable"
    cbgpRouteEntry.EntityData.SegmentPath = "cbgpRouteEntry" + types.AddKeyToken(cbgpRouteEntry.CbgpRouteAfi, "cbgpRouteAfi") + types.AddKeyToken(cbgpRouteEntry.CbgpRouteSafi, "cbgpRouteSafi") + types.AddKeyToken(cbgpRouteEntry.CbgpRoutePeerType, "cbgpRoutePeerType") + types.AddKeyToken(cbgpRouteEntry.CbgpRoutePeer, "cbgpRoutePeer") + types.AddKeyToken(cbgpRouteEntry.CbgpRouteAddrPrefix, "cbgpRouteAddrPrefix") + types.AddKeyToken(cbgpRouteEntry.CbgpRouteAddrPrefixLen, "cbgpRouteAddrPrefixLen")
    cbgpRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpRouteEntry.EntityData.Children = types.NewOrderedMap()
    cbgpRouteEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAfi", types.YLeaf{"CbgpRouteAfi", cbgpRouteEntry.CbgpRouteAfi})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteSafi", types.YLeaf{"CbgpRouteSafi", cbgpRouteEntry.CbgpRouteSafi})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRoutePeerType", types.YLeaf{"CbgpRoutePeerType", cbgpRouteEntry.CbgpRoutePeerType})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRoutePeer", types.YLeaf{"CbgpRoutePeer", cbgpRouteEntry.CbgpRoutePeer})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAddrPrefix", types.YLeaf{"CbgpRouteAddrPrefix", cbgpRouteEntry.CbgpRouteAddrPrefix})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAddrPrefixLen", types.YLeaf{"CbgpRouteAddrPrefixLen", cbgpRouteEntry.CbgpRouteAddrPrefixLen})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteOrigin", types.YLeaf{"CbgpRouteOrigin", cbgpRouteEntry.CbgpRouteOrigin})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteASPathSegment", types.YLeaf{"CbgpRouteASPathSegment", cbgpRouteEntry.CbgpRouteASPathSegment})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteNextHop", types.YLeaf{"CbgpRouteNextHop", cbgpRouteEntry.CbgpRouteNextHop})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteMedPresent", types.YLeaf{"CbgpRouteMedPresent", cbgpRouteEntry.CbgpRouteMedPresent})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteMultiExitDisc", types.YLeaf{"CbgpRouteMultiExitDisc", cbgpRouteEntry.CbgpRouteMultiExitDisc})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteLocalPrefPresent", types.YLeaf{"CbgpRouteLocalPrefPresent", cbgpRouteEntry.CbgpRouteLocalPrefPresent})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteLocalPref", types.YLeaf{"CbgpRouteLocalPref", cbgpRouteEntry.CbgpRouteLocalPref})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAtomicAggregate", types.YLeaf{"CbgpRouteAtomicAggregate", cbgpRouteEntry.CbgpRouteAtomicAggregate})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAggregatorAS", types.YLeaf{"CbgpRouteAggregatorAS", cbgpRouteEntry.CbgpRouteAggregatorAS})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAggregatorAddrType", types.YLeaf{"CbgpRouteAggregatorAddrType", cbgpRouteEntry.CbgpRouteAggregatorAddrType})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteAggregatorAddr", types.YLeaf{"CbgpRouteAggregatorAddr", cbgpRouteEntry.CbgpRouteAggregatorAddr})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteBest", types.YLeaf{"CbgpRouteBest", cbgpRouteEntry.CbgpRouteBest})
    cbgpRouteEntry.EntityData.Leafs.Append("cbgpRouteUnknownAttr", types.YLeaf{"CbgpRouteUnknownAttr", cbgpRouteEntry.CbgpRouteUnknownAttr})

    cbgpRouteEntry.EntityData.YListKeys = []string {"CbgpRouteAfi", "CbgpRouteSafi", "CbgpRoutePeerType", "CbgpRoutePeer", "CbgpRouteAddrPrefix", "CbgpRouteAddrPrefixLen"}

    return &(cbgpRouteEntry.EntityData)
}

// CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteAtomicAggregate represents route.
type CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteAtomicAggregate string

const (
    CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteAtomicAggregate_lessSpecificRouteNotSelected CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteAtomicAggregate = "lessSpecificRouteNotSelected"

    CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteAtomicAggregate_lessSpecificRouteSelected CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteAtomicAggregate = "lessSpecificRouteSelected"
)

// CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin represents The ultimate origin of the route information.
type CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin string

const (
    CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin_igp CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin = "igp"

    CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin_egp CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin = "egp"

    CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin_incomplete CISCOBGP4MIB_CbgpRouteTable_CbgpRouteEntry_CbgpRouteOrigin = "incomplete"
)

// CISCOBGP4MIB_CbgpPeerCapsTable
// This table contains the capabilities that are
// supported by a peer. Capabilities of a peer are 
// received during BGP connection establishment.
// Values corresponding to each received capability
// are stored in this table. When a new capability 
// is received, this table is updated with a new 
// entry. When an existing capability is not received 
// during the latest connection establishment, the 
// corresponding entry is deleted from the table.
type CISCOBGP4MIB_CbgpPeerCapsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a capability received from a peer with a particular
    // code and an index. When a  capability is received multiple times with
    // different values during a BGP connection establishment,  corresponding
    // entries are differentiated with indices. The type is slice of
    // CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry.
    CbgpPeerCapsEntry []*CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry
}

func (cbgpPeerCapsTable *CISCOBGP4MIB_CbgpPeerCapsTable) GetEntityData() *types.CommonEntityData {
    cbgpPeerCapsTable.EntityData.YFilter = cbgpPeerCapsTable.YFilter
    cbgpPeerCapsTable.EntityData.YangName = "cbgpPeerCapsTable"
    cbgpPeerCapsTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeerCapsTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeerCapsTable.EntityData.SegmentPath = "cbgpPeerCapsTable"
    cbgpPeerCapsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeerCapsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeerCapsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeerCapsTable.EntityData.Children = types.NewOrderedMap()
    cbgpPeerCapsTable.EntityData.Children.Append("cbgpPeerCapsEntry", types.YChild{"CbgpPeerCapsEntry", nil})
    for i := range cbgpPeerCapsTable.CbgpPeerCapsEntry {
        cbgpPeerCapsTable.EntityData.Children.Append(types.GetSegmentPath(cbgpPeerCapsTable.CbgpPeerCapsEntry[i]), types.YChild{"CbgpPeerCapsEntry", cbgpPeerCapsTable.CbgpPeerCapsEntry[i]})
    }
    cbgpPeerCapsTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeerCapsTable.EntityData.YListKeys = []string {}

    return &(cbgpPeerCapsTable.EntityData)
}

// CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry
// Each entry represents a capability received from a
// peer with a particular code and an index. When a 
// capability is received multiple times with different
// values during a BGP connection establishment, 
// corresponding entries are differentiated with indices.
type CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to bgp4_mib.BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerRemoteAddr
    BgpPeerRemoteAddr interface{}

    // This attribute is a key. The BGP Capability Advertisement Capability Code.
    // The type is CbgpPeerCapCode.
    CbgpPeerCapCode interface{}

    // This attribute is a key. Multiple instances of a given capability may be
    // sent by a BGP speaker.  This variable is used to index them. The type is
    // interface{} with range: 1..128.
    CbgpPeerCapIndex interface{}

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
    CbgpPeerCapValue interface{}
}

func (cbgpPeerCapsEntry *CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry) GetEntityData() *types.CommonEntityData {
    cbgpPeerCapsEntry.EntityData.YFilter = cbgpPeerCapsEntry.YFilter
    cbgpPeerCapsEntry.EntityData.YangName = "cbgpPeerCapsEntry"
    cbgpPeerCapsEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeerCapsEntry.EntityData.ParentYangName = "cbgpPeerCapsTable"
    cbgpPeerCapsEntry.EntityData.SegmentPath = "cbgpPeerCapsEntry" + types.AddKeyToken(cbgpPeerCapsEntry.BgpPeerRemoteAddr, "bgpPeerRemoteAddr") + types.AddKeyToken(cbgpPeerCapsEntry.CbgpPeerCapCode, "cbgpPeerCapCode") + types.AddKeyToken(cbgpPeerCapsEntry.CbgpPeerCapIndex, "cbgpPeerCapIndex")
    cbgpPeerCapsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeerCapsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeerCapsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeerCapsEntry.EntityData.Children = types.NewOrderedMap()
    cbgpPeerCapsEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeerCapsEntry.EntityData.Leafs.Append("bgpPeerRemoteAddr", types.YLeaf{"BgpPeerRemoteAddr", cbgpPeerCapsEntry.BgpPeerRemoteAddr})
    cbgpPeerCapsEntry.EntityData.Leafs.Append("cbgpPeerCapCode", types.YLeaf{"CbgpPeerCapCode", cbgpPeerCapsEntry.CbgpPeerCapCode})
    cbgpPeerCapsEntry.EntityData.Leafs.Append("cbgpPeerCapIndex", types.YLeaf{"CbgpPeerCapIndex", cbgpPeerCapsEntry.CbgpPeerCapIndex})
    cbgpPeerCapsEntry.EntityData.Leafs.Append("cbgpPeerCapValue", types.YLeaf{"CbgpPeerCapValue", cbgpPeerCapsEntry.CbgpPeerCapValue})

    cbgpPeerCapsEntry.EntityData.YListKeys = []string {"BgpPeerRemoteAddr", "CbgpPeerCapCode", "CbgpPeerCapIndex"}

    return &(cbgpPeerCapsEntry.EntityData)
}

// CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode represents The BGP Capability Advertisement Capability Code.
type CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode string

const (
    CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode_multiProtocol CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode = "multiProtocol"

    CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode_routeRefresh CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode = "routeRefresh"

    CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode_gracefulRestart CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode = "gracefulRestart"

    CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode_routeRefreshOld CISCOBGP4MIB_CbgpPeerCapsTable_CbgpPeerCapsEntry_CbgpPeerCapCode = "routeRefreshOld"
)

// CISCOBGP4MIB_CbgpPeerAddrFamilyTable
// This table contains information related to
// address families supported by a peer. Supported
// address families of a peer are known during BGP 
// connection establishment. When a new supported 
// address family is known, this table is updated 
// with a new entry. When an address family is not 
// supported any more, corresponding entry is deleted 
// from the table.
type CISCOBGP4MIB_CbgpPeerAddrFamilyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // names associated with an address family. The type is slice of
    // CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry.
    CbgpPeerAddrFamilyEntry []*CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry
}

func (cbgpPeerAddrFamilyTable *CISCOBGP4MIB_CbgpPeerAddrFamilyTable) GetEntityData() *types.CommonEntityData {
    cbgpPeerAddrFamilyTable.EntityData.YFilter = cbgpPeerAddrFamilyTable.YFilter
    cbgpPeerAddrFamilyTable.EntityData.YangName = "cbgpPeerAddrFamilyTable"
    cbgpPeerAddrFamilyTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeerAddrFamilyTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeerAddrFamilyTable.EntityData.SegmentPath = "cbgpPeerAddrFamilyTable"
    cbgpPeerAddrFamilyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeerAddrFamilyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeerAddrFamilyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeerAddrFamilyTable.EntityData.Children = types.NewOrderedMap()
    cbgpPeerAddrFamilyTable.EntityData.Children.Append("cbgpPeerAddrFamilyEntry", types.YChild{"CbgpPeerAddrFamilyEntry", nil})
    for i := range cbgpPeerAddrFamilyTable.CbgpPeerAddrFamilyEntry {
        cbgpPeerAddrFamilyTable.EntityData.Children.Append(types.GetSegmentPath(cbgpPeerAddrFamilyTable.CbgpPeerAddrFamilyEntry[i]), types.YChild{"CbgpPeerAddrFamilyEntry", cbgpPeerAddrFamilyTable.CbgpPeerAddrFamilyEntry[i]})
    }
    cbgpPeerAddrFamilyTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeerAddrFamilyTable.EntityData.YListKeys = []string {}

    return &(cbgpPeerAddrFamilyTable.EntityData)
}

// CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains names associated with
// an address family.
type CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to bgp4_mib.BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerRemoteAddr
    BgpPeerRemoteAddr interface{}

    // This attribute is a key. The AFI index of the entry. An implementation is
    // only required to support IPv4 unicast and  VPNv4 (Value - 1) address
    // families. The type is InetAddressType.
    CbgpPeerAddrFamilyAfi interface{}

    // This attribute is a key. The SAFI index of the entry. An implementation is
    // only required to support IPv4 unicast(Value  - 1) and VPNv4( Value - 128)
    // address families. The type is CbgpSafi.
    CbgpPeerAddrFamilySafi interface{}

    // Implementation specific Address Family name. The type is string.
    CbgpPeerAddrFamilyName interface{}
}

func (cbgpPeerAddrFamilyEntry *CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry) GetEntityData() *types.CommonEntityData {
    cbgpPeerAddrFamilyEntry.EntityData.YFilter = cbgpPeerAddrFamilyEntry.YFilter
    cbgpPeerAddrFamilyEntry.EntityData.YangName = "cbgpPeerAddrFamilyEntry"
    cbgpPeerAddrFamilyEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeerAddrFamilyEntry.EntityData.ParentYangName = "cbgpPeerAddrFamilyTable"
    cbgpPeerAddrFamilyEntry.EntityData.SegmentPath = "cbgpPeerAddrFamilyEntry" + types.AddKeyToken(cbgpPeerAddrFamilyEntry.BgpPeerRemoteAddr, "bgpPeerRemoteAddr") + types.AddKeyToken(cbgpPeerAddrFamilyEntry.CbgpPeerAddrFamilyAfi, "cbgpPeerAddrFamilyAfi") + types.AddKeyToken(cbgpPeerAddrFamilyEntry.CbgpPeerAddrFamilySafi, "cbgpPeerAddrFamilySafi")
    cbgpPeerAddrFamilyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeerAddrFamilyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeerAddrFamilyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeerAddrFamilyEntry.EntityData.Children = types.NewOrderedMap()
    cbgpPeerAddrFamilyEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeerAddrFamilyEntry.EntityData.Leafs.Append("bgpPeerRemoteAddr", types.YLeaf{"BgpPeerRemoteAddr", cbgpPeerAddrFamilyEntry.BgpPeerRemoteAddr})
    cbgpPeerAddrFamilyEntry.EntityData.Leafs.Append("cbgpPeerAddrFamilyAfi", types.YLeaf{"CbgpPeerAddrFamilyAfi", cbgpPeerAddrFamilyEntry.CbgpPeerAddrFamilyAfi})
    cbgpPeerAddrFamilyEntry.EntityData.Leafs.Append("cbgpPeerAddrFamilySafi", types.YLeaf{"CbgpPeerAddrFamilySafi", cbgpPeerAddrFamilyEntry.CbgpPeerAddrFamilySafi})
    cbgpPeerAddrFamilyEntry.EntityData.Leafs.Append("cbgpPeerAddrFamilyName", types.YLeaf{"CbgpPeerAddrFamilyName", cbgpPeerAddrFamilyEntry.CbgpPeerAddrFamilyName})

    cbgpPeerAddrFamilyEntry.EntityData.YListKeys = []string {"BgpPeerRemoteAddr", "CbgpPeerAddrFamilyAfi", "CbgpPeerAddrFamilySafi"}

    return &(cbgpPeerAddrFamilyEntry.EntityData)
}

// CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable
// This table contains prefix related information
// related to address families supported by a peer. 
// Supported address families of a peer are known 
// during BGP connection establishment. When a new 
// supported address family is known, this table 
// is updated with a new entry. When an address 
// family is not supported any more, corresponding 
// entry is deleted from the table.
type CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // information associated  with route prefixes belonging to an address family.
    // The type is slice of
    // CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable_CbgpPeerAddrFamilyPrefixEntry.
    CbgpPeerAddrFamilyPrefixEntry []*CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable_CbgpPeerAddrFamilyPrefixEntry
}

func (cbgpPeerAddrFamilyPrefixTable *CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable) GetEntityData() *types.CommonEntityData {
    cbgpPeerAddrFamilyPrefixTable.EntityData.YFilter = cbgpPeerAddrFamilyPrefixTable.YFilter
    cbgpPeerAddrFamilyPrefixTable.EntityData.YangName = "cbgpPeerAddrFamilyPrefixTable"
    cbgpPeerAddrFamilyPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeerAddrFamilyPrefixTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeerAddrFamilyPrefixTable.EntityData.SegmentPath = "cbgpPeerAddrFamilyPrefixTable"
    cbgpPeerAddrFamilyPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeerAddrFamilyPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeerAddrFamilyPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeerAddrFamilyPrefixTable.EntityData.Children = types.NewOrderedMap()
    cbgpPeerAddrFamilyPrefixTable.EntityData.Children.Append("cbgpPeerAddrFamilyPrefixEntry", types.YChild{"CbgpPeerAddrFamilyPrefixEntry", nil})
    for i := range cbgpPeerAddrFamilyPrefixTable.CbgpPeerAddrFamilyPrefixEntry {
        cbgpPeerAddrFamilyPrefixTable.EntityData.Children.Append(types.GetSegmentPath(cbgpPeerAddrFamilyPrefixTable.CbgpPeerAddrFamilyPrefixEntry[i]), types.YChild{"CbgpPeerAddrFamilyPrefixEntry", cbgpPeerAddrFamilyPrefixTable.CbgpPeerAddrFamilyPrefixEntry[i]})
    }
    cbgpPeerAddrFamilyPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeerAddrFamilyPrefixTable.EntityData.YListKeys = []string {}

    return &(cbgpPeerAddrFamilyPrefixTable.EntityData)
}

// CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable_CbgpPeerAddrFamilyPrefixEntry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains information associated 
// with route prefixes belonging to an address family.
type CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable_CbgpPeerAddrFamilyPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to bgp4_mib.BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerRemoteAddr
    BgpPeerRemoteAddr interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry_CbgpPeerAddrFamilyAfi
    CbgpPeerAddrFamilyAfi interface{}

    // This attribute is a key. The type is CbgpSafi. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeerAddrFamilyTable_CbgpPeerAddrFamilyEntry_CbgpPeerAddrFamilySafi
    CbgpPeerAddrFamilySafi interface{}

    // Number of accepted route prefixes on this connection, which belong to an
    // address family. The type is interface{} with range: 0..4294967295.
    CbgpPeerAcceptedPrefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, received on this  connection is denied. It is initialized
    // to zero when  the connection is undergone a hard reset. The type is
    // interface{} with range: 0..4294967295.
    CbgpPeerDeniedPrefixes interface{}

    // Max number of route prefixes accepted for an address family on this
    // connection. The type is interface{} with range: 1..4294967295.
    CbgpPeerPrefixAdminLimit interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which warning message stating the prefix count is crossed the threshold or 
    // corresponding SNMP notification is generated. The type is interface{} with
    // range: 1..100.
    CbgpPeerPrefixThreshold interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which SNMP clear notification is generated if prefix threshold notification
    // is already generated. The type is interface{} with range: 1..100.
    CbgpPeerPrefixClearThreshold interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is advertised on this connection. It is initialized to zero
    // when  the connection is undergone a hard reset. The type is interface{}
    // with range: 0..4294967295.
    CbgpPeerAdvertisedPrefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is suppressed from being sent on this connection. It is 
    // initialized to zero when the connection is undergone a hard reset. The type
    // is interface{} with range: 0..4294967295.
    CbgpPeerSuppressedPrefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, is withdrawn on this connection. It is initialized to zero
    // when the connection is undergone a hard reset. The type is interface{} with
    // range: 0..4294967295.
    CbgpPeerWithdrawnPrefixes interface{}
}

func (cbgpPeerAddrFamilyPrefixEntry *CISCOBGP4MIB_CbgpPeerAddrFamilyPrefixTable_CbgpPeerAddrFamilyPrefixEntry) GetEntityData() *types.CommonEntityData {
    cbgpPeerAddrFamilyPrefixEntry.EntityData.YFilter = cbgpPeerAddrFamilyPrefixEntry.YFilter
    cbgpPeerAddrFamilyPrefixEntry.EntityData.YangName = "cbgpPeerAddrFamilyPrefixEntry"
    cbgpPeerAddrFamilyPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeerAddrFamilyPrefixEntry.EntityData.ParentYangName = "cbgpPeerAddrFamilyPrefixTable"
    cbgpPeerAddrFamilyPrefixEntry.EntityData.SegmentPath = "cbgpPeerAddrFamilyPrefixEntry" + types.AddKeyToken(cbgpPeerAddrFamilyPrefixEntry.BgpPeerRemoteAddr, "bgpPeerRemoteAddr") + types.AddKeyToken(cbgpPeerAddrFamilyPrefixEntry.CbgpPeerAddrFamilyAfi, "cbgpPeerAddrFamilyAfi") + types.AddKeyToken(cbgpPeerAddrFamilyPrefixEntry.CbgpPeerAddrFamilySafi, "cbgpPeerAddrFamilySafi")
    cbgpPeerAddrFamilyPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeerAddrFamilyPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeerAddrFamilyPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeerAddrFamilyPrefixEntry.EntityData.Children = types.NewOrderedMap()
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("bgpPeerRemoteAddr", types.YLeaf{"BgpPeerRemoteAddr", cbgpPeerAddrFamilyPrefixEntry.BgpPeerRemoteAddr})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerAddrFamilyAfi", types.YLeaf{"CbgpPeerAddrFamilyAfi", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerAddrFamilyAfi})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerAddrFamilySafi", types.YLeaf{"CbgpPeerAddrFamilySafi", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerAddrFamilySafi})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerAcceptedPrefixes", types.YLeaf{"CbgpPeerAcceptedPrefixes", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerAcceptedPrefixes})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerDeniedPrefixes", types.YLeaf{"CbgpPeerDeniedPrefixes", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerDeniedPrefixes})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerPrefixAdminLimit", types.YLeaf{"CbgpPeerPrefixAdminLimit", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerPrefixAdminLimit})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerPrefixThreshold", types.YLeaf{"CbgpPeerPrefixThreshold", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerPrefixThreshold})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerPrefixClearThreshold", types.YLeaf{"CbgpPeerPrefixClearThreshold", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerPrefixClearThreshold})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerAdvertisedPrefixes", types.YLeaf{"CbgpPeerAdvertisedPrefixes", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerAdvertisedPrefixes})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerSuppressedPrefixes", types.YLeaf{"CbgpPeerSuppressedPrefixes", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerSuppressedPrefixes})
    cbgpPeerAddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeerWithdrawnPrefixes", types.YLeaf{"CbgpPeerWithdrawnPrefixes", cbgpPeerAddrFamilyPrefixEntry.CbgpPeerWithdrawnPrefixes})

    cbgpPeerAddrFamilyPrefixEntry.EntityData.YListKeys = []string {"BgpPeerRemoteAddr", "CbgpPeerAddrFamilyAfi", "CbgpPeerAddrFamilySafi"}

    return &(cbgpPeerAddrFamilyPrefixEntry.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2Table
// BGP peer table.  This table contains,
// one entry per BGP peer, information about
// the connections with BGP peers.
type CISCOBGP4MIB_CbgpPeer2Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry containing information about the connection with a BGP peer. The type
    // is slice of CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry.
    CbgpPeer2Entry []*CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry
}

func (cbgpPeer2Table *CISCOBGP4MIB_CbgpPeer2Table) GetEntityData() *types.CommonEntityData {
    cbgpPeer2Table.EntityData.YFilter = cbgpPeer2Table.YFilter
    cbgpPeer2Table.EntityData.YangName = "cbgpPeer2Table"
    cbgpPeer2Table.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2Table.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeer2Table.EntityData.SegmentPath = "cbgpPeer2Table"
    cbgpPeer2Table.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2Table.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2Table.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2Table.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2Table.EntityData.Children.Append("cbgpPeer2Entry", types.YChild{"CbgpPeer2Entry", nil})
    for i := range cbgpPeer2Table.CbgpPeer2Entry {
        cbgpPeer2Table.EntityData.Children.Append(types.GetSegmentPath(cbgpPeer2Table.CbgpPeer2Entry[i]), types.YChild{"CbgpPeer2Entry", cbgpPeer2Table.CbgpPeer2Entry[i]})
    }
    cbgpPeer2Table.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeer2Table.EntityData.YListKeys = []string {}

    return &(cbgpPeer2Table.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry
// Entry containing information about the
// connection with a BGP peer.
type CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Represents the type of Peer address stored in
    // cbgpPeer2Entry. The type is InetAddressType.
    CbgpPeer2Type interface{}

    // This attribute is a key. The remote IP address of this entry's BGP peer.
    // The type is string with length: 0..255.
    CbgpPeer2RemoteAddr interface{}

    // The BGP peer connection state. The type is CbgpPeer2State.
    CbgpPeer2State interface{}

    // The desired state of the BGP connection. A transition from 'stop' to
    // 'start' will cause the BGP Manual Start Event to be generated. A transition
    // from 'start' to 'stop' will cause the BGP Manual Stop Event to be
    // generated. This parameter can be used to restart BGP peer connections. 
    // Care should be used in providing write access to this object without
    // adequate authentication. The type is CbgpPeer2AdminStatus.
    CbgpPeer2AdminStatus interface{}

    // The negotiated version of BGP running between the two peers.  This entry
    // MUST be zero (0) unless the cbgpPeer2State is in the openconfirm or the
    // established state.  Note that legal values for this object are between 0
    // and 255. The type is interface{} with range: -2147483648..2147483647.
    CbgpPeer2NegotiatedVersion interface{}

    // The local IP address of this entry's BGP connection. The type is string
    // with length: 0..255.
    CbgpPeer2LocalAddr interface{}

    // The local port for the TCP connection between the BGP peers. The type is
    // interface{} with range: 0..65535.
    CbgpPeer2LocalPort interface{}

    // The local AS number for this session. The type is interface{} with range:
    // 0..4294967295.
    CbgpPeer2LocalAs interface{}

    // The BGP Identifier of this entry's BGP peer. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CbgpPeer2LocalIdentifier interface{}

    // The remote port for the TCP connection between the BGP peers.  Note that
    // the objects cbgpPeer2LocalAddr, cbgpPeer2LocalPort, cbgpPeer2RemoteAddr,
    // and cbgpPeer2RemotePort provide the appropriate reference to the standard
    // MIB TCP connection table. The type is interface{} with range: 0..65535.
    CbgpPeer2RemotePort interface{}

    // The remote autonomous system number received in the BGP OPEN message. The
    // type is interface{} with range: 0..4294967295.
    CbgpPeer2RemoteAs interface{}

    // The BGP Identifier of this entry's BGP peer. This entry MUST be 0.0.0.0
    // unless the cbgpPeer2State is in the openconfirm or the established state.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CbgpPeer2RemoteIdentifier interface{}

    // The number of BGP UPDATE messages received on this connection. The type is
    // interface{} with range: 0..4294967295.
    CbgpPeer2InUpdates interface{}

    // The number of BGP UPDATE messages transmitted on this connection. The type
    // is interface{} with range: 0..4294967295.
    CbgpPeer2OutUpdates interface{}

    // The total number of messages received from the remote peer on this
    // connection. The type is interface{} with range: 0..4294967295.
    CbgpPeer2InTotalMessages interface{}

    // The total number of messages transmitted to the remote peer on this
    // connection. The type is interface{} with range: 0..4294967295.
    CbgpPeer2OutTotalMessages interface{}

    // The last error code and subcode seen by this peer on this connection.  If
    // no error has occurred, this field is zero.  Otherwise, the first byte of
    // this two byte OCTET STRING contains the error code, and the second byte
    // contains the subcode. The type is string with length: 2.
    CbgpPeer2LastError interface{}

    // The total number of times the BGP FSM transitioned into the established
    // state for this peer. The type is interface{} with range: 0..4294967295.
    CbgpPeer2FsmEstablishedTransitions interface{}

    // This timer indicates how long (in seconds) this peer has been in the
    // established state or how long since this peer was last in the established
    // state.  It is set to zero when a new peer is configured or when the router
    // is booted. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CbgpPeer2FsmEstablishedTime interface{}

    // Time interval (in seconds) for the ConnectRetry timer.  The suggested value
    // for this timer is 120 seconds. The type is interface{} with range:
    // 1..65535. Units are seconds.
    CbgpPeer2ConnectRetryInterval interface{}

    // Time interval (in seconds) for the Hold Timer established with the peer. 
    // The value of this object is calculated by this BGP speaker, using the
    // smaller of the values in cbgpPeer2HoldTimeConfigured and the Hold Time
    // received in the OPEN message.  This value must be at least three seconds if
    // it is not zero (0).  If the Hold Timer has not been established with the
    // peer this object MUST have a value of zero (0).  If the
    // cbgpPeer2HoldTimeConfigured object has a value of (0), then this object
    // MUST have a value of (0). The type is interface{} with range: 0..None |
    // 3..65535. Units are seconds.
    CbgpPeer2HoldTime interface{}

    // Time interval (in seconds) for the KeepAlive timer established with the
    // peer.  The value of this object is calculated by this BGP speaker such
    // that, when compared with cbgpPeer2HoldTime, it has the same proportion that
    // cbgpPeer2KeepAliveConfigured has, compared with
    // cbgpPeer2HoldTimeConfigured.  If the KeepAlive timer has not been
    // established with the peer, this object MUST have a value of zero (0).  If
    // the of cbgpPeer2KeepAliveConfigured object has a value of (0), then this
    // object MUST have a value of (0). The type is interface{} with range:
    // 0..21845. Units are seconds.
    CbgpPeer2KeepAlive interface{}

    // Time interval (in seconds) for the Hold Time configured for this BGP
    // speaker with this peer.  This value is placed in an OPEN message sent to
    // this peer by this BGP speaker, and is compared with the Hold Time field in
    // an OPEN message received from the peer when determining the Hold Time
    // (cbgpPeer2HoldTime) with the peer. This value must not be less than three
    // seconds if it is not zero (0).  If it is zero (0), the Hold Time is NOT to
    // be established with the peer.  The suggested value for this timer is 90
    // seconds. The type is interface{} with range: 0..None | 3..65535. Units are
    // seconds.
    CbgpPeer2HoldTimeConfigured interface{}

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
    CbgpPeer2KeepAliveConfigured interface{}

    // Time interval (in seconds) for the MinASOriginationInterval timer. The
    // suggested value for this timer is 15 seconds. The type is interface{} with
    // range: 1..65535. Units are seconds.
    CbgpPeer2MinASOriginationInterval interface{}

    // Time interval (in seconds) for the MinRouteAdvertisementInterval timer. The
    // suggested value for this timer is 30 seconds for EBGP connections and 5
    // seconds for IBGP connections. The type is interface{} with range: 1..65535.
    // Units are seconds.
    CbgpPeer2MinRouteAdvertisementInterval interface{}

    // Elapsed time (in seconds) since the last BGP UPDATE message was received
    // from the peer. Each time cbgpPeer2InUpdates is incremented, the value of
    // this object is set to zero (0). The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    CbgpPeer2InUpdateElapsedTime interface{}

    // Implementation specific error description for bgpPeerLastErrorReceived. The
    // type is string.
    CbgpPeer2LastErrorTxt interface{}

    // The BGP peer connection previous state. The type is CbgpPeer2PrevState.
    CbgpPeer2PrevState interface{}
}

func (cbgpPeer2Entry *CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry) GetEntityData() *types.CommonEntityData {
    cbgpPeer2Entry.EntityData.YFilter = cbgpPeer2Entry.YFilter
    cbgpPeer2Entry.EntityData.YangName = "cbgpPeer2Entry"
    cbgpPeer2Entry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2Entry.EntityData.ParentYangName = "cbgpPeer2Table"
    cbgpPeer2Entry.EntityData.SegmentPath = "cbgpPeer2Entry" + types.AddKeyToken(cbgpPeer2Entry.CbgpPeer2Type, "cbgpPeer2Type") + types.AddKeyToken(cbgpPeer2Entry.CbgpPeer2RemoteAddr, "cbgpPeer2RemoteAddr")
    cbgpPeer2Entry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2Entry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2Entry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2Entry.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2Entry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2Type", types.YLeaf{"CbgpPeer2Type", cbgpPeer2Entry.CbgpPeer2Type})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2RemoteAddr", types.YLeaf{"CbgpPeer2RemoteAddr", cbgpPeer2Entry.CbgpPeer2RemoteAddr})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2State", types.YLeaf{"CbgpPeer2State", cbgpPeer2Entry.CbgpPeer2State})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2AdminStatus", types.YLeaf{"CbgpPeer2AdminStatus", cbgpPeer2Entry.CbgpPeer2AdminStatus})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2NegotiatedVersion", types.YLeaf{"CbgpPeer2NegotiatedVersion", cbgpPeer2Entry.CbgpPeer2NegotiatedVersion})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2LocalAddr", types.YLeaf{"CbgpPeer2LocalAddr", cbgpPeer2Entry.CbgpPeer2LocalAddr})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2LocalPort", types.YLeaf{"CbgpPeer2LocalPort", cbgpPeer2Entry.CbgpPeer2LocalPort})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2LocalAs", types.YLeaf{"CbgpPeer2LocalAs", cbgpPeer2Entry.CbgpPeer2LocalAs})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2LocalIdentifier", types.YLeaf{"CbgpPeer2LocalIdentifier", cbgpPeer2Entry.CbgpPeer2LocalIdentifier})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2RemotePort", types.YLeaf{"CbgpPeer2RemotePort", cbgpPeer2Entry.CbgpPeer2RemotePort})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2RemoteAs", types.YLeaf{"CbgpPeer2RemoteAs", cbgpPeer2Entry.CbgpPeer2RemoteAs})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2RemoteIdentifier", types.YLeaf{"CbgpPeer2RemoteIdentifier", cbgpPeer2Entry.CbgpPeer2RemoteIdentifier})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2InUpdates", types.YLeaf{"CbgpPeer2InUpdates", cbgpPeer2Entry.CbgpPeer2InUpdates})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2OutUpdates", types.YLeaf{"CbgpPeer2OutUpdates", cbgpPeer2Entry.CbgpPeer2OutUpdates})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2InTotalMessages", types.YLeaf{"CbgpPeer2InTotalMessages", cbgpPeer2Entry.CbgpPeer2InTotalMessages})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2OutTotalMessages", types.YLeaf{"CbgpPeer2OutTotalMessages", cbgpPeer2Entry.CbgpPeer2OutTotalMessages})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2LastError", types.YLeaf{"CbgpPeer2LastError", cbgpPeer2Entry.CbgpPeer2LastError})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2FsmEstablishedTransitions", types.YLeaf{"CbgpPeer2FsmEstablishedTransitions", cbgpPeer2Entry.CbgpPeer2FsmEstablishedTransitions})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2FsmEstablishedTime", types.YLeaf{"CbgpPeer2FsmEstablishedTime", cbgpPeer2Entry.CbgpPeer2FsmEstablishedTime})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2ConnectRetryInterval", types.YLeaf{"CbgpPeer2ConnectRetryInterval", cbgpPeer2Entry.CbgpPeer2ConnectRetryInterval})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2HoldTime", types.YLeaf{"CbgpPeer2HoldTime", cbgpPeer2Entry.CbgpPeer2HoldTime})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2KeepAlive", types.YLeaf{"CbgpPeer2KeepAlive", cbgpPeer2Entry.CbgpPeer2KeepAlive})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2HoldTimeConfigured", types.YLeaf{"CbgpPeer2HoldTimeConfigured", cbgpPeer2Entry.CbgpPeer2HoldTimeConfigured})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2KeepAliveConfigured", types.YLeaf{"CbgpPeer2KeepAliveConfigured", cbgpPeer2Entry.CbgpPeer2KeepAliveConfigured})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2MinASOriginationInterval", types.YLeaf{"CbgpPeer2MinASOriginationInterval", cbgpPeer2Entry.CbgpPeer2MinASOriginationInterval})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2MinRouteAdvertisementInterval", types.YLeaf{"CbgpPeer2MinRouteAdvertisementInterval", cbgpPeer2Entry.CbgpPeer2MinRouteAdvertisementInterval})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2InUpdateElapsedTime", types.YLeaf{"CbgpPeer2InUpdateElapsedTime", cbgpPeer2Entry.CbgpPeer2InUpdateElapsedTime})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2LastErrorTxt", types.YLeaf{"CbgpPeer2LastErrorTxt", cbgpPeer2Entry.CbgpPeer2LastErrorTxt})
    cbgpPeer2Entry.EntityData.Leafs.Append("cbgpPeer2PrevState", types.YLeaf{"CbgpPeer2PrevState", cbgpPeer2Entry.CbgpPeer2PrevState})

    cbgpPeer2Entry.EntityData.YListKeys = []string {"CbgpPeer2Type", "CbgpPeer2RemoteAddr"}

    return &(cbgpPeer2Entry.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2AdminStatus represents authentication.
type CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2AdminStatus string

const (
    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2AdminStatus_stop CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2AdminStatus = "stop"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2AdminStatus_start CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2AdminStatus = "start"
)

// CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState represents The BGP peer connection previous state.
type CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState string

const (
    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_none CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "none"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_idle CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "idle"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_connect CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "connect"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_active CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "active"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_opensent CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "opensent"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_openconfirm CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "openconfirm"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState_established CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2PrevState = "established"
)

// CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State represents The BGP peer connection state.
type CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State string

const (
    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State_idle CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State = "idle"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State_connect CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State = "connect"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State_active CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State = "active"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State_opensent CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State = "opensent"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State_openconfirm CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State = "openconfirm"

    CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State_established CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2State = "established"
)

// CISCOBGP4MIB_CbgpPeer2CapsTable
// This table contains the capabilities that are
// supported by a peer. Capabilities of a peer are
// received during BGP connection establishment.
// Values corresponding to each received capability
// are stored in this table. When a new capability
// is received, this table is updated with a new
// entry. When an existing capability is not received
// during the latest connection establishment, the
// corresponding entry is deleted from the table.
type CISCOBGP4MIB_CbgpPeer2CapsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a capability received from a peer with a particular
    // code and an index. When a capability is received multiple times with
    // different values during a BGP connection establishment, corresponding
    // entries are differentiated with indices. The type is slice of
    // CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry.
    CbgpPeer2CapsEntry []*CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry
}

func (cbgpPeer2CapsTable *CISCOBGP4MIB_CbgpPeer2CapsTable) GetEntityData() *types.CommonEntityData {
    cbgpPeer2CapsTable.EntityData.YFilter = cbgpPeer2CapsTable.YFilter
    cbgpPeer2CapsTable.EntityData.YangName = "cbgpPeer2CapsTable"
    cbgpPeer2CapsTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2CapsTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeer2CapsTable.EntityData.SegmentPath = "cbgpPeer2CapsTable"
    cbgpPeer2CapsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2CapsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2CapsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2CapsTable.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2CapsTable.EntityData.Children.Append("cbgpPeer2CapsEntry", types.YChild{"CbgpPeer2CapsEntry", nil})
    for i := range cbgpPeer2CapsTable.CbgpPeer2CapsEntry {
        cbgpPeer2CapsTable.EntityData.Children.Append(types.GetSegmentPath(cbgpPeer2CapsTable.CbgpPeer2CapsEntry[i]), types.YChild{"CbgpPeer2CapsEntry", cbgpPeer2CapsTable.CbgpPeer2CapsEntry[i]})
    }
    cbgpPeer2CapsTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeer2CapsTable.EntityData.YListKeys = []string {}

    return &(cbgpPeer2CapsTable.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry
// Each entry represents a capability received from a
// peer with a particular code and an index. When a
// capability is received multiple times with different
// values during a BGP connection establishment,
// corresponding entries are differentiated with indices.
type CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2Type
    CbgpPeer2Type interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2RemoteAddr
    CbgpPeer2RemoteAddr interface{}

    // This attribute is a key. The BGP Capability Advertisement Capability Code.
    // The type is CbgpPeer2CapCode.
    CbgpPeer2CapCode interface{}

    // This attribute is a key. Multiple instances of a given capability may be
    // sent by a BGP speaker.  This variable is used to index them. The type is
    // interface{} with range: 1..128.
    CbgpPeer2CapIndex interface{}

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
    CbgpPeer2CapValue interface{}
}

func (cbgpPeer2CapsEntry *CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry) GetEntityData() *types.CommonEntityData {
    cbgpPeer2CapsEntry.EntityData.YFilter = cbgpPeer2CapsEntry.YFilter
    cbgpPeer2CapsEntry.EntityData.YangName = "cbgpPeer2CapsEntry"
    cbgpPeer2CapsEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2CapsEntry.EntityData.ParentYangName = "cbgpPeer2CapsTable"
    cbgpPeer2CapsEntry.EntityData.SegmentPath = "cbgpPeer2CapsEntry" + types.AddKeyToken(cbgpPeer2CapsEntry.CbgpPeer2Type, "cbgpPeer2Type") + types.AddKeyToken(cbgpPeer2CapsEntry.CbgpPeer2RemoteAddr, "cbgpPeer2RemoteAddr") + types.AddKeyToken(cbgpPeer2CapsEntry.CbgpPeer2CapCode, "cbgpPeer2CapCode") + types.AddKeyToken(cbgpPeer2CapsEntry.CbgpPeer2CapIndex, "cbgpPeer2CapIndex")
    cbgpPeer2CapsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2CapsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2CapsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2CapsEntry.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2CapsEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeer2CapsEntry.EntityData.Leafs.Append("cbgpPeer2Type", types.YLeaf{"CbgpPeer2Type", cbgpPeer2CapsEntry.CbgpPeer2Type})
    cbgpPeer2CapsEntry.EntityData.Leafs.Append("cbgpPeer2RemoteAddr", types.YLeaf{"CbgpPeer2RemoteAddr", cbgpPeer2CapsEntry.CbgpPeer2RemoteAddr})
    cbgpPeer2CapsEntry.EntityData.Leafs.Append("cbgpPeer2CapCode", types.YLeaf{"CbgpPeer2CapCode", cbgpPeer2CapsEntry.CbgpPeer2CapCode})
    cbgpPeer2CapsEntry.EntityData.Leafs.Append("cbgpPeer2CapIndex", types.YLeaf{"CbgpPeer2CapIndex", cbgpPeer2CapsEntry.CbgpPeer2CapIndex})
    cbgpPeer2CapsEntry.EntityData.Leafs.Append("cbgpPeer2CapValue", types.YLeaf{"CbgpPeer2CapValue", cbgpPeer2CapsEntry.CbgpPeer2CapValue})

    cbgpPeer2CapsEntry.EntityData.YListKeys = []string {"CbgpPeer2Type", "CbgpPeer2RemoteAddr", "CbgpPeer2CapCode", "CbgpPeer2CapIndex"}

    return &(cbgpPeer2CapsEntry.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode represents The BGP Capability Advertisement Capability Code.
type CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode string

const (
    CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode_multiProtocol CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode = "multiProtocol"

    CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode_routeRefresh CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode = "routeRefresh"

    CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode_gracefulRestart CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode = "gracefulRestart"

    CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode_fourByteAs CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode = "fourByteAs"

    CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode_addPath CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode = "addPath"

    CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode_routeRefreshOld CISCOBGP4MIB_CbgpPeer2CapsTable_CbgpPeer2CapsEntry_CbgpPeer2CapCode = "routeRefreshOld"
)

// CISCOBGP4MIB_CbgpPeer2AddrFamilyTable
// This table contains information related to
// address families supported by a peer. Supported
// address families of a peer are known during BGP
// connection establishment. When a new supported
// address family is known, this table is updated
// with a new entry. When an address family is not
// supported any more, corresponding entry is deleted
// from the table.
type CISCOBGP4MIB_CbgpPeer2AddrFamilyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // names associated with an address family. The type is slice of
    // CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry.
    CbgpPeer2AddrFamilyEntry []*CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry
}

func (cbgpPeer2AddrFamilyTable *CISCOBGP4MIB_CbgpPeer2AddrFamilyTable) GetEntityData() *types.CommonEntityData {
    cbgpPeer2AddrFamilyTable.EntityData.YFilter = cbgpPeer2AddrFamilyTable.YFilter
    cbgpPeer2AddrFamilyTable.EntityData.YangName = "cbgpPeer2AddrFamilyTable"
    cbgpPeer2AddrFamilyTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2AddrFamilyTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeer2AddrFamilyTable.EntityData.SegmentPath = "cbgpPeer2AddrFamilyTable"
    cbgpPeer2AddrFamilyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2AddrFamilyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2AddrFamilyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2AddrFamilyTable.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2AddrFamilyTable.EntityData.Children.Append("cbgpPeer2AddrFamilyEntry", types.YChild{"CbgpPeer2AddrFamilyEntry", nil})
    for i := range cbgpPeer2AddrFamilyTable.CbgpPeer2AddrFamilyEntry {
        cbgpPeer2AddrFamilyTable.EntityData.Children.Append(types.GetSegmentPath(cbgpPeer2AddrFamilyTable.CbgpPeer2AddrFamilyEntry[i]), types.YChild{"CbgpPeer2AddrFamilyEntry", cbgpPeer2AddrFamilyTable.CbgpPeer2AddrFamilyEntry[i]})
    }
    cbgpPeer2AddrFamilyTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeer2AddrFamilyTable.EntityData.YListKeys = []string {}

    return &(cbgpPeer2AddrFamilyTable.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains names associated with
// an address family.
type CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2Type
    CbgpPeer2Type interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2RemoteAddr
    CbgpPeer2RemoteAddr interface{}

    // This attribute is a key. The AFI index of the entry. An implementation is
    // only required to support IPv4 unicast and VPNv4 (Value - 1) address
    // families. The type is InetAddressType.
    CbgpPeer2AddrFamilyAfi interface{}

    // This attribute is a key. The SAFI index of the entry. An implementation is
    // only required to support IPv4 unicast(Value - 1) and VPNv4( Value - 128)
    // address families. The type is CbgpSafi.
    CbgpPeer2AddrFamilySafi interface{}

    // Implementation specific Address Family name. The type is string.
    CbgpPeer2AddrFamilyName interface{}
}

func (cbgpPeer2AddrFamilyEntry *CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry) GetEntityData() *types.CommonEntityData {
    cbgpPeer2AddrFamilyEntry.EntityData.YFilter = cbgpPeer2AddrFamilyEntry.YFilter
    cbgpPeer2AddrFamilyEntry.EntityData.YangName = "cbgpPeer2AddrFamilyEntry"
    cbgpPeer2AddrFamilyEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2AddrFamilyEntry.EntityData.ParentYangName = "cbgpPeer2AddrFamilyTable"
    cbgpPeer2AddrFamilyEntry.EntityData.SegmentPath = "cbgpPeer2AddrFamilyEntry" + types.AddKeyToken(cbgpPeer2AddrFamilyEntry.CbgpPeer2Type, "cbgpPeer2Type") + types.AddKeyToken(cbgpPeer2AddrFamilyEntry.CbgpPeer2RemoteAddr, "cbgpPeer2RemoteAddr") + types.AddKeyToken(cbgpPeer2AddrFamilyEntry.CbgpPeer2AddrFamilyAfi, "cbgpPeer2AddrFamilyAfi") + types.AddKeyToken(cbgpPeer2AddrFamilyEntry.CbgpPeer2AddrFamilySafi, "cbgpPeer2AddrFamilySafi")
    cbgpPeer2AddrFamilyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2AddrFamilyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2AddrFamilyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2AddrFamilyEntry.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2AddrFamilyEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeer2AddrFamilyEntry.EntityData.Leafs.Append("cbgpPeer2Type", types.YLeaf{"CbgpPeer2Type", cbgpPeer2AddrFamilyEntry.CbgpPeer2Type})
    cbgpPeer2AddrFamilyEntry.EntityData.Leafs.Append("cbgpPeer2RemoteAddr", types.YLeaf{"CbgpPeer2RemoteAddr", cbgpPeer2AddrFamilyEntry.CbgpPeer2RemoteAddr})
    cbgpPeer2AddrFamilyEntry.EntityData.Leafs.Append("cbgpPeer2AddrFamilyAfi", types.YLeaf{"CbgpPeer2AddrFamilyAfi", cbgpPeer2AddrFamilyEntry.CbgpPeer2AddrFamilyAfi})
    cbgpPeer2AddrFamilyEntry.EntityData.Leafs.Append("cbgpPeer2AddrFamilySafi", types.YLeaf{"CbgpPeer2AddrFamilySafi", cbgpPeer2AddrFamilyEntry.CbgpPeer2AddrFamilySafi})
    cbgpPeer2AddrFamilyEntry.EntityData.Leafs.Append("cbgpPeer2AddrFamilyName", types.YLeaf{"CbgpPeer2AddrFamilyName", cbgpPeer2AddrFamilyEntry.CbgpPeer2AddrFamilyName})

    cbgpPeer2AddrFamilyEntry.EntityData.YListKeys = []string {"CbgpPeer2Type", "CbgpPeer2RemoteAddr", "CbgpPeer2AddrFamilyAfi", "CbgpPeer2AddrFamilySafi"}

    return &(cbgpPeer2AddrFamilyEntry.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable
// This table contains prefix related information
// related to address families supported by a peer.
// Supported address families of a peer are known
// during BGP connection establishment. When a new
// supported address family is known, this table
// is updated with a new entry. When an address
// family is not supported any more, corresponding
// entry is deleted from the table.
type CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // information associated with route prefixes belonging to an address family.
    // The type is slice of
    // CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable_CbgpPeer2AddrFamilyPrefixEntry.
    CbgpPeer2AddrFamilyPrefixEntry []*CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable_CbgpPeer2AddrFamilyPrefixEntry
}

func (cbgpPeer2AddrFamilyPrefixTable *CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable) GetEntityData() *types.CommonEntityData {
    cbgpPeer2AddrFamilyPrefixTable.EntityData.YFilter = cbgpPeer2AddrFamilyPrefixTable.YFilter
    cbgpPeer2AddrFamilyPrefixTable.EntityData.YangName = "cbgpPeer2AddrFamilyPrefixTable"
    cbgpPeer2AddrFamilyPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2AddrFamilyPrefixTable.EntityData.ParentYangName = "CISCO-BGP4-MIB"
    cbgpPeer2AddrFamilyPrefixTable.EntityData.SegmentPath = "cbgpPeer2AddrFamilyPrefixTable"
    cbgpPeer2AddrFamilyPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2AddrFamilyPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2AddrFamilyPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2AddrFamilyPrefixTable.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2AddrFamilyPrefixTable.EntityData.Children.Append("cbgpPeer2AddrFamilyPrefixEntry", types.YChild{"CbgpPeer2AddrFamilyPrefixEntry", nil})
    for i := range cbgpPeer2AddrFamilyPrefixTable.CbgpPeer2AddrFamilyPrefixEntry {
        cbgpPeer2AddrFamilyPrefixTable.EntityData.Children.Append(types.GetSegmentPath(cbgpPeer2AddrFamilyPrefixTable.CbgpPeer2AddrFamilyPrefixEntry[i]), types.YChild{"CbgpPeer2AddrFamilyPrefixEntry", cbgpPeer2AddrFamilyPrefixTable.CbgpPeer2AddrFamilyPrefixEntry[i]})
    }
    cbgpPeer2AddrFamilyPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    cbgpPeer2AddrFamilyPrefixTable.EntityData.YListKeys = []string {}

    return &(cbgpPeer2AddrFamilyPrefixTable.EntityData)
}

// CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable_CbgpPeer2AddrFamilyPrefixEntry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains information associated
// with route prefixes belonging to an address family.
type CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable_CbgpPeer2AddrFamilyPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2Type
    CbgpPeer2Type interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2Table_CbgpPeer2Entry_CbgpPeer2RemoteAddr
    CbgpPeer2RemoteAddr interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry_CbgpPeer2AddrFamilyAfi
    CbgpPeer2AddrFamilyAfi interface{}

    // This attribute is a key. The type is CbgpSafi. Refers to
    // cisco_bgp4_mib.CISCOBGP4MIB_CbgpPeer2AddrFamilyTable_CbgpPeer2AddrFamilyEntry_CbgpPeer2AddrFamilySafi
    CbgpPeer2AddrFamilySafi interface{}

    // Number of accepted route prefixes on this connection, which belong to an
    // address family. The type is interface{} with range: 0..4294967295.
    CbgpPeer2AcceptedPrefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, received on this connection is denied. It is initialized to
    // zero when the connection is undergone a hard reset. The type is interface{}
    // with range: 0..4294967295.
    CbgpPeer2DeniedPrefixes interface{}

    // Max number of route prefixes accepted for an address family on this
    // connection. The type is interface{} with range: 1..4294967295.
    CbgpPeer2PrefixAdminLimit interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which warning message stating the prefix count is crossed the threshold or
    // corresponding SNMP notification is generated. The type is interface{} with
    // range: 1..100. Units are percent.
    CbgpPeer2PrefixThreshold interface{}

    // Prefix threshold value (%) for an address family on this connection at
    // which SNMP clear notification is generated if prefix threshold notification
    // is already generated. The type is interface{} with range: 1..100. Units are
    // percent.
    CbgpPeer2PrefixClearThreshold interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is advertised on this connection. It is initialized to zero
    // when the connection is undergone a hard reset. The type is interface{} with
    // range: 0..4294967295.
    CbgpPeer2AdvertisedPrefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family is suppressed from being sent on this connection. It is
    // initialized to zero when the connection is undergone a hard reset. The type
    // is interface{} with range: 0..4294967295.
    CbgpPeer2SuppressedPrefixes interface{}

    // This counter is incremented when a route prefix, which belongs to an
    // address family, is withdrawn on this connection. It is initialized to zero
    // when the connection is undergone a hard reset. The type is interface{} with
    // range: 0..4294967295.
    CbgpPeer2WithdrawnPrefixes interface{}
}

func (cbgpPeer2AddrFamilyPrefixEntry *CISCOBGP4MIB_CbgpPeer2AddrFamilyPrefixTable_CbgpPeer2AddrFamilyPrefixEntry) GetEntityData() *types.CommonEntityData {
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.YFilter = cbgpPeer2AddrFamilyPrefixEntry.YFilter
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.YangName = "cbgpPeer2AddrFamilyPrefixEntry"
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.ParentYangName = "cbgpPeer2AddrFamilyPrefixTable"
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.SegmentPath = "cbgpPeer2AddrFamilyPrefixEntry" + types.AddKeyToken(cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2Type, "cbgpPeer2Type") + types.AddKeyToken(cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2RemoteAddr, "cbgpPeer2RemoteAddr") + types.AddKeyToken(cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2AddrFamilyAfi, "cbgpPeer2AddrFamilyAfi") + types.AddKeyToken(cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2AddrFamilySafi, "cbgpPeer2AddrFamilySafi")
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Children = types.NewOrderedMap()
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2Type", types.YLeaf{"CbgpPeer2Type", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2Type})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2RemoteAddr", types.YLeaf{"CbgpPeer2RemoteAddr", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2RemoteAddr})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2AddrFamilyAfi", types.YLeaf{"CbgpPeer2AddrFamilyAfi", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2AddrFamilyAfi})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2AddrFamilySafi", types.YLeaf{"CbgpPeer2AddrFamilySafi", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2AddrFamilySafi})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2AcceptedPrefixes", types.YLeaf{"CbgpPeer2AcceptedPrefixes", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2AcceptedPrefixes})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2DeniedPrefixes", types.YLeaf{"CbgpPeer2DeniedPrefixes", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2DeniedPrefixes})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2PrefixAdminLimit", types.YLeaf{"CbgpPeer2PrefixAdminLimit", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2PrefixAdminLimit})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2PrefixThreshold", types.YLeaf{"CbgpPeer2PrefixThreshold", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2PrefixThreshold})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2PrefixClearThreshold", types.YLeaf{"CbgpPeer2PrefixClearThreshold", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2PrefixClearThreshold})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2AdvertisedPrefixes", types.YLeaf{"CbgpPeer2AdvertisedPrefixes", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2AdvertisedPrefixes})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2SuppressedPrefixes", types.YLeaf{"CbgpPeer2SuppressedPrefixes", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2SuppressedPrefixes})
    cbgpPeer2AddrFamilyPrefixEntry.EntityData.Leafs.Append("cbgpPeer2WithdrawnPrefixes", types.YLeaf{"CbgpPeer2WithdrawnPrefixes", cbgpPeer2AddrFamilyPrefixEntry.CbgpPeer2WithdrawnPrefixes})

    cbgpPeer2AddrFamilyPrefixEntry.EntityData.YListKeys = []string {"CbgpPeer2Type", "CbgpPeer2RemoteAddr", "CbgpPeer2AddrFamilyAfi", "CbgpPeer2AddrFamilySafi"}

    return &(cbgpPeer2AddrFamilyPrefixEntry.EntityData)
}

