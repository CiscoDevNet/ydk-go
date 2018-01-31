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
    parent types.Entity
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

func (bGP4MIB *BGP4MIB) GetFilter() yfilter.YFilter { return bGP4MIB.YFilter }

func (bGP4MIB *BGP4MIB) SetFilter(yf yfilter.YFilter) { bGP4MIB.YFilter = yf }

func (bGP4MIB *BGP4MIB) GetGoName(yname string) string {
    if yname == "bgp" { return "Bgp" }
    if yname == "bgpPeerTable" { return "Bgppeertable" }
    if yname == "bgpRcvdPathAttrTable" { return "Bgprcvdpathattrtable" }
    if yname == "bgp4PathAttrTable" { return "Bgp4Pathattrtable" }
    return ""
}

func (bGP4MIB *BGP4MIB) GetSegmentPath() string {
    return "BGP4-MIB:BGP4-MIB"
}

func (bGP4MIB *BGP4MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp" {
        return &bGP4MIB.Bgp
    }
    if childYangName == "bgpPeerTable" {
        return &bGP4MIB.Bgppeertable
    }
    if childYangName == "bgpRcvdPathAttrTable" {
        return &bGP4MIB.Bgprcvdpathattrtable
    }
    if childYangName == "bgp4PathAttrTable" {
        return &bGP4MIB.Bgp4Pathattrtable
    }
    return nil
}

func (bGP4MIB *BGP4MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp"] = &bGP4MIB.Bgp
    children["bgpPeerTable"] = &bGP4MIB.Bgppeertable
    children["bgpRcvdPathAttrTable"] = &bGP4MIB.Bgprcvdpathattrtable
    children["bgp4PathAttrTable"] = &bGP4MIB.Bgp4Pathattrtable
    return children
}

func (bGP4MIB *BGP4MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bGP4MIB *BGP4MIB) GetBundleName() string { return "cisco_ios_xe" }

func (bGP4MIB *BGP4MIB) GetYangName() string { return "BGP4-MIB" }

func (bGP4MIB *BGP4MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bGP4MIB *BGP4MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bGP4MIB *BGP4MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bGP4MIB *BGP4MIB) SetParent(parent types.Entity) { bGP4MIB.parent = parent }

func (bGP4MIB *BGP4MIB) GetParent() types.Entity { return bGP4MIB.parent }

func (bGP4MIB *BGP4MIB) GetParentYangName() string { return "BGP4-MIB" }

// BGP4MIB_Bgp
type BGP4MIB_Bgp struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgpidentifier interface{}
}

func (bgp *BGP4MIB_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *BGP4MIB_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *BGP4MIB_Bgp) GetGoName(yname string) string {
    if yname == "bgpVersion" { return "Bgpversion" }
    if yname == "bgpLocalAs" { return "Bgplocalas" }
    if yname == "bgpIdentifier" { return "Bgpidentifier" }
    return ""
}

func (bgp *BGP4MIB_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *BGP4MIB_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *BGP4MIB_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *BGP4MIB_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgpVersion"] = bgp.Bgpversion
    leafs["bgpLocalAs"] = bgp.Bgplocalas
    leafs["bgpIdentifier"] = bgp.Bgpidentifier
    return leafs
}

func (bgp *BGP4MIB_Bgp) GetBundleName() string { return "cisco_ios_xe" }

func (bgp *BGP4MIB_Bgp) GetYangName() string { return "bgp" }

func (bgp *BGP4MIB_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgp *BGP4MIB_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgp *BGP4MIB_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgp *BGP4MIB_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *BGP4MIB_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *BGP4MIB_Bgp) GetParentYangName() string { return "BGP4-MIB" }

// BGP4MIB_Bgppeertable
// BGP peer table.  This table contains,
// one entry per BGP peer, information about
// the connections with BGP peers.
type BGP4MIB_Bgppeertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entry containing information about the connection with a BGP peer. The type
    // is slice of BGP4MIB_Bgppeertable_Bgppeerentry.
    Bgppeerentry []BGP4MIB_Bgppeertable_Bgppeerentry
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetFilter() yfilter.YFilter { return bgppeertable.YFilter }

func (bgppeertable *BGP4MIB_Bgppeertable) SetFilter(yf yfilter.YFilter) { bgppeertable.YFilter = yf }

func (bgppeertable *BGP4MIB_Bgppeertable) GetGoName(yname string) string {
    if yname == "bgpPeerEntry" { return "Bgppeerentry" }
    return ""
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetSegmentPath() string {
    return "bgpPeerTable"
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgpPeerEntry" {
        for _, c := range bgppeertable.Bgppeerentry {
            if bgppeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BGP4MIB_Bgppeertable_Bgppeerentry{}
        bgppeertable.Bgppeerentry = append(bgppeertable.Bgppeerentry, child)
        return &bgppeertable.Bgppeerentry[len(bgppeertable.Bgppeerentry)-1]
    }
    return nil
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgppeertable.Bgppeerentry {
        children[bgppeertable.Bgppeerentry[i].GetSegmentPath()] = &bgppeertable.Bgppeerentry[i]
    }
    return children
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgppeertable *BGP4MIB_Bgppeertable) GetBundleName() string { return "cisco_ios_xe" }

func (bgppeertable *BGP4MIB_Bgppeertable) GetYangName() string { return "bgpPeerTable" }

func (bgppeertable *BGP4MIB_Bgppeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgppeertable *BGP4MIB_Bgppeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgppeertable *BGP4MIB_Bgppeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgppeertable *BGP4MIB_Bgppeertable) SetParent(parent types.Entity) { bgppeertable.parent = parent }

func (bgppeertable *BGP4MIB_Bgppeertable) GetParent() types.Entity { return bgppeertable.parent }

func (bgppeertable *BGP4MIB_Bgppeertable) GetParentYangName() string { return "BGP4-MIB" }

// BGP4MIB_Bgppeertable_Bgppeerentry
// Entry containing information about the
// connection with a BGP peer.
type BGP4MIB_Bgppeertable_Bgppeerentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The remote IP address of this entry's BGP peer.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgppeerremoteaddr interface{}

    // The BGP Identifier of this entry's BGP peer. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetFilter() yfilter.YFilter { return bgppeerentry.YFilter }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) SetFilter(yf yfilter.YFilter) { bgppeerentry.YFilter = yf }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetGoName(yname string) string {
    if yname == "bgpPeerRemoteAddr" { return "Bgppeerremoteaddr" }
    if yname == "bgpPeerIdentifier" { return "Bgppeeridentifier" }
    if yname == "bgpPeerState" { return "Bgppeerstate" }
    if yname == "bgpPeerAdminStatus" { return "Bgppeeradminstatus" }
    if yname == "bgpPeerNegotiatedVersion" { return "Bgppeernegotiatedversion" }
    if yname == "bgpPeerLocalAddr" { return "Bgppeerlocaladdr" }
    if yname == "bgpPeerLocalPort" { return "Bgppeerlocalport" }
    if yname == "bgpPeerRemotePort" { return "Bgppeerremoteport" }
    if yname == "bgpPeerRemoteAs" { return "Bgppeerremoteas" }
    if yname == "bgpPeerInUpdates" { return "Bgppeerinupdates" }
    if yname == "bgpPeerOutUpdates" { return "Bgppeeroutupdates" }
    if yname == "bgpPeerInTotalMessages" { return "Bgppeerintotalmessages" }
    if yname == "bgpPeerOutTotalMessages" { return "Bgppeerouttotalmessages" }
    if yname == "bgpPeerLastError" { return "Bgppeerlasterror" }
    if yname == "bgpPeerFsmEstablishedTransitions" { return "Bgppeerfsmestablishedtransitions" }
    if yname == "bgpPeerFsmEstablishedTime" { return "Bgppeerfsmestablishedtime" }
    if yname == "bgpPeerConnectRetryInterval" { return "Bgppeerconnectretryinterval" }
    if yname == "bgpPeerHoldTime" { return "Bgppeerholdtime" }
    if yname == "bgpPeerKeepAlive" { return "Bgppeerkeepalive" }
    if yname == "bgpPeerHoldTimeConfigured" { return "Bgppeerholdtimeconfigured" }
    if yname == "bgpPeerKeepAliveConfigured" { return "Bgppeerkeepaliveconfigured" }
    if yname == "bgpPeerMinASOriginationInterval" { return "Bgppeerminasoriginationinterval" }
    if yname == "bgpPeerMinRouteAdvertisementInterval" { return "Bgppeerminrouteadvertisementinterval" }
    if yname == "bgpPeerInUpdateElapsedTime" { return "Bgppeerinupdateelapsedtime" }
    if yname == "cbgpPeerPrefixAccepted" { return "Cbgppeerprefixaccepted" }
    if yname == "cbgpPeerPrefixDenied" { return "Cbgppeerprefixdenied" }
    if yname == "cbgpPeerPrefixLimit" { return "Cbgppeerprefixlimit" }
    if yname == "cbgpPeerPrefixAdvertised" { return "Cbgppeerprefixadvertised" }
    if yname == "cbgpPeerPrefixSuppressed" { return "Cbgppeerprefixsuppressed" }
    if yname == "cbgpPeerPrefixWithdrawn" { return "Cbgppeerprefixwithdrawn" }
    if yname == "cbgpPeerLastErrorTxt" { return "Cbgppeerlasterrortxt" }
    if yname == "cbgpPeerPrevState" { return "Cbgppeerprevstate" }
    return ""
}

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetSegmentPath() string {
    return "bgpPeerEntry" + "[bgpPeerRemoteAddr='" + fmt.Sprintf("%v", bgppeerentry.Bgppeerremoteaddr) + "']"
}

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgpPeerRemoteAddr"] = bgppeerentry.Bgppeerremoteaddr
    leafs["bgpPeerIdentifier"] = bgppeerentry.Bgppeeridentifier
    leafs["bgpPeerState"] = bgppeerentry.Bgppeerstate
    leafs["bgpPeerAdminStatus"] = bgppeerentry.Bgppeeradminstatus
    leafs["bgpPeerNegotiatedVersion"] = bgppeerentry.Bgppeernegotiatedversion
    leafs["bgpPeerLocalAddr"] = bgppeerentry.Bgppeerlocaladdr
    leafs["bgpPeerLocalPort"] = bgppeerentry.Bgppeerlocalport
    leafs["bgpPeerRemotePort"] = bgppeerentry.Bgppeerremoteport
    leafs["bgpPeerRemoteAs"] = bgppeerentry.Bgppeerremoteas
    leafs["bgpPeerInUpdates"] = bgppeerentry.Bgppeerinupdates
    leafs["bgpPeerOutUpdates"] = bgppeerentry.Bgppeeroutupdates
    leafs["bgpPeerInTotalMessages"] = bgppeerentry.Bgppeerintotalmessages
    leafs["bgpPeerOutTotalMessages"] = bgppeerentry.Bgppeerouttotalmessages
    leafs["bgpPeerLastError"] = bgppeerentry.Bgppeerlasterror
    leafs["bgpPeerFsmEstablishedTransitions"] = bgppeerentry.Bgppeerfsmestablishedtransitions
    leafs["bgpPeerFsmEstablishedTime"] = bgppeerentry.Bgppeerfsmestablishedtime
    leafs["bgpPeerConnectRetryInterval"] = bgppeerentry.Bgppeerconnectretryinterval
    leafs["bgpPeerHoldTime"] = bgppeerentry.Bgppeerholdtime
    leafs["bgpPeerKeepAlive"] = bgppeerentry.Bgppeerkeepalive
    leafs["bgpPeerHoldTimeConfigured"] = bgppeerentry.Bgppeerholdtimeconfigured
    leafs["bgpPeerKeepAliveConfigured"] = bgppeerentry.Bgppeerkeepaliveconfigured
    leafs["bgpPeerMinASOriginationInterval"] = bgppeerentry.Bgppeerminasoriginationinterval
    leafs["bgpPeerMinRouteAdvertisementInterval"] = bgppeerentry.Bgppeerminrouteadvertisementinterval
    leafs["bgpPeerInUpdateElapsedTime"] = bgppeerentry.Bgppeerinupdateelapsedtime
    leafs["cbgpPeerPrefixAccepted"] = bgppeerentry.Cbgppeerprefixaccepted
    leafs["cbgpPeerPrefixDenied"] = bgppeerentry.Cbgppeerprefixdenied
    leafs["cbgpPeerPrefixLimit"] = bgppeerentry.Cbgppeerprefixlimit
    leafs["cbgpPeerPrefixAdvertised"] = bgppeerentry.Cbgppeerprefixadvertised
    leafs["cbgpPeerPrefixSuppressed"] = bgppeerentry.Cbgppeerprefixsuppressed
    leafs["cbgpPeerPrefixWithdrawn"] = bgppeerentry.Cbgppeerprefixwithdrawn
    leafs["cbgpPeerLastErrorTxt"] = bgppeerentry.Cbgppeerlasterrortxt
    leafs["cbgpPeerPrevState"] = bgppeerentry.Cbgppeerprevstate
    return leafs
}

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetYangName() string { return "bgpPeerEntry" }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) SetParent(parent types.Entity) { bgppeerentry.parent = parent }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetParent() types.Entity { return bgppeerentry.parent }

func (bgppeerentry *BGP4MIB_Bgppeertable_Bgppeerentry) GetParentYangName() string { return "bgpPeerTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a path to a network. The type is slice of
    // BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry.
    Bgppathattrentry []BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetFilter() yfilter.YFilter { return bgprcvdpathattrtable.YFilter }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) SetFilter(yf yfilter.YFilter) { bgprcvdpathattrtable.YFilter = yf }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetGoName(yname string) string {
    if yname == "bgpPathAttrEntry" { return "Bgppathattrentry" }
    return ""
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetSegmentPath() string {
    return "bgpRcvdPathAttrTable"
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgpPathAttrEntry" {
        for _, c := range bgprcvdpathattrtable.Bgppathattrentry {
            if bgprcvdpathattrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry{}
        bgprcvdpathattrtable.Bgppathattrentry = append(bgprcvdpathattrtable.Bgppathattrentry, child)
        return &bgprcvdpathattrtable.Bgppathattrentry[len(bgprcvdpathattrtable.Bgppathattrentry)-1]
    }
    return nil
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgprcvdpathattrtable.Bgppathattrentry {
        children[bgprcvdpathattrtable.Bgppathattrentry[i].GetSegmentPath()] = &bgprcvdpathattrtable.Bgppathattrentry[i]
    }
    return children
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetBundleName() string { return "cisco_ios_xe" }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetYangName() string { return "bgpRcvdPathAttrTable" }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) SetParent(parent types.Entity) { bgprcvdpathattrtable.parent = parent }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetParent() types.Entity { return bgprcvdpathattrtable.parent }

func (bgprcvdpathattrtable *BGP4MIB_Bgprcvdpathattrtable) GetParentYangName() string { return "BGP4-MIB" }

// BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry
// Information about a path to a network.
type BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The address of the destination network. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgppathattrdestnetwork interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgppathattrnexthop interface{}

    // The optional inter-AS metric.  If this attribute has not been provided for
    // this route, the value for this object is 0. The type is interface{} with
    // range: -2147483648..2147483647.
    Bgppathattrinterasmetric interface{}
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetFilter() yfilter.YFilter { return bgppathattrentry.YFilter }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) SetFilter(yf yfilter.YFilter) { bgppathattrentry.YFilter = yf }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetGoName(yname string) string {
    if yname == "bgpPathAttrDestNetwork" { return "Bgppathattrdestnetwork" }
    if yname == "bgpPathAttrPeer" { return "Bgppathattrpeer" }
    if yname == "bgpPathAttrOrigin" { return "Bgppathattrorigin" }
    if yname == "bgpPathAttrASPath" { return "Bgppathattraspath" }
    if yname == "bgpPathAttrNextHop" { return "Bgppathattrnexthop" }
    if yname == "bgpPathAttrInterASMetric" { return "Bgppathattrinterasmetric" }
    return ""
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetSegmentPath() string {
    return "bgpPathAttrEntry" + "[bgpPathAttrDestNetwork='" + fmt.Sprintf("%v", bgppathattrentry.Bgppathattrdestnetwork) + "']" + "[bgpPathAttrPeer='" + fmt.Sprintf("%v", bgppathattrentry.Bgppathattrpeer) + "']"
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgpPathAttrDestNetwork"] = bgppathattrentry.Bgppathattrdestnetwork
    leafs["bgpPathAttrPeer"] = bgppathattrentry.Bgppathattrpeer
    leafs["bgpPathAttrOrigin"] = bgppathattrentry.Bgppathattrorigin
    leafs["bgpPathAttrASPath"] = bgppathattrentry.Bgppathattraspath
    leafs["bgpPathAttrNextHop"] = bgppathattrentry.Bgppathattrnexthop
    leafs["bgpPathAttrInterASMetric"] = bgppathattrentry.Bgppathattrinterasmetric
    return leafs
}

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetBundleName() string { return "cisco_ios_xe" }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetYangName() string { return "bgpPathAttrEntry" }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) SetParent(parent types.Entity) { bgppathattrentry.parent = parent }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetParent() types.Entity { return bgppathattrentry.parent }

func (bgppathattrentry *BGP4MIB_Bgprcvdpathattrtable_Bgppathattrentry) GetParentYangName() string { return "bgpRcvdPathAttrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a path to a network. The type is slice of
    // BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry.
    Bgp4Pathattrentry []BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetFilter() yfilter.YFilter { return bgp4Pathattrtable.YFilter }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) SetFilter(yf yfilter.YFilter) { bgp4Pathattrtable.YFilter = yf }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetGoName(yname string) string {
    if yname == "bgp4PathAttrEntry" { return "Bgp4Pathattrentry" }
    return ""
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetSegmentPath() string {
    return "bgp4PathAttrTable"
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp4PathAttrEntry" {
        for _, c := range bgp4Pathattrtable.Bgp4Pathattrentry {
            if bgp4Pathattrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry{}
        bgp4Pathattrtable.Bgp4Pathattrentry = append(bgp4Pathattrtable.Bgp4Pathattrentry, child)
        return &bgp4Pathattrtable.Bgp4Pathattrentry[len(bgp4Pathattrtable.Bgp4Pathattrentry)-1]
    }
    return nil
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgp4Pathattrtable.Bgp4Pathattrentry {
        children[bgp4Pathattrtable.Bgp4Pathattrentry[i].GetSegmentPath()] = &bgp4Pathattrtable.Bgp4Pathattrentry[i]
    }
    return children
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetBundleName() string { return "cisco_ios_xe" }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetYangName() string { return "bgp4PathAttrTable" }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) SetParent(parent types.Entity) { bgp4Pathattrtable.parent = parent }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetParent() types.Entity { return bgp4Pathattrtable.parent }

func (bgp4Pathattrtable *BGP4MIB_Bgp4Pathattrtable) GetParentYangName() string { return "BGP4-MIB" }

// BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry
// Information about a path to a network.
type BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An IP address prefix in the Network Layer
    // Reachability Information field.  This object is an IP address containing
    // the prefix with length specified by bgp4PathAttrIpAddrPrefixLen. Any bits
    // beyond the length specified by bgp4PathAttrIpAddrPrefixLen are zeroed. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bgp4Pathattripaddrprefix interface{}

    // This attribute is a key. Length in bits of the IP address prefix in the
    // Network Layer Reachability Information field. The type is interface{} with
    // range: 0..32.
    Bgp4Pathattripaddrprefixlen interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetFilter() yfilter.YFilter { return bgp4Pathattrentry.YFilter }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) SetFilter(yf yfilter.YFilter) { bgp4Pathattrentry.YFilter = yf }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetGoName(yname string) string {
    if yname == "bgp4PathAttrIpAddrPrefix" { return "Bgp4Pathattripaddrprefix" }
    if yname == "bgp4PathAttrIpAddrPrefixLen" { return "Bgp4Pathattripaddrprefixlen" }
    if yname == "bgp4PathAttrPeer" { return "Bgp4Pathattrpeer" }
    if yname == "bgp4PathAttrOrigin" { return "Bgp4Pathattrorigin" }
    if yname == "bgp4PathAttrASPathSegment" { return "Bgp4Pathattraspathsegment" }
    if yname == "bgp4PathAttrNextHop" { return "Bgp4Pathattrnexthop" }
    if yname == "bgp4PathAttrMultiExitDisc" { return "Bgp4Pathattrmultiexitdisc" }
    if yname == "bgp4PathAttrLocalPref" { return "Bgp4Pathattrlocalpref" }
    if yname == "bgp4PathAttrAtomicAggregate" { return "Bgp4Pathattratomicaggregate" }
    if yname == "bgp4PathAttrAggregatorAS" { return "Bgp4Pathattraggregatoras" }
    if yname == "bgp4PathAttrAggregatorAddr" { return "Bgp4Pathattraggregatoraddr" }
    if yname == "bgp4PathAttrCalcLocalPref" { return "Bgp4Pathattrcalclocalpref" }
    if yname == "bgp4PathAttrBest" { return "Bgp4Pathattrbest" }
    if yname == "bgp4PathAttrUnknown" { return "Bgp4Pathattrunknown" }
    return ""
}

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetSegmentPath() string {
    return "bgp4PathAttrEntry" + "[bgp4PathAttrIpAddrPrefix='" + fmt.Sprintf("%v", bgp4Pathattrentry.Bgp4Pathattripaddrprefix) + "']" + "[bgp4PathAttrIpAddrPrefixLen='" + fmt.Sprintf("%v", bgp4Pathattrentry.Bgp4Pathattripaddrprefixlen) + "']" + "[bgp4PathAttrPeer='" + fmt.Sprintf("%v", bgp4Pathattrentry.Bgp4Pathattrpeer) + "']"
}

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bgp4PathAttrIpAddrPrefix"] = bgp4Pathattrentry.Bgp4Pathattripaddrprefix
    leafs["bgp4PathAttrIpAddrPrefixLen"] = bgp4Pathattrentry.Bgp4Pathattripaddrprefixlen
    leafs["bgp4PathAttrPeer"] = bgp4Pathattrentry.Bgp4Pathattrpeer
    leafs["bgp4PathAttrOrigin"] = bgp4Pathattrentry.Bgp4Pathattrorigin
    leafs["bgp4PathAttrASPathSegment"] = bgp4Pathattrentry.Bgp4Pathattraspathsegment
    leafs["bgp4PathAttrNextHop"] = bgp4Pathattrentry.Bgp4Pathattrnexthop
    leafs["bgp4PathAttrMultiExitDisc"] = bgp4Pathattrentry.Bgp4Pathattrmultiexitdisc
    leafs["bgp4PathAttrLocalPref"] = bgp4Pathattrentry.Bgp4Pathattrlocalpref
    leafs["bgp4PathAttrAtomicAggregate"] = bgp4Pathattrentry.Bgp4Pathattratomicaggregate
    leafs["bgp4PathAttrAggregatorAS"] = bgp4Pathattrentry.Bgp4Pathattraggregatoras
    leafs["bgp4PathAttrAggregatorAddr"] = bgp4Pathattrentry.Bgp4Pathattraggregatoraddr
    leafs["bgp4PathAttrCalcLocalPref"] = bgp4Pathattrentry.Bgp4Pathattrcalclocalpref
    leafs["bgp4PathAttrBest"] = bgp4Pathattrentry.Bgp4Pathattrbest
    leafs["bgp4PathAttrUnknown"] = bgp4Pathattrentry.Bgp4Pathattrunknown
    return leafs
}

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetBundleName() string { return "cisco_ios_xe" }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetYangName() string { return "bgp4PathAttrEntry" }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) SetParent(parent types.Entity) { bgp4Pathattrentry.parent = parent }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetParent() types.Entity { return bgp4Pathattrentry.parent }

func (bgp4Pathattrentry *BGP4MIB_Bgp4Pathattrtable_Bgp4Pathattrentry) GetParentYangName() string { return "bgp4PathAttrTable" }

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

