// The MIB module for BGP-4.
package bgp4_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp4_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:BGP4-MIB BGP4-MIB}", reflect.TypeOf(BGP4MIB{}))
    ydk.RegisterEntity("BGP4-MIB:BGP4-MIB", reflect.TypeOf(BGP4MIB{}))
}

// BGP4MIB
type BGP4MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Bgp BGP4MIB_Bgp

    // BGP peer table.  This table contains, one entry per BGP peer, information
    // about the connections with BGP peers.
    BgpPeerTable BGP4MIB_BgpPeerTable

    // The BGP Received Path Attribute Table contains information about paths to
    // destination networks received from all peers running BGP version 3 or less.
    BgpRcvdPathAttrTable BGP4MIB_BgpRcvdPathAttrTable

    // The BGP-4 Received Path Attribute Table contains information about paths to
    // destination networks received from all BGP4 peers.
    Bgp4PathAttrTable BGP4MIB_Bgp4PathAttrTable
}

func (bGP4MIB *BGP4MIB) GetEntityData() *types.CommonEntityData {
    bGP4MIB.EntityData.YFilter = bGP4MIB.YFilter
    bGP4MIB.EntityData.YangName = "BGP4-MIB"
    bGP4MIB.EntityData.BundleName = "cisco_ios_xe"
    bGP4MIB.EntityData.ParentYangName = "BGP4-MIB"
    bGP4MIB.EntityData.SegmentPath = "BGP4-MIB:BGP4-MIB"
    bGP4MIB.EntityData.AbsolutePath = bGP4MIB.EntityData.SegmentPath
    bGP4MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bGP4MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bGP4MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bGP4MIB.EntityData.Children = types.NewOrderedMap()
    bGP4MIB.EntityData.Children.Append("bgp", types.YChild{"Bgp", &bGP4MIB.Bgp})
    bGP4MIB.EntityData.Children.Append("bgpPeerTable", types.YChild{"BgpPeerTable", &bGP4MIB.BgpPeerTable})
    bGP4MIB.EntityData.Children.Append("bgpRcvdPathAttrTable", types.YChild{"BgpRcvdPathAttrTable", &bGP4MIB.BgpRcvdPathAttrTable})
    bGP4MIB.EntityData.Children.Append("bgp4PathAttrTable", types.YChild{"Bgp4PathAttrTable", &bGP4MIB.Bgp4PathAttrTable})
    bGP4MIB.EntityData.Leafs = types.NewOrderedMap()

    bGP4MIB.EntityData.YListKeys = []string {}

    return &(bGP4MIB.EntityData)
}

// BGP4MIB_Bgp
type BGP4MIB_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vector of supported BGP protocol version numbers.  Each peer negotiates the
    // version from this vector.  Versions are identified via the string of bits
    // contained within this object.  The first octet contains bits 0 to 7, the
    // second octet contains bits 8 to 15, and so on, with the most significant
    // bit referring to the lowest bit number in the octet (e.g., the MSB of the
    // first octet refers to bit 0).  If a bit, i, is present and set, then the
    // version (i+1) of the BGP is supported. The type is string with length:
    // 1..255.
    BgpVersion interface{}

    // The local autonomous system number. The type is interface{} with range:
    // 0..65535.
    BgpLocalAs interface{}

    // The BGP Identifier of local system. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpIdentifier interface{}
}

func (bgp *BGP4MIB_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xe"
    bgp.EntityData.ParentYangName = "BGP4-MIB"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/" + bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("bgpVersion", types.YLeaf{"BgpVersion", bgp.BgpVersion})
    bgp.EntityData.Leafs.Append("bgpLocalAs", types.YLeaf{"BgpLocalAs", bgp.BgpLocalAs})
    bgp.EntityData.Leafs.Append("bgpIdentifier", types.YLeaf{"BgpIdentifier", bgp.BgpIdentifier})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// BGP4MIB_BgpPeerTable
// BGP peer table.  This table contains,
// one entry per BGP peer, information about
// the connections with BGP peers.
type BGP4MIB_BgpPeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry containing information about the connection with a BGP peer. The type
    // is slice of BGP4MIB_BgpPeerTable_BgpPeerEntry.
    BgpPeerEntry []*BGP4MIB_BgpPeerTable_BgpPeerEntry
}

func (bgpPeerTable *BGP4MIB_BgpPeerTable) GetEntityData() *types.CommonEntityData {
    bgpPeerTable.EntityData.YFilter = bgpPeerTable.YFilter
    bgpPeerTable.EntityData.YangName = "bgpPeerTable"
    bgpPeerTable.EntityData.BundleName = "cisco_ios_xe"
    bgpPeerTable.EntityData.ParentYangName = "BGP4-MIB"
    bgpPeerTable.EntityData.SegmentPath = "bgpPeerTable"
    bgpPeerTable.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/" + bgpPeerTable.EntityData.SegmentPath
    bgpPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPeerTable.EntityData.Children = types.NewOrderedMap()
    bgpPeerTable.EntityData.Children.Append("bgpPeerEntry", types.YChild{"BgpPeerEntry", nil})
    for i := range bgpPeerTable.BgpPeerEntry {
        bgpPeerTable.EntityData.Children.Append(types.GetSegmentPath(bgpPeerTable.BgpPeerEntry[i]), types.YChild{"BgpPeerEntry", bgpPeerTable.BgpPeerEntry[i]})
    }
    bgpPeerTable.EntityData.Leafs = types.NewOrderedMap()

    bgpPeerTable.EntityData.YListKeys = []string {}

    return &(bgpPeerTable.EntityData)
}

// BGP4MIB_BgpPeerTable_BgpPeerEntry
// Entry containing information about the
// connection with a BGP peer.
type BGP4MIB_BgpPeerTable_BgpPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The remote IP address of this entry's BGP peer.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpPeerRemoteAddr interface{}

    // The BGP Identifier of this entry's BGP peer. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpPeerIdentifier interface{}

    // The BGP peer connection state. The type is BgpPeerState.
    BgpPeerState interface{}

    // The desired state of the BGP connection. A transition from 'stop' to
    // 'start' will cause the BGP Start Event to be generated. A transition from
    // 'start' to 'stop' will cause the BGP Stop Event to be generated. This
    // parameter can be used to restart BGP peer connections.  Care should be used
    // in providing write access to this object without adequate authentication.
    // The type is BgpPeerAdminStatus.
    BgpPeerAdminStatus interface{}

    // The negotiated version of BGP running between the two peers. The type is
    // interface{} with range: -2147483648..2147483647.
    BgpPeerNegotiatedVersion interface{}

    // The local IP address of this entry's BGP connection. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpPeerLocalAddr interface{}

    // The local port for the TCP connection between the BGP peers. The type is
    // interface{} with range: 0..65535.
    BgpPeerLocalPort interface{}

    // The remote port for the TCP connection between the BGP peers.  Note that
    // the objects bgpPeerLocalAddr, bgpPeerLocalPort, bgpPeerRemoteAddr and
    // bgpPeerRemotePort provide the appropriate reference to the standard MIB TCP
    // connection table. The type is interface{} with range: 0..65535.
    BgpPeerRemotePort interface{}

    // The remote autonomous system number. The type is interface{} with range:
    // 0..65535.
    BgpPeerRemoteAs interface{}

    // The number of BGP UPDATE messages received on this connection.  This object
    // should be initialized to zero (0) when the connection is established. The
    // type is interface{} with range: 0..4294967295.
    BgpPeerInUpdates interface{}

    // The number of BGP UPDATE messages transmitted on this connection.  This
    // object should be initialized to zero (0) when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    BgpPeerOutUpdates interface{}

    // The total number of messages received from the remote peer on this
    // connection. This object should be initialized to zero when the connection
    // is established. The type is interface{} with range: 0..4294967295.
    BgpPeerInTotalMessages interface{}

    // The total number of messages transmitted to the remote peer on this
    // connection.  This object should be initialized to zero when the connection
    // is established. The type is interface{} with range: 0..4294967295.
    BgpPeerOutTotalMessages interface{}

    // The last error code and subcode seen by this peer on this connection.  If
    // no error has occurred, this field is zero.  Otherwise, the first byte of
    // this two byte OCTET STRING contains the error code, and the second byte
    // contains the subcode. The type is string with length: 2.
    BgpPeerLastError interface{}

    // The total number of times the BGP FSM transitioned into the established
    // state. The type is interface{} with range: 0..4294967295.
    BgpPeerFsmEstablishedTransitions interface{}

    // This timer indicates how long (in seconds) this peer has been in the
    // Established state or how long since this peer was last in the Established
    // state.  It is set to zero when a new peer is configured or the router is
    // booted. The type is interface{} with range: 0..4294967295.
    BgpPeerFsmEstablishedTime interface{}

    // Time interval in seconds for the ConnectRetry timer.  The suggested value
    // for this timer is 120 seconds. The type is interface{} with range:
    // 1..65535.
    BgpPeerConnectRetryInterval interface{}

    // Time interval in seconds for the Hold Timer established with the peer.  The
    // value of this object is calculated by this BGP speaker by using the smaller
    // of the value in bgpPeerHoldTimeConfigured and the Hold Time received in the
    // OPEN message. This value must be at lease three seconds if it is not zero
    // (0) in which case the Hold Timer has not been established with the peer,
    // or, the value of bgpPeerHoldTimeConfigured is zero (0). The type is
    // interface{} with range: 0..None | 3..65535.
    BgpPeerHoldTime interface{}

    // Time interval in seconds for the KeepAlive timer established with the peer.
    // The value of this object is calculated by this BGP speaker such that, when
    // compared with bgpPeerHoldTime, it has the same proportion as what
    // bgpPeerKeepAliveConfigured has when compared with
    // bgpPeerHoldTimeConfigured. If the value of this object is zero (0), it
    // indicates that the KeepAlive timer has not been established with the peer,
    // or, the value of bgpPeerKeepAliveConfigured is zero (0). The type is
    // interface{} with range: 0..21845.
    BgpPeerKeepAlive interface{}

    // Time interval in seconds for the Hold Time configured for this BGP speaker
    // with this peer.  This value is placed in an OPEN message sent to this peer
    // by this BGP speaker, and is compared with the Hold Time field in an OPEN
    // message received from the peer when determining the Hold Time
    // (bgpPeerHoldTime) with the peer. This value must not be less than three
    // seconds if it is not zero (0) in which case the Hold Time is NOT to be
    // established with the peer.  The suggested value for this timer is 90
    // seconds. The type is interface{} with range: 0..None | 3..65535.
    BgpPeerHoldTimeConfigured interface{}

    // Time interval in seconds for the KeepAlive timer configured for this BGP
    // speaker with this peer.  The value of this object will only determine the
    // KEEPALIVE messages' frequency relative to the value specified in
    // bgpPeerHoldTimeConfigured; the actual time interval for the KEEPALIVE
    // messages is indicated by bgpPeerKeepAlive.  A reasonable maximum value for
    // this timer would be configured to be one third of that of
    // bgpPeerHoldTimeConfigured. If the value of this object is zero (0), no
    // periodical KEEPALIVE messages are sent to the peer after the BGP connection
    // has been established.  The suggested value for this timer is 30 seconds.
    // The type is interface{} with range: 0..21845.
    BgpPeerKeepAliveConfigured interface{}

    // Time interval in seconds for the MinASOriginationInterval timer. The
    // suggested value for this timer is 15 seconds. The type is interface{} with
    // range: 1..65535.
    BgpPeerMinASOriginationInterval interface{}

    // Time interval in seconds for the MinRouteAdvertisementInterval timer. The
    // suggested value for this timer is 30 seconds. The type is interface{} with
    // range: 1..65535.
    BgpPeerMinRouteAdvertisementInterval interface{}

    // Elapsed time in seconds since the last BGP UPDATE message was received from
    // the peer. Each time bgpPeerInUpdates is incremented, the value of this
    // object is set to zero (0). The type is interface{} with range:
    // 0..4294967295.
    BgpPeerInUpdateElapsedTime interface{}

    // Number of Route prefixes received on this connnection, which are accepted
    // after applying filters. Possible filters are route maps, prefix lists,
    // distributed lists, etc. The type is interface{} with range: 0..4294967295.
    CbgpPeerPrefixAccepted interface{}

    // Counter which gets incremented when a route prefix received on this
    // connection is denied  or when a route prefix is denied during soft reset of
    // this connection if 'soft-reconfiguration' is on . This object is 
    // initialized to zero when the peer is  configured or the router is rebooted.
    // The type is interface{} with range: 0..4294967295.
    CbgpPeerPrefixDenied interface{}

    // Max number of route prefixes accepted on this connection. The type is
    // interface{} with range: 1..4294967295.
    CbgpPeerPrefixLimit interface{}

    // Counter which gets incremented when a route prefix is advertised on this
    // connection. This object is initialized to zero when the peer is configured
    // or  the router is rebooted. The type is interface{} with range:
    // 0..4294967295.
    CbgpPeerPrefixAdvertised interface{}

    // Counter which gets incremented when a route prefix is suppressed from being
    // sent on this connection. This  object is initialized to zero when the peer
    // is  configured or the router is rebooted. The type is interface{} with
    // range: 0..4294967295.
    CbgpPeerPrefixSuppressed interface{}

    // Counter which gets incremented when a route prefix is withdrawn on this
    // connection. This object is initialized to zero when the peer is configured
    // or the router is rebooted. The type is interface{} with range:
    // 0..4294967295.
    CbgpPeerPrefixWithdrawn interface{}

    // Implementation specific error description for bgpPeerLastErrorReceived. The
    // type is string.
    CbgpPeerLastErrorTxt interface{}

    // The BGP peer connection previous state. The type is CbgpPeerPrevState.
    CbgpPeerPrevState interface{}
}

func (bgpPeerEntry *BGP4MIB_BgpPeerTable_BgpPeerEntry) GetEntityData() *types.CommonEntityData {
    bgpPeerEntry.EntityData.YFilter = bgpPeerEntry.YFilter
    bgpPeerEntry.EntityData.YangName = "bgpPeerEntry"
    bgpPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpPeerEntry.EntityData.ParentYangName = "bgpPeerTable"
    bgpPeerEntry.EntityData.SegmentPath = "bgpPeerEntry" + types.AddKeyToken(bgpPeerEntry.BgpPeerRemoteAddr, "bgpPeerRemoteAddr")
    bgpPeerEntry.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/bgpPeerTable/" + bgpPeerEntry.EntityData.SegmentPath
    bgpPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPeerEntry.EntityData.Children = types.NewOrderedMap()
    bgpPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerRemoteAddr", types.YLeaf{"BgpPeerRemoteAddr", bgpPeerEntry.BgpPeerRemoteAddr})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerIdentifier", types.YLeaf{"BgpPeerIdentifier", bgpPeerEntry.BgpPeerIdentifier})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerState", types.YLeaf{"BgpPeerState", bgpPeerEntry.BgpPeerState})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerAdminStatus", types.YLeaf{"BgpPeerAdminStatus", bgpPeerEntry.BgpPeerAdminStatus})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerNegotiatedVersion", types.YLeaf{"BgpPeerNegotiatedVersion", bgpPeerEntry.BgpPeerNegotiatedVersion})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerLocalAddr", types.YLeaf{"BgpPeerLocalAddr", bgpPeerEntry.BgpPeerLocalAddr})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerLocalPort", types.YLeaf{"BgpPeerLocalPort", bgpPeerEntry.BgpPeerLocalPort})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerRemotePort", types.YLeaf{"BgpPeerRemotePort", bgpPeerEntry.BgpPeerRemotePort})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerRemoteAs", types.YLeaf{"BgpPeerRemoteAs", bgpPeerEntry.BgpPeerRemoteAs})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerInUpdates", types.YLeaf{"BgpPeerInUpdates", bgpPeerEntry.BgpPeerInUpdates})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerOutUpdates", types.YLeaf{"BgpPeerOutUpdates", bgpPeerEntry.BgpPeerOutUpdates})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerInTotalMessages", types.YLeaf{"BgpPeerInTotalMessages", bgpPeerEntry.BgpPeerInTotalMessages})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerOutTotalMessages", types.YLeaf{"BgpPeerOutTotalMessages", bgpPeerEntry.BgpPeerOutTotalMessages})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerLastError", types.YLeaf{"BgpPeerLastError", bgpPeerEntry.BgpPeerLastError})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerFsmEstablishedTransitions", types.YLeaf{"BgpPeerFsmEstablishedTransitions", bgpPeerEntry.BgpPeerFsmEstablishedTransitions})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerFsmEstablishedTime", types.YLeaf{"BgpPeerFsmEstablishedTime", bgpPeerEntry.BgpPeerFsmEstablishedTime})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerConnectRetryInterval", types.YLeaf{"BgpPeerConnectRetryInterval", bgpPeerEntry.BgpPeerConnectRetryInterval})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerHoldTime", types.YLeaf{"BgpPeerHoldTime", bgpPeerEntry.BgpPeerHoldTime})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerKeepAlive", types.YLeaf{"BgpPeerKeepAlive", bgpPeerEntry.BgpPeerKeepAlive})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerHoldTimeConfigured", types.YLeaf{"BgpPeerHoldTimeConfigured", bgpPeerEntry.BgpPeerHoldTimeConfigured})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerKeepAliveConfigured", types.YLeaf{"BgpPeerKeepAliveConfigured", bgpPeerEntry.BgpPeerKeepAliveConfigured})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerMinASOriginationInterval", types.YLeaf{"BgpPeerMinASOriginationInterval", bgpPeerEntry.BgpPeerMinASOriginationInterval})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerMinRouteAdvertisementInterval", types.YLeaf{"BgpPeerMinRouteAdvertisementInterval", bgpPeerEntry.BgpPeerMinRouteAdvertisementInterval})
    bgpPeerEntry.EntityData.Leafs.Append("bgpPeerInUpdateElapsedTime", types.YLeaf{"BgpPeerInUpdateElapsedTime", bgpPeerEntry.BgpPeerInUpdateElapsedTime})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrefixAccepted", types.YLeaf{"CbgpPeerPrefixAccepted", bgpPeerEntry.CbgpPeerPrefixAccepted})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrefixDenied", types.YLeaf{"CbgpPeerPrefixDenied", bgpPeerEntry.CbgpPeerPrefixDenied})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrefixLimit", types.YLeaf{"CbgpPeerPrefixLimit", bgpPeerEntry.CbgpPeerPrefixLimit})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrefixAdvertised", types.YLeaf{"CbgpPeerPrefixAdvertised", bgpPeerEntry.CbgpPeerPrefixAdvertised})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrefixSuppressed", types.YLeaf{"CbgpPeerPrefixSuppressed", bgpPeerEntry.CbgpPeerPrefixSuppressed})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrefixWithdrawn", types.YLeaf{"CbgpPeerPrefixWithdrawn", bgpPeerEntry.CbgpPeerPrefixWithdrawn})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerLastErrorTxt", types.YLeaf{"CbgpPeerLastErrorTxt", bgpPeerEntry.CbgpPeerLastErrorTxt})
    bgpPeerEntry.EntityData.Leafs.Append("cbgpPeerPrevState", types.YLeaf{"CbgpPeerPrevState", bgpPeerEntry.CbgpPeerPrevState})

    bgpPeerEntry.EntityData.YListKeys = []string {"BgpPeerRemoteAddr"}

    return &(bgpPeerEntry.EntityData)
}

// BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerAdminStatus represents without adequate authentication.
type BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerAdminStatus string

const (
    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerAdminStatus_stop BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerAdminStatus = "stop"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerAdminStatus_start BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerAdminStatus = "start"
)

// BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState represents The BGP peer connection state.
type BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState string

const (
    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState_idle BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState = "idle"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState_connect BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState = "connect"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState_active BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState = "active"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState_opensent BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState = "opensent"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState_openconfirm BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState = "openconfirm"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState_established BGP4MIB_BgpPeerTable_BgpPeerEntry_BgpPeerState = "established"
)

// BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState represents The BGP peer connection previous state.
type BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState string

const (
    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_none BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "none"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_idle BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "idle"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_connect BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "connect"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_active BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "active"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_opensent BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "opensent"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_openconfirm BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "openconfirm"

    BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState_established BGP4MIB_BgpPeerTable_BgpPeerEntry_CbgpPeerPrevState = "established"
)

// BGP4MIB_BgpRcvdPathAttrTable
// The BGP Received Path Attribute Table
// contains information about paths to
// destination networks received from all
// peers running BGP version 3 or less.
type BGP4MIB_BgpRcvdPathAttrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a path to a network. The type is slice of
    // BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry.
    BgpPathAttrEntry []*BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry
}

func (bgpRcvdPathAttrTable *BGP4MIB_BgpRcvdPathAttrTable) GetEntityData() *types.CommonEntityData {
    bgpRcvdPathAttrTable.EntityData.YFilter = bgpRcvdPathAttrTable.YFilter
    bgpRcvdPathAttrTable.EntityData.YangName = "bgpRcvdPathAttrTable"
    bgpRcvdPathAttrTable.EntityData.BundleName = "cisco_ios_xe"
    bgpRcvdPathAttrTable.EntityData.ParentYangName = "BGP4-MIB"
    bgpRcvdPathAttrTable.EntityData.SegmentPath = "bgpRcvdPathAttrTable"
    bgpRcvdPathAttrTable.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/" + bgpRcvdPathAttrTable.EntityData.SegmentPath
    bgpRcvdPathAttrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRcvdPathAttrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRcvdPathAttrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRcvdPathAttrTable.EntityData.Children = types.NewOrderedMap()
    bgpRcvdPathAttrTable.EntityData.Children.Append("bgpPathAttrEntry", types.YChild{"BgpPathAttrEntry", nil})
    for i := range bgpRcvdPathAttrTable.BgpPathAttrEntry {
        bgpRcvdPathAttrTable.EntityData.Children.Append(types.GetSegmentPath(bgpRcvdPathAttrTable.BgpPathAttrEntry[i]), types.YChild{"BgpPathAttrEntry", bgpRcvdPathAttrTable.BgpPathAttrEntry[i]})
    }
    bgpRcvdPathAttrTable.EntityData.Leafs = types.NewOrderedMap()

    bgpRcvdPathAttrTable.EntityData.YListKeys = []string {}

    return &(bgpRcvdPathAttrTable.EntityData)
}

// BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry
// Information about a path to a network.
type BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address of the destination network. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpPathAttrDestNetwork interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpPathAttrPeer interface{}

    // The ultimate origin of the path information. The type is BgpPathAttrOrigin.
    BgpPathAttrOrigin interface{}

    // The set of ASs that must be traversed to reach the network.  This object is
    // probably best represented as SEQUENCE OF INTEGER.  For SMI compatibility,
    // though, it is represented as OCTET STRING.  Each AS is represented as a
    // pair of octets according to the following algorithm:     
    // first-byte-of-pair = ASNumber / 256;     second-byte-of-pair = ASNumber &
    // 255;. The type is string with length: 2..255.
    BgpPathAttrASPath interface{}

    // The address of the border router that should be used for the destination
    // network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BgpPathAttrNextHop interface{}

    // The optional inter-AS metric.  If this attribute has not been provided for
    // this route, the value for this object is 0. The type is interface{} with
    // range: -2147483648..2147483647.
    BgpPathAttrInterASMetric interface{}
}

func (bgpPathAttrEntry *BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry) GetEntityData() *types.CommonEntityData {
    bgpPathAttrEntry.EntityData.YFilter = bgpPathAttrEntry.YFilter
    bgpPathAttrEntry.EntityData.YangName = "bgpPathAttrEntry"
    bgpPathAttrEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpPathAttrEntry.EntityData.ParentYangName = "bgpRcvdPathAttrTable"
    bgpPathAttrEntry.EntityData.SegmentPath = "bgpPathAttrEntry" + types.AddKeyToken(bgpPathAttrEntry.BgpPathAttrDestNetwork, "bgpPathAttrDestNetwork") + types.AddKeyToken(bgpPathAttrEntry.BgpPathAttrPeer, "bgpPathAttrPeer")
    bgpPathAttrEntry.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/bgpRcvdPathAttrTable/" + bgpPathAttrEntry.EntityData.SegmentPath
    bgpPathAttrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPathAttrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPathAttrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPathAttrEntry.EntityData.Children = types.NewOrderedMap()
    bgpPathAttrEntry.EntityData.Leafs = types.NewOrderedMap()
    bgpPathAttrEntry.EntityData.Leafs.Append("bgpPathAttrDestNetwork", types.YLeaf{"BgpPathAttrDestNetwork", bgpPathAttrEntry.BgpPathAttrDestNetwork})
    bgpPathAttrEntry.EntityData.Leafs.Append("bgpPathAttrPeer", types.YLeaf{"BgpPathAttrPeer", bgpPathAttrEntry.BgpPathAttrPeer})
    bgpPathAttrEntry.EntityData.Leafs.Append("bgpPathAttrOrigin", types.YLeaf{"BgpPathAttrOrigin", bgpPathAttrEntry.BgpPathAttrOrigin})
    bgpPathAttrEntry.EntityData.Leafs.Append("bgpPathAttrASPath", types.YLeaf{"BgpPathAttrASPath", bgpPathAttrEntry.BgpPathAttrASPath})
    bgpPathAttrEntry.EntityData.Leafs.Append("bgpPathAttrNextHop", types.YLeaf{"BgpPathAttrNextHop", bgpPathAttrEntry.BgpPathAttrNextHop})
    bgpPathAttrEntry.EntityData.Leafs.Append("bgpPathAttrInterASMetric", types.YLeaf{"BgpPathAttrInterASMetric", bgpPathAttrEntry.BgpPathAttrInterASMetric})

    bgpPathAttrEntry.EntityData.YListKeys = []string {"BgpPathAttrDestNetwork", "BgpPathAttrPeer"}

    return &(bgpPathAttrEntry.EntityData)
}

// BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin represents The ultimate origin of the path information.
type BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin string

const (
    BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin_igp BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin = "igp"

    BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin_egp BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin = "egp"

    BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin_incomplete BGP4MIB_BgpRcvdPathAttrTable_BgpPathAttrEntry_BgpPathAttrOrigin = "incomplete"
)

// BGP4MIB_Bgp4PathAttrTable
// The BGP-4 Received Path Attribute Table
// contains information about paths to
// destination networks received from all
// BGP4 peers.
type BGP4MIB_Bgp4PathAttrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a path to a network. The type is slice of
    // BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry.
    Bgp4PathAttrEntry []*BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry
}

func (bgp4PathAttrTable *BGP4MIB_Bgp4PathAttrTable) GetEntityData() *types.CommonEntityData {
    bgp4PathAttrTable.EntityData.YFilter = bgp4PathAttrTable.YFilter
    bgp4PathAttrTable.EntityData.YangName = "bgp4PathAttrTable"
    bgp4PathAttrTable.EntityData.BundleName = "cisco_ios_xe"
    bgp4PathAttrTable.EntityData.ParentYangName = "BGP4-MIB"
    bgp4PathAttrTable.EntityData.SegmentPath = "bgp4PathAttrTable"
    bgp4PathAttrTable.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/" + bgp4PathAttrTable.EntityData.SegmentPath
    bgp4PathAttrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp4PathAttrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp4PathAttrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp4PathAttrTable.EntityData.Children = types.NewOrderedMap()
    bgp4PathAttrTable.EntityData.Children.Append("bgp4PathAttrEntry", types.YChild{"Bgp4PathAttrEntry", nil})
    for i := range bgp4PathAttrTable.Bgp4PathAttrEntry {
        bgp4PathAttrTable.EntityData.Children.Append(types.GetSegmentPath(bgp4PathAttrTable.Bgp4PathAttrEntry[i]), types.YChild{"Bgp4PathAttrEntry", bgp4PathAttrTable.Bgp4PathAttrEntry[i]})
    }
    bgp4PathAttrTable.EntityData.Leafs = types.NewOrderedMap()

    bgp4PathAttrTable.EntityData.YListKeys = []string {}

    return &(bgp4PathAttrTable.EntityData)
}

// BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry
// Information about a path to a network.
type BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An IP address prefix in the Network Layer
    // Reachability Information field.  This object is an IP address containing
    // the prefix with length specified by bgp4PathAttrIpAddrPrefixLen. Any bits
    // beyond the length specified by bgp4PathAttrIpAddrPrefixLen are zeroed. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgp4PathAttrIpAddrPrefix interface{}

    // This attribute is a key. Length in bits of the IP address prefix in the
    // Network Layer Reachability Information field. The type is interface{} with
    // range: 0..32.
    Bgp4PathAttrIpAddrPrefixLen interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgp4PathAttrPeer interface{}

    // The ultimate origin of the path information. The type is
    // Bgp4PathAttrOrigin.
    Bgp4PathAttrOrigin interface{}

    // The sequence of AS path segments.  Each AS path segment is represented by a
    // triple <type, length, value>.  The type is a 1-octet field which has two
    // possible values:      1      AS_SET: unordered set of ASs a                
    // route in the UPDATE                  message has traversed      2     
    // AS_SEQUENCE: ordered set of ASs                  a route in the UPDATE     
    // message has traversed.  The length is a 1-octet field containing the number
    // of ASs in the value field.  The value field contains one or more AS
    // numbers, each AS is represented in the octet string as a pair of octets
    // according to the following algorithm:      first-byte-of-pair = ASNumber /
    // 256;     second-byte-of-pair = ASNumber & 255;. The type is string with
    // length: 2..255.
    Bgp4PathAttrASPathSegment interface{}

    // The address of the border router that should be used for the destination
    // network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgp4PathAttrNextHop interface{}

    // This metric is used to discriminate between multiple exit points to an
    // adjacent autonomous system.  A value of -1 indicates the absence of this
    // attribute. The type is interface{} with range: -1..2147483647.
    Bgp4PathAttrMultiExitDisc interface{}

    // The originating BGP4 speaker's degree of preference for an advertised
    // route.  A value of -1 indicates the absence of this attribute. The type is
    // interface{} with range: -1..2147483647.
    Bgp4PathAttrLocalPref interface{}

    // Whether or not the local system has selected a less specific route without
    // selecting a more specific route. The type is Bgp4PathAttrAtomicAggregate.
    Bgp4PathAttrAtomicAggregate interface{}

    // The AS number of the last BGP4 speaker that performed route aggregation.  A
    // value of zero (0) indicates the absence of this attribute. The type is
    // interface{} with range: 0..65535.
    Bgp4PathAttrAggregatorAS interface{}

    // The IP address of the last BGP4 speaker that performed route aggregation. 
    // A value of 0.0.0.0 indicates the absence of this attribute. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgp4PathAttrAggregatorAddr interface{}

    // The degree of preference calculated by the receiving BGP4 speaker for an
    // advertised route.  A value of -1 indicates the absence of this attribute.
    // The type is interface{} with range: -1..2147483647.
    Bgp4PathAttrCalcLocalPref interface{}

    // An indication of whether or not this route was chosen as the best BGP4
    // route. The type is Bgp4PathAttrBest.
    Bgp4PathAttrBest interface{}

    // One or more path attributes not understood by this BGP4 speaker.  Size zero
    // (0) indicates the absence of such attribute(s).  Octets beyond the maximum
    // size, if any, are not recorded by this object. The type is string with
    // length: 0..255.
    Bgp4PathAttrUnknown interface{}
}

func (bgp4PathAttrEntry *BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry) GetEntityData() *types.CommonEntityData {
    bgp4PathAttrEntry.EntityData.YFilter = bgp4PathAttrEntry.YFilter
    bgp4PathAttrEntry.EntityData.YangName = "bgp4PathAttrEntry"
    bgp4PathAttrEntry.EntityData.BundleName = "cisco_ios_xe"
    bgp4PathAttrEntry.EntityData.ParentYangName = "bgp4PathAttrTable"
    bgp4PathAttrEntry.EntityData.SegmentPath = "bgp4PathAttrEntry" + types.AddKeyToken(bgp4PathAttrEntry.Bgp4PathAttrIpAddrPrefix, "bgp4PathAttrIpAddrPrefix") + types.AddKeyToken(bgp4PathAttrEntry.Bgp4PathAttrIpAddrPrefixLen, "bgp4PathAttrIpAddrPrefixLen") + types.AddKeyToken(bgp4PathAttrEntry.Bgp4PathAttrPeer, "bgp4PathAttrPeer")
    bgp4PathAttrEntry.EntityData.AbsolutePath = "BGP4-MIB:BGP4-MIB/bgp4PathAttrTable/" + bgp4PathAttrEntry.EntityData.SegmentPath
    bgp4PathAttrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp4PathAttrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp4PathAttrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp4PathAttrEntry.EntityData.Children = types.NewOrderedMap()
    bgp4PathAttrEntry.EntityData.Leafs = types.NewOrderedMap()
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrIpAddrPrefix", types.YLeaf{"Bgp4PathAttrIpAddrPrefix", bgp4PathAttrEntry.Bgp4PathAttrIpAddrPrefix})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrIpAddrPrefixLen", types.YLeaf{"Bgp4PathAttrIpAddrPrefixLen", bgp4PathAttrEntry.Bgp4PathAttrIpAddrPrefixLen})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrPeer", types.YLeaf{"Bgp4PathAttrPeer", bgp4PathAttrEntry.Bgp4PathAttrPeer})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrOrigin", types.YLeaf{"Bgp4PathAttrOrigin", bgp4PathAttrEntry.Bgp4PathAttrOrigin})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrASPathSegment", types.YLeaf{"Bgp4PathAttrASPathSegment", bgp4PathAttrEntry.Bgp4PathAttrASPathSegment})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrNextHop", types.YLeaf{"Bgp4PathAttrNextHop", bgp4PathAttrEntry.Bgp4PathAttrNextHop})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrMultiExitDisc", types.YLeaf{"Bgp4PathAttrMultiExitDisc", bgp4PathAttrEntry.Bgp4PathAttrMultiExitDisc})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrLocalPref", types.YLeaf{"Bgp4PathAttrLocalPref", bgp4PathAttrEntry.Bgp4PathAttrLocalPref})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrAtomicAggregate", types.YLeaf{"Bgp4PathAttrAtomicAggregate", bgp4PathAttrEntry.Bgp4PathAttrAtomicAggregate})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrAggregatorAS", types.YLeaf{"Bgp4PathAttrAggregatorAS", bgp4PathAttrEntry.Bgp4PathAttrAggregatorAS})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrAggregatorAddr", types.YLeaf{"Bgp4PathAttrAggregatorAddr", bgp4PathAttrEntry.Bgp4PathAttrAggregatorAddr})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrCalcLocalPref", types.YLeaf{"Bgp4PathAttrCalcLocalPref", bgp4PathAttrEntry.Bgp4PathAttrCalcLocalPref})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrBest", types.YLeaf{"Bgp4PathAttrBest", bgp4PathAttrEntry.Bgp4PathAttrBest})
    bgp4PathAttrEntry.EntityData.Leafs.Append("bgp4PathAttrUnknown", types.YLeaf{"Bgp4PathAttrUnknown", bgp4PathAttrEntry.Bgp4PathAttrUnknown})

    bgp4PathAttrEntry.EntityData.YListKeys = []string {"Bgp4PathAttrIpAddrPrefix", "Bgp4PathAttrIpAddrPrefixLen", "Bgp4PathAttrPeer"}

    return &(bgp4PathAttrEntry.EntityData)
}

// BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrAtomicAggregate represents selecting a more specific route.
type BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrAtomicAggregate string

const (
    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrAtomicAggregate_lessSpecificRrouteNotSelected BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrAtomicAggregate = "lessSpecificRrouteNotSelected"

    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrAtomicAggregate_lessSpecificRouteSelected BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrAtomicAggregate = "lessSpecificRouteSelected"
)

// BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrBest represents was chosen as the best BGP4 route.
type BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrBest string

const (
    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrBest_false_ BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrBest = "false"

    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrBest_true_ BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrBest = "true"
)

// BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin represents information.
type BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin string

const (
    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin_igp BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin = "igp"

    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin_egp BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin = "egp"

    BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin_incomplete BGP4MIB_Bgp4PathAttrTable_Bgp4PathAttrEntry_Bgp4PathAttrOrigin = "incomplete"
)

