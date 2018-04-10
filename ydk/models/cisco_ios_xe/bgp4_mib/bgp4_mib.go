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
    Bgppeertable BGP4MIB_Bgppeertable

    // The BGP Received Path Attribute Table contains information about paths to
    // destination networks received from all peers running BGP version 3 or less.
    Bgprcvdpathattrtable BGP4MIB_Bgprcvdpathattrtable

    // The BGP-4 Received Path Attribute Table contains information about paths to
    // destination networks received from all BGP4 peers.
    Bgp4Pathattrtable BGP4MIB_Bgp4Pathattrtable
}

func (bGP4MIB *BGP4MIB) GetEntityData() *types.CommonEntityData {
    bGP4MIB.EntityData.YFilter = bGP4MIB.YFilter
    bGP4MIB.EntityData.YangName = "BGP4-MIB"
    bGP4MIB.EntityData.BundleName = "cisco_ios_xe"
    bGP4MIB.EntityData.ParentYangName = "BGP4-MIB"
    bGP4MIB.EntityData.SegmentPath = "BGP4-MIB:BGP4-MIB"
    bGP4MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bGP4MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bGP4MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bGP4MIB.EntityData.Children = make(map[string]types.YChild)
    bGP4MIB.EntityData.Children["bgp"] = types.YChild{"Bgp", &bGP4MIB.Bgp}
    bGP4MIB.EntityData.Children["bgpPeerTable"] = types.YChild{"Bgppeertable", &bGP4MIB.Bgppeertable}
    bGP4MIB.EntityData.Children["bgpRcvdPathAttrTable"] = types.YChild{"Bgprcvdpathattrtable", &bGP4MIB.Bgprcvdpathattrtable}
    bGP4MIB.EntityData.Children["bgp4PathAttrTable"] = types.YChild{"Bgp4Pathattrtable", &bGP4MIB.Bgp4Pathattrtable}
    bGP4MIB.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Bgpversion interface{}

    // The local autonomous system number. The type is interface{} with range:
    // 0..65535.
    Bgplocalas interface{}

    // The BGP Identifier of local system. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgpidentifier interface{}
}

func (bgp *BGP4MIB_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xe"
    bgp.EntityData.ParentYangName = "BGP4-MIB"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["bgpVersion"] = types.YLeaf{"Bgpversion", bgp.Bgpversion}
    bgp.EntityData.Leafs["bgpLocalAs"] = types.YLeaf{"Bgplocalas", bgp.Bgplocalas}
    bgp.EntityData.Leafs["bgpIdentifier"] = types.YLeaf{"Bgpidentifier", bgp.Bgpidentifier}
    return &(bgp.EntityData)
}

// BGP4MIB_Bgppeertable
// BGP peer table.  This table contains,
// one entry per BGP peer, information about
// the connections with BGP peers.
type BGP4MIB_Bgppeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry containing information about the connection with a BGP peer. The type
    // is slice of BGP4MIB_Bgppeertable_Bgppeerentry.
    Bgppeerentry []BGP4MIB_Bgppeertable_Bgppeerentry
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetEntityData() *types.CommonEntityData {
    bgppeertable.EntityData.YFilter = bgppeertable.YFilter
    bgppeertable.EntityData.YangName = "bgpPeerTable"
    bgppeertable.EntityData.BundleName = "cisco_ios_xe"
    bgppeertable.EntityData.ParentYangName = "BGP4-MIB"
    bgppeertable.EntityData.SegmentPath = "bgpPeerTable"
    bgppeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgppeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgppeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgppeertable.EntityData.Children = make(map[string]types.YChild)
    bgppeertable.EntityData.Children["bgpPeerEntry"] = types.YChild{"Bgppeerentry", nil}
    for i := range bgppeertable.Bgppeerentry {
        bgppeertable.EntityData.Children[types.GetSegmentPath(&bgppeertable.Bgppeerentry[i])] = types.YChild{"Bgppeerentry", &bgppeertable.Bgppeerentry[i]}
    }
    bgppeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgppeertable.EntityData)
}

// BGP4MIB_Bgppeertable_Bgppeerentry
// Entry containing information about the
// connection with a BGP peer.
type BGP4MIB_Bgppeertable_Bgppeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The remote IP address of this entry's BGP peer.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgppeerremoteaddr interface{}

    // The BGP Identifier of this entry's BGP peer. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgppeeridentifier interface{}

    // The BGP peer connection state. The type is Bgppeerstate.
    Bgppeerstate interface{}

    // The desired state of the BGP connection. A transition from 'stop' to
    // 'start' will cause the BGP Start Event to be generated. A transition from
    // 'start' to 'stop' will cause the BGP Stop Event to be generated. This
    // parameter can be used to restart BGP peer connections.  Care should be used
    // in providing write access to this object without adequate authentication.
    // The type is Bgppeeradminstatus.
    Bgppeeradminstatus interface{}

    // The negotiated version of BGP running between the two peers. The type is
    // interface{} with range: -2147483648..2147483647.
    Bgppeernegotiatedversion interface{}

    // The local IP address of this entry's BGP connection. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgppeerlocaladdr interface{}

    // The local port for the TCP connection between the BGP peers. The type is
    // interface{} with range: 0..65535.
    Bgppeerlocalport interface{}

    // The remote port for the TCP connection between the BGP peers.  Note that
    // the objects bgpPeerLocalAddr, bgpPeerLocalPort, bgpPeerRemoteAddr and
    // bgpPeerRemotePort provide the appropriate reference to the standard MIB TCP
    // connection table. The type is interface{} with range: 0..65535.
    Bgppeerremoteport interface{}

    // The remote autonomous system number. The type is interface{} with range:
    // 0..65535.
    Bgppeerremoteas interface{}

    // The number of BGP UPDATE messages received on this connection.  This object
    // should be initialized to zero (0) when the connection is established. The
    // type is interface{} with range: 0..4294967295.
    Bgppeerinupdates interface{}

    // The number of BGP UPDATE messages transmitted on this connection.  This
    // object should be initialized to zero (0) when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    Bgppeeroutupdates interface{}

    // The total number of messages received from the remote peer on this
    // connection. This object should be initialized to zero when the connection
    // is established. The type is interface{} with range: 0..4294967295.
    Bgppeerintotalmessages interface{}

    // The total number of messages transmitted to the remote peer on this
    // connection.  This object should be initialized to zero when the connection
    // is established. The type is interface{} with range: 0..4294967295.
    Bgppeerouttotalmessages interface{}

    // The last error code and subcode seen by this peer on this connection.  If
    // no error has occurred, this field is zero.  Otherwise, the first byte of
    // this two byte OCTET STRING contains the error code, and the second byte
    // contains the subcode. The type is string with length: 2.
    Bgppeerlasterror interface{}

    // The total number of times the BGP FSM transitioned into the established
    // state. The type is interface{} with range: 0..4294967295.
    Bgppeerfsmestablishedtransitions interface{}

    // This timer indicates how long (in seconds) this peer has been in the
    // Established state or how long since this peer was last in the Established
    // state.  It is set to zero when a new peer is configured or the router is
    // booted. The type is interface{} with range: 0..4294967295.
    Bgppeerfsmestablishedtime interface{}

    // Time interval in seconds for the ConnectRetry timer.  The suggested value
    // for this timer is 120 seconds. The type is interface{} with range:
    // 1..65535.
    Bgppeerconnectretryinterval interface{}

    // Time interval in seconds for the Hold Timer established with the peer.  The
    // value of this object is calculated by this BGP speaker by using the smaller
    // of the value in bgpPeerHoldTimeConfigured and the Hold Time received in the
    // OPEN message. This value must be at lease three seconds if it is not zero
    // (0) in which case the Hold Timer has not been established with the peer,
    // or, the value of bgpPeerHoldTimeConfigured is zero (0). The type is
    // interface{} with range: 0..None | 3..65535.
    Bgppeerholdtime interface{}

    // Time interval in seconds for the KeepAlive timer established with the peer.
    // The value of this object is calculated by this BGP speaker such that, when
    // compared with bgpPeerHoldTime, it has the same proportion as what
    // bgpPeerKeepAliveConfigured has when compared with
    // bgpPeerHoldTimeConfigured. If the value of this object is zero (0), it
    // indicates that the KeepAlive timer has not been established with the peer,
    // or, the value of bgpPeerKeepAliveConfigured is zero (0). The type is
    // interface{} with range: 0..21845.
    Bgppeerkeepalive interface{}

    // Time interval in seconds for the Hold Time configured for this BGP speaker
    // with this peer.  This value is placed in an OPEN message sent to this peer
    // by this BGP speaker, and is compared with the Hold Time field in an OPEN
    // message received from the peer when determining the Hold Time
    // (bgpPeerHoldTime) with the peer. This value must not be less than three
    // seconds if it is not zero (0) in which case the Hold Time is NOT to be
    // established with the peer.  The suggested value for this timer is 90
    // seconds. The type is interface{} with range: 0..None | 3..65535.
    Bgppeerholdtimeconfigured interface{}

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
    Bgppeerkeepaliveconfigured interface{}

    // Time interval in seconds for the MinASOriginationInterval timer. The
    // suggested value for this timer is 15 seconds. The type is interface{} with
    // range: 1..65535.
    Bgppeerminasoriginationinterval interface{}

    // Time interval in seconds for the MinRouteAdvertisementInterval timer. The
    // suggested value for this timer is 30 seconds. The type is interface{} with
    // range: 1..65535.
    Bgppeerminrouteadvertisementinterval interface{}

    // Elapsed time in seconds since the last BGP UPDATE message was received from
    // the peer. Each time bgpPeerInUpdates is incremented, the value of this
    // object is set to zero (0). The type is interface{} with range:
    // 0..4294967295.
    Bgppeerinupdateelapsedtime interface{}

    // Number of Route prefixes received on this connnection, which are accepted
    // after applying filters. Possible filters are route maps, prefix lists,
    // distributed lists, etc. The type is interface{} with range: 0..4294967295.
    Cbgppeerprefixaccepted interface{}

    // Counter which gets incremented when a route prefix received on this
    // connection is denied  or when a route prefix is denied during soft reset of
    // this connection if 'soft-reconfiguration' is on . This object is 
    // initialized to zero when the peer is  configured or the router is rebooted.
    // The type is interface{} with range: 0..4294967295.
    Cbgppeerprefixdenied interface{}

    // Max number of route prefixes accepted on this connection. The type is
    // interface{} with range: 1..4294967295.
    Cbgppeerprefixlimit interface{}

    // Counter which gets incremented when a route prefix is advertised on this
    // connection. This object is initialized to zero when the peer is configured
    // or  the router is rebooted. The type is interface{} with range:
    // 0..4294967295.
    Cbgppeerprefixadvertised interface{}

    // Counter which gets incremented when a route prefix is suppressed from being
    // sent on this connection. This  object is initialized to zero when the peer
    // is  configured or the router is rebooted. The type is interface{} with
    // range: 0..4294967295.
    Cbgppeerprefixsuppressed interface{}

    // Counter which gets incremented when a route prefix is withdrawn on this
    // connection. This object is initialized to zero when the peer is configured
    // or the router is rebooted. The type is interface{} with range:
    // 0..4294967295.
    Cbgppeerprefixwithdrawn interface{}

    // Implementation specific error description for bgpPeerLastErrorReceived. The
    // type is string.
    Cbgppeerlasterrortxt interface{}

    // The BGP peer connection previous state. The type is Cbgppeerprevstate.
    Cbgppeerprevstate interface{}
}

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetEntityData() *types.CommonEntityData {
    bgppeerentry.EntityData.YFilter = bgppeerentry.YFilter
    bgppeerentry.EntityData.YangName = "bgpPeerEntry"
    bgppeerentry.EntityData.BundleName = "cisco_ios_xe"
    bgppeerentry.EntityData.ParentYangName = "bgpPeerTable"
    bgppeerentry.EntityData.SegmentPath = "bgpPeerEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", bgppeerentry.Bgppeerremoteaddr) + "']"
    bgppeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgppeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgppeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgppeerentry.EntityData.Children = make(map[string]types.YChild)
    bgppeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgppeerentry.EntityData.Leafs["bgpPeerRemoteAddr"] = types.YLeaf{"Bgppeerremoteaddr", bgppeerentry.Bgppeerremoteaddr}
    bgppeerentry.EntityData.Leafs["bgpPeerIdentifier"] = types.YLeaf{"Bgppeeridentifier", bgppeerentry.Bgppeeridentifier}
    bgppeerentry.EntityData.Leafs["bgpPeerState"] = types.YLeaf{"Bgppeerstate", bgppeerentry.Bgppeerstate}
    bgppeerentry.EntityData.Leafs["bgpPeerAdminStatus"] = types.YLeaf{"Bgppeeradminstatus", bgppeerentry.Bgppeeradminstatus}
    bgppeerentry.EntityData.Leafs["bgpPeerNegotiatedVersion"] = types.YLeaf{"Bgppeernegotiatedversion", bgppeerentry.Bgppeernegotiatedversion}
    bgppeerentry.EntityData.Leafs["bgpPeerLocalAddr"] = types.YLeaf{"Bgppeerlocaladdr", bgppeerentry.Bgppeerlocaladdr}
    bgppeerentry.EntityData.Leafs["bgpPeerLocalPort"] = types.YLeaf{"Bgppeerlocalport", bgppeerentry.Bgppeerlocalport}
    bgppeerentry.EntityData.Leafs["bgpPeerRemotePort"] = types.YLeaf{"Bgppeerremoteport", bgppeerentry.Bgppeerremoteport}
    bgppeerentry.EntityData.Leafs["bgpPeerRemoteAs"] = types.YLeaf{"Bgppeerremoteas", bgppeerentry.Bgppeerremoteas}
    bgppeerentry.EntityData.Leafs["bgpPeerInUpdates"] = types.YLeaf{"Bgppeerinupdates", bgppeerentry.Bgppeerinupdates}
    bgppeerentry.EntityData.Leafs["bgpPeerOutUpdates"] = types.YLeaf{"Bgppeeroutupdates", bgppeerentry.Bgppeeroutupdates}
    bgppeerentry.EntityData.Leafs["bgpPeerInTotalMessages"] = types.YLeaf{"Bgppeerintotalmessages", bgppeerentry.Bgppeerintotalmessages}
    bgppeerentry.EntityData.Leafs["bgpPeerOutTotalMessages"] = types.YLeaf{"Bgppeerouttotalmessages", bgppeerentry.Bgppeerouttotalmessages}
    bgppeerentry.EntityData.Leafs["bgpPeerLastError"] = types.YLeaf{"Bgppeerlasterror", bgppeerentry.Bgppeerlasterror}
    bgppeerentry.EntityData.Leafs["bgpPeerFsmEstablishedTransitions"] = types.YLeaf{"Bgppeerfsmestablishedtransitions", bgppeerentry.Bgppeerfsmestablishedtransitions}
    bgppeerentry.EntityData.Leafs["bgpPeerFsmEstablishedTime"] = types.YLeaf{"Bgppeerfsmestablishedtime", bgppeerentry.Bgppeerfsmestablishedtime}
    bgppeerentry.EntityData.Leafs["bgpPeerConnectRetryInterval"] = types.YLeaf{"Bgppeerconnectretryinterval", bgppeerentry.Bgppeerconnectretryinterval}
    bgppeerentry.EntityData.Leafs["bgpPeerHoldTime"] = types.YLeaf{"Bgppeerholdtime", bgppeerentry.Bgppeerholdtime}
    bgppeerentry.EntityData.Leafs["bgpPeerKeepAlive"] = types.YLeaf{"Bgppeerkeepalive", bgppeerentry.Bgppeerkeepalive}
    bgppeerentry.EntityData.Leafs["bgpPeerHoldTimeConfigured"] = types.YLeaf{"Bgppeerholdtimeconfigured", bgppeerentry.Bgppeerholdtimeconfigured}
    bgppeerentry.EntityData.Leafs["bgpPeerKeepAliveConfigured"] = types.YLeaf{"Bgppeerkeepaliveconfigured", bgppeerentry.Bgppeerkeepaliveconfigured}
    bgppeerentry.EntityData.Leafs["bgpPeerMinASOriginationInterval"] = types.YLeaf{"Bgppeerminasoriginationinterval", bgppeerentry.Bgppeerminasoriginationinterval}
    bgppeerentry.EntityData.Leafs["bgpPeerMinRouteAdvertisementInterval"] = types.YLeaf{"Bgppeerminrouteadvertisementinterval", bgppeerentry.Bgppeerminrouteadvertisementinterval}
    bgppeerentry.EntityData.Leafs["bgpPeerInUpdateElapsedTime"] = types.YLeaf{"Bgppeerinupdateelapsedtime", bgppeerentry.Bgppeerinupdateelapsedtime}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrefixAccepted"] = types.YLeaf{"Cbgppeerprefixaccepted", bgppeerentry.Cbgppeerprefixaccepted}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrefixDenied"] = types.YLeaf{"Cbgppeerprefixdenied", bgppeerentry.Cbgppeerprefixdenied}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrefixLimit"] = types.YLeaf{"Cbgppeerprefixlimit", bgppeerentry.Cbgppeerprefixlimit}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrefixAdvertised"] = types.YLeaf{"Cbgppeerprefixadvertised", bgppeerentry.Cbgppeerprefixadvertised}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrefixSuppressed"] = types.YLeaf{"Cbgppeerprefixsuppressed", bgppeerentry.Cbgppeerprefixsuppressed}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrefixWithdrawn"] = types.YLeaf{"Cbgppeerprefixwithdrawn", bgppeerentry.Cbgppeerprefixwithdrawn}
    bgppeerentry.EntityData.Leafs["cbgpPeerLastErrorTxt"] = types.YLeaf{"Cbgppeerlasterrortxt", bgppeerentry.Cbgppeerlasterrortxt}
    bgppeerentry.EntityData.Leafs["cbgpPeerPrevState"] = types.YLeaf{"Cbgppeerprevstate", bgppeerentry.Cbgppeerprevstate}
    return &(bgppeerentry.EntityData)
}

// BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeeradminstatus represents without adequate authentication.
type BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeeradminstatus string

const (
    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeeradminstatus_stop BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeeradminstatus = "stop"

    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeeradminstatus_start BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeeradminstatus = "start"
)

// BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate represents The BGP peer connection state.
type BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate string

const (
    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate_idle BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate = "idle"

    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate_connect BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate = "connect"

    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate_active BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate = "active"

    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate_opensent BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate = "opensent"

    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate_openconfirm BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate = "openconfirm"

    BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate_established BGP4MIB_Bgppeertable_Bgppeerentry_Bgppeerstate = "established"
)

// BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate represents The BGP peer connection previous state.
type BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate string

const (
    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_none BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "none"

    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_idle BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "idle"

    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_connect BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "connect"

    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_active BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "active"

    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_opensent BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "opensent"

    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_openconfirm BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "openconfirm"

    BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate_established BGP4MIB_Bgppeertable_Bgppeerentry_Cbgppeerprevstate = "established"
)

// BGP4MIB_Bgprcvdpathattrtable
// The BGP Received Path Attribute Table
// contains information about paths to
// destination networks received from all
// peers running BGP version 3 or less.
type BGP4MIB_Bgprcvdpathattrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a path to a network. The type is slice of
    // BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry.
    Bgppathattrentry []BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetEntityData() *types.CommonEntityData {
    bgprcvdpathattrtable.EntityData.YFilter = bgprcvdpathattrtable.YFilter
    bgprcvdpathattrtable.EntityData.YangName = "bgpRcvdPathAttrTable"
    bgprcvdpathattrtable.EntityData.BundleName = "cisco_ios_xe"
    bgprcvdpathattrtable.EntityData.ParentYangName = "BGP4-MIB"
    bgprcvdpathattrtable.EntityData.SegmentPath = "bgpRcvdPathAttrTable"
    bgprcvdpathattrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgprcvdpathattrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgprcvdpathattrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgprcvdpathattrtable.EntityData.Children = make(map[string]types.YChild)
    bgprcvdpathattrtable.EntityData.Children["bgpPathAttrEntry"] = types.YChild{"Bgppathattrentry", nil}
    for i := range bgprcvdpathattrtable.Bgppathattrentry {
        bgprcvdpathattrtable.EntityData.Children[types.GetSegmentPath(&bgprcvdpathattrtable.Bgppathattrentry[i])] = types.YChild{"Bgppathattrentry", &bgprcvdpathattrtable.Bgppathattrentry[i]}
    }
    bgprcvdpathattrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgprcvdpathattrtable.EntityData)
}

// BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry
// Information about a path to a network.
type BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The address of the destination network. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgppathattrdestnetwork interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgppathattrpeer interface{}

    // The ultimate origin of the path information. The type is Bgppathattrorigin.
    Bgppathattrorigin interface{}

    // The set of ASs that must be traversed to reach the network.  This object is
    // probably best represented as SEQUENCE OF INTEGER.  For SMI compatibility,
    // though, it is represented as OCTET STRING.  Each AS is represented as a
    // pair of octets according to the following algorithm:     
    // first-byte-of-pair = ASNumber / 256;     second-byte-of-pair = ASNumber &
    // 255;. The type is string with length: 2..255.
    Bgppathattraspath interface{}

    // The address of the border router that should be used for the destination
    // network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgppathattrnexthop interface{}

    // The optional inter-AS metric.  If this attribute has not been provided for
    // this route, the value for this object is 0. The type is interface{} with
    // range: -2147483648..2147483647.
    Bgppathattrinterasmetric interface{}
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetEntityData() *types.CommonEntityData {
    bgppathattrentry.EntityData.YFilter = bgppathattrentry.YFilter
    bgppathattrentry.EntityData.YangName = "bgpPathAttrEntry"
    bgppathattrentry.EntityData.BundleName = "cisco_ios_xe"
    bgppathattrentry.EntityData.ParentYangName = "bgpRcvdPathAttrTable"
    bgppathattrentry.EntityData.SegmentPath = "bgpPathAttrEntry" + "[bgpPathAttrDestNetwork='" + fmt.Sprintf("%v", bgppathattrentry.Bgppathattrdestnetwork) + "']" + "[bgpPathAttrPeer='" + fmt.Sprintf("%v", bgppathattrentry.Bgppathattrpeer) + "']"
    bgppathattrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgppathattrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgppathattrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgppathattrentry.EntityData.Children = make(map[string]types.YChild)
    bgppathattrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgppathattrentry.EntityData.Leafs["bgpPathAttrDestNetwork"] = types.YLeaf{"Bgppathattrdestnetwork", bgppathattrentry.Bgppathattrdestnetwork}
    bgppathattrentry.EntityData.Leafs["bgpPathAttrPeer"] = types.YLeaf{"Bgppathattrpeer", bgppathattrentry.Bgppathattrpeer}
    bgppathattrentry.EntityData.Leafs["bgpPathAttrOrigin"] = types.YLeaf{"Bgppathattrorigin", bgppathattrentry.Bgppathattrorigin}
    bgppathattrentry.EntityData.Leafs["bgpPathAttrASPath"] = types.YLeaf{"Bgppathattraspath", bgppathattrentry.Bgppathattraspath}
    bgppathattrentry.EntityData.Leafs["bgpPathAttrNextHop"] = types.YLeaf{"Bgppathattrnexthop", bgppathattrentry.Bgppathattrnexthop}
    bgppathattrentry.EntityData.Leafs["bgpPathAttrInterASMetric"] = types.YLeaf{"Bgppathattrinterasmetric", bgppathattrentry.Bgppathattrinterasmetric}
    return &(bgppathattrentry.EntityData)
}

// BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin represents The ultimate origin of the path information.
type BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin string

const (
    BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin_igp BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin = "igp"

    BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin_egp BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin = "egp"

    BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin_incomplete BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry_Bgppathattrorigin = "incomplete"
)

// BGP4MIB_Bgp4Pathattrtable
// The BGP-4 Received Path Attribute Table
// contains information about paths to
// destination networks received from all
// BGP4 peers.
type BGP4MIB_Bgp4Pathattrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a path to a network. The type is slice of
    // BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry.
    Bgp4Pathattrentry []BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetEntityData() *types.CommonEntityData {
    bgp4Pathattrtable.EntityData.YFilter = bgp4Pathattrtable.YFilter
    bgp4Pathattrtable.EntityData.YangName = "bgp4PathAttrTable"
    bgp4Pathattrtable.EntityData.BundleName = "cisco_ios_xe"
    bgp4Pathattrtable.EntityData.ParentYangName = "BGP4-MIB"
    bgp4Pathattrtable.EntityData.SegmentPath = "bgp4PathAttrTable"
    bgp4Pathattrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp4Pathattrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp4Pathattrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp4Pathattrtable.EntityData.Children = make(map[string]types.YChild)
    bgp4Pathattrtable.EntityData.Children["bgp4PathAttrEntry"] = types.YChild{"Bgp4Pathattrentry", nil}
    for i := range bgp4Pathattrtable.Bgp4Pathattrentry {
        bgp4Pathattrtable.EntityData.Children[types.GetSegmentPath(&bgp4Pathattrtable.Bgp4Pathattrentry[i])] = types.YChild{"Bgp4Pathattrentry", &bgp4Pathattrtable.Bgp4Pathattrentry[i]}
    }
    bgp4Pathattrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgp4Pathattrtable.EntityData)
}

// BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry
// Information about a path to a network.
type BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An IP address prefix in the Network Layer
    // Reachability Information field.  This object is an IP address containing
    // the prefix with length specified by bgp4PathAttrIpAddrPrefixLen. Any bits
    // beyond the length specified by bgp4PathAttrIpAddrPrefixLen are zeroed. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgp4Pathattripaddrprefix interface{}

    // This attribute is a key. Length in bits of the IP address prefix in the
    // Network Layer Reachability Information field. The type is interface{} with
    // range: 0..32.
    Bgp4Pathattripaddrprefixlen interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgp4Pathattrpeer interface{}

    // The ultimate origin of the path information. The type is
    // Bgp4Pathattrorigin.
    Bgp4Pathattrorigin interface{}

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
    Bgp4Pathattraspathsegment interface{}

    // The address of the border router that should be used for the destination
    // network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgp4Pathattrnexthop interface{}

    // This metric is used to discriminate between multiple exit points to an
    // adjacent autonomous system.  A value of -1 indicates the absence of this
    // attribute. The type is interface{} with range: -1..2147483647.
    Bgp4Pathattrmultiexitdisc interface{}

    // The originating BGP4 speaker's degree of preference for an advertised
    // route.  A value of -1 indicates the absence of this attribute. The type is
    // interface{} with range: -1..2147483647.
    Bgp4Pathattrlocalpref interface{}

    // Whether or not the local system has selected a less specific route without
    // selecting a more specific route. The type is Bgp4Pathattratomicaggregate.
    Bgp4Pathattratomicaggregate interface{}

    // The AS number of the last BGP4 speaker that performed route aggregation.  A
    // value of zero (0) indicates the absence of this attribute. The type is
    // interface{} with range: 0..65535.
    Bgp4Pathattraggregatoras interface{}

    // The IP address of the last BGP4 speaker that performed route aggregation. 
    // A value of 0.0.0.0 indicates the absence of this attribute. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bgp4Pathattraggregatoraddr interface{}

    // The degree of preference calculated by the receiving BGP4 speaker for an
    // advertised route.  A value of -1 indicates the absence of this attribute.
    // The type is interface{} with range: -1..2147483647.
    Bgp4Pathattrcalclocalpref interface{}

    // An indication of whether or not this route was chosen as the best BGP4
    // route. The type is Bgp4Pathattrbest.
    Bgp4Pathattrbest interface{}

    // One or more path attributes not understood by this BGP4 speaker.  Size zero
    // (0) indicates the absence of such attribute(s).  Octets beyond the maximum
    // size, if any, are not recorded by this object. The type is string with
    // length: 0..255.
    Bgp4Pathattrunknown interface{}
}

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetEntityData() *types.CommonEntityData {
    bgp4Pathattrentry.EntityData.YFilter = bgp4Pathattrentry.YFilter
    bgp4Pathattrentry.EntityData.YangName = "bgp4PathAttrEntry"
    bgp4Pathattrentry.EntityData.BundleName = "cisco_ios_xe"
    bgp4Pathattrentry.EntityData.ParentYangName = "bgp4PathAttrTable"
    bgp4Pathattrentry.EntityData.SegmentPath = "bgp4PathAttrEntry" + "[bgp4PathAttrIpAddrPrefix='" + fmt.Sprintf("%v", bgp4Pathattrentry.Bgp4Pathattripaddrprefix) + "']" + "[bgp4PathAttrIpAddrPrefixLen='" + fmt.Sprintf("%v", bgp4Pathattrentry.Bgp4Pathattripaddrprefixlen) + "']" + "[bgp4PathAttrPeer='" + fmt.Sprintf("%v", bgp4Pathattrentry.Bgp4Pathattrpeer) + "']"
    bgp4Pathattrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp4Pathattrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp4Pathattrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp4Pathattrentry.EntityData.Children = make(map[string]types.YChild)
    bgp4Pathattrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrIpAddrPrefix"] = types.YLeaf{"Bgp4Pathattripaddrprefix", bgp4Pathattrentry.Bgp4Pathattripaddrprefix}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrIpAddrPrefixLen"] = types.YLeaf{"Bgp4Pathattripaddrprefixlen", bgp4Pathattrentry.Bgp4Pathattripaddrprefixlen}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrPeer"] = types.YLeaf{"Bgp4Pathattrpeer", bgp4Pathattrentry.Bgp4Pathattrpeer}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrOrigin"] = types.YLeaf{"Bgp4Pathattrorigin", bgp4Pathattrentry.Bgp4Pathattrorigin}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrASPathSegment"] = types.YLeaf{"Bgp4Pathattraspathsegment", bgp4Pathattrentry.Bgp4Pathattraspathsegment}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrNextHop"] = types.YLeaf{"Bgp4Pathattrnexthop", bgp4Pathattrentry.Bgp4Pathattrnexthop}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrMultiExitDisc"] = types.YLeaf{"Bgp4Pathattrmultiexitdisc", bgp4Pathattrentry.Bgp4Pathattrmultiexitdisc}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrLocalPref"] = types.YLeaf{"Bgp4Pathattrlocalpref", bgp4Pathattrentry.Bgp4Pathattrlocalpref}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrAtomicAggregate"] = types.YLeaf{"Bgp4Pathattratomicaggregate", bgp4Pathattrentry.Bgp4Pathattratomicaggregate}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrAggregatorAS"] = types.YLeaf{"Bgp4Pathattraggregatoras", bgp4Pathattrentry.Bgp4Pathattraggregatoras}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrAggregatorAddr"] = types.YLeaf{"Bgp4Pathattraggregatoraddr", bgp4Pathattrentry.Bgp4Pathattraggregatoraddr}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrCalcLocalPref"] = types.YLeaf{"Bgp4Pathattrcalclocalpref", bgp4Pathattrentry.Bgp4Pathattrcalclocalpref}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrBest"] = types.YLeaf{"Bgp4Pathattrbest", bgp4Pathattrentry.Bgp4Pathattrbest}
    bgp4Pathattrentry.EntityData.Leafs["bgp4PathAttrUnknown"] = types.YLeaf{"Bgp4Pathattrunknown", bgp4Pathattrentry.Bgp4Pathattrunknown}
    return &(bgp4Pathattrentry.EntityData)
}

// BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattratomicaggregate represents selecting a more specific route.
type BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattratomicaggregate string

const (
    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattratomicaggregate_lessSpecificRrouteNotSelected BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattratomicaggregate = "lessSpecificRrouteNotSelected"

    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattratomicaggregate_lessSpecificRouteSelected BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattratomicaggregate = "lessSpecificRouteSelected"
)

// BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrbest represents was chosen as the best BGP4 route.
type BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrbest string

const (
    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrbest_false BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrbest = "false"

    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrbest_true BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrbest = "true"
)

// BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin represents information.
type BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin string

const (
    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin_igp BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin = "igp"

    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin_egp BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin = "egp"

    BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin_incomplete BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry_Bgp4Pathattrorigin = "incomplete"
)

