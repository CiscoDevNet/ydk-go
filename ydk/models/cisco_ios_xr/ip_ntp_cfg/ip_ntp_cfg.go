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

// Ntpdscp represents Ntpdscp
type Ntpdscp string

const (
    // Precedence Value
    Ntpdscp_ntp_precedence Ntpdscp = "ntp-precedence"

    // DSCP Value
    Ntpdscp_ntpdscp Ntpdscp = "ntpdscp"
)

// NtpPeer represents Ntp peer
type NtpPeer string

const (
    // Peer
    NtpPeer_peer NtpPeer = "peer"

    // Server
    NtpPeer_server NtpPeer = "server"
)

// NtpAccessAf represents Ntp access af
type NtpAccessAf string

const (
    // IPv4
    NtpAccessAf_ipv4 NtpAccessAf = "ipv4"

    // IPv6
    NtpAccessAf_ipv6 NtpAccessAf = "ipv6"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set maximum number of associations. The type is interface{} with range:
    // 0..4294967295.
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

    // NTP drift.
    Drift Ntp_Drift

    // Configure NTP Authentication keys.
    Authentication Ntp_Authentication

    // Configure NTP passive associations.
    Passive Ntp_Passive

    // NTP per interface configuration.
    InterfaceTables Ntp_InterfaceTables

    // Control NTP access.
    AccessGroupTables Ntp_AccessGroupTables
}

func (ntp *Ntp) GetEntityData() *types.CommonEntityData {
    ntp.EntityData.YFilter = ntp.YFilter
    ntp.EntityData.YangName = "ntp"
    ntp.EntityData.BundleName = "cisco_ios_xr"
    ntp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-ntp-cfg"
    ntp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-ntp-cfg:ntp"
    ntp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntp.EntityData.Children = make(map[string]types.YChild)
    ntp.EntityData.Children["peer-vrfs"] = types.YChild{"PeerVrfs", &ntp.PeerVrfs}
    ntp.EntityData.Children["dscp-ipv4"] = types.YChild{"DscpIpv4", &ntp.DscpIpv4}
    ntp.EntityData.Children["dscp-ipv6"] = types.YChild{"DscpIpv6", &ntp.DscpIpv6}
    ntp.EntityData.Children["sources"] = types.YChild{"Sources", &ntp.Sources}
    ntp.EntityData.Children["drift"] = types.YChild{"Drift", &ntp.Drift}
    ntp.EntityData.Children["authentication"] = types.YChild{"Authentication", &ntp.Authentication}
    ntp.EntityData.Children["passive"] = types.YChild{"Passive", &ntp.Passive}
    ntp.EntityData.Children["interface-tables"] = types.YChild{"InterfaceTables", &ntp.InterfaceTables}
    ntp.EntityData.Children["access-group-tables"] = types.YChild{"AccessGroupTables", &ntp.AccessGroupTables}
    ntp.EntityData.Leafs = make(map[string]types.YLeaf)
    ntp.EntityData.Leafs["max-associations"] = types.YLeaf{"MaxAssociations", ntp.MaxAssociations}
    ntp.EntityData.Leafs["master"] = types.YLeaf{"Master", ntp.Master}
    ntp.EntityData.Leafs["broadcast-delay"] = types.YLeaf{"BroadcastDelay", ntp.BroadcastDelay}
    ntp.EntityData.Leafs["log-internal-sync"] = types.YLeaf{"LogInternalSync", ntp.LogInternalSync}
    ntp.EntityData.Leafs["update-calendar"] = types.YLeaf{"UpdateCalendar", ntp.UpdateCalendar}
    return &(ntp.EntityData)
}

// Ntp_PeerVrfs
// Configures NTP Peers or Servers
type Ntp_PeerVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configures NTP Peers or Servers for a single VRF. The 'default' must also
    // be specified for default VRF. The type is slice of Ntp_PeerVrfs_PeerVrf.
    PeerVrf []Ntp_PeerVrfs_PeerVrf
}

func (peerVrfs *Ntp_PeerVrfs) GetEntityData() *types.CommonEntityData {
    peerVrfs.EntityData.YFilter = peerVrfs.YFilter
    peerVrfs.EntityData.YangName = "peer-vrfs"
    peerVrfs.EntityData.BundleName = "cisco_ios_xr"
    peerVrfs.EntityData.ParentYangName = "ntp"
    peerVrfs.EntityData.SegmentPath = "peer-vrfs"
    peerVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerVrfs.EntityData.Children = make(map[string]types.YChild)
    peerVrfs.EntityData.Children["peer-vrf"] = types.YChild{"PeerVrf", nil}
    for i := range peerVrfs.PeerVrf {
        peerVrfs.EntityData.Children[types.GetSegmentPath(&peerVrfs.PeerVrf[i])] = types.YChild{"PeerVrf", &peerVrfs.PeerVrf[i]}
    }
    peerVrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerVrfs.EntityData)
}

// Ntp_PeerVrfs_PeerVrf
// Configures NTP Peers or Servers for a single
// VRF. The 'default' must also be specified for
// default VRF
type Ntp_PeerVrfs_PeerVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Configures IPv4 NTP Peers or Servers.
    PeerIpv4S Ntp_PeerVrfs_PeerVrf_PeerIpv4S

    // Configuration NTP Peers or Servers of IPV6.
    PeerIpv6S Ntp_PeerVrfs_PeerVrf_PeerIpv6S
}

func (peerVrf *Ntp_PeerVrfs_PeerVrf) GetEntityData() *types.CommonEntityData {
    peerVrf.EntityData.YFilter = peerVrf.YFilter
    peerVrf.EntityData.YangName = "peer-vrf"
    peerVrf.EntityData.BundleName = "cisco_ios_xr"
    peerVrf.EntityData.ParentYangName = "peer-vrfs"
    peerVrf.EntityData.SegmentPath = "peer-vrf" + "[vrf-name='" + fmt.Sprintf("%v", peerVrf.VrfName) + "']"
    peerVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerVrf.EntityData.Children = make(map[string]types.YChild)
    peerVrf.EntityData.Children["peer-ipv4s"] = types.YChild{"PeerIpv4S", &peerVrf.PeerIpv4S}
    peerVrf.EntityData.Children["peer-ipv6s"] = types.YChild{"PeerIpv6S", &peerVrf.PeerIpv6S}
    peerVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    peerVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", peerVrf.VrfName}
    return &(peerVrf.EntityData)
}

// Ntp_PeerVrfs_PeerVrf_PeerIpv4S
// Configures IPv4 NTP Peers or Servers
type Ntp_PeerVrfs_PeerVrf_PeerIpv4S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure an IPv4 NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4.
    PeerIpv4 []Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4
}

func (peerIpv4S *Ntp_PeerVrfs_PeerVrf_PeerIpv4S) GetEntityData() *types.CommonEntityData {
    peerIpv4S.EntityData.YFilter = peerIpv4S.YFilter
    peerIpv4S.EntityData.YangName = "peer-ipv4s"
    peerIpv4S.EntityData.BundleName = "cisco_ios_xr"
    peerIpv4S.EntityData.ParentYangName = "peer-vrf"
    peerIpv4S.EntityData.SegmentPath = "peer-ipv4s"
    peerIpv4S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerIpv4S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerIpv4S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerIpv4S.EntityData.Children = make(map[string]types.YChild)
    peerIpv4S.EntityData.Children["peer-ipv4"] = types.YChild{"PeerIpv4", nil}
    for i := range peerIpv4S.PeerIpv4 {
        peerIpv4S.EntityData.Children[types.GetSegmentPath(&peerIpv4S.PeerIpv4[i])] = types.YChild{"PeerIpv4", &peerIpv4S.PeerIpv4[i]}
    }
    peerIpv4S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerIpv4S.EntityData)
}

// Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4
// Configure an IPv4 NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 Address of a peer. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AddressIpv4 interface{}

    // Configure an IPv4 NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4.
    PeerTypeIpv4 []Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4
}

func (peerIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4) GetEntityData() *types.CommonEntityData {
    peerIpv4.EntityData.YFilter = peerIpv4.YFilter
    peerIpv4.EntityData.YangName = "peer-ipv4"
    peerIpv4.EntityData.BundleName = "cisco_ios_xr"
    peerIpv4.EntityData.ParentYangName = "peer-ipv4s"
    peerIpv4.EntityData.SegmentPath = "peer-ipv4" + "[address-ipv4='" + fmt.Sprintf("%v", peerIpv4.AddressIpv4) + "']"
    peerIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerIpv4.EntityData.Children = make(map[string]types.YChild)
    peerIpv4.EntityData.Children["peer-type-ipv4"] = types.YChild{"PeerTypeIpv4", nil}
    for i := range peerIpv4.PeerTypeIpv4 {
        peerIpv4.EntityData.Children[types.GetSegmentPath(&peerIpv4.PeerTypeIpv4[i])] = types.YChild{"PeerTypeIpv4", &peerIpv4.PeerTypeIpv4[i]}
    }
    peerIpv4.EntityData.Leafs = make(map[string]types.YLeaf)
    peerIpv4.EntityData.Leafs["address-ipv4"] = types.YLeaf{"AddressIpv4", peerIpv4.AddressIpv4}
    return &(peerIpv4.EntityData)
}

// Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4
// Configure an IPv4 NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4 struct {
    EntityData types.CommonEntityData
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
    // b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Use burst mode. The type is interface{}.
    Burst interface{}

    // Use iburst mode. The type is interface{}.
    Iburst interface{}
}

func (peerTypeIpv4 *Ntp_PeerVrfs_PeerVrf_PeerIpv4S_PeerIpv4_PeerTypeIpv4) GetEntityData() *types.CommonEntityData {
    peerTypeIpv4.EntityData.YFilter = peerTypeIpv4.YFilter
    peerTypeIpv4.EntityData.YangName = "peer-type-ipv4"
    peerTypeIpv4.EntityData.BundleName = "cisco_ios_xr"
    peerTypeIpv4.EntityData.ParentYangName = "peer-ipv4"
    peerTypeIpv4.EntityData.SegmentPath = "peer-type-ipv4" + "[peer-type='" + fmt.Sprintf("%v", peerTypeIpv4.PeerType) + "']"
    peerTypeIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerTypeIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerTypeIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerTypeIpv4.EntityData.Children = make(map[string]types.YChild)
    peerTypeIpv4.EntityData.Leafs = make(map[string]types.YLeaf)
    peerTypeIpv4.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", peerTypeIpv4.PeerType}
    peerTypeIpv4.EntityData.Leafs["ntp-version"] = types.YLeaf{"NtpVersion", peerTypeIpv4.NtpVersion}
    peerTypeIpv4.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", peerTypeIpv4.AuthenticationKey}
    peerTypeIpv4.EntityData.Leafs["min-poll"] = types.YLeaf{"MinPoll", peerTypeIpv4.MinPoll}
    peerTypeIpv4.EntityData.Leafs["max-poll"] = types.YLeaf{"MaxPoll", peerTypeIpv4.MaxPoll}
    peerTypeIpv4.EntityData.Leafs["preferred-peer"] = types.YLeaf{"PreferredPeer", peerTypeIpv4.PreferredPeer}
    peerTypeIpv4.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", peerTypeIpv4.SourceInterface}
    peerTypeIpv4.EntityData.Leafs["burst"] = types.YLeaf{"Burst", peerTypeIpv4.Burst}
    peerTypeIpv4.EntityData.Leafs["iburst"] = types.YLeaf{"Iburst", peerTypeIpv4.Iburst}
    return &(peerTypeIpv4.EntityData)
}

// Ntp_PeerVrfs_PeerVrf_PeerIpv6S
// Configuration NTP Peers or Servers of IPV6
type Ntp_PeerVrfs_PeerVrf_PeerIpv6S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6.
    PeerIpv6 []Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6
}

func (peerIpv6S *Ntp_PeerVrfs_PeerVrf_PeerIpv6S) GetEntityData() *types.CommonEntityData {
    peerIpv6S.EntityData.YFilter = peerIpv6S.YFilter
    peerIpv6S.EntityData.YangName = "peer-ipv6s"
    peerIpv6S.EntityData.BundleName = "cisco_ios_xr"
    peerIpv6S.EntityData.ParentYangName = "peer-vrf"
    peerIpv6S.EntityData.SegmentPath = "peer-ipv6s"
    peerIpv6S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerIpv6S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerIpv6S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerIpv6S.EntityData.Children = make(map[string]types.YChild)
    peerIpv6S.EntityData.Children["peer-ipv6"] = types.YChild{"PeerIpv6", nil}
    for i := range peerIpv6S.PeerIpv6 {
        peerIpv6S.EntityData.Children[types.GetSegmentPath(&peerIpv6S.PeerIpv6[i])] = types.YChild{"PeerIpv6", &peerIpv6S.PeerIpv6[i]}
    }
    peerIpv6S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerIpv6S.EntityData)
}

// Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6
// Configure a NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address of a peer. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AddressIpv6 interface{}

    // Configure a NTP server or peer. The type is slice of
    // Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6.
    PeerTypeIpv6 []Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6
}

func (peerIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6) GetEntityData() *types.CommonEntityData {
    peerIpv6.EntityData.YFilter = peerIpv6.YFilter
    peerIpv6.EntityData.YangName = "peer-ipv6"
    peerIpv6.EntityData.BundleName = "cisco_ios_xr"
    peerIpv6.EntityData.ParentYangName = "peer-ipv6s"
    peerIpv6.EntityData.SegmentPath = "peer-ipv6" + "[address-ipv6='" + fmt.Sprintf("%v", peerIpv6.AddressIpv6) + "']"
    peerIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerIpv6.EntityData.Children = make(map[string]types.YChild)
    peerIpv6.EntityData.Children["peer-type-ipv6"] = types.YChild{"PeerTypeIpv6", nil}
    for i := range peerIpv6.PeerTypeIpv6 {
        peerIpv6.EntityData.Children[types.GetSegmentPath(&peerIpv6.PeerTypeIpv6[i])] = types.YChild{"PeerTypeIpv6", &peerIpv6.PeerTypeIpv6[i]}
    }
    peerIpv6.EntityData.Leafs = make(map[string]types.YLeaf)
    peerIpv6.EntityData.Leafs["address-ipv6"] = types.YLeaf{"AddressIpv6", peerIpv6.AddressIpv6}
    return &(peerIpv6.EntityData)
}

// Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6
// Configure a NTP server or peer
type Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6 struct {
    EntityData types.CommonEntityData
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
    // b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Use burst mode. The type is interface{}.
    Burst interface{}

    // Use iburst mode. The type is interface{}.
    Iburst interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AddressIpv6 interface{}
}

func (peerTypeIpv6 *Ntp_PeerVrfs_PeerVrf_PeerIpv6S_PeerIpv6_PeerTypeIpv6) GetEntityData() *types.CommonEntityData {
    peerTypeIpv6.EntityData.YFilter = peerTypeIpv6.YFilter
    peerTypeIpv6.EntityData.YangName = "peer-type-ipv6"
    peerTypeIpv6.EntityData.BundleName = "cisco_ios_xr"
    peerTypeIpv6.EntityData.ParentYangName = "peer-ipv6"
    peerTypeIpv6.EntityData.SegmentPath = "peer-type-ipv6" + "[peer-type='" + fmt.Sprintf("%v", peerTypeIpv6.PeerType) + "']"
    peerTypeIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerTypeIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerTypeIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerTypeIpv6.EntityData.Children = make(map[string]types.YChild)
    peerTypeIpv6.EntityData.Leafs = make(map[string]types.YLeaf)
    peerTypeIpv6.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", peerTypeIpv6.PeerType}
    peerTypeIpv6.EntityData.Leafs["ntp-version"] = types.YLeaf{"NtpVersion", peerTypeIpv6.NtpVersion}
    peerTypeIpv6.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", peerTypeIpv6.AuthenticationKey}
    peerTypeIpv6.EntityData.Leafs["min-poll"] = types.YLeaf{"MinPoll", peerTypeIpv6.MinPoll}
    peerTypeIpv6.EntityData.Leafs["max-poll"] = types.YLeaf{"MaxPoll", peerTypeIpv6.MaxPoll}
    peerTypeIpv6.EntityData.Leafs["preferred-peer"] = types.YLeaf{"PreferredPeer", peerTypeIpv6.PreferredPeer}
    peerTypeIpv6.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", peerTypeIpv6.SourceInterface}
    peerTypeIpv6.EntityData.Leafs["burst"] = types.YLeaf{"Burst", peerTypeIpv6.Burst}
    peerTypeIpv6.EntityData.Leafs["iburst"] = types.YLeaf{"Iburst", peerTypeIpv6.Iburst}
    peerTypeIpv6.EntityData.Leafs["address-ipv6"] = types.YLeaf{"AddressIpv6", peerTypeIpv6.AddressIpv6}
    return &(peerTypeIpv6.EntityData)
}

// Ntp_DscpIpv4
//  Set IP DSCP value for outgoing NTP IPV4 packets
// This type is a presence type.
type Ntp_DscpIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTPPRECEDENCE (0) to specify Precedence value  NTPDSCP (1) to specify DSCP
    // value. The type is Ntpdscp. This attribute is mandatory.
    Mode interface{}

    // If Mode is set to 'NTPPRECEDENCE(0)' specify Precedence value , if Mode is
    // set to 'NTPDSCP(1)' specify DSCP. The type is interface{} with range:
    // 0..63. This attribute is mandatory.
    DscpOrPrecedenceValue interface{}
}

func (dscpIpv4 *Ntp_DscpIpv4) GetEntityData() *types.CommonEntityData {
    dscpIpv4.EntityData.YFilter = dscpIpv4.YFilter
    dscpIpv4.EntityData.YangName = "dscp-ipv4"
    dscpIpv4.EntityData.BundleName = "cisco_ios_xr"
    dscpIpv4.EntityData.ParentYangName = "ntp"
    dscpIpv4.EntityData.SegmentPath = "dscp-ipv4"
    dscpIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscpIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscpIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscpIpv4.EntityData.Children = make(map[string]types.YChild)
    dscpIpv4.EntityData.Leafs = make(map[string]types.YLeaf)
    dscpIpv4.EntityData.Leafs["mode"] = types.YLeaf{"Mode", dscpIpv4.Mode}
    dscpIpv4.EntityData.Leafs["dscp-or-precedence-value"] = types.YLeaf{"DscpOrPrecedenceValue", dscpIpv4.DscpOrPrecedenceValue}
    return &(dscpIpv4.EntityData)
}

// Ntp_DscpIpv6
//  Set IP DSCP value for outgoing NTP IPV6 packets
// This type is a presence type.
type Ntp_DscpIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTPPRECEDENCE(0) to specify Precedence value NTPDSCP(1) to specify DSCP
    // value. The type is Ntpdscp. This attribute is mandatory.
    Mode interface{}

    // If Mode is set to 'NTPPRECEDENCE(0)' specify Precedence value , if Mode is
    // set to 'NTPDSCP(1)' specify DSCP. The type is interface{} with range:
    // 0..63. This attribute is mandatory.
    DscpOrPrecedenceValue interface{}
}

func (dscpIpv6 *Ntp_DscpIpv6) GetEntityData() *types.CommonEntityData {
    dscpIpv6.EntityData.YFilter = dscpIpv6.YFilter
    dscpIpv6.EntityData.YangName = "dscp-ipv6"
    dscpIpv6.EntityData.BundleName = "cisco_ios_xr"
    dscpIpv6.EntityData.ParentYangName = "ntp"
    dscpIpv6.EntityData.SegmentPath = "dscp-ipv6"
    dscpIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscpIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscpIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscpIpv6.EntityData.Children = make(map[string]types.YChild)
    dscpIpv6.EntityData.Leafs = make(map[string]types.YLeaf)
    dscpIpv6.EntityData.Leafs["mode"] = types.YLeaf{"Mode", dscpIpv6.Mode}
    dscpIpv6.EntityData.Leafs["dscp-or-precedence-value"] = types.YLeaf{"DscpOrPrecedenceValue", dscpIpv6.DscpOrPrecedenceValue}
    return &(dscpIpv6.EntityData)
}

// Ntp_Sources
// Configure  NTP source interface
type Ntp_Sources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure  NTP source interface. The type is slice of Ntp_Sources_Source.
    Source []Ntp_Sources_Source
}

func (sources *Ntp_Sources) GetEntityData() *types.CommonEntityData {
    sources.EntityData.YFilter = sources.YFilter
    sources.EntityData.YangName = "sources"
    sources.EntityData.BundleName = "cisco_ios_xr"
    sources.EntityData.ParentYangName = "ntp"
    sources.EntityData.SegmentPath = "sources"
    sources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sources.EntityData.Children = make(map[string]types.YChild)
    sources.EntityData.Children["source"] = types.YChild{"Source", nil}
    for i := range sources.Source {
        sources.EntityData.Children[types.GetSegmentPath(&sources.Source[i])] = types.YChild{"Source", &sources.Source[i]}
    }
    sources.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sources.EntityData)
}

// Ntp_Sources_Source
// Configure  NTP source interface
type Ntp_Sources_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Source Interface for NTP. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'. This attribute is mandatory.
    SourceInterface interface{}
}

func (source *Ntp_Sources_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "sources"
    source.EntityData.SegmentPath = "source" + "[vrf-name='" + fmt.Sprintf("%v", source.VrfName) + "']"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", source.VrfName}
    source.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", source.SourceInterface}
    return &(source.EntityData)
}

// Ntp_Drift
// NTP drift
type Ntp_Drift struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Drift Aging Time. The type is interface{} with range: 0..65535.
    AgingTime interface{}

    // File containing drift value.
    File Ntp_Drift_File
}

func (drift *Ntp_Drift) GetEntityData() *types.CommonEntityData {
    drift.EntityData.YFilter = drift.YFilter
    drift.EntityData.YangName = "drift"
    drift.EntityData.BundleName = "cisco_ios_xr"
    drift.EntityData.ParentYangName = "ntp"
    drift.EntityData.SegmentPath = "drift"
    drift.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drift.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drift.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drift.EntityData.Children = make(map[string]types.YChild)
    drift.EntityData.Children["file"] = types.YChild{"File", &drift.File}
    drift.EntityData.Leafs = make(map[string]types.YLeaf)
    drift.EntityData.Leafs["aging-time"] = types.YLeaf{"AgingTime", drift.AgingTime}
    return &(drift.EntityData)
}

// Ntp_Drift_File
// File containing drift value
type Ntp_Drift_File struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PWD or disk0 etc. The type is string. The default value is PWD.
    Location interface{}

    // File containing drift value. The type is string.
    Filename interface{}
}

func (file *Ntp_Drift_File) GetEntityData() *types.CommonEntityData {
    file.EntityData.YFilter = file.YFilter
    file.EntityData.YangName = "file"
    file.EntityData.BundleName = "cisco_ios_xr"
    file.EntityData.ParentYangName = "drift"
    file.EntityData.SegmentPath = "file"
    file.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    file.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    file.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    file.EntityData.Children = make(map[string]types.YChild)
    file.EntityData.Leafs = make(map[string]types.YLeaf)
    file.EntityData.Leafs["location"] = types.YLeaf{"Location", file.Location}
    file.EntityData.Leafs["filename"] = types.YLeaf{"Filename", file.Filename}
    return &(file.EntityData)
}

// Ntp_Authentication
// Configure NTP Authentication keys
type Ntp_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable NTP authentication keys. The type is interface{}.
    Enable interface{}

    // Authentication Key Table.
    Keies Ntp_Authentication_Keies

    // Key numbers for trusted time sources.
    TrustedKeies Ntp_Authentication_TrustedKeies
}

func (authentication *Ntp_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "ntp"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["keies"] = types.YChild{"Keies", &authentication.Keies}
    authentication.EntityData.Children["trusted-keies"] = types.YChild{"TrustedKeies", &authentication.TrustedKeies}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["enable"] = types.YLeaf{"Enable", authentication.Enable}
    return &(authentication.EntityData)
}

// Ntp_Authentication_Keies
// Authentication Key Table
type Ntp_Authentication_Keies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authentication key for trusted time sources. The type is slice of
    // Ntp_Authentication_Keies_Key.
    Key []Ntp_Authentication_Keies_Key
}

func (keies *Ntp_Authentication_Keies) GetEntityData() *types.CommonEntityData {
    keies.EntityData.YFilter = keies.YFilter
    keies.EntityData.YangName = "keies"
    keies.EntityData.BundleName = "cisco_ios_xr"
    keies.EntityData.ParentYangName = "authentication"
    keies.EntityData.SegmentPath = "keies"
    keies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keies.EntityData.Children = make(map[string]types.YChild)
    keies.EntityData.Children["key"] = types.YChild{"Key", nil}
    for i := range keies.Key {
        keies.EntityData.Children[types.GetSegmentPath(&keies.Key[i])] = types.YChild{"Key", &keies.Key[i]}
    }
    keies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(keies.EntityData)
}

// Ntp_Authentication_Keies_Key
// Authentication key for trusted time sources
type Ntp_Authentication_Keies_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Authentication Key number. The type is interface{}
    // with range: 1..65535.
    KeyNumber interface{}

    // Authentication key - maximum 32 characters. The type is string. This
    // attribute is mandatory.
    AuthenticationKey interface{}
}

func (key *Ntp_Authentication_Keies_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "keies"
    key.EntityData.SegmentPath = "key" + "[key-number='" + fmt.Sprintf("%v", key.KeyNumber) + "']"
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = make(map[string]types.YChild)
    key.EntityData.Leafs = make(map[string]types.YLeaf)
    key.EntityData.Leafs["key-number"] = types.YLeaf{"KeyNumber", key.KeyNumber}
    key.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", key.AuthenticationKey}
    return &(key.EntityData)
}

// Ntp_Authentication_TrustedKeies
// Key numbers for trusted time sources
type Ntp_Authentication_TrustedKeies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure NTP trusted key. The type is slice of
    // Ntp_Authentication_TrustedKeies_TrustedKey.
    TrustedKey []Ntp_Authentication_TrustedKeies_TrustedKey
}

func (trustedKeies *Ntp_Authentication_TrustedKeies) GetEntityData() *types.CommonEntityData {
    trustedKeies.EntityData.YFilter = trustedKeies.YFilter
    trustedKeies.EntityData.YangName = "trusted-keies"
    trustedKeies.EntityData.BundleName = "cisco_ios_xr"
    trustedKeies.EntityData.ParentYangName = "authentication"
    trustedKeies.EntityData.SegmentPath = "trusted-keies"
    trustedKeies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trustedKeies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trustedKeies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trustedKeies.EntityData.Children = make(map[string]types.YChild)
    trustedKeies.EntityData.Children["trusted-key"] = types.YChild{"TrustedKey", nil}
    for i := range trustedKeies.TrustedKey {
        trustedKeies.EntityData.Children[types.GetSegmentPath(&trustedKeies.TrustedKey[i])] = types.YChild{"TrustedKey", &trustedKeies.TrustedKey[i]}
    }
    trustedKeies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trustedKeies.EntityData)
}

// Ntp_Authentication_TrustedKeies_TrustedKey
// Configure NTP trusted key
type Ntp_Authentication_TrustedKeies_TrustedKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Key number. The type is interface{} with range:
    // 1..65535.
    KeyNumber interface{}
}

func (trustedKey *Ntp_Authentication_TrustedKeies_TrustedKey) GetEntityData() *types.CommonEntityData {
    trustedKey.EntityData.YFilter = trustedKey.YFilter
    trustedKey.EntityData.YangName = "trusted-key"
    trustedKey.EntityData.BundleName = "cisco_ios_xr"
    trustedKey.EntityData.ParentYangName = "trusted-keies"
    trustedKey.EntityData.SegmentPath = "trusted-key" + "[key-number='" + fmt.Sprintf("%v", trustedKey.KeyNumber) + "']"
    trustedKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trustedKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trustedKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trustedKey.EntityData.Children = make(map[string]types.YChild)
    trustedKey.EntityData.Leafs = make(map[string]types.YLeaf)
    trustedKey.EntityData.Leafs["key-number"] = types.YLeaf{"KeyNumber", trustedKey.KeyNumber}
    return &(trustedKey.EntityData)
}

// Ntp_Passive
// Configure NTP passive associations
type Ntp_Passive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable NTP Passive associations. The type is interface{}.
    Enable interface{}
}

func (passive *Ntp_Passive) GetEntityData() *types.CommonEntityData {
    passive.EntityData.YFilter = passive.YFilter
    passive.EntityData.YangName = "passive"
    passive.EntityData.BundleName = "cisco_ios_xr"
    passive.EntityData.ParentYangName = "ntp"
    passive.EntityData.SegmentPath = "passive"
    passive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passive.EntityData.Children = make(map[string]types.YChild)
    passive.EntityData.Leafs = make(map[string]types.YLeaf)
    passive.EntityData.Leafs["enable"] = types.YLeaf{"Enable", passive.Enable}
    return &(passive.EntityData)
}

// Ntp_InterfaceTables
// NTP per interface configuration
type Ntp_InterfaceTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP per interface configuration. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable.
    InterfaceTable []Ntp_InterfaceTables_InterfaceTable
}

func (interfaceTables *Ntp_InterfaceTables) GetEntityData() *types.CommonEntityData {
    interfaceTables.EntityData.YFilter = interfaceTables.YFilter
    interfaceTables.EntityData.YangName = "interface-tables"
    interfaceTables.EntityData.BundleName = "cisco_ios_xr"
    interfaceTables.EntityData.ParentYangName = "ntp"
    interfaceTables.EntityData.SegmentPath = "interface-tables"
    interfaceTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTables.EntityData.Children = make(map[string]types.YChild)
    interfaceTables.EntityData.Children["interface-table"] = types.YChild{"InterfaceTable", nil}
    for i := range interfaceTables.InterfaceTable {
        interfaceTables.EntityData.Children[types.GetSegmentPath(&interfaceTables.InterfaceTable[i])] = types.YChild{"InterfaceTable", &interfaceTables.InterfaceTable[i]}
    }
    interfaceTables.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceTables.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable
// NTP per interface configuration
type Ntp_InterfaceTables_InterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Name of the interface. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable_Interface_.
    Interface_ []Ntp_InterfaceTables_InterfaceTable_Interface
}

func (interfaceTable *Ntp_InterfaceTables_InterfaceTable) GetEntityData() *types.CommonEntityData {
    interfaceTable.EntityData.YFilter = interfaceTable.YFilter
    interfaceTable.EntityData.YangName = "interface-table"
    interfaceTable.EntityData.BundleName = "cisco_ios_xr"
    interfaceTable.EntityData.ParentYangName = "interface-tables"
    interfaceTable.EntityData.SegmentPath = "interface-table" + "[vrf-name='" + fmt.Sprintf("%v", interfaceTable.VrfName) + "']"
    interfaceTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTable.EntityData.Children = make(map[string]types.YChild)
    interfaceTable.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaceTable.Interface_ {
        interfaceTable.EntityData.Children[types.GetSegmentPath(&interfaceTable.Interface_[i])] = types.YChild{"Interface_", &interfaceTable.Interface_[i]}
    }
    interfaceTable.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTable.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", interfaceTable.VrfName}
    return &(interfaceTable.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface
// Name of the interface
type Ntp_InterfaceTables_InterfaceTable_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Disable NTP. The type is interface{}.
    Disable interface{}

    // Configure NTP multicast service.
    InterfaceMulticast Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast

    // Configure NTP broadcast service.
    InterfaceBroadcast Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast
}

func (self *Ntp_InterfaceTables_InterfaceTable_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-table"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["interface-multicast"] = types.YChild{"InterfaceMulticast", &self.InterfaceMulticast}
    self.EntityData.Children["interface-broadcast"] = types.YChild{"InterfaceBroadcast", &self.InterfaceBroadcast}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["disable"] = types.YLeaf{"Disable", self.Disable}
    return &(self.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast
// Configure NTP multicast service
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configures multicast client peers.
    MulticastClients Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients

    // Configures multicast server peers.
    MulticastServers Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers
}

func (interfaceMulticast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast) GetEntityData() *types.CommonEntityData {
    interfaceMulticast.EntityData.YFilter = interfaceMulticast.YFilter
    interfaceMulticast.EntityData.YangName = "interface-multicast"
    interfaceMulticast.EntityData.BundleName = "cisco_ios_xr"
    interfaceMulticast.EntityData.ParentYangName = "interface"
    interfaceMulticast.EntityData.SegmentPath = "interface-multicast"
    interfaceMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMulticast.EntityData.Children = make(map[string]types.YChild)
    interfaceMulticast.EntityData.Children["multicast-clients"] = types.YChild{"MulticastClients", &interfaceMulticast.MulticastClients}
    interfaceMulticast.EntityData.Children["multicast-servers"] = types.YChild{"MulticastServers", &interfaceMulticast.MulticastServers}
    interfaceMulticast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceMulticast.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients
// Configures multicast client peers
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Listen to NTP multicasts. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient.
    MulticastClient []Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient
}

func (multicastClients *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients) GetEntityData() *types.CommonEntityData {
    multicastClients.EntityData.YFilter = multicastClients.YFilter
    multicastClients.EntityData.YangName = "multicast-clients"
    multicastClients.EntityData.BundleName = "cisco_ios_xr"
    multicastClients.EntityData.ParentYangName = "interface-multicast"
    multicastClients.EntityData.SegmentPath = "multicast-clients"
    multicastClients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastClients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastClients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastClients.EntityData.Children = make(map[string]types.YChild)
    multicastClients.EntityData.Children["multicast-client"] = types.YChild{"MulticastClient", nil}
    for i := range multicastClients.MulticastClient {
        multicastClients.EntityData.Children[types.GetSegmentPath(&multicastClients.MulticastClient[i])] = types.YChild{"MulticastClient", &multicastClients.MulticastClient[i]}
    }
    multicastClients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multicastClients.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient
// Listen to NTP multicasts
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of a multicast group. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}
}

func (multicastClient *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastClients_MulticastClient) GetEntityData() *types.CommonEntityData {
    multicastClient.EntityData.YFilter = multicastClient.YFilter
    multicastClient.EntityData.YangName = "multicast-client"
    multicastClient.EntityData.BundleName = "cisco_ios_xr"
    multicastClient.EntityData.ParentYangName = "multicast-clients"
    multicastClient.EntityData.SegmentPath = "multicast-client" + "[ip-address='" + fmt.Sprintf("%v", multicastClient.IpAddress) + "']"
    multicastClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastClient.EntityData.Children = make(map[string]types.YChild)
    multicastClient.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastClient.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", multicastClient.IpAddress}
    return &(multicastClient.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers
// Configures multicast server peers
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure NTP multicast group server peer. The type is slice of
    // Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer.
    MulticastServer []Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer
}

func (multicastServers *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers) GetEntityData() *types.CommonEntityData {
    multicastServers.EntityData.YFilter = multicastServers.YFilter
    multicastServers.EntityData.YangName = "multicast-servers"
    multicastServers.EntityData.BundleName = "cisco_ios_xr"
    multicastServers.EntityData.ParentYangName = "interface-multicast"
    multicastServers.EntityData.SegmentPath = "multicast-servers"
    multicastServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastServers.EntityData.Children = make(map[string]types.YChild)
    multicastServers.EntityData.Children["multicast-server"] = types.YChild{"MulticastServer", nil}
    for i := range multicastServers.MulticastServer {
        multicastServers.EntityData.Children[types.GetSegmentPath(&multicastServers.MulticastServer[i])] = types.YChild{"MulticastServer", &multicastServers.MulticastServer[i]}
    }
    multicastServers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multicastServers.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer
// Configure NTP multicast group server peer
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of a multicast group. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Authentication key. The type is interface{} with range: 1..65535.
    AuthenticationKey interface{}

    // NTP version. The type is interface{} with range: 2..4.
    Version interface{}

    // TTL. The type is interface{} with range: 1..255.
    Ttl interface{}
}

func (multicastServer *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceMulticast_MulticastServers_MulticastServer) GetEntityData() *types.CommonEntityData {
    multicastServer.EntityData.YFilter = multicastServer.YFilter
    multicastServer.EntityData.YangName = "multicast-server"
    multicastServer.EntityData.BundleName = "cisco_ios_xr"
    multicastServer.EntityData.ParentYangName = "multicast-servers"
    multicastServer.EntityData.SegmentPath = "multicast-server" + "[ip-address='" + fmt.Sprintf("%v", multicastServer.IpAddress) + "']"
    multicastServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastServer.EntityData.Children = make(map[string]types.YChild)
    multicastServer.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastServer.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", multicastServer.IpAddress}
    multicastServer.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", multicastServer.AuthenticationKey}
    multicastServer.EntityData.Leafs["version"] = types.YLeaf{"Version", multicastServer.Version}
    multicastServer.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", multicastServer.Ttl}
    return &(multicastServer.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast
// Configure NTP broadcast service
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Listen to NTP broadcasts. The type is interface{}.
    BroadcastClient interface{}

    // Configure NTP broadcast.
    Broadcast Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast
}

func (interfaceBroadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast) GetEntityData() *types.CommonEntityData {
    interfaceBroadcast.EntityData.YFilter = interfaceBroadcast.YFilter
    interfaceBroadcast.EntityData.YangName = "interface-broadcast"
    interfaceBroadcast.EntityData.BundleName = "cisco_ios_xr"
    interfaceBroadcast.EntityData.ParentYangName = "interface"
    interfaceBroadcast.EntityData.SegmentPath = "interface-broadcast"
    interfaceBroadcast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceBroadcast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceBroadcast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceBroadcast.EntityData.Children = make(map[string]types.YChild)
    interfaceBroadcast.EntityData.Children["broadcast"] = types.YChild{"Broadcast", &interfaceBroadcast.Broadcast}
    interfaceBroadcast.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceBroadcast.EntityData.Leafs["broadcast-client"] = types.YLeaf{"BroadcastClient", interfaceBroadcast.BroadcastClient}
    return &(interfaceBroadcast.EntityData)
}

// Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast
// Configure NTP broadcast
type Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination broadcast IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Authentication key. The type is interface{} with range: 1..65535.
    AuthenticationKey interface{}

    // NTP version. The type is interface{} with range: 2..4.
    NtpVersion interface{}
}

func (broadcast *Ntp_InterfaceTables_InterfaceTable_Interface_InterfaceBroadcast_Broadcast) GetEntityData() *types.CommonEntityData {
    broadcast.EntityData.YFilter = broadcast.YFilter
    broadcast.EntityData.YangName = "broadcast"
    broadcast.EntityData.BundleName = "cisco_ios_xr"
    broadcast.EntityData.ParentYangName = "interface-broadcast"
    broadcast.EntityData.SegmentPath = "broadcast"
    broadcast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    broadcast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    broadcast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    broadcast.EntityData.Children = make(map[string]types.YChild)
    broadcast.EntityData.Leafs = make(map[string]types.YLeaf)
    broadcast.EntityData.Leafs["address"] = types.YLeaf{"Address", broadcast.Address}
    broadcast.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", broadcast.AuthenticationKey}
    broadcast.EntityData.Leafs["ntp-version"] = types.YLeaf{"NtpVersion", broadcast.NtpVersion}
    return &(broadcast.EntityData)
}

// Ntp_AccessGroupTables
// Control NTP access
type Ntp_AccessGroupTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control NTP access. The type is slice of
    // Ntp_AccessGroupTables_AccessGroupTable.
    AccessGroupTable []Ntp_AccessGroupTables_AccessGroupTable
}

func (accessGroupTables *Ntp_AccessGroupTables) GetEntityData() *types.CommonEntityData {
    accessGroupTables.EntityData.YFilter = accessGroupTables.YFilter
    accessGroupTables.EntityData.YangName = "access-group-tables"
    accessGroupTables.EntityData.BundleName = "cisco_ios_xr"
    accessGroupTables.EntityData.ParentYangName = "ntp"
    accessGroupTables.EntityData.SegmentPath = "access-group-tables"
    accessGroupTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessGroupTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessGroupTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessGroupTables.EntityData.Children = make(map[string]types.YChild)
    accessGroupTables.EntityData.Children["access-group-table"] = types.YChild{"AccessGroupTable", nil}
    for i := range accessGroupTables.AccessGroupTable {
        accessGroupTables.EntityData.Children[types.GetSegmentPath(&accessGroupTables.AccessGroupTable[i])] = types.YChild{"AccessGroupTable", &accessGroupTables.AccessGroupTable[i]}
    }
    accessGroupTables.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessGroupTables.EntityData)
}

// Ntp_AccessGroupTables_AccessGroupTable
// Control NTP access
type Ntp_AccessGroupTables_AccessGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Configure NTP access address family. The type is slice of
    // Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable.
    AccessGroupAfTable []Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable
}

func (accessGroupTable *Ntp_AccessGroupTables_AccessGroupTable) GetEntityData() *types.CommonEntityData {
    accessGroupTable.EntityData.YFilter = accessGroupTable.YFilter
    accessGroupTable.EntityData.YangName = "access-group-table"
    accessGroupTable.EntityData.BundleName = "cisco_ios_xr"
    accessGroupTable.EntityData.ParentYangName = "access-group-tables"
    accessGroupTable.EntityData.SegmentPath = "access-group-table" + "[vrf-name='" + fmt.Sprintf("%v", accessGroupTable.VrfName) + "']"
    accessGroupTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessGroupTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessGroupTable.EntityData.Children = make(map[string]types.YChild)
    accessGroupTable.EntityData.Children["access-group-af-table"] = types.YChild{"AccessGroupAfTable", nil}
    for i := range accessGroupTable.AccessGroupAfTable {
        accessGroupTable.EntityData.Children[types.GetSegmentPath(&accessGroupTable.AccessGroupAfTable[i])] = types.YChild{"AccessGroupAfTable", &accessGroupTable.AccessGroupAfTable[i]}
    }
    accessGroupTable.EntityData.Leafs = make(map[string]types.YLeaf)
    accessGroupTable.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", accessGroupTable.VrfName}
    return &(accessGroupTable.EntityData)
}

// Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable
// Configure NTP access address family
type Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is NtpAccessAf.
    Af interface{}

    // Configure NTP access group. The type is slice of
    // Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup.
    AccessGroup []Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup
}

func (accessGroupAfTable *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable) GetEntityData() *types.CommonEntityData {
    accessGroupAfTable.EntityData.YFilter = accessGroupAfTable.YFilter
    accessGroupAfTable.EntityData.YangName = "access-group-af-table"
    accessGroupAfTable.EntityData.BundleName = "cisco_ios_xr"
    accessGroupAfTable.EntityData.ParentYangName = "access-group-table"
    accessGroupAfTable.EntityData.SegmentPath = "access-group-af-table" + "[af='" + fmt.Sprintf("%v", accessGroupAfTable.Af) + "']"
    accessGroupAfTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessGroupAfTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessGroupAfTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessGroupAfTable.EntityData.Children = make(map[string]types.YChild)
    accessGroupAfTable.EntityData.Children["access-group"] = types.YChild{"AccessGroup", nil}
    for i := range accessGroupAfTable.AccessGroup {
        accessGroupAfTable.EntityData.Children[types.GetSegmentPath(&accessGroupAfTable.AccessGroup[i])] = types.YChild{"AccessGroup", &accessGroupAfTable.AccessGroup[i]}
    }
    accessGroupAfTable.EntityData.Leafs = make(map[string]types.YLeaf)
    accessGroupAfTable.EntityData.Leafs["af"] = types.YLeaf{"Af", accessGroupAfTable.Af}
    return &(accessGroupAfTable.EntityData)
}

// Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup
// Configure NTP access group
type Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Access group type. The type is NtpAccess.
    AccessGroupType interface{}

    // Access list name - maximum 32 characters. The type is string. This
    // attribute is mandatory.
    AccessListName interface{}
}

func (accessGroup *Ntp_AccessGroupTables_AccessGroupTable_AccessGroupAfTable_AccessGroup) GetEntityData() *types.CommonEntityData {
    accessGroup.EntityData.YFilter = accessGroup.YFilter
    accessGroup.EntityData.YangName = "access-group"
    accessGroup.EntityData.BundleName = "cisco_ios_xr"
    accessGroup.EntityData.ParentYangName = "access-group-af-table"
    accessGroup.EntityData.SegmentPath = "access-group" + "[access-group-type='" + fmt.Sprintf("%v", accessGroup.AccessGroupType) + "']"
    accessGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessGroup.EntityData.Children = make(map[string]types.YChild)
    accessGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    accessGroup.EntityData.Leafs["access-group-type"] = types.YLeaf{"AccessGroupType", accessGroup.AccessGroupType}
    accessGroup.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", accessGroup.AccessListName}
    return &(accessGroup.EntityData)
}

