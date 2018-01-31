// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-msdp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   msdp: MSDP Configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
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

func (msdp *Msdp) GetFilter() yfilter.YFilter { return msdp.YFilter }

func (msdp *Msdp) SetFilter(yf yfilter.YFilter) { msdp.YFilter = yf }

func (msdp *Msdp) GetGoName(yname string) string {
    if yname == "global-max-sa" { return "GlobalMaxSa" }
    if yname == "nsr-delay" { return "NsrDelay" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "default-context" { return "DefaultContext" }
    return ""
}

func (msdp *Msdp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-msdp-cfg:msdp"
}

func (msdp *Msdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &msdp.Vrfs
    }
    if childYangName == "default-context" {
        return &msdp.DefaultContext
    }
    return nil
}

func (msdp *Msdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &msdp.Vrfs
    children["default-context"] = &msdp.DefaultContext
    return children
}

func (msdp *Msdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-max-sa"] = msdp.GlobalMaxSa
    leafs["nsr-delay"] = msdp.NsrDelay
    return leafs
}

func (msdp *Msdp) GetBundleName() string { return "cisco_ios_xr" }

func (msdp *Msdp) GetYangName() string { return "msdp" }

func (msdp *Msdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (msdp *Msdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (msdp *Msdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (msdp *Msdp) SetParent(parent types.Entity) { msdp.parent = parent }

func (msdp *Msdp) GetParent() types.Entity { return msdp.parent }

func (msdp *Msdp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-msdp-cfg" }

// Msdp_Vrfs
// VRF Table
type Msdp_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of Msdp_Vrfs_Vrf.
    Vrf []Msdp_Vrfs_Vrf
}

func (vrfs *Msdp_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Msdp_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Msdp_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Msdp_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Msdp_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Msdp_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Msdp_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Msdp_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Msdp_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Msdp_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Msdp_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Msdp_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Msdp_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Msdp_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Msdp_Vrfs) GetParentYangName() string { return "msdp" }

// Msdp_Vrfs_Vrf
// VRF Name
type Msdp_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    // pattern: [a-zA-Z0-9./-]+.
    OriginatorId interface{}

    // Configure context's MAX SA state for the router. The type is interface{}
    // with range: 1..75000.
    MaxSa interface{}

    // Configure interface name used for MSDP connection. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (vrf *Msdp_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Msdp_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Msdp_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "max-peer-sa" { return "MaxPeerSa" }
    if yname == "default-peer" { return "DefaultPeer" }
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "max-sa" { return "MaxSa" }
    if yname == "connect-source" { return "ConnectSource" }
    if yname == "cache-state" { return "CacheState" }
    if yname == "keep-alive" { return "KeepAlive" }
    if yname == "peers" { return "Peers" }
    if yname == "sa-filters" { return "SaFilters" }
    return ""
}

func (vrf *Msdp_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Msdp_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cache-state" {
        return &vrf.CacheState
    }
    if childYangName == "keep-alive" {
        return &vrf.KeepAlive
    }
    if childYangName == "peers" {
        return &vrf.Peers
    }
    if childYangName == "sa-filters" {
        return &vrf.SaFilters
    }
    return nil
}

func (vrf *Msdp_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cache-state"] = &vrf.CacheState
    children["keep-alive"] = &vrf.KeepAlive
    children["peers"] = &vrf.Peers
    children["sa-filters"] = &vrf.SaFilters
    return children
}

func (vrf *Msdp_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["ttl-threshold"] = vrf.TtlThreshold
    leafs["max-peer-sa"] = vrf.MaxPeerSa
    leafs["default-peer"] = vrf.DefaultPeer
    leafs["originator-id"] = vrf.OriginatorId
    leafs["max-sa"] = vrf.MaxSa
    leafs["connect-source"] = vrf.ConnectSource
    return leafs
}

func (vrf *Msdp_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Msdp_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Msdp_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Msdp_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Msdp_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Msdp_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Msdp_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Msdp_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Msdp_Vrfs_Vrf_CacheState
// Configure this systems SA cache access-lists
type Msdp_Vrfs_Vrf_CacheState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA State Holdtime period. The type is interface{} with range: 150..3600.
    // Units are second. The default value is 150.
    SaHoldtime interface{}

    // Access list name. The type is string with length: 1..64.
    List interface{}

    // Access-list for originating RP. The type is string with length: 1..64.
    RpList interface{}
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetFilter() yfilter.YFilter { return cacheState.YFilter }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) SetFilter(yf yfilter.YFilter) { cacheState.YFilter = yf }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetGoName(yname string) string {
    if yname == "sa-holdtime" { return "SaHoldtime" }
    if yname == "list" { return "List" }
    if yname == "rp-list" { return "RpList" }
    return ""
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetSegmentPath() string {
    return "cache-state"
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa-holdtime"] = cacheState.SaHoldtime
    leafs["list"] = cacheState.List
    leafs["rp-list"] = cacheState.RpList
    return leafs
}

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetBundleName() string { return "cisco_ios_xr" }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetYangName() string { return "cache-state" }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) SetParent(parent types.Entity) { cacheState.parent = parent }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetParent() types.Entity { return cacheState.parent }

func (cacheState *Msdp_Vrfs_Vrf_CacheState) GetParentYangName() string { return "vrf" }

// Msdp_Vrfs_Vrf_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_Vrfs_Vrf_KeepAlive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetFilter() yfilter.YFilter { return keepAlive.YFilter }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) SetFilter(yf yfilter.YFilter) { keepAlive.YFilter = yf }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetGoName(yname string) string {
    if yname == "keep-alive-period" { return "KeepAlivePeriod" }
    if yname == "peer-timeout-period" { return "PeerTimeoutPeriod" }
    return ""
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetSegmentPath() string {
    return "keep-alive"
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keep-alive-period"] = keepAlive.KeepAlivePeriod
    leafs["peer-timeout-period"] = keepAlive.PeerTimeoutPeriod
    return leafs
}

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetBundleName() string { return "cisco_ios_xr" }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetYangName() string { return "keep-alive" }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) SetParent(parent types.Entity) { keepAlive.parent = parent }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetParent() types.Entity { return keepAlive.parent }

func (keepAlive *Msdp_Vrfs_Vrf_KeepAlive) GetParentYangName() string { return "vrf" }

// Msdp_Vrfs_Vrf_Peers
// Entering Peer Configuration
type Msdp_Vrfs_Vrf_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer address. The type is slice of Msdp_Vrfs_Vrf_Peers_Peer.
    Peer []Msdp_Vrfs_Vrf_Peers_Peer
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *Msdp_Vrfs_Vrf_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *Msdp_Vrfs_Vrf_Peers) GetGoName(yname string) string {
    if yname == "peer" { return "Peer" }
    return ""
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peers.Peer {
            if peers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_Vrfs_Vrf_Peers_Peer{}
        peers.Peer = append(peers.Peer, child)
        return &peers.Peer[len(peers.Peer)-1]
    }
    return nil
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peers.Peer {
        children[peers.Peer[i].GetSegmentPath()] = &peers.Peer[i]
    }
    return children
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peers *Msdp_Vrfs_Vrf_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *Msdp_Vrfs_Vrf_Peers) GetYangName() string { return "peers" }

func (peers *Msdp_Vrfs_Vrf_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *Msdp_Vrfs_Vrf_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *Msdp_Vrfs_Vrf_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *Msdp_Vrfs_Vrf_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *Msdp_Vrfs_Vrf_Peers) GetParent() types.Entity { return peers.parent }

func (peers *Msdp_Vrfs_Vrf_Peers) GetParentYangName() string { return "vrf" }

// Msdp_Vrfs_Vrf_Peers_Peer
// Peer address
type Msdp_Vrfs_Vrf_Peers_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    // pattern: [a-zA-Z0-9./-]+.
    ConnectSource interface{}

    // Configure the remote AS of this peer.
    RemoteAs Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs

    // MSDP keep alive period.
    KeepAlive Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive

    // Filter SA messages from peer.
    SaFilters Msdp_Vrfs_Vrf_Peers_Peer_SaFilters
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "shutdown" { return "Shutdown" }
    if yname == "description" { return "Description" }
    if yname == "enable" { return "Enable" }
    if yname == "max-sa" { return "MaxSa" }
    if yname == "nsr-down" { return "NsrDown" }
    if yname == "peer-password" { return "PeerPassword" }
    if yname == "mesh-group" { return "MeshGroup" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "connect-source" { return "ConnectSource" }
    if yname == "remote-as" { return "RemoteAs" }
    if yname == "keep-alive" { return "KeepAlive" }
    if yname == "sa-filters" { return "SaFilters" }
    return ""
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetSegmentPath() string {
    return "peer" + "[peer-address='" + fmt.Sprintf("%v", peer.PeerAddress) + "']"
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-as" {
        return &peer.RemoteAs
    }
    if childYangName == "keep-alive" {
        return &peer.KeepAlive
    }
    if childYangName == "sa-filters" {
        return &peer.SaFilters
    }
    return nil
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-as"] = &peer.RemoteAs
    children["keep-alive"] = &peer.KeepAlive
    children["sa-filters"] = &peer.SaFilters
    return children
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peer.PeerAddress
    leafs["shutdown"] = peer.Shutdown
    leafs["description"] = peer.Description
    leafs["enable"] = peer.Enable
    leafs["max-sa"] = peer.MaxSa
    leafs["nsr-down"] = peer.NsrDown
    leafs["peer-password"] = peer.PeerPassword
    leafs["mesh-group"] = peer.MeshGroup
    leafs["ttl-threshold"] = peer.TtlThreshold
    leafs["connect-source"] = peer.ConnectSource
    return leafs
}

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetYangName() string { return "peer" }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Msdp_Vrfs_Vrf_Peers_Peer) GetParentYangName() string { return "peers" }

// Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs
// Configure the remote AS of this peer
// This type is a presence type.
type Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First half of ASN in asdot format or 0 in asplain. The type is interface{}
    // with range: 0..65535. The default value is 0.
    AsXx interface{}

    // Second half of ASN in asdot format or complete ASN in asplain. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AsYy interface{}
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetFilter() yfilter.YFilter { return remoteAs.YFilter }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) SetFilter(yf yfilter.YFilter) { remoteAs.YFilter = yf }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    return ""
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetSegmentPath() string {
    return "remote-as"
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = remoteAs.AsXx
    leafs["as-yy"] = remoteAs.AsYy
    return leafs
}

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetBundleName() string { return "cisco_ios_xr" }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetYangName() string { return "remote-as" }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) SetParent(parent types.Entity) { remoteAs.parent = parent }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetParent() types.Entity { return remoteAs.parent }

func (remoteAs *Msdp_Vrfs_Vrf_Peers_Peer_RemoteAs) GetParentYangName() string { return "peer" }

// Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetFilter() yfilter.YFilter { return keepAlive.YFilter }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) SetFilter(yf yfilter.YFilter) { keepAlive.YFilter = yf }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetGoName(yname string) string {
    if yname == "keep-alive-period" { return "KeepAlivePeriod" }
    if yname == "peer-timeout-period" { return "PeerTimeoutPeriod" }
    return ""
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetSegmentPath() string {
    return "keep-alive"
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keep-alive-period"] = keepAlive.KeepAlivePeriod
    leafs["peer-timeout-period"] = keepAlive.PeerTimeoutPeriod
    return leafs
}

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetBundleName() string { return "cisco_ios_xr" }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetYangName() string { return "keep-alive" }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) SetParent(parent types.Entity) { keepAlive.parent = parent }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetParent() types.Entity { return keepAlive.parent }

func (keepAlive *Msdp_Vrfs_Vrf_Peers_Peer_KeepAlive) GetParentYangName() string { return "peer" }

// Msdp_Vrfs_Vrf_Peers_Peer_SaFilters
// Filter SA messages from peer
type Msdp_Vrfs_Vrf_Peers_Peer_SaFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter.
    SaFilter []Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetFilter() yfilter.YFilter { return saFilters.YFilter }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) SetFilter(yf yfilter.YFilter) { saFilters.YFilter = yf }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetGoName(yname string) string {
    if yname == "sa-filter" { return "SaFilter" }
    return ""
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetSegmentPath() string {
    return "sa-filters"
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sa-filter" {
        for _, c := range saFilters.SaFilter {
            if saFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter{}
        saFilters.SaFilter = append(saFilters.SaFilter, child)
        return &saFilters.SaFilter[len(saFilters.SaFilter)-1]
    }
    return nil
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range saFilters.SaFilter {
        children[saFilters.SaFilter[i].GetSegmentPath()] = &saFilters.SaFilter[i]
    }
    return children
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetBundleName() string { return "cisco_ios_xr" }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetYangName() string { return "sa-filters" }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) SetParent(parent types.Entity) { saFilters.parent = parent }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetParent() types.Entity { return saFilters.parent }

func (saFilters *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters) GetParentYangName() string { return "peer" }

// Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetFilter() yfilter.YFilter { return saFilter.YFilter }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) SetFilter(yf yfilter.YFilter) { saFilter.YFilter = yf }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetGoName(yname string) string {
    if yname == "list" { return "List" }
    if yname == "filter-type" { return "FilterType" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetSegmentPath() string {
    return "sa-filter" + "[list='" + fmt.Sprintf("%v", saFilter.List) + "']" + "[filter-type='" + fmt.Sprintf("%v", saFilter.FilterType) + "']"
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["list"] = saFilter.List
    leafs["filter-type"] = saFilter.FilterType
    leafs["access-list-name"] = saFilter.AccessListName
    return leafs
}

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetBundleName() string { return "cisco_ios_xr" }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetYangName() string { return "sa-filter" }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) SetParent(parent types.Entity) { saFilter.parent = parent }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetParent() types.Entity { return saFilter.parent }

func (saFilter *Msdp_Vrfs_Vrf_Peers_Peer_SaFilters_SaFilter) GetParentYangName() string { return "sa-filters" }

// Msdp_Vrfs_Vrf_SaFilters
// Filter SA messages from peer
type Msdp_Vrfs_Vrf_SaFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_Vrfs_Vrf_SaFilters_SaFilter.
    SaFilter []Msdp_Vrfs_Vrf_SaFilters_SaFilter
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetFilter() yfilter.YFilter { return saFilters.YFilter }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) SetFilter(yf yfilter.YFilter) { saFilters.YFilter = yf }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetGoName(yname string) string {
    if yname == "sa-filter" { return "SaFilter" }
    return ""
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetSegmentPath() string {
    return "sa-filters"
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sa-filter" {
        for _, c := range saFilters.SaFilter {
            if saFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_Vrfs_Vrf_SaFilters_SaFilter{}
        saFilters.SaFilter = append(saFilters.SaFilter, child)
        return &saFilters.SaFilter[len(saFilters.SaFilter)-1]
    }
    return nil
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range saFilters.SaFilter {
        children[saFilters.SaFilter[i].GetSegmentPath()] = &saFilters.SaFilter[i]
    }
    return children
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetBundleName() string { return "cisco_ios_xr" }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetYangName() string { return "sa-filters" }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) SetParent(parent types.Entity) { saFilters.parent = parent }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetParent() types.Entity { return saFilters.parent }

func (saFilters *Msdp_Vrfs_Vrf_SaFilters) GetParentYangName() string { return "vrf" }

// Msdp_Vrfs_Vrf_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_Vrfs_Vrf_SaFilters_SaFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetFilter() yfilter.YFilter { return saFilter.YFilter }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) SetFilter(yf yfilter.YFilter) { saFilter.YFilter = yf }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetGoName(yname string) string {
    if yname == "list" { return "List" }
    if yname == "filter-type" { return "FilterType" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetSegmentPath() string {
    return "sa-filter" + "[list='" + fmt.Sprintf("%v", saFilter.List) + "']" + "[filter-type='" + fmt.Sprintf("%v", saFilter.FilterType) + "']"
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["list"] = saFilter.List
    leafs["filter-type"] = saFilter.FilterType
    leafs["access-list-name"] = saFilter.AccessListName
    return leafs
}

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetBundleName() string { return "cisco_ios_xr" }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetYangName() string { return "sa-filter" }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) SetParent(parent types.Entity) { saFilter.parent = parent }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetParent() types.Entity { return saFilter.parent }

func (saFilter *Msdp_Vrfs_Vrf_SaFilters_SaFilter) GetParentYangName() string { return "sa-filters" }

// Msdp_DefaultContext
// Default Context
// This type is a presence type.
type Msdp_DefaultContext struct {
    parent types.Entity
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
    // pattern: [a-zA-Z0-9./-]+.
    OriginatorId interface{}

    // Configure context's MAX SA state for the router. The type is interface{}
    // with range: 1..75000.
    MaxSa interface{}

    // Configure interface name used for MSDP connection. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (defaultContext *Msdp_DefaultContext) GetFilter() yfilter.YFilter { return defaultContext.YFilter }

func (defaultContext *Msdp_DefaultContext) SetFilter(yf yfilter.YFilter) { defaultContext.YFilter = yf }

func (defaultContext *Msdp_DefaultContext) GetGoName(yname string) string {
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "max-peer-sa" { return "MaxPeerSa" }
    if yname == "default-peer" { return "DefaultPeer" }
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "max-sa" { return "MaxSa" }
    if yname == "connect-source" { return "ConnectSource" }
    if yname == "cache-state" { return "CacheState" }
    if yname == "keep-alive" { return "KeepAlive" }
    if yname == "peers" { return "Peers" }
    if yname == "sa-filters" { return "SaFilters" }
    return ""
}

func (defaultContext *Msdp_DefaultContext) GetSegmentPath() string {
    return "default-context"
}

func (defaultContext *Msdp_DefaultContext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cache-state" {
        return &defaultContext.CacheState
    }
    if childYangName == "keep-alive" {
        return &defaultContext.KeepAlive
    }
    if childYangName == "peers" {
        return &defaultContext.Peers
    }
    if childYangName == "sa-filters" {
        return &defaultContext.SaFilters
    }
    return nil
}

func (defaultContext *Msdp_DefaultContext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cache-state"] = &defaultContext.CacheState
    children["keep-alive"] = &defaultContext.KeepAlive
    children["peers"] = &defaultContext.Peers
    children["sa-filters"] = &defaultContext.SaFilters
    return children
}

func (defaultContext *Msdp_DefaultContext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ttl-threshold"] = defaultContext.TtlThreshold
    leafs["max-peer-sa"] = defaultContext.MaxPeerSa
    leafs["default-peer"] = defaultContext.DefaultPeer
    leafs["originator-id"] = defaultContext.OriginatorId
    leafs["max-sa"] = defaultContext.MaxSa
    leafs["connect-source"] = defaultContext.ConnectSource
    return leafs
}

func (defaultContext *Msdp_DefaultContext) GetBundleName() string { return "cisco_ios_xr" }

func (defaultContext *Msdp_DefaultContext) GetYangName() string { return "default-context" }

func (defaultContext *Msdp_DefaultContext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultContext *Msdp_DefaultContext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultContext *Msdp_DefaultContext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultContext *Msdp_DefaultContext) SetParent(parent types.Entity) { defaultContext.parent = parent }

func (defaultContext *Msdp_DefaultContext) GetParent() types.Entity { return defaultContext.parent }

func (defaultContext *Msdp_DefaultContext) GetParentYangName() string { return "msdp" }

// Msdp_DefaultContext_CacheState
// Configure this systems SA cache access-lists
type Msdp_DefaultContext_CacheState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA State Holdtime period. The type is interface{} with range: 150..3600.
    // Units are second. The default value is 150.
    SaHoldtime interface{}

    // Access list name. The type is string with length: 1..64.
    List interface{}

    // Access-list for originating RP. The type is string with length: 1..64.
    RpList interface{}
}

func (cacheState *Msdp_DefaultContext_CacheState) GetFilter() yfilter.YFilter { return cacheState.YFilter }

func (cacheState *Msdp_DefaultContext_CacheState) SetFilter(yf yfilter.YFilter) { cacheState.YFilter = yf }

func (cacheState *Msdp_DefaultContext_CacheState) GetGoName(yname string) string {
    if yname == "sa-holdtime" { return "SaHoldtime" }
    if yname == "list" { return "List" }
    if yname == "rp-list" { return "RpList" }
    return ""
}

func (cacheState *Msdp_DefaultContext_CacheState) GetSegmentPath() string {
    return "cache-state"
}

func (cacheState *Msdp_DefaultContext_CacheState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cacheState *Msdp_DefaultContext_CacheState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cacheState *Msdp_DefaultContext_CacheState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa-holdtime"] = cacheState.SaHoldtime
    leafs["list"] = cacheState.List
    leafs["rp-list"] = cacheState.RpList
    return leafs
}

func (cacheState *Msdp_DefaultContext_CacheState) GetBundleName() string { return "cisco_ios_xr" }

func (cacheState *Msdp_DefaultContext_CacheState) GetYangName() string { return "cache-state" }

func (cacheState *Msdp_DefaultContext_CacheState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cacheState *Msdp_DefaultContext_CacheState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cacheState *Msdp_DefaultContext_CacheState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cacheState *Msdp_DefaultContext_CacheState) SetParent(parent types.Entity) { cacheState.parent = parent }

func (cacheState *Msdp_DefaultContext_CacheState) GetParent() types.Entity { return cacheState.parent }

func (cacheState *Msdp_DefaultContext_CacheState) GetParentYangName() string { return "default-context" }

// Msdp_DefaultContext_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_DefaultContext_KeepAlive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetFilter() yfilter.YFilter { return keepAlive.YFilter }

func (keepAlive *Msdp_DefaultContext_KeepAlive) SetFilter(yf yfilter.YFilter) { keepAlive.YFilter = yf }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetGoName(yname string) string {
    if yname == "keep-alive-period" { return "KeepAlivePeriod" }
    if yname == "peer-timeout-period" { return "PeerTimeoutPeriod" }
    return ""
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetSegmentPath() string {
    return "keep-alive"
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keep-alive-period"] = keepAlive.KeepAlivePeriod
    leafs["peer-timeout-period"] = keepAlive.PeerTimeoutPeriod
    return leafs
}

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetBundleName() string { return "cisco_ios_xr" }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetYangName() string { return "keep-alive" }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepAlive *Msdp_DefaultContext_KeepAlive) SetParent(parent types.Entity) { keepAlive.parent = parent }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetParent() types.Entity { return keepAlive.parent }

func (keepAlive *Msdp_DefaultContext_KeepAlive) GetParentYangName() string { return "default-context" }

// Msdp_DefaultContext_Peers
// Entering Peer Configuration
type Msdp_DefaultContext_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer address. The type is slice of Msdp_DefaultContext_Peers_Peer.
    Peer []Msdp_DefaultContext_Peers_Peer
}

func (peers *Msdp_DefaultContext_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *Msdp_DefaultContext_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *Msdp_DefaultContext_Peers) GetGoName(yname string) string {
    if yname == "peer" { return "Peer" }
    return ""
}

func (peers *Msdp_DefaultContext_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *Msdp_DefaultContext_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peers.Peer {
            if peers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_DefaultContext_Peers_Peer{}
        peers.Peer = append(peers.Peer, child)
        return &peers.Peer[len(peers.Peer)-1]
    }
    return nil
}

func (peers *Msdp_DefaultContext_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peers.Peer {
        children[peers.Peer[i].GetSegmentPath()] = &peers.Peer[i]
    }
    return children
}

func (peers *Msdp_DefaultContext_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peers *Msdp_DefaultContext_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *Msdp_DefaultContext_Peers) GetYangName() string { return "peers" }

func (peers *Msdp_DefaultContext_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *Msdp_DefaultContext_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *Msdp_DefaultContext_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *Msdp_DefaultContext_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *Msdp_DefaultContext_Peers) GetParent() types.Entity { return peers.parent }

func (peers *Msdp_DefaultContext_Peers) GetParentYangName() string { return "default-context" }

// Msdp_DefaultContext_Peers_Peer
// Peer address
type Msdp_DefaultContext_Peers_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    // pattern: [a-zA-Z0-9./-]+.
    ConnectSource interface{}

    // Configure the remote AS of this peer.
    RemoteAs Msdp_DefaultContext_Peers_Peer_RemoteAs

    // MSDP keep alive period.
    KeepAlive Msdp_DefaultContext_Peers_Peer_KeepAlive

    // Filter SA messages from peer.
    SaFilters Msdp_DefaultContext_Peers_Peer_SaFilters
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Msdp_DefaultContext_Peers_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Msdp_DefaultContext_Peers_Peer) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "shutdown" { return "Shutdown" }
    if yname == "description" { return "Description" }
    if yname == "enable" { return "Enable" }
    if yname == "max-sa" { return "MaxSa" }
    if yname == "nsr-down" { return "NsrDown" }
    if yname == "peer-password" { return "PeerPassword" }
    if yname == "mesh-group" { return "MeshGroup" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "connect-source" { return "ConnectSource" }
    if yname == "remote-as" { return "RemoteAs" }
    if yname == "keep-alive" { return "KeepAlive" }
    if yname == "sa-filters" { return "SaFilters" }
    return ""
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetSegmentPath() string {
    return "peer" + "[peer-address='" + fmt.Sprintf("%v", peer.PeerAddress) + "']"
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-as" {
        return &peer.RemoteAs
    }
    if childYangName == "keep-alive" {
        return &peer.KeepAlive
    }
    if childYangName == "sa-filters" {
        return &peer.SaFilters
    }
    return nil
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-as"] = &peer.RemoteAs
    children["keep-alive"] = &peer.KeepAlive
    children["sa-filters"] = &peer.SaFilters
    return children
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peer.PeerAddress
    leafs["shutdown"] = peer.Shutdown
    leafs["description"] = peer.Description
    leafs["enable"] = peer.Enable
    leafs["max-sa"] = peer.MaxSa
    leafs["nsr-down"] = peer.NsrDown
    leafs["peer-password"] = peer.PeerPassword
    leafs["mesh-group"] = peer.MeshGroup
    leafs["ttl-threshold"] = peer.TtlThreshold
    leafs["connect-source"] = peer.ConnectSource
    return leafs
}

func (peer *Msdp_DefaultContext_Peers_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Msdp_DefaultContext_Peers_Peer) GetYangName() string { return "peer" }

func (peer *Msdp_DefaultContext_Peers_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Msdp_DefaultContext_Peers_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Msdp_DefaultContext_Peers_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Msdp_DefaultContext_Peers_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Msdp_DefaultContext_Peers_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Msdp_DefaultContext_Peers_Peer) GetParentYangName() string { return "peers" }

// Msdp_DefaultContext_Peers_Peer_RemoteAs
// Configure the remote AS of this peer
// This type is a presence type.
type Msdp_DefaultContext_Peers_Peer_RemoteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First half of ASN in asdot format or 0 in asplain. The type is interface{}
    // with range: 0..65535. The default value is 0.
    AsXx interface{}

    // Second half of ASN in asdot format or complete ASN in asplain. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AsYy interface{}
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetFilter() yfilter.YFilter { return remoteAs.YFilter }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) SetFilter(yf yfilter.YFilter) { remoteAs.YFilter = yf }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    return ""
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetSegmentPath() string {
    return "remote-as"
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = remoteAs.AsXx
    leafs["as-yy"] = remoteAs.AsYy
    return leafs
}

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetBundleName() string { return "cisco_ios_xr" }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetYangName() string { return "remote-as" }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) SetParent(parent types.Entity) { remoteAs.parent = parent }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetParent() types.Entity { return remoteAs.parent }

func (remoteAs *Msdp_DefaultContext_Peers_Peer_RemoteAs) GetParentYangName() string { return "peer" }

// Msdp_DefaultContext_Peers_Peer_KeepAlive
// MSDP keep alive period
// This type is a presence type.
type Msdp_DefaultContext_Peers_Peer_KeepAlive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Keep alive period in seconds. The type is interface{} with range: 1..60.
    // This attribute is mandatory. Units are second.
    KeepAlivePeriod interface{}

    // Peer timeout period in seconds. The type is interface{} with range: 1..75.
    // This attribute is mandatory. Units are second.
    PeerTimeoutPeriod interface{}
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetFilter() yfilter.YFilter { return keepAlive.YFilter }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) SetFilter(yf yfilter.YFilter) { keepAlive.YFilter = yf }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetGoName(yname string) string {
    if yname == "keep-alive-period" { return "KeepAlivePeriod" }
    if yname == "peer-timeout-period" { return "PeerTimeoutPeriod" }
    return ""
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetSegmentPath() string {
    return "keep-alive"
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keep-alive-period"] = keepAlive.KeepAlivePeriod
    leafs["peer-timeout-period"] = keepAlive.PeerTimeoutPeriod
    return leafs
}

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetBundleName() string { return "cisco_ios_xr" }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetYangName() string { return "keep-alive" }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) SetParent(parent types.Entity) { keepAlive.parent = parent }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetParent() types.Entity { return keepAlive.parent }

func (keepAlive *Msdp_DefaultContext_Peers_Peer_KeepAlive) GetParentYangName() string { return "peer" }

// Msdp_DefaultContext_Peers_Peer_SaFilters
// Filter SA messages from peer
type Msdp_DefaultContext_Peers_Peer_SaFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter.
    SaFilter []Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetFilter() yfilter.YFilter { return saFilters.YFilter }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) SetFilter(yf yfilter.YFilter) { saFilters.YFilter = yf }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetGoName(yname string) string {
    if yname == "sa-filter" { return "SaFilter" }
    return ""
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetSegmentPath() string {
    return "sa-filters"
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sa-filter" {
        for _, c := range saFilters.SaFilter {
            if saFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter{}
        saFilters.SaFilter = append(saFilters.SaFilter, child)
        return &saFilters.SaFilter[len(saFilters.SaFilter)-1]
    }
    return nil
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range saFilters.SaFilter {
        children[saFilters.SaFilter[i].GetSegmentPath()] = &saFilters.SaFilter[i]
    }
    return children
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetBundleName() string { return "cisco_ios_xr" }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetYangName() string { return "sa-filters" }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) SetParent(parent types.Entity) { saFilters.parent = parent }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetParent() types.Entity { return saFilters.parent }

func (saFilters *Msdp_DefaultContext_Peers_Peer_SaFilters) GetParentYangName() string { return "peer" }

// Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetFilter() yfilter.YFilter { return saFilter.YFilter }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) SetFilter(yf yfilter.YFilter) { saFilter.YFilter = yf }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetGoName(yname string) string {
    if yname == "list" { return "List" }
    if yname == "filter-type" { return "FilterType" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetSegmentPath() string {
    return "sa-filter" + "[list='" + fmt.Sprintf("%v", saFilter.List) + "']" + "[filter-type='" + fmt.Sprintf("%v", saFilter.FilterType) + "']"
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["list"] = saFilter.List
    leafs["filter-type"] = saFilter.FilterType
    leafs["access-list-name"] = saFilter.AccessListName
    return leafs
}

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetBundleName() string { return "cisco_ios_xr" }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetYangName() string { return "sa-filter" }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) SetParent(parent types.Entity) { saFilter.parent = parent }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetParent() types.Entity { return saFilter.parent }

func (saFilter *Msdp_DefaultContext_Peers_Peer_SaFilters_SaFilter) GetParentYangName() string { return "sa-filters" }

// Msdp_DefaultContext_SaFilters
// Filter SA messages from peer
type Msdp_DefaultContext_SaFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA-Filter incoming/outgoing list or RPlist. The type is slice of
    // Msdp_DefaultContext_SaFilters_SaFilter.
    SaFilter []Msdp_DefaultContext_SaFilters_SaFilter
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetFilter() yfilter.YFilter { return saFilters.YFilter }

func (saFilters *Msdp_DefaultContext_SaFilters) SetFilter(yf yfilter.YFilter) { saFilters.YFilter = yf }

func (saFilters *Msdp_DefaultContext_SaFilters) GetGoName(yname string) string {
    if yname == "sa-filter" { return "SaFilter" }
    return ""
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetSegmentPath() string {
    return "sa-filters"
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sa-filter" {
        for _, c := range saFilters.SaFilter {
            if saFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Msdp_DefaultContext_SaFilters_SaFilter{}
        saFilters.SaFilter = append(saFilters.SaFilter, child)
        return &saFilters.SaFilter[len(saFilters.SaFilter)-1]
    }
    return nil
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range saFilters.SaFilter {
        children[saFilters.SaFilter[i].GetSegmentPath()] = &saFilters.SaFilter[i]
    }
    return children
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (saFilters *Msdp_DefaultContext_SaFilters) GetBundleName() string { return "cisco_ios_xr" }

func (saFilters *Msdp_DefaultContext_SaFilters) GetYangName() string { return "sa-filters" }

func (saFilters *Msdp_DefaultContext_SaFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilters *Msdp_DefaultContext_SaFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilters *Msdp_DefaultContext_SaFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilters *Msdp_DefaultContext_SaFilters) SetParent(parent types.Entity) { saFilters.parent = parent }

func (saFilters *Msdp_DefaultContext_SaFilters) GetParent() types.Entity { return saFilters.parent }

func (saFilters *Msdp_DefaultContext_SaFilters) GetParentYangName() string { return "default-context" }

// Msdp_DefaultContext_SaFilters_SaFilter
// SA-Filter incoming/outgoing list or RPlist
type Msdp_DefaultContext_SaFilters_SaFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Src List/RP List. The type is MsdpListTypeVrf.
    List interface{}

    // This attribute is a key. Incoming/Outgoing ACL. The type is
    // MsdpFilterTypeVrf.
    FilterType interface{}

    // Access list name. The type is string with length: 1..64. This attribute is
    // mandatory.
    AccessListName interface{}
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetFilter() yfilter.YFilter { return saFilter.YFilter }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) SetFilter(yf yfilter.YFilter) { saFilter.YFilter = yf }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetGoName(yname string) string {
    if yname == "list" { return "List" }
    if yname == "filter-type" { return "FilterType" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetSegmentPath() string {
    return "sa-filter" + "[list='" + fmt.Sprintf("%v", saFilter.List) + "']" + "[filter-type='" + fmt.Sprintf("%v", saFilter.FilterType) + "']"
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["list"] = saFilter.List
    leafs["filter-type"] = saFilter.FilterType
    leafs["access-list-name"] = saFilter.AccessListName
    return leafs
}

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetBundleName() string { return "cisco_ios_xr" }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetYangName() string { return "sa-filter" }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) SetParent(parent types.Entity) { saFilter.parent = parent }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetParent() types.Entity { return saFilter.parent }

func (saFilter *Msdp_DefaultContext_SaFilters_SaFilter) GetParentYangName() string { return "sa-filters" }

