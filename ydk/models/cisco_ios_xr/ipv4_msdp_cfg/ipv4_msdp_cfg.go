// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-msdp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   msdp: MSDP Configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_msdp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_msdp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-msdp-cfg msdp}", reflect.TypeOf(Msdp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-msdp-cfg:msdp", reflect.TypeOf(Msdp{}))
}

// MsdpListTypeVrf represents Msdp list type vrf
type MsdpListTypeVrf string

const (
    // List
    MsdpListTypeVrf_list MsdpListTypeVrf = "list"

    // RPList
    MsdpListTypeVrf_rp_list MsdpListTypeVrf = "rp-list"
)

// MsdpFilterTypeVrf represents Msdp filter type vrf
type MsdpFilterTypeVrf string

const (
    // Incoming Mode
    MsdpFilterTypeVrf_incoming MsdpFilterTypeVrf = "incoming"

    // Outgoing Mode
    MsdpFilterTypeVrf_outgoing MsdpFilterTypeVrf = "outgoing"
)

// Msdp
// MSDP Configuration
type Msdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the global MAX SA state for the router. The type is interface{}
    // with range: 1..75000.
    GlobalMaxSa interface{}

    // NSR-Ready delay period for MSDP Peer. The type is interface{} with range:
    // 5..90. Units are second.
    NsrDelay interface{}

    // VRF Table.
    Vrfs Msdp_Vrfs

    // Default Context.
    DefaultContext Msdp_DefaultContext
}

func (msdp *Msdp) GetEntityData() *types.CommonEntityData {
    msdp.EntityData.YFilter = msdp.YFilter
    msdp.EntityData.YangName = "msdp"
    msdp.EntityData.BundleName = "cisco_ios_xr"
    msdp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-msdp-cfg"
    msdp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp"
    msdp.EntityData.AbsolutePath = msdp.EntityData.SegmentPath
    msdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    msdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    msdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    msdp.EntityData.Children = types.NewOrderedMap()
    msdp.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &msdp.Vrfs})
    msdp.EntityData.Children.Append("default-context", types.YChild{"DefaultContext", &msdp.DefaultContext})
    msdp.EntityData.Leafs = types.NewOrderedMap()
    msdp.EntityData.Leafs.Append("global-max-sa", types.YLeaf{"GlobalMaxSa", msdp.GlobalMaxSa})
    msdp.EntityData.Leafs.Append("nsr-delay", types.YLeaf{"NsrDelay", msdp.NsrDelay})

    msdp.EntityData.YListKeys = []string {}

    return &(msdp.EntityData)
}

// Msdp_Vrfs
// VRF Table
type Msdp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of Msdp_Vrfs_Vrf.
    Vrf []*Msdp_Vrfs_Vrf
}

func (vrfs *Msdp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "msdp"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Msdp_Vrfs_Vrf
// VRF Name
type Msdp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // Configure TTL Threshold for MSDP Peer. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Configure inheritable MAX SA state for peers. The type is interface{} with
    // range: 1..75000.
    MaxPeerSa interface{}

    // Configure default peers for the box. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultPeer interface{}

    // Configure interface name used as originator ID. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    OriginatorId interface{}

    // Configure context's MAX SA state for the router. The type is interface{}
    // with range: 1..75000.
    MaxSa interface{}

    // Configure interface name used for MSDP connection. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ConnectSource interface{}

    // Configure this systems SA cache access-lists.
    CacheState Msdp_Vrfs_Vrf_CacheState

    // MSDP keep alive period.
    KeepAlive Msdp_Vrfs_Vrf_KeepAlive

    // Entering Peer Configuration.
    Peers Msdp_Vrfs_Vrf_Peers

    // Filter SA messages from peer.
    SaFilters Msdp_Vrfs_Vrf_SaFilters
}

func (vrf *Msdp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("cache-state", types.YChild{"CacheState", &vrf.CacheState})
    vrf.EntityData.Children.Append("keep-alive", types.YChild{"KeepAlive", &vrf.KeepAlive})
    vrf.EntityData.Children.Append("peers", types.YChild{"Peers", &vrf.Peers})
    vrf.EntityData.Children.Append("sa-filters", types.YChild{"SaFilters", &vrf.SaFilters})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("ttl-threshold", types.YLeaf{"TtlThreshold", vrf.TtlThreshold})
    vrf.EntityData.Leafs.Append("max-peer-sa", types.YLeaf{"MaxPeerSa", vrf.MaxPeerSa})
    vrf.EntityData.Leafs.Append("default-peer", types.YLeaf{"DefaultPeer", vrf.DefaultPeer})
    vrf.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", vrf.OriginatorId})
    vrf.EntityData.Leafs.Append("max-sa", types.YLeaf{"MaxSa", vrf.MaxSa})
    vrf.EntityData.Leafs.Append("connect-source", types.YLeaf{"ConnectSource", vrf.ConnectSource})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Msdp_Vrfs_Vrf_CacheState
// Configure this systems SA cache access-lists
type Msdp_Vrfs_Vrf_CacheState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA State Holdtime period. The type is interface{} with range: 150..3600.
    // Units are second. The default value is 150.
    SaHoldtime interface{}

    // Access list name. The type is string with length: 1..64.
    List interface{}

    // Access-list for originating RP. The type is string with length: 1..64.
    RpList interface{}
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetEntityData() *types.CommonEntityData {
    cacheState.EntityData.YFilter = cacheState.YFilter
    cacheState.EntityData.YangName = "cache-state"
    cacheState.EntityData.BundleName = "cisco_ios_xr"
    cacheState.EntityData.ParentYangName = "vrf"
    cacheState.EntityData.SegmentPath = "cache-state"
    cacheState.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/" + cacheState.EntityData.SegmentPath
    cacheState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cacheState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cacheState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cacheState.EntityData.Children = types.NewOrderedMap()
    cacheState.EntityData.Leafs = types.NewOrderedMap()
    cacheState.EntityData.Leafs.Append("sa-holdtime", types.YLeaf{"SaHoldtime", cacheState.SaHoldtime})
    cacheState.EntityData.Leafs.Append("list", types.YLeaf{"List", cacheState.List})
    cacheState.EntityData.Leafs.Append("rp-list", types.YLeaf{"RpList", cacheState.RpList})

    cacheState.EntityData.YListKeys = []string {}

    return &(cacheState.EntityData)
}

// Msdp_Vrfs_Vrf_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_Vrfs_Vrf_KeepAlive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetEntityData() *types.CommonEntityData {
    keepAlive.EntityData.YFilter = keepAlive.YFilter
    keepAlive.EntityData.YangName = "keep-alive"
    keepAlive.EntityData.BundleName = "cisco_ios_xr"
    keepAlive.EntityData.ParentYangName = "vrf"
    keepAlive.EntityData.SegmentPath = "keep-alive"
    keepAlive.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/" + keepAlive.EntityData.SegmentPath
    keepAlive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepAlive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepAlive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepAlive.EntityData.Children = types.NewOrderedMap()
    keepAlive.EntityData.Leafs = types.NewOrderedMap()
    keepAlive.EntityData.Leafs.Append("keep-alive-period", types.YLeaf{"KeepAlivePeriod", keepAlive.KeepAlivePeriod})
    keepAlive.EntityData.Leafs.Append("peer-timeout-period", types.YLeaf{"PeerTimeoutPeriod", keepAlive.PeerTimeoutPeriod})

    keepAlive.EntityData.YListKeys = []string {}

    return &(keepAlive.EntityData)
}

// Msdp_Vrfs_Vrf_Peers
// Entering Peer Configuration
type Msdp_Vrfs_Vrf_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address. The type is slice of Msdp_Vrfs_Vrf_Peers_Peer.
    Peer []*Msdp_Vrfs_Vrf_Peers_Peer
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "vrf"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/" + peers.EntityData.SegmentPath
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// Msdp_Vrfs_Vrf_Peers_Peer
// Peer address
type Msdp_Vrfs_Vrf_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Peer address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // MSDP Peer Shutdown. The type is interface{}.
    Shutdown interface{}

    // Up to 80 characters describing this peer. The type is string with length:
    // 1..80.
    Description interface{}

    // Enabling Peer Configuration. The type is interface{}.
    Enable interface{}

    // Maximum SA accepted from this peer. The type is interface{} with range:
    // 1..75000.
    MaxSa interface{}

    // Disable NSR for the peer. The type is interface{}.
    NsrDown interface{}

    // Configuration of password of peer. The type is string with pattern:
    // (!.+)|([^!].+).
    PeerPassword interface{}

    // Configure an MSDP mesh-group. The type is string with length: 1..32.
    MeshGroup interface{}

    // Configure TTL Threshold for MSDP Peer. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Configure interface name used for MSDP connection. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ConnectSource interface{}

    // Configure the remote AS of this peer.
    RemoteAs Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs

    // MSDP keep alive period.
    KeepAlive Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive

    // Filter SA messages from peer.
    SaFilters Msdp_Vrfs_Vrf_Peers_Peer_SaFilters
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PeerAddress, "peer-address")
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/peers/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("remote-as", types.YChild{"RemoteAs", &peer.RemoteAs})
    peer.EntityData.Children.Append("keep-alive", types.YChild{"KeepAlive", &peer.KeepAlive})
    peer.EntityData.Children.Append("sa-filters", types.YChild{"SaFilters", &peer.SaFilters})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", peer.PeerAddress})
    peer.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", peer.Shutdown})
    peer.EntityData.Leafs.Append("description", types.YLeaf{"Description", peer.Description})
    peer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", peer.Enable})
    peer.EntityData.Leafs.Append("max-sa", types.YLeaf{"MaxSa", peer.MaxSa})
    peer.EntityData.Leafs.Append("nsr-down", types.YLeaf{"NsrDown", peer.NsrDown})
    peer.EntityData.Leafs.Append("peer-password", types.YLeaf{"PeerPassword", peer.PeerPassword})
    peer.EntityData.Leafs.Append("mesh-group", types.YLeaf{"MeshGroup", peer.MeshGroup})
    peer.EntityData.Leafs.Append("ttl-threshold", types.YLeaf{"TtlThreshold", peer.TtlThreshold})
    peer.EntityData.Leafs.Append("connect-source", types.YLeaf{"ConnectSource", peer.ConnectSource})

    peer.EntityData.YListKeys = []string {"PeerAddress"}

    return &(peer.EntityData)
}

// Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs
// Configure the remote AS of this peer
// This type is a presence type.
type Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // First half of ASN in asdot format or 0 in asplain. The type is interface{}
    // with range: 0..65535. The default value is 0.
    AsXx interface{}

    // Second half of ASN in asdot format or complete ASN in asplain. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AsYy interface{}
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetEntityData() *types.CommonEntityData {
    remoteAs.EntityData.YFilter = remoteAs.YFilter
    remoteAs.EntityData.YangName = "remote-as"
    remoteAs.EntityData.BundleName = "cisco_ios_xr"
    remoteAs.EntityData.ParentYangName = "peer"
    remoteAs.EntityData.SegmentPath = "remote-as"
    remoteAs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/peers/peer/" + remoteAs.EntityData.SegmentPath
    remoteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAs.EntityData.Children = types.NewOrderedMap()
    remoteAs.EntityData.Leafs = types.NewOrderedMap()
    remoteAs.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", remoteAs.AsXx})
    remoteAs.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", remoteAs.AsYy})

    remoteAs.EntityData.YListKeys = []string {}

    return &(remoteAs.EntityData)
}

// Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetEntityData() *types.CommonEntityData {
    keepAlive.EntityData.YFilter = keepAlive.YFilter
    keepAlive.EntityData.YangName = "keep-alive"
    keepAlive.EntityData.BundleName = "cisco_ios_xr"
    keepAlive.EntityData.ParentYangName = "peer"
    keepAlive.EntityData.SegmentPath = "keep-alive"
    keepAlive.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/peers/peer/" + keepAlive.EntityData.SegmentPath
    keepAlive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepAlive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepAlive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepAlive.EntityData.Children = types.NewOrderedMap()
    keepAlive.EntityData.Leafs = types.NewOrderedMap()
    keepAlive.EntityData.Leafs.Append("keep-alive-period", types.YLeaf{"KeepAlivePeriod", keepAlive.KeepAlivePeriod})
    keepAlive.EntityData.Leafs.Append("peer-timeout-period", types.YLeaf{"PeerTimeoutPeriod", keepAlive.PeerTimeoutPeriod})

    keepAlive.EntityData.YListKeys = []string {}

    return &(keepAlive.EntityData)
}

// Msdp_Vrfs_Vrf_Peers_Peer_SaFilters
// Filter SA messages from peer
type Msdp_Vrfs_Vrf_Peers_Peer_SaFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter.
    SaFilter []*Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetEntityData() *types.CommonEntityData {
    saFilters.EntityData.YFilter = saFilters.YFilter
    saFilters.EntityData.YangName = "sa-filters"
    saFilters.EntityData.BundleName = "cisco_ios_xr"
    saFilters.EntityData.ParentYangName = "peer"
    saFilters.EntityData.SegmentPath = "sa-filters"
    saFilters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/peers/peer/" + saFilters.EntityData.SegmentPath
    saFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilters.EntityData.Children = types.NewOrderedMap()
    saFilters.EntityData.Children.Append("sa-filter", types.YChild{"SaFilter", nil})
    for i := range saFilters.SaFilter {
        saFilters.EntityData.Children.Append(types.GetSegmentPath(saFilters.SaFilter[i]), types.YChild{"SaFilter", saFilters.SaFilter[i]})
    }
    saFilters.EntityData.Leafs = types.NewOrderedMap()

    saFilters.EntityData.YListKeys = []string {}

    return &(saFilters.EntityData)
}

// Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetEntityData() *types.CommonEntityData {
    saFilter.EntityData.YFilter = saFilter.YFilter
    saFilter.EntityData.YangName = "sa-filter"
    saFilter.EntityData.BundleName = "cisco_ios_xr"
    saFilter.EntityData.ParentYangName = "sa-filters"
    saFilter.EntityData.SegmentPath = "sa-filter" + types.AddKeyToken(saFilter.List, "list") + types.AddKeyToken(saFilter.FilterType, "filter-type")
    saFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/peers/peer/sa-filters/" + saFilter.EntityData.SegmentPath
    saFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilter.EntityData.Children = types.NewOrderedMap()
    saFilter.EntityData.Leafs = types.NewOrderedMap()
    saFilter.EntityData.Leafs.Append("list", types.YLeaf{"List", saFilter.List})
    saFilter.EntityData.Leafs.Append("filter-type", types.YLeaf{"FilterType", saFilter.FilterType})
    saFilter.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", saFilter.AccessListName})

    saFilter.EntityData.YListKeys = []string {"List", "FilterType"}

    return &(saFilter.EntityData)
}

// Msdp_Vrfs_Vrf_SaFilters
// Filter SA messages from peer
type Msdp_Vrfs_Vrf_SaFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_Vrfs_Vrf_SaFilters_SaFilter.
    SaFilter []*Msdp_Vrfs_Vrf_SaFilters_SaFilter
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetEntityData() *types.CommonEntityData {
    saFilters.EntityData.YFilter = saFilters.YFilter
    saFilters.EntityData.YangName = "sa-filters"
    saFilters.EntityData.BundleName = "cisco_ios_xr"
    saFilters.EntityData.ParentYangName = "vrf"
    saFilters.EntityData.SegmentPath = "sa-filters"
    saFilters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/" + saFilters.EntityData.SegmentPath
    saFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilters.EntityData.Children = types.NewOrderedMap()
    saFilters.EntityData.Children.Append("sa-filter", types.YChild{"SaFilter", nil})
    for i := range saFilters.SaFilter {
        saFilters.EntityData.Children.Append(types.GetSegmentPath(saFilters.SaFilter[i]), types.YChild{"SaFilter", saFilters.SaFilter[i]})
    }
    saFilters.EntityData.Leafs = types.NewOrderedMap()

    saFilters.EntityData.YListKeys = []string {}

    return &(saFilters.EntityData)
}

// Msdp_Vrfs_Vrf_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_Vrfs_Vrf_SaFilters_SaFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetEntityData() *types.CommonEntityData {
    saFilter.EntityData.YFilter = saFilter.YFilter
    saFilter.EntityData.YangName = "sa-filter"
    saFilter.EntityData.BundleName = "cisco_ios_xr"
    saFilter.EntityData.ParentYangName = "sa-filters"
    saFilter.EntityData.SegmentPath = "sa-filter" + types.AddKeyToken(saFilter.List, "list") + types.AddKeyToken(saFilter.FilterType, "filter-type")
    saFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/vrfs/vrf/sa-filters/" + saFilter.EntityData.SegmentPath
    saFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilter.EntityData.Children = types.NewOrderedMap()
    saFilter.EntityData.Leafs = types.NewOrderedMap()
    saFilter.EntityData.Leafs.Append("list", types.YLeaf{"List", saFilter.List})
    saFilter.EntityData.Leafs.Append("filter-type", types.YLeaf{"FilterType", saFilter.FilterType})
    saFilter.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", saFilter.AccessListName})

    saFilter.EntityData.YListKeys = []string {"List", "FilterType"}

    return &(saFilter.EntityData)
}

// Msdp_DefaultContext
// Default Context
type Msdp_DefaultContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure TTL Threshold for MSDP Peer. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Configure inheritable MAX SA state for peers. The type is interface{} with
    // range: 1..75000.
    MaxPeerSa interface{}

    // Configure default peers for the box. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultPeer interface{}

    // Configure interface name used as originator ID. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    OriginatorId interface{}

    // Configure context's MAX SA state for the router. The type is interface{}
    // with range: 1..75000.
    MaxSa interface{}

    // Configure interface name used for MSDP connection. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ConnectSource interface{}

    // Configure this systems SA cache access-lists.
    CacheState Msdp_DefaultContext_CacheState

    // MSDP keep alive period.
    KeepAlive Msdp_DefaultContext_KeepAlive

    // Entering Peer Configuration.
    Peers Msdp_DefaultContext_Peers

    // Filter SA messages from peer.
    SaFilters Msdp_DefaultContext_SaFilters
}

func (defaultContext *Msdp_DefaultContext) GetEntityData() *types.CommonEntityData {
    defaultContext.EntityData.YFilter = defaultContext.YFilter
    defaultContext.EntityData.YangName = "default-context"
    defaultContext.EntityData.BundleName = "cisco_ios_xr"
    defaultContext.EntityData.ParentYangName = "msdp"
    defaultContext.EntityData.SegmentPath = "default-context"
    defaultContext.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/" + defaultContext.EntityData.SegmentPath
    defaultContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultContext.EntityData.Children = types.NewOrderedMap()
    defaultContext.EntityData.Children.Append("cache-state", types.YChild{"CacheState", &defaultContext.CacheState})
    defaultContext.EntityData.Children.Append("keep-alive", types.YChild{"KeepAlive", &defaultContext.KeepAlive})
    defaultContext.EntityData.Children.Append("peers", types.YChild{"Peers", &defaultContext.Peers})
    defaultContext.EntityData.Children.Append("sa-filters", types.YChild{"SaFilters", &defaultContext.SaFilters})
    defaultContext.EntityData.Leafs = types.NewOrderedMap()
    defaultContext.EntityData.Leafs.Append("ttl-threshold", types.YLeaf{"TtlThreshold", defaultContext.TtlThreshold})
    defaultContext.EntityData.Leafs.Append("max-peer-sa", types.YLeaf{"MaxPeerSa", defaultContext.MaxPeerSa})
    defaultContext.EntityData.Leafs.Append("default-peer", types.YLeaf{"DefaultPeer", defaultContext.DefaultPeer})
    defaultContext.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", defaultContext.OriginatorId})
    defaultContext.EntityData.Leafs.Append("max-sa", types.YLeaf{"MaxSa", defaultContext.MaxSa})
    defaultContext.EntityData.Leafs.Append("connect-source", types.YLeaf{"ConnectSource", defaultContext.ConnectSource})

    defaultContext.EntityData.YListKeys = []string {}

    return &(defaultContext.EntityData)
}

// Msdp_DefaultContext_CacheState
// Configure this systems SA cache access-lists
type Msdp_DefaultContext_CacheState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA State Holdtime period. The type is interface{} with range: 150..3600.
    // Units are second. The default value is 150.
    SaHoldtime interface{}

    // Access list name. The type is string with length: 1..64.
    List interface{}

    // Access-list for originating RP. The type is string with length: 1..64.
    RpList interface{}
}

func (cacheState *Msdp_DefaultContext_CacheState) GetEntityData() *types.CommonEntityData {
    cacheState.EntityData.YFilter = cacheState.YFilter
    cacheState.EntityData.YangName = "cache-state"
    cacheState.EntityData.BundleName = "cisco_ios_xr"
    cacheState.EntityData.ParentYangName = "default-context"
    cacheState.EntityData.SegmentPath = "cache-state"
    cacheState.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/" + cacheState.EntityData.SegmentPath
    cacheState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cacheState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cacheState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cacheState.EntityData.Children = types.NewOrderedMap()
    cacheState.EntityData.Leafs = types.NewOrderedMap()
    cacheState.EntityData.Leafs.Append("sa-holdtime", types.YLeaf{"SaHoldtime", cacheState.SaHoldtime})
    cacheState.EntityData.Leafs.Append("list", types.YLeaf{"List", cacheState.List})
    cacheState.EntityData.Leafs.Append("rp-list", types.YLeaf{"RpList", cacheState.RpList})

    cacheState.EntityData.YListKeys = []string {}

    return &(cacheState.EntityData)
}

// Msdp_DefaultContext_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_DefaultContext_KeepAlive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetEntityData() *types.CommonEntityData {
    keepAlive.EntityData.YFilter = keepAlive.YFilter
    keepAlive.EntityData.YangName = "keep-alive"
    keepAlive.EntityData.BundleName = "cisco_ios_xr"
    keepAlive.EntityData.ParentYangName = "default-context"
    keepAlive.EntityData.SegmentPath = "keep-alive"
    keepAlive.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/" + keepAlive.EntityData.SegmentPath
    keepAlive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepAlive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepAlive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepAlive.EntityData.Children = types.NewOrderedMap()
    keepAlive.EntityData.Leafs = types.NewOrderedMap()
    keepAlive.EntityData.Leafs.Append("keep-alive-period", types.YLeaf{"KeepAlivePeriod", keepAlive.KeepAlivePeriod})
    keepAlive.EntityData.Leafs.Append("peer-timeout-period", types.YLeaf{"PeerTimeoutPeriod", keepAlive.PeerTimeoutPeriod})

    keepAlive.EntityData.YListKeys = []string {}

    return &(keepAlive.EntityData)
}

// Msdp_DefaultContext_Peers
// Entering Peer Configuration
type Msdp_DefaultContext_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address. The type is slice of Msdp_DefaultContext_Peers_Peer.
    Peer []*Msdp_DefaultContext_Peers_Peer
}

func (peers *Msdp_DefaultContext_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "default-context"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/" + peers.EntityData.SegmentPath
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// Msdp_DefaultContext_Peers_Peer
// Peer address
type Msdp_DefaultContext_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Peer address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // MSDP Peer Shutdown. The type is interface{}.
    Shutdown interface{}

    // Up to 80 characters describing this peer. The type is string with length:
    // 1..80.
    Description interface{}

    // Enabling Peer Configuration. The type is interface{}.
    Enable interface{}

    // Maximum SA accepted from this peer. The type is interface{} with range:
    // 1..75000.
    MaxSa interface{}

    // Disable NSR for the peer. The type is interface{}.
    NsrDown interface{}

    // Configuration of password of peer. The type is string with pattern:
    // (!.+)|([^!].+).
    PeerPassword interface{}

    // Configure an MSDP mesh-group. The type is string with length: 1..32.
    MeshGroup interface{}

    // Configure TTL Threshold for MSDP Peer. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Configure interface name used for MSDP connection. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ConnectSource interface{}

    // Configure the remote AS of this peer.
    RemoteAs Msdp_DefaultContext_Peers_Peer_RemoteAs

    // MSDP keep alive period.
    KeepAlive Msdp_DefaultContext_Peers_Peer_KeepAlive

    // Filter SA messages from peer.
    SaFilters Msdp_DefaultContext_Peers_Peer_SaFilters
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PeerAddress, "peer-address")
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/peers/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("remote-as", types.YChild{"RemoteAs", &peer.RemoteAs})
    peer.EntityData.Children.Append("keep-alive", types.YChild{"KeepAlive", &peer.KeepAlive})
    peer.EntityData.Children.Append("sa-filters", types.YChild{"SaFilters", &peer.SaFilters})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", peer.PeerAddress})
    peer.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", peer.Shutdown})
    peer.EntityData.Leafs.Append("description", types.YLeaf{"Description", peer.Description})
    peer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", peer.Enable})
    peer.EntityData.Leafs.Append("max-sa", types.YLeaf{"MaxSa", peer.MaxSa})
    peer.EntityData.Leafs.Append("nsr-down", types.YLeaf{"NsrDown", peer.NsrDown})
    peer.EntityData.Leafs.Append("peer-password", types.YLeaf{"PeerPassword", peer.PeerPassword})
    peer.EntityData.Leafs.Append("mesh-group", types.YLeaf{"MeshGroup", peer.MeshGroup})
    peer.EntityData.Leafs.Append("ttl-threshold", types.YLeaf{"TtlThreshold", peer.TtlThreshold})
    peer.EntityData.Leafs.Append("connect-source", types.YLeaf{"ConnectSource", peer.ConnectSource})

    peer.EntityData.YListKeys = []string {"PeerAddress"}

    return &(peer.EntityData)
}

// Msdp_DefaultContext_Peers_Peer_RemoteAs
// Configure the remote AS of this peer
// This type is a presence type.
type Msdp_DefaultContext_Peers_Peer_RemoteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // First half of ASN in asdot format or 0 in asplain. The type is interface{}
    // with range: 0..65535. The default value is 0.
    AsXx interface{}

    // Second half of ASN in asdot format or complete ASN in asplain. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AsYy interface{}
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetEntityData() *types.CommonEntityData {
    remoteAs.EntityData.YFilter = remoteAs.YFilter
    remoteAs.EntityData.YangName = "remote-as"
    remoteAs.EntityData.BundleName = "cisco_ios_xr"
    remoteAs.EntityData.ParentYangName = "peer"
    remoteAs.EntityData.SegmentPath = "remote-as"
    remoteAs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/peers/peer/" + remoteAs.EntityData.SegmentPath
    remoteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAs.EntityData.Children = types.NewOrderedMap()
    remoteAs.EntityData.Leafs = types.NewOrderedMap()
    remoteAs.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", remoteAs.AsXx})
    remoteAs.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", remoteAs.AsYy})

    remoteAs.EntityData.YListKeys = []string {}

    return &(remoteAs.EntityData)
}

// Msdp_DefaultContext_Peers_Peer_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_DefaultContext_Peers_Peer_KeepAlive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetEntityData() *types.CommonEntityData {
    keepAlive.EntityData.YFilter = keepAlive.YFilter
    keepAlive.EntityData.YangName = "keep-alive"
    keepAlive.EntityData.BundleName = "cisco_ios_xr"
    keepAlive.EntityData.ParentYangName = "peer"
    keepAlive.EntityData.SegmentPath = "keep-alive"
    keepAlive.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/peers/peer/" + keepAlive.EntityData.SegmentPath
    keepAlive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepAlive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepAlive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepAlive.EntityData.Children = types.NewOrderedMap()
    keepAlive.EntityData.Leafs = types.NewOrderedMap()
    keepAlive.EntityData.Leafs.Append("keep-alive-period", types.YLeaf{"KeepAlivePeriod", keepAlive.KeepAlivePeriod})
    keepAlive.EntityData.Leafs.Append("peer-timeout-period", types.YLeaf{"PeerTimeoutPeriod", keepAlive.PeerTimeoutPeriod})

    keepAlive.EntityData.YListKeys = []string {}

    return &(keepAlive.EntityData)
}

// Msdp_DefaultContext_Peers_Peer_SaFilters
// Filter SA messages from peer
type Msdp_DefaultContext_Peers_Peer_SaFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter.
    SaFilter []*Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetEntityData() *types.CommonEntityData {
    saFilters.EntityData.YFilter = saFilters.YFilter
    saFilters.EntityData.YangName = "sa-filters"
    saFilters.EntityData.BundleName = "cisco_ios_xr"
    saFilters.EntityData.ParentYangName = "peer"
    saFilters.EntityData.SegmentPath = "sa-filters"
    saFilters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/peers/peer/" + saFilters.EntityData.SegmentPath
    saFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilters.EntityData.Children = types.NewOrderedMap()
    saFilters.EntityData.Children.Append("sa-filter", types.YChild{"SaFilter", nil})
    for i := range saFilters.SaFilter {
        saFilters.EntityData.Children.Append(types.GetSegmentPath(saFilters.SaFilter[i]), types.YChild{"SaFilter", saFilters.SaFilter[i]})
    }
    saFilters.EntityData.Leafs = types.NewOrderedMap()

    saFilters.EntityData.YListKeys = []string {}

    return &(saFilters.EntityData)
}

// Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetEntityData() *types.CommonEntityData {
    saFilter.EntityData.YFilter = saFilter.YFilter
    saFilter.EntityData.YangName = "sa-filter"
    saFilter.EntityData.BundleName = "cisco_ios_xr"
    saFilter.EntityData.ParentYangName = "sa-filters"
    saFilter.EntityData.SegmentPath = "sa-filter" + types.AddKeyToken(saFilter.List, "list") + types.AddKeyToken(saFilter.FilterType, "filter-type")
    saFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/peers/peer/sa-filters/" + saFilter.EntityData.SegmentPath
    saFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilter.EntityData.Children = types.NewOrderedMap()
    saFilter.EntityData.Leafs = types.NewOrderedMap()
    saFilter.EntityData.Leafs.Append("list", types.YLeaf{"List", saFilter.List})
    saFilter.EntityData.Leafs.Append("filter-type", types.YLeaf{"FilterType", saFilter.FilterType})
    saFilter.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", saFilter.AccessListName})

    saFilter.EntityData.YListKeys = []string {"List", "FilterType"}

    return &(saFilter.EntityData)
}

// Msdp_DefaultContext_SaFilters
// Filter SA messages from peer
type Msdp_DefaultContext_SaFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_DefaultContext_SaFilters_SaFilter.
    SaFilter []*Msdp_DefaultContext_SaFilters_SaFilter
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetEntityData() *types.CommonEntityData {
    saFilters.EntityData.YFilter = saFilters.YFilter
    saFilters.EntityData.YangName = "sa-filters"
    saFilters.EntityData.BundleName = "cisco_ios_xr"
    saFilters.EntityData.ParentYangName = "default-context"
    saFilters.EntityData.SegmentPath = "sa-filters"
    saFilters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/" + saFilters.EntityData.SegmentPath
    saFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilters.EntityData.Children = types.NewOrderedMap()
    saFilters.EntityData.Children.Append("sa-filter", types.YChild{"SaFilter", nil})
    for i := range saFilters.SaFilter {
        saFilters.EntityData.Children.Append(types.GetSegmentPath(saFilters.SaFilter[i]), types.YChild{"SaFilter", saFilters.SaFilter[i]})
    }
    saFilters.EntityData.Leafs = types.NewOrderedMap()

    saFilters.EntityData.YListKeys = []string {}

    return &(saFilters.EntityData)
}

// Msdp_DefaultContext_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_DefaultContext_SaFilters_SaFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetEntityData() *types.CommonEntityData {
    saFilter.EntityData.YFilter = saFilter.YFilter
    saFilter.EntityData.YangName = "sa-filter"
    saFilter.EntityData.BundleName = "cisco_ios_xr"
    saFilter.EntityData.ParentYangName = "sa-filters"
    saFilter.EntityData.SegmentPath = "sa-filter" + types.AddKeyToken(saFilter.List, "list") + types.AddKeyToken(saFilter.FilterType, "filter-type")
    saFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-msdp-cfg:msdp/default-context/sa-filters/" + saFilter.EntityData.SegmentPath
    saFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    saFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    saFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    saFilter.EntityData.Children = types.NewOrderedMap()
    saFilter.EntityData.Leafs = types.NewOrderedMap()
    saFilter.EntityData.Leafs.Append("list", types.YLeaf{"List", saFilter.List})
    saFilter.EntityData.Leafs.Append("filter-type", types.YLeaf{"FilterType", saFilter.FilterType})
    saFilter.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", saFilter.AccessListName})

    saFilter.EntityData.YListKeys = []string {"List", "FilterType"}

    return &(saFilter.EntityData)
}

