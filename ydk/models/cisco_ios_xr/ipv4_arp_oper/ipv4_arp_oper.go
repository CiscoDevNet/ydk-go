// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-arp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   arp-gmp: ARP-GMP global operational data
//   arp: arp
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_arp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_arp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-oper arp-gmp}", reflect.TypeOf(ArpGmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-oper:arp-gmp", reflect.TypeOf(ArpGmp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-oper arp}", reflect.TypeOf(Arp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-oper:arp", reflect.TypeOf(Arp{}))
}

// ArpGmpBagEntry represents ARP Entry type
type ArpGmpBagEntry string

const (
    // No state
    ArpGmpBagEntry_null ArpGmpBagEntry = "null"

    // Static
    ArpGmpBagEntry_static ArpGmpBagEntry = "static"

    // Alias
    ArpGmpBagEntry_alias ArpGmpBagEntry = "alias"
)

// ArpGmpBagEncap represents ARP encapsulation
type ArpGmpBagEncap string

const (
    // No encapsulation
    ArpGmpBagEncap_none ArpGmpBagEncap = "none"

    // ARPA
    ArpGmpBagEncap_arpa ArpGmpBagEncap = "arpa"

    // SNAP
    ArpGmpBagEncap_snap ArpGmpBagEncap = "snap"

    // 802 1Q
    ArpGmpBagEncap_ieee802_1q ArpGmpBagEncap = "ieee802-1q"

    // SRP
    ArpGmpBagEncap_srp ArpGmpBagEncap = "srp"

    // SRPA
    ArpGmpBagEncap_srpa ArpGmpBagEncap = "srpa"

    // SRPB
    ArpGmpBagEncap_srpb ArpGmpBagEncap = "srpb"
)

// IpArpBagEncap represents ARP encapsulation
type IpArpBagEncap string

const (
    // No encapsulation
    IpArpBagEncap_none IpArpBagEncap = "none"

    // ARPA
    IpArpBagEncap_arpa IpArpBagEncap = "arpa"

    // SNAP
    IpArpBagEncap_snap IpArpBagEncap = "snap"

    // 802 1Q
    IpArpBagEncap_ieee802_1q IpArpBagEncap = "ieee802-1q"

    // SRP
    IpArpBagEncap_srp IpArpBagEncap = "srp"

    // SRPA
    IpArpBagEncap_srpa IpArpBagEncap = "srpa"

    // SRPB
    IpArpBagEncap_srpb IpArpBagEncap = "srpb"
)

// IpArpBagFlags represents ARP flags
type IpArpBagFlags string

const (
    // No Flag
    IpArpBagFlags_flag_none IpArpBagFlags = "flag-none"

    // Dynamic learnt entry
    IpArpBagFlags_flag_dynamic IpArpBagFlags = "flag-dynamic"

    // EVPN Synced entry
    IpArpBagFlags_flag_evpn_sync IpArpBagFlags = "flag-evpn-sync"

    // Maximum Flag number
    IpArpBagFlags_flag_max IpArpBagFlags = "flag-max"
)

// IpArpBagState represents ARP state
type IpArpBagState string

const (
    // No state
    IpArpBagState_state_none IpArpBagState = "state-none"

    // Interface
    IpArpBagState_state_interface IpArpBagState = "state-interface"

    // Standby
    IpArpBagState_state_standby IpArpBagState = "state-standby"

    // Static
    IpArpBagState_state_static IpArpBagState = "state-static"

    // Alias
    IpArpBagState_state_alias IpArpBagState = "state-alias"

    // Mobile
    IpArpBagState_state_mobile IpArpBagState = "state-mobile"

    // Incomplete
    IpArpBagState_state_incomplete IpArpBagState = "state-incomplete"

    // Deleted
    IpArpBagState_state_deleted IpArpBagState = "state-deleted"

    // Dynamic
    IpArpBagState_state_dynamic IpArpBagState = "state-dynamic"

    // Probe
    IpArpBagState_state_probe IpArpBagState = "state-probe"

    // Purge delayed
    IpArpBagState_state_purge_delayed IpArpBagState = "state-purge-delayed"

    // DHCP installed
    IpArpBagState_state_dhcp IpArpBagState = "state-dhcp"

    // VXLAN installed
    IpArpBagState_state_vxlan IpArpBagState = "state-vxlan"

    // EVPN-SYNC installed
    IpArpBagState_state_evpn_sync IpArpBagState = "state-evpn-sync"

    // Satellite installed
    IpArpBagState_state_sat IpArpBagState = "state-sat"

    // Geo-redundancy sync'ed
    IpArpBagState_state_r_sync IpArpBagState = "state-r-sync"

    // Maximum state number
    IpArpBagState_state_max IpArpBagState = "state-max"
)

// IpArpBagMedia represents ARP media type
type IpArpBagMedia string

const (
    // ARPA
    IpArpBagMedia_media_arpa IpArpBagMedia = "media-arpa"

    // SRP
    IpArpBagMedia_media_srp IpArpBagMedia = "media-srp"

    // Unknown
    IpArpBagMedia_media_unknown IpArpBagMedia = "media-unknown"
)

// ArpResolutionHistoryStatus represents Arp resolution history status
type ArpResolutionHistoryStatus string

const (
    // No Status
    ArpResolutionHistoryStatus_status_none ArpResolutionHistoryStatus = "status-none"

    // Resolution Request Received
    ArpResolutionHistoryStatus_status_resolution_request ArpResolutionHistoryStatus = "status-resolution-request"

    // Resolved with ARP reply
    ArpResolutionHistoryStatus_status_resolved_reply ArpResolutionHistoryStatus = "status-resolved-reply"

    // Resolved with Grat ARP
    ArpResolutionHistoryStatus_status_resolved_grat_arp ArpResolutionHistoryStatus = "status-resolved-grat-arp"

    // Resolved with ARP Request
    ArpResolutionHistoryStatus_status_resolved_request ArpResolutionHistoryStatus = "status-resolved-request"

    // Resolved via a Linecard sync
    ArpResolutionHistoryStatus_status_resolved_lc_sync ArpResolutionHistoryStatus = "status-resolved-lc-sync"

    // Resolved via a Linecard sync while purge
    // delayed
    ArpResolutionHistoryStatus_status_resolved_lc_sync_purge_delay ArpResolutionHistoryStatus = "status-resolved-lc-sync-purge-delay"

    // Resolved from an ARP API client
    ArpResolutionHistoryStatus_status_resolved_client ArpResolutionHistoryStatus = "status-resolved-client"

    // Removed by an ARP API client
    ArpResolutionHistoryStatus_status_removed_client ArpResolutionHistoryStatus = "status-removed-client"

    // Already Resolved
    ArpResolutionHistoryStatus_status_already_resolved ArpResolutionHistoryStatus = "status-already-resolved"

    // Resolution Failed
    ArpResolutionHistoryStatus_status_failed ArpResolutionHistoryStatus = "status-failed"

    // Dropped because the Interface was down
    ArpResolutionHistoryStatus_status_dropped_interface_down ArpResolutionHistoryStatus = "status-dropped-interface-down"

    // Dropped because the Interface was broadcast
    // disabled
    ArpResolutionHistoryStatus_status_dropped_broadcast_disabled ArpResolutionHistoryStatus = "status-dropped-broadcast-disabled"

    // Dropped because the Interface was unavailable
    // to arp
    ArpResolutionHistoryStatus_status_dropped_interface_unavailable ArpResolutionHistoryStatus = "status-dropped-interface-unavailable"

    // The requested IP address didn't belong to the
    // subnet
    ArpResolutionHistoryStatus_status_dropped_bad_subnet ArpResolutionHistoryStatus = "status-dropped-bad-subnet"

    // Dynamic learning of ARP entries is disabled on
    // the interface
    ArpResolutionHistoryStatus_status_dropped_dynamic_learning_disabled ArpResolutionHistoryStatus = "status-dropped-dynamic-learning-disabled"

    // Out of Subnet address learning is disabled on
    // the interface
    ArpResolutionHistoryStatus_status_dropped_out_of_subnet_disabled ArpResolutionHistoryStatus = "status-dropped-out-of-subnet-disabled"

    // Removed by an ARP API client during a resync
    ArpResolutionHistoryStatus_status_removed_client_sweep ArpResolutionHistoryStatus = "status-removed-client-sweep"

    // Added by an ARP API client
    ArpResolutionHistoryStatus_status_added_client ArpResolutionHistoryStatus = "status-added-client"

    // Added by replication from ARP V1 during ISSU
    ArpResolutionHistoryStatus_status_added_v1 ArpResolutionHistoryStatus = "status-added-v1"

    // Removed by replication from ARP V1 during ISSU
    ArpResolutionHistoryStatus_status_removed_v1 ArpResolutionHistoryStatus = "status-removed-v1"

    // Resolved via a Peer Router sync
    ArpResolutionHistoryStatus_status_resolved_peer_sync ArpResolutionHistoryStatus = "status-resolved-peer-sync"

    // Learning unsolicited ARP packets is disabled on
    // this Interface
    ArpResolutionHistoryStatus_status_dropped_unsolicited_pak ArpResolutionHistoryStatus = "status-dropped-unsolicited-pak"
)

// ArpGmp
// ARP-GMP global operational data
type ArpGmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of VRF related ARP-GMP operational data.
    VrfInfos ArpGmp_VrfInfos

    // Table of per VRF ARP-GMP operational data.
    Vrfs ArpGmp_Vrfs
}

func (arpGmp *ArpGmp) GetEntityData() *types.CommonEntityData {
    arpGmp.EntityData.YFilter = arpGmp.YFilter
    arpGmp.EntityData.YangName = "arp-gmp"
    arpGmp.EntityData.BundleName = "cisco_ios_xr"
    arpGmp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-oper"
    arpGmp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-oper:arp-gmp"
    arpGmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arpGmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arpGmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arpGmp.EntityData.Children = make(map[string]types.YChild)
    arpGmp.EntityData.Children["vrf-infos"] = types.YChild{"VrfInfos", &arpGmp.VrfInfos}
    arpGmp.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &arpGmp.Vrfs}
    arpGmp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(arpGmp.EntityData)
}

// ArpGmp_VrfInfos
// Table of VRF related ARP-GMP operational data
type ArpGmp_VrfInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF related ARP-GMP operational data. The type is slice of
    // ArpGmp_VrfInfos_VrfInfo.
    VrfInfo []ArpGmp_VrfInfos_VrfInfo
}

func (vrfInfos *ArpGmp_VrfInfos) GetEntityData() *types.CommonEntityData {
    vrfInfos.EntityData.YFilter = vrfInfos.YFilter
    vrfInfos.EntityData.YangName = "vrf-infos"
    vrfInfos.EntityData.BundleName = "cisco_ios_xr"
    vrfInfos.EntityData.ParentYangName = "arp-gmp"
    vrfInfos.EntityData.SegmentPath = "vrf-infos"
    vrfInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfInfos.EntityData.Children = make(map[string]types.YChild)
    vrfInfos.EntityData.Children["vrf-info"] = types.YChild{"VrfInfo", nil}
    for i := range vrfInfos.VrfInfo {
        vrfInfos.EntityData.Children[types.GetSegmentPath(&vrfInfos.VrfInfo[i])] = types.YChild{"VrfInfo", &vrfInfos.VrfInfo[i]}
    }
    vrfInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfInfos.EntityData)
}

// ArpGmp_VrfInfos_VrfInfo
// VRF related ARP-GMP operational data
type ArpGmp_VrfInfos_VrfInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name for the default VRF use 'default'. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfIdNumber interface{}

    // IPv4 unicast table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // RSI registration handle. The type is interface{} with range: 0..4294967295.
    RsiHandle interface{}

    // RSI registration handle (top 32-bits). The type is interface{} with range:
    // 0..4294967295.
    RsiHandleHigh interface{}
}

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetEntityData() *types.CommonEntityData {
    vrfInfo.EntityData.YFilter = vrfInfo.YFilter
    vrfInfo.EntityData.YangName = "vrf-info"
    vrfInfo.EntityData.BundleName = "cisco_ios_xr"
    vrfInfo.EntityData.ParentYangName = "vrf-infos"
    vrfInfo.EntityData.SegmentPath = "vrf-info" + "[vrf-name='" + fmt.Sprintf("%v", vrfInfo.VrfName) + "']"
    vrfInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfInfo.EntityData.Children = make(map[string]types.YChild)
    vrfInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfInfo.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfInfo.VrfName}
    vrfInfo.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", vrfInfo.VrfNameXr}
    vrfInfo.EntityData.Leafs["vrf-id-number"] = types.YLeaf{"VrfIdNumber", vrfInfo.VrfIdNumber}
    vrfInfo.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", vrfInfo.TableId}
    vrfInfo.EntityData.Leafs["rsi-handle"] = types.YLeaf{"RsiHandle", vrfInfo.RsiHandle}
    vrfInfo.EntityData.Leafs["rsi-handle-high"] = types.YLeaf{"RsiHandleHigh", vrfInfo.RsiHandleHigh}
    return &(vrfInfo.EntityData)
}

// ArpGmp_Vrfs
// Table of per VRF ARP-GMP operational data
type ArpGmp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per VRF ARP-GMP operational data. The type is slice of ArpGmp_Vrfs_Vrf.
    Vrf []ArpGmp_Vrfs_Vrf
}

func (vrfs *ArpGmp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "arp-gmp"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// ArpGmp_Vrfs_Vrf
// Per VRF ARP-GMP operational data
type ArpGmp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name for the default VRF use 'default'. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Table of ARP-GMP configured IP addresses information.
    ConfiguredIpAddresses ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses

    // Table of ARP GMP route information.
    Routes ArpGmp_Vrfs_Vrf_Routes

    // Table of ARP GMP interface and associated configured IP data.
    InterfaceConfiguredIps ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps
}

func (vrf *ArpGmp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["configured-ip-addresses"] = types.YChild{"ConfiguredIpAddresses", &vrf.ConfiguredIpAddresses}
    vrf.EntityData.Children["routes"] = types.YChild{"Routes", &vrf.Routes}
    vrf.EntityData.Children["interface-configured-ips"] = types.YChild{"InterfaceConfiguredIps", &vrf.InterfaceConfiguredIps}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses
// Table of ARP-GMP configured IP addresses
// information
type ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP-GMP configured IP address information. The type is slice of
    // ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress.
    ConfiguredIpAddress []ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetEntityData() *types.CommonEntityData {
    configuredIpAddresses.EntityData.YFilter = configuredIpAddresses.YFilter
    configuredIpAddresses.EntityData.YangName = "configured-ip-addresses"
    configuredIpAddresses.EntityData.BundleName = "cisco_ios_xr"
    configuredIpAddresses.EntityData.ParentYangName = "vrf"
    configuredIpAddresses.EntityData.SegmentPath = "configured-ip-addresses"
    configuredIpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredIpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredIpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredIpAddresses.EntityData.Children = make(map[string]types.YChild)
    configuredIpAddresses.EntityData.Children["configured-ip-address"] = types.YChild{"ConfiguredIpAddress", nil}
    for i := range configuredIpAddresses.ConfiguredIpAddress {
        configuredIpAddresses.EntityData.Children[types.GetSegmentPath(&configuredIpAddresses.ConfiguredIpAddress[i])] = types.YChild{"ConfiguredIpAddress", &configuredIpAddresses.ConfiguredIpAddress[i]}
    }
    configuredIpAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configuredIpAddresses.EntityData)
}

// ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress
// ARP-GMP configured IP address information
type ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Configured ARP-GMP IP. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Hardware address . The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    HardwareAddress interface{}

    // Encap type. The type is ArpGmpBagEncap.
    EncapsulationType interface{}

    // Entry type static/alias. The type is ArpGmpBagEntry.
    EntryType interface{}
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetEntityData() *types.CommonEntityData {
    configuredIpAddress.EntityData.YFilter = configuredIpAddress.YFilter
    configuredIpAddress.EntityData.YangName = "configured-ip-address"
    configuredIpAddress.EntityData.BundleName = "cisco_ios_xr"
    configuredIpAddress.EntityData.ParentYangName = "configured-ip-addresses"
    configuredIpAddress.EntityData.SegmentPath = "configured-ip-address" + "[address='" + fmt.Sprintf("%v", configuredIpAddress.Address) + "']"
    configuredIpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredIpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredIpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredIpAddress.EntityData.Children = make(map[string]types.YChild)
    configuredIpAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    configuredIpAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", configuredIpAddress.Address}
    configuredIpAddress.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", configuredIpAddress.IpAddress}
    configuredIpAddress.EntityData.Leafs["hardware-address"] = types.YLeaf{"HardwareAddress", configuredIpAddress.HardwareAddress}
    configuredIpAddress.EntityData.Leafs["encapsulation-type"] = types.YLeaf{"EncapsulationType", configuredIpAddress.EncapsulationType}
    configuredIpAddress.EntityData.Leafs["entry-type"] = types.YLeaf{"EntryType", configuredIpAddress.EntryType}
    return &(configuredIpAddress.EntityData)
}

// ArpGmp_Vrfs_Vrf_Routes
// Table of ARP GMP route information
type ArpGmp_Vrfs_Vrf_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP GMP route information. The type is slice of
    // ArpGmp_Vrfs_Vrf_Routes_Route.
    Route []ArpGmp_Vrfs_Vrf_Routes_Route
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "vrf"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = make(map[string]types.YChild)
    routes.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range routes.Route {
        routes.EntityData.Children[types.GetSegmentPath(&routes.Route[i])] = types.YChild{"Route", &routes.Route[i]}
    }
    routes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routes.EntityData)
}

// ArpGmp_Vrfs_Vrf_Routes_Route
// ARP GMP route information
type ArpGmp_Vrfs_Vrf_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // IP address length. The type is interface{} with range: 0..255.
    PrefixLengthXr interface{}

    // Interface name (first element of InterfaceNames array). The type is string
    // with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceNameXr interface{}

    // Interface names. The type is slice of string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName []interface{}
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["address"] = types.YLeaf{"Address", route.Address}
    route.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", route.PrefixLength}
    route.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", route.IpAddress}
    route.EntityData.Leafs["prefix-length-xr"] = types.YLeaf{"PrefixLengthXr", route.PrefixLengthXr}
    route.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", route.InterfaceNameXr}
    route.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", route.InterfaceName}
    return &(route.EntityData)
}

// ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps
// Table of ARP GMP interface and associated
// configured IP data
type ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP GMP interface and associated configured IP data. The type is slice of
    // ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp.
    InterfaceConfiguredIp []ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetEntityData() *types.CommonEntityData {
    interfaceConfiguredIps.EntityData.YFilter = interfaceConfiguredIps.YFilter
    interfaceConfiguredIps.EntityData.YangName = "interface-configured-ips"
    interfaceConfiguredIps.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfiguredIps.EntityData.ParentYangName = "vrf"
    interfaceConfiguredIps.EntityData.SegmentPath = "interface-configured-ips"
    interfaceConfiguredIps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfiguredIps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfiguredIps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfiguredIps.EntityData.Children = make(map[string]types.YChild)
    interfaceConfiguredIps.EntityData.Children["interface-configured-ip"] = types.YChild{"InterfaceConfiguredIp", nil}
    for i := range interfaceConfiguredIps.InterfaceConfiguredIp {
        interfaceConfiguredIps.EntityData.Children[types.GetSegmentPath(&interfaceConfiguredIps.InterfaceConfiguredIp[i])] = types.YChild{"InterfaceConfiguredIp", &interfaceConfiguredIps.InterfaceConfiguredIp[i]}
    }
    interfaceConfiguredIps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceConfiguredIps.EntityData)
}

// ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp
// ARP GMP interface and associated configured
// IP data
type ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Configured ARP-GMP IP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceNameXr interface{}

    // Route reference count. The type is interface{} with range: 0..4294967295.
    ReferenceCount interface{}

    // Associated configuration entry.
    AssociatedConfigurationEntry ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetEntityData() *types.CommonEntityData {
    interfaceConfiguredIp.EntityData.YFilter = interfaceConfiguredIp.YFilter
    interfaceConfiguredIp.EntityData.YangName = "interface-configured-ip"
    interfaceConfiguredIp.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfiguredIp.EntityData.ParentYangName = "interface-configured-ips"
    interfaceConfiguredIp.EntityData.SegmentPath = "interface-configured-ip"
    interfaceConfiguredIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfiguredIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfiguredIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfiguredIp.EntityData.Children = make(map[string]types.YChild)
    interfaceConfiguredIp.EntityData.Children["associated-configuration-entry"] = types.YChild{"AssociatedConfigurationEntry", &interfaceConfiguredIp.AssociatedConfigurationEntry}
    interfaceConfiguredIp.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConfiguredIp.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceConfiguredIp.InterfaceName}
    interfaceConfiguredIp.EntityData.Leafs["address"] = types.YLeaf{"Address", interfaceConfiguredIp.Address}
    interfaceConfiguredIp.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", interfaceConfiguredIp.InterfaceNameXr}
    interfaceConfiguredIp.EntityData.Leafs["reference-count"] = types.YLeaf{"ReferenceCount", interfaceConfiguredIp.ReferenceCount}
    return &(interfaceConfiguredIp.EntityData)
}

// ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry
// Associated configuration entry
type ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Hardware address . The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    HardwareAddress interface{}

    // Encap type. The type is ArpGmpBagEncap.
    EncapsulationType interface{}

    // Entry type static/alias. The type is ArpGmpBagEntry.
    EntryType interface{}
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetEntityData() *types.CommonEntityData {
    associatedConfigurationEntry.EntityData.YFilter = associatedConfigurationEntry.YFilter
    associatedConfigurationEntry.EntityData.YangName = "associated-configuration-entry"
    associatedConfigurationEntry.EntityData.BundleName = "cisco_ios_xr"
    associatedConfigurationEntry.EntityData.ParentYangName = "interface-configured-ip"
    associatedConfigurationEntry.EntityData.SegmentPath = "associated-configuration-entry"
    associatedConfigurationEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associatedConfigurationEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associatedConfigurationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associatedConfigurationEntry.EntityData.Children = make(map[string]types.YChild)
    associatedConfigurationEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    associatedConfigurationEntry.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", associatedConfigurationEntry.IpAddress}
    associatedConfigurationEntry.EntityData.Leafs["hardware-address"] = types.YLeaf{"HardwareAddress", associatedConfigurationEntry.HardwareAddress}
    associatedConfigurationEntry.EntityData.Leafs["encapsulation-type"] = types.YLeaf{"EncapsulationType", associatedConfigurationEntry.EncapsulationType}
    associatedConfigurationEntry.EntityData.Leafs["entry-type"] = types.YLeaf{"EntryType", associatedConfigurationEntry.EntryType}
    return &(associatedConfigurationEntry.EntityData)
}

// Arp
// arp
type Arp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of per-node ARP operational data.
    Nodes Arp_Nodes
}

func (arp *Arp) GetEntityData() *types.CommonEntityData {
    arp.EntityData.YFilter = arp.YFilter
    arp.EntityData.YangName = "arp"
    arp.EntityData.BundleName = "cisco_ios_xr"
    arp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-oper"
    arp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-oper:arp"
    arp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arp.EntityData.Children = make(map[string]types.YChild)
    arp.EntityData.Children["nodes"] = types.YChild{"Nodes", &arp.Nodes}
    arp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(arp.EntityData)
}

// Arp_Nodes
// Table of per-node ARP operational data
type Arp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-node ARP operational data. The type is slice of Arp_Nodes_Node.
    Node []Arp_Nodes_Node
}

func (nodes *Arp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "arp"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Arp_Nodes_Node
// Per-node ARP operational data
type Arp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Per node dynamically-resolved ARP resolution history data.
    ResolutionHistoryDynamic Arp_Nodes_Node_ResolutionHistoryDynamic

    // ARP Traffic information per VRF.
    TrafficVrfs Arp_Nodes_Node_TrafficVrfs

    // Per node ARP Traffic data.
    TrafficNode Arp_Nodes_Node_TrafficNode

    // Per node client-installed ARP resolution history data.
    ResolutionHistoryClient Arp_Nodes_Node_ResolutionHistoryClient

    // Table of ARP entries.
    Entries Arp_Nodes_Node_Entries

    // ARP Traffic information per interface.
    TrafficInterfaces Arp_Nodes_Node_TrafficInterfaces
}

func (node *Arp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["resolution-history-dynamic"] = types.YChild{"ResolutionHistoryDynamic", &node.ResolutionHistoryDynamic}
    node.EntityData.Children["traffic-vrfs"] = types.YChild{"TrafficVrfs", &node.TrafficVrfs}
    node.EntityData.Children["traffic-node"] = types.YChild{"TrafficNode", &node.TrafficNode}
    node.EntityData.Children["resolution-history-client"] = types.YChild{"ResolutionHistoryClient", &node.ResolutionHistoryClient}
    node.EntityData.Children["entries"] = types.YChild{"Entries", &node.Entries}
    node.EntityData.Children["traffic-interfaces"] = types.YChild{"TrafficInterfaces", &node.TrafficInterfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Arp_Nodes_Node_ResolutionHistoryDynamic
// Per node dynamically-resolved ARP resolution
// history data
type Arp_Nodes_Node_ResolutionHistoryDynamic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resolution history array. The type is slice of
    // Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry.
    ArpEntry []Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetEntityData() *types.CommonEntityData {
    resolutionHistoryDynamic.EntityData.YFilter = resolutionHistoryDynamic.YFilter
    resolutionHistoryDynamic.EntityData.YangName = "resolution-history-dynamic"
    resolutionHistoryDynamic.EntityData.BundleName = "cisco_ios_xr"
    resolutionHistoryDynamic.EntityData.ParentYangName = "node"
    resolutionHistoryDynamic.EntityData.SegmentPath = "resolution-history-dynamic"
    resolutionHistoryDynamic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resolutionHistoryDynamic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resolutionHistoryDynamic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resolutionHistoryDynamic.EntityData.Children = make(map[string]types.YChild)
    resolutionHistoryDynamic.EntityData.Children["arp-entry"] = types.YChild{"ArpEntry", nil}
    for i := range resolutionHistoryDynamic.ArpEntry {
        resolutionHistoryDynamic.EntityData.Children[types.GetSegmentPath(&resolutionHistoryDynamic.ArpEntry[i])] = types.YChild{"ArpEntry", &resolutionHistoryDynamic.ArpEntry[i]}
    }
    resolutionHistoryDynamic.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resolutionHistoryDynamic.EntityData)
}

// Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry
// Resolution history array
type Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp for entry in nanoseconds since Epoch, i.e. since 00:00:00 UTC,
    // January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    NsecTimestamp interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IdbInterfaceName interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Resolution status. The type is ArpResolutionHistoryStatus.
    Status interface{}

    // Resolving Client ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientId interface{}

    // ARP entry state. The type is interface{} with range:
    // -2147483648..2147483647.
    EntryState interface{}

    // Resolution Request count. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestCount interface{}
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetEntityData() *types.CommonEntityData {
    arpEntry.EntityData.YFilter = arpEntry.YFilter
    arpEntry.EntityData.YangName = "arp-entry"
    arpEntry.EntityData.BundleName = "cisco_ios_xr"
    arpEntry.EntityData.ParentYangName = "resolution-history-dynamic"
    arpEntry.EntityData.SegmentPath = "arp-entry"
    arpEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arpEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arpEntry.EntityData.Children = make(map[string]types.YChild)
    arpEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    arpEntry.EntityData.Leafs["nsec-timestamp"] = types.YLeaf{"NsecTimestamp", arpEntry.NsecTimestamp}
    arpEntry.EntityData.Leafs["idb-interface-name"] = types.YLeaf{"IdbInterfaceName", arpEntry.IdbInterfaceName}
    arpEntry.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", arpEntry.Ipv4Address}
    arpEntry.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", arpEntry.MacAddress}
    arpEntry.EntityData.Leafs["status"] = types.YLeaf{"Status", arpEntry.Status}
    arpEntry.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", arpEntry.ClientId}
    arpEntry.EntityData.Leafs["entry-state"] = types.YLeaf{"EntryState", arpEntry.EntryState}
    arpEntry.EntityData.Leafs["resolution-request-count"] = types.YLeaf{"ResolutionRequestCount", arpEntry.ResolutionRequestCount}
    return &(arpEntry.EntityData)
}

// Arp_Nodes_Node_TrafficVrfs
// ARP Traffic information per VRF
type Arp_Nodes_Node_TrafficVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per VRF traffic data. The type is slice of
    // Arp_Nodes_Node_TrafficVrfs_TrafficVrf.
    TrafficVrf []Arp_Nodes_Node_TrafficVrfs_TrafficVrf
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetEntityData() *types.CommonEntityData {
    trafficVrfs.EntityData.YFilter = trafficVrfs.YFilter
    trafficVrfs.EntityData.YangName = "traffic-vrfs"
    trafficVrfs.EntityData.BundleName = "cisco_ios_xr"
    trafficVrfs.EntityData.ParentYangName = "node"
    trafficVrfs.EntityData.SegmentPath = "traffic-vrfs"
    trafficVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficVrfs.EntityData.Children = make(map[string]types.YChild)
    trafficVrfs.EntityData.Children["traffic-vrf"] = types.YChild{"TrafficVrf", nil}
    for i := range trafficVrfs.TrafficVrf {
        trafficVrfs.EntityData.Children[types.GetSegmentPath(&trafficVrfs.TrafficVrf[i])] = types.YChild{"TrafficVrf", &trafficVrfs.TrafficVrf[i]}
    }
    trafficVrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trafficVrfs.EntityData)
}

// Arp_Nodes_Node_TrafficVrfs_TrafficVrf
// Per VRF traffic data
type Arp_Nodes_Node_TrafficVrfs_TrafficVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // Total ARP requests received. The type is interface{} with range:
    // 0..4294967295.
    RequestsReceived interface{}

    // Total ARP replies received. The type is interface{} with range:
    // 0..4294967295.
    RepliesReceived interface{}

    // Total ARP requests sent. The type is interface{} with range: 0..4294967295.
    RequestsSent interface{}

    // Total ARP replies sent. The type is interface{} with range: 0..4294967295.
    RepliesSent interface{}

    // Total Proxy ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    ProxyRepliesSent interface{}

    // Total ARP requests received over subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    SubscrRequestsReceived interface{}

    // Total ARP replies sent over subscriber interface. The type is interface{}
    // with range: 0..4294967295.
    SubscrRepliesSent interface{}

    // Total ARP grat replies sent over subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    SubscrRepliesGratgSent interface{}

    // Total Local Proxy ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    LocalProxyRepliesSent interface{}

    // Total Gratuituous ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    GratuitousRepliesSent interface{}

    // Total ARP resolution requests received. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestsReceived interface{}

    // Total ARP resolution replies received. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRepliesReceived interface{}

    // total ARP resolution requests dropped. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestsDropped interface{}

    // Total errors for out of memory. The type is interface{} with range:
    // 0..4294967295.
    OutOfMemoryErrors interface{}

    // Total errors for no buffer. The type is interface{} with range:
    // 0..4294967295.
    NoBufferErrors interface{}

    // Total ARP entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    TotalEntries interface{}

    // Total dynamic entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    DynamicEntries interface{}

    // Total static entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    StaticEntries interface{}

    // Total alias entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    AliasEntries interface{}

    // Total interface entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    InterfaceEntries interface{}

    // Total standby entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    StandbyEntries interface{}

    // Total DHCP entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    DhcpEntries interface{}

    // Total VXLAN entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    VxlanEntries interface{}

    // Total ip packets droped on this node. The type is interface{} with range:
    // 0..4294967295.
    IpPacketsDroppedNode interface{}

    // Total ARP packets on node due to out of subnet. The type is interface{}
    // with range: 0..4294967295.
    ArpPacketNodeOutOfSubnet interface{}

    // Total ip packets droped on this interface. The type is interface{} with
    // range: 0..4294967295.
    IpPacketsDroppedInterface interface{}

    // Total arp packets on interface due to out of subnet. The type is
    // interface{} with range: 0..4294967295.
    ArpPacketInterfaceOutOfSubnet interface{}

    // Total unsolicited arp packets dropped. The type is interface{} with range:
    // 0..4294967295.
    ArpPacketUnsolicitedPacket interface{}

    // Total idb structures on this node. The type is interface{} with range:
    // 0..4294967295.
    IdbStructures interface{}
}

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetEntityData() *types.CommonEntityData {
    trafficVrf.EntityData.YFilter = trafficVrf.YFilter
    trafficVrf.EntityData.YangName = "traffic-vrf"
    trafficVrf.EntityData.BundleName = "cisco_ios_xr"
    trafficVrf.EntityData.ParentYangName = "traffic-vrfs"
    trafficVrf.EntityData.SegmentPath = "traffic-vrf" + "[vrf-name='" + fmt.Sprintf("%v", trafficVrf.VrfName) + "']"
    trafficVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficVrf.EntityData.Children = make(map[string]types.YChild)
    trafficVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", trafficVrf.VrfName}
    trafficVrf.EntityData.Leafs["requests-received"] = types.YLeaf{"RequestsReceived", trafficVrf.RequestsReceived}
    trafficVrf.EntityData.Leafs["replies-received"] = types.YLeaf{"RepliesReceived", trafficVrf.RepliesReceived}
    trafficVrf.EntityData.Leafs["requests-sent"] = types.YLeaf{"RequestsSent", trafficVrf.RequestsSent}
    trafficVrf.EntityData.Leafs["replies-sent"] = types.YLeaf{"RepliesSent", trafficVrf.RepliesSent}
    trafficVrf.EntityData.Leafs["proxy-replies-sent"] = types.YLeaf{"ProxyRepliesSent", trafficVrf.ProxyRepliesSent}
    trafficVrf.EntityData.Leafs["subscr-requests-received"] = types.YLeaf{"SubscrRequestsReceived", trafficVrf.SubscrRequestsReceived}
    trafficVrf.EntityData.Leafs["subscr-replies-sent"] = types.YLeaf{"SubscrRepliesSent", trafficVrf.SubscrRepliesSent}
    trafficVrf.EntityData.Leafs["subscr-replies-gratg-sent"] = types.YLeaf{"SubscrRepliesGratgSent", trafficVrf.SubscrRepliesGratgSent}
    trafficVrf.EntityData.Leafs["local-proxy-replies-sent"] = types.YLeaf{"LocalProxyRepliesSent", trafficVrf.LocalProxyRepliesSent}
    trafficVrf.EntityData.Leafs["gratuitous-replies-sent"] = types.YLeaf{"GratuitousRepliesSent", trafficVrf.GratuitousRepliesSent}
    trafficVrf.EntityData.Leafs["resolution-requests-received"] = types.YLeaf{"ResolutionRequestsReceived", trafficVrf.ResolutionRequestsReceived}
    trafficVrf.EntityData.Leafs["resolution-replies-received"] = types.YLeaf{"ResolutionRepliesReceived", trafficVrf.ResolutionRepliesReceived}
    trafficVrf.EntityData.Leafs["resolution-requests-dropped"] = types.YLeaf{"ResolutionRequestsDropped", trafficVrf.ResolutionRequestsDropped}
    trafficVrf.EntityData.Leafs["out-of-memory-errors"] = types.YLeaf{"OutOfMemoryErrors", trafficVrf.OutOfMemoryErrors}
    trafficVrf.EntityData.Leafs["no-buffer-errors"] = types.YLeaf{"NoBufferErrors", trafficVrf.NoBufferErrors}
    trafficVrf.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", trafficVrf.TotalEntries}
    trafficVrf.EntityData.Leafs["dynamic-entries"] = types.YLeaf{"DynamicEntries", trafficVrf.DynamicEntries}
    trafficVrf.EntityData.Leafs["static-entries"] = types.YLeaf{"StaticEntries", trafficVrf.StaticEntries}
    trafficVrf.EntityData.Leafs["alias-entries"] = types.YLeaf{"AliasEntries", trafficVrf.AliasEntries}
    trafficVrf.EntityData.Leafs["interface-entries"] = types.YLeaf{"InterfaceEntries", trafficVrf.InterfaceEntries}
    trafficVrf.EntityData.Leafs["standby-entries"] = types.YLeaf{"StandbyEntries", trafficVrf.StandbyEntries}
    trafficVrf.EntityData.Leafs["dhcp-entries"] = types.YLeaf{"DhcpEntries", trafficVrf.DhcpEntries}
    trafficVrf.EntityData.Leafs["vxlan-entries"] = types.YLeaf{"VxlanEntries", trafficVrf.VxlanEntries}
    trafficVrf.EntityData.Leafs["ip-packets-dropped-node"] = types.YLeaf{"IpPacketsDroppedNode", trafficVrf.IpPacketsDroppedNode}
    trafficVrf.EntityData.Leafs["arp-packet-node-out-of-subnet"] = types.YLeaf{"ArpPacketNodeOutOfSubnet", trafficVrf.ArpPacketNodeOutOfSubnet}
    trafficVrf.EntityData.Leafs["ip-packets-dropped-interface"] = types.YLeaf{"IpPacketsDroppedInterface", trafficVrf.IpPacketsDroppedInterface}
    trafficVrf.EntityData.Leafs["arp-packet-interface-out-of-subnet"] = types.YLeaf{"ArpPacketInterfaceOutOfSubnet", trafficVrf.ArpPacketInterfaceOutOfSubnet}
    trafficVrf.EntityData.Leafs["arp-packet-unsolicited-packet"] = types.YLeaf{"ArpPacketUnsolicitedPacket", trafficVrf.ArpPacketUnsolicitedPacket}
    trafficVrf.EntityData.Leafs["idb-structures"] = types.YLeaf{"IdbStructures", trafficVrf.IdbStructures}
    return &(trafficVrf.EntityData)
}

// Arp_Nodes_Node_TrafficNode
// Per node ARP Traffic data
type Arp_Nodes_Node_TrafficNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total ARP requests received. The type is interface{} with range:
    // 0..4294967295.
    RequestsReceived interface{}

    // Total ARP replies received. The type is interface{} with range:
    // 0..4294967295.
    RepliesReceived interface{}

    // Total ARP requests sent. The type is interface{} with range: 0..4294967295.
    RequestsSent interface{}

    // Total ARP replies sent. The type is interface{} with range: 0..4294967295.
    RepliesSent interface{}

    // Total Proxy ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    ProxyRepliesSent interface{}

    // Total ARP requests received over subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    SubscrRequestsReceived interface{}

    // Total ARP replies sent over subscriber interface. The type is interface{}
    // with range: 0..4294967295.
    SubscrRepliesSent interface{}

    // Total ARP grat replies sent over subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    SubscrRepliesGratgSent interface{}

    // Total Local Proxy ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    LocalProxyRepliesSent interface{}

    // Total Gratuituous ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    GratuitousRepliesSent interface{}

    // Total ARP resolution requests received. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestsReceived interface{}

    // Total ARP resolution replies received. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRepliesReceived interface{}

    // total ARP resolution requests dropped. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestsDropped interface{}

    // Total errors for out of memory. The type is interface{} with range:
    // 0..4294967295.
    OutOfMemoryErrors interface{}

    // Total errors for no buffer. The type is interface{} with range:
    // 0..4294967295.
    NoBufferErrors interface{}

    // Total ARP entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    TotalEntries interface{}

    // Total dynamic entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    DynamicEntries interface{}

    // Total static entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    StaticEntries interface{}

    // Total alias entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    AliasEntries interface{}

    // Total interface entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    InterfaceEntries interface{}

    // Total standby entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    StandbyEntries interface{}

    // Total DHCP entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    DhcpEntries interface{}

    // Total VXLAN entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    VxlanEntries interface{}

    // Total ip packets droped on this node. The type is interface{} with range:
    // 0..4294967295.
    IpPacketsDroppedNode interface{}

    // Total ARP packets on node due to out of subnet. The type is interface{}
    // with range: 0..4294967295.
    ArpPacketNodeOutOfSubnet interface{}

    // Total ip packets droped on this interface. The type is interface{} with
    // range: 0..4294967295.
    IpPacketsDroppedInterface interface{}

    // Total arp packets on interface due to out of subnet. The type is
    // interface{} with range: 0..4294967295.
    ArpPacketInterfaceOutOfSubnet interface{}

    // Total unsolicited arp packets dropped. The type is interface{} with range:
    // 0..4294967295.
    ArpPacketUnsolicitedPacket interface{}

    // Total idb structures on this node. The type is interface{} with range:
    // 0..4294967295.
    IdbStructures interface{}
}

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetEntityData() *types.CommonEntityData {
    trafficNode.EntityData.YFilter = trafficNode.YFilter
    trafficNode.EntityData.YangName = "traffic-node"
    trafficNode.EntityData.BundleName = "cisco_ios_xr"
    trafficNode.EntityData.ParentYangName = "node"
    trafficNode.EntityData.SegmentPath = "traffic-node"
    trafficNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficNode.EntityData.Children = make(map[string]types.YChild)
    trafficNode.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficNode.EntityData.Leafs["requests-received"] = types.YLeaf{"RequestsReceived", trafficNode.RequestsReceived}
    trafficNode.EntityData.Leafs["replies-received"] = types.YLeaf{"RepliesReceived", trafficNode.RepliesReceived}
    trafficNode.EntityData.Leafs["requests-sent"] = types.YLeaf{"RequestsSent", trafficNode.RequestsSent}
    trafficNode.EntityData.Leafs["replies-sent"] = types.YLeaf{"RepliesSent", trafficNode.RepliesSent}
    trafficNode.EntityData.Leafs["proxy-replies-sent"] = types.YLeaf{"ProxyRepliesSent", trafficNode.ProxyRepliesSent}
    trafficNode.EntityData.Leafs["subscr-requests-received"] = types.YLeaf{"SubscrRequestsReceived", trafficNode.SubscrRequestsReceived}
    trafficNode.EntityData.Leafs["subscr-replies-sent"] = types.YLeaf{"SubscrRepliesSent", trafficNode.SubscrRepliesSent}
    trafficNode.EntityData.Leafs["subscr-replies-gratg-sent"] = types.YLeaf{"SubscrRepliesGratgSent", trafficNode.SubscrRepliesGratgSent}
    trafficNode.EntityData.Leafs["local-proxy-replies-sent"] = types.YLeaf{"LocalProxyRepliesSent", trafficNode.LocalProxyRepliesSent}
    trafficNode.EntityData.Leafs["gratuitous-replies-sent"] = types.YLeaf{"GratuitousRepliesSent", trafficNode.GratuitousRepliesSent}
    trafficNode.EntityData.Leafs["resolution-requests-received"] = types.YLeaf{"ResolutionRequestsReceived", trafficNode.ResolutionRequestsReceived}
    trafficNode.EntityData.Leafs["resolution-replies-received"] = types.YLeaf{"ResolutionRepliesReceived", trafficNode.ResolutionRepliesReceived}
    trafficNode.EntityData.Leafs["resolution-requests-dropped"] = types.YLeaf{"ResolutionRequestsDropped", trafficNode.ResolutionRequestsDropped}
    trafficNode.EntityData.Leafs["out-of-memory-errors"] = types.YLeaf{"OutOfMemoryErrors", trafficNode.OutOfMemoryErrors}
    trafficNode.EntityData.Leafs["no-buffer-errors"] = types.YLeaf{"NoBufferErrors", trafficNode.NoBufferErrors}
    trafficNode.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", trafficNode.TotalEntries}
    trafficNode.EntityData.Leafs["dynamic-entries"] = types.YLeaf{"DynamicEntries", trafficNode.DynamicEntries}
    trafficNode.EntityData.Leafs["static-entries"] = types.YLeaf{"StaticEntries", trafficNode.StaticEntries}
    trafficNode.EntityData.Leafs["alias-entries"] = types.YLeaf{"AliasEntries", trafficNode.AliasEntries}
    trafficNode.EntityData.Leafs["interface-entries"] = types.YLeaf{"InterfaceEntries", trafficNode.InterfaceEntries}
    trafficNode.EntityData.Leafs["standby-entries"] = types.YLeaf{"StandbyEntries", trafficNode.StandbyEntries}
    trafficNode.EntityData.Leafs["dhcp-entries"] = types.YLeaf{"DhcpEntries", trafficNode.DhcpEntries}
    trafficNode.EntityData.Leafs["vxlan-entries"] = types.YLeaf{"VxlanEntries", trafficNode.VxlanEntries}
    trafficNode.EntityData.Leafs["ip-packets-dropped-node"] = types.YLeaf{"IpPacketsDroppedNode", trafficNode.IpPacketsDroppedNode}
    trafficNode.EntityData.Leafs["arp-packet-node-out-of-subnet"] = types.YLeaf{"ArpPacketNodeOutOfSubnet", trafficNode.ArpPacketNodeOutOfSubnet}
    trafficNode.EntityData.Leafs["ip-packets-dropped-interface"] = types.YLeaf{"IpPacketsDroppedInterface", trafficNode.IpPacketsDroppedInterface}
    trafficNode.EntityData.Leafs["arp-packet-interface-out-of-subnet"] = types.YLeaf{"ArpPacketInterfaceOutOfSubnet", trafficNode.ArpPacketInterfaceOutOfSubnet}
    trafficNode.EntityData.Leafs["arp-packet-unsolicited-packet"] = types.YLeaf{"ArpPacketUnsolicitedPacket", trafficNode.ArpPacketUnsolicitedPacket}
    trafficNode.EntityData.Leafs["idb-structures"] = types.YLeaf{"IdbStructures", trafficNode.IdbStructures}
    return &(trafficNode.EntityData)
}

// Arp_Nodes_Node_ResolutionHistoryClient
// Per node client-installed ARP resolution
// history data
type Arp_Nodes_Node_ResolutionHistoryClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resolution history array. The type is slice of
    // Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry.
    ArpEntry []Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetEntityData() *types.CommonEntityData {
    resolutionHistoryClient.EntityData.YFilter = resolutionHistoryClient.YFilter
    resolutionHistoryClient.EntityData.YangName = "resolution-history-client"
    resolutionHistoryClient.EntityData.BundleName = "cisco_ios_xr"
    resolutionHistoryClient.EntityData.ParentYangName = "node"
    resolutionHistoryClient.EntityData.SegmentPath = "resolution-history-client"
    resolutionHistoryClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resolutionHistoryClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resolutionHistoryClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resolutionHistoryClient.EntityData.Children = make(map[string]types.YChild)
    resolutionHistoryClient.EntityData.Children["arp-entry"] = types.YChild{"ArpEntry", nil}
    for i := range resolutionHistoryClient.ArpEntry {
        resolutionHistoryClient.EntityData.Children[types.GetSegmentPath(&resolutionHistoryClient.ArpEntry[i])] = types.YChild{"ArpEntry", &resolutionHistoryClient.ArpEntry[i]}
    }
    resolutionHistoryClient.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resolutionHistoryClient.EntityData)
}

// Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry
// Resolution history array
type Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp for entry in nanoseconds since Epoch, i.e. since 00:00:00 UTC,
    // January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    NsecTimestamp interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IdbInterfaceName interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Resolution status. The type is ArpResolutionHistoryStatus.
    Status interface{}

    // Resolving Client ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientId interface{}

    // ARP entry state. The type is interface{} with range:
    // -2147483648..2147483647.
    EntryState interface{}

    // Resolution Request count. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestCount interface{}
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetEntityData() *types.CommonEntityData {
    arpEntry.EntityData.YFilter = arpEntry.YFilter
    arpEntry.EntityData.YangName = "arp-entry"
    arpEntry.EntityData.BundleName = "cisco_ios_xr"
    arpEntry.EntityData.ParentYangName = "resolution-history-client"
    arpEntry.EntityData.SegmentPath = "arp-entry"
    arpEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arpEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arpEntry.EntityData.Children = make(map[string]types.YChild)
    arpEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    arpEntry.EntityData.Leafs["nsec-timestamp"] = types.YLeaf{"NsecTimestamp", arpEntry.NsecTimestamp}
    arpEntry.EntityData.Leafs["idb-interface-name"] = types.YLeaf{"IdbInterfaceName", arpEntry.IdbInterfaceName}
    arpEntry.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", arpEntry.Ipv4Address}
    arpEntry.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", arpEntry.MacAddress}
    arpEntry.EntityData.Leafs["status"] = types.YLeaf{"Status", arpEntry.Status}
    arpEntry.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", arpEntry.ClientId}
    arpEntry.EntityData.Leafs["entry-state"] = types.YLeaf{"EntryState", arpEntry.EntryState}
    arpEntry.EntityData.Leafs["resolution-request-count"] = types.YLeaf{"ResolutionRequestCount", arpEntry.ResolutionRequestCount}
    return &(arpEntry.EntityData)
}

// Arp_Nodes_Node_Entries
// Table of ARP entries
type Arp_Nodes_Node_Entries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP entry. The type is slice of Arp_Nodes_Node_Entries_Entry.
    Entry []Arp_Nodes_Node_Entries_Entry
}

func (entries *Arp_Nodes_Node_Entries) GetEntityData() *types.CommonEntityData {
    entries.EntityData.YFilter = entries.YFilter
    entries.EntityData.YangName = "entries"
    entries.EntityData.BundleName = "cisco_ios_xr"
    entries.EntityData.ParentYangName = "node"
    entries.EntityData.SegmentPath = "entries"
    entries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entries.EntityData.Children = make(map[string]types.YChild)
    entries.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range entries.Entry {
        entries.EntityData.Children[types.GetSegmentPath(&entries.Entry[i])] = types.YChild{"Entry", &entries.Entry[i]}
    }
    entries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entries.EntityData)
}

// Arp_Nodes_Node_Entries_Entry
// ARP entry
type Arp_Nodes_Node_Entries_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP Address of ARP entry. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Media type for this entry. The type is IpArpBagMedia.
    MediaType interface{}

    // State of this entry. The type is IpArpBagState.
    State interface{}

    // Flags of this entry. The type is IpArpBagFlags.
    Flag interface{}

    // Age of this entry. The type is interface{} with range:
    // 0..18446744073709551615.
    Age interface{}

    // Source encapsulation type. The type is IpArpBagEncap.
    EncapsulationType interface{}

    // Source hardware length. The type is interface{} with range: 0..255.
    HardwareLength interface{}

    // Hardware address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    HardwareAddress interface{}
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "entries"
    entry.EntityData.SegmentPath = "entry" + "[address='" + fmt.Sprintf("%v", entry.Address) + "']" + "[interface-name='" + fmt.Sprintf("%v", entry.InterfaceName) + "']"
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = make(map[string]types.YChild)
    entry.EntityData.Leafs = make(map[string]types.YLeaf)
    entry.EntityData.Leafs["address"] = types.YLeaf{"Address", entry.Address}
    entry.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", entry.InterfaceName}
    entry.EntityData.Leafs["media-type"] = types.YLeaf{"MediaType", entry.MediaType}
    entry.EntityData.Leafs["state"] = types.YLeaf{"State", entry.State}
    entry.EntityData.Leafs["flag"] = types.YLeaf{"Flag", entry.Flag}
    entry.EntityData.Leafs["age"] = types.YLeaf{"Age", entry.Age}
    entry.EntityData.Leafs["encapsulation-type"] = types.YLeaf{"EncapsulationType", entry.EncapsulationType}
    entry.EntityData.Leafs["hardware-length"] = types.YLeaf{"HardwareLength", entry.HardwareLength}
    entry.EntityData.Leafs["hardware-address"] = types.YLeaf{"HardwareAddress", entry.HardwareAddress}
    return &(entry.EntityData)
}

// Arp_Nodes_Node_TrafficInterfaces
// ARP Traffic information per interface
type Arp_Nodes_Node_TrafficInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per interface traffic data. The type is slice of
    // Arp_Nodes_Node_TrafficInterfaces_TrafficInterface.
    TrafficInterface []Arp_Nodes_Node_TrafficInterfaces_TrafficInterface
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetEntityData() *types.CommonEntityData {
    trafficInterfaces.EntityData.YFilter = trafficInterfaces.YFilter
    trafficInterfaces.EntityData.YangName = "traffic-interfaces"
    trafficInterfaces.EntityData.BundleName = "cisco_ios_xr"
    trafficInterfaces.EntityData.ParentYangName = "node"
    trafficInterfaces.EntityData.SegmentPath = "traffic-interfaces"
    trafficInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficInterfaces.EntityData.Children = make(map[string]types.YChild)
    trafficInterfaces.EntityData.Children["traffic-interface"] = types.YChild{"TrafficInterface", nil}
    for i := range trafficInterfaces.TrafficInterface {
        trafficInterfaces.EntityData.Children[types.GetSegmentPath(&trafficInterfaces.TrafficInterface[i])] = types.YChild{"TrafficInterface", &trafficInterfaces.TrafficInterface[i]}
    }
    trafficInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trafficInterfaces.EntityData)
}

// Arp_Nodes_Node_TrafficInterfaces_TrafficInterface
// Per interface traffic data
type Arp_Nodes_Node_TrafficInterfaces_TrafficInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Total ARP requests received. The type is interface{} with range:
    // 0..4294967295.
    RequestsReceived interface{}

    // Total ARP replies received. The type is interface{} with range:
    // 0..4294967295.
    RepliesReceived interface{}

    // Total ARP requests sent. The type is interface{} with range: 0..4294967295.
    RequestsSent interface{}

    // Total ARP replies sent. The type is interface{} with range: 0..4294967295.
    RepliesSent interface{}

    // Total Proxy ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    ProxyRepliesSent interface{}

    // Total ARP requests received over subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    SubscrRequestsReceived interface{}

    // Total ARP replies sent over subscriber interface. The type is interface{}
    // with range: 0..4294967295.
    SubscrRepliesSent interface{}

    // Total ARP grat replies sent over subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    SubscrRepliesGratgSent interface{}

    // Total Local Proxy ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    LocalProxyRepliesSent interface{}

    // Total Gratuituous ARP replies sent. The type is interface{} with range:
    // 0..4294967295.
    GratuitousRepliesSent interface{}

    // Total ARP resolution requests received. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestsReceived interface{}

    // Total ARP resolution replies received. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRepliesReceived interface{}

    // total ARP resolution requests dropped. The type is interface{} with range:
    // 0..4294967295.
    ResolutionRequestsDropped interface{}

    // Total errors for out of memory. The type is interface{} with range:
    // 0..4294967295.
    OutOfMemoryErrors interface{}

    // Total errors for no buffer. The type is interface{} with range:
    // 0..4294967295.
    NoBufferErrors interface{}

    // Total ARP entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    TotalEntries interface{}

    // Total dynamic entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    DynamicEntries interface{}

    // Total static entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    StaticEntries interface{}

    // Total alias entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    AliasEntries interface{}

    // Total interface entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    InterfaceEntries interface{}

    // Total standby entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    StandbyEntries interface{}

    // Total DHCP entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    DhcpEntries interface{}

    // Total VXLAN entries in the cache. The type is interface{} with range:
    // 0..4294967295.
    VxlanEntries interface{}

    // Total ip packets droped on this node. The type is interface{} with range:
    // 0..4294967295.
    IpPacketsDroppedNode interface{}

    // Total ARP packets on node due to out of subnet. The type is interface{}
    // with range: 0..4294967295.
    ArpPacketNodeOutOfSubnet interface{}

    // Total ip packets droped on this interface. The type is interface{} with
    // range: 0..4294967295.
    IpPacketsDroppedInterface interface{}

    // Total arp packets on interface due to out of subnet. The type is
    // interface{} with range: 0..4294967295.
    ArpPacketInterfaceOutOfSubnet interface{}

    // Total unsolicited arp packets dropped. The type is interface{} with range:
    // 0..4294967295.
    ArpPacketUnsolicitedPacket interface{}

    // Total idb structures on this node. The type is interface{} with range:
    // 0..4294967295.
    IdbStructures interface{}
}

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetEntityData() *types.CommonEntityData {
    trafficInterface.EntityData.YFilter = trafficInterface.YFilter
    trafficInterface.EntityData.YangName = "traffic-interface"
    trafficInterface.EntityData.BundleName = "cisco_ios_xr"
    trafficInterface.EntityData.ParentYangName = "traffic-interfaces"
    trafficInterface.EntityData.SegmentPath = "traffic-interface" + "[interface-name='" + fmt.Sprintf("%v", trafficInterface.InterfaceName) + "']"
    trafficInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficInterface.EntityData.Children = make(map[string]types.YChild)
    trafficInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", trafficInterface.InterfaceName}
    trafficInterface.EntityData.Leafs["requests-received"] = types.YLeaf{"RequestsReceived", trafficInterface.RequestsReceived}
    trafficInterface.EntityData.Leafs["replies-received"] = types.YLeaf{"RepliesReceived", trafficInterface.RepliesReceived}
    trafficInterface.EntityData.Leafs["requests-sent"] = types.YLeaf{"RequestsSent", trafficInterface.RequestsSent}
    trafficInterface.EntityData.Leafs["replies-sent"] = types.YLeaf{"RepliesSent", trafficInterface.RepliesSent}
    trafficInterface.EntityData.Leafs["proxy-replies-sent"] = types.YLeaf{"ProxyRepliesSent", trafficInterface.ProxyRepliesSent}
    trafficInterface.EntityData.Leafs["subscr-requests-received"] = types.YLeaf{"SubscrRequestsReceived", trafficInterface.SubscrRequestsReceived}
    trafficInterface.EntityData.Leafs["subscr-replies-sent"] = types.YLeaf{"SubscrRepliesSent", trafficInterface.SubscrRepliesSent}
    trafficInterface.EntityData.Leafs["subscr-replies-gratg-sent"] = types.YLeaf{"SubscrRepliesGratgSent", trafficInterface.SubscrRepliesGratgSent}
    trafficInterface.EntityData.Leafs["local-proxy-replies-sent"] = types.YLeaf{"LocalProxyRepliesSent", trafficInterface.LocalProxyRepliesSent}
    trafficInterface.EntityData.Leafs["gratuitous-replies-sent"] = types.YLeaf{"GratuitousRepliesSent", trafficInterface.GratuitousRepliesSent}
    trafficInterface.EntityData.Leafs["resolution-requests-received"] = types.YLeaf{"ResolutionRequestsReceived", trafficInterface.ResolutionRequestsReceived}
    trafficInterface.EntityData.Leafs["resolution-replies-received"] = types.YLeaf{"ResolutionRepliesReceived", trafficInterface.ResolutionRepliesReceived}
    trafficInterface.EntityData.Leafs["resolution-requests-dropped"] = types.YLeaf{"ResolutionRequestsDropped", trafficInterface.ResolutionRequestsDropped}
    trafficInterface.EntityData.Leafs["out-of-memory-errors"] = types.YLeaf{"OutOfMemoryErrors", trafficInterface.OutOfMemoryErrors}
    trafficInterface.EntityData.Leafs["no-buffer-errors"] = types.YLeaf{"NoBufferErrors", trafficInterface.NoBufferErrors}
    trafficInterface.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", trafficInterface.TotalEntries}
    trafficInterface.EntityData.Leafs["dynamic-entries"] = types.YLeaf{"DynamicEntries", trafficInterface.DynamicEntries}
    trafficInterface.EntityData.Leafs["static-entries"] = types.YLeaf{"StaticEntries", trafficInterface.StaticEntries}
    trafficInterface.EntityData.Leafs["alias-entries"] = types.YLeaf{"AliasEntries", trafficInterface.AliasEntries}
    trafficInterface.EntityData.Leafs["interface-entries"] = types.YLeaf{"InterfaceEntries", trafficInterface.InterfaceEntries}
    trafficInterface.EntityData.Leafs["standby-entries"] = types.YLeaf{"StandbyEntries", trafficInterface.StandbyEntries}
    trafficInterface.EntityData.Leafs["dhcp-entries"] = types.YLeaf{"DhcpEntries", trafficInterface.DhcpEntries}
    trafficInterface.EntityData.Leafs["vxlan-entries"] = types.YLeaf{"VxlanEntries", trafficInterface.VxlanEntries}
    trafficInterface.EntityData.Leafs["ip-packets-dropped-node"] = types.YLeaf{"IpPacketsDroppedNode", trafficInterface.IpPacketsDroppedNode}
    trafficInterface.EntityData.Leafs["arp-packet-node-out-of-subnet"] = types.YLeaf{"ArpPacketNodeOutOfSubnet", trafficInterface.ArpPacketNodeOutOfSubnet}
    trafficInterface.EntityData.Leafs["ip-packets-dropped-interface"] = types.YLeaf{"IpPacketsDroppedInterface", trafficInterface.IpPacketsDroppedInterface}
    trafficInterface.EntityData.Leafs["arp-packet-interface-out-of-subnet"] = types.YLeaf{"ArpPacketInterfaceOutOfSubnet", trafficInterface.ArpPacketInterfaceOutOfSubnet}
    trafficInterface.EntityData.Leafs["arp-packet-unsolicited-packet"] = types.YLeaf{"ArpPacketUnsolicitedPacket", trafficInterface.ArpPacketUnsolicitedPacket}
    trafficInterface.EntityData.Leafs["idb-structures"] = types.YLeaf{"IdbStructures", trafficInterface.IdbStructures}
    return &(trafficInterface.EntityData)
}

