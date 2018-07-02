// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-vrrp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   vrrp: VRRP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_vrrp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_vrrp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-vrrp-oper vrrp}", reflect.TypeOf(Vrrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-vrrp-oper:vrrp", reflect.TypeOf(Vrrp{}))
}

// VrrpStateChangeReason represents Vrrp state change reason
type VrrpStateChangeReason string

const (
    // BFD session down
    VrrpStateChangeReason_state_change_bfd_down VrrpStateChangeReason = "state-change-bfd-down"

    // Virtual IP configured
    VrrpStateChangeReason_state_change_virtual_ip_configured VrrpStateChangeReason = "state-change-virtual-ip-configured"

    // Interface IP update
    VrrpStateChangeReason_state_change_interface_ip VrrpStateChangeReason = "state-change-interface-ip"

    // Delay timer expired
    VrrpStateChangeReason_state_change_delay_timer VrrpStateChangeReason = "state-change-delay-timer"

    // Ready on startup
    VrrpStateChangeReason_state_change_startup VrrpStateChangeReason = "state-change-startup"

    // Interface Up update
    VrrpStateChangeReason_state_change_interface_up VrrpStateChangeReason = "state-change-interface-up"

    // Interface Down update
    VrrpStateChangeReason_state_change_interface_down VrrpStateChangeReason = "state-change-interface-down"

    // Master down timer expired
    VrrpStateChangeReason_state_change_master_down_timer VrrpStateChangeReason = "state-change-master-down-timer"

    // Higher priority advert received
    VrrpStateChangeReason_state_change_higher_priority_master VrrpStateChangeReason = "state-change-higher-priority-master"

    // FHRP Admin state change
    VrrpStateChangeReason_state_change_fhrp_admin VrrpStateChangeReason = "state-change-fhrp-admin"

    // Change of MGO parent session
    VrrpStateChangeReason_state_change_mgo_parent VrrpStateChangeReason = "state-change-mgo-parent"

    // Checkpoint update from Primary VRRP instance
    VrrpStateChangeReason_state_change_chkpt_update VrrpStateChangeReason = "state-change-chkpt-update"

    // Resync following ISSU primary event
    VrrpStateChangeReason_state_change_issu_resync VrrpStateChangeReason = "state-change-issu-resync"
)

// VrrpVmacState represents Vrrp vmac state
type VrrpVmacState string

const (
    // VMAC stored locally
    VrrpVmacState_stored VrrpVmacState = "stored"

    // VMAC reserved in mac table
    VrrpVmacState_reserved VrrpVmacState = "reserved"

    // VMAC active in mac table
    VrrpVmacState_active VrrpVmacState = "active"

    // VMAC not yet reserved in mac table
    VrrpVmacState_reserving VrrpVmacState = "reserving"
)

// VrrpBAf represents Vrrp b af
type VrrpBAf string

const (
    // IPv4 Address Family
    VrrpBAf_address_family_ipv4 VrrpBAf = "address-family-ipv4"

    // IPv6 Address Family
    VrrpBAf_address_family_ipv6 VrrpBAf = "address-family-ipv6"

    // Number of Adddress Families
    VrrpBAf_vrrp_baf_count VrrpBAf = "vrrp-baf-count"
)

// VrrpVipState represents Vrrp vip state
type VrrpVipState string

const (
    // Down
    VrrpVipState_virtual_ip_state_down VrrpVipState = "virtual-ip-state-down"

    // Up
    VrrpVipState_virtual_ip_state_up VrrpVipState = "virtual-ip-state-up"
)

// VrrpProtAuth represents Vrrp prot auth
type VrrpProtAuth string

const (
    // Down
    VrrpProtAuth_authentication_none VrrpProtAuth = "authentication-none"

    // Simple Text
    VrrpProtAuth_authentication_text VrrpProtAuth = "authentication-text"

    // IP header
    VrrpProtAuth_authentication_ip VrrpProtAuth = "authentication-ip"
)

// VrrpBfdSessionState represents Vrrp bfd session state
type VrrpBfdSessionState string

const (
    // None
    VrrpBfdSessionState_bfd_state_none VrrpBfdSessionState = "bfd-state-none"

    // Inactive
    VrrpBfdSessionState_bfd_state_inactive VrrpBfdSessionState = "bfd-state-inactive"

    // Up
    VrrpBfdSessionState_bfd_state_up VrrpBfdSessionState = "bfd-state-up"

    // Down
    VrrpBfdSessionState_bfd_state_down VrrpBfdSessionState = "bfd-state-down"
)

// VrrpBagProtocolState represents VRRP protocol state
type VrrpBagProtocolState string

const (
    // Initial
    VrrpBagProtocolState_state_initial VrrpBagProtocolState = "state-initial"

    // Backup
    VrrpBagProtocolState_state_backup VrrpBagProtocolState = "state-backup"

    // Master
    VrrpBagProtocolState_state_master VrrpBagProtocolState = "state-master"
)

// Vrrp
// VRRP operational data
type Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRRP summary statistics.
    Summary Vrrp_Summary

    // IPv6 VRRP configuration.
    Ipv6 Vrrp_Ipv6

    // IPv4 VRRP configuration.
    Ipv4 Vrrp_Ipv4

    // VRRP MGO Session information.
    MgoSessions Vrrp_MgoSessions
}

func (vrrp *Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "cisco_ios_xr"
    vrrp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-vrrp-oper"
    vrrp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-vrrp-oper:vrrp"
    vrrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Children.Append("summary", types.YChild{"Summary", &vrrp.Summary})
    vrrp.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &vrrp.Ipv6})
    vrrp.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &vrrp.Ipv4})
    vrrp.EntityData.Children.Append("mgo-sessions", types.YChild{"MgoSessions", &vrrp.MgoSessions})
    vrrp.EntityData.Leafs = types.NewOrderedMap()

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Vrrp_Summary
// VRRP summary statistics
type Vrrp_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of IPv4 sessions in MASTER (owner) state. The type is interface{}
    // with range: 0..4294967295.
    Ipv4SessionsMasterOwner interface{}

    // Number of IPv4 sessions in MASTER state. The type is interface{} with
    // range: 0..4294967295.
    Ipv4SessionsMaster interface{}

    // Number of IPv4 sessions in BACKUP state. The type is interface{} with
    // range: 0..4294967295.
    Ipv4SessionsBackup interface{}

    // Number of IPv4 sessions in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SessionsInit interface{}

    // Number of IPv4 slaves in MASTER state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesMaster interface{}

    // Number of IPv4 slaves in BACKUP state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesBackup interface{}

    // Number of IPv4 slaves in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesInit interface{}

    // Number of UP IPv4 Virtual IP Addresses on virtual routers in MASTER (owner)
    // state. The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesMasterOwnerUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on virtual routers in MASTER
    // (owner) state. The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesMasterOwnerDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on virtual routers in MASTER state.
    // The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesMasterUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on virtual routers in MASTER
    // state. The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesMasterDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on virtual routers in BACKUP state.
    // The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesBackupUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on virtual routers in BACKUP
    // state. The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesBackupDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on virtual routers in INIT state.
    // The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesInitUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on virtual routers in INIT state.
    // The type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesInitDown interface{}

    // Number of IPv6 sessions in MASTER (owner) state. The type is interface{}
    // with range: 0..4294967295.
    Ipv6SessionsMasterOwner interface{}

    // Number of IPv6 sessions in MASTER state. The type is interface{} with
    // range: 0..4294967295.
    Ipv6SessionsMaster interface{}

    // Number of IPv6 sessions in BACKUP state. The type is interface{} with
    // range: 0..4294967295.
    Ipv6SessionsBackup interface{}

    // Number of IPv6 sessions in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SessionsInit interface{}

    // Number of IPv6 slaves in MASTER state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesMaster interface{}

    // Number of IPv6 slaves in BACKUP state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesBackup interface{}

    // Number of IPv6 slaves in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesInit interface{}

    // Number of UP IPv6 Virtual IP Addresses on virtual routers in MASTER (owner)
    // state. The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesMasterOwnerUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on virtual routers in MASTER
    // (owner) state. The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesMasterOwnerDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on virtual routers in MASTER state.
    // The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesMasterUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on virtual routers in MASTER
    // state. The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesMasterDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on virtual routers in BACKUP state.
    // The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesBackupUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on virtual routers in BACKUP
    // state. The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesBackupDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on virtual routers in INIT state.
    // The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesInitUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on virtual routers in INIT state.
    // The type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesInitDown interface{}

    // Number of VRRP interfaces with IPv4 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv4StateUp interface{}

    // Number of VRRP interfaces with IPv4 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv4StateDown interface{}

    // Number of tracked interfaces with IPv4 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv4StateUp interface{}

    // Number of tracked interfaces with IPv4 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv4StateDown interface{}

    // Number of VRRP interfaces with IPv6 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv6StateUp interface{}

    // Number of VRRP interfaces with IPv6 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv6StateDown interface{}

    // Number of tracked interfaces with IPv6 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv6StateUp interface{}

    // Number of tracked interfaces with IPv6 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv6StateDown interface{}

    // Number of tracked objects in UP state. The type is interface{} with range:
    // 0..4294967295.
    TrackedObjectsStateUp interface{}

    // Number of tracked objects in DOWN state. The type is interface{} with
    // range: 0..4294967295.
    TrackedObjectsStateDown interface{}

    // Number of VRRP IPv4 BFD sessions in UP state. The type is interface{} with
    // range: 0..4294967295.
    BfdSessionsUp interface{}

    // Number of VRRP IPv4 BFD sessions in DOWN state. The type is interface{}
    // with range: 0..4294967295.
    BfdSessionsDown interface{}

    // Number of VRRP IPv4 BFD sessions in INACTIVE state. The type is interface{}
    // with range: 0..4294967295.
    BfdSessionInactive interface{}

    // Number of VRRP IPv6 BFD sessions in UP state. The type is interface{} with
    // range: 0..4294967295.
    Ipv6bfdSessionsUp interface{}

    // Number of VRRP IPv6 BFD sessions in DOWN state. The type is interface{}
    // with range: 0..4294967295.
    Ipv6bfdSessionsDown interface{}

    // Number of VRRP IPv6 BFD sessions in INACTIVE state. The type is interface{}
    // with range: 0..4294967295.
    Ipv6bfdSessionInactive interface{}
}

func (summary *Vrrp_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "vrrp"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("ipv4-sessions-master-owner", types.YLeaf{"Ipv4SessionsMasterOwner", summary.Ipv4SessionsMasterOwner})
    summary.EntityData.Leafs.Append("ipv4-sessions-master", types.YLeaf{"Ipv4SessionsMaster", summary.Ipv4SessionsMaster})
    summary.EntityData.Leafs.Append("ipv4-sessions-backup", types.YLeaf{"Ipv4SessionsBackup", summary.Ipv4SessionsBackup})
    summary.EntityData.Leafs.Append("ipv4-sessions-init", types.YLeaf{"Ipv4SessionsInit", summary.Ipv4SessionsInit})
    summary.EntityData.Leafs.Append("ipv4-slaves-master", types.YLeaf{"Ipv4SlavesMaster", summary.Ipv4SlavesMaster})
    summary.EntityData.Leafs.Append("ipv4-slaves-backup", types.YLeaf{"Ipv4SlavesBackup", summary.Ipv4SlavesBackup})
    summary.EntityData.Leafs.Append("ipv4-slaves-init", types.YLeaf{"Ipv4SlavesInit", summary.Ipv4SlavesInit})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-master-owner-up", types.YLeaf{"Ipv4VirtualIpAddressesMasterOwnerUp", summary.Ipv4VirtualIpAddressesMasterOwnerUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-master-owner-down", types.YLeaf{"Ipv4VirtualIpAddressesMasterOwnerDown", summary.Ipv4VirtualIpAddressesMasterOwnerDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-master-up", types.YLeaf{"Ipv4VirtualIpAddressesMasterUp", summary.Ipv4VirtualIpAddressesMasterUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-master-down", types.YLeaf{"Ipv4VirtualIpAddressesMasterDown", summary.Ipv4VirtualIpAddressesMasterDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-backup-up", types.YLeaf{"Ipv4VirtualIpAddressesBackupUp", summary.Ipv4VirtualIpAddressesBackupUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-backup-down", types.YLeaf{"Ipv4VirtualIpAddressesBackupDown", summary.Ipv4VirtualIpAddressesBackupDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-init-up", types.YLeaf{"Ipv4VirtualIpAddressesInitUp", summary.Ipv4VirtualIpAddressesInitUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-init-down", types.YLeaf{"Ipv4VirtualIpAddressesInitDown", summary.Ipv4VirtualIpAddressesInitDown})
    summary.EntityData.Leafs.Append("ipv6-sessions-master-owner", types.YLeaf{"Ipv6SessionsMasterOwner", summary.Ipv6SessionsMasterOwner})
    summary.EntityData.Leafs.Append("ipv6-sessions-master", types.YLeaf{"Ipv6SessionsMaster", summary.Ipv6SessionsMaster})
    summary.EntityData.Leafs.Append("ipv6-sessions-backup", types.YLeaf{"Ipv6SessionsBackup", summary.Ipv6SessionsBackup})
    summary.EntityData.Leafs.Append("ipv6-sessions-init", types.YLeaf{"Ipv6SessionsInit", summary.Ipv6SessionsInit})
    summary.EntityData.Leafs.Append("ipv6-slaves-master", types.YLeaf{"Ipv6SlavesMaster", summary.Ipv6SlavesMaster})
    summary.EntityData.Leafs.Append("ipv6-slaves-backup", types.YLeaf{"Ipv6SlavesBackup", summary.Ipv6SlavesBackup})
    summary.EntityData.Leafs.Append("ipv6-slaves-init", types.YLeaf{"Ipv6SlavesInit", summary.Ipv6SlavesInit})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-master-owner-up", types.YLeaf{"Ipv6VirtualIpAddressesMasterOwnerUp", summary.Ipv6VirtualIpAddressesMasterOwnerUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-master-owner-down", types.YLeaf{"Ipv6VirtualIpAddressesMasterOwnerDown", summary.Ipv6VirtualIpAddressesMasterOwnerDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-master-up", types.YLeaf{"Ipv6VirtualIpAddressesMasterUp", summary.Ipv6VirtualIpAddressesMasterUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-master-down", types.YLeaf{"Ipv6VirtualIpAddressesMasterDown", summary.Ipv6VirtualIpAddressesMasterDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-backup-up", types.YLeaf{"Ipv6VirtualIpAddressesBackupUp", summary.Ipv6VirtualIpAddressesBackupUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-backup-down", types.YLeaf{"Ipv6VirtualIpAddressesBackupDown", summary.Ipv6VirtualIpAddressesBackupDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-init-up", types.YLeaf{"Ipv6VirtualIpAddressesInitUp", summary.Ipv6VirtualIpAddressesInitUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-init-down", types.YLeaf{"Ipv6VirtualIpAddressesInitDown", summary.Ipv6VirtualIpAddressesInitDown})
    summary.EntityData.Leafs.Append("interfaces-ipv4-state-up", types.YLeaf{"InterfacesIpv4StateUp", summary.InterfacesIpv4StateUp})
    summary.EntityData.Leafs.Append("interfaces-ipv4-state-down", types.YLeaf{"InterfacesIpv4StateDown", summary.InterfacesIpv4StateDown})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv4-state-up", types.YLeaf{"TrackedInterfacesIpv4StateUp", summary.TrackedInterfacesIpv4StateUp})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv4-state-down", types.YLeaf{"TrackedInterfacesIpv4StateDown", summary.TrackedInterfacesIpv4StateDown})
    summary.EntityData.Leafs.Append("interfaces-ipv6-state-up", types.YLeaf{"InterfacesIpv6StateUp", summary.InterfacesIpv6StateUp})
    summary.EntityData.Leafs.Append("interfaces-ipv6-state-down", types.YLeaf{"InterfacesIpv6StateDown", summary.InterfacesIpv6StateDown})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv6-state-up", types.YLeaf{"TrackedInterfacesIpv6StateUp", summary.TrackedInterfacesIpv6StateUp})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv6-state-down", types.YLeaf{"TrackedInterfacesIpv6StateDown", summary.TrackedInterfacesIpv6StateDown})
    summary.EntityData.Leafs.Append("tracked-objects-state-up", types.YLeaf{"TrackedObjectsStateUp", summary.TrackedObjectsStateUp})
    summary.EntityData.Leafs.Append("tracked-objects-state-down", types.YLeaf{"TrackedObjectsStateDown", summary.TrackedObjectsStateDown})
    summary.EntityData.Leafs.Append("bfd-sessions-up", types.YLeaf{"BfdSessionsUp", summary.BfdSessionsUp})
    summary.EntityData.Leafs.Append("bfd-sessions-down", types.YLeaf{"BfdSessionsDown", summary.BfdSessionsDown})
    summary.EntityData.Leafs.Append("bfd-session-inactive", types.YLeaf{"BfdSessionInactive", summary.BfdSessionInactive})
    summary.EntityData.Leafs.Append("ipv6bfd-sessions-up", types.YLeaf{"Ipv6bfdSessionsUp", summary.Ipv6bfdSessionsUp})
    summary.EntityData.Leafs.Append("ipv6bfd-sessions-down", types.YLeaf{"Ipv6bfdSessionsDown", summary.Ipv6bfdSessionsDown})
    summary.EntityData.Leafs.Append("ipv6bfd-session-inactive", types.YLeaf{"Ipv6bfdSessionInactive", summary.Ipv6bfdSessionInactive})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Vrrp_Ipv6
// IPv6 VRRP configuration
type Vrrp_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP tracked item table.
    TrackItems Vrrp_Ipv6_TrackItems

    // The VRRP virtual router table.
    VirtualRouters Vrrp_Ipv6_VirtualRouters

    // The VRRP interface table.
    Interfaces Vrrp_Ipv6_Interfaces
}

func (ipv6 *Vrrp_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "vrrp"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("track-items", types.YChild{"TrackItems", &ipv6.TrackItems})
    ipv6.EntityData.Children.Append("virtual-routers", types.YChild{"VirtualRouters", &ipv6.VirtualRouters})
    ipv6.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv6.Interfaces})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Vrrp_Ipv6_TrackItems
// The VRRP tracked item table
type Vrrp_Ipv6_TrackItems struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A configured VRRP IP address entry. The type is slice of
    // Vrrp_Ipv6_TrackItems_TrackItem.
    TrackItem []*Vrrp_Ipv6_TrackItems_TrackItem
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetEntityData() *types.CommonEntityData {
    trackItems.EntityData.YFilter = trackItems.YFilter
    trackItems.EntityData.YangName = "track-items"
    trackItems.EntityData.BundleName = "cisco_ios_xr"
    trackItems.EntityData.ParentYangName = "ipv6"
    trackItems.EntityData.SegmentPath = "track-items"
    trackItems.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackItems.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackItems.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackItems.EntityData.Children = types.NewOrderedMap()
    trackItems.EntityData.Children.Append("track-item", types.YChild{"TrackItem", nil})
    for i := range trackItems.TrackItem {
        trackItems.EntityData.Children.Append(types.GetSegmentPath(trackItems.TrackItem[i]), types.YChild{"TrackItem", trackItems.TrackItem[i]})
    }
    trackItems.EntityData.Leafs = types.NewOrderedMap()

    trackItems.EntityData.YListKeys = []string {}

    return &(trackItems.EntityData)
}

// Vrrp_Ipv6_TrackItems_TrackItem
// A configured VRRP IP address entry
type Vrrp_Ipv6_TrackItems_TrackItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name to track. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: 0..4294967295.
    VirtualRouterId interface{}

    // This attribute is a key. The name of the tracked interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    TrackedInterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    VirtualRouterIdXr interface{}

    // Type of tracked item. The type is interface{} with range: 0..65535.
    TrackedItemType interface{}

    // Tracked item index. The type is string with length: 0..32.
    TrackedItemIndex interface{}

    // State of the tracked item. The type is interface{} with range: 0..255.
    State interface{}

    // Priority weight of item. The type is interface{} with range: 0..255.
    Priority interface{}
}

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetEntityData() *types.CommonEntityData {
    trackItem.EntityData.YFilter = trackItem.YFilter
    trackItem.EntityData.YangName = "track-item"
    trackItem.EntityData.BundleName = "cisco_ios_xr"
    trackItem.EntityData.ParentYangName = "track-items"
    trackItem.EntityData.SegmentPath = "track-item" + types.AddKeyToken(trackItem.InterfaceName, "interface-name") + types.AddKeyToken(trackItem.VirtualRouterId, "virtual-router-id") + types.AddKeyToken(trackItem.TrackedInterfaceName, "tracked-interface-name")
    trackItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackItem.EntityData.Children = types.NewOrderedMap()
    trackItem.EntityData.Leafs = types.NewOrderedMap()
    trackItem.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", trackItem.InterfaceName})
    trackItem.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", trackItem.VirtualRouterId})
    trackItem.EntityData.Leafs.Append("tracked-interface-name", types.YLeaf{"TrackedInterfaceName", trackItem.TrackedInterfaceName})
    trackItem.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trackItem.Interface})
    trackItem.EntityData.Leafs.Append("virtual-router-id-xr", types.YLeaf{"VirtualRouterIdXr", trackItem.VirtualRouterIdXr})
    trackItem.EntityData.Leafs.Append("tracked-item-type", types.YLeaf{"TrackedItemType", trackItem.TrackedItemType})
    trackItem.EntityData.Leafs.Append("tracked-item-index", types.YLeaf{"TrackedItemIndex", trackItem.TrackedItemIndex})
    trackItem.EntityData.Leafs.Append("state", types.YLeaf{"State", trackItem.State})
    trackItem.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", trackItem.Priority})

    trackItem.EntityData.YListKeys = []string {"InterfaceName", "VirtualRouterId", "TrackedInterfaceName"}

    return &(trackItem.EntityData)
}

// Vrrp_Ipv6_VirtualRouters
// The VRRP virtual router table
type Vrrp_Ipv6_VirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP virtual router. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter.
    VirtualRouter []*Vrrp_Ipv6_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetEntityData() *types.CommonEntityData {
    virtualRouters.EntityData.YFilter = virtualRouters.YFilter
    virtualRouters.EntityData.YangName = "virtual-routers"
    virtualRouters.EntityData.BundleName = "cisco_ios_xr"
    virtualRouters.EntityData.ParentYangName = "ipv6"
    virtualRouters.EntityData.SegmentPath = "virtual-routers"
    virtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouters.EntityData.Children = types.NewOrderedMap()
    virtualRouters.EntityData.Children.Append("virtual-router", types.YChild{"VirtualRouter", nil})
    for i := range virtualRouters.VirtualRouter {
        virtualRouters.EntityData.Children.Append(types.GetSegmentPath(virtualRouters.VirtualRouter[i]), types.YChild{"VirtualRouter", virtualRouters.VirtualRouter[i]})
    }
    virtualRouters.EntityData.Leafs = types.NewOrderedMap()

    virtualRouters.EntityData.YListKeys = []string {}

    return &(virtualRouters.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter
// A VRRP virtual router
type Vrrp_Ipv6_VirtualRouters_VirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: 0..4294967295.
    VirtualRouterId interface{}

    // IM Interface Name. The type is string with length: 0..64.
    InterfaceNameXr interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    VirtualRouterIdXr interface{}

    // VRRP Protocol Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Address family. The type is VrrpBAf.
    AddressFamily interface{}

    // Session Name. The type is string with length: 0..16.
    SessionName interface{}

    // Number of slaves following state. The type is interface{} with range:
    // 0..4294967295.
    Slaves interface{}

    // Group is a slave group. The type is bool.
    IsSlave interface{}

    // Followed Session Name. The type is string with length: 0..16.
    FollowedSessionName interface{}

    // Configured VRRP secondary address count. The type is interface{} with
    // range: 0..255.
    SecondaryAddressCount interface{}

    // Operational VRRP address count. The type is interface{} with range: 0..255.
    OperationalAddressCount interface{}

    // Configured IPv4 Primary address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryVirtualIp interface{}

    // Configured but Down VRRP address count. The type is interface{} with range:
    // 0..255.
    ConfiguredDownAddressCount interface{}

    // Virtual linklocal IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinklocalIpv6Address interface{}

    // State of primary IP address. The type is VrrpVipState.
    PrimaryState interface{}

    // Master router real IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    MasterIpAddress interface{}

    // Master router real IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    MasterIpv6Address interface{}

    // Master router priority. The type is interface{} with range: 0..255.
    MasterPriority interface{}

    // VRRP state. The type is VrrpBagProtocolState.
    VrrpState interface{}

    // Authentication type. The type is VrrpProtAuth.
    AuthenticationType interface{}

    // Authentication data. The type is string.
    AuthenticationString interface{}

    // Configured advertize time. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredAdvertizeTime interface{}

    // Operational advertize time. The type is interface{} with range:
    // 0..4294967295.
    OperAdvertizeTime interface{}

    // Minimum delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    MinDelayTime interface{}

    // Reload delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ReloadDelayTime interface{}

    // Delay timer running flag. The type is bool.
    DelayTimerFlag interface{}

    // Delay timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DelayTimerSecs interface{}

    // Delay timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    DelayTimerMsecs interface{}

    // Text authentication configured flag. The type is bool.
    AuthenticationFlag interface{}

    // Configured timers forced flag. The type is bool.
    ForceTimerFlag interface{}

    // Preempt configured flag. The type is bool.
    PreemptFlag interface{}

    // IP address owner flag. The type is bool.
    IpAddressOwnerFlag interface{}

    // Is accept mode. The type is bool.
    IsAcceptMode interface{}

    // Preempt delay time. The type is interface{} with range: 0..65535.
    PreemptDelayTime interface{}

    // Configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // Operational priority. The type is interface{} with range: 0..255.
    OperationalPriority interface{}

    // Priority decrement. The type is interface{} with range: 0..4294967295.
    PriorityDecrement interface{}

    // Number of items tracked. The type is interface{} with range: 0..4294967295.
    TrackedInterfaceCount interface{}

    // Number of tracked items up. The type is interface{} with range:
    // 0..4294967295.
    TrackedInterfaceUpCount interface{}

    // Number of tracked items. The type is interface{} with range: 0..4294967295.
    TrackedItemCount interface{}

    // Number of tracked items in UP state. The type is interface{} with range:
    // 0..4294967295.
    TrackedItemUpCount interface{}

    // Time in current state secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TimeInCurrentState interface{}

    // Number of state changes. The type is interface{} with range: 0..4294967295.
    StateChangeCount interface{}

    // Time vrouter is up in centiseconds. The type is interface{} with range:
    // 0..4294967295. Units are centisecond.
    TimeVrouterUp interface{}

    // No. of times become Master. The type is interface{} with range:
    // 0..4294967295.
    MasterCount interface{}

    // No. of advertisements received. The type is interface{} with range:
    // 0..4294967295.
    AdvertsReceivedCount interface{}

    // Advertise interval errors. The type is interface{} with range:
    // 0..4294967295.
    AdvertIntervalErrorCount interface{}

    // No. of advertisements sent. The type is interface{} with range:
    // 0..4294967295.
    AdvertsSentCount interface{}

    // Authentication failures. The type is interface{} with range: 0..4294967295.
    AuthenticationFailCount interface{}

    // TTL errors. The type is interface{} with range: 0..4294967295.
    TtlErrorCount interface{}

    // No. priority 0 received. The type is interface{} with range: 0..4294967295.
    PriorityZeroReceivedCount interface{}

    // No. priority 0 sent. The type is interface{} with range: 0..4294967295.
    PriorityZeroSentCount interface{}

    // Invalid packets received. The type is interface{} with range:
    // 0..4294967295.
    InvalidPacketCount interface{}

    // Address list errors. The type is interface{} with range: 0..4294967295.
    AddressListErrorCount interface{}

    // Invalid authentication type. The type is interface{} with range:
    // 0..4294967295.
    InvalidAuthTypeCount interface{}

    // Authentication type mismatches. The type is interface{} with range:
    // 0..4294967295.
    AuthTypeMismatchCount interface{}

    // Packet length errors. The type is interface{} with range: 0..4294967295.
    PktLengthErrorsCount interface{}

    // Time since a statistics discontinuity in ticks (10ns units). The type is
    // interface{} with range: 0..4294967295.
    TimeStatsDiscontinuity interface{}

    // BFD session state. The type is VrrpBfdSessionState.
    BfdSessionState interface{}

    // BFD packet send interval. The type is interface{} with range:
    // 0..4294967295.
    BfdInterval interface{}

    // BFD multiplier. The type is interface{} with range: 0..4294967295.
    BfdMultiplier interface{}

    // BFD configured remote IP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BfdCfgRemoteIp interface{}

    // BFD configured remote IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    BfdConfiguredRemoteIpv6Address interface{}

    // Whether state recovered from checkpoint. The type is bool.
    StateFromCheckpoint interface{}

    // The Interface Primary IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    InterfaceIpv4Address interface{}

    // The Interface linklocal IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    InterfaceIpv6Address interface{}

    // Virtual mac address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // Virtual mac address state. The type is VrrpVmacState.
    VirtualMacAddressState interface{}

    // Operational IPv4 VRRP addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OperationalAddress []interface{}

    // IPv4 Configured but Down VRRP addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4ConfiguredDownAddress []interface{}

    // Time last resign was sent.
    ResignSentTime Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime

    // Time last resign was received.
    ResignReceivedTime Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime

    // IPv6 Operational VRRP addresses. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress.
    Ipv6OperationalAddress []*Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress

    // IPv6 Configured but Down VRRP addresses. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress.
    Ipv6ConfiguredDownAddress []*Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress

    // Track Item Info. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter_TrackItemInfo.
    TrackItemInfo []*Vrrp_Ipv6_VirtualRouters_VirtualRouter_TrackItemInfo

    // State change history. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory.
    StateChangeHistory []*Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetEntityData() *types.CommonEntityData {
    virtualRouter.EntityData.YFilter = virtualRouter.YFilter
    virtualRouter.EntityData.YangName = "virtual-router"
    virtualRouter.EntityData.BundleName = "cisco_ios_xr"
    virtualRouter.EntityData.ParentYangName = "virtual-routers"
    virtualRouter.EntityData.SegmentPath = "virtual-router" + types.AddKeyToken(virtualRouter.InterfaceName, "interface-name") + types.AddKeyToken(virtualRouter.VirtualRouterId, "virtual-router-id")
    virtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouter.EntityData.Children = types.NewOrderedMap()
    virtualRouter.EntityData.Children.Append("resign-sent-time", types.YChild{"ResignSentTime", &virtualRouter.ResignSentTime})
    virtualRouter.EntityData.Children.Append("resign-received-time", types.YChild{"ResignReceivedTime", &virtualRouter.ResignReceivedTime})
    virtualRouter.EntityData.Children.Append("ipv6-operational-address", types.YChild{"Ipv6OperationalAddress", nil})
    for i := range virtualRouter.Ipv6OperationalAddress {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.Ipv6OperationalAddress[i]), types.YChild{"Ipv6OperationalAddress", virtualRouter.Ipv6OperationalAddress[i]})
    }
    virtualRouter.EntityData.Children.Append("ipv6-configured-down-address", types.YChild{"Ipv6ConfiguredDownAddress", nil})
    for i := range virtualRouter.Ipv6ConfiguredDownAddress {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.Ipv6ConfiguredDownAddress[i]), types.YChild{"Ipv6ConfiguredDownAddress", virtualRouter.Ipv6ConfiguredDownAddress[i]})
    }
    virtualRouter.EntityData.Children.Append("track-item-info", types.YChild{"TrackItemInfo", nil})
    for i := range virtualRouter.TrackItemInfo {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.TrackItemInfo[i]), types.YChild{"TrackItemInfo", virtualRouter.TrackItemInfo[i]})
    }
    virtualRouter.EntityData.Children.Append("state-change-history", types.YChild{"StateChangeHistory", nil})
    for i := range virtualRouter.StateChangeHistory {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.StateChangeHistory[i]), types.YChild{"StateChangeHistory", virtualRouter.StateChangeHistory[i]})
    }
    virtualRouter.EntityData.Leafs = types.NewOrderedMap()
    virtualRouter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", virtualRouter.InterfaceName})
    virtualRouter.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", virtualRouter.VirtualRouterId})
    virtualRouter.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", virtualRouter.InterfaceNameXr})
    virtualRouter.EntityData.Leafs.Append("virtual-router-id-xr", types.YLeaf{"VirtualRouterIdXr", virtualRouter.VirtualRouterIdXr})
    virtualRouter.EntityData.Leafs.Append("version", types.YLeaf{"Version", virtualRouter.Version})
    virtualRouter.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", virtualRouter.AddressFamily})
    virtualRouter.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", virtualRouter.SessionName})
    virtualRouter.EntityData.Leafs.Append("slaves", types.YLeaf{"Slaves", virtualRouter.Slaves})
    virtualRouter.EntityData.Leafs.Append("is-slave", types.YLeaf{"IsSlave", virtualRouter.IsSlave})
    virtualRouter.EntityData.Leafs.Append("followed-session-name", types.YLeaf{"FollowedSessionName", virtualRouter.FollowedSessionName})
    virtualRouter.EntityData.Leafs.Append("secondary-address-count", types.YLeaf{"SecondaryAddressCount", virtualRouter.SecondaryAddressCount})
    virtualRouter.EntityData.Leafs.Append("operational-address-count", types.YLeaf{"OperationalAddressCount", virtualRouter.OperationalAddressCount})
    virtualRouter.EntityData.Leafs.Append("primary-virtual-ip", types.YLeaf{"PrimaryVirtualIp", virtualRouter.PrimaryVirtualIp})
    virtualRouter.EntityData.Leafs.Append("configured-down-address-count", types.YLeaf{"ConfiguredDownAddressCount", virtualRouter.ConfiguredDownAddressCount})
    virtualRouter.EntityData.Leafs.Append("virtual-linklocal-ipv6-address", types.YLeaf{"VirtualLinklocalIpv6Address", virtualRouter.VirtualLinklocalIpv6Address})
    virtualRouter.EntityData.Leafs.Append("primary-state", types.YLeaf{"PrimaryState", virtualRouter.PrimaryState})
    virtualRouter.EntityData.Leafs.Append("master-ip-address", types.YLeaf{"MasterIpAddress", virtualRouter.MasterIpAddress})
    virtualRouter.EntityData.Leafs.Append("master-ipv6-address", types.YLeaf{"MasterIpv6Address", virtualRouter.MasterIpv6Address})
    virtualRouter.EntityData.Leafs.Append("master-priority", types.YLeaf{"MasterPriority", virtualRouter.MasterPriority})
    virtualRouter.EntityData.Leafs.Append("vrrp-state", types.YLeaf{"VrrpState", virtualRouter.VrrpState})
    virtualRouter.EntityData.Leafs.Append("authentication-type", types.YLeaf{"AuthenticationType", virtualRouter.AuthenticationType})
    virtualRouter.EntityData.Leafs.Append("authentication-string", types.YLeaf{"AuthenticationString", virtualRouter.AuthenticationString})
    virtualRouter.EntityData.Leafs.Append("configured-advertize-time", types.YLeaf{"ConfiguredAdvertizeTime", virtualRouter.ConfiguredAdvertizeTime})
    virtualRouter.EntityData.Leafs.Append("oper-advertize-time", types.YLeaf{"OperAdvertizeTime", virtualRouter.OperAdvertizeTime})
    virtualRouter.EntityData.Leafs.Append("min-delay-time", types.YLeaf{"MinDelayTime", virtualRouter.MinDelayTime})
    virtualRouter.EntityData.Leafs.Append("reload-delay-time", types.YLeaf{"ReloadDelayTime", virtualRouter.ReloadDelayTime})
    virtualRouter.EntityData.Leafs.Append("delay-timer-flag", types.YLeaf{"DelayTimerFlag", virtualRouter.DelayTimerFlag})
    virtualRouter.EntityData.Leafs.Append("delay-timer-secs", types.YLeaf{"DelayTimerSecs", virtualRouter.DelayTimerSecs})
    virtualRouter.EntityData.Leafs.Append("delay-timer-msecs", types.YLeaf{"DelayTimerMsecs", virtualRouter.DelayTimerMsecs})
    virtualRouter.EntityData.Leafs.Append("authentication-flag", types.YLeaf{"AuthenticationFlag", virtualRouter.AuthenticationFlag})
    virtualRouter.EntityData.Leafs.Append("force-timer-flag", types.YLeaf{"ForceTimerFlag", virtualRouter.ForceTimerFlag})
    virtualRouter.EntityData.Leafs.Append("preempt-flag", types.YLeaf{"PreemptFlag", virtualRouter.PreemptFlag})
    virtualRouter.EntityData.Leafs.Append("ip-address-owner-flag", types.YLeaf{"IpAddressOwnerFlag", virtualRouter.IpAddressOwnerFlag})
    virtualRouter.EntityData.Leafs.Append("is-accept-mode", types.YLeaf{"IsAcceptMode", virtualRouter.IsAcceptMode})
    virtualRouter.EntityData.Leafs.Append("preempt-delay-time", types.YLeaf{"PreemptDelayTime", virtualRouter.PreemptDelayTime})
    virtualRouter.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", virtualRouter.ConfiguredPriority})
    virtualRouter.EntityData.Leafs.Append("operational-priority", types.YLeaf{"OperationalPriority", virtualRouter.OperationalPriority})
    virtualRouter.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", virtualRouter.PriorityDecrement})
    virtualRouter.EntityData.Leafs.Append("tracked-interface-count", types.YLeaf{"TrackedInterfaceCount", virtualRouter.TrackedInterfaceCount})
    virtualRouter.EntityData.Leafs.Append("tracked-interface-up-count", types.YLeaf{"TrackedInterfaceUpCount", virtualRouter.TrackedInterfaceUpCount})
    virtualRouter.EntityData.Leafs.Append("tracked-item-count", types.YLeaf{"TrackedItemCount", virtualRouter.TrackedItemCount})
    virtualRouter.EntityData.Leafs.Append("tracked-item-up-count", types.YLeaf{"TrackedItemUpCount", virtualRouter.TrackedItemUpCount})
    virtualRouter.EntityData.Leafs.Append("time-in-current-state", types.YLeaf{"TimeInCurrentState", virtualRouter.TimeInCurrentState})
    virtualRouter.EntityData.Leafs.Append("state-change-count", types.YLeaf{"StateChangeCount", virtualRouter.StateChangeCount})
    virtualRouter.EntityData.Leafs.Append("time-vrouter-up", types.YLeaf{"TimeVrouterUp", virtualRouter.TimeVrouterUp})
    virtualRouter.EntityData.Leafs.Append("master-count", types.YLeaf{"MasterCount", virtualRouter.MasterCount})
    virtualRouter.EntityData.Leafs.Append("adverts-received-count", types.YLeaf{"AdvertsReceivedCount", virtualRouter.AdvertsReceivedCount})
    virtualRouter.EntityData.Leafs.Append("advert-interval-error-count", types.YLeaf{"AdvertIntervalErrorCount", virtualRouter.AdvertIntervalErrorCount})
    virtualRouter.EntityData.Leafs.Append("adverts-sent-count", types.YLeaf{"AdvertsSentCount", virtualRouter.AdvertsSentCount})
    virtualRouter.EntityData.Leafs.Append("authentication-fail-count", types.YLeaf{"AuthenticationFailCount", virtualRouter.AuthenticationFailCount})
    virtualRouter.EntityData.Leafs.Append("ttl-error-count", types.YLeaf{"TtlErrorCount", virtualRouter.TtlErrorCount})
    virtualRouter.EntityData.Leafs.Append("priority-zero-received-count", types.YLeaf{"PriorityZeroReceivedCount", virtualRouter.PriorityZeroReceivedCount})
    virtualRouter.EntityData.Leafs.Append("priority-zero-sent-count", types.YLeaf{"PriorityZeroSentCount", virtualRouter.PriorityZeroSentCount})
    virtualRouter.EntityData.Leafs.Append("invalid-packet-count", types.YLeaf{"InvalidPacketCount", virtualRouter.InvalidPacketCount})
    virtualRouter.EntityData.Leafs.Append("address-list-error-count", types.YLeaf{"AddressListErrorCount", virtualRouter.AddressListErrorCount})
    virtualRouter.EntityData.Leafs.Append("invalid-auth-type-count", types.YLeaf{"InvalidAuthTypeCount", virtualRouter.InvalidAuthTypeCount})
    virtualRouter.EntityData.Leafs.Append("auth-type-mismatch-count", types.YLeaf{"AuthTypeMismatchCount", virtualRouter.AuthTypeMismatchCount})
    virtualRouter.EntityData.Leafs.Append("pkt-length-errors-count", types.YLeaf{"PktLengthErrorsCount", virtualRouter.PktLengthErrorsCount})
    virtualRouter.EntityData.Leafs.Append("time-stats-discontinuity", types.YLeaf{"TimeStatsDiscontinuity", virtualRouter.TimeStatsDiscontinuity})
    virtualRouter.EntityData.Leafs.Append("bfd-session-state", types.YLeaf{"BfdSessionState", virtualRouter.BfdSessionState})
    virtualRouter.EntityData.Leafs.Append("bfd-interval", types.YLeaf{"BfdInterval", virtualRouter.BfdInterval})
    virtualRouter.EntityData.Leafs.Append("bfd-multiplier", types.YLeaf{"BfdMultiplier", virtualRouter.BfdMultiplier})
    virtualRouter.EntityData.Leafs.Append("bfd-cfg-remote-ip", types.YLeaf{"BfdCfgRemoteIp", virtualRouter.BfdCfgRemoteIp})
    virtualRouter.EntityData.Leafs.Append("bfd-configured-remote-ipv6-address", types.YLeaf{"BfdConfiguredRemoteIpv6Address", virtualRouter.BfdConfiguredRemoteIpv6Address})
    virtualRouter.EntityData.Leafs.Append("state-from-checkpoint", types.YLeaf{"StateFromCheckpoint", virtualRouter.StateFromCheckpoint})
    virtualRouter.EntityData.Leafs.Append("interface-ipv4-address", types.YLeaf{"InterfaceIpv4Address", virtualRouter.InterfaceIpv4Address})
    virtualRouter.EntityData.Leafs.Append("interface-ipv6-address", types.YLeaf{"InterfaceIpv6Address", virtualRouter.InterfaceIpv6Address})
    virtualRouter.EntityData.Leafs.Append("virtual-mac-address", types.YLeaf{"VirtualMacAddress", virtualRouter.VirtualMacAddress})
    virtualRouter.EntityData.Leafs.Append("virtual-mac-address-state", types.YLeaf{"VirtualMacAddressState", virtualRouter.VirtualMacAddressState})
    virtualRouter.EntityData.Leafs.Append("operational-address", types.YLeaf{"OperationalAddress", virtualRouter.OperationalAddress})
    virtualRouter.EntityData.Leafs.Append("ipv4-configured-down-address", types.YLeaf{"Ipv4ConfiguredDownAddress", virtualRouter.Ipv4ConfiguredDownAddress})

    virtualRouter.EntityData.YListKeys = []string {"InterfaceName", "VirtualRouterId"}

    return &(virtualRouter.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime
// Time last resign was sent
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetEntityData() *types.CommonEntityData {
    resignSentTime.EntityData.YFilter = resignSentTime.YFilter
    resignSentTime.EntityData.YangName = "resign-sent-time"
    resignSentTime.EntityData.BundleName = "cisco_ios_xr"
    resignSentTime.EntityData.ParentYangName = "virtual-router"
    resignSentTime.EntityData.SegmentPath = "resign-sent-time"
    resignSentTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resignSentTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resignSentTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resignSentTime.EntityData.Children = types.NewOrderedMap()
    resignSentTime.EntityData.Leafs = types.NewOrderedMap()
    resignSentTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", resignSentTime.Seconds})
    resignSentTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", resignSentTime.Nanoseconds})

    resignSentTime.EntityData.YListKeys = []string {}

    return &(resignSentTime.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime
// Time last resign was received
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetEntityData() *types.CommonEntityData {
    resignReceivedTime.EntityData.YFilter = resignReceivedTime.YFilter
    resignReceivedTime.EntityData.YangName = "resign-received-time"
    resignReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    resignReceivedTime.EntityData.ParentYangName = "virtual-router"
    resignReceivedTime.EntityData.SegmentPath = "resign-received-time"
    resignReceivedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resignReceivedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resignReceivedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resignReceivedTime.EntityData.Children = types.NewOrderedMap()
    resignReceivedTime.EntityData.Leafs = types.NewOrderedMap()
    resignReceivedTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", resignReceivedTime.Seconds})
    resignReceivedTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", resignReceivedTime.Nanoseconds})

    resignReceivedTime.EntityData.YListKeys = []string {}

    return &(resignReceivedTime.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress
// IPv6 Operational VRRP addresses
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetEntityData() *types.CommonEntityData {
    ipv6OperationalAddress.EntityData.YFilter = ipv6OperationalAddress.YFilter
    ipv6OperationalAddress.EntityData.YangName = "ipv6-operational-address"
    ipv6OperationalAddress.EntityData.BundleName = "cisco_ios_xr"
    ipv6OperationalAddress.EntityData.ParentYangName = "virtual-router"
    ipv6OperationalAddress.EntityData.SegmentPath = "ipv6-operational-address"
    ipv6OperationalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6OperationalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6OperationalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6OperationalAddress.EntityData.Children = types.NewOrderedMap()
    ipv6OperationalAddress.EntityData.Leafs = types.NewOrderedMap()
    ipv6OperationalAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6OperationalAddress.Ipv6Address})

    ipv6OperationalAddress.EntityData.YListKeys = []string {}

    return &(ipv6OperationalAddress.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress
// IPv6 Configured but Down VRRP addresses
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetEntityData() *types.CommonEntityData {
    ipv6ConfiguredDownAddress.EntityData.YFilter = ipv6ConfiguredDownAddress.YFilter
    ipv6ConfiguredDownAddress.EntityData.YangName = "ipv6-configured-down-address"
    ipv6ConfiguredDownAddress.EntityData.BundleName = "cisco_ios_xr"
    ipv6ConfiguredDownAddress.EntityData.ParentYangName = "virtual-router"
    ipv6ConfiguredDownAddress.EntityData.SegmentPath = "ipv6-configured-down-address"
    ipv6ConfiguredDownAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6ConfiguredDownAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6ConfiguredDownAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6ConfiguredDownAddress.EntityData.Children = types.NewOrderedMap()
    ipv6ConfiguredDownAddress.EntityData.Leafs = types.NewOrderedMap()
    ipv6ConfiguredDownAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6ConfiguredDownAddress.Ipv6Address})

    ipv6ConfiguredDownAddress.EntityData.YListKeys = []string {}

    return &(ipv6ConfiguredDownAddress.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_TrackItemInfo
// Track Item Info
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_TrackItemInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    VirtualRouterIdXr interface{}

    // Type of tracked item. The type is interface{} with range: 0..65535.
    TrackedItemType interface{}

    // Tracked item index. The type is string with length: 0..32.
    TrackedItemIndex interface{}

    // State of the tracked item. The type is interface{} with range: 0..255.
    State interface{}

    // Priority weight of item. The type is interface{} with range: 0..255.
    Priority interface{}
}

func (trackItemInfo *Vrrp_Ipv6_VirtualRouters_VirtualRouter_TrackItemInfo) GetEntityData() *types.CommonEntityData {
    trackItemInfo.EntityData.YFilter = trackItemInfo.YFilter
    trackItemInfo.EntityData.YangName = "track-item-info"
    trackItemInfo.EntityData.BundleName = "cisco_ios_xr"
    trackItemInfo.EntityData.ParentYangName = "virtual-router"
    trackItemInfo.EntityData.SegmentPath = "track-item-info"
    trackItemInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackItemInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackItemInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackItemInfo.EntityData.Children = types.NewOrderedMap()
    trackItemInfo.EntityData.Leafs = types.NewOrderedMap()
    trackItemInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trackItemInfo.Interface})
    trackItemInfo.EntityData.Leafs.Append("virtual-router-id-xr", types.YLeaf{"VirtualRouterIdXr", trackItemInfo.VirtualRouterIdXr})
    trackItemInfo.EntityData.Leafs.Append("tracked-item-type", types.YLeaf{"TrackedItemType", trackItemInfo.TrackedItemType})
    trackItemInfo.EntityData.Leafs.Append("tracked-item-index", types.YLeaf{"TrackedItemIndex", trackItemInfo.TrackedItemIndex})
    trackItemInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", trackItemInfo.State})
    trackItemInfo.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", trackItemInfo.Priority})

    trackItemInfo.EntityData.YListKeys = []string {}

    return &(trackItemInfo.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory
// State change history
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Old State. The type is VrrpBagProtocolState.
    OldState interface{}

    // New State. The type is VrrpBagProtocolState.
    NewState interface{}

    // Reason for state change. The type is VrrpStateChangeReason.
    Reason interface{}

    // Time of state change.
    Time Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time
}

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetEntityData() *types.CommonEntityData {
    stateChangeHistory.EntityData.YFilter = stateChangeHistory.YFilter
    stateChangeHistory.EntityData.YangName = "state-change-history"
    stateChangeHistory.EntityData.BundleName = "cisco_ios_xr"
    stateChangeHistory.EntityData.ParentYangName = "virtual-router"
    stateChangeHistory.EntityData.SegmentPath = "state-change-history"
    stateChangeHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateChangeHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateChangeHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateChangeHistory.EntityData.Children = types.NewOrderedMap()
    stateChangeHistory.EntityData.Children.Append("time", types.YChild{"Time", &stateChangeHistory.Time})
    stateChangeHistory.EntityData.Leafs = types.NewOrderedMap()
    stateChangeHistory.EntityData.Leafs.Append("old-state", types.YLeaf{"OldState", stateChangeHistory.OldState})
    stateChangeHistory.EntityData.Leafs.Append("new-state", types.YLeaf{"NewState", stateChangeHistory.NewState})
    stateChangeHistory.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", stateChangeHistory.Reason})

    stateChangeHistory.EntityData.YListKeys = []string {}

    return &(stateChangeHistory.EntityData)
}

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time
// Time of state change
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetEntityData() *types.CommonEntityData {
    time.EntityData.YFilter = time.YFilter
    time.EntityData.YangName = "time"
    time.EntityData.BundleName = "cisco_ios_xr"
    time.EntityData.ParentYangName = "state-change-history"
    time.EntityData.SegmentPath = "time"
    time.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    time.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    time.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    time.EntityData.Children = types.NewOrderedMap()
    time.EntityData.Leafs = types.NewOrderedMap()
    time.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", time.Seconds})
    time.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", time.Nanoseconds})

    time.EntityData.YListKeys = []string {}

    return &(time.EntityData)
}

// Vrrp_Ipv6_Interfaces
// The VRRP interface table
type Vrrp_Ipv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP interface entry. The type is slice of
    // Vrrp_Ipv6_Interfaces_Interface.
    Interface []*Vrrp_Ipv6_Interfaces_Interface
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv6"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Vrrp_Ipv6_Interfaces_Interface
// A VRRP interface entry
type Vrrp_Ipv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Invalid checksum. The type is interface{} with range: 0..4294967295.
    InvalidChecksumCount interface{}

    // Unknown/unsupported version. The type is interface{} with range:
    // 0..4294967295.
    InvalidVersionCount interface{}

    // Invalid vrID. The type is interface{} with range: 0..4294967295.
    InvalidVridCount interface{}

    // Bad packet lengths. The type is interface{} with range: 0..4294967295.
    InvalidPacketLengthCount interface{}
}

func (self *Vrrp_Ipv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("invalid-checksum-count", types.YLeaf{"InvalidChecksumCount", self.InvalidChecksumCount})
    self.EntityData.Leafs.Append("invalid-version-count", types.YLeaf{"InvalidVersionCount", self.InvalidVersionCount})
    self.EntityData.Leafs.Append("invalid-vrid-count", types.YLeaf{"InvalidVridCount", self.InvalidVridCount})
    self.EntityData.Leafs.Append("invalid-packet-length-count", types.YLeaf{"InvalidPacketLengthCount", self.InvalidPacketLengthCount})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Vrrp_Ipv4
// IPv4 VRRP configuration
type Vrrp_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP interface table.
    Interfaces Vrrp_Ipv4_Interfaces

    // The VRRP tracked item table.
    TrackItems Vrrp_Ipv4_TrackItems

    // The VRRP virtual router table.
    VirtualRouters Vrrp_Ipv4_VirtualRouters
}

func (ipv4 *Vrrp_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrrp"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv4.Interfaces})
    ipv4.EntityData.Children.Append("track-items", types.YChild{"TrackItems", &ipv4.TrackItems})
    ipv4.EntityData.Children.Append("virtual-routers", types.YChild{"VirtualRouters", &ipv4.VirtualRouters})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Vrrp_Ipv4_Interfaces
// The VRRP interface table
type Vrrp_Ipv4_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP interface entry. The type is slice of
    // Vrrp_Ipv4_Interfaces_Interface.
    Interface []*Vrrp_Ipv4_Interfaces_Interface
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Vrrp_Ipv4_Interfaces_Interface
// A VRRP interface entry
type Vrrp_Ipv4_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Invalid checksum. The type is interface{} with range: 0..4294967295.
    InvalidChecksumCount interface{}

    // Unknown/unsupported version. The type is interface{} with range:
    // 0..4294967295.
    InvalidVersionCount interface{}

    // Invalid vrID. The type is interface{} with range: 0..4294967295.
    InvalidVridCount interface{}

    // Bad packet lengths. The type is interface{} with range: 0..4294967295.
    InvalidPacketLengthCount interface{}
}

func (self *Vrrp_Ipv4_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("invalid-checksum-count", types.YLeaf{"InvalidChecksumCount", self.InvalidChecksumCount})
    self.EntityData.Leafs.Append("invalid-version-count", types.YLeaf{"InvalidVersionCount", self.InvalidVersionCount})
    self.EntityData.Leafs.Append("invalid-vrid-count", types.YLeaf{"InvalidVridCount", self.InvalidVridCount})
    self.EntityData.Leafs.Append("invalid-packet-length-count", types.YLeaf{"InvalidPacketLengthCount", self.InvalidPacketLengthCount})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Vrrp_Ipv4_TrackItems
// The VRRP tracked item table
type Vrrp_Ipv4_TrackItems struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A configured VRRP IP address entry. The type is slice of
    // Vrrp_Ipv4_TrackItems_TrackItem.
    TrackItem []*Vrrp_Ipv4_TrackItems_TrackItem
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetEntityData() *types.CommonEntityData {
    trackItems.EntityData.YFilter = trackItems.YFilter
    trackItems.EntityData.YangName = "track-items"
    trackItems.EntityData.BundleName = "cisco_ios_xr"
    trackItems.EntityData.ParentYangName = "ipv4"
    trackItems.EntityData.SegmentPath = "track-items"
    trackItems.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackItems.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackItems.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackItems.EntityData.Children = types.NewOrderedMap()
    trackItems.EntityData.Children.Append("track-item", types.YChild{"TrackItem", nil})
    for i := range trackItems.TrackItem {
        trackItems.EntityData.Children.Append(types.GetSegmentPath(trackItems.TrackItem[i]), types.YChild{"TrackItem", trackItems.TrackItem[i]})
    }
    trackItems.EntityData.Leafs = types.NewOrderedMap()

    trackItems.EntityData.YListKeys = []string {}

    return &(trackItems.EntityData)
}

// Vrrp_Ipv4_TrackItems_TrackItem
// A configured VRRP IP address entry
type Vrrp_Ipv4_TrackItems_TrackItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name to track. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: 0..4294967295.
    VirtualRouterId interface{}

    // This attribute is a key. The name of the tracked interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    TrackedInterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    VirtualRouterIdXr interface{}

    // Type of tracked item. The type is interface{} with range: 0..65535.
    TrackedItemType interface{}

    // Tracked item index. The type is string with length: 0..32.
    TrackedItemIndex interface{}

    // State of the tracked item. The type is interface{} with range: 0..255.
    State interface{}

    // Priority weight of item. The type is interface{} with range: 0..255.
    Priority interface{}
}

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetEntityData() *types.CommonEntityData {
    trackItem.EntityData.YFilter = trackItem.YFilter
    trackItem.EntityData.YangName = "track-item"
    trackItem.EntityData.BundleName = "cisco_ios_xr"
    trackItem.EntityData.ParentYangName = "track-items"
    trackItem.EntityData.SegmentPath = "track-item" + types.AddKeyToken(trackItem.InterfaceName, "interface-name") + types.AddKeyToken(trackItem.VirtualRouterId, "virtual-router-id") + types.AddKeyToken(trackItem.TrackedInterfaceName, "tracked-interface-name")
    trackItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackItem.EntityData.Children = types.NewOrderedMap()
    trackItem.EntityData.Leafs = types.NewOrderedMap()
    trackItem.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", trackItem.InterfaceName})
    trackItem.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", trackItem.VirtualRouterId})
    trackItem.EntityData.Leafs.Append("tracked-interface-name", types.YLeaf{"TrackedInterfaceName", trackItem.TrackedInterfaceName})
    trackItem.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trackItem.Interface})
    trackItem.EntityData.Leafs.Append("virtual-router-id-xr", types.YLeaf{"VirtualRouterIdXr", trackItem.VirtualRouterIdXr})
    trackItem.EntityData.Leafs.Append("tracked-item-type", types.YLeaf{"TrackedItemType", trackItem.TrackedItemType})
    trackItem.EntityData.Leafs.Append("tracked-item-index", types.YLeaf{"TrackedItemIndex", trackItem.TrackedItemIndex})
    trackItem.EntityData.Leafs.Append("state", types.YLeaf{"State", trackItem.State})
    trackItem.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", trackItem.Priority})

    trackItem.EntityData.YListKeys = []string {"InterfaceName", "VirtualRouterId", "TrackedInterfaceName"}

    return &(trackItem.EntityData)
}

// Vrrp_Ipv4_VirtualRouters
// The VRRP virtual router table
type Vrrp_Ipv4_VirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP virtual router. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter.
    VirtualRouter []*Vrrp_Ipv4_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetEntityData() *types.CommonEntityData {
    virtualRouters.EntityData.YFilter = virtualRouters.YFilter
    virtualRouters.EntityData.YangName = "virtual-routers"
    virtualRouters.EntityData.BundleName = "cisco_ios_xr"
    virtualRouters.EntityData.ParentYangName = "ipv4"
    virtualRouters.EntityData.SegmentPath = "virtual-routers"
    virtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouters.EntityData.Children = types.NewOrderedMap()
    virtualRouters.EntityData.Children.Append("virtual-router", types.YChild{"VirtualRouter", nil})
    for i := range virtualRouters.VirtualRouter {
        virtualRouters.EntityData.Children.Append(types.GetSegmentPath(virtualRouters.VirtualRouter[i]), types.YChild{"VirtualRouter", virtualRouters.VirtualRouter[i]})
    }
    virtualRouters.EntityData.Leafs = types.NewOrderedMap()

    virtualRouters.EntityData.YListKeys = []string {}

    return &(virtualRouters.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter
// A VRRP virtual router
type Vrrp_Ipv4_VirtualRouters_VirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: 0..4294967295.
    VirtualRouterId interface{}

    // IM Interface Name. The type is string with length: 0..64.
    InterfaceNameXr interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    VirtualRouterIdXr interface{}

    // VRRP Protocol Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Address family. The type is VrrpBAf.
    AddressFamily interface{}

    // Session Name. The type is string with length: 0..16.
    SessionName interface{}

    // Number of slaves following state. The type is interface{} with range:
    // 0..4294967295.
    Slaves interface{}

    // Group is a slave group. The type is bool.
    IsSlave interface{}

    // Followed Session Name. The type is string with length: 0..16.
    FollowedSessionName interface{}

    // Configured VRRP secondary address count. The type is interface{} with
    // range: 0..255.
    SecondaryAddressCount interface{}

    // Operational VRRP address count. The type is interface{} with range: 0..255.
    OperationalAddressCount interface{}

    // Configured IPv4 Primary address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryVirtualIp interface{}

    // Configured but Down VRRP address count. The type is interface{} with range:
    // 0..255.
    ConfiguredDownAddressCount interface{}

    // Virtual linklocal IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinklocalIpv6Address interface{}

    // State of primary IP address. The type is VrrpVipState.
    PrimaryState interface{}

    // Master router real IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    MasterIpAddress interface{}

    // Master router real IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    MasterIpv6Address interface{}

    // Master router priority. The type is interface{} with range: 0..255.
    MasterPriority interface{}

    // VRRP state. The type is VrrpBagProtocolState.
    VrrpState interface{}

    // Authentication type. The type is VrrpProtAuth.
    AuthenticationType interface{}

    // Authentication data. The type is string.
    AuthenticationString interface{}

    // Configured advertize time. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredAdvertizeTime interface{}

    // Operational advertize time. The type is interface{} with range:
    // 0..4294967295.
    OperAdvertizeTime interface{}

    // Minimum delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    MinDelayTime interface{}

    // Reload delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ReloadDelayTime interface{}

    // Delay timer running flag. The type is bool.
    DelayTimerFlag interface{}

    // Delay timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DelayTimerSecs interface{}

    // Delay timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    DelayTimerMsecs interface{}

    // Text authentication configured flag. The type is bool.
    AuthenticationFlag interface{}

    // Configured timers forced flag. The type is bool.
    ForceTimerFlag interface{}

    // Preempt configured flag. The type is bool.
    PreemptFlag interface{}

    // IP address owner flag. The type is bool.
    IpAddressOwnerFlag interface{}

    // Is accept mode. The type is bool.
    IsAcceptMode interface{}

    // Preempt delay time. The type is interface{} with range: 0..65535.
    PreemptDelayTime interface{}

    // Configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // Operational priority. The type is interface{} with range: 0..255.
    OperationalPriority interface{}

    // Priority decrement. The type is interface{} with range: 0..4294967295.
    PriorityDecrement interface{}

    // Number of items tracked. The type is interface{} with range: 0..4294967295.
    TrackedInterfaceCount interface{}

    // Number of tracked items up. The type is interface{} with range:
    // 0..4294967295.
    TrackedInterfaceUpCount interface{}

    // Number of tracked items. The type is interface{} with range: 0..4294967295.
    TrackedItemCount interface{}

    // Number of tracked items in UP state. The type is interface{} with range:
    // 0..4294967295.
    TrackedItemUpCount interface{}

    // Time in current state secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TimeInCurrentState interface{}

    // Number of state changes. The type is interface{} with range: 0..4294967295.
    StateChangeCount interface{}

    // Time vrouter is up in centiseconds. The type is interface{} with range:
    // 0..4294967295. Units are centisecond.
    TimeVrouterUp interface{}

    // No. of times become Master. The type is interface{} with range:
    // 0..4294967295.
    MasterCount interface{}

    // No. of advertisements received. The type is interface{} with range:
    // 0..4294967295.
    AdvertsReceivedCount interface{}

    // Advertise interval errors. The type is interface{} with range:
    // 0..4294967295.
    AdvertIntervalErrorCount interface{}

    // No. of advertisements sent. The type is interface{} with range:
    // 0..4294967295.
    AdvertsSentCount interface{}

    // Authentication failures. The type is interface{} with range: 0..4294967295.
    AuthenticationFailCount interface{}

    // TTL errors. The type is interface{} with range: 0..4294967295.
    TtlErrorCount interface{}

    // No. priority 0 received. The type is interface{} with range: 0..4294967295.
    PriorityZeroReceivedCount interface{}

    // No. priority 0 sent. The type is interface{} with range: 0..4294967295.
    PriorityZeroSentCount interface{}

    // Invalid packets received. The type is interface{} with range:
    // 0..4294967295.
    InvalidPacketCount interface{}

    // Address list errors. The type is interface{} with range: 0..4294967295.
    AddressListErrorCount interface{}

    // Invalid authentication type. The type is interface{} with range:
    // 0..4294967295.
    InvalidAuthTypeCount interface{}

    // Authentication type mismatches. The type is interface{} with range:
    // 0..4294967295.
    AuthTypeMismatchCount interface{}

    // Packet length errors. The type is interface{} with range: 0..4294967295.
    PktLengthErrorsCount interface{}

    // Time since a statistics discontinuity in ticks (10ns units). The type is
    // interface{} with range: 0..4294967295.
    TimeStatsDiscontinuity interface{}

    // BFD session state. The type is VrrpBfdSessionState.
    BfdSessionState interface{}

    // BFD packet send interval. The type is interface{} with range:
    // 0..4294967295.
    BfdInterval interface{}

    // BFD multiplier. The type is interface{} with range: 0..4294967295.
    BfdMultiplier interface{}

    // BFD configured remote IP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BfdCfgRemoteIp interface{}

    // BFD configured remote IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    BfdConfiguredRemoteIpv6Address interface{}

    // Whether state recovered from checkpoint. The type is bool.
    StateFromCheckpoint interface{}

    // The Interface Primary IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    InterfaceIpv4Address interface{}

    // The Interface linklocal IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    InterfaceIpv6Address interface{}

    // Virtual mac address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // Virtual mac address state. The type is VrrpVmacState.
    VirtualMacAddressState interface{}

    // Operational IPv4 VRRP addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OperationalAddress []interface{}

    // IPv4 Configured but Down VRRP addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4ConfiguredDownAddress []interface{}

    // Time last resign was sent.
    ResignSentTime Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime

    // Time last resign was received.
    ResignReceivedTime Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime

    // IPv6 Operational VRRP addresses. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress.
    Ipv6OperationalAddress []*Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress

    // IPv6 Configured but Down VRRP addresses. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress.
    Ipv6ConfiguredDownAddress []*Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress

    // Track Item Info. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter_TrackItemInfo.
    TrackItemInfo []*Vrrp_Ipv4_VirtualRouters_VirtualRouter_TrackItemInfo

    // State change history. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory.
    StateChangeHistory []*Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetEntityData() *types.CommonEntityData {
    virtualRouter.EntityData.YFilter = virtualRouter.YFilter
    virtualRouter.EntityData.YangName = "virtual-router"
    virtualRouter.EntityData.BundleName = "cisco_ios_xr"
    virtualRouter.EntityData.ParentYangName = "virtual-routers"
    virtualRouter.EntityData.SegmentPath = "virtual-router" + types.AddKeyToken(virtualRouter.InterfaceName, "interface-name") + types.AddKeyToken(virtualRouter.VirtualRouterId, "virtual-router-id")
    virtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouter.EntityData.Children = types.NewOrderedMap()
    virtualRouter.EntityData.Children.Append("resign-sent-time", types.YChild{"ResignSentTime", &virtualRouter.ResignSentTime})
    virtualRouter.EntityData.Children.Append("resign-received-time", types.YChild{"ResignReceivedTime", &virtualRouter.ResignReceivedTime})
    virtualRouter.EntityData.Children.Append("ipv6-operational-address", types.YChild{"Ipv6OperationalAddress", nil})
    for i := range virtualRouter.Ipv6OperationalAddress {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.Ipv6OperationalAddress[i]), types.YChild{"Ipv6OperationalAddress", virtualRouter.Ipv6OperationalAddress[i]})
    }
    virtualRouter.EntityData.Children.Append("ipv6-configured-down-address", types.YChild{"Ipv6ConfiguredDownAddress", nil})
    for i := range virtualRouter.Ipv6ConfiguredDownAddress {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.Ipv6ConfiguredDownAddress[i]), types.YChild{"Ipv6ConfiguredDownAddress", virtualRouter.Ipv6ConfiguredDownAddress[i]})
    }
    virtualRouter.EntityData.Children.Append("track-item-info", types.YChild{"TrackItemInfo", nil})
    for i := range virtualRouter.TrackItemInfo {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.TrackItemInfo[i]), types.YChild{"TrackItemInfo", virtualRouter.TrackItemInfo[i]})
    }
    virtualRouter.EntityData.Children.Append("state-change-history", types.YChild{"StateChangeHistory", nil})
    for i := range virtualRouter.StateChangeHistory {
        virtualRouter.EntityData.Children.Append(types.GetSegmentPath(virtualRouter.StateChangeHistory[i]), types.YChild{"StateChangeHistory", virtualRouter.StateChangeHistory[i]})
    }
    virtualRouter.EntityData.Leafs = types.NewOrderedMap()
    virtualRouter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", virtualRouter.InterfaceName})
    virtualRouter.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", virtualRouter.VirtualRouterId})
    virtualRouter.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", virtualRouter.InterfaceNameXr})
    virtualRouter.EntityData.Leafs.Append("virtual-router-id-xr", types.YLeaf{"VirtualRouterIdXr", virtualRouter.VirtualRouterIdXr})
    virtualRouter.EntityData.Leafs.Append("version", types.YLeaf{"Version", virtualRouter.Version})
    virtualRouter.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", virtualRouter.AddressFamily})
    virtualRouter.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", virtualRouter.SessionName})
    virtualRouter.EntityData.Leafs.Append("slaves", types.YLeaf{"Slaves", virtualRouter.Slaves})
    virtualRouter.EntityData.Leafs.Append("is-slave", types.YLeaf{"IsSlave", virtualRouter.IsSlave})
    virtualRouter.EntityData.Leafs.Append("followed-session-name", types.YLeaf{"FollowedSessionName", virtualRouter.FollowedSessionName})
    virtualRouter.EntityData.Leafs.Append("secondary-address-count", types.YLeaf{"SecondaryAddressCount", virtualRouter.SecondaryAddressCount})
    virtualRouter.EntityData.Leafs.Append("operational-address-count", types.YLeaf{"OperationalAddressCount", virtualRouter.OperationalAddressCount})
    virtualRouter.EntityData.Leafs.Append("primary-virtual-ip", types.YLeaf{"PrimaryVirtualIp", virtualRouter.PrimaryVirtualIp})
    virtualRouter.EntityData.Leafs.Append("configured-down-address-count", types.YLeaf{"ConfiguredDownAddressCount", virtualRouter.ConfiguredDownAddressCount})
    virtualRouter.EntityData.Leafs.Append("virtual-linklocal-ipv6-address", types.YLeaf{"VirtualLinklocalIpv6Address", virtualRouter.VirtualLinklocalIpv6Address})
    virtualRouter.EntityData.Leafs.Append("primary-state", types.YLeaf{"PrimaryState", virtualRouter.PrimaryState})
    virtualRouter.EntityData.Leafs.Append("master-ip-address", types.YLeaf{"MasterIpAddress", virtualRouter.MasterIpAddress})
    virtualRouter.EntityData.Leafs.Append("master-ipv6-address", types.YLeaf{"MasterIpv6Address", virtualRouter.MasterIpv6Address})
    virtualRouter.EntityData.Leafs.Append("master-priority", types.YLeaf{"MasterPriority", virtualRouter.MasterPriority})
    virtualRouter.EntityData.Leafs.Append("vrrp-state", types.YLeaf{"VrrpState", virtualRouter.VrrpState})
    virtualRouter.EntityData.Leafs.Append("authentication-type", types.YLeaf{"AuthenticationType", virtualRouter.AuthenticationType})
    virtualRouter.EntityData.Leafs.Append("authentication-string", types.YLeaf{"AuthenticationString", virtualRouter.AuthenticationString})
    virtualRouter.EntityData.Leafs.Append("configured-advertize-time", types.YLeaf{"ConfiguredAdvertizeTime", virtualRouter.ConfiguredAdvertizeTime})
    virtualRouter.EntityData.Leafs.Append("oper-advertize-time", types.YLeaf{"OperAdvertizeTime", virtualRouter.OperAdvertizeTime})
    virtualRouter.EntityData.Leafs.Append("min-delay-time", types.YLeaf{"MinDelayTime", virtualRouter.MinDelayTime})
    virtualRouter.EntityData.Leafs.Append("reload-delay-time", types.YLeaf{"ReloadDelayTime", virtualRouter.ReloadDelayTime})
    virtualRouter.EntityData.Leafs.Append("delay-timer-flag", types.YLeaf{"DelayTimerFlag", virtualRouter.DelayTimerFlag})
    virtualRouter.EntityData.Leafs.Append("delay-timer-secs", types.YLeaf{"DelayTimerSecs", virtualRouter.DelayTimerSecs})
    virtualRouter.EntityData.Leafs.Append("delay-timer-msecs", types.YLeaf{"DelayTimerMsecs", virtualRouter.DelayTimerMsecs})
    virtualRouter.EntityData.Leafs.Append("authentication-flag", types.YLeaf{"AuthenticationFlag", virtualRouter.AuthenticationFlag})
    virtualRouter.EntityData.Leafs.Append("force-timer-flag", types.YLeaf{"ForceTimerFlag", virtualRouter.ForceTimerFlag})
    virtualRouter.EntityData.Leafs.Append("preempt-flag", types.YLeaf{"PreemptFlag", virtualRouter.PreemptFlag})
    virtualRouter.EntityData.Leafs.Append("ip-address-owner-flag", types.YLeaf{"IpAddressOwnerFlag", virtualRouter.IpAddressOwnerFlag})
    virtualRouter.EntityData.Leafs.Append("is-accept-mode", types.YLeaf{"IsAcceptMode", virtualRouter.IsAcceptMode})
    virtualRouter.EntityData.Leafs.Append("preempt-delay-time", types.YLeaf{"PreemptDelayTime", virtualRouter.PreemptDelayTime})
    virtualRouter.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", virtualRouter.ConfiguredPriority})
    virtualRouter.EntityData.Leafs.Append("operational-priority", types.YLeaf{"OperationalPriority", virtualRouter.OperationalPriority})
    virtualRouter.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", virtualRouter.PriorityDecrement})
    virtualRouter.EntityData.Leafs.Append("tracked-interface-count", types.YLeaf{"TrackedInterfaceCount", virtualRouter.TrackedInterfaceCount})
    virtualRouter.EntityData.Leafs.Append("tracked-interface-up-count", types.YLeaf{"TrackedInterfaceUpCount", virtualRouter.TrackedInterfaceUpCount})
    virtualRouter.EntityData.Leafs.Append("tracked-item-count", types.YLeaf{"TrackedItemCount", virtualRouter.TrackedItemCount})
    virtualRouter.EntityData.Leafs.Append("tracked-item-up-count", types.YLeaf{"TrackedItemUpCount", virtualRouter.TrackedItemUpCount})
    virtualRouter.EntityData.Leafs.Append("time-in-current-state", types.YLeaf{"TimeInCurrentState", virtualRouter.TimeInCurrentState})
    virtualRouter.EntityData.Leafs.Append("state-change-count", types.YLeaf{"StateChangeCount", virtualRouter.StateChangeCount})
    virtualRouter.EntityData.Leafs.Append("time-vrouter-up", types.YLeaf{"TimeVrouterUp", virtualRouter.TimeVrouterUp})
    virtualRouter.EntityData.Leafs.Append("master-count", types.YLeaf{"MasterCount", virtualRouter.MasterCount})
    virtualRouter.EntityData.Leafs.Append("adverts-received-count", types.YLeaf{"AdvertsReceivedCount", virtualRouter.AdvertsReceivedCount})
    virtualRouter.EntityData.Leafs.Append("advert-interval-error-count", types.YLeaf{"AdvertIntervalErrorCount", virtualRouter.AdvertIntervalErrorCount})
    virtualRouter.EntityData.Leafs.Append("adverts-sent-count", types.YLeaf{"AdvertsSentCount", virtualRouter.AdvertsSentCount})
    virtualRouter.EntityData.Leafs.Append("authentication-fail-count", types.YLeaf{"AuthenticationFailCount", virtualRouter.AuthenticationFailCount})
    virtualRouter.EntityData.Leafs.Append("ttl-error-count", types.YLeaf{"TtlErrorCount", virtualRouter.TtlErrorCount})
    virtualRouter.EntityData.Leafs.Append("priority-zero-received-count", types.YLeaf{"PriorityZeroReceivedCount", virtualRouter.PriorityZeroReceivedCount})
    virtualRouter.EntityData.Leafs.Append("priority-zero-sent-count", types.YLeaf{"PriorityZeroSentCount", virtualRouter.PriorityZeroSentCount})
    virtualRouter.EntityData.Leafs.Append("invalid-packet-count", types.YLeaf{"InvalidPacketCount", virtualRouter.InvalidPacketCount})
    virtualRouter.EntityData.Leafs.Append("address-list-error-count", types.YLeaf{"AddressListErrorCount", virtualRouter.AddressListErrorCount})
    virtualRouter.EntityData.Leafs.Append("invalid-auth-type-count", types.YLeaf{"InvalidAuthTypeCount", virtualRouter.InvalidAuthTypeCount})
    virtualRouter.EntityData.Leafs.Append("auth-type-mismatch-count", types.YLeaf{"AuthTypeMismatchCount", virtualRouter.AuthTypeMismatchCount})
    virtualRouter.EntityData.Leafs.Append("pkt-length-errors-count", types.YLeaf{"PktLengthErrorsCount", virtualRouter.PktLengthErrorsCount})
    virtualRouter.EntityData.Leafs.Append("time-stats-discontinuity", types.YLeaf{"TimeStatsDiscontinuity", virtualRouter.TimeStatsDiscontinuity})
    virtualRouter.EntityData.Leafs.Append("bfd-session-state", types.YLeaf{"BfdSessionState", virtualRouter.BfdSessionState})
    virtualRouter.EntityData.Leafs.Append("bfd-interval", types.YLeaf{"BfdInterval", virtualRouter.BfdInterval})
    virtualRouter.EntityData.Leafs.Append("bfd-multiplier", types.YLeaf{"BfdMultiplier", virtualRouter.BfdMultiplier})
    virtualRouter.EntityData.Leafs.Append("bfd-cfg-remote-ip", types.YLeaf{"BfdCfgRemoteIp", virtualRouter.BfdCfgRemoteIp})
    virtualRouter.EntityData.Leafs.Append("bfd-configured-remote-ipv6-address", types.YLeaf{"BfdConfiguredRemoteIpv6Address", virtualRouter.BfdConfiguredRemoteIpv6Address})
    virtualRouter.EntityData.Leafs.Append("state-from-checkpoint", types.YLeaf{"StateFromCheckpoint", virtualRouter.StateFromCheckpoint})
    virtualRouter.EntityData.Leafs.Append("interface-ipv4-address", types.YLeaf{"InterfaceIpv4Address", virtualRouter.InterfaceIpv4Address})
    virtualRouter.EntityData.Leafs.Append("interface-ipv6-address", types.YLeaf{"InterfaceIpv6Address", virtualRouter.InterfaceIpv6Address})
    virtualRouter.EntityData.Leafs.Append("virtual-mac-address", types.YLeaf{"VirtualMacAddress", virtualRouter.VirtualMacAddress})
    virtualRouter.EntityData.Leafs.Append("virtual-mac-address-state", types.YLeaf{"VirtualMacAddressState", virtualRouter.VirtualMacAddressState})
    virtualRouter.EntityData.Leafs.Append("operational-address", types.YLeaf{"OperationalAddress", virtualRouter.OperationalAddress})
    virtualRouter.EntityData.Leafs.Append("ipv4-configured-down-address", types.YLeaf{"Ipv4ConfiguredDownAddress", virtualRouter.Ipv4ConfiguredDownAddress})

    virtualRouter.EntityData.YListKeys = []string {"InterfaceName", "VirtualRouterId"}

    return &(virtualRouter.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime
// Time last resign was sent
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetEntityData() *types.CommonEntityData {
    resignSentTime.EntityData.YFilter = resignSentTime.YFilter
    resignSentTime.EntityData.YangName = "resign-sent-time"
    resignSentTime.EntityData.BundleName = "cisco_ios_xr"
    resignSentTime.EntityData.ParentYangName = "virtual-router"
    resignSentTime.EntityData.SegmentPath = "resign-sent-time"
    resignSentTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resignSentTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resignSentTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resignSentTime.EntityData.Children = types.NewOrderedMap()
    resignSentTime.EntityData.Leafs = types.NewOrderedMap()
    resignSentTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", resignSentTime.Seconds})
    resignSentTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", resignSentTime.Nanoseconds})

    resignSentTime.EntityData.YListKeys = []string {}

    return &(resignSentTime.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime
// Time last resign was received
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetEntityData() *types.CommonEntityData {
    resignReceivedTime.EntityData.YFilter = resignReceivedTime.YFilter
    resignReceivedTime.EntityData.YangName = "resign-received-time"
    resignReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    resignReceivedTime.EntityData.ParentYangName = "virtual-router"
    resignReceivedTime.EntityData.SegmentPath = "resign-received-time"
    resignReceivedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resignReceivedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resignReceivedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resignReceivedTime.EntityData.Children = types.NewOrderedMap()
    resignReceivedTime.EntityData.Leafs = types.NewOrderedMap()
    resignReceivedTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", resignReceivedTime.Seconds})
    resignReceivedTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", resignReceivedTime.Nanoseconds})

    resignReceivedTime.EntityData.YListKeys = []string {}

    return &(resignReceivedTime.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress
// IPv6 Operational VRRP addresses
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetEntityData() *types.CommonEntityData {
    ipv6OperationalAddress.EntityData.YFilter = ipv6OperationalAddress.YFilter
    ipv6OperationalAddress.EntityData.YangName = "ipv6-operational-address"
    ipv6OperationalAddress.EntityData.BundleName = "cisco_ios_xr"
    ipv6OperationalAddress.EntityData.ParentYangName = "virtual-router"
    ipv6OperationalAddress.EntityData.SegmentPath = "ipv6-operational-address"
    ipv6OperationalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6OperationalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6OperationalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6OperationalAddress.EntityData.Children = types.NewOrderedMap()
    ipv6OperationalAddress.EntityData.Leafs = types.NewOrderedMap()
    ipv6OperationalAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6OperationalAddress.Ipv6Address})

    ipv6OperationalAddress.EntityData.YListKeys = []string {}

    return &(ipv6OperationalAddress.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress
// IPv6 Configured but Down VRRP addresses
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetEntityData() *types.CommonEntityData {
    ipv6ConfiguredDownAddress.EntityData.YFilter = ipv6ConfiguredDownAddress.YFilter
    ipv6ConfiguredDownAddress.EntityData.YangName = "ipv6-configured-down-address"
    ipv6ConfiguredDownAddress.EntityData.BundleName = "cisco_ios_xr"
    ipv6ConfiguredDownAddress.EntityData.ParentYangName = "virtual-router"
    ipv6ConfiguredDownAddress.EntityData.SegmentPath = "ipv6-configured-down-address"
    ipv6ConfiguredDownAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6ConfiguredDownAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6ConfiguredDownAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6ConfiguredDownAddress.EntityData.Children = types.NewOrderedMap()
    ipv6ConfiguredDownAddress.EntityData.Leafs = types.NewOrderedMap()
    ipv6ConfiguredDownAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6ConfiguredDownAddress.Ipv6Address})

    ipv6ConfiguredDownAddress.EntityData.YListKeys = []string {}

    return &(ipv6ConfiguredDownAddress.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_TrackItemInfo
// Track Item Info
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_TrackItemInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Virtual Router ID. The type is interface{} with range: 0..4294967295.
    VirtualRouterIdXr interface{}

    // Type of tracked item. The type is interface{} with range: 0..65535.
    TrackedItemType interface{}

    // Tracked item index. The type is string with length: 0..32.
    TrackedItemIndex interface{}

    // State of the tracked item. The type is interface{} with range: 0..255.
    State interface{}

    // Priority weight of item. The type is interface{} with range: 0..255.
    Priority interface{}
}

func (trackItemInfo *Vrrp_Ipv4_VirtualRouters_VirtualRouter_TrackItemInfo) GetEntityData() *types.CommonEntityData {
    trackItemInfo.EntityData.YFilter = trackItemInfo.YFilter
    trackItemInfo.EntityData.YangName = "track-item-info"
    trackItemInfo.EntityData.BundleName = "cisco_ios_xr"
    trackItemInfo.EntityData.ParentYangName = "virtual-router"
    trackItemInfo.EntityData.SegmentPath = "track-item-info"
    trackItemInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackItemInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackItemInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackItemInfo.EntityData.Children = types.NewOrderedMap()
    trackItemInfo.EntityData.Leafs = types.NewOrderedMap()
    trackItemInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trackItemInfo.Interface})
    trackItemInfo.EntityData.Leafs.Append("virtual-router-id-xr", types.YLeaf{"VirtualRouterIdXr", trackItemInfo.VirtualRouterIdXr})
    trackItemInfo.EntityData.Leafs.Append("tracked-item-type", types.YLeaf{"TrackedItemType", trackItemInfo.TrackedItemType})
    trackItemInfo.EntityData.Leafs.Append("tracked-item-index", types.YLeaf{"TrackedItemIndex", trackItemInfo.TrackedItemIndex})
    trackItemInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", trackItemInfo.State})
    trackItemInfo.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", trackItemInfo.Priority})

    trackItemInfo.EntityData.YListKeys = []string {}

    return &(trackItemInfo.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory
// State change history
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Old State. The type is VrrpBagProtocolState.
    OldState interface{}

    // New State. The type is VrrpBagProtocolState.
    NewState interface{}

    // Reason for state change. The type is VrrpStateChangeReason.
    Reason interface{}

    // Time of state change.
    Time Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time
}

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetEntityData() *types.CommonEntityData {
    stateChangeHistory.EntityData.YFilter = stateChangeHistory.YFilter
    stateChangeHistory.EntityData.YangName = "state-change-history"
    stateChangeHistory.EntityData.BundleName = "cisco_ios_xr"
    stateChangeHistory.EntityData.ParentYangName = "virtual-router"
    stateChangeHistory.EntityData.SegmentPath = "state-change-history"
    stateChangeHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateChangeHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateChangeHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateChangeHistory.EntityData.Children = types.NewOrderedMap()
    stateChangeHistory.EntityData.Children.Append("time", types.YChild{"Time", &stateChangeHistory.Time})
    stateChangeHistory.EntityData.Leafs = types.NewOrderedMap()
    stateChangeHistory.EntityData.Leafs.Append("old-state", types.YLeaf{"OldState", stateChangeHistory.OldState})
    stateChangeHistory.EntityData.Leafs.Append("new-state", types.YLeaf{"NewState", stateChangeHistory.NewState})
    stateChangeHistory.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", stateChangeHistory.Reason})

    stateChangeHistory.EntityData.YListKeys = []string {}

    return &(stateChangeHistory.EntityData)
}

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time
// Time of state change
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetEntityData() *types.CommonEntityData {
    time.EntityData.YFilter = time.YFilter
    time.EntityData.YangName = "time"
    time.EntityData.BundleName = "cisco_ios_xr"
    time.EntityData.ParentYangName = "state-change-history"
    time.EntityData.SegmentPath = "time"
    time.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    time.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    time.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    time.EntityData.Children = types.NewOrderedMap()
    time.EntityData.Leafs = types.NewOrderedMap()
    time.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", time.Seconds})
    time.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", time.Nanoseconds})

    time.EntityData.YListKeys = []string {}

    return &(time.EntityData)
}

// Vrrp_MgoSessions
// VRRP MGO Session information
type Vrrp_MgoSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP MGO Session. The type is slice of Vrrp_MgoSessions_MgoSession.
    MgoSession []*Vrrp_MgoSessions_MgoSession
}

func (mgoSessions *Vrrp_MgoSessions) GetEntityData() *types.CommonEntityData {
    mgoSessions.EntityData.YFilter = mgoSessions.YFilter
    mgoSessions.EntityData.YangName = "mgo-sessions"
    mgoSessions.EntityData.BundleName = "cisco_ios_xr"
    mgoSessions.EntityData.ParentYangName = "vrrp"
    mgoSessions.EntityData.SegmentPath = "mgo-sessions"
    mgoSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mgoSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mgoSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mgoSessions.EntityData.Children = types.NewOrderedMap()
    mgoSessions.EntityData.Children.Append("mgo-session", types.YChild{"MgoSession", nil})
    for i := range mgoSessions.MgoSession {
        mgoSessions.EntityData.Children.Append(types.GetSegmentPath(mgoSessions.MgoSession[i]), types.YChild{"MgoSession", mgoSessions.MgoSession[i]})
    }
    mgoSessions.EntityData.Leafs = types.NewOrderedMap()

    mgoSessions.EntityData.YListKeys = []string {}

    return &(mgoSessions.EntityData)
}

// Vrrp_MgoSessions_MgoSession
// A VRRP MGO Session
type Vrrp_MgoSessions_MgoSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the session. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    SessionName interface{}

    // Session Name. The type is string with length: 0..16.
    PrimarySessionName interface{}

    // Interface of primary session. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    PrimarySessionInterface interface{}

    // Address family of primary session. The type is VrrpBAf.
    PrimaryAfName interface{}

    // VRID of primary session. The type is interface{} with range: 0..4294967295.
    PrimarySessionNumber interface{}

    // State of primary session. The type is VrrpBagProtocolState.
    PrimarySessionState interface{}

    // List of slaves following this primary session. The type is slice of
    // Vrrp_MgoSessions_MgoSession_Slave.
    Slave []*Vrrp_MgoSessions_MgoSession_Slave
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetEntityData() *types.CommonEntityData {
    mgoSession.EntityData.YFilter = mgoSession.YFilter
    mgoSession.EntityData.YangName = "mgo-session"
    mgoSession.EntityData.BundleName = "cisco_ios_xr"
    mgoSession.EntityData.ParentYangName = "mgo-sessions"
    mgoSession.EntityData.SegmentPath = "mgo-session" + types.AddKeyToken(mgoSession.SessionName, "session-name")
    mgoSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mgoSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mgoSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mgoSession.EntityData.Children = types.NewOrderedMap()
    mgoSession.EntityData.Children.Append("slave", types.YChild{"Slave", nil})
    for i := range mgoSession.Slave {
        mgoSession.EntityData.Children.Append(types.GetSegmentPath(mgoSession.Slave[i]), types.YChild{"Slave", mgoSession.Slave[i]})
    }
    mgoSession.EntityData.Leafs = types.NewOrderedMap()
    mgoSession.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", mgoSession.SessionName})
    mgoSession.EntityData.Leafs.Append("primary-session-name", types.YLeaf{"PrimarySessionName", mgoSession.PrimarySessionName})
    mgoSession.EntityData.Leafs.Append("primary-session-interface", types.YLeaf{"PrimarySessionInterface", mgoSession.PrimarySessionInterface})
    mgoSession.EntityData.Leafs.Append("primary-af-name", types.YLeaf{"PrimaryAfName", mgoSession.PrimaryAfName})
    mgoSession.EntityData.Leafs.Append("primary-session-number", types.YLeaf{"PrimarySessionNumber", mgoSession.PrimarySessionNumber})
    mgoSession.EntityData.Leafs.Append("primary-session-state", types.YLeaf{"PrimarySessionState", mgoSession.PrimarySessionState})

    mgoSession.EntityData.YListKeys = []string {"SessionName"}

    return &(mgoSession.EntityData)
}

// Vrrp_MgoSessions_MgoSession_Slave
// List of slaves following this primary session
type Vrrp_MgoSessions_MgoSession_Slave struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface of slave. The type is string with length: 0..64.
    SlaveInterface interface{}

    // VRID of slave. The type is interface{} with range: 0..4294967295.
    SlaveVirtualRouterId interface{}
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetEntityData() *types.CommonEntityData {
    slave.EntityData.YFilter = slave.YFilter
    slave.EntityData.YangName = "slave"
    slave.EntityData.BundleName = "cisco_ios_xr"
    slave.EntityData.ParentYangName = "mgo-session"
    slave.EntityData.SegmentPath = "slave"
    slave.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slave.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slave.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slave.EntityData.Children = types.NewOrderedMap()
    slave.EntityData.Leafs = types.NewOrderedMap()
    slave.EntityData.Leafs.Append("slave-interface", types.YLeaf{"SlaveInterface", slave.SlaveInterface})
    slave.EntityData.Leafs.Append("slave-virtual-router-id", types.YLeaf{"SlaveVirtualRouterId", slave.SlaveVirtualRouterId})

    slave.EntityData.YListKeys = []string {}

    return &(slave.EntityData)
}

