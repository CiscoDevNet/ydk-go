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
    parent types.Entity
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

func (cISCOBGP4MIB *CISCOBGP4MIB) GetFilter() yfilter.YFilter { return cISCOBGP4MIB.YFilter }

func (cISCOBGP4MIB *CISCOBGP4MIB) SetFilter(yf yfilter.YFilter) { cISCOBGP4MIB.YFilter = yf }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetGoName(yname string) string {
    if yname == "cbgpGlobal" { return "Cbgpglobal" }
    if yname == "cbgpRouteTable" { return "Cbgproutetable" }
    if yname == "cbgpPeerCapsTable" { return "Cbgppeercapstable" }
    if yname == "cbgpPeerAddrFamilyTable" { return "Cbgppeeraddrfamilytable" }
    if yname == "cbgpPeerAddrFamilyPrefixTable" { return "Cbgppeeraddrfamilyprefixtable" }
    if yname == "cbgpPeer2Table" { return "Cbgppeer2Table" }
    if yname == "cbgpPeer2CapsTable" { return "Cbgppeer2Capstable" }
    if yname == "cbgpPeer2AddrFamilyTable" { return "Cbgppeer2Addrfamilytable" }
    if yname == "cbgpPeer2AddrFamilyPrefixTable" { return "Cbgppeer2Addrfamilyprefixtable" }
    return ""
}

func (cISCOBGP4MIB *CISCOBGP4MIB) GetSegmentPath() string {
    return "CISCO-BGP4-MIB:CISCO-BGP4-MIB"
}

func (cISCOBGP4MIB *CISCOBGP4MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpGlobal" {
        return &cISCOBGP4MIB.Cbgpglobal
    }
    if childYangName == "cbgpRouteTable" {
        return &cISCOBGP4MIB.Cbgproutetable
    }
    if childYangName == "cbgpPeerCapsTable" {
        return &cISCOBGP4MIB.Cbgppeercapstable
    }
    if childYangName == "cbgpPeerAddrFamilyTable" {
        return &cISCOBGP4MIB.Cbgppeeraddrfamilytable
    }
    if childYangName == "cbgpPeerAddrFamilyPrefixTable" {
        return &cISCOBGP4MIB.Cbgppeeraddrfamilyprefixtable
    }
    if childYangName == "cbgpPeer2Table" {
        return &cISCOBGP4MIB.Cbgppeer2Table
    }
    if childYangName == "cbgpPeer2CapsTable" {
        return &cISCOBGP4MIB.Cbgppeer2Capstable
    }
    if childYangName == "cbgpPeer2AddrFamilyTable" {
        return &cISCOBGP4MIB.Cbgppeer2Addrfamilytable
    }
    if childYangName == "cbgpPeer2AddrFamilyPrefixTable" {
        return &cISCOBGP4MIB.Cbgppeer2Addrfamilyprefixtable
    }
    return nil
}

func (cISCOBGP4MIB *CISCOBGP4MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cbgpGlobal"] = &cISCOBGP4MIB.Cbgpglobal
    children["cbgpRouteTable"] = &cISCOBGP4MIB.Cbgproutetable
    children["cbgpPeerCapsTable"] = &cISCOBGP4MIB.Cbgppeercapstable
    children["cbgpPeerAddrFamilyTable"] = &cISCOBGP4MIB.Cbgppeeraddrfamilytable
    children["cbgpPeerAddrFamilyPrefixTable"] = &cISCOBGP4MIB.Cbgppeeraddrfamilyprefixtable
    children["cbgpPeer2Table"] = &cISCOBGP4MIB.Cbgppeer2Table
    children["cbgpPeer2CapsTable"] = &cISCOBGP4MIB.Cbgppeer2Capstable
    children["cbgpPeer2AddrFamilyTable"] = &cISCOBGP4MIB.Cbgppeer2Addrfamilytable
    children["cbgpPeer2AddrFamilyPrefixTable"] = &cISCOBGP4MIB.Cbgppeer2Addrfamilyprefixtable
    return children
}

func (cISCOBGP4MIB *CISCOBGP4MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOBGP4MIB *CISCOBGP4MIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetYangName() string { return "CISCO-BGP4-MIB" }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOBGP4MIB *CISCOBGP4MIB) SetParent(parent types.Entity) { cISCOBGP4MIB.parent = parent }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetParent() types.Entity { return cISCOBGP4MIB.parent }

func (cISCOBGP4MIB *CISCOBGP4MIB) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgpglobal
type CISCOBGP4MIB_Cbgpglobal struct {
    parent types.Entity
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

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetFilter() yfilter.YFilter { return cbgpglobal.YFilter }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) SetFilter(yf yfilter.YFilter) { cbgpglobal.YFilter = yf }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetGoName(yname string) string {
    if yname == "cbgpNotifsEnable" { return "Cbgpnotifsenable" }
    if yname == "cbgpLocalAs" { return "Cbgplocalas" }
    return ""
}

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetSegmentPath() string {
    return "cbgpGlobal"
}

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbgpNotifsEnable"] = cbgpglobal.Cbgpnotifsenable
    leafs["cbgpLocalAs"] = cbgpglobal.Cbgplocalas
    return leafs
}

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetBundleName() string { return "cisco_ios_xe" }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetYangName() string { return "cbgpGlobal" }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) SetParent(parent types.Entity) { cbgpglobal.parent = parent }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetParent() types.Entity { return cbgpglobal.parent }

func (cbgpglobal *CISCOBGP4MIB_Cbgpglobal) GetParentYangName() string { return "CISCO-BGP4-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a path to a network received from a peer. The type is
    // slice of CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry.
    Cbgprouteentry []CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetFilter() yfilter.YFilter { return cbgproutetable.YFilter }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) SetFilter(yf yfilter.YFilter) { cbgproutetable.YFilter = yf }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetGoName(yname string) string {
    if yname == "cbgpRouteEntry" { return "Cbgprouteentry" }
    return ""
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetSegmentPath() string {
    return "cbgpRouteTable"
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpRouteEntry" {
        for _, c := range cbgproutetable.Cbgprouteentry {
            if cbgproutetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry{}
        cbgproutetable.Cbgprouteentry = append(cbgproutetable.Cbgprouteentry, child)
        return &cbgproutetable.Cbgprouteentry[len(cbgproutetable.Cbgprouteentry)-1]
    }
    return nil
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgproutetable.Cbgprouteentry {
        children[cbgproutetable.Cbgprouteentry[i].GetSegmentPath()] = &cbgproutetable.Cbgprouteentry[i]
    }
    return children
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetYangName() string { return "cbgpRouteTable" }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) SetParent(parent types.Entity) { cbgproutetable.parent = parent }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetParent() types.Entity { return cbgproutetable.parent }

func (cbgproutetable *CISCOBGP4MIB_Cbgproutetable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry
// Information about a path to a network received from
// a peer.
type CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry struct {
    parent types.Entity
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

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetFilter() yfilter.YFilter { return cbgprouteentry.YFilter }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) SetFilter(yf yfilter.YFilter) { cbgprouteentry.YFilter = yf }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetGoName(yname string) string {
    if yname == "cbgpRouteAfi" { return "Cbgprouteafi" }
    if yname == "cbgpRouteSafi" { return "Cbgproutesafi" }
    if yname == "cbgpRoutePeerType" { return "Cbgproutepeertype" }
    if yname == "cbgpRoutePeer" { return "Cbgproutepeer" }
    if yname == "cbgpRouteAddrPrefix" { return "Cbgprouteaddrprefix" }
    if yname == "cbgpRouteAddrPrefixLen" { return "Cbgprouteaddrprefixlen" }
    if yname == "cbgpRouteOrigin" { return "Cbgprouteorigin" }
    if yname == "cbgpRouteASPathSegment" { return "Cbgprouteaspathsegment" }
    if yname == "cbgpRouteNextHop" { return "Cbgproutenexthop" }
    if yname == "cbgpRouteMedPresent" { return "Cbgproutemedpresent" }
    if yname == "cbgpRouteMultiExitDisc" { return "Cbgproutemultiexitdisc" }
    if yname == "cbgpRouteLocalPrefPresent" { return "Cbgproutelocalprefpresent" }
    if yname == "cbgpRouteLocalPref" { return "Cbgproutelocalpref" }
    if yname == "cbgpRouteAtomicAggregate" { return "Cbgprouteatomicaggregate" }
    if yname == "cbgpRouteAggregatorAS" { return "Cbgprouteaggregatoras" }
    if yname == "cbgpRouteAggregatorAddrType" { return "Cbgprouteaggregatoraddrtype" }
    if yname == "cbgpRouteAggregatorAddr" { return "Cbgprouteaggregatoraddr" }
    if yname == "cbgpRouteBest" { return "Cbgproutebest" }
    if yname == "cbgpRouteUnknownAttr" { return "Cbgprouteunknownattr" }
    return ""
}

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetSegmentPath() string {
    return "cbgpRouteEntry" + "[cbgpRouteAfi='" + fmt.Sprintf("%v", cbgprouteentry.Cbgprouteafi) + "']" + "[cbgpRouteSafi='" + fmt.Sprintf("%v", cbgprouteentry.Cbgproutesafi) + "']" + "[cbgpRoutePeerType='" + fmt.Sprintf("%v", cbgprouteentry.Cbgproutepeertype) + "']" + "[cbgpRoutePeer='" + fmt.Sprintf("%v", cbgprouteentry.Cbgproutepeer) + "']" + "[cbgpRouteAddrPrefix='" + fmt.Sprintf("%v", cbgprouteentry.Cbgprouteaddrprefix) + "']" + "[cbgpRouteAddrPrefixLen='" + fmt.Sprintf("%v", cbgprouteentry.Cbgprouteaddrprefixlen) + "']"
}

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbgpRouteAfi"] = cbgprouteentry.Cbgprouteafi
    leafs["cbgpRouteSafi"] = cbgprouteentry.Cbgproutesafi
    leafs["cbgpRoutePeerType"] = cbgprouteentry.Cbgproutepeertype
    leafs["cbgpRoutePeer"] = cbgprouteentry.Cbgproutepeer
    leafs["cbgpRouteAddrPrefix"] = cbgprouteentry.Cbgprouteaddrprefix
    leafs["cbgpRouteAddrPrefixLen"] = cbgprouteentry.Cbgprouteaddrprefixlen
    leafs["cbgpRouteOrigin"] = cbgprouteentry.Cbgprouteorigin
    leafs["cbgpRouteASPathSegment"] = cbgprouteentry.Cbgprouteaspathsegment
    leafs["cbgpRouteNextHop"] = cbgprouteentry.Cbgproutenexthop
    leafs["cbgpRouteMedPresent"] = cbgprouteentry.Cbgproutemedpresent
    leafs["cbgpRouteMultiExitDisc"] = cbgprouteentry.Cbgproutemultiexitdisc
    leafs["cbgpRouteLocalPrefPresent"] = cbgprouteentry.Cbgproutelocalprefpresent
    leafs["cbgpRouteLocalPref"] = cbgprouteentry.Cbgproutelocalpref
    leafs["cbgpRouteAtomicAggregate"] = cbgprouteentry.Cbgprouteatomicaggregate
    leafs["cbgpRouteAggregatorAS"] = cbgprouteentry.Cbgprouteaggregatoras
    leafs["cbgpRouteAggregatorAddrType"] = cbgprouteentry.Cbgprouteaggregatoraddrtype
    leafs["cbgpRouteAggregatorAddr"] = cbgprouteentry.Cbgprouteaggregatoraddr
    leafs["cbgpRouteBest"] = cbgprouteentry.Cbgproutebest
    leafs["cbgpRouteUnknownAttr"] = cbgprouteentry.Cbgprouteunknownattr
    return leafs
}

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetYangName() string { return "cbgpRouteEntry" }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) SetParent(parent types.Entity) { cbgprouteentry.parent = parent }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetParent() types.Entity { return cbgprouteentry.parent }

func (cbgprouteentry *CISCOBGP4MIB_Cbgproutetable_Cbgprouteentry) GetParentYangName() string { return "cbgpRouteTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a capability received from a peer with a particular
    // code and an index. When a  capability is received multiple times with
    // different values during a BGP connection establishment,  corresponding
    // entries are differentiated with indices. The type is slice of
    // CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry.
    Cbgppeercapsentry []CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetFilter() yfilter.YFilter { return cbgppeercapstable.YFilter }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) SetFilter(yf yfilter.YFilter) { cbgppeercapstable.YFilter = yf }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetGoName(yname string) string {
    if yname == "cbgpPeerCapsEntry" { return "Cbgppeercapsentry" }
    return ""
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetSegmentPath() string {
    return "cbgpPeerCapsTable"
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeerCapsEntry" {
        for _, c := range cbgppeercapstable.Cbgppeercapsentry {
            if cbgppeercapstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry{}
        cbgppeercapstable.Cbgppeercapsentry = append(cbgppeercapstable.Cbgppeercapsentry, child)
        return &cbgppeercapstable.Cbgppeercapsentry[len(cbgppeercapstable.Cbgppeercapsentry)-1]
    }
    return nil
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeercapstable.Cbgppeercapsentry {
        children[cbgppeercapstable.Cbgppeercapsentry[i].GetSegmentPath()] = &cbgppeercapstable.Cbgppeercapsentry[i]
    }
    return children
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetYangName() string { return "cbgpPeerCapsTable" }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) SetParent(parent types.Entity) { cbgppeercapstable.parent = parent }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetParent() types.Entity { return cbgppeercapstable.parent }

func (cbgppeercapstable *CISCOBGP4MIB_Cbgppeercapstable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry
// Each entry represents a capability received from a
// peer with a particular code and an index. When a 
// capability is received multiple times with different
// values during a BGP connection establishment, 
// corresponding entries are differentiated with indices.
type CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetFilter() yfilter.YFilter { return cbgppeercapsentry.YFilter }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) SetFilter(yf yfilter.YFilter) { cbgppeercapsentry.YFilter = yf }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetGoName(yname string) string {
    if yname == "bgpPeerRemoteAddr" { return "Bgppeerremoteaddr" }
    if yname == "cbgpPeerCapCode" { return "Cbgppeercapcode" }
    if yname == "cbgpPeerCapIndex" { return "Cbgppeercapindex" }
    if yname == "cbgpPeerCapValue" { return "Cbgppeercapvalue" }
    return ""
}

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetSegmentPath() string {
    return "cbgpPeerCapsEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", cbgppeercapsentry.Bgppeerremoteaddr) + "']" + "[cbgpPeerCapCode='" + fmt.Sprintf("%v", cbgppeercapsentry.Cbgppeercapcode) + "']" + "[cbgpPeerCapIndex='" + fmt.Sprintf("%v", cbgppeercapsentry.Cbgppeercapindex) + "']"
}

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgpPeerRemoteAddr"] = cbgppeercapsentry.Bgppeerremoteaddr
    leafs["cbgpPeerCapCode"] = cbgppeercapsentry.Cbgppeercapcode
    leafs["cbgpPeerCapIndex"] = cbgppeercapsentry.Cbgppeercapindex
    leafs["cbgpPeerCapValue"] = cbgppeercapsentry.Cbgppeercapvalue
    return leafs
}

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetYangName() string { return "cbgpPeerCapsEntry" }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) SetParent(parent types.Entity) { cbgppeercapsentry.parent = parent }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetParent() types.Entity { return cbgppeercapsentry.parent }

func (cbgppeercapsentry *CISCOBGP4MIB_Cbgppeercapstable_Cbgppeercapsentry) GetParentYangName() string { return "cbgpPeerCapsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // names associated with an address family. The type is slice of
    // CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry.
    Cbgppeeraddrfamilyentry []CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetFilter() yfilter.YFilter { return cbgppeeraddrfamilytable.YFilter }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) SetFilter(yf yfilter.YFilter) { cbgppeeraddrfamilytable.YFilter = yf }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetGoName(yname string) string {
    if yname == "cbgpPeerAddrFamilyEntry" { return "Cbgppeeraddrfamilyentry" }
    return ""
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetSegmentPath() string {
    return "cbgpPeerAddrFamilyTable"
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeerAddrFamilyEntry" {
        for _, c := range cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry {
            if cbgppeeraddrfamilytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry{}
        cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry = append(cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry, child)
        return &cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry[len(cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry)-1]
    }
    return nil
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry {
        children[cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry[i].GetSegmentPath()] = &cbgppeeraddrfamilytable.Cbgppeeraddrfamilyentry[i]
    }
    return children
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetYangName() string { return "cbgpPeerAddrFamilyTable" }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) SetParent(parent types.Entity) { cbgppeeraddrfamilytable.parent = parent }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetParent() types.Entity { return cbgppeeraddrfamilytable.parent }

func (cbgppeeraddrfamilytable *CISCOBGP4MIB_Cbgppeeraddrfamilytable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains names associated with
// an address family.
type CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetFilter() yfilter.YFilter { return cbgppeeraddrfamilyentry.YFilter }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) SetFilter(yf yfilter.YFilter) { cbgppeeraddrfamilyentry.YFilter = yf }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetGoName(yname string) string {
    if yname == "bgpPeerRemoteAddr" { return "Bgppeerremoteaddr" }
    if yname == "cbgpPeerAddrFamilyAfi" { return "Cbgppeeraddrfamilyafi" }
    if yname == "cbgpPeerAddrFamilySafi" { return "Cbgppeeraddrfamilysafi" }
    if yname == "cbgpPeerAddrFamilyName" { return "Cbgppeeraddrfamilyname" }
    return ""
}

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetSegmentPath() string {
    return "cbgpPeerAddrFamilyEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", cbgppeeraddrfamilyentry.Bgppeerremoteaddr) + "']" + "[cbgpPeerAddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilyafi) + "']" + "[cbgpPeerAddrFamilySafi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyentry.Cbgppeeraddrfamilysafi) + "']"
}

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgpPeerRemoteAddr"] = cbgppeeraddrfamilyentry.Bgppeerremoteaddr
    leafs["cbgpPeerAddrFamilyAfi"] = cbgppeeraddrfamilyentry.Cbgppeeraddrfamilyafi
    leafs["cbgpPeerAddrFamilySafi"] = cbgppeeraddrfamilyentry.Cbgppeeraddrfamilysafi
    leafs["cbgpPeerAddrFamilyName"] = cbgppeeraddrfamilyentry.Cbgppeeraddrfamilyname
    return leafs
}

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetYangName() string { return "cbgpPeerAddrFamilyEntry" }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) SetParent(parent types.Entity) { cbgppeeraddrfamilyentry.parent = parent }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetParent() types.Entity { return cbgppeeraddrfamilyentry.parent }

func (cbgppeeraddrfamilyentry *CISCOBGP4MIB_Cbgppeeraddrfamilytable_Cbgppeeraddrfamilyentry) GetParentYangName() string { return "cbgpPeerAddrFamilyTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // information associated  with route prefixes belonging to an address family.
    // The type is slice of
    // CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry.
    Cbgppeeraddrfamilyprefixentry []CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetFilter() yfilter.YFilter { return cbgppeeraddrfamilyprefixtable.YFilter }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) SetFilter(yf yfilter.YFilter) { cbgppeeraddrfamilyprefixtable.YFilter = yf }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetGoName(yname string) string {
    if yname == "cbgpPeerAddrFamilyPrefixEntry" { return "Cbgppeeraddrfamilyprefixentry" }
    return ""
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetSegmentPath() string {
    return "cbgpPeerAddrFamilyPrefixTable"
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeerAddrFamilyPrefixEntry" {
        for _, c := range cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry {
            if cbgppeeraddrfamilyprefixtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry{}
        cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry = append(cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry, child)
        return &cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry[len(cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry)-1]
    }
    return nil
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry {
        children[cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry[i].GetSegmentPath()] = &cbgppeeraddrfamilyprefixtable.Cbgppeeraddrfamilyprefixentry[i]
    }
    return children
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetYangName() string { return "cbgpPeerAddrFamilyPrefixTable" }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) SetParent(parent types.Entity) { cbgppeeraddrfamilyprefixtable.parent = parent }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetParent() types.Entity { return cbgppeeraddrfamilyprefixtable.parent }

func (cbgppeeraddrfamilyprefixtable *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains information associated 
// with route prefixes belonging to an address family.
type CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetFilter() yfilter.YFilter { return cbgppeeraddrfamilyprefixentry.YFilter }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) SetFilter(yf yfilter.YFilter) { cbgppeeraddrfamilyprefixentry.YFilter = yf }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetGoName(yname string) string {
    if yname == "bgpPeerRemoteAddr" { return "Bgppeerremoteaddr" }
    if yname == "cbgpPeerAddrFamilyAfi" { return "Cbgppeeraddrfamilyafi" }
    if yname == "cbgpPeerAddrFamilySafi" { return "Cbgppeeraddrfamilysafi" }
    if yname == "cbgpPeerAcceptedPrefixes" { return "Cbgppeeracceptedprefixes" }
    if yname == "cbgpPeerDeniedPrefixes" { return "Cbgppeerdeniedprefixes" }
    if yname == "cbgpPeerPrefixAdminLimit" { return "Cbgppeerprefixadminlimit" }
    if yname == "cbgpPeerPrefixThreshold" { return "Cbgppeerprefixthreshold" }
    if yname == "cbgpPeerPrefixClearThreshold" { return "Cbgppeerprefixclearthreshold" }
    if yname == "cbgpPeerAdvertisedPrefixes" { return "Cbgppeeradvertisedprefixes" }
    if yname == "cbgpPeerSuppressedPrefixes" { return "Cbgppeersuppressedprefixes" }
    if yname == "cbgpPeerWithdrawnPrefixes" { return "Cbgppeerwithdrawnprefixes" }
    return ""
}

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetSegmentPath() string {
    return "cbgpPeerAddrFamilyPrefixEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", cbgppeeraddrfamilyprefixentry.Bgppeerremoteaddr) + "']" + "[cbgpPeerAddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilyafi) + "']" + "[cbgpPeerAddrFamilySafi='" + fmt.Sprintf("%v", cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilysafi) + "']"
}

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgpPeerRemoteAddr"] = cbgppeeraddrfamilyprefixentry.Bgppeerremoteaddr
    leafs["cbgpPeerAddrFamilyAfi"] = cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilyafi
    leafs["cbgpPeerAddrFamilySafi"] = cbgppeeraddrfamilyprefixentry.Cbgppeeraddrfamilysafi
    leafs["cbgpPeerAcceptedPrefixes"] = cbgppeeraddrfamilyprefixentry.Cbgppeeracceptedprefixes
    leafs["cbgpPeerDeniedPrefixes"] = cbgppeeraddrfamilyprefixentry.Cbgppeerdeniedprefixes
    leafs["cbgpPeerPrefixAdminLimit"] = cbgppeeraddrfamilyprefixentry.Cbgppeerprefixadminlimit
    leafs["cbgpPeerPrefixThreshold"] = cbgppeeraddrfamilyprefixentry.Cbgppeerprefixthreshold
    leafs["cbgpPeerPrefixClearThreshold"] = cbgppeeraddrfamilyprefixentry.Cbgppeerprefixclearthreshold
    leafs["cbgpPeerAdvertisedPrefixes"] = cbgppeeraddrfamilyprefixentry.Cbgppeeradvertisedprefixes
    leafs["cbgpPeerSuppressedPrefixes"] = cbgppeeraddrfamilyprefixentry.Cbgppeersuppressedprefixes
    leafs["cbgpPeerWithdrawnPrefixes"] = cbgppeeraddrfamilyprefixentry.Cbgppeerwithdrawnprefixes
    return leafs
}

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetYangName() string { return "cbgpPeerAddrFamilyPrefixEntry" }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) SetParent(parent types.Entity) { cbgppeeraddrfamilyprefixentry.parent = parent }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetParent() types.Entity { return cbgppeeraddrfamilyprefixentry.parent }

func (cbgppeeraddrfamilyprefixentry *CISCOBGP4MIB_Cbgppeeraddrfamilyprefixtable_Cbgppeeraddrfamilyprefixentry) GetParentYangName() string { return "cbgpPeerAddrFamilyPrefixTable" }

// CISCOBGP4MIB_Cbgppeer2Table
// BGP peer table.  This table contains,
// one entry per BGP peer, information about
// the connections with BGP peers.
type CISCOBGP4MIB_Cbgppeer2Table struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entry containing information about the connection with a BGP peer. The type
    // is slice of CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry.
    Cbgppeer2Entry []CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetFilter() yfilter.YFilter { return cbgppeer2Table.YFilter }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) SetFilter(yf yfilter.YFilter) { cbgppeer2Table.YFilter = yf }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetGoName(yname string) string {
    if yname == "cbgpPeer2Entry" { return "Cbgppeer2Entry" }
    return ""
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetSegmentPath() string {
    return "cbgpPeer2Table"
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeer2Entry" {
        for _, c := range cbgppeer2Table.Cbgppeer2Entry {
            if cbgppeer2Table.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry{}
        cbgppeer2Table.Cbgppeer2Entry = append(cbgppeer2Table.Cbgppeer2Entry, child)
        return &cbgppeer2Table.Cbgppeer2Entry[len(cbgppeer2Table.Cbgppeer2Entry)-1]
    }
    return nil
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeer2Table.Cbgppeer2Entry {
        children[cbgppeer2Table.Cbgppeer2Entry[i].GetSegmentPath()] = &cbgppeer2Table.Cbgppeer2Entry[i]
    }
    return children
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetYangName() string { return "cbgpPeer2Table" }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) SetParent(parent types.Entity) { cbgppeer2Table.parent = parent }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetParent() types.Entity { return cbgppeer2Table.parent }

func (cbgppeer2Table *CISCOBGP4MIB_Cbgppeer2Table) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry
// Entry containing information about the
// connection with a BGP peer.
type CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetFilter() yfilter.YFilter { return cbgppeer2Entry.YFilter }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) SetFilter(yf yfilter.YFilter) { cbgppeer2Entry.YFilter = yf }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetGoName(yname string) string {
    if yname == "cbgpPeer2Type" { return "Cbgppeer2Type" }
    if yname == "cbgpPeer2RemoteAddr" { return "Cbgppeer2Remoteaddr" }
    if yname == "cbgpPeer2State" { return "Cbgppeer2State" }
    if yname == "cbgpPeer2AdminStatus" { return "Cbgppeer2Adminstatus" }
    if yname == "cbgpPeer2NegotiatedVersion" { return "Cbgppeer2Negotiatedversion" }
    if yname == "cbgpPeer2LocalAddr" { return "Cbgppeer2Localaddr" }
    if yname == "cbgpPeer2LocalPort" { return "Cbgppeer2Localport" }
    if yname == "cbgpPeer2LocalAs" { return "Cbgppeer2Localas" }
    if yname == "cbgpPeer2LocalIdentifier" { return "Cbgppeer2Localidentifier" }
    if yname == "cbgpPeer2RemotePort" { return "Cbgppeer2Remoteport" }
    if yname == "cbgpPeer2RemoteAs" { return "Cbgppeer2Remoteas" }
    if yname == "cbgpPeer2RemoteIdentifier" { return "Cbgppeer2Remoteidentifier" }
    if yname == "cbgpPeer2InUpdates" { return "Cbgppeer2Inupdates" }
    if yname == "cbgpPeer2OutUpdates" { return "Cbgppeer2Outupdates" }
    if yname == "cbgpPeer2InTotalMessages" { return "Cbgppeer2Intotalmessages" }
    if yname == "cbgpPeer2OutTotalMessages" { return "Cbgppeer2Outtotalmessages" }
    if yname == "cbgpPeer2LastError" { return "Cbgppeer2Lasterror" }
    if yname == "cbgpPeer2FsmEstablishedTransitions" { return "Cbgppeer2Fsmestablishedtransitions" }
    if yname == "cbgpPeer2FsmEstablishedTime" { return "Cbgppeer2Fsmestablishedtime" }
    if yname == "cbgpPeer2ConnectRetryInterval" { return "Cbgppeer2Connectretryinterval" }
    if yname == "cbgpPeer2HoldTime" { return "Cbgppeer2Holdtime" }
    if yname == "cbgpPeer2KeepAlive" { return "Cbgppeer2Keepalive" }
    if yname == "cbgpPeer2HoldTimeConfigured" { return "Cbgppeer2Holdtimeconfigured" }
    if yname == "cbgpPeer2KeepAliveConfigured" { return "Cbgppeer2Keepaliveconfigured" }
    if yname == "cbgpPeer2MinASOriginationInterval" { return "Cbgppeer2Minasoriginationinterval" }
    if yname == "cbgpPeer2MinRouteAdvertisementInterval" { return "Cbgppeer2Minrouteadvertisementinterval" }
    if yname == "cbgpPeer2InUpdateElapsedTime" { return "Cbgppeer2Inupdateelapsedtime" }
    if yname == "cbgpPeer2LastErrorTxt" { return "Cbgppeer2Lasterrortxt" }
    if yname == "cbgpPeer2PrevState" { return "Cbgppeer2Prevstate" }
    return ""
}

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetSegmentPath() string {
    return "cbgpPeer2Entry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Entry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Entry.Cbgppeer2Remoteaddr) + "']"
}

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbgpPeer2Type"] = cbgppeer2Entry.Cbgppeer2Type
    leafs["cbgpPeer2RemoteAddr"] = cbgppeer2Entry.Cbgppeer2Remoteaddr
    leafs["cbgpPeer2State"] = cbgppeer2Entry.Cbgppeer2State
    leafs["cbgpPeer2AdminStatus"] = cbgppeer2Entry.Cbgppeer2Adminstatus
    leafs["cbgpPeer2NegotiatedVersion"] = cbgppeer2Entry.Cbgppeer2Negotiatedversion
    leafs["cbgpPeer2LocalAddr"] = cbgppeer2Entry.Cbgppeer2Localaddr
    leafs["cbgpPeer2LocalPort"] = cbgppeer2Entry.Cbgppeer2Localport
    leafs["cbgpPeer2LocalAs"] = cbgppeer2Entry.Cbgppeer2Localas
    leafs["cbgpPeer2LocalIdentifier"] = cbgppeer2Entry.Cbgppeer2Localidentifier
    leafs["cbgpPeer2RemotePort"] = cbgppeer2Entry.Cbgppeer2Remoteport
    leafs["cbgpPeer2RemoteAs"] = cbgppeer2Entry.Cbgppeer2Remoteas
    leafs["cbgpPeer2RemoteIdentifier"] = cbgppeer2Entry.Cbgppeer2Remoteidentifier
    leafs["cbgpPeer2InUpdates"] = cbgppeer2Entry.Cbgppeer2Inupdates
    leafs["cbgpPeer2OutUpdates"] = cbgppeer2Entry.Cbgppeer2Outupdates
    leafs["cbgpPeer2InTotalMessages"] = cbgppeer2Entry.Cbgppeer2Intotalmessages
    leafs["cbgpPeer2OutTotalMessages"] = cbgppeer2Entry.Cbgppeer2Outtotalmessages
    leafs["cbgpPeer2LastError"] = cbgppeer2Entry.Cbgppeer2Lasterror
    leafs["cbgpPeer2FsmEstablishedTransitions"] = cbgppeer2Entry.Cbgppeer2Fsmestablishedtransitions
    leafs["cbgpPeer2FsmEstablishedTime"] = cbgppeer2Entry.Cbgppeer2Fsmestablishedtime
    leafs["cbgpPeer2ConnectRetryInterval"] = cbgppeer2Entry.Cbgppeer2Connectretryinterval
    leafs["cbgpPeer2HoldTime"] = cbgppeer2Entry.Cbgppeer2Holdtime
    leafs["cbgpPeer2KeepAlive"] = cbgppeer2Entry.Cbgppeer2Keepalive
    leafs["cbgpPeer2HoldTimeConfigured"] = cbgppeer2Entry.Cbgppeer2Holdtimeconfigured
    leafs["cbgpPeer2KeepAliveConfigured"] = cbgppeer2Entry.Cbgppeer2Keepaliveconfigured
    leafs["cbgpPeer2MinASOriginationInterval"] = cbgppeer2Entry.Cbgppeer2Minasoriginationinterval
    leafs["cbgpPeer2MinRouteAdvertisementInterval"] = cbgppeer2Entry.Cbgppeer2Minrouteadvertisementinterval
    leafs["cbgpPeer2InUpdateElapsedTime"] = cbgppeer2Entry.Cbgppeer2Inupdateelapsedtime
    leafs["cbgpPeer2LastErrorTxt"] = cbgppeer2Entry.Cbgppeer2Lasterrortxt
    leafs["cbgpPeer2PrevState"] = cbgppeer2Entry.Cbgppeer2Prevstate
    return leafs
}

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetYangName() string { return "cbgpPeer2Entry" }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) SetParent(parent types.Entity) { cbgppeer2Entry.parent = parent }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetParent() types.Entity { return cbgppeer2Entry.parent }

func (cbgppeer2Entry *CISCOBGP4MIB_Cbgppeer2Table_Cbgppeer2Entry) GetParentYangName() string { return "cbgpPeer2Table" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a capability received from a peer with a particular
    // code and an index. When a capability is received multiple times with
    // different values during a BGP connection establishment, corresponding
    // entries are differentiated with indices. The type is slice of
    // CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry.
    Cbgppeer2Capsentry []CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetFilter() yfilter.YFilter { return cbgppeer2Capstable.YFilter }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) SetFilter(yf yfilter.YFilter) { cbgppeer2Capstable.YFilter = yf }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetGoName(yname string) string {
    if yname == "cbgpPeer2CapsEntry" { return "Cbgppeer2Capsentry" }
    return ""
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetSegmentPath() string {
    return "cbgpPeer2CapsTable"
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeer2CapsEntry" {
        for _, c := range cbgppeer2Capstable.Cbgppeer2Capsentry {
            if cbgppeer2Capstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry{}
        cbgppeer2Capstable.Cbgppeer2Capsentry = append(cbgppeer2Capstable.Cbgppeer2Capsentry, child)
        return &cbgppeer2Capstable.Cbgppeer2Capsentry[len(cbgppeer2Capstable.Cbgppeer2Capsentry)-1]
    }
    return nil
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeer2Capstable.Cbgppeer2Capsentry {
        children[cbgppeer2Capstable.Cbgppeer2Capsentry[i].GetSegmentPath()] = &cbgppeer2Capstable.Cbgppeer2Capsentry[i]
    }
    return children
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetYangName() string { return "cbgpPeer2CapsTable" }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) SetParent(parent types.Entity) { cbgppeer2Capstable.parent = parent }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetParent() types.Entity { return cbgppeer2Capstable.parent }

func (cbgppeer2Capstable *CISCOBGP4MIB_Cbgppeer2Capstable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry
// Each entry represents a capability received from a
// peer with a particular code and an index. When a
// capability is received multiple times with different
// values during a BGP connection establishment,
// corresponding entries are differentiated with indices.
type CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry struct {
    parent types.Entity
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

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetFilter() yfilter.YFilter { return cbgppeer2Capsentry.YFilter }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) SetFilter(yf yfilter.YFilter) { cbgppeer2Capsentry.YFilter = yf }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetGoName(yname string) string {
    if yname == "cbgpPeer2Type" { return "Cbgppeer2Type" }
    if yname == "cbgpPeer2RemoteAddr" { return "Cbgppeer2Remoteaddr" }
    if yname == "cbgpPeer2CapCode" { return "Cbgppeer2Capcode" }
    if yname == "cbgpPeer2CapIndex" { return "Cbgppeer2Capindex" }
    if yname == "cbgpPeer2CapValue" { return "Cbgppeer2Capvalue" }
    return ""
}

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetSegmentPath() string {
    return "cbgpPeer2CapsEntry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Remoteaddr) + "']" + "[cbgpPeer2CapCode='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Capcode) + "']" + "[cbgpPeer2CapIndex='" + fmt.Sprintf("%v", cbgppeer2Capsentry.Cbgppeer2Capindex) + "']"
}

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbgpPeer2Type"] = cbgppeer2Capsentry.Cbgppeer2Type
    leafs["cbgpPeer2RemoteAddr"] = cbgppeer2Capsentry.Cbgppeer2Remoteaddr
    leafs["cbgpPeer2CapCode"] = cbgppeer2Capsentry.Cbgppeer2Capcode
    leafs["cbgpPeer2CapIndex"] = cbgppeer2Capsentry.Cbgppeer2Capindex
    leafs["cbgpPeer2CapValue"] = cbgppeer2Capsentry.Cbgppeer2Capvalue
    return leafs
}

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetYangName() string { return "cbgpPeer2CapsEntry" }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) SetParent(parent types.Entity) { cbgppeer2Capsentry.parent = parent }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetParent() types.Entity { return cbgppeer2Capsentry.parent }

func (cbgppeer2Capsentry *CISCOBGP4MIB_Cbgppeer2Capstable_Cbgppeer2Capsentry) GetParentYangName() string { return "cbgpPeer2CapsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // names associated with an address family. The type is slice of
    // CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry.
    Cbgppeer2Addrfamilyentry []CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetFilter() yfilter.YFilter { return cbgppeer2Addrfamilytable.YFilter }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) SetFilter(yf yfilter.YFilter) { cbgppeer2Addrfamilytable.YFilter = yf }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetGoName(yname string) string {
    if yname == "cbgpPeer2AddrFamilyEntry" { return "Cbgppeer2Addrfamilyentry" }
    return ""
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetSegmentPath() string {
    return "cbgpPeer2AddrFamilyTable"
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeer2AddrFamilyEntry" {
        for _, c := range cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry {
            if cbgppeer2Addrfamilytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry{}
        cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry = append(cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry, child)
        return &cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry[len(cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry)-1]
    }
    return nil
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry {
        children[cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry[i].GetSegmentPath()] = &cbgppeer2Addrfamilytable.Cbgppeer2Addrfamilyentry[i]
    }
    return children
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetYangName() string { return "cbgpPeer2AddrFamilyTable" }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) SetParent(parent types.Entity) { cbgppeer2Addrfamilytable.parent = parent }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetParent() types.Entity { return cbgppeer2Addrfamilytable.parent }

func (cbgppeer2Addrfamilytable *CISCOBGP4MIB_Cbgppeer2Addrfamilytable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains names associated with
// an address family.
type CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry struct {
    parent types.Entity
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

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetFilter() yfilter.YFilter { return cbgppeer2Addrfamilyentry.YFilter }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) SetFilter(yf yfilter.YFilter) { cbgppeer2Addrfamilyentry.YFilter = yf }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetGoName(yname string) string {
    if yname == "cbgpPeer2Type" { return "Cbgppeer2Type" }
    if yname == "cbgpPeer2RemoteAddr" { return "Cbgppeer2Remoteaddr" }
    if yname == "cbgpPeer2AddrFamilyAfi" { return "Cbgppeer2Addrfamilyafi" }
    if yname == "cbgpPeer2AddrFamilySafi" { return "Cbgppeer2Addrfamilysafi" }
    if yname == "cbgpPeer2AddrFamilyName" { return "Cbgppeer2Addrfamilyname" }
    return ""
}

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetSegmentPath() string {
    return "cbgpPeer2AddrFamilyEntry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Remoteaddr) + "']" + "[cbgpPeer2AddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilyafi) + "']" + "[cbgpPeer2AddrFamilySafi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilysafi) + "']"
}

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbgpPeer2Type"] = cbgppeer2Addrfamilyentry.Cbgppeer2Type
    leafs["cbgpPeer2RemoteAddr"] = cbgppeer2Addrfamilyentry.Cbgppeer2Remoteaddr
    leafs["cbgpPeer2AddrFamilyAfi"] = cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilyafi
    leafs["cbgpPeer2AddrFamilySafi"] = cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilysafi
    leafs["cbgpPeer2AddrFamilyName"] = cbgppeer2Addrfamilyentry.Cbgppeer2Addrfamilyname
    return leafs
}

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetYangName() string { return "cbgpPeer2AddrFamilyEntry" }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) SetParent(parent types.Entity) { cbgppeer2Addrfamilyentry.parent = parent }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetParent() types.Entity { return cbgppeer2Addrfamilyentry.parent }

func (cbgppeer2Addrfamilyentry *CISCOBGP4MIB_Cbgppeer2Addrfamilytable_Cbgppeer2Addrfamilyentry) GetParentYangName() string { return "cbgpPeer2AddrFamilyTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is identified by an AFI/SAFI pair and peer address. It contains
    // information associated with route prefixes belonging to an address family.
    // The type is slice of
    // CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry.
    Cbgppeer2Addrfamilyprefixentry []CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetFilter() yfilter.YFilter { return cbgppeer2Addrfamilyprefixtable.YFilter }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) SetFilter(yf yfilter.YFilter) { cbgppeer2Addrfamilyprefixtable.YFilter = yf }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetGoName(yname string) string {
    if yname == "cbgpPeer2AddrFamilyPrefixEntry" { return "Cbgppeer2Addrfamilyprefixentry" }
    return ""
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetSegmentPath() string {
    return "cbgpPeer2AddrFamilyPrefixTable"
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbgpPeer2AddrFamilyPrefixEntry" {
        for _, c := range cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry {
            if cbgppeer2Addrfamilyprefixtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry{}
        cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry = append(cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry, child)
        return &cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry[len(cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry)-1]
    }
    return nil
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry {
        children[cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry[i].GetSegmentPath()] = &cbgppeer2Addrfamilyprefixtable.Cbgppeer2Addrfamilyprefixentry[i]
    }
    return children
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetYangName() string { return "cbgpPeer2AddrFamilyPrefixTable" }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) SetParent(parent types.Entity) { cbgppeer2Addrfamilyprefixtable.parent = parent }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetParent() types.Entity { return cbgppeer2Addrfamilyprefixtable.parent }

func (cbgppeer2Addrfamilyprefixtable *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable) GetParentYangName() string { return "CISCO-BGP4-MIB" }

// CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry
// An entry is identified by an AFI/SAFI pair and
// peer address. It contains information associated
// with route prefixes belonging to an address family.
type CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry struct {
    parent types.Entity
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

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetFilter() yfilter.YFilter { return cbgppeer2Addrfamilyprefixentry.YFilter }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) SetFilter(yf yfilter.YFilter) { cbgppeer2Addrfamilyprefixentry.YFilter = yf }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetGoName(yname string) string {
    if yname == "cbgpPeer2Type" { return "Cbgppeer2Type" }
    if yname == "cbgpPeer2RemoteAddr" { return "Cbgppeer2Remoteaddr" }
    if yname == "cbgpPeer2AddrFamilyAfi" { return "Cbgppeer2Addrfamilyafi" }
    if yname == "cbgpPeer2AddrFamilySafi" { return "Cbgppeer2Addrfamilysafi" }
    if yname == "cbgpPeer2AcceptedPrefixes" { return "Cbgppeer2Acceptedprefixes" }
    if yname == "cbgpPeer2DeniedPrefixes" { return "Cbgppeer2Deniedprefixes" }
    if yname == "cbgpPeer2PrefixAdminLimit" { return "Cbgppeer2Prefixadminlimit" }
    if yname == "cbgpPeer2PrefixThreshold" { return "Cbgppeer2Prefixthreshold" }
    if yname == "cbgpPeer2PrefixClearThreshold" { return "Cbgppeer2Prefixclearthreshold" }
    if yname == "cbgpPeer2AdvertisedPrefixes" { return "Cbgppeer2Advertisedprefixes" }
    if yname == "cbgpPeer2SuppressedPrefixes" { return "Cbgppeer2Suppressedprefixes" }
    if yname == "cbgpPeer2WithdrawnPrefixes" { return "Cbgppeer2Withdrawnprefixes" }
    return ""
}

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetSegmentPath() string {
    return "cbgpPeer2AddrFamilyPrefixEntry" + "[cbgpPeer2Type='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Type) + "']" + "[cbgpPeer2RemoteAddr='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Remoteaddr) + "']" + "[cbgpPeer2AddrFamilyAfi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilyafi) + "']" + "[cbgpPeer2AddrFamilySafi='" + fmt.Sprintf("%v", cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilysafi) + "']"
}

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbgpPeer2Type"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Type
    leafs["cbgpPeer2RemoteAddr"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Remoteaddr
    leafs["cbgpPeer2AddrFamilyAfi"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilyafi
    leafs["cbgpPeer2AddrFamilySafi"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Addrfamilysafi
    leafs["cbgpPeer2AcceptedPrefixes"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Acceptedprefixes
    leafs["cbgpPeer2DeniedPrefixes"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Deniedprefixes
    leafs["cbgpPeer2PrefixAdminLimit"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Prefixadminlimit
    leafs["cbgpPeer2PrefixThreshold"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Prefixthreshold
    leafs["cbgpPeer2PrefixClearThreshold"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Prefixclearthreshold
    leafs["cbgpPeer2AdvertisedPrefixes"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Advertisedprefixes
    leafs["cbgpPeer2SuppressedPrefixes"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Suppressedprefixes
    leafs["cbgpPeer2WithdrawnPrefixes"] = cbgppeer2Addrfamilyprefixentry.Cbgppeer2Withdrawnprefixes
    return leafs
}

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetYangName() string { return "cbgpPeer2AddrFamilyPrefixEntry" }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) SetParent(parent types.Entity) { cbgppeer2Addrfamilyprefixentry.parent = parent }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetParent() types.Entity { return cbgppeer2Addrfamilyprefixentry.parent }

func (cbgppeer2Addrfamilyprefixentry *CISCOBGP4MIB_Cbgppeer2Addrfamilyprefixtable_Cbgppeer2Addrfamilyprefixentry) GetParentYangName() string { return "cbgpPeer2AddrFamilyPrefixTable" }

