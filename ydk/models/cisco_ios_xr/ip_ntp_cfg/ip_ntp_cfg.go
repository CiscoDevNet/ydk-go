// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-ntp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ntp: NTP configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_ntp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_ntp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-ntp-cfg ntp}", reflect.TypeOf(Ntp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-ntp-cfg:ntp", reflect.TypeOf(Ntp{}))
}

// NtpAccessAf represents Ntp access af
type NtpAccessAf string

const (
    // IPv4
    NtpAccessAf_ipv4 NtpAccessAf = "ipv4"

    // IPv6
    NtpAccessAf_ipv6 NtpAccessAf = "ipv6"
)

// NtpPeer represents Ntp peer
type NtpPeer string

const (
    // Peer
    NtpPeer_peer NtpPeer = "peer"

    // Server
    NtpPeer_server NtpPeer = "server"
)

// Ntpdscp represents Ntpdscp
type Ntpdscp string

const (
    // Precedence Value
    Ntpdscp_ntp_precedence Ntpdscp = "ntp-precedence"

    // DSCP Value
    Ntpdscp_ntpdscp Ntpdscp = "ntpdscp"
)

// NtpAccess represents Ntp access
type NtpAccess string

const (
    // Peer
    NtpAccess_peer NtpAccess = "peer"

    // Serve
    NtpAccess_serve NtpAccess = "serve"

    // Serve Only
    NtpAccess_serve_only NtpAccess = "serve-only"

    // Query Only
    NtpAccess_query_only NtpAccess = "query-only"
)

// Ntp
// NTP configuration
// This type is a presence type.
type Ntp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set maximum number of associations. The type is interface{} with range:
    // -2147483648..2147483647.
    MaxAssociations interface{}

    // Act as NTP master clock. The type is interface{} with range: 1..15. The
    // default value is 8.
    Master interface{}

    // Estimated round-trip delay. The type is interface{} with range: 1..999999.
    BroadcastDelay interface{}

    // To enable logging internal sync conflicts. The type is interface{}.
    LogInternalSync interface{}

    // To enable calendar update with NTP time. The type is interface{}.
    UpdateCalendar interface{}

    // Configures NTP Peers or Servers.
    PeerVrfs Ntp_PeerVrfs

    // Set IP DSCP value for outgoing NTP IPV4 packets.
    DscpIpv4 Ntp_DscpIpv4

    // Set IP DSCP value for outgoing NTP IPV6 packets.
    DscpIpv6 Ntp_DscpIpv6

    // Configure  NTP source interface.
    Sources Ntp_Sources

    // Configure NTP Authentication keys.
    Authentication Ntp_Authentication

    // Configure NTP passive associations.
    Passive Ntp_Passive

    // NTP per interface configuration.
    InterfaceTables Ntp_InterfaceTables

    // Control NTP access.
    AccessGroupTables Ntp_AccessGroupTables
}

func (ntp *Ntp) GetFilter() yfilter.YFilter { return ntp.YFilter }

func (ntp *Ntp) SetFilter(yf yfilter.YFilter) { ntp.YFilter = yf }

func (ntp *Ntp) GetGoName(yname string) string {
    if yname == "max-associations" { return "MaxAssociations" }
    if yname == "master" { return "Master" }
    if yname == "broadcast-delay" { return "BroadcastDelay" }
    if yname == "log-internal-sync" { return "LogInternalSync" }
    if yname == "update-calendar" { return "UpdateCalendar" }
    if yname == "peer-vrfs" { return "PeerVrfs" }
    if yname == "dscp-ipv4" { return "DscpIpv4" }
    if yname == "dscp-ipv6" { return "DscpIpv6" }
    if yname == "sources" { return "Sources" }
    if yname == "authentication" { return "Authentication" }
    if yname == "passive" { return "Passive" }
    if yname == "interface-tables" { return "InterfaceTables" }
    if yname == "access-group-tables" { return "AccessGroupTables" }
    return ""
}

func (ntp *Ntp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-ntp-cfg:ntp"
}

func (ntp *Ntp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-vrfs" {
        return &ntp.PeerVrfs
    }
    if childYangName == "dscp-ipv4" {
        return &ntp.DscpIpv4
    }
    if childYangName == "dscp-ipv6" {
        return &ntp.DscpIpv6
    }
    if childYangName == "sources" {
        return &ntp.Sources
    }
    if childYangName == "authentication" {
        return &ntp.Authentication
    }
    if childYangName == "passive" {
        return &ntp.Passive
    }
    if childYangName == "interface-tables" {
        return &ntp.InterfaceTables
    }
    if childYangName == "access-group-tables" {
        return &ntp.AccessGroupTables
    }
    return nil
}

func (ntp *Ntp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-vrfs"] = &ntp.PeerVrfs
    children["dscp-ipv4"] = &ntp.DscpIpv4
    children["dscp-ipv6"] = &ntp.DscpIpv6
    children["sources"] = &ntp.Sources
    children["authentication"] = &ntp.Authentication
    children["passive"] = &ntp.Passive
    children["interface-tables"] = &ntp.InterfaceTables
    children["access-group-tables"] = &ntp.AccessGroupTables
    return children
}

func (ntp *Ntp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-associations"] = ntp.MaxAssociations
    leafs["master"] = ntp.Master
    leafs["broadcast-delay"] = ntp.BroadcastDelay
    leafs["log-internal-sync"] = ntp.LogInternalSync
    leafs["update-calendar"] = ntp.UpdateCalendar
    return leafs
}

func (ntp *Ntp) GetBundleName() string { return "cisco_ios_xr" }

func (ntp *Ntp) GetYangName() string { return "ntp" }

func (ntp *Ntp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ntp *Ntp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ntp *Ntp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ntp *Ntp) SetParent(parent types.Entity) { ntp.parent = parent }

func (ntp *Ntp) GetParent() types.Entity { return ntp.parent }

func (ntp *Ntp) GetParentYangName() string { return "Cisco-IOS-XR-ip-ntp-cfg" }

// Ntp_PeerVrfs
// Configures NTP Peers or Servers
type Ntp_PeerVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configures NTP Peers or Servers for a single VRF. The 'default' must also
    // be specified for default VRF. The type is slice of Ntp_PeerVrfs_PeerVrf.
    PeerVrf []Ntp_PeerVrfs_PeerVrf
}

func (peerVrfs *Ntp_PeerVrfs) GetFilter() yfilter.YFilter { return peerVrfs.YFilter }

func (peerVrfs *Ntp_PeerVrfs) SetFilter(yf yfilter.YFilter) { peerVrfs.YFilter = yf }

func (peerVrfs *Ntp_PeerVrfs) GetGoName(yname string) string {
    if yname == "peer-vrf" { return "PeerVrf" }
    return ""
}

func (peerVrfs *Ntp_PeerVrfs) GetSegmentPath() string {
    return "peer-vrfs"
}

func (peerVrfs *Ntp_PeerVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-vrf" {
        for _, c := range peerVrfs.PeerVrf {
            if peerVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_PeerVrfs_PeerVrf{}
        peerVrfs.PeerVrf = append(peerVrfs.PeerVrf, child)
        return &peerVrfs.PeerVrf[len(peerVrfs.PeerVrf)-1]
    }
    return nil
}

func (peerVrfs *Ntp_PeerVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerVrfs.PeerVrf {
        children[peerVrfs.PeerVrf[i].GetSegmentPath()] = &peerVrfs.PeerVrf[i]
    }
    return children
}

func (peerVrfs *Ntp_PeerVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerVrfs *Ntp_PeerVrfs) GetBundleName() string { return "cisco_ios_xr" }

func (peerVrfs *Ntp_PeerVrfs) GetYangName() string { return "peer-vrfs" }

func (peerVrfs *Ntp_PeerVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerVrfs *Ntp_PeerVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerVrfs *Ntp_PeerVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerVrfs *Ntp_PeerVrfs) SetParent(parent types.Entity) { peerVrfs.parent = parent }

func (peerVrfs *Ntp_PeerVrfs) GetParent() types.Entity { return peerVrfs.parent }

func (peerVrfs *Ntp_PeerVrfs) GetParentYangName() string { return "ntp" }

// Ntp_PeerVrfs_PeerVrf
// Configures NTP Peers or Servers for a single
// VRF. The 'default' must also be specified for
// default VRF
type Ntp_PeerVrfs_PeerVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Configures IPv4 NTP Peers or Servers.
    PeerIpv4S Ntp_PeerVrfs_PeerVrf_PeerIpv4S

    // Configuration NTP Peers or Servers of IPV6.
    PeerIpv6S Ntp_PeerVrfs_PeerVrf_PeerIpv6S
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetFilter() yfilter.YFilter { return peerVrf.YFilter }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) SetFilter(yf yfilter.YFilter) { peerVrf.YFilter = yf }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "peer-ipv4s" { return "PeerIpv4S" }
    if yname == "peer-ipv6s" { return "PeerIpv6S" }
    return ""
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetSegmentPath() string {
    return "peer-vrf" + "[vrf-name='" + fmt.Sprintf("%v", peerVrf.VrfName) + "']"
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-ipv4s" {
        return &peerVrf.PeerIpv4S
    }
    if childYangName == "peer-ipv6s" {
        return &peerVrf.PeerIpv6S
    }
    return nil
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-ipv4s"] = &peerVrf.PeerIpv4S
    children["peer-ipv6s"] = &peerVrf.PeerIpv6S
    return children
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = peerVrf.VrfName
    return leafs
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetBundleName() string { return "cisco_ios_xr" }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetYangName() string { return "peer-vrf" }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) SetParent(parent types.Entity) { peerVrf.parent = parent }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetParent() types.Entity { return peerVrf.parent }

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetParentYangName() string { return "peer-vrfs" }

// Ntp_PeerVrfs_PeerVrf_PeerIpv4S
// Configures IPv4 NTP Peers or Servers
type Ntp_PeerVrfs_PeerVrf_PeerIpv4S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure an IPv4 NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4.
    PeerIpv4 []Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetFilter() yfilter.YFilter { return peerIpv4S.YFilter }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) SetFilter(yf yfilter.YFilter) { peerIpv4S.YFilter = yf }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetGoName(yname string) string {
    if yname == "peer-ipv4" { return "PeerIpv4" }
    return ""
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetSegmentPath() string {
    return "peer-ipv4s"
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-ipv4" {
        for _, c := range peerIpv4S.PeerIpv4 {
            if peerIpv4S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4{}
        peerIpv4S.PeerIpv4 = append(peerIpv4S.PeerIpv4, child)
        return &peerIpv4S.PeerIpv4[len(peerIpv4S.PeerIpv4)-1]
    }
    return nil
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerIpv4S.PeerIpv4 {
        children[peerIpv4S.PeerIpv4[i].GetSegmentPath()] = &peerIpv4S.PeerIpv4[i]
    }
    return children
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetBundleName() string { return "cisco_ios_xr" }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetYangName() string { return "peer-ipv4s" }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) SetParent(parent types.Entity) { peerIpv4S.parent = parent }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetParent() types.Entity { return peerIpv4S.parent }

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetParentYangName() string { return "peer-vrf" }

// Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4
// Configure an IPv4 NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 Address of a peer. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AddressIpv4 interface{}

    // Configure an IPv4 NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4.
    PeerTypeIpv4 []Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetFilter() yfilter.YFilter { return peerIpv4.YFilter }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) SetFilter(yf yfilter.YFilter) { peerIpv4.YFilter = yf }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetGoName(yname string) string {
    if yname == "address-ipv4" { return "AddressIpv4" }
    if yname == "peer-type-ipv4" { return "PeerTypeIpv4" }
    return ""
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetSegmentPath() string {
    return "peer-ipv4" + "[address-ipv4='" + fmt.Sprintf("%v", peerIpv4.AddressIpv4) + "']"
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-type-ipv4" {
        for _, c := range peerIpv4.PeerTypeIpv4 {
            if peerIpv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4{}
        peerIpv4.PeerTypeIpv4 = append(peerIpv4.PeerTypeIpv4, child)
        return &peerIpv4.PeerTypeIpv4[len(peerIpv4.PeerTypeIpv4)-1]
    }
    return nil
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerIpv4.PeerTypeIpv4 {
        children[peerIpv4.PeerTypeIpv4[i].GetSegmentPath()] = &peerIpv4.PeerTypeIpv4[i]
    }
    return children
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-ipv4"] = peerIpv4.AddressIpv4
    return leafs
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetYangName() string { return "peer-ipv4" }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) SetParent(parent types.Entity) { peerIpv4.parent = parent }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetParent() types.Entity { return peerIpv4.parent }

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetParentYangName() string { return "peer-ipv4s" }

// Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4
// Configure an IPv4 NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer or Server. The type is NtpPeer.
    PeerType interface{}

    // NTP version. The type is interface{} with range: 2..4.
    NtpVersion interface{}

    // Authentication Key. The type is interface{} with range: 1..65535.
    AuthenticationKey interface{}

    // Minimum poll interval. The type is interface{} with range: 4..17.
    MinPoll interface{}

    // Maxinum poll interval. The type is interface{} with range: 4..17.
    MaxPoll interface{}

    // Preferred peer. The type is interface{}.
    PreferredPeer interface{}

    // Source interface of this peer. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Use burst mode. The type is interface{}.
    Burst interface{}

    // Use iburst mode. The type is interface{}.
    Iburst interface{}
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetFilter() yfilter.YFilter { return peerTypeIpv4.YFilter }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) SetFilter(yf yfilter.YFilter) { peerTypeIpv4.YFilter = yf }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetGoName(yname string) string {
    if yname == "peer-type" { return "PeerType" }
    if yname == "ntp-version" { return "NtpVersion" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    if yname == "min-poll" { return "MinPoll" }
    if yname == "max-poll" { return "MaxPoll" }
    if yname == "preferred-peer" { return "PreferredPeer" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "burst" { return "Burst" }
    if yname == "iburst" { return "Iburst" }
    return ""
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetSegmentPath() string {
    return "peer-type-ipv4" + "[peer-type='" + fmt.Sprintf("%v", peerTypeIpv4.PeerType) + "']"
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-type"] = peerTypeIpv4.PeerType
    leafs["ntp-version"] = peerTypeIpv4.NtpVersion
    leafs["authentication-key"] = peerTypeIpv4.AuthenticationKey
    leafs["min-poll"] = peerTypeIpv4.MinPoll
    leafs["max-poll"] = peerTypeIpv4.MaxPoll
    leafs["preferred-peer"] = peerTypeIpv4.PreferredPeer
    leafs["source-interface"] = peerTypeIpv4.SourceInterface
    leafs["burst"] = peerTypeIpv4.Burst
    leafs["iburst"] = peerTypeIpv4.Iburst
    return leafs
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetYangName() string { return "peer-type-ipv4" }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) SetParent(parent types.Entity) { peerTypeIpv4.parent = parent }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetParent() types.Entity { return peerTypeIpv4.parent }

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetParentYangName() string { return "peer-ipv4" }

// Ntp_PeerVrfs_PeerVrf_PeerIpv6S
// Configuration NTP Peers or Servers of IPV6
type Ntp_PeerVrfs_PeerVrf_PeerIpv6S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure a NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6.
    PeerIpv6 []Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetFilter() yfilter.YFilter { return peerIpv6S.YFilter }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) SetFilter(yf yfilter.YFilter) { peerIpv6S.YFilter = yf }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetGoName(yname string) string {
    if yname == "peer-ipv6" { return "PeerIpv6" }
    return ""
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetSegmentPath() string {
    return "peer-ipv6s"
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-ipv6" {
        for _, c := range peerIpv6S.PeerIpv6 {
            if peerIpv6S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6{}
        peerIpv6S.PeerIpv6 = append(peerIpv6S.PeerIpv6, child)
        return &peerIpv6S.PeerIpv6[len(peerIpv6S.PeerIpv6)-1]
    }
    return nil
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerIpv6S.PeerIpv6 {
        children[peerIpv6S.PeerIpv6[i].GetSegmentPath()] = &peerIpv6S.PeerIpv6[i]
    }
    return children
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetBundleName() string { return "cisco_ios_xr" }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetYangName() string { return "peer-ipv6s" }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) SetParent(parent types.Entity) { peerIpv6S.parent = parent }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetParent() types.Entity { return peerIpv6S.parent }

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetParentYangName() string { return "peer-vrf" }

// Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6
// Configure a NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of a peer. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    AddressIpv6 interface{}

    // Configure a NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6.
    PeerTypeIpv6 []Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetFilter() yfilter.YFilter { return peerIpv6.YFilter }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) SetFilter(yf yfilter.YFilter) { peerIpv6.YFilter = yf }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetGoName(yname string) string {
    if yname == "address-ipv6" { return "AddressIpv6" }
    if yname == "peer-type-ipv6" { return "PeerTypeIpv6" }
    return ""
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetSegmentPath() string {
    return "peer-ipv6" + "[address-ipv6='" + fmt.Sprintf("%v", peerIpv6.AddressIpv6) + "']"
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-type-ipv6" {
        for _, c := range peerIpv6.PeerTypeIpv6 {
            if peerIpv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6{}
        peerIpv6.PeerTypeIpv6 = append(peerIpv6.PeerTypeIpv6, child)
        return &peerIpv6.PeerTypeIpv6[len(peerIpv6.PeerTypeIpv6)-1]
    }
    return nil
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerIpv6.PeerTypeIpv6 {
        children[peerIpv6.PeerTypeIpv6[i].GetSegmentPath()] = &peerIpv6.PeerTypeIpv6[i]
    }
    return children
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-ipv6"] = peerIpv6.AddressIpv6
    return leafs
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetYangName() string { return "peer-ipv6" }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) SetParent(parent types.Entity) { peerIpv6.parent = parent }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetParent() types.Entity { return peerIpv6.parent }

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetParentYangName() string { return "peer-ipv6s" }

// Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6
// Configure a NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer or Server. The type is NtpPeer.
    PeerType interface{}

    // NTP version. The type is interface{} with range: 2..4.
    NtpVersion interface{}

    // Authentication Key. The type is interface{} with range: 1..65535.
    AuthenticationKey interface{}

    // Minimum poll interval. The type is interface{} with range: 4..17.
    MinPoll interface{}

    // Maxinum poll interval. The type is interface{} with range: 4..17.
    MaxPoll interface{}

    // Preferred peer. The type is interface{}.
    PreferredPeer interface{}

    // Source interface of this peer. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Use burst mode. The type is interface{}.
    Burst interface{}

    // Use iburst mode. The type is interface{}.
    Iburst interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    AddressIpv6 interface{}
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetFilter() yfilter.YFilter { return peerTypeIpv6.YFilter }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) SetFilter(yf yfilter.YFilter) { peerTypeIpv6.YFilter = yf }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetGoName(yname string) string {
    if yname == "peer-type" { return "PeerType" }
    if yname == "ntp-version" { return "NtpVersion" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    if yname == "min-poll" { return "MinPoll" }
    if yname == "max-poll" { return "MaxPoll" }
    if yname == "preferred-peer" { return "PreferredPeer" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "burst" { return "Burst" }
    if yname == "iburst" { return "Iburst" }
    if yname == "address-ipv6" { return "AddressIpv6" }
    return ""
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetSegmentPath() string {
    return "peer-type-ipv6" + "[peer-type='" + fmt.Sprintf("%v", peerTypeIpv6.PeerType) + "']"
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-type"] = peerTypeIpv6.PeerType
    leafs["ntp-version"] = peerTypeIpv6.NtpVersion
    leafs["authentication-key"] = peerTypeIpv6.AuthenticationKey
    leafs["min-poll"] = peerTypeIpv6.MinPoll
    leafs["max-poll"] = peerTypeIpv6.MaxPoll
    leafs["preferred-peer"] = peerTypeIpv6.PreferredPeer
    leafs["source-interface"] = peerTypeIpv6.SourceInterface
    leafs["burst"] = peerTypeIpv6.Burst
    leafs["iburst"] = peerTypeIpv6.Iburst
    leafs["address-ipv6"] = peerTypeIpv6.AddressIpv6
    return leafs
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetYangName() string { return "peer-type-ipv6" }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) SetParent(parent types.Entity) { peerTypeIpv6.parent = parent }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetParent() types.Entity { return peerTypeIpv6.parent }

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetParentYangName() string { return "peer-ipv6" }

// Ntp_DscpIpv4
//  Set IP DSCP value for outgoing NTP IPV4 packets
// This type is a presence type.
type Ntp_DscpIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NTPPRECEDENCE (0) to specify Precedence value  NTPDSCP (1) to specify DSCP
    // value. The type is Ntpdscp. This attribute is mandatory.
    Mode interface{}

    // If Mode is set to 'NTPPRECEDENCE(0)' specify Precedence value , if Mode is
    // set to 'NTPDSCP(1)' specify DSCP. The type is interface{} with range:
    // 0..63. This attribute is mandatory.
    DscpOrPrecedenceValue interface{}
}

func (dscpIpv4 *Ntp_DscpIpv4) GetFilter() yfilter.YFilter { return dscpIpv4.YFilter }

func (dscpIpv4 *Ntp_DscpIpv4) SetFilter(yf yfilter.YFilter) { dscpIpv4.YFilter = yf }

func (dscpIpv4 *Ntp_DscpIpv4) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "dscp-or-precedence-value" { return "DscpOrPrecedenceValue" }
    return ""
}

func (dscpIpv4 *Ntp_DscpIpv4) GetSegmentPath() string {
    return "dscp-ipv4"
}

func (dscpIpv4 *Ntp_DscpIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscpIpv4 *Ntp_DscpIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscpIpv4 *Ntp_DscpIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = dscpIpv4.Mode
    leafs["dscp-or-precedence-value"] = dscpIpv4.DscpOrPrecedenceValue
    return leafs
}

func (dscpIpv4 *Ntp_DscpIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (dscpIpv4 *Ntp_DscpIpv4) GetYangName() string { return "dscp-ipv4" }

func (dscpIpv4 *Ntp_DscpIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscpIpv4 *Ntp_DscpIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscpIpv4 *Ntp_DscpIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscpIpv4 *Ntp_DscpIpv4) SetParent(parent types.Entity) { dscpIpv4.parent = parent }

func (dscpIpv4 *Ntp_DscpIpv4) GetParent() types.Entity { return dscpIpv4.parent }

func (dscpIpv4 *Ntp_DscpIpv4) GetParentYangName() string { return "ntp" }

// Ntp_DscpIpv6
//  Set IP DSCP value for outgoing NTP IPV6 packets
// This type is a presence type.
type Ntp_DscpIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NTPPRECEDENCE(0) to specify Precedence value NTPDSCP(1) to specify DSCP
    // value. The type is Ntpdscp. This attribute is mandatory.
    Mode interface{}

    // If Mode is set to 'NTPPRECEDENCE(0)' specify Precedence value , if Mode is
    // set to 'NTPDSCP(1)' specify DSCP. The type is interface{} with range:
    // 0..63. This attribute is mandatory.
    DscpOrPrecedenceValue interface{}
}

func (dscpIpv6 *Ntp_DscpIpv6) GetFilter() yfilter.YFilter { return dscpIpv6.YFilter }

func (dscpIpv6 *Ntp_DscpIpv6) SetFilter(yf yfilter.YFilter) { dscpIpv6.YFilter = yf }

func (dscpIpv6 *Ntp_DscpIpv6) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "dscp-or-precedence-value" { return "DscpOrPrecedenceValue" }
    return ""
}

func (dscpIpv6 *Ntp_DscpIpv6) GetSegmentPath() string {
    return "dscp-ipv6"
}

func (dscpIpv6 *Ntp_DscpIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscpIpv6 *Ntp_DscpIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscpIpv6 *Ntp_DscpIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = dscpIpv6.Mode
    leafs["dscp-or-precedence-value"] = dscpIpv6.DscpOrPrecedenceValue
    return leafs
}

func (dscpIpv6 *Ntp_DscpIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (dscpIpv6 *Ntp_DscpIpv6) GetYangName() string { return "dscp-ipv6" }

func (dscpIpv6 *Ntp_DscpIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscpIpv6 *Ntp_DscpIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscpIpv6 *Ntp_DscpIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscpIpv6 *Ntp_DscpIpv6) SetParent(parent types.Entity) { dscpIpv6.parent = parent }

func (dscpIpv6 *Ntp_DscpIpv6) GetParent() types.Entity { return dscpIpv6.parent }

func (dscpIpv6 *Ntp_DscpIpv6) GetParentYangName() string { return "ntp" }

// Ntp_Sources
// Configure  NTP source interface
type Ntp_Sources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure  NTP source interface. The type is slice of Ntp_Sources_Source.
    Source []Ntp_Sources_Source
}

func (sources *Ntp_Sources) GetFilter() yfilter.YFilter { return sources.YFilter }

func (sources *Ntp_Sources) SetFilter(yf yfilter.YFilter) { sources.YFilter = yf }

func (sources *Ntp_Sources) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    return ""
}

func (sources *Ntp_Sources) GetSegmentPath() string {
    return "sources"
}

func (sources *Ntp_Sources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        for _, c := range sources.Source {
            if sources.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Sources_Source{}
        sources.Source = append(sources.Source, child)
        return &sources.Source[len(sources.Source)-1]
    }
    return nil
}

func (sources *Ntp_Sources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sources.Source {
        children[sources.Source[i].GetSegmentPath()] = &sources.Source[i]
    }
    return children
}

func (sources *Ntp_Sources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sources *Ntp_Sources) GetBundleName() string { return "cisco_ios_xr" }

func (sources *Ntp_Sources) GetYangName() string { return "sources" }

func (sources *Ntp_Sources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sources *Ntp_Sources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sources *Ntp_Sources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sources *Ntp_Sources) SetParent(parent types.Entity) { sources.parent = parent }

func (sources *Ntp_Sources) GetParent() types.Entity { return sources.parent }

func (sources *Ntp_Sources) GetParentYangName() string { return "ntp" }

// Ntp_Sources_Source
// Configure  NTP source interface
type Ntp_Sources_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Source Interface for NTP. The type is string with pattern: [a-zA-Z0-9./-]+.
    // This attribute is mandatory.
    SourceInterface interface{}
}

func (source *Ntp_Sources_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *Ntp_Sources_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *Ntp_Sources_Source) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (source *Ntp_Sources_Source) GetSegmentPath() string {
    return "source" + "[vrf-name='" + fmt.Sprintf("%v", source.VrfName) + "']"
}

func (source *Ntp_Sources_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *Ntp_Sources_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *Ntp_Sources_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = source.VrfName
    leafs["source-interface"] = source.SourceInterface
    return leafs
}

func (source *Ntp_Sources_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *Ntp_Sources_Source) GetYangName() string { return "source" }

func (source *Ntp_Sources_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *Ntp_Sources_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *Ntp_Sources_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *Ntp_Sources_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *Ntp_Sources_Source) GetParent() types.Entity { return source.parent }

func (source *Ntp_Sources_Source) GetParentYangName() string { return "sources" }

// Ntp_Authentication
// Configure NTP Authentication keys
type Ntp_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable NTP authentication keys. The type is interface{}.
    Enable interface{}

    // Authentication Key Table.
    Keies Ntp_Authentication_Keies

    // Key numbers for trusted time sources.
    TrustedKeies Ntp_Authentication_TrustedKeies
}

func (authentication *Ntp_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Ntp_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Ntp_Authentication) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "keies" { return "Keies" }
    if yname == "trusted-keies" { return "TrustedKeies" }
    return ""
}

func (authentication *Ntp_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Ntp_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "keies" {
        return &authentication.Keies
    }
    if childYangName == "trusted-keies" {
        return &authentication.TrustedKeies
    }
    return nil
}

func (authentication *Ntp_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["keies"] = &authentication.Keies
    children["trusted-keies"] = &authentication.TrustedKeies
    return children
}

func (authentication *Ntp_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = authentication.Enable
    return leafs
}

func (authentication *Ntp_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Ntp_Authentication) GetYangName() string { return "authentication" }

func (authentication *Ntp_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Ntp_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Ntp_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Ntp_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Ntp_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Ntp_Authentication) GetParentYangName() string { return "ntp" }

// Ntp_Authentication_Keies
// Authentication Key Table
type Ntp_Authentication_Keies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Authentication key for trusted time sources. The type is slice of
    // Ntp_Authentication_Keies_Key.
    Key []Ntp_Authentication_Keies_Key
}

func (keies *Ntp_Authentication_Keies) GetFilter() yfilter.YFilter { return keies.YFilter }

func (keies *Ntp_Authentication_Keies) SetFilter(yf yfilter.YFilter) { keies.YFilter = yf }

func (keies *Ntp_Authentication_Keies) GetGoName(yname string) string {
    if yname == "key" { return "Key" }
    return ""
}

func (keies *Ntp_Authentication_Keies) GetSegmentPath() string {
    return "keies"
}

func (keies *Ntp_Authentication_Keies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key" {
        for _, c := range keies.Key {
            if keies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Authentication_Keies_Key{}
        keies.Key = append(keies.Key, child)
        return &keies.Key[len(keies.Key)-1]
    }
    return nil
}

func (keies *Ntp_Authentication_Keies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range keies.Key {
        children[keies.Key[i].GetSegmentPath()] = &keies.Key[i]
    }
    return children
}

func (keies *Ntp_Authentication_Keies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keies *Ntp_Authentication_Keies) GetBundleName() string { return "cisco_ios_xr" }

func (keies *Ntp_Authentication_Keies) GetYangName() string { return "keies" }

func (keies *Ntp_Authentication_Keies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keies *Ntp_Authentication_Keies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keies *Ntp_Authentication_Keies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keies *Ntp_Authentication_Keies) SetParent(parent types.Entity) { keies.parent = parent }

func (keies *Ntp_Authentication_Keies) GetParent() types.Entity { return keies.parent }

func (keies *Ntp_Authentication_Keies) GetParentYangName() string { return "authentication" }

// Ntp_Authentication_Keies_Key
// Authentication key for trusted time sources
type Ntp_Authentication_Keies_Key struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Authentication Key number. The type is interface{}
    // with range: 1..65535.
    KeyNumber interface{}

    // Authentication key - maximum 32 characters. The type is string. This
    // attribute is mandatory.
    AuthenticationKey interface{}
}

func (key *Ntp_Authentication_Keies_Key) GetFilter() yfilter.YFilter { return key.YFilter }

func (key *Ntp_Authentication_Keies_Key) SetFilter(yf yfilter.YFilter) { key.YFilter = yf }

func (key *Ntp_Authentication_Keies_Key) GetGoName(yname string) string {
    if yname == "key-number" { return "KeyNumber" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    return ""
}

func (key *Ntp_Authentication_Keies_Key) GetSegmentPath() string {
    return "key" + "[key-number='" + fmt.Sprintf("%v", key.KeyNumber) + "']"
}

func (key *Ntp_Authentication_Keies_Key) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (key *Ntp_Authentication_Keies_Key) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (key *Ntp_Authentication_Keies_Key) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-number"] = key.KeyNumber
    leafs["authentication-key"] = key.AuthenticationKey
    return leafs
}

func (key *Ntp_Authentication_Keies_Key) GetBundleName() string { return "cisco_ios_xr" }

func (key *Ntp_Authentication_Keies_Key) GetYangName() string { return "key" }

func (key *Ntp_Authentication_Keies_Key) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (key *Ntp_Authentication_Keies_Key) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (key *Ntp_Authentication_Keies_Key) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (key *Ntp_Authentication_Keies_Key) SetParent(parent types.Entity) { key.parent = parent }

func (key *Ntp_Authentication_Keies_Key) GetParent() types.Entity { return key.parent }

func (key *Ntp_Authentication_Keies_Key) GetParentYangName() string { return "keies" }

// Ntp_Authentication_TrustedKeies
// Key numbers for trusted time sources
type Ntp_Authentication_TrustedKeies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure NTP trusted key. The type is slice of
    // Ntp_Authentication_TrustedKeies_TrustedKey.
    TrustedKey []Ntp_Authentication_TrustedKeies_TrustedKey
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetFilter() yfilter.YFilter { return trustedKeies.YFilter }

func (trustedKeies *Ntp_Authentication_TrustedKeies) SetFilter(yf yfilter.YFilter) { trustedKeies.YFilter = yf }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetGoName(yname string) string {
    if yname == "trusted-key" { return "TrustedKey" }
    return ""
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetSegmentPath() string {
    return "trusted-keies"
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trusted-key" {
        for _, c := range trustedKeies.TrustedKey {
            if trustedKeies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Authentication_TrustedKeies_TrustedKey{}
        trustedKeies.TrustedKey = append(trustedKeies.TrustedKey, child)
        return &trustedKeies.TrustedKey[len(trustedKeies.TrustedKey)-1]
    }
    return nil
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trustedKeies.TrustedKey {
        children[trustedKeies.TrustedKey[i].GetSegmentPath()] = &trustedKeies.TrustedKey[i]
    }
    return children
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetBundleName() string { return "cisco_ios_xr" }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetYangName() string { return "trusted-keies" }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trustedKeies *Ntp_Authentication_TrustedKeies) SetParent(parent types.Entity) { trustedKeies.parent = parent }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetParent() types.Entity { return trustedKeies.parent }

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetParentYangName() string { return "authentication" }

// Ntp_Authentication_TrustedKeies_TrustedKey
// Configure NTP trusted key
type Ntp_Authentication_TrustedKeies_TrustedKey struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Key number. The type is interface{} with range:
    // 1..65535.
    KeyNumber interface{}
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetFilter() yfilter.YFilter { return trustedKey.YFilter }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) SetFilter(yf yfilter.YFilter) { trustedKey.YFilter = yf }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetGoName(yname string) string {
    if yname == "key-number" { return "KeyNumber" }
    return ""
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetSegmentPath() string {
    return "trusted-key" + "[key-number='" + fmt.Sprintf("%v", trustedKey.KeyNumber) + "']"
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-number"] = trustedKey.KeyNumber
    return leafs
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetBundleName() string { return "cisco_ios_xr" }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetYangName() string { return "trusted-key" }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) SetParent(parent types.Entity) { trustedKey.parent = parent }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetParent() types.Entity { return trustedKey.parent }

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetParentYangName() string { return "trusted-keies" }

// Ntp_Passive
// Configure NTP passive associations
type Ntp_Passive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable NTP Passive associations. The type is interface{}.
    Enable interface{}
}

func (passive *Ntp_Passive) GetFilter() yfilter.YFilter { return passive.YFilter }

func (passive *Ntp_Passive) SetFilter(yf yfilter.YFilter) { passive.YFilter = yf }

func (passive *Ntp_Passive) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (passive *Ntp_Passive) GetSegmentPath() string {
    return "passive"
}

func (passive *Ntp_Passive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (passive *Ntp_Passive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (passive *Ntp_Passive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = passive.Enable
    return leafs
}

func (passive *Ntp_Passive) GetBundleName() string { return "cisco_ios_xr" }

func (passive *Ntp_Passive) GetYangName() string { return "passive" }

func (passive *Ntp_Passive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (passive *Ntp_Passive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (passive *Ntp_Passive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (passive *Ntp_Passive) SetParent(parent types.Entity) { passive.parent = parent }

func (passive *Ntp_Passive) GetParent() types.Entity { return passive.parent }

func (passive *Ntp_Passive) GetParentYangName() string { return "ntp" }

// Ntp_InterfaceTables
// NTP per interface configuration
type Ntp_InterfaceTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NTP per interface configuration. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable.
    InterfaceTable []Ntp_InterfaceTables_InterfaceTable
}

func (interfaceTables *Ntp_InterfaceTables) GetFilter() yfilter.YFilter { return interfaceTables.YFilter }

func (interfaceTables *Ntp_InterfaceTables) SetFilter(yf yfilter.YFilter) { interfaceTables.YFilter = yf }

func (interfaceTables *Ntp_InterfaceTables) GetGoName(yname string) string {
    if yname == "interface-table" { return "InterfaceTable" }
    return ""
}

func (interfaceTables *Ntp_InterfaceTables) GetSegmentPath() string {
    return "interface-tables"
}

func (interfaceTables *Ntp_InterfaceTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-table" {
        for _, c := range interfaceTables.InterfaceTable {
            if interfaceTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_InterfaceTables_InterfaceTable{}
        interfaceTables.InterfaceTable = append(interfaceTables.InterfaceTable, child)
        return &interfaceTables.InterfaceTable[len(interfaceTables.InterfaceTable)-1]
    }
    return nil
}

func (interfaceTables *Ntp_InterfaceTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceTables.InterfaceTable {
        children[interfaceTables.InterfaceTable[i].GetSegmentPath()] = &interfaceTables.InterfaceTable[i]
    }
    return children
}

func (interfaceTables *Ntp_InterfaceTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceTables *Ntp_InterfaceTables) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTables *Ntp_InterfaceTables) GetYangName() string { return "interface-tables" }

func (interfaceTables *Ntp_InterfaceTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTables *Ntp_InterfaceTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTables *Ntp_InterfaceTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTables *Ntp_InterfaceTables) SetParent(parent types.Entity) { interfaceTables.parent = parent }

func (interfaceTables *Ntp_InterfaceTables) GetParent() types.Entity { return interfaceTables.parent }

func (interfaceTables *Ntp_InterfaceTables) GetParentYangName() string { return "ntp" }

// Ntp_InterfaceTables_InterfaceTable
// NTP per interface configuration
type Ntp_InterfaceTables_InterfaceTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Name of the interface. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable_Interface.
    Interface []Ntp_InterfaceTables_InterfaceTable_Interface
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetFilter() yfilter.YFilter { return interfaceTable.YFilter }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) SetFilter(yf yfilter.YFilter) { interfaceTable.YFilter = yf }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetSegmentPath() string {
    return "interface-table" + "[vrf-name='" + fmt.Sprintf("%v", interfaceTable.VrfName) + "']"
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaceTable.Interface {
            if interfaceTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_InterfaceTables_InterfaceTable_Interface{}
        interfaceTable.Interface = append(interfaceTable.Interface, child)
        return &interfaceTable.Interface[len(interfaceTable.Interface)-1]
    }
    return nil
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceTable.Interface {
        children[interfaceTable.Interface[i].GetSegmentPath()] = &interfaceTable.Interface[i]
    }
    return children
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = interfaceTable.VrfName
    return leafs
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetYangName() string { return "interface-table" }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) SetParent(parent types.Entity) { interfaceTable.parent = parent }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetParent() types.Entity { return interfaceTable.parent }

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetParentYangName() string { return "interface-tables" }

// Ntp_InterfaceTables_InterfaceTable_Interface
// Name of the interface
type Ntp_InterfaceTables_InterfaceTable_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Disable NTP. The type is interface{}.
    Disable interface{}

    // Configure NTP multicast service.
    InterfaceMulticast Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast

    // Configure NTP broadcast service.
    InterfaceBroadcast Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "disable" { return "Disable" }
    if yname == "interface-multicast" { return "InterfaceMulticast" }
    if yname == "interface-broadcast" { return "InterfaceBroadcast" }
    return ""
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-multicast" {
        return &self.InterfaceMulticast
    }
    if childYangName == "interface-broadcast" {
        return &self.InterfaceBroadcast
    }
    return nil
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-multicast"] = &self.InterfaceMulticast
    children["interface-broadcast"] = &self.InterfaceBroadcast
    return children
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["disable"] = self.Disable
    return leafs
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetYangName() string { return "interface" }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetParent() types.Entity { return self.parent }

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetParentYangName() string { return "interface-table" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast
// Configure NTP multicast service
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configures multicast client peers.
    MulticastClients Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients

    // Configures multicast server peers.
    MulticastServers Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetFilter() yfilter.YFilter { return interfaceMulticast.YFilter }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) SetFilter(yf yfilter.YFilter) { interfaceMulticast.YFilter = yf }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetGoName(yname string) string {
    if yname == "multicast-clients" { return "MulticastClients" }
    if yname == "multicast-servers" { return "MulticastServers" }
    return ""
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetSegmentPath() string {
    return "interface-multicast"
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multicast-clients" {
        return &interfaceMulticast.MulticastClients
    }
    if childYangName == "multicast-servers" {
        return &interfaceMulticast.MulticastServers
    }
    return nil
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["multicast-clients"] = &interfaceMulticast.MulticastClients
    children["multicast-servers"] = &interfaceMulticast.MulticastServers
    return children
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetYangName() string { return "interface-multicast" }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) SetParent(parent types.Entity) { interfaceMulticast.parent = parent }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetParent() types.Entity { return interfaceMulticast.parent }

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetParentYangName() string { return "interface" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients
// Configures multicast client peers
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Listen to NTP multicasts. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient.
    MulticastClient []Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetFilter() yfilter.YFilter { return multicastClients.YFilter }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) SetFilter(yf yfilter.YFilter) { multicastClients.YFilter = yf }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetGoName(yname string) string {
    if yname == "multicast-client" { return "MulticastClient" }
    return ""
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetSegmentPath() string {
    return "multicast-clients"
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multicast-client" {
        for _, c := range multicastClients.MulticastClient {
            if multicastClients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient{}
        multicastClients.MulticastClient = append(multicastClients.MulticastClient, child)
        return &multicastClients.MulticastClient[len(multicastClients.MulticastClient)-1]
    }
    return nil
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multicastClients.MulticastClient {
        children[multicastClients.MulticastClient[i].GetSegmentPath()] = &multicastClients.MulticastClient[i]
    }
    return children
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetBundleName() string { return "cisco_ios_xr" }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetYangName() string { return "multicast-clients" }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) SetParent(parent types.Entity) { multicastClients.parent = parent }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetParent() types.Entity { return multicastClients.parent }

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetParentYangName() string { return "interface-multicast" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient
// Listen to NTP multicasts
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of a multicast group. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetFilter() yfilter.YFilter { return multicastClient.YFilter }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) SetFilter(yf yfilter.YFilter) { multicastClient.YFilter = yf }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetSegmentPath() string {
    return "multicast-client" + "[ip-address='" + fmt.Sprintf("%v", multicastClient.IpAddress) + "']"
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = multicastClient.IpAddress
    return leafs
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetBundleName() string { return "cisco_ios_xr" }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetYangName() string { return "multicast-client" }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) SetParent(parent types.Entity) { multicastClient.parent = parent }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetParent() types.Entity { return multicastClient.parent }

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetParentYangName() string { return "multicast-clients" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers
// Configures multicast server peers
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure NTP multicast group server peer. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer.
    MulticastServer []Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetFilter() yfilter.YFilter { return multicastServers.YFilter }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) SetFilter(yf yfilter.YFilter) { multicastServers.YFilter = yf }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetGoName(yname string) string {
    if yname == "multicast-server" { return "MulticastServer" }
    return ""
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetSegmentPath() string {
    return "multicast-servers"
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multicast-server" {
        for _, c := range multicastServers.MulticastServer {
            if multicastServers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer{}
        multicastServers.MulticastServer = append(multicastServers.MulticastServer, child)
        return &multicastServers.MulticastServer[len(multicastServers.MulticastServer)-1]
    }
    return nil
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multicastServers.MulticastServer {
        children[multicastServers.MulticastServer[i].GetSegmentPath()] = &multicastServers.MulticastServer[i]
    }
    return children
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetBundleName() string { return "cisco_ios_xr" }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetYangName() string { return "multicast-servers" }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) SetParent(parent types.Entity) { multicastServers.parent = parent }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetParent() types.Entity { return multicastServers.parent }

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetParentYangName() string { return "interface-multicast" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer
// Configure NTP multicast group server peer
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of a multicast group. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Authentication key. The type is interface{} with range: 1..65535.
    AuthenticationKey interface{}

    // NTP version. The type is interface{} with range: 2..4.
    Version interface{}

    // TTL. The type is interface{} with range: 1..255.
    Ttl interface{}
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetFilter() yfilter.YFilter { return multicastServer.YFilter }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) SetFilter(yf yfilter.YFilter) { multicastServer.YFilter = yf }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    if yname == "version" { return "Version" }
    if yname == "ttl" { return "Ttl" }
    return ""
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetSegmentPath() string {
    return "multicast-server" + "[ip-address='" + fmt.Sprintf("%v", multicastServer.IpAddress) + "']"
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = multicastServer.IpAddress
    leafs["authentication-key"] = multicastServer.AuthenticationKey
    leafs["version"] = multicastServer.Version
    leafs["ttl"] = multicastServer.Ttl
    return leafs
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetBundleName() string { return "cisco_ios_xr" }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetYangName() string { return "multicast-server" }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) SetParent(parent types.Entity) { multicastServer.parent = parent }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetParent() types.Entity { return multicastServer.parent }

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetParentYangName() string { return "multicast-servers" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast
// Configure NTP broadcast service
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Listen to NTP broadcasts. The type is interface{}.
    BroadcastClient interface{}

    // Configure NTP broadcast.
    Broadcast Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetFilter() yfilter.YFilter { return interfaceBroadcast.YFilter }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) SetFilter(yf yfilter.YFilter) { interfaceBroadcast.YFilter = yf }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetGoName(yname string) string {
    if yname == "broadcast-client" { return "BroadcastClient" }
    if yname == "broadcast" { return "Broadcast" }
    return ""
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetSegmentPath() string {
    return "interface-broadcast"
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "broadcast" {
        return &interfaceBroadcast.Broadcast
    }
    return nil
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["broadcast"] = &interfaceBroadcast.Broadcast
    return children
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["broadcast-client"] = interfaceBroadcast.BroadcastClient
    return leafs
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetYangName() string { return "interface-broadcast" }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) SetParent(parent types.Entity) { interfaceBroadcast.parent = parent }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetParent() types.Entity { return interfaceBroadcast.parent }

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetParentYangName() string { return "interface" }

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast
// Configure NTP broadcast
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination broadcast IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Authentication key. The type is interface{} with range: 1..65535.
    AuthenticationKey interface{}

    // NTP version. The type is interface{} with range: 2..4.
    NtpVersion interface{}
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetFilter() yfilter.YFilter { return broadcast.YFilter }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) SetFilter(yf yfilter.YFilter) { broadcast.YFilter = yf }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    if yname == "ntp-version" { return "NtpVersion" }
    return ""
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetSegmentPath() string {
    return "broadcast"
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = broadcast.Address
    leafs["authentication-key"] = broadcast.AuthenticationKey
    leafs["ntp-version"] = broadcast.NtpVersion
    return leafs
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetBundleName() string { return "cisco_ios_xr" }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetYangName() string { return "broadcast" }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) SetParent(parent types.Entity) { broadcast.parent = parent }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetParent() types.Entity { return broadcast.parent }

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetParentYangName() string { return "interface-broadcast" }

// Ntp_AccessGroupTables
// Control NTP access
type Ntp_AccessGroupTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control NTP access. The type is slice of
    // Ntp_AccessGroupTables_AccessGroupTable.
    AccessGroupTable []Ntp_AccessGroupTables_AccessGroupTable
}

func (accessGroupTables *Ntp_AccessGroupTables) GetFilter() yfilter.YFilter { return accessGroupTables.YFilter }

func (accessGroupTables *Ntp_AccessGroupTables) SetFilter(yf yfilter.YFilter) { accessGroupTables.YFilter = yf }

func (accessGroupTables *Ntp_AccessGroupTables) GetGoName(yname string) string {
    if yname == "access-group-table" { return "AccessGroupTable" }
    return ""
}

func (accessGroupTables *Ntp_AccessGroupTables) GetSegmentPath() string {
    return "access-group-tables"
}

func (accessGroupTables *Ntp_AccessGroupTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-group-table" {
        for _, c := range accessGroupTables.AccessGroupTable {
            if accessGroupTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_AccessGroupTables_AccessGroupTable{}
        accessGroupTables.AccessGroupTable = append(accessGroupTables.AccessGroupTable, child)
        return &accessGroupTables.AccessGroupTable[len(accessGroupTables.AccessGroupTable)-1]
    }
    return nil
}

func (accessGroupTables *Ntp_AccessGroupTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessGroupTables.AccessGroupTable {
        children[accessGroupTables.AccessGroupTable[i].GetSegmentPath()] = &accessGroupTables.AccessGroupTable[i]
    }
    return children
}

func (accessGroupTables *Ntp_AccessGroupTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessGroupTables *Ntp_AccessGroupTables) GetBundleName() string { return "cisco_ios_xr" }

func (accessGroupTables *Ntp_AccessGroupTables) GetYangName() string { return "access-group-tables" }

func (accessGroupTables *Ntp_AccessGroupTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessGroupTables *Ntp_AccessGroupTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessGroupTables *Ntp_AccessGroupTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessGroupTables *Ntp_AccessGroupTables) SetParent(parent types.Entity) { accessGroupTables.parent = parent }

func (accessGroupTables *Ntp_AccessGroupTables) GetParent() types.Entity { return accessGroupTables.parent }

func (accessGroupTables *Ntp_AccessGroupTables) GetParentYangName() string { return "ntp" }

// Ntp_AccessGroupTables_AccessGroupTable
// Control NTP access
type Ntp_AccessGroupTables_AccessGroupTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Configure NTP access address family. The type is slice of
    // Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable.
    AccessGroupAfTable []Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetFilter() yfilter.YFilter { return accessGroupTable.YFilter }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) SetFilter(yf yfilter.YFilter) { accessGroupTable.YFilter = yf }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "access-group-af-table" { return "AccessGroupAfTable" }
    return ""
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetSegmentPath() string {
    return "access-group-table" + "[vrf-name='" + fmt.Sprintf("%v", accessGroupTable.VrfName) + "']"
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-group-af-table" {
        for _, c := range accessGroupTable.AccessGroupAfTable {
            if accessGroupTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable{}
        accessGroupTable.AccessGroupAfTable = append(accessGroupTable.AccessGroupAfTable, child)
        return &accessGroupTable.AccessGroupAfTable[len(accessGroupTable.AccessGroupAfTable)-1]
    }
    return nil
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessGroupTable.AccessGroupAfTable {
        children[accessGroupTable.AccessGroupAfTable[i].GetSegmentPath()] = &accessGroupTable.AccessGroupAfTable[i]
    }
    return children
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = accessGroupTable.VrfName
    return leafs
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetBundleName() string { return "cisco_ios_xr" }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetYangName() string { return "access-group-table" }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) SetParent(parent types.Entity) { accessGroupTable.parent = parent }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetParent() types.Entity { return accessGroupTable.parent }

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetParentYangName() string { return "access-group-tables" }

// Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable
// Configure NTP access address family
type Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is NtpAccessAf.
    Af interface{}

    // Configure NTP access group. The type is slice of
    // Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup.
    AccessGroup []Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetFilter() yfilter.YFilter { return accessGroupAfTable.YFilter }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) SetFilter(yf yfilter.YFilter) { accessGroupAfTable.YFilter = yf }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "access-group" { return "AccessGroup" }
    return ""
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetSegmentPath() string {
    return "access-group-af-table" + "[af='" + fmt.Sprintf("%v", accessGroupAfTable.Af) + "']"
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-group" {
        for _, c := range accessGroupAfTable.AccessGroup {
            if accessGroupAfTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup{}
        accessGroupAfTable.AccessGroup = append(accessGroupAfTable.AccessGroup, child)
        return &accessGroupAfTable.AccessGroup[len(accessGroupAfTable.AccessGroup)-1]
    }
    return nil
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessGroupAfTable.AccessGroup {
        children[accessGroupAfTable.AccessGroup[i].GetSegmentPath()] = &accessGroupAfTable.AccessGroup[i]
    }
    return children
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = accessGroupAfTable.Af
    return leafs
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetBundleName() string { return "cisco_ios_xr" }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetYangName() string { return "access-group-af-table" }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) SetParent(parent types.Entity) { accessGroupAfTable.parent = parent }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetParent() types.Entity { return accessGroupAfTable.parent }

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetParentYangName() string { return "access-group-table" }

// Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup
// Configure NTP access group
type Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Access group type. The type is NtpAccess.
    AccessGroupType interface{}

    // Access list name - maximum 32 characters. The type is string. This
    // attribute is mandatory.
    AccessListName interface{}
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetFilter() yfilter.YFilter { return accessGroup.YFilter }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) SetFilter(yf yfilter.YFilter) { accessGroup.YFilter = yf }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetGoName(yname string) string {
    if yname == "access-group-type" { return "AccessGroupType" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetSegmentPath() string {
    return "access-group" + "[access-group-type='" + fmt.Sprintf("%v", accessGroup.AccessGroupType) + "']"
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-group-type"] = accessGroup.AccessGroupType
    leafs["access-list-name"] = accessGroup.AccessListName
    return leafs
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetBundleName() string { return "cisco_ios_xr" }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetYangName() string { return "access-group" }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) SetParent(parent types.Entity) { accessGroup.parent = parent }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetParent() types.Entity { return accessGroup.parent }

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetParentYangName() string { return "access-group-af-table" }

