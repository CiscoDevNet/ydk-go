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

// ArpGmp
// ARP-GMP global operational data
type ArpGmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of VRF related ARP-GMP operational data.
    VrfInfos ArpGmp_VrfInfos

    // Table of per VRF ARP-GMP operational data.
    Vrfs ArpGmp_Vrfs
}

func (arpGmp *ArpGmp) GetFilter() yfilter.YFilter { return arpGmp.YFilter }

func (arpGmp *ArpGmp) SetFilter(yf yfilter.YFilter) { arpGmp.YFilter = yf }

func (arpGmp *ArpGmp) GetGoName(yname string) string {
    if yname == "vrf-infos" { return "VrfInfos" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (arpGmp *ArpGmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-arp-oper:arp-gmp"
}

func (arpGmp *ArpGmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-infos" {
        return &arpGmp.VrfInfos
    }
    if childYangName == "vrfs" {
        return &arpGmp.Vrfs
    }
    return nil
}

func (arpGmp *ArpGmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-infos"] = &arpGmp.VrfInfos
    children["vrfs"] = &arpGmp.Vrfs
    return children
}

func (arpGmp *ArpGmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (arpGmp *ArpGmp) GetBundleName() string { return "cisco_ios_xr" }

func (arpGmp *ArpGmp) GetYangName() string { return "arp-gmp" }

func (arpGmp *ArpGmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arpGmp *ArpGmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arpGmp *ArpGmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arpGmp *ArpGmp) SetParent(parent types.Entity) { arpGmp.parent = parent }

func (arpGmp *ArpGmp) GetParent() types.Entity { return arpGmp.parent }

func (arpGmp *ArpGmp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-arp-oper" }

// ArpGmp_VrfInfos
// Table of VRF related ARP-GMP operational data
type ArpGmp_VrfInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF related ARP-GMP operational data. The type is slice of
    // ArpGmp_VrfInfos_VrfInfo.
    VrfInfo []ArpGmp_VrfInfos_VrfInfo
}

func (vrfInfos *ArpGmp_VrfInfos) GetFilter() yfilter.YFilter { return vrfInfos.YFilter }

func (vrfInfos *ArpGmp_VrfInfos) SetFilter(yf yfilter.YFilter) { vrfInfos.YFilter = yf }

func (vrfInfos *ArpGmp_VrfInfos) GetGoName(yname string) string {
    if yname == "vrf-info" { return "VrfInfo" }
    return ""
}

func (vrfInfos *ArpGmp_VrfInfos) GetSegmentPath() string {
    return "vrf-infos"
}

func (vrfInfos *ArpGmp_VrfInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-info" {
        for _, c := range vrfInfos.VrfInfo {
            if vrfInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpGmp_VrfInfos_VrfInfo{}
        vrfInfos.VrfInfo = append(vrfInfos.VrfInfo, child)
        return &vrfInfos.VrfInfo[len(vrfInfos.VrfInfo)-1]
    }
    return nil
}

func (vrfInfos *ArpGmp_VrfInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfInfos.VrfInfo {
        children[vrfInfos.VrfInfo[i].GetSegmentPath()] = &vrfInfos.VrfInfo[i]
    }
    return children
}

func (vrfInfos *ArpGmp_VrfInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfInfos *ArpGmp_VrfInfos) GetBundleName() string { return "cisco_ios_xr" }

func (vrfInfos *ArpGmp_VrfInfos) GetYangName() string { return "vrf-infos" }

func (vrfInfos *ArpGmp_VrfInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfInfos *ArpGmp_VrfInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfInfos *ArpGmp_VrfInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfInfos *ArpGmp_VrfInfos) SetParent(parent types.Entity) { vrfInfos.parent = parent }

func (vrfInfos *ArpGmp_VrfInfos) GetParent() types.Entity { return vrfInfos.parent }

func (vrfInfos *ArpGmp_VrfInfos) GetParentYangName() string { return "arp-gmp" }

// ArpGmp_VrfInfos_VrfInfo
// VRF related ARP-GMP operational data
type ArpGmp_VrfInfos_VrfInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name for the default VRF use 'default'. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetFilter() yfilter.YFilter { return vrfInfo.YFilter }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) SetFilter(yf yfilter.YFilter) { vrfInfo.YFilter = yf }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    if yname == "vrf-id-number" { return "VrfIdNumber" }
    if yname == "table-id" { return "TableId" }
    if yname == "rsi-handle" { return "RsiHandle" }
    if yname == "rsi-handle-high" { return "RsiHandleHigh" }
    return ""
}

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetSegmentPath() string {
    return "vrf-info" + "[vrf-name='" + fmt.Sprintf("%v", vrfInfo.VrfName) + "']"
}

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfInfo.VrfName
    leafs["vrf-name-xr"] = vrfInfo.VrfNameXr
    leafs["vrf-id-number"] = vrfInfo.VrfIdNumber
    leafs["table-id"] = vrfInfo.TableId
    leafs["rsi-handle"] = vrfInfo.RsiHandle
    leafs["rsi-handle-high"] = vrfInfo.RsiHandleHigh
    return leafs
}

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetBundleName() string { return "cisco_ios_xr" }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetYangName() string { return "vrf-info" }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) SetParent(parent types.Entity) { vrfInfo.parent = parent }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetParent() types.Entity { return vrfInfo.parent }

func (vrfInfo *ArpGmp_VrfInfos_VrfInfo) GetParentYangName() string { return "vrf-infos" }

// ArpGmp_Vrfs
// Table of per VRF ARP-GMP operational data
type ArpGmp_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per VRF ARP-GMP operational data. The type is slice of ArpGmp_Vrfs_Vrf.
    Vrf []ArpGmp_Vrfs_Vrf
}

func (vrfs *ArpGmp_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *ArpGmp_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *ArpGmp_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *ArpGmp_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *ArpGmp_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpGmp_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *ArpGmp_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *ArpGmp_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *ArpGmp_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *ArpGmp_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *ArpGmp_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *ArpGmp_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *ArpGmp_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *ArpGmp_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *ArpGmp_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *ArpGmp_Vrfs) GetParentYangName() string { return "arp-gmp" }

// ArpGmp_Vrfs_Vrf
// Per VRF ARP-GMP operational data
type ArpGmp_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name for the default VRF use 'default'. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Table of ARP-GMP configured IP addresses information.
    ConfiguredIpAddresses ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses

    // Table of ARP GMP route information.
    Routes ArpGmp_Vrfs_Vrf_Routes

    // Table of ARP GMP interface and associated configured IP data.
    InterfaceConfiguredIps ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps
}

func (vrf *ArpGmp_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *ArpGmp_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *ArpGmp_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "configured-ip-addresses" { return "ConfiguredIpAddresses" }
    if yname == "routes" { return "Routes" }
    if yname == "interface-configured-ips" { return "InterfaceConfiguredIps" }
    return ""
}

func (vrf *ArpGmp_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *ArpGmp_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-ip-addresses" {
        return &vrf.ConfiguredIpAddresses
    }
    if childYangName == "routes" {
        return &vrf.Routes
    }
    if childYangName == "interface-configured-ips" {
        return &vrf.InterfaceConfiguredIps
    }
    return nil
}

func (vrf *ArpGmp_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-ip-addresses"] = &vrf.ConfiguredIpAddresses
    children["routes"] = &vrf.Routes
    children["interface-configured-ips"] = &vrf.InterfaceConfiguredIps
    return children
}

func (vrf *ArpGmp_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *ArpGmp_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *ArpGmp_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *ArpGmp_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *ArpGmp_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *ArpGmp_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *ArpGmp_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *ArpGmp_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *ArpGmp_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses
// Table of ARP-GMP configured IP addresses
// information
type ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP-GMP configured IP address information. The type is slice of
    // ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress.
    ConfiguredIpAddress []ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetFilter() yfilter.YFilter { return configuredIpAddresses.YFilter }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) SetFilter(yf yfilter.YFilter) { configuredIpAddresses.YFilter = yf }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetGoName(yname string) string {
    if yname == "configured-ip-address" { return "ConfiguredIpAddress" }
    return ""
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetSegmentPath() string {
    return "configured-ip-addresses"
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-ip-address" {
        for _, c := range configuredIpAddresses.ConfiguredIpAddress {
            if configuredIpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress{}
        configuredIpAddresses.ConfiguredIpAddress = append(configuredIpAddresses.ConfiguredIpAddress, child)
        return &configuredIpAddresses.ConfiguredIpAddress[len(configuredIpAddresses.ConfiguredIpAddress)-1]
    }
    return nil
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configuredIpAddresses.ConfiguredIpAddress {
        children[configuredIpAddresses.ConfiguredIpAddress[i].GetSegmentPath()] = &configuredIpAddresses.ConfiguredIpAddress[i]
    }
    return children
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetYangName() string { return "configured-ip-addresses" }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) SetParent(parent types.Entity) { configuredIpAddresses.parent = parent }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetParent() types.Entity { return configuredIpAddresses.parent }

func (configuredIpAddresses *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses) GetParentYangName() string { return "vrf" }

// ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress
// ARP-GMP configured IP address information
type ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Configured ARP-GMP IP. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Hardware address . The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HardwareAddress interface{}

    // Encap type. The type is ArpGmpBagEncap.
    EncapsulationType interface{}

    // Entry type static/alias. The type is ArpGmpBagEntry.
    EntryType interface{}
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetFilter() yfilter.YFilter { return configuredIpAddress.YFilter }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) SetFilter(yf yfilter.YFilter) { configuredIpAddress.YFilter = yf }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "hardware-address" { return "HardwareAddress" }
    if yname == "encapsulation-type" { return "EncapsulationType" }
    if yname == "entry-type" { return "EntryType" }
    return ""
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetSegmentPath() string {
    return "configured-ip-address" + "[address='" + fmt.Sprintf("%v", configuredIpAddress.Address) + "']"
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = configuredIpAddress.Address
    leafs["ip-address"] = configuredIpAddress.IpAddress
    leafs["hardware-address"] = configuredIpAddress.HardwareAddress
    leafs["encapsulation-type"] = configuredIpAddress.EncapsulationType
    leafs["entry-type"] = configuredIpAddress.EntryType
    return leafs
}

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetYangName() string { return "configured-ip-address" }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) SetParent(parent types.Entity) { configuredIpAddress.parent = parent }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetParent() types.Entity { return configuredIpAddress.parent }

func (configuredIpAddress *ArpGmp_Vrfs_Vrf_ConfiguredIpAddresses_ConfiguredIpAddress) GetParentYangName() string { return "configured-ip-addresses" }

// ArpGmp_Vrfs_Vrf_Routes
// Table of ARP GMP route information
type ArpGmp_Vrfs_Vrf_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP GMP route information. The type is slice of
    // ArpGmp_Vrfs_Vrf_Routes_Route.
    Route []ArpGmp_Vrfs_Vrf_Routes_Route
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *ArpGmp_Vrfs_Vrf_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpGmp_Vrfs_Vrf_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetYangName() string { return "routes" }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *ArpGmp_Vrfs_Vrf_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetParent() types.Entity { return routes.parent }

func (routes *ArpGmp_Vrfs_Vrf_Routes) GetParentYangName() string { return "vrf" }

// ArpGmp_Vrfs_Vrf_Routes_Route
// ARP GMP route information
type ArpGmp_Vrfs_Vrf_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // IP address length. The type is interface{} with range: 0..255.
    PrefixLengthXr interface{}

    // Interface name (first element of InterfaceNames array). The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // Interface names. The type is slice of string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName []interface{}
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "prefix-length-xr" { return "PrefixLengthXr" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = route.Address
    leafs["prefix-length"] = route.PrefixLength
    leafs["ip-address"] = route.IpAddress
    leafs["prefix-length-xr"] = route.PrefixLengthXr
    leafs["interface-name-xr"] = route.InterfaceNameXr
    leafs["interface-name"] = route.InterfaceName
    return leafs
}

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetYangName() string { return "route" }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *ArpGmp_Vrfs_Vrf_Routes_Route) GetParentYangName() string { return "routes" }

// ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps
// Table of ARP GMP interface and associated
// configured IP data
type ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP GMP interface and associated configured IP data. The type is slice of
    // ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp.
    InterfaceConfiguredIp []ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetFilter() yfilter.YFilter { return interfaceConfiguredIps.YFilter }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) SetFilter(yf yfilter.YFilter) { interfaceConfiguredIps.YFilter = yf }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetGoName(yname string) string {
    if yname == "interface-configured-ip" { return "InterfaceConfiguredIp" }
    return ""
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetSegmentPath() string {
    return "interface-configured-ips"
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-configured-ip" {
        for _, c := range interfaceConfiguredIps.InterfaceConfiguredIp {
            if interfaceConfiguredIps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp{}
        interfaceConfiguredIps.InterfaceConfiguredIp = append(interfaceConfiguredIps.InterfaceConfiguredIp, child)
        return &interfaceConfiguredIps.InterfaceConfiguredIp[len(interfaceConfiguredIps.InterfaceConfiguredIp)-1]
    }
    return nil
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceConfiguredIps.InterfaceConfiguredIp {
        children[interfaceConfiguredIps.InterfaceConfiguredIp[i].GetSegmentPath()] = &interfaceConfiguredIps.InterfaceConfiguredIp[i]
    }
    return children
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetYangName() string { return "interface-configured-ips" }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) SetParent(parent types.Entity) { interfaceConfiguredIps.parent = parent }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetParent() types.Entity { return interfaceConfiguredIps.parent }

func (interfaceConfiguredIps *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps) GetParentYangName() string { return "vrf" }

// ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp
// ARP GMP interface and associated configured
// IP data
type ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Configured ARP-GMP IP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // Route reference count. The type is interface{} with range: 0..4294967295.
    ReferenceCount interface{}

    // Associated configuration entry.
    AssociatedConfigurationEntry ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetFilter() yfilter.YFilter { return interfaceConfiguredIp.YFilter }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) SetFilter(yf yfilter.YFilter) { interfaceConfiguredIp.YFilter = yf }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "address" { return "Address" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "reference-count" { return "ReferenceCount" }
    if yname == "associated-configuration-entry" { return "AssociatedConfigurationEntry" }
    return ""
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetSegmentPath() string {
    return "interface-configured-ip"
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "associated-configuration-entry" {
        return &interfaceConfiguredIp.AssociatedConfigurationEntry
    }
    return nil
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["associated-configuration-entry"] = &interfaceConfiguredIp.AssociatedConfigurationEntry
    return children
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceConfiguredIp.InterfaceName
    leafs["address"] = interfaceConfiguredIp.Address
    leafs["interface-name-xr"] = interfaceConfiguredIp.InterfaceNameXr
    leafs["reference-count"] = interfaceConfiguredIp.ReferenceCount
    return leafs
}

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetYangName() string { return "interface-configured-ip" }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) SetParent(parent types.Entity) { interfaceConfiguredIp.parent = parent }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetParent() types.Entity { return interfaceConfiguredIp.parent }

func (interfaceConfiguredIp *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp) GetParentYangName() string { return "interface-configured-ips" }

// ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry
// Associated configuration entry
type ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Hardware address . The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HardwareAddress interface{}

    // Encap type. The type is ArpGmpBagEncap.
    EncapsulationType interface{}

    // Entry type static/alias. The type is ArpGmpBagEntry.
    EntryType interface{}
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetFilter() yfilter.YFilter { return associatedConfigurationEntry.YFilter }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) SetFilter(yf yfilter.YFilter) { associatedConfigurationEntry.YFilter = yf }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "hardware-address" { return "HardwareAddress" }
    if yname == "encapsulation-type" { return "EncapsulationType" }
    if yname == "entry-type" { return "EntryType" }
    return ""
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetSegmentPath() string {
    return "associated-configuration-entry"
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = associatedConfigurationEntry.IpAddress
    leafs["hardware-address"] = associatedConfigurationEntry.HardwareAddress
    leafs["encapsulation-type"] = associatedConfigurationEntry.EncapsulationType
    leafs["entry-type"] = associatedConfigurationEntry.EntryType
    return leafs
}

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetBundleName() string { return "cisco_ios_xr" }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetYangName() string { return "associated-configuration-entry" }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) SetParent(parent types.Entity) { associatedConfigurationEntry.parent = parent }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetParent() types.Entity { return associatedConfigurationEntry.parent }

func (associatedConfigurationEntry *ArpGmp_Vrfs_Vrf_InterfaceConfiguredIps_InterfaceConfiguredIp_AssociatedConfigurationEntry) GetParentYangName() string { return "interface-configured-ip" }

// Arp
// arp
type Arp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of per-node ARP operational data.
    Nodes Arp_Nodes
}

func (arp *Arp) GetFilter() yfilter.YFilter { return arp.YFilter }

func (arp *Arp) SetFilter(yf yfilter.YFilter) { arp.YFilter = yf }

func (arp *Arp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (arp *Arp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-arp-oper:arp"
}

func (arp *Arp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &arp.Nodes
    }
    return nil
}

func (arp *Arp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &arp.Nodes
    return children
}

func (arp *Arp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (arp *Arp) GetBundleName() string { return "cisco_ios_xr" }

func (arp *Arp) GetYangName() string { return "arp" }

func (arp *Arp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arp *Arp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arp *Arp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arp *Arp) SetParent(parent types.Entity) { arp.parent = parent }

func (arp *Arp) GetParent() types.Entity { return arp.parent }

func (arp *Arp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-arp-oper" }

// Arp_Nodes
// Table of per-node ARP operational data
type Arp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-node ARP operational data. The type is slice of Arp_Nodes_Node.
    Node []Arp_Nodes_Node
}

func (nodes *Arp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Arp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Arp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Arp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Arp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Arp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Arp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Arp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Arp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Arp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Arp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Arp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Arp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Arp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Arp_Nodes) GetParentYangName() string { return "arp" }

// Arp_Nodes_Node
// Per-node ARP operational data
type Arp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

func (node *Arp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Arp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Arp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "resolution-history-dynamic" { return "ResolutionHistoryDynamic" }
    if yname == "traffic-vrfs" { return "TrafficVrfs" }
    if yname == "traffic-node" { return "TrafficNode" }
    if yname == "resolution-history-client" { return "ResolutionHistoryClient" }
    if yname == "entries" { return "Entries" }
    if yname == "traffic-interfaces" { return "TrafficInterfaces" }
    return ""
}

func (node *Arp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Arp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resolution-history-dynamic" {
        return &node.ResolutionHistoryDynamic
    }
    if childYangName == "traffic-vrfs" {
        return &node.TrafficVrfs
    }
    if childYangName == "traffic-node" {
        return &node.TrafficNode
    }
    if childYangName == "resolution-history-client" {
        return &node.ResolutionHistoryClient
    }
    if childYangName == "entries" {
        return &node.Entries
    }
    if childYangName == "traffic-interfaces" {
        return &node.TrafficInterfaces
    }
    return nil
}

func (node *Arp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resolution-history-dynamic"] = &node.ResolutionHistoryDynamic
    children["traffic-vrfs"] = &node.TrafficVrfs
    children["traffic-node"] = &node.TrafficNode
    children["resolution-history-client"] = &node.ResolutionHistoryClient
    children["entries"] = &node.Entries
    children["traffic-interfaces"] = &node.TrafficInterfaces
    return children
}

func (node *Arp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Arp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Arp_Nodes_Node) GetYangName() string { return "node" }

func (node *Arp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Arp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Arp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Arp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Arp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Arp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Arp_Nodes_Node_ResolutionHistoryDynamic
// Per node dynamically-resolved ARP resolution
// history data
type Arp_Nodes_Node_ResolutionHistoryDynamic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resolution history array. The type is slice of
    // Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry.
    ArpEntry []Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetFilter() yfilter.YFilter { return resolutionHistoryDynamic.YFilter }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) SetFilter(yf yfilter.YFilter) { resolutionHistoryDynamic.YFilter = yf }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetGoName(yname string) string {
    if yname == "arp-entry" { return "ArpEntry" }
    return ""
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetSegmentPath() string {
    return "resolution-history-dynamic"
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "arp-entry" {
        for _, c := range resolutionHistoryDynamic.ArpEntry {
            if resolutionHistoryDynamic.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry{}
        resolutionHistoryDynamic.ArpEntry = append(resolutionHistoryDynamic.ArpEntry, child)
        return &resolutionHistoryDynamic.ArpEntry[len(resolutionHistoryDynamic.ArpEntry)-1]
    }
    return nil
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range resolutionHistoryDynamic.ArpEntry {
        children[resolutionHistoryDynamic.ArpEntry[i].GetSegmentPath()] = &resolutionHistoryDynamic.ArpEntry[i]
    }
    return children
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetBundleName() string { return "cisco_ios_xr" }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetYangName() string { return "resolution-history-dynamic" }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) SetParent(parent types.Entity) { resolutionHistoryDynamic.parent = parent }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetParent() types.Entity { return resolutionHistoryDynamic.parent }

func (resolutionHistoryDynamic *Arp_Nodes_Node_ResolutionHistoryDynamic) GetParentYangName() string { return "node" }

// Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry
// Resolution history array
type Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp for entry in nanoseconds since Epoch, i.e. since 00:00:00 UTC,
    // January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    NsecTimestamp interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    IdbInterfaceName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetFilter() yfilter.YFilter { return arpEntry.YFilter }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) SetFilter(yf yfilter.YFilter) { arpEntry.YFilter = yf }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetGoName(yname string) string {
    if yname == "nsec-timestamp" { return "NsecTimestamp" }
    if yname == "idb-interface-name" { return "IdbInterfaceName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "status" { return "Status" }
    if yname == "client-id" { return "ClientId" }
    if yname == "entry-state" { return "EntryState" }
    if yname == "resolution-request-count" { return "ResolutionRequestCount" }
    return ""
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetSegmentPath() string {
    return "arp-entry"
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nsec-timestamp"] = arpEntry.NsecTimestamp
    leafs["idb-interface-name"] = arpEntry.IdbInterfaceName
    leafs["ipv4-address"] = arpEntry.Ipv4Address
    leafs["mac-address"] = arpEntry.MacAddress
    leafs["status"] = arpEntry.Status
    leafs["client-id"] = arpEntry.ClientId
    leafs["entry-state"] = arpEntry.EntryState
    leafs["resolution-request-count"] = arpEntry.ResolutionRequestCount
    return leafs
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetBundleName() string { return "cisco_ios_xr" }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetYangName() string { return "arp-entry" }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) SetParent(parent types.Entity) { arpEntry.parent = parent }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetParent() types.Entity { return arpEntry.parent }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryDynamic_ArpEntry) GetParentYangName() string { return "resolution-history-dynamic" }

// Arp_Nodes_Node_TrafficVrfs
// ARP Traffic information per VRF
type Arp_Nodes_Node_TrafficVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per VRF traffic data. The type is slice of
    // Arp_Nodes_Node_TrafficVrfs_TrafficVrf.
    TrafficVrf []Arp_Nodes_Node_TrafficVrfs_TrafficVrf
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetFilter() yfilter.YFilter { return trafficVrfs.YFilter }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) SetFilter(yf yfilter.YFilter) { trafficVrfs.YFilter = yf }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetGoName(yname string) string {
    if yname == "traffic-vrf" { return "TrafficVrf" }
    return ""
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetSegmentPath() string {
    return "traffic-vrfs"
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic-vrf" {
        for _, c := range trafficVrfs.TrafficVrf {
            if trafficVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arp_Nodes_Node_TrafficVrfs_TrafficVrf{}
        trafficVrfs.TrafficVrf = append(trafficVrfs.TrafficVrf, child)
        return &trafficVrfs.TrafficVrf[len(trafficVrfs.TrafficVrf)-1]
    }
    return nil
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trafficVrfs.TrafficVrf {
        children[trafficVrfs.TrafficVrf[i].GetSegmentPath()] = &trafficVrfs.TrafficVrf[i]
    }
    return children
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetBundleName() string { return "cisco_ios_xr" }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetYangName() string { return "traffic-vrfs" }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) SetParent(parent types.Entity) { trafficVrfs.parent = parent }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetParent() types.Entity { return trafficVrfs.parent }

func (trafficVrfs *Arp_Nodes_Node_TrafficVrfs) GetParentYangName() string { return "node" }

// Arp_Nodes_Node_TrafficVrfs_TrafficVrf
// Per VRF traffic data
type Arp_Nodes_Node_TrafficVrfs_TrafficVrf struct {
    parent types.Entity
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

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetFilter() yfilter.YFilter { return trafficVrf.YFilter }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) SetFilter(yf yfilter.YFilter) { trafficVrf.YFilter = yf }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "requests-received" { return "RequestsReceived" }
    if yname == "replies-received" { return "RepliesReceived" }
    if yname == "requests-sent" { return "RequestsSent" }
    if yname == "replies-sent" { return "RepliesSent" }
    if yname == "proxy-replies-sent" { return "ProxyRepliesSent" }
    if yname == "subscr-requests-received" { return "SubscrRequestsReceived" }
    if yname == "subscr-replies-sent" { return "SubscrRepliesSent" }
    if yname == "subscr-replies-gratg-sent" { return "SubscrRepliesGratgSent" }
    if yname == "local-proxy-replies-sent" { return "LocalProxyRepliesSent" }
    if yname == "gratuitous-replies-sent" { return "GratuitousRepliesSent" }
    if yname == "resolution-requests-received" { return "ResolutionRequestsReceived" }
    if yname == "resolution-replies-received" { return "ResolutionRepliesReceived" }
    if yname == "resolution-requests-dropped" { return "ResolutionRequestsDropped" }
    if yname == "out-of-memory-errors" { return "OutOfMemoryErrors" }
    if yname == "no-buffer-errors" { return "NoBufferErrors" }
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "dynamic-entries" { return "DynamicEntries" }
    if yname == "static-entries" { return "StaticEntries" }
    if yname == "alias-entries" { return "AliasEntries" }
    if yname == "interface-entries" { return "InterfaceEntries" }
    if yname == "standby-entries" { return "StandbyEntries" }
    if yname == "dhcp-entries" { return "DhcpEntries" }
    if yname == "vxlan-entries" { return "VxlanEntries" }
    if yname == "ip-packets-dropped-node" { return "IpPacketsDroppedNode" }
    if yname == "arp-packet-node-out-of-subnet" { return "ArpPacketNodeOutOfSubnet" }
    if yname == "ip-packets-dropped-interface" { return "IpPacketsDroppedInterface" }
    if yname == "arp-packet-interface-out-of-subnet" { return "ArpPacketInterfaceOutOfSubnet" }
    if yname == "arp-packet-unsolicited-packet" { return "ArpPacketUnsolicitedPacket" }
    if yname == "idb-structures" { return "IdbStructures" }
    return ""
}

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetSegmentPath() string {
    return "traffic-vrf" + "[vrf-name='" + fmt.Sprintf("%v", trafficVrf.VrfName) + "']"
}

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = trafficVrf.VrfName
    leafs["requests-received"] = trafficVrf.RequestsReceived
    leafs["replies-received"] = trafficVrf.RepliesReceived
    leafs["requests-sent"] = trafficVrf.RequestsSent
    leafs["replies-sent"] = trafficVrf.RepliesSent
    leafs["proxy-replies-sent"] = trafficVrf.ProxyRepliesSent
    leafs["subscr-requests-received"] = trafficVrf.SubscrRequestsReceived
    leafs["subscr-replies-sent"] = trafficVrf.SubscrRepliesSent
    leafs["subscr-replies-gratg-sent"] = trafficVrf.SubscrRepliesGratgSent
    leafs["local-proxy-replies-sent"] = trafficVrf.LocalProxyRepliesSent
    leafs["gratuitous-replies-sent"] = trafficVrf.GratuitousRepliesSent
    leafs["resolution-requests-received"] = trafficVrf.ResolutionRequestsReceived
    leafs["resolution-replies-received"] = trafficVrf.ResolutionRepliesReceived
    leafs["resolution-requests-dropped"] = trafficVrf.ResolutionRequestsDropped
    leafs["out-of-memory-errors"] = trafficVrf.OutOfMemoryErrors
    leafs["no-buffer-errors"] = trafficVrf.NoBufferErrors
    leafs["total-entries"] = trafficVrf.TotalEntries
    leafs["dynamic-entries"] = trafficVrf.DynamicEntries
    leafs["static-entries"] = trafficVrf.StaticEntries
    leafs["alias-entries"] = trafficVrf.AliasEntries
    leafs["interface-entries"] = trafficVrf.InterfaceEntries
    leafs["standby-entries"] = trafficVrf.StandbyEntries
    leafs["dhcp-entries"] = trafficVrf.DhcpEntries
    leafs["vxlan-entries"] = trafficVrf.VxlanEntries
    leafs["ip-packets-dropped-node"] = trafficVrf.IpPacketsDroppedNode
    leafs["arp-packet-node-out-of-subnet"] = trafficVrf.ArpPacketNodeOutOfSubnet
    leafs["ip-packets-dropped-interface"] = trafficVrf.IpPacketsDroppedInterface
    leafs["arp-packet-interface-out-of-subnet"] = trafficVrf.ArpPacketInterfaceOutOfSubnet
    leafs["arp-packet-unsolicited-packet"] = trafficVrf.ArpPacketUnsolicitedPacket
    leafs["idb-structures"] = trafficVrf.IdbStructures
    return leafs
}

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetBundleName() string { return "cisco_ios_xr" }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetYangName() string { return "traffic-vrf" }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) SetParent(parent types.Entity) { trafficVrf.parent = parent }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetParent() types.Entity { return trafficVrf.parent }

func (trafficVrf *Arp_Nodes_Node_TrafficVrfs_TrafficVrf) GetParentYangName() string { return "traffic-vrfs" }

// Arp_Nodes_Node_TrafficNode
// Per node ARP Traffic data
type Arp_Nodes_Node_TrafficNode struct {
    parent types.Entity
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

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetFilter() yfilter.YFilter { return trafficNode.YFilter }

func (trafficNode *Arp_Nodes_Node_TrafficNode) SetFilter(yf yfilter.YFilter) { trafficNode.YFilter = yf }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetGoName(yname string) string {
    if yname == "requests-received" { return "RequestsReceived" }
    if yname == "replies-received" { return "RepliesReceived" }
    if yname == "requests-sent" { return "RequestsSent" }
    if yname == "replies-sent" { return "RepliesSent" }
    if yname == "proxy-replies-sent" { return "ProxyRepliesSent" }
    if yname == "subscr-requests-received" { return "SubscrRequestsReceived" }
    if yname == "subscr-replies-sent" { return "SubscrRepliesSent" }
    if yname == "subscr-replies-gratg-sent" { return "SubscrRepliesGratgSent" }
    if yname == "local-proxy-replies-sent" { return "LocalProxyRepliesSent" }
    if yname == "gratuitous-replies-sent" { return "GratuitousRepliesSent" }
    if yname == "resolution-requests-received" { return "ResolutionRequestsReceived" }
    if yname == "resolution-replies-received" { return "ResolutionRepliesReceived" }
    if yname == "resolution-requests-dropped" { return "ResolutionRequestsDropped" }
    if yname == "out-of-memory-errors" { return "OutOfMemoryErrors" }
    if yname == "no-buffer-errors" { return "NoBufferErrors" }
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "dynamic-entries" { return "DynamicEntries" }
    if yname == "static-entries" { return "StaticEntries" }
    if yname == "alias-entries" { return "AliasEntries" }
    if yname == "interface-entries" { return "InterfaceEntries" }
    if yname == "standby-entries" { return "StandbyEntries" }
    if yname == "dhcp-entries" { return "DhcpEntries" }
    if yname == "vxlan-entries" { return "VxlanEntries" }
    if yname == "ip-packets-dropped-node" { return "IpPacketsDroppedNode" }
    if yname == "arp-packet-node-out-of-subnet" { return "ArpPacketNodeOutOfSubnet" }
    if yname == "ip-packets-dropped-interface" { return "IpPacketsDroppedInterface" }
    if yname == "arp-packet-interface-out-of-subnet" { return "ArpPacketInterfaceOutOfSubnet" }
    if yname == "arp-packet-unsolicited-packet" { return "ArpPacketUnsolicitedPacket" }
    if yname == "idb-structures" { return "IdbStructures" }
    return ""
}

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetSegmentPath() string {
    return "traffic-node"
}

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["requests-received"] = trafficNode.RequestsReceived
    leafs["replies-received"] = trafficNode.RepliesReceived
    leafs["requests-sent"] = trafficNode.RequestsSent
    leafs["replies-sent"] = trafficNode.RepliesSent
    leafs["proxy-replies-sent"] = trafficNode.ProxyRepliesSent
    leafs["subscr-requests-received"] = trafficNode.SubscrRequestsReceived
    leafs["subscr-replies-sent"] = trafficNode.SubscrRepliesSent
    leafs["subscr-replies-gratg-sent"] = trafficNode.SubscrRepliesGratgSent
    leafs["local-proxy-replies-sent"] = trafficNode.LocalProxyRepliesSent
    leafs["gratuitous-replies-sent"] = trafficNode.GratuitousRepliesSent
    leafs["resolution-requests-received"] = trafficNode.ResolutionRequestsReceived
    leafs["resolution-replies-received"] = trafficNode.ResolutionRepliesReceived
    leafs["resolution-requests-dropped"] = trafficNode.ResolutionRequestsDropped
    leafs["out-of-memory-errors"] = trafficNode.OutOfMemoryErrors
    leafs["no-buffer-errors"] = trafficNode.NoBufferErrors
    leafs["total-entries"] = trafficNode.TotalEntries
    leafs["dynamic-entries"] = trafficNode.DynamicEntries
    leafs["static-entries"] = trafficNode.StaticEntries
    leafs["alias-entries"] = trafficNode.AliasEntries
    leafs["interface-entries"] = trafficNode.InterfaceEntries
    leafs["standby-entries"] = trafficNode.StandbyEntries
    leafs["dhcp-entries"] = trafficNode.DhcpEntries
    leafs["vxlan-entries"] = trafficNode.VxlanEntries
    leafs["ip-packets-dropped-node"] = trafficNode.IpPacketsDroppedNode
    leafs["arp-packet-node-out-of-subnet"] = trafficNode.ArpPacketNodeOutOfSubnet
    leafs["ip-packets-dropped-interface"] = trafficNode.IpPacketsDroppedInterface
    leafs["arp-packet-interface-out-of-subnet"] = trafficNode.ArpPacketInterfaceOutOfSubnet
    leafs["arp-packet-unsolicited-packet"] = trafficNode.ArpPacketUnsolicitedPacket
    leafs["idb-structures"] = trafficNode.IdbStructures
    return leafs
}

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetBundleName() string { return "cisco_ios_xr" }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetYangName() string { return "traffic-node" }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficNode *Arp_Nodes_Node_TrafficNode) SetParent(parent types.Entity) { trafficNode.parent = parent }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetParent() types.Entity { return trafficNode.parent }

func (trafficNode *Arp_Nodes_Node_TrafficNode) GetParentYangName() string { return "node" }

// Arp_Nodes_Node_ResolutionHistoryClient
// Per node client-installed ARP resolution
// history data
type Arp_Nodes_Node_ResolutionHistoryClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resolution history array. The type is slice of
    // Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry.
    ArpEntry []Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetFilter() yfilter.YFilter { return resolutionHistoryClient.YFilter }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) SetFilter(yf yfilter.YFilter) { resolutionHistoryClient.YFilter = yf }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetGoName(yname string) string {
    if yname == "arp-entry" { return "ArpEntry" }
    return ""
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetSegmentPath() string {
    return "resolution-history-client"
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "arp-entry" {
        for _, c := range resolutionHistoryClient.ArpEntry {
            if resolutionHistoryClient.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry{}
        resolutionHistoryClient.ArpEntry = append(resolutionHistoryClient.ArpEntry, child)
        return &resolutionHistoryClient.ArpEntry[len(resolutionHistoryClient.ArpEntry)-1]
    }
    return nil
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range resolutionHistoryClient.ArpEntry {
        children[resolutionHistoryClient.ArpEntry[i].GetSegmentPath()] = &resolutionHistoryClient.ArpEntry[i]
    }
    return children
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetBundleName() string { return "cisco_ios_xr" }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetYangName() string { return "resolution-history-client" }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) SetParent(parent types.Entity) { resolutionHistoryClient.parent = parent }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetParent() types.Entity { return resolutionHistoryClient.parent }

func (resolutionHistoryClient *Arp_Nodes_Node_ResolutionHistoryClient) GetParentYangName() string { return "node" }

// Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry
// Resolution history array
type Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp for entry in nanoseconds since Epoch, i.e. since 00:00:00 UTC,
    // January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    NsecTimestamp interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    IdbInterfaceName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetFilter() yfilter.YFilter { return arpEntry.YFilter }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) SetFilter(yf yfilter.YFilter) { arpEntry.YFilter = yf }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetGoName(yname string) string {
    if yname == "nsec-timestamp" { return "NsecTimestamp" }
    if yname == "idb-interface-name" { return "IdbInterfaceName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "status" { return "Status" }
    if yname == "client-id" { return "ClientId" }
    if yname == "entry-state" { return "EntryState" }
    if yname == "resolution-request-count" { return "ResolutionRequestCount" }
    return ""
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetSegmentPath() string {
    return "arp-entry"
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nsec-timestamp"] = arpEntry.NsecTimestamp
    leafs["idb-interface-name"] = arpEntry.IdbInterfaceName
    leafs["ipv4-address"] = arpEntry.Ipv4Address
    leafs["mac-address"] = arpEntry.MacAddress
    leafs["status"] = arpEntry.Status
    leafs["client-id"] = arpEntry.ClientId
    leafs["entry-state"] = arpEntry.EntryState
    leafs["resolution-request-count"] = arpEntry.ResolutionRequestCount
    return leafs
}

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetBundleName() string { return "cisco_ios_xr" }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetYangName() string { return "arp-entry" }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) SetParent(parent types.Entity) { arpEntry.parent = parent }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetParent() types.Entity { return arpEntry.parent }

func (arpEntry *Arp_Nodes_Node_ResolutionHistoryClient_ArpEntry) GetParentYangName() string { return "resolution-history-client" }

// Arp_Nodes_Node_Entries
// Table of ARP entries
type Arp_Nodes_Node_Entries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP entry. The type is slice of Arp_Nodes_Node_Entries_Entry.
    Entry []Arp_Nodes_Node_Entries_Entry
}

func (entries *Arp_Nodes_Node_Entries) GetFilter() yfilter.YFilter { return entries.YFilter }

func (entries *Arp_Nodes_Node_Entries) SetFilter(yf yfilter.YFilter) { entries.YFilter = yf }

func (entries *Arp_Nodes_Node_Entries) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (entries *Arp_Nodes_Node_Entries) GetSegmentPath() string {
    return "entries"
}

func (entries *Arp_Nodes_Node_Entries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry" {
        for _, c := range entries.Entry {
            if entries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arp_Nodes_Node_Entries_Entry{}
        entries.Entry = append(entries.Entry, child)
        return &entries.Entry[len(entries.Entry)-1]
    }
    return nil
}

func (entries *Arp_Nodes_Node_Entries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entries.Entry {
        children[entries.Entry[i].GetSegmentPath()] = &entries.Entry[i]
    }
    return children
}

func (entries *Arp_Nodes_Node_Entries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entries *Arp_Nodes_Node_Entries) GetBundleName() string { return "cisco_ios_xr" }

func (entries *Arp_Nodes_Node_Entries) GetYangName() string { return "entries" }

func (entries *Arp_Nodes_Node_Entries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entries *Arp_Nodes_Node_Entries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entries *Arp_Nodes_Node_Entries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entries *Arp_Nodes_Node_Entries) SetParent(parent types.Entity) { entries.parent = parent }

func (entries *Arp_Nodes_Node_Entries) GetParent() types.Entity { return entries.parent }

func (entries *Arp_Nodes_Node_Entries) GetParentYangName() string { return "node" }

// Arp_Nodes_Node_Entries_Entry
// ARP entry
type Arp_Nodes_Node_Entries_Entry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP Address of ARP entry. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HardwareAddress interface{}
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetFilter() yfilter.YFilter { return entry.YFilter }

func (entry *Arp_Nodes_Node_Entries_Entry) SetFilter(yf yfilter.YFilter) { entry.YFilter = yf }

func (entry *Arp_Nodes_Node_Entries_Entry) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "media-type" { return "MediaType" }
    if yname == "state" { return "State" }
    if yname == "flag" { return "Flag" }
    if yname == "age" { return "Age" }
    if yname == "encapsulation-type" { return "EncapsulationType" }
    if yname == "hardware-length" { return "HardwareLength" }
    if yname == "hardware-address" { return "HardwareAddress" }
    return ""
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetSegmentPath() string {
    return "entry" + "[address='" + fmt.Sprintf("%v", entry.Address) + "']" + "[interface-name='" + fmt.Sprintf("%v", entry.InterfaceName) + "']"
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = entry.Address
    leafs["interface-name"] = entry.InterfaceName
    leafs["media-type"] = entry.MediaType
    leafs["state"] = entry.State
    leafs["flag"] = entry.Flag
    leafs["age"] = entry.Age
    leafs["encapsulation-type"] = entry.EncapsulationType
    leafs["hardware-length"] = entry.HardwareLength
    leafs["hardware-address"] = entry.HardwareAddress
    return leafs
}

func (entry *Arp_Nodes_Node_Entries_Entry) GetBundleName() string { return "cisco_ios_xr" }

func (entry *Arp_Nodes_Node_Entries_Entry) GetYangName() string { return "entry" }

func (entry *Arp_Nodes_Node_Entries_Entry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entry *Arp_Nodes_Node_Entries_Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entry *Arp_Nodes_Node_Entries_Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entry *Arp_Nodes_Node_Entries_Entry) SetParent(parent types.Entity) { entry.parent = parent }

func (entry *Arp_Nodes_Node_Entries_Entry) GetParent() types.Entity { return entry.parent }

func (entry *Arp_Nodes_Node_Entries_Entry) GetParentYangName() string { return "entries" }

// Arp_Nodes_Node_TrafficInterfaces
// ARP Traffic information per interface
type Arp_Nodes_Node_TrafficInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per interface traffic data. The type is slice of
    // Arp_Nodes_Node_TrafficInterfaces_TrafficInterface.
    TrafficInterface []Arp_Nodes_Node_TrafficInterfaces_TrafficInterface
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetFilter() yfilter.YFilter { return trafficInterfaces.YFilter }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) SetFilter(yf yfilter.YFilter) { trafficInterfaces.YFilter = yf }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetGoName(yname string) string {
    if yname == "traffic-interface" { return "TrafficInterface" }
    return ""
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetSegmentPath() string {
    return "traffic-interfaces"
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic-interface" {
        for _, c := range trafficInterfaces.TrafficInterface {
            if trafficInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arp_Nodes_Node_TrafficInterfaces_TrafficInterface{}
        trafficInterfaces.TrafficInterface = append(trafficInterfaces.TrafficInterface, child)
        return &trafficInterfaces.TrafficInterface[len(trafficInterfaces.TrafficInterface)-1]
    }
    return nil
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trafficInterfaces.TrafficInterface {
        children[trafficInterfaces.TrafficInterface[i].GetSegmentPath()] = &trafficInterfaces.TrafficInterface[i]
    }
    return children
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetYangName() string { return "traffic-interfaces" }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) SetParent(parent types.Entity) { trafficInterfaces.parent = parent }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetParent() types.Entity { return trafficInterfaces.parent }

func (trafficInterfaces *Arp_Nodes_Node_TrafficInterfaces) GetParentYangName() string { return "node" }

// Arp_Nodes_Node_TrafficInterfaces_TrafficInterface
// Per interface traffic data
type Arp_Nodes_Node_TrafficInterfaces_TrafficInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetFilter() yfilter.YFilter { return trafficInterface.YFilter }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) SetFilter(yf yfilter.YFilter) { trafficInterface.YFilter = yf }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "requests-received" { return "RequestsReceived" }
    if yname == "replies-received" { return "RepliesReceived" }
    if yname == "requests-sent" { return "RequestsSent" }
    if yname == "replies-sent" { return "RepliesSent" }
    if yname == "proxy-replies-sent" { return "ProxyRepliesSent" }
    if yname == "subscr-requests-received" { return "SubscrRequestsReceived" }
    if yname == "subscr-replies-sent" { return "SubscrRepliesSent" }
    if yname == "subscr-replies-gratg-sent" { return "SubscrRepliesGratgSent" }
    if yname == "local-proxy-replies-sent" { return "LocalProxyRepliesSent" }
    if yname == "gratuitous-replies-sent" { return "GratuitousRepliesSent" }
    if yname == "resolution-requests-received" { return "ResolutionRequestsReceived" }
    if yname == "resolution-replies-received" { return "ResolutionRepliesReceived" }
    if yname == "resolution-requests-dropped" { return "ResolutionRequestsDropped" }
    if yname == "out-of-memory-errors" { return "OutOfMemoryErrors" }
    if yname == "no-buffer-errors" { return "NoBufferErrors" }
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "dynamic-entries" { return "DynamicEntries" }
    if yname == "static-entries" { return "StaticEntries" }
    if yname == "alias-entries" { return "AliasEntries" }
    if yname == "interface-entries" { return "InterfaceEntries" }
    if yname == "standby-entries" { return "StandbyEntries" }
    if yname == "dhcp-entries" { return "DhcpEntries" }
    if yname == "vxlan-entries" { return "VxlanEntries" }
    if yname == "ip-packets-dropped-node" { return "IpPacketsDroppedNode" }
    if yname == "arp-packet-node-out-of-subnet" { return "ArpPacketNodeOutOfSubnet" }
    if yname == "ip-packets-dropped-interface" { return "IpPacketsDroppedInterface" }
    if yname == "arp-packet-interface-out-of-subnet" { return "ArpPacketInterfaceOutOfSubnet" }
    if yname == "arp-packet-unsolicited-packet" { return "ArpPacketUnsolicitedPacket" }
    if yname == "idb-structures" { return "IdbStructures" }
    return ""
}

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetSegmentPath() string {
    return "traffic-interface" + "[interface-name='" + fmt.Sprintf("%v", trafficInterface.InterfaceName) + "']"
}

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trafficInterface.InterfaceName
    leafs["requests-received"] = trafficInterface.RequestsReceived
    leafs["replies-received"] = trafficInterface.RepliesReceived
    leafs["requests-sent"] = trafficInterface.RequestsSent
    leafs["replies-sent"] = trafficInterface.RepliesSent
    leafs["proxy-replies-sent"] = trafficInterface.ProxyRepliesSent
    leafs["subscr-requests-received"] = trafficInterface.SubscrRequestsReceived
    leafs["subscr-replies-sent"] = trafficInterface.SubscrRepliesSent
    leafs["subscr-replies-gratg-sent"] = trafficInterface.SubscrRepliesGratgSent
    leafs["local-proxy-replies-sent"] = trafficInterface.LocalProxyRepliesSent
    leafs["gratuitous-replies-sent"] = trafficInterface.GratuitousRepliesSent
    leafs["resolution-requests-received"] = trafficInterface.ResolutionRequestsReceived
    leafs["resolution-replies-received"] = trafficInterface.ResolutionRepliesReceived
    leafs["resolution-requests-dropped"] = trafficInterface.ResolutionRequestsDropped
    leafs["out-of-memory-errors"] = trafficInterface.OutOfMemoryErrors
    leafs["no-buffer-errors"] = trafficInterface.NoBufferErrors
    leafs["total-entries"] = trafficInterface.TotalEntries
    leafs["dynamic-entries"] = trafficInterface.DynamicEntries
    leafs["static-entries"] = trafficInterface.StaticEntries
    leafs["alias-entries"] = trafficInterface.AliasEntries
    leafs["interface-entries"] = trafficInterface.InterfaceEntries
    leafs["standby-entries"] = trafficInterface.StandbyEntries
    leafs["dhcp-entries"] = trafficInterface.DhcpEntries
    leafs["vxlan-entries"] = trafficInterface.VxlanEntries
    leafs["ip-packets-dropped-node"] = trafficInterface.IpPacketsDroppedNode
    leafs["arp-packet-node-out-of-subnet"] = trafficInterface.ArpPacketNodeOutOfSubnet
    leafs["ip-packets-dropped-interface"] = trafficInterface.IpPacketsDroppedInterface
    leafs["arp-packet-interface-out-of-subnet"] = trafficInterface.ArpPacketInterfaceOutOfSubnet
    leafs["arp-packet-unsolicited-packet"] = trafficInterface.ArpPacketUnsolicitedPacket
    leafs["idb-structures"] = trafficInterface.IdbStructures
    return leafs
}

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetYangName() string { return "traffic-interface" }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) SetParent(parent types.Entity) { trafficInterface.parent = parent }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetParent() types.Entity { return trafficInterface.parent }

func (trafficInterface *Arp_Nodes_Node_TrafficInterfaces_TrafficInterface) GetParentYangName() string { return "traffic-interfaces" }

